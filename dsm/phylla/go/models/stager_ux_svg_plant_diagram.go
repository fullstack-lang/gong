package models

import (
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

	if plantDiagram == nil {
		stager.svgStage.Commit()
		return
	}
	svgObject := stager.generateSvgObject(plantDiagram)

	svg.StageBranch(stager.svgStage, svgObject)
	stager.svgObject = svgObject

	stager.svgStage.Commit()
}

func (stager *Stager) generateSvgObject(plantDiagram *PlantDiagram) (svg_ *svg.SVG) {

	svg_ = new(svg.SVG)
	svg_.Name = "Plant Diagram" // or any name, if name is an attribute.
	svg_.IsEditable = true

	layer := &svg.Layer{Name: `Axis Shape Layer`}
	svg_.Layers = append(svg_.Layers, layer)

	// creation of 2 transparant rects, one at each ends of the vertical
	plantDiagram.draw(stager, layer)

	return
}

const AxisHandleBorderLength = 25

func (plantDiagram *PlantDiagram) draw(stager *Stager, layer *svg.Layer) {

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

	horizontalAxisLine.StartAnchorType = svg.ANCHOR_CENTER
	horizontalAxisLine.EndAnchorType = svg.ANCHOR_CENTER

	// right and top handle move with vertical bottom handle
	verticalAxisBottomHandle.Peers = append(verticalAxisBottomHandle.Peers,
		verticalAxisTopHandle,
		horizontalAxisRightHandle)
}
