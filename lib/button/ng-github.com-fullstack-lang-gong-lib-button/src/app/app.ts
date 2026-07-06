import { Component, signal } from '@angular/core';

import { SplitSpecificComponent } from '@vendored_components/github.com/fullstack-lang/gong/lib/split/ng-github.com-fullstack-lang-gong-lib-split/projects/splitspecific/src/lib/split-specific/split-specific.component'
import { ButtonSpecific } from '../../projects/buttonspecific/src/lib/button-specific/button-specific';

@Component({
  selector: 'app-root',
  imports: [SplitSpecificComponent, ButtonSpecific],
  template: `
      <lib-button-specific></lib-button-specific>
      <lib-split-specific></lib-split-specific>
  `,
  styles: [],
})
export class App {
  protected readonly title = signal('ng-button');
}
