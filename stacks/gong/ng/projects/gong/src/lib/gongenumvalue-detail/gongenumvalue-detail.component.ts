// generated from NgDetailTemplateTS
import { Component, OnInit } from '@angular/core';
import { FormControl } from '@angular/forms';

import { GongEnumValueDB } from '../gongenumvalue-db'
import { GongEnumValueService } from '../gongenumvalue.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'
import { MapOfComponents } from '../map-components'

// insertion point for imports

import { Router, RouterState, ActivatedRoute } from '@angular/router';

import { MatDialog, MAT_DIALOG_DATA, MatDialogRef, MatDialogConfig } from '@angular/material/dialog';

import { NullInt64 } from '../front-repo.service'

@Component({
	selector: 'app-gongenumvalue-detail',
	templateUrl: './gongenumvalue-detail.component.html',
	styleUrls: ['./gongenumvalue-detail.component.css'],
})
export class GongEnumValueDetailComponent implements OnInit {

	// insertion point for declarations

	// the GongEnumValueDB of interest
	gongenumvalue: GongEnumValueDB;

	// front repo
	frontRepo: FrontRepo

	constructor(
		private gongenumvalueService: GongEnumValueService,
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
		this.getGongEnumValue()

		// observable for changes in structs
		this.gongenumvalueService.GongEnumValueServiceChanged.subscribe(
			message => {
				if (message == "post" || message == "update" || message == "delete") {
					this.getGongEnumValue()
				}
			}
		)

		// insertion point for initialisation of enums list
	}

	getGongEnumValue(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		const association = this.route.snapshot.paramMap.get('association');

		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo
				console.log("front repo GongEnumValuePull returned")

				if (id != 0 && association == undefined) {
					this.gongenumvalue = frontRepo.GongEnumValues.get(id)
				} else {
					this.gongenumvalue = new (GongEnumValueDB)
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
			if (this.gongenumvalue.GongEnum_GongEnumValues_reverse != undefined) {
				this.gongenumvalue.GongEnum_GongEnumValuesDBID = new NullInt64
				this.gongenumvalue.GongEnum_GongEnumValuesDBID.Int64 = this.gongenumvalue.GongEnum_GongEnumValues_reverse.ID
				this.gongenumvalue.GongEnum_GongEnumValuesDBID.Valid = true
				this.gongenumvalue.GongEnum_GongEnumValues_reverse = undefined // very important, otherwise, circular JSON
			}

			this.gongenumvalueService.updateGongEnumValue(this.gongenumvalue)
				.subscribe(gongenumvalue => {
					this.gongenumvalueService.GongEnumValueServiceChanged.next("update")

					console.log("gongenumvalue saved")
				});
		} else {
			switch (association) {
				// insertion point for saving value of ONE_MANY association reverse pointer
				case "GongEnum_GongEnumValues":
					this.gongenumvalue.GongEnum_GongEnumValuesDBID = new NullInt64
					this.gongenumvalue.GongEnum_GongEnumValuesDBID.Int64 = id
					this.gongenumvalue.GongEnum_GongEnumValuesDBID.Valid = true
					break
			}
			this.gongenumvalueService.postGongEnumValue(this.gongenumvalue).subscribe(gongenumvalue => {

				this.gongenumvalueService.GongEnumValueServiceChanged.next("post")

				this.gongenumvalue = {} // reset fields
				console.log("gongenumvalue added")
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
			ID: this.gongenumvalue.ID,
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
