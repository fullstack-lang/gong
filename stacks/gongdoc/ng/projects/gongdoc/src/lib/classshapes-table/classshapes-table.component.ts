// generated from NgTableTemplateTS
import { Component, OnInit, OnChanges, Input, Output, EventEmitter, Inject, Optional } from '@angular/core';
import { MatTableDataSource } from '@angular/material/table';
import { MatButton } from '@angular/material/button'

import { MatDialogRef, MAT_DIALOG_DATA, MatDialog } from '@angular/material/dialog'
import { DialogData } from '../front-repo.service'
import { SelectionModel } from '@angular/cdk/collections';

const allowMultiSelect = true;

import { Router, RouterState } from '@angular/router';
import { ClassshapeDB } from '../classshape-db'
import { ClassshapeService } from '../classshape.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

// generated table component
@Component({
  selector: 'app-classshapes-table',
  templateUrl: './classshapes-table.component.html',
  styleUrls: ['./classshapes-table.component.css'],
})
export class ClassshapesTableComponent implements OnInit {

  // used if the component is called as a selection component of Classshape instances
  selection: SelectionModel<ClassshapeDB>;
  initialSelection = new Array<ClassshapeDB>();

  // the data source for the table
  classshapes: ClassshapeDB[];

  // front repo, that will be referenced by this.classshapes
  frontRepo: FrontRepo

  // displayedColumns is referenced by the MatTable component for specify what columns
  // have to be displayed and in what order
  displayedColumns: string[];

  constructor(
    private classshapeService: ClassshapeService,
    private frontRepoService: FrontRepoService,

    // not null if the component is called as a selection component of classshape instances
    public dialogRef: MatDialogRef<ClassshapesTableComponent>,
    @Optional() @Inject(MAT_DIALOG_DATA) public dialogData: DialogData,

    private router: Router,
  ) {
    // https://stackoverflow.com/questions/54627478/angular-7-routing-to-same-component-but-different-param-not-working
    // this is for routerLink on same component when only queryParameter changes
    this.router.routeReuseStrategy.shouldReuseRoute = function () {
      return false;
    };

    // observable for changes in structs
    this.classshapeService.ClassshapeServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.getClassshapes()
        }
      }
    )
    if (dialogData == undefined) {
  	  this.displayedColumns = ['ID', 'Edit', 'Delete', // insertion point for columns to display
        "Name",
        "Position",
        "Structname",
        "Width",
        "Heigth",
        "ClassshapeTargetType",
        "Classshapes",
      ]
    } else {
  	  this.displayedColumns = ['select', 'ID', // insertion point for columns to display
        "Name",
        "Position",
        "Structname",
        "Width",
        "Heigth",
        "ClassshapeTargetType",
        "Classshapes",
      ]
      this.selection = new SelectionModel<ClassshapeDB>(allowMultiSelect, this.initialSelection);
    }

  }

  ngOnInit(): void {
    this.getClassshapes()
  }

  getClassshapes(): void {
    this.frontRepoService.pull().subscribe(
      frontRepo => {
        this.frontRepo = frontRepo
        console.log("front repo pull returned")

        this.classshapes = this.frontRepo.Classshapes_array;
        // in case the component is called as a selection component
        if (this.dialogData != undefined) {
          this.classshapes.forEach(
            classshape => {
              let ID = this.dialogData.ID
              let revPointer = classshape[this.dialogData.ReversePointer]
              if (revPointer.Int64 == ID) {
                this.initialSelection.push(classshape)
              }
            }
          )
          this.selection = new SelectionModel<ClassshapeDB>(allowMultiSelect, this.initialSelection);
        }
      }
    )
  }

  // newClassshape initiate a new classshape
  // create a new Classshape objet
  newClassshape() {
  }

  deleteClassshape(classshapeID: number, classshape: ClassshapeDB) {
    // list of classshapes is truncated of classshape before the delete
    this.classshapes = this.classshapes.filter(h => h !== classshape);

    this.classshapeService.deleteClassshape(classshapeID).subscribe(
      classshape => {
        this.classshapeService.ClassshapeServiceChanged.next("delete")

        console.log("classshape deleted")
      }
    );
  }

  editClassshape(classshapeID: number, classshape: ClassshapeDB) {

  }

  // display classshape in router
  displayClassshapeInRouter(classshapeID: number) {
    this.router.navigate( ["classshape-display", classshapeID])
  }

  // set editor outlet
  setEditorRouterOutlet(classshapeID: number) {
    this.router.navigate([{
      outlets: {
        editor: ["classshape-detail", classshapeID]
      }
    }]);
  }

  // set presentation outlet
  setPresentationRouterOutlet(classshapeID: number) {
    this.router.navigate([{
      outlets: {
        presentation: ["classshape-presentation", classshapeID]
      }
    }]);
  }

  /** Whether the number of selected elements matches the total number of rows. */
  isAllSelected() {
    const numSelected = this.selection.selected.length;
    const numRows = this.classshapes.length;
    return numSelected === numRows;
  }

  /** Selects all rows if they are not all selected; otherwise clear selection. */
  masterToggle() {
    this.isAllSelected() ?
      this.selection.clear() :
      this.classshapes.forEach(row => this.selection.select(row));
  }

  save() {

    let toUpdate = new Set<ClassshapeDB>()

    // reset all initial selection of classshape that belong to aclass through Anarrayofb
    this.initialSelection.forEach(
      classshape => {
        classshape[this.dialogData.ReversePointer].Int64 = 0
        classshape[this.dialogData.ReversePointer].Valid = true
        toUpdate.add(classshape)
      }
    )

    // from selection, set classshape that belong to aclass through Anarrayofb
    this.selection.selected.forEach(
      classshape => {
        console.log("selection ID " + classshape.ID)
        let ID = +this.dialogData.ID
        classshape[this.dialogData.ReversePointer].Int64 = ID
        classshape[this.dialogData.ReversePointer].Valid = true
        toUpdate.add(classshape)
      }
    )

    // update all classshape (only update selection & initial selection)
    toUpdate.forEach(
      classshape => {
        this.classshapeService.updateClassshape(classshape)
          .subscribe(classshape => {
            this.classshapeService.ClassshapeServiceChanged.next("update")
            console.log("classshape saved")
          });
      }
    )
    this.dialogRef.close('Pizza!');
  }
}
