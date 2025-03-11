import { Component, OnInit } from '@angular/core';

import { AngularSplitModule } from 'angular-split';

import * as gantt from '../../projects/gantt/src/public-api'

import { GanttSpecificComponent } from '../../projects/ganttspecific/src/lib/gantt-specific/gantt-specific.component'

import { SplitSpecificComponent } from '@vendored_components/github.com/fullstack-lang/gong/lib/split/ng-github.com-fullstack-lang-gong-lib-split/projects/splitspecific/src/lib/split-specific/split-specific.component'


@Component({
  selector: 'app-root',
  standalone: true,
  imports: [
    AngularSplitModule,
    SplitSpecificComponent,
    GanttSpecificComponent
  ],

  templateUrl: './app.component.html',
})
export class AppComponent implements OnInit {

  StackName = "gantt"

  constructor(
  ) {
  }

  ngOnInit(): void {
  }
}
