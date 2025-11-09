import * as svg from '../../../svg/src/public-api'

export function controlPointToPoint(controlPoint: svg.ControlPoint): svg.Point {
  let result = new (svg.Point)

  result.X = controlPoint.ClosestRect!.X + controlPoint.X_Relative * controlPoint.ClosestRect!.Width
  result.Y = controlPoint.ClosestRect!.Y + controlPoint.Y_Relative * controlPoint.ClosestRect!.Height

  return result
}

export function pointToControlPoint(point: svg.Point, link: svg.Link): svg.ControlPoint {
  let result = new (svg.ControlPoint)

  if (calculateDistance(point.X, point.Y, link.Start!.X, link.Start!.Y) > calculateDistance(point.X, point.Y, link.End!.X, link.End!.Y)) {
    result.ClosestRect = link.End!
  } else {
    result.ClosestRect = link.Start!
  }

  result.X_Relative = (point.X - result.ClosestRect!.X) / (result.ClosestRect!.Width + 0.0001)
  result.Y_Relative = (point.Y - result.ClosestRect!.Y) / (result.ClosestRect!.Height + 0.0001)

  return result
}

export function calculateDistance(x1: number, y1: number, x2: number, y2: number): number {
  const deltaX = x2 - x1;
  const deltaY = y2 - y1;

  // Use Math.hypot() to safely calculate the square root of the sum of squares.
  // This is equivalent to Math.sqrt(deltaX * deltaX + deltaY * deltaY)
  return Math.hypot(deltaX, deltaY);
}