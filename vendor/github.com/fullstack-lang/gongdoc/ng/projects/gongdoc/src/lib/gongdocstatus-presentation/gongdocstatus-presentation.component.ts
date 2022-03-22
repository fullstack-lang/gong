import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';

import { GongdocStatusDB } from '../gongdocstatus-db'
import { GongdocStatusService } from '../gongdocstatus.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

// insertion point for additional imports

export interface gongdocstatusDummyElement {
}

const ELEMENT_DATA: gongdocstatusDummyElement[] = [
];

@Component({
	selector: 'app-gongdocstatus-presentation',
	templateUrl: './gongdocstatus-presentation.component.html',
	styleUrls: ['./gongdocstatus-presentation.component.css'],
})
export class GongdocStatusPresentationComponent implements OnInit {

	// insertion point for additionnal time duration declarations
	// insertion point for additionnal enum int field declarations

	displayedColumns: string[] = []
	dataSource = ELEMENT_DATA

	gongdocstatus: GongdocStatusDB = new (GongdocStatusDB)

	// front repo
	frontRepo: FrontRepo = new (FrontRepo)
 
	constructor(
		private gongdocstatusService: GongdocStatusService,
		private frontRepoService: FrontRepoService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getGongdocStatus();

		// observable for changes in 
		this.gongdocstatusService.GongdocStatusServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getGongdocStatus()
				}
			}
		)
	}

	getGongdocStatus(): void {
		const id = +this.route.snapshot.paramMap.get('id')!
		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				this.gongdocstatus = this.frontRepo.GongdocStatuss.get(id)!

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
				github_com_fullstack_lang_gongdoc_go_editor: ["github_com_fullstack_lang_gongdoc_go-" + "gongdocstatus-detail", ID]
			}
		}]);
	}
}
