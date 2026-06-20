package models

import (
	"log"

	"github.com/fullstack-lang/gong/lib/strutils"
	svg "github.com/fullstack-lang/gong/lib/svg/go/models"
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
		stager.drawPortShapes(diagramStructure, layer, rectOfOwningSystem)
	}
	stager.drawExternalPartShapes(diagramStructure, layer)

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

		// title is on top of the rect
		title := rect.RectAnchoredTexts[0]
		title.RectAnchorType = svg.RECT_TOP
		title.Y_Offset = +30

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
		rect.OnMove = onMoveRectElement(stager, systemShape, false)
		rect.OnResize = onResizeRectElement(stager, systemShape)
		diagramStructure.map_System_Rect[systemShape.System] = rect
	}
}

func (stager *Stager) drawPartShapes(diagramStructure *DiagramStructure, layer *svg.Layer, rectOfOwningSystem *svg.Rect) {
	root := stager.getRootLibrary()
	diagramStructure.map_Part_Rect = make(map[*Part]*svg.Rect)

	verticalTopMarginForTitle := 60.0

	for _, partShape := range diagramStructure.Part_Shapes {

		if partShape.IsHidden {
			continue
		}

		partRect := new(svg.Rect)
		layer.Rects = append(layer.Rects, partRect)
		diagramStructure.map_SvgRect_Part[partRect] = partShape.Part
		diagramStructure.map_Part_Rect[partShape.Part] = partRect

		// part rect cannot espace owning system shape
		partRect.EnclosingRect = rectOfOwningSystem
		rectOfOwningSystem.Peers = append(rectOfOwningSystem.Peers, partRect)

		partRect.Name = partShape.GetName()
		partRect.Stroke = "#E0E0E0"
		partRect.StrokeWidth = 1
		partRect.StrokeOpacity = 1

		partRect.X = partShape.GetX()
		partRect.Y = partShape.GetY()
		partRect.Width = partShape.GetWidth()
		partRect.Height = partShape.GetHeight()

		partRect.Color = "#FFFFFF"
		partRect.FillOpacity = 1.0

		// rect cannot move
		partRect.CanMoveHorizontaly = true
		partRect.CanMoveVerticaly = true

		partRect.CanHaveBottomHandle = true
		partRect.CanHaveLeftHandle = true
		partRect.CanHaveRightHandle = true
		partRect.CanHaveTopHandle = true

		// visuals
		partRect.RX = 0
		partRect.StrokeWidth = 1

		// override default behavior, we need to commit when the rect is moved
		partRect.OnSelect = func() {
			stager.probeForm.FillUpFormFromGongstruct(partShape.Part, GetPointerToGongstructName[*Part]())
		}
		partRect.OnMove = onMoveRectElement(stager, partShape, true)
		partRect.OnResize = onResizeRectElement(stager, partShape)

		{
			title := new(svg.RectAnchoredText)
			title.Name = partShape.GetAbstractElement().GetName()

			content := partShape.GetAbstractElement().GetName()
			if diagramStructure.GetIsShowPrefix() {
				content = partShape.GetAbstractElement().GetComputedPrefix() + " " + content
			}

			if partRect.Width > 0 {
				content = strutils.WrapStringPreservingNewlines(content, int(partRect.Width/root.NbPixPerCharacter))
			}
			title.Content = content
			title.StrokeWidth = 0
			title.StrokeOpacity = 1
			title.Color = "#333333"
			title.FillOpacity = 1

			title.FontSize = "16px"
			title.FontWeight = "500"
			title.X_Offset = 0
			title.Y_Offset = verticalTopMarginForTitle / 2.0
			title.RectAnchorType = svg.RECT_TOP
			title.TextAnchorType = svg.TEXT_ANCHOR_CENTER

			partRect.RectAnchoredTexts = append(partRect.RectAnchoredTexts, title)
		}
	}

	// make every part shape an obstable for the others
	for _, partShape1 := range diagramStructure.Part_Shapes {
		rect1 := diagramStructure.map_Part_Rect[partShape1.Part]

		for _, partShape2 := range diagramStructure.Part_Shapes {
			rect2 := diagramStructure.map_Part_Rect[partShape2.Part]

			if partShape1 != partShape2 {
				rect1.Obstacles = append(rect1.Obstacles, rect2)
				rect2.Obstacles = append(rect2.Obstacles, rect1)
			}
		}
	}

}

func (stager *Stager) drawExternalPartShapes(diagramStructure *DiagramStructure, layer *svg.Layer) {

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

		rect.OnSelect = onSelectRectElement(stager, externalPartShape.Part)
		rect.OnMove = onMoveRectElement(stager, externalPartShape, false)
		rect.OnResize = onResizeRectElement(stager, externalPartShape)

		externalPartWidth := 5.0
		if externalPartShape.TailHeigth == 0 {
			externalPartShape.TailHeigth = 100.0
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
			Height:              externalPartShape.TailHeigth,
			X:                   rect.X + (rect.Width-externalPartWidth)/2.0,
			Y:                   rect.Y + rect.Height,
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
			diffSize := externalPartShape.TailHeigth != height

			if diffSize {
				externalPartShape.TailHeigth = height
				stager.stage.CommitWithSuspendedCallbacks()
			}
		}
		layer.Rects = append(layer.Rects, tailRect)
	}
}

func (stager *Stager) drawPortShapes(diagramStructure *DiagramStructure, layer *svg.Layer, rectOfOwningSystem *svg.Rect) {
	rm := GetSliceOfPointersReverseMap[Part, Port](GetAssociationName[Part]().Ports[0].Name, stager.stage)

	diagramStructure.map_Port_Rect = make(map[*Port]*svg.Rect)
	for _, portShape := range diagramStructure.Port_Shapes {

		part := portShape.Port.owningPart
		partShape := diagramStructure.map_Part_PartShape[part]
		partRect := diagramStructure.map_Part_Rect[part]
		if partShape.IsHidden {
			continue
		}

		if portShape.IsHidden {
			continue
		}

		port := portShape.Port
		parts := rm[port]
		if len(parts) == 0 {
			continue
		}

		systemRect := diagramStructure.map_System_Rect[diagramStructure.owningSystem]
		if systemRect == nil {
			continue
		}

		portRect := svgRect(
			stager,
			diagramStructure,
			portShape,
			layer)
		diagramStructure.map_SvgRect_PortShape[portRect] = portShape

		portRect.URLPath = "../../../References/Ports/" + port.GetReferencePath() + "/index.html"
		portRect.URLTarget = svg.LINK_TARGET_BLANK

		portRect.RectAnchoredTexts[0].URLPath = portRect.URLPath
		portRect.RectAnchoredTexts[0].URLTarget = svg.LINK_TARGET_BLANK

		// make the port be anchored to the border of the part shape
		portRect.AnchoredTo = partRect
		partRect.Peers = append(partRect.Peers, portRect)
		rectOfOwningSystem.Peers = append(rectOfOwningSystem.Peers, portRect)

		portRect.Color = "#E3F2FD"
		portRect.FillOpacity = 1.0
		portRect.Stroke = "#90CAF9"
		portRect.StrokeWidth = 1.5
		portRect.RX = 5.0

		if len(portRect.RectAnchoredTexts) > 0 {
			portRect.RectAnchoredTexts[0].FontFamily = "sans-serif"
			portRect.RectAnchoredTexts[0].Color = "#333333"
			portRect.RectAnchoredTexts[0].RectAnchorType = svg.RECT_TOP
			portRect.RectAnchoredTexts[0].Y_Offset = -15
		}

		diagramStructure.map_Port_Rect[portShape.Port] = portRect
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
