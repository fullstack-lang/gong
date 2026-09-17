// generated code - do not edit
package models

// AfterCreateFromFront is the Stage method called after a create from front.
func (stage *Stage) AfterCreateFromFront(instance GongstructIF) {
	if instance != nil {
		instance.GongAfterCreateFromFront(stage)
	}
}

type Gong__MouseEvent struct {
	ShiftKey bool
}

// OnAfterUpdateFromFront is the Stage method called after an update from front.
func (stage *Stage) OnAfterUpdateFromFront(old, new GongstructIF) {
	if old != nil {
		old.GongOnAfterUpdateFromFront(stage, new)
	}
}

// AfterDeleteFromFront is the Stage method called after a delete from front.
func (stage *Stage) AfterDeleteFromFront(staged, front GongstructIF) {
	if staged != nil {
		staged.GongAfterDeleteFromFront(stage, front)
	}
}

// insertion point
func (action *Action) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterActionCreateCallback != nil {
		stage.OnAfterActionCreateCallback.OnAfterCreate(stage, action)
	}
}

func (action *Action) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterActionUpdateCallback != nil {
		var frontAction *Action
		if front != nil {
			frontAction, _ = front.(*Action)
		}
		stage.OnAfterActionUpdateCallback.OnAfterUpdate(stage, action, frontAction)
	}
}

func (action *Action) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterActionDeleteCallback != nil {
		var frontAction *Action
		if front != nil {
			frontAction, _ = front.(*Action)
		}
		stage.OnAfterActionDeleteCallback.OnAfterDelete(stage, action, frontAction)
	}
}

func (activities *Activities) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterActivitiesCreateCallback != nil {
		stage.OnAfterActivitiesCreateCallback.OnAfterCreate(stage, activities)
	}
}

func (activities *Activities) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterActivitiesUpdateCallback != nil {
		var frontActivities *Activities
		if front != nil {
			frontActivities, _ = front.(*Activities)
		}
		stage.OnAfterActivitiesUpdateCallback.OnAfterUpdate(stage, activities, frontActivities)
	}
}

func (activities *Activities) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterActivitiesDeleteCallback != nil {
		var frontActivities *Activities
		if front != nil {
			frontActivities, _ = front.(*Activities)
		}
		stage.OnAfterActivitiesDeleteCallback.OnAfterDelete(stage, activities, frontActivities)
	}
}

func (diagram *Diagram) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDiagramCreateCallback != nil {
		stage.OnAfterDiagramCreateCallback.OnAfterCreate(stage, diagram)
	}
}

func (diagram *Diagram) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDiagramUpdateCallback != nil {
		var frontDiagram *Diagram
		if front != nil {
			frontDiagram, _ = front.(*Diagram)
		}
		stage.OnAfterDiagramUpdateCallback.OnAfterUpdate(stage, diagram, frontDiagram)
	}
}

func (diagram *Diagram) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDiagramDeleteCallback != nil {
		var frontDiagram *Diagram
		if front != nil {
			frontDiagram, _ = front.(*Diagram)
		}
		stage.OnAfterDiagramDeleteCallback.OnAfterDelete(stage, diagram, frontDiagram)
	}
}

func (guard *Guard) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterGuardCreateCallback != nil {
		stage.OnAfterGuardCreateCallback.OnAfterCreate(stage, guard)
	}
}

func (guard *Guard) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGuardUpdateCallback != nil {
		var frontGuard *Guard
		if front != nil {
			frontGuard, _ = front.(*Guard)
		}
		stage.OnAfterGuardUpdateCallback.OnAfterUpdate(stage, guard, frontGuard)
	}
}

func (guard *Guard) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGuardDeleteCallback != nil {
		var frontGuard *Guard
		if front != nil {
			frontGuard, _ = front.(*Guard)
		}
		stage.OnAfterGuardDeleteCallback.OnAfterDelete(stage, guard, frontGuard)
	}
}

func (kill *Kill) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterKillCreateCallback != nil {
		stage.OnAfterKillCreateCallback.OnAfterCreate(stage, kill)
	}
}

func (kill *Kill) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterKillUpdateCallback != nil {
		var frontKill *Kill
		if front != nil {
			frontKill, _ = front.(*Kill)
		}
		stage.OnAfterKillUpdateCallback.OnAfterUpdate(stage, kill, frontKill)
	}
}

func (kill *Kill) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterKillDeleteCallback != nil {
		var frontKill *Kill
		if front != nil {
			frontKill, _ = front.(*Kill)
		}
		stage.OnAfterKillDeleteCallback.OnAfterDelete(stage, kill, frontKill)
	}
}

func (library *Library) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterLibraryCreateCallback != nil {
		stage.OnAfterLibraryCreateCallback.OnAfterCreate(stage, library)
	}
}

func (library *Library) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLibraryUpdateCallback != nil {
		var frontLibrary *Library
		if front != nil {
			frontLibrary, _ = front.(*Library)
		}
		stage.OnAfterLibraryUpdateCallback.OnAfterUpdate(stage, library, frontLibrary)
	}
}

func (library *Library) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLibraryDeleteCallback != nil {
		var frontLibrary *Library
		if front != nil {
			frontLibrary, _ = front.(*Library)
		}
		stage.OnAfterLibraryDeleteCallback.OnAfterDelete(stage, library, frontLibrary)
	}
}

func (message *Message) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMessageCreateCallback != nil {
		stage.OnAfterMessageCreateCallback.OnAfterCreate(stage, message)
	}
}

func (message *Message) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMessageUpdateCallback != nil {
		var frontMessage *Message
		if front != nil {
			frontMessage, _ = front.(*Message)
		}
		stage.OnAfterMessageUpdateCallback.OnAfterUpdate(stage, message, frontMessage)
	}
}

func (message *Message) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMessageDeleteCallback != nil {
		var frontMessage *Message
		if front != nil {
			frontMessage, _ = front.(*Message)
		}
		stage.OnAfterMessageDeleteCallback.OnAfterDelete(stage, message, frontMessage)
	}
}

func (messagetype *MessageType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMessageTypeCreateCallback != nil {
		stage.OnAfterMessageTypeCreateCallback.OnAfterCreate(stage, messagetype)
	}
}

func (messagetype *MessageType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMessageTypeUpdateCallback != nil {
		var frontMessageType *MessageType
		if front != nil {
			frontMessageType, _ = front.(*MessageType)
		}
		stage.OnAfterMessageTypeUpdateCallback.OnAfterUpdate(stage, messagetype, frontMessageType)
	}
}

func (messagetype *MessageType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMessageTypeDeleteCallback != nil {
		var frontMessageType *MessageType
		if front != nil {
			frontMessageType, _ = front.(*MessageType)
		}
		stage.OnAfterMessageTypeDeleteCallback.OnAfterDelete(stage, messagetype, frontMessageType)
	}
}

func (note *Note) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterNoteCreateCallback != nil {
		stage.OnAfterNoteCreateCallback.OnAfterCreate(stage, note)
	}
}

func (note *Note) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNoteUpdateCallback != nil {
		var frontNote *Note
		if front != nil {
			frontNote, _ = front.(*Note)
		}
		stage.OnAfterNoteUpdateCallback.OnAfterUpdate(stage, note, frontNote)
	}
}

func (note *Note) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNoteDeleteCallback != nil {
		var frontNote *Note
		if front != nil {
			frontNote, _ = front.(*Note)
		}
		stage.OnAfterNoteDeleteCallback.OnAfterDelete(stage, note, frontNote)
	}
}

func (noteshape *NoteShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterNoteShapeCreateCallback != nil {
		stage.OnAfterNoteShapeCreateCallback.OnAfterCreate(stage, noteshape)
	}
}

func (noteshape *NoteShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNoteShapeUpdateCallback != nil {
		var frontNoteShape *NoteShape
		if front != nil {
			frontNoteShape, _ = front.(*NoteShape)
		}
		stage.OnAfterNoteShapeUpdateCallback.OnAfterUpdate(stage, noteshape, frontNoteShape)
	}
}

func (noteshape *NoteShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNoteShapeDeleteCallback != nil {
		var frontNoteShape *NoteShape
		if front != nil {
			frontNoteShape, _ = front.(*NoteShape)
		}
		stage.OnAfterNoteShapeDeleteCallback.OnAfterDelete(stage, noteshape, frontNoteShape)
	}
}

func (notestateshape *NoteStateShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterNoteStateShapeCreateCallback != nil {
		stage.OnAfterNoteStateShapeCreateCallback.OnAfterCreate(stage, notestateshape)
	}
}

func (notestateshape *NoteStateShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNoteStateShapeUpdateCallback != nil {
		var frontNoteStateShape *NoteStateShape
		if front != nil {
			frontNoteStateShape, _ = front.(*NoteStateShape)
		}
		stage.OnAfterNoteStateShapeUpdateCallback.OnAfterUpdate(stage, notestateshape, frontNoteStateShape)
	}
}

func (notestateshape *NoteStateShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNoteStateShapeDeleteCallback != nil {
		var frontNoteStateShape *NoteStateShape
		if front != nil {
			frontNoteStateShape, _ = front.(*NoteStateShape)
		}
		stage.OnAfterNoteStateShapeDeleteCallback.OnAfterDelete(stage, notestateshape, frontNoteStateShape)
	}
}

func (object *Object) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterObjectCreateCallback != nil {
		stage.OnAfterObjectCreateCallback.OnAfterCreate(stage, object)
	}
}

func (object *Object) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterObjectUpdateCallback != nil {
		var frontObject *Object
		if front != nil {
			frontObject, _ = front.(*Object)
		}
		stage.OnAfterObjectUpdateCallback.OnAfterUpdate(stage, object, frontObject)
	}
}

func (object *Object) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterObjectDeleteCallback != nil {
		var frontObject *Object
		if front != nil {
			frontObject, _ = front.(*Object)
		}
		stage.OnAfterObjectDeleteCallback.OnAfterDelete(stage, object, frontObject)
	}
}

func (role *Role) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRoleCreateCallback != nil {
		stage.OnAfterRoleCreateCallback.OnAfterCreate(stage, role)
	}
}

func (role *Role) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRoleUpdateCallback != nil {
		var frontRole *Role
		if front != nil {
			frontRole, _ = front.(*Role)
		}
		stage.OnAfterRoleUpdateCallback.OnAfterUpdate(stage, role, frontRole)
	}
}

func (role *Role) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRoleDeleteCallback != nil {
		var frontRole *Role
		if front != nil {
			frontRole, _ = front.(*Role)
		}
		stage.OnAfterRoleDeleteCallback.OnAfterDelete(stage, role, frontRole)
	}
}

func (state *State) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterStateCreateCallback != nil {
		stage.OnAfterStateCreateCallback.OnAfterCreate(stage, state)
	}
}

func (state *State) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStateUpdateCallback != nil {
		var frontState *State
		if front != nil {
			frontState, _ = front.(*State)
		}
		stage.OnAfterStateUpdateCallback.OnAfterUpdate(stage, state, frontState)
	}
}

func (state *State) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStateDeleteCallback != nil {
		var frontState *State
		if front != nil {
			frontState, _ = front.(*State)
		}
		stage.OnAfterStateDeleteCallback.OnAfterDelete(stage, state, frontState)
	}
}

func (statemachine *StateMachine) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterStateMachineCreateCallback != nil {
		stage.OnAfterStateMachineCreateCallback.OnAfterCreate(stage, statemachine)
	}
}

func (statemachine *StateMachine) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStateMachineUpdateCallback != nil {
		var frontStateMachine *StateMachine
		if front != nil {
			frontStateMachine, _ = front.(*StateMachine)
		}
		stage.OnAfterStateMachineUpdateCallback.OnAfterUpdate(stage, statemachine, frontStateMachine)
	}
}

func (statemachine *StateMachine) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStateMachineDeleteCallback != nil {
		var frontStateMachine *StateMachine
		if front != nil {
			frontStateMachine, _ = front.(*StateMachine)
		}
		stage.OnAfterStateMachineDeleteCallback.OnAfterDelete(stage, statemachine, frontStateMachine)
	}
}

func (stateshape *StateShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterStateShapeCreateCallback != nil {
		stage.OnAfterStateShapeCreateCallback.OnAfterCreate(stage, stateshape)
	}
}

func (stateshape *StateShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStateShapeUpdateCallback != nil {
		var frontStateShape *StateShape
		if front != nil {
			frontStateShape, _ = front.(*StateShape)
		}
		stage.OnAfterStateShapeUpdateCallback.OnAfterUpdate(stage, stateshape, frontStateShape)
	}
}

func (stateshape *StateShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStateShapeDeleteCallback != nil {
		var frontStateShape *StateShape
		if front != nil {
			frontStateShape, _ = front.(*StateShape)
		}
		stage.OnAfterStateShapeDeleteCallback.OnAfterDelete(stage, stateshape, frontStateShape)
	}
}

func (transition *Transition) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTransitionCreateCallback != nil {
		stage.OnAfterTransitionCreateCallback.OnAfterCreate(stage, transition)
	}
}

func (transition *Transition) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTransitionUpdateCallback != nil {
		var frontTransition *Transition
		if front != nil {
			frontTransition, _ = front.(*Transition)
		}
		stage.OnAfterTransitionUpdateCallback.OnAfterUpdate(stage, transition, frontTransition)
	}
}

func (transition *Transition) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTransitionDeleteCallback != nil {
		var frontTransition *Transition
		if front != nil {
			frontTransition, _ = front.(*Transition)
		}
		stage.OnAfterTransitionDeleteCallback.OnAfterDelete(stage, transition, frontTransition)
	}
}

func (transition_shape *Transition_Shape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTransition_ShapeCreateCallback != nil {
		stage.OnAfterTransition_ShapeCreateCallback.OnAfterCreate(stage, transition_shape)
	}
}

func (transition_shape *Transition_Shape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTransition_ShapeUpdateCallback != nil {
		var frontTransition_Shape *Transition_Shape
		if front != nil {
			frontTransition_Shape, _ = front.(*Transition_Shape)
		}
		stage.OnAfterTransition_ShapeUpdateCallback.OnAfterUpdate(stage, transition_shape, frontTransition_Shape)
	}
}

func (transition_shape *Transition_Shape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTransition_ShapeDeleteCallback != nil {
		var frontTransition_Shape *Transition_Shape
		if front != nil {
			frontTransition_Shape, _ = front.(*Transition_Shape)
		}
		stage.OnAfterTransition_ShapeDeleteCallback.OnAfterDelete(stage, transition_shape, frontTransition_Shape)
	}
}

