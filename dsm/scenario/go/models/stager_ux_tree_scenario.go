package models

import (
	"slices"

	tree "github.com/fullstack-lang/gong/lib/tree/go/models"

	"github.com/fullstack-lang/gong/dsm/scenario/go/icons"
)

func (stager *Stager) treeScenario(treeInstance *tree.Tree, scenario *Scenario, parentNodes *[]*tree.Node) {
	scenarioNode := new(tree.Node)
	scenarioNode.Name = scenario.Name
	scenarioNode.IsExpanded = scenario.GetIsExpanded()
	scenarioNode.IsWithPreceedingIcon = false
	scenarioNode.PreceedingSVGIcon = icons.ScenarioIcon
	scenarioNode.IsNodeClickable = true

	*parentNodes = append(*parentNodes, scenarioNode)

	// Callback for when the scenario node is clicked
	scenarioNode.OnClick = func(frontNode *tree.Node) {
		stager.probeForm.FillUpFormFromGongstruct(scenario, GetPointerToGongstructName[*Scenario]())
		stager.stage.Commit()
	}
	scenarioNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&scenario.IsExpanded)
	scenarioNode.IsInEditMode = scenario.GetIsInRenameMode()
	scenarioNode.OnNameChange = stager.onNameChange(scenario)
	addRenameButton(scenario, scenarioNode, stager)

	// Categories
	stager.treeCategoryDiagrams(treeInstance, scenario, &scenarioNode.Children)
}

// -------------------------------------------------------------------------------------
// Diagrams
// -------------------------------------------------------------------------------------
func (stager *Stager) treeCategoryDiagrams(treeInstance *tree.Tree, scenario *Scenario, parentNodes *[]*tree.Node) {
	categoryNode := new(tree.Node)
	categoryNode.Name = "Diagrams"
	categoryNode.IsExpanded = scenario.IsDiagramsNodeExpanded
	categoryNode.IsNodeClickable = false
	categoryNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&scenario.IsDiagramsNodeExpanded)
	*parentNodes = append(*parentNodes, categoryNode)

	slices.SortFunc(scenario.Diagrams, CompareGongstructByName)
	for _, diagram := range scenario.Diagrams {
		diagramNode := new(tree.Node)
		diagramNode.Name = diagram.Name
		diagramNode.HasCheckboxButton = true
		diagramNode.IsChecked = diagram.IsChecked
		diagramNode.IsNodeClickable = true
		diagramNode.IsExpanded = diagram.GetIsExpanded()
		diagramNode.OnIsCheckedChanged = onIsCheckedChangedDiagram(stager, diagram)
		diagramNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&diagram.IsExpanded)
		diagramNode.OnClick = func(frontNode *tree.Node) {
			stager.probeForm.FillUpFormFromGongstruct(diagram, GetPointerToGongstructName[*Diagram]())
			stager.stage.Commit()
		}

		diagramNode.IsInEditMode = diagram.GetIsInRenameMode()
		diagramNode.OnNameChange = stager.onNameChange(diagram)
		addRenameButton(diagram, diagramNode, stager)

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
		parametersNode.IsExpanded = diagram.IsParametersNodeExpanded
		parametersNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&diagram.IsParametersNodeExpanded)
		parametersNode.IsNodeClickable = false
		diagramNode.Children = append(diagramNode.Children, parametersNode)

		slices.SortFunc(scenario.Parameters, CompareGongstructByName)
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
		actorStatesNode.IsExpanded = diagram.IsActorStatesNodeExpanded
		actorStatesNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&diagram.IsActorStatesNodeExpanded)
		actorStatesNode.IsNodeClickable = false
		diagramNode.Children = append(diagramNode.Children, actorStatesNode)

		slices.SortFunc(scenario.ActorStates, CompareGongstructByName)
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
		evolutionDirectionsNode.IsExpanded = diagram.IsEvolutionDirectionsNodeExpanded
		evolutionDirectionsNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&diagram.IsEvolutionDirectionsNodeExpanded)
		evolutionDirectionsNode.IsNodeClickable = false
		diagramNode.Children = append(diagramNode.Children, evolutionDirectionsNode)

		slices.SortFunc(scenario.EvolutionDirections, CompareGongstructByName)
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
		parametersAggregatesNode.IsExpanded = diagram.IsParametersAggregatesNodeExpanded
		parametersAggregatesNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&diagram.IsParametersAggregatesNodeExpanded)
		parametersAggregatesNode.IsNodeClickable = false
		diagramNode.Children = append(diagramNode.Children, parametersAggregatesNode)

		slices.SortFunc(scenario.ParametersAggretates, CompareGongstructByName)
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
		actorStateTransitionsNode.IsExpanded = diagram.IsActorStateTransitionsNodeExpanded
		actorStateTransitionsNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&diagram.IsActorStateTransitionsNodeExpanded)
		actorStateTransitionsNode.IsNodeClickable = false
		diagramNode.Children = append(diagramNode.Children, actorStateTransitionsNode)

		slices.SortFunc(scenario.ActorStateTransitions, CompareGongstructByName)
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
			for diagram_ := range *GetGongstructInstancesSet[Diagram](stager.stage) {
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
