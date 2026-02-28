// generated code - do not edit
package models

import (
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
	arrow.CopyBasicFields(newInstance)
	return newInstance
}

func (bar *Bar) GongCopy() GongstructIF {
	newInstance := new(Bar)
	bar.CopyBasicFields(newInstance)
	return newInstance
}

func (gantt *Gantt) GongCopy() GongstructIF {
	newInstance := new(Gantt)
	gantt.CopyBasicFields(newInstance)
	return newInstance
}

func (group *Group) GongCopy() GongstructIF {
	newInstance := new(Group)
	group.CopyBasicFields(newInstance)
	return newInstance
}

func (lane *Lane) GongCopy() GongstructIF {
	newInstance := new(Lane)
	lane.CopyBasicFields(newInstance)
	return newInstance
}

func (laneuse *LaneUse) GongCopy() GongstructIF {
	newInstance := new(LaneUse)
	laneuse.CopyBasicFields(newInstance)
	return newInstance
}

func (milestone *Milestone) GongCopy() GongstructIF {
	newInstance := new(Milestone)
	milestone.CopyBasicFields(newInstance)
	return newInstance
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
	var arrows_newInstances []*Arrow
	var arrows_deletedInstances []*Arrow

	// parse all staged instances and check if they have a reference
	for arrow := range stage.Arrows {
		if ref, ok := stage.Arrows_reference[arrow]; !ok {
			arrows_newInstances = append(arrows_newInstances, arrow)
			newInstancesSlice = append(newInstancesSlice, arrow.GongMarshallIdentifier(stage))
			if stage.Arrows_referenceOrder == nil {
				stage.Arrows_referenceOrder = make(map[*Arrow]uint)
			}
			stage.Arrows_referenceOrder[arrow] = stage.Arrow_stagedOrder[arrow]
			newInstancesReverseSlice = append(newInstancesReverseSlice, arrow.GongMarshallUnstaging(stage))
			// delete(stage.Arrows_referenceOrder, arrow)
			fieldInitializers, pointersInitializations := arrow.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Arrow_stagedOrder[ref] = stage.Arrow_stagedOrder[arrow]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := arrow.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, arrow)
			// delete(stage.Arrow_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", arrow.GetName())
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
	for _, ref := range stage.Arrows_reference {
		instance := stage.Arrows_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Arrows[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			arrows_deletedInstances = append(arrows_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(arrows_newInstances)
	lenDeletedInstances += len(arrows_deletedInstances)
	var bars_newInstances []*Bar
	var bars_deletedInstances []*Bar

	// parse all staged instances and check if they have a reference
	for bar := range stage.Bars {
		if ref, ok := stage.Bars_reference[bar]; !ok {
			bars_newInstances = append(bars_newInstances, bar)
			newInstancesSlice = append(newInstancesSlice, bar.GongMarshallIdentifier(stage))
			if stage.Bars_referenceOrder == nil {
				stage.Bars_referenceOrder = make(map[*Bar]uint)
			}
			stage.Bars_referenceOrder[bar] = stage.Bar_stagedOrder[bar]
			newInstancesReverseSlice = append(newInstancesReverseSlice, bar.GongMarshallUnstaging(stage))
			// delete(stage.Bars_referenceOrder, bar)
			fieldInitializers, pointersInitializations := bar.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Bar_stagedOrder[ref] = stage.Bar_stagedOrder[bar]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := bar.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, bar)
			// delete(stage.Bar_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", bar.GetName())
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
	for _, ref := range stage.Bars_reference {
		instance := stage.Bars_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Bars[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			bars_deletedInstances = append(bars_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(bars_newInstances)
	lenDeletedInstances += len(bars_deletedInstances)
	var gantts_newInstances []*Gantt
	var gantts_deletedInstances []*Gantt

	// parse all staged instances and check if they have a reference
	for gantt := range stage.Gantts {
		if ref, ok := stage.Gantts_reference[gantt]; !ok {
			gantts_newInstances = append(gantts_newInstances, gantt)
			newInstancesSlice = append(newInstancesSlice, gantt.GongMarshallIdentifier(stage))
			if stage.Gantts_referenceOrder == nil {
				stage.Gantts_referenceOrder = make(map[*Gantt]uint)
			}
			stage.Gantts_referenceOrder[gantt] = stage.Gantt_stagedOrder[gantt]
			newInstancesReverseSlice = append(newInstancesReverseSlice, gantt.GongMarshallUnstaging(stage))
			// delete(stage.Gantts_referenceOrder, gantt)
			fieldInitializers, pointersInitializations := gantt.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Gantt_stagedOrder[ref] = stage.Gantt_stagedOrder[gantt]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := gantt.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, gantt)
			// delete(stage.Gantt_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", gantt.GetName())
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
	for _, ref := range stage.Gantts_reference {
		instance := stage.Gantts_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Gantts[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			gantts_deletedInstances = append(gantts_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(gantts_newInstances)
	lenDeletedInstances += len(gantts_deletedInstances)
	var groups_newInstances []*Group
	var groups_deletedInstances []*Group

	// parse all staged instances and check if they have a reference
	for group := range stage.Groups {
		if ref, ok := stage.Groups_reference[group]; !ok {
			groups_newInstances = append(groups_newInstances, group)
			newInstancesSlice = append(newInstancesSlice, group.GongMarshallIdentifier(stage))
			if stage.Groups_referenceOrder == nil {
				stage.Groups_referenceOrder = make(map[*Group]uint)
			}
			stage.Groups_referenceOrder[group] = stage.Group_stagedOrder[group]
			newInstancesReverseSlice = append(newInstancesReverseSlice, group.GongMarshallUnstaging(stage))
			// delete(stage.Groups_referenceOrder, group)
			fieldInitializers, pointersInitializations := group.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Group_stagedOrder[ref] = stage.Group_stagedOrder[group]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := group.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, group)
			// delete(stage.Group_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", group.GetName())
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
	for _, ref := range stage.Groups_reference {
		instance := stage.Groups_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Groups[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			groups_deletedInstances = append(groups_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(groups_newInstances)
	lenDeletedInstances += len(groups_deletedInstances)
	var lanes_newInstances []*Lane
	var lanes_deletedInstances []*Lane

	// parse all staged instances and check if they have a reference
	for lane := range stage.Lanes {
		if ref, ok := stage.Lanes_reference[lane]; !ok {
			lanes_newInstances = append(lanes_newInstances, lane)
			newInstancesSlice = append(newInstancesSlice, lane.GongMarshallIdentifier(stage))
			if stage.Lanes_referenceOrder == nil {
				stage.Lanes_referenceOrder = make(map[*Lane]uint)
			}
			stage.Lanes_referenceOrder[lane] = stage.Lane_stagedOrder[lane]
			newInstancesReverseSlice = append(newInstancesReverseSlice, lane.GongMarshallUnstaging(stage))
			// delete(stage.Lanes_referenceOrder, lane)
			fieldInitializers, pointersInitializations := lane.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Lane_stagedOrder[ref] = stage.Lane_stagedOrder[lane]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := lane.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, lane)
			// delete(stage.Lane_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", lane.GetName())
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
	for _, ref := range stage.Lanes_reference {
		instance := stage.Lanes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Lanes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			lanes_deletedInstances = append(lanes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(lanes_newInstances)
	lenDeletedInstances += len(lanes_deletedInstances)
	var laneuses_newInstances []*LaneUse
	var laneuses_deletedInstances []*LaneUse

	// parse all staged instances and check if they have a reference
	for laneuse := range stage.LaneUses {
		if ref, ok := stage.LaneUses_reference[laneuse]; !ok {
			laneuses_newInstances = append(laneuses_newInstances, laneuse)
			newInstancesSlice = append(newInstancesSlice, laneuse.GongMarshallIdentifier(stage))
			if stage.LaneUses_referenceOrder == nil {
				stage.LaneUses_referenceOrder = make(map[*LaneUse]uint)
			}
			stage.LaneUses_referenceOrder[laneuse] = stage.LaneUse_stagedOrder[laneuse]
			newInstancesReverseSlice = append(newInstancesReverseSlice, laneuse.GongMarshallUnstaging(stage))
			// delete(stage.LaneUses_referenceOrder, laneuse)
			fieldInitializers, pointersInitializations := laneuse.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.LaneUse_stagedOrder[ref] = stage.LaneUse_stagedOrder[laneuse]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := laneuse.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, laneuse)
			// delete(stage.LaneUse_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", laneuse.GetName())
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
	for _, ref := range stage.LaneUses_reference {
		instance := stage.LaneUses_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.LaneUses[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			laneuses_deletedInstances = append(laneuses_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(laneuses_newInstances)
	lenDeletedInstances += len(laneuses_deletedInstances)
	var milestones_newInstances []*Milestone
	var milestones_deletedInstances []*Milestone

	// parse all staged instances and check if they have a reference
	for milestone := range stage.Milestones {
		if ref, ok := stage.Milestones_reference[milestone]; !ok {
			milestones_newInstances = append(milestones_newInstances, milestone)
			newInstancesSlice = append(newInstancesSlice, milestone.GongMarshallIdentifier(stage))
			if stage.Milestones_referenceOrder == nil {
				stage.Milestones_referenceOrder = make(map[*Milestone]uint)
			}
			stage.Milestones_referenceOrder[milestone] = stage.Milestone_stagedOrder[milestone]
			newInstancesReverseSlice = append(newInstancesReverseSlice, milestone.GongMarshallUnstaging(stage))
			// delete(stage.Milestones_referenceOrder, milestone)
			fieldInitializers, pointersInitializations := milestone.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Milestone_stagedOrder[ref] = stage.Milestone_stagedOrder[milestone]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := milestone.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, milestone)
			// delete(stage.Milestone_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", milestone.GetName())
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
	for _, ref := range stage.Milestones_reference {
		instance := stage.Milestones_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Milestones[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			milestones_deletedInstances = append(milestones_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(milestones_newInstances)
	lenDeletedInstances += len(milestones_deletedInstances)

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
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(arrow.Name))
	return
}

func (bar *Bar) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", bar.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Bar")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(bar.Name))
	return
}

func (gantt *Gantt) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", gantt.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Gantt")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(gantt.Name))
	return
}

func (group *Group) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", group.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Group")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(group.Name))
	return
}

func (lane *Lane) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", lane.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Lane")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(lane.Name))
	return
}

func (laneuse *LaneUse) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", laneuse.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "LaneUse")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(laneuse.Name))
	return
}

func (milestone *Milestone) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", milestone.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Milestone")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(milestone.Name))
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

// end of template
