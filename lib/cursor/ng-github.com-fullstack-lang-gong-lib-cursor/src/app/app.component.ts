import { Component, OnInit } from '@angular/core';

import { AngularSplitModule } from 'angular-split';

import * as cursor from '../../projects/cursor/src/public-api'

import { CursorSpecificComponent } from '../../projects/cursorspecific/src/lib/cursor-specific/cursor-specific.component'

import { SplitSpecificComponent } from '@vendored_components/github.com/fullstack-lang/gong/lib/split/ng-github.com-fullstack-lang-gong-lib-split/projects/splitspecific/src/lib/split-specific/split-specific.component'

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [
    AngularSplitModule,
    SplitSpecificComponent,
    CursorSpecificComponent
  ],

  templateUrl: './app.component.html',
})
export class AppComponent {

}
