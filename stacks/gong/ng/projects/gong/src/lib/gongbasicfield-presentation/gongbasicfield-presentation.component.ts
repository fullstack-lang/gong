import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup} from '@angular/forms';

import { GongBasicFieldDB } from '../gongbasicfield-db'
import { GongBasicFieldService } from '../gongbasicfield.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

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

	// insertion point for declarations

	displayedColumns: string[] = [];
	dataSource = ELEMENT_DATA;

	gongbasicfield: GongBasicFieldDB;
 
	constructor(
		private gongbasicfieldService: GongBasicFieldService,
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
		const id = +this.route.snapshot.paramMap.get('id');
		this.gongbasicfieldService.getGongBasicField(id)
			.subscribe(
				gongbasicfield => {
					this.gongbasicfield = gongbasicfield

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
				editor: ["gongbasicfield-detail", ID]
			}
		}]);
	}
}
