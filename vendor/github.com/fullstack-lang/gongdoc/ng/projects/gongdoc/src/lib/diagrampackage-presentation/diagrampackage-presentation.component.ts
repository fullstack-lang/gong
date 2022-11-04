import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';

import { DiagramPackageDB } from '../diagrampackage-db'
import { DiagramPackageService } from '../diagrampackage.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

// insertion point for additional imports

export interface diagrampackageDummyElement {
}

const ELEMENT_DATA: diagrampackageDummyElement[] = [
];

@Component({
	selector: 'app-diagrampackage-presentation',
	templateUrl: './diagrampackage-presentation.component.html',
	styleUrls: ['./diagrampackage-presentation.component.css'],
})
export class DiagramPackagePresentationComponent implements OnInit {

	// insertion point for additionnal time duration declarations
	// insertion point for additionnal enum int field declarations

	displayedColumns: string[] = []
	dataSource = ELEMENT_DATA

	diagrampackage: DiagramPackageDB = new (DiagramPackageDB)

	// front repo
	frontRepo: FrontRepo = new (FrontRepo)
 
	constructor(
		private diagrampackageService: DiagramPackageService,
		private frontRepoService: FrontRepoService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getDiagramPackage();

		// observable for changes in 
		this.diagrampackageService.DiagramPackageServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getDiagramPackage()
				}
			}
		)
	}

	getDiagramPackage(): void {
		const id = +this.route.snapshot.paramMap.get('id')!
		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				this.diagrampackage = this.frontRepo.DiagramPackages.get(id)!

				// insertion point for recovery of durations
				// insertion point for recovery of enum tint
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
				github_com_fullstack_lang_gongdoc_go_editor: ["github_com_fullstack_lang_gongdoc_go-" + "diagrampackage-detail", ID]
			}
		}]);
	}
}
