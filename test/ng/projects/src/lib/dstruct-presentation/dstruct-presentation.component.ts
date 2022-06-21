import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';

import { DstructDB } from '../dstruct-db'
import { DstructService } from '../dstruct.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

// insertion point for additional imports

export interface dstructDummyElement {
}

const ELEMENT_DATA: dstructDummyElement[] = [
];

@Component({
	selector: 'app-dstruct-presentation',
	templateUrl: './dstruct-presentation.component.html',
	styleUrls: ['./dstruct-presentation.component.css'],
})
export class DstructPresentationComponent implements OnInit {

	// insertion point for additionnal time duration declarations
	// insertion point for additionnal enum int field declarations

	displayedColumns: string[] = []
	dataSource = ELEMENT_DATA

	dstruct: DstructDB = new (DstructDB)

	// front repo
	frontRepo: FrontRepo = new (FrontRepo)
 
	constructor(
		private dstructService: DstructService,
		private frontRepoService: FrontRepoService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getDstruct();

		// observable for changes in 
		this.dstructService.DstructServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getDstruct()
				}
			}
		)
	}

	getDstruct(): void {
		const id = +this.route.snapshot.paramMap.get('id')!
		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				this.dstruct = this.frontRepo.Dstructs.get(id)!

				// insertion point for recovery of durations
				// insertion point for recovery of enum tint
			}
		);
	}

	// set presentation outlet
	setPresentationRouterOutlet(structName: string, ID: number) {
		this.router.navigate([{
			outlets: {
				github_com_fullstack_lang_gong_test_go_presentation: ["github_com_fullstack_lang_gong_test_go-" + structName + "-presentation", ID]
			}
		}]);
	}

	// set editor outlet
	setEditorRouterOutlet(ID: number) {
		this.router.navigate([{
			outlets: {
				github_com_fullstack_lang_gong_test_go_editor: ["github_com_fullstack_lang_gong_test_go-" + "dstruct-detail", ID]
			}
		}]);
	}
}
