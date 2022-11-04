import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';

import { NoteShapeDB } from '../noteshape-db'
import { NoteShapeService } from '../noteshape.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

// insertion point for additional imports

export interface noteshapeDummyElement {
}

const ELEMENT_DATA: noteshapeDummyElement[] = [
];

@Component({
	selector: 'app-noteshape-presentation',
	templateUrl: './noteshape-presentation.component.html',
	styleUrls: ['./noteshape-presentation.component.css'],
})
export class NoteShapePresentationComponent implements OnInit {

	// insertion point for additionnal time duration declarations
	// insertion point for additionnal enum int field declarations

	displayedColumns: string[] = []
	dataSource = ELEMENT_DATA

	noteshape: NoteShapeDB = new (NoteShapeDB)

	// front repo
	frontRepo: FrontRepo = new (FrontRepo)
 
	constructor(
		private noteshapeService: NoteShapeService,
		private frontRepoService: FrontRepoService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getNoteShape();

		// observable for changes in 
		this.noteshapeService.NoteShapeServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getNoteShape()
				}
			}
		)
	}

	getNoteShape(): void {
		const id = +this.route.snapshot.paramMap.get('id')!
		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				this.noteshape = this.frontRepo.NoteShapes.get(id)!

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
				github_com_fullstack_lang_gongdoc_go_editor: ["github_com_fullstack_lang_gongdoc_go-" + "noteshape-detail", ID]
			}
		}]);
	}
}
