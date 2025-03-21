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

import { GongEnumAPI } from './gongenum-api'
import { GongEnum, CopyGongEnumToGongEnumAPI } from './gongenum'

import { FrontRepo, FrontRepoService } from './front-repo.service';

// insertion point for imports
import { GongEnumValueAPI } from './gongenumvalue-api'

@Injectable({
  providedIn: 'root'
})
export class GongEnumService {

  // Kamar Raïmo: Adding a way to communicate between components that share information
  // so that they are notified of a change.
  GongEnumServiceChanged: BehaviorSubject<string> = new BehaviorSubject("");

  private gongenumsUrl: string

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
    this.gongenumsUrl = origin + '/api/github.com/fullstack-lang/gong/go/v1/gongenums';
  }

  /** GET gongenums from the server */
  // gets is more robust to refactoring
  gets(Name: string, frontRepo: FrontRepo): Observable<GongEnumAPI[]> {
    return this.getGongEnums(Name, frontRepo)
  }
  getGongEnums(Name: string, frontRepo: FrontRepo): Observable<GongEnumAPI[]> {

    let params = new HttpParams().set("Name", Name)

    return this.http.get<GongEnumAPI[]>(this.gongenumsUrl, { params: params })
      .pipe(
        tap(),
        catchError(this.handleError<GongEnumAPI[]>('getGongEnums', []))
      );
  }

  /** GET gongenum by id. Will 404 if id not found */
  // more robust API to refactoring
  get(id: number, Name: string, frontRepo: FrontRepo): Observable<GongEnumAPI> {
    return this.getGongEnum(id, Name, frontRepo)
  }
  getGongEnum(id: number, Name: string, frontRepo: FrontRepo): Observable<GongEnumAPI> {

    let params = new HttpParams().set("Name", Name)

    const url = `${this.gongenumsUrl}/${id}`;
    return this.http.get<GongEnumAPI>(url, { params: params }).pipe(
      // tap(_ => this.log(`fetched gongenum id=${id}`)),
      catchError(this.handleError<GongEnumAPI>(`getGongEnum id=${id}`))
    );
  }

  // postFront copy gongenum to a version with encoded pointers and post to the back
  postFront(gongenum: GongEnum, Name: string): Observable<GongEnumAPI> {
    let gongenumAPI = new GongEnumAPI
    CopyGongEnumToGongEnumAPI(gongenum, gongenumAPI)
    const id = typeof gongenumAPI === 'number' ? gongenumAPI : gongenumAPI.ID
    const url = `${this.gongenumsUrl}/${id}`;
    let params = new HttpParams().set("Name", Name)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.post<GongEnumAPI>(url, gongenumAPI, httpOptions).pipe(
      tap(_ => {
      }),
      catchError(this.handleError<GongEnumAPI>('postGongEnum'))
    );
  }
  
  /** POST: add a new gongenum to the server */
  post(gongenumdb: GongEnumAPI, Name: string, frontRepo: FrontRepo): Observable<GongEnumAPI> {
    return this.postGongEnum(gongenumdb, Name, frontRepo)
  }
  postGongEnum(gongenumdb: GongEnumAPI, Name: string, frontRepo: FrontRepo): Observable<GongEnumAPI> {

    let params = new HttpParams().set("Name", Name)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.post<GongEnumAPI>(this.gongenumsUrl, gongenumdb, httpOptions).pipe(
      tap(_ => {
        // this.log(`posted gongenumdb id=${gongenumdb.ID}`)
      }),
      catchError(this.handleError<GongEnumAPI>('postGongEnum'))
    );
  }

  /** DELETE: delete the gongenumdb from the server */
  delete(gongenumdb: GongEnumAPI | number, Name: string): Observable<GongEnumAPI> {
    return this.deleteGongEnum(gongenumdb, Name)
  }
  deleteGongEnum(gongenumdb: GongEnumAPI | number, Name: string): Observable<GongEnumAPI> {
    const id = typeof gongenumdb === 'number' ? gongenumdb : gongenumdb.ID;
    const url = `${this.gongenumsUrl}/${id}`;

    let params = new HttpParams().set("Name", Name)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.delete<GongEnumAPI>(url, httpOptions).pipe(
      tap(_ => this.log(`deleted gongenumdb id=${id}`)),
      catchError(this.handleError<GongEnumAPI>('deleteGongEnum'))
    );
  }

  // updateFront copy gongenum to a version with encoded pointers and update to the back
  updateFront(gongenum: GongEnum, Name: string): Observable<GongEnumAPI> {
    let gongenumAPI = new GongEnumAPI
    CopyGongEnumToGongEnumAPI(gongenum, gongenumAPI)
    const id = typeof gongenumAPI === 'number' ? gongenumAPI : gongenumAPI.ID
    const url = `${this.gongenumsUrl}/${id}`;
    let params = new HttpParams().set("Name", Name)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.put<GongEnumAPI>(url, gongenumAPI, httpOptions).pipe(
      tap(_ => {
      }),
      catchError(this.handleError<GongEnumAPI>('updateGongEnum'))
    );
  }

  /** PUT: update the gongenumdb on the server */
  update(gongenumdb: GongEnumAPI, Name: string, frontRepo: FrontRepo): Observable<GongEnumAPI> {
    return this.updateGongEnum(gongenumdb, Name, frontRepo)
  }
  updateGongEnum(gongenumdb: GongEnumAPI, Name: string, frontRepo: FrontRepo): Observable<GongEnumAPI> {
    const id = typeof gongenumdb === 'number' ? gongenumdb : gongenumdb.ID;
    const url = `${this.gongenumsUrl}/${id}`;


    let params = new HttpParams().set("Name", Name)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.put<GongEnumAPI>(url, gongenumdb, httpOptions).pipe(
      tap(_ => {
        // this.log(`updated gongenumdb id=${gongenumdb.ID}`)
      }),
      catchError(this.handleError<GongEnumAPI>('updateGongEnum'))
    );
  }

  /**
   * Handle Http operation that failed.
   * Let the app continue.
   * @param operation - name of the operation that failed
   * @param result - optional value to return as the observable result
   */
  private handleError<T>(operation = 'operation in GongEnumService', result?: T) {
    return (error: any): Observable<T> => {

      // TODO: send the error to remote logging infrastructure
      console.error("GongEnumService" + error); // log to console instead

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
