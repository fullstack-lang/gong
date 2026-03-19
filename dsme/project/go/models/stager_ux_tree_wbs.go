package models

import (
	_ "embed"
	"fmt"
	"slices"

	"github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

//go:embed visibility_off.svg
var visibilityOffSVG string

func (stager *Stager) treeWBSinDiagram(diagram *Diagram, task *Task, parentNode *tree.Node) {
	stage := stager.stage

	svgIconLinkVisibilityOn := &tree.SVGIcon{
		Name: "Show Visibility On",
		SVG:  `<svg xmlns="http://www.w3.org/2000/svg" height="24" viewBox="0 -960 960 960" width="24"><path d="M 40,-480 C 160,-220 800,-220 920,-480 C 800,-740 160,-740 40,-480 Z M 640,-480 C 640,-568.4 568.4,-640 480,-640 C 391.6,-640 320,-568.4 320,-480 C 320,-391.6 391.6,-320 480,-320 C 568.4,-320 640,-391.6 640,-480 Z M 354,-491.6 L 427.6,-430 A 40 40 0 0 0 479.2,-491.2 L 405.6,-552.8 A 40 40 0 0 0 354,-491.6 Z M 480.8,-468.8 L 554.4,-407.2 A 40 40 0 0 0 606,-468.4 L 532.4,-530 A 40 40 0 0 0 480.8,-468.8 Z" fill-rule="evenodd"/></svg>`,
	}
	_ = svgIconLinkVisibilityOn

	svgIconLinkVisibilityOff := &tree.SVGIcon{
		Name: "Show Visibility Off",
		SVG:  visibilityOffSVG,
	}
	_ = svgIconLinkVisibilityOff

	taskNode := addNodeToTree(
		stager,
		diagram,
		parentNode,
		task,
		task.parentTask,
		&diagram.TasksWhoseNodeIsExpanded,
		&diagram.Task_Shapes,
		&diagram.map_Task_TaskShape,
		diagram.map_Task_TaskCompositionShape,
		&diagram.TaskComposition_Shapes,
	)

	addAddItemButton(stager, &diagram.TasksWhoseNodeIsExpanded, task, nil, taskNode, &task.SubTasks, diagram, &diagram.Task_Shapes, &diagram.TaskComposition_Shapes)

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
				Name:                    product.GetName(),
				IsExpanded:              true,
				IsNodeClickable:         true,
				CheckboxHasToolTip:      true,
				CheckboxToolTipPosition: tree.Right,
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

					if ok {
						inputProductNode.CheckboxToolTipText = "Uncheck to remove shape from diagram"
					} else {
						inputProductNode.CheckboxToolTipText = "Check to shape to diagram"
					}

					inputProductNode.OnUpdate = func(stage *tree.Stage, stagedNode, frontNode *tree.Node) {
						if frontNode.IsChecked && !stagedNode.IsChecked {
							stagedNode.IsChecked = true
							addAssociationShapeToDiagram(stager, task, product, &diagram.TaskInputShapes)
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
							Name: diagram.GetName(),
							Icon: string(buttons.BUTTON_visibility_off),
							// SVGIcon:         svgIconLinkVisibilityOff,
							ToolTipText:     "Hide link from diagram",
							HasToolTip:      true,
							ToolTipPosition: tree.Right,
							OnUpdate: func(_ *tree.Stage, _ *tree.Button) {
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
			IsExpanded:           slices.Index(diagram.TasksWhoseOutputNodeIsExpanded, task) != -1,
			IsNodeClickable:      true,
			IsWithPreceedingIcon: true,
			PreceedingIcon:       string(buttons.BUTTON_output),
		}
		taskNode.Children = append(taskNode.Children, outputProductsNode)

		outputProductsNode.OnUpdate = onUpdateExpandableNode(stager, task, &diagram.TasksWhoseOutputNodeIsExpanded)

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
			if _, ok := diagram.map_Product_ProductShape[product]; ok {
				if _, ok := diagram.map_Task_TaskShape[task]; ok {

					outputProductNode.HasCheckboxButton = true

					taskProductKey := taskProductKey{
						Task:    task,
						Product: product,
					}
					taskOutputShape, ok := diagram.map_Task_TaskOutputShape[taskProductKey]
					outputProductNode.IsChecked = ok

					if ok {
						outputProductNode.CheckboxToolTipText = "Uncheck to remove shape from diagram"
					} else {
						outputProductNode.CheckboxToolTipText = "Check to shape to diagram"
					}

					outputProductNode.OnUpdate = func(stage *tree.Stage, stagedNode, frontNode *tree.Node) {
						if frontNode.IsChecked && !stagedNode.IsChecked {
							stagedNode.IsChecked = true
							addAssociationShapeToDiagram(stager, task, product, &diagram.TaskOutputShapes)
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
							Name:            diagram.GetName(),
							Icon:            string(buttons.BUTTON_visibility_off),
							ToolTipText:     "Hide link from diagram",
							HasToolTip:      true,
							ToolTipPosition: tree.Right,
							OnUpdate: func(_ *tree.Stage, _ *tree.Button) {
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
