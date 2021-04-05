import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup} from '@angular/forms';

import { GongEnumDB } from '../gongenum-db'
import { GongEnumService } from '../gongenum.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

export interface gongenumDummyElement {
}

const ELEMENT_DATA: gongenumDummyElement[] = [
];

@Component({
	selector: 'app-gongenum-presentation',
	templateUrl: './gongenum-presentation.component.html',
  styleUrls: ['./gongenum-presentation.component.css'],
})
export class GongEnumPresentationComponent implements OnInit {

	displayedColumns: string[] = [];
	dataSource = ELEMENT_DATA;

	gongenum: GongEnumDB;
 
	constructor(
		private gongenumService: GongEnumService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getGongEnum();

		// observable for changes in 
		this.gongenumService.GongEnumServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getGongEnum()
				}
			}
		)
	}

	getGongEnum(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		this.gongenumService.getGongEnum(id)
			.subscribe(
				gongenum => {
					this.gongenum = gongenum
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
				editor: ["gongenum-detail", ID]
			}
		}]);
	}
}
