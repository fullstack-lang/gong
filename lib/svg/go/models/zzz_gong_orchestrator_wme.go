// generated code - do not edit
package models

// insertion point
// RectOrchestratorWithMouseEvent
type RectOrchestratorWithMouseEvent struct {
}

func (orchestrator *RectOrchestratorWithMouseEvent) OnAfterUpdateWithMouseEvent(
	gongsvgStage *Stage,
	stagedRect, backRepoRect *Rect, mouseEvent *Gong__MouseEvent) {

	stagedRect.OnAfterUpdateWithMouseEvent(gongsvgStage, backRepoRect, mouseEvent)
}

func SetOrchestratorOnAfterUpdateWithMouseEvent[Type Gongstruct](stage *Stage) {

	var ret Type

	switch any(ret).(type) {
	// insertion point
	case Rect:
		stage.OnAfterRectUpdateWithMouseEventCallback = new(RectOrchestratorWithMouseEvent)

	}
}

