import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup} from '@angular/forms';

import { AclassDB } from '../aclass-db'
import { AclassService } from '../aclass.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

export interface aclassDummyElement {
}

const ELEMENT_DATA: aclassDummyElement[] = [
];

@Component({
	selector: 'app-aclass-presentation',
	templateUrl: './aclass-presentation.component.html',
  styleUrls: ['./aclass-presentation.component.css'],
})
export class AclassPresentationComponent implements OnInit {

	displayedColumns: string[] = [];
	dataSource = ELEMENT_DATA;

	aclass: AclassDB;
 
	constructor(
		private aclassService: AclassService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getAclass();

		// observable for changes in 
		this.aclassService.AclassServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getAclass()
				}
			}
		)
	}

	getAclass(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		this.aclassService.getAclass(id)
			.subscribe(
				aclass => {
					this.aclass = aclass
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
				editor: ["aclass-detail", ID]
			}
		}]);
	}
}
