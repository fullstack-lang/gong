import { Component, Input, OnInit } from '@angular/core';

import { Router, RouterState } from '@angular/router';

import { AngularSplitModule } from 'angular-split';

import { TreeComponent } from '@vendored_components/github.com/fullstack-lang/gongtree/ng-github.com-fullstack-lang-gongtree/projects/gongtreespecific/src/public-api'

import { GongsvgDiagrammingComponent } from '@vendored_components/github.com/fullstack-lang/gongsvg/ng-github.com-fullstack-lang-gongsvg/projects/gongsvgspecific/src/lib/gongsvg-diagramming/gongsvg-diagramming'

@Component({
  selector: 'lib-pkgelt-docs',
  templateUrl: './pkgelt-docs.component.html',
  styleUrls: ['./pkgelt-docs.component.css'],
  standalone: true,
  imports: [
    AngularSplitModule,
    TreeComponent,
    GongsvgDiagrammingComponent
  ]
})
export class PkgeltDocsComponent implements OnInit {

  @Input() GONG__StackPath: string = ""

  constructor(
    private router: Router,
  ) { }

  ngOnInit(): void {
    // console.log("PkgeltDocsComponent->GONG__StackPath : ", this.GONG__StackPath)
  }
}
