import { Component, OnInit } from '@angular/core';

import { AngularSplitModule } from 'angular-split';

import * as split from '../../projects/split/src/public-api'

import { SplitSpecificComponent } from '../../projects/splitspecific/src/lib/split-specific/split-specific.component'


@Component({
  selector: 'app-root',
  standalone: true,
  imports: [
    AngularSplitModule,
    SplitSpecificComponent,
  ],

  templateUrl: './app.component.html',
})
export class AppComponent implements OnInit {

  StackName = "split"
  StackProbeName = "split" + "-probe"

  constructor(
  ) {
  }

  ngOnInit(): void {
  }
}
