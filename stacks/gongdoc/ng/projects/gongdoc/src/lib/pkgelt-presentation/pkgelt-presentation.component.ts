import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup} from '@angular/forms';

import { PkgeltDB } from '../pkgelt-db'
import { PkgeltService } from '../pkgelt.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

export interface pkgeltDummyElement {
}

const ELEMENT_DATA: pkgeltDummyElement[] = [
];

@Component({
	selector: 'app-pkgelt-presentation',
	templateUrl: './pkgelt-presentation.component.html',
  styleUrls: ['./pkgelt-presentation.component.css'],
})
export class PkgeltPresentationComponent implements OnInit {

	displayedColumns: string[] = [];
	dataSource = ELEMENT_DATA;

	pkgelt: PkgeltDB;
 
	constructor(
		private pkgeltService: PkgeltService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getPkgelt();

		// observable for changes in 
		this.pkgeltService.PkgeltServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getPkgelt()
				}
			}
		)
	}

	getPkgelt(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		this.pkgeltService.getPkgelt(id)
			.subscribe(
				pkgelt => {
					this.pkgelt = pkgelt
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
				editor: ["pkgelt-detail", ID]
			}
		}]);
	}
}
