import { Injectable } from '@angular/core';
import { MatIconRegistry } from '@angular/material/icon';
import { DomSanitizer } from '@angular/platform-browser';

@Injectable({ providedIn: 'root' })
export class IconService {
  private registeredIcons = new Set<string>();

  constructor(
    private matIconRegistry: MatIconRegistry,
    private domSanitizer: DomSanitizer
  ) { }

  public registerIcon(iconName: string, svgContent: string): void {
    if (!this.registeredIcons.has(iconName)) {
      this.matIconRegistry.addSvgIconLiteral(
        iconName,
        this.domSanitizer.bypassSecurityTrustHtml(svgContent)
      );
      this.registeredIcons.add(iconName);
    }
  }

  public isIconRegistered(iconName: string): boolean {
    return this.registeredIcons.has(iconName);
  }
}
