import { Component, Input, OnInit } from '@angular/core';

import * as split from '../../../../split/src/public-api'

import { MatRadioModule } from '@angular/material/radio';
import { FormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';

import { AngularSplitModule } from 'angular-split';

import { ButtonSpecificComponent } from '../../../../../../../button/ng-github.com-fullstack-lang-gong-lib-button/projects/buttonspecific/src/lib/button-specific/button-specific.component'
import { CursorSpecificComponent } from '../../../../../../../cursor/ng-github.com-fullstack-lang-gong-lib-cursor/projects/cursorspecific/src/lib/cursor-specific/cursor-specific.component'
import { DocSpecificComponent } from '../../../../../../../doc/ng-github.com-fullstack-lang-gong-lib-doc/projects/docspecific/src/lib/doc-specific/doc-specific.component'
import { FormSpecificComponent } from '../../../../../../../table/ng-github.com-fullstack-lang-gong-lib-table/projects/tablespecific/src/lib/form-specific/form-specific.component'
import { LoadSpecificComponent } from '../../../../../../../load/ng-github.com-fullstack-lang-gong-lib-load/projects/loadspecific/src/lib/load-specific/load-specific.component'
import { SliderSpecificComponent } from '../../../../../../../slider/ng-github.com-fullstack-lang-gong-lib-slider/projects/sliderspecific/src/lib/slider-specific/slider-specific.component'
import { SvgSpecificComponent } from '../../../../../../../svg/ng-github.com-fullstack-lang-gong-lib-svg/projects/svgspecific/src/lib/svg-specific/svg-specific.component'
import { TableSpecificComponent } from '../../../../../../../table/ng-github.com-fullstack-lang-gong-lib-table/projects/tablespecific/src/lib/table-specific/table-specific.component'
import { ToneSpecificComponent } from '../../../../../../../tone/ng-github.com-fullstack-lang-gong-lib-tone/projects/tonespecific/src/lib/tone-specific/tone-specific.component'
import { TreeSpecificComponent } from '../../../../../../../tree/ng-github.com-fullstack-lang-gong-lib-tree/projects/treespecific/src/lib/tree-specific/tree-specific.component'
import { XlsxSpecificComponent } from '../../../../../../../xlsx/ng-github.com-fullstack-lang-gong-lib-xlsx/projects/xlsxspecific/src/lib/xlsx-specific/xlsx-specific.component'



@Component({
  selector: 'lib-split-specific',
  imports: [
    MatRadioModule,
    CommonModule,
    FormsModule,

    AngularSplitModule,

    ButtonSpecificComponent,
    CursorSpecificComponent,
    DocSpecificComponent,
    FormSpecificComponent,
    LoadSpecificComponent,
    SliderSpecificComponent,
    SvgSpecificComponent,
    TableSpecificComponent,
    ToneSpecificComponent,
    TreeSpecificComponent,
    XlsxSpecificComponent,

  
  ],
  templateUrl: './split-specific.component.html',
  styleUrl: './split-specific.component.css'
})
export class SplitSpecificComponent implements OnInit {
    @Input() Name: string = ""

    public frontRepo?: split.FrontRepo;

    view = ""

    constructor(
      private frontRepoService: split.FrontRepoService,
    ) { }

    ngOnInit(): void {
      console.log("ngOnInit");
  
      this.frontRepoService.connectToWebSocket(this.Name).subscribe({
        next: (frontRepo) => {
          this.frontRepo = frontRepo;
  
          if (this.frontRepo.array_Views.length > 0) {

            this.frontRepo.array_Views.sort((a: split.View, b : split.View) => {
              return a.ID - b.ID
            })

            this.view = this.frontRepo.array_Views[0].Name
          }

        }
      }
      )
    }
}