// generated by ng_file_service_ts
import { Injectable, Component, Inject } from '@angular/core';
import { HttpParams } from '@angular/common/http';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { DOCUMENT, Location } from '@angular/common'

/*
 * Behavior subject
 */
import { BehaviorSubject } from 'rxjs'
import { Observable, of } from 'rxjs'
import { catchError, map, tap } from 'rxjs/operators'

import { BstructAPI } from './bstruct-api'
import { Bstruct, CopyBstructToBstructAPI } from './bstruct'

import { FrontRepo, FrontRepoService } from './front-repo.service';

// insertion point for imports

@Injectable({
  providedIn: 'root'
})
export class BstructService {

  // Kamar Raïmo: Adding a way to communicate between components that share information
  // so that they are notified of a change.
  BstructServiceChanged: BehaviorSubject<string> = new BehaviorSubject("");

  private bstructsUrl: string

  constructor(
    private http: HttpClient,
    @Inject(DOCUMENT) private document: Document
  ) {
    // path to the service share the same origin with the path to the document
    // get the origin in the URL to the document
    let origin = this.document.location.origin

    // if debugging with ng, replace 4200 with 8080
    origin = origin.replace("4200", "8080")

    // compute path to the service
    this.bstructsUrl = origin + '/api/github.com/fullstack-lang/gong/test/test/go/v1/bstructs';
  }

  /** GET bstructs from the server */
  // gets is more robust to refactoring
  gets(Name: string, frontRepo: FrontRepo): Observable<BstructAPI[]> {
    return this.getBstructs(Name, frontRepo)
  }
  getBstructs(Name: string, frontRepo: FrontRepo): Observable<BstructAPI[]> {

    let params = new HttpParams().set("Name", Name)

    return this.http.get<BstructAPI[]>(this.bstructsUrl, { params: params })
      .pipe(
        tap(),
        catchError(this.handleError<BstructAPI[]>('getBstructs', []))
      );
  }

  /** GET bstruct by id. Will 404 if id not found */
  // more robust API to refactoring
  get(id: number, Name: string, frontRepo: FrontRepo): Observable<BstructAPI> {
    return this.getBstruct(id, Name, frontRepo)
  }
  getBstruct(id: number, Name: string, frontRepo: FrontRepo): Observable<BstructAPI> {

    let params = new HttpParams().set("Name", Name)

    const url = `${this.bstructsUrl}/${id}`;
    return this.http.get<BstructAPI>(url, { params: params }).pipe(
      // tap(_ => this.log(`fetched bstruct id=${id}`)),
      catchError(this.handleError<BstructAPI>(`getBstruct id=${id}`))
    );
  }

  // postFront copy bstruct to a version with encoded pointers and post to the back
  postFront(bstruct: Bstruct, Name: string): Observable<BstructAPI> {
    let bstructAPI = new BstructAPI
    CopyBstructToBstructAPI(bstruct, bstructAPI)
    const id = typeof bstructAPI === 'number' ? bstructAPI : bstructAPI.ID
    const url = `${this.bstructsUrl}/${id}`;
    let params = new HttpParams().set("Name", Name)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.post<BstructAPI>(url, bstructAPI, httpOptions).pipe(
      tap(_ => {
      }),
      catchError(this.handleError<BstructAPI>('postBstruct'))
    );
  }
  
  /** POST: add a new bstruct to the server */
  post(bstructdb: BstructAPI, Name: string, frontRepo: FrontRepo): Observable<BstructAPI> {
    return this.postBstruct(bstructdb, Name, frontRepo)
  }
  postBstruct(bstructdb: BstructAPI, Name: string, frontRepo: FrontRepo): Observable<BstructAPI> {

    let params = new HttpParams().set("Name", Name)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.post<BstructAPI>(this.bstructsUrl, bstructdb, httpOptions).pipe(
      tap(_ => {
        // this.log(`posted bstructdb id=${bstructdb.ID}`)
      }),
      catchError(this.handleError<BstructAPI>('postBstruct'))
    );
  }

  /** DELETE: delete the bstructdb from the server */
  delete(bstructdb: BstructAPI | number, Name: string): Observable<BstructAPI> {
    return this.deleteBstruct(bstructdb, Name)
  }
  deleteBstruct(bstructdb: BstructAPI | number, Name: string): Observable<BstructAPI> {
    const id = typeof bstructdb === 'number' ? bstructdb : bstructdb.ID;
    const url = `${this.bstructsUrl}/${id}`;

    let params = new HttpParams().set("Name", Name)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.delete<BstructAPI>(url, httpOptions).pipe(
      tap(_ => this.log(`deleted bstructdb id=${id}`)),
      catchError(this.handleError<BstructAPI>('deleteBstruct'))
    );
  }

  // updateFront copy bstruct to a version with encoded pointers and update to the back
  updateFront(bstruct: Bstruct, Name: string): Observable<BstructAPI> {
    let bstructAPI = new BstructAPI
    CopyBstructToBstructAPI(bstruct, bstructAPI)
    const id = typeof bstructAPI === 'number' ? bstructAPI : bstructAPI.ID
    const url = `${this.bstructsUrl}/${id}`;
    let params = new HttpParams().set("Name", Name)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.put<BstructAPI>(url, bstructAPI, httpOptions).pipe(
      tap(_ => {
      }),
      catchError(this.handleError<BstructAPI>('updateBstruct'))
    );
  }

  /** PUT: update the bstructdb on the server */
  update(bstructdb: BstructAPI, Name: string, frontRepo: FrontRepo): Observable<BstructAPI> {
    return this.updateBstruct(bstructdb, Name, frontRepo)
  }
  updateBstruct(bstructdb: BstructAPI, Name: string, frontRepo: FrontRepo): Observable<BstructAPI> {
    const id = typeof bstructdb === 'number' ? bstructdb : bstructdb.ID;
    const url = `${this.bstructsUrl}/${id}`;


    let params = new HttpParams().set("Name", Name)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.put<BstructAPI>(url, bstructdb, httpOptions).pipe(
      tap(_ => {
        // this.log(`updated bstructdb id=${bstructdb.ID}`)
      }),
      catchError(this.handleError<BstructAPI>('updateBstruct'))
    );
  }

  /**
   * Handle Http operation that failed.
   * Let the app continue.
   * @param operation - name of the operation that failed
   * @param result - optional value to return as the observable result
   */
  private handleError<T>(operation = 'operation in BstructService', result?: T) {
    return (error: any): Observable<T> => {

      // TODO: send the error to remote logging infrastructure
      console.error("BstructService" + error); // log to console instead

      // TODO: better job of transforming error for user consumption
      this.log(`${operation} failed: ${error.message}`);

      // Let the app keep running by returning an empty result.
      return of(result as T);
    };
  }

  private log(message: string) {
    console.log(message)
  }
}
