// generated from NgDetailTemplateTS
import { Component, OnInit } from '@angular/core';
import { FormControl } from '@angular/forms';

import { GongdocStatusDB } from '../gongdocstatus-db'
import { GongdocStatusService } from '../gongdocstatus.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'
import { MapOfComponents } from '../map-components'

// insertion point for imports
import { GongdocCommandTypeSelect, GongdocCommandTypeList } from '../GongdocCommandType'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

import { MatDialog, MAT_DIALOG_DATA, MatDialogRef, MatDialogConfig } from '@angular/material/dialog';

import { NullInt64 } from '../front-repo.service'

@Component({
	selector: 'app-gongdocstatus-detail',
	templateUrl: './gongdocstatus-detail.component.html',
	styleUrls: ['./gongdocstatus-detail.component.css'],
})
export class GongdocStatusDetailComponent implements OnInit {

	// insertion point for declarations
	GongdocCommandTypeList: GongdocCommandTypeSelect[]

	// the GongdocStatusDB of interest
	gongdocstatus: GongdocStatusDB;

	// front repo
	frontRepo: FrontRepo

	constructor(
		private gongdocstatusService: GongdocStatusService,
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
		this.getGongdocStatus()

		// observable for changes in structs
		this.gongdocstatusService.GongdocStatusServiceChanged.subscribe(
			message => {
				if (message == "post" || message == "update" || message == "delete") {
					this.getGongdocStatus()
				}
			}
		)

		// insertion point for initialisation of enums list
		this.GongdocCommandTypeList = GongdocCommandTypeList
	}

	getGongdocStatus(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		const association = this.route.snapshot.paramMap.get('association');

		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo
				console.log("front repo GongdocStatusPull returned")

				if (id != 0 && association == undefined) {
					this.gongdocstatus = frontRepo.GongdocStatuss.get(id)
				} else {
					this.gongdocstatus = new (GongdocStatusDB)
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

			this.gongdocstatusService.updateGongdocStatus(this.gongdocstatus)
				.subscribe(gongdocstatus => {
					this.gongdocstatusService.GongdocStatusServiceChanged.next("update")

					console.log("gongdocstatus saved")
				});
		} else {
			switch (association) {
				// insertion point for saving value of ONE_MANY association reverse pointer
			}
			this.gongdocstatusService.postGongdocStatus(this.gongdocstatus).subscribe(gongdocstatus => {

				this.gongdocstatusService.GongdocStatusServiceChanged.next("post")

				this.gongdocstatus = {} // reset fields
				console.log("gongdocstatus added")
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
			ID: this.gongdocstatus.ID,
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
