import { Component, signal } from '@angular/core';

import { Test4Specific } from '../../projects/test4specific/src/lib/test4-specific/test4-specific';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  standalone: true,
  imports: [Test4Specific],
  styles: [],
})
export class App {
  protected readonly title = signal('ng-helloworld');
}
