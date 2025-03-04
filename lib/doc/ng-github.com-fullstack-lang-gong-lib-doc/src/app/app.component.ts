import { Component, OnInit, ViewChild } from '@angular/core';

import { DocSpecificComponent } from '../../projects/docspecific/src/lib/doc-specific/doc-specific.component'

import * as doc from '../../projects/doc/src/public-api'

@Component({
    selector: 'app-root',
    templateUrl: './app.component.html',
    imports: [
      DocSpecificComponent,
    ]
})
export class AppComponent implements OnInit{
  StackType = doc.StackType

  ngOnInit(): void {
    console.log("AppComponent:ngOnInit")
  }
}
