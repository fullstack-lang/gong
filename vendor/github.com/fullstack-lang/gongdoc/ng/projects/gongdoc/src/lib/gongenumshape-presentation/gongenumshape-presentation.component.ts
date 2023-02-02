import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';

import { GongEnumShapeDB } from '../gongenumshape-db'
import { GongEnumShapeService } from '../gongenumshape.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

// insertion point for additional imports

export interface gongenumshapeDummyElement {
}

const ELEMENT_DATA: gongenumshapeDummyElement[] = [
];

@Component({
	selector: 'app-gongenumshape-presentation',
	templateUrl: './gongenumshape-presentation.component.html',
	styleUrls: ['./gongenumshape-presentation.component.css'],
})
export class GongEnumShapePresentationComponent implements OnInit {

	// insertion point for additionnal time duration declarations
	// insertion point for additionnal enum int field declarations

	displayedColumns: string[] = []
	dataSource = ELEMENT_DATA

	gongenumshape: GongEnumShapeDB = new (GongEnumShapeDB)

	// front repo
	frontRepo: FrontRepo = new (FrontRepo)
 
	constructor(
		private gongenumshapeService: GongEnumShapeService,
		private frontRepoService: FrontRepoService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getGongEnumShape();

		// observable for changes in 
		this.gongenumshapeService.GongEnumShapeServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getGongEnumShape()
				}
			}
		)
	}

	getGongEnumShape(): void {
		const id = +this.route.snapshot.paramMap.get('id')!
		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				this.gongenumshape = this.frontRepo.GongEnumShapes.get(id)!

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
				github_com_fullstack_lang_gongdoc_go_editor: ["github_com_fullstack_lang_gongdoc_go-" + "gongenumshape-detail", ID]
			}
		}]);
	}
}
