import { Component, Input, OnInit } from '@angular/core';
import { Observable, timer } from 'rxjs';

import * as test from 'test'


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

  // the component is refreshed when modification are performed in the back repo 
  // the checkCommiNbFromBagetCommitNbFromBackTimer polls the commit number of the back repo
  // if the commit number has increased, it pulls the front repo and redraw the diagram
  checkCommiNbFromBagetCommitNbFromBackTimer: Observable<number> = timer(500, 500);
  lastCommitNbFromBack = -1
  lastPushFromFrontNb = -1
  currTime: number = 0

  constructor(
    public frontRepoService: test.FrontRepoService,
    private astructService: test.AstructService,
    private commitNbFromBackService: test.CommitNbFromBackService,

  ) {

  }
  ngOnInit(): void {

    console.log("Material component->ngOnInit : GONG__StackPath, " + this.GONG__StackPath)


    // see above for the explanation
    this.commitNbFromBackService.getCommitNbFromBack(500, this.GONG__StackPath).subscribe(
      commiNbFromBagetCommitNbFromBack => {
        if (this.lastCommitNbFromBack < commiNbFromBagetCommitNbFromBack) {

          console.log("last commit nb " + this.lastCommitNbFromBack + " new: " + commiNbFromBagetCommitNbFromBack)
          this.lastCommitNbFromBack = commiNbFromBagetCommitNbFromBack

          this.frontRepoService.pull(this.GONG__StackPath).subscribe(
            frontRepo => {
              this.frontRepo = frontRepo
              let astructs = this.frontRepo.getFrontArray<test.Astruct>(test.Astruct.GONGSTRUCT_NAME)
              console.log("Nb of Astruct is ", astructs.length)

              let astruct = astructs[0]

              this.astructService.updateFront(astruct, this.GONG__StackPath).subscribe()
            }
          )
        }
      }
    )
  }
}
