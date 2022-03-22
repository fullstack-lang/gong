import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';

import { ClassdiagramDB } from '../classdiagram-db'
import { ClassdiagramService } from '../classdiagram.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

// insertion point for additional imports

export interface classdiagramDummyElement {
}

const ELEMENT_DATA: classdiagramDummyElement[] = [
];

@Component({
	selector: 'app-classdiagram-presentation',
	templateUrl: './classdiagram-presentation.component.html',
	styleUrls: ['./classdiagram-presentation.component.css'],
})
export class ClassdiagramPresentationComponent implements OnInit {

	// insertion point for additionnal time duration declarations
	// insertion point for additionnal enum int field declarations

	displayedColumns: string[] = []
	dataSource = ELEMENT_DATA

	classdiagram: ClassdiagramDB = new (ClassdiagramDB)

	// front repo
	frontRepo: FrontRepo = new (FrontRepo)
 
	constructor(
		private classdiagramService: ClassdiagramService,
		private frontRepoService: FrontRepoService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getClassdiagram();

		// observable for changes in 
		this.classdiagramService.ClassdiagramServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getClassdiagram()
				}
			}
		)
	}

	getClassdiagram(): void {
		const id = +this.route.snapshot.paramMap.get('id')!
		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				this.classdiagram = this.frontRepo.Classdiagrams.get(id)!

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
				github_com_fullstack_lang_gongdoc_go_editor: ["github_com_fullstack_lang_gongdoc_go-" + "classdiagram-detail", ID]
			}
		}]);
	}
}
