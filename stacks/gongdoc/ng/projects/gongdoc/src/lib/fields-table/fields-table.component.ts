// generated from NgTableTemplateTS
import { Component, OnInit, OnChanges, Input, Output, EventEmitter, Inject, Optional } from '@angular/core';
import { MatTableDataSource } from '@angular/material/table';
import { MatButton } from '@angular/material/button'

import { MatDialogRef, MAT_DIALOG_DATA, MatDialog } from '@angular/material/dialog'
import { DialogData } from '../front-repo.service'
import { SelectionModel } from '@angular/cdk/collections';

const allowMultiSelect = true;

import { Router, RouterState } from '@angular/router';
import { FieldDB } from '../field-db'
import { FieldService } from '../field.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

// generated table component
@Component({
  selector: 'app-fields-table',
  templateUrl: './fields-table.component.html',
  styleUrls: ['./fields-table.component.css'],
})
export class FieldsTableComponent implements OnInit {

  // used if the component is called as a selection component of Field instances
  selection: SelectionModel<FieldDB>;
  initialSelection = new Array<FieldDB>();

  // the data source for the table
  fields: FieldDB[];

  // front repo, that will be referenced by this.fields
  frontRepo: FrontRepo

  // displayedColumns is referenced by the MatTable component for specify what columns
  // have to be displayed and in what order
  displayedColumns: string[];

  constructor(
    private fieldService: FieldService,
    private frontRepoService: FrontRepoService,

    // not null if the component is called as a selection component of field instances
    public dialogRef: MatDialogRef<FieldsTableComponent>,
    @Optional() @Inject(MAT_DIALOG_DATA) public dialogData: DialogData,

    private router: Router,
  ) {
    // https://stackoverflow.com/questions/54627478/angular-7-routing-to-same-component-but-different-param-not-working
    // this is for routerLink on same component when only queryParameter changes
    this.router.routeReuseStrategy.shouldReuseRoute = function () {
      return false;
    };

    // observable for changes in structs
    this.fieldService.FieldServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.getFields()
        }
      }
    )
    if (dialogData == undefined) {
  	  this.displayedColumns = ['ID', 'Edit', 'Delete', // insertion point for columns to display
        "Name",
        "Fieldname",
        "FieldTypeAsString",
        "Structname",
        "Fieldtypename",
        "Fields",
      ]
    } else {
  	  this.displayedColumns = ['select', 'ID', // insertion point for columns to display
        "Name",
        "Fieldname",
        "FieldTypeAsString",
        "Structname",
        "Fieldtypename",
        "Fields",
      ]
      this.selection = new SelectionModel<FieldDB>(allowMultiSelect, this.initialSelection);
    }

  }

  ngOnInit(): void {
    this.getFields()
  }

  getFields(): void {
    this.frontRepoService.pull().subscribe(
      frontRepo => {
        this.frontRepo = frontRepo
        console.log("front repo pull returned")

        this.fields = this.frontRepo.Fields_array;
        // in case the component is called as a selection component
        if (this.dialogData != undefined) {
          this.fields.forEach(
            field => {
              let ID = this.dialogData.ID
              let revPointer = field[this.dialogData.ReversePointer]
              if (revPointer.Int64 == ID) {
                this.initialSelection.push(field)
              }
            }
          )
          this.selection = new SelectionModel<FieldDB>(allowMultiSelect, this.initialSelection);
        }
      }
    )
  }

  // newField initiate a new field
  // create a new Field objet
  newField() {
  }

  deleteField(fieldID: number, field: FieldDB) {
    // list of fields is truncated of field before the delete
    this.fields = this.fields.filter(h => h !== field);

    this.fieldService.deleteField(fieldID).subscribe(
      field => {
        this.fieldService.FieldServiceChanged.next("delete")

        console.log("field deleted")
      }
    );
  }

  editField(fieldID: number, field: FieldDB) {

  }

  // display field in router
  displayFieldInRouter(fieldID: number) {
    this.router.navigate( ["field-display", fieldID])
  }

  // set editor outlet
  setEditorRouterOutlet(fieldID: number) {
    this.router.navigate([{
      outlets: {
        editor: ["field-detail", fieldID]
      }
    }]);
  }

  // set presentation outlet
  setPresentationRouterOutlet(fieldID: number) {
    this.router.navigate([{
      outlets: {
        presentation: ["field-presentation", fieldID]
      }
    }]);
  }

  /** Whether the number of selected elements matches the total number of rows. */
  isAllSelected() {
    const numSelected = this.selection.selected.length;
    const numRows = this.fields.length;
    return numSelected === numRows;
  }

  /** Selects all rows if they are not all selected; otherwise clear selection. */
  masterToggle() {
    this.isAllSelected() ?
      this.selection.clear() :
      this.fields.forEach(row => this.selection.select(row));
  }

  save() {

    let toUpdate = new Set<FieldDB>()

    // reset all initial selection of field that belong to aclass through Anarrayofb
    this.initialSelection.forEach(
      field => {
        field[this.dialogData.ReversePointer].Int64 = 0
        field[this.dialogData.ReversePointer].Valid = true
        toUpdate.add(field)
      }
    )

    // from selection, set field that belong to aclass through Anarrayofb
    this.selection.selected.forEach(
      field => {
        console.log("selection ID " + field.ID)
        let ID = +this.dialogData.ID
        field[this.dialogData.ReversePointer].Int64 = ID
        field[this.dialogData.ReversePointer].Valid = true
        toUpdate.add(field)
      }
    )

    // update all field (only update selection & initial selection)
    toUpdate.forEach(
      field => {
        this.fieldService.updateField(field)
          .subscribe(field => {
            this.fieldService.FieldServiceChanged.next("update")
            console.log("field saved")
          });
      }
    )
    this.dialogRef.close('Pizza!');
  }
}
