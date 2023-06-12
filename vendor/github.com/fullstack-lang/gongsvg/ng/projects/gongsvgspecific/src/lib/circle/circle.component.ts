import { Component, OnInit, Input } from '@angular/core';

import * as gongsvg from 'gongsvg'
import { CircleDB } from 'gongsvg';
@Component({
  selector: 'lib-circle',
  templateUrl: './circle.component.svg',
  styleUrls: ['./circle.component.css']
})
export class CircleComponent implements OnInit {

  @Input() Circle?: gongsvg.CircleDB

  constructor() { }

  ngOnInit(): void {
  }

  onClick(event: MouseEvent) {
    console.log("Circle: ", this.Circle?.Name, " clicked")
    // event.stopPropagation(); // Prevent the event from bubbling up to the SVG element
  }
}
