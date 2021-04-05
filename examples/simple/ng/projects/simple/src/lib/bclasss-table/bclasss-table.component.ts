// generated from NgTableTemplateTS
import { Component, OnInit, OnChanges, Input, Output, EventEmitter, Inject, Optional } from '@angular/core';
import { MatTableDataSource } from '@angular/material/table';
import { MatButton } from '@angular/material/button'

import { MatDialogRef, MAT_DIALOG_DATA, MatDialog } from '@angular/material/dialog'
import { DialogData } from '../front-repo.service'
import { SelectionModel } from '@angular/cdk/collections';

const allowMultiSelect = true;

import { Router, RouterState } from '@angular/router';
import { BclassDB } from '../bclass-db'
import { BclassService } from '../bclass.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

// generated table component
@Component({
  selector: 'app-bclasss-table',
  templateUrl: './bclasss-table.component.html',
  styleUrls: ['./bclasss-table.component.css'],
})
export class BclasssTableComponent implements OnInit {

  // used if the component is called as a selection component of Bclass instances
  selection: SelectionModel<BclassDB>;
  initialSelection = new Array<BclassDB>();

  // the data source for the table
  bclasss: BclassDB[];

  // front repo, that will be referenced by this.bclasss
  frontRepo: FrontRepo

  // displayedColumns is referenced by the MatTable component for specify what columns
  // have to be displayed and in what order
  displayedColumns: string[];

  constructor(
    private bclassService: BclassService,
    private frontRepoService: FrontRepoService,

    // not null if the component is called as a selection component of bclass instances
    public dialogRef: MatDialogRef<BclasssTableComponent>,
    @Optional() @Inject(MAT_DIALOG_DATA) public dialogData: DialogData,

    private router: Router,
  ) {
    // https://stackoverflow.com/questions/54627478/angular-7-routing-to-same-component-but-different-param-not-working
    // this is for routerLink on same component when only queryParameter changes
    this.router.routeReuseStrategy.shouldReuseRoute = function () {
      return false;
    };

    // observable for changes in structs
    this.bclassService.BclassServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.getBclasss()
        }
      }
    )
    if (dialogData == undefined) {
  	  this.displayedColumns = ['ID', 'Edit', 'Delete', // insertion point for columns to display
        "Name",
        "Floatfield",
        "Intfield",
        "Anarrayofb",
        "Anotherarrayofb",
      ]
    } else {
  	  this.displayedColumns = ['select', 'ID', // insertion point for columns to display
        "Name",
        "Floatfield",
        "Intfield",
        "Anarrayofb",
        "Anotherarrayofb",
      ]
      this.selection = new SelectionModel<BclassDB>(allowMultiSelect, this.initialSelection);
    }

  }

  ngOnInit(): void {
    this.getBclasss()
  }

  getBclasss(): void {
    this.frontRepoService.pull().subscribe(
      frontRepo => {
        this.frontRepo = frontRepo
        console.log("front repo pull returned")

        this.bclasss = this.frontRepo.Bclasss_array;
        // in case the component is called as a selection component
        if (this.dialogData != undefined) {
          this.bclasss.forEach(
            bclass => {
              let ID = this.dialogData.ID
              let revPointer = bclass[this.dialogData.ReversePointer]
              if (revPointer.Int64 == ID) {
                this.initialSelection.push(bclass)
              }
            }
          )
          this.selection = new SelectionModel<BclassDB>(allowMultiSelect, this.initialSelection);
        }
      }
    )
  }

  // newBclass initiate a new bclass
  // create a new Bclass objet
  newBclass() {
  }

  deleteBclass(bclassID: number, bclass: BclassDB) {
    // list of bclasss is truncated of bclass before the delete
    this.bclasss = this.bclasss.filter(h => h !== bclass);

    this.bclassService.deleteBclass(bclassID).subscribe(
      bclass => {
        this.bclassService.BclassServiceChanged.next("delete")

        console.log("bclass deleted")
      }
    );
  }

  editBclass(bclassID: number, bclass: BclassDB) {

  }

  // display bclass in router
  displayBclassInRouter(bclassID: number) {
    this.router.navigate( ["bclass-display", bclassID])
  }

  // set editor outlet
  setEditorRouterOutlet(bclassID: number) {
    this.router.navigate([{
      outlets: {
        editor: ["bclass-detail", bclassID]
      }
    }]);
  }

  // set presentation outlet
  setPresentationRouterOutlet(bclassID: number) {
    this.router.navigate([{
      outlets: {
        presentation: ["bclass-presentation", bclassID]
      }
    }]);
  }

  /** Whether the number of selected elements matches the total number of rows. */
  isAllSelected() {
    const numSelected = this.selection.selected.length;
    const numRows = this.bclasss.length;
    return numSelected === numRows;
  }

  /** Selects all rows if they are not all selected; otherwise clear selection. */
  masterToggle() {
    this.isAllSelected() ?
      this.selection.clear() :
      this.bclasss.forEach(row => this.selection.select(row));
  }

  save() {

    let toUpdate = new Set<BclassDB>()

    // reset all initial selection of bclass that belong to aclass through Anarrayofb
    this.initialSelection.forEach(
      bclass => {
        bclass[this.dialogData.ReversePointer].Int64 = 0
        bclass[this.dialogData.ReversePointer].Valid = true
        toUpdate.add(bclass)
      }
    )

    // from selection, set bclass that belong to aclass through Anarrayofb
    this.selection.selected.forEach(
      bclass => {
        console.log("selection ID " + bclass.ID)
        let ID = +this.dialogData.ID
        bclass[this.dialogData.ReversePointer].Int64 = ID
        bclass[this.dialogData.ReversePointer].Valid = true
        toUpdate.add(bclass)
      }
    )

    // update all bclass (only update selection & initial selection)
    toUpdate.forEach(
      bclass => {
        this.bclassService.updateBclass(bclass)
          .subscribe(bclass => {
            this.bclassService.BclassServiceChanged.next("update")
            console.log("bclass saved")
          });
      }
    )
    this.dialogRef.close('Pizza!');
  }
}
