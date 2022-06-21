import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';

import { AstructBstructUseDB } from '../astructbstructuse-db'
import { AstructBstructUseService } from '../astructbstructuse.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

// insertion point for additional imports

export interface astructbstructuseDummyElement {
}

const ELEMENT_DATA: astructbstructuseDummyElement[] = [
];

@Component({
	selector: 'app-astructbstructuse-presentation',
	templateUrl: './astructbstructuse-presentation.component.html',
	styleUrls: ['./astructbstructuse-presentation.component.css'],
})
export class AstructBstructUsePresentationComponent implements OnInit {

	// insertion point for additionnal time duration declarations
	// insertion point for additionnal enum int field declarations

	displayedColumns: string[] = []
	dataSource = ELEMENT_DATA

	astructbstructuse: AstructBstructUseDB = new (AstructBstructUseDB)

	// front repo
	frontRepo: FrontRepo = new (FrontRepo)
 
	constructor(
		private astructbstructuseService: AstructBstructUseService,
		private frontRepoService: FrontRepoService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getAstructBstructUse();

		// observable for changes in 
		this.astructbstructuseService.AstructBstructUseServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getAstructBstructUse()
				}
			}
		)
	}

	getAstructBstructUse(): void {
		const id = +this.route.snapshot.paramMap.get('id')!
		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				this.astructbstructuse = this.frontRepo.AstructBstructUses.get(id)!

				// insertion point for recovery of durations
				// insertion point for recovery of enum tint
			}
		);
	}

	// set presentation outlet
	setPresentationRouterOutlet(structName: string, ID: number) {
		this.router.navigate([{
			outlets: {
				_presentation: ["-" + structName + "-presentation", ID]
			}
		}]);
	}

	// set editor outlet
	setEditorRouterOutlet(ID: number) {
		this.router.navigate([{
			outlets: {
				_editor: ["-" + "astructbstructuse-detail", ID]
			}
		}]);
	}
}
