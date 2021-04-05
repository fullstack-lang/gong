import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup} from '@angular/forms';

import { GongdocCommandDB } from '../gongdoccommand-db'
import { GongdocCommandService } from '../gongdoccommand.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

export interface gongdoccommandDummyElement {
}

const ELEMENT_DATA: gongdoccommandDummyElement[] = [
];

@Component({
	selector: 'app-gongdoccommand-presentation',
	templateUrl: './gongdoccommand-presentation.component.html',
  styleUrls: ['./gongdoccommand-presentation.component.css'],
})
export class GongdocCommandPresentationComponent implements OnInit {

	displayedColumns: string[] = [];
	dataSource = ELEMENT_DATA;

	gongdoccommand: GongdocCommandDB;
 
	constructor(
		private gongdoccommandService: GongdocCommandService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getGongdocCommand();

		// observable for changes in 
		this.gongdoccommandService.GongdocCommandServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getGongdocCommand()
				}
			}
		)
	}

	getGongdocCommand(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		this.gongdoccommandService.getGongdocCommand(id)
			.subscribe(
				gongdoccommand => {
					this.gongdoccommand = gongdoccommand
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
				editor: ["gongdoccommand-detail", ID]
			}
		}]);
	}
}
