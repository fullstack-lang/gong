// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	table "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gong/lib/button/go/models"
	"github.com/fullstack-lang/gong/lib/button/go/orm"
)

const __dummmy__time = time.Nanosecond

var __dummmy__letters = slices.Delete([]string{"a"}, 0, 1)
var __dummy_orm = orm.BackRepoStruct{}

// insertion point
func __gong__New__ButtonFormCallback(
	button *models.Button,
	probe *Probe,
	formGroup *table.FormGroup,
) (buttonFormCallback *ButtonFormCallback) {
	buttonFormCallback = new(ButtonFormCallback)
	buttonFormCallback.probe = probe
	buttonFormCallback.button = button
	buttonFormCallback.formGroup = formGroup

	buttonFormCallback.CreationMode = (button == nil)

	return
}

type ButtonFormCallback struct {
	button *models.Button

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (buttonFormCallback *ButtonFormCallback) OnSave() {

	log.Println("ButtonFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	buttonFormCallback.probe.formStage.Checkout()

	if buttonFormCallback.button == nil {
		buttonFormCallback.button = new(models.Button).Stage(buttonFormCallback.probe.stageOfInterest)
	}
	button_ := buttonFormCallback.button
	_ = button_

	for _, formDiv := range buttonFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(button_.Name), formDiv)
		case "Label":
			FormDivBasicFieldToField(&(button_.Label), formDiv)
		case "Icon":
			FormDivBasicFieldToField(&(button_.Icon), formDiv)
		case "Group:Buttons":
			// we need to retrieve the field owner before the change
			var pastGroupOwner *models.Group
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Group"
			rf.Fieldname = "Buttons"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				buttonFormCallback.probe.stageOfInterest,
				buttonFormCallback.probe.backRepoOfInterest,
				button_,
				&rf)

			if reverseFieldOwner != nil {
				pastGroupOwner = reverseFieldOwner.(*models.Group)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastGroupOwner != nil {
					idx := slices.Index(pastGroupOwner.Buttons, button_)
					pastGroupOwner.Buttons = slices.Delete(pastGroupOwner.Buttons, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _group := range *models.GetGongstructInstancesSet[models.Group](buttonFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _group.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newGroupOwner := _group // we have a match
						if pastGroupOwner != nil {
							if newGroupOwner != pastGroupOwner {
								idx := slices.Index(pastGroupOwner.Buttons, button_)
								pastGroupOwner.Buttons = slices.Delete(pastGroupOwner.Buttons, idx, idx+1)
								newGroupOwner.Buttons = append(newGroupOwner.Buttons, button_)
							}
						} else {
							newGroupOwner.Buttons = append(newGroupOwner.Buttons, button_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if buttonFormCallback.formGroup.HasSuppressButtonBeenPressed {
		button_.Unstage(buttonFormCallback.probe.stageOfInterest)
	}

	buttonFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Button](
		buttonFormCallback.probe,
	)
	buttonFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if buttonFormCallback.CreationMode || buttonFormCallback.formGroup.HasSuppressButtonBeenPressed {
		buttonFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(buttonFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ButtonFormCallback(
			nil,
			buttonFormCallback.probe,
			newFormGroup,
		)
		button := new(models.Button)
		FillUpForm(button, newFormGroup, buttonFormCallback.probe)
		buttonFormCallback.probe.formStage.Commit()
	}

	fillUpTree(buttonFormCallback.probe)
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
