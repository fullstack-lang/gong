package models

import (
	"fmt"
	"math"
	"strings"

	svg "github.com/fullstack-lang/gong/lib/svg/go/models"
)

func (stager *Stager) ux_svg_plant_diagram() {

	stager.svgStage.Reset()

	var plantDiagram *PlantDiagram
	{
		for d_ := range *GetGongstructInstancesSet[PlantDiagram](stager.stage) {
			if d_.IsChecked {
				plantDiagram = d_
			}
		}
	}

	var plant *Plant
	if plantDiagram != nil {
		for p_ := range *GetGongstructInstancesSet[Plant](stager.stage) {
			for _, d_ := range p_.PlantDiagrams {
				if d_ == plantDiagram {
					plant = p_
				}
			}
		}
	}

	if plantDiagram == nil || plant == nil {
		stager.svgStage.Commit()
		return
	}
	svgObject := stager.generateSvgObject(plantDiagram, plant)

	svg.StageBranch(stager.svgStage, svgObject)
	stager.svgObject = svgObject

	stager.svgStage.Commit()
}

func (stager *Stager) generateSvgObject(plantDiagram *PlantDiagram, plant *Plant) (svg_ *svg.SVG) {

	svg_ = new(svg.SVG)
	svg_.Name = "Plant Diagram" // or any name, if name is an attribute.
	svg_.IsEditable = true

	layer := &svg.Layer{Name: `Axis Shape Layer`}
	svg_.Layers = append(svg_.Layers, layer)

	// creation of 2 transparant rects, one at each ends of the vertical
	plantDiagram.drawAxes(stager, layer)
	plantDiagram.drawGrowthVectorShape(stager, layer)
	plantDiagram.drawReferenceRhombus(stager, layer, plant)
	plantDiagram.drawGridPathShape(stager, layer, plant)

	return
}

const AxisHandleBorderLength = 25

func (plantDiagram *PlantDiagram) drawAxes(stager *Stager, layer *svg.Layer) {

	if plantDiagram.AxesShape.IsHidden {
		return
	}

	handleSize := float64(AxisHandleBorderLength)
	if plantDiagram.AxesShape.GetIsWithHiddenHandle() {
		handleSize = 0.0
	}

	verticalAxisTopHandle := new(svg.Rect)
	verticalAxisTopHandle.Name = "Vertical axis bottom handle"
	if !plantDiagram.AxesShape.GetIsWithHiddenHandle() {
		layer.Rects = append(layer.Rects, verticalAxisTopHandle)
	}
	verticalAxisTopHandle.X = plantDiagram.OriginX - handleSize/2.0
	verticalAxisTopHandle.Y = plantDiagram.OriginY - plantDiagram.AxesShape.LengthY - handleSize
	verticalAxisTopHandle.Width = handleSize
	verticalAxisTopHandle.Height = handleSize
	verticalAxisTopHandle.CanMoveVerticaly = true
	verticalAxisTopHandle.CanMoveHorizontaly = true
	verticalAxisTopHandle.OnMove = func(x, y float64) {
		plantDiagram.AxesShape.LengthY = plantDiagram.OriginY - y - handleSize
		stager.stage.Commit()
	}

	verticalAxisTopHandle.Presentation.Stroke = "blue"
	verticalAxisTopHandle.Presentation.StrokeWidth = 1
	verticalAxisTopHandle.Presentation.StrokeOpacity = 1

	verticalAxisBottomHandle := new(svg.Rect)
	verticalAxisBottomHandle.Name = "Vertical axis top handle"
	if !plantDiagram.AxesShape.GetIsWithHiddenHandle() {
		layer.Rects = append(layer.Rects, verticalAxisBottomHandle)
	}

	verticalAxisBottomHandle.X = plantDiagram.OriginX - handleSize/2.0
	verticalAxisBottomHandle.Y = plantDiagram.OriginY - handleSize
	verticalAxisBottomHandle.Width = handleSize
	verticalAxisBottomHandle.Height = handleSize
	verticalAxisBottomHandle.CanMoveVerticaly = true
	verticalAxisBottomHandle.CanMoveHorizontaly = true
	verticalAxisBottomHandle.OnMove = func(x, y float64) {
		plantDiagram.OriginX = x + handleSize/2.0
		plantDiagram.OriginY = y + handleSize
		stager.stage.Commit() // the top handle will move with the commit
	}

	verticalAxisBottomHandle.Presentation.Stroke = "blue"
	verticalAxisBottomHandle.Presentation.StrokeWidth = 1
	verticalAxisBottomHandle.Presentation.StrokeOpacity = 1

	verticalAxisLine := new(svg.Link)
	layer.Links = append(layer.Links, verticalAxisLine)

	verticalAxisLine.StrokeWidth = 1
	verticalAxisLine.StrokeOpacity = 1
	verticalAxisLine.Name = "Vertical Axis"
	verticalAxisLine.Stroke = svg.Black.ToString()
	verticalAxisLine.Start = verticalAxisBottomHandle
	verticalAxisLine.End = verticalAxisTopHandle

	verticalAxisLine.HasStartArrow = false
	verticalAxisLine.HasEndArrow = true

	verticalAxisLine.CornerOffsetRatio = 2.0

	verticalAxisLine.EndArrowSize = 8
	verticalAxisLine.Type = svg.LINK_TYPE_LINE_WITH_CONTROL_POINTS

	verticalAxisLine.StartAnchorType = svg.ANCHOR_CENTER
	verticalAxisLine.EndAnchorType = svg.ANCHOR_CENTER

	horizontalAxisRightHandle := new(svg.Rect)
	horizontalAxisRightHandle.Name = "Horizontal axis right handle"
	if !plantDiagram.AxesShape.GetIsWithHiddenHandle() {
		layer.Rects = append(layer.Rects, horizontalAxisRightHandle)
	}
	horizontalAxisRightHandle.X = plantDiagram.OriginX + plantDiagram.AxesShape.LengthX - handleSize/2.0
	horizontalAxisRightHandle.Y = plantDiagram.OriginY - handleSize
	horizontalAxisRightHandle.Width = handleSize
	horizontalAxisRightHandle.Height = handleSize
	horizontalAxisRightHandle.CanMoveHorizontaly = true
	horizontalAxisRightHandle.CanMoveVerticaly = true
	horizontalAxisRightHandle.OnMove = func(x, y float64) {
		plantDiagram.AxesShape.LengthX = x - plantDiagram.OriginX + handleSize/2.0
		stager.stage.Commit()
	}

	horizontalAxisRightHandle.Presentation.Stroke = "blue"
	horizontalAxisRightHandle.Presentation.StrokeWidth = 1
	horizontalAxisRightHandle.Presentation.StrokeOpacity = 1

	horizontalAxisLine := new(svg.Link)
	layer.Links = append(layer.Links, horizontalAxisLine)

	horizontalAxisLine.StrokeWidth = 1
	horizontalAxisLine.StrokeOpacity = 1
	horizontalAxisLine.Name = "Horizontal Axis"
	horizontalAxisLine.Stroke = svg.Black.ToString()
	horizontalAxisLine.Start = verticalAxisBottomHandle
	horizontalAxisLine.End = horizontalAxisRightHandle

	horizontalAxisLine.HasStartArrow = false
	horizontalAxisLine.HasEndArrow = true

	horizontalAxisLine.CornerOffsetRatio = 2.0

	horizontalAxisLine.EndArrowSize = 8
	horizontalAxisLine.Type = svg.LINK_TYPE_LINE_WITH_CONTROL_POINTS

	horizontalAxisLine.StartAnchorType = svg.ANCHOR_BOTTOM
	horizontalAxisLine.EndAnchorType = svg.ANCHOR_BOTTOM

	// right and top handle move with vertical bottom handle
	verticalAxisBottomHandle.Peers = append(verticalAxisBottomHandle.Peers,
		verticalAxisTopHandle,
		horizontalAxisRightHandle)
}

func (plantDiagram *PlantDiagram) drawGrowthVectorShape(stager *Stager, layer *svg.Layer) {

	if plantDiagram.GrowthVectorShape.IsHidden {
		return
	}

	angleRad := plantDiagram.GrowthVectorShape.AngleDegree * math.Pi / 180.0
	length := plantDiagram.GrowthVectorShape.Length

	endX := plantDiagram.OriginX + length*math.Cos(angleRad)
	endY := plantDiagram.OriginY - length*math.Sin(angleRad) // minus because SVG y-axis is inverted

	line := new(svg.Line)
	layer.Lines = append(layer.Lines, line)

	line.Name = plantDiagram.GrowthVectorShape.Name
	line.X1 = plantDiagram.OriginX
	line.Y1 = plantDiagram.OriginY
	line.X2 = endX
	line.Y2 = endY
	line.Presentation.Stroke = "green"
	line.Presentation.StrokeWidth = 2.0
	line.Presentation.StrokeOpacity = 1.0
}

func (plantDiagram *PlantDiagram) drawReferenceRhombus(stager *Stager, layer *svg.Layer, plant *Plant) {

	if plantDiagram.ReferenceRhombus == nil || plantDiagram.ReferenceRhombus.IsHidden {
		return
	}

	angleRad := plant.RhombusInsideAngle * math.Pi / 180.0
	length := plant.RhombusSideLength

	// Vertices
	v0x := plantDiagram.OriginX
	v0y := plantDiagram.OriginY

	// Top vertex (SVG y-axis is inverted)
	v1x := v0x + length*math.Cos(angleRad/2.0)
	v1y := v0y - length*math.Sin(angleRad/2.0)

	// Right vertex
	v2x := v0x + 2.0*length*math.Cos(angleRad/2.0)
	v2y := v0y

	// Bottom vertex
	v3x := v0x + length*math.Cos(angleRad/2.0)
	v3y := v0y + length*math.Sin(angleRad/2.0)

	polygon := new(svg.Polygone)
	layer.Polygones = append(layer.Polygones, polygon)

	polygon.Name = plantDiagram.ReferenceRhombus.Name
	polygon.Points = fmt.Sprintf("%f,%f %f,%f %f,%f %f,%f",
		v0x, v0y,
		v1x, v1y,
		v2x, v2y,
		v3x, v3y)

	polygon.Presentation.Stroke = "blue"
	polygon.Presentation.StrokeWidth = 2.0
	polygon.Presentation.StrokeOpacity = 1.0
	polygon.Presentation.Color = "lightblue"
	polygon.Presentation.FillOpacity = 0.5
}

func (plantDiagram *PlantDiagram) drawGridPathShape(stager *Stager, layer *svg.Layer, plant *Plant) {

	if plantDiagram.GridPathShape == nil || plantDiagram.GridPathShape.IsHidden {
		return
	}

	angleRad := plant.RhombusInsideAngle * math.Pi / 180.0
	length := plant.RhombusSideLength

	// SVG Y-axis is inverted
	v1x := length * math.Cos(angleRad/2.0)
	v1y := -length * math.Sin(angleRad/2.0)

	v2x := -length * math.Cos(angleRad/2.0)
	v2y := -length * math.Sin(angleRad/2.0)

	polyline := new(svg.Polyline)
	layer.Polylines = append(layer.Polylines, polyline)

	polyline.Name = plantDiagram.GridPathShape.Name
	polyline.Presentation.Stroke = "red"
	polyline.Presentation.StrokeWidth = 2.0
	polyline.Presentation.StrokeOpacity = 1.0
	polyline.Presentation.FillOpacity = 0.0

	var points []string
	currX := plantDiagram.OriginX
	currY := plantDiagram.OriginY
	points = append(points, fmt.Sprintf("%f,%f", currX, currY))

	addCircle := func(x, y float64, stepIndex int, path string) {
		circle := new(svg.Circle)
		layer.Circles = append(layer.Circles, circle)
		circle.Name = fmt.Sprintf("%s-%s-step-%d", plantDiagram.GridPathShape.Name, path, stepIndex)
		circle.CX = x
		circle.CY = y
		circle.Radius = 4.0
		circle.Presentation.Stroke = "red"
		circle.Presentation.StrokeWidth = 1.0
		circle.Presentation.StrokeOpacity = 1.0
		circle.Presentation.Color = "white"
		circle.Presentation.FillOpacity = 1.0
	}

	addCircle(currX, currY, 0, "start")

	for i := 1; i <= plant.N; i++ {
		currX += v1x
		currY += v1y
		points = append(points, fmt.Sprintf("%f,%f", currX, currY))
		addCircle(currX, currY, i, "N")
	}

	for i := 1; i <= plant.M; i++ {
		currX += v2x
		currY += v2y
		points = append(points, fmt.Sprintf("%f,%f", currX, currY))
		addCircle(currX, currY, i, "M")
	}

	polyline.Points = strings.Join(points, " ")
}
