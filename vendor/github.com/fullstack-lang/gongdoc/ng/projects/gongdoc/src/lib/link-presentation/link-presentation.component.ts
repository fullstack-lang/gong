import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';

import { LinkDB } from '../link-db'
import { LinkService } from '../link.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

// insertion point for additional imports

export interface linkDummyElement {
}

const ELEMENT_DATA: linkDummyElement[] = [
];

@Component({
	selector: 'app-link-presentation',
	templateUrl: './link-presentation.component.html',
	styleUrls: ['./link-presentation.component.css'],
})
export class LinkPresentationComponent implements OnInit {

	// insertion point for additionnal time duration declarations
	// insertion point for additionnal enum int field declarations

	displayedColumns: string[] = []
	dataSource = ELEMENT_DATA

	link: LinkDB = new (LinkDB)

	// front repo
	frontRepo: FrontRepo = new (FrontRepo)
 
	constructor(
		private linkService: LinkService,
		private frontRepoService: FrontRepoService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getLink();

		// observable for changes in 
		this.linkService.LinkServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getLink()
				}
			}
		)
	}

	getLink(): void {
		const id = +this.route.snapshot.paramMap.get('id')!
		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				this.link = this.frontRepo.Links.get(id)!

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
				github_com_fullstack_lang_gongdoc_go_editor: ["github_com_fullstack_lang_gongdoc_go-" + "link-detail", ID]
			}
		}]);
	}
}
