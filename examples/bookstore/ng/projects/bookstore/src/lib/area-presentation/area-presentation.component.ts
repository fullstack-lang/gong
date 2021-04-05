import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup} from '@angular/forms';

import { AreaDB } from '../area-db'
import { AreaService } from '../area.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

export interface areaDummyElement {
}

const ELEMENT_DATA: areaDummyElement[] = [
];

@Component({
	selector: 'app-area-presentation',
	templateUrl: './area-presentation.component.html',
  styleUrls: ['./area-presentation.component.css'],
})
export class AreaPresentationComponent implements OnInit {

	displayedColumns: string[] = [];
	dataSource = ELEMENT_DATA;

	area: AreaDB;
 
	constructor(
		private areaService: AreaService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getArea();

		// observable for changes in 
		this.areaService.AreaServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getArea()
				}
			}
		)
	}

	getArea(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		this.areaService.getArea(id)
			.subscribe(
				area => {
					this.area = area
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
				editor: ["area-detail", ID]
			}
		}]);
	}
}
