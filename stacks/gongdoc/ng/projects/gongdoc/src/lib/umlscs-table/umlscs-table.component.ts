// generated from NgTableTemplateTS
import { Component, OnInit, OnChanges, Input, Output, EventEmitter, Inject, Optional } from '@angular/core';
import { MatTableDataSource } from '@angular/material/table';
import { MatButton } from '@angular/material/button'

import { MatDialogRef, MAT_DIALOG_DATA, MatDialog } from '@angular/material/dialog'
import { DialogData } from '../front-repo.service'
import { SelectionModel } from '@angular/cdk/collections';

const allowMultiSelect = true;

import { Router, RouterState } from '@angular/router';
import { UmlscDB } from '../umlsc-db'
import { UmlscService } from '../umlsc.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

// generated table component
@Component({
  selector: 'app-umlscs-table',
  templateUrl: './umlscs-table.component.html',
  styleUrls: ['./umlscs-table.component.css'],
})
export class UmlscsTableComponent implements OnInit {

  // used if the component is called as a selection component of Umlsc instances
  selection: SelectionModel<UmlscDB>;
  initialSelection = new Array<UmlscDB>();

  // the data source for the table
  umlscs: UmlscDB[];

  // front repo, that will be referenced by this.umlscs
  frontRepo: FrontRepo

  // displayedColumns is referenced by the MatTable component for specify what columns
  // have to be displayed and in what order
  displayedColumns: string[];

  constructor(
    private umlscService: UmlscService,
    private frontRepoService: FrontRepoService,

    // not null if the component is called as a selection component of umlsc instances
    public dialogRef: MatDialogRef<UmlscsTableComponent>,
    @Optional() @Inject(MAT_DIALOG_DATA) public dialogData: DialogData,

    private router: Router,
  ) {
    // https://stackoverflow.com/questions/54627478/angular-7-routing-to-same-component-but-different-param-not-working
    // this is for routerLink on same component when only queryParameter changes
    this.router.routeReuseStrategy.shouldReuseRoute = function () {
      return false;
    };

    // observable for changes in structs
    this.umlscService.UmlscServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.getUmlscs()
        }
      }
    )
    if (dialogData == undefined) {
  	  this.displayedColumns = ['ID', 'Edit', 'Delete', // insertion point for columns to display
        "Name",
        "Activestate",
        "Umlscs",
      ]
    } else {
  	  this.displayedColumns = ['select', 'ID', // insertion point for columns to display
        "Name",
        "Activestate",
        "Umlscs",
      ]
      this.selection = new SelectionModel<UmlscDB>(allowMultiSelect, this.initialSelection);
    }

  }

  ngOnInit(): void {
    this.getUmlscs()
  }

  getUmlscs(): void {
    this.frontRepoService.pull().subscribe(
      frontRepo => {
        this.frontRepo = frontRepo
        console.log("front repo pull returned")

        this.umlscs = this.frontRepo.Umlscs_array;
        // in case the component is called as a selection component
        if (this.dialogData != undefined) {
          this.umlscs.forEach(
            umlsc => {
              let ID = this.dialogData.ID
              let revPointer = umlsc[this.dialogData.ReversePointer]
              if (revPointer.Int64 == ID) {
                this.initialSelection.push(umlsc)
              }
            }
          )
          this.selection = new SelectionModel<UmlscDB>(allowMultiSelect, this.initialSelection);
        }
      }
    )
  }

  // newUmlsc initiate a new umlsc
  // create a new Umlsc objet
  newUmlsc() {
  }

  deleteUmlsc(umlscID: number, umlsc: UmlscDB) {
    // list of umlscs is truncated of umlsc before the delete
    this.umlscs = this.umlscs.filter(h => h !== umlsc);

    this.umlscService.deleteUmlsc(umlscID).subscribe(
      umlsc => {
        this.umlscService.UmlscServiceChanged.next("delete")

        console.log("umlsc deleted")
      }
    );
  }

  editUmlsc(umlscID: number, umlsc: UmlscDB) {

  }

  // display umlsc in router
  displayUmlscInRouter(umlscID: number) {
    this.router.navigate( ["umlsc-display", umlscID])
  }

  // set editor outlet
  setEditorRouterOutlet(umlscID: number) {
    this.router.navigate([{
      outlets: {
        editor: ["umlsc-detail", umlscID]
      }
    }]);
  }

  // set presentation outlet
  setPresentationRouterOutlet(umlscID: number) {
    this.router.navigate([{
      outlets: {
        presentation: ["umlsc-presentation", umlscID]
      }
    }]);
  }

  /** Whether the number of selected elements matches the total number of rows. */
  isAllSelected() {
    const numSelected = this.selection.selected.length;
    const numRows = this.umlscs.length;
    return numSelected === numRows;
  }

  /** Selects all rows if they are not all selected; otherwise clear selection. */
  masterToggle() {
    this.isAllSelected() ?
      this.selection.clear() :
      this.umlscs.forEach(row => this.selection.select(row));
  }

  save() {

    let toUpdate = new Set<UmlscDB>()

    // reset all initial selection of umlsc that belong to aclass through Anarrayofb
    this.initialSelection.forEach(
      umlsc => {
        umlsc[this.dialogData.ReversePointer].Int64 = 0
        umlsc[this.dialogData.ReversePointer].Valid = true
        toUpdate.add(umlsc)
      }
    )

    // from selection, set umlsc that belong to aclass through Anarrayofb
    this.selection.selected.forEach(
      umlsc => {
        console.log("selection ID " + umlsc.ID)
        let ID = +this.dialogData.ID
        umlsc[this.dialogData.ReversePointer].Int64 = ID
        umlsc[this.dialogData.ReversePointer].Valid = true
        toUpdate.add(umlsc)
      }
    )

    // update all umlsc (only update selection & initial selection)
    toUpdate.forEach(
      umlsc => {
        this.umlscService.updateUmlsc(umlsc)
          .subscribe(umlsc => {
            this.umlscService.UmlscServiceChanged.next("update")
            console.log("umlsc saved")
          });
      }
    )
    this.dialogRef.close('Pizza!');
  }
}
