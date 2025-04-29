import { Component, OnInit, ViewChild } from '@angular/core';

import { SplitSpecificComponent } from '@vendored_components/github.com/fullstack-lang/gong/lib/split/ng-github.com-fullstack-lang-gong-lib-split/projects/splitspecific/src/lib/split-specific/split-specific.component'


import * as doc from '../../projects/doc/src/public-api'

@Component({
    selector: 'app-root',
    templateUrl: './app.component.html',
    imports: [
      SplitSpecificComponent
    ]
})
export class AppComponent implements OnInit{
  StackType = doc.StackType

  ngOnInit(): void {
    console.log("AppComponent:ngOnInit")
  }
}
