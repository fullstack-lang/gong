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

func SetOrchestratorOnAfterUpdate[Type Gongstruct](stage *Stage) {

	var ret Type

	switch any(ret).(type) {
	// insertion point
	case Command:
		stage.OnAfterCommandUpdateCallback = new(CommandOrchestrator)

	}

}
