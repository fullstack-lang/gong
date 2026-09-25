// generated code - do not edit
package models

import (
	"fmt"
	"slices"
)

// IsStaged is the Stage method checking if a gongstruct instance is staged.
func (stage *Stage) IsStaged(instance GongstructIF) (ok bool) {
	if instance != nil {
		return instance.GongIsStaged(stage)
	}
	return false
}

// insertion point for stage per struct
func (arrow *Arrow) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Arrows[arrow]
	return ok
}

func (bar *Bar) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Bars[bar]
	return ok
}

func (gantt *Gantt) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Gantts[gantt]
	return ok
}

func (group *Group) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Groups[group]
	return ok
}

func (lane *Lane) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Lanes[lane]
	return ok
}

func (laneuse *LaneUse) GongIsStaged(stage *Stage) bool {
	_, ok := stage.LaneUses[laneuse]
	return ok
}

func (milestone *Milestone) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Milestones[milestone]
	return ok
}

// StageBranch is the Stage method that stages instance and applies StageBranch recursively.
func (stage *Stage) StageBranch(instance GongstructIF) {
	if instance != nil {
		instance.GongStageBranch(stage)
	}
}

// insertion point for stage branch per struct
func (arrow *Arrow) GongStageBranch(stage *Stage) {

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

	// check if instance is already staged
	if stage.IsStaged(bar) {
		return
	}

	bar.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (gantt *Gantt) GongStageBranch(stage *Stage) {

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
	var alreadyCopied bool
	arrowTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, arrowFrom)
	if alreadyCopied {
		return
	}
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
	var alreadyCopied bool
	barTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, barFrom)
	if alreadyCopied {
		return
	}
	barFrom.GongCopyBasicFields(barTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchGantt(mapOrigCopy map[any]any, ganttFrom *Gantt) (ganttTo *Gantt) {
	var alreadyCopied bool
	ganttTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, ganttFrom)
	if alreadyCopied {
		return
	}
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
	var alreadyCopied bool
	groupTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, groupFrom)
	if alreadyCopied {
		return
	}
	groupFrom.GongCopyBasicFields(groupTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _lane := range groupFrom.GroupLanes {
		groupTo.GroupLanes = append(groupTo.GroupLanes, GongCopyBranchLane(mapOrigCopy, _lane))
	}

	return
}

func GongCopyBranchLane(mapOrigCopy map[any]any, laneFrom *Lane) (laneTo *Lane) {
	var alreadyCopied bool
	laneTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, laneFrom)
	if alreadyCopied {
		return
	}
	laneFrom.GongCopyBasicFields(laneTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _bar := range laneFrom.Bars {
		laneTo.Bars = append(laneTo.Bars, GongCopyBranchBar(mapOrigCopy, _bar))
	}

	return
}

func GongCopyBranchLaneUse(mapOrigCopy map[any]any, laneuseFrom *LaneUse) (laneuseTo *LaneUse) {
	var alreadyCopied bool
	laneuseTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, laneuseFrom)
	if alreadyCopied {
		return
	}
	laneuseFrom.GongCopyBasicFields(laneuseTo)

	//insertion point for the staging of instances referenced by pointers
	if laneuseFrom.Lane != nil {
		laneuseTo.Lane = GongCopyBranchLane(mapOrigCopy, laneuseFrom.Lane)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchMilestone(mapOrigCopy map[any]any, milestoneFrom *Milestone) (milestoneTo *Milestone) {
	var alreadyCopied bool
	milestoneTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, milestoneFrom)
	if alreadyCopied {
		return
	}
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

// insertion point for unstage branch per struct
func (arrow *Arrow) GongUnstageBranch(stage *Stage) {

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

	// check if instance is already staged
	if !stage.IsStaged(bar) {
		return
	}

	bar.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (gantt *Gantt) GongUnstageBranch(stage *Stage) {

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
	__gong__reconstructPointer(&reference.From, stage.Bars_reference, instance.From)
	__gong__reconstructPointer(&reference.To, stage.Bars_reference, instance.To)
	// insertion point for slice of pointers field
}

func (reference *Bar) GongReconstructPointersFromReferences(stage *Stage, instance *Bar) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Gantt) GongReconstructPointersFromReferences(stage *Stage, instance *Gantt) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Lanes, stage.Lanes_reference, instance.Lanes)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Milestones, stage.Milestones_reference, instance.Milestones)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Groups, stage.Groups_reference, instance.Groups)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Arrows, stage.Arrows_reference, instance.Arrows)
}

func (reference *Group) GongReconstructPointersFromReferences(stage *Stage, instance *Group) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.GroupLanes, stage.Lanes_reference, instance.GroupLanes)
}

func (reference *Lane) GongReconstructPointersFromReferences(stage *Stage, instance *Lane) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Bars, stage.Bars_reference, instance.Bars)
}

func (reference *LaneUse) GongReconstructPointersFromReferences(stage *Stage, instance *LaneUse) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Lane, stage.Lanes_reference, instance.Lane)
	// insertion point for slice of pointers field
}

func (reference *Milestone) GongReconstructPointersFromReferences(stage *Stage, instance *Milestone) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.LanesToDisplay, stage.Lanes_reference, instance.LanesToDisplay)
}

// insertion point for pointer reconstruction from instances
func (reference *Arrow) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.From, stage.Bars_instance)
	__gong__reconstructPointerFromInstance(&reference.To, stage.Bars_instance)
	// insertion point for slice of pointers fields
}

func (reference *Bar) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Gantt) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Lanes, stage.Lanes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Milestones, stage.Milestones_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Groups, stage.Groups_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Arrows, stage.Arrows_instance)
}

func (reference *Group) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.GroupLanes, stage.Lanes_instance)
}

func (reference *Lane) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Bars, stage.Bars_instance)
}

func (reference *LaneUse) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Lane, stage.Lanes_instance)
	// insertion point for slice of pointers fields
}

func (reference *Milestone) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.LanesToDisplay, stage.Lanes_instance)
}

// insertion point for diff per struct
// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (arrow *Arrow) GongDiff(stage *Stage, arrowOther *Arrow) (diffs []string) {
	// insertion point for field diffs
	if arrow.Name != arrowOther.Name {
		diffs = append(diffs, arrow.GongMarshallField(stage, "Name"))
	}
	if arrow.From != arrowOther.From {
		diffs = append(diffs, arrow.GongMarshallField(stage, "From"))
	}
	if arrow.To != arrowOther.To {
		diffs = append(diffs, arrow.GongMarshallField(stage, "To"))
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
	if ops := __gong__diffSliceOfPointers(stage, gantt, "Lanes", ganttOther.Lanes, gantt.Lanes); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, gantt, "Milestones", ganttOther.Milestones, gantt.Milestones); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, gantt, "Groups", ganttOther.Groups, gantt.Groups); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, gantt, "Arrows", ganttOther.Arrows, gantt.Arrows); ops != "" {
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
	if ops := __gong__diffSliceOfPointers(stage, group, "GroupLanes", groupOther.GroupLanes, group.GroupLanes); ops != "" {
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
	if ops := __gong__diffSliceOfPointers(stage, lane, "Bars", laneOther.Bars, lane.Bars); ops != "" {
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
	if laneuse.Lane != laneuseOther.Lane {
		diffs = append(diffs, laneuse.GongMarshallField(stage, "Lane"))
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
	if ops := __gong__diffSliceOfPointers(stage, milestone, "LanesToDisplay", milestoneOther.LanesToDisplay, milestone.LanesToDisplay); ops != "" {
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

func __gong__copyBranchCheck[T any](mapOrigCopy map[any]any, from *T) (*T, bool) {
	if to, ok := mapOrigCopy[from]; ok {
		return to.(*T), true
	}
	to := new(T)
	mapOrigCopy[from] = to
	return to, false
}

func __gong__reconstructPointer[T comparable](field *T, refMap map[T]T, instanceField T) {
	var zero T
	if instanceField != zero {
		*field = refMap[instanceField]
	}
}

func __gong__reconstructPointerFromInstance[T comparable](field *T, instMap map[T]T) {
	ref := *field
	var zero T
	if ref != zero {
		*field = zero
		if inst, ok := instMap[ref]; ok {
			*field = inst
		}
	}
}

func __gong__reconstructSliceOfPointersFromReferences[T comparable](field *[]T, refMap map[T]T, instanceSlice []T) {
	*field = (*field)[:0]
	for _, b := range instanceSlice {
		*field = append(*field, refMap[b])
	}
}

func __gong__reconstructSliceOfPointersFromInstances[T comparable](field *[]T, instMap map[T]T) {
	var res []T
	for _, ref := range *field {
		if inst, ok := instMap[ref]; ok {
			res = append(res, inst)
		}
	}
	*field = res
}

func __gong__diffSliceOfPointers[T interface {
	comparable
	GongstructIF
}](
	stage *Stage,
	instance GongstructIF,
	fieldName string,
	oldSlice, newSlice []T,
) string {
	if slices.Equal(oldSlice, newSlice) {
		return ""
	}
	return stage.Diff(
		instance,
		fieldName,
		len(oldSlice),
		len(newSlice),
		func(i, j int) bool {
			return oldSlice[i] == newSlice[j]
		},
		func(j int) string {
			return newSlice[j].GongGetIdentifier(stage)
		},
	)
}
