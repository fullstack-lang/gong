import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';

import { VerticeDB } from '../vertice-db'
import { VerticeService } from '../vertice.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

export interface verticeDummyElement {
}

const ELEMENT_DATA: verticeDummyElement[] = [
];

@Component({
	selector: 'app-vertice-presentation',
	templateUrl: './vertice-presentation.component.html',
	styleUrls: ['./vertice-presentation.component.css'],
})
export class VerticePresentationComponent implements OnInit {

	// insertion point for declarations

	displayedColumns: string[] = []
	dataSource = ELEMENT_DATA

	vertice: VerticeDB = new (VerticeDB)

	// front repo
	frontRepo: FrontRepo = new (FrontRepo)
 
	constructor(
		private verticeService: VerticeService,
		private frontRepoService: FrontRepoService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getVertice();

		// observable for changes in 
		this.verticeService.VerticeServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getVertice()
				}
			}
		)
	}

	getVertice(): void {
		const id = +this.route.snapshot.paramMap.get('id')!
		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				this.vertice = this.frontRepo.Vertices.get(id)!

				// insertion point for recovery of durations
			}
		);
	}

	// set presentation outlet
	setPresentationRouterOutlet(structName: string, ID: number) {
		this.router.navigate([{
			outlets: {
				github_com_fullstack_lang_gongdoc_go_presentation: ["github_com_fullstack_lang_gongdoc_go-" + structName + "-presentation", ID]
			}
		}]);
	}

	// set editor outlet
	setEditorRouterOutlet(ID: number) {
		this.router.navigate([{
			outlets: {
				github_com_fullstack_lang_gongdoc_go_editor: ["github_com_fullstack_lang_gongdoc_go-" + "vertice-detail", ID]
			}
		}]);
	}
}
