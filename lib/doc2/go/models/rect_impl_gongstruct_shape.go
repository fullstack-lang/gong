package models

import (
	gongsvg_models "github.com/fullstack-lang/gong/lib/svg/go/models"
)

// RectImplGongstructShape
// it meets the interface of RectImplInterface
type RectImplGongstructShape struct {
	gongstructShape *GongStructShape
	gongdocStage    *Stage
}

func NewRectImplGongstructShape(
	gongstructShape *GongStructShape,
	gongdocStage *Stage) (rectImplGongstructShape *RectImplGongstructShape) {

	rectImplGongstructShape = new(RectImplGongstructShape)
	rectImplGongstructShape.gongstructShape = gongstructShape
	rectImplGongstructShape.gongdocStage = gongdocStage

	return
}

func (rectImplGongstructShape *RectImplGongstructShape) RectUpdated(updatedRect *gongsvg_models.Rect) {

	// log.Println("RectImplGongstructShape:RectUpdated")

	// update the shape
	rectImplGongstructShape.gongstructShape.X = updatedRect.X
	rectImplGongstructShape.gongstructShape.Y = updatedRect.Y
	rectImplGongstructShape.gongstructShape.Width = updatedRect.Width
	rectImplGongstructShape.gongstructShape.Height = updatedRect.Height

	rectImplGongstructShape.gongdocStage.CommitWithSuspendedCallbacks()
}
