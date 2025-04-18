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

import { LinkAnchoredTextAPI } from './linkanchoredtext-api'
import { LinkAnchoredText, CopyLinkAnchoredTextToLinkAnchoredTextAPI } from './linkanchoredtext'

import { FrontRepo, FrontRepoService } from './front-repo.service';

// insertion point for imports
import { AnimateAPI } from './animate-api'

@Injectable({
  providedIn: 'root'
})
export class LinkAnchoredTextService {

  // Kamar Raïmo: Adding a way to communicate between components that share information
  // so that they are notified of a change.
  LinkAnchoredTextServiceChanged: BehaviorSubject<string> = new BehaviorSubject("");

  private linkanchoredtextsUrl: string

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
    this.linkanchoredtextsUrl = origin + '/api/github.com/fullstack-lang/gong/lib/svg/go/v1/linkanchoredtexts';
  }

  /** GET linkanchoredtexts from the server */
  // gets is more robust to refactoring
  gets(Name: string, frontRepo: FrontRepo): Observable<LinkAnchoredTextAPI[]> {
    return this.getLinkAnchoredTexts(Name, frontRepo)
  }
  getLinkAnchoredTexts(Name: string, frontRepo: FrontRepo): Observable<LinkAnchoredTextAPI[]> {

    let params = new HttpParams().set("Name", Name)

    return this.http.get<LinkAnchoredTextAPI[]>(this.linkanchoredtextsUrl, { params: params })
      .pipe(
        tap(),
        catchError(this.handleError<LinkAnchoredTextAPI[]>('getLinkAnchoredTexts', []))
      );
  }

  /** GET linkanchoredtext by id. Will 404 if id not found */
  // more robust API to refactoring
  get(id: number, Name: string, frontRepo: FrontRepo): Observable<LinkAnchoredTextAPI> {
    return this.getLinkAnchoredText(id, Name, frontRepo)
  }
  getLinkAnchoredText(id: number, Name: string, frontRepo: FrontRepo): Observable<LinkAnchoredTextAPI> {

    let params = new HttpParams().set("Name", Name)

    const url = `${this.linkanchoredtextsUrl}/${id}`;
    return this.http.get<LinkAnchoredTextAPI>(url, { params: params }).pipe(
      // tap(_ => this.log(`fetched linkanchoredtext id=${id}`)),
      catchError(this.handleError<LinkAnchoredTextAPI>(`getLinkAnchoredText id=${id}`))
    );
  }

  // postFront copy linkanchoredtext to a version with encoded pointers and post to the back
  postFront(linkanchoredtext: LinkAnchoredText, Name: string): Observable<LinkAnchoredTextAPI> {
    let linkanchoredtextAPI = new LinkAnchoredTextAPI
    CopyLinkAnchoredTextToLinkAnchoredTextAPI(linkanchoredtext, linkanchoredtextAPI)
    const id = typeof linkanchoredtextAPI === 'number' ? linkanchoredtextAPI : linkanchoredtextAPI.ID
    const url = `${this.linkanchoredtextsUrl}/${id}`;
    let params = new HttpParams().set("Name", Name)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.post<LinkAnchoredTextAPI>(url, linkanchoredtextAPI, httpOptions).pipe(
      tap(_ => {
      }),
      catchError(this.handleError<LinkAnchoredTextAPI>('postLinkAnchoredText'))
    );
  }
  
  /** POST: add a new linkanchoredtext to the server */
  post(linkanchoredtextdb: LinkAnchoredTextAPI, Name: string, frontRepo: FrontRepo): Observable<LinkAnchoredTextAPI> {
    return this.postLinkAnchoredText(linkanchoredtextdb, Name, frontRepo)
  }
  postLinkAnchoredText(linkanchoredtextdb: LinkAnchoredTextAPI, Name: string, frontRepo: FrontRepo): Observable<LinkAnchoredTextAPI> {

    let params = new HttpParams().set("Name", Name)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.post<LinkAnchoredTextAPI>(this.linkanchoredtextsUrl, linkanchoredtextdb, httpOptions).pipe(
      tap(_ => {
        // this.log(`posted linkanchoredtextdb id=${linkanchoredtextdb.ID}`)
      }),
      catchError(this.handleError<LinkAnchoredTextAPI>('postLinkAnchoredText'))
    );
  }

  /** DELETE: delete the linkanchoredtextdb from the server */
  delete(linkanchoredtextdb: LinkAnchoredTextAPI | number, Name: string): Observable<LinkAnchoredTextAPI> {
    return this.deleteLinkAnchoredText(linkanchoredtextdb, Name)
  }
  deleteLinkAnchoredText(linkanchoredtextdb: LinkAnchoredTextAPI | number, Name: string): Observable<LinkAnchoredTextAPI> {
    const id = typeof linkanchoredtextdb === 'number' ? linkanchoredtextdb : linkanchoredtextdb.ID;
    const url = `${this.linkanchoredtextsUrl}/${id}`;

    let params = new HttpParams().set("Name", Name)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.delete<LinkAnchoredTextAPI>(url, httpOptions).pipe(
      tap(_ => this.log(`deleted linkanchoredtextdb id=${id}`)),
      catchError(this.handleError<LinkAnchoredTextAPI>('deleteLinkAnchoredText'))
    );
  }

  // updateFront copy linkanchoredtext to a version with encoded pointers and update to the back
  updateFront(linkanchoredtext: LinkAnchoredText, Name: string): Observable<LinkAnchoredTextAPI> {
    let linkanchoredtextAPI = new LinkAnchoredTextAPI
    CopyLinkAnchoredTextToLinkAnchoredTextAPI(linkanchoredtext, linkanchoredtextAPI)
    const id = typeof linkanchoredtextAPI === 'number' ? linkanchoredtextAPI : linkanchoredtextAPI.ID
    const url = `${this.linkanchoredtextsUrl}/${id}`;
    let params = new HttpParams().set("Name", Name)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.put<LinkAnchoredTextAPI>(url, linkanchoredtextAPI, httpOptions).pipe(
      tap(_ => {
      }),
      catchError(this.handleError<LinkAnchoredTextAPI>('updateLinkAnchoredText'))
    );
  }

  /** PUT: update the linkanchoredtextdb on the server */
  update(linkanchoredtextdb: LinkAnchoredTextAPI, Name: string, frontRepo: FrontRepo): Observable<LinkAnchoredTextAPI> {
    return this.updateLinkAnchoredText(linkanchoredtextdb, Name, frontRepo)
  }
  updateLinkAnchoredText(linkanchoredtextdb: LinkAnchoredTextAPI, Name: string, frontRepo: FrontRepo): Observable<LinkAnchoredTextAPI> {
    const id = typeof linkanchoredtextdb === 'number' ? linkanchoredtextdb : linkanchoredtextdb.ID;
    const url = `${this.linkanchoredtextsUrl}/${id}`;


    let params = new HttpParams().set("Name", Name)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.put<LinkAnchoredTextAPI>(url, linkanchoredtextdb, httpOptions).pipe(
      tap(_ => {
        // this.log(`updated linkanchoredtextdb id=${linkanchoredtextdb.ID}`)
      }),
      catchError(this.handleError<LinkAnchoredTextAPI>('updateLinkAnchoredText'))
    );
  }

  /**
   * Handle Http operation that failed.
   * Let the app continue.
   * @param operation - name of the operation that failed
   * @param result - optional value to return as the observable result
   */
  private handleError<T>(operation = 'operation in LinkAnchoredTextService', result?: T) {
    return (error: any): Observable<T> => {

      // TODO: send the error to remote logging infrastructure
      console.error("LinkAnchoredTextService" + error); // log to console instead

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
