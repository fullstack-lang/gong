import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';

import { GongStructDB } from '../gongstruct-db'
import { GongStructService } from '../gongstruct.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

// insertion point for additional imports

export interface gongstructDummyElement {
}

const ELEMENT_DATA: gongstructDummyElement[] = [
];

@Component({
	selector: 'app-gongstruct-presentation',
	templateUrl: './gongstruct-presentation.component.html',
	styleUrls: ['./gongstruct-presentation.component.css'],
})
export class GongStructPresentationComponent implements OnInit {

	// insertion point for additionnal time duration declarations
	// insertion point for additionnal enum int field declarations

	displayedColumns: string[] = []
	dataSource = ELEMENT_DATA

	gongstruct: GongStructDB = new (GongStructDB)

	// front repo
	frontRepo: FrontRepo = new (FrontRepo)
 
	constructor(
		private gongstructService: GongStructService,
		private frontRepoService: FrontRepoService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getGongStruct();

		// observable for changes in 
		this.gongstructService.GongStructServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getGongStruct()
				}
			}
		)
	}

	getGongStruct(): void {
		const id = +this.route.snapshot.paramMap.get('id')!
		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				this.gongstruct = this.frontRepo.GongStructs.get(id)!

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
				github_com_fullstack_lang_gong_go_editor: ["github_com_fullstack_lang_gong_go-" + "gongstruct-detail", ID]
			}
		}]);
	}
}
