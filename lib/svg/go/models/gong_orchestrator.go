// generated code - do not edit
package models

// insertion point
// LineOrchestrator
type LineOrchestrator struct {
}

func (orchestrator *LineOrchestrator) OnAfterUpdate(
	gongsvgStage *Stage,
	stagedLine, backRepoLine *Line) {

	stagedLine.OnAfterUpdate(gongsvgStage, stagedLine, backRepoLine)
}
// LinkOrchestrator
type LinkOrchestrator struct {
}

func (orchestrator *LinkOrchestrator) OnAfterUpdate(
	gongsvgStage *Stage,
	stagedLink, backRepoLink *Link) {

	stagedLink.OnAfterUpdate(gongsvgStage, stagedLink, backRepoLink)
}
// LinkAnchoredTextOrchestrator
type LinkAnchoredTextOrchestrator struct {
}

func (orchestrator *LinkAnchoredTextOrchestrator) OnAfterUpdate(
	gongsvgStage *Stage,
	stagedLinkAnchoredText, backRepoLinkAnchoredText *LinkAnchoredText) {

	stagedLinkAnchoredText.OnAfterUpdate(gongsvgStage, stagedLinkAnchoredText, backRepoLinkAnchoredText)
}
// RectOrchestrator
type RectOrchestrator struct {
}

func (orchestrator *RectOrchestrator) OnAfterUpdate(
	gongsvgStage *Stage,
	stagedRect, backRepoRect *Rect) {

	stagedRect.OnAfterUpdate(gongsvgStage, stagedRect, backRepoRect)
}
// SVGOrchestrator
type SVGOrchestrator struct {
}

func (orchestrator *SVGOrchestrator) OnAfterUpdate(
	gongsvgStage *Stage,
	stagedSVG, backRepoSVG *SVG) {

	stagedSVG.OnAfterUpdate(gongsvgStage, stagedSVG, backRepoSVG)
}

func SetOrchestratorOnAfterUpdate[Type Gongstruct](stage *Stage) {

	var ret Type

	switch any(ret).(type) {
	// insertion point
	case Line:
		stage.OnAfterLineUpdateCallback = new(LineOrchestrator)
	case Link:
		stage.OnAfterLinkUpdateCallback = new(LinkOrchestrator)
	case LinkAnchoredText:
		stage.OnAfterLinkAnchoredTextUpdateCallback = new(LinkAnchoredTextOrchestrator)
	case Rect:
		stage.OnAfterRectUpdateCallback = new(RectOrchestrator)
	case SVG:
		stage.OnAfterSVGUpdateCallback = new(SVGOrchestrator)

	}

}
