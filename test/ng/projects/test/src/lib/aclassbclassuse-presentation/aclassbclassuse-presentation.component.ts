import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';

import { AclassBclassUseDB } from '../aclassbclassuse-db'
import { AclassBclassUseService } from '../aclassbclassuse.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

export interface aclassbclassuseDummyElement {
}

const ELEMENT_DATA: aclassbclassuseDummyElement[] = [
];

@Component({
	selector: 'app-aclassbclassuse-presentation',
	templateUrl: './aclassbclassuse-presentation.component.html',
	styleUrls: ['./aclassbclassuse-presentation.component.css'],
})
export class AclassBclassUsePresentationComponent implements OnInit {

	// insertion point for declarations

	displayedColumns: string[] = [];
	dataSource = ELEMENT_DATA;

	aclassbclassuse: AclassBclassUseDB;

	// front repo
	frontRepo: FrontRepo
 
	constructor(
		private aclassbclassuseService: AclassBclassUseService,
		private frontRepoService: FrontRepoService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getAclassBclassUse();

		// observable for changes in 
		this.aclassbclassuseService.AclassBclassUseServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getAclassBclassUse()
				}
			}
		)
	}

	getAclassBclassUse(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				this.aclassbclassuse = this.frontRepo.AclassBclassUses.get(id)

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
				github_com_fullstack_lang_gong_test_go_editor: ["github_com_fullstack_lang_gong_test_go-" + "aclassbclassuse-detail", ID]
			}
		}]);
	}
}
