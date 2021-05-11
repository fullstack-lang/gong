// generated from NgDetailTemplateTS
import { Component, OnInit } from '@angular/core';
import { FormControl } from '@angular/forms';

import { BclassDB } from '../bclass-db'
import { BclassService } from '../bclass.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'
import { MapOfComponents } from '../map-components'
import { MapOfSortingComponents } from '../map-components'

// insertion point for imports

import { Router, RouterState, ActivatedRoute } from '@angular/router';

import { MatDialog, MAT_DIALOG_DATA, MatDialogRef, MatDialogConfig } from '@angular/material/dialog';

import { NullInt64 } from '../front-repo.service'

@Component({
	selector: 'app-bclass-detail',
	templateUrl: './bclass-detail.component.html',
	styleUrls: ['./bclass-detail.component.css'],
})
export class BclassDetailComponent implements OnInit {

	// insertion point for declarations

	// the BclassDB of interest
	bclass: BclassDB;

	// front repo
	frontRepo: FrontRepo

	constructor(
		private bclassService: BclassService,
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
		this.getBclass()

		// observable for changes in structs
		this.bclassService.BclassServiceChanged.subscribe(
			message => {
				if (message == "post" || message == "update" || message == "delete") {
					this.getBclass()
				}
			}
		)

		// insertion point for initialisation of enums list
	}

	getBclass(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		const association = this.route.snapshot.paramMap.get('association');

		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo
				console.log("front repo BclassPull returned")

				if (id != 0 && association == undefined) {
					this.bclass = frontRepo.Bclasss.get(id)
				} else {
					this.bclass = new (BclassDB)
				}

				// insertion point for recovery of form controls value for bool fields
			}
		)


	}

	save(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		const association = this.route.snapshot.paramMap.get('association');

		// some fields needs to be translated into serializable forms
		// pointers fields, after the translation, are nulled in order to perform serialization
		
		// insertion point for translation/nullation of each field
		
		// save from the front pointer space to the non pointer space for serialization
		if (association == undefined) {
			// insertion point for translation/nullation of each pointers
			if (this.bclass.Aclass_Anarrayofb_reverse != undefined) {
				this.bclass.Aclass_AnarrayofbDBID = new NullInt64
				this.bclass.Aclass_AnarrayofbDBID.Int64 = this.bclass.Aclass_Anarrayofb_reverse.ID
				this.bclass.Aclass_AnarrayofbDBID.Valid = true
				this.bclass.Aclass_AnarrayofbDBID_Index.Valid = true
				this.bclass.Aclass_Anarrayofb_reverse = undefined // very important, otherwise, circular JSON
			}
			if (this.bclass.Aclass_Anotherarrayofb_reverse != undefined) {
				this.bclass.Aclass_AnotherarrayofbDBID = new NullInt64
				this.bclass.Aclass_AnotherarrayofbDBID.Int64 = this.bclass.Aclass_Anotherarrayofb_reverse.ID
				this.bclass.Aclass_AnotherarrayofbDBID.Valid = true
				this.bclass.Aclass_AnotherarrayofbDBID_Index.Valid = true
				this.bclass.Aclass_Anotherarrayofb_reverse = undefined // very important, otherwise, circular JSON
			}
		}

		if (id != 0 && association == undefined) {

			this.bclassService.updateBclass(this.bclass)
				.subscribe(bclass => {
					this.bclassService.BclassServiceChanged.next("update")

					console.log("bclass saved")
				});
		} else {
			switch (association) {
				// insertion point for saving value of ONE_MANY association reverse pointer
				case "Aclass_Anarrayofb":
					this.bclass.Aclass_AnarrayofbDBID = new NullInt64
					this.bclass.Aclass_AnarrayofbDBID.Int64 = id
					this.bclass.Aclass_AnarrayofbDBID.Valid = true
					this.bclass.Aclass_AnarrayofbDBID_Index.Valid = true
					break
				case "Aclass_Anotherarrayofb":
					this.bclass.Aclass_AnotherarrayofbDBID = new NullInt64
					this.bclass.Aclass_AnotherarrayofbDBID.Int64 = id
					this.bclass.Aclass_AnotherarrayofbDBID.Valid = true
					this.bclass.Aclass_AnotherarrayofbDBID_Index.Valid = true
					break
			}
			this.bclassService.postBclass(this.bclass).subscribe(bclass => {

				this.bclassService.BclassServiceChanged.next("post")

				this.bclass = {} // reset fields
				console.log("bclass added")
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
			ID: this.bclass.ID,
			ReversePointer: reverseField,
			OrderingMode: false,
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

	openDragAndDropOrdering(AssociatedStruct: string, reverseField: string) {

		const dialogConfig = new MatDialogConfig();

		// dialogConfig.disableClose = true;
		dialogConfig.autoFocus = true;
		dialogConfig.data = {
			ID: this.bclass.ID,
			ReversePointer: reverseField,
			OrderingMode: true,
		};
		const dialogRef: MatDialogRef<string, any> = this.dialog.open(
			MapOfSortingComponents.get(AssociatedStruct).get(
				AssociatedStruct + 'SortingComponent'
			),
			dialogConfig
		);

		dialogRef.afterClosed().subscribe(result => {
			console.log('The dialog was closed');
		});
	}
}
