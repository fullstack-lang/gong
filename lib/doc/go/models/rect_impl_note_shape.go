package models

import (
	"log"

	gongsvg_models "github.com/fullstack-lang/gong/lib/svg/go/models"
)

// RectImplNoteShape
// it meets the interface of RectImplInterface
type RectImplNoteShape struct {
	noteShape *GongNoteShape
	stager    *Stager
}

func NewRectImplNoteShape(
	noteShape *GongNoteShape,
	stager *Stager) (rectImplNoteShape *RectImplNoteShape) {

	rectImplNoteShape = new(RectImplNoteShape)
	rectImplNoteShape.noteShape = noteShape
	rectImplNoteShape.stager = stager

	return
}

func (p *RectImplNoteShape) RectUpdated(updatedRect *gongsvg_models.Rect) {

	log.Println("RectImplNoteShape:RectUpdated")

	diffSize :=
		p.noteShape.Width != updatedRect.Width ||
			p.noteShape.Height != updatedRect.Height

	diffPosition :=
		p.noteShape.X != updatedRect.X ||
			p.noteShape.Y != updatedRect.Y

	// update the shape
	if diffSize {
		p.noteShape.Width = updatedRect.Width
		p.noteShape.Height = updatedRect.Height
		p.noteShape.X = updatedRect.X
		p.noteShape.Y = updatedRect.Y
		// one need to recomputes the text within the note shape
		p.stager.stage.CommitWithSuspendedCallbacks()
		p.stager.UpdateAndCommitSVGStage()

		return
	}

	if diffPosition {
		p.noteShape.X = updatedRect.X
		p.noteShape.Y = updatedRect.Y
		p.stager.stage.CommitWithSuspendedCallbacks()

		return
	}

}
