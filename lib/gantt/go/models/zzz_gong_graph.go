// generated code - do not edit
package models

func IsStagedPointerToGongstruct[Type PointerToGongstruct](stage *Stage, instance Type) (ok bool) {

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
	for _, _lane := range milestone.LanesToDisplay {
		StageBranch(stage, _lane)
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
	for _, _lane := range milestoneFrom.LanesToDisplay {
		milestoneTo.LanesToDisplay = append(milestoneTo.LanesToDisplay, CopyBranchLane(mapOrigCopy, _lane))
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
	for _, _lane := range milestone.LanesToDisplay {
		UnstageBranch(stage, _lane)
	}

}


// insertion point for diff per struct
// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (arrow *Arrow) GongDiff(arrowOther *Arrow) (diffs []string) {
	// insertion point for field diffs
	if arrow.Name != arrowOther.Name {
		diffs = append(diffs, "Name")
	}
	if (arrow.From == nil) != (arrowOther.From == nil) {
		diffs = append(diffs, "From")
	} else if arrow.From != nil && arrowOther.From != nil {
		if arrow.From != arrowOther.From {
			diffs = append(diffs, "From")
		}
	}
	if (arrow.To == nil) != (arrowOther.To == nil) {
		diffs = append(diffs, "To")
	} else if arrow.To != nil && arrowOther.To != nil {
		if arrow.To != arrowOther.To {
			diffs = append(diffs, "To")
		}
	}
	if arrow.OptionnalColor != arrowOther.OptionnalColor {
		diffs = append(diffs, "OptionnalColor")
	}
	if arrow.OptionnalStroke != arrowOther.OptionnalStroke {
		diffs = append(diffs, "OptionnalStroke")
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (bar *Bar) GongDiff(barOther *Bar) (diffs []string) {
	// insertion point for field diffs
	if bar.Name != barOther.Name {
		diffs = append(diffs, "Name")
	}
	if bar.ComputedDuration != barOther.ComputedDuration {
		diffs = append(diffs, "ComputedDuration")
	}
	if bar.OptionnalColor != barOther.OptionnalColor {
		diffs = append(diffs, "OptionnalColor")
	}
	if bar.OptionnalStroke != barOther.OptionnalStroke {
		diffs = append(diffs, "OptionnalStroke")
	}
	if bar.FillOpacity != barOther.FillOpacity {
		diffs = append(diffs, "FillOpacity")
	}
	if bar.StrokeWidth != barOther.StrokeWidth {
		diffs = append(diffs, "StrokeWidth")
	}
	if bar.StrokeDashArray != barOther.StrokeDashArray {
		diffs = append(diffs, "StrokeDashArray")
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (gantt *Gantt) GongDiff(ganttOther *Gantt) (diffs []string) {
	// insertion point for field diffs
	if gantt.Name != ganttOther.Name {
		diffs = append(diffs, "Name")
	}
	if gantt.ComputedDuration != ganttOther.ComputedDuration {
		diffs = append(diffs, "ComputedDuration")
	}
	if gantt.UseManualStartAndEndDates != ganttOther.UseManualStartAndEndDates {
		diffs = append(diffs, "UseManualStartAndEndDates")
	}
	if gantt.LaneHeight != ganttOther.LaneHeight {
		diffs = append(diffs, "LaneHeight")
	}
	if gantt.RatioBarToLaneHeight != ganttOther.RatioBarToLaneHeight {
		diffs = append(diffs, "RatioBarToLaneHeight")
	}
	if gantt.YTopMargin != ganttOther.YTopMargin {
		diffs = append(diffs, "YTopMargin")
	}
	if gantt.XLeftText != ganttOther.XLeftText {
		diffs = append(diffs, "XLeftText")
	}
	if gantt.TextHeight != ganttOther.TextHeight {
		diffs = append(diffs, "TextHeight")
	}
	if gantt.XLeftLanes != ganttOther.XLeftLanes {
		diffs = append(diffs, "XLeftLanes")
	}
	if gantt.XRightMargin != ganttOther.XRightMargin {
		diffs = append(diffs, "XRightMargin")
	}
	if gantt.ArrowLengthToTheRightOfStartBar != ganttOther.ArrowLengthToTheRightOfStartBar {
		diffs = append(diffs, "ArrowLengthToTheRightOfStartBar")
	}
	if gantt.ArrowTipLenght != ganttOther.ArrowTipLenght {
		diffs = append(diffs, "ArrowTipLenght")
	}
	if gantt.TimeLine_Color != ganttOther.TimeLine_Color {
		diffs = append(diffs, "TimeLine_Color")
	}
	if gantt.TimeLine_FillOpacity != ganttOther.TimeLine_FillOpacity {
		diffs = append(diffs, "TimeLine_FillOpacity")
	}
	if gantt.TimeLine_Stroke != ganttOther.TimeLine_Stroke {
		diffs = append(diffs, "TimeLine_Stroke")
	}
	if gantt.TimeLine_StrokeWidth != ganttOther.TimeLine_StrokeWidth {
		diffs = append(diffs, "TimeLine_StrokeWidth")
	}
	if gantt.Group_Stroke != ganttOther.Group_Stroke {
		diffs = append(diffs, "Group_Stroke")
	}
	if gantt.Group_StrokeWidth != ganttOther.Group_StrokeWidth {
		diffs = append(diffs, "Group_StrokeWidth")
	}
	if gantt.Group_StrokeDashArray != ganttOther.Group_StrokeDashArray {
		diffs = append(diffs, "Group_StrokeDashArray")
	}
	if gantt.DateYOffset != ganttOther.DateYOffset {
		diffs = append(diffs, "DateYOffset")
	}
	if gantt.AlignOnStartEndOnYearStart != ganttOther.AlignOnStartEndOnYearStart {
		diffs = append(diffs, "AlignOnStartEndOnYearStart")
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
                if len(gantt.Lanes[i].GongDiff(ganttOther.Lanes[i])) > 0 {
                    LanesDifferent = true
                    break
                }
            }
        }
    }
    if LanesDifferent {
        diffs = append(diffs, "Lanes")
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
                if len(gantt.Milestones[i].GongDiff(ganttOther.Milestones[i])) > 0 {
                    MilestonesDifferent = true
                    break
                }
            }
        }
    }
    if MilestonesDifferent {
        diffs = append(diffs, "Milestones")
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
                if len(gantt.Groups[i].GongDiff(ganttOther.Groups[i])) > 0 {
                    GroupsDifferent = true
                    break
                }
            }
        }
    }
    if GroupsDifferent {
        diffs = append(diffs, "Groups")
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
                if len(gantt.Arrows[i].GongDiff(ganttOther.Arrows[i])) > 0 {
                    ArrowsDifferent = true
                    break
                }
            }
        }
    }
    if ArrowsDifferent {
        diffs = append(diffs, "Arrows")
    }

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (group *Group) GongDiff(groupOther *Group) (diffs []string) {
	// insertion point for field diffs
	if group.Name != groupOther.Name {
		diffs = append(diffs, "Name")
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
                if len(group.GroupLanes[i].GongDiff(groupOther.GroupLanes[i])) > 0 {
                    GroupLanesDifferent = true
                    break
                }
            }
        }
    }
    if GroupLanesDifferent {
        diffs = append(diffs, "GroupLanes")
    }

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (lane *Lane) GongDiff(laneOther *Lane) (diffs []string) {
	// insertion point for field diffs
	if lane.Name != laneOther.Name {
		diffs = append(diffs, "Name")
	}
	if lane.Order != laneOther.Order {
		diffs = append(diffs, "Order")
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
                if len(lane.Bars[i].GongDiff(laneOther.Bars[i])) > 0 {
                    BarsDifferent = true
                    break
                }
            }
        }
    }
    if BarsDifferent {
        diffs = append(diffs, "Bars")
    }

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (laneuse *LaneUse) GongDiff(laneuseOther *LaneUse) (diffs []string) {
	// insertion point for field diffs
	if laneuse.Name != laneuseOther.Name {
		diffs = append(diffs, "Name")
	}
	if (laneuse.Lane == nil) != (laneuseOther.Lane == nil) {
		diffs = append(diffs, "Lane")
	} else if laneuse.Lane != nil && laneuseOther.Lane != nil {
		if laneuse.Lane != laneuseOther.Lane {
			diffs = append(diffs, "Lane")
		}
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (milestone *Milestone) GongDiff(milestoneOther *Milestone) (diffs []string) {
	// insertion point for field diffs
	if milestone.Name != milestoneOther.Name {
		diffs = append(diffs, "Name")
	}
	if milestone.DisplayVerticalBar != milestoneOther.DisplayVerticalBar {
		diffs = append(diffs, "DisplayVerticalBar")
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
                if len(milestone.LanesToDisplay[i].GongDiff(milestoneOther.LanesToDisplay[i])) > 0 {
                    LanesToDisplayDifferent = true
                    break
                }
            }
        }
    }
    if LanesToDisplayDifferent {
        diffs = append(diffs, "LanesToDisplay")
    }

	return
}
