package adapter

import (
	"log"

	gong_models "github.com/fullstack-lang/gong/go/models"
	"github.com/fullstack-lang/gongdoc/go/diagrammer"
	gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"
)

type GongEnumNode struct {
	ElementNodeBase
	gongEnum *gong_models.GongEnum
}

// RemoveFromDiagram implements diagrammer.ModelElementNode.
func (gongEnumNode *GongEnumNode) RemoveFromDiagram() {
	diagramPackage := gongEnumNode.portfolioAdapter.getDiagramPackage()

	var gongEnumShape *gongdoc_models.GongEnumShape
	shape, ok := gongEnumNode.portfolioAdapter.diagrammer.GetElementNodeDisplayed(gongEnumNode)

	if !ok {
		log.Fatalln("unknown gongenum shape")
		return
	}
	if gongEnumShape, ok = shape.(*gongdoc_models.GongEnumShape); !ok {
		log.Fatalln("not a gongenum shape")
		return
	}

	diagramPackage.SelectedClassdiagram.RemoveGongEnumShape(
		gongEnumNode.portfolioAdapter.gongdocStage,
		gongdoc_models.IdentifierToGongObjectName(gongEnumShape.Identifier))
}

// AddToDiagram implements diagrammer.ElementNode.
func (gongEnumNode *GongEnumNode) AddToDiagram() {
	diagramPackage := gongEnumNode.portfolioAdapter.getDiagramPackage()

	diagramPackage.SelectedClassdiagram.AddGongEnumShape(
		gongEnumNode.portfolioAdapter.gongdocStage, diagramPackage, gongEnumNode.GetName())
}

var _ diagrammer.ModelElementNode = &GongEnumNode{}

func NewGongEnumNode(
	portfolioAdapter *PortfolioAdapter,
	gongEnum *gong_models.GongEnum) (gongEnumNode *GongEnumNode) {
	gongEnumNode = &GongEnumNode{ElementNodeBase: ElementNodeBase{portfolioAdapter: portfolioAdapter}}

	gongEnumNode.gongEnum = gongEnum
	return
}

// GenerateProgeny implements diagrammer.Node.
func (gongEnumNode *GongEnumNode) GenerateProgeny() []diagrammer.ModelNode {

	for _, gongEnumValue := range gongEnumNode.gongEnum.GongEnumValues {
		enumValueNode := NewEnumValueNode(gongEnumNode.portfolioAdapter, gongEnumNode, gongEnumValue)
		gongEnumNode.children = append(gongEnumNode.children, enumValueNode)
	}

	return gongEnumNode.children
}

// GetName implements diagrammer.Node.
func (gongEnumNode *GongEnumNode) GetName() string {
	return gongEnumNode.gongEnum.GetName()
}

func (gongEnumNode *GongEnumNode) GetElement() any {
	return gongEnumNode.gongEnum
}
