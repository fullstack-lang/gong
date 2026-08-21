package models

import (
	"slices"

	buttons "github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) treeCompareAnalysisWithinLibrary(
	library *Library,
	compareAnalysis *CompareAnalysis,
	parentNode *tree.Node,
) {
	node := &tree.Node{
		Name:            compareAnalysis.GetName(),
		IsExpanded:      slices.Contains(library.CompareAnalysisWhoseNodeIsExpanded, compareAnalysis),
		IsNodeClickable: true,
		IsInEditMode:    compareAnalysis.GetIsInRenameMode(),
	}
	parentNode.Children = append(parentNode.Children, node)

	addRenameButton(compareAnalysis, node, stager)

	deleteButton := &tree.Button{
		Name:            "Delete",
		Icon:            string(buttons.BUTTON_delete),
		ToolTipText:     "Delete compare analysis",
		HasToolTip:      true,
		ToolTipPosition: tree.Above,
		OnClick: func() {
			library.RootCompareAnalysis = slices.DeleteFunc(library.RootCompareAnalysis, func(ca *CompareAnalysis) bool { return ca == compareAnalysis })
			compareAnalysis.Unstage(stager.stage)
			stager.stage.Commit()
		},
	}
	if node.Menu == nil {
		node.Menu = &tree.Menu{Name: "Menu"}
	}
	node.Menu.Buttons = append(node.Menu.Buttons, deleteButton)

	node.OnNameChange = stager.onNameChange(compareAnalysis)
	node.OnIsExpandedChange = onIsExpandedChangeSlice(stager, compareAnalysis, &library.CompareAnalysisWhoseNodeIsExpanded)
	node.OnClick = onNodeClicked(stager, compareAnalysis)
}
