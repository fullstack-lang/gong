import { Component, OnInit, ViewChild } from '@angular/core';

import { DocSpecificComponent } from '../../projects/docspecific/src/lib/doc-specific/doc-specific.component'

import * as gongtable from '@vendored_components/github.com/fullstack-lang/gongtable/ng-github.com-fullstack-lang-gongtable/projects/gongtable/src/public-api';

import * as doc from '../../projects/doc/src/public-api'

import * as svg from '../../../../svg/ng-github.com-fullstack-lang-gong-lib-svg/projects/svg/src/public-api'

@Component({
    selector: 'app-root',
    templateUrl: './app.component.html',
    imports: [
      DocSpecificComponent,
    ]
})
export class AppComponent implements OnInit {

  default = 'DOC Data/Model'
  diagrammer = 'Diagram Edition'
  probe = "probe"
  view = this.probe

  views: string[] = [this.diagrammer, this.probe];

  GONG__MODEL__StacksPath = "github.com/fullstack-lang/gong/lib/doc/go/models"
  GONG__DATA__StackPath = "gongdoc"

  loading = true

  scrollStyle = {
    'overflow- x': 'auto',
    'width': '100%',  // Ensure the div takes the full width of its parent container
  }

  StackType = doc.StackType

  StackName = "doc"
  TableExtraPathEnum = gongtable.TableExtraPathEnum
  StacksNames = doc.StacksNames
  SVGStackType = svg.StackType


  constructor(
  ) {

  }

  ngOnInit(): void {
    this.loading = false
  }
}
