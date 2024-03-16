package adapter

import (
	gong_models "github.com/fullstack-lang/gong/go/models"
	"github.com/fullstack-lang/gongdoc/go/diagrammer"
)

func NewGongStructNode(
	portfolioAdapter *PortfolioAdapter,
	gongStruct *gong_models.GongStruct) (gongStructNode *GongStructNode) {
	gongStructNode = &GongStructNode{
		ElementNodeBase: ElementNodeBase{portfolioAdapter: portfolioAdapter}}

	gongStructNode.gongStruct = gongStruct
	return
}

var _ diagrammer.ModelElementNode = &GongStructNode{}

type GongStructNode struct {
	ElementNodeBase
	gongStruct *gong_models.GongStruct
}

// RemoveFromDiagram implements diagrammer.ModelElementNode.
func (gongStructNode *GongStructNode) RemoveFromDiagram() {
	diagramPackage := gongStructNode.portfolioAdapter.getDiagramPackage()

	diagramPackage.SelectedClassdiagram.RemoveGongStructShape(
		gongStructNode.portfolioAdapter.gongdocStage, gongStructNode.gongStruct.Name)
}

// AddToDiagram implements diagrammer.ElementNode.
func (gongStructNode *GongStructNode) AddToDiagram() {
	diagramPackage := gongStructNode.portfolioAdapter.getDiagramPackage()

	diagramPackage.SelectedClassdiagram.AddGongStructShape(
		gongStructNode.portfolioAdapter.gongdocStage, diagramPackage, gongStructNode.gongStruct.Name)
}

// GenerateProgeny implements diagrammer.Node.
func (gongStructNode *GongStructNode) GenerateProgeny() []diagrammer.ModelNode {

	for _, field := range gongStructNode.gongStruct.Fields {
		fieldNode := NewFieldNode(gongStructNode.portfolioAdapter, gongStructNode, field)
		gongStructNode.children = append(gongStructNode.children, fieldNode)
	}

	return gongStructNode.children
}

// GetName implements diagrammer.Node.
func (gongStructNode *GongStructNode) GetName() string {
	return gongStructNode.gongStruct.GetName()
}

func (gongStructNode *GongStructNode) GetElement() any {
	return gongStructNode.gongStruct
}
