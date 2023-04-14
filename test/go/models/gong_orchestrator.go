package models

// insertion point
// AstructOrchestrator
type AstructOrchestrator struct {
}

func (orchestrator *AstructOrchestrator) OnAfterUpdate(
	gongsvgStage *StageStruct,
	stagedAstruct, backRepoAstruct *Astruct) {

	stagedAstruct.OnAfterUpdate(gongsvgStage, stagedAstruct, backRepoAstruct)
}
// BstructOrchestrator
type BstructOrchestrator struct {
}

func (orchestrator *BstructOrchestrator) OnAfterUpdate(
	gongsvgStage *StageStruct,
	stagedBstruct, backRepoBstruct *Bstruct) {

	stagedBstruct.OnAfterUpdate(gongsvgStage, stagedBstruct, backRepoBstruct)
}

func SetOrchestratorOnAfterUpdate[Type Gongstruct](stage *StageStruct) {

	var ret Type

	switch any(ret).(type) {
	// insertion point
	case Astruct:
		stage.OnAfterAstructUpdateCallback = new(AstructOrchestrator)
	case Bstruct:
		stage.OnAfterBstructUpdateCallback = new(BstructOrchestrator)

	}

}
