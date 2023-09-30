import { Component, Input, OnInit } from '@angular/core';

import * as gongdoc from 'gongdoc';

@Component({
  selector: 'lib-panel',
  templateUrl: './panel.component.html',
  styleUrls: ['./panel.component.css']
})
export class PanelComponent implements OnInit {

  // choices for the top radio button
  view = 'Diagrams'
  diagrams = 'Diagrams'
  views: string[] = [this.diagrams];

  stacks: string[] = []

  loading = true

  @Input() GONG__StackPath: string = ""

  constructor(
  ) {

  }

  ngOnInit(): void {

  }

  refresh() {
  }
}
