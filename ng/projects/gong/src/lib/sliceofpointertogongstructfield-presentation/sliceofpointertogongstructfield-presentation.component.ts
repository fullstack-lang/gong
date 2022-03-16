import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';

import { SliceOfPointerToGongStructFieldDB } from '../sliceofpointertogongstructfield-db'
import { SliceOfPointerToGongStructFieldService } from '../sliceofpointertogongstructfield.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

// insertion point for additional imports

export interface sliceofpointertogongstructfieldDummyElement {
}

const ELEMENT_DATA: sliceofpointertogongstructfieldDummyElement[] = [
];

@Component({
	selector: 'app-sliceofpointertogongstructfield-presentation',
	templateUrl: './sliceofpointertogongstructfield-presentation.component.html',
	styleUrls: ['./sliceofpointertogongstructfield-presentation.component.css'],
})
export class SliceOfPointerToGongStructFieldPresentationComponent implements OnInit {

	// insertion point for additionnal time duration declarations
	// insertion point for additionnal enum int field declarations

	displayedColumns: string[] = []
	dataSource = ELEMENT_DATA

	sliceofpointertogongstructfield: SliceOfPointerToGongStructFieldDB = new (SliceOfPointerToGongStructFieldDB)

	// front repo
	frontRepo: FrontRepo = new (FrontRepo)
 
	constructor(
		private sliceofpointertogongstructfieldService: SliceOfPointerToGongStructFieldService,
		private frontRepoService: FrontRepoService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getSliceOfPointerToGongStructField();

		// observable for changes in 
		this.sliceofpointertogongstructfieldService.SliceOfPointerToGongStructFieldServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getSliceOfPointerToGongStructField()
				}
			}
		)
	}

	getSliceOfPointerToGongStructField(): void {
		const id = +this.route.snapshot.paramMap.get('id')!
		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				this.sliceofpointertogongstructfield = this.frontRepo.SliceOfPointerToGongStructFields.get(id)!

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
				github_com_fullstack_lang_gong_go_editor: ["github_com_fullstack_lang_gong_go-" + "sliceofpointertogongstructfield-detail", ID]
			}
		}]);
	}
}
