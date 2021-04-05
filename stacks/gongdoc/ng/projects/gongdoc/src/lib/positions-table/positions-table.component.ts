// generated from NgTableTemplateTS
import { Component, OnInit, OnChanges, Input, Output, EventEmitter, Inject, Optional } from '@angular/core';
import { MatTableDataSource } from '@angular/material/table';
import { MatButton } from '@angular/material/button'

import { MatDialogRef, MAT_DIALOG_DATA, MatDialog } from '@angular/material/dialog'
import { DialogData } from '../front-repo.service'
import { SelectionModel } from '@angular/cdk/collections';

const allowMultiSelect = true;

import { Router, RouterState } from '@angular/router';
import { PositionDB } from '../position-db'
import { PositionService } from '../position.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

// generated table component
@Component({
  selector: 'app-positions-table',
  templateUrl: './positions-table.component.html',
  styleUrls: ['./positions-table.component.css'],
})
export class PositionsTableComponent implements OnInit {

  // used if the component is called as a selection component of Position instances
  selection: SelectionModel<PositionDB>;
  initialSelection = new Array<PositionDB>();

  // the data source for the table
  positions: PositionDB[];

  // front repo, that will be referenced by this.positions
  frontRepo: FrontRepo

  // displayedColumns is referenced by the MatTable component for specify what columns
  // have to be displayed and in what order
  displayedColumns: string[];

  constructor(
    private positionService: PositionService,
    private frontRepoService: FrontRepoService,

    // not null if the component is called as a selection component of position instances
    public dialogRef: MatDialogRef<PositionsTableComponent>,
    @Optional() @Inject(MAT_DIALOG_DATA) public dialogData: DialogData,

    private router: Router,
  ) {
    // https://stackoverflow.com/questions/54627478/angular-7-routing-to-same-component-but-different-param-not-working
    // this is for routerLink on same component when only queryParameter changes
    this.router.routeReuseStrategy.shouldReuseRoute = function () {
      return false;
    };

    // observable for changes in structs
    this.positionService.PositionServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.getPositions()
        }
      }
    )
    if (dialogData == undefined) {
  	  this.displayedColumns = ['ID', 'Edit', 'Delete', // insertion point for columns to display
        "X",
        "Y",
        "Name",
      ]
    } else {
  	  this.displayedColumns = ['select', 'ID', // insertion point for columns to display
        "X",
        "Y",
        "Name",
      ]
      this.selection = new SelectionModel<PositionDB>(allowMultiSelect, this.initialSelection);
    }

  }

  ngOnInit(): void {
    this.getPositions()
  }

  getPositions(): void {
    this.frontRepoService.pull().subscribe(
      frontRepo => {
        this.frontRepo = frontRepo
        console.log("front repo pull returned")

        this.positions = this.frontRepo.Positions_array;
        // in case the component is called as a selection component
        if (this.dialogData != undefined) {
          this.positions.forEach(
            position => {
              let ID = this.dialogData.ID
              let revPointer = position[this.dialogData.ReversePointer]
              if (revPointer.Int64 == ID) {
                this.initialSelection.push(position)
              }
            }
          )
          this.selection = new SelectionModel<PositionDB>(allowMultiSelect, this.initialSelection);
        }
      }
    )
  }

  // newPosition initiate a new position
  // create a new Position objet
  newPosition() {
  }

  deletePosition(positionID: number, position: PositionDB) {
    // list of positions is truncated of position before the delete
    this.positions = this.positions.filter(h => h !== position);

    this.positionService.deletePosition(positionID).subscribe(
      position => {
        this.positionService.PositionServiceChanged.next("delete")

        console.log("position deleted")
      }
    );
  }

  editPosition(positionID: number, position: PositionDB) {

  }

  // display position in router
  displayPositionInRouter(positionID: number) {
    this.router.navigate( ["position-display", positionID])
  }

  // set editor outlet
  setEditorRouterOutlet(positionID: number) {
    this.router.navigate([{
      outlets: {
        editor: ["position-detail", positionID]
      }
    }]);
  }

  // set presentation outlet
  setPresentationRouterOutlet(positionID: number) {
    this.router.navigate([{
      outlets: {
        presentation: ["position-presentation", positionID]
      }
    }]);
  }

  /** Whether the number of selected elements matches the total number of rows. */
  isAllSelected() {
    const numSelected = this.selection.selected.length;
    const numRows = this.positions.length;
    return numSelected === numRows;
  }

  /** Selects all rows if they are not all selected; otherwise clear selection. */
  masterToggle() {
    this.isAllSelected() ?
      this.selection.clear() :
      this.positions.forEach(row => this.selection.select(row));
  }

  save() {

    let toUpdate = new Set<PositionDB>()

    // reset all initial selection of position that belong to aclass through Anarrayofb
    this.initialSelection.forEach(
      position => {
        position[this.dialogData.ReversePointer].Int64 = 0
        position[this.dialogData.ReversePointer].Valid = true
        toUpdate.add(position)
      }
    )

    // from selection, set position that belong to aclass through Anarrayofb
    this.selection.selected.forEach(
      position => {
        console.log("selection ID " + position.ID)
        let ID = +this.dialogData.ID
        position[this.dialogData.ReversePointer].Int64 = ID
        position[this.dialogData.ReversePointer].Valid = true
        toUpdate.add(position)
      }
    )

    // update all position (only update selection & initial selection)
    toUpdate.forEach(
      position => {
        this.positionService.updatePosition(position)
          .subscribe(position => {
            this.positionService.PositionServiceChanged.next("update")
            console.log("position saved")
          });
      }
    )
    this.dialogRef.close('Pizza!');
  }
}
