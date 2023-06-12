import { Component, OnInit, Input } from '@angular/core';

import * as gongsvg from 'gongsvg'
@Component({
  selector: 'lib-text',
  templateUrl: './text.component.svg',
  styleUrls: ['./text.component.css']
})
export class TextComponent implements OnInit {

  @Input() Text?: gongsvg.TextDB

  @Input() X?: number
  @Input() Y?: number
  @Input() Content?: string

  @Input() Color?: string
  @Input() Stroke?: string
  @Input() StrokeWidth?: string
  @Input() FillOpacity?: number
  @Input() StrokeDashArray?: string
  @Input() Transform?: string

  constructor() { }

  ngOnInit(): void {
  }

}
