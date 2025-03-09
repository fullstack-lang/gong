import { Component, OnInit } from '@angular/core';

import { AngularSplitModule } from 'angular-split';

import * as gong from '../../projects/gong/src/public-api'

import { GongSpecificComponent } from '../../projects/gongspecific/src/lib/gong-specific/gong-specific.component'

import { SplitSpecificComponent } from '@vendored_components/github.com/fullstack-lang/gong/lib/split/ng-github.com-fullstack-lang-gong-lib-split/projects/splitspecific/src/lib/split-specific/split-specific.component'


@Component({
  selector: 'app-root',
  standalone: true,
  imports: [
    AngularSplitModule,
    SplitSpecificComponent,
    GongSpecificComponent
  ],

  templateUrl: './app.component.html',
})
export class AppComponent implements OnInit {

  StackName = "gong"

  constructor(
  ) {
  }

  ngOnInit(): void {
  }
}
