// generated from NgDetailTemplateTS
import { Component, OnInit } from '@angular/core';
import { FormControl } from '@angular/forms';

import { GongsimCommandDB } from '../gongsimcommand-db'
import { GongsimCommandService } from '../gongsimcommand.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'
import { MapOfComponents } from '../map-components'

// insertion point for imports
import { GongsimCommandTypeSelect, GongsimCommandTypeList } from '../GongsimCommandType'
import { SpeedCommandTypeSelect, SpeedCommandTypeList } from '../SpeedCommandType'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

import { MatDialog, MAT_DIALOG_DATA, MatDialogRef, MatDialogConfig } from '@angular/material/dialog';

import { NullInt64 } from '../front-repo.service'

@Component({
	selector: 'app-gongsimcommand-detail',
	templateUrl: './gongsimcommand-detail.component.html',
	styleUrls: ['./gongsimcommand-detail.component.css'],
})
export class GongsimCommandDetailComponent implements OnInit {

	// insertion point for declarations
	GongsimCommandTypeList: GongsimCommandTypeSelect[]
	SpeedCommandTypeList: SpeedCommandTypeSelect[]

	// the GongsimCommandDB of interest
	gongsimcommand: GongsimCommandDB;

	// front repo
	frontRepo: FrontRepo

	constructor(
		private gongsimcommandService: GongsimCommandService,
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
		this.getGongsimCommand()

		// observable for changes in structs
		this.gongsimcommandService.GongsimCommandServiceChanged.subscribe(
			message => {
				if (message == "post" || message == "update" || message == "delete") {
					this.getGongsimCommand()
				}
			}
		)

		// insertion point for initialisation of enums list
		this.GongsimCommandTypeList = GongsimCommandTypeList
		this.SpeedCommandTypeList = SpeedCommandTypeList
	}

	getGongsimCommand(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		const association = this.route.snapshot.paramMap.get('association');

		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo
				console.log("front repo GongsimCommandPull returned")

				if (id != 0 && association == undefined) {
					this.gongsimcommand = frontRepo.GongsimCommands.get(id)
				} else {
					this.gongsimcommand = new (GongsimCommandDB)
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

			this.gongsimcommandService.updateGongsimCommand(this.gongsimcommand)
				.subscribe(gongsimcommand => {
					this.gongsimcommandService.GongsimCommandServiceChanged.next("update")

					console.log("gongsimcommand saved")
				});
		} else {
			switch (association) {
				// insertion point for saving value of ONE_MANY association reverse pointer
			}
			this.gongsimcommandService.postGongsimCommand(this.gongsimcommand).subscribe(gongsimcommand => {

				this.gongsimcommandService.GongsimCommandServiceChanged.next("post")

				this.gongsimcommand = {} // reset fields
				console.log("gongsimcommand added")
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
			ID: this.gongsimcommand.ID,
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
