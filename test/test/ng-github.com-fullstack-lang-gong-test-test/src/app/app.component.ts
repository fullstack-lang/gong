import { Component, OnInit } from '@angular/core';

import { AngularSplitModule } from 'angular-split';

import * as test from '../../projects/test/src/public-api'

import { TestSpecificComponent } from '../../projects/testspecific/src/lib/test-specific/test-specific.component'

import { SplitSpecificComponent } from '@vendored_components/github.com/fullstack-lang/gong/lib/split/ng-github.com-fullstack-lang-gong-lib-split/projects/splitspecific/src/lib/split-specific/split-specific.component'


@Component({
  selector: 'app-root',
  standalone: true,
  imports: [
    AngularSplitModule,
    SplitSpecificComponent,
    TestSpecificComponent
  ],

  templateUrl: './app.component.html',
})
export class AppComponent implements OnInit {

  StackName = "test"

  constructor(
  ) {
  }

  ngOnInit(): void {
  }
}
