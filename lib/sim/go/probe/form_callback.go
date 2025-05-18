// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	table "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gong/lib/sim/go/models"
)

// to avoid errors when time and slices packages are not used in the generated code
const _ = time.Nanosecond

var _ = slices.Delete([]string{"a"}, 0, 1)

var _ = log.Panicf

// insertion point
func __gong__New__CommandFormCallback(
	command *models.Command,
	probe *Probe,
	formGroup *table.FormGroup,
) (commandFormCallback *CommandFormCallback) {
	commandFormCallback = new(CommandFormCallback)
	commandFormCallback.probe = probe
	commandFormCallback.command = command
	commandFormCallback.formGroup = formGroup

	commandFormCallback.CreationMode = (command == nil)

	return
}

type CommandFormCallback struct {
	command *models.Command

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (commandFormCallback *CommandFormCallback) OnSave() {

	// log.Println("CommandFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	commandFormCallback.probe.formStage.Checkout()

	if commandFormCallback.command == nil {
		commandFormCallback.command = new(models.Command).Stage(commandFormCallback.probe.stageOfInterest)
	}
	command_ := commandFormCallback.command
	_ = command_

	for _, formDiv := range commandFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(command_.Name), formDiv)
		case "Command":
			FormDivEnumStringFieldToField(&(command_.Command), formDiv)
		case "CommandDate":
			FormDivBasicFieldToField(&(command_.CommandDate), formDiv)
		case "Engine":
			FormDivSelectFieldToField(&(command_.Engine), commandFormCallback.probe.stageOfInterest, formDiv)
		}
	}

	// manage the suppress operation
	if commandFormCallback.formGroup.HasSuppressButtonBeenPressed {
		command_.Unstage(commandFormCallback.probe.stageOfInterest)
	}

	commandFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.Command](
		commandFormCallback.probe,
	)
	commandFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if commandFormCallback.CreationMode || commandFormCallback.formGroup.HasSuppressButtonBeenPressed {
		commandFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(commandFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__CommandFormCallback(
			nil,
			commandFormCallback.probe,
			newFormGroup,
		)
		command := new(models.Command)
		FillUpForm(command, newFormGroup, commandFormCallback.probe)
		commandFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(commandFormCallback.probe)
}
func __gong__New__DummyAgentFormCallback(
	dummyagent *models.DummyAgent,
	probe *Probe,
	formGroup *table.FormGroup,
) (dummyagentFormCallback *DummyAgentFormCallback) {
	dummyagentFormCallback = new(DummyAgentFormCallback)
	dummyagentFormCallback.probe = probe
	dummyagentFormCallback.dummyagent = dummyagent
	dummyagentFormCallback.formGroup = formGroup

	dummyagentFormCallback.CreationMode = (dummyagent == nil)

	return
}

type DummyAgentFormCallback struct {
	dummyagent *models.DummyAgent

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (dummyagentFormCallback *DummyAgentFormCallback) OnSave() {

	// log.Println("DummyAgentFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	dummyagentFormCallback.probe.formStage.Checkout()

	if dummyagentFormCallback.dummyagent == nil {
		dummyagentFormCallback.dummyagent = new(models.DummyAgent).Stage(dummyagentFormCallback.probe.stageOfInterest)
	}
	dummyagent_ := dummyagentFormCallback.dummyagent
	_ = dummyagent_

	for _, formDiv := range dummyagentFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "TechName":
			FormDivBasicFieldToField(&(dummyagent_.TechName), formDiv)
		case "Name":
			FormDivBasicFieldToField(&(dummyagent_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if dummyagentFormCallback.formGroup.HasSuppressButtonBeenPressed {
		dummyagent_.Unstage(dummyagentFormCallback.probe.stageOfInterest)
	}

	dummyagentFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.DummyAgent](
		dummyagentFormCallback.probe,
	)
	dummyagentFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if dummyagentFormCallback.CreationMode || dummyagentFormCallback.formGroup.HasSuppressButtonBeenPressed {
		dummyagentFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(dummyagentFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__DummyAgentFormCallback(
			nil,
			dummyagentFormCallback.probe,
			newFormGroup,
		)
		dummyagent := new(models.DummyAgent)
		FillUpForm(dummyagent, newFormGroup, dummyagentFormCallback.probe)
		dummyagentFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(dummyagentFormCallback.probe)
}
func __gong__New__EngineFormCallback(
	engine *models.Engine,
	probe *Probe,
	formGroup *table.FormGroup,
) (engineFormCallback *EngineFormCallback) {
	engineFormCallback = new(EngineFormCallback)
	engineFormCallback.probe = probe
	engineFormCallback.engine = engine
	engineFormCallback.formGroup = formGroup

	engineFormCallback.CreationMode = (engine == nil)

	return
}

type EngineFormCallback struct {
	engine *models.Engine

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (engineFormCallback *EngineFormCallback) OnSave() {

	// log.Println("EngineFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	engineFormCallback.probe.formStage.Checkout()

	if engineFormCallback.engine == nil {
		engineFormCallback.engine = new(models.Engine).Stage(engineFormCallback.probe.stageOfInterest)
	}
	engine_ := engineFormCallback.engine
	_ = engine_

	for _, formDiv := range engineFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(engine_.Name), formDiv)
		case "EndTime":
			FormDivBasicFieldToField(&(engine_.EndTime), formDiv)
		case "CurrentTime":
			FormDivBasicFieldToField(&(engine_.CurrentTime), formDiv)
		case "DisplayFormat":
			FormDivBasicFieldToField(&(engine_.DisplayFormat), formDiv)
		case "SecondsSinceStart":
			FormDivBasicFieldToField(&(engine_.SecondsSinceStart), formDiv)
		case "Fired":
			FormDivBasicFieldToField(&(engine_.Fired), formDiv)
		case "ControlMode":
			FormDivEnumStringFieldToField(&(engine_.ControlMode), formDiv)
		case "State":
			FormDivEnumStringFieldToField(&(engine_.State), formDiv)
		case "Speed":
			FormDivBasicFieldToField(&(engine_.Speed), formDiv)
		}
	}

	// manage the suppress operation
	if engineFormCallback.formGroup.HasSuppressButtonBeenPressed {
		engine_.Unstage(engineFormCallback.probe.stageOfInterest)
	}

	engineFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.Engine](
		engineFormCallback.probe,
	)
	engineFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if engineFormCallback.CreationMode || engineFormCallback.formGroup.HasSuppressButtonBeenPressed {
		engineFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(engineFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__EngineFormCallback(
			nil,
			engineFormCallback.probe,
			newFormGroup,
		)
		engine := new(models.Engine)
		FillUpForm(engine, newFormGroup, engineFormCallback.probe)
		engineFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(engineFormCallback.probe)
}
func __gong__New__EventFormCallback(
	event *models.Event,
	probe *Probe,
	formGroup *table.FormGroup,
) (eventFormCallback *EventFormCallback) {
	eventFormCallback = new(EventFormCallback)
	eventFormCallback.probe = probe
	eventFormCallback.event = event
	eventFormCallback.formGroup = formGroup

	eventFormCallback.CreationMode = (event == nil)

	return
}

type EventFormCallback struct {
	event *models.Event

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (eventFormCallback *EventFormCallback) OnSave() {

	// log.Println("EventFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	eventFormCallback.probe.formStage.Checkout()

	if eventFormCallback.event == nil {
		eventFormCallback.event = new(models.Event).Stage(eventFormCallback.probe.stageOfInterest)
	}
	event_ := eventFormCallback.event
	_ = event_

	for _, formDiv := range eventFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(event_.Name), formDiv)
		case "Duration":
			FormDivBasicFieldToField(&(event_.Duration), formDiv)
		}
	}

	// manage the suppress operation
	if eventFormCallback.formGroup.HasSuppressButtonBeenPressed {
		event_.Unstage(eventFormCallback.probe.stageOfInterest)
	}

	eventFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.Event](
		eventFormCallback.probe,
	)
	eventFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if eventFormCallback.CreationMode || eventFormCallback.formGroup.HasSuppressButtonBeenPressed {
		eventFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(eventFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__EventFormCallback(
			nil,
			eventFormCallback.probe,
			newFormGroup,
		)
		event := new(models.Event)
		FillUpForm(event, newFormGroup, eventFormCallback.probe)
		eventFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(eventFormCallback.probe)
}
func __gong__New__StatusFormCallback(
	status *models.Status,
	probe *Probe,
	formGroup *table.FormGroup,
) (statusFormCallback *StatusFormCallback) {
	statusFormCallback = new(StatusFormCallback)
	statusFormCallback.probe = probe
	statusFormCallback.status = status
	statusFormCallback.formGroup = formGroup

	statusFormCallback.CreationMode = (status == nil)

	return
}

type StatusFormCallback struct {
	status *models.Status

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (statusFormCallback *StatusFormCallback) OnSave() {

	// log.Println("StatusFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	statusFormCallback.probe.formStage.Checkout()

	if statusFormCallback.status == nil {
		statusFormCallback.status = new(models.Status).Stage(statusFormCallback.probe.stageOfInterest)
	}
	status_ := statusFormCallback.status
	_ = status_

	for _, formDiv := range statusFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(status_.Name), formDiv)
		case "CurrentCommand":
			FormDivEnumStringFieldToField(&(status_.CurrentCommand), formDiv)
		case "CompletionDate":
			FormDivBasicFieldToField(&(status_.CompletionDate), formDiv)
		case "CurrentSpeedCommand":
			FormDivEnumStringFieldToField(&(status_.CurrentSpeedCommand), formDiv)
		case "SpeedCommandCompletionDate":
			FormDivBasicFieldToField(&(status_.SpeedCommandCompletionDate), formDiv)
		}
	}

	// manage the suppress operation
	if statusFormCallback.formGroup.HasSuppressButtonBeenPressed {
		status_.Unstage(statusFormCallback.probe.stageOfInterest)
	}

	statusFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.Status](
		statusFormCallback.probe,
	)
	statusFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if statusFormCallback.CreationMode || statusFormCallback.formGroup.HasSuppressButtonBeenPressed {
		statusFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(statusFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__StatusFormCallback(
			nil,
			statusFormCallback.probe,
			newFormGroup,
		)
		status := new(models.Status)
		FillUpForm(status, newFormGroup, statusFormCallback.probe)
		statusFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(statusFormCallback.probe)
}
func __gong__New__UpdateStateFormCallback(
	updatestate *models.UpdateState,
	probe *Probe,
	formGroup *table.FormGroup,
) (updatestateFormCallback *UpdateStateFormCallback) {
	updatestateFormCallback = new(UpdateStateFormCallback)
	updatestateFormCallback.probe = probe
	updatestateFormCallback.updatestate = updatestate
	updatestateFormCallback.formGroup = formGroup

	updatestateFormCallback.CreationMode = (updatestate == nil)

	return
}

type UpdateStateFormCallback struct {
	updatestate *models.UpdateState

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (updatestateFormCallback *UpdateStateFormCallback) OnSave() {

	// log.Println("UpdateStateFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	updatestateFormCallback.probe.formStage.Checkout()

	if updatestateFormCallback.updatestate == nil {
		updatestateFormCallback.updatestate = new(models.UpdateState).Stage(updatestateFormCallback.probe.stageOfInterest)
	}
	updatestate_ := updatestateFormCallback.updatestate
	_ = updatestate_

	for _, formDiv := range updatestateFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(updatestate_.Name), formDiv)
		case "Duration":
			FormDivBasicFieldToField(&(updatestate_.Duration), formDiv)
		case "Period":
			FormDivBasicFieldToField(&(updatestate_.Period), formDiv)
		}
	}

	// manage the suppress operation
	if updatestateFormCallback.formGroup.HasSuppressButtonBeenPressed {
		updatestate_.Unstage(updatestateFormCallback.probe.stageOfInterest)
	}

	updatestateFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.UpdateState](
		updatestateFormCallback.probe,
	)
	updatestateFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if updatestateFormCallback.CreationMode || updatestateFormCallback.formGroup.HasSuppressButtonBeenPressed {
		updatestateFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(updatestateFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__UpdateStateFormCallback(
			nil,
			updatestateFormCallback.probe,
			newFormGroup,
		)
		updatestate := new(models.UpdateState)
		FillUpForm(updatestate, newFormGroup, updatestateFormCallback.probe)
		updatestateFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(updatestateFormCallback.probe)
}
