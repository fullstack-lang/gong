import { Component } from '@angular/core';

import { SplitSpecificComponent } from '@vendored_components/github.com/fullstack-lang/gong/lib/split/ng-github.com-fullstack-lang-gong-lib-split/projects/splitspecific/src/lib/split-specific/split-specific.component'
import { GongSplitComponent, GongSplitAreaComponent } from '@vendored_components/github.com/fullstack-lang/gong/lib/split/ng-github.com-fullstack-lang-gong-lib-split/projects/splitspecific/src/public-api';
import { TestSpecificComponent } from "../../projects/testspecific/src/lib/test-specific/test-specific.component";

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [
    SplitSpecificComponent,
    GongSplitComponent,
    GongSplitAreaComponent,
    TestSpecificComponent
  ],

  templateUrl: './app.component.html',
})
export class AppComponent {
}
