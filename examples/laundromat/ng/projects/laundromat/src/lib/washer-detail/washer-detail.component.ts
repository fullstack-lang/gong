// generated from NgDetailTemplateTS
import { Component, OnInit } from '@angular/core';
import { FormControl } from '@angular/forms';

import { WasherDB } from '../washer-db'
import { WasherService } from '../washer.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'
import { MapOfComponents } from '../map-components'

// insertion point for imports
import { WasherStateEnumSelect, WasherStateEnumList } from '../WasherStateEnum'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

import { MatDialog, MAT_DIALOG_DATA, MatDialogRef, MatDialogConfig } from '@angular/material/dialog';

import { NullInt64 } from '../front-repo.service'

@Component({
	selector: 'app-washer-detail',
	templateUrl: './washer-detail.component.html',
	styleUrls: ['./washer-detail.component.css'],
})
export class WasherDetailComponent implements OnInit {

	// insertion point for declarations
	WasherStateEnumList: WasherStateEnumSelect[]

	// the WasherDB of interest
	washer: WasherDB;

	// front repo
	frontRepo: FrontRepo

	constructor(
		private washerService: WasherService,
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
		this.getWasher()

		// observable for changes in structs
		this.washerService.WasherServiceChanged.subscribe(
			message => {
				if (message == "post" || message == "update" || message == "delete") {
					this.getWasher()
				}
			}
		)

		// insertion point for initialisation of enums list
		this.WasherStateEnumList = WasherStateEnumList
	}

	getWasher(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		const association = this.route.snapshot.paramMap.get('association');

		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo
				console.log("front repo WasherPull returned")

				if (id != 0 && association == undefined) {
					this.washer = frontRepo.Washers.get(id)
				} else {
					this.washer = new (WasherDB)
				}

				// insertion point for recovery of form controls value for bool fields
			}
		)


	}

	save(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		const association = this.route.snapshot.paramMap.get('association');

		// insertion point for saving value of form controls of boolean fields
		if (this.washer.MachineID == undefined) {
			this.washer.MachineID = new NullInt64
		}
		if (this.washer.Machine != undefined) {
			this.washer.MachineID.Int64 = this.washer.Machine.ID
			this.washer.MachineID.Valid = true
			this.washer.MachineName = this.washer.Machine.Name
		} else {
			this.washer.MachineID.Int64 = 0
			this.washer.MachineID.Valid = true
			this.washer.MachineName = ""
		}

		if (id != 0 && association == undefined) {
			// insertion point for saving value of reverse pointers

			this.washerService.updateWasher(this.washer)
				.subscribe(washer => {
					this.washerService.WasherServiceChanged.next("update")

					console.log("washer saved")
				});
		} else {
			switch (association) {
				// insertion point for saving value of ONE_MANY association reverse pointer
			}
			this.washerService.postWasher(this.washer).subscribe(washer => {

				this.washerService.WasherServiceChanged.next("post")

				this.washer = {} // reset fields
				console.log("washer added")
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
			ID: this.washer.ID,
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
