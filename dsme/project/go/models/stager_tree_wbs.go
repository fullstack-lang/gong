package models

import (
	"fmt"
	"slices"

	"github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

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

	addAddItemButton(stager, &diagram.TasksWhoseNodeIsExpanded, task, taskNode, &task.SubTasks)

	if _, ok := diagram.map_Task_TaskShape[task]; ok {
		taskNode.IsChecked = true
	}

	// what to do when the task node is clicked
	taskNode.Impl = &tree.FunctionalNodeProxy{
		OnUpdate: onUpdateElementInDiagram(
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
					Name:            task.GetName() + " add parent relation",
					HasToolTip:      true,
					ToolTipPosition: tree.Right,
				}

				if compositionShape, ok := diagram.map_Task_TaskCompositionShape[task]; !ok {
					showHideCompositionButton.Icon = string(buttons.BUTTON_unfold_more)
					showHideCompositionButton.ToolTipText = "Show link from \"" + parentTask.Name +
						"\" to \"" + task.Name + "\""

					showHideCompositionButton.Impl = &tree.FunctionalButtonProxy{
						OnUpdated: onAddAssociationShape(stager, diagram, parentTask, task, &diagram.TaskComposition_Shapes),
					}
				} else {
					showHideCompositionButton.Icon = string(buttons.BUTTON_unfold_less)
					showHideCompositionButton.ToolTipText = "Hide link from \"" + parentTask.Name +
						"\" to \"" + task.Name + "\""

					showHideCompositionButton.Impl = &tree.FunctionalButtonProxy{
						OnUpdated: onRemoveAssociationShape(stager, diagram, compositionShape, &diagram.TaskComposition_Shapes),
					}
				}
				taskNode.Buttons = append(taskNode.Buttons, showHideCompositionButton)
			}
		}
	}

	for _, task := range task.SubTasks {
		stager.treeWBSinDiagram(diagram, task, taskNode)
	}

	if len(task.Inputs) > 0 {
		inputProductsNode := &tree.Node{
			Name:                 fmt.Sprintf("(%d)", len(task.Inputs)),
			IsExpanded:           slices.Index(diagram.TasksWhoseInputNodeIsExpanded, task) != -1,
			IsNodeClickable:      true,
			IsWithPreceedingIcon: true,
			PreceedingIcon:       string(buttons.BUTTON_input),
		}
		taskNode.Children = append(taskNode.Children, inputProductsNode)

		inputProductsNode.Impl = &tree.FunctionalNodeProxy{
			OnUpdate: onUpdateExpandableNode(stager, task, &diagram.TasksWhoseInputNodeIsExpanded),
		}

		for _, product := range task.Inputs {
			inputProductNode := &tree.Node{
				Name:            product.GetName(),
				IsExpanded:      true,
				IsNodeClickable: true,
			}
			inputProductsNode.Children = append(inputProductsNode.Children, inputProductNode)

			// if input task is present in diagram as well as the input product
			// display the show/hide input relation button
			if _, ok := diagram.map_Product_ProductShape[product]; ok {
				if _, ok := diagram.map_Task_TaskShape[task]; ok {

					showHideTaskInputButton := &tree.Button{
						Name:            task.GetName() + " add input relation from " + product.GetName(),
						HasToolTip:      true,
						ToolTipPosition: tree.Right,
					}
					inputProductNode.Buttons = append(inputProductNode.Buttons, showHideTaskInputButton)

					taskProductKey := taskProductKey{
						Task:    task,
						Product: product,
					}
					if taskInputShape, ok := diagram.map_Task_TaskInputShape[taskProductKey]; ok {
						showHideTaskInputButton.Icon = string(buttons.BUTTON_unfold_less)
						showHideTaskInputButton.ToolTipText = "Hide link from \"" + task.Name +
							"\" to \"" + product.Name + "\""

						showHideTaskInputButton.Impl = &tree.FunctionalButtonProxy{
							OnUpdated: onRemoveAssociationShape(stager, diagram, taskInputShape, &diagram.TaskInputShapes),
						}
					} else {
						showHideTaskInputButton.Icon = string(buttons.BUTTON_unfold_more)
						showHideTaskInputButton.ToolTipText = "Show link from \"" + task.Name +
							"\" to \"" + product.Name + "\""

						showHideTaskInputButton.Impl = &tree.FunctionalButtonProxy{
							OnUpdated: onAddAssociationShape(stager, diagram, task, product, &diagram.TaskInputShapes),
						}
					}
				}
			}
		}

	}
	if len(task.Outputs) > 0 {
		outputProductsNode := &tree.Node{
			Name:                 fmt.Sprintf("(%d)", len(task.Outputs)),
			IsExpanded:           slices.Index(diagram.TasksWhoseOutputNodeIsExpanded, task) != -1,
			IsNodeClickable:      true,
			IsWithPreceedingIcon: true,
			PreceedingIcon:       string(buttons.BUTTON_output),
		}
		taskNode.Children = append(taskNode.Children, outputProductsNode)

		outputProductsNode.Impl = &tree.FunctionalNodeProxy{
			OnUpdate: onUpdateExpandableNode(stager, task, &diagram.TasksWhoseOutputNodeIsExpanded),
		}

		for _, product := range task.Outputs {
			outputProductNode := &tree.Node{
				Name:            product.GetName(),
				IsExpanded:      true,
				IsNodeClickable: true,
			}
			outputProductsNode.Children = append(outputProductsNode.Children, outputProductNode)

			// if output task is present in diagram as well as the output product
			// display the show/hide output relation button
			if _, ok := diagram.map_Product_ProductShape[product]; ok {
				if _, ok := diagram.map_Task_TaskShape[task]; ok {

					showHideTaskOutputButton := &tree.Button{
						Name:            task.GetName() + " add output relation from " + product.GetName(),
						HasToolTip:      true,
						ToolTipPosition: tree.Right,
					}
					outputProductNode.Buttons = append(outputProductNode.Buttons, showHideTaskOutputButton)

					taskProductKey := taskProductKey{
						Task:    task,
						Product: product,
					}
					if taskOutputShape, ok := diagram.map_Task_TaskOutputShape[taskProductKey]; ok {
						showHideTaskOutputButton.Icon = string(buttons.BUTTON_unfold_less)
						showHideTaskOutputButton.ToolTipText = "Hide link from \"" + task.Name +
							"\" to \"" + product.Name + "\""

						showHideTaskOutputButton.Impl = &tree.FunctionalButtonProxy{
							OnUpdated: onRemoveAssociationShape(stager, diagram, taskOutputShape, &diagram.TaskOutputShapes),
						}
					} else {
						showHideTaskOutputButton.Icon = string(buttons.BUTTON_unfold_more)
						showHideTaskOutputButton.ToolTipText = "Show link from \"" + task.Name +
							"\" to \"" + product.Name + "\""

						showHideTaskOutputButton.Impl = &tree.FunctionalButtonProxy{
							OnUpdated: onAddAssociationShape(stager, diagram, task, product, &diagram.TaskOutputShapes),
						}
					}
				}
			}
		}
	}
}
