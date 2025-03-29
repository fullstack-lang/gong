// generated code - do not edit
package models

func IsStaged[Type Gongstruct](stage *Stage, instance *Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
	case *Arrow:
		ok = stage.IsStagedArrow(target)

	case *Bar:
		ok = stage.IsStagedBar(target)

	case *Gantt:
		ok = stage.IsStagedGantt(target)

	case *Group:
		ok = stage.IsStagedGroup(target)

	case *Lane:
		ok = stage.IsStagedLane(target)

	case *LaneUse:
		ok = stage.IsStagedLaneUse(target)

	case *Milestone:
		ok = stage.IsStagedMilestone(target)

	default:
		_ = target
	}
	return
}

// insertion point for stage per struct
func (stage *Stage) IsStagedArrow(arrow *Arrow) (ok bool) {

	_, ok = stage.Arrows[arrow]

	return
}

func (stage *Stage) IsStagedBar(bar *Bar) (ok bool) {

	_, ok = stage.Bars[bar]

	return
}

func (stage *Stage) IsStagedGantt(gantt *Gantt) (ok bool) {

	_, ok = stage.Gantts[gantt]

	return
}

func (stage *Stage) IsStagedGroup(group *Group) (ok bool) {

	_, ok = stage.Groups[group]

	return
}

func (stage *Stage) IsStagedLane(lane *Lane) (ok bool) {

	_, ok = stage.Lanes[lane]

	return
}

func (stage *Stage) IsStagedLaneUse(laneuse *LaneUse) (ok bool) {

	_, ok = stage.LaneUses[laneuse]

	return
}

func (stage *Stage) IsStagedMilestone(milestone *Milestone) (ok bool) {

	_, ok = stage.Milestones[milestone]

	return
}

// StageBranch stages instance and apply StageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func StageBranch[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for stage branch
	case *Arrow:
		stage.StageBranchArrow(target)

	case *Bar:
		stage.StageBranchBar(target)

	case *Gantt:
		stage.StageBranchGantt(target)

	case *Group:
		stage.StageBranchGroup(target)

	case *Lane:
		stage.StageBranchLane(target)

	case *LaneUse:
		stage.StageBranchLaneUse(target)

	case *Milestone:
		stage.StageBranchMilestone(target)

	default:
		_ = target
	}
}

// insertion point for stage branch per struct
func (stage *Stage) StageBranchArrow(arrow *Arrow) {

	// check if instance is already staged
	if IsStaged(stage, arrow) {
		return
	}

	arrow.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if arrow.From != nil {
		StageBranch(stage, arrow.From)
	}
	if arrow.To != nil {
		StageBranch(stage, arrow.To)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchBar(bar *Bar) {

	// check if instance is already staged
	if IsStaged(stage, bar) {
		return
	}

	bar.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchGantt(gantt *Gantt) {

	// check if instance is already staged
	if IsStaged(stage, gantt) {
		return
	}

	gantt.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _lane := range gantt.Lanes {
		StageBranch(stage, _lane)
	}
	for _, _milestone := range gantt.Milestones {
		StageBranch(stage, _milestone)
	}
	for _, _group := range gantt.Groups {
		StageBranch(stage, _group)
	}
	for _, _arrow := range gantt.Arrows {
		StageBranch(stage, _arrow)
	}

}

func (stage *Stage) StageBranchGroup(group *Group) {

	// check if instance is already staged
	if IsStaged(stage, group) {
		return
	}

	group.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _lane := range group.GroupLanes {
		StageBranch(stage, _lane)
	}

}

func (stage *Stage) StageBranchLane(lane *Lane) {

	// check if instance is already staged
	if IsStaged(stage, lane) {
		return
	}

	lane.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _bar := range lane.Bars {
		StageBranch(stage, _bar)
	}

}

func (stage *Stage) StageBranchLaneUse(laneuse *LaneUse) {

	// check if instance is already staged
	if IsStaged(stage, laneuse) {
		return
	}

	laneuse.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if laneuse.Lane != nil {
		StageBranch(stage, laneuse.Lane)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchMilestone(milestone *Milestone) {

	// check if instance is already staged
	if IsStaged(stage, milestone) {
		return
	}

	milestone.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _laneuse := range milestone.LanesToDisplayMilestoneUse {
		StageBranch(stage, _laneuse)
	}

}

// CopyBranch stages instance and apply CopyBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func CopyBranch[Type Gongstruct](from *Type) (to *Type) {

	mapOrigCopy := make(map[any]any)
	_ = mapOrigCopy

	switch fromT := any(from).(type) {
	// insertion point for stage branch
	case *Arrow:
		toT := CopyBranchArrow(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Bar:
		toT := CopyBranchBar(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Gantt:
		toT := CopyBranchGantt(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Group:
		toT := CopyBranchGroup(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Lane:
		toT := CopyBranchLane(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *LaneUse:
		toT := CopyBranchLaneUse(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Milestone:
		toT := CopyBranchMilestone(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
func CopyBranchArrow(mapOrigCopy map[any]any, arrowFrom *Arrow) (arrowTo *Arrow) {

	// arrowFrom has already been copied
	if _arrowTo, ok := mapOrigCopy[arrowFrom]; ok {
		arrowTo = _arrowTo.(*Arrow)
		return
	}

	arrowTo = new(Arrow)
	mapOrigCopy[arrowFrom] = arrowTo
	arrowFrom.CopyBasicFields(arrowTo)

	//insertion point for the staging of instances referenced by pointers
	if arrowFrom.From != nil {
		arrowTo.From = CopyBranchBar(mapOrigCopy, arrowFrom.From)
	}
	if arrowFrom.To != nil {
		arrowTo.To = CopyBranchBar(mapOrigCopy, arrowFrom.To)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchBar(mapOrigCopy map[any]any, barFrom *Bar) (barTo *Bar) {

	// barFrom has already been copied
	if _barTo, ok := mapOrigCopy[barFrom]; ok {
		barTo = _barTo.(*Bar)
		return
	}

	barTo = new(Bar)
	mapOrigCopy[barFrom] = barTo
	barFrom.CopyBasicFields(barTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchGantt(mapOrigCopy map[any]any, ganttFrom *Gantt) (ganttTo *Gantt) {

	// ganttFrom has already been copied
	if _ganttTo, ok := mapOrigCopy[ganttFrom]; ok {
		ganttTo = _ganttTo.(*Gantt)
		return
	}

	ganttTo = new(Gantt)
	mapOrigCopy[ganttFrom] = ganttTo
	ganttFrom.CopyBasicFields(ganttTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _lane := range ganttFrom.Lanes {
		ganttTo.Lanes = append(ganttTo.Lanes, CopyBranchLane(mapOrigCopy, _lane))
	}
	for _, _milestone := range ganttFrom.Milestones {
		ganttTo.Milestones = append(ganttTo.Milestones, CopyBranchMilestone(mapOrigCopy, _milestone))
	}
	for _, _group := range ganttFrom.Groups {
		ganttTo.Groups = append(ganttTo.Groups, CopyBranchGroup(mapOrigCopy, _group))
	}
	for _, _arrow := range ganttFrom.Arrows {
		ganttTo.Arrows = append(ganttTo.Arrows, CopyBranchArrow(mapOrigCopy, _arrow))
	}

	return
}

func CopyBranchGroup(mapOrigCopy map[any]any, groupFrom *Group) (groupTo *Group) {

	// groupFrom has already been copied
	if _groupTo, ok := mapOrigCopy[groupFrom]; ok {
		groupTo = _groupTo.(*Group)
		return
	}

	groupTo = new(Group)
	mapOrigCopy[groupFrom] = groupTo
	groupFrom.CopyBasicFields(groupTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _lane := range groupFrom.GroupLanes {
		groupTo.GroupLanes = append(groupTo.GroupLanes, CopyBranchLane(mapOrigCopy, _lane))
	}

	return
}

func CopyBranchLane(mapOrigCopy map[any]any, laneFrom *Lane) (laneTo *Lane) {

	// laneFrom has already been copied
	if _laneTo, ok := mapOrigCopy[laneFrom]; ok {
		laneTo = _laneTo.(*Lane)
		return
	}

	laneTo = new(Lane)
	mapOrigCopy[laneFrom] = laneTo
	laneFrom.CopyBasicFields(laneTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _bar := range laneFrom.Bars {
		laneTo.Bars = append(laneTo.Bars, CopyBranchBar(mapOrigCopy, _bar))
	}

	return
}

func CopyBranchLaneUse(mapOrigCopy map[any]any, laneuseFrom *LaneUse) (laneuseTo *LaneUse) {

	// laneuseFrom has already been copied
	if _laneuseTo, ok := mapOrigCopy[laneuseFrom]; ok {
		laneuseTo = _laneuseTo.(*LaneUse)
		return
	}

	laneuseTo = new(LaneUse)
	mapOrigCopy[laneuseFrom] = laneuseTo
	laneuseFrom.CopyBasicFields(laneuseTo)

	//insertion point for the staging of instances referenced by pointers
	if laneuseFrom.Lane != nil {
		laneuseTo.Lane = CopyBranchLane(mapOrigCopy, laneuseFrom.Lane)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchMilestone(mapOrigCopy map[any]any, milestoneFrom *Milestone) (milestoneTo *Milestone) {

	// milestoneFrom has already been copied
	if _milestoneTo, ok := mapOrigCopy[milestoneFrom]; ok {
		milestoneTo = _milestoneTo.(*Milestone)
		return
	}

	milestoneTo = new(Milestone)
	mapOrigCopy[milestoneFrom] = milestoneTo
	milestoneFrom.CopyBasicFields(milestoneTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _laneuse := range milestoneFrom.LanesToDisplayMilestoneUse {
		milestoneTo.LanesToDisplayMilestoneUse = append(milestoneTo.LanesToDisplayMilestoneUse, CopyBranchLaneUse(mapOrigCopy, _laneuse))
	}

	return
}

// UnstageBranch stages instance and apply UnstageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the insance
//
// the algorithm stops along the course of graph if a vertex is already staged
func UnstageBranch[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for unstage branch
	case *Arrow:
		stage.UnstageBranchArrow(target)

	case *Bar:
		stage.UnstageBranchBar(target)

	case *Gantt:
		stage.UnstageBranchGantt(target)

	case *Group:
		stage.UnstageBranchGroup(target)

	case *Lane:
		stage.UnstageBranchLane(target)

	case *LaneUse:
		stage.UnstageBranchLaneUse(target)

	case *Milestone:
		stage.UnstageBranchMilestone(target)

	default:
		_ = target
	}
}

// insertion point for unstage branch per struct
func (stage *Stage) UnstageBranchArrow(arrow *Arrow) {

	// check if instance is already staged
	if !IsStaged(stage, arrow) {
		return
	}

	arrow.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if arrow.From != nil {
		UnstageBranch(stage, arrow.From)
	}
	if arrow.To != nil {
		UnstageBranch(stage, arrow.To)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchBar(bar *Bar) {

	// check if instance is already staged
	if !IsStaged(stage, bar) {
		return
	}

	bar.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchGantt(gantt *Gantt) {

	// check if instance is already staged
	if !IsStaged(stage, gantt) {
		return
	}

	gantt.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _lane := range gantt.Lanes {
		UnstageBranch(stage, _lane)
	}
	for _, _milestone := range gantt.Milestones {
		UnstageBranch(stage, _milestone)
	}
	for _, _group := range gantt.Groups {
		UnstageBranch(stage, _group)
	}
	for _, _arrow := range gantt.Arrows {
		UnstageBranch(stage, _arrow)
	}

}

func (stage *Stage) UnstageBranchGroup(group *Group) {

	// check if instance is already staged
	if !IsStaged(stage, group) {
		return
	}

	group.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _lane := range group.GroupLanes {
		UnstageBranch(stage, _lane)
	}

}

func (stage *Stage) UnstageBranchLane(lane *Lane) {

	// check if instance is already staged
	if !IsStaged(stage, lane) {
		return
	}

	lane.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _bar := range lane.Bars {
		UnstageBranch(stage, _bar)
	}

}

func (stage *Stage) UnstageBranchLaneUse(laneuse *LaneUse) {

	// check if instance is already staged
	if !IsStaged(stage, laneuse) {
		return
	}

	laneuse.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if laneuse.Lane != nil {
		UnstageBranch(stage, laneuse.Lane)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchMilestone(milestone *Milestone) {

	// check if instance is already staged
	if !IsStaged(stage, milestone) {
		return
	}

	milestone.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _laneuse := range milestone.LanesToDisplayMilestoneUse {
		UnstageBranch(stage, _laneuse)
	}

}
