package models

import (
	"fmt"

	tree "github.com/fullstack-lang/gong/lib/tree/go/models"

	"github.com/fullstack-lang/gong/lib/tree/go/buttons"
)

func (stager *Stager) tree() {

	stager.treeStage.Reset()

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

		pbsNode := &tree.Node{
			Name:            "PBS",
			IsExpanded:      project.IsPBSNodeExpanded,
			IsNodeClickable: true,
		}
		projectNode.Children = append(projectNode.Children, pbsNode)
		pbsNode.Impl = &expandableNodeProxy{
			node:           pbsNode,
			stager:         stager,
			isNodeExpanded: &project.IsPBSNodeExpanded,
		}

		addAddItemButton(stager, pbsNode, &project.RootProducts,
			func(items *[]*Product, item *Product) {
				*items = append(*items, item)
			})

		for _, product := range project.RootProducts {
			stager.generateTreeOfProduct(product, pbsNode)
		}

		wbsNode := &tree.Node{
			Name:            "WBS",
			IsExpanded:      project.IsWBSNodeExpanded,
			IsNodeClickable: true,
		}
		projectNode.Children = append(projectNode.Children, wbsNode)
		wbsNode.Impl = &expandableNodeProxy{
			node:           wbsNode,
			stager:         stager,
			isNodeExpanded: &project.IsWBSNodeExpanded,
		}

		addAddItemButton(stager, wbsNode, &project.RootTasks,
			func(items *[]*Task, item *Task) {
				*items = append(*items, item)
			})

		for _, task := range project.RootTasks {
			stager.generateTreeOfTask(task, wbsNode)
		}
	}

	if len(root.OrphanedProducts) > 0 {
		orphansProductNode := &tree.Node{Name: "Orphans Products", IsExpanded: true}
		treeInstance.RootNodes = append(treeInstance.RootNodes, orphansProductNode)
		for _, product := range root.OrphanedProducts {
			stager.generateTreeOfProduct(product, orphansProductNode)
		}
	}

	if len(root.OrphanedTasks) > 0 {
		orphansTaskNode := &tree.Node{Name: "Orphans Tasks", IsExpanded: true}
		treeInstance.RootNodes = append(treeInstance.RootNodes, orphansTaskNode)
		for _, task := range root.OrphanedTasks {
			stager.generateTreeOfTask(task, orphansTaskNode)
		}
	}

	tree.StageBranch(stager.treeStage, treeInstance)

	stager.treeStage.Commit()
}

func (stager *Stager) generateTreeOfProduct(product *Product, parentNode *tree.Node) {

	productNode := &tree.Node{
		Name:            product.ComputedPrefix + " " + product.Name,
		IsExpanded:      product.IsExpanded,
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

func (stager *Stager) generateTreeOfTask(task *Task, parentNode *tree.Node) {

	taskNode := &tree.Node{
		Name:            task.ComputedPrefix + " " + task.Name,
		IsExpanded:      task.IsExpanded,
		IsNodeClickable: true,
	}
	parentNode.Children = append(parentNode.Children, taskNode)

	taskNode.Impl = &NodeProxy[*Task]{
		stager:   stager,
		node:     taskNode,
		instance: task,
	}

	addAddItemButton(stager, taskNode, &task.SubTasks,
		func(items *[]*Task, item *Task) {
			*items = append(*items, item)
		})

	for _, task := range task.SubTasks {
		stager.generateTreeOfTask(task, taskNode)
	}

	if len(task.Inputs) > 0 {
		inputProductsNode := &tree.Node{
			Name:                 fmt.Sprintf("(%d)", len(task.Inputs)),
			IsExpanded:           task.IsInputsNodeExpanded,
			IsNodeClickable:      true,
			IsWithPreceedingIcon: true,
			PreceedingIcon:       string(buttons.BUTTON_input),
		}
		taskNode.Children = append(taskNode.Children, inputProductsNode)
		inputProductsNode.Impl = &expandableNodeProxy{
			node:           inputProductsNode,
			stager:         stager,
			isNodeExpanded: &task.IsInputsNodeExpanded,
		}

		for _, product := range task.Inputs {
			inputProductNode := &tree.Node{
				Name:            product.GetName(),
				IsExpanded:      true,
				IsNodeClickable: true,
			}
			inputProductsNode.Children = append(inputProductsNode.Children, inputProductNode)
			inputProductNode.Impl = &NodeProxy[*Product]{
				stager:   stager,
				node:     inputProductNode,
				instance: product,
			}
		}
	}

	if len(task.Outputs) > 0 {
		outputProductsNode := &tree.Node{
			Name:                 fmt.Sprintf("(%d)", len(task.Outputs)),
			IsExpanded:           task.IsOutputsNodeExpanded,
			IsNodeClickable:      true,
			IsWithPreceedingIcon: true,
			PreceedingIcon:       string(buttons.BUTTON_output),
		}
		taskNode.Children = append(taskNode.Children, outputProductsNode)
		outputProductsNode.Impl = &expandableNodeProxy{
			node:           outputProductsNode,
			stager:         stager,
			isNodeExpanded: &task.IsOutputsNodeExpanded,
		}

		for _, product := range task.Outputs {
			outputProductNode := &tree.Node{
				Name:            product.GetName(),
				IsExpanded:      true,
				IsNodeClickable: true,
			}
			outputProductsNode.Children = append(outputProductsNode.Children, outputProductNode)
			outputProductNode.Impl = &NodeProxy[*Product]{
				stager:   stager,
				node:     outputProductNode,
				instance: product,
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

type NodeProxy[T ProjectElementType] struct {
	stager   *Stager
	node     *tree.Node
	instance T
}

// OnAfterUpdate implements models.NodeImplInterface.
func (p *NodeProxy[T]) OnAfterUpdate(stage *tree.Stage, stagedNode *tree.Node, frontNode *tree.Node) {

	if frontNode.IsExpanded != stagedNode.IsExpanded {
		stagedNode.IsExpanded = frontNode.IsExpanded
		p.instance.SetIsExpanded(!p.instance.GetIsExpanded())
		return
	}

	p.stager.probeForm.FillUpFormFromGongstruct(p.instance, GetPointerToGongstructName[T]())

	p.stager.stage.Commit()
}
