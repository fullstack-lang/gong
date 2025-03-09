import { Component, OnInit } from '@angular/core';

import { AngularSplitModule } from 'angular-split';

import * as test3 from '../../projects/test3/src/public-api'

import { Test3SpecificComponent } from '../../projects/test3specific/src/lib/test3-specific/test3-specific.component'

import { SplitSpecificComponent } from '@vendored_components/github.com/fullstack-lang/gong/lib/split/ng-github.com-fullstack-lang-gong-lib-split/projects/splitspecific/src/lib/split-specific/split-specific.component'


@Component({
  selector: 'app-root',
  standalone: true,
  imports: [
    AngularSplitModule,
    SplitSpecificComponent,
    Test3SpecificComponent
  ],

  templateUrl: './app.component.html',
})
export class AppComponent implements OnInit {

  StackName = "test3"

  constructor(
  ) {
  }

  ngOnInit(): void {
  }
}
