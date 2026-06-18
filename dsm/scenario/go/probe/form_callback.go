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

// insertion point
func __gong__New__ActorStateFormCallback(
	actorstate *models.ActorState,
	probe *Probe,
	formGroup *form.FormGroup,
) (actorstateFormCallback *ActorStateFormCallback) {
	actorstateFormCallback = new(ActorStateFormCallback)
	actorstateFormCallback.probe = probe
	actorstateFormCallback.actorstate = actorstate
	actorstateFormCallback.formGroup = formGroup

	actorstateFormCallback.CreationMode = (actorstate == nil)

	return
}

type ActorStateFormCallback struct {
	actorstate *models.ActorState

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (actorstateFormCallback *ActorStateFormCallback) OnSave() {
	actorstateFormCallback.probe.stageOfInterest.Lock()
	defer actorstateFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("ActorStateFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	actorstateFormCallback.probe.formStage.Checkout()

	if actorstateFormCallback.actorstate == nil {
		actorstateFormCallback.actorstate = new(models.ActorState).Stage(actorstateFormCallback.probe.stageOfInterest)
	}
	actorstate_ := actorstateFormCallback.actorstate
	_ = actorstate_

	for _, formDiv := range actorstateFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(actorstate_.Name), formDiv)
		case "Description":
			FormDivBasicFieldToField(&(actorstate_.Description), formDiv)
		case "IsWithProbaility":
			FormDivBasicFieldToField(&(actorstate_.IsWithProbaility), formDiv)
		case "Probability":
			FormDivEnumStringFieldToField(&(actorstate_.Probability), formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(actorstate_.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(actorstate_.IsExpanded), formDiv)
		case "LayoutDirection":
			FormDivEnumIntFieldToField(&(actorstate_.LayoutDirection), formDiv)
		case "Diagram:ActorStatesWhoseNodeIsExpanded":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Diagram instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Diagram instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Diagram](actorstateFormCallback.probe.stageOfInterest)
			targetDiagramIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Diagram instances and update their ActorStatesWhoseNodeIsExpanded slice
			for _diagram := range *models.GetGongstructInstancesSetFromPointerType[*models.Diagram](actorstateFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(actorstateFormCallback.probe.stageOfInterest, _diagram)
				
				// if Diagram is selected
				if targetDiagramIDs[id] {
					// ensure actorstate_ is in _diagram.ActorStatesWhoseNodeIsExpanded
					found := false
					for _, _b := range _diagram.ActorStatesWhoseNodeIsExpanded {
						if _b == actorstate_ {
							found = true
							break
						}
					}
					if !found {
						_diagram.ActorStatesWhoseNodeIsExpanded = append(_diagram.ActorStatesWhoseNodeIsExpanded, actorstate_)
						actorstateFormCallback.probe.UpdateSliceOfPointersCallback(_diagram, "ActorStatesWhoseNodeIsExpanded", &_diagram.ActorStatesWhoseNodeIsExpanded)
					}
				} else {
					// ensure actorstate_ is NOT in _diagram.ActorStatesWhoseNodeIsExpanded
					idx := slices.Index(_diagram.ActorStatesWhoseNodeIsExpanded, actorstate_)
					if idx != -1 {
						_diagram.ActorStatesWhoseNodeIsExpanded = slices.Delete(_diagram.ActorStatesWhoseNodeIsExpanded, idx, idx+1)
						actorstateFormCallback.probe.UpdateSliceOfPointersCallback(_diagram, "ActorStatesWhoseNodeIsExpanded", &_diagram.ActorStatesWhoseNodeIsExpanded)
					}
				}
			}
		case "Scenario:ActorStates":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Scenario instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Scenario instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Scenario](actorstateFormCallback.probe.stageOfInterest)
			targetScenarioIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetScenarioIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Scenario instances and update their ActorStates slice
			for _scenario := range *models.GetGongstructInstancesSetFromPointerType[*models.Scenario](actorstateFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(actorstateFormCallback.probe.stageOfInterest, _scenario)
				
				// if Scenario is selected
				if targetScenarioIDs[id] {
					// ensure actorstate_ is in _scenario.ActorStates
					found := false
					for _, _b := range _scenario.ActorStates {
						if _b == actorstate_ {
							found = true
							break
						}
					}
					if !found {
						_scenario.ActorStates = append(_scenario.ActorStates, actorstate_)
						actorstateFormCallback.probe.UpdateSliceOfPointersCallback(_scenario, "ActorStates", &_scenario.ActorStates)
					}
				} else {
					// ensure actorstate_ is NOT in _scenario.ActorStates
					idx := slices.Index(_scenario.ActorStates, actorstate_)
					if idx != -1 {
						_scenario.ActorStates = slices.Delete(_scenario.ActorStates, idx, idx+1)
						actorstateFormCallback.probe.UpdateSliceOfPointersCallback(_scenario, "ActorStates", &_scenario.ActorStates)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if actorstateFormCallback.formGroup.HasSuppressButtonBeenPressed {
		actorstate_.Unstage(actorstateFormCallback.probe.stageOfInterest)
	}

	actorstateFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.ActorState](
		actorstateFormCallback.probe,
	)

	// display a new form by reset the form stage
	if actorstateFormCallback.CreationMode || actorstateFormCallback.formGroup.HasSuppressButtonBeenPressed {
		actorstateFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(actorstateFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ActorStateFormCallback(
			nil,
			actorstateFormCallback.probe,
			newFormGroup,
		)
		actorstate := new(models.ActorState)
		FillUpForm(actorstate, newFormGroup, actorstateFormCallback.probe)
		actorstateFormCallback.probe.formStage.Commit()
	}

	actorstateFormCallback.probe.ux_tree()
}
func __gong__New__ActorStateShapeFormCallback(
	actorstateshape *models.ActorStateShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (actorstateshapeFormCallback *ActorStateShapeFormCallback) {
	actorstateshapeFormCallback = new(ActorStateShapeFormCallback)
	actorstateshapeFormCallback.probe = probe
	actorstateshapeFormCallback.actorstateshape = actorstateshape
	actorstateshapeFormCallback.formGroup = formGroup

	actorstateshapeFormCallback.CreationMode = (actorstateshape == nil)

	return
}

type ActorStateShapeFormCallback struct {
	actorstateshape *models.ActorStateShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (actorstateshapeFormCallback *ActorStateShapeFormCallback) OnSave() {
	actorstateshapeFormCallback.probe.stageOfInterest.Lock()
	defer actorstateshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("ActorStateShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	actorstateshapeFormCallback.probe.formStage.Checkout()

	if actorstateshapeFormCallback.actorstateshape == nil {
		actorstateshapeFormCallback.actorstateshape = new(models.ActorStateShape).Stage(actorstateshapeFormCallback.probe.stageOfInterest)
	}
	actorstateshape_ := actorstateshapeFormCallback.actorstateshape
	_ = actorstateshape_

	for _, formDiv := range actorstateshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(actorstateshape_.Name), formDiv)
		case "ActorState":
			FormDivSelectFieldToField(&(actorstateshape_.ActorState), actorstateshapeFormCallback.probe.stageOfInterest, formDiv)
		case "X":
			FormDivBasicFieldToField(&(actorstateshape_.X), formDiv)
		case "Y":
			FormDivBasicFieldToField(&(actorstateshape_.Y), formDiv)
		case "Width":
			FormDivBasicFieldToField(&(actorstateshape_.Width), formDiv)
		case "Height":
			FormDivBasicFieldToField(&(actorstateshape_.Height), formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&(actorstateshape_.IsHidden), formDiv)
		case "Diagram:ActorStateShapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Diagram instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Diagram instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Diagram](actorstateshapeFormCallback.probe.stageOfInterest)
			targetDiagramIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Diagram instances and update their ActorStateShapes slice
			for _diagram := range *models.GetGongstructInstancesSetFromPointerType[*models.Diagram](actorstateshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(actorstateshapeFormCallback.probe.stageOfInterest, _diagram)
				
				// if Diagram is selected
				if targetDiagramIDs[id] {
					// ensure actorstateshape_ is in _diagram.ActorStateShapes
					found := false
					for _, _b := range _diagram.ActorStateShapes {
						if _b == actorstateshape_ {
							found = true
							break
						}
					}
					if !found {
						_diagram.ActorStateShapes = append(_diagram.ActorStateShapes, actorstateshape_)
						actorstateshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagram, "ActorStateShapes", &_diagram.ActorStateShapes)
					}
				} else {
					// ensure actorstateshape_ is NOT in _diagram.ActorStateShapes
					idx := slices.Index(_diagram.ActorStateShapes, actorstateshape_)
					if idx != -1 {
						_diagram.ActorStateShapes = slices.Delete(_diagram.ActorStateShapes, idx, idx+1)
						actorstateshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagram, "ActorStateShapes", &_diagram.ActorStateShapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if actorstateshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		actorstateshape_.Unstage(actorstateshapeFormCallback.probe.stageOfInterest)
	}

	actorstateshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.ActorStateShape](
		actorstateshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if actorstateshapeFormCallback.CreationMode || actorstateshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		actorstateshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(actorstateshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ActorStateShapeFormCallback(
			nil,
			actorstateshapeFormCallback.probe,
			newFormGroup,
		)
		actorstateshape := new(models.ActorStateShape)
		FillUpForm(actorstateshape, newFormGroup, actorstateshapeFormCallback.probe)
		actorstateshapeFormCallback.probe.formStage.Commit()
	}

	actorstateshapeFormCallback.probe.ux_tree()
}
func __gong__New__ActorStateTransitionFormCallback(
	actorstatetransition *models.ActorStateTransition,
	probe *Probe,
	formGroup *form.FormGroup,
) (actorstatetransitionFormCallback *ActorStateTransitionFormCallback) {
	actorstatetransitionFormCallback = new(ActorStateTransitionFormCallback)
	actorstatetransitionFormCallback.probe = probe
	actorstatetransitionFormCallback.actorstatetransition = actorstatetransition
	actorstatetransitionFormCallback.formGroup = formGroup

	actorstatetransitionFormCallback.CreationMode = (actorstatetransition == nil)

	return
}

type ActorStateTransitionFormCallback struct {
	actorstatetransition *models.ActorStateTransition

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (actorstatetransitionFormCallback *ActorStateTransitionFormCallback) OnSave() {
	actorstatetransitionFormCallback.probe.stageOfInterest.Lock()
	defer actorstatetransitionFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("ActorStateTransitionFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	actorstatetransitionFormCallback.probe.formStage.Checkout()

	if actorstatetransitionFormCallback.actorstatetransition == nil {
		actorstatetransitionFormCallback.actorstatetransition = new(models.ActorStateTransition).Stage(actorstatetransitionFormCallback.probe.stageOfInterest)
	}
	actorstatetransition_ := actorstatetransitionFormCallback.actorstatetransition
	_ = actorstatetransition_

	for _, formDiv := range actorstatetransitionFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(actorstatetransition_.Name), formDiv)
		case "StartState":
			FormDivSelectFieldToField(&(actorstatetransition_.StartState), actorstatetransitionFormCallback.probe.stageOfInterest, formDiv)
		case "EndState":
			FormDivSelectFieldToField(&(actorstatetransition_.EndState), actorstatetransitionFormCallback.probe.stageOfInterest, formDiv)
		case "Justifications":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Parameter](actorstatetransitionFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Parameter, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Parameter)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					actorstatetransitionFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Parameter](actorstatetransitionFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			actorstatetransition_.Justifications = instanceSlice
			actorstatetransitionFormCallback.probe.UpdateSliceOfPointersCallback(actorstatetransition_, "Justifications", &actorstatetransition_.Justifications)

		case "ComputedPrefix":
			FormDivBasicFieldToField(&(actorstatetransition_.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(actorstatetransition_.IsExpanded), formDiv)
		case "LayoutDirection":
			FormDivEnumIntFieldToField(&(actorstatetransition_.LayoutDirection), formDiv)
		case "Diagram:ActorStateTransitionsWhoseNodeIsExpanded":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Diagram instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Diagram instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Diagram](actorstatetransitionFormCallback.probe.stageOfInterest)
			targetDiagramIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Diagram instances and update their ActorStateTransitionsWhoseNodeIsExpanded slice
			for _diagram := range *models.GetGongstructInstancesSetFromPointerType[*models.Diagram](actorstatetransitionFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(actorstatetransitionFormCallback.probe.stageOfInterest, _diagram)
				
				// if Diagram is selected
				if targetDiagramIDs[id] {
					// ensure actorstatetransition_ is in _diagram.ActorStateTransitionsWhoseNodeIsExpanded
					found := false
					for _, _b := range _diagram.ActorStateTransitionsWhoseNodeIsExpanded {
						if _b == actorstatetransition_ {
							found = true
							break
						}
					}
					if !found {
						_diagram.ActorStateTransitionsWhoseNodeIsExpanded = append(_diagram.ActorStateTransitionsWhoseNodeIsExpanded, actorstatetransition_)
						actorstatetransitionFormCallback.probe.UpdateSliceOfPointersCallback(_diagram, "ActorStateTransitionsWhoseNodeIsExpanded", &_diagram.ActorStateTransitionsWhoseNodeIsExpanded)
					}
				} else {
					// ensure actorstatetransition_ is NOT in _diagram.ActorStateTransitionsWhoseNodeIsExpanded
					idx := slices.Index(_diagram.ActorStateTransitionsWhoseNodeIsExpanded, actorstatetransition_)
					if idx != -1 {
						_diagram.ActorStateTransitionsWhoseNodeIsExpanded = slices.Delete(_diagram.ActorStateTransitionsWhoseNodeIsExpanded, idx, idx+1)
						actorstatetransitionFormCallback.probe.UpdateSliceOfPointersCallback(_diagram, "ActorStateTransitionsWhoseNodeIsExpanded", &_diagram.ActorStateTransitionsWhoseNodeIsExpanded)
					}
				}
			}
		case "Scenario:ActorStateTransitions":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Scenario instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Scenario instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Scenario](actorstatetransitionFormCallback.probe.stageOfInterest)
			targetScenarioIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetScenarioIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Scenario instances and update their ActorStateTransitions slice
			for _scenario := range *models.GetGongstructInstancesSetFromPointerType[*models.Scenario](actorstatetransitionFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(actorstatetransitionFormCallback.probe.stageOfInterest, _scenario)
				
				// if Scenario is selected
				if targetScenarioIDs[id] {
					// ensure actorstatetransition_ is in _scenario.ActorStateTransitions
					found := false
					for _, _b := range _scenario.ActorStateTransitions {
						if _b == actorstatetransition_ {
							found = true
							break
						}
					}
					if !found {
						_scenario.ActorStateTransitions = append(_scenario.ActorStateTransitions, actorstatetransition_)
						actorstatetransitionFormCallback.probe.UpdateSliceOfPointersCallback(_scenario, "ActorStateTransitions", &_scenario.ActorStateTransitions)
					}
				} else {
					// ensure actorstatetransition_ is NOT in _scenario.ActorStateTransitions
					idx := slices.Index(_scenario.ActorStateTransitions, actorstatetransition_)
					if idx != -1 {
						_scenario.ActorStateTransitions = slices.Delete(_scenario.ActorStateTransitions, idx, idx+1)
						actorstatetransitionFormCallback.probe.UpdateSliceOfPointersCallback(_scenario, "ActorStateTransitions", &_scenario.ActorStateTransitions)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if actorstatetransitionFormCallback.formGroup.HasSuppressButtonBeenPressed {
		actorstatetransition_.Unstage(actorstatetransitionFormCallback.probe.stageOfInterest)
	}

	actorstatetransitionFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.ActorStateTransition](
		actorstatetransitionFormCallback.probe,
	)

	// display a new form by reset the form stage
	if actorstatetransitionFormCallback.CreationMode || actorstatetransitionFormCallback.formGroup.HasSuppressButtonBeenPressed {
		actorstatetransitionFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(actorstatetransitionFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ActorStateTransitionFormCallback(
			nil,
			actorstatetransitionFormCallback.probe,
			newFormGroup,
		)
		actorstatetransition := new(models.ActorStateTransition)
		FillUpForm(actorstatetransition, newFormGroup, actorstatetransitionFormCallback.probe)
		actorstatetransitionFormCallback.probe.formStage.Commit()
	}

	actorstatetransitionFormCallback.probe.ux_tree()
}
func __gong__New__ActorStateTransitionShapeFormCallback(
	actorstatetransitionshape *models.ActorStateTransitionShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (actorstatetransitionshapeFormCallback *ActorStateTransitionShapeFormCallback) {
	actorstatetransitionshapeFormCallback = new(ActorStateTransitionShapeFormCallback)
	actorstatetransitionshapeFormCallback.probe = probe
	actorstatetransitionshapeFormCallback.actorstatetransitionshape = actorstatetransitionshape
	actorstatetransitionshapeFormCallback.formGroup = formGroup

	actorstatetransitionshapeFormCallback.CreationMode = (actorstatetransitionshape == nil)

	return
}

type ActorStateTransitionShapeFormCallback struct {
	actorstatetransitionshape *models.ActorStateTransitionShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (actorstatetransitionshapeFormCallback *ActorStateTransitionShapeFormCallback) OnSave() {
	actorstatetransitionshapeFormCallback.probe.stageOfInterest.Lock()
	defer actorstatetransitionshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("ActorStateTransitionShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	actorstatetransitionshapeFormCallback.probe.formStage.Checkout()

	if actorstatetransitionshapeFormCallback.actorstatetransitionshape == nil {
		actorstatetransitionshapeFormCallback.actorstatetransitionshape = new(models.ActorStateTransitionShape).Stage(actorstatetransitionshapeFormCallback.probe.stageOfInterest)
	}
	actorstatetransitionshape_ := actorstatetransitionshapeFormCallback.actorstatetransitionshape
	_ = actorstatetransitionshape_

	for _, formDiv := range actorstatetransitionshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(actorstatetransitionshape_.Name), formDiv)
		case "ActorStateTransition":
			FormDivSelectFieldToField(&(actorstatetransitionshape_.ActorStateTransition), actorstatetransitionshapeFormCallback.probe.stageOfInterest, formDiv)
		case "Start":
			FormDivSelectFieldToField(&(actorstatetransitionshape_.Start), actorstatetransitionshapeFormCallback.probe.stageOfInterest, formDiv)
		case "End":
			FormDivSelectFieldToField(&(actorstatetransitionshape_.End), actorstatetransitionshapeFormCallback.probe.stageOfInterest, formDiv)
		case "X":
			FormDivBasicFieldToField(&(actorstatetransitionshape_.X), formDiv)
		case "Y":
			FormDivBasicFieldToField(&(actorstatetransitionshape_.Y), formDiv)
		case "Width":
			FormDivBasicFieldToField(&(actorstatetransitionshape_.Width), formDiv)
		case "Height":
			FormDivBasicFieldToField(&(actorstatetransitionshape_.Height), formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&(actorstatetransitionshape_.IsHidden), formDiv)
		case "ControlPointShapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.ControlPointShape](actorstatetransitionshapeFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.ControlPointShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.ControlPointShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					actorstatetransitionshapeFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.ControlPointShape](actorstatetransitionshapeFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			actorstatetransitionshape_.ControlPointShapes = instanceSlice
			actorstatetransitionshapeFormCallback.probe.UpdateSliceOfPointersCallback(actorstatetransitionshape_, "ControlPointShapes", &actorstatetransitionshape_.ControlPointShapes)

		case "Diagram:ActorStateTransitionShapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Diagram instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Diagram instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Diagram](actorstatetransitionshapeFormCallback.probe.stageOfInterest)
			targetDiagramIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Diagram instances and update their ActorStateTransitionShapes slice
			for _diagram := range *models.GetGongstructInstancesSetFromPointerType[*models.Diagram](actorstatetransitionshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(actorstatetransitionshapeFormCallback.probe.stageOfInterest, _diagram)
				
				// if Diagram is selected
				if targetDiagramIDs[id] {
					// ensure actorstatetransitionshape_ is in _diagram.ActorStateTransitionShapes
					found := false
					for _, _b := range _diagram.ActorStateTransitionShapes {
						if _b == actorstatetransitionshape_ {
							found = true
							break
						}
					}
					if !found {
						_diagram.ActorStateTransitionShapes = append(_diagram.ActorStateTransitionShapes, actorstatetransitionshape_)
						actorstatetransitionshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagram, "ActorStateTransitionShapes", &_diagram.ActorStateTransitionShapes)
					}
				} else {
					// ensure actorstatetransitionshape_ is NOT in _diagram.ActorStateTransitionShapes
					idx := slices.Index(_diagram.ActorStateTransitionShapes, actorstatetransitionshape_)
					if idx != -1 {
						_diagram.ActorStateTransitionShapes = slices.Delete(_diagram.ActorStateTransitionShapes, idx, idx+1)
						actorstatetransitionshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagram, "ActorStateTransitionShapes", &_diagram.ActorStateTransitionShapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if actorstatetransitionshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		actorstatetransitionshape_.Unstage(actorstatetransitionshapeFormCallback.probe.stageOfInterest)
	}

	actorstatetransitionshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.ActorStateTransitionShape](
		actorstatetransitionshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if actorstatetransitionshapeFormCallback.CreationMode || actorstatetransitionshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		actorstatetransitionshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(actorstatetransitionshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ActorStateTransitionShapeFormCallback(
			nil,
			actorstatetransitionshapeFormCallback.probe,
			newFormGroup,
		)
		actorstatetransitionshape := new(models.ActorStateTransitionShape)
		FillUpForm(actorstatetransitionshape, newFormGroup, actorstatetransitionshapeFormCallback.probe)
		actorstatetransitionshapeFormCallback.probe.formStage.Commit()
	}

	actorstatetransitionshapeFormCallback.probe.ux_tree()
}
func __gong__New__AnalysisFormCallback(
	analysis *models.Analysis,
	probe *Probe,
	formGroup *form.FormGroup,
) (analysisFormCallback *AnalysisFormCallback) {
	analysisFormCallback = new(AnalysisFormCallback)
	analysisFormCallback.probe = probe
	analysisFormCallback.analysis = analysis
	analysisFormCallback.formGroup = formGroup

	analysisFormCallback.CreationMode = (analysis == nil)

	return
}

type AnalysisFormCallback struct {
	analysis *models.Analysis

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (analysisFormCallback *AnalysisFormCallback) OnSave() {
	analysisFormCallback.probe.stageOfInterest.Lock()
	defer analysisFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("AnalysisFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	analysisFormCallback.probe.formStage.Checkout()

	if analysisFormCallback.analysis == nil {
		analysisFormCallback.analysis = new(models.Analysis).Stage(analysisFormCallback.probe.stageOfInterest)
	}
	analysis_ := analysisFormCallback.analysis
	_ = analysis_

	for _, formDiv := range analysisFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(analysis_.Name), formDiv)
		case "Description":
			FormDivBasicFieldToField(&(analysis_.Description), formDiv)
		case "Scenarios":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Scenario](analysisFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Scenario, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Scenario)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					analysisFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Scenario](analysisFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			analysis_.Scenarios = instanceSlice
			analysisFormCallback.probe.UpdateSliceOfPointersCallback(analysis_, "Scenarios", &analysis_.Scenarios)

		case "IsScenariosNodeExpanded":
			FormDivBasicFieldToField(&(analysis_.IsScenariosNodeExpanded), formDiv)
		case "GroupUse":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.GroupUse](analysisFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.GroupUse, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.GroupUse)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					analysisFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.GroupUse](analysisFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			analysis_.GroupUse = instanceSlice
			analysisFormCallback.probe.UpdateSliceOfPointersCallback(analysis_, "GroupUse", &analysis_.GroupUse)

		case "IsGroupUseNodeExpanded":
			FormDivBasicFieldToField(&(analysis_.IsGroupUseNodeExpanded), formDiv)
		case "GeoObjectUse":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.GeoObjectUse](analysisFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.GeoObjectUse, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.GeoObjectUse)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					analysisFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.GeoObjectUse](analysisFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			analysis_.GeoObjectUse = instanceSlice
			analysisFormCallback.probe.UpdateSliceOfPointersCallback(analysis_, "GeoObjectUse", &analysis_.GeoObjectUse)

		case "IsGeoObjectUseNodeExpanded":
			FormDivBasicFieldToField(&(analysis_.IsGeoObjectUseNodeExpanded), formDiv)
		case "MapUse":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.MapObjectUse](analysisFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.MapObjectUse, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.MapObjectUse)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					analysisFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.MapObjectUse](analysisFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			analysis_.MapUse = instanceSlice
			analysisFormCallback.probe.UpdateSliceOfPointersCallback(analysis_, "MapUse", &analysis_.MapUse)

		case "IsMapUseNodeExpanded":
			FormDivBasicFieldToField(&(analysis_.IsMapUseNodeExpanded), formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(analysis_.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(analysis_.IsExpanded), formDiv)
		case "LayoutDirection":
			FormDivEnumIntFieldToField(&(analysis_.LayoutDirection), formDiv)
		case "Library:Analyses":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Library instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Library instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Library](analysisFormCallback.probe.stageOfInterest)
			targetLibraryIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetLibraryIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Library instances and update their Analyses slice
			for _library := range *models.GetGongstructInstancesSetFromPointerType[*models.Library](analysisFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(analysisFormCallback.probe.stageOfInterest, _library)
				
				// if Library is selected
				if targetLibraryIDs[id] {
					// ensure analysis_ is in _library.Analyses
					found := false
					for _, _b := range _library.Analyses {
						if _b == analysis_ {
							found = true
							break
						}
					}
					if !found {
						_library.Analyses = append(_library.Analyses, analysis_)
						analysisFormCallback.probe.UpdateSliceOfPointersCallback(_library, "Analyses", &_library.Analyses)
					}
				} else {
					// ensure analysis_ is NOT in _library.Analyses
					idx := slices.Index(_library.Analyses, analysis_)
					if idx != -1 {
						_library.Analyses = slices.Delete(_library.Analyses, idx, idx+1)
						analysisFormCallback.probe.UpdateSliceOfPointersCallback(_library, "Analyses", &_library.Analyses)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if analysisFormCallback.formGroup.HasSuppressButtonBeenPressed {
		analysis_.Unstage(analysisFormCallback.probe.stageOfInterest)
	}

	analysisFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Analysis](
		analysisFormCallback.probe,
	)

	// display a new form by reset the form stage
	if analysisFormCallback.CreationMode || analysisFormCallback.formGroup.HasSuppressButtonBeenPressed {
		analysisFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(analysisFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__AnalysisFormCallback(
			nil,
			analysisFormCallback.probe,
			newFormGroup,
		)
		analysis := new(models.Analysis)
		FillUpForm(analysis, newFormGroup, analysisFormCallback.probe)
		analysisFormCallback.probe.formStage.Commit()
	}

	analysisFormCallback.probe.ux_tree()
}
func __gong__New__ControlPointShapeFormCallback(
	controlpointshape *models.ControlPointShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (controlpointshapeFormCallback *ControlPointShapeFormCallback) {
	controlpointshapeFormCallback = new(ControlPointShapeFormCallback)
	controlpointshapeFormCallback.probe = probe
	controlpointshapeFormCallback.controlpointshape = controlpointshape
	controlpointshapeFormCallback.formGroup = formGroup

	controlpointshapeFormCallback.CreationMode = (controlpointshape == nil)

	return
}

type ControlPointShapeFormCallback struct {
	controlpointshape *models.ControlPointShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (controlpointshapeFormCallback *ControlPointShapeFormCallback) OnSave() {
	controlpointshapeFormCallback.probe.stageOfInterest.Lock()
	defer controlpointshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("ControlPointShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	controlpointshapeFormCallback.probe.formStage.Checkout()

	if controlpointshapeFormCallback.controlpointshape == nil {
		controlpointshapeFormCallback.controlpointshape = new(models.ControlPointShape).Stage(controlpointshapeFormCallback.probe.stageOfInterest)
	}
	controlpointshape_ := controlpointshapeFormCallback.controlpointshape
	_ = controlpointshape_

	for _, formDiv := range controlpointshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(controlpointshape_.Name), formDiv)
		case "X_Relative":
			FormDivBasicFieldToField(&(controlpointshape_.X_Relative), formDiv)
		case "Y_Relative":
			FormDivBasicFieldToField(&(controlpointshape_.Y_Relative), formDiv)
		case "IsStartShapeTheClosestShape":
			FormDivBasicFieldToField(&(controlpointshape_.IsStartShapeTheClosestShape), formDiv)
		case "ActorStateTransitionShape:ControlPointShapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the ActorStateTransitionShape instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target ActorStateTransitionShape instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.ActorStateTransitionShape](controlpointshapeFormCallback.probe.stageOfInterest)
			targetActorStateTransitionShapeIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetActorStateTransitionShapeIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all ActorStateTransitionShape instances and update their ControlPointShapes slice
			for _actorstatetransitionshape := range *models.GetGongstructInstancesSetFromPointerType[*models.ActorStateTransitionShape](controlpointshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(controlpointshapeFormCallback.probe.stageOfInterest, _actorstatetransitionshape)
				
				// if ActorStateTransitionShape is selected
				if targetActorStateTransitionShapeIDs[id] {
					// ensure controlpointshape_ is in _actorstatetransitionshape.ControlPointShapes
					found := false
					for _, _b := range _actorstatetransitionshape.ControlPointShapes {
						if _b == controlpointshape_ {
							found = true
							break
						}
					}
					if !found {
						_actorstatetransitionshape.ControlPointShapes = append(_actorstatetransitionshape.ControlPointShapes, controlpointshape_)
						controlpointshapeFormCallback.probe.UpdateSliceOfPointersCallback(_actorstatetransitionshape, "ControlPointShapes", &_actorstatetransitionshape.ControlPointShapes)
					}
				} else {
					// ensure controlpointshape_ is NOT in _actorstatetransitionshape.ControlPointShapes
					idx := slices.Index(_actorstatetransitionshape.ControlPointShapes, controlpointshape_)
					if idx != -1 {
						_actorstatetransitionshape.ControlPointShapes = slices.Delete(_actorstatetransitionshape.ControlPointShapes, idx, idx+1)
						controlpointshapeFormCallback.probe.UpdateSliceOfPointersCallback(_actorstatetransitionshape, "ControlPointShapes", &_actorstatetransitionshape.ControlPointShapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if controlpointshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		controlpointshape_.Unstage(controlpointshapeFormCallback.probe.stageOfInterest)
	}

	controlpointshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.ControlPointShape](
		controlpointshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if controlpointshapeFormCallback.CreationMode || controlpointshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		controlpointshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(controlpointshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ControlPointShapeFormCallback(
			nil,
			controlpointshapeFormCallback.probe,
			newFormGroup,
		)
		controlpointshape := new(models.ControlPointShape)
		FillUpForm(controlpointshape, newFormGroup, controlpointshapeFormCallback.probe)
		controlpointshapeFormCallback.probe.formStage.Commit()
	}

	controlpointshapeFormCallback.probe.ux_tree()
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
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(diagram_.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(diagram_.IsExpanded), formDiv)
		case "LayoutDirection":
			FormDivEnumIntFieldToField(&(diagram_.LayoutDirection), formDiv)
		case "IsChecked":
			FormDivBasicFieldToField(&(diagram_.IsChecked), formDiv)
		case "IsShowPrefix":
			FormDivBasicFieldToField(&(diagram_.IsShowPrefix), formDiv)
		case "Description":
			FormDivBasicFieldToField(&(diagram_.Description), formDiv)
		case "EvolutionDirectionShapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.EvolutionDirectionShape](diagramFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.EvolutionDirectionShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.EvolutionDirectionShape)

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
			map_RowID_ID := GetMap_RowID_ID[*models.EvolutionDirectionShape](diagramFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagram_.EvolutionDirectionShapes = instanceSlice
			diagramFormCallback.probe.UpdateSliceOfPointersCallback(diagram_, "EvolutionDirectionShapes", &diagram_.EvolutionDirectionShapes)

		case "EvolutionDirectionsWhoseNodeIsExpanded":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.EvolutionDirection](diagramFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.EvolutionDirection, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.EvolutionDirection)

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
			map_RowID_ID := GetMap_RowID_ID[*models.EvolutionDirection](diagramFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagram_.EvolutionDirectionsWhoseNodeIsExpanded = instanceSlice
			diagramFormCallback.probe.UpdateSliceOfPointersCallback(diagram_, "EvolutionDirectionsWhoseNodeIsExpanded", &diagram_.EvolutionDirectionsWhoseNodeIsExpanded)

		case "IsEvolutionDirectionsNodeExpanded":
			FormDivBasicFieldToField(&(diagram_.IsEvolutionDirectionsNodeExpanded), formDiv)
		case "ActorStateShapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.ActorStateShape](diagramFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.ActorStateShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.ActorStateShape)

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
			map_RowID_ID := GetMap_RowID_ID[*models.ActorStateShape](diagramFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagram_.ActorStateShapes = instanceSlice
			diagramFormCallback.probe.UpdateSliceOfPointersCallback(diagram_, "ActorStateShapes", &diagram_.ActorStateShapes)

		case "ActorStatesWhoseNodeIsExpanded":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.ActorState](diagramFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.ActorState, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.ActorState)

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
			map_RowID_ID := GetMap_RowID_ID[*models.ActorState](diagramFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagram_.ActorStatesWhoseNodeIsExpanded = instanceSlice
			diagramFormCallback.probe.UpdateSliceOfPointersCallback(diagram_, "ActorStatesWhoseNodeIsExpanded", &diagram_.ActorStatesWhoseNodeIsExpanded)

		case "IsActorStatesNodeExpanded":
			FormDivBasicFieldToField(&(diagram_.IsActorStatesNodeExpanded), formDiv)
		case "ParameterShapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.ParameterShape](diagramFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.ParameterShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.ParameterShape)

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
			map_RowID_ID := GetMap_RowID_ID[*models.ParameterShape](diagramFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagram_.ParameterShapes = instanceSlice
			diagramFormCallback.probe.UpdateSliceOfPointersCallback(diagram_, "ParameterShapes", &diagram_.ParameterShapes)

		case "ParametersWhoseNodeIsExpanded":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Parameter](diagramFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Parameter, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Parameter)

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
			map_RowID_ID := GetMap_RowID_ID[*models.Parameter](diagramFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagram_.ParametersWhoseNodeIsExpanded = instanceSlice
			diagramFormCallback.probe.UpdateSliceOfPointersCallback(diagram_, "ParametersWhoseNodeIsExpanded", &diagram_.ParametersWhoseNodeIsExpanded)

		case "IsParametersNodeExpanded":
			FormDivBasicFieldToField(&(diagram_.IsParametersNodeExpanded), formDiv)
		case "ScenarioParameterShapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.ParametersAggregateShape](diagramFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.ParametersAggregateShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.ParametersAggregateShape)

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
			map_RowID_ID := GetMap_RowID_ID[*models.ParametersAggregateShape](diagramFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagram_.ScenarioParameterShapes = instanceSlice
			diagramFormCallback.probe.UpdateSliceOfPointersCallback(diagram_, "ScenarioParameterShapes", &diagram_.ScenarioParameterShapes)

		case "ParametersAggregatesWhoseNodeIsExpanded":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.ParametersAggregate](diagramFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.ParametersAggregate, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.ParametersAggregate)

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
			map_RowID_ID := GetMap_RowID_ID[*models.ParametersAggregate](diagramFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagram_.ParametersAggregatesWhoseNodeIsExpanded = instanceSlice
			diagramFormCallback.probe.UpdateSliceOfPointersCallback(diagram_, "ParametersAggregatesWhoseNodeIsExpanded", &diagram_.ParametersAggregatesWhoseNodeIsExpanded)

		case "IsParametersAggregatesNodeExpanded":
			FormDivBasicFieldToField(&(diagram_.IsParametersAggregatesNodeExpanded), formDiv)
		case "ActorStateTransitionShapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.ActorStateTransitionShape](diagramFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.ActorStateTransitionShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.ActorStateTransitionShape)

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
			map_RowID_ID := GetMap_RowID_ID[*models.ActorStateTransitionShape](diagramFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagram_.ActorStateTransitionShapes = instanceSlice
			diagramFormCallback.probe.UpdateSliceOfPointersCallback(diagram_, "ActorStateTransitionShapes", &diagram_.ActorStateTransitionShapes)

		case "ActorStateTransitionsWhoseNodeIsExpanded":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.ActorStateTransition](diagramFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.ActorStateTransition, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.ActorStateTransition)

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
			map_RowID_ID := GetMap_RowID_ID[*models.ActorStateTransition](diagramFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagram_.ActorStateTransitionsWhoseNodeIsExpanded = instanceSlice
			diagramFormCallback.probe.UpdateSliceOfPointersCallback(diagram_, "ActorStateTransitionsWhoseNodeIsExpanded", &diagram_.ActorStateTransitionsWhoseNodeIsExpanded)

		case "IsActorStateTransitionsNodeExpanded":
			FormDivBasicFieldToField(&(diagram_.IsActorStateTransitionsNodeExpanded), formDiv)
		case "AxisOrign_X":
			FormDivBasicFieldToField(&(diagram_.AxisOrign_X), formDiv)
		case "AxisOrign_Y":
			FormDivBasicFieldToField(&(diagram_.AxisOrign_Y), formDiv)
		case "VerticalAxis_Top_Y":
			FormDivBasicFieldToField(&(diagram_.VerticalAxis_Top_Y), formDiv)
		case "VerticalAxis_Bottom_Y":
			FormDivBasicFieldToField(&(diagram_.VerticalAxis_Bottom_Y), formDiv)
		case "VerticalAxis_StrokeWidth":
			FormDivBasicFieldToField(&(diagram_.VerticalAxis_StrokeWidth), formDiv)
		case "HorizontalAxis_Right_X":
			FormDivBasicFieldToField(&(diagram_.HorizontalAxis_Right_X), formDiv)
		case "Start":
			FormDivBasicFieldToField(&(diagram_.Start), formDiv)
		case "End":
			FormDivBasicFieldToField(&(diagram_.End), formDiv)
		case "NumberOfYearsBetweenTicks":
			FormDivBasicFieldToField(&(diagram_.NumberOfYearsBetweenTicks), formDiv)
		case "Scenario:Diagrams":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Scenario instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Scenario instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Scenario](diagramFormCallback.probe.stageOfInterest)
			targetScenarioIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetScenarioIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Scenario instances and update their Diagrams slice
			for _scenario := range *models.GetGongstructInstancesSetFromPointerType[*models.Scenario](diagramFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(diagramFormCallback.probe.stageOfInterest, _scenario)
				
				// if Scenario is selected
				if targetScenarioIDs[id] {
					// ensure diagram_ is in _scenario.Diagrams
					found := false
					for _, _b := range _scenario.Diagrams {
						if _b == diagram_ {
							found = true
							break
						}
					}
					if !found {
						_scenario.Diagrams = append(_scenario.Diagrams, diagram_)
						diagramFormCallback.probe.UpdateSliceOfPointersCallback(_scenario, "Diagrams", &_scenario.Diagrams)
					}
				} else {
					// ensure diagram_ is NOT in _scenario.Diagrams
					idx := slices.Index(_scenario.Diagrams, diagram_)
					if idx != -1 {
						_scenario.Diagrams = slices.Delete(_scenario.Diagrams, idx, idx+1)
						diagramFormCallback.probe.UpdateSliceOfPointersCallback(_scenario, "Diagrams", &_scenario.Diagrams)
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
func __gong__New__DocumentFormCallback(
	document *models.Document,
	probe *Probe,
	formGroup *form.FormGroup,
) (documentFormCallback *DocumentFormCallback) {
	documentFormCallback = new(DocumentFormCallback)
	documentFormCallback.probe = probe
	documentFormCallback.document = document
	documentFormCallback.formGroup = formGroup

	documentFormCallback.CreationMode = (document == nil)

	return
}

type DocumentFormCallback struct {
	document *models.Document

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (documentFormCallback *DocumentFormCallback) OnSave() {
	documentFormCallback.probe.stageOfInterest.Lock()
	defer documentFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("DocumentFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	documentFormCallback.probe.formStage.Checkout()

	if documentFormCallback.document == nil {
		documentFormCallback.document = new(models.Document).Stage(documentFormCallback.probe.stageOfInterest)
	}
	document_ := documentFormCallback.document
	_ = document_

	for _, formDiv := range documentFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(document_.Name), formDiv)
		case "GeoObjectUse":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.GeoObjectUse](documentFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.GeoObjectUse, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.GeoObjectUse)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					documentFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.GeoObjectUse](documentFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			document_.GeoObjectUse = instanceSlice
			documentFormCallback.probe.UpdateSliceOfPointersCallback(document_, "GeoObjectUse", &document_.GeoObjectUse)

		case "ComputedPrefix":
			FormDivBasicFieldToField(&(document_.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(document_.IsExpanded), formDiv)
		case "LayoutDirection":
			FormDivEnumIntFieldToField(&(document_.LayoutDirection), formDiv)
		}
	}

	// manage the suppress operation
	if documentFormCallback.formGroup.HasSuppressButtonBeenPressed {
		document_.Unstage(documentFormCallback.probe.stageOfInterest)
	}

	documentFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Document](
		documentFormCallback.probe,
	)

	// display a new form by reset the form stage
	if documentFormCallback.CreationMode || documentFormCallback.formGroup.HasSuppressButtonBeenPressed {
		documentFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(documentFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__DocumentFormCallback(
			nil,
			documentFormCallback.probe,
			newFormGroup,
		)
		document := new(models.Document)
		FillUpForm(document, newFormGroup, documentFormCallback.probe)
		documentFormCallback.probe.formStage.Commit()
	}

	documentFormCallback.probe.ux_tree()
}
func __gong__New__DocumentUseFormCallback(
	documentuse *models.DocumentUse,
	probe *Probe,
	formGroup *form.FormGroup,
) (documentuseFormCallback *DocumentUseFormCallback) {
	documentuseFormCallback = new(DocumentUseFormCallback)
	documentuseFormCallback.probe = probe
	documentuseFormCallback.documentuse = documentuse
	documentuseFormCallback.formGroup = formGroup

	documentuseFormCallback.CreationMode = (documentuse == nil)

	return
}

type DocumentUseFormCallback struct {
	documentuse *models.DocumentUse

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (documentuseFormCallback *DocumentUseFormCallback) OnSave() {
	documentuseFormCallback.probe.stageOfInterest.Lock()
	defer documentuseFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("DocumentUseFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	documentuseFormCallback.probe.formStage.Checkout()

	if documentuseFormCallback.documentuse == nil {
		documentuseFormCallback.documentuse = new(models.DocumentUse).Stage(documentuseFormCallback.probe.stageOfInterest)
	}
	documentuse_ := documentuseFormCallback.documentuse
	_ = documentuse_

	for _, formDiv := range documentuseFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(documentuse_.Name), formDiv)
		case "Document":
			FormDivSelectFieldToField(&(documentuse_.Document), documentuseFormCallback.probe.stageOfInterest, formDiv)
		case "Parameter:DocumentUse":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Parameter instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Parameter instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Parameter](documentuseFormCallback.probe.stageOfInterest)
			targetParameterIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetParameterIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Parameter instances and update their DocumentUse slice
			for _parameter := range *models.GetGongstructInstancesSetFromPointerType[*models.Parameter](documentuseFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(documentuseFormCallback.probe.stageOfInterest, _parameter)
				
				// if Parameter is selected
				if targetParameterIDs[id] {
					// ensure documentuse_ is in _parameter.DocumentUse
					found := false
					for _, _b := range _parameter.DocumentUse {
						if _b == documentuse_ {
							found = true
							break
						}
					}
					if !found {
						_parameter.DocumentUse = append(_parameter.DocumentUse, documentuse_)
						documentuseFormCallback.probe.UpdateSliceOfPointersCallback(_parameter, "DocumentUse", &_parameter.DocumentUse)
					}
				} else {
					// ensure documentuse_ is NOT in _parameter.DocumentUse
					idx := slices.Index(_parameter.DocumentUse, documentuse_)
					if idx != -1 {
						_parameter.DocumentUse = slices.Delete(_parameter.DocumentUse, idx, idx+1)
						documentuseFormCallback.probe.UpdateSliceOfPointersCallback(_parameter, "DocumentUse", &_parameter.DocumentUse)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if documentuseFormCallback.formGroup.HasSuppressButtonBeenPressed {
		documentuse_.Unstage(documentuseFormCallback.probe.stageOfInterest)
	}

	documentuseFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.DocumentUse](
		documentuseFormCallback.probe,
	)

	// display a new form by reset the form stage
	if documentuseFormCallback.CreationMode || documentuseFormCallback.formGroup.HasSuppressButtonBeenPressed {
		documentuseFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(documentuseFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__DocumentUseFormCallback(
			nil,
			documentuseFormCallback.probe,
			newFormGroup,
		)
		documentuse := new(models.DocumentUse)
		FillUpForm(documentuse, newFormGroup, documentuseFormCallback.probe)
		documentuseFormCallback.probe.formStage.Commit()
	}

	documentuseFormCallback.probe.ux_tree()
}
func __gong__New__EvolutionDirectionFormCallback(
	evolutiondirection *models.EvolutionDirection,
	probe *Probe,
	formGroup *form.FormGroup,
) (evolutiondirectionFormCallback *EvolutionDirectionFormCallback) {
	evolutiondirectionFormCallback = new(EvolutionDirectionFormCallback)
	evolutiondirectionFormCallback.probe = probe
	evolutiondirectionFormCallback.evolutiondirection = evolutiondirection
	evolutiondirectionFormCallback.formGroup = formGroup

	evolutiondirectionFormCallback.CreationMode = (evolutiondirection == nil)

	return
}

type EvolutionDirectionFormCallback struct {
	evolutiondirection *models.EvolutionDirection

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (evolutiondirectionFormCallback *EvolutionDirectionFormCallback) OnSave() {
	evolutiondirectionFormCallback.probe.stageOfInterest.Lock()
	defer evolutiondirectionFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("EvolutionDirectionFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	evolutiondirectionFormCallback.probe.formStage.Checkout()

	if evolutiondirectionFormCallback.evolutiondirection == nil {
		evolutiondirectionFormCallback.evolutiondirection = new(models.EvolutionDirection).Stage(evolutiondirectionFormCallback.probe.stageOfInterest)
	}
	evolutiondirection_ := evolutiondirectionFormCallback.evolutiondirection
	_ = evolutiondirection_

	for _, formDiv := range evolutiondirectionFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(evolutiondirection_.Name), formDiv)
		case "Description":
			FormDivBasicFieldToField(&(evolutiondirection_.Description), formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(evolutiondirection_.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(evolutiondirection_.IsExpanded), formDiv)
		case "LayoutDirection":
			FormDivEnumIntFieldToField(&(evolutiondirection_.LayoutDirection), formDiv)
		case "Diagram:EvolutionDirectionsWhoseNodeIsExpanded":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Diagram instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Diagram instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Diagram](evolutiondirectionFormCallback.probe.stageOfInterest)
			targetDiagramIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Diagram instances and update their EvolutionDirectionsWhoseNodeIsExpanded slice
			for _diagram := range *models.GetGongstructInstancesSetFromPointerType[*models.Diagram](evolutiondirectionFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(evolutiondirectionFormCallback.probe.stageOfInterest, _diagram)
				
				// if Diagram is selected
				if targetDiagramIDs[id] {
					// ensure evolutiondirection_ is in _diagram.EvolutionDirectionsWhoseNodeIsExpanded
					found := false
					for _, _b := range _diagram.EvolutionDirectionsWhoseNodeIsExpanded {
						if _b == evolutiondirection_ {
							found = true
							break
						}
					}
					if !found {
						_diagram.EvolutionDirectionsWhoseNodeIsExpanded = append(_diagram.EvolutionDirectionsWhoseNodeIsExpanded, evolutiondirection_)
						evolutiondirectionFormCallback.probe.UpdateSliceOfPointersCallback(_diagram, "EvolutionDirectionsWhoseNodeIsExpanded", &_diagram.EvolutionDirectionsWhoseNodeIsExpanded)
					}
				} else {
					// ensure evolutiondirection_ is NOT in _diagram.EvolutionDirectionsWhoseNodeIsExpanded
					idx := slices.Index(_diagram.EvolutionDirectionsWhoseNodeIsExpanded, evolutiondirection_)
					if idx != -1 {
						_diagram.EvolutionDirectionsWhoseNodeIsExpanded = slices.Delete(_diagram.EvolutionDirectionsWhoseNodeIsExpanded, idx, idx+1)
						evolutiondirectionFormCallback.probe.UpdateSliceOfPointersCallback(_diagram, "EvolutionDirectionsWhoseNodeIsExpanded", &_diagram.EvolutionDirectionsWhoseNodeIsExpanded)
					}
				}
			}
		case "Scenario:EvolutionDirections":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Scenario instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Scenario instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Scenario](evolutiondirectionFormCallback.probe.stageOfInterest)
			targetScenarioIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetScenarioIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Scenario instances and update their EvolutionDirections slice
			for _scenario := range *models.GetGongstructInstancesSetFromPointerType[*models.Scenario](evolutiondirectionFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(evolutiondirectionFormCallback.probe.stageOfInterest, _scenario)
				
				// if Scenario is selected
				if targetScenarioIDs[id] {
					// ensure evolutiondirection_ is in _scenario.EvolutionDirections
					found := false
					for _, _b := range _scenario.EvolutionDirections {
						if _b == evolutiondirection_ {
							found = true
							break
						}
					}
					if !found {
						_scenario.EvolutionDirections = append(_scenario.EvolutionDirections, evolutiondirection_)
						evolutiondirectionFormCallback.probe.UpdateSliceOfPointersCallback(_scenario, "EvolutionDirections", &_scenario.EvolutionDirections)
					}
				} else {
					// ensure evolutiondirection_ is NOT in _scenario.EvolutionDirections
					idx := slices.Index(_scenario.EvolutionDirections, evolutiondirection_)
					if idx != -1 {
						_scenario.EvolutionDirections = slices.Delete(_scenario.EvolutionDirections, idx, idx+1)
						evolutiondirectionFormCallback.probe.UpdateSliceOfPointersCallback(_scenario, "EvolutionDirections", &_scenario.EvolutionDirections)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if evolutiondirectionFormCallback.formGroup.HasSuppressButtonBeenPressed {
		evolutiondirection_.Unstage(evolutiondirectionFormCallback.probe.stageOfInterest)
	}

	evolutiondirectionFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.EvolutionDirection](
		evolutiondirectionFormCallback.probe,
	)

	// display a new form by reset the form stage
	if evolutiondirectionFormCallback.CreationMode || evolutiondirectionFormCallback.formGroup.HasSuppressButtonBeenPressed {
		evolutiondirectionFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(evolutiondirectionFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__EvolutionDirectionFormCallback(
			nil,
			evolutiondirectionFormCallback.probe,
			newFormGroup,
		)
		evolutiondirection := new(models.EvolutionDirection)
		FillUpForm(evolutiondirection, newFormGroup, evolutiondirectionFormCallback.probe)
		evolutiondirectionFormCallback.probe.formStage.Commit()
	}

	evolutiondirectionFormCallback.probe.ux_tree()
}
func __gong__New__EvolutionDirectionShapeFormCallback(
	evolutiondirectionshape *models.EvolutionDirectionShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (evolutiondirectionshapeFormCallback *EvolutionDirectionShapeFormCallback) {
	evolutiondirectionshapeFormCallback = new(EvolutionDirectionShapeFormCallback)
	evolutiondirectionshapeFormCallback.probe = probe
	evolutiondirectionshapeFormCallback.evolutiondirectionshape = evolutiondirectionshape
	evolutiondirectionshapeFormCallback.formGroup = formGroup

	evolutiondirectionshapeFormCallback.CreationMode = (evolutiondirectionshape == nil)

	return
}

type EvolutionDirectionShapeFormCallback struct {
	evolutiondirectionshape *models.EvolutionDirectionShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (evolutiondirectionshapeFormCallback *EvolutionDirectionShapeFormCallback) OnSave() {
	evolutiondirectionshapeFormCallback.probe.stageOfInterest.Lock()
	defer evolutiondirectionshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("EvolutionDirectionShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	evolutiondirectionshapeFormCallback.probe.formStage.Checkout()

	if evolutiondirectionshapeFormCallback.evolutiondirectionshape == nil {
		evolutiondirectionshapeFormCallback.evolutiondirectionshape = new(models.EvolutionDirectionShape).Stage(evolutiondirectionshapeFormCallback.probe.stageOfInterest)
	}
	evolutiondirectionshape_ := evolutiondirectionshapeFormCallback.evolutiondirectionshape
	_ = evolutiondirectionshape_

	for _, formDiv := range evolutiondirectionshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(evolutiondirectionshape_.Name), formDiv)
		case "EvolutionDirection":
			FormDivSelectFieldToField(&(evolutiondirectionshape_.EvolutionDirection), evolutiondirectionshapeFormCallback.probe.stageOfInterest, formDiv)
		case "X":
			FormDivBasicFieldToField(&(evolutiondirectionshape_.X), formDiv)
		case "Y":
			FormDivBasicFieldToField(&(evolutiondirectionshape_.Y), formDiv)
		case "Width":
			FormDivBasicFieldToField(&(evolutiondirectionshape_.Width), formDiv)
		case "Height":
			FormDivBasicFieldToField(&(evolutiondirectionshape_.Height), formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&(evolutiondirectionshape_.IsHidden), formDiv)
		case "Diagram:EvolutionDirectionShapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Diagram instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Diagram instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Diagram](evolutiondirectionshapeFormCallback.probe.stageOfInterest)
			targetDiagramIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Diagram instances and update their EvolutionDirectionShapes slice
			for _diagram := range *models.GetGongstructInstancesSetFromPointerType[*models.Diagram](evolutiondirectionshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(evolutiondirectionshapeFormCallback.probe.stageOfInterest, _diagram)
				
				// if Diagram is selected
				if targetDiagramIDs[id] {
					// ensure evolutiondirectionshape_ is in _diagram.EvolutionDirectionShapes
					found := false
					for _, _b := range _diagram.EvolutionDirectionShapes {
						if _b == evolutiondirectionshape_ {
							found = true
							break
						}
					}
					if !found {
						_diagram.EvolutionDirectionShapes = append(_diagram.EvolutionDirectionShapes, evolutiondirectionshape_)
						evolutiondirectionshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagram, "EvolutionDirectionShapes", &_diagram.EvolutionDirectionShapes)
					}
				} else {
					// ensure evolutiondirectionshape_ is NOT in _diagram.EvolutionDirectionShapes
					idx := slices.Index(_diagram.EvolutionDirectionShapes, evolutiondirectionshape_)
					if idx != -1 {
						_diagram.EvolutionDirectionShapes = slices.Delete(_diagram.EvolutionDirectionShapes, idx, idx+1)
						evolutiondirectionshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagram, "EvolutionDirectionShapes", &_diagram.EvolutionDirectionShapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if evolutiondirectionshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		evolutiondirectionshape_.Unstage(evolutiondirectionshapeFormCallback.probe.stageOfInterest)
	}

	evolutiondirectionshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.EvolutionDirectionShape](
		evolutiondirectionshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if evolutiondirectionshapeFormCallback.CreationMode || evolutiondirectionshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		evolutiondirectionshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(evolutiondirectionshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__EvolutionDirectionShapeFormCallback(
			nil,
			evolutiondirectionshapeFormCallback.probe,
			newFormGroup,
		)
		evolutiondirectionshape := new(models.EvolutionDirectionShape)
		FillUpForm(evolutiondirectionshape, newFormGroup, evolutiondirectionshapeFormCallback.probe)
		evolutiondirectionshapeFormCallback.probe.formStage.Commit()
	}

	evolutiondirectionshapeFormCallback.probe.ux_tree()
}
func __gong__New__FooFormCallback(
	foo *models.Foo,
	probe *Probe,
	formGroup *form.FormGroup,
) (fooFormCallback *FooFormCallback) {
	fooFormCallback = new(FooFormCallback)
	fooFormCallback.probe = probe
	fooFormCallback.foo = foo
	fooFormCallback.formGroup = formGroup

	fooFormCallback.CreationMode = (foo == nil)

	return
}

type FooFormCallback struct {
	foo *models.Foo

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (fooFormCallback *FooFormCallback) OnSave() {
	fooFormCallback.probe.stageOfInterest.Lock()
	defer fooFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("FooFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	fooFormCallback.probe.formStage.Checkout()

	if fooFormCallback.foo == nil {
		fooFormCallback.foo = new(models.Foo).Stage(fooFormCallback.probe.stageOfInterest)
	}
	foo_ := fooFormCallback.foo
	_ = foo_

	for _, formDiv := range fooFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(foo_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if fooFormCallback.formGroup.HasSuppressButtonBeenPressed {
		foo_.Unstage(fooFormCallback.probe.stageOfInterest)
	}

	fooFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Foo](
		fooFormCallback.probe,
	)

	// display a new form by reset the form stage
	if fooFormCallback.CreationMode || fooFormCallback.formGroup.HasSuppressButtonBeenPressed {
		fooFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(fooFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__FooFormCallback(
			nil,
			fooFormCallback.probe,
			newFormGroup,
		)
		foo := new(models.Foo)
		FillUpForm(foo, newFormGroup, fooFormCallback.probe)
		fooFormCallback.probe.formStage.Commit()
	}

	fooFormCallback.probe.ux_tree()
}
func __gong__New__GeoObjectFormCallback(
	geoobject *models.GeoObject,
	probe *Probe,
	formGroup *form.FormGroup,
) (geoobjectFormCallback *GeoObjectFormCallback) {
	geoobjectFormCallback = new(GeoObjectFormCallback)
	geoobjectFormCallback.probe = probe
	geoobjectFormCallback.geoobject = geoobject
	geoobjectFormCallback.formGroup = formGroup

	geoobjectFormCallback.CreationMode = (geoobject == nil)

	return
}

type GeoObjectFormCallback struct {
	geoobject *models.GeoObject

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (geoobjectFormCallback *GeoObjectFormCallback) OnSave() {
	geoobjectFormCallback.probe.stageOfInterest.Lock()
	defer geoobjectFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("GeoObjectFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	geoobjectFormCallback.probe.formStage.Checkout()

	if geoobjectFormCallback.geoobject == nil {
		geoobjectFormCallback.geoobject = new(models.GeoObject).Stage(geoobjectFormCallback.probe.stageOfInterest)
	}
	geoobject_ := geoobjectFormCallback.geoobject
	_ = geoobject_

	for _, formDiv := range geoobjectFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(geoobject_.Name), formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(geoobject_.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(geoobject_.IsExpanded), formDiv)
		case "LayoutDirection":
			FormDivEnumIntFieldToField(&(geoobject_.LayoutDirection), formDiv)
		}
	}

	// manage the suppress operation
	if geoobjectFormCallback.formGroup.HasSuppressButtonBeenPressed {
		geoobject_.Unstage(geoobjectFormCallback.probe.stageOfInterest)
	}

	geoobjectFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.GeoObject](
		geoobjectFormCallback.probe,
	)

	// display a new form by reset the form stage
	if geoobjectFormCallback.CreationMode || geoobjectFormCallback.formGroup.HasSuppressButtonBeenPressed {
		geoobjectFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(geoobjectFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__GeoObjectFormCallback(
			nil,
			geoobjectFormCallback.probe,
			newFormGroup,
		)
		geoobject := new(models.GeoObject)
		FillUpForm(geoobject, newFormGroup, geoobjectFormCallback.probe)
		geoobjectFormCallback.probe.formStage.Commit()
	}

	geoobjectFormCallback.probe.ux_tree()
}
func __gong__New__GeoObjectUseFormCallback(
	geoobjectuse *models.GeoObjectUse,
	probe *Probe,
	formGroup *form.FormGroup,
) (geoobjectuseFormCallback *GeoObjectUseFormCallback) {
	geoobjectuseFormCallback = new(GeoObjectUseFormCallback)
	geoobjectuseFormCallback.probe = probe
	geoobjectuseFormCallback.geoobjectuse = geoobjectuse
	geoobjectuseFormCallback.formGroup = formGroup

	geoobjectuseFormCallback.CreationMode = (geoobjectuse == nil)

	return
}

type GeoObjectUseFormCallback struct {
	geoobjectuse *models.GeoObjectUse

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (geoobjectuseFormCallback *GeoObjectUseFormCallback) OnSave() {
	geoobjectuseFormCallback.probe.stageOfInterest.Lock()
	defer geoobjectuseFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("GeoObjectUseFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	geoobjectuseFormCallback.probe.formStage.Checkout()

	if geoobjectuseFormCallback.geoobjectuse == nil {
		geoobjectuseFormCallback.geoobjectuse = new(models.GeoObjectUse).Stage(geoobjectuseFormCallback.probe.stageOfInterest)
	}
	geoobjectuse_ := geoobjectuseFormCallback.geoobjectuse
	_ = geoobjectuse_

	for _, formDiv := range geoobjectuseFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(geoobjectuse_.Name), formDiv)
		case "GeoObject":
			FormDivSelectFieldToField(&(geoobjectuse_.GeoObject), geoobjectuseFormCallback.probe.stageOfInterest, formDiv)
		case "Analysis:GeoObjectUse":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Analysis instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Analysis instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Analysis](geoobjectuseFormCallback.probe.stageOfInterest)
			targetAnalysisIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetAnalysisIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Analysis instances and update their GeoObjectUse slice
			for _analysis := range *models.GetGongstructInstancesSetFromPointerType[*models.Analysis](geoobjectuseFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(geoobjectuseFormCallback.probe.stageOfInterest, _analysis)
				
				// if Analysis is selected
				if targetAnalysisIDs[id] {
					// ensure geoobjectuse_ is in _analysis.GeoObjectUse
					found := false
					for _, _b := range _analysis.GeoObjectUse {
						if _b == geoobjectuse_ {
							found = true
							break
						}
					}
					if !found {
						_analysis.GeoObjectUse = append(_analysis.GeoObjectUse, geoobjectuse_)
						geoobjectuseFormCallback.probe.UpdateSliceOfPointersCallback(_analysis, "GeoObjectUse", &_analysis.GeoObjectUse)
					}
				} else {
					// ensure geoobjectuse_ is NOT in _analysis.GeoObjectUse
					idx := slices.Index(_analysis.GeoObjectUse, geoobjectuse_)
					if idx != -1 {
						_analysis.GeoObjectUse = slices.Delete(_analysis.GeoObjectUse, idx, idx+1)
						geoobjectuseFormCallback.probe.UpdateSliceOfPointersCallback(_analysis, "GeoObjectUse", &_analysis.GeoObjectUse)
					}
				}
			}
		case "Document:GeoObjectUse":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Document instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Document instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Document](geoobjectuseFormCallback.probe.stageOfInterest)
			targetDocumentIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDocumentIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Document instances and update their GeoObjectUse slice
			for _document := range *models.GetGongstructInstancesSetFromPointerType[*models.Document](geoobjectuseFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(geoobjectuseFormCallback.probe.stageOfInterest, _document)
				
				// if Document is selected
				if targetDocumentIDs[id] {
					// ensure geoobjectuse_ is in _document.GeoObjectUse
					found := false
					for _, _b := range _document.GeoObjectUse {
						if _b == geoobjectuse_ {
							found = true
							break
						}
					}
					if !found {
						_document.GeoObjectUse = append(_document.GeoObjectUse, geoobjectuse_)
						geoobjectuseFormCallback.probe.UpdateSliceOfPointersCallback(_document, "GeoObjectUse", &_document.GeoObjectUse)
					}
				} else {
					// ensure geoobjectuse_ is NOT in _document.GeoObjectUse
					idx := slices.Index(_document.GeoObjectUse, geoobjectuse_)
					if idx != -1 {
						_document.GeoObjectUse = slices.Delete(_document.GeoObjectUse, idx, idx+1)
						geoobjectuseFormCallback.probe.UpdateSliceOfPointersCallback(_document, "GeoObjectUse", &_document.GeoObjectUse)
					}
				}
			}
		case "Parameter:GeoObjectUse":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Parameter instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Parameter instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Parameter](geoobjectuseFormCallback.probe.stageOfInterest)
			targetParameterIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetParameterIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Parameter instances and update their GeoObjectUse slice
			for _parameter := range *models.GetGongstructInstancesSetFromPointerType[*models.Parameter](geoobjectuseFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(geoobjectuseFormCallback.probe.stageOfInterest, _parameter)
				
				// if Parameter is selected
				if targetParameterIDs[id] {
					// ensure geoobjectuse_ is in _parameter.GeoObjectUse
					found := false
					for _, _b := range _parameter.GeoObjectUse {
						if _b == geoobjectuse_ {
							found = true
							break
						}
					}
					if !found {
						_parameter.GeoObjectUse = append(_parameter.GeoObjectUse, geoobjectuse_)
						geoobjectuseFormCallback.probe.UpdateSliceOfPointersCallback(_parameter, "GeoObjectUse", &_parameter.GeoObjectUse)
					}
				} else {
					// ensure geoobjectuse_ is NOT in _parameter.GeoObjectUse
					idx := slices.Index(_parameter.GeoObjectUse, geoobjectuse_)
					if idx != -1 {
						_parameter.GeoObjectUse = slices.Delete(_parameter.GeoObjectUse, idx, idx+1)
						geoobjectuseFormCallback.probe.UpdateSliceOfPointersCallback(_parameter, "GeoObjectUse", &_parameter.GeoObjectUse)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if geoobjectuseFormCallback.formGroup.HasSuppressButtonBeenPressed {
		geoobjectuse_.Unstage(geoobjectuseFormCallback.probe.stageOfInterest)
	}

	geoobjectuseFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.GeoObjectUse](
		geoobjectuseFormCallback.probe,
	)

	// display a new form by reset the form stage
	if geoobjectuseFormCallback.CreationMode || geoobjectuseFormCallback.formGroup.HasSuppressButtonBeenPressed {
		geoobjectuseFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(geoobjectuseFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__GeoObjectUseFormCallback(
			nil,
			geoobjectuseFormCallback.probe,
			newFormGroup,
		)
		geoobjectuse := new(models.GeoObjectUse)
		FillUpForm(geoobjectuse, newFormGroup, geoobjectuseFormCallback.probe)
		geoobjectuseFormCallback.probe.formStage.Commit()
	}

	geoobjectuseFormCallback.probe.ux_tree()
}
func __gong__New__GroupFormCallback(
	group *models.Group,
	probe *Probe,
	formGroup *form.FormGroup,
) (groupFormCallback *GroupFormCallback) {
	groupFormCallback = new(GroupFormCallback)
	groupFormCallback.probe = probe
	groupFormCallback.group = group
	groupFormCallback.formGroup = formGroup

	groupFormCallback.CreationMode = (group == nil)

	return
}

type GroupFormCallback struct {
	group *models.Group

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (groupFormCallback *GroupFormCallback) OnSave() {
	groupFormCallback.probe.stageOfInterest.Lock()
	defer groupFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("GroupFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	groupFormCallback.probe.formStage.Checkout()

	if groupFormCallback.group == nil {
		groupFormCallback.group = new(models.Group).Stage(groupFormCallback.probe.stageOfInterest)
	}
	group_ := groupFormCallback.group
	_ = group_

	for _, formDiv := range groupFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(group_.Name), formDiv)
		case "UserUse":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.UserUse](groupFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.UserUse, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.UserUse)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					groupFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.UserUse](groupFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			group_.UserUse = instanceSlice
			groupFormCallback.probe.UpdateSliceOfPointersCallback(group_, "UserUse", &group_.UserUse)

		case "ComputedPrefix":
			FormDivBasicFieldToField(&(group_.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(group_.IsExpanded), formDiv)
		case "LayoutDirection":
			FormDivEnumIntFieldToField(&(group_.LayoutDirection), formDiv)
		}
	}

	// manage the suppress operation
	if groupFormCallback.formGroup.HasSuppressButtonBeenPressed {
		group_.Unstage(groupFormCallback.probe.stageOfInterest)
	}

	groupFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Group](
		groupFormCallback.probe,
	)

	// display a new form by reset the form stage
	if groupFormCallback.CreationMode || groupFormCallback.formGroup.HasSuppressButtonBeenPressed {
		groupFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(groupFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__GroupFormCallback(
			nil,
			groupFormCallback.probe,
			newFormGroup,
		)
		group := new(models.Group)
		FillUpForm(group, newFormGroup, groupFormCallback.probe)
		groupFormCallback.probe.formStage.Commit()
	}

	groupFormCallback.probe.ux_tree()
}
func __gong__New__GroupUseFormCallback(
	groupuse *models.GroupUse,
	probe *Probe,
	formGroup *form.FormGroup,
) (groupuseFormCallback *GroupUseFormCallback) {
	groupuseFormCallback = new(GroupUseFormCallback)
	groupuseFormCallback.probe = probe
	groupuseFormCallback.groupuse = groupuse
	groupuseFormCallback.formGroup = formGroup

	groupuseFormCallback.CreationMode = (groupuse == nil)

	return
}

type GroupUseFormCallback struct {
	groupuse *models.GroupUse

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (groupuseFormCallback *GroupUseFormCallback) OnSave() {
	groupuseFormCallback.probe.stageOfInterest.Lock()
	defer groupuseFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("GroupUseFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	groupuseFormCallback.probe.formStage.Checkout()

	if groupuseFormCallback.groupuse == nil {
		groupuseFormCallback.groupuse = new(models.GroupUse).Stage(groupuseFormCallback.probe.stageOfInterest)
	}
	groupuse_ := groupuseFormCallback.groupuse
	_ = groupuse_

	for _, formDiv := range groupuseFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(groupuse_.Name), formDiv)
		case "Group":
			FormDivSelectFieldToField(&(groupuse_.Group), groupuseFormCallback.probe.stageOfInterest, formDiv)
		case "Analysis:GroupUse":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Analysis instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Analysis instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Analysis](groupuseFormCallback.probe.stageOfInterest)
			targetAnalysisIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetAnalysisIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Analysis instances and update their GroupUse slice
			for _analysis := range *models.GetGongstructInstancesSetFromPointerType[*models.Analysis](groupuseFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(groupuseFormCallback.probe.stageOfInterest, _analysis)
				
				// if Analysis is selected
				if targetAnalysisIDs[id] {
					// ensure groupuse_ is in _analysis.GroupUse
					found := false
					for _, _b := range _analysis.GroupUse {
						if _b == groupuse_ {
							found = true
							break
						}
					}
					if !found {
						_analysis.GroupUse = append(_analysis.GroupUse, groupuse_)
						groupuseFormCallback.probe.UpdateSliceOfPointersCallback(_analysis, "GroupUse", &_analysis.GroupUse)
					}
				} else {
					// ensure groupuse_ is NOT in _analysis.GroupUse
					idx := slices.Index(_analysis.GroupUse, groupuse_)
					if idx != -1 {
						_analysis.GroupUse = slices.Delete(_analysis.GroupUse, idx, idx+1)
						groupuseFormCallback.probe.UpdateSliceOfPointersCallback(_analysis, "GroupUse", &_analysis.GroupUse)
					}
				}
			}
		case "Parameter:GroupUse":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Parameter instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Parameter instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Parameter](groupuseFormCallback.probe.stageOfInterest)
			targetParameterIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetParameterIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Parameter instances and update their GroupUse slice
			for _parameter := range *models.GetGongstructInstancesSetFromPointerType[*models.Parameter](groupuseFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(groupuseFormCallback.probe.stageOfInterest, _parameter)
				
				// if Parameter is selected
				if targetParameterIDs[id] {
					// ensure groupuse_ is in _parameter.GroupUse
					found := false
					for _, _b := range _parameter.GroupUse {
						if _b == groupuse_ {
							found = true
							break
						}
					}
					if !found {
						_parameter.GroupUse = append(_parameter.GroupUse, groupuse_)
						groupuseFormCallback.probe.UpdateSliceOfPointersCallback(_parameter, "GroupUse", &_parameter.GroupUse)
					}
				} else {
					// ensure groupuse_ is NOT in _parameter.GroupUse
					idx := slices.Index(_parameter.GroupUse, groupuse_)
					if idx != -1 {
						_parameter.GroupUse = slices.Delete(_parameter.GroupUse, idx, idx+1)
						groupuseFormCallback.probe.UpdateSliceOfPointersCallback(_parameter, "GroupUse", &_parameter.GroupUse)
					}
				}
			}
		case "Repository:GroupUse":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Repository instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Repository instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Repository](groupuseFormCallback.probe.stageOfInterest)
			targetRepositoryIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetRepositoryIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Repository instances and update their GroupUse slice
			for _repository := range *models.GetGongstructInstancesSetFromPointerType[*models.Repository](groupuseFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(groupuseFormCallback.probe.stageOfInterest, _repository)
				
				// if Repository is selected
				if targetRepositoryIDs[id] {
					// ensure groupuse_ is in _repository.GroupUse
					found := false
					for _, _b := range _repository.GroupUse {
						if _b == groupuse_ {
							found = true
							break
						}
					}
					if !found {
						_repository.GroupUse = append(_repository.GroupUse, groupuse_)
						groupuseFormCallback.probe.UpdateSliceOfPointersCallback(_repository, "GroupUse", &_repository.GroupUse)
					}
				} else {
					// ensure groupuse_ is NOT in _repository.GroupUse
					idx := slices.Index(_repository.GroupUse, groupuse_)
					if idx != -1 {
						_repository.GroupUse = slices.Delete(_repository.GroupUse, idx, idx+1)
						groupuseFormCallback.probe.UpdateSliceOfPointersCallback(_repository, "GroupUse", &_repository.GroupUse)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if groupuseFormCallback.formGroup.HasSuppressButtonBeenPressed {
		groupuse_.Unstage(groupuseFormCallback.probe.stageOfInterest)
	}

	groupuseFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.GroupUse](
		groupuseFormCallback.probe,
	)

	// display a new form by reset the form stage
	if groupuseFormCallback.CreationMode || groupuseFormCallback.formGroup.HasSuppressButtonBeenPressed {
		groupuseFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(groupuseFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__GroupUseFormCallback(
			nil,
			groupuseFormCallback.probe,
			newFormGroup,
		)
		groupuse := new(models.GroupUse)
		FillUpForm(groupuse, newFormGroup, groupuseFormCallback.probe)
		groupuseFormCallback.probe.formStage.Commit()
	}

	groupuseFormCallback.probe.ux_tree()
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
		case "Description":
			FormDivBasicFieldToField(&(library_.Description), formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(library_.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(library_.IsExpanded), formDiv)
		case "LayoutDirection":
			FormDivEnumIntFieldToField(&(library_.LayoutDirection), formDiv)
		case "IsRootLibrary":
			FormDivBasicFieldToField(&(library_.IsRootLibrary), formDiv)
		case "Analyses":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Analysis](libraryFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Analysis, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Analysis)

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
			map_RowID_ID := GetMap_RowID_ID[*models.Analysis](libraryFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			library_.Analyses = instanceSlice
			libraryFormCallback.probe.UpdateSliceOfPointersCallback(library_, "Analyses", &library_.Analyses)

		case "IsAnalysesNodeExpanded":
			FormDivBasicFieldToField(&(library_.IsAnalysesNodeExpanded), formDiv)
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

		case "NbPixPerCharacter":
			FormDivBasicFieldToField(&(library_.NbPixPerCharacter), formDiv)
		case "LogoSVGFile":
			FormDivBasicFieldToField(&(library_.LogoSVGFile), formDiv)
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
func __gong__New__MapObjectFormCallback(
	mapobject *models.MapObject,
	probe *Probe,
	formGroup *form.FormGroup,
) (mapobjectFormCallback *MapObjectFormCallback) {
	mapobjectFormCallback = new(MapObjectFormCallback)
	mapobjectFormCallback.probe = probe
	mapobjectFormCallback.mapobject = mapobject
	mapobjectFormCallback.formGroup = formGroup

	mapobjectFormCallback.CreationMode = (mapobject == nil)

	return
}

type MapObjectFormCallback struct {
	mapobject *models.MapObject

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (mapobjectFormCallback *MapObjectFormCallback) OnSave() {
	mapobjectFormCallback.probe.stageOfInterest.Lock()
	defer mapobjectFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("MapObjectFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	mapobjectFormCallback.probe.formStage.Checkout()

	if mapobjectFormCallback.mapobject == nil {
		mapobjectFormCallback.mapobject = new(models.MapObject).Stage(mapobjectFormCallback.probe.stageOfInterest)
	}
	mapobject_ := mapobjectFormCallback.mapobject
	_ = mapobject_

	for _, formDiv := range mapobjectFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(mapobject_.Name), formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(mapobject_.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(mapobject_.IsExpanded), formDiv)
		case "LayoutDirection":
			FormDivEnumIntFieldToField(&(mapobject_.LayoutDirection), formDiv)
		}
	}

	// manage the suppress operation
	if mapobjectFormCallback.formGroup.HasSuppressButtonBeenPressed {
		mapobject_.Unstage(mapobjectFormCallback.probe.stageOfInterest)
	}

	mapobjectFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.MapObject](
		mapobjectFormCallback.probe,
	)

	// display a new form by reset the form stage
	if mapobjectFormCallback.CreationMode || mapobjectFormCallback.formGroup.HasSuppressButtonBeenPressed {
		mapobjectFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(mapobjectFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__MapObjectFormCallback(
			nil,
			mapobjectFormCallback.probe,
			newFormGroup,
		)
		mapobject := new(models.MapObject)
		FillUpForm(mapobject, newFormGroup, mapobjectFormCallback.probe)
		mapobjectFormCallback.probe.formStage.Commit()
	}

	mapobjectFormCallback.probe.ux_tree()
}
func __gong__New__MapObjectUseFormCallback(
	mapobjectuse *models.MapObjectUse,
	probe *Probe,
	formGroup *form.FormGroup,
) (mapobjectuseFormCallback *MapObjectUseFormCallback) {
	mapobjectuseFormCallback = new(MapObjectUseFormCallback)
	mapobjectuseFormCallback.probe = probe
	mapobjectuseFormCallback.mapobjectuse = mapobjectuse
	mapobjectuseFormCallback.formGroup = formGroup

	mapobjectuseFormCallback.CreationMode = (mapobjectuse == nil)

	return
}

type MapObjectUseFormCallback struct {
	mapobjectuse *models.MapObjectUse

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (mapobjectuseFormCallback *MapObjectUseFormCallback) OnSave() {
	mapobjectuseFormCallback.probe.stageOfInterest.Lock()
	defer mapobjectuseFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("MapObjectUseFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	mapobjectuseFormCallback.probe.formStage.Checkout()

	if mapobjectuseFormCallback.mapobjectuse == nil {
		mapobjectuseFormCallback.mapobjectuse = new(models.MapObjectUse).Stage(mapobjectuseFormCallback.probe.stageOfInterest)
	}
	mapobjectuse_ := mapobjectuseFormCallback.mapobjectuse
	_ = mapobjectuse_

	for _, formDiv := range mapobjectuseFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(mapobjectuse_.Name), formDiv)
		case "Map":
			FormDivSelectFieldToField(&(mapobjectuse_.Map), mapobjectuseFormCallback.probe.stageOfInterest, formDiv)
		case "Analysis:MapUse":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Analysis instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Analysis instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Analysis](mapobjectuseFormCallback.probe.stageOfInterest)
			targetAnalysisIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetAnalysisIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Analysis instances and update their MapUse slice
			for _analysis := range *models.GetGongstructInstancesSetFromPointerType[*models.Analysis](mapobjectuseFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(mapobjectuseFormCallback.probe.stageOfInterest, _analysis)
				
				// if Analysis is selected
				if targetAnalysisIDs[id] {
					// ensure mapobjectuse_ is in _analysis.MapUse
					found := false
					for _, _b := range _analysis.MapUse {
						if _b == mapobjectuse_ {
							found = true
							break
						}
					}
					if !found {
						_analysis.MapUse = append(_analysis.MapUse, mapobjectuse_)
						mapobjectuseFormCallback.probe.UpdateSliceOfPointersCallback(_analysis, "MapUse", &_analysis.MapUse)
					}
				} else {
					// ensure mapobjectuse_ is NOT in _analysis.MapUse
					idx := slices.Index(_analysis.MapUse, mapobjectuse_)
					if idx != -1 {
						_analysis.MapUse = slices.Delete(_analysis.MapUse, idx, idx+1)
						mapobjectuseFormCallback.probe.UpdateSliceOfPointersCallback(_analysis, "MapUse", &_analysis.MapUse)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if mapobjectuseFormCallback.formGroup.HasSuppressButtonBeenPressed {
		mapobjectuse_.Unstage(mapobjectuseFormCallback.probe.stageOfInterest)
	}

	mapobjectuseFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.MapObjectUse](
		mapobjectuseFormCallback.probe,
	)

	// display a new form by reset the form stage
	if mapobjectuseFormCallback.CreationMode || mapobjectuseFormCallback.formGroup.HasSuppressButtonBeenPressed {
		mapobjectuseFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(mapobjectuseFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__MapObjectUseFormCallback(
			nil,
			mapobjectuseFormCallback.probe,
			newFormGroup,
		)
		mapobjectuse := new(models.MapObjectUse)
		FillUpForm(mapobjectuse, newFormGroup, mapobjectuseFormCallback.probe)
		mapobjectuseFormCallback.probe.formStage.Commit()
	}

	mapobjectuseFormCallback.probe.ux_tree()
}
func __gong__New__ParameterFormCallback(
	parameter *models.Parameter,
	probe *Probe,
	formGroup *form.FormGroup,
) (parameterFormCallback *ParameterFormCallback) {
	parameterFormCallback = new(ParameterFormCallback)
	parameterFormCallback.probe = probe
	parameterFormCallback.parameter = parameter
	parameterFormCallback.formGroup = formGroup

	parameterFormCallback.CreationMode = (parameter == nil)

	return
}

type ParameterFormCallback struct {
	parameter *models.Parameter

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (parameterFormCallback *ParameterFormCallback) OnSave() {
	parameterFormCallback.probe.stageOfInterest.Lock()
	defer parameterFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("ParameterFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	parameterFormCallback.probe.formStage.Checkout()

	if parameterFormCallback.parameter == nil {
		parameterFormCallback.parameter = new(models.Parameter).Stage(parameterFormCallback.probe.stageOfInterest)
	}
	parameter_ := parameterFormCallback.parameter
	_ = parameter_

	for _, formDiv := range parameterFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(parameter_.Name), formDiv)
		case "Description":
			FormDivBasicFieldToField(&(parameter_.Description), formDiv)
		case "IsResponse":
			FormDivBasicFieldToField(&(parameter_.IsResponse), formDiv)
		case "Start":
			FormDivBasicFieldToField(&(parameter_.Start), formDiv)
		case "End":
			FormDivBasicFieldToField(&(parameter_.End), formDiv)
		case "Force":
			FormDivBasicFieldToField(&(parameter_.Force), formDiv)
		case "GroupUse":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.GroupUse](parameterFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.GroupUse, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.GroupUse)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					parameterFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.GroupUse](parameterFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			parameter_.GroupUse = instanceSlice
			parameterFormCallback.probe.UpdateSliceOfPointersCallback(parameter_, "GroupUse", &parameter_.GroupUse)

		case "DocumentUse":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.DocumentUse](parameterFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.DocumentUse, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.DocumentUse)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					parameterFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.DocumentUse](parameterFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			parameter_.DocumentUse = instanceSlice
			parameterFormCallback.probe.UpdateSliceOfPointersCallback(parameter_, "DocumentUse", &parameter_.DocumentUse)

		case "GeoObjectUse":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.GeoObjectUse](parameterFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.GeoObjectUse, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.GeoObjectUse)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					parameterFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.GeoObjectUse](parameterFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			parameter_.GeoObjectUse = instanceSlice
			parameterFormCallback.probe.UpdateSliceOfPointersCallback(parameter_, "GeoObjectUse", &parameter_.GeoObjectUse)

		case "Tag":
			FormDivBasicFieldToField(&(parameter_.Tag), formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(parameter_.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(parameter_.IsExpanded), formDiv)
		case "LayoutDirection":
			FormDivEnumIntFieldToField(&(parameter_.LayoutDirection), formDiv)
		case "ActorStateTransition:Justifications":
			// 1. Decode the AssociationStorage which contains the rowIDs of the ActorStateTransition instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target ActorStateTransition instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.ActorStateTransition](parameterFormCallback.probe.stageOfInterest)
			targetActorStateTransitionIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetActorStateTransitionIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all ActorStateTransition instances and update their Justifications slice
			for _actorstatetransition := range *models.GetGongstructInstancesSetFromPointerType[*models.ActorStateTransition](parameterFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(parameterFormCallback.probe.stageOfInterest, _actorstatetransition)
				
				// if ActorStateTransition is selected
				if targetActorStateTransitionIDs[id] {
					// ensure parameter_ is in _actorstatetransition.Justifications
					found := false
					for _, _b := range _actorstatetransition.Justifications {
						if _b == parameter_ {
							found = true
							break
						}
					}
					if !found {
						_actorstatetransition.Justifications = append(_actorstatetransition.Justifications, parameter_)
						parameterFormCallback.probe.UpdateSliceOfPointersCallback(_actorstatetransition, "Justifications", &_actorstatetransition.Justifications)
					}
				} else {
					// ensure parameter_ is NOT in _actorstatetransition.Justifications
					idx := slices.Index(_actorstatetransition.Justifications, parameter_)
					if idx != -1 {
						_actorstatetransition.Justifications = slices.Delete(_actorstatetransition.Justifications, idx, idx+1)
						parameterFormCallback.probe.UpdateSliceOfPointersCallback(_actorstatetransition, "Justifications", &_actorstatetransition.Justifications)
					}
				}
			}
		case "Diagram:ParametersWhoseNodeIsExpanded":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Diagram instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Diagram instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Diagram](parameterFormCallback.probe.stageOfInterest)
			targetDiagramIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Diagram instances and update their ParametersWhoseNodeIsExpanded slice
			for _diagram := range *models.GetGongstructInstancesSetFromPointerType[*models.Diagram](parameterFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(parameterFormCallback.probe.stageOfInterest, _diagram)
				
				// if Diagram is selected
				if targetDiagramIDs[id] {
					// ensure parameter_ is in _diagram.ParametersWhoseNodeIsExpanded
					found := false
					for _, _b := range _diagram.ParametersWhoseNodeIsExpanded {
						if _b == parameter_ {
							found = true
							break
						}
					}
					if !found {
						_diagram.ParametersWhoseNodeIsExpanded = append(_diagram.ParametersWhoseNodeIsExpanded, parameter_)
						parameterFormCallback.probe.UpdateSliceOfPointersCallback(_diagram, "ParametersWhoseNodeIsExpanded", &_diagram.ParametersWhoseNodeIsExpanded)
					}
				} else {
					// ensure parameter_ is NOT in _diagram.ParametersWhoseNodeIsExpanded
					idx := slices.Index(_diagram.ParametersWhoseNodeIsExpanded, parameter_)
					if idx != -1 {
						_diagram.ParametersWhoseNodeIsExpanded = slices.Delete(_diagram.ParametersWhoseNodeIsExpanded, idx, idx+1)
						parameterFormCallback.probe.UpdateSliceOfPointersCallback(_diagram, "ParametersWhoseNodeIsExpanded", &_diagram.ParametersWhoseNodeIsExpanded)
					}
				}
			}
		case "ParametersAggregate:Parameters":
			// 1. Decode the AssociationStorage which contains the rowIDs of the ParametersAggregate instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target ParametersAggregate instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.ParametersAggregate](parameterFormCallback.probe.stageOfInterest)
			targetParametersAggregateIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetParametersAggregateIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all ParametersAggregate instances and update their Parameters slice
			for _parametersaggregate := range *models.GetGongstructInstancesSetFromPointerType[*models.ParametersAggregate](parameterFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(parameterFormCallback.probe.stageOfInterest, _parametersaggregate)
				
				// if ParametersAggregate is selected
				if targetParametersAggregateIDs[id] {
					// ensure parameter_ is in _parametersaggregate.Parameters
					found := false
					for _, _b := range _parametersaggregate.Parameters {
						if _b == parameter_ {
							found = true
							break
						}
					}
					if !found {
						_parametersaggregate.Parameters = append(_parametersaggregate.Parameters, parameter_)
						parameterFormCallback.probe.UpdateSliceOfPointersCallback(_parametersaggregate, "Parameters", &_parametersaggregate.Parameters)
					}
				} else {
					// ensure parameter_ is NOT in _parametersaggregate.Parameters
					idx := slices.Index(_parametersaggregate.Parameters, parameter_)
					if idx != -1 {
						_parametersaggregate.Parameters = slices.Delete(_parametersaggregate.Parameters, idx, idx+1)
						parameterFormCallback.probe.UpdateSliceOfPointersCallback(_parametersaggregate, "Parameters", &_parametersaggregate.Parameters)
					}
				}
			}
		case "Scenario:Parameters":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Scenario instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Scenario instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Scenario](parameterFormCallback.probe.stageOfInterest)
			targetScenarioIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetScenarioIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Scenario instances and update their Parameters slice
			for _scenario := range *models.GetGongstructInstancesSetFromPointerType[*models.Scenario](parameterFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(parameterFormCallback.probe.stageOfInterest, _scenario)
				
				// if Scenario is selected
				if targetScenarioIDs[id] {
					// ensure parameter_ is in _scenario.Parameters
					found := false
					for _, _b := range _scenario.Parameters {
						if _b == parameter_ {
							found = true
							break
						}
					}
					if !found {
						_scenario.Parameters = append(_scenario.Parameters, parameter_)
						parameterFormCallback.probe.UpdateSliceOfPointersCallback(_scenario, "Parameters", &_scenario.Parameters)
					}
				} else {
					// ensure parameter_ is NOT in _scenario.Parameters
					idx := slices.Index(_scenario.Parameters, parameter_)
					if idx != -1 {
						_scenario.Parameters = slices.Delete(_scenario.Parameters, idx, idx+1)
						parameterFormCallback.probe.UpdateSliceOfPointersCallback(_scenario, "Parameters", &_scenario.Parameters)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if parameterFormCallback.formGroup.HasSuppressButtonBeenPressed {
		parameter_.Unstage(parameterFormCallback.probe.stageOfInterest)
	}

	parameterFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Parameter](
		parameterFormCallback.probe,
	)

	// display a new form by reset the form stage
	if parameterFormCallback.CreationMode || parameterFormCallback.formGroup.HasSuppressButtonBeenPressed {
		parameterFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(parameterFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ParameterFormCallback(
			nil,
			parameterFormCallback.probe,
			newFormGroup,
		)
		parameter := new(models.Parameter)
		FillUpForm(parameter, newFormGroup, parameterFormCallback.probe)
		parameterFormCallback.probe.formStage.Commit()
	}

	parameterFormCallback.probe.ux_tree()
}
func __gong__New__ParameterCategoryFormCallback(
	parametercategory *models.ParameterCategory,
	probe *Probe,
	formGroup *form.FormGroup,
) (parametercategoryFormCallback *ParameterCategoryFormCallback) {
	parametercategoryFormCallback = new(ParameterCategoryFormCallback)
	parametercategoryFormCallback.probe = probe
	parametercategoryFormCallback.parametercategory = parametercategory
	parametercategoryFormCallback.formGroup = formGroup

	parametercategoryFormCallback.CreationMode = (parametercategory == nil)

	return
}

type ParameterCategoryFormCallback struct {
	parametercategory *models.ParameterCategory

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (parametercategoryFormCallback *ParameterCategoryFormCallback) OnSave() {
	parametercategoryFormCallback.probe.stageOfInterest.Lock()
	defer parametercategoryFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("ParameterCategoryFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	parametercategoryFormCallback.probe.formStage.Checkout()

	if parametercategoryFormCallback.parametercategory == nil {
		parametercategoryFormCallback.parametercategory = new(models.ParameterCategory).Stage(parametercategoryFormCallback.probe.stageOfInterest)
	}
	parametercategory_ := parametercategoryFormCallback.parametercategory
	_ = parametercategory_

	for _, formDiv := range parametercategoryFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(parametercategory_.Name), formDiv)
		case "ParameterUse":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.ParameterShape](parametercategoryFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.ParameterShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.ParameterShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					parametercategoryFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.ParameterShape](parametercategoryFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			parametercategory_.ParameterUse = instanceSlice
			parametercategoryFormCallback.probe.UpdateSliceOfPointersCallback(parametercategory_, "ParameterUse", &parametercategory_.ParameterUse)

		case "ComputedPrefix":
			FormDivBasicFieldToField(&(parametercategory_.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(parametercategory_.IsExpanded), formDiv)
		case "LayoutDirection":
			FormDivEnumIntFieldToField(&(parametercategory_.LayoutDirection), formDiv)
		}
	}

	// manage the suppress operation
	if parametercategoryFormCallback.formGroup.HasSuppressButtonBeenPressed {
		parametercategory_.Unstage(parametercategoryFormCallback.probe.stageOfInterest)
	}

	parametercategoryFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.ParameterCategory](
		parametercategoryFormCallback.probe,
	)

	// display a new form by reset the form stage
	if parametercategoryFormCallback.CreationMode || parametercategoryFormCallback.formGroup.HasSuppressButtonBeenPressed {
		parametercategoryFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(parametercategoryFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ParameterCategoryFormCallback(
			nil,
			parametercategoryFormCallback.probe,
			newFormGroup,
		)
		parametercategory := new(models.ParameterCategory)
		FillUpForm(parametercategory, newFormGroup, parametercategoryFormCallback.probe)
		parametercategoryFormCallback.probe.formStage.Commit()
	}

	parametercategoryFormCallback.probe.ux_tree()
}
func __gong__New__ParameterCategoryUseFormCallback(
	parametercategoryuse *models.ParameterCategoryUse,
	probe *Probe,
	formGroup *form.FormGroup,
) (parametercategoryuseFormCallback *ParameterCategoryUseFormCallback) {
	parametercategoryuseFormCallback = new(ParameterCategoryUseFormCallback)
	parametercategoryuseFormCallback.probe = probe
	parametercategoryuseFormCallback.parametercategoryuse = parametercategoryuse
	parametercategoryuseFormCallback.formGroup = formGroup

	parametercategoryuseFormCallback.CreationMode = (parametercategoryuse == nil)

	return
}

type ParameterCategoryUseFormCallback struct {
	parametercategoryuse *models.ParameterCategoryUse

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (parametercategoryuseFormCallback *ParameterCategoryUseFormCallback) OnSave() {
	parametercategoryuseFormCallback.probe.stageOfInterest.Lock()
	defer parametercategoryuseFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("ParameterCategoryUseFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	parametercategoryuseFormCallback.probe.formStage.Checkout()

	if parametercategoryuseFormCallback.parametercategoryuse == nil {
		parametercategoryuseFormCallback.parametercategoryuse = new(models.ParameterCategoryUse).Stage(parametercategoryuseFormCallback.probe.stageOfInterest)
	}
	parametercategoryuse_ := parametercategoryuseFormCallback.parametercategoryuse
	_ = parametercategoryuse_

	for _, formDiv := range parametercategoryuseFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(parametercategoryuse_.Name), formDiv)
		case "ParameterCategory":
			FormDivSelectFieldToField(&(parametercategoryuse_.ParameterCategory), parametercategoryuseFormCallback.probe.stageOfInterest, formDiv)
		}
	}

	// manage the suppress operation
	if parametercategoryuseFormCallback.formGroup.HasSuppressButtonBeenPressed {
		parametercategoryuse_.Unstage(parametercategoryuseFormCallback.probe.stageOfInterest)
	}

	parametercategoryuseFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.ParameterCategoryUse](
		parametercategoryuseFormCallback.probe,
	)

	// display a new form by reset the form stage
	if parametercategoryuseFormCallback.CreationMode || parametercategoryuseFormCallback.formGroup.HasSuppressButtonBeenPressed {
		parametercategoryuseFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(parametercategoryuseFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ParameterCategoryUseFormCallback(
			nil,
			parametercategoryuseFormCallback.probe,
			newFormGroup,
		)
		parametercategoryuse := new(models.ParameterCategoryUse)
		FillUpForm(parametercategoryuse, newFormGroup, parametercategoryuseFormCallback.probe)
		parametercategoryuseFormCallback.probe.formStage.Commit()
	}

	parametercategoryuseFormCallback.probe.ux_tree()
}
func __gong__New__ParameterShapeFormCallback(
	parametershape *models.ParameterShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (parametershapeFormCallback *ParameterShapeFormCallback) {
	parametershapeFormCallback = new(ParameterShapeFormCallback)
	parametershapeFormCallback.probe = probe
	parametershapeFormCallback.parametershape = parametershape
	parametershapeFormCallback.formGroup = formGroup

	parametershapeFormCallback.CreationMode = (parametershape == nil)

	return
}

type ParameterShapeFormCallback struct {
	parametershape *models.ParameterShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (parametershapeFormCallback *ParameterShapeFormCallback) OnSave() {
	parametershapeFormCallback.probe.stageOfInterest.Lock()
	defer parametershapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("ParameterShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	parametershapeFormCallback.probe.formStage.Checkout()

	if parametershapeFormCallback.parametershape == nil {
		parametershapeFormCallback.parametershape = new(models.ParameterShape).Stage(parametershapeFormCallback.probe.stageOfInterest)
	}
	parametershape_ := parametershapeFormCallback.parametershape
	_ = parametershape_

	for _, formDiv := range parametershapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(parametershape_.Name), formDiv)
		case "Parameter":
			FormDivSelectFieldToField(&(parametershape_.Parameter), parametershapeFormCallback.probe.stageOfInterest, formDiv)
		case "Direction":
			FormDivEnumStringFieldToField(&(parametershape_.Direction), formDiv)
		case "ShapeIsComputedFromModel":
			FormDivBasicFieldToField(&(parametershape_.ShapeIsComputedFromModel), formDiv)
		case "X":
			FormDivBasicFieldToField(&(parametershape_.X), formDiv)
		case "Y":
			FormDivBasicFieldToField(&(parametershape_.Y), formDiv)
		case "Width":
			FormDivBasicFieldToField(&(parametershape_.Width), formDiv)
		case "Height":
			FormDivBasicFieldToField(&(parametershape_.Height), formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&(parametershape_.IsHidden), formDiv)
		case "Diagram:ParameterShapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Diagram instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Diagram instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Diagram](parametershapeFormCallback.probe.stageOfInterest)
			targetDiagramIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Diagram instances and update their ParameterShapes slice
			for _diagram := range *models.GetGongstructInstancesSetFromPointerType[*models.Diagram](parametershapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(parametershapeFormCallback.probe.stageOfInterest, _diagram)
				
				// if Diagram is selected
				if targetDiagramIDs[id] {
					// ensure parametershape_ is in _diagram.ParameterShapes
					found := false
					for _, _b := range _diagram.ParameterShapes {
						if _b == parametershape_ {
							found = true
							break
						}
					}
					if !found {
						_diagram.ParameterShapes = append(_diagram.ParameterShapes, parametershape_)
						parametershapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagram, "ParameterShapes", &_diagram.ParameterShapes)
					}
				} else {
					// ensure parametershape_ is NOT in _diagram.ParameterShapes
					idx := slices.Index(_diagram.ParameterShapes, parametershape_)
					if idx != -1 {
						_diagram.ParameterShapes = slices.Delete(_diagram.ParameterShapes, idx, idx+1)
						parametershapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagram, "ParameterShapes", &_diagram.ParameterShapes)
					}
				}
			}
		case "ParameterCategory:ParameterUse":
			// 1. Decode the AssociationStorage which contains the rowIDs of the ParameterCategory instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target ParameterCategory instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.ParameterCategory](parametershapeFormCallback.probe.stageOfInterest)
			targetParameterCategoryIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetParameterCategoryIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all ParameterCategory instances and update their ParameterUse slice
			for _parametercategory := range *models.GetGongstructInstancesSetFromPointerType[*models.ParameterCategory](parametershapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(parametershapeFormCallback.probe.stageOfInterest, _parametercategory)
				
				// if ParameterCategory is selected
				if targetParameterCategoryIDs[id] {
					// ensure parametershape_ is in _parametercategory.ParameterUse
					found := false
					for _, _b := range _parametercategory.ParameterUse {
						if _b == parametershape_ {
							found = true
							break
						}
					}
					if !found {
						_parametercategory.ParameterUse = append(_parametercategory.ParameterUse, parametershape_)
						parametershapeFormCallback.probe.UpdateSliceOfPointersCallback(_parametercategory, "ParameterUse", &_parametercategory.ParameterUse)
					}
				} else {
					// ensure parametershape_ is NOT in _parametercategory.ParameterUse
					idx := slices.Index(_parametercategory.ParameterUse, parametershape_)
					if idx != -1 {
						_parametercategory.ParameterUse = slices.Delete(_parametercategory.ParameterUse, idx, idx+1)
						parametershapeFormCallback.probe.UpdateSliceOfPointersCallback(_parametercategory, "ParameterUse", &_parametercategory.ParameterUse)
					}
				}
			}
		case "Repository:ParameterUse":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Repository instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Repository instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Repository](parametershapeFormCallback.probe.stageOfInterest)
			targetRepositoryIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetRepositoryIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Repository instances and update their ParameterUse slice
			for _repository := range *models.GetGongstructInstancesSetFromPointerType[*models.Repository](parametershapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(parametershapeFormCallback.probe.stageOfInterest, _repository)
				
				// if Repository is selected
				if targetRepositoryIDs[id] {
					// ensure parametershape_ is in _repository.ParameterUse
					found := false
					for _, _b := range _repository.ParameterUse {
						if _b == parametershape_ {
							found = true
							break
						}
					}
					if !found {
						_repository.ParameterUse = append(_repository.ParameterUse, parametershape_)
						parametershapeFormCallback.probe.UpdateSliceOfPointersCallback(_repository, "ParameterUse", &_repository.ParameterUse)
					}
				} else {
					// ensure parametershape_ is NOT in _repository.ParameterUse
					idx := slices.Index(_repository.ParameterUse, parametershape_)
					if idx != -1 {
						_repository.ParameterUse = slices.Delete(_repository.ParameterUse, idx, idx+1)
						parametershapeFormCallback.probe.UpdateSliceOfPointersCallback(_repository, "ParameterUse", &_repository.ParameterUse)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if parametershapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		parametershape_.Unstage(parametershapeFormCallback.probe.stageOfInterest)
	}

	parametershapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.ParameterShape](
		parametershapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if parametershapeFormCallback.CreationMode || parametershapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		parametershapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(parametershapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ParameterShapeFormCallback(
			nil,
			parametershapeFormCallback.probe,
			newFormGroup,
		)
		parametershape := new(models.ParameterShape)
		FillUpForm(parametershape, newFormGroup, parametershapeFormCallback.probe)
		parametershapeFormCallback.probe.formStage.Commit()
	}

	parametershapeFormCallback.probe.ux_tree()
}
func __gong__New__ParametersAggregateFormCallback(
	parametersaggregate *models.ParametersAggregate,
	probe *Probe,
	formGroup *form.FormGroup,
) (parametersaggregateFormCallback *ParametersAggregateFormCallback) {
	parametersaggregateFormCallback = new(ParametersAggregateFormCallback)
	parametersaggregateFormCallback.probe = probe
	parametersaggregateFormCallback.parametersaggregate = parametersaggregate
	parametersaggregateFormCallback.formGroup = formGroup

	parametersaggregateFormCallback.CreationMode = (parametersaggregate == nil)

	return
}

type ParametersAggregateFormCallback struct {
	parametersaggregate *models.ParametersAggregate

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (parametersaggregateFormCallback *ParametersAggregateFormCallback) OnSave() {
	parametersaggregateFormCallback.probe.stageOfInterest.Lock()
	defer parametersaggregateFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("ParametersAggregateFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	parametersaggregateFormCallback.probe.formStage.Checkout()

	if parametersaggregateFormCallback.parametersaggregate == nil {
		parametersaggregateFormCallback.parametersaggregate = new(models.ParametersAggregate).Stage(parametersaggregateFormCallback.probe.stageOfInterest)
	}
	parametersaggregate_ := parametersaggregateFormCallback.parametersaggregate
	_ = parametersaggregate_

	for _, formDiv := range parametersaggregateFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(parametersaggregate_.Name), formDiv)
		case "Tag":
			FormDivBasicFieldToField(&(parametersaggregate_.Tag), formDiv)
		case "Description":
			FormDivBasicFieldToField(&(parametersaggregate_.Description), formDiv)
		case "Parameters":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Parameter](parametersaggregateFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Parameter, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Parameter)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					parametersaggregateFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Parameter](parametersaggregateFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			parametersaggregate_.Parameters = instanceSlice
			parametersaggregateFormCallback.probe.UpdateSliceOfPointersCallback(parametersaggregate_, "Parameters", &parametersaggregate_.Parameters)

		case "ComputedPrefix":
			FormDivBasicFieldToField(&(parametersaggregate_.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(parametersaggregate_.IsExpanded), formDiv)
		case "LayoutDirection":
			FormDivEnumIntFieldToField(&(parametersaggregate_.LayoutDirection), formDiv)
		case "Diagram:ParametersAggregatesWhoseNodeIsExpanded":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Diagram instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Diagram instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Diagram](parametersaggregateFormCallback.probe.stageOfInterest)
			targetDiagramIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Diagram instances and update their ParametersAggregatesWhoseNodeIsExpanded slice
			for _diagram := range *models.GetGongstructInstancesSetFromPointerType[*models.Diagram](parametersaggregateFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(parametersaggregateFormCallback.probe.stageOfInterest, _diagram)
				
				// if Diagram is selected
				if targetDiagramIDs[id] {
					// ensure parametersaggregate_ is in _diagram.ParametersAggregatesWhoseNodeIsExpanded
					found := false
					for _, _b := range _diagram.ParametersAggregatesWhoseNodeIsExpanded {
						if _b == parametersaggregate_ {
							found = true
							break
						}
					}
					if !found {
						_diagram.ParametersAggregatesWhoseNodeIsExpanded = append(_diagram.ParametersAggregatesWhoseNodeIsExpanded, parametersaggregate_)
						parametersaggregateFormCallback.probe.UpdateSliceOfPointersCallback(_diagram, "ParametersAggregatesWhoseNodeIsExpanded", &_diagram.ParametersAggregatesWhoseNodeIsExpanded)
					}
				} else {
					// ensure parametersaggregate_ is NOT in _diagram.ParametersAggregatesWhoseNodeIsExpanded
					idx := slices.Index(_diagram.ParametersAggregatesWhoseNodeIsExpanded, parametersaggregate_)
					if idx != -1 {
						_diagram.ParametersAggregatesWhoseNodeIsExpanded = slices.Delete(_diagram.ParametersAggregatesWhoseNodeIsExpanded, idx, idx+1)
						parametersaggregateFormCallback.probe.UpdateSliceOfPointersCallback(_diagram, "ParametersAggregatesWhoseNodeIsExpanded", &_diagram.ParametersAggregatesWhoseNodeIsExpanded)
					}
				}
			}
		case "Scenario:ParametersAggretates":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Scenario instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Scenario instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Scenario](parametersaggregateFormCallback.probe.stageOfInterest)
			targetScenarioIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetScenarioIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Scenario instances and update their ParametersAggretates slice
			for _scenario := range *models.GetGongstructInstancesSetFromPointerType[*models.Scenario](parametersaggregateFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(parametersaggregateFormCallback.probe.stageOfInterest, _scenario)
				
				// if Scenario is selected
				if targetScenarioIDs[id] {
					// ensure parametersaggregate_ is in _scenario.ParametersAggretates
					found := false
					for _, _b := range _scenario.ParametersAggretates {
						if _b == parametersaggregate_ {
							found = true
							break
						}
					}
					if !found {
						_scenario.ParametersAggretates = append(_scenario.ParametersAggretates, parametersaggregate_)
						parametersaggregateFormCallback.probe.UpdateSliceOfPointersCallback(_scenario, "ParametersAggretates", &_scenario.ParametersAggretates)
					}
				} else {
					// ensure parametersaggregate_ is NOT in _scenario.ParametersAggretates
					idx := slices.Index(_scenario.ParametersAggretates, parametersaggregate_)
					if idx != -1 {
						_scenario.ParametersAggretates = slices.Delete(_scenario.ParametersAggretates, idx, idx+1)
						parametersaggregateFormCallback.probe.UpdateSliceOfPointersCallback(_scenario, "ParametersAggretates", &_scenario.ParametersAggretates)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if parametersaggregateFormCallback.formGroup.HasSuppressButtonBeenPressed {
		parametersaggregate_.Unstage(parametersaggregateFormCallback.probe.stageOfInterest)
	}

	parametersaggregateFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.ParametersAggregate](
		parametersaggregateFormCallback.probe,
	)

	// display a new form by reset the form stage
	if parametersaggregateFormCallback.CreationMode || parametersaggregateFormCallback.formGroup.HasSuppressButtonBeenPressed {
		parametersaggregateFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(parametersaggregateFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ParametersAggregateFormCallback(
			nil,
			parametersaggregateFormCallback.probe,
			newFormGroup,
		)
		parametersaggregate := new(models.ParametersAggregate)
		FillUpForm(parametersaggregate, newFormGroup, parametersaggregateFormCallback.probe)
		parametersaggregateFormCallback.probe.formStage.Commit()
	}

	parametersaggregateFormCallback.probe.ux_tree()
}
func __gong__New__ParametersAggregateShapeFormCallback(
	parametersaggregateshape *models.ParametersAggregateShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (parametersaggregateshapeFormCallback *ParametersAggregateShapeFormCallback) {
	parametersaggregateshapeFormCallback = new(ParametersAggregateShapeFormCallback)
	parametersaggregateshapeFormCallback.probe = probe
	parametersaggregateshapeFormCallback.parametersaggregateshape = parametersaggregateshape
	parametersaggregateshapeFormCallback.formGroup = formGroup

	parametersaggregateshapeFormCallback.CreationMode = (parametersaggregateshape == nil)

	return
}

type ParametersAggregateShapeFormCallback struct {
	parametersaggregateshape *models.ParametersAggregateShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (parametersaggregateshapeFormCallback *ParametersAggregateShapeFormCallback) OnSave() {
	parametersaggregateshapeFormCallback.probe.stageOfInterest.Lock()
	defer parametersaggregateshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("ParametersAggregateShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	parametersaggregateshapeFormCallback.probe.formStage.Checkout()

	if parametersaggregateshapeFormCallback.parametersaggregateshape == nil {
		parametersaggregateshapeFormCallback.parametersaggregateshape = new(models.ParametersAggregateShape).Stage(parametersaggregateshapeFormCallback.probe.stageOfInterest)
	}
	parametersaggregateshape_ := parametersaggregateshapeFormCallback.parametersaggregateshape
	_ = parametersaggregateshape_

	for _, formDiv := range parametersaggregateshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(parametersaggregateshape_.Name), formDiv)
		case "ScenarioParameter":
			FormDivSelectFieldToField(&(parametersaggregateshape_.ScenarioParameter), parametersaggregateshapeFormCallback.probe.stageOfInterest, formDiv)
		case "Direction":
			FormDivEnumStringFieldToField(&(parametersaggregateshape_.Direction), formDiv)
		case "X":
			FormDivBasicFieldToField(&(parametersaggregateshape_.X), formDiv)
		case "Y":
			FormDivBasicFieldToField(&(parametersaggregateshape_.Y), formDiv)
		case "Width":
			FormDivBasicFieldToField(&(parametersaggregateshape_.Width), formDiv)
		case "Height":
			FormDivBasicFieldToField(&(parametersaggregateshape_.Height), formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&(parametersaggregateshape_.IsHidden), formDiv)
		case "Diagram:ScenarioParameterShapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Diagram instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Diagram instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Diagram](parametersaggregateshapeFormCallback.probe.stageOfInterest)
			targetDiagramIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Diagram instances and update their ScenarioParameterShapes slice
			for _diagram := range *models.GetGongstructInstancesSetFromPointerType[*models.Diagram](parametersaggregateshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(parametersaggregateshapeFormCallback.probe.stageOfInterest, _diagram)
				
				// if Diagram is selected
				if targetDiagramIDs[id] {
					// ensure parametersaggregateshape_ is in _diagram.ScenarioParameterShapes
					found := false
					for _, _b := range _diagram.ScenarioParameterShapes {
						if _b == parametersaggregateshape_ {
							found = true
							break
						}
					}
					if !found {
						_diagram.ScenarioParameterShapes = append(_diagram.ScenarioParameterShapes, parametersaggregateshape_)
						parametersaggregateshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagram, "ScenarioParameterShapes", &_diagram.ScenarioParameterShapes)
					}
				} else {
					// ensure parametersaggregateshape_ is NOT in _diagram.ScenarioParameterShapes
					idx := slices.Index(_diagram.ScenarioParameterShapes, parametersaggregateshape_)
					if idx != -1 {
						_diagram.ScenarioParameterShapes = slices.Delete(_diagram.ScenarioParameterShapes, idx, idx+1)
						parametersaggregateshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagram, "ScenarioParameterShapes", &_diagram.ScenarioParameterShapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if parametersaggregateshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		parametersaggregateshape_.Unstage(parametersaggregateshapeFormCallback.probe.stageOfInterest)
	}

	parametersaggregateshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.ParametersAggregateShape](
		parametersaggregateshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if parametersaggregateshapeFormCallback.CreationMode || parametersaggregateshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		parametersaggregateshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(parametersaggregateshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ParametersAggregateShapeFormCallback(
			nil,
			parametersaggregateshapeFormCallback.probe,
			newFormGroup,
		)
		parametersaggregateshape := new(models.ParametersAggregateShape)
		FillUpForm(parametersaggregateshape, newFormGroup, parametersaggregateshapeFormCallback.probe)
		parametersaggregateshapeFormCallback.probe.formStage.Commit()
	}

	parametersaggregateshapeFormCallback.probe.ux_tree()
}
func __gong__New__PositionFormCallback(
	position *models.Position,
	probe *Probe,
	formGroup *form.FormGroup,
) (positionFormCallback *PositionFormCallback) {
	positionFormCallback = new(PositionFormCallback)
	positionFormCallback.probe = probe
	positionFormCallback.position = position
	positionFormCallback.formGroup = formGroup

	positionFormCallback.CreationMode = (position == nil)

	return
}

type PositionFormCallback struct {
	position *models.Position

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (positionFormCallback *PositionFormCallback) OnSave() {
	positionFormCallback.probe.stageOfInterest.Lock()
	defer positionFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("PositionFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	positionFormCallback.probe.formStage.Checkout()

	if positionFormCallback.position == nil {
		positionFormCallback.position = new(models.Position).Stage(positionFormCallback.probe.stageOfInterest)
	}
	position_ := positionFormCallback.position
	_ = position_

	for _, formDiv := range positionFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(position_.Name), formDiv)
		case "Date":
			FormDivBasicFieldToField(&(position_.Date), formDiv)
		case "Ordinate":
			FormDivBasicFieldToField(&(position_.Ordinate), formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(position_.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(position_.IsExpanded), formDiv)
		case "LayoutDirection":
			FormDivEnumIntFieldToField(&(position_.LayoutDirection), formDiv)
		}
	}

	// manage the suppress operation
	if positionFormCallback.formGroup.HasSuppressButtonBeenPressed {
		position_.Unstage(positionFormCallback.probe.stageOfInterest)
	}

	positionFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Position](
		positionFormCallback.probe,
	)

	// display a new form by reset the form stage
	if positionFormCallback.CreationMode || positionFormCallback.formGroup.HasSuppressButtonBeenPressed {
		positionFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(positionFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__PositionFormCallback(
			nil,
			positionFormCallback.probe,
			newFormGroup,
		)
		position := new(models.Position)
		FillUpForm(position, newFormGroup, positionFormCallback.probe)
		positionFormCallback.probe.formStage.Commit()
	}

	positionFormCallback.probe.ux_tree()
}
func __gong__New__RepositoryFormCallback(
	repository *models.Repository,
	probe *Probe,
	formGroup *form.FormGroup,
) (repositoryFormCallback *RepositoryFormCallback) {
	repositoryFormCallback = new(RepositoryFormCallback)
	repositoryFormCallback.probe = probe
	repositoryFormCallback.repository = repository
	repositoryFormCallback.formGroup = formGroup

	repositoryFormCallback.CreationMode = (repository == nil)

	return
}

type RepositoryFormCallback struct {
	repository *models.Repository

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (repositoryFormCallback *RepositoryFormCallback) OnSave() {
	repositoryFormCallback.probe.stageOfInterest.Lock()
	defer repositoryFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("RepositoryFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	repositoryFormCallback.probe.formStage.Checkout()

	if repositoryFormCallback.repository == nil {
		repositoryFormCallback.repository = new(models.Repository).Stage(repositoryFormCallback.probe.stageOfInterest)
	}
	repository_ := repositoryFormCallback.repository
	_ = repository_

	for _, formDiv := range repositoryFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(repository_.Name), formDiv)
		case "ParameterUse":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.ParameterShape](repositoryFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.ParameterShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.ParameterShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					repositoryFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.ParameterShape](repositoryFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			repository_.ParameterUse = instanceSlice
			repositoryFormCallback.probe.UpdateSliceOfPointersCallback(repository_, "ParameterUse", &repository_.ParameterUse)

		case "GroupUse":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.GroupUse](repositoryFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.GroupUse, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.GroupUse)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					repositoryFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.GroupUse](repositoryFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			repository_.GroupUse = instanceSlice
			repositoryFormCallback.probe.UpdateSliceOfPointersCallback(repository_, "GroupUse", &repository_.GroupUse)

		case "ComputedPrefix":
			FormDivBasicFieldToField(&(repository_.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(repository_.IsExpanded), formDiv)
		case "LayoutDirection":
			FormDivEnumIntFieldToField(&(repository_.LayoutDirection), formDiv)
		}
	}

	// manage the suppress operation
	if repositoryFormCallback.formGroup.HasSuppressButtonBeenPressed {
		repository_.Unstage(repositoryFormCallback.probe.stageOfInterest)
	}

	repositoryFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Repository](
		repositoryFormCallback.probe,
	)

	// display a new form by reset the form stage
	if repositoryFormCallback.CreationMode || repositoryFormCallback.formGroup.HasSuppressButtonBeenPressed {
		repositoryFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(repositoryFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__RepositoryFormCallback(
			nil,
			repositoryFormCallback.probe,
			newFormGroup,
		)
		repository := new(models.Repository)
		FillUpForm(repository, newFormGroup, repositoryFormCallback.probe)
		repositoryFormCallback.probe.formStage.Commit()
	}

	repositoryFormCallback.probe.ux_tree()
}
func __gong__New__ScenarioFormCallback(
	scenario *models.Scenario,
	probe *Probe,
	formGroup *form.FormGroup,
) (scenarioFormCallback *ScenarioFormCallback) {
	scenarioFormCallback = new(ScenarioFormCallback)
	scenarioFormCallback.probe = probe
	scenarioFormCallback.scenario = scenario
	scenarioFormCallback.formGroup = formGroup

	scenarioFormCallback.CreationMode = (scenario == nil)

	return
}

type ScenarioFormCallback struct {
	scenario *models.Scenario

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (scenarioFormCallback *ScenarioFormCallback) OnSave() {
	scenarioFormCallback.probe.stageOfInterest.Lock()
	defer scenarioFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("ScenarioFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	scenarioFormCallback.probe.formStage.Checkout()

	if scenarioFormCallback.scenario == nil {
		scenarioFormCallback.scenario = new(models.Scenario).Stage(scenarioFormCallback.probe.stageOfInterest)
	}
	scenario_ := scenarioFormCallback.scenario
	_ = scenario_

	for _, formDiv := range scenarioFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(scenario_.Name), formDiv)
		case "Description":
			FormDivBasicFieldToField(&(scenario_.Description), formDiv)
		case "Diagrams":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Diagram](scenarioFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Diagram, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Diagram)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					scenarioFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Diagram](scenarioFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			scenario_.Diagrams = instanceSlice
			scenarioFormCallback.probe.UpdateSliceOfPointersCallback(scenario_, "Diagrams", &scenario_.Diagrams)

		case "IsDiagramsNodeExpanded":
			FormDivBasicFieldToField(&(scenario_.IsDiagramsNodeExpanded), formDiv)
		case "ActorStates":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.ActorState](scenarioFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.ActorState, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.ActorState)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					scenarioFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.ActorState](scenarioFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			scenario_.ActorStates = instanceSlice
			scenarioFormCallback.probe.UpdateSliceOfPointersCallback(scenario_, "ActorStates", &scenario_.ActorStates)

		case "IsActorStatesNodeExpanded":
			FormDivBasicFieldToField(&(scenario_.IsActorStatesNodeExpanded), formDiv)
		case "ActorStateTransitions":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.ActorStateTransition](scenarioFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.ActorStateTransition, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.ActorStateTransition)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					scenarioFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.ActorStateTransition](scenarioFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			scenario_.ActorStateTransitions = instanceSlice
			scenarioFormCallback.probe.UpdateSliceOfPointersCallback(scenario_, "ActorStateTransitions", &scenario_.ActorStateTransitions)

		case "IsActorStateTransitionsNodeExpanded":
			FormDivBasicFieldToField(&(scenario_.IsActorStateTransitionsNodeExpanded), formDiv)
		case "EvolutionDirections":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.EvolutionDirection](scenarioFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.EvolutionDirection, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.EvolutionDirection)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					scenarioFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.EvolutionDirection](scenarioFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			scenario_.EvolutionDirections = instanceSlice
			scenarioFormCallback.probe.UpdateSliceOfPointersCallback(scenario_, "EvolutionDirections", &scenario_.EvolutionDirections)

		case "IsEvolutionDirectionsNodeExpanded":
			FormDivBasicFieldToField(&(scenario_.IsEvolutionDirectionsNodeExpanded), formDiv)
		case "Parameters":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Parameter](scenarioFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Parameter, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Parameter)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					scenarioFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Parameter](scenarioFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			scenario_.Parameters = instanceSlice
			scenarioFormCallback.probe.UpdateSliceOfPointersCallback(scenario_, "Parameters", &scenario_.Parameters)

		case "IsParametersNodeExpanded":
			FormDivBasicFieldToField(&(scenario_.IsParametersNodeExpanded), formDiv)
		case "ParametersAggretates":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.ParametersAggregate](scenarioFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.ParametersAggregate, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.ParametersAggregate)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					scenarioFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.ParametersAggregate](scenarioFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			scenario_.ParametersAggretates = instanceSlice
			scenarioFormCallback.probe.UpdateSliceOfPointersCallback(scenario_, "ParametersAggretates", &scenario_.ParametersAggretates)

		case "IsParametersAggretatesNodeExpanded":
			FormDivBasicFieldToField(&(scenario_.IsParametersAggretatesNodeExpanded), formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(scenario_.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(scenario_.IsExpanded), formDiv)
		case "LayoutDirection":
			FormDivEnumIntFieldToField(&(scenario_.LayoutDirection), formDiv)
		case "Analysis:Scenarios":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Analysis instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Analysis instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Analysis](scenarioFormCallback.probe.stageOfInterest)
			targetAnalysisIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetAnalysisIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Analysis instances and update their Scenarios slice
			for _analysis := range *models.GetGongstructInstancesSetFromPointerType[*models.Analysis](scenarioFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(scenarioFormCallback.probe.stageOfInterest, _analysis)
				
				// if Analysis is selected
				if targetAnalysisIDs[id] {
					// ensure scenario_ is in _analysis.Scenarios
					found := false
					for _, _b := range _analysis.Scenarios {
						if _b == scenario_ {
							found = true
							break
						}
					}
					if !found {
						_analysis.Scenarios = append(_analysis.Scenarios, scenario_)
						scenarioFormCallback.probe.UpdateSliceOfPointersCallback(_analysis, "Scenarios", &_analysis.Scenarios)
					}
				} else {
					// ensure scenario_ is NOT in _analysis.Scenarios
					idx := slices.Index(_analysis.Scenarios, scenario_)
					if idx != -1 {
						_analysis.Scenarios = slices.Delete(_analysis.Scenarios, idx, idx+1)
						scenarioFormCallback.probe.UpdateSliceOfPointersCallback(_analysis, "Scenarios", &_analysis.Scenarios)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if scenarioFormCallback.formGroup.HasSuppressButtonBeenPressed {
		scenario_.Unstage(scenarioFormCallback.probe.stageOfInterest)
	}

	scenarioFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Scenario](
		scenarioFormCallback.probe,
	)

	// display a new form by reset the form stage
	if scenarioFormCallback.CreationMode || scenarioFormCallback.formGroup.HasSuppressButtonBeenPressed {
		scenarioFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(scenarioFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ScenarioFormCallback(
			nil,
			scenarioFormCallback.probe,
			newFormGroup,
		)
		scenario := new(models.Scenario)
		FillUpForm(scenario, newFormGroup, scenarioFormCallback.probe)
		scenarioFormCallback.probe.formStage.Commit()
	}

	scenarioFormCallback.probe.ux_tree()
}
func __gong__New__UserFormCallback(
	user *models.User,
	probe *Probe,
	formGroup *form.FormGroup,
) (userFormCallback *UserFormCallback) {
	userFormCallback = new(UserFormCallback)
	userFormCallback.probe = probe
	userFormCallback.user = user
	userFormCallback.formGroup = formGroup

	userFormCallback.CreationMode = (user == nil)

	return
}

type UserFormCallback struct {
	user *models.User

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (userFormCallback *UserFormCallback) OnSave() {
	userFormCallback.probe.stageOfInterest.Lock()
	defer userFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("UserFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	userFormCallback.probe.formStage.Checkout()

	if userFormCallback.user == nil {
		userFormCallback.user = new(models.User).Stage(userFormCallback.probe.stageOfInterest)
	}
	user_ := userFormCallback.user
	_ = user_

	for _, formDiv := range userFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(user_.Name), formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(user_.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(user_.IsExpanded), formDiv)
		case "LayoutDirection":
			FormDivEnumIntFieldToField(&(user_.LayoutDirection), formDiv)
		}
	}

	// manage the suppress operation
	if userFormCallback.formGroup.HasSuppressButtonBeenPressed {
		user_.Unstage(userFormCallback.probe.stageOfInterest)
	}

	userFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.User](
		userFormCallback.probe,
	)

	// display a new form by reset the form stage
	if userFormCallback.CreationMode || userFormCallback.formGroup.HasSuppressButtonBeenPressed {
		userFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(userFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__UserFormCallback(
			nil,
			userFormCallback.probe,
			newFormGroup,
		)
		user := new(models.User)
		FillUpForm(user, newFormGroup, userFormCallback.probe)
		userFormCallback.probe.formStage.Commit()
	}

	userFormCallback.probe.ux_tree()
}
func __gong__New__UserUseFormCallback(
	useruse *models.UserUse,
	probe *Probe,
	formGroup *form.FormGroup,
) (useruseFormCallback *UserUseFormCallback) {
	useruseFormCallback = new(UserUseFormCallback)
	useruseFormCallback.probe = probe
	useruseFormCallback.useruse = useruse
	useruseFormCallback.formGroup = formGroup

	useruseFormCallback.CreationMode = (useruse == nil)

	return
}

type UserUseFormCallback struct {
	useruse *models.UserUse

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (useruseFormCallback *UserUseFormCallback) OnSave() {
	useruseFormCallback.probe.stageOfInterest.Lock()
	defer useruseFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("UserUseFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	useruseFormCallback.probe.formStage.Checkout()

	if useruseFormCallback.useruse == nil {
		useruseFormCallback.useruse = new(models.UserUse).Stage(useruseFormCallback.probe.stageOfInterest)
	}
	useruse_ := useruseFormCallback.useruse
	_ = useruse_

	for _, formDiv := range useruseFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(useruse_.Name), formDiv)
		case "User":
			FormDivSelectFieldToField(&(useruse_.User), useruseFormCallback.probe.stageOfInterest, formDiv)
		case "Group:UserUse":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Group instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Group instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Group](useruseFormCallback.probe.stageOfInterest)
			targetGroupIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetGroupIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Group instances and update their UserUse slice
			for _group := range *models.GetGongstructInstancesSetFromPointerType[*models.Group](useruseFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(useruseFormCallback.probe.stageOfInterest, _group)
				
				// if Group is selected
				if targetGroupIDs[id] {
					// ensure useruse_ is in _group.UserUse
					found := false
					for _, _b := range _group.UserUse {
						if _b == useruse_ {
							found = true
							break
						}
					}
					if !found {
						_group.UserUse = append(_group.UserUse, useruse_)
						useruseFormCallback.probe.UpdateSliceOfPointersCallback(_group, "UserUse", &_group.UserUse)
					}
				} else {
					// ensure useruse_ is NOT in _group.UserUse
					idx := slices.Index(_group.UserUse, useruse_)
					if idx != -1 {
						_group.UserUse = slices.Delete(_group.UserUse, idx, idx+1)
						useruseFormCallback.probe.UpdateSliceOfPointersCallback(_group, "UserUse", &_group.UserUse)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if useruseFormCallback.formGroup.HasSuppressButtonBeenPressed {
		useruse_.Unstage(useruseFormCallback.probe.stageOfInterest)
	}

	useruseFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.UserUse](
		useruseFormCallback.probe,
	)

	// display a new form by reset the form stage
	if useruseFormCallback.CreationMode || useruseFormCallback.formGroup.HasSuppressButtonBeenPressed {
		useruseFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(useruseFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__UserUseFormCallback(
			nil,
			useruseFormCallback.probe,
			newFormGroup,
		)
		useruse := new(models.UserUse)
		FillUpForm(useruse, newFormGroup, useruseFormCallback.probe)
		useruseFormCallback.probe.formStage.Commit()
	}

	useruseFormCallback.probe.ux_tree()
}
func __gong__New__WorkspaceFormCallback(
	workspace *models.Workspace,
	probe *Probe,
	formGroup *form.FormGroup,
) (workspaceFormCallback *WorkspaceFormCallback) {
	workspaceFormCallback = new(WorkspaceFormCallback)
	workspaceFormCallback.probe = probe
	workspaceFormCallback.workspace = workspace
	workspaceFormCallback.formGroup = formGroup

	workspaceFormCallback.CreationMode = (workspace == nil)

	return
}

type WorkspaceFormCallback struct {
	workspace *models.Workspace

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (workspaceFormCallback *WorkspaceFormCallback) OnSave() {
	workspaceFormCallback.probe.stageOfInterest.Lock()
	defer workspaceFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("WorkspaceFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	workspaceFormCallback.probe.formStage.Checkout()

	if workspaceFormCallback.workspace == nil {
		workspaceFormCallback.workspace = new(models.Workspace).Stage(workspaceFormCallback.probe.stageOfInterest)
	}
	workspace_ := workspaceFormCallback.workspace
	_ = workspace_

	for _, formDiv := range workspaceFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(workspace_.Name), formDiv)
		case "SelectedDiagram":
			FormDivSelectFieldToField(&(workspace_.SelectedDiagram), workspaceFormCallback.probe.stageOfInterest, formDiv)
		case "Default_EvolutionDirectionShape":
			FormDivSelectFieldToField(&(workspace_.Default_EvolutionDirectionShape), workspaceFormCallback.probe.stageOfInterest, formDiv)
		case "Default_ParameterShape":
			FormDivSelectFieldToField(&(workspace_.Default_ParameterShape), workspaceFormCallback.probe.stageOfInterest, formDiv)
		case "Default_ScenarioParameterShape":
			FormDivSelectFieldToField(&(workspace_.Default_ScenarioParameterShape), workspaceFormCallback.probe.stageOfInterest, formDiv)
		case "Default_ActorStateShape":
			FormDivSelectFieldToField(&(workspace_.Default_ActorStateShape), workspaceFormCallback.probe.stageOfInterest, formDiv)
		case "Default_ActorStateTransitionShape":
			FormDivSelectFieldToField(&(workspace_.Default_ActorStateTransitionShape), workspaceFormCallback.probe.stageOfInterest, formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(workspace_.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(workspace_.IsExpanded), formDiv)
		case "LayoutDirection":
			FormDivEnumIntFieldToField(&(workspace_.LayoutDirection), formDiv)
		}
	}

	// manage the suppress operation
	if workspaceFormCallback.formGroup.HasSuppressButtonBeenPressed {
		workspace_.Unstage(workspaceFormCallback.probe.stageOfInterest)
	}

	workspaceFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Workspace](
		workspaceFormCallback.probe,
	)

	// display a new form by reset the form stage
	if workspaceFormCallback.CreationMode || workspaceFormCallback.formGroup.HasSuppressButtonBeenPressed {
		workspaceFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(workspaceFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__WorkspaceFormCallback(
			nil,
			workspaceFormCallback.probe,
			newFormGroup,
		)
		workspace := new(models.Workspace)
		FillUpForm(workspace, newFormGroup, workspaceFormCallback.probe)
		workspaceFormCallback.probe.formStage.Commit()
	}

	workspaceFormCallback.probe.ux_tree()
}
