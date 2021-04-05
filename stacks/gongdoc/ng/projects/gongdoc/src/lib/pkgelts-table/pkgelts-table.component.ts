// generated from NgTableTemplateTS
import { Component, OnInit, OnChanges, Input, Output, EventEmitter, Inject, Optional } from '@angular/core';
import { MatTableDataSource } from '@angular/material/table';
import { MatButton } from '@angular/material/button'

import { MatDialogRef, MAT_DIALOG_DATA, MatDialog } from '@angular/material/dialog'
import { DialogData } from '../front-repo.service'
import { SelectionModel } from '@angular/cdk/collections';

const allowMultiSelect = true;

import { Router, RouterState } from '@angular/router';
import { PkgeltDB } from '../pkgelt-db'
import { PkgeltService } from '../pkgelt.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

// generated table component
@Component({
  selector: 'app-pkgelts-table',
  templateUrl: './pkgelts-table.component.html',
  styleUrls: ['./pkgelts-table.component.css'],
})
export class PkgeltsTableComponent implements OnInit {

  // used if the component is called as a selection component of Pkgelt instances
  selection: SelectionModel<PkgeltDB>;
  initialSelection = new Array<PkgeltDB>();

  // the data source for the table
  pkgelts: PkgeltDB[];

  // front repo, that will be referenced by this.pkgelts
  frontRepo: FrontRepo

  // displayedColumns is referenced by the MatTable component for specify what columns
  // have to be displayed and in what order
  displayedColumns: string[];

  constructor(
    private pkgeltService: PkgeltService,
    private frontRepoService: FrontRepoService,

    // not null if the component is called as a selection component of pkgelt instances
    public dialogRef: MatDialogRef<PkgeltsTableComponent>,
    @Optional() @Inject(MAT_DIALOG_DATA) public dialogData: DialogData,

    private router: Router,
  ) {
    // https://stackoverflow.com/questions/54627478/angular-7-routing-to-same-component-but-different-param-not-working
    // this is for routerLink on same component when only queryParameter changes
    this.router.routeReuseStrategy.shouldReuseRoute = function () {
      return false;
    };

    // observable for changes in structs
    this.pkgeltService.PkgeltServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.getPkgelts()
        }
      }
    )
    if (dialogData == undefined) {
  	  this.displayedColumns = ['ID', 'Edit', 'Delete', // insertion point for columns to display
        "Name",
        "Path",
      ]
    } else {
  	  this.displayedColumns = ['select', 'ID', // insertion point for columns to display
        "Name",
        "Path",
      ]
      this.selection = new SelectionModel<PkgeltDB>(allowMultiSelect, this.initialSelection);
    }

  }

  ngOnInit(): void {
    this.getPkgelts()
  }

  getPkgelts(): void {
    this.frontRepoService.pull().subscribe(
      frontRepo => {
        this.frontRepo = frontRepo
        console.log("front repo pull returned")

        this.pkgelts = this.frontRepo.Pkgelts_array;
        // in case the component is called as a selection component
        if (this.dialogData != undefined) {
          this.pkgelts.forEach(
            pkgelt => {
              let ID = this.dialogData.ID
              let revPointer = pkgelt[this.dialogData.ReversePointer]
              if (revPointer.Int64 == ID) {
                this.initialSelection.push(pkgelt)
              }
            }
          )
          this.selection = new SelectionModel<PkgeltDB>(allowMultiSelect, this.initialSelection);
        }
      }
    )
  }

  // newPkgelt initiate a new pkgelt
  // create a new Pkgelt objet
  newPkgelt() {
  }

  deletePkgelt(pkgeltID: number, pkgelt: PkgeltDB) {
    // list of pkgelts is truncated of pkgelt before the delete
    this.pkgelts = this.pkgelts.filter(h => h !== pkgelt);

    this.pkgeltService.deletePkgelt(pkgeltID).subscribe(
      pkgelt => {
        this.pkgeltService.PkgeltServiceChanged.next("delete")

        console.log("pkgelt deleted")
      }
    );
  }

  editPkgelt(pkgeltID: number, pkgelt: PkgeltDB) {

  }

  // display pkgelt in router
  displayPkgeltInRouter(pkgeltID: number) {
    this.router.navigate( ["pkgelt-display", pkgeltID])
  }

  // set editor outlet
  setEditorRouterOutlet(pkgeltID: number) {
    this.router.navigate([{
      outlets: {
        editor: ["pkgelt-detail", pkgeltID]
      }
    }]);
  }

  // set presentation outlet
  setPresentationRouterOutlet(pkgeltID: number) {
    this.router.navigate([{
      outlets: {
        presentation: ["pkgelt-presentation", pkgeltID]
      }
    }]);
  }

  /** Whether the number of selected elements matches the total number of rows. */
  isAllSelected() {
    const numSelected = this.selection.selected.length;
    const numRows = this.pkgelts.length;
    return numSelected === numRows;
  }

  /** Selects all rows if they are not all selected; otherwise clear selection. */
  masterToggle() {
    this.isAllSelected() ?
      this.selection.clear() :
      this.pkgelts.forEach(row => this.selection.select(row));
  }

  save() {

    let toUpdate = new Set<PkgeltDB>()

    // reset all initial selection of pkgelt that belong to aclass through Anarrayofb
    this.initialSelection.forEach(
      pkgelt => {
        pkgelt[this.dialogData.ReversePointer].Int64 = 0
        pkgelt[this.dialogData.ReversePointer].Valid = true
        toUpdate.add(pkgelt)
      }
    )

    // from selection, set pkgelt that belong to aclass through Anarrayofb
    this.selection.selected.forEach(
      pkgelt => {
        console.log("selection ID " + pkgelt.ID)
        let ID = +this.dialogData.ID
        pkgelt[this.dialogData.ReversePointer].Int64 = ID
        pkgelt[this.dialogData.ReversePointer].Valid = true
        toUpdate.add(pkgelt)
      }
    )

    // update all pkgelt (only update selection & initial selection)
    toUpdate.forEach(
      pkgelt => {
        this.pkgeltService.updatePkgelt(pkgelt)
          .subscribe(pkgelt => {
            this.pkgeltService.PkgeltServiceChanged.next("update")
            console.log("pkgelt saved")
          });
      }
    )
    this.dialogRef.close('Pizza!');
  }
}
