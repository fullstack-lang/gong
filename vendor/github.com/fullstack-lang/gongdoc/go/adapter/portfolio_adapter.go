package adapter

import (
	gong_models "github.com/fullstack-lang/gong/go/models"

	gongsvg_models "github.com/fullstack-lang/gongsvg/go/models"

	gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"

	"github.com/fullstack-lang/gongdoc/go/diagrammer"
)

type PortfolioAdapter struct {
	gongStage    *gong_models.StageStruct
	gongdocStage *gongdoc_models.StageStruct
	gongsvgStage *gongsvg_models.StageStruct
	diagrammer   *diagrammer.Diagrammer

	rootNodes []diagrammer.PortfolioNode

	selectedPortfolioDiagramNode diagrammer.PortfolioDiagramNode
}

// IsInDrawingMode implements diagrammer.Portfolio.
func (portfolioAdapter *PortfolioAdapter) IsInDrawingMode() bool {
	selectedClassdiagram := portfolioAdapter.getDiagramPackage().SelectedClassdiagram
	if selectedClassdiagram == nil {
		return false
	}

	return selectedClassdiagram.IsInDrawMode
}

var _ diagrammer.Portfolio = &PortfolioAdapter{}

func NewPortfolioAdapter(
	gongStage *gong_models.StageStruct,
	gongdocStage *gongdoc_models.StageStruct,
	svgStage *gongsvg_models.StageStruct,
) (portfolioAdapter *PortfolioAdapter) {
	portfolioAdapter = new(PortfolioAdapter)

	portfolioAdapter.gongStage = gongStage
	portfolioAdapter.gongdocStage = gongdocStage
	portfolioAdapter.gongsvgStage = svgStage

	return
}

func (portfolioAdapter *PortfolioAdapter) SetDiagrammer(diagrammer *diagrammer.Diagrammer) {
	portfolioAdapter.diagrammer = diagrammer
}

// IsInSelectionMode implements diagrammer.Portfolio.
func (*PortfolioAdapter) IsInSelectionMode() bool {
	return true // temp, will be false when a diagram is selected
}

// GetSelectedPortfolioDiagramNode implements diagrammer.Portfolio.
func (portfolioAdapter *PortfolioAdapter) GetSelectedPortfolioDiagramNode() (portfolioDiagramNode diagrammer.PortfolioDiagramNode) {

	portfolioDiagramNode = portfolioAdapter.selectedPortfolioDiagramNode

	return
}

// GetRootNodes implements bridge.Portfolio.
func (portfolioAdapter *PortfolioAdapter) GeneratesProgeny() []diagrammer.PortfolioNode {

	classDiagramCategoryNode := NewClassDiagramCategoryNode(portfolioAdapter, "class diagrams")
	portfolioAdapter.rootNodes = append(portfolioAdapter.rootNodes, classDiagramCategoryNode)
	classDiagramCategoryNode.GeneratesProgeny()

	return portfolioAdapter.rootNodes
}

func (portfolioAdapter *PortfolioAdapter) GetChildren() []diagrammer.PortfolioNode {

	return portfolioAdapter.rootNodes
}

func (portfolioAdapter *PortfolioAdapter) getDiagramPackage() *gongdoc_models.DiagramPackage {
	gongdocStage := portfolioAdapter.gongdocStage
	var diagramPackage *gongdoc_models.DiagramPackage
	for diagramPackage_ := range *gongdoc_models.GetGongstructInstancesSet[gongdoc_models.DiagramPackage](gongdocStage) {
		diagramPackage = diagramPackage_
	}
	return diagramPackage
}

// AddElementToDiagram implements diagrammer.Portfolio.
func (portfolioAdapter *PortfolioAdapter) AddElementToDiagram(modelElementNode diagrammer.ModelElementNode) (
	setOfModelElementNode map[diagrammer.ModelElementNode]diagrammer.Shape) {

	modelElementNode.AddToDiagram()

	portfolioDiagramNode := portfolioAdapter.GetSelectedPortfolioDiagramNode()
	gongStage := portfolioAdapter.gongStage
	if classDiagramNode, ok := portfolioDiagramNode.(*ClassDiagramNode); ok {
		setOfModelElementNode = classDiagramNode.getSetOfModelElementNodesInDiagram(gongStage, classDiagramNode.classdiagram)
	}

	return
}

func (portfolioAdapter *PortfolioAdapter) RemoveElementFromDiagram(modelElementNode diagrammer.ModelElementNode) (
	setOfModelElementNode map[diagrammer.ModelElementNode]diagrammer.Shape) {

	modelElementNode.RemoveFromDiagram()

	portfolioDiagramNode := portfolioAdapter.GetSelectedPortfolioDiagramNode()
	gongStage := portfolioAdapter.gongStage

	if classDiagramNode, ok := portfolioDiagramNode.(*ClassDiagramNode); ok {
		setOfModelElementNode = classDiagramNode.getSetOfModelElementNodesInDiagram(gongStage, classDiagramNode.classdiagram)
	}
	return
}

func (portfolioAdapter *PortfolioAdapter) setSelectedClassdiagramNode(classDiagramNode *ClassDiagramNode) {
	portfolioAdapter.selectedPortfolioDiagramNode = classDiagramNode
}
