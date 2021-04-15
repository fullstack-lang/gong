// generated from NgTableTemplateTS
import { Component, OnInit, OnChanges, Input, Output, EventEmitter, Inject, Optional } from '@angular/core';
import { MatTableDataSource } from '@angular/material/table';
import { MatButton } from '@angular/material/button'

import { MatDialogRef, MAT_DIALOG_DATA, MatDialog } from '@angular/material/dialog'
import { DialogData } from '../front-repo.service'
import { SelectionModel } from '@angular/cdk/collections';

const allowMultiSelect = true;

import { Router, RouterState } from '@angular/router';
import { PointerToGongStructFieldDB } from '../pointertogongstructfield-db'
import { PointerToGongStructFieldService } from '../pointertogongstructfield.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

// generated table component
@Component({
  selector: 'app-pointertogongstructfields-table',
  templateUrl: './pointertogongstructfields-table.component.html',
  styleUrls: ['./pointertogongstructfields-table.component.css'],
})
export class PointerToGongStructFieldsTableComponent implements OnInit {

  // used if the component is called as a selection component of PointerToGongStructField instances
  selection: SelectionModel<PointerToGongStructFieldDB>;
  initialSelection = new Array<PointerToGongStructFieldDB>();

  // the data source for the table
  pointertogongstructfields: PointerToGongStructFieldDB[];

  // front repo, that will be referenced by this.pointertogongstructfields
  frontRepo: FrontRepo

  // displayedColumns is referenced by the MatTable component for specify what columns
  // have to be displayed and in what order
  displayedColumns: string[];

  constructor(
    private pointertogongstructfieldService: PointerToGongStructFieldService,
    private frontRepoService: FrontRepoService,

    // not null if the component is called as a selection component of pointertogongstructfield instances
    public dialogRef: MatDialogRef<PointerToGongStructFieldsTableComponent>,
    @Optional() @Inject(MAT_DIALOG_DATA) public dialogData: DialogData,

    private router: Router,
  ) {
    // https://stackoverflow.com/questions/54627478/angular-7-routing-to-same-component-but-different-param-not-working
    // this is for routerLink on same component when only queryParameter changes
    this.router.routeReuseStrategy.shouldReuseRoute = function () {
      return false;
    };

    // observable for changes in structs
    this.pointertogongstructfieldService.PointerToGongStructFieldServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.getPointerToGongStructFields()
        }
      }
    )
    if (dialogData == undefined) {
      this.displayedColumns = ['ID', 'Edit', 'Delete', // insertion point for columns to display
        "Name",
        "GongStruct",
        "PointerToGongStructFields",
      ]
    } else {
      this.displayedColumns = ['select', 'ID', // insertion point for columns to display
        "Name",
        "GongStruct",
        "PointerToGongStructFields",
      ]
      this.selection = new SelectionModel<PointerToGongStructFieldDB>(allowMultiSelect, this.initialSelection);
    }

  }

  ngOnInit(): void {
    this.getPointerToGongStructFields()
  }

  getPointerToGongStructFields(): void {
    this.frontRepoService.pull().subscribe(
      frontRepo => {
        this.frontRepo = frontRepo
        console.log("front repo pull returned")

        this.pointertogongstructfields = this.frontRepo.PointerToGongStructFields_array;

        // insertion point for variables Recoveries

        // in case the component is called as a selection component
        if (this.dialogData != undefined) {
          this.pointertogongstructfields.forEach(
            pointertogongstructfield => {
              let ID = this.dialogData.ID
              let revPointer = pointertogongstructfield[this.dialogData.ReversePointer]
              if (revPointer.Int64 == ID) {
                this.initialSelection.push(pointertogongstructfield)
              }
            }
          )
          this.selection = new SelectionModel<PointerToGongStructFieldDB>(allowMultiSelect, this.initialSelection);
        }
      }
    )
  }

  // newPointerToGongStructField initiate a new pointertogongstructfield
  // create a new PointerToGongStructField objet
  newPointerToGongStructField() {
  }

  deletePointerToGongStructField(pointertogongstructfieldID: number, pointertogongstructfield: PointerToGongStructFieldDB) {
    // list of pointertogongstructfields is truncated of pointertogongstructfield before the delete
    this.pointertogongstructfields = this.pointertogongstructfields.filter(h => h !== pointertogongstructfield);

    this.pointertogongstructfieldService.deletePointerToGongStructField(pointertogongstructfieldID).subscribe(
      pointertogongstructfield => {
        this.pointertogongstructfieldService.PointerToGongStructFieldServiceChanged.next("delete")

        console.log("pointertogongstructfield deleted")
      }
    );
  }

  editPointerToGongStructField(pointertogongstructfieldID: number, pointertogongstructfield: PointerToGongStructFieldDB) {

  }

  // display pointertogongstructfield in router
  displayPointerToGongStructFieldInRouter(pointertogongstructfieldID: number) {
    this.router.navigate( ["pointertogongstructfield-display", pointertogongstructfieldID])
  }

  // set editor outlet
  setEditorRouterOutlet(pointertogongstructfieldID: number) {
    this.router.navigate([{
      outlets: {
        editor: ["pointertogongstructfield-detail", pointertogongstructfieldID]
      }
    }]);
  }

  // set presentation outlet
  setPresentationRouterOutlet(pointertogongstructfieldID: number) {
    this.router.navigate([{
      outlets: {
        presentation: ["pointertogongstructfield-presentation", pointertogongstructfieldID]
      }
    }]);
  }

  /** Whether the number of selected elements matches the total number of rows. */
  isAllSelected() {
    const numSelected = this.selection.selected.length;
    const numRows = this.pointertogongstructfields.length;
    return numSelected === numRows;
  }

  /** Selects all rows if they are not all selected; otherwise clear selection. */
  masterToggle() {
    this.isAllSelected() ?
      this.selection.clear() :
      this.pointertogongstructfields.forEach(row => this.selection.select(row));
  }

  save() {

    let toUpdate = new Set<PointerToGongStructFieldDB>()

    // reset all initial selection of pointertogongstructfield that belong to pointertogongstructfield through Anarrayofb
    this.initialSelection.forEach(
      pointertogongstructfield => {
        pointertogongstructfield[this.dialogData.ReversePointer].Int64 = 0
        pointertogongstructfield[this.dialogData.ReversePointer].Valid = true
        toUpdate.add(pointertogongstructfield)
      }
    )

    // from selection, set pointertogongstructfield that belong to pointertogongstructfield through Anarrayofb
    this.selection.selected.forEach(
      pointertogongstructfield => {
        console.log("selection ID " + pointertogongstructfield.ID)
        let ID = +this.dialogData.ID
        pointertogongstructfield[this.dialogData.ReversePointer].Int64 = ID
        pointertogongstructfield[this.dialogData.ReversePointer].Valid = true
        toUpdate.add(pointertogongstructfield)
      }
    )

    // update all pointertogongstructfield (only update selection & initial selection)
    toUpdate.forEach(
      pointertogongstructfield => {
        this.pointertogongstructfieldService.updatePointerToGongStructField(pointertogongstructfield)
          .subscribe(pointertogongstructfield => {
            this.pointertogongstructfieldService.PointerToGongStructFieldServiceChanged.next("update")
            console.log("pointertogongstructfield saved")
          });
      }
    )
    this.dialogRef.close('Pizza!');
  }
}
