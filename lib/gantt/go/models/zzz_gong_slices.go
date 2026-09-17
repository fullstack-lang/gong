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
	// Compute reverse map for named struct Arrow
	// insertion point per field

	// Compute reverse map for named struct Bar
	// insertion point per field

	// Compute reverse map for named struct Gantt
	// insertion point per field
	stage.Gantt_Lanes_reverseMap = make(map[*Lane]*Gantt)
	for gantt := range stage.Gantts {
		_ = gantt
		for _, _lane := range gantt.Lanes {
			stage.Gantt_Lanes_reverseMap[_lane] = gantt
		}
	}
	stage.Gantt_Milestones_reverseMap = make(map[*Milestone]*Gantt)
	for gantt := range stage.Gantts {
		_ = gantt
		for _, _milestone := range gantt.Milestones {
			stage.Gantt_Milestones_reverseMap[_milestone] = gantt
		}
	}
	stage.Gantt_Groups_reverseMap = make(map[*Group]*Gantt)
	for gantt := range stage.Gantts {
		_ = gantt
		for _, _group := range gantt.Groups {
			stage.Gantt_Groups_reverseMap[_group] = gantt
		}
	}
	stage.Gantt_Arrows_reverseMap = make(map[*Arrow]*Gantt)
	for gantt := range stage.Gantts {
		_ = gantt
		for _, _arrow := range gantt.Arrows {
			stage.Gantt_Arrows_reverseMap[_arrow] = gantt
		}
	}

	// Compute reverse map for named struct Group
	// insertion point per field
	stage.Group_GroupLanes_reverseMap = make(map[*Lane]*Group)
	for group := range stage.Groups {
		_ = group
		for _, _lane := range group.GroupLanes {
			stage.Group_GroupLanes_reverseMap[_lane] = group
		}
	}

	// Compute reverse map for named struct Lane
	// insertion point per field
	stage.Lane_Bars_reverseMap = make(map[*Bar]*Lane)
	for lane := range stage.Lanes {
		_ = lane
		for _, _bar := range lane.Bars {
			stage.Lane_Bars_reverseMap[_bar] = lane
		}
	}

	// Compute reverse map for named struct LaneUse
	// insertion point per field

	// Compute reverse map for named struct Milestone
	// insertion point per field
	stage.Milestone_LanesToDisplay_reverseMap = make(map[*Lane]*Milestone)
	for milestone := range stage.Milestones {
		_ = milestone
		for _, _lane := range milestone.LanesToDisplay {
			stage.Milestone_LanesToDisplay_reverseMap[_lane] = milestone
		}
	}

	// end of insertion point per named struct
}

func (stage *Stage) GetInstances() (res []GongstructIF) {
	// insertion point per named struct
	for instance := range stage.Arrows {
		res = append(res, instance)
	}

	for instance := range stage.Bars {
		res = append(res, instance)
	}

	for instance := range stage.Gantts {
		res = append(res, instance)
	}

	for instance := range stage.Groups {
		res = append(res, instance)
	}

	for instance := range stage.Lanes {
		res = append(res, instance)
	}

	for instance := range stage.LaneUses {
		res = append(res, instance)
	}

	for instance := range stage.Milestones {
		res = append(res, instance)
	}

	return
}

// insertion point per named struct
func (arrow *Arrow) GongCopy() GongstructIF {
	newInstance := new(Arrow)
	arrow.GongCopyBasicFields(newInstance)
	return newInstance
}

func (bar *Bar) GongCopy() GongstructIF {
	newInstance := new(Bar)
	bar.GongCopyBasicFields(newInstance)
	return newInstance
}

func (gantt *Gantt) GongCopy() GongstructIF {
	newInstance := new(Gantt)
	gantt.GongCopyBasicFields(newInstance)
	return newInstance
}

func (group *Group) GongCopy() GongstructIF {
	newInstance := new(Group)
	group.GongCopyBasicFields(newInstance)
	return newInstance
}

func (lane *Lane) GongCopy() GongstructIF {
	newInstance := new(Lane)
	lane.GongCopyBasicFields(newInstance)
	return newInstance
}

func (laneuse *LaneUse) GongCopy() GongstructIF {
	newInstance := new(LaneUse)
	laneuse.GongCopyBasicFields(newInstance)
	return newInstance
}

func (milestone *Milestone) GongCopy() GongstructIF {
	newInstance := new(Milestone)
	milestone.GongCopyBasicFields(newInstance)
	return newInstance
}

// insertion point per named struct
func (arrow *Arrow) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(arrow).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(arrow), uint64(stage.GetOrder(arrow)))
	return
}

func (bar *Bar) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(bar).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(bar), uint64(stage.GetOrder(bar)))
	return
}

func (gantt *Gantt) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(gantt).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(gantt), uint64(stage.GetOrder(gantt)))
	return
}

func (group *Group) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(group).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(group), uint64(stage.GetOrder(group)))
	return
}

func (lane *Lane) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(lane).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(lane), uint64(stage.GetOrder(lane)))
	return
}

func (laneuse *LaneUse) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(laneuse).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(laneuse), uint64(stage.GetOrder(laneuse)))
	return
}

func (milestone *Milestone) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(milestone).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(milestone), uint64(stage.GetOrder(milestone)))
	return
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
		stage.Arrows,
		stage.Arrow_stagedOrder,
		stage.Arrows_reference,
		&stage.Arrows_referenceOrder,
		stage.Arrows_instance,
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
		stage.Bars,
		stage.Bar_stagedOrder,
		stage.Bars_reference,
		&stage.Bars_referenceOrder,
		stage.Bars_instance,
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
		stage.Gantts,
		stage.Gantt_stagedOrder,
		stage.Gantts_reference,
		&stage.Gantts_referenceOrder,
		stage.Gantts_instance,
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
		stage.Groups,
		stage.Group_stagedOrder,
		stage.Groups_reference,
		&stage.Groups_referenceOrder,
		stage.Groups_instance,
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
		stage.Lanes,
		stage.Lane_stagedOrder,
		stage.Lanes_reference,
		&stage.Lanes_referenceOrder,
		stage.Lanes_instance,
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
		stage.LaneUses,
		stage.LaneUse_stagedOrder,
		stage.LaneUses_reference,
		&stage.LaneUses_referenceOrder,
		stage.LaneUses_instance,
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
		stage.Milestones,
		stage.Milestone_stagedOrder,
		stage.Milestones_reference,
		&stage.Milestones_referenceOrder,
		stage.Milestones_instance,
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
	stage.Arrows_reference = make(map[*Arrow]*Arrow)
	stage.Arrows_referenceOrder = make(map[*Arrow]uint) // diff Unstage needs the reference order
	stage.Arrows_instance = make(map[*Arrow]*Arrow)
	for instance := range stage.Arrows {
		_copy := instance.GongCopy().(*Arrow)
		stage.Arrows_reference[instance] = _copy
		stage.Arrows_instance[_copy] = instance
		stage.Arrows_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Bars_reference = make(map[*Bar]*Bar)
	stage.Bars_referenceOrder = make(map[*Bar]uint) // diff Unstage needs the reference order
	stage.Bars_instance = make(map[*Bar]*Bar)
	for instance := range stage.Bars {
		_copy := instance.GongCopy().(*Bar)
		stage.Bars_reference[instance] = _copy
		stage.Bars_instance[_copy] = instance
		stage.Bars_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Gantts_reference = make(map[*Gantt]*Gantt)
	stage.Gantts_referenceOrder = make(map[*Gantt]uint) // diff Unstage needs the reference order
	stage.Gantts_instance = make(map[*Gantt]*Gantt)
	for instance := range stage.Gantts {
		_copy := instance.GongCopy().(*Gantt)
		stage.Gantts_reference[instance] = _copy
		stage.Gantts_instance[_copy] = instance
		stage.Gantts_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Groups_reference = make(map[*Group]*Group)
	stage.Groups_referenceOrder = make(map[*Group]uint) // diff Unstage needs the reference order
	stage.Groups_instance = make(map[*Group]*Group)
	for instance := range stage.Groups {
		_copy := instance.GongCopy().(*Group)
		stage.Groups_reference[instance] = _copy
		stage.Groups_instance[_copy] = instance
		stage.Groups_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Lanes_reference = make(map[*Lane]*Lane)
	stage.Lanes_referenceOrder = make(map[*Lane]uint) // diff Unstage needs the reference order
	stage.Lanes_instance = make(map[*Lane]*Lane)
	for instance := range stage.Lanes {
		_copy := instance.GongCopy().(*Lane)
		stage.Lanes_reference[instance] = _copy
		stage.Lanes_instance[_copy] = instance
		stage.Lanes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.LaneUses_reference = make(map[*LaneUse]*LaneUse)
	stage.LaneUses_referenceOrder = make(map[*LaneUse]uint) // diff Unstage needs the reference order
	stage.LaneUses_instance = make(map[*LaneUse]*LaneUse)
	for instance := range stage.LaneUses {
		_copy := instance.GongCopy().(*LaneUse)
		stage.LaneUses_reference[instance] = _copy
		stage.LaneUses_instance[_copy] = instance
		stage.LaneUses_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Milestones_reference = make(map[*Milestone]*Milestone)
	stage.Milestones_referenceOrder = make(map[*Milestone]uint) // diff Unstage needs the reference order
	stage.Milestones_instance = make(map[*Milestone]*Milestone)
	for instance := range stage.Milestones {
		_copy := instance.GongCopy().(*Milestone)
		stage.Milestones_reference[instance] = _copy
		stage.Milestones_instance[_copy] = instance
		stage.Milestones_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	// insertion point per named struct
	for instance := range stage.Arrows {
		reference := stage.Arrows_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Bars {
		reference := stage.Bars_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Gantts {
		reference := stage.Gantts_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Groups {
		reference := stage.Groups_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Lanes {
		reference := stage.Lanes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.LaneUses {
		reference := stage.LaneUses_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Milestones {
		reference := stage.Milestones_reference[instance]
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
func (arrow *Arrow) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Arrow_stagedOrder[arrow]; ok {
		return order
	}
	if order, ok := stage.Arrows_referenceOrder[arrow]; ok {
		return order
	} else {
		log.Printf("instance %p of type Arrow was not staged and does not have a reference order", arrow)
		return 0
	}
}

func (bar *Bar) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Bar_stagedOrder[bar]; ok {
		return order
	}
	if order, ok := stage.Bars_referenceOrder[bar]; ok {
		return order
	} else {
		log.Printf("instance %p of type Bar was not staged and does not have a reference order", bar)
		return 0
	}
}

func (gantt *Gantt) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Gantt_stagedOrder[gantt]; ok {
		return order
	}
	if order, ok := stage.Gantts_referenceOrder[gantt]; ok {
		return order
	} else {
		log.Printf("instance %p of type Gantt was not staged and does not have a reference order", gantt)
		return 0
	}
}

func (group *Group) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Group_stagedOrder[group]; ok {
		return order
	}
	if order, ok := stage.Groups_referenceOrder[group]; ok {
		return order
	} else {
		log.Printf("instance %p of type Group was not staged and does not have a reference order", group)
		return 0
	}
}

func (lane *Lane) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Lane_stagedOrder[lane]; ok {
		return order
	}
	if order, ok := stage.Lanes_referenceOrder[lane]; ok {
		return order
	} else {
		log.Printf("instance %p of type Lane was not staged and does not have a reference order", lane)
		return 0
	}
}

func (laneuse *LaneUse) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.LaneUse_stagedOrder[laneuse]; ok {
		return order
	}
	if order, ok := stage.LaneUses_referenceOrder[laneuse]; ok {
		return order
	} else {
		log.Printf("instance %p of type LaneUse was not staged and does not have a reference order", laneuse)
		return 0
	}
}

func (milestone *Milestone) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Milestone_stagedOrder[milestone]; ok {
		return order
	}
	if order, ok := stage.Milestones_referenceOrder[milestone]; ok {
		return order
	} else {
		log.Printf("instance %p of type Milestone was not staged and does not have a reference order", milestone)
		return 0
	}
}

// GongGetIdentifier returns a unique identifier of the instance in the staging area
// This identifier is composed of the Gongstruct name and the order of the instance
// in the staging area
// It is used to identify instances across sessions
// insertion point per named struct
func (arrow *Arrow) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", arrow.GongGetGongstructName(), arrow.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (arrow *Arrow) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", arrow.GongGetGongstructName(), arrow.GongGetOrder(stage))
}

func (bar *Bar) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", bar.GongGetGongstructName(), bar.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (bar *Bar) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", bar.GongGetGongstructName(), bar.GongGetOrder(stage))
}

func (gantt *Gantt) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", gantt.GongGetGongstructName(), gantt.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (gantt *Gantt) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", gantt.GongGetGongstructName(), gantt.GongGetOrder(stage))
}

func (group *Group) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", group.GongGetGongstructName(), group.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (group *Group) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", group.GongGetGongstructName(), group.GongGetOrder(stage))
}

func (lane *Lane) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", lane.GongGetGongstructName(), lane.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (lane *Lane) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", lane.GongGetGongstructName(), lane.GongGetOrder(stage))
}

func (laneuse *LaneUse) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", laneuse.GongGetGongstructName(), laneuse.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (laneuse *LaneUse) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", laneuse.GongGetGongstructName(), laneuse.GongGetOrder(stage))
}

func (milestone *Milestone) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", milestone.GongGetGongstructName(), milestone.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (milestone *Milestone) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", milestone.GongGetGongstructName(), milestone.GongGetOrder(stage))
}

// MarshallIdentifier returns the code to instantiate the instance
// in a marshalling file
// insertion point per named struct
func (arrow *Arrow) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", arrow.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Arrow")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(arrow.Name))
	return
}

func (bar *Bar) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", bar.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Bar")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(bar.Name))
	return
}

func (gantt *Gantt) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", gantt.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Gantt")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(gantt.Name))
	return
}

func (group *Group) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", group.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Group")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(group.Name))
	return
}

func (lane *Lane) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", lane.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Lane")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(lane.Name))
	return
}

func (laneuse *LaneUse) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", laneuse.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "LaneUse")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(laneuse.Name))
	return
}

func (milestone *Milestone) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", milestone.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Milestone")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(milestone.Name))
	return
}

// insertion point for unstaging
func (arrow *Arrow) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", arrow.GongGetReferenceIdentifier(stage))
	return
}

func (bar *Bar) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", bar.GongGetReferenceIdentifier(stage))
	return
}

func (gantt *Gantt) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", gantt.GongGetReferenceIdentifier(stage))
	return
}

func (group *Group) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", group.GongGetReferenceIdentifier(stage))
	return
}

func (lane *Lane) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", lane.GongGetReferenceIdentifier(stage))
	return
}

func (laneuse *LaneUse) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", laneuse.GongGetReferenceIdentifier(stage))
	return
}

func (milestone *Milestone) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", milestone.GongGetReferenceIdentifier(stage))
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
