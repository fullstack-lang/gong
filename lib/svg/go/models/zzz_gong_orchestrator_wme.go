// generated code - do not edit
package models

// insertion point
// LinkOrchestratorWithMouseEvent
type LinkOrchestratorWithMouseEvent struct {
}

func (orchestrator *LinkOrchestratorWithMouseEvent) OnAfterUpdateWithMouseEvent(
	gongsvgStage *Stage,
	stagedLink, backRepoLink *Link, mouseEvent *Gong__MouseEvent) {

	stagedLink.OnAfterUpdateWithMouseEvent(gongsvgStage, backRepoLink, mouseEvent)
}
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
	case Link:
		stage.OnAfterLinkUpdateWithMouseEventCallback = new(LinkOrchestratorWithMouseEvent)
	case Rect:
		stage.OnAfterRectUpdateWithMouseEventCallback = new(RectOrchestratorWithMouseEvent)

	}
}

