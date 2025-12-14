// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	table "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gong/test/statemachines/go/models"
)

// to avoid errors when time and slices packages are not used in the generated code
const _ = time.Nanosecond

var _ = slices.Delete([]string{"a"}, 0, 1)

var _ = log.Panicf

// insertion point
func __gong__New__ActionFormCallback(
	action *models.Action,
	probe *Probe,
	formGroup *table.FormGroup,
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

	formGroup *table.FormGroup
}

func (actionFormCallback *ActionFormCallback) OnSave() {

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
		newFormGroup := (&table.FormGroup{
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

	updateAndCommitTree(actionFormCallback.probe)
}
func __gong__New__ActivitiesFormCallback(
	activities *models.Activities,
	probe *Probe,
	formGroup *table.FormGroup,
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

	formGroup *table.FormGroup
}

func (activitiesFormCallback *ActivitiesFormCallback) OnSave() {

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
			// WARNING : this form deals with the N-N association "State.Activities []*Activities" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Activities). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.State
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "State"
				rf.Fieldname = "Activities"
				formerAssociationSource := activities_.GongGetReverseFieldOwner(
					activitiesFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.State)
					if !ok {
						log.Fatalln("Source of State.Activities []*Activities, is not an State instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.Activities, activities_)
					formerSource.Activities = slices.Delete(formerSource.Activities, idx, idx+1)
				}
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.State
			for _state := range *models.GetGongstructInstancesSet[models.State](activitiesFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _state.GetName() == newSourceName.GetName() {
					newSource = _state // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of State.Activities []*Activities, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.Activities = append(newSource.Activities, activities_)
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
		newFormGroup := (&table.FormGroup{
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

	updateAndCommitTree(activitiesFormCallback.probe)
}
func __gong__New__ArchitectureFormCallback(
	architecture *models.Architecture,
	probe *Probe,
	formGroup *table.FormGroup,
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

	formGroup *table.FormGroup
}

func (architectureFormCallback *ArchitectureFormCallback) OnSave() {

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

			ids, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			for _, id := range ids {
				instanceSlice = append(instanceSlice, map_id_instances[id])
			}
			architecture_.StateMachines = instanceSlice

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

			ids, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			for _, id := range ids {
				instanceSlice = append(instanceSlice, map_id_instances[id])
			}
			architecture_.Roles = instanceSlice

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
		newFormGroup := (&table.FormGroup{
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

	updateAndCommitTree(architectureFormCallback.probe)
}
func __gong__New__DiagramFormCallback(
	diagram *models.Diagram,
	probe *Probe,
	formGroup *table.FormGroup,
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

	formGroup *table.FormGroup
}

func (diagramFormCallback *DiagramFormCallback) OnSave() {

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
		case "IsInRenameMode":
			FormDivBasicFieldToField(&(diagram_.IsInRenameMode), formDiv)
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

			ids, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			for _, id := range ids {
				instanceSlice = append(instanceSlice, map_id_instances[id])
			}
			diagram_.State_Shapes = instanceSlice

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

			ids, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			for _, id := range ids {
				instanceSlice = append(instanceSlice, map_id_instances[id])
			}
			diagram_.Transition_Shapes = instanceSlice

		case "State:Diagrams":
			// WARNING : this form deals with the N-N association "State.Diagrams []*Diagram" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Diagram). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.State
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "State"
				rf.Fieldname = "Diagrams"
				formerAssociationSource := diagram_.GongGetReverseFieldOwner(
					diagramFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.State)
					if !ok {
						log.Fatalln("Source of State.Diagrams []*Diagram, is not an State instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.Diagrams, diagram_)
					formerSource.Diagrams = slices.Delete(formerSource.Diagrams, idx, idx+1)
				}
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.State
			for _state := range *models.GetGongstructInstancesSet[models.State](diagramFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _state.GetName() == newSourceName.GetName() {
					newSource = _state // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of State.Diagrams []*Diagram, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.Diagrams = append(newSource.Diagrams, diagram_)
		case "StateMachine:Diagrams":
			// WARNING : this form deals with the N-N association "StateMachine.Diagrams []*Diagram" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Diagram). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.StateMachine
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "StateMachine"
				rf.Fieldname = "Diagrams"
				formerAssociationSource := diagram_.GongGetReverseFieldOwner(
					diagramFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.StateMachine)
					if !ok {
						log.Fatalln("Source of StateMachine.Diagrams []*Diagram, is not an StateMachine instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.Diagrams, diagram_)
					formerSource.Diagrams = slices.Delete(formerSource.Diagrams, idx, idx+1)
				}
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.StateMachine
			for _statemachine := range *models.GetGongstructInstancesSet[models.StateMachine](diagramFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _statemachine.GetName() == newSourceName.GetName() {
					newSource = _statemachine // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of StateMachine.Diagrams []*Diagram, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.Diagrams = append(newSource.Diagrams, diagram_)
		case "Transition:Diagrams":
			// WARNING : this form deals with the N-N association "Transition.Diagrams []*Diagram" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Diagram). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Transition
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Transition"
				rf.Fieldname = "Diagrams"
				formerAssociationSource := diagram_.GongGetReverseFieldOwner(
					diagramFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Transition)
					if !ok {
						log.Fatalln("Source of Transition.Diagrams []*Diagram, is not an Transition instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.Diagrams, diagram_)
					formerSource.Diagrams = slices.Delete(formerSource.Diagrams, idx, idx+1)
				}
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.Transition
			for _transition := range *models.GetGongstructInstancesSet[models.Transition](diagramFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _transition.GetName() == newSourceName.GetName() {
					newSource = _transition // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Transition.Diagrams []*Diagram, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.Diagrams = append(newSource.Diagrams, diagram_)
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
		newFormGroup := (&table.FormGroup{
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

	updateAndCommitTree(diagramFormCallback.probe)
}
func __gong__New__KillFormCallback(
	kill *models.Kill,
	probe *Probe,
	formGroup *table.FormGroup,
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

	formGroup *table.FormGroup
}

func (killFormCallback *KillFormCallback) OnSave() {

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
		newFormGroup := (&table.FormGroup{
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

	updateAndCommitTree(killFormCallback.probe)
}
func __gong__New__MessageFormCallback(
	message *models.Message,
	probe *Probe,
	formGroup *table.FormGroup,
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

	formGroup *table.FormGroup
}

func (messageFormCallback *MessageFormCallback) OnSave() {

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
			// WARNING : this form deals with the N-N association "Object.Messages []*Message" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Message). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Object
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Object"
				rf.Fieldname = "Messages"
				formerAssociationSource := message_.GongGetReverseFieldOwner(
					messageFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Object)
					if !ok {
						log.Fatalln("Source of Object.Messages []*Message, is not an Object instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.Messages, message_)
					formerSource.Messages = slices.Delete(formerSource.Messages, idx, idx+1)
				}
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.Object
			for _object := range *models.GetGongstructInstancesSet[models.Object](messageFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _object.GetName() == newSourceName.GetName() {
					newSource = _object // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Object.Messages []*Message, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.Messages = append(newSource.Messages, message_)
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
		newFormGroup := (&table.FormGroup{
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

	updateAndCommitTree(messageFormCallback.probe)
}
func __gong__New__MessageTypeFormCallback(
	messagetype *models.MessageType,
	probe *Probe,
	formGroup *table.FormGroup,
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

	formGroup *table.FormGroup
}

func (messagetypeFormCallback *MessageTypeFormCallback) OnSave() {

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
			// WARNING : this form deals with the N-N association "Transition.GeneratedMessages []*MessageType" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of MessageType). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Transition
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Transition"
				rf.Fieldname = "GeneratedMessages"
				formerAssociationSource := messagetype_.GongGetReverseFieldOwner(
					messagetypeFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Transition)
					if !ok {
						log.Fatalln("Source of Transition.GeneratedMessages []*MessageType, is not an Transition instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.GeneratedMessages, messagetype_)
					formerSource.GeneratedMessages = slices.Delete(formerSource.GeneratedMessages, idx, idx+1)
				}
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.Transition
			for _transition := range *models.GetGongstructInstancesSet[models.Transition](messagetypeFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _transition.GetName() == newSourceName.GetName() {
					newSource = _transition // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Transition.GeneratedMessages []*MessageType, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.GeneratedMessages = append(newSource.GeneratedMessages, messagetype_)
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
		newFormGroup := (&table.FormGroup{
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

	updateAndCommitTree(messagetypeFormCallback.probe)
}
func __gong__New__ObjectFormCallback(
	object *models.Object,
	probe *Probe,
	formGroup *table.FormGroup,
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

	formGroup *table.FormGroup
}

func (objectFormCallback *ObjectFormCallback) OnSave() {

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

			ids, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			for _, id := range ids {
				instanceSlice = append(instanceSlice, map_id_instances[id])
			}
			object_.Messages = instanceSlice

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
		newFormGroup := (&table.FormGroup{
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

	updateAndCommitTree(objectFormCallback.probe)
}
func __gong__New__RoleFormCallback(
	role *models.Role,
	probe *Probe,
	formGroup *table.FormGroup,
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

	formGroup *table.FormGroup
}

func (roleFormCallback *RoleFormCallback) OnSave() {

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

			ids, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			for _, id := range ids {
				instanceSlice = append(instanceSlice, map_id_instances[id])
			}
			role_.RolesWithSamePermissions = instanceSlice

		case "Architecture:Roles":
			// WARNING : this form deals with the N-N association "Architecture.Roles []*Role" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Role). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Architecture
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Architecture"
				rf.Fieldname = "Roles"
				formerAssociationSource := role_.GongGetReverseFieldOwner(
					roleFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Architecture)
					if !ok {
						log.Fatalln("Source of Architecture.Roles []*Role, is not an Architecture instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.Roles, role_)
					formerSource.Roles = slices.Delete(formerSource.Roles, idx, idx+1)
				}
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.Architecture
			for _architecture := range *models.GetGongstructInstancesSet[models.Architecture](roleFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _architecture.GetName() == newSourceName.GetName() {
					newSource = _architecture // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Architecture.Roles []*Role, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.Roles = append(newSource.Roles, role_)
		case "Role:RolesWithSamePermissions":
			// WARNING : this form deals with the N-N association "Role.RolesWithSamePermissions []*Role" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Role). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Role
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Role"
				rf.Fieldname = "RolesWithSamePermissions"
				formerAssociationSource := role_.GongGetReverseFieldOwner(
					roleFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Role)
					if !ok {
						log.Fatalln("Source of Role.RolesWithSamePermissions []*Role, is not an Role instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.RolesWithSamePermissions, role_)
					formerSource.RolesWithSamePermissions = slices.Delete(formerSource.RolesWithSamePermissions, idx, idx+1)
				}
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.Role
			for _role := range *models.GetGongstructInstancesSet[models.Role](roleFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _role.GetName() == newSourceName.GetName() {
					newSource = _role // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Role.RolesWithSamePermissions []*Role, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.RolesWithSamePermissions = append(newSource.RolesWithSamePermissions, role_)
		case "Transition:RolesWithPermissions":
			// WARNING : this form deals with the N-N association "Transition.RolesWithPermissions []*Role" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Role). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Transition
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Transition"
				rf.Fieldname = "RolesWithPermissions"
				formerAssociationSource := role_.GongGetReverseFieldOwner(
					roleFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Transition)
					if !ok {
						log.Fatalln("Source of Transition.RolesWithPermissions []*Role, is not an Transition instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.RolesWithPermissions, role_)
					formerSource.RolesWithPermissions = slices.Delete(formerSource.RolesWithPermissions, idx, idx+1)
				}
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.Transition
			for _transition := range *models.GetGongstructInstancesSet[models.Transition](roleFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _transition.GetName() == newSourceName.GetName() {
					newSource = _transition // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Transition.RolesWithPermissions []*Role, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.RolesWithPermissions = append(newSource.RolesWithPermissions, role_)
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
		newFormGroup := (&table.FormGroup{
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

	updateAndCommitTree(roleFormCallback.probe)
}
func __gong__New__StateFormCallback(
	state *models.State,
	probe *Probe,
	formGroup *table.FormGroup,
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

	formGroup *table.FormGroup
}

func (stateFormCallback *StateFormCallback) OnSave() {

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
		case "IsFictif":
			FormDivBasicFieldToField(&(state_.IsFictif), formDiv)
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

			ids, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			for _, id := range ids {
				instanceSlice = append(instanceSlice, map_id_instances[id])
			}
			state_.SubStates = instanceSlice

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

			ids, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			for _, id := range ids {
				instanceSlice = append(instanceSlice, map_id_instances[id])
			}
			state_.Diagrams = instanceSlice

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

			ids, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			for _, id := range ids {
				instanceSlice = append(instanceSlice, map_id_instances[id])
			}
			state_.Activities = instanceSlice

		case "Exit":
			FormDivSelectFieldToField(&(state_.Exit), stateFormCallback.probe.stageOfInterest, formDiv)
		case "State:SubStates":
			// WARNING : this form deals with the N-N association "State.SubStates []*State" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of State). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.State
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "State"
				rf.Fieldname = "SubStates"
				formerAssociationSource := state_.GongGetReverseFieldOwner(
					stateFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.State)
					if !ok {
						log.Fatalln("Source of State.SubStates []*State, is not an State instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.SubStates, state_)
					formerSource.SubStates = slices.Delete(formerSource.SubStates, idx, idx+1)
				}
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.State
			for _state := range *models.GetGongstructInstancesSet[models.State](stateFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _state.GetName() == newSourceName.GetName() {
					newSource = _state // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of State.SubStates []*State, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.SubStates = append(newSource.SubStates, state_)
		case "StateMachine:States":
			// WARNING : this form deals with the N-N association "StateMachine.States []*State" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of State). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.StateMachine
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "StateMachine"
				rf.Fieldname = "States"
				formerAssociationSource := state_.GongGetReverseFieldOwner(
					stateFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.StateMachine)
					if !ok {
						log.Fatalln("Source of StateMachine.States []*State, is not an StateMachine instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.States, state_)
					formerSource.States = slices.Delete(formerSource.States, idx, idx+1)
				}
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.StateMachine
			for _statemachine := range *models.GetGongstructInstancesSet[models.StateMachine](stateFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _statemachine.GetName() == newSourceName.GetName() {
					newSource = _statemachine // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of StateMachine.States []*State, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.States = append(newSource.States, state_)
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
		newFormGroup := (&table.FormGroup{
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

	updateAndCommitTree(stateFormCallback.probe)
}
func __gong__New__StateMachineFormCallback(
	statemachine *models.StateMachine,
	probe *Probe,
	formGroup *table.FormGroup,
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

	formGroup *table.FormGroup
}

func (statemachineFormCallback *StateMachineFormCallback) OnSave() {

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
		case "IsNodeExpanded":
			FormDivBasicFieldToField(&(statemachine_.IsNodeExpanded), formDiv)
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

			ids, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			for _, id := range ids {
				instanceSlice = append(instanceSlice, map_id_instances[id])
			}
			statemachine_.States = instanceSlice

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

			ids, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			for _, id := range ids {
				instanceSlice = append(instanceSlice, map_id_instances[id])
			}
			statemachine_.Diagrams = instanceSlice

		case "InitialState":
			FormDivSelectFieldToField(&(statemachine_.InitialState), statemachineFormCallback.probe.stageOfInterest, formDiv)
		case "Architecture:StateMachines":
			// WARNING : this form deals with the N-N association "Architecture.StateMachines []*StateMachine" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of StateMachine). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Architecture
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Architecture"
				rf.Fieldname = "StateMachines"
				formerAssociationSource := statemachine_.GongGetReverseFieldOwner(
					statemachineFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Architecture)
					if !ok {
						log.Fatalln("Source of Architecture.StateMachines []*StateMachine, is not an Architecture instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.StateMachines, statemachine_)
					formerSource.StateMachines = slices.Delete(formerSource.StateMachines, idx, idx+1)
				}
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.Architecture
			for _architecture := range *models.GetGongstructInstancesSet[models.Architecture](statemachineFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _architecture.GetName() == newSourceName.GetName() {
					newSource = _architecture // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Architecture.StateMachines []*StateMachine, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.StateMachines = append(newSource.StateMachines, statemachine_)
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
		newFormGroup := (&table.FormGroup{
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

	updateAndCommitTree(statemachineFormCallback.probe)
}
func __gong__New__StateShapeFormCallback(
	stateshape *models.StateShape,
	probe *Probe,
	formGroup *table.FormGroup,
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

	formGroup *table.FormGroup
}

func (stateshapeFormCallback *StateShapeFormCallback) OnSave() {

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
		case "IsExpanded":
			FormDivBasicFieldToField(&(stateshape_.IsExpanded), formDiv)
		case "X":
			FormDivBasicFieldToField(&(stateshape_.X), formDiv)
		case "Y":
			FormDivBasicFieldToField(&(stateshape_.Y), formDiv)
		case "Width":
			FormDivBasicFieldToField(&(stateshape_.Width), formDiv)
		case "Height":
			FormDivBasicFieldToField(&(stateshape_.Height), formDiv)
		case "Diagram:State_Shapes":
			// WARNING : this form deals with the N-N association "Diagram.State_Shapes []*StateShape" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of StateShape). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Diagram
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Diagram"
				rf.Fieldname = "State_Shapes"
				formerAssociationSource := stateshape_.GongGetReverseFieldOwner(
					stateshapeFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Diagram)
					if !ok {
						log.Fatalln("Source of Diagram.State_Shapes []*StateShape, is not an Diagram instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.State_Shapes, stateshape_)
					formerSource.State_Shapes = slices.Delete(formerSource.State_Shapes, idx, idx+1)
				}
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.Diagram
			for _diagram := range *models.GetGongstructInstancesSet[models.Diagram](stateshapeFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _diagram.GetName() == newSourceName.GetName() {
					newSource = _diagram // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Diagram.State_Shapes []*StateShape, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.State_Shapes = append(newSource.State_Shapes, stateshape_)
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
		newFormGroup := (&table.FormGroup{
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

	updateAndCommitTree(stateshapeFormCallback.probe)
}
func __gong__New__TransitionFormCallback(
	transition *models.Transition,
	probe *Probe,
	formGroup *table.FormGroup,
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

	formGroup *table.FormGroup
}

func (transitionFormCallback *TransitionFormCallback) OnSave() {

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

			ids, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			for _, id := range ids {
				instanceSlice = append(instanceSlice, map_id_instances[id])
			}
			transition_.RolesWithPermissions = instanceSlice

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

			ids, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			for _, id := range ids {
				instanceSlice = append(instanceSlice, map_id_instances[id])
			}
			transition_.GeneratedMessages = instanceSlice

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

			ids, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			for _, id := range ids {
				instanceSlice = append(instanceSlice, map_id_instances[id])
			}
			transition_.Diagrams = instanceSlice

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
		newFormGroup := (&table.FormGroup{
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

	updateAndCommitTree(transitionFormCallback.probe)
}
func __gong__New__Transition_ShapeFormCallback(
	transition_shape *models.Transition_Shape,
	probe *Probe,
	formGroup *table.FormGroup,
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

	formGroup *table.FormGroup
}

func (transition_shapeFormCallback *Transition_ShapeFormCallback) OnSave() {

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
		case "Diagram:Transition_Shapes":
			// WARNING : this form deals with the N-N association "Diagram.Transition_Shapes []*Transition_Shape" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Transition_Shape). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Diagram
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Diagram"
				rf.Fieldname = "Transition_Shapes"
				formerAssociationSource := transition_shape_.GongGetReverseFieldOwner(
					transition_shapeFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Diagram)
					if !ok {
						log.Fatalln("Source of Diagram.Transition_Shapes []*Transition_Shape, is not an Diagram instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.Transition_Shapes, transition_shape_)
					formerSource.Transition_Shapes = slices.Delete(formerSource.Transition_Shapes, idx, idx+1)
				}
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.Diagram
			for _diagram := range *models.GetGongstructInstancesSet[models.Diagram](transition_shapeFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _diagram.GetName() == newSourceName.GetName() {
					newSource = _diagram // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Diagram.Transition_Shapes []*Transition_Shape, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.Transition_Shapes = append(newSource.Transition_Shapes, transition_shape_)
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
		newFormGroup := (&table.FormGroup{
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

	updateAndCommitTree(transition_shapeFormCallback.probe)
}
