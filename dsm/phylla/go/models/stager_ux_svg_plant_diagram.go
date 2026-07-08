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

	verticalAxisTopHandle := new(svg.Rect)
	verticalAxisTopHandle.Name = "Vertical axis bottom handle"
	layer.Rects = append(layer.Rects, verticalAxisTopHandle)
	verticalAxisTopHandle.X = plantDiagram.OriginX - AxisHandleBorderLength/2.0
	verticalAxisTopHandle.Y = plantDiagram.OriginY - plantDiagram.AxesShape.LengthY
	verticalAxisTopHandle.Width = AxisHandleBorderLength
	verticalAxisTopHandle.Height = AxisHandleBorderLength
	verticalAxisTopHandle.CanMoveVerticaly = true
	verticalAxisTopHandle.OnMove = func(x, y float64) {
		plantDiagram.AxesShape.LengthY = plantDiagram.OriginY - y
		stager.stage.CommitWithSuspendedCallbacks()
	}

	verticalAxisTopHandle.Presentation.Stroke = "blue"
	verticalAxisTopHandle.Presentation.StrokeWidth = 1
	verticalAxisTopHandle.Presentation.StrokeOpacity = 1

	verticalAxisBottomHandle := new(svg.Rect)
	verticalAxisBottomHandle.Name = "Vertical axis top handle"
	layer.Rects = append(layer.Rects, verticalAxisBottomHandle)

	verticalAxisBottomHandle.X = plantDiagram.OriginX - AxisHandleBorderLength/2.0
	verticalAxisBottomHandle.Y = plantDiagram.OriginY - AxisHandleBorderLength
	verticalAxisBottomHandle.Width = AxisHandleBorderLength
	verticalAxisBottomHandle.Height = AxisHandleBorderLength
	verticalAxisBottomHandle.CanMoveVerticaly = true
	verticalAxisBottomHandle.CanMoveHorizontaly = true
	verticalAxisBottomHandle.OnMove = func(x, y float64) {
		plantDiagram.OriginX = x
		plantDiagram.OriginY = y
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

}
