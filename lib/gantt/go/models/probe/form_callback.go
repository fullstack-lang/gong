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

type FormCallbackIF interface {
	GetCreationMode() bool
	GetInstance() any
	GetGongstructName() string
	OnSave()
}

type FormCallback[T models.PointerToGongstruct] struct {
	Instance     T
	CreationMode bool
	probe        *Probe
	formGroup    *form.FormGroup
	saveFields   func(instance T, probe *Probe, formGroup *form.FormGroup)
}

func NewFormCallback[T models.PointerToGongstruct](
	instance T,
	probe *Probe,
	formGroup *form.FormGroup,
	saveFields func(instance T, probe *Probe, formGroup *form.FormGroup),
) *FormCallback[T] {
	return &FormCallback[T]{
		Instance:     instance,
		CreationMode: any(instance) == nil,
		probe:        probe,
		formGroup:    formGroup,
		saveFields:   saveFields,
	}
}

func (cb *FormCallback[T]) GetCreationMode() bool     { return cb.CreationMode }
func (cb *FormCallback[T]) GetInstance() any           { return cb.Instance }
func (cb *FormCallback[T]) GetGongstructName() string { return models.GetPointerToGongstructName[T]() }

func (cb *FormCallback[T]) OnSave() {
	cb.probe.stageOfInterest.Lock()
	defer cb.probe.stageOfInterest.Unlock()

	cb.probe.formStage.Checkout()

	if any(cb.Instance) == nil {
		cb.Instance = cb.probe.stageOfInterest.GongNewInstance[T]()
	}

	cb.saveFields(cb.Instance, cb.probe, cb.formGroup)

	if cb.formGroup.HasSuppressButtonBeenPressed {
		cb.Instance.UnstageVoid(cb.probe.stageOfInterest)
	}

	cb.probe.stageOfInterest.Commit()
	updateProbeTable[T](cb.probe)

	if cb.CreationMode || cb.formGroup.HasSuppressButtonBeenPressed {
		cb.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(cb.probe.formStage)
		newFormGroup.OnSave = NewFormCallback[T](
			*new(T),
			cb.probe,
			newFormGroup,
			cb.saveFields,
		)
		newInstance := models.GongNewInstance[T]()
		FillUpForm(newInstance, newFormGroup, cb.probe)
		cb.probe.formStage.Commit()
	}

	cb.probe.ux_tree()
}

// insertion point
func __gong__New__ArrowFormCallback(
	_instance *models.Arrow,
	probe *Probe,
	formGroup *form.FormGroup,
) (arrowFormCallback *FormCallback[*models.Arrow]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveArrowFields,
	)
}

type ArrowFormCallback = FormCallback[*models.Arrow]

func saveArrowFields(
	_instance *models.Arrow,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "From":
			FormDivSelectFieldToField(&(_instance.From), probe.stageOfInterest, formDiv)
		case "To":
			FormDivSelectFieldToField(&(_instance.To), probe.stageOfInterest, formDiv)
		case "OptionnalColor":
			FormDivBasicFieldToField(&(_instance.OptionnalColor), formDiv)
		case "OptionnalStroke":
			FormDivBasicFieldToField(&(_instance.OptionnalStroke), formDiv)
		case "Gantt:Arrows":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Arrows", func(owner *models.Gantt) *[]*models.Arrow { return &owner.Arrows })
		}
	}
}

func __gong__New__BarFormCallback(
	_instance *models.Bar,
	probe *Probe,
	formGroup *form.FormGroup,
) (barFormCallback *FormCallback[*models.Bar]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveBarFields,
	)
}

type BarFormCallback = FormCallback[*models.Bar]

func saveBarFields(
	_instance *models.Bar,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Start":
			FormDivTimeFieldToField(&(_instance.Start), formDiv, false)
		case "End":
			FormDivTimeFieldToField(&(_instance.End), formDiv, false)
		case "ComputedDuration":
			FormDivBasicFieldToField(&(_instance.ComputedDuration), formDiv)
		case "OptionnalColor":
			FormDivBasicFieldToField(&(_instance.OptionnalColor), formDiv)
		case "OptionnalStroke":
			FormDivBasicFieldToField(&(_instance.OptionnalStroke), formDiv)
		case "FillOpacity":
			FormDivBasicFieldToField(&(_instance.FillOpacity), formDiv)
		case "StrokeWidth":
			FormDivBasicFieldToField(&(_instance.StrokeWidth), formDiv)
		case "StrokeDashArray":
			FormDivBasicFieldToField(&(_instance.StrokeDashArray), formDiv)
		case "Lane:Bars":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Bars", func(owner *models.Lane) *[]*models.Bar { return &owner.Bars })
		}
	}
}

func __gong__New__GanttFormCallback(
	_instance *models.Gantt,
	probe *Probe,
	formGroup *form.FormGroup,
) (ganttFormCallback *FormCallback[*models.Gantt]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveGanttFields,
	)
}

type GanttFormCallback = FormCallback[*models.Gantt]

func saveGanttFields(
	_instance *models.Gantt,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "ComputedStart":
			FormDivTimeFieldToField(&(_instance.ComputedStart), formDiv, false)
		case "ComputedEnd":
			FormDivTimeFieldToField(&(_instance.ComputedEnd), formDiv, false)
		case "ComputedDuration":
			FormDivBasicFieldToField(&(_instance.ComputedDuration), formDiv)
		case "UseManualStartAndEndDates":
			FormDivBasicFieldToField(&(_instance.UseManualStartAndEndDates), formDiv)
		case "ManualStart":
			FormDivTimeFieldToField(&(_instance.ManualStart), formDiv, false)
		case "ManualEnd":
			FormDivTimeFieldToField(&(_instance.ManualEnd), formDiv, false)
		case "LaneHeight":
			FormDivBasicFieldToField(&(_instance.LaneHeight), formDiv)
		case "RatioBarToLaneHeight":
			FormDivBasicFieldToField(&(_instance.RatioBarToLaneHeight), formDiv)
		case "YTopMargin":
			FormDivBasicFieldToField(&(_instance.YTopMargin), formDiv)
		case "XLeftText":
			FormDivBasicFieldToField(&(_instance.XLeftText), formDiv)
		case "TextHeight":
			FormDivBasicFieldToField(&(_instance.TextHeight), formDiv)
		case "XLeftLanes":
			FormDivBasicFieldToField(&(_instance.XLeftLanes), formDiv)
		case "XRightMargin":
			FormDivBasicFieldToField(&(_instance.XRightMargin), formDiv)
		case "ArrowLengthToTheRightOfStartBar":
			FormDivBasicFieldToField(&(_instance.ArrowLengthToTheRightOfStartBar), formDiv)
		case "ArrowTipLenght":
			FormDivBasicFieldToField(&(_instance.ArrowTipLenght), formDiv)
		case "TimeLine_Color":
			FormDivBasicFieldToField(&(_instance.TimeLine_Color), formDiv)
		case "TimeLine_FillOpacity":
			FormDivBasicFieldToField(&(_instance.TimeLine_FillOpacity), formDiv)
		case "TimeLine_Stroke":
			FormDivBasicFieldToField(&(_instance.TimeLine_Stroke), formDiv)
		case "TimeLine_StrokeWidth":
			FormDivBasicFieldToField(&(_instance.TimeLine_StrokeWidth), formDiv)
		case "Group_Stroke":
			FormDivBasicFieldToField(&(_instance.Group_Stroke), formDiv)
		case "Group_StrokeWidth":
			FormDivBasicFieldToField(&(_instance.Group_StrokeWidth), formDiv)
		case "Group_StrokeDashArray":
			FormDivBasicFieldToField(&(_instance.Group_StrokeDashArray), formDiv)
		case "DateYOffset":
			FormDivBasicFieldToField(&(_instance.DateYOffset), formDiv)
		case "AlignOnStartEndOnYearStart":
			FormDivBasicFieldToField(&(_instance.AlignOnStartEndOnYearStart), formDiv)
		case "Lanes":
			FormDivSliceOfPointersToField(_instance, "Lanes", &(_instance.Lanes), formDiv, probe)
		case "Milestones":
			FormDivSliceOfPointersToField(_instance, "Milestones", &(_instance.Milestones), formDiv, probe)
		case "Groups":
			FormDivSliceOfPointersToField(_instance, "Groups", &(_instance.Groups), formDiv, probe)
		case "Arrows":
			FormDivSliceOfPointersToField(_instance, "Arrows", &(_instance.Arrows), formDiv, probe)
		}
	}
}

func __gong__New__GroupFormCallback(
	_instance *models.Group,
	probe *Probe,
	formGroup *form.FormGroup,
) (groupFormCallback *FormCallback[*models.Group]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveGroupFields,
	)
}

type GroupFormCallback = FormCallback[*models.Group]

func saveGroupFields(
	_instance *models.Group,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "GroupLanes":
			FormDivSliceOfPointersToField(_instance, "GroupLanes", &(_instance.GroupLanes), formDiv, probe)
		case "Gantt:Groups":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Groups", func(owner *models.Gantt) *[]*models.Group { return &owner.Groups })
		}
	}
}

func __gong__New__LaneFormCallback(
	_instance *models.Lane,
	probe *Probe,
	formGroup *form.FormGroup,
) (laneFormCallback *FormCallback[*models.Lane]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveLaneFields,
	)
}

type LaneFormCallback = FormCallback[*models.Lane]

func saveLaneFields(
	_instance *models.Lane,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Order":
			FormDivBasicFieldToField(&(_instance.Order), formDiv)
		case "Bars":
			FormDivSliceOfPointersToField(_instance, "Bars", &(_instance.Bars), formDiv, probe)
		case "Gantt:Lanes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Lanes", func(owner *models.Gantt) *[]*models.Lane { return &owner.Lanes })
		case "Group:GroupLanes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "GroupLanes", func(owner *models.Group) *[]*models.Lane { return &owner.GroupLanes })
		case "Milestone:LanesToDisplay":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "LanesToDisplay", func(owner *models.Milestone) *[]*models.Lane { return &owner.LanesToDisplay })
		}
	}
}

func __gong__New__LaneUseFormCallback(
	_instance *models.LaneUse,
	probe *Probe,
	formGroup *form.FormGroup,
) (laneuseFormCallback *FormCallback[*models.LaneUse]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveLaneUseFields,
	)
}

type LaneUseFormCallback = FormCallback[*models.LaneUse]

func saveLaneUseFields(
	_instance *models.LaneUse,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Lane":
			FormDivSelectFieldToField(&(_instance.Lane), probe.stageOfInterest, formDiv)
		}
	}
}

func __gong__New__MilestoneFormCallback(
	_instance *models.Milestone,
	probe *Probe,
	formGroup *form.FormGroup,
) (milestoneFormCallback *FormCallback[*models.Milestone]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveMilestoneFields,
	)
}

type MilestoneFormCallback = FormCallback[*models.Milestone]

func saveMilestoneFields(
	_instance *models.Milestone,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Date":
			FormDivTimeFieldToField(&(_instance.Date), formDiv, false)
		case "DisplayVerticalBar":
			FormDivBasicFieldToField(&(_instance.DisplayVerticalBar), formDiv)
		case "LanesToDisplay":
			FormDivSliceOfPointersToField(_instance, "LanesToDisplay", &(_instance.LanesToDisplay), formDiv, probe)
		case "Gantt:Milestones":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Milestones", func(owner *models.Gantt) *[]*models.Milestone { return &owner.Milestones })
		}
	}
}

