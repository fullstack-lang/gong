package angular

const NgFileAppComponentTs = `import { Component, OnInit } from '@angular/core';

import { Observable, combineLatest, timer } from 'rxjs'

// for angular & angular material
import { FormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { MatRadioModule } from '@angular/material/radio';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';

import { AngularSplitModule } from 'angular-split';

import * as {{pkgname}} from '../../projects/{{pkgname}}/src/public-api'

import { {{TitlePkgName}}SpecificComponent } from '../../projects/{{pkgname}}specific/src/lib/{{pkgname}}-specific/{{pkgname}}-specific.component'

import { TreeSpecificComponent } from '@vendored_components/github.com/fullstack-lang/gong/lib/tree/ng-github.com-fullstack-lang-gong-lib-tree/projects/treespecific/src/lib/tree-specific/tree-specific.component'
import { TableSpecificComponent } from '@vendored_components/github.com/fullstack-lang/gong/lib/table/ng-github.com-fullstack-lang-gong-lib-table/projects/tablespecific/src/lib/table-specific/table-specific.component'
import { FormSpecificComponent } from '@vendored_components/github.com/fullstack-lang/gong/lib/table/ng-github.com-fullstack-lang-gong-lib-table/projects/tablespecific/src/lib/form-specific/form-specific.component'
import { DocSpecificComponent } from '@vendored_components/github.com/fullstack-lang/gong/lib/doc/ng-github.com-fullstack-lang-gong-lib-doc/projects/docspecific/src/lib/doc-specific/doc-specific.component'

import * as svg from '@vendored_components/github.com/fullstack-lang/gong/lib/svg/ng-github.com-fullstack-lang-gong-lib-svg/projects/svg/src/public-api'
import * as tree from '@vendored_components/github.com/fullstack-lang/gong/lib/tree/ng-github.com-fullstack-lang-gong-lib-tree/projects/tree/src/public-api'
import * as table from '@vendored_components/github.com/fullstack-lang/gong/lib/table/ng-github.com-fullstack-lang-gong-lib-table/projects/table/src/public-api'


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

    TreeSpecificComponent,
    TableSpecificComponent,
    FormSpecificComponent,
    DocSpecificComponent,

    {{TitlePkgName}}SpecificComponent

  ],

  templateUrl: './app.component.html',
})
export class AppComponent implements OnInit {

  {{pkgname}} = '{{TitlePkgName}}'
  probe = '{{TitlePkgName}} Data/Model'
  view = this.{{pkgname}}

  views: string[] = [this.{{pkgname}}, this.probe];

  scrollStyle = {
    'overflow- x': 'auto',
    'width': '100%',  // Ensure the div takes the full width of its parent container
  }

  StackName = "{{pkgname}}"
  StackType = {{pkgname}}.StackType

  TableExtraPathEnum = table.TableExtraPathEnum

  constructor(
  ) {

  }

  ngOnInit(): void {
  }
}
`
