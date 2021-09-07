import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';

import { AstructBstruct2UseDB } from '../astructbstruct2use-db'
import { AstructBstruct2UseService } from '../astructbstruct2use.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

export interface astructbstruct2useDummyElement {
}

const ELEMENT_DATA: astructbstruct2useDummyElement[] = [
];

@Component({
	selector: 'app-astructbstruct2use-presentation',
	templateUrl: './astructbstruct2use-presentation.component.html',
	styleUrls: ['./astructbstruct2use-presentation.component.css'],
})
export class AstructBstruct2UsePresentationComponent implements OnInit {

	// insertion point for declarations

	displayedColumns: string[] = [];
	dataSource = ELEMENT_DATA;

	astructbstruct2use: AstructBstruct2UseDB;

	// front repo
	frontRepo: FrontRepo
 
	constructor(
		private astructbstruct2useService: AstructBstruct2UseService,
		private frontRepoService: FrontRepoService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getAstructBstruct2Use();

		// observable for changes in 
		this.astructbstruct2useService.AstructBstruct2UseServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getAstructBstruct2Use()
				}
			}
		)
	}

	getAstructBstruct2Use(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				this.astructbstruct2use = this.frontRepo.AstructBstruct2Uses.get(id)

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
				github_com_fullstack_lang_gong_test_go_editor: ["github_com_fullstack_lang_gong_test_go-" + "astructbstruct2use-detail", ID]
			}
		}]);
	}
}
