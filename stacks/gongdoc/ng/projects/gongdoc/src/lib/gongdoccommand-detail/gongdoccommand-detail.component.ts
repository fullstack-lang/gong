// generated from NgDetailTemplateTS
import { Component, OnInit } from '@angular/core';
import { FormControl } from '@angular/forms';

import { GongdocCommandDB } from '../gongdoccommand-db'
import { GongdocCommandService } from '../gongdoccommand.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'
import { MapOfComponents } from '../map-components'

// insertion point for imports
import { GongdocCommandTypeSelect, GongdocCommandTypeList } from '../GongdocCommandType'
import { GongdocNodeTypeSelect, GongdocNodeTypeList } from '../GongdocNodeType'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

import { MatDialog, MAT_DIALOG_DATA, MatDialogRef, MatDialogConfig } from '@angular/material/dialog';

import { NullInt64 } from '../front-repo.service'

@Component({
	selector: 'app-gongdoccommand-detail',
	templateUrl: './gongdoccommand-detail.component.html',
	styleUrls: ['./gongdoccommand-detail.component.css'],
})
export class GongdocCommandDetailComponent implements OnInit {

	// insertion point for declarations
	GongdocCommandTypeList: GongdocCommandTypeSelect[]
	GongdocNodeTypeList: GongdocNodeTypeSelect[]

	// the GongdocCommandDB of interest
	gongdoccommand: GongdocCommandDB;

	// front repo
	frontRepo: FrontRepo

	constructor(
		private gongdoccommandService: GongdocCommandService,
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
		this.getGongdocCommand()

		// observable for changes in structs
		this.gongdoccommandService.GongdocCommandServiceChanged.subscribe(
			message => {
				if (message == "post" || message == "update" || message == "delete") {
					this.getGongdocCommand()
				}
			}
		)

		// insertion point for initialisation of enums list
		this.GongdocCommandTypeList = GongdocCommandTypeList
		this.GongdocNodeTypeList = GongdocNodeTypeList
	}

	getGongdocCommand(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		const association = this.route.snapshot.paramMap.get('association');

		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo
				console.log("front repo GongdocCommandPull returned")

				if (id != 0 && association == undefined) {
					this.gongdoccommand = frontRepo.GongdocCommands.get(id)
				} else {
					this.gongdoccommand = new (GongdocCommandDB)
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

			this.gongdoccommandService.updateGongdocCommand(this.gongdoccommand)
				.subscribe(gongdoccommand => {
					this.gongdoccommandService.GongdocCommandServiceChanged.next("update")

					console.log("gongdoccommand saved")
				});
		} else {
			switch (association) {
				// insertion point for saving value of ONE_MANY association reverse pointer
			}
			this.gongdoccommandService.postGongdocCommand(this.gongdoccommand).subscribe(gongdoccommand => {

				this.gongdoccommandService.GongdocCommandServiceChanged.next("post")

				this.gongdoccommand = {} // reset fields
				console.log("gongdoccommand added")
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
			ID: this.gongdoccommand.ID,
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
