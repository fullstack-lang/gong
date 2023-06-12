package doc2svg

import (
	"log"

	gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"
	gongsvg_models "github.com/fullstack-lang/gongsvg/go/models"
)

// RectImplGongenumShape
// it meets the interface of RectImplInterface
type RectImplGongenumShape struct {
	gongenumShape *gongdoc_models.GongEnumShape
	gongdocStage  *gongdoc_models.StageStruct
}

func NewRectImplGongenumShape(
	gongEnumShape *gongdoc_models.GongEnumShape,
	gongdocStage *gongdoc_models.StageStruct) (rectImplGongenumShape *RectImplGongenumShape) {

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
	rectImplGongenumShape.gongenumShape.Heigth = updatedRect.Height

	rectImplGongenumShape.gongdocStage.CommitWithSuspendedCallbacks()
}
