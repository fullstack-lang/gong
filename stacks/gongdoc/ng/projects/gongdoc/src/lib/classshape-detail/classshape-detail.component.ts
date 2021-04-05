// generated from NgDetailTemplateTS
import { Component, OnInit } from '@angular/core';
import { FormControl } from '@angular/forms';

import { ClassshapeDB } from '../classshape-db'
import { ClassshapeService } from '../classshape.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'
import { MapOfComponents } from '../map-components'

// insertion point for imports
import { ClassshapeTargetTypeSelect, ClassshapeTargetTypeList } from '../ClassshapeTargetType'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

import { MatDialog, MAT_DIALOG_DATA, MatDialogRef, MatDialogConfig } from '@angular/material/dialog';

import { NullInt64 } from '../front-repo.service'

@Component({
	selector: 'app-classshape-detail',
	templateUrl: './classshape-detail.component.html',
	styleUrls: ['./classshape-detail.component.css'],
})
export class ClassshapeDetailComponent implements OnInit {

	// insertion point for declarations
	ClassshapeTargetTypeList: ClassshapeTargetTypeSelect[]

	// the ClassshapeDB of interest
	classshape: ClassshapeDB;

	// front repo
	frontRepo: FrontRepo

	constructor(
		private classshapeService: ClassshapeService,
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
		this.getClassshape()

		// observable for changes in structs
		this.classshapeService.ClassshapeServiceChanged.subscribe(
			message => {
				if (message == "post" || message == "update" || message == "delete") {
					this.getClassshape()
				}
			}
		)

		// insertion point for initialisation of enums list
		this.ClassshapeTargetTypeList = ClassshapeTargetTypeList
	}

	getClassshape(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		const association = this.route.snapshot.paramMap.get('association');

		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo
				console.log("front repo ClassshapePull returned")

				if (id != 0 && association == undefined) {
					this.classshape = frontRepo.Classshapes.get(id)
				} else {
					this.classshape = new (ClassshapeDB)
				}

				// insertion point for recovery of form controls value for bool fields
			}
		)


	}

	save(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		const association = this.route.snapshot.paramMap.get('association');

		// insertion point for saving value of form controls of boolean fields
		if (this.classshape.PositionID == undefined) {
			this.classshape.PositionID = new NullInt64
		}
		if (this.classshape.Position != undefined) {
			this.classshape.PositionID.Int64 = this.classshape.Position.ID
			this.classshape.PositionID.Valid = true
			this.classshape.PositionName = this.classshape.Position.Name
		} else {
			this.classshape.PositionID.Int64 = 0
			this.classshape.PositionID.Valid = true
			this.classshape.PositionName = ""
		}

		if (id != 0 && association == undefined) {
			// insertion point for saving value of reverse pointers
			if (this.classshape.Classdiagram_Classshapes_reverse != undefined) {
				this.classshape.Classdiagram_ClassshapesDBID = new NullInt64
				this.classshape.Classdiagram_ClassshapesDBID.Int64 = this.classshape.Classdiagram_Classshapes_reverse.ID
				this.classshape.Classdiagram_ClassshapesDBID.Valid = true
				this.classshape.Classdiagram_Classshapes_reverse = undefined // very important, otherwise, circular JSON
			}

			this.classshapeService.updateClassshape(this.classshape)
				.subscribe(classshape => {
					this.classshapeService.ClassshapeServiceChanged.next("update")

					console.log("classshape saved")
				});
		} else {
			switch (association) {
				// insertion point for saving value of ONE_MANY association reverse pointer
				case "Classdiagram_Classshapes":
					this.classshape.Classdiagram_ClassshapesDBID = new NullInt64
					this.classshape.Classdiagram_ClassshapesDBID.Int64 = id
					this.classshape.Classdiagram_ClassshapesDBID.Valid = true
					break
			}
			this.classshapeService.postClassshape(this.classshape).subscribe(classshape => {

				this.classshapeService.ClassshapeServiceChanged.next("post")

				this.classshape = {} // reset fields
				console.log("classshape added")
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
			ID: this.classshape.ID,
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
