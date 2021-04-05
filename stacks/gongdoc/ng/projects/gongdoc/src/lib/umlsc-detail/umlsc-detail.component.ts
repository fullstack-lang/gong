// generated from NgDetailTemplateTS
import { Component, OnInit } from '@angular/core';
import { FormControl } from '@angular/forms';

import { UmlscDB } from '../umlsc-db'
import { UmlscService } from '../umlsc.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'
import { MapOfComponents } from '../map-components'

// insertion point for imports

import { Router, RouterState, ActivatedRoute } from '@angular/router';

import { MatDialog, MAT_DIALOG_DATA, MatDialogRef, MatDialogConfig } from '@angular/material/dialog';

import { NullInt64 } from '../front-repo.service'

@Component({
	selector: 'app-umlsc-detail',
	templateUrl: './umlsc-detail.component.html',
	styleUrls: ['./umlsc-detail.component.css'],
})
export class UmlscDetailComponent implements OnInit {

	// insertion point for declarations

	// the UmlscDB of interest
	umlsc: UmlscDB;

	// front repo
	frontRepo: FrontRepo

	constructor(
		private umlscService: UmlscService,
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
		this.getUmlsc()

		// observable for changes in structs
		this.umlscService.UmlscServiceChanged.subscribe(
			message => {
				if (message == "post" || message == "update" || message == "delete") {
					this.getUmlsc()
				}
			}
		)

		// insertion point for initialisation of enums list
	}

	getUmlsc(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		const association = this.route.snapshot.paramMap.get('association');

		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo
				console.log("front repo UmlscPull returned")

				if (id != 0 && association == undefined) {
					this.umlsc = frontRepo.Umlscs.get(id)
				} else {
					this.umlsc = new (UmlscDB)
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
			if (this.umlsc.Pkgelt_Umlscs_reverse != undefined) {
				this.umlsc.Pkgelt_UmlscsDBID = new NullInt64
				this.umlsc.Pkgelt_UmlscsDBID.Int64 = this.umlsc.Pkgelt_Umlscs_reverse.ID
				this.umlsc.Pkgelt_UmlscsDBID.Valid = true
				this.umlsc.Pkgelt_Umlscs_reverse = undefined // very important, otherwise, circular JSON
			}

			this.umlscService.updateUmlsc(this.umlsc)
				.subscribe(umlsc => {
					this.umlscService.UmlscServiceChanged.next("update")

					console.log("umlsc saved")
				});
		} else {
			switch (association) {
				// insertion point for saving value of ONE_MANY association reverse pointer
				case "Pkgelt_Umlscs":
					this.umlsc.Pkgelt_UmlscsDBID = new NullInt64
					this.umlsc.Pkgelt_UmlscsDBID.Int64 = id
					this.umlsc.Pkgelt_UmlscsDBID.Valid = true
					break
			}
			this.umlscService.postUmlsc(this.umlsc).subscribe(umlsc => {

				this.umlscService.UmlscServiceChanged.next("post")

				this.umlsc = {} // reset fields
				console.log("umlsc added")
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
			ID: this.umlsc.ID,
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
