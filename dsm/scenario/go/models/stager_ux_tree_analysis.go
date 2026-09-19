package models

import (
	"slices"
	"time"

	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
	"github.com/fullstack-lang/maticons/maticons"
)

func (stager *Stager) treeAnalysis(analysis *Analysis, parentNode *tree.Node) {
	analysisNode := &tree.Node{
		Name:                 analysis.Name,
		IsWithPreceedingIcon: true,
		PreceedingIcon:       string(maticons.BUTTON_library_books),
		IsExpanded:           analysis.GetIsExpanded(),
		IsNodeClickable:      true,
		IsInEditMode:         analysis.GetIsInRenameMode(),
	}

	parentNode.Children = append(parentNode.Children, analysisNode)

	analysisNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&analysis.IsExpanded)
	analysisNode.OnClick = onNodeClicked(stager, analysis)
	analysisNode.OnNameChange = stager.onNameChange(analysis)
	addRenameButton(analysis, analysisNode, stager)

	//
	// Scenarios
	//
	scenariosNode := &tree.Node{
		Name:            "Scenarios",
		FontStyle:       tree.ITALIC,
		IsExpanded:      analysis.IsScenariosNodeExpanded,
		IsNodeClickable: true,
	}
	analysisNode.Children = append(analysisNode.Children, scenariosNode)
	scenariosNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&analysis.IsScenariosNodeExpanded)
	scenariosNode.OnClick = onNodeClicked(stager, analysis)

	confScenarios := ItemButtonConfiguration[
		Scenario, *Scenario,
		Analysis, *Analysis,
	]{
		parentNode:                         scenariosNode,
		sliceForNewAddedItem:               &analysis.Scenarios,
		isParentNodeExpandedByAddOperation: true,
		parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
		parentNodeExpansionBooleanValue:    &analysis.IsScenariosNodeExpanded,
	}
	itemAdderCallback := addCreateItemButton(stager, confScenarios)
	itemAdderCallback.OnBeforeCommit = func() {
		newScenario := itemAdderCallback.createdItem
		newScenario.IsExpanded = true
		newScenario.IsDiagramsNodeExpanded = true

		// Create a default diagram for the new scenario
		defaultDiagram := (&Diagram{
			Name:                                "Default Diagram",
			IsChecked:                           true,
			IsExpanded:                          true,
			IsInDrawMode:                        true,
			AxisOrign_X:                         150,
			AxisOrign_Y:                         300,
			HorizontalAxis_Right_X:              950,
			VerticalAxis_Top_Y:                  100,
			VerticalAxis_Bottom_Y:               500,
			VerticalAxis_StrokeWidth:            2,
			Start:                               time.Now(),
			End:                                 time.Now().AddDate(1, 0, 0),
			IsParametersNodeExpanded:            true,
			IsActorStatesNodeExpanded:           true,
			IsEvolutionDirectionsNodeExpanded:   true,
			IsParametersAggregatesNodeExpanded:  true,
			IsActorStateTransitionsNodeExpanded: true,
		}).Stage(stager.stage)

		for d := range *stager.stage.GetInstancesSet[*Diagram]() {
			if d != defaultDiagram {
				d.IsChecked = false
			}
		}
		newScenario.Diagrams = append(newScenario.Diagrams, defaultDiagram)
	}

	slices.SortFunc(analysis.Scenarios, GongCompareGongstructByName)
	for _, scenario := range analysis.Scenarios {
		stager.treeScenario(scenario, scenariosNode)
	}
}
