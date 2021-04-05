// generated from NgTableTemplateTS
import { Component, OnInit, OnChanges, Input, Output, EventEmitter, Inject, Optional } from '@angular/core';
import { MatTableDataSource } from '@angular/material/table';
import { MatButton } from '@angular/material/button'

import { MatDialogRef, MAT_DIALOG_DATA, MatDialog } from '@angular/material/dialog'
import { DialogData } from '../front-repo.service'
import { SelectionModel } from '@angular/cdk/collections';

const allowMultiSelect = true;

import { Router, RouterState } from '@angular/router';
import { BookDB } from '../book-db'
import { BookService } from '../book.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

// generated table component
@Component({
  selector: 'app-books-table',
  templateUrl: './books-table.component.html',
  styleUrls: ['./books-table.component.css'],
})
export class BooksTableComponent implements OnInit {

  // used if the component is called as a selection component of Book instances
  selection: SelectionModel<BookDB>;
  initialSelection = new Array<BookDB>();

  // the data source for the table
  books: BookDB[];

  // front repo, that will be referenced by this.books
  frontRepo: FrontRepo

  // displayedColumns is referenced by the MatTable component for specify what columns
  // have to be displayed and in what order
  displayedColumns: string[];

  constructor(
    private bookService: BookService,
    private frontRepoService: FrontRepoService,

    // not null if the component is called as a selection component of book instances
    public dialogRef: MatDialogRef<BooksTableComponent>,
    @Optional() @Inject(MAT_DIALOG_DATA) public dialogData: DialogData,

    private router: Router,
  ) {
    // https://stackoverflow.com/questions/54627478/angular-7-routing-to-same-component-but-different-param-not-working
    // this is for routerLink on same component when only queryParameter changes
    this.router.routeReuseStrategy.shouldReuseRoute = function () {
      return false;
    };

    // observable for changes in structs
    this.bookService.BookServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.getBooks()
        }
      }
    )
    if (dialogData == undefined) {
  	  this.displayedColumns = ['ID', 'Edit', 'Delete', // insertion point for columns to display
        "Name",
        "Author",
        "City",
        "Year",
        "Price",
        "Recommanded",
        "Area",
        "Books",
      ]
    } else {
  	  this.displayedColumns = ['select', 'ID', // insertion point for columns to display
        "Name",
        "Author",
        "City",
        "Year",
        "Price",
        "Recommanded",
        "Area",
        "Books",
      ]
      this.selection = new SelectionModel<BookDB>(allowMultiSelect, this.initialSelection);
    }

  }

  ngOnInit(): void {
    this.getBooks()
  }

  getBooks(): void {
    this.frontRepoService.pull().subscribe(
      frontRepo => {
        this.frontRepo = frontRepo
        console.log("front repo pull returned")

        this.books = this.frontRepo.Books_array;
        // in case the component is called as a selection component
        if (this.dialogData != undefined) {
          this.books.forEach(
            book => {
              let ID = this.dialogData.ID
              let revPointer = book[this.dialogData.ReversePointer]
              if (revPointer.Int64 == ID) {
                this.initialSelection.push(book)
              }
            }
          )
          this.selection = new SelectionModel<BookDB>(allowMultiSelect, this.initialSelection);
        }
      }
    )
  }

  // newBook initiate a new book
  // create a new Book objet
  newBook() {
  }

  deleteBook(bookID: number, book: BookDB) {
    // list of books is truncated of book before the delete
    this.books = this.books.filter(h => h !== book);

    this.bookService.deleteBook(bookID).subscribe(
      book => {
        this.bookService.BookServiceChanged.next("delete")

        console.log("book deleted")
      }
    );
  }

  editBook(bookID: number, book: BookDB) {

  }

  // display book in router
  displayBookInRouter(bookID: number) {
    this.router.navigate( ["book-display", bookID])
  }

  // set editor outlet
  setEditorRouterOutlet(bookID: number) {
    this.router.navigate([{
      outlets: {
        editor: ["book-detail", bookID]
      }
    }]);
  }

  // set presentation outlet
  setPresentationRouterOutlet(bookID: number) {
    this.router.navigate([{
      outlets: {
        presentation: ["book-presentation", bookID]
      }
    }]);
  }

  /** Whether the number of selected elements matches the total number of rows. */
  isAllSelected() {
    const numSelected = this.selection.selected.length;
    const numRows = this.books.length;
    return numSelected === numRows;
  }

  /** Selects all rows if they are not all selected; otherwise clear selection. */
  masterToggle() {
    this.isAllSelected() ?
      this.selection.clear() :
      this.books.forEach(row => this.selection.select(row));
  }

  save() {

    let toUpdate = new Set<BookDB>()

    // reset all initial selection of book that belong to aclass through Anarrayofb
    this.initialSelection.forEach(
      book => {
        book[this.dialogData.ReversePointer].Int64 = 0
        book[this.dialogData.ReversePointer].Valid = true
        toUpdate.add(book)
      }
    )

    // from selection, set book that belong to aclass through Anarrayofb
    this.selection.selected.forEach(
      book => {
        console.log("selection ID " + book.ID)
        let ID = +this.dialogData.ID
        book[this.dialogData.ReversePointer].Int64 = ID
        book[this.dialogData.ReversePointer].Valid = true
        toUpdate.add(book)
      }
    )

    // update all book (only update selection & initial selection)
    toUpdate.forEach(
      book => {
        this.bookService.updateBook(book)
          .subscribe(book => {
            this.bookService.BookServiceChanged.next("update")
            console.log("book saved")
          });
      }
    )
    this.dialogRef.close('Pizza!');
  }
}
