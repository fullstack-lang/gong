// generated from NgDetailTemplateTS
import { Component, OnInit } from '@angular/core';
import { FormControl } from '@angular/forms';

import { GongBasicFieldDB } from '../gongbasicfield-db'
import { GongBasicFieldService } from '../gongbasicfield.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'
import { MapOfComponents } from '../map-components'

// insertion point for imports

import { Router, RouterState, ActivatedRoute } from '@angular/router';

import { MatDialog, MAT_DIALOG_DATA, MatDialogRef, MatDialogConfig } from '@angular/material/dialog';

import { NullInt64 } from '../front-repo.service'

@Component({
	selector: 'app-gongbasicfield-detail',
	templateUrl: './gongbasicfield-detail.component.html',
	styleUrls: ['./gongbasicfield-detail.component.css'],
})
export class GongBasicFieldDetailComponent implements OnInit {

	// insertion point for declarations

	// the GongBasicFieldDB of interest
	gongbasicfield: GongBasicFieldDB;

	// front repo
	frontRepo: FrontRepo

	constructor(
		private gongbasicfieldService: GongBasicFieldService,
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
		this.getGongBasicField()

		// observable for changes in structs
		this.gongbasicfieldService.GongBasicFieldServiceChanged.subscribe(
			message => {
				if (message == "post" || message == "update" || message == "delete") {
					this.getGongBasicField()
				}
			}
		)

		// insertion point for initialisation of enums list
	}

	getGongBasicField(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		const association = this.route.snapshot.paramMap.get('association');

		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo
				console.log("front repo GongBasicFieldPull returned")

				if (id != 0 && association == undefined) {
					this.gongbasicfield = frontRepo.GongBasicFields.get(id)
				} else {
					this.gongbasicfield = new (GongBasicFieldDB)
				}

				// insertion point for recovery of form controls value for bool fields
			}
		)


	}

	save(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		const association = this.route.snapshot.paramMap.get('association');

		// insertion point for saving value of form controls of boolean fields
		if (this.gongbasicfield.GongEnumID == undefined) {
			this.gongbasicfield.GongEnumID = new NullInt64
		}
		if (this.gongbasicfield.GongEnum != undefined) {
			this.gongbasicfield.GongEnumID.Int64 = this.gongbasicfield.GongEnum.ID
			this.gongbasicfield.GongEnumID.Valid = true
			this.gongbasicfield.GongEnumName = this.gongbasicfield.GongEnum.Name
		} else {
			this.gongbasicfield.GongEnumID.Int64 = 0
			this.gongbasicfield.GongEnumID.Valid = true
			this.gongbasicfield.GongEnumName = ""
		}

		if (id != 0 && association == undefined) {
			// insertion point for saving value of reverse pointers
			if (this.gongbasicfield.GongStruct_GongBasicFields_reverse != undefined) {
				this.gongbasicfield.GongStruct_GongBasicFieldsDBID = new NullInt64
				this.gongbasicfield.GongStruct_GongBasicFieldsDBID.Int64 = this.gongbasicfield.GongStruct_GongBasicFields_reverse.ID
				this.gongbasicfield.GongStruct_GongBasicFieldsDBID.Valid = true
				this.gongbasicfield.GongStruct_GongBasicFields_reverse = undefined // very important, otherwise, circular JSON
			}

			this.gongbasicfieldService.updateGongBasicField(this.gongbasicfield)
				.subscribe(gongbasicfield => {
					this.gongbasicfieldService.GongBasicFieldServiceChanged.next("update")

					console.log("gongbasicfield saved")
				});
		} else {
			switch (association) {
				// insertion point for saving value of ONE_MANY association reverse pointer
				case "GongStruct_GongBasicFields":
					this.gongbasicfield.GongStruct_GongBasicFieldsDBID = new NullInt64
					this.gongbasicfield.GongStruct_GongBasicFieldsDBID.Int64 = id
					this.gongbasicfield.GongStruct_GongBasicFieldsDBID.Valid = true
					break
			}
			this.gongbasicfieldService.postGongBasicField(this.gongbasicfield).subscribe(gongbasicfield => {

				this.gongbasicfieldService.GongBasicFieldServiceChanged.next("post")

				this.gongbasicfield = {} // reset fields
				console.log("gongbasicfield added")
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
			ID: this.gongbasicfield.ID,
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
