import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup} from '@angular/forms';

import { GongdocStatusDB } from '../gongdocstatus-db'
import { GongdocStatusService } from '../gongdocstatus.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

export interface gongdocstatusDummyElement {
}

const ELEMENT_DATA: gongdocstatusDummyElement[] = [
];

@Component({
	selector: 'app-gongdocstatus-presentation',
	templateUrl: './gongdocstatus-presentation.component.html',
  styleUrls: ['./gongdocstatus-presentation.component.css'],
})
export class GongdocStatusPresentationComponent implements OnInit {

	displayedColumns: string[] = [];
	dataSource = ELEMENT_DATA;

	gongdocstatus: GongdocStatusDB;
 
	constructor(
		private gongdocstatusService: GongdocStatusService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getGongdocStatus();

		// observable for changes in 
		this.gongdocstatusService.GongdocStatusServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getGongdocStatus()
				}
			}
		)
	}

	getGongdocStatus(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		this.gongdocstatusService.getGongdocStatus(id)
			.subscribe(
				gongdocstatus => {
					this.gongdocstatus = gongdocstatus
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
				editor: ["gongdocstatus-detail", ID]
			}
		}]);
	}
}
