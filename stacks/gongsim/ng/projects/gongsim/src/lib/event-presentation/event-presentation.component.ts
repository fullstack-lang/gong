import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup} from '@angular/forms';

import { EventDB } from '../event-db'
import { EventService } from '../event.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

export interface eventDummyElement {
}

const ELEMENT_DATA: eventDummyElement[] = [
];

@Component({
	selector: 'app-event-presentation',
	templateUrl: './event-presentation.component.html',
  styleUrls: ['./event-presentation.component.css'],
})
export class EventPresentationComponent implements OnInit {

	displayedColumns: string[] = [];
	dataSource = ELEMENT_DATA;

	event: EventDB;
 
	constructor(
		private eventService: EventService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getEvent();

		// observable for changes in 
		this.eventService.EventServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getEvent()
				}
			}
		)
	}

	getEvent(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		this.eventService.getEvent(id)
			.subscribe(
				event => {
					this.event = event
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
				editor: ["event-detail", ID]
			}
		}]);
	}
}
