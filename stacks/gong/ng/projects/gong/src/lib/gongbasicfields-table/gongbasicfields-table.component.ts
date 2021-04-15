// generated from NgTableTemplateTS
import { Component, OnInit, OnChanges, Input, Output, EventEmitter, Inject, Optional } from '@angular/core';
import { MatTableDataSource } from '@angular/material/table';
import { MatButton } from '@angular/material/button'

import { MatDialogRef, MAT_DIALOG_DATA, MatDialog } from '@angular/material/dialog'
import { DialogData } from '../front-repo.service'
import { SelectionModel } from '@angular/cdk/collections';

const allowMultiSelect = true;

import { Router, RouterState } from '@angular/router';
import { GongBasicFieldDB } from '../gongbasicfield-db'
import { GongBasicFieldService } from '../gongbasicfield.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

// generated table component
@Component({
  selector: 'app-gongbasicfields-table',
  templateUrl: './gongbasicfields-table.component.html',
  styleUrls: ['./gongbasicfields-table.component.css'],
})
export class GongBasicFieldsTableComponent implements OnInit {

  // used if the component is called as a selection component of GongBasicField instances
  selection: SelectionModel<GongBasicFieldDB>;
  initialSelection = new Array<GongBasicFieldDB>();

  // the data source for the table
  gongbasicfields: GongBasicFieldDB[];

  // front repo, that will be referenced by this.gongbasicfields
  frontRepo: FrontRepo

  // displayedColumns is referenced by the MatTable component for specify what columns
  // have to be displayed and in what order
  displayedColumns: string[];

  constructor(
    private gongbasicfieldService: GongBasicFieldService,
    private frontRepoService: FrontRepoService,

    // not null if the component is called as a selection component of gongbasicfield instances
    public dialogRef: MatDialogRef<GongBasicFieldsTableComponent>,
    @Optional() @Inject(MAT_DIALOG_DATA) public dialogData: DialogData,

    private router: Router,
  ) {
    // https://stackoverflow.com/questions/54627478/angular-7-routing-to-same-component-but-different-param-not-working
    // this is for routerLink on same component when only queryParameter changes
    this.router.routeReuseStrategy.shouldReuseRoute = function () {
      return false;
    };

    // observable for changes in structs
    this.gongbasicfieldService.GongBasicFieldServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.getGongBasicFields()
        }
      }
    )
    if (dialogData == undefined) {
      this.displayedColumns = ['ID', 'Edit', 'Delete', // insertion point for columns to display
        "Name",
        "BasicKindName",
        "GongEnum",
        "DeclaredType",
        "GongBasicFields",
      ]
    } else {
      this.displayedColumns = ['select', 'ID', // insertion point for columns to display
        "Name",
        "BasicKindName",
        "GongEnum",
        "DeclaredType",
        "GongBasicFields",
      ]
      this.selection = new SelectionModel<GongBasicFieldDB>(allowMultiSelect, this.initialSelection);
    }

  }

  ngOnInit(): void {
    this.getGongBasicFields()
  }

  getGongBasicFields(): void {
    this.frontRepoService.pull().subscribe(
      frontRepo => {
        this.frontRepo = frontRepo
        console.log("front repo pull returned")

        this.gongbasicfields = this.frontRepo.GongBasicFields_array;

        // insertion point for variables Recoveries

        // in case the component is called as a selection component
        if (this.dialogData != undefined) {
          this.gongbasicfields.forEach(
            gongbasicfield => {
              let ID = this.dialogData.ID
              let revPointer = gongbasicfield[this.dialogData.ReversePointer]
              if (revPointer.Int64 == ID) {
                this.initialSelection.push(gongbasicfield)
              }
            }
          )
          this.selection = new SelectionModel<GongBasicFieldDB>(allowMultiSelect, this.initialSelection);
        }
      }
    )
  }

  // newGongBasicField initiate a new gongbasicfield
  // create a new GongBasicField objet
  newGongBasicField() {
  }

  deleteGongBasicField(gongbasicfieldID: number, gongbasicfield: GongBasicFieldDB) {
    // list of gongbasicfields is truncated of gongbasicfield before the delete
    this.gongbasicfields = this.gongbasicfields.filter(h => h !== gongbasicfield);

    this.gongbasicfieldService.deleteGongBasicField(gongbasicfieldID).subscribe(
      gongbasicfield => {
        this.gongbasicfieldService.GongBasicFieldServiceChanged.next("delete")

        console.log("gongbasicfield deleted")
      }
    );
  }

  editGongBasicField(gongbasicfieldID: number, gongbasicfield: GongBasicFieldDB) {

  }

  // display gongbasicfield in router
  displayGongBasicFieldInRouter(gongbasicfieldID: number) {
    this.router.navigate( ["gongbasicfield-display", gongbasicfieldID])
  }

  // set editor outlet
  setEditorRouterOutlet(gongbasicfieldID: number) {
    this.router.navigate([{
      outlets: {
        editor: ["gongbasicfield-detail", gongbasicfieldID]
      }
    }]);
  }

  // set presentation outlet
  setPresentationRouterOutlet(gongbasicfieldID: number) {
    this.router.navigate([{
      outlets: {
        presentation: ["gongbasicfield-presentation", gongbasicfieldID]
      }
    }]);
  }

  /** Whether the number of selected elements matches the total number of rows. */
  isAllSelected() {
    const numSelected = this.selection.selected.length;
    const numRows = this.gongbasicfields.length;
    return numSelected === numRows;
  }

  /** Selects all rows if they are not all selected; otherwise clear selection. */
  masterToggle() {
    this.isAllSelected() ?
      this.selection.clear() :
      this.gongbasicfields.forEach(row => this.selection.select(row));
  }

  save() {

    let toUpdate = new Set<GongBasicFieldDB>()

    // reset all initial selection of gongbasicfield that belong to gongbasicfield through Anarrayofb
    this.initialSelection.forEach(
      gongbasicfield => {
        gongbasicfield[this.dialogData.ReversePointer].Int64 = 0
        gongbasicfield[this.dialogData.ReversePointer].Valid = true
        toUpdate.add(gongbasicfield)
      }
    )

    // from selection, set gongbasicfield that belong to gongbasicfield through Anarrayofb
    this.selection.selected.forEach(
      gongbasicfield => {
        console.log("selection ID " + gongbasicfield.ID)
        let ID = +this.dialogData.ID
        gongbasicfield[this.dialogData.ReversePointer].Int64 = ID
        gongbasicfield[this.dialogData.ReversePointer].Valid = true
        toUpdate.add(gongbasicfield)
      }
    )

    // update all gongbasicfield (only update selection & initial selection)
    toUpdate.forEach(
      gongbasicfield => {
        this.gongbasicfieldService.updateGongBasicField(gongbasicfield)
          .subscribe(gongbasicfield => {
            this.gongbasicfieldService.GongBasicFieldServiceChanged.next("update")
            console.log("gongbasicfield saved")
          });
      }
    )
    this.dialogRef.close('Pizza!');
  }
}
