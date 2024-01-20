import { Injectable } from '@angular/core';
import { Point } from 'gongsvg';
import { Subject } from 'rxjs';

export interface ShapeMouseEvent {
  ShapeID: number
  ShapeType: string
  Point: Point
}