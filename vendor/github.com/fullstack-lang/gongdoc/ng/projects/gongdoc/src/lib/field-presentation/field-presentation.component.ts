import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';

import { FieldDB } from '../field-db'
import { FieldService } from '../field.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

// insertion point for additional imports

export interface fieldDummyElement {
}

const ELEMENT_DATA: fieldDummyElement[] = [
];

@Component({
	selector: 'app-field-presentation',
	templateUrl: './field-presentation.component.html',
	styleUrls: ['./field-presentation.component.css'],
})
export class FieldPresentationComponent implements OnInit {

	// insertion point for additionnal time duration declarations
	// insertion point for additionnal enum int field declarations

	displayedColumns: string[] = []
	dataSource = ELEMENT_DATA

	field: FieldDB = new (FieldDB)

	// front repo
	frontRepo: FrontRepo = new (FrontRepo)
 
	constructor(
		private fieldService: FieldService,
		private frontRepoService: FrontRepoService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getField();

		// observable for changes in 
		this.fieldService.FieldServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getField()
				}
			}
		)
	}

	getField(): void {
		const id = +this.route.snapshot.paramMap.get('id')!
		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				this.field = this.frontRepo.Fields.get(id)!

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
				github_com_fullstack_lang_gongdoc_go_editor: ["github_com_fullstack_lang_gongdoc_go-" + "field-detail", ID]
			}
		}]);
	}
}
