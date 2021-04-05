// generated from NgTableTemplateTS
import { Component, OnInit, OnChanges, Input, Output, EventEmitter, Inject, Optional } from '@angular/core';
import { MatTableDataSource } from '@angular/material/table';
import { MatButton } from '@angular/material/button'

import { MatDialogRef, MAT_DIALOG_DATA, MatDialog } from '@angular/material/dialog'
import { DialogData } from '../front-repo.service'
import { SelectionModel } from '@angular/cdk/collections';

const allowMultiSelect = true;

import { Router, RouterState } from '@angular/router';
import { EventDB } from '../event-db'
import { EventService } from '../event.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

// generated table component
@Component({
  selector: 'app-events-table',
  templateUrl: './events-table.component.html',
  styleUrls: ['./events-table.component.css'],
})
export class EventsTableComponent implements OnInit {

  // used if the component is called as a selection component of Event instances
  selection: SelectionModel<EventDB>;
  initialSelection = new Array<EventDB>();

  // the data source for the table
  events: EventDB[];

  // front repo, that will be referenced by this.events
  frontRepo: FrontRepo

  // displayedColumns is referenced by the MatTable component for specify what columns
  // have to be displayed and in what order
  displayedColumns: string[];

  constructor(
    private eventService: EventService,
    private frontRepoService: FrontRepoService,

    // not null if the component is called as a selection component of event instances
    public dialogRef: MatDialogRef<EventsTableComponent>,
    @Optional() @Inject(MAT_DIALOG_DATA) public dialogData: DialogData,

    private router: Router,
  ) {
    // https://stackoverflow.com/questions/54627478/angular-7-routing-to-same-component-but-different-param-not-working
    // this is for routerLink on same component when only queryParameter changes
    this.router.routeReuseStrategy.shouldReuseRoute = function () {
      return false;
    };

    // observable for changes in structs
    this.eventService.EventServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.getEvents()
        }
      }
    )
    if (dialogData == undefined) {
  	  this.displayedColumns = ['ID', 'Edit', 'Delete', // insertion point for columns to display
        "Name",
        "Duration",
      ]
    } else {
  	  this.displayedColumns = ['select', 'ID', // insertion point for columns to display
        "Name",
        "Duration",
      ]
      this.selection = new SelectionModel<EventDB>(allowMultiSelect, this.initialSelection);
    }

  }

  ngOnInit(): void {
    this.getEvents()
  }

  getEvents(): void {
    this.frontRepoService.pull().subscribe(
      frontRepo => {
        this.frontRepo = frontRepo
        console.log("front repo pull returned")

        this.events = this.frontRepo.Events_array;
        // in case the component is called as a selection component
        if (this.dialogData != undefined) {
          this.events.forEach(
            event => {
              let ID = this.dialogData.ID
              let revPointer = event[this.dialogData.ReversePointer]
              if (revPointer.Int64 == ID) {
                this.initialSelection.push(event)
              }
            }
          )
          this.selection = new SelectionModel<EventDB>(allowMultiSelect, this.initialSelection);
        }
      }
    )
  }

  // newEvent initiate a new event
  // create a new Event objet
  newEvent() {
  }

  deleteEvent(eventID: number, event: EventDB) {
    // list of events is truncated of event before the delete
    this.events = this.events.filter(h => h !== event);

    this.eventService.deleteEvent(eventID).subscribe(
      event => {
        this.eventService.EventServiceChanged.next("delete")

        console.log("event deleted")
      }
    );
  }

  editEvent(eventID: number, event: EventDB) {

  }

  // display event in router
  displayEventInRouter(eventID: number) {
    this.router.navigate( ["event-display", eventID])
  }

  // set editor outlet
  setEditorRouterOutlet(eventID: number) {
    this.router.navigate([{
      outlets: {
        editor: ["event-detail", eventID]
      }
    }]);
  }

  // set presentation outlet
  setPresentationRouterOutlet(eventID: number) {
    this.router.navigate([{
      outlets: {
        presentation: ["event-presentation", eventID]
      }
    }]);
  }

  /** Whether the number of selected elements matches the total number of rows. */
  isAllSelected() {
    const numSelected = this.selection.selected.length;
    const numRows = this.events.length;
    return numSelected === numRows;
  }

  /** Selects all rows if they are not all selected; otherwise clear selection. */
  masterToggle() {
    this.isAllSelected() ?
      this.selection.clear() :
      this.events.forEach(row => this.selection.select(row));
  }

  save() {

    let toUpdate = new Set<EventDB>()

    // reset all initial selection of event that belong to aclass through Anarrayofb
    this.initialSelection.forEach(
      event => {
        event[this.dialogData.ReversePointer].Int64 = 0
        event[this.dialogData.ReversePointer].Valid = true
        toUpdate.add(event)
      }
    )

    // from selection, set event that belong to aclass through Anarrayofb
    this.selection.selected.forEach(
      event => {
        console.log("selection ID " + event.ID)
        let ID = +this.dialogData.ID
        event[this.dialogData.ReversePointer].Int64 = ID
        event[this.dialogData.ReversePointer].Valid = true
        toUpdate.add(event)
      }
    )

    // update all event (only update selection & initial selection)
    toUpdate.forEach(
      event => {
        this.eventService.updateEvent(event)
          .subscribe(event => {
            this.eventService.EventServiceChanged.next("update")
            console.log("event saved")
          });
      }
    )
    this.dialogRef.close('Pizza!');
  }
}
