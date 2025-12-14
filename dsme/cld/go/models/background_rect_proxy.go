package models

import svg "github.com/fullstack-lang/gong/lib/svg/go/models"

type BackgroundRectProxy struct {
	stager  *Stager
	diagram *Diagram
}

func (p *BackgroundRectProxy) RectUpdated(updatedRect *svg.Rect) {

	p.diagram.IsEditable = !p.diagram.IsEditable

	p.stager.stage.Commit()
}
