// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/go/models"
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
func __gong__New__GongBasicFieldFormCallback(
	_instance *models.GongBasicField,
	probe *Probe,
	formGroup *form.FormGroup,
) (gongbasicfieldFormCallback *FormCallback[*models.GongBasicField]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveGongBasicFieldFields,
	)
}

type GongBasicFieldFormCallback = FormCallback[*models.GongBasicField]

func saveGongBasicFieldFields(
	_instance *models.GongBasicField,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "BasicKindName":
			FormDivBasicFieldToField(&(_instance.BasicKindName), formDiv)
		case "GongEnum":
			FormDivSelectFieldToField(&(_instance.GongEnum), probe.stageOfInterest, formDiv)
		case "DeclaredType":
			FormDivBasicFieldToField(&(_instance.DeclaredType), formDiv)
		case "CompositeStructName":
			FormDivBasicFieldToField(&(_instance.CompositeStructName), formDiv)
		case "IsAccordionStart":
			FormDivBasicFieldToField(&(_instance.IsAccordionStart), formDiv)
		case "AccordionName":
			FormDivBasicFieldToField(&(_instance.AccordionName), formDiv)
		case "IsAccordionEnd":
			FormDivBasicFieldToField(&(_instance.IsAccordionEnd), formDiv)
		case "Index":
			FormDivBasicFieldToField(&(_instance.Index), formDiv)
		case "IsTextArea":
			FormDivBasicFieldToField(&(_instance.IsTextArea), formDiv)
		case "IsBespokeWidth":
			FormDivBasicFieldToField(&(_instance.IsBespokeWidth), formDiv)
		case "BespokeWidth":
			FormDivBasicFieldToField(&(_instance.BespokeWidth), formDiv)
		case "IsBespokeHeight":
			FormDivBasicFieldToField(&(_instance.IsBespokeHeight), formDiv)
		case "BespokeHeight":
			FormDivBasicFieldToField(&(_instance.BespokeHeight), formDiv)
		case "GongStruct:GongBasicFields":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "GongBasicFields", func(owner *models.GongStruct) *[]*models.GongBasicField { return &owner.GongBasicFields })
		}
	}
}

func __gong__New__GongEnumFormCallback(
	_instance *models.GongEnum,
	probe *Probe,
	formGroup *form.FormGroup,
) (gongenumFormCallback *FormCallback[*models.GongEnum]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveGongEnumFields,
	)
}

type GongEnumFormCallback = FormCallback[*models.GongEnum]

func saveGongEnumFields(
	_instance *models.GongEnum,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Type":
			FormDivEnumIntFieldToField(&(_instance.Type), formDiv)
		case "GongEnumValues":
			FormDivSliceOfPointersToField(_instance, "GongEnumValues", &(_instance.GongEnumValues), formDiv, probe)
		}
	}
}

func __gong__New__GongEnumValueFormCallback(
	_instance *models.GongEnumValue,
	probe *Probe,
	formGroup *form.FormGroup,
) (gongenumvalueFormCallback *FormCallback[*models.GongEnumValue]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveGongEnumValueFields,
	)
}

type GongEnumValueFormCallback = FormCallback[*models.GongEnumValue]

func saveGongEnumValueFields(
	_instance *models.GongEnumValue,
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
		case "GongEnum:GongEnumValues":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "GongEnumValues", func(owner *models.GongEnum) *[]*models.GongEnumValue { return &owner.GongEnumValues })
		}
	}
}

func __gong__New__GongLinkFormCallback(
	_instance *models.GongLink,
	probe *Probe,
	formGroup *form.FormGroup,
) (gonglinkFormCallback *FormCallback[*models.GongLink]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveGongLinkFields,
	)
}

type GongLinkFormCallback = FormCallback[*models.GongLink]

func saveGongLinkFields(
	_instance *models.GongLink,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Recv":
			FormDivBasicFieldToField(&(_instance.Recv), formDiv)
		case "ImportPath":
			FormDivBasicFieldToField(&(_instance.ImportPath), formDiv)
		case "GongNote:Links":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Links", func(owner *models.GongNote) *[]*models.GongLink { return &owner.Links })
		}
	}
}

func __gong__New__GongNoteFormCallback(
	_instance *models.GongNote,
	probe *Probe,
	formGroup *form.FormGroup,
) (gongnoteFormCallback *FormCallback[*models.GongNote]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveGongNoteFields,
	)
}

type GongNoteFormCallback = FormCallback[*models.GongNote]

func saveGongNoteFields(
	_instance *models.GongNote,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Body":
			FormDivBasicFieldToField(&(_instance.Body), formDiv)
		case "BodyHTML":
			FormDivBasicFieldToField(&(_instance.BodyHTML), formDiv)
		case "Links":
			FormDivSliceOfPointersToField(_instance, "Links", &(_instance.Links), formDiv, probe)
		}
	}
}

func __gong__New__GongStructFormCallback(
	_instance *models.GongStruct,
	probe *Probe,
	formGroup *form.FormGroup,
) (gongstructFormCallback *FormCallback[*models.GongStruct]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveGongStructFields,
	)
}

type GongStructFormCallback = FormCallback[*models.GongStruct]

func saveGongStructFields(
	_instance *models.GongStruct,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "GongBasicFields":
			FormDivSliceOfPointersToField(_instance, "GongBasicFields", &(_instance.GongBasicFields), formDiv, probe)
		case "GongTimeFields":
			FormDivSliceOfPointersToField(_instance, "GongTimeFields", &(_instance.GongTimeFields), formDiv, probe)
		case "PointerToGongStructFields":
			FormDivSliceOfPointersToField(_instance, "PointerToGongStructFields", &(_instance.PointerToGongStructFields), formDiv, probe)
		case "SliceOfPointerToGongStructFields":
			FormDivSliceOfPointersToField(_instance, "SliceOfPointerToGongStructFields", &(_instance.SliceOfPointerToGongStructFields), formDiv, probe)
		case "HasOnAfterUpdateSignature":
			FormDivBasicFieldToField(&(_instance.HasOnAfterUpdateSignature), formDiv)
		case "IsIgnoredForFront":
			FormDivBasicFieldToField(&(_instance.IsIgnoredForFront), formDiv)
		case "IsOmittedForMarshalling":
			FormDivBasicFieldToField(&(_instance.IsOmittedForMarshalling), formDiv)
		}
	}
}

func __gong__New__GongTimeFieldFormCallback(
	_instance *models.GongTimeField,
	probe *Probe,
	formGroup *form.FormGroup,
) (gongtimefieldFormCallback *FormCallback[*models.GongTimeField]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveGongTimeFieldFields,
	)
}

type GongTimeFieldFormCallback = FormCallback[*models.GongTimeField]

func saveGongTimeFieldFields(
	_instance *models.GongTimeField,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Index":
			FormDivBasicFieldToField(&(_instance.Index), formDiv)
		case "CompositeStructName":
			FormDivBasicFieldToField(&(_instance.CompositeStructName), formDiv)
		case "IsAccordionStart":
			FormDivBasicFieldToField(&(_instance.IsAccordionStart), formDiv)
		case "AccordionName":
			FormDivBasicFieldToField(&(_instance.AccordionName), formDiv)
		case "IsAccordionEnd":
			FormDivBasicFieldToField(&(_instance.IsAccordionEnd), formDiv)
		case "BespokeTimeFormat":
			FormDivBasicFieldToField(&(_instance.BespokeTimeFormat), formDiv)
		case "TimeFormOnly":
			FormDivBasicFieldToField(&(_instance.TimeFormOnly), formDiv)
		case "GongStruct:GongTimeFields":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "GongTimeFields", func(owner *models.GongStruct) *[]*models.GongTimeField { return &owner.GongTimeFields })
		}
	}
}

func __gong__New__MetaReferenceFormCallback(
	_instance *models.MetaReference,
	probe *Probe,
	formGroup *form.FormGroup,
) (metareferenceFormCallback *FormCallback[*models.MetaReference]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveMetaReferenceFields,
	)
}

type MetaReferenceFormCallback = FormCallback[*models.MetaReference]

func saveMetaReferenceFields(
	_instance *models.MetaReference,
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

func __gong__New__ModelPkgFormCallback(
	_instance *models.ModelPkg,
	probe *Probe,
	formGroup *form.FormGroup,
) (modelpkgFormCallback *FormCallback[*models.ModelPkg]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveModelPkgFields,
	)
}

type ModelPkgFormCallback = FormCallback[*models.ModelPkg]

func saveModelPkgFields(
	_instance *models.ModelPkg,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "PkgGoName":
			FormDivBasicFieldToField(&(_instance.PkgGoName), formDiv)
		case "PkgPath":
			FormDivBasicFieldToField(&(_instance.PkgPath), formDiv)
		case "PathToGoSubDirectory":
			FormDivBasicFieldToField(&(_instance.PathToGoSubDirectory), formDiv)
		case "OrmPkgGenPath":
			FormDivBasicFieldToField(&(_instance.OrmPkgGenPath), formDiv)
		case "DbOrmPkgGenPath":
			FormDivBasicFieldToField(&(_instance.DbOrmPkgGenPath), formDiv)
		case "DbLiteOrmPkgGenPath":
			FormDivBasicFieldToField(&(_instance.DbLiteOrmPkgGenPath), formDiv)
		case "DbPkgGenPath":
			FormDivBasicFieldToField(&(_instance.DbPkgGenPath), formDiv)
		case "ControllersPkgGenPath":
			FormDivBasicFieldToField(&(_instance.ControllersPkgGenPath), formDiv)
		case "FullstackPkgGenPath":
			FormDivBasicFieldToField(&(_instance.FullstackPkgGenPath), formDiv)
		case "StackPkgGenPath":
			FormDivBasicFieldToField(&(_instance.StackPkgGenPath), formDiv)
		case "Level1StackPkgGenPath":
			FormDivBasicFieldToField(&(_instance.Level1StackPkgGenPath), formDiv)
		case "StaticPkgGenPath":
			FormDivBasicFieldToField(&(_instance.StaticPkgGenPath), formDiv)
		case "ProbePkgGenPath":
			FormDivBasicFieldToField(&(_instance.ProbePkgGenPath), formDiv)
		case "NgWorkspacePath":
			FormDivBasicFieldToField(&(_instance.NgWorkspacePath), formDiv)
		case "NgWorkspaceName":
			FormDivBasicFieldToField(&(_instance.NgWorkspaceName), formDiv)
		case "NgDataLibrarySourceCodeDirectory":
			FormDivBasicFieldToField(&(_instance.NgDataLibrarySourceCodeDirectory), formDiv)
		case "NgSpecificLibrarySourceCodeDirectory":
			FormDivBasicFieldToField(&(_instance.NgSpecificLibrarySourceCodeDirectory), formDiv)
		case "MaterialLibDatamodelTargetPath":
			FormDivBasicFieldToField(&(_instance.MaterialLibDatamodelTargetPath), formDiv)
		}
	}
}

func __gong__New__PointerToGongStructFieldFormCallback(
	_instance *models.PointerToGongStructField,
	probe *Probe,
	formGroup *form.FormGroup,
) (pointertogongstructfieldFormCallback *FormCallback[*models.PointerToGongStructField]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		savePointerToGongStructFieldFields,
	)
}

type PointerToGongStructFieldFormCallback = FormCallback[*models.PointerToGongStructField]

func savePointerToGongStructFieldFields(
	_instance *models.PointerToGongStructField,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "GongStruct":
			FormDivSelectFieldToField(&(_instance.GongStruct), probe.stageOfInterest, formDiv)
		case "Index":
			FormDivBasicFieldToField(&(_instance.Index), formDiv)
		case "CompositeStructName":
			FormDivBasicFieldToField(&(_instance.CompositeStructName), formDiv)
		case "IsAccordionStart":
			FormDivBasicFieldToField(&(_instance.IsAccordionStart), formDiv)
		case "AccordionName":
			FormDivBasicFieldToField(&(_instance.AccordionName), formDiv)
		case "IsAccordionEnd":
			FormDivBasicFieldToField(&(_instance.IsAccordionEnd), formDiv)
		case "IsType":
			FormDivBasicFieldToField(&(_instance.IsType), formDiv)
		case "GongStruct:PointerToGongStructFields":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "PointerToGongStructFields", func(owner *models.GongStruct) *[]*models.PointerToGongStructField { return &owner.PointerToGongStructFields })
		}
	}
}

func __gong__New__SliceOfPointerToGongStructFieldFormCallback(
	_instance *models.SliceOfPointerToGongStructField,
	probe *Probe,
	formGroup *form.FormGroup,
) (sliceofpointertogongstructfieldFormCallback *FormCallback[*models.SliceOfPointerToGongStructField]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveSliceOfPointerToGongStructFieldFields,
	)
}

type SliceOfPointerToGongStructFieldFormCallback = FormCallback[*models.SliceOfPointerToGongStructField]

func saveSliceOfPointerToGongStructFieldFields(
	_instance *models.SliceOfPointerToGongStructField,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "GongStruct":
			FormDivSelectFieldToField(&(_instance.GongStruct), probe.stageOfInterest, formDiv)
		case "Index":
			FormDivBasicFieldToField(&(_instance.Index), formDiv)
		case "CompositeStructName":
			FormDivBasicFieldToField(&(_instance.CompositeStructName), formDiv)
		case "IsAccordionStart":
			FormDivBasicFieldToField(&(_instance.IsAccordionStart), formDiv)
		case "AccordionName":
			FormDivBasicFieldToField(&(_instance.AccordionName), formDiv)
		case "IsAccordionEnd":
			FormDivBasicFieldToField(&(_instance.IsAccordionEnd), formDiv)
		case "GongStruct:SliceOfPointerToGongStructFields":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "SliceOfPointerToGongStructFields", func(owner *models.GongStruct) *[]*models.SliceOfPointerToGongStructField { return &owner.SliceOfPointerToGongStructFields })
		}
	}
}

func __gong__New__StageSetFieldFormCallback(
	_instance *models.StageSetField,
	probe *Probe,
	formGroup *form.FormGroup,
) (stagesetfieldFormCallback *FormCallback[*models.StageSetField]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveStageSetFieldFields,
	)
}

type StageSetFieldFormCallback = FormCallback[*models.StageSetField]

func saveStageSetFieldFields(
	_instance *models.StageSetField,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "PackageName":
			FormDivBasicFieldToField(&(_instance.PackageName), formDiv)
		case "PackagePath":
			FormDivBasicFieldToField(&(_instance.PackagePath), formDiv)
		case "IsLocal":
			FormDivBasicFieldToField(&(_instance.IsLocal), formDiv)
		case "ImportAlias":
			FormDivBasicFieldToField(&(_instance.ImportAlias), formDiv)
		case "StageSetModel:Fields":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Fields", func(owner *models.StageSetModel) *[]*models.StageSetField { return &owner.Fields })
		}
	}
}

func __gong__New__StageSetModelFormCallback(
	_instance *models.StageSetModel,
	probe *Probe,
	formGroup *form.FormGroup,
) (stagesetmodelFormCallback *FormCallback[*models.StageSetModel]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveStageSetModelFields,
	)
}

type StageSetModelFormCallback = FormCallback[*models.StageSetModel]

func saveStageSetModelFields(
	_instance *models.StageSetModel,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Fields":
			FormDivSliceOfPointersToField(_instance, "Fields", &(_instance.Fields), formDiv, probe)
		case "IsManual":
			FormDivBasicFieldToField(&(_instance.IsManual), formDiv)
		}
	}
}

