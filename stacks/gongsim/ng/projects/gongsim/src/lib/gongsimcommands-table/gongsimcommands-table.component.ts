// generated from NgTableTemplateTS
import { Component, OnInit, OnChanges, Input, Output, EventEmitter, Inject, Optional } from '@angular/core';
import { MatTableDataSource } from '@angular/material/table';
import { MatButton } from '@angular/material/button'

import { MatDialogRef, MAT_DIALOG_DATA, MatDialog } from '@angular/material/dialog'
import { DialogData } from '../front-repo.service'
import { SelectionModel } from '@angular/cdk/collections';

const allowMultiSelect = true;

import { Router, RouterState } from '@angular/router';
import { GongsimCommandDB } from '../gongsimcommand-db'
import { GongsimCommandService } from '../gongsimcommand.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

// generated table component
@Component({
  selector: 'app-gongsimcommands-table',
  templateUrl: './gongsimcommands-table.component.html',
  styleUrls: ['./gongsimcommands-table.component.css'],
})
export class GongsimCommandsTableComponent implements OnInit {

  // used if the component is called as a selection component of GongsimCommand instances
  selection: SelectionModel<GongsimCommandDB>;
  initialSelection = new Array<GongsimCommandDB>();

  // the data source for the table
  gongsimcommands: GongsimCommandDB[];

  // front repo, that will be referenced by this.gongsimcommands
  frontRepo: FrontRepo

  // displayedColumns is referenced by the MatTable component for specify what columns
  // have to be displayed and in what order
  displayedColumns: string[];

  constructor(
    private gongsimcommandService: GongsimCommandService,
    private frontRepoService: FrontRepoService,

    // not null if the component is called as a selection component of gongsimcommand instances
    public dialogRef: MatDialogRef<GongsimCommandsTableComponent>,
    @Optional() @Inject(MAT_DIALOG_DATA) public dialogData: DialogData,

    private router: Router,
  ) {
    // https://stackoverflow.com/questions/54627478/angular-7-routing-to-same-component-but-different-param-not-working
    // this is for routerLink on same component when only queryParameter changes
    this.router.routeReuseStrategy.shouldReuseRoute = function () {
      return false;
    };

    // observable for changes in structs
    this.gongsimcommandService.GongsimCommandServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.getGongsimCommands()
        }
      }
    )
    if (dialogData == undefined) {
  	  this.displayedColumns = ['ID', 'Edit', 'Delete', // insertion point for columns to display
        "Name",
        "Command",
        "CommandDate",
        "SpeedCommandType",
        "DateSpeedCommand",
      ]
    } else {
  	  this.displayedColumns = ['select', 'ID', // insertion point for columns to display
        "Name",
        "Command",
        "CommandDate",
        "SpeedCommandType",
        "DateSpeedCommand",
      ]
      this.selection = new SelectionModel<GongsimCommandDB>(allowMultiSelect, this.initialSelection);
    }

  }

  ngOnInit(): void {
    this.getGongsimCommands()
  }

  getGongsimCommands(): void {
    this.frontRepoService.pull().subscribe(
      frontRepo => {
        this.frontRepo = frontRepo
        console.log("front repo pull returned")

        this.gongsimcommands = this.frontRepo.GongsimCommands_array;
        // in case the component is called as a selection component
        if (this.dialogData != undefined) {
          this.gongsimcommands.forEach(
            gongsimcommand => {
              let ID = this.dialogData.ID
              let revPointer = gongsimcommand[this.dialogData.ReversePointer]
              if (revPointer.Int64 == ID) {
                this.initialSelection.push(gongsimcommand)
              }
            }
          )
          this.selection = new SelectionModel<GongsimCommandDB>(allowMultiSelect, this.initialSelection);
        }
      }
    )
  }

  // newGongsimCommand initiate a new gongsimcommand
  // create a new GongsimCommand objet
  newGongsimCommand() {
  }

  deleteGongsimCommand(gongsimcommandID: number, gongsimcommand: GongsimCommandDB) {
    // list of gongsimcommands is truncated of gongsimcommand before the delete
    this.gongsimcommands = this.gongsimcommands.filter(h => h !== gongsimcommand);

    this.gongsimcommandService.deleteGongsimCommand(gongsimcommandID).subscribe(
      gongsimcommand => {
        this.gongsimcommandService.GongsimCommandServiceChanged.next("delete")

        console.log("gongsimcommand deleted")
      }
    );
  }

  editGongsimCommand(gongsimcommandID: number, gongsimcommand: GongsimCommandDB) {

  }

  // display gongsimcommand in router
  displayGongsimCommandInRouter(gongsimcommandID: number) {
    this.router.navigate( ["gongsimcommand-display", gongsimcommandID])
  }

  // set editor outlet
  setEditorRouterOutlet(gongsimcommandID: number) {
    this.router.navigate([{
      outlets: {
        editor: ["gongsimcommand-detail", gongsimcommandID]
      }
    }]);
  }

  // set presentation outlet
  setPresentationRouterOutlet(gongsimcommandID: number) {
    this.router.navigate([{
      outlets: {
        presentation: ["gongsimcommand-presentation", gongsimcommandID]
      }
    }]);
  }

  /** Whether the number of selected elements matches the total number of rows. */
  isAllSelected() {
    const numSelected = this.selection.selected.length;
    const numRows = this.gongsimcommands.length;
    return numSelected === numRows;
  }

  /** Selects all rows if they are not all selected; otherwise clear selection. */
  masterToggle() {
    this.isAllSelected() ?
      this.selection.clear() :
      this.gongsimcommands.forEach(row => this.selection.select(row));
  }

  save() {

    let toUpdate = new Set<GongsimCommandDB>()

    // reset all initial selection of gongsimcommand that belong to aclass through Anarrayofb
    this.initialSelection.forEach(
      gongsimcommand => {
        gongsimcommand[this.dialogData.ReversePointer].Int64 = 0
        gongsimcommand[this.dialogData.ReversePointer].Valid = true
        toUpdate.add(gongsimcommand)
      }
    )

    // from selection, set gongsimcommand that belong to aclass through Anarrayofb
    this.selection.selected.forEach(
      gongsimcommand => {
        console.log("selection ID " + gongsimcommand.ID)
        let ID = +this.dialogData.ID
        gongsimcommand[this.dialogData.ReversePointer].Int64 = ID
        gongsimcommand[this.dialogData.ReversePointer].Valid = true
        toUpdate.add(gongsimcommand)
      }
    )

    // update all gongsimcommand (only update selection & initial selection)
    toUpdate.forEach(
      gongsimcommand => {
        this.gongsimcommandService.updateGongsimCommand(gongsimcommand)
          .subscribe(gongsimcommand => {
            this.gongsimcommandService.GongsimCommandServiceChanged.next("update")
            console.log("gongsimcommand saved")
          });
      }
    )
    this.dialogRef.close('Pizza!');
  }
}
