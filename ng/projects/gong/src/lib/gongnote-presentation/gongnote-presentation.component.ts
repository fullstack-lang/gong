import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';

import { GongNoteDB } from '../gongnote-db'
import { GongNoteService } from '../gongnote.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

// insertion point for additional imports

export interface gongnoteDummyElement {
}

const ELEMENT_DATA: gongnoteDummyElement[] = [
];

@Component({
	selector: 'app-gongnote-presentation',
	templateUrl: './gongnote-presentation.component.html',
	styleUrls: ['./gongnote-presentation.component.css'],
})
export class GongNotePresentationComponent implements OnInit {

	// insertion point for additionnal time duration declarations
	// insertion point for additionnal enum int field declarations

	displayedColumns: string[] = []
	dataSource = ELEMENT_DATA

	gongnote: GongNoteDB = new (GongNoteDB)

	// front repo
	frontRepo: FrontRepo = new (FrontRepo)
 
	constructor(
		private gongnoteService: GongNoteService,
		private frontRepoService: FrontRepoService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getGongNote();

		// observable for changes in 
		this.gongnoteService.GongNoteServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getGongNote()
				}
			}
		)
	}

	getGongNote(): void {
		const id = +this.route.snapshot.paramMap.get('id')!
		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				this.gongnote = this.frontRepo.GongNotes.get(id)!

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
				github_com_fullstack_lang_gong_go_editor: ["github_com_fullstack_lang_gong_go-" + "gongnote-detail", ID]
			}
		}]);
	}
}
