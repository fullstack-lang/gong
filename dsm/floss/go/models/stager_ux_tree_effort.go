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

func (stager *Stager) treeEffortWithinDiagramFloss(
	diagramFloss *DiagramFloss,
	system *System,
	effort *Effort,
	parentNode *tree.Node,
) {
	shape, ok := diagramFloss.map_Effort_EffortShape[effort]
	node := &tree.Node{
		Name:                    effort.GetName(),
		IsExpanded:              slices.Contains(diagramFloss.EffortsWhoseNodeIsExpanded, effort),
		IsNodeClickable:         true,
		IsInEditMode:            effort.GetIsInRenameMode(),
		HasCheckboxButton:       true,
		IsChecked:               ok,
		CheckboxHasToolTip:      true,
		CheckboxToolTipPosition: tree.Left,
		CheckboxToolTipText: func() string {
			if ok {
				return "Click to remove effort shape from diagram"
			}
			return "Click to add effort shape to diagram"
		}(),
	}
	parentNode.Children = append(parentNode.Children, node)
	node.OnIsExpandedChange = onIsExpandedChangeSlice(stager, effort, &diagramFloss.EffortsWhoseNodeIsExpanded)
	node.OnNameChange = stager.onNameChange(effort)
	node.OnClick = onNodeClicked(stager, effort)

	addRenameButton(effort, node, stager)

	if ok {
		visibilityButton := &tree.Button{
			Name:            diagramFloss.GetName(),
			Icon:            string(buttons.BUTTON_visibility_off),
			ToolTipText:     "Hide from diagram",
			HasToolTip:      true,
			ToolTipPosition: tree.Right,
			OnClick: func() {
				shape.SetIsHidden(!shape.GetIsHidden())
				stager.stage.Commit()
			},
		}
		if shape.GetIsHidden() {
			visibilityButton.Icon = string(buttons.BUTTON_visibility)
			visibilityButton.ToolTipText = "Show on diagram"
		}
		node.Buttons = append(node.Buttons, visibilityButton)
	}

	node.OnIsCheckedChanged = func(isChecked bool) {
		if isChecked {
			newShape := newShapeToDiagram(effort, diagramFloss, &diagramFloss.Effort_Shapes, stager, node.ClientOnY)
			newShape.Effort = effort
			newShape.Width = 180
			newShape.Height = 40

			if systemShape, found := diagramFloss.map_System_SystemShape[system]; found {
				newShape.X = systemShape.X + 30
				newShape.Y = systemShape.Y + 70 + float64(len(diagramFloss.Effort_Shapes)-1)*50
			}
			stager.stage.Commit()
		} else {
			if shape, found := diagramFloss.map_Effort_EffortShape[effort]; found {
				shape.UnstageVoid(stager.stage)
				diagramFloss.Effort_Shapes = slices.DeleteFunc(diagramFloss.Effort_Shapes, func(s *EffortShape) bool { return s == shape })
				stager.stage.Commit()
			}
		}
	}
}
