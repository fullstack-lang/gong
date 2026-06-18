package models

import "slices"

func (stager *Stager) enforceDiagramMaps() {
	stager.map_Element_Diagrams = make(map[AbstractType][]*Diagram)

	for _, diagram := range GetGongstrucsSorted[*Diagram](stager.stage) {

		// computes all deliverables presents in the diagram
		diagram.map_Deliverable_DeliverableShape = make(map[*Deliverable]*DeliverableShape)
		for _, shape := range diagram.Deliverable_Shapes {
			if shape.Deliverable != nil {
				diagram.map_Deliverable_DeliverableShape[shape.Deliverable] = shape
				var diagrams = stager.map_Element_Diagrams[shape.Deliverable]

				if diagrams == nil {
					diagrams = []*Diagram{diagram}
				}

				if !slices.Contains(diagrams, diagram) {
					diagrams = append(diagrams, diagram)
				}
				stager.map_Element_Diagrams[shape.Deliverable] = diagrams
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

		diagram.map_Concern_TaskInputShape = make(map[concernDeliverableKey]*ConcernInputShape)
		for _, shape := range diagram.ConcernInputShapes {
			if shape.Concern != nil && shape.Deliverable != nil {
				diagram.map_Concern_TaskInputShape[concernDeliverableKey{Concern: shape.Concern, Deliverable: shape.Deliverable}] = shape
			}
		}

		diagram.map_Concern_ConcernOutputShape = make(map[concernDeliverableKey]*ConcernOutputShape)
		for _, shape := range diagram.ConcernOutputShapes {
			if shape.Task != nil && shape.Deliverable != nil {
				diagram.map_Concern_ConcernOutputShape[concernDeliverableKey{Concern: shape.Task, Deliverable: shape.Deliverable}] = shape
			}
		}

		diagram.map_Deliverable_DeliverableCompositionShape = make(map[*Deliverable]*DeliverableCompositionShape)
		for _, shape := range diagram.DeliverableComposition_Shapes {
			if shape.Deliverable != nil {
				diagram.map_Deliverable_DeliverableCompositionShape[shape.Deliverable] = shape
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

		diagram.map_Note_NoteDeliverableShape = make(map[noteDeliverableKey]*NoteDeliverableShape)
		for _, shape := range diagram.NoteDeliverableShapes {
			if shape.Note != nil {
				diagram.map_Note_NoteDeliverableShape[noteDeliverableKey{Note: shape.Note, Deliverable: shape.Deliverable}] = shape
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

		diagram.map_Requirement_RequirementShape = make(map[*Requirement]*RequirementShape)
		for _, shape := range diagram.Requirement_Shapes {
			if shape.Requirement != nil {
				diagram.map_Requirement_RequirementShape[shape.Requirement] = shape
				var diagrams = stager.map_Element_Diagrams[shape.Requirement]

				if diagrams == nil {
					diagrams = []*Diagram{diagram}
				}

				if !slices.Contains(diagrams, diagram) {
					diagrams = append(diagrams, diagram)
				}
				stager.map_Element_Diagrams[shape.Requirement] = diagrams
			}
		}

		diagram.map_Concept_ConceptShape = make(map[*Concept]*ConceptShape)
		for _, shape := range diagram.Concept_Shapes {
			if shape.Concept != nil {
				diagram.map_Concept_ConceptShape[shape.Concept] = shape
				var diagrams = stager.map_Element_Diagrams[shape.Concept]

				if diagrams == nil {
					diagrams = []*Diagram{diagram}
				}

				if !slices.Contains(diagrams, diagram) {
					diagrams = append(diagrams, diagram)
				}
				stager.map_Element_Diagrams[shape.Concept] = diagrams
			}
		}

		diagram.map_DeliverableConceptKey_DeliverableConceptShape = make(map[deliverableConceptKey]*DeliverableConceptShape)
		for _, shape := range diagram.DeliverableConceptShapes {
			if shape.Deliverable != nil && shape.Concept != nil {
				diagram.map_DeliverableConceptKey_DeliverableConceptShape[deliverableConceptKey{Deliverable: shape.Deliverable, Concept: shape.Concept}] = shape
			}
		}
	}
}
