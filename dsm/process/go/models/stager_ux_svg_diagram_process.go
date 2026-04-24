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

	diagramProcess.map_Process_Rect = make(map[*Process]*svg.Rect)
	diagramProcess.map_SvgRect_ProcessShape = make(map[*svg.Rect]*ProcessShape)

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
		diagramProcess.map_SvgRect_ProcessShape[rect] = processShape
	}

	for _, ProcessCompositionShape := range diagramProcess.ProcessComposition_Shapes {
		if ProcessCompositionShape.GetIsHidden() {
			continue
		}
		_ = ProcessCompositionShape
		subProcess := ProcessCompositionShape.Process
		parentProcess := subProcess.parentProcess

		if subProcess == nil || parentProcess == nil {
			log.Panic("There should be a subProcess and a parentProcess")
		}

		startRect := diagramProcess.map_Process_Rect[parentProcess]
		endRect := diagramProcess.map_Process_Rect[subProcess]

		if startRect == nil || endRect == nil {
			continue
		}

		svgAssociationLink(
			stager,
			startRect, endRect,
			ProcessCompositionShape,
			// when one clicks on the link, this is the form of the parent Process
			parentProcess,
			layer,
			false)
	}

	return svgObject
}
