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

// insertion point
func __gong__New__ActionFormCallback(
	action *models.Action,
	probe *Probe,
	formGroup *form.FormGroup,
) (actionFormCallback *ActionFormCallback) {
	actionFormCallback = new(ActionFormCallback)
	actionFormCallback.probe = probe
	actionFormCallback.action = action
	actionFormCallback.formGroup = formGroup

	actionFormCallback.CreationMode = (action == nil)

	return
}

type ActionFormCallback struct {
	action *models.Action

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (actionFormCallback *ActionFormCallback) OnSave() {
	actionFormCallback.probe.stageOfInterest.Lock()
	defer actionFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("ActionFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	actionFormCallback.probe.formStage.Checkout()

	if actionFormCallback.action == nil {
		actionFormCallback.action = new(models.Action).Stage(actionFormCallback.probe.stageOfInterest)
	}
	action_ := actionFormCallback.action
	_ = action_

	for _, formDiv := range actionFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(action_.Name), formDiv)
		case "Criticality":
			FormDivEnumStringFieldToField(&(action_.Criticality), formDiv)
		}
	}

	// manage the suppress operation
	if actionFormCallback.formGroup.HasSuppressButtonBeenPressed {
		action_.Unstage(actionFormCallback.probe.stageOfInterest)
	}

	actionFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Action](
		actionFormCallback.probe,
	)

	// display a new form by reset the form stage
	if actionFormCallback.CreationMode || actionFormCallback.formGroup.HasSuppressButtonBeenPressed {
		actionFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(actionFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ActionFormCallback(
			nil,
			actionFormCallback.probe,
			newFormGroup,
		)
		action := new(models.Action)
		FillUpForm(action, newFormGroup, actionFormCallback.probe)
		actionFormCallback.probe.formStage.Commit()
	}

	actionFormCallback.probe.ux_tree()
}
func __gong__New__ActivitiesFormCallback(
	activities *models.Activities,
	probe *Probe,
	formGroup *form.FormGroup,
) (activitiesFormCallback *ActivitiesFormCallback) {
	activitiesFormCallback = new(ActivitiesFormCallback)
	activitiesFormCallback.probe = probe
	activitiesFormCallback.activities = activities
	activitiesFormCallback.formGroup = formGroup

	activitiesFormCallback.CreationMode = (activities == nil)

	return
}

type ActivitiesFormCallback struct {
	activities *models.Activities

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (activitiesFormCallback *ActivitiesFormCallback) OnSave() {
	activitiesFormCallback.probe.stageOfInterest.Lock()
	defer activitiesFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("ActivitiesFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	activitiesFormCallback.probe.formStage.Checkout()

	if activitiesFormCallback.activities == nil {
		activitiesFormCallback.activities = new(models.Activities).Stage(activitiesFormCallback.probe.stageOfInterest)
	}
	activities_ := activitiesFormCallback.activities
	_ = activities_

	for _, formDiv := range activitiesFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(activities_.Name), formDiv)
		case "Criticality":
			FormDivEnumStringFieldToField(&(activities_.Criticality), formDiv)
		case "State:Activities":
			// 1. Decode the AssociationStorage which contains the rowIDs of the State instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target State instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.State](activitiesFormCallback.probe.stageOfInterest)
			targetStateIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetStateIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all State instances and update their Activities slice
			for _state := range *models.GetGongstructInstancesSetFromPointerType[*models.State](activitiesFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(activitiesFormCallback.probe.stageOfInterest, _state)
				
				// if State is selected
				if targetStateIDs[id] {
					// ensure activities_ is in _state.Activities
					found := false
					for _, _b := range _state.Activities {
						if _b == activities_ {
							found = true
							break
						}
					}
					if !found {
						_state.Activities = append(_state.Activities, activities_)
						activitiesFormCallback.probe.UpdateSliceOfPointersCallback(_state, "Activities", &_state.Activities)
					}
				} else {
					// ensure activities_ is NOT in _state.Activities
					idx := slices.Index(_state.Activities, activities_)
					if idx != -1 {
						_state.Activities = slices.Delete(_state.Activities, idx, idx+1)
						activitiesFormCallback.probe.UpdateSliceOfPointersCallback(_state, "Activities", &_state.Activities)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if activitiesFormCallback.formGroup.HasSuppressButtonBeenPressed {
		activities_.Unstage(activitiesFormCallback.probe.stageOfInterest)
	}

	activitiesFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Activities](
		activitiesFormCallback.probe,
	)

	// display a new form by reset the form stage
	if activitiesFormCallback.CreationMode || activitiesFormCallback.formGroup.HasSuppressButtonBeenPressed {
		activitiesFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(activitiesFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ActivitiesFormCallback(
			nil,
			activitiesFormCallback.probe,
			newFormGroup,
		)
		activities := new(models.Activities)
		FillUpForm(activities, newFormGroup, activitiesFormCallback.probe)
		activitiesFormCallback.probe.formStage.Commit()
	}

	activitiesFormCallback.probe.ux_tree()
}
func __gong__New__ArchitectureFormCallback(
	architecture *models.Architecture,
	probe *Probe,
	formGroup *form.FormGroup,
) (architectureFormCallback *ArchitectureFormCallback) {
	architectureFormCallback = new(ArchitectureFormCallback)
	architectureFormCallback.probe = probe
	architectureFormCallback.architecture = architecture
	architectureFormCallback.formGroup = formGroup

	architectureFormCallback.CreationMode = (architecture == nil)

	return
}

type ArchitectureFormCallback struct {
	architecture *models.Architecture

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (architectureFormCallback *ArchitectureFormCallback) OnSave() {
	architectureFormCallback.probe.stageOfInterest.Lock()
	defer architectureFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("ArchitectureFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	architectureFormCallback.probe.formStage.Checkout()

	if architectureFormCallback.architecture == nil {
		architectureFormCallback.architecture = new(models.Architecture).Stage(architectureFormCallback.probe.stageOfInterest)
	}
	architecture_ := architectureFormCallback.architecture
	_ = architecture_

	for _, formDiv := range architectureFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(architecture_.Name), formDiv)
		case "StateMachines":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.StateMachine](architectureFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.StateMachine, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.StateMachine)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					architectureFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.StateMachine](architectureFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			architecture_.StateMachines = instanceSlice
			architectureFormCallback.probe.UpdateSliceOfPointersCallback(architecture_, "StateMachines", &architecture_.StateMachines)

		case "Roles":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Role](architectureFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Role, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Role)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					architectureFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Role](architectureFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			architecture_.Roles = instanceSlice
			architectureFormCallback.probe.UpdateSliceOfPointersCallback(architecture_, "Roles", &architecture_.Roles)

		case "NbPixPerCharacter":
			FormDivBasicFieldToField(&(architecture_.NbPixPerCharacter), formDiv)
		}
	}

	// manage the suppress operation
	if architectureFormCallback.formGroup.HasSuppressButtonBeenPressed {
		architecture_.Unstage(architectureFormCallback.probe.stageOfInterest)
	}

	architectureFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Architecture](
		architectureFormCallback.probe,
	)

	// display a new form by reset the form stage
	if architectureFormCallback.CreationMode || architectureFormCallback.formGroup.HasSuppressButtonBeenPressed {
		architectureFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(architectureFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ArchitectureFormCallback(
			nil,
			architectureFormCallback.probe,
			newFormGroup,
		)
		architecture := new(models.Architecture)
		FillUpForm(architecture, newFormGroup, architectureFormCallback.probe)
		architectureFormCallback.probe.formStage.Commit()
	}

	architectureFormCallback.probe.ux_tree()
}
func __gong__New__DiagramFormCallback(
	diagram *models.Diagram,
	probe *Probe,
	formGroup *form.FormGroup,
) (diagramFormCallback *DiagramFormCallback) {
	diagramFormCallback = new(DiagramFormCallback)
	diagramFormCallback.probe = probe
	diagramFormCallback.diagram = diagram
	diagramFormCallback.formGroup = formGroup

	diagramFormCallback.CreationMode = (diagram == nil)

	return
}

type DiagramFormCallback struct {
	diagram *models.Diagram

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (diagramFormCallback *DiagramFormCallback) OnSave() {
	diagramFormCallback.probe.stageOfInterest.Lock()
	defer diagramFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("DiagramFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	diagramFormCallback.probe.formStage.Checkout()

	if diagramFormCallback.diagram == nil {
		diagramFormCallback.diagram = new(models.Diagram).Stage(diagramFormCallback.probe.stageOfInterest)
	}
	diagram_ := diagramFormCallback.diagram
	_ = diagram_

	for _, formDiv := range diagramFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(diagram_.Name), formDiv)
		case "IsChecked":
			FormDivBasicFieldToField(&(diagram_.IsChecked), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(diagram_.IsExpanded), formDiv)
		case "IsEditable_":
			FormDivBasicFieldToField(&(diagram_.IsEditable_), formDiv)
		case "IsStatesNodeExpanded":
			FormDivBasicFieldToField(&(diagram_.IsStatesNodeExpanded), formDiv)
		case "IsNotesNodeExpanded":
			FormDivBasicFieldToField(&(diagram_.IsNotesNodeExpanded), formDiv)
		case "NotesWhoseNodeIsExpanded":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Note](diagramFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Note, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Note)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Note](diagramFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagram_.NotesWhoseNodeIsExpanded = instanceSlice
			diagramFormCallback.probe.UpdateSliceOfPointersCallback(diagram_, "NotesWhoseNodeIsExpanded", &diagram_.NotesWhoseNodeIsExpanded)

		case "State_Shapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.StateShape](diagramFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.StateShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.StateShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.StateShape](diagramFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagram_.State_Shapes = instanceSlice
			diagramFormCallback.probe.UpdateSliceOfPointersCallback(diagram_, "State_Shapes", &diagram_.State_Shapes)

		case "StatesWhoseNodeIsExpanded":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.State](diagramFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.State, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.State)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.State](diagramFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagram_.StatesWhoseNodeIsExpanded = instanceSlice
			diagramFormCallback.probe.UpdateSliceOfPointersCallback(diagram_, "StatesWhoseNodeIsExpanded", &diagram_.StatesWhoseNodeIsExpanded)

		case "Transition_Shapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Transition_Shape](diagramFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Transition_Shape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Transition_Shape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Transition_Shape](diagramFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagram_.Transition_Shapes = instanceSlice
			diagramFormCallback.probe.UpdateSliceOfPointersCallback(diagram_, "Transition_Shapes", &diagram_.Transition_Shapes)

		case "Note_Shapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.NoteShape](diagramFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.NoteShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.NoteShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.NoteShape](diagramFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagram_.Note_Shapes = instanceSlice
			diagramFormCallback.probe.UpdateSliceOfPointersCallback(diagram_, "Note_Shapes", &diagram_.Note_Shapes)

		case "NoteState_Shapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.NoteStateShape](diagramFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.NoteStateShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.NoteStateShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.NoteStateShape](diagramFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagram_.NoteState_Shapes = instanceSlice
			diagramFormCallback.probe.UpdateSliceOfPointersCallback(diagram_, "NoteState_Shapes", &diagram_.NoteState_Shapes)

		case "NoteTransition_Shapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.NoteTransitionShape](diagramFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.NoteTransitionShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.NoteTransitionShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.NoteTransitionShape](diagramFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagram_.NoteTransition_Shapes = instanceSlice
			diagramFormCallback.probe.UpdateSliceOfPointersCallback(diagram_, "NoteTransition_Shapes", &diagram_.NoteTransition_Shapes)

		case "Library:Diagrams":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Library instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Library instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Library](diagramFormCallback.probe.stageOfInterest)
			targetLibraryIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetLibraryIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Library instances and update their Diagrams slice
			for _library := range *models.GetGongstructInstancesSetFromPointerType[*models.Library](diagramFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(diagramFormCallback.probe.stageOfInterest, _library)
				
				// if Library is selected
				if targetLibraryIDs[id] {
					// ensure diagram_ is in _library.Diagrams
					found := false
					for _, _b := range _library.Diagrams {
						if _b == diagram_ {
							found = true
							break
						}
					}
					if !found {
						_library.Diagrams = append(_library.Diagrams, diagram_)
						diagramFormCallback.probe.UpdateSliceOfPointersCallback(_library, "Diagrams", &_library.Diagrams)
					}
				} else {
					// ensure diagram_ is NOT in _library.Diagrams
					idx := slices.Index(_library.Diagrams, diagram_)
					if idx != -1 {
						_library.Diagrams = slices.Delete(_library.Diagrams, idx, idx+1)
						diagramFormCallback.probe.UpdateSliceOfPointersCallback(_library, "Diagrams", &_library.Diagrams)
					}
				}
			}
		case "State:Diagrams":
			// 1. Decode the AssociationStorage which contains the rowIDs of the State instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target State instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.State](diagramFormCallback.probe.stageOfInterest)
			targetStateIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetStateIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all State instances and update their Diagrams slice
			for _state := range *models.GetGongstructInstancesSetFromPointerType[*models.State](diagramFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(diagramFormCallback.probe.stageOfInterest, _state)
				
				// if State is selected
				if targetStateIDs[id] {
					// ensure diagram_ is in _state.Diagrams
					found := false
					for _, _b := range _state.Diagrams {
						if _b == diagram_ {
							found = true
							break
						}
					}
					if !found {
						_state.Diagrams = append(_state.Diagrams, diagram_)
						diagramFormCallback.probe.UpdateSliceOfPointersCallback(_state, "Diagrams", &_state.Diagrams)
					}
				} else {
					// ensure diagram_ is NOT in _state.Diagrams
					idx := slices.Index(_state.Diagrams, diagram_)
					if idx != -1 {
						_state.Diagrams = slices.Delete(_state.Diagrams, idx, idx+1)
						diagramFormCallback.probe.UpdateSliceOfPointersCallback(_state, "Diagrams", &_state.Diagrams)
					}
				}
			}
		case "StateMachine:Diagrams":
			// 1. Decode the AssociationStorage which contains the rowIDs of the StateMachine instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target StateMachine instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.StateMachine](diagramFormCallback.probe.stageOfInterest)
			targetStateMachineIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetStateMachineIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all StateMachine instances and update their Diagrams slice
			for _statemachine := range *models.GetGongstructInstancesSetFromPointerType[*models.StateMachine](diagramFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(diagramFormCallback.probe.stageOfInterest, _statemachine)
				
				// if StateMachine is selected
				if targetStateMachineIDs[id] {
					// ensure diagram_ is in _statemachine.Diagrams
					found := false
					for _, _b := range _statemachine.Diagrams {
						if _b == diagram_ {
							found = true
							break
						}
					}
					if !found {
						_statemachine.Diagrams = append(_statemachine.Diagrams, diagram_)
						diagramFormCallback.probe.UpdateSliceOfPointersCallback(_statemachine, "Diagrams", &_statemachine.Diagrams)
					}
				} else {
					// ensure diagram_ is NOT in _statemachine.Diagrams
					idx := slices.Index(_statemachine.Diagrams, diagram_)
					if idx != -1 {
						_statemachine.Diagrams = slices.Delete(_statemachine.Diagrams, idx, idx+1)
						diagramFormCallback.probe.UpdateSliceOfPointersCallback(_statemachine, "Diagrams", &_statemachine.Diagrams)
					}
				}
			}
		case "Transition:Diagrams":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Transition instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Transition instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Transition](diagramFormCallback.probe.stageOfInterest)
			targetTransitionIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetTransitionIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Transition instances and update their Diagrams slice
			for _transition := range *models.GetGongstructInstancesSetFromPointerType[*models.Transition](diagramFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(diagramFormCallback.probe.stageOfInterest, _transition)
				
				// if Transition is selected
				if targetTransitionIDs[id] {
					// ensure diagram_ is in _transition.Diagrams
					found := false
					for _, _b := range _transition.Diagrams {
						if _b == diagram_ {
							found = true
							break
						}
					}
					if !found {
						_transition.Diagrams = append(_transition.Diagrams, diagram_)
						diagramFormCallback.probe.UpdateSliceOfPointersCallback(_transition, "Diagrams", &_transition.Diagrams)
					}
				} else {
					// ensure diagram_ is NOT in _transition.Diagrams
					idx := slices.Index(_transition.Diagrams, diagram_)
					if idx != -1 {
						_transition.Diagrams = slices.Delete(_transition.Diagrams, idx, idx+1)
						diagramFormCallback.probe.UpdateSliceOfPointersCallback(_transition, "Diagrams", &_transition.Diagrams)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if diagramFormCallback.formGroup.HasSuppressButtonBeenPressed {
		diagram_.Unstage(diagramFormCallback.probe.stageOfInterest)
	}

	diagramFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Diagram](
		diagramFormCallback.probe,
	)

	// display a new form by reset the form stage
	if diagramFormCallback.CreationMode || diagramFormCallback.formGroup.HasSuppressButtonBeenPressed {
		diagramFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(diagramFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__DiagramFormCallback(
			nil,
			diagramFormCallback.probe,
			newFormGroup,
		)
		diagram := new(models.Diagram)
		FillUpForm(diagram, newFormGroup, diagramFormCallback.probe)
		diagramFormCallback.probe.formStage.Commit()
	}

	diagramFormCallback.probe.ux_tree()
}
func __gong__New__GuardFormCallback(
	guard *models.Guard,
	probe *Probe,
	formGroup *form.FormGroup,
) (guardFormCallback *GuardFormCallback) {
	guardFormCallback = new(GuardFormCallback)
	guardFormCallback.probe = probe
	guardFormCallback.guard = guard
	guardFormCallback.formGroup = formGroup

	guardFormCallback.CreationMode = (guard == nil)

	return
}

type GuardFormCallback struct {
	guard *models.Guard

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (guardFormCallback *GuardFormCallback) OnSave() {
	guardFormCallback.probe.stageOfInterest.Lock()
	defer guardFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("GuardFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	guardFormCallback.probe.formStage.Checkout()

	if guardFormCallback.guard == nil {
		guardFormCallback.guard = new(models.Guard).Stage(guardFormCallback.probe.stageOfInterest)
	}
	guard_ := guardFormCallback.guard
	_ = guard_

	for _, formDiv := range guardFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(guard_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if guardFormCallback.formGroup.HasSuppressButtonBeenPressed {
		guard_.Unstage(guardFormCallback.probe.stageOfInterest)
	}

	guardFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Guard](
		guardFormCallback.probe,
	)

	// display a new form by reset the form stage
	if guardFormCallback.CreationMode || guardFormCallback.formGroup.HasSuppressButtonBeenPressed {
		guardFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(guardFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__GuardFormCallback(
			nil,
			guardFormCallback.probe,
			newFormGroup,
		)
		guard := new(models.Guard)
		FillUpForm(guard, newFormGroup, guardFormCallback.probe)
		guardFormCallback.probe.formStage.Commit()
	}

	guardFormCallback.probe.ux_tree()
}
func __gong__New__KillFormCallback(
	kill *models.Kill,
	probe *Probe,
	formGroup *form.FormGroup,
) (killFormCallback *KillFormCallback) {
	killFormCallback = new(KillFormCallback)
	killFormCallback.probe = probe
	killFormCallback.kill = kill
	killFormCallback.formGroup = formGroup

	killFormCallback.CreationMode = (kill == nil)

	return
}

type KillFormCallback struct {
	kill *models.Kill

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (killFormCallback *KillFormCallback) OnSave() {
	killFormCallback.probe.stageOfInterest.Lock()
	defer killFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("KillFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	killFormCallback.probe.formStage.Checkout()

	if killFormCallback.kill == nil {
		killFormCallback.kill = new(models.Kill).Stage(killFormCallback.probe.stageOfInterest)
	}
	kill_ := killFormCallback.kill
	_ = kill_

	for _, formDiv := range killFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(kill_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if killFormCallback.formGroup.HasSuppressButtonBeenPressed {
		kill_.Unstage(killFormCallback.probe.stageOfInterest)
	}

	killFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Kill](
		killFormCallback.probe,
	)

	// display a new form by reset the form stage
	if killFormCallback.CreationMode || killFormCallback.formGroup.HasSuppressButtonBeenPressed {
		killFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(killFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__KillFormCallback(
			nil,
			killFormCallback.probe,
			newFormGroup,
		)
		kill := new(models.Kill)
		FillUpForm(kill, newFormGroup, killFormCallback.probe)
		killFormCallback.probe.formStage.Commit()
	}

	killFormCallback.probe.ux_tree()
}
func __gong__New__LibraryFormCallback(
	library *models.Library,
	probe *Probe,
	formGroup *form.FormGroup,
) (libraryFormCallback *LibraryFormCallback) {
	libraryFormCallback = new(LibraryFormCallback)
	libraryFormCallback.probe = probe
	libraryFormCallback.library = library
	libraryFormCallback.formGroup = formGroup

	libraryFormCallback.CreationMode = (library == nil)

	return
}

type LibraryFormCallback struct {
	library *models.Library

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (libraryFormCallback *LibraryFormCallback) OnSave() {
	libraryFormCallback.probe.stageOfInterest.Lock()
	defer libraryFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("LibraryFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	libraryFormCallback.probe.formStage.Checkout()

	if libraryFormCallback.library == nil {
		libraryFormCallback.library = new(models.Library).Stage(libraryFormCallback.probe.stageOfInterest)
	}
	library_ := libraryFormCallback.library
	_ = library_

	for _, formDiv := range libraryFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(library_.Name), formDiv)
		case "SubLibraries":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Library](libraryFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Library, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Library)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					libraryFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Library](libraryFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			library_.SubLibraries = instanceSlice
			libraryFormCallback.probe.UpdateSliceOfPointersCallback(library_, "SubLibraries", &library_.SubLibraries)

		case "NbPixPerCharacter":
			FormDivBasicFieldToField(&(library_.NbPixPerCharacter), formDiv)
		case "LogoSVGFile":
			FormDivBasicFieldToField(&(library_.LogoSVGFile), formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(library_.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(library_.IsExpanded), formDiv)
		case "LayoutDirection":
			FormDivEnumIntFieldToField(&(library_.LayoutDirection), formDiv)
		case "IsRootLibrary":
			FormDivBasicFieldToField(&(library_.IsRootLibrary), formDiv)
		case "Diagrams":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Diagram](libraryFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Diagram, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Diagram)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					libraryFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Diagram](libraryFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			library_.Diagrams = instanceSlice
			libraryFormCallback.probe.UpdateSliceOfPointersCallback(library_, "Diagrams", &library_.Diagrams)

		case "RootStateMachines":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.StateMachine](libraryFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.StateMachine, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.StateMachine)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					libraryFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.StateMachine](libraryFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			library_.RootStateMachines = instanceSlice
			libraryFormCallback.probe.UpdateSliceOfPointersCallback(library_, "RootStateMachines", &library_.RootStateMachines)

		case "IsStateMachinesNodeExpanded":
			FormDivBasicFieldToField(&(library_.IsStateMachinesNodeExpanded), formDiv)
		case "StateMachinesWhoseNodeIsExpanded":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.StateMachine](libraryFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.StateMachine, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.StateMachine)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					libraryFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.StateMachine](libraryFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			library_.StateMachinesWhoseNodeIsExpanded = instanceSlice
			libraryFormCallback.probe.UpdateSliceOfPointersCallback(library_, "StateMachinesWhoseNodeIsExpanded", &library_.StateMachinesWhoseNodeIsExpanded)

		case "RootNotes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Note](libraryFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Note, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Note)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					libraryFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Note](libraryFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			library_.RootNotes = instanceSlice
			libraryFormCallback.probe.UpdateSliceOfPointersCallback(library_, "RootNotes", &library_.RootNotes)

		case "IsNotesNodeExpanded":
			FormDivBasicFieldToField(&(library_.IsNotesNodeExpanded), formDiv)
		case "NotesWhoseNodeIsExpanded":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Note](libraryFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Note, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Note)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					libraryFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Note](libraryFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			library_.NotesWhoseNodeIsExpanded = instanceSlice
			libraryFormCallback.probe.UpdateSliceOfPointersCallback(library_, "NotesWhoseNodeIsExpanded", &library_.NotesWhoseNodeIsExpanded)

		case "IsSubLibrariesNodeExpanded":
			FormDivBasicFieldToField(&(library_.IsSubLibrariesNodeExpanded), formDiv)
		case "SubLibrariesWhoseNodeIsExpanded":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Library](libraryFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Library, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Library)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					libraryFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Library](libraryFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			library_.SubLibrariesWhoseNodeIsExpanded = instanceSlice
			libraryFormCallback.probe.UpdateSliceOfPointersCallback(library_, "SubLibrariesWhoseNodeIsExpanded", &library_.SubLibrariesWhoseNodeIsExpanded)

		case "IsExpandedTmp":
			FormDivBasicFieldToField(&(library_.IsExpandedTmp), formDiv)
		case "Library:SubLibraries":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Library instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Library instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Library](libraryFormCallback.probe.stageOfInterest)
			targetLibraryIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetLibraryIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Library instances and update their SubLibraries slice
			for _library := range *models.GetGongstructInstancesSetFromPointerType[*models.Library](libraryFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(libraryFormCallback.probe.stageOfInterest, _library)
				
				// if Library is selected
				if targetLibraryIDs[id] {
					// ensure library_ is in _library.SubLibraries
					found := false
					for _, _b := range _library.SubLibraries {
						if _b == library_ {
							found = true
							break
						}
					}
					if !found {
						_library.SubLibraries = append(_library.SubLibraries, library_)
						libraryFormCallback.probe.UpdateSliceOfPointersCallback(_library, "SubLibraries", &_library.SubLibraries)
					}
				} else {
					// ensure library_ is NOT in _library.SubLibraries
					idx := slices.Index(_library.SubLibraries, library_)
					if idx != -1 {
						_library.SubLibraries = slices.Delete(_library.SubLibraries, idx, idx+1)
						libraryFormCallback.probe.UpdateSliceOfPointersCallback(_library, "SubLibraries", &_library.SubLibraries)
					}
				}
			}
		case "Library:SubLibrariesWhoseNodeIsExpanded":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Library instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Library instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Library](libraryFormCallback.probe.stageOfInterest)
			targetLibraryIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetLibraryIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Library instances and update their SubLibrariesWhoseNodeIsExpanded slice
			for _library := range *models.GetGongstructInstancesSetFromPointerType[*models.Library](libraryFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(libraryFormCallback.probe.stageOfInterest, _library)
				
				// if Library is selected
				if targetLibraryIDs[id] {
					// ensure library_ is in _library.SubLibrariesWhoseNodeIsExpanded
					found := false
					for _, _b := range _library.SubLibrariesWhoseNodeIsExpanded {
						if _b == library_ {
							found = true
							break
						}
					}
					if !found {
						_library.SubLibrariesWhoseNodeIsExpanded = append(_library.SubLibrariesWhoseNodeIsExpanded, library_)
						libraryFormCallback.probe.UpdateSliceOfPointersCallback(_library, "SubLibrariesWhoseNodeIsExpanded", &_library.SubLibrariesWhoseNodeIsExpanded)
					}
				} else {
					// ensure library_ is NOT in _library.SubLibrariesWhoseNodeIsExpanded
					idx := slices.Index(_library.SubLibrariesWhoseNodeIsExpanded, library_)
					if idx != -1 {
						_library.SubLibrariesWhoseNodeIsExpanded = slices.Delete(_library.SubLibrariesWhoseNodeIsExpanded, idx, idx+1)
						libraryFormCallback.probe.UpdateSliceOfPointersCallback(_library, "SubLibrariesWhoseNodeIsExpanded", &_library.SubLibrariesWhoseNodeIsExpanded)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if libraryFormCallback.formGroup.HasSuppressButtonBeenPressed {
		library_.Unstage(libraryFormCallback.probe.stageOfInterest)
	}

	libraryFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Library](
		libraryFormCallback.probe,
	)

	// display a new form by reset the form stage
	if libraryFormCallback.CreationMode || libraryFormCallback.formGroup.HasSuppressButtonBeenPressed {
		libraryFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(libraryFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__LibraryFormCallback(
			nil,
			libraryFormCallback.probe,
			newFormGroup,
		)
		library := new(models.Library)
		FillUpForm(library, newFormGroup, libraryFormCallback.probe)
		libraryFormCallback.probe.formStage.Commit()
	}

	libraryFormCallback.probe.ux_tree()
}
func __gong__New__MessageFormCallback(
	message *models.Message,
	probe *Probe,
	formGroup *form.FormGroup,
) (messageFormCallback *MessageFormCallback) {
	messageFormCallback = new(MessageFormCallback)
	messageFormCallback.probe = probe
	messageFormCallback.message = message
	messageFormCallback.formGroup = formGroup

	messageFormCallback.CreationMode = (message == nil)

	return
}

type MessageFormCallback struct {
	message *models.Message

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (messageFormCallback *MessageFormCallback) OnSave() {
	messageFormCallback.probe.stageOfInterest.Lock()
	defer messageFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("MessageFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	messageFormCallback.probe.formStage.Checkout()

	if messageFormCallback.message == nil {
		messageFormCallback.message = new(models.Message).Stage(messageFormCallback.probe.stageOfInterest)
	}
	message_ := messageFormCallback.message
	_ = message_

	for _, formDiv := range messageFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(message_.Name), formDiv)
		case "IsSelected":
			FormDivBasicFieldToField(&(message_.IsSelected), formDiv)
		case "MessageType":
			FormDivSelectFieldToField(&(message_.MessageType), messageFormCallback.probe.stageOfInterest, formDiv)
		case "OriginTransition":
			FormDivSelectFieldToField(&(message_.OriginTransition), messageFormCallback.probe.stageOfInterest, formDiv)
		case "Object:Messages":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Object instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Object instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Object](messageFormCallback.probe.stageOfInterest)
			targetObjectIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetObjectIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Object instances and update their Messages slice
			for _object := range *models.GetGongstructInstancesSetFromPointerType[*models.Object](messageFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(messageFormCallback.probe.stageOfInterest, _object)
				
				// if Object is selected
				if targetObjectIDs[id] {
					// ensure message_ is in _object.Messages
					found := false
					for _, _b := range _object.Messages {
						if _b == message_ {
							found = true
							break
						}
					}
					if !found {
						_object.Messages = append(_object.Messages, message_)
						messageFormCallback.probe.UpdateSliceOfPointersCallback(_object, "Messages", &_object.Messages)
					}
				} else {
					// ensure message_ is NOT in _object.Messages
					idx := slices.Index(_object.Messages, message_)
					if idx != -1 {
						_object.Messages = slices.Delete(_object.Messages, idx, idx+1)
						messageFormCallback.probe.UpdateSliceOfPointersCallback(_object, "Messages", &_object.Messages)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if messageFormCallback.formGroup.HasSuppressButtonBeenPressed {
		message_.Unstage(messageFormCallback.probe.stageOfInterest)
	}

	messageFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Message](
		messageFormCallback.probe,
	)

	// display a new form by reset the form stage
	if messageFormCallback.CreationMode || messageFormCallback.formGroup.HasSuppressButtonBeenPressed {
		messageFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(messageFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__MessageFormCallback(
			nil,
			messageFormCallback.probe,
			newFormGroup,
		)
		message := new(models.Message)
		FillUpForm(message, newFormGroup, messageFormCallback.probe)
		messageFormCallback.probe.formStage.Commit()
	}

	messageFormCallback.probe.ux_tree()
}
func __gong__New__MessageTypeFormCallback(
	messagetype *models.MessageType,
	probe *Probe,
	formGroup *form.FormGroup,
) (messagetypeFormCallback *MessageTypeFormCallback) {
	messagetypeFormCallback = new(MessageTypeFormCallback)
	messagetypeFormCallback.probe = probe
	messagetypeFormCallback.messagetype = messagetype
	messagetypeFormCallback.formGroup = formGroup

	messagetypeFormCallback.CreationMode = (messagetype == nil)

	return
}

type MessageTypeFormCallback struct {
	messagetype *models.MessageType

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (messagetypeFormCallback *MessageTypeFormCallback) OnSave() {
	messagetypeFormCallback.probe.stageOfInterest.Lock()
	defer messagetypeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("MessageTypeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	messagetypeFormCallback.probe.formStage.Checkout()

	if messagetypeFormCallback.messagetype == nil {
		messagetypeFormCallback.messagetype = new(models.MessageType).Stage(messagetypeFormCallback.probe.stageOfInterest)
	}
	messagetype_ := messagetypeFormCallback.messagetype
	_ = messagetype_

	for _, formDiv := range messagetypeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(messagetype_.Name), formDiv)
		case "Description":
			FormDivBasicFieldToField(&(messagetype_.Description), formDiv)
		case "Transition:GeneratedMessages":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Transition instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Transition instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Transition](messagetypeFormCallback.probe.stageOfInterest)
			targetTransitionIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetTransitionIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Transition instances and update their GeneratedMessages slice
			for _transition := range *models.GetGongstructInstancesSetFromPointerType[*models.Transition](messagetypeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(messagetypeFormCallback.probe.stageOfInterest, _transition)
				
				// if Transition is selected
				if targetTransitionIDs[id] {
					// ensure messagetype_ is in _transition.GeneratedMessages
					found := false
					for _, _b := range _transition.GeneratedMessages {
						if _b == messagetype_ {
							found = true
							break
						}
					}
					if !found {
						_transition.GeneratedMessages = append(_transition.GeneratedMessages, messagetype_)
						messagetypeFormCallback.probe.UpdateSliceOfPointersCallback(_transition, "GeneratedMessages", &_transition.GeneratedMessages)
					}
				} else {
					// ensure messagetype_ is NOT in _transition.GeneratedMessages
					idx := slices.Index(_transition.GeneratedMessages, messagetype_)
					if idx != -1 {
						_transition.GeneratedMessages = slices.Delete(_transition.GeneratedMessages, idx, idx+1)
						messagetypeFormCallback.probe.UpdateSliceOfPointersCallback(_transition, "GeneratedMessages", &_transition.GeneratedMessages)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if messagetypeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		messagetype_.Unstage(messagetypeFormCallback.probe.stageOfInterest)
	}

	messagetypeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.MessageType](
		messagetypeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if messagetypeFormCallback.CreationMode || messagetypeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		messagetypeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(messagetypeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__MessageTypeFormCallback(
			nil,
			messagetypeFormCallback.probe,
			newFormGroup,
		)
		messagetype := new(models.MessageType)
		FillUpForm(messagetype, newFormGroup, messagetypeFormCallback.probe)
		messagetypeFormCallback.probe.formStage.Commit()
	}

	messagetypeFormCallback.probe.ux_tree()
}
func __gong__New__NoteFormCallback(
	note *models.Note,
	probe *Probe,
	formGroup *form.FormGroup,
) (noteFormCallback *NoteFormCallback) {
	noteFormCallback = new(NoteFormCallback)
	noteFormCallback.probe = probe
	noteFormCallback.note = note
	noteFormCallback.formGroup = formGroup

	noteFormCallback.CreationMode = (note == nil)

	return
}

type NoteFormCallback struct {
	note *models.Note

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (noteFormCallback *NoteFormCallback) OnSave() {
	noteFormCallback.probe.stageOfInterest.Lock()
	defer noteFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("NoteFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	noteFormCallback.probe.formStage.Checkout()

	if noteFormCallback.note == nil {
		noteFormCallback.note = new(models.Note).Stage(noteFormCallback.probe.stageOfInterest)
	}
	note_ := noteFormCallback.note
	_ = note_

	for _, formDiv := range noteFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(note_.Name), formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(note_.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(note_.IsExpanded), formDiv)
		case "LayoutDirection":
			FormDivEnumIntFieldToField(&(note_.LayoutDirection), formDiv)
		case "States":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.State](noteFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.State, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.State)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					noteFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.State](noteFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			note_.States = instanceSlice
			noteFormCallback.probe.UpdateSliceOfPointersCallback(note_, "States", &note_.States)

		case "Transitions":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Transition](noteFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Transition, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Transition)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					noteFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Transition](noteFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			note_.Transitions = instanceSlice
			noteFormCallback.probe.UpdateSliceOfPointersCallback(note_, "Transitions", &note_.Transitions)

		case "Diagram:NotesWhoseNodeIsExpanded":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Diagram instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Diagram instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Diagram](noteFormCallback.probe.stageOfInterest)
			targetDiagramIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Diagram instances and update their NotesWhoseNodeIsExpanded slice
			for _diagram := range *models.GetGongstructInstancesSetFromPointerType[*models.Diagram](noteFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(noteFormCallback.probe.stageOfInterest, _diagram)
				
				// if Diagram is selected
				if targetDiagramIDs[id] {
					// ensure note_ is in _diagram.NotesWhoseNodeIsExpanded
					found := false
					for _, _b := range _diagram.NotesWhoseNodeIsExpanded {
						if _b == note_ {
							found = true
							break
						}
					}
					if !found {
						_diagram.NotesWhoseNodeIsExpanded = append(_diagram.NotesWhoseNodeIsExpanded, note_)
						noteFormCallback.probe.UpdateSliceOfPointersCallback(_diagram, "NotesWhoseNodeIsExpanded", &_diagram.NotesWhoseNodeIsExpanded)
					}
				} else {
					// ensure note_ is NOT in _diagram.NotesWhoseNodeIsExpanded
					idx := slices.Index(_diagram.NotesWhoseNodeIsExpanded, note_)
					if idx != -1 {
						_diagram.NotesWhoseNodeIsExpanded = slices.Delete(_diagram.NotesWhoseNodeIsExpanded, idx, idx+1)
						noteFormCallback.probe.UpdateSliceOfPointersCallback(_diagram, "NotesWhoseNodeIsExpanded", &_diagram.NotesWhoseNodeIsExpanded)
					}
				}
			}
		case "Library:RootNotes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Library instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Library instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Library](noteFormCallback.probe.stageOfInterest)
			targetLibraryIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetLibraryIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Library instances and update their RootNotes slice
			for _library := range *models.GetGongstructInstancesSetFromPointerType[*models.Library](noteFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(noteFormCallback.probe.stageOfInterest, _library)
				
				// if Library is selected
				if targetLibraryIDs[id] {
					// ensure note_ is in _library.RootNotes
					found := false
					for _, _b := range _library.RootNotes {
						if _b == note_ {
							found = true
							break
						}
					}
					if !found {
						_library.RootNotes = append(_library.RootNotes, note_)
						noteFormCallback.probe.UpdateSliceOfPointersCallback(_library, "RootNotes", &_library.RootNotes)
					}
				} else {
					// ensure note_ is NOT in _library.RootNotes
					idx := slices.Index(_library.RootNotes, note_)
					if idx != -1 {
						_library.RootNotes = slices.Delete(_library.RootNotes, idx, idx+1)
						noteFormCallback.probe.UpdateSliceOfPointersCallback(_library, "RootNotes", &_library.RootNotes)
					}
				}
			}
		case "Library:NotesWhoseNodeIsExpanded":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Library instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Library instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Library](noteFormCallback.probe.stageOfInterest)
			targetLibraryIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetLibraryIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Library instances and update their NotesWhoseNodeIsExpanded slice
			for _library := range *models.GetGongstructInstancesSetFromPointerType[*models.Library](noteFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(noteFormCallback.probe.stageOfInterest, _library)
				
				// if Library is selected
				if targetLibraryIDs[id] {
					// ensure note_ is in _library.NotesWhoseNodeIsExpanded
					found := false
					for _, _b := range _library.NotesWhoseNodeIsExpanded {
						if _b == note_ {
							found = true
							break
						}
					}
					if !found {
						_library.NotesWhoseNodeIsExpanded = append(_library.NotesWhoseNodeIsExpanded, note_)
						noteFormCallback.probe.UpdateSliceOfPointersCallback(_library, "NotesWhoseNodeIsExpanded", &_library.NotesWhoseNodeIsExpanded)
					}
				} else {
					// ensure note_ is NOT in _library.NotesWhoseNodeIsExpanded
					idx := slices.Index(_library.NotesWhoseNodeIsExpanded, note_)
					if idx != -1 {
						_library.NotesWhoseNodeIsExpanded = slices.Delete(_library.NotesWhoseNodeIsExpanded, idx, idx+1)
						noteFormCallback.probe.UpdateSliceOfPointersCallback(_library, "NotesWhoseNodeIsExpanded", &_library.NotesWhoseNodeIsExpanded)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if noteFormCallback.formGroup.HasSuppressButtonBeenPressed {
		note_.Unstage(noteFormCallback.probe.stageOfInterest)
	}

	noteFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Note](
		noteFormCallback.probe,
	)

	// display a new form by reset the form stage
	if noteFormCallback.CreationMode || noteFormCallback.formGroup.HasSuppressButtonBeenPressed {
		noteFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(noteFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__NoteFormCallback(
			nil,
			noteFormCallback.probe,
			newFormGroup,
		)
		note := new(models.Note)
		FillUpForm(note, newFormGroup, noteFormCallback.probe)
		noteFormCallback.probe.formStage.Commit()
	}

	noteFormCallback.probe.ux_tree()
}
func __gong__New__NoteShapeFormCallback(
	noteshape *models.NoteShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (noteshapeFormCallback *NoteShapeFormCallback) {
	noteshapeFormCallback = new(NoteShapeFormCallback)
	noteshapeFormCallback.probe = probe
	noteshapeFormCallback.noteshape = noteshape
	noteshapeFormCallback.formGroup = formGroup

	noteshapeFormCallback.CreationMode = (noteshape == nil)

	return
}

type NoteShapeFormCallback struct {
	noteshape *models.NoteShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (noteshapeFormCallback *NoteShapeFormCallback) OnSave() {
	noteshapeFormCallback.probe.stageOfInterest.Lock()
	defer noteshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("NoteShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	noteshapeFormCallback.probe.formStage.Checkout()

	if noteshapeFormCallback.noteshape == nil {
		noteshapeFormCallback.noteshape = new(models.NoteShape).Stage(noteshapeFormCallback.probe.stageOfInterest)
	}
	noteshape_ := noteshapeFormCallback.noteshape
	_ = noteshape_

	for _, formDiv := range noteshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(noteshape_.Name), formDiv)
		case "Note":
			FormDivSelectFieldToField(&(noteshape_.Note), noteshapeFormCallback.probe.stageOfInterest, formDiv)
		case "OverideLayoutDirection":
			FormDivBasicFieldToField(&(noteshape_.OverideLayoutDirection), formDiv)
		case "LayoutDirection":
			FormDivEnumIntFieldToField(&(noteshape_.LayoutDirection), formDiv)
		case "X":
			FormDivBasicFieldToField(&(noteshape_.X), formDiv)
		case "Y":
			FormDivBasicFieldToField(&(noteshape_.Y), formDiv)
		case "Width":
			FormDivBasicFieldToField(&(noteshape_.Width), formDiv)
		case "Height":
			FormDivBasicFieldToField(&(noteshape_.Height), formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&(noteshape_.IsHidden), formDiv)
		case "Diagram:Note_Shapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Diagram instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Diagram instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Diagram](noteshapeFormCallback.probe.stageOfInterest)
			targetDiagramIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Diagram instances and update their Note_Shapes slice
			for _diagram := range *models.GetGongstructInstancesSetFromPointerType[*models.Diagram](noteshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(noteshapeFormCallback.probe.stageOfInterest, _diagram)
				
				// if Diagram is selected
				if targetDiagramIDs[id] {
					// ensure noteshape_ is in _diagram.Note_Shapes
					found := false
					for _, _b := range _diagram.Note_Shapes {
						if _b == noteshape_ {
							found = true
							break
						}
					}
					if !found {
						_diagram.Note_Shapes = append(_diagram.Note_Shapes, noteshape_)
						noteshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagram, "Note_Shapes", &_diagram.Note_Shapes)
					}
				} else {
					// ensure noteshape_ is NOT in _diagram.Note_Shapes
					idx := slices.Index(_diagram.Note_Shapes, noteshape_)
					if idx != -1 {
						_diagram.Note_Shapes = slices.Delete(_diagram.Note_Shapes, idx, idx+1)
						noteshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagram, "Note_Shapes", &_diagram.Note_Shapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if noteshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		noteshape_.Unstage(noteshapeFormCallback.probe.stageOfInterest)
	}

	noteshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.NoteShape](
		noteshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if noteshapeFormCallback.CreationMode || noteshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		noteshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(noteshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__NoteShapeFormCallback(
			nil,
			noteshapeFormCallback.probe,
			newFormGroup,
		)
		noteshape := new(models.NoteShape)
		FillUpForm(noteshape, newFormGroup, noteshapeFormCallback.probe)
		noteshapeFormCallback.probe.formStage.Commit()
	}

	noteshapeFormCallback.probe.ux_tree()
}
func __gong__New__NoteStateShapeFormCallback(
	notestateshape *models.NoteStateShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (notestateshapeFormCallback *NoteStateShapeFormCallback) {
	notestateshapeFormCallback = new(NoteStateShapeFormCallback)
	notestateshapeFormCallback.probe = probe
	notestateshapeFormCallback.notestateshape = notestateshape
	notestateshapeFormCallback.formGroup = formGroup

	notestateshapeFormCallback.CreationMode = (notestateshape == nil)

	return
}

type NoteStateShapeFormCallback struct {
	notestateshape *models.NoteStateShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (notestateshapeFormCallback *NoteStateShapeFormCallback) OnSave() {
	notestateshapeFormCallback.probe.stageOfInterest.Lock()
	defer notestateshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("NoteStateShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	notestateshapeFormCallback.probe.formStage.Checkout()

	if notestateshapeFormCallback.notestateshape == nil {
		notestateshapeFormCallback.notestateshape = new(models.NoteStateShape).Stage(notestateshapeFormCallback.probe.stageOfInterest)
	}
	notestateshape_ := notestateshapeFormCallback.notestateshape
	_ = notestateshape_

	for _, formDiv := range notestateshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(notestateshape_.Name), formDiv)
		case "Note":
			FormDivSelectFieldToField(&(notestateshape_.Note), notestateshapeFormCallback.probe.stageOfInterest, formDiv)
		case "State":
			FormDivSelectFieldToField(&(notestateshape_.State), notestateshapeFormCallback.probe.stageOfInterest, formDiv)
		case "StartRatio":
			FormDivBasicFieldToField(&(notestateshape_.StartRatio), formDiv)
		case "EndRatio":
			FormDivBasicFieldToField(&(notestateshape_.EndRatio), formDiv)
		case "StartOrientation":
			FormDivEnumStringFieldToField(&(notestateshape_.StartOrientation), formDiv)
		case "EndOrientation":
			FormDivEnumStringFieldToField(&(notestateshape_.EndOrientation), formDiv)
		case "CornerOffsetRatio":
			FormDivBasicFieldToField(&(notestateshape_.CornerOffsetRatio), formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&(notestateshape_.IsHidden), formDiv)
		case "Diagram:NoteState_Shapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Diagram instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Diagram instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Diagram](notestateshapeFormCallback.probe.stageOfInterest)
			targetDiagramIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Diagram instances and update their NoteState_Shapes slice
			for _diagram := range *models.GetGongstructInstancesSetFromPointerType[*models.Diagram](notestateshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(notestateshapeFormCallback.probe.stageOfInterest, _diagram)
				
				// if Diagram is selected
				if targetDiagramIDs[id] {
					// ensure notestateshape_ is in _diagram.NoteState_Shapes
					found := false
					for _, _b := range _diagram.NoteState_Shapes {
						if _b == notestateshape_ {
							found = true
							break
						}
					}
					if !found {
						_diagram.NoteState_Shapes = append(_diagram.NoteState_Shapes, notestateshape_)
						notestateshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagram, "NoteState_Shapes", &_diagram.NoteState_Shapes)
					}
				} else {
					// ensure notestateshape_ is NOT in _diagram.NoteState_Shapes
					idx := slices.Index(_diagram.NoteState_Shapes, notestateshape_)
					if idx != -1 {
						_diagram.NoteState_Shapes = slices.Delete(_diagram.NoteState_Shapes, idx, idx+1)
						notestateshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagram, "NoteState_Shapes", &_diagram.NoteState_Shapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if notestateshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		notestateshape_.Unstage(notestateshapeFormCallback.probe.stageOfInterest)
	}

	notestateshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.NoteStateShape](
		notestateshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if notestateshapeFormCallback.CreationMode || notestateshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		notestateshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(notestateshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__NoteStateShapeFormCallback(
			nil,
			notestateshapeFormCallback.probe,
			newFormGroup,
		)
		notestateshape := new(models.NoteStateShape)
		FillUpForm(notestateshape, newFormGroup, notestateshapeFormCallback.probe)
		notestateshapeFormCallback.probe.formStage.Commit()
	}

	notestateshapeFormCallback.probe.ux_tree()
}
func __gong__New__NoteTransitionShapeFormCallback(
	notetransitionshape *models.NoteTransitionShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (notetransitionshapeFormCallback *NoteTransitionShapeFormCallback) {
	notetransitionshapeFormCallback = new(NoteTransitionShapeFormCallback)
	notetransitionshapeFormCallback.probe = probe
	notetransitionshapeFormCallback.notetransitionshape = notetransitionshape
	notetransitionshapeFormCallback.formGroup = formGroup

	notetransitionshapeFormCallback.CreationMode = (notetransitionshape == nil)

	return
}

type NoteTransitionShapeFormCallback struct {
	notetransitionshape *models.NoteTransitionShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (notetransitionshapeFormCallback *NoteTransitionShapeFormCallback) OnSave() {
	notetransitionshapeFormCallback.probe.stageOfInterest.Lock()
	defer notetransitionshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("NoteTransitionShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	notetransitionshapeFormCallback.probe.formStage.Checkout()

	if notetransitionshapeFormCallback.notetransitionshape == nil {
		notetransitionshapeFormCallback.notetransitionshape = new(models.NoteTransitionShape).Stage(notetransitionshapeFormCallback.probe.stageOfInterest)
	}
	notetransitionshape_ := notetransitionshapeFormCallback.notetransitionshape
	_ = notetransitionshape_

	for _, formDiv := range notetransitionshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(notetransitionshape_.Name), formDiv)
		case "Note":
			FormDivSelectFieldToField(&(notetransitionshape_.Note), notetransitionshapeFormCallback.probe.stageOfInterest, formDiv)
		case "Transition":
			FormDivSelectFieldToField(&(notetransitionshape_.Transition), notetransitionshapeFormCallback.probe.stageOfInterest, formDiv)
		case "StartRatio":
			FormDivBasicFieldToField(&(notetransitionshape_.StartRatio), formDiv)
		case "EndRatio":
			FormDivBasicFieldToField(&(notetransitionshape_.EndRatio), formDiv)
		case "StartOrientation":
			FormDivEnumStringFieldToField(&(notetransitionshape_.StartOrientation), formDiv)
		case "EndOrientation":
			FormDivEnumStringFieldToField(&(notetransitionshape_.EndOrientation), formDiv)
		case "CornerOffsetRatio":
			FormDivBasicFieldToField(&(notetransitionshape_.CornerOffsetRatio), formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&(notetransitionshape_.IsHidden), formDiv)
		case "Diagram:NoteTransition_Shapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Diagram instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Diagram instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Diagram](notetransitionshapeFormCallback.probe.stageOfInterest)
			targetDiagramIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Diagram instances and update their NoteTransition_Shapes slice
			for _diagram := range *models.GetGongstructInstancesSetFromPointerType[*models.Diagram](notetransitionshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(notetransitionshapeFormCallback.probe.stageOfInterest, _diagram)
				
				// if Diagram is selected
				if targetDiagramIDs[id] {
					// ensure notetransitionshape_ is in _diagram.NoteTransition_Shapes
					found := false
					for _, _b := range _diagram.NoteTransition_Shapes {
						if _b == notetransitionshape_ {
							found = true
							break
						}
					}
					if !found {
						_diagram.NoteTransition_Shapes = append(_diagram.NoteTransition_Shapes, notetransitionshape_)
						notetransitionshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagram, "NoteTransition_Shapes", &_diagram.NoteTransition_Shapes)
					}
				} else {
					// ensure notetransitionshape_ is NOT in _diagram.NoteTransition_Shapes
					idx := slices.Index(_diagram.NoteTransition_Shapes, notetransitionshape_)
					if idx != -1 {
						_diagram.NoteTransition_Shapes = slices.Delete(_diagram.NoteTransition_Shapes, idx, idx+1)
						notetransitionshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagram, "NoteTransition_Shapes", &_diagram.NoteTransition_Shapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if notetransitionshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		notetransitionshape_.Unstage(notetransitionshapeFormCallback.probe.stageOfInterest)
	}

	notetransitionshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.NoteTransitionShape](
		notetransitionshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if notetransitionshapeFormCallback.CreationMode || notetransitionshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		notetransitionshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(notetransitionshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__NoteTransitionShapeFormCallback(
			nil,
			notetransitionshapeFormCallback.probe,
			newFormGroup,
		)
		notetransitionshape := new(models.NoteTransitionShape)
		FillUpForm(notetransitionshape, newFormGroup, notetransitionshapeFormCallback.probe)
		notetransitionshapeFormCallback.probe.formStage.Commit()
	}

	notetransitionshapeFormCallback.probe.ux_tree()
}
func __gong__New__ObjectFormCallback(
	object *models.Object,
	probe *Probe,
	formGroup *form.FormGroup,
) (objectFormCallback *ObjectFormCallback) {
	objectFormCallback = new(ObjectFormCallback)
	objectFormCallback.probe = probe
	objectFormCallback.object = object
	objectFormCallback.formGroup = formGroup

	objectFormCallback.CreationMode = (object == nil)

	return
}

type ObjectFormCallback struct {
	object *models.Object

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (objectFormCallback *ObjectFormCallback) OnSave() {
	objectFormCallback.probe.stageOfInterest.Lock()
	defer objectFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("ObjectFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	objectFormCallback.probe.formStage.Checkout()

	if objectFormCallback.object == nil {
		objectFormCallback.object = new(models.Object).Stage(objectFormCallback.probe.stageOfInterest)
	}
	object_ := objectFormCallback.object
	_ = object_

	for _, formDiv := range objectFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(object_.Name), formDiv)
		case "State":
			FormDivSelectFieldToField(&(object_.State), objectFormCallback.probe.stageOfInterest, formDiv)
		case "IsSelected":
			FormDivBasicFieldToField(&(object_.IsSelected), formDiv)
		case "Rank":
			FormDivBasicFieldToField(&(object_.Rank), formDiv)
		case "DOF":
			FormDivBasicFieldToField(&(object_.DOF), formDiv)
		case "Messages":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Message](objectFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Message, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Message)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					objectFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Message](objectFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			object_.Messages = instanceSlice
			objectFormCallback.probe.UpdateSliceOfPointersCallback(object_, "Messages", &object_.Messages)

		}
	}

	// manage the suppress operation
	if objectFormCallback.formGroup.HasSuppressButtonBeenPressed {
		object_.Unstage(objectFormCallback.probe.stageOfInterest)
	}

	objectFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Object](
		objectFormCallback.probe,
	)

	// display a new form by reset the form stage
	if objectFormCallback.CreationMode || objectFormCallback.formGroup.HasSuppressButtonBeenPressed {
		objectFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(objectFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ObjectFormCallback(
			nil,
			objectFormCallback.probe,
			newFormGroup,
		)
		object := new(models.Object)
		FillUpForm(object, newFormGroup, objectFormCallback.probe)
		objectFormCallback.probe.formStage.Commit()
	}

	objectFormCallback.probe.ux_tree()
}
func __gong__New__RoleFormCallback(
	role *models.Role,
	probe *Probe,
	formGroup *form.FormGroup,
) (roleFormCallback *RoleFormCallback) {
	roleFormCallback = new(RoleFormCallback)
	roleFormCallback.probe = probe
	roleFormCallback.role = role
	roleFormCallback.formGroup = formGroup

	roleFormCallback.CreationMode = (role == nil)

	return
}

type RoleFormCallback struct {
	role *models.Role

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (roleFormCallback *RoleFormCallback) OnSave() {
	roleFormCallback.probe.stageOfInterest.Lock()
	defer roleFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("RoleFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	roleFormCallback.probe.formStage.Checkout()

	if roleFormCallback.role == nil {
		roleFormCallback.role = new(models.Role).Stage(roleFormCallback.probe.stageOfInterest)
	}
	role_ := roleFormCallback.role
	_ = role_

	for _, formDiv := range roleFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(role_.Name), formDiv)
		case "Acronym":
			FormDivBasicFieldToField(&(role_.Acronym), formDiv)
		case "RolesWithSamePermissions":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Role](roleFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Role, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Role)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					roleFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Role](roleFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			role_.RolesWithSamePermissions = instanceSlice
			roleFormCallback.probe.UpdateSliceOfPointersCallback(role_, "RolesWithSamePermissions", &role_.RolesWithSamePermissions)

		case "Architecture:Roles":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Architecture instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Architecture instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Architecture](roleFormCallback.probe.stageOfInterest)
			targetArchitectureIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetArchitectureIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Architecture instances and update their Roles slice
			for _architecture := range *models.GetGongstructInstancesSetFromPointerType[*models.Architecture](roleFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(roleFormCallback.probe.stageOfInterest, _architecture)
				
				// if Architecture is selected
				if targetArchitectureIDs[id] {
					// ensure role_ is in _architecture.Roles
					found := false
					for _, _b := range _architecture.Roles {
						if _b == role_ {
							found = true
							break
						}
					}
					if !found {
						_architecture.Roles = append(_architecture.Roles, role_)
						roleFormCallback.probe.UpdateSliceOfPointersCallback(_architecture, "Roles", &_architecture.Roles)
					}
				} else {
					// ensure role_ is NOT in _architecture.Roles
					idx := slices.Index(_architecture.Roles, role_)
					if idx != -1 {
						_architecture.Roles = slices.Delete(_architecture.Roles, idx, idx+1)
						roleFormCallback.probe.UpdateSliceOfPointersCallback(_architecture, "Roles", &_architecture.Roles)
					}
				}
			}
		case "Role:RolesWithSamePermissions":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Role instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Role instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Role](roleFormCallback.probe.stageOfInterest)
			targetRoleIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetRoleIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Role instances and update their RolesWithSamePermissions slice
			for _role := range *models.GetGongstructInstancesSetFromPointerType[*models.Role](roleFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(roleFormCallback.probe.stageOfInterest, _role)
				
				// if Role is selected
				if targetRoleIDs[id] {
					// ensure role_ is in _role.RolesWithSamePermissions
					found := false
					for _, _b := range _role.RolesWithSamePermissions {
						if _b == role_ {
							found = true
							break
						}
					}
					if !found {
						_role.RolesWithSamePermissions = append(_role.RolesWithSamePermissions, role_)
						roleFormCallback.probe.UpdateSliceOfPointersCallback(_role, "RolesWithSamePermissions", &_role.RolesWithSamePermissions)
					}
				} else {
					// ensure role_ is NOT in _role.RolesWithSamePermissions
					idx := slices.Index(_role.RolesWithSamePermissions, role_)
					if idx != -1 {
						_role.RolesWithSamePermissions = slices.Delete(_role.RolesWithSamePermissions, idx, idx+1)
						roleFormCallback.probe.UpdateSliceOfPointersCallback(_role, "RolesWithSamePermissions", &_role.RolesWithSamePermissions)
					}
				}
			}
		case "Transition:RolesWithPermissions":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Transition instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Transition instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Transition](roleFormCallback.probe.stageOfInterest)
			targetTransitionIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetTransitionIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Transition instances and update their RolesWithPermissions slice
			for _transition := range *models.GetGongstructInstancesSetFromPointerType[*models.Transition](roleFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(roleFormCallback.probe.stageOfInterest, _transition)
				
				// if Transition is selected
				if targetTransitionIDs[id] {
					// ensure role_ is in _transition.RolesWithPermissions
					found := false
					for _, _b := range _transition.RolesWithPermissions {
						if _b == role_ {
							found = true
							break
						}
					}
					if !found {
						_transition.RolesWithPermissions = append(_transition.RolesWithPermissions, role_)
						roleFormCallback.probe.UpdateSliceOfPointersCallback(_transition, "RolesWithPermissions", &_transition.RolesWithPermissions)
					}
				} else {
					// ensure role_ is NOT in _transition.RolesWithPermissions
					idx := slices.Index(_transition.RolesWithPermissions, role_)
					if idx != -1 {
						_transition.RolesWithPermissions = slices.Delete(_transition.RolesWithPermissions, idx, idx+1)
						roleFormCallback.probe.UpdateSliceOfPointersCallback(_transition, "RolesWithPermissions", &_transition.RolesWithPermissions)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if roleFormCallback.formGroup.HasSuppressButtonBeenPressed {
		role_.Unstage(roleFormCallback.probe.stageOfInterest)
	}

	roleFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Role](
		roleFormCallback.probe,
	)

	// display a new form by reset the form stage
	if roleFormCallback.CreationMode || roleFormCallback.formGroup.HasSuppressButtonBeenPressed {
		roleFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(roleFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__RoleFormCallback(
			nil,
			roleFormCallback.probe,
			newFormGroup,
		)
		role := new(models.Role)
		FillUpForm(role, newFormGroup, roleFormCallback.probe)
		roleFormCallback.probe.formStage.Commit()
	}

	roleFormCallback.probe.ux_tree()
}
func __gong__New__StateFormCallback(
	state *models.State,
	probe *Probe,
	formGroup *form.FormGroup,
) (stateFormCallback *StateFormCallback) {
	stateFormCallback = new(StateFormCallback)
	stateFormCallback.probe = probe
	stateFormCallback.state = state
	stateFormCallback.formGroup = formGroup

	stateFormCallback.CreationMode = (state == nil)

	return
}

type StateFormCallback struct {
	state *models.State

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (stateFormCallback *StateFormCallback) OnSave() {
	stateFormCallback.probe.stageOfInterest.Lock()
	defer stateFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("StateFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	stateFormCallback.probe.formStage.Checkout()

	if stateFormCallback.state == nil {
		stateFormCallback.state = new(models.State).Stage(stateFormCallback.probe.stageOfInterest)
	}
	state_ := stateFormCallback.state
	_ = state_

	for _, formDiv := range stateFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(state_.Name), formDiv)
		case "Parent":
			FormDivSelectFieldToField(&(state_.Parent), stateFormCallback.probe.stageOfInterest, formDiv)
		case "IsDecisionNode":
			FormDivBasicFieldToField(&(state_.IsDecisionNode), formDiv)
		case "IsFictious":
			FormDivBasicFieldToField(&(state_.IsFictious), formDiv)
		case "IsEndState":
			FormDivBasicFieldToField(&(state_.IsEndState), formDiv)
		case "SubStates":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.State](stateFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.State, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.State)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					stateFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.State](stateFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			state_.SubStates = instanceSlice
			stateFormCallback.probe.UpdateSliceOfPointersCallback(state_, "SubStates", &state_.SubStates)

		case "Diagrams":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Diagram](stateFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Diagram, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Diagram)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					stateFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Diagram](stateFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			state_.Diagrams = instanceSlice
			stateFormCallback.probe.UpdateSliceOfPointersCallback(state_, "Diagrams", &state_.Diagrams)

		case "Entry":
			FormDivSelectFieldToField(&(state_.Entry), stateFormCallback.probe.stageOfInterest, formDiv)
		case "Activities":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Activities](stateFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Activities, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Activities)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					stateFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Activities](stateFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			state_.Activities = instanceSlice
			stateFormCallback.probe.UpdateSliceOfPointersCallback(state_, "Activities", &state_.Activities)

		case "Exit":
			FormDivSelectFieldToField(&(state_.Exit), stateFormCallback.probe.stageOfInterest, formDiv)
		case "Diagram:StatesWhoseNodeIsExpanded":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Diagram instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Diagram instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Diagram](stateFormCallback.probe.stageOfInterest)
			targetDiagramIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Diagram instances and update their StatesWhoseNodeIsExpanded slice
			for _diagram := range *models.GetGongstructInstancesSetFromPointerType[*models.Diagram](stateFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(stateFormCallback.probe.stageOfInterest, _diagram)
				
				// if Diagram is selected
				if targetDiagramIDs[id] {
					// ensure state_ is in _diagram.StatesWhoseNodeIsExpanded
					found := false
					for _, _b := range _diagram.StatesWhoseNodeIsExpanded {
						if _b == state_ {
							found = true
							break
						}
					}
					if !found {
						_diagram.StatesWhoseNodeIsExpanded = append(_diagram.StatesWhoseNodeIsExpanded, state_)
						stateFormCallback.probe.UpdateSliceOfPointersCallback(_diagram, "StatesWhoseNodeIsExpanded", &_diagram.StatesWhoseNodeIsExpanded)
					}
				} else {
					// ensure state_ is NOT in _diagram.StatesWhoseNodeIsExpanded
					idx := slices.Index(_diagram.StatesWhoseNodeIsExpanded, state_)
					if idx != -1 {
						_diagram.StatesWhoseNodeIsExpanded = slices.Delete(_diagram.StatesWhoseNodeIsExpanded, idx, idx+1)
						stateFormCallback.probe.UpdateSliceOfPointersCallback(_diagram, "StatesWhoseNodeIsExpanded", &_diagram.StatesWhoseNodeIsExpanded)
					}
				}
			}
		case "Note:States":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Note instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Note instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Note](stateFormCallback.probe.stageOfInterest)
			targetNoteIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetNoteIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Note instances and update their States slice
			for _note := range *models.GetGongstructInstancesSetFromPointerType[*models.Note](stateFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(stateFormCallback.probe.stageOfInterest, _note)
				
				// if Note is selected
				if targetNoteIDs[id] {
					// ensure state_ is in _note.States
					found := false
					for _, _b := range _note.States {
						if _b == state_ {
							found = true
							break
						}
					}
					if !found {
						_note.States = append(_note.States, state_)
						stateFormCallback.probe.UpdateSliceOfPointersCallback(_note, "States", &_note.States)
					}
				} else {
					// ensure state_ is NOT in _note.States
					idx := slices.Index(_note.States, state_)
					if idx != -1 {
						_note.States = slices.Delete(_note.States, idx, idx+1)
						stateFormCallback.probe.UpdateSliceOfPointersCallback(_note, "States", &_note.States)
					}
				}
			}
		case "State:SubStates":
			// 1. Decode the AssociationStorage which contains the rowIDs of the State instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target State instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.State](stateFormCallback.probe.stageOfInterest)
			targetStateIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetStateIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all State instances and update their SubStates slice
			for _state := range *models.GetGongstructInstancesSetFromPointerType[*models.State](stateFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(stateFormCallback.probe.stageOfInterest, _state)
				
				// if State is selected
				if targetStateIDs[id] {
					// ensure state_ is in _state.SubStates
					found := false
					for _, _b := range _state.SubStates {
						if _b == state_ {
							found = true
							break
						}
					}
					if !found {
						_state.SubStates = append(_state.SubStates, state_)
						stateFormCallback.probe.UpdateSliceOfPointersCallback(_state, "SubStates", &_state.SubStates)
					}
				} else {
					// ensure state_ is NOT in _state.SubStates
					idx := slices.Index(_state.SubStates, state_)
					if idx != -1 {
						_state.SubStates = slices.Delete(_state.SubStates, idx, idx+1)
						stateFormCallback.probe.UpdateSliceOfPointersCallback(_state, "SubStates", &_state.SubStates)
					}
				}
			}
		case "StateMachine:States":
			// 1. Decode the AssociationStorage which contains the rowIDs of the StateMachine instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target StateMachine instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.StateMachine](stateFormCallback.probe.stageOfInterest)
			targetStateMachineIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetStateMachineIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all StateMachine instances and update their States slice
			for _statemachine := range *models.GetGongstructInstancesSetFromPointerType[*models.StateMachine](stateFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(stateFormCallback.probe.stageOfInterest, _statemachine)
				
				// if StateMachine is selected
				if targetStateMachineIDs[id] {
					// ensure state_ is in _statemachine.States
					found := false
					for _, _b := range _statemachine.States {
						if _b == state_ {
							found = true
							break
						}
					}
					if !found {
						_statemachine.States = append(_statemachine.States, state_)
						stateFormCallback.probe.UpdateSliceOfPointersCallback(_statemachine, "States", &_statemachine.States)
					}
				} else {
					// ensure state_ is NOT in _statemachine.States
					idx := slices.Index(_statemachine.States, state_)
					if idx != -1 {
						_statemachine.States = slices.Delete(_statemachine.States, idx, idx+1)
						stateFormCallback.probe.UpdateSliceOfPointersCallback(_statemachine, "States", &_statemachine.States)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if stateFormCallback.formGroup.HasSuppressButtonBeenPressed {
		state_.Unstage(stateFormCallback.probe.stageOfInterest)
	}

	stateFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.State](
		stateFormCallback.probe,
	)

	// display a new form by reset the form stage
	if stateFormCallback.CreationMode || stateFormCallback.formGroup.HasSuppressButtonBeenPressed {
		stateFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(stateFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__StateFormCallback(
			nil,
			stateFormCallback.probe,
			newFormGroup,
		)
		state := new(models.State)
		FillUpForm(state, newFormGroup, stateFormCallback.probe)
		stateFormCallback.probe.formStage.Commit()
	}

	stateFormCallback.probe.ux_tree()
}
func __gong__New__StateMachineFormCallback(
	statemachine *models.StateMachine,
	probe *Probe,
	formGroup *form.FormGroup,
) (statemachineFormCallback *StateMachineFormCallback) {
	statemachineFormCallback = new(StateMachineFormCallback)
	statemachineFormCallback.probe = probe
	statemachineFormCallback.statemachine = statemachine
	statemachineFormCallback.formGroup = formGroup

	statemachineFormCallback.CreationMode = (statemachine == nil)

	return
}

type StateMachineFormCallback struct {
	statemachine *models.StateMachine

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (statemachineFormCallback *StateMachineFormCallback) OnSave() {
	statemachineFormCallback.probe.stageOfInterest.Lock()
	defer statemachineFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("StateMachineFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	statemachineFormCallback.probe.formStage.Checkout()

	if statemachineFormCallback.statemachine == nil {
		statemachineFormCallback.statemachine = new(models.StateMachine).Stage(statemachineFormCallback.probe.stageOfInterest)
	}
	statemachine_ := statemachineFormCallback.statemachine
	_ = statemachine_

	for _, formDiv := range statemachineFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(statemachine_.Name), formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(statemachine_.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(statemachine_.IsExpanded), formDiv)
		case "LayoutDirection":
			FormDivEnumIntFieldToField(&(statemachine_.LayoutDirection), formDiv)
		case "States":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.State](statemachineFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.State, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.State)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					statemachineFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.State](statemachineFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			statemachine_.States = instanceSlice
			statemachineFormCallback.probe.UpdateSliceOfPointersCallback(statemachine_, "States", &statemachine_.States)

		case "Diagrams":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Diagram](statemachineFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Diagram, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Diagram)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					statemachineFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Diagram](statemachineFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			statemachine_.Diagrams = instanceSlice
			statemachineFormCallback.probe.UpdateSliceOfPointersCallback(statemachine_, "Diagrams", &statemachine_.Diagrams)

		case "InitialState":
			FormDivSelectFieldToField(&(statemachine_.InitialState), statemachineFormCallback.probe.stageOfInterest, formDiv)
		case "Architecture:StateMachines":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Architecture instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Architecture instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Architecture](statemachineFormCallback.probe.stageOfInterest)
			targetArchitectureIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetArchitectureIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Architecture instances and update their StateMachines slice
			for _architecture := range *models.GetGongstructInstancesSetFromPointerType[*models.Architecture](statemachineFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(statemachineFormCallback.probe.stageOfInterest, _architecture)
				
				// if Architecture is selected
				if targetArchitectureIDs[id] {
					// ensure statemachine_ is in _architecture.StateMachines
					found := false
					for _, _b := range _architecture.StateMachines {
						if _b == statemachine_ {
							found = true
							break
						}
					}
					if !found {
						_architecture.StateMachines = append(_architecture.StateMachines, statemachine_)
						statemachineFormCallback.probe.UpdateSliceOfPointersCallback(_architecture, "StateMachines", &_architecture.StateMachines)
					}
				} else {
					// ensure statemachine_ is NOT in _architecture.StateMachines
					idx := slices.Index(_architecture.StateMachines, statemachine_)
					if idx != -1 {
						_architecture.StateMachines = slices.Delete(_architecture.StateMachines, idx, idx+1)
						statemachineFormCallback.probe.UpdateSliceOfPointersCallback(_architecture, "StateMachines", &_architecture.StateMachines)
					}
				}
			}
		case "Library:RootStateMachines":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Library instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Library instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Library](statemachineFormCallback.probe.stageOfInterest)
			targetLibraryIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetLibraryIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Library instances and update their RootStateMachines slice
			for _library := range *models.GetGongstructInstancesSetFromPointerType[*models.Library](statemachineFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(statemachineFormCallback.probe.stageOfInterest, _library)
				
				// if Library is selected
				if targetLibraryIDs[id] {
					// ensure statemachine_ is in _library.RootStateMachines
					found := false
					for _, _b := range _library.RootStateMachines {
						if _b == statemachine_ {
							found = true
							break
						}
					}
					if !found {
						_library.RootStateMachines = append(_library.RootStateMachines, statemachine_)
						statemachineFormCallback.probe.UpdateSliceOfPointersCallback(_library, "RootStateMachines", &_library.RootStateMachines)
					}
				} else {
					// ensure statemachine_ is NOT in _library.RootStateMachines
					idx := slices.Index(_library.RootStateMachines, statemachine_)
					if idx != -1 {
						_library.RootStateMachines = slices.Delete(_library.RootStateMachines, idx, idx+1)
						statemachineFormCallback.probe.UpdateSliceOfPointersCallback(_library, "RootStateMachines", &_library.RootStateMachines)
					}
				}
			}
		case "Library:StateMachinesWhoseNodeIsExpanded":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Library instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Library instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Library](statemachineFormCallback.probe.stageOfInterest)
			targetLibraryIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetLibraryIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Library instances and update their StateMachinesWhoseNodeIsExpanded slice
			for _library := range *models.GetGongstructInstancesSetFromPointerType[*models.Library](statemachineFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(statemachineFormCallback.probe.stageOfInterest, _library)
				
				// if Library is selected
				if targetLibraryIDs[id] {
					// ensure statemachine_ is in _library.StateMachinesWhoseNodeIsExpanded
					found := false
					for _, _b := range _library.StateMachinesWhoseNodeIsExpanded {
						if _b == statemachine_ {
							found = true
							break
						}
					}
					if !found {
						_library.StateMachinesWhoseNodeIsExpanded = append(_library.StateMachinesWhoseNodeIsExpanded, statemachine_)
						statemachineFormCallback.probe.UpdateSliceOfPointersCallback(_library, "StateMachinesWhoseNodeIsExpanded", &_library.StateMachinesWhoseNodeIsExpanded)
					}
				} else {
					// ensure statemachine_ is NOT in _library.StateMachinesWhoseNodeIsExpanded
					idx := slices.Index(_library.StateMachinesWhoseNodeIsExpanded, statemachine_)
					if idx != -1 {
						_library.StateMachinesWhoseNodeIsExpanded = slices.Delete(_library.StateMachinesWhoseNodeIsExpanded, idx, idx+1)
						statemachineFormCallback.probe.UpdateSliceOfPointersCallback(_library, "StateMachinesWhoseNodeIsExpanded", &_library.StateMachinesWhoseNodeIsExpanded)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if statemachineFormCallback.formGroup.HasSuppressButtonBeenPressed {
		statemachine_.Unstage(statemachineFormCallback.probe.stageOfInterest)
	}

	statemachineFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.StateMachine](
		statemachineFormCallback.probe,
	)

	// display a new form by reset the form stage
	if statemachineFormCallback.CreationMode || statemachineFormCallback.formGroup.HasSuppressButtonBeenPressed {
		statemachineFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(statemachineFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__StateMachineFormCallback(
			nil,
			statemachineFormCallback.probe,
			newFormGroup,
		)
		statemachine := new(models.StateMachine)
		FillUpForm(statemachine, newFormGroup, statemachineFormCallback.probe)
		statemachineFormCallback.probe.formStage.Commit()
	}

	statemachineFormCallback.probe.ux_tree()
}
func __gong__New__StateShapeFormCallback(
	stateshape *models.StateShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (stateshapeFormCallback *StateShapeFormCallback) {
	stateshapeFormCallback = new(StateShapeFormCallback)
	stateshapeFormCallback.probe = probe
	stateshapeFormCallback.stateshape = stateshape
	stateshapeFormCallback.formGroup = formGroup

	stateshapeFormCallback.CreationMode = (stateshape == nil)

	return
}

type StateShapeFormCallback struct {
	stateshape *models.StateShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (stateshapeFormCallback *StateShapeFormCallback) OnSave() {
	stateshapeFormCallback.probe.stageOfInterest.Lock()
	defer stateshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("StateShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	stateshapeFormCallback.probe.formStage.Checkout()

	if stateshapeFormCallback.stateshape == nil {
		stateshapeFormCallback.stateshape = new(models.StateShape).Stage(stateshapeFormCallback.probe.stageOfInterest)
	}
	stateshape_ := stateshapeFormCallback.stateshape
	_ = stateshape_

	for _, formDiv := range stateshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(stateshape_.Name), formDiv)
		case "State":
			FormDivSelectFieldToField(&(stateshape_.State), stateshapeFormCallback.probe.stageOfInterest, formDiv)
		case "X":
			FormDivBasicFieldToField(&(stateshape_.X), formDiv)
		case "Y":
			FormDivBasicFieldToField(&(stateshape_.Y), formDiv)
		case "Width":
			FormDivBasicFieldToField(&(stateshape_.Width), formDiv)
		case "Height":
			FormDivBasicFieldToField(&(stateshape_.Height), formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&(stateshape_.IsHidden), formDiv)
		case "Diagram:State_Shapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Diagram instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Diagram instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Diagram](stateshapeFormCallback.probe.stageOfInterest)
			targetDiagramIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Diagram instances and update their State_Shapes slice
			for _diagram := range *models.GetGongstructInstancesSetFromPointerType[*models.Diagram](stateshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(stateshapeFormCallback.probe.stageOfInterest, _diagram)
				
				// if Diagram is selected
				if targetDiagramIDs[id] {
					// ensure stateshape_ is in _diagram.State_Shapes
					found := false
					for _, _b := range _diagram.State_Shapes {
						if _b == stateshape_ {
							found = true
							break
						}
					}
					if !found {
						_diagram.State_Shapes = append(_diagram.State_Shapes, stateshape_)
						stateshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagram, "State_Shapes", &_diagram.State_Shapes)
					}
				} else {
					// ensure stateshape_ is NOT in _diagram.State_Shapes
					idx := slices.Index(_diagram.State_Shapes, stateshape_)
					if idx != -1 {
						_diagram.State_Shapes = slices.Delete(_diagram.State_Shapes, idx, idx+1)
						stateshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagram, "State_Shapes", &_diagram.State_Shapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if stateshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		stateshape_.Unstage(stateshapeFormCallback.probe.stageOfInterest)
	}

	stateshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.StateShape](
		stateshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if stateshapeFormCallback.CreationMode || stateshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		stateshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(stateshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__StateShapeFormCallback(
			nil,
			stateshapeFormCallback.probe,
			newFormGroup,
		)
		stateshape := new(models.StateShape)
		FillUpForm(stateshape, newFormGroup, stateshapeFormCallback.probe)
		stateshapeFormCallback.probe.formStage.Commit()
	}

	stateshapeFormCallback.probe.ux_tree()
}
func __gong__New__TransitionFormCallback(
	transition *models.Transition,
	probe *Probe,
	formGroup *form.FormGroup,
) (transitionFormCallback *TransitionFormCallback) {
	transitionFormCallback = new(TransitionFormCallback)
	transitionFormCallback.probe = probe
	transitionFormCallback.transition = transition
	transitionFormCallback.formGroup = formGroup

	transitionFormCallback.CreationMode = (transition == nil)

	return
}

type TransitionFormCallback struct {
	transition *models.Transition

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (transitionFormCallback *TransitionFormCallback) OnSave() {
	transitionFormCallback.probe.stageOfInterest.Lock()
	defer transitionFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("TransitionFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	transitionFormCallback.probe.formStage.Checkout()

	if transitionFormCallback.transition == nil {
		transitionFormCallback.transition = new(models.Transition).Stage(transitionFormCallback.probe.stageOfInterest)
	}
	transition_ := transitionFormCallback.transition
	_ = transition_

	for _, formDiv := range transitionFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(transition_.Name), formDiv)
		case "Start":
			FormDivSelectFieldToField(&(transition_.Start), transitionFormCallback.probe.stageOfInterest, formDiv)
		case "End":
			FormDivSelectFieldToField(&(transition_.End), transitionFormCallback.probe.stageOfInterest, formDiv)
		case "RolesWithPermissions":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Role](transitionFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Role, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Role)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					transitionFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Role](transitionFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			transition_.RolesWithPermissions = instanceSlice
			transitionFormCallback.probe.UpdateSliceOfPointersCallback(transition_, "RolesWithPermissions", &transition_.RolesWithPermissions)

		case "GeneratedMessages":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.MessageType](transitionFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.MessageType, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.MessageType)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					transitionFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.MessageType](transitionFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			transition_.GeneratedMessages = instanceSlice
			transitionFormCallback.probe.UpdateSliceOfPointersCallback(transition_, "GeneratedMessages", &transition_.GeneratedMessages)

		case "Guard":
			FormDivSelectFieldToField(&(transition_.Guard), transitionFormCallback.probe.stageOfInterest, formDiv)
		case "Diagrams":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Diagram](transitionFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Diagram, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Diagram)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					transitionFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Diagram](transitionFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			transition_.Diagrams = instanceSlice
			transitionFormCallback.probe.UpdateSliceOfPointersCallback(transition_, "Diagrams", &transition_.Diagrams)

		case "Note:Transitions":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Note instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Note instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Note](transitionFormCallback.probe.stageOfInterest)
			targetNoteIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetNoteIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Note instances and update their Transitions slice
			for _note := range *models.GetGongstructInstancesSetFromPointerType[*models.Note](transitionFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(transitionFormCallback.probe.stageOfInterest, _note)
				
				// if Note is selected
				if targetNoteIDs[id] {
					// ensure transition_ is in _note.Transitions
					found := false
					for _, _b := range _note.Transitions {
						if _b == transition_ {
							found = true
							break
						}
					}
					if !found {
						_note.Transitions = append(_note.Transitions, transition_)
						transitionFormCallback.probe.UpdateSliceOfPointersCallback(_note, "Transitions", &_note.Transitions)
					}
				} else {
					// ensure transition_ is NOT in _note.Transitions
					idx := slices.Index(_note.Transitions, transition_)
					if idx != -1 {
						_note.Transitions = slices.Delete(_note.Transitions, idx, idx+1)
						transitionFormCallback.probe.UpdateSliceOfPointersCallback(_note, "Transitions", &_note.Transitions)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if transitionFormCallback.formGroup.HasSuppressButtonBeenPressed {
		transition_.Unstage(transitionFormCallback.probe.stageOfInterest)
	}

	transitionFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Transition](
		transitionFormCallback.probe,
	)

	// display a new form by reset the form stage
	if transitionFormCallback.CreationMode || transitionFormCallback.formGroup.HasSuppressButtonBeenPressed {
		transitionFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(transitionFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__TransitionFormCallback(
			nil,
			transitionFormCallback.probe,
			newFormGroup,
		)
		transition := new(models.Transition)
		FillUpForm(transition, newFormGroup, transitionFormCallback.probe)
		transitionFormCallback.probe.formStage.Commit()
	}

	transitionFormCallback.probe.ux_tree()
}
func __gong__New__Transition_ShapeFormCallback(
	transition_shape *models.Transition_Shape,
	probe *Probe,
	formGroup *form.FormGroup,
) (transition_shapeFormCallback *Transition_ShapeFormCallback) {
	transition_shapeFormCallback = new(Transition_ShapeFormCallback)
	transition_shapeFormCallback.probe = probe
	transition_shapeFormCallback.transition_shape = transition_shape
	transition_shapeFormCallback.formGroup = formGroup

	transition_shapeFormCallback.CreationMode = (transition_shape == nil)

	return
}

type Transition_ShapeFormCallback struct {
	transition_shape *models.Transition_Shape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (transition_shapeFormCallback *Transition_ShapeFormCallback) OnSave() {
	transition_shapeFormCallback.probe.stageOfInterest.Lock()
	defer transition_shapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("Transition_ShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	transition_shapeFormCallback.probe.formStage.Checkout()

	if transition_shapeFormCallback.transition_shape == nil {
		transition_shapeFormCallback.transition_shape = new(models.Transition_Shape).Stage(transition_shapeFormCallback.probe.stageOfInterest)
	}
	transition_shape_ := transition_shapeFormCallback.transition_shape
	_ = transition_shape_

	for _, formDiv := range transition_shapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(transition_shape_.Name), formDiv)
		case "Transition":
			FormDivSelectFieldToField(&(transition_shape_.Transition), transition_shapeFormCallback.probe.stageOfInterest, formDiv)
		case "StartRatio":
			FormDivBasicFieldToField(&(transition_shape_.StartRatio), formDiv)
		case "EndRatio":
			FormDivBasicFieldToField(&(transition_shape_.EndRatio), formDiv)
		case "StartOrientation":
			FormDivEnumStringFieldToField(&(transition_shape_.StartOrientation), formDiv)
		case "EndOrientation":
			FormDivEnumStringFieldToField(&(transition_shape_.EndOrientation), formDiv)
		case "CornerOffsetRatio":
			FormDivBasicFieldToField(&(transition_shape_.CornerOffsetRatio), formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&(transition_shape_.IsHidden), formDiv)
		case "Diagram:Transition_Shapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Diagram instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Diagram instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Diagram](transition_shapeFormCallback.probe.stageOfInterest)
			targetDiagramIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Diagram instances and update their Transition_Shapes slice
			for _diagram := range *models.GetGongstructInstancesSetFromPointerType[*models.Diagram](transition_shapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(transition_shapeFormCallback.probe.stageOfInterest, _diagram)
				
				// if Diagram is selected
				if targetDiagramIDs[id] {
					// ensure transition_shape_ is in _diagram.Transition_Shapes
					found := false
					for _, _b := range _diagram.Transition_Shapes {
						if _b == transition_shape_ {
							found = true
							break
						}
					}
					if !found {
						_diagram.Transition_Shapes = append(_diagram.Transition_Shapes, transition_shape_)
						transition_shapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagram, "Transition_Shapes", &_diagram.Transition_Shapes)
					}
				} else {
					// ensure transition_shape_ is NOT in _diagram.Transition_Shapes
					idx := slices.Index(_diagram.Transition_Shapes, transition_shape_)
					if idx != -1 {
						_diagram.Transition_Shapes = slices.Delete(_diagram.Transition_Shapes, idx, idx+1)
						transition_shapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagram, "Transition_Shapes", &_diagram.Transition_Shapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if transition_shapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		transition_shape_.Unstage(transition_shapeFormCallback.probe.stageOfInterest)
	}

	transition_shapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Transition_Shape](
		transition_shapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if transition_shapeFormCallback.CreationMode || transition_shapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		transition_shapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(transition_shapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Transition_ShapeFormCallback(
			nil,
			transition_shapeFormCallback.probe,
			newFormGroup,
		)
		transition_shape := new(models.Transition_Shape)
		FillUpForm(transition_shape, newFormGroup, transition_shapeFormCallback.probe)
		transition_shapeFormCallback.probe.formStage.Commit()
	}

	transition_shapeFormCallback.probe.ux_tree()
}
