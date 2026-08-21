package models

import (
	"slices"
)

func (stager *Stager) enforceStagerMaps() {
	stager.map_Element_Diagrams = make(map[AbstractType][]*DiagramFloss)

	for _, diagramFloss := range GetGongstrucsSorted[*DiagramFloss](stager.stage) {
		updateMapElementDiagrams(stager, diagramFloss, diagramFloss.System_Shapes, &diagramFloss.map_System_SystemShape)
		updateMapElementDiagrams(stager, diagramFloss, diagramFloss.Complexity_Shapes, &diagramFloss.map_Complexity_ComplexityShape)
		updateMapElementDiagrams(stager, diagramFloss, diagramFloss.Performance_Shapes, &diagramFloss.map_Performance_PerformanceShape)
		updateMapElementDiagrams(stager, diagramFloss, diagramFloss.Effort_Shapes, &diagramFloss.map_Effort_EffortShape)
		updateMapElementDiagrams(stager, diagramFloss, diagramFloss.Note_Shapes, &diagramFloss.map_Note_NoteShape)

		diagramFloss.map_Note_NoteComplexityShape = make(map[noteComplexityKey]*NoteComplexityShape)
		for _, shape := range diagramFloss.NoteComplexityShapes {
			key := noteComplexityKey{
				Note:       shape.Note,
				Complexity: shape.Complexity,
			}
			diagramFloss.map_Note_NoteComplexityShape[key] = shape
		}

		diagramFloss.map_Note_NotePerformanceShape = make(map[notePerformanceKey]*NotePerformanceShape)
		for _, shape := range diagramFloss.NotePerformanceShapes {
			key := notePerformanceKey{
				Note:        shape.Note,
				Performance: shape.Performance,
			}
			diagramFloss.map_Note_NotePerformanceShape[key] = shape
		}

		diagramFloss.map_Note_NoteEffortShape = make(map[noteEffortKey]*NoteEffortShape)
		for _, shape := range diagramFloss.NoteEffortShapes {
			key := noteEffortKey{
				Note:   shape.Note,
				Effort: shape.Effort,
			}
			diagramFloss.map_Note_NoteEffortShape[key] = shape
		}
	}

	for _, diagramEquation := range GetGongstrucsSorted[*DiagramFlossEquation](stager.stage) {
		diagramEquation.map_Note_NoteShape = make(map[*Note]*NoteShape)
		for _, shape := range diagramEquation.Note_Shapes {
			if shape.Note != nil {
				diagramEquation.map_Note_NoteShape[shape.Note] = shape
			}
		}

		diagramEquation.map_Note_NoteComplexityShape = make(map[noteComplexityKey]*NoteComplexityShape)
		for _, shape := range diagramEquation.NoteComplexityShapes {
			key := noteComplexityKey{
				Note:       shape.Note,
				Complexity: shape.Complexity,
			}
			diagramEquation.map_Note_NoteComplexityShape[key] = shape
		}

		diagramEquation.map_Note_NotePerformanceShape = make(map[notePerformanceKey]*NotePerformanceShape)
		for _, shape := range diagramEquation.NotePerformanceShapes {
			key := notePerformanceKey{
				Note:        shape.Note,
				Performance: shape.Performance,
			}
			diagramEquation.map_Note_NotePerformanceShape[key] = shape
		}

		diagramEquation.map_Note_NoteEffortShape = make(map[noteEffortKey]*NoteEffortShape)
		for _, shape := range diagramEquation.NoteEffortShapes {
			key := noteEffortKey{
				Note:   shape.Note,
				Effort: shape.Effort,
			}
			diagramEquation.map_Note_NoteEffortShape[key] = shape
		}
	}
}

// updateMapElementDiagrams is a helper function to update the map of abstract elements to their shapes for a given diagram
// and track the diagrams where each element is displayed.
func updateMapElementDiagrams[
	AT interface {
		AbstractType
		comparable
	},
	CT ConcreteType,
](
	stager *Stager,
	diagram *DiagramFloss,
	shapes []CT,
	diagramMapPtr *map[AT]CT, // Now a pointer to the map
) {
	// 1. Initialize the map right on the struct
	*diagramMapPtr = make(map[AT]CT)

	for _, shape := range shapes {
		abstractElement, ok := shape.GetAbstractElement().(AT)
		if !ok {
			continue
		}

		// 2. Dereference the pointer to map the abstract element to its Shape
		(*diagramMapPtr)[abstractElement] = shape

		// track all diagrams that display this element across the stage
		diagrams := stager.map_Element_Diagrams[abstractElement]
		if diagrams == nil {
			diagrams = []*DiagramFloss{diagram}
		}
		if !slices.Contains(diagrams, diagram) {
			diagrams = append(diagrams, diagram)
		}
		stager.map_Element_Diagrams[abstractElement] = diagrams
	}
}
