import { Injectable, Component, Inject } from '@angular/core';
import { HttpClientModule, HttpParams } from '@angular/common/http';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { DOCUMENT, Location } from '@angular/common'

/*
 * Behavior subject
 */
import { BehaviorSubject, interval } from 'rxjs';
import { Observable, of } from 'rxjs';
import { catchError, map, switchMap, tap } from 'rxjs/operators';

@Injectable({
    providedIn: 'root'
})
export class PushFromFrontNbService {

    httpOptions = {
        headers: new HttpHeaders({ 'Content-Type': 'application/json' })
    };

    private pushFromFrontNbURL: string
    constructor(
        private http: HttpClient,
        private location: Location,
        @Inject(DOCUMENT) private document: Document
    ) {
        // path to the service share the same origin with the path to the document
        // get the origin in the URL to the document
        let origin = this.document.location.origin

        // if debugging with ng, replace 4200 with 8080
        origin = origin.replace("4200", "8080")

        // compute path to the service
        this.pushFromFrontNbURL = origin + '/api/github.com/fullstack-lang/gongtree/go/v1/pushfromfrontnb';
    }

    // observable of the commit nb getter
    public getPushFromFrontNb(): Observable<number> {
        return this.http.get<number>(this.pushFromFrontNbURL)
            .pipe(
                tap(_ => this.log('fetched commit nb')),
                catchError(this.handleError<number>('getPushFromFrontNb', -1))
            );
    }

    getPushNbFromFront(intervalMs: number, GONG__StackPath: string = ""): Observable<number> {

        let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

        return interval(intervalMs).pipe(
            switchMap(() => this.http.get<number>(this.pushFromFrontNbURL, { params: params }).pipe(
                catchError(error => {
                    // Handle the error here, e.g. log it, show a notification, etc.
                    console.error('Error fetching commit number:', error);

                    // Return a default value, a new Observable, or rethrow the error
                    return of(0); // Here, we return 0 as a default value
                })
            )
            )
        )
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
            this.log('${operation} failed: ${error.message}');

            // Let the app keep running by returning an empty result.
            return of(result as T);
        };
    }

    private log(message: string) {

    }
}
