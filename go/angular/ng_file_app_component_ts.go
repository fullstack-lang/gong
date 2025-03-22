package angular

const NgFileAppComponentTs = `import { Component } from '@angular/core';

import { SplitSpecificComponent } from '@vendored_components/github.com/fullstack-lang/gong/lib/split/ng-github.com-fullstack-lang-gong-lib-split/projects/splitspecific/src/lib/split-specific/split-specific.component'

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [
    SplitSpecificComponent,
  ],

  templateUrl: './app.component.html',
})
export class AppComponent {
}
`
