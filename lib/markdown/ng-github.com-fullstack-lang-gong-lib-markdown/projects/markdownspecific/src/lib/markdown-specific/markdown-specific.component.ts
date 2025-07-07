import { Component, Input } from '@angular/core';

import * as markdown from '../../../../markdown/src/public-api'

import { MarkdownModule, provideMarkdown } from 'ngx-markdown';
import { CommonModule } from '@angular/common';


@Component({
  selector: 'lib-markdown-specific',
  imports: [CommonModule, MarkdownModule],
  providers: [provideMarkdown()],
  templateUrl: './markdown-specific.component.html',
  styleUrl: './markdown-specific.component.css'
})
export class MarkdownSpecificComponent {
  @Input() Name: string = ""
  frontRepo: markdown.FrontRepo | undefined

  content = ``

  constructor(
    private frontRepoService: markdown.FrontRepoService,
  ) { }

  ngOnInit(): void {
    this.frontRepoService.connectToWebSocket(this.Name).subscribe(
      frontRepo => {
        this.frontRepo = frontRepo

        this.refresh()
      }
    )
  }

  refresh(): void {
    if (this.frontRepo === undefined) {
      return
    }
    if (this.frontRepo.getFrontArray(markdown.Content.GONGSTRUCT_NAME).length == 1) {
      let content = this.frontRepo.getFrontArray<markdown.Content>(markdown.Content.GONGSTRUCT_NAME)[0]
      this.content = content.Content
    }
  }
}
