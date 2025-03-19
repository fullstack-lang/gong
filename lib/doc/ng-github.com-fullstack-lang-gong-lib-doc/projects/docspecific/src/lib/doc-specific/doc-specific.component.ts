import { Component, Input, OnInit } from '@angular/core';

import { Router, RouterState } from '@angular/router';

import { AngularSplitModule } from 'angular-split';

import { TreeSpecificComponent } from '../../../../../../../tree/ng-github.com-fullstack-lang-gong-lib-tree/projects/treespecific/src/lib/tree-specific/tree-specific.component'
import { SvgSpecificComponent } from '../../../../../../../svg/ng-github.com-fullstack-lang-gong-lib-svg/projects/svgspecific/src/lib/svg-specific/svg-specific.component'

import * as doc from '../../../../doc/src/public-api'

@Component({
  selector: 'lib-doc-specific',
  imports: [
    AngularSplitModule,
    TreeSpecificComponent,
    SvgSpecificComponent,
  ],
  templateUrl: './doc-specific.component.html',
  styleUrl: './doc-specific.component.css'
})
export class DocSpecificComponent implements OnInit {
  @Input() Name: string = ""

  TreeNames = doc.TreeNames

  constructor(
  ) { }

  ngOnInit(): void {
    console.log("DocSpecificComponent:ngOnInit")
  }
}
