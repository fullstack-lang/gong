package models

import (
	"slices"

	buttons "github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) treePerformanceWithinLibrary(
	library *Library,
	performance *Performance,
	parentNode *tree.Node,
) {
	performanceNode := &tree.Node{
		Name:            performance.GetName(),
		IsExpanded:      slices.Contains(library.PerformancesWhoseNodeIsExpanded, performance),
		IsNodeClickable: true,
		IsInEditMode:    performance.GetIsInRenameMode(),
	}
	parentNode.Children = append(parentNode.Children, performanceNode)

	addRenameButton(performance, performanceNode, stager)

	deleteButton := &tree.Button{
		Name:            "Delete",
		Icon:            string(buttons.BUTTON_delete),
		ToolTipText:     "Delete",
		HasToolTip:      true,
		ToolTipPosition: tree.Above,
		OnClick: func() {
			library.RootPerformances = slices.DeleteFunc(library.RootPerformances, func(p *Performance) bool { return p == performance })
			performance.Unstage(stager.stage)
			stager.stage.Commit()
		},
	}
	if performanceNode.Menu == nil {
		performanceNode.Menu = &tree.Menu{Name: "Menu"}
	}
	performanceNode.Menu.Buttons = append(performanceNode.Menu.Buttons, deleteButton)

	performanceNode.OnNameChange = stager.onNameChange(performance)
	performanceNode.OnIsExpandedChange = onIsExpandedChangeSlice(stager, performance, &library.PerformancesWhoseNodeIsExpanded)
	performanceNode.OnClick = onNodeClicked(stager, performance)
}

func (stager *Stager) treePerformanceWithinSystem(
	system *System,
	performance *Performance,
	parentNode *tree.Node,
) {
	performanceNode := &tree.Node{
		Name:            performance.GetName(),
		IsExpanded:      slices.Contains(system.PerformancesWhoseNodeIsExpanded, performance),
		IsNodeClickable: true,
		IsInEditMode:    performance.GetIsInRenameMode(),
	}
	parentNode.Children = append(parentNode.Children, performanceNode)

	addRenameButton(performance, performanceNode, stager)

	deleteButton := &tree.Button{
		Name:            "Delete",
		Icon:            string(buttons.BUTTON_delete),
		ToolTipText:     "Remove from system",
		HasToolTip:      true,
		ToolTipPosition: tree.Above,
		OnClick: func() {
			system.Performances = slices.DeleteFunc(system.Performances, func(p *Performance) bool { return p == performance })
			stager.stage.Commit()
		},
	}
	if performanceNode.Menu == nil {
		performanceNode.Menu = &tree.Menu{Name: "Menu"}
	}
	performanceNode.Menu.Buttons = append(performanceNode.Menu.Buttons, deleteButton)

	performanceNode.OnNameChange = stager.onNameChange(performance)
	performanceNode.OnIsExpandedChange = onIsExpandedChangeSlice(stager, performance, &system.PerformancesWhoseNodeIsExpanded)
	performanceNode.OnClick = onNodeClicked(stager, performance)
}

