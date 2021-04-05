// generated from NgTableTemplateTS
import { Component, OnInit, OnChanges, Input, Output, EventEmitter, Inject, Optional } from '@angular/core';
import { MatTableDataSource } from '@angular/material/table';
import { MatButton } from '@angular/material/button'

import { MatDialogRef, MAT_DIALOG_DATA, MatDialog } from '@angular/material/dialog'
import { DialogData } from '../front-repo.service'
import { SelectionModel } from '@angular/cdk/collections';

const allowMultiSelect = true;

import { Router, RouterState } from '@angular/router';
import { SimulationDB } from '../simulation-db'
import { SimulationService } from '../simulation.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

// generated table component
@Component({
  selector: 'app-simulations-table',
  templateUrl: './simulations-table.component.html',
  styleUrls: ['./simulations-table.component.css'],
})
export class SimulationsTableComponent implements OnInit {

  // used if the component is called as a selection component of Simulation instances
  selection: SelectionModel<SimulationDB>;
  initialSelection = new Array<SimulationDB>();

  // the data source for the table
  simulations: SimulationDB[];

  // front repo, that will be referenced by this.simulations
  frontRepo: FrontRepo

  // displayedColumns is referenced by the MatTable component for specify what columns
  // have to be displayed and in what order
  displayedColumns: string[];

  constructor(
    private simulationService: SimulationService,
    private frontRepoService: FrontRepoService,

    // not null if the component is called as a selection component of simulation instances
    public dialogRef: MatDialogRef<SimulationsTableComponent>,
    @Optional() @Inject(MAT_DIALOG_DATA) public dialogData: DialogData,

    private router: Router,
  ) {
    // https://stackoverflow.com/questions/54627478/angular-7-routing-to-same-component-but-different-param-not-working
    // this is for routerLink on same component when only queryParameter changes
    this.router.routeReuseStrategy.shouldReuseRoute = function () {
      return false;
    };

    // observable for changes in structs
    this.simulationService.SimulationServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.getSimulations()
        }
      }
    )
    if (dialogData == undefined) {
  	  this.displayedColumns = ['ID', 'Edit', 'Delete', // insertion point for columns to display
        "Name",
        "Machine",
        "Washer",
        "LastCommitNb",
      ]
    } else {
  	  this.displayedColumns = ['select', 'ID', // insertion point for columns to display
        "Name",
        "Machine",
        "Washer",
        "LastCommitNb",
      ]
      this.selection = new SelectionModel<SimulationDB>(allowMultiSelect, this.initialSelection);
    }

  }

  ngOnInit(): void {
    this.getSimulations()
  }

  getSimulations(): void {
    this.frontRepoService.pull().subscribe(
      frontRepo => {
        this.frontRepo = frontRepo
        console.log("front repo pull returned")

        this.simulations = this.frontRepo.Simulations_array;
        // in case the component is called as a selection component
        if (this.dialogData != undefined) {
          this.simulations.forEach(
            simulation => {
              let ID = this.dialogData.ID
              let revPointer = simulation[this.dialogData.ReversePointer]
              if (revPointer.Int64 == ID) {
                this.initialSelection.push(simulation)
              }
            }
          )
          this.selection = new SelectionModel<SimulationDB>(allowMultiSelect, this.initialSelection);
        }
      }
    )
  }

  // newSimulation initiate a new simulation
  // create a new Simulation objet
  newSimulation() {
  }

  deleteSimulation(simulationID: number, simulation: SimulationDB) {
    // list of simulations is truncated of simulation before the delete
    this.simulations = this.simulations.filter(h => h !== simulation);

    this.simulationService.deleteSimulation(simulationID).subscribe(
      simulation => {
        this.simulationService.SimulationServiceChanged.next("delete")

        console.log("simulation deleted")
      }
    );
  }

  editSimulation(simulationID: number, simulation: SimulationDB) {

  }

  // display simulation in router
  displaySimulationInRouter(simulationID: number) {
    this.router.navigate( ["simulation-display", simulationID])
  }

  // set editor outlet
  setEditorRouterOutlet(simulationID: number) {
    this.router.navigate([{
      outlets: {
        editor: ["simulation-detail", simulationID]
      }
    }]);
  }

  // set presentation outlet
  setPresentationRouterOutlet(simulationID: number) {
    this.router.navigate([{
      outlets: {
        presentation: ["simulation-presentation", simulationID]
      }
    }]);
  }

  /** Whether the number of selected elements matches the total number of rows. */
  isAllSelected() {
    const numSelected = this.selection.selected.length;
    const numRows = this.simulations.length;
    return numSelected === numRows;
  }

  /** Selects all rows if they are not all selected; otherwise clear selection. */
  masterToggle() {
    this.isAllSelected() ?
      this.selection.clear() :
      this.simulations.forEach(row => this.selection.select(row));
  }

  save() {

    let toUpdate = new Set<SimulationDB>()

    // reset all initial selection of simulation that belong to aclass through Anarrayofb
    this.initialSelection.forEach(
      simulation => {
        simulation[this.dialogData.ReversePointer].Int64 = 0
        simulation[this.dialogData.ReversePointer].Valid = true
        toUpdate.add(simulation)
      }
    )

    // from selection, set simulation that belong to aclass through Anarrayofb
    this.selection.selected.forEach(
      simulation => {
        console.log("selection ID " + simulation.ID)
        let ID = +this.dialogData.ID
        simulation[this.dialogData.ReversePointer].Int64 = ID
        simulation[this.dialogData.ReversePointer].Valid = true
        toUpdate.add(simulation)
      }
    )

    // update all simulation (only update selection & initial selection)
    toUpdate.forEach(
      simulation => {
        this.simulationService.updateSimulation(simulation)
          .subscribe(simulation => {
            this.simulationService.SimulationServiceChanged.next("update")
            console.log("simulation saved")
          });
      }
    )
    this.dialogRef.close('Pizza!');
  }
}
