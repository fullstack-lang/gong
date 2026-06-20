package models

import (
	"slices"
)

func (stager *Stager) enforceStagerMaps() {
	stager.map_Element_Diagrams = make(map[AbstractType][]*DiagramStructure)

	for _, diagramStructure := range GetGongstrucsSorted[*DiagramStructure](stager.stage) {
		_ = diagramStructure

		updateMapElementDiagrams(stager, diagramStructure, diagramStructure.System_Shapes, &diagramStructure.map_System_SystemShape)
		updateMapElementDiagrams(stager, diagramStructure, diagramStructure.Part_Shapes, &diagramStructure.map_Part_PartShape)
		updateMapElementDiagrams(stager, diagramStructure, diagramStructure.ExternalPart_Shapes, &diagramStructure.map_Part_ExternalPartShape)
		updateMapElementDiagrams(stager, diagramStructure, diagramStructure.Port_Shapes, &diagramStructure.map_Port_PortShape)
		updateMapElementDiagrams(stager, diagramStructure, diagramStructure.ControlFlow_Shapes, &diagramStructure.map_ControlFlow_ControlFlowShape)
		updateMapElementDiagrams(stager, diagramStructure, diagramStructure.DataFlow_Shapes, &diagramStructure.map_DataFlow_DataFlowShape)
		updateMapElementDiagrams(stager, diagramStructure, diagramStructure.Note_Shapes, &diagramStructure.map_Note_NoteShape)

		diagramStructure.map_DataShapeKey_DataShape = make(map[dataShapeKey]*DataShape)
		for _, dataShape := range diagramStructure.Data_Shapes {
			key := dataShapeKey{
				dataFlow: dataShape.DataFlow,
				data:     dataShape.Data,
			}
			diagramStructure.map_DataShapeKey_DataShape[key] = dataShape
		}

		diagramStructure.map_AllocatedResourceShapeKey_AllocatedResourceShape = make(map[allocatedResourceShapeKey]*AllocatedResourceShape)
		for _, allocatedResourceShape := range diagramStructure.AllocatedResourceShapes {
			key := allocatedResourceShapeKey{
				part: allocatedResourceShape.Part,
				resource:    allocatedResourceShape.Resource,
			}
			diagramStructure.map_AllocatedResourceShapeKey_AllocatedResourceShape[key] = allocatedResourceShape
		}

		diagramStructure.map_AllocatedSystemShapeKey_AllocatedSystemShape = make(map[allocatedSystemShapeKey]*AllocatedSystemShape)
		for _, allocatedSystemShape := range diagramStructure.AllocatedSystemShapes {
			key := allocatedSystemShapeKey{
				part: allocatedSystemShape.Part,
				system:     allocatedSystemShape.System,
			}
			diagramStructure.map_AllocatedSystemShapeKey_AllocatedSystemShape[key] = allocatedSystemShape
		}

		diagramStructure.map_Note_NotePortShape = make(map[notePortKey]*NotePortShape)
		for _, notePortShape := range diagramStructure.NotePortShapes {
			key := notePortKey{
				Note: notePortShape.Note,
				Port: notePortShape.Port,
			}
			diagramStructure.map_Note_NotePortShape[key] = notePortShape
		}
	}

	stager.rm_Data_DataFlows = GetSliceOfPointersReverseMap[DataFlow, Data](GetAssociationName[DataFlow]().Datas[0].Name, stager.stage)
	stager.rm_Resource_Parts = GetSliceOfPointersReverseMap[Part, Resource](GetAssociationName[Part]().Resources[0].Name, stager.stage)
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
	diagram *DiagramStructure,
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
			diagrams = []*DiagramStructure{diagram}
		}
		if !slices.Contains(diagrams, diagram) {
			diagrams = append(diagrams, diagram)
		}
		stager.map_Element_Diagrams[abstractElement] = diagrams
	}
}
