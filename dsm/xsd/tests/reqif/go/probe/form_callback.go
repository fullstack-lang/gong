// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	table "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gong/dsm/xsd/tests/reqif/go/models"
)

// to avoid errors when time and slices packages are not used in the generated code
const _ = time.Nanosecond

var _ = slices.Delete([]string{"a"}, 0, 1)

var _ = log.Panicf

// insertion point
func __gong__New__ALTERNATIVE_IDFormCallback(
	alternative_id *models.ALTERNATIVE_ID,
	probe *Probe,
	formGroup *table.FormGroup,
) (alternative_idFormCallback *ALTERNATIVE_IDFormCallback) {
	alternative_idFormCallback = new(ALTERNATIVE_IDFormCallback)
	alternative_idFormCallback.probe = probe
	alternative_idFormCallback.alternative_id = alternative_id
	alternative_idFormCallback.formGroup = formGroup

	alternative_idFormCallback.CreationMode = (alternative_id == nil)

	return
}

type ALTERNATIVE_IDFormCallback struct {
	alternative_id *models.ALTERNATIVE_ID

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (alternative_idFormCallback *ALTERNATIVE_IDFormCallback) OnSave() {
	alternative_idFormCallback.probe.stageOfInterest.Lock()
	defer alternative_idFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("ALTERNATIVE_IDFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	alternative_idFormCallback.probe.formStage.Checkout()

	if alternative_idFormCallback.alternative_id == nil {
		alternative_idFormCallback.alternative_id = new(models.ALTERNATIVE_ID).Stage(alternative_idFormCallback.probe.stageOfInterest)
	}
	alternative_id_ := alternative_idFormCallback.alternative_id
	_ = alternative_id_

	for _, formDiv := range alternative_idFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(alternative_id_.Name), formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&(alternative_id_.IDENTIFIER), formDiv)
		}
	}

	// manage the suppress operation
	if alternative_idFormCallback.formGroup.HasSuppressButtonBeenPressed {
		alternative_id_.Unstage(alternative_idFormCallback.probe.stageOfInterest)
	}

	alternative_idFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.ALTERNATIVE_ID](
		alternative_idFormCallback.probe,
	)

	// display a new form by reset the form stage
	if alternative_idFormCallback.CreationMode || alternative_idFormCallback.formGroup.HasSuppressButtonBeenPressed {
		alternative_idFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(alternative_idFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ALTERNATIVE_IDFormCallback(
			nil,
			alternative_idFormCallback.probe,
			newFormGroup,
		)
		alternative_id := new(models.ALTERNATIVE_ID)
		FillUpForm(alternative_id, newFormGroup, alternative_idFormCallback.probe)
		alternative_idFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(alternative_idFormCallback.probe)
}

func __gong__New__ATTRIBUTE_DEFINITION_BOOLEANFormCallback(
	attribute_definition_boolean *models.ATTRIBUTE_DEFINITION_BOOLEAN,
	probe *Probe,
	formGroup *table.FormGroup,
) (attribute_definition_booleanFormCallback *ATTRIBUTE_DEFINITION_BOOLEANFormCallback) {
	attribute_definition_booleanFormCallback = new(ATTRIBUTE_DEFINITION_BOOLEANFormCallback)
	attribute_definition_booleanFormCallback.probe = probe
	attribute_definition_booleanFormCallback.attribute_definition_boolean = attribute_definition_boolean
	attribute_definition_booleanFormCallback.formGroup = formGroup

	attribute_definition_booleanFormCallback.CreationMode = (attribute_definition_boolean == nil)

	return
}

type ATTRIBUTE_DEFINITION_BOOLEANFormCallback struct {
	attribute_definition_boolean *models.ATTRIBUTE_DEFINITION_BOOLEAN

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (attribute_definition_booleanFormCallback *ATTRIBUTE_DEFINITION_BOOLEANFormCallback) OnSave() {
	attribute_definition_booleanFormCallback.probe.stageOfInterest.Lock()
	defer attribute_definition_booleanFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("ATTRIBUTE_DEFINITION_BOOLEANFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	attribute_definition_booleanFormCallback.probe.formStage.Checkout()

	if attribute_definition_booleanFormCallback.attribute_definition_boolean == nil {
		attribute_definition_booleanFormCallback.attribute_definition_boolean = new(models.ATTRIBUTE_DEFINITION_BOOLEAN).Stage(attribute_definition_booleanFormCallback.probe.stageOfInterest)
	}
	attribute_definition_boolean_ := attribute_definition_booleanFormCallback.attribute_definition_boolean
	_ = attribute_definition_boolean_

	for _, formDiv := range attribute_definition_booleanFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(attribute_definition_boolean_.Name), formDiv)
		case "DESC":
			FormDivBasicFieldToField(&(attribute_definition_boolean_.DESC), formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&(attribute_definition_boolean_.IDENTIFIER), formDiv)
		case "IS_EDITABLE":
			FormDivBasicFieldToField(&(attribute_definition_boolean_.IS_EDITABLE), formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&(attribute_definition_boolean_.LAST_CHANGE), formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&(attribute_definition_boolean_.LONG_NAME), formDiv)
		case "ALTERNATIVE_ID":
			FormDivSelectFieldToField(&(attribute_definition_boolean_.ALTERNATIVE_ID), attribute_definition_booleanFormCallback.probe.stageOfInterest, formDiv)
		case "DEFAULT_VALUE":
			FormDivSelectFieldToField(&(attribute_definition_boolean_.DEFAULT_VALUE), attribute_definition_booleanFormCallback.probe.stageOfInterest, formDiv)
		case "TYPE":
			FormDivSelectFieldToField(&(attribute_definition_boolean_.TYPE), attribute_definition_booleanFormCallback.probe.stageOfInterest, formDiv)
		case "A_SPEC_ATTRIBUTES:ATTRIBUTE_DEFINITION_BOOLEAN":
			// WARNING : this form deals with the N-N association "A_SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN []*ATTRIBUTE_DEFINITION_BOOLEAN" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of ATTRIBUTE_DEFINITION_BOOLEAN). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.A_SPEC_ATTRIBUTES
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "A_SPEC_ATTRIBUTES"
				rf.Fieldname = "ATTRIBUTE_DEFINITION_BOOLEAN"
				formerAssociationSource := attribute_definition_boolean_.GongGetReverseFieldOwner(
					attribute_definition_booleanFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.A_SPEC_ATTRIBUTES)
					if !ok {
						log.Fatalln("Source of A_SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN []*ATTRIBUTE_DEFINITION_BOOLEAN, is not an A_SPEC_ATTRIBUTES instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.ATTRIBUTE_DEFINITION_BOOLEAN, attribute_definition_boolean_)
					formerSource.ATTRIBUTE_DEFINITION_BOOLEAN = slices.Delete(formerSource.ATTRIBUTE_DEFINITION_BOOLEAN, idx, idx+1)
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
			var newSource *models.A_SPEC_ATTRIBUTES
			for _a_spec_attributes := range *models.GetGongstructInstancesSet[models.A_SPEC_ATTRIBUTES](attribute_definition_booleanFormCallback.probe.stageOfInterest) {
				// the match is base on the name
				if _a_spec_attributes.GetName() == newSourceName.GetName() {
					newSource = _a_spec_attributes // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of A_SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN []*ATTRIBUTE_DEFINITION_BOOLEAN, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.ATTRIBUTE_DEFINITION_BOOLEAN = append(newSource.ATTRIBUTE_DEFINITION_BOOLEAN, attribute_definition_boolean_)
		}
	}

	// manage the suppress operation
	if attribute_definition_booleanFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attribute_definition_boolean_.Unstage(attribute_definition_booleanFormCallback.probe.stageOfInterest)
	}

	attribute_definition_booleanFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.ATTRIBUTE_DEFINITION_BOOLEAN](
		attribute_definition_booleanFormCallback.probe,
	)

	// display a new form by reset the form stage
	if attribute_definition_booleanFormCallback.CreationMode || attribute_definition_booleanFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attribute_definition_booleanFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(attribute_definition_booleanFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ATTRIBUTE_DEFINITION_BOOLEANFormCallback(
			nil,
			attribute_definition_booleanFormCallback.probe,
			newFormGroup,
		)
		attribute_definition_boolean := new(models.ATTRIBUTE_DEFINITION_BOOLEAN)
		FillUpForm(attribute_definition_boolean, newFormGroup, attribute_definition_booleanFormCallback.probe)
		attribute_definition_booleanFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(attribute_definition_booleanFormCallback.probe)
}

func __gong__New__ATTRIBUTE_DEFINITION_DATEFormCallback(
	attribute_definition_date *models.ATTRIBUTE_DEFINITION_DATE,
	probe *Probe,
	formGroup *table.FormGroup,
) (attribute_definition_dateFormCallback *ATTRIBUTE_DEFINITION_DATEFormCallback) {
	attribute_definition_dateFormCallback = new(ATTRIBUTE_DEFINITION_DATEFormCallback)
	attribute_definition_dateFormCallback.probe = probe
	attribute_definition_dateFormCallback.attribute_definition_date = attribute_definition_date
	attribute_definition_dateFormCallback.formGroup = formGroup

	attribute_definition_dateFormCallback.CreationMode = (attribute_definition_date == nil)

	return
}

type ATTRIBUTE_DEFINITION_DATEFormCallback struct {
	attribute_definition_date *models.ATTRIBUTE_DEFINITION_DATE

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (attribute_definition_dateFormCallback *ATTRIBUTE_DEFINITION_DATEFormCallback) OnSave() {
	attribute_definition_dateFormCallback.probe.stageOfInterest.Lock()
	defer attribute_definition_dateFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("ATTRIBUTE_DEFINITION_DATEFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	attribute_definition_dateFormCallback.probe.formStage.Checkout()

	if attribute_definition_dateFormCallback.attribute_definition_date == nil {
		attribute_definition_dateFormCallback.attribute_definition_date = new(models.ATTRIBUTE_DEFINITION_DATE).Stage(attribute_definition_dateFormCallback.probe.stageOfInterest)
	}
	attribute_definition_date_ := attribute_definition_dateFormCallback.attribute_definition_date
	_ = attribute_definition_date_

	for _, formDiv := range attribute_definition_dateFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(attribute_definition_date_.Name), formDiv)
		case "DESC":
			FormDivBasicFieldToField(&(attribute_definition_date_.DESC), formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&(attribute_definition_date_.IDENTIFIER), formDiv)
		case "IS_EDITABLE":
			FormDivBasicFieldToField(&(attribute_definition_date_.IS_EDITABLE), formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&(attribute_definition_date_.LAST_CHANGE), formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&(attribute_definition_date_.LONG_NAME), formDiv)
		case "ALTERNATIVE_ID":
			FormDivSelectFieldToField(&(attribute_definition_date_.ALTERNATIVE_ID), attribute_definition_dateFormCallback.probe.stageOfInterest, formDiv)
		case "DEFAULT_VALUE":
			FormDivSelectFieldToField(&(attribute_definition_date_.DEFAULT_VALUE), attribute_definition_dateFormCallback.probe.stageOfInterest, formDiv)
		case "TYPE":
			FormDivSelectFieldToField(&(attribute_definition_date_.TYPE), attribute_definition_dateFormCallback.probe.stageOfInterest, formDiv)
		case "A_SPEC_ATTRIBUTES:ATTRIBUTE_DEFINITION_DATE":
			// WARNING : this form deals with the N-N association "A_SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE []*ATTRIBUTE_DEFINITION_DATE" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of ATTRIBUTE_DEFINITION_DATE). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.A_SPEC_ATTRIBUTES
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "A_SPEC_ATTRIBUTES"
				rf.Fieldname = "ATTRIBUTE_DEFINITION_DATE"
				formerAssociationSource := attribute_definition_date_.GongGetReverseFieldOwner(
					attribute_definition_dateFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.A_SPEC_ATTRIBUTES)
					if !ok {
						log.Fatalln("Source of A_SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE []*ATTRIBUTE_DEFINITION_DATE, is not an A_SPEC_ATTRIBUTES instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.ATTRIBUTE_DEFINITION_DATE, attribute_definition_date_)
					formerSource.ATTRIBUTE_DEFINITION_DATE = slices.Delete(formerSource.ATTRIBUTE_DEFINITION_DATE, idx, idx+1)
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
			var newSource *models.A_SPEC_ATTRIBUTES
			for _a_spec_attributes := range *models.GetGongstructInstancesSet[models.A_SPEC_ATTRIBUTES](attribute_definition_dateFormCallback.probe.stageOfInterest) {
				// the match is base on the name
				if _a_spec_attributes.GetName() == newSourceName.GetName() {
					newSource = _a_spec_attributes // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of A_SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE []*ATTRIBUTE_DEFINITION_DATE, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.ATTRIBUTE_DEFINITION_DATE = append(newSource.ATTRIBUTE_DEFINITION_DATE, attribute_definition_date_)
		}
	}

	// manage the suppress operation
	if attribute_definition_dateFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attribute_definition_date_.Unstage(attribute_definition_dateFormCallback.probe.stageOfInterest)
	}

	attribute_definition_dateFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.ATTRIBUTE_DEFINITION_DATE](
		attribute_definition_dateFormCallback.probe,
	)

	// display a new form by reset the form stage
	if attribute_definition_dateFormCallback.CreationMode || attribute_definition_dateFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attribute_definition_dateFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(attribute_definition_dateFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ATTRIBUTE_DEFINITION_DATEFormCallback(
			nil,
			attribute_definition_dateFormCallback.probe,
			newFormGroup,
		)
		attribute_definition_date := new(models.ATTRIBUTE_DEFINITION_DATE)
		FillUpForm(attribute_definition_date, newFormGroup, attribute_definition_dateFormCallback.probe)
		attribute_definition_dateFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(attribute_definition_dateFormCallback.probe)
}

func __gong__New__ATTRIBUTE_DEFINITION_ENUMERATIONFormCallback(
	attribute_definition_enumeration *models.ATTRIBUTE_DEFINITION_ENUMERATION,
	probe *Probe,
	formGroup *table.FormGroup,
) (attribute_definition_enumerationFormCallback *ATTRIBUTE_DEFINITION_ENUMERATIONFormCallback) {
	attribute_definition_enumerationFormCallback = new(ATTRIBUTE_DEFINITION_ENUMERATIONFormCallback)
	attribute_definition_enumerationFormCallback.probe = probe
	attribute_definition_enumerationFormCallback.attribute_definition_enumeration = attribute_definition_enumeration
	attribute_definition_enumerationFormCallback.formGroup = formGroup

	attribute_definition_enumerationFormCallback.CreationMode = (attribute_definition_enumeration == nil)

	return
}

type ATTRIBUTE_DEFINITION_ENUMERATIONFormCallback struct {
	attribute_definition_enumeration *models.ATTRIBUTE_DEFINITION_ENUMERATION

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (attribute_definition_enumerationFormCallback *ATTRIBUTE_DEFINITION_ENUMERATIONFormCallback) OnSave() {
	attribute_definition_enumerationFormCallback.probe.stageOfInterest.Lock()
	defer attribute_definition_enumerationFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("ATTRIBUTE_DEFINITION_ENUMERATIONFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	attribute_definition_enumerationFormCallback.probe.formStage.Checkout()

	if attribute_definition_enumerationFormCallback.attribute_definition_enumeration == nil {
		attribute_definition_enumerationFormCallback.attribute_definition_enumeration = new(models.ATTRIBUTE_DEFINITION_ENUMERATION).Stage(attribute_definition_enumerationFormCallback.probe.stageOfInterest)
	}
	attribute_definition_enumeration_ := attribute_definition_enumerationFormCallback.attribute_definition_enumeration
	_ = attribute_definition_enumeration_

	for _, formDiv := range attribute_definition_enumerationFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(attribute_definition_enumeration_.Name), formDiv)
		case "DESC":
			FormDivBasicFieldToField(&(attribute_definition_enumeration_.DESC), formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&(attribute_definition_enumeration_.IDENTIFIER), formDiv)
		case "IS_EDITABLE":
			FormDivBasicFieldToField(&(attribute_definition_enumeration_.IS_EDITABLE), formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&(attribute_definition_enumeration_.LAST_CHANGE), formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&(attribute_definition_enumeration_.LONG_NAME), formDiv)
		case "MULTI_VALUED":
			FormDivBasicFieldToField(&(attribute_definition_enumeration_.MULTI_VALUED), formDiv)
		case "ALTERNATIVE_ID":
			FormDivSelectFieldToField(&(attribute_definition_enumeration_.ALTERNATIVE_ID), attribute_definition_enumerationFormCallback.probe.stageOfInterest, formDiv)
		case "DEFAULT_VALUE":
			FormDivSelectFieldToField(&(attribute_definition_enumeration_.DEFAULT_VALUE), attribute_definition_enumerationFormCallback.probe.stageOfInterest, formDiv)
		case "TYPE":
			FormDivSelectFieldToField(&(attribute_definition_enumeration_.TYPE), attribute_definition_enumerationFormCallback.probe.stageOfInterest, formDiv)
		case "A_SPEC_ATTRIBUTES:ATTRIBUTE_DEFINITION_ENUMERATION":
			// WARNING : this form deals with the N-N association "A_SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION []*ATTRIBUTE_DEFINITION_ENUMERATION" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of ATTRIBUTE_DEFINITION_ENUMERATION). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.A_SPEC_ATTRIBUTES
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "A_SPEC_ATTRIBUTES"
				rf.Fieldname = "ATTRIBUTE_DEFINITION_ENUMERATION"
				formerAssociationSource := attribute_definition_enumeration_.GongGetReverseFieldOwner(
					attribute_definition_enumerationFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.A_SPEC_ATTRIBUTES)
					if !ok {
						log.Fatalln("Source of A_SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION []*ATTRIBUTE_DEFINITION_ENUMERATION, is not an A_SPEC_ATTRIBUTES instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.ATTRIBUTE_DEFINITION_ENUMERATION, attribute_definition_enumeration_)
					formerSource.ATTRIBUTE_DEFINITION_ENUMERATION = slices.Delete(formerSource.ATTRIBUTE_DEFINITION_ENUMERATION, idx, idx+1)
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
			var newSource *models.A_SPEC_ATTRIBUTES
			for _a_spec_attributes := range *models.GetGongstructInstancesSet[models.A_SPEC_ATTRIBUTES](attribute_definition_enumerationFormCallback.probe.stageOfInterest) {
				// the match is base on the name
				if _a_spec_attributes.GetName() == newSourceName.GetName() {
					newSource = _a_spec_attributes // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of A_SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION []*ATTRIBUTE_DEFINITION_ENUMERATION, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.ATTRIBUTE_DEFINITION_ENUMERATION = append(newSource.ATTRIBUTE_DEFINITION_ENUMERATION, attribute_definition_enumeration_)
		}
	}

	// manage the suppress operation
	if attribute_definition_enumerationFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attribute_definition_enumeration_.Unstage(attribute_definition_enumerationFormCallback.probe.stageOfInterest)
	}

	attribute_definition_enumerationFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.ATTRIBUTE_DEFINITION_ENUMERATION](
		attribute_definition_enumerationFormCallback.probe,
	)

	// display a new form by reset the form stage
	if attribute_definition_enumerationFormCallback.CreationMode || attribute_definition_enumerationFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attribute_definition_enumerationFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(attribute_definition_enumerationFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ATTRIBUTE_DEFINITION_ENUMERATIONFormCallback(
			nil,
			attribute_definition_enumerationFormCallback.probe,
			newFormGroup,
		)
		attribute_definition_enumeration := new(models.ATTRIBUTE_DEFINITION_ENUMERATION)
		FillUpForm(attribute_definition_enumeration, newFormGroup, attribute_definition_enumerationFormCallback.probe)
		attribute_definition_enumerationFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(attribute_definition_enumerationFormCallback.probe)
}

func __gong__New__ATTRIBUTE_DEFINITION_INTEGERFormCallback(
	attribute_definition_integer *models.ATTRIBUTE_DEFINITION_INTEGER,
	probe *Probe,
	formGroup *table.FormGroup,
) (attribute_definition_integerFormCallback *ATTRIBUTE_DEFINITION_INTEGERFormCallback) {
	attribute_definition_integerFormCallback = new(ATTRIBUTE_DEFINITION_INTEGERFormCallback)
	attribute_definition_integerFormCallback.probe = probe
	attribute_definition_integerFormCallback.attribute_definition_integer = attribute_definition_integer
	attribute_definition_integerFormCallback.formGroup = formGroup

	attribute_definition_integerFormCallback.CreationMode = (attribute_definition_integer == nil)

	return
}

type ATTRIBUTE_DEFINITION_INTEGERFormCallback struct {
	attribute_definition_integer *models.ATTRIBUTE_DEFINITION_INTEGER

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (attribute_definition_integerFormCallback *ATTRIBUTE_DEFINITION_INTEGERFormCallback) OnSave() {
	attribute_definition_integerFormCallback.probe.stageOfInterest.Lock()
	defer attribute_definition_integerFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("ATTRIBUTE_DEFINITION_INTEGERFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	attribute_definition_integerFormCallback.probe.formStage.Checkout()

	if attribute_definition_integerFormCallback.attribute_definition_integer == nil {
		attribute_definition_integerFormCallback.attribute_definition_integer = new(models.ATTRIBUTE_DEFINITION_INTEGER).Stage(attribute_definition_integerFormCallback.probe.stageOfInterest)
	}
	attribute_definition_integer_ := attribute_definition_integerFormCallback.attribute_definition_integer
	_ = attribute_definition_integer_

	for _, formDiv := range attribute_definition_integerFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(attribute_definition_integer_.Name), formDiv)
		case "DESC":
			FormDivBasicFieldToField(&(attribute_definition_integer_.DESC), formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&(attribute_definition_integer_.IDENTIFIER), formDiv)
		case "IS_EDITABLE":
			FormDivBasicFieldToField(&(attribute_definition_integer_.IS_EDITABLE), formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&(attribute_definition_integer_.LAST_CHANGE), formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&(attribute_definition_integer_.LONG_NAME), formDiv)
		case "ALTERNATIVE_ID":
			FormDivSelectFieldToField(&(attribute_definition_integer_.ALTERNATIVE_ID), attribute_definition_integerFormCallback.probe.stageOfInterest, formDiv)
		case "DEFAULT_VALUE":
			FormDivSelectFieldToField(&(attribute_definition_integer_.DEFAULT_VALUE), attribute_definition_integerFormCallback.probe.stageOfInterest, formDiv)
		case "TYPE":
			FormDivSelectFieldToField(&(attribute_definition_integer_.TYPE), attribute_definition_integerFormCallback.probe.stageOfInterest, formDiv)
		case "A_SPEC_ATTRIBUTES:ATTRIBUTE_DEFINITION_INTEGER":
			// WARNING : this form deals with the N-N association "A_SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER []*ATTRIBUTE_DEFINITION_INTEGER" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of ATTRIBUTE_DEFINITION_INTEGER). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.A_SPEC_ATTRIBUTES
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "A_SPEC_ATTRIBUTES"
				rf.Fieldname = "ATTRIBUTE_DEFINITION_INTEGER"
				formerAssociationSource := attribute_definition_integer_.GongGetReverseFieldOwner(
					attribute_definition_integerFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.A_SPEC_ATTRIBUTES)
					if !ok {
						log.Fatalln("Source of A_SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER []*ATTRIBUTE_DEFINITION_INTEGER, is not an A_SPEC_ATTRIBUTES instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.ATTRIBUTE_DEFINITION_INTEGER, attribute_definition_integer_)
					formerSource.ATTRIBUTE_DEFINITION_INTEGER = slices.Delete(formerSource.ATTRIBUTE_DEFINITION_INTEGER, idx, idx+1)
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
			var newSource *models.A_SPEC_ATTRIBUTES
			for _a_spec_attributes := range *models.GetGongstructInstancesSet[models.A_SPEC_ATTRIBUTES](attribute_definition_integerFormCallback.probe.stageOfInterest) {
				// the match is base on the name
				if _a_spec_attributes.GetName() == newSourceName.GetName() {
					newSource = _a_spec_attributes // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of A_SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER []*ATTRIBUTE_DEFINITION_INTEGER, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.ATTRIBUTE_DEFINITION_INTEGER = append(newSource.ATTRIBUTE_DEFINITION_INTEGER, attribute_definition_integer_)
		}
	}

	// manage the suppress operation
	if attribute_definition_integerFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attribute_definition_integer_.Unstage(attribute_definition_integerFormCallback.probe.stageOfInterest)
	}

	attribute_definition_integerFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.ATTRIBUTE_DEFINITION_INTEGER](
		attribute_definition_integerFormCallback.probe,
	)

	// display a new form by reset the form stage
	if attribute_definition_integerFormCallback.CreationMode || attribute_definition_integerFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attribute_definition_integerFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(attribute_definition_integerFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ATTRIBUTE_DEFINITION_INTEGERFormCallback(
			nil,
			attribute_definition_integerFormCallback.probe,
			newFormGroup,
		)
		attribute_definition_integer := new(models.ATTRIBUTE_DEFINITION_INTEGER)
		FillUpForm(attribute_definition_integer, newFormGroup, attribute_definition_integerFormCallback.probe)
		attribute_definition_integerFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(attribute_definition_integerFormCallback.probe)
}

func __gong__New__ATTRIBUTE_DEFINITION_REALFormCallback(
	attribute_definition_real *models.ATTRIBUTE_DEFINITION_REAL,
	probe *Probe,
	formGroup *table.FormGroup,
) (attribute_definition_realFormCallback *ATTRIBUTE_DEFINITION_REALFormCallback) {
	attribute_definition_realFormCallback = new(ATTRIBUTE_DEFINITION_REALFormCallback)
	attribute_definition_realFormCallback.probe = probe
	attribute_definition_realFormCallback.attribute_definition_real = attribute_definition_real
	attribute_definition_realFormCallback.formGroup = formGroup

	attribute_definition_realFormCallback.CreationMode = (attribute_definition_real == nil)

	return
}

type ATTRIBUTE_DEFINITION_REALFormCallback struct {
	attribute_definition_real *models.ATTRIBUTE_DEFINITION_REAL

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (attribute_definition_realFormCallback *ATTRIBUTE_DEFINITION_REALFormCallback) OnSave() {
	attribute_definition_realFormCallback.probe.stageOfInterest.Lock()
	defer attribute_definition_realFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("ATTRIBUTE_DEFINITION_REALFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	attribute_definition_realFormCallback.probe.formStage.Checkout()

	if attribute_definition_realFormCallback.attribute_definition_real == nil {
		attribute_definition_realFormCallback.attribute_definition_real = new(models.ATTRIBUTE_DEFINITION_REAL).Stage(attribute_definition_realFormCallback.probe.stageOfInterest)
	}
	attribute_definition_real_ := attribute_definition_realFormCallback.attribute_definition_real
	_ = attribute_definition_real_

	for _, formDiv := range attribute_definition_realFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(attribute_definition_real_.Name), formDiv)
		case "DESC":
			FormDivBasicFieldToField(&(attribute_definition_real_.DESC), formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&(attribute_definition_real_.IDENTIFIER), formDiv)
		case "IS_EDITABLE":
			FormDivBasicFieldToField(&(attribute_definition_real_.IS_EDITABLE), formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&(attribute_definition_real_.LAST_CHANGE), formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&(attribute_definition_real_.LONG_NAME), formDiv)
		case "ALTERNATIVE_ID":
			FormDivSelectFieldToField(&(attribute_definition_real_.ALTERNATIVE_ID), attribute_definition_realFormCallback.probe.stageOfInterest, formDiv)
		case "DEFAULT_VALUE":
			FormDivSelectFieldToField(&(attribute_definition_real_.DEFAULT_VALUE), attribute_definition_realFormCallback.probe.stageOfInterest, formDiv)
		case "TYPE":
			FormDivSelectFieldToField(&(attribute_definition_real_.TYPE), attribute_definition_realFormCallback.probe.stageOfInterest, formDiv)
		case "A_SPEC_ATTRIBUTES:ATTRIBUTE_DEFINITION_REAL":
			// WARNING : this form deals with the N-N association "A_SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL []*ATTRIBUTE_DEFINITION_REAL" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of ATTRIBUTE_DEFINITION_REAL). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.A_SPEC_ATTRIBUTES
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "A_SPEC_ATTRIBUTES"
				rf.Fieldname = "ATTRIBUTE_DEFINITION_REAL"
				formerAssociationSource := attribute_definition_real_.GongGetReverseFieldOwner(
					attribute_definition_realFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.A_SPEC_ATTRIBUTES)
					if !ok {
						log.Fatalln("Source of A_SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL []*ATTRIBUTE_DEFINITION_REAL, is not an A_SPEC_ATTRIBUTES instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.ATTRIBUTE_DEFINITION_REAL, attribute_definition_real_)
					formerSource.ATTRIBUTE_DEFINITION_REAL = slices.Delete(formerSource.ATTRIBUTE_DEFINITION_REAL, idx, idx+1)
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
			var newSource *models.A_SPEC_ATTRIBUTES
			for _a_spec_attributes := range *models.GetGongstructInstancesSet[models.A_SPEC_ATTRIBUTES](attribute_definition_realFormCallback.probe.stageOfInterest) {
				// the match is base on the name
				if _a_spec_attributes.GetName() == newSourceName.GetName() {
					newSource = _a_spec_attributes // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of A_SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL []*ATTRIBUTE_DEFINITION_REAL, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.ATTRIBUTE_DEFINITION_REAL = append(newSource.ATTRIBUTE_DEFINITION_REAL, attribute_definition_real_)
		}
	}

	// manage the suppress operation
	if attribute_definition_realFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attribute_definition_real_.Unstage(attribute_definition_realFormCallback.probe.stageOfInterest)
	}

	attribute_definition_realFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.ATTRIBUTE_DEFINITION_REAL](
		attribute_definition_realFormCallback.probe,
	)

	// display a new form by reset the form stage
	if attribute_definition_realFormCallback.CreationMode || attribute_definition_realFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attribute_definition_realFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(attribute_definition_realFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ATTRIBUTE_DEFINITION_REALFormCallback(
			nil,
			attribute_definition_realFormCallback.probe,
			newFormGroup,
		)
		attribute_definition_real := new(models.ATTRIBUTE_DEFINITION_REAL)
		FillUpForm(attribute_definition_real, newFormGroup, attribute_definition_realFormCallback.probe)
		attribute_definition_realFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(attribute_definition_realFormCallback.probe)
}

func __gong__New__ATTRIBUTE_DEFINITION_STRINGFormCallback(
	attribute_definition_string *models.ATTRIBUTE_DEFINITION_STRING,
	probe *Probe,
	formGroup *table.FormGroup,
) (attribute_definition_stringFormCallback *ATTRIBUTE_DEFINITION_STRINGFormCallback) {
	attribute_definition_stringFormCallback = new(ATTRIBUTE_DEFINITION_STRINGFormCallback)
	attribute_definition_stringFormCallback.probe = probe
	attribute_definition_stringFormCallback.attribute_definition_string = attribute_definition_string
	attribute_definition_stringFormCallback.formGroup = formGroup

	attribute_definition_stringFormCallback.CreationMode = (attribute_definition_string == nil)

	return
}

type ATTRIBUTE_DEFINITION_STRINGFormCallback struct {
	attribute_definition_string *models.ATTRIBUTE_DEFINITION_STRING

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (attribute_definition_stringFormCallback *ATTRIBUTE_DEFINITION_STRINGFormCallback) OnSave() {
	attribute_definition_stringFormCallback.probe.stageOfInterest.Lock()
	defer attribute_definition_stringFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("ATTRIBUTE_DEFINITION_STRINGFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	attribute_definition_stringFormCallback.probe.formStage.Checkout()

	if attribute_definition_stringFormCallback.attribute_definition_string == nil {
		attribute_definition_stringFormCallback.attribute_definition_string = new(models.ATTRIBUTE_DEFINITION_STRING).Stage(attribute_definition_stringFormCallback.probe.stageOfInterest)
	}
	attribute_definition_string_ := attribute_definition_stringFormCallback.attribute_definition_string
	_ = attribute_definition_string_

	for _, formDiv := range attribute_definition_stringFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(attribute_definition_string_.Name), formDiv)
		case "DESC":
			FormDivBasicFieldToField(&(attribute_definition_string_.DESC), formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&(attribute_definition_string_.IDENTIFIER), formDiv)
		case "IS_EDITABLE":
			FormDivBasicFieldToField(&(attribute_definition_string_.IS_EDITABLE), formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&(attribute_definition_string_.LAST_CHANGE), formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&(attribute_definition_string_.LONG_NAME), formDiv)
		case "ALTERNATIVE_ID":
			FormDivSelectFieldToField(&(attribute_definition_string_.ALTERNATIVE_ID), attribute_definition_stringFormCallback.probe.stageOfInterest, formDiv)
		case "DEFAULT_VALUE":
			FormDivSelectFieldToField(&(attribute_definition_string_.DEFAULT_VALUE), attribute_definition_stringFormCallback.probe.stageOfInterest, formDiv)
		case "TYPE":
			FormDivSelectFieldToField(&(attribute_definition_string_.TYPE), attribute_definition_stringFormCallback.probe.stageOfInterest, formDiv)
		case "A_SPEC_ATTRIBUTES:ATTRIBUTE_DEFINITION_STRING":
			// WARNING : this form deals with the N-N association "A_SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING []*ATTRIBUTE_DEFINITION_STRING" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of ATTRIBUTE_DEFINITION_STRING). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.A_SPEC_ATTRIBUTES
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "A_SPEC_ATTRIBUTES"
				rf.Fieldname = "ATTRIBUTE_DEFINITION_STRING"
				formerAssociationSource := attribute_definition_string_.GongGetReverseFieldOwner(
					attribute_definition_stringFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.A_SPEC_ATTRIBUTES)
					if !ok {
						log.Fatalln("Source of A_SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING []*ATTRIBUTE_DEFINITION_STRING, is not an A_SPEC_ATTRIBUTES instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.ATTRIBUTE_DEFINITION_STRING, attribute_definition_string_)
					formerSource.ATTRIBUTE_DEFINITION_STRING = slices.Delete(formerSource.ATTRIBUTE_DEFINITION_STRING, idx, idx+1)
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
			var newSource *models.A_SPEC_ATTRIBUTES
			for _a_spec_attributes := range *models.GetGongstructInstancesSet[models.A_SPEC_ATTRIBUTES](attribute_definition_stringFormCallback.probe.stageOfInterest) {
				// the match is base on the name
				if _a_spec_attributes.GetName() == newSourceName.GetName() {
					newSource = _a_spec_attributes // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of A_SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING []*ATTRIBUTE_DEFINITION_STRING, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.ATTRIBUTE_DEFINITION_STRING = append(newSource.ATTRIBUTE_DEFINITION_STRING, attribute_definition_string_)
		}
	}

	// manage the suppress operation
	if attribute_definition_stringFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attribute_definition_string_.Unstage(attribute_definition_stringFormCallback.probe.stageOfInterest)
	}

	attribute_definition_stringFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.ATTRIBUTE_DEFINITION_STRING](
		attribute_definition_stringFormCallback.probe,
	)

	// display a new form by reset the form stage
	if attribute_definition_stringFormCallback.CreationMode || attribute_definition_stringFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attribute_definition_stringFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(attribute_definition_stringFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ATTRIBUTE_DEFINITION_STRINGFormCallback(
			nil,
			attribute_definition_stringFormCallback.probe,
			newFormGroup,
		)
		attribute_definition_string := new(models.ATTRIBUTE_DEFINITION_STRING)
		FillUpForm(attribute_definition_string, newFormGroup, attribute_definition_stringFormCallback.probe)
		attribute_definition_stringFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(attribute_definition_stringFormCallback.probe)
}

func __gong__New__ATTRIBUTE_DEFINITION_XHTMLFormCallback(
	attribute_definition_xhtml *models.ATTRIBUTE_DEFINITION_XHTML,
	probe *Probe,
	formGroup *table.FormGroup,
) (attribute_definition_xhtmlFormCallback *ATTRIBUTE_DEFINITION_XHTMLFormCallback) {
	attribute_definition_xhtmlFormCallback = new(ATTRIBUTE_DEFINITION_XHTMLFormCallback)
	attribute_definition_xhtmlFormCallback.probe = probe
	attribute_definition_xhtmlFormCallback.attribute_definition_xhtml = attribute_definition_xhtml
	attribute_definition_xhtmlFormCallback.formGroup = formGroup

	attribute_definition_xhtmlFormCallback.CreationMode = (attribute_definition_xhtml == nil)

	return
}

type ATTRIBUTE_DEFINITION_XHTMLFormCallback struct {
	attribute_definition_xhtml *models.ATTRIBUTE_DEFINITION_XHTML

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (attribute_definition_xhtmlFormCallback *ATTRIBUTE_DEFINITION_XHTMLFormCallback) OnSave() {
	attribute_definition_xhtmlFormCallback.probe.stageOfInterest.Lock()
	defer attribute_definition_xhtmlFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("ATTRIBUTE_DEFINITION_XHTMLFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	attribute_definition_xhtmlFormCallback.probe.formStage.Checkout()

	if attribute_definition_xhtmlFormCallback.attribute_definition_xhtml == nil {
		attribute_definition_xhtmlFormCallback.attribute_definition_xhtml = new(models.ATTRIBUTE_DEFINITION_XHTML).Stage(attribute_definition_xhtmlFormCallback.probe.stageOfInterest)
	}
	attribute_definition_xhtml_ := attribute_definition_xhtmlFormCallback.attribute_definition_xhtml
	_ = attribute_definition_xhtml_

	for _, formDiv := range attribute_definition_xhtmlFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(attribute_definition_xhtml_.Name), formDiv)
		case "DESC":
			FormDivBasicFieldToField(&(attribute_definition_xhtml_.DESC), formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&(attribute_definition_xhtml_.IDENTIFIER), formDiv)
		case "IS_EDITABLE":
			FormDivBasicFieldToField(&(attribute_definition_xhtml_.IS_EDITABLE), formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&(attribute_definition_xhtml_.LAST_CHANGE), formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&(attribute_definition_xhtml_.LONG_NAME), formDiv)
		case "ALTERNATIVE_ID":
			FormDivSelectFieldToField(&(attribute_definition_xhtml_.ALTERNATIVE_ID), attribute_definition_xhtmlFormCallback.probe.stageOfInterest, formDiv)
		case "DEFAULT_VALUE":
			FormDivSelectFieldToField(&(attribute_definition_xhtml_.DEFAULT_VALUE), attribute_definition_xhtmlFormCallback.probe.stageOfInterest, formDiv)
		case "TYPE":
			FormDivSelectFieldToField(&(attribute_definition_xhtml_.TYPE), attribute_definition_xhtmlFormCallback.probe.stageOfInterest, formDiv)
		case "A_SPEC_ATTRIBUTES:ATTRIBUTE_DEFINITION_XHTML":
			// WARNING : this form deals with the N-N association "A_SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML []*ATTRIBUTE_DEFINITION_XHTML" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of ATTRIBUTE_DEFINITION_XHTML). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.A_SPEC_ATTRIBUTES
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "A_SPEC_ATTRIBUTES"
				rf.Fieldname = "ATTRIBUTE_DEFINITION_XHTML"
				formerAssociationSource := attribute_definition_xhtml_.GongGetReverseFieldOwner(
					attribute_definition_xhtmlFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.A_SPEC_ATTRIBUTES)
					if !ok {
						log.Fatalln("Source of A_SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML []*ATTRIBUTE_DEFINITION_XHTML, is not an A_SPEC_ATTRIBUTES instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.ATTRIBUTE_DEFINITION_XHTML, attribute_definition_xhtml_)
					formerSource.ATTRIBUTE_DEFINITION_XHTML = slices.Delete(formerSource.ATTRIBUTE_DEFINITION_XHTML, idx, idx+1)
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
			var newSource *models.A_SPEC_ATTRIBUTES
			for _a_spec_attributes := range *models.GetGongstructInstancesSet[models.A_SPEC_ATTRIBUTES](attribute_definition_xhtmlFormCallback.probe.stageOfInterest) {
				// the match is base on the name
				if _a_spec_attributes.GetName() == newSourceName.GetName() {
					newSource = _a_spec_attributes // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of A_SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML []*ATTRIBUTE_DEFINITION_XHTML, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.ATTRIBUTE_DEFINITION_XHTML = append(newSource.ATTRIBUTE_DEFINITION_XHTML, attribute_definition_xhtml_)
		}
	}

	// manage the suppress operation
	if attribute_definition_xhtmlFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attribute_definition_xhtml_.Unstage(attribute_definition_xhtmlFormCallback.probe.stageOfInterest)
	}

	attribute_definition_xhtmlFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.ATTRIBUTE_DEFINITION_XHTML](
		attribute_definition_xhtmlFormCallback.probe,
	)

	// display a new form by reset the form stage
	if attribute_definition_xhtmlFormCallback.CreationMode || attribute_definition_xhtmlFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attribute_definition_xhtmlFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(attribute_definition_xhtmlFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ATTRIBUTE_DEFINITION_XHTMLFormCallback(
			nil,
			attribute_definition_xhtmlFormCallback.probe,
			newFormGroup,
		)
		attribute_definition_xhtml := new(models.ATTRIBUTE_DEFINITION_XHTML)
		FillUpForm(attribute_definition_xhtml, newFormGroup, attribute_definition_xhtmlFormCallback.probe)
		attribute_definition_xhtmlFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(attribute_definition_xhtmlFormCallback.probe)
}

func __gong__New__ATTRIBUTE_VALUE_BOOLEANFormCallback(
	attribute_value_boolean *models.ATTRIBUTE_VALUE_BOOLEAN,
	probe *Probe,
	formGroup *table.FormGroup,
) (attribute_value_booleanFormCallback *ATTRIBUTE_VALUE_BOOLEANFormCallback) {
	attribute_value_booleanFormCallback = new(ATTRIBUTE_VALUE_BOOLEANFormCallback)
	attribute_value_booleanFormCallback.probe = probe
	attribute_value_booleanFormCallback.attribute_value_boolean = attribute_value_boolean
	attribute_value_booleanFormCallback.formGroup = formGroup

	attribute_value_booleanFormCallback.CreationMode = (attribute_value_boolean == nil)

	return
}

type ATTRIBUTE_VALUE_BOOLEANFormCallback struct {
	attribute_value_boolean *models.ATTRIBUTE_VALUE_BOOLEAN

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (attribute_value_booleanFormCallback *ATTRIBUTE_VALUE_BOOLEANFormCallback) OnSave() {
	attribute_value_booleanFormCallback.probe.stageOfInterest.Lock()
	defer attribute_value_booleanFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("ATTRIBUTE_VALUE_BOOLEANFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	attribute_value_booleanFormCallback.probe.formStage.Checkout()

	if attribute_value_booleanFormCallback.attribute_value_boolean == nil {
		attribute_value_booleanFormCallback.attribute_value_boolean = new(models.ATTRIBUTE_VALUE_BOOLEAN).Stage(attribute_value_booleanFormCallback.probe.stageOfInterest)
	}
	attribute_value_boolean_ := attribute_value_booleanFormCallback.attribute_value_boolean
	_ = attribute_value_boolean_

	for _, formDiv := range attribute_value_booleanFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(attribute_value_boolean_.Name), formDiv)
		case "THE_VALUE":
			FormDivBasicFieldToField(&(attribute_value_boolean_.THE_VALUE), formDiv)
		case "DEFINITION":
			FormDivSelectFieldToField(&(attribute_value_boolean_.DEFINITION), attribute_value_booleanFormCallback.probe.stageOfInterest, formDiv)
		case "A_ATTRIBUTE_VALUE_BOOLEAN:ATTRIBUTE_VALUE_BOOLEAN":
			// WARNING : this form deals with the N-N association "A_ATTRIBUTE_VALUE_BOOLEAN.ATTRIBUTE_VALUE_BOOLEAN []*ATTRIBUTE_VALUE_BOOLEAN" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of ATTRIBUTE_VALUE_BOOLEAN). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.A_ATTRIBUTE_VALUE_BOOLEAN
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "A_ATTRIBUTE_VALUE_BOOLEAN"
				rf.Fieldname = "ATTRIBUTE_VALUE_BOOLEAN"
				formerAssociationSource := attribute_value_boolean_.GongGetReverseFieldOwner(
					attribute_value_booleanFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.A_ATTRIBUTE_VALUE_BOOLEAN)
					if !ok {
						log.Fatalln("Source of A_ATTRIBUTE_VALUE_BOOLEAN.ATTRIBUTE_VALUE_BOOLEAN []*ATTRIBUTE_VALUE_BOOLEAN, is not an A_ATTRIBUTE_VALUE_BOOLEAN instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.ATTRIBUTE_VALUE_BOOLEAN, attribute_value_boolean_)
					formerSource.ATTRIBUTE_VALUE_BOOLEAN = slices.Delete(formerSource.ATTRIBUTE_VALUE_BOOLEAN, idx, idx+1)
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
			var newSource *models.A_ATTRIBUTE_VALUE_BOOLEAN
			for _a_attribute_value_boolean := range *models.GetGongstructInstancesSet[models.A_ATTRIBUTE_VALUE_BOOLEAN](attribute_value_booleanFormCallback.probe.stageOfInterest) {
				// the match is base on the name
				if _a_attribute_value_boolean.GetName() == newSourceName.GetName() {
					newSource = _a_attribute_value_boolean // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of A_ATTRIBUTE_VALUE_BOOLEAN.ATTRIBUTE_VALUE_BOOLEAN []*ATTRIBUTE_VALUE_BOOLEAN, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.ATTRIBUTE_VALUE_BOOLEAN = append(newSource.ATTRIBUTE_VALUE_BOOLEAN, attribute_value_boolean_)
		case "A_ATTRIBUTE_VALUE_XHTML_1:ATTRIBUTE_VALUE_BOOLEAN":
			// WARNING : this form deals with the N-N association "A_ATTRIBUTE_VALUE_XHTML_1.ATTRIBUTE_VALUE_BOOLEAN []*ATTRIBUTE_VALUE_BOOLEAN" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of ATTRIBUTE_VALUE_BOOLEAN). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.A_ATTRIBUTE_VALUE_XHTML_1
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "A_ATTRIBUTE_VALUE_XHTML_1"
				rf.Fieldname = "ATTRIBUTE_VALUE_BOOLEAN"
				formerAssociationSource := attribute_value_boolean_.GongGetReverseFieldOwner(
					attribute_value_booleanFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.A_ATTRIBUTE_VALUE_XHTML_1)
					if !ok {
						log.Fatalln("Source of A_ATTRIBUTE_VALUE_XHTML_1.ATTRIBUTE_VALUE_BOOLEAN []*ATTRIBUTE_VALUE_BOOLEAN, is not an A_ATTRIBUTE_VALUE_XHTML_1 instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.ATTRIBUTE_VALUE_BOOLEAN, attribute_value_boolean_)
					formerSource.ATTRIBUTE_VALUE_BOOLEAN = slices.Delete(formerSource.ATTRIBUTE_VALUE_BOOLEAN, idx, idx+1)
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
			var newSource *models.A_ATTRIBUTE_VALUE_XHTML_1
			for _a_attribute_value_xhtml_1 := range *models.GetGongstructInstancesSet[models.A_ATTRIBUTE_VALUE_XHTML_1](attribute_value_booleanFormCallback.probe.stageOfInterest) {
				// the match is base on the name
				if _a_attribute_value_xhtml_1.GetName() == newSourceName.GetName() {
					newSource = _a_attribute_value_xhtml_1 // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of A_ATTRIBUTE_VALUE_XHTML_1.ATTRIBUTE_VALUE_BOOLEAN []*ATTRIBUTE_VALUE_BOOLEAN, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.ATTRIBUTE_VALUE_BOOLEAN = append(newSource.ATTRIBUTE_VALUE_BOOLEAN, attribute_value_boolean_)
		}
	}

	// manage the suppress operation
	if attribute_value_booleanFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attribute_value_boolean_.Unstage(attribute_value_booleanFormCallback.probe.stageOfInterest)
	}

	attribute_value_booleanFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.ATTRIBUTE_VALUE_BOOLEAN](
		attribute_value_booleanFormCallback.probe,
	)

	// display a new form by reset the form stage
	if attribute_value_booleanFormCallback.CreationMode || attribute_value_booleanFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attribute_value_booleanFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(attribute_value_booleanFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ATTRIBUTE_VALUE_BOOLEANFormCallback(
			nil,
			attribute_value_booleanFormCallback.probe,
			newFormGroup,
		)
		attribute_value_boolean := new(models.ATTRIBUTE_VALUE_BOOLEAN)
		FillUpForm(attribute_value_boolean, newFormGroup, attribute_value_booleanFormCallback.probe)
		attribute_value_booleanFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(attribute_value_booleanFormCallback.probe)
}

func __gong__New__ATTRIBUTE_VALUE_DATEFormCallback(
	attribute_value_date *models.ATTRIBUTE_VALUE_DATE,
	probe *Probe,
	formGroup *table.FormGroup,
) (attribute_value_dateFormCallback *ATTRIBUTE_VALUE_DATEFormCallback) {
	attribute_value_dateFormCallback = new(ATTRIBUTE_VALUE_DATEFormCallback)
	attribute_value_dateFormCallback.probe = probe
	attribute_value_dateFormCallback.attribute_value_date = attribute_value_date
	attribute_value_dateFormCallback.formGroup = formGroup

	attribute_value_dateFormCallback.CreationMode = (attribute_value_date == nil)

	return
}

type ATTRIBUTE_VALUE_DATEFormCallback struct {
	attribute_value_date *models.ATTRIBUTE_VALUE_DATE

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (attribute_value_dateFormCallback *ATTRIBUTE_VALUE_DATEFormCallback) OnSave() {
	attribute_value_dateFormCallback.probe.stageOfInterest.Lock()
	defer attribute_value_dateFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("ATTRIBUTE_VALUE_DATEFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	attribute_value_dateFormCallback.probe.formStage.Checkout()

	if attribute_value_dateFormCallback.attribute_value_date == nil {
		attribute_value_dateFormCallback.attribute_value_date = new(models.ATTRIBUTE_VALUE_DATE).Stage(attribute_value_dateFormCallback.probe.stageOfInterest)
	}
	attribute_value_date_ := attribute_value_dateFormCallback.attribute_value_date
	_ = attribute_value_date_

	for _, formDiv := range attribute_value_dateFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(attribute_value_date_.Name), formDiv)
		case "THE_VALUE":
			FormDivBasicFieldToField(&(attribute_value_date_.THE_VALUE), formDiv)
		case "DEFINITION":
			FormDivSelectFieldToField(&(attribute_value_date_.DEFINITION), attribute_value_dateFormCallback.probe.stageOfInterest, formDiv)
		case "A_ATTRIBUTE_VALUE_DATE:ATTRIBUTE_VALUE_DATE":
			// WARNING : this form deals with the N-N association "A_ATTRIBUTE_VALUE_DATE.ATTRIBUTE_VALUE_DATE []*ATTRIBUTE_VALUE_DATE" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of ATTRIBUTE_VALUE_DATE). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.A_ATTRIBUTE_VALUE_DATE
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "A_ATTRIBUTE_VALUE_DATE"
				rf.Fieldname = "ATTRIBUTE_VALUE_DATE"
				formerAssociationSource := attribute_value_date_.GongGetReverseFieldOwner(
					attribute_value_dateFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.A_ATTRIBUTE_VALUE_DATE)
					if !ok {
						log.Fatalln("Source of A_ATTRIBUTE_VALUE_DATE.ATTRIBUTE_VALUE_DATE []*ATTRIBUTE_VALUE_DATE, is not an A_ATTRIBUTE_VALUE_DATE instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.ATTRIBUTE_VALUE_DATE, attribute_value_date_)
					formerSource.ATTRIBUTE_VALUE_DATE = slices.Delete(formerSource.ATTRIBUTE_VALUE_DATE, idx, idx+1)
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
			var newSource *models.A_ATTRIBUTE_VALUE_DATE
			for _a_attribute_value_date := range *models.GetGongstructInstancesSet[models.A_ATTRIBUTE_VALUE_DATE](attribute_value_dateFormCallback.probe.stageOfInterest) {
				// the match is base on the name
				if _a_attribute_value_date.GetName() == newSourceName.GetName() {
					newSource = _a_attribute_value_date // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of A_ATTRIBUTE_VALUE_DATE.ATTRIBUTE_VALUE_DATE []*ATTRIBUTE_VALUE_DATE, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.ATTRIBUTE_VALUE_DATE = append(newSource.ATTRIBUTE_VALUE_DATE, attribute_value_date_)
		case "A_ATTRIBUTE_VALUE_XHTML_1:ATTRIBUTE_VALUE_DATE":
			// WARNING : this form deals with the N-N association "A_ATTRIBUTE_VALUE_XHTML_1.ATTRIBUTE_VALUE_DATE []*ATTRIBUTE_VALUE_DATE" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of ATTRIBUTE_VALUE_DATE). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.A_ATTRIBUTE_VALUE_XHTML_1
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "A_ATTRIBUTE_VALUE_XHTML_1"
				rf.Fieldname = "ATTRIBUTE_VALUE_DATE"
				formerAssociationSource := attribute_value_date_.GongGetReverseFieldOwner(
					attribute_value_dateFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.A_ATTRIBUTE_VALUE_XHTML_1)
					if !ok {
						log.Fatalln("Source of A_ATTRIBUTE_VALUE_XHTML_1.ATTRIBUTE_VALUE_DATE []*ATTRIBUTE_VALUE_DATE, is not an A_ATTRIBUTE_VALUE_XHTML_1 instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.ATTRIBUTE_VALUE_DATE, attribute_value_date_)
					formerSource.ATTRIBUTE_VALUE_DATE = slices.Delete(formerSource.ATTRIBUTE_VALUE_DATE, idx, idx+1)
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
			var newSource *models.A_ATTRIBUTE_VALUE_XHTML_1
			for _a_attribute_value_xhtml_1 := range *models.GetGongstructInstancesSet[models.A_ATTRIBUTE_VALUE_XHTML_1](attribute_value_dateFormCallback.probe.stageOfInterest) {
				// the match is base on the name
				if _a_attribute_value_xhtml_1.GetName() == newSourceName.GetName() {
					newSource = _a_attribute_value_xhtml_1 // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of A_ATTRIBUTE_VALUE_XHTML_1.ATTRIBUTE_VALUE_DATE []*ATTRIBUTE_VALUE_DATE, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.ATTRIBUTE_VALUE_DATE = append(newSource.ATTRIBUTE_VALUE_DATE, attribute_value_date_)
		}
	}

	// manage the suppress operation
	if attribute_value_dateFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attribute_value_date_.Unstage(attribute_value_dateFormCallback.probe.stageOfInterest)
	}

	attribute_value_dateFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.ATTRIBUTE_VALUE_DATE](
		attribute_value_dateFormCallback.probe,
	)

	// display a new form by reset the form stage
	if attribute_value_dateFormCallback.CreationMode || attribute_value_dateFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attribute_value_dateFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(attribute_value_dateFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ATTRIBUTE_VALUE_DATEFormCallback(
			nil,
			attribute_value_dateFormCallback.probe,
			newFormGroup,
		)
		attribute_value_date := new(models.ATTRIBUTE_VALUE_DATE)
		FillUpForm(attribute_value_date, newFormGroup, attribute_value_dateFormCallback.probe)
		attribute_value_dateFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(attribute_value_dateFormCallback.probe)
}

func __gong__New__ATTRIBUTE_VALUE_ENUMERATIONFormCallback(
	attribute_value_enumeration *models.ATTRIBUTE_VALUE_ENUMERATION,
	probe *Probe,
	formGroup *table.FormGroup,
) (attribute_value_enumerationFormCallback *ATTRIBUTE_VALUE_ENUMERATIONFormCallback) {
	attribute_value_enumerationFormCallback = new(ATTRIBUTE_VALUE_ENUMERATIONFormCallback)
	attribute_value_enumerationFormCallback.probe = probe
	attribute_value_enumerationFormCallback.attribute_value_enumeration = attribute_value_enumeration
	attribute_value_enumerationFormCallback.formGroup = formGroup

	attribute_value_enumerationFormCallback.CreationMode = (attribute_value_enumeration == nil)

	return
}

type ATTRIBUTE_VALUE_ENUMERATIONFormCallback struct {
	attribute_value_enumeration *models.ATTRIBUTE_VALUE_ENUMERATION

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (attribute_value_enumerationFormCallback *ATTRIBUTE_VALUE_ENUMERATIONFormCallback) OnSave() {
	attribute_value_enumerationFormCallback.probe.stageOfInterest.Lock()
	defer attribute_value_enumerationFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("ATTRIBUTE_VALUE_ENUMERATIONFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	attribute_value_enumerationFormCallback.probe.formStage.Checkout()

	if attribute_value_enumerationFormCallback.attribute_value_enumeration == nil {
		attribute_value_enumerationFormCallback.attribute_value_enumeration = new(models.ATTRIBUTE_VALUE_ENUMERATION).Stage(attribute_value_enumerationFormCallback.probe.stageOfInterest)
	}
	attribute_value_enumeration_ := attribute_value_enumerationFormCallback.attribute_value_enumeration
	_ = attribute_value_enumeration_

	for _, formDiv := range attribute_value_enumerationFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(attribute_value_enumeration_.Name), formDiv)
		case "DEFINITION":
			FormDivSelectFieldToField(&(attribute_value_enumeration_.DEFINITION), attribute_value_enumerationFormCallback.probe.stageOfInterest, formDiv)
		case "VALUES":
			FormDivSelectFieldToField(&(attribute_value_enumeration_.VALUES), attribute_value_enumerationFormCallback.probe.stageOfInterest, formDiv)
		case "A_ATTRIBUTE_VALUE_ENUMERATION:ATTRIBUTE_VALUE_ENUMERATION":
			// WARNING : this form deals with the N-N association "A_ATTRIBUTE_VALUE_ENUMERATION.ATTRIBUTE_VALUE_ENUMERATION []*ATTRIBUTE_VALUE_ENUMERATION" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of ATTRIBUTE_VALUE_ENUMERATION). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.A_ATTRIBUTE_VALUE_ENUMERATION
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "A_ATTRIBUTE_VALUE_ENUMERATION"
				rf.Fieldname = "ATTRIBUTE_VALUE_ENUMERATION"
				formerAssociationSource := attribute_value_enumeration_.GongGetReverseFieldOwner(
					attribute_value_enumerationFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.A_ATTRIBUTE_VALUE_ENUMERATION)
					if !ok {
						log.Fatalln("Source of A_ATTRIBUTE_VALUE_ENUMERATION.ATTRIBUTE_VALUE_ENUMERATION []*ATTRIBUTE_VALUE_ENUMERATION, is not an A_ATTRIBUTE_VALUE_ENUMERATION instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.ATTRIBUTE_VALUE_ENUMERATION, attribute_value_enumeration_)
					formerSource.ATTRIBUTE_VALUE_ENUMERATION = slices.Delete(formerSource.ATTRIBUTE_VALUE_ENUMERATION, idx, idx+1)
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
			var newSource *models.A_ATTRIBUTE_VALUE_ENUMERATION
			for _a_attribute_value_enumeration := range *models.GetGongstructInstancesSet[models.A_ATTRIBUTE_VALUE_ENUMERATION](attribute_value_enumerationFormCallback.probe.stageOfInterest) {
				// the match is base on the name
				if _a_attribute_value_enumeration.GetName() == newSourceName.GetName() {
					newSource = _a_attribute_value_enumeration // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of A_ATTRIBUTE_VALUE_ENUMERATION.ATTRIBUTE_VALUE_ENUMERATION []*ATTRIBUTE_VALUE_ENUMERATION, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.ATTRIBUTE_VALUE_ENUMERATION = append(newSource.ATTRIBUTE_VALUE_ENUMERATION, attribute_value_enumeration_)
		case "A_ATTRIBUTE_VALUE_XHTML_1:ATTRIBUTE_VALUE_ENUMERATION":
			// WARNING : this form deals with the N-N association "A_ATTRIBUTE_VALUE_XHTML_1.ATTRIBUTE_VALUE_ENUMERATION []*ATTRIBUTE_VALUE_ENUMERATION" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of ATTRIBUTE_VALUE_ENUMERATION). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.A_ATTRIBUTE_VALUE_XHTML_1
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "A_ATTRIBUTE_VALUE_XHTML_1"
				rf.Fieldname = "ATTRIBUTE_VALUE_ENUMERATION"
				formerAssociationSource := attribute_value_enumeration_.GongGetReverseFieldOwner(
					attribute_value_enumerationFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.A_ATTRIBUTE_VALUE_XHTML_1)
					if !ok {
						log.Fatalln("Source of A_ATTRIBUTE_VALUE_XHTML_1.ATTRIBUTE_VALUE_ENUMERATION []*ATTRIBUTE_VALUE_ENUMERATION, is not an A_ATTRIBUTE_VALUE_XHTML_1 instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.ATTRIBUTE_VALUE_ENUMERATION, attribute_value_enumeration_)
					formerSource.ATTRIBUTE_VALUE_ENUMERATION = slices.Delete(formerSource.ATTRIBUTE_VALUE_ENUMERATION, idx, idx+1)
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
			var newSource *models.A_ATTRIBUTE_VALUE_XHTML_1
			for _a_attribute_value_xhtml_1 := range *models.GetGongstructInstancesSet[models.A_ATTRIBUTE_VALUE_XHTML_1](attribute_value_enumerationFormCallback.probe.stageOfInterest) {
				// the match is base on the name
				if _a_attribute_value_xhtml_1.GetName() == newSourceName.GetName() {
					newSource = _a_attribute_value_xhtml_1 // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of A_ATTRIBUTE_VALUE_XHTML_1.ATTRIBUTE_VALUE_ENUMERATION []*ATTRIBUTE_VALUE_ENUMERATION, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.ATTRIBUTE_VALUE_ENUMERATION = append(newSource.ATTRIBUTE_VALUE_ENUMERATION, attribute_value_enumeration_)
		}
	}

	// manage the suppress operation
	if attribute_value_enumerationFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attribute_value_enumeration_.Unstage(attribute_value_enumerationFormCallback.probe.stageOfInterest)
	}

	attribute_value_enumerationFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.ATTRIBUTE_VALUE_ENUMERATION](
		attribute_value_enumerationFormCallback.probe,
	)

	// display a new form by reset the form stage
	if attribute_value_enumerationFormCallback.CreationMode || attribute_value_enumerationFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attribute_value_enumerationFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(attribute_value_enumerationFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ATTRIBUTE_VALUE_ENUMERATIONFormCallback(
			nil,
			attribute_value_enumerationFormCallback.probe,
			newFormGroup,
		)
		attribute_value_enumeration := new(models.ATTRIBUTE_VALUE_ENUMERATION)
		FillUpForm(attribute_value_enumeration, newFormGroup, attribute_value_enumerationFormCallback.probe)
		attribute_value_enumerationFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(attribute_value_enumerationFormCallback.probe)
}

func __gong__New__ATTRIBUTE_VALUE_INTEGERFormCallback(
	attribute_value_integer *models.ATTRIBUTE_VALUE_INTEGER,
	probe *Probe,
	formGroup *table.FormGroup,
) (attribute_value_integerFormCallback *ATTRIBUTE_VALUE_INTEGERFormCallback) {
	attribute_value_integerFormCallback = new(ATTRIBUTE_VALUE_INTEGERFormCallback)
	attribute_value_integerFormCallback.probe = probe
	attribute_value_integerFormCallback.attribute_value_integer = attribute_value_integer
	attribute_value_integerFormCallback.formGroup = formGroup

	attribute_value_integerFormCallback.CreationMode = (attribute_value_integer == nil)

	return
}

type ATTRIBUTE_VALUE_INTEGERFormCallback struct {
	attribute_value_integer *models.ATTRIBUTE_VALUE_INTEGER

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (attribute_value_integerFormCallback *ATTRIBUTE_VALUE_INTEGERFormCallback) OnSave() {
	attribute_value_integerFormCallback.probe.stageOfInterest.Lock()
	defer attribute_value_integerFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("ATTRIBUTE_VALUE_INTEGERFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	attribute_value_integerFormCallback.probe.formStage.Checkout()

	if attribute_value_integerFormCallback.attribute_value_integer == nil {
		attribute_value_integerFormCallback.attribute_value_integer = new(models.ATTRIBUTE_VALUE_INTEGER).Stage(attribute_value_integerFormCallback.probe.stageOfInterest)
	}
	attribute_value_integer_ := attribute_value_integerFormCallback.attribute_value_integer
	_ = attribute_value_integer_

	for _, formDiv := range attribute_value_integerFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(attribute_value_integer_.Name), formDiv)
		case "THE_VALUE":
			FormDivBasicFieldToField(&(attribute_value_integer_.THE_VALUE), formDiv)
		case "DEFINITION":
			FormDivSelectFieldToField(&(attribute_value_integer_.DEFINITION), attribute_value_integerFormCallback.probe.stageOfInterest, formDiv)
		case "A_ATTRIBUTE_VALUE_INTEGER:ATTRIBUTE_VALUE_INTEGER":
			// WARNING : this form deals with the N-N association "A_ATTRIBUTE_VALUE_INTEGER.ATTRIBUTE_VALUE_INTEGER []*ATTRIBUTE_VALUE_INTEGER" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of ATTRIBUTE_VALUE_INTEGER). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.A_ATTRIBUTE_VALUE_INTEGER
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "A_ATTRIBUTE_VALUE_INTEGER"
				rf.Fieldname = "ATTRIBUTE_VALUE_INTEGER"
				formerAssociationSource := attribute_value_integer_.GongGetReverseFieldOwner(
					attribute_value_integerFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.A_ATTRIBUTE_VALUE_INTEGER)
					if !ok {
						log.Fatalln("Source of A_ATTRIBUTE_VALUE_INTEGER.ATTRIBUTE_VALUE_INTEGER []*ATTRIBUTE_VALUE_INTEGER, is not an A_ATTRIBUTE_VALUE_INTEGER instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.ATTRIBUTE_VALUE_INTEGER, attribute_value_integer_)
					formerSource.ATTRIBUTE_VALUE_INTEGER = slices.Delete(formerSource.ATTRIBUTE_VALUE_INTEGER, idx, idx+1)
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
			var newSource *models.A_ATTRIBUTE_VALUE_INTEGER
			for _a_attribute_value_integer := range *models.GetGongstructInstancesSet[models.A_ATTRIBUTE_VALUE_INTEGER](attribute_value_integerFormCallback.probe.stageOfInterest) {
				// the match is base on the name
				if _a_attribute_value_integer.GetName() == newSourceName.GetName() {
					newSource = _a_attribute_value_integer // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of A_ATTRIBUTE_VALUE_INTEGER.ATTRIBUTE_VALUE_INTEGER []*ATTRIBUTE_VALUE_INTEGER, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.ATTRIBUTE_VALUE_INTEGER = append(newSource.ATTRIBUTE_VALUE_INTEGER, attribute_value_integer_)
		case "A_ATTRIBUTE_VALUE_XHTML_1:ATTRIBUTE_VALUE_INTEGER":
			// WARNING : this form deals with the N-N association "A_ATTRIBUTE_VALUE_XHTML_1.ATTRIBUTE_VALUE_INTEGER []*ATTRIBUTE_VALUE_INTEGER" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of ATTRIBUTE_VALUE_INTEGER). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.A_ATTRIBUTE_VALUE_XHTML_1
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "A_ATTRIBUTE_VALUE_XHTML_1"
				rf.Fieldname = "ATTRIBUTE_VALUE_INTEGER"
				formerAssociationSource := attribute_value_integer_.GongGetReverseFieldOwner(
					attribute_value_integerFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.A_ATTRIBUTE_VALUE_XHTML_1)
					if !ok {
						log.Fatalln("Source of A_ATTRIBUTE_VALUE_XHTML_1.ATTRIBUTE_VALUE_INTEGER []*ATTRIBUTE_VALUE_INTEGER, is not an A_ATTRIBUTE_VALUE_XHTML_1 instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.ATTRIBUTE_VALUE_INTEGER, attribute_value_integer_)
					formerSource.ATTRIBUTE_VALUE_INTEGER = slices.Delete(formerSource.ATTRIBUTE_VALUE_INTEGER, idx, idx+1)
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
			var newSource *models.A_ATTRIBUTE_VALUE_XHTML_1
			for _a_attribute_value_xhtml_1 := range *models.GetGongstructInstancesSet[models.A_ATTRIBUTE_VALUE_XHTML_1](attribute_value_integerFormCallback.probe.stageOfInterest) {
				// the match is base on the name
				if _a_attribute_value_xhtml_1.GetName() == newSourceName.GetName() {
					newSource = _a_attribute_value_xhtml_1 // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of A_ATTRIBUTE_VALUE_XHTML_1.ATTRIBUTE_VALUE_INTEGER []*ATTRIBUTE_VALUE_INTEGER, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.ATTRIBUTE_VALUE_INTEGER = append(newSource.ATTRIBUTE_VALUE_INTEGER, attribute_value_integer_)
		}
	}

	// manage the suppress operation
	if attribute_value_integerFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attribute_value_integer_.Unstage(attribute_value_integerFormCallback.probe.stageOfInterest)
	}

	attribute_value_integerFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.ATTRIBUTE_VALUE_INTEGER](
		attribute_value_integerFormCallback.probe,
	)

	// display a new form by reset the form stage
	if attribute_value_integerFormCallback.CreationMode || attribute_value_integerFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attribute_value_integerFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(attribute_value_integerFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ATTRIBUTE_VALUE_INTEGERFormCallback(
			nil,
			attribute_value_integerFormCallback.probe,
			newFormGroup,
		)
		attribute_value_integer := new(models.ATTRIBUTE_VALUE_INTEGER)
		FillUpForm(attribute_value_integer, newFormGroup, attribute_value_integerFormCallback.probe)
		attribute_value_integerFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(attribute_value_integerFormCallback.probe)
}

func __gong__New__ATTRIBUTE_VALUE_REALFormCallback(
	attribute_value_real *models.ATTRIBUTE_VALUE_REAL,
	probe *Probe,
	formGroup *table.FormGroup,
) (attribute_value_realFormCallback *ATTRIBUTE_VALUE_REALFormCallback) {
	attribute_value_realFormCallback = new(ATTRIBUTE_VALUE_REALFormCallback)
	attribute_value_realFormCallback.probe = probe
	attribute_value_realFormCallback.attribute_value_real = attribute_value_real
	attribute_value_realFormCallback.formGroup = formGroup

	attribute_value_realFormCallback.CreationMode = (attribute_value_real == nil)

	return
}

type ATTRIBUTE_VALUE_REALFormCallback struct {
	attribute_value_real *models.ATTRIBUTE_VALUE_REAL

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (attribute_value_realFormCallback *ATTRIBUTE_VALUE_REALFormCallback) OnSave() {
	attribute_value_realFormCallback.probe.stageOfInterest.Lock()
	defer attribute_value_realFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("ATTRIBUTE_VALUE_REALFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	attribute_value_realFormCallback.probe.formStage.Checkout()

	if attribute_value_realFormCallback.attribute_value_real == nil {
		attribute_value_realFormCallback.attribute_value_real = new(models.ATTRIBUTE_VALUE_REAL).Stage(attribute_value_realFormCallback.probe.stageOfInterest)
	}
	attribute_value_real_ := attribute_value_realFormCallback.attribute_value_real
	_ = attribute_value_real_

	for _, formDiv := range attribute_value_realFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(attribute_value_real_.Name), formDiv)
		case "THE_VALUE":
			FormDivBasicFieldToField(&(attribute_value_real_.THE_VALUE), formDiv)
		case "DEFINITION":
			FormDivSelectFieldToField(&(attribute_value_real_.DEFINITION), attribute_value_realFormCallback.probe.stageOfInterest, formDiv)
		case "A_ATTRIBUTE_VALUE_REAL:ATTRIBUTE_VALUE_REAL":
			// WARNING : this form deals with the N-N association "A_ATTRIBUTE_VALUE_REAL.ATTRIBUTE_VALUE_REAL []*ATTRIBUTE_VALUE_REAL" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of ATTRIBUTE_VALUE_REAL). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.A_ATTRIBUTE_VALUE_REAL
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "A_ATTRIBUTE_VALUE_REAL"
				rf.Fieldname = "ATTRIBUTE_VALUE_REAL"
				formerAssociationSource := attribute_value_real_.GongGetReverseFieldOwner(
					attribute_value_realFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.A_ATTRIBUTE_VALUE_REAL)
					if !ok {
						log.Fatalln("Source of A_ATTRIBUTE_VALUE_REAL.ATTRIBUTE_VALUE_REAL []*ATTRIBUTE_VALUE_REAL, is not an A_ATTRIBUTE_VALUE_REAL instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.ATTRIBUTE_VALUE_REAL, attribute_value_real_)
					formerSource.ATTRIBUTE_VALUE_REAL = slices.Delete(formerSource.ATTRIBUTE_VALUE_REAL, idx, idx+1)
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
			var newSource *models.A_ATTRIBUTE_VALUE_REAL
			for _a_attribute_value_real := range *models.GetGongstructInstancesSet[models.A_ATTRIBUTE_VALUE_REAL](attribute_value_realFormCallback.probe.stageOfInterest) {
				// the match is base on the name
				if _a_attribute_value_real.GetName() == newSourceName.GetName() {
					newSource = _a_attribute_value_real // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of A_ATTRIBUTE_VALUE_REAL.ATTRIBUTE_VALUE_REAL []*ATTRIBUTE_VALUE_REAL, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.ATTRIBUTE_VALUE_REAL = append(newSource.ATTRIBUTE_VALUE_REAL, attribute_value_real_)
		case "A_ATTRIBUTE_VALUE_XHTML_1:ATTRIBUTE_VALUE_REAL":
			// WARNING : this form deals with the N-N association "A_ATTRIBUTE_VALUE_XHTML_1.ATTRIBUTE_VALUE_REAL []*ATTRIBUTE_VALUE_REAL" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of ATTRIBUTE_VALUE_REAL). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.A_ATTRIBUTE_VALUE_XHTML_1
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "A_ATTRIBUTE_VALUE_XHTML_1"
				rf.Fieldname = "ATTRIBUTE_VALUE_REAL"
				formerAssociationSource := attribute_value_real_.GongGetReverseFieldOwner(
					attribute_value_realFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.A_ATTRIBUTE_VALUE_XHTML_1)
					if !ok {
						log.Fatalln("Source of A_ATTRIBUTE_VALUE_XHTML_1.ATTRIBUTE_VALUE_REAL []*ATTRIBUTE_VALUE_REAL, is not an A_ATTRIBUTE_VALUE_XHTML_1 instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.ATTRIBUTE_VALUE_REAL, attribute_value_real_)
					formerSource.ATTRIBUTE_VALUE_REAL = slices.Delete(formerSource.ATTRIBUTE_VALUE_REAL, idx, idx+1)
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
			var newSource *models.A_ATTRIBUTE_VALUE_XHTML_1
			for _a_attribute_value_xhtml_1 := range *models.GetGongstructInstancesSet[models.A_ATTRIBUTE_VALUE_XHTML_1](attribute_value_realFormCallback.probe.stageOfInterest) {
				// the match is base on the name
				if _a_attribute_value_xhtml_1.GetName() == newSourceName.GetName() {
					newSource = _a_attribute_value_xhtml_1 // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of A_ATTRIBUTE_VALUE_XHTML_1.ATTRIBUTE_VALUE_REAL []*ATTRIBUTE_VALUE_REAL, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.ATTRIBUTE_VALUE_REAL = append(newSource.ATTRIBUTE_VALUE_REAL, attribute_value_real_)
		}
	}

	// manage the suppress operation
	if attribute_value_realFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attribute_value_real_.Unstage(attribute_value_realFormCallback.probe.stageOfInterest)
	}

	attribute_value_realFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.ATTRIBUTE_VALUE_REAL](
		attribute_value_realFormCallback.probe,
	)

	// display a new form by reset the form stage
	if attribute_value_realFormCallback.CreationMode || attribute_value_realFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attribute_value_realFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(attribute_value_realFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ATTRIBUTE_VALUE_REALFormCallback(
			nil,
			attribute_value_realFormCallback.probe,
			newFormGroup,
		)
		attribute_value_real := new(models.ATTRIBUTE_VALUE_REAL)
		FillUpForm(attribute_value_real, newFormGroup, attribute_value_realFormCallback.probe)
		attribute_value_realFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(attribute_value_realFormCallback.probe)
}

func __gong__New__ATTRIBUTE_VALUE_STRINGFormCallback(
	attribute_value_string *models.ATTRIBUTE_VALUE_STRING,
	probe *Probe,
	formGroup *table.FormGroup,
) (attribute_value_stringFormCallback *ATTRIBUTE_VALUE_STRINGFormCallback) {
	attribute_value_stringFormCallback = new(ATTRIBUTE_VALUE_STRINGFormCallback)
	attribute_value_stringFormCallback.probe = probe
	attribute_value_stringFormCallback.attribute_value_string = attribute_value_string
	attribute_value_stringFormCallback.formGroup = formGroup

	attribute_value_stringFormCallback.CreationMode = (attribute_value_string == nil)

	return
}

type ATTRIBUTE_VALUE_STRINGFormCallback struct {
	attribute_value_string *models.ATTRIBUTE_VALUE_STRING

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (attribute_value_stringFormCallback *ATTRIBUTE_VALUE_STRINGFormCallback) OnSave() {
	attribute_value_stringFormCallback.probe.stageOfInterest.Lock()
	defer attribute_value_stringFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("ATTRIBUTE_VALUE_STRINGFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	attribute_value_stringFormCallback.probe.formStage.Checkout()

	if attribute_value_stringFormCallback.attribute_value_string == nil {
		attribute_value_stringFormCallback.attribute_value_string = new(models.ATTRIBUTE_VALUE_STRING).Stage(attribute_value_stringFormCallback.probe.stageOfInterest)
	}
	attribute_value_string_ := attribute_value_stringFormCallback.attribute_value_string
	_ = attribute_value_string_

	for _, formDiv := range attribute_value_stringFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(attribute_value_string_.Name), formDiv)
		case "THE_VALUE":
			FormDivBasicFieldToField(&(attribute_value_string_.THE_VALUE), formDiv)
		case "DEFINITION":
			FormDivSelectFieldToField(&(attribute_value_string_.DEFINITION), attribute_value_stringFormCallback.probe.stageOfInterest, formDiv)
		case "A_ATTRIBUTE_VALUE_STRING:ATTRIBUTE_VALUE_STRING":
			// WARNING : this form deals with the N-N association "A_ATTRIBUTE_VALUE_STRING.ATTRIBUTE_VALUE_STRING []*ATTRIBUTE_VALUE_STRING" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of ATTRIBUTE_VALUE_STRING). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.A_ATTRIBUTE_VALUE_STRING
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "A_ATTRIBUTE_VALUE_STRING"
				rf.Fieldname = "ATTRIBUTE_VALUE_STRING"
				formerAssociationSource := attribute_value_string_.GongGetReverseFieldOwner(
					attribute_value_stringFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.A_ATTRIBUTE_VALUE_STRING)
					if !ok {
						log.Fatalln("Source of A_ATTRIBUTE_VALUE_STRING.ATTRIBUTE_VALUE_STRING []*ATTRIBUTE_VALUE_STRING, is not an A_ATTRIBUTE_VALUE_STRING instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.ATTRIBUTE_VALUE_STRING, attribute_value_string_)
					formerSource.ATTRIBUTE_VALUE_STRING = slices.Delete(formerSource.ATTRIBUTE_VALUE_STRING, idx, idx+1)
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
			var newSource *models.A_ATTRIBUTE_VALUE_STRING
			for _a_attribute_value_string := range *models.GetGongstructInstancesSet[models.A_ATTRIBUTE_VALUE_STRING](attribute_value_stringFormCallback.probe.stageOfInterest) {
				// the match is base on the name
				if _a_attribute_value_string.GetName() == newSourceName.GetName() {
					newSource = _a_attribute_value_string // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of A_ATTRIBUTE_VALUE_STRING.ATTRIBUTE_VALUE_STRING []*ATTRIBUTE_VALUE_STRING, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.ATTRIBUTE_VALUE_STRING = append(newSource.ATTRIBUTE_VALUE_STRING, attribute_value_string_)
		case "A_ATTRIBUTE_VALUE_XHTML_1:ATTRIBUTE_VALUE_STRING":
			// WARNING : this form deals with the N-N association "A_ATTRIBUTE_VALUE_XHTML_1.ATTRIBUTE_VALUE_STRING []*ATTRIBUTE_VALUE_STRING" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of ATTRIBUTE_VALUE_STRING). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.A_ATTRIBUTE_VALUE_XHTML_1
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "A_ATTRIBUTE_VALUE_XHTML_1"
				rf.Fieldname = "ATTRIBUTE_VALUE_STRING"
				formerAssociationSource := attribute_value_string_.GongGetReverseFieldOwner(
					attribute_value_stringFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.A_ATTRIBUTE_VALUE_XHTML_1)
					if !ok {
						log.Fatalln("Source of A_ATTRIBUTE_VALUE_XHTML_1.ATTRIBUTE_VALUE_STRING []*ATTRIBUTE_VALUE_STRING, is not an A_ATTRIBUTE_VALUE_XHTML_1 instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.ATTRIBUTE_VALUE_STRING, attribute_value_string_)
					formerSource.ATTRIBUTE_VALUE_STRING = slices.Delete(formerSource.ATTRIBUTE_VALUE_STRING, idx, idx+1)
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
			var newSource *models.A_ATTRIBUTE_VALUE_XHTML_1
			for _a_attribute_value_xhtml_1 := range *models.GetGongstructInstancesSet[models.A_ATTRIBUTE_VALUE_XHTML_1](attribute_value_stringFormCallback.probe.stageOfInterest) {
				// the match is base on the name
				if _a_attribute_value_xhtml_1.GetName() == newSourceName.GetName() {
					newSource = _a_attribute_value_xhtml_1 // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of A_ATTRIBUTE_VALUE_XHTML_1.ATTRIBUTE_VALUE_STRING []*ATTRIBUTE_VALUE_STRING, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.ATTRIBUTE_VALUE_STRING = append(newSource.ATTRIBUTE_VALUE_STRING, attribute_value_string_)
		}
	}

	// manage the suppress operation
	if attribute_value_stringFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attribute_value_string_.Unstage(attribute_value_stringFormCallback.probe.stageOfInterest)
	}

	attribute_value_stringFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.ATTRIBUTE_VALUE_STRING](
		attribute_value_stringFormCallback.probe,
	)

	// display a new form by reset the form stage
	if attribute_value_stringFormCallback.CreationMode || attribute_value_stringFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attribute_value_stringFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(attribute_value_stringFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ATTRIBUTE_VALUE_STRINGFormCallback(
			nil,
			attribute_value_stringFormCallback.probe,
			newFormGroup,
		)
		attribute_value_string := new(models.ATTRIBUTE_VALUE_STRING)
		FillUpForm(attribute_value_string, newFormGroup, attribute_value_stringFormCallback.probe)
		attribute_value_stringFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(attribute_value_stringFormCallback.probe)
}

func __gong__New__ATTRIBUTE_VALUE_XHTMLFormCallback(
	attribute_value_xhtml *models.ATTRIBUTE_VALUE_XHTML,
	probe *Probe,
	formGroup *table.FormGroup,
) (attribute_value_xhtmlFormCallback *ATTRIBUTE_VALUE_XHTMLFormCallback) {
	attribute_value_xhtmlFormCallback = new(ATTRIBUTE_VALUE_XHTMLFormCallback)
	attribute_value_xhtmlFormCallback.probe = probe
	attribute_value_xhtmlFormCallback.attribute_value_xhtml = attribute_value_xhtml
	attribute_value_xhtmlFormCallback.formGroup = formGroup

	attribute_value_xhtmlFormCallback.CreationMode = (attribute_value_xhtml == nil)

	return
}

type ATTRIBUTE_VALUE_XHTMLFormCallback struct {
	attribute_value_xhtml *models.ATTRIBUTE_VALUE_XHTML

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (attribute_value_xhtmlFormCallback *ATTRIBUTE_VALUE_XHTMLFormCallback) OnSave() {
	attribute_value_xhtmlFormCallback.probe.stageOfInterest.Lock()
	defer attribute_value_xhtmlFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("ATTRIBUTE_VALUE_XHTMLFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	attribute_value_xhtmlFormCallback.probe.formStage.Checkout()

	if attribute_value_xhtmlFormCallback.attribute_value_xhtml == nil {
		attribute_value_xhtmlFormCallback.attribute_value_xhtml = new(models.ATTRIBUTE_VALUE_XHTML).Stage(attribute_value_xhtmlFormCallback.probe.stageOfInterest)
	}
	attribute_value_xhtml_ := attribute_value_xhtmlFormCallback.attribute_value_xhtml
	_ = attribute_value_xhtml_

	for _, formDiv := range attribute_value_xhtmlFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(attribute_value_xhtml_.Name), formDiv)
		case "IS_SIMPLIFIED":
			FormDivBasicFieldToField(&(attribute_value_xhtml_.IS_SIMPLIFIED), formDiv)
		case "THE_VALUE":
			FormDivSelectFieldToField(&(attribute_value_xhtml_.THE_VALUE), attribute_value_xhtmlFormCallback.probe.stageOfInterest, formDiv)
		case "THE_ORIGINAL_VALUE":
			FormDivSelectFieldToField(&(attribute_value_xhtml_.THE_ORIGINAL_VALUE), attribute_value_xhtmlFormCallback.probe.stageOfInterest, formDiv)
		case "DEFINITION":
			FormDivSelectFieldToField(&(attribute_value_xhtml_.DEFINITION), attribute_value_xhtmlFormCallback.probe.stageOfInterest, formDiv)
		case "A_ATTRIBUTE_VALUE_XHTML:ATTRIBUTE_VALUE_XHTML":
			// WARNING : this form deals with the N-N association "A_ATTRIBUTE_VALUE_XHTML.ATTRIBUTE_VALUE_XHTML []*ATTRIBUTE_VALUE_XHTML" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of ATTRIBUTE_VALUE_XHTML). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.A_ATTRIBUTE_VALUE_XHTML
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "A_ATTRIBUTE_VALUE_XHTML"
				rf.Fieldname = "ATTRIBUTE_VALUE_XHTML"
				formerAssociationSource := attribute_value_xhtml_.GongGetReverseFieldOwner(
					attribute_value_xhtmlFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.A_ATTRIBUTE_VALUE_XHTML)
					if !ok {
						log.Fatalln("Source of A_ATTRIBUTE_VALUE_XHTML.ATTRIBUTE_VALUE_XHTML []*ATTRIBUTE_VALUE_XHTML, is not an A_ATTRIBUTE_VALUE_XHTML instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.ATTRIBUTE_VALUE_XHTML, attribute_value_xhtml_)
					formerSource.ATTRIBUTE_VALUE_XHTML = slices.Delete(formerSource.ATTRIBUTE_VALUE_XHTML, idx, idx+1)
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
			var newSource *models.A_ATTRIBUTE_VALUE_XHTML
			for _a_attribute_value_xhtml := range *models.GetGongstructInstancesSet[models.A_ATTRIBUTE_VALUE_XHTML](attribute_value_xhtmlFormCallback.probe.stageOfInterest) {
				// the match is base on the name
				if _a_attribute_value_xhtml.GetName() == newSourceName.GetName() {
					newSource = _a_attribute_value_xhtml // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of A_ATTRIBUTE_VALUE_XHTML.ATTRIBUTE_VALUE_XHTML []*ATTRIBUTE_VALUE_XHTML, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.ATTRIBUTE_VALUE_XHTML = append(newSource.ATTRIBUTE_VALUE_XHTML, attribute_value_xhtml_)
		case "A_ATTRIBUTE_VALUE_XHTML_1:ATTRIBUTE_VALUE_XHTML":
			// WARNING : this form deals with the N-N association "A_ATTRIBUTE_VALUE_XHTML_1.ATTRIBUTE_VALUE_XHTML []*ATTRIBUTE_VALUE_XHTML" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of ATTRIBUTE_VALUE_XHTML). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.A_ATTRIBUTE_VALUE_XHTML_1
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "A_ATTRIBUTE_VALUE_XHTML_1"
				rf.Fieldname = "ATTRIBUTE_VALUE_XHTML"
				formerAssociationSource := attribute_value_xhtml_.GongGetReverseFieldOwner(
					attribute_value_xhtmlFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.A_ATTRIBUTE_VALUE_XHTML_1)
					if !ok {
						log.Fatalln("Source of A_ATTRIBUTE_VALUE_XHTML_1.ATTRIBUTE_VALUE_XHTML []*ATTRIBUTE_VALUE_XHTML, is not an A_ATTRIBUTE_VALUE_XHTML_1 instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.ATTRIBUTE_VALUE_XHTML, attribute_value_xhtml_)
					formerSource.ATTRIBUTE_VALUE_XHTML = slices.Delete(formerSource.ATTRIBUTE_VALUE_XHTML, idx, idx+1)
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
			var newSource *models.A_ATTRIBUTE_VALUE_XHTML_1
			for _a_attribute_value_xhtml_1 := range *models.GetGongstructInstancesSet[models.A_ATTRIBUTE_VALUE_XHTML_1](attribute_value_xhtmlFormCallback.probe.stageOfInterest) {
				// the match is base on the name
				if _a_attribute_value_xhtml_1.GetName() == newSourceName.GetName() {
					newSource = _a_attribute_value_xhtml_1 // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of A_ATTRIBUTE_VALUE_XHTML_1.ATTRIBUTE_VALUE_XHTML []*ATTRIBUTE_VALUE_XHTML, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.ATTRIBUTE_VALUE_XHTML = append(newSource.ATTRIBUTE_VALUE_XHTML, attribute_value_xhtml_)
		}
	}

	// manage the suppress operation
	if attribute_value_xhtmlFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attribute_value_xhtml_.Unstage(attribute_value_xhtmlFormCallback.probe.stageOfInterest)
	}

	attribute_value_xhtmlFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.ATTRIBUTE_VALUE_XHTML](
		attribute_value_xhtmlFormCallback.probe,
	)

	// display a new form by reset the form stage
	if attribute_value_xhtmlFormCallback.CreationMode || attribute_value_xhtmlFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attribute_value_xhtmlFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(attribute_value_xhtmlFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ATTRIBUTE_VALUE_XHTMLFormCallback(
			nil,
			attribute_value_xhtmlFormCallback.probe,
			newFormGroup,
		)
		attribute_value_xhtml := new(models.ATTRIBUTE_VALUE_XHTML)
		FillUpForm(attribute_value_xhtml, newFormGroup, attribute_value_xhtmlFormCallback.probe)
		attribute_value_xhtmlFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(attribute_value_xhtmlFormCallback.probe)
}

func __gong__New__A_ALTERNATIVE_IDFormCallback(
	a_alternative_id *models.A_ALTERNATIVE_ID,
	probe *Probe,
	formGroup *table.FormGroup,
) (a_alternative_idFormCallback *A_ALTERNATIVE_IDFormCallback) {
	a_alternative_idFormCallback = new(A_ALTERNATIVE_IDFormCallback)
	a_alternative_idFormCallback.probe = probe
	a_alternative_idFormCallback.a_alternative_id = a_alternative_id
	a_alternative_idFormCallback.formGroup = formGroup

	a_alternative_idFormCallback.CreationMode = (a_alternative_id == nil)

	return
}

type A_ALTERNATIVE_IDFormCallback struct {
	a_alternative_id *models.A_ALTERNATIVE_ID

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (a_alternative_idFormCallback *A_ALTERNATIVE_IDFormCallback) OnSave() {
	a_alternative_idFormCallback.probe.stageOfInterest.Lock()
	defer a_alternative_idFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("A_ALTERNATIVE_IDFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	a_alternative_idFormCallback.probe.formStage.Checkout()

	if a_alternative_idFormCallback.a_alternative_id == nil {
		a_alternative_idFormCallback.a_alternative_id = new(models.A_ALTERNATIVE_ID).Stage(a_alternative_idFormCallback.probe.stageOfInterest)
	}
	a_alternative_id_ := a_alternative_idFormCallback.a_alternative_id
	_ = a_alternative_id_

	for _, formDiv := range a_alternative_idFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(a_alternative_id_.Name), formDiv)
		case "ALTERNATIVE_ID":
			FormDivSelectFieldToField(&(a_alternative_id_.ALTERNATIVE_ID), a_alternative_idFormCallback.probe.stageOfInterest, formDiv)
		}
	}

	// manage the suppress operation
	if a_alternative_idFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_alternative_id_.Unstage(a_alternative_idFormCallback.probe.stageOfInterest)
	}

	a_alternative_idFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.A_ALTERNATIVE_ID](
		a_alternative_idFormCallback.probe,
	)

	// display a new form by reset the form stage
	if a_alternative_idFormCallback.CreationMode || a_alternative_idFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_alternative_idFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(a_alternative_idFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__A_ALTERNATIVE_IDFormCallback(
			nil,
			a_alternative_idFormCallback.probe,
			newFormGroup,
		)
		a_alternative_id := new(models.A_ALTERNATIVE_ID)
		FillUpForm(a_alternative_id, newFormGroup, a_alternative_idFormCallback.probe)
		a_alternative_idFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(a_alternative_idFormCallback.probe)
}

func __gong__New__A_ATTRIBUTE_DEFINITION_BOOLEAN_REFFormCallback(
	a_attribute_definition_boolean_ref *models.A_ATTRIBUTE_DEFINITION_BOOLEAN_REF,
	probe *Probe,
	formGroup *table.FormGroup,
) (a_attribute_definition_boolean_refFormCallback *A_ATTRIBUTE_DEFINITION_BOOLEAN_REFFormCallback) {
	a_attribute_definition_boolean_refFormCallback = new(A_ATTRIBUTE_DEFINITION_BOOLEAN_REFFormCallback)
	a_attribute_definition_boolean_refFormCallback.probe = probe
	a_attribute_definition_boolean_refFormCallback.a_attribute_definition_boolean_ref = a_attribute_definition_boolean_ref
	a_attribute_definition_boolean_refFormCallback.formGroup = formGroup

	a_attribute_definition_boolean_refFormCallback.CreationMode = (a_attribute_definition_boolean_ref == nil)

	return
}

type A_ATTRIBUTE_DEFINITION_BOOLEAN_REFFormCallback struct {
	a_attribute_definition_boolean_ref *models.A_ATTRIBUTE_DEFINITION_BOOLEAN_REF

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (a_attribute_definition_boolean_refFormCallback *A_ATTRIBUTE_DEFINITION_BOOLEAN_REFFormCallback) OnSave() {
	a_attribute_definition_boolean_refFormCallback.probe.stageOfInterest.Lock()
	defer a_attribute_definition_boolean_refFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("A_ATTRIBUTE_DEFINITION_BOOLEAN_REFFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	a_attribute_definition_boolean_refFormCallback.probe.formStage.Checkout()

	if a_attribute_definition_boolean_refFormCallback.a_attribute_definition_boolean_ref == nil {
		a_attribute_definition_boolean_refFormCallback.a_attribute_definition_boolean_ref = new(models.A_ATTRIBUTE_DEFINITION_BOOLEAN_REF).Stage(a_attribute_definition_boolean_refFormCallback.probe.stageOfInterest)
	}
	a_attribute_definition_boolean_ref_ := a_attribute_definition_boolean_refFormCallback.a_attribute_definition_boolean_ref
	_ = a_attribute_definition_boolean_ref_

	for _, formDiv := range a_attribute_definition_boolean_refFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(a_attribute_definition_boolean_ref_.Name), formDiv)
		case "ATTRIBUTE_DEFINITION_BOOLEAN_REF":
			FormDivBasicFieldToField(&(a_attribute_definition_boolean_ref_.ATTRIBUTE_DEFINITION_BOOLEAN_REF), formDiv)
		}
	}

	// manage the suppress operation
	if a_attribute_definition_boolean_refFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_attribute_definition_boolean_ref_.Unstage(a_attribute_definition_boolean_refFormCallback.probe.stageOfInterest)
	}

	a_attribute_definition_boolean_refFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.A_ATTRIBUTE_DEFINITION_BOOLEAN_REF](
		a_attribute_definition_boolean_refFormCallback.probe,
	)

	// display a new form by reset the form stage
	if a_attribute_definition_boolean_refFormCallback.CreationMode || a_attribute_definition_boolean_refFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_attribute_definition_boolean_refFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(a_attribute_definition_boolean_refFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__A_ATTRIBUTE_DEFINITION_BOOLEAN_REFFormCallback(
			nil,
			a_attribute_definition_boolean_refFormCallback.probe,
			newFormGroup,
		)
		a_attribute_definition_boolean_ref := new(models.A_ATTRIBUTE_DEFINITION_BOOLEAN_REF)
		FillUpForm(a_attribute_definition_boolean_ref, newFormGroup, a_attribute_definition_boolean_refFormCallback.probe)
		a_attribute_definition_boolean_refFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(a_attribute_definition_boolean_refFormCallback.probe)
}

func __gong__New__A_ATTRIBUTE_DEFINITION_DATE_REFFormCallback(
	a_attribute_definition_date_ref *models.A_ATTRIBUTE_DEFINITION_DATE_REF,
	probe *Probe,
	formGroup *table.FormGroup,
) (a_attribute_definition_date_refFormCallback *A_ATTRIBUTE_DEFINITION_DATE_REFFormCallback) {
	a_attribute_definition_date_refFormCallback = new(A_ATTRIBUTE_DEFINITION_DATE_REFFormCallback)
	a_attribute_definition_date_refFormCallback.probe = probe
	a_attribute_definition_date_refFormCallback.a_attribute_definition_date_ref = a_attribute_definition_date_ref
	a_attribute_definition_date_refFormCallback.formGroup = formGroup

	a_attribute_definition_date_refFormCallback.CreationMode = (a_attribute_definition_date_ref == nil)

	return
}

type A_ATTRIBUTE_DEFINITION_DATE_REFFormCallback struct {
	a_attribute_definition_date_ref *models.A_ATTRIBUTE_DEFINITION_DATE_REF

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (a_attribute_definition_date_refFormCallback *A_ATTRIBUTE_DEFINITION_DATE_REFFormCallback) OnSave() {
	a_attribute_definition_date_refFormCallback.probe.stageOfInterest.Lock()
	defer a_attribute_definition_date_refFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("A_ATTRIBUTE_DEFINITION_DATE_REFFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	a_attribute_definition_date_refFormCallback.probe.formStage.Checkout()

	if a_attribute_definition_date_refFormCallback.a_attribute_definition_date_ref == nil {
		a_attribute_definition_date_refFormCallback.a_attribute_definition_date_ref = new(models.A_ATTRIBUTE_DEFINITION_DATE_REF).Stage(a_attribute_definition_date_refFormCallback.probe.stageOfInterest)
	}
	a_attribute_definition_date_ref_ := a_attribute_definition_date_refFormCallback.a_attribute_definition_date_ref
	_ = a_attribute_definition_date_ref_

	for _, formDiv := range a_attribute_definition_date_refFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(a_attribute_definition_date_ref_.Name), formDiv)
		case "ATTRIBUTE_DEFINITION_DATE_REF":
			FormDivBasicFieldToField(&(a_attribute_definition_date_ref_.ATTRIBUTE_DEFINITION_DATE_REF), formDiv)
		}
	}

	// manage the suppress operation
	if a_attribute_definition_date_refFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_attribute_definition_date_ref_.Unstage(a_attribute_definition_date_refFormCallback.probe.stageOfInterest)
	}

	a_attribute_definition_date_refFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.A_ATTRIBUTE_DEFINITION_DATE_REF](
		a_attribute_definition_date_refFormCallback.probe,
	)

	// display a new form by reset the form stage
	if a_attribute_definition_date_refFormCallback.CreationMode || a_attribute_definition_date_refFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_attribute_definition_date_refFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(a_attribute_definition_date_refFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__A_ATTRIBUTE_DEFINITION_DATE_REFFormCallback(
			nil,
			a_attribute_definition_date_refFormCallback.probe,
			newFormGroup,
		)
		a_attribute_definition_date_ref := new(models.A_ATTRIBUTE_DEFINITION_DATE_REF)
		FillUpForm(a_attribute_definition_date_ref, newFormGroup, a_attribute_definition_date_refFormCallback.probe)
		a_attribute_definition_date_refFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(a_attribute_definition_date_refFormCallback.probe)
}

func __gong__New__A_ATTRIBUTE_DEFINITION_ENUMERATION_REFFormCallback(
	a_attribute_definition_enumeration_ref *models.A_ATTRIBUTE_DEFINITION_ENUMERATION_REF,
	probe *Probe,
	formGroup *table.FormGroup,
) (a_attribute_definition_enumeration_refFormCallback *A_ATTRIBUTE_DEFINITION_ENUMERATION_REFFormCallback) {
	a_attribute_definition_enumeration_refFormCallback = new(A_ATTRIBUTE_DEFINITION_ENUMERATION_REFFormCallback)
	a_attribute_definition_enumeration_refFormCallback.probe = probe
	a_attribute_definition_enumeration_refFormCallback.a_attribute_definition_enumeration_ref = a_attribute_definition_enumeration_ref
	a_attribute_definition_enumeration_refFormCallback.formGroup = formGroup

	a_attribute_definition_enumeration_refFormCallback.CreationMode = (a_attribute_definition_enumeration_ref == nil)

	return
}

type A_ATTRIBUTE_DEFINITION_ENUMERATION_REFFormCallback struct {
	a_attribute_definition_enumeration_ref *models.A_ATTRIBUTE_DEFINITION_ENUMERATION_REF

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (a_attribute_definition_enumeration_refFormCallback *A_ATTRIBUTE_DEFINITION_ENUMERATION_REFFormCallback) OnSave() {
	a_attribute_definition_enumeration_refFormCallback.probe.stageOfInterest.Lock()
	defer a_attribute_definition_enumeration_refFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("A_ATTRIBUTE_DEFINITION_ENUMERATION_REFFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	a_attribute_definition_enumeration_refFormCallback.probe.formStage.Checkout()

	if a_attribute_definition_enumeration_refFormCallback.a_attribute_definition_enumeration_ref == nil {
		a_attribute_definition_enumeration_refFormCallback.a_attribute_definition_enumeration_ref = new(models.A_ATTRIBUTE_DEFINITION_ENUMERATION_REF).Stage(a_attribute_definition_enumeration_refFormCallback.probe.stageOfInterest)
	}
	a_attribute_definition_enumeration_ref_ := a_attribute_definition_enumeration_refFormCallback.a_attribute_definition_enumeration_ref
	_ = a_attribute_definition_enumeration_ref_

	for _, formDiv := range a_attribute_definition_enumeration_refFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(a_attribute_definition_enumeration_ref_.Name), formDiv)
		case "ATTRIBUTE_DEFINITION_ENUMERATION_REF":
			FormDivBasicFieldToField(&(a_attribute_definition_enumeration_ref_.ATTRIBUTE_DEFINITION_ENUMERATION_REF), formDiv)
		}
	}

	// manage the suppress operation
	if a_attribute_definition_enumeration_refFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_attribute_definition_enumeration_ref_.Unstage(a_attribute_definition_enumeration_refFormCallback.probe.stageOfInterest)
	}

	a_attribute_definition_enumeration_refFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.A_ATTRIBUTE_DEFINITION_ENUMERATION_REF](
		a_attribute_definition_enumeration_refFormCallback.probe,
	)

	// display a new form by reset the form stage
	if a_attribute_definition_enumeration_refFormCallback.CreationMode || a_attribute_definition_enumeration_refFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_attribute_definition_enumeration_refFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(a_attribute_definition_enumeration_refFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__A_ATTRIBUTE_DEFINITION_ENUMERATION_REFFormCallback(
			nil,
			a_attribute_definition_enumeration_refFormCallback.probe,
			newFormGroup,
		)
		a_attribute_definition_enumeration_ref := new(models.A_ATTRIBUTE_DEFINITION_ENUMERATION_REF)
		FillUpForm(a_attribute_definition_enumeration_ref, newFormGroup, a_attribute_definition_enumeration_refFormCallback.probe)
		a_attribute_definition_enumeration_refFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(a_attribute_definition_enumeration_refFormCallback.probe)
}

func __gong__New__A_ATTRIBUTE_DEFINITION_INTEGER_REFFormCallback(
	a_attribute_definition_integer_ref *models.A_ATTRIBUTE_DEFINITION_INTEGER_REF,
	probe *Probe,
	formGroup *table.FormGroup,
) (a_attribute_definition_integer_refFormCallback *A_ATTRIBUTE_DEFINITION_INTEGER_REFFormCallback) {
	a_attribute_definition_integer_refFormCallback = new(A_ATTRIBUTE_DEFINITION_INTEGER_REFFormCallback)
	a_attribute_definition_integer_refFormCallback.probe = probe
	a_attribute_definition_integer_refFormCallback.a_attribute_definition_integer_ref = a_attribute_definition_integer_ref
	a_attribute_definition_integer_refFormCallback.formGroup = formGroup

	a_attribute_definition_integer_refFormCallback.CreationMode = (a_attribute_definition_integer_ref == nil)

	return
}

type A_ATTRIBUTE_DEFINITION_INTEGER_REFFormCallback struct {
	a_attribute_definition_integer_ref *models.A_ATTRIBUTE_DEFINITION_INTEGER_REF

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (a_attribute_definition_integer_refFormCallback *A_ATTRIBUTE_DEFINITION_INTEGER_REFFormCallback) OnSave() {
	a_attribute_definition_integer_refFormCallback.probe.stageOfInterest.Lock()
	defer a_attribute_definition_integer_refFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("A_ATTRIBUTE_DEFINITION_INTEGER_REFFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	a_attribute_definition_integer_refFormCallback.probe.formStage.Checkout()

	if a_attribute_definition_integer_refFormCallback.a_attribute_definition_integer_ref == nil {
		a_attribute_definition_integer_refFormCallback.a_attribute_definition_integer_ref = new(models.A_ATTRIBUTE_DEFINITION_INTEGER_REF).Stage(a_attribute_definition_integer_refFormCallback.probe.stageOfInterest)
	}
	a_attribute_definition_integer_ref_ := a_attribute_definition_integer_refFormCallback.a_attribute_definition_integer_ref
	_ = a_attribute_definition_integer_ref_

	for _, formDiv := range a_attribute_definition_integer_refFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(a_attribute_definition_integer_ref_.Name), formDiv)
		case "ATTRIBUTE_DEFINITION_INTEGER_REF":
			FormDivBasicFieldToField(&(a_attribute_definition_integer_ref_.ATTRIBUTE_DEFINITION_INTEGER_REF), formDiv)
		}
	}

	// manage the suppress operation
	if a_attribute_definition_integer_refFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_attribute_definition_integer_ref_.Unstage(a_attribute_definition_integer_refFormCallback.probe.stageOfInterest)
	}

	a_attribute_definition_integer_refFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.A_ATTRIBUTE_DEFINITION_INTEGER_REF](
		a_attribute_definition_integer_refFormCallback.probe,
	)

	// display a new form by reset the form stage
	if a_attribute_definition_integer_refFormCallback.CreationMode || a_attribute_definition_integer_refFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_attribute_definition_integer_refFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(a_attribute_definition_integer_refFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__A_ATTRIBUTE_DEFINITION_INTEGER_REFFormCallback(
			nil,
			a_attribute_definition_integer_refFormCallback.probe,
			newFormGroup,
		)
		a_attribute_definition_integer_ref := new(models.A_ATTRIBUTE_DEFINITION_INTEGER_REF)
		FillUpForm(a_attribute_definition_integer_ref, newFormGroup, a_attribute_definition_integer_refFormCallback.probe)
		a_attribute_definition_integer_refFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(a_attribute_definition_integer_refFormCallback.probe)
}

func __gong__New__A_ATTRIBUTE_DEFINITION_REAL_REFFormCallback(
	a_attribute_definition_real_ref *models.A_ATTRIBUTE_DEFINITION_REAL_REF,
	probe *Probe,
	formGroup *table.FormGroup,
) (a_attribute_definition_real_refFormCallback *A_ATTRIBUTE_DEFINITION_REAL_REFFormCallback) {
	a_attribute_definition_real_refFormCallback = new(A_ATTRIBUTE_DEFINITION_REAL_REFFormCallback)
	a_attribute_definition_real_refFormCallback.probe = probe
	a_attribute_definition_real_refFormCallback.a_attribute_definition_real_ref = a_attribute_definition_real_ref
	a_attribute_definition_real_refFormCallback.formGroup = formGroup

	a_attribute_definition_real_refFormCallback.CreationMode = (a_attribute_definition_real_ref == nil)

	return
}

type A_ATTRIBUTE_DEFINITION_REAL_REFFormCallback struct {
	a_attribute_definition_real_ref *models.A_ATTRIBUTE_DEFINITION_REAL_REF

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (a_attribute_definition_real_refFormCallback *A_ATTRIBUTE_DEFINITION_REAL_REFFormCallback) OnSave() {
	a_attribute_definition_real_refFormCallback.probe.stageOfInterest.Lock()
	defer a_attribute_definition_real_refFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("A_ATTRIBUTE_DEFINITION_REAL_REFFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	a_attribute_definition_real_refFormCallback.probe.formStage.Checkout()

	if a_attribute_definition_real_refFormCallback.a_attribute_definition_real_ref == nil {
		a_attribute_definition_real_refFormCallback.a_attribute_definition_real_ref = new(models.A_ATTRIBUTE_DEFINITION_REAL_REF).Stage(a_attribute_definition_real_refFormCallback.probe.stageOfInterest)
	}
	a_attribute_definition_real_ref_ := a_attribute_definition_real_refFormCallback.a_attribute_definition_real_ref
	_ = a_attribute_definition_real_ref_

	for _, formDiv := range a_attribute_definition_real_refFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(a_attribute_definition_real_ref_.Name), formDiv)
		case "ATTRIBUTE_DEFINITION_REAL_REF":
			FormDivBasicFieldToField(&(a_attribute_definition_real_ref_.ATTRIBUTE_DEFINITION_REAL_REF), formDiv)
		}
	}

	// manage the suppress operation
	if a_attribute_definition_real_refFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_attribute_definition_real_ref_.Unstage(a_attribute_definition_real_refFormCallback.probe.stageOfInterest)
	}

	a_attribute_definition_real_refFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.A_ATTRIBUTE_DEFINITION_REAL_REF](
		a_attribute_definition_real_refFormCallback.probe,
	)

	// display a new form by reset the form stage
	if a_attribute_definition_real_refFormCallback.CreationMode || a_attribute_definition_real_refFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_attribute_definition_real_refFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(a_attribute_definition_real_refFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__A_ATTRIBUTE_DEFINITION_REAL_REFFormCallback(
			nil,
			a_attribute_definition_real_refFormCallback.probe,
			newFormGroup,
		)
		a_attribute_definition_real_ref := new(models.A_ATTRIBUTE_DEFINITION_REAL_REF)
		FillUpForm(a_attribute_definition_real_ref, newFormGroup, a_attribute_definition_real_refFormCallback.probe)
		a_attribute_definition_real_refFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(a_attribute_definition_real_refFormCallback.probe)
}

func __gong__New__A_ATTRIBUTE_DEFINITION_STRING_REFFormCallback(
	a_attribute_definition_string_ref *models.A_ATTRIBUTE_DEFINITION_STRING_REF,
	probe *Probe,
	formGroup *table.FormGroup,
) (a_attribute_definition_string_refFormCallback *A_ATTRIBUTE_DEFINITION_STRING_REFFormCallback) {
	a_attribute_definition_string_refFormCallback = new(A_ATTRIBUTE_DEFINITION_STRING_REFFormCallback)
	a_attribute_definition_string_refFormCallback.probe = probe
	a_attribute_definition_string_refFormCallback.a_attribute_definition_string_ref = a_attribute_definition_string_ref
	a_attribute_definition_string_refFormCallback.formGroup = formGroup

	a_attribute_definition_string_refFormCallback.CreationMode = (a_attribute_definition_string_ref == nil)

	return
}

type A_ATTRIBUTE_DEFINITION_STRING_REFFormCallback struct {
	a_attribute_definition_string_ref *models.A_ATTRIBUTE_DEFINITION_STRING_REF

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (a_attribute_definition_string_refFormCallback *A_ATTRIBUTE_DEFINITION_STRING_REFFormCallback) OnSave() {
	a_attribute_definition_string_refFormCallback.probe.stageOfInterest.Lock()
	defer a_attribute_definition_string_refFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("A_ATTRIBUTE_DEFINITION_STRING_REFFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	a_attribute_definition_string_refFormCallback.probe.formStage.Checkout()

	if a_attribute_definition_string_refFormCallback.a_attribute_definition_string_ref == nil {
		a_attribute_definition_string_refFormCallback.a_attribute_definition_string_ref = new(models.A_ATTRIBUTE_DEFINITION_STRING_REF).Stage(a_attribute_definition_string_refFormCallback.probe.stageOfInterest)
	}
	a_attribute_definition_string_ref_ := a_attribute_definition_string_refFormCallback.a_attribute_definition_string_ref
	_ = a_attribute_definition_string_ref_

	for _, formDiv := range a_attribute_definition_string_refFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(a_attribute_definition_string_ref_.Name), formDiv)
		case "ATTRIBUTE_DEFINITION_STRING_REF":
			FormDivBasicFieldToField(&(a_attribute_definition_string_ref_.ATTRIBUTE_DEFINITION_STRING_REF), formDiv)
		}
	}

	// manage the suppress operation
	if a_attribute_definition_string_refFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_attribute_definition_string_ref_.Unstage(a_attribute_definition_string_refFormCallback.probe.stageOfInterest)
	}

	a_attribute_definition_string_refFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.A_ATTRIBUTE_DEFINITION_STRING_REF](
		a_attribute_definition_string_refFormCallback.probe,
	)

	// display a new form by reset the form stage
	if a_attribute_definition_string_refFormCallback.CreationMode || a_attribute_definition_string_refFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_attribute_definition_string_refFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(a_attribute_definition_string_refFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__A_ATTRIBUTE_DEFINITION_STRING_REFFormCallback(
			nil,
			a_attribute_definition_string_refFormCallback.probe,
			newFormGroup,
		)
		a_attribute_definition_string_ref := new(models.A_ATTRIBUTE_DEFINITION_STRING_REF)
		FillUpForm(a_attribute_definition_string_ref, newFormGroup, a_attribute_definition_string_refFormCallback.probe)
		a_attribute_definition_string_refFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(a_attribute_definition_string_refFormCallback.probe)
}

func __gong__New__A_ATTRIBUTE_DEFINITION_XHTML_REFFormCallback(
	a_attribute_definition_xhtml_ref *models.A_ATTRIBUTE_DEFINITION_XHTML_REF,
	probe *Probe,
	formGroup *table.FormGroup,
) (a_attribute_definition_xhtml_refFormCallback *A_ATTRIBUTE_DEFINITION_XHTML_REFFormCallback) {
	a_attribute_definition_xhtml_refFormCallback = new(A_ATTRIBUTE_DEFINITION_XHTML_REFFormCallback)
	a_attribute_definition_xhtml_refFormCallback.probe = probe
	a_attribute_definition_xhtml_refFormCallback.a_attribute_definition_xhtml_ref = a_attribute_definition_xhtml_ref
	a_attribute_definition_xhtml_refFormCallback.formGroup = formGroup

	a_attribute_definition_xhtml_refFormCallback.CreationMode = (a_attribute_definition_xhtml_ref == nil)

	return
}

type A_ATTRIBUTE_DEFINITION_XHTML_REFFormCallback struct {
	a_attribute_definition_xhtml_ref *models.A_ATTRIBUTE_DEFINITION_XHTML_REF

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (a_attribute_definition_xhtml_refFormCallback *A_ATTRIBUTE_DEFINITION_XHTML_REFFormCallback) OnSave() {
	a_attribute_definition_xhtml_refFormCallback.probe.stageOfInterest.Lock()
	defer a_attribute_definition_xhtml_refFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("A_ATTRIBUTE_DEFINITION_XHTML_REFFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	a_attribute_definition_xhtml_refFormCallback.probe.formStage.Checkout()

	if a_attribute_definition_xhtml_refFormCallback.a_attribute_definition_xhtml_ref == nil {
		a_attribute_definition_xhtml_refFormCallback.a_attribute_definition_xhtml_ref = new(models.A_ATTRIBUTE_DEFINITION_XHTML_REF).Stage(a_attribute_definition_xhtml_refFormCallback.probe.stageOfInterest)
	}
	a_attribute_definition_xhtml_ref_ := a_attribute_definition_xhtml_refFormCallback.a_attribute_definition_xhtml_ref
	_ = a_attribute_definition_xhtml_ref_

	for _, formDiv := range a_attribute_definition_xhtml_refFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(a_attribute_definition_xhtml_ref_.Name), formDiv)
		case "ATTRIBUTE_DEFINITION_XHTML_REF":
			FormDivBasicFieldToField(&(a_attribute_definition_xhtml_ref_.ATTRIBUTE_DEFINITION_XHTML_REF), formDiv)
		}
	}

	// manage the suppress operation
	if a_attribute_definition_xhtml_refFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_attribute_definition_xhtml_ref_.Unstage(a_attribute_definition_xhtml_refFormCallback.probe.stageOfInterest)
	}

	a_attribute_definition_xhtml_refFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.A_ATTRIBUTE_DEFINITION_XHTML_REF](
		a_attribute_definition_xhtml_refFormCallback.probe,
	)

	// display a new form by reset the form stage
	if a_attribute_definition_xhtml_refFormCallback.CreationMode || a_attribute_definition_xhtml_refFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_attribute_definition_xhtml_refFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(a_attribute_definition_xhtml_refFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__A_ATTRIBUTE_DEFINITION_XHTML_REFFormCallback(
			nil,
			a_attribute_definition_xhtml_refFormCallback.probe,
			newFormGroup,
		)
		a_attribute_definition_xhtml_ref := new(models.A_ATTRIBUTE_DEFINITION_XHTML_REF)
		FillUpForm(a_attribute_definition_xhtml_ref, newFormGroup, a_attribute_definition_xhtml_refFormCallback.probe)
		a_attribute_definition_xhtml_refFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(a_attribute_definition_xhtml_refFormCallback.probe)
}

func __gong__New__A_ATTRIBUTE_VALUE_BOOLEANFormCallback(
	a_attribute_value_boolean *models.A_ATTRIBUTE_VALUE_BOOLEAN,
	probe *Probe,
	formGroup *table.FormGroup,
) (a_attribute_value_booleanFormCallback *A_ATTRIBUTE_VALUE_BOOLEANFormCallback) {
	a_attribute_value_booleanFormCallback = new(A_ATTRIBUTE_VALUE_BOOLEANFormCallback)
	a_attribute_value_booleanFormCallback.probe = probe
	a_attribute_value_booleanFormCallback.a_attribute_value_boolean = a_attribute_value_boolean
	a_attribute_value_booleanFormCallback.formGroup = formGroup

	a_attribute_value_booleanFormCallback.CreationMode = (a_attribute_value_boolean == nil)

	return
}

type A_ATTRIBUTE_VALUE_BOOLEANFormCallback struct {
	a_attribute_value_boolean *models.A_ATTRIBUTE_VALUE_BOOLEAN

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (a_attribute_value_booleanFormCallback *A_ATTRIBUTE_VALUE_BOOLEANFormCallback) OnSave() {
	a_attribute_value_booleanFormCallback.probe.stageOfInterest.Lock()
	defer a_attribute_value_booleanFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("A_ATTRIBUTE_VALUE_BOOLEANFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	a_attribute_value_booleanFormCallback.probe.formStage.Checkout()

	if a_attribute_value_booleanFormCallback.a_attribute_value_boolean == nil {
		a_attribute_value_booleanFormCallback.a_attribute_value_boolean = new(models.A_ATTRIBUTE_VALUE_BOOLEAN).Stage(a_attribute_value_booleanFormCallback.probe.stageOfInterest)
	}
	a_attribute_value_boolean_ := a_attribute_value_booleanFormCallback.a_attribute_value_boolean
	_ = a_attribute_value_boolean_

	for _, formDiv := range a_attribute_value_booleanFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(a_attribute_value_boolean_.Name), formDiv)
		case "ATTRIBUTE_VALUE_BOOLEAN":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.ATTRIBUTE_VALUE_BOOLEAN](a_attribute_value_booleanFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.ATTRIBUTE_VALUE_BOOLEAN, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.ATTRIBUTE_VALUE_BOOLEAN)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					a_attribute_value_booleanFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.ATTRIBUTE_VALUE_BOOLEAN](a_attribute_value_booleanFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			a_attribute_value_boolean_.ATTRIBUTE_VALUE_BOOLEAN = instanceSlice

		}
	}

	// manage the suppress operation
	if a_attribute_value_booleanFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_attribute_value_boolean_.Unstage(a_attribute_value_booleanFormCallback.probe.stageOfInterest)
	}

	a_attribute_value_booleanFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.A_ATTRIBUTE_VALUE_BOOLEAN](
		a_attribute_value_booleanFormCallback.probe,
	)

	// display a new form by reset the form stage
	if a_attribute_value_booleanFormCallback.CreationMode || a_attribute_value_booleanFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_attribute_value_booleanFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(a_attribute_value_booleanFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__A_ATTRIBUTE_VALUE_BOOLEANFormCallback(
			nil,
			a_attribute_value_booleanFormCallback.probe,
			newFormGroup,
		)
		a_attribute_value_boolean := new(models.A_ATTRIBUTE_VALUE_BOOLEAN)
		FillUpForm(a_attribute_value_boolean, newFormGroup, a_attribute_value_booleanFormCallback.probe)
		a_attribute_value_booleanFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(a_attribute_value_booleanFormCallback.probe)
}

func __gong__New__A_ATTRIBUTE_VALUE_DATEFormCallback(
	a_attribute_value_date *models.A_ATTRIBUTE_VALUE_DATE,
	probe *Probe,
	formGroup *table.FormGroup,
) (a_attribute_value_dateFormCallback *A_ATTRIBUTE_VALUE_DATEFormCallback) {
	a_attribute_value_dateFormCallback = new(A_ATTRIBUTE_VALUE_DATEFormCallback)
	a_attribute_value_dateFormCallback.probe = probe
	a_attribute_value_dateFormCallback.a_attribute_value_date = a_attribute_value_date
	a_attribute_value_dateFormCallback.formGroup = formGroup

	a_attribute_value_dateFormCallback.CreationMode = (a_attribute_value_date == nil)

	return
}

type A_ATTRIBUTE_VALUE_DATEFormCallback struct {
	a_attribute_value_date *models.A_ATTRIBUTE_VALUE_DATE

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (a_attribute_value_dateFormCallback *A_ATTRIBUTE_VALUE_DATEFormCallback) OnSave() {
	a_attribute_value_dateFormCallback.probe.stageOfInterest.Lock()
	defer a_attribute_value_dateFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("A_ATTRIBUTE_VALUE_DATEFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	a_attribute_value_dateFormCallback.probe.formStage.Checkout()

	if a_attribute_value_dateFormCallback.a_attribute_value_date == nil {
		a_attribute_value_dateFormCallback.a_attribute_value_date = new(models.A_ATTRIBUTE_VALUE_DATE).Stage(a_attribute_value_dateFormCallback.probe.stageOfInterest)
	}
	a_attribute_value_date_ := a_attribute_value_dateFormCallback.a_attribute_value_date
	_ = a_attribute_value_date_

	for _, formDiv := range a_attribute_value_dateFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(a_attribute_value_date_.Name), formDiv)
		case "ATTRIBUTE_VALUE_DATE":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.ATTRIBUTE_VALUE_DATE](a_attribute_value_dateFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.ATTRIBUTE_VALUE_DATE, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.ATTRIBUTE_VALUE_DATE)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					a_attribute_value_dateFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.ATTRIBUTE_VALUE_DATE](a_attribute_value_dateFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			a_attribute_value_date_.ATTRIBUTE_VALUE_DATE = instanceSlice

		}
	}

	// manage the suppress operation
	if a_attribute_value_dateFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_attribute_value_date_.Unstage(a_attribute_value_dateFormCallback.probe.stageOfInterest)
	}

	a_attribute_value_dateFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.A_ATTRIBUTE_VALUE_DATE](
		a_attribute_value_dateFormCallback.probe,
	)

	// display a new form by reset the form stage
	if a_attribute_value_dateFormCallback.CreationMode || a_attribute_value_dateFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_attribute_value_dateFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(a_attribute_value_dateFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__A_ATTRIBUTE_VALUE_DATEFormCallback(
			nil,
			a_attribute_value_dateFormCallback.probe,
			newFormGroup,
		)
		a_attribute_value_date := new(models.A_ATTRIBUTE_VALUE_DATE)
		FillUpForm(a_attribute_value_date, newFormGroup, a_attribute_value_dateFormCallback.probe)
		a_attribute_value_dateFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(a_attribute_value_dateFormCallback.probe)
}

func __gong__New__A_ATTRIBUTE_VALUE_ENUMERATIONFormCallback(
	a_attribute_value_enumeration *models.A_ATTRIBUTE_VALUE_ENUMERATION,
	probe *Probe,
	formGroup *table.FormGroup,
) (a_attribute_value_enumerationFormCallback *A_ATTRIBUTE_VALUE_ENUMERATIONFormCallback) {
	a_attribute_value_enumerationFormCallback = new(A_ATTRIBUTE_VALUE_ENUMERATIONFormCallback)
	a_attribute_value_enumerationFormCallback.probe = probe
	a_attribute_value_enumerationFormCallback.a_attribute_value_enumeration = a_attribute_value_enumeration
	a_attribute_value_enumerationFormCallback.formGroup = formGroup

	a_attribute_value_enumerationFormCallback.CreationMode = (a_attribute_value_enumeration == nil)

	return
}

type A_ATTRIBUTE_VALUE_ENUMERATIONFormCallback struct {
	a_attribute_value_enumeration *models.A_ATTRIBUTE_VALUE_ENUMERATION

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (a_attribute_value_enumerationFormCallback *A_ATTRIBUTE_VALUE_ENUMERATIONFormCallback) OnSave() {
	a_attribute_value_enumerationFormCallback.probe.stageOfInterest.Lock()
	defer a_attribute_value_enumerationFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("A_ATTRIBUTE_VALUE_ENUMERATIONFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	a_attribute_value_enumerationFormCallback.probe.formStage.Checkout()

	if a_attribute_value_enumerationFormCallback.a_attribute_value_enumeration == nil {
		a_attribute_value_enumerationFormCallback.a_attribute_value_enumeration = new(models.A_ATTRIBUTE_VALUE_ENUMERATION).Stage(a_attribute_value_enumerationFormCallback.probe.stageOfInterest)
	}
	a_attribute_value_enumeration_ := a_attribute_value_enumerationFormCallback.a_attribute_value_enumeration
	_ = a_attribute_value_enumeration_

	for _, formDiv := range a_attribute_value_enumerationFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(a_attribute_value_enumeration_.Name), formDiv)
		case "ATTRIBUTE_VALUE_ENUMERATION":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.ATTRIBUTE_VALUE_ENUMERATION](a_attribute_value_enumerationFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.ATTRIBUTE_VALUE_ENUMERATION, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.ATTRIBUTE_VALUE_ENUMERATION)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					a_attribute_value_enumerationFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.ATTRIBUTE_VALUE_ENUMERATION](a_attribute_value_enumerationFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			a_attribute_value_enumeration_.ATTRIBUTE_VALUE_ENUMERATION = instanceSlice

		}
	}

	// manage the suppress operation
	if a_attribute_value_enumerationFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_attribute_value_enumeration_.Unstage(a_attribute_value_enumerationFormCallback.probe.stageOfInterest)
	}

	a_attribute_value_enumerationFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.A_ATTRIBUTE_VALUE_ENUMERATION](
		a_attribute_value_enumerationFormCallback.probe,
	)

	// display a new form by reset the form stage
	if a_attribute_value_enumerationFormCallback.CreationMode || a_attribute_value_enumerationFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_attribute_value_enumerationFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(a_attribute_value_enumerationFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__A_ATTRIBUTE_VALUE_ENUMERATIONFormCallback(
			nil,
			a_attribute_value_enumerationFormCallback.probe,
			newFormGroup,
		)
		a_attribute_value_enumeration := new(models.A_ATTRIBUTE_VALUE_ENUMERATION)
		FillUpForm(a_attribute_value_enumeration, newFormGroup, a_attribute_value_enumerationFormCallback.probe)
		a_attribute_value_enumerationFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(a_attribute_value_enumerationFormCallback.probe)
}

func __gong__New__A_ATTRIBUTE_VALUE_INTEGERFormCallback(
	a_attribute_value_integer *models.A_ATTRIBUTE_VALUE_INTEGER,
	probe *Probe,
	formGroup *table.FormGroup,
) (a_attribute_value_integerFormCallback *A_ATTRIBUTE_VALUE_INTEGERFormCallback) {
	a_attribute_value_integerFormCallback = new(A_ATTRIBUTE_VALUE_INTEGERFormCallback)
	a_attribute_value_integerFormCallback.probe = probe
	a_attribute_value_integerFormCallback.a_attribute_value_integer = a_attribute_value_integer
	a_attribute_value_integerFormCallback.formGroup = formGroup

	a_attribute_value_integerFormCallback.CreationMode = (a_attribute_value_integer == nil)

	return
}

type A_ATTRIBUTE_VALUE_INTEGERFormCallback struct {
	a_attribute_value_integer *models.A_ATTRIBUTE_VALUE_INTEGER

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (a_attribute_value_integerFormCallback *A_ATTRIBUTE_VALUE_INTEGERFormCallback) OnSave() {
	a_attribute_value_integerFormCallback.probe.stageOfInterest.Lock()
	defer a_attribute_value_integerFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("A_ATTRIBUTE_VALUE_INTEGERFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	a_attribute_value_integerFormCallback.probe.formStage.Checkout()

	if a_attribute_value_integerFormCallback.a_attribute_value_integer == nil {
		a_attribute_value_integerFormCallback.a_attribute_value_integer = new(models.A_ATTRIBUTE_VALUE_INTEGER).Stage(a_attribute_value_integerFormCallback.probe.stageOfInterest)
	}
	a_attribute_value_integer_ := a_attribute_value_integerFormCallback.a_attribute_value_integer
	_ = a_attribute_value_integer_

	for _, formDiv := range a_attribute_value_integerFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(a_attribute_value_integer_.Name), formDiv)
		case "ATTRIBUTE_VALUE_INTEGER":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.ATTRIBUTE_VALUE_INTEGER](a_attribute_value_integerFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.ATTRIBUTE_VALUE_INTEGER, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.ATTRIBUTE_VALUE_INTEGER)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					a_attribute_value_integerFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.ATTRIBUTE_VALUE_INTEGER](a_attribute_value_integerFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			a_attribute_value_integer_.ATTRIBUTE_VALUE_INTEGER = instanceSlice

		}
	}

	// manage the suppress operation
	if a_attribute_value_integerFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_attribute_value_integer_.Unstage(a_attribute_value_integerFormCallback.probe.stageOfInterest)
	}

	a_attribute_value_integerFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.A_ATTRIBUTE_VALUE_INTEGER](
		a_attribute_value_integerFormCallback.probe,
	)

	// display a new form by reset the form stage
	if a_attribute_value_integerFormCallback.CreationMode || a_attribute_value_integerFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_attribute_value_integerFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(a_attribute_value_integerFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__A_ATTRIBUTE_VALUE_INTEGERFormCallback(
			nil,
			a_attribute_value_integerFormCallback.probe,
			newFormGroup,
		)
		a_attribute_value_integer := new(models.A_ATTRIBUTE_VALUE_INTEGER)
		FillUpForm(a_attribute_value_integer, newFormGroup, a_attribute_value_integerFormCallback.probe)
		a_attribute_value_integerFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(a_attribute_value_integerFormCallback.probe)
}

func __gong__New__A_ATTRIBUTE_VALUE_REALFormCallback(
	a_attribute_value_real *models.A_ATTRIBUTE_VALUE_REAL,
	probe *Probe,
	formGroup *table.FormGroup,
) (a_attribute_value_realFormCallback *A_ATTRIBUTE_VALUE_REALFormCallback) {
	a_attribute_value_realFormCallback = new(A_ATTRIBUTE_VALUE_REALFormCallback)
	a_attribute_value_realFormCallback.probe = probe
	a_attribute_value_realFormCallback.a_attribute_value_real = a_attribute_value_real
	a_attribute_value_realFormCallback.formGroup = formGroup

	a_attribute_value_realFormCallback.CreationMode = (a_attribute_value_real == nil)

	return
}

type A_ATTRIBUTE_VALUE_REALFormCallback struct {
	a_attribute_value_real *models.A_ATTRIBUTE_VALUE_REAL

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (a_attribute_value_realFormCallback *A_ATTRIBUTE_VALUE_REALFormCallback) OnSave() {
	a_attribute_value_realFormCallback.probe.stageOfInterest.Lock()
	defer a_attribute_value_realFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("A_ATTRIBUTE_VALUE_REALFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	a_attribute_value_realFormCallback.probe.formStage.Checkout()

	if a_attribute_value_realFormCallback.a_attribute_value_real == nil {
		a_attribute_value_realFormCallback.a_attribute_value_real = new(models.A_ATTRIBUTE_VALUE_REAL).Stage(a_attribute_value_realFormCallback.probe.stageOfInterest)
	}
	a_attribute_value_real_ := a_attribute_value_realFormCallback.a_attribute_value_real
	_ = a_attribute_value_real_

	for _, formDiv := range a_attribute_value_realFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(a_attribute_value_real_.Name), formDiv)
		case "ATTRIBUTE_VALUE_REAL":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.ATTRIBUTE_VALUE_REAL](a_attribute_value_realFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.ATTRIBUTE_VALUE_REAL, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.ATTRIBUTE_VALUE_REAL)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					a_attribute_value_realFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.ATTRIBUTE_VALUE_REAL](a_attribute_value_realFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			a_attribute_value_real_.ATTRIBUTE_VALUE_REAL = instanceSlice

		}
	}

	// manage the suppress operation
	if a_attribute_value_realFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_attribute_value_real_.Unstage(a_attribute_value_realFormCallback.probe.stageOfInterest)
	}

	a_attribute_value_realFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.A_ATTRIBUTE_VALUE_REAL](
		a_attribute_value_realFormCallback.probe,
	)

	// display a new form by reset the form stage
	if a_attribute_value_realFormCallback.CreationMode || a_attribute_value_realFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_attribute_value_realFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(a_attribute_value_realFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__A_ATTRIBUTE_VALUE_REALFormCallback(
			nil,
			a_attribute_value_realFormCallback.probe,
			newFormGroup,
		)
		a_attribute_value_real := new(models.A_ATTRIBUTE_VALUE_REAL)
		FillUpForm(a_attribute_value_real, newFormGroup, a_attribute_value_realFormCallback.probe)
		a_attribute_value_realFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(a_attribute_value_realFormCallback.probe)
}

func __gong__New__A_ATTRIBUTE_VALUE_STRINGFormCallback(
	a_attribute_value_string *models.A_ATTRIBUTE_VALUE_STRING,
	probe *Probe,
	formGroup *table.FormGroup,
) (a_attribute_value_stringFormCallback *A_ATTRIBUTE_VALUE_STRINGFormCallback) {
	a_attribute_value_stringFormCallback = new(A_ATTRIBUTE_VALUE_STRINGFormCallback)
	a_attribute_value_stringFormCallback.probe = probe
	a_attribute_value_stringFormCallback.a_attribute_value_string = a_attribute_value_string
	a_attribute_value_stringFormCallback.formGroup = formGroup

	a_attribute_value_stringFormCallback.CreationMode = (a_attribute_value_string == nil)

	return
}

type A_ATTRIBUTE_VALUE_STRINGFormCallback struct {
	a_attribute_value_string *models.A_ATTRIBUTE_VALUE_STRING

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (a_attribute_value_stringFormCallback *A_ATTRIBUTE_VALUE_STRINGFormCallback) OnSave() {
	a_attribute_value_stringFormCallback.probe.stageOfInterest.Lock()
	defer a_attribute_value_stringFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("A_ATTRIBUTE_VALUE_STRINGFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	a_attribute_value_stringFormCallback.probe.formStage.Checkout()

	if a_attribute_value_stringFormCallback.a_attribute_value_string == nil {
		a_attribute_value_stringFormCallback.a_attribute_value_string = new(models.A_ATTRIBUTE_VALUE_STRING).Stage(a_attribute_value_stringFormCallback.probe.stageOfInterest)
	}
	a_attribute_value_string_ := a_attribute_value_stringFormCallback.a_attribute_value_string
	_ = a_attribute_value_string_

	for _, formDiv := range a_attribute_value_stringFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(a_attribute_value_string_.Name), formDiv)
		case "ATTRIBUTE_VALUE_STRING":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.ATTRIBUTE_VALUE_STRING](a_attribute_value_stringFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.ATTRIBUTE_VALUE_STRING, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.ATTRIBUTE_VALUE_STRING)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					a_attribute_value_stringFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.ATTRIBUTE_VALUE_STRING](a_attribute_value_stringFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			a_attribute_value_string_.ATTRIBUTE_VALUE_STRING = instanceSlice

		}
	}

	// manage the suppress operation
	if a_attribute_value_stringFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_attribute_value_string_.Unstage(a_attribute_value_stringFormCallback.probe.stageOfInterest)
	}

	a_attribute_value_stringFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.A_ATTRIBUTE_VALUE_STRING](
		a_attribute_value_stringFormCallback.probe,
	)

	// display a new form by reset the form stage
	if a_attribute_value_stringFormCallback.CreationMode || a_attribute_value_stringFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_attribute_value_stringFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(a_attribute_value_stringFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__A_ATTRIBUTE_VALUE_STRINGFormCallback(
			nil,
			a_attribute_value_stringFormCallback.probe,
			newFormGroup,
		)
		a_attribute_value_string := new(models.A_ATTRIBUTE_VALUE_STRING)
		FillUpForm(a_attribute_value_string, newFormGroup, a_attribute_value_stringFormCallback.probe)
		a_attribute_value_stringFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(a_attribute_value_stringFormCallback.probe)
}

func __gong__New__A_ATTRIBUTE_VALUE_XHTMLFormCallback(
	a_attribute_value_xhtml *models.A_ATTRIBUTE_VALUE_XHTML,
	probe *Probe,
	formGroup *table.FormGroup,
) (a_attribute_value_xhtmlFormCallback *A_ATTRIBUTE_VALUE_XHTMLFormCallback) {
	a_attribute_value_xhtmlFormCallback = new(A_ATTRIBUTE_VALUE_XHTMLFormCallback)
	a_attribute_value_xhtmlFormCallback.probe = probe
	a_attribute_value_xhtmlFormCallback.a_attribute_value_xhtml = a_attribute_value_xhtml
	a_attribute_value_xhtmlFormCallback.formGroup = formGroup

	a_attribute_value_xhtmlFormCallback.CreationMode = (a_attribute_value_xhtml == nil)

	return
}

type A_ATTRIBUTE_VALUE_XHTMLFormCallback struct {
	a_attribute_value_xhtml *models.A_ATTRIBUTE_VALUE_XHTML

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (a_attribute_value_xhtmlFormCallback *A_ATTRIBUTE_VALUE_XHTMLFormCallback) OnSave() {
	a_attribute_value_xhtmlFormCallback.probe.stageOfInterest.Lock()
	defer a_attribute_value_xhtmlFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("A_ATTRIBUTE_VALUE_XHTMLFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	a_attribute_value_xhtmlFormCallback.probe.formStage.Checkout()

	if a_attribute_value_xhtmlFormCallback.a_attribute_value_xhtml == nil {
		a_attribute_value_xhtmlFormCallback.a_attribute_value_xhtml = new(models.A_ATTRIBUTE_VALUE_XHTML).Stage(a_attribute_value_xhtmlFormCallback.probe.stageOfInterest)
	}
	a_attribute_value_xhtml_ := a_attribute_value_xhtmlFormCallback.a_attribute_value_xhtml
	_ = a_attribute_value_xhtml_

	for _, formDiv := range a_attribute_value_xhtmlFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(a_attribute_value_xhtml_.Name), formDiv)
		case "ATTRIBUTE_VALUE_XHTML":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.ATTRIBUTE_VALUE_XHTML](a_attribute_value_xhtmlFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.ATTRIBUTE_VALUE_XHTML, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.ATTRIBUTE_VALUE_XHTML)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					a_attribute_value_xhtmlFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.ATTRIBUTE_VALUE_XHTML](a_attribute_value_xhtmlFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			a_attribute_value_xhtml_.ATTRIBUTE_VALUE_XHTML = instanceSlice

		}
	}

	// manage the suppress operation
	if a_attribute_value_xhtmlFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_attribute_value_xhtml_.Unstage(a_attribute_value_xhtmlFormCallback.probe.stageOfInterest)
	}

	a_attribute_value_xhtmlFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.A_ATTRIBUTE_VALUE_XHTML](
		a_attribute_value_xhtmlFormCallback.probe,
	)

	// display a new form by reset the form stage
	if a_attribute_value_xhtmlFormCallback.CreationMode || a_attribute_value_xhtmlFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_attribute_value_xhtmlFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(a_attribute_value_xhtmlFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__A_ATTRIBUTE_VALUE_XHTMLFormCallback(
			nil,
			a_attribute_value_xhtmlFormCallback.probe,
			newFormGroup,
		)
		a_attribute_value_xhtml := new(models.A_ATTRIBUTE_VALUE_XHTML)
		FillUpForm(a_attribute_value_xhtml, newFormGroup, a_attribute_value_xhtmlFormCallback.probe)
		a_attribute_value_xhtmlFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(a_attribute_value_xhtmlFormCallback.probe)
}

func __gong__New__A_ATTRIBUTE_VALUE_XHTML_1FormCallback(
	a_attribute_value_xhtml_1 *models.A_ATTRIBUTE_VALUE_XHTML_1,
	probe *Probe,
	formGroup *table.FormGroup,
) (a_attribute_value_xhtml_1FormCallback *A_ATTRIBUTE_VALUE_XHTML_1FormCallback) {
	a_attribute_value_xhtml_1FormCallback = new(A_ATTRIBUTE_VALUE_XHTML_1FormCallback)
	a_attribute_value_xhtml_1FormCallback.probe = probe
	a_attribute_value_xhtml_1FormCallback.a_attribute_value_xhtml_1 = a_attribute_value_xhtml_1
	a_attribute_value_xhtml_1FormCallback.formGroup = formGroup

	a_attribute_value_xhtml_1FormCallback.CreationMode = (a_attribute_value_xhtml_1 == nil)

	return
}

type A_ATTRIBUTE_VALUE_XHTML_1FormCallback struct {
	a_attribute_value_xhtml_1 *models.A_ATTRIBUTE_VALUE_XHTML_1

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (a_attribute_value_xhtml_1FormCallback *A_ATTRIBUTE_VALUE_XHTML_1FormCallback) OnSave() {
	a_attribute_value_xhtml_1FormCallback.probe.stageOfInterest.Lock()
	defer a_attribute_value_xhtml_1FormCallback.probe.stageOfInterest.Unlock()

	// log.Println("A_ATTRIBUTE_VALUE_XHTML_1FormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	a_attribute_value_xhtml_1FormCallback.probe.formStage.Checkout()

	if a_attribute_value_xhtml_1FormCallback.a_attribute_value_xhtml_1 == nil {
		a_attribute_value_xhtml_1FormCallback.a_attribute_value_xhtml_1 = new(models.A_ATTRIBUTE_VALUE_XHTML_1).Stage(a_attribute_value_xhtml_1FormCallback.probe.stageOfInterest)
	}
	a_attribute_value_xhtml_1_ := a_attribute_value_xhtml_1FormCallback.a_attribute_value_xhtml_1
	_ = a_attribute_value_xhtml_1_

	for _, formDiv := range a_attribute_value_xhtml_1FormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(a_attribute_value_xhtml_1_.Name), formDiv)
		case "ATTRIBUTE_VALUE_BOOLEAN":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.ATTRIBUTE_VALUE_BOOLEAN](a_attribute_value_xhtml_1FormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.ATTRIBUTE_VALUE_BOOLEAN, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.ATTRIBUTE_VALUE_BOOLEAN)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					a_attribute_value_xhtml_1FormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.ATTRIBUTE_VALUE_BOOLEAN](a_attribute_value_xhtml_1FormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			a_attribute_value_xhtml_1_.ATTRIBUTE_VALUE_BOOLEAN = instanceSlice

		case "ATTRIBUTE_VALUE_DATE":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.ATTRIBUTE_VALUE_DATE](a_attribute_value_xhtml_1FormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.ATTRIBUTE_VALUE_DATE, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.ATTRIBUTE_VALUE_DATE)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					a_attribute_value_xhtml_1FormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.ATTRIBUTE_VALUE_DATE](a_attribute_value_xhtml_1FormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			a_attribute_value_xhtml_1_.ATTRIBUTE_VALUE_DATE = instanceSlice

		case "ATTRIBUTE_VALUE_ENUMERATION":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.ATTRIBUTE_VALUE_ENUMERATION](a_attribute_value_xhtml_1FormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.ATTRIBUTE_VALUE_ENUMERATION, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.ATTRIBUTE_VALUE_ENUMERATION)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					a_attribute_value_xhtml_1FormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.ATTRIBUTE_VALUE_ENUMERATION](a_attribute_value_xhtml_1FormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			a_attribute_value_xhtml_1_.ATTRIBUTE_VALUE_ENUMERATION = instanceSlice

		case "ATTRIBUTE_VALUE_INTEGER":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.ATTRIBUTE_VALUE_INTEGER](a_attribute_value_xhtml_1FormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.ATTRIBUTE_VALUE_INTEGER, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.ATTRIBUTE_VALUE_INTEGER)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					a_attribute_value_xhtml_1FormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.ATTRIBUTE_VALUE_INTEGER](a_attribute_value_xhtml_1FormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			a_attribute_value_xhtml_1_.ATTRIBUTE_VALUE_INTEGER = instanceSlice

		case "ATTRIBUTE_VALUE_REAL":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.ATTRIBUTE_VALUE_REAL](a_attribute_value_xhtml_1FormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.ATTRIBUTE_VALUE_REAL, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.ATTRIBUTE_VALUE_REAL)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					a_attribute_value_xhtml_1FormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.ATTRIBUTE_VALUE_REAL](a_attribute_value_xhtml_1FormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			a_attribute_value_xhtml_1_.ATTRIBUTE_VALUE_REAL = instanceSlice

		case "ATTRIBUTE_VALUE_STRING":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.ATTRIBUTE_VALUE_STRING](a_attribute_value_xhtml_1FormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.ATTRIBUTE_VALUE_STRING, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.ATTRIBUTE_VALUE_STRING)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					a_attribute_value_xhtml_1FormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.ATTRIBUTE_VALUE_STRING](a_attribute_value_xhtml_1FormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			a_attribute_value_xhtml_1_.ATTRIBUTE_VALUE_STRING = instanceSlice

		case "ATTRIBUTE_VALUE_XHTML":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.ATTRIBUTE_VALUE_XHTML](a_attribute_value_xhtml_1FormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.ATTRIBUTE_VALUE_XHTML, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.ATTRIBUTE_VALUE_XHTML)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					a_attribute_value_xhtml_1FormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.ATTRIBUTE_VALUE_XHTML](a_attribute_value_xhtml_1FormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			a_attribute_value_xhtml_1_.ATTRIBUTE_VALUE_XHTML = instanceSlice

		}
	}

	// manage the suppress operation
	if a_attribute_value_xhtml_1FormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_attribute_value_xhtml_1_.Unstage(a_attribute_value_xhtml_1FormCallback.probe.stageOfInterest)
	}

	a_attribute_value_xhtml_1FormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.A_ATTRIBUTE_VALUE_XHTML_1](
		a_attribute_value_xhtml_1FormCallback.probe,
	)

	// display a new form by reset the form stage
	if a_attribute_value_xhtml_1FormCallback.CreationMode || a_attribute_value_xhtml_1FormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_attribute_value_xhtml_1FormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(a_attribute_value_xhtml_1FormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__A_ATTRIBUTE_VALUE_XHTML_1FormCallback(
			nil,
			a_attribute_value_xhtml_1FormCallback.probe,
			newFormGroup,
		)
		a_attribute_value_xhtml_1 := new(models.A_ATTRIBUTE_VALUE_XHTML_1)
		FillUpForm(a_attribute_value_xhtml_1, newFormGroup, a_attribute_value_xhtml_1FormCallback.probe)
		a_attribute_value_xhtml_1FormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(a_attribute_value_xhtml_1FormCallback.probe)
}

func __gong__New__A_CHILDRENFormCallback(
	a_children *models.A_CHILDREN,
	probe *Probe,
	formGroup *table.FormGroup,
) (a_childrenFormCallback *A_CHILDRENFormCallback) {
	a_childrenFormCallback = new(A_CHILDRENFormCallback)
	a_childrenFormCallback.probe = probe
	a_childrenFormCallback.a_children = a_children
	a_childrenFormCallback.formGroup = formGroup

	a_childrenFormCallback.CreationMode = (a_children == nil)

	return
}

type A_CHILDRENFormCallback struct {
	a_children *models.A_CHILDREN

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (a_childrenFormCallback *A_CHILDRENFormCallback) OnSave() {
	a_childrenFormCallback.probe.stageOfInterest.Lock()
	defer a_childrenFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("A_CHILDRENFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	a_childrenFormCallback.probe.formStage.Checkout()

	if a_childrenFormCallback.a_children == nil {
		a_childrenFormCallback.a_children = new(models.A_CHILDREN).Stage(a_childrenFormCallback.probe.stageOfInterest)
	}
	a_children_ := a_childrenFormCallback.a_children
	_ = a_children_

	for _, formDiv := range a_childrenFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(a_children_.Name), formDiv)
		case "SPEC_HIERARCHY":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.SPEC_HIERARCHY](a_childrenFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.SPEC_HIERARCHY, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.SPEC_HIERARCHY)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					a_childrenFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.SPEC_HIERARCHY](a_childrenFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			a_children_.SPEC_HIERARCHY = instanceSlice

		}
	}

	// manage the suppress operation
	if a_childrenFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_children_.Unstage(a_childrenFormCallback.probe.stageOfInterest)
	}

	a_childrenFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.A_CHILDREN](
		a_childrenFormCallback.probe,
	)

	// display a new form by reset the form stage
	if a_childrenFormCallback.CreationMode || a_childrenFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_childrenFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(a_childrenFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__A_CHILDRENFormCallback(
			nil,
			a_childrenFormCallback.probe,
			newFormGroup,
		)
		a_children := new(models.A_CHILDREN)
		FillUpForm(a_children, newFormGroup, a_childrenFormCallback.probe)
		a_childrenFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(a_childrenFormCallback.probe)
}

func __gong__New__A_CORE_CONTENTFormCallback(
	a_core_content *models.A_CORE_CONTENT,
	probe *Probe,
	formGroup *table.FormGroup,
) (a_core_contentFormCallback *A_CORE_CONTENTFormCallback) {
	a_core_contentFormCallback = new(A_CORE_CONTENTFormCallback)
	a_core_contentFormCallback.probe = probe
	a_core_contentFormCallback.a_core_content = a_core_content
	a_core_contentFormCallback.formGroup = formGroup

	a_core_contentFormCallback.CreationMode = (a_core_content == nil)

	return
}

type A_CORE_CONTENTFormCallback struct {
	a_core_content *models.A_CORE_CONTENT

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (a_core_contentFormCallback *A_CORE_CONTENTFormCallback) OnSave() {
	a_core_contentFormCallback.probe.stageOfInterest.Lock()
	defer a_core_contentFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("A_CORE_CONTENTFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	a_core_contentFormCallback.probe.formStage.Checkout()

	if a_core_contentFormCallback.a_core_content == nil {
		a_core_contentFormCallback.a_core_content = new(models.A_CORE_CONTENT).Stage(a_core_contentFormCallback.probe.stageOfInterest)
	}
	a_core_content_ := a_core_contentFormCallback.a_core_content
	_ = a_core_content_

	for _, formDiv := range a_core_contentFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(a_core_content_.Name), formDiv)
		case "REQ_IF_CONTENT":
			FormDivSelectFieldToField(&(a_core_content_.REQ_IF_CONTENT), a_core_contentFormCallback.probe.stageOfInterest, formDiv)
		}
	}

	// manage the suppress operation
	if a_core_contentFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_core_content_.Unstage(a_core_contentFormCallback.probe.stageOfInterest)
	}

	a_core_contentFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.A_CORE_CONTENT](
		a_core_contentFormCallback.probe,
	)

	// display a new form by reset the form stage
	if a_core_contentFormCallback.CreationMode || a_core_contentFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_core_contentFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(a_core_contentFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__A_CORE_CONTENTFormCallback(
			nil,
			a_core_contentFormCallback.probe,
			newFormGroup,
		)
		a_core_content := new(models.A_CORE_CONTENT)
		FillUpForm(a_core_content, newFormGroup, a_core_contentFormCallback.probe)
		a_core_contentFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(a_core_contentFormCallback.probe)
}

func __gong__New__A_DATATYPESFormCallback(
	a_datatypes *models.A_DATATYPES,
	probe *Probe,
	formGroup *table.FormGroup,
) (a_datatypesFormCallback *A_DATATYPESFormCallback) {
	a_datatypesFormCallback = new(A_DATATYPESFormCallback)
	a_datatypesFormCallback.probe = probe
	a_datatypesFormCallback.a_datatypes = a_datatypes
	a_datatypesFormCallback.formGroup = formGroup

	a_datatypesFormCallback.CreationMode = (a_datatypes == nil)

	return
}

type A_DATATYPESFormCallback struct {
	a_datatypes *models.A_DATATYPES

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (a_datatypesFormCallback *A_DATATYPESFormCallback) OnSave() {
	a_datatypesFormCallback.probe.stageOfInterest.Lock()
	defer a_datatypesFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("A_DATATYPESFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	a_datatypesFormCallback.probe.formStage.Checkout()

	if a_datatypesFormCallback.a_datatypes == nil {
		a_datatypesFormCallback.a_datatypes = new(models.A_DATATYPES).Stage(a_datatypesFormCallback.probe.stageOfInterest)
	}
	a_datatypes_ := a_datatypesFormCallback.a_datatypes
	_ = a_datatypes_

	for _, formDiv := range a_datatypesFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(a_datatypes_.Name), formDiv)
		case "DATATYPE_DEFINITION_BOOLEAN":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.DATATYPE_DEFINITION_BOOLEAN](a_datatypesFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.DATATYPE_DEFINITION_BOOLEAN, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.DATATYPE_DEFINITION_BOOLEAN)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					a_datatypesFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.DATATYPE_DEFINITION_BOOLEAN](a_datatypesFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			a_datatypes_.DATATYPE_DEFINITION_BOOLEAN = instanceSlice

		case "DATATYPE_DEFINITION_DATE":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.DATATYPE_DEFINITION_DATE](a_datatypesFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.DATATYPE_DEFINITION_DATE, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.DATATYPE_DEFINITION_DATE)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					a_datatypesFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.DATATYPE_DEFINITION_DATE](a_datatypesFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			a_datatypes_.DATATYPE_DEFINITION_DATE = instanceSlice

		case "DATATYPE_DEFINITION_ENUMERATION":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.DATATYPE_DEFINITION_ENUMERATION](a_datatypesFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.DATATYPE_DEFINITION_ENUMERATION, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.DATATYPE_DEFINITION_ENUMERATION)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					a_datatypesFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.DATATYPE_DEFINITION_ENUMERATION](a_datatypesFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			a_datatypes_.DATATYPE_DEFINITION_ENUMERATION = instanceSlice

		case "DATATYPE_DEFINITION_INTEGER":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.DATATYPE_DEFINITION_INTEGER](a_datatypesFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.DATATYPE_DEFINITION_INTEGER, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.DATATYPE_DEFINITION_INTEGER)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					a_datatypesFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.DATATYPE_DEFINITION_INTEGER](a_datatypesFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			a_datatypes_.DATATYPE_DEFINITION_INTEGER = instanceSlice

		case "DATATYPE_DEFINITION_REAL":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.DATATYPE_DEFINITION_REAL](a_datatypesFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.DATATYPE_DEFINITION_REAL, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.DATATYPE_DEFINITION_REAL)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					a_datatypesFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.DATATYPE_DEFINITION_REAL](a_datatypesFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			a_datatypes_.DATATYPE_DEFINITION_REAL = instanceSlice

		case "DATATYPE_DEFINITION_STRING":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.DATATYPE_DEFINITION_STRING](a_datatypesFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.DATATYPE_DEFINITION_STRING, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.DATATYPE_DEFINITION_STRING)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					a_datatypesFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.DATATYPE_DEFINITION_STRING](a_datatypesFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			a_datatypes_.DATATYPE_DEFINITION_STRING = instanceSlice

		case "DATATYPE_DEFINITION_XHTML":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.DATATYPE_DEFINITION_XHTML](a_datatypesFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.DATATYPE_DEFINITION_XHTML, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.DATATYPE_DEFINITION_XHTML)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					a_datatypesFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.DATATYPE_DEFINITION_XHTML](a_datatypesFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			a_datatypes_.DATATYPE_DEFINITION_XHTML = instanceSlice

		}
	}

	// manage the suppress operation
	if a_datatypesFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_datatypes_.Unstage(a_datatypesFormCallback.probe.stageOfInterest)
	}

	a_datatypesFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.A_DATATYPES](
		a_datatypesFormCallback.probe,
	)

	// display a new form by reset the form stage
	if a_datatypesFormCallback.CreationMode || a_datatypesFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_datatypesFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(a_datatypesFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__A_DATATYPESFormCallback(
			nil,
			a_datatypesFormCallback.probe,
			newFormGroup,
		)
		a_datatypes := new(models.A_DATATYPES)
		FillUpForm(a_datatypes, newFormGroup, a_datatypesFormCallback.probe)
		a_datatypesFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(a_datatypesFormCallback.probe)
}

func __gong__New__A_DATATYPE_DEFINITION_BOOLEAN_REFFormCallback(
	a_datatype_definition_boolean_ref *models.A_DATATYPE_DEFINITION_BOOLEAN_REF,
	probe *Probe,
	formGroup *table.FormGroup,
) (a_datatype_definition_boolean_refFormCallback *A_DATATYPE_DEFINITION_BOOLEAN_REFFormCallback) {
	a_datatype_definition_boolean_refFormCallback = new(A_DATATYPE_DEFINITION_BOOLEAN_REFFormCallback)
	a_datatype_definition_boolean_refFormCallback.probe = probe
	a_datatype_definition_boolean_refFormCallback.a_datatype_definition_boolean_ref = a_datatype_definition_boolean_ref
	a_datatype_definition_boolean_refFormCallback.formGroup = formGroup

	a_datatype_definition_boolean_refFormCallback.CreationMode = (a_datatype_definition_boolean_ref == nil)

	return
}

type A_DATATYPE_DEFINITION_BOOLEAN_REFFormCallback struct {
	a_datatype_definition_boolean_ref *models.A_DATATYPE_DEFINITION_BOOLEAN_REF

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (a_datatype_definition_boolean_refFormCallback *A_DATATYPE_DEFINITION_BOOLEAN_REFFormCallback) OnSave() {
	a_datatype_definition_boolean_refFormCallback.probe.stageOfInterest.Lock()
	defer a_datatype_definition_boolean_refFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("A_DATATYPE_DEFINITION_BOOLEAN_REFFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	a_datatype_definition_boolean_refFormCallback.probe.formStage.Checkout()

	if a_datatype_definition_boolean_refFormCallback.a_datatype_definition_boolean_ref == nil {
		a_datatype_definition_boolean_refFormCallback.a_datatype_definition_boolean_ref = new(models.A_DATATYPE_DEFINITION_BOOLEAN_REF).Stage(a_datatype_definition_boolean_refFormCallback.probe.stageOfInterest)
	}
	a_datatype_definition_boolean_ref_ := a_datatype_definition_boolean_refFormCallback.a_datatype_definition_boolean_ref
	_ = a_datatype_definition_boolean_ref_

	for _, formDiv := range a_datatype_definition_boolean_refFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(a_datatype_definition_boolean_ref_.Name), formDiv)
		case "DATATYPE_DEFINITION_BOOLEAN_REF":
			FormDivBasicFieldToField(&(a_datatype_definition_boolean_ref_.DATATYPE_DEFINITION_BOOLEAN_REF), formDiv)
		}
	}

	// manage the suppress operation
	if a_datatype_definition_boolean_refFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_datatype_definition_boolean_ref_.Unstage(a_datatype_definition_boolean_refFormCallback.probe.stageOfInterest)
	}

	a_datatype_definition_boolean_refFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.A_DATATYPE_DEFINITION_BOOLEAN_REF](
		a_datatype_definition_boolean_refFormCallback.probe,
	)

	// display a new form by reset the form stage
	if a_datatype_definition_boolean_refFormCallback.CreationMode || a_datatype_definition_boolean_refFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_datatype_definition_boolean_refFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(a_datatype_definition_boolean_refFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__A_DATATYPE_DEFINITION_BOOLEAN_REFFormCallback(
			nil,
			a_datatype_definition_boolean_refFormCallback.probe,
			newFormGroup,
		)
		a_datatype_definition_boolean_ref := new(models.A_DATATYPE_DEFINITION_BOOLEAN_REF)
		FillUpForm(a_datatype_definition_boolean_ref, newFormGroup, a_datatype_definition_boolean_refFormCallback.probe)
		a_datatype_definition_boolean_refFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(a_datatype_definition_boolean_refFormCallback.probe)
}

func __gong__New__A_DATATYPE_DEFINITION_DATE_REFFormCallback(
	a_datatype_definition_date_ref *models.A_DATATYPE_DEFINITION_DATE_REF,
	probe *Probe,
	formGroup *table.FormGroup,
) (a_datatype_definition_date_refFormCallback *A_DATATYPE_DEFINITION_DATE_REFFormCallback) {
	a_datatype_definition_date_refFormCallback = new(A_DATATYPE_DEFINITION_DATE_REFFormCallback)
	a_datatype_definition_date_refFormCallback.probe = probe
	a_datatype_definition_date_refFormCallback.a_datatype_definition_date_ref = a_datatype_definition_date_ref
	a_datatype_definition_date_refFormCallback.formGroup = formGroup

	a_datatype_definition_date_refFormCallback.CreationMode = (a_datatype_definition_date_ref == nil)

	return
}

type A_DATATYPE_DEFINITION_DATE_REFFormCallback struct {
	a_datatype_definition_date_ref *models.A_DATATYPE_DEFINITION_DATE_REF

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (a_datatype_definition_date_refFormCallback *A_DATATYPE_DEFINITION_DATE_REFFormCallback) OnSave() {
	a_datatype_definition_date_refFormCallback.probe.stageOfInterest.Lock()
	defer a_datatype_definition_date_refFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("A_DATATYPE_DEFINITION_DATE_REFFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	a_datatype_definition_date_refFormCallback.probe.formStage.Checkout()

	if a_datatype_definition_date_refFormCallback.a_datatype_definition_date_ref == nil {
		a_datatype_definition_date_refFormCallback.a_datatype_definition_date_ref = new(models.A_DATATYPE_DEFINITION_DATE_REF).Stage(a_datatype_definition_date_refFormCallback.probe.stageOfInterest)
	}
	a_datatype_definition_date_ref_ := a_datatype_definition_date_refFormCallback.a_datatype_definition_date_ref
	_ = a_datatype_definition_date_ref_

	for _, formDiv := range a_datatype_definition_date_refFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(a_datatype_definition_date_ref_.Name), formDiv)
		case "DATATYPE_DEFINITION_DATE_REF":
			FormDivBasicFieldToField(&(a_datatype_definition_date_ref_.DATATYPE_DEFINITION_DATE_REF), formDiv)
		}
	}

	// manage the suppress operation
	if a_datatype_definition_date_refFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_datatype_definition_date_ref_.Unstage(a_datatype_definition_date_refFormCallback.probe.stageOfInterest)
	}

	a_datatype_definition_date_refFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.A_DATATYPE_DEFINITION_DATE_REF](
		a_datatype_definition_date_refFormCallback.probe,
	)

	// display a new form by reset the form stage
	if a_datatype_definition_date_refFormCallback.CreationMode || a_datatype_definition_date_refFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_datatype_definition_date_refFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(a_datatype_definition_date_refFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__A_DATATYPE_DEFINITION_DATE_REFFormCallback(
			nil,
			a_datatype_definition_date_refFormCallback.probe,
			newFormGroup,
		)
		a_datatype_definition_date_ref := new(models.A_DATATYPE_DEFINITION_DATE_REF)
		FillUpForm(a_datatype_definition_date_ref, newFormGroup, a_datatype_definition_date_refFormCallback.probe)
		a_datatype_definition_date_refFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(a_datatype_definition_date_refFormCallback.probe)
}

func __gong__New__A_DATATYPE_DEFINITION_ENUMERATION_REFFormCallback(
	a_datatype_definition_enumeration_ref *models.A_DATATYPE_DEFINITION_ENUMERATION_REF,
	probe *Probe,
	formGroup *table.FormGroup,
) (a_datatype_definition_enumeration_refFormCallback *A_DATATYPE_DEFINITION_ENUMERATION_REFFormCallback) {
	a_datatype_definition_enumeration_refFormCallback = new(A_DATATYPE_DEFINITION_ENUMERATION_REFFormCallback)
	a_datatype_definition_enumeration_refFormCallback.probe = probe
	a_datatype_definition_enumeration_refFormCallback.a_datatype_definition_enumeration_ref = a_datatype_definition_enumeration_ref
	a_datatype_definition_enumeration_refFormCallback.formGroup = formGroup

	a_datatype_definition_enumeration_refFormCallback.CreationMode = (a_datatype_definition_enumeration_ref == nil)

	return
}

type A_DATATYPE_DEFINITION_ENUMERATION_REFFormCallback struct {
	a_datatype_definition_enumeration_ref *models.A_DATATYPE_DEFINITION_ENUMERATION_REF

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (a_datatype_definition_enumeration_refFormCallback *A_DATATYPE_DEFINITION_ENUMERATION_REFFormCallback) OnSave() {
	a_datatype_definition_enumeration_refFormCallback.probe.stageOfInterest.Lock()
	defer a_datatype_definition_enumeration_refFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("A_DATATYPE_DEFINITION_ENUMERATION_REFFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	a_datatype_definition_enumeration_refFormCallback.probe.formStage.Checkout()

	if a_datatype_definition_enumeration_refFormCallback.a_datatype_definition_enumeration_ref == nil {
		a_datatype_definition_enumeration_refFormCallback.a_datatype_definition_enumeration_ref = new(models.A_DATATYPE_DEFINITION_ENUMERATION_REF).Stage(a_datatype_definition_enumeration_refFormCallback.probe.stageOfInterest)
	}
	a_datatype_definition_enumeration_ref_ := a_datatype_definition_enumeration_refFormCallback.a_datatype_definition_enumeration_ref
	_ = a_datatype_definition_enumeration_ref_

	for _, formDiv := range a_datatype_definition_enumeration_refFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(a_datatype_definition_enumeration_ref_.Name), formDiv)
		case "DATATYPE_DEFINITION_ENUMERATION_REF":
			FormDivBasicFieldToField(&(a_datatype_definition_enumeration_ref_.DATATYPE_DEFINITION_ENUMERATION_REF), formDiv)
		}
	}

	// manage the suppress operation
	if a_datatype_definition_enumeration_refFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_datatype_definition_enumeration_ref_.Unstage(a_datatype_definition_enumeration_refFormCallback.probe.stageOfInterest)
	}

	a_datatype_definition_enumeration_refFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.A_DATATYPE_DEFINITION_ENUMERATION_REF](
		a_datatype_definition_enumeration_refFormCallback.probe,
	)

	// display a new form by reset the form stage
	if a_datatype_definition_enumeration_refFormCallback.CreationMode || a_datatype_definition_enumeration_refFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_datatype_definition_enumeration_refFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(a_datatype_definition_enumeration_refFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__A_DATATYPE_DEFINITION_ENUMERATION_REFFormCallback(
			nil,
			a_datatype_definition_enumeration_refFormCallback.probe,
			newFormGroup,
		)
		a_datatype_definition_enumeration_ref := new(models.A_DATATYPE_DEFINITION_ENUMERATION_REF)
		FillUpForm(a_datatype_definition_enumeration_ref, newFormGroup, a_datatype_definition_enumeration_refFormCallback.probe)
		a_datatype_definition_enumeration_refFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(a_datatype_definition_enumeration_refFormCallback.probe)
}

func __gong__New__A_DATATYPE_DEFINITION_INTEGER_REFFormCallback(
	a_datatype_definition_integer_ref *models.A_DATATYPE_DEFINITION_INTEGER_REF,
	probe *Probe,
	formGroup *table.FormGroup,
) (a_datatype_definition_integer_refFormCallback *A_DATATYPE_DEFINITION_INTEGER_REFFormCallback) {
	a_datatype_definition_integer_refFormCallback = new(A_DATATYPE_DEFINITION_INTEGER_REFFormCallback)
	a_datatype_definition_integer_refFormCallback.probe = probe
	a_datatype_definition_integer_refFormCallback.a_datatype_definition_integer_ref = a_datatype_definition_integer_ref
	a_datatype_definition_integer_refFormCallback.formGroup = formGroup

	a_datatype_definition_integer_refFormCallback.CreationMode = (a_datatype_definition_integer_ref == nil)

	return
}

type A_DATATYPE_DEFINITION_INTEGER_REFFormCallback struct {
	a_datatype_definition_integer_ref *models.A_DATATYPE_DEFINITION_INTEGER_REF

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (a_datatype_definition_integer_refFormCallback *A_DATATYPE_DEFINITION_INTEGER_REFFormCallback) OnSave() {
	a_datatype_definition_integer_refFormCallback.probe.stageOfInterest.Lock()
	defer a_datatype_definition_integer_refFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("A_DATATYPE_DEFINITION_INTEGER_REFFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	a_datatype_definition_integer_refFormCallback.probe.formStage.Checkout()

	if a_datatype_definition_integer_refFormCallback.a_datatype_definition_integer_ref == nil {
		a_datatype_definition_integer_refFormCallback.a_datatype_definition_integer_ref = new(models.A_DATATYPE_DEFINITION_INTEGER_REF).Stage(a_datatype_definition_integer_refFormCallback.probe.stageOfInterest)
	}
	a_datatype_definition_integer_ref_ := a_datatype_definition_integer_refFormCallback.a_datatype_definition_integer_ref
	_ = a_datatype_definition_integer_ref_

	for _, formDiv := range a_datatype_definition_integer_refFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(a_datatype_definition_integer_ref_.Name), formDiv)
		case "DATATYPE_DEFINITION_INTEGER_REF":
			FormDivBasicFieldToField(&(a_datatype_definition_integer_ref_.DATATYPE_DEFINITION_INTEGER_REF), formDiv)
		}
	}

	// manage the suppress operation
	if a_datatype_definition_integer_refFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_datatype_definition_integer_ref_.Unstage(a_datatype_definition_integer_refFormCallback.probe.stageOfInterest)
	}

	a_datatype_definition_integer_refFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.A_DATATYPE_DEFINITION_INTEGER_REF](
		a_datatype_definition_integer_refFormCallback.probe,
	)

	// display a new form by reset the form stage
	if a_datatype_definition_integer_refFormCallback.CreationMode || a_datatype_definition_integer_refFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_datatype_definition_integer_refFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(a_datatype_definition_integer_refFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__A_DATATYPE_DEFINITION_INTEGER_REFFormCallback(
			nil,
			a_datatype_definition_integer_refFormCallback.probe,
			newFormGroup,
		)
		a_datatype_definition_integer_ref := new(models.A_DATATYPE_DEFINITION_INTEGER_REF)
		FillUpForm(a_datatype_definition_integer_ref, newFormGroup, a_datatype_definition_integer_refFormCallback.probe)
		a_datatype_definition_integer_refFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(a_datatype_definition_integer_refFormCallback.probe)
}

func __gong__New__A_DATATYPE_DEFINITION_REAL_REFFormCallback(
	a_datatype_definition_real_ref *models.A_DATATYPE_DEFINITION_REAL_REF,
	probe *Probe,
	formGroup *table.FormGroup,
) (a_datatype_definition_real_refFormCallback *A_DATATYPE_DEFINITION_REAL_REFFormCallback) {
	a_datatype_definition_real_refFormCallback = new(A_DATATYPE_DEFINITION_REAL_REFFormCallback)
	a_datatype_definition_real_refFormCallback.probe = probe
	a_datatype_definition_real_refFormCallback.a_datatype_definition_real_ref = a_datatype_definition_real_ref
	a_datatype_definition_real_refFormCallback.formGroup = formGroup

	a_datatype_definition_real_refFormCallback.CreationMode = (a_datatype_definition_real_ref == nil)

	return
}

type A_DATATYPE_DEFINITION_REAL_REFFormCallback struct {
	a_datatype_definition_real_ref *models.A_DATATYPE_DEFINITION_REAL_REF

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (a_datatype_definition_real_refFormCallback *A_DATATYPE_DEFINITION_REAL_REFFormCallback) OnSave() {
	a_datatype_definition_real_refFormCallback.probe.stageOfInterest.Lock()
	defer a_datatype_definition_real_refFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("A_DATATYPE_DEFINITION_REAL_REFFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	a_datatype_definition_real_refFormCallback.probe.formStage.Checkout()

	if a_datatype_definition_real_refFormCallback.a_datatype_definition_real_ref == nil {
		a_datatype_definition_real_refFormCallback.a_datatype_definition_real_ref = new(models.A_DATATYPE_DEFINITION_REAL_REF).Stage(a_datatype_definition_real_refFormCallback.probe.stageOfInterest)
	}
	a_datatype_definition_real_ref_ := a_datatype_definition_real_refFormCallback.a_datatype_definition_real_ref
	_ = a_datatype_definition_real_ref_

	for _, formDiv := range a_datatype_definition_real_refFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(a_datatype_definition_real_ref_.Name), formDiv)
		case "DATATYPE_DEFINITION_REAL_REF":
			FormDivBasicFieldToField(&(a_datatype_definition_real_ref_.DATATYPE_DEFINITION_REAL_REF), formDiv)
		}
	}

	// manage the suppress operation
	if a_datatype_definition_real_refFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_datatype_definition_real_ref_.Unstage(a_datatype_definition_real_refFormCallback.probe.stageOfInterest)
	}

	a_datatype_definition_real_refFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.A_DATATYPE_DEFINITION_REAL_REF](
		a_datatype_definition_real_refFormCallback.probe,
	)

	// display a new form by reset the form stage
	if a_datatype_definition_real_refFormCallback.CreationMode || a_datatype_definition_real_refFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_datatype_definition_real_refFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(a_datatype_definition_real_refFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__A_DATATYPE_DEFINITION_REAL_REFFormCallback(
			nil,
			a_datatype_definition_real_refFormCallback.probe,
			newFormGroup,
		)
		a_datatype_definition_real_ref := new(models.A_DATATYPE_DEFINITION_REAL_REF)
		FillUpForm(a_datatype_definition_real_ref, newFormGroup, a_datatype_definition_real_refFormCallback.probe)
		a_datatype_definition_real_refFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(a_datatype_definition_real_refFormCallback.probe)
}

func __gong__New__A_DATATYPE_DEFINITION_STRING_REFFormCallback(
	a_datatype_definition_string_ref *models.A_DATATYPE_DEFINITION_STRING_REF,
	probe *Probe,
	formGroup *table.FormGroup,
) (a_datatype_definition_string_refFormCallback *A_DATATYPE_DEFINITION_STRING_REFFormCallback) {
	a_datatype_definition_string_refFormCallback = new(A_DATATYPE_DEFINITION_STRING_REFFormCallback)
	a_datatype_definition_string_refFormCallback.probe = probe
	a_datatype_definition_string_refFormCallback.a_datatype_definition_string_ref = a_datatype_definition_string_ref
	a_datatype_definition_string_refFormCallback.formGroup = formGroup

	a_datatype_definition_string_refFormCallback.CreationMode = (a_datatype_definition_string_ref == nil)

	return
}

type A_DATATYPE_DEFINITION_STRING_REFFormCallback struct {
	a_datatype_definition_string_ref *models.A_DATATYPE_DEFINITION_STRING_REF

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (a_datatype_definition_string_refFormCallback *A_DATATYPE_DEFINITION_STRING_REFFormCallback) OnSave() {
	a_datatype_definition_string_refFormCallback.probe.stageOfInterest.Lock()
	defer a_datatype_definition_string_refFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("A_DATATYPE_DEFINITION_STRING_REFFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	a_datatype_definition_string_refFormCallback.probe.formStage.Checkout()

	if a_datatype_definition_string_refFormCallback.a_datatype_definition_string_ref == nil {
		a_datatype_definition_string_refFormCallback.a_datatype_definition_string_ref = new(models.A_DATATYPE_DEFINITION_STRING_REF).Stage(a_datatype_definition_string_refFormCallback.probe.stageOfInterest)
	}
	a_datatype_definition_string_ref_ := a_datatype_definition_string_refFormCallback.a_datatype_definition_string_ref
	_ = a_datatype_definition_string_ref_

	for _, formDiv := range a_datatype_definition_string_refFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(a_datatype_definition_string_ref_.Name), formDiv)
		case "DATATYPE_DEFINITION_STRING_REF":
			FormDivBasicFieldToField(&(a_datatype_definition_string_ref_.DATATYPE_DEFINITION_STRING_REF), formDiv)
		}
	}

	// manage the suppress operation
	if a_datatype_definition_string_refFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_datatype_definition_string_ref_.Unstage(a_datatype_definition_string_refFormCallback.probe.stageOfInterest)
	}

	a_datatype_definition_string_refFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.A_DATATYPE_DEFINITION_STRING_REF](
		a_datatype_definition_string_refFormCallback.probe,
	)

	// display a new form by reset the form stage
	if a_datatype_definition_string_refFormCallback.CreationMode || a_datatype_definition_string_refFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_datatype_definition_string_refFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(a_datatype_definition_string_refFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__A_DATATYPE_DEFINITION_STRING_REFFormCallback(
			nil,
			a_datatype_definition_string_refFormCallback.probe,
			newFormGroup,
		)
		a_datatype_definition_string_ref := new(models.A_DATATYPE_DEFINITION_STRING_REF)
		FillUpForm(a_datatype_definition_string_ref, newFormGroup, a_datatype_definition_string_refFormCallback.probe)
		a_datatype_definition_string_refFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(a_datatype_definition_string_refFormCallback.probe)
}

func __gong__New__A_DATATYPE_DEFINITION_XHTML_REFFormCallback(
	a_datatype_definition_xhtml_ref *models.A_DATATYPE_DEFINITION_XHTML_REF,
	probe *Probe,
	formGroup *table.FormGroup,
) (a_datatype_definition_xhtml_refFormCallback *A_DATATYPE_DEFINITION_XHTML_REFFormCallback) {
	a_datatype_definition_xhtml_refFormCallback = new(A_DATATYPE_DEFINITION_XHTML_REFFormCallback)
	a_datatype_definition_xhtml_refFormCallback.probe = probe
	a_datatype_definition_xhtml_refFormCallback.a_datatype_definition_xhtml_ref = a_datatype_definition_xhtml_ref
	a_datatype_definition_xhtml_refFormCallback.formGroup = formGroup

	a_datatype_definition_xhtml_refFormCallback.CreationMode = (a_datatype_definition_xhtml_ref == nil)

	return
}

type A_DATATYPE_DEFINITION_XHTML_REFFormCallback struct {
	a_datatype_definition_xhtml_ref *models.A_DATATYPE_DEFINITION_XHTML_REF

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (a_datatype_definition_xhtml_refFormCallback *A_DATATYPE_DEFINITION_XHTML_REFFormCallback) OnSave() {
	a_datatype_definition_xhtml_refFormCallback.probe.stageOfInterest.Lock()
	defer a_datatype_definition_xhtml_refFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("A_DATATYPE_DEFINITION_XHTML_REFFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	a_datatype_definition_xhtml_refFormCallback.probe.formStage.Checkout()

	if a_datatype_definition_xhtml_refFormCallback.a_datatype_definition_xhtml_ref == nil {
		a_datatype_definition_xhtml_refFormCallback.a_datatype_definition_xhtml_ref = new(models.A_DATATYPE_DEFINITION_XHTML_REF).Stage(a_datatype_definition_xhtml_refFormCallback.probe.stageOfInterest)
	}
	a_datatype_definition_xhtml_ref_ := a_datatype_definition_xhtml_refFormCallback.a_datatype_definition_xhtml_ref
	_ = a_datatype_definition_xhtml_ref_

	for _, formDiv := range a_datatype_definition_xhtml_refFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(a_datatype_definition_xhtml_ref_.Name), formDiv)
		case "DATATYPE_DEFINITION_XHTML_REF":
			FormDivBasicFieldToField(&(a_datatype_definition_xhtml_ref_.DATATYPE_DEFINITION_XHTML_REF), formDiv)
		}
	}

	// manage the suppress operation
	if a_datatype_definition_xhtml_refFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_datatype_definition_xhtml_ref_.Unstage(a_datatype_definition_xhtml_refFormCallback.probe.stageOfInterest)
	}

	a_datatype_definition_xhtml_refFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.A_DATATYPE_DEFINITION_XHTML_REF](
		a_datatype_definition_xhtml_refFormCallback.probe,
	)

	// display a new form by reset the form stage
	if a_datatype_definition_xhtml_refFormCallback.CreationMode || a_datatype_definition_xhtml_refFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_datatype_definition_xhtml_refFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(a_datatype_definition_xhtml_refFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__A_DATATYPE_DEFINITION_XHTML_REFFormCallback(
			nil,
			a_datatype_definition_xhtml_refFormCallback.probe,
			newFormGroup,
		)
		a_datatype_definition_xhtml_ref := new(models.A_DATATYPE_DEFINITION_XHTML_REF)
		FillUpForm(a_datatype_definition_xhtml_ref, newFormGroup, a_datatype_definition_xhtml_refFormCallback.probe)
		a_datatype_definition_xhtml_refFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(a_datatype_definition_xhtml_refFormCallback.probe)
}

func __gong__New__A_EDITABLE_ATTSFormCallback(
	a_editable_atts *models.A_EDITABLE_ATTS,
	probe *Probe,
	formGroup *table.FormGroup,
) (a_editable_attsFormCallback *A_EDITABLE_ATTSFormCallback) {
	a_editable_attsFormCallback = new(A_EDITABLE_ATTSFormCallback)
	a_editable_attsFormCallback.probe = probe
	a_editable_attsFormCallback.a_editable_atts = a_editable_atts
	a_editable_attsFormCallback.formGroup = formGroup

	a_editable_attsFormCallback.CreationMode = (a_editable_atts == nil)

	return
}

type A_EDITABLE_ATTSFormCallback struct {
	a_editable_atts *models.A_EDITABLE_ATTS

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (a_editable_attsFormCallback *A_EDITABLE_ATTSFormCallback) OnSave() {
	a_editable_attsFormCallback.probe.stageOfInterest.Lock()
	defer a_editable_attsFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("A_EDITABLE_ATTSFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	a_editable_attsFormCallback.probe.formStage.Checkout()

	if a_editable_attsFormCallback.a_editable_atts == nil {
		a_editable_attsFormCallback.a_editable_atts = new(models.A_EDITABLE_ATTS).Stage(a_editable_attsFormCallback.probe.stageOfInterest)
	}
	a_editable_atts_ := a_editable_attsFormCallback.a_editable_atts
	_ = a_editable_atts_

	for _, formDiv := range a_editable_attsFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(a_editable_atts_.Name), formDiv)
		case "ATTRIBUTE_DEFINITION_BOOLEAN_REF":
			FormDivBasicFieldToField(&(a_editable_atts_.ATTRIBUTE_DEFINITION_BOOLEAN_REF), formDiv)
		case "ATTRIBUTE_DEFINITION_DATE_REF":
			FormDivBasicFieldToField(&(a_editable_atts_.ATTRIBUTE_DEFINITION_DATE_REF), formDiv)
		case "ATTRIBUTE_DEFINITION_ENUMERATION_REF":
			FormDivBasicFieldToField(&(a_editable_atts_.ATTRIBUTE_DEFINITION_ENUMERATION_REF), formDiv)
		case "ATTRIBUTE_DEFINITION_INTEGER_REF":
			FormDivBasicFieldToField(&(a_editable_atts_.ATTRIBUTE_DEFINITION_INTEGER_REF), formDiv)
		case "ATTRIBUTE_DEFINITION_REAL_REF":
			FormDivBasicFieldToField(&(a_editable_atts_.ATTRIBUTE_DEFINITION_REAL_REF), formDiv)
		case "ATTRIBUTE_DEFINITION_STRING_REF":
			FormDivBasicFieldToField(&(a_editable_atts_.ATTRIBUTE_DEFINITION_STRING_REF), formDiv)
		case "ATTRIBUTE_DEFINITION_XHTML_REF":
			FormDivBasicFieldToField(&(a_editable_atts_.ATTRIBUTE_DEFINITION_XHTML_REF), formDiv)
		}
	}

	// manage the suppress operation
	if a_editable_attsFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_editable_atts_.Unstage(a_editable_attsFormCallback.probe.stageOfInterest)
	}

	a_editable_attsFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.A_EDITABLE_ATTS](
		a_editable_attsFormCallback.probe,
	)

	// display a new form by reset the form stage
	if a_editable_attsFormCallback.CreationMode || a_editable_attsFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_editable_attsFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(a_editable_attsFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__A_EDITABLE_ATTSFormCallback(
			nil,
			a_editable_attsFormCallback.probe,
			newFormGroup,
		)
		a_editable_atts := new(models.A_EDITABLE_ATTS)
		FillUpForm(a_editable_atts, newFormGroup, a_editable_attsFormCallback.probe)
		a_editable_attsFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(a_editable_attsFormCallback.probe)
}

func __gong__New__A_ENUM_VALUE_REFFormCallback(
	a_enum_value_ref *models.A_ENUM_VALUE_REF,
	probe *Probe,
	formGroup *table.FormGroup,
) (a_enum_value_refFormCallback *A_ENUM_VALUE_REFFormCallback) {
	a_enum_value_refFormCallback = new(A_ENUM_VALUE_REFFormCallback)
	a_enum_value_refFormCallback.probe = probe
	a_enum_value_refFormCallback.a_enum_value_ref = a_enum_value_ref
	a_enum_value_refFormCallback.formGroup = formGroup

	a_enum_value_refFormCallback.CreationMode = (a_enum_value_ref == nil)

	return
}

type A_ENUM_VALUE_REFFormCallback struct {
	a_enum_value_ref *models.A_ENUM_VALUE_REF

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (a_enum_value_refFormCallback *A_ENUM_VALUE_REFFormCallback) OnSave() {
	a_enum_value_refFormCallback.probe.stageOfInterest.Lock()
	defer a_enum_value_refFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("A_ENUM_VALUE_REFFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	a_enum_value_refFormCallback.probe.formStage.Checkout()

	if a_enum_value_refFormCallback.a_enum_value_ref == nil {
		a_enum_value_refFormCallback.a_enum_value_ref = new(models.A_ENUM_VALUE_REF).Stage(a_enum_value_refFormCallback.probe.stageOfInterest)
	}
	a_enum_value_ref_ := a_enum_value_refFormCallback.a_enum_value_ref
	_ = a_enum_value_ref_

	for _, formDiv := range a_enum_value_refFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(a_enum_value_ref_.Name), formDiv)
		case "ENUM_VALUE_REF":
			FormDivBasicFieldToField(&(a_enum_value_ref_.ENUM_VALUE_REF), formDiv)
		}
	}

	// manage the suppress operation
	if a_enum_value_refFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_enum_value_ref_.Unstage(a_enum_value_refFormCallback.probe.stageOfInterest)
	}

	a_enum_value_refFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.A_ENUM_VALUE_REF](
		a_enum_value_refFormCallback.probe,
	)

	// display a new form by reset the form stage
	if a_enum_value_refFormCallback.CreationMode || a_enum_value_refFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_enum_value_refFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(a_enum_value_refFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__A_ENUM_VALUE_REFFormCallback(
			nil,
			a_enum_value_refFormCallback.probe,
			newFormGroup,
		)
		a_enum_value_ref := new(models.A_ENUM_VALUE_REF)
		FillUpForm(a_enum_value_ref, newFormGroup, a_enum_value_refFormCallback.probe)
		a_enum_value_refFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(a_enum_value_refFormCallback.probe)
}

func __gong__New__A_OBJECTFormCallback(
	a_object *models.A_OBJECT,
	probe *Probe,
	formGroup *table.FormGroup,
) (a_objectFormCallback *A_OBJECTFormCallback) {
	a_objectFormCallback = new(A_OBJECTFormCallback)
	a_objectFormCallback.probe = probe
	a_objectFormCallback.a_object = a_object
	a_objectFormCallback.formGroup = formGroup

	a_objectFormCallback.CreationMode = (a_object == nil)

	return
}

type A_OBJECTFormCallback struct {
	a_object *models.A_OBJECT

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (a_objectFormCallback *A_OBJECTFormCallback) OnSave() {
	a_objectFormCallback.probe.stageOfInterest.Lock()
	defer a_objectFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("A_OBJECTFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	a_objectFormCallback.probe.formStage.Checkout()

	if a_objectFormCallback.a_object == nil {
		a_objectFormCallback.a_object = new(models.A_OBJECT).Stage(a_objectFormCallback.probe.stageOfInterest)
	}
	a_object_ := a_objectFormCallback.a_object
	_ = a_object_

	for _, formDiv := range a_objectFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(a_object_.Name), formDiv)
		case "SPEC_OBJECT_REF":
			FormDivBasicFieldToField(&(a_object_.SPEC_OBJECT_REF), formDiv)
		}
	}

	// manage the suppress operation
	if a_objectFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_object_.Unstage(a_objectFormCallback.probe.stageOfInterest)
	}

	a_objectFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.A_OBJECT](
		a_objectFormCallback.probe,
	)

	// display a new form by reset the form stage
	if a_objectFormCallback.CreationMode || a_objectFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_objectFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(a_objectFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__A_OBJECTFormCallback(
			nil,
			a_objectFormCallback.probe,
			newFormGroup,
		)
		a_object := new(models.A_OBJECT)
		FillUpForm(a_object, newFormGroup, a_objectFormCallback.probe)
		a_objectFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(a_objectFormCallback.probe)
}

func __gong__New__A_PROPERTIESFormCallback(
	a_properties *models.A_PROPERTIES,
	probe *Probe,
	formGroup *table.FormGroup,
) (a_propertiesFormCallback *A_PROPERTIESFormCallback) {
	a_propertiesFormCallback = new(A_PROPERTIESFormCallback)
	a_propertiesFormCallback.probe = probe
	a_propertiesFormCallback.a_properties = a_properties
	a_propertiesFormCallback.formGroup = formGroup

	a_propertiesFormCallback.CreationMode = (a_properties == nil)

	return
}

type A_PROPERTIESFormCallback struct {
	a_properties *models.A_PROPERTIES

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (a_propertiesFormCallback *A_PROPERTIESFormCallback) OnSave() {
	a_propertiesFormCallback.probe.stageOfInterest.Lock()
	defer a_propertiesFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("A_PROPERTIESFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	a_propertiesFormCallback.probe.formStage.Checkout()

	if a_propertiesFormCallback.a_properties == nil {
		a_propertiesFormCallback.a_properties = new(models.A_PROPERTIES).Stage(a_propertiesFormCallback.probe.stageOfInterest)
	}
	a_properties_ := a_propertiesFormCallback.a_properties
	_ = a_properties_

	for _, formDiv := range a_propertiesFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(a_properties_.Name), formDiv)
		case "EMBEDDED_VALUE":
			FormDivSelectFieldToField(&(a_properties_.EMBEDDED_VALUE), a_propertiesFormCallback.probe.stageOfInterest, formDiv)
		}
	}

	// manage the suppress operation
	if a_propertiesFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_properties_.Unstage(a_propertiesFormCallback.probe.stageOfInterest)
	}

	a_propertiesFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.A_PROPERTIES](
		a_propertiesFormCallback.probe,
	)

	// display a new form by reset the form stage
	if a_propertiesFormCallback.CreationMode || a_propertiesFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_propertiesFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(a_propertiesFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__A_PROPERTIESFormCallback(
			nil,
			a_propertiesFormCallback.probe,
			newFormGroup,
		)
		a_properties := new(models.A_PROPERTIES)
		FillUpForm(a_properties, newFormGroup, a_propertiesFormCallback.probe)
		a_propertiesFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(a_propertiesFormCallback.probe)
}

func __gong__New__A_RELATION_GROUP_TYPE_REFFormCallback(
	a_relation_group_type_ref *models.A_RELATION_GROUP_TYPE_REF,
	probe *Probe,
	formGroup *table.FormGroup,
) (a_relation_group_type_refFormCallback *A_RELATION_GROUP_TYPE_REFFormCallback) {
	a_relation_group_type_refFormCallback = new(A_RELATION_GROUP_TYPE_REFFormCallback)
	a_relation_group_type_refFormCallback.probe = probe
	a_relation_group_type_refFormCallback.a_relation_group_type_ref = a_relation_group_type_ref
	a_relation_group_type_refFormCallback.formGroup = formGroup

	a_relation_group_type_refFormCallback.CreationMode = (a_relation_group_type_ref == nil)

	return
}

type A_RELATION_GROUP_TYPE_REFFormCallback struct {
	a_relation_group_type_ref *models.A_RELATION_GROUP_TYPE_REF

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (a_relation_group_type_refFormCallback *A_RELATION_GROUP_TYPE_REFFormCallback) OnSave() {
	a_relation_group_type_refFormCallback.probe.stageOfInterest.Lock()
	defer a_relation_group_type_refFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("A_RELATION_GROUP_TYPE_REFFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	a_relation_group_type_refFormCallback.probe.formStage.Checkout()

	if a_relation_group_type_refFormCallback.a_relation_group_type_ref == nil {
		a_relation_group_type_refFormCallback.a_relation_group_type_ref = new(models.A_RELATION_GROUP_TYPE_REF).Stage(a_relation_group_type_refFormCallback.probe.stageOfInterest)
	}
	a_relation_group_type_ref_ := a_relation_group_type_refFormCallback.a_relation_group_type_ref
	_ = a_relation_group_type_ref_

	for _, formDiv := range a_relation_group_type_refFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(a_relation_group_type_ref_.Name), formDiv)
		case "RELATION_GROUP_TYPE_REF":
			FormDivBasicFieldToField(&(a_relation_group_type_ref_.RELATION_GROUP_TYPE_REF), formDiv)
		}
	}

	// manage the suppress operation
	if a_relation_group_type_refFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_relation_group_type_ref_.Unstage(a_relation_group_type_refFormCallback.probe.stageOfInterest)
	}

	a_relation_group_type_refFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.A_RELATION_GROUP_TYPE_REF](
		a_relation_group_type_refFormCallback.probe,
	)

	// display a new form by reset the form stage
	if a_relation_group_type_refFormCallback.CreationMode || a_relation_group_type_refFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_relation_group_type_refFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(a_relation_group_type_refFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__A_RELATION_GROUP_TYPE_REFFormCallback(
			nil,
			a_relation_group_type_refFormCallback.probe,
			newFormGroup,
		)
		a_relation_group_type_ref := new(models.A_RELATION_GROUP_TYPE_REF)
		FillUpForm(a_relation_group_type_ref, newFormGroup, a_relation_group_type_refFormCallback.probe)
		a_relation_group_type_refFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(a_relation_group_type_refFormCallback.probe)
}

func __gong__New__A_SOURCE_1FormCallback(
	a_source_1 *models.A_SOURCE_1,
	probe *Probe,
	formGroup *table.FormGroup,
) (a_source_1FormCallback *A_SOURCE_1FormCallback) {
	a_source_1FormCallback = new(A_SOURCE_1FormCallback)
	a_source_1FormCallback.probe = probe
	a_source_1FormCallback.a_source_1 = a_source_1
	a_source_1FormCallback.formGroup = formGroup

	a_source_1FormCallback.CreationMode = (a_source_1 == nil)

	return
}

type A_SOURCE_1FormCallback struct {
	a_source_1 *models.A_SOURCE_1

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (a_source_1FormCallback *A_SOURCE_1FormCallback) OnSave() {
	a_source_1FormCallback.probe.stageOfInterest.Lock()
	defer a_source_1FormCallback.probe.stageOfInterest.Unlock()

	// log.Println("A_SOURCE_1FormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	a_source_1FormCallback.probe.formStage.Checkout()

	if a_source_1FormCallback.a_source_1 == nil {
		a_source_1FormCallback.a_source_1 = new(models.A_SOURCE_1).Stage(a_source_1FormCallback.probe.stageOfInterest)
	}
	a_source_1_ := a_source_1FormCallback.a_source_1
	_ = a_source_1_

	for _, formDiv := range a_source_1FormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(a_source_1_.Name), formDiv)
		case "SPEC_OBJECT_REF":
			FormDivBasicFieldToField(&(a_source_1_.SPEC_OBJECT_REF), formDiv)
		}
	}

	// manage the suppress operation
	if a_source_1FormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_source_1_.Unstage(a_source_1FormCallback.probe.stageOfInterest)
	}

	a_source_1FormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.A_SOURCE_1](
		a_source_1FormCallback.probe,
	)

	// display a new form by reset the form stage
	if a_source_1FormCallback.CreationMode || a_source_1FormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_source_1FormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(a_source_1FormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__A_SOURCE_1FormCallback(
			nil,
			a_source_1FormCallback.probe,
			newFormGroup,
		)
		a_source_1 := new(models.A_SOURCE_1)
		FillUpForm(a_source_1, newFormGroup, a_source_1FormCallback.probe)
		a_source_1FormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(a_source_1FormCallback.probe)
}

func __gong__New__A_SOURCE_SPECIFICATION_1FormCallback(
	a_source_specification_1 *models.A_SOURCE_SPECIFICATION_1,
	probe *Probe,
	formGroup *table.FormGroup,
) (a_source_specification_1FormCallback *A_SOURCE_SPECIFICATION_1FormCallback) {
	a_source_specification_1FormCallback = new(A_SOURCE_SPECIFICATION_1FormCallback)
	a_source_specification_1FormCallback.probe = probe
	a_source_specification_1FormCallback.a_source_specification_1 = a_source_specification_1
	a_source_specification_1FormCallback.formGroup = formGroup

	a_source_specification_1FormCallback.CreationMode = (a_source_specification_1 == nil)

	return
}

type A_SOURCE_SPECIFICATION_1FormCallback struct {
	a_source_specification_1 *models.A_SOURCE_SPECIFICATION_1

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (a_source_specification_1FormCallback *A_SOURCE_SPECIFICATION_1FormCallback) OnSave() {
	a_source_specification_1FormCallback.probe.stageOfInterest.Lock()
	defer a_source_specification_1FormCallback.probe.stageOfInterest.Unlock()

	// log.Println("A_SOURCE_SPECIFICATION_1FormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	a_source_specification_1FormCallback.probe.formStage.Checkout()

	if a_source_specification_1FormCallback.a_source_specification_1 == nil {
		a_source_specification_1FormCallback.a_source_specification_1 = new(models.A_SOURCE_SPECIFICATION_1).Stage(a_source_specification_1FormCallback.probe.stageOfInterest)
	}
	a_source_specification_1_ := a_source_specification_1FormCallback.a_source_specification_1
	_ = a_source_specification_1_

	for _, formDiv := range a_source_specification_1FormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(a_source_specification_1_.Name), formDiv)
		case "SPECIFICATION_REF":
			FormDivEnumStringFieldToField(&(a_source_specification_1_.SPECIFICATION_REF), formDiv)
		}
	}

	// manage the suppress operation
	if a_source_specification_1FormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_source_specification_1_.Unstage(a_source_specification_1FormCallback.probe.stageOfInterest)
	}

	a_source_specification_1FormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.A_SOURCE_SPECIFICATION_1](
		a_source_specification_1FormCallback.probe,
	)

	// display a new form by reset the form stage
	if a_source_specification_1FormCallback.CreationMode || a_source_specification_1FormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_source_specification_1FormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(a_source_specification_1FormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__A_SOURCE_SPECIFICATION_1FormCallback(
			nil,
			a_source_specification_1FormCallback.probe,
			newFormGroup,
		)
		a_source_specification_1 := new(models.A_SOURCE_SPECIFICATION_1)
		FillUpForm(a_source_specification_1, newFormGroup, a_source_specification_1FormCallback.probe)
		a_source_specification_1FormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(a_source_specification_1FormCallback.probe)
}

func __gong__New__A_SPECIFICATIONSFormCallback(
	a_specifications *models.A_SPECIFICATIONS,
	probe *Probe,
	formGroup *table.FormGroup,
) (a_specificationsFormCallback *A_SPECIFICATIONSFormCallback) {
	a_specificationsFormCallback = new(A_SPECIFICATIONSFormCallback)
	a_specificationsFormCallback.probe = probe
	a_specificationsFormCallback.a_specifications = a_specifications
	a_specificationsFormCallback.formGroup = formGroup

	a_specificationsFormCallback.CreationMode = (a_specifications == nil)

	return
}

type A_SPECIFICATIONSFormCallback struct {
	a_specifications *models.A_SPECIFICATIONS

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (a_specificationsFormCallback *A_SPECIFICATIONSFormCallback) OnSave() {
	a_specificationsFormCallback.probe.stageOfInterest.Lock()
	defer a_specificationsFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("A_SPECIFICATIONSFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	a_specificationsFormCallback.probe.formStage.Checkout()

	if a_specificationsFormCallback.a_specifications == nil {
		a_specificationsFormCallback.a_specifications = new(models.A_SPECIFICATIONS).Stage(a_specificationsFormCallback.probe.stageOfInterest)
	}
	a_specifications_ := a_specificationsFormCallback.a_specifications
	_ = a_specifications_

	for _, formDiv := range a_specificationsFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(a_specifications_.Name), formDiv)
		case "SPECIFICATION":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.SPECIFICATION](a_specificationsFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.SPECIFICATION, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.SPECIFICATION)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					a_specificationsFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.SPECIFICATION](a_specificationsFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			a_specifications_.SPECIFICATION = instanceSlice

		}
	}

	// manage the suppress operation
	if a_specificationsFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_specifications_.Unstage(a_specificationsFormCallback.probe.stageOfInterest)
	}

	a_specificationsFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.A_SPECIFICATIONS](
		a_specificationsFormCallback.probe,
	)

	// display a new form by reset the form stage
	if a_specificationsFormCallback.CreationMode || a_specificationsFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_specificationsFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(a_specificationsFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__A_SPECIFICATIONSFormCallback(
			nil,
			a_specificationsFormCallback.probe,
			newFormGroup,
		)
		a_specifications := new(models.A_SPECIFICATIONS)
		FillUpForm(a_specifications, newFormGroup, a_specificationsFormCallback.probe)
		a_specificationsFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(a_specificationsFormCallback.probe)
}

func __gong__New__A_SPECIFICATION_TYPE_REFFormCallback(
	a_specification_type_ref *models.A_SPECIFICATION_TYPE_REF,
	probe *Probe,
	formGroup *table.FormGroup,
) (a_specification_type_refFormCallback *A_SPECIFICATION_TYPE_REFFormCallback) {
	a_specification_type_refFormCallback = new(A_SPECIFICATION_TYPE_REFFormCallback)
	a_specification_type_refFormCallback.probe = probe
	a_specification_type_refFormCallback.a_specification_type_ref = a_specification_type_ref
	a_specification_type_refFormCallback.formGroup = formGroup

	a_specification_type_refFormCallback.CreationMode = (a_specification_type_ref == nil)

	return
}

type A_SPECIFICATION_TYPE_REFFormCallback struct {
	a_specification_type_ref *models.A_SPECIFICATION_TYPE_REF

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (a_specification_type_refFormCallback *A_SPECIFICATION_TYPE_REFFormCallback) OnSave() {
	a_specification_type_refFormCallback.probe.stageOfInterest.Lock()
	defer a_specification_type_refFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("A_SPECIFICATION_TYPE_REFFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	a_specification_type_refFormCallback.probe.formStage.Checkout()

	if a_specification_type_refFormCallback.a_specification_type_ref == nil {
		a_specification_type_refFormCallback.a_specification_type_ref = new(models.A_SPECIFICATION_TYPE_REF).Stage(a_specification_type_refFormCallback.probe.stageOfInterest)
	}
	a_specification_type_ref_ := a_specification_type_refFormCallback.a_specification_type_ref
	_ = a_specification_type_ref_

	for _, formDiv := range a_specification_type_refFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(a_specification_type_ref_.Name), formDiv)
		case "SPECIFICATION_TYPE_REF":
			FormDivBasicFieldToField(&(a_specification_type_ref_.SPECIFICATION_TYPE_REF), formDiv)
		}
	}

	// manage the suppress operation
	if a_specification_type_refFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_specification_type_ref_.Unstage(a_specification_type_refFormCallback.probe.stageOfInterest)
	}

	a_specification_type_refFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.A_SPECIFICATION_TYPE_REF](
		a_specification_type_refFormCallback.probe,
	)

	// display a new form by reset the form stage
	if a_specification_type_refFormCallback.CreationMode || a_specification_type_refFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_specification_type_refFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(a_specification_type_refFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__A_SPECIFICATION_TYPE_REFFormCallback(
			nil,
			a_specification_type_refFormCallback.probe,
			newFormGroup,
		)
		a_specification_type_ref := new(models.A_SPECIFICATION_TYPE_REF)
		FillUpForm(a_specification_type_ref, newFormGroup, a_specification_type_refFormCallback.probe)
		a_specification_type_refFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(a_specification_type_refFormCallback.probe)
}

func __gong__New__A_SPECIFIED_VALUESFormCallback(
	a_specified_values *models.A_SPECIFIED_VALUES,
	probe *Probe,
	formGroup *table.FormGroup,
) (a_specified_valuesFormCallback *A_SPECIFIED_VALUESFormCallback) {
	a_specified_valuesFormCallback = new(A_SPECIFIED_VALUESFormCallback)
	a_specified_valuesFormCallback.probe = probe
	a_specified_valuesFormCallback.a_specified_values = a_specified_values
	a_specified_valuesFormCallback.formGroup = formGroup

	a_specified_valuesFormCallback.CreationMode = (a_specified_values == nil)

	return
}

type A_SPECIFIED_VALUESFormCallback struct {
	a_specified_values *models.A_SPECIFIED_VALUES

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (a_specified_valuesFormCallback *A_SPECIFIED_VALUESFormCallback) OnSave() {
	a_specified_valuesFormCallback.probe.stageOfInterest.Lock()
	defer a_specified_valuesFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("A_SPECIFIED_VALUESFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	a_specified_valuesFormCallback.probe.formStage.Checkout()

	if a_specified_valuesFormCallback.a_specified_values == nil {
		a_specified_valuesFormCallback.a_specified_values = new(models.A_SPECIFIED_VALUES).Stage(a_specified_valuesFormCallback.probe.stageOfInterest)
	}
	a_specified_values_ := a_specified_valuesFormCallback.a_specified_values
	_ = a_specified_values_

	for _, formDiv := range a_specified_valuesFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(a_specified_values_.Name), formDiv)
		case "ENUM_VALUE":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.ENUM_VALUE](a_specified_valuesFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.ENUM_VALUE, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.ENUM_VALUE)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					a_specified_valuesFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.ENUM_VALUE](a_specified_valuesFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			a_specified_values_.ENUM_VALUE = instanceSlice

		}
	}

	// manage the suppress operation
	if a_specified_valuesFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_specified_values_.Unstage(a_specified_valuesFormCallback.probe.stageOfInterest)
	}

	a_specified_valuesFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.A_SPECIFIED_VALUES](
		a_specified_valuesFormCallback.probe,
	)

	// display a new form by reset the form stage
	if a_specified_valuesFormCallback.CreationMode || a_specified_valuesFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_specified_valuesFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(a_specified_valuesFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__A_SPECIFIED_VALUESFormCallback(
			nil,
			a_specified_valuesFormCallback.probe,
			newFormGroup,
		)
		a_specified_values := new(models.A_SPECIFIED_VALUES)
		FillUpForm(a_specified_values, newFormGroup, a_specified_valuesFormCallback.probe)
		a_specified_valuesFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(a_specified_valuesFormCallback.probe)
}

func __gong__New__A_SPEC_ATTRIBUTESFormCallback(
	a_spec_attributes *models.A_SPEC_ATTRIBUTES,
	probe *Probe,
	formGroup *table.FormGroup,
) (a_spec_attributesFormCallback *A_SPEC_ATTRIBUTESFormCallback) {
	a_spec_attributesFormCallback = new(A_SPEC_ATTRIBUTESFormCallback)
	a_spec_attributesFormCallback.probe = probe
	a_spec_attributesFormCallback.a_spec_attributes = a_spec_attributes
	a_spec_attributesFormCallback.formGroup = formGroup

	a_spec_attributesFormCallback.CreationMode = (a_spec_attributes == nil)

	return
}

type A_SPEC_ATTRIBUTESFormCallback struct {
	a_spec_attributes *models.A_SPEC_ATTRIBUTES

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (a_spec_attributesFormCallback *A_SPEC_ATTRIBUTESFormCallback) OnSave() {
	a_spec_attributesFormCallback.probe.stageOfInterest.Lock()
	defer a_spec_attributesFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("A_SPEC_ATTRIBUTESFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	a_spec_attributesFormCallback.probe.formStage.Checkout()

	if a_spec_attributesFormCallback.a_spec_attributes == nil {
		a_spec_attributesFormCallback.a_spec_attributes = new(models.A_SPEC_ATTRIBUTES).Stage(a_spec_attributesFormCallback.probe.stageOfInterest)
	}
	a_spec_attributes_ := a_spec_attributesFormCallback.a_spec_attributes
	_ = a_spec_attributes_

	for _, formDiv := range a_spec_attributesFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(a_spec_attributes_.Name), formDiv)
		case "ATTRIBUTE_DEFINITION_BOOLEAN":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.ATTRIBUTE_DEFINITION_BOOLEAN](a_spec_attributesFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.ATTRIBUTE_DEFINITION_BOOLEAN, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.ATTRIBUTE_DEFINITION_BOOLEAN)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					a_spec_attributesFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.ATTRIBUTE_DEFINITION_BOOLEAN](a_spec_attributesFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			a_spec_attributes_.ATTRIBUTE_DEFINITION_BOOLEAN = instanceSlice

		case "ATTRIBUTE_DEFINITION_DATE":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.ATTRIBUTE_DEFINITION_DATE](a_spec_attributesFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.ATTRIBUTE_DEFINITION_DATE, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.ATTRIBUTE_DEFINITION_DATE)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					a_spec_attributesFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.ATTRIBUTE_DEFINITION_DATE](a_spec_attributesFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			a_spec_attributes_.ATTRIBUTE_DEFINITION_DATE = instanceSlice

		case "ATTRIBUTE_DEFINITION_ENUMERATION":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.ATTRIBUTE_DEFINITION_ENUMERATION](a_spec_attributesFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.ATTRIBUTE_DEFINITION_ENUMERATION, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.ATTRIBUTE_DEFINITION_ENUMERATION)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					a_spec_attributesFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.ATTRIBUTE_DEFINITION_ENUMERATION](a_spec_attributesFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			a_spec_attributes_.ATTRIBUTE_DEFINITION_ENUMERATION = instanceSlice

		case "ATTRIBUTE_DEFINITION_INTEGER":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.ATTRIBUTE_DEFINITION_INTEGER](a_spec_attributesFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.ATTRIBUTE_DEFINITION_INTEGER, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.ATTRIBUTE_DEFINITION_INTEGER)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					a_spec_attributesFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.ATTRIBUTE_DEFINITION_INTEGER](a_spec_attributesFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			a_spec_attributes_.ATTRIBUTE_DEFINITION_INTEGER = instanceSlice

		case "ATTRIBUTE_DEFINITION_REAL":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.ATTRIBUTE_DEFINITION_REAL](a_spec_attributesFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.ATTRIBUTE_DEFINITION_REAL, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.ATTRIBUTE_DEFINITION_REAL)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					a_spec_attributesFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.ATTRIBUTE_DEFINITION_REAL](a_spec_attributesFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			a_spec_attributes_.ATTRIBUTE_DEFINITION_REAL = instanceSlice

		case "ATTRIBUTE_DEFINITION_STRING":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.ATTRIBUTE_DEFINITION_STRING](a_spec_attributesFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.ATTRIBUTE_DEFINITION_STRING, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.ATTRIBUTE_DEFINITION_STRING)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					a_spec_attributesFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.ATTRIBUTE_DEFINITION_STRING](a_spec_attributesFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			a_spec_attributes_.ATTRIBUTE_DEFINITION_STRING = instanceSlice

		case "ATTRIBUTE_DEFINITION_XHTML":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.ATTRIBUTE_DEFINITION_XHTML](a_spec_attributesFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.ATTRIBUTE_DEFINITION_XHTML, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.ATTRIBUTE_DEFINITION_XHTML)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					a_spec_attributesFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.ATTRIBUTE_DEFINITION_XHTML](a_spec_attributesFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			a_spec_attributes_.ATTRIBUTE_DEFINITION_XHTML = instanceSlice

		}
	}

	// manage the suppress operation
	if a_spec_attributesFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_spec_attributes_.Unstage(a_spec_attributesFormCallback.probe.stageOfInterest)
	}

	a_spec_attributesFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.A_SPEC_ATTRIBUTES](
		a_spec_attributesFormCallback.probe,
	)

	// display a new form by reset the form stage
	if a_spec_attributesFormCallback.CreationMode || a_spec_attributesFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_spec_attributesFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(a_spec_attributesFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__A_SPEC_ATTRIBUTESFormCallback(
			nil,
			a_spec_attributesFormCallback.probe,
			newFormGroup,
		)
		a_spec_attributes := new(models.A_SPEC_ATTRIBUTES)
		FillUpForm(a_spec_attributes, newFormGroup, a_spec_attributesFormCallback.probe)
		a_spec_attributesFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(a_spec_attributesFormCallback.probe)
}

func __gong__New__A_SPEC_OBJECTSFormCallback(
	a_spec_objects *models.A_SPEC_OBJECTS,
	probe *Probe,
	formGroup *table.FormGroup,
) (a_spec_objectsFormCallback *A_SPEC_OBJECTSFormCallback) {
	a_spec_objectsFormCallback = new(A_SPEC_OBJECTSFormCallback)
	a_spec_objectsFormCallback.probe = probe
	a_spec_objectsFormCallback.a_spec_objects = a_spec_objects
	a_spec_objectsFormCallback.formGroup = formGroup

	a_spec_objectsFormCallback.CreationMode = (a_spec_objects == nil)

	return
}

type A_SPEC_OBJECTSFormCallback struct {
	a_spec_objects *models.A_SPEC_OBJECTS

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (a_spec_objectsFormCallback *A_SPEC_OBJECTSFormCallback) OnSave() {
	a_spec_objectsFormCallback.probe.stageOfInterest.Lock()
	defer a_spec_objectsFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("A_SPEC_OBJECTSFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	a_spec_objectsFormCallback.probe.formStage.Checkout()

	if a_spec_objectsFormCallback.a_spec_objects == nil {
		a_spec_objectsFormCallback.a_spec_objects = new(models.A_SPEC_OBJECTS).Stage(a_spec_objectsFormCallback.probe.stageOfInterest)
	}
	a_spec_objects_ := a_spec_objectsFormCallback.a_spec_objects
	_ = a_spec_objects_

	for _, formDiv := range a_spec_objectsFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(a_spec_objects_.Name), formDiv)
		case "SPEC_OBJECT":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.SPEC_OBJECT](a_spec_objectsFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.SPEC_OBJECT, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.SPEC_OBJECT)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					a_spec_objectsFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.SPEC_OBJECT](a_spec_objectsFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			a_spec_objects_.SPEC_OBJECT = instanceSlice

		}
	}

	// manage the suppress operation
	if a_spec_objectsFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_spec_objects_.Unstage(a_spec_objectsFormCallback.probe.stageOfInterest)
	}

	a_spec_objectsFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.A_SPEC_OBJECTS](
		a_spec_objectsFormCallback.probe,
	)

	// display a new form by reset the form stage
	if a_spec_objectsFormCallback.CreationMode || a_spec_objectsFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_spec_objectsFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(a_spec_objectsFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__A_SPEC_OBJECTSFormCallback(
			nil,
			a_spec_objectsFormCallback.probe,
			newFormGroup,
		)
		a_spec_objects := new(models.A_SPEC_OBJECTS)
		FillUpForm(a_spec_objects, newFormGroup, a_spec_objectsFormCallback.probe)
		a_spec_objectsFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(a_spec_objectsFormCallback.probe)
}

func __gong__New__A_SPEC_OBJECT_TYPE_REFFormCallback(
	a_spec_object_type_ref *models.A_SPEC_OBJECT_TYPE_REF,
	probe *Probe,
	formGroup *table.FormGroup,
) (a_spec_object_type_refFormCallback *A_SPEC_OBJECT_TYPE_REFFormCallback) {
	a_spec_object_type_refFormCallback = new(A_SPEC_OBJECT_TYPE_REFFormCallback)
	a_spec_object_type_refFormCallback.probe = probe
	a_spec_object_type_refFormCallback.a_spec_object_type_ref = a_spec_object_type_ref
	a_spec_object_type_refFormCallback.formGroup = formGroup

	a_spec_object_type_refFormCallback.CreationMode = (a_spec_object_type_ref == nil)

	return
}

type A_SPEC_OBJECT_TYPE_REFFormCallback struct {
	a_spec_object_type_ref *models.A_SPEC_OBJECT_TYPE_REF

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (a_spec_object_type_refFormCallback *A_SPEC_OBJECT_TYPE_REFFormCallback) OnSave() {
	a_spec_object_type_refFormCallback.probe.stageOfInterest.Lock()
	defer a_spec_object_type_refFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("A_SPEC_OBJECT_TYPE_REFFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	a_spec_object_type_refFormCallback.probe.formStage.Checkout()

	if a_spec_object_type_refFormCallback.a_spec_object_type_ref == nil {
		a_spec_object_type_refFormCallback.a_spec_object_type_ref = new(models.A_SPEC_OBJECT_TYPE_REF).Stage(a_spec_object_type_refFormCallback.probe.stageOfInterest)
	}
	a_spec_object_type_ref_ := a_spec_object_type_refFormCallback.a_spec_object_type_ref
	_ = a_spec_object_type_ref_

	for _, formDiv := range a_spec_object_type_refFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(a_spec_object_type_ref_.Name), formDiv)
		case "SPEC_OBJECT_TYPE_REF":
			FormDivBasicFieldToField(&(a_spec_object_type_ref_.SPEC_OBJECT_TYPE_REF), formDiv)
		}
	}

	// manage the suppress operation
	if a_spec_object_type_refFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_spec_object_type_ref_.Unstage(a_spec_object_type_refFormCallback.probe.stageOfInterest)
	}

	a_spec_object_type_refFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.A_SPEC_OBJECT_TYPE_REF](
		a_spec_object_type_refFormCallback.probe,
	)

	// display a new form by reset the form stage
	if a_spec_object_type_refFormCallback.CreationMode || a_spec_object_type_refFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_spec_object_type_refFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(a_spec_object_type_refFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__A_SPEC_OBJECT_TYPE_REFFormCallback(
			nil,
			a_spec_object_type_refFormCallback.probe,
			newFormGroup,
		)
		a_spec_object_type_ref := new(models.A_SPEC_OBJECT_TYPE_REF)
		FillUpForm(a_spec_object_type_ref, newFormGroup, a_spec_object_type_refFormCallback.probe)
		a_spec_object_type_refFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(a_spec_object_type_refFormCallback.probe)
}

func __gong__New__A_SPEC_RELATIONSFormCallback(
	a_spec_relations *models.A_SPEC_RELATIONS,
	probe *Probe,
	formGroup *table.FormGroup,
) (a_spec_relationsFormCallback *A_SPEC_RELATIONSFormCallback) {
	a_spec_relationsFormCallback = new(A_SPEC_RELATIONSFormCallback)
	a_spec_relationsFormCallback.probe = probe
	a_spec_relationsFormCallback.a_spec_relations = a_spec_relations
	a_spec_relationsFormCallback.formGroup = formGroup

	a_spec_relationsFormCallback.CreationMode = (a_spec_relations == nil)

	return
}

type A_SPEC_RELATIONSFormCallback struct {
	a_spec_relations *models.A_SPEC_RELATIONS

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (a_spec_relationsFormCallback *A_SPEC_RELATIONSFormCallback) OnSave() {
	a_spec_relationsFormCallback.probe.stageOfInterest.Lock()
	defer a_spec_relationsFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("A_SPEC_RELATIONSFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	a_spec_relationsFormCallback.probe.formStage.Checkout()

	if a_spec_relationsFormCallback.a_spec_relations == nil {
		a_spec_relationsFormCallback.a_spec_relations = new(models.A_SPEC_RELATIONS).Stage(a_spec_relationsFormCallback.probe.stageOfInterest)
	}
	a_spec_relations_ := a_spec_relationsFormCallback.a_spec_relations
	_ = a_spec_relations_

	for _, formDiv := range a_spec_relationsFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(a_spec_relations_.Name), formDiv)
		case "SPEC_RELATION":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.SPEC_RELATION](a_spec_relationsFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.SPEC_RELATION, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.SPEC_RELATION)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					a_spec_relationsFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.SPEC_RELATION](a_spec_relationsFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			a_spec_relations_.SPEC_RELATION = instanceSlice

		}
	}

	// manage the suppress operation
	if a_spec_relationsFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_spec_relations_.Unstage(a_spec_relationsFormCallback.probe.stageOfInterest)
	}

	a_spec_relationsFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.A_SPEC_RELATIONS](
		a_spec_relationsFormCallback.probe,
	)

	// display a new form by reset the form stage
	if a_spec_relationsFormCallback.CreationMode || a_spec_relationsFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_spec_relationsFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(a_spec_relationsFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__A_SPEC_RELATIONSFormCallback(
			nil,
			a_spec_relationsFormCallback.probe,
			newFormGroup,
		)
		a_spec_relations := new(models.A_SPEC_RELATIONS)
		FillUpForm(a_spec_relations, newFormGroup, a_spec_relationsFormCallback.probe)
		a_spec_relationsFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(a_spec_relationsFormCallback.probe)
}

func __gong__New__A_SPEC_RELATION_GROUPSFormCallback(
	a_spec_relation_groups *models.A_SPEC_RELATION_GROUPS,
	probe *Probe,
	formGroup *table.FormGroup,
) (a_spec_relation_groupsFormCallback *A_SPEC_RELATION_GROUPSFormCallback) {
	a_spec_relation_groupsFormCallback = new(A_SPEC_RELATION_GROUPSFormCallback)
	a_spec_relation_groupsFormCallback.probe = probe
	a_spec_relation_groupsFormCallback.a_spec_relation_groups = a_spec_relation_groups
	a_spec_relation_groupsFormCallback.formGroup = formGroup

	a_spec_relation_groupsFormCallback.CreationMode = (a_spec_relation_groups == nil)

	return
}

type A_SPEC_RELATION_GROUPSFormCallback struct {
	a_spec_relation_groups *models.A_SPEC_RELATION_GROUPS

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (a_spec_relation_groupsFormCallback *A_SPEC_RELATION_GROUPSFormCallback) OnSave() {
	a_spec_relation_groupsFormCallback.probe.stageOfInterest.Lock()
	defer a_spec_relation_groupsFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("A_SPEC_RELATION_GROUPSFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	a_spec_relation_groupsFormCallback.probe.formStage.Checkout()

	if a_spec_relation_groupsFormCallback.a_spec_relation_groups == nil {
		a_spec_relation_groupsFormCallback.a_spec_relation_groups = new(models.A_SPEC_RELATION_GROUPS).Stage(a_spec_relation_groupsFormCallback.probe.stageOfInterest)
	}
	a_spec_relation_groups_ := a_spec_relation_groupsFormCallback.a_spec_relation_groups
	_ = a_spec_relation_groups_

	for _, formDiv := range a_spec_relation_groupsFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(a_spec_relation_groups_.Name), formDiv)
		case "RELATION_GROUP":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.RELATION_GROUP](a_spec_relation_groupsFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.RELATION_GROUP, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.RELATION_GROUP)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					a_spec_relation_groupsFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.RELATION_GROUP](a_spec_relation_groupsFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			a_spec_relation_groups_.RELATION_GROUP = instanceSlice

		}
	}

	// manage the suppress operation
	if a_spec_relation_groupsFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_spec_relation_groups_.Unstage(a_spec_relation_groupsFormCallback.probe.stageOfInterest)
	}

	a_spec_relation_groupsFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.A_SPEC_RELATION_GROUPS](
		a_spec_relation_groupsFormCallback.probe,
	)

	// display a new form by reset the form stage
	if a_spec_relation_groupsFormCallback.CreationMode || a_spec_relation_groupsFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_spec_relation_groupsFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(a_spec_relation_groupsFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__A_SPEC_RELATION_GROUPSFormCallback(
			nil,
			a_spec_relation_groupsFormCallback.probe,
			newFormGroup,
		)
		a_spec_relation_groups := new(models.A_SPEC_RELATION_GROUPS)
		FillUpForm(a_spec_relation_groups, newFormGroup, a_spec_relation_groupsFormCallback.probe)
		a_spec_relation_groupsFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(a_spec_relation_groupsFormCallback.probe)
}

func __gong__New__A_SPEC_RELATION_REFFormCallback(
	a_spec_relation_ref *models.A_SPEC_RELATION_REF,
	probe *Probe,
	formGroup *table.FormGroup,
) (a_spec_relation_refFormCallback *A_SPEC_RELATION_REFFormCallback) {
	a_spec_relation_refFormCallback = new(A_SPEC_RELATION_REFFormCallback)
	a_spec_relation_refFormCallback.probe = probe
	a_spec_relation_refFormCallback.a_spec_relation_ref = a_spec_relation_ref
	a_spec_relation_refFormCallback.formGroup = formGroup

	a_spec_relation_refFormCallback.CreationMode = (a_spec_relation_ref == nil)

	return
}

type A_SPEC_RELATION_REFFormCallback struct {
	a_spec_relation_ref *models.A_SPEC_RELATION_REF

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (a_spec_relation_refFormCallback *A_SPEC_RELATION_REFFormCallback) OnSave() {
	a_spec_relation_refFormCallback.probe.stageOfInterest.Lock()
	defer a_spec_relation_refFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("A_SPEC_RELATION_REFFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	a_spec_relation_refFormCallback.probe.formStage.Checkout()

	if a_spec_relation_refFormCallback.a_spec_relation_ref == nil {
		a_spec_relation_refFormCallback.a_spec_relation_ref = new(models.A_SPEC_RELATION_REF).Stage(a_spec_relation_refFormCallback.probe.stageOfInterest)
	}
	a_spec_relation_ref_ := a_spec_relation_refFormCallback.a_spec_relation_ref
	_ = a_spec_relation_ref_

	for _, formDiv := range a_spec_relation_refFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(a_spec_relation_ref_.Name), formDiv)
		case "SPEC_RELATION_REF":
			FormDivBasicFieldToField(&(a_spec_relation_ref_.SPEC_RELATION_REF), formDiv)
		}
	}

	// manage the suppress operation
	if a_spec_relation_refFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_spec_relation_ref_.Unstage(a_spec_relation_refFormCallback.probe.stageOfInterest)
	}

	a_spec_relation_refFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.A_SPEC_RELATION_REF](
		a_spec_relation_refFormCallback.probe,
	)

	// display a new form by reset the form stage
	if a_spec_relation_refFormCallback.CreationMode || a_spec_relation_refFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_spec_relation_refFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(a_spec_relation_refFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__A_SPEC_RELATION_REFFormCallback(
			nil,
			a_spec_relation_refFormCallback.probe,
			newFormGroup,
		)
		a_spec_relation_ref := new(models.A_SPEC_RELATION_REF)
		FillUpForm(a_spec_relation_ref, newFormGroup, a_spec_relation_refFormCallback.probe)
		a_spec_relation_refFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(a_spec_relation_refFormCallback.probe)
}

func __gong__New__A_SPEC_RELATION_TYPE_REFFormCallback(
	a_spec_relation_type_ref *models.A_SPEC_RELATION_TYPE_REF,
	probe *Probe,
	formGroup *table.FormGroup,
) (a_spec_relation_type_refFormCallback *A_SPEC_RELATION_TYPE_REFFormCallback) {
	a_spec_relation_type_refFormCallback = new(A_SPEC_RELATION_TYPE_REFFormCallback)
	a_spec_relation_type_refFormCallback.probe = probe
	a_spec_relation_type_refFormCallback.a_spec_relation_type_ref = a_spec_relation_type_ref
	a_spec_relation_type_refFormCallback.formGroup = formGroup

	a_spec_relation_type_refFormCallback.CreationMode = (a_spec_relation_type_ref == nil)

	return
}

type A_SPEC_RELATION_TYPE_REFFormCallback struct {
	a_spec_relation_type_ref *models.A_SPEC_RELATION_TYPE_REF

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (a_spec_relation_type_refFormCallback *A_SPEC_RELATION_TYPE_REFFormCallback) OnSave() {
	a_spec_relation_type_refFormCallback.probe.stageOfInterest.Lock()
	defer a_spec_relation_type_refFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("A_SPEC_RELATION_TYPE_REFFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	a_spec_relation_type_refFormCallback.probe.formStage.Checkout()

	if a_spec_relation_type_refFormCallback.a_spec_relation_type_ref == nil {
		a_spec_relation_type_refFormCallback.a_spec_relation_type_ref = new(models.A_SPEC_RELATION_TYPE_REF).Stage(a_spec_relation_type_refFormCallback.probe.stageOfInterest)
	}
	a_spec_relation_type_ref_ := a_spec_relation_type_refFormCallback.a_spec_relation_type_ref
	_ = a_spec_relation_type_ref_

	for _, formDiv := range a_spec_relation_type_refFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(a_spec_relation_type_ref_.Name), formDiv)
		case "SPEC_RELATION_TYPE_REF":
			FormDivBasicFieldToField(&(a_spec_relation_type_ref_.SPEC_RELATION_TYPE_REF), formDiv)
		}
	}

	// manage the suppress operation
	if a_spec_relation_type_refFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_spec_relation_type_ref_.Unstage(a_spec_relation_type_refFormCallback.probe.stageOfInterest)
	}

	a_spec_relation_type_refFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.A_SPEC_RELATION_TYPE_REF](
		a_spec_relation_type_refFormCallback.probe,
	)

	// display a new form by reset the form stage
	if a_spec_relation_type_refFormCallback.CreationMode || a_spec_relation_type_refFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_spec_relation_type_refFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(a_spec_relation_type_refFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__A_SPEC_RELATION_TYPE_REFFormCallback(
			nil,
			a_spec_relation_type_refFormCallback.probe,
			newFormGroup,
		)
		a_spec_relation_type_ref := new(models.A_SPEC_RELATION_TYPE_REF)
		FillUpForm(a_spec_relation_type_ref, newFormGroup, a_spec_relation_type_refFormCallback.probe)
		a_spec_relation_type_refFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(a_spec_relation_type_refFormCallback.probe)
}

func __gong__New__A_SPEC_TYPESFormCallback(
	a_spec_types *models.A_SPEC_TYPES,
	probe *Probe,
	formGroup *table.FormGroup,
) (a_spec_typesFormCallback *A_SPEC_TYPESFormCallback) {
	a_spec_typesFormCallback = new(A_SPEC_TYPESFormCallback)
	a_spec_typesFormCallback.probe = probe
	a_spec_typesFormCallback.a_spec_types = a_spec_types
	a_spec_typesFormCallback.formGroup = formGroup

	a_spec_typesFormCallback.CreationMode = (a_spec_types == nil)

	return
}

type A_SPEC_TYPESFormCallback struct {
	a_spec_types *models.A_SPEC_TYPES

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (a_spec_typesFormCallback *A_SPEC_TYPESFormCallback) OnSave() {
	a_spec_typesFormCallback.probe.stageOfInterest.Lock()
	defer a_spec_typesFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("A_SPEC_TYPESFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	a_spec_typesFormCallback.probe.formStage.Checkout()

	if a_spec_typesFormCallback.a_spec_types == nil {
		a_spec_typesFormCallback.a_spec_types = new(models.A_SPEC_TYPES).Stage(a_spec_typesFormCallback.probe.stageOfInterest)
	}
	a_spec_types_ := a_spec_typesFormCallback.a_spec_types
	_ = a_spec_types_

	for _, formDiv := range a_spec_typesFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(a_spec_types_.Name), formDiv)
		case "RELATION_GROUP_TYPE":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.RELATION_GROUP_TYPE](a_spec_typesFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.RELATION_GROUP_TYPE, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.RELATION_GROUP_TYPE)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					a_spec_typesFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.RELATION_GROUP_TYPE](a_spec_typesFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			a_spec_types_.RELATION_GROUP_TYPE = instanceSlice

		case "SPEC_OBJECT_TYPE":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.SPEC_OBJECT_TYPE](a_spec_typesFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.SPEC_OBJECT_TYPE, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.SPEC_OBJECT_TYPE)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					a_spec_typesFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.SPEC_OBJECT_TYPE](a_spec_typesFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			a_spec_types_.SPEC_OBJECT_TYPE = instanceSlice

		case "SPEC_RELATION_TYPE":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.SPEC_RELATION_TYPE](a_spec_typesFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.SPEC_RELATION_TYPE, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.SPEC_RELATION_TYPE)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					a_spec_typesFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.SPEC_RELATION_TYPE](a_spec_typesFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			a_spec_types_.SPEC_RELATION_TYPE = instanceSlice

		case "SPECIFICATION_TYPE":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.SPECIFICATION_TYPE](a_spec_typesFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.SPECIFICATION_TYPE, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.SPECIFICATION_TYPE)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					a_spec_typesFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.SPECIFICATION_TYPE](a_spec_typesFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			a_spec_types_.SPECIFICATION_TYPE = instanceSlice

		}
	}

	// manage the suppress operation
	if a_spec_typesFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_spec_types_.Unstage(a_spec_typesFormCallback.probe.stageOfInterest)
	}

	a_spec_typesFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.A_SPEC_TYPES](
		a_spec_typesFormCallback.probe,
	)

	// display a new form by reset the form stage
	if a_spec_typesFormCallback.CreationMode || a_spec_typesFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_spec_typesFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(a_spec_typesFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__A_SPEC_TYPESFormCallback(
			nil,
			a_spec_typesFormCallback.probe,
			newFormGroup,
		)
		a_spec_types := new(models.A_SPEC_TYPES)
		FillUpForm(a_spec_types, newFormGroup, a_spec_typesFormCallback.probe)
		a_spec_typesFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(a_spec_typesFormCallback.probe)
}

func __gong__New__A_THE_HEADERFormCallback(
	a_the_header *models.A_THE_HEADER,
	probe *Probe,
	formGroup *table.FormGroup,
) (a_the_headerFormCallback *A_THE_HEADERFormCallback) {
	a_the_headerFormCallback = new(A_THE_HEADERFormCallback)
	a_the_headerFormCallback.probe = probe
	a_the_headerFormCallback.a_the_header = a_the_header
	a_the_headerFormCallback.formGroup = formGroup

	a_the_headerFormCallback.CreationMode = (a_the_header == nil)

	return
}

type A_THE_HEADERFormCallback struct {
	a_the_header *models.A_THE_HEADER

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (a_the_headerFormCallback *A_THE_HEADERFormCallback) OnSave() {
	a_the_headerFormCallback.probe.stageOfInterest.Lock()
	defer a_the_headerFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("A_THE_HEADERFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	a_the_headerFormCallback.probe.formStage.Checkout()

	if a_the_headerFormCallback.a_the_header == nil {
		a_the_headerFormCallback.a_the_header = new(models.A_THE_HEADER).Stage(a_the_headerFormCallback.probe.stageOfInterest)
	}
	a_the_header_ := a_the_headerFormCallback.a_the_header
	_ = a_the_header_

	for _, formDiv := range a_the_headerFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(a_the_header_.Name), formDiv)
		case "REQ_IF_HEADER":
			FormDivSelectFieldToField(&(a_the_header_.REQ_IF_HEADER), a_the_headerFormCallback.probe.stageOfInterest, formDiv)
		}
	}

	// manage the suppress operation
	if a_the_headerFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_the_header_.Unstage(a_the_headerFormCallback.probe.stageOfInterest)
	}

	a_the_headerFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.A_THE_HEADER](
		a_the_headerFormCallback.probe,
	)

	// display a new form by reset the form stage
	if a_the_headerFormCallback.CreationMode || a_the_headerFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_the_headerFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(a_the_headerFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__A_THE_HEADERFormCallback(
			nil,
			a_the_headerFormCallback.probe,
			newFormGroup,
		)
		a_the_header := new(models.A_THE_HEADER)
		FillUpForm(a_the_header, newFormGroup, a_the_headerFormCallback.probe)
		a_the_headerFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(a_the_headerFormCallback.probe)
}

func __gong__New__A_TOOL_EXTENSIONSFormCallback(
	a_tool_extensions *models.A_TOOL_EXTENSIONS,
	probe *Probe,
	formGroup *table.FormGroup,
) (a_tool_extensionsFormCallback *A_TOOL_EXTENSIONSFormCallback) {
	a_tool_extensionsFormCallback = new(A_TOOL_EXTENSIONSFormCallback)
	a_tool_extensionsFormCallback.probe = probe
	a_tool_extensionsFormCallback.a_tool_extensions = a_tool_extensions
	a_tool_extensionsFormCallback.formGroup = formGroup

	a_tool_extensionsFormCallback.CreationMode = (a_tool_extensions == nil)

	return
}

type A_TOOL_EXTENSIONSFormCallback struct {
	a_tool_extensions *models.A_TOOL_EXTENSIONS

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (a_tool_extensionsFormCallback *A_TOOL_EXTENSIONSFormCallback) OnSave() {
	a_tool_extensionsFormCallback.probe.stageOfInterest.Lock()
	defer a_tool_extensionsFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("A_TOOL_EXTENSIONSFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	a_tool_extensionsFormCallback.probe.formStage.Checkout()

	if a_tool_extensionsFormCallback.a_tool_extensions == nil {
		a_tool_extensionsFormCallback.a_tool_extensions = new(models.A_TOOL_EXTENSIONS).Stage(a_tool_extensionsFormCallback.probe.stageOfInterest)
	}
	a_tool_extensions_ := a_tool_extensionsFormCallback.a_tool_extensions
	_ = a_tool_extensions_

	for _, formDiv := range a_tool_extensionsFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(a_tool_extensions_.Name), formDiv)
		case "REQ_IF_TOOL_EXTENSION":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.REQ_IF_TOOL_EXTENSION](a_tool_extensionsFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.REQ_IF_TOOL_EXTENSION, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.REQ_IF_TOOL_EXTENSION)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					a_tool_extensionsFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.REQ_IF_TOOL_EXTENSION](a_tool_extensionsFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			a_tool_extensions_.REQ_IF_TOOL_EXTENSION = instanceSlice

		}
	}

	// manage the suppress operation
	if a_tool_extensionsFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_tool_extensions_.Unstage(a_tool_extensionsFormCallback.probe.stageOfInterest)
	}

	a_tool_extensionsFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.A_TOOL_EXTENSIONS](
		a_tool_extensionsFormCallback.probe,
	)

	// display a new form by reset the form stage
	if a_tool_extensionsFormCallback.CreationMode || a_tool_extensionsFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_tool_extensionsFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(a_tool_extensionsFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__A_TOOL_EXTENSIONSFormCallback(
			nil,
			a_tool_extensionsFormCallback.probe,
			newFormGroup,
		)
		a_tool_extensions := new(models.A_TOOL_EXTENSIONS)
		FillUpForm(a_tool_extensions, newFormGroup, a_tool_extensionsFormCallback.probe)
		a_tool_extensionsFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(a_tool_extensionsFormCallback.probe)
}

func __gong__New__DATATYPE_DEFINITION_BOOLEANFormCallback(
	datatype_definition_boolean *models.DATATYPE_DEFINITION_BOOLEAN,
	probe *Probe,
	formGroup *table.FormGroup,
) (datatype_definition_booleanFormCallback *DATATYPE_DEFINITION_BOOLEANFormCallback) {
	datatype_definition_booleanFormCallback = new(DATATYPE_DEFINITION_BOOLEANFormCallback)
	datatype_definition_booleanFormCallback.probe = probe
	datatype_definition_booleanFormCallback.datatype_definition_boolean = datatype_definition_boolean
	datatype_definition_booleanFormCallback.formGroup = formGroup

	datatype_definition_booleanFormCallback.CreationMode = (datatype_definition_boolean == nil)

	return
}

type DATATYPE_DEFINITION_BOOLEANFormCallback struct {
	datatype_definition_boolean *models.DATATYPE_DEFINITION_BOOLEAN

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (datatype_definition_booleanFormCallback *DATATYPE_DEFINITION_BOOLEANFormCallback) OnSave() {
	datatype_definition_booleanFormCallback.probe.stageOfInterest.Lock()
	defer datatype_definition_booleanFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("DATATYPE_DEFINITION_BOOLEANFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	datatype_definition_booleanFormCallback.probe.formStage.Checkout()

	if datatype_definition_booleanFormCallback.datatype_definition_boolean == nil {
		datatype_definition_booleanFormCallback.datatype_definition_boolean = new(models.DATATYPE_DEFINITION_BOOLEAN).Stage(datatype_definition_booleanFormCallback.probe.stageOfInterest)
	}
	datatype_definition_boolean_ := datatype_definition_booleanFormCallback.datatype_definition_boolean
	_ = datatype_definition_boolean_

	for _, formDiv := range datatype_definition_booleanFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(datatype_definition_boolean_.Name), formDiv)
		case "DESC":
			FormDivBasicFieldToField(&(datatype_definition_boolean_.DESC), formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&(datatype_definition_boolean_.IDENTIFIER), formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&(datatype_definition_boolean_.LAST_CHANGE), formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&(datatype_definition_boolean_.LONG_NAME), formDiv)
		case "ALTERNATIVE_ID":
			FormDivSelectFieldToField(&(datatype_definition_boolean_.ALTERNATIVE_ID), datatype_definition_booleanFormCallback.probe.stageOfInterest, formDiv)
		case "A_DATATYPES:DATATYPE_DEFINITION_BOOLEAN":
			// WARNING : this form deals with the N-N association "A_DATATYPES.DATATYPE_DEFINITION_BOOLEAN []*DATATYPE_DEFINITION_BOOLEAN" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of DATATYPE_DEFINITION_BOOLEAN). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.A_DATATYPES
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "A_DATATYPES"
				rf.Fieldname = "DATATYPE_DEFINITION_BOOLEAN"
				formerAssociationSource := datatype_definition_boolean_.GongGetReverseFieldOwner(
					datatype_definition_booleanFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.A_DATATYPES)
					if !ok {
						log.Fatalln("Source of A_DATATYPES.DATATYPE_DEFINITION_BOOLEAN []*DATATYPE_DEFINITION_BOOLEAN, is not an A_DATATYPES instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.DATATYPE_DEFINITION_BOOLEAN, datatype_definition_boolean_)
					formerSource.DATATYPE_DEFINITION_BOOLEAN = slices.Delete(formerSource.DATATYPE_DEFINITION_BOOLEAN, idx, idx+1)
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
			var newSource *models.A_DATATYPES
			for _a_datatypes := range *models.GetGongstructInstancesSet[models.A_DATATYPES](datatype_definition_booleanFormCallback.probe.stageOfInterest) {
				// the match is base on the name
				if _a_datatypes.GetName() == newSourceName.GetName() {
					newSource = _a_datatypes // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of A_DATATYPES.DATATYPE_DEFINITION_BOOLEAN []*DATATYPE_DEFINITION_BOOLEAN, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.DATATYPE_DEFINITION_BOOLEAN = append(newSource.DATATYPE_DEFINITION_BOOLEAN, datatype_definition_boolean_)
		}
	}

	// manage the suppress operation
	if datatype_definition_booleanFormCallback.formGroup.HasSuppressButtonBeenPressed {
		datatype_definition_boolean_.Unstage(datatype_definition_booleanFormCallback.probe.stageOfInterest)
	}

	datatype_definition_booleanFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.DATATYPE_DEFINITION_BOOLEAN](
		datatype_definition_booleanFormCallback.probe,
	)

	// display a new form by reset the form stage
	if datatype_definition_booleanFormCallback.CreationMode || datatype_definition_booleanFormCallback.formGroup.HasSuppressButtonBeenPressed {
		datatype_definition_booleanFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(datatype_definition_booleanFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__DATATYPE_DEFINITION_BOOLEANFormCallback(
			nil,
			datatype_definition_booleanFormCallback.probe,
			newFormGroup,
		)
		datatype_definition_boolean := new(models.DATATYPE_DEFINITION_BOOLEAN)
		FillUpForm(datatype_definition_boolean, newFormGroup, datatype_definition_booleanFormCallback.probe)
		datatype_definition_booleanFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(datatype_definition_booleanFormCallback.probe)
}

func __gong__New__DATATYPE_DEFINITION_DATEFormCallback(
	datatype_definition_date *models.DATATYPE_DEFINITION_DATE,
	probe *Probe,
	formGroup *table.FormGroup,
) (datatype_definition_dateFormCallback *DATATYPE_DEFINITION_DATEFormCallback) {
	datatype_definition_dateFormCallback = new(DATATYPE_DEFINITION_DATEFormCallback)
	datatype_definition_dateFormCallback.probe = probe
	datatype_definition_dateFormCallback.datatype_definition_date = datatype_definition_date
	datatype_definition_dateFormCallback.formGroup = formGroup

	datatype_definition_dateFormCallback.CreationMode = (datatype_definition_date == nil)

	return
}

type DATATYPE_DEFINITION_DATEFormCallback struct {
	datatype_definition_date *models.DATATYPE_DEFINITION_DATE

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (datatype_definition_dateFormCallback *DATATYPE_DEFINITION_DATEFormCallback) OnSave() {
	datatype_definition_dateFormCallback.probe.stageOfInterest.Lock()
	defer datatype_definition_dateFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("DATATYPE_DEFINITION_DATEFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	datatype_definition_dateFormCallback.probe.formStage.Checkout()

	if datatype_definition_dateFormCallback.datatype_definition_date == nil {
		datatype_definition_dateFormCallback.datatype_definition_date = new(models.DATATYPE_DEFINITION_DATE).Stage(datatype_definition_dateFormCallback.probe.stageOfInterest)
	}
	datatype_definition_date_ := datatype_definition_dateFormCallback.datatype_definition_date
	_ = datatype_definition_date_

	for _, formDiv := range datatype_definition_dateFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(datatype_definition_date_.Name), formDiv)
		case "DESC":
			FormDivBasicFieldToField(&(datatype_definition_date_.DESC), formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&(datatype_definition_date_.IDENTIFIER), formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&(datatype_definition_date_.LAST_CHANGE), formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&(datatype_definition_date_.LONG_NAME), formDiv)
		case "ALTERNATIVE_ID":
			FormDivSelectFieldToField(&(datatype_definition_date_.ALTERNATIVE_ID), datatype_definition_dateFormCallback.probe.stageOfInterest, formDiv)
		case "A_DATATYPES:DATATYPE_DEFINITION_DATE":
			// WARNING : this form deals with the N-N association "A_DATATYPES.DATATYPE_DEFINITION_DATE []*DATATYPE_DEFINITION_DATE" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of DATATYPE_DEFINITION_DATE). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.A_DATATYPES
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "A_DATATYPES"
				rf.Fieldname = "DATATYPE_DEFINITION_DATE"
				formerAssociationSource := datatype_definition_date_.GongGetReverseFieldOwner(
					datatype_definition_dateFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.A_DATATYPES)
					if !ok {
						log.Fatalln("Source of A_DATATYPES.DATATYPE_DEFINITION_DATE []*DATATYPE_DEFINITION_DATE, is not an A_DATATYPES instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.DATATYPE_DEFINITION_DATE, datatype_definition_date_)
					formerSource.DATATYPE_DEFINITION_DATE = slices.Delete(formerSource.DATATYPE_DEFINITION_DATE, idx, idx+1)
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
			var newSource *models.A_DATATYPES
			for _a_datatypes := range *models.GetGongstructInstancesSet[models.A_DATATYPES](datatype_definition_dateFormCallback.probe.stageOfInterest) {
				// the match is base on the name
				if _a_datatypes.GetName() == newSourceName.GetName() {
					newSource = _a_datatypes // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of A_DATATYPES.DATATYPE_DEFINITION_DATE []*DATATYPE_DEFINITION_DATE, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.DATATYPE_DEFINITION_DATE = append(newSource.DATATYPE_DEFINITION_DATE, datatype_definition_date_)
		}
	}

	// manage the suppress operation
	if datatype_definition_dateFormCallback.formGroup.HasSuppressButtonBeenPressed {
		datatype_definition_date_.Unstage(datatype_definition_dateFormCallback.probe.stageOfInterest)
	}

	datatype_definition_dateFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.DATATYPE_DEFINITION_DATE](
		datatype_definition_dateFormCallback.probe,
	)

	// display a new form by reset the form stage
	if datatype_definition_dateFormCallback.CreationMode || datatype_definition_dateFormCallback.formGroup.HasSuppressButtonBeenPressed {
		datatype_definition_dateFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(datatype_definition_dateFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__DATATYPE_DEFINITION_DATEFormCallback(
			nil,
			datatype_definition_dateFormCallback.probe,
			newFormGroup,
		)
		datatype_definition_date := new(models.DATATYPE_DEFINITION_DATE)
		FillUpForm(datatype_definition_date, newFormGroup, datatype_definition_dateFormCallback.probe)
		datatype_definition_dateFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(datatype_definition_dateFormCallback.probe)
}

func __gong__New__DATATYPE_DEFINITION_ENUMERATIONFormCallback(
	datatype_definition_enumeration *models.DATATYPE_DEFINITION_ENUMERATION,
	probe *Probe,
	formGroup *table.FormGroup,
) (datatype_definition_enumerationFormCallback *DATATYPE_DEFINITION_ENUMERATIONFormCallback) {
	datatype_definition_enumerationFormCallback = new(DATATYPE_DEFINITION_ENUMERATIONFormCallback)
	datatype_definition_enumerationFormCallback.probe = probe
	datatype_definition_enumerationFormCallback.datatype_definition_enumeration = datatype_definition_enumeration
	datatype_definition_enumerationFormCallback.formGroup = formGroup

	datatype_definition_enumerationFormCallback.CreationMode = (datatype_definition_enumeration == nil)

	return
}

type DATATYPE_DEFINITION_ENUMERATIONFormCallback struct {
	datatype_definition_enumeration *models.DATATYPE_DEFINITION_ENUMERATION

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (datatype_definition_enumerationFormCallback *DATATYPE_DEFINITION_ENUMERATIONFormCallback) OnSave() {
	datatype_definition_enumerationFormCallback.probe.stageOfInterest.Lock()
	defer datatype_definition_enumerationFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("DATATYPE_DEFINITION_ENUMERATIONFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	datatype_definition_enumerationFormCallback.probe.formStage.Checkout()

	if datatype_definition_enumerationFormCallback.datatype_definition_enumeration == nil {
		datatype_definition_enumerationFormCallback.datatype_definition_enumeration = new(models.DATATYPE_DEFINITION_ENUMERATION).Stage(datatype_definition_enumerationFormCallback.probe.stageOfInterest)
	}
	datatype_definition_enumeration_ := datatype_definition_enumerationFormCallback.datatype_definition_enumeration
	_ = datatype_definition_enumeration_

	for _, formDiv := range datatype_definition_enumerationFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(datatype_definition_enumeration_.Name), formDiv)
		case "DESC":
			FormDivBasicFieldToField(&(datatype_definition_enumeration_.DESC), formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&(datatype_definition_enumeration_.IDENTIFIER), formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&(datatype_definition_enumeration_.LAST_CHANGE), formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&(datatype_definition_enumeration_.LONG_NAME), formDiv)
		case "ALTERNATIVE_ID":
			FormDivSelectFieldToField(&(datatype_definition_enumeration_.ALTERNATIVE_ID), datatype_definition_enumerationFormCallback.probe.stageOfInterest, formDiv)
		case "SPECIFIED_VALUES":
			FormDivSelectFieldToField(&(datatype_definition_enumeration_.SPECIFIED_VALUES), datatype_definition_enumerationFormCallback.probe.stageOfInterest, formDiv)
		case "A_DATATYPES:DATATYPE_DEFINITION_ENUMERATION":
			// WARNING : this form deals with the N-N association "A_DATATYPES.DATATYPE_DEFINITION_ENUMERATION []*DATATYPE_DEFINITION_ENUMERATION" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of DATATYPE_DEFINITION_ENUMERATION). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.A_DATATYPES
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "A_DATATYPES"
				rf.Fieldname = "DATATYPE_DEFINITION_ENUMERATION"
				formerAssociationSource := datatype_definition_enumeration_.GongGetReverseFieldOwner(
					datatype_definition_enumerationFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.A_DATATYPES)
					if !ok {
						log.Fatalln("Source of A_DATATYPES.DATATYPE_DEFINITION_ENUMERATION []*DATATYPE_DEFINITION_ENUMERATION, is not an A_DATATYPES instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.DATATYPE_DEFINITION_ENUMERATION, datatype_definition_enumeration_)
					formerSource.DATATYPE_DEFINITION_ENUMERATION = slices.Delete(formerSource.DATATYPE_DEFINITION_ENUMERATION, idx, idx+1)
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
			var newSource *models.A_DATATYPES
			for _a_datatypes := range *models.GetGongstructInstancesSet[models.A_DATATYPES](datatype_definition_enumerationFormCallback.probe.stageOfInterest) {
				// the match is base on the name
				if _a_datatypes.GetName() == newSourceName.GetName() {
					newSource = _a_datatypes // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of A_DATATYPES.DATATYPE_DEFINITION_ENUMERATION []*DATATYPE_DEFINITION_ENUMERATION, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.DATATYPE_DEFINITION_ENUMERATION = append(newSource.DATATYPE_DEFINITION_ENUMERATION, datatype_definition_enumeration_)
		}
	}

	// manage the suppress operation
	if datatype_definition_enumerationFormCallback.formGroup.HasSuppressButtonBeenPressed {
		datatype_definition_enumeration_.Unstage(datatype_definition_enumerationFormCallback.probe.stageOfInterest)
	}

	datatype_definition_enumerationFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.DATATYPE_DEFINITION_ENUMERATION](
		datatype_definition_enumerationFormCallback.probe,
	)

	// display a new form by reset the form stage
	if datatype_definition_enumerationFormCallback.CreationMode || datatype_definition_enumerationFormCallback.formGroup.HasSuppressButtonBeenPressed {
		datatype_definition_enumerationFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(datatype_definition_enumerationFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__DATATYPE_DEFINITION_ENUMERATIONFormCallback(
			nil,
			datatype_definition_enumerationFormCallback.probe,
			newFormGroup,
		)
		datatype_definition_enumeration := new(models.DATATYPE_DEFINITION_ENUMERATION)
		FillUpForm(datatype_definition_enumeration, newFormGroup, datatype_definition_enumerationFormCallback.probe)
		datatype_definition_enumerationFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(datatype_definition_enumerationFormCallback.probe)
}

func __gong__New__DATATYPE_DEFINITION_INTEGERFormCallback(
	datatype_definition_integer *models.DATATYPE_DEFINITION_INTEGER,
	probe *Probe,
	formGroup *table.FormGroup,
) (datatype_definition_integerFormCallback *DATATYPE_DEFINITION_INTEGERFormCallback) {
	datatype_definition_integerFormCallback = new(DATATYPE_DEFINITION_INTEGERFormCallback)
	datatype_definition_integerFormCallback.probe = probe
	datatype_definition_integerFormCallback.datatype_definition_integer = datatype_definition_integer
	datatype_definition_integerFormCallback.formGroup = formGroup

	datatype_definition_integerFormCallback.CreationMode = (datatype_definition_integer == nil)

	return
}

type DATATYPE_DEFINITION_INTEGERFormCallback struct {
	datatype_definition_integer *models.DATATYPE_DEFINITION_INTEGER

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (datatype_definition_integerFormCallback *DATATYPE_DEFINITION_INTEGERFormCallback) OnSave() {
	datatype_definition_integerFormCallback.probe.stageOfInterest.Lock()
	defer datatype_definition_integerFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("DATATYPE_DEFINITION_INTEGERFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	datatype_definition_integerFormCallback.probe.formStage.Checkout()

	if datatype_definition_integerFormCallback.datatype_definition_integer == nil {
		datatype_definition_integerFormCallback.datatype_definition_integer = new(models.DATATYPE_DEFINITION_INTEGER).Stage(datatype_definition_integerFormCallback.probe.stageOfInterest)
	}
	datatype_definition_integer_ := datatype_definition_integerFormCallback.datatype_definition_integer
	_ = datatype_definition_integer_

	for _, formDiv := range datatype_definition_integerFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(datatype_definition_integer_.Name), formDiv)
		case "DESC":
			FormDivBasicFieldToField(&(datatype_definition_integer_.DESC), formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&(datatype_definition_integer_.IDENTIFIER), formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&(datatype_definition_integer_.LAST_CHANGE), formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&(datatype_definition_integer_.LONG_NAME), formDiv)
		case "MAX":
			FormDivBasicFieldToField(&(datatype_definition_integer_.MAX), formDiv)
		case "MIN":
			FormDivBasicFieldToField(&(datatype_definition_integer_.MIN), formDiv)
		case "ALTERNATIVE_ID":
			FormDivSelectFieldToField(&(datatype_definition_integer_.ALTERNATIVE_ID), datatype_definition_integerFormCallback.probe.stageOfInterest, formDiv)
		case "A_DATATYPES:DATATYPE_DEFINITION_INTEGER":
			// WARNING : this form deals with the N-N association "A_DATATYPES.DATATYPE_DEFINITION_INTEGER []*DATATYPE_DEFINITION_INTEGER" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of DATATYPE_DEFINITION_INTEGER). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.A_DATATYPES
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "A_DATATYPES"
				rf.Fieldname = "DATATYPE_DEFINITION_INTEGER"
				formerAssociationSource := datatype_definition_integer_.GongGetReverseFieldOwner(
					datatype_definition_integerFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.A_DATATYPES)
					if !ok {
						log.Fatalln("Source of A_DATATYPES.DATATYPE_DEFINITION_INTEGER []*DATATYPE_DEFINITION_INTEGER, is not an A_DATATYPES instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.DATATYPE_DEFINITION_INTEGER, datatype_definition_integer_)
					formerSource.DATATYPE_DEFINITION_INTEGER = slices.Delete(formerSource.DATATYPE_DEFINITION_INTEGER, idx, idx+1)
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
			var newSource *models.A_DATATYPES
			for _a_datatypes := range *models.GetGongstructInstancesSet[models.A_DATATYPES](datatype_definition_integerFormCallback.probe.stageOfInterest) {
				// the match is base on the name
				if _a_datatypes.GetName() == newSourceName.GetName() {
					newSource = _a_datatypes // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of A_DATATYPES.DATATYPE_DEFINITION_INTEGER []*DATATYPE_DEFINITION_INTEGER, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.DATATYPE_DEFINITION_INTEGER = append(newSource.DATATYPE_DEFINITION_INTEGER, datatype_definition_integer_)
		}
	}

	// manage the suppress operation
	if datatype_definition_integerFormCallback.formGroup.HasSuppressButtonBeenPressed {
		datatype_definition_integer_.Unstage(datatype_definition_integerFormCallback.probe.stageOfInterest)
	}

	datatype_definition_integerFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.DATATYPE_DEFINITION_INTEGER](
		datatype_definition_integerFormCallback.probe,
	)

	// display a new form by reset the form stage
	if datatype_definition_integerFormCallback.CreationMode || datatype_definition_integerFormCallback.formGroup.HasSuppressButtonBeenPressed {
		datatype_definition_integerFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(datatype_definition_integerFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__DATATYPE_DEFINITION_INTEGERFormCallback(
			nil,
			datatype_definition_integerFormCallback.probe,
			newFormGroup,
		)
		datatype_definition_integer := new(models.DATATYPE_DEFINITION_INTEGER)
		FillUpForm(datatype_definition_integer, newFormGroup, datatype_definition_integerFormCallback.probe)
		datatype_definition_integerFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(datatype_definition_integerFormCallback.probe)
}

func __gong__New__DATATYPE_DEFINITION_REALFormCallback(
	datatype_definition_real *models.DATATYPE_DEFINITION_REAL,
	probe *Probe,
	formGroup *table.FormGroup,
) (datatype_definition_realFormCallback *DATATYPE_DEFINITION_REALFormCallback) {
	datatype_definition_realFormCallback = new(DATATYPE_DEFINITION_REALFormCallback)
	datatype_definition_realFormCallback.probe = probe
	datatype_definition_realFormCallback.datatype_definition_real = datatype_definition_real
	datatype_definition_realFormCallback.formGroup = formGroup

	datatype_definition_realFormCallback.CreationMode = (datatype_definition_real == nil)

	return
}

type DATATYPE_DEFINITION_REALFormCallback struct {
	datatype_definition_real *models.DATATYPE_DEFINITION_REAL

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (datatype_definition_realFormCallback *DATATYPE_DEFINITION_REALFormCallback) OnSave() {
	datatype_definition_realFormCallback.probe.stageOfInterest.Lock()
	defer datatype_definition_realFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("DATATYPE_DEFINITION_REALFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	datatype_definition_realFormCallback.probe.formStage.Checkout()

	if datatype_definition_realFormCallback.datatype_definition_real == nil {
		datatype_definition_realFormCallback.datatype_definition_real = new(models.DATATYPE_DEFINITION_REAL).Stage(datatype_definition_realFormCallback.probe.stageOfInterest)
	}
	datatype_definition_real_ := datatype_definition_realFormCallback.datatype_definition_real
	_ = datatype_definition_real_

	for _, formDiv := range datatype_definition_realFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(datatype_definition_real_.Name), formDiv)
		case "ACCURACY":
			FormDivBasicFieldToField(&(datatype_definition_real_.ACCURACY), formDiv)
		case "DESC":
			FormDivBasicFieldToField(&(datatype_definition_real_.DESC), formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&(datatype_definition_real_.IDENTIFIER), formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&(datatype_definition_real_.LAST_CHANGE), formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&(datatype_definition_real_.LONG_NAME), formDiv)
		case "MAX":
			FormDivBasicFieldToField(&(datatype_definition_real_.MAX), formDiv)
		case "MIN":
			FormDivBasicFieldToField(&(datatype_definition_real_.MIN), formDiv)
		case "ALTERNATIVE_ID":
			FormDivSelectFieldToField(&(datatype_definition_real_.ALTERNATIVE_ID), datatype_definition_realFormCallback.probe.stageOfInterest, formDiv)
		case "A_DATATYPES:DATATYPE_DEFINITION_REAL":
			// WARNING : this form deals with the N-N association "A_DATATYPES.DATATYPE_DEFINITION_REAL []*DATATYPE_DEFINITION_REAL" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of DATATYPE_DEFINITION_REAL). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.A_DATATYPES
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "A_DATATYPES"
				rf.Fieldname = "DATATYPE_DEFINITION_REAL"
				formerAssociationSource := datatype_definition_real_.GongGetReverseFieldOwner(
					datatype_definition_realFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.A_DATATYPES)
					if !ok {
						log.Fatalln("Source of A_DATATYPES.DATATYPE_DEFINITION_REAL []*DATATYPE_DEFINITION_REAL, is not an A_DATATYPES instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.DATATYPE_DEFINITION_REAL, datatype_definition_real_)
					formerSource.DATATYPE_DEFINITION_REAL = slices.Delete(formerSource.DATATYPE_DEFINITION_REAL, idx, idx+1)
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
			var newSource *models.A_DATATYPES
			for _a_datatypes := range *models.GetGongstructInstancesSet[models.A_DATATYPES](datatype_definition_realFormCallback.probe.stageOfInterest) {
				// the match is base on the name
				if _a_datatypes.GetName() == newSourceName.GetName() {
					newSource = _a_datatypes // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of A_DATATYPES.DATATYPE_DEFINITION_REAL []*DATATYPE_DEFINITION_REAL, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.DATATYPE_DEFINITION_REAL = append(newSource.DATATYPE_DEFINITION_REAL, datatype_definition_real_)
		}
	}

	// manage the suppress operation
	if datatype_definition_realFormCallback.formGroup.HasSuppressButtonBeenPressed {
		datatype_definition_real_.Unstage(datatype_definition_realFormCallback.probe.stageOfInterest)
	}

	datatype_definition_realFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.DATATYPE_DEFINITION_REAL](
		datatype_definition_realFormCallback.probe,
	)

	// display a new form by reset the form stage
	if datatype_definition_realFormCallback.CreationMode || datatype_definition_realFormCallback.formGroup.HasSuppressButtonBeenPressed {
		datatype_definition_realFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(datatype_definition_realFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__DATATYPE_DEFINITION_REALFormCallback(
			nil,
			datatype_definition_realFormCallback.probe,
			newFormGroup,
		)
		datatype_definition_real := new(models.DATATYPE_DEFINITION_REAL)
		FillUpForm(datatype_definition_real, newFormGroup, datatype_definition_realFormCallback.probe)
		datatype_definition_realFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(datatype_definition_realFormCallback.probe)
}

func __gong__New__DATATYPE_DEFINITION_STRINGFormCallback(
	datatype_definition_string *models.DATATYPE_DEFINITION_STRING,
	probe *Probe,
	formGroup *table.FormGroup,
) (datatype_definition_stringFormCallback *DATATYPE_DEFINITION_STRINGFormCallback) {
	datatype_definition_stringFormCallback = new(DATATYPE_DEFINITION_STRINGFormCallback)
	datatype_definition_stringFormCallback.probe = probe
	datatype_definition_stringFormCallback.datatype_definition_string = datatype_definition_string
	datatype_definition_stringFormCallback.formGroup = formGroup

	datatype_definition_stringFormCallback.CreationMode = (datatype_definition_string == nil)

	return
}

type DATATYPE_DEFINITION_STRINGFormCallback struct {
	datatype_definition_string *models.DATATYPE_DEFINITION_STRING

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (datatype_definition_stringFormCallback *DATATYPE_DEFINITION_STRINGFormCallback) OnSave() {
	datatype_definition_stringFormCallback.probe.stageOfInterest.Lock()
	defer datatype_definition_stringFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("DATATYPE_DEFINITION_STRINGFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	datatype_definition_stringFormCallback.probe.formStage.Checkout()

	if datatype_definition_stringFormCallback.datatype_definition_string == nil {
		datatype_definition_stringFormCallback.datatype_definition_string = new(models.DATATYPE_DEFINITION_STRING).Stage(datatype_definition_stringFormCallback.probe.stageOfInterest)
	}
	datatype_definition_string_ := datatype_definition_stringFormCallback.datatype_definition_string
	_ = datatype_definition_string_

	for _, formDiv := range datatype_definition_stringFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(datatype_definition_string_.Name), formDiv)
		case "DESC":
			FormDivBasicFieldToField(&(datatype_definition_string_.DESC), formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&(datatype_definition_string_.IDENTIFIER), formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&(datatype_definition_string_.LAST_CHANGE), formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&(datatype_definition_string_.LONG_NAME), formDiv)
		case "MAX_LENGTH":
			FormDivBasicFieldToField(&(datatype_definition_string_.MAX_LENGTH), formDiv)
		case "ALTERNATIVE_ID":
			FormDivSelectFieldToField(&(datatype_definition_string_.ALTERNATIVE_ID), datatype_definition_stringFormCallback.probe.stageOfInterest, formDiv)
		case "A_DATATYPES:DATATYPE_DEFINITION_STRING":
			// WARNING : this form deals with the N-N association "A_DATATYPES.DATATYPE_DEFINITION_STRING []*DATATYPE_DEFINITION_STRING" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of DATATYPE_DEFINITION_STRING). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.A_DATATYPES
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "A_DATATYPES"
				rf.Fieldname = "DATATYPE_DEFINITION_STRING"
				formerAssociationSource := datatype_definition_string_.GongGetReverseFieldOwner(
					datatype_definition_stringFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.A_DATATYPES)
					if !ok {
						log.Fatalln("Source of A_DATATYPES.DATATYPE_DEFINITION_STRING []*DATATYPE_DEFINITION_STRING, is not an A_DATATYPES instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.DATATYPE_DEFINITION_STRING, datatype_definition_string_)
					formerSource.DATATYPE_DEFINITION_STRING = slices.Delete(formerSource.DATATYPE_DEFINITION_STRING, idx, idx+1)
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
			var newSource *models.A_DATATYPES
			for _a_datatypes := range *models.GetGongstructInstancesSet[models.A_DATATYPES](datatype_definition_stringFormCallback.probe.stageOfInterest) {
				// the match is base on the name
				if _a_datatypes.GetName() == newSourceName.GetName() {
					newSource = _a_datatypes // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of A_DATATYPES.DATATYPE_DEFINITION_STRING []*DATATYPE_DEFINITION_STRING, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.DATATYPE_DEFINITION_STRING = append(newSource.DATATYPE_DEFINITION_STRING, datatype_definition_string_)
		}
	}

	// manage the suppress operation
	if datatype_definition_stringFormCallback.formGroup.HasSuppressButtonBeenPressed {
		datatype_definition_string_.Unstage(datatype_definition_stringFormCallback.probe.stageOfInterest)
	}

	datatype_definition_stringFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.DATATYPE_DEFINITION_STRING](
		datatype_definition_stringFormCallback.probe,
	)

	// display a new form by reset the form stage
	if datatype_definition_stringFormCallback.CreationMode || datatype_definition_stringFormCallback.formGroup.HasSuppressButtonBeenPressed {
		datatype_definition_stringFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(datatype_definition_stringFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__DATATYPE_DEFINITION_STRINGFormCallback(
			nil,
			datatype_definition_stringFormCallback.probe,
			newFormGroup,
		)
		datatype_definition_string := new(models.DATATYPE_DEFINITION_STRING)
		FillUpForm(datatype_definition_string, newFormGroup, datatype_definition_stringFormCallback.probe)
		datatype_definition_stringFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(datatype_definition_stringFormCallback.probe)
}

func __gong__New__DATATYPE_DEFINITION_XHTMLFormCallback(
	datatype_definition_xhtml *models.DATATYPE_DEFINITION_XHTML,
	probe *Probe,
	formGroup *table.FormGroup,
) (datatype_definition_xhtmlFormCallback *DATATYPE_DEFINITION_XHTMLFormCallback) {
	datatype_definition_xhtmlFormCallback = new(DATATYPE_DEFINITION_XHTMLFormCallback)
	datatype_definition_xhtmlFormCallback.probe = probe
	datatype_definition_xhtmlFormCallback.datatype_definition_xhtml = datatype_definition_xhtml
	datatype_definition_xhtmlFormCallback.formGroup = formGroup

	datatype_definition_xhtmlFormCallback.CreationMode = (datatype_definition_xhtml == nil)

	return
}

type DATATYPE_DEFINITION_XHTMLFormCallback struct {
	datatype_definition_xhtml *models.DATATYPE_DEFINITION_XHTML

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (datatype_definition_xhtmlFormCallback *DATATYPE_DEFINITION_XHTMLFormCallback) OnSave() {
	datatype_definition_xhtmlFormCallback.probe.stageOfInterest.Lock()
	defer datatype_definition_xhtmlFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("DATATYPE_DEFINITION_XHTMLFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	datatype_definition_xhtmlFormCallback.probe.formStage.Checkout()

	if datatype_definition_xhtmlFormCallback.datatype_definition_xhtml == nil {
		datatype_definition_xhtmlFormCallback.datatype_definition_xhtml = new(models.DATATYPE_DEFINITION_XHTML).Stage(datatype_definition_xhtmlFormCallback.probe.stageOfInterest)
	}
	datatype_definition_xhtml_ := datatype_definition_xhtmlFormCallback.datatype_definition_xhtml
	_ = datatype_definition_xhtml_

	for _, formDiv := range datatype_definition_xhtmlFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(datatype_definition_xhtml_.Name), formDiv)
		case "DESC":
			FormDivBasicFieldToField(&(datatype_definition_xhtml_.DESC), formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&(datatype_definition_xhtml_.IDENTIFIER), formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&(datatype_definition_xhtml_.LAST_CHANGE), formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&(datatype_definition_xhtml_.LONG_NAME), formDiv)
		case "ALTERNATIVE_ID":
			FormDivSelectFieldToField(&(datatype_definition_xhtml_.ALTERNATIVE_ID), datatype_definition_xhtmlFormCallback.probe.stageOfInterest, formDiv)
		case "A_DATATYPES:DATATYPE_DEFINITION_XHTML":
			// WARNING : this form deals with the N-N association "A_DATATYPES.DATATYPE_DEFINITION_XHTML []*DATATYPE_DEFINITION_XHTML" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of DATATYPE_DEFINITION_XHTML). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.A_DATATYPES
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "A_DATATYPES"
				rf.Fieldname = "DATATYPE_DEFINITION_XHTML"
				formerAssociationSource := datatype_definition_xhtml_.GongGetReverseFieldOwner(
					datatype_definition_xhtmlFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.A_DATATYPES)
					if !ok {
						log.Fatalln("Source of A_DATATYPES.DATATYPE_DEFINITION_XHTML []*DATATYPE_DEFINITION_XHTML, is not an A_DATATYPES instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.DATATYPE_DEFINITION_XHTML, datatype_definition_xhtml_)
					formerSource.DATATYPE_DEFINITION_XHTML = slices.Delete(formerSource.DATATYPE_DEFINITION_XHTML, idx, idx+1)
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
			var newSource *models.A_DATATYPES
			for _a_datatypes := range *models.GetGongstructInstancesSet[models.A_DATATYPES](datatype_definition_xhtmlFormCallback.probe.stageOfInterest) {
				// the match is base on the name
				if _a_datatypes.GetName() == newSourceName.GetName() {
					newSource = _a_datatypes // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of A_DATATYPES.DATATYPE_DEFINITION_XHTML []*DATATYPE_DEFINITION_XHTML, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.DATATYPE_DEFINITION_XHTML = append(newSource.DATATYPE_DEFINITION_XHTML, datatype_definition_xhtml_)
		}
	}

	// manage the suppress operation
	if datatype_definition_xhtmlFormCallback.formGroup.HasSuppressButtonBeenPressed {
		datatype_definition_xhtml_.Unstage(datatype_definition_xhtmlFormCallback.probe.stageOfInterest)
	}

	datatype_definition_xhtmlFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.DATATYPE_DEFINITION_XHTML](
		datatype_definition_xhtmlFormCallback.probe,
	)

	// display a new form by reset the form stage
	if datatype_definition_xhtmlFormCallback.CreationMode || datatype_definition_xhtmlFormCallback.formGroup.HasSuppressButtonBeenPressed {
		datatype_definition_xhtmlFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(datatype_definition_xhtmlFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__DATATYPE_DEFINITION_XHTMLFormCallback(
			nil,
			datatype_definition_xhtmlFormCallback.probe,
			newFormGroup,
		)
		datatype_definition_xhtml := new(models.DATATYPE_DEFINITION_XHTML)
		FillUpForm(datatype_definition_xhtml, newFormGroup, datatype_definition_xhtmlFormCallback.probe)
		datatype_definition_xhtmlFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(datatype_definition_xhtmlFormCallback.probe)
}

func __gong__New__EMBEDDED_VALUEFormCallback(
	embedded_value *models.EMBEDDED_VALUE,
	probe *Probe,
	formGroup *table.FormGroup,
) (embedded_valueFormCallback *EMBEDDED_VALUEFormCallback) {
	embedded_valueFormCallback = new(EMBEDDED_VALUEFormCallback)
	embedded_valueFormCallback.probe = probe
	embedded_valueFormCallback.embedded_value = embedded_value
	embedded_valueFormCallback.formGroup = formGroup

	embedded_valueFormCallback.CreationMode = (embedded_value == nil)

	return
}

type EMBEDDED_VALUEFormCallback struct {
	embedded_value *models.EMBEDDED_VALUE

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (embedded_valueFormCallback *EMBEDDED_VALUEFormCallback) OnSave() {
	embedded_valueFormCallback.probe.stageOfInterest.Lock()
	defer embedded_valueFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("EMBEDDED_VALUEFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	embedded_valueFormCallback.probe.formStage.Checkout()

	if embedded_valueFormCallback.embedded_value == nil {
		embedded_valueFormCallback.embedded_value = new(models.EMBEDDED_VALUE).Stage(embedded_valueFormCallback.probe.stageOfInterest)
	}
	embedded_value_ := embedded_valueFormCallback.embedded_value
	_ = embedded_value_

	for _, formDiv := range embedded_valueFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(embedded_value_.Name), formDiv)
		case "KEY":
			FormDivBasicFieldToField(&(embedded_value_.KEY), formDiv)
		case "OTHER_CONTENT":
			FormDivBasicFieldToField(&(embedded_value_.OTHER_CONTENT), formDiv)
		}
	}

	// manage the suppress operation
	if embedded_valueFormCallback.formGroup.HasSuppressButtonBeenPressed {
		embedded_value_.Unstage(embedded_valueFormCallback.probe.stageOfInterest)
	}

	embedded_valueFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.EMBEDDED_VALUE](
		embedded_valueFormCallback.probe,
	)

	// display a new form by reset the form stage
	if embedded_valueFormCallback.CreationMode || embedded_valueFormCallback.formGroup.HasSuppressButtonBeenPressed {
		embedded_valueFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(embedded_valueFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__EMBEDDED_VALUEFormCallback(
			nil,
			embedded_valueFormCallback.probe,
			newFormGroup,
		)
		embedded_value := new(models.EMBEDDED_VALUE)
		FillUpForm(embedded_value, newFormGroup, embedded_valueFormCallback.probe)
		embedded_valueFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(embedded_valueFormCallback.probe)
}

func __gong__New__ENUM_VALUEFormCallback(
	enum_value *models.ENUM_VALUE,
	probe *Probe,
	formGroup *table.FormGroup,
) (enum_valueFormCallback *ENUM_VALUEFormCallback) {
	enum_valueFormCallback = new(ENUM_VALUEFormCallback)
	enum_valueFormCallback.probe = probe
	enum_valueFormCallback.enum_value = enum_value
	enum_valueFormCallback.formGroup = formGroup

	enum_valueFormCallback.CreationMode = (enum_value == nil)

	return
}

type ENUM_VALUEFormCallback struct {
	enum_value *models.ENUM_VALUE

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (enum_valueFormCallback *ENUM_VALUEFormCallback) OnSave() {
	enum_valueFormCallback.probe.stageOfInterest.Lock()
	defer enum_valueFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("ENUM_VALUEFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	enum_valueFormCallback.probe.formStage.Checkout()

	if enum_valueFormCallback.enum_value == nil {
		enum_valueFormCallback.enum_value = new(models.ENUM_VALUE).Stage(enum_valueFormCallback.probe.stageOfInterest)
	}
	enum_value_ := enum_valueFormCallback.enum_value
	_ = enum_value_

	for _, formDiv := range enum_valueFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(enum_value_.Name), formDiv)
		case "DESC":
			FormDivBasicFieldToField(&(enum_value_.DESC), formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&(enum_value_.IDENTIFIER), formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&(enum_value_.LAST_CHANGE), formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&(enum_value_.LONG_NAME), formDiv)
		case "ALTERNATIVE_ID":
			FormDivSelectFieldToField(&(enum_value_.ALTERNATIVE_ID), enum_valueFormCallback.probe.stageOfInterest, formDiv)
		case "PROPERTIES":
			FormDivSelectFieldToField(&(enum_value_.PROPERTIES), enum_valueFormCallback.probe.stageOfInterest, formDiv)
		case "A_SPECIFIED_VALUES:ENUM_VALUE":
			// WARNING : this form deals with the N-N association "A_SPECIFIED_VALUES.ENUM_VALUE []*ENUM_VALUE" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of ENUM_VALUE). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.A_SPECIFIED_VALUES
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "A_SPECIFIED_VALUES"
				rf.Fieldname = "ENUM_VALUE"
				formerAssociationSource := enum_value_.GongGetReverseFieldOwner(
					enum_valueFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.A_SPECIFIED_VALUES)
					if !ok {
						log.Fatalln("Source of A_SPECIFIED_VALUES.ENUM_VALUE []*ENUM_VALUE, is not an A_SPECIFIED_VALUES instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.ENUM_VALUE, enum_value_)
					formerSource.ENUM_VALUE = slices.Delete(formerSource.ENUM_VALUE, idx, idx+1)
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
			var newSource *models.A_SPECIFIED_VALUES
			for _a_specified_values := range *models.GetGongstructInstancesSet[models.A_SPECIFIED_VALUES](enum_valueFormCallback.probe.stageOfInterest) {
				// the match is base on the name
				if _a_specified_values.GetName() == newSourceName.GetName() {
					newSource = _a_specified_values // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of A_SPECIFIED_VALUES.ENUM_VALUE []*ENUM_VALUE, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.ENUM_VALUE = append(newSource.ENUM_VALUE, enum_value_)
		}
	}

	// manage the suppress operation
	if enum_valueFormCallback.formGroup.HasSuppressButtonBeenPressed {
		enum_value_.Unstage(enum_valueFormCallback.probe.stageOfInterest)
	}

	enum_valueFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.ENUM_VALUE](
		enum_valueFormCallback.probe,
	)

	// display a new form by reset the form stage
	if enum_valueFormCallback.CreationMode || enum_valueFormCallback.formGroup.HasSuppressButtonBeenPressed {
		enum_valueFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(enum_valueFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ENUM_VALUEFormCallback(
			nil,
			enum_valueFormCallback.probe,
			newFormGroup,
		)
		enum_value := new(models.ENUM_VALUE)
		FillUpForm(enum_value, newFormGroup, enum_valueFormCallback.probe)
		enum_valueFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(enum_valueFormCallback.probe)
}

func __gong__New__RELATION_GROUPFormCallback(
	relation_group *models.RELATION_GROUP,
	probe *Probe,
	formGroup *table.FormGroup,
) (relation_groupFormCallback *RELATION_GROUPFormCallback) {
	relation_groupFormCallback = new(RELATION_GROUPFormCallback)
	relation_groupFormCallback.probe = probe
	relation_groupFormCallback.relation_group = relation_group
	relation_groupFormCallback.formGroup = formGroup

	relation_groupFormCallback.CreationMode = (relation_group == nil)

	return
}

type RELATION_GROUPFormCallback struct {
	relation_group *models.RELATION_GROUP

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (relation_groupFormCallback *RELATION_GROUPFormCallback) OnSave() {
	relation_groupFormCallback.probe.stageOfInterest.Lock()
	defer relation_groupFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("RELATION_GROUPFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	relation_groupFormCallback.probe.formStage.Checkout()

	if relation_groupFormCallback.relation_group == nil {
		relation_groupFormCallback.relation_group = new(models.RELATION_GROUP).Stage(relation_groupFormCallback.probe.stageOfInterest)
	}
	relation_group_ := relation_groupFormCallback.relation_group
	_ = relation_group_

	for _, formDiv := range relation_groupFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(relation_group_.Name), formDiv)
		case "DESC":
			FormDivBasicFieldToField(&(relation_group_.DESC), formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&(relation_group_.IDENTIFIER), formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&(relation_group_.LAST_CHANGE), formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&(relation_group_.LONG_NAME), formDiv)
		case "ALTERNATIVE_ID":
			FormDivSelectFieldToField(&(relation_group_.ALTERNATIVE_ID), relation_groupFormCallback.probe.stageOfInterest, formDiv)
		case "SOURCE_SPECIFICATION":
			FormDivSelectFieldToField(&(relation_group_.SOURCE_SPECIFICATION), relation_groupFormCallback.probe.stageOfInterest, formDiv)
		case "SPEC_RELATIONS":
			FormDivSelectFieldToField(&(relation_group_.SPEC_RELATIONS), relation_groupFormCallback.probe.stageOfInterest, formDiv)
		case "TARGET_SPECIFICATION":
			FormDivSelectFieldToField(&(relation_group_.TARGET_SPECIFICATION), relation_groupFormCallback.probe.stageOfInterest, formDiv)
		case "TYPE":
			FormDivSelectFieldToField(&(relation_group_.TYPE), relation_groupFormCallback.probe.stageOfInterest, formDiv)
		case "A_SPEC_RELATION_GROUPS:RELATION_GROUP":
			// WARNING : this form deals with the N-N association "A_SPEC_RELATION_GROUPS.RELATION_GROUP []*RELATION_GROUP" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of RELATION_GROUP). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.A_SPEC_RELATION_GROUPS
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "A_SPEC_RELATION_GROUPS"
				rf.Fieldname = "RELATION_GROUP"
				formerAssociationSource := relation_group_.GongGetReverseFieldOwner(
					relation_groupFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.A_SPEC_RELATION_GROUPS)
					if !ok {
						log.Fatalln("Source of A_SPEC_RELATION_GROUPS.RELATION_GROUP []*RELATION_GROUP, is not an A_SPEC_RELATION_GROUPS instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.RELATION_GROUP, relation_group_)
					formerSource.RELATION_GROUP = slices.Delete(formerSource.RELATION_GROUP, idx, idx+1)
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
			var newSource *models.A_SPEC_RELATION_GROUPS
			for _a_spec_relation_groups := range *models.GetGongstructInstancesSet[models.A_SPEC_RELATION_GROUPS](relation_groupFormCallback.probe.stageOfInterest) {
				// the match is base on the name
				if _a_spec_relation_groups.GetName() == newSourceName.GetName() {
					newSource = _a_spec_relation_groups // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of A_SPEC_RELATION_GROUPS.RELATION_GROUP []*RELATION_GROUP, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.RELATION_GROUP = append(newSource.RELATION_GROUP, relation_group_)
		}
	}

	// manage the suppress operation
	if relation_groupFormCallback.formGroup.HasSuppressButtonBeenPressed {
		relation_group_.Unstage(relation_groupFormCallback.probe.stageOfInterest)
	}

	relation_groupFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.RELATION_GROUP](
		relation_groupFormCallback.probe,
	)

	// display a new form by reset the form stage
	if relation_groupFormCallback.CreationMode || relation_groupFormCallback.formGroup.HasSuppressButtonBeenPressed {
		relation_groupFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(relation_groupFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__RELATION_GROUPFormCallback(
			nil,
			relation_groupFormCallback.probe,
			newFormGroup,
		)
		relation_group := new(models.RELATION_GROUP)
		FillUpForm(relation_group, newFormGroup, relation_groupFormCallback.probe)
		relation_groupFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(relation_groupFormCallback.probe)
}

func __gong__New__RELATION_GROUP_TYPEFormCallback(
	relation_group_type *models.RELATION_GROUP_TYPE,
	probe *Probe,
	formGroup *table.FormGroup,
) (relation_group_typeFormCallback *RELATION_GROUP_TYPEFormCallback) {
	relation_group_typeFormCallback = new(RELATION_GROUP_TYPEFormCallback)
	relation_group_typeFormCallback.probe = probe
	relation_group_typeFormCallback.relation_group_type = relation_group_type
	relation_group_typeFormCallback.formGroup = formGroup

	relation_group_typeFormCallback.CreationMode = (relation_group_type == nil)

	return
}

type RELATION_GROUP_TYPEFormCallback struct {
	relation_group_type *models.RELATION_GROUP_TYPE

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (relation_group_typeFormCallback *RELATION_GROUP_TYPEFormCallback) OnSave() {
	relation_group_typeFormCallback.probe.stageOfInterest.Lock()
	defer relation_group_typeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("RELATION_GROUP_TYPEFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	relation_group_typeFormCallback.probe.formStage.Checkout()

	if relation_group_typeFormCallback.relation_group_type == nil {
		relation_group_typeFormCallback.relation_group_type = new(models.RELATION_GROUP_TYPE).Stage(relation_group_typeFormCallback.probe.stageOfInterest)
	}
	relation_group_type_ := relation_group_typeFormCallback.relation_group_type
	_ = relation_group_type_

	for _, formDiv := range relation_group_typeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(relation_group_type_.Name), formDiv)
		case "DESC":
			FormDivBasicFieldToField(&(relation_group_type_.DESC), formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&(relation_group_type_.IDENTIFIER), formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&(relation_group_type_.LAST_CHANGE), formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&(relation_group_type_.LONG_NAME), formDiv)
		case "ALTERNATIVE_ID":
			FormDivSelectFieldToField(&(relation_group_type_.ALTERNATIVE_ID), relation_group_typeFormCallback.probe.stageOfInterest, formDiv)
		case "SPEC_ATTRIBUTES":
			FormDivSelectFieldToField(&(relation_group_type_.SPEC_ATTRIBUTES), relation_group_typeFormCallback.probe.stageOfInterest, formDiv)
		case "A_SPEC_TYPES:RELATION_GROUP_TYPE":
			// WARNING : this form deals with the N-N association "A_SPEC_TYPES.RELATION_GROUP_TYPE []*RELATION_GROUP_TYPE" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of RELATION_GROUP_TYPE). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.A_SPEC_TYPES
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "A_SPEC_TYPES"
				rf.Fieldname = "RELATION_GROUP_TYPE"
				formerAssociationSource := relation_group_type_.GongGetReverseFieldOwner(
					relation_group_typeFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.A_SPEC_TYPES)
					if !ok {
						log.Fatalln("Source of A_SPEC_TYPES.RELATION_GROUP_TYPE []*RELATION_GROUP_TYPE, is not an A_SPEC_TYPES instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.RELATION_GROUP_TYPE, relation_group_type_)
					formerSource.RELATION_GROUP_TYPE = slices.Delete(formerSource.RELATION_GROUP_TYPE, idx, idx+1)
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
			var newSource *models.A_SPEC_TYPES
			for _a_spec_types := range *models.GetGongstructInstancesSet[models.A_SPEC_TYPES](relation_group_typeFormCallback.probe.stageOfInterest) {
				// the match is base on the name
				if _a_spec_types.GetName() == newSourceName.GetName() {
					newSource = _a_spec_types // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of A_SPEC_TYPES.RELATION_GROUP_TYPE []*RELATION_GROUP_TYPE, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.RELATION_GROUP_TYPE = append(newSource.RELATION_GROUP_TYPE, relation_group_type_)
		}
	}

	// manage the suppress operation
	if relation_group_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		relation_group_type_.Unstage(relation_group_typeFormCallback.probe.stageOfInterest)
	}

	relation_group_typeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.RELATION_GROUP_TYPE](
		relation_group_typeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if relation_group_typeFormCallback.CreationMode || relation_group_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		relation_group_typeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(relation_group_typeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__RELATION_GROUP_TYPEFormCallback(
			nil,
			relation_group_typeFormCallback.probe,
			newFormGroup,
		)
		relation_group_type := new(models.RELATION_GROUP_TYPE)
		FillUpForm(relation_group_type, newFormGroup, relation_group_typeFormCallback.probe)
		relation_group_typeFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(relation_group_typeFormCallback.probe)
}

func __gong__New__REQ_IFFormCallback(
	req_if *models.REQ_IF,
	probe *Probe,
	formGroup *table.FormGroup,
) (req_ifFormCallback *REQ_IFFormCallback) {
	req_ifFormCallback = new(REQ_IFFormCallback)
	req_ifFormCallback.probe = probe
	req_ifFormCallback.req_if = req_if
	req_ifFormCallback.formGroup = formGroup

	req_ifFormCallback.CreationMode = (req_if == nil)

	return
}

type REQ_IFFormCallback struct {
	req_if *models.REQ_IF

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (req_ifFormCallback *REQ_IFFormCallback) OnSave() {
	req_ifFormCallback.probe.stageOfInterest.Lock()
	defer req_ifFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("REQ_IFFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	req_ifFormCallback.probe.formStage.Checkout()

	if req_ifFormCallback.req_if == nil {
		req_ifFormCallback.req_if = new(models.REQ_IF).Stage(req_ifFormCallback.probe.stageOfInterest)
	}
	req_if_ := req_ifFormCallback.req_if
	_ = req_if_

	for _, formDiv := range req_ifFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(req_if_.Name), formDiv)
		case "Lang":
			FormDivBasicFieldToField(&(req_if_.Lang), formDiv)
		case "THE_HEADER":
			FormDivSelectFieldToField(&(req_if_.THE_HEADER), req_ifFormCallback.probe.stageOfInterest, formDiv)
		case "CORE_CONTENT":
			FormDivSelectFieldToField(&(req_if_.CORE_CONTENT), req_ifFormCallback.probe.stageOfInterest, formDiv)
		case "TOOL_EXTENSIONS":
			FormDivSelectFieldToField(&(req_if_.TOOL_EXTENSIONS), req_ifFormCallback.probe.stageOfInterest, formDiv)
		}
	}

	// manage the suppress operation
	if req_ifFormCallback.formGroup.HasSuppressButtonBeenPressed {
		req_if_.Unstage(req_ifFormCallback.probe.stageOfInterest)
	}

	req_ifFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.REQ_IF](
		req_ifFormCallback.probe,
	)

	// display a new form by reset the form stage
	if req_ifFormCallback.CreationMode || req_ifFormCallback.formGroup.HasSuppressButtonBeenPressed {
		req_ifFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(req_ifFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__REQ_IFFormCallback(
			nil,
			req_ifFormCallback.probe,
			newFormGroup,
		)
		req_if := new(models.REQ_IF)
		FillUpForm(req_if, newFormGroup, req_ifFormCallback.probe)
		req_ifFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(req_ifFormCallback.probe)
}

func __gong__New__REQ_IF_CONTENTFormCallback(
	req_if_content *models.REQ_IF_CONTENT,
	probe *Probe,
	formGroup *table.FormGroup,
) (req_if_contentFormCallback *REQ_IF_CONTENTFormCallback) {
	req_if_contentFormCallback = new(REQ_IF_CONTENTFormCallback)
	req_if_contentFormCallback.probe = probe
	req_if_contentFormCallback.req_if_content = req_if_content
	req_if_contentFormCallback.formGroup = formGroup

	req_if_contentFormCallback.CreationMode = (req_if_content == nil)

	return
}

type REQ_IF_CONTENTFormCallback struct {
	req_if_content *models.REQ_IF_CONTENT

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (req_if_contentFormCallback *REQ_IF_CONTENTFormCallback) OnSave() {
	req_if_contentFormCallback.probe.stageOfInterest.Lock()
	defer req_if_contentFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("REQ_IF_CONTENTFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	req_if_contentFormCallback.probe.formStage.Checkout()

	if req_if_contentFormCallback.req_if_content == nil {
		req_if_contentFormCallback.req_if_content = new(models.REQ_IF_CONTENT).Stage(req_if_contentFormCallback.probe.stageOfInterest)
	}
	req_if_content_ := req_if_contentFormCallback.req_if_content
	_ = req_if_content_

	for _, formDiv := range req_if_contentFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(req_if_content_.Name), formDiv)
		case "DATATYPES":
			FormDivSelectFieldToField(&(req_if_content_.DATATYPES), req_if_contentFormCallback.probe.stageOfInterest, formDiv)
		case "SPEC_TYPES":
			FormDivSelectFieldToField(&(req_if_content_.SPEC_TYPES), req_if_contentFormCallback.probe.stageOfInterest, formDiv)
		case "SPEC_OBJECTS":
			FormDivSelectFieldToField(&(req_if_content_.SPEC_OBJECTS), req_if_contentFormCallback.probe.stageOfInterest, formDiv)
		case "SPEC_RELATIONS":
			FormDivSelectFieldToField(&(req_if_content_.SPEC_RELATIONS), req_if_contentFormCallback.probe.stageOfInterest, formDiv)
		case "SPECIFICATIONS":
			FormDivSelectFieldToField(&(req_if_content_.SPECIFICATIONS), req_if_contentFormCallback.probe.stageOfInterest, formDiv)
		case "SPEC_RELATION_GROUPS":
			FormDivSelectFieldToField(&(req_if_content_.SPEC_RELATION_GROUPS), req_if_contentFormCallback.probe.stageOfInterest, formDiv)
		}
	}

	// manage the suppress operation
	if req_if_contentFormCallback.formGroup.HasSuppressButtonBeenPressed {
		req_if_content_.Unstage(req_if_contentFormCallback.probe.stageOfInterest)
	}

	req_if_contentFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.REQ_IF_CONTENT](
		req_if_contentFormCallback.probe,
	)

	// display a new form by reset the form stage
	if req_if_contentFormCallback.CreationMode || req_if_contentFormCallback.formGroup.HasSuppressButtonBeenPressed {
		req_if_contentFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(req_if_contentFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__REQ_IF_CONTENTFormCallback(
			nil,
			req_if_contentFormCallback.probe,
			newFormGroup,
		)
		req_if_content := new(models.REQ_IF_CONTENT)
		FillUpForm(req_if_content, newFormGroup, req_if_contentFormCallback.probe)
		req_if_contentFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(req_if_contentFormCallback.probe)
}

func __gong__New__REQ_IF_HEADERFormCallback(
	req_if_header *models.REQ_IF_HEADER,
	probe *Probe,
	formGroup *table.FormGroup,
) (req_if_headerFormCallback *REQ_IF_HEADERFormCallback) {
	req_if_headerFormCallback = new(REQ_IF_HEADERFormCallback)
	req_if_headerFormCallback.probe = probe
	req_if_headerFormCallback.req_if_header = req_if_header
	req_if_headerFormCallback.formGroup = formGroup

	req_if_headerFormCallback.CreationMode = (req_if_header == nil)

	return
}

type REQ_IF_HEADERFormCallback struct {
	req_if_header *models.REQ_IF_HEADER

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (req_if_headerFormCallback *REQ_IF_HEADERFormCallback) OnSave() {
	req_if_headerFormCallback.probe.stageOfInterest.Lock()
	defer req_if_headerFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("REQ_IF_HEADERFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	req_if_headerFormCallback.probe.formStage.Checkout()

	if req_if_headerFormCallback.req_if_header == nil {
		req_if_headerFormCallback.req_if_header = new(models.REQ_IF_HEADER).Stage(req_if_headerFormCallback.probe.stageOfInterest)
	}
	req_if_header_ := req_if_headerFormCallback.req_if_header
	_ = req_if_header_

	for _, formDiv := range req_if_headerFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(req_if_header_.Name), formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&(req_if_header_.IDENTIFIER), formDiv)
		case "COMMENT":
			FormDivBasicFieldToField(&(req_if_header_.COMMENT), formDiv)
		case "CREATION_TIME":
			FormDivBasicFieldToField(&(req_if_header_.CREATION_TIME), formDiv)
		case "REPOSITORY_ID":
			FormDivBasicFieldToField(&(req_if_header_.REPOSITORY_ID), formDiv)
		case "REQ_IF_TOOL_ID":
			FormDivBasicFieldToField(&(req_if_header_.REQ_IF_TOOL_ID), formDiv)
		case "REQ_IF_VERSION":
			FormDivBasicFieldToField(&(req_if_header_.REQ_IF_VERSION), formDiv)
		case "SOURCE_TOOL_ID":
			FormDivBasicFieldToField(&(req_if_header_.SOURCE_TOOL_ID), formDiv)
		case "TITLE":
			FormDivBasicFieldToField(&(req_if_header_.TITLE), formDiv)
		}
	}

	// manage the suppress operation
	if req_if_headerFormCallback.formGroup.HasSuppressButtonBeenPressed {
		req_if_header_.Unstage(req_if_headerFormCallback.probe.stageOfInterest)
	}

	req_if_headerFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.REQ_IF_HEADER](
		req_if_headerFormCallback.probe,
	)

	// display a new form by reset the form stage
	if req_if_headerFormCallback.CreationMode || req_if_headerFormCallback.formGroup.HasSuppressButtonBeenPressed {
		req_if_headerFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(req_if_headerFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__REQ_IF_HEADERFormCallback(
			nil,
			req_if_headerFormCallback.probe,
			newFormGroup,
		)
		req_if_header := new(models.REQ_IF_HEADER)
		FillUpForm(req_if_header, newFormGroup, req_if_headerFormCallback.probe)
		req_if_headerFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(req_if_headerFormCallback.probe)
}

func __gong__New__REQ_IF_TOOL_EXTENSIONFormCallback(
	req_if_tool_extension *models.REQ_IF_TOOL_EXTENSION,
	probe *Probe,
	formGroup *table.FormGroup,
) (req_if_tool_extensionFormCallback *REQ_IF_TOOL_EXTENSIONFormCallback) {
	req_if_tool_extensionFormCallback = new(REQ_IF_TOOL_EXTENSIONFormCallback)
	req_if_tool_extensionFormCallback.probe = probe
	req_if_tool_extensionFormCallback.req_if_tool_extension = req_if_tool_extension
	req_if_tool_extensionFormCallback.formGroup = formGroup

	req_if_tool_extensionFormCallback.CreationMode = (req_if_tool_extension == nil)

	return
}

type REQ_IF_TOOL_EXTENSIONFormCallback struct {
	req_if_tool_extension *models.REQ_IF_TOOL_EXTENSION

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (req_if_tool_extensionFormCallback *REQ_IF_TOOL_EXTENSIONFormCallback) OnSave() {
	req_if_tool_extensionFormCallback.probe.stageOfInterest.Lock()
	defer req_if_tool_extensionFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("REQ_IF_TOOL_EXTENSIONFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	req_if_tool_extensionFormCallback.probe.formStage.Checkout()

	if req_if_tool_extensionFormCallback.req_if_tool_extension == nil {
		req_if_tool_extensionFormCallback.req_if_tool_extension = new(models.REQ_IF_TOOL_EXTENSION).Stage(req_if_tool_extensionFormCallback.probe.stageOfInterest)
	}
	req_if_tool_extension_ := req_if_tool_extensionFormCallback.req_if_tool_extension
	_ = req_if_tool_extension_

	for _, formDiv := range req_if_tool_extensionFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(req_if_tool_extension_.Name), formDiv)
		case "A_TOOL_EXTENSIONS:REQ_IF_TOOL_EXTENSION":
			// WARNING : this form deals with the N-N association "A_TOOL_EXTENSIONS.REQ_IF_TOOL_EXTENSION []*REQ_IF_TOOL_EXTENSION" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of REQ_IF_TOOL_EXTENSION). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.A_TOOL_EXTENSIONS
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "A_TOOL_EXTENSIONS"
				rf.Fieldname = "REQ_IF_TOOL_EXTENSION"
				formerAssociationSource := req_if_tool_extension_.GongGetReverseFieldOwner(
					req_if_tool_extensionFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.A_TOOL_EXTENSIONS)
					if !ok {
						log.Fatalln("Source of A_TOOL_EXTENSIONS.REQ_IF_TOOL_EXTENSION []*REQ_IF_TOOL_EXTENSION, is not an A_TOOL_EXTENSIONS instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.REQ_IF_TOOL_EXTENSION, req_if_tool_extension_)
					formerSource.REQ_IF_TOOL_EXTENSION = slices.Delete(formerSource.REQ_IF_TOOL_EXTENSION, idx, idx+1)
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
			var newSource *models.A_TOOL_EXTENSIONS
			for _a_tool_extensions := range *models.GetGongstructInstancesSet[models.A_TOOL_EXTENSIONS](req_if_tool_extensionFormCallback.probe.stageOfInterest) {
				// the match is base on the name
				if _a_tool_extensions.GetName() == newSourceName.GetName() {
					newSource = _a_tool_extensions // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of A_TOOL_EXTENSIONS.REQ_IF_TOOL_EXTENSION []*REQ_IF_TOOL_EXTENSION, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.REQ_IF_TOOL_EXTENSION = append(newSource.REQ_IF_TOOL_EXTENSION, req_if_tool_extension_)
		}
	}

	// manage the suppress operation
	if req_if_tool_extensionFormCallback.formGroup.HasSuppressButtonBeenPressed {
		req_if_tool_extension_.Unstage(req_if_tool_extensionFormCallback.probe.stageOfInterest)
	}

	req_if_tool_extensionFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.REQ_IF_TOOL_EXTENSION](
		req_if_tool_extensionFormCallback.probe,
	)

	// display a new form by reset the form stage
	if req_if_tool_extensionFormCallback.CreationMode || req_if_tool_extensionFormCallback.formGroup.HasSuppressButtonBeenPressed {
		req_if_tool_extensionFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(req_if_tool_extensionFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__REQ_IF_TOOL_EXTENSIONFormCallback(
			nil,
			req_if_tool_extensionFormCallback.probe,
			newFormGroup,
		)
		req_if_tool_extension := new(models.REQ_IF_TOOL_EXTENSION)
		FillUpForm(req_if_tool_extension, newFormGroup, req_if_tool_extensionFormCallback.probe)
		req_if_tool_extensionFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(req_if_tool_extensionFormCallback.probe)
}

func __gong__New__SPECIFICATIONFormCallback(
	specification *models.SPECIFICATION,
	probe *Probe,
	formGroup *table.FormGroup,
) (specificationFormCallback *SPECIFICATIONFormCallback) {
	specificationFormCallback = new(SPECIFICATIONFormCallback)
	specificationFormCallback.probe = probe
	specificationFormCallback.specification = specification
	specificationFormCallback.formGroup = formGroup

	specificationFormCallback.CreationMode = (specification == nil)

	return
}

type SPECIFICATIONFormCallback struct {
	specification *models.SPECIFICATION

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (specificationFormCallback *SPECIFICATIONFormCallback) OnSave() {
	specificationFormCallback.probe.stageOfInterest.Lock()
	defer specificationFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("SPECIFICATIONFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	specificationFormCallback.probe.formStage.Checkout()

	if specificationFormCallback.specification == nil {
		specificationFormCallback.specification = new(models.SPECIFICATION).Stage(specificationFormCallback.probe.stageOfInterest)
	}
	specification_ := specificationFormCallback.specification
	_ = specification_

	for _, formDiv := range specificationFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(specification_.Name), formDiv)
		case "DESC":
			FormDivBasicFieldToField(&(specification_.DESC), formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&(specification_.IDENTIFIER), formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&(specification_.LAST_CHANGE), formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&(specification_.LONG_NAME), formDiv)
		case "ALTERNATIVE_ID":
			FormDivSelectFieldToField(&(specification_.ALTERNATIVE_ID), specificationFormCallback.probe.stageOfInterest, formDiv)
		case "CHILDREN":
			FormDivSelectFieldToField(&(specification_.CHILDREN), specificationFormCallback.probe.stageOfInterest, formDiv)
		case "VALUES":
			FormDivSelectFieldToField(&(specification_.VALUES), specificationFormCallback.probe.stageOfInterest, formDiv)
		case "TYPE":
			FormDivSelectFieldToField(&(specification_.TYPE), specificationFormCallback.probe.stageOfInterest, formDiv)
		case "A_SPECIFICATIONS:SPECIFICATION":
			// WARNING : this form deals with the N-N association "A_SPECIFICATIONS.SPECIFICATION []*SPECIFICATION" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of SPECIFICATION). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.A_SPECIFICATIONS
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "A_SPECIFICATIONS"
				rf.Fieldname = "SPECIFICATION"
				formerAssociationSource := specification_.GongGetReverseFieldOwner(
					specificationFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.A_SPECIFICATIONS)
					if !ok {
						log.Fatalln("Source of A_SPECIFICATIONS.SPECIFICATION []*SPECIFICATION, is not an A_SPECIFICATIONS instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.SPECIFICATION, specification_)
					formerSource.SPECIFICATION = slices.Delete(formerSource.SPECIFICATION, idx, idx+1)
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
			var newSource *models.A_SPECIFICATIONS
			for _a_specifications := range *models.GetGongstructInstancesSet[models.A_SPECIFICATIONS](specificationFormCallback.probe.stageOfInterest) {
				// the match is base on the name
				if _a_specifications.GetName() == newSourceName.GetName() {
					newSource = _a_specifications // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of A_SPECIFICATIONS.SPECIFICATION []*SPECIFICATION, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.SPECIFICATION = append(newSource.SPECIFICATION, specification_)
		}
	}

	// manage the suppress operation
	if specificationFormCallback.formGroup.HasSuppressButtonBeenPressed {
		specification_.Unstage(specificationFormCallback.probe.stageOfInterest)
	}

	specificationFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.SPECIFICATION](
		specificationFormCallback.probe,
	)

	// display a new form by reset the form stage
	if specificationFormCallback.CreationMode || specificationFormCallback.formGroup.HasSuppressButtonBeenPressed {
		specificationFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(specificationFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__SPECIFICATIONFormCallback(
			nil,
			specificationFormCallback.probe,
			newFormGroup,
		)
		specification := new(models.SPECIFICATION)
		FillUpForm(specification, newFormGroup, specificationFormCallback.probe)
		specificationFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(specificationFormCallback.probe)
}

func __gong__New__SPECIFICATION_TYPEFormCallback(
	specification_type *models.SPECIFICATION_TYPE,
	probe *Probe,
	formGroup *table.FormGroup,
) (specification_typeFormCallback *SPECIFICATION_TYPEFormCallback) {
	specification_typeFormCallback = new(SPECIFICATION_TYPEFormCallback)
	specification_typeFormCallback.probe = probe
	specification_typeFormCallback.specification_type = specification_type
	specification_typeFormCallback.formGroup = formGroup

	specification_typeFormCallback.CreationMode = (specification_type == nil)

	return
}

type SPECIFICATION_TYPEFormCallback struct {
	specification_type *models.SPECIFICATION_TYPE

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (specification_typeFormCallback *SPECIFICATION_TYPEFormCallback) OnSave() {
	specification_typeFormCallback.probe.stageOfInterest.Lock()
	defer specification_typeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("SPECIFICATION_TYPEFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	specification_typeFormCallback.probe.formStage.Checkout()

	if specification_typeFormCallback.specification_type == nil {
		specification_typeFormCallback.specification_type = new(models.SPECIFICATION_TYPE).Stage(specification_typeFormCallback.probe.stageOfInterest)
	}
	specification_type_ := specification_typeFormCallback.specification_type
	_ = specification_type_

	for _, formDiv := range specification_typeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(specification_type_.Name), formDiv)
		case "DESC":
			FormDivBasicFieldToField(&(specification_type_.DESC), formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&(specification_type_.IDENTIFIER), formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&(specification_type_.LAST_CHANGE), formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&(specification_type_.LONG_NAME), formDiv)
		case "ALTERNATIVE_ID":
			FormDivSelectFieldToField(&(specification_type_.ALTERNATIVE_ID), specification_typeFormCallback.probe.stageOfInterest, formDiv)
		case "SPEC_ATTRIBUTES":
			FormDivSelectFieldToField(&(specification_type_.SPEC_ATTRIBUTES), specification_typeFormCallback.probe.stageOfInterest, formDiv)
		case "A_SPEC_TYPES:SPECIFICATION_TYPE":
			// WARNING : this form deals with the N-N association "A_SPEC_TYPES.SPECIFICATION_TYPE []*SPECIFICATION_TYPE" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of SPECIFICATION_TYPE). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.A_SPEC_TYPES
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "A_SPEC_TYPES"
				rf.Fieldname = "SPECIFICATION_TYPE"
				formerAssociationSource := specification_type_.GongGetReverseFieldOwner(
					specification_typeFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.A_SPEC_TYPES)
					if !ok {
						log.Fatalln("Source of A_SPEC_TYPES.SPECIFICATION_TYPE []*SPECIFICATION_TYPE, is not an A_SPEC_TYPES instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.SPECIFICATION_TYPE, specification_type_)
					formerSource.SPECIFICATION_TYPE = slices.Delete(formerSource.SPECIFICATION_TYPE, idx, idx+1)
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
			var newSource *models.A_SPEC_TYPES
			for _a_spec_types := range *models.GetGongstructInstancesSet[models.A_SPEC_TYPES](specification_typeFormCallback.probe.stageOfInterest) {
				// the match is base on the name
				if _a_spec_types.GetName() == newSourceName.GetName() {
					newSource = _a_spec_types // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of A_SPEC_TYPES.SPECIFICATION_TYPE []*SPECIFICATION_TYPE, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.SPECIFICATION_TYPE = append(newSource.SPECIFICATION_TYPE, specification_type_)
		}
	}

	// manage the suppress operation
	if specification_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		specification_type_.Unstage(specification_typeFormCallback.probe.stageOfInterest)
	}

	specification_typeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.SPECIFICATION_TYPE](
		specification_typeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if specification_typeFormCallback.CreationMode || specification_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		specification_typeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(specification_typeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__SPECIFICATION_TYPEFormCallback(
			nil,
			specification_typeFormCallback.probe,
			newFormGroup,
		)
		specification_type := new(models.SPECIFICATION_TYPE)
		FillUpForm(specification_type, newFormGroup, specification_typeFormCallback.probe)
		specification_typeFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(specification_typeFormCallback.probe)
}

func __gong__New__SPEC_HIERARCHYFormCallback(
	spec_hierarchy *models.SPEC_HIERARCHY,
	probe *Probe,
	formGroup *table.FormGroup,
) (spec_hierarchyFormCallback *SPEC_HIERARCHYFormCallback) {
	spec_hierarchyFormCallback = new(SPEC_HIERARCHYFormCallback)
	spec_hierarchyFormCallback.probe = probe
	spec_hierarchyFormCallback.spec_hierarchy = spec_hierarchy
	spec_hierarchyFormCallback.formGroup = formGroup

	spec_hierarchyFormCallback.CreationMode = (spec_hierarchy == nil)

	return
}

type SPEC_HIERARCHYFormCallback struct {
	spec_hierarchy *models.SPEC_HIERARCHY

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (spec_hierarchyFormCallback *SPEC_HIERARCHYFormCallback) OnSave() {
	spec_hierarchyFormCallback.probe.stageOfInterest.Lock()
	defer spec_hierarchyFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("SPEC_HIERARCHYFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	spec_hierarchyFormCallback.probe.formStage.Checkout()

	if spec_hierarchyFormCallback.spec_hierarchy == nil {
		spec_hierarchyFormCallback.spec_hierarchy = new(models.SPEC_HIERARCHY).Stage(spec_hierarchyFormCallback.probe.stageOfInterest)
	}
	spec_hierarchy_ := spec_hierarchyFormCallback.spec_hierarchy
	_ = spec_hierarchy_

	for _, formDiv := range spec_hierarchyFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(spec_hierarchy_.Name), formDiv)
		case "DESC":
			FormDivBasicFieldToField(&(spec_hierarchy_.DESC), formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&(spec_hierarchy_.IDENTIFIER), formDiv)
		case "IS_EDITABLE":
			FormDivBasicFieldToField(&(spec_hierarchy_.IS_EDITABLE), formDiv)
		case "IS_TABLE_INTERNAL":
			FormDivBasicFieldToField(&(spec_hierarchy_.IS_TABLE_INTERNAL), formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&(spec_hierarchy_.LAST_CHANGE), formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&(spec_hierarchy_.LONG_NAME), formDiv)
		case "ALTERNATIVE_ID":
			FormDivSelectFieldToField(&(spec_hierarchy_.ALTERNATIVE_ID), spec_hierarchyFormCallback.probe.stageOfInterest, formDiv)
		case "CHILDREN":
			FormDivSelectFieldToField(&(spec_hierarchy_.CHILDREN), spec_hierarchyFormCallback.probe.stageOfInterest, formDiv)
		case "EDITABLE_ATTS":
			FormDivSelectFieldToField(&(spec_hierarchy_.EDITABLE_ATTS), spec_hierarchyFormCallback.probe.stageOfInterest, formDiv)
		case "OBJECT":
			FormDivSelectFieldToField(&(spec_hierarchy_.OBJECT), spec_hierarchyFormCallback.probe.stageOfInterest, formDiv)
		case "A_CHILDREN:SPEC_HIERARCHY":
			// WARNING : this form deals with the N-N association "A_CHILDREN.SPEC_HIERARCHY []*SPEC_HIERARCHY" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of SPEC_HIERARCHY). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.A_CHILDREN
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "A_CHILDREN"
				rf.Fieldname = "SPEC_HIERARCHY"
				formerAssociationSource := spec_hierarchy_.GongGetReverseFieldOwner(
					spec_hierarchyFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.A_CHILDREN)
					if !ok {
						log.Fatalln("Source of A_CHILDREN.SPEC_HIERARCHY []*SPEC_HIERARCHY, is not an A_CHILDREN instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.SPEC_HIERARCHY, spec_hierarchy_)
					formerSource.SPEC_HIERARCHY = slices.Delete(formerSource.SPEC_HIERARCHY, idx, idx+1)
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
			var newSource *models.A_CHILDREN
			for _a_children := range *models.GetGongstructInstancesSet[models.A_CHILDREN](spec_hierarchyFormCallback.probe.stageOfInterest) {
				// the match is base on the name
				if _a_children.GetName() == newSourceName.GetName() {
					newSource = _a_children // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of A_CHILDREN.SPEC_HIERARCHY []*SPEC_HIERARCHY, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.SPEC_HIERARCHY = append(newSource.SPEC_HIERARCHY, spec_hierarchy_)
		}
	}

	// manage the suppress operation
	if spec_hierarchyFormCallback.formGroup.HasSuppressButtonBeenPressed {
		spec_hierarchy_.Unstage(spec_hierarchyFormCallback.probe.stageOfInterest)
	}

	spec_hierarchyFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.SPEC_HIERARCHY](
		spec_hierarchyFormCallback.probe,
	)

	// display a new form by reset the form stage
	if spec_hierarchyFormCallback.CreationMode || spec_hierarchyFormCallback.formGroup.HasSuppressButtonBeenPressed {
		spec_hierarchyFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(spec_hierarchyFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__SPEC_HIERARCHYFormCallback(
			nil,
			spec_hierarchyFormCallback.probe,
			newFormGroup,
		)
		spec_hierarchy := new(models.SPEC_HIERARCHY)
		FillUpForm(spec_hierarchy, newFormGroup, spec_hierarchyFormCallback.probe)
		spec_hierarchyFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(spec_hierarchyFormCallback.probe)
}

func __gong__New__SPEC_OBJECTFormCallback(
	spec_object *models.SPEC_OBJECT,
	probe *Probe,
	formGroup *table.FormGroup,
) (spec_objectFormCallback *SPEC_OBJECTFormCallback) {
	spec_objectFormCallback = new(SPEC_OBJECTFormCallback)
	spec_objectFormCallback.probe = probe
	spec_objectFormCallback.spec_object = spec_object
	spec_objectFormCallback.formGroup = formGroup

	spec_objectFormCallback.CreationMode = (spec_object == nil)

	return
}

type SPEC_OBJECTFormCallback struct {
	spec_object *models.SPEC_OBJECT

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (spec_objectFormCallback *SPEC_OBJECTFormCallback) OnSave() {
	spec_objectFormCallback.probe.stageOfInterest.Lock()
	defer spec_objectFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("SPEC_OBJECTFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	spec_objectFormCallback.probe.formStage.Checkout()

	if spec_objectFormCallback.spec_object == nil {
		spec_objectFormCallback.spec_object = new(models.SPEC_OBJECT).Stage(spec_objectFormCallback.probe.stageOfInterest)
	}
	spec_object_ := spec_objectFormCallback.spec_object
	_ = spec_object_

	for _, formDiv := range spec_objectFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(spec_object_.Name), formDiv)
		case "DESC":
			FormDivBasicFieldToField(&(spec_object_.DESC), formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&(spec_object_.IDENTIFIER), formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&(spec_object_.LAST_CHANGE), formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&(spec_object_.LONG_NAME), formDiv)
		case "ALTERNATIVE_ID":
			FormDivSelectFieldToField(&(spec_object_.ALTERNATIVE_ID), spec_objectFormCallback.probe.stageOfInterest, formDiv)
		case "VALUES":
			FormDivSelectFieldToField(&(spec_object_.VALUES), spec_objectFormCallback.probe.stageOfInterest, formDiv)
		case "TYPE":
			FormDivSelectFieldToField(&(spec_object_.TYPE), spec_objectFormCallback.probe.stageOfInterest, formDiv)
		case "A_SPEC_OBJECTS:SPEC_OBJECT":
			// WARNING : this form deals with the N-N association "A_SPEC_OBJECTS.SPEC_OBJECT []*SPEC_OBJECT" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of SPEC_OBJECT). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.A_SPEC_OBJECTS
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "A_SPEC_OBJECTS"
				rf.Fieldname = "SPEC_OBJECT"
				formerAssociationSource := spec_object_.GongGetReverseFieldOwner(
					spec_objectFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.A_SPEC_OBJECTS)
					if !ok {
						log.Fatalln("Source of A_SPEC_OBJECTS.SPEC_OBJECT []*SPEC_OBJECT, is not an A_SPEC_OBJECTS instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.SPEC_OBJECT, spec_object_)
					formerSource.SPEC_OBJECT = slices.Delete(formerSource.SPEC_OBJECT, idx, idx+1)
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
			var newSource *models.A_SPEC_OBJECTS
			for _a_spec_objects := range *models.GetGongstructInstancesSet[models.A_SPEC_OBJECTS](spec_objectFormCallback.probe.stageOfInterest) {
				// the match is base on the name
				if _a_spec_objects.GetName() == newSourceName.GetName() {
					newSource = _a_spec_objects // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of A_SPEC_OBJECTS.SPEC_OBJECT []*SPEC_OBJECT, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.SPEC_OBJECT = append(newSource.SPEC_OBJECT, spec_object_)
		}
	}

	// manage the suppress operation
	if spec_objectFormCallback.formGroup.HasSuppressButtonBeenPressed {
		spec_object_.Unstage(spec_objectFormCallback.probe.stageOfInterest)
	}

	spec_objectFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.SPEC_OBJECT](
		spec_objectFormCallback.probe,
	)

	// display a new form by reset the form stage
	if spec_objectFormCallback.CreationMode || spec_objectFormCallback.formGroup.HasSuppressButtonBeenPressed {
		spec_objectFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(spec_objectFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__SPEC_OBJECTFormCallback(
			nil,
			spec_objectFormCallback.probe,
			newFormGroup,
		)
		spec_object := new(models.SPEC_OBJECT)
		FillUpForm(spec_object, newFormGroup, spec_objectFormCallback.probe)
		spec_objectFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(spec_objectFormCallback.probe)
}

func __gong__New__SPEC_OBJECT_TYPEFormCallback(
	spec_object_type *models.SPEC_OBJECT_TYPE,
	probe *Probe,
	formGroup *table.FormGroup,
) (spec_object_typeFormCallback *SPEC_OBJECT_TYPEFormCallback) {
	spec_object_typeFormCallback = new(SPEC_OBJECT_TYPEFormCallback)
	spec_object_typeFormCallback.probe = probe
	spec_object_typeFormCallback.spec_object_type = spec_object_type
	spec_object_typeFormCallback.formGroup = formGroup

	spec_object_typeFormCallback.CreationMode = (spec_object_type == nil)

	return
}

type SPEC_OBJECT_TYPEFormCallback struct {
	spec_object_type *models.SPEC_OBJECT_TYPE

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (spec_object_typeFormCallback *SPEC_OBJECT_TYPEFormCallback) OnSave() {
	spec_object_typeFormCallback.probe.stageOfInterest.Lock()
	defer spec_object_typeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("SPEC_OBJECT_TYPEFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	spec_object_typeFormCallback.probe.formStage.Checkout()

	if spec_object_typeFormCallback.spec_object_type == nil {
		spec_object_typeFormCallback.spec_object_type = new(models.SPEC_OBJECT_TYPE).Stage(spec_object_typeFormCallback.probe.stageOfInterest)
	}
	spec_object_type_ := spec_object_typeFormCallback.spec_object_type
	_ = spec_object_type_

	for _, formDiv := range spec_object_typeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(spec_object_type_.Name), formDiv)
		case "DESC":
			FormDivBasicFieldToField(&(spec_object_type_.DESC), formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&(spec_object_type_.IDENTIFIER), formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&(spec_object_type_.LAST_CHANGE), formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&(spec_object_type_.LONG_NAME), formDiv)
		case "ALTERNATIVE_ID":
			FormDivSelectFieldToField(&(spec_object_type_.ALTERNATIVE_ID), spec_object_typeFormCallback.probe.stageOfInterest, formDiv)
		case "SPEC_ATTRIBUTES":
			FormDivSelectFieldToField(&(spec_object_type_.SPEC_ATTRIBUTES), spec_object_typeFormCallback.probe.stageOfInterest, formDiv)
		case "A_SPEC_TYPES:SPEC_OBJECT_TYPE":
			// WARNING : this form deals with the N-N association "A_SPEC_TYPES.SPEC_OBJECT_TYPE []*SPEC_OBJECT_TYPE" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of SPEC_OBJECT_TYPE). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.A_SPEC_TYPES
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "A_SPEC_TYPES"
				rf.Fieldname = "SPEC_OBJECT_TYPE"
				formerAssociationSource := spec_object_type_.GongGetReverseFieldOwner(
					spec_object_typeFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.A_SPEC_TYPES)
					if !ok {
						log.Fatalln("Source of A_SPEC_TYPES.SPEC_OBJECT_TYPE []*SPEC_OBJECT_TYPE, is not an A_SPEC_TYPES instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.SPEC_OBJECT_TYPE, spec_object_type_)
					formerSource.SPEC_OBJECT_TYPE = slices.Delete(formerSource.SPEC_OBJECT_TYPE, idx, idx+1)
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
			var newSource *models.A_SPEC_TYPES
			for _a_spec_types := range *models.GetGongstructInstancesSet[models.A_SPEC_TYPES](spec_object_typeFormCallback.probe.stageOfInterest) {
				// the match is base on the name
				if _a_spec_types.GetName() == newSourceName.GetName() {
					newSource = _a_spec_types // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of A_SPEC_TYPES.SPEC_OBJECT_TYPE []*SPEC_OBJECT_TYPE, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.SPEC_OBJECT_TYPE = append(newSource.SPEC_OBJECT_TYPE, spec_object_type_)
		}
	}

	// manage the suppress operation
	if spec_object_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		spec_object_type_.Unstage(spec_object_typeFormCallback.probe.stageOfInterest)
	}

	spec_object_typeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.SPEC_OBJECT_TYPE](
		spec_object_typeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if spec_object_typeFormCallback.CreationMode || spec_object_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		spec_object_typeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(spec_object_typeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__SPEC_OBJECT_TYPEFormCallback(
			nil,
			spec_object_typeFormCallback.probe,
			newFormGroup,
		)
		spec_object_type := new(models.SPEC_OBJECT_TYPE)
		FillUpForm(spec_object_type, newFormGroup, spec_object_typeFormCallback.probe)
		spec_object_typeFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(spec_object_typeFormCallback.probe)
}

func __gong__New__SPEC_RELATIONFormCallback(
	spec_relation *models.SPEC_RELATION,
	probe *Probe,
	formGroup *table.FormGroup,
) (spec_relationFormCallback *SPEC_RELATIONFormCallback) {
	spec_relationFormCallback = new(SPEC_RELATIONFormCallback)
	spec_relationFormCallback.probe = probe
	spec_relationFormCallback.spec_relation = spec_relation
	spec_relationFormCallback.formGroup = formGroup

	spec_relationFormCallback.CreationMode = (spec_relation == nil)

	return
}

type SPEC_RELATIONFormCallback struct {
	spec_relation *models.SPEC_RELATION

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (spec_relationFormCallback *SPEC_RELATIONFormCallback) OnSave() {
	spec_relationFormCallback.probe.stageOfInterest.Lock()
	defer spec_relationFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("SPEC_RELATIONFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	spec_relationFormCallback.probe.formStage.Checkout()

	if spec_relationFormCallback.spec_relation == nil {
		spec_relationFormCallback.spec_relation = new(models.SPEC_RELATION).Stage(spec_relationFormCallback.probe.stageOfInterest)
	}
	spec_relation_ := spec_relationFormCallback.spec_relation
	_ = spec_relation_

	for _, formDiv := range spec_relationFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(spec_relation_.Name), formDiv)
		case "DESC":
			FormDivBasicFieldToField(&(spec_relation_.DESC), formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&(spec_relation_.IDENTIFIER), formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&(spec_relation_.LAST_CHANGE), formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&(spec_relation_.LONG_NAME), formDiv)
		case "ALTERNATIVE_ID":
			FormDivSelectFieldToField(&(spec_relation_.ALTERNATIVE_ID), spec_relationFormCallback.probe.stageOfInterest, formDiv)
		case "VALUES":
			FormDivSelectFieldToField(&(spec_relation_.VALUES), spec_relationFormCallback.probe.stageOfInterest, formDiv)
		case "SOURCE":
			FormDivSelectFieldToField(&(spec_relation_.SOURCE), spec_relationFormCallback.probe.stageOfInterest, formDiv)
		case "TARGET":
			FormDivSelectFieldToField(&(spec_relation_.TARGET), spec_relationFormCallback.probe.stageOfInterest, formDiv)
		case "TYPE":
			FormDivSelectFieldToField(&(spec_relation_.TYPE), spec_relationFormCallback.probe.stageOfInterest, formDiv)
		case "A_SPEC_RELATIONS:SPEC_RELATION":
			// WARNING : this form deals with the N-N association "A_SPEC_RELATIONS.SPEC_RELATION []*SPEC_RELATION" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of SPEC_RELATION). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.A_SPEC_RELATIONS
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "A_SPEC_RELATIONS"
				rf.Fieldname = "SPEC_RELATION"
				formerAssociationSource := spec_relation_.GongGetReverseFieldOwner(
					spec_relationFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.A_SPEC_RELATIONS)
					if !ok {
						log.Fatalln("Source of A_SPEC_RELATIONS.SPEC_RELATION []*SPEC_RELATION, is not an A_SPEC_RELATIONS instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.SPEC_RELATION, spec_relation_)
					formerSource.SPEC_RELATION = slices.Delete(formerSource.SPEC_RELATION, idx, idx+1)
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
			var newSource *models.A_SPEC_RELATIONS
			for _a_spec_relations := range *models.GetGongstructInstancesSet[models.A_SPEC_RELATIONS](spec_relationFormCallback.probe.stageOfInterest) {
				// the match is base on the name
				if _a_spec_relations.GetName() == newSourceName.GetName() {
					newSource = _a_spec_relations // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of A_SPEC_RELATIONS.SPEC_RELATION []*SPEC_RELATION, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.SPEC_RELATION = append(newSource.SPEC_RELATION, spec_relation_)
		}
	}

	// manage the suppress operation
	if spec_relationFormCallback.formGroup.HasSuppressButtonBeenPressed {
		spec_relation_.Unstage(spec_relationFormCallback.probe.stageOfInterest)
	}

	spec_relationFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.SPEC_RELATION](
		spec_relationFormCallback.probe,
	)

	// display a new form by reset the form stage
	if spec_relationFormCallback.CreationMode || spec_relationFormCallback.formGroup.HasSuppressButtonBeenPressed {
		spec_relationFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(spec_relationFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__SPEC_RELATIONFormCallback(
			nil,
			spec_relationFormCallback.probe,
			newFormGroup,
		)
		spec_relation := new(models.SPEC_RELATION)
		FillUpForm(spec_relation, newFormGroup, spec_relationFormCallback.probe)
		spec_relationFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(spec_relationFormCallback.probe)
}

func __gong__New__SPEC_RELATION_TYPEFormCallback(
	spec_relation_type *models.SPEC_RELATION_TYPE,
	probe *Probe,
	formGroup *table.FormGroup,
) (spec_relation_typeFormCallback *SPEC_RELATION_TYPEFormCallback) {
	spec_relation_typeFormCallback = new(SPEC_RELATION_TYPEFormCallback)
	spec_relation_typeFormCallback.probe = probe
	spec_relation_typeFormCallback.spec_relation_type = spec_relation_type
	spec_relation_typeFormCallback.formGroup = formGroup

	spec_relation_typeFormCallback.CreationMode = (spec_relation_type == nil)

	return
}

type SPEC_RELATION_TYPEFormCallback struct {
	spec_relation_type *models.SPEC_RELATION_TYPE

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (spec_relation_typeFormCallback *SPEC_RELATION_TYPEFormCallback) OnSave() {
	spec_relation_typeFormCallback.probe.stageOfInterest.Lock()
	defer spec_relation_typeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("SPEC_RELATION_TYPEFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	spec_relation_typeFormCallback.probe.formStage.Checkout()

	if spec_relation_typeFormCallback.spec_relation_type == nil {
		spec_relation_typeFormCallback.spec_relation_type = new(models.SPEC_RELATION_TYPE).Stage(spec_relation_typeFormCallback.probe.stageOfInterest)
	}
	spec_relation_type_ := spec_relation_typeFormCallback.spec_relation_type
	_ = spec_relation_type_

	for _, formDiv := range spec_relation_typeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(spec_relation_type_.Name), formDiv)
		case "DESC":
			FormDivBasicFieldToField(&(spec_relation_type_.DESC), formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&(spec_relation_type_.IDENTIFIER), formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&(spec_relation_type_.LAST_CHANGE), formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&(spec_relation_type_.LONG_NAME), formDiv)
		case "ALTERNATIVE_ID":
			FormDivSelectFieldToField(&(spec_relation_type_.ALTERNATIVE_ID), spec_relation_typeFormCallback.probe.stageOfInterest, formDiv)
		case "SPEC_ATTRIBUTES":
			FormDivSelectFieldToField(&(spec_relation_type_.SPEC_ATTRIBUTES), spec_relation_typeFormCallback.probe.stageOfInterest, formDiv)
		case "A_SPEC_TYPES:SPEC_RELATION_TYPE":
			// WARNING : this form deals with the N-N association "A_SPEC_TYPES.SPEC_RELATION_TYPE []*SPEC_RELATION_TYPE" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of SPEC_RELATION_TYPE). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.A_SPEC_TYPES
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "A_SPEC_TYPES"
				rf.Fieldname = "SPEC_RELATION_TYPE"
				formerAssociationSource := spec_relation_type_.GongGetReverseFieldOwner(
					spec_relation_typeFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.A_SPEC_TYPES)
					if !ok {
						log.Fatalln("Source of A_SPEC_TYPES.SPEC_RELATION_TYPE []*SPEC_RELATION_TYPE, is not an A_SPEC_TYPES instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.SPEC_RELATION_TYPE, spec_relation_type_)
					formerSource.SPEC_RELATION_TYPE = slices.Delete(formerSource.SPEC_RELATION_TYPE, idx, idx+1)
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
			var newSource *models.A_SPEC_TYPES
			for _a_spec_types := range *models.GetGongstructInstancesSet[models.A_SPEC_TYPES](spec_relation_typeFormCallback.probe.stageOfInterest) {
				// the match is base on the name
				if _a_spec_types.GetName() == newSourceName.GetName() {
					newSource = _a_spec_types // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of A_SPEC_TYPES.SPEC_RELATION_TYPE []*SPEC_RELATION_TYPE, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.SPEC_RELATION_TYPE = append(newSource.SPEC_RELATION_TYPE, spec_relation_type_)
		}
	}

	// manage the suppress operation
	if spec_relation_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		spec_relation_type_.Unstage(spec_relation_typeFormCallback.probe.stageOfInterest)
	}

	spec_relation_typeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.SPEC_RELATION_TYPE](
		spec_relation_typeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if spec_relation_typeFormCallback.CreationMode || spec_relation_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		spec_relation_typeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(spec_relation_typeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__SPEC_RELATION_TYPEFormCallback(
			nil,
			spec_relation_typeFormCallback.probe,
			newFormGroup,
		)
		spec_relation_type := new(models.SPEC_RELATION_TYPE)
		FillUpForm(spec_relation_type, newFormGroup, spec_relation_typeFormCallback.probe)
		spec_relation_typeFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(spec_relation_typeFormCallback.probe)
}

func __gong__New__XHTML_CONTENTFormCallback(
	xhtml_content *models.XHTML_CONTENT,
	probe *Probe,
	formGroup *table.FormGroup,
) (xhtml_contentFormCallback *XHTML_CONTENTFormCallback) {
	xhtml_contentFormCallback = new(XHTML_CONTENTFormCallback)
	xhtml_contentFormCallback.probe = probe
	xhtml_contentFormCallback.xhtml_content = xhtml_content
	xhtml_contentFormCallback.formGroup = formGroup

	xhtml_contentFormCallback.CreationMode = (xhtml_content == nil)

	return
}

type XHTML_CONTENTFormCallback struct {
	xhtml_content *models.XHTML_CONTENT

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (xhtml_contentFormCallback *XHTML_CONTENTFormCallback) OnSave() {
	xhtml_contentFormCallback.probe.stageOfInterest.Lock()
	defer xhtml_contentFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("XHTML_CONTENTFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	xhtml_contentFormCallback.probe.formStage.Checkout()

	if xhtml_contentFormCallback.xhtml_content == nil {
		xhtml_contentFormCallback.xhtml_content = new(models.XHTML_CONTENT).Stage(xhtml_contentFormCallback.probe.stageOfInterest)
	}
	xhtml_content_ := xhtml_contentFormCallback.xhtml_content
	_ = xhtml_content_

	for _, formDiv := range xhtml_contentFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(xhtml_content_.Name), formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&(xhtml_content_.EnclosedText), formDiv)
		case "PureText":
			FormDivBasicFieldToField(&(xhtml_content_.PureText), formDiv)
		}
	}

	// manage the suppress operation
	if xhtml_contentFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xhtml_content_.Unstage(xhtml_contentFormCallback.probe.stageOfInterest)
	}

	xhtml_contentFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.XHTML_CONTENT](
		xhtml_contentFormCallback.probe,
	)

	// display a new form by reset the form stage
	if xhtml_contentFormCallback.CreationMode || xhtml_contentFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xhtml_contentFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(xhtml_contentFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__XHTML_CONTENTFormCallback(
			nil,
			xhtml_contentFormCallback.probe,
			newFormGroup,
		)
		xhtml_content := new(models.XHTML_CONTENT)
		FillUpForm(xhtml_content, newFormGroup, xhtml_contentFormCallback.probe)
		xhtml_contentFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(xhtml_contentFormCallback.probe)
}
