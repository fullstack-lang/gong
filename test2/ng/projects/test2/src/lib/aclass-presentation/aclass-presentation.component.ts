import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';

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
			}
		);
	}

	// set presentation outlet
	setPresentationRouterOutlet(structName: string, ID: number) {
		this.router.navigate([{
			outlets: {
				github_com_fullstack_lang_gong_test2_go_presentation: ["github_com_fullstack_lang_gong_test2_go-" + structName + "-presentation", ID]
			}
		}]);
	}

	// set editor outlet
	setEditorRouterOutlet(ID: number) {
		this.router.navigate([{
			outlets: {
				github_com_fullstack_lang_gong_test2_go_editor: ["github_com_fullstack_lang_gong_test2_go-" + "aclass-detail", ID]
			}
		}]);
	}
}
