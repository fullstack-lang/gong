import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup} from '@angular/forms';

import { GongTimeFieldDB } from '../gongtimefield-db'
import { GongTimeFieldService } from '../gongtimefield.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

export interface gongtimefieldDummyElement {
}

const ELEMENT_DATA: gongtimefieldDummyElement[] = [
];

@Component({
	selector: 'app-gongtimefield-presentation',
	templateUrl: './gongtimefield-presentation.component.html',
	styleUrls: ['./gongtimefield-presentation.component.css'],
})
export class GongTimeFieldPresentationComponent implements OnInit {

	// insertion point for declarations

	displayedColumns: string[] = [];
	dataSource = ELEMENT_DATA;

	gongtimefield: GongTimeFieldDB;
 
	constructor(
		private gongtimefieldService: GongTimeFieldService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getGongTimeField();

		// observable for changes in 
		this.gongtimefieldService.GongTimeFieldServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getGongTimeField()
				}
			}
		)
	}

	getGongTimeField(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		this.gongtimefieldService.getGongTimeField(id)
			.subscribe(
				gongtimefield => {
					this.gongtimefield = gongtimefield

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
				editor: ["gongtimefield-detail", ID]
			}
		}]);
	}
}
