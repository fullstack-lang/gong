// generated code - do not edit
package models

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"sort"
	"strings"
	"time"
)

var (
	__GongSliceTemplate_time__dummyDeclaration time.Duration
	_                                          = __GongSliceTemplate_time__dummyDeclaration
)

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

	// end of insertion point per named struct
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
	newInstance := new(Command)
	command.GongCopyBasicFields(newInstance)
	return newInstance
}

func (dummyagent *DummyAgent) GongCopy() GongstructIF {
	newInstance := new(DummyAgent)
	dummyagent.GongCopyBasicFields(newInstance)
	return newInstance
}

func (engine *Engine) GongCopy() GongstructIF {
	newInstance := new(Engine)
	engine.GongCopyBasicFields(newInstance)
	return newInstance
}

func (event *Event) GongCopy() GongstructIF {
	newInstance := new(Event)
	event.GongCopyBasicFields(newInstance)
	return newInstance
}

func (status *Status) GongCopy() GongstructIF {
	newInstance := new(Status)
	status.GongCopyBasicFields(newInstance)
	return newInstance
}

func (updatestate *UpdateState) GongCopy() GongstructIF {
	newInstance := new(UpdateState)
	updatestate.GongCopyBasicFields(newInstance)
	return newInstance
}

// insertion point per named struct
func (command *Command) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(command).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(command), uint64(stage.GetOrder(command)))
	return
}

func (dummyagent *DummyAgent) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(dummyagent).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(dummyagent), uint64(stage.GetOrder(dummyagent)))
	return
}

func (engine *Engine) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(engine).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(engine), uint64(stage.GetOrder(engine)))
	return
}

func (event *Event) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(event).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(event), uint64(stage.GetOrder(event)))
	return
}

func (status *Status) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(status).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(status), uint64(stage.GetOrder(status)))
	return
}

func (updatestate *UpdateState) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(updatestate).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(updatestate), uint64(stage.GetOrder(updatestate)))
	return
}


type GongstructDiffable[T any] interface {
	PointerToGongstruct
	GongMarshallIdentifier(stage *Stage) string
	GongMarshallUnstaging(stage *Stage) string
	GongMarshallAllFields(stage *Stage) (string, string)
	GongReconstructPointersFromInstances(stage *Stage)
	GongDiff(stage *Stage, other T) []string
}

func computeCommitsForType[T GongstructDiffable[T]](
	stage *Stage,
	stagedInstances map[T]struct{},
	stagedOrder map[T]uint,
	referenceInstances map[T]T,
	referenceOrder *map[T]uint,
	instancesMap map[T]T,
	newInstancesSlice *[]string,
	fieldsEditSlice *[]string,
	deletedInstancesSlice *[]string,
	newInstancesReverseSlice *[]string,
	fieldsEditReverseSlice *[]string,
	deletedInstancesReverseSlice *[]string,
	lenNewInstances *int,
	lenDeletedInstances *int,
	lenModifiedInstances *int,
) {
	var newInstances []T
	var deletedInstances []T

	// parse all staged instances and check if they have a reference
	for instance := range stagedInstances {
		if ref, ok := referenceInstances[instance]; !ok {
			newInstances = append(newInstances, instance)
			*newInstancesSlice = append(*newInstancesSlice, instance.GongMarshallIdentifier(stage))
			if *referenceOrder == nil {
				*referenceOrder = make(map[T]uint)
			}
			(*referenceOrder)[instance] = stagedOrder[instance]
			*newInstancesReverseSlice = append(*newInstancesReverseSlice, instance.GongMarshallUnstaging(stage))
			fieldInitializers, pointersInitializations := instance.GongMarshallAllFields(stage)
			*fieldsEditSlice = append(*fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stagedOrder[ref] = stagedOrder[instance]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := instance.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, instance)
			if len(diffs) > 0 {
				var fieldsEdit string
				if instance.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", instance.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				*fieldsEditSlice = append(*fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					*fieldsEditReverseSlice = append(*fieldsEditReverseSlice, reverseDiff)
				}
				*lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range referenceInstances {
		instance := instancesMap[ref] // get the instance corresponding to the reference
		if _, ok := stagedInstances[instance]; !ok { // if the instance is not staged anymore, it means it has been unstaged
			deletedInstances = append(deletedInstances, ref)
			*deletedInstancesSlice = append(*deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			*deletedInstancesReverseSlice = append(*deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			*fieldsEditReverseSlice = append(*fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	*lenNewInstances += len(newInstances)
	*lenDeletedInstances += len(deletedInstances)
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
	computeCommitsForType(
		stage,
		stage.Commands,
		stage.Command_stagedOrder,
		stage.Commands_reference,
		&stage.Commands_referenceOrder,
		stage.Commands_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.DummyAgents,
		stage.DummyAgent_stagedOrder,
		stage.DummyAgents_reference,
		&stage.DummyAgents_referenceOrder,
		stage.DummyAgents_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Engines,
		stage.Engine_stagedOrder,
		stage.Engines_reference,
		&stage.Engines_referenceOrder,
		stage.Engines_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Events,
		stage.Event_stagedOrder,
		stage.Events_reference,
		&stage.Events_referenceOrder,
		stage.Events_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Statuss,
		stage.Status_stagedOrder,
		stage.Statuss_reference,
		&stage.Statuss_referenceOrder,
		stage.Statuss_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.UpdateStates,
		stage.UpdateState_stagedOrder,
		stage.UpdateStates_reference,
		&stage.UpdateStates_referenceOrder,
		stage.UpdateStates_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)

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
		stage.modified = true
	} else {
		stage.modified = false
	}
}

// ComputeReferenceAndOrders will creates a deep copy of each of the staged elements
func (stage *Stage) ComputeReferenceAndOrders() {
	// insertion point per named struct
	stage.Commands_reference = make(map[*Command]*Command)
	stage.Commands_referenceOrder = make(map[*Command]uint) // diff Unstage needs the reference order
	stage.Commands_instance = make(map[*Command]*Command)
	for instance := range stage.Commands {
		_copy := instance.GongCopy().(*Command)
		stage.Commands_reference[instance] = _copy
		stage.Commands_instance[_copy] = instance
		stage.Commands_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.DummyAgents_reference = make(map[*DummyAgent]*DummyAgent)
	stage.DummyAgents_referenceOrder = make(map[*DummyAgent]uint) // diff Unstage needs the reference order
	stage.DummyAgents_instance = make(map[*DummyAgent]*DummyAgent)
	for instance := range stage.DummyAgents {
		_copy := instance.GongCopy().(*DummyAgent)
		stage.DummyAgents_reference[instance] = _copy
		stage.DummyAgents_instance[_copy] = instance
		stage.DummyAgents_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Engines_reference = make(map[*Engine]*Engine)
	stage.Engines_referenceOrder = make(map[*Engine]uint) // diff Unstage needs the reference order
	stage.Engines_instance = make(map[*Engine]*Engine)
	for instance := range stage.Engines {
		_copy := instance.GongCopy().(*Engine)
		stage.Engines_reference[instance] = _copy
		stage.Engines_instance[_copy] = instance
		stage.Engines_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Events_reference = make(map[*Event]*Event)
	stage.Events_referenceOrder = make(map[*Event]uint) // diff Unstage needs the reference order
	stage.Events_instance = make(map[*Event]*Event)
	for instance := range stage.Events {
		_copy := instance.GongCopy().(*Event)
		stage.Events_reference[instance] = _copy
		stage.Events_instance[_copy] = instance
		stage.Events_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Statuss_reference = make(map[*Status]*Status)
	stage.Statuss_referenceOrder = make(map[*Status]uint) // diff Unstage needs the reference order
	stage.Statuss_instance = make(map[*Status]*Status)
	for instance := range stage.Statuss {
		_copy := instance.GongCopy().(*Status)
		stage.Statuss_reference[instance] = _copy
		stage.Statuss_instance[_copy] = instance
		stage.Statuss_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.UpdateStates_reference = make(map[*UpdateState]*UpdateState)
	stage.UpdateStates_referenceOrder = make(map[*UpdateState]uint) // diff Unstage needs the reference order
	stage.UpdateStates_instance = make(map[*UpdateState]*UpdateState)
	for instance := range stage.UpdateStates {
		_copy := instance.GongCopy().(*UpdateState)
		stage.UpdateStates_reference[instance] = _copy
		stage.UpdateStates_instance[_copy] = instance
		stage.UpdateStates_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	// insertion point per named struct
	for instance := range stage.Commands {
		reference := stage.Commands_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.DummyAgents {
		reference := stage.DummyAgents_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Engines {
		reference := stage.Engines_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Events {
		reference := stage.Events_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Statuss {
		reference := stage.Statuss_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.UpdateStates {
		reference := stage.UpdateStates_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	stage.recomputeOrders()
}

// GongGetOrder returns the order of the instance in the staging area
// This order is set at staging time, and reflects the order of creation of the instances
// in the staging area
// It is used when rendering slices of GongstructIF to keep a deterministic order
// which is important for frontends such as web frontends
// to avoid unnecessary re-renderings
// insertion point per named struct
func (command *Command) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Command_stagedOrder[command]; ok {
		return order
	}
	if order, ok := stage.Commands_referenceOrder[command]; ok {
		return order
	} else {
		log.Printf("instance %p of type Command was not staged and does not have a reference order", command)
		return 0
	}
}

func (dummyagent *DummyAgent) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.DummyAgent_stagedOrder[dummyagent]; ok {
		return order
	}
	if order, ok := stage.DummyAgents_referenceOrder[dummyagent]; ok {
		return order
	} else {
		log.Printf("instance %p of type DummyAgent was not staged and does not have a reference order", dummyagent)
		return 0
	}
}

func (engine *Engine) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Engine_stagedOrder[engine]; ok {
		return order
	}
	if order, ok := stage.Engines_referenceOrder[engine]; ok {
		return order
	} else {
		log.Printf("instance %p of type Engine was not staged and does not have a reference order", engine)
		return 0
	}
}

func (event *Event) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Event_stagedOrder[event]; ok {
		return order
	}
	if order, ok := stage.Events_referenceOrder[event]; ok {
		return order
	} else {
		log.Printf("instance %p of type Event was not staged and does not have a reference order", event)
		return 0
	}
}

func (status *Status) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Status_stagedOrder[status]; ok {
		return order
	}
	if order, ok := stage.Statuss_referenceOrder[status]; ok {
		return order
	} else {
		log.Printf("instance %p of type Status was not staged and does not have a reference order", status)
		return 0
	}
}

func (updatestate *UpdateState) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.UpdateState_stagedOrder[updatestate]; ok {
		return order
	}
	if order, ok := stage.UpdateStates_referenceOrder[updatestate]; ok {
		return order
	} else {
		log.Printf("instance %p of type UpdateState was not staged and does not have a reference order", updatestate)
		return 0
	}
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
	return fmt.Sprintf("__%s__%08d_", command.GongGetGongstructName(), command.GongGetOrder(stage))
}

func (dummyagent *DummyAgent) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", dummyagent.GongGetGongstructName(), dummyagent.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (dummyagent *DummyAgent) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", dummyagent.GongGetGongstructName(), dummyagent.GongGetOrder(stage))
}

func (engine *Engine) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", engine.GongGetGongstructName(), engine.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (engine *Engine) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", engine.GongGetGongstructName(), engine.GongGetOrder(stage))
}

func (event *Event) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", event.GongGetGongstructName(), event.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (event *Event) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", event.GongGetGongstructName(), event.GongGetOrder(stage))
}

func (status *Status) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", status.GongGetGongstructName(), status.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (status *Status) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", status.GongGetGongstructName(), status.GongGetOrder(stage))
}

func (updatestate *UpdateState) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", updatestate.GongGetGongstructName(), updatestate.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (updatestate *UpdateState) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", updatestate.GongGetGongstructName(), updatestate.GongGetOrder(stage))
}

// MarshallIdentifier returns the code to instantiate the instance
// in a marshalling file
// insertion point per named struct
func (command *Command) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", command.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Command")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(command.Name))
	return
}

func (dummyagent *DummyAgent) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", dummyagent.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "DummyAgent")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(dummyagent.Name))
	return
}

func (engine *Engine) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", engine.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Engine")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(engine.Name))
	return
}

func (event *Event) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", event.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Event")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(event.Name))
	return
}

func (status *Status) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", status.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Status")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(status.Name))
	return
}

func (updatestate *UpdateState) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", updatestate.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "UpdateState")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(updatestate.Name))
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

func GongIntToLetters(number int32) (letters string) {
	number--
	if firstLetter := number / 26; firstLetter > 0 {
		letters += GongIntToLetters(firstLetter)
		letters += string('A' + number%26)
	} else {
		letters += string('A' + number)
	}

	return
}

// GongGenerateReproducibleUUIDv4 creates a deterministic UUIDv4 based on a string and a positive integer.
func GongGenerateReproducibleUUIDv4(seedStr string, seedInt uint64) string {
	// 1. Create a deterministic hash from the inputs using SHA-256
	h := sha256.New()

	// Write the string to the hash
	h.Write([]byte(seedStr))

	// Write the integer to the hash (using BigEndian to ensure consistency across architectures)
	intBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(intBytes, seedInt)
	h.Write(intBytes)

	// 2. Extract the first 16 bytes from our resulting hash
	hashBytes := h.Sum(nil)
	uuid := make([]byte, 16)
	copy(uuid, hashBytes[:16])

	// 3. Set the Version to 4 (0100 in binary)
	// We take the 7th byte, clear the top 4 bits with & 0x0f, and set the top bits to 0100 with | 0x40
	uuid[6] = (uuid[6] & 0x0f) | 0x40

	// 4. Set the Variant to RFC4122 (10 in binary)
	// We take the 9th byte, clear the top 2 bits with & 0x3f, and set the top bits to 10 with | 0x80
	uuid[8] = (uuid[8] & 0x3f) | 0x80

	// 5. Format and return the byte array as a standard UUID string
	return fmt.Sprintf("%08x-%04x-%04x-%04x-%012x",
		uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:16])
}

// end of template
