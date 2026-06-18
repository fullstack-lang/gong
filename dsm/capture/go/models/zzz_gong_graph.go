// generated code - do not edit
package models

import "fmt"

func IsStagedPointerToGongstruct[Type PointerToGongstruct](stage *Stage, instance Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
	case *AnalysisNeed:
		ok = stage.IsStagedAnalysisNeed(target)

	case *Concept:
		ok = stage.IsStagedConcept(target)

	case *ConceptShape:
		ok = stage.IsStagedConceptShape(target)

	case *Concern:
		ok = stage.IsStagedConcern(target)

	case *ConcernCompositionShape:
		ok = stage.IsStagedConcernCompositionShape(target)

	case *ConcernInputShape:
		ok = stage.IsStagedConcernInputShape(target)

	case *ConcernOutputShape:
		ok = stage.IsStagedConcernOutputShape(target)

	case *ConcernShape:
		ok = stage.IsStagedConcernShape(target)

	case *Deliverable:
		ok = stage.IsStagedDeliverable(target)

	case *Diagram:
		ok = stage.IsStagedDiagram(target)

	case *Library:
		ok = stage.IsStagedLibrary(target)

	case *Note:
		ok = stage.IsStagedNote(target)

	case *NoteProductShape:
		ok = stage.IsStagedNoteProductShape(target)

	case *NoteShape:
		ok = stage.IsStagedNoteShape(target)

	case *NoteStakeholderShape:
		ok = stage.IsStagedNoteStakeholderShape(target)

	case *NoteTaskShape:
		ok = stage.IsStagedNoteTaskShape(target)

	case *ProductCompositionShape:
		ok = stage.IsStagedProductCompositionShape(target)

	case *ProductShape:
		ok = stage.IsStagedProductShape(target)

	case *Requirement:
		ok = stage.IsStagedRequirement(target)

	case *RequirementShape:
		ok = stage.IsStagedRequirementShape(target)

	case *Stakeholder:
		ok = stage.IsStagedStakeholder(target)

	case *StakeholderCompositionShape:
		ok = stage.IsStagedStakeholderCompositionShape(target)

	case *StakeholderConcernShape:
		ok = stage.IsStagedStakeholderConcernShape(target)

	case *StakeholderShape:
		ok = stage.IsStagedStakeholderShape(target)

	case *SupportLevel:
		ok = stage.IsStagedSupportLevel(target)

	case *Tool:
		ok = stage.IsStagedTool(target)

	default:
		_ = target
	}
	return
}

func IsStaged[Type Gongstruct](stage *Stage, instance *Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
	case *AnalysisNeed:
		ok = stage.IsStagedAnalysisNeed(target)

	case *Concept:
		ok = stage.IsStagedConcept(target)

	case *ConceptShape:
		ok = stage.IsStagedConceptShape(target)

	case *Concern:
		ok = stage.IsStagedConcern(target)

	case *ConcernCompositionShape:
		ok = stage.IsStagedConcernCompositionShape(target)

	case *ConcernInputShape:
		ok = stage.IsStagedConcernInputShape(target)

	case *ConcernOutputShape:
		ok = stage.IsStagedConcernOutputShape(target)

	case *ConcernShape:
		ok = stage.IsStagedConcernShape(target)

	case *Deliverable:
		ok = stage.IsStagedDeliverable(target)

	case *Diagram:
		ok = stage.IsStagedDiagram(target)

	case *Library:
		ok = stage.IsStagedLibrary(target)

	case *Note:
		ok = stage.IsStagedNote(target)

	case *NoteProductShape:
		ok = stage.IsStagedNoteProductShape(target)

	case *NoteShape:
		ok = stage.IsStagedNoteShape(target)

	case *NoteStakeholderShape:
		ok = stage.IsStagedNoteStakeholderShape(target)

	case *NoteTaskShape:
		ok = stage.IsStagedNoteTaskShape(target)

	case *ProductCompositionShape:
		ok = stage.IsStagedProductCompositionShape(target)

	case *ProductShape:
		ok = stage.IsStagedProductShape(target)

	case *Requirement:
		ok = stage.IsStagedRequirement(target)

	case *RequirementShape:
		ok = stage.IsStagedRequirementShape(target)

	case *Stakeholder:
		ok = stage.IsStagedStakeholder(target)

	case *StakeholderCompositionShape:
		ok = stage.IsStagedStakeholderCompositionShape(target)

	case *StakeholderConcernShape:
		ok = stage.IsStagedStakeholderConcernShape(target)

	case *StakeholderShape:
		ok = stage.IsStagedStakeholderShape(target)

	case *SupportLevel:
		ok = stage.IsStagedSupportLevel(target)

	case *Tool:
		ok = stage.IsStagedTool(target)

	default:
		_ = target
	}
	return
}

// insertion point for stage per struct
func (stage *Stage) IsStagedAnalysisNeed(analysisneed *AnalysisNeed) (ok bool) {

	_, ok = stage.AnalysisNeeds[analysisneed]

	return
}

func (stage *Stage) IsStagedConcept(concept *Concept) (ok bool) {

	_, ok = stage.Concepts[concept]

	return
}

func (stage *Stage) IsStagedConceptShape(conceptshape *ConceptShape) (ok bool) {

	_, ok = stage.ConceptShapes[conceptshape]

	return
}

func (stage *Stage) IsStagedConcern(concern *Concern) (ok bool) {

	_, ok = stage.Concerns[concern]

	return
}

func (stage *Stage) IsStagedConcernCompositionShape(concerncompositionshape *ConcernCompositionShape) (ok bool) {

	_, ok = stage.ConcernCompositionShapes[concerncompositionshape]

	return
}

func (stage *Stage) IsStagedConcernInputShape(concerninputshape *ConcernInputShape) (ok bool) {

	_, ok = stage.ConcernInputShapes[concerninputshape]

	return
}

func (stage *Stage) IsStagedConcernOutputShape(concernoutputshape *ConcernOutputShape) (ok bool) {

	_, ok = stage.ConcernOutputShapes[concernoutputshape]

	return
}

func (stage *Stage) IsStagedConcernShape(concernshape *ConcernShape) (ok bool) {

	_, ok = stage.ConcernShapes[concernshape]

	return
}

func (stage *Stage) IsStagedDeliverable(deliverable *Deliverable) (ok bool) {

	_, ok = stage.Deliverables[deliverable]

	return
}

func (stage *Stage) IsStagedDiagram(diagram *Diagram) (ok bool) {

	_, ok = stage.Diagrams[diagram]

	return
}

func (stage *Stage) IsStagedLibrary(library *Library) (ok bool) {

	_, ok = stage.Librarys[library]

	return
}

func (stage *Stage) IsStagedNote(note *Note) (ok bool) {

	_, ok = stage.Notes[note]

	return
}

func (stage *Stage) IsStagedNoteProductShape(noteproductshape *NoteProductShape) (ok bool) {

	_, ok = stage.NoteProductShapes[noteproductshape]

	return
}

func (stage *Stage) IsStagedNoteShape(noteshape *NoteShape) (ok bool) {

	_, ok = stage.NoteShapes[noteshape]

	return
}

func (stage *Stage) IsStagedNoteStakeholderShape(notestakeholdershape *NoteStakeholderShape) (ok bool) {

	_, ok = stage.NoteStakeholderShapes[notestakeholdershape]

	return
}

func (stage *Stage) IsStagedNoteTaskShape(notetaskshape *NoteTaskShape) (ok bool) {

	_, ok = stage.NoteTaskShapes[notetaskshape]

	return
}

func (stage *Stage) IsStagedProductCompositionShape(productcompositionshape *ProductCompositionShape) (ok bool) {

	_, ok = stage.ProductCompositionShapes[productcompositionshape]

	return
}

func (stage *Stage) IsStagedProductShape(productshape *ProductShape) (ok bool) {

	_, ok = stage.ProductShapes[productshape]

	return
}

func (stage *Stage) IsStagedRequirement(requirement *Requirement) (ok bool) {

	_, ok = stage.Requirements[requirement]

	return
}

func (stage *Stage) IsStagedRequirementShape(requirementshape *RequirementShape) (ok bool) {

	_, ok = stage.RequirementShapes[requirementshape]

	return
}

func (stage *Stage) IsStagedStakeholder(stakeholder *Stakeholder) (ok bool) {

	_, ok = stage.Stakeholders[stakeholder]

	return
}

func (stage *Stage) IsStagedStakeholderCompositionShape(stakeholdercompositionshape *StakeholderCompositionShape) (ok bool) {

	_, ok = stage.StakeholderCompositionShapes[stakeholdercompositionshape]

	return
}

func (stage *Stage) IsStagedStakeholderConcernShape(stakeholderconcernshape *StakeholderConcernShape) (ok bool) {

	_, ok = stage.StakeholderConcernShapes[stakeholderconcernshape]

	return
}

func (stage *Stage) IsStagedStakeholderShape(stakeholdershape *StakeholderShape) (ok bool) {

	_, ok = stage.StakeholderShapes[stakeholdershape]

	return
}

func (stage *Stage) IsStagedSupportLevel(supportlevel *SupportLevel) (ok bool) {

	_, ok = stage.SupportLevels[supportlevel]

	return
}

func (stage *Stage) IsStagedTool(tool *Tool) (ok bool) {

	_, ok = stage.Tools[tool]

	return
}

// StageBranch stages instance and apply StageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func StageBranch[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for stage branch
	case *AnalysisNeed:
		stage.StageBranchAnalysisNeed(target)

	case *Concept:
		stage.StageBranchConcept(target)

	case *ConceptShape:
		stage.StageBranchConceptShape(target)

	case *Concern:
		stage.StageBranchConcern(target)

	case *ConcernCompositionShape:
		stage.StageBranchConcernCompositionShape(target)

	case *ConcernInputShape:
		stage.StageBranchConcernInputShape(target)

	case *ConcernOutputShape:
		stage.StageBranchConcernOutputShape(target)

	case *ConcernShape:
		stage.StageBranchConcernShape(target)

	case *Deliverable:
		stage.StageBranchDeliverable(target)

	case *Diagram:
		stage.StageBranchDiagram(target)

	case *Library:
		stage.StageBranchLibrary(target)

	case *Note:
		stage.StageBranchNote(target)

	case *NoteProductShape:
		stage.StageBranchNoteProductShape(target)

	case *NoteShape:
		stage.StageBranchNoteShape(target)

	case *NoteStakeholderShape:
		stage.StageBranchNoteStakeholderShape(target)

	case *NoteTaskShape:
		stage.StageBranchNoteTaskShape(target)

	case *ProductCompositionShape:
		stage.StageBranchProductCompositionShape(target)

	case *ProductShape:
		stage.StageBranchProductShape(target)

	case *Requirement:
		stage.StageBranchRequirement(target)

	case *RequirementShape:
		stage.StageBranchRequirementShape(target)

	case *Stakeholder:
		stage.StageBranchStakeholder(target)

	case *StakeholderCompositionShape:
		stage.StageBranchStakeholderCompositionShape(target)

	case *StakeholderConcernShape:
		stage.StageBranchStakeholderConcernShape(target)

	case *StakeholderShape:
		stage.StageBranchStakeholderShape(target)

	case *SupportLevel:
		stage.StageBranchSupportLevel(target)

	case *Tool:
		stage.StageBranchTool(target)

	default:
		_ = target
	}
}

// insertion point for stage branch per struct
func (stage *Stage) StageBranchAnalysisNeed(analysisneed *AnalysisNeed) {

	// check if instance is already staged
	if IsStaged(stage, analysisneed) {
		return
	}

	analysisneed.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchConcept(concept *Concept) {

	// check if instance is already staged
	if IsStaged(stage, concept) {
		return
	}

	concept.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _tool := range concept.Tools {
		StageBranch(stage, _tool)
	}

}

func (stage *Stage) StageBranchConceptShape(conceptshape *ConceptShape) {

	// check if instance is already staged
	if IsStaged(stage, conceptshape) {
		return
	}

	conceptshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if conceptshape.Concept != nil {
		StageBranch(stage, conceptshape.Concept)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchConcern(concern *Concern) {

	// check if instance is already staged
	if IsStaged(stage, concern) {
		return
	}

	concern.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _concern := range concern.SubConcerns {
		StageBranch(stage, _concern)
	}
	for _, _deliverable := range concern.Inputs {
		StageBranch(stage, _deliverable)
	}
	for _, _deliverable := range concern.Outputs {
		StageBranch(stage, _deliverable)
	}
	for _, _requirement := range concern.Requirements {
		StageBranch(stage, _requirement)
	}

}

func (stage *Stage) StageBranchConcernCompositionShape(concerncompositionshape *ConcernCompositionShape) {

	// check if instance is already staged
	if IsStaged(stage, concerncompositionshape) {
		return
	}

	concerncompositionshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if concerncompositionshape.Concern != nil {
		StageBranch(stage, concerncompositionshape.Concern)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchConcernInputShape(concerninputshape *ConcernInputShape) {

	// check if instance is already staged
	if IsStaged(stage, concerninputshape) {
		return
	}

	concerninputshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if concerninputshape.Deliverable != nil {
		StageBranch(stage, concerninputshape.Deliverable)
	}
	if concerninputshape.Concern != nil {
		StageBranch(stage, concerninputshape.Concern)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchConcernOutputShape(concernoutputshape *ConcernOutputShape) {

	// check if instance is already staged
	if IsStaged(stage, concernoutputshape) {
		return
	}

	concernoutputshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if concernoutputshape.Task != nil {
		StageBranch(stage, concernoutputshape.Task)
	}
	if concernoutputshape.Product != nil {
		StageBranch(stage, concernoutputshape.Product)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchConcernShape(concernshape *ConcernShape) {

	// check if instance is already staged
	if IsStaged(stage, concernshape) {
		return
	}

	concernshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if concernshape.Concern != nil {
		StageBranch(stage, concernshape.Concern)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchDeliverable(deliverable *Deliverable) {

	// check if instance is already staged
	if IsStaged(stage, deliverable) {
		return
	}

	deliverable.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _deliverable := range deliverable.SubProducts {
		StageBranch(stage, _deliverable)
	}
	for _, _concept := range deliverable.Concepts {
		StageBranch(stage, _concept)
	}

}

func (stage *Stage) StageBranchDiagram(diagram *Diagram) {

	// check if instance is already staged
	if IsStaged(stage, diagram) {
		return
	}

	diagram.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _concern := range diagram.ConcernsWhoseRequirementsNodeIsExpanded {
		StageBranch(stage, _concern)
	}
	for _, _productshape := range diagram.Product_Shapes {
		StageBranch(stage, _productshape)
	}
	for _, _deliverable := range diagram.ProductsWhoseNodeIsExpanded {
		StageBranch(stage, _deliverable)
	}
	for _, _productcompositionshape := range diagram.ProductComposition_Shapes {
		StageBranch(stage, _productcompositionshape)
	}
	for _, _concernshape := range diagram.Concern_Shapes {
		StageBranch(stage, _concernshape)
	}
	for _, _concern := range diagram.ConcernsWhoseNodeIsExpanded {
		StageBranch(stage, _concern)
	}
	for _, _concern := range diagram.ConcernsWhoseInputNodeIsExpanded {
		StageBranch(stage, _concern)
	}
	for _, _concern := range diagram.ConcernsWhoseStakeholderNodeIsExpanded {
		StageBranch(stage, _concern)
	}
	for _, _concern := range diagram.ConcernssWhoseOutputNodeIsExpanded {
		StageBranch(stage, _concern)
	}
	for _, _concerncompositionshape := range diagram.ConcernComposition_Shapes {
		StageBranch(stage, _concerncompositionshape)
	}
	for _, _concerninputshape := range diagram.ConcernInputShapes {
		StageBranch(stage, _concerninputshape)
	}
	for _, _concernoutputshape := range diagram.ConcernOutputShapes {
		StageBranch(stage, _concernoutputshape)
	}
	for _, _noteshape := range diagram.Note_Shapes {
		StageBranch(stage, _noteshape)
	}
	for _, _note := range diagram.NotesWhoseNodeIsExpanded {
		StageBranch(stage, _note)
	}
	for _, _noteproductshape := range diagram.NoteProductShapes {
		StageBranch(stage, _noteproductshape)
	}
	for _, _notetaskshape := range diagram.NoteTaskShapes {
		StageBranch(stage, _notetaskshape)
	}
	for _, _notestakeholdershape := range diagram.NoteResourceShapes {
		StageBranch(stage, _notestakeholdershape)
	}
	for _, _stakeholdershape := range diagram.Stakeholder_Shapes {
		StageBranch(stage, _stakeholdershape)
	}
	for _, _stakeholder := range diagram.ResourcesWhoseNodeIsExpanded {
		StageBranch(stage, _stakeholder)
	}
	for _, _stakeholdercompositionshape := range diagram.ResourceComposition_Shapes {
		StageBranch(stage, _stakeholdercompositionshape)
	}
	for _, _stakeholderconcernshape := range diagram.StakeholderConcernShapes {
		StageBranch(stage, _stakeholderconcernshape)
	}
	for _, _requirementshape := range diagram.Requirement_Shapes {
		StageBranch(stage, _requirementshape)
	}
	for _, _requirement := range diagram.RequirementsWhoseNodeIsExpanded {
		StageBranch(stage, _requirement)
	}
	for _, _conceptshape := range diagram.Concept_Shapes {
		StageBranch(stage, _conceptshape)
	}
	for _, _concept := range diagram.ConceptsWhoseNodeIsExpanded {
		StageBranch(stage, _concept)
	}

}

func (stage *Stage) StageBranchLibrary(library *Library) {

	// check if instance is already staged
	if IsStaged(stage, library) {
		return
	}

	library.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _deliverable := range library.RootDeliverables {
		StageBranch(stage, _deliverable)
	}
	for _, _concern := range library.RootConcerns {
		StageBranch(stage, _concern)
	}
	for _, _stakeholder := range library.RootStakeholders {
		StageBranch(stage, _stakeholder)
	}
	for _, _requirement := range library.RootRequirements {
		StageBranch(stage, _requirement)
	}
	for _, _concept := range library.RootConcepts {
		StageBranch(stage, _concept)
	}
	for _, _analysisneed := range library.AnalysisNeeds {
		StageBranch(stage, _analysisneed)
	}
	for _, _note := range library.Notes {
		StageBranch(stage, _note)
	}
	for _, _diagram := range library.Diagrams {
		StageBranch(stage, _diagram)
	}
	for _, _library := range library.SubLibraries {
		StageBranch(stage, _library)
	}

}

func (stage *Stage) StageBranchNote(note *Note) {

	// check if instance is already staged
	if IsStaged(stage, note) {
		return
	}

	note.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _deliverable := range note.Products {
		StageBranch(stage, _deliverable)
	}
	for _, _concern := range note.Tasks {
		StageBranch(stage, _concern)
	}
	for _, _stakeholder := range note.Resources {
		StageBranch(stage, _stakeholder)
	}

}

func (stage *Stage) StageBranchNoteProductShape(noteproductshape *NoteProductShape) {

	// check if instance is already staged
	if IsStaged(stage, noteproductshape) {
		return
	}

	noteproductshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if noteproductshape.Note != nil {
		StageBranch(stage, noteproductshape.Note)
	}
	if noteproductshape.Product != nil {
		StageBranch(stage, noteproductshape.Product)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchNoteShape(noteshape *NoteShape) {

	// check if instance is already staged
	if IsStaged(stage, noteshape) {
		return
	}

	noteshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if noteshape.Note != nil {
		StageBranch(stage, noteshape.Note)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchNoteStakeholderShape(notestakeholdershape *NoteStakeholderShape) {

	// check if instance is already staged
	if IsStaged(stage, notestakeholdershape) {
		return
	}

	notestakeholdershape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if notestakeholdershape.Note != nil {
		StageBranch(stage, notestakeholdershape.Note)
	}
	if notestakeholdershape.Stakeholder != nil {
		StageBranch(stage, notestakeholdershape.Stakeholder)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchNoteTaskShape(notetaskshape *NoteTaskShape) {

	// check if instance is already staged
	if IsStaged(stage, notetaskshape) {
		return
	}

	notetaskshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if notetaskshape.Note != nil {
		StageBranch(stage, notetaskshape.Note)
	}
	if notetaskshape.Task != nil {
		StageBranch(stage, notetaskshape.Task)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchProductCompositionShape(productcompositionshape *ProductCompositionShape) {

	// check if instance is already staged
	if IsStaged(stage, productcompositionshape) {
		return
	}

	productcompositionshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if productcompositionshape.Product != nil {
		StageBranch(stage, productcompositionshape.Product)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchProductShape(productshape *ProductShape) {

	// check if instance is already staged
	if IsStaged(stage, productshape) {
		return
	}

	productshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if productshape.Product != nil {
		StageBranch(stage, productshape.Product)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchRequirement(requirement *Requirement) {

	// check if instance is already staged
	if IsStaged(stage, requirement) {
		return
	}

	requirement.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _supportlevel := range requirement.SupportLevels {
		StageBranch(stage, _supportlevel)
	}
	for _, _concept := range requirement.Concepts {
		StageBranch(stage, _concept)
	}

}

func (stage *Stage) StageBranchRequirementShape(requirementshape *RequirementShape) {

	// check if instance is already staged
	if IsStaged(stage, requirementshape) {
		return
	}

	requirementshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if requirementshape.Requirement != nil {
		StageBranch(stage, requirementshape.Requirement)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchStakeholder(stakeholder *Stakeholder) {

	// check if instance is already staged
	if IsStaged(stage, stakeholder) {
		return
	}

	stakeholder.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _concern := range stakeholder.Concerns {
		StageBranch(stage, _concern)
	}
	for _, _stakeholder := range stakeholder.SubStakeholders {
		StageBranch(stage, _stakeholder)
	}

}

func (stage *Stage) StageBranchStakeholderCompositionShape(stakeholdercompositionshape *StakeholderCompositionShape) {

	// check if instance is already staged
	if IsStaged(stage, stakeholdercompositionshape) {
		return
	}

	stakeholdercompositionshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if stakeholdercompositionshape.Stakeholder != nil {
		StageBranch(stage, stakeholdercompositionshape.Stakeholder)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchStakeholderConcernShape(stakeholderconcernshape *StakeholderConcernShape) {

	// check if instance is already staged
	if IsStaged(stage, stakeholderconcernshape) {
		return
	}

	stakeholderconcernshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if stakeholderconcernshape.Stakeholder != nil {
		StageBranch(stage, stakeholderconcernshape.Stakeholder)
	}
	if stakeholderconcernshape.Concern != nil {
		StageBranch(stage, stakeholderconcernshape.Concern)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchStakeholderShape(stakeholdershape *StakeholderShape) {

	// check if instance is already staged
	if IsStaged(stage, stakeholdershape) {
		return
	}

	stakeholdershape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if stakeholdershape.Stakeholder != nil {
		StageBranch(stage, stakeholdershape.Stakeholder)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchSupportLevel(supportlevel *SupportLevel) {

	// check if instance is already staged
	if IsStaged(stage, supportlevel) {
		return
	}

	supportlevel.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if supportlevel.Tool != nil {
		StageBranch(stage, supportlevel.Tool)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchTool(tool *Tool) {

	// check if instance is already staged
	if IsStaged(stage, tool) {
		return
	}

	tool.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

// CopyBranch stages instance and apply CopyBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func CopyBranch[Type Gongstruct](from *Type) (to *Type) {

	mapOrigCopy := make(map[any]any)
	_ = mapOrigCopy

	switch fromT := any(from).(type) {
	// insertion point for stage branch
	case *AnalysisNeed:
		toT := CopyBranchAnalysisNeed(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Concept:
		toT := CopyBranchConcept(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ConceptShape:
		toT := CopyBranchConceptShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Concern:
		toT := CopyBranchConcern(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ConcernCompositionShape:
		toT := CopyBranchConcernCompositionShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ConcernInputShape:
		toT := CopyBranchConcernInputShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ConcernOutputShape:
		toT := CopyBranchConcernOutputShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ConcernShape:
		toT := CopyBranchConcernShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Deliverable:
		toT := CopyBranchDeliverable(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Diagram:
		toT := CopyBranchDiagram(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Library:
		toT := CopyBranchLibrary(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Note:
		toT := CopyBranchNote(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *NoteProductShape:
		toT := CopyBranchNoteProductShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *NoteShape:
		toT := CopyBranchNoteShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *NoteStakeholderShape:
		toT := CopyBranchNoteStakeholderShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *NoteTaskShape:
		toT := CopyBranchNoteTaskShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ProductCompositionShape:
		toT := CopyBranchProductCompositionShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ProductShape:
		toT := CopyBranchProductShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Requirement:
		toT := CopyBranchRequirement(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *RequirementShape:
		toT := CopyBranchRequirementShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Stakeholder:
		toT := CopyBranchStakeholder(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *StakeholderCompositionShape:
		toT := CopyBranchStakeholderCompositionShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *StakeholderConcernShape:
		toT := CopyBranchStakeholderConcernShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *StakeholderShape:
		toT := CopyBranchStakeholderShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *SupportLevel:
		toT := CopyBranchSupportLevel(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Tool:
		toT := CopyBranchTool(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
func CopyBranchAnalysisNeed(mapOrigCopy map[any]any, analysisneedFrom *AnalysisNeed) (analysisneedTo *AnalysisNeed) {

	// analysisneedFrom has already been copied
	if _analysisneedTo, ok := mapOrigCopy[analysisneedFrom]; ok {
		analysisneedTo = _analysisneedTo.(*AnalysisNeed)
		return
	}

	analysisneedTo = new(AnalysisNeed)
	mapOrigCopy[analysisneedFrom] = analysisneedTo
	analysisneedFrom.CopyBasicFields(analysisneedTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchConcept(mapOrigCopy map[any]any, conceptFrom *Concept) (conceptTo *Concept) {

	// conceptFrom has already been copied
	if _conceptTo, ok := mapOrigCopy[conceptFrom]; ok {
		conceptTo = _conceptTo.(*Concept)
		return
	}

	conceptTo = new(Concept)
	mapOrigCopy[conceptFrom] = conceptTo
	conceptFrom.CopyBasicFields(conceptTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _tool := range conceptFrom.Tools {
		conceptTo.Tools = append(conceptTo.Tools, CopyBranchTool(mapOrigCopy, _tool))
	}

	return
}

func CopyBranchConceptShape(mapOrigCopy map[any]any, conceptshapeFrom *ConceptShape) (conceptshapeTo *ConceptShape) {

	// conceptshapeFrom has already been copied
	if _conceptshapeTo, ok := mapOrigCopy[conceptshapeFrom]; ok {
		conceptshapeTo = _conceptshapeTo.(*ConceptShape)
		return
	}

	conceptshapeTo = new(ConceptShape)
	mapOrigCopy[conceptshapeFrom] = conceptshapeTo
	conceptshapeFrom.CopyBasicFields(conceptshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if conceptshapeFrom.Concept != nil {
		conceptshapeTo.Concept = CopyBranchConcept(mapOrigCopy, conceptshapeFrom.Concept)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchConcern(mapOrigCopy map[any]any, concernFrom *Concern) (concernTo *Concern) {

	// concernFrom has already been copied
	if _concernTo, ok := mapOrigCopy[concernFrom]; ok {
		concernTo = _concernTo.(*Concern)
		return
	}

	concernTo = new(Concern)
	mapOrigCopy[concernFrom] = concernTo
	concernFrom.CopyBasicFields(concernTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _concern := range concernFrom.SubConcerns {
		concernTo.SubConcerns = append(concernTo.SubConcerns, CopyBranchConcern(mapOrigCopy, _concern))
	}
	for _, _deliverable := range concernFrom.Inputs {
		concernTo.Inputs = append(concernTo.Inputs, CopyBranchDeliverable(mapOrigCopy, _deliverable))
	}
	for _, _deliverable := range concernFrom.Outputs {
		concernTo.Outputs = append(concernTo.Outputs, CopyBranchDeliverable(mapOrigCopy, _deliverable))
	}
	for _, _requirement := range concernFrom.Requirements {
		concernTo.Requirements = append(concernTo.Requirements, CopyBranchRequirement(mapOrigCopy, _requirement))
	}

	return
}

func CopyBranchConcernCompositionShape(mapOrigCopy map[any]any, concerncompositionshapeFrom *ConcernCompositionShape) (concerncompositionshapeTo *ConcernCompositionShape) {

	// concerncompositionshapeFrom has already been copied
	if _concerncompositionshapeTo, ok := mapOrigCopy[concerncompositionshapeFrom]; ok {
		concerncompositionshapeTo = _concerncompositionshapeTo.(*ConcernCompositionShape)
		return
	}

	concerncompositionshapeTo = new(ConcernCompositionShape)
	mapOrigCopy[concerncompositionshapeFrom] = concerncompositionshapeTo
	concerncompositionshapeFrom.CopyBasicFields(concerncompositionshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if concerncompositionshapeFrom.Concern != nil {
		concerncompositionshapeTo.Concern = CopyBranchConcern(mapOrigCopy, concerncompositionshapeFrom.Concern)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchConcernInputShape(mapOrigCopy map[any]any, concerninputshapeFrom *ConcernInputShape) (concerninputshapeTo *ConcernInputShape) {

	// concerninputshapeFrom has already been copied
	if _concerninputshapeTo, ok := mapOrigCopy[concerninputshapeFrom]; ok {
		concerninputshapeTo = _concerninputshapeTo.(*ConcernInputShape)
		return
	}

	concerninputshapeTo = new(ConcernInputShape)
	mapOrigCopy[concerninputshapeFrom] = concerninputshapeTo
	concerninputshapeFrom.CopyBasicFields(concerninputshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if concerninputshapeFrom.Deliverable != nil {
		concerninputshapeTo.Deliverable = CopyBranchDeliverable(mapOrigCopy, concerninputshapeFrom.Deliverable)
	}
	if concerninputshapeFrom.Concern != nil {
		concerninputshapeTo.Concern = CopyBranchConcern(mapOrigCopy, concerninputshapeFrom.Concern)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchConcernOutputShape(mapOrigCopy map[any]any, concernoutputshapeFrom *ConcernOutputShape) (concernoutputshapeTo *ConcernOutputShape) {

	// concernoutputshapeFrom has already been copied
	if _concernoutputshapeTo, ok := mapOrigCopy[concernoutputshapeFrom]; ok {
		concernoutputshapeTo = _concernoutputshapeTo.(*ConcernOutputShape)
		return
	}

	concernoutputshapeTo = new(ConcernOutputShape)
	mapOrigCopy[concernoutputshapeFrom] = concernoutputshapeTo
	concernoutputshapeFrom.CopyBasicFields(concernoutputshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if concernoutputshapeFrom.Task != nil {
		concernoutputshapeTo.Task = CopyBranchConcern(mapOrigCopy, concernoutputshapeFrom.Task)
	}
	if concernoutputshapeFrom.Product != nil {
		concernoutputshapeTo.Product = CopyBranchDeliverable(mapOrigCopy, concernoutputshapeFrom.Product)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchConcernShape(mapOrigCopy map[any]any, concernshapeFrom *ConcernShape) (concernshapeTo *ConcernShape) {

	// concernshapeFrom has already been copied
	if _concernshapeTo, ok := mapOrigCopy[concernshapeFrom]; ok {
		concernshapeTo = _concernshapeTo.(*ConcernShape)
		return
	}

	concernshapeTo = new(ConcernShape)
	mapOrigCopy[concernshapeFrom] = concernshapeTo
	concernshapeFrom.CopyBasicFields(concernshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if concernshapeFrom.Concern != nil {
		concernshapeTo.Concern = CopyBranchConcern(mapOrigCopy, concernshapeFrom.Concern)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchDeliverable(mapOrigCopy map[any]any, deliverableFrom *Deliverable) (deliverableTo *Deliverable) {

	// deliverableFrom has already been copied
	if _deliverableTo, ok := mapOrigCopy[deliverableFrom]; ok {
		deliverableTo = _deliverableTo.(*Deliverable)
		return
	}

	deliverableTo = new(Deliverable)
	mapOrigCopy[deliverableFrom] = deliverableTo
	deliverableFrom.CopyBasicFields(deliverableTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _deliverable := range deliverableFrom.SubProducts {
		deliverableTo.SubProducts = append(deliverableTo.SubProducts, CopyBranchDeliverable(mapOrigCopy, _deliverable))
	}
	for _, _concept := range deliverableFrom.Concepts {
		deliverableTo.Concepts = append(deliverableTo.Concepts, CopyBranchConcept(mapOrigCopy, _concept))
	}

	return
}

func CopyBranchDiagram(mapOrigCopy map[any]any, diagramFrom *Diagram) (diagramTo *Diagram) {

	// diagramFrom has already been copied
	if _diagramTo, ok := mapOrigCopy[diagramFrom]; ok {
		diagramTo = _diagramTo.(*Diagram)
		return
	}

	diagramTo = new(Diagram)
	mapOrigCopy[diagramFrom] = diagramTo
	diagramFrom.CopyBasicFields(diagramTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _concern := range diagramFrom.ConcernsWhoseRequirementsNodeIsExpanded {
		diagramTo.ConcernsWhoseRequirementsNodeIsExpanded = append(diagramTo.ConcernsWhoseRequirementsNodeIsExpanded, CopyBranchConcern(mapOrigCopy, _concern))
	}
	for _, _productshape := range diagramFrom.Product_Shapes {
		diagramTo.Product_Shapes = append(diagramTo.Product_Shapes, CopyBranchProductShape(mapOrigCopy, _productshape))
	}
	for _, _deliverable := range diagramFrom.ProductsWhoseNodeIsExpanded {
		diagramTo.ProductsWhoseNodeIsExpanded = append(diagramTo.ProductsWhoseNodeIsExpanded, CopyBranchDeliverable(mapOrigCopy, _deliverable))
	}
	for _, _productcompositionshape := range diagramFrom.ProductComposition_Shapes {
		diagramTo.ProductComposition_Shapes = append(diagramTo.ProductComposition_Shapes, CopyBranchProductCompositionShape(mapOrigCopy, _productcompositionshape))
	}
	for _, _concernshape := range diagramFrom.Concern_Shapes {
		diagramTo.Concern_Shapes = append(diagramTo.Concern_Shapes, CopyBranchConcernShape(mapOrigCopy, _concernshape))
	}
	for _, _concern := range diagramFrom.ConcernsWhoseNodeIsExpanded {
		diagramTo.ConcernsWhoseNodeIsExpanded = append(diagramTo.ConcernsWhoseNodeIsExpanded, CopyBranchConcern(mapOrigCopy, _concern))
	}
	for _, _concern := range diagramFrom.ConcernsWhoseInputNodeIsExpanded {
		diagramTo.ConcernsWhoseInputNodeIsExpanded = append(diagramTo.ConcernsWhoseInputNodeIsExpanded, CopyBranchConcern(mapOrigCopy, _concern))
	}
	for _, _concern := range diagramFrom.ConcernsWhoseStakeholderNodeIsExpanded {
		diagramTo.ConcernsWhoseStakeholderNodeIsExpanded = append(diagramTo.ConcernsWhoseStakeholderNodeIsExpanded, CopyBranchConcern(mapOrigCopy, _concern))
	}
	for _, _concern := range diagramFrom.ConcernssWhoseOutputNodeIsExpanded {
		diagramTo.ConcernssWhoseOutputNodeIsExpanded = append(diagramTo.ConcernssWhoseOutputNodeIsExpanded, CopyBranchConcern(mapOrigCopy, _concern))
	}
	for _, _concerncompositionshape := range diagramFrom.ConcernComposition_Shapes {
		diagramTo.ConcernComposition_Shapes = append(diagramTo.ConcernComposition_Shapes, CopyBranchConcernCompositionShape(mapOrigCopy, _concerncompositionshape))
	}
	for _, _concerninputshape := range diagramFrom.ConcernInputShapes {
		diagramTo.ConcernInputShapes = append(diagramTo.ConcernInputShapes, CopyBranchConcernInputShape(mapOrigCopy, _concerninputshape))
	}
	for _, _concernoutputshape := range diagramFrom.ConcernOutputShapes {
		diagramTo.ConcernOutputShapes = append(diagramTo.ConcernOutputShapes, CopyBranchConcernOutputShape(mapOrigCopy, _concernoutputshape))
	}
	for _, _noteshape := range diagramFrom.Note_Shapes {
		diagramTo.Note_Shapes = append(diagramTo.Note_Shapes, CopyBranchNoteShape(mapOrigCopy, _noteshape))
	}
	for _, _note := range diagramFrom.NotesWhoseNodeIsExpanded {
		diagramTo.NotesWhoseNodeIsExpanded = append(diagramTo.NotesWhoseNodeIsExpanded, CopyBranchNote(mapOrigCopy, _note))
	}
	for _, _noteproductshape := range diagramFrom.NoteProductShapes {
		diagramTo.NoteProductShapes = append(diagramTo.NoteProductShapes, CopyBranchNoteProductShape(mapOrigCopy, _noteproductshape))
	}
	for _, _notetaskshape := range diagramFrom.NoteTaskShapes {
		diagramTo.NoteTaskShapes = append(diagramTo.NoteTaskShapes, CopyBranchNoteTaskShape(mapOrigCopy, _notetaskshape))
	}
	for _, _notestakeholdershape := range diagramFrom.NoteResourceShapes {
		diagramTo.NoteResourceShapes = append(diagramTo.NoteResourceShapes, CopyBranchNoteStakeholderShape(mapOrigCopy, _notestakeholdershape))
	}
	for _, _stakeholdershape := range diagramFrom.Stakeholder_Shapes {
		diagramTo.Stakeholder_Shapes = append(diagramTo.Stakeholder_Shapes, CopyBranchStakeholderShape(mapOrigCopy, _stakeholdershape))
	}
	for _, _stakeholder := range diagramFrom.ResourcesWhoseNodeIsExpanded {
		diagramTo.ResourcesWhoseNodeIsExpanded = append(diagramTo.ResourcesWhoseNodeIsExpanded, CopyBranchStakeholder(mapOrigCopy, _stakeholder))
	}
	for _, _stakeholdercompositionshape := range diagramFrom.ResourceComposition_Shapes {
		diagramTo.ResourceComposition_Shapes = append(diagramTo.ResourceComposition_Shapes, CopyBranchStakeholderCompositionShape(mapOrigCopy, _stakeholdercompositionshape))
	}
	for _, _stakeholderconcernshape := range diagramFrom.StakeholderConcernShapes {
		diagramTo.StakeholderConcernShapes = append(diagramTo.StakeholderConcernShapes, CopyBranchStakeholderConcernShape(mapOrigCopy, _stakeholderconcernshape))
	}
	for _, _requirementshape := range diagramFrom.Requirement_Shapes {
		diagramTo.Requirement_Shapes = append(diagramTo.Requirement_Shapes, CopyBranchRequirementShape(mapOrigCopy, _requirementshape))
	}
	for _, _requirement := range diagramFrom.RequirementsWhoseNodeIsExpanded {
		diagramTo.RequirementsWhoseNodeIsExpanded = append(diagramTo.RequirementsWhoseNodeIsExpanded, CopyBranchRequirement(mapOrigCopy, _requirement))
	}
	for _, _conceptshape := range diagramFrom.Concept_Shapes {
		diagramTo.Concept_Shapes = append(diagramTo.Concept_Shapes, CopyBranchConceptShape(mapOrigCopy, _conceptshape))
	}
	for _, _concept := range diagramFrom.ConceptsWhoseNodeIsExpanded {
		diagramTo.ConceptsWhoseNodeIsExpanded = append(diagramTo.ConceptsWhoseNodeIsExpanded, CopyBranchConcept(mapOrigCopy, _concept))
	}

	return
}

func CopyBranchLibrary(mapOrigCopy map[any]any, libraryFrom *Library) (libraryTo *Library) {

	// libraryFrom has already been copied
	if _libraryTo, ok := mapOrigCopy[libraryFrom]; ok {
		libraryTo = _libraryTo.(*Library)
		return
	}

	libraryTo = new(Library)
	mapOrigCopy[libraryFrom] = libraryTo
	libraryFrom.CopyBasicFields(libraryTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _deliverable := range libraryFrom.RootDeliverables {
		libraryTo.RootDeliverables = append(libraryTo.RootDeliverables, CopyBranchDeliverable(mapOrigCopy, _deliverable))
	}
	for _, _concern := range libraryFrom.RootConcerns {
		libraryTo.RootConcerns = append(libraryTo.RootConcerns, CopyBranchConcern(mapOrigCopy, _concern))
	}
	for _, _stakeholder := range libraryFrom.RootStakeholders {
		libraryTo.RootStakeholders = append(libraryTo.RootStakeholders, CopyBranchStakeholder(mapOrigCopy, _stakeholder))
	}
	for _, _requirement := range libraryFrom.RootRequirements {
		libraryTo.RootRequirements = append(libraryTo.RootRequirements, CopyBranchRequirement(mapOrigCopy, _requirement))
	}
	for _, _concept := range libraryFrom.RootConcepts {
		libraryTo.RootConcepts = append(libraryTo.RootConcepts, CopyBranchConcept(mapOrigCopy, _concept))
	}
	for _, _analysisneed := range libraryFrom.AnalysisNeeds {
		libraryTo.AnalysisNeeds = append(libraryTo.AnalysisNeeds, CopyBranchAnalysisNeed(mapOrigCopy, _analysisneed))
	}
	for _, _note := range libraryFrom.Notes {
		libraryTo.Notes = append(libraryTo.Notes, CopyBranchNote(mapOrigCopy, _note))
	}
	for _, _diagram := range libraryFrom.Diagrams {
		libraryTo.Diagrams = append(libraryTo.Diagrams, CopyBranchDiagram(mapOrigCopy, _diagram))
	}
	for _, _library := range libraryFrom.SubLibraries {
		libraryTo.SubLibraries = append(libraryTo.SubLibraries, CopyBranchLibrary(mapOrigCopy, _library))
	}

	return
}

func CopyBranchNote(mapOrigCopy map[any]any, noteFrom *Note) (noteTo *Note) {

	// noteFrom has already been copied
	if _noteTo, ok := mapOrigCopy[noteFrom]; ok {
		noteTo = _noteTo.(*Note)
		return
	}

	noteTo = new(Note)
	mapOrigCopy[noteFrom] = noteTo
	noteFrom.CopyBasicFields(noteTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _deliverable := range noteFrom.Products {
		noteTo.Products = append(noteTo.Products, CopyBranchDeliverable(mapOrigCopy, _deliverable))
	}
	for _, _concern := range noteFrom.Tasks {
		noteTo.Tasks = append(noteTo.Tasks, CopyBranchConcern(mapOrigCopy, _concern))
	}
	for _, _stakeholder := range noteFrom.Resources {
		noteTo.Resources = append(noteTo.Resources, CopyBranchStakeholder(mapOrigCopy, _stakeholder))
	}

	return
}

func CopyBranchNoteProductShape(mapOrigCopy map[any]any, noteproductshapeFrom *NoteProductShape) (noteproductshapeTo *NoteProductShape) {

	// noteproductshapeFrom has already been copied
	if _noteproductshapeTo, ok := mapOrigCopy[noteproductshapeFrom]; ok {
		noteproductshapeTo = _noteproductshapeTo.(*NoteProductShape)
		return
	}

	noteproductshapeTo = new(NoteProductShape)
	mapOrigCopy[noteproductshapeFrom] = noteproductshapeTo
	noteproductshapeFrom.CopyBasicFields(noteproductshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if noteproductshapeFrom.Note != nil {
		noteproductshapeTo.Note = CopyBranchNote(mapOrigCopy, noteproductshapeFrom.Note)
	}
	if noteproductshapeFrom.Product != nil {
		noteproductshapeTo.Product = CopyBranchDeliverable(mapOrigCopy, noteproductshapeFrom.Product)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchNoteShape(mapOrigCopy map[any]any, noteshapeFrom *NoteShape) (noteshapeTo *NoteShape) {

	// noteshapeFrom has already been copied
	if _noteshapeTo, ok := mapOrigCopy[noteshapeFrom]; ok {
		noteshapeTo = _noteshapeTo.(*NoteShape)
		return
	}

	noteshapeTo = new(NoteShape)
	mapOrigCopy[noteshapeFrom] = noteshapeTo
	noteshapeFrom.CopyBasicFields(noteshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if noteshapeFrom.Note != nil {
		noteshapeTo.Note = CopyBranchNote(mapOrigCopy, noteshapeFrom.Note)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchNoteStakeholderShape(mapOrigCopy map[any]any, notestakeholdershapeFrom *NoteStakeholderShape) (notestakeholdershapeTo *NoteStakeholderShape) {

	// notestakeholdershapeFrom has already been copied
	if _notestakeholdershapeTo, ok := mapOrigCopy[notestakeholdershapeFrom]; ok {
		notestakeholdershapeTo = _notestakeholdershapeTo.(*NoteStakeholderShape)
		return
	}

	notestakeholdershapeTo = new(NoteStakeholderShape)
	mapOrigCopy[notestakeholdershapeFrom] = notestakeholdershapeTo
	notestakeholdershapeFrom.CopyBasicFields(notestakeholdershapeTo)

	//insertion point for the staging of instances referenced by pointers
	if notestakeholdershapeFrom.Note != nil {
		notestakeholdershapeTo.Note = CopyBranchNote(mapOrigCopy, notestakeholdershapeFrom.Note)
	}
	if notestakeholdershapeFrom.Stakeholder != nil {
		notestakeholdershapeTo.Stakeholder = CopyBranchStakeholder(mapOrigCopy, notestakeholdershapeFrom.Stakeholder)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchNoteTaskShape(mapOrigCopy map[any]any, notetaskshapeFrom *NoteTaskShape) (notetaskshapeTo *NoteTaskShape) {

	// notetaskshapeFrom has already been copied
	if _notetaskshapeTo, ok := mapOrigCopy[notetaskshapeFrom]; ok {
		notetaskshapeTo = _notetaskshapeTo.(*NoteTaskShape)
		return
	}

	notetaskshapeTo = new(NoteTaskShape)
	mapOrigCopy[notetaskshapeFrom] = notetaskshapeTo
	notetaskshapeFrom.CopyBasicFields(notetaskshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if notetaskshapeFrom.Note != nil {
		notetaskshapeTo.Note = CopyBranchNote(mapOrigCopy, notetaskshapeFrom.Note)
	}
	if notetaskshapeFrom.Task != nil {
		notetaskshapeTo.Task = CopyBranchConcern(mapOrigCopy, notetaskshapeFrom.Task)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchProductCompositionShape(mapOrigCopy map[any]any, productcompositionshapeFrom *ProductCompositionShape) (productcompositionshapeTo *ProductCompositionShape) {

	// productcompositionshapeFrom has already been copied
	if _productcompositionshapeTo, ok := mapOrigCopy[productcompositionshapeFrom]; ok {
		productcompositionshapeTo = _productcompositionshapeTo.(*ProductCompositionShape)
		return
	}

	productcompositionshapeTo = new(ProductCompositionShape)
	mapOrigCopy[productcompositionshapeFrom] = productcompositionshapeTo
	productcompositionshapeFrom.CopyBasicFields(productcompositionshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if productcompositionshapeFrom.Product != nil {
		productcompositionshapeTo.Product = CopyBranchDeliverable(mapOrigCopy, productcompositionshapeFrom.Product)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchProductShape(mapOrigCopy map[any]any, productshapeFrom *ProductShape) (productshapeTo *ProductShape) {

	// productshapeFrom has already been copied
	if _productshapeTo, ok := mapOrigCopy[productshapeFrom]; ok {
		productshapeTo = _productshapeTo.(*ProductShape)
		return
	}

	productshapeTo = new(ProductShape)
	mapOrigCopy[productshapeFrom] = productshapeTo
	productshapeFrom.CopyBasicFields(productshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if productshapeFrom.Product != nil {
		productshapeTo.Product = CopyBranchDeliverable(mapOrigCopy, productshapeFrom.Product)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchRequirement(mapOrigCopy map[any]any, requirementFrom *Requirement) (requirementTo *Requirement) {

	// requirementFrom has already been copied
	if _requirementTo, ok := mapOrigCopy[requirementFrom]; ok {
		requirementTo = _requirementTo.(*Requirement)
		return
	}

	requirementTo = new(Requirement)
	mapOrigCopy[requirementFrom] = requirementTo
	requirementFrom.CopyBasicFields(requirementTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _supportlevel := range requirementFrom.SupportLevels {
		requirementTo.SupportLevels = append(requirementTo.SupportLevels, CopyBranchSupportLevel(mapOrigCopy, _supportlevel))
	}
	for _, _concept := range requirementFrom.Concepts {
		requirementTo.Concepts = append(requirementTo.Concepts, CopyBranchConcept(mapOrigCopy, _concept))
	}

	return
}

func CopyBranchRequirementShape(mapOrigCopy map[any]any, requirementshapeFrom *RequirementShape) (requirementshapeTo *RequirementShape) {

	// requirementshapeFrom has already been copied
	if _requirementshapeTo, ok := mapOrigCopy[requirementshapeFrom]; ok {
		requirementshapeTo = _requirementshapeTo.(*RequirementShape)
		return
	}

	requirementshapeTo = new(RequirementShape)
	mapOrigCopy[requirementshapeFrom] = requirementshapeTo
	requirementshapeFrom.CopyBasicFields(requirementshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if requirementshapeFrom.Requirement != nil {
		requirementshapeTo.Requirement = CopyBranchRequirement(mapOrigCopy, requirementshapeFrom.Requirement)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchStakeholder(mapOrigCopy map[any]any, stakeholderFrom *Stakeholder) (stakeholderTo *Stakeholder) {

	// stakeholderFrom has already been copied
	if _stakeholderTo, ok := mapOrigCopy[stakeholderFrom]; ok {
		stakeholderTo = _stakeholderTo.(*Stakeholder)
		return
	}

	stakeholderTo = new(Stakeholder)
	mapOrigCopy[stakeholderFrom] = stakeholderTo
	stakeholderFrom.CopyBasicFields(stakeholderTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _concern := range stakeholderFrom.Concerns {
		stakeholderTo.Concerns = append(stakeholderTo.Concerns, CopyBranchConcern(mapOrigCopy, _concern))
	}
	for _, _stakeholder := range stakeholderFrom.SubStakeholders {
		stakeholderTo.SubStakeholders = append(stakeholderTo.SubStakeholders, CopyBranchStakeholder(mapOrigCopy, _stakeholder))
	}

	return
}

func CopyBranchStakeholderCompositionShape(mapOrigCopy map[any]any, stakeholdercompositionshapeFrom *StakeholderCompositionShape) (stakeholdercompositionshapeTo *StakeholderCompositionShape) {

	// stakeholdercompositionshapeFrom has already been copied
	if _stakeholdercompositionshapeTo, ok := mapOrigCopy[stakeholdercompositionshapeFrom]; ok {
		stakeholdercompositionshapeTo = _stakeholdercompositionshapeTo.(*StakeholderCompositionShape)
		return
	}

	stakeholdercompositionshapeTo = new(StakeholderCompositionShape)
	mapOrigCopy[stakeholdercompositionshapeFrom] = stakeholdercompositionshapeTo
	stakeholdercompositionshapeFrom.CopyBasicFields(stakeholdercompositionshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if stakeholdercompositionshapeFrom.Stakeholder != nil {
		stakeholdercompositionshapeTo.Stakeholder = CopyBranchStakeholder(mapOrigCopy, stakeholdercompositionshapeFrom.Stakeholder)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchStakeholderConcernShape(mapOrigCopy map[any]any, stakeholderconcernshapeFrom *StakeholderConcernShape) (stakeholderconcernshapeTo *StakeholderConcernShape) {

	// stakeholderconcernshapeFrom has already been copied
	if _stakeholderconcernshapeTo, ok := mapOrigCopy[stakeholderconcernshapeFrom]; ok {
		stakeholderconcernshapeTo = _stakeholderconcernshapeTo.(*StakeholderConcernShape)
		return
	}

	stakeholderconcernshapeTo = new(StakeholderConcernShape)
	mapOrigCopy[stakeholderconcernshapeFrom] = stakeholderconcernshapeTo
	stakeholderconcernshapeFrom.CopyBasicFields(stakeholderconcernshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if stakeholderconcernshapeFrom.Stakeholder != nil {
		stakeholderconcernshapeTo.Stakeholder = CopyBranchStakeholder(mapOrigCopy, stakeholderconcernshapeFrom.Stakeholder)
	}
	if stakeholderconcernshapeFrom.Concern != nil {
		stakeholderconcernshapeTo.Concern = CopyBranchConcern(mapOrigCopy, stakeholderconcernshapeFrom.Concern)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchStakeholderShape(mapOrigCopy map[any]any, stakeholdershapeFrom *StakeholderShape) (stakeholdershapeTo *StakeholderShape) {

	// stakeholdershapeFrom has already been copied
	if _stakeholdershapeTo, ok := mapOrigCopy[stakeholdershapeFrom]; ok {
		stakeholdershapeTo = _stakeholdershapeTo.(*StakeholderShape)
		return
	}

	stakeholdershapeTo = new(StakeholderShape)
	mapOrigCopy[stakeholdershapeFrom] = stakeholdershapeTo
	stakeholdershapeFrom.CopyBasicFields(stakeholdershapeTo)

	//insertion point for the staging of instances referenced by pointers
	if stakeholdershapeFrom.Stakeholder != nil {
		stakeholdershapeTo.Stakeholder = CopyBranchStakeholder(mapOrigCopy, stakeholdershapeFrom.Stakeholder)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchSupportLevel(mapOrigCopy map[any]any, supportlevelFrom *SupportLevel) (supportlevelTo *SupportLevel) {

	// supportlevelFrom has already been copied
	if _supportlevelTo, ok := mapOrigCopy[supportlevelFrom]; ok {
		supportlevelTo = _supportlevelTo.(*SupportLevel)
		return
	}

	supportlevelTo = new(SupportLevel)
	mapOrigCopy[supportlevelFrom] = supportlevelTo
	supportlevelFrom.CopyBasicFields(supportlevelTo)

	//insertion point for the staging of instances referenced by pointers
	if supportlevelFrom.Tool != nil {
		supportlevelTo.Tool = CopyBranchTool(mapOrigCopy, supportlevelFrom.Tool)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchTool(mapOrigCopy map[any]any, toolFrom *Tool) (toolTo *Tool) {

	// toolFrom has already been copied
	if _toolTo, ok := mapOrigCopy[toolFrom]; ok {
		toolTo = _toolTo.(*Tool)
		return
	}

	toolTo = new(Tool)
	mapOrigCopy[toolFrom] = toolTo
	toolFrom.CopyBasicFields(toolTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

// UnstageBranch stages instance and apply UnstageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the insance
//
// the algorithm stops along the course of graph if a vertex is already staged
func UnstageBranch[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for unstage branch
	case *AnalysisNeed:
		stage.UnstageBranchAnalysisNeed(target)

	case *Concept:
		stage.UnstageBranchConcept(target)

	case *ConceptShape:
		stage.UnstageBranchConceptShape(target)

	case *Concern:
		stage.UnstageBranchConcern(target)

	case *ConcernCompositionShape:
		stage.UnstageBranchConcernCompositionShape(target)

	case *ConcernInputShape:
		stage.UnstageBranchConcernInputShape(target)

	case *ConcernOutputShape:
		stage.UnstageBranchConcernOutputShape(target)

	case *ConcernShape:
		stage.UnstageBranchConcernShape(target)

	case *Deliverable:
		stage.UnstageBranchDeliverable(target)

	case *Diagram:
		stage.UnstageBranchDiagram(target)

	case *Library:
		stage.UnstageBranchLibrary(target)

	case *Note:
		stage.UnstageBranchNote(target)

	case *NoteProductShape:
		stage.UnstageBranchNoteProductShape(target)

	case *NoteShape:
		stage.UnstageBranchNoteShape(target)

	case *NoteStakeholderShape:
		stage.UnstageBranchNoteStakeholderShape(target)

	case *NoteTaskShape:
		stage.UnstageBranchNoteTaskShape(target)

	case *ProductCompositionShape:
		stage.UnstageBranchProductCompositionShape(target)

	case *ProductShape:
		stage.UnstageBranchProductShape(target)

	case *Requirement:
		stage.UnstageBranchRequirement(target)

	case *RequirementShape:
		stage.UnstageBranchRequirementShape(target)

	case *Stakeholder:
		stage.UnstageBranchStakeholder(target)

	case *StakeholderCompositionShape:
		stage.UnstageBranchStakeholderCompositionShape(target)

	case *StakeholderConcernShape:
		stage.UnstageBranchStakeholderConcernShape(target)

	case *StakeholderShape:
		stage.UnstageBranchStakeholderShape(target)

	case *SupportLevel:
		stage.UnstageBranchSupportLevel(target)

	case *Tool:
		stage.UnstageBranchTool(target)

	default:
		_ = target
	}
}

// insertion point for unstage branch per struct
func (stage *Stage) UnstageBranchAnalysisNeed(analysisneed *AnalysisNeed) {

	// check if instance is already staged
	if !IsStaged(stage, analysisneed) {
		return
	}

	analysisneed.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchConcept(concept *Concept) {

	// check if instance is already staged
	if !IsStaged(stage, concept) {
		return
	}

	concept.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _tool := range concept.Tools {
		UnstageBranch(stage, _tool)
	}

}

func (stage *Stage) UnstageBranchConceptShape(conceptshape *ConceptShape) {

	// check if instance is already staged
	if !IsStaged(stage, conceptshape) {
		return
	}

	conceptshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if conceptshape.Concept != nil {
		UnstageBranch(stage, conceptshape.Concept)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchConcern(concern *Concern) {

	// check if instance is already staged
	if !IsStaged(stage, concern) {
		return
	}

	concern.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _concern := range concern.SubConcerns {
		UnstageBranch(stage, _concern)
	}
	for _, _deliverable := range concern.Inputs {
		UnstageBranch(stage, _deliverable)
	}
	for _, _deliverable := range concern.Outputs {
		UnstageBranch(stage, _deliverable)
	}
	for _, _requirement := range concern.Requirements {
		UnstageBranch(stage, _requirement)
	}

}

func (stage *Stage) UnstageBranchConcernCompositionShape(concerncompositionshape *ConcernCompositionShape) {

	// check if instance is already staged
	if !IsStaged(stage, concerncompositionshape) {
		return
	}

	concerncompositionshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if concerncompositionshape.Concern != nil {
		UnstageBranch(stage, concerncompositionshape.Concern)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchConcernInputShape(concerninputshape *ConcernInputShape) {

	// check if instance is already staged
	if !IsStaged(stage, concerninputshape) {
		return
	}

	concerninputshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if concerninputshape.Deliverable != nil {
		UnstageBranch(stage, concerninputshape.Deliverable)
	}
	if concerninputshape.Concern != nil {
		UnstageBranch(stage, concerninputshape.Concern)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchConcernOutputShape(concernoutputshape *ConcernOutputShape) {

	// check if instance is already staged
	if !IsStaged(stage, concernoutputshape) {
		return
	}

	concernoutputshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if concernoutputshape.Task != nil {
		UnstageBranch(stage, concernoutputshape.Task)
	}
	if concernoutputshape.Product != nil {
		UnstageBranch(stage, concernoutputshape.Product)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchConcernShape(concernshape *ConcernShape) {

	// check if instance is already staged
	if !IsStaged(stage, concernshape) {
		return
	}

	concernshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if concernshape.Concern != nil {
		UnstageBranch(stage, concernshape.Concern)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchDeliverable(deliverable *Deliverable) {

	// check if instance is already staged
	if !IsStaged(stage, deliverable) {
		return
	}

	deliverable.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _deliverable := range deliverable.SubProducts {
		UnstageBranch(stage, _deliverable)
	}
	for _, _concept := range deliverable.Concepts {
		UnstageBranch(stage, _concept)
	}

}

func (stage *Stage) UnstageBranchDiagram(diagram *Diagram) {

	// check if instance is already staged
	if !IsStaged(stage, diagram) {
		return
	}

	diagram.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _concern := range diagram.ConcernsWhoseRequirementsNodeIsExpanded {
		UnstageBranch(stage, _concern)
	}
	for _, _productshape := range diagram.Product_Shapes {
		UnstageBranch(stage, _productshape)
	}
	for _, _deliverable := range diagram.ProductsWhoseNodeIsExpanded {
		UnstageBranch(stage, _deliverable)
	}
	for _, _productcompositionshape := range diagram.ProductComposition_Shapes {
		UnstageBranch(stage, _productcompositionshape)
	}
	for _, _concernshape := range diagram.Concern_Shapes {
		UnstageBranch(stage, _concernshape)
	}
	for _, _concern := range diagram.ConcernsWhoseNodeIsExpanded {
		UnstageBranch(stage, _concern)
	}
	for _, _concern := range diagram.ConcernsWhoseInputNodeIsExpanded {
		UnstageBranch(stage, _concern)
	}
	for _, _concern := range diagram.ConcernsWhoseStakeholderNodeIsExpanded {
		UnstageBranch(stage, _concern)
	}
	for _, _concern := range diagram.ConcernssWhoseOutputNodeIsExpanded {
		UnstageBranch(stage, _concern)
	}
	for _, _concerncompositionshape := range diagram.ConcernComposition_Shapes {
		UnstageBranch(stage, _concerncompositionshape)
	}
	for _, _concerninputshape := range diagram.ConcernInputShapes {
		UnstageBranch(stage, _concerninputshape)
	}
	for _, _concernoutputshape := range diagram.ConcernOutputShapes {
		UnstageBranch(stage, _concernoutputshape)
	}
	for _, _noteshape := range diagram.Note_Shapes {
		UnstageBranch(stage, _noteshape)
	}
	for _, _note := range diagram.NotesWhoseNodeIsExpanded {
		UnstageBranch(stage, _note)
	}
	for _, _noteproductshape := range diagram.NoteProductShapes {
		UnstageBranch(stage, _noteproductshape)
	}
	for _, _notetaskshape := range diagram.NoteTaskShapes {
		UnstageBranch(stage, _notetaskshape)
	}
	for _, _notestakeholdershape := range diagram.NoteResourceShapes {
		UnstageBranch(stage, _notestakeholdershape)
	}
	for _, _stakeholdershape := range diagram.Stakeholder_Shapes {
		UnstageBranch(stage, _stakeholdershape)
	}
	for _, _stakeholder := range diagram.ResourcesWhoseNodeIsExpanded {
		UnstageBranch(stage, _stakeholder)
	}
	for _, _stakeholdercompositionshape := range diagram.ResourceComposition_Shapes {
		UnstageBranch(stage, _stakeholdercompositionshape)
	}
	for _, _stakeholderconcernshape := range diagram.StakeholderConcernShapes {
		UnstageBranch(stage, _stakeholderconcernshape)
	}
	for _, _requirementshape := range diagram.Requirement_Shapes {
		UnstageBranch(stage, _requirementshape)
	}
	for _, _requirement := range diagram.RequirementsWhoseNodeIsExpanded {
		UnstageBranch(stage, _requirement)
	}
	for _, _conceptshape := range diagram.Concept_Shapes {
		UnstageBranch(stage, _conceptshape)
	}
	for _, _concept := range diagram.ConceptsWhoseNodeIsExpanded {
		UnstageBranch(stage, _concept)
	}

}

func (stage *Stage) UnstageBranchLibrary(library *Library) {

	// check if instance is already staged
	if !IsStaged(stage, library) {
		return
	}

	library.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _deliverable := range library.RootDeliverables {
		UnstageBranch(stage, _deliverable)
	}
	for _, _concern := range library.RootConcerns {
		UnstageBranch(stage, _concern)
	}
	for _, _stakeholder := range library.RootStakeholders {
		UnstageBranch(stage, _stakeholder)
	}
	for _, _requirement := range library.RootRequirements {
		UnstageBranch(stage, _requirement)
	}
	for _, _concept := range library.RootConcepts {
		UnstageBranch(stage, _concept)
	}
	for _, _analysisneed := range library.AnalysisNeeds {
		UnstageBranch(stage, _analysisneed)
	}
	for _, _note := range library.Notes {
		UnstageBranch(stage, _note)
	}
	for _, _diagram := range library.Diagrams {
		UnstageBranch(stage, _diagram)
	}
	for _, _library := range library.SubLibraries {
		UnstageBranch(stage, _library)
	}

}

func (stage *Stage) UnstageBranchNote(note *Note) {

	// check if instance is already staged
	if !IsStaged(stage, note) {
		return
	}

	note.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _deliverable := range note.Products {
		UnstageBranch(stage, _deliverable)
	}
	for _, _concern := range note.Tasks {
		UnstageBranch(stage, _concern)
	}
	for _, _stakeholder := range note.Resources {
		UnstageBranch(stage, _stakeholder)
	}

}

func (stage *Stage) UnstageBranchNoteProductShape(noteproductshape *NoteProductShape) {

	// check if instance is already staged
	if !IsStaged(stage, noteproductshape) {
		return
	}

	noteproductshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if noteproductshape.Note != nil {
		UnstageBranch(stage, noteproductshape.Note)
	}
	if noteproductshape.Product != nil {
		UnstageBranch(stage, noteproductshape.Product)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchNoteShape(noteshape *NoteShape) {

	// check if instance is already staged
	if !IsStaged(stage, noteshape) {
		return
	}

	noteshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if noteshape.Note != nil {
		UnstageBranch(stage, noteshape.Note)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchNoteStakeholderShape(notestakeholdershape *NoteStakeholderShape) {

	// check if instance is already staged
	if !IsStaged(stage, notestakeholdershape) {
		return
	}

	notestakeholdershape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if notestakeholdershape.Note != nil {
		UnstageBranch(stage, notestakeholdershape.Note)
	}
	if notestakeholdershape.Stakeholder != nil {
		UnstageBranch(stage, notestakeholdershape.Stakeholder)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchNoteTaskShape(notetaskshape *NoteTaskShape) {

	// check if instance is already staged
	if !IsStaged(stage, notetaskshape) {
		return
	}

	notetaskshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if notetaskshape.Note != nil {
		UnstageBranch(stage, notetaskshape.Note)
	}
	if notetaskshape.Task != nil {
		UnstageBranch(stage, notetaskshape.Task)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchProductCompositionShape(productcompositionshape *ProductCompositionShape) {

	// check if instance is already staged
	if !IsStaged(stage, productcompositionshape) {
		return
	}

	productcompositionshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if productcompositionshape.Product != nil {
		UnstageBranch(stage, productcompositionshape.Product)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchProductShape(productshape *ProductShape) {

	// check if instance is already staged
	if !IsStaged(stage, productshape) {
		return
	}

	productshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if productshape.Product != nil {
		UnstageBranch(stage, productshape.Product)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchRequirement(requirement *Requirement) {

	// check if instance is already staged
	if !IsStaged(stage, requirement) {
		return
	}

	requirement.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _supportlevel := range requirement.SupportLevels {
		UnstageBranch(stage, _supportlevel)
	}
	for _, _concept := range requirement.Concepts {
		UnstageBranch(stage, _concept)
	}

}

func (stage *Stage) UnstageBranchRequirementShape(requirementshape *RequirementShape) {

	// check if instance is already staged
	if !IsStaged(stage, requirementshape) {
		return
	}

	requirementshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if requirementshape.Requirement != nil {
		UnstageBranch(stage, requirementshape.Requirement)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchStakeholder(stakeholder *Stakeholder) {

	// check if instance is already staged
	if !IsStaged(stage, stakeholder) {
		return
	}

	stakeholder.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _concern := range stakeholder.Concerns {
		UnstageBranch(stage, _concern)
	}
	for _, _stakeholder := range stakeholder.SubStakeholders {
		UnstageBranch(stage, _stakeholder)
	}

}

func (stage *Stage) UnstageBranchStakeholderCompositionShape(stakeholdercompositionshape *StakeholderCompositionShape) {

	// check if instance is already staged
	if !IsStaged(stage, stakeholdercompositionshape) {
		return
	}

	stakeholdercompositionshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if stakeholdercompositionshape.Stakeholder != nil {
		UnstageBranch(stage, stakeholdercompositionshape.Stakeholder)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchStakeholderConcernShape(stakeholderconcernshape *StakeholderConcernShape) {

	// check if instance is already staged
	if !IsStaged(stage, stakeholderconcernshape) {
		return
	}

	stakeholderconcernshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if stakeholderconcernshape.Stakeholder != nil {
		UnstageBranch(stage, stakeholderconcernshape.Stakeholder)
	}
	if stakeholderconcernshape.Concern != nil {
		UnstageBranch(stage, stakeholderconcernshape.Concern)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchStakeholderShape(stakeholdershape *StakeholderShape) {

	// check if instance is already staged
	if !IsStaged(stage, stakeholdershape) {
		return
	}

	stakeholdershape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if stakeholdershape.Stakeholder != nil {
		UnstageBranch(stage, stakeholdershape.Stakeholder)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchSupportLevel(supportlevel *SupportLevel) {

	// check if instance is already staged
	if !IsStaged(stage, supportlevel) {
		return
	}

	supportlevel.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if supportlevel.Tool != nil {
		UnstageBranch(stage, supportlevel.Tool)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchTool(tool *Tool) {

	// check if instance is already staged
	if !IsStaged(stage, tool) {
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
	reference.Tools = reference.Tools[:0]
	for _, _b := range instance.Tools {
		reference.Tools = append(reference.Tools, stage.Tools_reference[_b])
	}
}

func (reference *ConceptShape) GongReconstructPointersFromReferences(stage *Stage, instance *ConceptShape) {
	// insertion point for pointers field
	if instance.Concept != nil {
		reference.Concept = stage.Concepts_reference[instance.Concept]
	}
	// insertion point for slice of pointers field
}

func (reference *Concern) GongReconstructPointersFromReferences(stage *Stage, instance *Concern) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.SubConcerns = reference.SubConcerns[:0]
	for _, _b := range instance.SubConcerns {
		reference.SubConcerns = append(reference.SubConcerns, stage.Concerns_reference[_b])
	}
	reference.Inputs = reference.Inputs[:0]
	for _, _b := range instance.Inputs {
		reference.Inputs = append(reference.Inputs, stage.Deliverables_reference[_b])
	}
	reference.Outputs = reference.Outputs[:0]
	for _, _b := range instance.Outputs {
		reference.Outputs = append(reference.Outputs, stage.Deliverables_reference[_b])
	}
	reference.Requirements = reference.Requirements[:0]
	for _, _b := range instance.Requirements {
		reference.Requirements = append(reference.Requirements, stage.Requirements_reference[_b])
	}
}

func (reference *ConcernCompositionShape) GongReconstructPointersFromReferences(stage *Stage, instance *ConcernCompositionShape) {
	// insertion point for pointers field
	if instance.Concern != nil {
		reference.Concern = stage.Concerns_reference[instance.Concern]
	}
	// insertion point for slice of pointers field
}

func (reference *ConcernInputShape) GongReconstructPointersFromReferences(stage *Stage, instance *ConcernInputShape) {
	// insertion point for pointers field
	if instance.Deliverable != nil {
		reference.Deliverable = stage.Deliverables_reference[instance.Deliverable]
	}
	if instance.Concern != nil {
		reference.Concern = stage.Concerns_reference[instance.Concern]
	}
	// insertion point for slice of pointers field
}

func (reference *ConcernOutputShape) GongReconstructPointersFromReferences(stage *Stage, instance *ConcernOutputShape) {
	// insertion point for pointers field
	if instance.Task != nil {
		reference.Task = stage.Concerns_reference[instance.Task]
	}
	if instance.Product != nil {
		reference.Product = stage.Deliverables_reference[instance.Product]
	}
	// insertion point for slice of pointers field
}

func (reference *ConcernShape) GongReconstructPointersFromReferences(stage *Stage, instance *ConcernShape) {
	// insertion point for pointers field
	if instance.Concern != nil {
		reference.Concern = stage.Concerns_reference[instance.Concern]
	}
	// insertion point for slice of pointers field
}

func (reference *Deliverable) GongReconstructPointersFromReferences(stage *Stage, instance *Deliverable) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.SubProducts = reference.SubProducts[:0]
	for _, _b := range instance.SubProducts {
		reference.SubProducts = append(reference.SubProducts, stage.Deliverables_reference[_b])
	}
	reference.Concepts = reference.Concepts[:0]
	for _, _b := range instance.Concepts {
		reference.Concepts = append(reference.Concepts, stage.Concepts_reference[_b])
	}
}

func (reference *Diagram) GongReconstructPointersFromReferences(stage *Stage, instance *Diagram) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.ConcernsWhoseRequirementsNodeIsExpanded = reference.ConcernsWhoseRequirementsNodeIsExpanded[:0]
	for _, _b := range instance.ConcernsWhoseRequirementsNodeIsExpanded {
		reference.ConcernsWhoseRequirementsNodeIsExpanded = append(reference.ConcernsWhoseRequirementsNodeIsExpanded, stage.Concerns_reference[_b])
	}
	reference.Product_Shapes = reference.Product_Shapes[:0]
	for _, _b := range instance.Product_Shapes {
		reference.Product_Shapes = append(reference.Product_Shapes, stage.ProductShapes_reference[_b])
	}
	reference.ProductsWhoseNodeIsExpanded = reference.ProductsWhoseNodeIsExpanded[:0]
	for _, _b := range instance.ProductsWhoseNodeIsExpanded {
		reference.ProductsWhoseNodeIsExpanded = append(reference.ProductsWhoseNodeIsExpanded, stage.Deliverables_reference[_b])
	}
	reference.ProductComposition_Shapes = reference.ProductComposition_Shapes[:0]
	for _, _b := range instance.ProductComposition_Shapes {
		reference.ProductComposition_Shapes = append(reference.ProductComposition_Shapes, stage.ProductCompositionShapes_reference[_b])
	}
	reference.Concern_Shapes = reference.Concern_Shapes[:0]
	for _, _b := range instance.Concern_Shapes {
		reference.Concern_Shapes = append(reference.Concern_Shapes, stage.ConcernShapes_reference[_b])
	}
	reference.ConcernsWhoseNodeIsExpanded = reference.ConcernsWhoseNodeIsExpanded[:0]
	for _, _b := range instance.ConcernsWhoseNodeIsExpanded {
		reference.ConcernsWhoseNodeIsExpanded = append(reference.ConcernsWhoseNodeIsExpanded, stage.Concerns_reference[_b])
	}
	reference.ConcernsWhoseInputNodeIsExpanded = reference.ConcernsWhoseInputNodeIsExpanded[:0]
	for _, _b := range instance.ConcernsWhoseInputNodeIsExpanded {
		reference.ConcernsWhoseInputNodeIsExpanded = append(reference.ConcernsWhoseInputNodeIsExpanded, stage.Concerns_reference[_b])
	}
	reference.ConcernsWhoseStakeholderNodeIsExpanded = reference.ConcernsWhoseStakeholderNodeIsExpanded[:0]
	for _, _b := range instance.ConcernsWhoseStakeholderNodeIsExpanded {
		reference.ConcernsWhoseStakeholderNodeIsExpanded = append(reference.ConcernsWhoseStakeholderNodeIsExpanded, stage.Concerns_reference[_b])
	}
	reference.ConcernssWhoseOutputNodeIsExpanded = reference.ConcernssWhoseOutputNodeIsExpanded[:0]
	for _, _b := range instance.ConcernssWhoseOutputNodeIsExpanded {
		reference.ConcernssWhoseOutputNodeIsExpanded = append(reference.ConcernssWhoseOutputNodeIsExpanded, stage.Concerns_reference[_b])
	}
	reference.ConcernComposition_Shapes = reference.ConcernComposition_Shapes[:0]
	for _, _b := range instance.ConcernComposition_Shapes {
		reference.ConcernComposition_Shapes = append(reference.ConcernComposition_Shapes, stage.ConcernCompositionShapes_reference[_b])
	}
	reference.ConcernInputShapes = reference.ConcernInputShapes[:0]
	for _, _b := range instance.ConcernInputShapes {
		reference.ConcernInputShapes = append(reference.ConcernInputShapes, stage.ConcernInputShapes_reference[_b])
	}
	reference.ConcernOutputShapes = reference.ConcernOutputShapes[:0]
	for _, _b := range instance.ConcernOutputShapes {
		reference.ConcernOutputShapes = append(reference.ConcernOutputShapes, stage.ConcernOutputShapes_reference[_b])
	}
	reference.Note_Shapes = reference.Note_Shapes[:0]
	for _, _b := range instance.Note_Shapes {
		reference.Note_Shapes = append(reference.Note_Shapes, stage.NoteShapes_reference[_b])
	}
	reference.NotesWhoseNodeIsExpanded = reference.NotesWhoseNodeIsExpanded[:0]
	for _, _b := range instance.NotesWhoseNodeIsExpanded {
		reference.NotesWhoseNodeIsExpanded = append(reference.NotesWhoseNodeIsExpanded, stage.Notes_reference[_b])
	}
	reference.NoteProductShapes = reference.NoteProductShapes[:0]
	for _, _b := range instance.NoteProductShapes {
		reference.NoteProductShapes = append(reference.NoteProductShapes, stage.NoteProductShapes_reference[_b])
	}
	reference.NoteTaskShapes = reference.NoteTaskShapes[:0]
	for _, _b := range instance.NoteTaskShapes {
		reference.NoteTaskShapes = append(reference.NoteTaskShapes, stage.NoteTaskShapes_reference[_b])
	}
	reference.NoteResourceShapes = reference.NoteResourceShapes[:0]
	for _, _b := range instance.NoteResourceShapes {
		reference.NoteResourceShapes = append(reference.NoteResourceShapes, stage.NoteStakeholderShapes_reference[_b])
	}
	reference.Stakeholder_Shapes = reference.Stakeholder_Shapes[:0]
	for _, _b := range instance.Stakeholder_Shapes {
		reference.Stakeholder_Shapes = append(reference.Stakeholder_Shapes, stage.StakeholderShapes_reference[_b])
	}
	reference.ResourcesWhoseNodeIsExpanded = reference.ResourcesWhoseNodeIsExpanded[:0]
	for _, _b := range instance.ResourcesWhoseNodeIsExpanded {
		reference.ResourcesWhoseNodeIsExpanded = append(reference.ResourcesWhoseNodeIsExpanded, stage.Stakeholders_reference[_b])
	}
	reference.ResourceComposition_Shapes = reference.ResourceComposition_Shapes[:0]
	for _, _b := range instance.ResourceComposition_Shapes {
		reference.ResourceComposition_Shapes = append(reference.ResourceComposition_Shapes, stage.StakeholderCompositionShapes_reference[_b])
	}
	reference.StakeholderConcernShapes = reference.StakeholderConcernShapes[:0]
	for _, _b := range instance.StakeholderConcernShapes {
		reference.StakeholderConcernShapes = append(reference.StakeholderConcernShapes, stage.StakeholderConcernShapes_reference[_b])
	}
	reference.Requirement_Shapes = reference.Requirement_Shapes[:0]
	for _, _b := range instance.Requirement_Shapes {
		reference.Requirement_Shapes = append(reference.Requirement_Shapes, stage.RequirementShapes_reference[_b])
	}
	reference.RequirementsWhoseNodeIsExpanded = reference.RequirementsWhoseNodeIsExpanded[:0]
	for _, _b := range instance.RequirementsWhoseNodeIsExpanded {
		reference.RequirementsWhoseNodeIsExpanded = append(reference.RequirementsWhoseNodeIsExpanded, stage.Requirements_reference[_b])
	}
	reference.Concept_Shapes = reference.Concept_Shapes[:0]
	for _, _b := range instance.Concept_Shapes {
		reference.Concept_Shapes = append(reference.Concept_Shapes, stage.ConceptShapes_reference[_b])
	}
	reference.ConceptsWhoseNodeIsExpanded = reference.ConceptsWhoseNodeIsExpanded[:0]
	for _, _b := range instance.ConceptsWhoseNodeIsExpanded {
		reference.ConceptsWhoseNodeIsExpanded = append(reference.ConceptsWhoseNodeIsExpanded, stage.Concepts_reference[_b])
	}
}

func (reference *Library) GongReconstructPointersFromReferences(stage *Stage, instance *Library) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.RootDeliverables = reference.RootDeliverables[:0]
	for _, _b := range instance.RootDeliverables {
		reference.RootDeliverables = append(reference.RootDeliverables, stage.Deliverables_reference[_b])
	}
	reference.RootConcerns = reference.RootConcerns[:0]
	for _, _b := range instance.RootConcerns {
		reference.RootConcerns = append(reference.RootConcerns, stage.Concerns_reference[_b])
	}
	reference.RootStakeholders = reference.RootStakeholders[:0]
	for _, _b := range instance.RootStakeholders {
		reference.RootStakeholders = append(reference.RootStakeholders, stage.Stakeholders_reference[_b])
	}
	reference.RootRequirements = reference.RootRequirements[:0]
	for _, _b := range instance.RootRequirements {
		reference.RootRequirements = append(reference.RootRequirements, stage.Requirements_reference[_b])
	}
	reference.RootConcepts = reference.RootConcepts[:0]
	for _, _b := range instance.RootConcepts {
		reference.RootConcepts = append(reference.RootConcepts, stage.Concepts_reference[_b])
	}
	reference.AnalysisNeeds = reference.AnalysisNeeds[:0]
	for _, _b := range instance.AnalysisNeeds {
		reference.AnalysisNeeds = append(reference.AnalysisNeeds, stage.AnalysisNeeds_reference[_b])
	}
	reference.Notes = reference.Notes[:0]
	for _, _b := range instance.Notes {
		reference.Notes = append(reference.Notes, stage.Notes_reference[_b])
	}
	reference.Diagrams = reference.Diagrams[:0]
	for _, _b := range instance.Diagrams {
		reference.Diagrams = append(reference.Diagrams, stage.Diagrams_reference[_b])
	}
	reference.SubLibraries = reference.SubLibraries[:0]
	for _, _b := range instance.SubLibraries {
		reference.SubLibraries = append(reference.SubLibraries, stage.Librarys_reference[_b])
	}
}

func (reference *Note) GongReconstructPointersFromReferences(stage *Stage, instance *Note) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.Products = reference.Products[:0]
	for _, _b := range instance.Products {
		reference.Products = append(reference.Products, stage.Deliverables_reference[_b])
	}
	reference.Tasks = reference.Tasks[:0]
	for _, _b := range instance.Tasks {
		reference.Tasks = append(reference.Tasks, stage.Concerns_reference[_b])
	}
	reference.Resources = reference.Resources[:0]
	for _, _b := range instance.Resources {
		reference.Resources = append(reference.Resources, stage.Stakeholders_reference[_b])
	}
}

func (reference *NoteProductShape) GongReconstructPointersFromReferences(stage *Stage, instance *NoteProductShape) {
	// insertion point for pointers field
	if instance.Note != nil {
		reference.Note = stage.Notes_reference[instance.Note]
	}
	if instance.Product != nil {
		reference.Product = stage.Deliverables_reference[instance.Product]
	}
	// insertion point for slice of pointers field
}

func (reference *NoteShape) GongReconstructPointersFromReferences(stage *Stage, instance *NoteShape) {
	// insertion point for pointers field
	if instance.Note != nil {
		reference.Note = stage.Notes_reference[instance.Note]
	}
	// insertion point for slice of pointers field
}

func (reference *NoteStakeholderShape) GongReconstructPointersFromReferences(stage *Stage, instance *NoteStakeholderShape) {
	// insertion point for pointers field
	if instance.Note != nil {
		reference.Note = stage.Notes_reference[instance.Note]
	}
	if instance.Stakeholder != nil {
		reference.Stakeholder = stage.Stakeholders_reference[instance.Stakeholder]
	}
	// insertion point for slice of pointers field
}

func (reference *NoteTaskShape) GongReconstructPointersFromReferences(stage *Stage, instance *NoteTaskShape) {
	// insertion point for pointers field
	if instance.Note != nil {
		reference.Note = stage.Notes_reference[instance.Note]
	}
	if instance.Task != nil {
		reference.Task = stage.Concerns_reference[instance.Task]
	}
	// insertion point for slice of pointers field
}

func (reference *ProductCompositionShape) GongReconstructPointersFromReferences(stage *Stage, instance *ProductCompositionShape) {
	// insertion point for pointers field
	if instance.Product != nil {
		reference.Product = stage.Deliverables_reference[instance.Product]
	}
	// insertion point for slice of pointers field
}

func (reference *ProductShape) GongReconstructPointersFromReferences(stage *Stage, instance *ProductShape) {
	// insertion point for pointers field
	if instance.Product != nil {
		reference.Product = stage.Deliverables_reference[instance.Product]
	}
	// insertion point for slice of pointers field
}

func (reference *Requirement) GongReconstructPointersFromReferences(stage *Stage, instance *Requirement) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.SupportLevels = reference.SupportLevels[:0]
	for _, _b := range instance.SupportLevels {
		reference.SupportLevels = append(reference.SupportLevels, stage.SupportLevels_reference[_b])
	}
	reference.Concepts = reference.Concepts[:0]
	for _, _b := range instance.Concepts {
		reference.Concepts = append(reference.Concepts, stage.Concepts_reference[_b])
	}
}

func (reference *RequirementShape) GongReconstructPointersFromReferences(stage *Stage, instance *RequirementShape) {
	// insertion point for pointers field
	if instance.Requirement != nil {
		reference.Requirement = stage.Requirements_reference[instance.Requirement]
	}
	// insertion point for slice of pointers field
}

func (reference *Stakeholder) GongReconstructPointersFromReferences(stage *Stage, instance *Stakeholder) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.Concerns = reference.Concerns[:0]
	for _, _b := range instance.Concerns {
		reference.Concerns = append(reference.Concerns, stage.Concerns_reference[_b])
	}
	reference.SubStakeholders = reference.SubStakeholders[:0]
	for _, _b := range instance.SubStakeholders {
		reference.SubStakeholders = append(reference.SubStakeholders, stage.Stakeholders_reference[_b])
	}
}

func (reference *StakeholderCompositionShape) GongReconstructPointersFromReferences(stage *Stage, instance *StakeholderCompositionShape) {
	// insertion point for pointers field
	if instance.Stakeholder != nil {
		reference.Stakeholder = stage.Stakeholders_reference[instance.Stakeholder]
	}
	// insertion point for slice of pointers field
}

func (reference *StakeholderConcernShape) GongReconstructPointersFromReferences(stage *Stage, instance *StakeholderConcernShape) {
	// insertion point for pointers field
	if instance.Stakeholder != nil {
		reference.Stakeholder = stage.Stakeholders_reference[instance.Stakeholder]
	}
	if instance.Concern != nil {
		reference.Concern = stage.Concerns_reference[instance.Concern]
	}
	// insertion point for slice of pointers field
}

func (reference *StakeholderShape) GongReconstructPointersFromReferences(stage *Stage, instance *StakeholderShape) {
	// insertion point for pointers field
	if instance.Stakeholder != nil {
		reference.Stakeholder = stage.Stakeholders_reference[instance.Stakeholder]
	}
	// insertion point for slice of pointers field
}

func (reference *SupportLevel) GongReconstructPointersFromReferences(stage *Stage, instance *SupportLevel) {
	// insertion point for pointers field
	if instance.Tool != nil {
		reference.Tool = stage.Tools_reference[instance.Tool]
	}
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
	var _Tools []*Tool
	for _, _reference := range reference.Tools {
		if _instance, ok := stage.Tools_instance[_reference]; ok {
			_Tools = append(_Tools, _instance)
		}
	}
	reference.Tools = _Tools
}

func (reference *ConceptShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.Concept; _reference != nil {
		reference.Concept = nil
		if _instance, ok := stage.Concepts_instance[_reference]; ok {
			reference.Concept = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *Concern) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _SubConcerns []*Concern
	for _, _reference := range reference.SubConcerns {
		if _instance, ok := stage.Concerns_instance[_reference]; ok {
			_SubConcerns = append(_SubConcerns, _instance)
		}
	}
	reference.SubConcerns = _SubConcerns
	var _Inputs []*Deliverable
	for _, _reference := range reference.Inputs {
		if _instance, ok := stage.Deliverables_instance[_reference]; ok {
			_Inputs = append(_Inputs, _instance)
		}
	}
	reference.Inputs = _Inputs
	var _Outputs []*Deliverable
	for _, _reference := range reference.Outputs {
		if _instance, ok := stage.Deliverables_instance[_reference]; ok {
			_Outputs = append(_Outputs, _instance)
		}
	}
	reference.Outputs = _Outputs
	var _Requirements []*Requirement
	for _, _reference := range reference.Requirements {
		if _instance, ok := stage.Requirements_instance[_reference]; ok {
			_Requirements = append(_Requirements, _instance)
		}
	}
	reference.Requirements = _Requirements
}

func (reference *ConcernCompositionShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.Concern; _reference != nil {
		reference.Concern = nil
		if _instance, ok := stage.Concerns_instance[_reference]; ok {
			reference.Concern = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *ConcernInputShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.Deliverable; _reference != nil {
		reference.Deliverable = nil
		if _instance, ok := stage.Deliverables_instance[_reference]; ok {
			reference.Deliverable = _instance
		}
	}
	if _reference := reference.Concern; _reference != nil {
		reference.Concern = nil
		if _instance, ok := stage.Concerns_instance[_reference]; ok {
			reference.Concern = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *ConcernOutputShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.Task; _reference != nil {
		reference.Task = nil
		if _instance, ok := stage.Concerns_instance[_reference]; ok {
			reference.Task = _instance
		}
	}
	if _reference := reference.Product; _reference != nil {
		reference.Product = nil
		if _instance, ok := stage.Deliverables_instance[_reference]; ok {
			reference.Product = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *ConcernShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.Concern; _reference != nil {
		reference.Concern = nil
		if _instance, ok := stage.Concerns_instance[_reference]; ok {
			reference.Concern = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *Deliverable) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _SubProducts []*Deliverable
	for _, _reference := range reference.SubProducts {
		if _instance, ok := stage.Deliverables_instance[_reference]; ok {
			_SubProducts = append(_SubProducts, _instance)
		}
	}
	reference.SubProducts = _SubProducts
	var _Concepts []*Concept
	for _, _reference := range reference.Concepts {
		if _instance, ok := stage.Concepts_instance[_reference]; ok {
			_Concepts = append(_Concepts, _instance)
		}
	}
	reference.Concepts = _Concepts
}

func (reference *Diagram) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _ConcernsWhoseRequirementsNodeIsExpanded []*Concern
	for _, _reference := range reference.ConcernsWhoseRequirementsNodeIsExpanded {
		if _instance, ok := stage.Concerns_instance[_reference]; ok {
			_ConcernsWhoseRequirementsNodeIsExpanded = append(_ConcernsWhoseRequirementsNodeIsExpanded, _instance)
		}
	}
	reference.ConcernsWhoseRequirementsNodeIsExpanded = _ConcernsWhoseRequirementsNodeIsExpanded
	var _Product_Shapes []*ProductShape
	for _, _reference := range reference.Product_Shapes {
		if _instance, ok := stage.ProductShapes_instance[_reference]; ok {
			_Product_Shapes = append(_Product_Shapes, _instance)
		}
	}
	reference.Product_Shapes = _Product_Shapes
	var _ProductsWhoseNodeIsExpanded []*Deliverable
	for _, _reference := range reference.ProductsWhoseNodeIsExpanded {
		if _instance, ok := stage.Deliverables_instance[_reference]; ok {
			_ProductsWhoseNodeIsExpanded = append(_ProductsWhoseNodeIsExpanded, _instance)
		}
	}
	reference.ProductsWhoseNodeIsExpanded = _ProductsWhoseNodeIsExpanded
	var _ProductComposition_Shapes []*ProductCompositionShape
	for _, _reference := range reference.ProductComposition_Shapes {
		if _instance, ok := stage.ProductCompositionShapes_instance[_reference]; ok {
			_ProductComposition_Shapes = append(_ProductComposition_Shapes, _instance)
		}
	}
	reference.ProductComposition_Shapes = _ProductComposition_Shapes
	var _Concern_Shapes []*ConcernShape
	for _, _reference := range reference.Concern_Shapes {
		if _instance, ok := stage.ConcernShapes_instance[_reference]; ok {
			_Concern_Shapes = append(_Concern_Shapes, _instance)
		}
	}
	reference.Concern_Shapes = _Concern_Shapes
	var _ConcernsWhoseNodeIsExpanded []*Concern
	for _, _reference := range reference.ConcernsWhoseNodeIsExpanded {
		if _instance, ok := stage.Concerns_instance[_reference]; ok {
			_ConcernsWhoseNodeIsExpanded = append(_ConcernsWhoseNodeIsExpanded, _instance)
		}
	}
	reference.ConcernsWhoseNodeIsExpanded = _ConcernsWhoseNodeIsExpanded
	var _ConcernsWhoseInputNodeIsExpanded []*Concern
	for _, _reference := range reference.ConcernsWhoseInputNodeIsExpanded {
		if _instance, ok := stage.Concerns_instance[_reference]; ok {
			_ConcernsWhoseInputNodeIsExpanded = append(_ConcernsWhoseInputNodeIsExpanded, _instance)
		}
	}
	reference.ConcernsWhoseInputNodeIsExpanded = _ConcernsWhoseInputNodeIsExpanded
	var _ConcernsWhoseStakeholderNodeIsExpanded []*Concern
	for _, _reference := range reference.ConcernsWhoseStakeholderNodeIsExpanded {
		if _instance, ok := stage.Concerns_instance[_reference]; ok {
			_ConcernsWhoseStakeholderNodeIsExpanded = append(_ConcernsWhoseStakeholderNodeIsExpanded, _instance)
		}
	}
	reference.ConcernsWhoseStakeholderNodeIsExpanded = _ConcernsWhoseStakeholderNodeIsExpanded
	var _ConcernssWhoseOutputNodeIsExpanded []*Concern
	for _, _reference := range reference.ConcernssWhoseOutputNodeIsExpanded {
		if _instance, ok := stage.Concerns_instance[_reference]; ok {
			_ConcernssWhoseOutputNodeIsExpanded = append(_ConcernssWhoseOutputNodeIsExpanded, _instance)
		}
	}
	reference.ConcernssWhoseOutputNodeIsExpanded = _ConcernssWhoseOutputNodeIsExpanded
	var _ConcernComposition_Shapes []*ConcernCompositionShape
	for _, _reference := range reference.ConcernComposition_Shapes {
		if _instance, ok := stage.ConcernCompositionShapes_instance[_reference]; ok {
			_ConcernComposition_Shapes = append(_ConcernComposition_Shapes, _instance)
		}
	}
	reference.ConcernComposition_Shapes = _ConcernComposition_Shapes
	var _ConcernInputShapes []*ConcernInputShape
	for _, _reference := range reference.ConcernInputShapes {
		if _instance, ok := stage.ConcernInputShapes_instance[_reference]; ok {
			_ConcernInputShapes = append(_ConcernInputShapes, _instance)
		}
	}
	reference.ConcernInputShapes = _ConcernInputShapes
	var _ConcernOutputShapes []*ConcernOutputShape
	for _, _reference := range reference.ConcernOutputShapes {
		if _instance, ok := stage.ConcernOutputShapes_instance[_reference]; ok {
			_ConcernOutputShapes = append(_ConcernOutputShapes, _instance)
		}
	}
	reference.ConcernOutputShapes = _ConcernOutputShapes
	var _Note_Shapes []*NoteShape
	for _, _reference := range reference.Note_Shapes {
		if _instance, ok := stage.NoteShapes_instance[_reference]; ok {
			_Note_Shapes = append(_Note_Shapes, _instance)
		}
	}
	reference.Note_Shapes = _Note_Shapes
	var _NotesWhoseNodeIsExpanded []*Note
	for _, _reference := range reference.NotesWhoseNodeIsExpanded {
		if _instance, ok := stage.Notes_instance[_reference]; ok {
			_NotesWhoseNodeIsExpanded = append(_NotesWhoseNodeIsExpanded, _instance)
		}
	}
	reference.NotesWhoseNodeIsExpanded = _NotesWhoseNodeIsExpanded
	var _NoteProductShapes []*NoteProductShape
	for _, _reference := range reference.NoteProductShapes {
		if _instance, ok := stage.NoteProductShapes_instance[_reference]; ok {
			_NoteProductShapes = append(_NoteProductShapes, _instance)
		}
	}
	reference.NoteProductShapes = _NoteProductShapes
	var _NoteTaskShapes []*NoteTaskShape
	for _, _reference := range reference.NoteTaskShapes {
		if _instance, ok := stage.NoteTaskShapes_instance[_reference]; ok {
			_NoteTaskShapes = append(_NoteTaskShapes, _instance)
		}
	}
	reference.NoteTaskShapes = _NoteTaskShapes
	var _NoteResourceShapes []*NoteStakeholderShape
	for _, _reference := range reference.NoteResourceShapes {
		if _instance, ok := stage.NoteStakeholderShapes_instance[_reference]; ok {
			_NoteResourceShapes = append(_NoteResourceShapes, _instance)
		}
	}
	reference.NoteResourceShapes = _NoteResourceShapes
	var _Stakeholder_Shapes []*StakeholderShape
	for _, _reference := range reference.Stakeholder_Shapes {
		if _instance, ok := stage.StakeholderShapes_instance[_reference]; ok {
			_Stakeholder_Shapes = append(_Stakeholder_Shapes, _instance)
		}
	}
	reference.Stakeholder_Shapes = _Stakeholder_Shapes
	var _ResourcesWhoseNodeIsExpanded []*Stakeholder
	for _, _reference := range reference.ResourcesWhoseNodeIsExpanded {
		if _instance, ok := stage.Stakeholders_instance[_reference]; ok {
			_ResourcesWhoseNodeIsExpanded = append(_ResourcesWhoseNodeIsExpanded, _instance)
		}
	}
	reference.ResourcesWhoseNodeIsExpanded = _ResourcesWhoseNodeIsExpanded
	var _ResourceComposition_Shapes []*StakeholderCompositionShape
	for _, _reference := range reference.ResourceComposition_Shapes {
		if _instance, ok := stage.StakeholderCompositionShapes_instance[_reference]; ok {
			_ResourceComposition_Shapes = append(_ResourceComposition_Shapes, _instance)
		}
	}
	reference.ResourceComposition_Shapes = _ResourceComposition_Shapes
	var _StakeholderConcernShapes []*StakeholderConcernShape
	for _, _reference := range reference.StakeholderConcernShapes {
		if _instance, ok := stage.StakeholderConcernShapes_instance[_reference]; ok {
			_StakeholderConcernShapes = append(_StakeholderConcernShapes, _instance)
		}
	}
	reference.StakeholderConcernShapes = _StakeholderConcernShapes
	var _Requirement_Shapes []*RequirementShape
	for _, _reference := range reference.Requirement_Shapes {
		if _instance, ok := stage.RequirementShapes_instance[_reference]; ok {
			_Requirement_Shapes = append(_Requirement_Shapes, _instance)
		}
	}
	reference.Requirement_Shapes = _Requirement_Shapes
	var _RequirementsWhoseNodeIsExpanded []*Requirement
	for _, _reference := range reference.RequirementsWhoseNodeIsExpanded {
		if _instance, ok := stage.Requirements_instance[_reference]; ok {
			_RequirementsWhoseNodeIsExpanded = append(_RequirementsWhoseNodeIsExpanded, _instance)
		}
	}
	reference.RequirementsWhoseNodeIsExpanded = _RequirementsWhoseNodeIsExpanded
	var _Concept_Shapes []*ConceptShape
	for _, _reference := range reference.Concept_Shapes {
		if _instance, ok := stage.ConceptShapes_instance[_reference]; ok {
			_Concept_Shapes = append(_Concept_Shapes, _instance)
		}
	}
	reference.Concept_Shapes = _Concept_Shapes
	var _ConceptsWhoseNodeIsExpanded []*Concept
	for _, _reference := range reference.ConceptsWhoseNodeIsExpanded {
		if _instance, ok := stage.Concepts_instance[_reference]; ok {
			_ConceptsWhoseNodeIsExpanded = append(_ConceptsWhoseNodeIsExpanded, _instance)
		}
	}
	reference.ConceptsWhoseNodeIsExpanded = _ConceptsWhoseNodeIsExpanded
}

func (reference *Library) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _RootDeliverables []*Deliverable
	for _, _reference := range reference.RootDeliverables {
		if _instance, ok := stage.Deliverables_instance[_reference]; ok {
			_RootDeliverables = append(_RootDeliverables, _instance)
		}
	}
	reference.RootDeliverables = _RootDeliverables
	var _RootConcerns []*Concern
	for _, _reference := range reference.RootConcerns {
		if _instance, ok := stage.Concerns_instance[_reference]; ok {
			_RootConcerns = append(_RootConcerns, _instance)
		}
	}
	reference.RootConcerns = _RootConcerns
	var _RootStakeholders []*Stakeholder
	for _, _reference := range reference.RootStakeholders {
		if _instance, ok := stage.Stakeholders_instance[_reference]; ok {
			_RootStakeholders = append(_RootStakeholders, _instance)
		}
	}
	reference.RootStakeholders = _RootStakeholders
	var _RootRequirements []*Requirement
	for _, _reference := range reference.RootRequirements {
		if _instance, ok := stage.Requirements_instance[_reference]; ok {
			_RootRequirements = append(_RootRequirements, _instance)
		}
	}
	reference.RootRequirements = _RootRequirements
	var _RootConcepts []*Concept
	for _, _reference := range reference.RootConcepts {
		if _instance, ok := stage.Concepts_instance[_reference]; ok {
			_RootConcepts = append(_RootConcepts, _instance)
		}
	}
	reference.RootConcepts = _RootConcepts
	var _AnalysisNeeds []*AnalysisNeed
	for _, _reference := range reference.AnalysisNeeds {
		if _instance, ok := stage.AnalysisNeeds_instance[_reference]; ok {
			_AnalysisNeeds = append(_AnalysisNeeds, _instance)
		}
	}
	reference.AnalysisNeeds = _AnalysisNeeds
	var _Notes []*Note
	for _, _reference := range reference.Notes {
		if _instance, ok := stage.Notes_instance[_reference]; ok {
			_Notes = append(_Notes, _instance)
		}
	}
	reference.Notes = _Notes
	var _Diagrams []*Diagram
	for _, _reference := range reference.Diagrams {
		if _instance, ok := stage.Diagrams_instance[_reference]; ok {
			_Diagrams = append(_Diagrams, _instance)
		}
	}
	reference.Diagrams = _Diagrams
	var _SubLibraries []*Library
	for _, _reference := range reference.SubLibraries {
		if _instance, ok := stage.Librarys_instance[_reference]; ok {
			_SubLibraries = append(_SubLibraries, _instance)
		}
	}
	reference.SubLibraries = _SubLibraries
}

func (reference *Note) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _Products []*Deliverable
	for _, _reference := range reference.Products {
		if _instance, ok := stage.Deliverables_instance[_reference]; ok {
			_Products = append(_Products, _instance)
		}
	}
	reference.Products = _Products
	var _Tasks []*Concern
	for _, _reference := range reference.Tasks {
		if _instance, ok := stage.Concerns_instance[_reference]; ok {
			_Tasks = append(_Tasks, _instance)
		}
	}
	reference.Tasks = _Tasks
	var _Resources []*Stakeholder
	for _, _reference := range reference.Resources {
		if _instance, ok := stage.Stakeholders_instance[_reference]; ok {
			_Resources = append(_Resources, _instance)
		}
	}
	reference.Resources = _Resources
}

func (reference *NoteProductShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.Note; _reference != nil {
		reference.Note = nil
		if _instance, ok := stage.Notes_instance[_reference]; ok {
			reference.Note = _instance
		}
	}
	if _reference := reference.Product; _reference != nil {
		reference.Product = nil
		if _instance, ok := stage.Deliverables_instance[_reference]; ok {
			reference.Product = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *NoteShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.Note; _reference != nil {
		reference.Note = nil
		if _instance, ok := stage.Notes_instance[_reference]; ok {
			reference.Note = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *NoteStakeholderShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.Note; _reference != nil {
		reference.Note = nil
		if _instance, ok := stage.Notes_instance[_reference]; ok {
			reference.Note = _instance
		}
	}
	if _reference := reference.Stakeholder; _reference != nil {
		reference.Stakeholder = nil
		if _instance, ok := stage.Stakeholders_instance[_reference]; ok {
			reference.Stakeholder = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *NoteTaskShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.Note; _reference != nil {
		reference.Note = nil
		if _instance, ok := stage.Notes_instance[_reference]; ok {
			reference.Note = _instance
		}
	}
	if _reference := reference.Task; _reference != nil {
		reference.Task = nil
		if _instance, ok := stage.Concerns_instance[_reference]; ok {
			reference.Task = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *ProductCompositionShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.Product; _reference != nil {
		reference.Product = nil
		if _instance, ok := stage.Deliverables_instance[_reference]; ok {
			reference.Product = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *ProductShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.Product; _reference != nil {
		reference.Product = nil
		if _instance, ok := stage.Deliverables_instance[_reference]; ok {
			reference.Product = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *Requirement) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _SupportLevels []*SupportLevel
	for _, _reference := range reference.SupportLevels {
		if _instance, ok := stage.SupportLevels_instance[_reference]; ok {
			_SupportLevels = append(_SupportLevels, _instance)
		}
	}
	reference.SupportLevels = _SupportLevels
	var _Concepts []*Concept
	for _, _reference := range reference.Concepts {
		if _instance, ok := stage.Concepts_instance[_reference]; ok {
			_Concepts = append(_Concepts, _instance)
		}
	}
	reference.Concepts = _Concepts
}

func (reference *RequirementShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.Requirement; _reference != nil {
		reference.Requirement = nil
		if _instance, ok := stage.Requirements_instance[_reference]; ok {
			reference.Requirement = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *Stakeholder) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _Concerns []*Concern
	for _, _reference := range reference.Concerns {
		if _instance, ok := stage.Concerns_instance[_reference]; ok {
			_Concerns = append(_Concerns, _instance)
		}
	}
	reference.Concerns = _Concerns
	var _SubStakeholders []*Stakeholder
	for _, _reference := range reference.SubStakeholders {
		if _instance, ok := stage.Stakeholders_instance[_reference]; ok {
			_SubStakeholders = append(_SubStakeholders, _instance)
		}
	}
	reference.SubStakeholders = _SubStakeholders
}

func (reference *StakeholderCompositionShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.Stakeholder; _reference != nil {
		reference.Stakeholder = nil
		if _instance, ok := stage.Stakeholders_instance[_reference]; ok {
			reference.Stakeholder = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *StakeholderConcernShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.Stakeholder; _reference != nil {
		reference.Stakeholder = nil
		if _instance, ok := stage.Stakeholders_instance[_reference]; ok {
			reference.Stakeholder = _instance
		}
	}
	if _reference := reference.Concern; _reference != nil {
		reference.Concern = nil
		if _instance, ok := stage.Concerns_instance[_reference]; ok {
			reference.Concern = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *StakeholderShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.Stakeholder; _reference != nil {
		reference.Stakeholder = nil
		if _instance, ok := stage.Stakeholders_instance[_reference]; ok {
			reference.Stakeholder = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *SupportLevel) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.Tool; _reference != nil {
		reference.Tool = nil
		if _instance, ok := stage.Tools_instance[_reference]; ok {
			reference.Tool = _instance
		}
	}
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
	if analysisneed.LayoutDirection != analysisneedOther.LayoutDirection {
		diffs = append(diffs, analysisneed.GongMarshallField(stage, "LayoutDirection"))
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
	if concept.LayoutDirection != conceptOther.LayoutDirection {
		diffs = append(diffs, concept.GongMarshallField(stage, "LayoutDirection"))
	}
	ToolsDifferent := false
	if len(concept.Tools) != len(conceptOther.Tools) {
		ToolsDifferent = true
	} else {
		for i := range concept.Tools {
			if (concept.Tools[i] == nil) != (conceptOther.Tools[i] == nil) {
				ToolsDifferent = true
				break
			} else if concept.Tools[i] != nil && conceptOther.Tools[i] != nil {
				// this is a pointer comparaison
				if concept.Tools[i] != conceptOther.Tools[i] {
					ToolsDifferent = true
					break
				}
			}
		}
	}
	if ToolsDifferent {
		ops := Diff(stage, concept, conceptOther, "Tools", conceptOther.Tools, concept.Tools)
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
	if (conceptshape.Concept == nil) != (conceptshapeOther.Concept == nil) {
		diffs = append(diffs, conceptshape.GongMarshallField(stage, "Concept"))
	} else if conceptshape.Concept != nil && conceptshapeOther.Concept != nil {
		if conceptshape.Concept != conceptshapeOther.Concept {
			diffs = append(diffs, conceptshape.GongMarshallField(stage, "Concept"))
		}
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
	if concern.LayoutDirection != concernOther.LayoutDirection {
		diffs = append(diffs, concern.GongMarshallField(stage, "LayoutDirection"))
	}
	if concern.Description != concernOther.Description {
		diffs = append(diffs, concern.GongMarshallField(stage, "Description"))
	}
	SubConcernsDifferent := false
	if len(concern.SubConcerns) != len(concernOther.SubConcerns) {
		SubConcernsDifferent = true
	} else {
		for i := range concern.SubConcerns {
			if (concern.SubConcerns[i] == nil) != (concernOther.SubConcerns[i] == nil) {
				SubConcernsDifferent = true
				break
			} else if concern.SubConcerns[i] != nil && concernOther.SubConcerns[i] != nil {
				// this is a pointer comparaison
				if concern.SubConcerns[i] != concernOther.SubConcerns[i] {
					SubConcernsDifferent = true
					break
				}
			}
		}
	}
	if SubConcernsDifferent {
		ops := Diff(stage, concern, concernOther, "SubConcerns", concernOther.SubConcerns, concern.SubConcerns)
		diffs = append(diffs, ops)
	}
	InputsDifferent := false
	if len(concern.Inputs) != len(concernOther.Inputs) {
		InputsDifferent = true
	} else {
		for i := range concern.Inputs {
			if (concern.Inputs[i] == nil) != (concernOther.Inputs[i] == nil) {
				InputsDifferent = true
				break
			} else if concern.Inputs[i] != nil && concernOther.Inputs[i] != nil {
				// this is a pointer comparaison
				if concern.Inputs[i] != concernOther.Inputs[i] {
					InputsDifferent = true
					break
				}
			}
		}
	}
	if InputsDifferent {
		ops := Diff(stage, concern, concernOther, "Inputs", concernOther.Inputs, concern.Inputs)
		diffs = append(diffs, ops)
	}
	if concern.IsInputsNodeExpanded != concernOther.IsInputsNodeExpanded {
		diffs = append(diffs, concern.GongMarshallField(stage, "IsInputsNodeExpanded"))
	}
	OutputsDifferent := false
	if len(concern.Outputs) != len(concernOther.Outputs) {
		OutputsDifferent = true
	} else {
		for i := range concern.Outputs {
			if (concern.Outputs[i] == nil) != (concernOther.Outputs[i] == nil) {
				OutputsDifferent = true
				break
			} else if concern.Outputs[i] != nil && concernOther.Outputs[i] != nil {
				// this is a pointer comparaison
				if concern.Outputs[i] != concernOther.Outputs[i] {
					OutputsDifferent = true
					break
				}
			}
		}
	}
	if OutputsDifferent {
		ops := Diff(stage, concern, concernOther, "Outputs", concernOther.Outputs, concern.Outputs)
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
	RequirementsDifferent := false
	if len(concern.Requirements) != len(concernOther.Requirements) {
		RequirementsDifferent = true
	} else {
		for i := range concern.Requirements {
			if (concern.Requirements[i] == nil) != (concernOther.Requirements[i] == nil) {
				RequirementsDifferent = true
				break
			} else if concern.Requirements[i] != nil && concernOther.Requirements[i] != nil {
				// this is a pointer comparaison
				if concern.Requirements[i] != concernOther.Requirements[i] {
					RequirementsDifferent = true
					break
				}
			}
		}
	}
	if RequirementsDifferent {
		ops := Diff(stage, concern, concernOther, "Requirements", concernOther.Requirements, concern.Requirements)
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
	if (concerncompositionshape.Concern == nil) != (concerncompositionshapeOther.Concern == nil) {
		diffs = append(diffs, concerncompositionshape.GongMarshallField(stage, "Concern"))
	} else if concerncompositionshape.Concern != nil && concerncompositionshapeOther.Concern != nil {
		if concerncompositionshape.Concern != concerncompositionshapeOther.Concern {
			diffs = append(diffs, concerncompositionshape.GongMarshallField(stage, "Concern"))
		}
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

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (concerninputshape *ConcernInputShape) GongDiff(stage *Stage, concerninputshapeOther *ConcernInputShape) (diffs []string) {
	// insertion point for field diffs
	if concerninputshape.Name != concerninputshapeOther.Name {
		diffs = append(diffs, concerninputshape.GongMarshallField(stage, "Name"))
	}
	if (concerninputshape.Deliverable == nil) != (concerninputshapeOther.Deliverable == nil) {
		diffs = append(diffs, concerninputshape.GongMarshallField(stage, "Deliverable"))
	} else if concerninputshape.Deliverable != nil && concerninputshapeOther.Deliverable != nil {
		if concerninputshape.Deliverable != concerninputshapeOther.Deliverable {
			diffs = append(diffs, concerninputshape.GongMarshallField(stage, "Deliverable"))
		}
	}
	if (concerninputshape.Concern == nil) != (concerninputshapeOther.Concern == nil) {
		diffs = append(diffs, concerninputshape.GongMarshallField(stage, "Concern"))
	} else if concerninputshape.Concern != nil && concerninputshapeOther.Concern != nil {
		if concerninputshape.Concern != concerninputshapeOther.Concern {
			diffs = append(diffs, concerninputshape.GongMarshallField(stage, "Concern"))
		}
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

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (concernoutputshape *ConcernOutputShape) GongDiff(stage *Stage, concernoutputshapeOther *ConcernOutputShape) (diffs []string) {
	// insertion point for field diffs
	if concernoutputshape.Name != concernoutputshapeOther.Name {
		diffs = append(diffs, concernoutputshape.GongMarshallField(stage, "Name"))
	}
	if (concernoutputshape.Task == nil) != (concernoutputshapeOther.Task == nil) {
		diffs = append(diffs, concernoutputshape.GongMarshallField(stage, "Task"))
	} else if concernoutputshape.Task != nil && concernoutputshapeOther.Task != nil {
		if concernoutputshape.Task != concernoutputshapeOther.Task {
			diffs = append(diffs, concernoutputshape.GongMarshallField(stage, "Task"))
		}
	}
	if (concernoutputshape.Product == nil) != (concernoutputshapeOther.Product == nil) {
		diffs = append(diffs, concernoutputshape.GongMarshallField(stage, "Product"))
	} else if concernoutputshape.Product != nil && concernoutputshapeOther.Product != nil {
		if concernoutputshape.Product != concernoutputshapeOther.Product {
			diffs = append(diffs, concernoutputshape.GongMarshallField(stage, "Product"))
		}
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

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (concernshape *ConcernShape) GongDiff(stage *Stage, concernshapeOther *ConcernShape) (diffs []string) {
	// insertion point for field diffs
	if concernshape.Name != concernshapeOther.Name {
		diffs = append(diffs, concernshape.GongMarshallField(stage, "Name"))
	}
	if (concernshape.Concern == nil) != (concernshapeOther.Concern == nil) {
		diffs = append(diffs, concernshape.GongMarshallField(stage, "Concern"))
	} else if concernshape.Concern != nil && concernshapeOther.Concern != nil {
		if concernshape.Concern != concernshapeOther.Concern {
			diffs = append(diffs, concernshape.GongMarshallField(stage, "Concern"))
		}
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
	if deliverable.LayoutDirection != deliverableOther.LayoutDirection {
		diffs = append(diffs, deliverable.GongMarshallField(stage, "LayoutDirection"))
	}
	if deliverable.Description != deliverableOther.Description {
		diffs = append(diffs, deliverable.GongMarshallField(stage, "Description"))
	}
	SubProductsDifferent := false
	if len(deliverable.SubProducts) != len(deliverableOther.SubProducts) {
		SubProductsDifferent = true
	} else {
		for i := range deliverable.SubProducts {
			if (deliverable.SubProducts[i] == nil) != (deliverableOther.SubProducts[i] == nil) {
				SubProductsDifferent = true
				break
			} else if deliverable.SubProducts[i] != nil && deliverableOther.SubProducts[i] != nil {
				// this is a pointer comparaison
				if deliverable.SubProducts[i] != deliverableOther.SubProducts[i] {
					SubProductsDifferent = true
					break
				}
			}
		}
	}
	if SubProductsDifferent {
		ops := Diff(stage, deliverable, deliverableOther, "SubProducts", deliverableOther.SubProducts, deliverable.SubProducts)
		diffs = append(diffs, ops)
	}
	if deliverable.IsProducersNodeExpanded != deliverableOther.IsProducersNodeExpanded {
		diffs = append(diffs, deliverable.GongMarshallField(stage, "IsProducersNodeExpanded"))
	}
	if deliverable.IsConsumersNodeExpanded != deliverableOther.IsConsumersNodeExpanded {
		diffs = append(diffs, deliverable.GongMarshallField(stage, "IsConsumersNodeExpanded"))
	}
	ConceptsDifferent := false
	if len(deliverable.Concepts) != len(deliverableOther.Concepts) {
		ConceptsDifferent = true
	} else {
		for i := range deliverable.Concepts {
			if (deliverable.Concepts[i] == nil) != (deliverableOther.Concepts[i] == nil) {
				ConceptsDifferent = true
				break
			} else if deliverable.Concepts[i] != nil && deliverableOther.Concepts[i] != nil {
				// this is a pointer comparaison
				if deliverable.Concepts[i] != deliverableOther.Concepts[i] {
					ConceptsDifferent = true
					break
				}
			}
		}
	}
	if ConceptsDifferent {
		ops := Diff(stage, deliverable, deliverableOther, "Concepts", deliverableOther.Concepts, deliverable.Concepts)
		diffs = append(diffs, ops)
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
	if diagram.LayoutDirection != diagramOther.LayoutDirection {
		diffs = append(diffs, diagram.GongMarshallField(stage, "LayoutDirection"))
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
	ConcernsWhoseRequirementsNodeIsExpandedDifferent := false
	if len(diagram.ConcernsWhoseRequirementsNodeIsExpanded) != len(diagramOther.ConcernsWhoseRequirementsNodeIsExpanded) {
		ConcernsWhoseRequirementsNodeIsExpandedDifferent = true
	} else {
		for i := range diagram.ConcernsWhoseRequirementsNodeIsExpanded {
			if (diagram.ConcernsWhoseRequirementsNodeIsExpanded[i] == nil) != (diagramOther.ConcernsWhoseRequirementsNodeIsExpanded[i] == nil) {
				ConcernsWhoseRequirementsNodeIsExpandedDifferent = true
				break
			} else if diagram.ConcernsWhoseRequirementsNodeIsExpanded[i] != nil && diagramOther.ConcernsWhoseRequirementsNodeIsExpanded[i] != nil {
				// this is a pointer comparaison
				if diagram.ConcernsWhoseRequirementsNodeIsExpanded[i] != diagramOther.ConcernsWhoseRequirementsNodeIsExpanded[i] {
					ConcernsWhoseRequirementsNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if ConcernsWhoseRequirementsNodeIsExpandedDifferent {
		ops := Diff(stage, diagram, diagramOther, "ConcernsWhoseRequirementsNodeIsExpanded", diagramOther.ConcernsWhoseRequirementsNodeIsExpanded, diagram.ConcernsWhoseRequirementsNodeIsExpanded)
		diffs = append(diffs, ops)
	}
	if diagram.IsRequirementsNodeExpanded != diagramOther.IsRequirementsNodeExpanded {
		diffs = append(diffs, diagram.GongMarshallField(stage, "IsRequirementsNodeExpanded"))
	}
	if diagram.IsConceptsNodeExpanded != diagramOther.IsConceptsNodeExpanded {
		diffs = append(diffs, diagram.GongMarshallField(stage, "IsConceptsNodeExpanded"))
	}
	Product_ShapesDifferent := false
	if len(diagram.Product_Shapes) != len(diagramOther.Product_Shapes) {
		Product_ShapesDifferent = true
	} else {
		for i := range diagram.Product_Shapes {
			if (diagram.Product_Shapes[i] == nil) != (diagramOther.Product_Shapes[i] == nil) {
				Product_ShapesDifferent = true
				break
			} else if diagram.Product_Shapes[i] != nil && diagramOther.Product_Shapes[i] != nil {
				// this is a pointer comparaison
				if diagram.Product_Shapes[i] != diagramOther.Product_Shapes[i] {
					Product_ShapesDifferent = true
					break
				}
			}
		}
	}
	if Product_ShapesDifferent {
		ops := Diff(stage, diagram, diagramOther, "Product_Shapes", diagramOther.Product_Shapes, diagram.Product_Shapes)
		diffs = append(diffs, ops)
	}
	ProductsWhoseNodeIsExpandedDifferent := false
	if len(diagram.ProductsWhoseNodeIsExpanded) != len(diagramOther.ProductsWhoseNodeIsExpanded) {
		ProductsWhoseNodeIsExpandedDifferent = true
	} else {
		for i := range diagram.ProductsWhoseNodeIsExpanded {
			if (diagram.ProductsWhoseNodeIsExpanded[i] == nil) != (diagramOther.ProductsWhoseNodeIsExpanded[i] == nil) {
				ProductsWhoseNodeIsExpandedDifferent = true
				break
			} else if diagram.ProductsWhoseNodeIsExpanded[i] != nil && diagramOther.ProductsWhoseNodeIsExpanded[i] != nil {
				// this is a pointer comparaison
				if diagram.ProductsWhoseNodeIsExpanded[i] != diagramOther.ProductsWhoseNodeIsExpanded[i] {
					ProductsWhoseNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if ProductsWhoseNodeIsExpandedDifferent {
		ops := Diff(stage, diagram, diagramOther, "ProductsWhoseNodeIsExpanded", diagramOther.ProductsWhoseNodeIsExpanded, diagram.ProductsWhoseNodeIsExpanded)
		diffs = append(diffs, ops)
	}
	if diagram.IsPBSNodeExpanded != diagramOther.IsPBSNodeExpanded {
		diffs = append(diffs, diagram.GongMarshallField(stage, "IsPBSNodeExpanded"))
	}
	ProductComposition_ShapesDifferent := false
	if len(diagram.ProductComposition_Shapes) != len(diagramOther.ProductComposition_Shapes) {
		ProductComposition_ShapesDifferent = true
	} else {
		for i := range diagram.ProductComposition_Shapes {
			if (diagram.ProductComposition_Shapes[i] == nil) != (diagramOther.ProductComposition_Shapes[i] == nil) {
				ProductComposition_ShapesDifferent = true
				break
			} else if diagram.ProductComposition_Shapes[i] != nil && diagramOther.ProductComposition_Shapes[i] != nil {
				// this is a pointer comparaison
				if diagram.ProductComposition_Shapes[i] != diagramOther.ProductComposition_Shapes[i] {
					ProductComposition_ShapesDifferent = true
					break
				}
			}
		}
	}
	if ProductComposition_ShapesDifferent {
		ops := Diff(stage, diagram, diagramOther, "ProductComposition_Shapes", diagramOther.ProductComposition_Shapes, diagram.ProductComposition_Shapes)
		diffs = append(diffs, ops)
	}
	if diagram.IsConcernsNodeExpanded != diagramOther.IsConcernsNodeExpanded {
		diffs = append(diffs, diagram.GongMarshallField(stage, "IsConcernsNodeExpanded"))
	}
	Concern_ShapesDifferent := false
	if len(diagram.Concern_Shapes) != len(diagramOther.Concern_Shapes) {
		Concern_ShapesDifferent = true
	} else {
		for i := range diagram.Concern_Shapes {
			if (diagram.Concern_Shapes[i] == nil) != (diagramOther.Concern_Shapes[i] == nil) {
				Concern_ShapesDifferent = true
				break
			} else if diagram.Concern_Shapes[i] != nil && diagramOther.Concern_Shapes[i] != nil {
				// this is a pointer comparaison
				if diagram.Concern_Shapes[i] != diagramOther.Concern_Shapes[i] {
					Concern_ShapesDifferent = true
					break
				}
			}
		}
	}
	if Concern_ShapesDifferent {
		ops := Diff(stage, diagram, diagramOther, "Concern_Shapes", diagramOther.Concern_Shapes, diagram.Concern_Shapes)
		diffs = append(diffs, ops)
	}
	ConcernsWhoseNodeIsExpandedDifferent := false
	if len(diagram.ConcernsWhoseNodeIsExpanded) != len(diagramOther.ConcernsWhoseNodeIsExpanded) {
		ConcernsWhoseNodeIsExpandedDifferent = true
	} else {
		for i := range diagram.ConcernsWhoseNodeIsExpanded {
			if (diagram.ConcernsWhoseNodeIsExpanded[i] == nil) != (diagramOther.ConcernsWhoseNodeIsExpanded[i] == nil) {
				ConcernsWhoseNodeIsExpandedDifferent = true
				break
			} else if diagram.ConcernsWhoseNodeIsExpanded[i] != nil && diagramOther.ConcernsWhoseNodeIsExpanded[i] != nil {
				// this is a pointer comparaison
				if diagram.ConcernsWhoseNodeIsExpanded[i] != diagramOther.ConcernsWhoseNodeIsExpanded[i] {
					ConcernsWhoseNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if ConcernsWhoseNodeIsExpandedDifferent {
		ops := Diff(stage, diagram, diagramOther, "ConcernsWhoseNodeIsExpanded", diagramOther.ConcernsWhoseNodeIsExpanded, diagram.ConcernsWhoseNodeIsExpanded)
		diffs = append(diffs, ops)
	}
	ConcernsWhoseInputNodeIsExpandedDifferent := false
	if len(diagram.ConcernsWhoseInputNodeIsExpanded) != len(diagramOther.ConcernsWhoseInputNodeIsExpanded) {
		ConcernsWhoseInputNodeIsExpandedDifferent = true
	} else {
		for i := range diagram.ConcernsWhoseInputNodeIsExpanded {
			if (diagram.ConcernsWhoseInputNodeIsExpanded[i] == nil) != (diagramOther.ConcernsWhoseInputNodeIsExpanded[i] == nil) {
				ConcernsWhoseInputNodeIsExpandedDifferent = true
				break
			} else if diagram.ConcernsWhoseInputNodeIsExpanded[i] != nil && diagramOther.ConcernsWhoseInputNodeIsExpanded[i] != nil {
				// this is a pointer comparaison
				if diagram.ConcernsWhoseInputNodeIsExpanded[i] != diagramOther.ConcernsWhoseInputNodeIsExpanded[i] {
					ConcernsWhoseInputNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if ConcernsWhoseInputNodeIsExpandedDifferent {
		ops := Diff(stage, diagram, diagramOther, "ConcernsWhoseInputNodeIsExpanded", diagramOther.ConcernsWhoseInputNodeIsExpanded, diagram.ConcernsWhoseInputNodeIsExpanded)
		diffs = append(diffs, ops)
	}
	ConcernsWhoseStakeholderNodeIsExpandedDifferent := false
	if len(diagram.ConcernsWhoseStakeholderNodeIsExpanded) != len(diagramOther.ConcernsWhoseStakeholderNodeIsExpanded) {
		ConcernsWhoseStakeholderNodeIsExpandedDifferent = true
	} else {
		for i := range diagram.ConcernsWhoseStakeholderNodeIsExpanded {
			if (diagram.ConcernsWhoseStakeholderNodeIsExpanded[i] == nil) != (diagramOther.ConcernsWhoseStakeholderNodeIsExpanded[i] == nil) {
				ConcernsWhoseStakeholderNodeIsExpandedDifferent = true
				break
			} else if diagram.ConcernsWhoseStakeholderNodeIsExpanded[i] != nil && diagramOther.ConcernsWhoseStakeholderNodeIsExpanded[i] != nil {
				// this is a pointer comparaison
				if diagram.ConcernsWhoseStakeholderNodeIsExpanded[i] != diagramOther.ConcernsWhoseStakeholderNodeIsExpanded[i] {
					ConcernsWhoseStakeholderNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if ConcernsWhoseStakeholderNodeIsExpandedDifferent {
		ops := Diff(stage, diagram, diagramOther, "ConcernsWhoseStakeholderNodeIsExpanded", diagramOther.ConcernsWhoseStakeholderNodeIsExpanded, diagram.ConcernsWhoseStakeholderNodeIsExpanded)
		diffs = append(diffs, ops)
	}
	ConcernssWhoseOutputNodeIsExpandedDifferent := false
	if len(diagram.ConcernssWhoseOutputNodeIsExpanded) != len(diagramOther.ConcernssWhoseOutputNodeIsExpanded) {
		ConcernssWhoseOutputNodeIsExpandedDifferent = true
	} else {
		for i := range diagram.ConcernssWhoseOutputNodeIsExpanded {
			if (diagram.ConcernssWhoseOutputNodeIsExpanded[i] == nil) != (diagramOther.ConcernssWhoseOutputNodeIsExpanded[i] == nil) {
				ConcernssWhoseOutputNodeIsExpandedDifferent = true
				break
			} else if diagram.ConcernssWhoseOutputNodeIsExpanded[i] != nil && diagramOther.ConcernssWhoseOutputNodeIsExpanded[i] != nil {
				// this is a pointer comparaison
				if diagram.ConcernssWhoseOutputNodeIsExpanded[i] != diagramOther.ConcernssWhoseOutputNodeIsExpanded[i] {
					ConcernssWhoseOutputNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if ConcernssWhoseOutputNodeIsExpandedDifferent {
		ops := Diff(stage, diagram, diagramOther, "ConcernssWhoseOutputNodeIsExpanded", diagramOther.ConcernssWhoseOutputNodeIsExpanded, diagram.ConcernssWhoseOutputNodeIsExpanded)
		diffs = append(diffs, ops)
	}
	ConcernComposition_ShapesDifferent := false
	if len(diagram.ConcernComposition_Shapes) != len(diagramOther.ConcernComposition_Shapes) {
		ConcernComposition_ShapesDifferent = true
	} else {
		for i := range diagram.ConcernComposition_Shapes {
			if (diagram.ConcernComposition_Shapes[i] == nil) != (diagramOther.ConcernComposition_Shapes[i] == nil) {
				ConcernComposition_ShapesDifferent = true
				break
			} else if diagram.ConcernComposition_Shapes[i] != nil && diagramOther.ConcernComposition_Shapes[i] != nil {
				// this is a pointer comparaison
				if diagram.ConcernComposition_Shapes[i] != diagramOther.ConcernComposition_Shapes[i] {
					ConcernComposition_ShapesDifferent = true
					break
				}
			}
		}
	}
	if ConcernComposition_ShapesDifferent {
		ops := Diff(stage, diagram, diagramOther, "ConcernComposition_Shapes", diagramOther.ConcernComposition_Shapes, diagram.ConcernComposition_Shapes)
		diffs = append(diffs, ops)
	}
	ConcernInputShapesDifferent := false
	if len(diagram.ConcernInputShapes) != len(diagramOther.ConcernInputShapes) {
		ConcernInputShapesDifferent = true
	} else {
		for i := range diagram.ConcernInputShapes {
			if (diagram.ConcernInputShapes[i] == nil) != (diagramOther.ConcernInputShapes[i] == nil) {
				ConcernInputShapesDifferent = true
				break
			} else if diagram.ConcernInputShapes[i] != nil && diagramOther.ConcernInputShapes[i] != nil {
				// this is a pointer comparaison
				if diagram.ConcernInputShapes[i] != diagramOther.ConcernInputShapes[i] {
					ConcernInputShapesDifferent = true
					break
				}
			}
		}
	}
	if ConcernInputShapesDifferent {
		ops := Diff(stage, diagram, diagramOther, "ConcernInputShapes", diagramOther.ConcernInputShapes, diagram.ConcernInputShapes)
		diffs = append(diffs, ops)
	}
	ConcernOutputShapesDifferent := false
	if len(diagram.ConcernOutputShapes) != len(diagramOther.ConcernOutputShapes) {
		ConcernOutputShapesDifferent = true
	} else {
		for i := range diagram.ConcernOutputShapes {
			if (diagram.ConcernOutputShapes[i] == nil) != (diagramOther.ConcernOutputShapes[i] == nil) {
				ConcernOutputShapesDifferent = true
				break
			} else if diagram.ConcernOutputShapes[i] != nil && diagramOther.ConcernOutputShapes[i] != nil {
				// this is a pointer comparaison
				if diagram.ConcernOutputShapes[i] != diagramOther.ConcernOutputShapes[i] {
					ConcernOutputShapesDifferent = true
					break
				}
			}
		}
	}
	if ConcernOutputShapesDifferent {
		ops := Diff(stage, diagram, diagramOther, "ConcernOutputShapes", diagramOther.ConcernOutputShapes, diagram.ConcernOutputShapes)
		diffs = append(diffs, ops)
	}
	Note_ShapesDifferent := false
	if len(diagram.Note_Shapes) != len(diagramOther.Note_Shapes) {
		Note_ShapesDifferent = true
	} else {
		for i := range diagram.Note_Shapes {
			if (diagram.Note_Shapes[i] == nil) != (diagramOther.Note_Shapes[i] == nil) {
				Note_ShapesDifferent = true
				break
			} else if diagram.Note_Shapes[i] != nil && diagramOther.Note_Shapes[i] != nil {
				// this is a pointer comparaison
				if diagram.Note_Shapes[i] != diagramOther.Note_Shapes[i] {
					Note_ShapesDifferent = true
					break
				}
			}
		}
	}
	if Note_ShapesDifferent {
		ops := Diff(stage, diagram, diagramOther, "Note_Shapes", diagramOther.Note_Shapes, diagram.Note_Shapes)
		diffs = append(diffs, ops)
	}
	NotesWhoseNodeIsExpandedDifferent := false
	if len(diagram.NotesWhoseNodeIsExpanded) != len(diagramOther.NotesWhoseNodeIsExpanded) {
		NotesWhoseNodeIsExpandedDifferent = true
	} else {
		for i := range diagram.NotesWhoseNodeIsExpanded {
			if (diagram.NotesWhoseNodeIsExpanded[i] == nil) != (diagramOther.NotesWhoseNodeIsExpanded[i] == nil) {
				NotesWhoseNodeIsExpandedDifferent = true
				break
			} else if diagram.NotesWhoseNodeIsExpanded[i] != nil && diagramOther.NotesWhoseNodeIsExpanded[i] != nil {
				// this is a pointer comparaison
				if diagram.NotesWhoseNodeIsExpanded[i] != diagramOther.NotesWhoseNodeIsExpanded[i] {
					NotesWhoseNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if NotesWhoseNodeIsExpandedDifferent {
		ops := Diff(stage, diagram, diagramOther, "NotesWhoseNodeIsExpanded", diagramOther.NotesWhoseNodeIsExpanded, diagram.NotesWhoseNodeIsExpanded)
		diffs = append(diffs, ops)
	}
	if diagram.IsNotesNodeExpanded != diagramOther.IsNotesNodeExpanded {
		diffs = append(diffs, diagram.GongMarshallField(stage, "IsNotesNodeExpanded"))
	}
	NoteProductShapesDifferent := false
	if len(diagram.NoteProductShapes) != len(diagramOther.NoteProductShapes) {
		NoteProductShapesDifferent = true
	} else {
		for i := range diagram.NoteProductShapes {
			if (diagram.NoteProductShapes[i] == nil) != (diagramOther.NoteProductShapes[i] == nil) {
				NoteProductShapesDifferent = true
				break
			} else if diagram.NoteProductShapes[i] != nil && diagramOther.NoteProductShapes[i] != nil {
				// this is a pointer comparaison
				if diagram.NoteProductShapes[i] != diagramOther.NoteProductShapes[i] {
					NoteProductShapesDifferent = true
					break
				}
			}
		}
	}
	if NoteProductShapesDifferent {
		ops := Diff(stage, diagram, diagramOther, "NoteProductShapes", diagramOther.NoteProductShapes, diagram.NoteProductShapes)
		diffs = append(diffs, ops)
	}
	NoteTaskShapesDifferent := false
	if len(diagram.NoteTaskShapes) != len(diagramOther.NoteTaskShapes) {
		NoteTaskShapesDifferent = true
	} else {
		for i := range diagram.NoteTaskShapes {
			if (diagram.NoteTaskShapes[i] == nil) != (diagramOther.NoteTaskShapes[i] == nil) {
				NoteTaskShapesDifferent = true
				break
			} else if diagram.NoteTaskShapes[i] != nil && diagramOther.NoteTaskShapes[i] != nil {
				// this is a pointer comparaison
				if diagram.NoteTaskShapes[i] != diagramOther.NoteTaskShapes[i] {
					NoteTaskShapesDifferent = true
					break
				}
			}
		}
	}
	if NoteTaskShapesDifferent {
		ops := Diff(stage, diagram, diagramOther, "NoteTaskShapes", diagramOther.NoteTaskShapes, diagram.NoteTaskShapes)
		diffs = append(diffs, ops)
	}
	NoteResourceShapesDifferent := false
	if len(diagram.NoteResourceShapes) != len(diagramOther.NoteResourceShapes) {
		NoteResourceShapesDifferent = true
	} else {
		for i := range diagram.NoteResourceShapes {
			if (diagram.NoteResourceShapes[i] == nil) != (diagramOther.NoteResourceShapes[i] == nil) {
				NoteResourceShapesDifferent = true
				break
			} else if diagram.NoteResourceShapes[i] != nil && diagramOther.NoteResourceShapes[i] != nil {
				// this is a pointer comparaison
				if diagram.NoteResourceShapes[i] != diagramOther.NoteResourceShapes[i] {
					NoteResourceShapesDifferent = true
					break
				}
			}
		}
	}
	if NoteResourceShapesDifferent {
		ops := Diff(stage, diagram, diagramOther, "NoteResourceShapes", diagramOther.NoteResourceShapes, diagram.NoteResourceShapes)
		diffs = append(diffs, ops)
	}
	Stakeholder_ShapesDifferent := false
	if len(diagram.Stakeholder_Shapes) != len(diagramOther.Stakeholder_Shapes) {
		Stakeholder_ShapesDifferent = true
	} else {
		for i := range diagram.Stakeholder_Shapes {
			if (diagram.Stakeholder_Shapes[i] == nil) != (diagramOther.Stakeholder_Shapes[i] == nil) {
				Stakeholder_ShapesDifferent = true
				break
			} else if diagram.Stakeholder_Shapes[i] != nil && diagramOther.Stakeholder_Shapes[i] != nil {
				// this is a pointer comparaison
				if diagram.Stakeholder_Shapes[i] != diagramOther.Stakeholder_Shapes[i] {
					Stakeholder_ShapesDifferent = true
					break
				}
			}
		}
	}
	if Stakeholder_ShapesDifferent {
		ops := Diff(stage, diagram, diagramOther, "Stakeholder_Shapes", diagramOther.Stakeholder_Shapes, diagram.Stakeholder_Shapes)
		diffs = append(diffs, ops)
	}
	ResourcesWhoseNodeIsExpandedDifferent := false
	if len(diagram.ResourcesWhoseNodeIsExpanded) != len(diagramOther.ResourcesWhoseNodeIsExpanded) {
		ResourcesWhoseNodeIsExpandedDifferent = true
	} else {
		for i := range diagram.ResourcesWhoseNodeIsExpanded {
			if (diagram.ResourcesWhoseNodeIsExpanded[i] == nil) != (diagramOther.ResourcesWhoseNodeIsExpanded[i] == nil) {
				ResourcesWhoseNodeIsExpandedDifferent = true
				break
			} else if diagram.ResourcesWhoseNodeIsExpanded[i] != nil && diagramOther.ResourcesWhoseNodeIsExpanded[i] != nil {
				// this is a pointer comparaison
				if diagram.ResourcesWhoseNodeIsExpanded[i] != diagramOther.ResourcesWhoseNodeIsExpanded[i] {
					ResourcesWhoseNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if ResourcesWhoseNodeIsExpandedDifferent {
		ops := Diff(stage, diagram, diagramOther, "ResourcesWhoseNodeIsExpanded", diagramOther.ResourcesWhoseNodeIsExpanded, diagram.ResourcesWhoseNodeIsExpanded)
		diffs = append(diffs, ops)
	}
	if diagram.IsStakeholdersNodeExpanded != diagramOther.IsStakeholdersNodeExpanded {
		diffs = append(diffs, diagram.GongMarshallField(stage, "IsStakeholdersNodeExpanded"))
	}
	ResourceComposition_ShapesDifferent := false
	if len(diagram.ResourceComposition_Shapes) != len(diagramOther.ResourceComposition_Shapes) {
		ResourceComposition_ShapesDifferent = true
	} else {
		for i := range diagram.ResourceComposition_Shapes {
			if (diagram.ResourceComposition_Shapes[i] == nil) != (diagramOther.ResourceComposition_Shapes[i] == nil) {
				ResourceComposition_ShapesDifferent = true
				break
			} else if diagram.ResourceComposition_Shapes[i] != nil && diagramOther.ResourceComposition_Shapes[i] != nil {
				// this is a pointer comparaison
				if diagram.ResourceComposition_Shapes[i] != diagramOther.ResourceComposition_Shapes[i] {
					ResourceComposition_ShapesDifferent = true
					break
				}
			}
		}
	}
	if ResourceComposition_ShapesDifferent {
		ops := Diff(stage, diagram, diagramOther, "ResourceComposition_Shapes", diagramOther.ResourceComposition_Shapes, diagram.ResourceComposition_Shapes)
		diffs = append(diffs, ops)
	}
	StakeholderConcernShapesDifferent := false
	if len(diagram.StakeholderConcernShapes) != len(diagramOther.StakeholderConcernShapes) {
		StakeholderConcernShapesDifferent = true
	} else {
		for i := range diagram.StakeholderConcernShapes {
			if (diagram.StakeholderConcernShapes[i] == nil) != (diagramOther.StakeholderConcernShapes[i] == nil) {
				StakeholderConcernShapesDifferent = true
				break
			} else if diagram.StakeholderConcernShapes[i] != nil && diagramOther.StakeholderConcernShapes[i] != nil {
				// this is a pointer comparaison
				if diagram.StakeholderConcernShapes[i] != diagramOther.StakeholderConcernShapes[i] {
					StakeholderConcernShapesDifferent = true
					break
				}
			}
		}
	}
	if StakeholderConcernShapesDifferent {
		ops := Diff(stage, diagram, diagramOther, "StakeholderConcernShapes", diagramOther.StakeholderConcernShapes, diagram.StakeholderConcernShapes)
		diffs = append(diffs, ops)
	}
	Requirement_ShapesDifferent := false
	if len(diagram.Requirement_Shapes) != len(diagramOther.Requirement_Shapes) {
		Requirement_ShapesDifferent = true
	} else {
		for i := range diagram.Requirement_Shapes {
			if (diagram.Requirement_Shapes[i] == nil) != (diagramOther.Requirement_Shapes[i] == nil) {
				Requirement_ShapesDifferent = true
				break
			} else if diagram.Requirement_Shapes[i] != nil && diagramOther.Requirement_Shapes[i] != nil {
				// this is a pointer comparaison
				if diagram.Requirement_Shapes[i] != diagramOther.Requirement_Shapes[i] {
					Requirement_ShapesDifferent = true
					break
				}
			}
		}
	}
	if Requirement_ShapesDifferent {
		ops := Diff(stage, diagram, diagramOther, "Requirement_Shapes", diagramOther.Requirement_Shapes, diagram.Requirement_Shapes)
		diffs = append(diffs, ops)
	}
	RequirementsWhoseNodeIsExpandedDifferent := false
	if len(diagram.RequirementsWhoseNodeIsExpanded) != len(diagramOther.RequirementsWhoseNodeIsExpanded) {
		RequirementsWhoseNodeIsExpandedDifferent = true
	} else {
		for i := range diagram.RequirementsWhoseNodeIsExpanded {
			if (diagram.RequirementsWhoseNodeIsExpanded[i] == nil) != (diagramOther.RequirementsWhoseNodeIsExpanded[i] == nil) {
				RequirementsWhoseNodeIsExpandedDifferent = true
				break
			} else if diagram.RequirementsWhoseNodeIsExpanded[i] != nil && diagramOther.RequirementsWhoseNodeIsExpanded[i] != nil {
				// this is a pointer comparaison
				if diagram.RequirementsWhoseNodeIsExpanded[i] != diagramOther.RequirementsWhoseNodeIsExpanded[i] {
					RequirementsWhoseNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if RequirementsWhoseNodeIsExpandedDifferent {
		ops := Diff(stage, diagram, diagramOther, "RequirementsWhoseNodeIsExpanded", diagramOther.RequirementsWhoseNodeIsExpanded, diagram.RequirementsWhoseNodeIsExpanded)
		diffs = append(diffs, ops)
	}
	Concept_ShapesDifferent := false
	if len(diagram.Concept_Shapes) != len(diagramOther.Concept_Shapes) {
		Concept_ShapesDifferent = true
	} else {
		for i := range diagram.Concept_Shapes {
			if (diagram.Concept_Shapes[i] == nil) != (diagramOther.Concept_Shapes[i] == nil) {
				Concept_ShapesDifferent = true
				break
			} else if diagram.Concept_Shapes[i] != nil && diagramOther.Concept_Shapes[i] != nil {
				// this is a pointer comparaison
				if diagram.Concept_Shapes[i] != diagramOther.Concept_Shapes[i] {
					Concept_ShapesDifferent = true
					break
				}
			}
		}
	}
	if Concept_ShapesDifferent {
		ops := Diff(stage, diagram, diagramOther, "Concept_Shapes", diagramOther.Concept_Shapes, diagram.Concept_Shapes)
		diffs = append(diffs, ops)
	}
	ConceptsWhoseNodeIsExpandedDifferent := false
	if len(diagram.ConceptsWhoseNodeIsExpanded) != len(diagramOther.ConceptsWhoseNodeIsExpanded) {
		ConceptsWhoseNodeIsExpandedDifferent = true
	} else {
		for i := range diagram.ConceptsWhoseNodeIsExpanded {
			if (diagram.ConceptsWhoseNodeIsExpanded[i] == nil) != (diagramOther.ConceptsWhoseNodeIsExpanded[i] == nil) {
				ConceptsWhoseNodeIsExpandedDifferent = true
				break
			} else if diagram.ConceptsWhoseNodeIsExpanded[i] != nil && diagramOther.ConceptsWhoseNodeIsExpanded[i] != nil {
				// this is a pointer comparaison
				if diagram.ConceptsWhoseNodeIsExpanded[i] != diagramOther.ConceptsWhoseNodeIsExpanded[i] {
					ConceptsWhoseNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if ConceptsWhoseNodeIsExpandedDifferent {
		ops := Diff(stage, diagram, diagramOther, "ConceptsWhoseNodeIsExpanded", diagramOther.ConceptsWhoseNodeIsExpanded, diagram.ConceptsWhoseNodeIsExpanded)
		diffs = append(diffs, ops)
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
	if library.LayoutDirection != libraryOther.LayoutDirection {
		diffs = append(diffs, library.GongMarshallField(stage, "LayoutDirection"))
	}
	RootDeliverablesDifferent := false
	if len(library.RootDeliverables) != len(libraryOther.RootDeliverables) {
		RootDeliverablesDifferent = true
	} else {
		for i := range library.RootDeliverables {
			if (library.RootDeliverables[i] == nil) != (libraryOther.RootDeliverables[i] == nil) {
				RootDeliverablesDifferent = true
				break
			} else if library.RootDeliverables[i] != nil && libraryOther.RootDeliverables[i] != nil {
				// this is a pointer comparaison
				if library.RootDeliverables[i] != libraryOther.RootDeliverables[i] {
					RootDeliverablesDifferent = true
					break
				}
			}
		}
	}
	if RootDeliverablesDifferent {
		ops := Diff(stage, library, libraryOther, "RootDeliverables", libraryOther.RootDeliverables, library.RootDeliverables)
		diffs = append(diffs, ops)
	}
	RootConcernsDifferent := false
	if len(library.RootConcerns) != len(libraryOther.RootConcerns) {
		RootConcernsDifferent = true
	} else {
		for i := range library.RootConcerns {
			if (library.RootConcerns[i] == nil) != (libraryOther.RootConcerns[i] == nil) {
				RootConcernsDifferent = true
				break
			} else if library.RootConcerns[i] != nil && libraryOther.RootConcerns[i] != nil {
				// this is a pointer comparaison
				if library.RootConcerns[i] != libraryOther.RootConcerns[i] {
					RootConcernsDifferent = true
					break
				}
			}
		}
	}
	if RootConcernsDifferent {
		ops := Diff(stage, library, libraryOther, "RootConcerns", libraryOther.RootConcerns, library.RootConcerns)
		diffs = append(diffs, ops)
	}
	RootStakeholdersDifferent := false
	if len(library.RootStakeholders) != len(libraryOther.RootStakeholders) {
		RootStakeholdersDifferent = true
	} else {
		for i := range library.RootStakeholders {
			if (library.RootStakeholders[i] == nil) != (libraryOther.RootStakeholders[i] == nil) {
				RootStakeholdersDifferent = true
				break
			} else if library.RootStakeholders[i] != nil && libraryOther.RootStakeholders[i] != nil {
				// this is a pointer comparaison
				if library.RootStakeholders[i] != libraryOther.RootStakeholders[i] {
					RootStakeholdersDifferent = true
					break
				}
			}
		}
	}
	if RootStakeholdersDifferent {
		ops := Diff(stage, library, libraryOther, "RootStakeholders", libraryOther.RootStakeholders, library.RootStakeholders)
		diffs = append(diffs, ops)
	}
	RootRequirementsDifferent := false
	if len(library.RootRequirements) != len(libraryOther.RootRequirements) {
		RootRequirementsDifferent = true
	} else {
		for i := range library.RootRequirements {
			if (library.RootRequirements[i] == nil) != (libraryOther.RootRequirements[i] == nil) {
				RootRequirementsDifferent = true
				break
			} else if library.RootRequirements[i] != nil && libraryOther.RootRequirements[i] != nil {
				// this is a pointer comparaison
				if library.RootRequirements[i] != libraryOther.RootRequirements[i] {
					RootRequirementsDifferent = true
					break
				}
			}
		}
	}
	if RootRequirementsDifferent {
		ops := Diff(stage, library, libraryOther, "RootRequirements", libraryOther.RootRequirements, library.RootRequirements)
		diffs = append(diffs, ops)
	}
	RootConceptsDifferent := false
	if len(library.RootConcepts) != len(libraryOther.RootConcepts) {
		RootConceptsDifferent = true
	} else {
		for i := range library.RootConcepts {
			if (library.RootConcepts[i] == nil) != (libraryOther.RootConcepts[i] == nil) {
				RootConceptsDifferent = true
				break
			} else if library.RootConcepts[i] != nil && libraryOther.RootConcepts[i] != nil {
				// this is a pointer comparaison
				if library.RootConcepts[i] != libraryOther.RootConcepts[i] {
					RootConceptsDifferent = true
					break
				}
			}
		}
	}
	if RootConceptsDifferent {
		ops := Diff(stage, library, libraryOther, "RootConcepts", libraryOther.RootConcepts, library.RootConcepts)
		diffs = append(diffs, ops)
	}
	AnalysisNeedsDifferent := false
	if len(library.AnalysisNeeds) != len(libraryOther.AnalysisNeeds) {
		AnalysisNeedsDifferent = true
	} else {
		for i := range library.AnalysisNeeds {
			if (library.AnalysisNeeds[i] == nil) != (libraryOther.AnalysisNeeds[i] == nil) {
				AnalysisNeedsDifferent = true
				break
			} else if library.AnalysisNeeds[i] != nil && libraryOther.AnalysisNeeds[i] != nil {
				// this is a pointer comparaison
				if library.AnalysisNeeds[i] != libraryOther.AnalysisNeeds[i] {
					AnalysisNeedsDifferent = true
					break
				}
			}
		}
	}
	if AnalysisNeedsDifferent {
		ops := Diff(stage, library, libraryOther, "AnalysisNeeds", libraryOther.AnalysisNeeds, library.AnalysisNeeds)
		diffs = append(diffs, ops)
	}
	NotesDifferent := false
	if len(library.Notes) != len(libraryOther.Notes) {
		NotesDifferent = true
	} else {
		for i := range library.Notes {
			if (library.Notes[i] == nil) != (libraryOther.Notes[i] == nil) {
				NotesDifferent = true
				break
			} else if library.Notes[i] != nil && libraryOther.Notes[i] != nil {
				// this is a pointer comparaison
				if library.Notes[i] != libraryOther.Notes[i] {
					NotesDifferent = true
					break
				}
			}
		}
	}
	if NotesDifferent {
		ops := Diff(stage, library, libraryOther, "Notes", libraryOther.Notes, library.Notes)
		diffs = append(diffs, ops)
	}
	DiagramsDifferent := false
	if len(library.Diagrams) != len(libraryOther.Diagrams) {
		DiagramsDifferent = true
	} else {
		for i := range library.Diagrams {
			if (library.Diagrams[i] == nil) != (libraryOther.Diagrams[i] == nil) {
				DiagramsDifferent = true
				break
			} else if library.Diagrams[i] != nil && libraryOther.Diagrams[i] != nil {
				// this is a pointer comparaison
				if library.Diagrams[i] != libraryOther.Diagrams[i] {
					DiagramsDifferent = true
					break
				}
			}
		}
	}
	if DiagramsDifferent {
		ops := Diff(stage, library, libraryOther, "Diagrams", libraryOther.Diagrams, library.Diagrams)
		diffs = append(diffs, ops)
	}
	SubLibrariesDifferent := false
	if len(library.SubLibraries) != len(libraryOther.SubLibraries) {
		SubLibrariesDifferent = true
	} else {
		for i := range library.SubLibraries {
			if (library.SubLibraries[i] == nil) != (libraryOther.SubLibraries[i] == nil) {
				SubLibrariesDifferent = true
				break
			} else if library.SubLibraries[i] != nil && libraryOther.SubLibraries[i] != nil {
				// this is a pointer comparaison
				if library.SubLibraries[i] != libraryOther.SubLibraries[i] {
					SubLibrariesDifferent = true
					break
				}
			}
		}
	}
	if SubLibrariesDifferent {
		ops := Diff(stage, library, libraryOther, "SubLibraries", libraryOther.SubLibraries, library.SubLibraries)
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
	if note.LayoutDirection != noteOther.LayoutDirection {
		diffs = append(diffs, note.GongMarshallField(stage, "LayoutDirection"))
	}
	ProductsDifferent := false
	if len(note.Products) != len(noteOther.Products) {
		ProductsDifferent = true
	} else {
		for i := range note.Products {
			if (note.Products[i] == nil) != (noteOther.Products[i] == nil) {
				ProductsDifferent = true
				break
			} else if note.Products[i] != nil && noteOther.Products[i] != nil {
				// this is a pointer comparaison
				if note.Products[i] != noteOther.Products[i] {
					ProductsDifferent = true
					break
				}
			}
		}
	}
	if ProductsDifferent {
		ops := Diff(stage, note, noteOther, "Products", noteOther.Products, note.Products)
		diffs = append(diffs, ops)
	}
	TasksDifferent := false
	if len(note.Tasks) != len(noteOther.Tasks) {
		TasksDifferent = true
	} else {
		for i := range note.Tasks {
			if (note.Tasks[i] == nil) != (noteOther.Tasks[i] == nil) {
				TasksDifferent = true
				break
			} else if note.Tasks[i] != nil && noteOther.Tasks[i] != nil {
				// this is a pointer comparaison
				if note.Tasks[i] != noteOther.Tasks[i] {
					TasksDifferent = true
					break
				}
			}
		}
	}
	if TasksDifferent {
		ops := Diff(stage, note, noteOther, "Tasks", noteOther.Tasks, note.Tasks)
		diffs = append(diffs, ops)
	}
	ResourcesDifferent := false
	if len(note.Resources) != len(noteOther.Resources) {
		ResourcesDifferent = true
	} else {
		for i := range note.Resources {
			if (note.Resources[i] == nil) != (noteOther.Resources[i] == nil) {
				ResourcesDifferent = true
				break
			} else if note.Resources[i] != nil && noteOther.Resources[i] != nil {
				// this is a pointer comparaison
				if note.Resources[i] != noteOther.Resources[i] {
					ResourcesDifferent = true
					break
				}
			}
		}
	}
	if ResourcesDifferent {
		ops := Diff(stage, note, noteOther, "Resources", noteOther.Resources, note.Resources)
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (noteproductshape *NoteProductShape) GongDiff(stage *Stage, noteproductshapeOther *NoteProductShape) (diffs []string) {
	// insertion point for field diffs
	if noteproductshape.Name != noteproductshapeOther.Name {
		diffs = append(diffs, noteproductshape.GongMarshallField(stage, "Name"))
	}
	if (noteproductshape.Note == nil) != (noteproductshapeOther.Note == nil) {
		diffs = append(diffs, noteproductshape.GongMarshallField(stage, "Note"))
	} else if noteproductshape.Note != nil && noteproductshapeOther.Note != nil {
		if noteproductshape.Note != noteproductshapeOther.Note {
			diffs = append(diffs, noteproductshape.GongMarshallField(stage, "Note"))
		}
	}
	if (noteproductshape.Product == nil) != (noteproductshapeOther.Product == nil) {
		diffs = append(diffs, noteproductshape.GongMarshallField(stage, "Product"))
	} else if noteproductshape.Product != nil && noteproductshapeOther.Product != nil {
		if noteproductshape.Product != noteproductshapeOther.Product {
			diffs = append(diffs, noteproductshape.GongMarshallField(stage, "Product"))
		}
	}
	if noteproductshape.StartRatio != noteproductshapeOther.StartRatio {
		diffs = append(diffs, noteproductshape.GongMarshallField(stage, "StartRatio"))
	}
	if noteproductshape.EndRatio != noteproductshapeOther.EndRatio {
		diffs = append(diffs, noteproductshape.GongMarshallField(stage, "EndRatio"))
	}
	if noteproductshape.StartOrientation != noteproductshapeOther.StartOrientation {
		diffs = append(diffs, noteproductshape.GongMarshallField(stage, "StartOrientation"))
	}
	if noteproductshape.EndOrientation != noteproductshapeOther.EndOrientation {
		diffs = append(diffs, noteproductshape.GongMarshallField(stage, "EndOrientation"))
	}
	if noteproductshape.CornerOffsetRatio != noteproductshapeOther.CornerOffsetRatio {
		diffs = append(diffs, noteproductshape.GongMarshallField(stage, "CornerOffsetRatio"))
	}
	if noteproductshape.IsHidden != noteproductshapeOther.IsHidden {
		diffs = append(diffs, noteproductshape.GongMarshallField(stage, "IsHidden"))
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
	if (noteshape.Note == nil) != (noteshapeOther.Note == nil) {
		diffs = append(diffs, noteshape.GongMarshallField(stage, "Note"))
	} else if noteshape.Note != nil && noteshapeOther.Note != nil {
		if noteshape.Note != noteshapeOther.Note {
			diffs = append(diffs, noteshape.GongMarshallField(stage, "Note"))
		}
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
	if (notestakeholdershape.Note == nil) != (notestakeholdershapeOther.Note == nil) {
		diffs = append(diffs, notestakeholdershape.GongMarshallField(stage, "Note"))
	} else if notestakeholdershape.Note != nil && notestakeholdershapeOther.Note != nil {
		if notestakeholdershape.Note != notestakeholdershapeOther.Note {
			diffs = append(diffs, notestakeholdershape.GongMarshallField(stage, "Note"))
		}
	}
	if (notestakeholdershape.Stakeholder == nil) != (notestakeholdershapeOther.Stakeholder == nil) {
		diffs = append(diffs, notestakeholdershape.GongMarshallField(stage, "Stakeholder"))
	} else if notestakeholdershape.Stakeholder != nil && notestakeholdershapeOther.Stakeholder != nil {
		if notestakeholdershape.Stakeholder != notestakeholdershapeOther.Stakeholder {
			diffs = append(diffs, notestakeholdershape.GongMarshallField(stage, "Stakeholder"))
		}
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

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (notetaskshape *NoteTaskShape) GongDiff(stage *Stage, notetaskshapeOther *NoteTaskShape) (diffs []string) {
	// insertion point for field diffs
	if notetaskshape.Name != notetaskshapeOther.Name {
		diffs = append(diffs, notetaskshape.GongMarshallField(stage, "Name"))
	}
	if (notetaskshape.Note == nil) != (notetaskshapeOther.Note == nil) {
		diffs = append(diffs, notetaskshape.GongMarshallField(stage, "Note"))
	} else if notetaskshape.Note != nil && notetaskshapeOther.Note != nil {
		if notetaskshape.Note != notetaskshapeOther.Note {
			diffs = append(diffs, notetaskshape.GongMarshallField(stage, "Note"))
		}
	}
	if (notetaskshape.Task == nil) != (notetaskshapeOther.Task == nil) {
		diffs = append(diffs, notetaskshape.GongMarshallField(stage, "Task"))
	} else if notetaskshape.Task != nil && notetaskshapeOther.Task != nil {
		if notetaskshape.Task != notetaskshapeOther.Task {
			diffs = append(diffs, notetaskshape.GongMarshallField(stage, "Task"))
		}
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

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (productcompositionshape *ProductCompositionShape) GongDiff(stage *Stage, productcompositionshapeOther *ProductCompositionShape) (diffs []string) {
	// insertion point for field diffs
	if productcompositionshape.Name != productcompositionshapeOther.Name {
		diffs = append(diffs, productcompositionshape.GongMarshallField(stage, "Name"))
	}
	if (productcompositionshape.Product == nil) != (productcompositionshapeOther.Product == nil) {
		diffs = append(diffs, productcompositionshape.GongMarshallField(stage, "Product"))
	} else if productcompositionshape.Product != nil && productcompositionshapeOther.Product != nil {
		if productcompositionshape.Product != productcompositionshapeOther.Product {
			diffs = append(diffs, productcompositionshape.GongMarshallField(stage, "Product"))
		}
	}
	if productcompositionshape.StartRatio != productcompositionshapeOther.StartRatio {
		diffs = append(diffs, productcompositionshape.GongMarshallField(stage, "StartRatio"))
	}
	if productcompositionshape.EndRatio != productcompositionshapeOther.EndRatio {
		diffs = append(diffs, productcompositionshape.GongMarshallField(stage, "EndRatio"))
	}
	if productcompositionshape.StartOrientation != productcompositionshapeOther.StartOrientation {
		diffs = append(diffs, productcompositionshape.GongMarshallField(stage, "StartOrientation"))
	}
	if productcompositionshape.EndOrientation != productcompositionshapeOther.EndOrientation {
		diffs = append(diffs, productcompositionshape.GongMarshallField(stage, "EndOrientation"))
	}
	if productcompositionshape.CornerOffsetRatio != productcompositionshapeOther.CornerOffsetRatio {
		diffs = append(diffs, productcompositionshape.GongMarshallField(stage, "CornerOffsetRatio"))
	}
	if productcompositionshape.IsHidden != productcompositionshapeOther.IsHidden {
		diffs = append(diffs, productcompositionshape.GongMarshallField(stage, "IsHidden"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (productshape *ProductShape) GongDiff(stage *Stage, productshapeOther *ProductShape) (diffs []string) {
	// insertion point for field diffs
	if productshape.Name != productshapeOther.Name {
		diffs = append(diffs, productshape.GongMarshallField(stage, "Name"))
	}
	if (productshape.Product == nil) != (productshapeOther.Product == nil) {
		diffs = append(diffs, productshape.GongMarshallField(stage, "Product"))
	} else if productshape.Product != nil && productshapeOther.Product != nil {
		if productshape.Product != productshapeOther.Product {
			diffs = append(diffs, productshape.GongMarshallField(stage, "Product"))
		}
	}
	if productshape.IsExpanded != productshapeOther.IsExpanded {
		diffs = append(diffs, productshape.GongMarshallField(stage, "IsExpanded"))
	}
	if productshape.X != productshapeOther.X {
		diffs = append(diffs, productshape.GongMarshallField(stage, "X"))
	}
	if productshape.Y != productshapeOther.Y {
		diffs = append(diffs, productshape.GongMarshallField(stage, "Y"))
	}
	if productshape.Width != productshapeOther.Width {
		diffs = append(diffs, productshape.GongMarshallField(stage, "Width"))
	}
	if productshape.Height != productshapeOther.Height {
		diffs = append(diffs, productshape.GongMarshallField(stage, "Height"))
	}
	if productshape.IsHidden != productshapeOther.IsHidden {
		diffs = append(diffs, productshape.GongMarshallField(stage, "IsHidden"))
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
	if requirement.LayoutDirection != requirementOther.LayoutDirection {
		diffs = append(diffs, requirement.GongMarshallField(stage, "LayoutDirection"))
	}
	SupportLevelsDifferent := false
	if len(requirement.SupportLevels) != len(requirementOther.SupportLevels) {
		SupportLevelsDifferent = true
	} else {
		for i := range requirement.SupportLevels {
			if (requirement.SupportLevels[i] == nil) != (requirementOther.SupportLevels[i] == nil) {
				SupportLevelsDifferent = true
				break
			} else if requirement.SupportLevels[i] != nil && requirementOther.SupportLevels[i] != nil {
				// this is a pointer comparaison
				if requirement.SupportLevels[i] != requirementOther.SupportLevels[i] {
					SupportLevelsDifferent = true
					break
				}
			}
		}
	}
	if SupportLevelsDifferent {
		ops := Diff(stage, requirement, requirementOther, "SupportLevels", requirementOther.SupportLevels, requirement.SupportLevels)
		diffs = append(diffs, ops)
	}
	ConceptsDifferent := false
	if len(requirement.Concepts) != len(requirementOther.Concepts) {
		ConceptsDifferent = true
	} else {
		for i := range requirement.Concepts {
			if (requirement.Concepts[i] == nil) != (requirementOther.Concepts[i] == nil) {
				ConceptsDifferent = true
				break
			} else if requirement.Concepts[i] != nil && requirementOther.Concepts[i] != nil {
				// this is a pointer comparaison
				if requirement.Concepts[i] != requirementOther.Concepts[i] {
					ConceptsDifferent = true
					break
				}
			}
		}
	}
	if ConceptsDifferent {
		ops := Diff(stage, requirement, requirementOther, "Concepts", requirementOther.Concepts, requirement.Concepts)
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
	if (requirementshape.Requirement == nil) != (requirementshapeOther.Requirement == nil) {
		diffs = append(diffs, requirementshape.GongMarshallField(stage, "Requirement"))
	} else if requirementshape.Requirement != nil && requirementshapeOther.Requirement != nil {
		if requirementshape.Requirement != requirementshapeOther.Requirement {
			diffs = append(diffs, requirementshape.GongMarshallField(stage, "Requirement"))
		}
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
	if stakeholder.LayoutDirection != stakeholderOther.LayoutDirection {
		diffs = append(diffs, stakeholder.GongMarshallField(stage, "LayoutDirection"))
	}
	if stakeholder.Description != stakeholderOther.Description {
		diffs = append(diffs, stakeholder.GongMarshallField(stage, "Description"))
	}
	ConcernsDifferent := false
	if len(stakeholder.Concerns) != len(stakeholderOther.Concerns) {
		ConcernsDifferent = true
	} else {
		for i := range stakeholder.Concerns {
			if (stakeholder.Concerns[i] == nil) != (stakeholderOther.Concerns[i] == nil) {
				ConcernsDifferent = true
				break
			} else if stakeholder.Concerns[i] != nil && stakeholderOther.Concerns[i] != nil {
				// this is a pointer comparaison
				if stakeholder.Concerns[i] != stakeholderOther.Concerns[i] {
					ConcernsDifferent = true
					break
				}
			}
		}
	}
	if ConcernsDifferent {
		ops := Diff(stage, stakeholder, stakeholderOther, "Concerns", stakeholderOther.Concerns, stakeholder.Concerns)
		diffs = append(diffs, ops)
	}
	SubStakeholdersDifferent := false
	if len(stakeholder.SubStakeholders) != len(stakeholderOther.SubStakeholders) {
		SubStakeholdersDifferent = true
	} else {
		for i := range stakeholder.SubStakeholders {
			if (stakeholder.SubStakeholders[i] == nil) != (stakeholderOther.SubStakeholders[i] == nil) {
				SubStakeholdersDifferent = true
				break
			} else if stakeholder.SubStakeholders[i] != nil && stakeholderOther.SubStakeholders[i] != nil {
				// this is a pointer comparaison
				if stakeholder.SubStakeholders[i] != stakeholderOther.SubStakeholders[i] {
					SubStakeholdersDifferent = true
					break
				}
			}
		}
	}
	if SubStakeholdersDifferent {
		ops := Diff(stage, stakeholder, stakeholderOther, "SubStakeholders", stakeholderOther.SubStakeholders, stakeholder.SubStakeholders)
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
	if (stakeholdercompositionshape.Stakeholder == nil) != (stakeholdercompositionshapeOther.Stakeholder == nil) {
		diffs = append(diffs, stakeholdercompositionshape.GongMarshallField(stage, "Stakeholder"))
	} else if stakeholdercompositionshape.Stakeholder != nil && stakeholdercompositionshapeOther.Stakeholder != nil {
		if stakeholdercompositionshape.Stakeholder != stakeholdercompositionshapeOther.Stakeholder {
			diffs = append(diffs, stakeholdercompositionshape.GongMarshallField(stage, "Stakeholder"))
		}
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

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (stakeholderconcernshape *StakeholderConcernShape) GongDiff(stage *Stage, stakeholderconcernshapeOther *StakeholderConcernShape) (diffs []string) {
	// insertion point for field diffs
	if stakeholderconcernshape.Name != stakeholderconcernshapeOther.Name {
		diffs = append(diffs, stakeholderconcernshape.GongMarshallField(stage, "Name"))
	}
	if (stakeholderconcernshape.Stakeholder == nil) != (stakeholderconcernshapeOther.Stakeholder == nil) {
		diffs = append(diffs, stakeholderconcernshape.GongMarshallField(stage, "Stakeholder"))
	} else if stakeholderconcernshape.Stakeholder != nil && stakeholderconcernshapeOther.Stakeholder != nil {
		if stakeholderconcernshape.Stakeholder != stakeholderconcernshapeOther.Stakeholder {
			diffs = append(diffs, stakeholderconcernshape.GongMarshallField(stage, "Stakeholder"))
		}
	}
	if (stakeholderconcernshape.Concern == nil) != (stakeholderconcernshapeOther.Concern == nil) {
		diffs = append(diffs, stakeholderconcernshape.GongMarshallField(stage, "Concern"))
	} else if stakeholderconcernshape.Concern != nil && stakeholderconcernshapeOther.Concern != nil {
		if stakeholderconcernshape.Concern != stakeholderconcernshapeOther.Concern {
			diffs = append(diffs, stakeholderconcernshape.GongMarshallField(stage, "Concern"))
		}
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

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (stakeholdershape *StakeholderShape) GongDiff(stage *Stage, stakeholdershapeOther *StakeholderShape) (diffs []string) {
	// insertion point for field diffs
	if stakeholdershape.Name != stakeholdershapeOther.Name {
		diffs = append(diffs, stakeholdershape.GongMarshallField(stage, "Name"))
	}
	if (stakeholdershape.Stakeholder == nil) != (stakeholdershapeOther.Stakeholder == nil) {
		diffs = append(diffs, stakeholdershape.GongMarshallField(stage, "Stakeholder"))
	} else if stakeholdershape.Stakeholder != nil && stakeholdershapeOther.Stakeholder != nil {
		if stakeholdershape.Stakeholder != stakeholdershapeOther.Stakeholder {
			diffs = append(diffs, stakeholdershape.GongMarshallField(stage, "Stakeholder"))
		}
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
	if (supportlevel.Tool == nil) != (supportlevelOther.Tool == nil) {
		diffs = append(diffs, supportlevel.GongMarshallField(stage, "Tool"))
	} else if supportlevel.Tool != nil && supportlevelOther.Tool != nil {
		if supportlevel.Tool != supportlevelOther.Tool {
			diffs = append(diffs, supportlevel.GongMarshallField(stage, "Tool"))
		}
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

// Diff returns the sequence of operations to transform oldSlice into newSlice.
// It requires type T to be comparable (e.g., pointers, ints, strings).
func Diff[T1, T2 PointerToGongstruct](stage *Stage, a, b T1, fieldName string, oldSlice, newSlice []T2) (ops string) {
	m, n := len(oldSlice), len(newSlice)

	// 1. Build the LCS (Longest Common Subsequence) Matrix
	// This helps us find the "anchor" elements that shouldn't move.
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if oldSlice[i] == newSlice[j] {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				// Take the maximum of previous options
				if dp[i][j+1] > dp[i+1][j] {
					dp[i+1][j+1] = dp[i][j+1]
				} else {
					dp[i+1][j+1] = dp[i+1][j]
				}
			}
		}
	}

	// 2. Backtrack to find which indices in oldSlice are part of the LCS
	// We use a map for O(1) lookups.
	keptIndices := make(map[int]bool)
	i, j := m, n
	for i > 0 && j > 0 {
		if oldSlice[i-1] == newSlice[j-1] {
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

	// Create a temporary view of what's left after deletions for tracking matches
	var currentLCS []T2
	for k := 0; k < m; k++ {
		if keptIndices[k] {
			currentLCS = append(currentLCS, oldSlice[k])
		}
	}

	lcsIdx := 0
	// Iterate through the NEW slice. If it matches the current LCS head, we keep it.
	// If it doesn't match, it must be inserted here.
	for k, targetVal := range newSlice {
		if lcsIdx < len(currentLCS) && currentLCS[lcsIdx] == targetVal {
			lcsIdx++
		} else {
			ops += fmt.Sprintf("\n\t%s.%s = slices.Insert( %s.%s, %d, %s)", a.GongGetIdentifier(stage), fieldName, a.GongGetIdentifier(stage), fieldName, k, targetVal.GongGetIdentifier(stage))
		}
	}

	return ops
}
