package models

import (
	"fmt"

	"github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) treeStakeholderBSinDiagram(diagram *Diagram, stakeholder *Stakeholder, parentNode *tree.Node) {
	resourceNode := addNodeToTree(
		stager,
		diagram,
		parentNode,
		stakeholder,
		stakeholder.parentStakeholder,
		&diagram.ResourcesWhoseNodeIsExpanded,
		&diagram.Stakeholder_Shapes,
		diagram.map_Stakeholder_StakeholderShape,
		diagram.map_Resource_ResourceCompositionShape,
		&diagram.ResourceComposition_Shapes,
	)

	addAddItemButton(stager, &diagram.ResourcesWhoseNodeIsExpanded, stakeholder, nil, resourceNode, &stakeholder.SubStakeholders, diagram, &diagram.Stakeholder_Shapes, &diagram.ResourceComposition_Shapes)

	for _, subResource := range stakeholder.SubStakeholders {
		stager.treeStakeholderBSinDiagram(diagram, subResource, resourceNode)
	}

	if len(stakeholder.Concerns) > 0 {
		tasksNode := &tree.Node{
			Name:                 fmt.Sprintf("Concerns (%d)", len(stakeholder.Concerns)),
			IsNodeClickable:      true,
			IsExpanded:           true,
			IsWithPreceedingIcon: true,
			PreceedingIcon:       string(buttons.BUTTON_output),
		}
		resourceNode.Children = append(resourceNode.Children, tasksNode)

		for _, concern := range stakeholder.Concerns {
			n := &tree.Node{
				Name:                    concern.IDAirbus + "- " + concern.Name,
				IsNodeClickable:         true,
				CheckboxHasToolTip:      true,
				CheckboxToolTipPosition: tree.Right,
				HasCheckboxButton:       true,
			}
			tasksNode.Children = append(tasksNode.Children, n)
			n.IsCheckboxDisabled = true

			if _, ok := diagram.map_Concern_ConcernShape[concern]; ok {
				if _, ok := diagram.map_Stakeholder_StakeholderShape[stakeholder]; ok {

					n.IsCheckboxDisabled = false

					key := stakeholderConcernKey{Stakeholder: stakeholder, Concern: concern}
					resourceTaskShape, ok := diagram.map_StakeholderConcernKey_StakeholderConcernShape[key]
					n.IsChecked = ok

					if ok {
						n.CheckboxToolTipText = "Uncheck to remove shape from diagram"
					} else {
						n.CheckboxToolTipText = "Check to add shape to diagram"
					}

					n.OnUpdate = func(stage *tree.Stage, stagedNode, frontNode *tree.Node) {
						if frontNode.IsChecked && !stagedNode.IsChecked {
							stagedNode.IsChecked = true
							addAssociationShapeToDiagram(stager, stakeholder, concern, &diagram.StakeholderConcernShapes)
							stager.stage.Commit()
						}
						if !frontNode.IsChecked && stagedNode.IsChecked {
							stagedNode.IsChecked = false
							resourceTaskShape.UnstageVoid(stager.stage)
							stager.stage.Commit()
						}
					}

					n.Buttons = []*tree.Button{
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
							n.Buttons[0].Icon = string(buttons.BUTTON_visibility)
							n.Buttons[0].ToolTipText = "Show link on diagram"
						}
					} else {
						n.Buttons[0].IsDisabled = true
					}
				}
			}
		}
	}
}
