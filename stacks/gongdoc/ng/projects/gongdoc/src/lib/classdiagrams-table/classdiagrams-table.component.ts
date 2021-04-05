// generated from NgTableTemplateTS
import { Component, OnInit, OnChanges, Input, Output, EventEmitter, Inject, Optional } from '@angular/core';
import { MatTableDataSource } from '@angular/material/table';
import { MatButton } from '@angular/material/button'

import { MatDialogRef, MAT_DIALOG_DATA, MatDialog } from '@angular/material/dialog'
import { DialogData } from '../front-repo.service'
import { SelectionModel } from '@angular/cdk/collections';

const allowMultiSelect = true;

import { Router, RouterState } from '@angular/router';
import { ClassdiagramDB } from '../classdiagram-db'
import { ClassdiagramService } from '../classdiagram.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

// generated table component
@Component({
  selector: 'app-classdiagrams-table',
  templateUrl: './classdiagrams-table.component.html',
  styleUrls: ['./classdiagrams-table.component.css'],
})
export class ClassdiagramsTableComponent implements OnInit {

  // used if the component is called as a selection component of Classdiagram instances
  selection: SelectionModel<ClassdiagramDB>;
  initialSelection = new Array<ClassdiagramDB>();

  // the data source for the table
  classdiagrams: ClassdiagramDB[];

  // front repo, that will be referenced by this.classdiagrams
  frontRepo: FrontRepo

  // displayedColumns is referenced by the MatTable component for specify what columns
  // have to be displayed and in what order
  displayedColumns: string[];

  constructor(
    private classdiagramService: ClassdiagramService,
    private frontRepoService: FrontRepoService,

    // not null if the component is called as a selection component of classdiagram instances
    public dialogRef: MatDialogRef<ClassdiagramsTableComponent>,
    @Optional() @Inject(MAT_DIALOG_DATA) public dialogData: DialogData,

    private router: Router,
  ) {
    // https://stackoverflow.com/questions/54627478/angular-7-routing-to-same-component-but-different-param-not-working
    // this is for routerLink on same component when only queryParameter changes
    this.router.routeReuseStrategy.shouldReuseRoute = function () {
      return false;
    };

    // observable for changes in structs
    this.classdiagramService.ClassdiagramServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.getClassdiagrams()
        }
      }
    )
    if (dialogData == undefined) {
  	  this.displayedColumns = ['ID', 'Edit', 'Delete', // insertion point for columns to display
        "Name",
        "Classdiagrams",
      ]
    } else {
  	  this.displayedColumns = ['select', 'ID', // insertion point for columns to display
        "Name",
        "Classdiagrams",
      ]
      this.selection = new SelectionModel<ClassdiagramDB>(allowMultiSelect, this.initialSelection);
    }

  }

  ngOnInit(): void {
    this.getClassdiagrams()
  }

  getClassdiagrams(): void {
    this.frontRepoService.pull().subscribe(
      frontRepo => {
        this.frontRepo = frontRepo
        console.log("front repo pull returned")

        this.classdiagrams = this.frontRepo.Classdiagrams_array;
        // in case the component is called as a selection component
        if (this.dialogData != undefined) {
          this.classdiagrams.forEach(
            classdiagram => {
              let ID = this.dialogData.ID
              let revPointer = classdiagram[this.dialogData.ReversePointer]
              if (revPointer.Int64 == ID) {
                this.initialSelection.push(classdiagram)
              }
            }
          )
          this.selection = new SelectionModel<ClassdiagramDB>(allowMultiSelect, this.initialSelection);
        }
      }
    )
  }

  // newClassdiagram initiate a new classdiagram
  // create a new Classdiagram objet
  newClassdiagram() {
  }

  deleteClassdiagram(classdiagramID: number, classdiagram: ClassdiagramDB) {
    // list of classdiagrams is truncated of classdiagram before the delete
    this.classdiagrams = this.classdiagrams.filter(h => h !== classdiagram);

    this.classdiagramService.deleteClassdiagram(classdiagramID).subscribe(
      classdiagram => {
        this.classdiagramService.ClassdiagramServiceChanged.next("delete")

        console.log("classdiagram deleted")
      }
    );
  }

  editClassdiagram(classdiagramID: number, classdiagram: ClassdiagramDB) {

  }

  // display classdiagram in router
  displayClassdiagramInRouter(classdiagramID: number) {
    this.router.navigate( ["classdiagram-display", classdiagramID])
  }

  // set editor outlet
  setEditorRouterOutlet(classdiagramID: number) {
    this.router.navigate([{
      outlets: {
        editor: ["classdiagram-detail", classdiagramID]
      }
    }]);
  }

  // set presentation outlet
  setPresentationRouterOutlet(classdiagramID: number) {
    this.router.navigate([{
      outlets: {
        presentation: ["classdiagram-presentation", classdiagramID]
      }
    }]);
  }

  /** Whether the number of selected elements matches the total number of rows. */
  isAllSelected() {
    const numSelected = this.selection.selected.length;
    const numRows = this.classdiagrams.length;
    return numSelected === numRows;
  }

  /** Selects all rows if they are not all selected; otherwise clear selection. */
  masterToggle() {
    this.isAllSelected() ?
      this.selection.clear() :
      this.classdiagrams.forEach(row => this.selection.select(row));
  }

  save() {

    let toUpdate = new Set<ClassdiagramDB>()

    // reset all initial selection of classdiagram that belong to aclass through Anarrayofb
    this.initialSelection.forEach(
      classdiagram => {
        classdiagram[this.dialogData.ReversePointer].Int64 = 0
        classdiagram[this.dialogData.ReversePointer].Valid = true
        toUpdate.add(classdiagram)
      }
    )

    // from selection, set classdiagram that belong to aclass through Anarrayofb
    this.selection.selected.forEach(
      classdiagram => {
        console.log("selection ID " + classdiagram.ID)
        let ID = +this.dialogData.ID
        classdiagram[this.dialogData.ReversePointer].Int64 = ID
        classdiagram[this.dialogData.ReversePointer].Valid = true
        toUpdate.add(classdiagram)
      }
    )

    // update all classdiagram (only update selection & initial selection)
    toUpdate.forEach(
      classdiagram => {
        this.classdiagramService.updateClassdiagram(classdiagram)
          .subscribe(classdiagram => {
            this.classdiagramService.ClassdiagramServiceChanged.next("update")
            console.log("classdiagram saved")
          });
      }
    )
    this.dialogRef.close('Pizza!');
  }
}
