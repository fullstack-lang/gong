import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup} from '@angular/forms';

import { SimulationDB } from '../simulation-db'
import { SimulationService } from '../simulation.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

export interface simulationDummyElement {
}

const ELEMENT_DATA: simulationDummyElement[] = [
];

@Component({
	selector: 'app-simulation-presentation',
	templateUrl: './simulation-presentation.component.html',
  styleUrls: ['./simulation-presentation.component.css'],
})
export class SimulationPresentationComponent implements OnInit {

	displayedColumns: string[] = [];
	dataSource = ELEMENT_DATA;

	simulation: SimulationDB;
 
	constructor(
		private simulationService: SimulationService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getSimulation();

		// observable for changes in 
		this.simulationService.SimulationServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getSimulation()
				}
			}
		)
	}

	getSimulation(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		this.simulationService.getSimulation(id)
			.subscribe(
				simulation => {
					this.simulation = simulation
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
				editor: ["simulation-detail", ID]
			}
		}]);
	}
}
