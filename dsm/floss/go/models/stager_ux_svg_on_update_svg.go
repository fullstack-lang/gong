package models

import (
	svg "github.com/fullstack-lang/gong/lib/svg/go/models"
)

// SVGUpdated implements SVGImplInterface.
func (stager *Stager) onUpdateSVG(frontSVG *svg.SVG) {
	svgObjectDiagramFloss := stager.svgObjectDiagramFloss
	if svgObjectDiagramFloss != nil && svgObjectDiagramFloss.DrawingState == frontSVG.DrawingState {
		return
	}

	// IMPORTANT : we are only interested when the updateSVG has finished drawing the connection
	if frontSVG.DrawingState == svg.DRAWING_LINK {
		return
	}

	startRect := frontSVG.StartRect
	endRect := frontSVG.EndRect

	if startRect == nil || endRect == nil {
		stager.stage.Commit()
		return
	}

	// Equation Diagram (DiagramFlossEquation) is active
	var diagramEquation *DiagramFlossEquation
	for d := range *GetGongstructInstancesSet[DiagramFlossEquation](stager.stage) {
		if d.IsChecked {
			diagramEquation = d
			break
		}
	}

	if diagramEquation != nil && diagramEquation.IsChecked {
		var note *Note
		var complexity *Complexity
		var performance *Performance
		var effort *Effort

		if noteShape, ok := diagramEquation.map_SvgRect_NoteShape[startRect]; ok {
			note = noteShape.Note
		} else if noteShape, ok := diagramEquation.map_SvgRect_NoteShape[endRect]; ok {
			note = noteShape.Note
		}

		if c, ok := diagramEquation.map_SvgRect_Complexity[startRect]; ok {
			complexity = c
		} else if c, ok := diagramEquation.map_SvgRect_Complexity[endRect]; ok {
			complexity = c
		}

		if p, ok := diagramEquation.map_SvgRect_Performance[startRect]; ok {
			performance = p
		} else if p, ok := diagramEquation.map_SvgRect_Performance[endRect]; ok {
			performance = p
		}

		if e, ok := diagramEquation.map_SvgRect_Effort[startRect]; ok {
			effort = e
		} else if e, ok := diagramEquation.map_SvgRect_Effort[endRect]; ok {
			effort = e
		}

		if note != nil && complexity != nil {
			note.Complexities = append(note.Complexities, complexity)
			addAssociationShapeToDiagram(stager, note, complexity, &diagramEquation.NoteComplexityShapes)
		} else if note != nil && performance != nil {
			note.Performances = append(note.Performances, performance)
			addAssociationShapeToDiagram(stager, note, performance, &diagramEquation.NotePerformanceShapes)
		} else if note != nil && effort != nil {
			note.Efforts = append(note.Efforts, effort)
			addAssociationShapeToDiagram(stager, note, effort, &diagramEquation.NoteEffortShapes)
		}
	}

	// commit to encode the result, this will generate a new SVG generation
	stager.stage.Commit()
}
