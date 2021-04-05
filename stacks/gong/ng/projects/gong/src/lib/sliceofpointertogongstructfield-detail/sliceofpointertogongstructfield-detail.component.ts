// generated from NgDetailTemplateTS
import { Component, OnInit } from '@angular/core';
import { FormControl } from '@angular/forms';

import { SliceOfPointerToGongStructFieldDB } from '../sliceofpointertogongstructfield-db'
import { SliceOfPointerToGongStructFieldService } from '../sliceofpointertogongstructfield.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'
import { MapOfComponents } from '../map-components'

// insertion point for imports

import { Router, RouterState, ActivatedRoute } from '@angular/router';

import { MatDialog, MAT_DIALOG_DATA, MatDialogRef, MatDialogConfig } from '@angular/material/dialog';

import { NullInt64 } from '../front-repo.service'

@Component({
	selector: 'app-sliceofpointertogongstructfield-detail',
	templateUrl: './sliceofpointertogongstructfield-detail.component.html',
	styleUrls: ['./sliceofpointertogongstructfield-detail.component.css'],
})
export class SliceOfPointerToGongStructFieldDetailComponent implements OnInit {

	// insertion point for declarations

	// the SliceOfPointerToGongStructFieldDB of interest
	sliceofpointertogongstructfield: SliceOfPointerToGongStructFieldDB;

	// front repo
	frontRepo: FrontRepo

	constructor(
		private sliceofpointertogongstructfieldService: SliceOfPointerToGongStructFieldService,
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
		this.getSliceOfPointerToGongStructField()

		// observable for changes in structs
		this.sliceofpointertogongstructfieldService.SliceOfPointerToGongStructFieldServiceChanged.subscribe(
			message => {
				if (message == "post" || message == "update" || message == "delete") {
					this.getSliceOfPointerToGongStructField()
				}
			}
		)

		// insertion point for initialisation of enums list
	}

	getSliceOfPointerToGongStructField(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		const association = this.route.snapshot.paramMap.get('association');

		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo
				console.log("front repo SliceOfPointerToGongStructFieldPull returned")

				if (id != 0 && association == undefined) {
					this.sliceofpointertogongstructfield = frontRepo.SliceOfPointerToGongStructFields.get(id)
				} else {
					this.sliceofpointertogongstructfield = new (SliceOfPointerToGongStructFieldDB)
				}

				// insertion point for recovery of form controls value for bool fields
			}
		)


	}

	save(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		const association = this.route.snapshot.paramMap.get('association');

		// insertion point for saving value of form controls of boolean fields
		if (this.sliceofpointertogongstructfield.GongStructID == undefined) {
			this.sliceofpointertogongstructfield.GongStructID = new NullInt64
		}
		if (this.sliceofpointertogongstructfield.GongStruct != undefined) {
			this.sliceofpointertogongstructfield.GongStructID.Int64 = this.sliceofpointertogongstructfield.GongStruct.ID
			this.sliceofpointertogongstructfield.GongStructID.Valid = true
			this.sliceofpointertogongstructfield.GongStructName = this.sliceofpointertogongstructfield.GongStruct.Name
		} else {
			this.sliceofpointertogongstructfield.GongStructID.Int64 = 0
			this.sliceofpointertogongstructfield.GongStructID.Valid = true
			this.sliceofpointertogongstructfield.GongStructName = ""
		}

		if (id != 0 && association == undefined) {
			// insertion point for saving value of reverse pointers
			if (this.sliceofpointertogongstructfield.GongStruct_SliceOfPointerToGongStructFields_reverse != undefined) {
				this.sliceofpointertogongstructfield.GongStruct_SliceOfPointerToGongStructFieldsDBID = new NullInt64
				this.sliceofpointertogongstructfield.GongStruct_SliceOfPointerToGongStructFieldsDBID.Int64 = this.sliceofpointertogongstructfield.GongStruct_SliceOfPointerToGongStructFields_reverse.ID
				this.sliceofpointertogongstructfield.GongStruct_SliceOfPointerToGongStructFieldsDBID.Valid = true
				this.sliceofpointertogongstructfield.GongStruct_SliceOfPointerToGongStructFields_reverse = undefined // very important, otherwise, circular JSON
			}

			this.sliceofpointertogongstructfieldService.updateSliceOfPointerToGongStructField(this.sliceofpointertogongstructfield)
				.subscribe(sliceofpointertogongstructfield => {
					this.sliceofpointertogongstructfieldService.SliceOfPointerToGongStructFieldServiceChanged.next("update")

					console.log("sliceofpointertogongstructfield saved")
				});
		} else {
			switch (association) {
				// insertion point for saving value of ONE_MANY association reverse pointer
				case "GongStruct_SliceOfPointerToGongStructFields":
					this.sliceofpointertogongstructfield.GongStruct_SliceOfPointerToGongStructFieldsDBID = new NullInt64
					this.sliceofpointertogongstructfield.GongStruct_SliceOfPointerToGongStructFieldsDBID.Int64 = id
					this.sliceofpointertogongstructfield.GongStruct_SliceOfPointerToGongStructFieldsDBID.Valid = true
					break
			}
			this.sliceofpointertogongstructfieldService.postSliceOfPointerToGongStructField(this.sliceofpointertogongstructfield).subscribe(sliceofpointertogongstructfield => {

				this.sliceofpointertogongstructfieldService.SliceOfPointerToGongStructFieldServiceChanged.next("post")

				this.sliceofpointertogongstructfield = {} // reset fields
				console.log("sliceofpointertogongstructfield added")
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
			ID: this.sliceofpointertogongstructfield.ID,
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
