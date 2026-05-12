package models

import (
	"fmt"
	"log"
	"strings"

	svg "github.com/fullstack-lang/gong/lib/svg/go/models"
	"github.com/fullstack-lang/gong/pkg/strutils"
)

func (stager *Stager) svg() {
	log.Println("svg")
	stager.processDiagramSvgStage.Reset()

	var diagramProcess *DiagramProcess
	{
		for diagramprocess_ := range *GetGongstructInstancesSet[DiagramProcess](stager.stage) {
			if diagramprocess_.IsChecked {
				diagramProcess = diagramprocess_
			}
		}
	}

	if diagramProcess == nil {
		stager.processDiagramSvgStage.Commit()
		return
	}
	svgObject := stager.generateSvgObject(diagramProcess)

	svg.StageBranch(stager.processDiagramSvgStage, svgObject)
	stager.svgObjectDiagramProcess = svgObject
	stager.svgObjectDiagramProcess.OnUpdate = stager.onUpdateSVG

	stager.processDiagramSvgStage.Commit()
}

// generateSvgObject creates and returns a new svg.SVG object representing the given diagram.
// It maps all visible domain shapes (Processs, Tasks, Notes, Resources) and their associations
// to SVG elements (Rects, Links, Paths) on a single layer. It also populates the diagram's
// internal maps to link abstract elements with their visual SVG counterparts.
func (stager *Stager) generateSvgObject(diagramProcess *DiagramProcess) *svg.SVG {
	svgObject := (&svg.SVG{Name: `SVG`})
	stager.diagramProcess = diagramProcess

	svgObject.OverrideWidth = true
	svgObject.OverriddenWidth = diagramProcess.Width
	svgObject.OverrideHeight = true
	svgObject.OverriddenHeight = diagramProcess.Height

	diagramProcess.map_SvgRect_TaskShape = make(map[*svg.Rect]*TaskShape)
	diagramProcess.map_SvgRect_ExternalParticipantShape = map[*svg.Rect]*ExternalParticipantShape{}
	diagramProcess.map_SvgRect_Participant = map[*svg.Rect]*Participant{}

	// // to implement association between abstract elements by mouse drag
	// svgImpl := &svgProxy{
	// 	stager:  stager,
	// 	svg_:
	// 	diagramProcess: diagramProcess,
	// }
	// Impl = svgImpl

	svgObject.Name = diagramProcess.Name
	svgObject.IsEditable = diagramProcess.IsEditable()

	layer := (&svg.Layer{Name: "Layer 1"})
	svgObject.Layers = append(svgObject.Layers, layer)

	stager.drawProcessShapes(diagramProcess, layer)

	rectOfOwningProcess := diagramProcess.map_Process_Rect[diagramProcess.owningProcess]
	if rectOfOwningProcess != nil {
		stager.drawParticipantShapes(diagramProcess, layer, rectOfOwningProcess)
	}
	stager.drawExternalParticipantShapes(diagramProcess, layer)
	stager.drawTaskShapes(diagramProcess, layer)
	stager.drawControlFlowShapes(diagramProcess, layer)
	stager.drawDataFlowShapes(diagramProcess, layer)
	map_Note_Rect := stager.drawNoteShapes(diagramProcess, layer)
	stager.drawNoteTaskShapes(diagramProcess, layer, map_Note_Rect)

	return svgObject
}

func (stager *Stager) drawProcessShapes(diagramProcess *DiagramProcess, layer *svg.Layer) {
	diagramProcess.map_Process_Rect = make(map[*Process]*svg.Rect)
	for _, processShape := range diagramProcess.Process_Shapes {
		if processShape.IsHidden {
			continue
		}

		rect := svgRect(
			stager,
			diagramProcess,
			processShape,
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
		// we need to update the position of the participants shapes and task shapes that are within the process
		rect.OnUpdate = func(updatedRect *svg.Rect) {
			diffX := processShape.GetX() != updatedRect.X
			diffY := processShape.GetY() != updatedRect.Y
			diffWidth := processShape.GetWidth() != updatedRect.Width
			diffHeight := processShape.GetHeight() != updatedRect.Height

			if diffX || diffY || diffWidth || diffHeight {
				deltaX := updatedRect.X - processShape.GetX()
				deltaY := updatedRect.Y - processShape.GetY()
				processShape.SetX(updatedRect.X)
				processShape.SetY(updatedRect.Y)
				processShape.SetWidth(updatedRect.Width)
				processShape.SetHeight(updatedRect.Height)

				// update the position of the participant shapes and task shapes that are within the process
				for _, participantShape := range diagramProcess.Participant_Shapes {
					if participantShape.Participant.GetOwningProcess() == processShape.GetAbstractElement() {
						participantShape.SetX(participantShape.GetX() + deltaX)
						participantShape.SetY(participantShape.GetY() + deltaY)
					}
				}

				for _, taskShape := range diagramProcess.Task_Shapes {
					if taskShape.Task.GetOwningParticipant().GetOwningProcess() == processShape.GetAbstractElement() {
						taskShape.SetX(taskShape.GetX() + deltaX)
						taskShape.SetY(taskShape.GetY() + deltaY)
					}
				}

				stager.stage.Commit()
			}
		}

		diagramProcess.map_Process_Rect[processShape.Process] = rect
	}
}

func (stager *Stager) drawParticipantShapes(diagramProcess *DiagramProcess, layer *svg.Layer, rectOfOwningProcess *svg.Rect) {
	root := stager.GetRootLibrary()

	diagramProcess.map_Participant_Rect = make(map[*Participant]*svg.Rect)
	horizontalMargin := 10.0
	verticalTopMargin := 50.0
	verticalTopMarginForTitle := 60.0
	verticalBottomMargin := 10.0

	participantsWidth := rectOfOwningProcess.Width - 2*horizontalMargin

	var totalWeight float64
	for _, pShape := range diagramProcess.Participant_Shapes {
		weight := pShape.WidthWeight
		if weight == 0 {
			weight = 1.0
		}
		totalWeight += weight
	}

	currentWeight := 0.0

	for idx, participantShape := range diagramProcess.Participant_Shapes {
		shapeWeight := participantShape.WidthWeight
		if shapeWeight == 0 {
			shapeWeight = 1.0
		}
		participantWidth := 0.0
		if totalWeight > 0 {
			participantWidth = shapeWeight * (participantsWidth / totalWeight)
		}

		if participantShape.IsHidden {
			currentWeight += shapeWeight
			continue
		}

		rect := new(svg.Rect)
		layer.Rects = append(layer.Rects, rect)
		diagramProcess.map_SvgRect_Participant[rect] = participantShape.Participant

		rect.Name = participantShape.GetName()
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

		// all but the last participant have right handle
		rect.CanHaveRightHandle = func() bool {
			if idx == len(diagramProcess.Participant_Shapes)-1 {
				return false
			}
			return true
		}()

		// visuals
		rect.RX = 0
		rect.StrokeWidth = 1

		if totalWeight > 0 {
			rect.X = rectOfOwningProcess.X + horizontalMargin + currentWeight*(participantsWidth/totalWeight)
		} else {
			rect.X = rectOfOwningProcess.X + horizontalMargin
		}
		rect.Width = participantWidth

		rect.Y = rectOfOwningProcess.Y + verticalTopMargin + verticalTopMarginForTitle
		rect.Height = rectOfOwningProcess.Height - verticalTopMargin - verticalBottomMargin - verticalTopMarginForTitle

		// override default behavior, we need to commit when the rect is moved
		rect.OnUpdate = func(updatedRect *svg.Rect) {
			if updatedRect.Width != participantWidth {
				othersWeight := totalWeight - shapeWeight
				if othersWeight > 0 {
					boundedWidth := updatedRect.Width
					if boundedWidth > participantsWidth-10 {
						boundedWidth = participantsWidth - 10
					}
					if boundedWidth < 10 {
						boundedWidth = 10
					}
					newWeight := (boundedWidth * othersWeight) / (participantsWidth - boundedWidth)

					participantShape.WidthWeight = newWeight
					stager.stage.Commit()
				}
			}
		}

		diagramProcess.map_Participant_Rect[participantShape.Participant] = rect

		{
			title := new(svg.RectAnchoredText)
			title.Name = participantShape.GetAbstractElement().GetName()

			content := participantShape.GetAbstractElement().GetName()
			if diagramProcess.GetIsShowPrefix() {
				content = participantShape.GetAbstractElement().GetComputedPrefix() + " " + content
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
		{
			titleBox := &svg.RectAnchoredRect{
				Name: participantShape.GetAbstractElement().GetName(),
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
		}

		stager.drawAllocatedResources(participantShape.Participant, diagramProcess, rect, svg.RECT_TOP)

		currentWeight += shapeWeight
	}
}

func (stager *Stager) drawExternalParticipantShapes(diagramProcess *DiagramProcess, layer *svg.Layer) {
	root := stager.GetRootLibrary()

	diagramProcess.map_ExternalParticipant_Rect = make(map[*Participant]*svg.Rect)
	for _, externalParticipantShape := range diagramProcess.ExternalParticipant_Shapes {
		if externalParticipantShape.IsHidden {
			continue
		}

		rect := new(svg.Rect)
		layer.Rects = append(layer.Rects, rect)

		rect.Name = externalParticipantShape.GetName()
		rect.X = externalParticipantShape.GetX()
		rect.Y = externalParticipantShape.GetY()
		rect.Width = externalParticipantShape.GetWidth()
		rect.Height = externalParticipantShape.GetHeight()

		rect.CanMoveHorizontaly = true
		rect.CanMoveVerticaly = true
		rect.CanHaveBottomHandle = true
		rect.CanHaveLeftHandle = true
		rect.CanHaveRightHandle = true
		rect.CanHaveTopHandle = true

		rect.RX = 0
		rect.StrokeWidth = 1.0
		rect.StrokeOpacity = 1.0
		rect.Color = "#FFFDE7"
		rect.FillOpacity = 1.0
		rect.Stroke = "#E0E0E0"

		{
			title := new(svg.RectAnchoredText)
			title.Name = externalParticipantShape.GetAbstractElement().GetName()

			content := externalParticipantShape.GetAbstractElement().GetName()
			if diagramProcess.GetIsShowPrefix() {
				content = externalParticipantShape.GetAbstractElement().GetComputedPrefix() + " " + content
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

		rect.OnUpdate = onUpdateRectElement(stager, externalParticipantShape.Participant, externalParticipantShape, false)

		externalParticipantWidth := 0.1
		if externalParticipantShape.TailHeigth == 0 {
			externalParticipantShape.TailHeigth = 100.0
		}

		boxHeight := stager.drawAllocatedResources(externalParticipantShape.Participant, diagramProcess, rect, svg.RECT_BOTTOM)

		tailRect := &svg.Rect{
			Name: "Tail" + rect.GetName(),
			Presentation: svg.Presentation{
				Stroke:          "#9E9E9E",
				StrokeWidth:     1.5,
				StrokeOpacity:   1,
				StrokeDashArray: "5 5",
				Color:           "transparent",
				FillOpacity:     0.0,
			},
			Width:               externalParticipantWidth,
			Height:              externalParticipantShape.TailHeigth - boxHeight,
			X:                   rect.X + (rect.Width-externalParticipantWidth)/2.0,
			Y:                   rect.Y + rect.Height + boxHeight,
			CanHaveBottomHandle: true,
		}
		diagramProcess.map_ExternalParticipant_Rect[externalParticipantShape.Participant] = tailRect

		// this is the tail rect that is used for drawing links between external participant and task, so we need to keep a reference of it in the diagram process
		diagramProcess.map_SvgRect_ExternalParticipantShape[tailRect] = externalParticipantShape
		// if the tailRect bottom handle is used, the heigth is updated
		tailRect.OnUpdate = func(updatedRect *svg.Rect) {
			diffSize := externalParticipantShape.TailHeigth != updatedRect.Height

			if diffSize {
				externalParticipantShape.TailHeigth = updatedRect.Height
				stager.stage.CommitWithSuspendedCallbacks()
			}
		}
		layer.Rects = append(layer.Rects, tailRect)

	}
}

func (stager *Stager) drawTaskShapes(diagramProcess *DiagramProcess, layer *svg.Layer) {
	rm := GetSliceOfPointersReverseMap[Participant, Task](GetAssociationName[Participant]().Tasks[0].Name, stager.stage)

	diagramProcess.map_Task_Rect = make(map[*Task]*svg.Rect)
	for _, taskShape := range diagramProcess.Task_Shapes {
		if taskShape.IsHidden {
			continue
		}

		task := taskShape.Task
		participants := rm[task]
		if len(participants) == 0 {
			continue
		}

		participant := participants[0]
		participantRect := diagramProcess.map_Participant_Rect[participant]
		if participantRect == nil {
			continue
		}

		rect := svgRect(
			stager,
			diagramProcess,
			taskShape,
			layer)
		diagramProcess.map_SvgRect_TaskShape[rect] = taskShape

		rect.EnclosingRect = participantRect

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
		if task.IsStartTask {
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
		if task.IsEndTask {
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
		diagramProcess.map_Task_Rect[taskShape.Task] = rect
	}
}

func (stager *Stager) drawControlFlowShapes(diagramProcess *DiagramProcess, layer *svg.Layer) {
	for _, controlFlowShape := range diagramProcess.ControlFlow_Shapes {
		if controlFlowShape.GetIsHidden() {
			continue
		}
		_ = controlFlowShape
		startTask := controlFlowShape.ControlFlow.Start
		endTask := controlFlowShape.ControlFlow.End

		if startTask == nil || endTask == nil {
			log.Panic("There should be a start task and a end task")
		}

		startRect := diagramProcess.map_Task_Rect[startTask]
		endRect := diagramProcess.map_Task_Rect[endTask]

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

func (stager *Stager) drawDataFlowShapes(diagramProcess *DiagramProcess, layer *svg.Layer) {
	for _, dataFlowShape := range diagramProcess.DataFlow_Shapes {
		if dataFlowShape.GetIsHidden() {
			continue
		}

		var startRect, endRect *svg.Rect
		df := dataFlowShape.DataFlow

		switch df.Type {
		case DataFlow_Task2Task:
			if df.StartTask == nil || df.EndTask == nil {
				log.Panic("There should be a start task and an end task")
			}
			startRect = diagramProcess.map_Task_Rect[df.StartTask]
			endRect = diagramProcess.map_Task_Rect[df.EndTask]
		case DataFlow_ExternalParticipant2Task:
			if df.StartExternalParticipant == nil || df.EndTask == nil {
				log.Panic("There should be a start external participant and an end task")
			}
			startRect = diagramProcess.map_ExternalParticipant_Rect[df.StartExternalParticipant]
			endRect = diagramProcess.map_Task_Rect[df.EndTask]
		case DataFlow_Task2ExternalParticipant:
			if df.StartTask == nil || df.EndExternalParticipant == nil {
				log.Panic("There should be a start task and an end external participant")
			}
			startRect = diagramProcess.map_Task_Rect[df.StartTask]
			endRect = diagramProcess.map_ExternalParticipant_Rect[df.EndExternalParticipant]
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
			rectAnchoredLink := &svg.LinkAnchoredText{
				Name:    dataShape.Name,
				Content: dataShape.Data.Name,
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
				X_Offset: 4.0,
			}
			link.TextAtCorner = append(link.TextAtCorner, rectAnchoredLink)
		}
	}
}

func (stager *Stager) drawNoteShapes(diagramProcess *DiagramProcess, layer *svg.Layer) map[*Note]*svg.Rect {
	map_Note_Rect := make(map[*Note]*svg.Rect)
	for _, noteShape := range diagramProcess.Note_Shapes {
		if noteShape.GetIsHidden() {
			continue
		}

		rect := svgRect(
			stager,
			diagramProcess,
			noteShape,
			layer)

		map_Note_Rect[noteShape.Note] = rect

		rect.Color = "#FFFDE7"
		rect.FillOpacity = 1.0
		rect.Stroke = "#FBC02D"
		rect.StrokeWidth = 1.0
		rect.RX = 0.0

		if len(rect.RectAnchoredTexts) > 0 {
			rect.RectAnchoredTexts[0].FontFamily = "sans-serif"
			rect.RectAnchoredTexts[0].Color = "#333333"

			if noteShape.Note != nil {
				content := noteShape.Note.GetName()
				if content == "" {
					content = noteShape.Note.GetName()
				}
				root := stager.GetRootLibrary()
				if rect.Width > 0 && root.NbPixPerCharacter > 0 {
					content = strutils.WrapStringPreservingNewlines(content, int(rect.Width/root.NbPixPerCharacter))
				}
				rect.RectAnchoredTexts[0].Content = content
			}
		}
	}
	return map_Note_Rect
}

func (stager *Stager) drawNoteTaskShapes(diagramProcess *DiagramProcess, layer *svg.Layer, map_Note_Rect map[*Note]*svg.Rect) {
	for _, noteTaskShape := range diagramProcess.NoteTaskShapes {
		if noteTaskShape.GetIsHidden() {
			continue
		}

		startNote := noteTaskShape.Note
		endTask := noteTaskShape.Task

		if startNote == nil || endTask == nil {
			continue
		}

		startRect := map_Note_Rect[startNote]
		endRect := diagramProcess.map_Task_Rect[endTask]

		if startRect == nil || endRect == nil {
			continue
		}

		link := svgAssociationLink(
			stager,
			startRect, endRect,
			noteTaskShape,
			startNote,
			layer,
			false)

		link.Type = svg.LINK_TYPE_LINE_WITH_CONTROL_POINTS
		link.StartAnchorType = svg.ANCHOR_CENTER
		link.EndAnchorType = svg.ANCHOR_CENTER
		link.HasEndArrow = false

		link.Presentation.StrokeDashArray = "3,3"
		link.Stroke = "#9E9E9E"
		link.StrokeWidth = 1.0
	}
}

func (stager *Stager) drawAllocatedResources(participant *Participant, diagramProcess *DiagramProcess, rect *svg.Rect, rectAnchorType svg.RectAnchorType) (boxHeight float64) {
	root := stager.GetRootLibrary()
	const HeightBetween2AttributeShapes = 16.0

	// draw allocated resource shapes that are within the participant
	totalLines := 0
	X_Offset := 30.0
	Y_Offset := 20.0
	for _, resource := range participant.Resources {
		key := allocatedResourceShapeKey{
			participant: participant,
			resource:    resource,
		}
		allocatedResourceShape := diagramProcess.map_AllocatedResourceShapeKey_AllocatedResourceShape[key]
		if allocatedResourceShape == nil {
			continue
		}

		content := resource.Name
		if rect.Width > 0 {
			content = strutils.WrapStringPreservingNewlinesScaled(content, rect.Width, float64(root.NbPixPerCharacter), 12.0, 16.0)
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
				FontSize:  "12px",
			},

			DominantBaseline: svg.DominantBaselineMiddle,
			X_Offset:         X_Offset,
			Y_Offset:         Y_Offset + float64(totalLines)*HeightBetween2AttributeShapes,
			RectAnchorType:   rectAnchorType,
			TextAnchorType:   svg.TEXT_ANCHOR_START,
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
		if resource.SVG_Path == "" {
			continue
		}
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
			AppliedScaling:      1.0 / resource.InverseAppliedScaling,
		}
		if rectAnchorType == svg.RECT_BOTTOM {
			allocatedResourcePath.RectAnchorType = svg.RECT_BOTTOM_LEFT
		} else {
			allocatedResourcePath.RectAnchorType = svg.RECT_TOP_LEFT
		}
		rect.RectAnchoredPaths = append(rect.RectAnchoredPaths, allocatedResourcePath)
	}
	// draw a rect around the allocated resource shapes if there is at least one allocated resource shape
	boxHeight = float64(totalLines)*HeightBetween2AttributeShapes + Y_Offset - 5.0
	if totalLines > 0 {
		lineWidth := 1.0
		allocatedResourceRect := &svg.RectAnchoredRect{
			Name: participant.Name + "_allocated_resources",
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
	return boxHeight
}
