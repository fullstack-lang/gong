package node2gongdoc

import (
	gong_models "github.com/fullstack-lang/gong/go/models"
	gongtree_models "github.com/fullstack-lang/gongtree/go/models"
)

type NodeImplGongstruct struct {
	NodeImplGongObjectAbstract
	gongStruct *gong_models.GongStruct
}

func NewNodeImplGongstruct(
	gongStruct *gong_models.GongStruct,
	nodeImplGongObjectAbstract NodeImplGongObjectAbstract,
) (nodeImplGongstruct *NodeImplGongstruct) {

	nodeImplGongstruct = new(NodeImplGongstruct)
	nodeImplGongstruct.NodeImplGongObjectAbstract = nodeImplGongObjectAbstract
	nodeImplGongstruct.gongStruct = gongStruct

	return
}

func (nodeImplGongstruct *NodeImplGongstruct) OnAfterUpdate(
	gongtreeStage *gongtree_models.StageStruct,
	stagedNode, frontNode *gongtree_models.Node) {

	gongdocStage := nodeImplGongstruct.diagramPackage.Stage_

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

		// remove the gongstructshape from the selected diagram
		classDiagram := nodeImplGongstruct.diagramPackage.SelectedClassdiagram
		classDiagram.RemoveGongStructShape(gongdocStage, stagedNode.Name)

	}

	// if node is checked, add gongstructshape
	if !stagedNode.IsChecked && frontNode.IsChecked {

		// get the latest version of the diagram before modifying it
		gongdocStage.Checkout()

		classDiagram := nodeImplGongstruct.diagramPackage.SelectedClassdiagram
		classDiagram.AddGongStructShape(gongdocStage, nodeImplGongstruct.diagramPackage, frontNode.Name)
	}

	computeGongNodesConfigurations(
		gongdocStage,
		nodeImplGongstruct.diagramPackage.SelectedClassdiagram,
		nodeImplGongstruct.treeOfGongObjects)

	gongdocStage.Commit()
}
