import { Injectable, Component, Inject } from '@angular/core';
import { HttpClientModule } from '@angular/common/http';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { DOCUMENT, Location } from '@angular/common'

/*
 * Behavior subject
 */
import { BehaviorSubject } from 'rxjs';
import { Observable, of } from 'rxjs';
import { catchError, map, tap } from 'rxjs/operators';

@Injectable({
    providedIn: 'root'
})
export class CommitNbService {

    httpOptions = {
        headers: new HttpHeaders({ 'Content-Type': 'application/json' })
    };

    private commitNbUrl: string

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
        this.commitNbUrl = origin + '/api/github.com/fullstack-lang/gongdoc/go/commitfrombacknb';
    }

    // observable of the commit nb getter
    public getCommitNb(): Observable<number> {
        return this.http.get<number>(this.commitNbUrl)
            .pipe(
                tap(_ => this.log('fetched commit nb')),
                catchError(this.handleError<number>('getCommitNb', -1))
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
            this.log('${operation} failed: ${error.message}');

            // Let the app keep running by returning an empty result.
            return of(result as T);
        };
    }

    private log(message: string) {

    }
}
