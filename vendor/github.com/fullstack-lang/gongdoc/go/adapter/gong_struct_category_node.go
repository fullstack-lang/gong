package adapter

import (
	"cmp"
	"slices"

	gong_models "github.com/fullstack-lang/gong/go/models"
	"github.com/fullstack-lang/gongdoc/go/diagrammer"
)

type GongStructCategoryNode struct {
	ModelCategoryNodeBase

	map_Gongtruct_GongstructNode map[*gong_models.GongStruct]*GongStructNode
}

func NewGongStructCategoryNode(portfolioAdapter *PortfolioAdapter, name string) (
	gongstructCategoryNode *GongStructCategoryNode) {

	gongstructCategoryNode = &GongStructCategoryNode{
		ModelCategoryNodeBase:        ModelCategoryNodeBase{portfolioAdapter: portfolioAdapter, Name: name},
		map_Gongtruct_GongstructNode: make(map[*gong_models.GongStruct]*GongStructNode),
	}

	return
}

// GenerateProgeny implements diagrammer.Node.
func (categoryNode *GongStructCategoryNode) GenerateProgeny() []diagrammer.ModelNode {

	for gongStruct := range *gong_models.GetGongstructInstancesSet[gong_models.GongStruct](categoryNode.portfolioAdapter.gongStage) {

		gongStructNode := NewGongStructNode(categoryNode.portfolioAdapter, gongStruct)
		gongStructNode.GenerateProgeny()
		categoryNode.children = append(categoryNode.children, gongStructNode)
		categoryNode.map_Gongtruct_GongstructNode[gongStruct] = gongStructNode
	}

	slices.SortFunc(categoryNode.children, func(a, b diagrammer.ModelNode) int {
		return cmp.Compare(a.GetName(), b.GetName())
	})

	return categoryNode.children
}
