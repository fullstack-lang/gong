import { Injectable } from '@angular/core';
import { PointDB } from 'gongsvg';
import { Subject } from 'rxjs';

export interface ShapeMouseEvent {
  ShapeID: number
  ShapeType: string
  Point: PointDB
}