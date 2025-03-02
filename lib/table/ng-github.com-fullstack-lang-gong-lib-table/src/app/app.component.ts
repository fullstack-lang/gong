import { Component, OnInit } from '@angular/core';

import { Observable, combineLatest, timer } from 'rxjs'

// for angular & angular material
import { FormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { MatRadioModule } from '@angular/material/radio';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';

import { AngularSplitModule } from 'angular-split';

import * as table from '../../projects/table/src/public-api'

import { TableSpecificComponent } from '../../projects/tablespecific/src/lib/table-specific/table-specific.component'

import { TreeComponent } from '@vendored_components/github.com/fullstack-lang/gongtree/ng-github.com-fullstack-lang-gongtree/projects/gongtreespecific/src/public-api'
import { MaterialTableComponent } from '@vendored_components/github.com/fullstack-lang/gongtable/ng-github.com-fullstack-lang-gongtable/projects/gongtablespecific/src/lib/material-table/material-table.component';
import { MaterialFormComponent } from '@vendored_components/github.com/fullstack-lang/gongtable/ng-github.com-fullstack-lang-gongtable/projects/gongtablespecific/src/lib/material-form/material-form.component';
import * as gongtable from '@vendored_components/github.com/fullstack-lang/gongtable/ng-github.com-fullstack-lang-gongtable/projects/gongtable/src/public-api';
import { PanelComponent } from '@vendored_components/github.com/fullstack-lang/gongdoc/ng-github.com-fullstack-lang-gongdoc/projects/gongdocspecific/src/public-api'
import { GongsvgDiagrammingComponent } from '@vendored_components/github.com/fullstack-lang/gongsvg/ng-github.com-fullstack-lang-gongsvg/projects/gongsvgspecific/src/lib/gongsvg-diagramming/gongsvg-diagramming'


@Component({
  selector: 'app-root',
  standalone: true,
  imports: [

    CommonModule,
    FormsModule,

    MatRadioModule,
    MatButtonModule,
    MatIconModule,

    AngularSplitModule,

    TreeComponent,
    MaterialTableComponent,
    MaterialFormComponent,
    PanelComponent,

    TableSpecificComponent

  ],

  templateUrl: './app.component.html',
})
export class AppComponent implements OnInit {

  table = 'Table'
  probe = 'Table Data/Model'
  view = this.table

  views: string[] = [this.table, this.probe];

  scrollStyle = {
    'overflow- x': 'auto',
    'width': '100%',  // Ensure the div takes the full width of its parent container
  }

  StackName = "table"
  StackType = table.StackType

  TableExtraPathEnum = gongtable.TableExtraPathEnum

  constructor(
  ) {

  }

  ngOnInit(): void {
  }
}
