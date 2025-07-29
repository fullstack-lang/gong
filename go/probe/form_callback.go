// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	table "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gong/go/models"
)

// to avoid errors when time and slices packages are not used in the generated code
const _ = time.Nanosecond

var _ = slices.Delete([]string{"a"}, 0, 1)

var _ = log.Panicf

// insertion point
func __gong__New__GongBasicFieldFormCallback(
	gongbasicfield *models.GongBasicField,
	probe *Probe,
	formGroup *table.FormGroup,
) (gongbasicfieldFormCallback *GongBasicFieldFormCallback) {
	gongbasicfieldFormCallback = new(GongBasicFieldFormCallback)
	gongbasicfieldFormCallback.probe = probe
	gongbasicfieldFormCallback.gongbasicfield = gongbasicfield
	gongbasicfieldFormCallback.formGroup = formGroup

	gongbasicfieldFormCallback.CreationMode = (gongbasicfield == nil)

	return
}

type GongBasicFieldFormCallback struct {
	gongbasicfield *models.GongBasicField

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (gongbasicfieldFormCallback *GongBasicFieldFormCallback) OnSave() {

	// log.Println("GongBasicFieldFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	gongbasicfieldFormCallback.probe.formStage.Checkout()

	if gongbasicfieldFormCallback.gongbasicfield == nil {
		gongbasicfieldFormCallback.gongbasicfield = new(models.GongBasicField).Stage(gongbasicfieldFormCallback.probe.stageOfInterest)
	}
	gongbasicfield_ := gongbasicfieldFormCallback.gongbasicfield
	_ = gongbasicfield_

	for _, formDiv := range gongbasicfieldFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(gongbasicfield_.Name), formDiv)
		case "BasicKindName":
			FormDivBasicFieldToField(&(gongbasicfield_.BasicKindName), formDiv)
		case "GongEnum":
			FormDivSelectFieldToField(&(gongbasicfield_.GongEnum), gongbasicfieldFormCallback.probe.stageOfInterest, formDiv)
		case "DeclaredType":
			FormDivBasicFieldToField(&(gongbasicfield_.DeclaredType), formDiv)
		case "CompositeStructName":
			FormDivBasicFieldToField(&(gongbasicfield_.CompositeStructName), formDiv)
		case "Index":
			FormDivBasicFieldToField(&(gongbasicfield_.Index), formDiv)
		case "IsTextArea":
			FormDivBasicFieldToField(&(gongbasicfield_.IsTextArea), formDiv)
		case "IsBespokeWidth":
			FormDivBasicFieldToField(&(gongbasicfield_.IsBespokeWidth), formDiv)
		case "BespokeWidth":
			FormDivBasicFieldToField(&(gongbasicfield_.BespokeWidth), formDiv)
		case "IsBespokeHeight":
			FormDivBasicFieldToField(&(gongbasicfield_.IsBespokeHeight), formDiv)
		case "BespokeHeight":
			FormDivBasicFieldToField(&(gongbasicfield_.BespokeHeight), formDiv)
		case "GongStruct:GongBasicFields":
			// WARNING : this form deals with the N-N association "GongStruct.GongBasicFields []*GongBasicField" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of GongBasicField). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.GongStruct
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "GongStruct"
				rf.Fieldname = "GongBasicFields"
				formerAssociationSource := models.GetReverseFieldOwner(
					gongbasicfieldFormCallback.probe.stageOfInterest,
					gongbasicfield_,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.GongStruct)
					if !ok {
						log.Fatalln("Source of GongStruct.GongBasicFields []*GongBasicField, is not an GongStruct instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.GongBasicFields, gongbasicfield_)
					formerSource.GongBasicFields = slices.Delete(formerSource.GongBasicFields, idx, idx+1)
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
			var newSource *models.GongStruct
			for _gongstruct := range *models.GetGongstructInstancesSet[models.GongStruct](gongbasicfieldFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _gongstruct.GetName() == newSourceName.GetName() {
					newSource = _gongstruct // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of GongStruct.GongBasicFields []*GongBasicField, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.GongBasicFields = append(newSource.GongBasicFields, gongbasicfield_)
		}
	}

	// manage the suppress operation
	if gongbasicfieldFormCallback.formGroup.HasSuppressButtonBeenPressed {
		gongbasicfield_.Unstage(gongbasicfieldFormCallback.probe.stageOfInterest)
	}

	gongbasicfieldFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.GongBasicField](
		gongbasicfieldFormCallback.probe,
	)
	gongbasicfieldFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if gongbasicfieldFormCallback.CreationMode || gongbasicfieldFormCallback.formGroup.HasSuppressButtonBeenPressed {
		gongbasicfieldFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(gongbasicfieldFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__GongBasicFieldFormCallback(
			nil,
			gongbasicfieldFormCallback.probe,
			newFormGroup,
		)
		gongbasicfield := new(models.GongBasicField)
		FillUpForm(gongbasicfield, newFormGroup, gongbasicfieldFormCallback.probe)
		gongbasicfieldFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(gongbasicfieldFormCallback.probe)
}
func __gong__New__GongEnumFormCallback(
	gongenum *models.GongEnum,
	probe *Probe,
	formGroup *table.FormGroup,
) (gongenumFormCallback *GongEnumFormCallback) {
	gongenumFormCallback = new(GongEnumFormCallback)
	gongenumFormCallback.probe = probe
	gongenumFormCallback.gongenum = gongenum
	gongenumFormCallback.formGroup = formGroup

	gongenumFormCallback.CreationMode = (gongenum == nil)

	return
}

type GongEnumFormCallback struct {
	gongenum *models.GongEnum

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (gongenumFormCallback *GongEnumFormCallback) OnSave() {

	// log.Println("GongEnumFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	gongenumFormCallback.probe.formStage.Checkout()

	if gongenumFormCallback.gongenum == nil {
		gongenumFormCallback.gongenum = new(models.GongEnum).Stage(gongenumFormCallback.probe.stageOfInterest)
	}
	gongenum_ := gongenumFormCallback.gongenum
	_ = gongenum_

	for _, formDiv := range gongenumFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(gongenum_.Name), formDiv)
		case "Type":
			FormDivEnumIntFieldToField(&(gongenum_.Type), formDiv)
		case "GongEnumValues":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.GongEnumValue](gongenumFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.GongEnumValue, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.GongEnumValue)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					gongenumFormCallback.probe.stageOfInterest,
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
			gongenum_.GongEnumValues = instanceSlice

		}
	}

	// manage the suppress operation
	if gongenumFormCallback.formGroup.HasSuppressButtonBeenPressed {
		gongenum_.Unstage(gongenumFormCallback.probe.stageOfInterest)
	}

	gongenumFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.GongEnum](
		gongenumFormCallback.probe,
	)
	gongenumFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if gongenumFormCallback.CreationMode || gongenumFormCallback.formGroup.HasSuppressButtonBeenPressed {
		gongenumFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(gongenumFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__GongEnumFormCallback(
			nil,
			gongenumFormCallback.probe,
			newFormGroup,
		)
		gongenum := new(models.GongEnum)
		FillUpForm(gongenum, newFormGroup, gongenumFormCallback.probe)
		gongenumFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(gongenumFormCallback.probe)
}
func __gong__New__GongEnumValueFormCallback(
	gongenumvalue *models.GongEnumValue,
	probe *Probe,
	formGroup *table.FormGroup,
) (gongenumvalueFormCallback *GongEnumValueFormCallback) {
	gongenumvalueFormCallback = new(GongEnumValueFormCallback)
	gongenumvalueFormCallback.probe = probe
	gongenumvalueFormCallback.gongenumvalue = gongenumvalue
	gongenumvalueFormCallback.formGroup = formGroup

	gongenumvalueFormCallback.CreationMode = (gongenumvalue == nil)

	return
}

type GongEnumValueFormCallback struct {
	gongenumvalue *models.GongEnumValue

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (gongenumvalueFormCallback *GongEnumValueFormCallback) OnSave() {

	// log.Println("GongEnumValueFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	gongenumvalueFormCallback.probe.formStage.Checkout()

	if gongenumvalueFormCallback.gongenumvalue == nil {
		gongenumvalueFormCallback.gongenumvalue = new(models.GongEnumValue).Stage(gongenumvalueFormCallback.probe.stageOfInterest)
	}
	gongenumvalue_ := gongenumvalueFormCallback.gongenumvalue
	_ = gongenumvalue_

	for _, formDiv := range gongenumvalueFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(gongenumvalue_.Name), formDiv)
		case "Value":
			FormDivBasicFieldToField(&(gongenumvalue_.Value), formDiv)
		case "GongEnum:GongEnumValues":
			// WARNING : this form deals with the N-N association "GongEnum.GongEnumValues []*GongEnumValue" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of GongEnumValue). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.GongEnum
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "GongEnum"
				rf.Fieldname = "GongEnumValues"
				formerAssociationSource := models.GetReverseFieldOwner(
					gongenumvalueFormCallback.probe.stageOfInterest,
					gongenumvalue_,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.GongEnum)
					if !ok {
						log.Fatalln("Source of GongEnum.GongEnumValues []*GongEnumValue, is not an GongEnum instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.GongEnumValues, gongenumvalue_)
					formerSource.GongEnumValues = slices.Delete(formerSource.GongEnumValues, idx, idx+1)
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
			var newSource *models.GongEnum
			for _gongenum := range *models.GetGongstructInstancesSet[models.GongEnum](gongenumvalueFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _gongenum.GetName() == newSourceName.GetName() {
					newSource = _gongenum // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of GongEnum.GongEnumValues []*GongEnumValue, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.GongEnumValues = append(newSource.GongEnumValues, gongenumvalue_)
		}
	}

	// manage the suppress operation
	if gongenumvalueFormCallback.formGroup.HasSuppressButtonBeenPressed {
		gongenumvalue_.Unstage(gongenumvalueFormCallback.probe.stageOfInterest)
	}

	gongenumvalueFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.GongEnumValue](
		gongenumvalueFormCallback.probe,
	)
	gongenumvalueFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if gongenumvalueFormCallback.CreationMode || gongenumvalueFormCallback.formGroup.HasSuppressButtonBeenPressed {
		gongenumvalueFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(gongenumvalueFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__GongEnumValueFormCallback(
			nil,
			gongenumvalueFormCallback.probe,
			newFormGroup,
		)
		gongenumvalue := new(models.GongEnumValue)
		FillUpForm(gongenumvalue, newFormGroup, gongenumvalueFormCallback.probe)
		gongenumvalueFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(gongenumvalueFormCallback.probe)
}
func __gong__New__GongLinkFormCallback(
	gonglink *models.GongLink,
	probe *Probe,
	formGroup *table.FormGroup,
) (gonglinkFormCallback *GongLinkFormCallback) {
	gonglinkFormCallback = new(GongLinkFormCallback)
	gonglinkFormCallback.probe = probe
	gonglinkFormCallback.gonglink = gonglink
	gonglinkFormCallback.formGroup = formGroup

	gonglinkFormCallback.CreationMode = (gonglink == nil)

	return
}

type GongLinkFormCallback struct {
	gonglink *models.GongLink

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (gonglinkFormCallback *GongLinkFormCallback) OnSave() {

	// log.Println("GongLinkFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	gonglinkFormCallback.probe.formStage.Checkout()

	if gonglinkFormCallback.gonglink == nil {
		gonglinkFormCallback.gonglink = new(models.GongLink).Stage(gonglinkFormCallback.probe.stageOfInterest)
	}
	gonglink_ := gonglinkFormCallback.gonglink
	_ = gonglink_

	for _, formDiv := range gonglinkFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(gonglink_.Name), formDiv)
		case "Recv":
			FormDivBasicFieldToField(&(gonglink_.Recv), formDiv)
		case "ImportPath":
			FormDivBasicFieldToField(&(gonglink_.ImportPath), formDiv)
		case "GongNote:Links":
			// WARNING : this form deals with the N-N association "GongNote.Links []*GongLink" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of GongLink). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.GongNote
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "GongNote"
				rf.Fieldname = "Links"
				formerAssociationSource := models.GetReverseFieldOwner(
					gonglinkFormCallback.probe.stageOfInterest,
					gonglink_,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.GongNote)
					if !ok {
						log.Fatalln("Source of GongNote.Links []*GongLink, is not an GongNote instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.Links, gonglink_)
					formerSource.Links = slices.Delete(formerSource.Links, idx, idx+1)
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
			var newSource *models.GongNote
			for _gongnote := range *models.GetGongstructInstancesSet[models.GongNote](gonglinkFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _gongnote.GetName() == newSourceName.GetName() {
					newSource = _gongnote // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of GongNote.Links []*GongLink, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.Links = append(newSource.Links, gonglink_)
		}
	}

	// manage the suppress operation
	if gonglinkFormCallback.formGroup.HasSuppressButtonBeenPressed {
		gonglink_.Unstage(gonglinkFormCallback.probe.stageOfInterest)
	}

	gonglinkFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.GongLink](
		gonglinkFormCallback.probe,
	)
	gonglinkFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if gonglinkFormCallback.CreationMode || gonglinkFormCallback.formGroup.HasSuppressButtonBeenPressed {
		gonglinkFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(gonglinkFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__GongLinkFormCallback(
			nil,
			gonglinkFormCallback.probe,
			newFormGroup,
		)
		gonglink := new(models.GongLink)
		FillUpForm(gonglink, newFormGroup, gonglinkFormCallback.probe)
		gonglinkFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(gonglinkFormCallback.probe)
}
func __gong__New__GongNoteFormCallback(
	gongnote *models.GongNote,
	probe *Probe,
	formGroup *table.FormGroup,
) (gongnoteFormCallback *GongNoteFormCallback) {
	gongnoteFormCallback = new(GongNoteFormCallback)
	gongnoteFormCallback.probe = probe
	gongnoteFormCallback.gongnote = gongnote
	gongnoteFormCallback.formGroup = formGroup

	gongnoteFormCallback.CreationMode = (gongnote == nil)

	return
}

type GongNoteFormCallback struct {
	gongnote *models.GongNote

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (gongnoteFormCallback *GongNoteFormCallback) OnSave() {

	// log.Println("GongNoteFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	gongnoteFormCallback.probe.formStage.Checkout()

	if gongnoteFormCallback.gongnote == nil {
		gongnoteFormCallback.gongnote = new(models.GongNote).Stage(gongnoteFormCallback.probe.stageOfInterest)
	}
	gongnote_ := gongnoteFormCallback.gongnote
	_ = gongnote_

	for _, formDiv := range gongnoteFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(gongnote_.Name), formDiv)
		case "Body":
			FormDivBasicFieldToField(&(gongnote_.Body), formDiv)
		case "BodyHTML":
			FormDivBasicFieldToField(&(gongnote_.BodyHTML), formDiv)
		case "Links":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.GongLink](gongnoteFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.GongLink, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.GongLink)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					gongnoteFormCallback.probe.stageOfInterest,
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
			gongnote_.Links = instanceSlice

		}
	}

	// manage the suppress operation
	if gongnoteFormCallback.formGroup.HasSuppressButtonBeenPressed {
		gongnote_.Unstage(gongnoteFormCallback.probe.stageOfInterest)
	}

	gongnoteFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.GongNote](
		gongnoteFormCallback.probe,
	)
	gongnoteFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if gongnoteFormCallback.CreationMode || gongnoteFormCallback.formGroup.HasSuppressButtonBeenPressed {
		gongnoteFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(gongnoteFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__GongNoteFormCallback(
			nil,
			gongnoteFormCallback.probe,
			newFormGroup,
		)
		gongnote := new(models.GongNote)
		FillUpForm(gongnote, newFormGroup, gongnoteFormCallback.probe)
		gongnoteFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(gongnoteFormCallback.probe)
}
func __gong__New__GongStructFormCallback(
	gongstruct *models.GongStruct,
	probe *Probe,
	formGroup *table.FormGroup,
) (gongstructFormCallback *GongStructFormCallback) {
	gongstructFormCallback = new(GongStructFormCallback)
	gongstructFormCallback.probe = probe
	gongstructFormCallback.gongstruct = gongstruct
	gongstructFormCallback.formGroup = formGroup

	gongstructFormCallback.CreationMode = (gongstruct == nil)

	return
}

type GongStructFormCallback struct {
	gongstruct *models.GongStruct

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (gongstructFormCallback *GongStructFormCallback) OnSave() {

	// log.Println("GongStructFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	gongstructFormCallback.probe.formStage.Checkout()

	if gongstructFormCallback.gongstruct == nil {
		gongstructFormCallback.gongstruct = new(models.GongStruct).Stage(gongstructFormCallback.probe.stageOfInterest)
	}
	gongstruct_ := gongstructFormCallback.gongstruct
	_ = gongstruct_

	for _, formDiv := range gongstructFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(gongstruct_.Name), formDiv)
		case "GongBasicFields":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.GongBasicField](gongstructFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.GongBasicField, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.GongBasicField)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					gongstructFormCallback.probe.stageOfInterest,
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
			gongstruct_.GongBasicFields = instanceSlice

		case "GongTimeFields":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.GongTimeField](gongstructFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.GongTimeField, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.GongTimeField)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					gongstructFormCallback.probe.stageOfInterest,
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
			gongstruct_.GongTimeFields = instanceSlice

		case "PointerToGongStructFields":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.PointerToGongStructField](gongstructFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.PointerToGongStructField, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.PointerToGongStructField)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					gongstructFormCallback.probe.stageOfInterest,
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
			gongstruct_.PointerToGongStructFields = instanceSlice

		case "SliceOfPointerToGongStructFields":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.SliceOfPointerToGongStructField](gongstructFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.SliceOfPointerToGongStructField, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.SliceOfPointerToGongStructField)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					gongstructFormCallback.probe.stageOfInterest,
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
			gongstruct_.SliceOfPointerToGongStructFields = instanceSlice

		case "HasOnAfterUpdateSignature":
			FormDivBasicFieldToField(&(gongstruct_.HasOnAfterUpdateSignature), formDiv)
		case "IsIgnoredForFront":
			FormDivBasicFieldToField(&(gongstruct_.IsIgnoredForFront), formDiv)
		}
	}

	// manage the suppress operation
	if gongstructFormCallback.formGroup.HasSuppressButtonBeenPressed {
		gongstruct_.Unstage(gongstructFormCallback.probe.stageOfInterest)
	}

	gongstructFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.GongStruct](
		gongstructFormCallback.probe,
	)
	gongstructFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if gongstructFormCallback.CreationMode || gongstructFormCallback.formGroup.HasSuppressButtonBeenPressed {
		gongstructFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(gongstructFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__GongStructFormCallback(
			nil,
			gongstructFormCallback.probe,
			newFormGroup,
		)
		gongstruct := new(models.GongStruct)
		FillUpForm(gongstruct, newFormGroup, gongstructFormCallback.probe)
		gongstructFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(gongstructFormCallback.probe)
}
func __gong__New__GongTimeFieldFormCallback(
	gongtimefield *models.GongTimeField,
	probe *Probe,
	formGroup *table.FormGroup,
) (gongtimefieldFormCallback *GongTimeFieldFormCallback) {
	gongtimefieldFormCallback = new(GongTimeFieldFormCallback)
	gongtimefieldFormCallback.probe = probe
	gongtimefieldFormCallback.gongtimefield = gongtimefield
	gongtimefieldFormCallback.formGroup = formGroup

	gongtimefieldFormCallback.CreationMode = (gongtimefield == nil)

	return
}

type GongTimeFieldFormCallback struct {
	gongtimefield *models.GongTimeField

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (gongtimefieldFormCallback *GongTimeFieldFormCallback) OnSave() {

	// log.Println("GongTimeFieldFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	gongtimefieldFormCallback.probe.formStage.Checkout()

	if gongtimefieldFormCallback.gongtimefield == nil {
		gongtimefieldFormCallback.gongtimefield = new(models.GongTimeField).Stage(gongtimefieldFormCallback.probe.stageOfInterest)
	}
	gongtimefield_ := gongtimefieldFormCallback.gongtimefield
	_ = gongtimefield_

	for _, formDiv := range gongtimefieldFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(gongtimefield_.Name), formDiv)
		case "Index":
			FormDivBasicFieldToField(&(gongtimefield_.Index), formDiv)
		case "CompositeStructName":
			FormDivBasicFieldToField(&(gongtimefield_.CompositeStructName), formDiv)
		case "BespokeTimeFormat":
			FormDivBasicFieldToField(&(gongtimefield_.BespokeTimeFormat), formDiv)
		case "GongStruct:GongTimeFields":
			// WARNING : this form deals with the N-N association "GongStruct.GongTimeFields []*GongTimeField" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of GongTimeField). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.GongStruct
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "GongStruct"
				rf.Fieldname = "GongTimeFields"
				formerAssociationSource := models.GetReverseFieldOwner(
					gongtimefieldFormCallback.probe.stageOfInterest,
					gongtimefield_,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.GongStruct)
					if !ok {
						log.Fatalln("Source of GongStruct.GongTimeFields []*GongTimeField, is not an GongStruct instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.GongTimeFields, gongtimefield_)
					formerSource.GongTimeFields = slices.Delete(formerSource.GongTimeFields, idx, idx+1)
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
			var newSource *models.GongStruct
			for _gongstruct := range *models.GetGongstructInstancesSet[models.GongStruct](gongtimefieldFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _gongstruct.GetName() == newSourceName.GetName() {
					newSource = _gongstruct // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of GongStruct.GongTimeFields []*GongTimeField, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.GongTimeFields = append(newSource.GongTimeFields, gongtimefield_)
		}
	}

	// manage the suppress operation
	if gongtimefieldFormCallback.formGroup.HasSuppressButtonBeenPressed {
		gongtimefield_.Unstage(gongtimefieldFormCallback.probe.stageOfInterest)
	}

	gongtimefieldFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.GongTimeField](
		gongtimefieldFormCallback.probe,
	)
	gongtimefieldFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if gongtimefieldFormCallback.CreationMode || gongtimefieldFormCallback.formGroup.HasSuppressButtonBeenPressed {
		gongtimefieldFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(gongtimefieldFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__GongTimeFieldFormCallback(
			nil,
			gongtimefieldFormCallback.probe,
			newFormGroup,
		)
		gongtimefield := new(models.GongTimeField)
		FillUpForm(gongtimefield, newFormGroup, gongtimefieldFormCallback.probe)
		gongtimefieldFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(gongtimefieldFormCallback.probe)
}
func __gong__New__MetaReferenceFormCallback(
	metareference *models.MetaReference,
	probe *Probe,
	formGroup *table.FormGroup,
) (metareferenceFormCallback *MetaReferenceFormCallback) {
	metareferenceFormCallback = new(MetaReferenceFormCallback)
	metareferenceFormCallback.probe = probe
	metareferenceFormCallback.metareference = metareference
	metareferenceFormCallback.formGroup = formGroup

	metareferenceFormCallback.CreationMode = (metareference == nil)

	return
}

type MetaReferenceFormCallback struct {
	metareference *models.MetaReference

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (metareferenceFormCallback *MetaReferenceFormCallback) OnSave() {

	// log.Println("MetaReferenceFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	metareferenceFormCallback.probe.formStage.Checkout()

	if metareferenceFormCallback.metareference == nil {
		metareferenceFormCallback.metareference = new(models.MetaReference).Stage(metareferenceFormCallback.probe.stageOfInterest)
	}
	metareference_ := metareferenceFormCallback.metareference
	_ = metareference_

	for _, formDiv := range metareferenceFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(metareference_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if metareferenceFormCallback.formGroup.HasSuppressButtonBeenPressed {
		metareference_.Unstage(metareferenceFormCallback.probe.stageOfInterest)
	}

	metareferenceFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.MetaReference](
		metareferenceFormCallback.probe,
	)
	metareferenceFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if metareferenceFormCallback.CreationMode || metareferenceFormCallback.formGroup.HasSuppressButtonBeenPressed {
		metareferenceFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(metareferenceFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__MetaReferenceFormCallback(
			nil,
			metareferenceFormCallback.probe,
			newFormGroup,
		)
		metareference := new(models.MetaReference)
		FillUpForm(metareference, newFormGroup, metareferenceFormCallback.probe)
		metareferenceFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(metareferenceFormCallback.probe)
}
func __gong__New__ModelPkgFormCallback(
	modelpkg *models.ModelPkg,
	probe *Probe,
	formGroup *table.FormGroup,
) (modelpkgFormCallback *ModelPkgFormCallback) {
	modelpkgFormCallback = new(ModelPkgFormCallback)
	modelpkgFormCallback.probe = probe
	modelpkgFormCallback.modelpkg = modelpkg
	modelpkgFormCallback.formGroup = formGroup

	modelpkgFormCallback.CreationMode = (modelpkg == nil)

	return
}

type ModelPkgFormCallback struct {
	modelpkg *models.ModelPkg

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (modelpkgFormCallback *ModelPkgFormCallback) OnSave() {

	// log.Println("ModelPkgFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	modelpkgFormCallback.probe.formStage.Checkout()

	if modelpkgFormCallback.modelpkg == nil {
		modelpkgFormCallback.modelpkg = new(models.ModelPkg).Stage(modelpkgFormCallback.probe.stageOfInterest)
	}
	modelpkg_ := modelpkgFormCallback.modelpkg
	_ = modelpkg_

	for _, formDiv := range modelpkgFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(modelpkg_.Name), formDiv)
		case "PkgPath":
			FormDivBasicFieldToField(&(modelpkg_.PkgPath), formDiv)
		case "PathToGoSubDirectory":
			FormDivBasicFieldToField(&(modelpkg_.PathToGoSubDirectory), formDiv)
		case "OrmPkgGenPath":
			FormDivBasicFieldToField(&(modelpkg_.OrmPkgGenPath), formDiv)
		case "DbOrmPkgGenPath":
			FormDivBasicFieldToField(&(modelpkg_.DbOrmPkgGenPath), formDiv)
		case "DbLiteOrmPkgGenPath":
			FormDivBasicFieldToField(&(modelpkg_.DbLiteOrmPkgGenPath), formDiv)
		case "DbPkgGenPath":
			FormDivBasicFieldToField(&(modelpkg_.DbPkgGenPath), formDiv)
		case "ControllersPkgGenPath":
			FormDivBasicFieldToField(&(modelpkg_.ControllersPkgGenPath), formDiv)
		case "FullstackPkgGenPath":
			FormDivBasicFieldToField(&(modelpkg_.FullstackPkgGenPath), formDiv)
		case "StackPkgGenPath":
			FormDivBasicFieldToField(&(modelpkg_.StackPkgGenPath), formDiv)
		case "StaticPkgGenPath":
			FormDivBasicFieldToField(&(modelpkg_.StaticPkgGenPath), formDiv)
		case "ProbePkgGenPath":
			FormDivBasicFieldToField(&(modelpkg_.ProbePkgGenPath), formDiv)
		case "NgWorkspacePath":
			FormDivBasicFieldToField(&(modelpkg_.NgWorkspacePath), formDiv)
		case "NgWorkspaceName":
			FormDivBasicFieldToField(&(modelpkg_.NgWorkspaceName), formDiv)
		case "NgDataLibrarySourceCodeDirectory":
			FormDivBasicFieldToField(&(modelpkg_.NgDataLibrarySourceCodeDirectory), formDiv)
		case "NgSpecificLibrarySourceCodeDirectory":
			FormDivBasicFieldToField(&(modelpkg_.NgSpecificLibrarySourceCodeDirectory), formDiv)
		case "MaterialLibDatamodelTargetPath":
			FormDivBasicFieldToField(&(modelpkg_.MaterialLibDatamodelTargetPath), formDiv)
		}
	}

	// manage the suppress operation
	if modelpkgFormCallback.formGroup.HasSuppressButtonBeenPressed {
		modelpkg_.Unstage(modelpkgFormCallback.probe.stageOfInterest)
	}

	modelpkgFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.ModelPkg](
		modelpkgFormCallback.probe,
	)
	modelpkgFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if modelpkgFormCallback.CreationMode || modelpkgFormCallback.formGroup.HasSuppressButtonBeenPressed {
		modelpkgFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(modelpkgFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ModelPkgFormCallback(
			nil,
			modelpkgFormCallback.probe,
			newFormGroup,
		)
		modelpkg := new(models.ModelPkg)
		FillUpForm(modelpkg, newFormGroup, modelpkgFormCallback.probe)
		modelpkgFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(modelpkgFormCallback.probe)
}
func __gong__New__PointerToGongStructFieldFormCallback(
	pointertogongstructfield *models.PointerToGongStructField,
	probe *Probe,
	formGroup *table.FormGroup,
) (pointertogongstructfieldFormCallback *PointerToGongStructFieldFormCallback) {
	pointertogongstructfieldFormCallback = new(PointerToGongStructFieldFormCallback)
	pointertogongstructfieldFormCallback.probe = probe
	pointertogongstructfieldFormCallback.pointertogongstructfield = pointertogongstructfield
	pointertogongstructfieldFormCallback.formGroup = formGroup

	pointertogongstructfieldFormCallback.CreationMode = (pointertogongstructfield == nil)

	return
}

type PointerToGongStructFieldFormCallback struct {
	pointertogongstructfield *models.PointerToGongStructField

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (pointertogongstructfieldFormCallback *PointerToGongStructFieldFormCallback) OnSave() {

	// log.Println("PointerToGongStructFieldFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	pointertogongstructfieldFormCallback.probe.formStage.Checkout()

	if pointertogongstructfieldFormCallback.pointertogongstructfield == nil {
		pointertogongstructfieldFormCallback.pointertogongstructfield = new(models.PointerToGongStructField).Stage(pointertogongstructfieldFormCallback.probe.stageOfInterest)
	}
	pointertogongstructfield_ := pointertogongstructfieldFormCallback.pointertogongstructfield
	_ = pointertogongstructfield_

	for _, formDiv := range pointertogongstructfieldFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(pointertogongstructfield_.Name), formDiv)
		case "GongStruct":
			FormDivSelectFieldToField(&(pointertogongstructfield_.GongStruct), pointertogongstructfieldFormCallback.probe.stageOfInterest, formDiv)
		case "Index":
			FormDivBasicFieldToField(&(pointertogongstructfield_.Index), formDiv)
		case "CompositeStructName":
			FormDivBasicFieldToField(&(pointertogongstructfield_.CompositeStructName), formDiv)
		case "IsType":
			FormDivBasicFieldToField(&(pointertogongstructfield_.IsType), formDiv)
		case "GongStruct:PointerToGongStructFields":
			// WARNING : this form deals with the N-N association "GongStruct.PointerToGongStructFields []*PointerToGongStructField" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of PointerToGongStructField). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.GongStruct
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "GongStruct"
				rf.Fieldname = "PointerToGongStructFields"
				formerAssociationSource := models.GetReverseFieldOwner(
					pointertogongstructfieldFormCallback.probe.stageOfInterest,
					pointertogongstructfield_,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.GongStruct)
					if !ok {
						log.Fatalln("Source of GongStruct.PointerToGongStructFields []*PointerToGongStructField, is not an GongStruct instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.PointerToGongStructFields, pointertogongstructfield_)
					formerSource.PointerToGongStructFields = slices.Delete(formerSource.PointerToGongStructFields, idx, idx+1)
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
			var newSource *models.GongStruct
			for _gongstruct := range *models.GetGongstructInstancesSet[models.GongStruct](pointertogongstructfieldFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _gongstruct.GetName() == newSourceName.GetName() {
					newSource = _gongstruct // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of GongStruct.PointerToGongStructFields []*PointerToGongStructField, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.PointerToGongStructFields = append(newSource.PointerToGongStructFields, pointertogongstructfield_)
		}
	}

	// manage the suppress operation
	if pointertogongstructfieldFormCallback.formGroup.HasSuppressButtonBeenPressed {
		pointertogongstructfield_.Unstage(pointertogongstructfieldFormCallback.probe.stageOfInterest)
	}

	pointertogongstructfieldFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.PointerToGongStructField](
		pointertogongstructfieldFormCallback.probe,
	)
	pointertogongstructfieldFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if pointertogongstructfieldFormCallback.CreationMode || pointertogongstructfieldFormCallback.formGroup.HasSuppressButtonBeenPressed {
		pointertogongstructfieldFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(pointertogongstructfieldFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__PointerToGongStructFieldFormCallback(
			nil,
			pointertogongstructfieldFormCallback.probe,
			newFormGroup,
		)
		pointertogongstructfield := new(models.PointerToGongStructField)
		FillUpForm(pointertogongstructfield, newFormGroup, pointertogongstructfieldFormCallback.probe)
		pointertogongstructfieldFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(pointertogongstructfieldFormCallback.probe)
}
func __gong__New__SliceOfPointerToGongStructFieldFormCallback(
	sliceofpointertogongstructfield *models.SliceOfPointerToGongStructField,
	probe *Probe,
	formGroup *table.FormGroup,
) (sliceofpointertogongstructfieldFormCallback *SliceOfPointerToGongStructFieldFormCallback) {
	sliceofpointertogongstructfieldFormCallback = new(SliceOfPointerToGongStructFieldFormCallback)
	sliceofpointertogongstructfieldFormCallback.probe = probe
	sliceofpointertogongstructfieldFormCallback.sliceofpointertogongstructfield = sliceofpointertogongstructfield
	sliceofpointertogongstructfieldFormCallback.formGroup = formGroup

	sliceofpointertogongstructfieldFormCallback.CreationMode = (sliceofpointertogongstructfield == nil)

	return
}

type SliceOfPointerToGongStructFieldFormCallback struct {
	sliceofpointertogongstructfield *models.SliceOfPointerToGongStructField

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (sliceofpointertogongstructfieldFormCallback *SliceOfPointerToGongStructFieldFormCallback) OnSave() {

	// log.Println("SliceOfPointerToGongStructFieldFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	sliceofpointertogongstructfieldFormCallback.probe.formStage.Checkout()

	if sliceofpointertogongstructfieldFormCallback.sliceofpointertogongstructfield == nil {
		sliceofpointertogongstructfieldFormCallback.sliceofpointertogongstructfield = new(models.SliceOfPointerToGongStructField).Stage(sliceofpointertogongstructfieldFormCallback.probe.stageOfInterest)
	}
	sliceofpointertogongstructfield_ := sliceofpointertogongstructfieldFormCallback.sliceofpointertogongstructfield
	_ = sliceofpointertogongstructfield_

	for _, formDiv := range sliceofpointertogongstructfieldFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(sliceofpointertogongstructfield_.Name), formDiv)
		case "GongStruct":
			FormDivSelectFieldToField(&(sliceofpointertogongstructfield_.GongStruct), sliceofpointertogongstructfieldFormCallback.probe.stageOfInterest, formDiv)
		case "Index":
			FormDivBasicFieldToField(&(sliceofpointertogongstructfield_.Index), formDiv)
		case "CompositeStructName":
			FormDivBasicFieldToField(&(sliceofpointertogongstructfield_.CompositeStructName), formDiv)
		case "GongStruct:SliceOfPointerToGongStructFields":
			// WARNING : this form deals with the N-N association "GongStruct.SliceOfPointerToGongStructFields []*SliceOfPointerToGongStructField" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of SliceOfPointerToGongStructField). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.GongStruct
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "GongStruct"
				rf.Fieldname = "SliceOfPointerToGongStructFields"
				formerAssociationSource := models.GetReverseFieldOwner(
					sliceofpointertogongstructfieldFormCallback.probe.stageOfInterest,
					sliceofpointertogongstructfield_,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.GongStruct)
					if !ok {
						log.Fatalln("Source of GongStruct.SliceOfPointerToGongStructFields []*SliceOfPointerToGongStructField, is not an GongStruct instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.SliceOfPointerToGongStructFields, sliceofpointertogongstructfield_)
					formerSource.SliceOfPointerToGongStructFields = slices.Delete(formerSource.SliceOfPointerToGongStructFields, idx, idx+1)
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
			var newSource *models.GongStruct
			for _gongstruct := range *models.GetGongstructInstancesSet[models.GongStruct](sliceofpointertogongstructfieldFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _gongstruct.GetName() == newSourceName.GetName() {
					newSource = _gongstruct // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of GongStruct.SliceOfPointerToGongStructFields []*SliceOfPointerToGongStructField, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.SliceOfPointerToGongStructFields = append(newSource.SliceOfPointerToGongStructFields, sliceofpointertogongstructfield_)
		}
	}

	// manage the suppress operation
	if sliceofpointertogongstructfieldFormCallback.formGroup.HasSuppressButtonBeenPressed {
		sliceofpointertogongstructfield_.Unstage(sliceofpointertogongstructfieldFormCallback.probe.stageOfInterest)
	}

	sliceofpointertogongstructfieldFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.SliceOfPointerToGongStructField](
		sliceofpointertogongstructfieldFormCallback.probe,
	)
	sliceofpointertogongstructfieldFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if sliceofpointertogongstructfieldFormCallback.CreationMode || sliceofpointertogongstructfieldFormCallback.formGroup.HasSuppressButtonBeenPressed {
		sliceofpointertogongstructfieldFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(sliceofpointertogongstructfieldFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__SliceOfPointerToGongStructFieldFormCallback(
			nil,
			sliceofpointertogongstructfieldFormCallback.probe,
			newFormGroup,
		)
		sliceofpointertogongstructfield := new(models.SliceOfPointerToGongStructField)
		FillUpForm(sliceofpointertogongstructfield, newFormGroup, sliceofpointertogongstructfieldFormCallback.probe)
		sliceofpointertogongstructfieldFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(sliceofpointertogongstructfieldFormCallback.probe)
}
