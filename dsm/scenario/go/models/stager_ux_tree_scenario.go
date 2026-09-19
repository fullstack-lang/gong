package models

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"

	"github.com/fullstack-lang/gong/dsm/scenario/go/icons"
)

func (stager *Stager) treeScenario(scenario *Scenario, parentNode *tree.Node) {
	scenarioNode := new(tree.Node)
	scenarioNode.Name = scenario.Name
	scenarioNode.IsExpanded = scenario.GetIsExpanded()
	scenarioNode.IsWithPreceedingIcon = false
	scenarioNode.PreceedingSVGIcon = icons.ScenarioIcon
	scenarioNode.IsNodeClickable = true
	scenarioNode.IsInEditMode = scenario.GetIsInRenameMode()

	parentNode.Children = append(parentNode.Children, scenarioNode)

	scenarioNode.OnClick = onNodeClicked(stager, scenario)
	scenarioNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&scenario.IsExpanded)
	scenarioNode.OnNameChange = stager.onNameChange(scenario)
	addRenameButton(scenario, scenarioNode, stager)

	// Categories
	stager.treeCategoryDiagrams(scenario, &scenarioNode.Children)
}

// -------------------------------------------------------------------------------------
// Diagrams
// -------------------------------------------------------------------------------------
func (stager *Stager) treeCategoryDiagrams(scenario *Scenario, parentNodes *[]*tree.Node) {
	categoryNode := new(tree.Node)
	categoryNode.Name = "Diagrams"
	categoryNode.FontStyle = tree.ITALIC
	categoryNode.IsExpanded = scenario.IsDiagramsNodeExpanded
	categoryNode.IsNodeClickable = true
	categoryNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&scenario.IsDiagramsNodeExpanded)
	categoryNode.OnClick = onNodeClicked(stager, scenario)
	*parentNodes = append(*parentNodes, categoryNode)

	confDiagrams := ItemButtonConfiguration[
		Diagram, *Diagram,
		Scenario, *Scenario,
	]{
		parentNode:                         categoryNode,
		sliceForNewAddedItem:               &scenario.Diagrams,
		isParentNodeExpandedByAddOperation: true,
		parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
		parentNodeExpansionBooleanValue:    &scenario.IsDiagramsNodeExpanded,
	}
	itemAdderCallback := addCreateItemButton(stager, confDiagrams)
	itemAdderCallback.OnBeforeCommit = func() {
		newDiagram := itemAdderCallback.createdItem
		newDiagram.IsExpanded = true
		newDiagram.AxisOrign_X = 150
		newDiagram.AxisOrign_Y = 300
		newDiagram.HorizontalAxis_Right_X = 950
		newDiagram.VerticalAxis_Top_Y = 100
		newDiagram.VerticalAxis_Bottom_Y = 500
		newDiagram.VerticalAxis_StrokeWidth = 2
		newDiagram.Start = time.Now()
		newDiagram.End = time.Now().AddDate(1, 0, 0)
		newDiagram.IsParametersNodeExpanded = true
		newDiagram.IsActorStatesNodeExpanded = true
		newDiagram.IsEvolutionDirectionsNodeExpanded = true
		newDiagram.IsParametersAggregatesNodeExpanded = true
		newDiagram.IsActorStateTransitionsNodeExpanded = true

		for diagram_ := range *stager.stage.GetInstancesSet[*Diagram]() {
			diagram_.IsChecked = false
		}
		newDiagram.IsChecked = true
		newDiagram.IsInDrawMode = true
	}

	slices.SortFunc(scenario.Diagrams, GongCompareGongstructByName)
	for _, diagram := range scenario.Diagrams {
		diagramNode := new(tree.Node)
		diagramNode.Name = diagram.Name
		diagramNode.HasCheckboxButton = true
		diagramNode.IsChecked = diagram.IsChecked
		diagramNode.IsNodeClickable = true
		diagramNode.IsExpanded = diagram.GetIsExpanded()
		diagramNode.OnIsCheckedChanged = onIsCheckedChangedDiagram(stager, diagram)
		diagramNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&diagram.IsExpanded)
		diagramNode.OnClick = onNodeClicked(stager, diagram)

		diagramNode.IsInEditMode = diagram.GetIsInRenameMode()
		diagramNode.OnNameChange = stager.onNameChange(diagram)
		addRenameButton(diagram, diagramNode, stager)

		// edit button
		{
			editButton := &tree.Button{
				Name:            "Diagram Editability",
				Icon:            string(buttons.BUTTON_edit),
				HasToolTip:      true,
				ToolTipPosition: tree.Above,
				OnClick: func() {
					diagram.IsInDrawMode = !diagram.IsInDrawMode
					stager.stage.Commit()
				},
			}
			if !diagram.IsInDrawMode {
				editButton.Icon = string(buttons.BUTTON_edit)
				editButton.ToolTipText = "Edit diagram"
			} else {
				editButton.Icon = string(buttons.BUTTON_edit_off)
				editButton.ToolTipText = "Stop editing diagram"
			}
			diagramNode.Buttons = append(diagramNode.Buttons, editButton)
		}

		// prefix button
		{
			showPrefixButton := &tree.Button{
				Name:            "Diagram Prefix",
				Icon:            string(buttons.BUTTON_show_chart),
				HasToolTip:      true,
				ToolTipPosition: tree.Above,
				OnClick: func() {
					diagram.IsShowPrefix = !diagram.IsShowPrefix
					stager.stage.Commit()
				},
			}
			if !diagram.IsShowPrefix {
				showPrefixButton.Icon = string(buttons.BUTTON_label)
				showPrefixButton.ToolTipText = "Show Prefix"
			} else {
				showPrefixButton.Icon = string(buttons.BUTTON_label_off)
				showPrefixButton.ToolTipText = "Hide Prefix"
			}
			diagramNode.Buttons = append(diagramNode.Buttons, showPrefixButton)
		}

		categoryNode.Children = append(categoryNode.Children, diagramNode)

		// Create maps for the generic functions
		map_Parameter_ParameterShape := make(map[*Parameter]*ParameterShape)
		for _, shape := range diagram.ParameterShapes {
			if shape.Parameter != nil {
				map_Parameter_ParameterShape[shape.Parameter] = shape
			}
		}

		map_ActorState_ActorStateShape := make(map[*ActorState]*ActorStateShape)
		for _, shape := range diagram.ActorStateShapes {
			if shape.ActorState != nil {
				map_ActorState_ActorStateShape[shape.ActorState] = shape
			}
		}

		map_EvolutionDirection_EvolutionDirectionShape := make(map[*EvolutionDirection]*EvolutionDirectionShape)
		for _, shape := range diagram.EvolutionDirectionShapes {
			if shape.EvolutionDirection != nil {
				map_EvolutionDirection_EvolutionDirectionShape[shape.EvolutionDirection] = shape
			}
		}

		map_ParametersAggregate_ParametersAggregateShape := make(map[*ParametersAggregate]*ParametersAggregateShape)
		for _, shape := range diagram.ScenarioParameterShapes {
			if shape.ScenarioParameter != nil {
				map_ParametersAggregate_ParametersAggregateShape[shape.ScenarioParameter] = shape
			}
		}

		map_ActorStateTransition_ActorStateTransitionShape := make(map[*ActorStateTransition]*ActorStateTransitionShape)
		for _, shape := range diagram.ActorStateTransitionShapes {
			if shape.ActorStateTransition != nil {
				map_ActorStateTransition_ActorStateTransitionShape[shape.ActorStateTransition] = shape
			}
		}

		// Dummy slices for elements whose node is expanded, since these are leaves without persistent expansion state
		var dummyExpandedParameters []*Parameter
		var dummyExpandedActorStates []*ActorState
		var dummyExpandedEvolutionDirections []*EvolutionDirection
		var dummyExpandedParametersAggregates []*ParametersAggregate
		var dummyExpandedActorStateTransitions []*ActorStateTransition

		// Parameters Node
		parametersNode := new(tree.Node)
		parametersNode.Name = "Parameters"
		parametersNode.FontStyle = tree.ITALIC
		parametersNode.IsExpanded = diagram.IsParametersNodeExpanded
		parametersNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&diagram.IsParametersNodeExpanded)
		parametersNode.IsNodeClickable = true
		parametersNode.OnClick = onNodeClicked(stager, diagram)
		diagramNode.Children = append(diagramNode.Children, parametersNode)

		confParameters := ItemButtonConfiguration[
			Parameter, *Parameter,
			Scenario, *Scenario,
		]{
			parentNode:                         parametersNode,
			sliceForNewAddedItem:               &scenario.Parameters,
			isParentNodeExpandedByAddOperation: true,
			parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
			parentNodeExpansionBooleanValue:    &diagram.IsParametersNodeExpanded,
		}
		addCreateItemButton(stager, confParameters)

		slices.SortFunc(scenario.Parameters, GongCompareGongstructByName)
		for _, p := range scenario.Parameters {
			conf := TreeNodeAndShapeConfigurationWithoutLink[
				*Parameter, Parameter, // AT, AT_
				*Parameter, Parameter, // ParentAT, ParentAT_ (dummy)
				*ParameterShape, ParameterShape, // CT, CT_
				*Diagram, // DiagramType
			]{
				diagram:                     diagram,
				parentNode:                  parametersNode,
				element:                     p,
				parentElement:               nil,
				elementsWhoseNodeIsExpanded: &dummyExpandedParameters,
				shapes:                      &diagram.ParameterShapes,
				shapesMap:                   map_Parameter_ParameterShape,
			}
			addNodeToTreeWithoutLink(stager, conf)
		}

		// Actor States Node
		actorStatesNode := new(tree.Node)
		actorStatesNode.Name = "Actor States"
		actorStatesNode.FontStyle = tree.ITALIC
		actorStatesNode.IsExpanded = diagram.IsActorStatesNodeExpanded
		actorStatesNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&diagram.IsActorStatesNodeExpanded)
		actorStatesNode.IsNodeClickable = true
		actorStatesNode.OnClick = onNodeClicked(stager, diagram)
		diagramNode.Children = append(diagramNode.Children, actorStatesNode)

		confActorStates := ItemButtonConfiguration[
			ActorState, *ActorState,
			Scenario, *Scenario,
		]{
			parentNode:                         actorStatesNode,
			sliceForNewAddedItem:               &scenario.ActorStates,
			isParentNodeExpandedByAddOperation: true,
			parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
			parentNodeExpansionBooleanValue:    &diagram.IsActorStatesNodeExpanded,
		}
		addCreateItemButton(stager, confActorStates)

		slices.SortFunc(scenario.ActorStates, GongCompareGongstructByName)
		for _, as := range scenario.ActorStates {
			conf := TreeNodeAndShapeConfigurationWithoutLink[
				*ActorState, ActorState,
				*ActorState, ActorState,
				*ActorStateShape, ActorStateShape,
				*Diagram,
			]{
				diagram:                     diagram,
				parentNode:                  actorStatesNode,
				element:                     as,
				parentElement:               nil,
				elementsWhoseNodeIsExpanded: &dummyExpandedActorStates,
				shapes:                      &diagram.ActorStateShapes,
				shapesMap:                   map_ActorState_ActorStateShape,
			}
			addNodeToTreeWithoutLink(stager, conf)
		}

		// Evolution Directions Node
		evolutionDirectionsNode := new(tree.Node)
		evolutionDirectionsNode.Name = "Evolution Directions"
		evolutionDirectionsNode.FontStyle = tree.ITALIC
		evolutionDirectionsNode.IsExpanded = diagram.IsEvolutionDirectionsNodeExpanded
		evolutionDirectionsNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&diagram.IsEvolutionDirectionsNodeExpanded)
		evolutionDirectionsNode.IsNodeClickable = true
		evolutionDirectionsNode.OnClick = onNodeClicked(stager, diagram)
		diagramNode.Children = append(diagramNode.Children, evolutionDirectionsNode)

		confEvolutionDirections := ItemButtonConfiguration[
			EvolutionDirection, *EvolutionDirection,
			Scenario, *Scenario,
		]{
			parentNode:                         evolutionDirectionsNode,
			sliceForNewAddedItem:               &scenario.EvolutionDirections,
			isParentNodeExpandedByAddOperation: true,
			parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
			parentNodeExpansionBooleanValue:    &diagram.IsEvolutionDirectionsNodeExpanded,
		}
		addCreateItemButton(stager, confEvolutionDirections)

		slices.SortFunc(scenario.EvolutionDirections, GongCompareGongstructByName)
		for _, ed := range scenario.EvolutionDirections {
			conf := TreeNodeAndShapeConfigurationWithoutLink[
				*EvolutionDirection, EvolutionDirection,
				*EvolutionDirection, EvolutionDirection,
				*EvolutionDirectionShape, EvolutionDirectionShape,
				*Diagram,
			]{
				diagram:                     diagram,
				parentNode:                  evolutionDirectionsNode,
				element:                     ed,
				parentElement:               nil,
				elementsWhoseNodeIsExpanded: &dummyExpandedEvolutionDirections,
				shapes:                      &diagram.EvolutionDirectionShapes,
				shapesMap:                   map_EvolutionDirection_EvolutionDirectionShape,
			}
			addNodeToTreeWithoutLink(stager, conf)
		}

		// Parameters Aggregates Node
		parametersAggregatesNode := new(tree.Node)
		parametersAggregatesNode.Name = "Parameters Aggregates"
		parametersAggregatesNode.FontStyle = tree.ITALIC
		parametersAggregatesNode.IsExpanded = diagram.IsParametersAggregatesNodeExpanded
		parametersAggregatesNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&diagram.IsParametersAggregatesNodeExpanded)
		parametersAggregatesNode.IsNodeClickable = true
		parametersAggregatesNode.OnClick = onNodeClicked(stager, diagram)
		diagramNode.Children = append(diagramNode.Children, parametersAggregatesNode)

		confParametersAggregates := ItemButtonConfiguration[
			ParametersAggregate, *ParametersAggregate,
			Scenario, *Scenario,
		]{
			parentNode:                         parametersAggregatesNode,
			sliceForNewAddedItem:               &scenario.ParametersAggretates,
			isParentNodeExpandedByAddOperation: true,
			parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
			parentNodeExpansionBooleanValue:    &diagram.IsParametersAggregatesNodeExpanded,
		}
		addCreateItemButton(stager, confParametersAggregates)

		slices.SortFunc(scenario.ParametersAggretates, GongCompareGongstructByName)
		for _, pa := range scenario.ParametersAggretates {
			conf := TreeNodeAndShapeConfigurationWithoutLink[
				*ParametersAggregate, ParametersAggregate,
				*ParametersAggregate, ParametersAggregate,
				*ParametersAggregateShape, ParametersAggregateShape,
				*Diagram,
			]{
				diagram:                     diagram,
				parentNode:                  parametersAggregatesNode,
				element:                     pa,
				parentElement:               nil,
				elementsWhoseNodeIsExpanded: &dummyExpandedParametersAggregates,
				shapes:                      &diagram.ScenarioParameterShapes,
				shapesMap:                   map_ParametersAggregate_ParametersAggregateShape,
			}
			addNodeToTreeWithoutLink(stager, conf)
		}

		// Actor State Transitions Node
		actorStateTransitionsNode := new(tree.Node)
		actorStateTransitionsNode.Name = "Actor State Transitions"
		actorStateTransitionsNode.FontStyle = tree.ITALIC
		actorStateTransitionsNode.IsExpanded = diagram.IsActorStateTransitionsNodeExpanded
		actorStateTransitionsNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&diagram.IsActorStateTransitionsNodeExpanded)
		actorStateTransitionsNode.IsNodeClickable = true
		actorStateTransitionsNode.OnClick = onNodeClicked(stager, diagram)
		diagramNode.Children = append(diagramNode.Children, actorStateTransitionsNode)

		confActorStateTransitions := ItemButtonConfiguration[
			ActorStateTransition, *ActorStateTransition,
			Scenario, *Scenario,
		]{
			parentNode:                         actorStateTransitionsNode,
			sliceForNewAddedItem:               &scenario.ActorStateTransitions,
			isParentNodeExpandedByAddOperation: true,
			parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
			parentNodeExpansionBooleanValue:    &diagram.IsActorStateTransitionsNodeExpanded,
		}
		addCreateItemButton(stager, confActorStateTransitions)

		slices.SortFunc(scenario.ActorStateTransitions, GongCompareGongstructByName)
		for _, ast := range scenario.ActorStateTransitions {
			conf := TreeNodeAndShapeConfigurationWithoutLink[
				*ActorStateTransition, ActorStateTransition,
				*ActorStateTransition, ActorStateTransition,
				*ActorStateTransitionShape, ActorStateTransitionShape,
				*Diagram,
			]{
				diagram:                     diagram,
				parentNode:                  actorStateTransitionsNode,
				element:                     ast,
				parentElement:               nil,
				elementsWhoseNodeIsExpanded: &dummyExpandedActorStateTransitions,
				shapes:                      &diagram.ActorStateTransitionShapes,
				shapesMap:                   map_ActorStateTransition_ActorStateTransitionShape,
			}
			addNodeToTreeWithoutLink(stager, conf)
		}
	}
}

func onIsCheckedChangedDiagram(stager *Stager, diagram *Diagram) func(isChecked bool) {
	return func(isChecked bool) {
		if isChecked {
			for diagram_ := range *stager.stage.GetInstancesSet[*Diagram]() {
				diagram_.IsChecked = false
			}
			diagram.IsChecked = true
			stager.stage.Commit()
		} else {
			diagram.IsChecked = false
			stager.stage.Commit()
		}
	}
}
