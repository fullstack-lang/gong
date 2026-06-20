package models

import (
	"fmt"

	"github.com/fullstack-lang/gong/lib/strutils"
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

	stager.drawSystemShape(diagram, layer)

	rectOfOwningSystem := diagram.map_System_Rect[diagram.owningSystem]
	if rectOfOwningSystem != nil {
		stager.drawPartShapes(diagram, layer, rectOfOwningSystem)
	}
	stager.drawLinkShapes(diagram, layer)

	return svgObject
}

func (stager *Stager) drawPartShapes(diagram *DiagramStructure, layer *svg.Layer, rectOfOwningSystem *svg.Rect) {
	root := stager.GetRootLibrary()

	diagram.map_Part_Rect = make(map[*Part]*svg.Rect)
	diagram.map_SvgRect_PartShape = make(map[*svg.Rect]*PartShape)

	horizontalMargin := 10.0
	verticalTopMargin := 50.0
	verticalTopMarginForTitle := 60.0
	verticalBottomMargin := 10.0

	partsWidth := rectOfOwningSystem.Width - 2*horizontalMargin

	var totalWeight float64
	for _, pShape := range diagram.Part_Shapes {
		weight := pShape.WidthWeight
		if weight == 0 {
			weight = 1.0
		}
		totalWeight += weight
	}

	currentWeight := 0.0

	for idx, partShape := range diagram.Part_Shapes {
		shapeWeight := partShape.WidthWeight
		if shapeWeight == 0 {
			shapeWeight = 1.0
		}
		partWidth := 0.0
		if totalWeight > 0 {
			partWidth = shapeWeight * (partsWidth / totalWeight)
		}

		if partShape.IsHidden {
			currentWeight += shapeWeight
			continue
		}

		rect := new(svg.Rect)
		layer.Rects = append(layer.Rects, rect)
		diagram.map_SvgRect_PartShape[rect] = partShape

		rect.Name = partShape.GetName()
		rect.Stroke = "#E0E0E0"
		rect.StrokeWidth = 1
		rect.StrokeOpacity = 1

		rect.Color = "#FFFFFF"
		rect.FillOpacity = 1.0

		// rect cannot move
		rect.CanMoveHorizontaly = true
		rect.CanMoveVerticaly = true
		rect.CanHaveBottomHandle = false
		rect.CanHaveLeftHandle = false
		rect.CanHaveTopHandle = false

		// all but the last part have right handle
		rect.CanHaveRightHandle = func() bool {
			if idx == len(diagram.Part_Shapes)-1 {
				return false
			}
			return true
		}()

		// visuals
		rect.RX = 0
		rect.StrokeWidth = 1

		if totalWeight > 0 {
			rect.X = rectOfOwningSystem.X + horizontalMargin + currentWeight*(partsWidth/totalWeight)
		} else {
			rect.X = rectOfOwningSystem.X + horizontalMargin
		}
		rect.Width = partWidth

		rect.Y = rectOfOwningSystem.Y + verticalTopMargin + verticalTopMarginForTitle
		rect.Height = rectOfOwningSystem.Height - verticalTopMargin - verticalBottomMargin - verticalTopMarginForTitle

		// make the part shape peer of the system shape
		rect.Peers = append(rect.Peers, rectOfOwningSystem)
		rectOfOwningSystem.Peers = append(rectOfOwningSystem.Peers, rect)

		// override default behavior, we need to commit when the rect is moved
		rect.OnSelect = func() {
			stager.probeForm.FillUpFormFromGongstruct(partShape.Part, GetPointerToGongstructName[*Part]())
		}
		rect.OnMove = func(x, y float64) {}
		rect.OnResize = func(x, y, width, height float64) {
			if width != partWidth {
				othersWeight := totalWeight - shapeWeight
				if othersWeight > 0 {
					boundedWidth := width
					if boundedWidth > partsWidth-10 {
						boundedWidth = partsWidth - 10
					}
					if boundedWidth < 10 {
						boundedWidth = 10
					}
					newWeight := (boundedWidth * othersWeight) / (partsWidth - boundedWidth)

					partShape.WidthWeight = newWeight
					stager.stage.Commit()
				}
			}
		}

		if partShape.Part != nil {
			diagram.map_Part_Rect[partShape.Part] = rect
		}

		title := new(svg.RectAnchoredText)
		title.Name = partShape.GetAbstractElement().GetName()

		content := partShape.GetAbstractElement().GetName()
		if diagram.GetIsShowPrefix() {
			content = partShape.GetAbstractElement().GetComputedPrefix() + " " + content
		}

		if rect.Width > 0 {
			content = strutils.WrapStringPreservingNewlines(content, int(rect.Width/root.NbPixPerCharacter))
		}
		title.Content = content
		title.StrokeWidth = 0
		title.StrokeOpacity = 1
		title.Color = "#333333"
		title.FillOpacity = 1
		title.FontSize = "16px"
		title.FontWeight = "500"
		title.X_Offset = 0
		title.Y_Offset = -verticalTopMarginForTitle / 2.0
		title.RectAnchorType = svg.RECT_TOP
		title.TextAnchorType = svg.TEXT_ANCHOR_CENTER

		rect.RectAnchoredTexts = append(rect.RectAnchoredTexts, title)

		titleBox := &svg.RectAnchoredRect{
			Name: partShape.GetAbstractElement().GetName(),
			Presentation: svg.Presentation{
				Stroke:        "#E0E0E0",
				StrokeWidth:   1,
				StrokeOpacity: 1,
			},
			X_Offset:       0,
			Y_Offset:       -verticalTopMarginForTitle,
			Height:         verticalTopMarginForTitle,
			Width:          rect.Width,
			RectAnchorType: svg.RECT_TOP_LEFT,
		}
		rect.RectAnchoredRects = append(rect.RectAnchoredRects, titleBox)

		currentWeight += shapeWeight
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


func (stager *Stager) drawSystemShape(diagram *DiagramStructure, layer *svg.Layer) {
	diagram.map_System_Rect = make(map[*System]*svg.Rect)
	for _, systemShape := range diagram.System_Shapes {

		rect := svgRect(
			stager,
			diagram,
			systemShape,
			layer)
		rect.RX = 3

		rect.Color = "#F8F9FA"
		rect.FillOpacity = 1.0
		rect.Stroke = "#E0E0E0"
		rect.StrokeWidth = 1.5

		if len(rect.RectAnchoredTexts) > 0 {
			title := rect.RectAnchoredTexts[0]
			title.FontWeight = "500"
			title.FontSize = "18px"
		}

		rect.OnSelect = func() {
			stager.probeForm.FillUpFormFromGongstruct(systemShape.System, GetPointerToGongstructName[*System]())
		}
		rect.OnMove = func(x, y float64) {
			systemShape.SetX(x)
			systemShape.SetY(y)
			stager.stage.CommitWithSuspendedCallbacks()
		}
		rect.OnResize = func(x, y, width, height float64) {
			systemShape.SetX(x)
			systemShape.SetY(y)
			systemShape.SetWidth(width)
			systemShape.SetHeight(height)
			stager.stage.Commit()
		}
		diagram.map_System_Rect[systemShape.System] = rect
	}
}
