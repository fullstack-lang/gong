package angular

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/fullstack-lang/gong/go/models"
)

const NgCommitNbFromBackTemplateTS = `import { Injectable, Component, Inject } from '@angular/core';
import { HttpClientModule, HttpParams } from '@angular/common/http';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { DOCUMENT, Location } from '@angular/common'

/*
 * Behavior subject
 */
import { BehaviorSubject, interval } from 'rxjs'
import { Observable, of } from 'rxjs'
import { catchError, map, switchMap, tap } from 'rxjs/operators'

@Injectable({
    providedIn: 'root'
})
export class CommitNbFromBackService {

    httpOptions = {
        headers: new HttpHeaders({ 'Content-Type': 'application/json' })
    };

    private commitNbFromBackUrl: string

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
        this.commitNbFromBackUrl = origin + '/api/{{PkgPathRoot}}/v1/commitfrombacknb';
    }

    getCommitNbFromBack(intervalMs: number, GONG__StackPath: string = ""): Observable<number> {

        let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

        return interval(intervalMs).pipe(
            switchMap(() => this.http.get<number>(this.commitNbFromBackUrl, { params: params }).pipe(
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
    private handleError<T>(operation = 'operation in CommitNbFromBackService', result?: T) {
        return (error: any): Observable<T> => {

            // TODO: send the error to remote logging infrastructure
            console.error("in CommitNbFromBackService" + error); // log to console instead

            // TODO: better job of transforming error for user consumption
            this.log('${operation} failed: ${error.message}');

            // Let the app keep running by returning an empty result.
            return of(result as T);
        };
    }

    private log(message: string) {
        console.log(message)
    }
}
`

// MultiCodeGeneratorNgCommitNb parses mdlPkg and generates the code for the
// CommitNb components
func CodeGeneratorNgCommitNbFromBack(
	mdlPkg *models.ModelPkg,
	pkgName string,
	matTargetPath string,
	pkgGoPath string,
	apiPath string) {

	codeTS := NgCommitNbFromBackTemplateTS
	codeTS = strings.ReplaceAll(codeTS, "{{addr}}", apiPath)

	// final replacement
	codeTS = models.Replace4(codeTS,
		"{{PkgName}}", mdlPkg.Name,
		"{{TitlePkgName}}", strings.Title(mdlPkg.Name),
		"{{pkgname}}", strings.ToLower(mdlPkg.Name),
		"{{PkgPathRoot}}", strings.ReplaceAll(mdlPkg.PkgPath, "/models", ""))

	file, err := os.Create(filepath.Join(matTargetPath, "commitnbfromback.service.ts"))
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()
	fmt.Fprint(file, codeTS)

}
