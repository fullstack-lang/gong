import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';

import { AclassBclass2UseDB } from '../aclassbclass2use-db'
import { AclassBclass2UseService } from '../aclassbclass2use.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

export interface aclassbclass2useDummyElement {
}

const ELEMENT_DATA: aclassbclass2useDummyElement[] = [
];

@Component({
	selector: 'app-aclassbclass2use-presentation',
	templateUrl: './aclassbclass2use-presentation.component.html',
	styleUrls: ['./aclassbclass2use-presentation.component.css'],
})
export class AclassBclass2UsePresentationComponent implements OnInit {

	// insertion point for declarations

	displayedColumns: string[] = [];
	dataSource = ELEMENT_DATA;

	aclassbclass2use: AclassBclass2UseDB;

	// front repo
	frontRepo: FrontRepo
 
	constructor(
		private aclassbclass2useService: AclassBclass2UseService,
		private frontRepoService: FrontRepoService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getAclassBclass2Use();

		// observable for changes in 
		this.aclassbclass2useService.AclassBclass2UseServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getAclassBclass2Use()
				}
			}
		)
	}

	getAclassBclass2Use(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				this.aclassbclass2use = this.frontRepo.AclassBclass2Uses.get(id)

				// insertion point for recovery of durations
			}
		);
	}

	// set presentation outlet
	setPresentationRouterOutlet(structName: string, ID: number) {
		this.router.navigate([{
			outlets: {
				github_com_fullstack_lang_gong_test_go_presentation: ["github_com_fullstack_lang_gong_test_go-" + structName + "-presentation", ID]
			}
		}]);
	}

	// set editor outlet
	setEditorRouterOutlet(ID: number) {
		this.router.navigate([{
			outlets: {
				github_com_fullstack_lang_gong_test_go_editor: ["github_com_fullstack_lang_gong_test_go-" + "aclassbclass2use-detail", ID]
			}
		}]);
	}
}
