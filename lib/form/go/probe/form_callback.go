// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/lib/form/go/models"
)

// to avoid errors when time and slices packages are not used in the generated code
const _ = time.Nanosecond

var _ = slices.Delete([]string{"a"}, 0, 1)

var _ = log.Panicf

// insertion point
func __gong__New__CheckBoxFormCallback(
	checkbox *models.CheckBox,
	probe *Probe,
	formGroup *form.FormGroup,
) (checkboxFormCallback *CheckBoxFormCallback) {
	checkboxFormCallback = new(CheckBoxFormCallback)
	checkboxFormCallback.probe = probe
	checkboxFormCallback.checkbox = checkbox
	checkboxFormCallback.formGroup = formGroup

	checkboxFormCallback.CreationMode = (checkbox == nil)

	return
}

type CheckBoxFormCallback struct {
	checkbox *models.CheckBox

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (checkboxFormCallback *CheckBoxFormCallback) OnSave() {
	checkboxFormCallback.probe.stageOfInterest.Lock()
	defer checkboxFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("CheckBoxFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	checkboxFormCallback.probe.formStage.Checkout()

	if checkboxFormCallback.checkbox == nil {
		checkboxFormCallback.checkbox = new(models.CheckBox).Stage(checkboxFormCallback.probe.stageOfInterest)
	}
	checkbox_ := checkboxFormCallback.checkbox
	_ = checkbox_

	for _, formDiv := range checkboxFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(checkbox_.Name), formDiv)
		case "Value":
			FormDivBasicFieldToField(&(checkbox_.Value), formDiv)
		case "FormDiv:CheckBoxs":
			// 1. Decode the AssociationStorage which contains the rowIDs of the FormDiv instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target FormDiv instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.FormDiv](checkboxFormCallback.probe.stageOfInterest)
			targetFormDivIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetFormDivIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all FormDiv instances and update their CheckBoxs slice
			for _formdiv := range *models.GetGongstructInstancesSetFromPointerType[*models.FormDiv](checkboxFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(checkboxFormCallback.probe.stageOfInterest, _formdiv)
				
				// if FormDiv is selected
				if targetFormDivIDs[id] {
					// ensure checkbox_ is in _formdiv.CheckBoxs
					found := false
					for _, _b := range _formdiv.CheckBoxs {
						if _b == checkbox_ {
							found = true
							break
						}
					}
					if !found {
						_formdiv.CheckBoxs = append(_formdiv.CheckBoxs, checkbox_)
						checkboxFormCallback.probe.UpdateSliceOfPointersCallback(_formdiv, "CheckBoxs", &_formdiv.CheckBoxs)
					}
				} else {
					// ensure checkbox_ is NOT in _formdiv.CheckBoxs
					idx := slices.Index(_formdiv.CheckBoxs, checkbox_)
					if idx != -1 {
						_formdiv.CheckBoxs = slices.Delete(_formdiv.CheckBoxs, idx, idx+1)
						checkboxFormCallback.probe.UpdateSliceOfPointersCallback(_formdiv, "CheckBoxs", &_formdiv.CheckBoxs)
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
	updateProbeTable[*models.CheckBox](
		checkboxFormCallback.probe,
	)

	// display a new form by reset the form stage
	if checkboxFormCallback.CreationMode || checkboxFormCallback.formGroup.HasSuppressButtonBeenPressed {
		checkboxFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(checkboxFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__CheckBoxFormCallback(
			nil,
			checkboxFormCallback.probe,
			newFormGroup,
		)
		checkbox := new(models.CheckBox)
		FillUpForm(checkbox, newFormGroup, checkboxFormCallback.probe)
		checkboxFormCallback.probe.formStage.Commit()
	}

	checkboxFormCallback.probe.ux_tree()
}
func __gong__New__FormDivFormCallback(
	formdiv *models.FormDiv,
	probe *Probe,
	formGroup *form.FormGroup,
) (formdivFormCallback *FormDivFormCallback) {
	formdivFormCallback = new(FormDivFormCallback)
	formdivFormCallback.probe = probe
	formdivFormCallback.formdiv = formdiv
	formdivFormCallback.formGroup = formGroup

	formdivFormCallback.CreationMode = (formdiv == nil)

	return
}

type FormDivFormCallback struct {
	formdiv *models.FormDiv

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (formdivFormCallback *FormDivFormCallback) OnSave() {
	formdivFormCallback.probe.stageOfInterest.Lock()
	defer formdivFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("FormDivFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	formdivFormCallback.probe.formStage.Checkout()

	if formdivFormCallback.formdiv == nil {
		formdivFormCallback.formdiv = new(models.FormDiv).Stage(formdivFormCallback.probe.stageOfInterest)
	}
	formdiv_ := formdivFormCallback.formdiv
	_ = formdiv_

	for _, formDiv := range formdivFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(formdiv_.Name), formDiv)
		case "FormFields":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.FormField](formdivFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.FormField, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.FormField)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					formdivFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.FormField](formdivFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			formdiv_.FormFields = instanceSlice
			formdivFormCallback.probe.UpdateSliceOfPointersCallback(formdiv_, "FormFields", &formdiv_.FormFields)

		case "CheckBoxs":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.CheckBox](formdivFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.CheckBox, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.CheckBox)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					formdivFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.CheckBox](formdivFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			formdiv_.CheckBoxs = instanceSlice
			formdivFormCallback.probe.UpdateSliceOfPointersCallback(formdiv_, "CheckBoxs", &formdiv_.CheckBoxs)

		case "FormEditAssocButton":
			FormDivSelectFieldToField(&(formdiv_.FormEditAssocButton), formdivFormCallback.probe.stageOfInterest, formDiv)
		case "FormSortAssocButton":
			FormDivSelectFieldToField(&(formdiv_.FormSortAssocButton), formdivFormCallback.probe.stageOfInterest, formDiv)
		case "IsADivider":
			FormDivBasicFieldToField(&(formdiv_.IsADivider), formDiv)
		case "IsAStartAccordionGroup":
			FormDivBasicFieldToField(&(formdiv_.IsAStartAccordionGroup), formDiv)
		case "AccordionGroupName":
			FormDivBasicFieldToField(&(formdiv_.AccordionGroupName), formDiv)
		case "IsAEndAccordionGroup":
			FormDivBasicFieldToField(&(formdiv_.IsAEndAccordionGroup), formDiv)
		case "FormGroup:FormDivs":
			// 1. Decode the AssociationStorage which contains the rowIDs of the FormGroup instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target FormGroup instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.FormGroup](formdivFormCallback.probe.stageOfInterest)
			targetFormGroupIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetFormGroupIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all FormGroup instances and update their FormDivs slice
			for _formgroup := range *models.GetGongstructInstancesSetFromPointerType[*models.FormGroup](formdivFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(formdivFormCallback.probe.stageOfInterest, _formgroup)
				
				// if FormGroup is selected
				if targetFormGroupIDs[id] {
					// ensure formdiv_ is in _formgroup.FormDivs
					found := false
					for _, _b := range _formgroup.FormDivs {
						if _b == formdiv_ {
							found = true
							break
						}
					}
					if !found {
						_formgroup.FormDivs = append(_formgroup.FormDivs, formdiv_)
						formdivFormCallback.probe.UpdateSliceOfPointersCallback(_formgroup, "FormDivs", &_formgroup.FormDivs)
					}
				} else {
					// ensure formdiv_ is NOT in _formgroup.FormDivs
					idx := slices.Index(_formgroup.FormDivs, formdiv_)
					if idx != -1 {
						_formgroup.FormDivs = slices.Delete(_formgroup.FormDivs, idx, idx+1)
						formdivFormCallback.probe.UpdateSliceOfPointersCallback(_formgroup, "FormDivs", &_formgroup.FormDivs)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if formdivFormCallback.formGroup.HasSuppressButtonBeenPressed {
		formdiv_.Unstage(formdivFormCallback.probe.stageOfInterest)
	}

	formdivFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.FormDiv](
		formdivFormCallback.probe,
	)

	// display a new form by reset the form stage
	if formdivFormCallback.CreationMode || formdivFormCallback.formGroup.HasSuppressButtonBeenPressed {
		formdivFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(formdivFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__FormDivFormCallback(
			nil,
			formdivFormCallback.probe,
			newFormGroup,
		)
		formdiv := new(models.FormDiv)
		FillUpForm(formdiv, newFormGroup, formdivFormCallback.probe)
		formdivFormCallback.probe.formStage.Commit()
	}

	formdivFormCallback.probe.ux_tree()
}
func __gong__New__FormEditAssocButtonFormCallback(
	formeditassocbutton *models.FormEditAssocButton,
	probe *Probe,
	formGroup *form.FormGroup,
) (formeditassocbuttonFormCallback *FormEditAssocButtonFormCallback) {
	formeditassocbuttonFormCallback = new(FormEditAssocButtonFormCallback)
	formeditassocbuttonFormCallback.probe = probe
	formeditassocbuttonFormCallback.formeditassocbutton = formeditassocbutton
	formeditassocbuttonFormCallback.formGroup = formGroup

	formeditassocbuttonFormCallback.CreationMode = (formeditassocbutton == nil)

	return
}

type FormEditAssocButtonFormCallback struct {
	formeditassocbutton *models.FormEditAssocButton

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (formeditassocbuttonFormCallback *FormEditAssocButtonFormCallback) OnSave() {
	formeditassocbuttonFormCallback.probe.stageOfInterest.Lock()
	defer formeditassocbuttonFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("FormEditAssocButtonFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	formeditassocbuttonFormCallback.probe.formStage.Checkout()

	if formeditassocbuttonFormCallback.formeditassocbutton == nil {
		formeditassocbuttonFormCallback.formeditassocbutton = new(models.FormEditAssocButton).Stage(formeditassocbuttonFormCallback.probe.stageOfInterest)
	}
	formeditassocbutton_ := formeditassocbuttonFormCallback.formeditassocbutton
	_ = formeditassocbutton_

	for _, formDiv := range formeditassocbuttonFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(formeditassocbutton_.Name), formDiv)
		case "Label":
			FormDivBasicFieldToField(&(formeditassocbutton_.Label), formDiv)
		case "AssociationStorage":
			FormDivBasicFieldToField(&(formeditassocbutton_.AssociationStorage), formDiv)
		case "HasChanged":
			FormDivBasicFieldToField(&(formeditassocbutton_.HasChanged), formDiv)
		case "IsForSavePurpose":
			FormDivBasicFieldToField(&(formeditassocbutton_.IsForSavePurpose), formDiv)
		case "HasToolTip":
			FormDivBasicFieldToField(&(formeditassocbutton_.HasToolTip), formDiv)
		case "ToolTipText":
			FormDivBasicFieldToField(&(formeditassocbutton_.ToolTipText), formDiv)
		case "MatTooltipShowDelay":
			FormDivBasicFieldToField(&(formeditassocbutton_.MatTooltipShowDelay), formDiv)
		}
	}

	// manage the suppress operation
	if formeditassocbuttonFormCallback.formGroup.HasSuppressButtonBeenPressed {
		formeditassocbutton_.Unstage(formeditassocbuttonFormCallback.probe.stageOfInterest)
	}

	formeditassocbuttonFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.FormEditAssocButton](
		formeditassocbuttonFormCallback.probe,
	)

	// display a new form by reset the form stage
	if formeditassocbuttonFormCallback.CreationMode || formeditassocbuttonFormCallback.formGroup.HasSuppressButtonBeenPressed {
		formeditassocbuttonFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(formeditassocbuttonFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__FormEditAssocButtonFormCallback(
			nil,
			formeditassocbuttonFormCallback.probe,
			newFormGroup,
		)
		formeditassocbutton := new(models.FormEditAssocButton)
		FillUpForm(formeditassocbutton, newFormGroup, formeditassocbuttonFormCallback.probe)
		formeditassocbuttonFormCallback.probe.formStage.Commit()
	}

	formeditassocbuttonFormCallback.probe.ux_tree()
}
func __gong__New__FormFieldFormCallback(
	formfield *models.FormField,
	probe *Probe,
	formGroup *form.FormGroup,
) (formfieldFormCallback *FormFieldFormCallback) {
	formfieldFormCallback = new(FormFieldFormCallback)
	formfieldFormCallback.probe = probe
	formfieldFormCallback.formfield = formfield
	formfieldFormCallback.formGroup = formGroup

	formfieldFormCallback.CreationMode = (formfield == nil)

	return
}

type FormFieldFormCallback struct {
	formfield *models.FormField

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (formfieldFormCallback *FormFieldFormCallback) OnSave() {
	formfieldFormCallback.probe.stageOfInterest.Lock()
	defer formfieldFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("FormFieldFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	formfieldFormCallback.probe.formStage.Checkout()

	if formfieldFormCallback.formfield == nil {
		formfieldFormCallback.formfield = new(models.FormField).Stage(formfieldFormCallback.probe.stageOfInterest)
	}
	formfield_ := formfieldFormCallback.formfield
	_ = formfield_

	for _, formDiv := range formfieldFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(formfield_.Name), formDiv)
		case "InputTypeEnum":
			FormDivEnumStringFieldToField(&(formfield_.InputTypeEnum), formDiv)
		case "Label":
			FormDivBasicFieldToField(&(formfield_.Label), formDiv)
		case "Placeholder":
			FormDivBasicFieldToField(&(formfield_.Placeholder), formDiv)
		case "FormFieldString":
			FormDivSelectFieldToField(&(formfield_.FormFieldString), formfieldFormCallback.probe.stageOfInterest, formDiv)
		case "FormFieldFloat64":
			FormDivSelectFieldToField(&(formfield_.FormFieldFloat64), formfieldFormCallback.probe.stageOfInterest, formDiv)
		case "FormFieldInt":
			FormDivSelectFieldToField(&(formfield_.FormFieldInt), formfieldFormCallback.probe.stageOfInterest, formDiv)
		case "FormFieldDate":
			FormDivSelectFieldToField(&(formfield_.FormFieldDate), formfieldFormCallback.probe.stageOfInterest, formDiv)
		case "FormFieldTime":
			FormDivSelectFieldToField(&(formfield_.FormFieldTime), formfieldFormCallback.probe.stageOfInterest, formDiv)
		case "FormFieldDateTime":
			FormDivSelectFieldToField(&(formfield_.FormFieldDateTime), formfieldFormCallback.probe.stageOfInterest, formDiv)
		case "FormFieldSelect":
			FormDivSelectFieldToField(&(formfield_.FormFieldSelect), formfieldFormCallback.probe.stageOfInterest, formDiv)
		case "HasBespokeWidth":
			FormDivBasicFieldToField(&(formfield_.HasBespokeWidth), formDiv)
		case "BespokeWidthPx":
			FormDivBasicFieldToField(&(formfield_.BespokeWidthPx), formDiv)
		case "HasBespokeHeight":
			FormDivBasicFieldToField(&(formfield_.HasBespokeHeight), formDiv)
		case "BespokeHeightPx":
			FormDivBasicFieldToField(&(formfield_.BespokeHeightPx), formDiv)
		case "FormDiv:FormFields":
			// 1. Decode the AssociationStorage which contains the rowIDs of the FormDiv instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target FormDiv instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.FormDiv](formfieldFormCallback.probe.stageOfInterest)
			targetFormDivIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetFormDivIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all FormDiv instances and update their FormFields slice
			for _formdiv := range *models.GetGongstructInstancesSetFromPointerType[*models.FormDiv](formfieldFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(formfieldFormCallback.probe.stageOfInterest, _formdiv)
				
				// if FormDiv is selected
				if targetFormDivIDs[id] {
					// ensure formfield_ is in _formdiv.FormFields
					found := false
					for _, _b := range _formdiv.FormFields {
						if _b == formfield_ {
							found = true
							break
						}
					}
					if !found {
						_formdiv.FormFields = append(_formdiv.FormFields, formfield_)
						formfieldFormCallback.probe.UpdateSliceOfPointersCallback(_formdiv, "FormFields", &_formdiv.FormFields)
					}
				} else {
					// ensure formfield_ is NOT in _formdiv.FormFields
					idx := slices.Index(_formdiv.FormFields, formfield_)
					if idx != -1 {
						_formdiv.FormFields = slices.Delete(_formdiv.FormFields, idx, idx+1)
						formfieldFormCallback.probe.UpdateSliceOfPointersCallback(_formdiv, "FormFields", &_formdiv.FormFields)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if formfieldFormCallback.formGroup.HasSuppressButtonBeenPressed {
		formfield_.Unstage(formfieldFormCallback.probe.stageOfInterest)
	}

	formfieldFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.FormField](
		formfieldFormCallback.probe,
	)

	// display a new form by reset the form stage
	if formfieldFormCallback.CreationMode || formfieldFormCallback.formGroup.HasSuppressButtonBeenPressed {
		formfieldFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(formfieldFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__FormFieldFormCallback(
			nil,
			formfieldFormCallback.probe,
			newFormGroup,
		)
		formfield := new(models.FormField)
		FillUpForm(formfield, newFormGroup, formfieldFormCallback.probe)
		formfieldFormCallback.probe.formStage.Commit()
	}

	formfieldFormCallback.probe.ux_tree()
}
func __gong__New__FormFieldDateFormCallback(
	formfielddate *models.FormFieldDate,
	probe *Probe,
	formGroup *form.FormGroup,
) (formfielddateFormCallback *FormFieldDateFormCallback) {
	formfielddateFormCallback = new(FormFieldDateFormCallback)
	formfielddateFormCallback.probe = probe
	formfielddateFormCallback.formfielddate = formfielddate
	formfielddateFormCallback.formGroup = formGroup

	formfielddateFormCallback.CreationMode = (formfielddate == nil)

	return
}

type FormFieldDateFormCallback struct {
	formfielddate *models.FormFieldDate

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (formfielddateFormCallback *FormFieldDateFormCallback) OnSave() {
	formfielddateFormCallback.probe.stageOfInterest.Lock()
	defer formfielddateFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("FormFieldDateFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	formfielddateFormCallback.probe.formStage.Checkout()

	if formfielddateFormCallback.formfielddate == nil {
		formfielddateFormCallback.formfielddate = new(models.FormFieldDate).Stage(formfielddateFormCallback.probe.stageOfInterest)
	}
	formfielddate_ := formfielddateFormCallback.formfielddate
	_ = formfielddate_

	for _, formDiv := range formfielddateFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(formfielddate_.Name), formDiv)
		case "Value":
			FormDivBasicFieldToField(&(formfielddate_.Value), formDiv)
		}
	}

	// manage the suppress operation
	if formfielddateFormCallback.formGroup.HasSuppressButtonBeenPressed {
		formfielddate_.Unstage(formfielddateFormCallback.probe.stageOfInterest)
	}

	formfielddateFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.FormFieldDate](
		formfielddateFormCallback.probe,
	)

	// display a new form by reset the form stage
	if formfielddateFormCallback.CreationMode || formfielddateFormCallback.formGroup.HasSuppressButtonBeenPressed {
		formfielddateFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(formfielddateFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__FormFieldDateFormCallback(
			nil,
			formfielddateFormCallback.probe,
			newFormGroup,
		)
		formfielddate := new(models.FormFieldDate)
		FillUpForm(formfielddate, newFormGroup, formfielddateFormCallback.probe)
		formfielddateFormCallback.probe.formStage.Commit()
	}

	formfielddateFormCallback.probe.ux_tree()
}
func __gong__New__FormFieldDateTimeFormCallback(
	formfielddatetime *models.FormFieldDateTime,
	probe *Probe,
	formGroup *form.FormGroup,
) (formfielddatetimeFormCallback *FormFieldDateTimeFormCallback) {
	formfielddatetimeFormCallback = new(FormFieldDateTimeFormCallback)
	formfielddatetimeFormCallback.probe = probe
	formfielddatetimeFormCallback.formfielddatetime = formfielddatetime
	formfielddatetimeFormCallback.formGroup = formGroup

	formfielddatetimeFormCallback.CreationMode = (formfielddatetime == nil)

	return
}

type FormFieldDateTimeFormCallback struct {
	formfielddatetime *models.FormFieldDateTime

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (formfielddatetimeFormCallback *FormFieldDateTimeFormCallback) OnSave() {
	formfielddatetimeFormCallback.probe.stageOfInterest.Lock()
	defer formfielddatetimeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("FormFieldDateTimeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	formfielddatetimeFormCallback.probe.formStage.Checkout()

	if formfielddatetimeFormCallback.formfielddatetime == nil {
		formfielddatetimeFormCallback.formfielddatetime = new(models.FormFieldDateTime).Stage(formfielddatetimeFormCallback.probe.stageOfInterest)
	}
	formfielddatetime_ := formfielddatetimeFormCallback.formfielddatetime
	_ = formfielddatetime_

	for _, formDiv := range formfielddatetimeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(formfielddatetime_.Name), formDiv)
		case "Value":
			FormDivBasicFieldToField(&(formfielddatetime_.Value), formDiv)
		}
	}

	// manage the suppress operation
	if formfielddatetimeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		formfielddatetime_.Unstage(formfielddatetimeFormCallback.probe.stageOfInterest)
	}

	formfielddatetimeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.FormFieldDateTime](
		formfielddatetimeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if formfielddatetimeFormCallback.CreationMode || formfielddatetimeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		formfielddatetimeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(formfielddatetimeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__FormFieldDateTimeFormCallback(
			nil,
			formfielddatetimeFormCallback.probe,
			newFormGroup,
		)
		formfielddatetime := new(models.FormFieldDateTime)
		FillUpForm(formfielddatetime, newFormGroup, formfielddatetimeFormCallback.probe)
		formfielddatetimeFormCallback.probe.formStage.Commit()
	}

	formfielddatetimeFormCallback.probe.ux_tree()
}
func __gong__New__FormFieldFloat64FormCallback(
	formfieldfloat64 *models.FormFieldFloat64,
	probe *Probe,
	formGroup *form.FormGroup,
) (formfieldfloat64FormCallback *FormFieldFloat64FormCallback) {
	formfieldfloat64FormCallback = new(FormFieldFloat64FormCallback)
	formfieldfloat64FormCallback.probe = probe
	formfieldfloat64FormCallback.formfieldfloat64 = formfieldfloat64
	formfieldfloat64FormCallback.formGroup = formGroup

	formfieldfloat64FormCallback.CreationMode = (formfieldfloat64 == nil)

	return
}

type FormFieldFloat64FormCallback struct {
	formfieldfloat64 *models.FormFieldFloat64

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (formfieldfloat64FormCallback *FormFieldFloat64FormCallback) OnSave() {
	formfieldfloat64FormCallback.probe.stageOfInterest.Lock()
	defer formfieldfloat64FormCallback.probe.stageOfInterest.Unlock()

	// log.Println("FormFieldFloat64FormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	formfieldfloat64FormCallback.probe.formStage.Checkout()

	if formfieldfloat64FormCallback.formfieldfloat64 == nil {
		formfieldfloat64FormCallback.formfieldfloat64 = new(models.FormFieldFloat64).Stage(formfieldfloat64FormCallback.probe.stageOfInterest)
	}
	formfieldfloat64_ := formfieldfloat64FormCallback.formfieldfloat64
	_ = formfieldfloat64_

	for _, formDiv := range formfieldfloat64FormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(formfieldfloat64_.Name), formDiv)
		case "Value":
			FormDivBasicFieldToField(&(formfieldfloat64_.Value), formDiv)
		case "HasMinValidator":
			FormDivBasicFieldToField(&(formfieldfloat64_.HasMinValidator), formDiv)
		case "MinValue":
			FormDivBasicFieldToField(&(formfieldfloat64_.MinValue), formDiv)
		case "HasMaxValidator":
			FormDivBasicFieldToField(&(formfieldfloat64_.HasMaxValidator), formDiv)
		case "MaxValue":
			FormDivBasicFieldToField(&(formfieldfloat64_.MaxValue), formDiv)
		}
	}

	// manage the suppress operation
	if formfieldfloat64FormCallback.formGroup.HasSuppressButtonBeenPressed {
		formfieldfloat64_.Unstage(formfieldfloat64FormCallback.probe.stageOfInterest)
	}

	formfieldfloat64FormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.FormFieldFloat64](
		formfieldfloat64FormCallback.probe,
	)

	// display a new form by reset the form stage
	if formfieldfloat64FormCallback.CreationMode || formfieldfloat64FormCallback.formGroup.HasSuppressButtonBeenPressed {
		formfieldfloat64FormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(formfieldfloat64FormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__FormFieldFloat64FormCallback(
			nil,
			formfieldfloat64FormCallback.probe,
			newFormGroup,
		)
		formfieldfloat64 := new(models.FormFieldFloat64)
		FillUpForm(formfieldfloat64, newFormGroup, formfieldfloat64FormCallback.probe)
		formfieldfloat64FormCallback.probe.formStage.Commit()
	}

	formfieldfloat64FormCallback.probe.ux_tree()
}
func __gong__New__FormFieldIntFormCallback(
	formfieldint *models.FormFieldInt,
	probe *Probe,
	formGroup *form.FormGroup,
) (formfieldintFormCallback *FormFieldIntFormCallback) {
	formfieldintFormCallback = new(FormFieldIntFormCallback)
	formfieldintFormCallback.probe = probe
	formfieldintFormCallback.formfieldint = formfieldint
	formfieldintFormCallback.formGroup = formGroup

	formfieldintFormCallback.CreationMode = (formfieldint == nil)

	return
}

type FormFieldIntFormCallback struct {
	formfieldint *models.FormFieldInt

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (formfieldintFormCallback *FormFieldIntFormCallback) OnSave() {
	formfieldintFormCallback.probe.stageOfInterest.Lock()
	defer formfieldintFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("FormFieldIntFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	formfieldintFormCallback.probe.formStage.Checkout()

	if formfieldintFormCallback.formfieldint == nil {
		formfieldintFormCallback.formfieldint = new(models.FormFieldInt).Stage(formfieldintFormCallback.probe.stageOfInterest)
	}
	formfieldint_ := formfieldintFormCallback.formfieldint
	_ = formfieldint_

	for _, formDiv := range formfieldintFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(formfieldint_.Name), formDiv)
		case "Value":
			FormDivBasicFieldToField(&(formfieldint_.Value), formDiv)
		case "HasMinValidator":
			FormDivBasicFieldToField(&(formfieldint_.HasMinValidator), formDiv)
		case "MinValue":
			FormDivBasicFieldToField(&(formfieldint_.MinValue), formDiv)
		case "HasMaxValidator":
			FormDivBasicFieldToField(&(formfieldint_.HasMaxValidator), formDiv)
		case "MaxValue":
			FormDivBasicFieldToField(&(formfieldint_.MaxValue), formDiv)
		}
	}

	// manage the suppress operation
	if formfieldintFormCallback.formGroup.HasSuppressButtonBeenPressed {
		formfieldint_.Unstage(formfieldintFormCallback.probe.stageOfInterest)
	}

	formfieldintFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.FormFieldInt](
		formfieldintFormCallback.probe,
	)

	// display a new form by reset the form stage
	if formfieldintFormCallback.CreationMode || formfieldintFormCallback.formGroup.HasSuppressButtonBeenPressed {
		formfieldintFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(formfieldintFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__FormFieldIntFormCallback(
			nil,
			formfieldintFormCallback.probe,
			newFormGroup,
		)
		formfieldint := new(models.FormFieldInt)
		FillUpForm(formfieldint, newFormGroup, formfieldintFormCallback.probe)
		formfieldintFormCallback.probe.formStage.Commit()
	}

	formfieldintFormCallback.probe.ux_tree()
}
func __gong__New__FormFieldSelectFormCallback(
	formfieldselect *models.FormFieldSelect,
	probe *Probe,
	formGroup *form.FormGroup,
) (formfieldselectFormCallback *FormFieldSelectFormCallback) {
	formfieldselectFormCallback = new(FormFieldSelectFormCallback)
	formfieldselectFormCallback.probe = probe
	formfieldselectFormCallback.formfieldselect = formfieldselect
	formfieldselectFormCallback.formGroup = formGroup

	formfieldselectFormCallback.CreationMode = (formfieldselect == nil)

	return
}

type FormFieldSelectFormCallback struct {
	formfieldselect *models.FormFieldSelect

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (formfieldselectFormCallback *FormFieldSelectFormCallback) OnSave() {
	formfieldselectFormCallback.probe.stageOfInterest.Lock()
	defer formfieldselectFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("FormFieldSelectFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	formfieldselectFormCallback.probe.formStage.Checkout()

	if formfieldselectFormCallback.formfieldselect == nil {
		formfieldselectFormCallback.formfieldselect = new(models.FormFieldSelect).Stage(formfieldselectFormCallback.probe.stageOfInterest)
	}
	formfieldselect_ := formfieldselectFormCallback.formfieldselect
	_ = formfieldselect_

	for _, formDiv := range formfieldselectFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(formfieldselect_.Name), formDiv)
		case "Value":
			FormDivSelectFieldToField(&(formfieldselect_.Value), formfieldselectFormCallback.probe.stageOfInterest, formDiv)
		case "Options":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Option](formfieldselectFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Option, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Option)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					formfieldselectFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Option](formfieldselectFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			formfieldselect_.Options = instanceSlice
			formfieldselectFormCallback.probe.UpdateSliceOfPointersCallback(formfieldselect_, "Options", &formfieldselect_.Options)

		case "CanBeEmpty":
			FormDivBasicFieldToField(&(formfieldselect_.CanBeEmpty), formDiv)
		case "PreserveInitialOrder":
			FormDivBasicFieldToField(&(formfieldselect_.PreserveInitialOrder), formDiv)
		}
	}

	// manage the suppress operation
	if formfieldselectFormCallback.formGroup.HasSuppressButtonBeenPressed {
		formfieldselect_.Unstage(formfieldselectFormCallback.probe.stageOfInterest)
	}

	formfieldselectFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.FormFieldSelect](
		formfieldselectFormCallback.probe,
	)

	// display a new form by reset the form stage
	if formfieldselectFormCallback.CreationMode || formfieldselectFormCallback.formGroup.HasSuppressButtonBeenPressed {
		formfieldselectFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(formfieldselectFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__FormFieldSelectFormCallback(
			nil,
			formfieldselectFormCallback.probe,
			newFormGroup,
		)
		formfieldselect := new(models.FormFieldSelect)
		FillUpForm(formfieldselect, newFormGroup, formfieldselectFormCallback.probe)
		formfieldselectFormCallback.probe.formStage.Commit()
	}

	formfieldselectFormCallback.probe.ux_tree()
}
func __gong__New__FormFieldStringFormCallback(
	formfieldstring *models.FormFieldString,
	probe *Probe,
	formGroup *form.FormGroup,
) (formfieldstringFormCallback *FormFieldStringFormCallback) {
	formfieldstringFormCallback = new(FormFieldStringFormCallback)
	formfieldstringFormCallback.probe = probe
	formfieldstringFormCallback.formfieldstring = formfieldstring
	formfieldstringFormCallback.formGroup = formGroup

	formfieldstringFormCallback.CreationMode = (formfieldstring == nil)

	return
}

type FormFieldStringFormCallback struct {
	formfieldstring *models.FormFieldString

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (formfieldstringFormCallback *FormFieldStringFormCallback) OnSave() {
	formfieldstringFormCallback.probe.stageOfInterest.Lock()
	defer formfieldstringFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("FormFieldStringFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	formfieldstringFormCallback.probe.formStage.Checkout()

	if formfieldstringFormCallback.formfieldstring == nil {
		formfieldstringFormCallback.formfieldstring = new(models.FormFieldString).Stage(formfieldstringFormCallback.probe.stageOfInterest)
	}
	formfieldstring_ := formfieldstringFormCallback.formfieldstring
	_ = formfieldstring_

	for _, formDiv := range formfieldstringFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(formfieldstring_.Name), formDiv)
		case "Value":
			FormDivBasicFieldToField(&(formfieldstring_.Value), formDiv)
		case "IsTextArea":
			FormDivBasicFieldToField(&(formfieldstring_.IsTextArea), formDiv)
		}
	}

	// manage the suppress operation
	if formfieldstringFormCallback.formGroup.HasSuppressButtonBeenPressed {
		formfieldstring_.Unstage(formfieldstringFormCallback.probe.stageOfInterest)
	}

	formfieldstringFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.FormFieldString](
		formfieldstringFormCallback.probe,
	)

	// display a new form by reset the form stage
	if formfieldstringFormCallback.CreationMode || formfieldstringFormCallback.formGroup.HasSuppressButtonBeenPressed {
		formfieldstringFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(formfieldstringFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__FormFieldStringFormCallback(
			nil,
			formfieldstringFormCallback.probe,
			newFormGroup,
		)
		formfieldstring := new(models.FormFieldString)
		FillUpForm(formfieldstring, newFormGroup, formfieldstringFormCallback.probe)
		formfieldstringFormCallback.probe.formStage.Commit()
	}

	formfieldstringFormCallback.probe.ux_tree()
}
func __gong__New__FormFieldTimeFormCallback(
	formfieldtime *models.FormFieldTime,
	probe *Probe,
	formGroup *form.FormGroup,
) (formfieldtimeFormCallback *FormFieldTimeFormCallback) {
	formfieldtimeFormCallback = new(FormFieldTimeFormCallback)
	formfieldtimeFormCallback.probe = probe
	formfieldtimeFormCallback.formfieldtime = formfieldtime
	formfieldtimeFormCallback.formGroup = formGroup

	formfieldtimeFormCallback.CreationMode = (formfieldtime == nil)

	return
}

type FormFieldTimeFormCallback struct {
	formfieldtime *models.FormFieldTime

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (formfieldtimeFormCallback *FormFieldTimeFormCallback) OnSave() {
	formfieldtimeFormCallback.probe.stageOfInterest.Lock()
	defer formfieldtimeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("FormFieldTimeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	formfieldtimeFormCallback.probe.formStage.Checkout()

	if formfieldtimeFormCallback.formfieldtime == nil {
		formfieldtimeFormCallback.formfieldtime = new(models.FormFieldTime).Stage(formfieldtimeFormCallback.probe.stageOfInterest)
	}
	formfieldtime_ := formfieldtimeFormCallback.formfieldtime
	_ = formfieldtime_

	for _, formDiv := range formfieldtimeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(formfieldtime_.Name), formDiv)
		case "Value":
			FormDivBasicFieldToField(&(formfieldtime_.Value), formDiv)
		case "Step":
			FormDivBasicFieldToField(&(formfieldtime_.Step), formDiv)
		}
	}

	// manage the suppress operation
	if formfieldtimeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		formfieldtime_.Unstage(formfieldtimeFormCallback.probe.stageOfInterest)
	}

	formfieldtimeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.FormFieldTime](
		formfieldtimeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if formfieldtimeFormCallback.CreationMode || formfieldtimeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		formfieldtimeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(formfieldtimeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__FormFieldTimeFormCallback(
			nil,
			formfieldtimeFormCallback.probe,
			newFormGroup,
		)
		formfieldtime := new(models.FormFieldTime)
		FillUpForm(formfieldtime, newFormGroup, formfieldtimeFormCallback.probe)
		formfieldtimeFormCallback.probe.formStage.Commit()
	}

	formfieldtimeFormCallback.probe.ux_tree()
}
func __gong__New__FormGroupFormCallback(
	formgroup *models.FormGroup,
	probe *Probe,
	formGroup *form.FormGroup,
) (formgroupFormCallback *FormGroupFormCallback) {
	formgroupFormCallback = new(FormGroupFormCallback)
	formgroupFormCallback.probe = probe
	formgroupFormCallback.formgroup = formgroup
	formgroupFormCallback.formGroup = formGroup

	formgroupFormCallback.CreationMode = (formgroup == nil)

	return
}

type FormGroupFormCallback struct {
	formgroup *models.FormGroup

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (formgroupFormCallback *FormGroupFormCallback) OnSave() {
	formgroupFormCallback.probe.stageOfInterest.Lock()
	defer formgroupFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("FormGroupFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	formgroupFormCallback.probe.formStage.Checkout()

	if formgroupFormCallback.formgroup == nil {
		formgroupFormCallback.formgroup = new(models.FormGroup).Stage(formgroupFormCallback.probe.stageOfInterest)
	}
	formgroup_ := formgroupFormCallback.formgroup
	_ = formgroup_

	for _, formDiv := range formgroupFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(formgroup_.Name), formDiv)
		case "Label":
			FormDivBasicFieldToField(&(formgroup_.Label), formDiv)
		case "TypeLabel":
			FormDivBasicFieldToField(&(formgroup_.TypeLabel), formDiv)
		case "FormDivs":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.FormDiv](formgroupFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.FormDiv, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.FormDiv)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					formgroupFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.FormDiv](formgroupFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			formgroup_.FormDivs = instanceSlice
			formgroupFormCallback.probe.UpdateSliceOfPointersCallback(formgroup_, "FormDivs", &formgroup_.FormDivs)

		case "HasSuppressButton":
			FormDivBasicFieldToField(&(formgroup_.HasSuppressButton), formDiv)
		case "HasSuppressButtonBeenPressed":
			FormDivBasicFieldToField(&(formgroup_.HasSuppressButtonBeenPressed), formDiv)
		}
	}

	// manage the suppress operation
	if formgroupFormCallback.formGroup.HasSuppressButtonBeenPressed {
		formgroup_.Unstage(formgroupFormCallback.probe.stageOfInterest)
	}

	formgroupFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.FormGroup](
		formgroupFormCallback.probe,
	)

	// display a new form by reset the form stage
	if formgroupFormCallback.CreationMode || formgroupFormCallback.formGroup.HasSuppressButtonBeenPressed {
		formgroupFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(formgroupFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__FormGroupFormCallback(
			nil,
			formgroupFormCallback.probe,
			newFormGroup,
		)
		formgroup := new(models.FormGroup)
		FillUpForm(formgroup, newFormGroup, formgroupFormCallback.probe)
		formgroupFormCallback.probe.formStage.Commit()
	}

	formgroupFormCallback.probe.ux_tree()
}
func __gong__New__FormSortAssocButtonFormCallback(
	formsortassocbutton *models.FormSortAssocButton,
	probe *Probe,
	formGroup *form.FormGroup,
) (formsortassocbuttonFormCallback *FormSortAssocButtonFormCallback) {
	formsortassocbuttonFormCallback = new(FormSortAssocButtonFormCallback)
	formsortassocbuttonFormCallback.probe = probe
	formsortassocbuttonFormCallback.formsortassocbutton = formsortassocbutton
	formsortassocbuttonFormCallback.formGroup = formGroup

	formsortassocbuttonFormCallback.CreationMode = (formsortassocbutton == nil)

	return
}

type FormSortAssocButtonFormCallback struct {
	formsortassocbutton *models.FormSortAssocButton

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (formsortassocbuttonFormCallback *FormSortAssocButtonFormCallback) OnSave() {
	formsortassocbuttonFormCallback.probe.stageOfInterest.Lock()
	defer formsortassocbuttonFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("FormSortAssocButtonFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	formsortassocbuttonFormCallback.probe.formStage.Checkout()

	if formsortassocbuttonFormCallback.formsortassocbutton == nil {
		formsortassocbuttonFormCallback.formsortassocbutton = new(models.FormSortAssocButton).Stage(formsortassocbuttonFormCallback.probe.stageOfInterest)
	}
	formsortassocbutton_ := formsortassocbuttonFormCallback.formsortassocbutton
	_ = formsortassocbutton_

	for _, formDiv := range formsortassocbuttonFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(formsortassocbutton_.Name), formDiv)
		case "Label":
			FormDivBasicFieldToField(&(formsortassocbutton_.Label), formDiv)
		case "HasToolTip":
			FormDivBasicFieldToField(&(formsortassocbutton_.HasToolTip), formDiv)
		case "ToolTipText":
			FormDivBasicFieldToField(&(formsortassocbutton_.ToolTipText), formDiv)
		case "MatTooltipShowDelay":
			FormDivBasicFieldToField(&(formsortassocbutton_.MatTooltipShowDelay), formDiv)
		case "FormEditAssocButton":
			FormDivSelectFieldToField(&(formsortassocbutton_.FormEditAssocButton), formsortassocbuttonFormCallback.probe.stageOfInterest, formDiv)
		}
	}

	// manage the suppress operation
	if formsortassocbuttonFormCallback.formGroup.HasSuppressButtonBeenPressed {
		formsortassocbutton_.Unstage(formsortassocbuttonFormCallback.probe.stageOfInterest)
	}

	formsortassocbuttonFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.FormSortAssocButton](
		formsortassocbuttonFormCallback.probe,
	)

	// display a new form by reset the form stage
	if formsortassocbuttonFormCallback.CreationMode || formsortassocbuttonFormCallback.formGroup.HasSuppressButtonBeenPressed {
		formsortassocbuttonFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(formsortassocbuttonFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__FormSortAssocButtonFormCallback(
			nil,
			formsortassocbuttonFormCallback.probe,
			newFormGroup,
		)
		formsortassocbutton := new(models.FormSortAssocButton)
		FillUpForm(formsortassocbutton, newFormGroup, formsortassocbuttonFormCallback.probe)
		formsortassocbuttonFormCallback.probe.formStage.Commit()
	}

	formsortassocbuttonFormCallback.probe.ux_tree()
}
func __gong__New__OptionFormCallback(
	option *models.Option,
	probe *Probe,
	formGroup *form.FormGroup,
) (optionFormCallback *OptionFormCallback) {
	optionFormCallback = new(OptionFormCallback)
	optionFormCallback.probe = probe
	optionFormCallback.option = option
	optionFormCallback.formGroup = formGroup

	optionFormCallback.CreationMode = (option == nil)

	return
}

type OptionFormCallback struct {
	option *models.Option

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (optionFormCallback *OptionFormCallback) OnSave() {
	optionFormCallback.probe.stageOfInterest.Lock()
	defer optionFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("OptionFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	optionFormCallback.probe.formStage.Checkout()

	if optionFormCallback.option == nil {
		optionFormCallback.option = new(models.Option).Stage(optionFormCallback.probe.stageOfInterest)
	}
	option_ := optionFormCallback.option
	_ = option_

	for _, formDiv := range optionFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(option_.Name), formDiv)
		case "FormFieldSelect:Options":
			// 1. Decode the AssociationStorage which contains the rowIDs of the FormFieldSelect instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target FormFieldSelect instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.FormFieldSelect](optionFormCallback.probe.stageOfInterest)
			targetFormFieldSelectIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetFormFieldSelectIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all FormFieldSelect instances and update their Options slice
			for _formfieldselect := range *models.GetGongstructInstancesSetFromPointerType[*models.FormFieldSelect](optionFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(optionFormCallback.probe.stageOfInterest, _formfieldselect)
				
				// if FormFieldSelect is selected
				if targetFormFieldSelectIDs[id] {
					// ensure option_ is in _formfieldselect.Options
					found := false
					for _, _b := range _formfieldselect.Options {
						if _b == option_ {
							found = true
							break
						}
					}
					if !found {
						_formfieldselect.Options = append(_formfieldselect.Options, option_)
						optionFormCallback.probe.UpdateSliceOfPointersCallback(_formfieldselect, "Options", &_formfieldselect.Options)
					}
				} else {
					// ensure option_ is NOT in _formfieldselect.Options
					idx := slices.Index(_formfieldselect.Options, option_)
					if idx != -1 {
						_formfieldselect.Options = slices.Delete(_formfieldselect.Options, idx, idx+1)
						optionFormCallback.probe.UpdateSliceOfPointersCallback(_formfieldselect, "Options", &_formfieldselect.Options)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if optionFormCallback.formGroup.HasSuppressButtonBeenPressed {
		option_.Unstage(optionFormCallback.probe.stageOfInterest)
	}

	optionFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Option](
		optionFormCallback.probe,
	)

	// display a new form by reset the form stage
	if optionFormCallback.CreationMode || optionFormCallback.formGroup.HasSuppressButtonBeenPressed {
		optionFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(optionFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__OptionFormCallback(
			nil,
			optionFormCallback.probe,
			newFormGroup,
		)
		option := new(models.Option)
		FillUpForm(option, newFormGroup, optionFormCallback.probe)
		optionFormCallback.probe.formStage.Commit()
	}

	optionFormCallback.probe.ux_tree()
}
