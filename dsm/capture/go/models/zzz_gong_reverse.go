// generated code - do not edit
package models

// insertion point
func (inst *AnalysisNeed) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Library":
		switch reverseField.Fieldname {
		case "AnalysisNeeds":
			if _library, ok := stage.Library_AnalysisNeeds_reverseMap[inst]; ok {
				res = _library.Name
			}
		}
	}
	return
}

func (inst *Concept) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Deliverable":
		switch reverseField.Fieldname {
		case "Concepts":
			if _deliverable, ok := stage.Deliverable_Concepts_reverseMap[inst]; ok {
				res = _deliverable.Name
			}
		}
	case "Diagram":
		switch reverseField.Fieldname {
		case "ConceptsWhoseNodeIsExpanded":
			if _diagram, ok := stage.Diagram_ConceptsWhoseNodeIsExpanded_reverseMap[inst]; ok {
				res = _diagram.Name
			}
		case "ConceptsWhoseDeliverablesNodeIsExpanded":
			if _diagram, ok := stage.Diagram_ConceptsWhoseDeliverablesNodeIsExpanded_reverseMap[inst]; ok {
				res = _diagram.Name
			}
		}
	case "Library":
		switch reverseField.Fieldname {
		case "RootConcepts":
			if _library, ok := stage.Library_RootConcepts_reverseMap[inst]; ok {
				res = _library.Name
			}
		}
	case "Requirement":
		switch reverseField.Fieldname {
		case "Concepts":
			if _requirement, ok := stage.Requirement_Concepts_reverseMap[inst]; ok {
				res = _requirement.Name
			}
		}
	}
	return
}

func (inst *ConceptShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Diagram":
		switch reverseField.Fieldname {
		case "Concept_Shapes":
			if _diagram, ok := stage.Diagram_Concept_Shapes_reverseMap[inst]; ok {
				res = _diagram.Name
			}
		}
	}
	return
}

func (inst *Concern) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Concern":
		switch reverseField.Fieldname {
		case "SubConcerns":
			if _concern, ok := stage.Concern_SubConcerns_reverseMap[inst]; ok {
				res = _concern.Name
			}
		}
	case "Diagram":
		switch reverseField.Fieldname {
		case "ConcernsWhoseRequirementsNodeIsExpanded":
			if _diagram, ok := stage.Diagram_ConcernsWhoseRequirementsNodeIsExpanded_reverseMap[inst]; ok {
				res = _diagram.Name
			}
		case "ConcernsWhoseNodeIsExpanded":
			if _diagram, ok := stage.Diagram_ConcernsWhoseNodeIsExpanded_reverseMap[inst]; ok {
				res = _diagram.Name
			}
		case "ConcernsWhoseInputNodeIsExpanded":
			if _diagram, ok := stage.Diagram_ConcernsWhoseInputNodeIsExpanded_reverseMap[inst]; ok {
				res = _diagram.Name
			}
		case "ConcernsWhoseStakeholderNodeIsExpanded":
			if _diagram, ok := stage.Diagram_ConcernsWhoseStakeholderNodeIsExpanded_reverseMap[inst]; ok {
				res = _diagram.Name
			}
		case "ConcernssWhoseOutputNodeIsExpanded":
			if _diagram, ok := stage.Diagram_ConcernssWhoseOutputNodeIsExpanded_reverseMap[inst]; ok {
				res = _diagram.Name
			}
		}
	case "Library":
		switch reverseField.Fieldname {
		case "RootConcerns":
			if _library, ok := stage.Library_RootConcerns_reverseMap[inst]; ok {
				res = _library.Name
			}
		}
	case "Note":
		switch reverseField.Fieldname {
		case "Tasks":
			if _note, ok := stage.Note_Tasks_reverseMap[inst]; ok {
				res = _note.Name
			}
		}
	case "Stakeholder":
		switch reverseField.Fieldname {
		case "Concerns":
			if _stakeholder, ok := stage.Stakeholder_Concerns_reverseMap[inst]; ok {
				res = _stakeholder.Name
			}
		}
	}
	return
}

func (inst *ConcernCompositionShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Diagram":
		switch reverseField.Fieldname {
		case "ConcernComposition_Shapes":
			if _diagram, ok := stage.Diagram_ConcernComposition_Shapes_reverseMap[inst]; ok {
				res = _diagram.Name
			}
		}
	}
	return
}

func (inst *ConcernInputShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Diagram":
		switch reverseField.Fieldname {
		case "ConcernInputShapes":
			if _diagram, ok := stage.Diagram_ConcernInputShapes_reverseMap[inst]; ok {
				res = _diagram.Name
			}
		}
	}
	return
}

func (inst *ConcernOutputShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Diagram":
		switch reverseField.Fieldname {
		case "ConcernOutputShapes":
			if _diagram, ok := stage.Diagram_ConcernOutputShapes_reverseMap[inst]; ok {
				res = _diagram.Name
			}
		}
	}
	return
}

func (inst *ConcernShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Diagram":
		switch reverseField.Fieldname {
		case "Concern_Shapes":
			if _diagram, ok := stage.Diagram_Concern_Shapes_reverseMap[inst]; ok {
				res = _diagram.Name
			}
		}
	}
	return
}

func (inst *ControlPointShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ConcernCompositionShape":
		switch reverseField.Fieldname {
		case "ControlPointShapes":
			if _concerncompositionshape, ok := stage.ConcernCompositionShape_ControlPointShapes_reverseMap[inst]; ok {
				res = _concerncompositionshape.Name
			}
		}
	case "ConcernInputShape":
		switch reverseField.Fieldname {
		case "ControlPointShapes":
			if _concerninputshape, ok := stage.ConcernInputShape_ControlPointShapes_reverseMap[inst]; ok {
				res = _concerninputshape.Name
			}
		}
	case "ConcernOutputShape":
		switch reverseField.Fieldname {
		case "ControlPointShapes":
			if _concernoutputshape, ok := stage.ConcernOutputShape_ControlPointShapes_reverseMap[inst]; ok {
				res = _concernoutputshape.Name
			}
		}
	case "DeliverableCompositionShape":
		switch reverseField.Fieldname {
		case "ControlPointShapes":
			if _deliverablecompositionshape, ok := stage.DeliverableCompositionShape_ControlPointShapes_reverseMap[inst]; ok {
				res = _deliverablecompositionshape.Name
			}
		}
	case "DeliverableConceptShape":
		switch reverseField.Fieldname {
		case "ControlPointShapes":
			if _deliverableconceptshape, ok := stage.DeliverableConceptShape_ControlPointShapes_reverseMap[inst]; ok {
				res = _deliverableconceptshape.Name
			}
		}
	case "NoteDeliverableShape":
		switch reverseField.Fieldname {
		case "ControlPointShapes":
			if _notedeliverableshape, ok := stage.NoteDeliverableShape_ControlPointShapes_reverseMap[inst]; ok {
				res = _notedeliverableshape.Name
			}
		}
	case "NoteStakeholderShape":
		switch reverseField.Fieldname {
		case "ControlPointShapes":
			if _notestakeholdershape, ok := stage.NoteStakeholderShape_ControlPointShapes_reverseMap[inst]; ok {
				res = _notestakeholdershape.Name
			}
		}
	case "NoteTaskShape":
		switch reverseField.Fieldname {
		case "ControlPointShapes":
			if _notetaskshape, ok := stage.NoteTaskShape_ControlPointShapes_reverseMap[inst]; ok {
				res = _notetaskshape.Name
			}
		}
	case "StakeholderCompositionShape":
		switch reverseField.Fieldname {
		case "ControlPointShapes":
			if _stakeholdercompositionshape, ok := stage.StakeholderCompositionShape_ControlPointShapes_reverseMap[inst]; ok {
				res = _stakeholdercompositionshape.Name
			}
		}
	case "StakeholderConcernShape":
		switch reverseField.Fieldname {
		case "ControlPointShapes":
			if _stakeholderconcernshape, ok := stage.StakeholderConcernShape_ControlPointShapes_reverseMap[inst]; ok {
				res = _stakeholderconcernshape.Name
			}
		}
	}
	return
}

func (inst *Deliverable) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Concern":
		switch reverseField.Fieldname {
		case "Inputs":
			if _concern, ok := stage.Concern_Inputs_reverseMap[inst]; ok {
				res = _concern.Name
			}
		case "Outputs":
			if _concern, ok := stage.Concern_Outputs_reverseMap[inst]; ok {
				res = _concern.Name
			}
		}
	case "Deliverable":
		switch reverseField.Fieldname {
		case "SubDeliverables":
			if _deliverable, ok := stage.Deliverable_SubDeliverables_reverseMap[inst]; ok {
				res = _deliverable.Name
			}
		}
	case "Diagram":
		switch reverseField.Fieldname {
		case "DeliverablesWhoseNodeIsExpanded":
			if _diagram, ok := stage.Diagram_DeliverablesWhoseNodeIsExpanded_reverseMap[inst]; ok {
				res = _diagram.Name
			}
		case "DeliverablesWhoseConceptsNodeIsExpanded":
			if _diagram, ok := stage.Diagram_DeliverablesWhoseConceptsNodeIsExpanded_reverseMap[inst]; ok {
				res = _diagram.Name
			}
		}
	case "Library":
		switch reverseField.Fieldname {
		case "RootDeliverables":
			if _library, ok := stage.Library_RootDeliverables_reverseMap[inst]; ok {
				res = _library.Name
			}
		}
	case "Note":
		switch reverseField.Fieldname {
		case "Deliverables":
			if _note, ok := stage.Note_Deliverables_reverseMap[inst]; ok {
				res = _note.Name
			}
		}
	}
	return
}

func (inst *DeliverableCompositionShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Diagram":
		switch reverseField.Fieldname {
		case "DeliverableComposition_Shapes":
			if _diagram, ok := stage.Diagram_DeliverableComposition_Shapes_reverseMap[inst]; ok {
				res = _diagram.Name
			}
		}
	}
	return
}

func (inst *DeliverableConceptShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Diagram":
		switch reverseField.Fieldname {
		case "DeliverableConceptShapes":
			if _diagram, ok := stage.Diagram_DeliverableConceptShapes_reverseMap[inst]; ok {
				res = _diagram.Name
			}
		}
	}
	return
}

func (inst *DeliverableShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Diagram":
		switch reverseField.Fieldname {
		case "Deliverable_Shapes":
			if _diagram, ok := stage.Diagram_Deliverable_Shapes_reverseMap[inst]; ok {
				res = _diagram.Name
			}
		}
	}
	return
}

func (inst *Diagram) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Diagram":
		switch reverseField.Fieldname {
		case "DiagramsWhoseNodeIsExpanded":
			if _diagram, ok := stage.Diagram_DiagramsWhoseNodeIsExpanded_reverseMap[inst]; ok {
				res = _diagram.Name
			}
		}
	case "Library":
		switch reverseField.Fieldname {
		case "Diagrams":
			if _library, ok := stage.Library_Diagrams_reverseMap[inst]; ok {
				res = _library.Name
			}
		}
	}
	return
}

func (inst *DiagramShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Diagram":
		switch reverseField.Fieldname {
		case "Diagram_Shapes":
			if _diagram, ok := stage.Diagram_Diagram_Shapes_reverseMap[inst]; ok {
				res = _diagram.Name
			}
		}
	}
	return
}

func (inst *Library) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Library":
		switch reverseField.Fieldname {
		case "SubLibraries":
			if _library, ok := stage.Library_SubLibraries_reverseMap[inst]; ok {
				res = _library.Name
			}
		}
	}
	return
}

func (inst *Note) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Diagram":
		switch reverseField.Fieldname {
		case "NotesWhoseNodeIsExpanded":
			if _diagram, ok := stage.Diagram_NotesWhoseNodeIsExpanded_reverseMap[inst]; ok {
				res = _diagram.Name
			}
		}
	case "Library":
		switch reverseField.Fieldname {
		case "Notes":
			if _library, ok := stage.Library_Notes_reverseMap[inst]; ok {
				res = _library.Name
			}
		}
	}
	return
}

func (inst *NoteDeliverableShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Diagram":
		switch reverseField.Fieldname {
		case "NoteDeliverableShapes":
			if _diagram, ok := stage.Diagram_NoteDeliverableShapes_reverseMap[inst]; ok {
				res = _diagram.Name
			}
		}
	}
	return
}

func (inst *NoteShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Diagram":
		switch reverseField.Fieldname {
		case "Note_Shapes":
			if _diagram, ok := stage.Diagram_Note_Shapes_reverseMap[inst]; ok {
				res = _diagram.Name
			}
		}
	}
	return
}

func (inst *NoteStakeholderShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Diagram":
		switch reverseField.Fieldname {
		case "NoteResourceShapes":
			if _diagram, ok := stage.Diagram_NoteResourceShapes_reverseMap[inst]; ok {
				res = _diagram.Name
			}
		}
	}
	return
}

func (inst *NoteTaskShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Diagram":
		switch reverseField.Fieldname {
		case "NoteTaskShapes":
			if _diagram, ok := stage.Diagram_NoteTaskShapes_reverseMap[inst]; ok {
				res = _diagram.Name
			}
		}
	}
	return
}

func (inst *Requirement) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Concern":
		switch reverseField.Fieldname {
		case "Requirements":
			if _concern, ok := stage.Concern_Requirements_reverseMap[inst]; ok {
				res = _concern.Name
			}
		}
	case "Diagram":
		switch reverseField.Fieldname {
		case "RequirementsWhoseNodeIsExpanded":
			if _diagram, ok := stage.Diagram_RequirementsWhoseNodeIsExpanded_reverseMap[inst]; ok {
				res = _diagram.Name
			}
		}
	case "Library":
		switch reverseField.Fieldname {
		case "RootRequirements":
			if _library, ok := stage.Library_RootRequirements_reverseMap[inst]; ok {
				res = _library.Name
			}
		}
	}
	return
}

func (inst *RequirementShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Diagram":
		switch reverseField.Fieldname {
		case "Requirement_Shapes":
			if _diagram, ok := stage.Diagram_Requirement_Shapes_reverseMap[inst]; ok {
				res = _diagram.Name
			}
		}
	}
	return
}

func (inst *Stakeholder) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Diagram":
		switch reverseField.Fieldname {
		case "ResourcesWhoseNodeIsExpanded":
			if _diagram, ok := stage.Diagram_ResourcesWhoseNodeIsExpanded_reverseMap[inst]; ok {
				res = _diagram.Name
			}
		}
	case "Library":
		switch reverseField.Fieldname {
		case "RootStakeholders":
			if _library, ok := stage.Library_RootStakeholders_reverseMap[inst]; ok {
				res = _library.Name
			}
		}
	case "Note":
		switch reverseField.Fieldname {
		case "Resources":
			if _note, ok := stage.Note_Resources_reverseMap[inst]; ok {
				res = _note.Name
			}
		}
	case "Stakeholder":
		switch reverseField.Fieldname {
		case "SubStakeholders":
			if _stakeholder, ok := stage.Stakeholder_SubStakeholders_reverseMap[inst]; ok {
				res = _stakeholder.Name
			}
		}
	}
	return
}

func (inst *StakeholderCompositionShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Diagram":
		switch reverseField.Fieldname {
		case "ResourceComposition_Shapes":
			if _diagram, ok := stage.Diagram_ResourceComposition_Shapes_reverseMap[inst]; ok {
				res = _diagram.Name
			}
		}
	}
	return
}

func (inst *StakeholderConcernShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Diagram":
		switch reverseField.Fieldname {
		case "StakeholderConcernShapes":
			if _diagram, ok := stage.Diagram_StakeholderConcernShapes_reverseMap[inst]; ok {
				res = _diagram.Name
			}
		}
	}
	return
}

func (inst *StakeholderShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Diagram":
		switch reverseField.Fieldname {
		case "Stakeholder_Shapes":
			if _diagram, ok := stage.Diagram_Stakeholder_Shapes_reverseMap[inst]; ok {
				res = _diagram.Name
			}
		}
	}
	return
}

func (inst *SupportLevel) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Requirement":
		switch reverseField.Fieldname {
		case "SupportLevels":
			if _requirement, ok := stage.Requirement_SupportLevels_reverseMap[inst]; ok {
				res = _requirement.Name
			}
		}
	}
	return
}

func (inst *Tool) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Concept":
		switch reverseField.Fieldname {
		case "Tools":
			if _concept, ok := stage.Concept_Tools_reverseMap[inst]; ok {
				res = _concept.Name
			}
		}
	}
	return
}
