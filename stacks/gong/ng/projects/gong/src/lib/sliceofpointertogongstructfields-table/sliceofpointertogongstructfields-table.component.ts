// generated from NgTableTemplateTS
import { Component, OnInit, OnChanges, Input, Output, EventEmitter, Inject, Optional } from '@angular/core';
import { MatTableDataSource } from '@angular/material/table';
import { MatButton } from '@angular/material/button'

import { MatDialogRef, MAT_DIALOG_DATA, MatDialog } from '@angular/material/dialog'
import { DialogData } from '../front-repo.service'
import { SelectionModel } from '@angular/cdk/collections';

const allowMultiSelect = true;

import { Router, RouterState } from '@angular/router';
import { SliceOfPointerToGongStructFieldDB } from '../sliceofpointertogongstructfield-db'
import { SliceOfPointerToGongStructFieldService } from '../sliceofpointertogongstructfield.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

// generated table component
@Component({
  selector: 'app-sliceofpointertogongstructfields-table',
  templateUrl: './sliceofpointertogongstructfields-table.component.html',
  styleUrls: ['./sliceofpointertogongstructfields-table.component.css'],
})
export class SliceOfPointerToGongStructFieldsTableComponent implements OnInit {

  // used if the component is called as a selection component of SliceOfPointerToGongStructField instances
  selection: SelectionModel<SliceOfPointerToGongStructFieldDB>;
  initialSelection = new Array<SliceOfPointerToGongStructFieldDB>();

  // the data source for the table
  sliceofpointertogongstructfields: SliceOfPointerToGongStructFieldDB[];

  // front repo, that will be referenced by this.sliceofpointertogongstructfields
  frontRepo: FrontRepo

  // displayedColumns is referenced by the MatTable component for specify what columns
  // have to be displayed and in what order
  displayedColumns: string[];

  constructor(
    private sliceofpointertogongstructfieldService: SliceOfPointerToGongStructFieldService,
    private frontRepoService: FrontRepoService,

    // not null if the component is called as a selection component of sliceofpointertogongstructfield instances
    public dialogRef: MatDialogRef<SliceOfPointerToGongStructFieldsTableComponent>,
    @Optional() @Inject(MAT_DIALOG_DATA) public dialogData: DialogData,

    private router: Router,
  ) {
    // https://stackoverflow.com/questions/54627478/angular-7-routing-to-same-component-but-different-param-not-working
    // this is for routerLink on same component when only queryParameter changes
    this.router.routeReuseStrategy.shouldReuseRoute = function () {
      return false;
    };

    // observable for changes in structs
    this.sliceofpointertogongstructfieldService.SliceOfPointerToGongStructFieldServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.getSliceOfPointerToGongStructFields()
        }
      }
    )
    if (dialogData == undefined) {
  	  this.displayedColumns = ['ID', 'Edit', 'Delete', // insertion point for columns to display
        "Name",
        "GongStruct",
        "SliceOfPointerToGongStructFields",
      ]
    } else {
  	  this.displayedColumns = ['select', 'ID', // insertion point for columns to display
        "Name",
        "GongStruct",
        "SliceOfPointerToGongStructFields",
      ]
      this.selection = new SelectionModel<SliceOfPointerToGongStructFieldDB>(allowMultiSelect, this.initialSelection);
    }

  }

  ngOnInit(): void {
    this.getSliceOfPointerToGongStructFields()
  }

  getSliceOfPointerToGongStructFields(): void {
    this.frontRepoService.pull().subscribe(
      frontRepo => {
        this.frontRepo = frontRepo
        console.log("front repo pull returned")

        this.sliceofpointertogongstructfields = this.frontRepo.SliceOfPointerToGongStructFields_array;
        // in case the component is called as a selection component
        if (this.dialogData != undefined) {
          this.sliceofpointertogongstructfields.forEach(
            sliceofpointertogongstructfield => {
              let ID = this.dialogData.ID
              let revPointer = sliceofpointertogongstructfield[this.dialogData.ReversePointer]
              if (revPointer.Int64 == ID) {
                this.initialSelection.push(sliceofpointertogongstructfield)
              }
            }
          )
          this.selection = new SelectionModel<SliceOfPointerToGongStructFieldDB>(allowMultiSelect, this.initialSelection);
        }
      }
    )
  }

  // newSliceOfPointerToGongStructField initiate a new sliceofpointertogongstructfield
  // create a new SliceOfPointerToGongStructField objet
  newSliceOfPointerToGongStructField() {
  }

  deleteSliceOfPointerToGongStructField(sliceofpointertogongstructfieldID: number, sliceofpointertogongstructfield: SliceOfPointerToGongStructFieldDB) {
    // list of sliceofpointertogongstructfields is truncated of sliceofpointertogongstructfield before the delete
    this.sliceofpointertogongstructfields = this.sliceofpointertogongstructfields.filter(h => h !== sliceofpointertogongstructfield);

    this.sliceofpointertogongstructfieldService.deleteSliceOfPointerToGongStructField(sliceofpointertogongstructfieldID).subscribe(
      sliceofpointertogongstructfield => {
        this.sliceofpointertogongstructfieldService.SliceOfPointerToGongStructFieldServiceChanged.next("delete")

        console.log("sliceofpointertogongstructfield deleted")
      }
    );
  }

  editSliceOfPointerToGongStructField(sliceofpointertogongstructfieldID: number, sliceofpointertogongstructfield: SliceOfPointerToGongStructFieldDB) {

  }

  // display sliceofpointertogongstructfield in router
  displaySliceOfPointerToGongStructFieldInRouter(sliceofpointertogongstructfieldID: number) {
    this.router.navigate( ["sliceofpointertogongstructfield-display", sliceofpointertogongstructfieldID])
  }

  // set editor outlet
  setEditorRouterOutlet(sliceofpointertogongstructfieldID: number) {
    this.router.navigate([{
      outlets: {
        editor: ["sliceofpointertogongstructfield-detail", sliceofpointertogongstructfieldID]
      }
    }]);
  }

  // set presentation outlet
  setPresentationRouterOutlet(sliceofpointertogongstructfieldID: number) {
    this.router.navigate([{
      outlets: {
        presentation: ["sliceofpointertogongstructfield-presentation", sliceofpointertogongstructfieldID]
      }
    }]);
  }

  /** Whether the number of selected elements matches the total number of rows. */
  isAllSelected() {
    const numSelected = this.selection.selected.length;
    const numRows = this.sliceofpointertogongstructfields.length;
    return numSelected === numRows;
  }

  /** Selects all rows if they are not all selected; otherwise clear selection. */
  masterToggle() {
    this.isAllSelected() ?
      this.selection.clear() :
      this.sliceofpointertogongstructfields.forEach(row => this.selection.select(row));
  }

  save() {

    let toUpdate = new Set<SliceOfPointerToGongStructFieldDB>()

    // reset all initial selection of sliceofpointertogongstructfield that belong to aclass through Anarrayofb
    this.initialSelection.forEach(
      sliceofpointertogongstructfield => {
        sliceofpointertogongstructfield[this.dialogData.ReversePointer].Int64 = 0
        sliceofpointertogongstructfield[this.dialogData.ReversePointer].Valid = true
        toUpdate.add(sliceofpointertogongstructfield)
      }
    )

    // from selection, set sliceofpointertogongstructfield that belong to aclass through Anarrayofb
    this.selection.selected.forEach(
      sliceofpointertogongstructfield => {
        console.log("selection ID " + sliceofpointertogongstructfield.ID)
        let ID = +this.dialogData.ID
        sliceofpointertogongstructfield[this.dialogData.ReversePointer].Int64 = ID
        sliceofpointertogongstructfield[this.dialogData.ReversePointer].Valid = true
        toUpdate.add(sliceofpointertogongstructfield)
      }
    )

    // update all sliceofpointertogongstructfield (only update selection & initial selection)
    toUpdate.forEach(
      sliceofpointertogongstructfield => {
        this.sliceofpointertogongstructfieldService.updateSliceOfPointerToGongStructField(sliceofpointertogongstructfield)
          .subscribe(sliceofpointertogongstructfield => {
            this.sliceofpointertogongstructfieldService.SliceOfPointerToGongStructFieldServiceChanged.next("update")
            console.log("sliceofpointertogongstructfield saved")
          });
      }
    )
    this.dialogRef.close('Pizza!');
  }
}
