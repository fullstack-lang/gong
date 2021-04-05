// generated from NgDetailTemplateTS
import { Component, OnInit } from '@angular/core';
import { FormControl } from '@angular/forms';

import { EngineDB } from '../engine-db'
import { EngineService } from '../engine.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'
import { MapOfComponents } from '../map-components'

// insertion point for imports
import { ControlModeSelect, ControlModeList } from '../ControlMode'
import { EngineStateSelect, EngineStateList } from '../EngineState'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

import { MatDialog, MAT_DIALOG_DATA, MatDialogRef, MatDialogConfig } from '@angular/material/dialog';

import { NullInt64 } from '../front-repo.service'

@Component({
	selector: 'app-engine-detail',
	templateUrl: './engine-detail.component.html',
	styleUrls: ['./engine-detail.component.css'],
})
export class EngineDetailComponent implements OnInit {

	// insertion point for declarations
	ControlModeList: ControlModeSelect[]
	EngineStateList: EngineStateSelect[]

	// the EngineDB of interest
	engine: EngineDB;

	// front repo
	frontRepo: FrontRepo

	constructor(
		private engineService: EngineService,
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
		this.getEngine()

		// observable for changes in structs
		this.engineService.EngineServiceChanged.subscribe(
			message => {
				if (message == "post" || message == "update" || message == "delete") {
					this.getEngine()
				}
			}
		)

		// insertion point for initialisation of enums list
		this.ControlModeList = ControlModeList
		this.EngineStateList = EngineStateList
	}

	getEngine(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		const association = this.route.snapshot.paramMap.get('association');

		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo
				console.log("front repo EnginePull returned")

				if (id != 0 && association == undefined) {
					this.engine = frontRepo.Engines.get(id)
				} else {
					this.engine = new (EngineDB)
				}

				// insertion point for recovery of form controls value for bool fields
			}
		)


	}

	save(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		const association = this.route.snapshot.paramMap.get('association');

		// insertion point for saving value of form controls of boolean fields

		if (id != 0 && association == undefined) {
			// insertion point for saving value of reverse pointers

			this.engineService.updateEngine(this.engine)
				.subscribe(engine => {
					this.engineService.EngineServiceChanged.next("update")

					console.log("engine saved")
				});
		} else {
			switch (association) {
				// insertion point for saving value of ONE_MANY association reverse pointer
			}
			this.engineService.postEngine(this.engine).subscribe(engine => {

				this.engineService.EngineServiceChanged.next("post")

				this.engine = {} // reset fields
				console.log("engine added")
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
			ID: this.engine.ID,
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
