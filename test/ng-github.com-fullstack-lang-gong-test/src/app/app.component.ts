import { Component, OnInit } from '@angular/core';

import { Observable, combineLatest, timer } from 'rxjs'

import { CommonModule } from '@angular/common';

// for angular material
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { MatRadioModule } from '@angular/material/radio';

import { FormsModule } from '@angular/forms';

import { AngularSplitModule } from 'angular-split';

import * as test from '../../projects/test/src/public-api'


import { TreeComponent } from '@vendored_components/github.com/fullstack-lang/gongtree/ng-github.com-fullstack-lang-gongtree/projects/gongtreespecific/src/public-api'
import { MaterialTableComponent } from '@vendored_components/github.com/fullstack-lang/gongtable/ng-github.com-fullstack-lang-gongtable/projects/gongtablespecific/src/lib/material-table/material-table.component';
import { MaterialFormComponent } from '@vendored_components/github.com/fullstack-lang/gongtable/ng-github.com-fullstack-lang-gongtable/projects/gongtablespecific/src/lib/material-form/material-form.component';
import * as gongtable from '@vendored_components/github.com/fullstack-lang/gongtable/ng-github.com-fullstack-lang-gongtable/projects/gongtable/src/public-api';
import { PanelComponent } from '@vendored_components/github.com/fullstack-lang/gongdoc/ng-github.com-fullstack-lang-gongdoc/projects/gongdocspecific/src/public-api'

import { TestspecificComponent } from '../../projects/testspecific/src/public-api'

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [
    CommonModule,

    MatRadioModule,
    FormsModule,
    CommonModule,
    MatButtonModule,
    MatIconModule,

    AngularSplitModule,

    TreeComponent,
    MaterialTableComponent,
    MaterialFormComponent,
    PanelComponent,

    TestspecificComponent,
  ],

  templateUrl: './app.component.html',
})
export class AppComponent implements OnInit {

  test = 'Test'
  probe = 'Test Data/Model'
  view = this.test

  views: string[] = [this.test, this.probe];

  scrollStyle = {
    'overflow- x': 'auto',
    'width': '100%',  // Ensure the div takes the full width of its parent container
  }

  StackName = "test"
  StackType = test.StackType

  TableExtraPathEnum = gongtable.TableExtraPathEnum

  constructor(
  ) {

  }

  ngOnInit(): void {
  }
}
