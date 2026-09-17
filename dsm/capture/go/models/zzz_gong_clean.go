// generated code - do not edit
package models

import "time"

// CleanSlice is the Stage method that removes unstaged elements from a slice of pointers.
func (stage *Stage) CleanSlice[T PointerToGongstruct](slice *[]T) (modified bool) {
	if *slice == nil {
		return false
	}

	var cleanedSlice []T
	for _, element := range *slice {
		if stage.IsStaged(element) {
			cleanedSlice = append(cleanedSlice, element)
		}
	}
	modified = len(cleanedSlice) != len(*slice)
	if modified {
		*slice = cleanedSlice
	}
	return
}

// CleanPointer is the Stage method that sets the pointer to nil if the referenced element is not staged.
func (stage *Stage) CleanPointer[T PointerToGongstruct](element *T) (modified bool) {
	var zero T
	if *element == zero {
		return
	}

	if !stage.IsStaged(*element) {
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
	modified = stage.CleanSlice(&concept.Tools) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by ConceptShape
func (conceptshape *ConceptShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&conceptshape.Concept) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Concern
func (concern *Concern) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&concern.SubConcerns) || modified
	modified = stage.CleanSlice(&concern.Inputs) || modified
	modified = stage.CleanSlice(&concern.Outputs) || modified
	modified = stage.CleanSlice(&concern.Requirements) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by ConcernCompositionShape
func (concerncompositionshape *ConcernCompositionShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&concerncompositionshape.ControlPointShapes) || modified
	// insertion point per field
	modified = stage.CleanPointer(&concerncompositionshape.Concern) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ConcernInputShape
func (concerninputshape *ConcernInputShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&concerninputshape.ControlPointShapes) || modified
	// insertion point per field
	modified = stage.CleanPointer(&concerninputshape.Deliverable) || modified
	modified = stage.CleanPointer(&concerninputshape.Concern) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ConcernOutputShape
func (concernoutputshape *ConcernOutputShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&concernoutputshape.ControlPointShapes) || modified
	// insertion point per field
	modified = stage.CleanPointer(&concernoutputshape.Concern) || modified
	modified = stage.CleanPointer(&concernoutputshape.Deliverable) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ConcernShape
func (concernshape *ConcernShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&concernshape.Concern) || modified
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
	modified = stage.CleanSlice(&deliverable.SubDeliverables) || modified
	modified = stage.CleanSlice(&deliverable.Concepts) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by DeliverableCompositionShape
func (deliverablecompositionshape *DeliverableCompositionShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&deliverablecompositionshape.ControlPointShapes) || modified
	// insertion point per field
	modified = stage.CleanPointer(&deliverablecompositionshape.Deliverable) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DeliverableConceptShape
func (deliverableconceptshape *DeliverableConceptShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&deliverableconceptshape.ControlPointShapes) || modified
	// insertion point per field
	modified = stage.CleanPointer(&deliverableconceptshape.Deliverable) || modified
	modified = stage.CleanPointer(&deliverableconceptshape.Concept) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DeliverableShape
func (deliverableshape *DeliverableShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&deliverableshape.Deliverable) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Diagram
func (diagram *Diagram) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&diagram.ConcernsWhoseRequirementsNodeIsExpanded) || modified
	modified = stage.CleanSlice(&diagram.Deliverable_Shapes) || modified
	modified = stage.CleanSlice(&diagram.DeliverablesWhoseNodeIsExpanded) || modified
	modified = stage.CleanSlice(&diagram.DeliverablesWhoseConceptsNodeIsExpanded) || modified
	modified = stage.CleanSlice(&diagram.DeliverableComposition_Shapes) || modified
	modified = stage.CleanSlice(&diagram.Concern_Shapes) || modified
	modified = stage.CleanSlice(&diagram.ConcernsWhoseNodeIsExpanded) || modified
	modified = stage.CleanSlice(&diagram.ConcernsWhoseInputNodeIsExpanded) || modified
	modified = stage.CleanSlice(&diagram.ConcernsWhoseStakeholderNodeIsExpanded) || modified
	modified = stage.CleanSlice(&diagram.ConcernssWhoseOutputNodeIsExpanded) || modified
	modified = stage.CleanSlice(&diagram.ConcernComposition_Shapes) || modified
	modified = stage.CleanSlice(&diagram.ConcernInputShapes) || modified
	modified = stage.CleanSlice(&diagram.ConcernOutputShapes) || modified
	modified = stage.CleanSlice(&diagram.Note_Shapes) || modified
	modified = stage.CleanSlice(&diagram.NotesWhoseNodeIsExpanded) || modified
	modified = stage.CleanSlice(&diagram.NoteDeliverableShapes) || modified
	modified = stage.CleanSlice(&diagram.NoteTaskShapes) || modified
	modified = stage.CleanSlice(&diagram.NoteResourceShapes) || modified
	modified = stage.CleanSlice(&diagram.Stakeholder_Shapes) || modified
	modified = stage.CleanSlice(&diagram.ResourcesWhoseNodeIsExpanded) || modified
	modified = stage.CleanSlice(&diagram.ResourceComposition_Shapes) || modified
	modified = stage.CleanSlice(&diagram.StakeholderConcernShapes) || modified
	modified = stage.CleanSlice(&diagram.Requirement_Shapes) || modified
	modified = stage.CleanSlice(&diagram.RequirementsWhoseNodeIsExpanded) || modified
	modified = stage.CleanSlice(&diagram.Concept_Shapes) || modified
	modified = stage.CleanSlice(&diagram.ConceptsWhoseNodeIsExpanded) || modified
	modified = stage.CleanSlice(&diagram.ConceptsWhoseDeliverablesNodeIsExpanded) || modified
	modified = stage.CleanSlice(&diagram.DeliverableConceptShapes) || modified
	modified = stage.CleanSlice(&diagram.Diagram_Shapes) || modified
	modified = stage.CleanSlice(&diagram.DiagramsWhoseNodeIsExpanded) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by DiagramShape
func (diagramshape *DiagramShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&diagramshape.Diagram) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Library
func (library *Library) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&library.RootDeliverables) || modified
	modified = stage.CleanSlice(&library.RootConcerns) || modified
	modified = stage.CleanSlice(&library.RootStakeholders) || modified
	modified = stage.CleanSlice(&library.RootRequirements) || modified
	modified = stage.CleanSlice(&library.RootConcepts) || modified
	modified = stage.CleanSlice(&library.AnalysisNeeds) || modified
	modified = stage.CleanSlice(&library.Notes) || modified
	modified = stage.CleanSlice(&library.Diagrams) || modified
	modified = stage.CleanSlice(&library.SubLibraries) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Note
func (note *Note) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&note.Deliverables) || modified
	modified = stage.CleanSlice(&note.Tasks) || modified
	modified = stage.CleanSlice(&note.Resources) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by NoteDeliverableShape
func (notedeliverableshape *NoteDeliverableShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&notedeliverableshape.ControlPointShapes) || modified
	// insertion point per field
	modified = stage.CleanPointer(&notedeliverableshape.Note) || modified
	modified = stage.CleanPointer(&notedeliverableshape.Deliverable) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by NoteShape
func (noteshape *NoteShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&noteshape.Note) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by NoteStakeholderShape
func (notestakeholdershape *NoteStakeholderShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&notestakeholdershape.ControlPointShapes) || modified
	// insertion point per field
	modified = stage.CleanPointer(&notestakeholdershape.Note) || modified
	modified = stage.CleanPointer(&notestakeholdershape.Stakeholder) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by NoteTaskShape
func (notetaskshape *NoteTaskShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&notetaskshape.ControlPointShapes) || modified
	// insertion point per field
	modified = stage.CleanPointer(&notetaskshape.Note) || modified
	modified = stage.CleanPointer(&notetaskshape.Task) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Requirement
func (requirement *Requirement) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&requirement.SupportLevels) || modified
	modified = stage.CleanSlice(&requirement.Concepts) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by RequirementShape
func (requirementshape *RequirementShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&requirementshape.Requirement) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Stakeholder
func (stakeholder *Stakeholder) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&stakeholder.Concerns) || modified
	modified = stage.CleanSlice(&stakeholder.SubStakeholders) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by StakeholderCompositionShape
func (stakeholdercompositionshape *StakeholderCompositionShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&stakeholdercompositionshape.ControlPointShapes) || modified
	// insertion point per field
	modified = stage.CleanPointer(&stakeholdercompositionshape.Stakeholder) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by StakeholderConcernShape
func (stakeholderconcernshape *StakeholderConcernShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&stakeholderconcernshape.ControlPointShapes) || modified
	// insertion point per field
	modified = stage.CleanPointer(&stakeholderconcernshape.Stakeholder) || modified
	modified = stage.CleanPointer(&stakeholderconcernshape.Concern) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by StakeholderShape
func (stakeholdershape *StakeholderShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&stakeholdershape.Stakeholder) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by SupportLevel
func (supportlevel *SupportLevel) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&supportlevel.Tool) || modified
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
