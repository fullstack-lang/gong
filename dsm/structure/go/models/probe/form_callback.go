// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/dsm/structure/go/models"
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
func __gong__New__AllocatedResourceShapeFormCallback(
	_instance *models.AllocatedResourceShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (allocatedresourceshapeFormCallback *FormCallback[*models.AllocatedResourceShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveAllocatedResourceShapeFields,
	)
}

type AllocatedResourceShapeFormCallback = FormCallback[*models.AllocatedResourceShape]

func saveAllocatedResourceShapeFields(
	_instance *models.AllocatedResourceShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Part":
			FormDivSelectFieldToField(&(_instance.Part), probe.stageOfInterest, formDiv)
		case "Resource":
			FormDivSelectFieldToField(&(_instance.Resource), probe.stageOfInterest, formDiv)
		case "DiagramStructure:AllocatedResourceShapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "AllocatedResourceShapes", func(owner *models.DiagramStructure) *[]*models.AllocatedResourceShape { return &owner.AllocatedResourceShapes })
		}
	}
}

func __gong__New__AllocatedSystemShapeFormCallback(
	_instance *models.AllocatedSystemShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (allocatedsystemshapeFormCallback *FormCallback[*models.AllocatedSystemShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveAllocatedSystemShapeFields,
	)
}

type AllocatedSystemShapeFormCallback = FormCallback[*models.AllocatedSystemShape]

func saveAllocatedSystemShapeFields(
	_instance *models.AllocatedSystemShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Part":
			FormDivSelectFieldToField(&(_instance.Part), probe.stageOfInterest, formDiv)
		case "System":
			FormDivSelectFieldToField(&(_instance.System), probe.stageOfInterest, formDiv)
		case "DiagramStructure:AllocatedSystemShapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "AllocatedSystemShapes", func(owner *models.DiagramStructure) *[]*models.AllocatedSystemShape { return &owner.AllocatedSystemShapes })
		}
	}
}

func __gong__New__ControlFlowFormCallback(
	_instance *models.ControlFlow,
	probe *Probe,
	formGroup *form.FormGroup,
) (controlflowFormCallback *FormCallback[*models.ControlFlow]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveControlFlowFields,
	)
}

type ControlFlowFormCallback = FormCallback[*models.ControlFlow]

func saveControlFlowFields(
	_instance *models.ControlFlow,
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
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(_instance.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(_instance.IsExpanded), formDiv)
		case "Start":
			FormDivSelectFieldToField(&(_instance.Start), probe.stageOfInterest, formDiv)
		case "End":
			FormDivSelectFieldToField(&(_instance.End), probe.stageOfInterest, formDiv)
		case "DiagramStructure:ControlFlowsWhoseNodeIsExpanded":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "ControlFlowsWhoseNodeIsExpanded", func(owner *models.DiagramStructure) *[]*models.ControlFlow { return &owner.ControlFlowsWhoseNodeIsExpanded })
		case "Part:ControlFlows":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "ControlFlows", func(owner *models.Part) *[]*models.ControlFlow { return &owner.ControlFlows })
		}
	}
}

func __gong__New__ControlFlowShapeFormCallback(
	_instance *models.ControlFlowShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (controlflowshapeFormCallback *FormCallback[*models.ControlFlowShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveControlFlowShapeFields,
	)
}

type ControlFlowShapeFormCallback = FormCallback[*models.ControlFlowShape]

func saveControlFlowShapeFields(
	_instance *models.ControlFlowShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "ControlFlow":
			FormDivSelectFieldToField(&(_instance.ControlFlow), probe.stageOfInterest, formDiv)
		case "StartRatio":
			FormDivBasicFieldToField(&(_instance.StartRatio), formDiv)
		case "EndRatio":
			FormDivBasicFieldToField(&(_instance.EndRatio), formDiv)
		case "StartOrientation":
			FormDivEnumStringFieldToField(&(_instance.StartOrientation), formDiv)
		case "EndOrientation":
			FormDivEnumStringFieldToField(&(_instance.EndOrientation), formDiv)
		case "CornerOffsetRatio":
			FormDivBasicFieldToField(&(_instance.CornerOffsetRatio), formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&(_instance.IsHidden), formDiv)
		case "DiagramStructure:ControlFlow_Shapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "ControlFlow_Shapes", func(owner *models.DiagramStructure) *[]*models.ControlFlowShape { return &owner.ControlFlow_Shapes })
		}
	}
}

func __gong__New__DataFormCallback(
	_instance *models.Data,
	probe *Probe,
	formGroup *form.FormGroup,
) (dataFormCallback *FormCallback[*models.Data]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveDataFields,
	)
}

type DataFormCallback = FormCallback[*models.Data]

func saveDataFields(
	_instance *models.Data,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Acronym":
			FormDivBasicFieldToField(&(_instance.Acronym), formDiv)
		case "Description":
			FormDivBasicFieldToField(&(_instance.Description), formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(_instance.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(_instance.IsExpanded), formDiv)
		case "SVG_Path":
			FormDivBasicFieldToField(&(_instance.SVG_Path), formDiv)
		case "InverseAppliedScaling":
			FormDivBasicFieldToField(&(_instance.InverseAppliedScaling), formDiv)
		case "DataFlow:Datas":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Datas", func(owner *models.DataFlow) *[]*models.Data { return &owner.Datas })
		case "DiagramStructure:DatasWhoseNodeIsExpanded":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "DatasWhoseNodeIsExpanded", func(owner *models.DiagramStructure) *[]*models.Data { return &owner.DatasWhoseNodeIsExpanded })
		case "Library:RootDatas":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "RootDatas", func(owner *models.Library) *[]*models.Data { return &owner.RootDatas })
		case "Library:DatasWhoseNodeIsExpanded":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "DatasWhoseNodeIsExpanded", func(owner *models.Library) *[]*models.Data { return &owner.DatasWhoseNodeIsExpanded })
		}
	}
}

func __gong__New__DataFlowFormCallback(
	_instance *models.DataFlow,
	probe *Probe,
	formGroup *form.FormGroup,
) (dataflowFormCallback *FormCallback[*models.DataFlow]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveDataFlowFields,
	)
}

type DataFlowFormCallback = FormCallback[*models.DataFlow]

func saveDataFlowFields(
	_instance *models.DataFlow,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "StartPort":
			FormDivSelectFieldToField(&(_instance.StartPort), probe.stageOfInterest, formDiv)
		case "EndPort":
			FormDivSelectFieldToField(&(_instance.EndPort), probe.stageOfInterest, formDiv)
		case "StartExternalPart":
			FormDivSelectFieldToField(&(_instance.StartExternalPart), probe.stageOfInterest, formDiv)
		case "EndExternalPart":
			FormDivSelectFieldToField(&(_instance.EndExternalPart), probe.stageOfInterest, formDiv)
		case "Datas":
			FormDivSliceOfPointersToField(_instance, "Datas", &(_instance.Datas), formDiv, probe)
		case "Description":
			FormDivBasicFieldToField(&(_instance.Description), formDiv)
		case "Direction":
			FormDivEnumStringFieldToField(&(_instance.Direction), formDiv)
		case "IsDatasNodeExpanded":
			FormDivBasicFieldToField(&(_instance.IsDatasNodeExpanded), formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(_instance.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(_instance.IsExpanded), formDiv)
		case "Type":
			FormDivEnumStringFieldToField(&(_instance.Type), formDiv)
		case "DiagramStructure:DataFlowsWhoseNodeIsExpanded":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "DataFlowsWhoseNodeIsExpanded", func(owner *models.DiagramStructure) *[]*models.DataFlow { return &owner.DataFlowsWhoseNodeIsExpanded })
		case "DiagramStructure:DataFlowsWhoseDataNodeIsExpanded":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "DataFlowsWhoseDataNodeIsExpanded", func(owner *models.DiagramStructure) *[]*models.DataFlow { return &owner.DataFlowsWhoseDataNodeIsExpanded })
		case "Library:RootDataFlows":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "RootDataFlows", func(owner *models.Library) *[]*models.DataFlow { return &owner.RootDataFlows })
		case "Library:DataFlowsWhoseNodeIsExpanded":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "DataFlowsWhoseNodeIsExpanded", func(owner *models.Library) *[]*models.DataFlow { return &owner.DataFlowsWhoseNodeIsExpanded })
		case "System:DataFlows":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "DataFlows", func(owner *models.System) *[]*models.DataFlow { return &owner.DataFlows })
		}
	}
}

func __gong__New__DataFlowShapeFormCallback(
	_instance *models.DataFlowShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (dataflowshapeFormCallback *FormCallback[*models.DataFlowShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveDataFlowShapeFields,
	)
}

type DataFlowShapeFormCallback = FormCallback[*models.DataFlowShape]

func saveDataFlowShapeFields(
	_instance *models.DataFlowShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "DataFlow":
			FormDivSelectFieldToField(&(_instance.DataFlow), probe.stageOfInterest, formDiv)
		case "StartRatio":
			FormDivBasicFieldToField(&(_instance.StartRatio), formDiv)
		case "EndRatio":
			FormDivBasicFieldToField(&(_instance.EndRatio), formDiv)
		case "StartOrientation":
			FormDivEnumStringFieldToField(&(_instance.StartOrientation), formDiv)
		case "EndOrientation":
			FormDivEnumStringFieldToField(&(_instance.EndOrientation), formDiv)
		case "CornerOffsetRatio":
			FormDivBasicFieldToField(&(_instance.CornerOffsetRatio), formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&(_instance.IsHidden), formDiv)
		case "DiagramStructure:DataFlow_Shapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "DataFlow_Shapes", func(owner *models.DiagramStructure) *[]*models.DataFlowShape { return &owner.DataFlow_Shapes })
		}
	}
}

func __gong__New__DataShapeFormCallback(
	_instance *models.DataShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (datashapeFormCallback *FormCallback[*models.DataShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveDataShapeFields,
	)
}

type DataShapeFormCallback = FormCallback[*models.DataShape]

func saveDataShapeFields(
	_instance *models.DataShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Data":
			FormDivSelectFieldToField(&(_instance.Data), probe.stageOfInterest, formDiv)
		case "DataFlow":
			FormDivSelectFieldToField(&(_instance.DataFlow), probe.stageOfInterest, formDiv)
		case "DiagramStructure:Data_Shapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Data_Shapes", func(owner *models.DiagramStructure) *[]*models.DataShape { return &owner.Data_Shapes })
		}
	}
}

func __gong__New__DiagramLayerStateFormCallback(
	_instance *models.DiagramLayerState,
	probe *Probe,
	formGroup *form.FormGroup,
) (diagramlayerstateFormCallback *FormCallback[*models.DiagramLayerState]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveDiagramLayerStateFields,
	)
}

type DiagramLayerStateFormCallback = FormCallback[*models.DiagramLayerState]

func saveDiagramLayerStateFields(
	_instance *models.DiagramLayerState,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "DiagramStructure":
			FormDivSelectFieldToField(&(_instance.DiagramStructure), probe.stageOfInterest, formDiv)
		case "LayerDefinition":
			FormDivSelectFieldToField(&(_instance.LayerDefinition), probe.stageOfInterest, formDiv)
		}
	}
}

func __gong__New__DiagramStructureFormCallback(
	_instance *models.DiagramStructure,
	probe *Probe,
	formGroup *form.FormGroup,
) (diagramstructureFormCallback *FormCallback[*models.DiagramStructure]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveDiagramStructureFields,
	)
}

type DiagramStructureFormCallback = FormCallback[*models.DiagramStructure]

func saveDiagramStructureFields(
	_instance *models.DiagramStructure,
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
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(_instance.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(_instance.IsExpanded), formDiv)
		case "IsChecked":
			FormDivBasicFieldToField(&(_instance.IsChecked), formDiv)
		case "IsEditable_":
			FormDivBasicFieldToField(&(_instance.IsEditable_), formDiv)
		case "IsShowPrefix":
			FormDivBasicFieldToField(&(_instance.IsShowPrefix), formDiv)
		case "DefaultBoxWidth":
			FormDivBasicFieldToField(&(_instance.DefaultBoxWidth), formDiv)
		case "DefaultBoxHeigth":
			FormDivBasicFieldToField(&(_instance.DefaultBoxHeigth), formDiv)
		case "IsWithDiscretePorts":
			FormDivBasicFieldToField(&(_instance.IsWithDiscretePorts), formDiv)
		case "Width":
			FormDivBasicFieldToField(&(_instance.Width), formDiv)
		case "Height":
			FormDivBasicFieldToField(&(_instance.Height), formDiv)
		case "System_Shapes":
			FormDivSliceOfPointersToField(_instance, "System_Shapes", &(_instance.System_Shapes), formDiv, probe)
		case "IsSystemsNodeExpanded":
			FormDivBasicFieldToField(&(_instance.IsSystemsNodeExpanded), formDiv)
		case "SystemsWhoseNodeIsExpanded":
			FormDivSliceOfPointersToField(_instance, "SystemsWhoseNodeIsExpanded", &(_instance.SystemsWhoseNodeIsExpanded), formDiv, probe)
		case "Part_Shapes":
			FormDivSliceOfPointersToField(_instance, "Part_Shapes", &(_instance.Part_Shapes), formDiv, probe)
		case "IsPartsNodeExpanded":
			FormDivBasicFieldToField(&(_instance.IsPartsNodeExpanded), formDiv)
		case "PartWhoseNodeIsExpanded":
			FormDivSliceOfPointersToField(_instance, "PartWhoseNodeIsExpanded", &(_instance.PartWhoseNodeIsExpanded), formDiv, probe)
		case "ExternalPart_Shapes":
			FormDivSliceOfPointersToField(_instance, "ExternalPart_Shapes", &(_instance.ExternalPart_Shapes), formDiv, probe)
		case "IsExternalPartsNodeExpanded":
			FormDivBasicFieldToField(&(_instance.IsExternalPartsNodeExpanded), formDiv)
		case "ExternalPartWhoseNodeIsExpanded":
			FormDivSliceOfPointersToField(_instance, "ExternalPartWhoseNodeIsExpanded", &(_instance.ExternalPartWhoseNodeIsExpanded), formDiv, probe)
		case "ExternalPartsWhoseOutDataFlowsNodeIsExpanded":
			FormDivSliceOfPointersToField(_instance, "ExternalPartsWhoseOutDataFlowsNodeIsExpanded", &(_instance.ExternalPartsWhoseOutDataFlowsNodeIsExpanded), formDiv, probe)
		case "ExternalPartsWhoseInDataFlowsNodeIsExpanded":
			FormDivSliceOfPointersToField(_instance, "ExternalPartsWhoseInDataFlowsNodeIsExpanded", &(_instance.ExternalPartsWhoseInDataFlowsNodeIsExpanded), formDiv, probe)
		case "PortsWhoseNodeIsExpanded":
			FormDivSliceOfPointersToField(_instance, "PortsWhoseNodeIsExpanded", &(_instance.PortsWhoseNodeIsExpanded), formDiv, probe)
		case "Port_Shapes":
			FormDivSliceOfPointersToField(_instance, "Port_Shapes", &(_instance.Port_Shapes), formDiv, probe)
		case "ControlFlowsWhoseNodeIsExpanded":
			FormDivSliceOfPointersToField(_instance, "ControlFlowsWhoseNodeIsExpanded", &(_instance.ControlFlowsWhoseNodeIsExpanded), formDiv, probe)
		case "ControlFlow_Shapes":
			FormDivSliceOfPointersToField(_instance, "ControlFlow_Shapes", &(_instance.ControlFlow_Shapes), formDiv, probe)
		case "DataFlowsWhoseNodeIsExpanded":
			FormDivSliceOfPointersToField(_instance, "DataFlowsWhoseNodeIsExpanded", &(_instance.DataFlowsWhoseNodeIsExpanded), formDiv, probe)
		case "DataFlow_Shapes":
			FormDivSliceOfPointersToField(_instance, "DataFlow_Shapes", &(_instance.DataFlow_Shapes), formDiv, probe)
		case "DatasWhoseNodeIsExpanded":
			FormDivSliceOfPointersToField(_instance, "DatasWhoseNodeIsExpanded", &(_instance.DatasWhoseNodeIsExpanded), formDiv, probe)
		case "Data_Shapes":
			FormDivSliceOfPointersToField(_instance, "Data_Shapes", &(_instance.Data_Shapes), formDiv, probe)
		case "DataFlowsWhoseDataNodeIsExpanded":
			FormDivSliceOfPointersToField(_instance, "DataFlowsWhoseDataNodeIsExpanded", &(_instance.DataFlowsWhoseDataNodeIsExpanded), formDiv, probe)
		case "AllocatedResourcesWhoseNodeIsExpanded":
			FormDivSliceOfPointersToField(_instance, "AllocatedResourcesWhoseNodeIsExpanded", &(_instance.AllocatedResourcesWhoseNodeIsExpanded), formDiv, probe)
		case "AllocatedResourceShapes":
			FormDivSliceOfPointersToField(_instance, "AllocatedResourceShapes", &(_instance.AllocatedResourceShapes), formDiv, probe)
		case "AllocatedSystemesWhoseNodeIsExpanded":
			FormDivSliceOfPointersToField(_instance, "AllocatedSystemesWhoseNodeIsExpanded", &(_instance.AllocatedSystemesWhoseNodeIsExpanded), formDiv, probe)
		case "AllocatedSystemShapes":
			FormDivSliceOfPointersToField(_instance, "AllocatedSystemShapes", &(_instance.AllocatedSystemShapes), formDiv, probe)
		case "Note_Shapes":
			FormDivSliceOfPointersToField(_instance, "Note_Shapes", &(_instance.Note_Shapes), formDiv, probe)
		case "NotesWhoseNodeIsExpanded":
			FormDivSliceOfPointersToField(_instance, "NotesWhoseNodeIsExpanded", &(_instance.NotesWhoseNodeIsExpanded), formDiv, probe)
		case "IsNotesNodeExpanded":
			FormDivBasicFieldToField(&(_instance.IsNotesNodeExpanded), formDiv)
		case "NotePortShapes":
			FormDivSliceOfPointersToField(_instance, "NotePortShapes", &(_instance.NotePortShapes), formDiv, probe)
		case "NotePartShapes":
			FormDivSliceOfPointersToField(_instance, "NotePartShapes", &(_instance.NotePartShapes), formDiv, probe)
		case "System:DiagramStructures":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "DiagramStructures", func(owner *models.System) *[]*models.DiagramStructure { return &owner.DiagramStructures })
		case "System:DiagramStructureWhoseNodeIsExpanded":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "DiagramStructureWhoseNodeIsExpanded", func(owner *models.System) *[]*models.DiagramStructure { return &owner.DiagramStructureWhoseNodeIsExpanded })
		}
	}
}

func __gong__New__ExternalPartShapeFormCallback(
	_instance *models.ExternalPartShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (externalpartshapeFormCallback *FormCallback[*models.ExternalPartShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveExternalPartShapeFields,
	)
}

type ExternalPartShapeFormCallback = FormCallback[*models.ExternalPartShape]

func saveExternalPartShapeFields(
	_instance *models.ExternalPartShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Part":
			FormDivSelectFieldToField(&(_instance.Part), probe.stageOfInterest, formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(_instance.IsExpanded), formDiv)
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
		case "TailHeigth":
			FormDivBasicFieldToField(&(_instance.TailHeigth), formDiv)
		case "DiagramStructure:ExternalPart_Shapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "ExternalPart_Shapes", func(owner *models.DiagramStructure) *[]*models.ExternalPartShape { return &owner.ExternalPart_Shapes })
		}
	}
}

func __gong__New__LayerDefinitionFormCallback(
	_instance *models.LayerDefinition,
	probe *Probe,
	formGroup *form.FormGroup,
) (layerdefinitionFormCallback *FormCallback[*models.LayerDefinition]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveLayerDefinitionFields,
	)
}

type LayerDefinitionFormCallback = FormCallback[*models.LayerDefinition]

func saveLayerDefinitionFields(
	_instance *models.LayerDefinition,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Query":
			FormDivSliceOfPointersToField(_instance, "Query", &(_instance.Query), formDiv, probe)
		}
	}
}

func __gong__New__LibraryFormCallback(
	_instance *models.Library,
	probe *Probe,
	formGroup *form.FormGroup,
) (libraryFormCallback *FormCallback[*models.Library]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveLibraryFields,
	)
}

type LibraryFormCallback = FormCallback[*models.Library]

func saveLibraryFields(
	_instance *models.Library,
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
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(_instance.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(_instance.IsExpanded), formDiv)
		case "IsRootLibrary":
			FormDivBasicFieldToField(&(_instance.IsRootLibrary), formDiv)
		case "SubLibraries":
			FormDivSliceOfPointersToField(_instance, "SubLibraries", &(_instance.SubLibraries), formDiv, probe)
		case "IsSubLibrariesNodeExpanded":
			FormDivBasicFieldToField(&(_instance.IsSubLibrariesNodeExpanded), formDiv)
		case "SubLibrariesWhoseNodeIsExpanded":
			FormDivSliceOfPointersToField(_instance, "SubLibrariesWhoseNodeIsExpanded", &(_instance.SubLibrariesWhoseNodeIsExpanded), formDiv, probe)
		case "NbPixPerCharacter":
			FormDivBasicFieldToField(&(_instance.NbPixPerCharacter), formDiv)
		case "LogoSVGFile":
			FormDivBasicFieldToField(&(_instance.LogoSVGFile), formDiv)
		case "RootSystemes":
			FormDivSliceOfPointersToField(_instance, "RootSystemes", &(_instance.RootSystemes), formDiv, probe)
		case "IsSystemesNodeExpanded":
			FormDivBasicFieldToField(&(_instance.IsSystemesNodeExpanded), formDiv)
		case "SystemsWhoseNodeIsExpanded":
			FormDivSliceOfPointersToField(_instance, "SystemsWhoseNodeIsExpanded", &(_instance.SystemsWhoseNodeIsExpanded), formDiv, probe)
		case "RootDataFlows":
			FormDivSliceOfPointersToField(_instance, "RootDataFlows", &(_instance.RootDataFlows), formDiv, probe)
		case "IsDataFlowsNodeExpanded":
			FormDivBasicFieldToField(&(_instance.IsDataFlowsNodeExpanded), formDiv)
		case "DataFlowsWhoseNodeIsExpanded":
			FormDivSliceOfPointersToField(_instance, "DataFlowsWhoseNodeIsExpanded", &(_instance.DataFlowsWhoseNodeIsExpanded), formDiv, probe)
		case "RootDatas":
			FormDivSliceOfPointersToField(_instance, "RootDatas", &(_instance.RootDatas), formDiv, probe)
		case "IsDatasNodeExpanded":
			FormDivBasicFieldToField(&(_instance.IsDatasNodeExpanded), formDiv)
		case "DatasWhoseNodeIsExpanded":
			FormDivSliceOfPointersToField(_instance, "DatasWhoseNodeIsExpanded", &(_instance.DatasWhoseNodeIsExpanded), formDiv, probe)
		case "RootResources":
			FormDivSliceOfPointersToField(_instance, "RootResources", &(_instance.RootResources), formDiv, probe)
		case "IsResourcesNodeExpanded":
			FormDivBasicFieldToField(&(_instance.IsResourcesNodeExpanded), formDiv)
		case "ResourcesWhoseNodeIsExpanded":
			FormDivSliceOfPointersToField(_instance, "ResourcesWhoseNodeIsExpanded", &(_instance.ResourcesWhoseNodeIsExpanded), formDiv, probe)
		case "PartsWhoseNodeIsExpanded":
			FormDivSliceOfPointersToField(_instance, "PartsWhoseNodeIsExpanded", &(_instance.PartsWhoseNodeIsExpanded), formDiv, probe)
		case "RootNotes":
			FormDivSliceOfPointersToField(_instance, "RootNotes", &(_instance.RootNotes), formDiv, probe)
		case "IsNotesNodeExpanded":
			FormDivBasicFieldToField(&(_instance.IsNotesNodeExpanded), formDiv)
		case "NotesWhoseNodeIsExpanded":
			FormDivSliceOfPointersToField(_instance, "NotesWhoseNodeIsExpanded", &(_instance.NotesWhoseNodeIsExpanded), formDiv, probe)
		case "IsExpandedTmp":
			FormDivBasicFieldToField(&(_instance.IsExpandedTmp), formDiv)
		case "Library:SubLibraries":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "SubLibraries", func(owner *models.Library) *[]*models.Library { return &owner.SubLibraries })
		case "Library:SubLibrariesWhoseNodeIsExpanded":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "SubLibrariesWhoseNodeIsExpanded", func(owner *models.Library) *[]*models.Library { return &owner.SubLibrariesWhoseNodeIsExpanded })
		}
	}
}

func __gong__New__NoteFormCallback(
	_instance *models.Note,
	probe *Probe,
	formGroup *form.FormGroup,
) (noteFormCallback *FormCallback[*models.Note]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveNoteFields,
	)
}

type NoteFormCallback = FormCallback[*models.Note]

func saveNoteFields(
	_instance *models.Note,
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
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(_instance.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(_instance.IsExpanded), formDiv)
		case "IsPartsNodeExpanded":
			FormDivBasicFieldToField(&(_instance.IsPartsNodeExpanded), formDiv)
		case "Parts":
			FormDivSliceOfPointersToField(_instance, "Parts", &(_instance.Parts), formDiv, probe)
		case "IsPortsNodeExpanded":
			FormDivBasicFieldToField(&(_instance.IsPortsNodeExpanded), formDiv)
		case "Ports":
			FormDivSliceOfPointersToField(_instance, "Ports", &(_instance.Ports), formDiv, probe)
		case "DiagramStructure:NotesWhoseNodeIsExpanded":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "NotesWhoseNodeIsExpanded", func(owner *models.DiagramStructure) *[]*models.Note { return &owner.NotesWhoseNodeIsExpanded })
		case "Library:RootNotes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "RootNotes", func(owner *models.Library) *[]*models.Note { return &owner.RootNotes })
		case "Library:NotesWhoseNodeIsExpanded":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "NotesWhoseNodeIsExpanded", func(owner *models.Library) *[]*models.Note { return &owner.NotesWhoseNodeIsExpanded })
		}
	}
}

func __gong__New__NotePartShapeFormCallback(
	_instance *models.NotePartShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (notepartshapeFormCallback *FormCallback[*models.NotePartShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveNotePartShapeFields,
	)
}

type NotePartShapeFormCallback = FormCallback[*models.NotePartShape]

func saveNotePartShapeFields(
	_instance *models.NotePartShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Note":
			FormDivSelectFieldToField(&(_instance.Note), probe.stageOfInterest, formDiv)
		case "Part":
			FormDivSelectFieldToField(&(_instance.Part), probe.stageOfInterest, formDiv)
		case "StartRatio":
			FormDivBasicFieldToField(&(_instance.StartRatio), formDiv)
		case "EndRatio":
			FormDivBasicFieldToField(&(_instance.EndRatio), formDiv)
		case "StartOrientation":
			FormDivEnumStringFieldToField(&(_instance.StartOrientation), formDiv)
		case "EndOrientation":
			FormDivEnumStringFieldToField(&(_instance.EndOrientation), formDiv)
		case "CornerOffsetRatio":
			FormDivBasicFieldToField(&(_instance.CornerOffsetRatio), formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&(_instance.IsHidden), formDiv)
		case "DiagramStructure:NotePartShapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "NotePartShapes", func(owner *models.DiagramStructure) *[]*models.NotePartShape { return &owner.NotePartShapes })
		}
	}
}

func __gong__New__NotePortShapeFormCallback(
	_instance *models.NotePortShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (noteportshapeFormCallback *FormCallback[*models.NotePortShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveNotePortShapeFields,
	)
}

type NotePortShapeFormCallback = FormCallback[*models.NotePortShape]

func saveNotePortShapeFields(
	_instance *models.NotePortShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Note":
			FormDivSelectFieldToField(&(_instance.Note), probe.stageOfInterest, formDiv)
		case "Port":
			FormDivSelectFieldToField(&(_instance.Port), probe.stageOfInterest, formDiv)
		case "StartRatio":
			FormDivBasicFieldToField(&(_instance.StartRatio), formDiv)
		case "EndRatio":
			FormDivBasicFieldToField(&(_instance.EndRatio), formDiv)
		case "StartOrientation":
			FormDivEnumStringFieldToField(&(_instance.StartOrientation), formDiv)
		case "EndOrientation":
			FormDivEnumStringFieldToField(&(_instance.EndOrientation), formDiv)
		case "CornerOffsetRatio":
			FormDivBasicFieldToField(&(_instance.CornerOffsetRatio), formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&(_instance.IsHidden), formDiv)
		case "DiagramStructure:NotePortShapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "NotePortShapes", func(owner *models.DiagramStructure) *[]*models.NotePortShape { return &owner.NotePortShapes })
		}
	}
}

func __gong__New__NoteShapeFormCallback(
	_instance *models.NoteShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (noteshapeFormCallback *FormCallback[*models.NoteShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveNoteShapeFields,
	)
}

type NoteShapeFormCallback = FormCallback[*models.NoteShape]

func saveNoteShapeFields(
	_instance *models.NoteShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Note":
			FormDivSelectFieldToField(&(_instance.Note), probe.stageOfInterest, formDiv)
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
		case "DiagramStructure:Note_Shapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Note_Shapes", func(owner *models.DiagramStructure) *[]*models.NoteShape { return &owner.Note_Shapes })
		}
	}
}

func __gong__New__PartFormCallback(
	_instance *models.Part,
	probe *Probe,
	formGroup *form.FormGroup,
) (partFormCallback *FormCallback[*models.Part]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		savePartFields,
	)
}

type PartFormCallback = FormCallback[*models.Part]

func savePartFields(
	_instance *models.Part,
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
		case "Ports":
			FormDivSliceOfPointersToField(_instance, "Ports", &(_instance.Ports), formDiv, probe)
		case "TypeOfPart":
			FormDivSelectFieldToField(&(_instance.TypeOfPart), probe.stageOfInterest, formDiv)
		case "IsPartNameNotSystemName":
			FormDivBasicFieldToField(&(_instance.IsPartNameNotSystemName), formDiv)
		case "IsControlFlowsNodeExpanded":
			FormDivBasicFieldToField(&(_instance.IsControlFlowsNodeExpanded), formDiv)
		case "ControlFlows":
			FormDivSliceOfPointersToField(_instance, "ControlFlows", &(_instance.ControlFlows), formDiv, probe)
		case "PortWhoseOutControlFlowsNodeIsExpanded":
			FormDivSliceOfPointersToField(_instance, "PortWhoseOutControlFlowsNodeIsExpanded", &(_instance.PortWhoseOutControlFlowsNodeIsExpanded), formDiv, probe)
		case "PortWhoseInControlFlowsNodeIsExpanded":
			FormDivSliceOfPointersToField(_instance, "PortWhoseInControlFlowsNodeIsExpanded", &(_instance.PortWhoseInControlFlowsNodeIsExpanded), formDiv, probe)
		case "IsDataFlowsNodeExpanded":
			FormDivBasicFieldToField(&(_instance.IsDataFlowsNodeExpanded), formDiv)
		case "PortWhoseOutDataFlowsNodeIsExpanded":
			FormDivSliceOfPointersToField(_instance, "PortWhoseOutDataFlowsNodeIsExpanded", &(_instance.PortWhoseOutDataFlowsNodeIsExpanded), formDiv, probe)
		case "PortWhoseInDataFlowsNodeIsExpanded":
			FormDivSliceOfPointersToField(_instance, "PortWhoseInDataFlowsNodeIsExpanded", &(_instance.PortWhoseInDataFlowsNodeIsExpanded), formDiv, probe)
		case "PartAnchoredPath":
			FormDivSliceOfPointersToField(_instance, "PartAnchoredPath", &(_instance.PartAnchoredPath), formDiv, probe)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(_instance.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(_instance.IsExpanded), formDiv)
		case "IsPortsNodeExpanded":
			FormDivBasicFieldToField(&(_instance.IsPortsNodeExpanded), formDiv)
		case "DiagramStructure:PartWhoseNodeIsExpanded":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "PartWhoseNodeIsExpanded", func(owner *models.DiagramStructure) *[]*models.Part { return &owner.PartWhoseNodeIsExpanded })
		case "DiagramStructure:ExternalPartWhoseNodeIsExpanded":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "ExternalPartWhoseNodeIsExpanded", func(owner *models.DiagramStructure) *[]*models.Part { return &owner.ExternalPartWhoseNodeIsExpanded })
		case "DiagramStructure:ExternalPartsWhoseOutDataFlowsNodeIsExpanded":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "ExternalPartsWhoseOutDataFlowsNodeIsExpanded", func(owner *models.DiagramStructure) *[]*models.Part { return &owner.ExternalPartsWhoseOutDataFlowsNodeIsExpanded })
		case "DiagramStructure:ExternalPartsWhoseInDataFlowsNodeIsExpanded":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "ExternalPartsWhoseInDataFlowsNodeIsExpanded", func(owner *models.DiagramStructure) *[]*models.Part { return &owner.ExternalPartsWhoseInDataFlowsNodeIsExpanded })
		case "Library:PartsWhoseNodeIsExpanded":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "PartsWhoseNodeIsExpanded", func(owner *models.Library) *[]*models.Part { return &owner.PartsWhoseNodeIsExpanded })
		case "Note:Parts":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Parts", func(owner *models.Note) *[]*models.Part { return &owner.Parts })
		case "SemanticTag:Parts":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Parts", func(owner *models.SemanticTag) *[]*models.Part { return &owner.Parts })
		case "System:Parts":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Parts", func(owner *models.System) *[]*models.Part { return &owner.Parts })
		case "System:PartWhoseNodeIsExpanded":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "PartWhoseNodeIsExpanded", func(owner *models.System) *[]*models.Part { return &owner.PartWhoseNodeIsExpanded })
		case "System:ExternalParts":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "ExternalParts", func(owner *models.System) *[]*models.Part { return &owner.ExternalParts })
		case "System:ExternalPartWhoseNodeIsExpanded":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "ExternalPartWhoseNodeIsExpanded", func(owner *models.System) *[]*models.Part { return &owner.ExternalPartWhoseNodeIsExpanded })
		}
	}
}

func __gong__New__PartAnchoredPathFormCallback(
	_instance *models.PartAnchoredPath,
	probe *Probe,
	formGroup *form.FormGroup,
) (partanchoredpathFormCallback *FormCallback[*models.PartAnchoredPath]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		savePartAnchoredPathFields,
	)
}

type PartAnchoredPathFormCallback = FormCallback[*models.PartAnchoredPath]

func savePartAnchoredPathFields(
	_instance *models.PartAnchoredPath,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Definition":
			FormDivBasicFieldToField(&(_instance.Definition), formDiv)
		case "X_Offset":
			FormDivBasicFieldToField(&(_instance.X_Offset), formDiv)
		case "Y_Offset":
			FormDivBasicFieldToField(&(_instance.Y_Offset), formDiv)
		case "RectAnchorType":
			FormDivEnumStringFieldToField(&(_instance.RectAnchorType), formDiv)
		case "ScalePropotionnally":
			FormDivBasicFieldToField(&(_instance.ScalePropotionnally), formDiv)
		case "AppliedScaling":
			FormDivBasicFieldToField(&(_instance.AppliedScaling), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(_instance.Color), formDiv)
		case "FillOpacity":
			FormDivBasicFieldToField(&(_instance.FillOpacity), formDiv)
		case "Stroke":
			FormDivBasicFieldToField(&(_instance.Stroke), formDiv)
		case "StrokeOpacity":
			FormDivBasicFieldToField(&(_instance.StrokeOpacity), formDiv)
		case "StrokeWidth":
			FormDivBasicFieldToField(&(_instance.StrokeWidth), formDiv)
		case "StrokeDashArray":
			FormDivBasicFieldToField(&(_instance.StrokeDashArray), formDiv)
		case "StrokeDashArrayWhenSelected":
			FormDivBasicFieldToField(&(_instance.StrokeDashArrayWhenSelected), formDiv)
		case "Transform":
			FormDivBasicFieldToField(&(_instance.Transform), formDiv)
		case "Part:PartAnchoredPath":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "PartAnchoredPath", func(owner *models.Part) *[]*models.PartAnchoredPath { return &owner.PartAnchoredPath })
		}
	}
}

func __gong__New__PartShapeFormCallback(
	_instance *models.PartShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (partshapeFormCallback *FormCallback[*models.PartShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		savePartShapeFields,
	)
}

type PartShapeFormCallback = FormCallback[*models.PartShape]

func savePartShapeFields(
	_instance *models.PartShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Part":
			FormDivSelectFieldToField(&(_instance.Part), probe.stageOfInterest, formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(_instance.IsExpanded), formDiv)
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
		case "DiagramStructure:Part_Shapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Part_Shapes", func(owner *models.DiagramStructure) *[]*models.PartShape { return &owner.Part_Shapes })
		}
	}
}

func __gong__New__PortFormCallback(
	_instance *models.Port,
	probe *Probe,
	formGroup *form.FormGroup,
) (portFormCallback *FormCallback[*models.Port]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		savePortFields,
	)
}

type PortFormCallback = FormCallback[*models.Port]

func savePortFields(
	_instance *models.Port,
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
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(_instance.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(_instance.IsExpanded), formDiv)
		case "DiagramStructure:PortsWhoseNodeIsExpanded":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "PortsWhoseNodeIsExpanded", func(owner *models.DiagramStructure) *[]*models.Port { return &owner.PortsWhoseNodeIsExpanded })
		case "Note:Ports":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Ports", func(owner *models.Note) *[]*models.Port { return &owner.Ports })
		case "Part:Ports":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Ports", func(owner *models.Part) *[]*models.Port { return &owner.Ports })
		case "Part:PortWhoseOutControlFlowsNodeIsExpanded":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "PortWhoseOutControlFlowsNodeIsExpanded", func(owner *models.Part) *[]*models.Port { return &owner.PortWhoseOutControlFlowsNodeIsExpanded })
		case "Part:PortWhoseInControlFlowsNodeIsExpanded":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "PortWhoseInControlFlowsNodeIsExpanded", func(owner *models.Part) *[]*models.Port { return &owner.PortWhoseInControlFlowsNodeIsExpanded })
		case "Part:PortWhoseOutDataFlowsNodeIsExpanded":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "PortWhoseOutDataFlowsNodeIsExpanded", func(owner *models.Part) *[]*models.Port { return &owner.PortWhoseOutDataFlowsNodeIsExpanded })
		case "Part:PortWhoseInDataFlowsNodeIsExpanded":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "PortWhoseInDataFlowsNodeIsExpanded", func(owner *models.Part) *[]*models.Port { return &owner.PortWhoseInDataFlowsNodeIsExpanded })
		}
	}
}

func __gong__New__PortShapeFormCallback(
	_instance *models.PortShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (portshapeFormCallback *FormCallback[*models.PortShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		savePortShapeFields,
	)
}

type PortShapeFormCallback = FormCallback[*models.PortShape]

func savePortShapeFields(
	_instance *models.PortShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Port":
			FormDivSelectFieldToField(&(_instance.Port), probe.stageOfInterest, formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(_instance.IsExpanded), formDiv)
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
		case "DiagramStructure:Port_Shapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Port_Shapes", func(owner *models.DiagramStructure) *[]*models.PortShape { return &owner.Port_Shapes })
		}
	}
}

func __gong__New__ResourceFormCallback(
	_instance *models.Resource,
	probe *Probe,
	formGroup *form.FormGroup,
) (resourceFormCallback *FormCallback[*models.Resource]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveResourceFields,
	)
}

type ResourceFormCallback = FormCallback[*models.Resource]

func saveResourceFields(
	_instance *models.Resource,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Acronym":
			FormDivBasicFieldToField(&(_instance.Acronym), formDiv)
		case "Description":
			FormDivBasicFieldToField(&(_instance.Description), formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(_instance.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(_instance.IsExpanded), formDiv)
		case "SVG_Path":
			FormDivBasicFieldToField(&(_instance.SVG_Path), formDiv)
		case "InverseAppliedScaling":
			FormDivBasicFieldToField(&(_instance.InverseAppliedScaling), formDiv)
		case "DiagramStructure:AllocatedResourcesWhoseNodeIsExpanded":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "AllocatedResourcesWhoseNodeIsExpanded", func(owner *models.DiagramStructure) *[]*models.Resource { return &owner.AllocatedResourcesWhoseNodeIsExpanded })
		case "Library:RootResources":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "RootResources", func(owner *models.Library) *[]*models.Resource { return &owner.RootResources })
		case "Library:ResourcesWhoseNodeIsExpanded":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "ResourcesWhoseNodeIsExpanded", func(owner *models.Library) *[]*models.Resource { return &owner.ResourcesWhoseNodeIsExpanded })
		}
	}
}

func __gong__New__SemanticTagFormCallback(
	_instance *models.SemanticTag,
	probe *Probe,
	formGroup *form.FormGroup,
) (semantictagFormCallback *FormCallback[*models.SemanticTag]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveSemanticTagFields,
	)
}

type SemanticTagFormCallback = FormCallback[*models.SemanticTag]

func saveSemanticTagFields(
	_instance *models.SemanticTag,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Parts":
			FormDivSliceOfPointersToField(_instance, "Parts", &(_instance.Parts), formDiv, probe)
		case "LayerDefinition:Query":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Query", func(owner *models.LayerDefinition) *[]*models.SemanticTag { return &owner.Query })
		}
	}
}

func __gong__New__SystemFormCallback(
	_instance *models.System,
	probe *Probe,
	formGroup *form.FormGroup,
) (systemFormCallback *FormCallback[*models.System]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveSystemFields,
	)
}

type SystemFormCallback = FormCallback[*models.System]

func saveSystemFields(
	_instance *models.System,
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
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(_instance.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(_instance.IsExpanded), formDiv)
		case "SVG_Path":
			FormDivBasicFieldToField(&(_instance.SVG_Path), formDiv)
		case "InverseAppliedScaling":
			FormDivBasicFieldToField(&(_instance.InverseAppliedScaling), formDiv)
		case "DiagramStructures":
			FormDivSliceOfPointersToField(_instance, "DiagramStructures", &(_instance.DiagramStructures), formDiv, probe)
		case "DiagramStructureWhoseNodeIsExpanded":
			FormDivSliceOfPointersToField(_instance, "DiagramStructureWhoseNodeIsExpanded", &(_instance.DiagramStructureWhoseNodeIsExpanded), formDiv, probe)
		case "IsSubSystemNodeExpanded":
			FormDivBasicFieldToField(&(_instance.IsSubSystemNodeExpanded), formDiv)
		case "SubSystemes":
			FormDivSliceOfPointersToField(_instance, "SubSystemes", &(_instance.SubSystemes), formDiv, probe)
		case "Parts":
			FormDivSliceOfPointersToField(_instance, "Parts", &(_instance.Parts), formDiv, probe)
		case "PartWhoseNodeIsExpanded":
			FormDivSliceOfPointersToField(_instance, "PartWhoseNodeIsExpanded", &(_instance.PartWhoseNodeIsExpanded), formDiv, probe)
		case "DataFlows":
			FormDivSliceOfPointersToField(_instance, "DataFlows", &(_instance.DataFlows), formDiv, probe)
		case "IsDataFlowsNodeExpanded":
			FormDivBasicFieldToField(&(_instance.IsDataFlowsNodeExpanded), formDiv)
		case "ExternalParts":
			FormDivSliceOfPointersToField(_instance, "ExternalParts", &(_instance.ExternalParts), formDiv, probe)
		case "ExternalPartWhoseNodeIsExpanded":
			FormDivSliceOfPointersToField(_instance, "ExternalPartWhoseNodeIsExpanded", &(_instance.ExternalPartWhoseNodeIsExpanded), formDiv, probe)
		case "DiagramStructure:SystemsWhoseNodeIsExpanded":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "SystemsWhoseNodeIsExpanded", func(owner *models.DiagramStructure) *[]*models.System { return &owner.SystemsWhoseNodeIsExpanded })
		case "DiagramStructure:AllocatedSystemesWhoseNodeIsExpanded":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "AllocatedSystemesWhoseNodeIsExpanded", func(owner *models.DiagramStructure) *[]*models.System { return &owner.AllocatedSystemesWhoseNodeIsExpanded })
		case "Library:RootSystemes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "RootSystemes", func(owner *models.Library) *[]*models.System { return &owner.RootSystemes })
		case "Library:SystemsWhoseNodeIsExpanded":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "SystemsWhoseNodeIsExpanded", func(owner *models.Library) *[]*models.System { return &owner.SystemsWhoseNodeIsExpanded })
		case "System:SubSystemes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "SubSystemes", func(owner *models.System) *[]*models.System { return &owner.SubSystemes })
		}
	}
}

func __gong__New__SystemShapeFormCallback(
	_instance *models.SystemShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (systemshapeFormCallback *FormCallback[*models.SystemShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveSystemShapeFields,
	)
}

type SystemShapeFormCallback = FormCallback[*models.SystemShape]

func saveSystemShapeFields(
	_instance *models.SystemShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "System":
			FormDivSelectFieldToField(&(_instance.System), probe.stageOfInterest, formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(_instance.IsExpanded), formDiv)
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
		case "DiagramStructure:System_Shapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "System_Shapes", func(owner *models.DiagramStructure) *[]*models.SystemShape { return &owner.System_Shapes })
		}
	}
}

