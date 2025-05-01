package models

import (
	"log"

	gongsvg_models "github.com/fullstack-lang/gong/lib/svg/go/models"
)

// RectImplGongenumShape
// it meets the interface of RectImplInterface
type RectImplGongenumShape struct {
	gongenumShape *GongEnumShape
	gongdocStage  *Stage
}

func NewRectImplGongenumShape(
	gongEnumShape *GongEnumShape,
	gongdocStage *Stage) (rectImplGongenumShape *RectImplGongenumShape) {

	rectImplGongenumShape = new(RectImplGongenumShape)
	rectImplGongenumShape.gongenumShape = gongEnumShape
	rectImplGongenumShape.gongdocStage = gongdocStage

	return
}

func (rectImplGongenumShape *RectImplGongenumShape) RectUpdated(updatedRect *gongsvg_models.Rect) {

	log.Println("RectImplGongenumShape:RectUpdated")

	// update the shape
	rectImplGongenumShape.gongenumShape.Position.X = updatedRect.X
	rectImplGongenumShape.gongenumShape.Position.Y = updatedRect.Y
	rectImplGongenumShape.gongenumShape.Width = updatedRect.Width
	rectImplGongenumShape.gongenumShape.Height = updatedRect.Height

	rectImplGongenumShape.gongdocStage.CommitWithSuspendedCallbacks()
}
