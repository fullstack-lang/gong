package doc2svg

import (
	gongdoc_models "github.com/fullstack-lang/gong/lib/doc/go/models"
	gongsvg_models "github.com/fullstack-lang/gong/lib/svg/go/models"
)

// RectImplGongstructShape
// it meets the interface of RectImplInterface
type RectImplGongstructShape struct {
	gongstructShape *gongdoc_models.GongStructShape
	gongdocStage    *gongdoc_models.Stage
}

func NewRectImplGongstructShape(
	gongstructShape *gongdoc_models.GongStructShape,
	gongdocStage *gongdoc_models.Stage) (rectImplGongstructShape *RectImplGongstructShape) {

	rectImplGongstructShape = new(RectImplGongstructShape)
	rectImplGongstructShape.gongstructShape = gongstructShape
	rectImplGongstructShape.gongdocStage = gongdocStage

	return
}

func (rectImplGongstructShape *RectImplGongstructShape) RectUpdated(updatedRect *gongsvg_models.Rect) {

	// log.Println("RectImplGongstructShape:RectUpdated")

	// update the shape
	rectImplGongstructShape.gongstructShape.Position.X = updatedRect.X
	rectImplGongstructShape.gongstructShape.Position.Y = updatedRect.Y
	rectImplGongstructShape.gongstructShape.Width = updatedRect.Width
	rectImplGongstructShape.gongstructShape.Height = updatedRect.Height

	rectImplGongstructShape.gongdocStage.CommitWithSuspendedCallbacks()
}
