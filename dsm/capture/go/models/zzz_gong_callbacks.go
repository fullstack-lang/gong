// generated code - do not edit
package models

// AfterCreateFromFront is the Stage method called after a create from front.
func (stage *Stage) AfterCreateFromFront[Type Gongstruct](instance *Type) {

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
	case *ControlPointShape:
		if stage.OnAfterControlPointShapeCreateCallback != nil {
			stage.OnAfterControlPointShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Deliverable:
		if stage.OnAfterDeliverableCreateCallback != nil {
			stage.OnAfterDeliverableCreateCallback.OnAfterCreate(stage, target)
		}
	case *DeliverableCompositionShape:
		if stage.OnAfterDeliverableCompositionShapeCreateCallback != nil {
			stage.OnAfterDeliverableCompositionShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *DeliverableConceptShape:
		if stage.OnAfterDeliverableConceptShapeCreateCallback != nil {
			stage.OnAfterDeliverableConceptShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *DeliverableShape:
		if stage.OnAfterDeliverableShapeCreateCallback != nil {
			stage.OnAfterDeliverableShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Diagram:
		if stage.OnAfterDiagramCreateCallback != nil {
			stage.OnAfterDiagramCreateCallback.OnAfterCreate(stage, target)
		}
	case *DiagramShape:
		if stage.OnAfterDiagramShapeCreateCallback != nil {
			stage.OnAfterDiagramShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Library:
		if stage.OnAfterLibraryCreateCallback != nil {
			stage.OnAfterLibraryCreateCallback.OnAfterCreate(stage, target)
		}
	case *Note:
		if stage.OnAfterNoteCreateCallback != nil {
			stage.OnAfterNoteCreateCallback.OnAfterCreate(stage, target)
		}
	case *NoteDeliverableShape:
		if stage.OnAfterNoteDeliverableShapeCreateCallback != nil {
			stage.OnAfterNoteDeliverableShapeCreateCallback.OnAfterCreate(stage, target)
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

// AfterCreateFromFront is a backward-compatible package-level forwarder.
func AfterCreateFromFront[Type Gongstruct](stage *Stage, instance *Type) {
	stage.AfterCreateFromFront(instance)
}

type Gong__MouseEvent struct {
	ShiftKey bool
}

// OnAfterUpdateFromFront is the Stage method called after an update from front.
func (stage *Stage) OnAfterUpdateFromFront[Type Gongstruct](old, new *Type) {

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
	case *ControlPointShape:
		newTarget := any(new).(*ControlPointShape)
		if stage.OnAfterControlPointShapeUpdateCallback != nil {
			stage.OnAfterControlPointShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Deliverable:
		newTarget := any(new).(*Deliverable)
		if stage.OnAfterDeliverableUpdateCallback != nil {
			stage.OnAfterDeliverableUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *DeliverableCompositionShape:
		newTarget := any(new).(*DeliverableCompositionShape)
		if stage.OnAfterDeliverableCompositionShapeUpdateCallback != nil {
			stage.OnAfterDeliverableCompositionShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *DeliverableConceptShape:
		newTarget := any(new).(*DeliverableConceptShape)
		if stage.OnAfterDeliverableConceptShapeUpdateCallback != nil {
			stage.OnAfterDeliverableConceptShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *DeliverableShape:
		newTarget := any(new).(*DeliverableShape)
		if stage.OnAfterDeliverableShapeUpdateCallback != nil {
			stage.OnAfterDeliverableShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Diagram:
		newTarget := any(new).(*Diagram)
		if stage.OnAfterDiagramUpdateCallback != nil {
			stage.OnAfterDiagramUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *DiagramShape:
		newTarget := any(new).(*DiagramShape)
		if stage.OnAfterDiagramShapeUpdateCallback != nil {
			stage.OnAfterDiagramShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
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
	case *NoteDeliverableShape:
		newTarget := any(new).(*NoteDeliverableShape)
		if stage.OnAfterNoteDeliverableShapeUpdateCallback != nil {
			stage.OnAfterNoteDeliverableShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
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

// OnAfterUpdateFromFront is a backward-compatible package-level forwarder.
func OnAfterUpdateFromFront[Type Gongstruct](stage *Stage, old, new *Type) {
	stage.OnAfterUpdateFromFront(old, new)
}

// AfterDeleteFromFront is the Stage method called after a delete from front.
func (stage *Stage) AfterDeleteFromFront[Type Gongstruct](staged, front *Type) {

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
	case *ControlPointShape:
		if stage.OnAfterControlPointShapeDeleteCallback != nil {
			staged := any(staged).(*ControlPointShape)
			stage.OnAfterControlPointShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Deliverable:
		if stage.OnAfterDeliverableDeleteCallback != nil {
			staged := any(staged).(*Deliverable)
			stage.OnAfterDeliverableDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *DeliverableCompositionShape:
		if stage.OnAfterDeliverableCompositionShapeDeleteCallback != nil {
			staged := any(staged).(*DeliverableCompositionShape)
			stage.OnAfterDeliverableCompositionShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *DeliverableConceptShape:
		if stage.OnAfterDeliverableConceptShapeDeleteCallback != nil {
			staged := any(staged).(*DeliverableConceptShape)
			stage.OnAfterDeliverableConceptShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *DeliverableShape:
		if stage.OnAfterDeliverableShapeDeleteCallback != nil {
			staged := any(staged).(*DeliverableShape)
			stage.OnAfterDeliverableShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Diagram:
		if stage.OnAfterDiagramDeleteCallback != nil {
			staged := any(staged).(*Diagram)
			stage.OnAfterDiagramDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *DiagramShape:
		if stage.OnAfterDiagramShapeDeleteCallback != nil {
			staged := any(staged).(*DiagramShape)
			stage.OnAfterDiagramShapeDeleteCallback.OnAfterDelete(stage, staged, front)
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
	case *NoteDeliverableShape:
		if stage.OnAfterNoteDeliverableShapeDeleteCallback != nil {
			staged := any(staged).(*NoteDeliverableShape)
			stage.OnAfterNoteDeliverableShapeDeleteCallback.OnAfterDelete(stage, staged, front)
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

// AfterDeleteFromFront is a backward-compatible package-level forwarder.
func AfterDeleteFromFront[Type Gongstruct](stage *Stage, staged, front *Type) {
	stage.AfterDeleteFromFront(staged, front)
}
