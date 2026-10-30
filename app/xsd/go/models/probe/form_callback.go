// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	form "github.com/fullstack-lang/gong/lib/form/go/models"
	gongprobe "github.com/fullstack-lang/gong/pkg/runtime/probe"

	"github.com/fullstack-lang/gong/app/xsd/go/models"
)

// to avoid errors when time and slices packages are not used in the generated code
const _ = time.Nanosecond

var _ = slices.Delete([]string{"a"}, 0, 1)

var _ = log.Panicf

type FormCallbackIF = gongprobe.FormCallbackIF


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
func __gong__New__AllFormCallback(
	_instance *models.All,
	probe *Probe,
	formGroup *form.FormGroup,
) (allFormCallback *FormCallback[*models.All]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveAllFields,
	)
}

type AllFormCallback = FormCallback[*models.All]

func saveAllFields(
	_instance *models.All,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Annotation":
			FormDivSelectFieldToField(&(_instance.Annotation), probe.stageOfInterest, formDiv)
		case "OuterElementName":
			FormDivBasicFieldToField(&(_instance.OuterElementName), formDiv)
		case "Sequences":
			FormDivSliceOfPointersToField(_instance, "Sequences", &(_instance.Sequences), formDiv, probe)
		case "Alls":
			FormDivSliceOfPointersToField(_instance, "Alls", &(_instance.Alls), formDiv, probe)
		case "Choices":
			FormDivSliceOfPointersToField(_instance, "Choices", &(_instance.Choices), formDiv, probe)
		case "Groups":
			FormDivSliceOfPointersToField(_instance, "Groups", &(_instance.Groups), formDiv, probe)
		case "Elements":
			FormDivSliceOfPointersToField(_instance, "Elements", &(_instance.Elements), formDiv, probe)
		case "Order":
			FormDivBasicFieldToField(&(_instance.Order), formDiv)
		case "Depth":
			FormDivBasicFieldToField(&(_instance.Depth), formDiv)
		case "MinOccurs":
			FormDivBasicFieldToField(&(_instance.MinOccurs), formDiv)
		case "MaxOccurs":
			FormDivBasicFieldToField(&(_instance.MaxOccurs), formDiv)
		case "All:Alls":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Alls", func(owner *models.All) *[]*models.All { return &owner.Alls })
		case "Choice:Alls":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Alls", func(owner *models.Choice) *[]*models.All { return &owner.Alls })
		case "ComplexType:Alls":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Alls", func(owner *models.ComplexType) *[]*models.All { return &owner.Alls })
		case "Extension:Alls":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Alls", func(owner *models.Extension) *[]*models.All { return &owner.Alls })
		case "Group:Alls":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Alls", func(owner *models.Group) *[]*models.All { return &owner.Alls })
		case "Restriction:Alls":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Alls", func(owner *models.Restriction) *[]*models.All { return &owner.Alls })
		case "Sequence:Alls":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Alls", func(owner *models.Sequence) *[]*models.All { return &owner.Alls })
		}
	}
}

func __gong__New__AnnotationFormCallback(
	_instance *models.Annotation,
	probe *Probe,
	formGroup *form.FormGroup,
) (annotationFormCallback *FormCallback[*models.Annotation]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveAnnotationFields,
	)
}

type AnnotationFormCallback = FormCallback[*models.Annotation]

func saveAnnotationFields(
	_instance *models.Annotation,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Documentations":
			FormDivSliceOfPointersToField(_instance, "Documentations", &(_instance.Documentations), formDiv, probe)
		}
	}
}

func __gong__New__AttributeFormCallback(
	_instance *models.Attribute,
	probe *Probe,
	formGroup *form.FormGroup,
) (attributeFormCallback *FormCallback[*models.Attribute]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveAttributeFields,
	)
}

type AttributeFormCallback = FormCallback[*models.Attribute]

func saveAttributeFields(
	_instance *models.Attribute,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "NameXSD":
			FormDivBasicFieldToField(&(_instance.NameXSD), formDiv)
		case "Type":
			FormDivBasicFieldToField(&(_instance.Type), formDiv)
		case "Annotation":
			FormDivSelectFieldToField(&(_instance.Annotation), probe.stageOfInterest, formDiv)
		case "HasNameConflict":
			FormDivBasicFieldToField(&(_instance.HasNameConflict), formDiv)
		case "GoIdentifier":
			FormDivBasicFieldToField(&(_instance.GoIdentifier), formDiv)
		case "Default":
			FormDivBasicFieldToField(&(_instance.Default), formDiv)
		case "Use":
			FormDivBasicFieldToField(&(_instance.Use), formDiv)
		case "Form":
			FormDivBasicFieldToField(&(_instance.Form), formDiv)
		case "Fixed":
			FormDivBasicFieldToField(&(_instance.Fixed), formDiv)
		case "Ref":
			FormDivBasicFieldToField(&(_instance.Ref), formDiv)
		case "TargetNamespace":
			FormDivBasicFieldToField(&(_instance.TargetNamespace), formDiv)
		case "SimpleType":
			FormDivBasicFieldToField(&(_instance.SimpleType), formDiv)
		case "InnerSimpleType":
			FormDivSelectFieldToField(&(_instance.InnerSimpleType), probe.stageOfInterest, formDiv)
		case "IDXSD":
			FormDivBasicFieldToField(&(_instance.IDXSD), formDiv)
		case "Order":
			FormDivBasicFieldToField(&(_instance.Order), formDiv)
		case "Depth":
			FormDivBasicFieldToField(&(_instance.Depth), formDiv)
		case "AttributeGroup:Attributes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Attributes", func(owner *models.AttributeGroup) *[]*models.Attribute { return &owner.Attributes })
		case "ComplexType:Attributes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Attributes", func(owner *models.ComplexType) *[]*models.Attribute { return &owner.Attributes })
		case "Extension:Attributes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Attributes", func(owner *models.Extension) *[]*models.Attribute { return &owner.Attributes })
		case "Restriction:Attributes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Attributes", func(owner *models.Restriction) *[]*models.Attribute { return &owner.Attributes })
		case "Schema:Attributes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Attributes", func(owner *models.Schema) *[]*models.Attribute { return &owner.Attributes })
		}
	}
}

func __gong__New__AttributeGroupFormCallback(
	_instance *models.AttributeGroup,
	probe *Probe,
	formGroup *form.FormGroup,
) (attributegroupFormCallback *FormCallback[*models.AttributeGroup]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveAttributeGroupFields,
	)
}

type AttributeGroupFormCallback = FormCallback[*models.AttributeGroup]

func saveAttributeGroupFields(
	_instance *models.AttributeGroup,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "NameXSD":
			FormDivBasicFieldToField(&(_instance.NameXSD), formDiv)
		case "Annotation":
			FormDivSelectFieldToField(&(_instance.Annotation), probe.stageOfInterest, formDiv)
		case "HasNameConflict":
			FormDivBasicFieldToField(&(_instance.HasNameConflict), formDiv)
		case "GoIdentifier":
			FormDivBasicFieldToField(&(_instance.GoIdentifier), formDiv)
		case "AttributeGroups":
			FormDivSliceOfPointersToField(_instance, "AttributeGroups", &(_instance.AttributeGroups), formDiv, probe)
		case "Ref":
			FormDivBasicFieldToField(&(_instance.Ref), formDiv)
		case "Attributes":
			FormDivSliceOfPointersToField(_instance, "Attributes", &(_instance.Attributes), formDiv, probe)
		case "Order":
			FormDivBasicFieldToField(&(_instance.Order), formDiv)
		case "Depth":
			FormDivBasicFieldToField(&(_instance.Depth), formDiv)
		case "AttributeGroup:AttributeGroups":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "AttributeGroups", func(owner *models.AttributeGroup) *[]*models.AttributeGroup { return &owner.AttributeGroups })
		case "ComplexType:AttributeGroups":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "AttributeGroups", func(owner *models.ComplexType) *[]*models.AttributeGroup { return &owner.AttributeGroups })
		case "Extension:AttributeGroups":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "AttributeGroups", func(owner *models.Extension) *[]*models.AttributeGroup { return &owner.AttributeGroups })
		case "Restriction:AttributeGroups":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "AttributeGroups", func(owner *models.Restriction) *[]*models.AttributeGroup { return &owner.AttributeGroups })
		case "Schema:AttributeGroups":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "AttributeGroups", func(owner *models.Schema) *[]*models.AttributeGroup { return &owner.AttributeGroups })
		}
	}
}

func __gong__New__ChoiceFormCallback(
	_instance *models.Choice,
	probe *Probe,
	formGroup *form.FormGroup,
) (choiceFormCallback *FormCallback[*models.Choice]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveChoiceFields,
	)
}

type ChoiceFormCallback = FormCallback[*models.Choice]

func saveChoiceFields(
	_instance *models.Choice,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Annotation":
			FormDivSelectFieldToField(&(_instance.Annotation), probe.stageOfInterest, formDiv)
		case "OuterElementName":
			FormDivBasicFieldToField(&(_instance.OuterElementName), formDiv)
		case "Sequences":
			FormDivSliceOfPointersToField(_instance, "Sequences", &(_instance.Sequences), formDiv, probe)
		case "Alls":
			FormDivSliceOfPointersToField(_instance, "Alls", &(_instance.Alls), formDiv, probe)
		case "Choices":
			FormDivSliceOfPointersToField(_instance, "Choices", &(_instance.Choices), formDiv, probe)
		case "Groups":
			FormDivSliceOfPointersToField(_instance, "Groups", &(_instance.Groups), formDiv, probe)
		case "Elements":
			FormDivSliceOfPointersToField(_instance, "Elements", &(_instance.Elements), formDiv, probe)
		case "Order":
			FormDivBasicFieldToField(&(_instance.Order), formDiv)
		case "Depth":
			FormDivBasicFieldToField(&(_instance.Depth), formDiv)
		case "MinOccurs":
			FormDivBasicFieldToField(&(_instance.MinOccurs), formDiv)
		case "MaxOccurs":
			FormDivBasicFieldToField(&(_instance.MaxOccurs), formDiv)
		case "IsDuplicatedInXSD":
			FormDivBasicFieldToField(&(_instance.IsDuplicatedInXSD), formDiv)
		case "All:Choices":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Choices", func(owner *models.All) *[]*models.Choice { return &owner.Choices })
		case "Choice:Choices":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Choices", func(owner *models.Choice) *[]*models.Choice { return &owner.Choices })
		case "ComplexType:Choices":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Choices", func(owner *models.ComplexType) *[]*models.Choice { return &owner.Choices })
		case "Extension:Choices":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Choices", func(owner *models.Extension) *[]*models.Choice { return &owner.Choices })
		case "Group:Choices":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Choices", func(owner *models.Group) *[]*models.Choice { return &owner.Choices })
		case "Restriction:Choices":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Choices", func(owner *models.Restriction) *[]*models.Choice { return &owner.Choices })
		case "Sequence:Choices":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Choices", func(owner *models.Sequence) *[]*models.Choice { return &owner.Choices })
		}
	}
}

func __gong__New__ComplexContentFormCallback(
	_instance *models.ComplexContent,
	probe *Probe,
	formGroup *form.FormGroup,
) (complexcontentFormCallback *FormCallback[*models.ComplexContent]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveComplexContentFields,
	)
}

type ComplexContentFormCallback = FormCallback[*models.ComplexContent]

func saveComplexContentFields(
	_instance *models.ComplexContent,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Annotation":
			FormDivSelectFieldToField(&(_instance.Annotation), probe.stageOfInterest, formDiv)
		case "Mixed":
			FormDivBasicFieldToField(&(_instance.Mixed), formDiv)
		case "Extension":
			FormDivSelectFieldToField(&(_instance.Extension), probe.stageOfInterest, formDiv)
		case "Restriction":
			FormDivSelectFieldToField(&(_instance.Restriction), probe.stageOfInterest, formDiv)
		}
	}
}

func __gong__New__ComplexTypeFormCallback(
	_instance *models.ComplexType,
	probe *Probe,
	formGroup *form.FormGroup,
) (complextypeFormCallback *FormCallback[*models.ComplexType]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveComplexTypeFields,
	)
}

type ComplexTypeFormCallback = FormCallback[*models.ComplexType]

func saveComplexTypeFields(
	_instance *models.ComplexType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "HasNameConflict":
			FormDivBasicFieldToField(&(_instance.HasNameConflict), formDiv)
		case "GoIdentifier":
			FormDivBasicFieldToField(&(_instance.GoIdentifier), formDiv)
		case "IsAnonymous":
			FormDivBasicFieldToField(&(_instance.IsAnonymous), formDiv)
		case "OuterElement":
			FormDivSelectFieldToField(&(_instance.OuterElement), probe.stageOfInterest, formDiv)
		case "Annotation":
			FormDivSelectFieldToField(&(_instance.Annotation), probe.stageOfInterest, formDiv)
		case "NameXSD":
			FormDivBasicFieldToField(&(_instance.NameXSD), formDiv)
		case "OuterElementName":
			FormDivBasicFieldToField(&(_instance.OuterElementName), formDiv)
		case "Sequences":
			FormDivSliceOfPointersToField(_instance, "Sequences", &(_instance.Sequences), formDiv, probe)
		case "Alls":
			FormDivSliceOfPointersToField(_instance, "Alls", &(_instance.Alls), formDiv, probe)
		case "Choices":
			FormDivSliceOfPointersToField(_instance, "Choices", &(_instance.Choices), formDiv, probe)
		case "Groups":
			FormDivSliceOfPointersToField(_instance, "Groups", &(_instance.Groups), formDiv, probe)
		case "Elements":
			FormDivSliceOfPointersToField(_instance, "Elements", &(_instance.Elements), formDiv, probe)
		case "Order":
			FormDivBasicFieldToField(&(_instance.Order), formDiv)
		case "Depth":
			FormDivBasicFieldToField(&(_instance.Depth), formDiv)
		case "MinOccurs":
			FormDivBasicFieldToField(&(_instance.MinOccurs), formDiv)
		case "MaxOccurs":
			FormDivBasicFieldToField(&(_instance.MaxOccurs), formDiv)
		case "Extension":
			FormDivSelectFieldToField(&(_instance.Extension), probe.stageOfInterest, formDiv)
		case "SimpleContent":
			FormDivSelectFieldToField(&(_instance.SimpleContent), probe.stageOfInterest, formDiv)
		case "ComplexContent":
			FormDivSelectFieldToField(&(_instance.ComplexContent), probe.stageOfInterest, formDiv)
		case "Attributes":
			FormDivSliceOfPointersToField(_instance, "Attributes", &(_instance.Attributes), formDiv, probe)
		case "AttributeGroups":
			FormDivSliceOfPointersToField(_instance, "AttributeGroups", &(_instance.AttributeGroups), formDiv, probe)
		case "IsDuplicatedInXSD":
			FormDivBasicFieldToField(&(_instance.IsDuplicatedInXSD), formDiv)
		case "Schema:ComplexTypes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "ComplexTypes", func(owner *models.Schema) *[]*models.ComplexType { return &owner.ComplexTypes })
		}
	}
}

func __gong__New__DocumentationFormCallback(
	_instance *models.Documentation,
	probe *Probe,
	formGroup *form.FormGroup,
) (documentationFormCallback *FormCallback[*models.Documentation]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveDocumentationFields,
	)
}

type DocumentationFormCallback = FormCallback[*models.Documentation]

func saveDocumentationFields(
	_instance *models.Documentation,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Text":
			FormDivBasicFieldToField(&(_instance.Text), formDiv)
		case "Source":
			FormDivBasicFieldToField(&(_instance.Source), formDiv)
		case "Lang":
			FormDivBasicFieldToField(&(_instance.Lang), formDiv)
		case "Annotation:Documentations":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Documentations", func(owner *models.Annotation) *[]*models.Documentation { return &owner.Documentations })
		}
	}
}

func __gong__New__ElementFormCallback(
	_instance *models.Element,
	probe *Probe,
	formGroup *form.FormGroup,
) (elementFormCallback *FormCallback[*models.Element]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveElementFields,
	)
}

type ElementFormCallback = FormCallback[*models.Element]

func saveElementFields(
	_instance *models.Element,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Order":
			FormDivBasicFieldToField(&(_instance.Order), formDiv)
		case "Depth":
			FormDivBasicFieldToField(&(_instance.Depth), formDiv)
		case "HasNameConflict":
			FormDivBasicFieldToField(&(_instance.HasNameConflict), formDiv)
		case "GoIdentifier":
			FormDivBasicFieldToField(&(_instance.GoIdentifier), formDiv)
		case "Annotation":
			FormDivSelectFieldToField(&(_instance.Annotation), probe.stageOfInterest, formDiv)
		case "NameXSD":
			FormDivBasicFieldToField(&(_instance.NameXSD), formDiv)
		case "Type":
			FormDivBasicFieldToField(&(_instance.Type), formDiv)
		case "MinOccurs":
			FormDivBasicFieldToField(&(_instance.MinOccurs), formDiv)
		case "MaxOccurs":
			FormDivBasicFieldToField(&(_instance.MaxOccurs), formDiv)
		case "Default":
			FormDivBasicFieldToField(&(_instance.Default), formDiv)
		case "Fixed":
			FormDivBasicFieldToField(&(_instance.Fixed), formDiv)
		case "Nillable":
			FormDivBasicFieldToField(&(_instance.Nillable), formDiv)
		case "Ref":
			FormDivBasicFieldToField(&(_instance.Ref), formDiv)
		case "Abstract":
			FormDivBasicFieldToField(&(_instance.Abstract), formDiv)
		case "Form":
			FormDivBasicFieldToField(&(_instance.Form), formDiv)
		case "Block":
			FormDivBasicFieldToField(&(_instance.Block), formDiv)
		case "Final":
			FormDivBasicFieldToField(&(_instance.Final), formDiv)
		case "SubstitutionGroup":
			FormDivBasicFieldToField(&(_instance.SubstitutionGroup), formDiv)
		case "SimpleType":
			FormDivSelectFieldToField(&(_instance.SimpleType), probe.stageOfInterest, formDiv)
		case "ComplexType":
			FormDivSelectFieldToField(&(_instance.ComplexType), probe.stageOfInterest, formDiv)
		case "Groups":
			FormDivSliceOfPointersToField(_instance, "Groups", &(_instance.Groups), formDiv, probe)
		case "Keys":
			FormDivSliceOfPointersToField(_instance, "Keys", &(_instance.Keys), formDiv, probe)
		case "Uniques":
			FormDivSliceOfPointersToField(_instance, "Uniques", &(_instance.Uniques), formDiv, probe)
		case "KeyRefs":
			FormDivSliceOfPointersToField(_instance, "KeyRefs", &(_instance.KeyRefs), formDiv, probe)
		case "IsDuplicatedInXSD":
			FormDivBasicFieldToField(&(_instance.IsDuplicatedInXSD), formDiv)
		case "All:Elements":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Elements", func(owner *models.All) *[]*models.Element { return &owner.Elements })
		case "Choice:Elements":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Elements", func(owner *models.Choice) *[]*models.Element { return &owner.Elements })
		case "ComplexType:Elements":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Elements", func(owner *models.ComplexType) *[]*models.Element { return &owner.Elements })
		case "Extension:Elements":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Elements", func(owner *models.Extension) *[]*models.Element { return &owner.Elements })
		case "Group:Elements":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Elements", func(owner *models.Group) *[]*models.Element { return &owner.Elements })
		case "Restriction:Elements":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Elements", func(owner *models.Restriction) *[]*models.Element { return &owner.Elements })
		case "Schema:Elements":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Elements", func(owner *models.Schema) *[]*models.Element { return &owner.Elements })
		case "Sequence:Elements":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Elements", func(owner *models.Sequence) *[]*models.Element { return &owner.Elements })
		}
	}
}

func __gong__New__EnumerationFormCallback(
	_instance *models.Enumeration,
	probe *Probe,
	formGroup *form.FormGroup,
) (enumerationFormCallback *FormCallback[*models.Enumeration]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveEnumerationFields,
	)
}

type EnumerationFormCallback = FormCallback[*models.Enumeration]

func saveEnumerationFields(
	_instance *models.Enumeration,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Annotation":
			FormDivSelectFieldToField(&(_instance.Annotation), probe.stageOfInterest, formDiv)
		case "Value":
			FormDivBasicFieldToField(&(_instance.Value), formDiv)
		case "Restriction:Enumerations":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Enumerations", func(owner *models.Restriction) *[]*models.Enumeration { return &owner.Enumerations })
		}
	}
}

func __gong__New__ExtensionFormCallback(
	_instance *models.Extension,
	probe *Probe,
	formGroup *form.FormGroup,
) (extensionFormCallback *FormCallback[*models.Extension]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveExtensionFields,
	)
}

type ExtensionFormCallback = FormCallback[*models.Extension]

func saveExtensionFields(
	_instance *models.Extension,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "OuterElementName":
			FormDivBasicFieldToField(&(_instance.OuterElementName), formDiv)
		case "Sequences":
			FormDivSliceOfPointersToField(_instance, "Sequences", &(_instance.Sequences), formDiv, probe)
		case "Alls":
			FormDivSliceOfPointersToField(_instance, "Alls", &(_instance.Alls), formDiv, probe)
		case "Choices":
			FormDivSliceOfPointersToField(_instance, "Choices", &(_instance.Choices), formDiv, probe)
		case "Groups":
			FormDivSliceOfPointersToField(_instance, "Groups", &(_instance.Groups), formDiv, probe)
		case "Elements":
			FormDivSliceOfPointersToField(_instance, "Elements", &(_instance.Elements), formDiv, probe)
		case "Order":
			FormDivBasicFieldToField(&(_instance.Order), formDiv)
		case "Depth":
			FormDivBasicFieldToField(&(_instance.Depth), formDiv)
		case "MinOccurs":
			FormDivBasicFieldToField(&(_instance.MinOccurs), formDiv)
		case "MaxOccurs":
			FormDivBasicFieldToField(&(_instance.MaxOccurs), formDiv)
		case "Base":
			FormDivBasicFieldToField(&(_instance.Base), formDiv)
		case "Ref":
			FormDivBasicFieldToField(&(_instance.Ref), formDiv)
		case "Attributes":
			FormDivSliceOfPointersToField(_instance, "Attributes", &(_instance.Attributes), formDiv, probe)
		case "AttributeGroups":
			FormDivSliceOfPointersToField(_instance, "AttributeGroups", &(_instance.AttributeGroups), formDiv, probe)
		}
	}
}

func __gong__New__FieldXPathFormCallback(
	_instance *models.FieldXPath,
	probe *Probe,
	formGroup *form.FormGroup,
) (fieldxpathFormCallback *FormCallback[*models.FieldXPath]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveFieldXPathFields,
	)
}

type FieldXPathFormCallback = FormCallback[*models.FieldXPath]

func saveFieldXPathFields(
	_instance *models.FieldXPath,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "XPath":
			FormDivBasicFieldToField(&(_instance.XPath), formDiv)
		case "Key:Fields":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Fields", func(owner *models.Key) *[]*models.FieldXPath { return &owner.Fields })
		case "KeyRef:Fields":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Fields", func(owner *models.KeyRef) *[]*models.FieldXPath { return &owner.Fields })
		case "Unique:Fields":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Fields", func(owner *models.Unique) *[]*models.FieldXPath { return &owner.Fields })
		}
	}
}

func __gong__New__FractionDigitFormCallback(
	_instance *models.FractionDigit,
	probe *Probe,
	formGroup *form.FormGroup,
) (fractiondigitFormCallback *FormCallback[*models.FractionDigit]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveFractionDigitFields,
	)
}

type FractionDigitFormCallback = FormCallback[*models.FractionDigit]

func saveFractionDigitFields(
	_instance *models.FractionDigit,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Annotation":
			FormDivSelectFieldToField(&(_instance.Annotation), probe.stageOfInterest, formDiv)
		case "Value":
			FormDivBasicFieldToField(&(_instance.Value), formDiv)
		}
	}
}

func __gong__New__GroupFormCallback(
	_instance *models.Group,
	probe *Probe,
	formGroup *form.FormGroup,
) (groupFormCallback *FormCallback[*models.Group]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveGroupFields,
	)
}

type GroupFormCallback = FormCallback[*models.Group]

func saveGroupFields(
	_instance *models.Group,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Annotation":
			FormDivSelectFieldToField(&(_instance.Annotation), probe.stageOfInterest, formDiv)
		case "NameXSD":
			FormDivBasicFieldToField(&(_instance.NameXSD), formDiv)
		case "Ref":
			FormDivBasicFieldToField(&(_instance.Ref), formDiv)
		case "IsAnonymous":
			FormDivBasicFieldToField(&(_instance.IsAnonymous), formDiv)
		case "OuterElement":
			FormDivSelectFieldToField(&(_instance.OuterElement), probe.stageOfInterest, formDiv)
		case "HasNameConflict":
			FormDivBasicFieldToField(&(_instance.HasNameConflict), formDiv)
		case "GoIdentifier":
			FormDivBasicFieldToField(&(_instance.GoIdentifier), formDiv)
		case "OuterElementName":
			FormDivBasicFieldToField(&(_instance.OuterElementName), formDiv)
		case "Sequences":
			FormDivSliceOfPointersToField(_instance, "Sequences", &(_instance.Sequences), formDiv, probe)
		case "Alls":
			FormDivSliceOfPointersToField(_instance, "Alls", &(_instance.Alls), formDiv, probe)
		case "Choices":
			FormDivSliceOfPointersToField(_instance, "Choices", &(_instance.Choices), formDiv, probe)
		case "Groups":
			FormDivSliceOfPointersToField(_instance, "Groups", &(_instance.Groups), formDiv, probe)
		case "Elements":
			FormDivSliceOfPointersToField(_instance, "Elements", &(_instance.Elements), formDiv, probe)
		case "Order":
			FormDivBasicFieldToField(&(_instance.Order), formDiv)
		case "Depth":
			FormDivBasicFieldToField(&(_instance.Depth), formDiv)
		case "MinOccurs":
			FormDivBasicFieldToField(&(_instance.MinOccurs), formDiv)
		case "MaxOccurs":
			FormDivBasicFieldToField(&(_instance.MaxOccurs), formDiv)
		case "All:Groups":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Groups", func(owner *models.All) *[]*models.Group { return &owner.Groups })
		case "Choice:Groups":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Groups", func(owner *models.Choice) *[]*models.Group { return &owner.Groups })
		case "ComplexType:Groups":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Groups", func(owner *models.ComplexType) *[]*models.Group { return &owner.Groups })
		case "Element:Groups":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Groups", func(owner *models.Element) *[]*models.Group { return &owner.Groups })
		case "Extension:Groups":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Groups", func(owner *models.Extension) *[]*models.Group { return &owner.Groups })
		case "Group:Groups":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Groups", func(owner *models.Group) *[]*models.Group { return &owner.Groups })
		case "Restriction:Groups":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Groups", func(owner *models.Restriction) *[]*models.Group { return &owner.Groups })
		case "Schema:Groups":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Groups", func(owner *models.Schema) *[]*models.Group { return &owner.Groups })
		case "Sequence:Groups":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Groups", func(owner *models.Sequence) *[]*models.Group { return &owner.Groups })
		}
	}
}

func __gong__New__KeyFormCallback(
	_instance *models.Key,
	probe *Probe,
	formGroup *form.FormGroup,
) (keyFormCallback *FormCallback[*models.Key]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveKeyFields,
	)
}

type KeyFormCallback = FormCallback[*models.Key]

func saveKeyFields(
	_instance *models.Key,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Annotation":
			FormDivSelectFieldToField(&(_instance.Annotation), probe.stageOfInterest, formDiv)
		case "NameXSD":
			FormDivBasicFieldToField(&(_instance.NameXSD), formDiv)
		case "Selector":
			FormDivSelectFieldToField(&(_instance.Selector), probe.stageOfInterest, formDiv)
		case "Fields":
			FormDivSliceOfPointersToField(_instance, "Fields", &(_instance.Fields), formDiv, probe)
		case "Element:Keys":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Keys", func(owner *models.Element) *[]*models.Key { return &owner.Keys })
		case "Schema:Keys":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Keys", func(owner *models.Schema) *[]*models.Key { return &owner.Keys })
		}
	}
}

func __gong__New__KeyRefFormCallback(
	_instance *models.KeyRef,
	probe *Probe,
	formGroup *form.FormGroup,
) (keyrefFormCallback *FormCallback[*models.KeyRef]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveKeyRefFields,
	)
}

type KeyRefFormCallback = FormCallback[*models.KeyRef]

func saveKeyRefFields(
	_instance *models.KeyRef,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Annotation":
			FormDivSelectFieldToField(&(_instance.Annotation), probe.stageOfInterest, formDiv)
		case "NameXSD":
			FormDivBasicFieldToField(&(_instance.NameXSD), formDiv)
		case "Refer":
			FormDivBasicFieldToField(&(_instance.Refer), formDiv)
		case "Selector":
			FormDivSelectFieldToField(&(_instance.Selector), probe.stageOfInterest, formDiv)
		case "Fields":
			FormDivSliceOfPointersToField(_instance, "Fields", &(_instance.Fields), formDiv, probe)
		case "Element:KeyRefs":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "KeyRefs", func(owner *models.Element) *[]*models.KeyRef { return &owner.KeyRefs })
		case "Schema:KeyRefs":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "KeyRefs", func(owner *models.Schema) *[]*models.KeyRef { return &owner.KeyRefs })
		}
	}
}

func __gong__New__LengthFormCallback(
	_instance *models.Length,
	probe *Probe,
	formGroup *form.FormGroup,
) (lengthFormCallback *FormCallback[*models.Length]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveLengthFields,
	)
}

type LengthFormCallback = FormCallback[*models.Length]

func saveLengthFields(
	_instance *models.Length,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Annotation":
			FormDivSelectFieldToField(&(_instance.Annotation), probe.stageOfInterest, formDiv)
		case "Value":
			FormDivBasicFieldToField(&(_instance.Value), formDiv)
		}
	}
}

func __gong__New__ListFormCallback(
	_instance *models.List,
	probe *Probe,
	formGroup *form.FormGroup,
) (listFormCallback *FormCallback[*models.List]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveListFields,
	)
}

type ListFormCallback = FormCallback[*models.List]

func saveListFields(
	_instance *models.List,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Annotation":
			FormDivSelectFieldToField(&(_instance.Annotation), probe.stageOfInterest, formDiv)
		case "ItemType":
			FormDivBasicFieldToField(&(_instance.ItemType), formDiv)
		case "SimpleType":
			FormDivSelectFieldToField(&(_instance.SimpleType), probe.stageOfInterest, formDiv)
		}
	}
}

func __gong__New__MaxExclusiveFormCallback(
	_instance *models.MaxExclusive,
	probe *Probe,
	formGroup *form.FormGroup,
) (maxexclusiveFormCallback *FormCallback[*models.MaxExclusive]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveMaxExclusiveFields,
	)
}

type MaxExclusiveFormCallback = FormCallback[*models.MaxExclusive]

func saveMaxExclusiveFields(
	_instance *models.MaxExclusive,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Annotation":
			FormDivSelectFieldToField(&(_instance.Annotation), probe.stageOfInterest, formDiv)
		case "Value":
			FormDivBasicFieldToField(&(_instance.Value), formDiv)
		}
	}
}

func __gong__New__MaxInclusiveFormCallback(
	_instance *models.MaxInclusive,
	probe *Probe,
	formGroup *form.FormGroup,
) (maxinclusiveFormCallback *FormCallback[*models.MaxInclusive]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveMaxInclusiveFields,
	)
}

type MaxInclusiveFormCallback = FormCallback[*models.MaxInclusive]

func saveMaxInclusiveFields(
	_instance *models.MaxInclusive,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Annotation":
			FormDivSelectFieldToField(&(_instance.Annotation), probe.stageOfInterest, formDiv)
		case "Value":
			FormDivBasicFieldToField(&(_instance.Value), formDiv)
		}
	}
}

func __gong__New__MaxLengthFormCallback(
	_instance *models.MaxLength,
	probe *Probe,
	formGroup *form.FormGroup,
) (maxlengthFormCallback *FormCallback[*models.MaxLength]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveMaxLengthFields,
	)
}

type MaxLengthFormCallback = FormCallback[*models.MaxLength]

func saveMaxLengthFields(
	_instance *models.MaxLength,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Annotation":
			FormDivSelectFieldToField(&(_instance.Annotation), probe.stageOfInterest, formDiv)
		case "Value":
			FormDivBasicFieldToField(&(_instance.Value), formDiv)
		}
	}
}

func __gong__New__MinExclusiveFormCallback(
	_instance *models.MinExclusive,
	probe *Probe,
	formGroup *form.FormGroup,
) (minexclusiveFormCallback *FormCallback[*models.MinExclusive]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveMinExclusiveFields,
	)
}

type MinExclusiveFormCallback = FormCallback[*models.MinExclusive]

func saveMinExclusiveFields(
	_instance *models.MinExclusive,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Annotation":
			FormDivSelectFieldToField(&(_instance.Annotation), probe.stageOfInterest, formDiv)
		case "Value":
			FormDivBasicFieldToField(&(_instance.Value), formDiv)
		}
	}
}

func __gong__New__MinInclusiveFormCallback(
	_instance *models.MinInclusive,
	probe *Probe,
	formGroup *form.FormGroup,
) (mininclusiveFormCallback *FormCallback[*models.MinInclusive]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveMinInclusiveFields,
	)
}

type MinInclusiveFormCallback = FormCallback[*models.MinInclusive]

func saveMinInclusiveFields(
	_instance *models.MinInclusive,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Annotation":
			FormDivSelectFieldToField(&(_instance.Annotation), probe.stageOfInterest, formDiv)
		case "Value":
			FormDivBasicFieldToField(&(_instance.Value), formDiv)
		}
	}
}

func __gong__New__MinLengthFormCallback(
	_instance *models.MinLength,
	probe *Probe,
	formGroup *form.FormGroup,
) (minlengthFormCallback *FormCallback[*models.MinLength]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveMinLengthFields,
	)
}

type MinLengthFormCallback = FormCallback[*models.MinLength]

func saveMinLengthFields(
	_instance *models.MinLength,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Annotation":
			FormDivSelectFieldToField(&(_instance.Annotation), probe.stageOfInterest, formDiv)
		case "Value":
			FormDivBasicFieldToField(&(_instance.Value), formDiv)
		}
	}
}

func __gong__New__PatternFormCallback(
	_instance *models.Pattern,
	probe *Probe,
	formGroup *form.FormGroup,
) (patternFormCallback *FormCallback[*models.Pattern]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		savePatternFields,
	)
}

type PatternFormCallback = FormCallback[*models.Pattern]

func savePatternFields(
	_instance *models.Pattern,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Annotation":
			FormDivSelectFieldToField(&(_instance.Annotation), probe.stageOfInterest, formDiv)
		case "Value":
			FormDivBasicFieldToField(&(_instance.Value), formDiv)
		}
	}
}

func __gong__New__RestrictionFormCallback(
	_instance *models.Restriction,
	probe *Probe,
	formGroup *form.FormGroup,
) (restrictionFormCallback *FormCallback[*models.Restriction]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveRestrictionFields,
	)
}

type RestrictionFormCallback = FormCallback[*models.Restriction]

func saveRestrictionFields(
	_instance *models.Restriction,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Annotation":
			FormDivSelectFieldToField(&(_instance.Annotation), probe.stageOfInterest, formDiv)
		case "Base":
			FormDivBasicFieldToField(&(_instance.Base), formDiv)
		case "SimpleType":
			FormDivSelectFieldToField(&(_instance.SimpleType), probe.stageOfInterest, formDiv)
		case "Enumerations":
			FormDivSliceOfPointersToField(_instance, "Enumerations", &(_instance.Enumerations), formDiv, probe)
		case "MinInclusive":
			FormDivSelectFieldToField(&(_instance.MinInclusive), probe.stageOfInterest, formDiv)
		case "MaxInclusive":
			FormDivSelectFieldToField(&(_instance.MaxInclusive), probe.stageOfInterest, formDiv)
		case "MinExclusive":
			FormDivSelectFieldToField(&(_instance.MinExclusive), probe.stageOfInterest, formDiv)
		case "MaxExclusive":
			FormDivSelectFieldToField(&(_instance.MaxExclusive), probe.stageOfInterest, formDiv)
		case "Pattern":
			FormDivSelectFieldToField(&(_instance.Pattern), probe.stageOfInterest, formDiv)
		case "WhiteSpace":
			FormDivSelectFieldToField(&(_instance.WhiteSpace), probe.stageOfInterest, formDiv)
		case "MinLength":
			FormDivSelectFieldToField(&(_instance.MinLength), probe.stageOfInterest, formDiv)
		case "MaxLength":
			FormDivSelectFieldToField(&(_instance.MaxLength), probe.stageOfInterest, formDiv)
		case "Length":
			FormDivSelectFieldToField(&(_instance.Length), probe.stageOfInterest, formDiv)
		case "TotalDigit":
			FormDivSelectFieldToField(&(_instance.TotalDigit), probe.stageOfInterest, formDiv)
		case "FractionDigit":
			FormDivSelectFieldToField(&(_instance.FractionDigit), probe.stageOfInterest, formDiv)
		case "OuterElementName":
			FormDivBasicFieldToField(&(_instance.OuterElementName), formDiv)
		case "Sequences":
			FormDivSliceOfPointersToField(_instance, "Sequences", &(_instance.Sequences), formDiv, probe)
		case "Alls":
			FormDivSliceOfPointersToField(_instance, "Alls", &(_instance.Alls), formDiv, probe)
		case "Choices":
			FormDivSliceOfPointersToField(_instance, "Choices", &(_instance.Choices), formDiv, probe)
		case "Groups":
			FormDivSliceOfPointersToField(_instance, "Groups", &(_instance.Groups), formDiv, probe)
		case "Elements":
			FormDivSliceOfPointersToField(_instance, "Elements", &(_instance.Elements), formDiv, probe)
		case "Order":
			FormDivBasicFieldToField(&(_instance.Order), formDiv)
		case "Depth":
			FormDivBasicFieldToField(&(_instance.Depth), formDiv)
		case "MinOccurs":
			FormDivBasicFieldToField(&(_instance.MinOccurs), formDiv)
		case "MaxOccurs":
			FormDivBasicFieldToField(&(_instance.MaxOccurs), formDiv)
		case "Attributes":
			FormDivSliceOfPointersToField(_instance, "Attributes", &(_instance.Attributes), formDiv, probe)
		case "AttributeGroups":
			FormDivSliceOfPointersToField(_instance, "AttributeGroups", &(_instance.AttributeGroups), formDiv, probe)
		}
	}
}

func __gong__New__SchemaFormCallback(
	_instance *models.Schema,
	probe *Probe,
	formGroup *form.FormGroup,
) (schemaFormCallback *FormCallback[*models.Schema]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveSchemaFields,
	)
}

type SchemaFormCallback = FormCallback[*models.Schema]

func saveSchemaFields(
	_instance *models.Schema,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Xs":
			FormDivBasicFieldToField(&(_instance.Xs), formDiv)
		case "TargetNamespace":
			FormDivBasicFieldToField(&(_instance.TargetNamespace), formDiv)
		case "ElementFormDefault":
			FormDivBasicFieldToField(&(_instance.ElementFormDefault), formDiv)
		case "AttributeFormDefault":
			FormDivBasicFieldToField(&(_instance.AttributeFormDefault), formDiv)
		case "Version":
			FormDivBasicFieldToField(&(_instance.Version), formDiv)
		case "FilePath":
			FormDivBasicFieldToField(&(_instance.FilePath), formDiv)
		case "PackageName":
			FormDivBasicFieldToField(&(_instance.PackageName), formDiv)
		case "PackageRelDir":
			FormDivBasicFieldToField(&(_instance.PackageRelDir), formDiv)
		case "IsRootSchema":
			FormDivBasicFieldToField(&(_instance.IsRootSchema), formDiv)
		case "Annotation":
			FormDivSelectFieldToField(&(_instance.Annotation), probe.stageOfInterest, formDiv)
		case "Imports":
			FormDivSliceOfPointersToField(_instance, "Imports", &(_instance.Imports), formDiv, probe)
		case "Includes":
			FormDivSliceOfPointersToField(_instance, "Includes", &(_instance.Includes), formDiv, probe)
		case "Elements":
			FormDivSliceOfPointersToField(_instance, "Elements", &(_instance.Elements), formDiv, probe)
		case "Attributes":
			FormDivSliceOfPointersToField(_instance, "Attributes", &(_instance.Attributes), formDiv, probe)
		case "SimpleTypes":
			FormDivSliceOfPointersToField(_instance, "SimpleTypes", &(_instance.SimpleTypes), formDiv, probe)
		case "ComplexTypes":
			FormDivSliceOfPointersToField(_instance, "ComplexTypes", &(_instance.ComplexTypes), formDiv, probe)
		case "AttributeGroups":
			FormDivSliceOfPointersToField(_instance, "AttributeGroups", &(_instance.AttributeGroups), formDiv, probe)
		case "Groups":
			FormDivSliceOfPointersToField(_instance, "Groups", &(_instance.Groups), formDiv, probe)
		case "Keys":
			FormDivSliceOfPointersToField(_instance, "Keys", &(_instance.Keys), formDiv, probe)
		case "Uniques":
			FormDivSliceOfPointersToField(_instance, "Uniques", &(_instance.Uniques), formDiv, probe)
		case "KeyRefs":
			FormDivSliceOfPointersToField(_instance, "KeyRefs", &(_instance.KeyRefs), formDiv, probe)
		case "Order":
			FormDivBasicFieldToField(&(_instance.Order), formDiv)
		case "Depth":
			FormDivBasicFieldToField(&(_instance.Depth), formDiv)
		}
	}
}

func __gong__New__SelectorFormCallback(
	_instance *models.Selector,
	probe *Probe,
	formGroup *form.FormGroup,
) (selectorFormCallback *FormCallback[*models.Selector]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveSelectorFields,
	)
}

type SelectorFormCallback = FormCallback[*models.Selector]

func saveSelectorFields(
	_instance *models.Selector,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "XPath":
			FormDivBasicFieldToField(&(_instance.XPath), formDiv)
		}
	}
}

func __gong__New__SequenceFormCallback(
	_instance *models.Sequence,
	probe *Probe,
	formGroup *form.FormGroup,
) (sequenceFormCallback *FormCallback[*models.Sequence]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveSequenceFields,
	)
}

type SequenceFormCallback = FormCallback[*models.Sequence]

func saveSequenceFields(
	_instance *models.Sequence,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Annotation":
			FormDivSelectFieldToField(&(_instance.Annotation), probe.stageOfInterest, formDiv)
		case "OuterElementName":
			FormDivBasicFieldToField(&(_instance.OuterElementName), formDiv)
		case "Sequences":
			FormDivSliceOfPointersToField(_instance, "Sequences", &(_instance.Sequences), formDiv, probe)
		case "Alls":
			FormDivSliceOfPointersToField(_instance, "Alls", &(_instance.Alls), formDiv, probe)
		case "Choices":
			FormDivSliceOfPointersToField(_instance, "Choices", &(_instance.Choices), formDiv, probe)
		case "Groups":
			FormDivSliceOfPointersToField(_instance, "Groups", &(_instance.Groups), formDiv, probe)
		case "Elements":
			FormDivSliceOfPointersToField(_instance, "Elements", &(_instance.Elements), formDiv, probe)
		case "Order":
			FormDivBasicFieldToField(&(_instance.Order), formDiv)
		case "Depth":
			FormDivBasicFieldToField(&(_instance.Depth), formDiv)
		case "MinOccurs":
			FormDivBasicFieldToField(&(_instance.MinOccurs), formDiv)
		case "MaxOccurs":
			FormDivBasicFieldToField(&(_instance.MaxOccurs), formDiv)
		case "All:Sequences":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Sequences", func(owner *models.All) *[]*models.Sequence { return &owner.Sequences })
		case "Choice:Sequences":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Sequences", func(owner *models.Choice) *[]*models.Sequence { return &owner.Sequences })
		case "ComplexType:Sequences":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Sequences", func(owner *models.ComplexType) *[]*models.Sequence { return &owner.Sequences })
		case "Extension:Sequences":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Sequences", func(owner *models.Extension) *[]*models.Sequence { return &owner.Sequences })
		case "Group:Sequences":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Sequences", func(owner *models.Group) *[]*models.Sequence { return &owner.Sequences })
		case "Restriction:Sequences":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Sequences", func(owner *models.Restriction) *[]*models.Sequence { return &owner.Sequences })
		case "Sequence:Sequences":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Sequences", func(owner *models.Sequence) *[]*models.Sequence { return &owner.Sequences })
		}
	}
}

func __gong__New__SimpleContentFormCallback(
	_instance *models.SimpleContent,
	probe *Probe,
	formGroup *form.FormGroup,
) (simplecontentFormCallback *FormCallback[*models.SimpleContent]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveSimpleContentFields,
	)
}

type SimpleContentFormCallback = FormCallback[*models.SimpleContent]

func saveSimpleContentFields(
	_instance *models.SimpleContent,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Extension":
			FormDivSelectFieldToField(&(_instance.Extension), probe.stageOfInterest, formDiv)
		case "Restriction":
			FormDivSelectFieldToField(&(_instance.Restriction), probe.stageOfInterest, formDiv)
		}
	}
}

func __gong__New__SimpleTypeFormCallback(
	_instance *models.SimpleType,
	probe *Probe,
	formGroup *form.FormGroup,
) (simpletypeFormCallback *FormCallback[*models.SimpleType]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveSimpleTypeFields,
	)
}

type SimpleTypeFormCallback = FormCallback[*models.SimpleType]

func saveSimpleTypeFields(
	_instance *models.SimpleType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Annotation":
			FormDivSelectFieldToField(&(_instance.Annotation), probe.stageOfInterest, formDiv)
		case "NameXSD":
			FormDivBasicFieldToField(&(_instance.NameXSD), formDiv)
		case "Restriction":
			FormDivSelectFieldToField(&(_instance.Restriction), probe.stageOfInterest, formDiv)
		case "Union":
			FormDivSelectFieldToField(&(_instance.Union), probe.stageOfInterest, formDiv)
		case "List":
			FormDivSelectFieldToField(&(_instance.List), probe.stageOfInterest, formDiv)
		case "HasNameConflict":
			FormDivBasicFieldToField(&(_instance.HasNameConflict), formDiv)
		case "GoIdentifier":
			FormDivBasicFieldToField(&(_instance.GoIdentifier), formDiv)
		case "Order":
			FormDivBasicFieldToField(&(_instance.Order), formDiv)
		case "Depth":
			FormDivBasicFieldToField(&(_instance.Depth), formDiv)
		case "Schema:SimpleTypes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "SimpleTypes", func(owner *models.Schema) *[]*models.SimpleType { return &owner.SimpleTypes })
		case "Union:SimpleTypes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "SimpleTypes", func(owner *models.Union) *[]*models.SimpleType { return &owner.SimpleTypes })
		}
	}
}

func __gong__New__TotalDigitFormCallback(
	_instance *models.TotalDigit,
	probe *Probe,
	formGroup *form.FormGroup,
) (totaldigitFormCallback *FormCallback[*models.TotalDigit]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveTotalDigitFields,
	)
}

type TotalDigitFormCallback = FormCallback[*models.TotalDigit]

func saveTotalDigitFields(
	_instance *models.TotalDigit,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Annotation":
			FormDivSelectFieldToField(&(_instance.Annotation), probe.stageOfInterest, formDiv)
		case "Value":
			FormDivBasicFieldToField(&(_instance.Value), formDiv)
		}
	}
}

func __gong__New__UnionFormCallback(
	_instance *models.Union,
	probe *Probe,
	formGroup *form.FormGroup,
) (unionFormCallback *FormCallback[*models.Union]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveUnionFields,
	)
}

type UnionFormCallback = FormCallback[*models.Union]

func saveUnionFields(
	_instance *models.Union,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Annotation":
			FormDivSelectFieldToField(&(_instance.Annotation), probe.stageOfInterest, formDiv)
		case "MemberTypes":
			FormDivBasicFieldToField(&(_instance.MemberTypes), formDiv)
		case "SimpleTypes":
			FormDivSliceOfPointersToField(_instance, "SimpleTypes", &(_instance.SimpleTypes), formDiv, probe)
		}
	}
}

func __gong__New__UniqueFormCallback(
	_instance *models.Unique,
	probe *Probe,
	formGroup *form.FormGroup,
) (uniqueFormCallback *FormCallback[*models.Unique]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveUniqueFields,
	)
}

type UniqueFormCallback = FormCallback[*models.Unique]

func saveUniqueFields(
	_instance *models.Unique,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Annotation":
			FormDivSelectFieldToField(&(_instance.Annotation), probe.stageOfInterest, formDiv)
		case "NameXSD":
			FormDivBasicFieldToField(&(_instance.NameXSD), formDiv)
		case "Selector":
			FormDivSelectFieldToField(&(_instance.Selector), probe.stageOfInterest, formDiv)
		case "Fields":
			FormDivSliceOfPointersToField(_instance, "Fields", &(_instance.Fields), formDiv, probe)
		case "Element:Uniques":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Uniques", func(owner *models.Element) *[]*models.Unique { return &owner.Uniques })
		case "Schema:Uniques":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Uniques", func(owner *models.Schema) *[]*models.Unique { return &owner.Uniques })
		}
	}
}

func __gong__New__WhiteSpaceFormCallback(
	_instance *models.WhiteSpace,
	probe *Probe,
	formGroup *form.FormGroup,
) (whitespaceFormCallback *FormCallback[*models.WhiteSpace]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveWhiteSpaceFields,
	)
}

type WhiteSpaceFormCallback = FormCallback[*models.WhiteSpace]

func saveWhiteSpaceFields(
	_instance *models.WhiteSpace,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Annotation":
			FormDivSelectFieldToField(&(_instance.Annotation), probe.stageOfInterest, formDiv)
		case "Value":
			FormDivBasicFieldToField(&(_instance.Value), formDiv)
		}
	}
}

func __gong__New__XSDImportFormCallback(
	_instance *models.XSDImport,
	probe *Probe,
	formGroup *form.FormGroup,
) (xsdimportFormCallback *FormCallback[*models.XSDImport]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveXSDImportFields,
	)
}

type XSDImportFormCallback = FormCallback[*models.XSDImport]

func saveXSDImportFields(
	_instance *models.XSDImport,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Annotation":
			FormDivSelectFieldToField(&(_instance.Annotation), probe.stageOfInterest, formDiv)
		case "Namespace":
			FormDivBasicFieldToField(&(_instance.Namespace), formDiv)
		case "SchemaLocation":
			FormDivBasicFieldToField(&(_instance.SchemaLocation), formDiv)
		case "Schema:Imports":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Imports", func(owner *models.Schema) *[]*models.XSDImport { return &owner.Imports })
		}
	}
}

func __gong__New__XSDIncludeFormCallback(
	_instance *models.XSDInclude,
	probe *Probe,
	formGroup *form.FormGroup,
) (xsdincludeFormCallback *FormCallback[*models.XSDInclude]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveXSDIncludeFields,
	)
}

type XSDIncludeFormCallback = FormCallback[*models.XSDInclude]

func saveXSDIncludeFields(
	_instance *models.XSDInclude,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Annotation":
			FormDivSelectFieldToField(&(_instance.Annotation), probe.stageOfInterest, formDiv)
		case "SchemaLocation":
			FormDivBasicFieldToField(&(_instance.SchemaLocation), formDiv)
		case "Schema:Includes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Includes", func(owner *models.Schema) *[]*models.XSDInclude { return &owner.Includes })
		}
	}
}

