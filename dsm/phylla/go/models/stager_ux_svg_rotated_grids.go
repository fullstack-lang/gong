package models

import (
	"fmt"
	"math"
	"sort"

	svg "github.com/fullstack-lang/gong/lib/svg/go/models"
)

type coord struct {
	x, y float64
}

func (plantDiagram *PlantDiagram) drawRotatedRhombusGridShape(stager *Stager, layer *svg.Layer, plant *Plant) {
	if plantDiagram.RotatedRhombusGridShape == nil || plantDiagram.RotatedRhombusGridShape.IsHidden {
		return
	}

	insideAngleRad := plant.RhombusInsideAngle * math.Pi / 180.0
	sinHalf := math.Sin(insideAngleRad / 2.0)
	cosHalf := math.Cos(insideAngleRad / 2.0)
	sideLength := plant.RhombusSideLength

	deltaY := float64(plant.N)*sideLength*sinHalf + float64(plant.M)*sideLength*sinHalf
	deltaX := float64(plant.N)*sideLength*cosHalf - float64(plant.M)*sideLength*cosHalf

	L := plantDiagram.AxesShape.LengthX

	for i := 0; i < plant.Z; i++ {
		rawX := float64(i) * deltaX
		rawY := float64(i) * deltaY

		var wrappedX float64
		if L > 0 {
			wrappedX = math.Mod(math.Mod(rawX, L)+L, L)
		} else {
			wrappedX = rawX
		}

		cx := plantDiagram.OriginX + wrappedX
		cy := plantDiagram.OriginY - rawY

		v0x := cx
		v0y := cy
		v1x := cx + sideLength*cosHalf
		v1y := cy - sideLength*sinHalf
		v2x := cx + 2.0*sideLength*cosHalf
		v2y := cy
		v3x := cx + sideLength*cosHalf
		v3y := cy + sideLength*sinHalf

		polygon := new(svg.Polygone).Stage(stager.svgStage)
		layer.Polygones = append(layer.Polygones, polygon)
		polygon.Name = fmt.Sprintf("%s-%d", plantDiagram.RotatedRhombusGridShape.Name, i)
		polygon.Points = fmt.Sprintf("%f,%f %f,%f %f,%f %f,%f", v0x, v0y, v1x, v1y, v2x, v2y, v3x, v3y)

		polygon.Presentation.Stroke = "darkgreen" // From screenshot
		polygon.Presentation.StrokeWidth = 1.0
		polygon.Presentation.StrokeOpacity = 1.0
		polygon.Presentation.Color = "lightblue"
		polygon.Presentation.FillOpacity = 0.2 // slight fill like the screenshot
	}
}

func (plantDiagram *PlantDiagram) drawRotatedCircleGridShape(stager *Stager, layer *svg.Layer, plant *Plant) {
	if plantDiagram.RotatedCircleGridShape == nil || plantDiagram.RotatedCircleGridShape.IsHidden {
		return
	}

	insideAngleRad := plant.RhombusInsideAngle * math.Pi / 180.0
	sinHalf := math.Sin(insideAngleRad / 2.0)
	cosHalf := math.Cos(insideAngleRad / 2.0)
	sideLength := plant.RhombusSideLength

	deltaY := float64(plant.N)*sideLength*sinHalf + float64(plant.M)*sideLength*sinHalf
	deltaX := float64(plant.N)*sideLength*cosHalf - float64(plant.M)*sideLength*cosHalf

	L := plantDiagram.AxesShape.LengthX
	radius := sideLength * math.Sin(insideAngleRad) / 2.0

	for i := 0; i < plant.Z; i++ {
		rawX := float64(i) * deltaX
		rawY := float64(i) * deltaY

		var wrappedX float64
		if L > 0 {
			wrappedX = math.Mod(math.Mod(rawX, L)+L, L)
		} else {
			wrappedX = rawX
		}

		// left vertex
		cx := plantDiagram.OriginX + wrappedX
		cy := plantDiagram.OriginY - rawY

		// circle center is rhombus center
		circleX := cx + sideLength*cosHalf
		circleY := cy

		circle := new(svg.Circle).Stage(stager.svgStage)
		layer.Circles = append(layer.Circles, circle)
		circle.Name = fmt.Sprintf("%s-%d", plantDiagram.RotatedCircleGridShape.Name, i)
		circle.CX = circleX
		circle.CY = circleY
		circle.Radius = radius
		circle.Presentation.Stroke = "lightblue" // From screenshot, circles are blueish or greenish
		circle.Presentation.StrokeWidth = 2.0
		circle.Presentation.StrokeOpacity = 1.0
		circle.Presentation.Color = "white"
		circle.Presentation.FillOpacity = 0.0
	}
}

func (plantDiagram *PlantDiagram) drawRotatedNextCircleShape(stager *Stager, layer *svg.Layer, plant *Plant) {
	if plantDiagram.RotatedNextCircleShape == nil || plantDiagram.RotatedNextCircleShape.IsHidden {
		return
	}

	insideAngleRad := plant.RhombusInsideAngle * math.Pi / 180.0
	sinHalf := math.Sin(insideAngleRad / 2.0)
	cosHalf := math.Cos(insideAngleRad / 2.0)
	sideLength := plant.RhombusSideLength

	deltaY := float64(plant.N)*sideLength*sinHalf + float64(plant.M)*sideLength*sinHalf
	deltaX := float64(plant.N)*sideLength*cosHalf - float64(plant.M)*sideLength*cosHalf

	L := plantDiagram.AxesShape.LengthX
	radius := sideLength * math.Sin(insideAngleRad) / 2.0

	var circles []coord
	for i := 0; i < plant.Z; i++ {
		rawX := float64(i) * deltaX
		rawY := float64(i) * deltaY

		var wrappedX float64
		if L > 0 {
			wrappedX = math.Mod(math.Mod(rawX, L)+L, L)
		} else {
			wrappedX = rawX
		}

		cx := plantDiagram.OriginX + wrappedX
		cy := plantDiagram.OriginY - rawY
		circleX := cx + sideLength*cosHalf
		circleY := cy
		circles = append(circles, coord{x: circleX, y: circleY})
	}

	if len(circles) < 3 {
		return // not enough circles to pick the 3rd one
	}

	sort.Slice(circles, func(i, j int) bool {
		return circles[i].y > circles[j].y
	})

	circleX := circles[2].x
	circleY := circles[2].y

	circle := new(svg.Circle).Stage(stager.svgStage)
	layer.Circles = append(layer.Circles, circle)
	circle.Name = plantDiagram.RotatedNextCircleShape.Name
	circle.CX = circleX
	circle.CY = circleY
	circle.Radius = radius
	circle.Presentation.Stroke = "yellow"
	circle.Presentation.StrokeWidth = 4.0
	circle.Presentation.StrokeOpacity = 1.0
	circle.Presentation.Color = "none"
	circle.Presentation.FillOpacity = 0.0
}
