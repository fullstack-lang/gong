// generated code - do not edit
package models

import (
	"fmt"
	"slices"
)

// IsStaged is the Stage method checking if a gongstruct instance is staged.
func (stage *Stage) IsStaged(instance GongstructIF) (ok bool) {
	if instance != nil {
		return instance.GongIsStaged(stage)
	}
	return false
}

// insertion point for stage per struct
func (analysisneed *AnalysisNeed) GongIsStaged(stage *Stage) bool {
	_, ok := stage.AnalysisNeeds[analysisneed]
	return ok
}

func (concept *Concept) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Concepts[concept]
	return ok
}

func (conceptshape *ConceptShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ConceptShapes[conceptshape]
	return ok
}

func (concern *Concern) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Concerns[concern]
	return ok
}

func (concerncompositionshape *ConcernCompositionShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ConcernCompositionShapes[concerncompositionshape]
	return ok
}

func (concerninputshape *ConcernInputShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ConcernInputShapes[concerninputshape]
	return ok
}

func (concernoutputshape *ConcernOutputShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ConcernOutputShapes[concernoutputshape]
	return ok
}

func (concernshape *ConcernShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ConcernShapes[concernshape]
	return ok
}

func (controlpointshape *ControlPointShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ControlPointShapes[controlpointshape]
	return ok
}

func (deliverable *Deliverable) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Deliverables[deliverable]
	return ok
}

func (deliverablecompositionshape *DeliverableCompositionShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.DeliverableCompositionShapes[deliverablecompositionshape]
	return ok
}

func (deliverableconceptshape *DeliverableConceptShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.DeliverableConceptShapes[deliverableconceptshape]
	return ok
}

func (deliverableshape *DeliverableShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.DeliverableShapes[deliverableshape]
	return ok
}

func (diagram *Diagram) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Diagrams[diagram]
	return ok
}

func (diagramshape *DiagramShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.DiagramShapes[diagramshape]
	return ok
}

func (library *Library) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Librarys[library]
	return ok
}

func (note *Note) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Notes[note]
	return ok
}

func (notedeliverableshape *NoteDeliverableShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.NoteDeliverableShapes[notedeliverableshape]
	return ok
}

func (noteshape *NoteShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.NoteShapes[noteshape]
	return ok
}

func (notestakeholdershape *NoteStakeholderShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.NoteStakeholderShapes[notestakeholdershape]
	return ok
}

func (notetaskshape *NoteTaskShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.NoteTaskShapes[notetaskshape]
	return ok
}

func (requirement *Requirement) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Requirements[requirement]
	return ok
}

func (requirementshape *RequirementShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.RequirementShapes[requirementshape]
	return ok
}

func (stakeholder *Stakeholder) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Stakeholders[stakeholder]
	return ok
}

func (stakeholdercompositionshape *StakeholderCompositionShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.StakeholderCompositionShapes[stakeholdercompositionshape]
	return ok
}

func (stakeholderconcernshape *StakeholderConcernShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.StakeholderConcernShapes[stakeholderconcernshape]
	return ok
}

func (stakeholdershape *StakeholderShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.StakeholderShapes[stakeholdershape]
	return ok
}

func (supportlevel *SupportLevel) GongIsStaged(stage *Stage) bool {
	_, ok := stage.SupportLevels[supportlevel]
	return ok
}

func (tool *Tool) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Tools[tool]
	return ok
}

// StageBranch is the Stage method that stages instance and applies StageBranch recursively.
func (stage *Stage) StageBranch(instance GongstructIF) {
	if instance != nil {
		instance.GongStageBranch(stage)
	}
}

// insertion point for stage branch per struct
func (analysisneed *AnalysisNeed) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(analysisneed) {
		return
	}

	analysisneed.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (concept *Concept) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(concept) {
		return
	}

	concept.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _tool := range concept.Tools {
		stage.StageBranch(_tool)
	}

}

func (conceptshape *ConceptShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(conceptshape) {
		return
	}

	conceptshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if conceptshape.Concept != nil {
		stage.StageBranch(conceptshape.Concept)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (concern *Concern) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(concern) {
		return
	}

	concern.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _concern := range concern.SubConcerns {
		stage.StageBranch(_concern)
	}
	for _, _deliverable := range concern.Inputs {
		stage.StageBranch(_deliverable)
	}
	for _, _deliverable := range concern.Outputs {
		stage.StageBranch(_deliverable)
	}
	for _, _requirement := range concern.Requirements {
		stage.StageBranch(_requirement)
	}

}

func (concerncompositionshape *ConcernCompositionShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(concerncompositionshape) {
		return
	}

	concerncompositionshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if concerncompositionshape.Concern != nil {
		stage.StageBranch(concerncompositionshape.Concern)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _controlpointshape := range concerncompositionshape.ControlPointShapes {
		stage.StageBranch(_controlpointshape)
	}

}

func (concerninputshape *ConcernInputShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(concerninputshape) {
		return
	}

	concerninputshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if concerninputshape.Deliverable != nil {
		stage.StageBranch(concerninputshape.Deliverable)
	}
	if concerninputshape.Concern != nil {
		stage.StageBranch(concerninputshape.Concern)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _controlpointshape := range concerninputshape.ControlPointShapes {
		stage.StageBranch(_controlpointshape)
	}

}

func (concernoutputshape *ConcernOutputShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(concernoutputshape) {
		return
	}

	concernoutputshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if concernoutputshape.Concern != nil {
		stage.StageBranch(concernoutputshape.Concern)
	}
	if concernoutputshape.Deliverable != nil {
		stage.StageBranch(concernoutputshape.Deliverable)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _controlpointshape := range concernoutputshape.ControlPointShapes {
		stage.StageBranch(_controlpointshape)
	}

}

func (concernshape *ConcernShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(concernshape) {
		return
	}

	concernshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if concernshape.Concern != nil {
		stage.StageBranch(concernshape.Concern)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (controlpointshape *ControlPointShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(controlpointshape) {
		return
	}

	controlpointshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (deliverable *Deliverable) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(deliverable) {
		return
	}

	deliverable.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _deliverable := range deliverable.SubDeliverables {
		stage.StageBranch(_deliverable)
	}
	for _, _concept := range deliverable.Concepts {
		stage.StageBranch(_concept)
	}

}

func (deliverablecompositionshape *DeliverableCompositionShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(deliverablecompositionshape) {
		return
	}

	deliverablecompositionshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if deliverablecompositionshape.Deliverable != nil {
		stage.StageBranch(deliverablecompositionshape.Deliverable)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _controlpointshape := range deliverablecompositionshape.ControlPointShapes {
		stage.StageBranch(_controlpointshape)
	}

}

func (deliverableconceptshape *DeliverableConceptShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(deliverableconceptshape) {
		return
	}

	deliverableconceptshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if deliverableconceptshape.Deliverable != nil {
		stage.StageBranch(deliverableconceptshape.Deliverable)
	}
	if deliverableconceptshape.Concept != nil {
		stage.StageBranch(deliverableconceptshape.Concept)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _controlpointshape := range deliverableconceptshape.ControlPointShapes {
		stage.StageBranch(_controlpointshape)
	}

}

func (deliverableshape *DeliverableShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(deliverableshape) {
		return
	}

	deliverableshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if deliverableshape.Deliverable != nil {
		stage.StageBranch(deliverableshape.Deliverable)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (diagram *Diagram) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(diagram) {
		return
	}

	diagram.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _concern := range diagram.ConcernsWhoseRequirementsNodeIsExpanded {
		stage.StageBranch(_concern)
	}
	for _, _deliverableshape := range diagram.Deliverable_Shapes {
		stage.StageBranch(_deliverableshape)
	}
	for _, _deliverable := range diagram.DeliverablesWhoseNodeIsExpanded {
		stage.StageBranch(_deliverable)
	}
	for _, _deliverable := range diagram.DeliverablesWhoseConceptsNodeIsExpanded {
		stage.StageBranch(_deliverable)
	}
	for _, _deliverablecompositionshape := range diagram.DeliverableComposition_Shapes {
		stage.StageBranch(_deliverablecompositionshape)
	}
	for _, _concernshape := range diagram.Concern_Shapes {
		stage.StageBranch(_concernshape)
	}
	for _, _concern := range diagram.ConcernsWhoseNodeIsExpanded {
		stage.StageBranch(_concern)
	}
	for _, _concern := range diagram.ConcernsWhoseInputNodeIsExpanded {
		stage.StageBranch(_concern)
	}
	for _, _concern := range diagram.ConcernsWhoseStakeholderNodeIsExpanded {
		stage.StageBranch(_concern)
	}
	for _, _concern := range diagram.ConcernssWhoseOutputNodeIsExpanded {
		stage.StageBranch(_concern)
	}
	for _, _concerncompositionshape := range diagram.ConcernComposition_Shapes {
		stage.StageBranch(_concerncompositionshape)
	}
	for _, _concerninputshape := range diagram.ConcernInputShapes {
		stage.StageBranch(_concerninputshape)
	}
	for _, _concernoutputshape := range diagram.ConcernOutputShapes {
		stage.StageBranch(_concernoutputshape)
	}
	for _, _noteshape := range diagram.Note_Shapes {
		stage.StageBranch(_noteshape)
	}
	for _, _note := range diagram.NotesWhoseNodeIsExpanded {
		stage.StageBranch(_note)
	}
	for _, _notedeliverableshape := range diagram.NoteDeliverableShapes {
		stage.StageBranch(_notedeliverableshape)
	}
	for _, _notetaskshape := range diagram.NoteTaskShapes {
		stage.StageBranch(_notetaskshape)
	}
	for _, _notestakeholdershape := range diagram.NoteResourceShapes {
		stage.StageBranch(_notestakeholdershape)
	}
	for _, _stakeholdershape := range diagram.Stakeholder_Shapes {
		stage.StageBranch(_stakeholdershape)
	}
	for _, _stakeholder := range diagram.ResourcesWhoseNodeIsExpanded {
		stage.StageBranch(_stakeholder)
	}
	for _, _stakeholdercompositionshape := range diagram.ResourceComposition_Shapes {
		stage.StageBranch(_stakeholdercompositionshape)
	}
	for _, _stakeholderconcernshape := range diagram.StakeholderConcernShapes {
		stage.StageBranch(_stakeholderconcernshape)
	}
	for _, _requirementshape := range diagram.Requirement_Shapes {
		stage.StageBranch(_requirementshape)
	}
	for _, _requirement := range diagram.RequirementsWhoseNodeIsExpanded {
		stage.StageBranch(_requirement)
	}
	for _, _conceptshape := range diagram.Concept_Shapes {
		stage.StageBranch(_conceptshape)
	}
	for _, _concept := range diagram.ConceptsWhoseNodeIsExpanded {
		stage.StageBranch(_concept)
	}
	for _, _concept := range diagram.ConceptsWhoseDeliverablesNodeIsExpanded {
		stage.StageBranch(_concept)
	}
	for _, _deliverableconceptshape := range diagram.DeliverableConceptShapes {
		stage.StageBranch(_deliverableconceptshape)
	}
	for _, _diagramshape := range diagram.Diagram_Shapes {
		stage.StageBranch(_diagramshape)
	}
	for _, _diagram := range diagram.DiagramsWhoseNodeIsExpanded {
		stage.StageBranch(_diagram)
	}

}

func (diagramshape *DiagramShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(diagramshape) {
		return
	}

	diagramshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if diagramshape.Diagram != nil {
		stage.StageBranch(diagramshape.Diagram)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (library *Library) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(library) {
		return
	}

	library.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _deliverable := range library.RootDeliverables {
		stage.StageBranch(_deliverable)
	}
	for _, _concern := range library.RootConcerns {
		stage.StageBranch(_concern)
	}
	for _, _stakeholder := range library.RootStakeholders {
		stage.StageBranch(_stakeholder)
	}
	for _, _requirement := range library.RootRequirements {
		stage.StageBranch(_requirement)
	}
	for _, _concept := range library.RootConcepts {
		stage.StageBranch(_concept)
	}
	for _, _analysisneed := range library.AnalysisNeeds {
		stage.StageBranch(_analysisneed)
	}
	for _, _note := range library.Notes {
		stage.StageBranch(_note)
	}
	for _, _diagram := range library.Diagrams {
		stage.StageBranch(_diagram)
	}
	for _, _library := range library.SubLibraries {
		stage.StageBranch(_library)
	}

}

func (note *Note) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(note) {
		return
	}

	note.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _deliverable := range note.Deliverables {
		stage.StageBranch(_deliverable)
	}
	for _, _concern := range note.Tasks {
		stage.StageBranch(_concern)
	}
	for _, _stakeholder := range note.Resources {
		stage.StageBranch(_stakeholder)
	}

}

func (notedeliverableshape *NoteDeliverableShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(notedeliverableshape) {
		return
	}

	notedeliverableshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if notedeliverableshape.Note != nil {
		stage.StageBranch(notedeliverableshape.Note)
	}
	if notedeliverableshape.Deliverable != nil {
		stage.StageBranch(notedeliverableshape.Deliverable)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _controlpointshape := range notedeliverableshape.ControlPointShapes {
		stage.StageBranch(_controlpointshape)
	}

}

func (noteshape *NoteShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(noteshape) {
		return
	}

	noteshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if noteshape.Note != nil {
		stage.StageBranch(noteshape.Note)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (notestakeholdershape *NoteStakeholderShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(notestakeholdershape) {
		return
	}

	notestakeholdershape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if notestakeholdershape.Note != nil {
		stage.StageBranch(notestakeholdershape.Note)
	}
	if notestakeholdershape.Stakeholder != nil {
		stage.StageBranch(notestakeholdershape.Stakeholder)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _controlpointshape := range notestakeholdershape.ControlPointShapes {
		stage.StageBranch(_controlpointshape)
	}

}

func (notetaskshape *NoteTaskShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(notetaskshape) {
		return
	}

	notetaskshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if notetaskshape.Note != nil {
		stage.StageBranch(notetaskshape.Note)
	}
	if notetaskshape.Task != nil {
		stage.StageBranch(notetaskshape.Task)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _controlpointshape := range notetaskshape.ControlPointShapes {
		stage.StageBranch(_controlpointshape)
	}

}

func (requirement *Requirement) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(requirement) {
		return
	}

	requirement.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _supportlevel := range requirement.SupportLevels {
		stage.StageBranch(_supportlevel)
	}
	for _, _concept := range requirement.Concepts {
		stage.StageBranch(_concept)
	}

}

func (requirementshape *RequirementShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(requirementshape) {
		return
	}

	requirementshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if requirementshape.Requirement != nil {
		stage.StageBranch(requirementshape.Requirement)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stakeholder *Stakeholder) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(stakeholder) {
		return
	}

	stakeholder.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _concern := range stakeholder.Concerns {
		stage.StageBranch(_concern)
	}
	for _, _stakeholder := range stakeholder.SubStakeholders {
		stage.StageBranch(_stakeholder)
	}

}

func (stakeholdercompositionshape *StakeholderCompositionShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(stakeholdercompositionshape) {
		return
	}

	stakeholdercompositionshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if stakeholdercompositionshape.Stakeholder != nil {
		stage.StageBranch(stakeholdercompositionshape.Stakeholder)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _controlpointshape := range stakeholdercompositionshape.ControlPointShapes {
		stage.StageBranch(_controlpointshape)
	}

}

func (stakeholderconcernshape *StakeholderConcernShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(stakeholderconcernshape) {
		return
	}

	stakeholderconcernshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if stakeholderconcernshape.Stakeholder != nil {
		stage.StageBranch(stakeholderconcernshape.Stakeholder)
	}
	if stakeholderconcernshape.Concern != nil {
		stage.StageBranch(stakeholderconcernshape.Concern)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _controlpointshape := range stakeholderconcernshape.ControlPointShapes {
		stage.StageBranch(_controlpointshape)
	}

}

func (stakeholdershape *StakeholderShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(stakeholdershape) {
		return
	}

	stakeholdershape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if stakeholdershape.Stakeholder != nil {
		stage.StageBranch(stakeholdershape.Stakeholder)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (supportlevel *SupportLevel) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(supportlevel) {
		return
	}

	supportlevel.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if supportlevel.Tool != nil {
		stage.StageBranch(supportlevel.Tool)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (tool *Tool) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(tool) {
		return
	}

	tool.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

// GongCopyBranch stages instance and apply GongCopyBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func GongCopyBranch[Type Gongstruct](from *Type) (to *Type) {

	mapOrigCopy := make(map[any]any)
	_ = mapOrigCopy

	switch fromT := any(from).(type) {
	// insertion point for stage branch
	case *AnalysisNeed:
		toT := GongCopyBranchAnalysisNeed(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Concept:
		toT := GongCopyBranchConcept(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ConceptShape:
		toT := GongCopyBranchConceptShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Concern:
		toT := GongCopyBranchConcern(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ConcernCompositionShape:
		toT := GongCopyBranchConcernCompositionShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ConcernInputShape:
		toT := GongCopyBranchConcernInputShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ConcernOutputShape:
		toT := GongCopyBranchConcernOutputShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ConcernShape:
		toT := GongCopyBranchConcernShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ControlPointShape:
		toT := GongCopyBranchControlPointShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Deliverable:
		toT := GongCopyBranchDeliverable(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *DeliverableCompositionShape:
		toT := GongCopyBranchDeliverableCompositionShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *DeliverableConceptShape:
		toT := GongCopyBranchDeliverableConceptShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *DeliverableShape:
		toT := GongCopyBranchDeliverableShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Diagram:
		toT := GongCopyBranchDiagram(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *DiagramShape:
		toT := GongCopyBranchDiagramShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Library:
		toT := GongCopyBranchLibrary(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Note:
		toT := GongCopyBranchNote(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *NoteDeliverableShape:
		toT := GongCopyBranchNoteDeliverableShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *NoteShape:
		toT := GongCopyBranchNoteShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *NoteStakeholderShape:
		toT := GongCopyBranchNoteStakeholderShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *NoteTaskShape:
		toT := GongCopyBranchNoteTaskShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Requirement:
		toT := GongCopyBranchRequirement(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *RequirementShape:
		toT := GongCopyBranchRequirementShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Stakeholder:
		toT := GongCopyBranchStakeholder(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *StakeholderCompositionShape:
		toT := GongCopyBranchStakeholderCompositionShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *StakeholderConcernShape:
		toT := GongCopyBranchStakeholderConcernShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *StakeholderShape:
		toT := GongCopyBranchStakeholderShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *SupportLevel:
		toT := GongCopyBranchSupportLevel(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Tool:
		toT := GongCopyBranchTool(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
func GongCopyBranchAnalysisNeed(mapOrigCopy map[any]any, analysisneedFrom *AnalysisNeed) (analysisneedTo *AnalysisNeed) {
	var alreadyCopied bool
	analysisneedTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, analysisneedFrom)
	if alreadyCopied {
		return
	}
	analysisneedFrom.GongCopyBasicFields(analysisneedTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchConcept(mapOrigCopy map[any]any, conceptFrom *Concept) (conceptTo *Concept) {
	var alreadyCopied bool
	conceptTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, conceptFrom)
	if alreadyCopied {
		return
	}
	conceptFrom.GongCopyBasicFields(conceptTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _tool := range conceptFrom.Tools {
		conceptTo.Tools = append(conceptTo.Tools, GongCopyBranchTool(mapOrigCopy, _tool))
	}

	return
}

func GongCopyBranchConceptShape(mapOrigCopy map[any]any, conceptshapeFrom *ConceptShape) (conceptshapeTo *ConceptShape) {
	var alreadyCopied bool
	conceptshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, conceptshapeFrom)
	if alreadyCopied {
		return
	}
	conceptshapeFrom.GongCopyBasicFields(conceptshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if conceptshapeFrom.Concept != nil {
		conceptshapeTo.Concept = GongCopyBranchConcept(mapOrigCopy, conceptshapeFrom.Concept)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchConcern(mapOrigCopy map[any]any, concernFrom *Concern) (concernTo *Concern) {
	var alreadyCopied bool
	concernTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, concernFrom)
	if alreadyCopied {
		return
	}
	concernFrom.GongCopyBasicFields(concernTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _concern := range concernFrom.SubConcerns {
		concernTo.SubConcerns = append(concernTo.SubConcerns, GongCopyBranchConcern(mapOrigCopy, _concern))
	}
	for _, _deliverable := range concernFrom.Inputs {
		concernTo.Inputs = append(concernTo.Inputs, GongCopyBranchDeliverable(mapOrigCopy, _deliverable))
	}
	for _, _deliverable := range concernFrom.Outputs {
		concernTo.Outputs = append(concernTo.Outputs, GongCopyBranchDeliverable(mapOrigCopy, _deliverable))
	}
	for _, _requirement := range concernFrom.Requirements {
		concernTo.Requirements = append(concernTo.Requirements, GongCopyBranchRequirement(mapOrigCopy, _requirement))
	}

	return
}

func GongCopyBranchConcernCompositionShape(mapOrigCopy map[any]any, concerncompositionshapeFrom *ConcernCompositionShape) (concerncompositionshapeTo *ConcernCompositionShape) {
	var alreadyCopied bool
	concerncompositionshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, concerncompositionshapeFrom)
	if alreadyCopied {
		return
	}
	concerncompositionshapeFrom.GongCopyBasicFields(concerncompositionshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if concerncompositionshapeFrom.Concern != nil {
		concerncompositionshapeTo.Concern = GongCopyBranchConcern(mapOrigCopy, concerncompositionshapeFrom.Concern)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _controlpointshape := range concerncompositionshapeFrom.ControlPointShapes {
		concerncompositionshapeTo.ControlPointShapes = append(concerncompositionshapeTo.ControlPointShapes, GongCopyBranchControlPointShape(mapOrigCopy, _controlpointshape))
	}

	return
}

func GongCopyBranchConcernInputShape(mapOrigCopy map[any]any, concerninputshapeFrom *ConcernInputShape) (concerninputshapeTo *ConcernInputShape) {
	var alreadyCopied bool
	concerninputshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, concerninputshapeFrom)
	if alreadyCopied {
		return
	}
	concerninputshapeFrom.GongCopyBasicFields(concerninputshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if concerninputshapeFrom.Deliverable != nil {
		concerninputshapeTo.Deliverable = GongCopyBranchDeliverable(mapOrigCopy, concerninputshapeFrom.Deliverable)
	}
	if concerninputshapeFrom.Concern != nil {
		concerninputshapeTo.Concern = GongCopyBranchConcern(mapOrigCopy, concerninputshapeFrom.Concern)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _controlpointshape := range concerninputshapeFrom.ControlPointShapes {
		concerninputshapeTo.ControlPointShapes = append(concerninputshapeTo.ControlPointShapes, GongCopyBranchControlPointShape(mapOrigCopy, _controlpointshape))
	}

	return
}

func GongCopyBranchConcernOutputShape(mapOrigCopy map[any]any, concernoutputshapeFrom *ConcernOutputShape) (concernoutputshapeTo *ConcernOutputShape) {
	var alreadyCopied bool
	concernoutputshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, concernoutputshapeFrom)
	if alreadyCopied {
		return
	}
	concernoutputshapeFrom.GongCopyBasicFields(concernoutputshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if concernoutputshapeFrom.Concern != nil {
		concernoutputshapeTo.Concern = GongCopyBranchConcern(mapOrigCopy, concernoutputshapeFrom.Concern)
	}
	if concernoutputshapeFrom.Deliverable != nil {
		concernoutputshapeTo.Deliverable = GongCopyBranchDeliverable(mapOrigCopy, concernoutputshapeFrom.Deliverable)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _controlpointshape := range concernoutputshapeFrom.ControlPointShapes {
		concernoutputshapeTo.ControlPointShapes = append(concernoutputshapeTo.ControlPointShapes, GongCopyBranchControlPointShape(mapOrigCopy, _controlpointshape))
	}

	return
}

func GongCopyBranchConcernShape(mapOrigCopy map[any]any, concernshapeFrom *ConcernShape) (concernshapeTo *ConcernShape) {
	var alreadyCopied bool
	concernshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, concernshapeFrom)
	if alreadyCopied {
		return
	}
	concernshapeFrom.GongCopyBasicFields(concernshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if concernshapeFrom.Concern != nil {
		concernshapeTo.Concern = GongCopyBranchConcern(mapOrigCopy, concernshapeFrom.Concern)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchControlPointShape(mapOrigCopy map[any]any, controlpointshapeFrom *ControlPointShape) (controlpointshapeTo *ControlPointShape) {
	var alreadyCopied bool
	controlpointshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, controlpointshapeFrom)
	if alreadyCopied {
		return
	}
	controlpointshapeFrom.GongCopyBasicFields(controlpointshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchDeliverable(mapOrigCopy map[any]any, deliverableFrom *Deliverable) (deliverableTo *Deliverable) {
	var alreadyCopied bool
	deliverableTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, deliverableFrom)
	if alreadyCopied {
		return
	}
	deliverableFrom.GongCopyBasicFields(deliverableTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _deliverable := range deliverableFrom.SubDeliverables {
		deliverableTo.SubDeliverables = append(deliverableTo.SubDeliverables, GongCopyBranchDeliverable(mapOrigCopy, _deliverable))
	}
	for _, _concept := range deliverableFrom.Concepts {
		deliverableTo.Concepts = append(deliverableTo.Concepts, GongCopyBranchConcept(mapOrigCopy, _concept))
	}

	return
}

func GongCopyBranchDeliverableCompositionShape(mapOrigCopy map[any]any, deliverablecompositionshapeFrom *DeliverableCompositionShape) (deliverablecompositionshapeTo *DeliverableCompositionShape) {
	var alreadyCopied bool
	deliverablecompositionshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, deliverablecompositionshapeFrom)
	if alreadyCopied {
		return
	}
	deliverablecompositionshapeFrom.GongCopyBasicFields(deliverablecompositionshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if deliverablecompositionshapeFrom.Deliverable != nil {
		deliverablecompositionshapeTo.Deliverable = GongCopyBranchDeliverable(mapOrigCopy, deliverablecompositionshapeFrom.Deliverable)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _controlpointshape := range deliverablecompositionshapeFrom.ControlPointShapes {
		deliverablecompositionshapeTo.ControlPointShapes = append(deliverablecompositionshapeTo.ControlPointShapes, GongCopyBranchControlPointShape(mapOrigCopy, _controlpointshape))
	}

	return
}

func GongCopyBranchDeliverableConceptShape(mapOrigCopy map[any]any, deliverableconceptshapeFrom *DeliverableConceptShape) (deliverableconceptshapeTo *DeliverableConceptShape) {
	var alreadyCopied bool
	deliverableconceptshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, deliverableconceptshapeFrom)
	if alreadyCopied {
		return
	}
	deliverableconceptshapeFrom.GongCopyBasicFields(deliverableconceptshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if deliverableconceptshapeFrom.Deliverable != nil {
		deliverableconceptshapeTo.Deliverable = GongCopyBranchDeliverable(mapOrigCopy, deliverableconceptshapeFrom.Deliverable)
	}
	if deliverableconceptshapeFrom.Concept != nil {
		deliverableconceptshapeTo.Concept = GongCopyBranchConcept(mapOrigCopy, deliverableconceptshapeFrom.Concept)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _controlpointshape := range deliverableconceptshapeFrom.ControlPointShapes {
		deliverableconceptshapeTo.ControlPointShapes = append(deliverableconceptshapeTo.ControlPointShapes, GongCopyBranchControlPointShape(mapOrigCopy, _controlpointshape))
	}

	return
}

func GongCopyBranchDeliverableShape(mapOrigCopy map[any]any, deliverableshapeFrom *DeliverableShape) (deliverableshapeTo *DeliverableShape) {
	var alreadyCopied bool
	deliverableshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, deliverableshapeFrom)
	if alreadyCopied {
		return
	}
	deliverableshapeFrom.GongCopyBasicFields(deliverableshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if deliverableshapeFrom.Deliverable != nil {
		deliverableshapeTo.Deliverable = GongCopyBranchDeliverable(mapOrigCopy, deliverableshapeFrom.Deliverable)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchDiagram(mapOrigCopy map[any]any, diagramFrom *Diagram) (diagramTo *Diagram) {
	var alreadyCopied bool
	diagramTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, diagramFrom)
	if alreadyCopied {
		return
	}
	diagramFrom.GongCopyBasicFields(diagramTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _concern := range diagramFrom.ConcernsWhoseRequirementsNodeIsExpanded {
		diagramTo.ConcernsWhoseRequirementsNodeIsExpanded = append(diagramTo.ConcernsWhoseRequirementsNodeIsExpanded, GongCopyBranchConcern(mapOrigCopy, _concern))
	}
	for _, _deliverableshape := range diagramFrom.Deliverable_Shapes {
		diagramTo.Deliverable_Shapes = append(diagramTo.Deliverable_Shapes, GongCopyBranchDeliverableShape(mapOrigCopy, _deliverableshape))
	}
	for _, _deliverable := range diagramFrom.DeliverablesWhoseNodeIsExpanded {
		diagramTo.DeliverablesWhoseNodeIsExpanded = append(diagramTo.DeliverablesWhoseNodeIsExpanded, GongCopyBranchDeliverable(mapOrigCopy, _deliverable))
	}
	for _, _deliverable := range diagramFrom.DeliverablesWhoseConceptsNodeIsExpanded {
		diagramTo.DeliverablesWhoseConceptsNodeIsExpanded = append(diagramTo.DeliverablesWhoseConceptsNodeIsExpanded, GongCopyBranchDeliverable(mapOrigCopy, _deliverable))
	}
	for _, _deliverablecompositionshape := range diagramFrom.DeliverableComposition_Shapes {
		diagramTo.DeliverableComposition_Shapes = append(diagramTo.DeliverableComposition_Shapes, GongCopyBranchDeliverableCompositionShape(mapOrigCopy, _deliverablecompositionshape))
	}
	for _, _concernshape := range diagramFrom.Concern_Shapes {
		diagramTo.Concern_Shapes = append(diagramTo.Concern_Shapes, GongCopyBranchConcernShape(mapOrigCopy, _concernshape))
	}
	for _, _concern := range diagramFrom.ConcernsWhoseNodeIsExpanded {
		diagramTo.ConcernsWhoseNodeIsExpanded = append(diagramTo.ConcernsWhoseNodeIsExpanded, GongCopyBranchConcern(mapOrigCopy, _concern))
	}
	for _, _concern := range diagramFrom.ConcernsWhoseInputNodeIsExpanded {
		diagramTo.ConcernsWhoseInputNodeIsExpanded = append(diagramTo.ConcernsWhoseInputNodeIsExpanded, GongCopyBranchConcern(mapOrigCopy, _concern))
	}
	for _, _concern := range diagramFrom.ConcernsWhoseStakeholderNodeIsExpanded {
		diagramTo.ConcernsWhoseStakeholderNodeIsExpanded = append(diagramTo.ConcernsWhoseStakeholderNodeIsExpanded, GongCopyBranchConcern(mapOrigCopy, _concern))
	}
	for _, _concern := range diagramFrom.ConcernssWhoseOutputNodeIsExpanded {
		diagramTo.ConcernssWhoseOutputNodeIsExpanded = append(diagramTo.ConcernssWhoseOutputNodeIsExpanded, GongCopyBranchConcern(mapOrigCopy, _concern))
	}
	for _, _concerncompositionshape := range diagramFrom.ConcernComposition_Shapes {
		diagramTo.ConcernComposition_Shapes = append(diagramTo.ConcernComposition_Shapes, GongCopyBranchConcernCompositionShape(mapOrigCopy, _concerncompositionshape))
	}
	for _, _concerninputshape := range diagramFrom.ConcernInputShapes {
		diagramTo.ConcernInputShapes = append(diagramTo.ConcernInputShapes, GongCopyBranchConcernInputShape(mapOrigCopy, _concerninputshape))
	}
	for _, _concernoutputshape := range diagramFrom.ConcernOutputShapes {
		diagramTo.ConcernOutputShapes = append(diagramTo.ConcernOutputShapes, GongCopyBranchConcernOutputShape(mapOrigCopy, _concernoutputshape))
	}
	for _, _noteshape := range diagramFrom.Note_Shapes {
		diagramTo.Note_Shapes = append(diagramTo.Note_Shapes, GongCopyBranchNoteShape(mapOrigCopy, _noteshape))
	}
	for _, _note := range diagramFrom.NotesWhoseNodeIsExpanded {
		diagramTo.NotesWhoseNodeIsExpanded = append(diagramTo.NotesWhoseNodeIsExpanded, GongCopyBranchNote(mapOrigCopy, _note))
	}
	for _, _notedeliverableshape := range diagramFrom.NoteDeliverableShapes {
		diagramTo.NoteDeliverableShapes = append(diagramTo.NoteDeliverableShapes, GongCopyBranchNoteDeliverableShape(mapOrigCopy, _notedeliverableshape))
	}
	for _, _notetaskshape := range diagramFrom.NoteTaskShapes {
		diagramTo.NoteTaskShapes = append(diagramTo.NoteTaskShapes, GongCopyBranchNoteTaskShape(mapOrigCopy, _notetaskshape))
	}
	for _, _notestakeholdershape := range diagramFrom.NoteResourceShapes {
		diagramTo.NoteResourceShapes = append(diagramTo.NoteResourceShapes, GongCopyBranchNoteStakeholderShape(mapOrigCopy, _notestakeholdershape))
	}
	for _, _stakeholdershape := range diagramFrom.Stakeholder_Shapes {
		diagramTo.Stakeholder_Shapes = append(diagramTo.Stakeholder_Shapes, GongCopyBranchStakeholderShape(mapOrigCopy, _stakeholdershape))
	}
	for _, _stakeholder := range diagramFrom.ResourcesWhoseNodeIsExpanded {
		diagramTo.ResourcesWhoseNodeIsExpanded = append(diagramTo.ResourcesWhoseNodeIsExpanded, GongCopyBranchStakeholder(mapOrigCopy, _stakeholder))
	}
	for _, _stakeholdercompositionshape := range diagramFrom.ResourceComposition_Shapes {
		diagramTo.ResourceComposition_Shapes = append(diagramTo.ResourceComposition_Shapes, GongCopyBranchStakeholderCompositionShape(mapOrigCopy, _stakeholdercompositionshape))
	}
	for _, _stakeholderconcernshape := range diagramFrom.StakeholderConcernShapes {
		diagramTo.StakeholderConcernShapes = append(diagramTo.StakeholderConcernShapes, GongCopyBranchStakeholderConcernShape(mapOrigCopy, _stakeholderconcernshape))
	}
	for _, _requirementshape := range diagramFrom.Requirement_Shapes {
		diagramTo.Requirement_Shapes = append(diagramTo.Requirement_Shapes, GongCopyBranchRequirementShape(mapOrigCopy, _requirementshape))
	}
	for _, _requirement := range diagramFrom.RequirementsWhoseNodeIsExpanded {
		diagramTo.RequirementsWhoseNodeIsExpanded = append(diagramTo.RequirementsWhoseNodeIsExpanded, GongCopyBranchRequirement(mapOrigCopy, _requirement))
	}
	for _, _conceptshape := range diagramFrom.Concept_Shapes {
		diagramTo.Concept_Shapes = append(diagramTo.Concept_Shapes, GongCopyBranchConceptShape(mapOrigCopy, _conceptshape))
	}
	for _, _concept := range diagramFrom.ConceptsWhoseNodeIsExpanded {
		diagramTo.ConceptsWhoseNodeIsExpanded = append(diagramTo.ConceptsWhoseNodeIsExpanded, GongCopyBranchConcept(mapOrigCopy, _concept))
	}
	for _, _concept := range diagramFrom.ConceptsWhoseDeliverablesNodeIsExpanded {
		diagramTo.ConceptsWhoseDeliverablesNodeIsExpanded = append(diagramTo.ConceptsWhoseDeliverablesNodeIsExpanded, GongCopyBranchConcept(mapOrigCopy, _concept))
	}
	for _, _deliverableconceptshape := range diagramFrom.DeliverableConceptShapes {
		diagramTo.DeliverableConceptShapes = append(diagramTo.DeliverableConceptShapes, GongCopyBranchDeliverableConceptShape(mapOrigCopy, _deliverableconceptshape))
	}
	for _, _diagramshape := range diagramFrom.Diagram_Shapes {
		diagramTo.Diagram_Shapes = append(diagramTo.Diagram_Shapes, GongCopyBranchDiagramShape(mapOrigCopy, _diagramshape))
	}
	for _, _diagram := range diagramFrom.DiagramsWhoseNodeIsExpanded {
		diagramTo.DiagramsWhoseNodeIsExpanded = append(diagramTo.DiagramsWhoseNodeIsExpanded, GongCopyBranchDiagram(mapOrigCopy, _diagram))
	}

	return
}

func GongCopyBranchDiagramShape(mapOrigCopy map[any]any, diagramshapeFrom *DiagramShape) (diagramshapeTo *DiagramShape) {
	var alreadyCopied bool
	diagramshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, diagramshapeFrom)
	if alreadyCopied {
		return
	}
	diagramshapeFrom.GongCopyBasicFields(diagramshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if diagramshapeFrom.Diagram != nil {
		diagramshapeTo.Diagram = GongCopyBranchDiagram(mapOrigCopy, diagramshapeFrom.Diagram)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchLibrary(mapOrigCopy map[any]any, libraryFrom *Library) (libraryTo *Library) {
	var alreadyCopied bool
	libraryTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, libraryFrom)
	if alreadyCopied {
		return
	}
	libraryFrom.GongCopyBasicFields(libraryTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _deliverable := range libraryFrom.RootDeliverables {
		libraryTo.RootDeliverables = append(libraryTo.RootDeliverables, GongCopyBranchDeliverable(mapOrigCopy, _deliverable))
	}
	for _, _concern := range libraryFrom.RootConcerns {
		libraryTo.RootConcerns = append(libraryTo.RootConcerns, GongCopyBranchConcern(mapOrigCopy, _concern))
	}
	for _, _stakeholder := range libraryFrom.RootStakeholders {
		libraryTo.RootStakeholders = append(libraryTo.RootStakeholders, GongCopyBranchStakeholder(mapOrigCopy, _stakeholder))
	}
	for _, _requirement := range libraryFrom.RootRequirements {
		libraryTo.RootRequirements = append(libraryTo.RootRequirements, GongCopyBranchRequirement(mapOrigCopy, _requirement))
	}
	for _, _concept := range libraryFrom.RootConcepts {
		libraryTo.RootConcepts = append(libraryTo.RootConcepts, GongCopyBranchConcept(mapOrigCopy, _concept))
	}
	for _, _analysisneed := range libraryFrom.AnalysisNeeds {
		libraryTo.AnalysisNeeds = append(libraryTo.AnalysisNeeds, GongCopyBranchAnalysisNeed(mapOrigCopy, _analysisneed))
	}
	for _, _note := range libraryFrom.Notes {
		libraryTo.Notes = append(libraryTo.Notes, GongCopyBranchNote(mapOrigCopy, _note))
	}
	for _, _diagram := range libraryFrom.Diagrams {
		libraryTo.Diagrams = append(libraryTo.Diagrams, GongCopyBranchDiagram(mapOrigCopy, _diagram))
	}
	for _, _library := range libraryFrom.SubLibraries {
		libraryTo.SubLibraries = append(libraryTo.SubLibraries, GongCopyBranchLibrary(mapOrigCopy, _library))
	}

	return
}

func GongCopyBranchNote(mapOrigCopy map[any]any, noteFrom *Note) (noteTo *Note) {
	var alreadyCopied bool
	noteTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, noteFrom)
	if alreadyCopied {
		return
	}
	noteFrom.GongCopyBasicFields(noteTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _deliverable := range noteFrom.Deliverables {
		noteTo.Deliverables = append(noteTo.Deliverables, GongCopyBranchDeliverable(mapOrigCopy, _deliverable))
	}
	for _, _concern := range noteFrom.Tasks {
		noteTo.Tasks = append(noteTo.Tasks, GongCopyBranchConcern(mapOrigCopy, _concern))
	}
	for _, _stakeholder := range noteFrom.Resources {
		noteTo.Resources = append(noteTo.Resources, GongCopyBranchStakeholder(mapOrigCopy, _stakeholder))
	}

	return
}

func GongCopyBranchNoteDeliverableShape(mapOrigCopy map[any]any, notedeliverableshapeFrom *NoteDeliverableShape) (notedeliverableshapeTo *NoteDeliverableShape) {
	var alreadyCopied bool
	notedeliverableshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, notedeliverableshapeFrom)
	if alreadyCopied {
		return
	}
	notedeliverableshapeFrom.GongCopyBasicFields(notedeliverableshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if notedeliverableshapeFrom.Note != nil {
		notedeliverableshapeTo.Note = GongCopyBranchNote(mapOrigCopy, notedeliverableshapeFrom.Note)
	}
	if notedeliverableshapeFrom.Deliverable != nil {
		notedeliverableshapeTo.Deliverable = GongCopyBranchDeliverable(mapOrigCopy, notedeliverableshapeFrom.Deliverable)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _controlpointshape := range notedeliverableshapeFrom.ControlPointShapes {
		notedeliverableshapeTo.ControlPointShapes = append(notedeliverableshapeTo.ControlPointShapes, GongCopyBranchControlPointShape(mapOrigCopy, _controlpointshape))
	}

	return
}

func GongCopyBranchNoteShape(mapOrigCopy map[any]any, noteshapeFrom *NoteShape) (noteshapeTo *NoteShape) {
	var alreadyCopied bool
	noteshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, noteshapeFrom)
	if alreadyCopied {
		return
	}
	noteshapeFrom.GongCopyBasicFields(noteshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if noteshapeFrom.Note != nil {
		noteshapeTo.Note = GongCopyBranchNote(mapOrigCopy, noteshapeFrom.Note)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchNoteStakeholderShape(mapOrigCopy map[any]any, notestakeholdershapeFrom *NoteStakeholderShape) (notestakeholdershapeTo *NoteStakeholderShape) {
	var alreadyCopied bool
	notestakeholdershapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, notestakeholdershapeFrom)
	if alreadyCopied {
		return
	}
	notestakeholdershapeFrom.GongCopyBasicFields(notestakeholdershapeTo)

	//insertion point for the staging of instances referenced by pointers
	if notestakeholdershapeFrom.Note != nil {
		notestakeholdershapeTo.Note = GongCopyBranchNote(mapOrigCopy, notestakeholdershapeFrom.Note)
	}
	if notestakeholdershapeFrom.Stakeholder != nil {
		notestakeholdershapeTo.Stakeholder = GongCopyBranchStakeholder(mapOrigCopy, notestakeholdershapeFrom.Stakeholder)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _controlpointshape := range notestakeholdershapeFrom.ControlPointShapes {
		notestakeholdershapeTo.ControlPointShapes = append(notestakeholdershapeTo.ControlPointShapes, GongCopyBranchControlPointShape(mapOrigCopy, _controlpointshape))
	}

	return
}

func GongCopyBranchNoteTaskShape(mapOrigCopy map[any]any, notetaskshapeFrom *NoteTaskShape) (notetaskshapeTo *NoteTaskShape) {
	var alreadyCopied bool
	notetaskshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, notetaskshapeFrom)
	if alreadyCopied {
		return
	}
	notetaskshapeFrom.GongCopyBasicFields(notetaskshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if notetaskshapeFrom.Note != nil {
		notetaskshapeTo.Note = GongCopyBranchNote(mapOrigCopy, notetaskshapeFrom.Note)
	}
	if notetaskshapeFrom.Task != nil {
		notetaskshapeTo.Task = GongCopyBranchConcern(mapOrigCopy, notetaskshapeFrom.Task)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _controlpointshape := range notetaskshapeFrom.ControlPointShapes {
		notetaskshapeTo.ControlPointShapes = append(notetaskshapeTo.ControlPointShapes, GongCopyBranchControlPointShape(mapOrigCopy, _controlpointshape))
	}

	return
}

func GongCopyBranchRequirement(mapOrigCopy map[any]any, requirementFrom *Requirement) (requirementTo *Requirement) {
	var alreadyCopied bool
	requirementTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, requirementFrom)
	if alreadyCopied {
		return
	}
	requirementFrom.GongCopyBasicFields(requirementTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _supportlevel := range requirementFrom.SupportLevels {
		requirementTo.SupportLevels = append(requirementTo.SupportLevels, GongCopyBranchSupportLevel(mapOrigCopy, _supportlevel))
	}
	for _, _concept := range requirementFrom.Concepts {
		requirementTo.Concepts = append(requirementTo.Concepts, GongCopyBranchConcept(mapOrigCopy, _concept))
	}

	return
}

func GongCopyBranchRequirementShape(mapOrigCopy map[any]any, requirementshapeFrom *RequirementShape) (requirementshapeTo *RequirementShape) {
	var alreadyCopied bool
	requirementshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, requirementshapeFrom)
	if alreadyCopied {
		return
	}
	requirementshapeFrom.GongCopyBasicFields(requirementshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if requirementshapeFrom.Requirement != nil {
		requirementshapeTo.Requirement = GongCopyBranchRequirement(mapOrigCopy, requirementshapeFrom.Requirement)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchStakeholder(mapOrigCopy map[any]any, stakeholderFrom *Stakeholder) (stakeholderTo *Stakeholder) {
	var alreadyCopied bool
	stakeholderTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, stakeholderFrom)
	if alreadyCopied {
		return
	}
	stakeholderFrom.GongCopyBasicFields(stakeholderTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _concern := range stakeholderFrom.Concerns {
		stakeholderTo.Concerns = append(stakeholderTo.Concerns, GongCopyBranchConcern(mapOrigCopy, _concern))
	}
	for _, _stakeholder := range stakeholderFrom.SubStakeholders {
		stakeholderTo.SubStakeholders = append(stakeholderTo.SubStakeholders, GongCopyBranchStakeholder(mapOrigCopy, _stakeholder))
	}

	return
}

func GongCopyBranchStakeholderCompositionShape(mapOrigCopy map[any]any, stakeholdercompositionshapeFrom *StakeholderCompositionShape) (stakeholdercompositionshapeTo *StakeholderCompositionShape) {
	var alreadyCopied bool
	stakeholdercompositionshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, stakeholdercompositionshapeFrom)
	if alreadyCopied {
		return
	}
	stakeholdercompositionshapeFrom.GongCopyBasicFields(stakeholdercompositionshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if stakeholdercompositionshapeFrom.Stakeholder != nil {
		stakeholdercompositionshapeTo.Stakeholder = GongCopyBranchStakeholder(mapOrigCopy, stakeholdercompositionshapeFrom.Stakeholder)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _controlpointshape := range stakeholdercompositionshapeFrom.ControlPointShapes {
		stakeholdercompositionshapeTo.ControlPointShapes = append(stakeholdercompositionshapeTo.ControlPointShapes, GongCopyBranchControlPointShape(mapOrigCopy, _controlpointshape))
	}

	return
}

func GongCopyBranchStakeholderConcernShape(mapOrigCopy map[any]any, stakeholderconcernshapeFrom *StakeholderConcernShape) (stakeholderconcernshapeTo *StakeholderConcernShape) {
	var alreadyCopied bool
	stakeholderconcernshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, stakeholderconcernshapeFrom)
	if alreadyCopied {
		return
	}
	stakeholderconcernshapeFrom.GongCopyBasicFields(stakeholderconcernshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if stakeholderconcernshapeFrom.Stakeholder != nil {
		stakeholderconcernshapeTo.Stakeholder = GongCopyBranchStakeholder(mapOrigCopy, stakeholderconcernshapeFrom.Stakeholder)
	}
	if stakeholderconcernshapeFrom.Concern != nil {
		stakeholderconcernshapeTo.Concern = GongCopyBranchConcern(mapOrigCopy, stakeholderconcernshapeFrom.Concern)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _controlpointshape := range stakeholderconcernshapeFrom.ControlPointShapes {
		stakeholderconcernshapeTo.ControlPointShapes = append(stakeholderconcernshapeTo.ControlPointShapes, GongCopyBranchControlPointShape(mapOrigCopy, _controlpointshape))
	}

	return
}

func GongCopyBranchStakeholderShape(mapOrigCopy map[any]any, stakeholdershapeFrom *StakeholderShape) (stakeholdershapeTo *StakeholderShape) {
	var alreadyCopied bool
	stakeholdershapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, stakeholdershapeFrom)
	if alreadyCopied {
		return
	}
	stakeholdershapeFrom.GongCopyBasicFields(stakeholdershapeTo)

	//insertion point for the staging of instances referenced by pointers
	if stakeholdershapeFrom.Stakeholder != nil {
		stakeholdershapeTo.Stakeholder = GongCopyBranchStakeholder(mapOrigCopy, stakeholdershapeFrom.Stakeholder)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchSupportLevel(mapOrigCopy map[any]any, supportlevelFrom *SupportLevel) (supportlevelTo *SupportLevel) {
	var alreadyCopied bool
	supportlevelTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, supportlevelFrom)
	if alreadyCopied {
		return
	}
	supportlevelFrom.GongCopyBasicFields(supportlevelTo)

	//insertion point for the staging of instances referenced by pointers
	if supportlevelFrom.Tool != nil {
		supportlevelTo.Tool = GongCopyBranchTool(mapOrigCopy, supportlevelFrom.Tool)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTool(mapOrigCopy map[any]any, toolFrom *Tool) (toolTo *Tool) {
	var alreadyCopied bool
	toolTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, toolFrom)
	if alreadyCopied {
		return
	}
	toolFrom.GongCopyBasicFields(toolTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

// UnstageBranch stages instance and apply UnstageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the insance
//
// the algorithm stops along the course of graph if a vertex is already staged
// UnstageBranch is the Stage method that unstages instance and applies UnstageBranch recursively.
func (stage *Stage) UnstageBranch(instance GongstructIF) {
	if instance != nil {
		instance.GongUnstageBranch(stage)
	}
}

// insertion point for unstage branch per struct
func (analysisneed *AnalysisNeed) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(analysisneed) {
		return
	}

	analysisneed.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (concept *Concept) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(concept) {
		return
	}

	concept.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _tool := range concept.Tools {
		stage.UnstageBranch(_tool)
	}

}

func (conceptshape *ConceptShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(conceptshape) {
		return
	}

	conceptshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if conceptshape.Concept != nil {
		stage.UnstageBranch(conceptshape.Concept)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (concern *Concern) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(concern) {
		return
	}

	concern.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _concern := range concern.SubConcerns {
		stage.UnstageBranch(_concern)
	}
	for _, _deliverable := range concern.Inputs {
		stage.UnstageBranch(_deliverable)
	}
	for _, _deliverable := range concern.Outputs {
		stage.UnstageBranch(_deliverable)
	}
	for _, _requirement := range concern.Requirements {
		stage.UnstageBranch(_requirement)
	}

}

func (concerncompositionshape *ConcernCompositionShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(concerncompositionshape) {
		return
	}

	concerncompositionshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if concerncompositionshape.Concern != nil {
		stage.UnstageBranch(concerncompositionshape.Concern)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _controlpointshape := range concerncompositionshape.ControlPointShapes {
		stage.UnstageBranch(_controlpointshape)
	}

}

func (concerninputshape *ConcernInputShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(concerninputshape) {
		return
	}

	concerninputshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if concerninputshape.Deliverable != nil {
		stage.UnstageBranch(concerninputshape.Deliverable)
	}
	if concerninputshape.Concern != nil {
		stage.UnstageBranch(concerninputshape.Concern)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _controlpointshape := range concerninputshape.ControlPointShapes {
		stage.UnstageBranch(_controlpointshape)
	}

}

func (concernoutputshape *ConcernOutputShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(concernoutputshape) {
		return
	}

	concernoutputshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if concernoutputshape.Concern != nil {
		stage.UnstageBranch(concernoutputshape.Concern)
	}
	if concernoutputshape.Deliverable != nil {
		stage.UnstageBranch(concernoutputshape.Deliverable)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _controlpointshape := range concernoutputshape.ControlPointShapes {
		stage.UnstageBranch(_controlpointshape)
	}

}

func (concernshape *ConcernShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(concernshape) {
		return
	}

	concernshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if concernshape.Concern != nil {
		stage.UnstageBranch(concernshape.Concern)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (controlpointshape *ControlPointShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(controlpointshape) {
		return
	}

	controlpointshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (deliverable *Deliverable) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(deliverable) {
		return
	}

	deliverable.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _deliverable := range deliverable.SubDeliverables {
		stage.UnstageBranch(_deliverable)
	}
	for _, _concept := range deliverable.Concepts {
		stage.UnstageBranch(_concept)
	}

}

func (deliverablecompositionshape *DeliverableCompositionShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(deliverablecompositionshape) {
		return
	}

	deliverablecompositionshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if deliverablecompositionshape.Deliverable != nil {
		stage.UnstageBranch(deliverablecompositionshape.Deliverable)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _controlpointshape := range deliverablecompositionshape.ControlPointShapes {
		stage.UnstageBranch(_controlpointshape)
	}

}

func (deliverableconceptshape *DeliverableConceptShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(deliverableconceptshape) {
		return
	}

	deliverableconceptshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if deliverableconceptshape.Deliverable != nil {
		stage.UnstageBranch(deliverableconceptshape.Deliverable)
	}
	if deliverableconceptshape.Concept != nil {
		stage.UnstageBranch(deliverableconceptshape.Concept)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _controlpointshape := range deliverableconceptshape.ControlPointShapes {
		stage.UnstageBranch(_controlpointshape)
	}

}

func (deliverableshape *DeliverableShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(deliverableshape) {
		return
	}

	deliverableshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if deliverableshape.Deliverable != nil {
		stage.UnstageBranch(deliverableshape.Deliverable)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (diagram *Diagram) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(diagram) {
		return
	}

	diagram.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _concern := range diagram.ConcernsWhoseRequirementsNodeIsExpanded {
		stage.UnstageBranch(_concern)
	}
	for _, _deliverableshape := range diagram.Deliverable_Shapes {
		stage.UnstageBranch(_deliverableshape)
	}
	for _, _deliverable := range diagram.DeliverablesWhoseNodeIsExpanded {
		stage.UnstageBranch(_deliverable)
	}
	for _, _deliverable := range diagram.DeliverablesWhoseConceptsNodeIsExpanded {
		stage.UnstageBranch(_deliverable)
	}
	for _, _deliverablecompositionshape := range diagram.DeliverableComposition_Shapes {
		stage.UnstageBranch(_deliverablecompositionshape)
	}
	for _, _concernshape := range diagram.Concern_Shapes {
		stage.UnstageBranch(_concernshape)
	}
	for _, _concern := range diagram.ConcernsWhoseNodeIsExpanded {
		stage.UnstageBranch(_concern)
	}
	for _, _concern := range diagram.ConcernsWhoseInputNodeIsExpanded {
		stage.UnstageBranch(_concern)
	}
	for _, _concern := range diagram.ConcernsWhoseStakeholderNodeIsExpanded {
		stage.UnstageBranch(_concern)
	}
	for _, _concern := range diagram.ConcernssWhoseOutputNodeIsExpanded {
		stage.UnstageBranch(_concern)
	}
	for _, _concerncompositionshape := range diagram.ConcernComposition_Shapes {
		stage.UnstageBranch(_concerncompositionshape)
	}
	for _, _concerninputshape := range diagram.ConcernInputShapes {
		stage.UnstageBranch(_concerninputshape)
	}
	for _, _concernoutputshape := range diagram.ConcernOutputShapes {
		stage.UnstageBranch(_concernoutputshape)
	}
	for _, _noteshape := range diagram.Note_Shapes {
		stage.UnstageBranch(_noteshape)
	}
	for _, _note := range diagram.NotesWhoseNodeIsExpanded {
		stage.UnstageBranch(_note)
	}
	for _, _notedeliverableshape := range diagram.NoteDeliverableShapes {
		stage.UnstageBranch(_notedeliverableshape)
	}
	for _, _notetaskshape := range diagram.NoteTaskShapes {
		stage.UnstageBranch(_notetaskshape)
	}
	for _, _notestakeholdershape := range diagram.NoteResourceShapes {
		stage.UnstageBranch(_notestakeholdershape)
	}
	for _, _stakeholdershape := range diagram.Stakeholder_Shapes {
		stage.UnstageBranch(_stakeholdershape)
	}
	for _, _stakeholder := range diagram.ResourcesWhoseNodeIsExpanded {
		stage.UnstageBranch(_stakeholder)
	}
	for _, _stakeholdercompositionshape := range diagram.ResourceComposition_Shapes {
		stage.UnstageBranch(_stakeholdercompositionshape)
	}
	for _, _stakeholderconcernshape := range diagram.StakeholderConcernShapes {
		stage.UnstageBranch(_stakeholderconcernshape)
	}
	for _, _requirementshape := range diagram.Requirement_Shapes {
		stage.UnstageBranch(_requirementshape)
	}
	for _, _requirement := range diagram.RequirementsWhoseNodeIsExpanded {
		stage.UnstageBranch(_requirement)
	}
	for _, _conceptshape := range diagram.Concept_Shapes {
		stage.UnstageBranch(_conceptshape)
	}
	for _, _concept := range diagram.ConceptsWhoseNodeIsExpanded {
		stage.UnstageBranch(_concept)
	}
	for _, _concept := range diagram.ConceptsWhoseDeliverablesNodeIsExpanded {
		stage.UnstageBranch(_concept)
	}
	for _, _deliverableconceptshape := range diagram.DeliverableConceptShapes {
		stage.UnstageBranch(_deliverableconceptshape)
	}
	for _, _diagramshape := range diagram.Diagram_Shapes {
		stage.UnstageBranch(_diagramshape)
	}
	for _, _diagram := range diagram.DiagramsWhoseNodeIsExpanded {
		stage.UnstageBranch(_diagram)
	}

}

func (diagramshape *DiagramShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(diagramshape) {
		return
	}

	diagramshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if diagramshape.Diagram != nil {
		stage.UnstageBranch(diagramshape.Diagram)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (library *Library) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(library) {
		return
	}

	library.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _deliverable := range library.RootDeliverables {
		stage.UnstageBranch(_deliverable)
	}
	for _, _concern := range library.RootConcerns {
		stage.UnstageBranch(_concern)
	}
	for _, _stakeholder := range library.RootStakeholders {
		stage.UnstageBranch(_stakeholder)
	}
	for _, _requirement := range library.RootRequirements {
		stage.UnstageBranch(_requirement)
	}
	for _, _concept := range library.RootConcepts {
		stage.UnstageBranch(_concept)
	}
	for _, _analysisneed := range library.AnalysisNeeds {
		stage.UnstageBranch(_analysisneed)
	}
	for _, _note := range library.Notes {
		stage.UnstageBranch(_note)
	}
	for _, _diagram := range library.Diagrams {
		stage.UnstageBranch(_diagram)
	}
	for _, _library := range library.SubLibraries {
		stage.UnstageBranch(_library)
	}

}

func (note *Note) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(note) {
		return
	}

	note.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _deliverable := range note.Deliverables {
		stage.UnstageBranch(_deliverable)
	}
	for _, _concern := range note.Tasks {
		stage.UnstageBranch(_concern)
	}
	for _, _stakeholder := range note.Resources {
		stage.UnstageBranch(_stakeholder)
	}

}

func (notedeliverableshape *NoteDeliverableShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(notedeliverableshape) {
		return
	}

	notedeliverableshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if notedeliverableshape.Note != nil {
		stage.UnstageBranch(notedeliverableshape.Note)
	}
	if notedeliverableshape.Deliverable != nil {
		stage.UnstageBranch(notedeliverableshape.Deliverable)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _controlpointshape := range notedeliverableshape.ControlPointShapes {
		stage.UnstageBranch(_controlpointshape)
	}

}

func (noteshape *NoteShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(noteshape) {
		return
	}

	noteshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if noteshape.Note != nil {
		stage.UnstageBranch(noteshape.Note)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (notestakeholdershape *NoteStakeholderShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(notestakeholdershape) {
		return
	}

	notestakeholdershape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if notestakeholdershape.Note != nil {
		stage.UnstageBranch(notestakeholdershape.Note)
	}
	if notestakeholdershape.Stakeholder != nil {
		stage.UnstageBranch(notestakeholdershape.Stakeholder)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _controlpointshape := range notestakeholdershape.ControlPointShapes {
		stage.UnstageBranch(_controlpointshape)
	}

}

func (notetaskshape *NoteTaskShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(notetaskshape) {
		return
	}

	notetaskshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if notetaskshape.Note != nil {
		stage.UnstageBranch(notetaskshape.Note)
	}
	if notetaskshape.Task != nil {
		stage.UnstageBranch(notetaskshape.Task)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _controlpointshape := range notetaskshape.ControlPointShapes {
		stage.UnstageBranch(_controlpointshape)
	}

}

func (requirement *Requirement) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(requirement) {
		return
	}

	requirement.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _supportlevel := range requirement.SupportLevels {
		stage.UnstageBranch(_supportlevel)
	}
	for _, _concept := range requirement.Concepts {
		stage.UnstageBranch(_concept)
	}

}

func (requirementshape *RequirementShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(requirementshape) {
		return
	}

	requirementshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if requirementshape.Requirement != nil {
		stage.UnstageBranch(requirementshape.Requirement)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stakeholder *Stakeholder) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(stakeholder) {
		return
	}

	stakeholder.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _concern := range stakeholder.Concerns {
		stage.UnstageBranch(_concern)
	}
	for _, _stakeholder := range stakeholder.SubStakeholders {
		stage.UnstageBranch(_stakeholder)
	}

}

func (stakeholdercompositionshape *StakeholderCompositionShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(stakeholdercompositionshape) {
		return
	}

	stakeholdercompositionshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if stakeholdercompositionshape.Stakeholder != nil {
		stage.UnstageBranch(stakeholdercompositionshape.Stakeholder)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _controlpointshape := range stakeholdercompositionshape.ControlPointShapes {
		stage.UnstageBranch(_controlpointshape)
	}

}

func (stakeholderconcernshape *StakeholderConcernShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(stakeholderconcernshape) {
		return
	}

	stakeholderconcernshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if stakeholderconcernshape.Stakeholder != nil {
		stage.UnstageBranch(stakeholderconcernshape.Stakeholder)
	}
	if stakeholderconcernshape.Concern != nil {
		stage.UnstageBranch(stakeholderconcernshape.Concern)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _controlpointshape := range stakeholderconcernshape.ControlPointShapes {
		stage.UnstageBranch(_controlpointshape)
	}

}

func (stakeholdershape *StakeholderShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(stakeholdershape) {
		return
	}

	stakeholdershape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if stakeholdershape.Stakeholder != nil {
		stage.UnstageBranch(stakeholdershape.Stakeholder)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (supportlevel *SupportLevel) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(supportlevel) {
		return
	}

	supportlevel.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if supportlevel.Tool != nil {
		stage.UnstageBranch(supportlevel.Tool)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (tool *Tool) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(tool) {
		return
	}

	tool.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

// insertion point for pointer reconstruction from references
func (reference *AnalysisNeed) GongReconstructPointersFromReferences(stage *Stage, instance *AnalysisNeed) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Concept) GongReconstructPointersFromReferences(stage *Stage, instance *Concept) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Tools, stage.Tools_reference, instance.Tools)
}

func (reference *ConceptShape) GongReconstructPointersFromReferences(stage *Stage, instance *ConceptShape) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Concept, stage.Concepts_reference, instance.Concept)
	// insertion point for slice of pointers field
}

func (reference *Concern) GongReconstructPointersFromReferences(stage *Stage, instance *Concern) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.SubConcerns, stage.Concerns_reference, instance.SubConcerns)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Inputs, stage.Deliverables_reference, instance.Inputs)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Outputs, stage.Deliverables_reference, instance.Outputs)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Requirements, stage.Requirements_reference, instance.Requirements)
}

func (reference *ConcernCompositionShape) GongReconstructPointersFromReferences(stage *Stage, instance *ConcernCompositionShape) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Concern, stage.Concerns_reference, instance.Concern)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.ControlPointShapes, stage.ControlPointShapes_reference, instance.ControlPointShapes)
}

func (reference *ConcernInputShape) GongReconstructPointersFromReferences(stage *Stage, instance *ConcernInputShape) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Deliverable, stage.Deliverables_reference, instance.Deliverable)
	__gong__reconstructPointer(&reference.Concern, stage.Concerns_reference, instance.Concern)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.ControlPointShapes, stage.ControlPointShapes_reference, instance.ControlPointShapes)
}

func (reference *ConcernOutputShape) GongReconstructPointersFromReferences(stage *Stage, instance *ConcernOutputShape) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Concern, stage.Concerns_reference, instance.Concern)
	__gong__reconstructPointer(&reference.Deliverable, stage.Deliverables_reference, instance.Deliverable)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.ControlPointShapes, stage.ControlPointShapes_reference, instance.ControlPointShapes)
}

func (reference *ConcernShape) GongReconstructPointersFromReferences(stage *Stage, instance *ConcernShape) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Concern, stage.Concerns_reference, instance.Concern)
	// insertion point for slice of pointers field
}

func (reference *ControlPointShape) GongReconstructPointersFromReferences(stage *Stage, instance *ControlPointShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Deliverable) GongReconstructPointersFromReferences(stage *Stage, instance *Deliverable) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.SubDeliverables, stage.Deliverables_reference, instance.SubDeliverables)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Concepts, stage.Concepts_reference, instance.Concepts)
}

func (reference *DeliverableCompositionShape) GongReconstructPointersFromReferences(stage *Stage, instance *DeliverableCompositionShape) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Deliverable, stage.Deliverables_reference, instance.Deliverable)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.ControlPointShapes, stage.ControlPointShapes_reference, instance.ControlPointShapes)
}

func (reference *DeliverableConceptShape) GongReconstructPointersFromReferences(stage *Stage, instance *DeliverableConceptShape) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Deliverable, stage.Deliverables_reference, instance.Deliverable)
	__gong__reconstructPointer(&reference.Concept, stage.Concepts_reference, instance.Concept)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.ControlPointShapes, stage.ControlPointShapes_reference, instance.ControlPointShapes)
}

func (reference *DeliverableShape) GongReconstructPointersFromReferences(stage *Stage, instance *DeliverableShape) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Deliverable, stage.Deliverables_reference, instance.Deliverable)
	// insertion point for slice of pointers field
}

func (reference *Diagram) GongReconstructPointersFromReferences(stage *Stage, instance *Diagram) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.ConcernsWhoseRequirementsNodeIsExpanded, stage.Concerns_reference, instance.ConcernsWhoseRequirementsNodeIsExpanded)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Deliverable_Shapes, stage.DeliverableShapes_reference, instance.Deliverable_Shapes)
	__gong__reconstructSliceOfPointersFromReferences(&reference.DeliverablesWhoseNodeIsExpanded, stage.Deliverables_reference, instance.DeliverablesWhoseNodeIsExpanded)
	__gong__reconstructSliceOfPointersFromReferences(&reference.DeliverablesWhoseConceptsNodeIsExpanded, stage.Deliverables_reference, instance.DeliverablesWhoseConceptsNodeIsExpanded)
	__gong__reconstructSliceOfPointersFromReferences(&reference.DeliverableComposition_Shapes, stage.DeliverableCompositionShapes_reference, instance.DeliverableComposition_Shapes)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Concern_Shapes, stage.ConcernShapes_reference, instance.Concern_Shapes)
	__gong__reconstructSliceOfPointersFromReferences(&reference.ConcernsWhoseNodeIsExpanded, stage.Concerns_reference, instance.ConcernsWhoseNodeIsExpanded)
	__gong__reconstructSliceOfPointersFromReferences(&reference.ConcernsWhoseInputNodeIsExpanded, stage.Concerns_reference, instance.ConcernsWhoseInputNodeIsExpanded)
	__gong__reconstructSliceOfPointersFromReferences(&reference.ConcernsWhoseStakeholderNodeIsExpanded, stage.Concerns_reference, instance.ConcernsWhoseStakeholderNodeIsExpanded)
	__gong__reconstructSliceOfPointersFromReferences(&reference.ConcernssWhoseOutputNodeIsExpanded, stage.Concerns_reference, instance.ConcernssWhoseOutputNodeIsExpanded)
	__gong__reconstructSliceOfPointersFromReferences(&reference.ConcernComposition_Shapes, stage.ConcernCompositionShapes_reference, instance.ConcernComposition_Shapes)
	__gong__reconstructSliceOfPointersFromReferences(&reference.ConcernInputShapes, stage.ConcernInputShapes_reference, instance.ConcernInputShapes)
	__gong__reconstructSliceOfPointersFromReferences(&reference.ConcernOutputShapes, stage.ConcernOutputShapes_reference, instance.ConcernOutputShapes)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Note_Shapes, stage.NoteShapes_reference, instance.Note_Shapes)
	__gong__reconstructSliceOfPointersFromReferences(&reference.NotesWhoseNodeIsExpanded, stage.Notes_reference, instance.NotesWhoseNodeIsExpanded)
	__gong__reconstructSliceOfPointersFromReferences(&reference.NoteDeliverableShapes, stage.NoteDeliverableShapes_reference, instance.NoteDeliverableShapes)
	__gong__reconstructSliceOfPointersFromReferences(&reference.NoteTaskShapes, stage.NoteTaskShapes_reference, instance.NoteTaskShapes)
	__gong__reconstructSliceOfPointersFromReferences(&reference.NoteResourceShapes, stage.NoteStakeholderShapes_reference, instance.NoteResourceShapes)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Stakeholder_Shapes, stage.StakeholderShapes_reference, instance.Stakeholder_Shapes)
	__gong__reconstructSliceOfPointersFromReferences(&reference.ResourcesWhoseNodeIsExpanded, stage.Stakeholders_reference, instance.ResourcesWhoseNodeIsExpanded)
	__gong__reconstructSliceOfPointersFromReferences(&reference.ResourceComposition_Shapes, stage.StakeholderCompositionShapes_reference, instance.ResourceComposition_Shapes)
	__gong__reconstructSliceOfPointersFromReferences(&reference.StakeholderConcernShapes, stage.StakeholderConcernShapes_reference, instance.StakeholderConcernShapes)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Requirement_Shapes, stage.RequirementShapes_reference, instance.Requirement_Shapes)
	__gong__reconstructSliceOfPointersFromReferences(&reference.RequirementsWhoseNodeIsExpanded, stage.Requirements_reference, instance.RequirementsWhoseNodeIsExpanded)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Concept_Shapes, stage.ConceptShapes_reference, instance.Concept_Shapes)
	__gong__reconstructSliceOfPointersFromReferences(&reference.ConceptsWhoseNodeIsExpanded, stage.Concepts_reference, instance.ConceptsWhoseNodeIsExpanded)
	__gong__reconstructSliceOfPointersFromReferences(&reference.ConceptsWhoseDeliverablesNodeIsExpanded, stage.Concepts_reference, instance.ConceptsWhoseDeliverablesNodeIsExpanded)
	__gong__reconstructSliceOfPointersFromReferences(&reference.DeliverableConceptShapes, stage.DeliverableConceptShapes_reference, instance.DeliverableConceptShapes)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Diagram_Shapes, stage.DiagramShapes_reference, instance.Diagram_Shapes)
	__gong__reconstructSliceOfPointersFromReferences(&reference.DiagramsWhoseNodeIsExpanded, stage.Diagrams_reference, instance.DiagramsWhoseNodeIsExpanded)
}

func (reference *DiagramShape) GongReconstructPointersFromReferences(stage *Stage, instance *DiagramShape) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Diagram, stage.Diagrams_reference, instance.Diagram)
	// insertion point for slice of pointers field
}

func (reference *Library) GongReconstructPointersFromReferences(stage *Stage, instance *Library) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.RootDeliverables, stage.Deliverables_reference, instance.RootDeliverables)
	__gong__reconstructSliceOfPointersFromReferences(&reference.RootConcerns, stage.Concerns_reference, instance.RootConcerns)
	__gong__reconstructSliceOfPointersFromReferences(&reference.RootStakeholders, stage.Stakeholders_reference, instance.RootStakeholders)
	__gong__reconstructSliceOfPointersFromReferences(&reference.RootRequirements, stage.Requirements_reference, instance.RootRequirements)
	__gong__reconstructSliceOfPointersFromReferences(&reference.RootConcepts, stage.Concepts_reference, instance.RootConcepts)
	__gong__reconstructSliceOfPointersFromReferences(&reference.AnalysisNeeds, stage.AnalysisNeeds_reference, instance.AnalysisNeeds)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Notes, stage.Notes_reference, instance.Notes)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Diagrams, stage.Diagrams_reference, instance.Diagrams)
	__gong__reconstructSliceOfPointersFromReferences(&reference.SubLibraries, stage.Librarys_reference, instance.SubLibraries)
}

func (reference *Note) GongReconstructPointersFromReferences(stage *Stage, instance *Note) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Deliverables, stage.Deliverables_reference, instance.Deliverables)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Tasks, stage.Concerns_reference, instance.Tasks)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Resources, stage.Stakeholders_reference, instance.Resources)
}

func (reference *NoteDeliverableShape) GongReconstructPointersFromReferences(stage *Stage, instance *NoteDeliverableShape) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Note, stage.Notes_reference, instance.Note)
	__gong__reconstructPointer(&reference.Deliverable, stage.Deliverables_reference, instance.Deliverable)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.ControlPointShapes, stage.ControlPointShapes_reference, instance.ControlPointShapes)
}

func (reference *NoteShape) GongReconstructPointersFromReferences(stage *Stage, instance *NoteShape) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Note, stage.Notes_reference, instance.Note)
	// insertion point for slice of pointers field
}

func (reference *NoteStakeholderShape) GongReconstructPointersFromReferences(stage *Stage, instance *NoteStakeholderShape) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Note, stage.Notes_reference, instance.Note)
	__gong__reconstructPointer(&reference.Stakeholder, stage.Stakeholders_reference, instance.Stakeholder)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.ControlPointShapes, stage.ControlPointShapes_reference, instance.ControlPointShapes)
}

func (reference *NoteTaskShape) GongReconstructPointersFromReferences(stage *Stage, instance *NoteTaskShape) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Note, stage.Notes_reference, instance.Note)
	__gong__reconstructPointer(&reference.Task, stage.Concerns_reference, instance.Task)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.ControlPointShapes, stage.ControlPointShapes_reference, instance.ControlPointShapes)
}

func (reference *Requirement) GongReconstructPointersFromReferences(stage *Stage, instance *Requirement) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.SupportLevels, stage.SupportLevels_reference, instance.SupportLevels)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Concepts, stage.Concepts_reference, instance.Concepts)
}

func (reference *RequirementShape) GongReconstructPointersFromReferences(stage *Stage, instance *RequirementShape) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Requirement, stage.Requirements_reference, instance.Requirement)
	// insertion point for slice of pointers field
}

func (reference *Stakeholder) GongReconstructPointersFromReferences(stage *Stage, instance *Stakeholder) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Concerns, stage.Concerns_reference, instance.Concerns)
	__gong__reconstructSliceOfPointersFromReferences(&reference.SubStakeholders, stage.Stakeholders_reference, instance.SubStakeholders)
}

func (reference *StakeholderCompositionShape) GongReconstructPointersFromReferences(stage *Stage, instance *StakeholderCompositionShape) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Stakeholder, stage.Stakeholders_reference, instance.Stakeholder)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.ControlPointShapes, stage.ControlPointShapes_reference, instance.ControlPointShapes)
}

func (reference *StakeholderConcernShape) GongReconstructPointersFromReferences(stage *Stage, instance *StakeholderConcernShape) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Stakeholder, stage.Stakeholders_reference, instance.Stakeholder)
	__gong__reconstructPointer(&reference.Concern, stage.Concerns_reference, instance.Concern)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.ControlPointShapes, stage.ControlPointShapes_reference, instance.ControlPointShapes)
}

func (reference *StakeholderShape) GongReconstructPointersFromReferences(stage *Stage, instance *StakeholderShape) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Stakeholder, stage.Stakeholders_reference, instance.Stakeholder)
	// insertion point for slice of pointers field
}

func (reference *SupportLevel) GongReconstructPointersFromReferences(stage *Stage, instance *SupportLevel) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Tool, stage.Tools_reference, instance.Tool)
	// insertion point for slice of pointers field
}

func (reference *Tool) GongReconstructPointersFromReferences(stage *Stage, instance *Tool) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

// insertion point for pointer reconstruction from instances
func (reference *AnalysisNeed) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Concept) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Tools, stage.Tools_instance)
}

func (reference *ConceptShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Concept, stage.Concepts_instance)
	// insertion point for slice of pointers fields
}

func (reference *Concern) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.SubConcerns, stage.Concerns_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Inputs, stage.Deliverables_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Outputs, stage.Deliverables_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Requirements, stage.Requirements_instance)
}

func (reference *ConcernCompositionShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Concern, stage.Concerns_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.ControlPointShapes, stage.ControlPointShapes_instance)
}

func (reference *ConcernInputShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Deliverable, stage.Deliverables_instance)
	__gong__reconstructPointerFromInstance(&reference.Concern, stage.Concerns_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.ControlPointShapes, stage.ControlPointShapes_instance)
}

func (reference *ConcernOutputShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Concern, stage.Concerns_instance)
	__gong__reconstructPointerFromInstance(&reference.Deliverable, stage.Deliverables_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.ControlPointShapes, stage.ControlPointShapes_instance)
}

func (reference *ConcernShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Concern, stage.Concerns_instance)
	// insertion point for slice of pointers fields
}

func (reference *ControlPointShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Deliverable) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.SubDeliverables, stage.Deliverables_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Concepts, stage.Concepts_instance)
}

func (reference *DeliverableCompositionShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Deliverable, stage.Deliverables_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.ControlPointShapes, stage.ControlPointShapes_instance)
}

func (reference *DeliverableConceptShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Deliverable, stage.Deliverables_instance)
	__gong__reconstructPointerFromInstance(&reference.Concept, stage.Concepts_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.ControlPointShapes, stage.ControlPointShapes_instance)
}

func (reference *DeliverableShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Deliverable, stage.Deliverables_instance)
	// insertion point for slice of pointers fields
}

func (reference *Diagram) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.ConcernsWhoseRequirementsNodeIsExpanded, stage.Concerns_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Deliverable_Shapes, stage.DeliverableShapes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.DeliverablesWhoseNodeIsExpanded, stage.Deliverables_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.DeliverablesWhoseConceptsNodeIsExpanded, stage.Deliverables_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.DeliverableComposition_Shapes, stage.DeliverableCompositionShapes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Concern_Shapes, stage.ConcernShapes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.ConcernsWhoseNodeIsExpanded, stage.Concerns_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.ConcernsWhoseInputNodeIsExpanded, stage.Concerns_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.ConcernsWhoseStakeholderNodeIsExpanded, stage.Concerns_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.ConcernssWhoseOutputNodeIsExpanded, stage.Concerns_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.ConcernComposition_Shapes, stage.ConcernCompositionShapes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.ConcernInputShapes, stage.ConcernInputShapes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.ConcernOutputShapes, stage.ConcernOutputShapes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Note_Shapes, stage.NoteShapes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.NotesWhoseNodeIsExpanded, stage.Notes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.NoteDeliverableShapes, stage.NoteDeliverableShapes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.NoteTaskShapes, stage.NoteTaskShapes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.NoteResourceShapes, stage.NoteStakeholderShapes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Stakeholder_Shapes, stage.StakeholderShapes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.ResourcesWhoseNodeIsExpanded, stage.Stakeholders_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.ResourceComposition_Shapes, stage.StakeholderCompositionShapes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.StakeholderConcernShapes, stage.StakeholderConcernShapes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Requirement_Shapes, stage.RequirementShapes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.RequirementsWhoseNodeIsExpanded, stage.Requirements_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Concept_Shapes, stage.ConceptShapes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.ConceptsWhoseNodeIsExpanded, stage.Concepts_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.ConceptsWhoseDeliverablesNodeIsExpanded, stage.Concepts_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.DeliverableConceptShapes, stage.DeliverableConceptShapes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Diagram_Shapes, stage.DiagramShapes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.DiagramsWhoseNodeIsExpanded, stage.Diagrams_instance)
}

func (reference *DiagramShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Diagram, stage.Diagrams_instance)
	// insertion point for slice of pointers fields
}

func (reference *Library) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.RootDeliverables, stage.Deliverables_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.RootConcerns, stage.Concerns_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.RootStakeholders, stage.Stakeholders_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.RootRequirements, stage.Requirements_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.RootConcepts, stage.Concepts_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.AnalysisNeeds, stage.AnalysisNeeds_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Notes, stage.Notes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Diagrams, stage.Diagrams_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.SubLibraries, stage.Librarys_instance)
}

func (reference *Note) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Deliverables, stage.Deliverables_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Tasks, stage.Concerns_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Resources, stage.Stakeholders_instance)
}

func (reference *NoteDeliverableShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Note, stage.Notes_instance)
	__gong__reconstructPointerFromInstance(&reference.Deliverable, stage.Deliverables_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.ControlPointShapes, stage.ControlPointShapes_instance)
}

func (reference *NoteShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Note, stage.Notes_instance)
	// insertion point for slice of pointers fields
}

func (reference *NoteStakeholderShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Note, stage.Notes_instance)
	__gong__reconstructPointerFromInstance(&reference.Stakeholder, stage.Stakeholders_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.ControlPointShapes, stage.ControlPointShapes_instance)
}

func (reference *NoteTaskShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Note, stage.Notes_instance)
	__gong__reconstructPointerFromInstance(&reference.Task, stage.Concerns_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.ControlPointShapes, stage.ControlPointShapes_instance)
}

func (reference *Requirement) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.SupportLevels, stage.SupportLevels_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Concepts, stage.Concepts_instance)
}

func (reference *RequirementShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Requirement, stage.Requirements_instance)
	// insertion point for slice of pointers fields
}

func (reference *Stakeholder) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Concerns, stage.Concerns_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.SubStakeholders, stage.Stakeholders_instance)
}

func (reference *StakeholderCompositionShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Stakeholder, stage.Stakeholders_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.ControlPointShapes, stage.ControlPointShapes_instance)
}

func (reference *StakeholderConcernShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Stakeholder, stage.Stakeholders_instance)
	__gong__reconstructPointerFromInstance(&reference.Concern, stage.Concerns_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.ControlPointShapes, stage.ControlPointShapes_instance)
}

func (reference *StakeholderShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Stakeholder, stage.Stakeholders_instance)
	// insertion point for slice of pointers fields
}

func (reference *SupportLevel) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Tool, stage.Tools_instance)
	// insertion point for slice of pointers fields
}

func (reference *Tool) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

// insertion point for diff per struct
// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (analysisneed *AnalysisNeed) GongDiff(stage *Stage, analysisneedOther *AnalysisNeed) (diffs []string) {
	// insertion point for field diffs
	if analysisneed.Name != analysisneedOther.Name {
		diffs = append(diffs, analysisneed.GongMarshallField(stage, "Name"))
	}
	if analysisneed.ComputedPrefix != analysisneedOther.ComputedPrefix {
		diffs = append(diffs, analysisneed.GongMarshallField(stage, "ComputedPrefix"))
	}
	if analysisneed.IsExpanded != analysisneedOther.IsExpanded {
		diffs = append(diffs, analysisneed.GongMarshallField(stage, "IsExpanded"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (concept *Concept) GongDiff(stage *Stage, conceptOther *Concept) (diffs []string) {
	// insertion point for field diffs
	if concept.Name != conceptOther.Name {
		diffs = append(diffs, concept.GongMarshallField(stage, "Name"))
	}
	if concept.ComputedPrefix != conceptOther.ComputedPrefix {
		diffs = append(diffs, concept.GongMarshallField(stage, "ComputedPrefix"))
	}
	if concept.IsExpanded != conceptOther.IsExpanded {
		diffs = append(diffs, concept.GongMarshallField(stage, "IsExpanded"))
	}
	if ops := __gong__diffSliceOfPointers(stage, concept, "Tools", conceptOther.Tools, concept.Tools); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (conceptshape *ConceptShape) GongDiff(stage *Stage, conceptshapeOther *ConceptShape) (diffs []string) {
	// insertion point for field diffs
	if conceptshape.Name != conceptshapeOther.Name {
		diffs = append(diffs, conceptshape.GongMarshallField(stage, "Name"))
	}
	if conceptshape.Concept != conceptshapeOther.Concept {
		diffs = append(diffs, conceptshape.GongMarshallField(stage, "Concept"))
	}
	if conceptshape.IsExpanded != conceptshapeOther.IsExpanded {
		diffs = append(diffs, conceptshape.GongMarshallField(stage, "IsExpanded"))
	}
	if conceptshape.X != conceptshapeOther.X {
		diffs = append(diffs, conceptshape.GongMarshallField(stage, "X"))
	}
	if conceptshape.Y != conceptshapeOther.Y {
		diffs = append(diffs, conceptshape.GongMarshallField(stage, "Y"))
	}
	if conceptshape.Width != conceptshapeOther.Width {
		diffs = append(diffs, conceptshape.GongMarshallField(stage, "Width"))
	}
	if conceptshape.Height != conceptshapeOther.Height {
		diffs = append(diffs, conceptshape.GongMarshallField(stage, "Height"))
	}
	if conceptshape.IsHidden != conceptshapeOther.IsHidden {
		diffs = append(diffs, conceptshape.GongMarshallField(stage, "IsHidden"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (concern *Concern) GongDiff(stage *Stage, concernOther *Concern) (diffs []string) {
	// insertion point for field diffs
	if concern.Name != concernOther.Name {
		diffs = append(diffs, concern.GongMarshallField(stage, "Name"))
	}
	if concern.IDAirbus != concernOther.IDAirbus {
		diffs = append(diffs, concern.GongMarshallField(stage, "IDAirbus"))
	}
	if concern.Priority != concernOther.Priority {
		diffs = append(diffs, concern.GongMarshallField(stage, "Priority"))
	}
	if concern.ComputedPrefix != concernOther.ComputedPrefix {
		diffs = append(diffs, concern.GongMarshallField(stage, "ComputedPrefix"))
	}
	if concern.IsExpanded != concernOther.IsExpanded {
		diffs = append(diffs, concern.GongMarshallField(stage, "IsExpanded"))
	}
	if concern.Description != concernOther.Description {
		diffs = append(diffs, concern.GongMarshallField(stage, "Description"))
	}
	if ops := __gong__diffSliceOfPointers(stage, concern, "SubConcerns", concernOther.SubConcerns, concern.SubConcerns); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, concern, "Inputs", concernOther.Inputs, concern.Inputs); ops != "" {
		diffs = append(diffs, ops)
	}
	if concern.IsInputsNodeExpanded != concernOther.IsInputsNodeExpanded {
		diffs = append(diffs, concern.GongMarshallField(stage, "IsInputsNodeExpanded"))
	}
	if ops := __gong__diffSliceOfPointers(stage, concern, "Outputs", concernOther.Outputs, concern.Outputs); ops != "" {
		diffs = append(diffs, ops)
	}
	if concern.IsOutputsNodeExpanded != concernOther.IsOutputsNodeExpanded {
		diffs = append(diffs, concern.GongMarshallField(stage, "IsOutputsNodeExpanded"))
	}
	if concern.IsWithCompletion != concernOther.IsWithCompletion {
		diffs = append(diffs, concern.GongMarshallField(stage, "IsWithCompletion"))
	}
	if concern.Completion != concernOther.Completion {
		diffs = append(diffs, concern.GongMarshallField(stage, "Completion"))
	}
	if ops := __gong__diffSliceOfPointers(stage, concern, "Requirements", concernOther.Requirements, concern.Requirements); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (concerncompositionshape *ConcernCompositionShape) GongDiff(stage *Stage, concerncompositionshapeOther *ConcernCompositionShape) (diffs []string) {
	// insertion point for field diffs
	if concerncompositionshape.Name != concerncompositionshapeOther.Name {
		diffs = append(diffs, concerncompositionshape.GongMarshallField(stage, "Name"))
	}
	if concerncompositionshape.Concern != concerncompositionshapeOther.Concern {
		diffs = append(diffs, concerncompositionshape.GongMarshallField(stage, "Concern"))
	}
	if concerncompositionshape.StartRatio != concerncompositionshapeOther.StartRatio {
		diffs = append(diffs, concerncompositionshape.GongMarshallField(stage, "StartRatio"))
	}
	if concerncompositionshape.EndRatio != concerncompositionshapeOther.EndRatio {
		diffs = append(diffs, concerncompositionshape.GongMarshallField(stage, "EndRatio"))
	}
	if concerncompositionshape.StartOrientation != concerncompositionshapeOther.StartOrientation {
		diffs = append(diffs, concerncompositionshape.GongMarshallField(stage, "StartOrientation"))
	}
	if concerncompositionshape.EndOrientation != concerncompositionshapeOther.EndOrientation {
		diffs = append(diffs, concerncompositionshape.GongMarshallField(stage, "EndOrientation"))
	}
	if concerncompositionshape.CornerOffsetRatio != concerncompositionshapeOther.CornerOffsetRatio {
		diffs = append(diffs, concerncompositionshape.GongMarshallField(stage, "CornerOffsetRatio"))
	}
	if concerncompositionshape.IsHidden != concerncompositionshapeOther.IsHidden {
		diffs = append(diffs, concerncompositionshape.GongMarshallField(stage, "IsHidden"))
	}
	if ops := __gong__diffSliceOfPointers(stage, concerncompositionshape, "ControlPointShapes", concerncompositionshapeOther.ControlPointShapes, concerncompositionshape.ControlPointShapes); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (concerninputshape *ConcernInputShape) GongDiff(stage *Stage, concerninputshapeOther *ConcernInputShape) (diffs []string) {
	// insertion point for field diffs
	if concerninputshape.Name != concerninputshapeOther.Name {
		diffs = append(diffs, concerninputshape.GongMarshallField(stage, "Name"))
	}
	if concerninputshape.Deliverable != concerninputshapeOther.Deliverable {
		diffs = append(diffs, concerninputshape.GongMarshallField(stage, "Deliverable"))
	}
	if concerninputshape.Concern != concerninputshapeOther.Concern {
		diffs = append(diffs, concerninputshape.GongMarshallField(stage, "Concern"))
	}
	if concerninputshape.StartRatio != concerninputshapeOther.StartRatio {
		diffs = append(diffs, concerninputshape.GongMarshallField(stage, "StartRatio"))
	}
	if concerninputshape.EndRatio != concerninputshapeOther.EndRatio {
		diffs = append(diffs, concerninputshape.GongMarshallField(stage, "EndRatio"))
	}
	if concerninputshape.StartOrientation != concerninputshapeOther.StartOrientation {
		diffs = append(diffs, concerninputshape.GongMarshallField(stage, "StartOrientation"))
	}
	if concerninputshape.EndOrientation != concerninputshapeOther.EndOrientation {
		diffs = append(diffs, concerninputshape.GongMarshallField(stage, "EndOrientation"))
	}
	if concerninputshape.CornerOffsetRatio != concerninputshapeOther.CornerOffsetRatio {
		diffs = append(diffs, concerninputshape.GongMarshallField(stage, "CornerOffsetRatio"))
	}
	if concerninputshape.IsHidden != concerninputshapeOther.IsHidden {
		diffs = append(diffs, concerninputshape.GongMarshallField(stage, "IsHidden"))
	}
	if ops := __gong__diffSliceOfPointers(stage, concerninputshape, "ControlPointShapes", concerninputshapeOther.ControlPointShapes, concerninputshape.ControlPointShapes); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (concernoutputshape *ConcernOutputShape) GongDiff(stage *Stage, concernoutputshapeOther *ConcernOutputShape) (diffs []string) {
	// insertion point for field diffs
	if concernoutputshape.Name != concernoutputshapeOther.Name {
		diffs = append(diffs, concernoutputshape.GongMarshallField(stage, "Name"))
	}
	if concernoutputshape.Concern != concernoutputshapeOther.Concern {
		diffs = append(diffs, concernoutputshape.GongMarshallField(stage, "Concern"))
	}
	if concernoutputshape.Deliverable != concernoutputshapeOther.Deliverable {
		diffs = append(diffs, concernoutputshape.GongMarshallField(stage, "Deliverable"))
	}
	if concernoutputshape.StartRatio != concernoutputshapeOther.StartRatio {
		diffs = append(diffs, concernoutputshape.GongMarshallField(stage, "StartRatio"))
	}
	if concernoutputshape.EndRatio != concernoutputshapeOther.EndRatio {
		diffs = append(diffs, concernoutputshape.GongMarshallField(stage, "EndRatio"))
	}
	if concernoutputshape.StartOrientation != concernoutputshapeOther.StartOrientation {
		diffs = append(diffs, concernoutputshape.GongMarshallField(stage, "StartOrientation"))
	}
	if concernoutputshape.EndOrientation != concernoutputshapeOther.EndOrientation {
		diffs = append(diffs, concernoutputshape.GongMarshallField(stage, "EndOrientation"))
	}
	if concernoutputshape.CornerOffsetRatio != concernoutputshapeOther.CornerOffsetRatio {
		diffs = append(diffs, concernoutputshape.GongMarshallField(stage, "CornerOffsetRatio"))
	}
	if concernoutputshape.IsHidden != concernoutputshapeOther.IsHidden {
		diffs = append(diffs, concernoutputshape.GongMarshallField(stage, "IsHidden"))
	}
	if ops := __gong__diffSliceOfPointers(stage, concernoutputshape, "ControlPointShapes", concernoutputshapeOther.ControlPointShapes, concernoutputshape.ControlPointShapes); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (concernshape *ConcernShape) GongDiff(stage *Stage, concernshapeOther *ConcernShape) (diffs []string) {
	// insertion point for field diffs
	if concernshape.Name != concernshapeOther.Name {
		diffs = append(diffs, concernshape.GongMarshallField(stage, "Name"))
	}
	if concernshape.Concern != concernshapeOther.Concern {
		diffs = append(diffs, concernshape.GongMarshallField(stage, "Concern"))
	}
	if concernshape.IsExpanded != concernshapeOther.IsExpanded {
		diffs = append(diffs, concernshape.GongMarshallField(stage, "IsExpanded"))
	}
	if concernshape.X != concernshapeOther.X {
		diffs = append(diffs, concernshape.GongMarshallField(stage, "X"))
	}
	if concernshape.Y != concernshapeOther.Y {
		diffs = append(diffs, concernshape.GongMarshallField(stage, "Y"))
	}
	if concernshape.Width != concernshapeOther.Width {
		diffs = append(diffs, concernshape.GongMarshallField(stage, "Width"))
	}
	if concernshape.Height != concernshapeOther.Height {
		diffs = append(diffs, concernshape.GongMarshallField(stage, "Height"))
	}
	if concernshape.IsHidden != concernshapeOther.IsHidden {
		diffs = append(diffs, concernshape.GongMarshallField(stage, "IsHidden"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (controlpointshape *ControlPointShape) GongDiff(stage *Stage, controlpointshapeOther *ControlPointShape) (diffs []string) {
	// insertion point for field diffs
	if controlpointshape.Name != controlpointshapeOther.Name {
		diffs = append(diffs, controlpointshape.GongMarshallField(stage, "Name"))
	}
	if controlpointshape.X_Relative != controlpointshapeOther.X_Relative {
		diffs = append(diffs, controlpointshape.GongMarshallField(stage, "X_Relative"))
	}
	if controlpointshape.Y_Relative != controlpointshapeOther.Y_Relative {
		diffs = append(diffs, controlpointshape.GongMarshallField(stage, "Y_Relative"))
	}
	if controlpointshape.IsStartShapeTheClosestShape != controlpointshapeOther.IsStartShapeTheClosestShape {
		diffs = append(diffs, controlpointshape.GongMarshallField(stage, "IsStartShapeTheClosestShape"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (deliverable *Deliverable) GongDiff(stage *Stage, deliverableOther *Deliverable) (diffs []string) {
	// insertion point for field diffs
	if deliverable.Name != deliverableOther.Name {
		diffs = append(diffs, deliverable.GongMarshallField(stage, "Name"))
	}
	if deliverable.ComputedPrefix != deliverableOther.ComputedPrefix {
		diffs = append(diffs, deliverable.GongMarshallField(stage, "ComputedPrefix"))
	}
	if deliverable.IsExpanded != deliverableOther.IsExpanded {
		diffs = append(diffs, deliverable.GongMarshallField(stage, "IsExpanded"))
	}
	if deliverable.Description != deliverableOther.Description {
		diffs = append(diffs, deliverable.GongMarshallField(stage, "Description"))
	}
	if ops := __gong__diffSliceOfPointers(stage, deliverable, "SubDeliverables", deliverableOther.SubDeliverables, deliverable.SubDeliverables); ops != "" {
		diffs = append(diffs, ops)
	}
	if deliverable.IsProducersNodeExpanded != deliverableOther.IsProducersNodeExpanded {
		diffs = append(diffs, deliverable.GongMarshallField(stage, "IsProducersNodeExpanded"))
	}
	if deliverable.IsConsumersNodeExpanded != deliverableOther.IsConsumersNodeExpanded {
		diffs = append(diffs, deliverable.GongMarshallField(stage, "IsConsumersNodeExpanded"))
	}
	if ops := __gong__diffSliceOfPointers(stage, deliverable, "Concepts", deliverableOther.Concepts, deliverable.Concepts); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (deliverablecompositionshape *DeliverableCompositionShape) GongDiff(stage *Stage, deliverablecompositionshapeOther *DeliverableCompositionShape) (diffs []string) {
	// insertion point for field diffs
	if deliverablecompositionshape.Name != deliverablecompositionshapeOther.Name {
		diffs = append(diffs, deliverablecompositionshape.GongMarshallField(stage, "Name"))
	}
	if deliverablecompositionshape.Deliverable != deliverablecompositionshapeOther.Deliverable {
		diffs = append(diffs, deliverablecompositionshape.GongMarshallField(stage, "Deliverable"))
	}
	if deliverablecompositionshape.StartRatio != deliverablecompositionshapeOther.StartRatio {
		diffs = append(diffs, deliverablecompositionshape.GongMarshallField(stage, "StartRatio"))
	}
	if deliverablecompositionshape.EndRatio != deliverablecompositionshapeOther.EndRatio {
		diffs = append(diffs, deliverablecompositionshape.GongMarshallField(stage, "EndRatio"))
	}
	if deliverablecompositionshape.StartOrientation != deliverablecompositionshapeOther.StartOrientation {
		diffs = append(diffs, deliverablecompositionshape.GongMarshallField(stage, "StartOrientation"))
	}
	if deliverablecompositionshape.EndOrientation != deliverablecompositionshapeOther.EndOrientation {
		diffs = append(diffs, deliverablecompositionshape.GongMarshallField(stage, "EndOrientation"))
	}
	if deliverablecompositionshape.CornerOffsetRatio != deliverablecompositionshapeOther.CornerOffsetRatio {
		diffs = append(diffs, deliverablecompositionshape.GongMarshallField(stage, "CornerOffsetRatio"))
	}
	if deliverablecompositionshape.IsHidden != deliverablecompositionshapeOther.IsHidden {
		diffs = append(diffs, deliverablecompositionshape.GongMarshallField(stage, "IsHidden"))
	}
	if ops := __gong__diffSliceOfPointers(stage, deliverablecompositionshape, "ControlPointShapes", deliverablecompositionshapeOther.ControlPointShapes, deliverablecompositionshape.ControlPointShapes); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (deliverableconceptshape *DeliverableConceptShape) GongDiff(stage *Stage, deliverableconceptshapeOther *DeliverableConceptShape) (diffs []string) {
	// insertion point for field diffs
	if deliverableconceptshape.Name != deliverableconceptshapeOther.Name {
		diffs = append(diffs, deliverableconceptshape.GongMarshallField(stage, "Name"))
	}
	if deliverableconceptshape.Deliverable != deliverableconceptshapeOther.Deliverable {
		diffs = append(diffs, deliverableconceptshape.GongMarshallField(stage, "Deliverable"))
	}
	if deliverableconceptshape.Concept != deliverableconceptshapeOther.Concept {
		diffs = append(diffs, deliverableconceptshape.GongMarshallField(stage, "Concept"))
	}
	if deliverableconceptshape.StartRatio != deliverableconceptshapeOther.StartRatio {
		diffs = append(diffs, deliverableconceptshape.GongMarshallField(stage, "StartRatio"))
	}
	if deliverableconceptshape.EndRatio != deliverableconceptshapeOther.EndRatio {
		diffs = append(diffs, deliverableconceptshape.GongMarshallField(stage, "EndRatio"))
	}
	if deliverableconceptshape.StartOrientation != deliverableconceptshapeOther.StartOrientation {
		diffs = append(diffs, deliverableconceptshape.GongMarshallField(stage, "StartOrientation"))
	}
	if deliverableconceptshape.EndOrientation != deliverableconceptshapeOther.EndOrientation {
		diffs = append(diffs, deliverableconceptshape.GongMarshallField(stage, "EndOrientation"))
	}
	if deliverableconceptshape.CornerOffsetRatio != deliverableconceptshapeOther.CornerOffsetRatio {
		diffs = append(diffs, deliverableconceptshape.GongMarshallField(stage, "CornerOffsetRatio"))
	}
	if deliverableconceptshape.IsHidden != deliverableconceptshapeOther.IsHidden {
		diffs = append(diffs, deliverableconceptshape.GongMarshallField(stage, "IsHidden"))
	}
	if ops := __gong__diffSliceOfPointers(stage, deliverableconceptshape, "ControlPointShapes", deliverableconceptshapeOther.ControlPointShapes, deliverableconceptshape.ControlPointShapes); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (deliverableshape *DeliverableShape) GongDiff(stage *Stage, deliverableshapeOther *DeliverableShape) (diffs []string) {
	// insertion point for field diffs
	if deliverableshape.Name != deliverableshapeOther.Name {
		diffs = append(diffs, deliverableshape.GongMarshallField(stage, "Name"))
	}
	if deliverableshape.Deliverable != deliverableshapeOther.Deliverable {
		diffs = append(diffs, deliverableshape.GongMarshallField(stage, "Deliverable"))
	}
	if deliverableshape.IsExpanded != deliverableshapeOther.IsExpanded {
		diffs = append(diffs, deliverableshape.GongMarshallField(stage, "IsExpanded"))
	}
	if deliverableshape.X != deliverableshapeOther.X {
		diffs = append(diffs, deliverableshape.GongMarshallField(stage, "X"))
	}
	if deliverableshape.Y != deliverableshapeOther.Y {
		diffs = append(diffs, deliverableshape.GongMarshallField(stage, "Y"))
	}
	if deliverableshape.Width != deliverableshapeOther.Width {
		diffs = append(diffs, deliverableshape.GongMarshallField(stage, "Width"))
	}
	if deliverableshape.Height != deliverableshapeOther.Height {
		diffs = append(diffs, deliverableshape.GongMarshallField(stage, "Height"))
	}
	if deliverableshape.IsHidden != deliverableshapeOther.IsHidden {
		diffs = append(diffs, deliverableshape.GongMarshallField(stage, "IsHidden"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (diagram *Diagram) GongDiff(stage *Stage, diagramOther *Diagram) (diffs []string) {
	// insertion point for field diffs
	if diagram.Name != diagramOther.Name {
		diffs = append(diffs, diagram.GongMarshallField(stage, "Name"))
	}
	if diagram.ComputedPrefix != diagramOther.ComputedPrefix {
		diffs = append(diffs, diagram.GongMarshallField(stage, "ComputedPrefix"))
	}
	if diagram.IsExpanded != diagramOther.IsExpanded {
		diffs = append(diffs, diagram.GongMarshallField(stage, "IsExpanded"))
	}
	if diagram.IsChecked != diagramOther.IsChecked {
		diffs = append(diffs, diagram.GongMarshallField(stage, "IsChecked"))
	}
	if diagram.IsEditable_ != diagramOther.IsEditable_ {
		diffs = append(diffs, diagram.GongMarshallField(stage, "IsEditable_"))
	}
	if diagram.ShowPrefix != diagramOther.ShowPrefix {
		diffs = append(diffs, diagram.GongMarshallField(stage, "ShowPrefix"))
	}
	if diagram.DefaultBoxWidth != diagramOther.DefaultBoxWidth {
		diffs = append(diffs, diagram.GongMarshallField(stage, "DefaultBoxWidth"))
	}
	if diagram.DefaultBoxHeigth != diagramOther.DefaultBoxHeigth {
		diffs = append(diffs, diagram.GongMarshallField(stage, "DefaultBoxHeigth"))
	}
	if diagram.Width != diagramOther.Width {
		diffs = append(diffs, diagram.GongMarshallField(stage, "Width"))
	}
	if diagram.Height != diagramOther.Height {
		diffs = append(diffs, diagram.GongMarshallField(stage, "Height"))
	}
	if ops := __gong__diffSliceOfPointers(stage, diagram, "ConcernsWhoseRequirementsNodeIsExpanded", diagramOther.ConcernsWhoseRequirementsNodeIsExpanded, diagram.ConcernsWhoseRequirementsNodeIsExpanded); ops != "" {
		diffs = append(diffs, ops)
	}
	if diagram.IsRequirementsNodeExpanded != diagramOther.IsRequirementsNodeExpanded {
		diffs = append(diffs, diagram.GongMarshallField(stage, "IsRequirementsNodeExpanded"))
	}
	if diagram.IsConceptsNodeExpanded != diagramOther.IsConceptsNodeExpanded {
		diffs = append(diffs, diagram.GongMarshallField(stage, "IsConceptsNodeExpanded"))
	}
	if ops := __gong__diffSliceOfPointers(stage, diagram, "Deliverable_Shapes", diagramOther.Deliverable_Shapes, diagram.Deliverable_Shapes); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, diagram, "DeliverablesWhoseNodeIsExpanded", diagramOther.DeliverablesWhoseNodeIsExpanded, diagram.DeliverablesWhoseNodeIsExpanded); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, diagram, "DeliverablesWhoseConceptsNodeIsExpanded", diagramOther.DeliverablesWhoseConceptsNodeIsExpanded, diagram.DeliverablesWhoseConceptsNodeIsExpanded); ops != "" {
		diffs = append(diffs, ops)
	}
	if diagram.IsPBSNodeExpanded != diagramOther.IsPBSNodeExpanded {
		diffs = append(diffs, diagram.GongMarshallField(stage, "IsPBSNodeExpanded"))
	}
	if ops := __gong__diffSliceOfPointers(stage, diagram, "DeliverableComposition_Shapes", diagramOther.DeliverableComposition_Shapes, diagram.DeliverableComposition_Shapes); ops != "" {
		diffs = append(diffs, ops)
	}
	if diagram.IsConcernsNodeExpanded != diagramOther.IsConcernsNodeExpanded {
		diffs = append(diffs, diagram.GongMarshallField(stage, "IsConcernsNodeExpanded"))
	}
	if ops := __gong__diffSliceOfPointers(stage, diagram, "Concern_Shapes", diagramOther.Concern_Shapes, diagram.Concern_Shapes); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, diagram, "ConcernsWhoseNodeIsExpanded", diagramOther.ConcernsWhoseNodeIsExpanded, diagram.ConcernsWhoseNodeIsExpanded); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, diagram, "ConcernsWhoseInputNodeIsExpanded", diagramOther.ConcernsWhoseInputNodeIsExpanded, diagram.ConcernsWhoseInputNodeIsExpanded); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, diagram, "ConcernsWhoseStakeholderNodeIsExpanded", diagramOther.ConcernsWhoseStakeholderNodeIsExpanded, diagram.ConcernsWhoseStakeholderNodeIsExpanded); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, diagram, "ConcernssWhoseOutputNodeIsExpanded", diagramOther.ConcernssWhoseOutputNodeIsExpanded, diagram.ConcernssWhoseOutputNodeIsExpanded); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, diagram, "ConcernComposition_Shapes", diagramOther.ConcernComposition_Shapes, diagram.ConcernComposition_Shapes); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, diagram, "ConcernInputShapes", diagramOther.ConcernInputShapes, diagram.ConcernInputShapes); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, diagram, "ConcernOutputShapes", diagramOther.ConcernOutputShapes, diagram.ConcernOutputShapes); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, diagram, "Note_Shapes", diagramOther.Note_Shapes, diagram.Note_Shapes); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, diagram, "NotesWhoseNodeIsExpanded", diagramOther.NotesWhoseNodeIsExpanded, diagram.NotesWhoseNodeIsExpanded); ops != "" {
		diffs = append(diffs, ops)
	}
	if diagram.IsNotesNodeExpanded != diagramOther.IsNotesNodeExpanded {
		diffs = append(diffs, diagram.GongMarshallField(stage, "IsNotesNodeExpanded"))
	}
	if ops := __gong__diffSliceOfPointers(stage, diagram, "NoteDeliverableShapes", diagramOther.NoteDeliverableShapes, diagram.NoteDeliverableShapes); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, diagram, "NoteTaskShapes", diagramOther.NoteTaskShapes, diagram.NoteTaskShapes); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, diagram, "NoteResourceShapes", diagramOther.NoteResourceShapes, diagram.NoteResourceShapes); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, diagram, "Stakeholder_Shapes", diagramOther.Stakeholder_Shapes, diagram.Stakeholder_Shapes); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, diagram, "ResourcesWhoseNodeIsExpanded", diagramOther.ResourcesWhoseNodeIsExpanded, diagram.ResourcesWhoseNodeIsExpanded); ops != "" {
		diffs = append(diffs, ops)
	}
	if diagram.IsStakeholdersNodeExpanded != diagramOther.IsStakeholdersNodeExpanded {
		diffs = append(diffs, diagram.GongMarshallField(stage, "IsStakeholdersNodeExpanded"))
	}
	if ops := __gong__diffSliceOfPointers(stage, diagram, "ResourceComposition_Shapes", diagramOther.ResourceComposition_Shapes, diagram.ResourceComposition_Shapes); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, diagram, "StakeholderConcernShapes", diagramOther.StakeholderConcernShapes, diagram.StakeholderConcernShapes); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, diagram, "Requirement_Shapes", diagramOther.Requirement_Shapes, diagram.Requirement_Shapes); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, diagram, "RequirementsWhoseNodeIsExpanded", diagramOther.RequirementsWhoseNodeIsExpanded, diagram.RequirementsWhoseNodeIsExpanded); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, diagram, "Concept_Shapes", diagramOther.Concept_Shapes, diagram.Concept_Shapes); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, diagram, "ConceptsWhoseNodeIsExpanded", diagramOther.ConceptsWhoseNodeIsExpanded, diagram.ConceptsWhoseNodeIsExpanded); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, diagram, "ConceptsWhoseDeliverablesNodeIsExpanded", diagramOther.ConceptsWhoseDeliverablesNodeIsExpanded, diagram.ConceptsWhoseDeliverablesNodeIsExpanded); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, diagram, "DeliverableConceptShapes", diagramOther.DeliverableConceptShapes, diagram.DeliverableConceptShapes); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, diagram, "Diagram_Shapes", diagramOther.Diagram_Shapes, diagram.Diagram_Shapes); ops != "" {
		diffs = append(diffs, ops)
	}
	if diagram.IsDiagramsNodeExpanded != diagramOther.IsDiagramsNodeExpanded {
		diffs = append(diffs, diagram.GongMarshallField(stage, "IsDiagramsNodeExpanded"))
	}
	if ops := __gong__diffSliceOfPointers(stage, diagram, "DiagramsWhoseNodeIsExpanded", diagramOther.DiagramsWhoseNodeIsExpanded, diagram.DiagramsWhoseNodeIsExpanded); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (diagramshape *DiagramShape) GongDiff(stage *Stage, diagramshapeOther *DiagramShape) (diffs []string) {
	// insertion point for field diffs
	if diagramshape.Name != diagramshapeOther.Name {
		diffs = append(diffs, diagramshape.GongMarshallField(stage, "Name"))
	}
	if diagramshape.Diagram != diagramshapeOther.Diagram {
		diffs = append(diffs, diagramshape.GongMarshallField(stage, "Diagram"))
	}
	if diagramshape.IsExpanded != diagramshapeOther.IsExpanded {
		diffs = append(diffs, diagramshape.GongMarshallField(stage, "IsExpanded"))
	}
	if diagramshape.X != diagramshapeOther.X {
		diffs = append(diffs, diagramshape.GongMarshallField(stage, "X"))
	}
	if diagramshape.Y != diagramshapeOther.Y {
		diffs = append(diffs, diagramshape.GongMarshallField(stage, "Y"))
	}
	if diagramshape.Width != diagramshapeOther.Width {
		diffs = append(diffs, diagramshape.GongMarshallField(stage, "Width"))
	}
	if diagramshape.Height != diagramshapeOther.Height {
		diffs = append(diffs, diagramshape.GongMarshallField(stage, "Height"))
	}
	if diagramshape.IsHidden != diagramshapeOther.IsHidden {
		diffs = append(diffs, diagramshape.GongMarshallField(stage, "IsHidden"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (library *Library) GongDiff(stage *Stage, libraryOther *Library) (diffs []string) {
	// insertion point for field diffs
	if library.Name != libraryOther.Name {
		diffs = append(diffs, library.GongMarshallField(stage, "Name"))
	}
	if library.IsRootLibrary != libraryOther.IsRootLibrary {
		diffs = append(diffs, library.GongMarshallField(stage, "IsRootLibrary"))
	}
	if library.ComputedPrefix != libraryOther.ComputedPrefix {
		diffs = append(diffs, library.GongMarshallField(stage, "ComputedPrefix"))
	}
	if library.IsExpanded != libraryOther.IsExpanded {
		diffs = append(diffs, library.GongMarshallField(stage, "IsExpanded"))
	}
	if ops := __gong__diffSliceOfPointers(stage, library, "RootDeliverables", libraryOther.RootDeliverables, library.RootDeliverables); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, library, "RootConcerns", libraryOther.RootConcerns, library.RootConcerns); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, library, "RootStakeholders", libraryOther.RootStakeholders, library.RootStakeholders); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, library, "RootRequirements", libraryOther.RootRequirements, library.RootRequirements); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, library, "RootConcepts", libraryOther.RootConcepts, library.RootConcepts); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, library, "AnalysisNeeds", libraryOther.AnalysisNeeds, library.AnalysisNeeds); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, library, "Notes", libraryOther.Notes, library.Notes); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, library, "Diagrams", libraryOther.Diagrams, library.Diagrams); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, library, "SubLibraries", libraryOther.SubLibraries, library.SubLibraries); ops != "" {
		diffs = append(diffs, ops)
	}
	if library.NbPixPerCharacter != libraryOther.NbPixPerCharacter {
		diffs = append(diffs, library.GongMarshallField(stage, "NbPixPerCharacter"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (note *Note) GongDiff(stage *Stage, noteOther *Note) (diffs []string) {
	// insertion point for field diffs
	if note.Name != noteOther.Name {
		diffs = append(diffs, note.GongMarshallField(stage, "Name"))
	}
	if note.ComputedPrefix != noteOther.ComputedPrefix {
		diffs = append(diffs, note.GongMarshallField(stage, "ComputedPrefix"))
	}
	if note.IsExpanded != noteOther.IsExpanded {
		diffs = append(diffs, note.GongMarshallField(stage, "IsExpanded"))
	}
	if ops := __gong__diffSliceOfPointers(stage, note, "Deliverables", noteOther.Deliverables, note.Deliverables); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, note, "Tasks", noteOther.Tasks, note.Tasks); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, note, "Resources", noteOther.Resources, note.Resources); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (notedeliverableshape *NoteDeliverableShape) GongDiff(stage *Stage, notedeliverableshapeOther *NoteDeliverableShape) (diffs []string) {
	// insertion point for field diffs
	if notedeliverableshape.Name != notedeliverableshapeOther.Name {
		diffs = append(diffs, notedeliverableshape.GongMarshallField(stage, "Name"))
	}
	if notedeliverableshape.Note != notedeliverableshapeOther.Note {
		diffs = append(diffs, notedeliverableshape.GongMarshallField(stage, "Note"))
	}
	if notedeliverableshape.Deliverable != notedeliverableshapeOther.Deliverable {
		diffs = append(diffs, notedeliverableshape.GongMarshallField(stage, "Deliverable"))
	}
	if notedeliverableshape.StartRatio != notedeliverableshapeOther.StartRatio {
		diffs = append(diffs, notedeliverableshape.GongMarshallField(stage, "StartRatio"))
	}
	if notedeliverableshape.EndRatio != notedeliverableshapeOther.EndRatio {
		diffs = append(diffs, notedeliverableshape.GongMarshallField(stage, "EndRatio"))
	}
	if notedeliverableshape.StartOrientation != notedeliverableshapeOther.StartOrientation {
		diffs = append(diffs, notedeliverableshape.GongMarshallField(stage, "StartOrientation"))
	}
	if notedeliverableshape.EndOrientation != notedeliverableshapeOther.EndOrientation {
		diffs = append(diffs, notedeliverableshape.GongMarshallField(stage, "EndOrientation"))
	}
	if notedeliverableshape.CornerOffsetRatio != notedeliverableshapeOther.CornerOffsetRatio {
		diffs = append(diffs, notedeliverableshape.GongMarshallField(stage, "CornerOffsetRatio"))
	}
	if notedeliverableshape.IsHidden != notedeliverableshapeOther.IsHidden {
		diffs = append(diffs, notedeliverableshape.GongMarshallField(stage, "IsHidden"))
	}
	if ops := __gong__diffSliceOfPointers(stage, notedeliverableshape, "ControlPointShapes", notedeliverableshapeOther.ControlPointShapes, notedeliverableshape.ControlPointShapes); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (noteshape *NoteShape) GongDiff(stage *Stage, noteshapeOther *NoteShape) (diffs []string) {
	// insertion point for field diffs
	if noteshape.Name != noteshapeOther.Name {
		diffs = append(diffs, noteshape.GongMarshallField(stage, "Name"))
	}
	if noteshape.Note != noteshapeOther.Note {
		diffs = append(diffs, noteshape.GongMarshallField(stage, "Note"))
	}
	if noteshape.IsExpanded != noteshapeOther.IsExpanded {
		diffs = append(diffs, noteshape.GongMarshallField(stage, "IsExpanded"))
	}
	if noteshape.X != noteshapeOther.X {
		diffs = append(diffs, noteshape.GongMarshallField(stage, "X"))
	}
	if noteshape.Y != noteshapeOther.Y {
		diffs = append(diffs, noteshape.GongMarshallField(stage, "Y"))
	}
	if noteshape.Width != noteshapeOther.Width {
		diffs = append(diffs, noteshape.GongMarshallField(stage, "Width"))
	}
	if noteshape.Height != noteshapeOther.Height {
		diffs = append(diffs, noteshape.GongMarshallField(stage, "Height"))
	}
	if noteshape.IsHidden != noteshapeOther.IsHidden {
		diffs = append(diffs, noteshape.GongMarshallField(stage, "IsHidden"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (notestakeholdershape *NoteStakeholderShape) GongDiff(stage *Stage, notestakeholdershapeOther *NoteStakeholderShape) (diffs []string) {
	// insertion point for field diffs
	if notestakeholdershape.Name != notestakeholdershapeOther.Name {
		diffs = append(diffs, notestakeholdershape.GongMarshallField(stage, "Name"))
	}
	if notestakeholdershape.Note != notestakeholdershapeOther.Note {
		diffs = append(diffs, notestakeholdershape.GongMarshallField(stage, "Note"))
	}
	if notestakeholdershape.Stakeholder != notestakeholdershapeOther.Stakeholder {
		diffs = append(diffs, notestakeholdershape.GongMarshallField(stage, "Stakeholder"))
	}
	if notestakeholdershape.StartRatio != notestakeholdershapeOther.StartRatio {
		diffs = append(diffs, notestakeholdershape.GongMarshallField(stage, "StartRatio"))
	}
	if notestakeholdershape.EndRatio != notestakeholdershapeOther.EndRatio {
		diffs = append(diffs, notestakeholdershape.GongMarshallField(stage, "EndRatio"))
	}
	if notestakeholdershape.StartOrientation != notestakeholdershapeOther.StartOrientation {
		diffs = append(diffs, notestakeholdershape.GongMarshallField(stage, "StartOrientation"))
	}
	if notestakeholdershape.EndOrientation != notestakeholdershapeOther.EndOrientation {
		diffs = append(diffs, notestakeholdershape.GongMarshallField(stage, "EndOrientation"))
	}
	if notestakeholdershape.CornerOffsetRatio != notestakeholdershapeOther.CornerOffsetRatio {
		diffs = append(diffs, notestakeholdershape.GongMarshallField(stage, "CornerOffsetRatio"))
	}
	if notestakeholdershape.IsHidden != notestakeholdershapeOther.IsHidden {
		diffs = append(diffs, notestakeholdershape.GongMarshallField(stage, "IsHidden"))
	}
	if ops := __gong__diffSliceOfPointers(stage, notestakeholdershape, "ControlPointShapes", notestakeholdershapeOther.ControlPointShapes, notestakeholdershape.ControlPointShapes); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (notetaskshape *NoteTaskShape) GongDiff(stage *Stage, notetaskshapeOther *NoteTaskShape) (diffs []string) {
	// insertion point for field diffs
	if notetaskshape.Name != notetaskshapeOther.Name {
		diffs = append(diffs, notetaskshape.GongMarshallField(stage, "Name"))
	}
	if notetaskshape.Note != notetaskshapeOther.Note {
		diffs = append(diffs, notetaskshape.GongMarshallField(stage, "Note"))
	}
	if notetaskshape.Task != notetaskshapeOther.Task {
		diffs = append(diffs, notetaskshape.GongMarshallField(stage, "Task"))
	}
	if notetaskshape.StartRatio != notetaskshapeOther.StartRatio {
		diffs = append(diffs, notetaskshape.GongMarshallField(stage, "StartRatio"))
	}
	if notetaskshape.EndRatio != notetaskshapeOther.EndRatio {
		diffs = append(diffs, notetaskshape.GongMarshallField(stage, "EndRatio"))
	}
	if notetaskshape.StartOrientation != notetaskshapeOther.StartOrientation {
		diffs = append(diffs, notetaskshape.GongMarshallField(stage, "StartOrientation"))
	}
	if notetaskshape.EndOrientation != notetaskshapeOther.EndOrientation {
		diffs = append(diffs, notetaskshape.GongMarshallField(stage, "EndOrientation"))
	}
	if notetaskshape.CornerOffsetRatio != notetaskshapeOther.CornerOffsetRatio {
		diffs = append(diffs, notetaskshape.GongMarshallField(stage, "CornerOffsetRatio"))
	}
	if notetaskshape.IsHidden != notetaskshapeOther.IsHidden {
		diffs = append(diffs, notetaskshape.GongMarshallField(stage, "IsHidden"))
	}
	if ops := __gong__diffSliceOfPointers(stage, notetaskshape, "ControlPointShapes", notetaskshapeOther.ControlPointShapes, notetaskshape.ControlPointShapes); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (requirement *Requirement) GongDiff(stage *Stage, requirementOther *Requirement) (diffs []string) {
	// insertion point for field diffs
	if requirement.Name != requirementOther.Name {
		diffs = append(diffs, requirement.GongMarshallField(stage, "Name"))
	}
	if requirement.ComputedPrefix != requirementOther.ComputedPrefix {
		diffs = append(diffs, requirement.GongMarshallField(stage, "ComputedPrefix"))
	}
	if requirement.IsExpanded != requirementOther.IsExpanded {
		diffs = append(diffs, requirement.GongMarshallField(stage, "IsExpanded"))
	}
	if ops := __gong__diffSliceOfPointers(stage, requirement, "SupportLevels", requirementOther.SupportLevels, requirement.SupportLevels); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, requirement, "Concepts", requirementOther.Concepts, requirement.Concepts); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (requirementshape *RequirementShape) GongDiff(stage *Stage, requirementshapeOther *RequirementShape) (diffs []string) {
	// insertion point for field diffs
	if requirementshape.Name != requirementshapeOther.Name {
		diffs = append(diffs, requirementshape.GongMarshallField(stage, "Name"))
	}
	if requirementshape.Requirement != requirementshapeOther.Requirement {
		diffs = append(diffs, requirementshape.GongMarshallField(stage, "Requirement"))
	}
	if requirementshape.IsExpanded != requirementshapeOther.IsExpanded {
		diffs = append(diffs, requirementshape.GongMarshallField(stage, "IsExpanded"))
	}
	if requirementshape.X != requirementshapeOther.X {
		diffs = append(diffs, requirementshape.GongMarshallField(stage, "X"))
	}
	if requirementshape.Y != requirementshapeOther.Y {
		diffs = append(diffs, requirementshape.GongMarshallField(stage, "Y"))
	}
	if requirementshape.Width != requirementshapeOther.Width {
		diffs = append(diffs, requirementshape.GongMarshallField(stage, "Width"))
	}
	if requirementshape.Height != requirementshapeOther.Height {
		diffs = append(diffs, requirementshape.GongMarshallField(stage, "Height"))
	}
	if requirementshape.IsHidden != requirementshapeOther.IsHidden {
		diffs = append(diffs, requirementshape.GongMarshallField(stage, "IsHidden"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (stakeholder *Stakeholder) GongDiff(stage *Stage, stakeholderOther *Stakeholder) (diffs []string) {
	// insertion point for field diffs
	if stakeholder.Name != stakeholderOther.Name {
		diffs = append(diffs, stakeholder.GongMarshallField(stage, "Name"))
	}
	if stakeholder.IDAirbus != stakeholderOther.IDAirbus {
		diffs = append(diffs, stakeholder.GongMarshallField(stage, "IDAirbus"))
	}
	if stakeholder.ComputedPrefix != stakeholderOther.ComputedPrefix {
		diffs = append(diffs, stakeholder.GongMarshallField(stage, "ComputedPrefix"))
	}
	if stakeholder.IsExpanded != stakeholderOther.IsExpanded {
		diffs = append(diffs, stakeholder.GongMarshallField(stage, "IsExpanded"))
	}
	if stakeholder.Description != stakeholderOther.Description {
		diffs = append(diffs, stakeholder.GongMarshallField(stage, "Description"))
	}
	if ops := __gong__diffSliceOfPointers(stage, stakeholder, "Concerns", stakeholderOther.Concerns, stakeholder.Concerns); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, stakeholder, "SubStakeholders", stakeholderOther.SubStakeholders, stakeholder.SubStakeholders); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (stakeholdercompositionshape *StakeholderCompositionShape) GongDiff(stage *Stage, stakeholdercompositionshapeOther *StakeholderCompositionShape) (diffs []string) {
	// insertion point for field diffs
	if stakeholdercompositionshape.Name != stakeholdercompositionshapeOther.Name {
		diffs = append(diffs, stakeholdercompositionshape.GongMarshallField(stage, "Name"))
	}
	if stakeholdercompositionshape.Stakeholder != stakeholdercompositionshapeOther.Stakeholder {
		diffs = append(diffs, stakeholdercompositionshape.GongMarshallField(stage, "Stakeholder"))
	}
	if stakeholdercompositionshape.StartRatio != stakeholdercompositionshapeOther.StartRatio {
		diffs = append(diffs, stakeholdercompositionshape.GongMarshallField(stage, "StartRatio"))
	}
	if stakeholdercompositionshape.EndRatio != stakeholdercompositionshapeOther.EndRatio {
		diffs = append(diffs, stakeholdercompositionshape.GongMarshallField(stage, "EndRatio"))
	}
	if stakeholdercompositionshape.StartOrientation != stakeholdercompositionshapeOther.StartOrientation {
		diffs = append(diffs, stakeholdercompositionshape.GongMarshallField(stage, "StartOrientation"))
	}
	if stakeholdercompositionshape.EndOrientation != stakeholdercompositionshapeOther.EndOrientation {
		diffs = append(diffs, stakeholdercompositionshape.GongMarshallField(stage, "EndOrientation"))
	}
	if stakeholdercompositionshape.CornerOffsetRatio != stakeholdercompositionshapeOther.CornerOffsetRatio {
		diffs = append(diffs, stakeholdercompositionshape.GongMarshallField(stage, "CornerOffsetRatio"))
	}
	if stakeholdercompositionshape.IsHidden != stakeholdercompositionshapeOther.IsHidden {
		diffs = append(diffs, stakeholdercompositionshape.GongMarshallField(stage, "IsHidden"))
	}
	if ops := __gong__diffSliceOfPointers(stage, stakeholdercompositionshape, "ControlPointShapes", stakeholdercompositionshapeOther.ControlPointShapes, stakeholdercompositionshape.ControlPointShapes); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (stakeholderconcernshape *StakeholderConcernShape) GongDiff(stage *Stage, stakeholderconcernshapeOther *StakeholderConcernShape) (diffs []string) {
	// insertion point for field diffs
	if stakeholderconcernshape.Name != stakeholderconcernshapeOther.Name {
		diffs = append(diffs, stakeholderconcernshape.GongMarshallField(stage, "Name"))
	}
	if stakeholderconcernshape.Stakeholder != stakeholderconcernshapeOther.Stakeholder {
		diffs = append(diffs, stakeholderconcernshape.GongMarshallField(stage, "Stakeholder"))
	}
	if stakeholderconcernshape.Concern != stakeholderconcernshapeOther.Concern {
		diffs = append(diffs, stakeholderconcernshape.GongMarshallField(stage, "Concern"))
	}
	if stakeholderconcernshape.StartRatio != stakeholderconcernshapeOther.StartRatio {
		diffs = append(diffs, stakeholderconcernshape.GongMarshallField(stage, "StartRatio"))
	}
	if stakeholderconcernshape.EndRatio != stakeholderconcernshapeOther.EndRatio {
		diffs = append(diffs, stakeholderconcernshape.GongMarshallField(stage, "EndRatio"))
	}
	if stakeholderconcernshape.StartOrientation != stakeholderconcernshapeOther.StartOrientation {
		diffs = append(diffs, stakeholderconcernshape.GongMarshallField(stage, "StartOrientation"))
	}
	if stakeholderconcernshape.EndOrientation != stakeholderconcernshapeOther.EndOrientation {
		diffs = append(diffs, stakeholderconcernshape.GongMarshallField(stage, "EndOrientation"))
	}
	if stakeholderconcernshape.CornerOffsetRatio != stakeholderconcernshapeOther.CornerOffsetRatio {
		diffs = append(diffs, stakeholderconcernshape.GongMarshallField(stage, "CornerOffsetRatio"))
	}
	if stakeholderconcernshape.IsHidden != stakeholderconcernshapeOther.IsHidden {
		diffs = append(diffs, stakeholderconcernshape.GongMarshallField(stage, "IsHidden"))
	}
	if ops := __gong__diffSliceOfPointers(stage, stakeholderconcernshape, "ControlPointShapes", stakeholderconcernshapeOther.ControlPointShapes, stakeholderconcernshape.ControlPointShapes); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (stakeholdershape *StakeholderShape) GongDiff(stage *Stage, stakeholdershapeOther *StakeholderShape) (diffs []string) {
	// insertion point for field diffs
	if stakeholdershape.Name != stakeholdershapeOther.Name {
		diffs = append(diffs, stakeholdershape.GongMarshallField(stage, "Name"))
	}
	if stakeholdershape.Stakeholder != stakeholdershapeOther.Stakeholder {
		diffs = append(diffs, stakeholdershape.GongMarshallField(stage, "Stakeholder"))
	}
	if stakeholdershape.IsExpanded != stakeholdershapeOther.IsExpanded {
		diffs = append(diffs, stakeholdershape.GongMarshallField(stage, "IsExpanded"))
	}
	if stakeholdershape.X != stakeholdershapeOther.X {
		diffs = append(diffs, stakeholdershape.GongMarshallField(stage, "X"))
	}
	if stakeholdershape.Y != stakeholdershapeOther.Y {
		diffs = append(diffs, stakeholdershape.GongMarshallField(stage, "Y"))
	}
	if stakeholdershape.Width != stakeholdershapeOther.Width {
		diffs = append(diffs, stakeholdershape.GongMarshallField(stage, "Width"))
	}
	if stakeholdershape.Height != stakeholdershapeOther.Height {
		diffs = append(diffs, stakeholdershape.GongMarshallField(stage, "Height"))
	}
	if stakeholdershape.IsHidden != stakeholdershapeOther.IsHidden {
		diffs = append(diffs, stakeholdershape.GongMarshallField(stage, "IsHidden"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (supportlevel *SupportLevel) GongDiff(stage *Stage, supportlevelOther *SupportLevel) (diffs []string) {
	// insertion point for field diffs
	if supportlevel.Name != supportlevelOther.Name {
		diffs = append(diffs, supportlevel.GongMarshallField(stage, "Name"))
	}
	if supportlevel.Tool != supportlevelOther.Tool {
		diffs = append(diffs, supportlevel.GongMarshallField(stage, "Tool"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (tool *Tool) GongDiff(stage *Stage, toolOther *Tool) (diffs []string) {
	// insertion point for field diffs
	if tool.Name != toolOther.Name {
		diffs = append(diffs, tool.GongMarshallField(stage, "Name"))
	}

	return
}

// Diff is the Stage method that returns the sequence of operations to transform oldSlice into newSlice.
func (stage *Stage) Diff(
	a GongstructIF,
	fieldName string,
	lenOld, lenNew int,
	equal func(i, j int) bool,
	getNewIdentifier func(j int) string,
) (ops string) {
	m, n := lenOld, lenNew

	// 1. Build the LCS (Longest Common Subsequence) Matrix
	// This helps us find the "anchor" elements that shouldn't move.
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := range m {
		for j := range n {
			if equal(i, j) {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				// Take the maximum of previous options
				dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
			}
		}
	}

	// 2. Backtrack to find which indices in oldSlice are part of the LCS
	// We use a map for O(1) lookups.
	keptIndices := make(map[int]bool)
	i, j := m, n
	for i > 0 && j > 0 {
		if equal(i-1, j-1) {
			keptIndices[i-1] = true
			i--
			j--
		} else if dp[i-1][j] > dp[i][j-1] {
			i--
		} else {
			j--
		}
	}

	// 3. PHASE 1: Generate Deletions
	// MUST go from High Index -> Low Index to preserve validity of lower indices.
	for k := m - 1; k >= 0; k-- {
		if !keptIndices[k] {
			ops += fmt.Sprintf("\n\t%s.%s = slices.Delete( %s.%s, %d, %d)", a.GongGetReferenceIdentifier(stage), fieldName, a.GongGetReferenceIdentifier(stage), fieldName, k, k+1)
		}
	}

	// 4. PHASE 2: Generate Insertions
	// We simulate the state of the slice after deletions to determine insertion points.
	// The 'current' slice essentially consists of only the kept LCS items.

	// Track kept indices in old slice
	keptOldIndices := make([]int, 0, len(keptIndices))
	for k := range m {
		if keptIndices[k] {
			keptOldIndices = append(keptOldIndices, k)
		}
	}

	lcsIdx := 0
	// Iterate through the NEW slice. If it matches the current LCS head, we keep it.
	// If it doesn't match, it must be inserted here.
	for k := range n {
		if lcsIdx < len(keptOldIndices) && equal(keptOldIndices[lcsIdx], k) {
			lcsIdx++
		} else {
			ops += fmt.Sprintf("\n\t%s.%s = slices.Insert( %s.%s, %d, %s)", a.GongGetIdentifier(stage), fieldName, a.GongGetIdentifier(stage), fieldName, k, getNewIdentifier(k))
		}
	}

	return ops
}

func __gong__copyBranchCheck[T any](mapOrigCopy map[any]any, from *T) (*T, bool) {
	if to, ok := mapOrigCopy[from]; ok {
		return to.(*T), true
	}
	to := new(T)
	mapOrigCopy[from] = to
	return to, false
}

func __gong__reconstructPointer[T comparable](field *T, refMap map[T]T, instanceField T) {
	var zero T
	if instanceField != zero {
		*field = refMap[instanceField]
	}
}

func __gong__reconstructPointerFromInstance[T comparable](field *T, instMap map[T]T) {
	ref := *field
	var zero T
	if ref != zero {
		*field = zero
		if inst, ok := instMap[ref]; ok {
			*field = inst
		}
	}
}

func __gong__reconstructSliceOfPointersFromReferences[T comparable](field *[]T, refMap map[T]T, instanceSlice []T) {
	*field = (*field)[:0]
	for _, b := range instanceSlice {
		*field = append(*field, refMap[b])
	}
}

func __gong__reconstructSliceOfPointersFromInstances[T comparable](field *[]T, instMap map[T]T) {
	var res []T
	for _, ref := range *field {
		if inst, ok := instMap[ref]; ok {
			res = append(res, inst)
		}
	}
	*field = res
}

func __gong__diffSliceOfPointers[T interface {
	comparable
	GongstructIF
}](
	stage *Stage,
	instance GongstructIF,
	fieldName string,
	oldSlice, newSlice []T,
) string {
	if slices.Equal(oldSlice, newSlice) {
		return ""
	}
	return stage.Diff(
		instance,
		fieldName,
		len(oldSlice),
		len(newSlice),
		func(i, j int) bool {
			return oldSlice[i] == newSlice[j]
		},
		func(j int) string {
			return newSlice[j].GongGetIdentifier(stage)
		},
	)
}
