package models

// insertion point
// LineOrchestrator
type LineOrchestrator struct {
}

func (orchestrator *LineOrchestrator) OnAfterUpdate(
	gongsvgStage *StageStruct,
	stagedLine, backRepoLine *Line) {

	stagedLine.OnAfterUpdate(gongsvgStage, stagedLine, backRepoLine)
}
// LinkOrchestrator
type LinkOrchestrator struct {
}

func (orchestrator *LinkOrchestrator) OnAfterUpdate(
	gongsvgStage *StageStruct,
	stagedLink, backRepoLink *Link) {

	stagedLink.OnAfterUpdate(gongsvgStage, stagedLink, backRepoLink)
}
// LinkAnchoredTextOrchestrator
type LinkAnchoredTextOrchestrator struct {
}

func (orchestrator *LinkAnchoredTextOrchestrator) OnAfterUpdate(
	gongsvgStage *StageStruct,
	stagedLinkAnchoredText, backRepoLinkAnchoredText *LinkAnchoredText) {

	stagedLinkAnchoredText.OnAfterUpdate(gongsvgStage, stagedLinkAnchoredText, backRepoLinkAnchoredText)
}
// RectOrchestrator
type RectOrchestrator struct {
}

func (orchestrator *RectOrchestrator) OnAfterUpdate(
	gongsvgStage *StageStruct,
	stagedRect, backRepoRect *Rect) {

	stagedRect.OnAfterUpdate(gongsvgStage, stagedRect, backRepoRect)
}
// SVGOrchestrator
type SVGOrchestrator struct {
}

func (orchestrator *SVGOrchestrator) OnAfterUpdate(
	gongsvgStage *StageStruct,
	stagedSVG, backRepoSVG *SVG) {

	stagedSVG.OnAfterUpdate(gongsvgStage, stagedSVG, backRepoSVG)
}

func SetOrchestratorOnAfterUpdate[Type Gongstruct](stage *StageStruct) {

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
