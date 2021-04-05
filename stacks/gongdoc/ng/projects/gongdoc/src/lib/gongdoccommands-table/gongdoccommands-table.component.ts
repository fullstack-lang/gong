// generated from NgTableTemplateTS
import { Component, OnInit, OnChanges, Input, Output, EventEmitter, Inject, Optional } from '@angular/core';
import { MatTableDataSource } from '@angular/material/table';
import { MatButton } from '@angular/material/button'

import { MatDialogRef, MAT_DIALOG_DATA, MatDialog } from '@angular/material/dialog'
import { DialogData } from '../front-repo.service'
import { SelectionModel } from '@angular/cdk/collections';

const allowMultiSelect = true;

import { Router, RouterState } from '@angular/router';
import { GongdocCommandDB } from '../gongdoccommand-db'
import { GongdocCommandService } from '../gongdoccommand.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

// generated table component
@Component({
  selector: 'app-gongdoccommands-table',
  templateUrl: './gongdoccommands-table.component.html',
  styleUrls: ['./gongdoccommands-table.component.css'],
})
export class GongdocCommandsTableComponent implements OnInit {

  // used if the component is called as a selection component of GongdocCommand instances
  selection: SelectionModel<GongdocCommandDB>;
  initialSelection = new Array<GongdocCommandDB>();

  // the data source for the table
  gongdoccommands: GongdocCommandDB[];

  // front repo, that will be referenced by this.gongdoccommands
  frontRepo: FrontRepo

  // displayedColumns is referenced by the MatTable component for specify what columns
  // have to be displayed and in what order
  displayedColumns: string[];

  constructor(
    private gongdoccommandService: GongdocCommandService,
    private frontRepoService: FrontRepoService,

    // not null if the component is called as a selection component of gongdoccommand instances
    public dialogRef: MatDialogRef<GongdocCommandsTableComponent>,
    @Optional() @Inject(MAT_DIALOG_DATA) public dialogData: DialogData,

    private router: Router,
  ) {
    // https://stackoverflow.com/questions/54627478/angular-7-routing-to-same-component-but-different-param-not-working
    // this is for routerLink on same component when only queryParameter changes
    this.router.routeReuseStrategy.shouldReuseRoute = function () {
      return false;
    };

    // observable for changes in structs
    this.gongdoccommandService.GongdocCommandServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.getGongdocCommands()
        }
      }
    )
    if (dialogData == undefined) {
  	  this.displayedColumns = ['ID', 'Edit', 'Delete', // insertion point for columns to display
        "Name",
        "Command",
        "DiagramName",
        "Date",
        "GongdocNodeType",
        "StructName",
        "FieldName",
        "FieldTypeName",
        "PositionX",
        "PositionY",
      ]
    } else {
  	  this.displayedColumns = ['select', 'ID', // insertion point for columns to display
        "Name",
        "Command",
        "DiagramName",
        "Date",
        "GongdocNodeType",
        "StructName",
        "FieldName",
        "FieldTypeName",
        "PositionX",
        "PositionY",
      ]
      this.selection = new SelectionModel<GongdocCommandDB>(allowMultiSelect, this.initialSelection);
    }

  }

  ngOnInit(): void {
    this.getGongdocCommands()
  }

  getGongdocCommands(): void {
    this.frontRepoService.pull().subscribe(
      frontRepo => {
        this.frontRepo = frontRepo
        console.log("front repo pull returned")

        this.gongdoccommands = this.frontRepo.GongdocCommands_array;
        // in case the component is called as a selection component
        if (this.dialogData != undefined) {
          this.gongdoccommands.forEach(
            gongdoccommand => {
              let ID = this.dialogData.ID
              let revPointer = gongdoccommand[this.dialogData.ReversePointer]
              if (revPointer.Int64 == ID) {
                this.initialSelection.push(gongdoccommand)
              }
            }
          )
          this.selection = new SelectionModel<GongdocCommandDB>(allowMultiSelect, this.initialSelection);
        }
      }
    )
  }

  // newGongdocCommand initiate a new gongdoccommand
  // create a new GongdocCommand objet
  newGongdocCommand() {
  }

  deleteGongdocCommand(gongdoccommandID: number, gongdoccommand: GongdocCommandDB) {
    // list of gongdoccommands is truncated of gongdoccommand before the delete
    this.gongdoccommands = this.gongdoccommands.filter(h => h !== gongdoccommand);

    this.gongdoccommandService.deleteGongdocCommand(gongdoccommandID).subscribe(
      gongdoccommand => {
        this.gongdoccommandService.GongdocCommandServiceChanged.next("delete")

        console.log("gongdoccommand deleted")
      }
    );
  }

  editGongdocCommand(gongdoccommandID: number, gongdoccommand: GongdocCommandDB) {

  }

  // display gongdoccommand in router
  displayGongdocCommandInRouter(gongdoccommandID: number) {
    this.router.navigate( ["gongdoccommand-display", gongdoccommandID])
  }

  // set editor outlet
  setEditorRouterOutlet(gongdoccommandID: number) {
    this.router.navigate([{
      outlets: {
        editor: ["gongdoccommand-detail", gongdoccommandID]
      }
    }]);
  }

  // set presentation outlet
  setPresentationRouterOutlet(gongdoccommandID: number) {
    this.router.navigate([{
      outlets: {
        presentation: ["gongdoccommand-presentation", gongdoccommandID]
      }
    }]);
  }

  /** Whether the number of selected elements matches the total number of rows. */
  isAllSelected() {
    const numSelected = this.selection.selected.length;
    const numRows = this.gongdoccommands.length;
    return numSelected === numRows;
  }

  /** Selects all rows if they are not all selected; otherwise clear selection. */
  masterToggle() {
    this.isAllSelected() ?
      this.selection.clear() :
      this.gongdoccommands.forEach(row => this.selection.select(row));
  }

  save() {

    let toUpdate = new Set<GongdocCommandDB>()

    // reset all initial selection of gongdoccommand that belong to aclass through Anarrayofb
    this.initialSelection.forEach(
      gongdoccommand => {
        gongdoccommand[this.dialogData.ReversePointer].Int64 = 0
        gongdoccommand[this.dialogData.ReversePointer].Valid = true
        toUpdate.add(gongdoccommand)
      }
    )

    // from selection, set gongdoccommand that belong to aclass through Anarrayofb
    this.selection.selected.forEach(
      gongdoccommand => {
        console.log("selection ID " + gongdoccommand.ID)
        let ID = +this.dialogData.ID
        gongdoccommand[this.dialogData.ReversePointer].Int64 = ID
        gongdoccommand[this.dialogData.ReversePointer].Valid = true
        toUpdate.add(gongdoccommand)
      }
    )

    // update all gongdoccommand (only update selection & initial selection)
    toUpdate.forEach(
      gongdoccommand => {
        this.gongdoccommandService.updateGongdocCommand(gongdoccommand)
          .subscribe(gongdoccommand => {
            this.gongdoccommandService.GongdocCommandServiceChanged.next("update")
            console.log("gongdoccommand saved")
          });
      }
    )
    this.dialogRef.close('Pizza!');
  }
}
