// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/dsm/capture/go/models"
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
func __gong__New__AnalysisNeedFormCallback(
	_instance *models.AnalysisNeed,
	probe *Probe,
	formGroup *form.FormGroup,
) (analysisneedFormCallback *FormCallback[*models.AnalysisNeed]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveAnalysisNeedFields,
	)
}

type AnalysisNeedFormCallback = FormCallback[*models.AnalysisNeed]

func saveAnalysisNeedFields(
	_instance *models.AnalysisNeed,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(_instance.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(_instance.IsExpanded), formDiv)
		case "Library:AnalysisNeeds":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "AnalysisNeeds", func(owner *models.Library) *[]*models.AnalysisNeed { return &owner.AnalysisNeeds })
		}
	}
}

func __gong__New__ConceptFormCallback(
	_instance *models.Concept,
	probe *Probe,
	formGroup *form.FormGroup,
) (conceptFormCallback *FormCallback[*models.Concept]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveConceptFields,
	)
}

type ConceptFormCallback = FormCallback[*models.Concept]

func saveConceptFields(
	_instance *models.Concept,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(_instance.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(_instance.IsExpanded), formDiv)
		case "Tools":
			FormDivSliceOfPointersToField(_instance, "Tools", &(_instance.Tools), formDiv, probe)
		case "Deliverable:Concepts":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Concepts", func(owner *models.Deliverable) *[]*models.Concept { return &owner.Concepts })
		case "Diagram:ConceptsWhoseNodeIsExpanded":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "ConceptsWhoseNodeIsExpanded", func(owner *models.Diagram) *[]*models.Concept { return &owner.ConceptsWhoseNodeIsExpanded })
		case "Diagram:ConceptsWhoseDeliverablesNodeIsExpanded":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "ConceptsWhoseDeliverablesNodeIsExpanded", func(owner *models.Diagram) *[]*models.Concept { return &owner.ConceptsWhoseDeliverablesNodeIsExpanded })
		case "Library:RootConcepts":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "RootConcepts", func(owner *models.Library) *[]*models.Concept { return &owner.RootConcepts })
		case "Requirement:Concepts":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Concepts", func(owner *models.Requirement) *[]*models.Concept { return &owner.Concepts })
		}
	}
}

func __gong__New__ConceptShapeFormCallback(
	_instance *models.ConceptShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (conceptshapeFormCallback *FormCallback[*models.ConceptShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveConceptShapeFields,
	)
}

type ConceptShapeFormCallback = FormCallback[*models.ConceptShape]

func saveConceptShapeFields(
	_instance *models.ConceptShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Concept":
			FormDivSelectFieldToField(&(_instance.Concept), probe.stageOfInterest, formDiv)
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
		case "Diagram:Concept_Shapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Concept_Shapes", func(owner *models.Diagram) *[]*models.ConceptShape { return &owner.Concept_Shapes })
		}
	}
}

func __gong__New__ConcernFormCallback(
	_instance *models.Concern,
	probe *Probe,
	formGroup *form.FormGroup,
) (concernFormCallback *FormCallback[*models.Concern]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveConcernFields,
	)
}

type ConcernFormCallback = FormCallback[*models.Concern]

func saveConcernFields(
	_instance *models.Concern,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "IDAirbus":
			FormDivBasicFieldToField(&(_instance.IDAirbus), formDiv)
		case "Priority":
			FormDivEnumStringFieldToField(&(_instance.Priority), formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(_instance.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(_instance.IsExpanded), formDiv)
		case "Description":
			FormDivBasicFieldToField(&(_instance.Description), formDiv)
		case "SubConcerns":
			FormDivSliceOfPointersToField(_instance, "SubConcerns", &(_instance.SubConcerns), formDiv, probe)
		case "Inputs":
			FormDivSliceOfPointersToField(_instance, "Inputs", &(_instance.Inputs), formDiv, probe)
		case "IsInputsNodeExpanded":
			FormDivBasicFieldToField(&(_instance.IsInputsNodeExpanded), formDiv)
		case "Outputs":
			FormDivSliceOfPointersToField(_instance, "Outputs", &(_instance.Outputs), formDiv, probe)
		case "IsOutputsNodeExpanded":
			FormDivBasicFieldToField(&(_instance.IsOutputsNodeExpanded), formDiv)
		case "IsWithCompletion":
			FormDivBasicFieldToField(&(_instance.IsWithCompletion), formDiv)
		case "Completion":
			FormDivEnumStringFieldToField(&(_instance.Completion), formDiv)
		case "Requirements":
			FormDivSliceOfPointersToField(_instance, "Requirements", &(_instance.Requirements), formDiv, probe)
		case "Concern:SubConcerns":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "SubConcerns", func(owner *models.Concern) *[]*models.Concern { return &owner.SubConcerns })
		case "Diagram:ConcernsWhoseRequirementsNodeIsExpanded":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "ConcernsWhoseRequirementsNodeIsExpanded", func(owner *models.Diagram) *[]*models.Concern { return &owner.ConcernsWhoseRequirementsNodeIsExpanded })
		case "Diagram:ConcernsWhoseNodeIsExpanded":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "ConcernsWhoseNodeIsExpanded", func(owner *models.Diagram) *[]*models.Concern { return &owner.ConcernsWhoseNodeIsExpanded })
		case "Diagram:ConcernsWhoseInputNodeIsExpanded":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "ConcernsWhoseInputNodeIsExpanded", func(owner *models.Diagram) *[]*models.Concern { return &owner.ConcernsWhoseInputNodeIsExpanded })
		case "Diagram:ConcernsWhoseStakeholderNodeIsExpanded":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "ConcernsWhoseStakeholderNodeIsExpanded", func(owner *models.Diagram) *[]*models.Concern { return &owner.ConcernsWhoseStakeholderNodeIsExpanded })
		case "Diagram:ConcernssWhoseOutputNodeIsExpanded":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "ConcernssWhoseOutputNodeIsExpanded", func(owner *models.Diagram) *[]*models.Concern { return &owner.ConcernssWhoseOutputNodeIsExpanded })
		case "Library:RootConcerns":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "RootConcerns", func(owner *models.Library) *[]*models.Concern { return &owner.RootConcerns })
		case "Note:Tasks":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Tasks", func(owner *models.Note) *[]*models.Concern { return &owner.Tasks })
		case "Stakeholder:Concerns":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Concerns", func(owner *models.Stakeholder) *[]*models.Concern { return &owner.Concerns })
		}
	}
}

func __gong__New__ConcernCompositionShapeFormCallback(
	_instance *models.ConcernCompositionShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (concerncompositionshapeFormCallback *FormCallback[*models.ConcernCompositionShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveConcernCompositionShapeFields,
	)
}

type ConcernCompositionShapeFormCallback = FormCallback[*models.ConcernCompositionShape]

func saveConcernCompositionShapeFields(
	_instance *models.ConcernCompositionShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Concern":
			FormDivSelectFieldToField(&(_instance.Concern), probe.stageOfInterest, formDiv)
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
		case "ControlPointShapes":
			FormDivSliceOfPointersToField(_instance, "ControlPointShapes", &(_instance.ControlPointShapes), formDiv, probe)
		case "Diagram:ConcernComposition_Shapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "ConcernComposition_Shapes", func(owner *models.Diagram) *[]*models.ConcernCompositionShape { return &owner.ConcernComposition_Shapes })
		}
	}
}

func __gong__New__ConcernInputShapeFormCallback(
	_instance *models.ConcernInputShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (concerninputshapeFormCallback *FormCallback[*models.ConcernInputShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveConcernInputShapeFields,
	)
}

type ConcernInputShapeFormCallback = FormCallback[*models.ConcernInputShape]

func saveConcernInputShapeFields(
	_instance *models.ConcernInputShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Deliverable":
			FormDivSelectFieldToField(&(_instance.Deliverable), probe.stageOfInterest, formDiv)
		case "Concern":
			FormDivSelectFieldToField(&(_instance.Concern), probe.stageOfInterest, formDiv)
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
		case "ControlPointShapes":
			FormDivSliceOfPointersToField(_instance, "ControlPointShapes", &(_instance.ControlPointShapes), formDiv, probe)
		case "Diagram:ConcernInputShapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "ConcernInputShapes", func(owner *models.Diagram) *[]*models.ConcernInputShape { return &owner.ConcernInputShapes })
		}
	}
}

func __gong__New__ConcernOutputShapeFormCallback(
	_instance *models.ConcernOutputShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (concernoutputshapeFormCallback *FormCallback[*models.ConcernOutputShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveConcernOutputShapeFields,
	)
}

type ConcernOutputShapeFormCallback = FormCallback[*models.ConcernOutputShape]

func saveConcernOutputShapeFields(
	_instance *models.ConcernOutputShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Concern":
			FormDivSelectFieldToField(&(_instance.Concern), probe.stageOfInterest, formDiv)
		case "Deliverable":
			FormDivSelectFieldToField(&(_instance.Deliverable), probe.stageOfInterest, formDiv)
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
		case "ControlPointShapes":
			FormDivSliceOfPointersToField(_instance, "ControlPointShapes", &(_instance.ControlPointShapes), formDiv, probe)
		case "Diagram:ConcernOutputShapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "ConcernOutputShapes", func(owner *models.Diagram) *[]*models.ConcernOutputShape { return &owner.ConcernOutputShapes })
		}
	}
}

func __gong__New__ConcernShapeFormCallback(
	_instance *models.ConcernShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (concernshapeFormCallback *FormCallback[*models.ConcernShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveConcernShapeFields,
	)
}

type ConcernShapeFormCallback = FormCallback[*models.ConcernShape]

func saveConcernShapeFields(
	_instance *models.ConcernShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Concern":
			FormDivSelectFieldToField(&(_instance.Concern), probe.stageOfInterest, formDiv)
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
		case "Diagram:Concern_Shapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Concern_Shapes", func(owner *models.Diagram) *[]*models.ConcernShape { return &owner.Concern_Shapes })
		}
	}
}

func __gong__New__ControlPointShapeFormCallback(
	_instance *models.ControlPointShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (controlpointshapeFormCallback *FormCallback[*models.ControlPointShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveControlPointShapeFields,
	)
}

type ControlPointShapeFormCallback = FormCallback[*models.ControlPointShape]

func saveControlPointShapeFields(
	_instance *models.ControlPointShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "X_Relative":
			FormDivBasicFieldToField(&(_instance.X_Relative), formDiv)
		case "Y_Relative":
			FormDivBasicFieldToField(&(_instance.Y_Relative), formDiv)
		case "IsStartShapeTheClosestShape":
			FormDivBasicFieldToField(&(_instance.IsStartShapeTheClosestShape), formDiv)
		case "ConcernCompositionShape:ControlPointShapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "ControlPointShapes", func(owner *models.ConcernCompositionShape) *[]*models.ControlPointShape { return &owner.ControlPointShapes })
		case "ConcernInputShape:ControlPointShapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "ControlPointShapes", func(owner *models.ConcernInputShape) *[]*models.ControlPointShape { return &owner.ControlPointShapes })
		case "ConcernOutputShape:ControlPointShapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "ControlPointShapes", func(owner *models.ConcernOutputShape) *[]*models.ControlPointShape { return &owner.ControlPointShapes })
		case "DeliverableCompositionShape:ControlPointShapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "ControlPointShapes", func(owner *models.DeliverableCompositionShape) *[]*models.ControlPointShape { return &owner.ControlPointShapes })
		case "DeliverableConceptShape:ControlPointShapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "ControlPointShapes", func(owner *models.DeliverableConceptShape) *[]*models.ControlPointShape { return &owner.ControlPointShapes })
		case "NoteDeliverableShape:ControlPointShapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "ControlPointShapes", func(owner *models.NoteDeliverableShape) *[]*models.ControlPointShape { return &owner.ControlPointShapes })
		case "NoteStakeholderShape:ControlPointShapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "ControlPointShapes", func(owner *models.NoteStakeholderShape) *[]*models.ControlPointShape { return &owner.ControlPointShapes })
		case "NoteTaskShape:ControlPointShapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "ControlPointShapes", func(owner *models.NoteTaskShape) *[]*models.ControlPointShape { return &owner.ControlPointShapes })
		case "StakeholderCompositionShape:ControlPointShapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "ControlPointShapes", func(owner *models.StakeholderCompositionShape) *[]*models.ControlPointShape { return &owner.ControlPointShapes })
		case "StakeholderConcernShape:ControlPointShapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "ControlPointShapes", func(owner *models.StakeholderConcernShape) *[]*models.ControlPointShape { return &owner.ControlPointShapes })
		}
	}
}

func __gong__New__DeliverableFormCallback(
	_instance *models.Deliverable,
	probe *Probe,
	formGroup *form.FormGroup,
) (deliverableFormCallback *FormCallback[*models.Deliverable]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveDeliverableFields,
	)
}

type DeliverableFormCallback = FormCallback[*models.Deliverable]

func saveDeliverableFields(
	_instance *models.Deliverable,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(_instance.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(_instance.IsExpanded), formDiv)
		case "Description":
			FormDivBasicFieldToField(&(_instance.Description), formDiv)
		case "SubDeliverables":
			FormDivSliceOfPointersToField(_instance, "SubDeliverables", &(_instance.SubDeliverables), formDiv, probe)
		case "IsProducersNodeExpanded":
			FormDivBasicFieldToField(&(_instance.IsProducersNodeExpanded), formDiv)
		case "IsConsumersNodeExpanded":
			FormDivBasicFieldToField(&(_instance.IsConsumersNodeExpanded), formDiv)
		case "Concepts":
			FormDivSliceOfPointersToField(_instance, "Concepts", &(_instance.Concepts), formDiv, probe)
		case "Concern:Inputs":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Inputs", func(owner *models.Concern) *[]*models.Deliverable { return &owner.Inputs })
		case "Concern:Outputs":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Outputs", func(owner *models.Concern) *[]*models.Deliverable { return &owner.Outputs })
		case "Deliverable:SubDeliverables":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "SubDeliverables", func(owner *models.Deliverable) *[]*models.Deliverable { return &owner.SubDeliverables })
		case "Diagram:DeliverablesWhoseNodeIsExpanded":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "DeliverablesWhoseNodeIsExpanded", func(owner *models.Diagram) *[]*models.Deliverable { return &owner.DeliverablesWhoseNodeIsExpanded })
		case "Diagram:DeliverablesWhoseConceptsNodeIsExpanded":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "DeliverablesWhoseConceptsNodeIsExpanded", func(owner *models.Diagram) *[]*models.Deliverable { return &owner.DeliverablesWhoseConceptsNodeIsExpanded })
		case "Library:RootDeliverables":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "RootDeliverables", func(owner *models.Library) *[]*models.Deliverable { return &owner.RootDeliverables })
		case "Note:Deliverables":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Deliverables", func(owner *models.Note) *[]*models.Deliverable { return &owner.Deliverables })
		}
	}
}

func __gong__New__DeliverableCompositionShapeFormCallback(
	_instance *models.DeliverableCompositionShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (deliverablecompositionshapeFormCallback *FormCallback[*models.DeliverableCompositionShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveDeliverableCompositionShapeFields,
	)
}

type DeliverableCompositionShapeFormCallback = FormCallback[*models.DeliverableCompositionShape]

func saveDeliverableCompositionShapeFields(
	_instance *models.DeliverableCompositionShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Deliverable":
			FormDivSelectFieldToField(&(_instance.Deliverable), probe.stageOfInterest, formDiv)
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
		case "ControlPointShapes":
			FormDivSliceOfPointersToField(_instance, "ControlPointShapes", &(_instance.ControlPointShapes), formDiv, probe)
		case "Diagram:DeliverableComposition_Shapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "DeliverableComposition_Shapes", func(owner *models.Diagram) *[]*models.DeliverableCompositionShape { return &owner.DeliverableComposition_Shapes })
		}
	}
}

func __gong__New__DeliverableConceptShapeFormCallback(
	_instance *models.DeliverableConceptShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (deliverableconceptshapeFormCallback *FormCallback[*models.DeliverableConceptShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveDeliverableConceptShapeFields,
	)
}

type DeliverableConceptShapeFormCallback = FormCallback[*models.DeliverableConceptShape]

func saveDeliverableConceptShapeFields(
	_instance *models.DeliverableConceptShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Deliverable":
			FormDivSelectFieldToField(&(_instance.Deliverable), probe.stageOfInterest, formDiv)
		case "Concept":
			FormDivSelectFieldToField(&(_instance.Concept), probe.stageOfInterest, formDiv)
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
		case "ControlPointShapes":
			FormDivSliceOfPointersToField(_instance, "ControlPointShapes", &(_instance.ControlPointShapes), formDiv, probe)
		case "Diagram:DeliverableConceptShapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "DeliverableConceptShapes", func(owner *models.Diagram) *[]*models.DeliverableConceptShape { return &owner.DeliverableConceptShapes })
		}
	}
}

func __gong__New__DeliverableShapeFormCallback(
	_instance *models.DeliverableShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (deliverableshapeFormCallback *FormCallback[*models.DeliverableShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveDeliverableShapeFields,
	)
}

type DeliverableShapeFormCallback = FormCallback[*models.DeliverableShape]

func saveDeliverableShapeFields(
	_instance *models.DeliverableShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Deliverable":
			FormDivSelectFieldToField(&(_instance.Deliverable), probe.stageOfInterest, formDiv)
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
		case "Diagram:Deliverable_Shapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Deliverable_Shapes", func(owner *models.Diagram) *[]*models.DeliverableShape { return &owner.Deliverable_Shapes })
		}
	}
}

func __gong__New__DiagramFormCallback(
	_instance *models.Diagram,
	probe *Probe,
	formGroup *form.FormGroup,
) (diagramFormCallback *FormCallback[*models.Diagram]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveDiagramFields,
	)
}

type DiagramFormCallback = FormCallback[*models.Diagram]

func saveDiagramFields(
	_instance *models.Diagram,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(_instance.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(_instance.IsExpanded), formDiv)
		case "IsChecked":
			FormDivBasicFieldToField(&(_instance.IsChecked), formDiv)
		case "IsEditable_":
			FormDivBasicFieldToField(&(_instance.IsEditable_), formDiv)
		case "ShowPrefix":
			FormDivBasicFieldToField(&(_instance.ShowPrefix), formDiv)
		case "DefaultBoxWidth":
			FormDivBasicFieldToField(&(_instance.DefaultBoxWidth), formDiv)
		case "DefaultBoxHeigth":
			FormDivBasicFieldToField(&(_instance.DefaultBoxHeigth), formDiv)
		case "Width":
			FormDivBasicFieldToField(&(_instance.Width), formDiv)
		case "Height":
			FormDivBasicFieldToField(&(_instance.Height), formDiv)
		case "ConcernsWhoseRequirementsNodeIsExpanded":
			FormDivSliceOfPointersToField(_instance, "ConcernsWhoseRequirementsNodeIsExpanded", &(_instance.ConcernsWhoseRequirementsNodeIsExpanded), formDiv, probe)
		case "IsRequirementsNodeExpanded":
			FormDivBasicFieldToField(&(_instance.IsRequirementsNodeExpanded), formDiv)
		case "IsConceptsNodeExpanded":
			FormDivBasicFieldToField(&(_instance.IsConceptsNodeExpanded), formDiv)
		case "Deliverable_Shapes":
			FormDivSliceOfPointersToField(_instance, "Deliverable_Shapes", &(_instance.Deliverable_Shapes), formDiv, probe)
		case "DeliverablesWhoseNodeIsExpanded":
			FormDivSliceOfPointersToField(_instance, "DeliverablesWhoseNodeIsExpanded", &(_instance.DeliverablesWhoseNodeIsExpanded), formDiv, probe)
		case "DeliverablesWhoseConceptsNodeIsExpanded":
			FormDivSliceOfPointersToField(_instance, "DeliverablesWhoseConceptsNodeIsExpanded", &(_instance.DeliverablesWhoseConceptsNodeIsExpanded), formDiv, probe)
		case "IsPBSNodeExpanded":
			FormDivBasicFieldToField(&(_instance.IsPBSNodeExpanded), formDiv)
		case "DeliverableComposition_Shapes":
			FormDivSliceOfPointersToField(_instance, "DeliverableComposition_Shapes", &(_instance.DeliverableComposition_Shapes), formDiv, probe)
		case "IsConcernsNodeExpanded":
			FormDivBasicFieldToField(&(_instance.IsConcernsNodeExpanded), formDiv)
		case "Concern_Shapes":
			FormDivSliceOfPointersToField(_instance, "Concern_Shapes", &(_instance.Concern_Shapes), formDiv, probe)
		case "ConcernsWhoseNodeIsExpanded":
			FormDivSliceOfPointersToField(_instance, "ConcernsWhoseNodeIsExpanded", &(_instance.ConcernsWhoseNodeIsExpanded), formDiv, probe)
		case "ConcernsWhoseInputNodeIsExpanded":
			FormDivSliceOfPointersToField(_instance, "ConcernsWhoseInputNodeIsExpanded", &(_instance.ConcernsWhoseInputNodeIsExpanded), formDiv, probe)
		case "ConcernsWhoseStakeholderNodeIsExpanded":
			FormDivSliceOfPointersToField(_instance, "ConcernsWhoseStakeholderNodeIsExpanded", &(_instance.ConcernsWhoseStakeholderNodeIsExpanded), formDiv, probe)
		case "ConcernssWhoseOutputNodeIsExpanded":
			FormDivSliceOfPointersToField(_instance, "ConcernssWhoseOutputNodeIsExpanded", &(_instance.ConcernssWhoseOutputNodeIsExpanded), formDiv, probe)
		case "ConcernComposition_Shapes":
			FormDivSliceOfPointersToField(_instance, "ConcernComposition_Shapes", &(_instance.ConcernComposition_Shapes), formDiv, probe)
		case "ConcernInputShapes":
			FormDivSliceOfPointersToField(_instance, "ConcernInputShapes", &(_instance.ConcernInputShapes), formDiv, probe)
		case "ConcernOutputShapes":
			FormDivSliceOfPointersToField(_instance, "ConcernOutputShapes", &(_instance.ConcernOutputShapes), formDiv, probe)
		case "Note_Shapes":
			FormDivSliceOfPointersToField(_instance, "Note_Shapes", &(_instance.Note_Shapes), formDiv, probe)
		case "NotesWhoseNodeIsExpanded":
			FormDivSliceOfPointersToField(_instance, "NotesWhoseNodeIsExpanded", &(_instance.NotesWhoseNodeIsExpanded), formDiv, probe)
		case "IsNotesNodeExpanded":
			FormDivBasicFieldToField(&(_instance.IsNotesNodeExpanded), formDiv)
		case "NoteDeliverableShapes":
			FormDivSliceOfPointersToField(_instance, "NoteDeliverableShapes", &(_instance.NoteDeliverableShapes), formDiv, probe)
		case "NoteTaskShapes":
			FormDivSliceOfPointersToField(_instance, "NoteTaskShapes", &(_instance.NoteTaskShapes), formDiv, probe)
		case "NoteResourceShapes":
			FormDivSliceOfPointersToField(_instance, "NoteResourceShapes", &(_instance.NoteResourceShapes), formDiv, probe)
		case "Stakeholder_Shapes":
			FormDivSliceOfPointersToField(_instance, "Stakeholder_Shapes", &(_instance.Stakeholder_Shapes), formDiv, probe)
		case "ResourcesWhoseNodeIsExpanded":
			FormDivSliceOfPointersToField(_instance, "ResourcesWhoseNodeIsExpanded", &(_instance.ResourcesWhoseNodeIsExpanded), formDiv, probe)
		case "IsStakeholdersNodeExpanded":
			FormDivBasicFieldToField(&(_instance.IsStakeholdersNodeExpanded), formDiv)
		case "ResourceComposition_Shapes":
			FormDivSliceOfPointersToField(_instance, "ResourceComposition_Shapes", &(_instance.ResourceComposition_Shapes), formDiv, probe)
		case "StakeholderConcernShapes":
			FormDivSliceOfPointersToField(_instance, "StakeholderConcernShapes", &(_instance.StakeholderConcernShapes), formDiv, probe)
		case "Requirement_Shapes":
			FormDivSliceOfPointersToField(_instance, "Requirement_Shapes", &(_instance.Requirement_Shapes), formDiv, probe)
		case "RequirementsWhoseNodeIsExpanded":
			FormDivSliceOfPointersToField(_instance, "RequirementsWhoseNodeIsExpanded", &(_instance.RequirementsWhoseNodeIsExpanded), formDiv, probe)
		case "Concept_Shapes":
			FormDivSliceOfPointersToField(_instance, "Concept_Shapes", &(_instance.Concept_Shapes), formDiv, probe)
		case "ConceptsWhoseNodeIsExpanded":
			FormDivSliceOfPointersToField(_instance, "ConceptsWhoseNodeIsExpanded", &(_instance.ConceptsWhoseNodeIsExpanded), formDiv, probe)
		case "ConceptsWhoseDeliverablesNodeIsExpanded":
			FormDivSliceOfPointersToField(_instance, "ConceptsWhoseDeliverablesNodeIsExpanded", &(_instance.ConceptsWhoseDeliverablesNodeIsExpanded), formDiv, probe)
		case "DeliverableConceptShapes":
			FormDivSliceOfPointersToField(_instance, "DeliverableConceptShapes", &(_instance.DeliverableConceptShapes), formDiv, probe)
		case "Diagram_Shapes":
			FormDivSliceOfPointersToField(_instance, "Diagram_Shapes", &(_instance.Diagram_Shapes), formDiv, probe)
		case "IsDiagramsNodeExpanded":
			FormDivBasicFieldToField(&(_instance.IsDiagramsNodeExpanded), formDiv)
		case "DiagramsWhoseNodeIsExpanded":
			FormDivSliceOfPointersToField(_instance, "DiagramsWhoseNodeIsExpanded", &(_instance.DiagramsWhoseNodeIsExpanded), formDiv, probe)
		case "Diagram:DiagramsWhoseNodeIsExpanded":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "DiagramsWhoseNodeIsExpanded", func(owner *models.Diagram) *[]*models.Diagram { return &owner.DiagramsWhoseNodeIsExpanded })
		case "Library:Diagrams":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Diagrams", func(owner *models.Library) *[]*models.Diagram { return &owner.Diagrams })
		}
	}
}

func __gong__New__DiagramShapeFormCallback(
	_instance *models.DiagramShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (diagramshapeFormCallback *FormCallback[*models.DiagramShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveDiagramShapeFields,
	)
}

type DiagramShapeFormCallback = FormCallback[*models.DiagramShape]

func saveDiagramShapeFields(
	_instance *models.DiagramShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Diagram":
			FormDivSelectFieldToField(&(_instance.Diagram), probe.stageOfInterest, formDiv)
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
		case "Diagram:Diagram_Shapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Diagram_Shapes", func(owner *models.Diagram) *[]*models.DiagramShape { return &owner.Diagram_Shapes })
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
		case "IsRootLibrary":
			FormDivBasicFieldToField(&(_instance.IsRootLibrary), formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(_instance.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(_instance.IsExpanded), formDiv)
		case "RootDeliverables":
			FormDivSliceOfPointersToField(_instance, "RootDeliverables", &(_instance.RootDeliverables), formDiv, probe)
		case "RootConcerns":
			FormDivSliceOfPointersToField(_instance, "RootConcerns", &(_instance.RootConcerns), formDiv, probe)
		case "RootStakeholders":
			FormDivSliceOfPointersToField(_instance, "RootStakeholders", &(_instance.RootStakeholders), formDiv, probe)
		case "RootRequirements":
			FormDivSliceOfPointersToField(_instance, "RootRequirements", &(_instance.RootRequirements), formDiv, probe)
		case "RootConcepts":
			FormDivSliceOfPointersToField(_instance, "RootConcepts", &(_instance.RootConcepts), formDiv, probe)
		case "AnalysisNeeds":
			FormDivSliceOfPointersToField(_instance, "AnalysisNeeds", &(_instance.AnalysisNeeds), formDiv, probe)
		case "Notes":
			FormDivSliceOfPointersToField(_instance, "Notes", &(_instance.Notes), formDiv, probe)
		case "Diagrams":
			FormDivSliceOfPointersToField(_instance, "Diagrams", &(_instance.Diagrams), formDiv, probe)
		case "SubLibraries":
			FormDivSliceOfPointersToField(_instance, "SubLibraries", &(_instance.SubLibraries), formDiv, probe)
		case "NbPixPerCharacter":
			FormDivBasicFieldToField(&(_instance.NbPixPerCharacter), formDiv)
		case "Library:SubLibraries":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "SubLibraries", func(owner *models.Library) *[]*models.Library { return &owner.SubLibraries })
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
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(_instance.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(_instance.IsExpanded), formDiv)
		case "Deliverables":
			FormDivSliceOfPointersToField(_instance, "Deliverables", &(_instance.Deliverables), formDiv, probe)
		case "Tasks":
			FormDivSliceOfPointersToField(_instance, "Tasks", &(_instance.Tasks), formDiv, probe)
		case "Resources":
			FormDivSliceOfPointersToField(_instance, "Resources", &(_instance.Resources), formDiv, probe)
		case "Diagram:NotesWhoseNodeIsExpanded":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "NotesWhoseNodeIsExpanded", func(owner *models.Diagram) *[]*models.Note { return &owner.NotesWhoseNodeIsExpanded })
		case "Library:Notes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Notes", func(owner *models.Library) *[]*models.Note { return &owner.Notes })
		}
	}
}

func __gong__New__NoteDeliverableShapeFormCallback(
	_instance *models.NoteDeliverableShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (notedeliverableshapeFormCallback *FormCallback[*models.NoteDeliverableShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveNoteDeliverableShapeFields,
	)
}

type NoteDeliverableShapeFormCallback = FormCallback[*models.NoteDeliverableShape]

func saveNoteDeliverableShapeFields(
	_instance *models.NoteDeliverableShape,
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
		case "Deliverable":
			FormDivSelectFieldToField(&(_instance.Deliverable), probe.stageOfInterest, formDiv)
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
		case "ControlPointShapes":
			FormDivSliceOfPointersToField(_instance, "ControlPointShapes", &(_instance.ControlPointShapes), formDiv, probe)
		case "Diagram:NoteDeliverableShapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "NoteDeliverableShapes", func(owner *models.Diagram) *[]*models.NoteDeliverableShape { return &owner.NoteDeliverableShapes })
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
		case "Diagram:Note_Shapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Note_Shapes", func(owner *models.Diagram) *[]*models.NoteShape { return &owner.Note_Shapes })
		}
	}
}

func __gong__New__NoteStakeholderShapeFormCallback(
	_instance *models.NoteStakeholderShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (notestakeholdershapeFormCallback *FormCallback[*models.NoteStakeholderShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveNoteStakeholderShapeFields,
	)
}

type NoteStakeholderShapeFormCallback = FormCallback[*models.NoteStakeholderShape]

func saveNoteStakeholderShapeFields(
	_instance *models.NoteStakeholderShape,
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
		case "Stakeholder":
			FormDivSelectFieldToField(&(_instance.Stakeholder), probe.stageOfInterest, formDiv)
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
		case "ControlPointShapes":
			FormDivSliceOfPointersToField(_instance, "ControlPointShapes", &(_instance.ControlPointShapes), formDiv, probe)
		case "Diagram:NoteResourceShapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "NoteResourceShapes", func(owner *models.Diagram) *[]*models.NoteStakeholderShape { return &owner.NoteResourceShapes })
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
		case "ControlPointShapes":
			FormDivSliceOfPointersToField(_instance, "ControlPointShapes", &(_instance.ControlPointShapes), formDiv, probe)
		case "Diagram:NoteTaskShapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "NoteTaskShapes", func(owner *models.Diagram) *[]*models.NoteTaskShape { return &owner.NoteTaskShapes })
		}
	}
}

func __gong__New__RequirementFormCallback(
	_instance *models.Requirement,
	probe *Probe,
	formGroup *form.FormGroup,
) (requirementFormCallback *FormCallback[*models.Requirement]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveRequirementFields,
	)
}

type RequirementFormCallback = FormCallback[*models.Requirement]

func saveRequirementFields(
	_instance *models.Requirement,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(_instance.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(_instance.IsExpanded), formDiv)
		case "SupportLevels":
			FormDivSliceOfPointersToField(_instance, "SupportLevels", &(_instance.SupportLevels), formDiv, probe)
		case "Concepts":
			FormDivSliceOfPointersToField(_instance, "Concepts", &(_instance.Concepts), formDiv, probe)
		case "Concern:Requirements":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Requirements", func(owner *models.Concern) *[]*models.Requirement { return &owner.Requirements })
		case "Diagram:RequirementsWhoseNodeIsExpanded":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "RequirementsWhoseNodeIsExpanded", func(owner *models.Diagram) *[]*models.Requirement { return &owner.RequirementsWhoseNodeIsExpanded })
		case "Library:RootRequirements":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "RootRequirements", func(owner *models.Library) *[]*models.Requirement { return &owner.RootRequirements })
		}
	}
}

func __gong__New__RequirementShapeFormCallback(
	_instance *models.RequirementShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (requirementshapeFormCallback *FormCallback[*models.RequirementShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveRequirementShapeFields,
	)
}

type RequirementShapeFormCallback = FormCallback[*models.RequirementShape]

func saveRequirementShapeFields(
	_instance *models.RequirementShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Requirement":
			FormDivSelectFieldToField(&(_instance.Requirement), probe.stageOfInterest, formDiv)
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
		case "Diagram:Requirement_Shapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Requirement_Shapes", func(owner *models.Diagram) *[]*models.RequirementShape { return &owner.Requirement_Shapes })
		}
	}
}

func __gong__New__StakeholderFormCallback(
	_instance *models.Stakeholder,
	probe *Probe,
	formGroup *form.FormGroup,
) (stakeholderFormCallback *FormCallback[*models.Stakeholder]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveStakeholderFields,
	)
}

type StakeholderFormCallback = FormCallback[*models.Stakeholder]

func saveStakeholderFields(
	_instance *models.Stakeholder,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "IDAirbus":
			FormDivBasicFieldToField(&(_instance.IDAirbus), formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(_instance.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(_instance.IsExpanded), formDiv)
		case "Description":
			FormDivBasicFieldToField(&(_instance.Description), formDiv)
		case "Concerns":
			FormDivSliceOfPointersToField(_instance, "Concerns", &(_instance.Concerns), formDiv, probe)
		case "SubStakeholders":
			FormDivSliceOfPointersToField(_instance, "SubStakeholders", &(_instance.SubStakeholders), formDiv, probe)
		case "Diagram:ResourcesWhoseNodeIsExpanded":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "ResourcesWhoseNodeIsExpanded", func(owner *models.Diagram) *[]*models.Stakeholder { return &owner.ResourcesWhoseNodeIsExpanded })
		case "Library:RootStakeholders":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "RootStakeholders", func(owner *models.Library) *[]*models.Stakeholder { return &owner.RootStakeholders })
		case "Note:Resources":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Resources", func(owner *models.Note) *[]*models.Stakeholder { return &owner.Resources })
		case "Stakeholder:SubStakeholders":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "SubStakeholders", func(owner *models.Stakeholder) *[]*models.Stakeholder { return &owner.SubStakeholders })
		}
	}
}

func __gong__New__StakeholderCompositionShapeFormCallback(
	_instance *models.StakeholderCompositionShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (stakeholdercompositionshapeFormCallback *FormCallback[*models.StakeholderCompositionShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveStakeholderCompositionShapeFields,
	)
}

type StakeholderCompositionShapeFormCallback = FormCallback[*models.StakeholderCompositionShape]

func saveStakeholderCompositionShapeFields(
	_instance *models.StakeholderCompositionShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Stakeholder":
			FormDivSelectFieldToField(&(_instance.Stakeholder), probe.stageOfInterest, formDiv)
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
		case "ControlPointShapes":
			FormDivSliceOfPointersToField(_instance, "ControlPointShapes", &(_instance.ControlPointShapes), formDiv, probe)
		case "Diagram:ResourceComposition_Shapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "ResourceComposition_Shapes", func(owner *models.Diagram) *[]*models.StakeholderCompositionShape { return &owner.ResourceComposition_Shapes })
		}
	}
}

func __gong__New__StakeholderConcernShapeFormCallback(
	_instance *models.StakeholderConcernShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (stakeholderconcernshapeFormCallback *FormCallback[*models.StakeholderConcernShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveStakeholderConcernShapeFields,
	)
}

type StakeholderConcernShapeFormCallback = FormCallback[*models.StakeholderConcernShape]

func saveStakeholderConcernShapeFields(
	_instance *models.StakeholderConcernShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Stakeholder":
			FormDivSelectFieldToField(&(_instance.Stakeholder), probe.stageOfInterest, formDiv)
		case "Concern":
			FormDivSelectFieldToField(&(_instance.Concern), probe.stageOfInterest, formDiv)
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
		case "ControlPointShapes":
			FormDivSliceOfPointersToField(_instance, "ControlPointShapes", &(_instance.ControlPointShapes), formDiv, probe)
		case "Diagram:StakeholderConcernShapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "StakeholderConcernShapes", func(owner *models.Diagram) *[]*models.StakeholderConcernShape { return &owner.StakeholderConcernShapes })
		}
	}
}

func __gong__New__StakeholderShapeFormCallback(
	_instance *models.StakeholderShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (stakeholdershapeFormCallback *FormCallback[*models.StakeholderShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveStakeholderShapeFields,
	)
}

type StakeholderShapeFormCallback = FormCallback[*models.StakeholderShape]

func saveStakeholderShapeFields(
	_instance *models.StakeholderShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Stakeholder":
			FormDivSelectFieldToField(&(_instance.Stakeholder), probe.stageOfInterest, formDiv)
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
		case "Diagram:Stakeholder_Shapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Stakeholder_Shapes", func(owner *models.Diagram) *[]*models.StakeholderShape { return &owner.Stakeholder_Shapes })
		}
	}
}

func __gong__New__SupportLevelFormCallback(
	_instance *models.SupportLevel,
	probe *Probe,
	formGroup *form.FormGroup,
) (supportlevelFormCallback *FormCallback[*models.SupportLevel]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveSupportLevelFields,
	)
}

type SupportLevelFormCallback = FormCallback[*models.SupportLevel]

func saveSupportLevelFields(
	_instance *models.SupportLevel,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Tool":
			FormDivSelectFieldToField(&(_instance.Tool), probe.stageOfInterest, formDiv)
		case "Requirement:SupportLevels":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "SupportLevels", func(owner *models.Requirement) *[]*models.SupportLevel { return &owner.SupportLevels })
		}
	}
}

func __gong__New__ToolFormCallback(
	_instance *models.Tool,
	probe *Probe,
	formGroup *form.FormGroup,
) (toolFormCallback *FormCallback[*models.Tool]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveToolFields,
	)
}

type ToolFormCallback = FormCallback[*models.Tool]

func saveToolFields(
	_instance *models.Tool,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Concept:Tools":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Tools", func(owner *models.Concept) *[]*models.Tool { return &owner.Tools })
		}
	}
}

