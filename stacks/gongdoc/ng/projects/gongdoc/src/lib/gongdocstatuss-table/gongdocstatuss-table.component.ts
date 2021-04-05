// generated from NgTableTemplateTS
import { Component, OnInit, OnChanges, Input, Output, EventEmitter, Inject, Optional } from '@angular/core';
import { MatTableDataSource } from '@angular/material/table';
import { MatButton } from '@angular/material/button'

import { MatDialogRef, MAT_DIALOG_DATA, MatDialog } from '@angular/material/dialog'
import { DialogData } from '../front-repo.service'
import { SelectionModel } from '@angular/cdk/collections';

const allowMultiSelect = true;

import { Router, RouterState } from '@angular/router';
import { GongdocStatusDB } from '../gongdocstatus-db'
import { GongdocStatusService } from '../gongdocstatus.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

// generated table component
@Component({
  selector: 'app-gongdocstatuss-table',
  templateUrl: './gongdocstatuss-table.component.html',
  styleUrls: ['./gongdocstatuss-table.component.css'],
})
export class GongdocStatussTableComponent implements OnInit {

  // used if the component is called as a selection component of GongdocStatus instances
  selection: SelectionModel<GongdocStatusDB>;
  initialSelection = new Array<GongdocStatusDB>();

  // the data source for the table
  gongdocstatuss: GongdocStatusDB[];

  // front repo, that will be referenced by this.gongdocstatuss
  frontRepo: FrontRepo

  // displayedColumns is referenced by the MatTable component for specify what columns
  // have to be displayed and in what order
  displayedColumns: string[];

  constructor(
    private gongdocstatusService: GongdocStatusService,
    private frontRepoService: FrontRepoService,

    // not null if the component is called as a selection component of gongdocstatus instances
    public dialogRef: MatDialogRef<GongdocStatussTableComponent>,
    @Optional() @Inject(MAT_DIALOG_DATA) public dialogData: DialogData,

    private router: Router,
  ) {
    // https://stackoverflow.com/questions/54627478/angular-7-routing-to-same-component-but-different-param-not-working
    // this is for routerLink on same component when only queryParameter changes
    this.router.routeReuseStrategy.shouldReuseRoute = function () {
      return false;
    };

    // observable for changes in structs
    this.gongdocstatusService.GongdocStatusServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.getGongdocStatuss()
        }
      }
    )
    if (dialogData == undefined) {
  	  this.displayedColumns = ['ID', 'Edit', 'Delete', // insertion point for columns to display
        "Name",
        "Status",
        "CommandCompletionDate",
      ]
    } else {
  	  this.displayedColumns = ['select', 'ID', // insertion point for columns to display
        "Name",
        "Status",
        "CommandCompletionDate",
      ]
      this.selection = new SelectionModel<GongdocStatusDB>(allowMultiSelect, this.initialSelection);
    }

  }

  ngOnInit(): void {
    this.getGongdocStatuss()
  }

  getGongdocStatuss(): void {
    this.frontRepoService.pull().subscribe(
      frontRepo => {
        this.frontRepo = frontRepo
        console.log("front repo pull returned")

        this.gongdocstatuss = this.frontRepo.GongdocStatuss_array;
        // in case the component is called as a selection component
        if (this.dialogData != undefined) {
          this.gongdocstatuss.forEach(
            gongdocstatus => {
              let ID = this.dialogData.ID
              let revPointer = gongdocstatus[this.dialogData.ReversePointer]
              if (revPointer.Int64 == ID) {
                this.initialSelection.push(gongdocstatus)
              }
            }
          )
          this.selection = new SelectionModel<GongdocStatusDB>(allowMultiSelect, this.initialSelection);
        }
      }
    )
  }

  // newGongdocStatus initiate a new gongdocstatus
  // create a new GongdocStatus objet
  newGongdocStatus() {
  }

  deleteGongdocStatus(gongdocstatusID: number, gongdocstatus: GongdocStatusDB) {
    // list of gongdocstatuss is truncated of gongdocstatus before the delete
    this.gongdocstatuss = this.gongdocstatuss.filter(h => h !== gongdocstatus);

    this.gongdocstatusService.deleteGongdocStatus(gongdocstatusID).subscribe(
      gongdocstatus => {
        this.gongdocstatusService.GongdocStatusServiceChanged.next("delete")

        console.log("gongdocstatus deleted")
      }
    );
  }

  editGongdocStatus(gongdocstatusID: number, gongdocstatus: GongdocStatusDB) {

  }

  // display gongdocstatus in router
  displayGongdocStatusInRouter(gongdocstatusID: number) {
    this.router.navigate( ["gongdocstatus-display", gongdocstatusID])
  }

  // set editor outlet
  setEditorRouterOutlet(gongdocstatusID: number) {
    this.router.navigate([{
      outlets: {
        editor: ["gongdocstatus-detail", gongdocstatusID]
      }
    }]);
  }

  // set presentation outlet
  setPresentationRouterOutlet(gongdocstatusID: number) {
    this.router.navigate([{
      outlets: {
        presentation: ["gongdocstatus-presentation", gongdocstatusID]
      }
    }]);
  }

  /** Whether the number of selected elements matches the total number of rows. */
  isAllSelected() {
    const numSelected = this.selection.selected.length;
    const numRows = this.gongdocstatuss.length;
    return numSelected === numRows;
  }

  /** Selects all rows if they are not all selected; otherwise clear selection. */
  masterToggle() {
    this.isAllSelected() ?
      this.selection.clear() :
      this.gongdocstatuss.forEach(row => this.selection.select(row));
  }

  save() {

    let toUpdate = new Set<GongdocStatusDB>()

    // reset all initial selection of gongdocstatus that belong to aclass through Anarrayofb
    this.initialSelection.forEach(
      gongdocstatus => {
        gongdocstatus[this.dialogData.ReversePointer].Int64 = 0
        gongdocstatus[this.dialogData.ReversePointer].Valid = true
        toUpdate.add(gongdocstatus)
      }
    )

    // from selection, set gongdocstatus that belong to aclass through Anarrayofb
    this.selection.selected.forEach(
      gongdocstatus => {
        console.log("selection ID " + gongdocstatus.ID)
        let ID = +this.dialogData.ID
        gongdocstatus[this.dialogData.ReversePointer].Int64 = ID
        gongdocstatus[this.dialogData.ReversePointer].Valid = true
        toUpdate.add(gongdocstatus)
      }
    )

    // update all gongdocstatus (only update selection & initial selection)
    toUpdate.forEach(
      gongdocstatus => {
        this.gongdocstatusService.updateGongdocStatus(gongdocstatus)
          .subscribe(gongdocstatus => {
            this.gongdocstatusService.GongdocStatusServiceChanged.next("update")
            console.log("gongdocstatus saved")
          });
      }
    )
    this.dialogRef.close('Pizza!');
  }
}
