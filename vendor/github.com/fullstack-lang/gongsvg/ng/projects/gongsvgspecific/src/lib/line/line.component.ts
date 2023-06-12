import { Component, OnInit, Input } from '@angular/core';

import * as gongsvg from 'gongsvg'
@Component({
  selector: 'lib-line',
  templateUrl: './line.component.svg',
  styleUrls: ['./line.component.css']
})
export class LineComponent implements OnInit {

  @Input() Line?: gongsvg.LineDB
  @Input() GONG__StackPath: string = ""

  constructor(private lineService: gongsvg.LineService) { }

  ngOnInit(): void {

  }

  onClick(event: MouseEvent) {
    event.stopPropagation(); // Prevent the event from bubbling up to the SVG element

    console.log("Line on click", this.Line!.Name)

    this.Line!.MouseClickX = event.clientX
    this.Line!.MouseClickY = event.clientY

    this.lineService.updateLine(this.Line!, this.GONG__StackPath).subscribe()

  }

}
