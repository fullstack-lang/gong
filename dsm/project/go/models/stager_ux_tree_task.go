package models

import (
	_ "embed"
	"fmt"
	"slices"

	"github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) treeTask(diagramHierarchy *DiagramHierarchy, task *Task, parentNode *tree.Node) {
	stage := stager.stage

	/*
		taskNode := addNodeToTree(stager,
			TreeNodeShapeAndLinkConfiguration{
					diagram:                      diagram,
					parentNode:                   parentNode,
					element:                      task,
					parentElement:                task.parentTask,
					elementsWhoseNodeIsExpanded:  &diagram.TasksWhoseNodeIsExpanded,
					shapes:    					  &diagram.Task_Shapes,
					shapesMap: 					  diagram.map_Task_TaskShape,
					map_Element_CompositionShape: diagram.map_Task_TaskCompositionShape,
					compositionShapes:            &diagram.TaskComposition_Shapes,
		})
	*/
	taskNode := addNodeToTree(stager,
		TreeNodeShapeAndLinkConfiguration[
			*Task, Task, // AT, AT_
			*TaskShape, TaskShape, // CT, CT_
			*TaskCompositionShape, TaskCompositionShape, // ACT, ACT_
			*DiagramHierarchy, // DiagramType
		]{
			TreeNodeAndShapeConfiguration: TreeNodeAndShapeConfiguration[
				*Task, Task, // AT, AT_
				*TaskShape, TaskShape, // CT, CT_
				*DiagramHierarchy, // DiagramType
			]{
				TreeNodeConfiguration: TreeNodeConfiguration[
					*Task, Task, // AT, AT_
					*DiagramHierarchy, // DiagramType
				]{
					diagram:            diagramHierarchy,
					parentNode:                  parentNode,
					element:                     task,
					parentElement:               task.parentTask,
					elementsWhoseNodeIsExpanded: &diagramHierarchy.TasksWhoseNodeIsExpanded,
				},
				shapes:    &diagramHierarchy.Task_Shapes,
				shapesMap: diagramHierarchy.map_Task_TaskShape,
			},
			map_Element_CompositionShape: diagramHierarchy.map_Task_TaskCompositionShape,
			compositionShapes:            &diagramHierarchy.TaskComposition_Shapes,
		})

	if task.IsImport && task.ReferencedTask != nil {
		taskNode.Name = "🔗 " + task.ReferencedTask.Name
		taskNode.CheckboxToolTipText = "Add imported task to diagram"
	}

	conf := ItemShapeAndLinkButtonConfiguration[
		Task, *Task, // AT, PAT (Added Element)
		Task, *Task, // ParentAT, PParentAT (Parent Element)
		TaskShape, *TaskShape, // CT, PCT (Concrete Shape)
		TaskCompositionShape, *TaskCompositionShape, // ACT, PACT (Association Shape)
	]{
		ItemAndShapeButtonConfiguration: ItemAndShapeButtonConfiguration[
			Task, *Task, // AT, PAT (Added Element)
			Task, *Task, // ParentAT, PParentAT (Parent Element)
			TaskShape, *TaskShape, // CT, PCT (Concrete Shape)
		]{
			ItemButtonConfiguration: ItemButtonConfiguration[
				Task, *Task, // AT, PAT (Added Element)
				Task, *Task, // ParentAT, PParentAT (Parent Element)
			]{
				parentNode:                         taskNode,
				sliceForNewAddedItem:               &task.SubTasks,
				isParentNodeExpandedByAddOperation: true,
				parentNodeExpansionType:            parentNodeExpansionTypeBySlice,
				parentNodeExpansionSliceEncoding:   &diagramHierarchy.TasksWhoseNodeIsExpanded,
				parentElement:                      task,
			},
			receivingDiagram:      diagramHierarchy,
			sliceForNewAddedShape: &diagramHierarchy.Task_Shapes,
		},
		sliceForNewCompositionShapes: &diagramHierarchy.TaskComposition_Shapes,
	}
	addCreateItemShapeAndLinkButton(stager, conf)

	if taskShape, ok := diagramHierarchy.map_Task_TaskShape[task]; ok {
		button := &tree.Button{
			Name:            "Show dates on diagram",
			Icon:            string(buttons.BUTTON_calendar_month),
			ToolTipText:     "Show dates on diagram",
			HasToolTip:      true,
			ToolTipPosition: tree.Right,
			OnClick: func() {
				taskShape.IsShowDate = !taskShape.IsShowDate
				stage.Commit()
			},
		}
		if taskShape.IsShowDate {
			button.Name = "Hide dates on diagram"
			button.ToolTipText = "Hide dates on diagram"
			button.Icon = string(buttons.BUTTON_calendar_today)
		}
		if taskNode.Menu == nil {
			taskNode.Menu = &tree.Menu{Name: "Menu"}
		}
		taskNode.Menu.Buttons = append(taskNode.Menu.Buttons, button)
	}

	for _, task := range task.SubTasks {
		stager.treeTask(diagramHierarchy, task, taskNode)
	}

	if len(task.Inputs) > 0 {
		inputProductsNode := &tree.Node{
			Name:                 fmt.Sprintf("In (%d)", len(task.Inputs)),
			IsExpanded:           slices.Contains(diagramHierarchy.TasksWhoseInputNodeIsExpanded, task),
			IsNodeClickable:      true,
			IsWithPreceedingIcon: true,
			PreceedingIcon:       string(buttons.BUTTON_input),
		}
		taskNode.Children = append(taskNode.Children, inputProductsNode)

		inputProductsNode.OnUpdate = onUpdateExpandableNode(stager, task, &diagramHierarchy.TasksWhoseInputNodeIsExpanded)

		for _, product := range task.Inputs {
			inputProductNode := &tree.Node{
				Name:                    product.GetName(),
				IsExpanded:              true,
				IsNodeClickable:         true,
				CheckboxHasToolTip:      true,
				CheckboxToolTipPosition: tree.Right,
			}
			inputProductsNode.Children = append(inputProductsNode.Children, inputProductNode)

			// if input task is present in diagram as well as the input product
			// display the show/hide input relation button
			if _, ok := diagramHierarchy.map_Product_ProductShape[product]; ok {
				if _, ok := diagramHierarchy.map_Task_TaskShape[task]; ok {

					inputProductNode.HasCheckboxButton = true

					taskProductKey := taskProductKey{
						Task:    task,
						Product: product,
					}
					taskInputShape, ok := diagramHierarchy.map_Task_TaskInputShape[taskProductKey]
					inputProductNode.IsChecked = ok

					if ok {
						inputProductNode.CheckboxToolTipText = "Uncheck to remove shape from diagram"
					} else {
						inputProductNode.CheckboxToolTipText = "Check to shape to diagram"
					}

					inputProductNode.OnUpdate = func(stage *tree.Stage, stagedNode, frontNode *tree.Node) {
						if frontNode.IsChecked && !stagedNode.IsChecked {
							stagedNode.IsChecked = true
							addAssociationShapeToDiagram(stager, task, product, &diagramHierarchy.TaskInputShapes)
							stager.stage.Commit()
						}
						if !frontNode.IsChecked && stagedNode.IsChecked {
							stagedNode.IsChecked = false
							taskInputShape.UnstageVoid(stager.stage)
							stager.stage.Commit()
						}
					}

					inputProductNode.Buttons = []*tree.Button{
						{
							Name: diagramHierarchy.GetName(),
							Icon: string(buttons.BUTTON_visibility_off),
							// SVGIcon:         svgIconLinkVisibilityOff,
							ToolTipText:     "Hide link from diagram",
							HasToolTip:      true,
							ToolTipPosition: tree.Right,
							OnClick: func() {
								taskInputShape.SetIsHidden(!taskInputShape.GetIsHidden())
								stage.Commit()
							},
						},
					}
					if ok {
						if taskInputShape.GetIsHidden() {
							inputProductNode.Buttons[0].Icon = string(buttons.BUTTON_visibility)
							// inputProductNode.Buttons[0].SVGIcon = svgIconLinkVisibilityOn
							inputProductNode.Buttons[0].ToolTipText = "Show link on diagram"
						}
					} else {
						inputProductNode.Buttons[0].IsDisabled = true
					}
				}
			}
		}

	}
	if len(task.Outputs) > 0 {
		outputProductsNode := &tree.Node{
			Name:                 fmt.Sprintf("Out (%d)", len(task.Outputs)),
			IsExpanded:           slices.Contains(diagramHierarchy.TasksWhoseOutputNodeIsExpanded, task),
			IsNodeClickable:      true,
			IsWithPreceedingIcon: true,
			PreceedingIcon:       string(buttons.BUTTON_output),
		}
		taskNode.Children = append(taskNode.Children, outputProductsNode)

		outputProductsNode.OnUpdate = onUpdateExpandableNode(stager, task, &diagramHierarchy.TasksWhoseOutputNodeIsExpanded)

		for _, product := range task.Outputs {
			outputProductNode := &tree.Node{
				Name:                    product.GetName(),
				IsExpanded:              true,
				IsNodeClickable:         true,
				CheckboxHasToolTip:      true,
				CheckboxToolTipPosition: tree.Right,
			}
			outputProductsNode.Children = append(outputProductsNode.Children, outputProductNode)

			// if output task is present in diagram as well as the output product
			// display the show/hide output relation button
			if _, ok := diagramHierarchy.map_Product_ProductShape[product]; ok {
				if _, ok := diagramHierarchy.map_Task_TaskShape[task]; ok {

					outputProductNode.HasCheckboxButton = true

					taskProductKey := taskProductKey{
						Task:    task,
						Product: product,
					}
					taskOutputShape, ok := diagramHierarchy.map_Task_TaskOutputShape[taskProductKey]
					outputProductNode.IsChecked = ok

					if ok {
						outputProductNode.CheckboxToolTipText = "Uncheck to remove shape from diagram"
					} else {
						outputProductNode.CheckboxToolTipText = "Check to shape to diagram"
					}

					outputProductNode.OnUpdate = func(stage *tree.Stage, stagedNode, frontNode *tree.Node) {
						if frontNode.IsChecked && !stagedNode.IsChecked {
							stagedNode.IsChecked = true
							addAssociationShapeToDiagram(stager, task, product, &diagramHierarchy.TaskOutputShapes)
							stager.stage.Commit()
						}
						if !frontNode.IsChecked && stagedNode.IsChecked {
							stagedNode.IsChecked = false
							taskOutputShape.UnstageVoid(stager.stage)
							stager.stage.Commit()
						}
					}

					outputProductNode.Buttons = []*tree.Button{
						{
							Name:            diagramHierarchy.GetName(),
							Icon:            string(buttons.BUTTON_visibility_off),
							ToolTipText:     "Hide link from diagram",
							HasToolTip:      true,
							ToolTipPosition: tree.Right,
							OnClick: func() {
								taskOutputShape.SetIsHidden(!taskOutputShape.GetIsHidden())
								stage.Commit()
							},
						},
					}
					if ok {
						if taskOutputShape.GetIsHidden() {
							outputProductNode.Buttons[0].Icon = string(buttons.BUTTON_visibility)
							outputProductNode.Buttons[0].ToolTipText = "Show link on diagram"
						}
					} else {
						outputProductNode.Buttons[0].IsDisabled = true
					}
				}
			}
		}
	}
}
