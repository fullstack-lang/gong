// generated code - do not edit
package models

// AfterCreateFromFront is the Stage method called after a create from front.
func (stage *Stage) AfterCreateFromFront[Type Gongstruct](instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *Action:
		if stage.OnAfterActionCreateCallback != nil {
			stage.OnAfterActionCreateCallback.OnAfterCreate(stage, target)
		}
	case *Activities:
		if stage.OnAfterActivitiesCreateCallback != nil {
			stage.OnAfterActivitiesCreateCallback.OnAfterCreate(stage, target)
		}
	case *Diagram:
		if stage.OnAfterDiagramCreateCallback != nil {
			stage.OnAfterDiagramCreateCallback.OnAfterCreate(stage, target)
		}
	case *Guard:
		if stage.OnAfterGuardCreateCallback != nil {
			stage.OnAfterGuardCreateCallback.OnAfterCreate(stage, target)
		}
	case *Kill:
		if stage.OnAfterKillCreateCallback != nil {
			stage.OnAfterKillCreateCallback.OnAfterCreate(stage, target)
		}
	case *Library:
		if stage.OnAfterLibraryCreateCallback != nil {
			stage.OnAfterLibraryCreateCallback.OnAfterCreate(stage, target)
		}
	case *Message:
		if stage.OnAfterMessageCreateCallback != nil {
			stage.OnAfterMessageCreateCallback.OnAfterCreate(stage, target)
		}
	case *MessageType:
		if stage.OnAfterMessageTypeCreateCallback != nil {
			stage.OnAfterMessageTypeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Note:
		if stage.OnAfterNoteCreateCallback != nil {
			stage.OnAfterNoteCreateCallback.OnAfterCreate(stage, target)
		}
	case *NoteShape:
		if stage.OnAfterNoteShapeCreateCallback != nil {
			stage.OnAfterNoteShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *NoteStateShape:
		if stage.OnAfterNoteStateShapeCreateCallback != nil {
			stage.OnAfterNoteStateShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Object:
		if stage.OnAfterObjectCreateCallback != nil {
			stage.OnAfterObjectCreateCallback.OnAfterCreate(stage, target)
		}
	case *Role:
		if stage.OnAfterRoleCreateCallback != nil {
			stage.OnAfterRoleCreateCallback.OnAfterCreate(stage, target)
		}
	case *State:
		if stage.OnAfterStateCreateCallback != nil {
			stage.OnAfterStateCreateCallback.OnAfterCreate(stage, target)
		}
	case *StateMachine:
		if stage.OnAfterStateMachineCreateCallback != nil {
			stage.OnAfterStateMachineCreateCallback.OnAfterCreate(stage, target)
		}
	case *StateShape:
		if stage.OnAfterStateShapeCreateCallback != nil {
			stage.OnAfterStateShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Transition:
		if stage.OnAfterTransitionCreateCallback != nil {
			stage.OnAfterTransitionCreateCallback.OnAfterCreate(stage, target)
		}
	case *Transition_Shape:
		if stage.OnAfterTransition_ShapeCreateCallback != nil {
			stage.OnAfterTransition_ShapeCreateCallback.OnAfterCreate(stage, target)
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
	case *Action:
		newTarget := any(new).(*Action)
		if stage.OnAfterActionUpdateCallback != nil {
			stage.OnAfterActionUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Activities:
		newTarget := any(new).(*Activities)
		if stage.OnAfterActivitiesUpdateCallback != nil {
			stage.OnAfterActivitiesUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Diagram:
		newTarget := any(new).(*Diagram)
		if stage.OnAfterDiagramUpdateCallback != nil {
			stage.OnAfterDiagramUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Guard:
		newTarget := any(new).(*Guard)
		if stage.OnAfterGuardUpdateCallback != nil {
			stage.OnAfterGuardUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Kill:
		newTarget := any(new).(*Kill)
		if stage.OnAfterKillUpdateCallback != nil {
			stage.OnAfterKillUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Library:
		newTarget := any(new).(*Library)
		if stage.OnAfterLibraryUpdateCallback != nil {
			stage.OnAfterLibraryUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Message:
		newTarget := any(new).(*Message)
		if stage.OnAfterMessageUpdateCallback != nil {
			stage.OnAfterMessageUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *MessageType:
		newTarget := any(new).(*MessageType)
		if stage.OnAfterMessageTypeUpdateCallback != nil {
			stage.OnAfterMessageTypeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Note:
		newTarget := any(new).(*Note)
		if stage.OnAfterNoteUpdateCallback != nil {
			stage.OnAfterNoteUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *NoteShape:
		newTarget := any(new).(*NoteShape)
		if stage.OnAfterNoteShapeUpdateCallback != nil {
			stage.OnAfterNoteShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *NoteStateShape:
		newTarget := any(new).(*NoteStateShape)
		if stage.OnAfterNoteStateShapeUpdateCallback != nil {
			stage.OnAfterNoteStateShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Object:
		newTarget := any(new).(*Object)
		if stage.OnAfterObjectUpdateCallback != nil {
			stage.OnAfterObjectUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Role:
		newTarget := any(new).(*Role)
		if stage.OnAfterRoleUpdateCallback != nil {
			stage.OnAfterRoleUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *State:
		newTarget := any(new).(*State)
		if stage.OnAfterStateUpdateCallback != nil {
			stage.OnAfterStateUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *StateMachine:
		newTarget := any(new).(*StateMachine)
		if stage.OnAfterStateMachineUpdateCallback != nil {
			stage.OnAfterStateMachineUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *StateShape:
		newTarget := any(new).(*StateShape)
		if stage.OnAfterStateShapeUpdateCallback != nil {
			stage.OnAfterStateShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Transition:
		newTarget := any(new).(*Transition)
		if stage.OnAfterTransitionUpdateCallback != nil {
			stage.OnAfterTransitionUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Transition_Shape:
		newTarget := any(new).(*Transition_Shape)
		if stage.OnAfterTransition_ShapeUpdateCallback != nil {
			stage.OnAfterTransition_ShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
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
	case *Action:
		if stage.OnAfterActionDeleteCallback != nil {
			staged := any(staged).(*Action)
			stage.OnAfterActionDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Activities:
		if stage.OnAfterActivitiesDeleteCallback != nil {
			staged := any(staged).(*Activities)
			stage.OnAfterActivitiesDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Diagram:
		if stage.OnAfterDiagramDeleteCallback != nil {
			staged := any(staged).(*Diagram)
			stage.OnAfterDiagramDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Guard:
		if stage.OnAfterGuardDeleteCallback != nil {
			staged := any(staged).(*Guard)
			stage.OnAfterGuardDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Kill:
		if stage.OnAfterKillDeleteCallback != nil {
			staged := any(staged).(*Kill)
			stage.OnAfterKillDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Library:
		if stage.OnAfterLibraryDeleteCallback != nil {
			staged := any(staged).(*Library)
			stage.OnAfterLibraryDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Message:
		if stage.OnAfterMessageDeleteCallback != nil {
			staged := any(staged).(*Message)
			stage.OnAfterMessageDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *MessageType:
		if stage.OnAfterMessageTypeDeleteCallback != nil {
			staged := any(staged).(*MessageType)
			stage.OnAfterMessageTypeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Note:
		if stage.OnAfterNoteDeleteCallback != nil {
			staged := any(staged).(*Note)
			stage.OnAfterNoteDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *NoteShape:
		if stage.OnAfterNoteShapeDeleteCallback != nil {
			staged := any(staged).(*NoteShape)
			stage.OnAfterNoteShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *NoteStateShape:
		if stage.OnAfterNoteStateShapeDeleteCallback != nil {
			staged := any(staged).(*NoteStateShape)
			stage.OnAfterNoteStateShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Object:
		if stage.OnAfterObjectDeleteCallback != nil {
			staged := any(staged).(*Object)
			stage.OnAfterObjectDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Role:
		if stage.OnAfterRoleDeleteCallback != nil {
			staged := any(staged).(*Role)
			stage.OnAfterRoleDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *State:
		if stage.OnAfterStateDeleteCallback != nil {
			staged := any(staged).(*State)
			stage.OnAfterStateDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *StateMachine:
		if stage.OnAfterStateMachineDeleteCallback != nil {
			staged := any(staged).(*StateMachine)
			stage.OnAfterStateMachineDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *StateShape:
		if stage.OnAfterStateShapeDeleteCallback != nil {
			staged := any(staged).(*StateShape)
			stage.OnAfterStateShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Transition:
		if stage.OnAfterTransitionDeleteCallback != nil {
			staged := any(staged).(*Transition)
			stage.OnAfterTransitionDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Transition_Shape:
		if stage.OnAfterTransition_ShapeDeleteCallback != nil {
			staged := any(staged).(*Transition_Shape)
			stage.OnAfterTransition_ShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	default:
		_ = front
	}
}

// AfterDeleteFromFront is a backward-compatible package-level forwarder.
func AfterDeleteFromFront[Type Gongstruct](stage *Stage, staged, front *Type) {
	stage.AfterDeleteFromFront(staged, front)
}
