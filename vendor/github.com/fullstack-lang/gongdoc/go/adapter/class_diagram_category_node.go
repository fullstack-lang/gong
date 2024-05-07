package adapter

import (
	"cmp"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"slices"

	gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"

	"github.com/fullstack-lang/gongdoc/go/diagrammer"
)

type ClassDiagramCategoryNode struct {
	portfolioAdapter *PortfolioAdapter
	Name             string
	children         []diagrammer.PortfolioNode
}

// static check that it meets the intended interface
var _ diagrammer.PortfolioCategoryNode = &(ClassDiagramCategoryNode{})

func NewClassDiagramCategoryNode(
	portfolioAdapter *PortfolioAdapter,
	name string,
) (categoryNode *ClassDiagramCategoryNode) {
	categoryNode = new(ClassDiagramCategoryNode)

	categoryNode.portfolioAdapter = portfolioAdapter
	categoryNode.Name = name

	return
}

// GetChildren implements diagrammer.PortfolioCategoryNode.
func (classDiagramCategoryNode *ClassDiagramCategoryNode) GetChildren() []diagrammer.PortfolioNode {
	return classDiagramCategoryNode.children
}

// RemoveChildren implements diagrammer.PortfolioCategoryNode.
func (classDiagramCategoryNode *ClassDiagramCategoryNode) RemoveChildren(portfolioNode diagrammer.PortfolioNode) {
	idx := diagrammer.IndexOf(classDiagramCategoryNode.children, portfolioNode)
	classDiagramCategoryNode.children = diagrammer.Remove(classDiagramCategoryNode.children, idx)
}

// AppendChildren implements diagrammer.PortfolioCategoryNode.
func (classDiagramCategoryNode *ClassDiagramCategoryNode) AppendChildren(children diagrammer.PortfolioNode) {
	classDiagramCategoryNode.children = append(classDiagramCategoryNode.children, children)
}

// GetParent implements diagrammer.PortfolioCategoryNode.
func (classDiagramCategoryNode *ClassDiagramCategoryNode) GetParent() diagrammer.PortfolioNode {
	return nil
}

func (ClassDiagramCategoryNode *ClassDiagramCategoryNode) IsExpanded() bool {
	return true
}

func (categoryNode *ClassDiagramCategoryNode) GeneratesProgeny() []diagrammer.PortfolioNode {

	gongdocStage := categoryNode.portfolioAdapter.gongdocStage
	for classDiagram := range *gongdoc_models.GetGongstructInstancesSet[gongdoc_models.Classdiagram](gongdocStage) {

		classDiagramNode := NewClassDiagramNode(
			categoryNode.portfolioAdapter,
			categoryNode,
			classDiagram)
		categoryNode.AppendChildren(classDiagramNode)
		classDiagramNode.GeneratesProgeny()
	}

	slices.SortFunc(categoryNode.children, func(a, b diagrammer.PortfolioNode) int {
		return cmp.Compare(a.GetName(), b.GetName())
	})

	return categoryNode.children
}

// GetName implements diagrammer.Node.
func (classDiagramCategoryNode *ClassDiagramCategoryNode) GetName() string {
	return classDiagramCategoryNode.Name
}

// HasAddDiagramButton
func (classDiagramCategoryNode *ClassDiagramCategoryNode) HasAddDiagramButton() bool {

	diagramPackage := classDiagramCategoryNode.portfolioAdapter.getDiagramPackage()

	return diagramPackage.IsEditable
}

// AddDiagram implements diagrammer.Portfolio.
func (classDiagramCategoryNode *ClassDiagramCategoryNode) AddDiagram() diagrammer.PortfolioDiagramNode {

	gongdocStage := classDiagramCategoryNode.portfolioAdapter.gongdocStage
	diagramPackage := classDiagramCategoryNode.portfolioAdapter.getDiagramPackage()

	// check unicity of name, otherwise, add an index
	var hasNameCollision bool
	initialName := "Default"
	newClassdiagramName := initialName
	index := 0
	// loop until the name of the new diagram is not in collision with an existing
	// diagram name
	for index == 0 || hasNameCollision {
		index++
		hasNameCollision = false
		for classdiagram := range *gongdoc_models.GetGongstructInstancesSet[gongdoc_models.Classdiagram](gongdocStage) {
			if classdiagram.Name == newClassdiagramName {
				hasNameCollision = true
			}
		}
		if hasNameCollision {
			newClassdiagramName = initialName + fmt.Sprintf("_%d", index)
		}
	}

	newClassdiagram := (&gongdoc_models.Classdiagram{Name: newClassdiagramName}).Stage(gongdocStage)
	diagramPackage.Classdiagrams = append(diagramPackage.Classdiagrams, newClassdiagram)

	filepath := filepath.Join(
		filepath.Join(diagramPackage.AbsolutePathToDiagramPackage, "../diagrams"),
		newClassdiagram.Name) + ".go"
	file, err := os.Create(filepath)
	if err != nil {
		log.Fatal("Cannot open diagram file" + err.Error())
	}
	defer file.Close()

	// the gongstage has now the new class diagram within
	gongdocStage.Commit()

	// now save the diagram into a file
	// to do that, one need to empty the stage of all diagramms but the
	// new classdiagram
	gongdocStage.Checkout()
	gongdocStage.Unstage()
	gongdoc_models.StageBranch(gongdocStage, newClassdiagram)

	gongdoc_models.SetupMapDocLinkRenaming(diagramPackage.ModelPkg.Stage_, gongdocStage)

	gongdocStage.Marshall(file, "github.com/fullstack-lang/gongdoc/go/models", "diagrams")

	// restore the original stage
	gongdocStage.Unstage()
	gongdocStage.Checkout()

	classDiagramNode := NewClassDiagramNode(classDiagramCategoryNode.portfolioAdapter, classDiagramCategoryNode, newClassdiagram)

	classDiagramCategoryNode.AppendChildren(classDiagramNode)

	return classDiagramNode
}
