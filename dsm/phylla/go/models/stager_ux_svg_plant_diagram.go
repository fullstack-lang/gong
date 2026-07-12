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
	plantDiagram.drawAxes(stager, layer, plant)
	plantDiagram.drawPlantCircumferenceShape(stager, layer, plant)
	plantDiagram.drawReferenceRhombus(stager, layer, plant)
	plantDiagram.drawGridPathShape(stager, layer, plant)
	plantDiagram.drawRhombusGridShape(stager, layer, plant)
	plantDiagram.drawExplanationTextShape(stager, layer, plant)
	plantDiagram.drawRotatedPlantCircumferenceShape(stager, layer, plant)
	plantDiagram.drawRotatedReferenceRhombus(stager, layer, plant)
	plantDiagram.drawRotatedGridPathShape(stager, layer, plant)
	plantDiagram.drawRotatedRhombusGridShape(stager, layer, plant)
	plantDiagram.drawGrowthPathRhombusGridShape(stager, layer, plant)
	plantDiagram.drawGrowthVectorShape(stager, layer, plant)
	plantDiagram.drawPerpendicularVectorGrid(stager, layer, plant)
	plantDiagram.drawPerpendicularVectorGridHalfway(stager, layer, plant)
	plantDiagram.drawStartArcShapeGrid(stager, layer, plant)
	plantDiagram.drawEndArcShapeGrid(stager, layer, plant)
	plantDiagram.drawGrowthCurveBezierShapeGrid(stager, layer, plant)
	plantDiagram.drawStackOfGrowthCurve(stager, layer, plant)

	return
}

const AxisHandleBorderLength = 25

func (plantDiagram *PlantDiagram) drawAxes(stager *Stager, layer *svg.Layer, plant *Plant) {

	if plantDiagram.IsHiddenAxesShape {
		return
	}

	handleSize := float64(AxisHandleBorderLength)
	if plant.AxesShape.GetIsWithHiddenHandle() {
		handleSize = 0.0
	}

	verticalAxisTopHandle := new(svg.Rect)
	verticalAxisTopHandle.Name = "Vertical axis bottom handle"
	if !plant.AxesShape.GetIsWithHiddenHandle() {
		layer.Rects = append(layer.Rects, verticalAxisTopHandle)
	}
	verticalAxisTopHandle.X = plantDiagram.OriginX - handleSize/2.0
	verticalAxisTopHandle.Y = plantDiagram.OriginY - plant.AxesShape.LengthY - handleSize
	verticalAxisTopHandle.Width = handleSize
	verticalAxisTopHandle.Height = handleSize
	verticalAxisTopHandle.CanMoveVerticaly = true
	verticalAxisTopHandle.CanMoveHorizontaly = true
	verticalAxisTopHandle.OnMove = func(x, y float64) {
		plant.AxesShape.LengthY = plantDiagram.OriginY - y - handleSize
		stager.stage.Commit()
	}

	verticalAxisTopHandle.Presentation.Stroke = "blue"
	verticalAxisTopHandle.Presentation.StrokeWidth = 1
	verticalAxisTopHandle.Presentation.StrokeOpacity = 1

	verticalAxisBottomHandle := new(svg.Rect)
	verticalAxisBottomHandle.Name = "Vertical axis top handle"
	if !plant.AxesShape.GetIsWithHiddenHandle() {
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
	if !plant.AxesShape.GetIsWithHiddenHandle() {
		layer.Rects = append(layer.Rects, horizontalAxisRightHandle)
	}
	horizontalAxisRightHandle.X = plantDiagram.OriginX + plant.AxesShape.LengthX - handleSize/2.0
	horizontalAxisRightHandle.Y = plantDiagram.OriginY - handleSize
	horizontalAxisRightHandle.Width = handleSize
	horizontalAxisRightHandle.Height = handleSize
	horizontalAxisRightHandle.CanMoveHorizontaly = true
	horizontalAxisRightHandle.CanMoveVerticaly = true
	horizontalAxisRightHandle.OnMove = func(x, y float64) {
		plant.AxesShape.LengthX = x - plantDiagram.OriginX + handleSize/2.0
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

func (plantDiagram *PlantDiagram) drawExplanationTextShape(stager *Stager, layer *svg.Layer, plant *Plant) {
	if plant.ExplanationTextShape == nil || plantDiagram.IsHiddenExplanationTextShape {
		return
	}

	lines := []string{
		"this diagram is the construction of the growth curve.",
		"First, we know that the plant circonference will cross all spirals",
		"constructing the spirals starts with the reference rombus grid that represents",
		fmt.Sprintf("the N (%d) spiral in one direction and the M (%d) spiral in the other direction.", plant.N, plant.M),
	}

	var endY float64 = plantDiagram.OriginY
	angleRad := plant.PlantCircumferenceShape.AngleDegree * math.Pi / 180.0
	length := plant.PlantCircumferenceShape.Length
	endY = plantDiagram.OriginY - length*math.Sin(angleRad)

	startY := endY - float64(len(lines))*20.0 - 20.0 // Above the circumference vector
	startX := plantDiagram.OriginX + 50.0

	for i, lineText := range lines {
		text := new(svg.Text)
		layer.Texts = append(layer.Texts, text)
		text.Name = fmt.Sprintf("%s-line-%d", plant.ExplanationTextShape.Name, i)
		text.Content = lineText
		text.X = startX
		text.Y = startY + float64(i)*20.0
		text.Presentation.FillOpacity = 1.0
		text.Presentation.Color = "black"
		text.TextAttributes.FontSize = "14"
	}
}

func (plantDiagram *PlantDiagram) drawPlantCircumferenceShape(stager *Stager, layer *svg.Layer, plant *Plant) {

	if plantDiagram.IsHiddenPlantCircumferenceShape {
		return
	}

	angleRad := plant.PlantCircumferenceShape.AngleDegree * math.Pi / 180.0
	length := plant.PlantCircumferenceShape.Length

	endX := plantDiagram.OriginX + length*math.Cos(angleRad)
	endY := plantDiagram.OriginY - length*math.Sin(angleRad) // minus because SVG y-axis is inverted

	line := new(svg.Line)
	layer.Lines = append(layer.Lines, line)

	line.Name = plant.PlantCircumferenceShape.Name
	line.X1 = plantDiagram.OriginX
	line.Y1 = plantDiagram.OriginY
	line.X2 = endX
	line.Y2 = endY
	line.Presentation.Stroke = "green"
	line.Presentation.StrokeWidth = 2.0
	line.Presentation.StrokeOpacity = 1.0
}

func (plantDiagram *PlantDiagram) drawReferenceRhombus(stager *Stager, layer *svg.Layer, plant *Plant) {

	if plant.ReferenceRhombus == nil || plantDiagram.IsHiddenReferenceRhombus {
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

	polygon.Name = plant.ReferenceRhombus.Name
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

	if plant.GridPathShape == nil || plantDiagram.IsHiddenGridPathShape {
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

	polyline.Name = plant.GridPathShape.Name
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
		circle.Name = fmt.Sprintf("%s-%s-step-%d", plant.GridPathShape.Name, path, stepIndex)
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

func (plantDiagram *PlantDiagram) drawRotatedPlantCircumferenceShape(stager *Stager, layer *svg.Layer, plant *Plant) {

	if plant.RotatedPlantCircumferenceShape == nil || plantDiagram.IsHiddenRotatedPlantCircumferenceShape {
		return
	}

	angleRad := plant.RotatedPlantCircumferenceShape.AngleDegree * math.Pi / 180.0
	length := plant.RotatedPlantCircumferenceShape.Length

	// SVG Y-axis is inverted
	endX := plantDiagram.OriginX + length*math.Cos(angleRad)
	endY := plantDiagram.OriginY - length*math.Sin(angleRad)

	line := new(svg.Line)
	layer.Lines = append(layer.Lines, line)

	line.Name = plant.RotatedPlantCircumferenceShape.Name
	line.X1 = plantDiagram.OriginX
	line.Y1 = plantDiagram.OriginY
	line.X2 = endX
	line.Y2 = endY

	line.Presentation.Stroke = "darkgreen" // slightly different color to distinguish
	line.Presentation.StrokeWidth = 2.0
	line.Presentation.StrokeOpacity = 1.0
	line.Presentation.StrokeDashArray = "5, 5" // make it dashed
}

func (plantDiagram *PlantDiagram) drawRotatedReferenceRhombus(stager *Stager, layer *svg.Layer, plant *Plant) {

	if plant.RotatedReferenceRhombus == nil || plantDiagram.IsHiddenRotatedReferenceRhombus {
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

	polygon.Name = plant.RotatedReferenceRhombus.Name
	polygon.Points = fmt.Sprintf("%f,%f %f,%f %f,%f %f,%f",
		v0x, v0y,
		v1x, v1y,
		v2x, v2y,
		v3x, v3y)

	polygon.Presentation.Stroke = "darkblue"
	polygon.Presentation.StrokeWidth = 2.0
	polygon.Presentation.StrokeOpacity = 1.0
	polygon.Presentation.Color = "lightblue"
	polygon.Presentation.FillOpacity = 0.2 // more transparent
	polygon.Presentation.StrokeDashArray = "5, 5"

	// Add rotation transform
	polygon.Presentation.Transform = fmt.Sprintf("rotate(%f %f %f)", plant.PlantCircumferenceShape.AngleDegree, plantDiagram.OriginX, plantDiagram.OriginY)
}

func (plantDiagram *PlantDiagram) drawRotatedGridPathShape(stager *Stager, layer *svg.Layer, plant *Plant) {

	if plant.RotatedGridPathShape == nil || plantDiagram.IsHiddenRotatedGridPathShape {
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

	polyline.Name = plant.RotatedGridPathShape.Name
	polyline.Presentation.Stroke = "darkred"
	polyline.Presentation.StrokeWidth = 2.0
	polyline.Presentation.StrokeOpacity = 1.0
	polyline.Presentation.FillOpacity = 0.0
	polyline.Presentation.StrokeDashArray = "5, 5"
	polyline.Presentation.Transform = fmt.Sprintf("rotate(%f %f %f)", plant.PlantCircumferenceShape.AngleDegree, plantDiagram.OriginX, plantDiagram.OriginY)

	var points []string
	currX := plantDiagram.OriginX
	currY := plantDiagram.OriginY
	points = append(points, fmt.Sprintf("%f,%f", currX, currY))

	addCircle := func(x, y float64, stepIndex int, path string) {
		circle := new(svg.Circle)
		layer.Circles = append(layer.Circles, circle)
		circle.Name = fmt.Sprintf("%s-%s-step-%d", plant.RotatedGridPathShape.Name, path, stepIndex)
		circle.CX = x
		circle.CY = y
		circle.Radius = 4.0
		circle.Presentation.Stroke = "darkred"
		circle.Presentation.StrokeWidth = 1.0
		circle.Presentation.StrokeOpacity = 1.0
		circle.Presentation.Color = "white"
		circle.Presentation.FillOpacity = 1.0
		circle.Presentation.Transform = fmt.Sprintf("rotate(%f %f %f)", plant.PlantCircumferenceShape.AngleDegree, plantDiagram.OriginX, plantDiagram.OriginY)
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

func (plantDiagram *PlantDiagram) drawRhombusGridShape(stager *Stager, layer *svg.Layer, plant *Plant) {

	if plant.InitialRhombusGridShape == nil || plantDiagram.IsHiddenRhombusGridShape {
		return
	}

	angleRad := plant.RhombusInsideAngle * math.Pi / 180.0
	length := plant.RhombusSideLength

	// SVG Y-axis is inverted
	v1x := length * math.Cos(angleRad/2.0)
	v1y := -length * math.Sin(angleRad/2.0)

	v2x := -length * math.Cos(angleRad/2.0)
	v2y := -length * math.Sin(angleRad/2.0)

	for _, rhombus := range plant.InitialRhombusGridShape.InitialRhombusShapes {
		polygon := new(svg.Polygone)
		layer.Polygones = append(layer.Polygones, polygon)

		polygon.Name = rhombus.Name

		// r.X and r.Y are Cartesian center coordinates
		svg_cx := plantDiagram.OriginX + rhombus.X
		svg_cy := plantDiagram.OriginY - rhombus.Y

		// Calculate v0 (bottom-left vertex in visual SVG space) from the center
		v0x := svg_cx - (v1x + v2x)/2.0
		v0y := svg_cy - (v1y + v2y)/2.0

		v1_vertex_x := v0x + v1x
		v1_vertex_y := v0y + v1y

		v2_vertex_x := v0x + v1x + v2x
		v2_vertex_y := v0y + v1y + v2y

		v3_vertex_x := v0x + v2x
		v3_vertex_y := v0y + v2y

		polygon.Points = fmt.Sprintf("%f,%f %f,%f %f,%f %f,%f",
			v0x, v0y,
			v1_vertex_x, v1_vertex_y,
			v2_vertex_x, v2_vertex_y,
			v3_vertex_x, v3_vertex_y)

		polygon.Presentation.Stroke = "blue"
		polygon.Presentation.StrokeWidth = 1.0
		polygon.Presentation.StrokeOpacity = 0.5
		polygon.Presentation.Color = "lightblue"
		polygon.Presentation.FillOpacity = 0.2

		// Draw a little cross at the center (which is simply svg_cx, svg_cy)
		cx := svg_cx
		cy := svg_cy

		line1 := new(svg.Line)
		layer.Lines = append(layer.Lines, line1)
		line1.Name = rhombus.Name + "-cross-h"
		line1.X1 = cx - 3.0
		line1.Y1 = cy
		line1.X2 = cx + 3.0
		line1.Y2 = cy
		line1.Presentation.Stroke = "black"
		line1.Presentation.StrokeWidth = 1.0
		line1.Presentation.StrokeOpacity = 1.0

		line2 := new(svg.Line)
		layer.Lines = append(layer.Lines, line2)
		line2.Name = rhombus.Name + "-cross-v"
		line2.X1 = cx
		line2.Y1 = cy - 3.0
		line2.X2 = cx
		line2.Y2 = cy + 3.0
		line2.Presentation.Stroke = "black"
		line2.Presentation.StrokeWidth = 1.0
		line2.Presentation.StrokeOpacity = 1.0
	}
}

func (plantDiagram *PlantDiagram) drawRotatedRhombusGridShape(stager *Stager, layer *svg.Layer, plant *Plant) {

	if plant.RotatedRhombusGridShape2 == nil || plantDiagram.IsHiddenRotatedRhombusGridShape {
		return
	}

	angleRad := plant.RhombusInsideAngle * math.Pi / 180.0
	length := plant.RhombusSideLength

	// Cartesian vectors
	v1x := length * math.Cos(angleRad/2.0)
	v1y := length * math.Sin(angleRad/2.0)

	v2x := -length * math.Cos(angleRad/2.0)
	v2y := length * math.Sin(angleRad/2.0)

	rotRad := -plant.PlantCircumferenceShape.AngleDegree * math.Pi / 180.0
	cosA := math.Cos(rotRad)
	sinA := math.Sin(rotRad)

	// Rotate in Cartesian space
	v1_cart_rot_x := v1x*cosA - v1y*sinA
	v1_cart_rot_y := v1x*sinA + v1y*cosA
	v2_cart_rot_x := v2x*cosA - v2y*sinA
	v2_cart_rot_y := v2x*sinA + v2y*cosA

	// Map to SVG space (invert Y)
	v1_rot_x := v1_cart_rot_x
	v1_rot_y := -v1_cart_rot_y
	v2_rot_x := v2_cart_rot_x
	v2_rot_y := -v2_cart_rot_y

	for _, rhombus := range plant.RotatedRhombusGridShape2.RotatedRhombusShapes {
		polygon := new(svg.Polygone)
		layer.Polygones = append(layer.Polygones, polygon)

		polygon.Name = rhombus.Name

		// r.X and r.Y are Cartesian center coordinates
		svg_cx := plantDiagram.OriginX + rhombus.X
		svg_cy := plantDiagram.OriginY - rhombus.Y

		// Calculate v0 (bottom-left vertex in visual SVG space) from the center
		v0x := svg_cx - (v1_rot_x + v2_rot_x)/2.0
		v0y := svg_cy - (v1_rot_y + v2_rot_y)/2.0

		v1_vertex_x := v0x + v1_rot_x
		v1_vertex_y := v0y + v1_rot_y

		v2_vertex_x := v0x + v1_rot_x + v2_rot_x
		v2_vertex_y := v0y + v1_rot_y + v2_rot_y

		v3_vertex_x := v0x + v2_rot_x
		v3_vertex_y := v0y + v2_rot_y

		polygon.Points = fmt.Sprintf("%f,%f %f,%f %f,%f %f,%f",
			v0x, v0y,
			v1_vertex_x, v1_vertex_y,
			v2_vertex_x, v2_vertex_y,
			v3_vertex_x, v3_vertex_y)

		polygon.Presentation.Stroke = "darkblue"
		polygon.Presentation.StrokeWidth = 1.0
		polygon.Presentation.StrokeOpacity = 0.5
		polygon.Presentation.Color = "lightblue"
		polygon.Presentation.FillOpacity = 0.1
		polygon.Presentation.StrokeDashArray = "5, 5"
		
		// Draw a little cross at the center (which is simply svg_cx, svg_cy)
		cx := svg_cx
		cy := svg_cy

		
		line1 := new(svg.Line)
		layer.Lines = append(layer.Lines, line1)
		line1.Name = rhombus.Name + "-cross-h"
		line1.X1 = cx - 3.0
		line1.Y1 = cy
		line1.X2 = cx + 3.0
		line1.Y2 = cy
		line1.Presentation.Stroke = "black"
		line1.Presentation.StrokeWidth = 1.0
		line1.Presentation.StrokeOpacity = 1.0
		
		line2 := new(svg.Line)
		layer.Lines = append(layer.Lines, line2)
		line2.Name = rhombus.Name + "-cross-v"
		line2.X1 = cx
		line2.Y1 = cy - 3.0
		line2.X2 = cx
		line2.Y2 = cy + 3.0
		line2.Presentation.Stroke = "black"
		line2.Presentation.StrokeWidth = 1.0
		line2.Presentation.StrokeOpacity = 1.0
			}
}

func (plantDiagram *PlantDiagram) drawGrowthPathRhombusGridShape(stager *Stager, layer *svg.Layer, plant *Plant) {

	if plant.GrowthCurveRhombusGridShape == nil || plantDiagram.IsHiddenGrowthPathRhombusGridShape {
		return
	}

	angleRad := plant.RhombusInsideAngle * math.Pi / 180.0
	length := plant.RhombusSideLength

	// Cartesian vectors
	v1x := length * math.Cos(angleRad/2.0)
	v1y := length * math.Sin(angleRad/2.0)

	v2x := -length * math.Cos(angleRad/2.0)
	v2y := length * math.Sin(angleRad/2.0)

	rotRad := -plant.PlantCircumferenceShape.AngleDegree * math.Pi / 180.0
	cosA := math.Cos(rotRad)
	sinA := math.Sin(rotRad)

	// Rotate in Cartesian space
	v1_cart_rot_x := v1x*cosA - v1y*sinA
	v1_cart_rot_y := v1x*sinA + v1y*cosA
	v2_cart_rot_x := v2x*cosA - v2y*sinA
	v2_cart_rot_y := v2x*sinA + v2y*cosA

	// Map to SVG space (invert Y)
	v1_rot_x := v1_cart_rot_x
	v1_rot_y := -v1_cart_rot_y
	v2_rot_x := v2_cart_rot_x
	v2_rot_y := -v2_cart_rot_y

	for _, rhombus := range plant.GrowthCurveRhombusGridShape.GrowthCurveRhombusShapes {
		polygon := new(svg.Polygone)
		layer.Polygones = append(layer.Polygones, polygon)

		polygon.Name = rhombus.Name

		// r.X and r.Y are Cartesian center coordinates
		svg_cx := plantDiagram.OriginX + rhombus.X
		svg_cy := plantDiagram.OriginY - rhombus.Y

		// Calculate v0 (bottom-left vertex in visual SVG space) from the center
		v0x := svg_cx - (v1_rot_x + v2_rot_x)/2.0
		v0y := svg_cy - (v1_rot_y + v2_rot_y)/2.0

		v1_vertex_x := v0x + v1_rot_x
		v1_vertex_y := v0y + v1_rot_y

		v2_vertex_x := v0x + v1_rot_x + v2_rot_x
		v2_vertex_y := v0y + v1_rot_y + v2_rot_y

		v3_vertex_x := v0x + v2_rot_x
		v3_vertex_y := v0y + v2_rot_y

		polygon.Points = fmt.Sprintf("%f,%f %f,%f %f,%f %f,%f",
			v0x, v0y,
			v1_vertex_x, v1_vertex_y,
			v2_vertex_x, v2_vertex_y,
			v3_vertex_x, v3_vertex_y)

		polygon.Presentation.Stroke = "red"
		polygon.Presentation.StrokeWidth = 4.0
		polygon.Presentation.StrokeOpacity = 1.0
		polygon.Presentation.Color = "lightblue"
		polygon.Presentation.FillOpacity = 0.0
		
		// Draw a little cross at the center (which is simply svg_cx, svg_cy)
		cx := svg_cx
		cy := svg_cy

		
		line1 := new(svg.Line)
		layer.Lines = append(layer.Lines, line1)
		line1.Name = rhombus.Name + "-cross-h"
		line1.X1 = cx - 3.0
		line1.Y1 = cy
		line1.X2 = cx + 3.0
		line1.Y2 = cy
		line1.Presentation.Stroke = "red"
		line1.Presentation.StrokeWidth = 2.0
		line1.Presentation.StrokeOpacity = 1.0
		
		line2 := new(svg.Line)
		layer.Lines = append(layer.Lines, line2)
		line2.Name = rhombus.Name + "-cross-v"
		line2.X1 = cx
		line2.Y1 = cy - 3.0
		line2.X2 = cx
		line2.Y2 = cy + 3.0
		line2.Presentation.Stroke = "red"
		line2.Presentation.StrokeWidth = 2.0
		line2.Presentation.StrokeOpacity = 1.0
			}
}

func (plantDiagram *PlantDiagram) drawGrowthVectorShape(stager *Stager, layer *svg.Layer, plant *Plant) {
	if plant.GrowthVectorShape == nil || plantDiagram.IsHiddenGrowthVectorShape {
		return
	}
	if plant.GrowthCurveRhombusGridShape == nil || len(plant.GrowthCurveRhombusGridShape.GrowthCurveRhombusShapes) < 2 {
		return
	}

	rhombuses := plant.GrowthCurveRhombusGridShape.GrowthCurveRhombusShapes
	first := rhombuses[0]

	line := new(svg.Line)
	layer.Lines = append(layer.Lines, line)
	line.Name = plant.GrowthVectorShape.Name

	svg_x1 := plantDiagram.OriginX + first.X
	svg_y1 := plantDiagram.OriginY - first.Y

	svg_x2 := svg_x1 + plant.GrowthVectorShape.X
	svg_y2 := svg_y1 - plant.GrowthVectorShape.Y

	line.X1 = svg_x1
	line.Y1 = svg_y1
	line.X2 = svg_x2
	line.Y2 = svg_y2

	line.Presentation.Stroke = "blue"
	line.Presentation.StrokeWidth = 4.0
	line.Presentation.StrokeOpacity = 1.0
}

func (plantDiagram *PlantDiagram) drawPerpendicularVectorGrid(stager *Stager, layer *svg.Layer, plant *Plant) {
	if plant.PerpendicularVectorGrid == nil || plantDiagram.IsHiddenPerpendicularVectorGrid {
		return
	}

	for _, vec := range plant.PerpendicularVectorGrid.PerpendicularVectors {
		line := new(svg.Line)
		layer.Lines = append(layer.Lines, line)
		line.Name = vec.Name

		svg_x1 := plantDiagram.OriginX + vec.StartX
		svg_y1 := plantDiagram.OriginY - vec.StartY

		svg_x2 := plantDiagram.OriginX + vec.EndX
		svg_y2 := plantDiagram.OriginY - vec.EndY

		line.X1 = svg_x1
		line.Y1 = svg_y1
		line.X2 = svg_x2
		line.Y2 = svg_y2

		line.Presentation.Stroke = "green"
		line.Presentation.StrokeWidth = 2.0
		line.Presentation.StrokeOpacity = 1.0
	}
}

func (plantDiagram *PlantDiagram) drawPerpendicularVectorGridHalfway(stager *Stager, layer *svg.Layer, plant *Plant) {
	if plant.PerpendicularVectorGridHalfway == nil || plantDiagram.IsHiddenPerpendicularVectorGridHalfway {
		return
	}

	for _, vec := range plant.PerpendicularVectorGridHalfway.PerpendicularVectorHalfways {
		line := new(svg.Line)
		layer.Lines = append(layer.Lines, line)
		line.Name = vec.Name

		svg_x1 := plantDiagram.OriginX + vec.StartX
		svg_y1 := plantDiagram.OriginY - vec.StartY

		svg_x2 := plantDiagram.OriginX + vec.EndX
		svg_y2 := plantDiagram.OriginY - vec.EndY

		line.X1 = svg_x1
		line.Y1 = svg_y1
		line.X2 = svg_x2
		line.Y2 = svg_y2

		line.Presentation.Stroke = "purple" // Distinct color for halfway
		line.Presentation.StrokeWidth = 2.0
		line.Presentation.StrokeOpacity = 1.0
	}
}

func (plantDiagram *PlantDiagram) drawStartArcShapeGrid(stager *Stager, layer *svg.Layer, plant *Plant) {
	if plant.StartArcShapeGrid == nil || plantDiagram.IsHiddenStartArcShapeGrid {
		return
	}

	for _, arc := range plant.StartArcShapeGrid.StartArcShapes {
		path := new(svg.Path)
		layer.Paths = append(layer.Paths, path)
		path.Name = arc.Name

		sweepFlag := 0
		if arc.SweepFlag {
			sweepFlag = 1
		}
		largeArcFlag := 0
		if arc.LargeArcFlag {
			largeArcFlag = 1
		}

		path.Definition = fmt.Sprintf("M %0.1f %0.1f A %0.1f %0.1f %0.1f %d %d %0.1f %0.1f",
			plantDiagram.OriginX+arc.StartX, plantDiagram.OriginY-arc.StartY,
			arc.RadiusX, arc.RadiusY, arc.XAxisRotation, largeArcFlag, sweepFlag,
			plantDiagram.OriginX+arc.EndX, plantDiagram.OriginY-arc.EndY,
		)

		path.Presentation.Stroke = "cyan"
		path.Presentation.StrokeWidth = 3.0
		path.Presentation.StrokeOpacity = 1.0
		path.Presentation.FillOpacity = 0.0
	}
}

func (plantDiagram *PlantDiagram) drawEndArcShapeGrid(stager *Stager, layer *svg.Layer, plant *Plant) {
	if plant.EndArcShapeGrid == nil || plantDiagram.IsHiddenEndArcShapeGrid {
		return
	}

	for _, arc := range plant.EndArcShapeGrid.EndArcShapes {
		path := new(svg.Path)
		layer.Paths = append(layer.Paths, path)
		path.Name = arc.Name

		sweepFlag := 0
		if arc.SweepFlag {
			sweepFlag = 1
		}
		largeArcFlag := 0
		if arc.LargeArcFlag {
			largeArcFlag = 1
		}

		path.Definition = fmt.Sprintf("M %0.1f %0.1f A %0.1f %0.1f %0.1f %d %d %0.1f %0.1f",
			plantDiagram.OriginX+arc.StartX, plantDiagram.OriginY-arc.StartY,
			arc.RadiusX, arc.RadiusY, arc.XAxisRotation, largeArcFlag, sweepFlag,
			plantDiagram.OriginX+arc.EndX, plantDiagram.OriginY-arc.EndY,
		)

		path.Presentation.Stroke = "magenta"
		path.Presentation.StrokeWidth = 3.0
		path.Presentation.StrokeOpacity = 1.0
		path.Presentation.FillOpacity = 0.0
	}
}


func (plantDiagram *PlantDiagram) drawGrowthCurveBezierShapeGrid(stager *Stager, layer *svg.Layer, plant *Plant) {
	if plant.GrowthCurveBezierShapeGrid == nil || plantDiagram.IsHiddenGrowthCurveBezierShapeGrid {
		return
	}

	for _, b := range plant.GrowthCurveBezierShapeGrid.GrowthCurveBezierShapes {
		path := new(svg.Path)
		layer.Paths = append(layer.Paths, path)
		path.Name = b.Name

		path.Definition = fmt.Sprintf("M %0.1f %0.1f C %0.1f %0.1f, %0.1f %0.1f, %0.1f %0.1f",
			plantDiagram.OriginX+b.StartX, plantDiagram.OriginY-b.StartY,
			plantDiagram.OriginX+b.ControlPointStartX, plantDiagram.OriginY-b.ControlPointStartY,
			plantDiagram.OriginX+b.ControlPointEndX, plantDiagram.OriginY-b.ControlPointEndY,
			plantDiagram.OriginX+b.EndX, plantDiagram.OriginY-b.EndY,
		)

		path.Presentation.Stroke = "blue"
		path.Presentation.StrokeWidth = 3.0
		path.Presentation.FillOpacity = 0.0
		path.Presentation.StrokeOpacity = 0.8
	}
}

func (plantDiagram *PlantDiagram) drawStackOfGrowthCurve(stager *Stager, layer *svg.Layer, plant *Plant) {
	if plant.StackOfGrowthCurve == nil || plantDiagram.IsHiddenStackOfGrowthCurve {
		return
	}

	for _, b := range plant.StackOfGrowthCurve.StackGrowthCurveBezierShapes {
		path := new(svg.Path)
		layer.Paths = append(layer.Paths, path)
		path.Name = b.Name

		path.Definition = fmt.Sprintf("M %0.1f %0.1f C %0.1f %0.1f, %0.1f %0.1f, %0.1f %0.1f",
			plantDiagram.OriginX+b.StartX, plantDiagram.OriginY-b.StartY,
			plantDiagram.OriginX+b.ControlPointStartX, plantDiagram.OriginY-b.ControlPointStartY,
			plantDiagram.OriginX+b.ControlPointEndX, plantDiagram.OriginY-b.ControlPointEndY,
			plantDiagram.OriginX+b.EndX, plantDiagram.OriginY-b.EndY,
		)

		path.Presentation.Stroke = "darkcyan"
		path.Presentation.StrokeWidth = 2.0
		path.Presentation.FillOpacity = 0.0
		path.Presentation.StrokeOpacity = 0.6
	}
}
