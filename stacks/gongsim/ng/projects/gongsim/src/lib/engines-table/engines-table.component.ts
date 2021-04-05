// generated from NgTableTemplateTS
import { Component, OnInit, OnChanges, Input, Output, EventEmitter, Inject, Optional } from '@angular/core';
import { MatTableDataSource } from '@angular/material/table';
import { MatButton } from '@angular/material/button'

import { MatDialogRef, MAT_DIALOG_DATA, MatDialog } from '@angular/material/dialog'
import { DialogData } from '../front-repo.service'
import { SelectionModel } from '@angular/cdk/collections';

const allowMultiSelect = true;

import { Router, RouterState } from '@angular/router';
import { EngineDB } from '../engine-db'
import { EngineService } from '../engine.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

// generated table component
@Component({
  selector: 'app-engines-table',
  templateUrl: './engines-table.component.html',
  styleUrls: ['./engines-table.component.css'],
})
export class EnginesTableComponent implements OnInit {

  // used if the component is called as a selection component of Engine instances
  selection: SelectionModel<EngineDB>;
  initialSelection = new Array<EngineDB>();

  // the data source for the table
  engines: EngineDB[];

  // front repo, that will be referenced by this.engines
  frontRepo: FrontRepo

  // displayedColumns is referenced by the MatTable component for specify what columns
  // have to be displayed and in what order
  displayedColumns: string[];

  constructor(
    private engineService: EngineService,
    private frontRepoService: FrontRepoService,

    // not null if the component is called as a selection component of engine instances
    public dialogRef: MatDialogRef<EnginesTableComponent>,
    @Optional() @Inject(MAT_DIALOG_DATA) public dialogData: DialogData,

    private router: Router,
  ) {
    // https://stackoverflow.com/questions/54627478/angular-7-routing-to-same-component-but-different-param-not-working
    // this is for routerLink on same component when only queryParameter changes
    this.router.routeReuseStrategy.shouldReuseRoute = function () {
      return false;
    };

    // observable for changes in structs
    this.engineService.EngineServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.getEngines()
        }
      }
    )
    if (dialogData == undefined) {
  	  this.displayedColumns = ['ID', 'Edit', 'Delete', // insertion point for columns to display
        "Name",
        "EndTime",
        "CurrentTime",
        "SecondsSinceStart",
        "Fired",
        "ControlMode",
        "State",
        "Speed",
      ]
    } else {
  	  this.displayedColumns = ['select', 'ID', // insertion point for columns to display
        "Name",
        "EndTime",
        "CurrentTime",
        "SecondsSinceStart",
        "Fired",
        "ControlMode",
        "State",
        "Speed",
      ]
      this.selection = new SelectionModel<EngineDB>(allowMultiSelect, this.initialSelection);
    }

  }

  ngOnInit(): void {
    this.getEngines()
  }

  getEngines(): void {
    this.frontRepoService.pull().subscribe(
      frontRepo => {
        this.frontRepo = frontRepo
        console.log("front repo pull returned")

        this.engines = this.frontRepo.Engines_array;
        // in case the component is called as a selection component
        if (this.dialogData != undefined) {
          this.engines.forEach(
            engine => {
              let ID = this.dialogData.ID
              let revPointer = engine[this.dialogData.ReversePointer]
              if (revPointer.Int64 == ID) {
                this.initialSelection.push(engine)
              }
            }
          )
          this.selection = new SelectionModel<EngineDB>(allowMultiSelect, this.initialSelection);
        }
      }
    )
  }

  // newEngine initiate a new engine
  // create a new Engine objet
  newEngine() {
  }

  deleteEngine(engineID: number, engine: EngineDB) {
    // list of engines is truncated of engine before the delete
    this.engines = this.engines.filter(h => h !== engine);

    this.engineService.deleteEngine(engineID).subscribe(
      engine => {
        this.engineService.EngineServiceChanged.next("delete")

        console.log("engine deleted")
      }
    );
  }

  editEngine(engineID: number, engine: EngineDB) {

  }

  // display engine in router
  displayEngineInRouter(engineID: number) {
    this.router.navigate( ["engine-display", engineID])
  }

  // set editor outlet
  setEditorRouterOutlet(engineID: number) {
    this.router.navigate([{
      outlets: {
        editor: ["engine-detail", engineID]
      }
    }]);
  }

  // set presentation outlet
  setPresentationRouterOutlet(engineID: number) {
    this.router.navigate([{
      outlets: {
        presentation: ["engine-presentation", engineID]
      }
    }]);
  }

  /** Whether the number of selected elements matches the total number of rows. */
  isAllSelected() {
    const numSelected = this.selection.selected.length;
    const numRows = this.engines.length;
    return numSelected === numRows;
  }

  /** Selects all rows if they are not all selected; otherwise clear selection. */
  masterToggle() {
    this.isAllSelected() ?
      this.selection.clear() :
      this.engines.forEach(row => this.selection.select(row));
  }

  save() {

    let toUpdate = new Set<EngineDB>()

    // reset all initial selection of engine that belong to aclass through Anarrayofb
    this.initialSelection.forEach(
      engine => {
        engine[this.dialogData.ReversePointer].Int64 = 0
        engine[this.dialogData.ReversePointer].Valid = true
        toUpdate.add(engine)
      }
    )

    // from selection, set engine that belong to aclass through Anarrayofb
    this.selection.selected.forEach(
      engine => {
        console.log("selection ID " + engine.ID)
        let ID = +this.dialogData.ID
        engine[this.dialogData.ReversePointer].Int64 = ID
        engine[this.dialogData.ReversePointer].Valid = true
        toUpdate.add(engine)
      }
    )

    // update all engine (only update selection & initial selection)
    toUpdate.forEach(
      engine => {
        this.engineService.updateEngine(engine)
          .subscribe(engine => {
            this.engineService.EngineServiceChanged.next("update")
            console.log("engine saved")
          });
      }
    )
    this.dialogRef.close('Pizza!');
  }
}
