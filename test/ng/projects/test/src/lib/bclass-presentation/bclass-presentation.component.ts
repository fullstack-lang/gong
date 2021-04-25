import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup} from '@angular/forms';

import { BclassDB } from '../bclass-db'
import { BclassService } from '../bclass.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

export interface bclassDummyElement {
}

const ELEMENT_DATA: bclassDummyElement[] = [
];

@Component({
	selector: 'app-bclass-presentation',
	templateUrl: './bclass-presentation.component.html',
	styleUrls: ['./bclass-presentation.component.css'],
})
export class BclassPresentationComponent implements OnInit {

	// insertion point for declarations

	displayedColumns: string[] = [];
	dataSource = ELEMENT_DATA;

	bclass: BclassDB;
 
	constructor(
		private bclassService: BclassService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getBclass();

		// observable for changes in 
		this.bclassService.BclassServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getBclass()
				}
			}
		)
	}

	getBclass(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		this.bclassService.getBclass(id)
			.subscribe(
				bclass => {
					this.bclass = bclass

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
				editor: ["bclass-detail", ID]
			}
		}]);
	}
}
