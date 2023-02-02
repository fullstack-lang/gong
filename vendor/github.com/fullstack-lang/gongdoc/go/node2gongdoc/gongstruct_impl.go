package node2gongdoc

import (
	gong_models "github.com/fullstack-lang/gong/go/models"
	gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"
)

type GongStructImpl struct {
	gongStruct *gong_models.GongStruct
	NodeImpl
}

func (gongStructImpl *GongStructImpl) OnAfterUpdate(
	stage *gongdoc_models.StageStruct,
	stagedNode, frontNode *gongdoc_models.Node) {
	// if node is unchecked
	if stagedNode.IsChecked && !frontNode.IsChecked {

		// get the latest version of the diagram before modifying it
		stage.Checkout()

		// remove the classshape from the selected diagram
		classDiagram := gongStructImpl.nodeCb.GetSelectedClassdiagram()
		classDiagram.RemoveClassshape(stagedNode.Name)

	}

	// if node is checked, add classshape
	if !stagedNode.IsChecked && frontNode.IsChecked {

		// get the latest version of the diagram before modifying it
		stage.Checkout()

		classDiagram := gongStructImpl.nodeCb.GetSelectedClassdiagram()
		classDiagram.AddClassshape(frontNode.Name)

	}
}

func (gongStructImpl *GongStructImpl) OnAfterDelete(
	stage *gongdoc_models.StageStruct,
	stagedNode, frontNode *gongdoc_models.Node) {
}
