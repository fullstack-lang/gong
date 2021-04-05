// generated from NgTableTemplateTS
import { Component, OnInit, OnChanges, Input, Output, EventEmitter, Inject, Optional } from '@angular/core';
import { MatTableDataSource } from '@angular/material/table';
import { MatButton } from '@angular/material/button'

import { MatDialogRef, MAT_DIALOG_DATA, MatDialog } from '@angular/material/dialog'
import { DialogData } from '../front-repo.service'
import { SelectionModel } from '@angular/cdk/collections';

const allowMultiSelect = true;

import { Router, RouterState } from '@angular/router';
import { StateDB } from '../state-db'
import { StateService } from '../state.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

// generated table component
@Component({
  selector: 'app-states-table',
  templateUrl: './states-table.component.html',
  styleUrls: ['./states-table.component.css'],
})
export class StatesTableComponent implements OnInit {

  // used if the component is called as a selection component of State instances
  selection: SelectionModel<StateDB>;
  initialSelection = new Array<StateDB>();

  // the data source for the table
  states: StateDB[];

  // front repo, that will be referenced by this.states
  frontRepo: FrontRepo

  // displayedColumns is referenced by the MatTable component for specify what columns
  // have to be displayed and in what order
  displayedColumns: string[];

  constructor(
    private stateService: StateService,
    private frontRepoService: FrontRepoService,

    // not null if the component is called as a selection component of state instances
    public dialogRef: MatDialogRef<StatesTableComponent>,
    @Optional() @Inject(MAT_DIALOG_DATA) public dialogData: DialogData,

    private router: Router,
  ) {
    // https://stackoverflow.com/questions/54627478/angular-7-routing-to-same-component-but-different-param-not-working
    // this is for routerLink on same component when only queryParameter changes
    this.router.routeReuseStrategy.shouldReuseRoute = function () {
      return false;
    };

    // observable for changes in structs
    this.stateService.StateServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.getStates()
        }
      }
    )
    if (dialogData == undefined) {
  	  this.displayedColumns = ['ID', 'Edit', 'Delete', // insertion point for columns to display
        "Name",
        "X",
        "Y",
        "States",
      ]
    } else {
  	  this.displayedColumns = ['select', 'ID', // insertion point for columns to display
        "Name",
        "X",
        "Y",
        "States",
      ]
      this.selection = new SelectionModel<StateDB>(allowMultiSelect, this.initialSelection);
    }

  }

  ngOnInit(): void {
    this.getStates()
  }

  getStates(): void {
    this.frontRepoService.pull().subscribe(
      frontRepo => {
        this.frontRepo = frontRepo
        console.log("front repo pull returned")

        this.states = this.frontRepo.States_array;
        // in case the component is called as a selection component
        if (this.dialogData != undefined) {
          this.states.forEach(
            state => {
              let ID = this.dialogData.ID
              let revPointer = state[this.dialogData.ReversePointer]
              if (revPointer.Int64 == ID) {
                this.initialSelection.push(state)
              }
            }
          )
          this.selection = new SelectionModel<StateDB>(allowMultiSelect, this.initialSelection);
        }
      }
    )
  }

  // newState initiate a new state
  // create a new State objet
  newState() {
  }

  deleteState(stateID: number, state: StateDB) {
    // list of states is truncated of state before the delete
    this.states = this.states.filter(h => h !== state);

    this.stateService.deleteState(stateID).subscribe(
      state => {
        this.stateService.StateServiceChanged.next("delete")

        console.log("state deleted")
      }
    );
  }

  editState(stateID: number, state: StateDB) {

  }

  // display state in router
  displayStateInRouter(stateID: number) {
    this.router.navigate( ["state-display", stateID])
  }

  // set editor outlet
  setEditorRouterOutlet(stateID: number) {
    this.router.navigate([{
      outlets: {
        editor: ["state-detail", stateID]
      }
    }]);
  }

  // set presentation outlet
  setPresentationRouterOutlet(stateID: number) {
    this.router.navigate([{
      outlets: {
        presentation: ["state-presentation", stateID]
      }
    }]);
  }

  /** Whether the number of selected elements matches the total number of rows. */
  isAllSelected() {
    const numSelected = this.selection.selected.length;
    const numRows = this.states.length;
    return numSelected === numRows;
  }

  /** Selects all rows if they are not all selected; otherwise clear selection. */
  masterToggle() {
    this.isAllSelected() ?
      this.selection.clear() :
      this.states.forEach(row => this.selection.select(row));
  }

  save() {

    let toUpdate = new Set<StateDB>()

    // reset all initial selection of state that belong to aclass through Anarrayofb
    this.initialSelection.forEach(
      state => {
        state[this.dialogData.ReversePointer].Int64 = 0
        state[this.dialogData.ReversePointer].Valid = true
        toUpdate.add(state)
      }
    )

    // from selection, set state that belong to aclass through Anarrayofb
    this.selection.selected.forEach(
      state => {
        console.log("selection ID " + state.ID)
        let ID = +this.dialogData.ID
        state[this.dialogData.ReversePointer].Int64 = ID
        state[this.dialogData.ReversePointer].Valid = true
        toUpdate.add(state)
      }
    )

    // update all state (only update selection & initial selection)
    toUpdate.forEach(
      state => {
        this.stateService.updateState(state)
          .subscribe(state => {
            this.stateService.StateServiceChanged.next("update")
            console.log("state saved")
          });
      }
    )
    this.dialogRef.close('Pizza!');
  }
}
