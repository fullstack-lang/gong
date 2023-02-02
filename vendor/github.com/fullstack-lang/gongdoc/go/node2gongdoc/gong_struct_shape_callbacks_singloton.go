package node2gongdoc

import (
	"log"

	gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"
)

type GongStructShapeCallbacksSingloton struct {
	GongStructShapeCallback GongStructShapeCallback
}

func (gongStructShapeCallbacksSingloton *GongStructShapeCallbacksSingloton) OnAfterUpdate(
	stage *gongdoc_models.StageStruct,
	stagedGongStructShape, frontGongStructShape *gongdoc_models.GongStructShape) {

	if stagedGongStructShape.IsSelected != frontGongStructShape.IsSelected {

		// reset the IsSelected to false
		stagedGongStructShape.Checkout()
		stagedGongStructShape.IsSelected = false
		stagedGongStructShape.Commit()

		log.Println("UML Shape selected ", stagedGongStructShape.Identifier)
		if gongStructShapeCallbacksSingloton.GongStructShapeCallback != nil {
			gongStructShapeCallbacksSingloton.GongStructShapeCallback.HasSelected(stagedGongStructShape.Identifier)
		}
	}
}

type GongStructShapeCallback interface {
	HasSelected(gongstructName string)
}
