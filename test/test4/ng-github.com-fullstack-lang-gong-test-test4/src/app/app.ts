import { Component, signal } from '@angular/core';

import { SplitSpecificComponent } from '@vendored_components/github.com/fullstack-lang/gong/lib/split/ng-github.com-fullstack-lang-gong-lib-split/projects/splitspecific/src/lib/split-specific/split-specific.component'

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  standalone: true,
  imports: [SplitSpecificComponent],
  styles: [],
})
export class App {
  protected readonly title = signal('ng-helloworld');
}
