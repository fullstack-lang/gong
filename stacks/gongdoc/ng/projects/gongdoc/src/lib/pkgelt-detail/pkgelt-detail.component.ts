// generated from NgDetailTemplateTS
import { Component, OnInit } from '@angular/core';
import { FormControl } from '@angular/forms';

import { PkgeltDB } from '../pkgelt-db'
import { PkgeltService } from '../pkgelt.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'
import { MapOfComponents } from '../map-components'

// insertion point for imports

import { Router, RouterState, ActivatedRoute } from '@angular/router';

import { MatDialog, MAT_DIALOG_DATA, MatDialogRef, MatDialogConfig } from '@angular/material/dialog';

import { NullInt64 } from '../front-repo.service'

@Component({
	selector: 'app-pkgelt-detail',
	templateUrl: './pkgelt-detail.component.html',
	styleUrls: ['./pkgelt-detail.component.css'],
})
export class PkgeltDetailComponent implements OnInit {

	// insertion point for declarations

	// the PkgeltDB of interest
	pkgelt: PkgeltDB;

	// front repo
	frontRepo: FrontRepo

	constructor(
		private pkgeltService: PkgeltService,
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
		this.getPkgelt()

		// observable for changes in structs
		this.pkgeltService.PkgeltServiceChanged.subscribe(
			message => {
				if (message == "post" || message == "update" || message == "delete") {
					this.getPkgelt()
				}
			}
		)

		// insertion point for initialisation of enums list
	}

	getPkgelt(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		const association = this.route.snapshot.paramMap.get('association');

		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo
				console.log("front repo PkgeltPull returned")

				if (id != 0 && association == undefined) {
					this.pkgelt = frontRepo.Pkgelts.get(id)
				} else {
					this.pkgelt = new (PkgeltDB)
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

			this.pkgeltService.updatePkgelt(this.pkgelt)
				.subscribe(pkgelt => {
					this.pkgeltService.PkgeltServiceChanged.next("update")

					console.log("pkgelt saved")
				});
		} else {
			switch (association) {
				// insertion point for saving value of ONE_MANY association reverse pointer
			}
			this.pkgeltService.postPkgelt(this.pkgelt).subscribe(pkgelt => {

				this.pkgeltService.PkgeltServiceChanged.next("post")

				this.pkgelt = {} // reset fields
				console.log("pkgelt added")
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
			ID: this.pkgelt.ID,
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
