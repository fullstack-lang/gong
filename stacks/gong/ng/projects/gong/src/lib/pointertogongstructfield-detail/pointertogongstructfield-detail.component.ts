// generated from NgDetailTemplateTS
import { Component, OnInit } from '@angular/core';
import { FormControl } from '@angular/forms';

import { PointerToGongStructFieldDB } from '../pointertogongstructfield-db'
import { PointerToGongStructFieldService } from '../pointertogongstructfield.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'
import { MapOfComponents } from '../map-components'

// insertion point for imports

import { Router, RouterState, ActivatedRoute } from '@angular/router';

import { MatDialog, MAT_DIALOG_DATA, MatDialogRef, MatDialogConfig } from '@angular/material/dialog';

import { NullInt64 } from '../front-repo.service'

@Component({
	selector: 'app-pointertogongstructfield-detail',
	templateUrl: './pointertogongstructfield-detail.component.html',
	styleUrls: ['./pointertogongstructfield-detail.component.css'],
})
export class PointerToGongStructFieldDetailComponent implements OnInit {

	// insertion point for declarations

	// the PointerToGongStructFieldDB of interest
	pointertogongstructfield: PointerToGongStructFieldDB;

	// front repo
	frontRepo: FrontRepo

	constructor(
		private pointertogongstructfieldService: PointerToGongStructFieldService,
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
		this.getPointerToGongStructField()

		// observable for changes in structs
		this.pointertogongstructfieldService.PointerToGongStructFieldServiceChanged.subscribe(
			message => {
				if (message == "post" || message == "update" || message == "delete") {
					this.getPointerToGongStructField()
				}
			}
		)

		// insertion point for initialisation of enums list
	}

	getPointerToGongStructField(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		const association = this.route.snapshot.paramMap.get('association');

		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo
				console.log("front repo PointerToGongStructFieldPull returned")

				if (id != 0 && association == undefined) {
					this.pointertogongstructfield = frontRepo.PointerToGongStructFields.get(id)
				} else {
					this.pointertogongstructfield = new (PointerToGongStructFieldDB)
				}

				// insertion point for recovery of form controls value for bool fields
			}
		)


	}

	save(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		const association = this.route.snapshot.paramMap.get('association');

		// insertion point for saving value of form controls of boolean fields
		if (this.pointertogongstructfield.GongStructID == undefined) {
			this.pointertogongstructfield.GongStructID = new NullInt64
		}
		if (this.pointertogongstructfield.GongStruct != undefined) {
			this.pointertogongstructfield.GongStructID.Int64 = this.pointertogongstructfield.GongStruct.ID
			this.pointertogongstructfield.GongStructID.Valid = true
			this.pointertogongstructfield.GongStructName = this.pointertogongstructfield.GongStruct.Name
		} else {
			this.pointertogongstructfield.GongStructID.Int64 = 0
			this.pointertogongstructfield.GongStructID.Valid = true
			this.pointertogongstructfield.GongStructName = ""
		}

		if (id != 0 && association == undefined) {
			// insertion point for saving value of reverse pointers
			if (this.pointertogongstructfield.GongStruct_PointerToGongStructFields_reverse != undefined) {
				this.pointertogongstructfield.GongStruct_PointerToGongStructFieldsDBID = new NullInt64
				this.pointertogongstructfield.GongStruct_PointerToGongStructFieldsDBID.Int64 = this.pointertogongstructfield.GongStruct_PointerToGongStructFields_reverse.ID
				this.pointertogongstructfield.GongStruct_PointerToGongStructFieldsDBID.Valid = true
				this.pointertogongstructfield.GongStruct_PointerToGongStructFields_reverse = undefined // very important, otherwise, circular JSON
			}

			this.pointertogongstructfieldService.updatePointerToGongStructField(this.pointertogongstructfield)
				.subscribe(pointertogongstructfield => {
					this.pointertogongstructfieldService.PointerToGongStructFieldServiceChanged.next("update")

					console.log("pointertogongstructfield saved")
				});
		} else {
			switch (association) {
				// insertion point for saving value of ONE_MANY association reverse pointer
				case "GongStruct_PointerToGongStructFields":
					this.pointertogongstructfield.GongStruct_PointerToGongStructFieldsDBID = new NullInt64
					this.pointertogongstructfield.GongStruct_PointerToGongStructFieldsDBID.Int64 = id
					this.pointertogongstructfield.GongStruct_PointerToGongStructFieldsDBID.Valid = true
					break
			}
			this.pointertogongstructfieldService.postPointerToGongStructField(this.pointertogongstructfield).subscribe(pointertogongstructfield => {

				this.pointertogongstructfieldService.PointerToGongStructFieldServiceChanged.next("post")

				this.pointertogongstructfield = {} // reset fields
				console.log("pointertogongstructfield added")
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
			ID: this.pointertogongstructfield.ID,
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
