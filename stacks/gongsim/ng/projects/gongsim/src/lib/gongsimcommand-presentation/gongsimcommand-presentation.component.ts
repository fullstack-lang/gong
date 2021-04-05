import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup} from '@angular/forms';

import { GongsimCommandDB } from '../gongsimcommand-db'
import { GongsimCommandService } from '../gongsimcommand.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

export interface gongsimcommandDummyElement {
}

const ELEMENT_DATA: gongsimcommandDummyElement[] = [
];

@Component({
	selector: 'app-gongsimcommand-presentation',
	templateUrl: './gongsimcommand-presentation.component.html',
  styleUrls: ['./gongsimcommand-presentation.component.css'],
})
export class GongsimCommandPresentationComponent implements OnInit {

	displayedColumns: string[] = [];
	dataSource = ELEMENT_DATA;

	gongsimcommand: GongsimCommandDB;
 
	constructor(
		private gongsimcommandService: GongsimCommandService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getGongsimCommand();

		// observable for changes in 
		this.gongsimcommandService.GongsimCommandServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getGongsimCommand()
				}
			}
		)
	}

	getGongsimCommand(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		this.gongsimcommandService.getGongsimCommand(id)
			.subscribe(
				gongsimcommand => {
					this.gongsimcommand = gongsimcommand
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
				editor: ["gongsimcommand-detail", ID]
			}
		}]);
	}
}
