// generated from NgDetailTemplateTS
import { Component, OnInit } from '@angular/core';
import { FormControl } from '@angular/forms';

import { ClassdiagramDB } from '../classdiagram-db'
import { ClassdiagramService } from '../classdiagram.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'
import { MapOfComponents } from '../map-components'

// insertion point for imports

import { Router, RouterState, ActivatedRoute } from '@angular/router';

import { MatDialog, MAT_DIALOG_DATA, MatDialogRef, MatDialogConfig } from '@angular/material/dialog';

import { NullInt64 } from '../front-repo.service'

@Component({
	selector: 'app-classdiagram-detail',
	templateUrl: './classdiagram-detail.component.html',
	styleUrls: ['./classdiagram-detail.component.css'],
})
export class ClassdiagramDetailComponent implements OnInit {

	// insertion point for declarations

	// the ClassdiagramDB of interest
	classdiagram: ClassdiagramDB;

	// front repo
	frontRepo: FrontRepo

	constructor(
		private classdiagramService: ClassdiagramService,
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
		this.getClassdiagram()

		// observable for changes in structs
		this.classdiagramService.ClassdiagramServiceChanged.subscribe(
			message => {
				if (message == "post" || message == "update" || message == "delete") {
					this.getClassdiagram()
				}
			}
		)

		// insertion point for initialisation of enums list
	}

	getClassdiagram(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		const association = this.route.snapshot.paramMap.get('association');

		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo
				console.log("front repo ClassdiagramPull returned")

				if (id != 0 && association == undefined) {
					this.classdiagram = frontRepo.Classdiagrams.get(id)
				} else {
					this.classdiagram = new (ClassdiagramDB)
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
			if (this.classdiagram.Pkgelt_Classdiagrams_reverse != undefined) {
				this.classdiagram.Pkgelt_ClassdiagramsDBID = new NullInt64
				this.classdiagram.Pkgelt_ClassdiagramsDBID.Int64 = this.classdiagram.Pkgelt_Classdiagrams_reverse.ID
				this.classdiagram.Pkgelt_ClassdiagramsDBID.Valid = true
				this.classdiagram.Pkgelt_Classdiagrams_reverse = undefined // very important, otherwise, circular JSON
			}

			this.classdiagramService.updateClassdiagram(this.classdiagram)
				.subscribe(classdiagram => {
					this.classdiagramService.ClassdiagramServiceChanged.next("update")

					console.log("classdiagram saved")
				});
		} else {
			switch (association) {
				// insertion point for saving value of ONE_MANY association reverse pointer
				case "Pkgelt_Classdiagrams":
					this.classdiagram.Pkgelt_ClassdiagramsDBID = new NullInt64
					this.classdiagram.Pkgelt_ClassdiagramsDBID.Int64 = id
					this.classdiagram.Pkgelt_ClassdiagramsDBID.Valid = true
					break
			}
			this.classdiagramService.postClassdiagram(this.classdiagram).subscribe(classdiagram => {

				this.classdiagramService.ClassdiagramServiceChanged.next("post")

				this.classdiagram = {} // reset fields
				console.log("classdiagram added")
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
			ID: this.classdiagram.ID,
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
