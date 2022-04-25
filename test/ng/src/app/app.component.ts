import { Component } from '@angular/core';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
})
export class AppComponent {

  view = 'Default view'
  default = 'Default view'
  diagrams = 'Diagrams view'
  meta = 'Meta view'

  views: string[] = [this.default, this.diagrams, this.meta];
}
