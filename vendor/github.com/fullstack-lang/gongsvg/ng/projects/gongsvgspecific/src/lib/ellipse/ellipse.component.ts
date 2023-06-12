import { Component, Input, OnInit } from '@angular/core';

import * as gongsvg from 'gongsvg'
@Component({
  selector: 'lib-ellipse',
  templateUrl: './ellipse.component.svg',
  styleUrls: ['./ellipse.component.css']
})
export class EllipseComponent implements OnInit {

  @Input() Ellipse?: gongsvg.EllipseDB

  constructor() { }

  ngOnInit(): void {
  }

}
