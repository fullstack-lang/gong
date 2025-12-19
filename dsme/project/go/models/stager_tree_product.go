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

		stager.addAddProductButton(projectNode, &project.RootProducts)

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

	productNode.Impl = &ProductNodeProxy{
		stager:  stager,
		node:    productNode,
		product: product,
	}

	stager.addAddProductButton(productNode, &product.SubProducts)

	for _, product := range product.SubProducts {
		stager.generateTreeOfProduct(product, productNode)
	}
}

func (stager *Stager) addAddProductButton(productNode *tree.Node, products *[]*Product) {
	addButton := &tree.Button{
		Name:            "Product" + " " + string(buttons.BUTTON_add),
		Icon:            string(buttons.BUTTON_add),
		ToolTipText:     "Add a product to \"" + productNode.Name + "\"",
		HasToolTip:      true,
		ToolTipPosition: tree.Right,
	}
	productNode.Buttons = append(productNode.Buttons, addButton)
	addButton.Impl = &AddProductButtonNodeProxy{
		stager:   stager,
		products: products,
	}
}

type AddProductButtonNodeProxy struct {
	stager   *Stager
	products *[]*Product
}

// ButtonUpdated implements models.ButtonImplInterface.
func (p *AddProductButtonNodeProxy) ButtonUpdated(stage *tree.Stage, button *tree.Button, updatedButton *tree.Button) {

	product := (&Product{
		Name: "New Product",
	}).Stage(p.stager.stage)

	*p.products = append(*p.products, product)

	p.stager.stage.Commit()
}

type ProjectNodeProxy struct {
	stager  *Stager
	node    *tree.Node
	project *Project
}

// OnAfterUpdate implements models.NodeImplInterface.
func (p *ProjectNodeProxy) OnAfterUpdate(stage *tree.Stage, stagedNode *tree.Node, frontNode *tree.Node) {

	if frontNode.IsExpanded != stagedNode.IsExpanded {
		p.project.IsExpanded = !p.project.IsExpanded
		return
	}

	p.stager.probeForm.FillUpFormFromGongstruct(p.project, GetPointerToGongstructName[*Product]())

	p.stager.stage.Commit()
}

type ProductNodeProxy struct {
	stager  *Stager
	node    *tree.Node
	product *Product
}

// OnAfterUpdate implements models.NodeImplInterface.
func (p *ProductNodeProxy) OnAfterUpdate(stage *tree.Stage, stagedNode *tree.Node, frontNode *tree.Node) {

	if frontNode.IsExpanded != stagedNode.IsExpanded {
		p.product.IsExpanded = !p.product.IsExpanded
		return
	}

	p.stager.probeForm.FillUpFormFromGongstruct(p.product, GetPointerToGongstructName[*Product]())

	p.stager.stage.Commit()
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
