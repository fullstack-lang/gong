package angular

const NgFileAppComponentTs = `import { Component, OnInit } from '@angular/core';

import { Observable, combineLatest, timer } from 'rxjs'

import * as gongdoc from 'gongdoc'
import * as {{pkgname}} from '{{pkgname}}'

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

  // variable that enables pooling of selected gongstruct
  obsTimer: Observable<number> = timer(1000, 1000)
  lastSelectionDate: string = ''

  constructor(
  ) {

  }

  ngOnInit(): void {

    // pool the gongdoc command and check wether a gongstruct has been selected
    this.obsTimer.subscribe(
      currTime => {
      }
    )
  }
}
`
