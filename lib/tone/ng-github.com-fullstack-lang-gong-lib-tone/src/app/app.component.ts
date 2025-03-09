import { Component, OnInit } from '@angular/core';

import { AngularSplitModule } from 'angular-split';

import * as tone from '../../projects/tone/src/public-api'

import { ToneSpecificComponent } from '../../projects/tonespecific/src/lib/tone-specific/tone-specific.component'

import { SplitSpecificComponent } from '@vendored_components/github.com/fullstack-lang/gong/lib/split/ng-github.com-fullstack-lang-gong-lib-split/projects/splitspecific/src/lib/split-specific/split-specific.component'


@Component({
  selector: 'app-root',
  standalone: true,
  imports: [
    AngularSplitModule,
    SplitSpecificComponent,
    ToneSpecificComponent
  ],

  templateUrl: './app.component.html',
})
export class AppComponent implements OnInit {

  StackName = tone.StacksNames.Tone

  constructor(
  ) {
  }

  ngOnInit(): void {
  }
}
