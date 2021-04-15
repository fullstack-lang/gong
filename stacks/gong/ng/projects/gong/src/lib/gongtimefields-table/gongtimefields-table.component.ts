// generated from NgTableTemplateTS
import { Component, OnInit, OnChanges, Input, Output, EventEmitter, Inject, Optional } from '@angular/core';
import { MatTableDataSource } from '@angular/material/table';
import { MatButton } from '@angular/material/button'

import { MatDialogRef, MAT_DIALOG_DATA, MatDialog } from '@angular/material/dialog'
import { DialogData } from '../front-repo.service'
import { SelectionModel } from '@angular/cdk/collections';

const allowMultiSelect = true;

import { Router, RouterState } from '@angular/router';
import { GongTimeFieldDB } from '../gongtimefield-db'
import { GongTimeFieldService } from '../gongtimefield.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

// generated table component
@Component({
  selector: 'app-gongtimefields-table',
  templateUrl: './gongtimefields-table.component.html',
  styleUrls: ['./gongtimefields-table.component.css'],
})
export class GongTimeFieldsTableComponent implements OnInit {

  // used if the component is called as a selection component of GongTimeField instances
  selection: SelectionModel<GongTimeFieldDB>;
  initialSelection = new Array<GongTimeFieldDB>();

  // the data source for the table
  gongtimefields: GongTimeFieldDB[];

  // front repo, that will be referenced by this.gongtimefields
  frontRepo: FrontRepo

  // displayedColumns is referenced by the MatTable component for specify what columns
  // have to be displayed and in what order
  displayedColumns: string[];

  constructor(
    private gongtimefieldService: GongTimeFieldService,
    private frontRepoService: FrontRepoService,

    // not null if the component is called as a selection component of gongtimefield instances
    public dialogRef: MatDialogRef<GongTimeFieldsTableComponent>,
    @Optional() @Inject(MAT_DIALOG_DATA) public dialogData: DialogData,

    private router: Router,
  ) {
    // https://stackoverflow.com/questions/54627478/angular-7-routing-to-same-component-but-different-param-not-working
    // this is for routerLink on same component when only queryParameter changes
    this.router.routeReuseStrategy.shouldReuseRoute = function () {
      return false;
    };

    // observable for changes in structs
    this.gongtimefieldService.GongTimeFieldServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.getGongTimeFields()
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
      this.selection = new SelectionModel<GongTimeFieldDB>(allowMultiSelect, this.initialSelection);
    }

  }

  ngOnInit(): void {
    this.getGongTimeFields()
  }

  getGongTimeFields(): void {
    this.frontRepoService.pull().subscribe(
      frontRepo => {
        this.frontRepo = frontRepo
        console.log("front repo pull returned")

        this.gongtimefields = this.frontRepo.GongTimeFields_array;

        // insertion point for variables Recoveries

        // in case the component is called as a selection component
        if (this.dialogData != undefined) {
          this.gongtimefields.forEach(
            gongtimefield => {
              let ID = this.dialogData.ID
              let revPointer = gongtimefield[this.dialogData.ReversePointer]
              if (revPointer.Int64 == ID) {
                this.initialSelection.push(gongtimefield)
              }
            }
          )
          this.selection = new SelectionModel<GongTimeFieldDB>(allowMultiSelect, this.initialSelection);
        }
      }
    )
  }

  // newGongTimeField initiate a new gongtimefield
  // create a new GongTimeField objet
  newGongTimeField() {
  }

  deleteGongTimeField(gongtimefieldID: number, gongtimefield: GongTimeFieldDB) {
    // list of gongtimefields is truncated of gongtimefield before the delete
    this.gongtimefields = this.gongtimefields.filter(h => h !== gongtimefield);

    this.gongtimefieldService.deleteGongTimeField(gongtimefieldID).subscribe(
      gongtimefield => {
        this.gongtimefieldService.GongTimeFieldServiceChanged.next("delete")

        console.log("gongtimefield deleted")
      }
    );
  }

  editGongTimeField(gongtimefieldID: number, gongtimefield: GongTimeFieldDB) {

  }

  // display gongtimefield in router
  displayGongTimeFieldInRouter(gongtimefieldID: number) {
    this.router.navigate( ["gongtimefield-display", gongtimefieldID])
  }

  // set editor outlet
  setEditorRouterOutlet(gongtimefieldID: number) {
    this.router.navigate([{
      outlets: {
        editor: ["gongtimefield-detail", gongtimefieldID]
      }
    }]);
  }

  // set presentation outlet
  setPresentationRouterOutlet(gongtimefieldID: number) {
    this.router.navigate([{
      outlets: {
        presentation: ["gongtimefield-presentation", gongtimefieldID]
      }
    }]);
  }

  /** Whether the number of selected elements matches the total number of rows. */
  isAllSelected() {
    const numSelected = this.selection.selected.length;
    const numRows = this.gongtimefields.length;
    return numSelected === numRows;
  }

  /** Selects all rows if they are not all selected; otherwise clear selection. */
  masterToggle() {
    this.isAllSelected() ?
      this.selection.clear() :
      this.gongtimefields.forEach(row => this.selection.select(row));
  }

  save() {

    let toUpdate = new Set<GongTimeFieldDB>()

    // reset all initial selection of gongtimefield that belong to gongtimefield through Anarrayofb
    this.initialSelection.forEach(
      gongtimefield => {
        gongtimefield[this.dialogData.ReversePointer].Int64 = 0
        gongtimefield[this.dialogData.ReversePointer].Valid = true
        toUpdate.add(gongtimefield)
      }
    )

    // from selection, set gongtimefield that belong to gongtimefield through Anarrayofb
    this.selection.selected.forEach(
      gongtimefield => {
        console.log("selection ID " + gongtimefield.ID)
        let ID = +this.dialogData.ID
        gongtimefield[this.dialogData.ReversePointer].Int64 = ID
        gongtimefield[this.dialogData.ReversePointer].Valid = true
        toUpdate.add(gongtimefield)
      }
    )

    // update all gongtimefield (only update selection & initial selection)
    toUpdate.forEach(
      gongtimefield => {
        this.gongtimefieldService.updateGongTimeField(gongtimefield)
          .subscribe(gongtimefield => {
            this.gongtimefieldService.GongTimeFieldServiceChanged.next("update")
            console.log("gongtimefield saved")
          });
      }
    )
    this.dialogRef.close('Pizza!');
  }
}
