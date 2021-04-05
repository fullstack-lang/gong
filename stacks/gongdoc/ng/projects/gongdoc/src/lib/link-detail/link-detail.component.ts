// generated from NgDetailTemplateTS
import { Component, OnInit } from '@angular/core';
import { FormControl } from '@angular/forms';

import { LinkDB } from '../link-db'
import { LinkService } from '../link.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'
import { MapOfComponents } from '../map-components'

// insertion point for imports
import { MultiplicityTypeSelect, MultiplicityTypeList } from '../MultiplicityType'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

import { MatDialog, MAT_DIALOG_DATA, MatDialogRef, MatDialogConfig } from '@angular/material/dialog';

import { NullInt64 } from '../front-repo.service'

@Component({
	selector: 'app-link-detail',
	templateUrl: './link-detail.component.html',
	styleUrls: ['./link-detail.component.css'],
})
export class LinkDetailComponent implements OnInit {

	// insertion point for declarations
	MultiplicityTypeList: MultiplicityTypeSelect[]

	// the LinkDB of interest
	link: LinkDB;

	// front repo
	frontRepo: FrontRepo

	constructor(
		private linkService: LinkService,
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
		this.getLink()

		// observable for changes in structs
		this.linkService.LinkServiceChanged.subscribe(
			message => {
				if (message == "post" || message == "update" || message == "delete") {
					this.getLink()
				}
			}
		)

		// insertion point for initialisation of enums list
		this.MultiplicityTypeList = MultiplicityTypeList
	}

	getLink(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		const association = this.route.snapshot.paramMap.get('association');

		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo
				console.log("front repo LinkPull returned")

				if (id != 0 && association == undefined) {
					this.link = frontRepo.Links.get(id)
				} else {
					this.link = new (LinkDB)
				}

				// insertion point for recovery of form controls value for bool fields
			}
		)


	}

	save(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		const association = this.route.snapshot.paramMap.get('association');

		// insertion point for saving value of form controls of boolean fields
		if (this.link.MiddleverticeID == undefined) {
			this.link.MiddleverticeID = new NullInt64
		}
		if (this.link.Middlevertice != undefined) {
			this.link.MiddleverticeID.Int64 = this.link.Middlevertice.ID
			this.link.MiddleverticeID.Valid = true
			this.link.MiddleverticeName = this.link.Middlevertice.Name
		} else {
			this.link.MiddleverticeID.Int64 = 0
			this.link.MiddleverticeID.Valid = true
			this.link.MiddleverticeName = ""
		}

		if (id != 0 && association == undefined) {
			// insertion point for saving value of reverse pointers
			if (this.link.Classshape_Links_reverse != undefined) {
				this.link.Classshape_LinksDBID = new NullInt64
				this.link.Classshape_LinksDBID.Int64 = this.link.Classshape_Links_reverse.ID
				this.link.Classshape_LinksDBID.Valid = true
				this.link.Classshape_Links_reverse = undefined // very important, otherwise, circular JSON
			}

			this.linkService.updateLink(this.link)
				.subscribe(link => {
					this.linkService.LinkServiceChanged.next("update")

					console.log("link saved")
				});
		} else {
			switch (association) {
				// insertion point for saving value of ONE_MANY association reverse pointer
				case "Classshape_Links":
					this.link.Classshape_LinksDBID = new NullInt64
					this.link.Classshape_LinksDBID.Int64 = id
					this.link.Classshape_LinksDBID.Valid = true
					break
			}
			this.linkService.postLink(this.link).subscribe(link => {

				this.linkService.LinkServiceChanged.next("post")

				this.link = {} // reset fields
				console.log("link added")
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
			ID: this.link.ID,
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
