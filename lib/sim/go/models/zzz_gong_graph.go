// generated code - do not edit
package models

import "fmt"

// IsStaged is the Stage method checking if a gongstruct instance is staged.
func (stage *Stage) IsStaged(instance GongstructIF) (ok bool) {
	if instance != nil {
		return instance.GongIsStaged(stage)
	}
	return false
}

// insertion point for stage per struct
func (command *Command) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Commands[command]

	return
}

func (stage *Stage) IsStagedCommand(command *Command) (ok bool) {

	return command.GongIsStaged(stage)
}

func (dummyagent *DummyAgent) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.DummyAgents[dummyagent]

	return
}

func (stage *Stage) IsStagedDummyAgent(dummyagent *DummyAgent) (ok bool) {

	return dummyagent.GongIsStaged(stage)
}

func (engine *Engine) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Engines[engine]

	return
}

func (stage *Stage) IsStagedEngine(engine *Engine) (ok bool) {

	return engine.GongIsStaged(stage)
}

func (event *Event) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Events[event]

	return
}

func (stage *Stage) IsStagedEvent(event *Event) (ok bool) {

	return event.GongIsStaged(stage)
}

func (status *Status) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Statuss[status]

	return
}

func (stage *Stage) IsStagedStatus(status *Status) (ok bool) {

	return status.GongIsStaged(stage)
}

func (updatestate *UpdateState) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.UpdateStates[updatestate]

	return
}

func (stage *Stage) IsStagedUpdateState(updatestate *UpdateState) (ok bool) {

	return updatestate.GongIsStaged(stage)
}

// StageBranch is the Stage method that stages instance and applies StageBranch recursively.
func (stage *Stage) StageBranch(instance GongstructIF) {
	if instance != nil {
		instance.GongStageBranch(stage)
	}
}

// StageBranch is a backward-compatible package-level forwarder.
func StageBranch(stage *Stage, instance GongstructIF) {
	stage.StageBranch(instance)
}

// insertion point for stage branch per struct
func (command *Command) GongStageBranch(stage *Stage) {
	stage.StageBranchCommand(command)
}

func (stage *Stage) StageBranchCommand(command *Command) {

	// check if instance is already staged
	if stage.IsStaged(command) {
		return
	}

	command.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if command.Engine != nil {
		stage.StageBranch(command.Engine)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (dummyagent *DummyAgent) GongStageBranch(stage *Stage) {
	stage.StageBranchDummyAgent(dummyagent)
}

func (stage *Stage) StageBranchDummyAgent(dummyagent *DummyAgent) {

	// check if instance is already staged
	if stage.IsStaged(dummyagent) {
		return
	}

	dummyagent.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (engine *Engine) GongStageBranch(stage *Stage) {
	stage.StageBranchEngine(engine)
}

func (stage *Stage) StageBranchEngine(engine *Engine) {

	// check if instance is already staged
	if stage.IsStaged(engine) {
		return
	}

	engine.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (event *Event) GongStageBranch(stage *Stage) {
	stage.StageBranchEvent(event)
}

func (stage *Stage) StageBranchEvent(event *Event) {

	// check if instance is already staged
	if stage.IsStaged(event) {
		return
	}

	event.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (status *Status) GongStageBranch(stage *Stage) {
	stage.StageBranchStatus(status)
}

func (stage *Stage) StageBranchStatus(status *Status) {

	// check if instance is already staged
	if stage.IsStaged(status) {
		return
	}

	status.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (updatestate *UpdateState) GongStageBranch(stage *Stage) {
	stage.StageBranchUpdateState(updatestate)
}

func (stage *Stage) StageBranchUpdateState(updatestate *UpdateState) {

	// check if instance is already staged
	if stage.IsStaged(updatestate) {
		return
	}

	updatestate.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

// GongCopyBranch stages instance and apply GongCopyBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func GongCopyBranch[Type Gongstruct](from *Type) (to *Type) {

	mapOrigCopy := make(map[any]any)
	_ = mapOrigCopy

	switch fromT := any(from).(type) {
	// insertion point for stage branch
	case *Command:
		toT := GongCopyBranchCommand(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *DummyAgent:
		toT := GongCopyBranchDummyAgent(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Engine:
		toT := GongCopyBranchEngine(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Event:
		toT := GongCopyBranchEvent(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Status:
		toT := GongCopyBranchStatus(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *UpdateState:
		toT := GongCopyBranchUpdateState(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
func GongCopyBranchCommand(mapOrigCopy map[any]any, commandFrom *Command) (commandTo *Command) {

	// commandFrom has already been copied
	if _commandTo, ok := mapOrigCopy[commandFrom]; ok {
		commandTo = _commandTo.(*Command)
		return
	}

	commandTo = new(Command)
	mapOrigCopy[commandFrom] = commandTo
	commandFrom.GongCopyBasicFields(commandTo)

	//insertion point for the staging of instances referenced by pointers
	if commandFrom.Engine != nil {
		commandTo.Engine = GongCopyBranchEngine(mapOrigCopy, commandFrom.Engine)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchDummyAgent(mapOrigCopy map[any]any, dummyagentFrom *DummyAgent) (dummyagentTo *DummyAgent) {

	// dummyagentFrom has already been copied
	if _dummyagentTo, ok := mapOrigCopy[dummyagentFrom]; ok {
		dummyagentTo = _dummyagentTo.(*DummyAgent)
		return
	}

	dummyagentTo = new(DummyAgent)
	mapOrigCopy[dummyagentFrom] = dummyagentTo
	dummyagentFrom.GongCopyBasicFields(dummyagentTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchEngine(mapOrigCopy map[any]any, engineFrom *Engine) (engineTo *Engine) {

	// engineFrom has already been copied
	if _engineTo, ok := mapOrigCopy[engineFrom]; ok {
		engineTo = _engineTo.(*Engine)
		return
	}

	engineTo = new(Engine)
	mapOrigCopy[engineFrom] = engineTo
	engineFrom.GongCopyBasicFields(engineTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchEvent(mapOrigCopy map[any]any, eventFrom *Event) (eventTo *Event) {

	// eventFrom has already been copied
	if _eventTo, ok := mapOrigCopy[eventFrom]; ok {
		eventTo = _eventTo.(*Event)
		return
	}

	eventTo = new(Event)
	mapOrigCopy[eventFrom] = eventTo
	eventFrom.GongCopyBasicFields(eventTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchStatus(mapOrigCopy map[any]any, statusFrom *Status) (statusTo *Status) {

	// statusFrom has already been copied
	if _statusTo, ok := mapOrigCopy[statusFrom]; ok {
		statusTo = _statusTo.(*Status)
		return
	}

	statusTo = new(Status)
	mapOrigCopy[statusFrom] = statusTo
	statusFrom.GongCopyBasicFields(statusTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchUpdateState(mapOrigCopy map[any]any, updatestateFrom *UpdateState) (updatestateTo *UpdateState) {

	// updatestateFrom has already been copied
	if _updatestateTo, ok := mapOrigCopy[updatestateFrom]; ok {
		updatestateTo = _updatestateTo.(*UpdateState)
		return
	}

	updatestateTo = new(UpdateState)
	mapOrigCopy[updatestateFrom] = updatestateTo
	updatestateFrom.GongCopyBasicFields(updatestateTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

// UnstageBranch stages instance and apply UnstageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the insance
//
// the algorithm stops along the course of graph if a vertex is already staged
// UnstageBranch is the Stage method that unstages instance and applies UnstageBranch recursively.
func (stage *Stage) UnstageBranch(instance GongstructIF) {
	if instance != nil {
		instance.GongUnstageBranch(stage)
	}
}

// UnstageBranch is a backward-compatible package-level forwarder.
func UnstageBranch(stage *Stage, instance GongstructIF) {
	stage.UnstageBranch(instance)
}

// insertion point for unstage branch per struct
func (command *Command) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchCommand(command)
}

func (stage *Stage) UnstageBranchCommand(command *Command) {

	// check if instance is already staged
	if !stage.IsStaged(command) {
		return
	}

	command.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if command.Engine != nil {
		stage.UnstageBranch(command.Engine)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (dummyagent *DummyAgent) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchDummyAgent(dummyagent)
}

func (stage *Stage) UnstageBranchDummyAgent(dummyagent *DummyAgent) {

	// check if instance is already staged
	if !stage.IsStaged(dummyagent) {
		return
	}

	dummyagent.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (engine *Engine) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchEngine(engine)
}

func (stage *Stage) UnstageBranchEngine(engine *Engine) {

	// check if instance is already staged
	if !stage.IsStaged(engine) {
		return
	}

	engine.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (event *Event) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchEvent(event)
}

func (stage *Stage) UnstageBranchEvent(event *Event) {

	// check if instance is already staged
	if !stage.IsStaged(event) {
		return
	}

	event.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (status *Status) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchStatus(status)
}

func (stage *Stage) UnstageBranchStatus(status *Status) {

	// check if instance is already staged
	if !stage.IsStaged(status) {
		return
	}

	status.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (updatestate *UpdateState) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchUpdateState(updatestate)
}

func (stage *Stage) UnstageBranchUpdateState(updatestate *UpdateState) {

	// check if instance is already staged
	if !stage.IsStaged(updatestate) {
		return
	}

	updatestate.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

// insertion point for pointer reconstruction from references
func (reference *Command) GongReconstructPointersFromReferences(stage *Stage, instance *Command) {
	// insertion point for pointers field
	if instance.Engine != nil {
		reference.Engine = stage.Engines_reference[instance.Engine]
	}
	// insertion point for slice of pointers field
}

func (reference *DummyAgent) GongReconstructPointersFromReferences(stage *Stage, instance *DummyAgent) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Engine) GongReconstructPointersFromReferences(stage *Stage, instance *Engine) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Event) GongReconstructPointersFromReferences(stage *Stage, instance *Event) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Status) GongReconstructPointersFromReferences(stage *Stage, instance *Status) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *UpdateState) GongReconstructPointersFromReferences(stage *Stage, instance *UpdateState) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

// insertion point for pointer reconstruction from instances
func (reference *Command) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.Engine; _reference != nil {
		reference.Engine = nil
		if _instance, ok := stage.Engines_instance[_reference]; ok {
			reference.Engine = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *DummyAgent) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Engine) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Event) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Status) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *UpdateState) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

// insertion point for diff per struct
// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (command *Command) GongDiff(stage *Stage, commandOther *Command) (diffs []string) {
	// insertion point for field diffs
	if command.Name != commandOther.Name {
		diffs = append(diffs, command.GongMarshallField(stage, "Name"))
	}
	if command.Command != commandOther.Command {
		diffs = append(diffs, command.GongMarshallField(stage, "Command"))
	}
	if command.CommandDate != commandOther.CommandDate {
		diffs = append(diffs, command.GongMarshallField(stage, "CommandDate"))
	}
	if (command.Engine == nil) != (commandOther.Engine == nil) {
		diffs = append(diffs, command.GongMarshallField(stage, "Engine"))
	} else if command.Engine != nil && commandOther.Engine != nil {
		if command.Engine != commandOther.Engine {
			diffs = append(diffs, command.GongMarshallField(stage, "Engine"))
		}
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (dummyagent *DummyAgent) GongDiff(stage *Stage, dummyagentOther *DummyAgent) (diffs []string) {
	// insertion point for field diffs
	if dummyagent.TechName != dummyagentOther.TechName {
		diffs = append(diffs, dummyagent.GongMarshallField(stage, "TechName"))
	}
	if dummyagent.Name != dummyagentOther.Name {
		diffs = append(diffs, dummyagent.GongMarshallField(stage, "Name"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (engine *Engine) GongDiff(stage *Stage, engineOther *Engine) (diffs []string) {
	// insertion point for field diffs
	if engine.Name != engineOther.Name {
		diffs = append(diffs, engine.GongMarshallField(stage, "Name"))
	}
	if engine.EndTime != engineOther.EndTime {
		diffs = append(diffs, engine.GongMarshallField(stage, "EndTime"))
	}
	if engine.CurrentTime != engineOther.CurrentTime {
		diffs = append(diffs, engine.GongMarshallField(stage, "CurrentTime"))
	}
	if engine.DisplayFormat != engineOther.DisplayFormat {
		diffs = append(diffs, engine.GongMarshallField(stage, "DisplayFormat"))
	}
	if engine.SecondsSinceStart != engineOther.SecondsSinceStart {
		diffs = append(diffs, engine.GongMarshallField(stage, "SecondsSinceStart"))
	}
	if engine.Fired != engineOther.Fired {
		diffs = append(diffs, engine.GongMarshallField(stage, "Fired"))
	}
	if engine.ControlMode != engineOther.ControlMode {
		diffs = append(diffs, engine.GongMarshallField(stage, "ControlMode"))
	}
	if engine.State != engineOther.State {
		diffs = append(diffs, engine.GongMarshallField(stage, "State"))
	}
	if engine.Speed != engineOther.Speed {
		diffs = append(diffs, engine.GongMarshallField(stage, "Speed"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (event *Event) GongDiff(stage *Stage, eventOther *Event) (diffs []string) {
	// insertion point for field diffs
	if event.Name != eventOther.Name {
		diffs = append(diffs, event.GongMarshallField(stage, "Name"))
	}
	if event.Duration != eventOther.Duration {
		diffs = append(diffs, event.GongMarshallField(stage, "Duration"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (status *Status) GongDiff(stage *Stage, statusOther *Status) (diffs []string) {
	// insertion point for field diffs
	if status.Name != statusOther.Name {
		diffs = append(diffs, status.GongMarshallField(stage, "Name"))
	}
	if status.CurrentCommand != statusOther.CurrentCommand {
		diffs = append(diffs, status.GongMarshallField(stage, "CurrentCommand"))
	}
	if status.CompletionDate != statusOther.CompletionDate {
		diffs = append(diffs, status.GongMarshallField(stage, "CompletionDate"))
	}
	if status.CurrentSpeedCommand != statusOther.CurrentSpeedCommand {
		diffs = append(diffs, status.GongMarshallField(stage, "CurrentSpeedCommand"))
	}
	if status.SpeedCommandCompletionDate != statusOther.SpeedCommandCompletionDate {
		diffs = append(diffs, status.GongMarshallField(stage, "SpeedCommandCompletionDate"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (updatestate *UpdateState) GongDiff(stage *Stage, updatestateOther *UpdateState) (diffs []string) {
	// insertion point for field diffs
	if updatestate.Name != updatestateOther.Name {
		diffs = append(diffs, updatestate.GongMarshallField(stage, "Name"))
	}
	if updatestate.Duration != updatestateOther.Duration {
		diffs = append(diffs, updatestate.GongMarshallField(stage, "Duration"))
	}
	if updatestate.Period != updatestateOther.Period {
		diffs = append(diffs, updatestate.GongMarshallField(stage, "Period"))
	}

	return
}

// Diff is the Stage method that returns the sequence of operations to transform oldSlice into newSlice.
func (stage *Stage) Diff(
	a GongstructIF,
	fieldName string,
	lenOld, lenNew int,
	equal func(i, j int) bool,
	getNewIdentifier func(j int) string,
) (ops string) {
	m, n := lenOld, lenNew

	// 1. Build the LCS (Longest Common Subsequence) Matrix
	// This helps us find the "anchor" elements that shouldn't move.
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if equal(i, j) {
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
		if equal(i-1, j-1) {
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

	// Track kept indices in old slice
	keptOldIndices := make([]int, 0, len(keptIndices))
	for k := 0; k < m; k++ {
		if keptIndices[k] {
			keptOldIndices = append(keptOldIndices, k)
		}
	}

	lcsIdx := 0
	// Iterate through the NEW slice. If it matches the current LCS head, we keep it.
	// If it doesn't match, it must be inserted here.
	for k := 0; k < n; k++ {
		if lcsIdx < len(keptOldIndices) && equal(keptOldIndices[lcsIdx], k) {
			lcsIdx++
		} else {
			ops += fmt.Sprintf("\n\t%s.%s = slices.Insert( %s.%s, %d, %s)", a.GongGetIdentifier(stage), fieldName, a.GongGetIdentifier(stage), fieldName, k, getNewIdentifier(k))
		}
	}

	return ops
}
