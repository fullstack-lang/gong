// generated from NgTableTemplateTS
import { Component, OnInit, OnChanges, Input, Output, EventEmitter, Inject, Optional } from '@angular/core';
import { MatTableDataSource } from '@angular/material/table';
import { MatButton } from '@angular/material/button'

import { MatDialogRef, MAT_DIALOG_DATA, MatDialog } from '@angular/material/dialog'
import { DialogData } from '../front-repo.service'
import { SelectionModel } from '@angular/cdk/collections';

const allowMultiSelect = true;

import { Router, RouterState } from '@angular/router';
import { GongsimStatusDB } from '../gongsimstatus-db'
import { GongsimStatusService } from '../gongsimstatus.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

// generated table component
@Component({
  selector: 'app-gongsimstatuss-table',
  templateUrl: './gongsimstatuss-table.component.html',
  styleUrls: ['./gongsimstatuss-table.component.css'],
})
export class GongsimStatussTableComponent implements OnInit {

  // used if the component is called as a selection component of GongsimStatus instances
  selection: SelectionModel<GongsimStatusDB>;
  initialSelection = new Array<GongsimStatusDB>();

  // the data source for the table
  gongsimstatuss: GongsimStatusDB[];

  // front repo, that will be referenced by this.gongsimstatuss
  frontRepo: FrontRepo

  // displayedColumns is referenced by the MatTable component for specify what columns
  // have to be displayed and in what order
  displayedColumns: string[];

  constructor(
    private gongsimstatusService: GongsimStatusService,
    private frontRepoService: FrontRepoService,

    // not null if the component is called as a selection component of gongsimstatus instances
    public dialogRef: MatDialogRef<GongsimStatussTableComponent>,
    @Optional() @Inject(MAT_DIALOG_DATA) public dialogData: DialogData,

    private router: Router,
  ) {
    // https://stackoverflow.com/questions/54627478/angular-7-routing-to-same-component-but-different-param-not-working
    // this is for routerLink on same component when only queryParameter changes
    this.router.routeReuseStrategy.shouldReuseRoute = function () {
      return false;
    };

    // observable for changes in structs
    this.gongsimstatusService.GongsimStatusServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.getGongsimStatuss()
        }
      }
    )
    if (dialogData == undefined) {
  	  this.displayedColumns = ['ID', 'Edit', 'Delete', // insertion point for columns to display
        "Name",
        "CurrentCommand",
        "CompletionDate",
        "CurrentSpeedCommand",
        "SpeedCommandCompletionDate",
      ]
    } else {
  	  this.displayedColumns = ['select', 'ID', // insertion point for columns to display
        "Name",
        "CurrentCommand",
        "CompletionDate",
        "CurrentSpeedCommand",
        "SpeedCommandCompletionDate",
      ]
      this.selection = new SelectionModel<GongsimStatusDB>(allowMultiSelect, this.initialSelection);
    }

  }

  ngOnInit(): void {
    this.getGongsimStatuss()
  }

  getGongsimStatuss(): void {
    this.frontRepoService.pull().subscribe(
      frontRepo => {
        this.frontRepo = frontRepo
        console.log("front repo pull returned")

        this.gongsimstatuss = this.frontRepo.GongsimStatuss_array;
        // in case the component is called as a selection component
        if (this.dialogData != undefined) {
          this.gongsimstatuss.forEach(
            gongsimstatus => {
              let ID = this.dialogData.ID
              let revPointer = gongsimstatus[this.dialogData.ReversePointer]
              if (revPointer.Int64 == ID) {
                this.initialSelection.push(gongsimstatus)
              }
            }
          )
          this.selection = new SelectionModel<GongsimStatusDB>(allowMultiSelect, this.initialSelection);
        }
      }
    )
  }

  // newGongsimStatus initiate a new gongsimstatus
  // create a new GongsimStatus objet
  newGongsimStatus() {
  }

  deleteGongsimStatus(gongsimstatusID: number, gongsimstatus: GongsimStatusDB) {
    // list of gongsimstatuss is truncated of gongsimstatus before the delete
    this.gongsimstatuss = this.gongsimstatuss.filter(h => h !== gongsimstatus);

    this.gongsimstatusService.deleteGongsimStatus(gongsimstatusID).subscribe(
      gongsimstatus => {
        this.gongsimstatusService.GongsimStatusServiceChanged.next("delete")

        console.log("gongsimstatus deleted")
      }
    );
  }

  editGongsimStatus(gongsimstatusID: number, gongsimstatus: GongsimStatusDB) {

  }

  // display gongsimstatus in router
  displayGongsimStatusInRouter(gongsimstatusID: number) {
    this.router.navigate( ["gongsimstatus-display", gongsimstatusID])
  }

  // set editor outlet
  setEditorRouterOutlet(gongsimstatusID: number) {
    this.router.navigate([{
      outlets: {
        editor: ["gongsimstatus-detail", gongsimstatusID]
      }
    }]);
  }

  // set presentation outlet
  setPresentationRouterOutlet(gongsimstatusID: number) {
    this.router.navigate([{
      outlets: {
        presentation: ["gongsimstatus-presentation", gongsimstatusID]
      }
    }]);
  }

  /** Whether the number of selected elements matches the total number of rows. */
  isAllSelected() {
    const numSelected = this.selection.selected.length;
    const numRows = this.gongsimstatuss.length;
    return numSelected === numRows;
  }

  /** Selects all rows if they are not all selected; otherwise clear selection. */
  masterToggle() {
    this.isAllSelected() ?
      this.selection.clear() :
      this.gongsimstatuss.forEach(row => this.selection.select(row));
  }

  save() {

    let toUpdate = new Set<GongsimStatusDB>()

    // reset all initial selection of gongsimstatus that belong to aclass through Anarrayofb
    this.initialSelection.forEach(
      gongsimstatus => {
        gongsimstatus[this.dialogData.ReversePointer].Int64 = 0
        gongsimstatus[this.dialogData.ReversePointer].Valid = true
        toUpdate.add(gongsimstatus)
      }
    )

    // from selection, set gongsimstatus that belong to aclass through Anarrayofb
    this.selection.selected.forEach(
      gongsimstatus => {
        console.log("selection ID " + gongsimstatus.ID)
        let ID = +this.dialogData.ID
        gongsimstatus[this.dialogData.ReversePointer].Int64 = ID
        gongsimstatus[this.dialogData.ReversePointer].Valid = true
        toUpdate.add(gongsimstatus)
      }
    )

    // update all gongsimstatus (only update selection & initial selection)
    toUpdate.forEach(
      gongsimstatus => {
        this.gongsimstatusService.updateGongsimStatus(gongsimstatus)
          .subscribe(gongsimstatus => {
            this.gongsimstatusService.GongsimStatusServiceChanged.next("update")
            console.log("gongsimstatus saved")
          });
      }
    )
    this.dialogRef.close('Pizza!');
  }
}
