import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';

import { ModelPkgDB } from '../modelpkg-db'
import { ModelPkgService } from '../modelpkg.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

// insertion point for additional imports

export interface modelpkgDummyElement {
}

const ELEMENT_DATA: modelpkgDummyElement[] = [
];

@Component({
	selector: 'app-modelpkg-presentation',
	templateUrl: './modelpkg-presentation.component.html',
	styleUrls: ['./modelpkg-presentation.component.css'],
})
export class ModelPkgPresentationComponent implements OnInit {

	// insertion point for additionnal time duration declarations
	// insertion point for additionnal enum int field declarations

	displayedColumns: string[] = []
	dataSource = ELEMENT_DATA

	modelpkg: ModelPkgDB = new (ModelPkgDB)

	// front repo
	frontRepo: FrontRepo = new (FrontRepo)
 
	constructor(
		private modelpkgService: ModelPkgService,
		private frontRepoService: FrontRepoService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getModelPkg();

		// observable for changes in 
		this.modelpkgService.ModelPkgServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getModelPkg()
				}
			}
		)
	}

	getModelPkg(): void {
		const id = +this.route.snapshot.paramMap.get('id')!
		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				this.modelpkg = this.frontRepo.ModelPkgs.get(id)!

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
				github_com_fullstack_lang_gong_go_editor: ["github_com_fullstack_lang_gong_go-" + "modelpkg-detail", ID]
			}
		}]);
	}
}
