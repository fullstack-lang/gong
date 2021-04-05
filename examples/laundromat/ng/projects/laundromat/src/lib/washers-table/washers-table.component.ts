// generated from NgTableTemplateTS
import { Component, OnInit, OnChanges, Input, Output, EventEmitter, Inject, Optional } from '@angular/core';
import { MatTableDataSource } from '@angular/material/table';
import { MatButton } from '@angular/material/button'

import { MatDialogRef, MAT_DIALOG_DATA, MatDialog } from '@angular/material/dialog'
import { DialogData } from '../front-repo.service'
import { SelectionModel } from '@angular/cdk/collections';

const allowMultiSelect = true;

import { Router, RouterState } from '@angular/router';
import { WasherDB } from '../washer-db'
import { WasherService } from '../washer.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

// generated table component
@Component({
  selector: 'app-washers-table',
  templateUrl: './washers-table.component.html',
  styleUrls: ['./washers-table.component.css'],
})
export class WashersTableComponent implements OnInit {

  // used if the component is called as a selection component of Washer instances
  selection: SelectionModel<WasherDB>;
  initialSelection = new Array<WasherDB>();

  // the data source for the table
  washers: WasherDB[];

  // front repo, that will be referenced by this.washers
  frontRepo: FrontRepo

  // displayedColumns is referenced by the MatTable component for specify what columns
  // have to be displayed and in what order
  displayedColumns: string[];

  constructor(
    private washerService: WasherService,
    private frontRepoService: FrontRepoService,

    // not null if the component is called as a selection component of washer instances
    public dialogRef: MatDialogRef<WashersTableComponent>,
    @Optional() @Inject(MAT_DIALOG_DATA) public dialogData: DialogData,

    private router: Router,
  ) {
    // https://stackoverflow.com/questions/54627478/angular-7-routing-to-same-component-but-different-param-not-working
    // this is for routerLink on same component when only queryParameter changes
    this.router.routeReuseStrategy.shouldReuseRoute = function () {
      return false;
    };

    // observable for changes in structs
    this.washerService.WasherServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.getWashers()
        }
      }
    )
    if (dialogData == undefined) {
  	  this.displayedColumns = ['ID', 'Edit', 'Delete', // insertion point for columns to display
        "TechName",
        "Name",
        "DirtyLaundryWeight",
        "State",
        "Machine",
        "CleanedLaundryWeight",
      ]
    } else {
  	  this.displayedColumns = ['select', 'ID', // insertion point for columns to display
        "TechName",
        "Name",
        "DirtyLaundryWeight",
        "State",
        "Machine",
        "CleanedLaundryWeight",
      ]
      this.selection = new SelectionModel<WasherDB>(allowMultiSelect, this.initialSelection);
    }

  }

  ngOnInit(): void {
    this.getWashers()
  }

  getWashers(): void {
    this.frontRepoService.pull().subscribe(
      frontRepo => {
        this.frontRepo = frontRepo
        console.log("front repo pull returned")

        this.washers = this.frontRepo.Washers_array;
        // in case the component is called as a selection component
        if (this.dialogData != undefined) {
          this.washers.forEach(
            washer => {
              let ID = this.dialogData.ID
              let revPointer = washer[this.dialogData.ReversePointer]
              if (revPointer.Int64 == ID) {
                this.initialSelection.push(washer)
              }
            }
          )
          this.selection = new SelectionModel<WasherDB>(allowMultiSelect, this.initialSelection);
        }
      }
    )
  }

  // newWasher initiate a new washer
  // create a new Washer objet
  newWasher() {
  }

  deleteWasher(washerID: number, washer: WasherDB) {
    // list of washers is truncated of washer before the delete
    this.washers = this.washers.filter(h => h !== washer);

    this.washerService.deleteWasher(washerID).subscribe(
      washer => {
        this.washerService.WasherServiceChanged.next("delete")

        console.log("washer deleted")
      }
    );
  }

  editWasher(washerID: number, washer: WasherDB) {

  }

  // display washer in router
  displayWasherInRouter(washerID: number) {
    this.router.navigate( ["washer-display", washerID])
  }

  // set editor outlet
  setEditorRouterOutlet(washerID: number) {
    this.router.navigate([{
      outlets: {
        editor: ["washer-detail", washerID]
      }
    }]);
  }

  // set presentation outlet
  setPresentationRouterOutlet(washerID: number) {
    this.router.navigate([{
      outlets: {
        presentation: ["washer-presentation", washerID]
      }
    }]);
  }

  /** Whether the number of selected elements matches the total number of rows. */
  isAllSelected() {
    const numSelected = this.selection.selected.length;
    const numRows = this.washers.length;
    return numSelected === numRows;
  }

  /** Selects all rows if they are not all selected; otherwise clear selection. */
  masterToggle() {
    this.isAllSelected() ?
      this.selection.clear() :
      this.washers.forEach(row => this.selection.select(row));
  }

  save() {

    let toUpdate = new Set<WasherDB>()

    // reset all initial selection of washer that belong to aclass through Anarrayofb
    this.initialSelection.forEach(
      washer => {
        washer[this.dialogData.ReversePointer].Int64 = 0
        washer[this.dialogData.ReversePointer].Valid = true
        toUpdate.add(washer)
      }
    )

    // from selection, set washer that belong to aclass through Anarrayofb
    this.selection.selected.forEach(
      washer => {
        console.log("selection ID " + washer.ID)
        let ID = +this.dialogData.ID
        washer[this.dialogData.ReversePointer].Int64 = ID
        washer[this.dialogData.ReversePointer].Valid = true
        toUpdate.add(washer)
      }
    )

    // update all washer (only update selection & initial selection)
    toUpdate.forEach(
      washer => {
        this.washerService.updateWasher(washer)
          .subscribe(washer => {
            this.washerService.WasherServiceChanged.next("update")
            console.log("washer saved")
          });
      }
    )
    this.dialogRef.close('Pizza!');
  }
}
