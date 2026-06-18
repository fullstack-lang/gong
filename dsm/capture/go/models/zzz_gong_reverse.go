// generated code - do not edit
package models

// insertion point
func (inst *AnalysisNeed) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Concept) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *ConceptShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Concern) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *ConcernCompositionShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *ConcernInputShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *ConcernOutputShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *ConcernShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *ControlPointShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Deliverable) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *DeliverableCompositionShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *DeliverableConceptShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *DeliverableShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Diagram) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
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

func (inst *Library) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Note) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *NoteDeliverableShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *NoteShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *NoteStakeholderShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *NoteTaskShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Requirement) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *RequirementShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Stakeholder) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *StakeholderCompositionShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *StakeholderConcernShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *StakeholderShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *SupportLevel) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Tool) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

// insertion point
func (inst *AnalysisNeed) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Library":
		switch reverseField.Fieldname {
		case "AnalysisNeeds":
			res = stage.Library_AnalysisNeeds_reverseMap[inst]
		}
	}
	return res
}

func (inst *Concept) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Deliverable":
		switch reverseField.Fieldname {
		case "Concepts":
			res = stage.Deliverable_Concepts_reverseMap[inst]
		}
	case "Diagram":
		switch reverseField.Fieldname {
		case "ConceptsWhoseNodeIsExpanded":
			res = stage.Diagram_ConceptsWhoseNodeIsExpanded_reverseMap[inst]
		case "ConceptsWhoseDeliverablesNodeIsExpanded":
			res = stage.Diagram_ConceptsWhoseDeliverablesNodeIsExpanded_reverseMap[inst]
		}
	case "Library":
		switch reverseField.Fieldname {
		case "RootConcepts":
			res = stage.Library_RootConcepts_reverseMap[inst]
		}
	case "Requirement":
		switch reverseField.Fieldname {
		case "Concepts":
			res = stage.Requirement_Concepts_reverseMap[inst]
		}
	}
	return res
}

func (inst *ConceptShape) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Diagram":
		switch reverseField.Fieldname {
		case "Concept_Shapes":
			res = stage.Diagram_Concept_Shapes_reverseMap[inst]
		}
	}
	return res
}

func (inst *Concern) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Concern":
		switch reverseField.Fieldname {
		case "SubConcerns":
			res = stage.Concern_SubConcerns_reverseMap[inst]
		}
	case "Diagram":
		switch reverseField.Fieldname {
		case "ConcernsWhoseRequirementsNodeIsExpanded":
			res = stage.Diagram_ConcernsWhoseRequirementsNodeIsExpanded_reverseMap[inst]
		case "ConcernsWhoseNodeIsExpanded":
			res = stage.Diagram_ConcernsWhoseNodeIsExpanded_reverseMap[inst]
		case "ConcernsWhoseInputNodeIsExpanded":
			res = stage.Diagram_ConcernsWhoseInputNodeIsExpanded_reverseMap[inst]
		case "ConcernsWhoseStakeholderNodeIsExpanded":
			res = stage.Diagram_ConcernsWhoseStakeholderNodeIsExpanded_reverseMap[inst]
		case "ConcernssWhoseOutputNodeIsExpanded":
			res = stage.Diagram_ConcernssWhoseOutputNodeIsExpanded_reverseMap[inst]
		}
	case "Library":
		switch reverseField.Fieldname {
		case "RootConcerns":
			res = stage.Library_RootConcerns_reverseMap[inst]
		}
	case "Note":
		switch reverseField.Fieldname {
		case "Tasks":
			res = stage.Note_Tasks_reverseMap[inst]
		}
	case "Stakeholder":
		switch reverseField.Fieldname {
		case "Concerns":
			res = stage.Stakeholder_Concerns_reverseMap[inst]
		}
	}
	return res
}

func (inst *ConcernCompositionShape) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Diagram":
		switch reverseField.Fieldname {
		case "ConcernComposition_Shapes":
			res = stage.Diagram_ConcernComposition_Shapes_reverseMap[inst]
		}
	}
	return res
}

func (inst *ConcernInputShape) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Diagram":
		switch reverseField.Fieldname {
		case "ConcernInputShapes":
			res = stage.Diagram_ConcernInputShapes_reverseMap[inst]
		}
	}
	return res
}

func (inst *ConcernOutputShape) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Diagram":
		switch reverseField.Fieldname {
		case "ConcernOutputShapes":
			res = stage.Diagram_ConcernOutputShapes_reverseMap[inst]
		}
	}
	return res
}

func (inst *ConcernShape) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Diagram":
		switch reverseField.Fieldname {
		case "Concern_Shapes":
			res = stage.Diagram_Concern_Shapes_reverseMap[inst]
		}
	}
	return res
}

func (inst *ControlPointShape) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "ConcernCompositionShape":
		switch reverseField.Fieldname {
		case "ControlPointShapes":
			res = stage.ConcernCompositionShape_ControlPointShapes_reverseMap[inst]
		}
	case "ConcernInputShape":
		switch reverseField.Fieldname {
		case "ControlPointShapes":
			res = stage.ConcernInputShape_ControlPointShapes_reverseMap[inst]
		}
	case "ConcernOutputShape":
		switch reverseField.Fieldname {
		case "ControlPointShapes":
			res = stage.ConcernOutputShape_ControlPointShapes_reverseMap[inst]
		}
	case "DeliverableCompositionShape":
		switch reverseField.Fieldname {
		case "ControlPointShapes":
			res = stage.DeliverableCompositionShape_ControlPointShapes_reverseMap[inst]
		}
	case "DeliverableConceptShape":
		switch reverseField.Fieldname {
		case "ControlPointShapes":
			res = stage.DeliverableConceptShape_ControlPointShapes_reverseMap[inst]
		}
	case "NoteDeliverableShape":
		switch reverseField.Fieldname {
		case "ControlPointShapes":
			res = stage.NoteDeliverableShape_ControlPointShapes_reverseMap[inst]
		}
	case "NoteStakeholderShape":
		switch reverseField.Fieldname {
		case "ControlPointShapes":
			res = stage.NoteStakeholderShape_ControlPointShapes_reverseMap[inst]
		}
	case "NoteTaskShape":
		switch reverseField.Fieldname {
		case "ControlPointShapes":
			res = stage.NoteTaskShape_ControlPointShapes_reverseMap[inst]
		}
	case "StakeholderCompositionShape":
		switch reverseField.Fieldname {
		case "ControlPointShapes":
			res = stage.StakeholderCompositionShape_ControlPointShapes_reverseMap[inst]
		}
	case "StakeholderConcernShape":
		switch reverseField.Fieldname {
		case "ControlPointShapes":
			res = stage.StakeholderConcernShape_ControlPointShapes_reverseMap[inst]
		}
	}
	return res
}

func (inst *Deliverable) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Concern":
		switch reverseField.Fieldname {
		case "Inputs":
			res = stage.Concern_Inputs_reverseMap[inst]
		case "Outputs":
			res = stage.Concern_Outputs_reverseMap[inst]
		}
	case "Deliverable":
		switch reverseField.Fieldname {
		case "SubDeliverables":
			res = stage.Deliverable_SubDeliverables_reverseMap[inst]
		}
	case "Diagram":
		switch reverseField.Fieldname {
		case "DeliverablesWhoseNodeIsExpanded":
			res = stage.Diagram_DeliverablesWhoseNodeIsExpanded_reverseMap[inst]
		case "DeliverablesWhoseConceptsNodeIsExpanded":
			res = stage.Diagram_DeliverablesWhoseConceptsNodeIsExpanded_reverseMap[inst]
		}
	case "Library":
		switch reverseField.Fieldname {
		case "RootDeliverables":
			res = stage.Library_RootDeliverables_reverseMap[inst]
		}
	case "Note":
		switch reverseField.Fieldname {
		case "Deliverables":
			res = stage.Note_Deliverables_reverseMap[inst]
		}
	}
	return res
}

func (inst *DeliverableCompositionShape) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Diagram":
		switch reverseField.Fieldname {
		case "DeliverableComposition_Shapes":
			res = stage.Diagram_DeliverableComposition_Shapes_reverseMap[inst]
		}
	}
	return res
}

func (inst *DeliverableConceptShape) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Diagram":
		switch reverseField.Fieldname {
		case "DeliverableConceptShapes":
			res = stage.Diagram_DeliverableConceptShapes_reverseMap[inst]
		}
	}
	return res
}

func (inst *DeliverableShape) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Diagram":
		switch reverseField.Fieldname {
		case "Deliverable_Shapes":
			res = stage.Diagram_Deliverable_Shapes_reverseMap[inst]
		}
	}
	return res
}

func (inst *Diagram) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Library":
		switch reverseField.Fieldname {
		case "Diagrams":
			res = stage.Library_Diagrams_reverseMap[inst]
		}
	}
	return res
}

func (inst *Library) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Library":
		switch reverseField.Fieldname {
		case "SubLibraries":
			res = stage.Library_SubLibraries_reverseMap[inst]
		}
	}
	return res
}

func (inst *Note) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Diagram":
		switch reverseField.Fieldname {
		case "NotesWhoseNodeIsExpanded":
			res = stage.Diagram_NotesWhoseNodeIsExpanded_reverseMap[inst]
		}
	case "Library":
		switch reverseField.Fieldname {
		case "Notes":
			res = stage.Library_Notes_reverseMap[inst]
		}
	}
	return res
}

func (inst *NoteDeliverableShape) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Diagram":
		switch reverseField.Fieldname {
		case "NoteDeliverableShapes":
			res = stage.Diagram_NoteDeliverableShapes_reverseMap[inst]
		}
	}
	return res
}

func (inst *NoteShape) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Diagram":
		switch reverseField.Fieldname {
		case "Note_Shapes":
			res = stage.Diagram_Note_Shapes_reverseMap[inst]
		}
	}
	return res
}

func (inst *NoteStakeholderShape) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Diagram":
		switch reverseField.Fieldname {
		case "NoteResourceShapes":
			res = stage.Diagram_NoteResourceShapes_reverseMap[inst]
		}
	}
	return res
}

func (inst *NoteTaskShape) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Diagram":
		switch reverseField.Fieldname {
		case "NoteTaskShapes":
			res = stage.Diagram_NoteTaskShapes_reverseMap[inst]
		}
	}
	return res
}

func (inst *Requirement) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Concern":
		switch reverseField.Fieldname {
		case "Requirements":
			res = stage.Concern_Requirements_reverseMap[inst]
		}
	case "Diagram":
		switch reverseField.Fieldname {
		case "RequirementsWhoseNodeIsExpanded":
			res = stage.Diagram_RequirementsWhoseNodeIsExpanded_reverseMap[inst]
		}
	case "Library":
		switch reverseField.Fieldname {
		case "RootRequirements":
			res = stage.Library_RootRequirements_reverseMap[inst]
		}
	}
	return res
}

func (inst *RequirementShape) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Diagram":
		switch reverseField.Fieldname {
		case "Requirement_Shapes":
			res = stage.Diagram_Requirement_Shapes_reverseMap[inst]
		}
	}
	return res
}

func (inst *Stakeholder) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Diagram":
		switch reverseField.Fieldname {
		case "ResourcesWhoseNodeIsExpanded":
			res = stage.Diagram_ResourcesWhoseNodeIsExpanded_reverseMap[inst]
		}
	case "Library":
		switch reverseField.Fieldname {
		case "RootStakeholders":
			res = stage.Library_RootStakeholders_reverseMap[inst]
		}
	case "Note":
		switch reverseField.Fieldname {
		case "Resources":
			res = stage.Note_Resources_reverseMap[inst]
		}
	case "Stakeholder":
		switch reverseField.Fieldname {
		case "SubStakeholders":
			res = stage.Stakeholder_SubStakeholders_reverseMap[inst]
		}
	}
	return res
}

func (inst *StakeholderCompositionShape) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Diagram":
		switch reverseField.Fieldname {
		case "ResourceComposition_Shapes":
			res = stage.Diagram_ResourceComposition_Shapes_reverseMap[inst]
		}
	}
	return res
}

func (inst *StakeholderConcernShape) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Diagram":
		switch reverseField.Fieldname {
		case "StakeholderConcernShapes":
			res = stage.Diagram_StakeholderConcernShapes_reverseMap[inst]
		}
	}
	return res
}

func (inst *StakeholderShape) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Diagram":
		switch reverseField.Fieldname {
		case "Stakeholder_Shapes":
			res = stage.Diagram_Stakeholder_Shapes_reverseMap[inst]
		}
	}
	return res
}

func (inst *SupportLevel) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Requirement":
		switch reverseField.Fieldname {
		case "SupportLevels":
			res = stage.Requirement_SupportLevels_reverseMap[inst]
		}
	}
	return res
}

func (inst *Tool) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Concept":
		switch reverseField.Fieldname {
		case "Tools":
			res = stage.Concept_Tools_reverseMap[inst]
		}
	}
	return res
}
