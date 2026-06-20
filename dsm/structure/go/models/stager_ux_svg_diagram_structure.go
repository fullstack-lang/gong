package models

import (
	"fmt"
	"log"
	"strings"

	svg "github.com/fullstack-lang/gong/lib/svg/go/models"
	"github.com/fullstack-lang/gong/lib/strutils"
)

func (stager *Stager) svg() {
	log.Println("svg")
	stager.systemDiagramSvgStage.Reset()

	var diagramStructure *DiagramStructure
	{
		for diagramsystem_ := range *GetGongstructInstancesSet[DiagramStructure](stager.stage) {
			if diagramsystem_.IsChecked {
				diagramStructure = diagramsystem_
			}
		}
	}

	if diagramStructure == nil {
		stager.systemDiagramSvgStage.Commit()
		return
	}
	svgObject := stager.generateSvgObject(diagramStructure)

	svg.StageBranch(stager.systemDiagramSvgStage, svgObject)
	stager.svgObjectDiagramStructure = svgObject
	stager.svgObjectDiagramStructure.OnUpdate = stager.onUpdateSVG

	stager.systemDiagramSvgStage.Commit()
}

// generateSvgObject creates and returns a new svg.SVG object representing the given diagram.
// It maps all visible domain shapes (Systems, Ports, Notes, Resources) and their associations
// to SVG elements (Rects, Links, Paths) on a single layer. It also populates the diagram's
// internal maps to link abstract elements with their visual SVG counterparts.
func (stager *Stager) generateSvgObject(diagramStructure *DiagramStructure) *svg.SVG {
	svgObject := (&svg.SVG{Name: `SVG`})
	stager.diagramStructure = diagramStructure

	svgObject.OverrideWidth = true
	svgObject.OverriddenWidth = diagramStructure.Width
	svgObject.OverrideHeight = true
	svgObject.OverriddenHeight = diagramStructure.Height

	diagramStructure.map_SvgRect_PortShape = make(map[*svg.Rect]*PortShape)
	diagramStructure.map_SvgRect_ExternalPartShape = map[*svg.Rect]*ExternalPartShape{}
	diagramStructure.map_SvgRect_Part = map[*svg.Rect]*Part{}

	// // to implement association between abstract elements by mouse drag
	// svgImpl := &svgProxy{
	// 	stager:  stager,
	// 	svg_:
	// 	diagramStructure: diagramStructure,
	// }
	// Impl = svgImpl

	svgObject.Name = diagramStructure.Name
	svgObject.IsEditable = diagramStructure.IsEditable()

	layer := (&svg.Layer{Name: "Layer 1"})
	svgObject.Layers = append(svgObject.Layers, layer)

	stager.drawSystemShapes(diagramStructure, layer)

	rectOfOwningSystem := diagramStructure.map_System_Rect[diagramStructure.owningSystem]
	if rectOfOwningSystem != nil {
		stager.drawPartShapes(diagramStructure, layer, rectOfOwningSystem)
	}
	stager.drawExternalPartShapes(diagramStructure, layer)
	stager.drawPortShapes(diagramStructure, layer)
	stager.drawControlFlowShapes(diagramStructure, layer)
	stager.drawDataFlowShapes(diagramStructure, layer)
	map_Note_Rect := stager.drawNoteShapes(diagramStructure, layer)
	stager.drawNotePortShapes(diagramStructure, layer, map_Note_Rect)

	return svgObject
}

func (stager *Stager) drawSystemShapes(diagramStructure *DiagramStructure, layer *svg.Layer) {
	diagramStructure.map_System_Rect = make(map[*System]*svg.Rect)
	for _, systemShape := range diagramStructure.System_Shapes {
		if systemShape.IsHidden {
			continue
		}

		rect := svgRect(
			stager,
			diagramStructure,
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

		// override default behavior, we need to commit when the rect is moved
		// we need to update the position of the parts shapes and port shapes that are within the system
		rect.OnSelect = func() {
			stager.probeForm.FillUpFormFromGongstruct(systemShape.System, GetPointerToGongstructName[*System]())
		}
		rect.OnMove = func(x, y float64) {
			systemShape.SetX(x)
			systemShape.SetY(y)

			// no need to update the position of the part shapes and port shapes that are within the system
			// because they are peers of the system shape, so they will move together
			stager.stage.CommitWithSuspendedCallbacks()
		}
		rect.OnResize = func(x, y, width, height float64) {
			systemShape.SetX(x)
			systemShape.SetY(y)
			systemShape.SetWidth(width)
			systemShape.SetHeight(height)

			// when heigth or width is updated, we need to update the position/size of the parts shapes and port shapes that are within the system
			// this is performed in the semantic pass of the commit
			stager.stage.Commit()
		}
		diagramStructure.map_System_Rect[systemShape.System] = rect
	}
}

func (stager *Stager) drawPartShapes(diagramStructure *DiagramStructure, layer *svg.Layer, rectOfOwningSystem *svg.Rect) {
	root := stager.GetRootLibrary()

	diagramStructure.map_Part_Rect = make(map[*Part]*svg.Rect)
	horizontalMargin := 10.0
	verticalTopMargin := 50.0
	verticalTopMarginForTitle := 60.0
	verticalBottomMargin := 10.0

	partsWidth := rectOfOwningSystem.Width - 2*horizontalMargin

	var totalWeight float64
	for _, pShape := range diagramStructure.Part_Shapes {
		weight := pShape.WidthWeight
		if weight == 0 {
			weight = 1.0
		}
		totalWeight += weight
	}

	currentWeight := 0.0

	for idx, partShape := range diagramStructure.Part_Shapes {
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
		diagramStructure.map_SvgRect_Part[rect] = partShape.Part

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
			if idx == len(diagramStructure.Part_Shapes)-1 {
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

		if partShape.Part.IsSystemResource {
			rect.Y = rectOfOwningSystem.Y + verticalTopMargin
			rect.Height = rectOfOwningSystem.Height - verticalTopMargin - verticalBottomMargin
		}

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

		diagramStructure.map_Part_Rect[partShape.Part] = rect

		if !partShape.Part.IsSystemResource {
			title := new(svg.RectAnchoredText)
			title.Name = partShape.GetAbstractElement().GetName()

			content := partShape.GetAbstractElement().GetName()
			if diagramStructure.GetIsShowPrefix() {
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
		}
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

		boxHeigth, anchoredRect := stager.drawAllocatedSystemesAndResources(partShape.Part, diagramStructure, rect, svg.RECT_TOP)
		_ = boxHeigth
		_ = anchoredRect

		if partShape.Part.IsSystemResource {
			if boxHeigth > 0 {
				titleBox.Y_Offset = 0
				titleBox.Height = boxHeigth
			} else {
				titleBox.Height = 0
				titleBox.StrokeWidth = 0
			}
		}

		currentWeight += shapeWeight
	}

	// make all part shapes peers together, so that if one of them move, all the others move as well
	for _, partShape := range diagramStructure.Part_Shapes {
		if partShape.IsHidden {
			continue
		}
		rect := diagramStructure.map_Part_Rect[partShape.Part]
		if rect == nil {
			continue
		}
		for _, otherPartShape := range diagramStructure.Part_Shapes {
			if partShape == otherPartShape || otherPartShape.IsHidden {
				continue
			}
			otherRect := diagramStructure.map_Part_Rect[otherPartShape.Part]
			if otherRect == nil {
				continue
			}
			rect.Peers = append(rect.Peers, otherRect)
			otherRect.Peers = append(otherRect.Peers, rect)
		}
	}
}

func (stager *Stager) drawExternalPartShapes(diagramStructure *DiagramStructure, layer *svg.Layer) {
	root := stager.GetRootLibrary()

	diagramStructure.map_ExternalPart_Rect = make(map[*Part]*svg.Rect)
	for _, externalPartShape := range diagramStructure.ExternalPart_Shapes {
		if externalPartShape.IsHidden {
			continue
		}

		rect := new(svg.Rect)
		layer.Rects = append(layer.Rects, rect)

		rect.Name = externalPartShape.GetName()
		rect.X = externalPartShape.GetX()
		rect.Y = externalPartShape.GetY()
		rect.Width = externalPartShape.GetWidth()
		rect.Height = externalPartShape.GetHeight()

		rect.CanMoveHorizontaly = true
		rect.CanMoveVerticaly = true
		rect.CanHaveBottomHandle = true
		rect.CanHaveLeftHandle = true
		rect.CanHaveRightHandle = true
		rect.CanHaveTopHandle = true

		rect.RX = 0
		rect.StrokeWidth = 1.0
		rect.StrokeOpacity = 1.0

		rect.StrokeWidth = 1
		rect.StrokeOpacity = 1

		rect.FillOpacity = 1.0

		rect.Color = "#F8F9FA"
		rect.FillOpacity = 1.0
		rect.Stroke = "#E0E0E0"

		if !externalPartShape.Part.IsSystemResource {
			title := new(svg.RectAnchoredText)
			title.Name = externalPartShape.GetAbstractElement().GetName()

			content := externalPartShape.GetAbstractElement().GetName()
			if diagramStructure.GetIsShowPrefix() {
				content = externalPartShape.GetAbstractElement().GetComputedPrefix() + " " + content
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
			title.Y_Offset = 0
			title.RectAnchorType = svg.RECT_CENTER
			title.TextAnchorType = svg.TEXT_ANCHOR_CENTER

			rect.RectAnchoredTexts = append(rect.RectAnchoredTexts, title)
		}

		rect.OnSelect = onSelectRectElement(stager, externalPartShape.Part)
		rect.OnMove = onMoveRectElement(stager, externalPartShape, false)
		rect.OnResize = onResizeRectElement(stager, externalPartShape)

		externalPartWidth := 5.0
		if externalPartShape.TailHeigth == 0 {
			externalPartShape.TailHeigth = 100.0
		}

		anchorType := svg.RECT_BOTTOM
		if externalPartShape.Part.IsSystemResource {
			anchorType = svg.RECT_TOP
		}

		boxHeight, rectAnchoredRect := stager.drawAllocatedSystemesAndResources(
			externalPartShape.Part,
			diagramStructure,
			rect,
			anchorType)
		_ = rectAnchoredRect

		if externalPartShape.Part.IsSystemResource && boxHeight > 0 {
			rect.Height = boxHeight
			boxHeight = 0
		}

		// have the system rect be an obstacle to the rect
		systemRect := diagramStructure.map_System_Rect[diagramStructure.owningSystem]
		rect.Obstacles = append(rect.Obstacles, systemRect)
		systemRect.Obstacles = append(systemRect.Obstacles, rect)

		tailRect := &svg.Rect{
			Name:               "Tail" + rect.GetName(),
			CanMoveHorizontaly: true,
			CanMoveVerticaly:   true,
			Presentation: svg.Presentation{
				Stroke:          "#9E9E9E",
				StrokeWidth:     1.5,
				StrokeOpacity:   1,
				StrokeDashArray: "5 5",
				Color:           "transparent",
				FillOpacity:     0.0,
			},
			Width:               externalPartWidth,
			Height:              externalPartShape.TailHeigth - boxHeight,
			X:                   rect.X + (rect.Width-externalPartWidth)/2.0,
			Y:                   rect.Y + rect.Height + boxHeight,
			CanHaveBottomHandle: true,
		}
		diagramStructure.map_ExternalPart_Rect[externalPartShape.Part] = tailRect

		// tail rect and rect are moved togethere, therefore they are peers
		rect.Peers = append(rect.Peers, tailRect)
		tailRect.Peers = append(tailRect.Peers, rect)

		// this is the tail rect that is used for drawing links between external part and port, so we need to keep a reference of it in the diagram system
		diagramStructure.map_SvgRect_ExternalPartShape[tailRect] = externalPartShape
		// if the tailRect bottom handle is used, the heigth is updated
		tailRect.OnSelect = func() {}
		tailRect.OnMove = func(x, y float64) {}
		tailRect.OnResize = func(x, y, width, height float64) {
			diffSize := externalPartShape.TailHeigth != height+boxHeight

			if diffSize {
				externalPartShape.TailHeigth = height + boxHeight
				stager.stage.CommitWithSuspendedCallbacks()
			}
		}
		layer.Rects = append(layer.Rects, tailRect)
	}
}

func (stager *Stager) drawPortShapes(diagramStructure *DiagramStructure, layer *svg.Layer) {
	rm := GetSliceOfPointersReverseMap[Part, Port](GetAssociationName[Part]().Ports[0].Name, stager.stage)

	diagramStructure.map_Port_Rect = make(map[*Port]*svg.Rect)
	for _, portShape := range diagramStructure.Port_Shapes {
		if portShape.IsHidden {
			continue
		}

		port := portShape.Port
		parts := rm[port]
		if len(parts) == 0 {
			continue
		}

		part := parts[0]
		partRect := diagramStructure.map_Part_Rect[part]
		if partRect == nil {
			continue
		}

		systemRect := diagramStructure.map_System_Rect[diagramStructure.owningSystem]
		if systemRect == nil {
			continue
		}

		rect := svgRect(
			stager,
			diagramStructure,
			portShape,
			layer)
		diagramStructure.map_SvgRect_PortShape[rect] = portShape

		rect.URLPath = "../../../References/Ports/" + port.GetReferencePath() + "/index.html"
		rect.URLTarget = svg.LINK_TARGET_BLANK

		rect.RectAnchoredTexts[0].URLPath = rect.URLPath
		rect.RectAnchoredTexts[0].URLTarget = svg.LINK_TARGET_BLANK

		// make the rect of the port move with alls part rect and the system rect
		// not the opposite !
		systemRect.Peers = append(systemRect.Peers, rect)

		// we range over the non hidden parts and have the port shape appended as peer
		for _, pShape := range diagramStructure.Part_Shapes {
			if pShape.IsHidden {
				continue
			}
			partRect := diagramStructure.map_Part_Rect[pShape.Part]
			if partRect == nil {
				continue
			}
			partRect.Peers = append(partRect.Peers, rect)
		}

		rect.EnclosingRect = partRect

		rect.Color = "#E3F2FD"
		rect.FillOpacity = 1.0
		rect.Stroke = "#90CAF9"
		rect.StrokeWidth = 1.5
		rect.RX = 5.0

		if len(rect.RectAnchoredTexts) > 0 {
			rect.RectAnchoredTexts[0].FontFamily = "sans-serif"
			rect.RectAnchoredTexts[0].Color = "#333333"
		}

		// pick up the title of the rect
		stateTitleText := rect.RectAnchoredTexts[0]
		smallRadius := 10.0
		if port.IsStartPort {
			stateTitleText.TextAnchorType = svg.TEXT_ANCHOR_START
			stateTitleText.RectAnchorType = svg.RECT_TOP_LEFT
			stateTitleText.DominantBaseline = svg.DominantBaselineCentral
			stateTitleText.WhiteSpace = svg.WhiteSpaceEnumPre
			stateTitleText.X_Offset = 0
			stateTitleText.Y_Offset = 0

			circle := new(svg.RectAnchoredPath)
			circle.Stroke = "#90CAF9"
			circle.StrokeWidth = 2
			circle.StrokeOpacity = 1

			circle.Color = "#E3F2FD"
			circle.FillOpacity = 1.0

			// force size
			rect.CanHaveBottomHandle = false
			rect.CanHaveTopHandle = false

			// we allow resizing for the sake of the text width
			if rect.Width < 2*smallRadius {
				rect.Width = 2 * smallRadius
			}
			rect.Height = 2 * smallRadius

			circle.Definition = fmt.Sprintf("M %f 0 A %f %f 0 0 1 %f %f A %f %f 0 0 1 %f 0 Z",
				smallRadius, smallRadius, smallRadius, smallRadius, 2*smallRadius, smallRadius, smallRadius, smallRadius)
			circle.X_Offset = -smallRadius
			circle.Y_Offset = -smallRadius
			circle.RectAnchorType = svg.RECT_RIGHT
			rect.RectAnchoredPaths = append(rect.RectAnchoredPaths, circle)

			rect.StrokeOpacity = 0.0
			rect.FillOpacity = 0.0
		}

		bigRadius := 18.0
		if port.IsEndPort {
			stateTitleText.TextAnchorType = svg.TEXT_ANCHOR_START
			stateTitleText.RectAnchorType = svg.RECT_TOP_LEFT
			stateTitleText.DominantBaseline = svg.DominantBaselineCentral
			stateTitleText.WhiteSpace = svg.WhiteSpaceEnumPre
			stateTitleText.X_Offset = 0
			stateTitleText.Y_Offset = 0

			rect.CanHaveBottomHandle = false
			rect.CanHaveTopHandle = false
			if rect.Width < 2*bigRadius {
				rect.Width = 2 * bigRadius
			}
			rect.Height = 2 * bigRadius

			{
				circle := new(svg.RectAnchoredPath)

				circle.Stroke = "#90CAF9"
				circle.StrokeWidth = 2
				circle.StrokeOpacity = 1.0

				circle.Definition = fmt.Sprintf("M %f 0 A %f %f 0 0 1 %f %f A %f %f 0 0 1 %f 0 Z",
					bigRadius, bigRadius, bigRadius, bigRadius, 2*bigRadius, bigRadius, bigRadius, bigRadius)
				circle.X_Offset = -2 * bigRadius
				circle.Y_Offset = -bigRadius
				circle.RectAnchorType = svg.RECT_RIGHT
				rect.RectAnchoredPaths = append(rect.RectAnchoredPaths, circle)
			}

			{
				circle := new(svg.RectAnchoredPath)
				circle.Stroke = "#90CAF9"
				circle.StrokeWidth = 2
				circle.StrokeOpacity = 1

				circle.Color = "#90CAF9"
				circle.FillOpacity = 1.0

				circle.Definition = fmt.Sprintf("M %f 0 A %f %f 0 0 1 %f %f A %f %f 0 0 1 %f 0 Z",
					smallRadius, smallRadius, smallRadius, smallRadius, 2*smallRadius, smallRadius, smallRadius, smallRadius)
				circle.X_Offset = -smallRadius - bigRadius
				circle.Y_Offset = -smallRadius
				circle.RectAnchorType = svg.RECT_RIGHT
				rect.RectAnchoredPaths = append(rect.RectAnchoredPaths, circle)
			}

			rect.StrokeOpacity = 0.0
			rect.FillOpacity = 0.0

		}
		diagramStructure.map_Port_Rect[portShape.Port] = rect
	}
}

func (stager *Stager) drawControlFlowShapes(diagramStructure *DiagramStructure, layer *svg.Layer) {
	for _, controlFlowShape := range diagramStructure.ControlFlow_Shapes {
		if controlFlowShape.GetIsHidden() {
			continue
		}
		_ = controlFlowShape
		startPort := controlFlowShape.ControlFlow.Start
		endPort := controlFlowShape.ControlFlow.End

		if startPort == nil || endPort == nil {
			log.Panic("There should be a start port and a end port")
		}

		startRect := diagramStructure.map_Port_Rect[startPort]
		endRect := diagramStructure.map_Port_Rect[endPort]

		if startRect == nil || endRect == nil {
			continue
		}

		link := svgAssociationLinkAsCT[AbstractType](
			stager,
			startRect, endRect,
			controlFlowShape,
			layer,
			false)

		link.Stroke = "#333333"
		link.StrokeWidth = 1.5
		_ = link
	}
}

func (stager *Stager) drawDataFlowShapes(diagramStructure *DiagramStructure, layer *svg.Layer) {
	for _, dataFlowShape := range diagramStructure.DataFlow_Shapes {
		if dataFlowShape.GetIsHidden() {
			continue
		}

		var startRect, endRect *svg.Rect
		df := dataFlowShape.DataFlow

		switch df.Type {
		case DataFlow_Port2Port:
			if df.StartPort == nil || df.EndPort == nil {
				log.Panic("There should be a start port and an end port")
			}
			startRect = diagramStructure.map_Port_Rect[df.StartPort]
			endRect = diagramStructure.map_Port_Rect[df.EndPort]
		case DataFlow_ExternalPart2Port:
			if df.StartExternalPart == nil || df.EndPort == nil {
				log.Panic("There should be a start external part and an end port")
			}
			startRect = diagramStructure.map_ExternalPart_Rect[df.StartExternalPart]
			endRect = diagramStructure.map_Port_Rect[df.EndPort]
		case DataFlow_Port2ExternalPart:
			if df.StartPort == nil || df.EndExternalPart == nil {
				log.Panic("There should be a start port and an end external part")
			}
			startRect = diagramStructure.map_Port_Rect[df.StartPort]
			endRect = diagramStructure.map_ExternalPart_Rect[df.EndExternalPart]
		}

		if startRect == nil || endRect == nil {
			continue
		}

		link := svgAssociationLinkAsCT[AbstractType](
			stager,
			startRect, endRect,
			dataFlowShape,
			layer,
			false)

		link.Presentation.StrokeDashArray = "5,5"
		link.Stroke = "#4CAF50"
		link.StrokeWidth = 1.5

		nbDataShapes := len(dataFlowShape.dataShapes)
		for idx, dataShape := range dataFlowShape.dataShapes {
			data := dataShape.Data

			textXOffset := 4.0
			if data.SVG_Path != "" {
				textXOffset = 24.0
			}

			textAnchoredText := &svg.LinkAnchoredText{
				Name: func() string {
					if dataShape.Data.Acronym != "" {
						return dataShape.Data.Acronym
					}
					return dataShape.Data.Name
				}(),
				Content: data.Name,
				Presentation: svg.Presentation{
					Color:       "#333333",
					FillOpacity: 1.0,
					Stroke:      "#FFFFFF",
					StrokeWidth: 2.0,
				},
				TextAttributes: svg.TextAttributes{
					FontWeight: "600",
				},

				Y_Offset: float64(-nbDataShapes+idx+1)*18.0 - 4.0,
				X_Offset: textXOffset,
			}
			link.TextAtCorner = append(link.TextAtCorner, textAnchoredText)

			if data.SVG_Path != "" {
				if data.InverseAppliedScaling == 0 {
					data.InverseAppliedScaling = 1.0
				}
				path := &svg.LinkAnchoredPath{
					Name:       data.GetName(),
					Definition: data.SVG_Path,
					Presentation: svg.Presentation{
						StrokeWidth: 0,
						Color:       "#757575",
						FillOpacity: 1,
					},
					X_Offset:            4.0,
					Y_Offset:            float64(-nbDataShapes+idx+1)*18.0 - 4.0,
					ScalePropotionnally: true,
					AppliedScaling:      1.0 / data.InverseAppliedScaling,
				}
				link.PathAtCorner = append(link.PathAtCorner, path)
			}
		}
	}
}

func (stager *Stager) drawNoteShapes(diagramStructure *DiagramStructure, layer *svg.Layer) map[*Note]*svg.Rect {
	map_Note_Rect := make(map[*Note]*svg.Rect)
	for _, noteShape := range diagramStructure.Note_Shapes {
		if noteShape.GetIsHidden() {
			continue
		}

		rect := svgRect(
			stager,
			diagramStructure,
			noteShape,
			layer)

		map_Note_Rect[noteShape.Note] = rect

		rect.RX = 0
		rect.Color = "#FFF9C4"
		rect.Stroke = "#FBC02D"
		rect.StrokeWidth = 1.0

		if len(rect.RectAnchoredTexts) > 0 {
			abstractElement := noteShape.GetAbstractElement()
			content := abstractElement.GetName()
			if diagramStructure.GetIsShowPrefix() {
				content = abstractElement.GetComputedPrefix() + " " + content
			}
			content = "📝 " + content

			margin := 20.0
			wrapWidth := rect.Width - margin
			if wrapWidth > 0 {
				content = strutils.WrapStringPreservingNewlines(content, int(wrapWidth/(stager.GetRootLibrary().NbPixPerCharacter*0.7)))
			}

			rect.RectAnchoredTexts[0].Content = content
			rect.RectAnchoredTexts[0].FontWeight = "normal"
			rect.RectAnchoredTexts[0].FontStyle = "italic"
			rect.RectAnchoredTexts[0].TextAnchorType = svg.TEXT_ANCHOR_START
			rect.RectAnchoredTexts[0].RectAnchorType = svg.RECT_TOP_LEFT
			rect.RectAnchoredTexts[0].X_Offset = 10
			rect.RectAnchoredTexts[0].Y_Offset = 20
		}
	}
	return map_Note_Rect
}

func (stager *Stager) drawNotePortShapes(diagramStructure *DiagramStructure, layer *svg.Layer, map_Note_Rect map[*Note]*svg.Rect) {
	for _, notePortShape := range diagramStructure.NotePortShapes {
		if notePortShape.GetIsHidden() {
			continue
		}

		startNote := notePortShape.Note
		endPort := notePortShape.Port

		if startNote == nil || endPort == nil {
			continue
		}

		startRect := map_Note_Rect[startNote]
		endRect := diagramStructure.map_Port_Rect[endPort]

		if startRect == nil || endRect == nil {
			continue
		}

		link := svgAssociationLink(
			stager,
			startRect, endRect,
			notePortShape,
			startNote,
			layer,
			false)

		link.Type = svg.LINK_TYPE_LINE_WITH_CONTROL_POINTS
		link.StartAnchorType = svg.ANCHOR_CENTER
		link.EndAnchorType = svg.ANCHOR_CENTER
		link.HasEndArrow = false

		link.Presentation.StrokeDashArray = "5,5"
		link.Stroke = "#9E9E9E"
		link.StrokeWidth = 1.5
	}
}

func (stager *Stager) drawAllocatedSystemesAndResources(
	part *Part,
	diagramStructure *DiagramStructure,
	rect *svg.Rect,
	rectAnchorType svg.RectAnchorType) (boxHeight float64, allocatedResourceRect *svg.RectAnchoredRect) {
	root := stager.GetRootLibrary()
	const HeightBetween2AttributeShapes = 20.0

	totalLines := 0
	X_Offset := 35.0
	Y_Offset := 20.0

	// draw allocated system shapes that are within the part
	for _, system := range part.Systemes {
		key := allocatedSystemShapeKey{
			part: part,
			system:     system,
		}
		allocatedSystemShape := diagramStructure.map_AllocatedSystemShapeKey_AllocatedSystemShape[key]
		if allocatedSystemShape == nil {
			continue
		}

		content := system.Name
		if rect.Width > 0 {
			content = strutils.WrapStringPreservingNewlinesScaled(content, rect.Width, float64(root.NbPixPerCharacter), 15.0, 16.0)
		}

		allocatedSystemText := &svg.RectAnchoredText{
			Name:    allocatedSystemShape.Name,
			Content: content,
			Presentation: svg.Presentation{
				StrokeWidth: 0,
				Color:       "#757575",
				FillOpacity: 1,
			},
			TextAttributes: svg.TextAttributes{
				FontStyle: "italic",
				FontSize:  "15px",
			},

			DominantBaseline: svg.DominantBaselineMiddle,
			X_Offset:         X_Offset,
			Y_Offset:         Y_Offset + float64(totalLines)*HeightBetween2AttributeShapes,
			RectAnchorType:   rectAnchorType,
			TextAnchorType:   svg.TEXT_ANCHOR_START,
			URLPath:          "../../../References/Systemes/" + system.GetReferencePath() + "/index.html",
			URLTarget:        svg.LINK_TARGET_BLANK,
		}
		if rectAnchorType == svg.RECT_BOTTOM {
			allocatedSystemText.RectAnchorType = svg.RECT_BOTTOM_LEFT
		} else {
			allocatedSystemText.RectAnchorType = svg.RECT_TOP_LEFT
		}
		rect.RectAnchoredTexts = append(rect.RectAnchoredTexts, allocatedSystemText)

		lines := strings.Split(content, "\n")
		totalLines += len(lines)

		// add a rect anchored icon
		if system.SVG_Path != "" {
			if system.InverseAppliedScaling == 0 {
				system.InverseAppliedScaling = 1.0
			}
			allocatedSystemPath := &svg.RectAnchoredPath{
				Name:       system.GetName(),
				Definition: system.SVG_Path,
				Presentation: svg.Presentation{
					StrokeWidth: 0,
					Color:       "#757575",
					FillOpacity: 1,
				},
				X_Offset:            10,
				Y_Offset:            10 + float64(totalLines)*HeightBetween2AttributeShapes,
				ScalePropotionnally: true,
				AppliedScaling:      1.25 / system.InverseAppliedScaling,
			}
			if rectAnchorType == svg.RECT_BOTTOM {
				allocatedSystemPath.RectAnchorType = svg.RECT_BOTTOM_LEFT
			} else {
				allocatedSystemPath.RectAnchorType = svg.RECT_TOP_LEFT
			}
			rect.RectAnchoredPaths = append(rect.RectAnchoredPaths, allocatedSystemPath)
		}
	}

	// draw allocated resource shapes that are within the part
	for _, resource := range part.Resources {
		key := allocatedResourceShapeKey{
			part: part,
			resource:    resource,
		}
		allocatedResourceShape := diagramStructure.map_AllocatedResourceShapeKey_AllocatedResourceShape[key]
		if allocatedResourceShape == nil {
			continue
		}

		content := resource.Name
		if rect.Width > 0 {
			content = strutils.WrapStringPreservingNewlinesScaled(content, rect.Width, float64(root.NbPixPerCharacter), 15.0, 16.0)
		}

		allocatedResourceText := &svg.RectAnchoredText{
			Name:    allocatedResourceShape.Name,
			Content: content,
			Presentation: svg.Presentation{
				StrokeWidth: 0,
				Color:       "#757575",
				FillOpacity: 1,
			},
			TextAttributes: svg.TextAttributes{
				FontStyle: "italic",
				FontSize:  "15px",
			},

			DominantBaseline: svg.DominantBaselineMiddle,
			X_Offset:         X_Offset,
			Y_Offset:         Y_Offset + float64(totalLines)*HeightBetween2AttributeShapes,
			RectAnchorType:   rectAnchorType,
			TextAnchorType:   svg.TEXT_ANCHOR_START,
			URLPath:          "../../../References/Resources/" + resource.GetReferencePath() + "/index.html",
			URLTarget:        svg.LINK_TARGET_BLANK,
		}
		if rectAnchorType == svg.RECT_BOTTOM {
			allocatedResourceText.RectAnchorType = svg.RECT_BOTTOM_LEFT
		} else {
			allocatedResourceText.RectAnchorType = svg.RECT_TOP_LEFT
		}
		rect.RectAnchoredTexts = append(rect.RectAnchoredTexts, allocatedResourceText)

		lines := strings.Split(content, "\n")
		totalLines += len(lines)

		// add a rect anchored icon
		if resource.SVG_Path != "" {
			if resource.InverseAppliedScaling == 0 {
				resource.InverseAppliedScaling = 1.0
			}
			allocatedResourcePath := &svg.RectAnchoredPath{
				Name:       resource.GetName(),
				Definition: resource.SVG_Path,
				Presentation: svg.Presentation{
					StrokeWidth: 0,
					Color:       "#757575",
					FillOpacity: 1,
				},
				X_Offset:            10,
				Y_Offset:            10 + float64(totalLines)*HeightBetween2AttributeShapes,
				ScalePropotionnally: true,
				AppliedScaling:      1.25 / resource.InverseAppliedScaling,
			}
			if rectAnchorType == svg.RECT_BOTTOM {
				allocatedResourcePath.RectAnchorType = svg.RECT_BOTTOM_LEFT
			} else {
				allocatedResourcePath.RectAnchorType = svg.RECT_TOP_LEFT
			}
			rect.RectAnchoredPaths = append(rect.RectAnchoredPaths, allocatedResourcePath)
		}
	}
	// draw a rect around the allocated resource and system shapes if there is at least one
	boxHeight = float64(totalLines)*HeightBetween2AttributeShapes + Y_Offset - 5.0
	if totalLines > 0 {
		lineWidth := 1.0
		allocatedResourceRect = &svg.RectAnchoredRect{
			Name: part.Name + "_allocated_resources_and_systemes",
			Presentation: svg.Presentation{
				Stroke:        "#CCCCCC",
				StrokeWidth:   lineWidth,
				StrokeOpacity: 1,
			},
			X_Offset:       lineWidth,
			Y_Offset:       lineWidth,
			Height:         boxHeight - 2*lineWidth,
			Width:          rect.Width - 2*lineWidth,
			RectAnchorType: svg.RECT_TOP_LEFT,
		}

		if rectAnchorType == svg.RECT_BOTTOM {
			allocatedResourceRect.RectAnchorType = svg.RECT_BOTTOM_LEFT
		}
		rect.RectAnchoredRects = append(rect.RectAnchoredRects, allocatedResourceRect)
	}
	return
}
