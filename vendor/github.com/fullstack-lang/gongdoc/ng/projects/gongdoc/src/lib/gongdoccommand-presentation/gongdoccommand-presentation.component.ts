import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';

import { GongdocCommandDB } from '../gongdoccommand-db'
import { GongdocCommandService } from '../gongdoccommand.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

// insertion point for additional imports

export interface gongdoccommandDummyElement {
}

const ELEMENT_DATA: gongdoccommandDummyElement[] = [
];

@Component({
	selector: 'app-gongdoccommand-presentation',
	templateUrl: './gongdoccommand-presentation.component.html',
	styleUrls: ['./gongdoccommand-presentation.component.css'],
})
export class GongdocCommandPresentationComponent implements OnInit {

	// insertion point for additionnal time duration declarations
	// insertion point for additionnal enum int field declarations

	displayedColumns: string[] = []
	dataSource = ELEMENT_DATA

	gongdoccommand: GongdocCommandDB = new (GongdocCommandDB)

	// front repo
	frontRepo: FrontRepo = new (FrontRepo)
 
	constructor(
		private gongdoccommandService: GongdocCommandService,
		private frontRepoService: FrontRepoService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getGongdocCommand();

		// observable for changes in 
		this.gongdoccommandService.GongdocCommandServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getGongdocCommand()
				}
			}
		)
	}

	getGongdocCommand(): void {
		const id = +this.route.snapshot.paramMap.get('id')!
		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				this.gongdoccommand = this.frontRepo.GongdocCommands.get(id)!

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
				github_com_fullstack_lang_gongdoc_go_editor: ["github_com_fullstack_lang_gongdoc_go-" + "gongdoccommand-detail", ID]
			}
		}]);
	}
}
