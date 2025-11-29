// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	table "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gong/lib/button/go/models"
)

// to avoid errors when time and slices packages are not used in the generated code
const _ = time.Nanosecond

var _ = slices.Delete([]string{"a"}, 0, 1)

var _ = log.Panicf

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

	// log.Println("ButtonFormCallback, OnSave")

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
		case "IsDisabled":
			FormDivBasicFieldToField(&(button_.IsDisabled), formDiv)
		case "Color":
			FormDivEnumStringFieldToField(&(button_.Color), formDiv)
		case "MatButtonType":
			FormDivEnumStringFieldToField(&(button_.MatButtonType), formDiv)
		case "MatButtonAppearance":
			FormDivEnumStringFieldToField(&(button_.MatButtonAppearance), formDiv)
		case "Group:Buttons":
			// WARNING : this form deals with the N-N association "Group.Buttons []*Button" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Button). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Group
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Group"
				rf.Fieldname = "Buttons"
				formerAssociationSource := button_.GongGetReverseFieldOwner(
					buttonFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Group)
					if !ok {
						log.Fatalln("Source of Group.Buttons []*Button, is not an Group instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.Buttons, button_)
					formerSource.Buttons = slices.Delete(formerSource.Buttons, idx, idx+1)
				}
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
			for _group := range *models.GetGongstructInstancesSet[models.Group](buttonFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _group.GetName() == newSourceName.GetName() {
					newSource = _group // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Group.Buttons []*Button, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.Buttons = append(newSource.Buttons, button_)
		}
	}

	// manage the suppress operation
	if buttonFormCallback.formGroup.HasSuppressButtonBeenPressed {
		button_.Unstage(buttonFormCallback.probe.stageOfInterest)
	}

	buttonFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[*models.Button](
		buttonFormCallback.probe,
	)
	buttonFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if buttonFormCallback.CreationMode || buttonFormCallback.formGroup.HasSuppressButtonBeenPressed {
		buttonFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
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

	updateAndCommitTree(buttonFormCallback.probe)
}
func __gong__New__ButtonToggleFormCallback(
	buttontoggle *models.ButtonToggle,
	probe *Probe,
	formGroup *table.FormGroup,
) (buttontoggleFormCallback *ButtonToggleFormCallback) {
	buttontoggleFormCallback = new(ButtonToggleFormCallback)
	buttontoggleFormCallback.probe = probe
	buttontoggleFormCallback.buttontoggle = buttontoggle
	buttontoggleFormCallback.formGroup = formGroup

	buttontoggleFormCallback.CreationMode = (buttontoggle == nil)

	return
}

type ButtonToggleFormCallback struct {
	buttontoggle *models.ButtonToggle

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (buttontoggleFormCallback *ButtonToggleFormCallback) OnSave() {

	// log.Println("ButtonToggleFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	buttontoggleFormCallback.probe.formStage.Checkout()

	if buttontoggleFormCallback.buttontoggle == nil {
		buttontoggleFormCallback.buttontoggle = new(models.ButtonToggle).Stage(buttontoggleFormCallback.probe.stageOfInterest)
	}
	buttontoggle_ := buttontoggleFormCallback.buttontoggle
	_ = buttontoggle_

	for _, formDiv := range buttontoggleFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(buttontoggle_.Name), formDiv)
		case "Label":
			FormDivBasicFieldToField(&(buttontoggle_.Label), formDiv)
		case "Icon":
			FormDivBasicFieldToField(&(buttontoggle_.Icon), formDiv)
		case "IsDisabled":
			FormDivBasicFieldToField(&(buttontoggle_.IsDisabled), formDiv)
		case "IsChecked":
			FormDivBasicFieldToField(&(buttontoggle_.IsChecked), formDiv)
		case "GroupToogle:ButtonToggles":
			// WARNING : this form deals with the N-N association "GroupToogle.ButtonToggles []*ButtonToggle" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of ButtonToggle). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.GroupToogle
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "GroupToogle"
				rf.Fieldname = "ButtonToggles"
				formerAssociationSource := buttontoggle_.GongGetReverseFieldOwner(
					buttontoggleFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.GroupToogle)
					if !ok {
						log.Fatalln("Source of GroupToogle.ButtonToggles []*ButtonToggle, is not an GroupToogle instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.ButtonToggles, buttontoggle_)
					formerSource.ButtonToggles = slices.Delete(formerSource.ButtonToggles, idx, idx+1)
				}
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.GroupToogle
			for _grouptoogle := range *models.GetGongstructInstancesSet[models.GroupToogle](buttontoggleFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _grouptoogle.GetName() == newSourceName.GetName() {
					newSource = _grouptoogle // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of GroupToogle.ButtonToggles []*ButtonToggle, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.ButtonToggles = append(newSource.ButtonToggles, buttontoggle_)
		}
	}

	// manage the suppress operation
	if buttontoggleFormCallback.formGroup.HasSuppressButtonBeenPressed {
		buttontoggle_.Unstage(buttontoggleFormCallback.probe.stageOfInterest)
	}

	buttontoggleFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[*models.ButtonToggle](
		buttontoggleFormCallback.probe,
	)
	buttontoggleFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if buttontoggleFormCallback.CreationMode || buttontoggleFormCallback.formGroup.HasSuppressButtonBeenPressed {
		buttontoggleFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(buttontoggleFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ButtonToggleFormCallback(
			nil,
			buttontoggleFormCallback.probe,
			newFormGroup,
		)
		buttontoggle := new(models.ButtonToggle)
		FillUpForm(buttontoggle, newFormGroup, buttontoggleFormCallback.probe)
		buttontoggleFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(buttontoggleFormCallback.probe)
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
		case "Buttons":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Button](groupFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Button, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Button)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					groupFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			ids, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			for _, id := range ids {
				instanceSlice = append(instanceSlice, map_id_instances[id])
			}
			group_.Buttons = instanceSlice

		case "NbColumns":
			FormDivBasicFieldToField(&(group_.NbColumns), formDiv)
		case "Layout:Groups":
			// WARNING : this form deals with the N-N association "Layout.Groups []*Group" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Group). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Layout
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Layout"
				rf.Fieldname = "Groups"
				formerAssociationSource := group_.GongGetReverseFieldOwner(
					groupFormCallback.probe.stageOfInterest,
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
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.Groups, group_)
					formerSource.Groups = slices.Delete(formerSource.Groups, idx, idx+1)
				}
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
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
	updateAndCommitTable[*models.Group](
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
func __gong__New__GroupToogleFormCallback(
	grouptoogle *models.GroupToogle,
	probe *Probe,
	formGroup *table.FormGroup,
) (grouptoogleFormCallback *GroupToogleFormCallback) {
	grouptoogleFormCallback = new(GroupToogleFormCallback)
	grouptoogleFormCallback.probe = probe
	grouptoogleFormCallback.grouptoogle = grouptoogle
	grouptoogleFormCallback.formGroup = formGroup

	grouptoogleFormCallback.CreationMode = (grouptoogle == nil)

	return
}

type GroupToogleFormCallback struct {
	grouptoogle *models.GroupToogle

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (grouptoogleFormCallback *GroupToogleFormCallback) OnSave() {

	// log.Println("GroupToogleFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	grouptoogleFormCallback.probe.formStage.Checkout()

	if grouptoogleFormCallback.grouptoogle == nil {
		grouptoogleFormCallback.grouptoogle = new(models.GroupToogle).Stage(grouptoogleFormCallback.probe.stageOfInterest)
	}
	grouptoogle_ := grouptoogleFormCallback.grouptoogle
	_ = grouptoogle_

	for _, formDiv := range grouptoogleFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(grouptoogle_.Name), formDiv)
		case "Percentage":
			FormDivBasicFieldToField(&(grouptoogle_.Percentage), formDiv)
		case "ButtonToggles":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.ButtonToggle](grouptoogleFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.ButtonToggle, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.ButtonToggle)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					grouptoogleFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			ids, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			for _, id := range ids {
				instanceSlice = append(instanceSlice, map_id_instances[id])
			}
			grouptoogle_.ButtonToggles = instanceSlice

		case "IsSingleSelector":
			FormDivBasicFieldToField(&(grouptoogle_.IsSingleSelector), formDiv)
		case "Layout:GroupToogles":
			// WARNING : this form deals with the N-N association "Layout.GroupToogles []*GroupToogle" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of GroupToogle). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Layout
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Layout"
				rf.Fieldname = "GroupToogles"
				formerAssociationSource := grouptoogle_.GongGetReverseFieldOwner(
					grouptoogleFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Layout)
					if !ok {
						log.Fatalln("Source of Layout.GroupToogles []*GroupToogle, is not an Layout instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.GroupToogles, grouptoogle_)
					formerSource.GroupToogles = slices.Delete(formerSource.GroupToogles, idx, idx+1)
				}
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.Layout
			for _layout := range *models.GetGongstructInstancesSet[models.Layout](grouptoogleFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _layout.GetName() == newSourceName.GetName() {
					newSource = _layout // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Layout.GroupToogles []*GroupToogle, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.GroupToogles = append(newSource.GroupToogles, grouptoogle_)
		}
	}

	// manage the suppress operation
	if grouptoogleFormCallback.formGroup.HasSuppressButtonBeenPressed {
		grouptoogle_.Unstage(grouptoogleFormCallback.probe.stageOfInterest)
	}

	grouptoogleFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[*models.GroupToogle](
		grouptoogleFormCallback.probe,
	)
	grouptoogleFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if grouptoogleFormCallback.CreationMode || grouptoogleFormCallback.formGroup.HasSuppressButtonBeenPressed {
		grouptoogleFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(grouptoogleFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__GroupToogleFormCallback(
			nil,
			grouptoogleFormCallback.probe,
			newFormGroup,
		)
		grouptoogle := new(models.GroupToogle)
		FillUpForm(grouptoogle, newFormGroup, grouptoogleFormCallback.probe)
		grouptoogleFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(grouptoogleFormCallback.probe)
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
		case "Groups":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Group](layoutFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Group, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Group)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					layoutFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			ids, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			for _, id := range ids {
				instanceSlice = append(instanceSlice, map_id_instances[id])
			}
			layout_.Groups = instanceSlice

		case "GroupToogles":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.GroupToogle](layoutFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.GroupToogle, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.GroupToogle)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					layoutFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			ids, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			for _, id := range ids {
				instanceSlice = append(instanceSlice, map_id_instances[id])
			}
			layout_.GroupToogles = instanceSlice

		}
	}

	// manage the suppress operation
	if layoutFormCallback.formGroup.HasSuppressButtonBeenPressed {
		layout_.Unstage(layoutFormCallback.probe.stageOfInterest)
	}

	layoutFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[*models.Layout](
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
