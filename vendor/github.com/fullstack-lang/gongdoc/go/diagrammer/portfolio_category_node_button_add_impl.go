package diagrammer

import (
	gongtree_models "github.com/fullstack-lang/gongtree/go/models"
)

type PortfolioCategoryNodeButtonAddImpl struct {
	portfolioCategoryNode PortfolioCategoryNode

	diagrammer *Diagrammer

	treeNode *gongtree_models.Node

	treeStage *gongtree_models.StageStruct
}

func NewDiagramButtonAddImpl(
	portfolioCategoryNode PortfolioCategoryNode,
	diagrammer *Diagrammer,
	treeNode *gongtree_models.Node,
	treeStage *gongtree_models.StageStruct,
) (portfolioCategoryNodeButtonAddImpl *PortfolioCategoryNodeButtonAddImpl) {

	portfolioCategoryNodeButtonAddImpl = new(PortfolioCategoryNodeButtonAddImpl)

	portfolioCategoryNodeButtonAddImpl.portfolioCategoryNode = portfolioCategoryNode
	portfolioCategoryNodeButtonAddImpl.diagrammer = diagrammer
	portfolioCategoryNodeButtonAddImpl.treeNode = treeNode
	portfolioCategoryNodeButtonAddImpl.treeStage = treeStage

	return
}

func (buttonImpl *PortfolioCategoryNodeButtonAddImpl) ButtonUpdated(
	gongtreeStage *gongtree_models.StageStruct,
	stageButton, front *gongtree_models.Button) {

	if childrenPortfolioNode := buttonImpl.portfolioCategoryNode.AddDiagram(); childrenPortfolioNode != nil {
		childrenTreeNode := buttonImpl.diagrammer.portfolioNode2NodeTree(childrenPortfolioNode, buttonImpl.treeStage)
		buttonImpl.treeNode.Children = append(buttonImpl.treeNode.Children, childrenTreeNode)
		buttonImpl.diagrammer.AddPortfiolioNodeTreeNodeEntry(childrenPortfolioNode, childrenTreeNode)
		buttonImpl.diagrammer.generatePortfolioNodesStatusAndButtons()
		buttonImpl.treeStage.Commit()
	}
}
