package models

import (
	"cmp"
	"fmt"
	"slices"

	"github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) treeViewBSinDiagram(diagram *Diagram, view *View, parentNode *tree.Node) {
	confNode := TreeNodeAndShapeConfigurationWithoutLink[
		*View, View,
		*View, View, // Parent is not used
		*ViewShape, ViewShape,
		*Diagram,
	]{
		diagram:                     diagram,
		parentNode:                  parentNode,
		element:                     view,
		parentElement:               nil,
		elementsWhoseNodeIsExpanded: &diagram.ViewsWhoseNodeIsExpanded,
		shapes:                      &diagram.View_Shapes,
		shapesMap:                   diagram.map_View_ViewShape,
	}
	viewNode := addNodeToTreeWithoutLink(stager, confNode)

	var deliverables []*Deliverable
	for deliverable := range *stager.stage.GetInstancesSet[*Deliverable]() {
		if slices.Contains(deliverable.Views, view) {
			deliverables = append(deliverables, deliverable)
		}
	}

	if len(deliverables) > 0 {
		slices.SortFunc(deliverables, func(a, b *Deliverable) int {
			return cmp.Compare(a.Name, b.Name)
		})

		deliverablesNode := &tree.Node{
			Name:                 fmt.Sprintf("Deliverables (%d)", len(deliverables)),
			IsExpanded:           slices.Index(diagram.ViewsWhoseDeliverablesNodeIsExpanded, view) != -1,
			IsNodeClickable:      true,
			IsWithPreceedingIcon: true,
			PreceedingIcon:       string(buttons.BUTTON_input),
		}
		viewNode.Children = append(viewNode.Children, deliverablesNode)

		setCallbacksExpandableNode(stager, deliverablesNode, view, &diagram.ViewsWhoseDeliverablesNodeIsExpanded)

		for _, deliverable := range deliverables {
			deliverableNode := &tree.Node{
				Name:                    deliverable.GetName(),
				IsExpanded:              true,
				IsNodeClickable:         true,
				IsInEditMode:            deliverable.GetIsInRenameMode(),
				CheckboxHasToolTip:      true,
				CheckboxToolTipPosition: tree.Right,
			}
			deliverablesNode.Children = append(deliverablesNode.Children, deliverableNode)
			addRenameButton(deliverable, deliverableNode, stager)
			deliverableNode.OnNameChange = func(newName string) {
				deliverable.SetName(newName)
				deliverable.SetIsInRenameMode(false)
				stager.stage.Commit()
			}

			if _, ok := diagram.map_Deliverable_DeliverableShape[deliverable]; ok {
				if _, ok := diagram.map_View_ViewShape[view]; ok {

					deliverableNode.HasCheckboxButton = true

					key := deliverableViewKey{
						Deliverable: deliverable,
						View:        view,
					}
					associationShape, ok := diagram.map_DeliverableViewKey_DeliverableViewShape[key]
					deliverableNode.IsChecked = ok

					if ok {
						deliverableNode.CheckboxToolTipText = "Uncheck to remove shape from diagram"
					} else {
						deliverableNode.CheckboxToolTipText = "Check to add shape to diagram"
					}

					deliverableNode.OnIsCheckedChanged = func(isChecked bool) {
						if isChecked {
							addAssociationShapeToDiagram(stager, deliverable, view, &diagram.DeliverableViewShapes)
							stager.stage.Commit()
						} else {
							associationShape.UnstageVoid(stager.stage)
							stager.stage.Commit()
						}
					}

					deliverableNode.Buttons = []*tree.Button{
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
							deliverableNode.Buttons[0].Icon = string(buttons.BUTTON_visibility)
							deliverableNode.Buttons[0].ToolTipText = "Show link on diagram"
						}
					} else {
						deliverableNode.Buttons[0].IsDisabled = true
					}
				}
			}
		}
	}

	if len(view.Concepts) > 0 {
		conceptsNode := &tree.Node{
			Name:                 fmt.Sprintf("Concepts (%d)", len(view.Concepts)),
			IsExpanded:           slices.Index(diagram.ViewsWhoseConceptsNodeIsExpanded, view) != -1,
			IsNodeClickable:      true,
			IsWithPreceedingIcon: true,
			PreceedingIcon:       string(buttons.BUTTON_output),
		}
		viewNode.Children = append(viewNode.Children, conceptsNode)

		setCallbacksExpandableNode(stager, conceptsNode, view, &diagram.ViewsWhoseConceptsNodeIsExpanded)

		for _, concept := range view.Concepts {
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

			if _, ok := diagram.map_View_ViewShape[view]; ok {
				if _, ok := diagram.map_Concept_ConceptShape[concept]; ok {

					conceptNode.HasCheckboxButton = true

					key := viewConceptKey{
						View:    view,
						Concept: concept,
					}
					associationShape, ok := diagram.map_ViewConceptKey_ViewConceptShape[key]
					conceptNode.IsChecked = ok

					if ok {
						conceptNode.CheckboxToolTipText = "Uncheck to remove shape from diagram"
					} else {
						conceptNode.CheckboxToolTipText = "Check to add shape to diagram"
					}

					conceptNode.OnIsCheckedChanged = func(isChecked bool) {
						if isChecked {
							addAssociationShapeToDiagram(stager, view, concept, &diagram.ViewConceptShapes)
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

	var requirements []*Requirement
	for requirement := range *stager.stage.GetInstancesSet[*Requirement]() {
		if slices.Contains(requirement.Views, view) {
			requirements = append(requirements, requirement)
		}
	}

	if len(requirements) > 0 {
		slices.SortFunc(requirements, func(a, b *Requirement) int {
			return cmp.Compare(a.Name, b.Name)
		})

		requirementsNode := &tree.Node{
			Name:                 fmt.Sprintf("Requirements (%d)", len(requirements)),
			IsExpanded:           slices.Index(diagram.ViewsWhoseRequirementsNodeIsExpanded, view) != -1,
			IsNodeClickable:      true,
			IsWithPreceedingIcon: true,
			PreceedingIcon:       string(buttons.BUTTON_input),
		}
		viewNode.Children = append(viewNode.Children, requirementsNode)

		setCallbacksExpandableNode(stager, requirementsNode, view, &diagram.ViewsWhoseRequirementsNodeIsExpanded)

		for _, requirement := range requirements {
			requirementNode := &tree.Node{
				Name:                    requirement.GetName(),
				IsExpanded:              true,
				IsNodeClickable:         true,
				IsInEditMode:            requirement.GetIsInRenameMode(),
				CheckboxHasToolTip:      true,
				CheckboxToolTipPosition: tree.Right,
			}
			requirementsNode.Children = append(requirementsNode.Children, requirementNode)
			addRenameButton(requirement, requirementNode, stager)
			requirementNode.OnNameChange = func(newName string) {
				requirement.SetName(newName)
				requirement.SetIsInRenameMode(false)
				stager.stage.Commit()
			}

			if _, ok := diagram.map_Requirement_RequirementShape[requirement]; ok {
				if _, ok := diagram.map_View_ViewShape[view]; ok {

					requirementNode.HasCheckboxButton = true

					key := requirementViewKey{
						Requirement: requirement,
						View:        view,
					}
					associationShape, ok := diagram.map_RequirementViewKey_RequirementViewShape[key]
					requirementNode.IsChecked = ok

					if ok {
						requirementNode.CheckboxToolTipText = "Uncheck to remove shape from diagram"
					} else {
						requirementNode.CheckboxToolTipText = "Check to add shape to diagram"
					}

					requirementNode.OnIsCheckedChanged = func(isChecked bool) {
						if isChecked {
							addAssociationShapeToDiagram(stager, requirement, view, &diagram.RequirementViewShapes)
							stager.stage.Commit()
						} else {
							associationShape.UnstageVoid(stager.stage)
							stager.stage.Commit()
						}
					}

					requirementNode.Buttons = []*tree.Button{
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
							requirementNode.Buttons[0].Icon = string(buttons.BUTTON_visibility)
							requirementNode.Buttons[0].ToolTipText = "Show link on diagram"
						}
					} else {
						requirementNode.Buttons[0].IsDisabled = true
					}
				}
			}
		}
	}
}

