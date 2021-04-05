// generated from NgTableTemplateTS
import { Component, OnInit, OnChanges, Input, Output, EventEmitter, Inject, Optional } from '@angular/core';
import { MatTableDataSource } from '@angular/material/table';
import { MatButton } from '@angular/material/button'

import { MatDialogRef, MAT_DIALOG_DATA, MatDialog } from '@angular/material/dialog'
import { DialogData } from '../front-repo.service'
import { SelectionModel } from '@angular/cdk/collections';

const allowMultiSelect = true;

import { Router, RouterState } from '@angular/router';
import { AreaDB } from '../area-db'
import { AreaService } from '../area.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

// generated table component
@Component({
  selector: 'app-areas-table',
  templateUrl: './areas-table.component.html',
  styleUrls: ['./areas-table.component.css'],
})
export class AreasTableComponent implements OnInit {

  // used if the component is called as a selection component of Area instances
  selection: SelectionModel<AreaDB>;
  initialSelection = new Array<AreaDB>();

  // the data source for the table
  areas: AreaDB[];

  // front repo, that will be referenced by this.areas
  frontRepo: FrontRepo

  // displayedColumns is referenced by the MatTable component for specify what columns
  // have to be displayed and in what order
  displayedColumns: string[];

  constructor(
    private areaService: AreaService,
    private frontRepoService: FrontRepoService,

    // not null if the component is called as a selection component of area instances
    public dialogRef: MatDialogRef<AreasTableComponent>,
    @Optional() @Inject(MAT_DIALOG_DATA) public dialogData: DialogData,

    private router: Router,
  ) {
    // https://stackoverflow.com/questions/54627478/angular-7-routing-to-same-component-but-different-param-not-working
    // this is for routerLink on same component when only queryParameter changes
    this.router.routeReuseStrategy.shouldReuseRoute = function () {
      return false;
    };

    // observable for changes in structs
    this.areaService.AreaServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.getAreas()
        }
      }
    )
    if (dialogData == undefined) {
  	  this.displayedColumns = ['ID', 'Edit', 'Delete', // insertion point for columns to display
        "Name",
        "Subarea",
      ]
    } else {
  	  this.displayedColumns = ['select', 'ID', // insertion point for columns to display
        "Name",
        "Subarea",
      ]
      this.selection = new SelectionModel<AreaDB>(allowMultiSelect, this.initialSelection);
    }

  }

  ngOnInit(): void {
    this.getAreas()
  }

  getAreas(): void {
    this.frontRepoService.pull().subscribe(
      frontRepo => {
        this.frontRepo = frontRepo
        console.log("front repo pull returned")

        this.areas = this.frontRepo.Areas_array;
        // in case the component is called as a selection component
        if (this.dialogData != undefined) {
          this.areas.forEach(
            area => {
              let ID = this.dialogData.ID
              let revPointer = area[this.dialogData.ReversePointer]
              if (revPointer.Int64 == ID) {
                this.initialSelection.push(area)
              }
            }
          )
          this.selection = new SelectionModel<AreaDB>(allowMultiSelect, this.initialSelection);
        }
      }
    )
  }

  // newArea initiate a new area
  // create a new Area objet
  newArea() {
  }

  deleteArea(areaID: number, area: AreaDB) {
    // list of areas is truncated of area before the delete
    this.areas = this.areas.filter(h => h !== area);

    this.areaService.deleteArea(areaID).subscribe(
      area => {
        this.areaService.AreaServiceChanged.next("delete")

        console.log("area deleted")
      }
    );
  }

  editArea(areaID: number, area: AreaDB) {

  }

  // display area in router
  displayAreaInRouter(areaID: number) {
    this.router.navigate( ["area-display", areaID])
  }

  // set editor outlet
  setEditorRouterOutlet(areaID: number) {
    this.router.navigate([{
      outlets: {
        editor: ["area-detail", areaID]
      }
    }]);
  }

  // set presentation outlet
  setPresentationRouterOutlet(areaID: number) {
    this.router.navigate([{
      outlets: {
        presentation: ["area-presentation", areaID]
      }
    }]);
  }

  /** Whether the number of selected elements matches the total number of rows. */
  isAllSelected() {
    const numSelected = this.selection.selected.length;
    const numRows = this.areas.length;
    return numSelected === numRows;
  }

  /** Selects all rows if they are not all selected; otherwise clear selection. */
  masterToggle() {
    this.isAllSelected() ?
      this.selection.clear() :
      this.areas.forEach(row => this.selection.select(row));
  }

  save() {

    let toUpdate = new Set<AreaDB>()

    // reset all initial selection of area that belong to aclass through Anarrayofb
    this.initialSelection.forEach(
      area => {
        area[this.dialogData.ReversePointer].Int64 = 0
        area[this.dialogData.ReversePointer].Valid = true
        toUpdate.add(area)
      }
    )

    // from selection, set area that belong to aclass through Anarrayofb
    this.selection.selected.forEach(
      area => {
        console.log("selection ID " + area.ID)
        let ID = +this.dialogData.ID
        area[this.dialogData.ReversePointer].Int64 = ID
        area[this.dialogData.ReversePointer].Valid = true
        toUpdate.add(area)
      }
    )

    // update all area (only update selection & initial selection)
    toUpdate.forEach(
      area => {
        this.areaService.updateArea(area)
          .subscribe(area => {
            this.areaService.AreaServiceChanged.next("update")
            console.log("area saved")
          });
      }
    )
    this.dialogRef.close('Pizza!');
  }
}
