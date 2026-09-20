// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/app/reqif/go/models"
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
	var zero T
	return &FormCallback[T]{
		Instance:     instance,
		CreationMode: instance == zero,
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

	var zero T
	if cb.Instance == zero {
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
func __gong__New__ALTERNATIVE_IDFormCallback(
	_instance *models.ALTERNATIVE_ID,
	probe *Probe,
	formGroup *form.FormGroup,
) (alternative_idFormCallback *FormCallback[*models.ALTERNATIVE_ID]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveALTERNATIVE_IDFields,
	)
}

type ALTERNATIVE_IDFormCallback = FormCallback[*models.ALTERNATIVE_ID]

func saveALTERNATIVE_IDFields(
	_instance *models.ALTERNATIVE_ID,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&(_instance.IDENTIFIER), formDiv)
		}
	}
}

func __gong__New__ATTRIBUTE_DEFINITION_BOOLEANFormCallback(
	_instance *models.ATTRIBUTE_DEFINITION_BOOLEAN,
	probe *Probe,
	formGroup *form.FormGroup,
) (attribute_definition_booleanFormCallback *FormCallback[*models.ATTRIBUTE_DEFINITION_BOOLEAN]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveATTRIBUTE_DEFINITION_BOOLEANFields,
	)
}

type ATTRIBUTE_DEFINITION_BOOLEANFormCallback = FormCallback[*models.ATTRIBUTE_DEFINITION_BOOLEAN]

func saveATTRIBUTE_DEFINITION_BOOLEANFields(
	_instance *models.ATTRIBUTE_DEFINITION_BOOLEAN,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "DESC":
			FormDivBasicFieldToField(&(_instance.DESC), formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&(_instance.IDENTIFIER), formDiv)
		case "IS_EDITABLE":
			FormDivBasicFieldToField(&(_instance.IS_EDITABLE), formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&(_instance.LAST_CHANGE), formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&(_instance.LONG_NAME), formDiv)
		case "ALTERNATIVE_ID":
			FormDivSelectFieldToField(&(_instance.ALTERNATIVE_ID), probe.stageOfInterest, formDiv)
		case "DEFAULT_VALUE":
			FormDivSelectFieldToField(&(_instance.DEFAULT_VALUE), probe.stageOfInterest, formDiv)
		case "TYPE":
			FormDivSelectFieldToField(&(_instance.TYPE), probe.stageOfInterest, formDiv)
		case "A_SPEC_ATTRIBUTES:ATTRIBUTE_DEFINITION_BOOLEAN":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "ATTRIBUTE_DEFINITION_BOOLEAN", func(owner *models.A_SPEC_ATTRIBUTES) *[]*models.ATTRIBUTE_DEFINITION_BOOLEAN { return &owner.ATTRIBUTE_DEFINITION_BOOLEAN })
		}
	}
}

func __gong__New__ATTRIBUTE_DEFINITION_BOOLEAN_RenderingFormCallback(
	_instance *models.ATTRIBUTE_DEFINITION_BOOLEAN_Rendering,
	probe *Probe,
	formGroup *form.FormGroup,
) (attribute_definition_boolean_renderingFormCallback *FormCallback[*models.ATTRIBUTE_DEFINITION_BOOLEAN_Rendering]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveATTRIBUTE_DEFINITION_BOOLEAN_RenderingFields,
	)
}

type ATTRIBUTE_DEFINITION_BOOLEAN_RenderingFormCallback = FormCallback[*models.ATTRIBUTE_DEFINITION_BOOLEAN_Rendering]

func saveATTRIBUTE_DEFINITION_BOOLEAN_RenderingFields(
	_instance *models.ATTRIBUTE_DEFINITION_BOOLEAN_Rendering,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "ShowInTable":
			FormDivBasicFieldToField(&(_instance.ShowInTable), formDiv)
		case "ShowInTitle":
			FormDivBasicFieldToField(&(_instance.ShowInTitle), formDiv)
		case "ShowInSubject":
			FormDivBasicFieldToField(&(_instance.ShowInSubject), formDiv)
		case "Rank":
			FormDivBasicFieldToField(&(_instance.Rank), formDiv)
		}
	}
}

func __gong__New__ATTRIBUTE_DEFINITION_DATEFormCallback(
	_instance *models.ATTRIBUTE_DEFINITION_DATE,
	probe *Probe,
	formGroup *form.FormGroup,
) (attribute_definition_dateFormCallback *FormCallback[*models.ATTRIBUTE_DEFINITION_DATE]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveATTRIBUTE_DEFINITION_DATEFields,
	)
}

type ATTRIBUTE_DEFINITION_DATEFormCallback = FormCallback[*models.ATTRIBUTE_DEFINITION_DATE]

func saveATTRIBUTE_DEFINITION_DATEFields(
	_instance *models.ATTRIBUTE_DEFINITION_DATE,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "DESC":
			FormDivBasicFieldToField(&(_instance.DESC), formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&(_instance.IDENTIFIER), formDiv)
		case "IS_EDITABLE":
			FormDivBasicFieldToField(&(_instance.IS_EDITABLE), formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&(_instance.LAST_CHANGE), formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&(_instance.LONG_NAME), formDiv)
		case "ALTERNATIVE_ID":
			FormDivSelectFieldToField(&(_instance.ALTERNATIVE_ID), probe.stageOfInterest, formDiv)
		case "DEFAULT_VALUE":
			FormDivSelectFieldToField(&(_instance.DEFAULT_VALUE), probe.stageOfInterest, formDiv)
		case "TYPE":
			FormDivSelectFieldToField(&(_instance.TYPE), probe.stageOfInterest, formDiv)
		case "A_SPEC_ATTRIBUTES:ATTRIBUTE_DEFINITION_DATE":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "ATTRIBUTE_DEFINITION_DATE", func(owner *models.A_SPEC_ATTRIBUTES) *[]*models.ATTRIBUTE_DEFINITION_DATE { return &owner.ATTRIBUTE_DEFINITION_DATE })
		}
	}
}

func __gong__New__ATTRIBUTE_DEFINITION_DATE_RenderingFormCallback(
	_instance *models.ATTRIBUTE_DEFINITION_DATE_Rendering,
	probe *Probe,
	formGroup *form.FormGroup,
) (attribute_definition_date_renderingFormCallback *FormCallback[*models.ATTRIBUTE_DEFINITION_DATE_Rendering]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveATTRIBUTE_DEFINITION_DATE_RenderingFields,
	)
}

type ATTRIBUTE_DEFINITION_DATE_RenderingFormCallback = FormCallback[*models.ATTRIBUTE_DEFINITION_DATE_Rendering]

func saveATTRIBUTE_DEFINITION_DATE_RenderingFields(
	_instance *models.ATTRIBUTE_DEFINITION_DATE_Rendering,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "ShowInTable":
			FormDivBasicFieldToField(&(_instance.ShowInTable), formDiv)
		case "ShowInTitle":
			FormDivBasicFieldToField(&(_instance.ShowInTitle), formDiv)
		case "ShowInSubject":
			FormDivBasicFieldToField(&(_instance.ShowInSubject), formDiv)
		case "Rank":
			FormDivBasicFieldToField(&(_instance.Rank), formDiv)
		}
	}
}

func __gong__New__ATTRIBUTE_DEFINITION_ENUMERATIONFormCallback(
	_instance *models.ATTRIBUTE_DEFINITION_ENUMERATION,
	probe *Probe,
	formGroup *form.FormGroup,
) (attribute_definition_enumerationFormCallback *FormCallback[*models.ATTRIBUTE_DEFINITION_ENUMERATION]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveATTRIBUTE_DEFINITION_ENUMERATIONFields,
	)
}

type ATTRIBUTE_DEFINITION_ENUMERATIONFormCallback = FormCallback[*models.ATTRIBUTE_DEFINITION_ENUMERATION]

func saveATTRIBUTE_DEFINITION_ENUMERATIONFields(
	_instance *models.ATTRIBUTE_DEFINITION_ENUMERATION,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "DESC":
			FormDivBasicFieldToField(&(_instance.DESC), formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&(_instance.IDENTIFIER), formDiv)
		case "IS_EDITABLE":
			FormDivBasicFieldToField(&(_instance.IS_EDITABLE), formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&(_instance.LAST_CHANGE), formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&(_instance.LONG_NAME), formDiv)
		case "MULTI_VALUED":
			FormDivBasicFieldToField(&(_instance.MULTI_VALUED), formDiv)
		case "ALTERNATIVE_ID":
			FormDivSelectFieldToField(&(_instance.ALTERNATIVE_ID), probe.stageOfInterest, formDiv)
		case "DEFAULT_VALUE":
			FormDivSelectFieldToField(&(_instance.DEFAULT_VALUE), probe.stageOfInterest, formDiv)
		case "TYPE":
			FormDivSelectFieldToField(&(_instance.TYPE), probe.stageOfInterest, formDiv)
		case "A_SPEC_ATTRIBUTES:ATTRIBUTE_DEFINITION_ENUMERATION":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "ATTRIBUTE_DEFINITION_ENUMERATION", func(owner *models.A_SPEC_ATTRIBUTES) *[]*models.ATTRIBUTE_DEFINITION_ENUMERATION { return &owner.ATTRIBUTE_DEFINITION_ENUMERATION })
		}
	}
}

func __gong__New__ATTRIBUTE_DEFINITION_ENUMERATION_RenderingFormCallback(
	_instance *models.ATTRIBUTE_DEFINITION_ENUMERATION_Rendering,
	probe *Probe,
	formGroup *form.FormGroup,
) (attribute_definition_enumeration_renderingFormCallback *FormCallback[*models.ATTRIBUTE_DEFINITION_ENUMERATION_Rendering]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveATTRIBUTE_DEFINITION_ENUMERATION_RenderingFields,
	)
}

type ATTRIBUTE_DEFINITION_ENUMERATION_RenderingFormCallback = FormCallback[*models.ATTRIBUTE_DEFINITION_ENUMERATION_Rendering]

func saveATTRIBUTE_DEFINITION_ENUMERATION_RenderingFields(
	_instance *models.ATTRIBUTE_DEFINITION_ENUMERATION_Rendering,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "ShowInTable":
			FormDivBasicFieldToField(&(_instance.ShowInTable), formDiv)
		case "ShowInTitle":
			FormDivBasicFieldToField(&(_instance.ShowInTitle), formDiv)
		case "ShowInSubject":
			FormDivBasicFieldToField(&(_instance.ShowInSubject), formDiv)
		case "Rank":
			FormDivBasicFieldToField(&(_instance.Rank), formDiv)
		}
	}
}

func __gong__New__ATTRIBUTE_DEFINITION_INTEGERFormCallback(
	_instance *models.ATTRIBUTE_DEFINITION_INTEGER,
	probe *Probe,
	formGroup *form.FormGroup,
) (attribute_definition_integerFormCallback *FormCallback[*models.ATTRIBUTE_DEFINITION_INTEGER]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveATTRIBUTE_DEFINITION_INTEGERFields,
	)
}

type ATTRIBUTE_DEFINITION_INTEGERFormCallback = FormCallback[*models.ATTRIBUTE_DEFINITION_INTEGER]

func saveATTRIBUTE_DEFINITION_INTEGERFields(
	_instance *models.ATTRIBUTE_DEFINITION_INTEGER,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "DESC":
			FormDivBasicFieldToField(&(_instance.DESC), formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&(_instance.IDENTIFIER), formDiv)
		case "IS_EDITABLE":
			FormDivBasicFieldToField(&(_instance.IS_EDITABLE), formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&(_instance.LAST_CHANGE), formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&(_instance.LONG_NAME), formDiv)
		case "ALTERNATIVE_ID":
			FormDivSelectFieldToField(&(_instance.ALTERNATIVE_ID), probe.stageOfInterest, formDiv)
		case "DEFAULT_VALUE":
			FormDivSelectFieldToField(&(_instance.DEFAULT_VALUE), probe.stageOfInterest, formDiv)
		case "TYPE":
			FormDivSelectFieldToField(&(_instance.TYPE), probe.stageOfInterest, formDiv)
		case "A_SPEC_ATTRIBUTES:ATTRIBUTE_DEFINITION_INTEGER":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "ATTRIBUTE_DEFINITION_INTEGER", func(owner *models.A_SPEC_ATTRIBUTES) *[]*models.ATTRIBUTE_DEFINITION_INTEGER { return &owner.ATTRIBUTE_DEFINITION_INTEGER })
		}
	}
}

func __gong__New__ATTRIBUTE_DEFINITION_INTEGER_RenderingFormCallback(
	_instance *models.ATTRIBUTE_DEFINITION_INTEGER_Rendering,
	probe *Probe,
	formGroup *form.FormGroup,
) (attribute_definition_integer_renderingFormCallback *FormCallback[*models.ATTRIBUTE_DEFINITION_INTEGER_Rendering]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveATTRIBUTE_DEFINITION_INTEGER_RenderingFields,
	)
}

type ATTRIBUTE_DEFINITION_INTEGER_RenderingFormCallback = FormCallback[*models.ATTRIBUTE_DEFINITION_INTEGER_Rendering]

func saveATTRIBUTE_DEFINITION_INTEGER_RenderingFields(
	_instance *models.ATTRIBUTE_DEFINITION_INTEGER_Rendering,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "ShowInTable":
			FormDivBasicFieldToField(&(_instance.ShowInTable), formDiv)
		case "ShowInTitle":
			FormDivBasicFieldToField(&(_instance.ShowInTitle), formDiv)
		case "ShowInSubject":
			FormDivBasicFieldToField(&(_instance.ShowInSubject), formDiv)
		case "Rank":
			FormDivBasicFieldToField(&(_instance.Rank), formDiv)
		}
	}
}

func __gong__New__ATTRIBUTE_DEFINITION_REALFormCallback(
	_instance *models.ATTRIBUTE_DEFINITION_REAL,
	probe *Probe,
	formGroup *form.FormGroup,
) (attribute_definition_realFormCallback *FormCallback[*models.ATTRIBUTE_DEFINITION_REAL]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveATTRIBUTE_DEFINITION_REALFields,
	)
}

type ATTRIBUTE_DEFINITION_REALFormCallback = FormCallback[*models.ATTRIBUTE_DEFINITION_REAL]

func saveATTRIBUTE_DEFINITION_REALFields(
	_instance *models.ATTRIBUTE_DEFINITION_REAL,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "DESC":
			FormDivBasicFieldToField(&(_instance.DESC), formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&(_instance.IDENTIFIER), formDiv)
		case "IS_EDITABLE":
			FormDivBasicFieldToField(&(_instance.IS_EDITABLE), formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&(_instance.LAST_CHANGE), formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&(_instance.LONG_NAME), formDiv)
		case "ALTERNATIVE_ID":
			FormDivSelectFieldToField(&(_instance.ALTERNATIVE_ID), probe.stageOfInterest, formDiv)
		case "DEFAULT_VALUE":
			FormDivSelectFieldToField(&(_instance.DEFAULT_VALUE), probe.stageOfInterest, formDiv)
		case "TYPE":
			FormDivSelectFieldToField(&(_instance.TYPE), probe.stageOfInterest, formDiv)
		case "A_SPEC_ATTRIBUTES:ATTRIBUTE_DEFINITION_REAL":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "ATTRIBUTE_DEFINITION_REAL", func(owner *models.A_SPEC_ATTRIBUTES) *[]*models.ATTRIBUTE_DEFINITION_REAL { return &owner.ATTRIBUTE_DEFINITION_REAL })
		}
	}
}

func __gong__New__ATTRIBUTE_DEFINITION_REAL_RenderingFormCallback(
	_instance *models.ATTRIBUTE_DEFINITION_REAL_Rendering,
	probe *Probe,
	formGroup *form.FormGroup,
) (attribute_definition_real_renderingFormCallback *FormCallback[*models.ATTRIBUTE_DEFINITION_REAL_Rendering]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveATTRIBUTE_DEFINITION_REAL_RenderingFields,
	)
}

type ATTRIBUTE_DEFINITION_REAL_RenderingFormCallback = FormCallback[*models.ATTRIBUTE_DEFINITION_REAL_Rendering]

func saveATTRIBUTE_DEFINITION_REAL_RenderingFields(
	_instance *models.ATTRIBUTE_DEFINITION_REAL_Rendering,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "ShowInTable":
			FormDivBasicFieldToField(&(_instance.ShowInTable), formDiv)
		case "ShowInTitle":
			FormDivBasicFieldToField(&(_instance.ShowInTitle), formDiv)
		case "ShowInSubject":
			FormDivBasicFieldToField(&(_instance.ShowInSubject), formDiv)
		case "Rank":
			FormDivBasicFieldToField(&(_instance.Rank), formDiv)
		}
	}
}

func __gong__New__ATTRIBUTE_DEFINITION_RenderingFormCallback(
	_instance *models.ATTRIBUTE_DEFINITION_Rendering,
	probe *Probe,
	formGroup *form.FormGroup,
) (attribute_definition_renderingFormCallback *FormCallback[*models.ATTRIBUTE_DEFINITION_Rendering]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveATTRIBUTE_DEFINITION_RenderingFields,
	)
}

type ATTRIBUTE_DEFINITION_RenderingFormCallback = FormCallback[*models.ATTRIBUTE_DEFINITION_Rendering]

func saveATTRIBUTE_DEFINITION_RenderingFields(
	_instance *models.ATTRIBUTE_DEFINITION_Rendering,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "ShowInTable":
			FormDivBasicFieldToField(&(_instance.ShowInTable), formDiv)
		case "ShowInTitle":
			FormDivBasicFieldToField(&(_instance.ShowInTitle), formDiv)
		case "ShowInSubject":
			FormDivBasicFieldToField(&(_instance.ShowInSubject), formDiv)
		case "Rank":
			FormDivBasicFieldToField(&(_instance.Rank), formDiv)
		}
	}
}

func __gong__New__ATTRIBUTE_DEFINITION_STRINGFormCallback(
	_instance *models.ATTRIBUTE_DEFINITION_STRING,
	probe *Probe,
	formGroup *form.FormGroup,
) (attribute_definition_stringFormCallback *FormCallback[*models.ATTRIBUTE_DEFINITION_STRING]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveATTRIBUTE_DEFINITION_STRINGFields,
	)
}

type ATTRIBUTE_DEFINITION_STRINGFormCallback = FormCallback[*models.ATTRIBUTE_DEFINITION_STRING]

func saveATTRIBUTE_DEFINITION_STRINGFields(
	_instance *models.ATTRIBUTE_DEFINITION_STRING,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "DESC":
			FormDivBasicFieldToField(&(_instance.DESC), formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&(_instance.IDENTIFIER), formDiv)
		case "IS_EDITABLE":
			FormDivBasicFieldToField(&(_instance.IS_EDITABLE), formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&(_instance.LAST_CHANGE), formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&(_instance.LONG_NAME), formDiv)
		case "ALTERNATIVE_ID":
			FormDivSelectFieldToField(&(_instance.ALTERNATIVE_ID), probe.stageOfInterest, formDiv)
		case "DEFAULT_VALUE":
			FormDivSelectFieldToField(&(_instance.DEFAULT_VALUE), probe.stageOfInterest, formDiv)
		case "TYPE":
			FormDivSelectFieldToField(&(_instance.TYPE), probe.stageOfInterest, formDiv)
		case "A_SPEC_ATTRIBUTES:ATTRIBUTE_DEFINITION_STRING":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "ATTRIBUTE_DEFINITION_STRING", func(owner *models.A_SPEC_ATTRIBUTES) *[]*models.ATTRIBUTE_DEFINITION_STRING { return &owner.ATTRIBUTE_DEFINITION_STRING })
		}
	}
}

func __gong__New__ATTRIBUTE_DEFINITION_STRING_RenderingFormCallback(
	_instance *models.ATTRIBUTE_DEFINITION_STRING_Rendering,
	probe *Probe,
	formGroup *form.FormGroup,
) (attribute_definition_string_renderingFormCallback *FormCallback[*models.ATTRIBUTE_DEFINITION_STRING_Rendering]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveATTRIBUTE_DEFINITION_STRING_RenderingFields,
	)
}

type ATTRIBUTE_DEFINITION_STRING_RenderingFormCallback = FormCallback[*models.ATTRIBUTE_DEFINITION_STRING_Rendering]

func saveATTRIBUTE_DEFINITION_STRING_RenderingFields(
	_instance *models.ATTRIBUTE_DEFINITION_STRING_Rendering,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "ShowInTable":
			FormDivBasicFieldToField(&(_instance.ShowInTable), formDiv)
		case "ShowInTitle":
			FormDivBasicFieldToField(&(_instance.ShowInTitle), formDiv)
		case "ShowInSubject":
			FormDivBasicFieldToField(&(_instance.ShowInSubject), formDiv)
		case "Rank":
			FormDivBasicFieldToField(&(_instance.Rank), formDiv)
		}
	}
}

func __gong__New__ATTRIBUTE_DEFINITION_XHTMLFormCallback(
	_instance *models.ATTRIBUTE_DEFINITION_XHTML,
	probe *Probe,
	formGroup *form.FormGroup,
) (attribute_definition_xhtmlFormCallback *FormCallback[*models.ATTRIBUTE_DEFINITION_XHTML]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveATTRIBUTE_DEFINITION_XHTMLFields,
	)
}

type ATTRIBUTE_DEFINITION_XHTMLFormCallback = FormCallback[*models.ATTRIBUTE_DEFINITION_XHTML]

func saveATTRIBUTE_DEFINITION_XHTMLFields(
	_instance *models.ATTRIBUTE_DEFINITION_XHTML,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "DESC":
			FormDivBasicFieldToField(&(_instance.DESC), formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&(_instance.IDENTIFIER), formDiv)
		case "IS_EDITABLE":
			FormDivBasicFieldToField(&(_instance.IS_EDITABLE), formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&(_instance.LAST_CHANGE), formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&(_instance.LONG_NAME), formDiv)
		case "ALTERNATIVE_ID":
			FormDivSelectFieldToField(&(_instance.ALTERNATIVE_ID), probe.stageOfInterest, formDiv)
		case "DEFAULT_VALUE":
			FormDivSelectFieldToField(&(_instance.DEFAULT_VALUE), probe.stageOfInterest, formDiv)
		case "TYPE":
			FormDivSelectFieldToField(&(_instance.TYPE), probe.stageOfInterest, formDiv)
		case "A_SPEC_ATTRIBUTES:ATTRIBUTE_DEFINITION_XHTML":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "ATTRIBUTE_DEFINITION_XHTML", func(owner *models.A_SPEC_ATTRIBUTES) *[]*models.ATTRIBUTE_DEFINITION_XHTML { return &owner.ATTRIBUTE_DEFINITION_XHTML })
		}
	}
}

func __gong__New__ATTRIBUTE_DEFINITION_XHTML_RenderingFormCallback(
	_instance *models.ATTRIBUTE_DEFINITION_XHTML_Rendering,
	probe *Probe,
	formGroup *form.FormGroup,
) (attribute_definition_xhtml_renderingFormCallback *FormCallback[*models.ATTRIBUTE_DEFINITION_XHTML_Rendering]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveATTRIBUTE_DEFINITION_XHTML_RenderingFields,
	)
}

type ATTRIBUTE_DEFINITION_XHTML_RenderingFormCallback = FormCallback[*models.ATTRIBUTE_DEFINITION_XHTML_Rendering]

func saveATTRIBUTE_DEFINITION_XHTML_RenderingFields(
	_instance *models.ATTRIBUTE_DEFINITION_XHTML_Rendering,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "ShowInTable":
			FormDivBasicFieldToField(&(_instance.ShowInTable), formDiv)
		case "ShowInTitle":
			FormDivBasicFieldToField(&(_instance.ShowInTitle), formDiv)
		case "ShowInSubject":
			FormDivBasicFieldToField(&(_instance.ShowInSubject), formDiv)
		case "Rank":
			FormDivBasicFieldToField(&(_instance.Rank), formDiv)
		}
	}
}

func __gong__New__ATTRIBUTE_VALUE_BOOLEANFormCallback(
	_instance *models.ATTRIBUTE_VALUE_BOOLEAN,
	probe *Probe,
	formGroup *form.FormGroup,
) (attribute_value_booleanFormCallback *FormCallback[*models.ATTRIBUTE_VALUE_BOOLEAN]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveATTRIBUTE_VALUE_BOOLEANFields,
	)
}

type ATTRIBUTE_VALUE_BOOLEANFormCallback = FormCallback[*models.ATTRIBUTE_VALUE_BOOLEAN]

func saveATTRIBUTE_VALUE_BOOLEANFields(
	_instance *models.ATTRIBUTE_VALUE_BOOLEAN,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "DEFINITION":
			FormDivSelectFieldToField(&(_instance.DEFINITION), probe.stageOfInterest, formDiv)
		case "THE_VALUE":
			FormDivBasicFieldToField(&(_instance.THE_VALUE), formDiv)
		case "A_ATTRIBUTE_VALUE_BOOLEAN:ATTRIBUTE_VALUE_BOOLEAN":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "ATTRIBUTE_VALUE_BOOLEAN", func(owner *models.A_ATTRIBUTE_VALUE_BOOLEAN) *[]*models.ATTRIBUTE_VALUE_BOOLEAN { return &owner.ATTRIBUTE_VALUE_BOOLEAN })
		case "A_ATTRIBUTE_VALUE_XHTML_1:ATTRIBUTE_VALUE_BOOLEAN":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "ATTRIBUTE_VALUE_BOOLEAN", func(owner *models.A_ATTRIBUTE_VALUE_XHTML_1) *[]*models.ATTRIBUTE_VALUE_BOOLEAN { return &owner.ATTRIBUTE_VALUE_BOOLEAN })
		}
	}
}

func __gong__New__ATTRIBUTE_VALUE_DATEFormCallback(
	_instance *models.ATTRIBUTE_VALUE_DATE,
	probe *Probe,
	formGroup *form.FormGroup,
) (attribute_value_dateFormCallback *FormCallback[*models.ATTRIBUTE_VALUE_DATE]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveATTRIBUTE_VALUE_DATEFields,
	)
}

type ATTRIBUTE_VALUE_DATEFormCallback = FormCallback[*models.ATTRIBUTE_VALUE_DATE]

func saveATTRIBUTE_VALUE_DATEFields(
	_instance *models.ATTRIBUTE_VALUE_DATE,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "DEFINITION":
			FormDivSelectFieldToField(&(_instance.DEFINITION), probe.stageOfInterest, formDiv)
		case "THE_VALUE":
			FormDivBasicFieldToField(&(_instance.THE_VALUE), formDiv)
		case "A_ATTRIBUTE_VALUE_DATE:ATTRIBUTE_VALUE_DATE":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "ATTRIBUTE_VALUE_DATE", func(owner *models.A_ATTRIBUTE_VALUE_DATE) *[]*models.ATTRIBUTE_VALUE_DATE { return &owner.ATTRIBUTE_VALUE_DATE })
		case "A_ATTRIBUTE_VALUE_XHTML_1:ATTRIBUTE_VALUE_DATE":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "ATTRIBUTE_VALUE_DATE", func(owner *models.A_ATTRIBUTE_VALUE_XHTML_1) *[]*models.ATTRIBUTE_VALUE_DATE { return &owner.ATTRIBUTE_VALUE_DATE })
		}
	}
}

func __gong__New__ATTRIBUTE_VALUE_ENUMERATIONFormCallback(
	_instance *models.ATTRIBUTE_VALUE_ENUMERATION,
	probe *Probe,
	formGroup *form.FormGroup,
) (attribute_value_enumerationFormCallback *FormCallback[*models.ATTRIBUTE_VALUE_ENUMERATION]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveATTRIBUTE_VALUE_ENUMERATIONFields,
	)
}

type ATTRIBUTE_VALUE_ENUMERATIONFormCallback = FormCallback[*models.ATTRIBUTE_VALUE_ENUMERATION]

func saveATTRIBUTE_VALUE_ENUMERATIONFields(
	_instance *models.ATTRIBUTE_VALUE_ENUMERATION,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "DEFINITION":
			FormDivSelectFieldToField(&(_instance.DEFINITION), probe.stageOfInterest, formDiv)
		case "VALUES":
			FormDivSelectFieldToField(&(_instance.VALUES), probe.stageOfInterest, formDiv)
		case "A_ATTRIBUTE_VALUE_ENUMERATION:ATTRIBUTE_VALUE_ENUMERATION":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "ATTRIBUTE_VALUE_ENUMERATION", func(owner *models.A_ATTRIBUTE_VALUE_ENUMERATION) *[]*models.ATTRIBUTE_VALUE_ENUMERATION { return &owner.ATTRIBUTE_VALUE_ENUMERATION })
		case "A_ATTRIBUTE_VALUE_XHTML_1:ATTRIBUTE_VALUE_ENUMERATION":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "ATTRIBUTE_VALUE_ENUMERATION", func(owner *models.A_ATTRIBUTE_VALUE_XHTML_1) *[]*models.ATTRIBUTE_VALUE_ENUMERATION { return &owner.ATTRIBUTE_VALUE_ENUMERATION })
		}
	}
}

func __gong__New__ATTRIBUTE_VALUE_INTEGERFormCallback(
	_instance *models.ATTRIBUTE_VALUE_INTEGER,
	probe *Probe,
	formGroup *form.FormGroup,
) (attribute_value_integerFormCallback *FormCallback[*models.ATTRIBUTE_VALUE_INTEGER]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveATTRIBUTE_VALUE_INTEGERFields,
	)
}

type ATTRIBUTE_VALUE_INTEGERFormCallback = FormCallback[*models.ATTRIBUTE_VALUE_INTEGER]

func saveATTRIBUTE_VALUE_INTEGERFields(
	_instance *models.ATTRIBUTE_VALUE_INTEGER,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "DEFINITION":
			FormDivSelectFieldToField(&(_instance.DEFINITION), probe.stageOfInterest, formDiv)
		case "THE_VALUE":
			FormDivBasicFieldToField(&(_instance.THE_VALUE), formDiv)
		case "A_ATTRIBUTE_VALUE_INTEGER:ATTRIBUTE_VALUE_INTEGER":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "ATTRIBUTE_VALUE_INTEGER", func(owner *models.A_ATTRIBUTE_VALUE_INTEGER) *[]*models.ATTRIBUTE_VALUE_INTEGER { return &owner.ATTRIBUTE_VALUE_INTEGER })
		case "A_ATTRIBUTE_VALUE_XHTML_1:ATTRIBUTE_VALUE_INTEGER":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "ATTRIBUTE_VALUE_INTEGER", func(owner *models.A_ATTRIBUTE_VALUE_XHTML_1) *[]*models.ATTRIBUTE_VALUE_INTEGER { return &owner.ATTRIBUTE_VALUE_INTEGER })
		}
	}
}

func __gong__New__ATTRIBUTE_VALUE_REALFormCallback(
	_instance *models.ATTRIBUTE_VALUE_REAL,
	probe *Probe,
	formGroup *form.FormGroup,
) (attribute_value_realFormCallback *FormCallback[*models.ATTRIBUTE_VALUE_REAL]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveATTRIBUTE_VALUE_REALFields,
	)
}

type ATTRIBUTE_VALUE_REALFormCallback = FormCallback[*models.ATTRIBUTE_VALUE_REAL]

func saveATTRIBUTE_VALUE_REALFields(
	_instance *models.ATTRIBUTE_VALUE_REAL,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "DEFINITION":
			FormDivSelectFieldToField(&(_instance.DEFINITION), probe.stageOfInterest, formDiv)
		case "THE_VALUE":
			FormDivBasicFieldToField(&(_instance.THE_VALUE), formDiv)
		case "A_ATTRIBUTE_VALUE_REAL:ATTRIBUTE_VALUE_REAL":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "ATTRIBUTE_VALUE_REAL", func(owner *models.A_ATTRIBUTE_VALUE_REAL) *[]*models.ATTRIBUTE_VALUE_REAL { return &owner.ATTRIBUTE_VALUE_REAL })
		case "A_ATTRIBUTE_VALUE_XHTML_1:ATTRIBUTE_VALUE_REAL":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "ATTRIBUTE_VALUE_REAL", func(owner *models.A_ATTRIBUTE_VALUE_XHTML_1) *[]*models.ATTRIBUTE_VALUE_REAL { return &owner.ATTRIBUTE_VALUE_REAL })
		}
	}
}

func __gong__New__ATTRIBUTE_VALUE_STRINGFormCallback(
	_instance *models.ATTRIBUTE_VALUE_STRING,
	probe *Probe,
	formGroup *form.FormGroup,
) (attribute_value_stringFormCallback *FormCallback[*models.ATTRIBUTE_VALUE_STRING]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveATTRIBUTE_VALUE_STRINGFields,
	)
}

type ATTRIBUTE_VALUE_STRINGFormCallback = FormCallback[*models.ATTRIBUTE_VALUE_STRING]

func saveATTRIBUTE_VALUE_STRINGFields(
	_instance *models.ATTRIBUTE_VALUE_STRING,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "DEFINITION":
			FormDivSelectFieldToField(&(_instance.DEFINITION), probe.stageOfInterest, formDiv)
		case "THE_VALUE":
			FormDivBasicFieldToField(&(_instance.THE_VALUE), formDiv)
		case "A_ATTRIBUTE_VALUE_STRING:ATTRIBUTE_VALUE_STRING":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "ATTRIBUTE_VALUE_STRING", func(owner *models.A_ATTRIBUTE_VALUE_STRING) *[]*models.ATTRIBUTE_VALUE_STRING { return &owner.ATTRIBUTE_VALUE_STRING })
		case "A_ATTRIBUTE_VALUE_XHTML_1:ATTRIBUTE_VALUE_STRING":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "ATTRIBUTE_VALUE_STRING", func(owner *models.A_ATTRIBUTE_VALUE_XHTML_1) *[]*models.ATTRIBUTE_VALUE_STRING { return &owner.ATTRIBUTE_VALUE_STRING })
		}
	}
}

func __gong__New__ATTRIBUTE_VALUE_XHTMLFormCallback(
	_instance *models.ATTRIBUTE_VALUE_XHTML,
	probe *Probe,
	formGroup *form.FormGroup,
) (attribute_value_xhtmlFormCallback *FormCallback[*models.ATTRIBUTE_VALUE_XHTML]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveATTRIBUTE_VALUE_XHTMLFields,
	)
}

type ATTRIBUTE_VALUE_XHTMLFormCallback = FormCallback[*models.ATTRIBUTE_VALUE_XHTML]

func saveATTRIBUTE_VALUE_XHTMLFields(
	_instance *models.ATTRIBUTE_VALUE_XHTML,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "DEFINITION":
			FormDivSelectFieldToField(&(_instance.DEFINITION), probe.stageOfInterest, formDiv)
		case "IS_SIMPLIFIED":
			FormDivBasicFieldToField(&(_instance.IS_SIMPLIFIED), formDiv)
		case "THE_VALUE":
			FormDivSelectFieldToField(&(_instance.THE_VALUE), probe.stageOfInterest, formDiv)
		case "THE_ORIGINAL_VALUE":
			FormDivSelectFieldToField(&(_instance.THE_ORIGINAL_VALUE), probe.stageOfInterest, formDiv)
		case "A_ATTRIBUTE_VALUE_XHTML:ATTRIBUTE_VALUE_XHTML":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "ATTRIBUTE_VALUE_XHTML", func(owner *models.A_ATTRIBUTE_VALUE_XHTML) *[]*models.ATTRIBUTE_VALUE_XHTML { return &owner.ATTRIBUTE_VALUE_XHTML })
		case "A_ATTRIBUTE_VALUE_XHTML_1:ATTRIBUTE_VALUE_XHTML":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "ATTRIBUTE_VALUE_XHTML", func(owner *models.A_ATTRIBUTE_VALUE_XHTML_1) *[]*models.ATTRIBUTE_VALUE_XHTML { return &owner.ATTRIBUTE_VALUE_XHTML })
		}
	}
}

func __gong__New__A_ALTERNATIVE_IDFormCallback(
	_instance *models.A_ALTERNATIVE_ID,
	probe *Probe,
	formGroup *form.FormGroup,
) (a_alternative_idFormCallback *FormCallback[*models.A_ALTERNATIVE_ID]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveA_ALTERNATIVE_IDFields,
	)
}

type A_ALTERNATIVE_IDFormCallback = FormCallback[*models.A_ALTERNATIVE_ID]

func saveA_ALTERNATIVE_IDFields(
	_instance *models.A_ALTERNATIVE_ID,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "ALTERNATIVE_ID":
			FormDivSelectFieldToField(&(_instance.ALTERNATIVE_ID), probe.stageOfInterest, formDiv)
		}
	}
}

func __gong__New__A_ATTRIBUTE_DEFINITION_BOOLEAN_REFFormCallback(
	_instance *models.A_ATTRIBUTE_DEFINITION_BOOLEAN_REF,
	probe *Probe,
	formGroup *form.FormGroup,
) (a_attribute_definition_boolean_refFormCallback *FormCallback[*models.A_ATTRIBUTE_DEFINITION_BOOLEAN_REF]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveA_ATTRIBUTE_DEFINITION_BOOLEAN_REFFields,
	)
}

type A_ATTRIBUTE_DEFINITION_BOOLEAN_REFFormCallback = FormCallback[*models.A_ATTRIBUTE_DEFINITION_BOOLEAN_REF]

func saveA_ATTRIBUTE_DEFINITION_BOOLEAN_REFFields(
	_instance *models.A_ATTRIBUTE_DEFINITION_BOOLEAN_REF,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "ATTRIBUTE_DEFINITION_BOOLEAN_REF":
			FormDivBasicFieldToField(&(_instance.ATTRIBUTE_DEFINITION_BOOLEAN_REF), formDiv)
		}
	}
}

func __gong__New__A_ATTRIBUTE_DEFINITION_DATE_REFFormCallback(
	_instance *models.A_ATTRIBUTE_DEFINITION_DATE_REF,
	probe *Probe,
	formGroup *form.FormGroup,
) (a_attribute_definition_date_refFormCallback *FormCallback[*models.A_ATTRIBUTE_DEFINITION_DATE_REF]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveA_ATTRIBUTE_DEFINITION_DATE_REFFields,
	)
}

type A_ATTRIBUTE_DEFINITION_DATE_REFFormCallback = FormCallback[*models.A_ATTRIBUTE_DEFINITION_DATE_REF]

func saveA_ATTRIBUTE_DEFINITION_DATE_REFFields(
	_instance *models.A_ATTRIBUTE_DEFINITION_DATE_REF,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "ATTRIBUTE_DEFINITION_DATE_REF":
			FormDivBasicFieldToField(&(_instance.ATTRIBUTE_DEFINITION_DATE_REF), formDiv)
		}
	}
}

func __gong__New__A_ATTRIBUTE_DEFINITION_ENUMERATION_REFFormCallback(
	_instance *models.A_ATTRIBUTE_DEFINITION_ENUMERATION_REF,
	probe *Probe,
	formGroup *form.FormGroup,
) (a_attribute_definition_enumeration_refFormCallback *FormCallback[*models.A_ATTRIBUTE_DEFINITION_ENUMERATION_REF]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveA_ATTRIBUTE_DEFINITION_ENUMERATION_REFFields,
	)
}

type A_ATTRIBUTE_DEFINITION_ENUMERATION_REFFormCallback = FormCallback[*models.A_ATTRIBUTE_DEFINITION_ENUMERATION_REF]

func saveA_ATTRIBUTE_DEFINITION_ENUMERATION_REFFields(
	_instance *models.A_ATTRIBUTE_DEFINITION_ENUMERATION_REF,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "ATTRIBUTE_DEFINITION_ENUMERATION_REF":
			FormDivBasicFieldToField(&(_instance.ATTRIBUTE_DEFINITION_ENUMERATION_REF), formDiv)
		}
	}
}

func __gong__New__A_ATTRIBUTE_DEFINITION_INTEGER_REFFormCallback(
	_instance *models.A_ATTRIBUTE_DEFINITION_INTEGER_REF,
	probe *Probe,
	formGroup *form.FormGroup,
) (a_attribute_definition_integer_refFormCallback *FormCallback[*models.A_ATTRIBUTE_DEFINITION_INTEGER_REF]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveA_ATTRIBUTE_DEFINITION_INTEGER_REFFields,
	)
}

type A_ATTRIBUTE_DEFINITION_INTEGER_REFFormCallback = FormCallback[*models.A_ATTRIBUTE_DEFINITION_INTEGER_REF]

func saveA_ATTRIBUTE_DEFINITION_INTEGER_REFFields(
	_instance *models.A_ATTRIBUTE_DEFINITION_INTEGER_REF,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "ATTRIBUTE_DEFINITION_INTEGER_REF":
			FormDivBasicFieldToField(&(_instance.ATTRIBUTE_DEFINITION_INTEGER_REF), formDiv)
		}
	}
}

func __gong__New__A_ATTRIBUTE_DEFINITION_REAL_REFFormCallback(
	_instance *models.A_ATTRIBUTE_DEFINITION_REAL_REF,
	probe *Probe,
	formGroup *form.FormGroup,
) (a_attribute_definition_real_refFormCallback *FormCallback[*models.A_ATTRIBUTE_DEFINITION_REAL_REF]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveA_ATTRIBUTE_DEFINITION_REAL_REFFields,
	)
}

type A_ATTRIBUTE_DEFINITION_REAL_REFFormCallback = FormCallback[*models.A_ATTRIBUTE_DEFINITION_REAL_REF]

func saveA_ATTRIBUTE_DEFINITION_REAL_REFFields(
	_instance *models.A_ATTRIBUTE_DEFINITION_REAL_REF,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "ATTRIBUTE_DEFINITION_REAL_REF":
			FormDivBasicFieldToField(&(_instance.ATTRIBUTE_DEFINITION_REAL_REF), formDiv)
		}
	}
}

func __gong__New__A_ATTRIBUTE_DEFINITION_STRING_REFFormCallback(
	_instance *models.A_ATTRIBUTE_DEFINITION_STRING_REF,
	probe *Probe,
	formGroup *form.FormGroup,
) (a_attribute_definition_string_refFormCallback *FormCallback[*models.A_ATTRIBUTE_DEFINITION_STRING_REF]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveA_ATTRIBUTE_DEFINITION_STRING_REFFields,
	)
}

type A_ATTRIBUTE_DEFINITION_STRING_REFFormCallback = FormCallback[*models.A_ATTRIBUTE_DEFINITION_STRING_REF]

func saveA_ATTRIBUTE_DEFINITION_STRING_REFFields(
	_instance *models.A_ATTRIBUTE_DEFINITION_STRING_REF,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "ATTRIBUTE_DEFINITION_STRING_REF":
			FormDivBasicFieldToField(&(_instance.ATTRIBUTE_DEFINITION_STRING_REF), formDiv)
		}
	}
}

func __gong__New__A_ATTRIBUTE_DEFINITION_XHTML_REFFormCallback(
	_instance *models.A_ATTRIBUTE_DEFINITION_XHTML_REF,
	probe *Probe,
	formGroup *form.FormGroup,
) (a_attribute_definition_xhtml_refFormCallback *FormCallback[*models.A_ATTRIBUTE_DEFINITION_XHTML_REF]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveA_ATTRIBUTE_DEFINITION_XHTML_REFFields,
	)
}

type A_ATTRIBUTE_DEFINITION_XHTML_REFFormCallback = FormCallback[*models.A_ATTRIBUTE_DEFINITION_XHTML_REF]

func saveA_ATTRIBUTE_DEFINITION_XHTML_REFFields(
	_instance *models.A_ATTRIBUTE_DEFINITION_XHTML_REF,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "ATTRIBUTE_DEFINITION_XHTML_REF":
			FormDivBasicFieldToField(&(_instance.ATTRIBUTE_DEFINITION_XHTML_REF), formDiv)
		}
	}
}

func __gong__New__A_ATTRIBUTE_VALUE_BOOLEANFormCallback(
	_instance *models.A_ATTRIBUTE_VALUE_BOOLEAN,
	probe *Probe,
	formGroup *form.FormGroup,
) (a_attribute_value_booleanFormCallback *FormCallback[*models.A_ATTRIBUTE_VALUE_BOOLEAN]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveA_ATTRIBUTE_VALUE_BOOLEANFields,
	)
}

type A_ATTRIBUTE_VALUE_BOOLEANFormCallback = FormCallback[*models.A_ATTRIBUTE_VALUE_BOOLEAN]

func saveA_ATTRIBUTE_VALUE_BOOLEANFields(
	_instance *models.A_ATTRIBUTE_VALUE_BOOLEAN,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "ATTRIBUTE_VALUE_BOOLEAN":
			FormDivSliceOfPointersToField(_instance, "ATTRIBUTE_VALUE_BOOLEAN", &(_instance.ATTRIBUTE_VALUE_BOOLEAN), formDiv, probe)
		}
	}
}

func __gong__New__A_ATTRIBUTE_VALUE_DATEFormCallback(
	_instance *models.A_ATTRIBUTE_VALUE_DATE,
	probe *Probe,
	formGroup *form.FormGroup,
) (a_attribute_value_dateFormCallback *FormCallback[*models.A_ATTRIBUTE_VALUE_DATE]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveA_ATTRIBUTE_VALUE_DATEFields,
	)
}

type A_ATTRIBUTE_VALUE_DATEFormCallback = FormCallback[*models.A_ATTRIBUTE_VALUE_DATE]

func saveA_ATTRIBUTE_VALUE_DATEFields(
	_instance *models.A_ATTRIBUTE_VALUE_DATE,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "ATTRIBUTE_VALUE_DATE":
			FormDivSliceOfPointersToField(_instance, "ATTRIBUTE_VALUE_DATE", &(_instance.ATTRIBUTE_VALUE_DATE), formDiv, probe)
		}
	}
}

func __gong__New__A_ATTRIBUTE_VALUE_ENUMERATIONFormCallback(
	_instance *models.A_ATTRIBUTE_VALUE_ENUMERATION,
	probe *Probe,
	formGroup *form.FormGroup,
) (a_attribute_value_enumerationFormCallback *FormCallback[*models.A_ATTRIBUTE_VALUE_ENUMERATION]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveA_ATTRIBUTE_VALUE_ENUMERATIONFields,
	)
}

type A_ATTRIBUTE_VALUE_ENUMERATIONFormCallback = FormCallback[*models.A_ATTRIBUTE_VALUE_ENUMERATION]

func saveA_ATTRIBUTE_VALUE_ENUMERATIONFields(
	_instance *models.A_ATTRIBUTE_VALUE_ENUMERATION,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "ATTRIBUTE_VALUE_ENUMERATION":
			FormDivSliceOfPointersToField(_instance, "ATTRIBUTE_VALUE_ENUMERATION", &(_instance.ATTRIBUTE_VALUE_ENUMERATION), formDiv, probe)
		}
	}
}

func __gong__New__A_ATTRIBUTE_VALUE_INTEGERFormCallback(
	_instance *models.A_ATTRIBUTE_VALUE_INTEGER,
	probe *Probe,
	formGroup *form.FormGroup,
) (a_attribute_value_integerFormCallback *FormCallback[*models.A_ATTRIBUTE_VALUE_INTEGER]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveA_ATTRIBUTE_VALUE_INTEGERFields,
	)
}

type A_ATTRIBUTE_VALUE_INTEGERFormCallback = FormCallback[*models.A_ATTRIBUTE_VALUE_INTEGER]

func saveA_ATTRIBUTE_VALUE_INTEGERFields(
	_instance *models.A_ATTRIBUTE_VALUE_INTEGER,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "ATTRIBUTE_VALUE_INTEGER":
			FormDivSliceOfPointersToField(_instance, "ATTRIBUTE_VALUE_INTEGER", &(_instance.ATTRIBUTE_VALUE_INTEGER), formDiv, probe)
		}
	}
}

func __gong__New__A_ATTRIBUTE_VALUE_REALFormCallback(
	_instance *models.A_ATTRIBUTE_VALUE_REAL,
	probe *Probe,
	formGroup *form.FormGroup,
) (a_attribute_value_realFormCallback *FormCallback[*models.A_ATTRIBUTE_VALUE_REAL]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveA_ATTRIBUTE_VALUE_REALFields,
	)
}

type A_ATTRIBUTE_VALUE_REALFormCallback = FormCallback[*models.A_ATTRIBUTE_VALUE_REAL]

func saveA_ATTRIBUTE_VALUE_REALFields(
	_instance *models.A_ATTRIBUTE_VALUE_REAL,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "ATTRIBUTE_VALUE_REAL":
			FormDivSliceOfPointersToField(_instance, "ATTRIBUTE_VALUE_REAL", &(_instance.ATTRIBUTE_VALUE_REAL), formDiv, probe)
		}
	}
}

func __gong__New__A_ATTRIBUTE_VALUE_STRINGFormCallback(
	_instance *models.A_ATTRIBUTE_VALUE_STRING,
	probe *Probe,
	formGroup *form.FormGroup,
) (a_attribute_value_stringFormCallback *FormCallback[*models.A_ATTRIBUTE_VALUE_STRING]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveA_ATTRIBUTE_VALUE_STRINGFields,
	)
}

type A_ATTRIBUTE_VALUE_STRINGFormCallback = FormCallback[*models.A_ATTRIBUTE_VALUE_STRING]

func saveA_ATTRIBUTE_VALUE_STRINGFields(
	_instance *models.A_ATTRIBUTE_VALUE_STRING,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "ATTRIBUTE_VALUE_STRING":
			FormDivSliceOfPointersToField(_instance, "ATTRIBUTE_VALUE_STRING", &(_instance.ATTRIBUTE_VALUE_STRING), formDiv, probe)
		}
	}
}

func __gong__New__A_ATTRIBUTE_VALUE_XHTMLFormCallback(
	_instance *models.A_ATTRIBUTE_VALUE_XHTML,
	probe *Probe,
	formGroup *form.FormGroup,
) (a_attribute_value_xhtmlFormCallback *FormCallback[*models.A_ATTRIBUTE_VALUE_XHTML]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveA_ATTRIBUTE_VALUE_XHTMLFields,
	)
}

type A_ATTRIBUTE_VALUE_XHTMLFormCallback = FormCallback[*models.A_ATTRIBUTE_VALUE_XHTML]

func saveA_ATTRIBUTE_VALUE_XHTMLFields(
	_instance *models.A_ATTRIBUTE_VALUE_XHTML,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "ATTRIBUTE_VALUE_XHTML":
			FormDivSliceOfPointersToField(_instance, "ATTRIBUTE_VALUE_XHTML", &(_instance.ATTRIBUTE_VALUE_XHTML), formDiv, probe)
		}
	}
}

func __gong__New__A_ATTRIBUTE_VALUE_XHTML_1FormCallback(
	_instance *models.A_ATTRIBUTE_VALUE_XHTML_1,
	probe *Probe,
	formGroup *form.FormGroup,
) (a_attribute_value_xhtml_1FormCallback *FormCallback[*models.A_ATTRIBUTE_VALUE_XHTML_1]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveA_ATTRIBUTE_VALUE_XHTML_1Fields,
	)
}

type A_ATTRIBUTE_VALUE_XHTML_1FormCallback = FormCallback[*models.A_ATTRIBUTE_VALUE_XHTML_1]

func saveA_ATTRIBUTE_VALUE_XHTML_1Fields(
	_instance *models.A_ATTRIBUTE_VALUE_XHTML_1,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "ATTRIBUTE_VALUE_BOOLEAN":
			FormDivSliceOfPointersToField(_instance, "ATTRIBUTE_VALUE_BOOLEAN", &(_instance.ATTRIBUTE_VALUE_BOOLEAN), formDiv, probe)
		case "ATTRIBUTE_VALUE_DATE":
			FormDivSliceOfPointersToField(_instance, "ATTRIBUTE_VALUE_DATE", &(_instance.ATTRIBUTE_VALUE_DATE), formDiv, probe)
		case "ATTRIBUTE_VALUE_ENUMERATION":
			FormDivSliceOfPointersToField(_instance, "ATTRIBUTE_VALUE_ENUMERATION", &(_instance.ATTRIBUTE_VALUE_ENUMERATION), formDiv, probe)
		case "ATTRIBUTE_VALUE_INTEGER":
			FormDivSliceOfPointersToField(_instance, "ATTRIBUTE_VALUE_INTEGER", &(_instance.ATTRIBUTE_VALUE_INTEGER), formDiv, probe)
		case "ATTRIBUTE_VALUE_REAL":
			FormDivSliceOfPointersToField(_instance, "ATTRIBUTE_VALUE_REAL", &(_instance.ATTRIBUTE_VALUE_REAL), formDiv, probe)
		case "ATTRIBUTE_VALUE_STRING":
			FormDivSliceOfPointersToField(_instance, "ATTRIBUTE_VALUE_STRING", &(_instance.ATTRIBUTE_VALUE_STRING), formDiv, probe)
		case "ATTRIBUTE_VALUE_XHTML":
			FormDivSliceOfPointersToField(_instance, "ATTRIBUTE_VALUE_XHTML", &(_instance.ATTRIBUTE_VALUE_XHTML), formDiv, probe)
		}
	}
}

func __gong__New__A_CHILDRENFormCallback(
	_instance *models.A_CHILDREN,
	probe *Probe,
	formGroup *form.FormGroup,
) (a_childrenFormCallback *FormCallback[*models.A_CHILDREN]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveA_CHILDRENFields,
	)
}

type A_CHILDRENFormCallback = FormCallback[*models.A_CHILDREN]

func saveA_CHILDRENFields(
	_instance *models.A_CHILDREN,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "SPEC_HIERARCHY":
			FormDivSliceOfPointersToField(_instance, "SPEC_HIERARCHY", &(_instance.SPEC_HIERARCHY), formDiv, probe)
		}
	}
}

func __gong__New__A_CORE_CONTENTFormCallback(
	_instance *models.A_CORE_CONTENT,
	probe *Probe,
	formGroup *form.FormGroup,
) (a_core_contentFormCallback *FormCallback[*models.A_CORE_CONTENT]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveA_CORE_CONTENTFields,
	)
}

type A_CORE_CONTENTFormCallback = FormCallback[*models.A_CORE_CONTENT]

func saveA_CORE_CONTENTFields(
	_instance *models.A_CORE_CONTENT,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "REQ_IF_CONTENT":
			FormDivSelectFieldToField(&(_instance.REQ_IF_CONTENT), probe.stageOfInterest, formDiv)
		}
	}
}

func __gong__New__A_DATATYPESFormCallback(
	_instance *models.A_DATATYPES,
	probe *Probe,
	formGroup *form.FormGroup,
) (a_datatypesFormCallback *FormCallback[*models.A_DATATYPES]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveA_DATATYPESFields,
	)
}

type A_DATATYPESFormCallback = FormCallback[*models.A_DATATYPES]

func saveA_DATATYPESFields(
	_instance *models.A_DATATYPES,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "DATATYPE_DEFINITION_BOOLEAN":
			FormDivSliceOfPointersToField(_instance, "DATATYPE_DEFINITION_BOOLEAN", &(_instance.DATATYPE_DEFINITION_BOOLEAN), formDiv, probe)
		case "DATATYPE_DEFINITION_DATE":
			FormDivSliceOfPointersToField(_instance, "DATATYPE_DEFINITION_DATE", &(_instance.DATATYPE_DEFINITION_DATE), formDiv, probe)
		case "DATATYPE_DEFINITION_ENUMERATION":
			FormDivSliceOfPointersToField(_instance, "DATATYPE_DEFINITION_ENUMERATION", &(_instance.DATATYPE_DEFINITION_ENUMERATION), formDiv, probe)
		case "DATATYPE_DEFINITION_INTEGER":
			FormDivSliceOfPointersToField(_instance, "DATATYPE_DEFINITION_INTEGER", &(_instance.DATATYPE_DEFINITION_INTEGER), formDiv, probe)
		case "DATATYPE_DEFINITION_REAL":
			FormDivSliceOfPointersToField(_instance, "DATATYPE_DEFINITION_REAL", &(_instance.DATATYPE_DEFINITION_REAL), formDiv, probe)
		case "DATATYPE_DEFINITION_STRING":
			FormDivSliceOfPointersToField(_instance, "DATATYPE_DEFINITION_STRING", &(_instance.DATATYPE_DEFINITION_STRING), formDiv, probe)
		case "DATATYPE_DEFINITION_XHTML":
			FormDivSliceOfPointersToField(_instance, "DATATYPE_DEFINITION_XHTML", &(_instance.DATATYPE_DEFINITION_XHTML), formDiv, probe)
		}
	}
}

func __gong__New__A_DATATYPE_DEFINITION_BOOLEAN_REFFormCallback(
	_instance *models.A_DATATYPE_DEFINITION_BOOLEAN_REF,
	probe *Probe,
	formGroup *form.FormGroup,
) (a_datatype_definition_boolean_refFormCallback *FormCallback[*models.A_DATATYPE_DEFINITION_BOOLEAN_REF]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveA_DATATYPE_DEFINITION_BOOLEAN_REFFields,
	)
}

type A_DATATYPE_DEFINITION_BOOLEAN_REFFormCallback = FormCallback[*models.A_DATATYPE_DEFINITION_BOOLEAN_REF]

func saveA_DATATYPE_DEFINITION_BOOLEAN_REFFields(
	_instance *models.A_DATATYPE_DEFINITION_BOOLEAN_REF,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "DATATYPE_DEFINITION_BOOLEAN_REF":
			FormDivBasicFieldToField(&(_instance.DATATYPE_DEFINITION_BOOLEAN_REF), formDiv)
		}
	}
}

func __gong__New__A_DATATYPE_DEFINITION_DATE_REFFormCallback(
	_instance *models.A_DATATYPE_DEFINITION_DATE_REF,
	probe *Probe,
	formGroup *form.FormGroup,
) (a_datatype_definition_date_refFormCallback *FormCallback[*models.A_DATATYPE_DEFINITION_DATE_REF]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveA_DATATYPE_DEFINITION_DATE_REFFields,
	)
}

type A_DATATYPE_DEFINITION_DATE_REFFormCallback = FormCallback[*models.A_DATATYPE_DEFINITION_DATE_REF]

func saveA_DATATYPE_DEFINITION_DATE_REFFields(
	_instance *models.A_DATATYPE_DEFINITION_DATE_REF,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "DATATYPE_DEFINITION_DATE_REF":
			FormDivBasicFieldToField(&(_instance.DATATYPE_DEFINITION_DATE_REF), formDiv)
		}
	}
}

func __gong__New__A_DATATYPE_DEFINITION_ENUMERATION_REFFormCallback(
	_instance *models.A_DATATYPE_DEFINITION_ENUMERATION_REF,
	probe *Probe,
	formGroup *form.FormGroup,
) (a_datatype_definition_enumeration_refFormCallback *FormCallback[*models.A_DATATYPE_DEFINITION_ENUMERATION_REF]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveA_DATATYPE_DEFINITION_ENUMERATION_REFFields,
	)
}

type A_DATATYPE_DEFINITION_ENUMERATION_REFFormCallback = FormCallback[*models.A_DATATYPE_DEFINITION_ENUMERATION_REF]

func saveA_DATATYPE_DEFINITION_ENUMERATION_REFFields(
	_instance *models.A_DATATYPE_DEFINITION_ENUMERATION_REF,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "DATATYPE_DEFINITION_ENUMERATION_REF":
			FormDivBasicFieldToField(&(_instance.DATATYPE_DEFINITION_ENUMERATION_REF), formDiv)
		}
	}
}

func __gong__New__A_DATATYPE_DEFINITION_INTEGER_REFFormCallback(
	_instance *models.A_DATATYPE_DEFINITION_INTEGER_REF,
	probe *Probe,
	formGroup *form.FormGroup,
) (a_datatype_definition_integer_refFormCallback *FormCallback[*models.A_DATATYPE_DEFINITION_INTEGER_REF]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveA_DATATYPE_DEFINITION_INTEGER_REFFields,
	)
}

type A_DATATYPE_DEFINITION_INTEGER_REFFormCallback = FormCallback[*models.A_DATATYPE_DEFINITION_INTEGER_REF]

func saveA_DATATYPE_DEFINITION_INTEGER_REFFields(
	_instance *models.A_DATATYPE_DEFINITION_INTEGER_REF,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "DATATYPE_DEFINITION_INTEGER_REF":
			FormDivBasicFieldToField(&(_instance.DATATYPE_DEFINITION_INTEGER_REF), formDiv)
		}
	}
}

func __gong__New__A_DATATYPE_DEFINITION_REAL_REFFormCallback(
	_instance *models.A_DATATYPE_DEFINITION_REAL_REF,
	probe *Probe,
	formGroup *form.FormGroup,
) (a_datatype_definition_real_refFormCallback *FormCallback[*models.A_DATATYPE_DEFINITION_REAL_REF]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveA_DATATYPE_DEFINITION_REAL_REFFields,
	)
}

type A_DATATYPE_DEFINITION_REAL_REFFormCallback = FormCallback[*models.A_DATATYPE_DEFINITION_REAL_REF]

func saveA_DATATYPE_DEFINITION_REAL_REFFields(
	_instance *models.A_DATATYPE_DEFINITION_REAL_REF,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "DATATYPE_DEFINITION_REAL_REF":
			FormDivBasicFieldToField(&(_instance.DATATYPE_DEFINITION_REAL_REF), formDiv)
		}
	}
}

func __gong__New__A_DATATYPE_DEFINITION_STRING_REFFormCallback(
	_instance *models.A_DATATYPE_DEFINITION_STRING_REF,
	probe *Probe,
	formGroup *form.FormGroup,
) (a_datatype_definition_string_refFormCallback *FormCallback[*models.A_DATATYPE_DEFINITION_STRING_REF]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveA_DATATYPE_DEFINITION_STRING_REFFields,
	)
}

type A_DATATYPE_DEFINITION_STRING_REFFormCallback = FormCallback[*models.A_DATATYPE_DEFINITION_STRING_REF]

func saveA_DATATYPE_DEFINITION_STRING_REFFields(
	_instance *models.A_DATATYPE_DEFINITION_STRING_REF,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "DATATYPE_DEFINITION_STRING_REF":
			FormDivBasicFieldToField(&(_instance.DATATYPE_DEFINITION_STRING_REF), formDiv)
		}
	}
}

func __gong__New__A_DATATYPE_DEFINITION_XHTML_REFFormCallback(
	_instance *models.A_DATATYPE_DEFINITION_XHTML_REF,
	probe *Probe,
	formGroup *form.FormGroup,
) (a_datatype_definition_xhtml_refFormCallback *FormCallback[*models.A_DATATYPE_DEFINITION_XHTML_REF]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveA_DATATYPE_DEFINITION_XHTML_REFFields,
	)
}

type A_DATATYPE_DEFINITION_XHTML_REFFormCallback = FormCallback[*models.A_DATATYPE_DEFINITION_XHTML_REF]

func saveA_DATATYPE_DEFINITION_XHTML_REFFields(
	_instance *models.A_DATATYPE_DEFINITION_XHTML_REF,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "DATATYPE_DEFINITION_XHTML_REF":
			FormDivBasicFieldToField(&(_instance.DATATYPE_DEFINITION_XHTML_REF), formDiv)
		}
	}
}

func __gong__New__A_EDITABLE_ATTSFormCallback(
	_instance *models.A_EDITABLE_ATTS,
	probe *Probe,
	formGroup *form.FormGroup,
) (a_editable_attsFormCallback *FormCallback[*models.A_EDITABLE_ATTS]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveA_EDITABLE_ATTSFields,
	)
}

type A_EDITABLE_ATTSFormCallback = FormCallback[*models.A_EDITABLE_ATTS]

func saveA_EDITABLE_ATTSFields(
	_instance *models.A_EDITABLE_ATTS,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "ATTRIBUTE_DEFINITION_BOOLEAN_REF":
			FormDivBasicFieldToField(&(_instance.ATTRIBUTE_DEFINITION_BOOLEAN_REF), formDiv)
		case "ATTRIBUTE_DEFINITION_DATE_REF":
			FormDivBasicFieldToField(&(_instance.ATTRIBUTE_DEFINITION_DATE_REF), formDiv)
		case "ATTRIBUTE_DEFINITION_ENUMERATION_REF":
			FormDivBasicFieldToField(&(_instance.ATTRIBUTE_DEFINITION_ENUMERATION_REF), formDiv)
		case "ATTRIBUTE_DEFINITION_INTEGER_REF":
			FormDivBasicFieldToField(&(_instance.ATTRIBUTE_DEFINITION_INTEGER_REF), formDiv)
		case "ATTRIBUTE_DEFINITION_REAL_REF":
			FormDivBasicFieldToField(&(_instance.ATTRIBUTE_DEFINITION_REAL_REF), formDiv)
		case "ATTRIBUTE_DEFINITION_STRING_REF":
			FormDivBasicFieldToField(&(_instance.ATTRIBUTE_DEFINITION_STRING_REF), formDiv)
		case "ATTRIBUTE_DEFINITION_XHTML_REF":
			FormDivBasicFieldToField(&(_instance.ATTRIBUTE_DEFINITION_XHTML_REF), formDiv)
		}
	}
}

func __gong__New__A_ENUM_VALUE_REFFormCallback(
	_instance *models.A_ENUM_VALUE_REF,
	probe *Probe,
	formGroup *form.FormGroup,
) (a_enum_value_refFormCallback *FormCallback[*models.A_ENUM_VALUE_REF]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveA_ENUM_VALUE_REFFields,
	)
}

type A_ENUM_VALUE_REFFormCallback = FormCallback[*models.A_ENUM_VALUE_REF]

func saveA_ENUM_VALUE_REFFields(
	_instance *models.A_ENUM_VALUE_REF,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "ENUM_VALUE_REF":
			FormDivBasicFieldToField(&(_instance.ENUM_VALUE_REF), formDiv)
		}
	}
}

func __gong__New__A_OBJECTFormCallback(
	_instance *models.A_OBJECT,
	probe *Probe,
	formGroup *form.FormGroup,
) (a_objectFormCallback *FormCallback[*models.A_OBJECT]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveA_OBJECTFields,
	)
}

type A_OBJECTFormCallback = FormCallback[*models.A_OBJECT]

func saveA_OBJECTFields(
	_instance *models.A_OBJECT,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "SPEC_OBJECT_REF":
			FormDivBasicFieldToField(&(_instance.SPEC_OBJECT_REF), formDiv)
		}
	}
}

func __gong__New__A_PROPERTIESFormCallback(
	_instance *models.A_PROPERTIES,
	probe *Probe,
	formGroup *form.FormGroup,
) (a_propertiesFormCallback *FormCallback[*models.A_PROPERTIES]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveA_PROPERTIESFields,
	)
}

type A_PROPERTIESFormCallback = FormCallback[*models.A_PROPERTIES]

func saveA_PROPERTIESFields(
	_instance *models.A_PROPERTIES,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "EMBEDDED_VALUE":
			FormDivSelectFieldToField(&(_instance.EMBEDDED_VALUE), probe.stageOfInterest, formDiv)
		}
	}
}

func __gong__New__A_RELATION_GROUP_TYPE_REFFormCallback(
	_instance *models.A_RELATION_GROUP_TYPE_REF,
	probe *Probe,
	formGroup *form.FormGroup,
) (a_relation_group_type_refFormCallback *FormCallback[*models.A_RELATION_GROUP_TYPE_REF]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveA_RELATION_GROUP_TYPE_REFFields,
	)
}

type A_RELATION_GROUP_TYPE_REFFormCallback = FormCallback[*models.A_RELATION_GROUP_TYPE_REF]

func saveA_RELATION_GROUP_TYPE_REFFields(
	_instance *models.A_RELATION_GROUP_TYPE_REF,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "RELATION_GROUP_TYPE_REF":
			FormDivBasicFieldToField(&(_instance.RELATION_GROUP_TYPE_REF), formDiv)
		}
	}
}

func __gong__New__A_SOURCE_1FormCallback(
	_instance *models.A_SOURCE_1,
	probe *Probe,
	formGroup *form.FormGroup,
) (a_source_1FormCallback *FormCallback[*models.A_SOURCE_1]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveA_SOURCE_1Fields,
	)
}

type A_SOURCE_1FormCallback = FormCallback[*models.A_SOURCE_1]

func saveA_SOURCE_1Fields(
	_instance *models.A_SOURCE_1,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "SPEC_OBJECT_REF":
			FormDivBasicFieldToField(&(_instance.SPEC_OBJECT_REF), formDiv)
		}
	}
}

func __gong__New__A_SOURCE_SPECIFICATION_1FormCallback(
	_instance *models.A_SOURCE_SPECIFICATION_1,
	probe *Probe,
	formGroup *form.FormGroup,
) (a_source_specification_1FormCallback *FormCallback[*models.A_SOURCE_SPECIFICATION_1]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveA_SOURCE_SPECIFICATION_1Fields,
	)
}

type A_SOURCE_SPECIFICATION_1FormCallback = FormCallback[*models.A_SOURCE_SPECIFICATION_1]

func saveA_SOURCE_SPECIFICATION_1Fields(
	_instance *models.A_SOURCE_SPECIFICATION_1,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "SPECIFICATION_REF":
			FormDivEnumStringFieldToField(&(_instance.SPECIFICATION_REF), formDiv)
		}
	}
}

func __gong__New__A_SPECIFICATIONSFormCallback(
	_instance *models.A_SPECIFICATIONS,
	probe *Probe,
	formGroup *form.FormGroup,
) (a_specificationsFormCallback *FormCallback[*models.A_SPECIFICATIONS]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveA_SPECIFICATIONSFields,
	)
}

type A_SPECIFICATIONSFormCallback = FormCallback[*models.A_SPECIFICATIONS]

func saveA_SPECIFICATIONSFields(
	_instance *models.A_SPECIFICATIONS,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "SPECIFICATION":
			FormDivSliceOfPointersToField(_instance, "SPECIFICATION", &(_instance.SPECIFICATION), formDiv, probe)
		}
	}
}

func __gong__New__A_SPECIFICATION_TYPE_REFFormCallback(
	_instance *models.A_SPECIFICATION_TYPE_REF,
	probe *Probe,
	formGroup *form.FormGroup,
) (a_specification_type_refFormCallback *FormCallback[*models.A_SPECIFICATION_TYPE_REF]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveA_SPECIFICATION_TYPE_REFFields,
	)
}

type A_SPECIFICATION_TYPE_REFFormCallback = FormCallback[*models.A_SPECIFICATION_TYPE_REF]

func saveA_SPECIFICATION_TYPE_REFFields(
	_instance *models.A_SPECIFICATION_TYPE_REF,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "SPECIFICATION_TYPE_REF":
			FormDivBasicFieldToField(&(_instance.SPECIFICATION_TYPE_REF), formDiv)
		}
	}
}

func __gong__New__A_SPECIFIED_VALUESFormCallback(
	_instance *models.A_SPECIFIED_VALUES,
	probe *Probe,
	formGroup *form.FormGroup,
) (a_specified_valuesFormCallback *FormCallback[*models.A_SPECIFIED_VALUES]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveA_SPECIFIED_VALUESFields,
	)
}

type A_SPECIFIED_VALUESFormCallback = FormCallback[*models.A_SPECIFIED_VALUES]

func saveA_SPECIFIED_VALUESFields(
	_instance *models.A_SPECIFIED_VALUES,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "ENUM_VALUE":
			FormDivSliceOfPointersToField(_instance, "ENUM_VALUE", &(_instance.ENUM_VALUE), formDiv, probe)
		}
	}
}

func __gong__New__A_SPEC_ATTRIBUTESFormCallback(
	_instance *models.A_SPEC_ATTRIBUTES,
	probe *Probe,
	formGroup *form.FormGroup,
) (a_spec_attributesFormCallback *FormCallback[*models.A_SPEC_ATTRIBUTES]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveA_SPEC_ATTRIBUTESFields,
	)
}

type A_SPEC_ATTRIBUTESFormCallback = FormCallback[*models.A_SPEC_ATTRIBUTES]

func saveA_SPEC_ATTRIBUTESFields(
	_instance *models.A_SPEC_ATTRIBUTES,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "ATTRIBUTE_DEFINITION_BOOLEAN":
			FormDivSliceOfPointersToField(_instance, "ATTRIBUTE_DEFINITION_BOOLEAN", &(_instance.ATTRIBUTE_DEFINITION_BOOLEAN), formDiv, probe)
		case "ATTRIBUTE_DEFINITION_DATE":
			FormDivSliceOfPointersToField(_instance, "ATTRIBUTE_DEFINITION_DATE", &(_instance.ATTRIBUTE_DEFINITION_DATE), formDiv, probe)
		case "ATTRIBUTE_DEFINITION_ENUMERATION":
			FormDivSliceOfPointersToField(_instance, "ATTRIBUTE_DEFINITION_ENUMERATION", &(_instance.ATTRIBUTE_DEFINITION_ENUMERATION), formDiv, probe)
		case "ATTRIBUTE_DEFINITION_INTEGER":
			FormDivSliceOfPointersToField(_instance, "ATTRIBUTE_DEFINITION_INTEGER", &(_instance.ATTRIBUTE_DEFINITION_INTEGER), formDiv, probe)
		case "ATTRIBUTE_DEFINITION_REAL":
			FormDivSliceOfPointersToField(_instance, "ATTRIBUTE_DEFINITION_REAL", &(_instance.ATTRIBUTE_DEFINITION_REAL), formDiv, probe)
		case "ATTRIBUTE_DEFINITION_STRING":
			FormDivSliceOfPointersToField(_instance, "ATTRIBUTE_DEFINITION_STRING", &(_instance.ATTRIBUTE_DEFINITION_STRING), formDiv, probe)
		case "ATTRIBUTE_DEFINITION_XHTML":
			FormDivSliceOfPointersToField(_instance, "ATTRIBUTE_DEFINITION_XHTML", &(_instance.ATTRIBUTE_DEFINITION_XHTML), formDiv, probe)
		}
	}
}

func __gong__New__A_SPEC_OBJECTSFormCallback(
	_instance *models.A_SPEC_OBJECTS,
	probe *Probe,
	formGroup *form.FormGroup,
) (a_spec_objectsFormCallback *FormCallback[*models.A_SPEC_OBJECTS]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveA_SPEC_OBJECTSFields,
	)
}

type A_SPEC_OBJECTSFormCallback = FormCallback[*models.A_SPEC_OBJECTS]

func saveA_SPEC_OBJECTSFields(
	_instance *models.A_SPEC_OBJECTS,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "SPEC_OBJECT":
			FormDivSliceOfPointersToField(_instance, "SPEC_OBJECT", &(_instance.SPEC_OBJECT), formDiv, probe)
		}
	}
}

func __gong__New__A_SPEC_OBJECT_TYPE_REFFormCallback(
	_instance *models.A_SPEC_OBJECT_TYPE_REF,
	probe *Probe,
	formGroup *form.FormGroup,
) (a_spec_object_type_refFormCallback *FormCallback[*models.A_SPEC_OBJECT_TYPE_REF]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveA_SPEC_OBJECT_TYPE_REFFields,
	)
}

type A_SPEC_OBJECT_TYPE_REFFormCallback = FormCallback[*models.A_SPEC_OBJECT_TYPE_REF]

func saveA_SPEC_OBJECT_TYPE_REFFields(
	_instance *models.A_SPEC_OBJECT_TYPE_REF,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "SPEC_OBJECT_TYPE_REF":
			FormDivBasicFieldToField(&(_instance.SPEC_OBJECT_TYPE_REF), formDiv)
		}
	}
}

func __gong__New__A_SPEC_RELATIONSFormCallback(
	_instance *models.A_SPEC_RELATIONS,
	probe *Probe,
	formGroup *form.FormGroup,
) (a_spec_relationsFormCallback *FormCallback[*models.A_SPEC_RELATIONS]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveA_SPEC_RELATIONSFields,
	)
}

type A_SPEC_RELATIONSFormCallback = FormCallback[*models.A_SPEC_RELATIONS]

func saveA_SPEC_RELATIONSFields(
	_instance *models.A_SPEC_RELATIONS,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "SPEC_RELATION":
			FormDivSliceOfPointersToField(_instance, "SPEC_RELATION", &(_instance.SPEC_RELATION), formDiv, probe)
		}
	}
}

func __gong__New__A_SPEC_RELATION_GROUPSFormCallback(
	_instance *models.A_SPEC_RELATION_GROUPS,
	probe *Probe,
	formGroup *form.FormGroup,
) (a_spec_relation_groupsFormCallback *FormCallback[*models.A_SPEC_RELATION_GROUPS]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveA_SPEC_RELATION_GROUPSFields,
	)
}

type A_SPEC_RELATION_GROUPSFormCallback = FormCallback[*models.A_SPEC_RELATION_GROUPS]

func saveA_SPEC_RELATION_GROUPSFields(
	_instance *models.A_SPEC_RELATION_GROUPS,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "RELATION_GROUP":
			FormDivSliceOfPointersToField(_instance, "RELATION_GROUP", &(_instance.RELATION_GROUP), formDiv, probe)
		}
	}
}

func __gong__New__A_SPEC_RELATION_REFFormCallback(
	_instance *models.A_SPEC_RELATION_REF,
	probe *Probe,
	formGroup *form.FormGroup,
) (a_spec_relation_refFormCallback *FormCallback[*models.A_SPEC_RELATION_REF]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveA_SPEC_RELATION_REFFields,
	)
}

type A_SPEC_RELATION_REFFormCallback = FormCallback[*models.A_SPEC_RELATION_REF]

func saveA_SPEC_RELATION_REFFields(
	_instance *models.A_SPEC_RELATION_REF,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "SPEC_RELATION_REF":
			FormDivBasicFieldToField(&(_instance.SPEC_RELATION_REF), formDiv)
		}
	}
}

func __gong__New__A_SPEC_RELATION_TYPE_REFFormCallback(
	_instance *models.A_SPEC_RELATION_TYPE_REF,
	probe *Probe,
	formGroup *form.FormGroup,
) (a_spec_relation_type_refFormCallback *FormCallback[*models.A_SPEC_RELATION_TYPE_REF]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveA_SPEC_RELATION_TYPE_REFFields,
	)
}

type A_SPEC_RELATION_TYPE_REFFormCallback = FormCallback[*models.A_SPEC_RELATION_TYPE_REF]

func saveA_SPEC_RELATION_TYPE_REFFields(
	_instance *models.A_SPEC_RELATION_TYPE_REF,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "SPEC_RELATION_TYPE_REF":
			FormDivBasicFieldToField(&(_instance.SPEC_RELATION_TYPE_REF), formDiv)
		}
	}
}

func __gong__New__A_SPEC_TYPESFormCallback(
	_instance *models.A_SPEC_TYPES,
	probe *Probe,
	formGroup *form.FormGroup,
) (a_spec_typesFormCallback *FormCallback[*models.A_SPEC_TYPES]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveA_SPEC_TYPESFields,
	)
}

type A_SPEC_TYPESFormCallback = FormCallback[*models.A_SPEC_TYPES]

func saveA_SPEC_TYPESFields(
	_instance *models.A_SPEC_TYPES,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "RELATION_GROUP_TYPE":
			FormDivSliceOfPointersToField(_instance, "RELATION_GROUP_TYPE", &(_instance.RELATION_GROUP_TYPE), formDiv, probe)
		case "SPEC_OBJECT_TYPE":
			FormDivSliceOfPointersToField(_instance, "SPEC_OBJECT_TYPE", &(_instance.SPEC_OBJECT_TYPE), formDiv, probe)
		case "SPEC_RELATION_TYPE":
			FormDivSliceOfPointersToField(_instance, "SPEC_RELATION_TYPE", &(_instance.SPEC_RELATION_TYPE), formDiv, probe)
		case "SPECIFICATION_TYPE":
			FormDivSliceOfPointersToField(_instance, "SPECIFICATION_TYPE", &(_instance.SPECIFICATION_TYPE), formDiv, probe)
		}
	}
}

func __gong__New__A_THE_HEADERFormCallback(
	_instance *models.A_THE_HEADER,
	probe *Probe,
	formGroup *form.FormGroup,
) (a_the_headerFormCallback *FormCallback[*models.A_THE_HEADER]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveA_THE_HEADERFields,
	)
}

type A_THE_HEADERFormCallback = FormCallback[*models.A_THE_HEADER]

func saveA_THE_HEADERFields(
	_instance *models.A_THE_HEADER,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "REQ_IF_HEADER":
			FormDivSelectFieldToField(&(_instance.REQ_IF_HEADER), probe.stageOfInterest, formDiv)
		}
	}
}

func __gong__New__A_TOOL_EXTENSIONSFormCallback(
	_instance *models.A_TOOL_EXTENSIONS,
	probe *Probe,
	formGroup *form.FormGroup,
) (a_tool_extensionsFormCallback *FormCallback[*models.A_TOOL_EXTENSIONS]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveA_TOOL_EXTENSIONSFields,
	)
}

type A_TOOL_EXTENSIONSFormCallback = FormCallback[*models.A_TOOL_EXTENSIONS]

func saveA_TOOL_EXTENSIONSFields(
	_instance *models.A_TOOL_EXTENSIONS,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "REQ_IF_TOOL_EXTENSION":
			FormDivSliceOfPointersToField(_instance, "REQ_IF_TOOL_EXTENSION", &(_instance.REQ_IF_TOOL_EXTENSION), formDiv, probe)
		}
	}
}

func __gong__New__DATATYPE_DEFINITION_BOOLEANFormCallback(
	_instance *models.DATATYPE_DEFINITION_BOOLEAN,
	probe *Probe,
	formGroup *form.FormGroup,
) (datatype_definition_booleanFormCallback *FormCallback[*models.DATATYPE_DEFINITION_BOOLEAN]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveDATATYPE_DEFINITION_BOOLEANFields,
	)
}

type DATATYPE_DEFINITION_BOOLEANFormCallback = FormCallback[*models.DATATYPE_DEFINITION_BOOLEAN]

func saveDATATYPE_DEFINITION_BOOLEANFields(
	_instance *models.DATATYPE_DEFINITION_BOOLEAN,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "DESC":
			FormDivBasicFieldToField(&(_instance.DESC), formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&(_instance.IDENTIFIER), formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&(_instance.LAST_CHANGE), formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&(_instance.LONG_NAME), formDiv)
		case "ALTERNATIVE_ID":
			FormDivSelectFieldToField(&(_instance.ALTERNATIVE_ID), probe.stageOfInterest, formDiv)
		case "A_DATATYPES:DATATYPE_DEFINITION_BOOLEAN":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "DATATYPE_DEFINITION_BOOLEAN", func(owner *models.A_DATATYPES) *[]*models.DATATYPE_DEFINITION_BOOLEAN { return &owner.DATATYPE_DEFINITION_BOOLEAN })
		}
	}
}

func __gong__New__DATATYPE_DEFINITION_DATEFormCallback(
	_instance *models.DATATYPE_DEFINITION_DATE,
	probe *Probe,
	formGroup *form.FormGroup,
) (datatype_definition_dateFormCallback *FormCallback[*models.DATATYPE_DEFINITION_DATE]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveDATATYPE_DEFINITION_DATEFields,
	)
}

type DATATYPE_DEFINITION_DATEFormCallback = FormCallback[*models.DATATYPE_DEFINITION_DATE]

func saveDATATYPE_DEFINITION_DATEFields(
	_instance *models.DATATYPE_DEFINITION_DATE,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "DESC":
			FormDivBasicFieldToField(&(_instance.DESC), formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&(_instance.IDENTIFIER), formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&(_instance.LAST_CHANGE), formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&(_instance.LONG_NAME), formDiv)
		case "ALTERNATIVE_ID":
			FormDivSelectFieldToField(&(_instance.ALTERNATIVE_ID), probe.stageOfInterest, formDiv)
		case "A_DATATYPES:DATATYPE_DEFINITION_DATE":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "DATATYPE_DEFINITION_DATE", func(owner *models.A_DATATYPES) *[]*models.DATATYPE_DEFINITION_DATE { return &owner.DATATYPE_DEFINITION_DATE })
		}
	}
}

func __gong__New__DATATYPE_DEFINITION_ENUMERATIONFormCallback(
	_instance *models.DATATYPE_DEFINITION_ENUMERATION,
	probe *Probe,
	formGroup *form.FormGroup,
) (datatype_definition_enumerationFormCallback *FormCallback[*models.DATATYPE_DEFINITION_ENUMERATION]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveDATATYPE_DEFINITION_ENUMERATIONFields,
	)
}

type DATATYPE_DEFINITION_ENUMERATIONFormCallback = FormCallback[*models.DATATYPE_DEFINITION_ENUMERATION]

func saveDATATYPE_DEFINITION_ENUMERATIONFields(
	_instance *models.DATATYPE_DEFINITION_ENUMERATION,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "DESC":
			FormDivBasicFieldToField(&(_instance.DESC), formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&(_instance.IDENTIFIER), formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&(_instance.LAST_CHANGE), formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&(_instance.LONG_NAME), formDiv)
		case "ALTERNATIVE_ID":
			FormDivSelectFieldToField(&(_instance.ALTERNATIVE_ID), probe.stageOfInterest, formDiv)
		case "SPECIFIED_VALUES":
			FormDivSelectFieldToField(&(_instance.SPECIFIED_VALUES), probe.stageOfInterest, formDiv)
		case "A_DATATYPES:DATATYPE_DEFINITION_ENUMERATION":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "DATATYPE_DEFINITION_ENUMERATION", func(owner *models.A_DATATYPES) *[]*models.DATATYPE_DEFINITION_ENUMERATION { return &owner.DATATYPE_DEFINITION_ENUMERATION })
		}
	}
}

func __gong__New__DATATYPE_DEFINITION_INTEGERFormCallback(
	_instance *models.DATATYPE_DEFINITION_INTEGER,
	probe *Probe,
	formGroup *form.FormGroup,
) (datatype_definition_integerFormCallback *FormCallback[*models.DATATYPE_DEFINITION_INTEGER]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveDATATYPE_DEFINITION_INTEGERFields,
	)
}

type DATATYPE_DEFINITION_INTEGERFormCallback = FormCallback[*models.DATATYPE_DEFINITION_INTEGER]

func saveDATATYPE_DEFINITION_INTEGERFields(
	_instance *models.DATATYPE_DEFINITION_INTEGER,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "DESC":
			FormDivBasicFieldToField(&(_instance.DESC), formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&(_instance.IDENTIFIER), formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&(_instance.LAST_CHANGE), formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&(_instance.LONG_NAME), formDiv)
		case "MAX":
			FormDivBasicFieldToField(&(_instance.MAX), formDiv)
		case "MIN":
			FormDivBasicFieldToField(&(_instance.MIN), formDiv)
		case "ALTERNATIVE_ID":
			FormDivSelectFieldToField(&(_instance.ALTERNATIVE_ID), probe.stageOfInterest, formDiv)
		case "A_DATATYPES:DATATYPE_DEFINITION_INTEGER":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "DATATYPE_DEFINITION_INTEGER", func(owner *models.A_DATATYPES) *[]*models.DATATYPE_DEFINITION_INTEGER { return &owner.DATATYPE_DEFINITION_INTEGER })
		}
	}
}

func __gong__New__DATATYPE_DEFINITION_REALFormCallback(
	_instance *models.DATATYPE_DEFINITION_REAL,
	probe *Probe,
	formGroup *form.FormGroup,
) (datatype_definition_realFormCallback *FormCallback[*models.DATATYPE_DEFINITION_REAL]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveDATATYPE_DEFINITION_REALFields,
	)
}

type DATATYPE_DEFINITION_REALFormCallback = FormCallback[*models.DATATYPE_DEFINITION_REAL]

func saveDATATYPE_DEFINITION_REALFields(
	_instance *models.DATATYPE_DEFINITION_REAL,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "ACCURACY":
			FormDivBasicFieldToField(&(_instance.ACCURACY), formDiv)
		case "DESC":
			FormDivBasicFieldToField(&(_instance.DESC), formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&(_instance.IDENTIFIER), formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&(_instance.LAST_CHANGE), formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&(_instance.LONG_NAME), formDiv)
		case "MAX":
			FormDivBasicFieldToField(&(_instance.MAX), formDiv)
		case "MIN":
			FormDivBasicFieldToField(&(_instance.MIN), formDiv)
		case "ALTERNATIVE_ID":
			FormDivSelectFieldToField(&(_instance.ALTERNATIVE_ID), probe.stageOfInterest, formDiv)
		case "A_DATATYPES:DATATYPE_DEFINITION_REAL":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "DATATYPE_DEFINITION_REAL", func(owner *models.A_DATATYPES) *[]*models.DATATYPE_DEFINITION_REAL { return &owner.DATATYPE_DEFINITION_REAL })
		}
	}
}

func __gong__New__DATATYPE_DEFINITION_STRINGFormCallback(
	_instance *models.DATATYPE_DEFINITION_STRING,
	probe *Probe,
	formGroup *form.FormGroup,
) (datatype_definition_stringFormCallback *FormCallback[*models.DATATYPE_DEFINITION_STRING]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveDATATYPE_DEFINITION_STRINGFields,
	)
}

type DATATYPE_DEFINITION_STRINGFormCallback = FormCallback[*models.DATATYPE_DEFINITION_STRING]

func saveDATATYPE_DEFINITION_STRINGFields(
	_instance *models.DATATYPE_DEFINITION_STRING,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "DESC":
			FormDivBasicFieldToField(&(_instance.DESC), formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&(_instance.IDENTIFIER), formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&(_instance.LAST_CHANGE), formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&(_instance.LONG_NAME), formDiv)
		case "MAX_LENGTH":
			FormDivBasicFieldToField(&(_instance.MAX_LENGTH), formDiv)
		case "ALTERNATIVE_ID":
			FormDivSelectFieldToField(&(_instance.ALTERNATIVE_ID), probe.stageOfInterest, formDiv)
		case "A_DATATYPES:DATATYPE_DEFINITION_STRING":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "DATATYPE_DEFINITION_STRING", func(owner *models.A_DATATYPES) *[]*models.DATATYPE_DEFINITION_STRING { return &owner.DATATYPE_DEFINITION_STRING })
		}
	}
}

func __gong__New__DATATYPE_DEFINITION_XHTMLFormCallback(
	_instance *models.DATATYPE_DEFINITION_XHTML,
	probe *Probe,
	formGroup *form.FormGroup,
) (datatype_definition_xhtmlFormCallback *FormCallback[*models.DATATYPE_DEFINITION_XHTML]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveDATATYPE_DEFINITION_XHTMLFields,
	)
}

type DATATYPE_DEFINITION_XHTMLFormCallback = FormCallback[*models.DATATYPE_DEFINITION_XHTML]

func saveDATATYPE_DEFINITION_XHTMLFields(
	_instance *models.DATATYPE_DEFINITION_XHTML,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "DESC":
			FormDivBasicFieldToField(&(_instance.DESC), formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&(_instance.IDENTIFIER), formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&(_instance.LAST_CHANGE), formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&(_instance.LONG_NAME), formDiv)
		case "ALTERNATIVE_ID":
			FormDivSelectFieldToField(&(_instance.ALTERNATIVE_ID), probe.stageOfInterest, formDiv)
		case "A_DATATYPES:DATATYPE_DEFINITION_XHTML":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "DATATYPE_DEFINITION_XHTML", func(owner *models.A_DATATYPES) *[]*models.DATATYPE_DEFINITION_XHTML { return &owner.DATATYPE_DEFINITION_XHTML })
		}
	}
}

func __gong__New__EMBEDDED_VALUEFormCallback(
	_instance *models.EMBEDDED_VALUE,
	probe *Probe,
	formGroup *form.FormGroup,
) (embedded_valueFormCallback *FormCallback[*models.EMBEDDED_VALUE]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveEMBEDDED_VALUEFields,
	)
}

type EMBEDDED_VALUEFormCallback = FormCallback[*models.EMBEDDED_VALUE]

func saveEMBEDDED_VALUEFields(
	_instance *models.EMBEDDED_VALUE,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "KEY":
			FormDivBasicFieldToField(&(_instance.KEY), formDiv)
		case "OTHER_CONTENT":
			FormDivBasicFieldToField(&(_instance.OTHER_CONTENT), formDiv)
		}
	}
}

func __gong__New__ENUM_VALUEFormCallback(
	_instance *models.ENUM_VALUE,
	probe *Probe,
	formGroup *form.FormGroup,
) (enum_valueFormCallback *FormCallback[*models.ENUM_VALUE]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveENUM_VALUEFields,
	)
}

type ENUM_VALUEFormCallback = FormCallback[*models.ENUM_VALUE]

func saveENUM_VALUEFields(
	_instance *models.ENUM_VALUE,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "DESC":
			FormDivBasicFieldToField(&(_instance.DESC), formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&(_instance.IDENTIFIER), formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&(_instance.LAST_CHANGE), formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&(_instance.LONG_NAME), formDiv)
		case "ALTERNATIVE_ID":
			FormDivSelectFieldToField(&(_instance.ALTERNATIVE_ID), probe.stageOfInterest, formDiv)
		case "PROPERTIES":
			FormDivSelectFieldToField(&(_instance.PROPERTIES), probe.stageOfInterest, formDiv)
		case "A_SPECIFIED_VALUES:ENUM_VALUE":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "ENUM_VALUE", func(owner *models.A_SPECIFIED_VALUES) *[]*models.ENUM_VALUE { return &owner.ENUM_VALUE })
		}
	}
}

func __gong__New__EmbeddedJpgImageFormCallback(
	_instance *models.EmbeddedJpgImage,
	probe *Probe,
	formGroup *form.FormGroup,
) (embeddedjpgimageFormCallback *FormCallback[*models.EmbeddedJpgImage]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveEmbeddedJpgImageFields,
	)
}

type EmbeddedJpgImageFormCallback = FormCallback[*models.EmbeddedJpgImage]

func saveEmbeddedJpgImageFields(
	_instance *models.EmbeddedJpgImage,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Base64Content":
			FormDivBasicFieldToField(&(_instance.Base64Content), formDiv)
		}
	}
}

func __gong__New__EmbeddedPngImageFormCallback(
	_instance *models.EmbeddedPngImage,
	probe *Probe,
	formGroup *form.FormGroup,
) (embeddedpngimageFormCallback *FormCallback[*models.EmbeddedPngImage]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveEmbeddedPngImageFields,
	)
}

type EmbeddedPngImageFormCallback = FormCallback[*models.EmbeddedPngImage]

func saveEmbeddedPngImageFields(
	_instance *models.EmbeddedPngImage,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Base64Content":
			FormDivBasicFieldToField(&(_instance.Base64Content), formDiv)
		}
	}
}

func __gong__New__EmbeddedSvgImageFormCallback(
	_instance *models.EmbeddedSvgImage,
	probe *Probe,
	formGroup *form.FormGroup,
) (embeddedsvgimageFormCallback *FormCallback[*models.EmbeddedSvgImage]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveEmbeddedSvgImageFields,
	)
}

type EmbeddedSvgImageFormCallback = FormCallback[*models.EmbeddedSvgImage]

func saveEmbeddedSvgImageFields(
	_instance *models.EmbeddedSvgImage,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Content":
			FormDivBasicFieldToField(&(_instance.Content), formDiv)
		}
	}
}

func __gong__New__KillFormCallback(
	_instance *models.Kill,
	probe *Probe,
	formGroup *form.FormGroup,
) (killFormCallback *FormCallback[*models.Kill]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveKillFields,
	)
}

type KillFormCallback = FormCallback[*models.Kill]

func saveKillFields(
	_instance *models.Kill,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		}
	}
}

func __gong__New__Map_identifier_boolFormCallback(
	_instance *models.Map_identifier_bool,
	probe *Probe,
	formGroup *form.FormGroup,
) (map_identifier_boolFormCallback *FormCallback[*models.Map_identifier_bool]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveMap_identifier_boolFields,
	)
}

type Map_identifier_boolFormCallback = FormCallback[*models.Map_identifier_bool]

func saveMap_identifier_boolFields(
	_instance *models.Map_identifier_bool,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Value":
			FormDivBasicFieldToField(&(_instance.Value), formDiv)
		}
	}
}

func __gong__New__RELATION_GROUPFormCallback(
	_instance *models.RELATION_GROUP,
	probe *Probe,
	formGroup *form.FormGroup,
) (relation_groupFormCallback *FormCallback[*models.RELATION_GROUP]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveRELATION_GROUPFields,
	)
}

type RELATION_GROUPFormCallback = FormCallback[*models.RELATION_GROUP]

func saveRELATION_GROUPFields(
	_instance *models.RELATION_GROUP,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "DESC":
			FormDivBasicFieldToField(&(_instance.DESC), formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&(_instance.IDENTIFIER), formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&(_instance.LAST_CHANGE), formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&(_instance.LONG_NAME), formDiv)
		case "ALTERNATIVE_ID":
			FormDivSelectFieldToField(&(_instance.ALTERNATIVE_ID), probe.stageOfInterest, formDiv)
		case "SOURCE_SPECIFICATION":
			FormDivSelectFieldToField(&(_instance.SOURCE_SPECIFICATION), probe.stageOfInterest, formDiv)
		case "SPEC_RELATIONS":
			FormDivSelectFieldToField(&(_instance.SPEC_RELATIONS), probe.stageOfInterest, formDiv)
		case "TARGET_SPECIFICATION":
			FormDivSelectFieldToField(&(_instance.TARGET_SPECIFICATION), probe.stageOfInterest, formDiv)
		case "TYPE":
			FormDivSelectFieldToField(&(_instance.TYPE), probe.stageOfInterest, formDiv)
		case "A_SPEC_RELATION_GROUPS:RELATION_GROUP":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "RELATION_GROUP", func(owner *models.A_SPEC_RELATION_GROUPS) *[]*models.RELATION_GROUP { return &owner.RELATION_GROUP })
		}
	}
}

func __gong__New__RELATION_GROUP_TYPEFormCallback(
	_instance *models.RELATION_GROUP_TYPE,
	probe *Probe,
	formGroup *form.FormGroup,
) (relation_group_typeFormCallback *FormCallback[*models.RELATION_GROUP_TYPE]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveRELATION_GROUP_TYPEFields,
	)
}

type RELATION_GROUP_TYPEFormCallback = FormCallback[*models.RELATION_GROUP_TYPE]

func saveRELATION_GROUP_TYPEFields(
	_instance *models.RELATION_GROUP_TYPE,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "DESC":
			FormDivBasicFieldToField(&(_instance.DESC), formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&(_instance.IDENTIFIER), formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&(_instance.LAST_CHANGE), formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&(_instance.LONG_NAME), formDiv)
		case "ALTERNATIVE_ID":
			FormDivSelectFieldToField(&(_instance.ALTERNATIVE_ID), probe.stageOfInterest, formDiv)
		case "SPEC_ATTRIBUTES":
			FormDivSelectFieldToField(&(_instance.SPEC_ATTRIBUTES), probe.stageOfInterest, formDiv)
		case "A_SPEC_TYPES:RELATION_GROUP_TYPE":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "RELATION_GROUP_TYPE", func(owner *models.A_SPEC_TYPES) *[]*models.RELATION_GROUP_TYPE { return &owner.RELATION_GROUP_TYPE })
		}
	}
}

func __gong__New__REQ_IFFormCallback(
	_instance *models.REQ_IF,
	probe *Probe,
	formGroup *form.FormGroup,
) (req_ifFormCallback *FormCallback[*models.REQ_IF]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveREQ_IFFields,
	)
}

type REQ_IFFormCallback = FormCallback[*models.REQ_IF]

func saveREQ_IFFields(
	_instance *models.REQ_IF,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Lang":
			FormDivBasicFieldToField(&(_instance.Lang), formDiv)
		case "THE_HEADER":
			FormDivSelectFieldToField(&(_instance.THE_HEADER), probe.stageOfInterest, formDiv)
		case "CORE_CONTENT":
			FormDivSelectFieldToField(&(_instance.CORE_CONTENT), probe.stageOfInterest, formDiv)
		case "TOOL_EXTENSIONS":
			FormDivSelectFieldToField(&(_instance.TOOL_EXTENSIONS), probe.stageOfInterest, formDiv)
		}
	}
}

func __gong__New__REQ_IF_CONTENTFormCallback(
	_instance *models.REQ_IF_CONTENT,
	probe *Probe,
	formGroup *form.FormGroup,
) (req_if_contentFormCallback *FormCallback[*models.REQ_IF_CONTENT]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveREQ_IF_CONTENTFields,
	)
}

type REQ_IF_CONTENTFormCallback = FormCallback[*models.REQ_IF_CONTENT]

func saveREQ_IF_CONTENTFields(
	_instance *models.REQ_IF_CONTENT,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "DATATYPES":
			FormDivSelectFieldToField(&(_instance.DATATYPES), probe.stageOfInterest, formDiv)
		case "SPEC_TYPES":
			FormDivSelectFieldToField(&(_instance.SPEC_TYPES), probe.stageOfInterest, formDiv)
		case "SPEC_OBJECTS":
			FormDivSelectFieldToField(&(_instance.SPEC_OBJECTS), probe.stageOfInterest, formDiv)
		case "SPEC_RELATIONS":
			FormDivSelectFieldToField(&(_instance.SPEC_RELATIONS), probe.stageOfInterest, formDiv)
		case "SPECIFICATIONS":
			FormDivSelectFieldToField(&(_instance.SPECIFICATIONS), probe.stageOfInterest, formDiv)
		case "SPEC_RELATION_GROUPS":
			FormDivSelectFieldToField(&(_instance.SPEC_RELATION_GROUPS), probe.stageOfInterest, formDiv)
		}
	}
}

func __gong__New__REQ_IF_HEADERFormCallback(
	_instance *models.REQ_IF_HEADER,
	probe *Probe,
	formGroup *form.FormGroup,
) (req_if_headerFormCallback *FormCallback[*models.REQ_IF_HEADER]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveREQ_IF_HEADERFields,
	)
}

type REQ_IF_HEADERFormCallback = FormCallback[*models.REQ_IF_HEADER]

func saveREQ_IF_HEADERFields(
	_instance *models.REQ_IF_HEADER,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&(_instance.IDENTIFIER), formDiv)
		case "COMMENT":
			FormDivBasicFieldToField(&(_instance.COMMENT), formDiv)
		case "CREATION_TIME":
			FormDivBasicFieldToField(&(_instance.CREATION_TIME), formDiv)
		case "REPOSITORY_ID":
			FormDivBasicFieldToField(&(_instance.REPOSITORY_ID), formDiv)
		case "REQ_IF_TOOL_ID":
			FormDivBasicFieldToField(&(_instance.REQ_IF_TOOL_ID), formDiv)
		case "REQ_IF_VERSION":
			FormDivBasicFieldToField(&(_instance.REQ_IF_VERSION), formDiv)
		case "SOURCE_TOOL_ID":
			FormDivBasicFieldToField(&(_instance.SOURCE_TOOL_ID), formDiv)
		case "TITLE":
			FormDivBasicFieldToField(&(_instance.TITLE), formDiv)
		}
	}
}

func __gong__New__REQ_IF_TOOL_EXTENSIONFormCallback(
	_instance *models.REQ_IF_TOOL_EXTENSION,
	probe *Probe,
	formGroup *form.FormGroup,
) (req_if_tool_extensionFormCallback *FormCallback[*models.REQ_IF_TOOL_EXTENSION]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveREQ_IF_TOOL_EXTENSIONFields,
	)
}

type REQ_IF_TOOL_EXTENSIONFormCallback = FormCallback[*models.REQ_IF_TOOL_EXTENSION]

func saveREQ_IF_TOOL_EXTENSIONFields(
	_instance *models.REQ_IF_TOOL_EXTENSION,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "A_TOOL_EXTENSIONS:REQ_IF_TOOL_EXTENSION":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "REQ_IF_TOOL_EXTENSION", func(owner *models.A_TOOL_EXTENSIONS) *[]*models.REQ_IF_TOOL_EXTENSION { return &owner.REQ_IF_TOOL_EXTENSION })
		}
	}
}

func __gong__New__SPECIFICATIONFormCallback(
	_instance *models.SPECIFICATION,
	probe *Probe,
	formGroup *form.FormGroup,
) (specificationFormCallback *FormCallback[*models.SPECIFICATION]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveSPECIFICATIONFields,
	)
}

type SPECIFICATIONFormCallback = FormCallback[*models.SPECIFICATION]

func saveSPECIFICATIONFields(
	_instance *models.SPECIFICATION,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "DESC":
			FormDivBasicFieldToField(&(_instance.DESC), formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&(_instance.IDENTIFIER), formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&(_instance.LAST_CHANGE), formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&(_instance.LONG_NAME), formDiv)
		case "ALTERNATIVE_ID":
			FormDivSelectFieldToField(&(_instance.ALTERNATIVE_ID), probe.stageOfInterest, formDiv)
		case "TYPE":
			FormDivSelectFieldToField(&(_instance.TYPE), probe.stageOfInterest, formDiv)
		case "CHILDREN":
			FormDivSelectFieldToField(&(_instance.CHILDREN), probe.stageOfInterest, formDiv)
		case "VALUES":
			FormDivSelectFieldToField(&(_instance.VALUES), probe.stageOfInterest, formDiv)
		case "A_SPECIFICATIONS:SPECIFICATION":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "SPECIFICATION", func(owner *models.A_SPECIFICATIONS) *[]*models.SPECIFICATION { return &owner.SPECIFICATION })
		}
	}
}

func __gong__New__SPECIFICATION_RenderingFormCallback(
	_instance *models.SPECIFICATION_Rendering,
	probe *Probe,
	formGroup *form.FormGroup,
) (specification_renderingFormCallback *FormCallback[*models.SPECIFICATION_Rendering]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveSPECIFICATION_RenderingFields,
	)
}

type SPECIFICATION_RenderingFormCallback = FormCallback[*models.SPECIFICATION_Rendering]

func saveSPECIFICATION_RenderingFields(
	_instance *models.SPECIFICATION_Rendering,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "IsNodeExpanded":
			FormDivBasicFieldToField(&(_instance.IsNodeExpanded), formDiv)
		case "IsSelected":
			FormDivBasicFieldToField(&(_instance.IsSelected), formDiv)
		case "IsWithHeadingNumbering":
			FormDivBasicFieldToField(&(_instance.IsWithHeadingNumbering), formDiv)
		}
	}
}

func __gong__New__SPECIFICATION_TYPEFormCallback(
	_instance *models.SPECIFICATION_TYPE,
	probe *Probe,
	formGroup *form.FormGroup,
) (specification_typeFormCallback *FormCallback[*models.SPECIFICATION_TYPE]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveSPECIFICATION_TYPEFields,
	)
}

type SPECIFICATION_TYPEFormCallback = FormCallback[*models.SPECIFICATION_TYPE]

func saveSPECIFICATION_TYPEFields(
	_instance *models.SPECIFICATION_TYPE,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "DESC":
			FormDivBasicFieldToField(&(_instance.DESC), formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&(_instance.IDENTIFIER), formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&(_instance.LAST_CHANGE), formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&(_instance.LONG_NAME), formDiv)
		case "ALTERNATIVE_ID":
			FormDivSelectFieldToField(&(_instance.ALTERNATIVE_ID), probe.stageOfInterest, formDiv)
		case "SPEC_ATTRIBUTES":
			FormDivSelectFieldToField(&(_instance.SPEC_ATTRIBUTES), probe.stageOfInterest, formDiv)
		case "A_SPEC_TYPES:SPECIFICATION_TYPE":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "SPECIFICATION_TYPE", func(owner *models.A_SPEC_TYPES) *[]*models.SPECIFICATION_TYPE { return &owner.SPECIFICATION_TYPE })
		}
	}
}

func __gong__New__SPEC_HIERARCHYFormCallback(
	_instance *models.SPEC_HIERARCHY,
	probe *Probe,
	formGroup *form.FormGroup,
) (spec_hierarchyFormCallback *FormCallback[*models.SPEC_HIERARCHY]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveSPEC_HIERARCHYFields,
	)
}

type SPEC_HIERARCHYFormCallback = FormCallback[*models.SPEC_HIERARCHY]

func saveSPEC_HIERARCHYFields(
	_instance *models.SPEC_HIERARCHY,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "DESC":
			FormDivBasicFieldToField(&(_instance.DESC), formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&(_instance.IDENTIFIER), formDiv)
		case "IS_EDITABLE":
			FormDivBasicFieldToField(&(_instance.IS_EDITABLE), formDiv)
		case "IS_TABLE_INTERNAL":
			FormDivBasicFieldToField(&(_instance.IS_TABLE_INTERNAL), formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&(_instance.LAST_CHANGE), formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&(_instance.LONG_NAME), formDiv)
		case "ALTERNATIVE_ID":
			FormDivSelectFieldToField(&(_instance.ALTERNATIVE_ID), probe.stageOfInterest, formDiv)
		case "OBJECT":
			FormDivSelectFieldToField(&(_instance.OBJECT), probe.stageOfInterest, formDiv)
		case "CHILDREN":
			FormDivSelectFieldToField(&(_instance.CHILDREN), probe.stageOfInterest, formDiv)
		case "EDITABLE_ATTS":
			FormDivSelectFieldToField(&(_instance.EDITABLE_ATTS), probe.stageOfInterest, formDiv)
		case "A_CHILDREN:SPEC_HIERARCHY":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "SPEC_HIERARCHY", func(owner *models.A_CHILDREN) *[]*models.SPEC_HIERARCHY { return &owner.SPEC_HIERARCHY })
		}
	}
}

func __gong__New__SPEC_OBJECTFormCallback(
	_instance *models.SPEC_OBJECT,
	probe *Probe,
	formGroup *form.FormGroup,
) (spec_objectFormCallback *FormCallback[*models.SPEC_OBJECT]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveSPEC_OBJECTFields,
	)
}

type SPEC_OBJECTFormCallback = FormCallback[*models.SPEC_OBJECT]

func saveSPEC_OBJECTFields(
	_instance *models.SPEC_OBJECT,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "DESC":
			FormDivBasicFieldToField(&(_instance.DESC), formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&(_instance.IDENTIFIER), formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&(_instance.LAST_CHANGE), formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&(_instance.LONG_NAME), formDiv)
		case "ALTERNATIVE_ID":
			FormDivSelectFieldToField(&(_instance.ALTERNATIVE_ID), probe.stageOfInterest, formDiv)
		case "VALUES":
			FormDivSelectFieldToField(&(_instance.VALUES), probe.stageOfInterest, formDiv)
		case "TYPE":
			FormDivSelectFieldToField(&(_instance.TYPE), probe.stageOfInterest, formDiv)
		case "A_SPEC_OBJECTS:SPEC_OBJECT":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "SPEC_OBJECT", func(owner *models.A_SPEC_OBJECTS) *[]*models.SPEC_OBJECT { return &owner.SPEC_OBJECT })
		}
	}
}

func __gong__New__SPEC_OBJECT_TYPEFormCallback(
	_instance *models.SPEC_OBJECT_TYPE,
	probe *Probe,
	formGroup *form.FormGroup,
) (spec_object_typeFormCallback *FormCallback[*models.SPEC_OBJECT_TYPE]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveSPEC_OBJECT_TYPEFields,
	)
}

type SPEC_OBJECT_TYPEFormCallback = FormCallback[*models.SPEC_OBJECT_TYPE]

func saveSPEC_OBJECT_TYPEFields(
	_instance *models.SPEC_OBJECT_TYPE,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "DESC":
			FormDivBasicFieldToField(&(_instance.DESC), formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&(_instance.IDENTIFIER), formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&(_instance.LAST_CHANGE), formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&(_instance.LONG_NAME), formDiv)
		case "ALTERNATIVE_ID":
			FormDivSelectFieldToField(&(_instance.ALTERNATIVE_ID), probe.stageOfInterest, formDiv)
		case "SPEC_ATTRIBUTES":
			FormDivSelectFieldToField(&(_instance.SPEC_ATTRIBUTES), probe.stageOfInterest, formDiv)
		case "A_SPEC_TYPES:SPEC_OBJECT_TYPE":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "SPEC_OBJECT_TYPE", func(owner *models.A_SPEC_TYPES) *[]*models.SPEC_OBJECT_TYPE { return &owner.SPEC_OBJECT_TYPE })
		}
	}
}

func __gong__New__SPEC_OBJECT_TYPE_RenderingFormCallback(
	_instance *models.SPEC_OBJECT_TYPE_Rendering,
	probe *Probe,
	formGroup *form.FormGroup,
) (spec_object_type_renderingFormCallback *FormCallback[*models.SPEC_OBJECT_TYPE_Rendering]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveSPEC_OBJECT_TYPE_RenderingFields,
	)
}

type SPEC_OBJECT_TYPE_RenderingFormCallback = FormCallback[*models.SPEC_OBJECT_TYPE_Rendering]

func saveSPEC_OBJECT_TYPE_RenderingFields(
	_instance *models.SPEC_OBJECT_TYPE_Rendering,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "IsNodeExpanded":
			FormDivBasicFieldToField(&(_instance.IsNodeExpanded), formDiv)
		case "ShowIdentifier":
			FormDivBasicFieldToField(&(_instance.ShowIdentifier), formDiv)
		case "ShowName":
			FormDivBasicFieldToField(&(_instance.ShowName), formDiv)
		case "ShowRelations":
			FormDivBasicFieldToField(&(_instance.ShowRelations), formDiv)
		case "IsHeading":
			FormDivBasicFieldToField(&(_instance.IsHeading), formDiv)
		}
	}
}

func __gong__New__SPEC_RELATIONFormCallback(
	_instance *models.SPEC_RELATION,
	probe *Probe,
	formGroup *form.FormGroup,
) (spec_relationFormCallback *FormCallback[*models.SPEC_RELATION]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveSPEC_RELATIONFields,
	)
}

type SPEC_RELATIONFormCallback = FormCallback[*models.SPEC_RELATION]

func saveSPEC_RELATIONFields(
	_instance *models.SPEC_RELATION,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "DESC":
			FormDivBasicFieldToField(&(_instance.DESC), formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&(_instance.IDENTIFIER), formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&(_instance.LAST_CHANGE), formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&(_instance.LONG_NAME), formDiv)
		case "ALTERNATIVE_ID":
			FormDivSelectFieldToField(&(_instance.ALTERNATIVE_ID), probe.stageOfInterest, formDiv)
		case "VALUES":
			FormDivSelectFieldToField(&(_instance.VALUES), probe.stageOfInterest, formDiv)
		case "SOURCE":
			FormDivSelectFieldToField(&(_instance.SOURCE), probe.stageOfInterest, formDiv)
		case "TARGET":
			FormDivSelectFieldToField(&(_instance.TARGET), probe.stageOfInterest, formDiv)
		case "TYPE":
			FormDivSelectFieldToField(&(_instance.TYPE), probe.stageOfInterest, formDiv)
		case "A_SPEC_RELATIONS:SPEC_RELATION":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "SPEC_RELATION", func(owner *models.A_SPEC_RELATIONS) *[]*models.SPEC_RELATION { return &owner.SPEC_RELATION })
		}
	}
}

func __gong__New__SPEC_RELATION_TYPEFormCallback(
	_instance *models.SPEC_RELATION_TYPE,
	probe *Probe,
	formGroup *form.FormGroup,
) (spec_relation_typeFormCallback *FormCallback[*models.SPEC_RELATION_TYPE]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveSPEC_RELATION_TYPEFields,
	)
}

type SPEC_RELATION_TYPEFormCallback = FormCallback[*models.SPEC_RELATION_TYPE]

func saveSPEC_RELATION_TYPEFields(
	_instance *models.SPEC_RELATION_TYPE,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "DESC":
			FormDivBasicFieldToField(&(_instance.DESC), formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&(_instance.IDENTIFIER), formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&(_instance.LAST_CHANGE), formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&(_instance.LONG_NAME), formDiv)
		case "ALTERNATIVE_ID":
			FormDivSelectFieldToField(&(_instance.ALTERNATIVE_ID), probe.stageOfInterest, formDiv)
		case "SPEC_ATTRIBUTES":
			FormDivSelectFieldToField(&(_instance.SPEC_ATTRIBUTES), probe.stageOfInterest, formDiv)
		case "A_SPEC_TYPES:SPEC_RELATION_TYPE":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "SPEC_RELATION_TYPE", func(owner *models.A_SPEC_TYPES) *[]*models.SPEC_RELATION_TYPE { return &owner.SPEC_RELATION_TYPE })
		}
	}
}

func __gong__New__StaticWebSiteFormCallback(
	_instance *models.StaticWebSite,
	probe *Probe,
	formGroup *form.FormGroup,
) (staticwebsiteFormCallback *FormCallback[*models.StaticWebSite]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveStaticWebSiteFields,
	)
}

type StaticWebSiteFormCallback = FormCallback[*models.StaticWebSite]

func saveStaticWebSiteFields(
	_instance *models.StaticWebSite,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "MarkdownContent":
			FormDivBasicFieldToField(&(_instance.MarkdownContent), formDiv)
		case "Chapters":
			FormDivSliceOfPointersToField(_instance, "Chapters", &(_instance.Chapters), formDiv, probe)
		case "InputImagesDir":
			FormDivBasicFieldToField(&(_instance.InputImagesDir), formDiv)
		case "OutputStaticWebDir":
			FormDivBasicFieldToField(&(_instance.OutputStaticWebDir), formDiv)
		case "VersionInfo":
			FormDivBasicFieldToField(&(_instance.VersionInfo), formDiv)
		}
	}
}

func __gong__New__StaticWebSiteChapterFormCallback(
	_instance *models.StaticWebSiteChapter,
	probe *Probe,
	formGroup *form.FormGroup,
) (staticwebsitechapterFormCallback *FormCallback[*models.StaticWebSiteChapter]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveStaticWebSiteChapterFields,
	)
}

type StaticWebSiteChapterFormCallback = FormCallback[*models.StaticWebSiteChapter]

func saveStaticWebSiteChapterFields(
	_instance *models.StaticWebSiteChapter,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "MarkdownContent":
			FormDivBasicFieldToField(&(_instance.MarkdownContent), formDiv)
		case "Paragraphs":
			FormDivSliceOfPointersToField(_instance, "Paragraphs", &(_instance.Paragraphs), formDiv, probe)
		case "StaticWebSite:Chapters":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Chapters", func(owner *models.StaticWebSite) *[]*models.StaticWebSiteChapter { return &owner.Chapters })
		}
	}
}

func __gong__New__StaticWebSiteGeneratedImageFormCallback(
	_instance *models.StaticWebSiteGeneratedImage,
	probe *Probe,
	formGroup *form.FormGroup,
) (staticwebsitegeneratedimageFormCallback *FormCallback[*models.StaticWebSiteGeneratedImage]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveStaticWebSiteGeneratedImageFields,
	)
}

type StaticWebSiteGeneratedImageFormCallback = FormCallback[*models.StaticWebSiteGeneratedImage]

func saveStaticWebSiteGeneratedImageFields(
	_instance *models.StaticWebSiteGeneratedImage,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "SourceDirectoryPath":
			FormDivBasicFieldToField(&(_instance.SourceDirectoryPath), formDiv)
		case "Width":
			FormDivBasicFieldToField(&(_instance.Width), formDiv)
		case "Height":
			FormDivBasicFieldToField(&(_instance.Height), formDiv)
		}
	}
}

func __gong__New__StaticWebSiteImageFormCallback(
	_instance *models.StaticWebSiteImage,
	probe *Probe,
	formGroup *form.FormGroup,
) (staticwebsiteimageFormCallback *FormCallback[*models.StaticWebSiteImage]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveStaticWebSiteImageFields,
	)
}

type StaticWebSiteImageFormCallback = FormCallback[*models.StaticWebSiteImage]

func saveStaticWebSiteImageFields(
	_instance *models.StaticWebSiteImage,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "SourceDirectoryPath":
			FormDivBasicFieldToField(&(_instance.SourceDirectoryPath), formDiv)
		case "Width":
			FormDivBasicFieldToField(&(_instance.Width), formDiv)
		case "Height":
			FormDivBasicFieldToField(&(_instance.Height), formDiv)
		}
	}
}

func __gong__New__StaticWebSiteParagraphFormCallback(
	_instance *models.StaticWebSiteParagraph,
	probe *Probe,
	formGroup *form.FormGroup,
) (staticwebsiteparagraphFormCallback *FormCallback[*models.StaticWebSiteParagraph]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveStaticWebSiteParagraphFields,
	)
}

type StaticWebSiteParagraphFormCallback = FormCallback[*models.StaticWebSiteParagraph]

func saveStaticWebSiteParagraphFields(
	_instance *models.StaticWebSiteParagraph,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "LegendMarkdownContent":
			FormDivBasicFieldToField(&(_instance.LegendMarkdownContent), formDiv)
		case "Image":
			FormDivSelectFieldToField(&(_instance.Image), probe.stageOfInterest, formDiv)
		case "StaticWebSiteChapter:Paragraphs":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Paragraphs", func(owner *models.StaticWebSiteChapter) *[]*models.StaticWebSiteParagraph { return &owner.Paragraphs })
		}
	}
}

func __gong__New__XHTML_CONTENTFormCallback(
	_instance *models.XHTML_CONTENT,
	probe *Probe,
	formGroup *form.FormGroup,
) (xhtml_contentFormCallback *FormCallback[*models.XHTML_CONTENT]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveXHTML_CONTENTFields,
	)
}

type XHTML_CONTENTFormCallback = FormCallback[*models.XHTML_CONTENT]

func saveXHTML_CONTENTFields(
	_instance *models.XHTML_CONTENT,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&(_instance.EnclosedText), formDiv)
		case "PureText":
			FormDivBasicFieldToField(&(_instance.PureText), formDiv)
		}
	}
}

