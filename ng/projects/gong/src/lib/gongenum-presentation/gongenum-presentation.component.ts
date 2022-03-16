import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';

import { GongEnumDB } from '../gongenum-db'
import { GongEnumService } from '../gongenum.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

// insertion point for additional imports
import { GongEnumTypeList } from '../GongEnumType'

export interface gongenumDummyElement {
}

const ELEMENT_DATA: gongenumDummyElement[] = [
];

@Component({
	selector: 'app-gongenum-presentation',
	templateUrl: './gongenum-presentation.component.html',
	styleUrls: ['./gongenum-presentation.component.css'],
})
export class GongEnumPresentationComponent implements OnInit {

	// insertion point for additionnal time duration declarations
	// insertion point for additionnal enum int field declarations
	Type_Value : string = ""

	displayedColumns: string[] = []
	dataSource = ELEMENT_DATA

	gongenum: GongEnumDB = new (GongEnumDB)

	// front repo
	frontRepo: FrontRepo = new (FrontRepo)
 
	constructor(
		private gongenumService: GongEnumService,
		private frontRepoService: FrontRepoService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getGongEnum();

		// observable for changes in 
		this.gongenumService.GongEnumServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getGongEnum()
				}
			}
		)
	}

	getGongEnum(): void {
		const id = +this.route.snapshot.paramMap.get('id')!
		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				this.gongenum = this.frontRepo.GongEnums.get(id)!

				// insertion point for recovery of durations
				// insertion point for recovery of enum tint
				this.Type_Value = GongEnumTypeList[this.gongenum.Type].viewValue
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
				github_com_fullstack_lang_gong_go_editor: ["github_com_fullstack_lang_gong_go-" + "gongenum-detail", ID]
			}
		}]);
	}
}
