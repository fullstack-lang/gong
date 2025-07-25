import { AfterViewInit, Component, Inject, Input, OnInit, OnDestroy, Optional, ViewChild } from '@angular/core'
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
import { ConfirmationDialogComponent } from '../dialog/dialog.component'
import { MatIconModule } from '@angular/material/icon'
import { MatTooltipModule } from '@angular/material/tooltip'


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
    MatIconModule,
    MatTooltipModule,
  ],
  templateUrl: './table-specific.component.html',
  styleUrl: './table-specific.component.css'
})
export class TableSpecificComponent implements OnInit, AfterViewInit, OnDestroy {

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

  // Subscription management
  private subscriptions: Subscription = new Subscription()
  private webSocketSubscription?: Subscription
  private filterSubscription?: Subscription

  public gongtableFrontRepo?: table.FrontRepo

  // SelectionModel to hold the selected rows
  selection: SelectionModel<table.Row> = new SelectionModel<table.Row>(allowMultiSelect, [])
  initialSelection = new Array<table.Row>()

  // NEW: for toggling text truncation
  public isTextTruncated = true

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

    // Subscribe to filter control changes and add to subscriptions
    this.filterSubscription = this.filterControl.valueChanges
      .pipe(
        debounceTime(200),
        distinctUntilChanged()
      )
      .subscribe(value => {
        this.dataSource.filter = value ? value.trim().toLowerCase() : ''
      })

    // Add filter subscription to the main subscriptions container
    if (this.filterSubscription) {
      this.subscriptions.add(this.filterSubscription)
    }
  }

  ngOnDestroy(): void {
    // Unsubscribe from all subscriptions to prevent memory leaks
    this.subscriptions.unsubscribe()

    // Explicitly unsubscribe from the WebSocket connection if it exists
    if (this.webSocketSubscription) {
      this.webSocketSubscription.unsubscribe()
    }

    // Unsubscribe from the commit number subscription
    if (this.commutNbFromBackSubscription) {
      this.commutNbFromBackSubscription.unsubscribe()
    }
  }

  // NEW: toggle function for the text wrap button
  toggleTextTruncation(): void {
    this.isTextTruncated = !this.isTextTruncated
  }

  refresh(): void {
    // Unsubscribe from previous WebSocket connection if it exists
    if (this.webSocketSubscription) {
      this.webSocketSubscription.unsubscribe()
    }

    this.webSocketSubscription = this.gongtableFrontRepoService.connectToWebSocket(this.Name).subscribe(
      gongtablesFrontRepo => {
        this.gongtableFrontRepo = gongtablesFrontRepo
        this.selectedTable = undefined

        if (!this.gongtableFrontRepo) { // Guard against undefined front repo
          console.error("Gongtable front repo is not available.")
          return
        }

        let tableNames = new (Array<string>)
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

        // for sorting, we reorder the items according to the order in the association storage
        // and we remove the items that are not in the association storage
        if (this.selectedTable.CanDragDropRows && this.tableDialogData.AssociationStorage) {
          const sliceOfIDs = decodeStringToIntArray_json(this.tableDialogData.AssociationStorage)

          // --- DEBUGGING START ---
          // Step 1: Check the IDs you expect to find.
          console.log("Association IDs to order by:", sliceOfIDs)
          // --- DEBUGGING END ---

          // Create a Map for quick lookup of rows by their ID.
          const rowMapByID = new Map<number, table.Row>()
          for (const row of this.selectedTable.Rows) {
            if (row.Cells.length > 0 && row.Cells[0]?.CellInt) {

              // FIX: Explicitly convert the ID from the cell to a number to prevent type mismatch (e.g., "4" vs 4).
              const id = Number(row.Cells[0].CellInt.Value)

              // Only add the row to the map if its ID is a valid number.
              if (!isNaN(id)) {
                rowMapByID.set(id, row)
              }
            }
          }

          // --- DEBUGGING START ---
          // Step 2: Check the map that was created. Does it contain keys 4 and 1?
          // If not, the original `this.selectedTable.Rows` does not contain rows with these IDs.
          console.log("Map of available rows (ID -> Row):", rowMapByID)
          // --- DEBUGGING END ---

          // Build the new `Rows` array by iterating through the ordered IDs.
          const orderedRows: table.Row[] = sliceOfIDs
            .map(id => rowMapByID.get(id)) // Look up the row for each ID (as a number)
            .filter((row): row is table.Row => row !== undefined) // Filter out any IDs that didn't have a matching row

          // --- DEBUGGING START ---
          // Step 3: Check the final result.
          console.log("Final ordered and filtered rows:", orderedRows)
          // --- DEBUGGING END ---

          this.selectedTable.Rows = orderedRows
        }

        // Create the new dataSource
        this.dataSource = new MatTableDataSource(this.selectedTable.Rows || [])

        // NOW, set the filterPredicate on the new dataSource
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

        // IMPORTANT: Re-apply sort and paginator after creating new data source
        if (this.selectedTable.HasColumnSorting && this.sort) {
          this.dataSource.sort = this.sort;
          this.setupSortingDataAccessor()
        }

        if (this.selectedTable.HasPaginator && this.paginator) {
          this.dataSource.paginator = this.paginator;
        }

      }
    )

    // Add the WebSocket subscription to the main subscriptions container
    if (this.webSocketSubscription) {
      this.subscriptions.add(this.webSocketSubscription)
    }
  }

  ngAfterViewInit() {
    // Wait for next tick to ensure all view children are available
    setTimeout(() => {
      if (this.sort && this.selectedTable?.HasColumnSorting) {
        this.dataSource.sort = this.sort;

        // Set up the sorting data accessor
        this.dataSource.sortingDataAccessor = (row: table.Row, sortHeaderId: string): string | number => {
          if (!row.Cells) return "";
          const index = this.mapHeaderIdIndex.get(sortHeaderId)
          if (index === undefined) return "";

          const cell: table.Cell = row.Cells[index];
          if (cell.CellInt) return cell.CellInt.Value;
          if (cell.CellFloat64) return cell.CellFloat64.Value;
          if (cell.CellString) return cell.CellString.Value;
          if (cell.CellIcon) return cell.CellIcon.Icon || "";
          if (cell.CellBool) return cell.CellBool.Value ? 1 : 0; // Use numbers for boolean sorting
          return "";
        };
      }

      if (this.paginator && this.selectedTable?.HasPaginator) {
        this.dataSource.paginator = this.paginator;
      }
    })

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

    if (this.tableDialogData && this.dialogRef) {
      // Pass the current selection (which might be empty or unchanged)
      this.dialogRef.close(this.selection.selected)
    }
    return
  }

  drop(event: CdkDragDrop<string[]>) {
    if (!this.selectedTable || !this.selectedTable.Rows) {
      return
    }

    // This moves the item in the local array, reordering the rows
    moveItemInArray(this.selectedTable.Rows, event.previousIndex, event.currentIndex)

    // This updates the MatTableDataSource to reflect the new order in the UI
    this.dataSource.data = [...this.selectedTable.Rows]

    // one need to update the associated storage
    // [START] Completed Code
    if (this.selectedTable?.CanDragDropRows && this.tableDialogData?.AssociationStorage) {

      // 1. Extract the IDs from the rows in their new order.
      // We must ensure that we only process rows that have a valid ID structure,
      // which is assumed to be in the first cell as an integer (CellInt).
      const newOrderedIDs: number[] = this.selectedTable.Rows
        .map(row => {
          // Safely access the ID from the first cell
          if (row.Cells && row.Cells.length > 0 && row.Cells[0]?.CellInt) {
            return row.Cells[0].CellInt.Value
          }
          return null // Return null for rows that don't match the structure
        })
        .filter((id): id is number => id !== null) // Filter out any nulls to get a clean number array

      // 2. Convert the new array of IDs into a JSON string.
      const newAssociationStorage = JSON.stringify(newOrderedIDs)

      // 3. Update the AssociationStorage property in the dialog data.
      // This ensures that when the dialog is saved or closed, the new order is available
      // to the component that opened it.
      this.tableDialogData.AssociationStorage = newAssociationStorage

      if (this.sort) {
        this.sort.active = '';
        this.sort.direction = '';
      }

      console.log("Updated association order:", this.tableDialogData.AssociationStorage)
    }
    // [END] Completed Code
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
    const updateSubscription = this.rowService.updateFront(row, this.Name).subscribe(
      () => {
        console.log("Row updated on click:", row.Name || 'Unnamed Row')
      }
    )

    // Add the subscription to be cleaned up
    this.subscriptions.add(updateSubscription)
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

    if (cellIcon.NeedsConfirmation) {
      const dialogRef = this.dialog.open(ConfirmationDialogComponent, {
        width: '250px',
        data: { message: cellIcon.ConfirmationMessage }
      })

      const dialogSubscription = dialogRef.afterClosed().subscribe(result => {
        if (!result) {
          return
        } else {
          const updateSubscription = this.celliconService.updateFront(cellIcon, this.Name).subscribe(
            () => {
              console.log("Cell icon updated after confirmation:", cellIcon.Name || 'Unnamed Icon')
            }
          )
          // Add the subscription to be cleaned up
          this.subscriptions.add(updateSubscription)
        }
      })

      // Add the dialog subscription to be cleaned up
      this.subscriptions.add(dialogSubscription)
    } else {
      const updateSubscription = this.celliconService.updateFront(cellIcon, this.Name).subscribe(
        () => {
          console.log("Cell icon updated:", cellIcon.Name || 'Unnamed Icon')
        }
      )

      // Add the subscription to be cleaned up
      this.subscriptions.add(updateSubscription)
    }
  }

  private setupSortingDataAccessor(): void {
    this.dataSource.sortingDataAccessor = (row: table.Row, sortHeaderId: string): string | number => {
      if (!row.Cells) return "";

      const index = this.mapHeaderIdIndex.get(sortHeaderId)
      if (index === undefined) {
        console.warn(`Column index not found for: ${sortHeaderId}`)
        return "";
      }

      const cell: table.Cell = row.Cells[index];
      if (!cell) return "";

      // Handle different cell types with proper type conversion
      if (cell.CellInt) return Number(cell.CellInt.Value) || 0;
      if (cell.CellFloat64) return Number(cell.CellFloat64.Value) || 0;
      if (cell.CellString) return String(cell.CellString.Value || "").toLowerCase()
      if (cell.CellIcon) return String(cell.CellIcon.Icon || "").toLowerCase()
      if (cell.CellBool) return cell.CellBool.Value ? 1 : 0;

      return "";
    };
  }

  debugSortingState(): void {
    console.log('=== Sorting Debug Info ===')
    console.log('selectedTable.HasColumnSorting:', this.selectedTable?.HasColumnSorting)
    console.log('sort ViewChild:', this.sort)
    console.log('dataSource.sort:', this.dataSource.sort)
    console.log('displayedColumns:', this.displayedColumns)
    console.log('mapHeaderIdIndex:', this.mapHeaderIdIndex)
    console.log('dataSource.data length:', this.dataSource.data.length)

    if (this.sort) {
      console.log('sort.active:', this.sort.active)
      console.log('sort.direction:', this.sort.direction)
    }
  }
}