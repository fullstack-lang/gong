import { AfterViewInit, Component, Inject, Input, OnInit, Optional, ViewChild } from '@angular/core'
import { Subscription, debounceTime, distinctUntilChanged, forkJoin } from 'rxjs'
import { CdkDragDrop, moveItemInArray } from '@angular/cdk/drag-drop'

import * as table from '../../../../table/src/public-api' // Assuming this path is correct

import { MatTableDataSource } from '@angular/material/table'
import { FormControl } from '@angular/forms'

import { SelectionModel } from '@angular/cdk/collections'
import { TableDialogData } from '../table-dialog-data' // Assuming this path is correct

const allowMultiSelect = true

// Import MatDialogRef
import { MAT_DIALOG_DATA, MatDialog, MatDialogModule, MatDialogRef } from '@angular/material/dialog'
import { FormsModule, ReactiveFormsModule } from '@angular/forms'
import { MatSort, MatSortModule } from '@angular/material/sort'
import { DragDropModule } from '@angular/cdk/drag-drop'
import { CommonModule } from '@angular/common'

import { MatFormFieldModule } from '@angular/material/form-field'
import { MatPaginator, MatPaginatorModule } from '@angular/material/paginator'
import { MatTableModule } from '@angular/material/table'
import { MatCheckboxModule } from '@angular/material/checkbox'
import { MatInputModule } from '@angular/material/input'
import { MatButtonModule } from '@angular/material/button'
import { decodeStringToIntArray_json } from '../association-storage'


@Component({
  selector: 'lib-table-specific',
  // Ensure these are in the NgModule that declares this component if not using standalone components
  // For Angular 17+ standalone is the default, so 'imports' array is used directly here.
  imports: [
    FormsModule,
    ReactiveFormsModule,
    MatButtonModule,
    MatFormFieldModule,
    MatPaginatorModule,
    MatTableModule,
    MatCheckboxModule,
    MatDialogModule,
    MatInputModule,
    MatSortModule,
    CommonModule,
    DragDropModule,
  ],
  templateUrl: './table-specific.component.html',
  styleUrl: './table-specific.component.css'
})
export class TableSpecificComponent implements OnInit, AfterViewInit {

  displayedColumns: string[] = []
  allDisplayedColumns: string[] = [] // in case there is a checkbox

  mapHeaderIdIndex = new Map<string, number>()

  dataSource = new MatTableDataSource<table.Row>()

  stickyStyle = {
    position: 'sticky',
    left: '0',
    zIndex: '1'
  }
  // for selection
  selectedTable: table.Table | undefined = undefined

  @Input() Name: string = ""
  @Input() TableName: string = ""

  // for filtering
  filterControl = new FormControl()

  // for sorting
  @ViewChild(MatSort) sort!: MatSort // Use definite assignment assertion if sure it's always there after view init
  matSortDirective: string = ""

  // for pagination
  @ViewChild(MatPaginator)
  paginator!: MatPaginator // Use definite assignment assertion

  // The component is refreshed when modifications are performed in the back repo
  private commutNbFromBackSubscription: Subscription = new Subscription()
  lastCommitNbFromBack = -1
  lastPushFromFrontNb = -1
  currTime: number = 0
  dateOfLastTimerEmission: Date = new Date()


  public gongtableFrontRepo?: table.FrontRepo

  // SelectionModel to hold the selected rows
  selection: SelectionModel<table.Row> = new SelectionModel<table.Row>(allowMultiSelect, [])
  initialSelection = new Array<table.Row>()

  constructor(
    private gongtableFrontRepoService: table.FrontRepoService,
    private gongtableCommitNbFromBackService: table.CommitNbFromBackService,
    private rowService: table.RowService,
    private tableService: table.TableService,
    private celliconService: table.CellIconService,

    // MatDialog service for opening OTHER dialogs (if needed)
    public dialog: MatDialog,

    // Inject MatDialogRef for THIS dialog instance
    @Optional() public dialogRef: MatDialogRef<TableSpecificComponent>,

    // Data passed into this dialog
    @Optional() @Inject(MAT_DIALOG_DATA) public tableDialogData: TableDialogData,
  ) {
    // Initialize selection if needed, or it will be done in refresh
    // The following 'if' condition (line 112 from your error) caused a type error
    // because 'initialSelection' was not defined on the 'TableDialogData' type.
    // To fix this, ensure your 'TableDialogData' interface (in '../table-dialog-data.ts')
    // includes 'initialSelection?: table.Row[]' as an optional property.
    if (this.tableDialogData && this.tableDialogData.initialSelection) {
      // Assuming TableDialogData can carry an initial selection
      // this.initialSelection = this.tableDialogData.initialSelection
      // this.selection = new SelectionModel<table.Row>(allowMultiSelect, this.initialSelection)
    }
  }

  ngOnInit(): void {
    if (this.tableDialogData) {
      this.Name = this.tableDialogData.Name
      this.TableName = this.tableDialogData.TableName
    }

    this.refresh()

    this.filterControl.valueChanges
      .pipe(
        debounceTime(200),
        distinctUntilChanged()
      )
      .subscribe(value => {
        this.dataSource.filter = value ? value.trim().toLowerCase() : ''
      })
  }

  refresh(): void {
    this.gongtableFrontRepoService.connectToWebSocket(this.Name).subscribe(
      gongtablesFrontRepo => {
        this.gongtableFrontRepo = gongtablesFrontRepo
        this.selectedTable = undefined

        if (!this.gongtableFrontRepo) { // Guard against undefined front repo
          console.error("Gongtable front repo is not available.")
          return
        }

        let tableNames = new(Array<string>)
        for (let item of this.gongtableFrontRepo.getFrontArray<table.Table>(table.Table.GONGSTRUCT_NAME)) {
          tableNames.push(item.Name)
          if (item.Name == this.TableName) {
            this.selectedTable = item
            break // Found the table, no need to continue loop
          }
        }

        if (!this.selectedTable) {
          console.error(this.selectedTable, "not found among table names", tableNames)
          return
        }

        this.dataSource = new MatTableDataSource(this.selectedTable.Rows || [])

        this.mapHeaderIdIndex = new Map<string, number>()
        this.displayedColumns = []
        if (this.selectedTable.DisplayedColumns) {
          this.selectedTable.DisplayedColumns.forEach((column, index) => {
            if (column.Name) { // Ensure column name exists
              this.mapHeaderIdIndex.set(column.Name, index)
              this.displayedColumns.push(column.Name)
            }
          })
        }


        this.allDisplayedColumns = [...this.displayedColumns] // Initialize with displayed columns

        if (this.selectedTable.HasCheckableRows) {
          this.allDisplayedColumns.unshift('select') // Add 'select' column at the beginning

          this.initialSelection = []
          if (this.selectedTable.Rows) {
            this.selectedTable.Rows.forEach(Row => {
              if (Row.IsChecked) {
                this.initialSelection.push(Row)
              }
            })
          }

          if (this.tableDialogData && this.tableDialogData.AssociationStorage) {

            let sliceOfIDs = decodeStringToIntArray_json(this.tableDialogData.AssociationStorage)

            // parse all rows, get the ID (first cell), and if the ID is in the association storage,
            // push the row into the initial selection
            this.selectedTable.Rows.forEach(Row => {
              if (Row.Cells.length > 0 && Row.Cells[0].CellInt) {
                let id = Row.Cells[0].CellInt.Value

                if (sliceOfIDs.includes(id)) {
                  this.initialSelection.push(Row)
                }
              }
            })

          }
          // Re-initialize selection model with current data
          this.selection = new SelectionModel<table.Row>(allowMultiSelect, this.initialSelection)
        }


        if (this.selectedTable.HasFiltering) {
          this.dataSource.filterPredicate = (row: table.Row, filter: string) => {
            let mergedContent = ""
            if (row.Cells) {
              for (let cell of row.Cells) {
                if (cell.CellInt) mergedContent += cell.CellInt.Value
                if (cell.CellFloat64) mergedContent += cell.CellFloat64.Value
                if (cell.CellString) mergedContent += cell.CellString.Value
                // Add other cell types if necessary
              }
            }
            mergedContent = mergedContent.toLowerCase()
            return mergedContent.includes(filter.toLowerCase())
          }
        }

        this.matSortDirective = ""
        if (this.selectedTable.HasColumnSorting && this.sort) {
          this.dataSource.sort = this.sort
          this.matSortDirective = "mat-sort"

          this.dataSource.sortingDataAccessor = (row: table.Row, sortHeaderId: string): string | number => {
            if (!row.Cells) return ""
            const index = this.mapHeaderIdIndex.get(sortHeaderId)
            if (index === undefined) return ""

            const cell: table.Cell = row.Cells[index]
            if (cell.CellInt) return cell.CellInt.Value
            if (cell.CellFloat64) return cell.CellFloat64.Value
            if (cell.CellString) return cell.CellString.Value
            if (cell.CellIcon) return cell.CellIcon.Icon || "" // Ensure string for sorting
            if (cell.CellBool) return cell.CellBool.Value ? "true" : "false"
            return ""
          }
        }

        if (this.selectedTable.HasPaginator && this.paginator) {
          this.dataSource.paginator = this.paginator
        }
      }
    )
  }

  ngAfterViewInit() {
    // Ensure sort and paginator are assigned if they exist
    if (this.sort) {
      this.dataSource.sort = this.sort
    }
    if (this.paginator) {
      this.dataSource.paginator = this.paginator
    }
  }

  applyFilter(event: Event) {
    const filterValue = (event.target as HTMLInputElement).value
    this.dataSource.filter = filterValue.trim().toLowerCase()

    if (this.dataSource.paginator) {
      this.dataSource.paginator.firstPage()
    }
  }

  /** Whether the number of selected elements matches the total number of rows. */
  isAllSelected() {
    const numSelected = this.selection.selected.length
    const numRows = this.dataSource.data.length // Use dataSource.data for current rows after filter
    return numSelected === numRows && numRows > 0
  }

  /** Selects all rows if they are not all selected; otherwise clear selection. */
  masterToggle() {
    if (this.isAllSelected()) {
      this.selection.clear()
    } else {
      this.dataSource.data.forEach(row => this.selection.select(row)) // Select all rows in current view
    }
  }

  save(): void {
    if (!this.selectedTable || !this.selectedTable.Rows) {
      console.warn("Save called without a selected table or rows.")
      return
    }

    const modifiedRows = new Set<table.Row>()
    this.selectedTable.Rows.forEach(row => {
      const isCurrentlySelected = this.selection.isSelected(row)
      if (row.IsChecked !== isCurrentlySelected) {
        row.IsChecked = isCurrentlySelected
        modifiedRows.add(row)
      }
    })

    this.selectedTable.SavingInProgress = true // Indicate saving started

    if (modifiedRows.size === 0) {
      // Even if no rows changed selection status, we might want to save the table state
      // or simply close if that's the desired behavior.
      this.tableService.updateFront(this.selectedTable, this.Name).subscribe({
        next: () => {
          console.log('Table state saved (no selection changes).')
          if (this.selectedTable) this.selectedTable.SavingInProgress = false // Ensure selectedTable is defined
          if (this.tableDialogData && this.dialogRef) {
            // Pass the current selection (which might be empty or unchanged)
            this.dialogRef.close(this.selection.selected)
          }
        },
        error: (err) => {
          console.error('Error saving table state:', err)
          if (this.selectedTable) this.selectedTable.SavingInProgress = false // Ensure selectedTable is defined
        }
      })
      return
    }

    const promises = Array.from(modifiedRows).map(row =>
      this.rowService.updateFront(row, this.Name)
    )

    forkJoin(promises).subscribe({
      next: () => {
        console.log('Modified rows updated.')
        // After rows are updated, update the table itself
        if (!this.selectedTable) { // Guard against undefined selectedTable
          console.error("Selected table became undefined during save operation.")
          return
        }
        this.tableService.updateFront(this.selectedTable, this.Name).subscribe({
          next: () => {
            console.log('Table updated after row modifications.')
            if (this.selectedTable) this.selectedTable.SavingInProgress = false // Ensure selectedTable is defined
            if (this.tableDialogData && this.dialogRef) {
              // Pass the final selection state
              this.dialogRef.close(this.selection.selected)
            }
          },
          error: (err) => {
            console.error('Error updating table after row modifications:', err)
            if (this.selectedTable) this.selectedTable.SavingInProgress = false // Ensure selectedTable is defined
          }
        })
      },
      error: (err) => {
        console.error('Error updating one or more rows:', err)
        if (this.selectedTable) this.selectedTable.SavingInProgress = false // Ensure selectedTable is defined
      }
    })
  }

  drop(event: CdkDragDrop<string[]>) {
    if (!this.selectedTable || !this.selectedTable.Rows) return

    moveItemInArray(this.selectedTable.Rows, event.previousIndex, event.currentIndex)
    this.dataSource.data = [...this.selectedTable.Rows] // Update data source data

    this.tableService.updateFront(this.selectedTable, this.Name).subscribe(
      () => {
        console.log("Table rows shuffled and updated:", this.selectedTable?.Name)
      }
    )
  }

  isDraggableRow = (index: number, item: table.Row): boolean => !!this.selectedTable?.CanDragDropRows

  close(): void {
    if (this.tableDialogData && this.dialogRef) {
      this.dialogRef.close(this.selection.selected)
    } else if (this.dialogRef) {
      this.dialogRef.close()
    }
  }

  onClick(row: table.Row): void {
    console.log("Material Table: onClick: Stack: `" + this.Name + "`table:`" + this.TableName + "`row:" + (row.Name || 'Unnamed Row'))

    const originalCells = row.Cells
    this.rowService.updateFront(row, this.Name).subscribe(
      () => {
        console.log("Row updated on click:", row.Name || 'Unnamed Row')
      }
    )
  }

  getDynamicStyles(columnIndex: number): { [key: string]: any } {
    const styles: { [key: string]: any } = {}
    if (!this.selectedTable) {
      return styles
    }

    if (this.selectedTable.NbOfStickyColumns && columnIndex < this.selectedTable.NbOfStickyColumns) {
      styles['position'] = 'sticky'
      styles['left'] = '0px'
      styles['background'] = 'white'
      styles['z-index'] = '1'
    }
    return styles
  }

  onClickCellIcon(cellIcon: table.CellIcon): void {
    console.log("Cell Icon clicked:", cellIcon.Name || 'Unnamed Icon')
    this.celliconService.updateFront(cellIcon, this.Name).subscribe(
      () => {
        console.log("Cell icon updated:", cellIcon.Name || 'Unnamed Icon')
      }
    )
  }
}
