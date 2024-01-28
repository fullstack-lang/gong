import { DOCUMENT } from '@angular/common';
import { HttpParams } from '@angular/common/http';
import { Component, Inject, Input, OnInit } from '@angular/core';
import { ErrorObserver, Observable, timer } from 'rxjs';

import { webSocket, WebSocketSubject } from 'rxjs/webSocket'

import * as test from 'test'

// Create an error handling function
const handleError: ErrorObserver<Event> = {
  error: (err: Event) => {
    console.error('WebSocket error:', err)
    // Display the error message
    const message = `WebSocket error: ${err}`
    // Handle the error display logic here, e.g., show an alert, update a variable to show the error on the UI, etc.
  },
}
@Component({
  selector: 'lib-testspecific',
  template: `
    <p>
      testspecific works!
    </p>
  `,
  styles: [
  ]
})
export class TestspecificComponent implements OnInit {

  @Input() GONG__StackPath: string = ""

  frontRepo = new test.FrontRepo

  private wsSubject: WebSocketSubject<any> | undefined

  // the component is refreshed when modification are performed in the back repo 
  // the checkCommiNbFromBagetCommitNbFromBackTimer polls the commit number of the back repo
  // if the commit number has increased, it pulls the front repo and redraw the diagram
  checkCommiNbFromBagetCommitNbFromBackTimer: Observable<number> = timer(500, 500);
  lastCommitNbFromBack = -1
  lastPushFromFrontNb = -1
  currTime: number = 0

  constructor(
    private frontRepoService: test.FrontRepoService,
    private astructService: test.AstructService,
    private commitNbFromBackService: test.CommitNbFromBackService,
    @Inject(DOCUMENT) private document: Document

  ) {

  }
  ngOnInit(): void {

    console.log("Material component->ngOnInit : GONG__StackPath, " + this.GONG__StackPath)

    let params = new HttpParams().set("GONG__StackPath", this.GONG__StackPath)

    let origin = this.document.location.origin

    // if debugging with ng, replace 4200 with 8080
    origin = origin.replace("4200", "8080")

    // compute path to the service
    let basePath = origin + '/api/github.com/fullstack-lang/gong/test/go/v1/ws';

    // Add params to the WebSocket URL
    // let paramString = params.toString(); // Convert HttpParams to string
    // let url = `${basePath}?${paramString}`;

    basePath = 'ws://localhost:8080/api/github.com/fullstack-lang/gong/test/go/v1/ws'
    let paramString = params.toString()
    let url = `${basePath}?${paramString}`
    this.wsSubject = webSocket(url)

    if (this.wsSubject != undefined) {
      this.wsSubject.subscribe({
        next: (data) => {
          let nbCommit = parseInt(data, 10)
          console.log("test specific", nbCommit)
        },
        error: handleError.error,
      });
    }


    // see above for the explanation
    // this.commitNbFromBackService.getCommitNbFromBack(500, this.GONG__StackPath).subscribe(
    //   commiNbFromBagetCommitNbFromBack => {
    //     if (this.lastCommitNbFromBack < commiNbFromBagetCommitNbFromBack) {

    //       console.log("last commit nb " + this.lastCommitNbFromBack + " new: " + commiNbFromBagetCommitNbFromBack)
    //       this.lastCommitNbFromBack = commiNbFromBagetCommitNbFromBack

    //       this.frontRepoService.pull(this.GONG__StackPath).subscribe(
    //         frontRepo => {
    //           this.frontRepo = frontRepo
    //           let astructs = this.frontRepo.getFrontArray<test.Astruct>(test.Astruct.GONGSTRUCT_NAME)
    //           console.log("Nb of Astruct is ", astructs.length)

    //           let astruct = astructs[0]

    //           this.astructService.updateFront(astruct, this.GONG__StackPath).subscribe()
    //         }
    //       )
    //     }
    //   }
    // )
  }
}
