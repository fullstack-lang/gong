import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';

import { GongTimeFieldDB } from '../gongtimefield-db'
import { GongTimeFieldService } from '../gongtimefield.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

// insertion point for additional imports

export interface gongtimefieldDummyElement {
}

const ELEMENT_DATA: gongtimefieldDummyElement[] = [
];

@Component({
	selector: 'app-gongtimefield-presentation',
	templateUrl: './gongtimefield-presentation.component.html',
	styleUrls: ['./gongtimefield-presentation.component.css'],
})
export class GongTimeFieldPresentationComponent implements OnInit {

	// insertion point for additionnal time duration declarations
	// insertion point for additionnal enum int field declarations

	displayedColumns: string[] = []
	dataSource = ELEMENT_DATA

	gongtimefield: GongTimeFieldDB = new (GongTimeFieldDB)

	// front repo
	frontRepo: FrontRepo = new (FrontRepo)
 
	constructor(
		private gongtimefieldService: GongTimeFieldService,
		private frontRepoService: FrontRepoService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getGongTimeField();

		// observable for changes in 
		this.gongtimefieldService.GongTimeFieldServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getGongTimeField()
				}
			}
		)
	}

	getGongTimeField(): void {
		const id = +this.route.snapshot.paramMap.get('id')!
		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				this.gongtimefield = this.frontRepo.GongTimeFields.get(id)!

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
				github_com_fullstack_lang_gong_go_editor: ["github_com_fullstack_lang_gong_go-" + "gongtimefield-detail", ID]
			}
		}]);
	}
}
