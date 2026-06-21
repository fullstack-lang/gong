import { Component } from '@angular/core';

import { SplitSpecificComponent } from '@vendored_components/github.com/fullstack-lang/gong/lib/split/ng-github.com-fullstack-lang-gong-lib-split/projects/splitspecific/src/lib/split-specific/split-specific.component'
import { AngularSplitModule } from 'angular-split';
import { TestSpecificComponent } from "../../projects/testspecific/src/lib/test-specific/test-specific.component";

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [
    SplitSpecificComponent,
    AngularSplitModule,
    TestSpecificComponent
  ],

  templateUrl: './app.component.html',
})
export class AppComponent {
}
