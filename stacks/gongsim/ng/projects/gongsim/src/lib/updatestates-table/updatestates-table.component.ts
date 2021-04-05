// generated from NgTableTemplateTS
import { Component, OnInit, OnChanges, Input, Output, EventEmitter, Inject, Optional } from '@angular/core';
import { MatTableDataSource } from '@angular/material/table';
import { MatButton } from '@angular/material/button'

import { MatDialogRef, MAT_DIALOG_DATA, MatDialog } from '@angular/material/dialog'
import { DialogData } from '../front-repo.service'
import { SelectionModel } from '@angular/cdk/collections';

const allowMultiSelect = true;

import { Router, RouterState } from '@angular/router';
import { UpdateStateDB } from '../updatestate-db'
import { UpdateStateService } from '../updatestate.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

// generated table component
@Component({
  selector: 'app-updatestates-table',
  templateUrl: './updatestates-table.component.html',
  styleUrls: ['./updatestates-table.component.css'],
})
export class UpdateStatesTableComponent implements OnInit {

  // used if the component is called as a selection component of UpdateState instances
  selection: SelectionModel<UpdateStateDB>;
  initialSelection = new Array<UpdateStateDB>();

  // the data source for the table
  updatestates: UpdateStateDB[];

  // front repo, that will be referenced by this.updatestates
  frontRepo: FrontRepo

  // displayedColumns is referenced by the MatTable component for specify what columns
  // have to be displayed and in what order
  displayedColumns: string[];

  constructor(
    private updatestateService: UpdateStateService,
    private frontRepoService: FrontRepoService,

    // not null if the component is called as a selection component of updatestate instances
    public dialogRef: MatDialogRef<UpdateStatesTableComponent>,
    @Optional() @Inject(MAT_DIALOG_DATA) public dialogData: DialogData,

    private router: Router,
  ) {
    // https://stackoverflow.com/questions/54627478/angular-7-routing-to-same-component-but-different-param-not-working
    // this is for routerLink on same component when only queryParameter changes
    this.router.routeReuseStrategy.shouldReuseRoute = function () {
      return false;
    };

    // observable for changes in structs
    this.updatestateService.UpdateStateServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.getUpdateStates()
        }
      }
    )
    if (dialogData == undefined) {
  	  this.displayedColumns = ['ID', 'Edit', 'Delete', // insertion point for columns to display
        "Name",
        "Duration",
        "Period",
      ]
    } else {
  	  this.displayedColumns = ['select', 'ID', // insertion point for columns to display
        "Name",
        "Duration",
        "Period",
      ]
      this.selection = new SelectionModel<UpdateStateDB>(allowMultiSelect, this.initialSelection);
    }

  }

  ngOnInit(): void {
    this.getUpdateStates()
  }

  getUpdateStates(): void {
    this.frontRepoService.pull().subscribe(
      frontRepo => {
        this.frontRepo = frontRepo
        console.log("front repo pull returned")

        this.updatestates = this.frontRepo.UpdateStates_array;
        // in case the component is called as a selection component
        if (this.dialogData != undefined) {
          this.updatestates.forEach(
            updatestate => {
              let ID = this.dialogData.ID
              let revPointer = updatestate[this.dialogData.ReversePointer]
              if (revPointer.Int64 == ID) {
                this.initialSelection.push(updatestate)
              }
            }
          )
          this.selection = new SelectionModel<UpdateStateDB>(allowMultiSelect, this.initialSelection);
        }
      }
    )
  }

  // newUpdateState initiate a new updatestate
  // create a new UpdateState objet
  newUpdateState() {
  }

  deleteUpdateState(updatestateID: number, updatestate: UpdateStateDB) {
    // list of updatestates is truncated of updatestate before the delete
    this.updatestates = this.updatestates.filter(h => h !== updatestate);

    this.updatestateService.deleteUpdateState(updatestateID).subscribe(
      updatestate => {
        this.updatestateService.UpdateStateServiceChanged.next("delete")

        console.log("updatestate deleted")
      }
    );
  }

  editUpdateState(updatestateID: number, updatestate: UpdateStateDB) {

  }

  // display updatestate in router
  displayUpdateStateInRouter(updatestateID: number) {
    this.router.navigate( ["updatestate-display", updatestateID])
  }

  // set editor outlet
  setEditorRouterOutlet(updatestateID: number) {
    this.router.navigate([{
      outlets: {
        editor: ["updatestate-detail", updatestateID]
      }
    }]);
  }

  // set presentation outlet
  setPresentationRouterOutlet(updatestateID: number) {
    this.router.navigate([{
      outlets: {
        presentation: ["updatestate-presentation", updatestateID]
      }
    }]);
  }

  /** Whether the number of selected elements matches the total number of rows. */
  isAllSelected() {
    const numSelected = this.selection.selected.length;
    const numRows = this.updatestates.length;
    return numSelected === numRows;
  }

  /** Selects all rows if they are not all selected; otherwise clear selection. */
  masterToggle() {
    this.isAllSelected() ?
      this.selection.clear() :
      this.updatestates.forEach(row => this.selection.select(row));
  }

  save() {

    let toUpdate = new Set<UpdateStateDB>()

    // reset all initial selection of updatestate that belong to aclass through Anarrayofb
    this.initialSelection.forEach(
      updatestate => {
        updatestate[this.dialogData.ReversePointer].Int64 = 0
        updatestate[this.dialogData.ReversePointer].Valid = true
        toUpdate.add(updatestate)
      }
    )

    // from selection, set updatestate that belong to aclass through Anarrayofb
    this.selection.selected.forEach(
      updatestate => {
        console.log("selection ID " + updatestate.ID)
        let ID = +this.dialogData.ID
        updatestate[this.dialogData.ReversePointer].Int64 = ID
        updatestate[this.dialogData.ReversePointer].Valid = true
        toUpdate.add(updatestate)
      }
    )

    // update all updatestate (only update selection & initial selection)
    toUpdate.forEach(
      updatestate => {
        this.updatestateService.updateUpdateState(updatestate)
          .subscribe(updatestate => {
            this.updatestateService.UpdateStateServiceChanged.next("update")
            console.log("updatestate saved")
          });
      }
    )
    this.dialogRef.close('Pizza!');
  }
}
