// generated code - do not edit
package models

// insertion point
// AstructOrchestrator
type AstructOrchestrator struct {
}

func (orchestrator *AstructOrchestrator) OnAfterUpdate(
	gongsvgStage *Stage,
	stagedAstruct, backRepoAstruct *Astruct) {

	stagedAstruct.OnAfterUpdate(gongsvgStage, stagedAstruct, backRepoAstruct)
}

type AstructOrchestratorWithMouseEvent struct {
}

func (orchestrator *AstructOrchestratorWithMouseEvent) OnAfterUpdateWithMouseEvent(
	gongsvgStage *Stage,
	stagedAstruct, backRepoAstruct *Astruct, mouseEvent *MouseEvent) {

	stagedAstruct.OnAfterUpdateWithMouseEvent(gongsvgStage, backRepoAstruct, mouseEvent)
}

// BstructOrchestrator
type BstructOrchestrator struct {
}

func (orchestrator *BstructOrchestrator) OnAfterUpdate(
	gongsvgStage *Stage,
	stagedBstruct, backRepoBstruct *Bstruct) {

	stagedBstruct.OnAfterUpdate(gongsvgStage, stagedBstruct, backRepoBstruct)
}

func SetOrchestratorOnAfterUpdate[Type Gongstruct](stage *Stage) {

	var ret Type

	switch any(ret).(type) {
	// insertion point
	case Astruct:
		stage.OnAfterAstructUpdateCallback = new(AstructOrchestrator)
	case Bstruct:
		stage.OnAfterBstructUpdateCallback = new(BstructOrchestrator)

	}

}

func SetOrchestratorOnAfterUpdateWithMouseEvent[Type Gongstruct](stage *Stage) {

	var ret Type

	switch any(ret).(type) {
	// insertion point
	case Astruct:
		stage.OnAfterAstructUpdateWithMouseEventCallback = new(AstructOrchestratorWithMouseEvent)

	}

}
