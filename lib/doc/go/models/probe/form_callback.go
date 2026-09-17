// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/lib/doc/go/models"
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
func __gong__New__AttributeShapeFormCallback(
	_instance *models.AttributeShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (attributeshapeFormCallback *FormCallback[*models.AttributeShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveAttributeShapeFields,
	)
}

type AttributeShapeFormCallback = FormCallback[*models.AttributeShape]

func saveAttributeShapeFields(
	_instance *models.AttributeShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "FieldTypeAsString":
			FormDivBasicFieldToField(&(_instance.FieldTypeAsString), formDiv)
		case "Structname":
			FormDivBasicFieldToField(&(_instance.Structname), formDiv)
		case "Fieldtypename":
			FormDivBasicFieldToField(&(_instance.Fieldtypename), formDiv)
		case "GongStructShape:AttributeShapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "AttributeShapes", func(owner *models.GongStructShape) *[]*models.AttributeShape { return &owner.AttributeShapes })
		}
	}
}

func __gong__New__ClassdiagramFormCallback(
	_instance *models.Classdiagram,
	probe *Probe,
	formGroup *form.FormGroup,
) (classdiagramFormCallback *FormCallback[*models.Classdiagram]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveClassdiagramFields,
	)
}

type ClassdiagramFormCallback = FormCallback[*models.Classdiagram]

func saveClassdiagramFields(
	_instance *models.Classdiagram,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Description":
			FormDivBasicFieldToField(&(_instance.Description), formDiv)
		case "IsIncludedInStaticWebSite":
			FormDivBasicFieldToField(&(_instance.IsIncludedInStaticWebSite), formDiv)
		case "GongStructShapes":
			FormDivSliceOfPointersToField(_instance, "GongStructShapes", &(_instance.GongStructShapes), formDiv, probe)
		case "GongEnumShapes":
			FormDivSliceOfPointersToField(_instance, "GongEnumShapes", &(_instance.GongEnumShapes), formDiv, probe)
		case "GongNoteShapes":
			FormDivSliceOfPointersToField(_instance, "GongNoteShapes", &(_instance.GongNoteShapes), formDiv, probe)
		case "ShowNbInstances":
			FormDivBasicFieldToField(&(_instance.ShowNbInstances), formDiv)
		case "ShowMultiplicity":
			FormDivBasicFieldToField(&(_instance.ShowMultiplicity), formDiv)
		case "ShowLinkNames":
			FormDivBasicFieldToField(&(_instance.ShowLinkNames), formDiv)
		case "IsInRenameMode":
			FormDivBasicFieldToField(&(_instance.IsInRenameMode), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(_instance.IsExpanded), formDiv)
		case "NodeGongStructsIsExpanded":
			FormDivBasicFieldToField(&(_instance.NodeGongStructsIsExpanded), formDiv)
		case "NodeGongStructNodeExpansion":
			FormDivBasicFieldToField(&(_instance.NodeGongStructNodeExpansion), formDiv)
		case "NodeGongEnumsIsExpanded":
			FormDivBasicFieldToField(&(_instance.NodeGongEnumsIsExpanded), formDiv)
		case "NodeGongEnumNodeExpansion":
			FormDivBasicFieldToField(&(_instance.NodeGongEnumNodeExpansion), formDiv)
		case "NodeGongNotesIsExpanded":
			FormDivBasicFieldToField(&(_instance.NodeGongNotesIsExpanded), formDiv)
		case "NodeGongNoteNodeExpansion":
			FormDivBasicFieldToField(&(_instance.NodeGongNoteNodeExpansion), formDiv)
		case "DiagramPackage:Classdiagrams":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Classdiagrams", func(owner *models.DiagramPackage) *[]*models.Classdiagram { return &owner.Classdiagrams })
		}
	}
}

func __gong__New__DiagramPackageFormCallback(
	_instance *models.DiagramPackage,
	probe *Probe,
	formGroup *form.FormGroup,
) (diagrampackageFormCallback *FormCallback[*models.DiagramPackage]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveDiagramPackageFields,
	)
}

type DiagramPackageFormCallback = FormCallback[*models.DiagramPackage]

func saveDiagramPackageFields(
	_instance *models.DiagramPackage,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Path":
			FormDivBasicFieldToField(&(_instance.Path), formDiv)
		case "GongModelPath":
			FormDivBasicFieldToField(&(_instance.GongModelPath), formDiv)
		case "Classdiagrams":
			FormDivSliceOfPointersToField(_instance, "Classdiagrams", &(_instance.Classdiagrams), formDiv, probe)
		case "SelectedClassdiagram":
			FormDivSelectFieldToField(&(_instance.SelectedClassdiagram), probe.stageOfInterest, formDiv)
		case "AbsolutePathToDiagramPackage":
			FormDivBasicFieldToField(&(_instance.AbsolutePathToDiagramPackage), formDiv)
		}
	}
}

func __gong__New__GongEnumShapeFormCallback(
	_instance *models.GongEnumShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (gongenumshapeFormCallback *FormCallback[*models.GongEnumShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveGongEnumShapeFields,
	)
}

type GongEnumShapeFormCallback = FormCallback[*models.GongEnumShape]

func saveGongEnumShapeFields(
	_instance *models.GongEnumShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "X":
			FormDivBasicFieldToField(&(_instance.X), formDiv)
		case "Y":
			FormDivBasicFieldToField(&(_instance.Y), formDiv)
		case "Width":
			FormDivBasicFieldToField(&(_instance.Width), formDiv)
		case "Height":
			FormDivBasicFieldToField(&(_instance.Height), formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&(_instance.IsHidden), formDiv)
		case "GongEnumValueShapes":
			FormDivSliceOfPointersToField(_instance, "GongEnumValueShapes", &(_instance.GongEnumValueShapes), formDiv, probe)
		case "IsExpanded":
			FormDivBasicFieldToField(&(_instance.IsExpanded), formDiv)
		case "Classdiagram:GongEnumShapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "GongEnumShapes", func(owner *models.Classdiagram) *[]*models.GongEnumShape { return &owner.GongEnumShapes })
		}
	}
}

func __gong__New__GongEnumValueShapeFormCallback(
	_instance *models.GongEnumValueShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (gongenumvalueshapeFormCallback *FormCallback[*models.GongEnumValueShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveGongEnumValueShapeFields,
	)
}

type GongEnumValueShapeFormCallback = FormCallback[*models.GongEnumValueShape]

func saveGongEnumValueShapeFields(
	_instance *models.GongEnumValueShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "GongEnumShape:GongEnumValueShapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "GongEnumValueShapes", func(owner *models.GongEnumShape) *[]*models.GongEnumValueShape { return &owner.GongEnumValueShapes })
		}
	}
}

func __gong__New__GongNoteLinkShapeFormCallback(
	_instance *models.GongNoteLinkShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (gongnotelinkshapeFormCallback *FormCallback[*models.GongNoteLinkShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveGongNoteLinkShapeFields,
	)
}

type GongNoteLinkShapeFormCallback = FormCallback[*models.GongNoteLinkShape]

func saveGongNoteLinkShapeFields(
	_instance *models.GongNoteLinkShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Identifier":
			FormDivBasicFieldToField(&(_instance.Identifier), formDiv)
		case "Type":
			FormDivEnumStringFieldToField(&(_instance.Type), formDiv)
		case "GongNoteShape:GongNoteLinkShapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "GongNoteLinkShapes", func(owner *models.GongNoteShape) *[]*models.GongNoteLinkShape { return &owner.GongNoteLinkShapes })
		}
	}
}

func __gong__New__GongNoteShapeFormCallback(
	_instance *models.GongNoteShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (gongnoteshapeFormCallback *FormCallback[*models.GongNoteShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveGongNoteShapeFields,
	)
}

type GongNoteShapeFormCallback = FormCallback[*models.GongNoteShape]

func saveGongNoteShapeFields(
	_instance *models.GongNoteShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Identifier":
			FormDivBasicFieldToField(&(_instance.Identifier), formDiv)
		case "Body":
			FormDivBasicFieldToField(&(_instance.Body), formDiv)
		case "BodyHTML":
			FormDivBasicFieldToField(&(_instance.BodyHTML), formDiv)
		case "X":
			FormDivBasicFieldToField(&(_instance.X), formDiv)
		case "Y":
			FormDivBasicFieldToField(&(_instance.Y), formDiv)
		case "Width":
			FormDivBasicFieldToField(&(_instance.Width), formDiv)
		case "Height":
			FormDivBasicFieldToField(&(_instance.Height), formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&(_instance.IsHidden), formDiv)
		case "Matched":
			FormDivBasicFieldToField(&(_instance.Matched), formDiv)
		case "GongNoteLinkShapes":
			FormDivSliceOfPointersToField(_instance, "GongNoteLinkShapes", &(_instance.GongNoteLinkShapes), formDiv, probe)
		case "IsExpanded":
			FormDivBasicFieldToField(&(_instance.IsExpanded), formDiv)
		case "Classdiagram:GongNoteShapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "GongNoteShapes", func(owner *models.Classdiagram) *[]*models.GongNoteShape { return &owner.GongNoteShapes })
		}
	}
}

func __gong__New__GongStructShapeFormCallback(
	_instance *models.GongStructShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (gongstructshapeFormCallback *FormCallback[*models.GongStructShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveGongStructShapeFields,
	)
}

type GongStructShapeFormCallback = FormCallback[*models.GongStructShape]

func saveGongStructShapeFields(
	_instance *models.GongStructShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "X":
			FormDivBasicFieldToField(&(_instance.X), formDiv)
		case "Y":
			FormDivBasicFieldToField(&(_instance.Y), formDiv)
		case "Width":
			FormDivBasicFieldToField(&(_instance.Width), formDiv)
		case "Height":
			FormDivBasicFieldToField(&(_instance.Height), formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&(_instance.IsHidden), formDiv)
		case "AttributeShapes":
			FormDivSliceOfPointersToField(_instance, "AttributeShapes", &(_instance.AttributeShapes), formDiv, probe)
		case "LinkShapes":
			FormDivSliceOfPointersToField(_instance, "LinkShapes", &(_instance.LinkShapes), formDiv, probe)
		case "IsSelected":
			FormDivBasicFieldToField(&(_instance.IsSelected), formDiv)
		case "Classdiagram:GongStructShapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "GongStructShapes", func(owner *models.Classdiagram) *[]*models.GongStructShape { return &owner.GongStructShapes })
		}
	}
}

func __gong__New__LinkShapeFormCallback(
	_instance *models.LinkShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (linkshapeFormCallback *FormCallback[*models.LinkShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveLinkShapeFields,
	)
}

type LinkShapeFormCallback = FormCallback[*models.LinkShape]

func saveLinkShapeFields(
	_instance *models.LinkShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "FieldOffsetX":
			FormDivBasicFieldToField(&(_instance.FieldOffsetX), formDiv)
		case "FieldOffsetY":
			FormDivBasicFieldToField(&(_instance.FieldOffsetY), formDiv)
		case "TargetMultiplicity":
			FormDivEnumStringFieldToField(&(_instance.TargetMultiplicity), formDiv)
		case "TargetMultiplicityOffsetX":
			FormDivBasicFieldToField(&(_instance.TargetMultiplicityOffsetX), formDiv)
		case "TargetMultiplicityOffsetY":
			FormDivBasicFieldToField(&(_instance.TargetMultiplicityOffsetY), formDiv)
		case "SourceMultiplicity":
			FormDivEnumStringFieldToField(&(_instance.SourceMultiplicity), formDiv)
		case "SourceMultiplicityOffsetX":
			FormDivBasicFieldToField(&(_instance.SourceMultiplicityOffsetX), formDiv)
		case "SourceMultiplicityOffsetY":
			FormDivBasicFieldToField(&(_instance.SourceMultiplicityOffsetY), formDiv)
		case "X":
			FormDivBasicFieldToField(&(_instance.X), formDiv)
		case "Y":
			FormDivBasicFieldToField(&(_instance.Y), formDiv)
		case "StartOrientation":
			FormDivEnumStringFieldToField(&(_instance.StartOrientation), formDiv)
		case "StartRatio":
			FormDivBasicFieldToField(&(_instance.StartRatio), formDiv)
		case "EndOrientation":
			FormDivEnumStringFieldToField(&(_instance.EndOrientation), formDiv)
		case "EndRatio":
			FormDivBasicFieldToField(&(_instance.EndRatio), formDiv)
		case "CornerOffsetRatio":
			FormDivBasicFieldToField(&(_instance.CornerOffsetRatio), formDiv)
		case "GongStructShape:LinkShapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "LinkShapes", func(owner *models.GongStructShape) *[]*models.LinkShape { return &owner.LinkShapes })
		}
	}
}

