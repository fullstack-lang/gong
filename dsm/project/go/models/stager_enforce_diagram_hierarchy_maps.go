package models

import "slices"

func (stager *Stager) enforceDiagramMaps() {
	stager.map_Element_Diagrams = make(map[AbstractType][]DiagramIF)

	for _, diagramHierarchy := range GetGongstrucsSorted[*DiagramHierarchy](stager.stage) {
		// computes all products presents in the diagram
		diagramHierarchy.map_Product_ProductShape = make(map[*Product]*ProductShape)
		for _, shape := range diagramHierarchy.Product_Shapes {
			if shape.Product != nil {
				diagramHierarchy.map_Product_ProductShape[shape.Product] = shape
				diagrams := stager.map_Element_Diagrams[shape.Product]

				if diagrams == nil {
					diagrams = []DiagramIF{diagramHierarchy}
				}

				if !slices.Contains(diagrams, DiagramIF(diagramHierarchy)) {
					diagrams = append(diagrams, diagramHierarchy)
				}
				stager.map_Element_Diagrams[shape.Product] = diagrams
			}
		}

		diagramHierarchy.map_Task_TaskShape = make(map[*Task]*TaskShape)
		for _, shape := range diagramHierarchy.Task_Shapes {
			if shape.Task != nil {
				diagramHierarchy.map_Task_TaskShape[shape.Task] = shape
				diagrams := stager.map_Element_Diagrams[shape.Task]

				if diagrams == nil {
					diagrams = []DiagramIF{diagramHierarchy}
				}

				if !slices.Contains(diagrams, DiagramIF(diagramHierarchy)) {
					diagrams = append(diagrams, diagramHierarchy)
				}
				stager.map_Element_Diagrams[shape.Task] = diagrams
			}
		}

		diagramHierarchy.map_TaskGroup_TaskGroupShape = make(map[*TaskGroup]*TaskGroupShape)
		for _, shape := range diagramHierarchy.TaskGroupShapes {
			if shape.TaskGroup != nil {
				diagramHierarchy.map_TaskGroup_TaskGroupShape[shape.TaskGroup] = shape
				diagrams := stager.map_Element_Diagrams[shape.TaskGroup]

				if diagrams == nil {
					diagrams = []DiagramIF{diagramHierarchy}
				}

				if !slices.Contains(diagrams, DiagramIF(diagramHierarchy)) {
					diagrams = append(diagrams, diagramHierarchy)
				}
				stager.map_Element_Diagrams[shape.TaskGroup] = diagrams
			}
		}

		diagramHierarchy.map_Milestone_MilestoneShape = make(map[*Milestone]*MilestoneShape)
		for _, shape := range diagramHierarchy.MilestoneShapes {
			if shape.Milestone != nil {
				diagramHierarchy.map_Milestone_MilestoneShape[shape.Milestone] = shape
				diagrams := stager.map_Element_Diagrams[shape.Milestone]

				if diagrams == nil {
					diagrams = []DiagramIF{diagramHierarchy}
				}

				if !slices.Contains(diagrams, DiagramIF(diagramHierarchy)) {
					diagrams = append(diagrams, diagramHierarchy)
				}
				stager.map_Element_Diagrams[shape.Milestone] = diagrams
			}
		}

		diagramHierarchy.map_Task_TaskInputShape = make(map[taskProductKey]*TaskInputShape)
		for _, shape := range diagramHierarchy.TaskInputShapes {
			if shape.Task != nil && shape.Product != nil {
				diagramHierarchy.map_Task_TaskInputShape[taskProductKey{Task: shape.Task, Product: shape.Product}] = shape
			}
		}

		diagramHierarchy.map_Task_TaskOutputShape = make(map[taskProductKey]*TaskOutputShape)
		for _, shape := range diagramHierarchy.TaskOutputShapes {
			if shape.Task != nil && shape.Product != nil {
				diagramHierarchy.map_Task_TaskOutputShape[taskProductKey{Task: shape.Task, Product: shape.Product}] = shape
			}
		}

		diagramHierarchy.map_Product_ProductCompositionShape = make(map[*Product]*ProductCompositionShape)
		for _, shape := range diagramHierarchy.ProductComposition_Shapes {
			if shape.Product != nil {
				diagramHierarchy.map_Product_ProductCompositionShape[shape.Product] = shape
			}
		}

		diagramHierarchy.map_Task_TaskCompositionShape = make(map[*Task]*TaskCompositionShape)
		for _, shape := range diagramHierarchy.TaskComposition_Shapes {
			if shape.Task != nil {
				diagramHierarchy.map_Task_TaskCompositionShape[shape.Task] = shape
			}
		}

		diagramHierarchy.map_Note_NoteShape = make(map[*Note]*NoteShape)
		for _, shape := range diagramHierarchy.Note_Shapes {
			if shape.Note != nil {
				diagramHierarchy.map_Note_NoteShape[shape.Note] = shape
			}
		}

		diagramHierarchy.map_Note_NoteProductShape = make(map[noteProductKey]*NoteProductShape)
		for _, shape := range diagramHierarchy.NoteProductShapes {
			if shape.Note != nil {
				diagramHierarchy.map_Note_NoteProductShape[noteProductKey{Note: shape.Note, Product: shape.Product}] = shape
			}
		}

		diagramHierarchy.map_Note_NoteTaskShape = make(map[noteTaskKey]*NoteTaskShape)
		for _, shape := range diagramHierarchy.NoteTaskShapes {
			if shape.Note != nil && shape.Task != nil {
				diagramHierarchy.map_Note_NoteTaskShape[noteTaskKey{Note: shape.Note, Task: shape.Task}] = shape
			}
		}

		diagramHierarchy.map_Note_NoteResourceShape = make(map[noteResourceKey]*NoteResourceShape)
		for _, shape := range diagramHierarchy.NoteResourceShapes {
			if shape.Note != nil && shape.Resource != nil {
				diagramHierarchy.map_Note_NoteResourceShape[noteResourceKey{Note: shape.Note, Resource: shape.Resource}] = shape
			}
		}

		diagramHierarchy.map_Resource_ResourceShape = make(map[*Resource]*ResourceShape)
		for _, shape := range diagramHierarchy.Resource_Shapes {
			if shape.Resource != nil {
				diagramHierarchy.map_Resource_ResourceShape[shape.Resource] = shape
				diagrams := stager.map_Element_Diagrams[shape.Resource]

				if diagrams == nil {
					diagrams = []DiagramIF{diagramHierarchy}
				}

				if !slices.Contains(diagrams, DiagramIF(diagramHierarchy)) {
					diagrams = append(diagrams, diagramHierarchy)
				}
				stager.map_Element_Diagrams[shape.Resource] = diagrams
			}
		}

		diagramHierarchy.map_Resource_ResourceCompositionShape = make(map[*Resource]*ResourceCompositionShape)
		for _, shape := range diagramHierarchy.ResourceComposition_Shapes {
			if shape.Resource != nil {
				diagramHierarchy.map_Resource_ResourceCompositionShape[shape.Resource] = shape
			}
		}

		diagramHierarchy.map_Resource_ResourceTaskShape = make(map[resourceTaskKey]*ResourceTaskShape)
		for _, shape := range diagramHierarchy.ResourceTaskShapes {
			if shape.Resource != nil && shape.Task != nil {
				diagramHierarchy.map_Resource_ResourceTaskShape[resourceTaskKey{Resource: shape.Resource, Task: shape.Task}] = shape
			}
		}
	}
}
