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
	res = __gong__appendInstances(res, stage.Arrows)

	res = __gong__appendInstances(res, stage.Bars)

	res = __gong__appendInstances(res, stage.Gantts)

	res = __gong__appendInstances(res, stage.Groups)

	res = __gong__appendInstances(res, stage.Lanes)

	res = __gong__appendInstances(res, stage.LaneUses)

	res = __gong__appendInstances(res, stage.Milestones)

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
func (arrow *Arrow) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, arrow)
}

func (bar *Bar) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, bar)
}

func (gantt *Gantt) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, gantt)
}

func (group *Group) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, group)
}

func (lane *Lane) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, lane)
}

func (laneuse *LaneUse) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, laneuse)
}

func (milestone *Milestone) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, milestone)
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
	__gong__computeReferencePass1(stage, stage.Arrows, &stage.Arrows_reference, &stage.Arrows_referenceOrder, &stage.Arrows_instance)

	__gong__computeReferencePass1(stage, stage.Bars, &stage.Bars_reference, &stage.Bars_referenceOrder, &stage.Bars_instance)

	__gong__computeReferencePass1(stage, stage.Gantts, &stage.Gantts_reference, &stage.Gantts_referenceOrder, &stage.Gantts_instance)

	__gong__computeReferencePass1(stage, stage.Groups, &stage.Groups_reference, &stage.Groups_referenceOrder, &stage.Groups_instance)

	__gong__computeReferencePass1(stage, stage.Lanes, &stage.Lanes_reference, &stage.Lanes_referenceOrder, &stage.Lanes_instance)

	__gong__computeReferencePass1(stage, stage.LaneUses, &stage.LaneUses_reference, &stage.LaneUses_referenceOrder, &stage.LaneUses_instance)

	__gong__computeReferencePass1(stage, stage.Milestones, &stage.Milestones_reference, &stage.Milestones_referenceOrder, &stage.Milestones_instance)

	// insertion point per named struct
	__gong__computeReferencePass2(stage.Arrows, stage.Arrows_reference, stage)

	__gong__computeReferencePass2(stage.Bars, stage.Bars_reference, stage)

	__gong__computeReferencePass2(stage.Gantts, stage.Gantts_reference, stage)

	__gong__computeReferencePass2(stage.Groups, stage.Groups_reference, stage)

	__gong__computeReferencePass2(stage.Lanes, stage.Lanes_reference, stage)

	__gong__computeReferencePass2(stage.LaneUses, stage.LaneUses_reference, stage)

	__gong__computeReferencePass2(stage.Milestones, stage.Milestones_reference, stage)

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
	return __gong__getOrder(stage.Arrow_stagedOrder, stage.Arrows_referenceOrder, arrow, "Arrow")
}

func (bar *Bar) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Bar_stagedOrder, stage.Bars_referenceOrder, bar, "Bar")
}

func (gantt *Gantt) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Gantt_stagedOrder, stage.Gantts_referenceOrder, gantt, "Gantt")
}

func (group *Group) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Group_stagedOrder, stage.Groups_referenceOrder, group, "Group")
}

func (lane *Lane) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Lane_stagedOrder, stage.Lanes_referenceOrder, lane, "Lane")
}

func (laneuse *LaneUse) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.LaneUse_stagedOrder, stage.LaneUses_referenceOrder, laneuse, "LaneUse")
}

func (milestone *Milestone) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Milestone_stagedOrder, stage.Milestones_referenceOrder, milestone, "Milestone")
}

// GongGetIdentifier returns a unique identifier of the instance in the staging area
// This identifier is composed of the Gongstruct name and the order of the instance
// in the staging area
// It is used to identify instances across sessions
// insertion point per named struct
func (arrow *Arrow) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(arrow, arrow.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (arrow *Arrow) GongGetReferenceIdentifier(stage *Stage) string {
	return arrow.GongGetIdentifier(stage)
}

func (bar *Bar) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(bar, bar.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (bar *Bar) GongGetReferenceIdentifier(stage *Stage) string {
	return bar.GongGetIdentifier(stage)
}

func (gantt *Gantt) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(gantt, gantt.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (gantt *Gantt) GongGetReferenceIdentifier(stage *Stage) string {
	return gantt.GongGetIdentifier(stage)
}

func (group *Group) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(group, group.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (group *Group) GongGetReferenceIdentifier(stage *Stage) string {
	return group.GongGetIdentifier(stage)
}

func (lane *Lane) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(lane, lane.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (lane *Lane) GongGetReferenceIdentifier(stage *Stage) string {
	return lane.GongGetIdentifier(stage)
}

func (laneuse *LaneUse) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(laneuse, laneuse.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (laneuse *LaneUse) GongGetReferenceIdentifier(stage *Stage) string {
	return laneuse.GongGetIdentifier(stage)
}

func (milestone *Milestone) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(milestone, milestone.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (milestone *Milestone) GongGetReferenceIdentifier(stage *Stage) string {
	return milestone.GongGetIdentifier(stage)
}

// MarshallIdentifier returns the code to instantiate the instance
// in a marshalling file
// insertion point per named struct
func (arrow *Arrow) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(arrow.GongGetIdentifier(stage), "Arrow", arrow.Name)
}

func (bar *Bar) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(bar.GongGetIdentifier(stage), "Bar", bar.Name)
}

func (gantt *Gantt) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(gantt.GongGetIdentifier(stage), "Gantt", gantt.Name)
}

func (group *Group) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(group.GongGetIdentifier(stage), "Group", group.Name)
}

func (lane *Lane) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(lane.GongGetIdentifier(stage), "Lane", lane.Name)
}

func (laneuse *LaneUse) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(laneuse.GongGetIdentifier(stage), "LaneUse", laneuse.Name)
}

func (milestone *Milestone) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(milestone.GongGetIdentifier(stage), "Milestone", milestone.Name)
}

// insertion point for unstaging
func (arrow *Arrow) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(arrow.GongGetReferenceIdentifier(stage))
}

func (bar *Bar) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(bar.GongGetReferenceIdentifier(stage))
}

func (gantt *Gantt) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(gantt.GongGetReferenceIdentifier(stage))
}

func (group *Group) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(group.GongGetReferenceIdentifier(stage))
}

func (lane *Lane) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(lane.GongGetReferenceIdentifier(stage))
}

func (laneuse *LaneUse) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(laneuse.GongGetReferenceIdentifier(stage))
}

func (milestone *Milestone) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(milestone.GongGetReferenceIdentifier(stage))
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
