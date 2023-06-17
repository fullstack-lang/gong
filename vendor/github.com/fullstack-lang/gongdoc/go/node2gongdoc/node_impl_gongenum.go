package node2gongdoc

import (
	gong_models "github.com/fullstack-lang/gong/go/models"
	gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"
	gongtree_models "github.com/fullstack-lang/gongtree/go/models"
)

type NodeImplGongEnum struct {
	NodeImplGongObjectAbstract
	gongEnum *gong_models.GongEnum
}

func NewNodeImplGongEnum(
	gongEnum *gong_models.GongEnum,
	nodeImplGongObjectAbstract NodeImplGongObjectAbstract,
) (nodeImplGongEnum *NodeImplGongEnum) {

	nodeImplGongEnum = new(NodeImplGongEnum)
	nodeImplGongEnum.NodeImplGongObjectAbstract = nodeImplGongObjectAbstract
	nodeImplGongEnum.gongEnum = gongEnum

	return
}

func (nodeImplGongEnum *NodeImplGongEnum) OnAfterUpdate(
	gongtreeStage *gongtree_models.StageStruct,
	stagedNode, frontNode *gongtree_models.Node) {

	classDiagram := nodeImplGongEnum.diagramPackage.SelectedClassdiagram
	gongdocStage := nodeImplGongEnum.diagramPackage.Stage_

	// setting the value of the staged node	to the new value
	// otherwise, the expansion memory is lost
	if stagedNode.IsExpanded != frontNode.IsExpanded {
		stagedNode.IsExpanded = frontNode.IsExpanded
		return
	}

	// if node is unchecked
	if stagedNode.IsChecked && !frontNode.IsChecked {

		// get the latest version of the diagram before modifying it
		gongdocStage.Checkout()

		// get the referenced gongstructs
		for _, gongEnumShape := range classDiagram.GongEnumShapes {
			if gongdoc_models.IdentifierToGongObjectName(gongEnumShape.Identifier) == stagedNode.Name {
				classDiagram.RemoveGongEnumShape(gongdocStage, gongdoc_models.IdentifierToGongObjectName(gongEnumShape.Identifier))
			}

		}

	}

	// if node is checked, add gongstructshape
	if !stagedNode.IsChecked && frontNode.IsChecked {

		// get the latest version of the diagram before modifying it
		gongdocStage.Checkout()

		classDiagram.AddGongEnumShape(gongdocStage, nodeImplGongEnum.diagramPackage, frontNode.Name)
	}

	computeGongNodesConfigurations(
		gongdocStage,
		classDiagram,
		nodeImplGongEnum.treeOfGongObjects)
	gongdocStage.Commit()
}
