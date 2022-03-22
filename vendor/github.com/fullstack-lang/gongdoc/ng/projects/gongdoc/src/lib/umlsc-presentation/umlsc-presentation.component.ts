import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';

import { UmlscDB } from '../umlsc-db'
import { UmlscService } from '../umlsc.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

// insertion point for additional imports

export interface umlscDummyElement {
}

const ELEMENT_DATA: umlscDummyElement[] = [
];

@Component({
	selector: 'app-umlsc-presentation',
	templateUrl: './umlsc-presentation.component.html',
	styleUrls: ['./umlsc-presentation.component.css'],
})
export class UmlscPresentationComponent implements OnInit {

	// insertion point for additionnal time duration declarations
	// insertion point for additionnal enum int field declarations

	displayedColumns: string[] = []
	dataSource = ELEMENT_DATA

	umlsc: UmlscDB = new (UmlscDB)

	// front repo
	frontRepo: FrontRepo = new (FrontRepo)
 
	constructor(
		private umlscService: UmlscService,
		private frontRepoService: FrontRepoService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getUmlsc();

		// observable for changes in 
		this.umlscService.UmlscServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getUmlsc()
				}
			}
		)
	}

	getUmlsc(): void {
		const id = +this.route.snapshot.paramMap.get('id')!
		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				this.umlsc = this.frontRepo.Umlscs.get(id)!

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
				github_com_fullstack_lang_gongdoc_go_editor: ["github_com_fullstack_lang_gongdoc_go-" + "umlsc-detail", ID]
			}
		}]);
	}
}
