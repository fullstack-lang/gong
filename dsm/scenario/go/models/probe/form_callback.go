// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/dsm/scenario/go/models"
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
func __gong__New__ActorStateFormCallback(
	_instance *models.ActorState,
	probe *Probe,
	formGroup *form.FormGroup,
) (actorstateFormCallback *FormCallback[*models.ActorState]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveActorStateFields,
	)
}

type ActorStateFormCallback = FormCallback[*models.ActorState]

func saveActorStateFields(
	_instance *models.ActorState,
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
		case "IsWithProbaility":
			FormDivBasicFieldToField(&(_instance.IsWithProbaility), formDiv)
		case "Probability":
			FormDivEnumStringFieldToField(&(_instance.Probability), formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(_instance.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(_instance.IsExpanded), formDiv)
		case "Diagram:ActorStatesWhoseNodeIsExpanded":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "ActorStatesWhoseNodeIsExpanded", func(owner *models.Diagram) *[]*models.ActorState { return &owner.ActorStatesWhoseNodeIsExpanded })
		case "Scenario:ActorStates":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "ActorStates", func(owner *models.Scenario) *[]*models.ActorState { return &owner.ActorStates })
		}
	}
}

func __gong__New__ActorStateShapeFormCallback(
	_instance *models.ActorStateShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (actorstateshapeFormCallback *FormCallback[*models.ActorStateShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveActorStateShapeFields,
	)
}

type ActorStateShapeFormCallback = FormCallback[*models.ActorStateShape]

func saveActorStateShapeFields(
	_instance *models.ActorStateShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "ActorState":
			FormDivSelectFieldToField(&(_instance.ActorState), probe.stageOfInterest, formDiv)
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
		case "Diagram:ActorStateShapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "ActorStateShapes", func(owner *models.Diagram) *[]*models.ActorStateShape { return &owner.ActorStateShapes })
		}
	}
}

func __gong__New__ActorStateTransitionFormCallback(
	_instance *models.ActorStateTransition,
	probe *Probe,
	formGroup *form.FormGroup,
) (actorstatetransitionFormCallback *FormCallback[*models.ActorStateTransition]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveActorStateTransitionFields,
	)
}

type ActorStateTransitionFormCallback = FormCallback[*models.ActorStateTransition]

func saveActorStateTransitionFields(
	_instance *models.ActorStateTransition,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "StartState":
			FormDivSelectFieldToField(&(_instance.StartState), probe.stageOfInterest, formDiv)
		case "EndState":
			FormDivSelectFieldToField(&(_instance.EndState), probe.stageOfInterest, formDiv)
		case "Justifications":
			FormDivSliceOfPointersToField(_instance, "Justifications", &(_instance.Justifications), formDiv, probe)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(_instance.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(_instance.IsExpanded), formDiv)
		case "Diagram:ActorStateTransitionsWhoseNodeIsExpanded":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "ActorStateTransitionsWhoseNodeIsExpanded", func(owner *models.Diagram) *[]*models.ActorStateTransition { return &owner.ActorStateTransitionsWhoseNodeIsExpanded })
		case "Scenario:ActorStateTransitions":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "ActorStateTransitions", func(owner *models.Scenario) *[]*models.ActorStateTransition { return &owner.ActorStateTransitions })
		}
	}
}

func __gong__New__ActorStateTransitionShapeFormCallback(
	_instance *models.ActorStateTransitionShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (actorstatetransitionshapeFormCallback *FormCallback[*models.ActorStateTransitionShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveActorStateTransitionShapeFields,
	)
}

type ActorStateTransitionShapeFormCallback = FormCallback[*models.ActorStateTransitionShape]

func saveActorStateTransitionShapeFields(
	_instance *models.ActorStateTransitionShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "ActorStateTransition":
			FormDivSelectFieldToField(&(_instance.ActorStateTransition), probe.stageOfInterest, formDiv)
		case "Start":
			FormDivSelectFieldToField(&(_instance.Start), probe.stageOfInterest, formDiv)
		case "End":
			FormDivSelectFieldToField(&(_instance.End), probe.stageOfInterest, formDiv)
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
		case "ControlPointShapes":
			FormDivSliceOfPointersToField(_instance, "ControlPointShapes", &(_instance.ControlPointShapes), formDiv, probe)
		case "Diagram:ActorStateTransitionShapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "ActorStateTransitionShapes", func(owner *models.Diagram) *[]*models.ActorStateTransitionShape { return &owner.ActorStateTransitionShapes })
		}
	}
}

func __gong__New__AnalysisFormCallback(
	_instance *models.Analysis,
	probe *Probe,
	formGroup *form.FormGroup,
) (analysisFormCallback *FormCallback[*models.Analysis]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveAnalysisFields,
	)
}

type AnalysisFormCallback = FormCallback[*models.Analysis]

func saveAnalysisFields(
	_instance *models.Analysis,
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
		case "Scenarios":
			FormDivSliceOfPointersToField(_instance, "Scenarios", &(_instance.Scenarios), formDiv, probe)
		case "IsScenariosNodeExpanded":
			FormDivBasicFieldToField(&(_instance.IsScenariosNodeExpanded), formDiv)
		case "GroupUse":
			FormDivSliceOfPointersToField(_instance, "GroupUse", &(_instance.GroupUse), formDiv, probe)
		case "IsGroupUseNodeExpanded":
			FormDivBasicFieldToField(&(_instance.IsGroupUseNodeExpanded), formDiv)
		case "GeoObjectUse":
			FormDivSliceOfPointersToField(_instance, "GeoObjectUse", &(_instance.GeoObjectUse), formDiv, probe)
		case "IsGeoObjectUseNodeExpanded":
			FormDivBasicFieldToField(&(_instance.IsGeoObjectUseNodeExpanded), formDiv)
		case "MapUse":
			FormDivSliceOfPointersToField(_instance, "MapUse", &(_instance.MapUse), formDiv, probe)
		case "IsMapUseNodeExpanded":
			FormDivBasicFieldToField(&(_instance.IsMapUseNodeExpanded), formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(_instance.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(_instance.IsExpanded), formDiv)
		case "Library:Analyses":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Analyses", func(owner *models.Library) *[]*models.Analysis { return &owner.Analyses })
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
		case "ActorStateTransitionShape:ControlPointShapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "ControlPointShapes", func(owner *models.ActorStateTransitionShape) *[]*models.ControlPointShape { return &owner.ControlPointShapes })
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
		case "IsShowPrefix":
			FormDivBasicFieldToField(&(_instance.IsShowPrefix), formDiv)
		case "Description":
			FormDivBasicFieldToField(&(_instance.Description), formDiv)
		case "EvolutionDirectionShapes":
			FormDivSliceOfPointersToField(_instance, "EvolutionDirectionShapes", &(_instance.EvolutionDirectionShapes), formDiv, probe)
		case "EvolutionDirectionsWhoseNodeIsExpanded":
			FormDivSliceOfPointersToField(_instance, "EvolutionDirectionsWhoseNodeIsExpanded", &(_instance.EvolutionDirectionsWhoseNodeIsExpanded), formDiv, probe)
		case "IsEvolutionDirectionsNodeExpanded":
			FormDivBasicFieldToField(&(_instance.IsEvolutionDirectionsNodeExpanded), formDiv)
		case "ActorStateShapes":
			FormDivSliceOfPointersToField(_instance, "ActorStateShapes", &(_instance.ActorStateShapes), formDiv, probe)
		case "ActorStatesWhoseNodeIsExpanded":
			FormDivSliceOfPointersToField(_instance, "ActorStatesWhoseNodeIsExpanded", &(_instance.ActorStatesWhoseNodeIsExpanded), formDiv, probe)
		case "IsActorStatesNodeExpanded":
			FormDivBasicFieldToField(&(_instance.IsActorStatesNodeExpanded), formDiv)
		case "ParameterShapes":
			FormDivSliceOfPointersToField(_instance, "ParameterShapes", &(_instance.ParameterShapes), formDiv, probe)
		case "ParametersWhoseNodeIsExpanded":
			FormDivSliceOfPointersToField(_instance, "ParametersWhoseNodeIsExpanded", &(_instance.ParametersWhoseNodeIsExpanded), formDiv, probe)
		case "IsParametersNodeExpanded":
			FormDivBasicFieldToField(&(_instance.IsParametersNodeExpanded), formDiv)
		case "ScenarioParameterShapes":
			FormDivSliceOfPointersToField(_instance, "ScenarioParameterShapes", &(_instance.ScenarioParameterShapes), formDiv, probe)
		case "ParametersAggregatesWhoseNodeIsExpanded":
			FormDivSliceOfPointersToField(_instance, "ParametersAggregatesWhoseNodeIsExpanded", &(_instance.ParametersAggregatesWhoseNodeIsExpanded), formDiv, probe)
		case "IsParametersAggregatesNodeExpanded":
			FormDivBasicFieldToField(&(_instance.IsParametersAggregatesNodeExpanded), formDiv)
		case "ActorStateTransitionShapes":
			FormDivSliceOfPointersToField(_instance, "ActorStateTransitionShapes", &(_instance.ActorStateTransitionShapes), formDiv, probe)
		case "ActorStateTransitionsWhoseNodeIsExpanded":
			FormDivSliceOfPointersToField(_instance, "ActorStateTransitionsWhoseNodeIsExpanded", &(_instance.ActorStateTransitionsWhoseNodeIsExpanded), formDiv, probe)
		case "IsActorStateTransitionsNodeExpanded":
			FormDivBasicFieldToField(&(_instance.IsActorStateTransitionsNodeExpanded), formDiv)
		case "AxisOrign_X":
			FormDivBasicFieldToField(&(_instance.AxisOrign_X), formDiv)
		case "AxisOrign_Y":
			FormDivBasicFieldToField(&(_instance.AxisOrign_Y), formDiv)
		case "VerticalAxis_Top_Y":
			FormDivBasicFieldToField(&(_instance.VerticalAxis_Top_Y), formDiv)
		case "VerticalAxis_Bottom_Y":
			FormDivBasicFieldToField(&(_instance.VerticalAxis_Bottom_Y), formDiv)
		case "VerticalAxis_StrokeWidth":
			FormDivBasicFieldToField(&(_instance.VerticalAxis_StrokeWidth), formDiv)
		case "HorizontalAxis_Right_X":
			FormDivBasicFieldToField(&(_instance.HorizontalAxis_Right_X), formDiv)
		case "Start":
			FormDivTimeFieldToField(&(_instance.Start), formDiv, false)
		case "End":
			FormDivTimeFieldToField(&(_instance.End), formDiv, false)
		case "NumberOfYearsBetweenTicks":
			FormDivBasicFieldToField(&(_instance.NumberOfYearsBetweenTicks), formDiv)
		case "Scenario:Diagrams":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Diagrams", func(owner *models.Scenario) *[]*models.Diagram { return &owner.Diagrams })
		}
	}
}

func __gong__New__DocumentFormCallback(
	_instance *models.Document,
	probe *Probe,
	formGroup *form.FormGroup,
) (documentFormCallback *FormCallback[*models.Document]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveDocumentFields,
	)
}

type DocumentFormCallback = FormCallback[*models.Document]

func saveDocumentFields(
	_instance *models.Document,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "GeoObjectUse":
			FormDivSliceOfPointersToField(_instance, "GeoObjectUse", &(_instance.GeoObjectUse), formDiv, probe)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(_instance.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(_instance.IsExpanded), formDiv)
		}
	}
}

func __gong__New__DocumentUseFormCallback(
	_instance *models.DocumentUse,
	probe *Probe,
	formGroup *form.FormGroup,
) (documentuseFormCallback *FormCallback[*models.DocumentUse]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveDocumentUseFields,
	)
}

type DocumentUseFormCallback = FormCallback[*models.DocumentUse]

func saveDocumentUseFields(
	_instance *models.DocumentUse,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Document":
			FormDivSelectFieldToField(&(_instance.Document), probe.stageOfInterest, formDiv)
		case "Parameter:DocumentUse":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "DocumentUse", func(owner *models.Parameter) *[]*models.DocumentUse { return &owner.DocumentUse })
		}
	}
}

func __gong__New__EvolutionDirectionFormCallback(
	_instance *models.EvolutionDirection,
	probe *Probe,
	formGroup *form.FormGroup,
) (evolutiondirectionFormCallback *FormCallback[*models.EvolutionDirection]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveEvolutionDirectionFields,
	)
}

type EvolutionDirectionFormCallback = FormCallback[*models.EvolutionDirection]

func saveEvolutionDirectionFields(
	_instance *models.EvolutionDirection,
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
		case "Diagram:EvolutionDirectionsWhoseNodeIsExpanded":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "EvolutionDirectionsWhoseNodeIsExpanded", func(owner *models.Diagram) *[]*models.EvolutionDirection { return &owner.EvolutionDirectionsWhoseNodeIsExpanded })
		case "Scenario:EvolutionDirections":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "EvolutionDirections", func(owner *models.Scenario) *[]*models.EvolutionDirection { return &owner.EvolutionDirections })
		}
	}
}

func __gong__New__EvolutionDirectionShapeFormCallback(
	_instance *models.EvolutionDirectionShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (evolutiondirectionshapeFormCallback *FormCallback[*models.EvolutionDirectionShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveEvolutionDirectionShapeFields,
	)
}

type EvolutionDirectionShapeFormCallback = FormCallback[*models.EvolutionDirectionShape]

func saveEvolutionDirectionShapeFields(
	_instance *models.EvolutionDirectionShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "EvolutionDirection":
			FormDivSelectFieldToField(&(_instance.EvolutionDirection), probe.stageOfInterest, formDiv)
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
		case "Diagram:EvolutionDirectionShapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "EvolutionDirectionShapes", func(owner *models.Diagram) *[]*models.EvolutionDirectionShape { return &owner.EvolutionDirectionShapes })
		}
	}
}

func __gong__New__FooFormCallback(
	_instance *models.Foo,
	probe *Probe,
	formGroup *form.FormGroup,
) (fooFormCallback *FormCallback[*models.Foo]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveFooFields,
	)
}

type FooFormCallback = FormCallback[*models.Foo]

func saveFooFields(
	_instance *models.Foo,
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

func __gong__New__GeoObjectFormCallback(
	_instance *models.GeoObject,
	probe *Probe,
	formGroup *form.FormGroup,
) (geoobjectFormCallback *FormCallback[*models.GeoObject]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveGeoObjectFields,
	)
}

type GeoObjectFormCallback = FormCallback[*models.GeoObject]

func saveGeoObjectFields(
	_instance *models.GeoObject,
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
		}
	}
}

func __gong__New__GeoObjectUseFormCallback(
	_instance *models.GeoObjectUse,
	probe *Probe,
	formGroup *form.FormGroup,
) (geoobjectuseFormCallback *FormCallback[*models.GeoObjectUse]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveGeoObjectUseFields,
	)
}

type GeoObjectUseFormCallback = FormCallback[*models.GeoObjectUse]

func saveGeoObjectUseFields(
	_instance *models.GeoObjectUse,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "GeoObject":
			FormDivSelectFieldToField(&(_instance.GeoObject), probe.stageOfInterest, formDiv)
		case "Analysis:GeoObjectUse":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "GeoObjectUse", func(owner *models.Analysis) *[]*models.GeoObjectUse { return &owner.GeoObjectUse })
		case "Document:GeoObjectUse":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "GeoObjectUse", func(owner *models.Document) *[]*models.GeoObjectUse { return &owner.GeoObjectUse })
		case "Parameter:GeoObjectUse":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "GeoObjectUse", func(owner *models.Parameter) *[]*models.GeoObjectUse { return &owner.GeoObjectUse })
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
		case "UserUse":
			FormDivSliceOfPointersToField(_instance, "UserUse", &(_instance.UserUse), formDiv, probe)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(_instance.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(_instance.IsExpanded), formDiv)
		}
	}
}

func __gong__New__GroupUseFormCallback(
	_instance *models.GroupUse,
	probe *Probe,
	formGroup *form.FormGroup,
) (groupuseFormCallback *FormCallback[*models.GroupUse]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveGroupUseFields,
	)
}

type GroupUseFormCallback = FormCallback[*models.GroupUse]

func saveGroupUseFields(
	_instance *models.GroupUse,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Group":
			FormDivSelectFieldToField(&(_instance.Group), probe.stageOfInterest, formDiv)
		case "Analysis:GroupUse":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "GroupUse", func(owner *models.Analysis) *[]*models.GroupUse { return &owner.GroupUse })
		case "Parameter:GroupUse":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "GroupUse", func(owner *models.Parameter) *[]*models.GroupUse { return &owner.GroupUse })
		case "Repository:GroupUse":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "GroupUse", func(owner *models.Repository) *[]*models.GroupUse { return &owner.GroupUse })
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
		case "Analyses":
			FormDivSliceOfPointersToField(_instance, "Analyses", &(_instance.Analyses), formDiv, probe)
		case "IsAnalysesNodeExpanded":
			FormDivBasicFieldToField(&(_instance.IsAnalysesNodeExpanded), formDiv)
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
		case "IsExpandedTmp":
			FormDivBasicFieldToField(&(_instance.IsExpandedTmp), formDiv)
		case "Library:SubLibraries":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "SubLibraries", func(owner *models.Library) *[]*models.Library { return &owner.SubLibraries })
		case "Library:SubLibrariesWhoseNodeIsExpanded":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "SubLibrariesWhoseNodeIsExpanded", func(owner *models.Library) *[]*models.Library { return &owner.SubLibrariesWhoseNodeIsExpanded })
		}
	}
}

func __gong__New__MapObjectFormCallback(
	_instance *models.MapObject,
	probe *Probe,
	formGroup *form.FormGroup,
) (mapobjectFormCallback *FormCallback[*models.MapObject]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveMapObjectFields,
	)
}

type MapObjectFormCallback = FormCallback[*models.MapObject]

func saveMapObjectFields(
	_instance *models.MapObject,
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
		}
	}
}

func __gong__New__MapObjectUseFormCallback(
	_instance *models.MapObjectUse,
	probe *Probe,
	formGroup *form.FormGroup,
) (mapobjectuseFormCallback *FormCallback[*models.MapObjectUse]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveMapObjectUseFields,
	)
}

type MapObjectUseFormCallback = FormCallback[*models.MapObjectUse]

func saveMapObjectUseFields(
	_instance *models.MapObjectUse,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Map":
			FormDivSelectFieldToField(&(_instance.Map), probe.stageOfInterest, formDiv)
		case "Analysis:MapUse":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "MapUse", func(owner *models.Analysis) *[]*models.MapObjectUse { return &owner.MapUse })
		}
	}
}

func __gong__New__ParameterFormCallback(
	_instance *models.Parameter,
	probe *Probe,
	formGroup *form.FormGroup,
) (parameterFormCallback *FormCallback[*models.Parameter]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveParameterFields,
	)
}

type ParameterFormCallback = FormCallback[*models.Parameter]

func saveParameterFields(
	_instance *models.Parameter,
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
		case "IsResponse":
			FormDivBasicFieldToField(&(_instance.IsResponse), formDiv)
		case "Start":
			FormDivTimeFieldToField(&(_instance.Start), formDiv, false)
		case "End":
			FormDivTimeFieldToField(&(_instance.End), formDiv, false)
		case "Force":
			FormDivBasicFieldToField(&(_instance.Force), formDiv)
		case "GroupUse":
			FormDivSliceOfPointersToField(_instance, "GroupUse", &(_instance.GroupUse), formDiv, probe)
		case "DocumentUse":
			FormDivSliceOfPointersToField(_instance, "DocumentUse", &(_instance.DocumentUse), formDiv, probe)
		case "GeoObjectUse":
			FormDivSliceOfPointersToField(_instance, "GeoObjectUse", &(_instance.GeoObjectUse), formDiv, probe)
		case "Tag":
			FormDivBasicFieldToField(&(_instance.Tag), formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(_instance.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(_instance.IsExpanded), formDiv)
		case "ActorStateTransition:Justifications":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Justifications", func(owner *models.ActorStateTransition) *[]*models.Parameter { return &owner.Justifications })
		case "Diagram:ParametersWhoseNodeIsExpanded":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "ParametersWhoseNodeIsExpanded", func(owner *models.Diagram) *[]*models.Parameter { return &owner.ParametersWhoseNodeIsExpanded })
		case "ParametersAggregate:Parameters":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Parameters", func(owner *models.ParametersAggregate) *[]*models.Parameter { return &owner.Parameters })
		case "Scenario:Parameters":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Parameters", func(owner *models.Scenario) *[]*models.Parameter { return &owner.Parameters })
		}
	}
}

func __gong__New__ParameterCategoryFormCallback(
	_instance *models.ParameterCategory,
	probe *Probe,
	formGroup *form.FormGroup,
) (parametercategoryFormCallback *FormCallback[*models.ParameterCategory]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveParameterCategoryFields,
	)
}

type ParameterCategoryFormCallback = FormCallback[*models.ParameterCategory]

func saveParameterCategoryFields(
	_instance *models.ParameterCategory,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "ParameterUse":
			FormDivSliceOfPointersToField(_instance, "ParameterUse", &(_instance.ParameterUse), formDiv, probe)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(_instance.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(_instance.IsExpanded), formDiv)
		}
	}
}

func __gong__New__ParameterCategoryUseFormCallback(
	_instance *models.ParameterCategoryUse,
	probe *Probe,
	formGroup *form.FormGroup,
) (parametercategoryuseFormCallback *FormCallback[*models.ParameterCategoryUse]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveParameterCategoryUseFields,
	)
}

type ParameterCategoryUseFormCallback = FormCallback[*models.ParameterCategoryUse]

func saveParameterCategoryUseFields(
	_instance *models.ParameterCategoryUse,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "ParameterCategory":
			FormDivSelectFieldToField(&(_instance.ParameterCategory), probe.stageOfInterest, formDiv)
		}
	}
}

func __gong__New__ParameterShapeFormCallback(
	_instance *models.ParameterShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (parametershapeFormCallback *FormCallback[*models.ParameterShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveParameterShapeFields,
	)
}

type ParameterShapeFormCallback = FormCallback[*models.ParameterShape]

func saveParameterShapeFields(
	_instance *models.ParameterShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Parameter":
			FormDivSelectFieldToField(&(_instance.Parameter), probe.stageOfInterest, formDiv)
		case "Direction":
			FormDivEnumStringFieldToField(&(_instance.Direction), formDiv)
		case "ShapeIsComputedFromModel":
			FormDivBasicFieldToField(&(_instance.ShapeIsComputedFromModel), formDiv)
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
		case "Diagram:ParameterShapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "ParameterShapes", func(owner *models.Diagram) *[]*models.ParameterShape { return &owner.ParameterShapes })
		case "ParameterCategory:ParameterUse":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "ParameterUse", func(owner *models.ParameterCategory) *[]*models.ParameterShape { return &owner.ParameterUse })
		case "Repository:ParameterUse":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "ParameterUse", func(owner *models.Repository) *[]*models.ParameterShape { return &owner.ParameterUse })
		}
	}
}

func __gong__New__ParametersAggregateFormCallback(
	_instance *models.ParametersAggregate,
	probe *Probe,
	formGroup *form.FormGroup,
) (parametersaggregateFormCallback *FormCallback[*models.ParametersAggregate]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveParametersAggregateFields,
	)
}

type ParametersAggregateFormCallback = FormCallback[*models.ParametersAggregate]

func saveParametersAggregateFields(
	_instance *models.ParametersAggregate,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Tag":
			FormDivBasicFieldToField(&(_instance.Tag), formDiv)
		case "Description":
			FormDivBasicFieldToField(&(_instance.Description), formDiv)
		case "Parameters":
			FormDivSliceOfPointersToField(_instance, "Parameters", &(_instance.Parameters), formDiv, probe)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(_instance.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(_instance.IsExpanded), formDiv)
		case "Diagram:ParametersAggregatesWhoseNodeIsExpanded":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "ParametersAggregatesWhoseNodeIsExpanded", func(owner *models.Diagram) *[]*models.ParametersAggregate { return &owner.ParametersAggregatesWhoseNodeIsExpanded })
		case "Scenario:ParametersAggretates":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "ParametersAggretates", func(owner *models.Scenario) *[]*models.ParametersAggregate { return &owner.ParametersAggretates })
		}
	}
}

func __gong__New__ParametersAggregateShapeFormCallback(
	_instance *models.ParametersAggregateShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (parametersaggregateshapeFormCallback *FormCallback[*models.ParametersAggregateShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveParametersAggregateShapeFields,
	)
}

type ParametersAggregateShapeFormCallback = FormCallback[*models.ParametersAggregateShape]

func saveParametersAggregateShapeFields(
	_instance *models.ParametersAggregateShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "ScenarioParameter":
			FormDivSelectFieldToField(&(_instance.ScenarioParameter), probe.stageOfInterest, formDiv)
		case "Direction":
			FormDivEnumStringFieldToField(&(_instance.Direction), formDiv)
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
		case "Diagram:ScenarioParameterShapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "ScenarioParameterShapes", func(owner *models.Diagram) *[]*models.ParametersAggregateShape { return &owner.ScenarioParameterShapes })
		}
	}
}

func __gong__New__PositionFormCallback(
	_instance *models.Position,
	probe *Probe,
	formGroup *form.FormGroup,
) (positionFormCallback *FormCallback[*models.Position]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		savePositionFields,
	)
}

type PositionFormCallback = FormCallback[*models.Position]

func savePositionFields(
	_instance *models.Position,
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
		case "Ordinate":
			FormDivBasicFieldToField(&(_instance.Ordinate), formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(_instance.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(_instance.IsExpanded), formDiv)
		}
	}
}

func __gong__New__RepositoryFormCallback(
	_instance *models.Repository,
	probe *Probe,
	formGroup *form.FormGroup,
) (repositoryFormCallback *FormCallback[*models.Repository]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveRepositoryFields,
	)
}

type RepositoryFormCallback = FormCallback[*models.Repository]

func saveRepositoryFields(
	_instance *models.Repository,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "ParameterUse":
			FormDivSliceOfPointersToField(_instance, "ParameterUse", &(_instance.ParameterUse), formDiv, probe)
		case "GroupUse":
			FormDivSliceOfPointersToField(_instance, "GroupUse", &(_instance.GroupUse), formDiv, probe)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(_instance.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(_instance.IsExpanded), formDiv)
		}
	}
}

func __gong__New__ScenarioFormCallback(
	_instance *models.Scenario,
	probe *Probe,
	formGroup *form.FormGroup,
) (scenarioFormCallback *FormCallback[*models.Scenario]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveScenarioFields,
	)
}

type ScenarioFormCallback = FormCallback[*models.Scenario]

func saveScenarioFields(
	_instance *models.Scenario,
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
		case "Diagrams":
			FormDivSliceOfPointersToField(_instance, "Diagrams", &(_instance.Diagrams), formDiv, probe)
		case "IsDiagramsNodeExpanded":
			FormDivBasicFieldToField(&(_instance.IsDiagramsNodeExpanded), formDiv)
		case "ActorStates":
			FormDivSliceOfPointersToField(_instance, "ActorStates", &(_instance.ActorStates), formDiv, probe)
		case "IsActorStatesNodeExpanded":
			FormDivBasicFieldToField(&(_instance.IsActorStatesNodeExpanded), formDiv)
		case "ActorStateTransitions":
			FormDivSliceOfPointersToField(_instance, "ActorStateTransitions", &(_instance.ActorStateTransitions), formDiv, probe)
		case "IsActorStateTransitionsNodeExpanded":
			FormDivBasicFieldToField(&(_instance.IsActorStateTransitionsNodeExpanded), formDiv)
		case "EvolutionDirections":
			FormDivSliceOfPointersToField(_instance, "EvolutionDirections", &(_instance.EvolutionDirections), formDiv, probe)
		case "IsEvolutionDirectionsNodeExpanded":
			FormDivBasicFieldToField(&(_instance.IsEvolutionDirectionsNodeExpanded), formDiv)
		case "Parameters":
			FormDivSliceOfPointersToField(_instance, "Parameters", &(_instance.Parameters), formDiv, probe)
		case "IsParametersNodeExpanded":
			FormDivBasicFieldToField(&(_instance.IsParametersNodeExpanded), formDiv)
		case "ParametersAggretates":
			FormDivSliceOfPointersToField(_instance, "ParametersAggretates", &(_instance.ParametersAggretates), formDiv, probe)
		case "IsParametersAggretatesNodeExpanded":
			FormDivBasicFieldToField(&(_instance.IsParametersAggretatesNodeExpanded), formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(_instance.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(_instance.IsExpanded), formDiv)
		case "Analysis:Scenarios":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Scenarios", func(owner *models.Analysis) *[]*models.Scenario { return &owner.Scenarios })
		}
	}
}

func __gong__New__UserFormCallback(
	_instance *models.User,
	probe *Probe,
	formGroup *form.FormGroup,
) (userFormCallback *FormCallback[*models.User]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveUserFields,
	)
}

type UserFormCallback = FormCallback[*models.User]

func saveUserFields(
	_instance *models.User,
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
		}
	}
}

func __gong__New__UserUseFormCallback(
	_instance *models.UserUse,
	probe *Probe,
	formGroup *form.FormGroup,
) (useruseFormCallback *FormCallback[*models.UserUse]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveUserUseFields,
	)
}

type UserUseFormCallback = FormCallback[*models.UserUse]

func saveUserUseFields(
	_instance *models.UserUse,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "User":
			FormDivSelectFieldToField(&(_instance.User), probe.stageOfInterest, formDiv)
		case "Group:UserUse":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "UserUse", func(owner *models.Group) *[]*models.UserUse { return &owner.UserUse })
		}
	}
}

func __gong__New__WorkspaceFormCallback(
	_instance *models.Workspace,
	probe *Probe,
	formGroup *form.FormGroup,
) (workspaceFormCallback *FormCallback[*models.Workspace]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveWorkspaceFields,
	)
}

type WorkspaceFormCallback = FormCallback[*models.Workspace]

func saveWorkspaceFields(
	_instance *models.Workspace,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "SelectedDiagram":
			FormDivSelectFieldToField(&(_instance.SelectedDiagram), probe.stageOfInterest, formDiv)
		case "Default_EvolutionDirectionShape":
			FormDivSelectFieldToField(&(_instance.Default_EvolutionDirectionShape), probe.stageOfInterest, formDiv)
		case "Default_ParameterShape":
			FormDivSelectFieldToField(&(_instance.Default_ParameterShape), probe.stageOfInterest, formDiv)
		case "Default_ScenarioParameterShape":
			FormDivSelectFieldToField(&(_instance.Default_ScenarioParameterShape), probe.stageOfInterest, formDiv)
		case "Default_ActorStateShape":
			FormDivSelectFieldToField(&(_instance.Default_ActorStateShape), probe.stageOfInterest, formDiv)
		case "Default_ActorStateTransitionShape":
			FormDivSelectFieldToField(&(_instance.Default_ActorStateTransitionShape), probe.stageOfInterest, formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(_instance.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(_instance.IsExpanded), formDiv)
		}
	}
}

