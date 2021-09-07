import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';

import { BstructDB } from '../bstruct-db'
import { BstructService } from '../bstruct.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

export interface bstructDummyElement {
}

const ELEMENT_DATA: bstructDummyElement[] = [
];

@Component({
	selector: 'app-bstruct-presentation',
	templateUrl: './bstruct-presentation.component.html',
	styleUrls: ['./bstruct-presentation.component.css'],
})
export class BstructPresentationComponent implements OnInit {

	// insertion point for declarations

	displayedColumns: string[] = [];
	dataSource = ELEMENT_DATA;

	bstruct: BstructDB;

	// front repo
	frontRepo: FrontRepo
 
	constructor(
		private bstructService: BstructService,
		private frontRepoService: FrontRepoService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getBstruct();

		// observable for changes in 
		this.bstructService.BstructServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getBstruct()
				}
			}
		)
	}

	getBstruct(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				this.bstruct = this.frontRepo.Bstructs.get(id)

				// insertion point for recovery of durations
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
				github_com_fullstack_lang_gong_test_go_editor: ["github_com_fullstack_lang_gong_test_go-" + "bstruct-detail", ID]
			}
		}]);
	}
}
