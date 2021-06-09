import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';

import { DclassDB } from '../dclass-db'
import { DclassService } from '../dclass.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

export interface dclassDummyElement {
}

const ELEMENT_DATA: dclassDummyElement[] = [
];

@Component({
	selector: 'app-dclass-presentation',
	templateUrl: './dclass-presentation.component.html',
	styleUrls: ['./dclass-presentation.component.css'],
})
export class DclassPresentationComponent implements OnInit {

	// insertion point for declarations

	displayedColumns: string[] = [];
	dataSource = ELEMENT_DATA;

	dclass: DclassDB;

	// front repo
	frontRepo: FrontRepo
 
	constructor(
		private dclassService: DclassService,
		private frontRepoService: FrontRepoService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getDclass();

		// observable for changes in 
		this.dclassService.DclassServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getDclass()
				}
			}
		)
	}

	getDclass(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				this.dclass = this.frontRepo.Dclasss.get(id)

				// insertion point for recovery of durations
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
				editor: ["dclass-detail", ID]
			}
		}]);
	}
}
