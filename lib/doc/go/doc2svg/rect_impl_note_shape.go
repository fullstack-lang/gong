package doc2svg

import (
	"log"

	gongdoc_models "github.com/fullstack-lang/gong/lib/doc/go/models"
	gongsvg_models "github.com/fullstack-lang/gong/lib/svg/go/models"
)

// RectImplNoteShape
// it meets the interface of RectImplInterface
type RectImplNoteShape struct {
	noteShape    *gongdoc_models.NoteShape
	gongdocStage *gongdoc_models.Stage
}

func NewRectImplNoteShape(
	noteShape *gongdoc_models.NoteShape,
	gongdocStage *gongdoc_models.Stage) (rectImplNoteShape *RectImplNoteShape) {

	rectImplNoteShape = new(RectImplNoteShape)
	rectImplNoteShape.noteShape = noteShape
	rectImplNoteShape.gongdocStage = gongdocStage

	return
}

func (rectImplNoteShape *RectImplNoteShape) RectUpdated(updatedRect *gongsvg_models.Rect) {

	log.Println("RectImplNoteShape:RectUpdated")

	// update the shape
	rectImplNoteShape.noteShape.X = updatedRect.X
	rectImplNoteShape.noteShape.Y = updatedRect.Y
	rectImplNoteShape.noteShape.Width = updatedRect.Width
	rectImplNoteShape.noteShape.Height = updatedRect.Height

	rectImplNoteShape.gongdocStage.CommitWithSuspendedCallbacks()
}
