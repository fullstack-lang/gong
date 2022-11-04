import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';

import { NodeDB } from '../node-db'
import { NodeService } from '../node.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

// insertion point for additional imports

export interface nodeDummyElement {
}

const ELEMENT_DATA: nodeDummyElement[] = [
];

@Component({
	selector: 'app-node-presentation',
	templateUrl: './node-presentation.component.html',
	styleUrls: ['./node-presentation.component.css'],
})
export class NodePresentationComponent implements OnInit {

	// insertion point for additionnal time duration declarations
	// insertion point for additionnal enum int field declarations

	displayedColumns: string[] = []
	dataSource = ELEMENT_DATA

	node: NodeDB = new (NodeDB)

	// front repo
	frontRepo: FrontRepo = new (FrontRepo)
 
	constructor(
		private nodeService: NodeService,
		private frontRepoService: FrontRepoService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getNode();

		// observable for changes in 
		this.nodeService.NodeServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getNode()
				}
			}
		)
	}

	getNode(): void {
		const id = +this.route.snapshot.paramMap.get('id')!
		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				this.node = this.frontRepo.Nodes.get(id)!

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
				github_com_fullstack_lang_gongdoc_go_editor: ["github_com_fullstack_lang_gongdoc_go-" + "node-detail", ID]
			}
		}]);
	}
}
