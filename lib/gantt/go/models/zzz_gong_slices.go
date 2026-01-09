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
	newInstance := *arrow
	return &newInstance
}

func (bar *Bar) GongCopy() GongstructIF {
	newInstance := *bar
	return &newInstance
}

func (gantt *Gantt) GongCopy() GongstructIF {
	newInstance := *gantt
	return &newInstance
}

func (group *Group) GongCopy() GongstructIF {
	newInstance := *group
	return &newInstance
}

func (lane *Lane) GongCopy() GongstructIF {
	newInstance := *lane
	return &newInstance
}

func (laneuse *LaneUse) GongCopy() GongstructIF {
	newInstance := *laneuse
	return &newInstance
}

func (milestone *Milestone) GongCopy() GongstructIF {
	newInstance := *milestone
	return &newInstance
}

func (stage *Stage) ComputeDifference() {
	var lenNewInstances int
	var lenModifiedInstances int
	var lenDeletedInstances int

	var newInstancesStmt string
	_ = newInstancesStmt
	var fieldsEditStmt string
	_ = fieldsEditStmt
	var deletedInstancesStmt string
	_ = deletedInstancesStmt

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
			newInstancesStmt += arrow.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := arrow.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := arrow.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", arrow.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for arrow := range stage.Arrows_reference {
		if _, ok := stage.Arrows[arrow]; !ok {
			arrows_deletedInstances = append(arrows_deletedInstances, arrow)
			deletedInstancesStmt += arrow.GongMarshallUnstaging(stage)
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
			newInstancesStmt += bar.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := bar.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := bar.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", bar.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for bar := range stage.Bars_reference {
		if _, ok := stage.Bars[bar]; !ok {
			bars_deletedInstances = append(bars_deletedInstances, bar)
			deletedInstancesStmt += bar.GongMarshallUnstaging(stage)
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
			newInstancesStmt += gantt.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := gantt.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := gantt.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", gantt.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for gantt := range stage.Gantts_reference {
		if _, ok := stage.Gantts[gantt]; !ok {
			gantts_deletedInstances = append(gantts_deletedInstances, gantt)
			deletedInstancesStmt += gantt.GongMarshallUnstaging(stage)
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
			newInstancesStmt += group.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := group.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := group.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", group.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for group := range stage.Groups_reference {
		if _, ok := stage.Groups[group]; !ok {
			groups_deletedInstances = append(groups_deletedInstances, group)
			deletedInstancesStmt += group.GongMarshallUnstaging(stage)
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
			newInstancesStmt += lane.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := lane.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := lane.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", lane.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for lane := range stage.Lanes_reference {
		if _, ok := stage.Lanes[lane]; !ok {
			lanes_deletedInstances = append(lanes_deletedInstances, lane)
			deletedInstancesStmt += lane.GongMarshallUnstaging(stage)
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
			newInstancesStmt += laneuse.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := laneuse.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := laneuse.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", laneuse.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for laneuse := range stage.LaneUses_reference {
		if _, ok := stage.LaneUses[laneuse]; !ok {
			laneuses_deletedInstances = append(laneuses_deletedInstances, laneuse)
			deletedInstancesStmt += laneuse.GongMarshallUnstaging(stage)
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
			newInstancesStmt += milestone.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := milestone.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := milestone.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", milestone.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for milestone := range stage.Milestones_reference {
		if _, ok := stage.Milestones[milestone]; !ok {
			milestones_deletedInstances = append(milestones_deletedInstances, milestone)
			deletedInstancesStmt += milestone.GongMarshallUnstaging(stage)
		}
	}

	lenNewInstances += len(milestones_newInstances)
	lenDeletedInstances += len(milestones_deletedInstances)

	if lenNewInstances > 0 || lenDeletedInstances > 0 || lenModifiedInstances > 0 {
		notif := newInstancesStmt + fieldsEditStmt + deletedInstancesStmt
		notif += fmt.Sprintf("\n\t// %s", time.Now().Format(time.RFC3339Nano))
		notif += "\n\tstage.Commit()"
		if stage.GetProbeIF() != nil {
			stage.GetProbeIF().AddNotification(
				time.Now(),
				notif,
			)
			stage.GetProbeIF().CommitNotificationTable()
		}
	}
}

// ComputeReference will creates a deep copy of each of the staged elements
func (stage *Stage) ComputeReference() {

	// insertion point per named struct
	stage.Arrows_reference = make(map[*Arrow]*Arrow)
	stage.Arrows_referenceOrder = make(map[*Arrow]uint) // diff Unstage needs the reference order
	for instance := range stage.Arrows {
		stage.Arrows_reference[instance] = instance.GongCopy().(*Arrow)
		stage.Arrows_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.Bars_reference = make(map[*Bar]*Bar)
	stage.Bars_referenceOrder = make(map[*Bar]uint) // diff Unstage needs the reference order
	for instance := range stage.Bars {
		stage.Bars_reference[instance] = instance.GongCopy().(*Bar)
		stage.Bars_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.Gantts_reference = make(map[*Gantt]*Gantt)
	stage.Gantts_referenceOrder = make(map[*Gantt]uint) // diff Unstage needs the reference order
	for instance := range stage.Gantts {
		stage.Gantts_reference[instance] = instance.GongCopy().(*Gantt)
		stage.Gantts_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.Groups_reference = make(map[*Group]*Group)
	stage.Groups_referenceOrder = make(map[*Group]uint) // diff Unstage needs the reference order
	for instance := range stage.Groups {
		stage.Groups_reference[instance] = instance.GongCopy().(*Group)
		stage.Groups_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.Lanes_reference = make(map[*Lane]*Lane)
	stage.Lanes_referenceOrder = make(map[*Lane]uint) // diff Unstage needs the reference order
	for instance := range stage.Lanes {
		stage.Lanes_reference[instance] = instance.GongCopy().(*Lane)
		stage.Lanes_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.LaneUses_reference = make(map[*LaneUse]*LaneUse)
	stage.LaneUses_referenceOrder = make(map[*LaneUse]uint) // diff Unstage needs the reference order
	for instance := range stage.LaneUses {
		stage.LaneUses_reference[instance] = instance.GongCopy().(*LaneUse)
		stage.LaneUses_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.Milestones_reference = make(map[*Milestone]*Milestone)
	stage.Milestones_referenceOrder = make(map[*Milestone]uint) // diff Unstage needs the reference order
	for instance := range stage.Milestones {
		stage.Milestones_reference[instance] = instance.GongCopy().(*Milestone)
		stage.Milestones_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

}

// GongGetOrder returns the order of the instance in the staging area
// This order is set at staging time, and reflects the order of creation of the instances
// in the staging area
// It is used when rendering slices of GongstructIF to keep a deterministic order
// which is important for frontends such as web frontends
// to avoid unnecessary re-renderings
// insertion point per named struct
func (arrow *Arrow) GongGetOrder(stage *Stage) uint {
	return stage.ArrowMap_Staged_Order[arrow]
}

func (arrow *Arrow) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Arrows_referenceOrder[arrow]
}

func (bar *Bar) GongGetOrder(stage *Stage) uint {
	return stage.BarMap_Staged_Order[bar]
}

func (bar *Bar) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Bars_referenceOrder[bar]
}

func (gantt *Gantt) GongGetOrder(stage *Stage) uint {
	return stage.GanttMap_Staged_Order[gantt]
}

func (gantt *Gantt) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Gantts_referenceOrder[gantt]
}

func (group *Group) GongGetOrder(stage *Stage) uint {
	return stage.GroupMap_Staged_Order[group]
}

func (group *Group) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Groups_referenceOrder[group]
}

func (lane *Lane) GongGetOrder(stage *Stage) uint {
	return stage.LaneMap_Staged_Order[lane]
}

func (lane *Lane) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Lanes_referenceOrder[lane]
}

func (laneuse *LaneUse) GongGetOrder(stage *Stage) uint {
	return stage.LaneUseMap_Staged_Order[laneuse]
}

func (laneuse *LaneUse) GongGetReferenceOrder(stage *Stage) uint {
	return stage.LaneUses_referenceOrder[laneuse]
}

func (milestone *Milestone) GongGetOrder(stage *Stage) uint {
	return stage.MilestoneMap_Staged_Order[milestone]
}

func (milestone *Milestone) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Milestones_referenceOrder[milestone]
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
	return fmt.Sprintf("__%s__%08d_", arrow.GongGetGongstructName(), arrow.GongGetReferenceOrder(stage))
}

func (bar *Bar) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", bar.GongGetGongstructName(), bar.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (bar *Bar) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", bar.GongGetGongstructName(), bar.GongGetReferenceOrder(stage))
}

func (gantt *Gantt) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", gantt.GongGetGongstructName(), gantt.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (gantt *Gantt) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", gantt.GongGetGongstructName(), gantt.GongGetReferenceOrder(stage))
}

func (group *Group) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", group.GongGetGongstructName(), group.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (group *Group) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", group.GongGetGongstructName(), group.GongGetReferenceOrder(stage))
}

func (lane *Lane) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", lane.GongGetGongstructName(), lane.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (lane *Lane) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", lane.GongGetGongstructName(), lane.GongGetReferenceOrder(stage))
}

func (laneuse *LaneUse) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", laneuse.GongGetGongstructName(), laneuse.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (laneuse *LaneUse) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", laneuse.GongGetGongstructName(), laneuse.GongGetReferenceOrder(stage))
}

func (milestone *Milestone) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", milestone.GongGetGongstructName(), milestone.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (milestone *Milestone) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", milestone.GongGetGongstructName(), milestone.GongGetReferenceOrder(stage))
}

// MarshallIdentifier returns the code to instantiate the instance
// in a marshalling file
// insertion point per named struct
func (arrow *Arrow) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", arrow.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Arrow")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", arrow.Name)
	return
}
func (bar *Bar) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", bar.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Bar")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", bar.Name)
	return
}
func (gantt *Gantt) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", gantt.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Gantt")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", gantt.Name)
	return
}
func (group *Group) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", group.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Group")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", group.Name)
	return
}
func (lane *Lane) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", lane.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Lane")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", lane.Name)
	return
}
func (laneuse *LaneUse) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", laneuse.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "LaneUse")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", laneuse.Name)
	return
}
func (milestone *Milestone) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", milestone.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Milestone")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", milestone.Name)
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
