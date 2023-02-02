package node2gongdoc

import (
	gong_models "github.com/fullstack-lang/gong/go/models"
	gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"
)

type GongEnumImpl struct {
	gongEnum *gong_models.GongEnum
	NodeImpl
}

func (gongEnumImpl *GongEnumImpl) OnAfterUpdate(
	stage *gongdoc_models.StageStruct,
	stagedNode, frontNode *gongdoc_models.Node) {

	// if node is unchecked
	if stagedNode.IsChecked && !frontNode.IsChecked {

		// get the latest version of the diagram before modifying it
		stage.Checkout()

		// remove the classshape from the selected diagram
		classDiagram := gongEnumImpl.nodeCb.GetSelectedClassdiagram()

		// get the referenced gongstructs
		for _, classshape := range classDiagram.GongStructShapes {
			if gongdoc_models.IdentifierToGongStructName(classshape.Identifier) == stagedNode.Name {
				classDiagram.RemoveClassshape(gongdoc_models.IdentifierToGongStructName(classshape.Identifier))
			}

		}

	}

	// if node is checked, add classshape
	if !stagedNode.IsChecked && frontNode.IsChecked {

		// get the latest version of the diagram before modifying it
		stage.Checkout()

		classDiagram := gongEnumImpl.nodeCb.GetSelectedClassdiagram()
		classDiagram.AddClassshape(frontNode.Name)
	}

}

func (gongEnumImpl *GongEnumImpl) OnAfterDelete(
	stage *gongdoc_models.StageStruct,
	stagedNode, frontNode *gongdoc_models.Node) {
}
