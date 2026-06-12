// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/lib/gantt/go/models"
)

// to avoid errors when time and slices packages are not used in the generated code
const _ = time.Nanosecond

var _ = slices.Delete([]string{"a"}, 0, 1)

var _ = log.Panicf

// insertion point
func __gong__New__ArrowFormCallback(
	arrow *models.Arrow,
	probe *Probe,
	formGroup *form.FormGroup,
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

	formGroup *form.FormGroup
}

func (arrowFormCallback *ArrowFormCallback) OnSave() {
	arrowFormCallback.probe.stageOfInterest.Lock()
	defer arrowFormCallback.probe.stageOfInterest.Unlock()

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
			// 1. Decode the AssociationStorage which contains the rowIDs of the Gantt instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Gantt instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Gantt](arrowFormCallback.probe.stageOfInterest)
			targetGanttIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetGanttIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Gantt instances and update their Arrows slice
			for _gantt := range *models.GetGongstructInstancesSetFromPointerType[*models.Gantt](arrowFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(arrowFormCallback.probe.stageOfInterest, _gantt)
				
				// if Gantt is selected
				if targetGanttIDs[id] {
					// ensure arrow_ is in _gantt.Arrows
					found := false
					for _, _b := range _gantt.Arrows {
						if _b == arrow_ {
							found = true
							break
						}
					}
					if !found {
						_gantt.Arrows = append(_gantt.Arrows, arrow_)
						arrowFormCallback.probe.UpdateSliceOfPointersCallback(_gantt, "Arrows", &_gantt.Arrows)
					}
				} else {
					// ensure arrow_ is NOT in _gantt.Arrows
					idx := slices.Index(_gantt.Arrows, arrow_)
					if idx != -1 {
						_gantt.Arrows = slices.Delete(_gantt.Arrows, idx, idx+1)
						arrowFormCallback.probe.UpdateSliceOfPointersCallback(_gantt, "Arrows", &_gantt.Arrows)
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
	updateProbeTable[*models.Arrow](
		arrowFormCallback.probe,
	)

	// display a new form by reset the form stage
	if arrowFormCallback.CreationMode || arrowFormCallback.formGroup.HasSuppressButtonBeenPressed {
		arrowFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
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

	arrowFormCallback.probe.ux_tree()
}
func __gong__New__BarFormCallback(
	bar *models.Bar,
	probe *Probe,
	formGroup *form.FormGroup,
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

	formGroup *form.FormGroup
}

func (barFormCallback *BarFormCallback) OnSave() {
	barFormCallback.probe.stageOfInterest.Lock()
	defer barFormCallback.probe.stageOfInterest.Unlock()

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
			// 1. Decode the AssociationStorage which contains the rowIDs of the Lane instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Lane instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Lane](barFormCallback.probe.stageOfInterest)
			targetLaneIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetLaneIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Lane instances and update their Bars slice
			for _lane := range *models.GetGongstructInstancesSetFromPointerType[*models.Lane](barFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(barFormCallback.probe.stageOfInterest, _lane)
				
				// if Lane is selected
				if targetLaneIDs[id] {
					// ensure bar_ is in _lane.Bars
					found := false
					for _, _b := range _lane.Bars {
						if _b == bar_ {
							found = true
							break
						}
					}
					if !found {
						_lane.Bars = append(_lane.Bars, bar_)
						barFormCallback.probe.UpdateSliceOfPointersCallback(_lane, "Bars", &_lane.Bars)
					}
				} else {
					// ensure bar_ is NOT in _lane.Bars
					idx := slices.Index(_lane.Bars, bar_)
					if idx != -1 {
						_lane.Bars = slices.Delete(_lane.Bars, idx, idx+1)
						barFormCallback.probe.UpdateSliceOfPointersCallback(_lane, "Bars", &_lane.Bars)
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
	updateProbeTable[*models.Bar](
		barFormCallback.probe,
	)

	// display a new form by reset the form stage
	if barFormCallback.CreationMode || barFormCallback.formGroup.HasSuppressButtonBeenPressed {
		barFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
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

	barFormCallback.probe.ux_tree()
}
func __gong__New__GanttFormCallback(
	gantt *models.Gantt,
	probe *Probe,
	formGroup *form.FormGroup,
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

	formGroup *form.FormGroup
}

func (ganttFormCallback *GanttFormCallback) OnSave() {
	ganttFormCallback.probe.stageOfInterest.Lock()
	defer ganttFormCallback.probe.stageOfInterest.Unlock()

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
		case "Lanes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Lane](ganttFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Lane, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Lane)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					ganttFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Lane](ganttFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			gantt_.Lanes = instanceSlice
			ganttFormCallback.probe.UpdateSliceOfPointersCallback(gantt_, "Lanes", &gantt_.Lanes)

		case "Milestones":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Milestone](ganttFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Milestone, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Milestone)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					ganttFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Milestone](ganttFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			gantt_.Milestones = instanceSlice
			ganttFormCallback.probe.UpdateSliceOfPointersCallback(gantt_, "Milestones", &gantt_.Milestones)

		case "Groups":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Group](ganttFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Group, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Group)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					ganttFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Group](ganttFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			gantt_.Groups = instanceSlice
			ganttFormCallback.probe.UpdateSliceOfPointersCallback(gantt_, "Groups", &gantt_.Groups)

		case "Arrows":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Arrow](ganttFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Arrow, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Arrow)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					ganttFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Arrow](ganttFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			gantt_.Arrows = instanceSlice
			ganttFormCallback.probe.UpdateSliceOfPointersCallback(gantt_, "Arrows", &gantt_.Arrows)

		}
	}

	// manage the suppress operation
	if ganttFormCallback.formGroup.HasSuppressButtonBeenPressed {
		gantt_.Unstage(ganttFormCallback.probe.stageOfInterest)
	}

	ganttFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Gantt](
		ganttFormCallback.probe,
	)

	// display a new form by reset the form stage
	if ganttFormCallback.CreationMode || ganttFormCallback.formGroup.HasSuppressButtonBeenPressed {
		ganttFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
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

	ganttFormCallback.probe.ux_tree()
}
func __gong__New__GroupFormCallback(
	group *models.Group,
	probe *Probe,
	formGroup *form.FormGroup,
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

	formGroup *form.FormGroup
}

func (groupFormCallback *GroupFormCallback) OnSave() {
	groupFormCallback.probe.stageOfInterest.Lock()
	defer groupFormCallback.probe.stageOfInterest.Unlock()

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
		case "GroupLanes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Lane](groupFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Lane, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Lane)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					groupFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Lane](groupFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			group_.GroupLanes = instanceSlice
			groupFormCallback.probe.UpdateSliceOfPointersCallback(group_, "GroupLanes", &group_.GroupLanes)

		case "Gantt:Groups":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Gantt instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Gantt instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Gantt](groupFormCallback.probe.stageOfInterest)
			targetGanttIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetGanttIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Gantt instances and update their Groups slice
			for _gantt := range *models.GetGongstructInstancesSetFromPointerType[*models.Gantt](groupFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(groupFormCallback.probe.stageOfInterest, _gantt)
				
				// if Gantt is selected
				if targetGanttIDs[id] {
					// ensure group_ is in _gantt.Groups
					found := false
					for _, _b := range _gantt.Groups {
						if _b == group_ {
							found = true
							break
						}
					}
					if !found {
						_gantt.Groups = append(_gantt.Groups, group_)
						groupFormCallback.probe.UpdateSliceOfPointersCallback(_gantt, "Groups", &_gantt.Groups)
					}
				} else {
					// ensure group_ is NOT in _gantt.Groups
					idx := slices.Index(_gantt.Groups, group_)
					if idx != -1 {
						_gantt.Groups = slices.Delete(_gantt.Groups, idx, idx+1)
						groupFormCallback.probe.UpdateSliceOfPointersCallback(_gantt, "Groups", &_gantt.Groups)
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
	updateProbeTable[*models.Group](
		groupFormCallback.probe,
	)

	// display a new form by reset the form stage
	if groupFormCallback.CreationMode || groupFormCallback.formGroup.HasSuppressButtonBeenPressed {
		groupFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
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

	groupFormCallback.probe.ux_tree()
}
func __gong__New__LaneFormCallback(
	lane *models.Lane,
	probe *Probe,
	formGroup *form.FormGroup,
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

	formGroup *form.FormGroup
}

func (laneFormCallback *LaneFormCallback) OnSave() {
	laneFormCallback.probe.stageOfInterest.Lock()
	defer laneFormCallback.probe.stageOfInterest.Unlock()

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
		case "Bars":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Bar](laneFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Bar, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Bar)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					laneFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Bar](laneFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			lane_.Bars = instanceSlice
			laneFormCallback.probe.UpdateSliceOfPointersCallback(lane_, "Bars", &lane_.Bars)

		case "Gantt:Lanes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Gantt instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Gantt instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Gantt](laneFormCallback.probe.stageOfInterest)
			targetGanttIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetGanttIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Gantt instances and update their Lanes slice
			for _gantt := range *models.GetGongstructInstancesSetFromPointerType[*models.Gantt](laneFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(laneFormCallback.probe.stageOfInterest, _gantt)
				
				// if Gantt is selected
				if targetGanttIDs[id] {
					// ensure lane_ is in _gantt.Lanes
					found := false
					for _, _b := range _gantt.Lanes {
						if _b == lane_ {
							found = true
							break
						}
					}
					if !found {
						_gantt.Lanes = append(_gantt.Lanes, lane_)
						laneFormCallback.probe.UpdateSliceOfPointersCallback(_gantt, "Lanes", &_gantt.Lanes)
					}
				} else {
					// ensure lane_ is NOT in _gantt.Lanes
					idx := slices.Index(_gantt.Lanes, lane_)
					if idx != -1 {
						_gantt.Lanes = slices.Delete(_gantt.Lanes, idx, idx+1)
						laneFormCallback.probe.UpdateSliceOfPointersCallback(_gantt, "Lanes", &_gantt.Lanes)
					}
				}
			}
		case "Group:GroupLanes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Group instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Group instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Group](laneFormCallback.probe.stageOfInterest)
			targetGroupIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetGroupIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Group instances and update their GroupLanes slice
			for _group := range *models.GetGongstructInstancesSetFromPointerType[*models.Group](laneFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(laneFormCallback.probe.stageOfInterest, _group)
				
				// if Group is selected
				if targetGroupIDs[id] {
					// ensure lane_ is in _group.GroupLanes
					found := false
					for _, _b := range _group.GroupLanes {
						if _b == lane_ {
							found = true
							break
						}
					}
					if !found {
						_group.GroupLanes = append(_group.GroupLanes, lane_)
						laneFormCallback.probe.UpdateSliceOfPointersCallback(_group, "GroupLanes", &_group.GroupLanes)
					}
				} else {
					// ensure lane_ is NOT in _group.GroupLanes
					idx := slices.Index(_group.GroupLanes, lane_)
					if idx != -1 {
						_group.GroupLanes = slices.Delete(_group.GroupLanes, idx, idx+1)
						laneFormCallback.probe.UpdateSliceOfPointersCallback(_group, "GroupLanes", &_group.GroupLanes)
					}
				}
			}
		case "Milestone:LanesToDisplay":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Milestone instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Milestone instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Milestone](laneFormCallback.probe.stageOfInterest)
			targetMilestoneIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetMilestoneIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Milestone instances and update their LanesToDisplay slice
			for _milestone := range *models.GetGongstructInstancesSetFromPointerType[*models.Milestone](laneFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(laneFormCallback.probe.stageOfInterest, _milestone)
				
				// if Milestone is selected
				if targetMilestoneIDs[id] {
					// ensure lane_ is in _milestone.LanesToDisplay
					found := false
					for _, _b := range _milestone.LanesToDisplay {
						if _b == lane_ {
							found = true
							break
						}
					}
					if !found {
						_milestone.LanesToDisplay = append(_milestone.LanesToDisplay, lane_)
						laneFormCallback.probe.UpdateSliceOfPointersCallback(_milestone, "LanesToDisplay", &_milestone.LanesToDisplay)
					}
				} else {
					// ensure lane_ is NOT in _milestone.LanesToDisplay
					idx := slices.Index(_milestone.LanesToDisplay, lane_)
					if idx != -1 {
						_milestone.LanesToDisplay = slices.Delete(_milestone.LanesToDisplay, idx, idx+1)
						laneFormCallback.probe.UpdateSliceOfPointersCallback(_milestone, "LanesToDisplay", &_milestone.LanesToDisplay)
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
	updateProbeTable[*models.Lane](
		laneFormCallback.probe,
	)

	// display a new form by reset the form stage
	if laneFormCallback.CreationMode || laneFormCallback.formGroup.HasSuppressButtonBeenPressed {
		laneFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
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

	laneFormCallback.probe.ux_tree()
}
func __gong__New__LaneUseFormCallback(
	laneuse *models.LaneUse,
	probe *Probe,
	formGroup *form.FormGroup,
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

	formGroup *form.FormGroup
}

func (laneuseFormCallback *LaneUseFormCallback) OnSave() {
	laneuseFormCallback.probe.stageOfInterest.Lock()
	defer laneuseFormCallback.probe.stageOfInterest.Unlock()

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
		}
	}

	// manage the suppress operation
	if laneuseFormCallback.formGroup.HasSuppressButtonBeenPressed {
		laneuse_.Unstage(laneuseFormCallback.probe.stageOfInterest)
	}

	laneuseFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.LaneUse](
		laneuseFormCallback.probe,
	)

	// display a new form by reset the form stage
	if laneuseFormCallback.CreationMode || laneuseFormCallback.formGroup.HasSuppressButtonBeenPressed {
		laneuseFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
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

	laneuseFormCallback.probe.ux_tree()
}
func __gong__New__MilestoneFormCallback(
	milestone *models.Milestone,
	probe *Probe,
	formGroup *form.FormGroup,
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

	formGroup *form.FormGroup
}

func (milestoneFormCallback *MilestoneFormCallback) OnSave() {
	milestoneFormCallback.probe.stageOfInterest.Lock()
	defer milestoneFormCallback.probe.stageOfInterest.Unlock()

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
		case "LanesToDisplay":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Lane](milestoneFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Lane, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Lane)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					milestoneFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Lane](milestoneFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			milestone_.LanesToDisplay = instanceSlice
			milestoneFormCallback.probe.UpdateSliceOfPointersCallback(milestone_, "LanesToDisplay", &milestone_.LanesToDisplay)

		case "Gantt:Milestones":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Gantt instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Gantt instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Gantt](milestoneFormCallback.probe.stageOfInterest)
			targetGanttIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetGanttIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Gantt instances and update their Milestones slice
			for _gantt := range *models.GetGongstructInstancesSetFromPointerType[*models.Gantt](milestoneFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(milestoneFormCallback.probe.stageOfInterest, _gantt)
				
				// if Gantt is selected
				if targetGanttIDs[id] {
					// ensure milestone_ is in _gantt.Milestones
					found := false
					for _, _b := range _gantt.Milestones {
						if _b == milestone_ {
							found = true
							break
						}
					}
					if !found {
						_gantt.Milestones = append(_gantt.Milestones, milestone_)
						milestoneFormCallback.probe.UpdateSliceOfPointersCallback(_gantt, "Milestones", &_gantt.Milestones)
					}
				} else {
					// ensure milestone_ is NOT in _gantt.Milestones
					idx := slices.Index(_gantt.Milestones, milestone_)
					if idx != -1 {
						_gantt.Milestones = slices.Delete(_gantt.Milestones, idx, idx+1)
						milestoneFormCallback.probe.UpdateSliceOfPointersCallback(_gantt, "Milestones", &_gantt.Milestones)
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
	updateProbeTable[*models.Milestone](
		milestoneFormCallback.probe,
	)

	// display a new form by reset the form stage
	if milestoneFormCallback.CreationMode || milestoneFormCallback.formGroup.HasSuppressButtonBeenPressed {
		milestoneFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
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

	milestoneFormCallback.probe.ux_tree()
}
