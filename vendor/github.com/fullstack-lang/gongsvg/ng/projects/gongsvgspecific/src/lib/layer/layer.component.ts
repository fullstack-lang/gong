import { Component, Input, OnDestroy, OnInit, ViewChild } from '@angular/core';
import { Observable, Subscription, timer } from 'rxjs';
import { Coordinate, RectangleEventService } from '../rectangle-event.service';

import * as gongsvg from 'gongsvg'

@Component({
  selector: 'lib-layer',
  templateUrl: './layer.component.html',
  styleUrls: ['./layer.component.css']
})
export class LayerComponent implements OnInit {

  @Input() GONG__StackPath: string = ""
  @Input() Layer?: gongsvg.LayerDB

  ngOnInit(): void {

    // console.log("Layer component->ngOnInit : Layer, ", this.Layer?.Name, ", GONG__StackPath, " + this.GONG__StackPath)
  }

  onDivClick(event: MouseEvent) {
    console.log("Layer: on div click")
  }
}
