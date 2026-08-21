package models

import (
	"slices"

	buttons "github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) treeEffortWithinLibrary(
	library *Library,
	effort *Effort,
	parentNode *tree.Node,
) {
	effortNode := &tree.Node{
		Name:            effort.GetName(),
		IsExpanded:      slices.Contains(library.EffortsWhoseNodeIsExpanded, effort),
		IsNodeClickable: true,
		IsInEditMode:    effort.GetIsInRenameMode(),
	}
	parentNode.Children = append(parentNode.Children, effortNode)

	addRenameButton(effort, effortNode, stager)

	deleteButton := &tree.Button{
		Name:            "Delete",
		Icon:            string(buttons.BUTTON_delete),
		ToolTipText:     "Delete",
		HasToolTip:      true,
		ToolTipPosition: tree.Above,
		OnClick: func() {
			library.RootEfforts = slices.DeleteFunc(library.RootEfforts, func(e *Effort) bool { return e == effort })
			effort.Unstage(stager.stage)
			stager.stage.Commit()
		},
	}
	if effortNode.Menu == nil {
		effortNode.Menu = &tree.Menu{Name: "Menu"}
	}
	effortNode.Menu.Buttons = append(effortNode.Menu.Buttons, deleteButton)

	effortNode.OnNameChange = stager.onNameChange(effort)
	effortNode.OnIsExpandedChange = onIsExpandedChangeSlice(stager, effort, &library.EffortsWhoseNodeIsExpanded)
	effortNode.OnClick = onNodeClicked(stager, effort)
}

func (stager *Stager) treeEffortWithinSystem(
	system *System,
	effort *Effort,
	parentNode *tree.Node,
) {
	effortNode := &tree.Node{
		Name:            effort.GetName(),
		IsExpanded:      slices.Contains(system.EffortsWhoseNodeIsExpanded, effort),
		IsNodeClickable: true,
		IsInEditMode:    effort.GetIsInRenameMode(),
	}
	parentNode.Children = append(parentNode.Children, effortNode)

	addRenameButton(effort, effortNode, stager)

	deleteButton := &tree.Button{
		Name:            "Delete",
		Icon:            string(buttons.BUTTON_delete),
		ToolTipText:     "Remove from system",
		HasToolTip:      true,
		ToolTipPosition: tree.Above,
		OnClick: func() {
			system.Efforts = slices.DeleteFunc(system.Efforts, func(e *Effort) bool { return e == effort })
			stager.stage.Commit()
		},
	}
	if effortNode.Menu == nil {
		effortNode.Menu = &tree.Menu{Name: "Menu"}
	}
	effortNode.Menu.Buttons = append(effortNode.Menu.Buttons, deleteButton)

	effortNode.OnNameChange = stager.onNameChange(effort)
	effortNode.OnIsExpandedChange = onIsExpandedChangeSlice(stager, effort, &system.EffortsWhoseNodeIsExpanded)
	effortNode.OnClick = onNodeClicked(stager, effort)
}
