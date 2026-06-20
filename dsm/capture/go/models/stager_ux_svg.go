package models

import (
	"log"

	"github.com/fullstack-lang/gong/lib/strutils"
	svg "github.com/fullstack-lang/gong/lib/svg/go/models"
)

func (stager *Stager) svg() {
	log.Println("svg")
	stager.svgStage.Reset()
	stage := stager.stage
	_ = stage

	var diagram *Diagram
	{
		for diagram_ := range *GetGongstructInstancesSet[Diagram](stager.stage) {
			if diagram_.IsChecked {
				diagram = diagram_
			}
		}
	}

	if diagram == nil {
		stager.svgStage.Commit()
		return
	}

	svgObject := stager.generateSvgObject(diagram)

	svg.StageBranch(stager.svgStage, svgObject)
	stager.svgObject = svgObject
	stager.svgObject.OnUpdate = stager.onUpdateSVG

	stager.svgStage.Commit()
}

// generateSvgObject creates and returns a new svg.SVG object representing the given diagram.
// It maps all visible domain shapes (Deliverables, Tasks, Notes, Resources) and their associations
// to SVG elements (Rects, Links, Paths) on a single layer. It also populates the diagram's
// internal maps to link abstract elements with their visual SVG counterparts.
func (stager *Stager) generateSvgObject(diagram *Diagram) *svg.SVG {
	root := stager.GetRootLibrary()

	svgStage := stager.svgStage

	stager.svgObject = (&svg.SVG{Name: `SVG`})
	stager.diagram = diagram
	stager.svgObject.OnUpdate = stager.onUpdateSVG

	stager.svgObject.OverrideWidth = true
	stager.svgObject.OverriddenWidth = diagram.Width
	stager.svgObject.OverrideHeight = true
	stager.svgObject.OverriddenHeight = diagram.Height

	diagram.map_Deliverable_Rect = make(map[*Deliverable]*svg.Rect)
	diagram.map_Task_Rect = make(map[*Concern]*svg.Rect)
	diagram.map_Note_Rect = make(map[*Note]*svg.Rect)
	diagram.map_Stakeholder_Rect = make(map[*Stakeholder]*svg.Rect)
	diagram.map_Requirement_Rect = make(map[*Requirement]*svg.Rect)
	diagram.map_Concept_Rect = make(map[*Concept]*svg.Rect)
	diagram.map_Diagram_Rect = make(map[*Diagram]*svg.Rect)

	diagram.map_SvgRect_DeliverableShape = make(map[*svg.Rect]*DeliverableShape)
	diagram.map_SvgRect_ConcernShape = make(map[*svg.Rect]*ConcernShape)
	diagram.map_SvgRect_NoteShape = make(map[*svg.Rect]*NoteShape)
	diagram.map_SvgRect_StakeholderShape = make(map[*svg.Rect]*StakeholderShape)
	diagram.map_SvgRect_RequirementShape = make(map[*svg.Rect]*RequirementShape)
	diagram.map_SvgRect_ConceptShape = make(map[*svg.Rect]*ConceptShape)
	diagram.map_SvgRect_DiagramShape = make(map[*svg.Rect]*DiagramShape)

	// // to implement association between abstract elements by mouse drag
	// svgImpl := &svgProxy{
	// 	stager:  stager,
	// 	svg_:    stager.svgObject,
	// 	diagram: diagram,
	// }
	// stager.svgObject.Impl = svgImpl

	stager.svgObject.Name = diagram.Name
	stager.svgObject.IsEditable = diagram.IsEditable()

	layer := (&svg.Layer{Name: "Layer 1"})
	stager.svgObject.Layers = append(stager.svgObject.Layers, layer)

	backgroundRect := &svg.Rect{
		Name:   diagram.Name + " background",
		X:      0.0,
		Y:      0.0,
		Width:  diagram.Width,
		Height: diagram.Height,
		Presentation: svg.Presentation{
			Color:       "transparent",
			FillOpacity: 0.0,
		},
		OnUpdate: func(frontRect *svg.Rect) {
			diagram.IsEditable_ = !diagram.IsEditable_
			stager.stage.Commit()
		},
	}
	layer.Rects = append(layer.Rects, backgroundRect)

	for _, deliverableShape := range diagram.Deliverable_Shapes {
		if deliverableShape.IsHidden {
			continue
		}

		rect := svgRect(
			stager,
			diagram,
			deliverableShape,
			layer)
		rect.RX = 3
		rect.Color = "#F3E5F5"
		rect.FillOpacity = 1.0
		rect.Stroke = "#CE93D8"
		rect.StrokeWidth = 1.5

		diagram.map_Deliverable_Rect[deliverableShape.Deliverable] = rect
		diagram.map_SvgRect_DeliverableShape[rect] = deliverableShape

		deliverableLogo := new(svg.RectAnchoredPath)
		deliverableLogo.Name = "Document"
		deliverableLogo.Stroke = svg.Black.ToString()
		deliverableLogo.StrokeWidth = 1
		deliverableLogo.StrokeOpacity = 1
		deliverableLogo.Color = svg.Black.ToString()
		deliverableLogo.FillOpacity = 0.2
		deliverableLogo.ScalePropotionnally = true
		deliverableLogo.AppliedScaling = 1.0

		// Material 'description' document icon path
		deliverableLogo.Definition = "M14 2H6c-1.1 0-1.99.9-1.99 2L4 20c0 1.1.89 2 1.99 2H18c1.1 0 2-.9 2-2V8l-6-6zm2 16H8v-2h8v2zm0-4H8v-2h8v2zm-3-5V3.5L18.5 9H13z"

		distanceFromBorder := 5.0
		iconWidth := 20.0
		padding := 10.0

		deliverableLogo.X_Offset = distanceFromBorder
		deliverableLogo.Y_Offset = distanceFromBorder
		deliverableLogo.RectAnchorType = svg.RECT_TOP_LEFT
		rect.RectAnchoredPaths = append(rect.RectAnchoredPaths, deliverableLogo)

		if len(rect.RectAnchoredTexts) > 0 {
			title := rect.RectAnchoredTexts[0]
			title.Color = "#333333"
			title.FontWeight = "500"
			title.FontSize = "16px"
			title.FontFamily = "sans-serif"
			if rect.Width > (distanceFromBorder + iconWidth + padding) {
				title.Content = strutils.WrapStringPreservingNewlines(title.Content, int((rect.Width-(distanceFromBorder+iconWidth+padding))/root.NbPixPerCharacter))
			}
			title.X_Offset = (distanceFromBorder + iconWidth) / 2.0
		}
	}

	for _, deliverableCompositionShape := range diagram.DeliverableComposition_Shapes {
		if deliverableCompositionShape.GetIsHidden() {
			continue
		}
		_ = deliverableCompositionShape
		subDeliverable := deliverableCompositionShape.Deliverable
		parentDeliverable := subDeliverable.parentDeliverable

		if subDeliverable == nil || parentDeliverable == nil {
			log.Panic("There should be a subDeliverable and a parentDeliverable")
		}

		startRect := diagram.map_Deliverable_Rect[parentDeliverable]
		endRect := diagram.map_Deliverable_Rect[subDeliverable]

		if startRect == nil || endRect == nil {
			continue
		}

		link := svgAssociationLink(
			stager,
			startRect, endRect,
			deliverableCompositionShape,
			// when one clicks on the link, this is the form of the parent deliverable
			parentDeliverable,
			layer,
			false)
		configureLinkWithControlPoints(stager, link, &deliverableCompositionShape.ControlPointShapes, startRect.Name)
	}

	for _, concernShape := range diagram.Concern_Shapes {
		if concernShape.IsHidden {
			continue
		}

		rect := svgRect(
			stager,
			diagram,
			concernShape,
			layer)

		rect.RX = 3
		rect.Color = "#E8F5E9"
		rect.FillOpacity = 1.0
		rect.Stroke = "#81C784"
		rect.StrokeWidth = 1.5

		if len(rect.RectAnchoredTexts) > 0 {
			title := rect.RectAnchoredTexts[0]
			title.Color = "#333333"
			title.FontWeight = "500"
			title.FontSize = "16px"
			title.FontFamily = "sans-serif"
		}

		diagram.map_Task_Rect[concernShape.Concern] = rect
		diagram.map_SvgRect_ConcernShape[rect] = concernShape

		distanceFromBorder := 5.0
		iconWidth := 20.0
		padding := 10.0

		if concernShape.Concern.IsWithCompletion {
			distanceFromBorder = 10.0
			iconWidth = 25.0
			rectAnchoredPath := new(svg.RectAnchoredPath)
			rect.IsScalingProportionally = false
			rect.RectAnchoredPaths = append(rect.RectAnchoredPaths, rectAnchoredPath)

			switch concernShape.Concern.Completion {
			case PERCENT_100:
				rectAnchoredPath.Definition = "M 5.56,0 c -3.08,0 -5.56,2.48 -5.56,5.56 v 42.97 c 0,3.08 2.48,5.56 5.56,5.56 h 13.89 c 3.08,0 5.56,-2.48 5.56,-5.56 v -42.97 c 0,-3.08 -2.48,-5.56 -5.56,-5.56 z M 5.56,2.70 h 13.89 c 1.54,0 2.78,1.24 2.78,2.78 v 43.06 c 0,1.54 -1.24,2.78 -2.78,2.78 H 5.56 c -1.54,0 -2.78,-1.24 -2.78,-2.78 v -43.06 c 0,-1.54 1.24,-2.78 2.78,-2.78 z m 1.38,4.16 v 6.85 h 11.11 v -6.85 z m 0,11.35 v 6.68 h 11.11 v -6.68 z m 0,11.08 v 6.60 h 11.11 v -6.60 z m 0,11.05 v 6.67 h 11.11 v -6.67 z"
			case PERCENT_075:
				rectAnchoredPath.Definition = "M 5.56,0 c -3.08,0 -5.56,2.48 -5.56,5.56 v 42.97 c 0,3.08 2.48,5.56 5.56,5.56 h 13.89 c 3.08,0 5.56,-2.48 5.56,-5.56 v -42.97 c 0,-3.08 -2.48,-5.56 -5.56,-5.56 z M 5.56,2.70 h 13.89 c 1.54,0 2.78,1.24 2.78,2.78 v 43.06 c 0,1.54 -1.24,2.78 -2.78,2.78 H 5.56 c -1.54,0 -2.78,-1.24 -2.78,-2.78 v -43.06 c 0,-1.54 1.24,-2.78 2.78,-2.78 z m 1.38,15.51 v 6.68 h 11.11 v -6.68 z m 0,11.09 v 6.60 h 11.11 v -6.60 z m 0,11.04 v 6.67 h 11.11 v -6.67 z"
			case PERCENT_050:
				rectAnchoredPath.Definition = "M 5.56,0 c -3.08,0 -5.56,2.48 -5.56,5.56 v 42.97 c 0,3.08 2.48,5.56 5.56,5.56 h 13.89 c 3.08,0 5.56,-2.48 5.56,-5.56 v -42.97 c 0,-3.08 -2.48,-5.56 -5.56,-5.56 z M 5.56,2.70 h 13.89 c 1.54,0 2.78,1.24 2.78,2.78 v 43.06 c 0,1.54 -1.24,2.78 -2.78,2.78 H 5.56 c -1.54,0 -2.78,-1.24 -2.78,-2.78 v -43.06 c 0,-1.54 1.24,-2.78 2.78,-2.78 z m 1.39,26.6 v 6.6 h 11.11 v -6.6 z m 0,11.05 v 6.67 h 11.11 v -6.67 z"
			case PERCENT_025:
				rectAnchoredPath.Definition = "M 5.56,0 c -3.08,0 -5.56,2.48 -5.56,5.56 v 42.97 c 0,3.08 2.48,5.56 5.56,5.56 h 13.89 c 3.08,0 5.56,-2.48 5.56,-5.56 v -42.97 c 0,-3.08 -2.48,-5.56 -5.56,-5.56 z M 5.56,2.70 h 13.89 c 1.54,0 2.78,1.24 2.78,2.78 v 43.06 c 0,1.54 -1.24,2.78 -2.78,2.78 H 5.56 c -1.54,0 -2.78,-1.24 -2.78,-2.78 v -43.06 c 0,-1.54 1.24,-2.78 2.78,-2.78 z M 6.94,40.34 v 6.67 h 11.11 v -6.67 z"
			case PERCENT_000:
				rectAnchoredPath.Definition = "M 4.99,0 c -2.81,0.28 -4.99,2.65 -4.99,5.53 v 42.95 c 0,3.08 2.51,5.53 5.59,5.53 h 13.88 c 3.08,0 5.53,-2.45 5.53,-5.53 v -42.95 c 0,-3.08 -2.46,-5.53 -5.53,-5.53 h -13.88 c -0.19,0 -0.41,-0.02 -0.60,0 z m 0.60,2.66 h 13.88 c 1.54,0 2.77,1.23 2.77,2.77 v 43.06 c 0,1.54 -1.23,2.77 -2.77,2.77 h -13.88 c -1.54,0 -2.82,-1.23 -2.82,-2.77 v -43.06 c 0,-1.54 1.28,-2.77 2.82,-2.77 z"
			}

			rectAnchoredPath.Name = "Percentage"

			// rectAnchoredPath.ScalePropotionnally = true
			// rectAnchoredPath.AppliedScaling = 0.08

			rectAnchoredPath.Stroke = svg.Lightgray.ToString()
			rectAnchoredPath.StrokeWidth = 2
			rectAnchoredPath.StrokeOpacity = 1

			rectAnchoredPath.Color = svg.Lightgray.ToString()
			rectAnchoredPath.FillOpacity = 0.8

			rectAnchoredPath.X_Offset = distanceFromBorder
			rectAnchoredPath.Y_Offset = distanceFromBorder

			rectAnchoredPath.RectAnchorType = svg.RECT_TOP_LEFT

			// shift the text on the right
			title := rect.RectAnchoredTexts[0]
			if rect.Width > (distanceFromBorder + iconWidth + padding) {
				title.Content = strutils.WrapStringPreservingNewlines(title.Content, int((rect.Width-(distanceFromBorder+iconWidth+padding))/root.NbPixPerCharacter))
			}
			title.X_Offset = (distanceFromBorder + iconWidth) / 2.0
		} else {
			concernLogo := new(svg.RectAnchoredPath)
			concernLogo.Name = "Gear"
			concernLogo.Stroke = svg.Black.ToString()
			concernLogo.StrokeWidth = 1
			concernLogo.StrokeOpacity = 1
			concernLogo.Color = svg.Black.ToString()
			concernLogo.FillOpacity = 0.2
			concernLogo.ScalePropotionnally = true
			concernLogo.AppliedScaling = 1.0

			// Material 'settings' icon path
			concernLogo.Definition = "M19.14 12.94c.04-.3.06-.61.06-.94 0-.32-.02-.64-.06-.94l2.03-1.58c.18-.14.23-.41.12-.61l-1.92-3.32c-.12-.22-.37-.29-.59-.22l-2.39.96c-.5-.38-1.03-.7-1.62-.94l-.36-2.54c-.04-.24-.24-.41-.48-.41h-3.84c-.24 0-.43.17-.47.41l-.36 2.54c-.59.24-1.13.57-1.62.94l-2.39-.96c-.22-.08-.47 0-.59.22L2.74 8.87c-.12.21-.08.47.12.61l2.03 1.58c-.05.3-.09.63-.09.94s.02.64.06.94l-2.03 1.58c-.18.14-.23.41-.12.61l1.92 3.32c.12.22.37.29.59.22l2.39-.96c.5.38 1.03.7 1.62.94l.36 2.54c.05.24.24.41.48.41h3.84c.24 0 .44-.17.47-.41l.36-2.54c.59-.24 1.13-.56 1.62-.94l2.39.96c.22.08.47 0 .59-.22l1.92-3.32c.12-.22.07-.47-.12-.61l-2.01-1.58zM12 15.6c-1.98 0-3.6-1.62-3.6-3.6s1.62-3.6 3.6-3.6 3.6 1.62 3.6 3.6-1.62 3.6-3.6 3.6z"

			concernLogo.X_Offset = distanceFromBorder
			concernLogo.Y_Offset = distanceFromBorder
			concernLogo.RectAnchorType = svg.RECT_TOP_LEFT
			rect.RectAnchoredPaths = append(rect.RectAnchoredPaths, concernLogo)

			if len(rect.RectAnchoredTexts) > 0 {
				title := rect.RectAnchoredTexts[0]
				if rect.Width > (distanceFromBorder + iconWidth + padding) {
					title.Content = strutils.WrapStringPreservingNewlines(title.Content, int((rect.Width-(distanceFromBorder+iconWidth+padding))/root.NbPixPerCharacter))
				}
				title.X_Offset = (distanceFromBorder + iconWidth) / 2.0
			}
		}

		if concernShape.Concern.IDAirbus != "" {

			idAirbusText := svg.RectAnchoredText{
				Name:           concernShape.Concern.IDAirbus,
				Content:        concernShape.Concern.IDAirbus,
				RectAnchorType: svg.RECT_TOP_LEFT,
				TextAnchorType: svg.TEXT_ANCHOR_START,
				X_Offset:       10,
				Y_Offset:       -10,
			}
			applyStyleToRectAnchoredText(&idAirbusText)
			rect.RectAnchoredTexts = append(rect.RectAnchoredTexts, &idAirbusText)
		}
	}

	for _, concernCompositionShape := range diagram.ConcernComposition_Shapes {
		if concernCompositionShape.GetIsHidden() {
			continue
		}
		_ = concernCompositionShape
		subTask := concernCompositionShape.Concern
		parentTask := subTask.parentConcern

		if subTask == nil || parentTask == nil {
			log.Panic("There should be a subTask and a parentTask")
		}

		startRect := diagram.map_Task_Rect[parentTask]
		endRect := diagram.map_Task_Rect[subTask]

		if startRect == nil || endRect == nil {
			continue
		}

		link := svgAssociationLink(
			stager,
			startRect, endRect,
			concernCompositionShape,
			parentTask,
			layer,
			false)
		configureLinkWithControlPoints(stager, link, &concernCompositionShape.ControlPointShapes, startRect.Name)
	}

	for _, taskInputShape := range diagram.ConcernInputShapes {
		if taskInputShape.GetIsHidden() {
			continue
		}
		concern := taskInputShape.Concern
		deliverable := taskInputShape.Deliverable

		if concern == nil || deliverable == nil {
			log.Panic("There should be a task and a deliverable")
		}

		startRect := diagram.map_Deliverable_Rect[deliverable]
		endRect := diagram.map_Task_Rect[concern]

		if startRect == nil || endRect == nil {
			continue
		}

		link := svgAssociationLink(
			stager,
			startRect, endRect,
			taskInputShape,
			concern,
			layer,
			false,
		)
		link.Stroke = "#555555"
		link.StrokeWidth = 1.5
		link.HasEndArrow = true
		configureLinkWithControlPoints(stager, link, &taskInputShape.ControlPointShapes, startRect.Name)
	}

	for _, taskOutputShape := range diagram.ConcernOutputShapes {
		if taskOutputShape.GetIsHidden() {
			continue
		}
		task := taskOutputShape.Concern
		deliverable := taskOutputShape.Deliverable

		if task == nil || deliverable == nil {
			log.Panic("There should be a task and a deliverable")
		}

		startRect := diagram.map_Task_Rect[task]
		endRect := diagram.map_Deliverable_Rect[deliverable]

		if startRect == nil || endRect == nil {
			continue
		}

		link := svgAssociationLink(
			stager,
			startRect, endRect,
			taskOutputShape,
			task,
			layer,
			false,
		)
		link.Stroke = "#555555"
		link.StrokeWidth = 1.5
		link.HasEndArrow = true
		configureLinkWithControlPoints(stager, link, &taskOutputShape.ControlPointShapes, startRect.Name)
	}

	for _, noteShape := range diagram.Note_Shapes {
		if noteShape.IsHidden {
			continue
		}

		note := noteShape.Note

		if note == nil {
			log.Println("There should be a note")
			continue
		}

		rect := svgRect(
			stager,
			diagram,
			noteShape,
			layer)
		rect.RX = 6
		rect.Color = "#FFF9C4"
		rect.FillOpacity = 1.0
		rect.Stroke = "#FBC02D"
		rect.StrokeWidth = 1.0

		if len(rect.RectAnchoredTexts) > 0 {
			title := rect.RectAnchoredTexts[0]
			title.Color = "#333333"
			title.FontWeight = "500"
			title.FontSize = "14px"
			title.FontFamily = "sans-serif"
		}

		diagram.map_Note_Rect[note] = rect
		diagram.map_SvgRect_NoteShape[rect] = noteShape

		penLogo := new(svg.RectAnchoredPath)
		penLogo.Stroke = svg.Black.ToString()
		penLogo.StrokeWidth = 2
		penLogo.StrokeOpacity = 1
		penLogo.ScalePropotionnally = true
		penLogo.AppliedScaling = 1

		penLogo.Definition = "M 5 16 L 9 20 L 20 9 L 25 0 L 16 5 Z"
		distanceFromBorder := 5.0
		iconWidth := 20.0
		padding := 10.0

		penLogo.X_Offset = distanceFromBorder
		penLogo.Y_Offset = distanceFromBorder
		penLogo.RectAnchorType = svg.RECT_TOP_LEFT
		rect.RectAnchoredPaths = append(rect.RectAnchoredPaths, penLogo)

		if len(rect.RectAnchoredTexts) > 0 {
			title := rect.RectAnchoredTexts[0]
			if rect.Width > (distanceFromBorder + iconWidth + padding) {
				title.Content = strutils.WrapStringPreservingNewlines(title.Content, int((rect.Width-(distanceFromBorder+iconWidth+padding))/root.NbPixPerCharacter))
			}
			title.X_Offset = (distanceFromBorder + iconWidth) / 2.0
		}
	}

	for _, noteDeliverableShape := range diagram.NoteDeliverableShapes {
		note := noteDeliverableShape.Note
		deliverable := noteDeliverableShape.Deliverable

		startRect := diagram.map_Note_Rect[note]
		endRect := diagram.map_Deliverable_Rect[deliverable]

		if startRect == nil || endRect == nil {
			continue
		}

		link := svgAssociationLink(
			stager,
			startRect, endRect,
			noteDeliverableShape,
			note,
			layer,
			true,
		)
		link.HasEndArrow = false
		configureLinkWithControlPoints(stager, link, &noteDeliverableShape.ControlPointShapes, startRect.Name)
	}

	for _, noteTaskShape := range diagram.NoteTaskShapes {
		note := noteTaskShape.Note
		task := noteTaskShape.Task
		startRect := diagram.map_Note_Rect[note]
		endRect := diagram.map_Task_Rect[task]

		if startRect == nil || endRect == nil {
			continue
		}

		link := svgAssociationLink(
			stager,
			startRect, endRect,
			noteTaskShape,
			note,
			layer,
			true,
		)
		link.HasEndArrow = false
		configureLinkWithControlPoints(stager, link, &noteTaskShape.ControlPointShapes, startRect.Name)
	}

	for _, s := range diagram.Stakeholder_Shapes {
		if s.IsHidden {
			continue
		}

		rect := svgRect(
			stager,
			diagram,
			s,
			layer)

		rect.RX = 3
		rect.Color = "#E3F2FD"
		rect.FillOpacity = 1.0
		rect.Stroke = "#90CAF9"
		rect.StrokeWidth = 1.5

		if len(rect.RectAnchoredTexts) > 0 {
			title := rect.RectAnchoredTexts[0]
			title.Color = "#333333"
			title.FontWeight = "500"
			title.FontSize = "16px"
			title.FontFamily = "sans-serif"
		}

		diagram.map_Stakeholder_Rect[s.Stakeholder] = rect
		diagram.map_SvgRect_StakeholderShape[rect] = s
		{
			rectAnchoredPath := new(svg.RectAnchoredPath)
			rect.IsScalingProportionally = false
			rect.RectAnchoredPaths = append(rect.RectAnchoredPaths, rectAnchoredPath)

			rectAnchoredPath.Definition = "M 30, 11 a 10, 10 0 1, 1 -20, 0 a 10, 10 0 1, 1 20, 0 M 1, 49 h 38 v -9 c 0 -6.6 -5.4 -12 -12 -12 H 13 c -6.6 0 -12 5.4 -12 12 z"

			rectAnchoredPath.Name = "Person"

			rectAnchoredPath.Stroke = svg.Lightgray.ToString()
			rectAnchoredPath.StrokeWidth = 4
			rectAnchoredPath.StrokeOpacity = 1

			rectAnchoredPath.FillOpacity = 0.0

			distanceFromBorder := 10.0
			iconWidth := 40.0

			rectAnchoredPath.X_Offset = distanceFromBorder
			rectAnchoredPath.Y_Offset = distanceFromBorder

			rectAnchoredPath.RectAnchorType = svg.RECT_TOP_LEFT

			// shift the text on the right
			padding := 10.0
			title := rect.RectAnchoredTexts[0]
			if rect.Width > (distanceFromBorder + iconWidth + padding) {
				title.Content = strutils.WrapStringPreservingNewlines(title.Content, int((rect.Width-(distanceFromBorder+iconWidth+padding))/root.NbPixPerCharacter))
			}
			title.X_Offset = (distanceFromBorder + iconWidth) / 2.0
		}

		if s.Stakeholder.IDAirbus != "" {

			idAirbusText := svg.RectAnchoredText{
				Name:           s.Stakeholder.IDAirbus,
				Content:        s.Stakeholder.IDAirbus,
				RectAnchorType: svg.RECT_TOP_LEFT,
				TextAnchorType: svg.TEXT_ANCHOR_START,
				X_Offset:       10,
				Y_Offset:       -10,
			}
			applyStyleToRectAnchoredText(&idAirbusText)
			rect.RectAnchoredTexts = append(rect.RectAnchoredTexts, &idAirbusText)
		}
	}

	for _, noteResourceShape := range diagram.NoteResourceShapes {
		note := noteResourceShape.Note
		resource := noteResourceShape.Stakeholder
		startRect := diagram.map_Note_Rect[note]
		endRect := diagram.map_Stakeholder_Rect[resource]
		if startRect == nil || endRect == nil {
			continue
		}
		link := svgAssociationLink(
			stager,
			startRect, endRect,
			noteResourceShape,
			note,
			layer,
			true,
		)
		link.HasEndArrow = false
		configureLinkWithControlPoints(stager, link, &noteResourceShape.ControlPointShapes, startRect.Name)
	}

	for _, resourceCompositionShape := range diagram.ResourceComposition_Shapes {
		if resourceCompositionShape.GetIsHidden() {
			continue
		}
		subResource := resourceCompositionShape.Stakeholder
		parentResource := subResource.parentStakeholder
		if subResource == nil || parentResource == nil {
			log.Panic("There should be a subResource and a parentResource")
		}
		startRect := diagram.map_Stakeholder_Rect[parentResource]
		endRect := diagram.map_Stakeholder_Rect[subResource]
		if startRect == nil || endRect == nil {
			continue
		}
		link := svgAssociationLink(
			stager,
			startRect, endRect,
			resourceCompositionShape,
			parentResource,
			layer,
			false,
		)
		configureLinkWithControlPoints(stager, link, &resourceCompositionShape.ControlPointShapes, startRect.Name)
	}

	for _, resourceTaskShape := range diagram.StakeholderConcernShapes {
		if resourceTaskShape.GetIsHidden() {
			continue
		}
		resource := resourceTaskShape.Stakeholder
		task := resourceTaskShape.Concern
		if resource == nil || task == nil {
			log.Panic("There should be a resource and a task")
		}
		startRect := diagram.map_Stakeholder_Rect[resource]
		endRect := diagram.map_Task_Rect[task]
		if startRect == nil || endRect == nil {
			continue
		}
		link := svgAssociationLink(
			stager,
			startRect, endRect,
			resourceTaskShape,
			resource,
			layer,
			false,
		)
		link.Stroke = "#555555"
		link.StrokeWidth = 1.5
		link.HasEndArrow = true
		configureLinkWithControlPoints(stager, link, &resourceTaskShape.ControlPointShapes, startRect.Name)
	}

	for _, reqShape := range diagram.Requirement_Shapes {
		if reqShape.IsHidden {
			continue
		}

		rect := svgRect(
			stager,
			diagram,
			reqShape,
			layer)
		rect.RX = 0 // Sharp corners for Requirements
		rect.Color = "#F5F5F5"
		rect.FillOpacity = 1.0
		rect.Stroke = "#E0E0E0"
		rect.StrokeWidth = 1.5

		if len(rect.RectAnchoredTexts) > 0 {
			title := rect.RectAnchoredTexts[0]
			title.Color = "#333333"
			title.FontWeight = "500"
			title.FontSize = "16px"
			title.FontFamily = "sans-serif"
		}

		diagram.map_Requirement_Rect[reqShape.Requirement] = rect
		diagram.map_SvgRect_RequirementShape[rect] = reqShape

		reqLogo := new(svg.RectAnchoredPath)
		reqLogo.Name = "Assignment"
		reqLogo.Stroke = svg.Black.ToString()
		reqLogo.StrokeWidth = 1
		reqLogo.StrokeOpacity = 1
		reqLogo.Color = svg.Black.ToString()
		reqLogo.FillOpacity = 0.2
		reqLogo.ScalePropotionnally = true
		reqLogo.AppliedScaling = 1.0

		// Material 'assignment' icon path
		reqLogo.Definition = "M19 3h-4.18C14.4 1.84 13.3 1 12 1c-1.3 0-2.4.84-2.82 2H5c-1.1 0-2 .9-2 2v14c0 1.1.9 2 2 2h14c1.1 0 2-.9 2-2V5c0-1.1-.9-2-2-2zm-7 0c.55 0 1 .45 1 1s-.45 1-1 1-1-.45-1-1 .45-1 1-1zm2 14H7v-2h7v2zm3-4H7v-2h10v2zm0-4H7V7h10v2z"

		distanceFromBorder := 5.0
		iconWidth := 20.0

		reqLogo.X_Offset = distanceFromBorder
		reqLogo.Y_Offset = distanceFromBorder
		reqLogo.RectAnchorType = svg.RECT_TOP_LEFT
		rect.RectAnchoredPaths = append(rect.RectAnchoredPaths, reqLogo)

		if len(rect.RectAnchoredTexts) > 0 {
			padding := 10.0
			title := rect.RectAnchoredTexts[0]
			if rect.Width > (distanceFromBorder + iconWidth + padding) {
				title.Content = strutils.WrapStringPreservingNewlines(title.Content, int((rect.Width-(distanceFromBorder+iconWidth+padding))/root.NbPixPerCharacter))
			}
			title.X_Offset = (distanceFromBorder + iconWidth) / 2.0
		}
	}

	for _, diagramShape := range diagram.Diagram_Shapes {
		if diagramShape.IsHidden {
			continue
		}

		rect := svgRect(
			stager,
			diagram,
			diagramShape,
			layer)
		rect.RX = 5
		rect.Color = "#E1F5FE"
		rect.FillOpacity = 1.0
		rect.Stroke = "#81D4FA"
		rect.StrokeWidth = 1.5

		if len(rect.RectAnchoredTexts) > 0 {
			title := rect.RectAnchoredTexts[0]
			title.Color = "#333333"
			title.FontWeight = "500"
			title.FontSize = "16px"
			title.FontFamily = "sans-serif"
		}

		diagram.map_Diagram_Rect[diagramShape.Diagram] = rect
		diagram.map_SvgRect_DiagramShape[rect] = diagramShape

		rect.OnSelect = func() {
			for diagram_ := range *GetGongstructInstancesSet[Diagram](stager.stage) {
				diagram_.IsChecked = false
			}
			diagramShape.Diagram.IsChecked = true
			stager.stage.Commit()
		}
		rect.OnMove = func(x, y float64) {
			diagramShape.SetX(x)
			diagramShape.SetY(y)

			stager.stage.CommitWithSuspendedCallbacks()
			stager.ux_tree()
		}
		rect.OnResize = func(x, y, width, height float64) {
			diagramShape.SetX(x)
			diagramShape.SetY(y)
			diagramShape.SetWidth(width)
			diagramShape.SetHeight(height)

			stager.stage.Commit()
		}

		diagramLogo := new(svg.RectAnchoredPath)
		diagramLogo.Name = "Folder"
		diagramLogo.Stroke = svg.Black.ToString()
		diagramLogo.StrokeWidth = 1
		diagramLogo.StrokeOpacity = 1
		diagramLogo.Color = svg.Black.ToString()
		diagramLogo.FillOpacity = 0.2
		diagramLogo.ScalePropotionnally = true
		diagramLogo.AppliedScaling = 1.0

		// Material 'folder' outline icon path
		diagramLogo.Definition = "M10 4H4c-1.1 0-1.99.9-1.99 2L2 18c0 1.1.9 2 2 2h16c1.1 0 2-.9 2-2V8c0-1.1-.9-2-2-2h-8l-2-2z"

		distanceFromBorder := 5.0
		iconWidth := 20.0

		diagramLogo.X_Offset = distanceFromBorder
		diagramLogo.Y_Offset = distanceFromBorder
		diagramLogo.RectAnchorType = svg.RECT_TOP_LEFT
		rect.RectAnchoredPaths = append(rect.RectAnchoredPaths, diagramLogo)

		if len(rect.RectAnchoredTexts) > 0 {
			padding := 10.0
			title := rect.RectAnchoredTexts[0]
			if rect.Width > (distanceFromBorder + iconWidth + padding) {
				title.Content = strutils.WrapStringPreservingNewlines(title.Content, int((rect.Width-(distanceFromBorder+iconWidth+padding))/root.NbPixPerCharacter))
			}
			title.X_Offset = (distanceFromBorder + iconWidth) / 2.0
		}
	}

	for _, conceptShape := range diagram.Concept_Shapes {
		if conceptShape.IsHidden {
			continue
		}

		rect := svgRect(
			stager,
			diagram,
			conceptShape,
			layer)
		rect.RX = 15                  // Very rounded corners for Concepts
		rect.StrokeDashArray = "5, 5" // Dashed outline for concepts
		rect.Color = "#FFF3E0"
		rect.FillOpacity = 1.0
		rect.Stroke = "#FFB74D"
		rect.StrokeWidth = 1.5

		if len(rect.RectAnchoredTexts) > 0 {
			title := rect.RectAnchoredTexts[0]
			title.Color = "#333333"
			title.FontWeight = "500"
			title.FontSize = "16px"
			title.FontFamily = "sans-serif"
		}

		diagram.map_Concept_Rect[conceptShape.Concept] = rect
		diagram.map_SvgRect_ConceptShape[rect] = conceptShape

		conceptLogo := new(svg.RectAnchoredPath)
		conceptLogo.Name = "Lightbulb"
		conceptLogo.Stroke = svg.Black.ToString()
		conceptLogo.StrokeWidth = 1
		conceptLogo.StrokeOpacity = 1
		conceptLogo.Color = svg.Black.ToString()
		conceptLogo.FillOpacity = 0.2
		conceptLogo.ScalePropotionnally = true
		conceptLogo.AppliedScaling = 1.0

		// Material 'lightbulb' outline icon path
		conceptLogo.Definition = "M9 21c0 .55.45 1 1 1h4c.55 0 1-.45 1-1v-1H9v1zm3-19C8.14 2 5 5.14 5 9c0 2.38 1.19 4.47 3 5.74V17c0 .55.45 1 1 1h6c.55 0 1-.45 1-1v-2.26c1.81-1.27 3-3.36 3-5.74 0-3.86-3.14-7-7-7zm2.85 11.1l-.85.6V16h-4v-2.3l-.85-.6C7.8 12.16 7 10.63 7 9c0-2.76 2.24-5 5-5s5 2.24 5 5c0 1.63-.8 3.16-2.15 4.1z"

		distanceFromBorder := 5.0
		iconWidth := 20.0

		conceptLogo.X_Offset = distanceFromBorder
		conceptLogo.Y_Offset = distanceFromBorder
		conceptLogo.RectAnchorType = svg.RECT_TOP_LEFT
		rect.RectAnchoredPaths = append(rect.RectAnchoredPaths, conceptLogo)

		if len(rect.RectAnchoredTexts) > 0 {
			padding := 10.0
			title := rect.RectAnchoredTexts[0]
			if rect.Width > (distanceFromBorder + iconWidth + padding) {
				title.Content = strutils.WrapStringPreservingNewlines(title.Content, int((rect.Width-(distanceFromBorder+iconWidth+padding))/root.NbPixPerCharacter))
			}
			title.X_Offset = (distanceFromBorder + iconWidth) / 2.0
		}
	}

	for _, deliverableConceptShape := range diagram.DeliverableConceptShapes {
		if deliverableConceptShape.GetIsHidden() {
			continue
		}
		deliverable := deliverableConceptShape.Deliverable
		concept := deliverableConceptShape.Concept
		if deliverable == nil || concept == nil {
			log.Panic("There should be a deliverable and a concept")
		}
		startRect := diagram.map_Deliverable_Rect[deliverable]
		endRect := diagram.map_Concept_Rect[concept]
		if startRect == nil || endRect == nil {
			continue
		}
		link := svgAssociationLink(
			stager,
			startRect, endRect,
			deliverableConceptShape,
			deliverable,
			layer,
			true,
		)
		configureLinkWithControlPoints(stager, link, &deliverableConceptShape.ControlPointShapes, startRect.Name)
	}

	svg.StageBranch(svgStage, stager.svgObject)
	stager.svgStage.Commit()

	return stager.svgObject
}

func applyStyleToRectAnchoredText(title *svg.RectAnchoredText) {
	title.Stroke = svg.Black.ToString()
	title.StrokeWidth = 1
	title.StrokeOpacity = 1
	title.Color = svg.Black.ToString()
	title.FillOpacity = 1

	title.FontSize = "12px"
}

func configureLinkWithControlPoints(
	stager *Stager,
	link *svg.Link,
	controlPointShapes *[]*ControlPointShape,
	startRectName string,
) {
	link.Type = svg.LINK_TYPE_LINE_WITH_CONTROL_POINTS
	link.StartAnchorType = svg.ANCHOR_CENTER
	link.EndAnchorType = svg.ANCHOR_CENTER
	link.StartArrowOffset = 15.0
	link.EndArrowOffset = 15.0

	for _, controlPointShape := range *controlPointShapes {
		closestRect := link.Start
		if !controlPointShape.IsStartShapeTheClosestShape {
			closestRect = link.End
		}

		link.ControlPoints = append(link.ControlPoints, &svg.ControlPoint{
			Name:        controlPointShape.Name,
			X_Relative:  controlPointShape.X_Relative,
			Y_Relative:  controlPointShape.Y_Relative,
			ClosestRect: closestRect,
			Impl: &ControlPointShapeProxy{
				stager:             stager,
				controlPointShapes: controlPointShapes,
				controlPointShape:  controlPointShape,
				sourceName:         startRectName,
			},
		})
	}
	link.Impl = &LinkShapeProxy{
		stager:             stager,
		controlPointShapes: controlPointShapes,
		linkName:           link.Name,
		sourceName:         startRectName,
	}
}
