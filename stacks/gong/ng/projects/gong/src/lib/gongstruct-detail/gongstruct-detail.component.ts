// generated from NgDetailTemplateTS
import { Component, OnInit } from '@angular/core';
import { FormControl } from '@angular/forms';

import { GongStructDB } from '../gongstruct-db'
import { GongStructService } from '../gongstruct.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'
import { MapOfComponents } from '../map-components'

// insertion point for imports

import { Router, RouterState, ActivatedRoute } from '@angular/router';

import { MatDialog, MAT_DIALOG_DATA, MatDialogRef, MatDialogConfig } from '@angular/material/dialog';

import { NullInt64 } from '../front-repo.service'

@Component({
	selector: 'app-gongstruct-detail',
	templateUrl: './gongstruct-detail.component.html',
	styleUrls: ['./gongstruct-detail.component.css'],
})
export class GongStructDetailComponent implements OnInit {

	// insertion point for declarations

	// the GongStructDB of interest
	gongstruct: GongStructDB;

	// front repo
	frontRepo: FrontRepo

	constructor(
		private gongstructService: GongStructService,
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
		this.getGongStruct()

		// observable for changes in structs
		this.gongstructService.GongStructServiceChanged.subscribe(
			message => {
				if (message == "post" || message == "update" || message == "delete") {
					this.getGongStruct()
				}
			}
		)

		// insertion point for initialisation of enums list
	}

	getGongStruct(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		const association = this.route.snapshot.paramMap.get('association');

		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo
				console.log("front repo GongStructPull returned")

				if (id != 0 && association == undefined) {
					this.gongstruct = frontRepo.GongStructs.get(id)
				} else {
					this.gongstruct = new (GongStructDB)
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

			this.gongstructService.updateGongStruct(this.gongstruct)
				.subscribe(gongstruct => {
					this.gongstructService.GongStructServiceChanged.next("update")

					console.log("gongstruct saved")
				});
		} else {
			switch (association) {
				// insertion point for saving value of ONE_MANY association reverse pointer
			}
			this.gongstructService.postGongStruct(this.gongstruct).subscribe(gongstruct => {

				this.gongstructService.GongStructServiceChanged.next("post")

				this.gongstruct = {} // reset fields
				console.log("gongstruct added")
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
			ID: this.gongstruct.ID,
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
