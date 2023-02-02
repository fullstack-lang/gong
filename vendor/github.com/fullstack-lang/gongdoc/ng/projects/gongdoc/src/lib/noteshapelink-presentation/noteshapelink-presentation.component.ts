import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';

import { NoteShapeLinkDB } from '../noteshapelink-db'
import { NoteShapeLinkService } from '../noteshapelink.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

// insertion point for additional imports

export interface noteshapelinkDummyElement {
}

const ELEMENT_DATA: noteshapelinkDummyElement[] = [
];

@Component({
	selector: 'app-noteshapelink-presentation',
	templateUrl: './noteshapelink-presentation.component.html',
	styleUrls: ['./noteshapelink-presentation.component.css'],
})
export class NoteShapeLinkPresentationComponent implements OnInit {

	// insertion point for additionnal time duration declarations
	// insertion point for additionnal enum int field declarations

	displayedColumns: string[] = []
	dataSource = ELEMENT_DATA

	noteshapelink: NoteShapeLinkDB = new (NoteShapeLinkDB)

	// front repo
	frontRepo: FrontRepo = new (FrontRepo)
 
	constructor(
		private noteshapelinkService: NoteShapeLinkService,
		private frontRepoService: FrontRepoService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getNoteShapeLink();

		// observable for changes in 
		this.noteshapelinkService.NoteShapeLinkServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getNoteShapeLink()
				}
			}
		)
	}

	getNoteShapeLink(): void {
		const id = +this.route.snapshot.paramMap.get('id')!
		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				this.noteshapelink = this.frontRepo.NoteShapeLinks.get(id)!

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
				github_com_fullstack_lang_gongdoc_go_editor: ["github_com_fullstack_lang_gongdoc_go-" + "noteshapelink-detail", ID]
			}
		}]);
	}
}
