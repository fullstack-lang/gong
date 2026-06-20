package models

import (
	"fmt"

	svg "github.com/fullstack-lang/gong/lib/svg/go/models"
)

func (stager *Stager) svg() {
	svgStage := stager.svgStage

	// reset the stage
	svgStage.Reset()

	var diagram *DiagramStructure
	{
		for diagram_ := range *GetGongstructInstancesSet[DiagramStructure](stager.stage) {
			if diagram_.IsChecked {
				diagram = diagram_
			}
		}
	}

	if diagram == nil {
		svgStage.Commit()
		return
	}

	svgObject := stager.generateSvgObject(diagram)

	svg.StageBranch(svgStage, svgObject)

	svgStage.Commit()
}

func (stager *Stager) generateSvgObject(diagram *DiagramStructure) *svg.SVG {
	svgObject := (&svg.SVG{Name: "SVG"})
	stager.diagram = diagram

	svgObject.OverrideWidth = true
	svgObject.OverriddenWidth = diagram.Width
	svgObject.OverrideHeight = true
	svgObject.OverriddenHeight = diagram.Height

	svgObject.Name = diagram.Name
	svgObject.IsEditable = diagram.IsEditable()

	layer := (&svg.Layer{Name: "Default"})
	svgObject.Layers = append(svgObject.Layers, layer)

	stager.drawPartShapes(diagram, layer)
	stager.drawLinkShapes(diagram, layer)

	return svgObject
}

func (stager *Stager) drawPartShapes(diagram *DiagramStructure, layer *svg.Layer) {
	diagram.map_Part_Rect = make(map[*Part]*svg.Rect)
	diagram.map_SvgRect_PartShape = make(map[*svg.Rect]*PartShape)

	for _, partShape := range diagram.Part_Shapes {
		if partShape.IsHidden {
			continue
		}

		rect := new(svg.Rect)
		layer.Rects = append(layer.Rects, rect)

		rect.Name = partShape.Name
		rect.X = partShape.X
		rect.Y = partShape.Y
		rect.Width = partShape.Width
		if rect.Width == 0 {
			rect.Width = 200 // default width
		}
		rect.Height = partShape.Height
		if rect.Height == 0 {
			rect.Height = 100 // default height
		}

		rect.Color = "white"
		rect.FillOpacity = 100
		rect.Stroke = "black"
		rect.StrokeWidth = 2
		rect.RX = 5

		// Setup interactivity
		rect.IsSelectable = true
		rect.CanHaveLeftHandle = true
		rect.CanHaveRightHandle = true
		rect.CanHaveTopHandle = true
		rect.CanHaveBottomHandle = true

		rect.OnSelect = func() {
			if partShape.Part != nil {
				stager.probeForm.FillUpFormFromGongstruct(partShape.Part, GetPointerToGongstructName[*Part]())
			}
		}
		rect.OnMove = func(x, y float64) {
			partShape.SetX(x)
			partShape.SetY(y)
			stager.stage.CommitWithSuspendedCallbacks()
		}
		rect.OnResize = func(x, y, width, height float64) {
			partShape.SetX(x)
			partShape.SetY(y)
			partShape.SetWidth(width)
			partShape.SetHeight(height)
			stager.stage.Commit()
		}

		if partShape.Part != nil {
			diagram.map_Part_Rect[partShape.Part] = rect
			diagram.map_SvgRect_PartShape[rect] = partShape
		}
	}
}

func (stager *Stager) drawLinkShapes(diagram *DiagramStructure, layer *svg.Layer) {
	for _, linkShape := range diagram.Link_Shapes {
		if linkShape.IsHidden {
			continue
		}
		if linkShape.Link == nil || linkShape.Link.Source == nil || linkShape.Link.Target == nil {
			continue
		}

		sourceRect, ok1 := diagram.map_Part_Rect[linkShape.Link.Source]
		targetRect, ok2 := diagram.map_Part_Rect[linkShape.Link.Target]

		if ok1 && ok2 {
			svgLink := new(svg.Link)
			layer.Links = append(layer.Links, svgLink)

			svgLink.Name = fmt.Sprintf("Link from %s to %s", linkShape.Link.Source.Name, linkShape.Link.Target.Name)
			
			svgLink.Start = sourceRect
			svgLink.End = targetRect

			svgLink.StartOrientation = svg.OrientationType(linkShape.StartOrientation)
			svgLink.EndOrientation = svg.OrientationType(linkShape.EndOrientation)

			svgLink.StartRatio = linkShape.StartRatio
			if svgLink.StartRatio == 0 {
				svgLink.StartRatio = 0.5
			}
			svgLink.EndRatio = linkShape.EndRatio
			if svgLink.EndRatio == 0 {
				svgLink.EndRatio = 0.5
			}
			svgLink.CornerOffsetRatio = linkShape.CornerOffsetRatio
			if svgLink.CornerOffsetRatio == 0 {
				svgLink.CornerOffsetRatio = 1.0
			}

			svgLink.Color = "black"
			svgLink.StrokeWidth = 2
			svgLink.Type = svg.LINK_TYPE_FLOATING_ORTHOGONAL
			svgLink.HasEndArrow = true
			svgLink.EndArrowSize = 10

			// Text for the link
			text := new(svg.LinkAnchoredText)
			text.Name = linkShape.Link.Name
			text.Content = linkShape.Link.Name
			text.X_Offset = 10
			text.Y_Offset = -10
			text.Color = "black"
			
			svgLink.TextAtArrowEnd = append(svgLink.TextAtArrowEnd, text)
		}
	}
}

