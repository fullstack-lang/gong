import { Component, signal } from '@angular/core';

import { SplitSpecificComponent } from '@vendored_components/github.com/fullstack-lang/gong/lib/split/ng-github.com-fullstack-lang-gong-lib-split/projects/splitspecific/src/lib/split-specific/split-specific.component'
import { GongthreejsSpecific } from '../../projects/gong-lib-threejsspecific/src/lib/gongthreejs-specific/gongthreejs-specific';

@Component({
  selector: 'app-root',
  imports: [SplitSpecificComponent, GongthreejsSpecific],
  template: `
      <lib-threejs-specific Name="threejs"></lib-threejs-specific>
      <lib-split-specific></lib-split-specific>
  `,
  styles: [],
})
export class App {
  protected readonly title = signal('ng-gong-lib-threejs');
}
