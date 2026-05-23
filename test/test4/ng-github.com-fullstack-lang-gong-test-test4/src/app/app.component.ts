import { Component } from '@angular/core';
import { Test4Specific } from "../../projects/test4specific/src/lib/test4-specific/test4-specific";


@Component({
  selector: 'app-root',
  standalone: true,
  imports: [
    Test4Specific
],

  templateUrl: './app.component.html',
})
export class AppComponent {
}
