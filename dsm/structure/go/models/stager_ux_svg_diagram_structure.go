package models

import (
	"fmt"
	"log"
	"math"
	"slices"
	"strings"

	"github.com/fullstack-lang/gong/lib/strutils"
	svg "github.com/fullstack-lang/gong/lib/svg/go/models"
)

func (stager *Stager) svg() {
	log.Println("svg")
	stager.systemDiagramSvgStage.Reset()

	var diagramStructure *DiagramStructure
	{
		for diagramsystem_ := range *stager.stage.GetInstancesSet[*DiagramStructure]() {
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

	stager.systemDiagramSvgStage.StageBranch(svgObject)
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
	diagramStructure.map_SvgRect_NoteShape = map[*svg.Rect]*NoteShape{}

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
	map_DataFlow_Link := stager.drawDataFlowShapes(diagramStructure, layer)
	map_Note_Rect := stager.drawNoteShapes(diagramStructure, layer)
	stager.drawNotePortShapes(diagramStructure, layer, map_Note_Rect)
	stager.drawNotePartShapes(diagramStructure, layer, map_Note_Rect)
	stager.drawNoteDataFlowShapes(diagramStructure, layer, map_Note_Rect, map_DataFlow_Link)

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
		syncRectPresentation(systemShape, rect)

		// override default behavior, we need to commit when the rect is moved
		// we need to update the position of the parts shapes and port shapes that are within the system
		rect.OnSelect = onSelectRectShape(stager, systemShape.System, systemShape)
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
		partRect.OnSelect = onSelectRectShape(stager, partShape.Part, partShape)
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
			title.Y_Offset = 0
			title.RectAnchorType = svg.RECT_TOP
			title.RectAnchorType = svg.RECT_CENTER_MIDDLE
			title.TextAnchorType = svg.TEXT_ANCHOR_CENTER

			partRect.RectAnchoredTexts = append(partRect.RectAnchoredTexts, title)
		}
		syncRectPresentation(partShape, partRect)

		for _, path := range partShape.Part.PartAnchoredPath {
			rectAnchoredPath := &svg.RectAnchoredPath{
				Name:                        path.Name,
				Definition:                  path.Definition,
				X_Offset:                    path.X_Offset,
				Y_Offset:                    path.Y_Offset,
				RectAnchorType:              svg.RectAnchorType(path.RectAnchorType),
				ScalePropotionnally:         path.ScalePropotionnally,
				AppliedScaling:              path.AppliedScaling,
				Color:                       path.Color,
				FillOpacity:                 path.FillOpacity,
				Stroke:                      path.Stroke,
				StrokeOpacity:               path.StrokeOpacity,
				StrokeWidth:                 path.StrokeWidth,
				StrokeDashArray:             path.StrokeDashArray,
				StrokeDashArrayWhenSelected: path.StrokeDashArrayWhenSelected,
				Transform:                   path.Transform,
			}
			partRect.RectAnchoredPaths = append(partRect.RectAnchoredPaths, rectAnchoredPath)
		}
	}

	// make every part shape an obstable for the others
	for _, partShape1 := range diagramStructure.Part_Shapes {
		rect1 := diagramStructure.map_Part_Rect[partShape1.Part]
		if rect1 == nil {
			continue
		}

		for _, partShape2 := range diagramStructure.Part_Shapes {
			rect2 := diagramStructure.map_Part_Rect[partShape2.Part]
			if rect2 == nil {
				continue
			}

			if partShape1 != partShape2 {
				rect1.Obstacles = append(rect1.Obstacles, rect2)
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
		syncRectPresentation(externalPartShape, rect)

		rect.OnSelect = onSelectRectShape(stager, externalPartShape.Part, externalPartShape)
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
			Name:                "Tail" + rect.GetName(),
			CanMoveHorizontaly:  true,
			CanMoveVerticaly:    true,
			Stroke:              "#9E9E9E",
			StrokeWidth:         1.5,
			StrokeOpacity:       1,
			StrokeDashArray:     "5 5",
			Color:               "transparent",
			FillOpacity:         0.0,
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
	root := stager.GetRootLibrary()
	_, partToLibrary := stager.buildSystemAndPartLibraryMaps()
	rm := stager.stage.GetSliceOfPointersReverseMap[Part, Port](GongGetAssociationName[Part]().Ports[0].Name)

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

		// make the port be anchored to the border of the part shape
		portRect.AnchoredTo = partRect
		partRect.Peers = append(partRect.Peers, portRect)
		rectOfOwningSystem.Peers = append(rectOfOwningSystem.Peers, portRect)

		portRect.Color = "#E3F2FD"
		portRect.FillOpacity = 1.0
		portRect.Stroke = "#90CAF9"
		portRect.StrokeWidth = 1.5
		portRect.RX = 5.0

		// if the diagram has discrete port, we remove the port title
		if diagramStructure.IsWithDiscretePorts {
			portRect.RectAnchoredTexts = slices.Delete(portRect.RectAnchoredTexts, 0, len(portRect.RectAnchoredTexts))
			portRect.FillOpacity = 0.1
			portRect.StrokeOpacity = 0.3

			portRect.X += portRect.Width / 4.0
			portRect.Y += portRect.Height / 4.0
			portRect.Width /= 2.0
			portRect.Height /= 2.0

			originalOnMove := portRect.OnMove
			if originalOnMove != nil {
				portRect.OnMove = func(x float64, y float64) {
					originalOnMove(x-portRect.Width/2.0, y-portRect.Height/2.0)
				}
			}
			originalOnResize := portRect.OnResize
			if originalOnResize != nil {
				portRect.OnResize = func(x float64, y float64, width float64, height float64) {
					originalOnResize(x-width/2.0, y-height/2.0, width*2.0, height*2.0)
				}
			}
		}

		if len(portRect.RectAnchoredTexts) > 0 {
			portRect.URLPath = "../../../References/Ports/" + port.GetReferencePath() + "/index.html"
			portRect.URLTarget = svg.LINK_TARGET_BLANK

			portLabel := port.GetName()
			lib := partToLibrary[part]
			if lib == nil && part != nil {
				lib = part.GetOwningLibrary()
			}
			if lib != nil && lib.IsPortAndDataFlowNamesAutoGenerated && part != nil {
				if idx := slices.Index(part.Ports, port); idx >= 0 {
					portLabel = fmt.Sprintf("P%d", idx+1)
				}
			}
			if diagramStructure.GetIsShowPrefix() && port.GetComputedPrefix() != "" {
				portLabel = port.GetComputedPrefix() + " " + portLabel
			}
			wrapWidth := math.Max(portRect.Width, 120.0)
			if wrapWidth > 0 && root != nil && root.NbPixPerCharacter > 0 {
				portLabel = strutils.WrapStringPreservingNewlines(portLabel, int(wrapWidth/root.NbPixPerCharacter))
			}
			lines := strings.Split(portLabel, "\n")

			portTitle := portRect.RectAnchoredTexts[0]
			portTitle.Content = portLabel
			portTitle.URLPath = portRect.URLPath
			portTitle.URLTarget = svg.LINK_TARGET_BLANK
			portTitle.FontFamily = "sans-serif"
			portTitle.FontSize = "11px"
			portTitle.FontWeight = "400"
			portTitle.Color = "#666666"
			portTitle.Stroke = "#666666"
			portTitle.StrokeWidth = 0
			portTitle.StrokeOpacity = 1.0
			portTitle.RectAnchorType = svg.RECT_TOP
			portTitle.Y_Offset = -6.0 - float64(len(lines)-1)*13.0
		}
		if !diagramStructure.IsWithDiscretePorts {
			syncRectPresentation(portShape, portRect)
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
		syncLinkPresentation(controlFlowShape, link)
		_ = link
	}
}

func (stager *Stager) drawNotePartShapes(diagramStructure *DiagramStructure, layer *svg.Layer, map_Note_Rect map[*Note]*svg.Rect) {
	for _, notePartShape := range diagramStructure.NotePartShapes {
		if notePartShape.GetIsHidden() {
			continue
		}

		startNote := notePartShape.Note
		endPart := notePartShape.Part

		if startNote == nil || endPart == nil {
			continue
		}

		startRect := map_Note_Rect[startNote]
		endRect := diagramStructure.map_Part_Rect[endPart]
		if endRect == nil {
			endRect = diagramStructure.map_ExternalPart_Rect[endPart]
		}

		if startRect == nil || endRect == nil {
			continue
		}

		link := svgAssociationLink(
			stager,
			startRect, endRect,
			notePartShape,
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
		syncLinkPresentation(notePartShape, link)

		_ = link
	}
}

func rectBorderPointToward(rect *svg.Rect, targetX, targetY float64) (float64, float64) {
	centerX := rect.X + rect.Width/2.0
	centerY := rect.Y + rect.Height/2.0

	dx := targetX - centerX
	dy := targetY - centerY
	mag := math.Sqrt(dx*dx + dy*dy)
	if mag == 0 {
		return centerX, centerY
	}

	ux := dx / mag
	uy := dy / mag

	borderX := centerX
	borderY := centerY

	if uy < 0 {
		t := (rect.Y - centerY) / uy
		px := centerX + t*ux
		if px >= rect.X && px <= rect.X+rect.Width {
			borderX = px
			borderY = rect.Y
		}
	}
	if uy > 0 {
		t := (rect.Y + rect.Height - centerY) / uy
		px := centerX + t*ux
		if px >= rect.X && px <= rect.X+rect.Width {
			borderX = px
			borderY = rect.Y + rect.Height
		}
	}
	if ux < 0 {
		t := (rect.X - centerX) / ux
		py := centerY + t*uy
		if py >= rect.Y && py <= rect.Y+rect.Height {
			borderX = rect.X
			borderY = py
		}
	}
	if ux > 0 {
		t := (rect.X + rect.Width - centerX) / ux
		py := centerY + t*uy
		if py >= rect.Y && py <= rect.Y+rect.Height {
			borderX = rect.X + rect.Width
			borderY = py
		}
	}

	return borderX, borderY
}

func (stager *Stager) drawDataFlowShapes(diagramStructure *DiagramStructure, layer *svg.Layer) map[*DataFlow]*svg.Link {
	root := stager.GetRootLibrary()
	map_DataFlow_Link := make(map[*DataFlow]*svg.Link)

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
		link.StrokeOpacity = 1.0
		syncLinkPresentation(dataFlowShape, link)

		link.Type = svg.LINK_TYPE_LINE_WITH_CONTROL_POINTS
		link.StartAnchorType = svg.ANCHOR_CENTER
		link.EndAnchorType = svg.ANCHOR_CENTER

		link.StartArrowSize = link.EndArrowSize
		switch dataFlowShape.DataFlow.Direction {
		case Unspecified:
			link.HasEndArrow = false
		case BothWays:
			link.HasStartArrow = true
		case Backward:
			link.HasEndArrow = false
			link.HasStartArrow = true
		}

		map_DataFlow_Link[df] = link

		var visibleDataShapes []*DataShape
		for _, ds := range dataFlowShape.dataShapes {
			if ds != nil && !ds.GetIsHidden() && ds.Data != nil {
				visibleDataShapes = append(visibleDataShapes, ds)
			}
		}

		var visibleProtocolShape *ProtocolShape
		if ps := dataFlowShape.protocolShape; ps != nil && !ps.GetIsHidden() && ps.Protocol != nil {
			visibleProtocolShape = ps
		}

		var visibleProtocolLayerShapes []*ProtocolLayerShape
		for _, pls := range dataFlowShape.protocolLayerShapes {
			if pls != nil && !pls.GetIsHidden() && pls.ProtocolLayer != nil {
				visibleProtocolLayerShapes = append(visibleProtocolLayerShapes, pls)
			}
		}

		if len(visibleDataShapes) == 0 && visibleProtocolShape == nil && len(visibleProtocolLayerShapes) == 0 {
			continue
		}

		startCenterX := startRect.X + startRect.Width/2.0
		startCenterY := startRect.Y + startRect.Height/2.0
		endCenterX := endRect.X + endRect.Width/2.0
		endCenterY := endRect.Y + endRect.Height/2.0

		startBorderX, startBorderY := rectBorderPointToward(startRect, endCenterX, endCenterY)
		endBorderX, endBorderY := rectBorderPointToward(endRect, startCenterX, startCenterY)

		anchorX := (startBorderX + endBorderX) / 2.0
		anchorY := (startBorderY + endBorderY) / 2.0

		appendPeerUnique := func(host *svg.Rect, peer *svg.Rect) {
			if host != nil && !slices.Contains(host.Peers, peer) {
				host.Peers = append(host.Peers, peer)
			}
		}
		attachMidpointBoxPeers := func(boxRect *svg.Rect) {
			appendPeerUnique(startRect, boxRect)
			appendPeerUnique(endRect, boxRect)
			for _, r := range layer.Rects {
				if slices.Contains(r.Peers, startRect) || slices.Contains(r.Peers, endRect) {
					appendPeerUnique(r, boxRect)
				}
			}
		}

		type preparedDataItem struct {
			dataShape  *DataShape
			content    string
			leftMargin float64
			width      float64
			height     float64
		}

		const lineHeight = 18.0
		const verticalPadding = 6.0
		const stackGap = 2.0

		var items []preparedDataItem
		totalStackHeight := 0.0

		for idx, dataShape := range visibleDataShapes {
			data := dataShape.Data

			width := dataShape.GetWidth()
			if width <= 0 {
				width = 180.0
				dataShape.SetWidth(width)
			}

			leftMargin := 6.0
			rightMargin := 6.0
			if data.SVG_Path != "" {
				leftMargin = 24.0
			}

			content := data.Name
			if data.Acronym != "" {
				content = data.Acronym
			}
			if diagramStructure.GetIsShowPrefix() && data.GetComputedPrefix() != "" {
				content = data.GetComputedPrefix() + " " + content
			}

			availableWidth := width - (leftMargin + rightMargin)
			if availableWidth > 0 && root != nil && root.NbPixPerCharacter > 0 {
				content = strutils.WrapStringPreservingNewlines(content, int(availableWidth/root.NbPixPerCharacter))
			}

			lines := strings.Split(content, "\n")
			height := float64(len(lines))*lineHeight + verticalPadding
			dataShape.SetHeight(height)

			items = append(items, preparedDataItem{
				dataShape:  dataShape,
				content:    content,
				leftMargin: leftMargin,
				width:      width,
				height:     height,
			})
			totalStackHeight += height
			if idx > 0 {
				totalStackHeight += stackGap
			}
		}

		currentY := anchorY - 4.0 - totalStackHeight
		for _, item := range items {
			dataShape := item.dataShape
			data := dataShape.Data

			rectX := anchorX + 4.0
			rectY := currentY
			currentY += item.height + stackGap

			dataShape.SetX(rectX)
			dataShape.SetY(rectY)

			dataRect := &svg.Rect{
				Name:          data.GetName(),
				X:             rectX,
				Y:             rectY,
				Width:         item.width,
				Height:        item.height,
				RX:            3.0,
				Color:         "#FFFFFF",
				FillOpacity:   0.85,
				Stroke:        "#A5D6A7",
				StrokeWidth:   1.0,
				StrokeOpacity: 1.0,
				IsSelectable:  true,
			}

			title := &svg.RectAnchoredText{
				Name:             data.GetName(),
				Content:          item.content,
				Color:            "#333333",
				FillOpacity:      1.0,
				Stroke:           "#333333",
				StrokeWidth:      0,
				StrokeOpacity:    1.0,
				FontWeight:       "600",
				FontSize:         "14px",
				X_Offset:         item.leftMargin,
				Y_Offset:         0,
				RectAnchorType:   svg.RECT_LEFT_MIDDLE,
				TextAnchorType:   svg.TEXT_ANCHOR_START,
				DominantBaseline: svg.DominantBaselineCentral,
			}
			dataRect.RectAnchoredTexts = append(dataRect.RectAnchoredTexts, title)

			if data.SVG_Path != "" {
				if data.InverseAppliedScaling == 0 {
					data.InverseAppliedScaling = 1.0
				}
				path := &svg.RectAnchoredPath{
					Name:                data.GetName(),
					Definition:          data.SVG_Path,
					StrokeWidth:         0,
					StrokeOpacity:       1.0,
					Color:               "#757575",
					FillOpacity:         1,
					X_Offset:            4.0,
					Y_Offset:            -8.0,
					RectAnchorType:      svg.RECT_LEFT,
					ScalePropotionnally: true,
					AppliedScaling:      1.0 / data.InverseAppliedScaling,
				}
				dataRect.RectAnchoredPaths = append(dataRect.RectAnchoredPaths, path)
			}

			syncRectPresentation(dataShape, dataRect)

			dataRect.OnSelect = onSelectRectShape(stager, data, dataShape)
			if diagramStructure.IsEditable() {
				dataRect.CanHaveLeftHandle = true
				dataRect.CanHaveRightHandle = true
				dataRect.CanHaveTopHandle = false
				dataRect.CanHaveBottomHandle = false

				dataRect.CanMoveHorizontaly = true
				dataRect.CanMoveVerticaly = true

				dataShape.SetReceiver(dataShape)
				dataRect.OnMove = func(x, y float64) {
					stager.stage.Commit()
				}
				dataRect.OnResize = func(x, y, width, height float64) {
					if width < 30.0 {
						width = 30.0
					}
					dataShape.SetWidth(width)
					stager.stage.Commit()
				}
			}

			attachMidpointBoxPeers(dataRect)
			layer.Rects = append(layer.Rects, dataRect)
		}

		// Render ProtocolShape and ProtocolLayerShapes stacked below the DataFlow midpoint
		currentProtoY := anchorY + 4.0
		const protoLineHeight = 15.0
		const protoVerticalPadding = 4.0
		renderProtocolBox := func(
			abstractElem AbstractType,
			shape interface {
				ConcreteType
				RectShapeInterface
			},
			onSelect func(),
			rawLabel string,
			bgColor string,
			strokeColor string,
			textColor string,
			fontWeight string,
			fontSize string,
		) {
			width := shape.GetWidth()
			if width <= 0 {
				width = 140.0
				shape.SetWidth(width)
			}
			const sideMargin = 6.0
			content := rawLabel
			if diagramStructure.GetIsShowPrefix() && abstractElem.GetComputedPrefix() != "" {
				content = abstractElem.GetComputedPrefix() + " " + content
			}
			availableWidth := width - 2*sideMargin
			if availableWidth > 0 && root != nil && root.NbPixPerCharacter > 0 {
				content = strutils.WrapStringPreservingNewlines(content, int(availableWidth/root.NbPixPerCharacter))
			}
			lines := strings.Split(content, "\n")
			height := float64(len(lines))*protoLineHeight + protoVerticalPadding
			shape.SetHeight(height)

			rectX := anchorX + 4.0
			rectY := currentProtoY
			currentProtoY += height + stackGap

			shape.SetX(rectX)
			shape.SetY(rectY)

			boxRect := &svg.Rect{
				Name:          abstractElem.GetName(),
				X:             rectX,
				Y:             rectY,
				Width:         width,
				Height:        height,
				RX:            3.0,
				Color:         bgColor,
				FillOpacity:   0.85,
				Stroke:        strokeColor,
				StrokeWidth:   0.75,
				StrokeOpacity: 1.0,
				IsSelectable:  true,
			}

			title := &svg.RectAnchoredText{
				Name:             abstractElem.GetName(),
				Content:          content,
				Color:            textColor,
				FillOpacity:      1.0,
				Stroke:           textColor,
				StrokeWidth:      0,
				StrokeOpacity:    1.0,
				FontWeight:       fontWeight,
				FontSize:         fontSize,
				X_Offset:         sideMargin,
				Y_Offset:         0,
				RectAnchorType:   svg.RECT_LEFT_MIDDLE,
				TextAnchorType:   svg.TEXT_ANCHOR_START,
				DominantBaseline: svg.DominantBaselineCentral,
			}
			boxRect.RectAnchoredTexts = append(boxRect.RectAnchoredTexts, title)

			syncRectPresentation(shape, boxRect)

			boxRect.OnSelect = onSelect
			if diagramStructure.IsEditable() {
				boxRect.CanHaveLeftHandle = true
				boxRect.CanHaveRightHandle = true
				boxRect.CanHaveTopHandle = false
				boxRect.CanHaveBottomHandle = false

				boxRect.CanMoveHorizontaly = true
				boxRect.CanMoveVerticaly = true

				shape.SetReceiver(shape)
				boxRect.OnMove = func(x, y float64) {
					stager.stage.Commit()
				}
				boxRect.OnResize = func(x, y, w, h float64) {
					if w < 30.0 {
						w = 30.0
					}
					shape.SetWidth(w)
					stager.stage.Commit()
				}
			}

			attachMidpointBoxPeers(boxRect)
			layer.Rects = append(layer.Rects, boxRect)
		}

		if visibleProtocolShape != nil {
			proto := visibleProtocolShape.Protocol
			label := proto.Name
			if proto.Acronym != "" {
				label = proto.Acronym
			}
			label = "«" + label + "»"
			renderProtocolBox(
				proto,
				visibleProtocolShape,
				onSelectRectShape(stager, proto, visibleProtocolShape),
				label,
				"#FAFAFA",
				"#CFD8DC",
				"#607D8B",
				"400",
				"11px",
			)
		}
		for _, layerShape := range visibleProtocolLayerShapes {
			protoLayer := layerShape.ProtocolLayer
			label := protoLayer.Name
			if protoLayer.Acronym != "" {
				label = protoLayer.Acronym
			}
			renderProtocolBox(
				protoLayer,
				layerShape,
				onSelectRectShape(stager, protoLayer, layerShape),
				label,
				"#FFFFFF",
				"#E0E0E0",
				"#757575",
				"400",
				"11px",
			)
		}
	}

	return map_DataFlow_Link
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
		diagramStructure.map_SvgRect_NoteShape[rect] = noteShape

		FormatNoteRect(stager, diagramStructure, noteShape, rect)
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
		syncLinkPresentation(notePortShape, link)
	}
}

func (stager *Stager) drawNoteDataFlowShapes(
	diagramStructure *DiagramStructure,
	layer *svg.Layer,
	map_Note_Rect map[*Note]*svg.Rect,
	map_DataFlow_Link map[*DataFlow]*svg.Link,
) {
	for _, noteDataFlowShape := range diagramStructure.NoteDataFlowShapes {
		if noteDataFlowShape.GetIsHidden() {
			continue
		}

		startNote := noteDataFlowShape.Note
		endDataFlow := noteDataFlowShape.DataFlow

		if startNote == nil || endDataFlow == nil {
			continue
		}

		startRect := map_Note_Rect[startNote]
		endLink := map_DataFlow_Link[endDataFlow]

		if startRect == nil || endLink == nil {
			continue
		}

		rectLinkLink := new(svg.RectLinkLink)
		rectLinkLink.Name = startRect.Name + " to " + endLink.Name
		rectLinkLink.Start = startRect
		rectLinkLink.End = endLink
		rectLinkLink.TargetAnchorPosition = 0.5

		rectLinkLink.Stroke = "#9E9E9E"
		rectLinkLink.StrokeWidth = 1.5
		rectLinkLink.StrokeOpacity = 1.0
		rectLinkLink.StrokeDashArray = "5,5"

		if noteDataFlowShape.GetIsPresentationOverridden() {
			if noteDataFlowShape.GetStroke() != "" {
				rectLinkLink.Stroke = noteDataFlowShape.GetStroke()
			}
			if noteDataFlowShape.GetStrokeWidth() != 0 {
				rectLinkLink.StrokeWidth = noteDataFlowShape.GetStrokeWidth()
			}
			if noteDataFlowShape.GetStrokeOpacity() != 0 {
				rectLinkLink.StrokeOpacity = noteDataFlowShape.GetStrokeOpacity()
			}
			if noteDataFlowShape.GetStrokeDashArray() != "" {
				rectLinkLink.StrokeDashArray = noteDataFlowShape.GetStrokeDashArray()
			}
		} else {
			noteDataFlowShape.SetStroke(rectLinkLink.Stroke)
			noteDataFlowShape.SetStrokeWidth(rectLinkLink.StrokeWidth)
			noteDataFlowShape.SetStrokeOpacity(rectLinkLink.StrokeOpacity)
			noteDataFlowShape.SetStrokeDashArray(rectLinkLink.StrokeDashArray)
		}

		layer.RectLinkLinks = append(layer.RectLinkLinks, rectLinkLink)
	}
}
