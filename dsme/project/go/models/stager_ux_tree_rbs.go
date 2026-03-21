package models

import (
	"fmt"
	"slices"

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
			Name:            fmt.Sprintf("Tasks (%d)", len(resource.Tasks)),
			IsNodeClickable: true,
			IsExpanded:      true,
		}
		resourceNode.Children = append(resourceNode.Children, tasksNode)

		for _, task := range resource.Tasks {
			taskNode := &tree.Node{
				Name:            task.Name,
				IsNodeClickable: true,
			}
			tasksNode.Children = append(tasksNode.Children, taskNode)

			if _, ok := diagram.map_Task_TaskShape[task]; ok {
				if _, ok := diagram.map_Resource_ResourceShape[resource]; ok {

					showHideRelationButton := &tree.Button{
						Name:            task.GetName() + " add resource relation",
						HasToolTip:      true,
						ToolTipPosition: tree.Right,
					}
					taskNode.Buttons = append(taskNode.Buttons, showHideRelationButton)

					key := resourceTaskKey{Resource: resource, Task: task}
					if shape, ok := diagram.map_Resource_ResourceTaskShape[key]; ok {
						showHideRelationButton.Icon = string(buttons.BUTTON_visibility_off)
						showHideRelationButton.ToolTipText = "Remove link from \"" + resource.Name +
							"\" to \"" + task.Name + "\""
						showHideRelationButton.OnUpdate = func(_ *tree.Stage, _ *tree.Button) {
							shape.UnstageVoid(stager.stage)
							idx := slices.Index(diagram.ResourceTaskShapes, shape)
							diagram.ResourceTaskShapes = slices.Delete(diagram.ResourceTaskShapes, idx, idx+1)
							stager.stage.Commit()
						}
					} else {
						showHideRelationButton.Icon = string(buttons.BUTTON_visibility)
						showHideRelationButton.ToolTipText = "Insert link from \"" + resource.Name +
							"\" to \"" + task.Name + "\""
						showHideRelationButton.OnUpdate = func(_ *tree.Stage, _ *tree.Button) {
							addAssociationShapeToDiagram(stager, resource, task, &diagram.ResourceTaskShapes)
							stager.stage.Commit()
						}
					}
				}
			}
		}
	}
}
