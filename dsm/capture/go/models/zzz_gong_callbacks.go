// generated code - do not edit
package models

// AfterCreateFromFront is called after a create from front
func AfterCreateFromFront[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *AnalysisNeed:
		if stage.OnAfterAnalysisNeedCreateCallback != nil {
			stage.OnAfterAnalysisNeedCreateCallback.OnAfterCreate(stage, target)
		}
	case *Concept:
		if stage.OnAfterConceptCreateCallback != nil {
			stage.OnAfterConceptCreateCallback.OnAfterCreate(stage, target)
		}
	case *ConceptShape:
		if stage.OnAfterConceptShapeCreateCallback != nil {
			stage.OnAfterConceptShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Concern:
		if stage.OnAfterConcernCreateCallback != nil {
			stage.OnAfterConcernCreateCallback.OnAfterCreate(stage, target)
		}
	case *ConcernCompositionShape:
		if stage.OnAfterConcernCompositionShapeCreateCallback != nil {
			stage.OnAfterConcernCompositionShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *ConcernInputShape:
		if stage.OnAfterConcernInputShapeCreateCallback != nil {
			stage.OnAfterConcernInputShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *ConcernOutputShape:
		if stage.OnAfterConcernOutputShapeCreateCallback != nil {
			stage.OnAfterConcernOutputShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *ConcernShape:
		if stage.OnAfterConcernShapeCreateCallback != nil {
			stage.OnAfterConcernShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Deliverable:
		if stage.OnAfterDeliverableCreateCallback != nil {
			stage.OnAfterDeliverableCreateCallback.OnAfterCreate(stage, target)
		}
	case *Diagram:
		if stage.OnAfterDiagramCreateCallback != nil {
			stage.OnAfterDiagramCreateCallback.OnAfterCreate(stage, target)
		}
	case *Library:
		if stage.OnAfterLibraryCreateCallback != nil {
			stage.OnAfterLibraryCreateCallback.OnAfterCreate(stage, target)
		}
	case *Note:
		if stage.OnAfterNoteCreateCallback != nil {
			stage.OnAfterNoteCreateCallback.OnAfterCreate(stage, target)
		}
	case *NoteProductShape:
		if stage.OnAfterNoteProductShapeCreateCallback != nil {
			stage.OnAfterNoteProductShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *NoteShape:
		if stage.OnAfterNoteShapeCreateCallback != nil {
			stage.OnAfterNoteShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *NoteStakeholderShape:
		if stage.OnAfterNoteStakeholderShapeCreateCallback != nil {
			stage.OnAfterNoteStakeholderShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *NoteTaskShape:
		if stage.OnAfterNoteTaskShapeCreateCallback != nil {
			stage.OnAfterNoteTaskShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *ProductCompositionShape:
		if stage.OnAfterProductCompositionShapeCreateCallback != nil {
			stage.OnAfterProductCompositionShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *ProductShape:
		if stage.OnAfterProductShapeCreateCallback != nil {
			stage.OnAfterProductShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Requirement:
		if stage.OnAfterRequirementCreateCallback != nil {
			stage.OnAfterRequirementCreateCallback.OnAfterCreate(stage, target)
		}
	case *RequirementShape:
		if stage.OnAfterRequirementShapeCreateCallback != nil {
			stage.OnAfterRequirementShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Stakeholder:
		if stage.OnAfterStakeholderCreateCallback != nil {
			stage.OnAfterStakeholderCreateCallback.OnAfterCreate(stage, target)
		}
	case *StakeholderCompositionShape:
		if stage.OnAfterStakeholderCompositionShapeCreateCallback != nil {
			stage.OnAfterStakeholderCompositionShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *StakeholderConcernShape:
		if stage.OnAfterStakeholderConcernShapeCreateCallback != nil {
			stage.OnAfterStakeholderConcernShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *StakeholderShape:
		if stage.OnAfterStakeholderShapeCreateCallback != nil {
			stage.OnAfterStakeholderShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *SupportLevel:
		if stage.OnAfterSupportLevelCreateCallback != nil {
			stage.OnAfterSupportLevelCreateCallback.OnAfterCreate(stage, target)
		}
	case *Tool:
		if stage.OnAfterToolCreateCallback != nil {
			stage.OnAfterToolCreateCallback.OnAfterCreate(stage, target)
		}
	default:
		_ = target
	}
}

type Gong__MouseEvent struct {
	ShiftKey bool
}

// OnAfterUpdateFromFront is called after a update from front
func OnAfterUpdateFromFront[Type Gongstruct](stage *Stage, old, new *Type) {

	switch oldTarget := any(old).(type) {
	// insertion point
	case *AnalysisNeed:
		newTarget := any(new).(*AnalysisNeed)
		if stage.OnAfterAnalysisNeedUpdateCallback != nil {
			stage.OnAfterAnalysisNeedUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Concept:
		newTarget := any(new).(*Concept)
		if stage.OnAfterConceptUpdateCallback != nil {
			stage.OnAfterConceptUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ConceptShape:
		newTarget := any(new).(*ConceptShape)
		if stage.OnAfterConceptShapeUpdateCallback != nil {
			stage.OnAfterConceptShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Concern:
		newTarget := any(new).(*Concern)
		if stage.OnAfterConcernUpdateCallback != nil {
			stage.OnAfterConcernUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ConcernCompositionShape:
		newTarget := any(new).(*ConcernCompositionShape)
		if stage.OnAfterConcernCompositionShapeUpdateCallback != nil {
			stage.OnAfterConcernCompositionShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ConcernInputShape:
		newTarget := any(new).(*ConcernInputShape)
		if stage.OnAfterConcernInputShapeUpdateCallback != nil {
			stage.OnAfterConcernInputShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ConcernOutputShape:
		newTarget := any(new).(*ConcernOutputShape)
		if stage.OnAfterConcernOutputShapeUpdateCallback != nil {
			stage.OnAfterConcernOutputShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ConcernShape:
		newTarget := any(new).(*ConcernShape)
		if stage.OnAfterConcernShapeUpdateCallback != nil {
			stage.OnAfterConcernShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Deliverable:
		newTarget := any(new).(*Deliverable)
		if stage.OnAfterDeliverableUpdateCallback != nil {
			stage.OnAfterDeliverableUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Diagram:
		newTarget := any(new).(*Diagram)
		if stage.OnAfterDiagramUpdateCallback != nil {
			stage.OnAfterDiagramUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Library:
		newTarget := any(new).(*Library)
		if stage.OnAfterLibraryUpdateCallback != nil {
			stage.OnAfterLibraryUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Note:
		newTarget := any(new).(*Note)
		if stage.OnAfterNoteUpdateCallback != nil {
			stage.OnAfterNoteUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *NoteProductShape:
		newTarget := any(new).(*NoteProductShape)
		if stage.OnAfterNoteProductShapeUpdateCallback != nil {
			stage.OnAfterNoteProductShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *NoteShape:
		newTarget := any(new).(*NoteShape)
		if stage.OnAfterNoteShapeUpdateCallback != nil {
			stage.OnAfterNoteShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *NoteStakeholderShape:
		newTarget := any(new).(*NoteStakeholderShape)
		if stage.OnAfterNoteStakeholderShapeUpdateCallback != nil {
			stage.OnAfterNoteStakeholderShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *NoteTaskShape:
		newTarget := any(new).(*NoteTaskShape)
		if stage.OnAfterNoteTaskShapeUpdateCallback != nil {
			stage.OnAfterNoteTaskShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ProductCompositionShape:
		newTarget := any(new).(*ProductCompositionShape)
		if stage.OnAfterProductCompositionShapeUpdateCallback != nil {
			stage.OnAfterProductCompositionShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ProductShape:
		newTarget := any(new).(*ProductShape)
		if stage.OnAfterProductShapeUpdateCallback != nil {
			stage.OnAfterProductShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Requirement:
		newTarget := any(new).(*Requirement)
		if stage.OnAfterRequirementUpdateCallback != nil {
			stage.OnAfterRequirementUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *RequirementShape:
		newTarget := any(new).(*RequirementShape)
		if stage.OnAfterRequirementShapeUpdateCallback != nil {
			stage.OnAfterRequirementShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Stakeholder:
		newTarget := any(new).(*Stakeholder)
		if stage.OnAfterStakeholderUpdateCallback != nil {
			stage.OnAfterStakeholderUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *StakeholderCompositionShape:
		newTarget := any(new).(*StakeholderCompositionShape)
		if stage.OnAfterStakeholderCompositionShapeUpdateCallback != nil {
			stage.OnAfterStakeholderCompositionShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *StakeholderConcernShape:
		newTarget := any(new).(*StakeholderConcernShape)
		if stage.OnAfterStakeholderConcernShapeUpdateCallback != nil {
			stage.OnAfterStakeholderConcernShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *StakeholderShape:
		newTarget := any(new).(*StakeholderShape)
		if stage.OnAfterStakeholderShapeUpdateCallback != nil {
			stage.OnAfterStakeholderShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *SupportLevel:
		newTarget := any(new).(*SupportLevel)
		if stage.OnAfterSupportLevelUpdateCallback != nil {
			stage.OnAfterSupportLevelUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Tool:
		newTarget := any(new).(*Tool)
		if stage.OnAfterToolUpdateCallback != nil {
			stage.OnAfterToolUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	default:
		_ = oldTarget
	}
}

// AfterDeleteFromFront is called after a delete from front
func AfterDeleteFromFront[Type Gongstruct](stage *Stage, staged, front *Type) {

	switch front := any(front).(type) {
	// insertion point
	case *AnalysisNeed:
		if stage.OnAfterAnalysisNeedDeleteCallback != nil {
			staged := any(staged).(*AnalysisNeed)
			stage.OnAfterAnalysisNeedDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Concept:
		if stage.OnAfterConceptDeleteCallback != nil {
			staged := any(staged).(*Concept)
			stage.OnAfterConceptDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ConceptShape:
		if stage.OnAfterConceptShapeDeleteCallback != nil {
			staged := any(staged).(*ConceptShape)
			stage.OnAfterConceptShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Concern:
		if stage.OnAfterConcernDeleteCallback != nil {
			staged := any(staged).(*Concern)
			stage.OnAfterConcernDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ConcernCompositionShape:
		if stage.OnAfterConcernCompositionShapeDeleteCallback != nil {
			staged := any(staged).(*ConcernCompositionShape)
			stage.OnAfterConcernCompositionShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ConcernInputShape:
		if stage.OnAfterConcernInputShapeDeleteCallback != nil {
			staged := any(staged).(*ConcernInputShape)
			stage.OnAfterConcernInputShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ConcernOutputShape:
		if stage.OnAfterConcernOutputShapeDeleteCallback != nil {
			staged := any(staged).(*ConcernOutputShape)
			stage.OnAfterConcernOutputShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ConcernShape:
		if stage.OnAfterConcernShapeDeleteCallback != nil {
			staged := any(staged).(*ConcernShape)
			stage.OnAfterConcernShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Deliverable:
		if stage.OnAfterDeliverableDeleteCallback != nil {
			staged := any(staged).(*Deliverable)
			stage.OnAfterDeliverableDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Diagram:
		if stage.OnAfterDiagramDeleteCallback != nil {
			staged := any(staged).(*Diagram)
			stage.OnAfterDiagramDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Library:
		if stage.OnAfterLibraryDeleteCallback != nil {
			staged := any(staged).(*Library)
			stage.OnAfterLibraryDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Note:
		if stage.OnAfterNoteDeleteCallback != nil {
			staged := any(staged).(*Note)
			stage.OnAfterNoteDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *NoteProductShape:
		if stage.OnAfterNoteProductShapeDeleteCallback != nil {
			staged := any(staged).(*NoteProductShape)
			stage.OnAfterNoteProductShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *NoteShape:
		if stage.OnAfterNoteShapeDeleteCallback != nil {
			staged := any(staged).(*NoteShape)
			stage.OnAfterNoteShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *NoteStakeholderShape:
		if stage.OnAfterNoteStakeholderShapeDeleteCallback != nil {
			staged := any(staged).(*NoteStakeholderShape)
			stage.OnAfterNoteStakeholderShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *NoteTaskShape:
		if stage.OnAfterNoteTaskShapeDeleteCallback != nil {
			staged := any(staged).(*NoteTaskShape)
			stage.OnAfterNoteTaskShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ProductCompositionShape:
		if stage.OnAfterProductCompositionShapeDeleteCallback != nil {
			staged := any(staged).(*ProductCompositionShape)
			stage.OnAfterProductCompositionShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ProductShape:
		if stage.OnAfterProductShapeDeleteCallback != nil {
			staged := any(staged).(*ProductShape)
			stage.OnAfterProductShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Requirement:
		if stage.OnAfterRequirementDeleteCallback != nil {
			staged := any(staged).(*Requirement)
			stage.OnAfterRequirementDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *RequirementShape:
		if stage.OnAfterRequirementShapeDeleteCallback != nil {
			staged := any(staged).(*RequirementShape)
			stage.OnAfterRequirementShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Stakeholder:
		if stage.OnAfterStakeholderDeleteCallback != nil {
			staged := any(staged).(*Stakeholder)
			stage.OnAfterStakeholderDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *StakeholderCompositionShape:
		if stage.OnAfterStakeholderCompositionShapeDeleteCallback != nil {
			staged := any(staged).(*StakeholderCompositionShape)
			stage.OnAfterStakeholderCompositionShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *StakeholderConcernShape:
		if stage.OnAfterStakeholderConcernShapeDeleteCallback != nil {
			staged := any(staged).(*StakeholderConcernShape)
			stage.OnAfterStakeholderConcernShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *StakeholderShape:
		if stage.OnAfterStakeholderShapeDeleteCallback != nil {
			staged := any(staged).(*StakeholderShape)
			stage.OnAfterStakeholderShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *SupportLevel:
		if stage.OnAfterSupportLevelDeleteCallback != nil {
			staged := any(staged).(*SupportLevel)
			stage.OnAfterSupportLevelDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Tool:
		if stage.OnAfterToolDeleteCallback != nil {
			staged := any(staged).(*Tool)
			stage.OnAfterToolDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	default:
		_ = front
	}
}

// AfterReadFromFront is called after a Read from front
func AfterReadFromFront[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *AnalysisNeed:
		if stage.OnAfterAnalysisNeedReadCallback != nil {
			stage.OnAfterAnalysisNeedReadCallback.OnAfterRead(stage, target)
		}
	case *Concept:
		if stage.OnAfterConceptReadCallback != nil {
			stage.OnAfterConceptReadCallback.OnAfterRead(stage, target)
		}
	case *ConceptShape:
		if stage.OnAfterConceptShapeReadCallback != nil {
			stage.OnAfterConceptShapeReadCallback.OnAfterRead(stage, target)
		}
	case *Concern:
		if stage.OnAfterConcernReadCallback != nil {
			stage.OnAfterConcernReadCallback.OnAfterRead(stage, target)
		}
	case *ConcernCompositionShape:
		if stage.OnAfterConcernCompositionShapeReadCallback != nil {
			stage.OnAfterConcernCompositionShapeReadCallback.OnAfterRead(stage, target)
		}
	case *ConcernInputShape:
		if stage.OnAfterConcernInputShapeReadCallback != nil {
			stage.OnAfterConcernInputShapeReadCallback.OnAfterRead(stage, target)
		}
	case *ConcernOutputShape:
		if stage.OnAfterConcernOutputShapeReadCallback != nil {
			stage.OnAfterConcernOutputShapeReadCallback.OnAfterRead(stage, target)
		}
	case *ConcernShape:
		if stage.OnAfterConcernShapeReadCallback != nil {
			stage.OnAfterConcernShapeReadCallback.OnAfterRead(stage, target)
		}
	case *Deliverable:
		if stage.OnAfterDeliverableReadCallback != nil {
			stage.OnAfterDeliverableReadCallback.OnAfterRead(stage, target)
		}
	case *Diagram:
		if stage.OnAfterDiagramReadCallback != nil {
			stage.OnAfterDiagramReadCallback.OnAfterRead(stage, target)
		}
	case *Library:
		if stage.OnAfterLibraryReadCallback != nil {
			stage.OnAfterLibraryReadCallback.OnAfterRead(stage, target)
		}
	case *Note:
		if stage.OnAfterNoteReadCallback != nil {
			stage.OnAfterNoteReadCallback.OnAfterRead(stage, target)
		}
	case *NoteProductShape:
		if stage.OnAfterNoteProductShapeReadCallback != nil {
			stage.OnAfterNoteProductShapeReadCallback.OnAfterRead(stage, target)
		}
	case *NoteShape:
		if stage.OnAfterNoteShapeReadCallback != nil {
			stage.OnAfterNoteShapeReadCallback.OnAfterRead(stage, target)
		}
	case *NoteStakeholderShape:
		if stage.OnAfterNoteStakeholderShapeReadCallback != nil {
			stage.OnAfterNoteStakeholderShapeReadCallback.OnAfterRead(stage, target)
		}
	case *NoteTaskShape:
		if stage.OnAfterNoteTaskShapeReadCallback != nil {
			stage.OnAfterNoteTaskShapeReadCallback.OnAfterRead(stage, target)
		}
	case *ProductCompositionShape:
		if stage.OnAfterProductCompositionShapeReadCallback != nil {
			stage.OnAfterProductCompositionShapeReadCallback.OnAfterRead(stage, target)
		}
	case *ProductShape:
		if stage.OnAfterProductShapeReadCallback != nil {
			stage.OnAfterProductShapeReadCallback.OnAfterRead(stage, target)
		}
	case *Requirement:
		if stage.OnAfterRequirementReadCallback != nil {
			stage.OnAfterRequirementReadCallback.OnAfterRead(stage, target)
		}
	case *RequirementShape:
		if stage.OnAfterRequirementShapeReadCallback != nil {
			stage.OnAfterRequirementShapeReadCallback.OnAfterRead(stage, target)
		}
	case *Stakeholder:
		if stage.OnAfterStakeholderReadCallback != nil {
			stage.OnAfterStakeholderReadCallback.OnAfterRead(stage, target)
		}
	case *StakeholderCompositionShape:
		if stage.OnAfterStakeholderCompositionShapeReadCallback != nil {
			stage.OnAfterStakeholderCompositionShapeReadCallback.OnAfterRead(stage, target)
		}
	case *StakeholderConcernShape:
		if stage.OnAfterStakeholderConcernShapeReadCallback != nil {
			stage.OnAfterStakeholderConcernShapeReadCallback.OnAfterRead(stage, target)
		}
	case *StakeholderShape:
		if stage.OnAfterStakeholderShapeReadCallback != nil {
			stage.OnAfterStakeholderShapeReadCallback.OnAfterRead(stage, target)
		}
	case *SupportLevel:
		if stage.OnAfterSupportLevelReadCallback != nil {
			stage.OnAfterSupportLevelReadCallback.OnAfterRead(stage, target)
		}
	case *Tool:
		if stage.OnAfterToolReadCallback != nil {
			stage.OnAfterToolReadCallback.OnAfterRead(stage, target)
		}
	default:
		_ = target
	}
}

// SetCallbackAfterUpdateFromFront is a function to set up callback that is robust to refactoring
func SetCallbackAfterUpdateFromFront[Type Gongstruct](stage *Stage, callback OnAfterUpdateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
	// insertion point
	case *AnalysisNeed:
		stage.OnAfterAnalysisNeedUpdateCallback = any(callback).(OnAfterUpdateInterface[AnalysisNeed])
	case *Concept:
		stage.OnAfterConceptUpdateCallback = any(callback).(OnAfterUpdateInterface[Concept])
	case *ConceptShape:
		stage.OnAfterConceptShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[ConceptShape])
	case *Concern:
		stage.OnAfterConcernUpdateCallback = any(callback).(OnAfterUpdateInterface[Concern])
	case *ConcernCompositionShape:
		stage.OnAfterConcernCompositionShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[ConcernCompositionShape])
	case *ConcernInputShape:
		stage.OnAfterConcernInputShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[ConcernInputShape])
	case *ConcernOutputShape:
		stage.OnAfterConcernOutputShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[ConcernOutputShape])
	case *ConcernShape:
		stage.OnAfterConcernShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[ConcernShape])
	case *Deliverable:
		stage.OnAfterDeliverableUpdateCallback = any(callback).(OnAfterUpdateInterface[Deliverable])
	case *Diagram:
		stage.OnAfterDiagramUpdateCallback = any(callback).(OnAfterUpdateInterface[Diagram])
	case *Library:
		stage.OnAfterLibraryUpdateCallback = any(callback).(OnAfterUpdateInterface[Library])
	case *Note:
		stage.OnAfterNoteUpdateCallback = any(callback).(OnAfterUpdateInterface[Note])
	case *NoteProductShape:
		stage.OnAfterNoteProductShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[NoteProductShape])
	case *NoteShape:
		stage.OnAfterNoteShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[NoteShape])
	case *NoteStakeholderShape:
		stage.OnAfterNoteStakeholderShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[NoteStakeholderShape])
	case *NoteTaskShape:
		stage.OnAfterNoteTaskShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[NoteTaskShape])
	case *ProductCompositionShape:
		stage.OnAfterProductCompositionShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[ProductCompositionShape])
	case *ProductShape:
		stage.OnAfterProductShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[ProductShape])
	case *Requirement:
		stage.OnAfterRequirementUpdateCallback = any(callback).(OnAfterUpdateInterface[Requirement])
	case *RequirementShape:
		stage.OnAfterRequirementShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[RequirementShape])
	case *Stakeholder:
		stage.OnAfterStakeholderUpdateCallback = any(callback).(OnAfterUpdateInterface[Stakeholder])
	case *StakeholderCompositionShape:
		stage.OnAfterStakeholderCompositionShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[StakeholderCompositionShape])
	case *StakeholderConcernShape:
		stage.OnAfterStakeholderConcernShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[StakeholderConcernShape])
	case *StakeholderShape:
		stage.OnAfterStakeholderShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[StakeholderShape])
	case *SupportLevel:
		stage.OnAfterSupportLevelUpdateCallback = any(callback).(OnAfterUpdateInterface[SupportLevel])
	case *Tool:
		stage.OnAfterToolUpdateCallback = any(callback).(OnAfterUpdateInterface[Tool])
	}
}
func SetCallbackAfterCreateFromFront[Type Gongstruct](stage *Stage, callback OnAfterCreateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
	// insertion point
	case *AnalysisNeed:
		stage.OnAfterAnalysisNeedCreateCallback = any(callback).(OnAfterCreateInterface[AnalysisNeed])
	case *Concept:
		stage.OnAfterConceptCreateCallback = any(callback).(OnAfterCreateInterface[Concept])
	case *ConceptShape:
		stage.OnAfterConceptShapeCreateCallback = any(callback).(OnAfterCreateInterface[ConceptShape])
	case *Concern:
		stage.OnAfterConcernCreateCallback = any(callback).(OnAfterCreateInterface[Concern])
	case *ConcernCompositionShape:
		stage.OnAfterConcernCompositionShapeCreateCallback = any(callback).(OnAfterCreateInterface[ConcernCompositionShape])
	case *ConcernInputShape:
		stage.OnAfterConcernInputShapeCreateCallback = any(callback).(OnAfterCreateInterface[ConcernInputShape])
	case *ConcernOutputShape:
		stage.OnAfterConcernOutputShapeCreateCallback = any(callback).(OnAfterCreateInterface[ConcernOutputShape])
	case *ConcernShape:
		stage.OnAfterConcernShapeCreateCallback = any(callback).(OnAfterCreateInterface[ConcernShape])
	case *Deliverable:
		stage.OnAfterDeliverableCreateCallback = any(callback).(OnAfterCreateInterface[Deliverable])
	case *Diagram:
		stage.OnAfterDiagramCreateCallback = any(callback).(OnAfterCreateInterface[Diagram])
	case *Library:
		stage.OnAfterLibraryCreateCallback = any(callback).(OnAfterCreateInterface[Library])
	case *Note:
		stage.OnAfterNoteCreateCallback = any(callback).(OnAfterCreateInterface[Note])
	case *NoteProductShape:
		stage.OnAfterNoteProductShapeCreateCallback = any(callback).(OnAfterCreateInterface[NoteProductShape])
	case *NoteShape:
		stage.OnAfterNoteShapeCreateCallback = any(callback).(OnAfterCreateInterface[NoteShape])
	case *NoteStakeholderShape:
		stage.OnAfterNoteStakeholderShapeCreateCallback = any(callback).(OnAfterCreateInterface[NoteStakeholderShape])
	case *NoteTaskShape:
		stage.OnAfterNoteTaskShapeCreateCallback = any(callback).(OnAfterCreateInterface[NoteTaskShape])
	case *ProductCompositionShape:
		stage.OnAfterProductCompositionShapeCreateCallback = any(callback).(OnAfterCreateInterface[ProductCompositionShape])
	case *ProductShape:
		stage.OnAfterProductShapeCreateCallback = any(callback).(OnAfterCreateInterface[ProductShape])
	case *Requirement:
		stage.OnAfterRequirementCreateCallback = any(callback).(OnAfterCreateInterface[Requirement])
	case *RequirementShape:
		stage.OnAfterRequirementShapeCreateCallback = any(callback).(OnAfterCreateInterface[RequirementShape])
	case *Stakeholder:
		stage.OnAfterStakeholderCreateCallback = any(callback).(OnAfterCreateInterface[Stakeholder])
	case *StakeholderCompositionShape:
		stage.OnAfterStakeholderCompositionShapeCreateCallback = any(callback).(OnAfterCreateInterface[StakeholderCompositionShape])
	case *StakeholderConcernShape:
		stage.OnAfterStakeholderConcernShapeCreateCallback = any(callback).(OnAfterCreateInterface[StakeholderConcernShape])
	case *StakeholderShape:
		stage.OnAfterStakeholderShapeCreateCallback = any(callback).(OnAfterCreateInterface[StakeholderShape])
	case *SupportLevel:
		stage.OnAfterSupportLevelCreateCallback = any(callback).(OnAfterCreateInterface[SupportLevel])
	case *Tool:
		stage.OnAfterToolCreateCallback = any(callback).(OnAfterCreateInterface[Tool])
	}
}
func SetCallbackAfterDeleteFromFront[Type Gongstruct](stage *Stage, callback OnAfterDeleteInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
	// insertion point
	case *AnalysisNeed:
		stage.OnAfterAnalysisNeedDeleteCallback = any(callback).(OnAfterDeleteInterface[AnalysisNeed])
	case *Concept:
		stage.OnAfterConceptDeleteCallback = any(callback).(OnAfterDeleteInterface[Concept])
	case *ConceptShape:
		stage.OnAfterConceptShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[ConceptShape])
	case *Concern:
		stage.OnAfterConcernDeleteCallback = any(callback).(OnAfterDeleteInterface[Concern])
	case *ConcernCompositionShape:
		stage.OnAfterConcernCompositionShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[ConcernCompositionShape])
	case *ConcernInputShape:
		stage.OnAfterConcernInputShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[ConcernInputShape])
	case *ConcernOutputShape:
		stage.OnAfterConcernOutputShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[ConcernOutputShape])
	case *ConcernShape:
		stage.OnAfterConcernShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[ConcernShape])
	case *Deliverable:
		stage.OnAfterDeliverableDeleteCallback = any(callback).(OnAfterDeleteInterface[Deliverable])
	case *Diagram:
		stage.OnAfterDiagramDeleteCallback = any(callback).(OnAfterDeleteInterface[Diagram])
	case *Library:
		stage.OnAfterLibraryDeleteCallback = any(callback).(OnAfterDeleteInterface[Library])
	case *Note:
		stage.OnAfterNoteDeleteCallback = any(callback).(OnAfterDeleteInterface[Note])
	case *NoteProductShape:
		stage.OnAfterNoteProductShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[NoteProductShape])
	case *NoteShape:
		stage.OnAfterNoteShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[NoteShape])
	case *NoteStakeholderShape:
		stage.OnAfterNoteStakeholderShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[NoteStakeholderShape])
	case *NoteTaskShape:
		stage.OnAfterNoteTaskShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[NoteTaskShape])
	case *ProductCompositionShape:
		stage.OnAfterProductCompositionShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[ProductCompositionShape])
	case *ProductShape:
		stage.OnAfterProductShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[ProductShape])
	case *Requirement:
		stage.OnAfterRequirementDeleteCallback = any(callback).(OnAfterDeleteInterface[Requirement])
	case *RequirementShape:
		stage.OnAfterRequirementShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[RequirementShape])
	case *Stakeholder:
		stage.OnAfterStakeholderDeleteCallback = any(callback).(OnAfterDeleteInterface[Stakeholder])
	case *StakeholderCompositionShape:
		stage.OnAfterStakeholderCompositionShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[StakeholderCompositionShape])
	case *StakeholderConcernShape:
		stage.OnAfterStakeholderConcernShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[StakeholderConcernShape])
	case *StakeholderShape:
		stage.OnAfterStakeholderShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[StakeholderShape])
	case *SupportLevel:
		stage.OnAfterSupportLevelDeleteCallback = any(callback).(OnAfterDeleteInterface[SupportLevel])
	case *Tool:
		stage.OnAfterToolDeleteCallback = any(callback).(OnAfterDeleteInterface[Tool])
	}
}
func SetCallbackAfterReadFromFront[Type Gongstruct](stage *Stage, callback OnAfterReadInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
	// insertion point
	case *AnalysisNeed:
		stage.OnAfterAnalysisNeedReadCallback = any(callback).(OnAfterReadInterface[AnalysisNeed])
	case *Concept:
		stage.OnAfterConceptReadCallback = any(callback).(OnAfterReadInterface[Concept])
	case *ConceptShape:
		stage.OnAfterConceptShapeReadCallback = any(callback).(OnAfterReadInterface[ConceptShape])
	case *Concern:
		stage.OnAfterConcernReadCallback = any(callback).(OnAfterReadInterface[Concern])
	case *ConcernCompositionShape:
		stage.OnAfterConcernCompositionShapeReadCallback = any(callback).(OnAfterReadInterface[ConcernCompositionShape])
	case *ConcernInputShape:
		stage.OnAfterConcernInputShapeReadCallback = any(callback).(OnAfterReadInterface[ConcernInputShape])
	case *ConcernOutputShape:
		stage.OnAfterConcernOutputShapeReadCallback = any(callback).(OnAfterReadInterface[ConcernOutputShape])
	case *ConcernShape:
		stage.OnAfterConcernShapeReadCallback = any(callback).(OnAfterReadInterface[ConcernShape])
	case *Deliverable:
		stage.OnAfterDeliverableReadCallback = any(callback).(OnAfterReadInterface[Deliverable])
	case *Diagram:
		stage.OnAfterDiagramReadCallback = any(callback).(OnAfterReadInterface[Diagram])
	case *Library:
		stage.OnAfterLibraryReadCallback = any(callback).(OnAfterReadInterface[Library])
	case *Note:
		stage.OnAfterNoteReadCallback = any(callback).(OnAfterReadInterface[Note])
	case *NoteProductShape:
		stage.OnAfterNoteProductShapeReadCallback = any(callback).(OnAfterReadInterface[NoteProductShape])
	case *NoteShape:
		stage.OnAfterNoteShapeReadCallback = any(callback).(OnAfterReadInterface[NoteShape])
	case *NoteStakeholderShape:
		stage.OnAfterNoteStakeholderShapeReadCallback = any(callback).(OnAfterReadInterface[NoteStakeholderShape])
	case *NoteTaskShape:
		stage.OnAfterNoteTaskShapeReadCallback = any(callback).(OnAfterReadInterface[NoteTaskShape])
	case *ProductCompositionShape:
		stage.OnAfterProductCompositionShapeReadCallback = any(callback).(OnAfterReadInterface[ProductCompositionShape])
	case *ProductShape:
		stage.OnAfterProductShapeReadCallback = any(callback).(OnAfterReadInterface[ProductShape])
	case *Requirement:
		stage.OnAfterRequirementReadCallback = any(callback).(OnAfterReadInterface[Requirement])
	case *RequirementShape:
		stage.OnAfterRequirementShapeReadCallback = any(callback).(OnAfterReadInterface[RequirementShape])
	case *Stakeholder:
		stage.OnAfterStakeholderReadCallback = any(callback).(OnAfterReadInterface[Stakeholder])
	case *StakeholderCompositionShape:
		stage.OnAfterStakeholderCompositionShapeReadCallback = any(callback).(OnAfterReadInterface[StakeholderCompositionShape])
	case *StakeholderConcernShape:
		stage.OnAfterStakeholderConcernShapeReadCallback = any(callback).(OnAfterReadInterface[StakeholderConcernShape])
	case *StakeholderShape:
		stage.OnAfterStakeholderShapeReadCallback = any(callback).(OnAfterReadInterface[StakeholderShape])
	case *SupportLevel:
		stage.OnAfterSupportLevelReadCallback = any(callback).(OnAfterReadInterface[SupportLevel])
	case *Tool:
		stage.OnAfterToolReadCallback = any(callback).(OnAfterReadInterface[Tool])
	}
}
