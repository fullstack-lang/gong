package models

import (
	"slices"

	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
	"github.com/fullstack-lang/maticons/maticons"
)

func (stager *Stager) treeAnalysis(treeInstance *tree.Tree, analysis *Analysis, parentNodes *[]*tree.Node) {
	analysisNode := new(tree.Node)
	analysisNode.Name = analysis.Name
	analysisNode.IsWithPreceedingIcon = true
	analysisNode.PreceedingIcon = string(maticons.BUTTON_library_books)
	analysisNode.IsExpanded = analysis.GetIsExpanded()
	analysisNode.IsNodeClickable = true

	*parentNodes = append(*parentNodes, analysisNode)

	// Callback for when the analysis node is clicked or expanded
	analysisNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&analysis.IsExpanded)
	analysisNode.OnClick = func(frontNode *tree.Node) {
		stager.probeForm.FillUpFormFromGongstruct(analysis, GetPointerToGongstructName[*Analysis]())
		stager.stage.Commit()
	}

	analysisNode.IsInEditMode = analysis.GetIsInRenameMode()
	analysisNode.OnNameChange = stager.onNameChange(analysis)
	addRenameButton(analysis, analysisNode, stager)

	slices.SortFunc(analysis.Scenarios, CompareGongstructByName)
	for _, scenario := range analysis.Scenarios {
		stager.treeScenario(treeInstance, scenario, &analysisNode.Children)
	}
}
