import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup} from '@angular/forms';

import { GongEnumValueDB } from '../gongenumvalue-db'
import { GongEnumValueService } from '../gongenumvalue.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

export interface gongenumvalueDummyElement {
}

const ELEMENT_DATA: gongenumvalueDummyElement[] = [
];

@Component({
	selector: 'app-gongenumvalue-presentation',
	templateUrl: './gongenumvalue-presentation.component.html',
  styleUrls: ['./gongenumvalue-presentation.component.css'],
})
export class GongEnumValuePresentationComponent implements OnInit {

	displayedColumns: string[] = [];
	dataSource = ELEMENT_DATA;

	gongenumvalue: GongEnumValueDB;
 
	constructor(
		private gongenumvalueService: GongEnumValueService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getGongEnumValue();

		// observable for changes in 
		this.gongenumvalueService.GongEnumValueServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getGongEnumValue()
				}
			}
		)
	}

	getGongEnumValue(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		this.gongenumvalueService.getGongEnumValue(id)
			.subscribe(
				gongenumvalue => {
					this.gongenumvalue = gongenumvalue
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
				editor: ["gongenumvalue-detail", ID]
			}
		}]);
	}
}
