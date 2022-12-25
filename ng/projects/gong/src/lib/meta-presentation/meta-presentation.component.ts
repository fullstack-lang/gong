import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';

import { MetaDB } from '../meta-db'
import { MetaService } from '../meta.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

// insertion point for additional imports

export interface metaDummyElement {
}

const ELEMENT_DATA: metaDummyElement[] = [
];

@Component({
	selector: 'app-meta-presentation',
	templateUrl: './meta-presentation.component.html',
	styleUrls: ['./meta-presentation.component.css'],
})
export class MetaPresentationComponent implements OnInit {

	// insertion point for additionnal time duration declarations
	// insertion point for additionnal enum int field declarations

	displayedColumns: string[] = []
	dataSource = ELEMENT_DATA

	meta: MetaDB = new (MetaDB)

	// front repo
	frontRepo: FrontRepo = new (FrontRepo)
 
	constructor(
		private metaService: MetaService,
		private frontRepoService: FrontRepoService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getMeta();

		// observable for changes in 
		this.metaService.MetaServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getMeta()
				}
			}
		)
	}

	getMeta(): void {
		const id = +this.route.snapshot.paramMap.get('id')!
		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				this.meta = this.frontRepo.Metas.get(id)!

				// insertion point for recovery of durations
				// insertion point for recovery of enum tint
			}
		);
	}

	// set presentation outlet
	setPresentationRouterOutlet(structName: string, ID: number) {
		this.router.navigate([{
			outlets: {
				github_com_fullstack_lang_gong_go_presentation: ["github_com_fullstack_lang_gong_go-" + structName + "-presentation", ID]
			}
		}]);
	}

	// set editor outlet
	setEditorRouterOutlet(ID: number) {
		this.router.navigate([{
			outlets: {
				github_com_fullstack_lang_gong_go_editor: ["github_com_fullstack_lang_gong_go-" + "meta-detail", ID]
			}
		}]);
	}
}
