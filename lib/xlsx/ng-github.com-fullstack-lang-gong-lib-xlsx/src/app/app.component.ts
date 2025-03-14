import { Component, OnInit } from '@angular/core';

import { AngularSplitModule } from 'angular-split';

import * as xlsx from '../../projects/xlsx/src/public-api'

import { XlsxSpecificComponent } from '../../projects/xlsxspecific/src/lib/xlsx-specific/xlsx-specific.component'

import { SplitSpecificComponent } from '@vendored_components/github.com/fullstack-lang/gong/lib/split/ng-github.com-fullstack-lang-gong-lib-split/projects/splitspecific/src/lib/split-specific/split-specific.component'


@Component({
  selector: 'app-root',
  standalone: true,
  imports: [
    AngularSplitModule,
    SplitSpecificComponent,
    XlsxSpecificComponent
  ],

  templateUrl: './app.component.html',
})
export class AppComponent implements OnInit {

  StackName = "xlsx"

  constructor(
  ) {
  }

  ngOnInit(): void {
  }
}
