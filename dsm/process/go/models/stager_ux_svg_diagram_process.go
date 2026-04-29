package models

import (
	"fmt"
	"log"

	svg "github.com/fullstack-lang/gong/lib/svg/go/models"
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

	//
	// Process
	//
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
		diagramProcess.map_Process_Rect[processShape.Process] = rect
	}

	rectOfOwningProcess := diagramProcess.map_Process_Rect[diagramProcess.owningProcess]

	//
	// Participant
	//
	diagramProcess.map_Participant_Rect = make(map[*Participant]*svg.Rect)
	horizontalMargin := 10.0
	verticalTopMargin := 50.0
	verticalBottomMargin := 10.0

	participantsWidth := rectOfOwningProcess.Width - 2*horizontalMargin
	participantWidth := participantsWidth / float64(len(diagramProcess.Participant_Shapes))

	for idx, participantShape := range diagramProcess.Participant_Shapes {
		if participantShape.IsHidden {
			continue
		}

		rect := svgRect(
			stager,
			diagramProcess,
			participantShape,
			layer)

		// rect cannot move
		rect.CanMoveHorizontaly = false
		rect.CanMoveVerticaly = false
		rect.CanHaveBottomHandle = false
		rect.CanHaveLeftHandle = false
		rect.CanHaveRightHandle = false
		rect.CanHaveTopHandle = false

		// visuals
		rect.RX = 0
		rect.StrokeWidth = 1

		rect.X = rectOfOwningProcess.X + horizontalMargin + float64(idx)*(participantWidth)
		rect.Width = participantWidth

		rect.Y = rectOfOwningProcess.Y + verticalTopMargin
		rect.Height = rectOfOwningProcess.Height - verticalTopMargin - verticalBottomMargin

		diagramProcess.map_Participant_Rect[participantShape.Participant] = rect
	}

	//
	// Task
	//
	rm := GetSliceOfPointersReverseMap[Participant, Task](GetAssociationName[Participant]().Tasks[0].Name, stager.stage)

	diagramProcess.map_Task_Rect = make(map[*Task]*svg.Rect)
	for _, taskShape := range diagramProcess.TaskShapes {
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

		rect.EnclosingRect = participantRect

		// pick up the title of the rect
		stateTitleText := rect.RectAnchoredTexts[0]
		var smallRadius = 10.0
		if task.IsStartTask {
			stateTitleText.TextAnchorType = svg.TEXT_ANCHOR_START
			stateTitleText.RectAnchorType = svg.RECT_TOP_LEFT
			stateTitleText.DominantBaseline = svg.DominantBaselineCentral
			stateTitleText.WhiteSpace = svg.WhiteSpaceEnumPre
			stateTitleText.X_Offset = 0
			stateTitleText.Y_Offset = 0

			circle := new(svg.RectAnchoredPath)
			circle.Stroke = svg.Black.ToString()
			circle.StrokeWidth = 2
			circle.StrokeOpacity = 1

			circle.Color = svg.Black.ToString()
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

		var bigRadius = 18.0
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

				circle.Stroke = svg.Black.ToString()
				circle.StrokeWidth = 1
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
				circle.Stroke = svg.Black.ToString()
				circle.StrokeWidth = 2
				circle.StrokeOpacity = 1

				circle.Color = svg.Black.ToString()
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

	return svgObject
}
