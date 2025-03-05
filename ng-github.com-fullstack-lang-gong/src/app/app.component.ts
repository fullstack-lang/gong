import { Component, OnInit } from '@angular/core';

import { Observable, combineLatest, timer } from 'rxjs'

// for angular & angular material
import { FormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { MatRadioModule } from '@angular/material/radio';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';

import { AngularSplitModule } from 'angular-split';

import * as gong from '../../projects/gong/src/public-api'

import { TreeSpecificComponent } from '../../../lib/tree/ng-github.com-fullstack-lang-gong-lib-tree/projects/treespecific/src/lib/tree-specific/tree-specific.component'
import { TableSpecificComponent } from '../../../lib/table/ng-github.com-fullstack-lang-gong-lib-table/projects/tablespecific/src/lib/table-specific/table-specific.component'
import { FormSpecificComponent } from '../../../lib/table/ng-github.com-fullstack-lang-gong-lib-table/projects/tablespecific/src/lib/form-specific/form-specific.component'
// import { SvgSpecificComponent } from '../../../lib/svg/ng-github.com-fullstack-lang-gong-lib-svg/projects/svgspecific/src/lib/svg-specific/svg-specific.component'
import { DocSpecificComponent } from '../../../lib/doc/ng-github.com-fullstack-lang-gong-lib-doc/projects/docspecific/src/lib/doc-specific/doc-specific.component'

import * as table from '../../../lib/table/ng-github.com-fullstack-lang-gong-lib-table/projects/table/src/public-api'

@Component({
  selector: 'app-root',
  imports: [
    CommonModule,
    FormsModule,
    MatRadioModule,
    MatButtonModule,
    MatIconModule,
    AngularSplitModule,

    TreeSpecificComponent,
    TableSpecificComponent,
    FormSpecificComponent,
    // SvgSpecificComponent,
    DocSpecificComponent,

  ],
  templateUrl: './app.component.html'
})
export class AppComponent implements OnInit {

  gong = 'Gong'
  probe = 'Gong Data/Model'
  view = this.gong

  views: string[] = [this.gong, this.probe];

  scrollStyle = {
    'overflow- x': 'auto',
    'width': '100%',  // Ensure the div takes the full width of its parent container
  }

  StackName = "gong"
  StackType = gong.StackType

  TableExtraPathEnum = table.TableExtraPathEnum

  constructor(
  ) {

  }

  ngOnInit(): void {
  }
}
