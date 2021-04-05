import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup} from '@angular/forms';

import { LinkDB } from '../link-db'
import { LinkService } from '../link.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

export interface linkDummyElement {
}

const ELEMENT_DATA: linkDummyElement[] = [
];

@Component({
	selector: 'app-link-presentation',
	templateUrl: './link-presentation.component.html',
  styleUrls: ['./link-presentation.component.css'],
})
export class LinkPresentationComponent implements OnInit {

	displayedColumns: string[] = [];
	dataSource = ELEMENT_DATA;

	link: LinkDB;
 
	constructor(
		private linkService: LinkService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getLink();

		// observable for changes in 
		this.linkService.LinkServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getLink()
				}
			}
		)
	}

	getLink(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		this.linkService.getLink(id)
			.subscribe(
				link => {
					this.link = link
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
				editor: ["link-detail", ID]
			}
		}]);
	}
}
