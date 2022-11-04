import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';

import { ReferenceDB } from '../reference-db'
import { ReferenceService } from '../reference.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

// insertion point for additional imports

export interface referenceDummyElement {
}

const ELEMENT_DATA: referenceDummyElement[] = [
];

@Component({
	selector: 'app-reference-presentation',
	templateUrl: './reference-presentation.component.html',
	styleUrls: ['./reference-presentation.component.css'],
})
export class ReferencePresentationComponent implements OnInit {

	// insertion point for additionnal time duration declarations
	// insertion point for additionnal enum int field declarations

	displayedColumns: string[] = []
	dataSource = ELEMENT_DATA

	reference: ReferenceDB = new (ReferenceDB)

	// front repo
	frontRepo: FrontRepo = new (FrontRepo)
 
	constructor(
		private referenceService: ReferenceService,
		private frontRepoService: FrontRepoService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getReference();

		// observable for changes in 
		this.referenceService.ReferenceServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getReference()
				}
			}
		)
	}

	getReference(): void {
		const id = +this.route.snapshot.paramMap.get('id')!
		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				this.reference = this.frontRepo.References.get(id)!

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
				github_com_fullstack_lang_gongdoc_go_editor: ["github_com_fullstack_lang_gongdoc_go-" + "reference-detail", ID]
			}
		}]);
	}
}
