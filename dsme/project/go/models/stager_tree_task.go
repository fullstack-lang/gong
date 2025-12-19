package models

import (
	"fmt"

	"github.com/fullstack-lang/gong/lib/tree/go/buttons"

	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) updateTaskTreeStage() {

	stager.treeTasksStage.Reset()

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

		addAddItemButton(stager, projectNode, &project.RootTasks,
			func(items *[]*Task, item *Task) {
				*items = append(*items, item)
			})

		for _, task := range project.RootTasks {
			stager.generateTreeOfTask(task, projectNode)
		}
	}

	if len(root.OrphanedTasks) > 0 {
		orphansTaskNode := &tree.Node{Name: "Orphans Tasks", IsExpanded: true}
		treeInstance.RootNodes = append(treeInstance.RootNodes, orphansTaskNode)
		for _, task := range root.OrphanedTasks {
			stager.generateTreeOfTask(task, orphansTaskNode)
		}
	}

	tree.StageBranch(stager.treeTasksStage, treeInstance)

	stager.treeTasksStage.Commit()
}

func (stager *Stager) generateTreeOfTask(task *Task, parentNode *tree.Node) {

	taskNode := &tree.Node{
		Name:            task.ComputedPrefix + " " + task.Name,
		IsExpanded:      true,
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

	if len(task.InputProducts) > 0 {
		inputProductsNode := &tree.Node{
			Name:                 fmt.Sprintf("(%d)", len(task.InputProducts)),
			IsExpanded:           task.IsInputProducsNodeExpanded,
			IsNodeClickable:      true,
			IsWithPreceedingIcon: true,
			PreceedingIcon:       string(buttons.BUTTON_input),
		}
		taskNode.Children = append(taskNode.Children, inputProductsNode)
		inputProductsNode.Impl = &expandableNodeProxy{
			node:           inputProductsNode,
			stager:         stager,
			isNodeExpanded: &task.IsInputProducsNodeExpanded,
		}

		for _, product := range task.InputProducts {
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

	if len(task.OutputProducts) > 0 {
		outputProductsNode := &tree.Node{
			Name:                 fmt.Sprintf("(%d)", len(task.OutputProducts)),
			IsExpanded:           task.IsOutputProducsNodeExpanded,
			IsNodeClickable:      true,
			IsWithPreceedingIcon: true,
			PreceedingIcon:       string(buttons.BUTTON_output),
		}
		taskNode.Children = append(taskNode.Children, outputProductsNode)
		outputProductsNode.Impl = &expandableNodeProxy{
			node:           outputProductsNode,
			stager:         stager,
			isNodeExpanded: &task.IsOutputProducsNodeExpanded,
		}

		for _, product := range task.OutputProducts {
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
