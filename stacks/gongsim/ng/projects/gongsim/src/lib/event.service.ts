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

import { EventAPI } from './event-api';
import { EventDB } from './event-db';

@Injectable({
  providedIn: 'root'
})
export class EventService {

  httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' })
  };

  // Kamar Raïmo: Adding a way to communicate between components that share information
  // so that they are notified of a change.
  EventServiceChanged: BehaviorSubject<string> = new BehaviorSubject("");

  private eventsUrl = 'http://localhost:8080/api/github.com/fullstack-lang/gong/stacks/gongsim/go/v1/events';

  constructor(
    private http: HttpClient
  ) { }

  /** GET events from the server */
  getEvents(): Observable<EventDB[]> {
    return this.http.get<EventDB[]>(this.eventsUrl)
      .pipe(
        tap(_ => this.log('fetched events')),
        catchError(this.handleError<EventDB[]>('getEvents', []))
      );
  }

  /** GET event by id. Will 404 if id not found */
  getEvent(id: number): Observable<EventDB> {
    const url = `${this.eventsUrl}/${id}`;
    return this.http.get<EventDB>(url).pipe(
      tap(_ => this.log(`fetched event id=${id}`)),
      catchError(this.handleError<EventDB>(`getEvent id=${id}`))
    );
  }

  //////// Save methods //////////

  /** POST: add a new event to the server */
  postEvent(eventAPI: EventAPI): Observable<EventDB> {
    return this.http.post<EventDB>(this.eventsUrl, eventAPI, this.httpOptions).pipe(
      tap((newEvent: EventDB) => {})
    );
  }

  /** DELETE: delete the eventdb from the server */
  deleteEvent(eventdb: EventDB | number): Observable<EventDB> {
    const id = typeof eventdb === 'number' ? eventdb : eventdb.ID;
    const url = `${this.eventsUrl}/${id}`;

    return this.http.delete<EventDB>(url, this.httpOptions).pipe(
      tap(_ => this.log(`deleted eventdb id=${id}`)),
      catchError(this.handleError<EventDB>('deleteEvent'))
    );
  }

  /** PUT: update the eventdb on the server */
  updateEvent(eventdb: EventDB): Observable<EventDB> {
    const id = typeof eventdb === 'number' ? eventdb : eventdb.ID;
    const url = `${this.eventsUrl}/${id}`;

    // insertion point for reset of reverse pointers (to avoid circular JSON)

    return this.http.put(url, eventdb, this.httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers
        this.log(`updated eventdb id=${eventdb.ID}`)
      }),
      catchError(this.handleError<EventDB>('updateEvent'))
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