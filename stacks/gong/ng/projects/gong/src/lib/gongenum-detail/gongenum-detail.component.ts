// generated from NgDetailTemplateTS
import { Component, OnInit } from '@angular/core';
import { FormControl } from '@angular/forms';

import { GongEnumDB } from '../gongenum-db'
import { GongEnumService } from '../gongenum.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'
import { MapOfComponents } from '../map-components'

// insertion point for imports

import { Router, RouterState, ActivatedRoute } from '@angular/router';

import { MatDialog, MAT_DIALOG_DATA, MatDialogRef, MatDialogConfig } from '@angular/material/dialog';

import { NullInt64 } from '../front-repo.service'

@Component({
	selector: 'app-gongenum-detail',
	templateUrl: './gongenum-detail.component.html',
	styleUrls: ['./gongenum-detail.component.css'],
})
export class GongEnumDetailComponent implements OnInit {

	// insertion point for declarations

	// the GongEnumDB of interest
	gongenum: GongEnumDB;

	// front repo
	frontRepo: FrontRepo

	constructor(
		private gongenumService: GongEnumService,
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
		this.getGongEnum()

		// observable for changes in structs
		this.gongenumService.GongEnumServiceChanged.subscribe(
			message => {
				if (message == "post" || message == "update" || message == "delete") {
					this.getGongEnum()
				}
			}
		)

		// insertion point for initialisation of enums list
	}

	getGongEnum(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		const association = this.route.snapshot.paramMap.get('association');

		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo
				console.log("front repo GongEnumPull returned")

				if (id != 0 && association == undefined) {
					this.gongenum = frontRepo.GongEnums.get(id)
				} else {
					this.gongenum = new (GongEnumDB)
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

			this.gongenumService.updateGongEnum(this.gongenum)
				.subscribe(gongenum => {
					this.gongenumService.GongEnumServiceChanged.next("update")

					console.log("gongenum saved")
				});
		} else {
			switch (association) {
				// insertion point for saving value of ONE_MANY association reverse pointer
			}
			this.gongenumService.postGongEnum(this.gongenum).subscribe(gongenum => {

				this.gongenumService.GongEnumServiceChanged.next("post")

				this.gongenum = {} // reset fields
				console.log("gongenum added")
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
			ID: this.gongenum.ID,
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
