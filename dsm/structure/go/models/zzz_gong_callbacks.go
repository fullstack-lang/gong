// generated code - do not edit
package models

// AfterCreateFromFront is called after a create from front
func AfterCreateFromFront[Type Gongstruct](stage *Stage, instance *Type) {

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
	case *DiagramStructure:
		if stage.OnAfterDiagramStructureCreateCallback != nil {
			stage.OnAfterDiagramStructureCreateCallback.OnAfterCreate(stage, target)
		}
	case *ExternalPartShape:
		if stage.OnAfterExternalPartShapeCreateCallback != nil {
			stage.OnAfterExternalPartShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Library:
		if stage.OnAfterLibraryCreateCallback != nil {
			stage.OnAfterLibraryCreateCallback.OnAfterCreate(stage, target)
		}
	case *Note:
		if stage.OnAfterNoteCreateCallback != nil {
			stage.OnAfterNoteCreateCallback.OnAfterCreate(stage, target)
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

type Gong__MouseEvent struct {
	ShiftKey bool
}

// OnAfterUpdateFromFront is called after a update from front
func OnAfterUpdateFromFront[Type Gongstruct](stage *Stage, old, new *Type) {

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

// AfterDeleteFromFront is called after a delete from front
func AfterDeleteFromFront[Type Gongstruct](stage *Stage, staged, front *Type) {

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

// AfterReadFromFront is called after a Read from front
func AfterReadFromFront[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *AllocatedResourceShape:
		if stage.OnAfterAllocatedResourceShapeReadCallback != nil {
			stage.OnAfterAllocatedResourceShapeReadCallback.OnAfterRead(stage, target)
		}
	case *AllocatedSystemShape:
		if stage.OnAfterAllocatedSystemShapeReadCallback != nil {
			stage.OnAfterAllocatedSystemShapeReadCallback.OnAfterRead(stage, target)
		}
	case *ControlFlow:
		if stage.OnAfterControlFlowReadCallback != nil {
			stage.OnAfterControlFlowReadCallback.OnAfterRead(stage, target)
		}
	case *ControlFlowShape:
		if stage.OnAfterControlFlowShapeReadCallback != nil {
			stage.OnAfterControlFlowShapeReadCallback.OnAfterRead(stage, target)
		}
	case *Data:
		if stage.OnAfterDataReadCallback != nil {
			stage.OnAfterDataReadCallback.OnAfterRead(stage, target)
		}
	case *DataFlow:
		if stage.OnAfterDataFlowReadCallback != nil {
			stage.OnAfterDataFlowReadCallback.OnAfterRead(stage, target)
		}
	case *DataFlowShape:
		if stage.OnAfterDataFlowShapeReadCallback != nil {
			stage.OnAfterDataFlowShapeReadCallback.OnAfterRead(stage, target)
		}
	case *DataShape:
		if stage.OnAfterDataShapeReadCallback != nil {
			stage.OnAfterDataShapeReadCallback.OnAfterRead(stage, target)
		}
	case *DiagramStructure:
		if stage.OnAfterDiagramStructureReadCallback != nil {
			stage.OnAfterDiagramStructureReadCallback.OnAfterRead(stage, target)
		}
	case *ExternalPartShape:
		if stage.OnAfterExternalPartShapeReadCallback != nil {
			stage.OnAfterExternalPartShapeReadCallback.OnAfterRead(stage, target)
		}
	case *Library:
		if stage.OnAfterLibraryReadCallback != nil {
			stage.OnAfterLibraryReadCallback.OnAfterRead(stage, target)
		}
	case *Note:
		if stage.OnAfterNoteReadCallback != nil {
			stage.OnAfterNoteReadCallback.OnAfterRead(stage, target)
		}
	case *NotePortShape:
		if stage.OnAfterNotePortShapeReadCallback != nil {
			stage.OnAfterNotePortShapeReadCallback.OnAfterRead(stage, target)
		}
	case *NoteShape:
		if stage.OnAfterNoteShapeReadCallback != nil {
			stage.OnAfterNoteShapeReadCallback.OnAfterRead(stage, target)
		}
	case *Part:
		if stage.OnAfterPartReadCallback != nil {
			stage.OnAfterPartReadCallback.OnAfterRead(stage, target)
		}
	case *PartAnchoredPath:
		if stage.OnAfterPartAnchoredPathReadCallback != nil {
			stage.OnAfterPartAnchoredPathReadCallback.OnAfterRead(stage, target)
		}
	case *PartShape:
		if stage.OnAfterPartShapeReadCallback != nil {
			stage.OnAfterPartShapeReadCallback.OnAfterRead(stage, target)
		}
	case *Port:
		if stage.OnAfterPortReadCallback != nil {
			stage.OnAfterPortReadCallback.OnAfterRead(stage, target)
		}
	case *PortShape:
		if stage.OnAfterPortShapeReadCallback != nil {
			stage.OnAfterPortShapeReadCallback.OnAfterRead(stage, target)
		}
	case *Resource:
		if stage.OnAfterResourceReadCallback != nil {
			stage.OnAfterResourceReadCallback.OnAfterRead(stage, target)
		}
	case *System:
		if stage.OnAfterSystemReadCallback != nil {
			stage.OnAfterSystemReadCallback.OnAfterRead(stage, target)
		}
	case *SystemShape:
		if stage.OnAfterSystemShapeReadCallback != nil {
			stage.OnAfterSystemShapeReadCallback.OnAfterRead(stage, target)
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
	case *AllocatedResourceShape:
		stage.OnAfterAllocatedResourceShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[AllocatedResourceShape])
	case *AllocatedSystemShape:
		stage.OnAfterAllocatedSystemShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[AllocatedSystemShape])
	case *ControlFlow:
		stage.OnAfterControlFlowUpdateCallback = any(callback).(OnAfterUpdateInterface[ControlFlow])
	case *ControlFlowShape:
		stage.OnAfterControlFlowShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[ControlFlowShape])
	case *Data:
		stage.OnAfterDataUpdateCallback = any(callback).(OnAfterUpdateInterface[Data])
	case *DataFlow:
		stage.OnAfterDataFlowUpdateCallback = any(callback).(OnAfterUpdateInterface[DataFlow])
	case *DataFlowShape:
		stage.OnAfterDataFlowShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[DataFlowShape])
	case *DataShape:
		stage.OnAfterDataShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[DataShape])
	case *DiagramStructure:
		stage.OnAfterDiagramStructureUpdateCallback = any(callback).(OnAfterUpdateInterface[DiagramStructure])
	case *ExternalPartShape:
		stage.OnAfterExternalPartShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[ExternalPartShape])
	case *Library:
		stage.OnAfterLibraryUpdateCallback = any(callback).(OnAfterUpdateInterface[Library])
	case *Note:
		stage.OnAfterNoteUpdateCallback = any(callback).(OnAfterUpdateInterface[Note])
	case *NotePortShape:
		stage.OnAfterNotePortShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[NotePortShape])
	case *NoteShape:
		stage.OnAfterNoteShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[NoteShape])
	case *Part:
		stage.OnAfterPartUpdateCallback = any(callback).(OnAfterUpdateInterface[Part])
	case *PartAnchoredPath:
		stage.OnAfterPartAnchoredPathUpdateCallback = any(callback).(OnAfterUpdateInterface[PartAnchoredPath])
	case *PartShape:
		stage.OnAfterPartShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[PartShape])
	case *Port:
		stage.OnAfterPortUpdateCallback = any(callback).(OnAfterUpdateInterface[Port])
	case *PortShape:
		stage.OnAfterPortShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[PortShape])
	case *Resource:
		stage.OnAfterResourceUpdateCallback = any(callback).(OnAfterUpdateInterface[Resource])
	case *System:
		stage.OnAfterSystemUpdateCallback = any(callback).(OnAfterUpdateInterface[System])
	case *SystemShape:
		stage.OnAfterSystemShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[SystemShape])
	}
}
func SetCallbackAfterCreateFromFront[Type Gongstruct](stage *Stage, callback OnAfterCreateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
	// insertion point
	case *AllocatedResourceShape:
		stage.OnAfterAllocatedResourceShapeCreateCallback = any(callback).(OnAfterCreateInterface[AllocatedResourceShape])
	case *AllocatedSystemShape:
		stage.OnAfterAllocatedSystemShapeCreateCallback = any(callback).(OnAfterCreateInterface[AllocatedSystemShape])
	case *ControlFlow:
		stage.OnAfterControlFlowCreateCallback = any(callback).(OnAfterCreateInterface[ControlFlow])
	case *ControlFlowShape:
		stage.OnAfterControlFlowShapeCreateCallback = any(callback).(OnAfterCreateInterface[ControlFlowShape])
	case *Data:
		stage.OnAfterDataCreateCallback = any(callback).(OnAfterCreateInterface[Data])
	case *DataFlow:
		stage.OnAfterDataFlowCreateCallback = any(callback).(OnAfterCreateInterface[DataFlow])
	case *DataFlowShape:
		stage.OnAfterDataFlowShapeCreateCallback = any(callback).(OnAfterCreateInterface[DataFlowShape])
	case *DataShape:
		stage.OnAfterDataShapeCreateCallback = any(callback).(OnAfterCreateInterface[DataShape])
	case *DiagramStructure:
		stage.OnAfterDiagramStructureCreateCallback = any(callback).(OnAfterCreateInterface[DiagramStructure])
	case *ExternalPartShape:
		stage.OnAfterExternalPartShapeCreateCallback = any(callback).(OnAfterCreateInterface[ExternalPartShape])
	case *Library:
		stage.OnAfterLibraryCreateCallback = any(callback).(OnAfterCreateInterface[Library])
	case *Note:
		stage.OnAfterNoteCreateCallback = any(callback).(OnAfterCreateInterface[Note])
	case *NotePortShape:
		stage.OnAfterNotePortShapeCreateCallback = any(callback).(OnAfterCreateInterface[NotePortShape])
	case *NoteShape:
		stage.OnAfterNoteShapeCreateCallback = any(callback).(OnAfterCreateInterface[NoteShape])
	case *Part:
		stage.OnAfterPartCreateCallback = any(callback).(OnAfterCreateInterface[Part])
	case *PartAnchoredPath:
		stage.OnAfterPartAnchoredPathCreateCallback = any(callback).(OnAfterCreateInterface[PartAnchoredPath])
	case *PartShape:
		stage.OnAfterPartShapeCreateCallback = any(callback).(OnAfterCreateInterface[PartShape])
	case *Port:
		stage.OnAfterPortCreateCallback = any(callback).(OnAfterCreateInterface[Port])
	case *PortShape:
		stage.OnAfterPortShapeCreateCallback = any(callback).(OnAfterCreateInterface[PortShape])
	case *Resource:
		stage.OnAfterResourceCreateCallback = any(callback).(OnAfterCreateInterface[Resource])
	case *System:
		stage.OnAfterSystemCreateCallback = any(callback).(OnAfterCreateInterface[System])
	case *SystemShape:
		stage.OnAfterSystemShapeCreateCallback = any(callback).(OnAfterCreateInterface[SystemShape])
	}
}
func SetCallbackAfterDeleteFromFront[Type Gongstruct](stage *Stage, callback OnAfterDeleteInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
	// insertion point
	case *AllocatedResourceShape:
		stage.OnAfterAllocatedResourceShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[AllocatedResourceShape])
	case *AllocatedSystemShape:
		stage.OnAfterAllocatedSystemShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[AllocatedSystemShape])
	case *ControlFlow:
		stage.OnAfterControlFlowDeleteCallback = any(callback).(OnAfterDeleteInterface[ControlFlow])
	case *ControlFlowShape:
		stage.OnAfterControlFlowShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[ControlFlowShape])
	case *Data:
		stage.OnAfterDataDeleteCallback = any(callback).(OnAfterDeleteInterface[Data])
	case *DataFlow:
		stage.OnAfterDataFlowDeleteCallback = any(callback).(OnAfterDeleteInterface[DataFlow])
	case *DataFlowShape:
		stage.OnAfterDataFlowShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[DataFlowShape])
	case *DataShape:
		stage.OnAfterDataShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[DataShape])
	case *DiagramStructure:
		stage.OnAfterDiagramStructureDeleteCallback = any(callback).(OnAfterDeleteInterface[DiagramStructure])
	case *ExternalPartShape:
		stage.OnAfterExternalPartShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[ExternalPartShape])
	case *Library:
		stage.OnAfterLibraryDeleteCallback = any(callback).(OnAfterDeleteInterface[Library])
	case *Note:
		stage.OnAfterNoteDeleteCallback = any(callback).(OnAfterDeleteInterface[Note])
	case *NotePortShape:
		stage.OnAfterNotePortShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[NotePortShape])
	case *NoteShape:
		stage.OnAfterNoteShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[NoteShape])
	case *Part:
		stage.OnAfterPartDeleteCallback = any(callback).(OnAfterDeleteInterface[Part])
	case *PartAnchoredPath:
		stage.OnAfterPartAnchoredPathDeleteCallback = any(callback).(OnAfterDeleteInterface[PartAnchoredPath])
	case *PartShape:
		stage.OnAfterPartShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[PartShape])
	case *Port:
		stage.OnAfterPortDeleteCallback = any(callback).(OnAfterDeleteInterface[Port])
	case *PortShape:
		stage.OnAfterPortShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[PortShape])
	case *Resource:
		stage.OnAfterResourceDeleteCallback = any(callback).(OnAfterDeleteInterface[Resource])
	case *System:
		stage.OnAfterSystemDeleteCallback = any(callback).(OnAfterDeleteInterface[System])
	case *SystemShape:
		stage.OnAfterSystemShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[SystemShape])
	}
}
func SetCallbackAfterReadFromFront[Type Gongstruct](stage *Stage, callback OnAfterReadInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
	// insertion point
	case *AllocatedResourceShape:
		stage.OnAfterAllocatedResourceShapeReadCallback = any(callback).(OnAfterReadInterface[AllocatedResourceShape])
	case *AllocatedSystemShape:
		stage.OnAfterAllocatedSystemShapeReadCallback = any(callback).(OnAfterReadInterface[AllocatedSystemShape])
	case *ControlFlow:
		stage.OnAfterControlFlowReadCallback = any(callback).(OnAfterReadInterface[ControlFlow])
	case *ControlFlowShape:
		stage.OnAfterControlFlowShapeReadCallback = any(callback).(OnAfterReadInterface[ControlFlowShape])
	case *Data:
		stage.OnAfterDataReadCallback = any(callback).(OnAfterReadInterface[Data])
	case *DataFlow:
		stage.OnAfterDataFlowReadCallback = any(callback).(OnAfterReadInterface[DataFlow])
	case *DataFlowShape:
		stage.OnAfterDataFlowShapeReadCallback = any(callback).(OnAfterReadInterface[DataFlowShape])
	case *DataShape:
		stage.OnAfterDataShapeReadCallback = any(callback).(OnAfterReadInterface[DataShape])
	case *DiagramStructure:
		stage.OnAfterDiagramStructureReadCallback = any(callback).(OnAfterReadInterface[DiagramStructure])
	case *ExternalPartShape:
		stage.OnAfterExternalPartShapeReadCallback = any(callback).(OnAfterReadInterface[ExternalPartShape])
	case *Library:
		stage.OnAfterLibraryReadCallback = any(callback).(OnAfterReadInterface[Library])
	case *Note:
		stage.OnAfterNoteReadCallback = any(callback).(OnAfterReadInterface[Note])
	case *NotePortShape:
		stage.OnAfterNotePortShapeReadCallback = any(callback).(OnAfterReadInterface[NotePortShape])
	case *NoteShape:
		stage.OnAfterNoteShapeReadCallback = any(callback).(OnAfterReadInterface[NoteShape])
	case *Part:
		stage.OnAfterPartReadCallback = any(callback).(OnAfterReadInterface[Part])
	case *PartAnchoredPath:
		stage.OnAfterPartAnchoredPathReadCallback = any(callback).(OnAfterReadInterface[PartAnchoredPath])
	case *PartShape:
		stage.OnAfterPartShapeReadCallback = any(callback).(OnAfterReadInterface[PartShape])
	case *Port:
		stage.OnAfterPortReadCallback = any(callback).(OnAfterReadInterface[Port])
	case *PortShape:
		stage.OnAfterPortShapeReadCallback = any(callback).(OnAfterReadInterface[PortShape])
	case *Resource:
		stage.OnAfterResourceReadCallback = any(callback).(OnAfterReadInterface[Resource])
	case *System:
		stage.OnAfterSystemReadCallback = any(callback).(OnAfterReadInterface[System])
	case *SystemShape:
		stage.OnAfterSystemShapeReadCallback = any(callback).(OnAfterReadInterface[SystemShape])
	}
}
