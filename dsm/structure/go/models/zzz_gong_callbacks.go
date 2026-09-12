// generated code - do not edit
package models

// AfterCreateFromFront is the Stage method called after a create from front.
func (stage *Stage) AfterCreateFromFront[Type Gongstruct](instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *AllocatedResourceShape:
		if stage.OnAfterAllocatedResourceShapeCreateCallback != nil {
			stage.OnAfterAllocatedResourceShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *AllocatedSystemShape:
		if stage.OnAfterAllocatedSystemShapeCreateCallback != nil {
			stage.OnAfterAllocatedSystemShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *ControlFlow:
		if stage.OnAfterControlFlowCreateCallback != nil {
			stage.OnAfterControlFlowCreateCallback.OnAfterCreate(stage, target)
		}
	case *ControlFlowShape:
		if stage.OnAfterControlFlowShapeCreateCallback != nil {
			stage.OnAfterControlFlowShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Data:
		if stage.OnAfterDataCreateCallback != nil {
			stage.OnAfterDataCreateCallback.OnAfterCreate(stage, target)
		}
	case *DataFlow:
		if stage.OnAfterDataFlowCreateCallback != nil {
			stage.OnAfterDataFlowCreateCallback.OnAfterCreate(stage, target)
		}
	case *DataFlowShape:
		if stage.OnAfterDataFlowShapeCreateCallback != nil {
			stage.OnAfterDataFlowShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *DataShape:
		if stage.OnAfterDataShapeCreateCallback != nil {
			stage.OnAfterDataShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *DiagramLayerState:
		if stage.OnAfterDiagramLayerStateCreateCallback != nil {
			stage.OnAfterDiagramLayerStateCreateCallback.OnAfterCreate(stage, target)
		}
	case *DiagramStructure:
		if stage.OnAfterDiagramStructureCreateCallback != nil {
			stage.OnAfterDiagramStructureCreateCallback.OnAfterCreate(stage, target)
		}
	case *ExternalPartShape:
		if stage.OnAfterExternalPartShapeCreateCallback != nil {
			stage.OnAfterExternalPartShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *LayerDefinition:
		if stage.OnAfterLayerDefinitionCreateCallback != nil {
			stage.OnAfterLayerDefinitionCreateCallback.OnAfterCreate(stage, target)
		}
	case *Library:
		if stage.OnAfterLibraryCreateCallback != nil {
			stage.OnAfterLibraryCreateCallback.OnAfterCreate(stage, target)
		}
	case *Note:
		if stage.OnAfterNoteCreateCallback != nil {
			stage.OnAfterNoteCreateCallback.OnAfterCreate(stage, target)
		}
	case *NotePartShape:
		if stage.OnAfterNotePartShapeCreateCallback != nil {
			stage.OnAfterNotePartShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *NotePortShape:
		if stage.OnAfterNotePortShapeCreateCallback != nil {
			stage.OnAfterNotePortShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *NoteShape:
		if stage.OnAfterNoteShapeCreateCallback != nil {
			stage.OnAfterNoteShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Part:
		if stage.OnAfterPartCreateCallback != nil {
			stage.OnAfterPartCreateCallback.OnAfterCreate(stage, target)
		}
	case *PartAnchoredPath:
		if stage.OnAfterPartAnchoredPathCreateCallback != nil {
			stage.OnAfterPartAnchoredPathCreateCallback.OnAfterCreate(stage, target)
		}
	case *PartShape:
		if stage.OnAfterPartShapeCreateCallback != nil {
			stage.OnAfterPartShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Port:
		if stage.OnAfterPortCreateCallback != nil {
			stage.OnAfterPortCreateCallback.OnAfterCreate(stage, target)
		}
	case *PortShape:
		if stage.OnAfterPortShapeCreateCallback != nil {
			stage.OnAfterPortShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Resource:
		if stage.OnAfterResourceCreateCallback != nil {
			stage.OnAfterResourceCreateCallback.OnAfterCreate(stage, target)
		}
	case *SemanticTag:
		if stage.OnAfterSemanticTagCreateCallback != nil {
			stage.OnAfterSemanticTagCreateCallback.OnAfterCreate(stage, target)
		}
	case *System:
		if stage.OnAfterSystemCreateCallback != nil {
			stage.OnAfterSystemCreateCallback.OnAfterCreate(stage, target)
		}
	case *SystemShape:
		if stage.OnAfterSystemShapeCreateCallback != nil {
			stage.OnAfterSystemShapeCreateCallback.OnAfterCreate(stage, target)
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
	case *AllocatedResourceShape:
		newTarget := any(new).(*AllocatedResourceShape)
		if stage.OnAfterAllocatedResourceShapeUpdateCallback != nil {
			stage.OnAfterAllocatedResourceShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *AllocatedSystemShape:
		newTarget := any(new).(*AllocatedSystemShape)
		if stage.OnAfterAllocatedSystemShapeUpdateCallback != nil {
			stage.OnAfterAllocatedSystemShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ControlFlow:
		newTarget := any(new).(*ControlFlow)
		if stage.OnAfterControlFlowUpdateCallback != nil {
			stage.OnAfterControlFlowUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ControlFlowShape:
		newTarget := any(new).(*ControlFlowShape)
		if stage.OnAfterControlFlowShapeUpdateCallback != nil {
			stage.OnAfterControlFlowShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Data:
		newTarget := any(new).(*Data)
		if stage.OnAfterDataUpdateCallback != nil {
			stage.OnAfterDataUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *DataFlow:
		newTarget := any(new).(*DataFlow)
		if stage.OnAfterDataFlowUpdateCallback != nil {
			stage.OnAfterDataFlowUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *DataFlowShape:
		newTarget := any(new).(*DataFlowShape)
		if stage.OnAfterDataFlowShapeUpdateCallback != nil {
			stage.OnAfterDataFlowShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *DataShape:
		newTarget := any(new).(*DataShape)
		if stage.OnAfterDataShapeUpdateCallback != nil {
			stage.OnAfterDataShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *DiagramLayerState:
		newTarget := any(new).(*DiagramLayerState)
		if stage.OnAfterDiagramLayerStateUpdateCallback != nil {
			stage.OnAfterDiagramLayerStateUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *DiagramStructure:
		newTarget := any(new).(*DiagramStructure)
		if stage.OnAfterDiagramStructureUpdateCallback != nil {
			stage.OnAfterDiagramStructureUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ExternalPartShape:
		newTarget := any(new).(*ExternalPartShape)
		if stage.OnAfterExternalPartShapeUpdateCallback != nil {
			stage.OnAfterExternalPartShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *LayerDefinition:
		newTarget := any(new).(*LayerDefinition)
		if stage.OnAfterLayerDefinitionUpdateCallback != nil {
			stage.OnAfterLayerDefinitionUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
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
	case *NotePartShape:
		newTarget := any(new).(*NotePartShape)
		if stage.OnAfterNotePartShapeUpdateCallback != nil {
			stage.OnAfterNotePartShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *NotePortShape:
		newTarget := any(new).(*NotePortShape)
		if stage.OnAfterNotePortShapeUpdateCallback != nil {
			stage.OnAfterNotePortShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *NoteShape:
		newTarget := any(new).(*NoteShape)
		if stage.OnAfterNoteShapeUpdateCallback != nil {
			stage.OnAfterNoteShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Part:
		newTarget := any(new).(*Part)
		if stage.OnAfterPartUpdateCallback != nil {
			stage.OnAfterPartUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *PartAnchoredPath:
		newTarget := any(new).(*PartAnchoredPath)
		if stage.OnAfterPartAnchoredPathUpdateCallback != nil {
			stage.OnAfterPartAnchoredPathUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *PartShape:
		newTarget := any(new).(*PartShape)
		if stage.OnAfterPartShapeUpdateCallback != nil {
			stage.OnAfterPartShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Port:
		newTarget := any(new).(*Port)
		if stage.OnAfterPortUpdateCallback != nil {
			stage.OnAfterPortUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *PortShape:
		newTarget := any(new).(*PortShape)
		if stage.OnAfterPortShapeUpdateCallback != nil {
			stage.OnAfterPortShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Resource:
		newTarget := any(new).(*Resource)
		if stage.OnAfterResourceUpdateCallback != nil {
			stage.OnAfterResourceUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *SemanticTag:
		newTarget := any(new).(*SemanticTag)
		if stage.OnAfterSemanticTagUpdateCallback != nil {
			stage.OnAfterSemanticTagUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *System:
		newTarget := any(new).(*System)
		if stage.OnAfterSystemUpdateCallback != nil {
			stage.OnAfterSystemUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *SystemShape:
		newTarget := any(new).(*SystemShape)
		if stage.OnAfterSystemShapeUpdateCallback != nil {
			stage.OnAfterSystemShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
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
	case *AllocatedResourceShape:
		if stage.OnAfterAllocatedResourceShapeDeleteCallback != nil {
			staged := any(staged).(*AllocatedResourceShape)
			stage.OnAfterAllocatedResourceShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *AllocatedSystemShape:
		if stage.OnAfterAllocatedSystemShapeDeleteCallback != nil {
			staged := any(staged).(*AllocatedSystemShape)
			stage.OnAfterAllocatedSystemShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ControlFlow:
		if stage.OnAfterControlFlowDeleteCallback != nil {
			staged := any(staged).(*ControlFlow)
			stage.OnAfterControlFlowDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ControlFlowShape:
		if stage.OnAfterControlFlowShapeDeleteCallback != nil {
			staged := any(staged).(*ControlFlowShape)
			stage.OnAfterControlFlowShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Data:
		if stage.OnAfterDataDeleteCallback != nil {
			staged := any(staged).(*Data)
			stage.OnAfterDataDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *DataFlow:
		if stage.OnAfterDataFlowDeleteCallback != nil {
			staged := any(staged).(*DataFlow)
			stage.OnAfterDataFlowDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *DataFlowShape:
		if stage.OnAfterDataFlowShapeDeleteCallback != nil {
			staged := any(staged).(*DataFlowShape)
			stage.OnAfterDataFlowShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *DataShape:
		if stage.OnAfterDataShapeDeleteCallback != nil {
			staged := any(staged).(*DataShape)
			stage.OnAfterDataShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *DiagramLayerState:
		if stage.OnAfterDiagramLayerStateDeleteCallback != nil {
			staged := any(staged).(*DiagramLayerState)
			stage.OnAfterDiagramLayerStateDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *DiagramStructure:
		if stage.OnAfterDiagramStructureDeleteCallback != nil {
			staged := any(staged).(*DiagramStructure)
			stage.OnAfterDiagramStructureDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ExternalPartShape:
		if stage.OnAfterExternalPartShapeDeleteCallback != nil {
			staged := any(staged).(*ExternalPartShape)
			stage.OnAfterExternalPartShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *LayerDefinition:
		if stage.OnAfterLayerDefinitionDeleteCallback != nil {
			staged := any(staged).(*LayerDefinition)
			stage.OnAfterLayerDefinitionDeleteCallback.OnAfterDelete(stage, staged, front)
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
	case *NotePartShape:
		if stage.OnAfterNotePartShapeDeleteCallback != nil {
			staged := any(staged).(*NotePartShape)
			stage.OnAfterNotePartShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *NotePortShape:
		if stage.OnAfterNotePortShapeDeleteCallback != nil {
			staged := any(staged).(*NotePortShape)
			stage.OnAfterNotePortShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *NoteShape:
		if stage.OnAfterNoteShapeDeleteCallback != nil {
			staged := any(staged).(*NoteShape)
			stage.OnAfterNoteShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Part:
		if stage.OnAfterPartDeleteCallback != nil {
			staged := any(staged).(*Part)
			stage.OnAfterPartDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *PartAnchoredPath:
		if stage.OnAfterPartAnchoredPathDeleteCallback != nil {
			staged := any(staged).(*PartAnchoredPath)
			stage.OnAfterPartAnchoredPathDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *PartShape:
		if stage.OnAfterPartShapeDeleteCallback != nil {
			staged := any(staged).(*PartShape)
			stage.OnAfterPartShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Port:
		if stage.OnAfterPortDeleteCallback != nil {
			staged := any(staged).(*Port)
			stage.OnAfterPortDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *PortShape:
		if stage.OnAfterPortShapeDeleteCallback != nil {
			staged := any(staged).(*PortShape)
			stage.OnAfterPortShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Resource:
		if stage.OnAfterResourceDeleteCallback != nil {
			staged := any(staged).(*Resource)
			stage.OnAfterResourceDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *SemanticTag:
		if stage.OnAfterSemanticTagDeleteCallback != nil {
			staged := any(staged).(*SemanticTag)
			stage.OnAfterSemanticTagDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *System:
		if stage.OnAfterSystemDeleteCallback != nil {
			staged := any(staged).(*System)
			stage.OnAfterSystemDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *SystemShape:
		if stage.OnAfterSystemShapeDeleteCallback != nil {
			staged := any(staged).(*SystemShape)
			stage.OnAfterSystemShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	default:
		_ = front
	}
}

// AfterDeleteFromFront is a backward-compatible package-level forwarder.
func AfterDeleteFromFront[Type Gongstruct](stage *Stage, staged, front *Type) {
	stage.AfterDeleteFromFront(staged, front)
}
