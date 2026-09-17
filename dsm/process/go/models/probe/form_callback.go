// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/dsm/process/go/models"
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
func __gong__New__AllocatedProcessShapeFormCallback(
	_instance *models.AllocatedProcessShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (allocatedprocessshapeFormCallback *FormCallback[*models.AllocatedProcessShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveAllocatedProcessShapeFields,
	)
}

type AllocatedProcessShapeFormCallback = FormCallback[*models.AllocatedProcessShape]

func saveAllocatedProcessShapeFields(
	_instance *models.AllocatedProcessShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Participant":
			FormDivSelectFieldToField(&(_instance.Participant), probe.stageOfInterest, formDiv)
		case "Process":
			FormDivSelectFieldToField(&(_instance.Process), probe.stageOfInterest, formDiv)
		case "DiagramProcess:AllocatedProcessShapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "AllocatedProcessShapes", func(owner *models.DiagramProcess) *[]*models.AllocatedProcessShape { return &owner.AllocatedProcessShapes })
		}
	}
}

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
		case "Participant":
			FormDivSelectFieldToField(&(_instance.Participant), probe.stageOfInterest, formDiv)
		case "Resource":
			FormDivSelectFieldToField(&(_instance.Resource), probe.stageOfInterest, formDiv)
		case "DiagramProcess:AllocatedResourceShapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "AllocatedResourceShapes", func(owner *models.DiagramProcess) *[]*models.AllocatedResourceShape { return &owner.AllocatedResourceShapes })
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
		case "DiagramProcess:ControlFlowsWhoseNodeIsExpanded":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "ControlFlowsWhoseNodeIsExpanded", func(owner *models.DiagramProcess) *[]*models.ControlFlow { return &owner.ControlFlowsWhoseNodeIsExpanded })
		case "Participant:ControlFlows":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "ControlFlows", func(owner *models.Participant) *[]*models.ControlFlow { return &owner.ControlFlows })
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
		case "DiagramProcess:ControlFlow_Shapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "ControlFlow_Shapes", func(owner *models.DiagramProcess) *[]*models.ControlFlowShape { return &owner.ControlFlow_Shapes })
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
		case "DiagramProcess:DatasWhoseNodeIsExpanded":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "DatasWhoseNodeIsExpanded", func(owner *models.DiagramProcess) *[]*models.Data { return &owner.DatasWhoseNodeIsExpanded })
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
		case "Datas":
			FormDivSliceOfPointersToField(_instance, "Datas", &(_instance.Datas), formDiv, probe)
		case "Description":
			FormDivBasicFieldToField(&(_instance.Description), formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(_instance.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(_instance.IsExpanded), formDiv)
		case "Type":
			FormDivEnumStringFieldToField(&(_instance.Type), formDiv)
		case "StartTask":
			FormDivSelectFieldToField(&(_instance.StartTask), probe.stageOfInterest, formDiv)
		case "EndTask":
			FormDivSelectFieldToField(&(_instance.EndTask), probe.stageOfInterest, formDiv)
		case "StartExternalParticipant":
			FormDivSelectFieldToField(&(_instance.StartExternalParticipant), probe.stageOfInterest, formDiv)
		case "EndExternalParticipant":
			FormDivSelectFieldToField(&(_instance.EndExternalParticipant), probe.stageOfInterest, formDiv)
		case "IsDatasNodeExpanded":
			FormDivBasicFieldToField(&(_instance.IsDatasNodeExpanded), formDiv)
		case "DiagramProcess:DataFlowsWhoseNodeIsExpanded":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "DataFlowsWhoseNodeIsExpanded", func(owner *models.DiagramProcess) *[]*models.DataFlow { return &owner.DataFlowsWhoseNodeIsExpanded })
		case "DiagramProcess:DataFlowsWhoseDataNodeIsExpanded":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "DataFlowsWhoseDataNodeIsExpanded", func(owner *models.DiagramProcess) *[]*models.DataFlow { return &owner.DataFlowsWhoseDataNodeIsExpanded })
		case "Library:RootDataFlows":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "RootDataFlows", func(owner *models.Library) *[]*models.DataFlow { return &owner.RootDataFlows })
		case "Library:DataFlowsWhoseNodeIsExpanded":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "DataFlowsWhoseNodeIsExpanded", func(owner *models.Library) *[]*models.DataFlow { return &owner.DataFlowsWhoseNodeIsExpanded })
		case "Process:DataFlows":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "DataFlows", func(owner *models.Process) *[]*models.DataFlow { return &owner.DataFlows })
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
		case "DiagramProcess:DataFlow_Shapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "DataFlow_Shapes", func(owner *models.DiagramProcess) *[]*models.DataFlowShape { return &owner.DataFlow_Shapes })
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
		case "DiagramProcess:Data_Shapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Data_Shapes", func(owner *models.DiagramProcess) *[]*models.DataShape { return &owner.Data_Shapes })
		}
	}
}

func __gong__New__DiagramProcessFormCallback(
	_instance *models.DiagramProcess,
	probe *Probe,
	formGroup *form.FormGroup,
) (diagramprocessFormCallback *FormCallback[*models.DiagramProcess]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveDiagramProcessFields,
	)
}

type DiagramProcessFormCallback = FormCallback[*models.DiagramProcess]

func saveDiagramProcessFields(
	_instance *models.DiagramProcess,
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
		case "Width":
			FormDivBasicFieldToField(&(_instance.Width), formDiv)
		case "Height":
			FormDivBasicFieldToField(&(_instance.Height), formDiv)
		case "Process_Shapes":
			FormDivSliceOfPointersToField(_instance, "Process_Shapes", &(_instance.Process_Shapes), formDiv, probe)
		case "IsProcesssNodeExpanded":
			FormDivBasicFieldToField(&(_instance.IsProcesssNodeExpanded), formDiv)
		case "ProcesssWhoseNodeIsExpanded":
			FormDivSliceOfPointersToField(_instance, "ProcesssWhoseNodeIsExpanded", &(_instance.ProcesssWhoseNodeIsExpanded), formDiv, probe)
		case "Participant_Shapes":
			FormDivSliceOfPointersToField(_instance, "Participant_Shapes", &(_instance.Participant_Shapes), formDiv, probe)
		case "IsParticipantsNodeExpanded":
			FormDivBasicFieldToField(&(_instance.IsParticipantsNodeExpanded), formDiv)
		case "ParticipantWhoseNodeIsExpanded":
			FormDivSliceOfPointersToField(_instance, "ParticipantWhoseNodeIsExpanded", &(_instance.ParticipantWhoseNodeIsExpanded), formDiv, probe)
		case "ExternalParticipant_Shapes":
			FormDivSliceOfPointersToField(_instance, "ExternalParticipant_Shapes", &(_instance.ExternalParticipant_Shapes), formDiv, probe)
		case "IsExternalParticipantsNodeExpanded":
			FormDivBasicFieldToField(&(_instance.IsExternalParticipantsNodeExpanded), formDiv)
		case "ExternalParticipantWhoseNodeIsExpanded":
			FormDivSliceOfPointersToField(_instance, "ExternalParticipantWhoseNodeIsExpanded", &(_instance.ExternalParticipantWhoseNodeIsExpanded), formDiv, probe)
		case "ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded":
			FormDivSliceOfPointersToField(_instance, "ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded", &(_instance.ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded), formDiv, probe)
		case "ExternalParticipantsWhoseInDataFlowsNodeIsExpanded":
			FormDivSliceOfPointersToField(_instance, "ExternalParticipantsWhoseInDataFlowsNodeIsExpanded", &(_instance.ExternalParticipantsWhoseInDataFlowsNodeIsExpanded), formDiv, probe)
		case "TasksWhoseNodeIsExpanded":
			FormDivSliceOfPointersToField(_instance, "TasksWhoseNodeIsExpanded", &(_instance.TasksWhoseNodeIsExpanded), formDiv, probe)
		case "Task_Shapes":
			FormDivSliceOfPointersToField(_instance, "Task_Shapes", &(_instance.Task_Shapes), formDiv, probe)
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
		case "AllocatedProcessesWhoseNodeIsExpanded":
			FormDivSliceOfPointersToField(_instance, "AllocatedProcessesWhoseNodeIsExpanded", &(_instance.AllocatedProcessesWhoseNodeIsExpanded), formDiv, probe)
		case "AllocatedProcessShapes":
			FormDivSliceOfPointersToField(_instance, "AllocatedProcessShapes", &(_instance.AllocatedProcessShapes), formDiv, probe)
		case "Note_Shapes":
			FormDivSliceOfPointersToField(_instance, "Note_Shapes", &(_instance.Note_Shapes), formDiv, probe)
		case "NotesWhoseNodeIsExpanded":
			FormDivSliceOfPointersToField(_instance, "NotesWhoseNodeIsExpanded", &(_instance.NotesWhoseNodeIsExpanded), formDiv, probe)
		case "IsNotesNodeExpanded":
			FormDivBasicFieldToField(&(_instance.IsNotesNodeExpanded), formDiv)
		case "NoteTaskShapes":
			FormDivSliceOfPointersToField(_instance, "NoteTaskShapes", &(_instance.NoteTaskShapes), formDiv, probe)
		case "Process:DiagramProcesss":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "DiagramProcesss", func(owner *models.Process) *[]*models.DiagramProcess { return &owner.DiagramProcesss })
		case "Process:DiagramProcessWhoseNodeIsExpanded":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "DiagramProcessWhoseNodeIsExpanded", func(owner *models.Process) *[]*models.DiagramProcess { return &owner.DiagramProcessWhoseNodeIsExpanded })
		}
	}
}

func __gong__New__ExternalParticipantShapeFormCallback(
	_instance *models.ExternalParticipantShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (externalparticipantshapeFormCallback *FormCallback[*models.ExternalParticipantShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveExternalParticipantShapeFields,
	)
}

type ExternalParticipantShapeFormCallback = FormCallback[*models.ExternalParticipantShape]

func saveExternalParticipantShapeFields(
	_instance *models.ExternalParticipantShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Participant":
			FormDivSelectFieldToField(&(_instance.Participant), probe.stageOfInterest, formDiv)
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
		case "DiagramProcess:ExternalParticipant_Shapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "ExternalParticipant_Shapes", func(owner *models.DiagramProcess) *[]*models.ExternalParticipantShape { return &owner.ExternalParticipant_Shapes })
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
		case "RootProcesses":
			FormDivSliceOfPointersToField(_instance, "RootProcesses", &(_instance.RootProcesses), formDiv, probe)
		case "IsProcessesNodeExpanded":
			FormDivBasicFieldToField(&(_instance.IsProcessesNodeExpanded), formDiv)
		case "ProcesssWhoseNodeIsExpanded":
			FormDivSliceOfPointersToField(_instance, "ProcesssWhoseNodeIsExpanded", &(_instance.ProcesssWhoseNodeIsExpanded), formDiv, probe)
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
		case "ParticipantsWhoseNodeIsExpanded":
			FormDivSliceOfPointersToField(_instance, "ParticipantsWhoseNodeIsExpanded", &(_instance.ParticipantsWhoseNodeIsExpanded), formDiv, probe)
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
		case "IsTasksNodeExpanded":
			FormDivBasicFieldToField(&(_instance.IsTasksNodeExpanded), formDiv)
		case "Tasks":
			FormDivSliceOfPointersToField(_instance, "Tasks", &(_instance.Tasks), formDiv, probe)
		case "DiagramProcess:NotesWhoseNodeIsExpanded":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "NotesWhoseNodeIsExpanded", func(owner *models.DiagramProcess) *[]*models.Note { return &owner.NotesWhoseNodeIsExpanded })
		case "Library:RootNotes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "RootNotes", func(owner *models.Library) *[]*models.Note { return &owner.RootNotes })
		case "Library:NotesWhoseNodeIsExpanded":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "NotesWhoseNodeIsExpanded", func(owner *models.Library) *[]*models.Note { return &owner.NotesWhoseNodeIsExpanded })
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
		case "DiagramProcess:Note_Shapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Note_Shapes", func(owner *models.DiagramProcess) *[]*models.NoteShape { return &owner.Note_Shapes })
		}
	}
}

func __gong__New__NoteTaskShapeFormCallback(
	_instance *models.NoteTaskShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (notetaskshapeFormCallback *FormCallback[*models.NoteTaskShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveNoteTaskShapeFields,
	)
}

type NoteTaskShapeFormCallback = FormCallback[*models.NoteTaskShape]

func saveNoteTaskShapeFields(
	_instance *models.NoteTaskShape,
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
		case "Task":
			FormDivSelectFieldToField(&(_instance.Task), probe.stageOfInterest, formDiv)
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
		case "DiagramProcess:NoteTaskShapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "NoteTaskShapes", func(owner *models.DiagramProcess) *[]*models.NoteTaskShape { return &owner.NoteTaskShapes })
		}
	}
}

func __gong__New__ParticipantFormCallback(
	_instance *models.Participant,
	probe *Probe,
	formGroup *form.FormGroup,
) (participantFormCallback *FormCallback[*models.Participant]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveParticipantFields,
	)
}

type ParticipantFormCallback = FormCallback[*models.Participant]

func saveParticipantFields(
	_instance *models.Participant,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "IsProcessResource":
			FormDivBasicFieldToField(&(_instance.IsProcessResource), formDiv)
		case "Description":
			FormDivBasicFieldToField(&(_instance.Description), formDiv)
		case "Resources":
			FormDivSliceOfPointersToField(_instance, "Resources", &(_instance.Resources), formDiv, probe)
		case "IsResourcesNodeExpanded":
			FormDivBasicFieldToField(&(_instance.IsResourcesNodeExpanded), formDiv)
		case "Processes":
			FormDivSliceOfPointersToField(_instance, "Processes", &(_instance.Processes), formDiv, probe)
		case "IsProcessesNodeExpanded":
			FormDivBasicFieldToField(&(_instance.IsProcessesNodeExpanded), formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(_instance.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(_instance.IsExpanded), formDiv)
		case "IsTasksNodeExpanded":
			FormDivBasicFieldToField(&(_instance.IsTasksNodeExpanded), formDiv)
		case "Tasks":
			FormDivSliceOfPointersToField(_instance, "Tasks", &(_instance.Tasks), formDiv, probe)
		case "IsControlFlowsNodeExpanded":
			FormDivBasicFieldToField(&(_instance.IsControlFlowsNodeExpanded), formDiv)
		case "ControlFlows":
			FormDivSliceOfPointersToField(_instance, "ControlFlows", &(_instance.ControlFlows), formDiv, probe)
		case "TaskWhoseOutControlFlowsNodeIsExpanded":
			FormDivSliceOfPointersToField(_instance, "TaskWhoseOutControlFlowsNodeIsExpanded", &(_instance.TaskWhoseOutControlFlowsNodeIsExpanded), formDiv, probe)
		case "TaskWhoseInControlFlowsNodeIsExpanded":
			FormDivSliceOfPointersToField(_instance, "TaskWhoseInControlFlowsNodeIsExpanded", &(_instance.TaskWhoseInControlFlowsNodeIsExpanded), formDiv, probe)
		case "IsDataFlowsNodeExpanded":
			FormDivBasicFieldToField(&(_instance.IsDataFlowsNodeExpanded), formDiv)
		case "TaskWhoseOutDataFlowsNodeIsExpanded":
			FormDivSliceOfPointersToField(_instance, "TaskWhoseOutDataFlowsNodeIsExpanded", &(_instance.TaskWhoseOutDataFlowsNodeIsExpanded), formDiv, probe)
		case "TaskWhoseInDataFlowsNodeIsExpanded":
			FormDivSliceOfPointersToField(_instance, "TaskWhoseInDataFlowsNodeIsExpanded", &(_instance.TaskWhoseInDataFlowsNodeIsExpanded), formDiv, probe)
		case "DiagramProcess:ParticipantWhoseNodeIsExpanded":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "ParticipantWhoseNodeIsExpanded", func(owner *models.DiagramProcess) *[]*models.Participant { return &owner.ParticipantWhoseNodeIsExpanded })
		case "DiagramProcess:ExternalParticipantWhoseNodeIsExpanded":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "ExternalParticipantWhoseNodeIsExpanded", func(owner *models.DiagramProcess) *[]*models.Participant { return &owner.ExternalParticipantWhoseNodeIsExpanded })
		case "DiagramProcess:ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded", func(owner *models.DiagramProcess) *[]*models.Participant { return &owner.ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded })
		case "DiagramProcess:ExternalParticipantsWhoseInDataFlowsNodeIsExpanded":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "ExternalParticipantsWhoseInDataFlowsNodeIsExpanded", func(owner *models.DiagramProcess) *[]*models.Participant { return &owner.ExternalParticipantsWhoseInDataFlowsNodeIsExpanded })
		case "Library:ParticipantsWhoseNodeIsExpanded":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "ParticipantsWhoseNodeIsExpanded", func(owner *models.Library) *[]*models.Participant { return &owner.ParticipantsWhoseNodeIsExpanded })
		case "Process:Participants":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Participants", func(owner *models.Process) *[]*models.Participant { return &owner.Participants })
		case "Process:ParticipantWhoseNodeIsExpanded":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "ParticipantWhoseNodeIsExpanded", func(owner *models.Process) *[]*models.Participant { return &owner.ParticipantWhoseNodeIsExpanded })
		case "Process:ExternalParticipants":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "ExternalParticipants", func(owner *models.Process) *[]*models.Participant { return &owner.ExternalParticipants })
		case "Process:ExternalParticipantWhoseNodeIsExpanded":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "ExternalParticipantWhoseNodeIsExpanded", func(owner *models.Process) *[]*models.Participant { return &owner.ExternalParticipantWhoseNodeIsExpanded })
		}
	}
}

func __gong__New__ParticipantShapeFormCallback(
	_instance *models.ParticipantShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (participantshapeFormCallback *FormCallback[*models.ParticipantShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveParticipantShapeFields,
	)
}

type ParticipantShapeFormCallback = FormCallback[*models.ParticipantShape]

func saveParticipantShapeFields(
	_instance *models.ParticipantShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Participant":
			FormDivSelectFieldToField(&(_instance.Participant), probe.stageOfInterest, formDiv)
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
		case "WidthWeight":
			FormDivBasicFieldToField(&(_instance.WidthWeight), formDiv)
		case "DiagramProcess:Participant_Shapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Participant_Shapes", func(owner *models.DiagramProcess) *[]*models.ParticipantShape { return &owner.Participant_Shapes })
		}
	}
}

func __gong__New__ProcessFormCallback(
	_instance *models.Process,
	probe *Probe,
	formGroup *form.FormGroup,
) (processFormCallback *FormCallback[*models.Process]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveProcessFields,
	)
}

type ProcessFormCallback = FormCallback[*models.Process]

func saveProcessFields(
	_instance *models.Process,
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
		case "DiagramProcesss":
			FormDivSliceOfPointersToField(_instance, "DiagramProcesss", &(_instance.DiagramProcesss), formDiv, probe)
		case "DiagramProcessWhoseNodeIsExpanded":
			FormDivSliceOfPointersToField(_instance, "DiagramProcessWhoseNodeIsExpanded", &(_instance.DiagramProcessWhoseNodeIsExpanded), formDiv, probe)
		case "IsSubProcessNodeExpanded":
			FormDivBasicFieldToField(&(_instance.IsSubProcessNodeExpanded), formDiv)
		case "SubProcesses":
			FormDivSliceOfPointersToField(_instance, "SubProcesses", &(_instance.SubProcesses), formDiv, probe)
		case "Participants":
			FormDivSliceOfPointersToField(_instance, "Participants", &(_instance.Participants), formDiv, probe)
		case "ParticipantWhoseNodeIsExpanded":
			FormDivSliceOfPointersToField(_instance, "ParticipantWhoseNodeIsExpanded", &(_instance.ParticipantWhoseNodeIsExpanded), formDiv, probe)
		case "DataFlows":
			FormDivSliceOfPointersToField(_instance, "DataFlows", &(_instance.DataFlows), formDiv, probe)
		case "IsDataFlowsNodeExpanded":
			FormDivBasicFieldToField(&(_instance.IsDataFlowsNodeExpanded), formDiv)
		case "ExternalParticipants":
			FormDivSliceOfPointersToField(_instance, "ExternalParticipants", &(_instance.ExternalParticipants), formDiv, probe)
		case "ExternalParticipantWhoseNodeIsExpanded":
			FormDivSliceOfPointersToField(_instance, "ExternalParticipantWhoseNodeIsExpanded", &(_instance.ExternalParticipantWhoseNodeIsExpanded), formDiv, probe)
		case "DiagramProcess:ProcesssWhoseNodeIsExpanded":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "ProcesssWhoseNodeIsExpanded", func(owner *models.DiagramProcess) *[]*models.Process { return &owner.ProcesssWhoseNodeIsExpanded })
		case "DiagramProcess:AllocatedProcessesWhoseNodeIsExpanded":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "AllocatedProcessesWhoseNodeIsExpanded", func(owner *models.DiagramProcess) *[]*models.Process { return &owner.AllocatedProcessesWhoseNodeIsExpanded })
		case "Library:RootProcesses":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "RootProcesses", func(owner *models.Library) *[]*models.Process { return &owner.RootProcesses })
		case "Library:ProcesssWhoseNodeIsExpanded":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "ProcesssWhoseNodeIsExpanded", func(owner *models.Library) *[]*models.Process { return &owner.ProcesssWhoseNodeIsExpanded })
		case "Participant:Processes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Processes", func(owner *models.Participant) *[]*models.Process { return &owner.Processes })
		case "Process:SubProcesses":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "SubProcesses", func(owner *models.Process) *[]*models.Process { return &owner.SubProcesses })
		}
	}
}

func __gong__New__ProcessShapeFormCallback(
	_instance *models.ProcessShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (processshapeFormCallback *FormCallback[*models.ProcessShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveProcessShapeFields,
	)
}

type ProcessShapeFormCallback = FormCallback[*models.ProcessShape]

func saveProcessShapeFields(
	_instance *models.ProcessShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Process":
			FormDivSelectFieldToField(&(_instance.Process), probe.stageOfInterest, formDiv)
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
		case "DiagramProcess:Process_Shapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Process_Shapes", func(owner *models.DiagramProcess) *[]*models.ProcessShape { return &owner.Process_Shapes })
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
		case "DiagramProcess:AllocatedResourcesWhoseNodeIsExpanded":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "AllocatedResourcesWhoseNodeIsExpanded", func(owner *models.DiagramProcess) *[]*models.Resource { return &owner.AllocatedResourcesWhoseNodeIsExpanded })
		case "Library:RootResources":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "RootResources", func(owner *models.Library) *[]*models.Resource { return &owner.RootResources })
		case "Library:ResourcesWhoseNodeIsExpanded":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "ResourcesWhoseNodeIsExpanded", func(owner *models.Library) *[]*models.Resource { return &owner.ResourcesWhoseNodeIsExpanded })
		case "Participant:Resources":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Resources", func(owner *models.Participant) *[]*models.Resource { return &owner.Resources })
		}
	}
}

func __gong__New__TaskFormCallback(
	_instance *models.Task,
	probe *Probe,
	formGroup *form.FormGroup,
) (taskFormCallback *FormCallback[*models.Task]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveTaskFields,
	)
}

type TaskFormCallback = FormCallback[*models.Task]

func saveTaskFields(
	_instance *models.Task,
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
		case "IsStartTask":
			FormDivBasicFieldToField(&(_instance.IsStartTask), formDiv)
		case "IsEndTask":
			FormDivBasicFieldToField(&(_instance.IsEndTask), formDiv)
		case "Type":
			FormDivSelectFieldToField(&(_instance.Type), probe.stageOfInterest, formDiv)
		case "IsTaskNameNotProcessName":
			FormDivBasicFieldToField(&(_instance.IsTaskNameNotProcessName), formDiv)
		case "DiagramProcess:TasksWhoseNodeIsExpanded":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "TasksWhoseNodeIsExpanded", func(owner *models.DiagramProcess) *[]*models.Task { return &owner.TasksWhoseNodeIsExpanded })
		case "Note:Tasks":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Tasks", func(owner *models.Note) *[]*models.Task { return &owner.Tasks })
		case "Participant:Tasks":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Tasks", func(owner *models.Participant) *[]*models.Task { return &owner.Tasks })
		case "Participant:TaskWhoseOutControlFlowsNodeIsExpanded":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "TaskWhoseOutControlFlowsNodeIsExpanded", func(owner *models.Participant) *[]*models.Task { return &owner.TaskWhoseOutControlFlowsNodeIsExpanded })
		case "Participant:TaskWhoseInControlFlowsNodeIsExpanded":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "TaskWhoseInControlFlowsNodeIsExpanded", func(owner *models.Participant) *[]*models.Task { return &owner.TaskWhoseInControlFlowsNodeIsExpanded })
		case "Participant:TaskWhoseOutDataFlowsNodeIsExpanded":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "TaskWhoseOutDataFlowsNodeIsExpanded", func(owner *models.Participant) *[]*models.Task { return &owner.TaskWhoseOutDataFlowsNodeIsExpanded })
		case "Participant:TaskWhoseInDataFlowsNodeIsExpanded":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "TaskWhoseInDataFlowsNodeIsExpanded", func(owner *models.Participant) *[]*models.Task { return &owner.TaskWhoseInDataFlowsNodeIsExpanded })
		}
	}
}

func __gong__New__TaskShapeFormCallback(
	_instance *models.TaskShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (taskshapeFormCallback *FormCallback[*models.TaskShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveTaskShapeFields,
	)
}

type TaskShapeFormCallback = FormCallback[*models.TaskShape]

func saveTaskShapeFields(
	_instance *models.TaskShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Task":
			FormDivSelectFieldToField(&(_instance.Task), probe.stageOfInterest, formDiv)
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
		case "DiagramProcess:Task_Shapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Task_Shapes", func(owner *models.DiagramProcess) *[]*models.TaskShape { return &owner.Task_Shapes })
		}
	}
}

