import { Component, Input, OnInit } from '@angular/core';

import * as gongsvg from 'gongsvg'
@Component({
  selector: 'lib-polygone',
  templateUrl: './polygone.component.svg',
  styleUrls: ['./polygone.component.css']
})
export class PolygoneComponent implements OnInit {

  @Input() Polygone?: gongsvg.PolygoneDB

  constructor() { }

  ngOnInit(): void {
  }

}
