import { Component, Input, ViewChild, ElementRef } from '@angular/core';

import * as markdown from '../../../../markdown/src/public-api';

import { MarkdownModule, provideMarkdown } from 'ngx-markdown';
import { CommonModule } from '@angular/common';
import { MatButtonModule } from '@angular/material/button';

@Component({
  selector: 'lib-markdown-specific',
  imports: [CommonModule, MarkdownModule, MatButtonModule],
  providers: [provideMarkdown()],
  templateUrl: './markdown-specific.component.html',
  styleUrl: './markdown-specific.component.css'
})
export class MarkdownSpecificComponent {
  @Input() Name: string = ""
  @ViewChild('markdownWrapper', { read: ElementRef }) markdownWrapper!: ElementRef;

  frontRepo: markdown.FrontRepo | undefined

  content = ``
  contentName = ``

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
      const contentInstance = this.frontRepo.getFrontArray<markdown.Content>(markdown.Content.GONGSTRUCT_NAME)[0];
      this.contentName = contentInstance.Name;

      // Process the content to embed SVG images
      this.processSvgImages(contentInstance.Content);
    }
  }

  private processSvgImages(markdownContent: string): void {
    if (!this.frontRepo) {
      return;
    }

    const svgImages = this.frontRepo.getFrontArray<markdown.SvgImage>(markdown.SvgImage.GONGSTRUCT_NAME);
    const svgImageMap = new Map(svgImages.map(img => [img.Name, img.Content]));

    // Regex to find all instances of ![alt](svg:svg-name)
    const svgRegex = /!\[(.*?)\]\(svg:(.*?)\)/g;

    this.content = markdownContent.replace(svgRegex, (match, altText, svgName) => {
      const svgContent = svgImageMap.get(svgName);
      if (svgContent) {
        // Embed SVG as a data URI
        const encodedSvg = btoa(svgContent);
        return `![${altText}](data:image/svg+xml;base64,${encodedSvg})`;
      }
      // If SVG is not found, leave the original reference
      return match;
    });
  }


  downloadHtml(): void {
    if (!this.markdownWrapper?.nativeElement && this.content != undefined) {
      console.error('Markdown container not found.');
      return;
    }

    if (this.content == undefined) {
      console.error('No content to download.');
      return;
    }

    const renderedHtml = this.markdownWrapper.nativeElement.innerHTML;

    // Create a full HTML document for better standalone viewing
    const fullHtml = `
      <!DOCTYPE html>
      <html lang="en">
      <head>
        <meta charset="UTF-8">
        <title>${this.contentName || 'Markdown Content'}</title>
        <style>
          body { font-family: sans-serif; line-height: 1.6; padding: 20px; max-width: 800px; margin: 0 auto; }
          pre, code { background-color: #f4f4f4; padding: 2px 5px; border-radius: 4px; }
          pre { padding: 1em; }
          blockquote { border-left: 4px solid #ddd; padding-left: 1em; color: #666; }
        </style>
      </head>
      <body>
        ${renderedHtml}
      </body>
      </html>
    `;

    const blob = new Blob([fullHtml], { type: 'text/html' });
    const url = window.URL.createObjectURL(blob);
    const a = document.createElement('a');
    a.href = url;
    a.download = `${this.contentName || 'markdown'}.html`;
    document.body.appendChild(a);
    a.click();
    document.body.removeChild(a);
    window.URL.revokeObjectURL(url);
  }
}