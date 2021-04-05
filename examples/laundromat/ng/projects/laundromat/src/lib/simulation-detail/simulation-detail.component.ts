// generated from NgDetailTemplateTS
import { Component, OnInit } from '@angular/core';
import { FormControl } from '@angular/forms';

import { SimulationDB } from '../simulation-db'
import { SimulationService } from '../simulation.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'
import { MapOfComponents } from '../map-components'

// insertion point for imports

import { Router, RouterState, ActivatedRoute } from '@angular/router';

import { MatDialog, MAT_DIALOG_DATA, MatDialogRef, MatDialogConfig } from '@angular/material/dialog';

import { NullInt64 } from '../front-repo.service'

@Component({
	selector: 'app-simulation-detail',
	templateUrl: './simulation-detail.component.html',
	styleUrls: ['./simulation-detail.component.css'],
})
export class SimulationDetailComponent implements OnInit {

	// insertion point for declarations

	// the SimulationDB of interest
	simulation: SimulationDB;

	// front repo
	frontRepo: FrontRepo

	constructor(
		private simulationService: SimulationService,
		private frontRepoService: FrontRepoService,
		public dialog: MatDialog,
		private route: ActivatedRoute,
		private router: Router,
	) {
		// https://stackoverflow.com/questions/54627478/angular-7-routing-to-same-component-but-different-param-not-working
		// this is for routerLink on same component when only queryParameter changes
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getSimulation()

		// observable for changes in structs
		this.simulationService.SimulationServiceChanged.subscribe(
			message => {
				if (message == "post" || message == "update" || message == "delete") {
					this.getSimulation()
				}
			}
		)

		// insertion point for initialisation of enums list
	}

	getSimulation(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		const association = this.route.snapshot.paramMap.get('association');

		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo
				console.log("front repo SimulationPull returned")

				if (id != 0 && association == undefined) {
					this.simulation = frontRepo.Simulations.get(id)
				} else {
					this.simulation = new (SimulationDB)
				}

				// insertion point for recovery of form controls value for bool fields
			}
		)


	}

	save(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		const association = this.route.snapshot.paramMap.get('association');

		// insertion point for saving value of form controls of boolean fields
		if (this.simulation.MachineID == undefined) {
			this.simulation.MachineID = new NullInt64
		}
		if (this.simulation.Machine != undefined) {
			this.simulation.MachineID.Int64 = this.simulation.Machine.ID
			this.simulation.MachineID.Valid = true
			this.simulation.MachineName = this.simulation.Machine.Name
		} else {
			this.simulation.MachineID.Int64 = 0
			this.simulation.MachineID.Valid = true
			this.simulation.MachineName = ""
		}
		if (this.simulation.WasherID == undefined) {
			this.simulation.WasherID = new NullInt64
		}
		if (this.simulation.Washer != undefined) {
			this.simulation.WasherID.Int64 = this.simulation.Washer.ID
			this.simulation.WasherID.Valid = true
			this.simulation.WasherName = this.simulation.Washer.Name
		} else {
			this.simulation.WasherID.Int64 = 0
			this.simulation.WasherID.Valid = true
			this.simulation.WasherName = ""
		}

		if (id != 0 && association == undefined) {
			// insertion point for saving value of reverse pointers

			this.simulationService.updateSimulation(this.simulation)
				.subscribe(simulation => {
					this.simulationService.SimulationServiceChanged.next("update")

					console.log("simulation saved")
				});
		} else {
			switch (association) {
				// insertion point for saving value of ONE_MANY association reverse pointer
			}
			this.simulationService.postSimulation(this.simulation).subscribe(simulation => {

				this.simulationService.SimulationServiceChanged.next("post")

				this.simulation = {} // reset fields
				console.log("simulation added")
			});
		}
	}

	// openReverseSelection is a generic function that calls dialog for the edition of 
	// ONE-MANY association
	// It uses the MapOfComponent provided by the front repo
	openReverseSelection(AssociatedStruct: string, reverseField: string) {

		const dialogConfig = new MatDialogConfig();

		// dialogConfig.disableClose = true;
		dialogConfig.autoFocus = true;
		dialogConfig.data = {
			ID: this.simulation.ID,
			ReversePointer: reverseField,
		};
		const dialogRef: MatDialogRef<string, any> = this.dialog.open(
			MapOfComponents.get(AssociatedStruct).get(
				AssociatedStruct + 'sTableComponent'
			),
			dialogConfig
		);

		dialogRef.afterClosed().subscribe(result => {
			console.log('The dialog was closed');
		});
	}
}
