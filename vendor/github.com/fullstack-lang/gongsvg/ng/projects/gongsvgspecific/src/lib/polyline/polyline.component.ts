import { Component, Input, OnInit } from '@angular/core';

import * as gongsvg from 'gongsvg'
@Component({
  selector: 'lib-polyline',
  templateUrl: './polyline.component.svg',
  styleUrls: ['./polyline.component.css']
})
export class PolylineComponent implements OnInit {

  @Input() Polyline?: gongsvg.PolylineDB

  constructor() { }

  ngOnInit(): void {
  }

}
