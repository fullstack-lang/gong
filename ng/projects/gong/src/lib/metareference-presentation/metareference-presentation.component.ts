import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';

import { MetaReferenceDB } from '../metareference-db'
import { MetaReferenceService } from '../metareference.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

// insertion point for additional imports

export interface metareferenceDummyElement {
}

const ELEMENT_DATA: metareferenceDummyElement[] = [
];

@Component({
	selector: 'app-metareference-presentation',
	templateUrl: './metareference-presentation.component.html',
	styleUrls: ['./metareference-presentation.component.css'],
})
export class MetaReferencePresentationComponent implements OnInit {

	// insertion point for additionnal time duration declarations
	// insertion point for additionnal enum int field declarations

	displayedColumns: string[] = []
	dataSource = ELEMENT_DATA

	metareference: MetaReferenceDB = new (MetaReferenceDB)

	// front repo
	frontRepo: FrontRepo = new (FrontRepo)
 
	constructor(
		private metareferenceService: MetaReferenceService,
		private frontRepoService: FrontRepoService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getMetaReference();

		// observable for changes in 
		this.metareferenceService.MetaReferenceServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getMetaReference()
				}
			}
		)
	}

	getMetaReference(): void {
		const id = +this.route.snapshot.paramMap.get('id')!
		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				this.metareference = this.frontRepo.MetaReferences.get(id)!

				// insertion point for recovery of durations
				// insertion point for recovery of enum tint
			}
		);
	}

	// set presentation outlet
	setPresentationRouterOutlet(structName: string, ID: number) {
		this.router.navigate([{
			outlets: {
				github_com_fullstack_lang_gong_go_presentation: ["github_com_fullstack_lang_gong_go-" + structName + "-presentation", ID]
			}
		}]);
	}

	// set editor outlet
	setEditorRouterOutlet(ID: number) {
		this.router.navigate([{
			outlets: {
				github_com_fullstack_lang_gong_go_editor: ["github_com_fullstack_lang_gong_go-" + "metareference-detail", ID]
			}
		}]);
	}
}
