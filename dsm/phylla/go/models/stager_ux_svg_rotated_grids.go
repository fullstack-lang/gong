package models

import (
	"fmt"
	"math"
	"sort"

	svg "github.com/fullstack-lang/gong/lib/svg/go/models"
)

func (plantDiagram *PlantDiagram) drawRotatedRhombusGridShape(stager *Stager, layer *svg.Layer, plant *Plant) {
	if plantDiagram.RotatedRhombusGridShape == nil || plantDiagram.RotatedRhombusGridShape.IsHidden {
		return
	}

	insideAngleRad := plant.RhombusInsideAngle * math.Pi / 180.0
	sinHalf := math.Sin(insideAngleRad / 2.0)
	cosHalf := math.Cos(insideAngleRad / 2.0)
	sideLength := plant.RhombusSideLength
	angleRad := plantDiagram.GrowthVectorShape.AngleDegree * math.Pi / 180.0

	for i := 0; i <= plant.N; i++ {
		for j := 0; j <= plant.M; j++ {
			rawX := float64(i)*sideLength*cosHalf - float64(j)*sideLength*cosHalf
			rawY := float64(i)*sideLength*sinHalf + float64(j)*sideLength*sinHalf

			// Rotation by -angleRad to find the Cartesian Y coordinate of the center in the rotated frame
			rotY := rawX*math.Sin(-angleRad) + rawY*math.Cos(-angleRad)
			if rotY < -0.00001 {
				continue // discard shapes below the growth vector axis
			}

			// Center of the rhombus
			cx := plantDiagram.OriginX + rawX
			cy := plantDiagram.OriginY - rawY

			// Vertices relative to the center
			v0x := cx - sideLength*cosHalf
			v0y := cy
			v1x := cx
			v1y := cy - sideLength*sinHalf
			v2x := cx + sideLength*cosHalf
			v2y := cy
			v3x := cx
			v3y := cy + sideLength*sinHalf

			polygon := new(svg.Polygone).Stage(stager.svgStage)
			layer.Polygones = append(layer.Polygones, polygon)
			polygon.Name = fmt.Sprintf("%s-%d-%d", plantDiagram.RotatedRhombusGridShape.Name, i, j)
			polygon.Points = fmt.Sprintf("%f,%f %f,%f %f,%f %f,%f", v0x, v0y, v1x, v1y, v2x, v2y, v3x, v3y)

			polygon.Presentation.Stroke = "darkgreen"
			polygon.Presentation.StrokeWidth = 1.0
			polygon.Presentation.StrokeOpacity = 1.0
			polygon.Presentation.Color = "lightblue"
			polygon.Presentation.FillOpacity = 0.2
			polygon.Presentation.Transform = fmt.Sprintf("rotate(%f %f %f)", plantDiagram.GrowthVectorShape.AngleDegree, plantDiagram.OriginX, plantDiagram.OriginY)
		}
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
	radius := sideLength * math.Sin(insideAngleRad) / 2.0
	angleRad := plantDiagram.GrowthVectorShape.AngleDegree * math.Pi / 180.0

	for i := 0; i <= plant.N; i++ {
		for j := 0; j <= plant.M; j++ {
			rawX := float64(i)*sideLength*cosHalf - float64(j)*sideLength*cosHalf
			rawY := float64(i)*sideLength*sinHalf + float64(j)*sideLength*sinHalf

			rotY := rawX*math.Sin(-angleRad) + rawY*math.Cos(-angleRad)
			if rotY < -0.00001 {
				continue
			}

			// Center of the circle
			cx := plantDiagram.OriginX + rawX
			cy := plantDiagram.OriginY - rawY

			circle := new(svg.Circle).Stage(stager.svgStage)
			layer.Circles = append(layer.Circles, circle)
			circle.Name = fmt.Sprintf("%s-%d-%d", plantDiagram.RotatedCircleGridShape.Name, i, j)
			circle.CX = cx
			circle.CY = cy
			circle.Radius = radius
			circle.Presentation.Stroke = "lightblue"
			circle.Presentation.StrokeWidth = 2.0
			circle.Presentation.StrokeOpacity = 1.0
			circle.Presentation.Color = "white"
			circle.Presentation.FillOpacity = 0.0
			circle.Presentation.Transform = fmt.Sprintf("rotate(%f %f %f)", plantDiagram.GrowthVectorShape.AngleDegree, plantDiagram.OriginX, plantDiagram.OriginY)
		}
	}
}

type rotatedCircleData struct {
	circleX, circleY float64
	rotY             float64
}

func (plantDiagram *PlantDiagram) drawRotatedNextCircleShape(stager *Stager, layer *svg.Layer, plant *Plant) {
	if plantDiagram.RotatedNextCircleShape == nil || plantDiagram.RotatedNextCircleShape.IsHidden {
		return
	}

	insideAngleRad := plant.RhombusInsideAngle * math.Pi / 180.0
	sinHalf := math.Sin(insideAngleRad / 2.0)
	cosHalf := math.Cos(insideAngleRad / 2.0)
	sideLength := plant.RhombusSideLength
	radius := sideLength * math.Sin(insideAngleRad) / 2.0
	angleRad := plantDiagram.GrowthVectorShape.AngleDegree * math.Pi / 180.0

	var circles []rotatedCircleData

	for i := 0; i <= plant.N; i++ {
		for j := 0; j <= plant.M; j++ {
			rawX := float64(i)*sideLength*cosHalf - float64(j)*sideLength*cosHalf
			rawY := float64(i)*sideLength*sinHalf + float64(j)*sideLength*sinHalf

			rotY := rawX*math.Sin(-angleRad) + rawY*math.Cos(-angleRad)
			if rotY < -0.00001 {
				continue
			}

			cx := plantDiagram.OriginX + rawX
			cy := plantDiagram.OriginY - rawY

			circles = append(circles, rotatedCircleData{circleX: cx, circleY: cy, rotY: rotY})
		}
	}

	if len(circles) < 3 {
		return
	}

	sort.Slice(circles, func(i, j int) bool {
		return circles[i].rotY < circles[j].rotY
	})

	circle := new(svg.Circle).Stage(stager.svgStage)
	layer.Circles = append(layer.Circles, circle)
	circle.Name = plantDiagram.RotatedNextCircleShape.Name
	circle.CX = circles[2].circleX
	circle.CY = circles[2].circleY
	circle.Radius = radius
	circle.Presentation.Stroke = "yellow"
	circle.Presentation.StrokeWidth = 4.0
	circle.Presentation.StrokeOpacity = 1.0
	circle.Presentation.Color = "none"
	circle.Presentation.FillOpacity = 0.0
	circle.Presentation.Transform = fmt.Sprintf("rotate(%f %f %f)", plantDiagram.GrowthVectorShape.AngleDegree, plantDiagram.OriginX, plantDiagram.OriginY)
}
