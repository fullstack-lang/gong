package models

import "slices"

func (stager *Stager) enforceDiagramMaps() {
	stager.map_Element_Diagrams = make(map[AbstractType][]*Diagram)

	for _, diagram := range GetGongstrucsSorted[*Diagram](stager.stage) {

		// computes all products presents in the diagram
		diagram.map_Product_ProductShape = make(map[*Deliverable]*ProductShape)
		for _, shape := range diagram.Product_Shapes {
			if shape.Product != nil {
				diagram.map_Product_ProductShape[shape.Product] = shape
				var diagrams = stager.map_Element_Diagrams[shape.Product]

				if diagrams == nil {
					diagrams = []*Diagram{diagram}
				}

				if !slices.Contains(diagrams, diagram) {
					diagrams = append(diagrams, diagram)
				}
				stager.map_Element_Diagrams[shape.Product] = diagrams
			}
		}

		diagram.map_Concern_ConcernShape = make(map[*Concern]*ConcernShape)
		for _, shape := range diagram.Concern_Shapes {
			if shape.Concern != nil {
				diagram.map_Concern_ConcernShape[shape.Concern] = shape
				var diagrams = stager.map_Element_Diagrams[shape.Concern]

				if diagrams == nil {
					diagrams = []*Diagram{diagram}
				}

				if !slices.Contains(diagrams, diagram) {
					diagrams = append(diagrams, diagram)
				}
				stager.map_Element_Diagrams[shape.Concern] = diagrams
			}
		}

		diagram.map_Concern_TaskInputShape = make(map[concernProductKey]*ConcernInputShape)
		for _, shape := range diagram.ConcernInputShapes {
			if shape.Concern != nil && shape.Deliverable != nil {
				diagram.map_Concern_TaskInputShape[concernProductKey{Concern: shape.Concern, Product: shape.Deliverable}] = shape
			}
		}

		diagram.map_Concern_ConcernOutputShape = make(map[concernProductKey]*ConcernOutputShape)
		for _, shape := range diagram.ConcernOutputShapes {
			if shape.Task != nil && shape.Product != nil {
				diagram.map_Concern_ConcernOutputShape[concernProductKey{Concern: shape.Task, Product: shape.Product}] = shape
			}
		}

		diagram.map_Product_ProductCompositionShape = make(map[*Deliverable]*ProductCompositionShape)
		for _, shape := range diagram.ProductComposition_Shapes {
			if shape.Product != nil {
				diagram.map_Product_ProductCompositionShape[shape.Product] = shape
			}
		}

		diagram.map_Concern_ConcernCompositionShape = make(map[*Concern]*ConcernCompositionShape)
		for _, shape := range diagram.ConcernComposition_Shapes {
			if shape.Concern != nil {
				diagram.map_Concern_ConcernCompositionShape[shape.Concern] = shape
			}
		}

		diagram.map_Note_NoteShape = make(map[*Note]*NoteShape)
		for _, shape := range diagram.Note_Shapes {
			if shape.Note != nil {
				diagram.map_Note_NoteShape[shape.Note] = shape
			}
		}

		diagram.map_Note_NoteProductShape = make(map[noteProductKey]*NoteProductShape)
		for _, shape := range diagram.NoteProductShapes {
			if shape.Note != nil {
				diagram.map_Note_NoteProductShape[noteProductKey{Note: shape.Note, Product: shape.Product}] = shape
			}
		}

		diagram.map_Note_NoteTaskShape = make(map[noteTaskKey]*NoteTaskShape)
		for _, shape := range diagram.NoteTaskShapes {
			if shape.Note != nil && shape.Task != nil {
				diagram.map_Note_NoteTaskShape[noteTaskKey{Note: shape.Note, Task: shape.Task}] = shape
			}
		}

		diagram.map_Note_NoteResourceShape = make(map[noteResourceKey]*NoteStakeholderShape)
		for _, shape := range diagram.NoteResourceShapes {
			if shape.Note != nil && shape.Stakeholder != nil {
				diagram.map_Note_NoteResourceShape[noteResourceKey{Note: shape.Note, Stakeholder: shape.Stakeholder}] = shape
			}
		}

		diagram.map_Stakeholder_StakeholderShape = make(map[*Stakeholder]*StakeholderShape)
		for _, shape := range diagram.Stakeholder_Shapes {
			if shape.Stakeholder != nil {
				diagram.map_Stakeholder_StakeholderShape[shape.Stakeholder] = shape
				var diagrams = stager.map_Element_Diagrams[shape.Stakeholder]

				if diagrams == nil {
					diagrams = []*Diagram{diagram}
				}

				if !slices.Contains(diagrams, diagram) {
					diagrams = append(diagrams, diagram)
				}
				stager.map_Element_Diagrams[shape.Stakeholder] = diagrams
			}
		}

		diagram.map_Resource_ResourceCompositionShape = make(map[*Stakeholder]*StakeholderCompositionShape)
		for _, shape := range diagram.ResourceComposition_Shapes {
			if shape.Stakeholder != nil {
				diagram.map_Resource_ResourceCompositionShape[shape.Stakeholder] = shape
			}
		}

		diagram.map_StakeholderConcernKey_StakeholderConcernShape = make(map[stakeholderConcernKey]*StakeholderConcernShape)
		for _, shape := range diagram.StakeholderConcernShapes {
			if shape.Stakeholder != nil && shape.Concern != nil {
				diagram.map_StakeholderConcernKey_StakeholderConcernShape[stakeholderConcernKey{Stakeholder: shape.Stakeholder, Concern: shape.Concern}] = shape
			}
		}
	}
}
