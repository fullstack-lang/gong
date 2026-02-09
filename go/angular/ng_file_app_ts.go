package angular

const NgFileAppTs = `import { Component, signal } from '@angular/core';

import { SplitSpecificComponent } from '@vendored_components/github.com/fullstack-lang/gong/lib/split/ng-github.com-fullstack-lang-gong-lib-split/projects/splitspecific/src/lib/split-specific/split-specific.component'

@Component({
  selector: 'app-root',
  imports: [SplitSpecificComponent],
  template: ` + "`" + `<lib-split-specific></lib-split-specific>` + "`" + `,
  styles: [],
})
export class App {
  protected readonly title = signal('ng-helloworld');
}
`
