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
      this.processImages(contentInstance.Content);
    }
  }

  private processImages(markdownContent: string): void {
    if (!this.frontRepo) {
        return;
    }

    let processedContent = markdownContent;

    // Process SVGs
    const svgImages = this.frontRepo.getFrontArray<markdown.SvgImage>(markdown.SvgImage.GONGSTRUCT_NAME);
    const svgImageMap = new Map(svgImages.map(img => [img.Name, img.Content]));
    const svgRegex = /!\[(.*?)\]\(svg:([^?)]+?)(?:\?(.*?))?\)/g;

    processedContent = processedContent.replace(svgRegex, (match, altText, svgName, queryParams) => {
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
                if (svgTagRegex.test(svgContent)) {
                    svgContent = svgContent.replace(svgTagRegex, (tagMatch, attributes) => {
                        let newAttributes = attributes.replace(/width="[^"]*"/g, '').trim();
                        return `<svg width="${widthValue}" ${newAttributes}>`;
                    });
                }
            }
            
            const encodedSvg = btoa(svgContent);
            const dataUri = `data:image/svg+xml;base64,${encodedSvg}`;
            return `![${altText}](${dataUri})`;
        }
        return match;
    });

    // Process PNGs
    const pngImages = this.frontRepo.getFrontArray<markdown.PngImage>(markdown.PngImage.GONGSTRUCT_NAME);
    const pngImageMap = new Map(pngImages.map(img => [img.Name, img.Base64Content]));
    const pngRegex = /!\[(.*?)\]\(png:([^?)]+?)(?:\?(.*?))?\)/g;

    processedContent = processedContent.replace(pngRegex, (match, altText, pngName, queryParams) => {
      const trimmedPngName = pngName.trim();
      const base64Content = pngImageMap.get(trimmedPngName);

      if (base64Content) {
        const dataUri = `data:image/png;base64,${base64Content}`;
        return `![${altText}](${dataUri})`;
      }
      return match;
    });

    // Process JPGs
    const jpgImages = this.frontRepo.getFrontArray<markdown.JpgImage>(markdown.JpgImage.GONGSTRUCT_NAME);
    const jpgImageMap = new Map(jpgImages.map(img => [img.Name, img.Base64Content]));
    const jpgRegex = /!\[(.*?)\]\(jpg:([^?)]+?)(?:\?(.*?))?\)/g;

    processedContent = processedContent.replace(jpgRegex, (match, altText, jpgName, queryParams) => {
      const trimmedJpgName = jpgName.trim();
      const base64Content = jpgImageMap.get(trimmedJpgName);

      if (base64Content) {
        const dataUri = `data:image/jpeg;base64,${base64Content}`;
        return `![${altText}](${dataUri})`;
      }
      return match;
    });

    this.content = processedContent;
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