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
			system.Complexities = slices.DeleteFunc(system.Complexities, func(c *Complexity) bool { return c == complexity })
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

func (stager *Stager) treeComplexityWithinDiagramFloss(
	diagramFloss *DiagramFloss,
	system *System,
	complexity *Complexity,
	parentNode *tree.Node,
) {
	shape, ok := diagramFloss.map_Complexity_ComplexityShape[complexity]
	node := &tree.Node{
		Name:                    complexity.GetName(),
		IsExpanded:              slices.Contains(diagramFloss.ComplexitysWhoseNodeIsExpanded, complexity),
		IsNodeClickable:         true,
		IsInEditMode:            complexity.GetIsInRenameMode(),
		HasCheckboxButton:       true,
		IsChecked:               ok,
		CheckboxHasToolTip:      true,
		CheckboxToolTipPosition: tree.Left,
		CheckboxToolTipText: func() string {
			if ok {
				return "Click to remove complexity shape from diagram"
			}
			return "Click to add complexity shape to diagram"
		}(),
	}
	parentNode.Children = append(parentNode.Children, node)
	node.OnIsExpandedChange = onIsExpandedChangeSlice(stager, complexity, &diagramFloss.ComplexitysWhoseNodeIsExpanded)
	node.OnNameChange = stager.onNameChange(complexity)
	node.OnClick = onNodeClicked(stager, complexity)

	addRenameButton(complexity, node, stager)

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
			newShape := newShapeToDiagram(complexity, diagramFloss, &diagramFloss.Complexity_Shapes, stager, node.ClientOnY)
			newShape.Complexity = complexity
			newShape.Width = 180
			newShape.Height = 40

			if systemShape, found := diagramFloss.map_System_SystemShape[system]; found {
				newShape.X = systemShape.X + 30
				newShape.Y = systemShape.Y + 70 + float64(len(diagramFloss.Complexity_Shapes)-1)*50
			}
			stager.stage.Commit()
		} else {
			if shape, found := diagramFloss.map_Complexity_ComplexityShape[complexity]; found {
				shape.UnstageVoid(stager.stage)
				diagramFloss.Complexity_Shapes = slices.DeleteFunc(diagramFloss.Complexity_Shapes, func(s *ComplexityShape) bool { return s == shape })
				stager.stage.Commit()
			}
		}
	}
}
