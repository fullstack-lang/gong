// generated from NgTableTemplateTS
import { Component, OnInit, OnChanges, Input, Output, EventEmitter, Inject, Optional } from '@angular/core';
import { MatTableDataSource } from '@angular/material/table';
import { MatButton } from '@angular/material/button'

import { MatDialogRef, MAT_DIALOG_DATA, MatDialog } from '@angular/material/dialog'
import { DialogData } from '../front-repo.service'
import { SelectionModel } from '@angular/cdk/collections';

const allowMultiSelect = true;

import { Router, RouterState } from '@angular/router';
import { GongEnumValueDB } from '../gongenumvalue-db'
import { GongEnumValueService } from '../gongenumvalue.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

// generated table component
@Component({
  selector: 'app-gongenumvalues-table',
  templateUrl: './gongenumvalues-table.component.html',
  styleUrls: ['./gongenumvalues-table.component.css'],
})
export class GongEnumValuesTableComponent implements OnInit {

  // used if the component is called as a selection component of GongEnumValue instances
  selection: SelectionModel<GongEnumValueDB>;
  initialSelection = new Array<GongEnumValueDB>();

  // the data source for the table
  gongenumvalues: GongEnumValueDB[];

  // front repo, that will be referenced by this.gongenumvalues
  frontRepo: FrontRepo

  // displayedColumns is referenced by the MatTable component for specify what columns
  // have to be displayed and in what order
  displayedColumns: string[];

  constructor(
    private gongenumvalueService: GongEnumValueService,
    private frontRepoService: FrontRepoService,

    // not null if the component is called as a selection component of gongenumvalue instances
    public dialogRef: MatDialogRef<GongEnumValuesTableComponent>,
    @Optional() @Inject(MAT_DIALOG_DATA) public dialogData: DialogData,

    private router: Router,
  ) {
    // https://stackoverflow.com/questions/54627478/angular-7-routing-to-same-component-but-different-param-not-working
    // this is for routerLink on same component when only queryParameter changes
    this.router.routeReuseStrategy.shouldReuseRoute = function () {
      return false;
    };

    // observable for changes in structs
    this.gongenumvalueService.GongEnumValueServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.getGongEnumValues()
        }
      }
    )
    if (dialogData == undefined) {
  	  this.displayedColumns = ['ID', 'Edit', 'Delete', // insertion point for columns to display
        "Name",
        "Value",
        "GongEnumValues",
      ]
    } else {
  	  this.displayedColumns = ['select', 'ID', // insertion point for columns to display
        "Name",
        "Value",
        "GongEnumValues",
      ]
      this.selection = new SelectionModel<GongEnumValueDB>(allowMultiSelect, this.initialSelection);
    }

  }

  ngOnInit(): void {
    this.getGongEnumValues()
  }

  getGongEnumValues(): void {
    this.frontRepoService.pull().subscribe(
      frontRepo => {
        this.frontRepo = frontRepo
        console.log("front repo pull returned")

        this.gongenumvalues = this.frontRepo.GongEnumValues_array;
        // in case the component is called as a selection component
        if (this.dialogData != undefined) {
          this.gongenumvalues.forEach(
            gongenumvalue => {
              let ID = this.dialogData.ID
              let revPointer = gongenumvalue[this.dialogData.ReversePointer]
              if (revPointer.Int64 == ID) {
                this.initialSelection.push(gongenumvalue)
              }
            }
          )
          this.selection = new SelectionModel<GongEnumValueDB>(allowMultiSelect, this.initialSelection);
        }
      }
    )
  }

  // newGongEnumValue initiate a new gongenumvalue
  // create a new GongEnumValue objet
  newGongEnumValue() {
  }

  deleteGongEnumValue(gongenumvalueID: number, gongenumvalue: GongEnumValueDB) {
    // list of gongenumvalues is truncated of gongenumvalue before the delete
    this.gongenumvalues = this.gongenumvalues.filter(h => h !== gongenumvalue);

    this.gongenumvalueService.deleteGongEnumValue(gongenumvalueID).subscribe(
      gongenumvalue => {
        this.gongenumvalueService.GongEnumValueServiceChanged.next("delete")

        console.log("gongenumvalue deleted")
      }
    );
  }

  editGongEnumValue(gongenumvalueID: number, gongenumvalue: GongEnumValueDB) {

  }

  // display gongenumvalue in router
  displayGongEnumValueInRouter(gongenumvalueID: number) {
    this.router.navigate( ["gongenumvalue-display", gongenumvalueID])
  }

  // set editor outlet
  setEditorRouterOutlet(gongenumvalueID: number) {
    this.router.navigate([{
      outlets: {
        editor: ["gongenumvalue-detail", gongenumvalueID]
      }
    }]);
  }

  // set presentation outlet
  setPresentationRouterOutlet(gongenumvalueID: number) {
    this.router.navigate([{
      outlets: {
        presentation: ["gongenumvalue-presentation", gongenumvalueID]
      }
    }]);
  }

  /** Whether the number of selected elements matches the total number of rows. */
  isAllSelected() {
    const numSelected = this.selection.selected.length;
    const numRows = this.gongenumvalues.length;
    return numSelected === numRows;
  }

  /** Selects all rows if they are not all selected; otherwise clear selection. */
  masterToggle() {
    this.isAllSelected() ?
      this.selection.clear() :
      this.gongenumvalues.forEach(row => this.selection.select(row));
  }

  save() {

    let toUpdate = new Set<GongEnumValueDB>()

    // reset all initial selection of gongenumvalue that belong to aclass through Anarrayofb
    this.initialSelection.forEach(
      gongenumvalue => {
        gongenumvalue[this.dialogData.ReversePointer].Int64 = 0
        gongenumvalue[this.dialogData.ReversePointer].Valid = true
        toUpdate.add(gongenumvalue)
      }
    )

    // from selection, set gongenumvalue that belong to aclass through Anarrayofb
    this.selection.selected.forEach(
      gongenumvalue => {
        console.log("selection ID " + gongenumvalue.ID)
        let ID = +this.dialogData.ID
        gongenumvalue[this.dialogData.ReversePointer].Int64 = ID
        gongenumvalue[this.dialogData.ReversePointer].Valid = true
        toUpdate.add(gongenumvalue)
      }
    )

    // update all gongenumvalue (only update selection & initial selection)
    toUpdate.forEach(
      gongenumvalue => {
        this.gongenumvalueService.updateGongEnumValue(gongenumvalue)
          .subscribe(gongenumvalue => {
            this.gongenumvalueService.GongEnumValueServiceChanged.next("update")
            console.log("gongenumvalue saved")
          });
      }
    )
    this.dialogRef.close('Pizza!');
  }
}
