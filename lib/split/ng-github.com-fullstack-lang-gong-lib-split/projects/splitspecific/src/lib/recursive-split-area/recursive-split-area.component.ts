import { Component, Input } from '@angular/core';

import * as split from '../../../../split/src/public-api'

import { TreeSpecificComponent } from '../../../../../../../tree/ng-github.com-fullstack-lang-gong-lib-tree/projects/treespecific/src/lib/tree-specific/tree-specific.component'
import { TableSpecificComponent } from '../../../../../../../table/ng-github.com-fullstack-lang-gong-lib-table/projects/tablespecific/src/lib/table-specific/table-specific.component'
import { FormSpecificComponent } from '../../../../../../../table/ng-github.com-fullstack-lang-gong-lib-table/projects/tablespecific/src/lib/form-specific/form-specific.component'
import { DocSpecificComponent } from '../../../../../../../doc/ng-github.com-fullstack-lang-gong-lib-doc/projects/docspecific/src/lib/doc-specific/doc-specific.component'
import { SvgSpecificComponent } from '../../../../../../../svg/ng-github.com-fullstack-lang-gong-lib-svg/projects/svgspecific/src/lib/svg-specific/svg-specific.component'
import { CommonModule } from '@angular/common';
import { AngularSplitModule } from 'angular-split';


@Component({
  selector: 'lib-recursive-split-area',
  imports: [

    CommonModule,
    AngularSplitModule,
    
    TreeSpecificComponent,
    TableSpecificComponent,
    FormSpecificComponent,
    // SvgSpecificComponent,
    // DocSpecificComponent,

    RecursiveSplitAreaComponent, // self import for recursion
  ],
  templateUrl: './recursive-split-area.component.html',
  styleUrl: './recursive-split-area.component.css'
})
export class RecursiveSplitAreaComponent {
  @Input() splitArea!: split.AsSplitArea;

}
