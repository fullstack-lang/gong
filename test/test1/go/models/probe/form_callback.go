// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/test/test1/go/models"
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
func __gong__New__AstructFormCallback(
	_instance *models.Astruct,
	probe *Probe,
	formGroup *form.FormGroup,
) (astructFormCallback *FormCallback[*models.Astruct]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveAstructFields,
	)
}

type AstructFormCallback = FormCallback[*models.Astruct]

func saveAstructFields(
	_instance *models.Astruct,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Associationtob":
			FormDivSelectFieldToField(&(_instance.Associationtob), probe.stageOfInterest, formDiv)
		case "Anarrayofb":
			FormDivSliceOfPointersToField(_instance, "Anarrayofb", &(_instance.Anarrayofb), formDiv, probe)
		case "Anotherassociationtob_2":
			FormDivSelectFieldToField(&(_instance.Anotherassociationtob_2), probe.stageOfInterest, formDiv)
		case "Date":
			FormDivTimeFieldToField(&(_instance.Date), formDiv, false)
		case "Date2":
			FormDivTimeFieldToField(&(_instance.Date2), formDiv, false)
		case "Booleanfield":
			FormDivBasicFieldToField(&(_instance.Booleanfield), formDiv)
		case "Aenum":
			FormDivEnumStringFieldToField(&(_instance.Aenum), formDiv)
		case "Aenum_2":
			FormDivEnumStringFieldToField(&(_instance.Aenum_2), formDiv)
		case "Benum":
			FormDivEnumStringFieldToField(&(_instance.Benum), formDiv)
		case "CEnum":
			FormDivEnumIntFieldToField(&(_instance.CEnum), formDiv)
		case "CName":
			FormDivBasicFieldToField(&(_instance.CName), formDiv)
		case "CFloatfield":
			FormDivBasicFieldToField(&(_instance.CFloatfield), formDiv)
		case "Bstruct":
			FormDivSelectFieldToField(&(_instance.Bstruct), probe.stageOfInterest, formDiv)
		case "Bstruct2":
			FormDivSelectFieldToField(&(_instance.Bstruct2), probe.stageOfInterest, formDiv)
		case "Dstruct":
			FormDivSelectFieldToField(&(_instance.Dstruct), probe.stageOfInterest, formDiv)
		case "Dstruct2":
			FormDivSelectFieldToField(&(_instance.Dstruct2), probe.stageOfInterest, formDiv)
		case "Dstruct3":
			FormDivSelectFieldToField(&(_instance.Dstruct3), probe.stageOfInterest, formDiv)
		case "Dstruct4":
			FormDivSelectFieldToField(&(_instance.Dstruct4), probe.stageOfInterest, formDiv)
		case "Dstruct4s":
			FormDivSliceOfPointersToField(_instance, "Dstruct4s", &(_instance.Dstruct4s), formDiv, probe)
		case "Floatfield":
			FormDivBasicFieldToField(&(_instance.Floatfield), formDiv)
		case "Intfield":
			FormDivBasicFieldToField(&(_instance.Intfield), formDiv)
		case "Anotherbooleanfield":
			FormDivBasicFieldToField(&(_instance.Anotherbooleanfield), formDiv)
		case "Duration1":
			FormDivBasicFieldToField(&(_instance.Duration1), formDiv)
		case "Anarrayofa":
			FormDivSliceOfPointersToField(_instance, "Anarrayofa", &(_instance.Anarrayofa), formDiv, probe)
		case "Anotherarrayofb":
			FormDivSliceOfPointersToField(_instance, "Anotherarrayofb", &(_instance.Anotherarrayofb), formDiv, probe)
		case "AnarrayofbUse":
			FormDivSliceOfPointersToField(_instance, "AnarrayofbUse", &(_instance.AnarrayofbUse), formDiv, probe)
		case "Anarrayofb2Use":
			FormDivSliceOfPointersToField(_instance, "Anarrayofb2Use", &(_instance.Anarrayofb2Use), formDiv, probe)
		case "AnAstruct":
			FormDivSelectFieldToField(&(_instance.AnAstruct), probe.stageOfInterest, formDiv)
		case "TextFieldBespokeSize":
			FormDivBasicFieldToField(&(_instance.TextFieldBespokeSize), formDiv)
		case "TextArea":
			FormDivBasicFieldToField(&(_instance.TextArea), formDiv)
		case "Astruct:Anarrayofa":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Anarrayofa", func(owner *models.Astruct) *[]*models.Astruct { return &owner.Anarrayofa })
		}
	}
}

func __gong__New__AstructBstruct2UseFormCallback(
	_instance *models.AstructBstruct2Use,
	probe *Probe,
	formGroup *form.FormGroup,
) (astructbstruct2useFormCallback *FormCallback[*models.AstructBstruct2Use]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveAstructBstruct2UseFields,
	)
}

type AstructBstruct2UseFormCallback = FormCallback[*models.AstructBstruct2Use]

func saveAstructBstruct2UseFields(
	_instance *models.AstructBstruct2Use,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Bstrcut2":
			FormDivSelectFieldToField(&(_instance.Bstrcut2), probe.stageOfInterest, formDiv)
		case "Astruct:Anarrayofb2Use":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Anarrayofb2Use", func(owner *models.Astruct) *[]*models.AstructBstruct2Use { return &owner.Anarrayofb2Use })
		}
	}
}

func __gong__New__AstructBstructUseFormCallback(
	_instance *models.AstructBstructUse,
	probe *Probe,
	formGroup *form.FormGroup,
) (astructbstructuseFormCallback *FormCallback[*models.AstructBstructUse]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveAstructBstructUseFields,
	)
}

type AstructBstructUseFormCallback = FormCallback[*models.AstructBstructUse]

func saveAstructBstructUseFields(
	_instance *models.AstructBstructUse,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Bstruct2":
			FormDivSelectFieldToField(&(_instance.Bstruct2), probe.stageOfInterest, formDiv)
		case "Astruct:AnarrayofbUse":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "AnarrayofbUse", func(owner *models.Astruct) *[]*models.AstructBstructUse { return &owner.AnarrayofbUse })
		}
	}
}

func __gong__New__BstructFormCallback(
	_instance *models.Bstruct,
	probe *Probe,
	formGroup *form.FormGroup,
) (bstructFormCallback *FormCallback[*models.Bstruct]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveBstructFields,
	)
}

type BstructFormCallback = FormCallback[*models.Bstruct]

func saveBstructFields(
	_instance *models.Bstruct,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Floatfield":
			FormDivBasicFieldToField(&(_instance.Floatfield), formDiv)
		case "Floatfield2":
			FormDivBasicFieldToField(&(_instance.Floatfield2), formDiv)
		case "Intfield":
			FormDivBasicFieldToField(&(_instance.Intfield), formDiv)
		case "Astruct:Anarrayofb":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Anarrayofb", func(owner *models.Astruct) *[]*models.Bstruct { return &owner.Anarrayofb })
		case "Astruct:Anotherarrayofb":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Anotherarrayofb", func(owner *models.Astruct) *[]*models.Bstruct { return &owner.Anotherarrayofb })
		case "Dstruct:Anarrayofb":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Anarrayofb", func(owner *models.Dstruct) *[]*models.Bstruct { return &owner.Anarrayofb })
		}
	}
}

func __gong__New__DstructFormCallback(
	_instance *models.Dstruct,
	probe *Probe,
	formGroup *form.FormGroup,
) (dstructFormCallback *FormCallback[*models.Dstruct]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveDstructFields,
	)
}

type DstructFormCallback = FormCallback[*models.Dstruct]

func saveDstructFields(
	_instance *models.Dstruct,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Anarrayofb":
			FormDivSliceOfPointersToField(_instance, "Anarrayofb", &(_instance.Anarrayofb), formDiv, probe)
		case "Gstruct":
			FormDivSelectFieldToField(&(_instance.Gstruct), probe.stageOfInterest, formDiv)
		case "Gstructs":
			FormDivSliceOfPointersToField(_instance, "Gstructs", &(_instance.Gstructs), formDiv, probe)
		case "Astruct:Dstruct4s":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Dstruct4s", func(owner *models.Astruct) *[]*models.Dstruct { return &owner.Dstruct4s })
		}
	}
}

func __gong__New__F0123456789012345678901234567890FormCallback(
	_instance *models.F0123456789012345678901234567890,
	probe *Probe,
	formGroup *form.FormGroup,
) (f0123456789012345678901234567890FormCallback *FormCallback[*models.F0123456789012345678901234567890]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveF0123456789012345678901234567890Fields,
	)
}

type F0123456789012345678901234567890FormCallback = FormCallback[*models.F0123456789012345678901234567890]

func saveF0123456789012345678901234567890Fields(
	_instance *models.F0123456789012345678901234567890,
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
		}
	}
}

func __gong__New__GstructFormCallback(
	_instance *models.Gstruct,
	probe *Probe,
	formGroup *form.FormGroup,
) (gstructFormCallback *FormCallback[*models.Gstruct]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveGstructFields,
	)
}

type GstructFormCallback = FormCallback[*models.Gstruct]

func saveGstructFields(
	_instance *models.Gstruct,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Floatfield":
			FormDivBasicFieldToField(&(_instance.Floatfield), formDiv)
		case "Floatfield2":
			FormDivBasicFieldToField(&(_instance.Floatfield2), formDiv)
		case "Intfield":
			FormDivBasicFieldToField(&(_instance.Intfield), formDiv)
		case "Dstruct:Gstructs":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Gstructs", func(owner *models.Dstruct) *[]*models.Gstruct { return &owner.Gstructs })
		}
	}
}

