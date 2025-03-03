import { Component, Input, OnInit } from '@angular/core';

import { Router, RouterState } from '@angular/router';

import { AngularSplitModule } from 'angular-split';

import { TreeSpecificComponent } from '../../../../../../../tree/ng-github.com-fullstack-lang-gong-lib-tree/projects/treespecific/src/lib/tree-specific/tree-specific.component'
import { SvgSpecificComponent } from '../../../../../../../svg/ng-github.com-fullstack-lang-gong-lib-svg/projects/svgspecific/src/lib/svg-specific/svg-specific.component'


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
  @Input() GONG__StackPath: string = ""

  constructor(
    private router: Router,
  ) { }

  ngOnInit(): void {
    console.log("PkgeltDocsComponent->GONG__StackPath : ", this.GONG__StackPath)
  }
}
