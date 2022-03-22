import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';

import { UmlStateDB } from '../umlstate-db'
import { UmlStateService } from '../umlstate.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

// insertion point for additional imports

export interface umlstateDummyElement {
}

const ELEMENT_DATA: umlstateDummyElement[] = [
];

@Component({
	selector: 'app-umlstate-presentation',
	templateUrl: './umlstate-presentation.component.html',
	styleUrls: ['./umlstate-presentation.component.css'],
})
export class UmlStatePresentationComponent implements OnInit {

	// insertion point for additionnal time duration declarations
	// insertion point for additionnal enum int field declarations

	displayedColumns: string[] = []
	dataSource = ELEMENT_DATA

	umlstate: UmlStateDB = new (UmlStateDB)

	// front repo
	frontRepo: FrontRepo = new (FrontRepo)
 
	constructor(
		private umlstateService: UmlStateService,
		private frontRepoService: FrontRepoService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getUmlState();

		// observable for changes in 
		this.umlstateService.UmlStateServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getUmlState()
				}
			}
		)
	}

	getUmlState(): void {
		const id = +this.route.snapshot.paramMap.get('id')!
		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				this.umlstate = this.frontRepo.UmlStates.get(id)!

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
				github_com_fullstack_lang_gongdoc_go_editor: ["github_com_fullstack_lang_gongdoc_go-" + "umlstate-detail", ID]
			}
		}]);
	}
}
