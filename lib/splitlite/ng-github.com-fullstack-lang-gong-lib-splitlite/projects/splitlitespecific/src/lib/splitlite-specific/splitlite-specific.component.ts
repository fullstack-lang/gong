import { Component, Renderer2, Input, OnInit, Inject, DOCUMENT, ChangeDetectorRef, forwardRef } from '@angular/core';

import * as splitlite from '../../../../splitlite/src/public-api';

import { FormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { MatButtonModule } from '@angular/material/button';
import { MatMenuModule } from '@angular/material/menu';

import { GongSplitComponent } from '../native-split/gong-split.component';
import { GongSplitAreaComponent } from '../native-split/gong-split-area.component';

import { ButtonSpecificComponent } from '../../../../../../../button/ng-github.com-fullstack-lang-gong-lib-button/projects/buttonspecific/src/lib/button-specific/button-specific.component';
import { FormSpecific } from '../../../../../../../form/ng-github.com-fullstack-lang-gong-lib-form/projects/formspecific/src/lib/form-specific/form-specific';
import { LoadSpecificComponent } from '../../../../../../../load/ng-github.com-fullstack-lang-gong-lib-load/projects/loadspecific/src/lib/load-specific/load-specific.component';
import { SvgSpecificComponent } from '../../../../../../../svg/ng-github.com-fullstack-lang-gong-lib-svg/projects/svgspecific/src/lib/svg-specific/svg-specific.component';
import { TableSpecificComponent } from '../../../../../../../table/ng-github.com-fullstack-lang-gong-lib-table/projects/tablespecific/src/lib/table-specific/table-specific.component';
import { TreeSpecificComponent } from '../../../../../../../tree/ng-github.com-fullstack-lang-gong-lib-tree/projects/treespecific/src/lib/tree-specific/tree-specific.component';

// to set the title of the application
import { DomSanitizer, SafeHtml, Title } from '@angular/platform-browser';

@Component({
  selector: 'lib-splitlite-specific',
  standalone: true,
  imports: [
    CommonModule,
    FormsModule,
    MatButtonModule,
    MatMenuModule,

    GongSplitComponent,
    GongSplitAreaComponent,

    ButtonSpecificComponent,
    FormSpecific,
    LoadSpecificComponent,
    SvgSpecificComponent,
    TableSpecificComponent,
    TreeSpecificComponent,

    forwardRef(() => SplitliteSpecificComponent),
  ],
  templateUrl: './splitlite-specific.component.html',
  styleUrl: './splitlite-specific.component.css'
})
export class SplitliteSpecificComponent implements OnInit {
  @Input() Name: string = "";

  public frontRepo?: splitlite.FrontRepo;
  public view: splitlite.View | undefined;
  topBarHeight = 56;

  constructor(
    private frontRepoService: splitlite.FrontRepoService,
    private viewService: splitlite.ViewService,
    private titleService: Title,
    private renderer: Renderer2,
    private sanitizer: DomSanitizer,
    private cdr: ChangeDetectorRef,
    @Inject(DOCUMENT) private document: Document
  ) { }

  selectView(selectedView: splitlite.View) {
    this.view = selectedView;
    if (this.frontRepo && this.frontRepo.array_Views) {
      for (let v of this.frontRepo.array_Views) {
        v.IsSelectedView = (v === selectedView);
      }
    }
    this.cdr.markForCheck();
    this.viewService.updateFront(selectedView, this.Name).subscribe(
      () => {
        // front update sent to backend
      }
    );
  }

  ngOnInit(): void {
    this.frontRepoService.connectToWebSocket(this.Name).subscribe({
      next: (frontRepo) => {
        this.frontRepo = frontRepo;

        if (this.frontRepo.array_Titles && this.frontRepo.array_Titles.length > 0) {
          this.titleService.setTitle(this.frontRepo.array_Titles[0].Name);
        }

        if (this.frontRepo.array_FavIcons && this.frontRepo.array_FavIcons.length > 0) {
          this.setSvgFavicon(this.frontRepo.array_FavIcons[0].SVG);
        }

        if (this.frontRepo.array_LogoOnTheLefts && this.frontRepo.array_LogoOnTheLefts.length > 0) {
          this.frontRepo.array_LogoOnTheLefts.sort((a: splitlite.LogoOnTheLeft, b: splitlite.LogoOnTheLeft) => a.ID - b.ID);
        }

        if (this.frontRepo.array_LogoOnTheRights && this.frontRepo.array_LogoOnTheRights.length > 0) {
          this.frontRepo.array_LogoOnTheRights.sort((a: splitlite.LogoOnTheRight, b: splitlite.LogoOnTheRight) => a.ID - b.ID);
        }

        if (this.frontRepo.array_LogoOnTheLefts) {
          for (let logo of this.frontRepo.array_LogoOnTheLefts) {
            if (this.topBarHeight < logo.Height) {
              this.topBarHeight = logo.Height;
            }
          }
        }

        if (this.frontRepo.array_LogoOnTheRights) {
          for (let logo of this.frontRepo.array_LogoOnTheRights) {
            if (this.topBarHeight < logo.Height) {
              this.topBarHeight = logo.Height;
            }
          }
        }

        if (this.frontRepo.array_Views && this.frontRepo.array_Views.length > 0) {
          this.frontRepo.array_Views.sort((a: splitlite.View, b: splitlite.View) => a.ID - b.ID);
          this.view = this.frontRepo.array_Views[0];

          for (let view_ of this.frontRepo.array_Views) {
            if (view_.IsSelectedView) {
              this.view = view_;
            }
          }
        }

        this.cdr.markForCheck();
      },
      error: (err) => {
        console.error(`[lib-splitlite-specific] connectToWebSocket failed for Name: "${this.Name}"`, err);
      }
    });
  }

  setSvgFavicon(svgString: string) {
    const svgBlob = new Blob([svgString], { type: 'image/svg+xml' });
    const svgUrl = URL.createObjectURL(svgBlob);

    let link: HTMLLinkElement | null = this.document.querySelector("link[rel*='icon']");
    if (link) {
      link.href = svgUrl;
    } else {
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

  get currentView(): splitlite.View | undefined {
    return this.frontRepo?.array_Views.find(v => v === this.view);
  }

  isSecondaryViewSelected(): boolean {
    return this.view?.IsSecondaryView ?? false;
  }

  public asSplitDirection(direction: string): 'horizontal' | 'vertical' {
    if (direction === '' || direction === undefined || direction === null) {
      return 'vertical';
    }
    return direction as 'horizontal' | 'vertical';
  }
}
