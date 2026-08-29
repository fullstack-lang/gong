import { Component, OnInit } from '@angular/core';

import { GongSplitComponent, GongSplitAreaComponent } from '@vendored_components/github.com/fullstack-lang/gong/lib/split/ng-github.com-fullstack-lang-gong-lib-split/projects/splitspecific/src/public-api';

import * as test2 from '../../projects/test2/src/public-api'

import { Test2SpecificComponent } from '../../projects/test2specific/src/lib/test2-specific/test2-specific.component'

import { SplitSpecificComponent } from '@vendored_components/github.com/fullstack-lang/gong/lib/split/ng-github.com-fullstack-lang-gong-lib-split/projects/splitspecific/src/lib/split-specific/split-specific.component'


@Component({
  selector: 'app-root',
  standalone: true,
  imports: [
    GongSplitComponent,
    GongSplitAreaComponent,
    SplitSpecificComponent,
    Test2SpecificComponent
  ],

  templateUrl: './app.component.html',
})
export class AppComponent implements OnInit {

  StackName = "test2"

  constructor(
  ) {
  }

  ngOnInit(): void {
  }
}
