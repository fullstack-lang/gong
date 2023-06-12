import { Component, Input, OnInit } from '@angular/core';

import * as gongsvg from 'gongsvg'
@Component({
  selector: 'lib-path',
  templateUrl: './path.component.svg',
  styleUrls: ['./path.component.css']
})
export class PathComponent implements OnInit {

  @Input() Path?: gongsvg.PathDB

  constructor() { }

  ngOnInit(): void {
  }

}
