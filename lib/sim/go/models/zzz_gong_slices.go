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
	// end of insertion point per named struct
}

func (stage *Stage) GetInstances() (res []GongstructIF) {
	// insertion point per named struct
	res = __gong__appendInstances(res, stage.Commands)

	res = __gong__appendInstances(res, stage.DummyAgents)

	res = __gong__appendInstances(res, stage.Engines)

	res = __gong__appendInstances(res, stage.Events)

	res = __gong__appendInstances(res, stage.Statuss)

	res = __gong__appendInstances(res, stage.UpdateStates)

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
func (command *Command) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, command)
}

func (dummyagent *DummyAgent) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, dummyagent)
}

func (engine *Engine) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, engine)
}

func (event *Event) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, event)
}

func (status *Status) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, status)
}

func (updatestate *UpdateState) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, updatestate)
}


type GongstructDiffable[T any] interface {
	GongstructPtr
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
	__gong__computeReferencePass1(stage, stage.Commands, &stage.Commands_reference, &stage.Commands_referenceOrder, &stage.Commands_instance)

	__gong__computeReferencePass1(stage, stage.DummyAgents, &stage.DummyAgents_reference, &stage.DummyAgents_referenceOrder, &stage.DummyAgents_instance)

	__gong__computeReferencePass1(stage, stage.Engines, &stage.Engines_reference, &stage.Engines_referenceOrder, &stage.Engines_instance)

	__gong__computeReferencePass1(stage, stage.Events, &stage.Events_reference, &stage.Events_referenceOrder, &stage.Events_instance)

	__gong__computeReferencePass1(stage, stage.Statuss, &stage.Statuss_reference, &stage.Statuss_referenceOrder, &stage.Statuss_instance)

	__gong__computeReferencePass1(stage, stage.UpdateStates, &stage.UpdateStates_reference, &stage.UpdateStates_referenceOrder, &stage.UpdateStates_instance)

	// insertion point per named struct
	__gong__computeReferencePass2(stage.Commands, stage.Commands_reference, stage)

	__gong__computeReferencePass2(stage.DummyAgents, stage.DummyAgents_reference, stage)

	__gong__computeReferencePass2(stage.Engines, stage.Engines_reference, stage)

	__gong__computeReferencePass2(stage.Events, stage.Events_reference, stage)

	__gong__computeReferencePass2(stage.Statuss, stage.Statuss_reference, stage)

	__gong__computeReferencePass2(stage.UpdateStates, stage.UpdateStates_reference, stage)

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
	return __gong__getOrder(stage.Command_stagedOrder, stage.Commands_referenceOrder, command, "Command")
}

func (dummyagent *DummyAgent) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.DummyAgent_stagedOrder, stage.DummyAgents_referenceOrder, dummyagent, "DummyAgent")
}

func (engine *Engine) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Engine_stagedOrder, stage.Engines_referenceOrder, engine, "Engine")
}

func (event *Event) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Event_stagedOrder, stage.Events_referenceOrder, event, "Event")
}

func (status *Status) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Status_stagedOrder, stage.Statuss_referenceOrder, status, "Status")
}

func (updatestate *UpdateState) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.UpdateState_stagedOrder, stage.UpdateStates_referenceOrder, updatestate, "UpdateState")
}

// GongGetIdentifier returns a unique identifier of the instance in the staging area
// This identifier is composed of the Gongstruct name and the order of the instance
// in the staging area
// It is used to identify instances across sessions
// insertion point per named struct
func (command *Command) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(command, command.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (command *Command) GongGetReferenceIdentifier(stage *Stage) string {
	return command.GongGetIdentifier(stage)
}

func (dummyagent *DummyAgent) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(dummyagent, dummyagent.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (dummyagent *DummyAgent) GongGetReferenceIdentifier(stage *Stage) string {
	return dummyagent.GongGetIdentifier(stage)
}

func (engine *Engine) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(engine, engine.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (engine *Engine) GongGetReferenceIdentifier(stage *Stage) string {
	return engine.GongGetIdentifier(stage)
}

func (event *Event) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(event, event.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (event *Event) GongGetReferenceIdentifier(stage *Stage) string {
	return event.GongGetIdentifier(stage)
}

func (status *Status) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(status, status.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (status *Status) GongGetReferenceIdentifier(stage *Stage) string {
	return status.GongGetIdentifier(stage)
}

func (updatestate *UpdateState) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(updatestate, updatestate.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (updatestate *UpdateState) GongGetReferenceIdentifier(stage *Stage) string {
	return updatestate.GongGetIdentifier(stage)
}

// MarshallIdentifier returns the code to instantiate the instance
// in a marshalling file
// insertion point per named struct
func (command *Command) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(command.GongGetIdentifier(stage), "Command", command.Name)
}

func (dummyagent *DummyAgent) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(dummyagent.GongGetIdentifier(stage), "DummyAgent", dummyagent.Name)
}

func (engine *Engine) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(engine.GongGetIdentifier(stage), "Engine", engine.Name)
}

func (event *Event) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(event.GongGetIdentifier(stage), "Event", event.Name)
}

func (status *Status) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(status.GongGetIdentifier(stage), "Status", status.Name)
}

func (updatestate *UpdateState) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(updatestate.GongGetIdentifier(stage), "UpdateState", updatestate.Name)
}

// insertion point for unstaging
func (command *Command) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(command.GongGetReferenceIdentifier(stage))
}

func (dummyagent *DummyAgent) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(dummyagent.GongGetReferenceIdentifier(stage))
}

func (engine *Engine) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(engine.GongGetReferenceIdentifier(stage))
}

func (event *Event) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(event.GongGetReferenceIdentifier(stage))
}

func (status *Status) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(status.GongGetReferenceIdentifier(stage))
}

func (updatestate *UpdateState) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(updatestate.GongGetReferenceIdentifier(stage))
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

func __gong__appendInstances[T interface {
	comparable
	GongstructIF
}](res []GongstructIF, m map[T]struct{}) []GongstructIF {
	for instance := range m {
		res = append(res, instance)
	}
	return res
}

func __gong__getUUID(stage *Stage, instance GongstructIF) string {
	if __gong__, ok := any(instance).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}
	return GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(instance), uint64(stage.GetOrder(instance)))
}

func __gong__computeReferencePass1[T interface {
	comparable
	GongstructIF
}](
	stage *Stage,
	staged map[T]struct{},
	ref *map[T]T,
	refOrder *map[T]uint,
	inst *map[T]T,
) {
	*ref = make(map[T]T, len(staged))
	*refOrder = make(map[T]uint, len(staged))
	*inst = make(map[T]T, len(staged))
	for instance := range staged {
		_copy := instance.GongCopy().(T)
		(*ref)[instance] = _copy
		(*inst)[_copy] = instance
		(*refOrder)[_copy] = instance.GongGetOrder(stage)
	}
}

func __gong__computeReferencePass2[T interface {
	comparable
	GongstructIF
	GongReconstructPointersFromReferences(*Stage, T)
}](staged map[T]struct{}, reference map[T]T, stage *Stage) {
	for instance := range staged {
		reference[instance].GongReconstructPointersFromReferences(stage, instance)
	}
}

func __gong__getOrder[T comparable](stagedOrder, refOrder map[T]uint, instance T, typeName string) uint {
	if order, ok := stagedOrder[instance]; ok {
		return order
	}
	if order, ok := refOrder[instance]; ok {
		return order
	}
	log.Printf("instance %p of type %s was not staged and does not have a reference order", any(instance), typeName)
	return 0
}

func __gong__formatIdentifier(s GongstructIF, order uint) string {
	return fmt.Sprintf("__%s__%08d_", s.GongGetGongstructName(), order)
}

func __gong__marshallIdentifier(identifier, structName, name string) string {
	decl := strings.ReplaceAll(GongIdentifiersDecls, "{{Identifier}}", identifier)
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", structName)
	return strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(name))
}

func __gong__marshallUnstaging(identifier string) string {
	return strings.ReplaceAll(GongUnstageStmt, "{{Identifier}}", identifier)
}

// end of template
