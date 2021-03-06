import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';

import { AstructDB } from '../astruct-db'
import { AstructService } from '../astruct.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

// insertion point for additional imports
import { CEnumTypeIntList } from '../CEnumTypeInt'

export interface astructDummyElement {
}

const ELEMENT_DATA: astructDummyElement[] = [
];

@Component({
	selector: 'app-astruct-presentation',
	templateUrl: './astruct-presentation.component.html',
	styleUrls: ['./astruct-presentation.component.css'],
})
export class AstructPresentationComponent implements OnInit {

	// insertion point for additionnal time duration declarations
	// fields from Duration1
	Duration1_Hours: number = 0
	Duration1_Minutes: number = 0
	Duration1_Seconds: number = 0
	// insertion point for additionnal enum int field declarations
	CEnum_Value : string = ""

	displayedColumns: string[] = []
	dataSource = ELEMENT_DATA

	astruct: AstructDB = new (AstructDB)

	// front repo
	frontRepo: FrontRepo = new (FrontRepo)
 
	constructor(
		private astructService: AstructService,
		private frontRepoService: FrontRepoService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getAstruct();

		// observable for changes in 
		this.astructService.AstructServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getAstruct()
				}
			}
		)
	}

	getAstruct(): void {
		const id = +this.route.snapshot.paramMap.get('id')!
		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				this.astruct = this.frontRepo.Astructs.get(id)!

				// insertion point for recovery of durations
				// computation of Hours, Minutes, Seconds for Duration1
				this.Duration1_Hours = Math.floor(this.astruct.Duration1 / (3600 * 1000 * 1000 * 1000))
				this.Duration1_Minutes = Math.floor(this.astruct.Duration1 % (3600 * 1000 * 1000 * 1000) / (60 * 1000 * 1000 * 1000))
				this.Duration1_Seconds = this.astruct.Duration1 % (60 * 1000 * 1000 * 1000) / (1000 * 1000 * 1000)
				// insertion point for recovery of enum tint
				this.CEnum_Value = CEnumTypeIntList[this.astruct.CEnum].viewValue
			}
		);
	}

	// set presentation outlet
	setPresentationRouterOutlet(structName: string, ID: number) {
		this.router.navigate([{
			outlets: {
				github_com_fullstack_lang_gong_test_go_presentation: ["github_com_fullstack_lang_gong_test_go-" + structName + "-presentation", ID]
			}
		}]);
	}

	// set editor outlet
	setEditorRouterOutlet(ID: number) {
		this.router.navigate([{
			outlets: {
				github_com_fullstack_lang_gong_test_go_editor: ["github_com_fullstack_lang_gong_test_go-" + "astruct-detail", ID]
			}
		}]);
	}
}
