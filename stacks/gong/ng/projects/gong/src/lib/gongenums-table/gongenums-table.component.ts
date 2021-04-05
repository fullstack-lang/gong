// generated from NgTableTemplateTS
import { Component, OnInit, OnChanges, Input, Output, EventEmitter, Inject, Optional } from '@angular/core';
import { MatTableDataSource } from '@angular/material/table';
import { MatButton } from '@angular/material/button'

import { MatDialogRef, MAT_DIALOG_DATA, MatDialog } from '@angular/material/dialog'
import { DialogData } from '../front-repo.service'
import { SelectionModel } from '@angular/cdk/collections';

const allowMultiSelect = true;

import { Router, RouterState } from '@angular/router';
import { GongEnumDB } from '../gongenum-db'
import { GongEnumService } from '../gongenum.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

// generated table component
@Component({
  selector: 'app-gongenums-table',
  templateUrl: './gongenums-table.component.html',
  styleUrls: ['./gongenums-table.component.css'],
})
export class GongEnumsTableComponent implements OnInit {

  // used if the component is called as a selection component of GongEnum instances
  selection: SelectionModel<GongEnumDB>;
  initialSelection = new Array<GongEnumDB>();

  // the data source for the table
  gongenums: GongEnumDB[];

  // front repo, that will be referenced by this.gongenums
  frontRepo: FrontRepo

  // displayedColumns is referenced by the MatTable component for specify what columns
  // have to be displayed and in what order
  displayedColumns: string[];

  constructor(
    private gongenumService: GongEnumService,
    private frontRepoService: FrontRepoService,

    // not null if the component is called as a selection component of gongenum instances
    public dialogRef: MatDialogRef<GongEnumsTableComponent>,
    @Optional() @Inject(MAT_DIALOG_DATA) public dialogData: DialogData,

    private router: Router,
  ) {
    // https://stackoverflow.com/questions/54627478/angular-7-routing-to-same-component-but-different-param-not-working
    // this is for routerLink on same component when only queryParameter changes
    this.router.routeReuseStrategy.shouldReuseRoute = function () {
      return false;
    };

    // observable for changes in structs
    this.gongenumService.GongEnumServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.getGongEnums()
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
      this.selection = new SelectionModel<GongEnumDB>(allowMultiSelect, this.initialSelection);
    }

  }

  ngOnInit(): void {
    this.getGongEnums()
  }

  getGongEnums(): void {
    this.frontRepoService.pull().subscribe(
      frontRepo => {
        this.frontRepo = frontRepo
        console.log("front repo pull returned")

        this.gongenums = this.frontRepo.GongEnums_array;
        // in case the component is called as a selection component
        if (this.dialogData != undefined) {
          this.gongenums.forEach(
            gongenum => {
              let ID = this.dialogData.ID
              let revPointer = gongenum[this.dialogData.ReversePointer]
              if (revPointer.Int64 == ID) {
                this.initialSelection.push(gongenum)
              }
            }
          )
          this.selection = new SelectionModel<GongEnumDB>(allowMultiSelect, this.initialSelection);
        }
      }
    )
  }

  // newGongEnum initiate a new gongenum
  // create a new GongEnum objet
  newGongEnum() {
  }

  deleteGongEnum(gongenumID: number, gongenum: GongEnumDB) {
    // list of gongenums is truncated of gongenum before the delete
    this.gongenums = this.gongenums.filter(h => h !== gongenum);

    this.gongenumService.deleteGongEnum(gongenumID).subscribe(
      gongenum => {
        this.gongenumService.GongEnumServiceChanged.next("delete")

        console.log("gongenum deleted")
      }
    );
  }

  editGongEnum(gongenumID: number, gongenum: GongEnumDB) {

  }

  // display gongenum in router
  displayGongEnumInRouter(gongenumID: number) {
    this.router.navigate( ["gongenum-display", gongenumID])
  }

  // set editor outlet
  setEditorRouterOutlet(gongenumID: number) {
    this.router.navigate([{
      outlets: {
        editor: ["gongenum-detail", gongenumID]
      }
    }]);
  }

  // set presentation outlet
  setPresentationRouterOutlet(gongenumID: number) {
    this.router.navigate([{
      outlets: {
        presentation: ["gongenum-presentation", gongenumID]
      }
    }]);
  }

  /** Whether the number of selected elements matches the total number of rows. */
  isAllSelected() {
    const numSelected = this.selection.selected.length;
    const numRows = this.gongenums.length;
    return numSelected === numRows;
  }

  /** Selects all rows if they are not all selected; otherwise clear selection. */
  masterToggle() {
    this.isAllSelected() ?
      this.selection.clear() :
      this.gongenums.forEach(row => this.selection.select(row));
  }

  save() {

    let toUpdate = new Set<GongEnumDB>()

    // reset all initial selection of gongenum that belong to aclass through Anarrayofb
    this.initialSelection.forEach(
      gongenum => {
        gongenum[this.dialogData.ReversePointer].Int64 = 0
        gongenum[this.dialogData.ReversePointer].Valid = true
        toUpdate.add(gongenum)
      }
    )

    // from selection, set gongenum that belong to aclass through Anarrayofb
    this.selection.selected.forEach(
      gongenum => {
        console.log("selection ID " + gongenum.ID)
        let ID = +this.dialogData.ID
        gongenum[this.dialogData.ReversePointer].Int64 = ID
        gongenum[this.dialogData.ReversePointer].Valid = true
        toUpdate.add(gongenum)
      }
    )

    // update all gongenum (only update selection & initial selection)
    toUpdate.forEach(
      gongenum => {
        this.gongenumService.updateGongEnum(gongenum)
          .subscribe(gongenum => {
            this.gongenumService.GongEnumServiceChanged.next("update")
            console.log("gongenum saved")
          });
      }
    )
    this.dialogRef.close('Pizza!');
  }
}
