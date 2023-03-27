import { Component, OnInit } from '@angular/core';

import { Observable, combineLatest, timer } from 'rxjs'

import * as gongdoc from 'gongdoc'
import * as test from 'test'

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
})
export class AppComponent implements OnInit {

  view = 'Data view'
  default = 'Data view'
  model = 'Model view'

  views: string[] = [this.default, this.model];

  GONG__StackPath = "github.com/fullstack-lang/gong/test/go/models"

  // variable that enables pooling of selected gongstruct
  obsTimer: Observable<number> = timer(1000, 1000)
  lastSelectionDate: string = ''

  constructor(private gongdocGongStructShapeService: gongdoc.GongStructShapeService,
    private gongstructSelectionService: test.GongstructSelectionService
  ) {

  }

  ngOnInit(): void {

    // pool the gongdoc command and check wether a gongstruct has been selected
    this.obsTimer.subscribe(
      currTime => {

        // pool all GongStructShapes and find which one is selected
        this.gongdocGongStructShapeService.getGongStructShapes().subscribe(
          GongStructShapes => {
            for (let GongStructShape of GongStructShapes) {
              if (GongStructShape.IsSelected) {
                GongStructShape.IsSelected = false
                // console.log("GongStructShape " + GongStructShape.ReferenceName + " is selected")
                this.gongdocGongStructShapeService.updateGongStructShape(GongStructShape, "").subscribe(
                  GongStructShape2 => {
                    // console.log("GongStructShape has been unselected")
                  }
                )
                this.gongstructSelectionService.gongstructSelected(GongStructShape.Identifier)
              }
            }
          }
        )
      }
    )
  }
}
