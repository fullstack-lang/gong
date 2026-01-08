// generated code - do not edit
package models

import "fmt"

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
		ops := Diff(stage, gantt, ganttOther, "Lanes", ganttOther.Lanes, gantt.Lanes)
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
		ops := Diff(stage, gantt, ganttOther, "Milestones", ganttOther.Milestones, gantt.Milestones)
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
		ops := Diff(stage, gantt, ganttOther, "Groups", ganttOther.Groups, gantt.Groups)
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
		ops := Diff(stage, gantt, ganttOther, "Arrows", ganttOther.Arrows, gantt.Arrows)
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
		ops := Diff(stage, group, groupOther, "GroupLanes", groupOther.GroupLanes, group.GroupLanes)
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
		ops := Diff(stage, lane, laneOther, "Bars", laneOther.Bars, lane.Bars)
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
		ops := Diff(stage, milestone, milestoneOther, "LanesToDisplay", milestoneOther.LanesToDisplay, milestone.LanesToDisplay)
		diffs = append(diffs, ops)
	}

	return
}

// Diff returns the sequence of operations to transform oldSlice into newSlice.
// It requires type T to be comparable (e.g., pointers, ints, strings).
func Diff[T1, T2 PointerToGongstruct](stage *Stage, a, b T1, fieldName string, oldSlice, newSlice []T2) (ops string) {
	m, n := len(oldSlice), len(newSlice)

	// 1. Build the LCS (Longest Common Subsequence) Matrix
	// This helps us find the "anchor" elements that shouldn't move.
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if oldSlice[i] == newSlice[j] {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				// Take the maximum of previous options
				if dp[i][j+1] > dp[i+1][j] {
					dp[i+1][j+1] = dp[i][j+1]
				} else {
					dp[i+1][j+1] = dp[i+1][j]
				}
			}
		}
	}

	// 2. Backtrack to find which indices in oldSlice are part of the LCS
	// We use a map for O(1) lookups.
	keptIndices := make(map[int]bool)
	i, j := m, n
	for i > 0 && j > 0 {
		if oldSlice[i-1] == newSlice[j-1] {
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
			ops += fmt.Sprintf("\n\t%s.%s = slices.Delete( %s.%s, %d, %d)", a.GongGetIdentifier(stage), fieldName, a.GongGetIdentifier(stage), fieldName, k, k+1)
		}
	}

	// 4. PHASE 2: Generate Insertions
	// We simulate the state of the slice after deletions to determine insertion points.
	// The 'current' slice essentially consists of only the kept LCS items.

	// Create a temporary view of what's left after deletions for tracking matches
	var currentLCS []T2
	for k := 0; k < m; k++ {
		if keptIndices[k] {
			currentLCS = append(currentLCS, oldSlice[k])
		}
	}

	lcsIdx := 0
	// Iterate through the NEW slice. If it matches the current LCS head, we keep it.
	// If it doesn't match, it must be inserted here.
	for k, targetVal := range newSlice {
		if lcsIdx < len(currentLCS) && currentLCS[lcsIdx] == targetVal {
			lcsIdx++
		} else {
			ops += fmt.Sprintf("\n\t%s.%s = slices.Insert( %s.%s, %d, %s)", a.GongGetIdentifier(stage), fieldName, a.GongGetIdentifier(stage), fieldName, k, targetVal.GongGetIdentifier(stage))
		}
	}

	return ops
}
