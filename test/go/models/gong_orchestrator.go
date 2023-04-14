package models

// AstructOrchestrator
type AstructOrchestrator struct {
}

func (orchestrator *AstructOrchestrator) OnAfterUpdate(
	gongsvgStage *StageStruct,
	stagedAstruct, backRepoAstruct *Astruct) {

	stagedAstruct.OnAfterUpdate(gongsvgStage, stagedAstruct, backRepoAstruct)
}

func SetOrchestratorOnAfterUpdate[Type Gongstruct](stage *StageStruct) {

	var ret Type

	switch any(ret).(type) {

	case Astruct:
		stage.OnAfterAstructUpdateCallback = new(AstructOrchestrator)
	}

}
