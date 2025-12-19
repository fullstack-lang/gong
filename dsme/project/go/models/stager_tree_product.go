package models

import (
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"

	"github.com/fullstack-lang/gong/lib/tree/go/buttons"
)

func (stager *Stager) updateProductTreeStage() {

	stager.treeProductsStage.Reset()

	root := stager.root

	treeInstance := &tree.Tree{Name: "PBS"}

	allProjectsNode := &tree.Node{
		Name:       "Projects",
		IsExpanded: true,
	}
	treeInstance.RootNodes = append(treeInstance.RootNodes, allProjectsNode)

	for _, project := range root.Projects {
		projectNode := &tree.Node{
			Name:            project.Name,
			IsExpanded:      project.IsExpanded,
			IsNodeClickable: true,
		}
		treeInstance.RootNodes = append(treeInstance.RootNodes, projectNode)
		projectNode.Impl = &NodeProxy[*Project]{
			stager:   stager,
			node:     projectNode,
			instance: project,
		}

		addAddItemButton(stager, projectNode, &project.RootProducts,
			func(items *[]*Product, item *Product) {
				*items = append(*items, item)
			})

		for _, product := range project.RootProducts {
			stager.generateTreeOfProduct(product, projectNode)
		}
	}

	if len(root.OrphanedProducts) > 0 {
		orphansProductNode := &tree.Node{Name: "Orphans Products", IsExpanded: true}
		treeInstance.RootNodes = append(treeInstance.RootNodes, orphansProductNode)
		for _, product := range root.OrphanedProducts {
			stager.generateTreeOfProduct(product, orphansProductNode)
		}
	}

	tree.StageBranch(stager.treeProductsStage, treeInstance)

	stager.treeProductsStage.Commit()
}

func (stager *Stager) generateTreeOfProduct(product *Product, parentNode *tree.Node) {

	productNode := &tree.Node{
		Name:            product.Name,
		IsExpanded:      true,
		IsNodeClickable: true,
	}
	parentNode.Children = append(parentNode.Children, productNode)

	productNode.Impl = &NodeProxy[*Product]{
		stager:   stager,
		node:     productNode,
		instance: product,
	}

	addAddItemButton(stager, productNode, &product.SubProducts,
		func(items *[]*Product, item *Product) {
			*items = append(*items, item)
		})

	for _, product := range product.SubProducts {
		stager.generateTreeOfProduct(product, productNode)
	}
}

func addAddItemButton[T Gongstruct, PT interface {
	*T
	GongstructIF
}](stager *Stager, productNode *tree.Node, items *[]*T, appendItem func(items *[]*T, item *T)) {
	addButton := &tree.Button{
		Name:            "Product" + " " + string(buttons.BUTTON_add),
		Icon:            string(buttons.BUTTON_add),
		ToolTipText:     "Add a product to \"" + productNode.Name + "\"",
		HasToolTip:      true,
		ToolTipPosition: tree.Right,
	}
	productNode.Buttons = append(productNode.Buttons, addButton)
	addButton.Impl = &tree.FunctionalButtonProxy{
		OnUpdated: func(stage *tree.Stage, button *tree.Button, updatedButton *tree.Button) {
			item := PT(new(T))
			appendItem(items, item)

			stager.stage.Commit()
		},
	}
}

type NodeProxy[T NodeType] struct {
	stager   *Stager
	node     *tree.Node
	instance T
}

// OnAfterUpdate implements models.NodeImplInterface.
func (p *NodeProxy[T]) OnAfterUpdate(stage *tree.Stage, stagedNode *tree.Node, frontNode *tree.Node) {

	if frontNode.IsExpanded != stagedNode.IsExpanded {
		p.instance.SetIsExpanded(!p.instance.GetIsExpanded())
		return
	}

	p.stager.probeForm.FillUpFormFromGongstruct(p.instance, GetPointerToGongstructName[T]())

	p.stager.stage.Commit()
}
