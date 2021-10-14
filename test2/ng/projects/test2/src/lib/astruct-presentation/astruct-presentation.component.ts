import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';

import { AstructDB } from '../astruct-db'
import { AstructService } from '../astruct.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

export interface astructDummyElement {
}

const ELEMENT_DATA: astructDummyElement[] = [
];

@Component({
	selector: 'app-astruct-presentation',
	templateUrl: './astruct-presentation.component.html',
	styleUrls: ['./astruct-presentation.component.css'],
})
export class AstructPresentationComponent implements OnInit {

	// insertion point for declarations

	displayedColumns: string[] = []
	dataSource = ELEMENT_DATA

	astruct: AstructDB = new (AstructDB)

	// front repo
	frontRepo: FrontRepo = new (FrontRepo)
 
	constructor(
		private astructService: AstructService,
		private frontRepoService: FrontRepoService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getAstruct();

		// observable for changes in 
		this.astructService.AstructServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getAstruct()
				}
			}
		)
	}

	getAstruct(): void {
		const id = +this.route.snapshot.paramMap.get('id')!
		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				this.astruct = this.frontRepo.Astructs.get(id)!

				// insertion point for recovery of durations
			}
		);
	}

	// set presentation outlet
	setPresentationRouterOutlet(structName: string, ID: number) {
		this.router.navigate([{
			outlets: {
				github_com_fullstack_lang_gong_test2_go_presentation: ["github_com_fullstack_lang_gong_test2_go-" + structName + "-presentation", ID]
			}
		}]);
	}

	// set editor outlet
	setEditorRouterOutlet(ID: number) {
		this.router.navigate([{
			outlets: {
				github_com_fullstack_lang_gong_test2_go_editor: ["github_com_fullstack_lang_gong_test2_go-" + "astruct-detail", ID]
			}
		}]);
	}
}
