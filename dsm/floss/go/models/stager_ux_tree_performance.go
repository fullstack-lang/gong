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

func (stager *Stager) treePerformanceWithinDiagramFloss(
	diagramFloss *DiagramFloss,
	system *System,
	performance *Performance,
	parentNode *tree.Node,
) {
	shape, ok := diagramFloss.map_Performance_PerformanceShape[performance]
	node := &tree.Node{
		Name:                    performance.GetName(),
		IsExpanded:              slices.Contains(diagramFloss.PerformancesWhoseNodeIsExpanded, performance),
		IsNodeClickable:         true,
		IsInEditMode:            performance.GetIsInRenameMode(),
		HasCheckboxButton:       true,
		IsChecked:               ok,
		CheckboxHasToolTip:      true,
		CheckboxToolTipPosition: tree.Left,
		CheckboxToolTipText: func() string {
			if ok {
				return "Click to remove performance shape from diagram"
			}
			return "Click to add performance shape to diagram"
		}(),
	}
	parentNode.Children = append(parentNode.Children, node)
	node.OnIsExpandedChange = onIsExpandedChangeSlice(stager, performance, &diagramFloss.PerformancesWhoseNodeIsExpanded)
	node.OnNameChange = stager.onNameChange(performance)
	node.OnClick = onNodeClicked(stager, performance)

	addRenameButton(performance, node, stager)

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
			newShape := newShapeToDiagram(performance, diagramFloss, &diagramFloss.Performance_Shapes, stager, node.ClientOnY)
			newShape.Performance = performance
			newShape.Width = 180
			newShape.Height = 40

			if systemShape, found := diagramFloss.map_System_SystemShape[system]; found {
				newShape.X = systemShape.X + 30
				newShape.Y = systemShape.Y + 70 + float64(len(diagramFloss.Performance_Shapes)-1)*50
			}
			stager.stage.Commit()
		} else {
			if shape, found := diagramFloss.map_Performance_PerformanceShape[performance]; found {
				shape.UnstageVoid(stager.stage)
				diagramFloss.Performance_Shapes = slices.DeleteFunc(diagramFloss.Performance_Shapes, func(s *PerformanceShape) bool { return s == shape })
				stager.stage.Commit()
			}
		}
	}
}
