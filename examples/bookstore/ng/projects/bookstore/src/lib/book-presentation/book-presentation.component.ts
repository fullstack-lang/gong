import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup} from '@angular/forms';

import { BookDB } from '../book-db'
import { BookService } from '../book.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

export interface bookDummyElement {
}

const ELEMENT_DATA: bookDummyElement[] = [
];

@Component({
	selector: 'app-book-presentation',
	templateUrl: './book-presentation.component.html',
  styleUrls: ['./book-presentation.component.css'],
})
export class BookPresentationComponent implements OnInit {

	displayedColumns: string[] = [];
	dataSource = ELEMENT_DATA;

	book: BookDB;
 
	constructor(
		private bookService: BookService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getBook();

		// observable for changes in 
		this.bookService.BookServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getBook()
				}
			}
		)
	}

	getBook(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		this.bookService.getBook(id)
			.subscribe(
				book => {
					this.book = book
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
				editor: ["book-detail", ID]
			}
		}]);
	}
}
