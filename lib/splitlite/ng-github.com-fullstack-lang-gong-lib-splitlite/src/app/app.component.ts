import { Component } from '@angular/core';
import { SplitliteSpecificComponent } from 'splitlitespecific';

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [SplitliteSpecificComponent],
  templateUrl: './app.component.html'
})
export class AppComponent {
  title = 'splitlite';
}
