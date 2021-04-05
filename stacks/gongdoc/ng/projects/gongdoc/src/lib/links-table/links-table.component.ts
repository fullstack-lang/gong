// generated from NgTableTemplateTS
import { Component, OnInit, OnChanges, Input, Output, EventEmitter, Inject, Optional } from '@angular/core';
import { MatTableDataSource } from '@angular/material/table';
import { MatButton } from '@angular/material/button'

import { MatDialogRef, MAT_DIALOG_DATA, MatDialog } from '@angular/material/dialog'
import { DialogData } from '../front-repo.service'
import { SelectionModel } from '@angular/cdk/collections';

const allowMultiSelect = true;

import { Router, RouterState } from '@angular/router';
import { LinkDB } from '../link-db'
import { LinkService } from '../link.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

// generated table component
@Component({
  selector: 'app-links-table',
  templateUrl: './links-table.component.html',
  styleUrls: ['./links-table.component.css'],
})
export class LinksTableComponent implements OnInit {

  // used if the component is called as a selection component of Link instances
  selection: SelectionModel<LinkDB>;
  initialSelection = new Array<LinkDB>();

  // the data source for the table
  links: LinkDB[];

  // front repo, that will be referenced by this.links
  frontRepo: FrontRepo

  // displayedColumns is referenced by the MatTable component for specify what columns
  // have to be displayed and in what order
  displayedColumns: string[];

  constructor(
    private linkService: LinkService,
    private frontRepoService: FrontRepoService,

    // not null if the component is called as a selection component of link instances
    public dialogRef: MatDialogRef<LinksTableComponent>,
    @Optional() @Inject(MAT_DIALOG_DATA) public dialogData: DialogData,

    private router: Router,
  ) {
    // https://stackoverflow.com/questions/54627478/angular-7-routing-to-same-component-but-different-param-not-working
    // this is for routerLink on same component when only queryParameter changes
    this.router.routeReuseStrategy.shouldReuseRoute = function () {
      return false;
    };

    // observable for changes in structs
    this.linkService.LinkServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.getLinks()
        }
      }
    )
    if (dialogData == undefined) {
  	  this.displayedColumns = ['ID', 'Edit', 'Delete', // insertion point for columns to display
        "Name",
        "Fieldname",
        "Structname",
        "Fieldtypename",
        "Multiplicity",
        "Middlevertice",
        "Links",
      ]
    } else {
  	  this.displayedColumns = ['select', 'ID', // insertion point for columns to display
        "Name",
        "Fieldname",
        "Structname",
        "Fieldtypename",
        "Multiplicity",
        "Middlevertice",
        "Links",
      ]
      this.selection = new SelectionModel<LinkDB>(allowMultiSelect, this.initialSelection);
    }

  }

  ngOnInit(): void {
    this.getLinks()
  }

  getLinks(): void {
    this.frontRepoService.pull().subscribe(
      frontRepo => {
        this.frontRepo = frontRepo
        console.log("front repo pull returned")

        this.links = this.frontRepo.Links_array;
        // in case the component is called as a selection component
        if (this.dialogData != undefined) {
          this.links.forEach(
            link => {
              let ID = this.dialogData.ID
              let revPointer = link[this.dialogData.ReversePointer]
              if (revPointer.Int64 == ID) {
                this.initialSelection.push(link)
              }
            }
          )
          this.selection = new SelectionModel<LinkDB>(allowMultiSelect, this.initialSelection);
        }
      }
    )
  }

  // newLink initiate a new link
  // create a new Link objet
  newLink() {
  }

  deleteLink(linkID: number, link: LinkDB) {
    // list of links is truncated of link before the delete
    this.links = this.links.filter(h => h !== link);

    this.linkService.deleteLink(linkID).subscribe(
      link => {
        this.linkService.LinkServiceChanged.next("delete")

        console.log("link deleted")
      }
    );
  }

  editLink(linkID: number, link: LinkDB) {

  }

  // display link in router
  displayLinkInRouter(linkID: number) {
    this.router.navigate( ["link-display", linkID])
  }

  // set editor outlet
  setEditorRouterOutlet(linkID: number) {
    this.router.navigate([{
      outlets: {
        editor: ["link-detail", linkID]
      }
    }]);
  }

  // set presentation outlet
  setPresentationRouterOutlet(linkID: number) {
    this.router.navigate([{
      outlets: {
        presentation: ["link-presentation", linkID]
      }
    }]);
  }

  /** Whether the number of selected elements matches the total number of rows. */
  isAllSelected() {
    const numSelected = this.selection.selected.length;
    const numRows = this.links.length;
    return numSelected === numRows;
  }

  /** Selects all rows if they are not all selected; otherwise clear selection. */
  masterToggle() {
    this.isAllSelected() ?
      this.selection.clear() :
      this.links.forEach(row => this.selection.select(row));
  }

  save() {

    let toUpdate = new Set<LinkDB>()

    // reset all initial selection of link that belong to aclass through Anarrayofb
    this.initialSelection.forEach(
      link => {
        link[this.dialogData.ReversePointer].Int64 = 0
        link[this.dialogData.ReversePointer].Valid = true
        toUpdate.add(link)
      }
    )

    // from selection, set link that belong to aclass through Anarrayofb
    this.selection.selected.forEach(
      link => {
        console.log("selection ID " + link.ID)
        let ID = +this.dialogData.ID
        link[this.dialogData.ReversePointer].Int64 = ID
        link[this.dialogData.ReversePointer].Valid = true
        toUpdate.add(link)
      }
    )

    // update all link (only update selection & initial selection)
    toUpdate.forEach(
      link => {
        this.linkService.updateLink(link)
          .subscribe(link => {
            this.linkService.LinkServiceChanged.next("update")
            console.log("link saved")
          });
      }
    )
    this.dialogRef.close('Pizza!');
  }
}
