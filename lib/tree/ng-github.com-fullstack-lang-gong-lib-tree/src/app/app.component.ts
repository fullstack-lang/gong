import { Component, OnInit } from '@angular/core';

import { AngularSplitModule } from 'angular-split';

import * as tree from '../../projects/tree/src/public-api'

import { TreeSpecificComponent } from '../../projects/treespecific/src/lib/tree-specific/tree-specific.component'

import { SplitSpecificComponent } from '@vendored_components/github.com/fullstack-lang/gong/lib/split/ng-github.com-fullstack-lang-gong-lib-split/projects/splitspecific/src/lib/split-specific/split-specific.component'


@Component({
  selector: 'app-root',
  standalone: true,
  imports: [
    AngularSplitModule,
    SplitSpecificComponent,
    TreeSpecificComponent
  ],

  templateUrl: './app.component.html',
})
export class AppComponent implements OnInit {

  StackName = tree.TreeStackName.TreeStackDefaultName

  constructor(
  ) {
  }

  ngOnInit(): void {
  }
}
