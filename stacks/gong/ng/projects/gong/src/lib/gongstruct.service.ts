// generated by MultiCodeGeneratorNgService
import { Injectable } from '@angular/core';
import { HttpClientModule } from '@angular/common/http';
import { HttpClient, HttpHeaders } from '@angular/common/http';

/*
 * Behavior subject
 */
import { BehaviorSubject } from 'rxjs';
import { Observable, of } from 'rxjs';
import { catchError, map, tap } from 'rxjs/operators';

import { GongStructAPI } from './gongstruct-api';
import { GongStructDB } from './gongstruct-db';

@Injectable({
  providedIn: 'root'
})
export class GongStructService {

  httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' })
  };

  // Kamar Raïmo: Adding a way to communicate between components that share information
  // so that they are notified of a change.
  GongStructServiceChanged: BehaviorSubject<string> = new BehaviorSubject("");

  private gongstructsUrl = 'http://localhost:8080/api/github.com/fullstack-lang/gong/stacks/gong/go/v1/gongstructs';

  constructor(
    private http: HttpClient
  ) { }

  /** GET gongstructs from the server */
  getGongStructs(): Observable<GongStructDB[]> {
    return this.http.get<GongStructDB[]>(this.gongstructsUrl)
      .pipe(
        tap(_ => this.log('fetched gongstructs')),
        catchError(this.handleError<GongStructDB[]>('getGongStructs', []))
      );
  }

  /** GET gongstruct by id. Will 404 if id not found */
  getGongStruct(id: number): Observable<GongStructDB> {
    const url = `${this.gongstructsUrl}/${id}`;
    return this.http.get<GongStructDB>(url).pipe(
      tap(_ => this.log(`fetched gongstruct id=${id}`)),
      catchError(this.handleError<GongStructDB>(`getGongStruct id=${id}`))
    );
  }

  //////// Save methods //////////

  /** POST: add a new gongstruct to the server */
  postGongStruct(gongstructAPI: GongStructAPI): Observable<GongStructDB> {
    return this.http.post<GongStructDB>(this.gongstructsUrl, gongstructAPI, this.httpOptions).pipe(
      tap((newGongStruct: GongStructDB) => {})
    );
  }

  /** DELETE: delete the gongstructdb from the server */
  deleteGongStruct(gongstructdb: GongStructDB | number): Observable<GongStructDB> {
    const id = typeof gongstructdb === 'number' ? gongstructdb : gongstructdb.ID;
    const url = `${this.gongstructsUrl}/${id}`;

    return this.http.delete<GongStructDB>(url, this.httpOptions).pipe(
      tap(_ => this.log(`deleted gongstructdb id=${id}`)),
      catchError(this.handleError<GongStructDB>('deleteGongStruct'))
    );
  }

  /** PUT: update the gongstructdb on the server */
  updateGongStruct(gongstructdb: GongStructDB): Observable<GongStructDB> {
    const id = typeof gongstructdb === 'number' ? gongstructdb : gongstructdb.ID;
    const url = `${this.gongstructsUrl}/${id}`;

    // insertion point for reset of reverse pointers (to avoid circular JSON)
    gongstructdb.GongBasicFields = []
    gongstructdb.PointerToGongStructFields = []
    gongstructdb.SliceOfPointerToGongStructFields = []

    return this.http.put(url, gongstructdb, this.httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers
        this.log(`updated gongstructdb id=${gongstructdb.ID}`)
      }),
      catchError(this.handleError<GongStructDB>('updateGongStruct'))
    );
  }

  /**
   * Handle Http operation that failed.
   * Let the app continue.
   * @param operation - name of the operation that failed
   * @param result - optional value to return as the observable result
   */
  private handleError<T>(operation = 'operation', result?: T) {
    return (error: any): Observable<T> => {

      // TODO: send the error to remote logging infrastructure
      console.error(error); // log to console instead

      // TODO: better job of transforming error for user consumption
      this.log(`${operation} failed: ${error.message}`);

      // Let the app keep running by returning an empty result.
      return of(result as T);
    };
  }

  private log(message: string) {

  }
}