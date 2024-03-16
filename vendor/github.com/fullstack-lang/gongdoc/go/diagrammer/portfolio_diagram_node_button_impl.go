package diagrammer

import (
	"slices"

	gongtree_models "github.com/fullstack-lang/gongtree/go/models"
)

type ButtonType int

const (
	DUPLICATE ButtonType = iota
	EDIT_CANCEL
	EDIT
	REMOVE
	RENAME_CANCEL
	RENAME
	SAVE
)

type PortfolioDiagramNodeButtonImpl struct {
	portfolioDiagramNode PortfolioDiagramNode

	diagrammer *Diagrammer

	treeNode *gongtree_models.Node

	treeStage *gongtree_models.StageStruct

	ButtonType
}

func NewPortfolioDiagramNodeButtonImpl(
	portfolioDiagramNode PortfolioDiagramNode,
	diagrammer *Diagrammer,
	treeNode *gongtree_models.Node,
	treeStage *gongtree_models.StageStruct,
	buttonType ButtonType,
) (buttonImpl *PortfolioDiagramNodeButtonImpl) {

	buttonImpl = new(PortfolioDiagramNodeButtonImpl)

	buttonImpl.portfolioDiagramNode = portfolioDiagramNode
	buttonImpl.diagrammer = diagrammer
	buttonImpl.treeNode = treeNode
	buttonImpl.treeStage = treeStage
	buttonImpl.ButtonType = buttonType

	return
}

func (buttonImpl *PortfolioDiagramNodeButtonImpl) ButtonUpdated(
	gongtreeStage *gongtree_models.StageStruct,
	stageButton, front *gongtree_models.Button) {

	switch buttonImpl.ButtonType {
	case DUPLICATE:
		parentTreeNode := buttonImpl.diagrammer.GetPortfiolioNodeFromTreeNode(buttonImpl.portfolioDiagramNode.GetParent())

		if childrenPortfolioNode := buttonImpl.portfolioDiagramNode.DuplicateDiagram(); childrenPortfolioNode != nil {
			childrenTreeNode := buttonImpl.diagrammer.portfolioNode2NodeTree(childrenPortfolioNode, buttonImpl.treeStage)
			parentTreeNode.Children = append(parentTreeNode.Children, childrenTreeNode)
			buttonImpl.diagrammer.AddPortfiolioNodeTreeNodeEntry(childrenPortfolioNode, childrenTreeNode)
			buttonImpl.diagrammer.generatePortfolioNodesStatusAndButtons()

			buttonImpl.treeStage.Commit()
		}
	case EDIT_CANCEL:
		map_ModelNode_Shape := buttonImpl.portfolioDiagramNode.CancelEdit()

		buttonImpl.diagrammer.generatePortfolioNodesStatusAndButtons()
		buttonImpl.diagrammer.computeModelNodeStatus(map_ModelNode_Shape)
		buttonImpl.treeStage.Commit()
	case EDIT:
		map_ModelNode_Shape := buttonImpl.portfolioDiagramNode.DrawDiagram()

		buttonImpl.diagrammer.generatePortfolioNodesStatusAndButtons()
		buttonImpl.diagrammer.computeModelNodeStatus(map_ModelNode_Shape)
		buttonImpl.treeStage.Commit()
	case REMOVE:
		parentTreeNode := buttonImpl.diagrammer.GetPortfiolioNodeFromTreeNode(buttonImpl.portfolioDiagramNode.GetParent())

		buttonImpl.portfolioDiagramNode.DeleteDiagram()

		// remove the node from the parent
		buttonImpl.portfolioDiagramNode.GetParent().RemoveChildren(buttonImpl.portfolioDiagramNode)

		idx := slices.Index(parentTreeNode.Children, buttonImpl.treeNode)
		parentTreeNode.Children = slices.Delete(parentTreeNode.Children, idx, idx+1)

		buttonImpl.diagrammer.RemovePortfiolioNodeTreeNodeEntry(buttonImpl.portfolioDiagramNode)
		buttonImpl.treeStage.Commit()

		buttonImpl.diagrammer.generatePortfolioNodesStatusAndButtons()

		// generate empty model tree
		map_ModelElementNode_Shape := make(map[ModelElementNode]Shape)
		buttonImpl.diagrammer.computeModelNodeStatus(map_ModelElementNode_Shape)
		buttonImpl.treeStage.Commit()
	case RENAME_CANCEL:
		buttonImpl.portfolioDiagramNode.SetIsInRenameMode(false)
		buttonImpl.treeNode.IsInEditMode = false

		buttonImpl.diagrammer.generatePortfolioNodesStatusAndButtons()
		buttonImpl.treeStage.Commit()
	case RENAME:
		buttonImpl.portfolioDiagramNode.SetIsInRenameMode(true)
		buttonImpl.treeNode.IsInEditMode = true

		buttonImpl.diagrammer.generatePortfolioNodesStatusAndButtons()
		buttonImpl.treeStage.Commit()
	case SAVE:
		map_ModelNode_Shape := buttonImpl.portfolioDiagramNode.SaveDiagram()

		buttonImpl.diagrammer.generatePortfolioNodesStatusAndButtons()
		buttonImpl.diagrammer.computeModelNodeStatus(map_ModelNode_Shape)
		buttonImpl.treeStage.Commit()

	}

}
