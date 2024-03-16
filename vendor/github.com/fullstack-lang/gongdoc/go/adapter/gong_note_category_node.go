package adapter

import (
	"cmp"
	"slices"

	gong_models "github.com/fullstack-lang/gong/go/models"
	"github.com/fullstack-lang/gongdoc/go/diagrammer"
)

type GongNoteCategoryNode struct {
	ModelCategoryNodeBase
}

func NewGongNoteCategoryNode(portfolioAdapter *PortfolioAdapter, name string) *GongNoteCategoryNode {
	return &GongNoteCategoryNode{ModelCategoryNodeBase: ModelCategoryNodeBase{portfolioAdapter: portfolioAdapter, Name: name}}
}

// GenerateProgeny implements diagrammer.Node.
func (categoryNode *GongNoteCategoryNode) GenerateProgeny() []diagrammer.ModelNode {

	for gongNote := range *gong_models.GetGongstructInstancesSet[gong_models.GongNote](categoryNode.portfolioAdapter.gongStage) {

		gongNoteNode := NewGongNoteNode(categoryNode.portfolioAdapter, gongNote)
		gongNoteNode.GenerateProgeny()
		categoryNode.children = append(categoryNode.children, gongNoteNode)
	}

	slices.SortFunc(categoryNode.children, func(a, b diagrammer.ModelNode) int {
		return cmp.Compare(a.GetName(), b.GetName())
	})

	return categoryNode.children
}
