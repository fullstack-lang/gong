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

	// insertion point per named struct
	var arrows_newInstances []*Arrow
	var arrows_deletedInstances []*Arrow

	// parse all staged instances and check if they have a reference
	for arrow := range stage.Arrows {
		if ref, ok := stage.Arrows_reference[arrow]; !ok {
			arrows_newInstances = append(arrows_newInstances, arrow)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of Arrow "+arrow.Name,
				)
			}
		} else {
			diffs := arrow.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of Arrow \""+arrow.Name+"\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for arrow := range stage.Arrows_reference {
		if _, ok := stage.Arrows[arrow]; !ok {
			arrows_deletedInstances = append(arrows_deletedInstances, arrow)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of Arrow "+arrow.Name,
				)
			}
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
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of Bar "+bar.Name,
				)
			}
		} else {
			diffs := bar.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of Bar \""+bar.Name+"\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for bar := range stage.Bars_reference {
		if _, ok := stage.Bars[bar]; !ok {
			bars_deletedInstances = append(bars_deletedInstances, bar)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of Bar "+bar.Name,
				)
			}
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
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of Gantt "+gantt.Name,
				)
			}
		} else {
			diffs := gantt.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of Gantt \""+gantt.Name+"\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for gantt := range stage.Gantts_reference {
		if _, ok := stage.Gantts[gantt]; !ok {
			gantts_deletedInstances = append(gantts_deletedInstances, gantt)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of Gantt "+gantt.Name,
				)
			}
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
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of Group "+group.Name,
				)
			}
		} else {
			diffs := group.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of Group \""+group.Name+"\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for group := range stage.Groups_reference {
		if _, ok := stage.Groups[group]; !ok {
			groups_deletedInstances = append(groups_deletedInstances, group)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of Group "+group.Name,
				)
			}
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
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of Lane "+lane.Name,
				)
			}
		} else {
			diffs := lane.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of Lane \""+lane.Name+"\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for lane := range stage.Lanes_reference {
		if _, ok := stage.Lanes[lane]; !ok {
			lanes_deletedInstances = append(lanes_deletedInstances, lane)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of Lane "+lane.Name,
				)
			}
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
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of LaneUse "+laneuse.Name,
				)
			}
		} else {
			diffs := laneuse.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of LaneUse \""+laneuse.Name+"\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for laneuse := range stage.LaneUses_reference {
		if _, ok := stage.LaneUses[laneuse]; !ok {
			laneuses_deletedInstances = append(laneuses_deletedInstances, laneuse)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of LaneUse "+laneuse.Name,
				)
			}
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
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of Milestone "+milestone.Name,
				)
			}
		} else {
			diffs := milestone.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of Milestone \""+milestone.Name+"\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for milestone := range stage.Milestones_reference {
		if _, ok := stage.Milestones[milestone]; !ok {
			milestones_deletedInstances = append(milestones_deletedInstances, milestone)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of Milestone "+milestone.Name,
				)
			}
		}
	}

	lenNewInstances += len(milestones_newInstances)
	lenDeletedInstances += len(milestones_deletedInstances)

	if lenNewInstances > 0 || lenDeletedInstances > 0 || lenModifiedInstances > 0 {
		// if stage.GetProbeIF() != nil {
		// 	stage.GetProbeIF().CommitNotificationTable()
		// }
	}
}

// ComputeReference will creates a deep copy of each of the staged elements
func (stage *Stage) ComputeReference() {

	// insertion point per named struct
	stage.Arrows_reference = make(map[*Arrow]*Arrow)
	for instance := range stage.Arrows {
		stage.Arrows_reference[instance] = instance.GongCopy().(*Arrow)
	}

	stage.Bars_reference = make(map[*Bar]*Bar)
	for instance := range stage.Bars {
		stage.Bars_reference[instance] = instance.GongCopy().(*Bar)
	}

	stage.Gantts_reference = make(map[*Gantt]*Gantt)
	for instance := range stage.Gantts {
		stage.Gantts_reference[instance] = instance.GongCopy().(*Gantt)
	}

	stage.Groups_reference = make(map[*Group]*Group)
	for instance := range stage.Groups {
		stage.Groups_reference[instance] = instance.GongCopy().(*Group)
	}

	stage.Lanes_reference = make(map[*Lane]*Lane)
	for instance := range stage.Lanes {
		stage.Lanes_reference[instance] = instance.GongCopy().(*Lane)
	}

	stage.LaneUses_reference = make(map[*LaneUse]*LaneUse)
	for instance := range stage.LaneUses {
		stage.LaneUses_reference[instance] = instance.GongCopy().(*LaneUse)
	}

	stage.Milestones_reference = make(map[*Milestone]*Milestone)
	for instance := range stage.Milestones {
		stage.Milestones_reference[instance] = instance.GongCopy().(*Milestone)
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

func (bar *Bar) GongGetOrder(stage *Stage) uint {
	return stage.BarMap_Staged_Order[bar]
}

func (gantt *Gantt) GongGetOrder(stage *Stage) uint {
	return stage.GanttMap_Staged_Order[gantt]
}

func (group *Group) GongGetOrder(stage *Stage) uint {
	return stage.GroupMap_Staged_Order[group]
}

func (lane *Lane) GongGetOrder(stage *Stage) uint {
	return stage.LaneMap_Staged_Order[lane]
}

func (laneuse *LaneUse) GongGetOrder(stage *Stage) uint {
	return stage.LaneUseMap_Staged_Order[laneuse]
}

func (milestone *Milestone) GongGetOrder(stage *Stage) uint {
	return stage.MilestoneMap_Staged_Order[milestone]
}


// GongGetIdentifier returns a unique identifier of the instance in the staging area
// This identifier is composed of the Gongstruct name and the order of the instance
// in the staging area
// It is used to identify instances across sessions
// insertion point per named struct
func (arrow *Arrow) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", arrow.GongGetGongstructName(), arrow.GongGetOrder(stage))
}

func (bar *Bar) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", bar.GongGetGongstructName(), bar.GongGetOrder(stage))
}

func (gantt *Gantt) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", gantt.GongGetGongstructName(), gantt.GongGetOrder(stage))
}

func (group *Group) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", group.GongGetGongstructName(), group.GongGetOrder(stage))
}

func (lane *Lane) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", lane.GongGetGongstructName(), lane.GongGetOrder(stage))
}

func (laneuse *LaneUse) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", laneuse.GongGetGongstructName(), laneuse.GongGetOrder(stage))
}

func (milestone *Milestone) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", milestone.GongGetGongstructName(), milestone.GongGetOrder(stage))
}

// MarshallIdentifier returns the code to instantiate the instance
// in a marshalling file
// insertion point per named struct
func (arrow *Arrow) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = IdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", arrow.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Arrow")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", arrow.Name)
	return
}
func (bar *Bar) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = IdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", bar.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Bar")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", bar.Name)
	return
}
func (gantt *Gantt) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = IdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", gantt.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Gantt")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", gantt.Name)
	return
}
func (group *Group) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = IdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", group.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Group")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", group.Name)
	return
}
func (lane *Lane) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = IdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", lane.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Lane")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", lane.Name)
	return
}
func (laneuse *LaneUse) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = IdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", laneuse.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "LaneUse")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", laneuse.Name)
	return
}
func (milestone *Milestone) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = IdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", milestone.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Milestone")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", milestone.Name)
	return
}
