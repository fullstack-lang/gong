import { Component, Input, OnInit } from '@angular/core';
import { Router } from '@angular/router';

import * as gongrouter from 'gongrouter'
import { Subscription } from 'rxjs';


@Component({
  selector: 'lib-gongrouter-outlet',
  templateUrl: './gongrouter-outlet.component.html',
  styleUrls: ['./gongrouter-outlet.component.css']
})
export class GongrouterOutletComponent {

  @Input() DataStack: string = ""
  @Input() OutletName: string = ""

  // the component is refreshed when modification are performed in the back repo 
  // 
  // the checkCommitNbFromBackTimer polls the commit number of the back repo
  // if the commit number has increased, it pulls the front repo and redraw the diagram
  private commutNbFromBackSubscription: Subscription = new Subscription
  lastCommitNbFromBack = -1
  lastPushFromFrontNb = -1
  currTime: number = 0
  dateOfLastTimerEmission: Date = new Date


  public gongrouterFrontRepo?: gongrouter.FrontRepo

  constructor(
    private router: Router,
    private routeService: gongrouter.RouteService,
    private gongrouterFrontRepoService: gongrouter.FrontRepoService,
    private gongrouterCommitNbFromBackService: gongrouter.CommitNbFromBackService,
  ) {

  }

  ngOnInit(): void {
    this.startAutoRefresh(500); // Refresh every 500 ms (half second)
  }

  ngOnDestroy(): void {
    this.stopAutoRefresh();
  }


  stopAutoRefresh(): void {
    if (this.commutNbFromBackSubscription) {
      this.commutNbFromBackSubscription.unsubscribe();
    }
  }

  startAutoRefresh(intervalMs: number): void {
    this.commutNbFromBackSubscription = this.gongrouterCommitNbFromBackService
      .getCommitNbFromBack(intervalMs, this.DataStack)
      .subscribe((commitNbFromBack: number) => {
        // console.log("OutletComponent, last commit nb " + this.lastCommitNbFromBack + " new: " + commitNbFromBack)

        if (this.lastCommitNbFromBack < commitNbFromBack) {
          const d = new Date()
          console.log("OutletComponent. OutletName: " + this.OutletName + ", Stack: " + this.DataStack, d.toLocaleTimeString() + `.${d.getMilliseconds()}` +
            ", last commit increased nb " + this.lastCommitNbFromBack + " new: " + commitNbFromBack)
          this.lastCommitNbFromBack = commitNbFromBack
          this.refresh()
        }
      }
      )
  }

  refresh(): void {

    this.gongrouterFrontRepoService.pull(this.DataStack).subscribe(
      gongroutersFrontRepo => {
        this.gongrouterFrontRepo = gongroutersFrontRepo

        var outlet: gongrouter.OutletDB = new (gongrouter.OutletDB)
        var selected: boolean = false
        for (var outlet_ of this.gongrouterFrontRepo.Outlets_array) {
          if (outlet_.Name == this.OutletName) {
            outlet = outlet_
            selected = true
          }
        }
        if (!selected) {
          console.log("no outlet selected")
          return
        }

        if (outlet.Path == "") {
          return
        }
        this.setOutlet(outlet.Name, outlet.Path)

      }
    )
  }

  setOutlet(outletName: string, path: string) {
    let outletConf: any = {}

    outletConf[outletName] = [path, this.DataStack]

    this.router.navigate([{ outlets: outletConf }]);

  }

}
