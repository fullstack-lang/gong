package models

import (
	"fmt"
	"slices"

	"github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) treeWBSinDiagram(diagram *Diagram, task *Task, parentNode *tree.Node) {
	taskNode := addNodeToTree(
		stager,
		diagram,
		parentNode,
		task,
		&diagram.TasksWhoseNodeIsExpanded,
		&diagram.Task_Shapes,
		&diagram.map_Task_TaskShape,
	)

	addAddItemButton(stager, &diagram.TasksWhoseNodeIsExpanded, task, nil, taskNode, &task.SubTasks, diagram, &diagram.Task_Shapes, &diagram.TaskComposition_Shapes)

	// if task has a parent task, add a button to show/hide the link to the parent
	addShowHideCompositionButton(
		stager,
		task,
		task.parentTask,
		taskNode,
		diagram.map_Task_TaskShape,
		diagram.map_Task_TaskCompositionShape,
		&diagram.TaskComposition_Shapes,
	)

	for _, task := range task.SubTasks {
		stager.treeWBSinDiagram(diagram, task, taskNode)
	}

	if len(task.Inputs) > 0 {
		inputProductsNode := &tree.Node{
			Name:                 fmt.Sprintf("In (%d)", len(task.Inputs)),
			IsExpanded:           slices.Index(diagram.TasksWhoseInputNodeIsExpanded, task) != -1,
			IsNodeClickable:      true,
			IsWithPreceedingIcon: true,
			PreceedingIcon:       string(buttons.BUTTON_input),
		}
		taskNode.Children = append(taskNode.Children, inputProductsNode)

		inputProductsNode.OnUpdate = onUpdateExpandableNode(stager, task, &diagram.TasksWhoseInputNodeIsExpanded)

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

					inputProductNode.HasCheckboxButton = true

					taskProductKey := taskProductKey{
						Task:    task,
						Product: product,
					}
					taskInputShape, ok := diagram.map_Task_TaskInputShape[taskProductKey]
					inputProductNode.IsChecked = ok

					inputProductNode.OnUpdate = func(stage *tree.Stage, stagedNode, frontNode *tree.Node) {
						if frontNode.IsChecked && !stagedNode.IsChecked {
							stagedNode.IsChecked = true
							addAssociationShapeToDiagram(stager, task, product, &diagram.TaskInputShapes)
							stager.stage.Commit()
						}
						if !frontNode.IsChecked && stagedNode.IsChecked {
							stagedNode.IsChecked = false
							if taskInputShape != nil {
								taskInputShape.UnstageVoid(stager.stage)
								idx := slices.Index(diagram.TaskInputShapes, taskInputShape)
								if idx != -1 {
									diagram.TaskInputShapes = slices.Delete(diagram.TaskInputShapes, idx, idx+1)
								}
							}
							stager.stage.Commit()
						}
					}
				}
			}
		}

	}
	if len(task.Outputs) > 0 {
		outputProductsNode := &tree.Node{
			Name:                 fmt.Sprintf("Out (%d)", len(task.Outputs)),
			IsExpanded:           slices.Index(diagram.TasksWhoseOutputNodeIsExpanded, task) != -1,
			IsNodeClickable:      true,
			IsWithPreceedingIcon: true,
			PreceedingIcon:       string(buttons.BUTTON_output),
		}
		taskNode.Children = append(taskNode.Children, outputProductsNode)

		outputProductsNode.OnUpdate = onUpdateExpandableNode(stager, task, &diagram.TasksWhoseOutputNodeIsExpanded)

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

					outputProductNode.HasCheckboxButton = true

					taskProductKey := taskProductKey{
						Task:    task,
						Product: product,
					}
					taskOutputShape, ok := diagram.map_Task_TaskOutputShape[taskProductKey]
					outputProductNode.IsChecked = ok

					outputProductNode.OnUpdate = func(stage *tree.Stage, stagedNode, frontNode *tree.Node) {
						if frontNode.IsChecked && !stagedNode.IsChecked {
							stagedNode.IsChecked = true
							addAssociationShapeToDiagram(stager, task, product, &diagram.TaskOutputShapes)
							stager.stage.Commit()
						}
						if !frontNode.IsChecked && stagedNode.IsChecked {
							stagedNode.IsChecked = false
							if taskOutputShape != nil {
								taskOutputShape.UnstageVoid(stager.stage)
								idx := slices.Index(diagram.TaskOutputShapes, taskOutputShape)
								if idx != -1 {
									diagram.TaskOutputShapes = slices.Delete(diagram.TaskOutputShapes, idx, idx+1)
								}
							}
							stager.stage.Commit()
						}
					}
				}
			}
		}
	}
}
