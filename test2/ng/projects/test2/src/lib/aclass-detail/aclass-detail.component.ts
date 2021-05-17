// generated from NgDetailTemplateTS
import { Component, OnInit } from '@angular/core';
import { FormControl } from '@angular/forms';

import { AclassDB } from '../aclass-db'
import { AclassService } from '../aclass.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'
import { MapOfComponents } from '../map-components'

// insertion point for imports

import { Router, RouterState, ActivatedRoute } from '@angular/router';

import { MatDialog, MAT_DIALOG_DATA, MatDialogRef, MatDialogConfig } from '@angular/material/dialog';

import { NullInt64 } from '../front-repo.service'

@Component({
	selector: 'app-aclass-detail',
	templateUrl: './aclass-detail.component.html',
	styleUrls: ['./aclass-detail.component.css'],
})
export class AclassDetailComponent implements OnInit {

	// insertion point for declarations
	BooleanfieldFormControl = new FormControl(false);
	AnotherbooleanfieldFormControl = new FormControl(false);
	Duration1_Hours: number
	Duration1_Minutes: number
	Duration1_Seconds: number

	// the AclassDB of interest
	aclass: AclassDB;

	// front repo
	frontRepo: FrontRepo

	constructor(
		private aclassService: AclassService,
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
		this.getAclass()

		// observable for changes in structs
		this.aclassService.AclassServiceChanged.subscribe(
			message => {
				if (message == "post" || message == "update" || message == "delete") {
					this.getAclass()
				}
			}
		)

		// insertion point for initialisation of enums list
	}

	getAclass(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		const association = this.route.snapshot.paramMap.get('association');

		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo
				console.log("front repo AclassPull returned")

				if (id != 0 && association == undefined) {
					this.aclass = frontRepo.Aclasss.get(id)
				} else {
					this.aclass = new (AclassDB)
				}

				// insertion point for recovery of form controls value for bool fields
				this.BooleanfieldFormControl.setValue(this.aclass.Booleanfield)
				this.AnotherbooleanfieldFormControl.setValue(this.aclass.Anotherbooleanfield)
				// computation of Hours, Minutes, Seconds for Duration1
				this.Duration1_Hours = Math.floor(this.aclass.Duration1 / (3600 * 1000 * 1000 * 1000))
				this.Duration1_Minutes = Math.floor(this.aclass.Duration1 % (3600 * 1000 * 1000 * 1000) / (60 * 1000 * 1000 * 1000))
				this.Duration1_Seconds = this.aclass.Duration1 % (60 * 1000 * 1000 * 1000) / (1000 * 1000 * 1000)
			}
		)


	}

	save(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		const association = this.route.snapshot.paramMap.get('association');

		// some fields needs to be translated into serializable forms
		// pointers fields, after the translation, are nulled in order to perform serialization
		
		// insertion point for translation/nullation of each field
		this.aclass.Booleanfield = this.BooleanfieldFormControl.value
		this.aclass.Anotherbooleanfield = this.AnotherbooleanfieldFormControl.value
		this.aclass.Duration1 =
			this.Duration1_Hours * (3600 * 1000 * 1000 * 1000) +
			this.Duration1_Minutes * (60 * 1000 * 1000 * 1000) +
			this.Duration1_Seconds * (1000 * 1000 * 1000)
		
		// save from the front pointer space to the non pointer space for serialization
		if (association == undefined) {
			// insertion point for translation/nullation of each pointers
			if (this.aclass.Aclass_Anarrayofa_reverse != undefined) {
				this.aclass.Aclass_AnarrayofaDBID = new NullInt64
				this.aclass.Aclass_AnarrayofaDBID.Int64 = this.aclass.Aclass_Anarrayofa_reverse.ID
				this.aclass.Aclass_AnarrayofaDBID.Valid = true
				this.aclass.Aclass_Anarrayofa_reverse = undefined // very important, otherwise, circular JSON
			}
		}

		if (id != 0 && association == undefined) {

			this.aclassService.updateAclass(this.aclass)
				.subscribe(aclass => {
					this.aclassService.AclassServiceChanged.next("update")

					console.log("aclass saved")
				});
		} else {
			switch (association) {
				// insertion point for saving value of ONE_MANY association reverse pointer
				case "Aclass_Anarrayofa":
					this.aclass.Aclass_AnarrayofaDBID = new NullInt64
					this.aclass.Aclass_AnarrayofaDBID.Int64 = id
					this.aclass.Aclass_AnarrayofaDBID.Valid = true
					break
			}
			this.aclassService.postAclass(this.aclass).subscribe(aclass => {

				this.aclassService.AclassServiceChanged.next("post")

				this.aclass = {} // reset fields
				console.log("aclass added")
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
		dialogConfig.width = "50%"
		dialogConfig.height = "50%"
		dialogConfig.data = {
			ID: this.aclass.ID,
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
