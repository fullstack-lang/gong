// generated from NgDetailTemplateTS
import { Component, OnInit } from '@angular/core';
import { FormControl } from '@angular/forms';

import { GongsimStatusDB } from '../gongsimstatus-db'
import { GongsimStatusService } from '../gongsimstatus.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'
import { MapOfComponents } from '../map-components'

// insertion point for imports
import { GongsimCommandTypeSelect, GongsimCommandTypeList } from '../GongsimCommandType'
import { SpeedCommandTypeSelect, SpeedCommandTypeList } from '../SpeedCommandType'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

import { MatDialog, MAT_DIALOG_DATA, MatDialogRef, MatDialogConfig } from '@angular/material/dialog';

import { NullInt64 } from '../front-repo.service'

@Component({
	selector: 'app-gongsimstatus-detail',
	templateUrl: './gongsimstatus-detail.component.html',
	styleUrls: ['./gongsimstatus-detail.component.css'],
})
export class GongsimStatusDetailComponent implements OnInit {

	// insertion point for declarations
	GongsimCommandTypeList: GongsimCommandTypeSelect[]
	SpeedCommandTypeList: SpeedCommandTypeSelect[]

	// the GongsimStatusDB of interest
	gongsimstatus: GongsimStatusDB;

	// front repo
	frontRepo: FrontRepo

	constructor(
		private gongsimstatusService: GongsimStatusService,
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
		this.getGongsimStatus()

		// observable for changes in structs
		this.gongsimstatusService.GongsimStatusServiceChanged.subscribe(
			message => {
				if (message == "post" || message == "update" || message == "delete") {
					this.getGongsimStatus()
				}
			}
		)

		// insertion point for initialisation of enums list
		this.GongsimCommandTypeList = GongsimCommandTypeList
		this.SpeedCommandTypeList = SpeedCommandTypeList
	}

	getGongsimStatus(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		const association = this.route.snapshot.paramMap.get('association');

		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo
				console.log("front repo GongsimStatusPull returned")

				if (id != 0 && association == undefined) {
					this.gongsimstatus = frontRepo.GongsimStatuss.get(id)
				} else {
					this.gongsimstatus = new (GongsimStatusDB)
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

			this.gongsimstatusService.updateGongsimStatus(this.gongsimstatus)
				.subscribe(gongsimstatus => {
					this.gongsimstatusService.GongsimStatusServiceChanged.next("update")

					console.log("gongsimstatus saved")
				});
		} else {
			switch (association) {
				// insertion point for saving value of ONE_MANY association reverse pointer
			}
			this.gongsimstatusService.postGongsimStatus(this.gongsimstatus).subscribe(gongsimstatus => {

				this.gongsimstatusService.GongsimStatusServiceChanged.next("post")

				this.gongsimstatus = {} // reset fields
				console.log("gongsimstatus added")
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
			ID: this.gongsimstatus.ID,
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
