package doc2svg

import (
	"log"

	gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"
	gongsvg_models "github.com/fullstack-lang/gongsvg/go/models"
)

// RectImplNoteShape
// it meets the interface of RectImplInterface
type RectImplNoteShape struct {
	noteShape    *gongdoc_models.NoteShape
	gongdocStage *gongdoc_models.StageStruct
}

func NewRectImplNoteShape(
	noteShape *gongdoc_models.NoteShape,
	gongdocStage *gongdoc_models.StageStruct) (rectImplNoteShape *RectImplNoteShape) {

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
	rectImplNoteShape.noteShape.Heigth = updatedRect.Height

	rectImplNoteShape.gongdocStage.CommitWithSuspendedCallbacks()
}
