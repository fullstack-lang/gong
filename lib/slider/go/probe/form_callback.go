// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	table "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gong/lib/slider/go/models"
	"github.com/fullstack-lang/gong/lib/slider/go/orm"
)

const __dummmy__time = time.Nanosecond

var __dummmy__letters = slices.Delete([]string{"a"}, 0, 1)
var __dummy_orm = orm.BackRepoStruct{}

// insertion point
func __gong__New__CheckboxFormCallback(
	checkbox *models.Checkbox,
	probe *Probe,
	formGroup *table.FormGroup,
) (checkboxFormCallback *CheckboxFormCallback) {
	checkboxFormCallback = new(CheckboxFormCallback)
	checkboxFormCallback.probe = probe
	checkboxFormCallback.checkbox = checkbox
	checkboxFormCallback.formGroup = formGroup

	checkboxFormCallback.CreationMode = (checkbox == nil)

	return
}

type CheckboxFormCallback struct {
	checkbox *models.Checkbox

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (checkboxFormCallback *CheckboxFormCallback) OnSave() {

	log.Println("CheckboxFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	checkboxFormCallback.probe.formStage.Checkout()

	if checkboxFormCallback.checkbox == nil {
		checkboxFormCallback.checkbox = new(models.Checkbox).Stage(checkboxFormCallback.probe.stageOfInterest)
	}
	checkbox_ := checkboxFormCallback.checkbox
	_ = checkbox_

	for _, formDiv := range checkboxFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(checkbox_.Name), formDiv)
		case "ValueBool":
			FormDivBasicFieldToField(&(checkbox_.ValueBool), formDiv)
		case "LabelForTrue":
			FormDivBasicFieldToField(&(checkbox_.LabelForTrue), formDiv)
		case "LabelForFalse":
			FormDivBasicFieldToField(&(checkbox_.LabelForFalse), formDiv)
		case "Group:Checkboxes":
			// we need to retrieve the field owner before the change
			var pastGroupOwner *models.Group
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Group"
			rf.Fieldname = "Checkboxes"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				checkboxFormCallback.probe.stageOfInterest,
				checkboxFormCallback.probe.backRepoOfInterest,
				checkbox_,
				&rf)

			if reverseFieldOwner != nil {
				pastGroupOwner = reverseFieldOwner.(*models.Group)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastGroupOwner != nil {
					idx := slices.Index(pastGroupOwner.Checkboxes, checkbox_)
					pastGroupOwner.Checkboxes = slices.Delete(pastGroupOwner.Checkboxes, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _group := range *models.GetGongstructInstancesSet[models.Group](checkboxFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _group.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newGroupOwner := _group // we have a match
						if pastGroupOwner != nil {
							if newGroupOwner != pastGroupOwner {
								idx := slices.Index(pastGroupOwner.Checkboxes, checkbox_)
								pastGroupOwner.Checkboxes = slices.Delete(pastGroupOwner.Checkboxes, idx, idx+1)
								newGroupOwner.Checkboxes = append(newGroupOwner.Checkboxes, checkbox_)
							}
						} else {
							newGroupOwner.Checkboxes = append(newGroupOwner.Checkboxes, checkbox_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if checkboxFormCallback.formGroup.HasSuppressButtonBeenPressed {
		checkbox_.Unstage(checkboxFormCallback.probe.stageOfInterest)
	}

	checkboxFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Checkbox](
		checkboxFormCallback.probe,
	)
	checkboxFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if checkboxFormCallback.CreationMode || checkboxFormCallback.formGroup.HasSuppressButtonBeenPressed {
		checkboxFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(checkboxFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__CheckboxFormCallback(
			nil,
			checkboxFormCallback.probe,
			newFormGroup,
		)
		checkbox := new(models.Checkbox)
		FillUpForm(checkbox, newFormGroup, checkboxFormCallback.probe)
		checkboxFormCallback.probe.formStage.Commit()
	}

	fillUpTree(checkboxFormCallback.probe)
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
		case "Percentage":
			FormDivBasicFieldToField(&(group_.Percentage), formDiv)
		case "Layout:Groups":
			// we need to retrieve the field owner before the change
			var pastLayoutOwner *models.Layout
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Layout"
			rf.Fieldname = "Groups"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				groupFormCallback.probe.stageOfInterest,
				groupFormCallback.probe.backRepoOfInterest,
				group_,
				&rf)

			if reverseFieldOwner != nil {
				pastLayoutOwner = reverseFieldOwner.(*models.Layout)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastLayoutOwner != nil {
					idx := slices.Index(pastLayoutOwner.Groups, group_)
					pastLayoutOwner.Groups = slices.Delete(pastLayoutOwner.Groups, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _layout := range *models.GetGongstructInstancesSet[models.Layout](groupFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _layout.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newLayoutOwner := _layout // we have a match
						if pastLayoutOwner != nil {
							if newLayoutOwner != pastLayoutOwner {
								idx := slices.Index(pastLayoutOwner.Groups, group_)
								pastLayoutOwner.Groups = slices.Delete(pastLayoutOwner.Groups, idx, idx+1)
								newLayoutOwner.Groups = append(newLayoutOwner.Groups, group_)
							}
						} else {
							newLayoutOwner.Groups = append(newLayoutOwner.Groups, group_)
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
			Name: table.FormGroupDefaultName.ToString(),
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
func __gong__New__LayoutFormCallback(
	layout *models.Layout,
	probe *Probe,
	formGroup *table.FormGroup,
) (layoutFormCallback *LayoutFormCallback) {
	layoutFormCallback = new(LayoutFormCallback)
	layoutFormCallback.probe = probe
	layoutFormCallback.layout = layout
	layoutFormCallback.formGroup = formGroup

	layoutFormCallback.CreationMode = (layout == nil)

	return
}

type LayoutFormCallback struct {
	layout *models.Layout

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (layoutFormCallback *LayoutFormCallback) OnSave() {

	log.Println("LayoutFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	layoutFormCallback.probe.formStage.Checkout()

	if layoutFormCallback.layout == nil {
		layoutFormCallback.layout = new(models.Layout).Stage(layoutFormCallback.probe.stageOfInterest)
	}
	layout_ := layoutFormCallback.layout
	_ = layout_

	for _, formDiv := range layoutFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(layout_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if layoutFormCallback.formGroup.HasSuppressButtonBeenPressed {
		layout_.Unstage(layoutFormCallback.probe.stageOfInterest)
	}

	layoutFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Layout](
		layoutFormCallback.probe,
	)
	layoutFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if layoutFormCallback.CreationMode || layoutFormCallback.formGroup.HasSuppressButtonBeenPressed {
		layoutFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(layoutFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__LayoutFormCallback(
			nil,
			layoutFormCallback.probe,
			newFormGroup,
		)
		layout := new(models.Layout)
		FillUpForm(layout, newFormGroup, layoutFormCallback.probe)
		layoutFormCallback.probe.formStage.Commit()
	}

	fillUpTree(layoutFormCallback.probe)
}
func __gong__New__SliderFormCallback(
	slider *models.Slider,
	probe *Probe,
	formGroup *table.FormGroup,
) (sliderFormCallback *SliderFormCallback) {
	sliderFormCallback = new(SliderFormCallback)
	sliderFormCallback.probe = probe
	sliderFormCallback.slider = slider
	sliderFormCallback.formGroup = formGroup

	sliderFormCallback.CreationMode = (slider == nil)

	return
}

type SliderFormCallback struct {
	slider *models.Slider

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (sliderFormCallback *SliderFormCallback) OnSave() {

	log.Println("SliderFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	sliderFormCallback.probe.formStage.Checkout()

	if sliderFormCallback.slider == nil {
		sliderFormCallback.slider = new(models.Slider).Stage(sliderFormCallback.probe.stageOfInterest)
	}
	slider_ := sliderFormCallback.slider
	_ = slider_

	for _, formDiv := range sliderFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(slider_.Name), formDiv)
		case "IsFloat64":
			FormDivBasicFieldToField(&(slider_.IsFloat64), formDiv)
		case "IsInt":
			FormDivBasicFieldToField(&(slider_.IsInt), formDiv)
		case "MinInt":
			FormDivBasicFieldToField(&(slider_.MinInt), formDiv)
		case "MaxInt":
			FormDivBasicFieldToField(&(slider_.MaxInt), formDiv)
		case "StepInt":
			FormDivBasicFieldToField(&(slider_.StepInt), formDiv)
		case "ValueInt":
			FormDivBasicFieldToField(&(slider_.ValueInt), formDiv)
		case "MinFloat64":
			FormDivBasicFieldToField(&(slider_.MinFloat64), formDiv)
		case "MaxFloat64":
			FormDivBasicFieldToField(&(slider_.MaxFloat64), formDiv)
		case "StepFloat64":
			FormDivBasicFieldToField(&(slider_.StepFloat64), formDiv)
		case "ValueFloat64":
			FormDivBasicFieldToField(&(slider_.ValueFloat64), formDiv)
		case "Group:Sliders":
			// we need to retrieve the field owner before the change
			var pastGroupOwner *models.Group
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Group"
			rf.Fieldname = "Sliders"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				sliderFormCallback.probe.stageOfInterest,
				sliderFormCallback.probe.backRepoOfInterest,
				slider_,
				&rf)

			if reverseFieldOwner != nil {
				pastGroupOwner = reverseFieldOwner.(*models.Group)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastGroupOwner != nil {
					idx := slices.Index(pastGroupOwner.Sliders, slider_)
					pastGroupOwner.Sliders = slices.Delete(pastGroupOwner.Sliders, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _group := range *models.GetGongstructInstancesSet[models.Group](sliderFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _group.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newGroupOwner := _group // we have a match
						if pastGroupOwner != nil {
							if newGroupOwner != pastGroupOwner {
								idx := slices.Index(pastGroupOwner.Sliders, slider_)
								pastGroupOwner.Sliders = slices.Delete(pastGroupOwner.Sliders, idx, idx+1)
								newGroupOwner.Sliders = append(newGroupOwner.Sliders, slider_)
							}
						} else {
							newGroupOwner.Sliders = append(newGroupOwner.Sliders, slider_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if sliderFormCallback.formGroup.HasSuppressButtonBeenPressed {
		slider_.Unstage(sliderFormCallback.probe.stageOfInterest)
	}

	sliderFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Slider](
		sliderFormCallback.probe,
	)
	sliderFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if sliderFormCallback.CreationMode || sliderFormCallback.formGroup.HasSuppressButtonBeenPressed {
		sliderFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(sliderFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__SliderFormCallback(
			nil,
			sliderFormCallback.probe,
			newFormGroup,
		)
		slider := new(models.Slider)
		FillUpForm(slider, newFormGroup, sliderFormCallback.probe)
		sliderFormCallback.probe.formStage.Commit()
	}

	fillUpTree(sliderFormCallback.probe)
}
