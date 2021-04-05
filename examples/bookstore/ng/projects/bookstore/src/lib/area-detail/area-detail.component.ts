// generated from NgDetailTemplateTS
import { Component, OnInit } from '@angular/core';
import { FormControl } from '@angular/forms';

import { AreaDB } from '../area-db'
import { AreaService } from '../area.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'
import { MapOfComponents } from '../map-components'

// insertion point for imports

import { Router, RouterState, ActivatedRoute } from '@angular/router';

import { MatDialog, MAT_DIALOG_DATA, MatDialogRef, MatDialogConfig } from '@angular/material/dialog';

import { NullInt64 } from '../front-repo.service'

@Component({
	selector: 'app-area-detail',
	templateUrl: './area-detail.component.html',
	styleUrls: ['./area-detail.component.css'],
})
export class AreaDetailComponent implements OnInit {

	// insertion point for declarations

	// the AreaDB of interest
	area: AreaDB;

	// front repo
	frontRepo: FrontRepo

	constructor(
		private areaService: AreaService,
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
		this.getArea()

		// observable for changes in structs
		this.areaService.AreaServiceChanged.subscribe(
			message => {
				if (message == "post" || message == "update" || message == "delete") {
					this.getArea()
				}
			}
		)

		// insertion point for initialisation of enums list
	}

	getArea(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		const association = this.route.snapshot.paramMap.get('association');

		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo
				console.log("front repo AreaPull returned")

				if (id != 0 && association == undefined) {
					this.area = frontRepo.Areas.get(id)
				} else {
					this.area = new (AreaDB)
				}

				// insertion point for recovery of form controls value for bool fields
			}
		)


	}

	save(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		const association = this.route.snapshot.paramMap.get('association');

		// insertion point for saving value of form controls of boolean fields
		if (this.area.SubareaID == undefined) {
			this.area.SubareaID = new NullInt64
		}
		if (this.area.Subarea != undefined) {
			this.area.SubareaID.Int64 = this.area.Subarea.ID
			this.area.SubareaID.Valid = true
			this.area.SubareaName = this.area.Subarea.Name
		} else {
			this.area.SubareaID.Int64 = 0
			this.area.SubareaID.Valid = true
			this.area.SubareaName = ""
		}

		if (id != 0 && association == undefined) {
			// insertion point for saving value of reverse pointers

			this.areaService.updateArea(this.area)
				.subscribe(area => {
					this.areaService.AreaServiceChanged.next("update")

					console.log("area saved")
				});
		} else {
			switch (association) {
				// insertion point for saving value of ONE_MANY association reverse pointer
			}
			this.areaService.postArea(this.area).subscribe(area => {

				this.areaService.AreaServiceChanged.next("post")

				this.area = {} // reset fields
				console.log("area added")
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
			ID: this.area.ID,
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
