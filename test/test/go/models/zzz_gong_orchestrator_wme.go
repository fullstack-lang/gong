// generated code - do not edit
package models

// insertion point
// AstructOrchestratorWithMouseEvent
type AstructOrchestratorWithMouseEvent struct {
}

func (orchestrator *AstructOrchestratorWithMouseEvent) OnAfterUpdateWithMouseEvent(
	gongsvgStage *Stage,
	stagedAstruct, backRepoAstruct *Astruct, mouseEvent *Gong__MouseEvent) {

	stagedAstruct.OnAfterUpdateWithMouseEvent(gongsvgStage, backRepoAstruct, mouseEvent)
}

func SetOrchestratorOnAfterUpdateWithMouseEvent[Type Gongstruct](stage *Stage) {

	var ret Type

	switch any(ret).(type) {
	// insertion point
	case Astruct:
		stage.OnAfterAstructUpdateWithMouseEventCallback = new(AstructOrchestratorWithMouseEvent)

	}
}

