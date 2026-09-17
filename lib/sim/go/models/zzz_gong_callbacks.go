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
func (command *Command) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCommandCreateCallback != nil {
		stage.OnAfterCommandCreateCallback.OnAfterCreate(stage, command)
	}
}

func (command *Command) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCommandUpdateCallback != nil {
		var frontCommand *Command
		if front != nil {
			frontCommand, _ = front.(*Command)
		}
		stage.OnAfterCommandUpdateCallback.OnAfterUpdate(stage, command, frontCommand)
	}
}

func (command *Command) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCommandDeleteCallback != nil {
		var frontCommand *Command
		if front != nil {
			frontCommand, _ = front.(*Command)
		}
		stage.OnAfterCommandDeleteCallback.OnAfterDelete(stage, command, frontCommand)
	}
}

func (dummyagent *DummyAgent) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDummyAgentCreateCallback != nil {
		stage.OnAfterDummyAgentCreateCallback.OnAfterCreate(stage, dummyagent)
	}
}

func (dummyagent *DummyAgent) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDummyAgentUpdateCallback != nil {
		var frontDummyAgent *DummyAgent
		if front != nil {
			frontDummyAgent, _ = front.(*DummyAgent)
		}
		stage.OnAfterDummyAgentUpdateCallback.OnAfterUpdate(stage, dummyagent, frontDummyAgent)
	}
}

func (dummyagent *DummyAgent) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDummyAgentDeleteCallback != nil {
		var frontDummyAgent *DummyAgent
		if front != nil {
			frontDummyAgent, _ = front.(*DummyAgent)
		}
		stage.OnAfterDummyAgentDeleteCallback.OnAfterDelete(stage, dummyagent, frontDummyAgent)
	}
}

func (engine *Engine) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterEngineCreateCallback != nil {
		stage.OnAfterEngineCreateCallback.OnAfterCreate(stage, engine)
	}
}

func (engine *Engine) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEngineUpdateCallback != nil {
		var frontEngine *Engine
		if front != nil {
			frontEngine, _ = front.(*Engine)
		}
		stage.OnAfterEngineUpdateCallback.OnAfterUpdate(stage, engine, frontEngine)
	}
}

func (engine *Engine) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEngineDeleteCallback != nil {
		var frontEngine *Engine
		if front != nil {
			frontEngine, _ = front.(*Engine)
		}
		stage.OnAfterEngineDeleteCallback.OnAfterDelete(stage, engine, frontEngine)
	}
}

func (event *Event) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterEventCreateCallback != nil {
		stage.OnAfterEventCreateCallback.OnAfterCreate(stage, event)
	}
}

func (event *Event) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEventUpdateCallback != nil {
		var frontEvent *Event
		if front != nil {
			frontEvent, _ = front.(*Event)
		}
		stage.OnAfterEventUpdateCallback.OnAfterUpdate(stage, event, frontEvent)
	}
}

func (event *Event) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEventDeleteCallback != nil {
		var frontEvent *Event
		if front != nil {
			frontEvent, _ = front.(*Event)
		}
		stage.OnAfterEventDeleteCallback.OnAfterDelete(stage, event, frontEvent)
	}
}

func (status *Status) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterStatusCreateCallback != nil {
		stage.OnAfterStatusCreateCallback.OnAfterCreate(stage, status)
	}
}

func (status *Status) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStatusUpdateCallback != nil {
		var frontStatus *Status
		if front != nil {
			frontStatus, _ = front.(*Status)
		}
		stage.OnAfterStatusUpdateCallback.OnAfterUpdate(stage, status, frontStatus)
	}
}

func (status *Status) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStatusDeleteCallback != nil {
		var frontStatus *Status
		if front != nil {
			frontStatus, _ = front.(*Status)
		}
		stage.OnAfterStatusDeleteCallback.OnAfterDelete(stage, status, frontStatus)
	}
}

func (updatestate *UpdateState) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterUpdateStateCreateCallback != nil {
		stage.OnAfterUpdateStateCreateCallback.OnAfterCreate(stage, updatestate)
	}
}

func (updatestate *UpdateState) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterUpdateStateUpdateCallback != nil {
		var frontUpdateState *UpdateState
		if front != nil {
			frontUpdateState, _ = front.(*UpdateState)
		}
		stage.OnAfterUpdateStateUpdateCallback.OnAfterUpdate(stage, updatestate, frontUpdateState)
	}
}

func (updatestate *UpdateState) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterUpdateStateDeleteCallback != nil {
		var frontUpdateState *UpdateState
		if front != nil {
			frontUpdateState, _ = front.(*UpdateState)
		}
		stage.OnAfterUpdateStateDeleteCallback.OnAfterDelete(stage, updatestate, frontUpdateState)
	}
}

