import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup} from '@angular/forms';

import { ModelPkgDB } from '../modelpkg-db'
import { ModelPkgService } from '../modelpkg.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

export interface modelpkgDummyElement {
}

const ELEMENT_DATA: modelpkgDummyElement[] = [
];

@Component({
	selector: 'app-modelpkg-presentation',
	templateUrl: './modelpkg-presentation.component.html',
	styleUrls: ['./modelpkg-presentation.component.css'],
})
export class ModelPkgPresentationComponent implements OnInit {

	// insertion point for declarations

	displayedColumns: string[] = [];
	dataSource = ELEMENT_DATA;

	modelpkg: ModelPkgDB;
 
	constructor(
		private modelpkgService: ModelPkgService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getModelPkg();

		// observable for changes in 
		this.modelpkgService.ModelPkgServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getModelPkg()
				}
			}
		)
	}

	getModelPkg(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		this.modelpkgService.getModelPkg(id)
			.subscribe(
				modelpkg => {
					this.modelpkg = modelpkg

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
				editor: ["modelpkg-detail", ID]
			}
		}]);
	}
}
