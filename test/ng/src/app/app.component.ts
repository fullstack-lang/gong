import { Component, OnInit } from '@angular/core';

import { Observable, combineLatest, timer } from 'rxjs'

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
})
export class AppComponent implements OnInit {

  view = 'Default view'
  default = 'Default view'
  diagrams = 'Diagrams view'
  meta = 'Meta view'

  views: string[] = [this.default, this.diagrams, this.meta];

  obsTimer: Observable<number> = timer(1000, 1000)

  ngOnInit(): void {

    this.obsTimer.subscribe(
      currTime => {
        console.log(currTime)
      }
    )

  }
}
