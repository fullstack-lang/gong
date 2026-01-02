// generated code - do not edit
package models

import (
	"fmt"
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

func (stage *Stage) ComputeDifference() {
	var lenNewInstances int
	var lenModifiedInstances int
	var lenDeletedInstances int

	var pointersInitializesStatements string

	// insertion point per named struct
	var commands_newInstances []*Command
	var commands_deletedInstances []*Command

	// parse all staged instances and check if they have a reference
	for command := range stage.Commands {
		if ref, ok := stage.Commands_reference[command]; !ok {
			commands_newInstances = append(commands_newInstances, command)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of Command "+command.Name,
				)
				stage.GetProbeIF().AddNotification(
					time.Now(),
					command.GongMarshallIdentifier(stage),
				)
				basicFieldInitializers, pointersInitializations := command.GongMarshallAllFields(stage)
				stage.GetProbeIF().AddNotification(
					time.Now(),
					basicFieldInitializers,
				)
				pointersInitializesStatements += pointersInitializations
			}
		} else {
			diffs := command.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of Command \""+command.Name+"\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for command := range stage.Commands_reference {
		if _, ok := stage.Commands[command]; !ok {
			commands_deletedInstances = append(commands_deletedInstances, command)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of Command "+command.Name,
				)
			}
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
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of DummyAgent "+dummyagent.Name,
				)
				stage.GetProbeIF().AddNotification(
					time.Now(),
					dummyagent.GongMarshallIdentifier(stage),
				)
				basicFieldInitializers, pointersInitializations := dummyagent.GongMarshallAllFields(stage)
				stage.GetProbeIF().AddNotification(
					time.Now(),
					basicFieldInitializers,
				)
				pointersInitializesStatements += pointersInitializations
			}
		} else {
			diffs := dummyagent.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of DummyAgent \""+dummyagent.Name+"\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for dummyagent := range stage.DummyAgents_reference {
		if _, ok := stage.DummyAgents[dummyagent]; !ok {
			dummyagents_deletedInstances = append(dummyagents_deletedInstances, dummyagent)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of DummyAgent "+dummyagent.Name,
				)
			}
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
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of Engine "+engine.Name,
				)
				stage.GetProbeIF().AddNotification(
					time.Now(),
					engine.GongMarshallIdentifier(stage),
				)
				basicFieldInitializers, pointersInitializations := engine.GongMarshallAllFields(stage)
				stage.GetProbeIF().AddNotification(
					time.Now(),
					basicFieldInitializers,
				)
				pointersInitializesStatements += pointersInitializations
			}
		} else {
			diffs := engine.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of Engine \""+engine.Name+"\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for engine := range stage.Engines_reference {
		if _, ok := stage.Engines[engine]; !ok {
			engines_deletedInstances = append(engines_deletedInstances, engine)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of Engine "+engine.Name,
				)
			}
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
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of Event "+event.Name,
				)
				stage.GetProbeIF().AddNotification(
					time.Now(),
					event.GongMarshallIdentifier(stage),
				)
				basicFieldInitializers, pointersInitializations := event.GongMarshallAllFields(stage)
				stage.GetProbeIF().AddNotification(
					time.Now(),
					basicFieldInitializers,
				)
				pointersInitializesStatements += pointersInitializations
			}
		} else {
			diffs := event.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of Event \""+event.Name+"\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for event := range stage.Events_reference {
		if _, ok := stage.Events[event]; !ok {
			events_deletedInstances = append(events_deletedInstances, event)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of Event "+event.Name,
				)
			}
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
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of Status "+status.Name,
				)
				stage.GetProbeIF().AddNotification(
					time.Now(),
					status.GongMarshallIdentifier(stage),
				)
				basicFieldInitializers, pointersInitializations := status.GongMarshallAllFields(stage)
				stage.GetProbeIF().AddNotification(
					time.Now(),
					basicFieldInitializers,
				)
				pointersInitializesStatements += pointersInitializations
			}
		} else {
			diffs := status.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of Status \""+status.Name+"\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for status := range stage.Statuss_reference {
		if _, ok := stage.Statuss[status]; !ok {
			statuss_deletedInstances = append(statuss_deletedInstances, status)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of Status "+status.Name,
				)
			}
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
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of UpdateState "+updatestate.Name,
				)
				stage.GetProbeIF().AddNotification(
					time.Now(),
					updatestate.GongMarshallIdentifier(stage),
				)
				basicFieldInitializers, pointersInitializations := updatestate.GongMarshallAllFields(stage)
				stage.GetProbeIF().AddNotification(
					time.Now(),
					basicFieldInitializers,
				)
				pointersInitializesStatements += pointersInitializations
			}
		} else {
			diffs := updatestate.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of UpdateState \""+updatestate.Name+"\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for updatestate := range stage.UpdateStates_reference {
		if _, ok := stage.UpdateStates[updatestate]; !ok {
			updatestates_deletedInstances = append(updatestates_deletedInstances, updatestate)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of UpdateState "+updatestate.Name,
				)
			}
		}
	}

	lenNewInstances += len(updatestates_newInstances)
	lenDeletedInstances += len(updatestates_deletedInstances)

	if lenNewInstances > 0 || lenDeletedInstances > 0 || lenModifiedInstances > 0 {
		// if stage.GetProbeIF() != nil {
		// 	stage.GetProbeIF().CommitNotificationTable()
		// }
	}

	if pointersInitializesStatements != "" {
		if stage.GetProbeIF() != nil {
			stage.GetProbeIF().AddNotification(
				time.Now(),
				pointersInitializesStatements,
			)
		}
	}
}

// ComputeReference will creates a deep copy of each of the staged elements
func (stage *Stage) ComputeReference() {

	// insertion point per named struct
	stage.Commands_reference = make(map[*Command]*Command)
	for instance := range stage.Commands {
		stage.Commands_reference[instance] = instance.GongCopy().(*Command)
	}

	stage.DummyAgents_reference = make(map[*DummyAgent]*DummyAgent)
	for instance := range stage.DummyAgents {
		stage.DummyAgents_reference[instance] = instance.GongCopy().(*DummyAgent)
	}

	stage.Engines_reference = make(map[*Engine]*Engine)
	for instance := range stage.Engines {
		stage.Engines_reference[instance] = instance.GongCopy().(*Engine)
	}

	stage.Events_reference = make(map[*Event]*Event)
	for instance := range stage.Events {
		stage.Events_reference[instance] = instance.GongCopy().(*Event)
	}

	stage.Statuss_reference = make(map[*Status]*Status)
	for instance := range stage.Statuss {
		stage.Statuss_reference[instance] = instance.GongCopy().(*Status)
	}

	stage.UpdateStates_reference = make(map[*UpdateState]*UpdateState)
	for instance := range stage.UpdateStates {
		stage.UpdateStates_reference[instance] = instance.GongCopy().(*UpdateState)
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
	return stage.CommandMap_Staged_Order[command]
}

func (dummyagent *DummyAgent) GongGetOrder(stage *Stage) uint {
	return stage.DummyAgentMap_Staged_Order[dummyagent]
}

func (engine *Engine) GongGetOrder(stage *Stage) uint {
	return stage.EngineMap_Staged_Order[engine]
}

func (event *Event) GongGetOrder(stage *Stage) uint {
	return stage.EventMap_Staged_Order[event]
}

func (status *Status) GongGetOrder(stage *Stage) uint {
	return stage.StatusMap_Staged_Order[status]
}

func (updatestate *UpdateState) GongGetOrder(stage *Stage) uint {
	return stage.UpdateStateMap_Staged_Order[updatestate]
}

// GongGetIdentifier returns a unique identifier of the instance in the staging area
// This identifier is composed of the Gongstruct name and the order of the instance
// in the staging area
// It is used to identify instances across sessions
// insertion point per named struct
func (command *Command) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", command.GongGetGongstructName(), command.GongGetOrder(stage))
}

func (dummyagent *DummyAgent) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", dummyagent.GongGetGongstructName(), dummyagent.GongGetOrder(stage))
}

func (engine *Engine) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", engine.GongGetGongstructName(), engine.GongGetOrder(stage))
}

func (event *Event) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", event.GongGetGongstructName(), event.GongGetOrder(stage))
}

func (status *Status) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", status.GongGetGongstructName(), status.GongGetOrder(stage))
}

func (updatestate *UpdateState) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", updatestate.GongGetGongstructName(), updatestate.GongGetOrder(stage))
}

// MarshallIdentifier returns the code to instantiate the instance
// in a marshalling file
// insertion point per named struct
func (command *Command) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = IdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", command.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Command")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", command.Name)
	return
}
func (dummyagent *DummyAgent) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = IdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", dummyagent.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "DummyAgent")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", dummyagent.Name)
	return
}
func (engine *Engine) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = IdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", engine.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Engine")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", engine.Name)
	return
}
func (event *Event) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = IdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", event.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Event")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", event.Name)
	return
}
func (status *Status) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = IdentifiersDecls
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
