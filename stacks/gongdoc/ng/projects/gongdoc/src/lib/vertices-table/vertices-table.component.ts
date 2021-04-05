// generated from NgTableTemplateTS
import { Component, OnInit, OnChanges, Input, Output, EventEmitter, Inject, Optional } from '@angular/core';
import { MatTableDataSource } from '@angular/material/table';
import { MatButton } from '@angular/material/button'

import { MatDialogRef, MAT_DIALOG_DATA, MatDialog } from '@angular/material/dialog'
import { DialogData } from '../front-repo.service'
import { SelectionModel } from '@angular/cdk/collections';

const allowMultiSelect = true;

import { Router, RouterState } from '@angular/router';
import { VerticeDB } from '../vertice-db'
import { VerticeService } from '../vertice.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

// generated table component
@Component({
  selector: 'app-vertices-table',
  templateUrl: './vertices-table.component.html',
  styleUrls: ['./vertices-table.component.css'],
})
export class VerticesTableComponent implements OnInit {

  // used if the component is called as a selection component of Vertice instances
  selection: SelectionModel<VerticeDB>;
  initialSelection = new Array<VerticeDB>();

  // the data source for the table
  vertices: VerticeDB[];

  // front repo, that will be referenced by this.vertices
  frontRepo: FrontRepo

  // displayedColumns is referenced by the MatTable component for specify what columns
  // have to be displayed and in what order
  displayedColumns: string[];

  constructor(
    private verticeService: VerticeService,
    private frontRepoService: FrontRepoService,

    // not null if the component is called as a selection component of vertice instances
    public dialogRef: MatDialogRef<VerticesTableComponent>,
    @Optional() @Inject(MAT_DIALOG_DATA) public dialogData: DialogData,

    private router: Router,
  ) {
    // https://stackoverflow.com/questions/54627478/angular-7-routing-to-same-component-but-different-param-not-working
    // this is for routerLink on same component when only queryParameter changes
    this.router.routeReuseStrategy.shouldReuseRoute = function () {
      return false;
    };

    // observable for changes in structs
    this.verticeService.VerticeServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.getVertices()
        }
      }
    )
    if (dialogData == undefined) {
  	  this.displayedColumns = ['ID', 'Edit', 'Delete', // insertion point for columns to display
        "X",
        "Y",
        "Name",
      ]
    } else {
  	  this.displayedColumns = ['select', 'ID', // insertion point for columns to display
        "X",
        "Y",
        "Name",
      ]
      this.selection = new SelectionModel<VerticeDB>(allowMultiSelect, this.initialSelection);
    }

  }

  ngOnInit(): void {
    this.getVertices()
  }

  getVertices(): void {
    this.frontRepoService.pull().subscribe(
      frontRepo => {
        this.frontRepo = frontRepo
        console.log("front repo pull returned")

        this.vertices = this.frontRepo.Vertices_array;
        // in case the component is called as a selection component
        if (this.dialogData != undefined) {
          this.vertices.forEach(
            vertice => {
              let ID = this.dialogData.ID
              let revPointer = vertice[this.dialogData.ReversePointer]
              if (revPointer.Int64 == ID) {
                this.initialSelection.push(vertice)
              }
            }
          )
          this.selection = new SelectionModel<VerticeDB>(allowMultiSelect, this.initialSelection);
        }
      }
    )
  }

  // newVertice initiate a new vertice
  // create a new Vertice objet
  newVertice() {
  }

  deleteVertice(verticeID: number, vertice: VerticeDB) {
    // list of vertices is truncated of vertice before the delete
    this.vertices = this.vertices.filter(h => h !== vertice);

    this.verticeService.deleteVertice(verticeID).subscribe(
      vertice => {
        this.verticeService.VerticeServiceChanged.next("delete")

        console.log("vertice deleted")
      }
    );
  }

  editVertice(verticeID: number, vertice: VerticeDB) {

  }

  // display vertice in router
  displayVerticeInRouter(verticeID: number) {
    this.router.navigate( ["vertice-display", verticeID])
  }

  // set editor outlet
  setEditorRouterOutlet(verticeID: number) {
    this.router.navigate([{
      outlets: {
        editor: ["vertice-detail", verticeID]
      }
    }]);
  }

  // set presentation outlet
  setPresentationRouterOutlet(verticeID: number) {
    this.router.navigate([{
      outlets: {
        presentation: ["vertice-presentation", verticeID]
      }
    }]);
  }

  /** Whether the number of selected elements matches the total number of rows. */
  isAllSelected() {
    const numSelected = this.selection.selected.length;
    const numRows = this.vertices.length;
    return numSelected === numRows;
  }

  /** Selects all rows if they are not all selected; otherwise clear selection. */
  masterToggle() {
    this.isAllSelected() ?
      this.selection.clear() :
      this.vertices.forEach(row => this.selection.select(row));
  }

  save() {

    let toUpdate = new Set<VerticeDB>()

    // reset all initial selection of vertice that belong to aclass through Anarrayofb
    this.initialSelection.forEach(
      vertice => {
        vertice[this.dialogData.ReversePointer].Int64 = 0
        vertice[this.dialogData.ReversePointer].Valid = true
        toUpdate.add(vertice)
      }
    )

    // from selection, set vertice that belong to aclass through Anarrayofb
    this.selection.selected.forEach(
      vertice => {
        console.log("selection ID " + vertice.ID)
        let ID = +this.dialogData.ID
        vertice[this.dialogData.ReversePointer].Int64 = ID
        vertice[this.dialogData.ReversePointer].Valid = true
        toUpdate.add(vertice)
      }
    )

    // update all vertice (only update selection & initial selection)
    toUpdate.forEach(
      vertice => {
        this.verticeService.updateVertice(vertice)
          .subscribe(vertice => {
            this.verticeService.VerticeServiceChanged.next("update")
            console.log("vertice saved")
          });
      }
    )
    this.dialogRef.close('Pizza!');
  }
}
