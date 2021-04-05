import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup} from '@angular/forms';

import { WasherDB } from '../washer-db'
import { WasherService } from '../washer.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

export interface washerDummyElement {
}

const ELEMENT_DATA: washerDummyElement[] = [
];

@Component({
	selector: 'app-washer-presentation',
	templateUrl: './washer-presentation.component.html',
  styleUrls: ['./washer-presentation.component.css'],
})
export class WasherPresentationComponent implements OnInit {

	displayedColumns: string[] = [];
	dataSource = ELEMENT_DATA;

	washer: WasherDB;
 
	constructor(
		private washerService: WasherService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getWasher();

		// observable for changes in 
		this.washerService.WasherServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getWasher()
				}
			}
		)
	}

	getWasher(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		this.washerService.getWasher(id)
			.subscribe(
				washer => {
					this.washer = washer
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
				editor: ["washer-detail", ID]
			}
		}]);
	}
}
