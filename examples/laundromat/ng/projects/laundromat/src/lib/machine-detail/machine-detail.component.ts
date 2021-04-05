// generated from NgDetailTemplateTS
import { Component, OnInit } from '@angular/core';
import { FormControl } from '@angular/forms';

import { MachineDB } from '../machine-db'
import { MachineService } from '../machine.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'
import { MapOfComponents } from '../map-components'

// insertion point for imports
import { MachineStateEnumSelect, MachineStateEnumList } from '../MachineStateEnum'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

import { MatDialog, MAT_DIALOG_DATA, MatDialogRef, MatDialogConfig } from '@angular/material/dialog';

import { NullInt64 } from '../front-repo.service'

@Component({
	selector: 'app-machine-detail',
	templateUrl: './machine-detail.component.html',
	styleUrls: ['./machine-detail.component.css'],
})
export class MachineDetailComponent implements OnInit {

	// insertion point for declarations
	CleanedlaundryFormControl = new FormControl(false);
	MachineStateEnumList: MachineStateEnumSelect[]

	// the MachineDB of interest
	machine: MachineDB;

	// front repo
	frontRepo: FrontRepo

	constructor(
		private machineService: MachineService,
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
		this.getMachine()

		// observable for changes in structs
		this.machineService.MachineServiceChanged.subscribe(
			message => {
				if (message == "post" || message == "update" || message == "delete") {
					this.getMachine()
				}
			}
		)

		// insertion point for initialisation of enums list
		this.MachineStateEnumList = MachineStateEnumList
	}

	getMachine(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		const association = this.route.snapshot.paramMap.get('association');

		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo
				console.log("front repo MachinePull returned")

				if (id != 0 && association == undefined) {
					this.machine = frontRepo.Machines.get(id)
				} else {
					this.machine = new (MachineDB)
				}

				// insertion point for recovery of form controls value for bool fields
				this.CleanedlaundryFormControl.setValue(this.machine.Cleanedlaundry)
			}
		)


	}

	save(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		const association = this.route.snapshot.paramMap.get('association');

		// insertion point for saving value of form controls of boolean fields
		this.machine.Cleanedlaundry = this.CleanedlaundryFormControl.value

		if (id != 0 && association == undefined) {
			// insertion point for saving value of reverse pointers

			this.machineService.updateMachine(this.machine)
				.subscribe(machine => {
					this.machineService.MachineServiceChanged.next("update")

					console.log("machine saved")
				});
		} else {
			switch (association) {
				// insertion point for saving value of ONE_MANY association reverse pointer
			}
			this.machineService.postMachine(this.machine).subscribe(machine => {

				this.machineService.MachineServiceChanged.next("post")

				this.machine = {} // reset fields
				console.log("machine added")
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
			ID: this.machine.ID,
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
