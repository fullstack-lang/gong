import { Component, Input, OnDestroy, OnInit } from '@angular/core';
import { Subscription } from 'rxjs';

import * as test2 from 'test2'

@Component({
  selector: 'lib-test2specific',
  templateUrl: './test2specific.component.html',
  styles: [
  ]
})
export class Test2specificComponent implements OnInit, OnDestroy {

  @Input() name: string = ""
  @Input() GONG__StackPath: string = ""

  private commutNbFromBackSubscription: Subscription = new Subscription
  lastCommitNbFromBack = -1
  lastPushFromFrontNb = -1
  currTime: number = 0
  dateOfLastTimerEmission: Date = new Date

  // 
  frontRepo: test2.FrontRepo = new (test2.FrontRepo)

  constructor(
    private frontRepoService: test2.FrontRepoService,
    private commitNbFromBackService: test2.CommitNbFromBackService,
    private aService: test2.AService,
    private bService: test2.BService,

  ) {

  }

  ngOnInit(): void {
    console.log("Component->name : ", this.name)
    console.log("Component->GONG__StackPath : ", this.GONG__StackPath)
    this.startAutoRefresh(500); // Refresh every 500 ms (half second)
  }

  ngOnDestroy(): void {
    this.stopAutoRefresh();
  }

  startAutoRefresh(intervalMs: number): void {

    this.commutNbFromBackSubscription = this.commitNbFromBackService
      .getCommitNbFromBack(intervalMs, this.GONG__StackPath)
      .subscribe((commitNbFromBack: number) => {
        // console.log("TreeComponent, last commit nb " + this.lastCommitNbFromBack + " new: " + commitNbFromBack)

        if (this.lastCommitNbFromBack < commitNbFromBack) {
          const d = new Date()
          console.log("Test2 Specific Component, ", this.GONG__StackPath, " name ", this.name + d.toLocaleTimeString() + `.${d.getMilliseconds()}` +
            ", last commit increased nb " + this.lastCommitNbFromBack + " new: " + commitNbFromBack)
          this.lastCommitNbFromBack = commitNbFromBack
          this.refresh()
        }
      }
      )
  }

  stopAutoRefresh(): void {
    if (this.commutNbFromBackSubscription) {
      this.commutNbFromBackSubscription.unsubscribe();
    }
  }

  refresh(): void {

    this.frontRepoService.pull(this.GONG__StackPath).subscribe(
      frontRepo => {
        this.frontRepo = frontRepo
      }
    )
  }

  onClickUpdateBasicField(A: test2.ADB) {
    A.NumberField = Math.floor(Math.random() * 100)

    this.aService.update(A, "test2", this.frontRepoService.frontRepo).subscribe(
      a => {
        console.log("a", a)
      }
    )
  }

  onClickUpdatePointerField(A: test2.ADB) {
    let rank = Math.floor(Math.random() * this.frontRepo.Bs_array.length)
    A.B = this.frontRepo.Bs_array[rank]

    this.aService.update(A, "test2", this.frontRepoService.frontRepo).subscribe()

    A.B = this.frontRepo.Bs_array[rank]
  }

  onClickUpdateSliceOfPointerField(A: test2.ADB) {
    shuffleBs(A)

    this.aService.update(A, "test2", this.frontRepoService.frontRepo).subscribe()
  }


}

function shuffleBs(A: test2.ADB) {
  for (let i = A.Bs.length - 1; i > 0; i--) {
    const j = Math.floor(Math.random() * (i + 1));
    [A.Bs[i], A.Bs[j]] = [A.Bs[j], A.Bs[i]]; // Swap elements
  }
}
