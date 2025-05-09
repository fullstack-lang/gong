// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	table "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gong/lib/gantt/go/models"
)

// code to avoid error when generated code does not need to import packages
const __dummmy__time = time.Nanosecond

var _ = __dummmy__time

var __dummmy__letters = slices.Delete([]string{"a"}, 0, 1)

var _ = __dummmy__letters

const __dummy__log = log.Ldate

var _ = __dummy__log

// end of code to avoid error when generated code does not need to import packages

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

	// log.Println("ArrowFormCallback, OnSave")

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
			// WARNING : this form deals with the N-N association "Gantt.Arrows []*Arrow" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Arrow). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Gantt
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Gantt"
				rf.Fieldname = "Arrows"
				formerAssociationSource := models.GetReverseFieldOwner(
					arrowFormCallback.probe.stageOfInterest,
					arrow_,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Gantt)
					if !ok {
						log.Fatalln("Source of Gantt.Arrows []*Arrow, is not an Gantt instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				if formerSource != nil {
					idx := slices.Index(formerSource.Arrows, arrow_)
					formerSource.Arrows = slices.Delete(formerSource.Arrows, idx, idx+1)
				}
				// That could mean we clear the assocation for all source instances
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.Gantt
			for _gantt := range *models.GetGongstructInstancesSet[models.Gantt](arrowFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _gantt.GetName() == newSourceName.GetName() {
					newSource = _gantt // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Gantt.Arrows []*Arrow, with name", newSourceName, ", does not exist")
				break
			}

			// append the value to the new source field
			newSource.Arrows = append(newSource.Arrows, arrow_)
		}
	}

	// manage the suppress operation
	if arrowFormCallback.formGroup.HasSuppressButtonBeenPressed {
		arrow_.Unstage(arrowFormCallback.probe.stageOfInterest)
	}

	arrowFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.Arrow](
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

	updateAndCommitTree(arrowFormCallback.probe)
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

	// log.Println("BarFormCallback, OnSave")

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
			// WARNING : this form deals with the N-N association "Lane.Bars []*Bar" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Bar). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Lane
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Lane"
				rf.Fieldname = "Bars"
				formerAssociationSource := models.GetReverseFieldOwner(
					barFormCallback.probe.stageOfInterest,
					bar_,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Lane)
					if !ok {
						log.Fatalln("Source of Lane.Bars []*Bar, is not an Lane instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				if formerSource != nil {
					idx := slices.Index(formerSource.Bars, bar_)
					formerSource.Bars = slices.Delete(formerSource.Bars, idx, idx+1)
				}
				// That could mean we clear the assocation for all source instances
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.Lane
			for _lane := range *models.GetGongstructInstancesSet[models.Lane](barFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _lane.GetName() == newSourceName.GetName() {
					newSource = _lane // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Lane.Bars []*Bar, with name", newSourceName, ", does not exist")
				break
			}

			// append the value to the new source field
			newSource.Bars = append(newSource.Bars, bar_)
		}
	}

	// manage the suppress operation
	if barFormCallback.formGroup.HasSuppressButtonBeenPressed {
		bar_.Unstage(barFormCallback.probe.stageOfInterest)
	}

	barFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.Bar](
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

	updateAndCommitTree(barFormCallback.probe)
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

	// log.Println("GanttFormCallback, OnSave")

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
	updateAndCommitTable[models.Gantt](
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

	updateAndCommitTree(ganttFormCallback.probe)
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

	// log.Println("GroupFormCallback, OnSave")

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
			// WARNING : this form deals with the N-N association "Gantt.Groups []*Group" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Group). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Gantt
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Gantt"
				rf.Fieldname = "Groups"
				formerAssociationSource := models.GetReverseFieldOwner(
					groupFormCallback.probe.stageOfInterest,
					group_,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Gantt)
					if !ok {
						log.Fatalln("Source of Gantt.Groups []*Group, is not an Gantt instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				if formerSource != nil {
					idx := slices.Index(formerSource.Groups, group_)
					formerSource.Groups = slices.Delete(formerSource.Groups, idx, idx+1)
				}
				// That could mean we clear the assocation for all source instances
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.Gantt
			for _gantt := range *models.GetGongstructInstancesSet[models.Gantt](groupFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _gantt.GetName() == newSourceName.GetName() {
					newSource = _gantt // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Gantt.Groups []*Group, with name", newSourceName, ", does not exist")
				break
			}

			// append the value to the new source field
			newSource.Groups = append(newSource.Groups, group_)
		}
	}

	// manage the suppress operation
	if groupFormCallback.formGroup.HasSuppressButtonBeenPressed {
		group_.Unstage(groupFormCallback.probe.stageOfInterest)
	}

	groupFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.Group](
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

	updateAndCommitTree(groupFormCallback.probe)
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

	// log.Println("LaneFormCallback, OnSave")

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
			// WARNING : this form deals with the N-N association "Gantt.Lanes []*Lane" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Lane). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Gantt
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Gantt"
				rf.Fieldname = "Lanes"
				formerAssociationSource := models.GetReverseFieldOwner(
					laneFormCallback.probe.stageOfInterest,
					lane_,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Gantt)
					if !ok {
						log.Fatalln("Source of Gantt.Lanes []*Lane, is not an Gantt instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				if formerSource != nil {
					idx := slices.Index(formerSource.Lanes, lane_)
					formerSource.Lanes = slices.Delete(formerSource.Lanes, idx, idx+1)
				}
				// That could mean we clear the assocation for all source instances
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.Gantt
			for _gantt := range *models.GetGongstructInstancesSet[models.Gantt](laneFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _gantt.GetName() == newSourceName.GetName() {
					newSource = _gantt // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Gantt.Lanes []*Lane, with name", newSourceName, ", does not exist")
				break
			}

			// append the value to the new source field
			newSource.Lanes = append(newSource.Lanes, lane_)
		case "Group:GroupLanes":
			// WARNING : this form deals with the N-N association "Group.GroupLanes []*Lane" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Lane). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Group
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Group"
				rf.Fieldname = "GroupLanes"
				formerAssociationSource := models.GetReverseFieldOwner(
					laneFormCallback.probe.stageOfInterest,
					lane_,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Group)
					if !ok {
						log.Fatalln("Source of Group.GroupLanes []*Lane, is not an Group instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				if formerSource != nil {
					idx := slices.Index(formerSource.GroupLanes, lane_)
					formerSource.GroupLanes = slices.Delete(formerSource.GroupLanes, idx, idx+1)
				}
				// That could mean we clear the assocation for all source instances
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.Group
			for _group := range *models.GetGongstructInstancesSet[models.Group](laneFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _group.GetName() == newSourceName.GetName() {
					newSource = _group // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Group.GroupLanes []*Lane, with name", newSourceName, ", does not exist")
				break
			}

			// append the value to the new source field
			newSource.GroupLanes = append(newSource.GroupLanes, lane_)
		}
	}

	// manage the suppress operation
	if laneFormCallback.formGroup.HasSuppressButtonBeenPressed {
		lane_.Unstage(laneFormCallback.probe.stageOfInterest)
	}

	laneFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.Lane](
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

	updateAndCommitTree(laneFormCallback.probe)
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

	// log.Println("LaneUseFormCallback, OnSave")

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
			// WARNING : this form deals with the N-N association "Milestone.LanesToDisplayMilestoneUse []*LaneUse" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of LaneUse). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Milestone
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Milestone"
				rf.Fieldname = "LanesToDisplayMilestoneUse"
				formerAssociationSource := models.GetReverseFieldOwner(
					laneuseFormCallback.probe.stageOfInterest,
					laneuse_,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Milestone)
					if !ok {
						log.Fatalln("Source of Milestone.LanesToDisplayMilestoneUse []*LaneUse, is not an Milestone instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				if formerSource != nil {
					idx := slices.Index(formerSource.LanesToDisplayMilestoneUse, laneuse_)
					formerSource.LanesToDisplayMilestoneUse = slices.Delete(formerSource.LanesToDisplayMilestoneUse, idx, idx+1)
				}
				// That could mean we clear the assocation for all source instances
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.Milestone
			for _milestone := range *models.GetGongstructInstancesSet[models.Milestone](laneuseFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _milestone.GetName() == newSourceName.GetName() {
					newSource = _milestone // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Milestone.LanesToDisplayMilestoneUse []*LaneUse, with name", newSourceName, ", does not exist")
				break
			}

			// append the value to the new source field
			newSource.LanesToDisplayMilestoneUse = append(newSource.LanesToDisplayMilestoneUse, laneuse_)
		}
	}

	// manage the suppress operation
	if laneuseFormCallback.formGroup.HasSuppressButtonBeenPressed {
		laneuse_.Unstage(laneuseFormCallback.probe.stageOfInterest)
	}

	laneuseFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.LaneUse](
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

	updateAndCommitTree(laneuseFormCallback.probe)
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

	// log.Println("MilestoneFormCallback, OnSave")

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
			// WARNING : this form deals with the N-N association "Gantt.Milestones []*Milestone" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Milestone). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Gantt
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Gantt"
				rf.Fieldname = "Milestones"
				formerAssociationSource := models.GetReverseFieldOwner(
					milestoneFormCallback.probe.stageOfInterest,
					milestone_,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Gantt)
					if !ok {
						log.Fatalln("Source of Gantt.Milestones []*Milestone, is not an Gantt instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				if formerSource != nil {
					idx := slices.Index(formerSource.Milestones, milestone_)
					formerSource.Milestones = slices.Delete(formerSource.Milestones, idx, idx+1)
				}
				// That could mean we clear the assocation for all source instances
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.Gantt
			for _gantt := range *models.GetGongstructInstancesSet[models.Gantt](milestoneFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _gantt.GetName() == newSourceName.GetName() {
					newSource = _gantt // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Gantt.Milestones []*Milestone, with name", newSourceName, ", does not exist")
				break
			}

			// append the value to the new source field
			newSource.Milestones = append(newSource.Milestones, milestone_)
		}
	}

	// manage the suppress operation
	if milestoneFormCallback.formGroup.HasSuppressButtonBeenPressed {
		milestone_.Unstage(milestoneFormCallback.probe.stageOfInterest)
	}

	milestoneFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.Milestone](
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

	updateAndCommitTree(milestoneFormCallback.probe)
}
