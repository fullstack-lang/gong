// generated code - do not edit
package models

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

var __GongSliceTemplate_time__dummyDeclaration time.Duration
var _ = __GongSliceTemplate_time__dummyDeclaration

// ComputeReverseMaps computes the reverse map, for all intances, for all slice to pointers field
// Its complexity is in O(n)O(p) where p is the number of pointers
func (stage *Stage) ComputeReverseMaps() {
	// insertion point per named struct
	// Compute reverse map for named struct Command
	// insertion point per field

	// Compute reverse map for named struct DummyAgent
	// insertion point per field

	// Compute reverse map for named struct Engine
	// insertion point per field

	// Compute reverse map for named struct Event
	// insertion point per field

	// Compute reverse map for named struct Status
	// insertion point per field

	// Compute reverse map for named struct UpdateState
	// insertion point per field

}

func (stage *Stage) GetInstances() (res []GongstructIF) {

	// insertion point per named struct
	for instance := range stage.Commands {
		res = append(res, instance)
	}

	for instance := range stage.DummyAgents {
		res = append(res, instance)
	}

	for instance := range stage.Engines {
		res = append(res, instance)
	}

	for instance := range stage.Events {
		res = append(res, instance)
	}

	for instance := range stage.Statuss {
		res = append(res, instance)
	}

	for instance := range stage.UpdateStates {
		res = append(res, instance)
	}

	return
}

// insertion point per named struct
func (command *Command) GongCopy() GongstructIF {
	newInstance := *command
	return &newInstance
}

func (dummyagent *DummyAgent) GongCopy() GongstructIF {
	newInstance := *dummyagent
	return &newInstance
}

func (engine *Engine) GongCopy() GongstructIF {
	newInstance := *engine
	return &newInstance
}

func (event *Event) GongCopy() GongstructIF {
	newInstance := *event
	return &newInstance
}

func (status *Status) GongCopy() GongstructIF {
	newInstance := *status
	return &newInstance
}

func (updatestate *UpdateState) GongCopy() GongstructIF {
	newInstance := *updatestate
	return &newInstance
}

func (stage *Stage) ComputeForwardAndBackwardCommits() {
	var lenNewInstances int
	var lenModifiedInstances int
	var lenDeletedInstances int

	var newInstancesSlice []string
	var fieldsEditSlice []string
	var deletedInstancesSlice []string

	var newInstancesReverseSlice []string
	var fieldsEditReverseSlice []string
	var deletedInstancesReverseSlice []string

	// first clean the staging area to remove non staged instances
	// from pointers fields and slices of pointers fields
	stage.Clean()

	// insertion point per named struct
	var commands_newInstances []*Command
	var commands_deletedInstances []*Command

	// parse all staged instances and check if they have a reference
	for command := range stage.Commands {
		if ref, ok := stage.Commands_reference[command]; !ok {
			commands_newInstances = append(commands_newInstances, command)
			newInstancesSlice = append(newInstancesSlice, command.GongMarshallIdentifier(stage))
			if stage.Commands_referenceOrder == nil {
				stage.Commands_referenceOrder = make(map[*Command]uint)
			}
			stage.Commands_referenceOrder[command] = stage.CommandMap_Staged_Order[command]
			newInstancesReverseSlice = append(newInstancesReverseSlice, command.GongMarshallUnstaging(stage))
			delete(stage.Commands_referenceOrder, command)
			fieldInitializers, pointersInitializations := command.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.CommandMap_Staged_Order[ref] = stage.CommandMap_Staged_Order[command]
			diffs := command.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, command)
			delete(stage.CommandMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", command.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.Commands_reference {
		if _, ok := stage.Commands[ref]; !ok {
			commands_deletedInstances = append(commands_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(commands_newInstances)
	lenDeletedInstances += len(commands_deletedInstances)
	var dummyagents_newInstances []*DummyAgent
	var dummyagents_deletedInstances []*DummyAgent

	// parse all staged instances and check if they have a reference
	for dummyagent := range stage.DummyAgents {
		if ref, ok := stage.DummyAgents_reference[dummyagent]; !ok {
			dummyagents_newInstances = append(dummyagents_newInstances, dummyagent)
			newInstancesSlice = append(newInstancesSlice, dummyagent.GongMarshallIdentifier(stage))
			if stage.DummyAgents_referenceOrder == nil {
				stage.DummyAgents_referenceOrder = make(map[*DummyAgent]uint)
			}
			stage.DummyAgents_referenceOrder[dummyagent] = stage.DummyAgentMap_Staged_Order[dummyagent]
			newInstancesReverseSlice = append(newInstancesReverseSlice, dummyagent.GongMarshallUnstaging(stage))
			delete(stage.DummyAgents_referenceOrder, dummyagent)
			fieldInitializers, pointersInitializations := dummyagent.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.DummyAgentMap_Staged_Order[ref] = stage.DummyAgentMap_Staged_Order[dummyagent]
			diffs := dummyagent.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, dummyagent)
			delete(stage.DummyAgentMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", dummyagent.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.DummyAgents_reference {
		if _, ok := stage.DummyAgents[ref]; !ok {
			dummyagents_deletedInstances = append(dummyagents_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(dummyagents_newInstances)
	lenDeletedInstances += len(dummyagents_deletedInstances)
	var engines_newInstances []*Engine
	var engines_deletedInstances []*Engine

	// parse all staged instances and check if they have a reference
	for engine := range stage.Engines {
		if ref, ok := stage.Engines_reference[engine]; !ok {
			engines_newInstances = append(engines_newInstances, engine)
			newInstancesSlice = append(newInstancesSlice, engine.GongMarshallIdentifier(stage))
			if stage.Engines_referenceOrder == nil {
				stage.Engines_referenceOrder = make(map[*Engine]uint)
			}
			stage.Engines_referenceOrder[engine] = stage.EngineMap_Staged_Order[engine]
			newInstancesReverseSlice = append(newInstancesReverseSlice, engine.GongMarshallUnstaging(stage))
			delete(stage.Engines_referenceOrder, engine)
			fieldInitializers, pointersInitializations := engine.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.EngineMap_Staged_Order[ref] = stage.EngineMap_Staged_Order[engine]
			diffs := engine.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, engine)
			delete(stage.EngineMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", engine.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.Engines_reference {
		if _, ok := stage.Engines[ref]; !ok {
			engines_deletedInstances = append(engines_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(engines_newInstances)
	lenDeletedInstances += len(engines_deletedInstances)
	var events_newInstances []*Event
	var events_deletedInstances []*Event

	// parse all staged instances and check if they have a reference
	for event := range stage.Events {
		if ref, ok := stage.Events_reference[event]; !ok {
			events_newInstances = append(events_newInstances, event)
			newInstancesSlice = append(newInstancesSlice, event.GongMarshallIdentifier(stage))
			if stage.Events_referenceOrder == nil {
				stage.Events_referenceOrder = make(map[*Event]uint)
			}
			stage.Events_referenceOrder[event] = stage.EventMap_Staged_Order[event]
			newInstancesReverseSlice = append(newInstancesReverseSlice, event.GongMarshallUnstaging(stage))
			delete(stage.Events_referenceOrder, event)
			fieldInitializers, pointersInitializations := event.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.EventMap_Staged_Order[ref] = stage.EventMap_Staged_Order[event]
			diffs := event.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, event)
			delete(stage.EventMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", event.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.Events_reference {
		if _, ok := stage.Events[ref]; !ok {
			events_deletedInstances = append(events_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(events_newInstances)
	lenDeletedInstances += len(events_deletedInstances)
	var statuss_newInstances []*Status
	var statuss_deletedInstances []*Status

	// parse all staged instances and check if they have a reference
	for status := range stage.Statuss {
		if ref, ok := stage.Statuss_reference[status]; !ok {
			statuss_newInstances = append(statuss_newInstances, status)
			newInstancesSlice = append(newInstancesSlice, status.GongMarshallIdentifier(stage))
			if stage.Statuss_referenceOrder == nil {
				stage.Statuss_referenceOrder = make(map[*Status]uint)
			}
			stage.Statuss_referenceOrder[status] = stage.StatusMap_Staged_Order[status]
			newInstancesReverseSlice = append(newInstancesReverseSlice, status.GongMarshallUnstaging(stage))
			delete(stage.Statuss_referenceOrder, status)
			fieldInitializers, pointersInitializations := status.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.StatusMap_Staged_Order[ref] = stage.StatusMap_Staged_Order[status]
			diffs := status.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, status)
			delete(stage.StatusMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", status.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.Statuss_reference {
		if _, ok := stage.Statuss[ref]; !ok {
			statuss_deletedInstances = append(statuss_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(statuss_newInstances)
	lenDeletedInstances += len(statuss_deletedInstances)
	var updatestates_newInstances []*UpdateState
	var updatestates_deletedInstances []*UpdateState

	// parse all staged instances and check if they have a reference
	for updatestate := range stage.UpdateStates {
		if ref, ok := stage.UpdateStates_reference[updatestate]; !ok {
			updatestates_newInstances = append(updatestates_newInstances, updatestate)
			newInstancesSlice = append(newInstancesSlice, updatestate.GongMarshallIdentifier(stage))
			if stage.UpdateStates_referenceOrder == nil {
				stage.UpdateStates_referenceOrder = make(map[*UpdateState]uint)
			}
			stage.UpdateStates_referenceOrder[updatestate] = stage.UpdateStateMap_Staged_Order[updatestate]
			newInstancesReverseSlice = append(newInstancesReverseSlice, updatestate.GongMarshallUnstaging(stage))
			delete(stage.UpdateStates_referenceOrder, updatestate)
			fieldInitializers, pointersInitializations := updatestate.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.UpdateStateMap_Staged_Order[ref] = stage.UpdateStateMap_Staged_Order[updatestate]
			diffs := updatestate.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, updatestate)
			delete(stage.UpdateStateMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", updatestate.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.UpdateStates_reference {
		if _, ok := stage.UpdateStates[ref]; !ok {
			updatestates_deletedInstances = append(updatestates_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(updatestates_newInstances)
	lenDeletedInstances += len(updatestates_deletedInstances)

	if lenNewInstances > 0 || lenDeletedInstances > 0 || lenModifiedInstances > 0 {

		// sort the stmt to have reproductible forward/backward commit
		sort.Strings(newInstancesSlice)
		newInstancesStmt := strings.Join(newInstancesSlice, "")
		sort.Strings(fieldsEditSlice)
		fieldsEditStmt := strings.Join(fieldsEditSlice, "")
		sort.Strings(deletedInstancesSlice)
		deletedInstancesStmt := strings.Join(deletedInstancesSlice, "")

		sort.Strings(newInstancesReverseSlice)
		newInstancesReverseStmt := strings.Join(newInstancesReverseSlice, "")
		sort.Strings(fieldsEditReverseSlice)
		fieldsEditReverseStmt := strings.Join(fieldsEditReverseSlice, "")
		sort.Strings(deletedInstancesReverseSlice)
		deletedInstancesReverseStmt := strings.Join(deletedInstancesReverseSlice, "")

		forwardCommit := newInstancesStmt + fieldsEditStmt + deletedInstancesStmt
		forwardCommit += "\n\tstage.Commit()"
		stage.forwardCommits = append(stage.forwardCommits, forwardCommit)

		backwardCommit := deletedInstancesReverseStmt + fieldsEditReverseStmt + newInstancesReverseStmt
		backwardCommit += "\n\tstage.Commit()"
		// append to the end of the backward commits slice
		stage.backwardCommits = append(stage.backwardCommits, backwardCommit)

		if stage.GetProbeIF() != nil {
			var mergedCommits string
			for _, commit := range stage.forwardCommits {
				mergedCommits += commit
			}
			stage.GetProbeIF().AddNotification(
				time.Now(),
				"	// Forward commits:\n"+
					mergedCommits,
			)

			var reverseMergedCommits string
			for _, reverserCommit := range stage.backwardCommits {
				reverseMergedCommits += reverserCommit
			}
			stage.GetProbeIF().AddNotification(
				time.Now(),
				"	// Backward commits:\n"+
					reverseMergedCommits,
			)

			stage.GetProbeIF().CommitNotificationTable()
		}
	}
}

// ComputeReference will creates a deep copy of each of the staged elements
func (stage *Stage) ComputeReference() {

	// insertion point per named struct
	stage.Commands_reference = make(map[*Command]*Command)
	stage.Commands_referenceOrder = make(map[*Command]uint) // diff Unstage needs the reference order
	for instance := range stage.Commands {
		stage.Commands_reference[instance] = instance.GongCopy().(*Command)
		stage.Commands_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.DummyAgents_reference = make(map[*DummyAgent]*DummyAgent)
	stage.DummyAgents_referenceOrder = make(map[*DummyAgent]uint) // diff Unstage needs the reference order
	for instance := range stage.DummyAgents {
		stage.DummyAgents_reference[instance] = instance.GongCopy().(*DummyAgent)
		stage.DummyAgents_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.Engines_reference = make(map[*Engine]*Engine)
	stage.Engines_referenceOrder = make(map[*Engine]uint) // diff Unstage needs the reference order
	for instance := range stage.Engines {
		stage.Engines_reference[instance] = instance.GongCopy().(*Engine)
		stage.Engines_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.Events_reference = make(map[*Event]*Event)
	stage.Events_referenceOrder = make(map[*Event]uint) // diff Unstage needs the reference order
	for instance := range stage.Events {
		stage.Events_reference[instance] = instance.GongCopy().(*Event)
		stage.Events_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.Statuss_reference = make(map[*Status]*Status)
	stage.Statuss_referenceOrder = make(map[*Status]uint) // diff Unstage needs the reference order
	for instance := range stage.Statuss {
		stage.Statuss_reference[instance] = instance.GongCopy().(*Status)
		stage.Statuss_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.UpdateStates_reference = make(map[*UpdateState]*UpdateState)
	stage.UpdateStates_referenceOrder = make(map[*UpdateState]uint) // diff Unstage needs the reference order
	for instance := range stage.UpdateStates {
		stage.UpdateStates_reference[instance] = instance.GongCopy().(*UpdateState)
		stage.UpdateStates_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

}

// GongGetOrder returns the order of the instance in the staging area
// This order is set at staging time, and reflects the order of creation of the instances
// in the staging area
// It is used when rendering slices of GongstructIF to keep a deterministic order
// which is important for frontends such as web frontends
// to avoid unnecessary re-renderings
// insertion point per named struct
func (command *Command) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.CommandMap_Staged_Order[command]; ok {
		return order
	}
	return stage.Commands_referenceOrder[command]
}

func (command *Command) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Commands_referenceOrder[command]
}

func (dummyagent *DummyAgent) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.DummyAgentMap_Staged_Order[dummyagent]; ok {
		return order
	}
	return stage.DummyAgents_referenceOrder[dummyagent]
}

func (dummyagent *DummyAgent) GongGetReferenceOrder(stage *Stage) uint {
	return stage.DummyAgents_referenceOrder[dummyagent]
}

func (engine *Engine) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.EngineMap_Staged_Order[engine]; ok {
		return order
	}
	return stage.Engines_referenceOrder[engine]
}

func (engine *Engine) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Engines_referenceOrder[engine]
}

func (event *Event) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.EventMap_Staged_Order[event]; ok {
		return order
	}
	return stage.Events_referenceOrder[event]
}

func (event *Event) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Events_referenceOrder[event]
}

func (status *Status) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.StatusMap_Staged_Order[status]; ok {
		return order
	}
	return stage.Statuss_referenceOrder[status]
}

func (status *Status) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Statuss_referenceOrder[status]
}

func (updatestate *UpdateState) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.UpdateStateMap_Staged_Order[updatestate]; ok {
		return order
	}
	return stage.UpdateStates_referenceOrder[updatestate]
}

func (updatestate *UpdateState) GongGetReferenceOrder(stage *Stage) uint {
	return stage.UpdateStates_referenceOrder[updatestate]
}

// GongGetIdentifier returns a unique identifier of the instance in the staging area
// This identifier is composed of the Gongstruct name and the order of the instance
// in the staging area
// It is used to identify instances across sessions
// insertion point per named struct
func (command *Command) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", command.GongGetGongstructName(), command.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (command *Command) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", command.GongGetGongstructName(), command.GongGetReferenceOrder(stage))
}

func (dummyagent *DummyAgent) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", dummyagent.GongGetGongstructName(), dummyagent.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (dummyagent *DummyAgent) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", dummyagent.GongGetGongstructName(), dummyagent.GongGetReferenceOrder(stage))
}

func (engine *Engine) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", engine.GongGetGongstructName(), engine.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (engine *Engine) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", engine.GongGetGongstructName(), engine.GongGetReferenceOrder(stage))
}

func (event *Event) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", event.GongGetGongstructName(), event.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (event *Event) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", event.GongGetGongstructName(), event.GongGetReferenceOrder(stage))
}

func (status *Status) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", status.GongGetGongstructName(), status.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (status *Status) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", status.GongGetGongstructName(), status.GongGetReferenceOrder(stage))
}

func (updatestate *UpdateState) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", updatestate.GongGetGongstructName(), updatestate.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (updatestate *UpdateState) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", updatestate.GongGetGongstructName(), updatestate.GongGetReferenceOrder(stage))
}

// MarshallIdentifier returns the code to instantiate the instance
// in a marshalling file
// insertion point per named struct
func (command *Command) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", command.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Command")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", command.Name)
	return
}
func (dummyagent *DummyAgent) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", dummyagent.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "DummyAgent")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", dummyagent.Name)
	return
}
func (engine *Engine) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", engine.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Engine")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", engine.Name)
	return
}
func (event *Event) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", event.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Event")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", event.Name)
	return
}
func (status *Status) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", status.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Status")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", status.Name)
	return
}
func (updatestate *UpdateState) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = IdentifiersDeclsWithoutNameInit
	decl = strings.ReplaceAll(decl, "{{Identifier}}", updatestate.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "UpdateState")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", updatestate.Name)
	return
}

// insertion point for unstaging
func (command *Command) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", command.GongGetReferenceIdentifier(stage))
	return
}
func (dummyagent *DummyAgent) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", dummyagent.GongGetReferenceIdentifier(stage))
	return
}
func (engine *Engine) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", engine.GongGetReferenceIdentifier(stage))
	return
}
func (event *Event) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", event.GongGetReferenceIdentifier(stage))
	return
}
func (status *Status) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", status.GongGetReferenceIdentifier(stage))
	return
}
func (updatestate *UpdateState) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", updatestate.GongGetReferenceIdentifier(stage))
	return
}
