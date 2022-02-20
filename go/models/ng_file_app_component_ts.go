package models

const NgFileAppComponentTs = `import { Component } from '@angular/core';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
})
export class AppComponent {

  // choices for the top radio button
  view = 'Default view'
  default = 'Default view'
  diagrams = 'Diagrams view'
  meta = 'Meta view'
  views: string[] = [this.default, this.diagrams, this.meta];
}
`
