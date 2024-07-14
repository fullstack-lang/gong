import { AfterViewInit, Component, Inject, Input, OnInit, Optional, ViewChild } from '@angular/core';
import { Subscription, debounceTime, distinctUntilChanged, forkJoin } from 'rxjs';
import { CdkDragDrop, moveItemInArray } from '@angular/cdk/drag-drop';

import * as gongtable from '../../../../gongtable/src/public-api'

import { MatTableDataSource } from '@angular/material/table';
import { FormControl } from '@angular/forms';

import { SelectionModel } from '@angular/cdk/collections';
import { TableDialogData } from '../table-dialog-data';

const allowMultiSelect = true

import { MAT_DIALOG_DATA, MatDialog, MatDialogModule, MatDialogRef } from '@angular/material/dialog';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { MatSort, MatSortModule } from '@angular/material/sort';
import { DragDropModule } from '@angular/cdk/drag-drop';
import { CommonModule } from '@angular/common';

import { MatFormFieldModule } from '@angular/material/form-field';
import { MatPaginator, MatPaginatorModule } from '@angular/material/paginator';
import { MatTableModule } from '@angular/material/table';
import { MatCheckboxModule } from '@angular/material/checkbox';
import { MatInputModule } from '@angular/material/input';
import { MatButtonModule } from '@angular/material/button';


@Component({
  selector: 'lib-material-table',
  templateUrl: './material-table.component.html',
  styleUrls: ['./material-table.component.css'],
  standalone: true,
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
})
export class MaterialTableComponent implements OnInit, AfterViewInit {

  displayedColumns: string[] = []
  allDisplayedColumns: string[] = [] // in case there is a checkbox

  mapHeaderIdIndex = new Map<string, number>()

  dataSource = new MatTableDataSource<gongtable.Row>()

  stickyStyle = {
    position: 'sticky',
    left: '0',
    zIndex: '1'
  }
  // for selection
  selectedTable: gongtable.Table | undefined = undefined;

  @Input() DataStack: string = ""
  @Input() TableName: string = ""

  // for filtering
  filterControl = new FormControl()

  // for sorting
  @ViewChild(MatSort) sort: MatSort | undefined
  matSortDirective: string = ""

  // for pagination
  @ViewChild(MatPaginator)
  paginator: MatPaginator | undefined;

  // the component is refreshed when modification are performed in the back repo 
  // 
  // the checkCommitNbFromBackTimer polls the commit number of the back repo
  // if the commit number has increased, it pulls the front repo and redraw the diagram
  private commutNbFromBackSubscription: Subscription = new Subscription
  lastCommitNbFromBack = -1
  lastPushFromFrontNb = -1
  currTime: number = 0
  dateOfLastTimerEmission: Date = new Date


  public gongtableFrontRepo?: gongtable.FrontRepo

  constructor(
    private gongtableFrontRepoService: gongtable.FrontRepoService,
    private gongtableCommitNbFromBackService: gongtable.CommitNbFromBackService,
    private rowService: gongtable.RowService,
    private tableService: gongtable.TableService,
    private celliconService: gongtable.CellIconService,


    // not null if the component is called as a selection component of cellboolean instances
    public dialog: MatDialog,
    @Optional() @Inject(MAT_DIALOG_DATA) public tableDialogData: TableDialogData,
  ) {

  }

  ngOnInit(): void {

    // if the component is started via component, one needs to fetch DataStack and TableName from
    // the dialog data
    if (this.tableDialogData) {
      this.DataStack = this.tableDialogData.DataStack
      this.TableName = this.tableDialogData.TableName
    }

    this.refresh()

    this.filterControl.valueChanges
      .pipe(
        debounceTime(200), // Optional. To reduce number of requests.
        distinctUntilChanged() // Optional. To prevent same filter fire multiple times.
      )
      .subscribe(value => {
        this.dataSource.filter = value;
      })

    this.dataSource.filterPredicate = (data: any, filter: string) => {
      const dataStr = JSON.stringify(data).toLowerCase(); // Convert the data to a lower case string.
      return dataStr.indexOf(filter) !== -1;
    }
  }

  refresh(): void {

    this.gongtableFrontRepoService.connectToWebSocket(this.DataStack).subscribe(
      gongtablesFrontRepo => {
        this.gongtableFrontRepo = gongtablesFrontRepo

        this.selectedTable = undefined

        // use of the refactorable version
        for (let table of this.gongtableFrontRepo.getFrontArray<gongtable.Table>(gongtable.Table.GONGSTRUCT_NAME)) {
          if (table.Name == this.TableName) {
            this.selectedTable = table
          }
        }

        if (this.selectedTable == undefined) {
          return
        }

        this.dataSource = new MatTableDataSource(this.selectedTable.Rows!)

        if (this.selectedTable.HasCheckableRows) {

        }

        // enable filtering on all fields (including pointers and reverse pointer, which is not done by default)

        if (this.selectedTable.DisplayedColumns == undefined) {
          return
        }

        this.mapHeaderIdIndex = new Map<string, number>()
        let index = 0

        this.displayedColumns = []
        for (let column of this.selectedTable.DisplayedColumns) {
          this.mapHeaderIdIndex.set(column.Name, index)
          this.displayedColumns.push(column.Name)
          index++
        }
        this.allDisplayedColumns = []
        if (this.selectedTable.HasCheckableRows) {
          this.allDisplayedColumns = ['select']

          if (this.selectedTable.Rows != undefined) {
            this.initialSelection = []
            for (let Row of this.selectedTable.Rows) {
              if (Row.IsChecked) {
                this.initialSelection.push(Row)
              }
            }
            this.selection = new SelectionModel<gongtable.Row>(allowMultiSelect, this.initialSelection)
          }

        }
        this.allDisplayedColumns = this.allDisplayedColumns.concat(this.displayedColumns)

        if (this.selectedTable.HasFiltering) {
          this.dataSource.filterPredicate = (Row: gongtable.Row, filter: string) => {

            // filtering is based on finding a lower case filter into a concatenated string
            // the Cell properties
            let mergedContent = ""

            for (let cell of Row.Cells!) {
              if (cell.CellInt) {
                mergedContent += cell.CellInt.Value
              }
              if (cell.CellFloat64) {
                mergedContent += cell.CellFloat64.Value
              }
              if (cell.CellString) {
                mergedContent += cell.CellString.Value
              }
            }

            mergedContent = mergedContent.toLowerCase()
            let isSelected = mergedContent.includes(filter.toLowerCase())
            return isSelected
          }
        }

        this.matSortDirective = ""
        if (this.selectedTable.HasColumnSorting) {
          this.dataSource.sort = this.sort!
          this.matSortDirective = "mat-sort"

          // enable sorting on all fields (including pointers and reverse pointer)
          this.dataSource.sortingDataAccessor = (Row: gongtable.Row, sortHeaderId: string) => {

            if (Row.Cells == undefined) {
              return ""
            }
            let index = this.mapHeaderIdIndex.get(sortHeaderId)
            if (index == undefined) {
              return ""
            }

            let cell: gongtable.Cell = Row.Cells[index]
            if (cell.CellInt) {
              return cell.CellInt.Value
            }
            if (cell.CellFloat64) {
              return cell.CellFloat64.Value
            }
            if (cell.CellString) {
              return cell.CellString.Value
            }
            if (cell.CellIcon) {
              return cell.CellIcon.Icon
            }
            if (cell.CellBool) {
              if (cell.CellBool.Value) {
                return "true"
              } else {
                return "false"
              }
            }

            return "";
          };
        }

        if (this.selectedTable.HasPaginator) {
          this.dataSource.paginator = this.paginator!
        }
      }
    )
  }

  ngAfterViewInit() {
    this.dataSource.sort = this.sort!
  }

  applyFilter(event: Event) {
    const filterValue = (event.target as HTMLInputElement).value;
    this.dataSource.filter = filterValue.trim().toLowerCase();
  }

  selection: SelectionModel<gongtable.Row> = new (SelectionModel)
  initialSelection = new Array<gongtable.Row>()

  isAllSelected() {
    const numSelected = this.selection.selected.length;
    const numRows = this.selectedTable?.Rows?.length
    return numSelected === numRows;
  }


  masterToggle() {
    this.isAllSelected() ?
      this.selection.clear() :
      this.selectedTable?.Rows?.forEach(row => this.selection.select(row));
  }

  save() {

    if (this.selectedTable == undefined) {
      return
    }

    // map of modified rows to be be updated
    let modifiedRows = new Set<gongtable.Row>
    for (let row of this.selectedTable.Rows!) {
      if ((row.IsChecked && !this.selection.isSelected(row)) ||
        (!row.IsChecked && this.selection.isSelected(row))) {
        row.IsChecked = !row.IsChecked
        modifiedRows.add(row)
      }
    }

    if (modifiedRows.size == 0) {
      // in case this component is called as a modal window (MatDialog)
      // exits,
      this.selectedTable.SavingInProgress = true
      this.tableService.updateFront(this.selectedTable!, this.DataStack).subscribe(
        () => {
          // in case this component is called as a modal window (MatDialog)
          // exits,
          if (this.tableDialogData) {
            this.dialog.closeAll()
          }
        }
      )
      return
    }

    // inform the back that the saving is some rows is in progress
    this.selectedTable.SavingInProgress = true


    const promises = []
    for (let row of modifiedRows) {
      promises.push(this.rowService.updateFront(row, this.DataStack))
    }

    forkJoin(promises).subscribe(
      () => {

        this.selectedTable!.SavingInProgress = false
        this.tableService.updateFront(this.selectedTable!, this.DataStack).subscribe(
          () => {
            // in case this component is called as a modal window (MatDialog)
            // exits,
            if (this.tableDialogData) {
              this.dialog.closeAll()
            }
          }
        )
      }

    )


  }

  drop(event: CdkDragDrop<string[]>) {
    moveItemInArray(this.selectedTable?.Rows!, event.previousIndex, event.currentIndex)

    this.tableService.updateFront(this.selectedTable!, this.DataStack).subscribe(
      () => {
        console.log("table", this.selectedTable?.Name, "rows shuffled")
      }
    )
  }

  isDraggableRow = (index: number, item: gongtable.Row) => this.selectedTable?.CanDragDropRows

  close() {
    if (this.tableDialogData) {
      this.dialog.closeAll()
    }
  }

  // onClick performs an update of the clicked row (without any property change)
  // this minimalist design will hopefully be sufficient for the backend to interpret
  // that the row has been clicked
  onClick(row: gongtable.Row) {
    console.log("Material Table: onClick: Stack: `" + this.DataStack + "`table:`" + this.TableName + "`row:" + row.Name)

    let cells = row.Cells

    this.rowService.updateFront(row, this.DataStack).subscribe(
      () => {
        console.log("row updated")
        row.Cells = cells
      }
    )
  }

  getDynamicStyles(columnIndex: number): { [key: string]: any } {
    const styles: { [key: string]: any } = {} // Explicitly define the type here 
    if (this.selectedTable == undefined) {
      return styles
    }

    if (columnIndex <= this.selectedTable.NbOfStickyColumns) {
      styles['position'] = 'sticky'
      return styles
    }


    return styles
  }

  onClickCellIcon(cellIcon: gongtable.CellIcon) {
    console.log("Cell Icon clicked")
    this.celliconService.updateFront(cellIcon, this.DataStack).subscribe(
      () => {

      }
    )
  }
}
