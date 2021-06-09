import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup} from '@angular/forms';

import { AclassDB } from '../aclass-db'
import { AclassService } from '../aclass.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

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

	// insertion point for declarations
	// fields from Duration1
	Duration1_Hours: number
	Duration1_Minutes: number
	Duration1_Seconds: number

	displayedColumns: string[] = [];
	dataSource = ELEMENT_DATA;

	aclass: AclassDB;

	// front repo
	frontRepo: FrontRepo
 
	constructor(
		private aclassService: AclassService,
		private frontRepoService: FrontRepoService,
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
		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				this.aclass = this.frontRepo.Aclasss.get(id)

				// insertion point for recovery of durations
				// computation of Hours, Minutes, Seconds for Duration1
				this.Duration1_Hours = Math.floor(this.aclass.Duration1 / (3600 * 1000 * 1000 * 1000))
				this.Duration1_Minutes = Math.floor(this.aclass.Duration1 % (3600 * 1000 * 1000 * 1000) / (60 * 1000 * 1000 * 1000))
				this.Duration1_Seconds = this.aclass.Duration1 % (60 * 1000 * 1000 * 1000) / (1000 * 1000 * 1000)
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
