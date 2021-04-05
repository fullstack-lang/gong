import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup} from '@angular/forms';

import { GongsimStatusDB } from '../gongsimstatus-db'
import { GongsimStatusService } from '../gongsimstatus.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

export interface gongsimstatusDummyElement {
}

const ELEMENT_DATA: gongsimstatusDummyElement[] = [
];

@Component({
	selector: 'app-gongsimstatus-presentation',
	templateUrl: './gongsimstatus-presentation.component.html',
  styleUrls: ['./gongsimstatus-presentation.component.css'],
})
export class GongsimStatusPresentationComponent implements OnInit {

	displayedColumns: string[] = [];
	dataSource = ELEMENT_DATA;

	gongsimstatus: GongsimStatusDB;
 
	constructor(
		private gongsimstatusService: GongsimStatusService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getGongsimStatus();

		// observable for changes in 
		this.gongsimstatusService.GongsimStatusServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getGongsimStatus()
				}
			}
		)
	}

	getGongsimStatus(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		this.gongsimstatusService.getGongsimStatus(id)
			.subscribe(
				gongsimstatus => {
					this.gongsimstatus = gongsimstatus
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
				editor: ["gongsimstatus-detail", ID]
			}
		}]);
	}
}
