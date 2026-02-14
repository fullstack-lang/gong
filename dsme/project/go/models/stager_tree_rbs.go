package models

import (
	"fmt"
	"slices"

	"github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) treeRBSinDiagram(diagram *Diagram, resource *Resource, parentNode *tree.Node) {
	resourceNode := &tree.Node{
		Name:       resource.ComputedPrefix + " " + resource.Name,
		IsExpanded: slices.Index(diagram.ResourcesWhoseNodeIsExpanded, resource) != -1,

		HasCheckboxButton:  true,
		IsCheckboxDisabled: !diagram.IsChecked,

		HasToolTip:      true,
		ToolTipPosition: tree.Above,
		ToolTipText:     "Add resource to diagram",

		IsNodeClickable: true,
	}
	parentNode.Children = append(parentNode.Children, resourceNode)

	addAddItemButton(stager, &diagram.ResourcesWhoseNodeIsExpanded, resource, nil, resourceNode, &resource.SubResources, diagram, &diagram.Resource_Shapes, &diagram.ResourceComposition_Shapes)

	if _, ok := diagram.map_Resource_ResourceShape[resource]; ok {
		resourceNode.IsChecked = true
	}

	// what to do when the resource node is clicked
	resourceNode.Impl = &tree.FunctionalNodeProxy{
		OnUpdate: onUpdateElementInDiagram(
			stager,
			diagram,
			resource,
			&diagram.ResourcesWhoseNodeIsExpanded,
			&diagram.Resource_Shapes,
			&diagram.map_Resource_ResourceShape),
	}

	// if resource has a parent resource, add a button to show/hide the link to the parent
	if parentResource := resource.parentResource; parentResource != nil {
		if _, ok := diagram.map_Resource_ResourceShape[parentResource]; ok {
			if _, ok := diagram.map_Resource_ResourceShape[resource]; ok {

				showHideCompositionButton := &tree.Button{
					Name: GetGongstructNameFromPointer(resource) + " " + string(buttons.BUTTON_add),

					HasToolTip:      true,
					ToolTipPosition: tree.Right,
				}

				if compositionShape, ok := diagram.map_Resource_ResourceCompositionShape[resource]; !ok {
					showHideCompositionButton.Icon = string(buttons.BUTTON_visibility)
					showHideCompositionButton.ToolTipText = "Show link from \"" + parentResource.Name +
						"\" to \"" + resource.Name + "\""

					showHideCompositionButton.Impl = &tree.FunctionalButtonProxy{
						OnUpdated: onAddAssociationShape(stager, parentResource, resource, &diagram.ResourceComposition_Shapes),
					}
				} else {
					showHideCompositionButton.Icon = string(buttons.BUTTON_visibility_off)
					showHideCompositionButton.ToolTipText = "Hide link from \"" + parentResource.Name +
						"\" to \"" + resource.Name + "\""

					showHideCompositionButton.Impl = &tree.FunctionalButtonProxy{
						OnUpdated: onRemoveAssociationShape(stager, compositionShape, &diagram.ResourceComposition_Shapes),
					}
				}
				resourceNode.Buttons = append(resourceNode.Buttons, showHideCompositionButton)
			}
		}
	}

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
