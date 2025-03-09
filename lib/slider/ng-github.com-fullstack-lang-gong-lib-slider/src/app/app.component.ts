import { Component, OnInit } from '@angular/core';

import { AngularSplitModule } from 'angular-split';

import * as slider from '../../projects/slider/src/public-api'

import { SliderSpecificComponent } from '../../projects/sliderspecific/src/lib/slider-specific/slider-specific.component'

import { SplitSpecificComponent } from '@vendored_components/github.com/fullstack-lang/gong/lib/split/ng-github.com-fullstack-lang-gong-lib-split/projects/splitspecific/src/lib/split-specific/split-specific.component'


@Component({
  selector: 'app-root',
  standalone: true,
  imports: [
    AngularSplitModule,
    SplitSpecificComponent,
    SliderSpecificComponent
  ],

  templateUrl: './app.component.html',
})
export class AppComponent implements OnInit {

  StackName = "slider"

  constructor(
  ) {
  }

  ngOnInit(): void {
  }
}
