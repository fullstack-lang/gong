// generated from NgDetailTemplateTS
import { Component, OnInit } from '@angular/core';
import { FormControl } from '@angular/forms';

import { ClassdiagramDB } from 'gongdoc'
import { ClassdiagramService } from 'gongdoc'

import { PkgeltDB } from 'gongdoc'
import { PkgeltService } from 'gongdoc'



import { FrontRepoService, FrontRepo } from 'gongdoc'
import { NullInt64 } from 'gongdoc'


// insertion point for import of enums lists 

import { Router, RouterState, ActivatedRoute } from '@angular/router';

import { MatDialog, MAT_DIALOG_DATA, MatDialogRef, MatDialogConfig } from '@angular/material/dialog';


@Component({
	styleUrls: ['./classdiagram-detail.css'],
	selector: 'app-classdiagram-detail',
	templateUrl: './classdiagram-detail.component.html'
})
export class ClassdiagramDetailComponent implements OnInit {

	// insertion point for declaration of enums list 

	// insertion point for declation of form controls for bool 

	// the ClassdiagramDB of interest
	classdiagram: ClassdiagramDB = new ClassdiagramDB;

	// front repo
	frontRepo: FrontRepo = new FrontRepo

	constructor(
		private pkgeltService: PkgeltService,
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
		const id = +this.route.snapshot.paramMap.get('id')!

		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo
				console.log("front repo ClassdiagramPull returned")

				if (id != 0) {
					this.classdiagram = frontRepo.Classdiagrams.get(id)!
				} else {
					this.classdiagram = new (ClassdiagramDB)
				}

				// insertion point for recovery of form controls value for bool fields
			}
		)


	}

	save(): void {
		const id = +this.route.snapshot.paramMap.get('id')!

		if (id != 0) {

			this.classdiagramService.updateClassdiagram(this.classdiagram)
				.subscribe(classdiagram => {
					this.classdiagramService.ClassdiagramServiceChanged.next("update")

					console.log("classdiagram saved")
				});

		} else {

			let singlotonPkg = this.frontRepo.Pkgelts_array[0]
			if (singlotonPkg == undefined) {
				console.log("problem !")
				return
			}
			if (singlotonPkg.Classdiagrams == undefined) {
				singlotonPkg.Classdiagrams = new Array<ClassdiagramDB>()
			}
			singlotonPkg.Classdiagrams.concat(this.classdiagram)
			this.classdiagram.Pkgelt_ClassdiagramsDBID = new NullInt64
			this.classdiagram.Pkgelt_ClassdiagramsDBID.Int64 = singlotonPkg.ID
			this.classdiagram.Pkgelt_ClassdiagramsDBID.Valid = true

			this.classdiagramService.postClassdiagram(this.classdiagram)
				.subscribe(classdiagram => {
					this.classdiagramService.ClassdiagramServiceChanged.next("post")

					this.classdiagram = new ClassdiagramDB // reset fields
					console.log("classdiagram added")
				});
		}
	}
}
