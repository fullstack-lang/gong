import { Component, signal } from '@angular/core';

import { Test4Specific } from '../../projects/test4specific/src/lib/test4-specific/test4-specific';
import { SplitSpecificComponent } from '@vendored_components/github.com/fullstack-lang/gong/lib/split/ng-github.com-fullstack-lang-gong-lib-split/projects/splitspecific/src/lib/split-specific/split-specific.component'

import { AngularSplitModule } from 'angular-split';


@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  standalone: true,
  imports: [Test4Specific, AngularSplitModule, SplitSpecificComponent],
  styles: [],
})
export class App {
  protected readonly title = signal('ng-helloworld');
}
