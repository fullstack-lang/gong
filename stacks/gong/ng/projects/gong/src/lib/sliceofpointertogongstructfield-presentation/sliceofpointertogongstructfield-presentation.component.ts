import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup} from '@angular/forms';

import { SliceOfPointerToGongStructFieldDB } from '../sliceofpointertogongstructfield-db'
import { SliceOfPointerToGongStructFieldService } from '../sliceofpointertogongstructfield.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

export interface sliceofpointertogongstructfieldDummyElement {
}

const ELEMENT_DATA: sliceofpointertogongstructfieldDummyElement[] = [
];

@Component({
	selector: 'app-sliceofpointertogongstructfield-presentation',
	templateUrl: './sliceofpointertogongstructfield-presentation.component.html',
  styleUrls: ['./sliceofpointertogongstructfield-presentation.component.css'],
})
export class SliceOfPointerToGongStructFieldPresentationComponent implements OnInit {

	displayedColumns: string[] = [];
	dataSource = ELEMENT_DATA;

	sliceofpointertogongstructfield: SliceOfPointerToGongStructFieldDB;
 
	constructor(
		private sliceofpointertogongstructfieldService: SliceOfPointerToGongStructFieldService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getSliceOfPointerToGongStructField();

		// observable for changes in 
		this.sliceofpointertogongstructfieldService.SliceOfPointerToGongStructFieldServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getSliceOfPointerToGongStructField()
				}
			}
		)
	}

	getSliceOfPointerToGongStructField(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		this.sliceofpointertogongstructfieldService.getSliceOfPointerToGongStructField(id)
			.subscribe(
				sliceofpointertogongstructfield => {
					this.sliceofpointertogongstructfield = sliceofpointertogongstructfield
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
				editor: ["sliceofpointertogongstructfield-detail", ID]
			}
		}]);
	}
}
