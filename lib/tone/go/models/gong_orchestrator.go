// generated code - do not edit
package models

// insertion point
// PlayerOrchestrator
type PlayerOrchestrator struct {
}

func (orchestrator *PlayerOrchestrator) OnAfterUpdate(
	gongsvgStage *StageStruct,
	stagedPlayer, backRepoPlayer *Player) {

	stagedPlayer.OnAfterUpdate(gongsvgStage, stagedPlayer, backRepoPlayer)
}

func SetOrchestratorOnAfterUpdate[Type Gongstruct](stage *StageStruct) {

	var ret Type

	switch any(ret).(type) {
	// insertion point
	case Player:
		stage.OnAfterPlayerUpdateCallback = new(PlayerOrchestrator)

	}

}
