import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup} from '@angular/forms';

import { UpdateStateDB } from '../updatestate-db'
import { UpdateStateService } from '../updatestate.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

export interface updatestateDummyElement {
}

const ELEMENT_DATA: updatestateDummyElement[] = [
];

@Component({
	selector: 'app-updatestate-presentation',
	templateUrl: './updatestate-presentation.component.html',
  styleUrls: ['./updatestate-presentation.component.css'],
})
export class UpdateStatePresentationComponent implements OnInit {

	displayedColumns: string[] = [];
	dataSource = ELEMENT_DATA;

	updatestate: UpdateStateDB;
 
	constructor(
		private updatestateService: UpdateStateService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getUpdateState();

		// observable for changes in 
		this.updatestateService.UpdateStateServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getUpdateState()
				}
			}
		)
	}

	getUpdateState(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		this.updatestateService.getUpdateState(id)
			.subscribe(
				updatestate => {
					this.updatestate = updatestate
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
				editor: ["updatestate-detail", ID]
			}
		}]);
	}
}
