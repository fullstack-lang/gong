// generated code - do not edit
package models

// AfterCreateFromFront is the Stage method called after a create from front.
func (stage *Stage) AfterCreateFromFront[Type Gongstruct](instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
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
	case *NoteResourceShape:
		if stage.OnAfterNoteResourceShapeCreateCallback != nil {
			stage.OnAfterNoteResourceShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *NoteShape:
		if stage.OnAfterNoteShapeCreateCallback != nil {
			stage.OnAfterNoteShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *NoteTaskShape:
		if stage.OnAfterNoteTaskShapeCreateCallback != nil {
			stage.OnAfterNoteTaskShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Product:
		if stage.OnAfterProductCreateCallback != nil {
			stage.OnAfterProductCreateCallback.OnAfterCreate(stage, target)
		}
	case *ProductCompositionShape:
		if stage.OnAfterProductCompositionShapeCreateCallback != nil {
			stage.OnAfterProductCompositionShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *ProductShape:
		if stage.OnAfterProductShapeCreateCallback != nil {
			stage.OnAfterProductShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Resource:
		if stage.OnAfterResourceCreateCallback != nil {
			stage.OnAfterResourceCreateCallback.OnAfterCreate(stage, target)
		}
	case *ResourceCompositionShape:
		if stage.OnAfterResourceCompositionShapeCreateCallback != nil {
			stage.OnAfterResourceCompositionShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *ResourceShape:
		if stage.OnAfterResourceShapeCreateCallback != nil {
			stage.OnAfterResourceShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *ResourceTaskShape:
		if stage.OnAfterResourceTaskShapeCreateCallback != nil {
			stage.OnAfterResourceTaskShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Task:
		if stage.OnAfterTaskCreateCallback != nil {
			stage.OnAfterTaskCreateCallback.OnAfterCreate(stage, target)
		}
	case *TaskCompositionShape:
		if stage.OnAfterTaskCompositionShapeCreateCallback != nil {
			stage.OnAfterTaskCompositionShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *TaskGroup:
		if stage.OnAfterTaskGroupCreateCallback != nil {
			stage.OnAfterTaskGroupCreateCallback.OnAfterCreate(stage, target)
		}
	case *TaskGroupShape:
		if stage.OnAfterTaskGroupShapeCreateCallback != nil {
			stage.OnAfterTaskGroupShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *TaskInputShape:
		if stage.OnAfterTaskInputShapeCreateCallback != nil {
			stage.OnAfterTaskInputShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *TaskOutputShape:
		if stage.OnAfterTaskOutputShapeCreateCallback != nil {
			stage.OnAfterTaskOutputShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *TaskPredecessorShape:
		if stage.OnAfterTaskPredecessorShapeCreateCallback != nil {
			stage.OnAfterTaskPredecessorShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *TaskShape:
		if stage.OnAfterTaskShapeCreateCallback != nil {
			stage.OnAfterTaskShapeCreateCallback.OnAfterCreate(stage, target)
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
	case *NoteResourceShape:
		newTarget := any(new).(*NoteResourceShape)
		if stage.OnAfterNoteResourceShapeUpdateCallback != nil {
			stage.OnAfterNoteResourceShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *NoteShape:
		newTarget := any(new).(*NoteShape)
		if stage.OnAfterNoteShapeUpdateCallback != nil {
			stage.OnAfterNoteShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *NoteTaskShape:
		newTarget := any(new).(*NoteTaskShape)
		if stage.OnAfterNoteTaskShapeUpdateCallback != nil {
			stage.OnAfterNoteTaskShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Product:
		newTarget := any(new).(*Product)
		if stage.OnAfterProductUpdateCallback != nil {
			stage.OnAfterProductUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
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
	case *Resource:
		newTarget := any(new).(*Resource)
		if stage.OnAfterResourceUpdateCallback != nil {
			stage.OnAfterResourceUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ResourceCompositionShape:
		newTarget := any(new).(*ResourceCompositionShape)
		if stage.OnAfterResourceCompositionShapeUpdateCallback != nil {
			stage.OnAfterResourceCompositionShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ResourceShape:
		newTarget := any(new).(*ResourceShape)
		if stage.OnAfterResourceShapeUpdateCallback != nil {
			stage.OnAfterResourceShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ResourceTaskShape:
		newTarget := any(new).(*ResourceTaskShape)
		if stage.OnAfterResourceTaskShapeUpdateCallback != nil {
			stage.OnAfterResourceTaskShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Task:
		newTarget := any(new).(*Task)
		if stage.OnAfterTaskUpdateCallback != nil {
			stage.OnAfterTaskUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *TaskCompositionShape:
		newTarget := any(new).(*TaskCompositionShape)
		if stage.OnAfterTaskCompositionShapeUpdateCallback != nil {
			stage.OnAfterTaskCompositionShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *TaskGroup:
		newTarget := any(new).(*TaskGroup)
		if stage.OnAfterTaskGroupUpdateCallback != nil {
			stage.OnAfterTaskGroupUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *TaskGroupShape:
		newTarget := any(new).(*TaskGroupShape)
		if stage.OnAfterTaskGroupShapeUpdateCallback != nil {
			stage.OnAfterTaskGroupShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *TaskInputShape:
		newTarget := any(new).(*TaskInputShape)
		if stage.OnAfterTaskInputShapeUpdateCallback != nil {
			stage.OnAfterTaskInputShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *TaskOutputShape:
		newTarget := any(new).(*TaskOutputShape)
		if stage.OnAfterTaskOutputShapeUpdateCallback != nil {
			stage.OnAfterTaskOutputShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *TaskPredecessorShape:
		newTarget := any(new).(*TaskPredecessorShape)
		if stage.OnAfterTaskPredecessorShapeUpdateCallback != nil {
			stage.OnAfterTaskPredecessorShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *TaskShape:
		newTarget := any(new).(*TaskShape)
		if stage.OnAfterTaskShapeUpdateCallback != nil {
			stage.OnAfterTaskShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
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
	case *NoteResourceShape:
		if stage.OnAfterNoteResourceShapeDeleteCallback != nil {
			staged := any(staged).(*NoteResourceShape)
			stage.OnAfterNoteResourceShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *NoteShape:
		if stage.OnAfterNoteShapeDeleteCallback != nil {
			staged := any(staged).(*NoteShape)
			stage.OnAfterNoteShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *NoteTaskShape:
		if stage.OnAfterNoteTaskShapeDeleteCallback != nil {
			staged := any(staged).(*NoteTaskShape)
			stage.OnAfterNoteTaskShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Product:
		if stage.OnAfterProductDeleteCallback != nil {
			staged := any(staged).(*Product)
			stage.OnAfterProductDeleteCallback.OnAfterDelete(stage, staged, front)
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
	case *Resource:
		if stage.OnAfterResourceDeleteCallback != nil {
			staged := any(staged).(*Resource)
			stage.OnAfterResourceDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ResourceCompositionShape:
		if stage.OnAfterResourceCompositionShapeDeleteCallback != nil {
			staged := any(staged).(*ResourceCompositionShape)
			stage.OnAfterResourceCompositionShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ResourceShape:
		if stage.OnAfterResourceShapeDeleteCallback != nil {
			staged := any(staged).(*ResourceShape)
			stage.OnAfterResourceShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ResourceTaskShape:
		if stage.OnAfterResourceTaskShapeDeleteCallback != nil {
			staged := any(staged).(*ResourceTaskShape)
			stage.OnAfterResourceTaskShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Task:
		if stage.OnAfterTaskDeleteCallback != nil {
			staged := any(staged).(*Task)
			stage.OnAfterTaskDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *TaskCompositionShape:
		if stage.OnAfterTaskCompositionShapeDeleteCallback != nil {
			staged := any(staged).(*TaskCompositionShape)
			stage.OnAfterTaskCompositionShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *TaskGroup:
		if stage.OnAfterTaskGroupDeleteCallback != nil {
			staged := any(staged).(*TaskGroup)
			stage.OnAfterTaskGroupDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *TaskGroupShape:
		if stage.OnAfterTaskGroupShapeDeleteCallback != nil {
			staged := any(staged).(*TaskGroupShape)
			stage.OnAfterTaskGroupShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *TaskInputShape:
		if stage.OnAfterTaskInputShapeDeleteCallback != nil {
			staged := any(staged).(*TaskInputShape)
			stage.OnAfterTaskInputShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *TaskOutputShape:
		if stage.OnAfterTaskOutputShapeDeleteCallback != nil {
			staged := any(staged).(*TaskOutputShape)
			stage.OnAfterTaskOutputShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *TaskPredecessorShape:
		if stage.OnAfterTaskPredecessorShapeDeleteCallback != nil {
			staged := any(staged).(*TaskPredecessorShape)
			stage.OnAfterTaskPredecessorShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *TaskShape:
		if stage.OnAfterTaskShapeDeleteCallback != nil {
			staged := any(staged).(*TaskShape)
			stage.OnAfterTaskShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	default:
		_ = front
	}
}

// AfterDeleteFromFront is a backward-compatible package-level forwarder.
func AfterDeleteFromFront[Type Gongstruct](stage *Stage, staged, front *Type) {
	stage.AfterDeleteFromFront(staged, front)
}
