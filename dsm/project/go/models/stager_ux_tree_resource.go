package models

import (
	"fmt"

	"github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) treeResourceinDiagram(diagram *Diagram, resource *Resource, parentNode *tree.Node) {
	resourceNodeConf := TreeNodeShapeAndLinkConfiguration[
		*Resource, Resource, // AT, AT_
		*ResourceShape, ResourceShape, // CT, CT_
		*ResourceCompositionShape, ResourceCompositionShape, // ACT, ACT_
		*Diagram, // DiagramType
	]{
		TreeNodeAndShapeConfiguration: TreeNodeAndShapeConfiguration[
			*Resource, Resource, // AT, AT_
			*ResourceShape, ResourceShape, // CT, CT_
			*Diagram, // DiagramType
		]{
			TreeNodeConfiguration: TreeNodeConfiguration[
				*Resource, Resource, // AT, AT_
				*Diagram, // DiagramType
			]{
				diagram:                     diagram,
				parentNode:                  parentNode,
				element:                     resource,
				parentElement:               resource.parentResource,
				elementsWhoseNodeIsExpanded: &diagram.ResourcesWhoseNodeIsExpanded,
			},
			shapes:    &diagram.Resource_Shapes,
			shapesMap: diagram.map_Resource_ResourceShape,
		},
		map_Element_CompositionShape: diagram.map_Resource_ResourceCompositionShape,
		compositionShapes:            &diagram.ResourceComposition_Shapes,
	}
	resourceNode := addNodeToTreeWithConf(stager, resourceNodeConf)

	conf := ItemShapeAndLinkButtonConfiguration[
		Resource, *Resource, // AT, PAT (Added Element)
		Resource, *Resource, // ParentAT, PParentAT (Parent Element)
		ResourceShape, *ResourceShape, // CT, PCT (Concrete Shape)
		ResourceCompositionShape, *ResourceCompositionShape, // ACT, PACT (Association Shape)
	]{
		ItemAndShapeButtonConfiguration: ItemAndShapeButtonConfiguration[
			Resource, *Resource, // AT, PAT (Added Element)
			Resource, *Resource, // ParentAT, PParentAT (Parent Element)
			ResourceShape, *ResourceShape, // CT, PCT (Concrete Shape)
		]{
			ItemButtonConfiguration: ItemButtonConfiguration[
				Resource, *Resource, // AT, PAT (Added Element)
				Resource, *Resource, // ParentAT, PParentAT (Parent Element)
			]{
				parentNode:                         resourceNode,
				sliceForNewAddedItem:               &resource.SubResources,
				isParentNodeExpandedByAddOperation: true,
				parentNodeExpansionType:            parentNodeExpansionTypeBySlice,
				parentNodeExpansionSliceEncoding:   &diagram.ResourcesWhoseNodeIsExpanded,
				parentElement:                      resource,
			},
			receivingDiagram:      diagram,
			sliceForNewAddedShape: &diagram.Resource_Shapes,
		},
		sliceForNewCompositionShapes: &diagram.ResourceComposition_Shapes,
	}
	addCreateItemShapeAndLinkButton(stager, conf)

	for _, subResource := range resource.SubResources {
		stager.treeResourceinDiagram(diagram, subResource, resourceNode)
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
							OnClick: func() {
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
