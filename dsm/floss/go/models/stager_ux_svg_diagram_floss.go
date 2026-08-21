package models

import (
	"fmt"

	"github.com/fullstack-lang/gong/lib/strutils"
	svg "github.com/fullstack-lang/gong/lib/svg/go/models"
)

func (stager *Stager) svg() {
	stager.systemDiagramSvgStage.Reset()

	var diagramFloss *DiagramFloss
	for diagramsystem_ := range *GetGongstructInstancesSet[DiagramFloss](stager.stage) {
		if diagramsystem_.IsChecked {
			diagramFloss = diagramsystem_
			break
		}
	}

	var diagramFlossEquation *DiagramFlossEquation
	for diagram_ := range *GetGongstructInstancesSet[DiagramFlossEquation](stager.stage) {
		if diagram_.IsChecked {
			diagramFlossEquation = diagram_
			break
		}
	}

	if diagramFloss != nil {
		svgObject := stager.generateSvgObject(diagramFloss)
		svg.StageBranch(stager.systemDiagramSvgStage, svgObject)
		stager.svgObjectDiagramFloss = svgObject
		stager.svgObjectDiagramFloss.OnUpdate = stager.onUpdateSVG
	} else if diagramFlossEquation != nil {
		svgObject := stager.generateSvgObjectFlossEquation(diagramFlossEquation)
		svg.StageBranch(stager.systemDiagramSvgStage, svgObject)
		stager.svgObjectDiagramFloss = svgObject
		stager.svgObjectDiagramFloss.OnUpdate = stager.onUpdateSVG
	}

	stager.systemDiagramSvgStage.Commit()
}

func (stager *Stager) generateSvgObject(diagramFloss *DiagramFloss) *svg.SVG {

	svgObject := (&svg.SVG{Name: `SVG`})
	stager.diagramFloss = diagramFloss

	svgObject.OverrideWidth = true
	svgObject.OverriddenWidth = diagramFloss.Width
	svgObject.OverrideHeight = true
	svgObject.OverriddenHeight = diagramFloss.Height

	svgObject.Name = diagramFloss.Name
	svgObject.IsEditable = diagramFloss.IsEditable()

	layer := (&svg.Layer{Name: "Layer 1"})
	svgObject.Layers = append(svgObject.Layers, layer)

	stager.drawSystemShapes(diagramFloss, layer)
	stager.drawComplexityShapes(diagramFloss, layer)
	stager.drawPerformanceShapes(diagramFloss, layer)
	stager.drawEffortShapes(diagramFloss, layer)
	stager.drawNoteShapes(diagramFloss, layer)
	stager.drawNoteComplexityShapes(diagramFloss, layer)
	stager.drawNotePerformanceShapes(diagramFloss, layer)
	stager.drawNoteEffortShapes(diagramFloss, layer)

	return svgObject
}


func (stager *Stager) drawSystemShapes(diagramFloss *DiagramFloss, layer *svg.Layer) {
	diagramFloss.map_System_Rect = make(map[*System]*svg.Rect)
	for _, systemShape := range diagramFloss.System_Shapes {
		if systemShape.IsHidden {
			continue
		}

		rect := new(svg.Rect)
		layer.Rects = append(layer.Rects, rect)

		rect.Name = systemShape.GetName()
		rect.X = systemShape.GetX()
		rect.Y = systemShape.GetY()
		rect.Width = systemShape.GetWidth()
		rect.Height = systemShape.GetHeight()

		rect.Color = "#FAFAFB"
		rect.FillOpacity = 1.0
		rect.Stroke = "#B0BEC5"
		rect.StrokeWidth = 2.0
		rect.StrokeOpacity = 1.0

		rect.CanHaveBottomHandle = true
		rect.CanHaveLeftHandle = true
		rect.CanHaveRightHandle = true
		rect.CanHaveTopHandle = true

		title := new(svg.RectAnchoredText)
		title.Name = systemShape.System.Name
		title.Color = "black"
		title.FillOpacity = 1.0
		title.Stroke = "black"
		title.StrokeWidth = 0
		title.StrokeOpacity = 1.0
		title.FontSize = "18px"
		title.FontWeight = "bold"
		title.RectAnchorType = svg.RECT_TOP
		title.TextAnchorType = svg.TEXT_ANCHOR_CENTER
		title.Y_Offset = 28
		rect.RectAnchoredTexts = append(rect.RectAnchoredTexts, title)


		root := stager.getRootLibrary()
		nbPixPerChar := root.NbPixPerCharacter
		if nbPixPerChar <= 0 {
			nbPixPerChar = 8.0
		}

		{
			content := systemShape.System.Name
			if diagramFloss.IsShowPrefix {
				content = systemShape.System.ComputedPrefix + " " + content
			}
			if rect.Width > 0 {
				content = strutils.WrapStringPreservingNewlinesScaled(content, rect.Width, nbPixPerChar, 18.0, 16.0)
			}
			title.Content = content
		}


		rect.OnSelect = func() {
			stager.probeForm.FillUpFormFromGongstruct(systemShape.System, GetPointerToGongstructName[*System]())
		}
		rect.OnMove = onMoveRectElement(stager, systemShape, false)
		rect.OnResize = onResizeRectElement(stager, systemShape)
		diagramFloss.map_System_Rect[systemShape.System] = rect
	}
}

func (stager *Stager) drawComplexityShapes(diagramFloss *DiagramFloss, layer *svg.Layer) {
	diagramFloss.map_Complexity_Rect = make(map[*Complexity]*svg.Rect)
	diagramFloss.map_SvgRect_ComplexityShape = make(map[*svg.Rect]*ComplexityShape)
	for _, complexityShape := range diagramFloss.Complexity_Shapes {
		if complexityShape.IsHidden {
			continue
		}

		rect := new(svg.Rect)
		layer.Rects = append(layer.Rects, rect)
		diagramFloss.map_Complexity_Rect[complexityShape.Complexity] = rect
		diagramFloss.map_SvgRect_ComplexityShape[rect] = complexityShape

		if diagramFloss.owningSystem != nil {
			if systemRect, ok := diagramFloss.map_System_Rect[diagramFloss.owningSystem]; ok {
				rect.EnclosingRect = systemRect
				systemRect.Peers = append(systemRect.Peers, rect)
			}
		}

		rect.Name = complexityShape.GetName()
		rect.X = complexityShape.GetX()
		rect.Y = complexityShape.GetY()
		rect.Width = complexityShape.GetWidth()
		rect.Height = complexityShape.GetHeight()

		// Amber / warm theme for Complexity
		rect.Color = "#FFF8E1"
		rect.FillOpacity = 1.0
		rect.Stroke = "#FFA000"
		rect.StrokeWidth = 1.5
		rect.StrokeOpacity = 1.0
		rect.RX = 6

		rect.CanMoveHorizontaly = true
		rect.CanMoveVerticaly = true
		rect.CanHaveBottomHandle = true
		rect.CanHaveLeftHandle = true
		rect.CanHaveRightHandle = true
		rect.CanHaveTopHandle = true

		rect.OnSelect = func() {
			stager.probeForm.FillUpFormFromGongstruct(complexityShape.Complexity, GetPointerToGongstructName[*Complexity]())
		}
		rect.OnMove = onMoveRectElement(stager, complexityShape, true)
		rect.OnResize = onResizeRectElement(stager, complexityShape)

		title := new(svg.RectAnchoredText)
		title.Name = complexityShape.Complexity.Name
		content := complexityShape.Complexity.Name
		if complexityShape.Complexity.Strength != 0 {
			content = fmt.Sprintf("%s (%.1f)", complexityShape.Complexity.Name, complexityShape.Complexity.Strength)
		}

		root := stager.getRootLibrary()
		nbPixPerChar := root.NbPixPerCharacter
		if nbPixPerChar <= 0 {
			nbPixPerChar = 8.0
		}
		if rect.Width > 0 {
			content = strutils.WrapStringPreservingNewlinesScaled(content, rect.Width, nbPixPerChar, 14.0, 16.0)
		}

		title.Content = content
		title.Color = "#E65100"
		title.FillOpacity = 1.0
		title.Stroke = "#E65100"
		title.StrokeWidth = 0
		title.StrokeOpacity = 1.0
		title.FontSize = "14px"
		title.FontWeight = "600"
		title.RectAnchorType = svg.RECT_CENTER_MIDDLE
		title.TextAnchorType = svg.TEXT_ANCHOR_CENTER
		rect.RectAnchoredTexts = append(rect.RectAnchoredTexts, title)
	}
}

func (stager *Stager) drawPerformanceShapes(diagramFloss *DiagramFloss, layer *svg.Layer) {
	diagramFloss.map_Performance_Rect = make(map[*Performance]*svg.Rect)
	diagramFloss.map_SvgRect_PerformanceShape = make(map[*svg.Rect]*PerformanceShape)
	for _, performanceShape := range diagramFloss.Performance_Shapes {
		if performanceShape.IsHidden {
			continue
		}

		rect := new(svg.Rect)
		layer.Rects = append(layer.Rects, rect)
		diagramFloss.map_Performance_Rect[performanceShape.Performance] = rect
		diagramFloss.map_SvgRect_PerformanceShape[rect] = performanceShape

		if diagramFloss.owningSystem != nil {
			if systemRect, ok := diagramFloss.map_System_Rect[diagramFloss.owningSystem]; ok {
				rect.EnclosingRect = systemRect
				systemRect.Peers = append(systemRect.Peers, rect)
			}
		}

		rect.Name = performanceShape.GetName()
		rect.X = performanceShape.GetX()
		rect.Y = performanceShape.GetY()
		rect.Width = performanceShape.GetWidth()
		rect.Height = performanceShape.GetHeight()

		// Emerald green capsule theme for Performance
		rect.Color = "#E8F5E9"
		rect.FillOpacity = 1.0
		rect.Stroke = "#2E7D32"
		rect.StrokeWidth = 1.5
		rect.StrokeOpacity = 1.0
		rect.RX = 14

		rect.CanMoveHorizontaly = true
		rect.CanMoveVerticaly = true
		rect.CanHaveBottomHandle = true
		rect.CanHaveLeftHandle = true
		rect.CanHaveRightHandle = true
		rect.CanHaveTopHandle = true

		rect.OnSelect = func() {
			stager.probeForm.FillUpFormFromGongstruct(performanceShape.Performance, GetPointerToGongstructName[*Performance]())
		}
		rect.OnMove = onMoveRectElement(stager, performanceShape, true)
		rect.OnResize = onResizeRectElement(stager, performanceShape)

		title := new(svg.RectAnchoredText)
		title.Name = performanceShape.Performance.Name
		content := performanceShape.Performance.Name
		if performanceShape.Performance.Strength != 0 {
			content = fmt.Sprintf("%s (%.1f)", performanceShape.Performance.Name, performanceShape.Performance.Strength)
		}

		root := stager.getRootLibrary()
		nbPixPerChar := root.NbPixPerCharacter
		if nbPixPerChar <= 0 {
			nbPixPerChar = 8.0
		}
		if rect.Width > 0 {
			content = strutils.WrapStringPreservingNewlinesScaled(content, rect.Width, nbPixPerChar, 14.0, 16.0)
		}

		title.Content = content
		title.Color = "#1B5E20"
		title.FillOpacity = 1.0
		title.Stroke = "#1B5E20"
		title.StrokeWidth = 0
		title.StrokeOpacity = 1.0
		title.FontSize = "14px"
		title.FontWeight = "600"
		title.RectAnchorType = svg.RECT_CENTER_MIDDLE
		title.TextAnchorType = svg.TEXT_ANCHOR_CENTER
		rect.RectAnchoredTexts = append(rect.RectAnchoredTexts, title)
	}
}

func (stager *Stager) drawEffortShapes(diagramFloss *DiagramFloss, layer *svg.Layer) {
	diagramFloss.map_Effort_Rect = make(map[*Effort]*svg.Rect)
	diagramFloss.map_SvgRect_EffortShape = make(map[*svg.Rect]*EffortShape)
	for _, effortShape := range diagramFloss.Effort_Shapes {
		if effortShape.IsHidden {
			continue
		}

		rect := new(svg.Rect)
		layer.Rects = append(layer.Rects, rect)
		diagramFloss.map_Effort_Rect[effortShape.Effort] = rect
		diagramFloss.map_SvgRect_EffortShape[rect] = effortShape

		if diagramFloss.owningSystem != nil {
			if systemRect, ok := diagramFloss.map_System_Rect[diagramFloss.owningSystem]; ok {
				rect.EnclosingRect = systemRect
				systemRect.Peers = append(systemRect.Peers, rect)
			}
		}

		rect.Name = effortShape.GetName()
		rect.X = effortShape.GetX()
		rect.Y = effortShape.GetY()
		rect.Width = effortShape.GetWidth()
		rect.Height = effortShape.GetHeight()

		// Crisp slate blue theme for Effort
		rect.Color = "#E3F2FD"
		rect.FillOpacity = 1.0
		rect.Stroke = "#1976D2"
		rect.StrokeWidth = 1.5
		rect.StrokeOpacity = 1.0
		rect.RX = 2

		rect.CanMoveHorizontaly = true
		rect.CanMoveVerticaly = true
		rect.CanHaveBottomHandle = true
		rect.CanHaveLeftHandle = true
		rect.CanHaveRightHandle = true
		rect.CanHaveTopHandle = true

		rect.OnSelect = func() {
			stager.probeForm.FillUpFormFromGongstruct(effortShape.Effort, GetPointerToGongstructName[*Effort]())
		}
		rect.OnMove = onMoveRectElement(stager, effortShape, true)
		rect.OnResize = onResizeRectElement(stager, effortShape)

		title := new(svg.RectAnchoredText)
		title.Name = effortShape.Effort.Name
		content := effortShape.Effort.Name
		if effortShape.Effort.Strength != 0 {
			content = fmt.Sprintf("%s (%.1f)", effortShape.Effort.Name, effortShape.Effort.Strength)
		}

		root := stager.getRootLibrary()
		nbPixPerChar := root.NbPixPerCharacter
		if nbPixPerChar <= 0 {
			nbPixPerChar = 8.0
		}
		if rect.Width > 0 {
			content = strutils.WrapStringPreservingNewlinesScaled(content, rect.Width, nbPixPerChar, 14.0, 16.0)
		}

		title.Content = content
		title.Color = "#0D47A1"
		title.FillOpacity = 1.0
		title.Stroke = "#0D47A1"
		title.StrokeWidth = 0
		title.StrokeOpacity = 1.0
		title.FontSize = "14px"
		title.FontWeight = "600"
		title.RectAnchorType = svg.RECT_CENTER_MIDDLE
		title.TextAnchorType = svg.TEXT_ANCHOR_CENTER
		rect.RectAnchoredTexts = append(rect.RectAnchoredTexts, title)
	}
}

func (stager *Stager) drawNoteShapes(diagramFloss *DiagramFloss, layer *svg.Layer) {
	diagramFloss.map_Note_Rect = make(map[*Note]*svg.Rect)
	diagramFloss.map_SvgRect_NoteShape = make(map[*svg.Rect]*NoteShape)
	for _, noteShape := range diagramFloss.Note_Shapes {
		if noteShape.IsHidden || noteShape.Note == nil {
			continue
		}

		rect := new(svg.Rect)
		layer.Rects = append(layer.Rects, rect)
		diagramFloss.map_Note_Rect[noteShape.Note] = rect
		diagramFloss.map_SvgRect_NoteShape[rect] = noteShape


		rect.Name = noteShape.GetName()
		rect.X = noteShape.GetX()
		rect.Y = noteShape.GetY()
		rect.Width = noteShape.GetWidth()
		rect.Height = noteShape.GetHeight()

		rect.Color = "#FFF9C4"
		rect.FillOpacity = 1.0
		rect.Stroke = "#FBC02D"
		rect.StrokeWidth = 1.5
		rect.StrokeOpacity = 1.0
		rect.RX = 2

		rect.CanMoveHorizontaly = true
		rect.CanMoveVerticaly = true
		rect.CanHaveBottomHandle = true
		rect.CanHaveLeftHandle = true
		rect.CanHaveRightHandle = true
		rect.CanHaveTopHandle = true

		rect.OnSelect = func() {
			stager.probeForm.FillUpFormFromGongstruct(noteShape.Note, GetPointerToGongstructName[*Note]())
		}
		rect.OnMove = onMoveRectElement(stager, noteShape, true)
		rect.OnResize = onResizeRectElement(stager, noteShape)

		title := new(svg.RectAnchoredText)
		title.Name = noteShape.Note.Name
		content := "📝 " + noteShape.Note.Name
		if noteShape.Note.Description != "" {
			content += "\n" + noteShape.Note.Description
		}

		root := stager.getRootLibrary()
		nbPixPerChar := root.NbPixPerCharacter
		if nbPixPerChar <= 0 {
			nbPixPerChar = 8.0
		}
		if rect.Width > 0 {
			content = strutils.WrapStringPreservingNewlinesScaled(content, rect.Width-20, nbPixPerChar, 14.0, 16.0)
		}

		title.Content = content
		title.Color = "#5D4037"
		title.FillOpacity = 1.0
		title.Stroke = "#5D4037"
		title.StrokeWidth = 0
		title.StrokeOpacity = 1.0
		title.FontSize = "13px"
		title.FontStyle = "italic"
		title.FontWeight = "normal"
		title.RectAnchorType = svg.RECT_TOP_LEFT
		title.TextAnchorType = svg.TEXT_ANCHOR_START
		title.X_Offset = 10
		title.Y_Offset = 18
		rect.RectAnchoredTexts = append(rect.RectAnchoredTexts, title)
	}
}

func (stager *Stager) drawNoteComplexityShapes(diagramFloss *DiagramFloss, layer *svg.Layer) {
	for _, shape := range diagramFloss.NoteComplexityShapes {
		if shape.IsHidden || shape.Note == nil || shape.Complexity == nil {
			continue
		}
		startRect := diagramFloss.map_Note_Rect[shape.Note]
		endRect := diagramFloss.map_Complexity_Rect[shape.Complexity]
		if startRect == nil || endRect == nil {
			continue
		}

		link := new(svg.Link)
		layer.Links = append(layer.Links, link)
		link.Name = startRect.Name + " to " + endRect.Name
		link.Start = startRect
		link.End = endRect
		link.StartAnchorType = svg.ANCHOR_CENTER
		link.EndAnchorType = svg.ANCHOR_CENTER
		link.Type = svg.LINK_TYPE_LINE_WITH_CONTROL_POINTS
		link.HasEndArrow = false
		link.Stroke = "#FFA000"
		link.StrokeWidth = 1.5
		link.StrokeDashArray = "5 5"
		link.StrokeOpacity = 1.0
	}
}

func (stager *Stager) drawNotePerformanceShapes(diagramFloss *DiagramFloss, layer *svg.Layer) {
	for _, shape := range diagramFloss.NotePerformanceShapes {
		if shape.IsHidden || shape.Note == nil || shape.Performance == nil {
			continue
		}
		startRect := diagramFloss.map_Note_Rect[shape.Note]
		endRect := diagramFloss.map_Performance_Rect[shape.Performance]
		if startRect == nil || endRect == nil {
			continue
		}

		link := new(svg.Link)
		layer.Links = append(layer.Links, link)
		link.Name = startRect.Name + " to " + endRect.Name
		link.Start = startRect
		link.End = endRect
		link.StartAnchorType = svg.ANCHOR_CENTER
		link.EndAnchorType = svg.ANCHOR_CENTER
		link.Type = svg.LINK_TYPE_LINE_WITH_CONTROL_POINTS
		link.HasEndArrow = false
		link.Stroke = "#2E7D32"
		link.StrokeWidth = 1.5
		link.StrokeDashArray = "5 5"
		link.StrokeOpacity = 1.0
	}
}

func (stager *Stager) drawNoteEffortShapes(diagramFloss *DiagramFloss, layer *svg.Layer) {
	for _, shape := range diagramFloss.NoteEffortShapes {
		if shape.IsHidden || shape.Note == nil || shape.Effort == nil {
			continue
		}
		startRect := diagramFloss.map_Note_Rect[shape.Note]
		endRect := diagramFloss.map_Effort_Rect[shape.Effort]
		if startRect == nil || endRect == nil {
			continue
		}

		link := new(svg.Link)
		layer.Links = append(layer.Links, link)
		link.Name = startRect.Name + " to " + endRect.Name
		link.Start = startRect
		link.End = endRect
		link.StartAnchorType = svg.ANCHOR_CENTER
		link.EndAnchorType = svg.ANCHOR_CENTER
		link.Type = svg.LINK_TYPE_LINE_WITH_CONTROL_POINTS
		link.HasEndArrow = false
		link.Stroke = "#1976D2"
		link.StrokeWidth = 1.5
		link.StrokeDashArray = "5 5"
		link.StrokeOpacity = 1.0
	}
}



