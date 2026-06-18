package models

import (
	"math"
)

// -------------------------------------------------------------------------------------
// Trajectory Math Helpers
// -------------------------------------------------------------------------------------

type Point struct {
	x float64
	y float64
}

func normalize(v Point) Point {
	length := math.Sqrt(v.x*v.x + v.y*v.y)
	return Point{v.x / length, v.y / length}
}

func perpendicular(v Point) Point {
	return Point{-v.y, v.x}
}

func arrowTipPoints(beforeLastPoint, lastPoint Point, length float64) [3]Point {
	direction := Point{lastPoint.x - beforeLastPoint.x, lastPoint.y - beforeLastPoint.y}
	normalizedDirection := normalize(direction)
	perp := perpendicular(normalizedDirection)

	tipPoint := Point{lastPoint.x + normalizedDirection.x*length/1.2, lastPoint.y + normalizedDirection.y*length/1.2}
	leftPoint := Point{lastPoint.x + perp.x*(length/2), lastPoint.y + perp.y*(length/2)}
	rightPoint := Point{lastPoint.x - perp.x*(length/2), lastPoint.y - perp.y*(length/2)}

	return [3]Point{tipPoint, leftPoint, rightPoint}
}

func Newton(n int, k int) float64 {
	var rez float64 = 1
	for i := 1; i < k; i++ {
		rez = rez * float64(n-i+1.0) / float64(i)
	}
	return rez
}

func Bernstein(n int, i int, t float64) float64 {
	var rez = math.Pow(float64(t), float64(i)) * math.Pow(float64(1-t), float64(n-i))
	return Newton(n, i) * rez
}

func generateBezierPoints(points []Point, nbPlots int) []Point {
	if len(points) == 0 {
		return nil
	}
	n := len(points) - 1
	pts := make([]Point, nbPlots+1)

	for z := 0; z <= nbPlots; z++ {
		var current_x, current_y, t float64 = 0.0, 0.0, float64(z) / float64(nbPlots)
		var temp float64
		var denominator float64 = 0.0

		for i := 0; i <= n; i++ {
			temp = points[i].x * Bernstein(n, i, t)
			current_x += temp
			temp = points[i].y * Bernstein(n, i, t)
			current_y += temp
			temp = Bernstein(n, i, t)
			denominator += temp
		}

		pts[z].x = current_x / denominator
		pts[z].y = current_y / denominator
	}

	return pts
}
