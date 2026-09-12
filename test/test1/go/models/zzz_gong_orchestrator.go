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

// BstructOrchestrator
type BstructOrchestrator struct {
}

func (orchestrator *BstructOrchestrator) OnAfterUpdate(
	gongsvgStage *Stage,
	stagedBstruct, backRepoBstruct *Bstruct) {

	stagedBstruct.OnAfterUpdate(gongsvgStage, stagedBstruct, backRepoBstruct)
}

// SetOrchestratorOnAfterUpdate is the Stage method for setting orchestrators.
func (stage *Stage) SetOrchestratorOnAfterUpdate[Type Gongstruct]() {

	var ret Type

	switch any(ret).(type) {
	// insertion point
	case Astruct:
		stage.OnAfterAstructUpdateCallback = new(AstructOrchestrator)
	case Bstruct:
		stage.OnAfterBstructUpdateCallback = new(BstructOrchestrator)

	}

}

// SetOrchestratorOnAfterUpdate is a backward-compatible package-level forwarder.
func SetOrchestratorOnAfterUpdate[Type Gongstruct](stage *Stage) {
	stage.SetOrchestratorOnAfterUpdate[Type]()
}
