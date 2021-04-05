// generated from NgDetailTemplateTS
import { Component, OnInit } from '@angular/core';
import { FormControl } from '@angular/forms';

import { FieldDB } from '../field-db'
import { FieldService } from '../field.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'
import { MapOfComponents } from '../map-components'

// insertion point for imports

import { Router, RouterState, ActivatedRoute } from '@angular/router';

import { MatDialog, MAT_DIALOG_DATA, MatDialogRef, MatDialogConfig } from '@angular/material/dialog';

import { NullInt64 } from '../front-repo.service'

@Component({
	selector: 'app-field-detail',
	templateUrl: './field-detail.component.html',
	styleUrls: ['./field-detail.component.css'],
})
export class FieldDetailComponent implements OnInit {

	// insertion point for declarations

	// the FieldDB of interest
	field: FieldDB;

	// front repo
	frontRepo: FrontRepo

	constructor(
		private fieldService: FieldService,
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
		this.getField()

		// observable for changes in structs
		this.fieldService.FieldServiceChanged.subscribe(
			message => {
				if (message == "post" || message == "update" || message == "delete") {
					this.getField()
				}
			}
		)

		// insertion point for initialisation of enums list
	}

	getField(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		const association = this.route.snapshot.paramMap.get('association');

		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo
				console.log("front repo FieldPull returned")

				if (id != 0 && association == undefined) {
					this.field = frontRepo.Fields.get(id)
				} else {
					this.field = new (FieldDB)
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
			if (this.field.Classshape_Fields_reverse != undefined) {
				this.field.Classshape_FieldsDBID = new NullInt64
				this.field.Classshape_FieldsDBID.Int64 = this.field.Classshape_Fields_reverse.ID
				this.field.Classshape_FieldsDBID.Valid = true
				this.field.Classshape_Fields_reverse = undefined // very important, otherwise, circular JSON
			}

			this.fieldService.updateField(this.field)
				.subscribe(field => {
					this.fieldService.FieldServiceChanged.next("update")

					console.log("field saved")
				});
		} else {
			switch (association) {
				// insertion point for saving value of ONE_MANY association reverse pointer
				case "Classshape_Fields":
					this.field.Classshape_FieldsDBID = new NullInt64
					this.field.Classshape_FieldsDBID.Int64 = id
					this.field.Classshape_FieldsDBID.Valid = true
					break
			}
			this.fieldService.postField(this.field).subscribe(field => {

				this.fieldService.FieldServiceChanged.next("post")

				this.field = {} // reset fields
				console.log("field added")
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
			ID: this.field.ID,
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
