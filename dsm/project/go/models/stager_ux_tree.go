package models

import (
	"embed"
	"go/parser"
	"go/token"
	"strings"

	buttons "github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

var DataFS *embed.FS

func (stager *Stager) ux_tree() {
	stager.treeStage.Reset()

	rootLibrary := stager.getRootLibrary()
	_ = rootLibrary

	treeInstance := &tree.Tree{
		Name:       "Library Tree",
		HaveSearch: true,
	}

	stager.probeForm.AddCommitNavigationNode(func(gni GongNodeIF) {
		treeInstance.RootNodes = append(treeInstance.RootNodes, gni.(*tree.Node))
	})

	stager.treeLibrary(treeInstance, rootLibrary, &treeInstance.RootNodes)

	examplesNode := &tree.Node{
		Name:            "Examples",
		IsNodeClickable: true,
		IsExpanded:      true,
	}

	examplesNode.Menu = &tree.Menu{
		Name: "Examples Menu",
	}

	if DataFS != nil {
		entries, err := DataFS.ReadDir("data")
		if err == nil {
			for _, entry := range entries {
				if entry.IsDir() {
					continue
				}
				entryName := entry.Name()
				if !strings.HasSuffix(entryName, ".go") {
					continue
				}
				examplesNode.Menu.Buttons = append(examplesNode.Menu.Buttons, &tree.Button{
					Name: entryName,
					Icon: string(buttons.BUTTON_file_open),
					OnClick: func() {
						content, err := DataFS.ReadFile("data/" + entryName)
						if err == nil {
							// if the user loads an example file, we don't want the file to be automatically overwritten
							stager.stage.OnInitCommitCallback = nil

							stager.stage.Reset()

							fset := token.NewFileSet()
							file, err := parser.ParseFile(fset, "", string(content), parser.ParseComments)
							if err == nil {
								ParseAstFileFromAst(stager.stage, file, fset, true)
								stager.stage.ComputeReverseMaps()
								stager.stage.ComputeInstancesNb()
								stager.stage.ComputeReferenceAndOrders()
								stager.stage.Commit()
								stager.probeForm.Refresh()
							}
						}
					},
				})
			}
		}
	}
	treeInstance.RootNodes = append(treeInstance.RootNodes, examplesNode)

	tree.StageBranch(stager.treeStage, treeInstance)

	stager.treeStage.Commit()
}

func addLayoutButtons[AT interface {
	AbstractType
	comparable
}, CT LayoutConcreteType](
	stager *Stager,
	node *tree.Node,
	abstractElement AT,
	concreteShape CT,
	hasConcreteShape bool,
) {
	if treeNode, ok := any(abstractElement).(TreeAbstractType); ok {
		toggleAbstractLayoutButton := &tree.Button{
			Name: "Toggle Abstract Layout Direction to " + func() string {
				if treeNode.GetLayoutDirection() == Vertical {
					return "Horizontal"
				} else {
					return "Vertical"
				}
			}(),
			HasToolTip:      true,
			ToolTipPosition: tree.Above,
			OnClick: func() {
				if treeNode.GetLayoutDirection() == Vertical {
					treeNode.SetLayoutDirection(Horizontal)
				} else {
					treeNode.SetLayoutDirection(Vertical)
				}
				stager.stage.Commit()
			},
		}

		if treeNode.GetLayoutDirection() == Vertical {
			toggleAbstractLayoutButton.Icon = string(buttons.BUTTON_swap_horiz)
			toggleAbstractLayoutButton.ToolTipText = "Set layout to Horizontal"
		} else {
			toggleAbstractLayoutButton.Icon = string(buttons.BUTTON_swap_vert)
			toggleAbstractLayoutButton.ToolTipText = "Set layout to Vertical"
		}

		if node.Menu == nil {
			node.Menu = &tree.Menu{Name: "Menu"}
		}
		node.Menu.Buttons = append(node.Menu.Buttons, toggleAbstractLayoutButton)
	}

	if hasConcreteShape {
		toggleLayoutButton := &tree.Button{
			Name: "Toggle Concrete Layout Direction to " + func() string {
				if concreteShape.GetConcreteLayoutDirection() == Vertical {
					return "Horizontal"
				} else {
					return "Vertical"
				}
			}(),
			HasToolTip:      true,
			ToolTipPosition: tree.Above,
			OnClick: func() {
				if concreteShape.GetConcreteLayoutDirection() == Vertical {
					concreteShape.SetConcreteLayoutDirection(Horizontal)
				} else {
					concreteShape.SetConcreteLayoutDirection(Vertical)
				}
				stager.stage.Commit()
			},
		}

		if concreteShape.GetConcreteLayoutDirection() == Vertical {
			toggleLayoutButton.Icon = string(buttons.BUTTON_swap_horiz)
			toggleLayoutButton.ToolTipText = "Set concrete layout to Horizontal"
		} else {
			toggleLayoutButton.Icon = string(buttons.BUTTON_swap_vert)
			toggleLayoutButton.ToolTipText = "Set concrete layout to Vertical"
		}

		toggleOverrideButton := &tree.Button{
			Name:            "Toggle Override Layout Direction",
			HasToolTip:      true,
			ToolTipPosition: tree.Above,
			OnClick: func() {
				concreteShape.SetOverideLayoutDirection(!concreteShape.GetOverideLayoutDirection())
				stager.stage.Commit()
			},
		}

		if concreteShape.GetOverideLayoutDirection() {
			toggleOverrideButton.Icon = string(buttons.BUTTON_check_box)
			toggleOverrideButton.ToolTipText = "Disable layout override"
		} else {
			toggleOverrideButton.Icon = string(buttons.BUTTON_check_box_outline_blank)
			toggleOverrideButton.ToolTipText = "Enable layout override"
		}

		node.Menu.Buttons = append(node.Menu.Buttons, toggleLayoutButton)
		node.Menu.Buttons = append(node.Menu.Buttons, toggleOverrideButton)
	}
}
