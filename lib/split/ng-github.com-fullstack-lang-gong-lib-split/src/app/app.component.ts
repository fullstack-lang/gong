import { Component, OnInit } from '@angular/core';

import { SplitSpecificComponent } from '../../projects/splitspecific/src/lib/split-specific/split-specific.component'


@Component({
  selector: 'app-root',
  standalone: true,
  imports: [
    SplitSpecificComponent,
  ],

  templateUrl: './app.component.html',
})
export class AppComponent implements OnInit {
  
  constructor(
  ) {
  }

  ngOnInit(): void {
  }
}
