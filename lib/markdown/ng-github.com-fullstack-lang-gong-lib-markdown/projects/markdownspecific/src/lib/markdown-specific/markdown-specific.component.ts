import { Component, Input, ViewChild, ElementRef } from '@angular/core';
import * as markdown from '../../../../markdown/src/public-api';

// 1. IMPORT provideMarkdown again
import { MarkdownModule, provideMarkdown } from 'ngx-markdown';
import { CommonModule } from '@angular/common';
import { MatButtonModule } from '@angular/material/button';

@Component({
  selector: 'lib-markdown-specific',
  imports: [CommonModule, MarkdownModule, MatButtonModule],
  // 2. ADD the providers array back with the default provider
  providers: [provideMarkdown()],
  templateUrl: './markdown-specific.component.html',
  styleUrl: './markdown-specific.component.css'
})
export class MarkdownSpecificComponent {
  // ... the rest of your component code remains the same as the previous step
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
      this.processSvgImages(contentInstance.Content);
    }
  }

  private processSvgImages(markdownContent: string): void {
    if (!this.frontRepo) {
        return;
    }

    const svgImages = this.frontRepo.getFrontArray<markdown.SvgImage>(markdown.SvgImage.GONGSTRUCT_NAME);
    const svgImageMap = new Map(svgImages.map(img => [img.Name, img.Content]));

    const svgRegex = /!\[(.*?)\]\(svg:([^?)]+?)(?:\?(.*?))?\)/g;

    this.content = markdownContent.replace(svgRegex, (match, altText, svgName, queryParams) => {
        const trimmedSvgName = svgName.trim();
        let svgContent = svgImageMap.get(trimmedSvgName);

        if (svgContent) {
            let widthValue = '';
            if (queryParams) {
                const params = new URLSearchParams(queryParams);
                widthValue = params.get('width') || '';
            }

            if (widthValue) {
                const svgTagRegex = /<svg([^>]*?)>/;
                const tagMatch = svgContent.match(svgTagRegex);

                if (tagMatch) {
                    let attributes = tagMatch[1] || '';
                    attributes = attributes.replace(/width="[^"]*"/g, '').trim();
                    const newAttributes = `width="${widthValue}" ${attributes}`;
                    const newSvgTag = `<svg ${newAttributes}>`;
                    svgContent = svgContent.replace(svgTagRegex, newSvgTag);
                }
            }

            const encodedSvg = btoa(svgContent);
            const dataUri = `data:image/svg+xml;base64,${encodedSvg}`;
            
            return `![${altText}](${dataUri})`;
        }
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
          img { max-width: 100%; height: auto; display: block; }
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