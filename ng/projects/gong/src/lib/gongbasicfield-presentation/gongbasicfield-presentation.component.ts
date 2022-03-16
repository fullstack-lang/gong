import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';

import { GongBasicFieldDB } from '../gongbasicfield-db'
import { GongBasicFieldService } from '../gongbasicfield.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

// insertion point for additional imports

export interface gongbasicfieldDummyElement {
}

const ELEMENT_DATA: gongbasicfieldDummyElement[] = [
];

@Component({
	selector: 'app-gongbasicfield-presentation',
	templateUrl: './gongbasicfield-presentation.component.html',
	styleUrls: ['./gongbasicfield-presentation.component.css'],
})
export class GongBasicFieldPresentationComponent implements OnInit {

	// insertion point for additionnal time duration declarations
	// insertion point for additionnal enum int field declarations

	displayedColumns: string[] = []
	dataSource = ELEMENT_DATA

	gongbasicfield: GongBasicFieldDB = new (GongBasicFieldDB)

	// front repo
	frontRepo: FrontRepo = new (FrontRepo)
 
	constructor(
		private gongbasicfieldService: GongBasicFieldService,
		private frontRepoService: FrontRepoService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getGongBasicField();

		// observable for changes in 
		this.gongbasicfieldService.GongBasicFieldServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getGongBasicField()
				}
			}
		)
	}

	getGongBasicField(): void {
		const id = +this.route.snapshot.paramMap.get('id')!
		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				this.gongbasicfield = this.frontRepo.GongBasicFields.get(id)!

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
				github_com_fullstack_lang_gong_go_editor: ["github_com_fullstack_lang_gong_go-" + "gongbasicfield-detail", ID]
			}
		}]);
	}
}
