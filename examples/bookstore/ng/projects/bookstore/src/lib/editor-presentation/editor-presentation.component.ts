import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup} from '@angular/forms';

import { EditorDB } from '../editor-db'
import { EditorService } from '../editor.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

export interface editorDummyElement {
}

const ELEMENT_DATA: editorDummyElement[] = [
];

@Component({
	selector: 'app-editor-presentation',
	templateUrl: './editor-presentation.component.html',
  styleUrls: ['./editor-presentation.component.css'],
})
export class EditorPresentationComponent implements OnInit {

	displayedColumns: string[] = [];
	dataSource = ELEMENT_DATA;

	editor: EditorDB;
 
	constructor(
		private editorService: EditorService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getEditor();

		// observable for changes in 
		this.editorService.EditorServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getEditor()
				}
			}
		)
	}

	getEditor(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		this.editorService.getEditor(id)
			.subscribe(
				editor => {
					this.editor = editor
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
				editor: ["editor-detail", ID]
			}
		}]);
	}
}
