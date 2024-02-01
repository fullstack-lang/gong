import { DOCUMENT } from '@angular/common';
import { HttpParams } from '@angular/common/http';
import { Component, Inject, Input, OnInit } from '@angular/core';
import { ErrorObserver, Observable, timer } from 'rxjs';

import { webSocket, WebSocketSubject } from 'rxjs/webSocket'

import * as test from 'test'
import { WebSocketService } from './web-socket-service';

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
    @Inject(DOCUMENT) private document: Document,
    private webSocketService: WebSocketService

  ) {

  }
  ngOnInit(): void {

    console.log("Material component->ngOnInit : GONG__StackPath, " + this.GONG__StackPath)


    if (this.wsSubject == undefined) {
      console.log("webSocket(url) failed")
    }

    this.webSocketService.connect(this.GONG__StackPath).subscribe(data => {
      const response = JSON.parse(data.data) as test.BackRepoData
      console.log(response)
    })


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
