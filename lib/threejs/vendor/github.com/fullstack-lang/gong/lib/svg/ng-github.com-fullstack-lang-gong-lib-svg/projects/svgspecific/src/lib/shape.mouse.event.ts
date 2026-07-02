import * as svg from '../../../svg/src/public-api'

export interface ShapeMouseEvent {
  ShapeID: number
  ShapeType: string
  Point: svg.Point
}