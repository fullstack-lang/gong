// generated code - do not edit
package models

import "time"

// GongCleanSlice removes unstaged elements from a slice of pointers of type T.
// T must be a pointer to a struct that implements PointerToGongstruct.
func GongCleanSlice[T PointerToGongstruct](stage *Stage, slice *[]T) (modified bool) {
	if *slice == nil {
		return false
	}

	var cleanedSlice []T
	for _, element := range *slice {
		if IsStagedPointerToGongstruct(stage, element) {
			cleanedSlice = append(cleanedSlice, element)
		}
	}
	modified = len(cleanedSlice) != len(*slice)
	if modified {
		*slice = cleanedSlice
	}
	return
}

// GongCleanPointer sets the pointer to nil if the referenced element is not staged.
// T must be a pointer to a struct that implements PointerToGongstruct.
func GongCleanPointer[T PointerToGongstruct](stage *Stage, element *T) (modified bool) {
	var zero T
	if *element == zero {
		return
	}

	if !IsStagedPointerToGongstruct(stage, *element) {
		*element = zero
		modified = true
		return
	}
	return
}

// insertion point per named struct
// Clean garbage collect unstaged instances that are referenced by AnalysisNeed
func (analysisneed *AnalysisNeed) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Concept
func (concept *Concept) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &concept.Tools) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by ConceptShape
func (conceptshape *ConceptShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &conceptshape.Concept) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Concern
func (concern *Concern) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &concern.SubConcerns) || modified
	modified = GongCleanSlice(stage, &concern.Inputs) || modified
	modified = GongCleanSlice(stage, &concern.Outputs) || modified
	modified = GongCleanSlice(stage, &concern.Requirements) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by ConcernCompositionShape
func (concerncompositionshape *ConcernCompositionShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &concerncompositionshape.ControlPointShapes) || modified
	// insertion point per field
	modified = GongCleanPointer(stage, &concerncompositionshape.Concern) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ConcernInputShape
func (concerninputshape *ConcernInputShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &concerninputshape.ControlPointShapes) || modified
	// insertion point per field
	modified = GongCleanPointer(stage, &concerninputshape.Deliverable) || modified
	modified = GongCleanPointer(stage, &concerninputshape.Concern) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ConcernOutputShape
func (concernoutputshape *ConcernOutputShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &concernoutputshape.ControlPointShapes) || modified
	// insertion point per field
	modified = GongCleanPointer(stage, &concernoutputshape.Concern) || modified
	modified = GongCleanPointer(stage, &concernoutputshape.Deliverable) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ConcernShape
func (concernshape *ConcernShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &concernshape.Concern) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ControlPointShape
func (controlpointshape *ControlPointShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Deliverable
func (deliverable *Deliverable) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &deliverable.SubDeliverables) || modified
	modified = GongCleanSlice(stage, &deliverable.Concepts) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by DeliverableCompositionShape
func (deliverablecompositionshape *DeliverableCompositionShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &deliverablecompositionshape.ControlPointShapes) || modified
	// insertion point per field
	modified = GongCleanPointer(stage, &deliverablecompositionshape.Deliverable) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DeliverableConceptShape
func (deliverableconceptshape *DeliverableConceptShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &deliverableconceptshape.ControlPointShapes) || modified
	// insertion point per field
	modified = GongCleanPointer(stage, &deliverableconceptshape.Deliverable) || modified
	modified = GongCleanPointer(stage, &deliverableconceptshape.Concept) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DeliverableShape
func (deliverableshape *DeliverableShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &deliverableshape.Deliverable) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Diagram
func (diagram *Diagram) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &diagram.ConcernsWhoseRequirementsNodeIsExpanded) || modified
	modified = GongCleanSlice(stage, &diagram.Deliverable_Shapes) || modified
	modified = GongCleanSlice(stage, &diagram.DeliverablesWhoseNodeIsExpanded) || modified
	modified = GongCleanSlice(stage, &diagram.DeliverablesWhoseConceptsNodeIsExpanded) || modified
	modified = GongCleanSlice(stage, &diagram.DeliverableComposition_Shapes) || modified
	modified = GongCleanSlice(stage, &diagram.Concern_Shapes) || modified
	modified = GongCleanSlice(stage, &diagram.ConcernsWhoseNodeIsExpanded) || modified
	modified = GongCleanSlice(stage, &diagram.ConcernsWhoseInputNodeIsExpanded) || modified
	modified = GongCleanSlice(stage, &diagram.ConcernsWhoseStakeholderNodeIsExpanded) || modified
	modified = GongCleanSlice(stage, &diagram.ConcernssWhoseOutputNodeIsExpanded) || modified
	modified = GongCleanSlice(stage, &diagram.ConcernComposition_Shapes) || modified
	modified = GongCleanSlice(stage, &diagram.ConcernInputShapes) || modified
	modified = GongCleanSlice(stage, &diagram.ConcernOutputShapes) || modified
	modified = GongCleanSlice(stage, &diagram.Note_Shapes) || modified
	modified = GongCleanSlice(stage, &diagram.NotesWhoseNodeIsExpanded) || modified
	modified = GongCleanSlice(stage, &diagram.NoteDeliverableShapes) || modified
	modified = GongCleanSlice(stage, &diagram.NoteTaskShapes) || modified
	modified = GongCleanSlice(stage, &diagram.NoteResourceShapes) || modified
	modified = GongCleanSlice(stage, &diagram.Stakeholder_Shapes) || modified
	modified = GongCleanSlice(stage, &diagram.ResourcesWhoseNodeIsExpanded) || modified
	modified = GongCleanSlice(stage, &diagram.ResourceComposition_Shapes) || modified
	modified = GongCleanSlice(stage, &diagram.StakeholderConcernShapes) || modified
	modified = GongCleanSlice(stage, &diagram.Requirement_Shapes) || modified
	modified = GongCleanSlice(stage, &diagram.RequirementsWhoseNodeIsExpanded) || modified
	modified = GongCleanSlice(stage, &diagram.Concept_Shapes) || modified
	modified = GongCleanSlice(stage, &diagram.ConceptsWhoseNodeIsExpanded) || modified
	modified = GongCleanSlice(stage, &diagram.ConceptsWhoseDeliverablesNodeIsExpanded) || modified
	modified = GongCleanSlice(stage, &diagram.DeliverableConceptShapes) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Library
func (library *Library) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &library.RootDeliverables) || modified
	modified = GongCleanSlice(stage, &library.RootConcerns) || modified
	modified = GongCleanSlice(stage, &library.RootStakeholders) || modified
	modified = GongCleanSlice(stage, &library.RootRequirements) || modified
	modified = GongCleanSlice(stage, &library.RootConcepts) || modified
	modified = GongCleanSlice(stage, &library.AnalysisNeeds) || modified
	modified = GongCleanSlice(stage, &library.Notes) || modified
	modified = GongCleanSlice(stage, &library.Diagrams) || modified
	modified = GongCleanSlice(stage, &library.SubLibraries) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Note
func (note *Note) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &note.Deliverables) || modified
	modified = GongCleanSlice(stage, &note.Tasks) || modified
	modified = GongCleanSlice(stage, &note.Resources) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by NoteDeliverableShape
func (notedeliverableshape *NoteDeliverableShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &notedeliverableshape.ControlPointShapes) || modified
	// insertion point per field
	modified = GongCleanPointer(stage, &notedeliverableshape.Note) || modified
	modified = GongCleanPointer(stage, &notedeliverableshape.Deliverable) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by NoteShape
func (noteshape *NoteShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &noteshape.Note) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by NoteStakeholderShape
func (notestakeholdershape *NoteStakeholderShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &notestakeholdershape.ControlPointShapes) || modified
	// insertion point per field
	modified = GongCleanPointer(stage, &notestakeholdershape.Note) || modified
	modified = GongCleanPointer(stage, &notestakeholdershape.Stakeholder) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by NoteTaskShape
func (notetaskshape *NoteTaskShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &notetaskshape.ControlPointShapes) || modified
	// insertion point per field
	modified = GongCleanPointer(stage, &notetaskshape.Note) || modified
	modified = GongCleanPointer(stage, &notetaskshape.Task) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Requirement
func (requirement *Requirement) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &requirement.SupportLevels) || modified
	modified = GongCleanSlice(stage, &requirement.Concepts) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by RequirementShape
func (requirementshape *RequirementShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &requirementshape.Requirement) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Stakeholder
func (stakeholder *Stakeholder) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &stakeholder.Concerns) || modified
	modified = GongCleanSlice(stage, &stakeholder.SubStakeholders) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by StakeholderCompositionShape
func (stakeholdercompositionshape *StakeholderCompositionShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &stakeholdercompositionshape.ControlPointShapes) || modified
	// insertion point per field
	modified = GongCleanPointer(stage, &stakeholdercompositionshape.Stakeholder) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by StakeholderConcernShape
func (stakeholderconcernshape *StakeholderConcernShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &stakeholderconcernshape.ControlPointShapes) || modified
	// insertion point per field
	modified = GongCleanPointer(stage, &stakeholderconcernshape.Stakeholder) || modified
	modified = GongCleanPointer(stage, &stakeholderconcernshape.Concern) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by StakeholderShape
func (stakeholdershape *StakeholderShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &stakeholdershape.Stakeholder) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by SupportLevel
func (supportlevel *SupportLevel) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &supportlevel.Tool) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Tool
func (tool *Tool) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by staged elements
func (stage *Stage) Clean() (modified bool) {
	for _, instance := range stage.GetInstances() {
		modified = instance.GongClean(stage) || modified
	}
	if modified {
		if stage.probeIF != nil {
			stage.probeIF.AddNotification(time.Now(), "Stage clean generated a modification")
		}
	}
	return
}
