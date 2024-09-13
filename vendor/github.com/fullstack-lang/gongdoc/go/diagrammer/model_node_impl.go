package diagrammer

import gongtree_models "github.com/fullstack-lang/gongtree/go/models"

type ModelNodeImpl struct {
	diagrammer *Diagrammer
	modelNode  ModelNode
}

// OnAfterUpdate implements models.NodeImplInterface
func (modelNodeImpl *ModelNodeImpl) OnAfterUpdate(stage *gongtree_models.StageStruct, stagedNode *gongtree_models.Node, frontNode *gongtree_models.Node) {

	if frontNode.IsExpanded != stagedNode.IsExpanded {
		modelNodeImpl.modelNode.SetIsExpanded(!modelNodeImpl.modelNode.IsExpanded())

		// we need to update the stage node because when the tree stage is committed later
		// the expanded configuration of the node will be correct
		// if not, the front node will be in the configuration of the stage node
		// whatever action the end user performs
		stagedNode.IsExpanded = frontNode.IsExpanded
	}

	if !stagedNode.IsChecked && frontNode.IsChecked {

		if modelElementNode, ok := modelNodeImpl.modelNode.(ModelElementNode); ok {
			map_ModelNode_Shape := modelNodeImpl.diagrammer.portfolio.AddElementToDiagram(modelElementNode)

			stagedNode.IsChecked = frontNode.IsChecked
			modelNodeImpl.diagrammer.generatePortfolioNodesStatusAndButtons()
			modelNodeImpl.diagrammer.computeModelNodeStatus(map_ModelNode_Shape)
			modelNodeImpl.diagrammer.treeStage.Commit()
		}

	}

	if stagedNode.IsChecked && !frontNode.IsChecked {

		if modelElementNode, ok := modelNodeImpl.modelNode.(ModelElementNode); ok {
			map_ModelNode_Shape := modelNodeImpl.diagrammer.portfolio.RemoveElementFromDiagram(modelElementNode)

			stagedNode.IsChecked = frontNode.IsChecked
			modelNodeImpl.diagrammer.generatePortfolioNodesStatusAndButtons()
			modelNodeImpl.diagrammer.computeModelNodeStatus(map_ModelNode_Shape)
			modelNodeImpl.diagrammer.treeStage.Commit()
		}
	}

	if stagedNode.IsSecondCheckboxChecked && !frontNode.IsSecondCheckboxChecked {

		if modelElementNode, ok := modelNodeImpl.modelNode.(ModelElementNode); ok {
			_ = modelElementNode
		}

	}

	if !stagedNode.IsSecondCheckboxChecked && frontNode.IsSecondCheckboxChecked {

		if modelElementNode, ok := modelNodeImpl.modelNode.(ModelElementNode); ok {
			_ = modelElementNode

			map_ModelNode_Shape := modelNodeImpl.diagrammer.portfolio.AddElementToDiagram(
				modelElementNode)

			stagedNode.IsChecked = frontNode.IsChecked
			modelNodeImpl.diagrammer.generatePortfolioNodesStatusAndButtons()
			modelNodeImpl.diagrammer.computeModelNodeStatus(map_ModelNode_Shape)
			modelNodeImpl.diagrammer.treeStage.Commit()
		}

	}

}
