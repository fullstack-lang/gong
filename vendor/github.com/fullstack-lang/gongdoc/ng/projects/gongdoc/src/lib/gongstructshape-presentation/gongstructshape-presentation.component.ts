import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';

import { GongStructShapeDB } from '../gongstructshape-db'
import { GongStructShapeService } from '../gongstructshape.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

// insertion point for additional imports

export interface gongstructshapeDummyElement {
}

const ELEMENT_DATA: gongstructshapeDummyElement[] = [
];

@Component({
	selector: 'app-gongstructshape-presentation',
	templateUrl: './gongstructshape-presentation.component.html',
	styleUrls: ['./gongstructshape-presentation.component.css'],
})
export class GongStructShapePresentationComponent implements OnInit {

	// insertion point for additionnal time duration declarations
	// insertion point for additionnal enum int field declarations

	displayedColumns: string[] = []
	dataSource = ELEMENT_DATA

	gongstructshape: GongStructShapeDB = new (GongStructShapeDB)

	// front repo
	frontRepo: FrontRepo = new (FrontRepo)
 
	constructor(
		private gongstructshapeService: GongStructShapeService,
		private frontRepoService: FrontRepoService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getGongStructShape();

		// observable for changes in 
		this.gongstructshapeService.GongStructShapeServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getGongStructShape()
				}
			}
		)
	}

	getGongStructShape(): void {
		const id = +this.route.snapshot.paramMap.get('id')!
		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				this.gongstructshape = this.frontRepo.GongStructShapes.get(id)!

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
				github_com_fullstack_lang_gongdoc_go_editor: ["github_com_fullstack_lang_gongdoc_go-" + "gongstructshape-detail", ID]
			}
		}]);
	}
}
