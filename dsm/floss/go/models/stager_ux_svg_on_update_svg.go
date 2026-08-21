package models

import (
	svg "github.com/fullstack-lang/gong/lib/svg/go/models"
)

// SVGUpdated implements SVGImplInterface.
func (stager *Stager) onUpdateSVG(frontSVG *svg.SVG) {
	stage := stager.stage
	diagramFloss := stager.diagramFloss
	svgObjectDiagramFloss := stager.svgObjectDiagramFloss

	if svgObjectDiagramFloss.DrawingState == frontSVG.DrawingState {
		if diagramFloss != nil {
			stager.probeForm.FillUpFormFromGongstruct(diagramFloss, "Diagram")
		}
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

	// Case 1: System Diagram (DiagramFloss) is active
	if diagramFloss != nil && diagramFloss.IsChecked {
		var note *Note
		var complexity *Complexity
		var performance *Performance
		var effort *Effort

		if noteShape, ok := diagramFloss.map_SvgRect_NoteShape[startRect]; ok {
			note = noteShape.Note
		} else if noteShape, ok := diagramFloss.map_SvgRect_NoteShape[endRect]; ok {
			note = noteShape.Note
		}

		if compShape, ok := diagramFloss.map_SvgRect_ComplexityShape[startRect]; ok {
			complexity = compShape.Complexity
		} else if compShape, ok := diagramFloss.map_SvgRect_ComplexityShape[endRect]; ok {
			complexity = compShape.Complexity
		}

		if perfShape, ok := diagramFloss.map_SvgRect_PerformanceShape[startRect]; ok {
			performance = perfShape.Performance
		} else if perfShape, ok := diagramFloss.map_SvgRect_PerformanceShape[endRect]; ok {
			performance = perfShape.Performance
		}

		if effShape, ok := diagramFloss.map_SvgRect_EffortShape[startRect]; ok {
			effort = effShape.Effort
		} else if effShape, ok := diagramFloss.map_SvgRect_EffortShape[endRect]; ok {
			effort = effShape.Effort
		}

		if note != nil && complexity != nil {
			note.Complexities = append(note.Complexities, complexity)
			addAssociationShapeToDiagram(stager, note, complexity, &diagramFloss.NoteComplexityShapes)
		} else if note != nil && performance != nil {
			note.Performances = append(note.Performances, performance)
			addAssociationShapeToDiagram(stager, note, performance, &diagramFloss.NotePerformanceShapes)
		} else if note != nil && effort != nil {
			note.Efforts = append(note.Efforts, effort)
			addAssociationShapeToDiagram(stager, note, effort, &diagramFloss.NoteEffortShapes)
		}
	}

	// Case 2: Equation Diagram (DiagramFlossEquation) is active
	var diagramEquation *DiagramFlossEquation
	for d := range *GetGongstructInstancesSet[DiagramFlossEquation](stage) {
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
