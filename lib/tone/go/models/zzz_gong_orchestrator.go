// generated code - do not edit
package models

// insertion point
// PlayerOrchestrator
type PlayerOrchestrator struct {
}

func (orchestrator *PlayerOrchestrator) OnAfterUpdate(
	gongsvgStage *Stage,
	stagedPlayer, backRepoPlayer *Player) {

	stagedPlayer.OnAfterUpdate(gongsvgStage, stagedPlayer, backRepoPlayer)
}

// SetOrchestratorOnAfterUpdate is the Stage method for setting orchestrators.
func (stage *Stage) SetOrchestratorOnAfterUpdate[Type Gongstruct]() {

	var ret Type

	switch any(ret).(type) {
	// insertion point
	case Player:
		stage.OnAfterPlayerUpdateCallback = new(PlayerOrchestrator)

	}

}

// SetOrchestratorOnAfterUpdate is a backward-compatible package-level forwarder.
func SetOrchestratorOnAfterUpdate[Type Gongstruct](stage *Stage) {
	stage.SetOrchestratorOnAfterUpdate[Type]()
}
