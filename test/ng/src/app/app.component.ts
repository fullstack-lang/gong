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
  gongdocCommandSingloton: gongdoc.GongdocCommandDB = new gongdoc.GongdocCommandDB
  lastSelectionDate: string = ''

  constructor(private gongdocCommandService: gongdoc.GongdocCommandService,
    private gongstructSelectionService: test.GongstructSelectionService
  ) {

  }

  ngOnInit(): void {

    // pool the gongdoc command and check wether a gongstruct has been selected
    this.obsTimer.subscribe(
      currTime => {

        // fetch the command singloton
        this.gongdocCommandService.getGongdocCommands().subscribe(
          gongdocCommands => {
            for (let gongdocCommand_ of gongdocCommands) {
              this.gongdocCommandSingloton = gongdocCommand_
            }

            // check the type of command
            if (this.gongdocCommandSingloton.Command == gongdoc.GongdocCommandType.DIAGRAM_GONGSTRUCT_SELECT) {
              if (this.lastSelectionDate != this.gongdocCommandSingloton.Date) {
                // console.log("New user selection of gongstruct " + this.gongdocCommandSingloton.StructName)

                this.gongstructSelectionService.gongstructSelected(this.gongdocCommandSingloton.StructName)
                this.lastSelectionDate = this.gongdocCommandSingloton.Date
              }
            }
          }
        )
      }
    )
  }
}
