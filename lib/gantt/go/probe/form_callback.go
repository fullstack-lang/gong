// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	table "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gong/lib/gantt/go/models"
	"github.com/fullstack-lang/gong/lib/gantt/go/orm"
)

const __dummmy__time = time.Nanosecond

var __dummmy__letters = slices.Delete([]string{"a"}, 0, 1)
var __dummy_orm = orm.BackRepoStruct{}

// insertion point
func __gong__New__ArrowFormCallback(
	arrow *models.Arrow,
	probe *Probe,
	formGroup *table.FormGroup,
) (arrowFormCallback *ArrowFormCallback) {
	arrowFormCallback = new(ArrowFormCallback)
	arrowFormCallback.probe = probe
	arrowFormCallback.arrow = arrow
	arrowFormCallback.formGroup = formGroup

	arrowFormCallback.CreationMode = (arrow == nil)

	return
}

type ArrowFormCallback struct {
	arrow *models.Arrow

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (arrowFormCallback *ArrowFormCallback) OnSave() {

	log.Println("ArrowFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	arrowFormCallback.probe.formStage.Checkout()

	if arrowFormCallback.arrow == nil {
		arrowFormCallback.arrow = new(models.Arrow).Stage(arrowFormCallback.probe.stageOfInterest)
	}
	arrow_ := arrowFormCallback.arrow
	_ = arrow_

	for _, formDiv := range arrowFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(arrow_.Name), formDiv)
		case "From":
			FormDivSelectFieldToField(&(arrow_.From), arrowFormCallback.probe.stageOfInterest, formDiv)
		case "To":
			FormDivSelectFieldToField(&(arrow_.To), arrowFormCallback.probe.stageOfInterest, formDiv)
		case "OptionnalColor":
			FormDivBasicFieldToField(&(arrow_.OptionnalColor), formDiv)
		case "OptionnalStroke":
			FormDivBasicFieldToField(&(arrow_.OptionnalStroke), formDiv)
		case "Gantt:Arrows":
			// we need to retrieve the field owner before the change
			var pastGanttOwner *models.Gantt
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Gantt"
			rf.Fieldname = "Arrows"
			reverseFieldOwner := models.GetReverseFieldOwner(
				arrowFormCallback.probe.stageOfInterest,
				arrow_,
				&rf)

			if reverseFieldOwner != nil {
				pastGanttOwner = reverseFieldOwner.(*models.Gantt)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastGanttOwner != nil {
					idx := slices.Index(pastGanttOwner.Arrows, arrow_)
					pastGanttOwner.Arrows = slices.Delete(pastGanttOwner.Arrows, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _gantt := range *models.GetGongstructInstancesSet[models.Gantt](arrowFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _gantt.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newGanttOwner := _gantt // we have a match
						if pastGanttOwner != nil {
							if newGanttOwner != pastGanttOwner {
								idx := slices.Index(pastGanttOwner.Arrows, arrow_)
								pastGanttOwner.Arrows = slices.Delete(pastGanttOwner.Arrows, idx, idx+1)
								newGanttOwner.Arrows = append(newGanttOwner.Arrows, arrow_)
							}
						} else {
							newGanttOwner.Arrows = append(newGanttOwner.Arrows, arrow_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if arrowFormCallback.formGroup.HasSuppressButtonBeenPressed {
		arrow_.Unstage(arrowFormCallback.probe.stageOfInterest)
	}

	arrowFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Arrow](
		arrowFormCallback.probe,
	)
	arrowFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if arrowFormCallback.CreationMode || arrowFormCallback.formGroup.HasSuppressButtonBeenPressed {
		arrowFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(arrowFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ArrowFormCallback(
			nil,
			arrowFormCallback.probe,
			newFormGroup,
		)
		arrow := new(models.Arrow)
		FillUpForm(arrow, newFormGroup, arrowFormCallback.probe)
		arrowFormCallback.probe.formStage.Commit()
	}

	fillUpTree(arrowFormCallback.probe)
}
func __gong__New__BarFormCallback(
	bar *models.Bar,
	probe *Probe,
	formGroup *table.FormGroup,
) (barFormCallback *BarFormCallback) {
	barFormCallback = new(BarFormCallback)
	barFormCallback.probe = probe
	barFormCallback.bar = bar
	barFormCallback.formGroup = formGroup

	barFormCallback.CreationMode = (bar == nil)

	return
}

type BarFormCallback struct {
	bar *models.Bar

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (barFormCallback *BarFormCallback) OnSave() {

	log.Println("BarFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	barFormCallback.probe.formStage.Checkout()

	if barFormCallback.bar == nil {
		barFormCallback.bar = new(models.Bar).Stage(barFormCallback.probe.stageOfInterest)
	}
	bar_ := barFormCallback.bar
	_ = bar_

	for _, formDiv := range barFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(bar_.Name), formDiv)
		case "Start":
			FormDivBasicFieldToField(&(bar_.Start), formDiv)
		case "End":
			FormDivBasicFieldToField(&(bar_.End), formDiv)
		case "ComputedDuration":
			FormDivBasicFieldToField(&(bar_.ComputedDuration), formDiv)
		case "OptionnalColor":
			FormDivBasicFieldToField(&(bar_.OptionnalColor), formDiv)
		case "OptionnalStroke":
			FormDivBasicFieldToField(&(bar_.OptionnalStroke), formDiv)
		case "FillOpacity":
			FormDivBasicFieldToField(&(bar_.FillOpacity), formDiv)
		case "StrokeWidth":
			FormDivBasicFieldToField(&(bar_.StrokeWidth), formDiv)
		case "StrokeDashArray":
			FormDivBasicFieldToField(&(bar_.StrokeDashArray), formDiv)
		case "Lane:Bars":
			// we need to retrieve the field owner before the change
			var pastLaneOwner *models.Lane
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Lane"
			rf.Fieldname = "Bars"
			reverseFieldOwner := models.GetReverseFieldOwner(
				barFormCallback.probe.stageOfInterest,
				bar_,
				&rf)

			if reverseFieldOwner != nil {
				pastLaneOwner = reverseFieldOwner.(*models.Lane)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastLaneOwner != nil {
					idx := slices.Index(pastLaneOwner.Bars, bar_)
					pastLaneOwner.Bars = slices.Delete(pastLaneOwner.Bars, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _lane := range *models.GetGongstructInstancesSet[models.Lane](barFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _lane.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newLaneOwner := _lane // we have a match
						if pastLaneOwner != nil {
							if newLaneOwner != pastLaneOwner {
								idx := slices.Index(pastLaneOwner.Bars, bar_)
								pastLaneOwner.Bars = slices.Delete(pastLaneOwner.Bars, idx, idx+1)
								newLaneOwner.Bars = append(newLaneOwner.Bars, bar_)
							}
						} else {
							newLaneOwner.Bars = append(newLaneOwner.Bars, bar_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if barFormCallback.formGroup.HasSuppressButtonBeenPressed {
		bar_.Unstage(barFormCallback.probe.stageOfInterest)
	}

	barFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Bar](
		barFormCallback.probe,
	)
	barFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if barFormCallback.CreationMode || barFormCallback.formGroup.HasSuppressButtonBeenPressed {
		barFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(barFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__BarFormCallback(
			nil,
			barFormCallback.probe,
			newFormGroup,
		)
		bar := new(models.Bar)
		FillUpForm(bar, newFormGroup, barFormCallback.probe)
		barFormCallback.probe.formStage.Commit()
	}

	fillUpTree(barFormCallback.probe)
}
func __gong__New__GanttFormCallback(
	gantt *models.Gantt,
	probe *Probe,
	formGroup *table.FormGroup,
) (ganttFormCallback *GanttFormCallback) {
	ganttFormCallback = new(GanttFormCallback)
	ganttFormCallback.probe = probe
	ganttFormCallback.gantt = gantt
	ganttFormCallback.formGroup = formGroup

	ganttFormCallback.CreationMode = (gantt == nil)

	return
}

type GanttFormCallback struct {
	gantt *models.Gantt

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (ganttFormCallback *GanttFormCallback) OnSave() {

	log.Println("GanttFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	ganttFormCallback.probe.formStage.Checkout()

	if ganttFormCallback.gantt == nil {
		ganttFormCallback.gantt = new(models.Gantt).Stage(ganttFormCallback.probe.stageOfInterest)
	}
	gantt_ := ganttFormCallback.gantt
	_ = gantt_

	for _, formDiv := range ganttFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(gantt_.Name), formDiv)
		case "ComputedStart":
			FormDivBasicFieldToField(&(gantt_.ComputedStart), formDiv)
		case "ComputedEnd":
			FormDivBasicFieldToField(&(gantt_.ComputedEnd), formDiv)
		case "ComputedDuration":
			FormDivBasicFieldToField(&(gantt_.ComputedDuration), formDiv)
		case "UseManualStartAndEndDates":
			FormDivBasicFieldToField(&(gantt_.UseManualStartAndEndDates), formDiv)
		case "ManualStart":
			FormDivBasicFieldToField(&(gantt_.ManualStart), formDiv)
		case "ManualEnd":
			FormDivBasicFieldToField(&(gantt_.ManualEnd), formDiv)
		case "LaneHeight":
			FormDivBasicFieldToField(&(gantt_.LaneHeight), formDiv)
		case "RatioBarToLaneHeight":
			FormDivBasicFieldToField(&(gantt_.RatioBarToLaneHeight), formDiv)
		case "YTopMargin":
			FormDivBasicFieldToField(&(gantt_.YTopMargin), formDiv)
		case "XLeftText":
			FormDivBasicFieldToField(&(gantt_.XLeftText), formDiv)
		case "TextHeight":
			FormDivBasicFieldToField(&(gantt_.TextHeight), formDiv)
		case "XLeftLanes":
			FormDivBasicFieldToField(&(gantt_.XLeftLanes), formDiv)
		case "XRightMargin":
			FormDivBasicFieldToField(&(gantt_.XRightMargin), formDiv)
		case "ArrowLengthToTheRightOfStartBar":
			FormDivBasicFieldToField(&(gantt_.ArrowLengthToTheRightOfStartBar), formDiv)
		case "ArrowTipLenght":
			FormDivBasicFieldToField(&(gantt_.ArrowTipLenght), formDiv)
		case "TimeLine_Color":
			FormDivBasicFieldToField(&(gantt_.TimeLine_Color), formDiv)
		case "TimeLine_FillOpacity":
			FormDivBasicFieldToField(&(gantt_.TimeLine_FillOpacity), formDiv)
		case "TimeLine_Stroke":
			FormDivBasicFieldToField(&(gantt_.TimeLine_Stroke), formDiv)
		case "TimeLine_StrokeWidth":
			FormDivBasicFieldToField(&(gantt_.TimeLine_StrokeWidth), formDiv)
		case "Group_Stroke":
			FormDivBasicFieldToField(&(gantt_.Group_Stroke), formDiv)
		case "Group_StrokeWidth":
			FormDivBasicFieldToField(&(gantt_.Group_StrokeWidth), formDiv)
		case "Group_StrokeDashArray":
			FormDivBasicFieldToField(&(gantt_.Group_StrokeDashArray), formDiv)
		case "DateYOffset":
			FormDivBasicFieldToField(&(gantt_.DateYOffset), formDiv)
		case "AlignOnStartEndOnYearStart":
			FormDivBasicFieldToField(&(gantt_.AlignOnStartEndOnYearStart), formDiv)
		}
	}

	// manage the suppress operation
	if ganttFormCallback.formGroup.HasSuppressButtonBeenPressed {
		gantt_.Unstage(ganttFormCallback.probe.stageOfInterest)
	}

	ganttFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Gantt](
		ganttFormCallback.probe,
	)
	ganttFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if ganttFormCallback.CreationMode || ganttFormCallback.formGroup.HasSuppressButtonBeenPressed {
		ganttFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(ganttFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__GanttFormCallback(
			nil,
			ganttFormCallback.probe,
			newFormGroup,
		)
		gantt := new(models.Gantt)
		FillUpForm(gantt, newFormGroup, ganttFormCallback.probe)
		ganttFormCallback.probe.formStage.Commit()
	}

	fillUpTree(ganttFormCallback.probe)
}
func __gong__New__GroupFormCallback(
	group *models.Group,
	probe *Probe,
	formGroup *table.FormGroup,
) (groupFormCallback *GroupFormCallback) {
	groupFormCallback = new(GroupFormCallback)
	groupFormCallback.probe = probe
	groupFormCallback.group = group
	groupFormCallback.formGroup = formGroup

	groupFormCallback.CreationMode = (group == nil)

	return
}

type GroupFormCallback struct {
	group *models.Group

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (groupFormCallback *GroupFormCallback) OnSave() {

	log.Println("GroupFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	groupFormCallback.probe.formStage.Checkout()

	if groupFormCallback.group == nil {
		groupFormCallback.group = new(models.Group).Stage(groupFormCallback.probe.stageOfInterest)
	}
	group_ := groupFormCallback.group
	_ = group_

	for _, formDiv := range groupFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(group_.Name), formDiv)
		case "Gantt:Groups":
			// we need to retrieve the field owner before the change
			var pastGanttOwner *models.Gantt
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Gantt"
			rf.Fieldname = "Groups"
			reverseFieldOwner := models.GetReverseFieldOwner(
				groupFormCallback.probe.stageOfInterest,
				group_,
				&rf)

			if reverseFieldOwner != nil {
				pastGanttOwner = reverseFieldOwner.(*models.Gantt)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastGanttOwner != nil {
					idx := slices.Index(pastGanttOwner.Groups, group_)
					pastGanttOwner.Groups = slices.Delete(pastGanttOwner.Groups, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _gantt := range *models.GetGongstructInstancesSet[models.Gantt](groupFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _gantt.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newGanttOwner := _gantt // we have a match
						if pastGanttOwner != nil {
							if newGanttOwner != pastGanttOwner {
								idx := slices.Index(pastGanttOwner.Groups, group_)
								pastGanttOwner.Groups = slices.Delete(pastGanttOwner.Groups, idx, idx+1)
								newGanttOwner.Groups = append(newGanttOwner.Groups, group_)
							}
						} else {
							newGanttOwner.Groups = append(newGanttOwner.Groups, group_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if groupFormCallback.formGroup.HasSuppressButtonBeenPressed {
		group_.Unstage(groupFormCallback.probe.stageOfInterest)
	}

	groupFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Group](
		groupFormCallback.probe,
	)
	groupFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if groupFormCallback.CreationMode || groupFormCallback.formGroup.HasSuppressButtonBeenPressed {
		groupFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(groupFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__GroupFormCallback(
			nil,
			groupFormCallback.probe,
			newFormGroup,
		)
		group := new(models.Group)
		FillUpForm(group, newFormGroup, groupFormCallback.probe)
		groupFormCallback.probe.formStage.Commit()
	}

	fillUpTree(groupFormCallback.probe)
}
func __gong__New__LaneFormCallback(
	lane *models.Lane,
	probe *Probe,
	formGroup *table.FormGroup,
) (laneFormCallback *LaneFormCallback) {
	laneFormCallback = new(LaneFormCallback)
	laneFormCallback.probe = probe
	laneFormCallback.lane = lane
	laneFormCallback.formGroup = formGroup

	laneFormCallback.CreationMode = (lane == nil)

	return
}

type LaneFormCallback struct {
	lane *models.Lane

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (laneFormCallback *LaneFormCallback) OnSave() {

	log.Println("LaneFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	laneFormCallback.probe.formStage.Checkout()

	if laneFormCallback.lane == nil {
		laneFormCallback.lane = new(models.Lane).Stage(laneFormCallback.probe.stageOfInterest)
	}
	lane_ := laneFormCallback.lane
	_ = lane_

	for _, formDiv := range laneFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(lane_.Name), formDiv)
		case "Order":
			FormDivBasicFieldToField(&(lane_.Order), formDiv)
		case "Gantt:Lanes":
			// we need to retrieve the field owner before the change
			var pastGanttOwner *models.Gantt
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Gantt"
			rf.Fieldname = "Lanes"
			reverseFieldOwner := models.GetReverseFieldOwner(
				laneFormCallback.probe.stageOfInterest,
				lane_,
				&rf)

			if reverseFieldOwner != nil {
				pastGanttOwner = reverseFieldOwner.(*models.Gantt)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastGanttOwner != nil {
					idx := slices.Index(pastGanttOwner.Lanes, lane_)
					pastGanttOwner.Lanes = slices.Delete(pastGanttOwner.Lanes, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _gantt := range *models.GetGongstructInstancesSet[models.Gantt](laneFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _gantt.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newGanttOwner := _gantt // we have a match
						if pastGanttOwner != nil {
							if newGanttOwner != pastGanttOwner {
								idx := slices.Index(pastGanttOwner.Lanes, lane_)
								pastGanttOwner.Lanes = slices.Delete(pastGanttOwner.Lanes, idx, idx+1)
								newGanttOwner.Lanes = append(newGanttOwner.Lanes, lane_)
							}
						} else {
							newGanttOwner.Lanes = append(newGanttOwner.Lanes, lane_)
						}
					}
				}
			}
		case "Group:GroupLanes":
			// we need to retrieve the field owner before the change
			var pastGroupOwner *models.Group
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Group"
			rf.Fieldname = "GroupLanes"
			reverseFieldOwner := models.GetReverseFieldOwner(
				laneFormCallback.probe.stageOfInterest,
				lane_,
				&rf)

			if reverseFieldOwner != nil {
				pastGroupOwner = reverseFieldOwner.(*models.Group)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastGroupOwner != nil {
					idx := slices.Index(pastGroupOwner.GroupLanes, lane_)
					pastGroupOwner.GroupLanes = slices.Delete(pastGroupOwner.GroupLanes, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _group := range *models.GetGongstructInstancesSet[models.Group](laneFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _group.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newGroupOwner := _group // we have a match
						if pastGroupOwner != nil {
							if newGroupOwner != pastGroupOwner {
								idx := slices.Index(pastGroupOwner.GroupLanes, lane_)
								pastGroupOwner.GroupLanes = slices.Delete(pastGroupOwner.GroupLanes, idx, idx+1)
								newGroupOwner.GroupLanes = append(newGroupOwner.GroupLanes, lane_)
							}
						} else {
							newGroupOwner.GroupLanes = append(newGroupOwner.GroupLanes, lane_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if laneFormCallback.formGroup.HasSuppressButtonBeenPressed {
		lane_.Unstage(laneFormCallback.probe.stageOfInterest)
	}

	laneFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Lane](
		laneFormCallback.probe,
	)
	laneFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if laneFormCallback.CreationMode || laneFormCallback.formGroup.HasSuppressButtonBeenPressed {
		laneFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(laneFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__LaneFormCallback(
			nil,
			laneFormCallback.probe,
			newFormGroup,
		)
		lane := new(models.Lane)
		FillUpForm(lane, newFormGroup, laneFormCallback.probe)
		laneFormCallback.probe.formStage.Commit()
	}

	fillUpTree(laneFormCallback.probe)
}
func __gong__New__LaneUseFormCallback(
	laneuse *models.LaneUse,
	probe *Probe,
	formGroup *table.FormGroup,
) (laneuseFormCallback *LaneUseFormCallback) {
	laneuseFormCallback = new(LaneUseFormCallback)
	laneuseFormCallback.probe = probe
	laneuseFormCallback.laneuse = laneuse
	laneuseFormCallback.formGroup = formGroup

	laneuseFormCallback.CreationMode = (laneuse == nil)

	return
}

type LaneUseFormCallback struct {
	laneuse *models.LaneUse

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (laneuseFormCallback *LaneUseFormCallback) OnSave() {

	log.Println("LaneUseFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	laneuseFormCallback.probe.formStage.Checkout()

	if laneuseFormCallback.laneuse == nil {
		laneuseFormCallback.laneuse = new(models.LaneUse).Stage(laneuseFormCallback.probe.stageOfInterest)
	}
	laneuse_ := laneuseFormCallback.laneuse
	_ = laneuse_

	for _, formDiv := range laneuseFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(laneuse_.Name), formDiv)
		case "Lane":
			FormDivSelectFieldToField(&(laneuse_.Lane), laneuseFormCallback.probe.stageOfInterest, formDiv)
		case "Milestone:LanesToDisplayMilestoneUse":
			// we need to retrieve the field owner before the change
			var pastMilestoneOwner *models.Milestone
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Milestone"
			rf.Fieldname = "LanesToDisplayMilestoneUse"
			reverseFieldOwner := models.GetReverseFieldOwner(
				laneuseFormCallback.probe.stageOfInterest,
				laneuse_,
				&rf)

			if reverseFieldOwner != nil {
				pastMilestoneOwner = reverseFieldOwner.(*models.Milestone)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastMilestoneOwner != nil {
					idx := slices.Index(pastMilestoneOwner.LanesToDisplayMilestoneUse, laneuse_)
					pastMilestoneOwner.LanesToDisplayMilestoneUse = slices.Delete(pastMilestoneOwner.LanesToDisplayMilestoneUse, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _milestone := range *models.GetGongstructInstancesSet[models.Milestone](laneuseFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _milestone.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newMilestoneOwner := _milestone // we have a match
						if pastMilestoneOwner != nil {
							if newMilestoneOwner != pastMilestoneOwner {
								idx := slices.Index(pastMilestoneOwner.LanesToDisplayMilestoneUse, laneuse_)
								pastMilestoneOwner.LanesToDisplayMilestoneUse = slices.Delete(pastMilestoneOwner.LanesToDisplayMilestoneUse, idx, idx+1)
								newMilestoneOwner.LanesToDisplayMilestoneUse = append(newMilestoneOwner.LanesToDisplayMilestoneUse, laneuse_)
							}
						} else {
							newMilestoneOwner.LanesToDisplayMilestoneUse = append(newMilestoneOwner.LanesToDisplayMilestoneUse, laneuse_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if laneuseFormCallback.formGroup.HasSuppressButtonBeenPressed {
		laneuse_.Unstage(laneuseFormCallback.probe.stageOfInterest)
	}

	laneuseFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.LaneUse](
		laneuseFormCallback.probe,
	)
	laneuseFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if laneuseFormCallback.CreationMode || laneuseFormCallback.formGroup.HasSuppressButtonBeenPressed {
		laneuseFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(laneuseFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__LaneUseFormCallback(
			nil,
			laneuseFormCallback.probe,
			newFormGroup,
		)
		laneuse := new(models.LaneUse)
		FillUpForm(laneuse, newFormGroup, laneuseFormCallback.probe)
		laneuseFormCallback.probe.formStage.Commit()
	}

	fillUpTree(laneuseFormCallback.probe)
}
func __gong__New__MilestoneFormCallback(
	milestone *models.Milestone,
	probe *Probe,
	formGroup *table.FormGroup,
) (milestoneFormCallback *MilestoneFormCallback) {
	milestoneFormCallback = new(MilestoneFormCallback)
	milestoneFormCallback.probe = probe
	milestoneFormCallback.milestone = milestone
	milestoneFormCallback.formGroup = formGroup

	milestoneFormCallback.CreationMode = (milestone == nil)

	return
}

type MilestoneFormCallback struct {
	milestone *models.Milestone

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (milestoneFormCallback *MilestoneFormCallback) OnSave() {

	log.Println("MilestoneFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	milestoneFormCallback.probe.formStage.Checkout()

	if milestoneFormCallback.milestone == nil {
		milestoneFormCallback.milestone = new(models.Milestone).Stage(milestoneFormCallback.probe.stageOfInterest)
	}
	milestone_ := milestoneFormCallback.milestone
	_ = milestone_

	for _, formDiv := range milestoneFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(milestone_.Name), formDiv)
		case "Date":
			FormDivBasicFieldToField(&(milestone_.Date), formDiv)
		case "DisplayVerticalBar":
			FormDivBasicFieldToField(&(milestone_.DisplayVerticalBar), formDiv)
		case "Gantt:Milestones":
			// we need to retrieve the field owner before the change
			var pastGanttOwner *models.Gantt
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Gantt"
			rf.Fieldname = "Milestones"
			reverseFieldOwner := models.GetReverseFieldOwner(
				milestoneFormCallback.probe.stageOfInterest,
				milestone_,
				&rf)

			if reverseFieldOwner != nil {
				pastGanttOwner = reverseFieldOwner.(*models.Gantt)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastGanttOwner != nil {
					idx := slices.Index(pastGanttOwner.Milestones, milestone_)
					pastGanttOwner.Milestones = slices.Delete(pastGanttOwner.Milestones, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _gantt := range *models.GetGongstructInstancesSet[models.Gantt](milestoneFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _gantt.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newGanttOwner := _gantt // we have a match
						if pastGanttOwner != nil {
							if newGanttOwner != pastGanttOwner {
								idx := slices.Index(pastGanttOwner.Milestones, milestone_)
								pastGanttOwner.Milestones = slices.Delete(pastGanttOwner.Milestones, idx, idx+1)
								newGanttOwner.Milestones = append(newGanttOwner.Milestones, milestone_)
							}
						} else {
							newGanttOwner.Milestones = append(newGanttOwner.Milestones, milestone_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if milestoneFormCallback.formGroup.HasSuppressButtonBeenPressed {
		milestone_.Unstage(milestoneFormCallback.probe.stageOfInterest)
	}

	milestoneFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Milestone](
		milestoneFormCallback.probe,
	)
	milestoneFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if milestoneFormCallback.CreationMode || milestoneFormCallback.formGroup.HasSuppressButtonBeenPressed {
		milestoneFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(milestoneFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__MilestoneFormCallback(
			nil,
			milestoneFormCallback.probe,
			newFormGroup,
		)
		milestone := new(models.Milestone)
		FillUpForm(milestone, newFormGroup, milestoneFormCallback.probe)
		milestoneFormCallback.probe.formStage.Commit()
	}

	fillUpTree(milestoneFormCallback.probe)
}
