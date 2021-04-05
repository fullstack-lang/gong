// generated from NgDetailTemplateTS
import { Component, OnInit } from '@angular/core';
import { FormControl } from '@angular/forms';

import { BookDB } from '../book-db'
import { BookService } from '../book.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'
import { MapOfComponents } from '../map-components'

// insertion point for imports

import { Router, RouterState, ActivatedRoute } from '@angular/router';

import { MatDialog, MAT_DIALOG_DATA, MatDialogRef, MatDialogConfig } from '@angular/material/dialog';

import { NullInt64 } from '../front-repo.service'

@Component({
	selector: 'app-book-detail',
	templateUrl: './book-detail.component.html',
	styleUrls: ['./book-detail.component.css'],
})
export class BookDetailComponent implements OnInit {

	// insertion point for declarations
	RecommandedFormControl = new FormControl(false);

	// the BookDB of interest
	book: BookDB;

	// front repo
	frontRepo: FrontRepo

	constructor(
		private bookService: BookService,
		private frontRepoService: FrontRepoService,
		public dialog: MatDialog,
		private route: ActivatedRoute,
		private router: Router,
	) {
		// https://stackoverflow.com/questions/54627478/angular-7-routing-to-same-component-but-different-param-not-working
		// this is for routerLink on same component when only queryParameter changes
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getBook()

		// observable for changes in structs
		this.bookService.BookServiceChanged.subscribe(
			message => {
				if (message == "post" || message == "update" || message == "delete") {
					this.getBook()
				}
			}
		)

		// insertion point for initialisation of enums list
	}

	getBook(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		const association = this.route.snapshot.paramMap.get('association');

		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo
				console.log("front repo BookPull returned")

				if (id != 0 && association == undefined) {
					this.book = frontRepo.Books.get(id)
				} else {
					this.book = new (BookDB)
				}

				// insertion point for recovery of form controls value for bool fields
				this.RecommandedFormControl.setValue(this.book.Recommanded)
			}
		)


	}

	save(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		const association = this.route.snapshot.paramMap.get('association');

		// insertion point for saving value of form controls of boolean fields
		this.book.Recommanded = this.RecommandedFormControl.value
		if (this.book.AreaID == undefined) {
			this.book.AreaID = new NullInt64
		}
		if (this.book.Area != undefined) {
			this.book.AreaID.Int64 = this.book.Area.ID
			this.book.AreaID.Valid = true
			this.book.AreaName = this.book.Area.Name
		} else {
			this.book.AreaID.Int64 = 0
			this.book.AreaID.Valid = true
			this.book.AreaName = ""
		}

		if (id != 0 && association == undefined) {
			// insertion point for saving value of reverse pointers
			if (this.book.Editor_Books_reverse != undefined) {
				this.book.Editor_BooksDBID = new NullInt64
				this.book.Editor_BooksDBID.Int64 = this.book.Editor_Books_reverse.ID
				this.book.Editor_BooksDBID.Valid = true
				this.book.Editor_Books_reverse = undefined // very important, otherwise, circular JSON
			}

			this.bookService.updateBook(this.book)
				.subscribe(book => {
					this.bookService.BookServiceChanged.next("update")

					console.log("book saved")
				});
		} else {
			switch (association) {
				// insertion point for saving value of ONE_MANY association reverse pointer
				case "Editor_Books":
					this.book.Editor_BooksDBID = new NullInt64
					this.book.Editor_BooksDBID.Int64 = id
					this.book.Editor_BooksDBID.Valid = true
					break
			}
			this.bookService.postBook(this.book).subscribe(book => {

				this.bookService.BookServiceChanged.next("post")

				this.book = {} // reset fields
				console.log("book added")
			});
		}
	}

	// openReverseSelection is a generic function that calls dialog for the edition of 
	// ONE-MANY association
	// It uses the MapOfComponent provided by the front repo
	openReverseSelection(AssociatedStruct: string, reverseField: string) {

		const dialogConfig = new MatDialogConfig();

		// dialogConfig.disableClose = true;
		dialogConfig.autoFocus = true;
		dialogConfig.data = {
			ID: this.book.ID,
			ReversePointer: reverseField,
		};
		const dialogRef: MatDialogRef<string, any> = this.dialog.open(
			MapOfComponents.get(AssociatedStruct).get(
				AssociatedStruct + 'sTableComponent'
			),
			dialogConfig
		);

		dialogRef.afterClosed().subscribe(result => {
			console.log('The dialog was closed');
		});
	}
}
