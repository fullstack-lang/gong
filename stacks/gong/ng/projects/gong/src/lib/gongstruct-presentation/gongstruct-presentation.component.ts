import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup} from '@angular/forms';

import { GongStructDB } from '../gongstruct-db'
import { GongStructService } from '../gongstruct.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

export interface gongstructDummyElement {
}

const ELEMENT_DATA: gongstructDummyElement[] = [
];

@Component({
	selector: 'app-gongstruct-presentation',
	templateUrl: './gongstruct-presentation.component.html',
	styleUrls: ['./gongstruct-presentation.component.css'],
})
export class GongStructPresentationComponent implements OnInit {

	// insertion point for declarations

	displayedColumns: string[] = [];
	dataSource = ELEMENT_DATA;

	gongstruct: GongStructDB;
 
	constructor(
		private gongstructService: GongStructService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getGongStruct();

		// observable for changes in 
		this.gongstructService.GongStructServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getGongStruct()
				}
			}
		)
	}

	getGongStruct(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		this.gongstructService.getGongStruct(id)
			.subscribe(
				gongstruct => {
					this.gongstruct = gongstruct

					// insertion point for recovery of durations

				}
			);
	}

	// set presentation outlet
	setPresentationRouterOutlet(structName: string, ID: number) {
		this.router.navigate([{
			outlets: {
				presentation: [structName + "-presentation", ID]
			}
		}]);
	}

	// set editor outlet
	setEditorRouterOutlet(ID: number) {
		this.router.navigate([{
			outlets: {
				editor: ["gongstruct-detail", ID]
			}
		}]);
	}
}
