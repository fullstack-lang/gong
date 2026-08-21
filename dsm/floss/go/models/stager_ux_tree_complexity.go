package models

import (
	"slices"

	buttons "github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) treeComplexityWithinLibrary(
	library *Library,
	complexity *Complexity,
	parentNode *tree.Node,
) {
	complexityNode := &tree.Node{
		Name:            complexity.GetName(),
		IsExpanded:      slices.Contains(library.ComplexitysWhoseNodeIsExpanded, complexity),
		IsNodeClickable: true,
		IsInEditMode:    complexity.GetIsInRenameMode(),
	}
	parentNode.Children = append(parentNode.Children, complexityNode)

	addRenameButton(complexity, complexityNode, stager)

	deleteButton := &tree.Button{
		Name:            "Delete",
		Icon:            string(buttons.BUTTON_delete),
		ToolTipText:     "Delete",
		HasToolTip:      true,
		ToolTipPosition: tree.Above,
		OnClick: func() {
			library.RootComplexitys = slices.DeleteFunc(library.RootComplexitys, func(c *Complexity) bool { return c == complexity })
			complexity.Unstage(stager.stage)
			stager.stage.Commit()
		},
	}
	if complexityNode.Menu == nil {
		complexityNode.Menu = &tree.Menu{Name: "Menu"}
	}
	complexityNode.Menu.Buttons = append(complexityNode.Menu.Buttons, deleteButton)

	complexityNode.OnNameChange = stager.onNameChange(complexity)
	complexityNode.OnIsExpandedChange = onIsExpandedChangeSlice(stager, complexity, &library.ComplexitysWhoseNodeIsExpanded)
	complexityNode.OnClick = onNodeClicked(stager, complexity)
}

func (stager *Stager) treeComplexityWithinSystem(
	system *System,
	complexity *Complexity,
	parentNode *tree.Node,
) {
	complexityNode := &tree.Node{
		Name:            complexity.GetName(),
		IsExpanded:      slices.Contains(system.ComplexitysWhoseNodeIsExpanded, complexity),
		IsNodeClickable: true,
		IsInEditMode:    complexity.GetIsInRenameMode(),
	}
	parentNode.Children = append(parentNode.Children, complexityNode)

	addRenameButton(complexity, complexityNode, stager)

	deleteButton := &tree.Button{
		Name:            "Delete",
		Icon:            string(buttons.BUTTON_delete),
		ToolTipText:     "Remove from system",
		HasToolTip:      true,
		ToolTipPosition: tree.Above,
		OnClick: func() {
			system.Complexitys = slices.DeleteFunc(system.Complexitys, func(c *Complexity) bool { return c == complexity })
			stager.stage.Commit()
		},
	}
	if complexityNode.Menu == nil {
		complexityNode.Menu = &tree.Menu{Name: "Menu"}
	}
	complexityNode.Menu.Buttons = append(complexityNode.Menu.Buttons, deleteButton)

	complexityNode.OnNameChange = stager.onNameChange(complexity)
	complexityNode.OnIsExpandedChange = onIsExpandedChangeSlice(stager, complexity, &system.ComplexitysWhoseNodeIsExpanded)
	complexityNode.OnClick = onNodeClicked(stager, complexity)
}
