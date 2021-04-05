import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup} from '@angular/forms';

import { PointerToGongStructFieldDB } from '../pointertogongstructfield-db'
import { PointerToGongStructFieldService } from '../pointertogongstructfield.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

export interface pointertogongstructfieldDummyElement {
}

const ELEMENT_DATA: pointertogongstructfieldDummyElement[] = [
];

@Component({
	selector: 'app-pointertogongstructfield-presentation',
	templateUrl: './pointertogongstructfield-presentation.component.html',
  styleUrls: ['./pointertogongstructfield-presentation.component.css'],
})
export class PointerToGongStructFieldPresentationComponent implements OnInit {

	displayedColumns: string[] = [];
	dataSource = ELEMENT_DATA;

	pointertogongstructfield: PointerToGongStructFieldDB;
 
	constructor(
		private pointertogongstructfieldService: PointerToGongStructFieldService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getPointerToGongStructField();

		// observable for changes in 
		this.pointertogongstructfieldService.PointerToGongStructFieldServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getPointerToGongStructField()
				}
			}
		)
	}

	getPointerToGongStructField(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		this.pointertogongstructfieldService.getPointerToGongStructField(id)
			.subscribe(
				pointertogongstructfield => {
					this.pointertogongstructfield = pointertogongstructfield
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
				editor: ["pointertogongstructfield-detail", ID]
			}
		}]);
	}
}
