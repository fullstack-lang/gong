import { Component, OnInit } from '@angular/core';

import { AngularSplitModule } from 'angular-split';

import * as svg from '../../projects/svg/src/public-api'

import { SvgSpecificComponent } from '../../projects/svgspecific/src/lib/svg-specific/svg-specific.component'

import { SplitSpecificComponent } from '@vendored_components/github.com/fullstack-lang/gong/lib/split/ng-github.com-fullstack-lang-gong-lib-split/projects/splitspecific/src/lib/split-specific/split-specific.component'


@Component({
  selector: 'app-root',
  standalone: true,
  imports: [
    AngularSplitModule,
    SplitSpecificComponent,
    SvgSpecificComponent
  ],

  templateUrl: './app.component.html',
})
export class AppComponent implements OnInit {

  StackName = svg.StackName.StackNameDefault

  constructor(
  ) {
  }

  ngOnInit(): void {
  }
}
