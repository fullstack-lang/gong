import { Component, OnInit } from '@angular/core';

import { Observable, combineLatest, timer } from 'rxjs'

import * as gongdoc from 'gongdoc'
import * as test from 'test'

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
})
export class AppComponent implements OnInit {

  view = 'Default view'
  default = 'Default view'
  diagrams = 'Diagrams view'
  meta = 'Meta view'

  views: string[] = [this.default, this.diagrams, this.meta];

  // variable that enables pooling of selected gongstruct
  obsTimer: Observable<number> = timer(1000, 1000)
  lastSelectionDate: string = ''

  constructor(private gongdocClassshapeService: gongdoc.ClassshapeService,
    private gongstructSelectionService: test.GongstructSelectionService
  ) {

  }

  ngOnInit(): void {

    // pool the gongdoc command and check wether a gongstruct has been selected
    this.obsTimer.subscribe(
      currTime => {

        // pool all classshapes and find which one is selected
        this.gongdocClassshapeService.getClassshapes().subscribe(
          classshapes => {
            for (let classshape of classshapes) {
              if (classshape.IsSelected) {
                classshape.IsSelected = false
                // console.log("classshape " + classshape.ReferenceName + " is selected")
                this.gongdocClassshapeService.updateClassshape(classshape).subscribe(
                  classshape2 => {
                    // console.log("classshape has been unselected")
                  }
                )
                this.gongstructSelectionService.gongstructSelected( classshape.ReferenceName)
              }
            }
          }
        )
      }
    )
  }
}
