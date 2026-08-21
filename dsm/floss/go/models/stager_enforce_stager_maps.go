package models

import (
	"slices"
)

func (stager *Stager) enforceStagerMaps() {
	stager.map_Element_Diagrams = make(map[AbstractType][]*DiagramFloss)

	for _, diagramFloss := range GetGongstrucsSorted[*DiagramFloss](stager.stage) {
		updateMapElementDiagrams(stager, diagramFloss, diagramFloss.System_Shapes, &diagramFloss.map_System_SystemShape)
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
