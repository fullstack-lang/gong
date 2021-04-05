// generated from NgTableTemplateTS
import { Component, OnInit, OnChanges, Input, Output, EventEmitter, Inject, Optional } from '@angular/core';
import { MatTableDataSource } from '@angular/material/table';
import { MatButton } from '@angular/material/button'

import { MatDialogRef, MAT_DIALOG_DATA, MatDialog } from '@angular/material/dialog'
import { DialogData } from '../front-repo.service'
import { SelectionModel } from '@angular/cdk/collections';

const allowMultiSelect = true;

import { Router, RouterState } from '@angular/router';
import { EditorDB } from '../editor-db'
import { EditorService } from '../editor.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

// generated table component
@Component({
  selector: 'app-editors-table',
  templateUrl: './editors-table.component.html',
  styleUrls: ['./editors-table.component.css'],
})
export class EditorsTableComponent implements OnInit {

  // used if the component is called as a selection component of Editor instances
  selection: SelectionModel<EditorDB>;
  initialSelection = new Array<EditorDB>();

  // the data source for the table
  editors: EditorDB[];

  // front repo, that will be referenced by this.editors
  frontRepo: FrontRepo

  // displayedColumns is referenced by the MatTable component for specify what columns
  // have to be displayed and in what order
  displayedColumns: string[];

  constructor(
    private editorService: EditorService,
    private frontRepoService: FrontRepoService,

    // not null if the component is called as a selection component of editor instances
    public dialogRef: MatDialogRef<EditorsTableComponent>,
    @Optional() @Inject(MAT_DIALOG_DATA) public dialogData: DialogData,

    private router: Router,
  ) {
    // https://stackoverflow.com/questions/54627478/angular-7-routing-to-same-component-but-different-param-not-working
    // this is for routerLink on same component when only queryParameter changes
    this.router.routeReuseStrategy.shouldReuseRoute = function () {
      return false;
    };

    // observable for changes in structs
    this.editorService.EditorServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.getEditors()
        }
      }
    )
    if (dialogData == undefined) {
  	  this.displayedColumns = ['ID', 'Edit', 'Delete', // insertion point for columns to display
        "Name",
      ]
    } else {
  	  this.displayedColumns = ['select', 'ID', // insertion point for columns to display
        "Name",
      ]
      this.selection = new SelectionModel<EditorDB>(allowMultiSelect, this.initialSelection);
    }

  }

  ngOnInit(): void {
    this.getEditors()
  }

  getEditors(): void {
    this.frontRepoService.pull().subscribe(
      frontRepo => {
        this.frontRepo = frontRepo
        console.log("front repo pull returned")

        this.editors = this.frontRepo.Editors_array;
        // in case the component is called as a selection component
        if (this.dialogData != undefined) {
          this.editors.forEach(
            editor => {
              let ID = this.dialogData.ID
              let revPointer = editor[this.dialogData.ReversePointer]
              if (revPointer.Int64 == ID) {
                this.initialSelection.push(editor)
              }
            }
          )
          this.selection = new SelectionModel<EditorDB>(allowMultiSelect, this.initialSelection);
        }
      }
    )
  }

  // newEditor initiate a new editor
  // create a new Editor objet
  newEditor() {
  }

  deleteEditor(editorID: number, editor: EditorDB) {
    // list of editors is truncated of editor before the delete
    this.editors = this.editors.filter(h => h !== editor);

    this.editorService.deleteEditor(editorID).subscribe(
      editor => {
        this.editorService.EditorServiceChanged.next("delete")

        console.log("editor deleted")
      }
    );
  }

  editEditor(editorID: number, editor: EditorDB) {

  }

  // display editor in router
  displayEditorInRouter(editorID: number) {
    this.router.navigate( ["editor-display", editorID])
  }

  // set editor outlet
  setEditorRouterOutlet(editorID: number) {
    this.router.navigate([{
      outlets: {
        editor: ["editor-detail", editorID]
      }
    }]);
  }

  // set presentation outlet
  setPresentationRouterOutlet(editorID: number) {
    this.router.navigate([{
      outlets: {
        presentation: ["editor-presentation", editorID]
      }
    }]);
  }

  /** Whether the number of selected elements matches the total number of rows. */
  isAllSelected() {
    const numSelected = this.selection.selected.length;
    const numRows = this.editors.length;
    return numSelected === numRows;
  }

  /** Selects all rows if they are not all selected; otherwise clear selection. */
  masterToggle() {
    this.isAllSelected() ?
      this.selection.clear() :
      this.editors.forEach(row => this.selection.select(row));
  }

  save() {

    let toUpdate = new Set<EditorDB>()

    // reset all initial selection of editor that belong to aclass through Anarrayofb
    this.initialSelection.forEach(
      editor => {
        editor[this.dialogData.ReversePointer].Int64 = 0
        editor[this.dialogData.ReversePointer].Valid = true
        toUpdate.add(editor)
      }
    )

    // from selection, set editor that belong to aclass through Anarrayofb
    this.selection.selected.forEach(
      editor => {
        console.log("selection ID " + editor.ID)
        let ID = +this.dialogData.ID
        editor[this.dialogData.ReversePointer].Int64 = ID
        editor[this.dialogData.ReversePointer].Valid = true
        toUpdate.add(editor)
      }
    )

    // update all editor (only update selection & initial selection)
    toUpdate.forEach(
      editor => {
        this.editorService.updateEditor(editor)
          .subscribe(editor => {
            this.editorService.EditorServiceChanged.next("update")
            console.log("editor saved")
          });
      }
    )
    this.dialogRef.close('Pizza!');
  }
}
