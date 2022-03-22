import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';

import { PositionDB } from '../position-db'
import { PositionService } from '../position.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

// insertion point for additional imports

export interface positionDummyElement {
}

const ELEMENT_DATA: positionDummyElement[] = [
];

@Component({
	selector: 'app-position-presentation',
	templateUrl: './position-presentation.component.html',
	styleUrls: ['./position-presentation.component.css'],
})
export class PositionPresentationComponent implements OnInit {

	// insertion point for additionnal time duration declarations
	// insertion point for additionnal enum int field declarations

	displayedColumns: string[] = []
	dataSource = ELEMENT_DATA

	position: PositionDB = new (PositionDB)

	// front repo
	frontRepo: FrontRepo = new (FrontRepo)
 
	constructor(
		private positionService: PositionService,
		private frontRepoService: FrontRepoService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getPosition();

		// observable for changes in 
		this.positionService.PositionServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getPosition()
				}
			}
		)
	}

	getPosition(): void {
		const id = +this.route.snapshot.paramMap.get('id')!
		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				this.position = this.frontRepo.Positions.get(id)!

				// insertion point for recovery of durations
				// insertion point for recovery of enum tint
			}
		);
	}

	// set presentation outlet
	setPresentationRouterOutlet(structName: string, ID: number) {
		this.router.navigate([{
			outlets: {
				github_com_fullstack_lang_gongdoc_go_presentation: ["github_com_fullstack_lang_gongdoc_go-" + structName + "-presentation", ID]
			}
		}]);
	}

	// set editor outlet
	setEditorRouterOutlet(ID: number) {
		this.router.navigate([{
			outlets: {
				github_com_fullstack_lang_gongdoc_go_editor: ["github_com_fullstack_lang_gongdoc_go-" + "position-detail", ID]
			}
		}]);
	}
}
