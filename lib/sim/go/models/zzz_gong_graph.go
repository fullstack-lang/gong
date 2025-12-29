// generated code - do not edit
package models

func IsStagedPointerToGongstruct[Type PointerToGongstruct](stage *Stage, instance Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
	case *Command:
		ok = stage.IsStagedCommand(target)

	case *DummyAgent:
		ok = stage.IsStagedDummyAgent(target)

	case *Engine:
		ok = stage.IsStagedEngine(target)

	case *Event:
		ok = stage.IsStagedEvent(target)

	case *Status:
		ok = stage.IsStagedStatus(target)

	case *UpdateState:
		ok = stage.IsStagedUpdateState(target)

	default:
		_ = target
	}
	return
}

func IsStaged[Type Gongstruct](stage *Stage, instance *Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
	case *Command:
		ok = stage.IsStagedCommand(target)

	case *DummyAgent:
		ok = stage.IsStagedDummyAgent(target)

	case *Engine:
		ok = stage.IsStagedEngine(target)

	case *Event:
		ok = stage.IsStagedEvent(target)

	case *Status:
		ok = stage.IsStagedStatus(target)

	case *UpdateState:
		ok = stage.IsStagedUpdateState(target)

	default:
		_ = target
	}
	return
}

// insertion point for stage per struct
func (stage *Stage) IsStagedCommand(command *Command) (ok bool) {

	_, ok = stage.Commands[command]

	return
}

func (stage *Stage) IsStagedDummyAgent(dummyagent *DummyAgent) (ok bool) {

	_, ok = stage.DummyAgents[dummyagent]

	return
}

func (stage *Stage) IsStagedEngine(engine *Engine) (ok bool) {

	_, ok = stage.Engines[engine]

	return
}

func (stage *Stage) IsStagedEvent(event *Event) (ok bool) {

	_, ok = stage.Events[event]

	return
}

func (stage *Stage) IsStagedStatus(status *Status) (ok bool) {

	_, ok = stage.Statuss[status]

	return
}

func (stage *Stage) IsStagedUpdateState(updatestate *UpdateState) (ok bool) {

	_, ok = stage.UpdateStates[updatestate]

	return
}

// StageBranch stages instance and apply StageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func StageBranch[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for stage branch
	case *Command:
		stage.StageBranchCommand(target)

	case *DummyAgent:
		stage.StageBranchDummyAgent(target)

	case *Engine:
		stage.StageBranchEngine(target)

	case *Event:
		stage.StageBranchEvent(target)

	case *Status:
		stage.StageBranchStatus(target)

	case *UpdateState:
		stage.StageBranchUpdateState(target)

	default:
		_ = target
	}
}

// insertion point for stage branch per struct
func (stage *Stage) StageBranchCommand(command *Command) {

	// check if instance is already staged
	if IsStaged(stage, command) {
		return
	}

	command.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if command.Engine != nil {
		StageBranch(stage, command.Engine)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchDummyAgent(dummyagent *DummyAgent) {

	// check if instance is already staged
	if IsStaged(stage, dummyagent) {
		return
	}

	dummyagent.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchEngine(engine *Engine) {

	// check if instance is already staged
	if IsStaged(stage, engine) {
		return
	}

	engine.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchEvent(event *Event) {

	// check if instance is already staged
	if IsStaged(stage, event) {
		return
	}

	event.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchStatus(status *Status) {

	// check if instance is already staged
	if IsStaged(stage, status) {
		return
	}

	status.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchUpdateState(updatestate *UpdateState) {

	// check if instance is already staged
	if IsStaged(stage, updatestate) {
		return
	}

	updatestate.Stage(stage)

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
	case *Command:
		toT := CopyBranchCommand(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *DummyAgent:
		toT := CopyBranchDummyAgent(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Engine:
		toT := CopyBranchEngine(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Event:
		toT := CopyBranchEvent(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Status:
		toT := CopyBranchStatus(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *UpdateState:
		toT := CopyBranchUpdateState(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
func CopyBranchCommand(mapOrigCopy map[any]any, commandFrom *Command) (commandTo *Command) {

	// commandFrom has already been copied
	if _commandTo, ok := mapOrigCopy[commandFrom]; ok {
		commandTo = _commandTo.(*Command)
		return
	}

	commandTo = new(Command)
	mapOrigCopy[commandFrom] = commandTo
	commandFrom.CopyBasicFields(commandTo)

	//insertion point for the staging of instances referenced by pointers
	if commandFrom.Engine != nil {
		commandTo.Engine = CopyBranchEngine(mapOrigCopy, commandFrom.Engine)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchDummyAgent(mapOrigCopy map[any]any, dummyagentFrom *DummyAgent) (dummyagentTo *DummyAgent) {

	// dummyagentFrom has already been copied
	if _dummyagentTo, ok := mapOrigCopy[dummyagentFrom]; ok {
		dummyagentTo = _dummyagentTo.(*DummyAgent)
		return
	}

	dummyagentTo = new(DummyAgent)
	mapOrigCopy[dummyagentFrom] = dummyagentTo
	dummyagentFrom.CopyBasicFields(dummyagentTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchEngine(mapOrigCopy map[any]any, engineFrom *Engine) (engineTo *Engine) {

	// engineFrom has already been copied
	if _engineTo, ok := mapOrigCopy[engineFrom]; ok {
		engineTo = _engineTo.(*Engine)
		return
	}

	engineTo = new(Engine)
	mapOrigCopy[engineFrom] = engineTo
	engineFrom.CopyBasicFields(engineTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchEvent(mapOrigCopy map[any]any, eventFrom *Event) (eventTo *Event) {

	// eventFrom has already been copied
	if _eventTo, ok := mapOrigCopy[eventFrom]; ok {
		eventTo = _eventTo.(*Event)
		return
	}

	eventTo = new(Event)
	mapOrigCopy[eventFrom] = eventTo
	eventFrom.CopyBasicFields(eventTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchStatus(mapOrigCopy map[any]any, statusFrom *Status) (statusTo *Status) {

	// statusFrom has already been copied
	if _statusTo, ok := mapOrigCopy[statusFrom]; ok {
		statusTo = _statusTo.(*Status)
		return
	}

	statusTo = new(Status)
	mapOrigCopy[statusFrom] = statusTo
	statusFrom.CopyBasicFields(statusTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchUpdateState(mapOrigCopy map[any]any, updatestateFrom *UpdateState) (updatestateTo *UpdateState) {

	// updatestateFrom has already been copied
	if _updatestateTo, ok := mapOrigCopy[updatestateFrom]; ok {
		updatestateTo = _updatestateTo.(*UpdateState)
		return
	}

	updatestateTo = new(UpdateState)
	mapOrigCopy[updatestateFrom] = updatestateTo
	updatestateFrom.CopyBasicFields(updatestateTo)

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
	case *Command:
		stage.UnstageBranchCommand(target)

	case *DummyAgent:
		stage.UnstageBranchDummyAgent(target)

	case *Engine:
		stage.UnstageBranchEngine(target)

	case *Event:
		stage.UnstageBranchEvent(target)

	case *Status:
		stage.UnstageBranchStatus(target)

	case *UpdateState:
		stage.UnstageBranchUpdateState(target)

	default:
		_ = target
	}
}

// insertion point for unstage branch per struct
func (stage *Stage) UnstageBranchCommand(command *Command) {

	// check if instance is already staged
	if !IsStaged(stage, command) {
		return
	}

	command.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if command.Engine != nil {
		UnstageBranch(stage, command.Engine)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchDummyAgent(dummyagent *DummyAgent) {

	// check if instance is already staged
	if !IsStaged(stage, dummyagent) {
		return
	}

	dummyagent.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchEngine(engine *Engine) {

	// check if instance is already staged
	if !IsStaged(stage, engine) {
		return
	}

	engine.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchEvent(event *Event) {

	// check if instance is already staged
	if !IsStaged(stage, event) {
		return
	}

	event.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchStatus(status *Status) {

	// check if instance is already staged
	if !IsStaged(stage, status) {
		return
	}

	status.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchUpdateState(updatestate *UpdateState) {

	// check if instance is already staged
	if !IsStaged(stage, updatestate) {
		return
	}

	updatestate.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}


// insertion point for diff per struct
// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (command *Command) GongDiff(commandOther *Command) (diffs []string) {
	// insertion point for field diffs
	if command.Name != commandOther.Name {
		diffs = append(diffs, "Name")
	}
	if command.Command != commandOther.Command {
		diffs = append(diffs, "Command")
	}
	if command.CommandDate != commandOther.CommandDate {
		diffs = append(diffs, "CommandDate")
	}
	if (command.Engine == nil) != (commandOther.Engine == nil) {
		diffs = append(diffs, "Engine")
	} else if command.Engine != nil && commandOther.Engine != nil {
		if command.Engine != commandOther.Engine {
			diffs = append(diffs, "Engine")
		}
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (dummyagent *DummyAgent) GongDiff(dummyagentOther *DummyAgent) (diffs []string) {
	// insertion point for field diffs
	if dummyagent.TechName != dummyagentOther.TechName {
		diffs = append(diffs, "TechName")
	}
	if dummyagent.Name != dummyagentOther.Name {
		diffs = append(diffs, "Name")
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (engine *Engine) GongDiff(engineOther *Engine) (diffs []string) {
	// insertion point for field diffs
	if engine.Name != engineOther.Name {
		diffs = append(diffs, "Name")
	}
	if engine.EndTime != engineOther.EndTime {
		diffs = append(diffs, "EndTime")
	}
	if engine.CurrentTime != engineOther.CurrentTime {
		diffs = append(diffs, "CurrentTime")
	}
	if engine.DisplayFormat != engineOther.DisplayFormat {
		diffs = append(diffs, "DisplayFormat")
	}
	if engine.SecondsSinceStart != engineOther.SecondsSinceStart {
		diffs = append(diffs, "SecondsSinceStart")
	}
	if engine.Fired != engineOther.Fired {
		diffs = append(diffs, "Fired")
	}
	if engine.ControlMode != engineOther.ControlMode {
		diffs = append(diffs, "ControlMode")
	}
	if engine.State != engineOther.State {
		diffs = append(diffs, "State")
	}
	if engine.Speed != engineOther.Speed {
		diffs = append(diffs, "Speed")
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (event *Event) GongDiff(eventOther *Event) (diffs []string) {
	// insertion point for field diffs
	if event.Name != eventOther.Name {
		diffs = append(diffs, "Name")
	}
	if event.Duration != eventOther.Duration {
		diffs = append(diffs, "Duration")
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (status *Status) GongDiff(statusOther *Status) (diffs []string) {
	// insertion point for field diffs
	if status.Name != statusOther.Name {
		diffs = append(diffs, "Name")
	}
	if status.CurrentCommand != statusOther.CurrentCommand {
		diffs = append(diffs, "CurrentCommand")
	}
	if status.CompletionDate != statusOther.CompletionDate {
		diffs = append(diffs, "CompletionDate")
	}
	if status.CurrentSpeedCommand != statusOther.CurrentSpeedCommand {
		diffs = append(diffs, "CurrentSpeedCommand")
	}
	if status.SpeedCommandCompletionDate != statusOther.SpeedCommandCompletionDate {
		diffs = append(diffs, "SpeedCommandCompletionDate")
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (updatestate *UpdateState) GongDiff(updatestateOther *UpdateState) (diffs []string) {
	// insertion point for field diffs
	if updatestate.Name != updatestateOther.Name {
		diffs = append(diffs, "Name")
	}
	if updatestate.Duration != updatestateOther.Duration {
		diffs = append(diffs, "Duration")
	}
	if updatestate.Period != updatestateOther.Period {
		diffs = append(diffs, "Period")
	}

	return
}
