import { Component, OnInit } from '@angular/core';

import { AngularSplitModule } from 'angular-split';

import * as sim from '../../projects/sim/src/public-api'

import { SimSpecificComponent } from '../../projects/simspecific/src/lib/sim-specific/sim-specific.component'

import { SplitSpecificComponent } from '@vendored_components/github.com/fullstack-lang/gong/lib/split/ng-github.com-fullstack-lang-gong-lib-split/projects/splitspecific/src/lib/split-specific/split-specific.component'


@Component({
  selector: 'app-root',
  standalone: true,
  imports: [
    AngularSplitModule,
    SplitSpecificComponent,
    SimSpecificComponent
  ],

  templateUrl: './app.component.html',
})
export class AppComponent implements OnInit {

  StackName = "sim"

  constructor(
  ) {
  }

  ngOnInit(): void {
  }
}
