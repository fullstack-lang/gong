package models

import (
	"fmt"

	"github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) treeRBSinDiagram(diagram *Diagram, resource *Resource, parentNode *tree.Node) {
	resourceNode := addNodeToTree(
		stager,
		diagram,
		parentNode,
		resource,
		resource.parentResource,
		&diagram.ResourcesWhoseNodeIsExpanded,
		&diagram.Resource_Shapes,
		diagram.map_Resource_ResourceShape,
		diagram.map_Resource_ResourceCompositionShape,
		&diagram.ResourceComposition_Shapes,
	)

	addAddItemButton(stager, &diagram.ResourcesWhoseNodeIsExpanded, resource, nil, resourceNode, &resource.SubResources, diagram, &diagram.Resource_Shapes, &diagram.ResourceComposition_Shapes)

	for _, subResource := range resource.SubResources {
		stager.treeRBSinDiagram(diagram, subResource, resourceNode)
	}

	if len(resource.Tasks) > 0 {
		tasksNode := &tree.Node{
			Name:                 fmt.Sprintf("Tasks (%d)", len(resource.Tasks)),
			IsNodeClickable:      true,
			IsExpanded:           true,
			IsWithPreceedingIcon: true,
			PreceedingIcon:       string(buttons.BUTTON_output),
		}
		resourceNode.Children = append(resourceNode.Children, tasksNode)

		for _, task := range resource.Tasks {
			taskNode := &tree.Node{
				Name:                    task.Name,
				IsNodeClickable:         true,
				CheckboxHasToolTip:      true,
				CheckboxToolTipPosition: tree.Right,
				HasCheckboxButton:       true,
			}
			tasksNode.Children = append(tasksNode.Children, taskNode)
			taskNode.IsCheckboxDisabled = true

			if _, ok := diagram.map_Task_TaskShape[task]; ok {
				if _, ok := diagram.map_Resource_ResourceShape[resource]; ok {

					taskNode.IsCheckboxDisabled = false

					key := resourceTaskKey{Resource: resource, Task: task}
					resourceTaskShape, ok := diagram.map_Resource_ResourceTaskShape[key]
					taskNode.IsChecked = ok

					if ok {
						taskNode.CheckboxToolTipText = "Uncheck to remove shape from diagram"
					} else {
						taskNode.CheckboxToolTipText = "Check to add shape to diagram"
					}

					taskNode.OnUpdate = func(stage *tree.Stage, stagedNode, frontNode *tree.Node) {
						if frontNode.IsChecked && !stagedNode.IsChecked {
							stagedNode.IsChecked = true
							addAssociationShapeToDiagram(stager, resource, task, &diagram.ResourceTaskShapes)
							stager.stage.Commit()
						}
						if !frontNode.IsChecked && stagedNode.IsChecked {
							stagedNode.IsChecked = false
							resourceTaskShape.UnstageVoid(stager.stage)
							stager.stage.Commit()
						}
					}

					taskNode.Buttons = []*tree.Button{
						{
							Name:            diagram.GetName(),
							Icon:            string(buttons.BUTTON_visibility_off),
							ToolTipText:     "Hide link from diagram",
							HasToolTip:      true,
							ToolTipPosition: tree.Right,
							OnUpdate: func(_ *tree.Stage, _ *tree.Button) {
								resourceTaskShape.SetIsHidden(!resourceTaskShape.GetIsHidden())
								stager.stage.Commit()
							},
						},
					}
					if ok {
						if resourceTaskShape.GetIsHidden() {
							taskNode.Buttons[0].Icon = string(buttons.BUTTON_visibility)
							taskNode.Buttons[0].ToolTipText = "Show link on diagram"
						}
					} else {
						taskNode.Buttons[0].IsDisabled = true
					}
				}
			}
		}
	}
}
