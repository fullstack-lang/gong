import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';

import { ClassshapeDB } from '../classshape-db'
import { ClassshapeService } from '../classshape.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

// insertion point for additional imports

export interface classshapeDummyElement {
}

const ELEMENT_DATA: classshapeDummyElement[] = [
];

@Component({
	selector: 'app-classshape-presentation',
	templateUrl: './classshape-presentation.component.html',
	styleUrls: ['./classshape-presentation.component.css'],
})
export class ClassshapePresentationComponent implements OnInit {

	// insertion point for additionnal time duration declarations
	// insertion point for additionnal enum int field declarations

	displayedColumns: string[] = []
	dataSource = ELEMENT_DATA

	classshape: ClassshapeDB = new (ClassshapeDB)

	// front repo
	frontRepo: FrontRepo = new (FrontRepo)
 
	constructor(
		private classshapeService: ClassshapeService,
		private frontRepoService: FrontRepoService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getClassshape();

		// observable for changes in 
		this.classshapeService.ClassshapeServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getClassshape()
				}
			}
		)
	}

	getClassshape(): void {
		const id = +this.route.snapshot.paramMap.get('id')!
		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				this.classshape = this.frontRepo.Classshapes.get(id)!

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
				github_com_fullstack_lang_gongdoc_go_editor: ["github_com_fullstack_lang_gongdoc_go-" + "classshape-detail", ID]
			}
		}]);
	}
}
