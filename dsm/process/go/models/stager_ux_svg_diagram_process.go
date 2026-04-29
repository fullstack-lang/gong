package models

import (
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
	diagramProcess.map_Task_Rect = make(map[*Task]*svg.Rect)
	for _, taskShape := range diagramProcess.TaskShapes {
		if taskShape.IsHidden {
			continue
		}

		rect := svgRect(
			stager,
			diagramProcess,
			taskShape,
			layer)

		diagramProcess.map_Task_Rect[taskShape.Task] = rect
	}

	return svgObject
}
