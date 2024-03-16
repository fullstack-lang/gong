package adapter

import (
	"github.com/fullstack-lang/gongdoc/go/diagrammer"
)

type ModelAdapter struct {
	portfolioAdapter *PortfolioAdapter
	rootNodes        []diagrammer.ModelNode
}

var _ diagrammer.Model = &ModelAdapter{}

func NewModelAdapter(portfolioAdapter *PortfolioAdapter) (adapter *ModelAdapter) {
	adapter = new(ModelAdapter)
	adapter.portfolioAdapter = portfolioAdapter
	return
}

// GetRootNodes implements bridge.Model.
func (modelAdapter *ModelAdapter) GenerateProgeny() []diagrammer.ModelNode {
	gongStructCategoryNode := NewGongStructCategoryNode(modelAdapter.portfolioAdapter, "gongstructs")
	gongStructCategoryNode.GenerateProgeny()
	modelAdapter.rootNodes = append(modelAdapter.rootNodes, gongStructCategoryNode)

	gongEnumCategoryNode := NewGongEnumCategoryNode(modelAdapter.portfolioAdapter, "gongenums")
	gongEnumCategoryNode.GenerateProgeny()
	modelAdapter.rootNodes = append(modelAdapter.rootNodes, gongEnumCategoryNode)

	gongNoteCategoryNode := NewGongNoteCategoryNode(modelAdapter.portfolioAdapter, "gongnotes")
	gongNoteCategoryNode.GenerateProgeny()
	modelAdapter.rootNodes = append(modelAdapter.rootNodes, gongNoteCategoryNode)

	return modelAdapter.rootNodes
}

// GetChildren implements diagrammer.Model.
func (modelAdapter *ModelAdapter) GetChildren() []diagrammer.ModelNode {
	return modelAdapter.rootNodes
}
