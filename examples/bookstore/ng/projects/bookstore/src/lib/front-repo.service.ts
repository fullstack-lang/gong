import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';

import { Observable, combineLatest } from 'rxjs';

// insertion point sub template for services imports 
import { AreaDB } from './area-db'
import { AreaService } from './area.service'

import { BookDB } from './book-db'
import { BookService } from './book.service'

import { EditorDB } from './editor-db'
import { EditorService } from './editor.service'


// FrontRepo stores all instances in a front repository (design pattern repository)
export class FrontRepo { // insertion point sub template 
  Areas_array = new Array<AreaDB>(); // array of repo instances
  Areas = new Map<number, AreaDB>(); // map of repo instances
  Areas_batch = new Map<number, AreaDB>(); // same but only in last GET (for finding repo instances to delete)
  Books_array = new Array<BookDB>(); // array of repo instances
  Books = new Map<number, BookDB>(); // map of repo instances
  Books_batch = new Map<number, BookDB>(); // same but only in last GET (for finding repo instances to delete)
  Editors_array = new Array<EditorDB>(); // array of repo instances
  Editors = new Map<number, EditorDB>(); // map of repo instances
  Editors_batch = new Map<number, EditorDB>(); // same but only in last GET (for finding repo instances to delete)
}

//
// Store of all instances of the stack
//
export const FrontRepoSingloton = new (FrontRepo)

// define the type of nullable Int64 in order to support back pointers IDs
export class NullInt64 {
    Int64: number
    Valid: boolean
}

// define the interface for information that is forwarded from the calling instance to 
// the select table
export interface DialogData {
  ID: number; // ID of the calling instance
  ReversePointer: string; // field of {{Structname}} that serve as reverse pointer
}

//
// observable that fetch all elements of the stack and store them in the FrontRepo
//
@Injectable({
  providedIn: 'root'
})
export class FrontRepoService {

  httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' })
  };

  constructor(
    private http: HttpClient, // insertion point sub template 
    private areaService: AreaService,
    private bookService: BookService,
    private editorService: EditorService,
  ) { }

  // typing of observable can be messy in typescript. Therefore, one force the type
  observableFrontRepo: [ // insertion point sub template 
    Observable<AreaDB[]>,
    Observable<BookDB[]>,
    Observable<EditorDB[]>,
  ] = [ // insertion point sub template 
      this.areaService.getAreas(),
      this.bookService.getBooks(),
      this.editorService.getEditors(),
    ];

  //
  // pull performs a GET on all struct of the stack and redeem association pointers 
  //
  // This is an observable. Therefore, the control flow forks with
  // - pull() return immediatly the observable
  // - the observable observer, if it subscribe, is called when all GET calls are performs
  pull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest(
          this.observableFrontRepo
        ).subscribe(
          ([ // insertion point sub template for declarations 
            areas_,
            books_,
            editors_,
          ]) => {
            // Typing can be messy with many items. Therefore, type casting is necessary here
            // insertion point sub template for type casting 
            var areas: AreaDB[]
            areas = areas_
            var books: BookDB[]
            books = books_
            var editors: EditorDB[]
            editors = editors_

            // 
            // First Step: init map of instances
            // insertion point sub template for init 
            // init the array
            FrontRepoSingloton.Areas_array = areas

            // clear the map that counts Area in the GET
            FrontRepoSingloton.Areas_batch.clear()
            
            areas.forEach(
              area => {
                FrontRepoSingloton.Areas.set(area.ID, area)
                FrontRepoSingloton.Areas_batch.set(area.ID, area)
              }
            )
            
            // clear areas that are absent from the batch
            FrontRepoSingloton.Areas.forEach(
              area => {
                if (FrontRepoSingloton.Areas_batch.get(area.ID) == undefined) {
                  FrontRepoSingloton.Areas.delete(area.ID)
                }
              }
            )
            // init the array
            FrontRepoSingloton.Books_array = books

            // clear the map that counts Book in the GET
            FrontRepoSingloton.Books_batch.clear()
            
            books.forEach(
              book => {
                FrontRepoSingloton.Books.set(book.ID, book)
                FrontRepoSingloton.Books_batch.set(book.ID, book)
              }
            )
            
            // clear books that are absent from the batch
            FrontRepoSingloton.Books.forEach(
              book => {
                if (FrontRepoSingloton.Books_batch.get(book.ID) == undefined) {
                  FrontRepoSingloton.Books.delete(book.ID)
                }
              }
            )
            // init the array
            FrontRepoSingloton.Editors_array = editors

            // clear the map that counts Editor in the GET
            FrontRepoSingloton.Editors_batch.clear()
            
            editors.forEach(
              editor => {
                FrontRepoSingloton.Editors.set(editor.ID, editor)
                FrontRepoSingloton.Editors_batch.set(editor.ID, editor)
              }
            )
            
            // clear editors that are absent from the batch
            FrontRepoSingloton.Editors.forEach(
              editor => {
                if (FrontRepoSingloton.Editors_batch.get(editor.ID) == undefined) {
                  FrontRepoSingloton.Editors.delete(editor.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template for redeem 
            areas.forEach(
              area => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointer field Subarea redeeming
                {
                  let _area = FrontRepoSingloton.Areas.get(area.SubareaID.Int64)
                  if (_area) {
                    area.Subarea = _area
                  }
                }

                // insertion point for redeeming ONE-MANY associations
              }
            )
            books.forEach(
              book => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointer field Area redeeming
                {
                  let _area = FrontRepoSingloton.Areas.get(book.AreaID.Int64)
                  if (_area) {
                    book.Area = _area
                  }
                }

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Editor.Books redeeming
                {
                  let _editor = FrontRepoSingloton.Editors.get(book.Editor_BooksDBID.Int64)
                  if (_editor) {
                    if (_editor.Books == undefined) {
                      _editor.Books = new Array<EditorDB>()
                    }
                    _editor.Books.push(book)
                    if (book.Editor_Books_reverse == undefined) {
                      book.Editor_Books_reverse = _editor
                    }
                  }
                }
              }
            )
            editors.forEach(
              editor => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
              }
            )

            // hand over control flow to observer
            observer.next(FrontRepoSingloton)
          }
        )
      }
    )
  }

  // insertion point for pull per struct 

  // AreaPull performs a GET on Area of the stack and redeem association pointers 
  AreaPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.areaService.getAreas()
        ]).subscribe(
          ([ // insertion point sub template 
            areas,
          ]) => {
            // init the array
            FrontRepoSingloton.Areas_array = areas

            // clear the map that counts Area in the GET
            FrontRepoSingloton.Areas_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            areas.forEach(
              area => {
                FrontRepoSingloton.Areas.set(area.ID, area)
                FrontRepoSingloton.Areas_batch.set(area.ID, area)

                // insertion point for redeeming ONE/ZERO-ONE associations 
                // insertion point for pointer field Subarea redeeming
                {
                  let _area = FrontRepoSingloton.Areas.get(area.SubareaID.Int64)
                  if (_area) {
                    area.Subarea = _area
                  }
                }

                // insertion point for redeeming ONE-MANY associations 
              }
            )

            // clear areas that are absent from the GET
            FrontRepoSingloton.Areas.forEach(
              area => {
                if (FrontRepoSingloton.Areas_batch.get(area.ID) == undefined) {
                  FrontRepoSingloton.Areas.delete(area.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(FrontRepoSingloton)
          }
        )
      }
    )
  }

  // BookPull performs a GET on Book of the stack and redeem association pointers 
  BookPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.bookService.getBooks()
        ]).subscribe(
          ([ // insertion point sub template 
            books,
          ]) => {
            // init the array
            FrontRepoSingloton.Books_array = books

            // clear the map that counts Book in the GET
            FrontRepoSingloton.Books_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            books.forEach(
              book => {
                FrontRepoSingloton.Books.set(book.ID, book)
                FrontRepoSingloton.Books_batch.set(book.ID, book)

                // insertion point for redeeming ONE/ZERO-ONE associations 
                // insertion point for pointer field Area redeeming
                {
                  let _area = FrontRepoSingloton.Areas.get(book.AreaID.Int64)
                  if (_area) {
                    book.Area = _area
                  }
                }

                // insertion point for redeeming ONE-MANY associations 
                // insertion point for slice of pointer field Editor.Books redeeming
                {
                  let _editor = FrontRepoSingloton.Editors.get(book.Editor_BooksDBID.Int64)
                  if (_editor) {
                    if (_editor.Books == undefined) {
                      _editor.Books = new Array<EditorDB>()
                    }
                    _editor.Books.push(book)
                    if (book.Editor_Books_reverse == undefined) {
                      book.Editor_Books_reverse = _editor
                    }
                  }
                }
              }
            )

            // clear books that are absent from the GET
            FrontRepoSingloton.Books.forEach(
              book => {
                if (FrontRepoSingloton.Books_batch.get(book.ID) == undefined) {
                  FrontRepoSingloton.Books.delete(book.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(FrontRepoSingloton)
          }
        )
      }
    )
  }

  // EditorPull performs a GET on Editor of the stack and redeem association pointers 
  EditorPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.editorService.getEditors()
        ]).subscribe(
          ([ // insertion point sub template 
            editors,
          ]) => {
            // init the array
            FrontRepoSingloton.Editors_array = editors

            // clear the map that counts Editor in the GET
            FrontRepoSingloton.Editors_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            editors.forEach(
              editor => {
                FrontRepoSingloton.Editors.set(editor.ID, editor)
                FrontRepoSingloton.Editors_batch.set(editor.ID, editor)

                // insertion point for redeeming ONE/ZERO-ONE associations 

                // insertion point for redeeming ONE-MANY associations 
              }
            )

            // clear editors that are absent from the GET
            FrontRepoSingloton.Editors.forEach(
              editor => {
                if (FrontRepoSingloton.Editors_batch.get(editor.ID) == undefined) {
                  FrontRepoSingloton.Editors.delete(editor.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(FrontRepoSingloton)
          }
        )
      }
    )
  }
}

// insertion point for get unique ID per struct 
export function getAreaUniqueID(id: number): number {
  return 31 * id
}
export function getBookUniqueID(id: number): number {
  return 37 * id
}
export function getEditorUniqueID(id: number): number {
  return 41 * id
}
