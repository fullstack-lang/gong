// generated from NgDetailTemplateTS
import { Component, OnInit } from '@angular/core';
import { FormControl } from '@angular/forms';

import { ModelPkgDB } from '../modelpkg-db'
import { ModelPkgService } from '../modelpkg.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'
import { MapOfComponents } from '../map-components'

// insertion point for imports

import { Router, RouterState, ActivatedRoute } from '@angular/router';

import { MatDialog, MAT_DIALOG_DATA, MatDialogRef, MatDialogConfig } from '@angular/material/dialog';

import { NullInt64 } from '../front-repo.service'

@Component({
	selector: 'app-modelpkg-detail',
	templateUrl: './modelpkg-detail.component.html',
	styleUrls: ['./modelpkg-detail.component.css'],
})
export class ModelPkgDetailComponent implements OnInit {

	// insertion point for declarations

	// the ModelPkgDB of interest
	modelpkg: ModelPkgDB;

	// front repo
	frontRepo: FrontRepo

	constructor(
		private modelpkgService: ModelPkgService,
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
		this.getModelPkg()

		// observable for changes in structs
		this.modelpkgService.ModelPkgServiceChanged.subscribe(
			message => {
				if (message == "post" || message == "update" || message == "delete") {
					this.getModelPkg()
				}
			}
		)

		// insertion point for initialisation of enums list
	}

	getModelPkg(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		const association = this.route.snapshot.paramMap.get('association');

		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo
				console.log("front repo ModelPkgPull returned")

				if (id != 0 && association == undefined) {
					this.modelpkg = frontRepo.ModelPkgs.get(id)
				} else {
					this.modelpkg = new (ModelPkgDB)
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

			this.modelpkgService.updateModelPkg(this.modelpkg)
				.subscribe(modelpkg => {
					this.modelpkgService.ModelPkgServiceChanged.next("update")

					console.log("modelpkg saved")
				});
		} else {
			switch (association) {
				// insertion point for saving value of ONE_MANY association reverse pointer
			}
			this.modelpkgService.postModelPkg(this.modelpkg).subscribe(modelpkg => {

				this.modelpkgService.ModelPkgServiceChanged.next("post")

				this.modelpkg = {} // reset fields
				console.log("modelpkg added")
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
			ID: this.modelpkg.ID,
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
