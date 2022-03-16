import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';

import { GongEnumValueDB } from '../gongenumvalue-db'
import { GongEnumValueService } from '../gongenumvalue.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

// insertion point for additional imports

export interface gongenumvalueDummyElement {
}

const ELEMENT_DATA: gongenumvalueDummyElement[] = [
];

@Component({
	selector: 'app-gongenumvalue-presentation',
	templateUrl: './gongenumvalue-presentation.component.html',
	styleUrls: ['./gongenumvalue-presentation.component.css'],
})
export class GongEnumValuePresentationComponent implements OnInit {

	// insertion point for additionnal time duration declarations
	// insertion point for additionnal enum int field declarations

	displayedColumns: string[] = []
	dataSource = ELEMENT_DATA

	gongenumvalue: GongEnumValueDB = new (GongEnumValueDB)

	// front repo
	frontRepo: FrontRepo = new (FrontRepo)
 
	constructor(
		private gongenumvalueService: GongEnumValueService,
		private frontRepoService: FrontRepoService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getGongEnumValue();

		// observable for changes in 
		this.gongenumvalueService.GongEnumValueServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getGongEnumValue()
				}
			}
		)
	}

	getGongEnumValue(): void {
		const id = +this.route.snapshot.paramMap.get('id')!
		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				this.gongenumvalue = this.frontRepo.GongEnumValues.get(id)!

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
				github_com_fullstack_lang_gong_go_editor: ["github_com_fullstack_lang_gong_go-" + "gongenumvalue-detail", ID]
			}
		}]);
	}
}
