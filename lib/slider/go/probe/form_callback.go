// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	table "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gong/lib/slider/go/models"
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

	// log.Println("CheckboxFormCallback, OnSave")

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
			// WARNING : this form deals with the N-N association "Group.Checkboxes []*Checkbox" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Checkbox). Setting up a value
			// will discard the former value is there is one.
			//
			// the algorithm is
			// 1/ get the former source of the association
			var formerSource *models.Group
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Group"
				rf.Fieldname = "Checkboxes"
				formerAssociationSource := models.GetReverseFieldOwner(
					checkboxFormCallback.probe.stageOfInterest,
					checkbox_,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Group)
					if !ok {
						log.Fatalln("Source of Group.Checkboxes []*Checkbox, is not an Group instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				if formerSource != nil {
					idx := slices.Index(formerSource.Checkboxes, checkbox_)
					formerSource.Checkboxes = slices.Delete(formerSource.Checkboxes, idx, idx+1)
				}
				break // nothing else to do for this field
			}

			// we need to deal with the 2 cases:
			// 1 the field source is unchanged
			// 2 the field source is changed

			// 1 field source is unchanged
			if formerSource != nil && formerSource.GetName() == newSourceName.GetName() {
				break // nothing else to do for this field
			}

			// 2 field source is changed -->
			// (1) clear the source slice field if it exist
			// (2) find the new source
			// (3) append the new value to the new source field

			// (1) clear the source slice field if it exist
			if formerSource != nil {
				idx := slices.Index(formerSource.Checkboxes, checkbox_)
				formerSource.Checkboxes = slices.Delete(formerSource.Checkboxes, idx, idx+1)
			}

			// (2) find the source
			var newSource *models.Group
			for _group := range *models.GetGongstructInstancesSet[models.Group](checkboxFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _group.GetName() == newSourceName.GetName() {
					newSource = _group // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Group.Checkboxes []*Checkbox, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.Checkboxes = append(newSource.Checkboxes, checkbox_)
		}
	}

	// manage the suppress operation
	if checkboxFormCallback.formGroup.HasSuppressButtonBeenPressed {
		checkbox_.Unstage(checkboxFormCallback.probe.stageOfInterest)
	}

	checkboxFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.Checkbox](
		checkboxFormCallback.probe,
	)
	checkboxFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if checkboxFormCallback.CreationMode || checkboxFormCallback.formGroup.HasSuppressButtonBeenPressed {
		checkboxFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
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

	updateAndCommitTree(checkboxFormCallback.probe)
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
		case "Percentage":
			FormDivBasicFieldToField(&(group_.Percentage), formDiv)
		case "Layout:Groups":
			// WARNING : this form deals with the N-N association "Layout.Groups []*Group" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Group). Setting up a value
			// will discard the former value is there is one.
			//
			// the algorithm is
			// 1/ get the former source of the association
			var formerSource *models.Layout
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Layout"
				rf.Fieldname = "Groups"
				formerAssociationSource := models.GetReverseFieldOwner(
					groupFormCallback.probe.stageOfInterest,
					group_,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Layout)
					if !ok {
						log.Fatalln("Source of Layout.Groups []*Group, is not an Layout instance")
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
				break // nothing else to do for this field
			}

			// we need to deal with the 2 cases:
			// 1 the field source is unchanged
			// 2 the field source is changed

			// 1 field source is unchanged
			if formerSource != nil && formerSource.GetName() == newSourceName.GetName() {
				break // nothing else to do for this field
			}

			// 2 field source is changed -->
			// (1) clear the source slice field if it exist
			// (2) find the new source
			// (3) append the new value to the new source field

			// (1) clear the source slice field if it exist
			if formerSource != nil {
				idx := slices.Index(formerSource.Groups, group_)
				formerSource.Groups = slices.Delete(formerSource.Groups, idx, idx+1)
			}

			// (2) find the source
			var newSource *models.Layout
			for _layout := range *models.GetGongstructInstancesSet[models.Layout](groupFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _layout.GetName() == newSourceName.GetName() {
					newSource = _layout // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Layout.Groups []*Group, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
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

	// log.Println("LayoutFormCallback, OnSave")

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
	updateAndCommitTable[models.Layout](
		layoutFormCallback.probe,
	)
	layoutFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if layoutFormCallback.CreationMode || layoutFormCallback.formGroup.HasSuppressButtonBeenPressed {
		layoutFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
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

	updateAndCommitTree(layoutFormCallback.probe)
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

	// log.Println("SliderFormCallback, OnSave")

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
			// WARNING : this form deals with the N-N association "Group.Sliders []*Slider" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Slider). Setting up a value
			// will discard the former value is there is one.
			//
			// the algorithm is
			// 1/ get the former source of the association
			var formerSource *models.Group
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Group"
				rf.Fieldname = "Sliders"
				formerAssociationSource := models.GetReverseFieldOwner(
					sliderFormCallback.probe.stageOfInterest,
					slider_,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Group)
					if !ok {
						log.Fatalln("Source of Group.Sliders []*Slider, is not an Group instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				if formerSource != nil {
					idx := slices.Index(formerSource.Sliders, slider_)
					formerSource.Sliders = slices.Delete(formerSource.Sliders, idx, idx+1)
				}
				break // nothing else to do for this field
			}

			// we need to deal with the 2 cases:
			// 1 the field source is unchanged
			// 2 the field source is changed

			// 1 field source is unchanged
			if formerSource != nil && formerSource.GetName() == newSourceName.GetName() {
				break // nothing else to do for this field
			}

			// 2 field source is changed -->
			// (1) clear the source slice field if it exist
			// (2) find the new source
			// (3) append the new value to the new source field

			// (1) clear the source slice field if it exist
			if formerSource != nil {
				idx := slices.Index(formerSource.Sliders, slider_)
				formerSource.Sliders = slices.Delete(formerSource.Sliders, idx, idx+1)
			}

			// (2) find the source
			var newSource *models.Group
			for _group := range *models.GetGongstructInstancesSet[models.Group](sliderFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _group.GetName() == newSourceName.GetName() {
					newSource = _group // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Group.Sliders []*Slider, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.Sliders = append(newSource.Sliders, slider_)
		}
	}

	// manage the suppress operation
	if sliderFormCallback.formGroup.HasSuppressButtonBeenPressed {
		slider_.Unstage(sliderFormCallback.probe.stageOfInterest)
	}

	sliderFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.Slider](
		sliderFormCallback.probe,
	)
	sliderFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if sliderFormCallback.CreationMode || sliderFormCallback.formGroup.HasSuppressButtonBeenPressed {
		sliderFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
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

	updateAndCommitTree(sliderFormCallback.probe)
}
