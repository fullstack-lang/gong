// generated from NgDetailTemplateTS
import { Component, OnInit } from '@angular/core';
import { FormControl } from '@angular/forms';

import { AclassDB } from '../aclass-db'
import { AclassService } from '../aclass.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'
import { MapOfComponents } from '../map-components'

// insertion point for imports
import { AEnumTypeSelect, AEnumTypeList } from '../AEnumType'
import { BEnumTypeSelect, BEnumTypeList } from '../BEnumType'

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
	AEnumTypeList: AEnumTypeSelect[]
	BEnumTypeList: BEnumTypeSelect[]
	AnotherbooleanfieldFormControl = new FormControl(false);

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
		this.AEnumTypeList = AEnumTypeList
		this.BEnumTypeList = BEnumTypeList
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
			}
		)


	}

	save(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		const association = this.route.snapshot.paramMap.get('association');

		// insertion point for saving value of form controls of boolean fields
		this.aclass.Booleanfield = this.BooleanfieldFormControl.value
		this.aclass.Anotherbooleanfield = this.AnotherbooleanfieldFormControl.value
		if (this.aclass.AssociationtobID == undefined) {
			this.aclass.AssociationtobID = new NullInt64
		}
		if (this.aclass.Associationtob != undefined) {
			this.aclass.AssociationtobID.Int64 = this.aclass.Associationtob.ID
			this.aclass.AssociationtobID.Valid = true
			this.aclass.AssociationtobName = this.aclass.Associationtob.Name
		} else {
			this.aclass.AssociationtobID.Int64 = 0
			this.aclass.AssociationtobID.Valid = true
			this.aclass.AssociationtobName = ""
		}
		if (this.aclass.Anotherassociationtob_2ID == undefined) {
			this.aclass.Anotherassociationtob_2ID = new NullInt64
		}
		if (this.aclass.Anotherassociationtob_2 != undefined) {
			this.aclass.Anotherassociationtob_2ID.Int64 = this.aclass.Anotherassociationtob_2.ID
			this.aclass.Anotherassociationtob_2ID.Valid = true
			this.aclass.Anotherassociationtob_2Name = this.aclass.Anotherassociationtob_2.Name
		} else {
			this.aclass.Anotherassociationtob_2ID.Int64 = 0
			this.aclass.Anotherassociationtob_2ID.Valid = true
			this.aclass.Anotherassociationtob_2Name = ""
		}

		if (id != 0 && association == undefined) {
			// insertion point for saving value of reverse pointers
			if (this.aclass.Aclass_Anarrayofa_reverse != undefined) {
				this.aclass.Aclass_AnarrayofaDBID = new NullInt64
				this.aclass.Aclass_AnarrayofaDBID.Int64 = this.aclass.Aclass_Anarrayofa_reverse.ID
				this.aclass.Aclass_AnarrayofaDBID.Valid = true
				this.aclass.Aclass_Anarrayofa_reverse = undefined // very important, otherwise, circular JSON
			}

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
