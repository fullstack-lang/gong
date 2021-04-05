import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup} from '@angular/forms';

import { VerticeDB } from '../vertice-db'
import { VerticeService } from '../vertice.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

export interface verticeDummyElement {
}

const ELEMENT_DATA: verticeDummyElement[] = [
];

@Component({
	selector: 'app-vertice-presentation',
	templateUrl: './vertice-presentation.component.html',
  styleUrls: ['./vertice-presentation.component.css'],
})
export class VerticePresentationComponent implements OnInit {

	displayedColumns: string[] = [];
	dataSource = ELEMENT_DATA;

	vertice: VerticeDB;
 
	constructor(
		private verticeService: VerticeService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getVertice();

		// observable for changes in 
		this.verticeService.VerticeServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getVertice()
				}
			}
		)
	}

	getVertice(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		this.verticeService.getVertice(id)
			.subscribe(
				vertice => {
					this.vertice = vertice
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
				editor: ["vertice-detail", ID]
			}
		}]);
	}
}
