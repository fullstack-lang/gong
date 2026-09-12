// generated code - do not edit
package models

// insertion point
// CommandOrchestrator
type CommandOrchestrator struct {
}

func (orchestrator *CommandOrchestrator) OnAfterUpdate(
	gongsvgStage *Stage,
	stagedCommand, backRepoCommand *Command) {

	stagedCommand.OnAfterUpdate(gongsvgStage, stagedCommand, backRepoCommand)
}

// SetOrchestratorOnAfterUpdate is the Stage method for setting orchestrators.
func (stage *Stage) SetOrchestratorOnAfterUpdate[Type Gongstruct]() {

	var ret Type

	switch any(ret).(type) {
	// insertion point
	case Command:
		stage.OnAfterCommandUpdateCallback = new(CommandOrchestrator)

	}

}

// SetOrchestratorOnAfterUpdate is a backward-compatible package-level forwarder.
func SetOrchestratorOnAfterUpdate[Type Gongstruct](stage *Stage) {
	stage.SetOrchestratorOnAfterUpdate[Type]()
}
