// generated from NgTableTemplateTS
import { Component, OnInit, OnChanges, Input, Output, EventEmitter, Inject, Optional } from '@angular/core';
import { MatTableDataSource } from '@angular/material/table';
import { MatButton } from '@angular/material/button'

import { MatDialogRef, MAT_DIALOG_DATA, MatDialog } from '@angular/material/dialog'
import { DialogData } from '../front-repo.service'
import { SelectionModel } from '@angular/cdk/collections';

const allowMultiSelect = true;

import { Router, RouterState } from '@angular/router';
import { ModelPkgDB } from '../modelpkg-db'
import { ModelPkgService } from '../modelpkg.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

// generated table component
@Component({
  selector: 'app-modelpkgs-table',
  templateUrl: './modelpkgs-table.component.html',
  styleUrls: ['./modelpkgs-table.component.css'],
})
export class ModelPkgsTableComponent implements OnInit {

  // used if the component is called as a selection component of ModelPkg instances
  selection: SelectionModel<ModelPkgDB>;
  initialSelection = new Array<ModelPkgDB>();

  // the data source for the table
  modelpkgs: ModelPkgDB[];

  // front repo, that will be referenced by this.modelpkgs
  frontRepo: FrontRepo

  // displayedColumns is referenced by the MatTable component for specify what columns
  // have to be displayed and in what order
  displayedColumns: string[];

  constructor(
    private modelpkgService: ModelPkgService,
    private frontRepoService: FrontRepoService,

    // not null if the component is called as a selection component of modelpkg instances
    public dialogRef: MatDialogRef<ModelPkgsTableComponent>,
    @Optional() @Inject(MAT_DIALOG_DATA) public dialogData: DialogData,

    private router: Router,
  ) {
    // https://stackoverflow.com/questions/54627478/angular-7-routing-to-same-component-but-different-param-not-working
    // this is for routerLink on same component when only queryParameter changes
    this.router.routeReuseStrategy.shouldReuseRoute = function () {
      return false;
    };

    // observable for changes in structs
    this.modelpkgService.ModelPkgServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.getModelPkgs()
        }
      }
    )
    if (dialogData == undefined) {
  	  this.displayedColumns = ['ID', 'Edit', 'Delete', // insertion point for columns to display
        "Name",
        "PkgPath",
      ]
    } else {
  	  this.displayedColumns = ['select', 'ID', // insertion point for columns to display
        "Name",
        "PkgPath",
      ]
      this.selection = new SelectionModel<ModelPkgDB>(allowMultiSelect, this.initialSelection);
    }

  }

  ngOnInit(): void {
    this.getModelPkgs()
  }

  getModelPkgs(): void {
    this.frontRepoService.pull().subscribe(
      frontRepo => {
        this.frontRepo = frontRepo
        console.log("front repo pull returned")

        this.modelpkgs = this.frontRepo.ModelPkgs_array;
        // in case the component is called as a selection component
        if (this.dialogData != undefined) {
          this.modelpkgs.forEach(
            modelpkg => {
              let ID = this.dialogData.ID
              let revPointer = modelpkg[this.dialogData.ReversePointer]
              if (revPointer.Int64 == ID) {
                this.initialSelection.push(modelpkg)
              }
            }
          )
          this.selection = new SelectionModel<ModelPkgDB>(allowMultiSelect, this.initialSelection);
        }
      }
    )
  }

  // newModelPkg initiate a new modelpkg
  // create a new ModelPkg objet
  newModelPkg() {
  }

  deleteModelPkg(modelpkgID: number, modelpkg: ModelPkgDB) {
    // list of modelpkgs is truncated of modelpkg before the delete
    this.modelpkgs = this.modelpkgs.filter(h => h !== modelpkg);

    this.modelpkgService.deleteModelPkg(modelpkgID).subscribe(
      modelpkg => {
        this.modelpkgService.ModelPkgServiceChanged.next("delete")

        console.log("modelpkg deleted")
      }
    );
  }

  editModelPkg(modelpkgID: number, modelpkg: ModelPkgDB) {

  }

  // display modelpkg in router
  displayModelPkgInRouter(modelpkgID: number) {
    this.router.navigate( ["modelpkg-display", modelpkgID])
  }

  // set editor outlet
  setEditorRouterOutlet(modelpkgID: number) {
    this.router.navigate([{
      outlets: {
        editor: ["modelpkg-detail", modelpkgID]
      }
    }]);
  }

  // set presentation outlet
  setPresentationRouterOutlet(modelpkgID: number) {
    this.router.navigate([{
      outlets: {
        presentation: ["modelpkg-presentation", modelpkgID]
      }
    }]);
  }

  /** Whether the number of selected elements matches the total number of rows. */
  isAllSelected() {
    const numSelected = this.selection.selected.length;
    const numRows = this.modelpkgs.length;
    return numSelected === numRows;
  }

  /** Selects all rows if they are not all selected; otherwise clear selection. */
  masterToggle() {
    this.isAllSelected() ?
      this.selection.clear() :
      this.modelpkgs.forEach(row => this.selection.select(row));
  }

  save() {

    let toUpdate = new Set<ModelPkgDB>()

    // reset all initial selection of modelpkg that belong to aclass through Anarrayofb
    this.initialSelection.forEach(
      modelpkg => {
        modelpkg[this.dialogData.ReversePointer].Int64 = 0
        modelpkg[this.dialogData.ReversePointer].Valid = true
        toUpdate.add(modelpkg)
      }
    )

    // from selection, set modelpkg that belong to aclass through Anarrayofb
    this.selection.selected.forEach(
      modelpkg => {
        console.log("selection ID " + modelpkg.ID)
        let ID = +this.dialogData.ID
        modelpkg[this.dialogData.ReversePointer].Int64 = ID
        modelpkg[this.dialogData.ReversePointer].Valid = true
        toUpdate.add(modelpkg)
      }
    )

    // update all modelpkg (only update selection & initial selection)
    toUpdate.forEach(
      modelpkg => {
        this.modelpkgService.updateModelPkg(modelpkg)
          .subscribe(modelpkg => {
            this.modelpkgService.ModelPkgServiceChanged.next("update")
            console.log("modelpkg saved")
          });
      }
    )
    this.dialogRef.close('Pizza!');
  }
}
