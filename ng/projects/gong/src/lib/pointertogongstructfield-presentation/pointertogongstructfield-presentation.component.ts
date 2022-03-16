import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';

import { PointerToGongStructFieldDB } from '../pointertogongstructfield-db'
import { PointerToGongStructFieldService } from '../pointertogongstructfield.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

// insertion point for additional imports

export interface pointertogongstructfieldDummyElement {
}

const ELEMENT_DATA: pointertogongstructfieldDummyElement[] = [
];

@Component({
	selector: 'app-pointertogongstructfield-presentation',
	templateUrl: './pointertogongstructfield-presentation.component.html',
	styleUrls: ['./pointertogongstructfield-presentation.component.css'],
})
export class PointerToGongStructFieldPresentationComponent implements OnInit {

	// insertion point for additionnal time duration declarations
	// insertion point for additionnal enum int field declarations

	displayedColumns: string[] = []
	dataSource = ELEMENT_DATA

	pointertogongstructfield: PointerToGongStructFieldDB = new (PointerToGongStructFieldDB)

	// front repo
	frontRepo: FrontRepo = new (FrontRepo)
 
	constructor(
		private pointertogongstructfieldService: PointerToGongStructFieldService,
		private frontRepoService: FrontRepoService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getPointerToGongStructField();

		// observable for changes in 
		this.pointertogongstructfieldService.PointerToGongStructFieldServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getPointerToGongStructField()
				}
			}
		)
	}

	getPointerToGongStructField(): void {
		const id = +this.route.snapshot.paramMap.get('id')!
		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				this.pointertogongstructfield = this.frontRepo.PointerToGongStructFields.get(id)!

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
				github_com_fullstack_lang_gong_go_editor: ["github_com_fullstack_lang_gong_go-" + "pointertogongstructfield-detail", ID]
			}
		}]);
	}
}
