package models

import (
	"fmt"
	"slices"

	"github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) treeWBS(task *Task, parentNode *tree.Node) {
	taskNode := &tree.Node{
		Name:            task.ComputedPrefix + " " + task.Name,
		IsExpanded:      task.IsExpanded,
		IsNodeClickable: true,
	}
	parentNode.Children = append(parentNode.Children, taskNode)

	taskNode.Impl = &tree.FunctionalNodeProxy{
		OnUpdate: OnUpdateAbstractElement(stager, task),
	}

	addAddItemButton(stager, taskNode, &task.SubTasks)

	for _, task := range task.SubTasks {
		stager.treeWBS(task, taskNode)
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

		inputProductsNode.Impl = &tree.FunctionalNodeProxy{
			OnUpdate: stager.OnUpdateExpansion(&task.IsInputsNodeExpanded),
		}

		for _, product := range task.Inputs {
			inputProductNode := &tree.Node{
				Name:            product.GetName(),
				IsExpanded:      true,
				IsNodeClickable: true,
			}
			inputProductsNode.Children = append(inputProductsNode.Children, inputProductNode)

			inputProductNode.Impl = &tree.FunctionalNodeProxy{
				OnUpdate: OnUpdateAbstractElement(stager, product),
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

		outputProductsNode.Impl = &tree.FunctionalNodeProxy{
			OnUpdate: stager.OnUpdateExpansion(&task.IsOutputsNodeExpanded),
		}

		for _, product := range task.Outputs {
			outputProductNode := &tree.Node{
				Name:            product.GetName(),
				IsExpanded:      true,
				IsNodeClickable: true,
			}
			outputProductsNode.Children = append(outputProductsNode.Children, outputProductNode)

			outputProductNode.Impl = &tree.FunctionalNodeProxy{
				OnUpdate: OnUpdateAbstractElement(stager, product),
			}
		}
	}
}

func (stager *Stager) treeWBSinDiagram(diagram *Diagram, task *Task, parentNode *tree.Node) {
	taskNode := &tree.Node{
		Name:       task.ComputedPrefix + " " + task.Name,
		IsExpanded: slices.Index(diagram.TasksWhoseNodeIsExpanded, task) != -1,

		HasCheckboxButton:  true,
		IsCheckboxDisabled: !diagram.IsChecked,

		HasToolTip:      true,
		ToolTipPosition: tree.Above,
		ToolTipText:     "Add task to diagram",

		IsNodeClickable: true,
	}
	parentNode.Children = append(parentNode.Children, taskNode)

	if _, ok := diagram.map_Task_TaskShape[task]; ok {
		taskNode.IsChecked = true
	}

	// what to do when the task node is clicked
	taskNode.Impl = &tree.FunctionalNodeProxy{
		OnUpdate: OnUpdateElementInDiagram(
			stager,
			diagram,
			task,
			&diagram.TasksWhoseNodeIsExpanded,
			&diagram.Task_Shapes,
			&diagram.map_Task_TaskShape),
	}

	// if task has a parent task, add a button to show/hide the link to the parent
	if parentTask := task.parentTask; parentTask != nil {
		if _, ok := diagram.map_Task_TaskShape[parentTask]; ok {
			if _, ok := diagram.map_Task_TaskShape[task]; ok {

				showHideCompositionButton := &tree.Button{
					Name: GetGongstructNameFromPointer(task) + " " + string(buttons.BUTTON_add),

					HasToolTip:      true,
					ToolTipPosition: tree.Right,
				}

				if compositionShape, ok := diagram.map_Task_TaskCompositionShape[task]; !ok {
					showHideCompositionButton.Icon = string(buttons.BUTTON_unfold_more)
					showHideCompositionButton.ToolTipText = "Show link from \"" + parentTask.Name +
						"\" to \"" + task.Name + "\""

					showHideCompositionButton.Impl = &tree.FunctionalButtonProxy{
						OnUpdated: OnAddCompositionShape(stager, diagram, parentTask, task, &diagram.TaskComposition_Shapes),
					}
				} else {
					showHideCompositionButton.Icon = string(buttons.BUTTON_unfold_less)
					showHideCompositionButton.ToolTipText = "Hide link from \"" + parentTask.Name +
						"\" to \"" + task.Name + "\""

					showHideCompositionButton.Impl = &tree.FunctionalButtonProxy{
						OnUpdated: OnRemoveCompositionShape(stager, diagram, compositionShape, &diagram.TaskComposition_Shapes),
					}
				}
				taskNode.Buttons = append(taskNode.Buttons, showHideCompositionButton)
			}
		}
	}

	for _, task := range task.SubTasks {
		stager.treeWBSinDiagram(diagram, task, taskNode)
	}
}
