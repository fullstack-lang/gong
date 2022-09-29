import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';

import { NoteDB } from '../note-db'
import { NoteService } from '../note.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

// insertion point for additional imports

export interface noteDummyElement {
}

const ELEMENT_DATA: noteDummyElement[] = [
];

@Component({
	selector: 'app-note-presentation',
	templateUrl: './note-presentation.component.html',
	styleUrls: ['./note-presentation.component.css'],
})
export class NotePresentationComponent implements OnInit {

	// insertion point for additionnal time duration declarations
	// insertion point for additionnal enum int field declarations

	displayedColumns: string[] = []
	dataSource = ELEMENT_DATA

	note: NoteDB = new (NoteDB)

	// front repo
	frontRepo: FrontRepo = new (FrontRepo)
 
	constructor(
		private noteService: NoteService,
		private frontRepoService: FrontRepoService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getNote();

		// observable for changes in 
		this.noteService.NoteServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getNote()
				}
			}
		)
	}

	getNote(): void {
		const id = +this.route.snapshot.paramMap.get('id')!
		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				this.note = this.frontRepo.Notes.get(id)!

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
				github_com_fullstack_lang_gongdoc_go_editor: ["github_com_fullstack_lang_gongdoc_go-" + "note-detail", ID]
			}
		}]);
	}
}
