// generated from NgTableTemplateTS
import { Component, OnInit, OnChanges, Input, Output, EventEmitter, Inject, Optional } from '@angular/core';
import { MatTableDataSource } from '@angular/material/table';
import { MatButton } from '@angular/material/button'

import { MatDialogRef, MAT_DIALOG_DATA, MatDialog } from '@angular/material/dialog'
import { DialogData } from '../front-repo.service'
import { SelectionModel } from '@angular/cdk/collections';

const allowMultiSelect = true;

import { Router, RouterState } from '@angular/router';
import { AclassDB } from '../aclass-db'
import { AclassService } from '../aclass.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

// generated table component
@Component({
  selector: 'app-aclasss-table',
  templateUrl: './aclasss-table.component.html',
  styleUrls: ['./aclasss-table.component.css'],
})
export class AclasssTableComponent implements OnInit {

  // used if the component is called as a selection component of Aclass instances
  selection: SelectionModel<AclassDB>;
  initialSelection = new Array<AclassDB>();

  // the data source for the table
  aclasss: AclassDB[];

  // front repo, that will be referenced by this.aclasss
  frontRepo: FrontRepo

  // displayedColumns is referenced by the MatTable component for specify what columns
  // have to be displayed and in what order
  displayedColumns: string[];

  constructor(
    private aclassService: AclassService,
    private frontRepoService: FrontRepoService,

    // not null if the component is called as a selection component of aclass instances
    public dialogRef: MatDialogRef<AclasssTableComponent>,
    @Optional() @Inject(MAT_DIALOG_DATA) public dialogData: DialogData,

    private router: Router,
  ) {
    // https://stackoverflow.com/questions/54627478/angular-7-routing-to-same-component-but-different-param-not-working
    // this is for routerLink on same component when only queryParameter changes
    this.router.routeReuseStrategy.shouldReuseRoute = function () {
      return false;
    };

    // observable for changes in structs
    this.aclassService.AclassServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.getAclasss()
        }
      }
    )
    if (dialogData == undefined) {
  	  this.displayedColumns = ['ID', 'Edit', 'Delete', // insertion point for columns to display
        "Name",
        "Booleanfield",
        "Aenum",
        "Aenum_2",
        "Benum",
        "CName",
        "CFloatfield",
        "Floatfield",
        "Intfield",
        "Anotherbooleanfield",
        "Associationtob",
        "Anotherassociationtob_2",
        "Anarrayofa",
      ]
    } else {
  	  this.displayedColumns = ['select', 'ID', // insertion point for columns to display
        "Name",
        "Booleanfield",
        "Aenum",
        "Aenum_2",
        "Benum",
        "CName",
        "CFloatfield",
        "Floatfield",
        "Intfield",
        "Anotherbooleanfield",
        "Associationtob",
        "Anotherassociationtob_2",
        "Anarrayofa",
      ]
      this.selection = new SelectionModel<AclassDB>(allowMultiSelect, this.initialSelection);
    }

  }

  ngOnInit(): void {
    this.getAclasss()
  }

  getAclasss(): void {
    this.frontRepoService.pull().subscribe(
      frontRepo => {
        this.frontRepo = frontRepo
        console.log("front repo pull returned")

        this.aclasss = this.frontRepo.Aclasss_array;
        // in case the component is called as a selection component
        if (this.dialogData != undefined) {
          this.aclasss.forEach(
            aclass => {
              let ID = this.dialogData.ID
              let revPointer = aclass[this.dialogData.ReversePointer]
              if (revPointer.Int64 == ID) {
                this.initialSelection.push(aclass)
              }
            }
          )
          this.selection = new SelectionModel<AclassDB>(allowMultiSelect, this.initialSelection);
        }
      }
    )
  }

  // newAclass initiate a new aclass
  // create a new Aclass objet
  newAclass() {
  }

  deleteAclass(aclassID: number, aclass: AclassDB) {
    // list of aclasss is truncated of aclass before the delete
    this.aclasss = this.aclasss.filter(h => h !== aclass);

    this.aclassService.deleteAclass(aclassID).subscribe(
      aclass => {
        this.aclassService.AclassServiceChanged.next("delete")

        console.log("aclass deleted")
      }
    );
  }

  editAclass(aclassID: number, aclass: AclassDB) {

  }

  // display aclass in router
  displayAclassInRouter(aclassID: number) {
    this.router.navigate( ["aclass-display", aclassID])
  }

  // set editor outlet
  setEditorRouterOutlet(aclassID: number) {
    this.router.navigate([{
      outlets: {
        editor: ["aclass-detail", aclassID]
      }
    }]);
  }

  // set presentation outlet
  setPresentationRouterOutlet(aclassID: number) {
    this.router.navigate([{
      outlets: {
        presentation: ["aclass-presentation", aclassID]
      }
    }]);
  }

  /** Whether the number of selected elements matches the total number of rows. */
  isAllSelected() {
    const numSelected = this.selection.selected.length;
    const numRows = this.aclasss.length;
    return numSelected === numRows;
  }

  /** Selects all rows if they are not all selected; otherwise clear selection. */
  masterToggle() {
    this.isAllSelected() ?
      this.selection.clear() :
      this.aclasss.forEach(row => this.selection.select(row));
  }

  save() {

    let toUpdate = new Set<AclassDB>()

    // reset all initial selection of aclass that belong to aclass through Anarrayofb
    this.initialSelection.forEach(
      aclass => {
        aclass[this.dialogData.ReversePointer].Int64 = 0
        aclass[this.dialogData.ReversePointer].Valid = true
        toUpdate.add(aclass)
      }
    )

    // from selection, set aclass that belong to aclass through Anarrayofb
    this.selection.selected.forEach(
      aclass => {
        console.log("selection ID " + aclass.ID)
        let ID = +this.dialogData.ID
        aclass[this.dialogData.ReversePointer].Int64 = ID
        aclass[this.dialogData.ReversePointer].Valid = true
        toUpdate.add(aclass)
      }
    )

    // update all aclass (only update selection & initial selection)
    toUpdate.forEach(
      aclass => {
        this.aclassService.updateAclass(aclass)
          .subscribe(aclass => {
            this.aclassService.AclassServiceChanged.next("update")
            console.log("aclass saved")
          });
      }
    )
    this.dialogRef.close('Pizza!');
  }
}