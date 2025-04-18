package doc2svg

import (
	"log"

	gongdoc_models "github.com/fullstack-lang/gong/lib/doc/go/models"
	gongsvg_models "github.com/fullstack-lang/gong/lib/svg/go/models"
)

// RectImplGongenumShape
// it meets the interface of RectImplInterface
type RectImplGongenumShape struct {
	gongenumShape *gongdoc_models.GongEnumShape
	gongdocStage  *gongdoc_models.Stage
}

func NewRectImplGongenumShape(
	gongEnumShape *gongdoc_models.GongEnumShape,
	gongdocStage *gongdoc_models.Stage) (rectImplGongenumShape *RectImplGongenumShape) {

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
