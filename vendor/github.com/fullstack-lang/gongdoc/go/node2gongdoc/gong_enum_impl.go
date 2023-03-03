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

		// remove the gongstructshape from the selected diagram
		classDiagram := gongEnumImpl.nodeCb.GetSelectedClassdiagram()

		// get the referenced gongstructs
		for _, gongEnumShape := range classDiagram.GongEnumShapes {
			if gongdoc_models.IdentifierToGongObjectName(gongEnumShape.Identifier) == stagedNode.Name {
				classDiagram.RemoveGongEnumShape(gongdoc_models.IdentifierToGongObjectName(gongEnumShape.Identifier))
			}

		}

	}

	// if node is checked, add gongstructshape
	if !stagedNode.IsChecked && frontNode.IsChecked {

		// get the latest version of the diagram before modifying it
		stage.Checkout()

		classDiagram := gongEnumImpl.nodeCb.GetSelectedClassdiagram()
		classDiagram.AddGongEnumShape(gongEnumImpl.nodeCb.diagramPackage, frontNode.Name)
	}

}

func (gongEnumImpl *GongEnumImpl) OnAfterDelete(
	stage *gongdoc_models.StageStruct,
	stagedNode, frontNode *gongdoc_models.Node) {
}
