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
		&diagram.ResourcesWhoseNodeIsExpanded,
		&diagram.Resource_Shapes,
		&diagram.map_Resource_ResourceShape,
	)

	addAddItemButton(stager, &diagram.ResourcesWhoseNodeIsExpanded, resource, nil, resourceNode, &resource.SubResources, diagram, &diagram.Resource_Shapes, &diagram.ResourceComposition_Shapes)

	// if resource has a parent resource, add a button to show/hide the link to the parent
	addShowHideCompositionButton(
		stager,
		diagram,
		resource,
		resource.parentResource,
		resourceNode,
		diagram.map_Resource_ResourceShape,
		diagram.map_Resource_ResourceCompositionShape,
		&diagram.ResourceComposition_Shapes,
	)

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
						showHideRelationButton.ToolTipText = "Hide link from \"" + resource.Name +
							"\" to \"" + task.Name + "\""
						showHideRelationButton.Impl = &tree.FunctionalButtonProxy{
							OnUpdated: onRemoveAssociationShape(stager, shape, &diagram.ResourceTaskShapes),
						}
					} else {
						showHideRelationButton.Icon = string(buttons.BUTTON_visibility)
						showHideRelationButton.ToolTipText = "Show link from \"" + resource.Name +
							"\" to \"" + task.Name + "\""
						showHideRelationButton.Impl = &tree.FunctionalButtonProxy{
							OnUpdated: onAddAssociationShape(stager, resource, task, &diagram.ResourceTaskShapes),
						}
					}
				}
			}
		}
	}
}
