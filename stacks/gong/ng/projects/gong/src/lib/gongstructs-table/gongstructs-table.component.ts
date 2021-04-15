// generated from NgTableTemplateTS
import { Component, OnInit, OnChanges, Input, Output, EventEmitter, Inject, Optional } from '@angular/core';
import { MatTableDataSource } from '@angular/material/table';
import { MatButton } from '@angular/material/button'

import { MatDialogRef, MAT_DIALOG_DATA, MatDialog } from '@angular/material/dialog'
import { DialogData } from '../front-repo.service'
import { SelectionModel } from '@angular/cdk/collections';

const allowMultiSelect = true;

import { Router, RouterState } from '@angular/router';
import { GongStructDB } from '../gongstruct-db'
import { GongStructService } from '../gongstruct.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

// generated table component
@Component({
  selector: 'app-gongstructs-table',
  templateUrl: './gongstructs-table.component.html',
  styleUrls: ['./gongstructs-table.component.css'],
})
export class GongStructsTableComponent implements OnInit {

  // used if the component is called as a selection component of GongStruct instances
  selection: SelectionModel<GongStructDB>;
  initialSelection = new Array<GongStructDB>();

  // the data source for the table
  gongstructs: GongStructDB[];

  // front repo, that will be referenced by this.gongstructs
  frontRepo: FrontRepo

  // displayedColumns is referenced by the MatTable component for specify what columns
  // have to be displayed and in what order
  displayedColumns: string[];

  constructor(
    private gongstructService: GongStructService,
    private frontRepoService: FrontRepoService,

    // not null if the component is called as a selection component of gongstruct instances
    public dialogRef: MatDialogRef<GongStructsTableComponent>,
    @Optional() @Inject(MAT_DIALOG_DATA) public dialogData: DialogData,

    private router: Router,
  ) {
    // https://stackoverflow.com/questions/54627478/angular-7-routing-to-same-component-but-different-param-not-working
    // this is for routerLink on same component when only queryParameter changes
    this.router.routeReuseStrategy.shouldReuseRoute = function () {
      return false;
    };

    // observable for changes in structs
    this.gongstructService.GongStructServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.getGongStructs()
        }
      }
    )
    if (dialogData == undefined) {
      this.displayedColumns = ['ID', 'Edit', 'Delete', // insertion point for columns to display
        "Name",
      ]
    } else {
      this.displayedColumns = ['select', 'ID', // insertion point for columns to display
        "Name",
      ]
      this.selection = new SelectionModel<GongStructDB>(allowMultiSelect, this.initialSelection);
    }

  }

  ngOnInit(): void {
    this.getGongStructs()
  }

  getGongStructs(): void {
    this.frontRepoService.pull().subscribe(
      frontRepo => {
        this.frontRepo = frontRepo
        console.log("front repo pull returned")

        this.gongstructs = this.frontRepo.GongStructs_array;

        // insertion point for variables Recoveries

        // in case the component is called as a selection component
        if (this.dialogData != undefined) {
          this.gongstructs.forEach(
            gongstruct => {
              let ID = this.dialogData.ID
              let revPointer = gongstruct[this.dialogData.ReversePointer]
              if (revPointer.Int64 == ID) {
                this.initialSelection.push(gongstruct)
              }
            }
          )
          this.selection = new SelectionModel<GongStructDB>(allowMultiSelect, this.initialSelection);
        }
      }
    )
  }

  // newGongStruct initiate a new gongstruct
  // create a new GongStruct objet
  newGongStruct() {
  }

  deleteGongStruct(gongstructID: number, gongstruct: GongStructDB) {
    // list of gongstructs is truncated of gongstruct before the delete
    this.gongstructs = this.gongstructs.filter(h => h !== gongstruct);

    this.gongstructService.deleteGongStruct(gongstructID).subscribe(
      gongstruct => {
        this.gongstructService.GongStructServiceChanged.next("delete")

        console.log("gongstruct deleted")
      }
    );
  }

  editGongStruct(gongstructID: number, gongstruct: GongStructDB) {

  }

  // display gongstruct in router
  displayGongStructInRouter(gongstructID: number) {
    this.router.navigate( ["gongstruct-display", gongstructID])
  }

  // set editor outlet
  setEditorRouterOutlet(gongstructID: number) {
    this.router.navigate([{
      outlets: {
        editor: ["gongstruct-detail", gongstructID]
      }
    }]);
  }

  // set presentation outlet
  setPresentationRouterOutlet(gongstructID: number) {
    this.router.navigate([{
      outlets: {
        presentation: ["gongstruct-presentation", gongstructID]
      }
    }]);
  }

  /** Whether the number of selected elements matches the total number of rows. */
  isAllSelected() {
    const numSelected = this.selection.selected.length;
    const numRows = this.gongstructs.length;
    return numSelected === numRows;
  }

  /** Selects all rows if they are not all selected; otherwise clear selection. */
  masterToggle() {
    this.isAllSelected() ?
      this.selection.clear() :
      this.gongstructs.forEach(row => this.selection.select(row));
  }

  save() {

    let toUpdate = new Set<GongStructDB>()

    // reset all initial selection of gongstruct that belong to gongstruct through Anarrayofb
    this.initialSelection.forEach(
      gongstruct => {
        gongstruct[this.dialogData.ReversePointer].Int64 = 0
        gongstruct[this.dialogData.ReversePointer].Valid = true
        toUpdate.add(gongstruct)
      }
    )

    // from selection, set gongstruct that belong to gongstruct through Anarrayofb
    this.selection.selected.forEach(
      gongstruct => {
        console.log("selection ID " + gongstruct.ID)
        let ID = +this.dialogData.ID
        gongstruct[this.dialogData.ReversePointer].Int64 = ID
        gongstruct[this.dialogData.ReversePointer].Valid = true
        toUpdate.add(gongstruct)
      }
    )

    // update all gongstruct (only update selection & initial selection)
    toUpdate.forEach(
      gongstruct => {
        this.gongstructService.updateGongStruct(gongstruct)
          .subscribe(gongstruct => {
            this.gongstructService.GongStructServiceChanged.next("update")
            console.log("gongstruct saved")
          });
      }
    )
    this.dialogRef.close('Pizza!');
  }
}
