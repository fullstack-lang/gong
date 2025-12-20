package models

import (
	"fmt"

	tree "github.com/fullstack-lang/gong/lib/tree/go/models"

	"github.com/fullstack-lang/gong/lib/tree/go/buttons"
)

func (stager *Stager) updateProductTreeStage() {

	stager.treeProductsStage.Reset()

	root := stager.root

	treeInstance := &tree.Tree{Name: "PBS"}

	allProjectsNode := &tree.Node{
		Name:       "** Tree of Projects **",
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
		Name:            product.ComputedPrefix + " " + product.Name,
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

	if len(product.producers) > 0 {
		producersNode := &tree.Node{
			Name:                 fmt.Sprintf("(%d)", len(product.producers)),
			IsExpanded:           product.IsProducersNodeExpanded,
			IsNodeClickable:      true,
			IsWithPreceedingIcon: true,
			PreceedingIcon:       string(buttons.BUTTON_input),
		}
		productNode.Children = append(productNode.Children, producersNode)
		producersNode.Impl = &expandableNodeProxy{
			node:           producersNode,
			stager:         stager,
			isNodeExpanded: &product.IsProducersNodeExpanded,
		}

		for _, task := range product.producers {
			inputProductNode := &tree.Node{
				Name:            task.GetName(),
				IsExpanded:      true,
				IsNodeClickable: true,
			}
			producersNode.Children = append(producersNode.Children, inputProductNode)
			inputProductNode.Impl = &NodeProxy[*Task]{
				stager:   stager,
				node:     inputProductNode,
				instance: task,
			}
		}
	}

	if len(product.consumers) > 0 {
		consumersNode := &tree.Node{
			Name:                 fmt.Sprintf("(%d)", len(product.consumers)),
			IsExpanded:           product.IsConsumersNodeExpanded,
			IsNodeClickable:      true,
			IsWithPreceedingIcon: true,
			PreceedingIcon:       string(buttons.BUTTON_output),
		}
		productNode.Children = append(productNode.Children, consumersNode)
		consumersNode.Impl = &expandableNodeProxy{
			node:           consumersNode,
			stager:         stager,
			isNodeExpanded: &product.IsConsumersNodeExpanded,
		}

		for _, task := range product.consumers {
			outputTaskNode := &tree.Node{
				Name:            task.GetName(),
				IsExpanded:      true,
				IsNodeClickable: true,
			}
			consumersNode.Children = append(consumersNode.Children, outputTaskNode)
			outputTaskNode.Impl = &NodeProxy[*Task]{
				stager:   stager,
				node:     outputTaskNode,
				instance: task,
			}
		}
	}
}

func addAddItemButton[T Gongstruct, PT interface {
	*T
	GongstructIF
}](stager *Stager, node *tree.Node, items *[]*T, appendItem func(items *[]*T, item *T)) {

	var item PT
	addButton := &tree.Button{
		Name:            GetGongstructNameFromPointer(item) + " " + string(buttons.BUTTON_add),
		Icon:            string(buttons.BUTTON_add),
		ToolTipText:     "Add a " + GetGongstructNameFromPointer(item) + " to \"" + node.Name + "\"",
		HasToolTip:      true,
		ToolTipPosition: tree.Right,
	}
	node.Buttons = append(node.Buttons, addButton)
	addButton.Impl = &tree.FunctionalButtonProxy{
		OnUpdated: func(stage *tree.Stage, button *tree.Button, updatedButton *tree.Button) {
			item := PT(new(T))
			item.SetName("New" + GetGongstructNameFromPointer(item))
			item.StageVoid(stager.stage)
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
