// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/dsm/statemachines/go/models"
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
func __gong__New__ActionFormCallback(
	_instance *models.Action,
	probe *Probe,
	formGroup *form.FormGroup,
) (actionFormCallback *FormCallback[*models.Action]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveActionFields,
	)
}

type ActionFormCallback = FormCallback[*models.Action]

func saveActionFields(
	_instance *models.Action,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Criticality":
			FormDivEnumStringFieldToField(&(_instance.Criticality), formDiv)
		}
	}
}

func __gong__New__ActivitiesFormCallback(
	_instance *models.Activities,
	probe *Probe,
	formGroup *form.FormGroup,
) (activitiesFormCallback *FormCallback[*models.Activities]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveActivitiesFields,
	)
}

type ActivitiesFormCallback = FormCallback[*models.Activities]

func saveActivitiesFields(
	_instance *models.Activities,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Criticality":
			FormDivEnumStringFieldToField(&(_instance.Criticality), formDiv)
		case "State:Activities":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Activities", func(owner *models.State) *[]*models.Activities { return &owner.Activities })
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
		case "IsChecked":
			FormDivBasicFieldToField(&(_instance.IsChecked), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(_instance.IsExpanded), formDiv)
		case "IsEditable_":
			FormDivBasicFieldToField(&(_instance.IsEditable_), formDiv)
		case "IsStatesNodeExpanded":
			FormDivBasicFieldToField(&(_instance.IsStatesNodeExpanded), formDiv)
		case "State_Shapes":
			FormDivSliceOfPointersToField(_instance, "State_Shapes", &(_instance.State_Shapes), formDiv, probe)
		case "StatesWhoseNodeIsExpanded":
			FormDivSliceOfPointersToField(_instance, "StatesWhoseNodeIsExpanded", &(_instance.StatesWhoseNodeIsExpanded), formDiv, probe)
		case "Transition_Shapes":
			FormDivSliceOfPointersToField(_instance, "Transition_Shapes", &(_instance.Transition_Shapes), formDiv, probe)
		case "Note_Shapes":
			FormDivSliceOfPointersToField(_instance, "Note_Shapes", &(_instance.Note_Shapes), formDiv, probe)
		case "NoteState_Shapes":
			FormDivSliceOfPointersToField(_instance, "NoteState_Shapes", &(_instance.NoteState_Shapes), formDiv, probe)
		case "ShowRoles":
			FormDivBasicFieldToField(&(_instance.ShowRoles), formDiv)
		case "ShowMessages":
			FormDivBasicFieldToField(&(_instance.ShowMessages), formDiv)
		case "Library:Diagrams":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Diagrams", func(owner *models.Library) *[]*models.Diagram { return &owner.Diagrams })
		case "State:Diagrams":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Diagrams", func(owner *models.State) *[]*models.Diagram { return &owner.Diagrams })
		case "StateMachine:Diagrams":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Diagrams", func(owner *models.StateMachine) *[]*models.Diagram { return &owner.Diagrams })
		case "Transition:Diagrams":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Diagrams", func(owner *models.Transition) *[]*models.Diagram { return &owner.Diagrams })
		}
	}
}

func __gong__New__GuardFormCallback(
	_instance *models.Guard,
	probe *Probe,
	formGroup *form.FormGroup,
) (guardFormCallback *FormCallback[*models.Guard]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveGuardFields,
	)
}

type GuardFormCallback = FormCallback[*models.Guard]

func saveGuardFields(
	_instance *models.Guard,
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
		case "SubLibraries":
			FormDivSliceOfPointersToField(_instance, "SubLibraries", &(_instance.SubLibraries), formDiv, probe)
		case "NbPixPerCharacter":
			FormDivBasicFieldToField(&(_instance.NbPixPerCharacter), formDiv)
		case "LogoSVGFile":
			FormDivBasicFieldToField(&(_instance.LogoSVGFile), formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(_instance.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(_instance.IsExpanded), formDiv)
		case "IsRootLibrary":
			FormDivBasicFieldToField(&(_instance.IsRootLibrary), formDiv)
		case "Diagrams":
			FormDivSliceOfPointersToField(_instance, "Diagrams", &(_instance.Diagrams), formDiv, probe)
		case "RootStateMachines":
			FormDivSliceOfPointersToField(_instance, "RootStateMachines", &(_instance.RootStateMachines), formDiv, probe)
		case "IsStateMachinesNodeExpanded":
			FormDivBasicFieldToField(&(_instance.IsStateMachinesNodeExpanded), formDiv)
		case "StateMachinesWhoseNodeIsExpanded":
			FormDivSliceOfPointersToField(_instance, "StateMachinesWhoseNodeIsExpanded", &(_instance.StateMachinesWhoseNodeIsExpanded), formDiv, probe)
		case "IsSubLibrariesNodeExpanded":
			FormDivBasicFieldToField(&(_instance.IsSubLibrariesNodeExpanded), formDiv)
		case "SubLibrariesWhoseNodeIsExpanded":
			FormDivSliceOfPointersToField(_instance, "SubLibrariesWhoseNodeIsExpanded", &(_instance.SubLibrariesWhoseNodeIsExpanded), formDiv, probe)
		case "IsExpandedTmp":
			FormDivBasicFieldToField(&(_instance.IsExpandedTmp), formDiv)
		case "Roles":
			FormDivSliceOfPointersToField(_instance, "Roles", &(_instance.Roles), formDiv, probe)
		case "IsRolesNodeExpanded":
			FormDivBasicFieldToField(&(_instance.IsRolesNodeExpanded), formDiv)
		case "MessageTypes":
			FormDivSliceOfPointersToField(_instance, "MessageTypes", &(_instance.MessageTypes), formDiv, probe)
		case "IsMessageTypesNodeExpanded":
			FormDivBasicFieldToField(&(_instance.IsMessageTypesNodeExpanded), formDiv)
		case "Library:SubLibraries":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "SubLibraries", func(owner *models.Library) *[]*models.Library { return &owner.SubLibraries })
		case "Library:SubLibrariesWhoseNodeIsExpanded":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "SubLibrariesWhoseNodeIsExpanded", func(owner *models.Library) *[]*models.Library { return &owner.SubLibrariesWhoseNodeIsExpanded })
		}
	}
}

func __gong__New__MessageFormCallback(
	_instance *models.Message,
	probe *Probe,
	formGroup *form.FormGroup,
) (messageFormCallback *FormCallback[*models.Message]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveMessageFields,
	)
}

type MessageFormCallback = FormCallback[*models.Message]

func saveMessageFields(
	_instance *models.Message,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "IsSelected":
			FormDivBasicFieldToField(&(_instance.IsSelected), formDiv)
		case "MessageType":
			FormDivSelectFieldToField(&(_instance.MessageType), probe.stageOfInterest, formDiv)
		case "OriginTransition":
			FormDivSelectFieldToField(&(_instance.OriginTransition), probe.stageOfInterest, formDiv)
		case "Object:Messages":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Messages", func(owner *models.Object) *[]*models.Message { return &owner.Messages })
		}
	}
}

func __gong__New__MessageTypeFormCallback(
	_instance *models.MessageType,
	probe *Probe,
	formGroup *form.FormGroup,
) (messagetypeFormCallback *FormCallback[*models.MessageType]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveMessageTypeFields,
	)
}

type MessageTypeFormCallback = FormCallback[*models.MessageType]

func saveMessageTypeFields(
	_instance *models.MessageType,
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
		case "Library:MessageTypes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "MessageTypes", func(owner *models.Library) *[]*models.MessageType { return &owner.MessageTypes })
		case "Transition:GeneratedMessages":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "GeneratedMessages", func(owner *models.Transition) *[]*models.MessageType { return &owner.GeneratedMessages })
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
		case "State":
			FormDivSelectFieldToField(&(_instance.State), probe.stageOfInterest, formDiv)
		case "State:Notes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Notes", func(owner *models.State) *[]*models.Note { return &owner.Notes })
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
		case "OverideLayoutDirection":
			FormDivBasicFieldToField(&(_instance.OverideLayoutDirection), formDiv)
		case "LayoutDirection":
			FormDivEnumIntFieldToField(&(_instance.LayoutDirection), formDiv)
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

func __gong__New__NoteStateShapeFormCallback(
	_instance *models.NoteStateShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (notestateshapeFormCallback *FormCallback[*models.NoteStateShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveNoteStateShapeFields,
	)
}

type NoteStateShapeFormCallback = FormCallback[*models.NoteStateShape]

func saveNoteStateShapeFields(
	_instance *models.NoteStateShape,
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
		case "State":
			FormDivSelectFieldToField(&(_instance.State), probe.stageOfInterest, formDiv)
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
		case "Diagram:NoteState_Shapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "NoteState_Shapes", func(owner *models.Diagram) *[]*models.NoteStateShape { return &owner.NoteState_Shapes })
		}
	}
}

func __gong__New__ObjectFormCallback(
	_instance *models.Object,
	probe *Probe,
	formGroup *form.FormGroup,
) (objectFormCallback *FormCallback[*models.Object]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveObjectFields,
	)
}

type ObjectFormCallback = FormCallback[*models.Object]

func saveObjectFields(
	_instance *models.Object,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "State":
			FormDivSelectFieldToField(&(_instance.State), probe.stageOfInterest, formDiv)
		case "IsSelected":
			FormDivBasicFieldToField(&(_instance.IsSelected), formDiv)
		case "Rank":
			FormDivBasicFieldToField(&(_instance.Rank), formDiv)
		case "DOF":
			FormDivTimeFieldToField(&(_instance.DOF), formDiv, false)
		case "Messages":
			FormDivSliceOfPointersToField(_instance, "Messages", &(_instance.Messages), formDiv, probe)
		}
	}
}

func __gong__New__RoleFormCallback(
	_instance *models.Role,
	probe *Probe,
	formGroup *form.FormGroup,
) (roleFormCallback *FormCallback[*models.Role]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveRoleFields,
	)
}

type RoleFormCallback = FormCallback[*models.Role]

func saveRoleFields(
	_instance *models.Role,
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
		case "RolesWithSamePermissions":
			FormDivSliceOfPointersToField(_instance, "RolesWithSamePermissions", &(_instance.RolesWithSamePermissions), formDiv, probe)
		case "Library:Roles":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Roles", func(owner *models.Library) *[]*models.Role { return &owner.Roles })
		case "Role:RolesWithSamePermissions":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "RolesWithSamePermissions", func(owner *models.Role) *[]*models.Role { return &owner.RolesWithSamePermissions })
		case "Transition:RolesWithPermissions":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "RolesWithPermissions", func(owner *models.Transition) *[]*models.Role { return &owner.RolesWithPermissions })
		}
	}
}

func __gong__New__StateFormCallback(
	_instance *models.State,
	probe *Probe,
	formGroup *form.FormGroup,
) (stateFormCallback *FormCallback[*models.State]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveStateFields,
	)
}

type StateFormCallback = FormCallback[*models.State]

func saveStateFields(
	_instance *models.State,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "IsEndState":
			FormDivBasicFieldToField(&(_instance.IsEndState), formDiv)
		case "IsDecisionNode":
			FormDivBasicFieldToField(&(_instance.IsDecisionNode), formDiv)
		case "SubStates":
			FormDivSliceOfPointersToField(_instance, "SubStates", &(_instance.SubStates), formDiv, probe)
		case "Entry":
			FormDivSelectFieldToField(&(_instance.Entry), probe.stageOfInterest, formDiv)
		case "Activities":
			FormDivSliceOfPointersToField(_instance, "Activities", &(_instance.Activities), formDiv, probe)
		case "Exit":
			FormDivSelectFieldToField(&(_instance.Exit), probe.stageOfInterest, formDiv)
		case "Parent":
			FormDivSelectFieldToField(&(_instance.Parent), probe.stageOfInterest, formDiv)
		case "IsFictious":
			FormDivBasicFieldToField(&(_instance.IsFictious), formDiv)
		case "Diagrams":
			FormDivSliceOfPointersToField(_instance, "Diagrams", &(_instance.Diagrams), formDiv, probe)
		case "Notes":
			FormDivSliceOfPointersToField(_instance, "Notes", &(_instance.Notes), formDiv, probe)
		case "Diagram:StatesWhoseNodeIsExpanded":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "StatesWhoseNodeIsExpanded", func(owner *models.Diagram) *[]*models.State { return &owner.StatesWhoseNodeIsExpanded })
		case "State:SubStates":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "SubStates", func(owner *models.State) *[]*models.State { return &owner.SubStates })
		case "StateMachine:States":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "States", func(owner *models.StateMachine) *[]*models.State { return &owner.States })
		}
	}
}

func __gong__New__StateMachineFormCallback(
	_instance *models.StateMachine,
	probe *Probe,
	formGroup *form.FormGroup,
) (statemachineFormCallback *FormCallback[*models.StateMachine]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveStateMachineFields,
	)
}

type StateMachineFormCallback = FormCallback[*models.StateMachine]

func saveStateMachineFields(
	_instance *models.StateMachine,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "InitialState":
			FormDivSelectFieldToField(&(_instance.InitialState), probe.stageOfInterest, formDiv)
		case "States":
			FormDivSliceOfPointersToField(_instance, "States", &(_instance.States), formDiv, probe)
		case "Diagrams":
			FormDivSliceOfPointersToField(_instance, "Diagrams", &(_instance.Diagrams), formDiv, probe)
		case "IsWithTransitionNameAutonamticalyGenerated":
			FormDivBasicFieldToField(&(_instance.IsWithTransitionNameAutonamticalyGenerated), formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(_instance.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(_instance.IsExpanded), formDiv)
		case "Library:RootStateMachines":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "RootStateMachines", func(owner *models.Library) *[]*models.StateMachine { return &owner.RootStateMachines })
		case "Library:StateMachinesWhoseNodeIsExpanded":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "StateMachinesWhoseNodeIsExpanded", func(owner *models.Library) *[]*models.StateMachine { return &owner.StateMachinesWhoseNodeIsExpanded })
		}
	}
}

func __gong__New__StateShapeFormCallback(
	_instance *models.StateShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (stateshapeFormCallback *FormCallback[*models.StateShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveStateShapeFields,
	)
}

type StateShapeFormCallback = FormCallback[*models.StateShape]

func saveStateShapeFields(
	_instance *models.StateShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "State":
			FormDivSelectFieldToField(&(_instance.State), probe.stageOfInterest, formDiv)
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
		case "Diagram:State_Shapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "State_Shapes", func(owner *models.Diagram) *[]*models.StateShape { return &owner.State_Shapes })
		}
	}
}

func __gong__New__TransitionFormCallback(
	_instance *models.Transition,
	probe *Probe,
	formGroup *form.FormGroup,
) (transitionFormCallback *FormCallback[*models.Transition]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveTransitionFields,
	)
}

type TransitionFormCallback = FormCallback[*models.Transition]

func saveTransitionFields(
	_instance *models.Transition,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Start":
			FormDivSelectFieldToField(&(_instance.Start), probe.stageOfInterest, formDiv)
		case "End":
			FormDivSelectFieldToField(&(_instance.End), probe.stageOfInterest, formDiv)
		case "RolesWithPermissions":
			FormDivSliceOfPointersToField(_instance, "RolesWithPermissions", &(_instance.RolesWithPermissions), formDiv, probe)
		case "GeneratedMessages":
			FormDivSliceOfPointersToField(_instance, "GeneratedMessages", &(_instance.GeneratedMessages), formDiv, probe)
		case "Guard":
			FormDivSelectFieldToField(&(_instance.Guard), probe.stageOfInterest, formDiv)
		case "Diagrams":
			FormDivSliceOfPointersToField(_instance, "Diagrams", &(_instance.Diagrams), formDiv, probe)
		case "IsExpanded":
			FormDivBasicFieldToField(&(_instance.IsExpanded), formDiv)
		case "IsRolesNodeExpanded":
			FormDivBasicFieldToField(&(_instance.IsRolesNodeExpanded), formDiv)
		case "IsMessagesNodeExpanded":
			FormDivBasicFieldToField(&(_instance.IsMessagesNodeExpanded), formDiv)
		}
	}
}

func __gong__New__Transition_ShapeFormCallback(
	_instance *models.Transition_Shape,
	probe *Probe,
	formGroup *form.FormGroup,
) (transition_shapeFormCallback *FormCallback[*models.Transition_Shape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveTransition_ShapeFields,
	)
}

type Transition_ShapeFormCallback = FormCallback[*models.Transition_Shape]

func saveTransition_ShapeFields(
	_instance *models.Transition_Shape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Transition":
			FormDivSelectFieldToField(&(_instance.Transition), probe.stageOfInterest, formDiv)
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
		case "Diagram:Transition_Shapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Transition_Shapes", func(owner *models.Diagram) *[]*models.Transition_Shape { return &owner.Transition_Shapes })
		}
	}
}

