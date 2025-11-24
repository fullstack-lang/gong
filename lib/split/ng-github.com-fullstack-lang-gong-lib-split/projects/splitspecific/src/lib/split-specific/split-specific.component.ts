import { Component, Renderer2, Input, OnInit, Inject } from '@angular/core';
import { DOCUMENT } from '@angular/common';

import * as split from '../../../../split/src/public-api'

import { MatRadioModule } from '@angular/material/radio';
import { FormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';

import { AngularSplitModule } from 'angular-split';

import { ButtonSpecificComponent } from '../../../../../../../button/ng-github.com-fullstack-lang-gong-lib-button/projects/buttonspecific/src/lib/button-specific/button-specific.component'
import { CursorSpecificComponent } from '../../../../../../../cursor/ng-github.com-fullstack-lang-gong-lib-cursor/projects/cursorspecific/src/lib/cursor-specific/cursor-specific.component'
import { FormSpecificComponent } from '../../../../../../../table/ng-github.com-fullstack-lang-gong-lib-table/projects/tablespecific/src/lib/form-specific/form-specific.component'
import { LoadSpecificComponent } from '../../../../../../../load/ng-github.com-fullstack-lang-gong-lib-load/projects/loadspecific/src/lib/load-specific/load-specific.component'
import { MarkdownSpecificComponent } from '../../../../../../../markdown/ng-github.com-fullstack-lang-gong-lib-markdown/projects/markdownspecific/src/lib/markdown-specific/markdown-specific.component'
import { SliderSpecificComponent } from '../../../../../../../slider/ng-github.com-fullstack-lang-gong-lib-slider/projects/sliderspecific/src/lib/slider-specific/slider-specific.component'
import { SvgSpecificComponent } from '../../../../../../../svg/ng-github.com-fullstack-lang-gong-lib-svg/projects/svgspecific/src/lib/svg-specific/svg-specific.component'
import { TableSpecificComponent } from '../../../../../../../table/ng-github.com-fullstack-lang-gong-lib-table/projects/tablespecific/src/lib/table-specific/table-specific.component'
import { ToneSpecificComponent } from '../../../../../../../tone/ng-github.com-fullstack-lang-gong-lib-tone/projects/tonespecific/src/lib/tone-specific/tone-specific.component'
import { TreeSpecificComponent } from '../../../../../../../tree/ng-github.com-fullstack-lang-gong-lib-tree/projects/treespecific/src/lib/tree-specific/tree-specific.component'

// to set the title of the application
import { DomSanitizer, SafeHtml, Title } from '@angular/platform-browser';

@Component({
  selector: 'lib-split-specific',
  imports: [
    MatRadioModule,
    CommonModule,
    FormsModule,

    AngularSplitModule,

    ButtonSpecificComponent,
    CursorSpecificComponent,
    FormSpecificComponent,
    LoadSpecificComponent,
    MarkdownSpecificComponent,
    SliderSpecificComponent,
    SvgSpecificComponent,
    TableSpecificComponent,
    ToneSpecificComponent,
    TreeSpecificComponent,

  ],
  templateUrl: './split-specific.component.html',
  styleUrl: './split-specific.component.css'
})
export class SplitSpecificComponent implements OnInit {
  @Input() Name: string = ""

  public frontRepo?: split.FrontRepo;

  public view: split.View | undefined

  radioButtonHeight = 40

  constructor(
    private frontRepoService: split.FrontRepoService,
    private titleService: Title,
    private renderer: Renderer2,
    private sanitizer: DomSanitizer,
    @Inject(DOCUMENT) private document: Document
  ) { }

  ngOnInit(): void {
    console.log("ngOnInit");

    this.frontRepoService.connectToWebSocket(this.Name).subscribe({
      next: (frontRepo) => {
        this.frontRepo = frontRepo;

        if (this.frontRepo.array_Titles.length > 0) {
          this.titleService.setTitle(this.frontRepo.array_Titles[0].Name)
        }


        if (this.frontRepo.array_FavIcons.length > 0) {
          this.setSvgFavicon(this.frontRepo.array_FavIcons[0].SVG)
        }

        if (this.frontRepo.array_LogoOnTheLefts.length > 0) {

          this.frontRepo.array_LogoOnTheLefts.sort((a: split.LogoOnTheLeft, b: split.LogoOnTheLeft) => {
            return a.ID - b.ID
          })
        }

        if (this.frontRepo.array_LogoOnTheRights.length > 0) {

          this.frontRepo.array_LogoOnTheRights.sort((a: split.LogoOnTheRight, b: split.LogoOnTheRight) => {
            return a.ID - b.ID
          })
        }

        for (let logo of this.frontRepo.array_LogoOnTheLefts) {
          if (this.radioButtonHeight < logo.Height) {
            this.radioButtonHeight = logo.Height
          }
        }
        for (let logo of this.frontRepo.array_LogoOnTheRights) {
          if (this.radioButtonHeight < logo.Height) {
            this.radioButtonHeight = logo.Height
          }
        }

        if (this.frontRepo.array_Views.length > 0) {

          this.frontRepo.array_Views.sort((a: split.View, b: split.View) => {
            return a.ID - b.ID
          })

          // in case no view has the field IsSelectedView set to true,
          // the first view is selected
          this.view = this.frontRepo.array_Views[0]
        }

        // set the view to the view with IsSelectedView to true
        for (let view_ of this.frontRepo.array_Views) {
          if (view_.IsSelectedView) {
            this.view = view_
          }
        }

      }
    }
    )
  }

  setSvgFavicon(svgString: string) {
    // Create data URL from SVG string
    const svgBlob = new Blob([svgString], { type: 'image/svg+xml' });
    const svgUrl = URL.createObjectURL(svgBlob);

    // Alternative: use data URL directly
    // const svgUrl = `data:image/svg+xml;base64,${btoa(svgString)}`;

    // Find existing favicon or create new one
    let link: HTMLLinkElement | null = this.document.querySelector("link[rel*='icon']");

    if (link) {
      // Existing favicon found, update it
      link.href = svgUrl;
    } else {
      // No existing favicon, create new one
      link = this.renderer.createElement('link');
      if (link) {
        link.rel = 'icon';
        link.type = 'image/svg+xml';
        link.href = svgUrl;
        this.renderer.appendChild(this.document.head, link);
      }

    }
  }

  getSafeHtml(svgContent: string): SafeHtml {
    return this.sanitizer.bypassSecurityTrustHtml(svgContent);
  }

  get currentView(): split.View | undefined {
    return this.frontRepo?.array_Views.find(v => v === this.view);
  }

  public asSplitDirection(direction: string): 'horizontal' | 'vertical' {
    if (direction === '' || direction === undefined || direction === null) {
      return 'vertical';
    }
    return direction as 'horizontal' | 'vertical';
  }
}