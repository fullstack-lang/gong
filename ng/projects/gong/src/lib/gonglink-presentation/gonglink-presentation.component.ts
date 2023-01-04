import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';

import { GongLinkDB } from '../gonglink-db'
import { GongLinkService } from '../gonglink.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

// insertion point for additional imports

export interface gonglinkDummyElement {
}

const ELEMENT_DATA: gonglinkDummyElement[] = [
];

@Component({
	selector: 'app-gonglink-presentation',
	templateUrl: './gonglink-presentation.component.html',
	styleUrls: ['./gonglink-presentation.component.css'],
})
export class GongLinkPresentationComponent implements OnInit {

	// insertion point for additionnal time duration declarations
	// insertion point for additionnal enum int field declarations

	displayedColumns: string[] = []
	dataSource = ELEMENT_DATA

	gonglink: GongLinkDB = new (GongLinkDB)

	// front repo
	frontRepo: FrontRepo = new (FrontRepo)
 
	constructor(
		private gonglinkService: GongLinkService,
		private frontRepoService: FrontRepoService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getGongLink();

		// observable for changes in 
		this.gonglinkService.GongLinkServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getGongLink()
				}
			}
		)
	}

	getGongLink(): void {
		const id = +this.route.snapshot.paramMap.get('id')!
		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				this.gonglink = this.frontRepo.GongLinks.get(id)!

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
				github_com_fullstack_lang_gong_go_editor: ["github_com_fullstack_lang_gong_go-" + "gonglink-detail", ID]
			}
		}]);
	}
}
