import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';

import { PkgeltDB } from '../pkgelt-db'
import { PkgeltService } from '../pkgelt.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

// insertion point for additional imports

export interface pkgeltDummyElement {
}

const ELEMENT_DATA: pkgeltDummyElement[] = [
];

@Component({
	selector: 'app-pkgelt-presentation',
	templateUrl: './pkgelt-presentation.component.html',
	styleUrls: ['./pkgelt-presentation.component.css'],
})
export class PkgeltPresentationComponent implements OnInit {

	// insertion point for additionnal time duration declarations
	// insertion point for additionnal enum int field declarations

	displayedColumns: string[] = []
	dataSource = ELEMENT_DATA

	pkgelt: PkgeltDB = new (PkgeltDB)

	// front repo
	frontRepo: FrontRepo = new (FrontRepo)
 
	constructor(
		private pkgeltService: PkgeltService,
		private frontRepoService: FrontRepoService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getPkgelt();

		// observable for changes in 
		this.pkgeltService.PkgeltServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getPkgelt()
				}
			}
		)
	}

	getPkgelt(): void {
		const id = +this.route.snapshot.paramMap.get('id')!
		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				this.pkgelt = this.frontRepo.Pkgelts.get(id)!

				// insertion point for recovery of durations
				// insertion point for recovery of enum tint
			}
		);
	}

	// set presentation outlet
	setPresentationRouterOutlet(structName: string, ID: number) {
		this.router.navigate([{
			outlets: {
				github_com_fullstack_lang_gongdoc_go_presentation: ["github_com_fullstack_lang_gongdoc_go-" + structName + "-presentation", ID]
			}
		}]);
	}

	// set editor outlet
	setEditorRouterOutlet(ID: number) {
		this.router.navigate([{
			outlets: {
				github_com_fullstack_lang_gongdoc_go_editor: ["github_com_fullstack_lang_gongdoc_go-" + "pkgelt-detail", ID]
			}
		}]);
	}
}
