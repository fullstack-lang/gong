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
func (arrow *Arrow) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Arrows[arrow]

	return
}

func (stage *Stage) IsStagedArrow(arrow *Arrow) (ok bool) {

	return arrow.GongIsStaged(stage)
}

func (bar *Bar) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Bars[bar]

	return
}

func (stage *Stage) IsStagedBar(bar *Bar) (ok bool) {

	return bar.GongIsStaged(stage)
}

func (gantt *Gantt) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Gantts[gantt]

	return
}

func (stage *Stage) IsStagedGantt(gantt *Gantt) (ok bool) {

	return gantt.GongIsStaged(stage)
}

func (group *Group) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Groups[group]

	return
}

func (stage *Stage) IsStagedGroup(group *Group) (ok bool) {

	return group.GongIsStaged(stage)
}

func (lane *Lane) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Lanes[lane]

	return
}

func (stage *Stage) IsStagedLane(lane *Lane) (ok bool) {

	return lane.GongIsStaged(stage)
}

func (laneuse *LaneUse) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.LaneUses[laneuse]

	return
}

func (stage *Stage) IsStagedLaneUse(laneuse *LaneUse) (ok bool) {

	return laneuse.GongIsStaged(stage)
}

func (milestone *Milestone) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Milestones[milestone]

	return
}

func (stage *Stage) IsStagedMilestone(milestone *Milestone) (ok bool) {

	return milestone.GongIsStaged(stage)
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
func (arrow *Arrow) GongStageBranch(stage *Stage) {
	stage.StageBranchArrow(arrow)
}

func (stage *Stage) StageBranchArrow(arrow *Arrow) {

	// check if instance is already staged
	if stage.IsStaged(arrow) {
		return
	}

	arrow.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if arrow.From != nil {
		stage.StageBranch(arrow.From)
	}
	if arrow.To != nil {
		stage.StageBranch(arrow.To)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (bar *Bar) GongStageBranch(stage *Stage) {
	stage.StageBranchBar(bar)
}

func (stage *Stage) StageBranchBar(bar *Bar) {

	// check if instance is already staged
	if stage.IsStaged(bar) {
		return
	}

	bar.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (gantt *Gantt) GongStageBranch(stage *Stage) {
	stage.StageBranchGantt(gantt)
}

func (stage *Stage) StageBranchGantt(gantt *Gantt) {

	// check if instance is already staged
	if stage.IsStaged(gantt) {
		return
	}

	gantt.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _lane := range gantt.Lanes {
		stage.StageBranch(_lane)
	}
	for _, _milestone := range gantt.Milestones {
		stage.StageBranch(_milestone)
	}
	for _, _group := range gantt.Groups {
		stage.StageBranch(_group)
	}
	for _, _arrow := range gantt.Arrows {
		stage.StageBranch(_arrow)
	}

}

func (group *Group) GongStageBranch(stage *Stage) {
	stage.StageBranchGroup(group)
}

func (stage *Stage) StageBranchGroup(group *Group) {

	// check if instance is already staged
	if stage.IsStaged(group) {
		return
	}

	group.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _lane := range group.GroupLanes {
		stage.StageBranch(_lane)
	}

}

func (lane *Lane) GongStageBranch(stage *Stage) {
	stage.StageBranchLane(lane)
}

func (stage *Stage) StageBranchLane(lane *Lane) {

	// check if instance is already staged
	if stage.IsStaged(lane) {
		return
	}

	lane.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _bar := range lane.Bars {
		stage.StageBranch(_bar)
	}

}

func (laneuse *LaneUse) GongStageBranch(stage *Stage) {
	stage.StageBranchLaneUse(laneuse)
}

func (stage *Stage) StageBranchLaneUse(laneuse *LaneUse) {

	// check if instance is already staged
	if stage.IsStaged(laneuse) {
		return
	}

	laneuse.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if laneuse.Lane != nil {
		stage.StageBranch(laneuse.Lane)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (milestone *Milestone) GongStageBranch(stage *Stage) {
	stage.StageBranchMilestone(milestone)
}

func (stage *Stage) StageBranchMilestone(milestone *Milestone) {

	// check if instance is already staged
	if stage.IsStaged(milestone) {
		return
	}

	milestone.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _lane := range milestone.LanesToDisplay {
		stage.StageBranch(_lane)
	}

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
	case *Arrow:
		toT := GongCopyBranchArrow(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Bar:
		toT := GongCopyBranchBar(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Gantt:
		toT := GongCopyBranchGantt(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Group:
		toT := GongCopyBranchGroup(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Lane:
		toT := GongCopyBranchLane(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *LaneUse:
		toT := GongCopyBranchLaneUse(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Milestone:
		toT := GongCopyBranchMilestone(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
func GongCopyBranchArrow(mapOrigCopy map[any]any, arrowFrom *Arrow) (arrowTo *Arrow) {

	// arrowFrom has already been copied
	if _arrowTo, ok := mapOrigCopy[arrowFrom]; ok {
		arrowTo = _arrowTo.(*Arrow)
		return
	}

	arrowTo = new(Arrow)
	mapOrigCopy[arrowFrom] = arrowTo
	arrowFrom.GongCopyBasicFields(arrowTo)

	//insertion point for the staging of instances referenced by pointers
	if arrowFrom.From != nil {
		arrowTo.From = GongCopyBranchBar(mapOrigCopy, arrowFrom.From)
	}
	if arrowFrom.To != nil {
		arrowTo.To = GongCopyBranchBar(mapOrigCopy, arrowFrom.To)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchBar(mapOrigCopy map[any]any, barFrom *Bar) (barTo *Bar) {

	// barFrom has already been copied
	if _barTo, ok := mapOrigCopy[barFrom]; ok {
		barTo = _barTo.(*Bar)
		return
	}

	barTo = new(Bar)
	mapOrigCopy[barFrom] = barTo
	barFrom.GongCopyBasicFields(barTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchGantt(mapOrigCopy map[any]any, ganttFrom *Gantt) (ganttTo *Gantt) {

	// ganttFrom has already been copied
	if _ganttTo, ok := mapOrigCopy[ganttFrom]; ok {
		ganttTo = _ganttTo.(*Gantt)
		return
	}

	ganttTo = new(Gantt)
	mapOrigCopy[ganttFrom] = ganttTo
	ganttFrom.GongCopyBasicFields(ganttTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _lane := range ganttFrom.Lanes {
		ganttTo.Lanes = append(ganttTo.Lanes, GongCopyBranchLane(mapOrigCopy, _lane))
	}
	for _, _milestone := range ganttFrom.Milestones {
		ganttTo.Milestones = append(ganttTo.Milestones, GongCopyBranchMilestone(mapOrigCopy, _milestone))
	}
	for _, _group := range ganttFrom.Groups {
		ganttTo.Groups = append(ganttTo.Groups, GongCopyBranchGroup(mapOrigCopy, _group))
	}
	for _, _arrow := range ganttFrom.Arrows {
		ganttTo.Arrows = append(ganttTo.Arrows, GongCopyBranchArrow(mapOrigCopy, _arrow))
	}

	return
}

func GongCopyBranchGroup(mapOrigCopy map[any]any, groupFrom *Group) (groupTo *Group) {

	// groupFrom has already been copied
	if _groupTo, ok := mapOrigCopy[groupFrom]; ok {
		groupTo = _groupTo.(*Group)
		return
	}

	groupTo = new(Group)
	mapOrigCopy[groupFrom] = groupTo
	groupFrom.GongCopyBasicFields(groupTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _lane := range groupFrom.GroupLanes {
		groupTo.GroupLanes = append(groupTo.GroupLanes, GongCopyBranchLane(mapOrigCopy, _lane))
	}

	return
}

func GongCopyBranchLane(mapOrigCopy map[any]any, laneFrom *Lane) (laneTo *Lane) {

	// laneFrom has already been copied
	if _laneTo, ok := mapOrigCopy[laneFrom]; ok {
		laneTo = _laneTo.(*Lane)
		return
	}

	laneTo = new(Lane)
	mapOrigCopy[laneFrom] = laneTo
	laneFrom.GongCopyBasicFields(laneTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _bar := range laneFrom.Bars {
		laneTo.Bars = append(laneTo.Bars, GongCopyBranchBar(mapOrigCopy, _bar))
	}

	return
}

func GongCopyBranchLaneUse(mapOrigCopy map[any]any, laneuseFrom *LaneUse) (laneuseTo *LaneUse) {

	// laneuseFrom has already been copied
	if _laneuseTo, ok := mapOrigCopy[laneuseFrom]; ok {
		laneuseTo = _laneuseTo.(*LaneUse)
		return
	}

	laneuseTo = new(LaneUse)
	mapOrigCopy[laneuseFrom] = laneuseTo
	laneuseFrom.GongCopyBasicFields(laneuseTo)

	//insertion point for the staging of instances referenced by pointers
	if laneuseFrom.Lane != nil {
		laneuseTo.Lane = GongCopyBranchLane(mapOrigCopy, laneuseFrom.Lane)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchMilestone(mapOrigCopy map[any]any, milestoneFrom *Milestone) (milestoneTo *Milestone) {

	// milestoneFrom has already been copied
	if _milestoneTo, ok := mapOrigCopy[milestoneFrom]; ok {
		milestoneTo = _milestoneTo.(*Milestone)
		return
	}

	milestoneTo = new(Milestone)
	mapOrigCopy[milestoneFrom] = milestoneTo
	milestoneFrom.GongCopyBasicFields(milestoneTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _lane := range milestoneFrom.LanesToDisplay {
		milestoneTo.LanesToDisplay = append(milestoneTo.LanesToDisplay, GongCopyBranchLane(mapOrigCopy, _lane))
	}

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
func (arrow *Arrow) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchArrow(arrow)
}

func (stage *Stage) UnstageBranchArrow(arrow *Arrow) {

	// check if instance is already staged
	if !stage.IsStaged(arrow) {
		return
	}

	arrow.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if arrow.From != nil {
		stage.UnstageBranch(arrow.From)
	}
	if arrow.To != nil {
		stage.UnstageBranch(arrow.To)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (bar *Bar) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchBar(bar)
}

func (stage *Stage) UnstageBranchBar(bar *Bar) {

	// check if instance is already staged
	if !stage.IsStaged(bar) {
		return
	}

	bar.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (gantt *Gantt) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchGantt(gantt)
}

func (stage *Stage) UnstageBranchGantt(gantt *Gantt) {

	// check if instance is already staged
	if !stage.IsStaged(gantt) {
		return
	}

	gantt.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _lane := range gantt.Lanes {
		stage.UnstageBranch(_lane)
	}
	for _, _milestone := range gantt.Milestones {
		stage.UnstageBranch(_milestone)
	}
	for _, _group := range gantt.Groups {
		stage.UnstageBranch(_group)
	}
	for _, _arrow := range gantt.Arrows {
		stage.UnstageBranch(_arrow)
	}

}

func (group *Group) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchGroup(group)
}

func (stage *Stage) UnstageBranchGroup(group *Group) {

	// check if instance is already staged
	if !stage.IsStaged(group) {
		return
	}

	group.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _lane := range group.GroupLanes {
		stage.UnstageBranch(_lane)
	}

}

func (lane *Lane) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchLane(lane)
}

func (stage *Stage) UnstageBranchLane(lane *Lane) {

	// check if instance is already staged
	if !stage.IsStaged(lane) {
		return
	}

	lane.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _bar := range lane.Bars {
		stage.UnstageBranch(_bar)
	}

}

func (laneuse *LaneUse) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchLaneUse(laneuse)
}

func (stage *Stage) UnstageBranchLaneUse(laneuse *LaneUse) {

	// check if instance is already staged
	if !stage.IsStaged(laneuse) {
		return
	}

	laneuse.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if laneuse.Lane != nil {
		stage.UnstageBranch(laneuse.Lane)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (milestone *Milestone) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchMilestone(milestone)
}

func (stage *Stage) UnstageBranchMilestone(milestone *Milestone) {

	// check if instance is already staged
	if !stage.IsStaged(milestone) {
		return
	}

	milestone.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _lane := range milestone.LanesToDisplay {
		stage.UnstageBranch(_lane)
	}

}

// insertion point for pointer reconstruction from references
func (reference *Arrow) GongReconstructPointersFromReferences(stage *Stage, instance *Arrow) {
	// insertion point for pointers field
	if instance.From != nil {
		reference.From = stage.Bars_reference[instance.From]
	}
	if instance.To != nil {
		reference.To = stage.Bars_reference[instance.To]
	}
	// insertion point for slice of pointers field
}

func (reference *Bar) GongReconstructPointersFromReferences(stage *Stage, instance *Bar) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Gantt) GongReconstructPointersFromReferences(stage *Stage, instance *Gantt) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.Lanes = reference.Lanes[:0]
	for _, _b := range instance.Lanes {
		reference.Lanes = append(reference.Lanes, stage.Lanes_reference[_b])
	}
	reference.Milestones = reference.Milestones[:0]
	for _, _b := range instance.Milestones {
		reference.Milestones = append(reference.Milestones, stage.Milestones_reference[_b])
	}
	reference.Groups = reference.Groups[:0]
	for _, _b := range instance.Groups {
		reference.Groups = append(reference.Groups, stage.Groups_reference[_b])
	}
	reference.Arrows = reference.Arrows[:0]
	for _, _b := range instance.Arrows {
		reference.Arrows = append(reference.Arrows, stage.Arrows_reference[_b])
	}
}

func (reference *Group) GongReconstructPointersFromReferences(stage *Stage, instance *Group) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.GroupLanes = reference.GroupLanes[:0]
	for _, _b := range instance.GroupLanes {
		reference.GroupLanes = append(reference.GroupLanes, stage.Lanes_reference[_b])
	}
}

func (reference *Lane) GongReconstructPointersFromReferences(stage *Stage, instance *Lane) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.Bars = reference.Bars[:0]
	for _, _b := range instance.Bars {
		reference.Bars = append(reference.Bars, stage.Bars_reference[_b])
	}
}

func (reference *LaneUse) GongReconstructPointersFromReferences(stage *Stage, instance *LaneUse) {
	// insertion point for pointers field
	if instance.Lane != nil {
		reference.Lane = stage.Lanes_reference[instance.Lane]
	}
	// insertion point for slice of pointers field
}

func (reference *Milestone) GongReconstructPointersFromReferences(stage *Stage, instance *Milestone) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.LanesToDisplay = reference.LanesToDisplay[:0]
	for _, _b := range instance.LanesToDisplay {
		reference.LanesToDisplay = append(reference.LanesToDisplay, stage.Lanes_reference[_b])
	}
}

// insertion point for pointer reconstruction from instances
func (reference *Arrow) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.From; _reference != nil {
		reference.From = nil
		if _instance, ok := stage.Bars_instance[_reference]; ok {
			reference.From = _instance
		}
	}
	if _reference := reference.To; _reference != nil {
		reference.To = nil
		if _instance, ok := stage.Bars_instance[_reference]; ok {
			reference.To = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *Bar) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Gantt) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _Lanes []*Lane
	for _, _reference := range reference.Lanes {
		if _instance, ok := stage.Lanes_instance[_reference]; ok {
			_Lanes = append(_Lanes, _instance)
		}
	}
	reference.Lanes = _Lanes
	var _Milestones []*Milestone
	for _, _reference := range reference.Milestones {
		if _instance, ok := stage.Milestones_instance[_reference]; ok {
			_Milestones = append(_Milestones, _instance)
		}
	}
	reference.Milestones = _Milestones
	var _Groups []*Group
	for _, _reference := range reference.Groups {
		if _instance, ok := stage.Groups_instance[_reference]; ok {
			_Groups = append(_Groups, _instance)
		}
	}
	reference.Groups = _Groups
	var _Arrows []*Arrow
	for _, _reference := range reference.Arrows {
		if _instance, ok := stage.Arrows_instance[_reference]; ok {
			_Arrows = append(_Arrows, _instance)
		}
	}
	reference.Arrows = _Arrows
}

func (reference *Group) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _GroupLanes []*Lane
	for _, _reference := range reference.GroupLanes {
		if _instance, ok := stage.Lanes_instance[_reference]; ok {
			_GroupLanes = append(_GroupLanes, _instance)
		}
	}
	reference.GroupLanes = _GroupLanes
}

func (reference *Lane) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _Bars []*Bar
	for _, _reference := range reference.Bars {
		if _instance, ok := stage.Bars_instance[_reference]; ok {
			_Bars = append(_Bars, _instance)
		}
	}
	reference.Bars = _Bars
}

func (reference *LaneUse) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.Lane; _reference != nil {
		reference.Lane = nil
		if _instance, ok := stage.Lanes_instance[_reference]; ok {
			reference.Lane = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *Milestone) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _LanesToDisplay []*Lane
	for _, _reference := range reference.LanesToDisplay {
		if _instance, ok := stage.Lanes_instance[_reference]; ok {
			_LanesToDisplay = append(_LanesToDisplay, _instance)
		}
	}
	reference.LanesToDisplay = _LanesToDisplay
}

// insertion point for diff per struct
// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (arrow *Arrow) GongDiff(stage *Stage, arrowOther *Arrow) (diffs []string) {
	// insertion point for field diffs
	if arrow.Name != arrowOther.Name {
		diffs = append(diffs, arrow.GongMarshallField(stage, "Name"))
	}
	if (arrow.From == nil) != (arrowOther.From == nil) {
		diffs = append(diffs, arrow.GongMarshallField(stage, "From"))
	} else if arrow.From != nil && arrowOther.From != nil {
		if arrow.From != arrowOther.From {
			diffs = append(diffs, arrow.GongMarshallField(stage, "From"))
		}
	}
	if (arrow.To == nil) != (arrowOther.To == nil) {
		diffs = append(diffs, arrow.GongMarshallField(stage, "To"))
	} else if arrow.To != nil && arrowOther.To != nil {
		if arrow.To != arrowOther.To {
			diffs = append(diffs, arrow.GongMarshallField(stage, "To"))
		}
	}
	if arrow.OptionnalColor != arrowOther.OptionnalColor {
		diffs = append(diffs, arrow.GongMarshallField(stage, "OptionnalColor"))
	}
	if arrow.OptionnalStroke != arrowOther.OptionnalStroke {
		diffs = append(diffs, arrow.GongMarshallField(stage, "OptionnalStroke"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (bar *Bar) GongDiff(stage *Stage, barOther *Bar) (diffs []string) {
	// insertion point for field diffs
	if bar.Name != barOther.Name {
		diffs = append(diffs, bar.GongMarshallField(stage, "Name"))
	}
	if bar.Start != barOther.Start {
		diffs = append(diffs, bar.GongMarshallField(stage, "Start"))
	}
	if bar.End != barOther.End {
		diffs = append(diffs, bar.GongMarshallField(stage, "End"))
	}
	if bar.ComputedDuration != barOther.ComputedDuration {
		diffs = append(diffs, bar.GongMarshallField(stage, "ComputedDuration"))
	}
	if bar.OptionnalColor != barOther.OptionnalColor {
		diffs = append(diffs, bar.GongMarshallField(stage, "OptionnalColor"))
	}
	if bar.OptionnalStroke != barOther.OptionnalStroke {
		diffs = append(diffs, bar.GongMarshallField(stage, "OptionnalStroke"))
	}
	if bar.FillOpacity != barOther.FillOpacity {
		diffs = append(diffs, bar.GongMarshallField(stage, "FillOpacity"))
	}
	if bar.StrokeWidth != barOther.StrokeWidth {
		diffs = append(diffs, bar.GongMarshallField(stage, "StrokeWidth"))
	}
	if bar.StrokeDashArray != barOther.StrokeDashArray {
		diffs = append(diffs, bar.GongMarshallField(stage, "StrokeDashArray"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (gantt *Gantt) GongDiff(stage *Stage, ganttOther *Gantt) (diffs []string) {
	// insertion point for field diffs
	if gantt.Name != ganttOther.Name {
		diffs = append(diffs, gantt.GongMarshallField(stage, "Name"))
	}
	if gantt.ComputedStart != ganttOther.ComputedStart {
		diffs = append(diffs, gantt.GongMarshallField(stage, "ComputedStart"))
	}
	if gantt.ComputedEnd != ganttOther.ComputedEnd {
		diffs = append(diffs, gantt.GongMarshallField(stage, "ComputedEnd"))
	}
	if gantt.ComputedDuration != ganttOther.ComputedDuration {
		diffs = append(diffs, gantt.GongMarshallField(stage, "ComputedDuration"))
	}
	if gantt.UseManualStartAndEndDates != ganttOther.UseManualStartAndEndDates {
		diffs = append(diffs, gantt.GongMarshallField(stage, "UseManualStartAndEndDates"))
	}
	if gantt.ManualStart != ganttOther.ManualStart {
		diffs = append(diffs, gantt.GongMarshallField(stage, "ManualStart"))
	}
	if gantt.ManualEnd != ganttOther.ManualEnd {
		diffs = append(diffs, gantt.GongMarshallField(stage, "ManualEnd"))
	}
	if gantt.LaneHeight != ganttOther.LaneHeight {
		diffs = append(diffs, gantt.GongMarshallField(stage, "LaneHeight"))
	}
	if gantt.RatioBarToLaneHeight != ganttOther.RatioBarToLaneHeight {
		diffs = append(diffs, gantt.GongMarshallField(stage, "RatioBarToLaneHeight"))
	}
	if gantt.YTopMargin != ganttOther.YTopMargin {
		diffs = append(diffs, gantt.GongMarshallField(stage, "YTopMargin"))
	}
	if gantt.XLeftText != ganttOther.XLeftText {
		diffs = append(diffs, gantt.GongMarshallField(stage, "XLeftText"))
	}
	if gantt.TextHeight != ganttOther.TextHeight {
		diffs = append(diffs, gantt.GongMarshallField(stage, "TextHeight"))
	}
	if gantt.XLeftLanes != ganttOther.XLeftLanes {
		diffs = append(diffs, gantt.GongMarshallField(stage, "XLeftLanes"))
	}
	if gantt.XRightMargin != ganttOther.XRightMargin {
		diffs = append(diffs, gantt.GongMarshallField(stage, "XRightMargin"))
	}
	if gantt.ArrowLengthToTheRightOfStartBar != ganttOther.ArrowLengthToTheRightOfStartBar {
		diffs = append(diffs, gantt.GongMarshallField(stage, "ArrowLengthToTheRightOfStartBar"))
	}
	if gantt.ArrowTipLenght != ganttOther.ArrowTipLenght {
		diffs = append(diffs, gantt.GongMarshallField(stage, "ArrowTipLenght"))
	}
	if gantt.TimeLine_Color != ganttOther.TimeLine_Color {
		diffs = append(diffs, gantt.GongMarshallField(stage, "TimeLine_Color"))
	}
	if gantt.TimeLine_FillOpacity != ganttOther.TimeLine_FillOpacity {
		diffs = append(diffs, gantt.GongMarshallField(stage, "TimeLine_FillOpacity"))
	}
	if gantt.TimeLine_Stroke != ganttOther.TimeLine_Stroke {
		diffs = append(diffs, gantt.GongMarshallField(stage, "TimeLine_Stroke"))
	}
	if gantt.TimeLine_StrokeWidth != ganttOther.TimeLine_StrokeWidth {
		diffs = append(diffs, gantt.GongMarshallField(stage, "TimeLine_StrokeWidth"))
	}
	if gantt.Group_Stroke != ganttOther.Group_Stroke {
		diffs = append(diffs, gantt.GongMarshallField(stage, "Group_Stroke"))
	}
	if gantt.Group_StrokeWidth != ganttOther.Group_StrokeWidth {
		diffs = append(diffs, gantt.GongMarshallField(stage, "Group_StrokeWidth"))
	}
	if gantt.Group_StrokeDashArray != ganttOther.Group_StrokeDashArray {
		diffs = append(diffs, gantt.GongMarshallField(stage, "Group_StrokeDashArray"))
	}
	if gantt.DateYOffset != ganttOther.DateYOffset {
		diffs = append(diffs, gantt.GongMarshallField(stage, "DateYOffset"))
	}
	if gantt.AlignOnStartEndOnYearStart != ganttOther.AlignOnStartEndOnYearStart {
		diffs = append(diffs, gantt.GongMarshallField(stage, "AlignOnStartEndOnYearStart"))
	}
	LanesDifferent := false
	if len(gantt.Lanes) != len(ganttOther.Lanes) {
		LanesDifferent = true
	} else {
		for i := range gantt.Lanes {
			if (gantt.Lanes[i] == nil) != (ganttOther.Lanes[i] == nil) {
				LanesDifferent = true
				break
			} else if gantt.Lanes[i] != nil && ganttOther.Lanes[i] != nil {
				// this is a pointer comparaison
				if gantt.Lanes[i] != ganttOther.Lanes[i] {
					LanesDifferent = true
					break
				}
			}
		}
	}
	if LanesDifferent {
		ops := stage.Diff(
			gantt,
			"Lanes",
			len(ganttOther.Lanes),
			len(gantt.Lanes),
			func(i, j int) bool {
				return ganttOther.Lanes[i] == gantt.Lanes[j]
			},
			func(j int) string {
				return gantt.Lanes[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	MilestonesDifferent := false
	if len(gantt.Milestones) != len(ganttOther.Milestones) {
		MilestonesDifferent = true
	} else {
		for i := range gantt.Milestones {
			if (gantt.Milestones[i] == nil) != (ganttOther.Milestones[i] == nil) {
				MilestonesDifferent = true
				break
			} else if gantt.Milestones[i] != nil && ganttOther.Milestones[i] != nil {
				// this is a pointer comparaison
				if gantt.Milestones[i] != ganttOther.Milestones[i] {
					MilestonesDifferent = true
					break
				}
			}
		}
	}
	if MilestonesDifferent {
		ops := stage.Diff(
			gantt,
			"Milestones",
			len(ganttOther.Milestones),
			len(gantt.Milestones),
			func(i, j int) bool {
				return ganttOther.Milestones[i] == gantt.Milestones[j]
			},
			func(j int) string {
				return gantt.Milestones[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	GroupsDifferent := false
	if len(gantt.Groups) != len(ganttOther.Groups) {
		GroupsDifferent = true
	} else {
		for i := range gantt.Groups {
			if (gantt.Groups[i] == nil) != (ganttOther.Groups[i] == nil) {
				GroupsDifferent = true
				break
			} else if gantt.Groups[i] != nil && ganttOther.Groups[i] != nil {
				// this is a pointer comparaison
				if gantt.Groups[i] != ganttOther.Groups[i] {
					GroupsDifferent = true
					break
				}
			}
		}
	}
	if GroupsDifferent {
		ops := stage.Diff(
			gantt,
			"Groups",
			len(ganttOther.Groups),
			len(gantt.Groups),
			func(i, j int) bool {
				return ganttOther.Groups[i] == gantt.Groups[j]
			},
			func(j int) string {
				return gantt.Groups[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	ArrowsDifferent := false
	if len(gantt.Arrows) != len(ganttOther.Arrows) {
		ArrowsDifferent = true
	} else {
		for i := range gantt.Arrows {
			if (gantt.Arrows[i] == nil) != (ganttOther.Arrows[i] == nil) {
				ArrowsDifferent = true
				break
			} else if gantt.Arrows[i] != nil && ganttOther.Arrows[i] != nil {
				// this is a pointer comparaison
				if gantt.Arrows[i] != ganttOther.Arrows[i] {
					ArrowsDifferent = true
					break
				}
			}
		}
	}
	if ArrowsDifferent {
		ops := stage.Diff(
			gantt,
			"Arrows",
			len(ganttOther.Arrows),
			len(gantt.Arrows),
			func(i, j int) bool {
				return ganttOther.Arrows[i] == gantt.Arrows[j]
			},
			func(j int) string {
				return gantt.Arrows[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (group *Group) GongDiff(stage *Stage, groupOther *Group) (diffs []string) {
	// insertion point for field diffs
	if group.Name != groupOther.Name {
		diffs = append(diffs, group.GongMarshallField(stage, "Name"))
	}
	GroupLanesDifferent := false
	if len(group.GroupLanes) != len(groupOther.GroupLanes) {
		GroupLanesDifferent = true
	} else {
		for i := range group.GroupLanes {
			if (group.GroupLanes[i] == nil) != (groupOther.GroupLanes[i] == nil) {
				GroupLanesDifferent = true
				break
			} else if group.GroupLanes[i] != nil && groupOther.GroupLanes[i] != nil {
				// this is a pointer comparaison
				if group.GroupLanes[i] != groupOther.GroupLanes[i] {
					GroupLanesDifferent = true
					break
				}
			}
		}
	}
	if GroupLanesDifferent {
		ops := stage.Diff(
			group,
			"GroupLanes",
			len(groupOther.GroupLanes),
			len(group.GroupLanes),
			func(i, j int) bool {
				return groupOther.GroupLanes[i] == group.GroupLanes[j]
			},
			func(j int) string {
				return group.GroupLanes[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (lane *Lane) GongDiff(stage *Stage, laneOther *Lane) (diffs []string) {
	// insertion point for field diffs
	if lane.Name != laneOther.Name {
		diffs = append(diffs, lane.GongMarshallField(stage, "Name"))
	}
	if lane.Order != laneOther.Order {
		diffs = append(diffs, lane.GongMarshallField(stage, "Order"))
	}
	BarsDifferent := false
	if len(lane.Bars) != len(laneOther.Bars) {
		BarsDifferent = true
	} else {
		for i := range lane.Bars {
			if (lane.Bars[i] == nil) != (laneOther.Bars[i] == nil) {
				BarsDifferent = true
				break
			} else if lane.Bars[i] != nil && laneOther.Bars[i] != nil {
				// this is a pointer comparaison
				if lane.Bars[i] != laneOther.Bars[i] {
					BarsDifferent = true
					break
				}
			}
		}
	}
	if BarsDifferent {
		ops := stage.Diff(
			lane,
			"Bars",
			len(laneOther.Bars),
			len(lane.Bars),
			func(i, j int) bool {
				return laneOther.Bars[i] == lane.Bars[j]
			},
			func(j int) string {
				return lane.Bars[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (laneuse *LaneUse) GongDiff(stage *Stage, laneuseOther *LaneUse) (diffs []string) {
	// insertion point for field diffs
	if laneuse.Name != laneuseOther.Name {
		diffs = append(diffs, laneuse.GongMarshallField(stage, "Name"))
	}
	if (laneuse.Lane == nil) != (laneuseOther.Lane == nil) {
		diffs = append(diffs, laneuse.GongMarshallField(stage, "Lane"))
	} else if laneuse.Lane != nil && laneuseOther.Lane != nil {
		if laneuse.Lane != laneuseOther.Lane {
			diffs = append(diffs, laneuse.GongMarshallField(stage, "Lane"))
		}
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (milestone *Milestone) GongDiff(stage *Stage, milestoneOther *Milestone) (diffs []string) {
	// insertion point for field diffs
	if milestone.Name != milestoneOther.Name {
		diffs = append(diffs, milestone.GongMarshallField(stage, "Name"))
	}
	if milestone.Date != milestoneOther.Date {
		diffs = append(diffs, milestone.GongMarshallField(stage, "Date"))
	}
	if milestone.DisplayVerticalBar != milestoneOther.DisplayVerticalBar {
		diffs = append(diffs, milestone.GongMarshallField(stage, "DisplayVerticalBar"))
	}
	LanesToDisplayDifferent := false
	if len(milestone.LanesToDisplay) != len(milestoneOther.LanesToDisplay) {
		LanesToDisplayDifferent = true
	} else {
		for i := range milestone.LanesToDisplay {
			if (milestone.LanesToDisplay[i] == nil) != (milestoneOther.LanesToDisplay[i] == nil) {
				LanesToDisplayDifferent = true
				break
			} else if milestone.LanesToDisplay[i] != nil && milestoneOther.LanesToDisplay[i] != nil {
				// this is a pointer comparaison
				if milestone.LanesToDisplay[i] != milestoneOther.LanesToDisplay[i] {
					LanesToDisplayDifferent = true
					break
				}
			}
		}
	}
	if LanesToDisplayDifferent {
		ops := stage.Diff(
			milestone,
			"LanesToDisplay",
			len(milestoneOther.LanesToDisplay),
			len(milestone.LanesToDisplay),
			func(i, j int) bool {
				return milestoneOther.LanesToDisplay[i] == milestone.LanesToDisplay[j]
			},
			func(j int) string {
				return milestone.LanesToDisplay[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
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

	for i := range m {
		for j := range n {
			if equal(i, j) {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				// Take the maximum of previous options
				dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
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
	for k := range m {
		if keptIndices[k] {
			keptOldIndices = append(keptOldIndices, k)
		}
	}

	lcsIdx := 0
	// Iterate through the NEW slice. If it matches the current LCS head, we keep it.
	// If it doesn't match, it must be inserted here.
	for k := range n {
		if lcsIdx < len(keptOldIndices) && equal(keptOldIndices[lcsIdx], k) {
			lcsIdx++
		} else {
			ops += fmt.Sprintf("\n\t%s.%s = slices.Insert( %s.%s, %d, %s)", a.GongGetIdentifier(stage), fieldName, a.GongGetIdentifier(stage), fieldName, k, getNewIdentifier(k))
		}
	}

	return ops
}
