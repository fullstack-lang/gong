package models

import (
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"

	"github.com/fullstack-lang/gong/lib/tree/go/buttons"
)

func (stager *Stager) updateProductTreeStage() {

	stager.treeProductsStage.Reset()

	root := stager.root

	treeInstance := new(tree.Tree).Stage(stager.treeProductsStage)

	for _, project := range root.Projects {

		nodeForAddButton := new(tree.Node).Stage(stager.treeProductsStage)
		nodeForAddButton.Name = project.Name
		treeInstance.RootNodes = append(treeInstance.RootNodes, nodeForAddButton)

		addButton := (&tree.Button{
			Name: "Product" + " " + string(buttons.BUTTON_add),
			Icon: string(buttons.BUTTON_add)}).Stage(stager.treeProductsStage)
		nodeForAddButton.Buttons = append(nodeForAddButton.Buttons, addButton)
		addButton.Impl = &ProductAddButtonProxy{
			stager:   stager,
			products: project.RootProducts,
		}

		for _, product := range project.RootProducts {
			nodeProduct := new(tree.Node).Stage(stager.treeProductsStage)
			nodeProduct.Name = product.Name
			nodeProduct.IsExpanded = true

			nodeProduct.Impl = &expandableNodeProxy{
				node:           nodeProduct,
				stager:         stager,
				isNodeExpanded: &nodeProduct.IsExpanded,
			}
			treeInstance.RootNodes = append(treeInstance.RootNodes, nodeProduct)
		}
	}

	tree.StageBranch(stager.treeProductsStage, treeInstance)

	stager.treeProductsStage.Commit()
}
