import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup} from '@angular/forms';

import { ClassshapeDB } from '../classshape-db'
import { ClassshapeService } from '../classshape.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

export interface classshapeDummyElement {
}

const ELEMENT_DATA: classshapeDummyElement[] = [
];

@Component({
	selector: 'app-classshape-presentation',
	templateUrl: './classshape-presentation.component.html',
  styleUrls: ['./classshape-presentation.component.css'],
})
export class ClassshapePresentationComponent implements OnInit {

	displayedColumns: string[] = [];
	dataSource = ELEMENT_DATA;

	classshape: ClassshapeDB;
 
	constructor(
		private classshapeService: ClassshapeService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getClassshape();

		// observable for changes in 
		this.classshapeService.ClassshapeServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getClassshape()
				}
			}
		)
	}

	getClassshape(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		this.classshapeService.getClassshape(id)
			.subscribe(
				classshape => {
					this.classshape = classshape
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
				editor: ["classshape-detail", ID]
			}
		}]);
	}
}
