package models

import (
	"fmt"
	"slices"

	"github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) treeRequirementBSinDiagram(diagram *Diagram, requirement *Requirement, parentNode *tree.Node) {
	confNode := TreeNodeAndShapeConfigurationWithoutLink[
		*Requirement, Requirement,
		*Requirement, Requirement, // Parent is not used
		*RequirementShape, RequirementShape,
		*Diagram,
	]{
		diagram:                     diagram,
		parentNode:                  parentNode,
		element:                     requirement,
		parentElement:               nil,
		elementsWhoseNodeIsExpanded: &diagram.RequirementsWhoseNodeIsExpanded,
		shapes:                      &diagram.Requirement_Shapes,
		shapesMap:                   diagram.map_Requirement_RequirementShape,
	}
	requirementNode := addNodeToTreeWithoutLink(stager, confNode)

	if len(requirement.Concerns) > 0 {
		concernsNode := &tree.Node{
			Name:                 fmt.Sprintf("Concerns (%d)", len(requirement.Concerns)),
			IsExpanded:           slices.Index(diagram.RequirementsWhoseConcernsNodeIsExpanded, requirement) != -1,
			IsNodeClickable:      true,
			IsWithPreceedingIcon: true,
			PreceedingIcon:       string(buttons.BUTTON_output),
		}
		requirementNode.Children = append(requirementNode.Children, concernsNode)

		setCallbacksExpandableNode(stager, concernsNode, requirement, &diagram.RequirementsWhoseConcernsNodeIsExpanded)

		for _, concern := range requirement.Concerns {
			concernNode := &tree.Node{
				Name:                    concern.GetName(),
				IsExpanded:              true,
				IsNodeClickable:         true,
				IsInEditMode:            concern.GetIsInRenameMode(),
				CheckboxHasToolTip:      true,
				CheckboxToolTipPosition: tree.Right,
			}
			concernsNode.Children = append(concernsNode.Children, concernNode)
			addRenameButton(concern, concernNode, stager)
			concernNode.OnNameChange = func(newName string) {
				concern.SetName(newName)
				concern.SetIsInRenameMode(false)
				stager.stage.Commit()
			}

			if _, ok := diagram.map_Requirement_RequirementShape[requirement]; ok {
				if _, ok := diagram.map_Concern_ConcernShape[concern]; ok {

					concernNode.HasCheckboxButton = true

					key := requirementConcernKey{
						Requirement: requirement,
						Concern:     concern,
					}
					associationShape, ok := diagram.map_RequirementConcernKey_RequirementConcernShape[key]
					concernNode.IsChecked = ok

					if ok {
						concernNode.CheckboxToolTipText = "Uncheck to remove shape from diagram"
					} else {
						concernNode.CheckboxToolTipText = "Check to add shape to diagram"
					}

					concernNode.OnIsCheckedChanged = func(isChecked bool) {
						if isChecked {
							addAssociationShapeToDiagram(stager, requirement, concern, &diagram.RequirementConcernShapes)
							stager.stage.Commit()
						} else {
							associationShape.UnstageVoid(stager.stage)
							stager.stage.Commit()
						}
					}

					concernNode.Buttons = []*tree.Button{
						{
							Name:            diagram.GetName(),
							Icon:            string(buttons.BUTTON_visibility_off),
							ToolTipText:     "Hide link from diagram",
							HasToolTip:      true,
							ToolTipPosition: tree.Right,
							OnClick: func() {
								associationShape.SetIsHidden(!associationShape.GetIsHidden())
								stager.stage.Commit()
							},
						},
					}
					if ok {
						if associationShape.GetIsHidden() {
							concernNode.Buttons[0].Icon = string(buttons.BUTTON_visibility)
							concernNode.Buttons[0].ToolTipText = "Show link on diagram"
						}
					} else {
						concernNode.Buttons[0].IsDisabled = true
					}
				}
			}
		}
	}

	if len(requirement.Views) > 0 {
		viewsNode := &tree.Node{
			Name:                 fmt.Sprintf("Views (%d)", len(requirement.Views)),
			IsExpanded:           slices.Index(diagram.RequirementsWhoseViewsNodeIsExpanded, requirement) != -1,
			IsNodeClickable:      true,
			IsWithPreceedingIcon: true,
			PreceedingIcon:       string(buttons.BUTTON_output),
		}
		requirementNode.Children = append(requirementNode.Children, viewsNode)

		setCallbacksExpandableNode(stager, viewsNode, requirement, &diagram.RequirementsWhoseViewsNodeIsExpanded)

		for _, view := range requirement.Views {
			viewNode := &tree.Node{
				Name:                    view.GetName(),
				IsExpanded:              true,
				IsNodeClickable:         true,
				IsInEditMode:            view.GetIsInRenameMode(),
				CheckboxHasToolTip:      true,
				CheckboxToolTipPosition: tree.Right,
			}
			viewsNode.Children = append(viewsNode.Children, viewNode)
			addRenameButton(view, viewNode, stager)
			viewNode.OnNameChange = func(newName string) {
				view.SetName(newName)
				view.SetIsInRenameMode(false)
				stager.stage.Commit()
			}

			if _, ok := diagram.map_Requirement_RequirementShape[requirement]; ok {
				if _, ok := diagram.map_View_ViewShape[view]; ok {

					viewNode.HasCheckboxButton = true

					key := requirementViewKey{
						Requirement: requirement,
						View:        view,
					}
					associationShape, ok := diagram.map_RequirementViewKey_RequirementViewShape[key]
					viewNode.IsChecked = ok

					if ok {
						viewNode.CheckboxToolTipText = "Uncheck to remove shape from diagram"
					} else {
						viewNode.CheckboxToolTipText = "Check to add shape to diagram"
					}

					viewNode.OnIsCheckedChanged = func(isChecked bool) {
						if isChecked {
							addAssociationShapeToDiagram(stager, requirement, view, &diagram.RequirementViewShapes)
							stager.stage.Commit()
						} else {
							associationShape.UnstageVoid(stager.stage)
							stager.stage.Commit()
						}
					}

					viewNode.Buttons = []*tree.Button{
						{
							Name:            diagram.GetName(),
							Icon:            string(buttons.BUTTON_visibility_off),
							ToolTipText:     "Hide link from diagram",
							HasToolTip:      true,
							ToolTipPosition: tree.Right,
							OnClick: func() {
								associationShape.SetIsHidden(!associationShape.GetIsHidden())
								stager.stage.Commit()
							},
						},
					}
					if ok {
						if associationShape.GetIsHidden() {
							viewNode.Buttons[0].Icon = string(buttons.BUTTON_visibility)
							viewNode.Buttons[0].ToolTipText = "Show link on diagram"
						}
					} else {
						viewNode.Buttons[0].IsDisabled = true
					}
				}
			}
		}
	}

	if len(requirement.Concepts) > 0 {
		conceptsNode := &tree.Node{
			Name:                 fmt.Sprintf("Concepts (%d)", len(requirement.Concepts)),
			IsExpanded:           slices.Index(diagram.RequirementsWhoseConceptsNodeIsExpanded, requirement) != -1,
			IsNodeClickable:      true,
			IsWithPreceedingIcon: true,
			PreceedingIcon:       string(buttons.BUTTON_output),
		}
		requirementNode.Children = append(requirementNode.Children, conceptsNode)

		setCallbacksExpandableNode(stager, conceptsNode, requirement, &diagram.RequirementsWhoseConceptsNodeIsExpanded)

		for _, concept := range requirement.Concepts {
			conceptNode := &tree.Node{
				Name:                    concept.GetName(),
				IsExpanded:              true,
				IsNodeClickable:         true,
				IsInEditMode:            concept.GetIsInRenameMode(),
				CheckboxHasToolTip:      true,
				CheckboxToolTipPosition: tree.Right,
			}
			conceptsNode.Children = append(conceptsNode.Children, conceptNode)
			addRenameButton(concept, conceptNode, stager)
			conceptNode.OnNameChange = func(newName string) {
				concept.SetName(newName)
				concept.SetIsInRenameMode(false)
				stager.stage.Commit()
			}

			if _, ok := diagram.map_Requirement_RequirementShape[requirement]; ok {
				if _, ok := diagram.map_Concept_ConceptShape[concept]; ok {

					conceptNode.HasCheckboxButton = true

					key := requirementConceptKey{
						Requirement: requirement,
						Concept:     concept,
					}
					associationShape, ok := diagram.map_RequirementConceptKey_RequirementConceptShape[key]
					conceptNode.IsChecked = ok

					if ok {
						conceptNode.CheckboxToolTipText = "Uncheck to remove shape from diagram"
					} else {
						conceptNode.CheckboxToolTipText = "Check to add shape to diagram"
					}

					conceptNode.OnIsCheckedChanged = func(isChecked bool) {
						if isChecked {
							addAssociationShapeToDiagram(stager, requirement, concept, &diagram.RequirementConceptShapes)
							stager.stage.Commit()
						} else {
							associationShape.UnstageVoid(stager.stage)
							stager.stage.Commit()
						}
					}

					conceptNode.Buttons = []*tree.Button{
						{
							Name:            diagram.GetName(),
							Icon:            string(buttons.BUTTON_visibility_off),
							ToolTipText:     "Hide link from diagram",
							HasToolTip:      true,
							ToolTipPosition: tree.Right,
							OnClick: func() {
								associationShape.SetIsHidden(!associationShape.GetIsHidden())
								stager.stage.Commit()
							},
						},
					}
					if ok {
						if associationShape.GetIsHidden() {
							conceptNode.Buttons[0].Icon = string(buttons.BUTTON_visibility)
							conceptNode.Buttons[0].ToolTipText = "Show link on diagram"
						}
					} else {
						conceptNode.Buttons[0].IsDisabled = true
					}
				}
			}
		}
	}
}
