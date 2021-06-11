import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';

import { BclassDB } from '../bclass-db'
import { BclassService } from '../bclass.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

export interface bclassDummyElement {
}

const ELEMENT_DATA: bclassDummyElement[] = [
];

@Component({
	selector: 'app-bclass-presentation',
	templateUrl: './bclass-presentation.component.html',
	styleUrls: ['./bclass-presentation.component.css'],
})
export class BclassPresentationComponent implements OnInit {

	// insertion point for declarations

	displayedColumns: string[] = [];
	dataSource = ELEMENT_DATA;

	bclass: BclassDB;

	// front repo
	frontRepo: FrontRepo
 
	constructor(
		private bclassService: BclassService,
		private frontRepoService: FrontRepoService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getBclass();

		// observable for changes in 
		this.bclassService.BclassServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getBclass()
				}
			}
		)
	}

	getBclass(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				this.bclass = this.frontRepo.Bclasss.get(id)

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
				github_com_fullstack_lang_gong_test_go_editor: ["github_com_fullstack_lang_gong_test_go-" + "bclass-detail", ID]
			}
		}]);
	}
}
