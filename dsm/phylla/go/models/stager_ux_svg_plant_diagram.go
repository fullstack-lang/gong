package models

import (
	"fmt"
	"math"
	"strings"

	svg "github.com/fullstack-lang/gong/lib/svg/go/models"
)

func (stager *Stager) ux_svg_plant_diagram() {

	stager.svgPlantStage.Reset()
	stager.svgVaseStage.Reset()

	var plant2DDiagram *Plant2DDiagram
	var vase2DDiagram *Vase2DDiagram
	for plant := range *GetGongstructInstancesSetFromPointerType[*PlantAbstract](stager.stage) {
		if plant.IsSelected {
			for _, d_ := range plant.Plant2DDiagrams {
				if d_.IsChecked {
					plant2DDiagram = d_
				}
			}
			for _, d_ := range plant.Vase2DDiagrams {
				if d_.IsChecked {
					vase2DDiagram = d_
				}
			}
			if plant2DDiagram == nil && len(plant.Plant2DDiagrams) > 0 {
				plant2DDiagram = plant.Plant2DDiagrams[0]
			}
			if vase2DDiagram == nil && len(plant.Vase2DDiagrams) > 0 {
				vase2DDiagram = plant.Vase2DDiagrams[0]
			}
		}
	}

	var plant *PlantAbstract
	for p_ := range *GetGongstructInstancesSetFromPointerType[*PlantAbstract](stager.stage) {
		if p_.IsSelected {
			plant = p_
		}
	}

	if plant2DDiagram == nil || plant == nil {
		stager.svgPlantStage.Commit()
		stager.svgVaseStage.Commit()
		return
	}

	// 1. Generate and stage Plant 2D SVG
	svgPlantObject := stager.generateSvgPlantObject(plant2DDiagram, plant)
	svg.StageBranch(stager.svgPlantStage, svgPlantObject)
	stager.svgPlantStage.Commit()

	// 2. Generate and stage Vase 2D SVG
	if vase2DDiagram != nil {
		svgVaseObject := stager.generateSvgVaseObject(plant2DDiagram, vase2DDiagram, plant)
		svg.StageBranch(stager.svgVaseStage, svgVaseObject)
		stager.svgObject = svgVaseObject
		stager.svgVaseStage.Commit()
	} else {
		stager.svgVaseStage.Commit()
	}
}

func (plant2DDiagram *Plant2DDiagram) drawCommonPlant(stager *Stager, layer *svg.Layer, plant *PlantAbstract) {
	// creation of 2 transparant rects, one at each ends of the vertical
	plant2DDiagram.drawAxes(stager, layer, plant, plant2DDiagram.IsHiddenAxesShape)
	plant2DDiagram.drawPlantCircumferenceShape(stager, layer, plant)
	plant2DDiagram.drawReferenceRhombus(stager, layer, plant)
	plant2DDiagram.drawGridPathShape(stager, layer, plant)
	plant2DDiagram.drawRhombusGridShape(stager, layer, plant)
	plant2DDiagram.drawExplanationTextShape(stager, layer, plant)
	plant2DDiagram.drawRotatedPlantCircumferenceShape(stager, layer, plant)
	plant2DDiagram.drawRotatedReferenceRhombus(stager, layer, plant)
	plant2DDiagram.drawRotatedGridPathShape(stager, layer, plant)
	plant2DDiagram.drawRotatedRhombusGridShape(stager, layer, plant)
	plant2DDiagram.drawGrowthPathRhombusGridShape(stager, layer, plant)
	plant2DDiagram.drawGrowthVectorShape(stager, layer, plant)
	plant2DDiagram.drawPerpendicularVectorGrid(stager, layer, plant)
	plant2DDiagram.drawBaseVectorShapeGrid(stager, layer, plant)
	plant2DDiagram.drawArcNormalVectorShapeGrid(stager, layer, plant)
	plant2DDiagram.drawStartArcShapeV2Grid(stager, layer, plant)
	plant2DDiagram.drawMidArcVectorShapeGrid(stager, layer, plant)
	plant2DDiagram.drawEndArcShapeV2Grid(stager, layer, plant)
	plant2DDiagram.drawGrowthCurve2D(stager, layer, plant)
	plant2DDiagram.drawStackOfGrowthCurve2DByGrowthVector(stager, layer, plant)
}

func (vase2DDiagram *Vase2DDiagram) drawVaseDiagram(stager *Stager, layer *svg.Layer, plant *PlantAbstract, plant2DDiagram *Plant2DDiagram) {
	if vase2DDiagram == nil {
		return
	}
	plant2DDiagram.drawPerpendicularVectorGridHalfway(stager, layer, plant, vase2DDiagram)
	plant2DDiagram.drawTopStartArcShapeV2Grid(stager, layer, plant, vase2DDiagram)
	plant2DDiagram.drawShiftedBottomTopStartArcShapeV2Grid(stager, layer, plant, vase2DDiagram)
	plant2DDiagram.drawTopMidArcVectorShapeGrid(stager, layer, plant, vase2DDiagram)
	plant2DDiagram.drawStartHalfwayArcShapeGrid(stager, layer, plant, vase2DDiagram)
	plant2DDiagram.drawTopStartHalfwayArcShapeGrid(stager, layer, plant, vase2DDiagram)
	plant2DDiagram.drawEndHalfwayArcShapeGrid(stager, layer, plant, vase2DDiagram)
	plant2DDiagram.drawTopEndHalfwayArcShapeGrid(stager, layer, plant, vase2DDiagram)
	plant2DDiagram.drawTopEndArcShapeV2Grid(stager, layer, plant, vase2DDiagram)
	plant2DDiagram.drawStackOfGrowthCurveV2(stager, layer, plant, vase2DDiagram)
	plant2DDiagram.drawShiftedLeftStackOfGrowthCurveV2(stager, layer, plant, vase2DDiagram)
	plant2DDiagram.drawTopStackOfGrowthCurveV2(stager, layer, plant, vase2DDiagram)
	plant2DDiagram.drawTopGrowthCurve2D(stager, layer, plant, vase2DDiagram)
	plant2DDiagram.drawStackOfGrowthCurve2D(stager, layer, plant, vase2DDiagram)
	plant2DDiagram.drawTopStackOfGrowthCurve2D(stager, layer, plant, vase2DDiagram)
	plant2DDiagram.drawStackOfGrowthCurve2DRibbon(stager, layer, plant, vase2DDiagram)
	plant2DDiagram.drawStackOfRotatedGrowthCurve2DRibbon(stager, layer, plant, vase2DDiagram)
	plant2DDiagram.drawPartiallyGrowthCurve2DRibbon(stager, layer, plant, vase2DDiagram)
	plant2DDiagram.drawShiftedLeftPartiallyGrowthCurve2DRibbon(stager, layer, plant, vase2DDiagram)
	plant2DDiagram.drawPartiallyGrowthCurve2DTrajectory(stager, layer, plant, vase2DDiagram)
	plant2DDiagram.drawPartiallyGrowthCurve2DTrajectoryP1P2(stager, layer, plant, vase2DDiagram)
	plant2DDiagram.drawPxShape(stager, layer, plant, vase2DDiagram)
	plant2DDiagram.drawChosenP1P2PairShape(stager, layer, plant, vase2DDiagram)
	plant2DDiagram.drawKeyHoleShape(stager, layer, plant, vase2DDiagram)
	plant2DDiagram.drawGrowthCurve2DRibbon(stager, layer, plant, vase2DDiagram)
	plant2DDiagram.drawShiftedRightGrowthCurve2DRibbon(stager, layer, plant, vase2DDiagram)
	plant2DDiagram.drawShiftedLeftGrowthCurve2DRibbon(stager, layer, plant, vase2DDiagram)
}

func (stager *Stager) generateSvgPlantObject(plant2DDiagram *Plant2DDiagram, plant *PlantAbstract) (svg_ *svg.SVG) {
	svg_ = new(svg.SVG)
	svg_.Name = "Plant Diagram"
	svg_.IsEditable = true

	layer := &svg.Layer{Name: `Axis Shape Layer`}
	svg_.Layers = append(svg_.Layers, layer)

	plant2DDiagram.drawCommonPlant(stager, layer, plant)

	return
}

func (stager *Stager) generateSvgVaseObject(plant2DDiagram *Plant2DDiagram, vase2DDiagram *Vase2DDiagram, plant *PlantAbstract) (svg_ *svg.SVG) {
	svg_ = new(svg.SVG)
	svg_.Name = "Vase Diagram"
	svg_.IsEditable = true

	layer := &svg.Layer{Name: `Axis Shape Layer`}
	svg_.Layers = append(svg_.Layers, layer)

	plant2DDiagram.drawAxes(stager, layer, plant, vase2DDiagram.IsHiddenAxesShape)
	vase2DDiagram.drawVaseDiagram(stager, layer, plant, plant2DDiagram)

	return
}

const AxisHandleBorderLength = 25

func (plant2DDiagram *Plant2DDiagram) drawAxes(stager *Stager, layer *svg.Layer, plant *PlantAbstract, isHidden bool) {

	if isHidden {
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
	verticalAxisTopHandle.X = plant2DDiagram.OriginX - handleSize/2.0
	verticalAxisTopHandle.Y = plant2DDiagram.OriginY - plant.AxesShape.LengthY - handleSize
	verticalAxisTopHandle.Width = handleSize
	verticalAxisTopHandle.Height = handleSize
	verticalAxisTopHandle.CanMoveVerticaly = true
	verticalAxisTopHandle.CanMoveHorizontaly = true
	verticalAxisTopHandle.OnMove = func(x, y float64) {
		plant.AxesShape.LengthY = plant2DDiagram.OriginY - y - handleSize
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

	verticalAxisBottomHandle.X = plant2DDiagram.OriginX - handleSize/2.0
	verticalAxisBottomHandle.Y = plant2DDiagram.OriginY - handleSize
	verticalAxisBottomHandle.Width = handleSize
	verticalAxisBottomHandle.Height = handleSize
	verticalAxisBottomHandle.CanMoveVerticaly = true
	verticalAxisBottomHandle.CanMoveHorizontaly = true
	verticalAxisBottomHandle.OnMove = func(x, y float64) {
		plant2DDiagram.OriginX = x + handleSize/2.0
		plant2DDiagram.OriginY = y + handleSize
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
	horizontalAxisRightHandle.X = plant2DDiagram.OriginX + plant.AxesShape.LengthX - handleSize/2.0
	horizontalAxisRightHandle.Y = plant2DDiagram.OriginY - handleSize
	horizontalAxisRightHandle.Width = handleSize
	horizontalAxisRightHandle.Height = handleSize
	horizontalAxisRightHandle.CanMoveHorizontaly = true
	horizontalAxisRightHandle.CanMoveVerticaly = true
	horizontalAxisRightHandle.OnMove = func(x, y float64) {
		plant.AxesShape.LengthX = x - plant2DDiagram.OriginX + handleSize/2.0
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

func (plant2DDiagram *Plant2DDiagram) drawExplanationTextShape(stager *Stager, layer *svg.Layer, plant *PlantAbstract) {
	if plant2DDiagram.IsHiddenExplanationTextShape {
		return
	}

	lines := []string{
		"this diagram is the construction of the growth curve.",
		"First, we know that the plant circonference will cross all spirals",
		"constructing the spirals starts with the reference rombus grid that represents",
		fmt.Sprintf("the N (%d) spiral in one direction and the M (%d) spiral in the other direction.", plant.N, plant.M),
	}

	var endY float64 = plant2DDiagram.OriginY
	angleRad := plant.RhombusStuff.PlantCircumferenceShape.AngleDegree * math.Pi / 180.0
	length := plant.RhombusStuff.PlantCircumferenceShape.Length
	endY = plant2DDiagram.OriginY - length*math.Sin(angleRad)

	startY := endY - float64(len(lines))*20.0 - 20.0 // Above the circumference vector
	startX := plant2DDiagram.OriginX + 50.0

	for i, lineText := range lines {
		text := new(svg.Text)
		layer.Texts = append(layer.Texts, text)
		text.Name = fmt.Sprintf("%s-line-%d", plant.RhombusStuff.ExplanationTextShape.Name, i)
		text.Content = lineText
		text.X = startX
		text.Y = startY + float64(i)*20.0
		text.Presentation.FillOpacity = 1.0
		text.Presentation.Color = "black"
		text.TextAttributes.FontSize = "14"
	}
}

func (plant2DDiagram *Plant2DDiagram) drawPlantCircumferenceShape(stager *Stager, layer *svg.Layer, plant *PlantAbstract) {

	if plant2DDiagram.IsHiddenPlantCircumferenceShape {
		return
	}

	angleRad := plant.RhombusStuff.PlantCircumferenceShape.AngleDegree * math.Pi / 180.0
	length := plant.RhombusStuff.PlantCircumferenceShape.Length

	endX := plant2DDiagram.OriginX + length*math.Cos(angleRad)
	endY := plant2DDiagram.OriginY - length*math.Sin(angleRad) // minus because SVG y-axis is inverted

	line := new(svg.Line)
	layer.Lines = append(layer.Lines, line)

	line.Name = plant.RhombusStuff.PlantCircumferenceShape.Name
	line.X1 = plant2DDiagram.OriginX
	line.Y1 = plant2DDiagram.OriginY
	line.X2 = endX
	line.Y2 = endY
	line.Presentation.Stroke = "green"
	line.Presentation.StrokeWidth = 2.0
	line.Presentation.StrokeOpacity = 1.0
}

func (plant2DDiagram *Plant2DDiagram) getZoom() float64 {
	if plant2DDiagram == nil || plant2DDiagram.Zoom <= 0 {
		return 1.0
	}
	return plant2DDiagram.Zoom
}

func (plant2DDiagram *Plant2DDiagram) drawReferenceRhombus(stager *Stager, layer *svg.Layer, plant *PlantAbstract) {

	if plant2DDiagram.IsHiddenReferenceRhombus {
		return
	}

	angleRad := plant.RhombusInsideAngle * math.Pi / 180.0
	length := plant.RhombusSideLength * plant2DDiagram.getZoom()

	// Vertices
	v0x := plant2DDiagram.OriginX
	v0y := plant2DDiagram.OriginY

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

	polygon.Name = plant.RhombusStuff.ReferenceRhombus.Name
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

func (plant2DDiagram *Plant2DDiagram) drawGridPathShape(stager *Stager, layer *svg.Layer, plant *PlantAbstract) {

	if plant2DDiagram.IsHiddenGridPathShape {
		return
	}

	angleRad := plant.RhombusInsideAngle * math.Pi / 180.0
	length := plant.RhombusSideLength * plant2DDiagram.getZoom()

	// SVG Y-axis is inverted
	v1x := length * math.Cos(angleRad/2.0)
	v1y := -length * math.Sin(angleRad/2.0)

	v2x := -length * math.Cos(angleRad/2.0)
	v2y := -length * math.Sin(angleRad/2.0)

	polyline := new(svg.Polyline)
	layer.Polylines = append(layer.Polylines, polyline)

	polyline.Name = plant.RhombusStuff.GridPathShape.Name
	polyline.Presentation.Stroke = "red"
	polyline.Presentation.StrokeWidth = 2.0
	polyline.Presentation.StrokeOpacity = 1.0
	polyline.Presentation.FillOpacity = 0.0

	var points []string
	currX := plant2DDiagram.OriginX
	currY := plant2DDiagram.OriginY
	points = append(points, fmt.Sprintf("%f,%f", currX, currY))

	addCircle := func(x, y float64, stepIndex int, path string) {
		circle := new(svg.Circle)
		layer.Circles = append(layer.Circles, circle)
		circle.Name = fmt.Sprintf("%s-%s-step-%d", plant.RhombusStuff.GridPathShape.Name, path, stepIndex)
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

func (plant2DDiagram *Plant2DDiagram) drawRotatedPlantCircumferenceShape(stager *Stager, layer *svg.Layer, plant *PlantAbstract) {

	if plant2DDiagram.IsHiddenRotatedPlantCircumferenceShape {
		return
	}

	angleRad := plant.RhombusStuff.RotatedPlantCircumferenceShape.AngleDegree * math.Pi / 180.0
	length := plant.RhombusStuff.RotatedPlantCircumferenceShape.Length

	// SVG Y-axis is inverted
	endX := plant2DDiagram.OriginX + length*math.Cos(angleRad)
	endY := plant2DDiagram.OriginY - length*math.Sin(angleRad)

	line := new(svg.Line)
	layer.Lines = append(layer.Lines, line)

	line.Name = plant.RhombusStuff.RotatedPlantCircumferenceShape.Name
	line.X1 = plant2DDiagram.OriginX
	line.Y1 = plant2DDiagram.OriginY
	line.X2 = endX
	line.Y2 = endY

	line.Presentation.Stroke = "darkgreen" // slightly different color to distinguish
	line.Presentation.StrokeWidth = 2.0
	line.Presentation.StrokeOpacity = 1.0
	line.Presentation.StrokeDashArray = "5, 5" // make it dashed
}

func (plant2DDiagram *Plant2DDiagram) drawRotatedReferenceRhombus(stager *Stager, layer *svg.Layer, plant *PlantAbstract) {

	if plant2DDiagram.IsHiddenRotatedReferenceRhombus {
		return
	}

	angleRad := plant.RhombusInsideAngle * math.Pi / 180.0
	length := plant.RhombusSideLength * plant2DDiagram.getZoom()

	// Vertices
	v0x := plant2DDiagram.OriginX
	v0y := plant2DDiagram.OriginY

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

	polygon.Name = plant.RhombusStuff.RotatedReferenceRhombus.Name
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
	polygon.Presentation.Transform = fmt.Sprintf("rotate(%f %f %f)", plant.RhombusStuff.PlantCircumferenceShape.AngleDegree, plant2DDiagram.OriginX, plant2DDiagram.OriginY)
}

func (plant2DDiagram *Plant2DDiagram) drawRotatedGridPathShape(stager *Stager, layer *svg.Layer, plant *PlantAbstract) {

	if plant2DDiagram.IsHiddenRotatedGridPathShape {
		return
	}

	angleRad := plant.RhombusInsideAngle * math.Pi / 180.0
	length := plant.RhombusSideLength * plant2DDiagram.getZoom()

	// SVG Y-axis is inverted
	v1x := length * math.Cos(angleRad/2.0)
	v1y := -length * math.Sin(angleRad/2.0)

	v2x := -length * math.Cos(angleRad/2.0)
	v2y := -length * math.Sin(angleRad/2.0)

	polyline := new(svg.Polyline)
	layer.Polylines = append(layer.Polylines, polyline)

	polyline.Name = plant.RhombusStuff.RotatedGridPathShape.Name
	polyline.Presentation.Stroke = "darkred"
	polyline.Presentation.StrokeWidth = 2.0
	polyline.Presentation.StrokeOpacity = 1.0
	polyline.Presentation.FillOpacity = 0.0
	polyline.Presentation.StrokeDashArray = "5, 5"
	polyline.Presentation.Transform = fmt.Sprintf("rotate(%f %f %f)", plant.RhombusStuff.PlantCircumferenceShape.AngleDegree, plant2DDiagram.OriginX, plant2DDiagram.OriginY)

	var points []string
	currX := plant2DDiagram.OriginX
	currY := plant2DDiagram.OriginY
	points = append(points, fmt.Sprintf("%f,%f", currX, currY))

	addCircle := func(x, y float64, stepIndex int, path string) {
		circle := new(svg.Circle)
		layer.Circles = append(layer.Circles, circle)
		circle.Name = fmt.Sprintf("%s-%s-step-%d", plant.RhombusStuff.RotatedGridPathShape.Name, path, stepIndex)
		circle.CX = x
		circle.CY = y
		circle.Radius = 4.0
		circle.Presentation.Stroke = "darkred"
		circle.Presentation.StrokeWidth = 1.0
		circle.Presentation.StrokeOpacity = 1.0
		circle.Presentation.Color = "white"
		circle.Presentation.FillOpacity = 1.0
		circle.Presentation.Transform = fmt.Sprintf("rotate(%f %f %f)", plant.RhombusStuff.PlantCircumferenceShape.AngleDegree, plant2DDiagram.OriginX, plant2DDiagram.OriginY)
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

func (plant2DDiagram *Plant2DDiagram) drawRhombusGridShape(stager *Stager, layer *svg.Layer, plant *PlantAbstract) {

	if plant2DDiagram.IsHiddenRhombusGridShape {
		return
	}

	angleRad := plant.RhombusInsideAngle * math.Pi / 180.0
	length := plant.RhombusSideLength * plant2DDiagram.getZoom()

	// SVG Y-axis is inverted
	v1x := length * math.Cos(angleRad/2.0)
	v1y := -length * math.Sin(angleRad/2.0)

	v2x := -length * math.Cos(angleRad/2.0)
	v2y := -length * math.Sin(angleRad/2.0)

	for _, rhombus := range plant.RhombusStuff.InitialRhombusGridShape.InitialRhombusShapes {
		polygon := new(svg.Polygone)
		layer.Polygones = append(layer.Polygones, polygon)

		polygon.Name = rhombus.Name

		// r.X and r.Y are Cartesian center coordinates
		svg_cx := plant2DDiagram.OriginX + rhombus.X
		svg_cy := plant2DDiagram.OriginY - rhombus.Y

		// Calculate v0 (bottom-left vertex in visual SVG space) from the center
		v0x := svg_cx - (v1x+v2x)/2.0
		v0y := svg_cy - (v1y+v2y)/2.0

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

func (plant2DDiagram *Plant2DDiagram) drawRotatedRhombusGridShape(stager *Stager, layer *svg.Layer, plant *PlantAbstract) {

	if plant2DDiagram.IsHiddenRotatedRhombusGridShape {
		return
	}

	angleRad := plant.RhombusInsideAngle * math.Pi / 180.0
	length := plant.RhombusSideLength * plant2DDiagram.getZoom()

	// Cartesian vectors
	v1x := length * math.Cos(angleRad/2.0)
	v1y := length * math.Sin(angleRad/2.0)

	v2x := -length * math.Cos(angleRad/2.0)
	v2y := length * math.Sin(angleRad/2.0)

	rotRad := -plant.RhombusStuff.PlantCircumferenceShape.AngleDegree * math.Pi / 180.0
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

	for _, rhombus := range plant.RhombusStuff.RotatedRhombusGridShape2.RotatedRhombusShapes {
		polygon := new(svg.Polygone)
		layer.Polygones = append(layer.Polygones, polygon)

		polygon.Name = rhombus.Name

		// r.X and r.Y are Cartesian center coordinates
		svg_cx := plant2DDiagram.OriginX + rhombus.X
		svg_cy := plant2DDiagram.OriginY - rhombus.Y

		// Calculate v0 (bottom-left vertex in visual SVG space) from the center
		v0x := svg_cx - (v1_rot_x+v2_rot_x)/2.0
		v0y := svg_cy - (v1_rot_y+v2_rot_y)/2.0

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

func (plant2DDiagram *Plant2DDiagram) drawGrowthPathRhombusGridShape(stager *Stager, layer *svg.Layer, plant *PlantAbstract) {

	if plant2DDiagram.IsHiddenGrowthPathRhombusGridShape {
		return
	}

	angleRad := plant.RhombusInsideAngle * math.Pi / 180.0
	length := plant.RhombusSideLength * plant2DDiagram.getZoom()

	// Cartesian vectors
	v1x := length * math.Cos(angleRad/2.0)
	v1y := length * math.Sin(angleRad/2.0)

	v2x := -length * math.Cos(angleRad/2.0)
	v2y := length * math.Sin(angleRad/2.0)

	rotRad := -plant.RhombusStuff.PlantCircumferenceShape.AngleDegree * math.Pi / 180.0
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

	for _, rhombus := range plant.RhombusStuff.GrowthCurveRhombusGridShape.GrowthCurveRhombusShapes {
		polygon := new(svg.Polygone)
		layer.Polygones = append(layer.Polygones, polygon)

		polygon.Name = rhombus.Name

		// r.X and r.Y are Cartesian center coordinates
		svg_cx := plant2DDiagram.OriginX + rhombus.X
		svg_cy := plant2DDiagram.OriginY - rhombus.Y

		// Calculate v0 (bottom-left vertex in visual SVG space) from the center
		v0x := svg_cx - (v1_rot_x+v2_rot_x)/2.0
		v0y := svg_cy - (v1_rot_y+v2_rot_y)/2.0

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

func (plant2DDiagram *Plant2DDiagram) drawGrowthVectorShape(stager *Stager, layer *svg.Layer, plant *PlantAbstract) {
	if plant2DDiagram.IsHiddenGrowthVectorShape {
		return
	}
	if len(plant.RhombusStuff.GrowthCurveRhombusGridShape.GrowthCurveRhombusShapes) < 2 {
		return
	}

	rhombuses := plant.RhombusStuff.GrowthCurveRhombusGridShape.GrowthCurveRhombusShapes
	first := rhombuses[0]

	line := new(svg.Line)
	layer.Lines = append(layer.Lines, line)
	line.Name = plant.GrowthVectorShape.Name

	svg_x1 := plant2DDiagram.OriginX + first.X
	svg_y1 := plant2DDiagram.OriginY - first.Y

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

func (plant2DDiagram *Plant2DDiagram) drawPerpendicularVectorGrid(stager *Stager, layer *svg.Layer, plant *PlantAbstract) {
	if plant2DDiagram.IsHiddenPerpendicularVectorGrid {
		return
	}

	for _, vec := range plant.PerpendicularVectorGrid.PerpendicularVectors {
		line := new(svg.Line)
		layer.Lines = append(layer.Lines, line)
		line.Name = vec.Name

		svg_x1 := plant2DDiagram.OriginX + vec.StartX
		svg_y1 := plant2DDiagram.OriginY - vec.StartY

		svg_x2 := plant2DDiagram.OriginX + vec.EndX
		svg_y2 := plant2DDiagram.OriginY - vec.EndY

		line.X1 = svg_x1
		line.Y1 = svg_y1
		line.X2 = svg_x2
		line.Y2 = svg_y2

		line.Presentation.Stroke = "green"
		line.Presentation.StrokeWidth = 2.0
		line.Presentation.StrokeOpacity = 1.0
	}
}

func (plant2DDiagram *Plant2DDiagram) drawPerpendicularVectorGridHalfway(stager *Stager, layer *svg.Layer, plant *PlantAbstract, vase2DDiagram *Vase2DDiagram) {
	if vase2DDiagram == nil || vase2DDiagram.IsHiddenPerpendicularVectorGridHalfway {
		return
	}
	if plant.VaseAbstract == nil || plant.VaseAbstract.PerpendicularVectorGridHalfway == nil {
		return
	}

	for _, vec := range plant.VaseAbstract.PerpendicularVectorGridHalfway.PerpendicularVectorHalfways {
		line := new(svg.Line)
		layer.Lines = append(layer.Lines, line)
		line.Name = vec.Name

		svg_x1 := plant2DDiagram.OriginX + vec.StartX
		svg_y1 := plant2DDiagram.OriginY - vec.StartY

		svg_x2 := plant2DDiagram.OriginX + vec.EndX
		svg_y2 := plant2DDiagram.OriginY - vec.EndY

		line.X1 = svg_x1
		line.Y1 = svg_y1
		line.X2 = svg_x2
		line.Y2 = svg_y2

		line.Presentation.Stroke = "purple" // Distinct color for halfway
		line.Presentation.StrokeWidth = 2.0
		line.Presentation.StrokeOpacity = 1.0
	}
}

func (plant2DDiagram *Plant2DDiagram) drawBaseVectorShapeGrid(stager *Stager, layer *svg.Layer, plant *PlantAbstract) {
	if plant2DDiagram.IsHiddenBaseVectorShapeGrid {
		return
	}

	for _, base := range plant.BaseVectorShapeGrid.BaseVectorShapes {
		line := new(svg.Line)
		layer.Lines = append(layer.Lines, line)

		line.Name = base.Name

		line.X1 = plant2DDiagram.OriginX + base.StartX
		line.Y1 = plant2DDiagram.OriginY - base.StartY
		line.X2 = plant2DDiagram.OriginX + base.EndX
		line.Y2 = plant2DDiagram.OriginY - base.EndY

		line.Presentation.Stroke = "blue"
		line.Presentation.StrokeWidth = 2.0
		line.Presentation.StrokeOpacity = 1.0
	}
}

func (plant2DDiagram *Plant2DDiagram) drawArcNormalVectorShapeGrid(stager *Stager, layer *svg.Layer, plant *PlantAbstract) {
	if plant2DDiagram.IsHiddenArcNormalVectorShapeGrid {
		return
	}

	for _, shape := range plant.ArcNormalVectorShapeGrid.ArcNormalVectorShapes {
		line := new(svg.Line)
		layer.Lines = append(layer.Lines, line)
		line.Name = shape.Name
		line.X1 = plant2DDiagram.OriginX + shape.StartX
		line.Y1 = plant2DDiagram.OriginY - shape.StartY
		line.X2 = plant2DDiagram.OriginX + shape.EndX
		line.Y2 = plant2DDiagram.OriginY - shape.EndY

		line.Presentation.Stroke = "dodgerblue"
		line.Presentation.StrokeWidth = 2.0
		line.Presentation.StrokeOpacity = 1.0
	}
}

func (plant2DDiagram *Plant2DDiagram) drawStartArcShapeV2Grid(stager *Stager, layer *svg.Layer, plant *PlantAbstract) {
	if plant2DDiagram.IsHiddenStartArcShapeGrid {
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
			plant2DDiagram.OriginX+arc.StartX, plant2DDiagram.OriginY-arc.StartY,
			arc.RadiusX, arc.RadiusY, arc.XAxisRotation, largeArcFlag, sweepFlag,
			plant2DDiagram.OriginX+arc.EndX, plant2DDiagram.OriginY-arc.EndY,
		)

		path.Presentation.Stroke = "darkorange"
		path.Presentation.StrokeWidth = 3.0
		path.Presentation.StrokeOpacity = 1.0
		path.Presentation.FillOpacity = 0.0
	}
}

func (plant2DDiagram *Plant2DDiagram) drawTopStartArcShapeV2Grid(stager *Stager, layer *svg.Layer, plant *PlantAbstract, vase2DDiagram *Vase2DDiagram) {
	if vase2DDiagram == nil || vase2DDiagram.IsHiddenTopStartArcShapeGrid {
		return
	}
	if plant.VaseAbstract == nil || plant.VaseAbstract.TopStartArcShapeGrid == nil {
		return
	}

	for _, arc := range plant.VaseAbstract.TopStartArcShapeGrid.TopStartArcShapes {
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
			plant2DDiagram.OriginX+arc.StartX, plant2DDiagram.OriginY-arc.StartY,
			arc.RadiusX, arc.RadiusY, arc.XAxisRotation, largeArcFlag, sweepFlag,
			plant2DDiagram.OriginX+arc.EndX, plant2DDiagram.OriginY-arc.EndY,
		)

		path.Presentation.Stroke = "cyan"
		path.Presentation.StrokeWidth = 3.0
		path.Presentation.StrokeOpacity = 1.0
		path.Presentation.FillOpacity = 0.0
	}
}

func (plant2DDiagram *Plant2DDiagram) drawEndArcShapeV2Grid(stager *Stager, layer *svg.Layer, plant *PlantAbstract) {
	if plant2DDiagram.IsHiddenEndArcShapeGrid {
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
			plant2DDiagram.OriginX+arc.StartX, plant2DDiagram.OriginY-arc.StartY,
			arc.RadiusX, arc.RadiusY, arc.XAxisRotation, largeArcFlag, sweepFlag,
			plant2DDiagram.OriginX+arc.EndX, plant2DDiagram.OriginY-arc.EndY,
		)

		path.Presentation.Stroke = "purple"
		path.Presentation.StrokeWidth = 3.0
		path.Presentation.StrokeOpacity = 1.0
		path.Presentation.FillOpacity = 0.0
	}
}

func (plant2DDiagram *Plant2DDiagram) drawTopEndArcShapeV2Grid(stager *Stager, layer *svg.Layer, plant *PlantAbstract, vase2DDiagram *Vase2DDiagram) {
	if vase2DDiagram == nil || vase2DDiagram.IsHiddenTopEndArcShapeGrid {
		return
	}
	if plant.VaseAbstract == nil || plant.VaseAbstract.TopEndArcShapeGrid == nil {
		return
	}

	for _, arc := range plant.VaseAbstract.TopEndArcShapeGrid.TopEndArcShapes {
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
			plant2DDiagram.OriginX+arc.StartX, plant2DDiagram.OriginY-arc.StartY,
			arc.RadiusX, arc.RadiusY, arc.XAxisRotation, largeArcFlag, sweepFlag,
			plant2DDiagram.OriginX+arc.EndX, plant2DDiagram.OriginY-arc.EndY,
		)

		path.Presentation.Stroke = "cyan"
		path.Presentation.StrokeWidth = 3.0
		path.Presentation.StrokeOpacity = 1.0
		path.Presentation.FillOpacity = 0.0
	}
}

func (plant2DDiagram *Plant2DDiagram) drawStackOfGrowthCurveV2(stager *Stager, layer *svg.Layer, plant *PlantAbstract, vase2DDiagram *Vase2DDiagram) {
	if vase2DDiagram == nil || vase2DDiagram.IsHiddenStackOfGrowthCurve {
		return
	}
	if plant.VaseAbstract == nil || plant.VaseAbstract.StackOfRotatedGrowthCurve2D == nil {
		return
	}

	for _, sa := range plant.VaseAbstract.StackOfRotatedGrowthCurve2D.StackRotatedGrowthCurve2DStartArcShapes {
		path := new(svg.Path)
		layer.Paths = append(layer.Paths, path)
		path.Name = sa.Name

		sweepFlagStr := "0"
		if sa.SweepFlag {
			sweepFlagStr = "1"
		}
		largeArcFlagStr := "0"
		if sa.LargeArcFlag {
			largeArcFlagStr = "1"
		}

		path.Definition = fmt.Sprintf("M %0.1f %0.1f A %0.1f %0.1f %0.1f %s %s %0.1f %0.1f",
			plant2DDiagram.OriginX+sa.StartX, plant2DDiagram.OriginY-sa.StartY,
			sa.RadiusX, sa.RadiusY,
			sa.XAxisRotation, largeArcFlagStr, sweepFlagStr,
			plant2DDiagram.OriginX+sa.EndX, plant2DDiagram.OriginY-sa.EndY,
		)
		path.Presentation.Stroke = "blue"
		path.Presentation.StrokeWidth = 2.0
		path.Presentation.FillOpacity = 0.0
		path.Presentation.StrokeOpacity = 0.6
	}

	for _, ea := range plant.VaseAbstract.StackOfRotatedGrowthCurve2D.StackRotatedGrowthCurve2DEndArcShapes {
		path := new(svg.Path)
		layer.Paths = append(layer.Paths, path)
		path.Name = ea.Name

		sweepFlagStr := "0"
		if ea.SweepFlag {
			sweepFlagStr = "1"
		}
		largeArcFlagStr := "0"
		if ea.LargeArcFlag {
			largeArcFlagStr = "1"
		}

		path.Definition = fmt.Sprintf("M %0.1f %0.1f A %0.1f %0.1f %0.1f %s %s %0.1f %0.1f",
			plant2DDiagram.OriginX+ea.StartX, plant2DDiagram.OriginY-ea.StartY,
			ea.RadiusX, ea.RadiusY,
			ea.XAxisRotation, largeArcFlagStr, sweepFlagStr,
			plant2DDiagram.OriginX+ea.EndX, plant2DDiagram.OriginY-ea.EndY,
		)
		path.Presentation.Stroke = "purple"
		path.Presentation.StrokeWidth = 2.0
		path.Presentation.FillOpacity = 0.0
		path.Presentation.StrokeOpacity = 0.6
	}
}

func (plant2DDiagram *Plant2DDiagram) drawTopStackOfGrowthCurveV2(stager *Stager, layer *svg.Layer, plant *PlantAbstract, vase2DDiagram *Vase2DDiagram) {
	if vase2DDiagram == nil || vase2DDiagram.IsHiddenTopStackOfGrowthCurve {
		return
	}
	if plant.VaseAbstract == nil || plant.VaseAbstract.TopStackOfRotatedGrowthCurve2D == nil {
		return
	}

	for _, sa := range plant.VaseAbstract.TopStackOfRotatedGrowthCurve2D.TopStackOfRotatedGrowthCurve2DStartArcShapes {
		path := new(svg.Path)
		layer.Paths = append(layer.Paths, path)
		path.Name = sa.Name

		sweepFlagStr := "0"
		if sa.SweepFlag {
			sweepFlagStr = "1"
		}
		largeArcFlagStr := "0"
		if sa.LargeArcFlag {
			largeArcFlagStr = "1"
		}

		path.Definition = fmt.Sprintf("M %0.1f %0.1f A %0.1f %0.1f %0.1f %s %s %0.1f %0.1f",
			plant2DDiagram.OriginX+sa.StartX, plant2DDiagram.OriginY-sa.StartY,
			sa.RadiusX, sa.RadiusY,
			sa.XAxisRotation, largeArcFlagStr, sweepFlagStr,
			plant2DDiagram.OriginX+sa.EndX, plant2DDiagram.OriginY-sa.EndY,
		)
		path.Presentation.Stroke = "blue"
		path.Presentation.StrokeWidth = 2.0
		path.Presentation.FillOpacity = 0.0
		path.Presentation.StrokeOpacity = 0.6
	}

	for _, ea := range plant.VaseAbstract.TopStackOfRotatedGrowthCurve2D.TopStackOfRotatedGrowthCurve2DEndArcShapes {
		path := new(svg.Path)
		layer.Paths = append(layer.Paths, path)
		path.Name = ea.Name

		sweepFlagStr := "0"
		if ea.SweepFlag {
			sweepFlagStr = "1"
		}
		largeArcFlagStr := "0"
		if ea.LargeArcFlag {
			largeArcFlagStr = "1"
		}

		path.Definition = fmt.Sprintf("M %0.1f %0.1f A %0.1f %0.1f %0.1f %s %s %0.1f %0.1f",
			plant2DDiagram.OriginX+ea.StartX, plant2DDiagram.OriginY-ea.StartY,
			ea.RadiusX, ea.RadiusY,
			ea.XAxisRotation, largeArcFlagStr, sweepFlagStr,
			plant2DDiagram.OriginX+ea.EndX, plant2DDiagram.OriginY-ea.EndY,
		)
		path.Presentation.Stroke = "purple"
		path.Presentation.StrokeWidth = 2.0
		path.Presentation.FillOpacity = 0.0
		path.Presentation.StrokeOpacity = 0.6
	}
}

func (plant2DDiagram *Plant2DDiagram) drawGrowthCurve2D(stager *Stager, layer *svg.Layer, plant *PlantAbstract) {
	if plant2DDiagram.IsHiddenGrowthCurve2D {
		return
	}

	originalStartHidden := plant2DDiagram.IsHiddenStartArcShapeGrid
	originalEndHidden := plant2DDiagram.IsHiddenEndArcShapeGrid

	plant2DDiagram.IsHiddenStartArcShapeGrid = false
	plant2DDiagram.IsHiddenEndArcShapeGrid = false

	plant2DDiagram.drawStartArcShapeV2Grid(stager, layer, plant)
	plant2DDiagram.drawEndArcShapeV2Grid(stager, layer, plant)

	plant2DDiagram.IsHiddenStartArcShapeGrid = originalStartHidden
	plant2DDiagram.IsHiddenEndArcShapeGrid = originalEndHidden
}

// drawStackOfGrowthCurve2DByGrowthVector draws plant.StackHeight copies of GrowthCurve2D,
// each translated by k * GrowthVectorShape (k = 0..StackHeight-1).
func (plant2DDiagram *Plant2DDiagram) drawStackOfGrowthCurve2DByGrowthVector(stager *Stager, layer *svg.Layer, plant *PlantAbstract) {
	if plant2DDiagram.IsHiddenStackOfGrowthCurve2DByGrowthVector {
		return
	}
	if plant.GrowthVectorShape == nil || plant.StartArcShapeGrid == nil || plant.EndArcShapeGrid == nil {
		return
	}

	originalOriginX := plant2DDiagram.OriginX
	originalOriginY := plant2DDiagram.OriginY
	originalStartHidden := plant2DDiagram.IsHiddenStartArcShapeGrid
	originalEndHidden := plant2DDiagram.IsHiddenEndArcShapeGrid

	plant2DDiagram.IsHiddenStartArcShapeGrid = false
	plant2DDiagram.IsHiddenEndArcShapeGrid = false

	for k := 0; k < plant.StackHeight; k++ {
		// shift the origin by k growth-vector steps (SVG y-axis is inverted, hence the minus)
		plant2DDiagram.OriginX = originalOriginX + float64(k)*plant.GrowthVectorShape.X
		plant2DDiagram.OriginY = originalOriginY - float64(k)*plant.GrowthVectorShape.Y
		plant2DDiagram.drawStartArcShapeV2Grid(stager, layer, plant)
		plant2DDiagram.drawEndArcShapeV2Grid(stager, layer, plant)
	}

	plant2DDiagram.OriginX = originalOriginX
	plant2DDiagram.OriginY = originalOriginY
	plant2DDiagram.IsHiddenStartArcShapeGrid = originalStartHidden
	plant2DDiagram.IsHiddenEndArcShapeGrid = originalEndHidden
}

func (plant2DDiagram *Plant2DDiagram) drawTopGrowthCurve2D(stager *Stager, layer *svg.Layer, plant *PlantAbstract, vase2DDiagram *Vase2DDiagram) {
	if vase2DDiagram == nil || vase2DDiagram.IsHiddenTopGrowthCurve2D {
		return
	}

	originalStartHidden := vase2DDiagram.IsHiddenTopStartArcShapeGrid
	originalEndHidden := vase2DDiagram.IsHiddenTopEndArcShapeGrid

	vase2DDiagram.IsHiddenTopStartArcShapeGrid = false
	vase2DDiagram.IsHiddenTopEndArcShapeGrid = false

	plant2DDiagram.drawTopStartArcShapeV2Grid(stager, layer, plant, vase2DDiagram)
	plant2DDiagram.drawTopEndArcShapeV2Grid(stager, layer, plant, vase2DDiagram)

	vase2DDiagram.IsHiddenTopStartArcShapeGrid = originalStartHidden
	vase2DDiagram.IsHiddenTopEndArcShapeGrid = originalEndHidden
}

func (plant2DDiagram *Plant2DDiagram) drawShiftedLeftStackOfGrowthCurveV2(stager *Stager, layer *svg.Layer, plant *PlantAbstract, vase2DDiagram *Vase2DDiagram) {
	if vase2DDiagram == nil || vase2DDiagram.IsHiddenShiftedLeftStackOfGrowthCurve {
		return
	}
}

func (plant2DDiagram *Plant2DDiagram) drawShiftedBottomTopStartArcShapeV2Grid(stager *Stager, layer *svg.Layer, plant *PlantAbstract, vase2DDiagram *Vase2DDiagram) {
	if vase2DDiagram == nil || vase2DDiagram.IsHiddenShiftedBottomTopStartArcShapeGrid {
		return
	}
	if plant.VaseAbstract == nil || plant.VaseAbstract.ShiftedBottomTopStartArcShapeGrid == nil {
		return
	}

	for _, arc := range plant.VaseAbstract.ShiftedBottomTopStartArcShapeGrid.ShiftedBottomTopStartArcShapes {
		path := new(svg.Path)
		layer.Paths = append(layer.Paths, path)
		path.Name = arc.Name

		sweepFlagStr := "0"
		if arc.SweepFlag {
			sweepFlagStr = "1"
		}
		largeArcFlagStr := "0"
		if arc.LargeArcFlag {
			largeArcFlagStr = "1"
		}

		path.Definition = fmt.Sprintf("M %0.1f %0.1f A %0.1f %0.1f %0.1f %s %s %0.1f %0.1f",
			plant2DDiagram.OriginX+arc.StartX, plant2DDiagram.OriginY-arc.StartY,
			arc.RadiusX, arc.RadiusY,
			arc.XAxisRotation, largeArcFlagStr, sweepFlagStr,
			plant2DDiagram.OriginX+arc.EndX, plant2DDiagram.OriginY-arc.EndY,
		)
		path.Presentation.Stroke = "blue"
		path.Presentation.StrokeWidth = 2.0
		path.Presentation.FillOpacity = 0.0
		path.Presentation.StrokeOpacity = 1.0
	}
}

func (plant2DDiagram *Plant2DDiagram) drawMidArcVectorShapeGrid(stager *Stager, layer *svg.Layer, plant *PlantAbstract) {
	if plant2DDiagram.IsHiddenMidArcVectorShapeGrid {
		return
	}

	for _, base := range plant.MidArcVectorShapeGrid.MidArcVectorShapes {
		line := new(svg.Line)
		layer.Lines = append(layer.Lines, line)

		line.Name = base.Name

		line.X1 = plant2DDiagram.OriginX + base.StartX
		line.Y1 = plant2DDiagram.OriginY - base.StartY
		line.X2 = plant2DDiagram.OriginX + base.EndX
		line.Y2 = plant2DDiagram.OriginY - base.EndY

		line.Presentation.Stroke = "black"
		line.Presentation.StrokeWidth = 2.0
		line.Presentation.StrokeOpacity = 1.0
	}
}

func (plant2DDiagram *Plant2DDiagram) drawTopMidArcVectorShapeGrid(stager *Stager, layer *svg.Layer, plant *PlantAbstract, vase2DDiagram *Vase2DDiagram) {
	if vase2DDiagram == nil || vase2DDiagram.IsHiddenTopMidArcVectorShapeGrid {
		return
	}
	if plant.VaseAbstract == nil || plant.VaseAbstract.TopMidArcVectorShapeGrid == nil {
		return
	}

	for _, base := range plant.VaseAbstract.TopMidArcVectorShapeGrid.TopMidArcVectorShapes {
		line := new(svg.Line)
		layer.Lines = append(layer.Lines, line)

		line.Name = base.Name

		line.X1 = plant2DDiagram.OriginX + base.StartX
		line.Y1 = plant2DDiagram.OriginY - base.StartY
		line.X2 = plant2DDiagram.OriginX + base.EndX
		line.Y2 = plant2DDiagram.OriginY - base.EndY

		line.Presentation.Stroke = "black"
		line.Presentation.StrokeWidth = 2.0
		line.Presentation.StrokeOpacity = 1.0
	}
}

func (plant2DDiagram *Plant2DDiagram) drawStartHalfwayArcShapeGrid(stager *Stager, layer *svg.Layer, plant *PlantAbstract, vase2DDiagram *Vase2DDiagram) {
	if vase2DDiagram == nil || vase2DDiagram.IsHiddenStartHalfwayArcShapeGrid {
		return
	}
	if plant.VaseAbstract == nil || plant.VaseAbstract.StartHalfwayArcShapeGrid == nil {
		return
	}

	for _, base := range plant.VaseAbstract.StartHalfwayArcShapeGrid.StartHalfwayArcShapes {
		path := new(svg.Path)
		layer.Paths = append(layer.Paths, path)

		path.Name = base.Name

		pathStartX := plant2DDiagram.OriginX + base.StartX
		pathStartY := plant2DDiagram.OriginY - base.StartY
		pathEndX := plant2DDiagram.OriginX + base.EndX
		pathEndY := plant2DDiagram.OriginY - base.EndY

		largeArcFlag := 0
		if base.LargeArcFlag {
			largeArcFlag = 1
		}
		sweepFlag := 0
		if base.SweepFlag {
			sweepFlag = 1
		}

		path.Definition = fmt.Sprintf("M %f %f A %f %f %f %d %d %f %f",
			pathStartX, pathStartY,
			base.RadiusX, base.RadiusY,
			base.XAxisRotation,
			largeArcFlag, sweepFlag,
			pathEndX, pathEndY)

		path.Presentation.Stroke = "green"
		path.Presentation.StrokeWidth = 1.5
		path.Presentation.FillOpacity = 0.0
		path.Presentation.StrokeOpacity = 1.0
	}
}

func (plant2DDiagram *Plant2DDiagram) drawEndHalfwayArcShapeGrid(stager *Stager, layer *svg.Layer, plant *PlantAbstract, vase2DDiagram *Vase2DDiagram) {
	if vase2DDiagram == nil || vase2DDiagram.IsHiddenEndHalfwayArcShapeGrid {
		return
	}
	if plant.VaseAbstract == nil || plant.VaseAbstract.EndHalfwayArcShapeGrid == nil {
		return
	}

	for _, base := range plant.VaseAbstract.EndHalfwayArcShapeGrid.EndHalfwayArcShapes {
		path := new(svg.Path)
		layer.Paths = append(layer.Paths, path)

		path.Name = base.Name

		pathStartX := plant2DDiagram.OriginX + base.StartX
		pathStartY := plant2DDiagram.OriginY - base.StartY
		pathEndX := plant2DDiagram.OriginX + base.EndX
		pathEndY := plant2DDiagram.OriginY - base.EndY

		largeArcFlag := 0
		if base.LargeArcFlag {
			largeArcFlag = 1
		}
		sweepFlag := 0
		if base.SweepFlag {
			sweepFlag = 1
		}

		path.Definition = fmt.Sprintf("M %f %f A %f %f %f %d %d %f %f",
			pathStartX, pathStartY,
			base.RadiusX, base.RadiusY,
			base.XAxisRotation,
			largeArcFlag, sweepFlag,
			pathEndX, pathEndY)

		path.Presentation.Stroke = "green"
		path.Presentation.StrokeWidth = 1.5
		path.Presentation.FillOpacity = 0.0
		path.Presentation.StrokeOpacity = 1.0
	}
}

func (plant2DDiagram *Plant2DDiagram) drawTopStartHalfwayArcShapeGrid(stager *Stager, layer *svg.Layer, plant *PlantAbstract, vase2DDiagram *Vase2DDiagram) {
	if vase2DDiagram == nil || vase2DDiagram.IsHiddenTopStartHalfwayArcShapeGrid {
		return
	}
	if plant.VaseAbstract == nil || plant.VaseAbstract.TopStartHalfwayArcShapeGrid == nil {
		return
	}

	for _, base := range plant.VaseAbstract.TopStartHalfwayArcShapeGrid.TopStartHalfwayArcShapes {
		path := new(svg.Path)
		layer.Paths = append(layer.Paths, path)

		path.Name = base.Name

		pathStartX := plant2DDiagram.OriginX + base.StartX
		pathStartY := plant2DDiagram.OriginY - base.StartY
		pathEndX := plant2DDiagram.OriginX + base.EndX
		pathEndY := plant2DDiagram.OriginY - base.EndY

		largeArcFlag := 0
		if base.LargeArcFlag {
			largeArcFlag = 1
		}
		sweepFlag := 0
		if base.SweepFlag {
			sweepFlag = 1
		}

		path.Definition = fmt.Sprintf("M %f %f A %f %f %f %d %d %f %f",
			pathStartX, pathStartY,
			base.RadiusX, base.RadiusY,
			base.XAxisRotation,
			largeArcFlag, sweepFlag,
			pathEndX, pathEndY)

		path.Presentation.Stroke = "green"
		path.Presentation.StrokeWidth = 1.5
		path.Presentation.FillOpacity = 0.0
		path.Presentation.StrokeOpacity = 1.0
	}
}

func (plant2DDiagram *Plant2DDiagram) drawTopEndHalfwayArcShapeGrid(stager *Stager, layer *svg.Layer, plant *PlantAbstract, vase2DDiagram *Vase2DDiagram) {
	if vase2DDiagram == nil || vase2DDiagram.IsHiddenTopEndHalfwayArcShapeGrid {
		return
	}
	if plant.VaseAbstract == nil || plant.VaseAbstract.TopEndHalfwayArcShapeGrid == nil {
		return
	}

	for _, base := range plant.VaseAbstract.TopEndHalfwayArcShapeGrid.TopEndHalfwayArcShapes {
		path := new(svg.Path)
		layer.Paths = append(layer.Paths, path)

		path.Name = base.Name

		pathStartX := plant2DDiagram.OriginX + base.StartX
		pathStartY := plant2DDiagram.OriginY - base.StartY
		pathEndX := plant2DDiagram.OriginX + base.EndX
		pathEndY := plant2DDiagram.OriginY - base.EndY

		largeArcFlag := 0
		if base.LargeArcFlag {
			largeArcFlag = 1
		}
		sweepFlag := 0
		if base.SweepFlag {
			sweepFlag = 1
		}

		path.Definition = fmt.Sprintf("M %f %f A %f %f %f %d %d %f %f",
			pathStartX, pathStartY,
			base.RadiusX, base.RadiusY,
			base.XAxisRotation,
			largeArcFlag, sweepFlag,
			pathEndX, pathEndY)

		path.Presentation.Stroke = "green"
		path.Presentation.StrokeWidth = 1.5
		path.Presentation.FillOpacity = 0.0
		path.Presentation.StrokeOpacity = 1.0
	}
}

func (plant2DDiagram *Plant2DDiagram) drawStackOfGrowthCurve2D(stager *Stager, layer *svg.Layer, plant *PlantAbstract, vase2DDiagram *Vase2DDiagram) {
	if vase2DDiagram == nil || vase2DDiagram.IsHiddenStackOfGrowthCurve2D {
		return
	}
	if plant.VaseAbstract == nil || plant.VaseAbstract.StackOfGrowthCurve2D == nil {
		return
	}

	for _, start := range plant.VaseAbstract.StackOfGrowthCurve2D.StackGrowthCurve2DStartHalfwayArcShapes {
		path := new(svg.Path)
		layer.Paths = append(layer.Paths, path)

		path.Name = start.Name

		pathStartX := plant2DDiagram.OriginX + start.StartX
		pathStartY := plant2DDiagram.OriginY - start.StartY
		pathEndX := plant2DDiagram.OriginX + start.EndX
		pathEndY := plant2DDiagram.OriginY - start.EndY

		largeArcFlag := 0
		if start.LargeArcFlag {
			largeArcFlag = 1
		}
		sweepFlag := 0
		if start.SweepFlag {
			sweepFlag = 1
		}

		path.Definition = fmt.Sprintf("M %f %f A %f %f %f %d %d %f %f",
			pathStartX, pathStartY,
			start.RadiusX, start.RadiusY,
			start.XAxisRotation,
			largeArcFlag, sweepFlag,
			pathEndX, pathEndY)

		path.Presentation.Stroke = "blue"
		path.Presentation.StrokeWidth = 1.5
		path.Presentation.FillOpacity = 0.0
		path.Presentation.StrokeOpacity = 1.0
	}

	for _, end := range plant.VaseAbstract.StackOfGrowthCurve2D.StackGrowthCurve2DEndHalfwayArcShapes {
		path := new(svg.Path)
		layer.Paths = append(layer.Paths, path)

		path.Name = end.Name

		pathStartX := plant2DDiagram.OriginX + end.StartX
		pathStartY := plant2DDiagram.OriginY - end.StartY
		pathEndX := plant2DDiagram.OriginX + end.EndX
		pathEndY := plant2DDiagram.OriginY - end.EndY

		largeArcFlag := 0
		if end.LargeArcFlag {
			largeArcFlag = 1
		}
		sweepFlag := 0
		if end.SweepFlag {
			sweepFlag = 1
		}

		path.Definition = fmt.Sprintf("M %f %f A %f %f %f %d %d %f %f",
			pathStartX, pathStartY,
			end.RadiusX, end.RadiusY,
			end.XAxisRotation,
			largeArcFlag, sweepFlag,
			pathEndX, pathEndY)

		path.Presentation.Stroke = "blue"
		path.Presentation.StrokeWidth = 1.5
		path.Presentation.FillOpacity = 0.0
		path.Presentation.StrokeOpacity = 1.0
	}
}

func (plant2DDiagram *Plant2DDiagram) drawTopStackOfGrowthCurve2D(stager *Stager, layer *svg.Layer, plant *PlantAbstract, vase2DDiagram *Vase2DDiagram) {
	if vase2DDiagram == nil || vase2DDiagram.IsHiddenTopStackOfGrowthCurve2D {
		return
	}
	if plant.VaseAbstract == nil || plant.VaseAbstract.TopStackOfGrowthCurve2D == nil {
		return
	}

	for _, start := range plant.VaseAbstract.TopStackOfGrowthCurve2D.TopStackGrowthCurve2DStartHalfwayArcShapes {
		path := new(svg.Path)
		layer.Paths = append(layer.Paths, path)

		path.Name = start.Name

		pathStartX := plant2DDiagram.OriginX + start.StartX
		pathStartY := plant2DDiagram.OriginY - start.StartY
		pathEndX := plant2DDiagram.OriginX + start.EndX
		pathEndY := plant2DDiagram.OriginY - start.EndY

		largeArcFlag := 0
		if start.LargeArcFlag {
			largeArcFlag = 1
		}
		sweepFlag := 0
		if start.SweepFlag {
			sweepFlag = 1
		}

		path.Definition = fmt.Sprintf("M %f %f A %f %f %f %d %d %f %f",
			pathStartX, pathStartY,
			start.RadiusX, start.RadiusY,
			start.XAxisRotation,
			largeArcFlag, sweepFlag,
			pathEndX, pathEndY)

		path.Presentation.Stroke = "red"
		path.Presentation.StrokeWidth = 1.5
		path.Presentation.FillOpacity = 0.0
		path.Presentation.StrokeOpacity = 1.0
	}

	for _, end := range plant.VaseAbstract.TopStackOfGrowthCurve2D.TopStackGrowthCurve2DEndHalfwayArcShapes {
		path := new(svg.Path)
		layer.Paths = append(layer.Paths, path)

		path.Name = end.Name

		pathStartX := plant2DDiagram.OriginX + end.StartX
		pathStartY := plant2DDiagram.OriginY - end.StartY
		pathEndX := plant2DDiagram.OriginX + end.EndX
		pathEndY := plant2DDiagram.OriginY - end.EndY

		largeArcFlag := 0
		if end.LargeArcFlag {
			largeArcFlag = 1
		}
		sweepFlag := 0
		if end.SweepFlag {
			sweepFlag = 1
		}

		path.Definition = fmt.Sprintf("M %f %f A %f %f %f %d %d %f %f",
			pathStartX, pathStartY,
			end.RadiusX, end.RadiusY,
			end.XAxisRotation,
			largeArcFlag, sweepFlag,
			pathEndX, pathEndY)

		path.Presentation.Stroke = "red"
		path.Presentation.StrokeWidth = 1.5
		path.Presentation.FillOpacity = 0.0
		path.Presentation.StrokeOpacity = 1.0
	}
}

func (plant2DDiagram *Plant2DDiagram) drawStackOfGrowthCurve2DRibbon(stager *Stager, layer *svg.Layer, plant *PlantAbstract, vase2DDiagram *Vase2DDiagram) {
	if vase2DDiagram == nil || vase2DDiagram.IsHiddenStackOfGrowthCurve2DRibbon {
		return
	}
	if plant.VaseAbstract == nil || plant.VaseAbstract.StackOfGrowthCurve2DRibbon == nil {
		return
	}

	for i, start := range plant.VaseAbstract.StackOfGrowthCurve2DRibbon.StackGrowthCurve2DRibbonStartShapes {
		path := new(svg.Path)
		layer.Paths = append(layer.Paths, path)

		path.Name = start.Name

		bottomStartX := plant2DDiagram.OriginX + start.BottomStartX
		bottomStartY := plant2DDiagram.OriginY - start.BottomStartY
		bottomEndX := plant2DDiagram.OriginX + start.BottomEndX
		bottomEndY := plant2DDiagram.OriginY - start.BottomEndY

		topStartX := plant2DDiagram.OriginX + start.TopStartX
		topStartY := plant2DDiagram.OriginY - start.TopStartY
		topEndX := plant2DDiagram.OriginX + start.TopEndX
		topEndY := plant2DDiagram.OriginY - start.TopEndY

		bottomLargeArcFlag := 0
		if start.BottomLargeArcFlag {
			bottomLargeArcFlag = 1
		}
		bottomSweepFlag := 0
		if start.BottomSweepFlag {
			bottomSweepFlag = 1
		}

		topLargeArcFlag := 0
		if start.TopLargeArcFlag {
			topLargeArcFlag = 1
		}
		topSweepFlagRev := 1
		if start.TopSweepFlag {
			topSweepFlagRev = 0
		}

		path.Definition = fmt.Sprintf("M %f %f A %f %f %f %d %d %f %f L %f %f A %f %f %f %d %d %f %f Z",
			bottomStartX, bottomStartY,
			start.BottomRadiusX, start.BottomRadiusY,
			start.BottomXAxisRotation,
			bottomLargeArcFlag, bottomSweepFlag,
			bottomEndX, bottomEndY,
			topEndX, topEndY,
			start.TopRadiusX, start.TopRadiusY,
			start.TopXAxisRotation,
			topLargeArcFlag, topSweepFlagRev,
			topStartX, topStartY)

		path.Presentation.FillOpacity = 0.3
		path.Presentation.Color = "blue"
		if plant.PlantType == Vase && plant.VaseAbstract.HasAlternatingRingColors && i%2 != 0 {
			path.Presentation.Color = "saddlebrown"
		}
		path.Presentation.Stroke = "none"
	}

	for i, end := range plant.VaseAbstract.StackOfGrowthCurve2DRibbon.StackGrowthCurve2DRibbonEndShapes {
		path := new(svg.Path)
		layer.Paths = append(layer.Paths, path)

		path.Name = end.Name

		bottomStartX := plant2DDiagram.OriginX + end.BottomStartX
		bottomStartY := plant2DDiagram.OriginY - end.BottomStartY
		bottomEndX := plant2DDiagram.OriginX + end.BottomEndX
		bottomEndY := plant2DDiagram.OriginY - end.BottomEndY

		topStartX := plant2DDiagram.OriginX + end.TopStartX
		topStartY := plant2DDiagram.OriginY - end.TopStartY
		topEndX := plant2DDiagram.OriginX + end.TopEndX
		topEndY := plant2DDiagram.OriginY - end.TopEndY

		bottomLargeArcFlag := 0
		if end.BottomLargeArcFlag {
			bottomLargeArcFlag = 1
		}
		bottomSweepFlag := 0
		if end.BottomSweepFlag {
			bottomSweepFlag = 1
		}

		topLargeArcFlag := 0
		if end.TopLargeArcFlag {
			topLargeArcFlag = 1
		}
		topSweepFlagRev := 1
		if end.TopSweepFlag {
			topSweepFlagRev = 0
		}

		path.Definition = fmt.Sprintf("M %f %f A %f %f %f %d %d %f %f L %f %f A %f %f %f %d %d %f %f Z",
			bottomStartX, bottomStartY,
			end.BottomRadiusX, end.BottomRadiusY,
			end.BottomXAxisRotation,
			bottomLargeArcFlag, bottomSweepFlag,
			bottomEndX, bottomEndY,
			topEndX, topEndY,
			end.TopRadiusX, end.TopRadiusY,
			end.TopXAxisRotation,
			topLargeArcFlag, topSweepFlagRev,
			topStartX, topStartY)

		path.Presentation.FillOpacity = 0.3
		path.Presentation.Color = "blue"
		if plant.PlantType == Vase && plant.VaseAbstract.HasAlternatingRingColors && i%2 != 0 {
			path.Presentation.Color = "saddlebrown"
		}
		path.Presentation.Stroke = "none"
	}
}

func (plant2DDiagram *Plant2DDiagram) drawStackOfRotatedGrowthCurve2DRibbon(stager *Stager, layer *svg.Layer, plant *PlantAbstract, vase2DDiagram *Vase2DDiagram) {
	if vase2DDiagram == nil || vase2DDiagram.IsHiddenStackOfRotatedGrowthCurve2DRibbon {
		return
	}
	if plant.VaseAbstract == nil || plant.VaseAbstract.StackOfRotatedGrowthCurve2DRibbon == nil {
		return
	}

	for i, start := range plant.VaseAbstract.StackOfRotatedGrowthCurve2DRibbon.StackRotatedGrowthCurve2DRibbonStartShapes {
		path := new(svg.Path)
		layer.Paths = append(layer.Paths, path)

		path.Name = start.Name

		bottomStartX := plant2DDiagram.OriginX + start.BottomStartX
		bottomStartY := plant2DDiagram.OriginY - start.BottomStartY
		bottomEndX := plant2DDiagram.OriginX + start.BottomEndX
		bottomEndY := plant2DDiagram.OriginY - start.BottomEndY

		topStartX := plant2DDiagram.OriginX + start.TopStartX
		topStartY := plant2DDiagram.OriginY - start.TopStartY
		topEndX := plant2DDiagram.OriginX + start.TopEndX
		topEndY := plant2DDiagram.OriginY - start.TopEndY

		bottomLargeArcFlag := 0
		if start.BottomLargeArcFlag {
			bottomLargeArcFlag = 1
		}
		bottomSweepFlag := 0
		if start.BottomSweepFlag {
			bottomSweepFlag = 1
		}

		topLargeArcFlag := 0
		if start.TopLargeArcFlag {
			topLargeArcFlag = 1
		}
		topSweepFlagRev := 1
		if start.TopSweepFlag {
			topSweepFlagRev = 0
		}

		path.Definition = fmt.Sprintf("M %f %f A %f %f %f %d %d %f %f L %f %f A %f %f %f %d %d %f %f Z",
			bottomStartX, bottomStartY,
			start.BottomRadiusX, start.BottomRadiusY,
			start.BottomXAxisRotation,
			bottomLargeArcFlag, bottomSweepFlag,
			bottomEndX, bottomEndY,
			topEndX, topEndY,
			start.TopRadiusX, start.TopRadiusY,
			start.TopXAxisRotation,
			topLargeArcFlag, topSweepFlagRev,
			topStartX, topStartY)

		path.Presentation.FillOpacity = 0.3
		path.Presentation.Color = "purple"
		if plant.PlantType == Vase && plant.VaseAbstract.HasAlternatingRingColors && i%2 != 0 {
			path.Presentation.Color = "saddlebrown"
		}
		path.Presentation.Stroke = "none"
	}

	for i, end := range plant.VaseAbstract.StackOfRotatedGrowthCurve2DRibbon.StackRotatedGrowthCurve2DRibbonEndShapes {
		path := new(svg.Path)
		layer.Paths = append(layer.Paths, path)

		path.Name = end.Name

		bottomStartX := plant2DDiagram.OriginX + end.BottomStartX
		bottomStartY := plant2DDiagram.OriginY - end.BottomStartY
		bottomEndX := plant2DDiagram.OriginX + end.BottomEndX
		bottomEndY := plant2DDiagram.OriginY - end.BottomEndY

		topStartX := plant2DDiagram.OriginX + end.TopStartX
		topStartY := plant2DDiagram.OriginY - end.TopStartY
		topEndX := plant2DDiagram.OriginX + end.TopEndX
		topEndY := plant2DDiagram.OriginY - end.TopEndY

		bottomLargeArcFlag := 0
		if end.BottomLargeArcFlag {
			bottomLargeArcFlag = 1
		}
		bottomSweepFlag := 0
		if end.BottomSweepFlag {
			bottomSweepFlag = 1
		}

		topLargeArcFlag := 0
		if end.TopLargeArcFlag {
			topLargeArcFlag = 1
		}
		topSweepFlagRev := 1
		if end.TopSweepFlag {
			topSweepFlagRev = 0
		}

		path.Definition = fmt.Sprintf("M %f %f A %f %f %f %d %d %f %f L %f %f A %f %f %f %d %d %f %f Z",
			bottomStartX, bottomStartY,
			end.BottomRadiusX, end.BottomRadiusY,
			end.BottomXAxisRotation,
			bottomLargeArcFlag, bottomSweepFlag,
			bottomEndX, bottomEndY,
			topEndX, topEndY,
			end.TopRadiusX, end.TopRadiusY,
			end.TopXAxisRotation,
			topLargeArcFlag, topSweepFlagRev,
			topStartX, topStartY)

		path.Presentation.FillOpacity = 0.3
		path.Presentation.Color = "purple"
		if plant.PlantType == Vase && plant.VaseAbstract.HasAlternatingRingColors && i%2 != 0 {
			path.Presentation.Color = "saddlebrown"
		}
		path.Presentation.Stroke = "none"
	}
}

func (plant2DDiagram *Plant2DDiagram) drawPartiallyGrowthCurve2DRibbon(stager *Stager, layer *svg.Layer, plant *PlantAbstract, vase2DDiagram *Vase2DDiagram) {
	if vase2DDiagram == nil || vase2DDiagram.IsHiddenPartiallyGrowthCurve2DRibbon {
		return
	}
	if plant.VaseAbstract == nil || plant.VaseAbstract.PartiallyGrowthCurve2DRibbon == nil {
		return
	}

	for _, start := range plant.VaseAbstract.PartiallyGrowthCurve2DRibbon.PartiallyGrowthCurve2DRibbonStartShapes {
		path := new(svg.Path)
		layer.Paths = append(layer.Paths, path)

		path.Name = start.Name

		bottomStartX := plant2DDiagram.OriginX + start.BottomStartX
		bottomStartY := plant2DDiagram.OriginY - start.BottomStartY
		bottomEndX := plant2DDiagram.OriginX + start.BottomEndX
		bottomEndY := plant2DDiagram.OriginY - start.BottomEndY

		topStartX := plant2DDiagram.OriginX + start.TopStartX
		topStartY := plant2DDiagram.OriginY - start.TopStartY
		topEndX := plant2DDiagram.OriginX + start.TopEndX
		topEndY := plant2DDiagram.OriginY - start.TopEndY

		bottomLargeArcFlag := 0
		if start.BottomLargeArcFlag {
			bottomLargeArcFlag = 1
		}
		bottomSweepFlag := 0
		if start.BottomSweepFlag {
			bottomSweepFlag = 1
		}

		topLargeArcFlag := 0
		if start.TopLargeArcFlag {
			topLargeArcFlag = 1
		}
		topSweepFlagRev := 1
		if start.TopSweepFlag {
			topSweepFlagRev = 0
		}

		path.Definition = fmt.Sprintf("M %f %f A %f %f %f %d %d %f %f L %f %f A %f %f %f %d %d %f %f Z",
			bottomStartX, bottomStartY,
			start.BottomRadiusX, start.BottomRadiusY,
			start.BottomXAxisRotation,
			bottomLargeArcFlag, bottomSweepFlag,
			bottomEndX, bottomEndY,
			topEndX, topEndY,
			start.TopRadiusX, start.TopRadiusY,
			start.TopXAxisRotation,
			topLargeArcFlag, topSweepFlagRev,
			topStartX, topStartY)

		path.Presentation.FillOpacity = 0.3
		path.Presentation.Color = "green"
		path.Presentation.Stroke = "none"
	}

	for _, end := range plant.VaseAbstract.PartiallyGrowthCurve2DRibbon.PartiallyGrowthCurve2DRibbonEndShapes {
		path := new(svg.Path)
		layer.Paths = append(layer.Paths, path)

		path.Name = end.Name

		bottomStartX := plant2DDiagram.OriginX + end.BottomStartX
		bottomStartY := plant2DDiagram.OriginY - end.BottomStartY
		bottomEndX := plant2DDiagram.OriginX + end.BottomEndX
		bottomEndY := plant2DDiagram.OriginY - end.BottomEndY

		topStartX := plant2DDiagram.OriginX + end.TopStartX
		topStartY := plant2DDiagram.OriginY - end.TopStartY
		topEndX := plant2DDiagram.OriginX + end.TopEndX
		topEndY := plant2DDiagram.OriginY - end.TopEndY

		bottomLargeArcFlag := 0
		if end.BottomLargeArcFlag {
			bottomLargeArcFlag = 1
		}
		bottomSweepFlag := 0
		if end.BottomSweepFlag {
			bottomSweepFlag = 1
		}

		topLargeArcFlag := 0
		if end.TopLargeArcFlag {
			topLargeArcFlag = 1
		}
		topSweepFlagRev := 1
		if end.TopSweepFlag {
			topSweepFlagRev = 0
		}

		path.Definition = fmt.Sprintf("M %f %f A %f %f %f %d %d %f %f L %f %f A %f %f %f %d %d %f %f Z",
			bottomStartX, bottomStartY,
			end.BottomRadiusX, end.BottomRadiusY,
			end.BottomXAxisRotation,
			bottomLargeArcFlag, bottomSweepFlag,
			bottomEndX, bottomEndY,
			topEndX, topEndY,
			end.TopRadiusX, end.TopRadiusY,
			end.TopXAxisRotation,
			topLargeArcFlag, topSweepFlagRev,
			topStartX, topStartY)

		path.Presentation.FillOpacity = 0.3
		path.Presentation.Color = "green"
		path.Presentation.Stroke = "none"
	}
}

func (plant2DDiagram *Plant2DDiagram) drawShiftedLeftPartiallyGrowthCurve2DRibbon(stager *Stager, layer *svg.Layer, plant *PlantAbstract, vase2DDiagram *Vase2DDiagram) {
	if vase2DDiagram == nil || vase2DDiagram.IsHiddenShiftedLeftPartiallyGrowthCurve2DRibbon {
		return
	}
	if plant.VaseAbstract == nil || plant.VaseAbstract.ShiftedLeftPartiallyGrowthCurve2DRibbon == nil {
		return
	}

	// Translation by -PlantCircumferenceShape.Length
	dx := 0.0
	if plant.RhombusStuff != nil && plant.RhombusStuff.PlantCircumferenceShape != nil {
		dx = -plant.RhombusStuff.PlantCircumferenceShape.Length
	}

	for _, start := range plant.VaseAbstract.ShiftedLeftPartiallyGrowthCurve2DRibbon.ShiftedLeftPartiallyGrowthCurve2DRibbonStartShapes {
		path := new(svg.Path)
		layer.Paths = append(layer.Paths, path)

		path.Name = start.Name

		bottomStartX := plant2DDiagram.OriginX + start.BottomStartX + dx
		bottomStartY := plant2DDiagram.OriginY - start.BottomStartY
		bottomEndX := plant2DDiagram.OriginX + start.BottomEndX + dx
		bottomEndY := plant2DDiagram.OriginY - start.BottomEndY

		topStartX := plant2DDiagram.OriginX + start.TopStartX + dx
		topStartY := plant2DDiagram.OriginY - start.TopStartY
		topEndX := plant2DDiagram.OriginX + start.TopEndX + dx
		topEndY := plant2DDiagram.OriginY - start.TopEndY

		bottomLargeArcFlag := 0
		if start.BottomLargeArcFlag {
			bottomLargeArcFlag = 1
		}
		bottomSweepFlag := 0
		if start.BottomSweepFlag {
			bottomSweepFlag = 1
		}

		topLargeArcFlag := 0
		if start.TopLargeArcFlag {
			topLargeArcFlag = 1
		}
		topSweepFlagRev := 1
		if start.TopSweepFlag {
			topSweepFlagRev = 0
		}

		path.Definition = fmt.Sprintf("M %f %f A %f %f %f %d %d %f %f L %f %f A %f %f %f %d %d %f %f Z",
			bottomStartX, bottomStartY,
			start.BottomRadiusX, start.BottomRadiusY,
			start.BottomXAxisRotation,
			bottomLargeArcFlag, bottomSweepFlag,
			bottomEndX, bottomEndY,
			topEndX, topEndY,
			start.TopRadiusX, start.TopRadiusY,
			start.TopXAxisRotation,
			topLargeArcFlag, topSweepFlagRev,
			topStartX, topStartY)

		path.Presentation.FillOpacity = 0.3
		path.Presentation.Color = "green"
		path.Presentation.Stroke = "none"
	}

	for _, end := range plant.VaseAbstract.ShiftedLeftPartiallyGrowthCurve2DRibbon.ShiftedLeftPartiallyGrowthCurve2DRibbonEndShapes {
		path := new(svg.Path)
		layer.Paths = append(layer.Paths, path)

		path.Name = end.Name

		bottomStartX := plant2DDiagram.OriginX + end.BottomStartX + dx
		bottomStartY := plant2DDiagram.OriginY - end.BottomStartY
		bottomEndX := plant2DDiagram.OriginX + end.BottomEndX + dx
		bottomEndY := plant2DDiagram.OriginY - end.BottomEndY

		topStartX := plant2DDiagram.OriginX + end.TopStartX + dx
		topStartY := plant2DDiagram.OriginY - end.TopStartY
		topEndX := plant2DDiagram.OriginX + end.TopEndX + dx
		topEndY := plant2DDiagram.OriginY - end.TopEndY

		bottomLargeArcFlag := 0
		if end.BottomLargeArcFlag {
			bottomLargeArcFlag = 1
		}
		bottomSweepFlag := 0
		if end.BottomSweepFlag {
			bottomSweepFlag = 1
		}

		topLargeArcFlag := 0
		if end.TopLargeArcFlag {
			topLargeArcFlag = 1
		}
		topSweepFlagRev := 1
		if end.TopSweepFlag {
			topSweepFlagRev = 0
		}

		path.Definition = fmt.Sprintf("M %f %f A %f %f %f %d %d %f %f L %f %f A %f %f %f %d %d %f %f Z",
			bottomStartX, bottomStartY,
			end.BottomRadiusX, end.BottomRadiusY,
			end.BottomXAxisRotation,
			bottomLargeArcFlag, bottomSweepFlag,
			bottomEndX, bottomEndY,
			topEndX, topEndY,
			end.TopRadiusX, end.TopRadiusY,
			end.TopXAxisRotation,
			topLargeArcFlag, topSweepFlagRev,
			topStartX, topStartY)

		path.Presentation.FillOpacity = 0.3
		path.Presentation.Color = "green"
		path.Presentation.Stroke = "none"
	}
}

func (plant2DDiagram *Plant2DDiagram) drawGrowthCurve2DRibbon(stager *Stager, layer *svg.Layer, plant *PlantAbstract, vase2DDiagram *Vase2DDiagram) {
	if vase2DDiagram == nil || vase2DDiagram.IsHiddenGrowthCurve2DRibbon {
		return
	}
	if plant.VaseAbstract == nil || plant.VaseAbstract.GrowthCurve2DRibbon == nil {
		return
	}

	for _, start := range plant.VaseAbstract.GrowthCurve2DRibbon.GrowthCurve2DRibbonStartShapes {
		path := new(svg.Path)
		layer.Paths = append(layer.Paths, path)

		path.Name = start.Name

		bottomStartX := plant2DDiagram.OriginX + start.BottomStartX
		bottomStartY := plant2DDiagram.OriginY - start.BottomStartY
		bottomEndX := plant2DDiagram.OriginX + start.BottomEndX
		bottomEndY := plant2DDiagram.OriginY - start.BottomEndY

		topStartX := plant2DDiagram.OriginX + start.TopStartX
		topStartY := plant2DDiagram.OriginY - start.TopStartY
		topEndX := plant2DDiagram.OriginX + start.TopEndX
		topEndY := plant2DDiagram.OriginY - start.TopEndY

		bottomLargeArcFlag := 0
		if start.BottomLargeArcFlag {
			bottomLargeArcFlag = 1
		}
		bottomSweepFlag := 0
		if start.BottomSweepFlag {
			bottomSweepFlag = 1
		}

		topLargeArcFlag := 0
		if start.TopLargeArcFlag {
			topLargeArcFlag = 1
		}
		topSweepFlagRev := 1
		if start.TopSweepFlag {
			topSweepFlagRev = 0
		}

		path.Definition = fmt.Sprintf("M %f %f A %f %f %f %d %d %f %f L %f %f A %f %f %f %d %d %f %f Z",
			bottomStartX, bottomStartY,
			start.BottomRadiusX, start.BottomRadiusY,
			start.BottomXAxisRotation,
			bottomLargeArcFlag, bottomSweepFlag,
			bottomEndX, bottomEndY,
			topEndX, topEndY,
			start.TopRadiusX, start.TopRadiusY,
			start.TopXAxisRotation,
			topLargeArcFlag, topSweepFlagRev,
			topStartX, topStartY)

		path.Presentation.FillOpacity = 0.3
		path.Presentation.Color = "lightgreen"
		path.Presentation.Stroke = "none"
	}

	for _, end := range plant.VaseAbstract.GrowthCurve2DRibbon.GrowthCurve2DRibbonEndShapes {
		path := new(svg.Path)
		layer.Paths = append(layer.Paths, path)

		path.Name = end.Name

		bottomStartX := plant2DDiagram.OriginX + end.BottomStartX
		bottomStartY := plant2DDiagram.OriginY - end.BottomStartY
		bottomEndX := plant2DDiagram.OriginX + end.BottomEndX
		bottomEndY := plant2DDiagram.OriginY - end.BottomEndY

		topStartX := plant2DDiagram.OriginX + end.TopStartX
		topStartY := plant2DDiagram.OriginY - end.TopStartY
		topEndX := plant2DDiagram.OriginX + end.TopEndX
		topEndY := plant2DDiagram.OriginY - end.TopEndY

		bottomLargeArcFlag := 0
		if end.BottomLargeArcFlag {
			bottomLargeArcFlag = 1
		}
		bottomSweepFlag := 0
		if end.BottomSweepFlag {
			bottomSweepFlag = 1
		}

		topLargeArcFlag := 0
		if end.TopLargeArcFlag {
			topLargeArcFlag = 1
		}
		topSweepFlagRev := 1
		if end.TopSweepFlag {
			topSweepFlagRev = 0
		}

		path.Definition = fmt.Sprintf("M %f %f A %f %f %f %d %d %f %f L %f %f A %f %f %f %d %d %f %f Z",
			bottomStartX, bottomStartY,
			end.BottomRadiusX, end.BottomRadiusY,
			end.BottomXAxisRotation,
			bottomLargeArcFlag, bottomSweepFlag,
			bottomEndX, bottomEndY,
			topEndX, topEndY,
			end.TopRadiusX, end.TopRadiusY,
			end.TopXAxisRotation,
			topLargeArcFlag, topSweepFlagRev,
			topStartX, topStartY)

		path.Presentation.FillOpacity = 0.3
		path.Presentation.Color = "lightgreen"
		path.Presentation.Stroke = "none"
	}
}

func (plant2DDiagram *Plant2DDiagram) drawShiftedRightGrowthCurve2DRibbon(stager *Stager, layer *svg.Layer, plant *PlantAbstract, vase2DDiagram *Vase2DDiagram) {
	if vase2DDiagram == nil || vase2DDiagram.IsHiddenShiftedRightGrowthCurve2DRibbon {
		return
	}
	if plant.VaseAbstract == nil || plant.VaseAbstract.ShiftedRightGrowthCurve2DRibbon == nil {
		return
	}

	// Translation by PlantCircumferenceShape.Length
	dx := 0.0
	if plant.RhombusStuff != nil && plant.RhombusStuff.PlantCircumferenceShape != nil {
		dx = plant.RhombusStuff.PlantCircumferenceShape.Length
	}

	for _, start := range plant.VaseAbstract.ShiftedRightGrowthCurve2DRibbon.ShiftedRightGrowthCurve2DRibbonStartShapes {
		path := new(svg.Path)
		layer.Paths = append(layer.Paths, path)

		path.Name = start.Name

		bottomStartX := plant2DDiagram.OriginX + start.BottomStartX + dx
		bottomStartY := plant2DDiagram.OriginY - start.BottomStartY
		bottomEndX := plant2DDiagram.OriginX + start.BottomEndX + dx
		bottomEndY := plant2DDiagram.OriginY - start.BottomEndY

		topStartX := plant2DDiagram.OriginX + start.TopStartX + dx
		topStartY := plant2DDiagram.OriginY - start.TopStartY
		topEndX := plant2DDiagram.OriginX + start.TopEndX + dx
		topEndY := plant2DDiagram.OriginY - start.TopEndY

		bottomLargeArcFlag := 0
		if start.BottomLargeArcFlag {
			bottomLargeArcFlag = 1
		}
		bottomSweepFlag := 0
		if start.BottomSweepFlag {
			bottomSweepFlag = 1
		}

		topLargeArcFlag := 0
		if start.TopLargeArcFlag {
			topLargeArcFlag = 1
		}
		topSweepFlagRev := 1
		if start.TopSweepFlag {
			topSweepFlagRev = 0
		}

		path.Definition = fmt.Sprintf("M %f %f A %f %f %f %d %d %f %f L %f %f A %f %f %f %d %d %f %f Z",
			bottomStartX, bottomStartY,
			start.BottomRadiusX, start.BottomRadiusY,
			start.BottomXAxisRotation,
			bottomLargeArcFlag, bottomSweepFlag,
			bottomEndX, bottomEndY,
			topEndX, topEndY,
			start.TopRadiusX, start.TopRadiusY,
			start.TopXAxisRotation,
			topLargeArcFlag, topSweepFlagRev,
			topStartX, topStartY)

		path.Presentation.FillOpacity = 0.5
		path.Presentation.Color = "rosybrown"
		path.Presentation.Stroke = "none"
	}

	for _, end := range plant.VaseAbstract.ShiftedRightGrowthCurve2DRibbon.ShiftedRightGrowthCurve2DRibbonEndShapes {
		path := new(svg.Path)
		layer.Paths = append(layer.Paths, path)

		path.Name = end.Name

		bottomStartX := plant2DDiagram.OriginX + end.BottomStartX + dx
		bottomStartY := plant2DDiagram.OriginY - end.BottomStartY
		bottomEndX := plant2DDiagram.OriginX + end.BottomEndX + dx
		bottomEndY := plant2DDiagram.OriginY - end.BottomEndY

		topStartX := plant2DDiagram.OriginX + end.TopStartX + dx
		topStartY := plant2DDiagram.OriginY - end.TopStartY
		topEndX := plant2DDiagram.OriginX + end.TopEndX + dx
		topEndY := plant2DDiagram.OriginY - end.TopEndY

		bottomLargeArcFlag := 0
		if end.BottomLargeArcFlag {
			bottomLargeArcFlag = 1
		}
		bottomSweepFlag := 0
		if end.BottomSweepFlag {
			bottomSweepFlag = 1
		}

		topLargeArcFlag := 0
		if end.TopLargeArcFlag {
			topLargeArcFlag = 1
		}
		topSweepFlagRev := 1
		if end.TopSweepFlag {
			topSweepFlagRev = 0
		}

		path.Definition = fmt.Sprintf("M %f %f A %f %f %f %d %d %f %f L %f %f A %f %f %f %d %d %f %f Z",
			bottomStartX, bottomStartY,
			end.BottomRadiusX, end.BottomRadiusY,
			end.BottomXAxisRotation,
			bottomLargeArcFlag, bottomSweepFlag,
			bottomEndX, bottomEndY,
			topEndX, topEndY,
			end.TopRadiusX, end.TopRadiusY,
			end.TopXAxisRotation,
			topLargeArcFlag, topSweepFlagRev,
			topStartX, topStartY)

		path.Presentation.FillOpacity = 0.5
		path.Presentation.Color = "rosybrown"
		path.Presentation.Stroke = "none"
	}
}

func (plant2DDiagram *Plant2DDiagram) drawShiftedLeftGrowthCurve2DRibbon(stager *Stager, layer *svg.Layer, plant *PlantAbstract, vase2DDiagram *Vase2DDiagram) {
	if vase2DDiagram == nil || vase2DDiagram.IsHiddenShiftedLeftGrowthCurve2DRibbon {
		return
	}
	if plant.VaseAbstract == nil || plant.VaseAbstract.ShiftedLeftGrowthCurve2DRibbon == nil {
		return
	}

	// Translation by -PlantCircumferenceShape.Length
	dx := 0.0
	if plant.RhombusStuff != nil && plant.RhombusStuff.PlantCircumferenceShape != nil {
		dx = -plant.RhombusStuff.PlantCircumferenceShape.Length
	}

	for _, start := range plant.VaseAbstract.ShiftedLeftGrowthCurve2DRibbon.ShiftedLeftGrowthCurve2DRibbonStartShapes {
		path := new(svg.Path)
		layer.Paths = append(layer.Paths, path)

		path.Name = start.Name

		bottomStartX := plant2DDiagram.OriginX + start.BottomStartX + dx
		bottomStartY := plant2DDiagram.OriginY - start.BottomStartY
		bottomEndX := plant2DDiagram.OriginX + start.BottomEndX + dx
		bottomEndY := plant2DDiagram.OriginY - start.BottomEndY

		topStartX := plant2DDiagram.OriginX + start.TopStartX + dx
		topStartY := plant2DDiagram.OriginY - start.TopStartY
		topEndX := plant2DDiagram.OriginX + start.TopEndX + dx
		topEndY := plant2DDiagram.OriginY - start.TopEndY

		bottomLargeArcFlag := 0
		if start.BottomLargeArcFlag {
			bottomLargeArcFlag = 1
		}
		bottomSweepFlag := 0
		if start.BottomSweepFlag {
			bottomSweepFlag = 1
		}

		topLargeArcFlag := 0
		if start.TopLargeArcFlag {
			topLargeArcFlag = 1
		}
		topSweepFlagRev := 1
		if start.TopSweepFlag {
			topSweepFlagRev = 0
		}

		path.Definition = fmt.Sprintf("M %f %f A %f %f %f %d %d %f %f L %f %f A %f %f %f %d %d %f %f Z",
			bottomStartX, bottomStartY,
			start.BottomRadiusX, start.BottomRadiusY,
			start.BottomXAxisRotation,
			bottomLargeArcFlag, bottomSweepFlag,
			bottomEndX, bottomEndY,
			topEndX, topEndY,
			start.TopRadiusX, start.TopRadiusY,
			start.TopXAxisRotation,
			topLargeArcFlag, topSweepFlagRev,
			topStartX, topStartY)

		path.Presentation.FillOpacity = 0.5
		path.Presentation.Color = "rosybrown"
		path.Presentation.Stroke = "none"
	}

	for _, end := range plant.VaseAbstract.ShiftedLeftGrowthCurve2DRibbon.ShiftedLeftGrowthCurve2DRibbonEndShapes {
		path := new(svg.Path)
		layer.Paths = append(layer.Paths, path)

		path.Name = end.Name

		bottomStartX := plant2DDiagram.OriginX + end.BottomStartX + dx
		bottomStartY := plant2DDiagram.OriginY - end.BottomStartY
		bottomEndX := plant2DDiagram.OriginX + end.BottomEndX + dx
		bottomEndY := plant2DDiagram.OriginY - end.BottomEndY

		topStartX := plant2DDiagram.OriginX + end.TopStartX + dx
		topStartY := plant2DDiagram.OriginY - end.TopStartY
		topEndX := plant2DDiagram.OriginX + end.TopEndX + dx
		topEndY := plant2DDiagram.OriginY - end.TopEndY

		bottomLargeArcFlag := 0
		if end.BottomLargeArcFlag {
			bottomLargeArcFlag = 1
		}
		bottomSweepFlag := 0
		if end.BottomSweepFlag {
			bottomSweepFlag = 1
		}

		topLargeArcFlag := 0
		if end.TopLargeArcFlag {
			topLargeArcFlag = 1
		}
		topSweepFlagRev := 1
		if end.TopSweepFlag {
			topSweepFlagRev = 0
		}

		path.Definition = fmt.Sprintf("M %f %f A %f %f %f %d %d %f %f L %f %f A %f %f %f %d %d %f %f Z",
			bottomStartX, bottomStartY,
			end.BottomRadiusX, end.BottomRadiusY,
			end.BottomXAxisRotation,
			bottomLargeArcFlag, bottomSweepFlag,
			bottomEndX, bottomEndY,
			topEndX, topEndY,
			end.TopRadiusX, end.TopRadiusY,
			end.TopXAxisRotation,
			topLargeArcFlag, topSweepFlagRev,
			topStartX, topStartY)

		path.Presentation.FillOpacity = 0.5
		path.Presentation.Color = "rosybrown"
		path.Presentation.Stroke = "none"
	}
}

func (plant2DDiagram *Plant2DDiagram) drawPartiallyGrowthCurve2DTrajectory(stager *Stager, layer *svg.Layer, plant *PlantAbstract, vase2DDiagram *Vase2DDiagram) {
	if vase2DDiagram == nil || vase2DDiagram.IsHiddenPartiallyGrowthCurve2DTrajectory {
		return
	}
	if plant.VaseAbstract == nil || plant.VaseAbstract.PartiallyGrowthCurve2DTrajectory == nil {
		return
	}

	for _, shape := range plant.VaseAbstract.PartiallyGrowthCurve2DTrajectory.PartiallyGrowthCurve2DTrajectoryShapes {
		line := new(svg.Line)
		layer.Lines = append(layer.Lines, line)

		line.Name = shape.Name

		line.X1 = plant2DDiagram.OriginX + shape.StartX
		line.Y1 = plant2DDiagram.OriginY - shape.StartY
		line.X2 = plant2DDiagram.OriginX + shape.EndX
		line.Y2 = plant2DDiagram.OriginY - shape.EndY

		line.Presentation.Stroke = "purple"
		line.Presentation.StrokeWidth = 2.0
		line.Presentation.StrokeOpacity = 1.0
	}
}

func (plant2DDiagram *Plant2DDiagram) drawPartiallyGrowthCurve2DTrajectoryP1P2(stager *Stager, layer *svg.Layer, plant *PlantAbstract, vase2DDiagram *Vase2DDiagram) {
	if vase2DDiagram == nil || vase2DDiagram.IsHiddenPartiallyGrowthCurve2DTrajectoryP1P2 {
		return
	}
	if plant.VaseAbstract == nil || plant.VaseAbstract.PartiallyGrowthCurve2DTrajectoryP1P2 == nil {
		return
	}

	p1p2 := plant.VaseAbstract.PartiallyGrowthCurve2DTrajectoryP1P2

	// Draw P1 Dots
	for _, shape := range p1p2.P1PointShapes {
		circle := new(svg.Circle)
		layer.Circles = append(layer.Circles, circle)
		circle.Name = shape.Name
		circle.CX = plant2DDiagram.OriginX + shape.X
		circle.CY = plant2DDiagram.OriginY - shape.Y
		circle.Radius = 3.0
		circle.Presentation.Color = "magenta"
		circle.Presentation.FillOpacity = 1.0
		circle.Presentation.Stroke = "darkmagenta"
		circle.Presentation.StrokeWidth = 1.0
		circle.Presentation.StrokeOpacity = 1.0
	}

	// Draw P2 Dots
	for _, shape := range p1p2.P2PointShapes {
		circle := new(svg.Circle)
		layer.Circles = append(layer.Circles, circle)
		circle.Name = shape.Name
		circle.CX = plant2DDiagram.OriginX + shape.X
		circle.CY = plant2DDiagram.OriginY - shape.Y
		circle.Radius = 3.0
		circle.Presentation.Color = "magenta"
		circle.Presentation.FillOpacity = 1.0
		circle.Presentation.Stroke = "darkmagenta"
		circle.Presentation.StrokeWidth = 1.0
		circle.Presentation.StrokeOpacity = 1.0
	}

	// Draw P1 Text Label at step 0
	if len(p1p2.P1PointShapes) > 0 {
		p1_0 := p1p2.P1PointShapes[0]
		textP1 := new(svg.Text)
		layer.Texts = append(layer.Texts, textP1)
		textP1.Name = plant.Name + "-P1-Text"
		textP1.X = plant2DDiagram.OriginX + p1_0.X - 14
		textP1.Y = plant2DDiagram.OriginY - p1_0.Y - 6
		textP1.Content = "P1"
		textP1.Presentation.Color = "magenta"
		textP1.Presentation.FillOpacity = 1.0
	}

	// Draw P2 Text Label at step 0
	if len(p1p2.P2PointShapes) > 0 {
		p2_0 := p1p2.P2PointShapes[0]
		textP2 := new(svg.Text)
		layer.Texts = append(layer.Texts, textP2)
		textP2.Name = plant.Name + "-P2-Text"
		textP2.X = plant2DDiagram.OriginX + p2_0.X + 6
		textP2.Y = plant2DDiagram.OriginY - p2_0.Y - 6
		textP2.Content = "P2"
		textP2.Presentation.Color = "magenta"
		textP2.Presentation.FillOpacity = 1.0
	}

	// Draw P1 Curve Lines
	for _, shape := range p1p2.P1CurveShapes {
		line := new(svg.Line)
		layer.Lines = append(layer.Lines, line)
		line.Name = shape.Name
		line.X1 = plant2DDiagram.OriginX + shape.StartX
		line.Y1 = plant2DDiagram.OriginY - shape.StartY
		line.X2 = plant2DDiagram.OriginX + shape.EndX
		line.Y2 = plant2DDiagram.OriginY - shape.EndY
		line.Presentation.Stroke = "magenta"
		line.Presentation.StrokeWidth = 1.5
		line.Presentation.StrokeOpacity = 0.8
	}

	// Draw P2 Curve Lines
	for _, shape := range p1p2.P2CurveShapes {
		line := new(svg.Line)
		layer.Lines = append(layer.Lines, line)
		line.Name = shape.Name
		line.X1 = plant2DDiagram.OriginX + shape.StartX
		line.Y1 = plant2DDiagram.OriginY - shape.StartY
		line.X2 = plant2DDiagram.OriginX + shape.EndX
		line.Y2 = plant2DDiagram.OriginY - shape.EndY
		line.Presentation.Stroke = "magenta"
		line.Presentation.StrokeWidth = 1.5
		line.Presentation.StrokeOpacity = 0.8
	}

	// Draw P1-P2 Pair Lines
	for _, shape := range p1p2.P1P2PairLineShapes {
		line := new(svg.Line)
		layer.Lines = append(layer.Lines, line)
		line.Name = shape.Name
		line.X1 = plant2DDiagram.OriginX + shape.StartX
		line.Y1 = plant2DDiagram.OriginY - shape.StartY
		line.X2 = plant2DDiagram.OriginX + shape.EndX
		line.Y2 = plant2DDiagram.OriginY - shape.EndY
		line.Presentation.Stroke = "magenta"
		line.Presentation.StrokeWidth = 1.0
		line.Presentation.StrokeOpacity = 0.8
	}
}

func (plant2DDiagram *Plant2DDiagram) drawPxShape(stager *Stager, layer *svg.Layer, plant *PlantAbstract, vase2DDiagram *Vase2DDiagram) {
	if vase2DDiagram == nil || vase2DDiagram.IsHiddenPxShape {
		return
	}
	if plant.VaseAbstract == nil || plant.VaseAbstract.PxShape == nil {
		return
	}

	px := plant.VaseAbstract.PxShape

	circLen := 0.0
	if plant.RhombusStuff != nil && plant.RhombusStuff.PlantCircumferenceShape != nil {
		circLen = plant.RhombusStuff.PlantCircumferenceShape.Length
	}

	_, dy, currentDX := ComputePartiallyGrowthCurveDY(plant)

	drawPxAt := func(nameSuffix string, pxX, pxY float64) {
		circle := new(svg.Circle)
		layer.Circles = append(layer.Circles, circle)
		circle.Name = px.Name + nameSuffix
		circle.CX = plant2DDiagram.OriginX + pxX
		circle.CY = plant2DDiagram.OriginY - pxY
		circle.Radius = 4.0
		circle.Presentation.Color = "magenta"
		circle.Presentation.FillOpacity = 1.0
		circle.Presentation.Stroke = "darkmagenta"
		circle.Presentation.StrokeWidth = 1.0
		circle.Presentation.StrokeOpacity = 1.0

		text := new(svg.Text)
		layer.Texts = append(layer.Texts, text)
		text.Name = px.Name + nameSuffix + "-Text"
		text.X = plant2DDiagram.OriginX + pxX - 14
		text.Y = plant2DDiagram.OriginY - pxY - 6
		text.Content = "Px"
		text.Presentation.Color = "magenta"
		text.Presentation.FillOpacity = 1.0
	}

	if !vase2DDiagram.IsHiddenPartiallyGrowthCurve2DRibbon {
		drawPxAt("-Partially", px.X, px.Y)
	}

	if !vase2DDiagram.IsHiddenShiftedLeftPartiallyGrowthCurve2DRibbon && circLen > 0 {
		drawPxAt("-ShiftedLeftPartially", px.X-circLen, px.Y)
	}

	if !vase2DDiagram.IsHiddenGrowthCurve2DRibbon {
		drawPxAt("-GrowthCurve2DRibbon", px.X-currentDX, px.Y-dy)
	}

	if !vase2DDiagram.IsHiddenShiftedLeftGrowthCurve2DRibbon && circLen > 0 {
		drawPxAt("-ShiftedLeft", px.X-currentDX-circLen, px.Y-dy)
	}

	if !vase2DDiagram.IsHiddenShiftedRightGrowthCurve2DRibbon && circLen > 0 {
		drawPxAt("-ShiftedRight", px.X-currentDX+circLen, px.Y-dy)
	}
}

func (plant2DDiagram *Plant2DDiagram) drawKeyHoleShape(stager *Stager, layer *svg.Layer, plant *PlantAbstract, vase2DDiagram *Vase2DDiagram) {
	if vase2DDiagram == nil || vase2DDiagram.IsHiddenKeyHoleShape {
		return
	}
	if plant.VaseAbstract == nil || plant.VaseAbstract.KeyHoleShape == nil {
		return
	}

	keyHole := plant.VaseAbstract.KeyHoleShape
	if plant.PlantType == Vase {
		keyHole.X = plant.VaseAbstract.OffsetKeyX
		keyHole.Y = plant.VaseAbstract.OffsetKeyY
		keyHole.Width = plant.VaseAbstract.WidthKey
		keyHole.Height = plant.VaseAbstract.HeightKey
	}

	vThickness := 0.0
	widthKey := 0.0
	heightKey := 0.0
	offsetKeyX := 0.0
	offsetKeyY := 0.0
	if plant.PlantType == Vase {
		vThickness = plant.VaseAbstract.RelativeVerticalThickness * plant.RhombusSideLength * plant2DDiagram.getZoom()
		widthKey = plant.VaseAbstract.WidthKey
		heightKey = plant.VaseAbstract.HeightKey
		offsetKeyX = plant.VaseAbstract.OffsetKeyX
		offsetKeyY = plant.VaseAbstract.OffsetKeyY
	}

	drawRect := func(name string, offsetX, offsetY float64, withLine bool) {
		rect := new(svg.Rect)
		layer.Rects = append(layer.Rects, rect)
		rect.Name = name
		rect.X = plant2DDiagram.OriginX + offsetX - widthKey/2.0
		rect.Y = plant2DDiagram.OriginY - offsetY - heightKey/2.0
		rect.Width = widthKey
		rect.Height = heightKey

		rect.Presentation.Stroke = "darkred"
		rect.Presentation.StrokeWidth = 1.5
		rect.Presentation.StrokeOpacity = 1.0
		rect.Presentation.Color = "pink"
		rect.Presentation.FillOpacity = 0.3

		if withLine && vThickness > 0 {
			leftLine := new(svg.Line)
			layer.Lines = append(layer.Lines, leftLine)
			leftLine.Name = name + "-LeftLine"
			leftLine.X1 = rect.X
			leftLine.Y1 = rect.Y + rect.Height
			leftLine.X2 = rect.X
			leftLine.Y2 = rect.Y + rect.Height + vThickness

			leftLine.Presentation.Stroke = "darkred"
			leftLine.Presentation.StrokeWidth = 1.0
			leftLine.Presentation.StrokeOpacity = 1.0
		}
	}

	circLen := 0.0
	if plant.RhombusStuff != nil && plant.RhombusStuff.PlantCircumferenceShape != nil {
		circLen = plant.RhombusStuff.PlantCircumferenceShape.Length
	}

	_, dy, currentDX := ComputePartiallyGrowthCurveDY(plant)

	if !vase2DDiagram.IsHiddenPartiallyGrowthCurve2DRibbon {
		drawRect(keyHole.Name+"-Partially", offsetKeyX+currentDX, offsetKeyY+dy, true)
	}

	if !vase2DDiagram.IsHiddenShiftedLeftPartiallyGrowthCurve2DRibbon && circLen > 0 {
		drawRect(keyHole.Name+"-ShiftedLeftPartially", offsetKeyX+currentDX-circLen, offsetKeyY+dy, false)
	}

	if !vase2DDiagram.IsHiddenGrowthCurve2DRibbon {
		drawRect(keyHole.Name, offsetKeyX, offsetKeyY, false)
	}

	if !vase2DDiagram.IsHiddenShiftedLeftGrowthCurve2DRibbon && circLen > 0 {
		drawRect(keyHole.Name+"-ShiftedLeft", offsetKeyX-circLen, offsetKeyY, false)
	}

	if !vase2DDiagram.IsHiddenShiftedRightGrowthCurve2DRibbon && circLen > 0 {
		drawRect(keyHole.Name+"-ShiftedRight", offsetKeyX+circLen, offsetKeyY, false)
	}
}

func (plant2DDiagram *Plant2DDiagram) drawChosenP1P2PairShape(stager *Stager, layer *svg.Layer, plant *PlantAbstract, vase2DDiagram *Vase2DDiagram) {
	if vase2DDiagram == nil || vase2DDiagram.IsHiddenChosenP1P2PairShape {
		return
	}
	if plant.VaseAbstract == nil || plant.VaseAbstract.ChosenP1P2PairShape == nil {
		return
	}

	chosen := plant.VaseAbstract.ChosenP1P2PairShape

	// Line connecting Chosen P1 and P2
	line := new(svg.Line)
	layer.Lines = append(layer.Lines, line)
	line.Name = chosen.Name + "-Line"
	line.X1 = plant2DDiagram.OriginX + chosen.P1X
	line.Y1 = plant2DDiagram.OriginY - chosen.P1Y
	line.X2 = plant2DDiagram.OriginX + chosen.P2X
	line.Y2 = plant2DDiagram.OriginY - chosen.P2Y
	line.Presentation.Stroke = "darkred"
	line.Presentation.StrokeWidth = 2.5
	line.Presentation.StrokeOpacity = 1.0

	// Line connecting Chosen P1 and Px
	lineP1Px := new(svg.Line)
	layer.Lines = append(layer.Lines, lineP1Px)
	lineP1Px.Name = chosen.Name + "-P1-Px-Line"
	lineP1Px.X1 = plant2DDiagram.OriginX + chosen.P1X
	lineP1Px.Y1 = plant2DDiagram.OriginY - chosen.P1Y
	lineP1Px.X2 = plant2DDiagram.OriginX + chosen.PxX
	lineP1Px.Y2 = plant2DDiagram.OriginY - chosen.PxY
	lineP1Px.Presentation.Stroke = "darkred"
	lineP1Px.Presentation.StrokeWidth = 1.5
	lineP1Px.Presentation.StrokeOpacity = 0.9

	// Line connecting Chosen P2 and Px
	lineP2Px := new(svg.Line)
	layer.Lines = append(layer.Lines, lineP2Px)
	lineP2Px.Name = chosen.Name + "-P2-Px-Line"
	lineP2Px.X1 = plant2DDiagram.OriginX + chosen.P2X
	lineP2Px.Y1 = plant2DDiagram.OriginY - chosen.P2Y
	lineP2Px.X2 = plant2DDiagram.OriginX + chosen.PxX
	lineP2Px.Y2 = plant2DDiagram.OriginY - chosen.PxY
	lineP2Px.Presentation.Stroke = "darkred"
	lineP2Px.Presentation.StrokeWidth = 1.5
	lineP2Px.Presentation.StrokeOpacity = 0.9

	// Chosen P1 Dot
	circleP1 := new(svg.Circle)
	layer.Circles = append(layer.Circles, circleP1)
	circleP1.Name = chosen.Name + "-P1-Dot"
	circleP1.CX = plant2DDiagram.OriginX + chosen.P1X
	circleP1.CY = plant2DDiagram.OriginY - chosen.P1Y
	circleP1.Radius = 4.0
	circleP1.Presentation.Color = "red"
	circleP1.Presentation.FillOpacity = 1.0
	circleP1.Presentation.Stroke = "darkred"
	circleP1.Presentation.StrokeWidth = 1.5
	circleP1.Presentation.StrokeOpacity = 1.0

	// Chosen P2 Dot
	circleP2 := new(svg.Circle)
	layer.Circles = append(layer.Circles, circleP2)
	circleP2.Name = chosen.Name + "-P2-Dot"
	circleP2.CX = plant2DDiagram.OriginX + chosen.P2X
	circleP2.CY = plant2DDiagram.OriginY - chosen.P2Y
	circleP2.Radius = 4.0
	circleP2.Presentation.Color = "red"
	circleP2.Presentation.FillOpacity = 1.0
	circleP2.Presentation.Stroke = "darkred"
	circleP2.Presentation.StrokeWidth = 1.5
	circleP2.Presentation.StrokeOpacity = 1.0

	// Text showing distances: P1 to Px, P2 to Px, and Sum
	textDist := new(svg.Text)
	layer.Texts = append(layer.Texts, textDist)
	textDist.Name = chosen.Name + "-Distances-Text"
	textDist.X = plant2DDiagram.OriginX + chosen.P1X - 30
	textDist.Y = plant2DDiagram.OriginY - chosen.P1Y + 18
	textDist.Content = fmt.Sprintf("P1-Px: %.2f  P2-Px: %.2f  Sum: %.2f", chosen.DistanceP1Px, chosen.DistanceP2Px, chosen.DistanceSum)
	textDist.Presentation.Color = "darkred"
	textDist.Presentation.FillOpacity = 1.0

	// Draw Partial Ellipse defined ONLY by P1 & P2 (independent of ratio)
	if plant.VaseAbstract != nil && plant.VaseAbstract.PartiallyGrowthCurve2DTrajectory != nil {
		traj := plant.VaseAbstract.PartiallyGrowthCurve2DTrajectory
		nShapes := len(traj.PartiallyGrowthCurve2DTrajectoryShapes)
		if nShapes > 0 {
			x1 := traj.PartiallyGrowthCurve2DTrajectoryShapes[0].StartX
			y1 := traj.PartiallyGrowthCurve2DTrajectoryShapes[0].StartY
			x3 := traj.PartiallyGrowthCurve2DTrajectoryShapes[nShapes-1].EndX
			y3 := traj.PartiallyGrowthCurve2DTrajectoryShapes[nShapes-1].EndY
			midIdx := nShapes / 2
			x2 := traj.PartiallyGrowthCurve2DTrajectoryShapes[midIdx].EndX
			y2 := traj.PartiallyGrowthCurve2DTrajectoryShapes[midIdx].EndY

			// 3-point circle fit (x1, y1), (x2, y2), (x3, y3)
			D := 2 * (x1*(y2-y3) + x2*(y3-y1) + x3*(y1-y2))
			if math.Abs(D) > 1e-6 {
				sq1 := x1*x1 + y1*y1
				sq2 := x2*x2 + y2*y2
				sq3 := x3*x3 + y3*y3

				cx := (sq1*(y2-y3) + sq2*(y3-y1) + sq3*(y1-y2)) / D
				cy := (sq1*(x3-x2) + sq2*(x1-x3) + sq3*(x2-x1)) / D
				radius := math.Hypot(x1-cx, y1-cy)

				if radius > 0 {
					dxChord := x3 - x1
					dyChord := y3 - y1
					chordLen := math.Hypot(dxChord, dyChord)
					if chordLen > 0 {
						ux := dxChord / chordLen
						uy := dyChord / chordLen

						mx := (x1 + x3) / 2.0
						my := (y1 + y3) / 2.0

						vxRaw := -uy
						vyRaw := ux
						dotMid := (x2-mx)*vxRaw + (y2-my)*vyRaw
						if dotMid < 0 {
							vxRaw = -vxRaw
							vyRaw = -vyRaw
						}
						vx := vxRaw
						vy := vyRaw

						R1 := math.Abs((mx-cx)*vx + (my-cy)*vy)
						R2 := radius - R1

						if R1 >= 0 && R2 > 0 {
							refSteps := 10
							chosenK := 0
							if plant.PlantType == Vase {
								refSteps = plant.VaseAbstract.NbStepP1P2
								chosenK = plant.VaseAbstract.ChosenStep
							}
							if refSteps <= 0 {
								refSteps = 10
							}
							if chosenK < 0 {
								chosenK = 0
							}
							if chosenK > refSteps {
								chosenK = refSteps
							}

							yValChosen := float64(chosenK) * R1 / float64(refSteps)
							yLineChosen := -yValChosen

							cVal := math.Sqrt(2 * (R1 + yLineChosen) * (R2 - yLineChosen))
							bVal := R2 - yLineChosen
							aVal := math.Sqrt(bVal*bVal + cVal*cVal)

							ellipseCenterX := mx - yValChosen*vx
							ellipseCenterY := my - yValChosen*vy

							sinCutoff := yValChosen / bVal
							if sinCutoff < 0.0 {
								sinCutoff = 0.0
							}
							if sinCutoff > 1.0 {
								sinCutoff = 1.0
							}
							tMin := math.Asin(sinCutoff)
							tMax := math.Pi - tMin

							numEllipseSteps := 80
							ellipsePtsX := make([]float64, numEllipseSteps+1)
							ellipsePtsY := make([]float64, numEllipseSteps+1)

							for i := 0; i <= numEllipseSteps; i++ {
								t := tMin + float64(i)*(tMax-tMin)/float64(numEllipseSteps)
								ellipsePtsX[i] = plant2DDiagram.OriginX + (ellipseCenterX + aVal*math.Cos(t)*ux + bVal*math.Sin(t)*vx)
								ellipsePtsY[i] = plant2DDiagram.OriginY - (ellipseCenterY + aVal*math.Cos(t)*uy + bVal*math.Sin(t)*vy)
							}

							for i := 0; i < numEllipseSteps; i++ {
								lineE := new(svg.Line)
								layer.Lines = append(layer.Lines, lineE)
								lineE.Name = fmt.Sprintf("%s-Partial-Ellipse-Seg-%d", chosen.Name, i)
								lineE.X1 = ellipsePtsX[i]
								lineE.Y1 = ellipsePtsY[i]
								lineE.X2 = ellipsePtsX[i+1]
								lineE.Y2 = ellipsePtsY[i+1]
								lineE.Presentation.Stroke = "darkorange"
								lineE.Presentation.StrokeWidth = 2.0
								lineE.Presentation.StrokeOpacity = 1.0
							}
						}
					}
				}
			}
		}
	}
}
