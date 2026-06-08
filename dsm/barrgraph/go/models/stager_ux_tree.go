package models

import (
	"embed"
	"go/parser"
	"go/token"
	"math/rand/v2"
	"slices"
	"strings"

	"github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

var DataFS *embed.FS

func (stager *Stager) ux_tree() {
	stage := stager.stage
	stager.treeStage.Reset()

	tree_ := &tree.Tree{
		Name: "Tree of objects",
	}

	stager.probeForm.AddCommitNavigationNode(func(gni GongNodeIF) {
		tree_.RootNodes = append(tree_.RootNodes, gni.(*tree.Node))
	})

	galleryNode := &tree.Node{
		Name:            "Gallery",
		IsNodeClickable: true,
		IsExpanded:      true,
	}

	galleryNode.Menu = &tree.Menu{
		Name: "Gallery Menu",
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
				galleryNode.Menu.Buttons = append(galleryNode.Menu.Buttons, &tree.Button{
					Name: entryName,
					Icon: string(buttons.BUTTON_file_open),
					OnClick: func() {
						content, err := DataFS.ReadFile("data/" + entryName)
						if err == nil {
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
	tree_.RootNodes = append(tree_.RootNodes, galleryNode)

	for _, diagram := range GetGongstrucsSorted[*Diagram](stager.stage) {

		diagramNode := &tree.Node{
			Name:              diagram.Name,
			HasCheckboxButton: true,
			IsExpanded:        diagram.IsNodeExpanded,
			IsNodeClickable:   true,
		}
		if stager.desk.SelectedDiagram == diagram {
			diagramNode.IsChecked = true
		}

		diagramNode.Buttons = []*tree.Button{
			{
				Name:            diagram.GetName(),
				Icon:            string(buttons.BUTTON_edit),
				HasToolTip:      true,
				ToolTipText:     "Edit diagram",
				ToolTipPosition: tree.Right,
				OnClick: func() {
					diagram.IsEditable = !diagram.IsEditable
					stage.Commit()
				},
			},
		}
		if diagram.IsEditable {
			diagramNode.Buttons[0].Icon = string(buttons.BUTTON_edit_off)
			diagramNode.Buttons[0].ToolTipText = "Stop editing diagram"
		}

		tree_.RootNodes = append(tree_.RootNodes, diagramNode)

		diagramNode.OnIsCheckedChanged = func(isChecked bool) {
			if isChecked {
				stager.desk.SelectedDiagram = diagram
				stage.Commit()
			}
		}

		diagramNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&diagram.IsNodeExpanded)

		diagramNode.OnClick = onNodeClicked(stager, diagram)

		movementCategoryNode := &tree.Node{
			Name:               "Movements",
			HasCheckboxButton:  false,
			IsExpanded:         diagram.IsMovementCategoryNodeExpanded,
			OnIsExpandedChange: stager.onIsExpandedChangeBool(&diagram.IsMovementCategoryNodeExpanded),
		}
		movementCategoryNode.Buttons = []*tree.Button{
			{
				Name:            diagram.GetName(),
				Icon:            string(buttons.BUTTON_visibility),
				HasToolTip:      true,
				ToolTipText:     "Show movements on diagram",
				ToolTipPosition: tree.Right,
				OnClick: func() {
					diagram.IsMovementCategoryHidden = !diagram.IsMovementCategoryHidden
					stage.Commit()
				},
			},
		}
		if !diagram.IsMovementCategoryHidden {
			movementCategoryNode.Buttons[0].Icon = string(buttons.BUTTON_visibility_off)
			movementCategoryNode.Buttons[0].ToolTipText = "Hide movements from diagram"
		}

		movementCategoryNode.Buttons = append(movementCategoryNode.Buttons,
			&tree.Button{
				Name:            diagram.GetName() + " add movement",
				Icon:            string(buttons.BUTTON_add),
				HasToolTip:      true,
				ToolTipText:     "Add a Movement",
				ToolTipPosition: tree.Right,
				OnClick: func() {
					movement := (&Movement{Name: ""}).Stage(stage)
					movement.SetIsInRenameMode(true)

					movementShape := &MovementShape{
						Movement: movement,
						PositionAndSize: PositionAndSize{
							Width:  240,
							Height: 80,
							X:      float64(int(rand.Float32()*100) + 10),
							Y:      float64(int(rand.Float32()*100) + 10),
						},
					}
					movementShape.Stage(stage)
					diagram.MovementShapes = append(diagram.MovementShapes, movementShape)
					stage.Commit()
				},
			})
		diagramNode.Children = append(diagramNode.Children, movementCategoryNode)

		map_Movement_MovementShape := make(map[*Movement]*MovementShape, 0)
		for _, _shape := range diagram.MovementShapes {
			map_Movement_MovementShape[_shape.Movement] = _shape
		}

		for _, movement := range GetGongstrucsSorted[*Movement](stager.stage) {
			isInDiagram := false
			var shape *MovementShape
			shape = map_Movement_MovementShape[movement]
			if shape != nil {
				isInDiagram = true
			}

			movementNode := &tree.Node{
				Name:              movement.Name,
				HasCheckboxButton: true,
				IsChecked:         isInDiagram,
				IsNodeClickable:   true,
				IsInEditMode:      movement.GetIsInRenameMode(),
				OnIsCheckedChanged: func(isChecked bool) {
					if isChecked {
						movementShape := &MovementShape{
							Movement: movement,
							PositionAndSize: PositionAndSize{
								Width:  240,
								Height: 80,
								X:      float64(int(rand.Float32()*100) + 10),
								Y:      float64(int(rand.Float32()*100) + 10),
							},
						}
						movementShape.Stage(stage)
						diagram.MovementShapes = append(diagram.MovementShapes, movementShape)
					} else {
						for idx, shape := range diagram.MovementShapes {
							if shape.Movement == movement {
								diagram.MovementShapes = slices.Delete(diagram.MovementShapes, idx, idx+1)
								continue
							}
						}
					}
					stage.Commit()
				},
			}
			movementNode.OnClick = onNodeClicked(stager, movement)
			movementNode.OnNameChange = stager.onNameChange(movement)
			movementCategoryNode.Children = append(movementCategoryNode.Children, movementNode)

			addRenameButton(movement, movementNode, stager)

			if shape != nil {
				movementNode.Buttons = []*tree.Button{
					{
						Name:            diagram.GetName(),
						Icon:            string(buttons.BUTTON_visibility_off),
						ToolTipText:     "Hide from diagram",
						HasToolTip:      true,
						ToolTipPosition: tree.Right,
						OnClick: func() {
							shape.IsHidden = !shape.IsHidden
							stage.Commit()
						},
					},
				}
				if shape.IsHidden {
					movementNode.Buttons[0].Icon = string(buttons.BUTTON_visibility)
					movementNode.Buttons[0].ToolTipText = "Show on diagram"
				}
			}
		}

		artefactTypeCategoryNode := &tree.Node{
			Name:               "ArtefactTypes",
			HasCheckboxButton:  false,
			IsExpanded:         diagram.IsArtefactTypeCategoryNodeExpanded,
			OnIsExpandedChange: stager.onIsExpandedChangeBool(&diagram.IsArtefactTypeCategoryNodeExpanded),
		}
		diagramNode.Children = append(diagramNode.Children, artefactTypeCategoryNode)
		artefactTypeCategoryNode.Buttons = []*tree.Button{
			{
				Name:            diagram.GetName(),
				Icon:            string(buttons.BUTTON_visibility),
				HasToolTip:      true,
				ToolTipText:     "Show artefact types on diagram",
				ToolTipPosition: tree.Right,
				OnClick: func() {
					diagram.IsArtefactTypeCategoryHidden = !diagram.IsArtefactTypeCategoryHidden
					stage.Commit()
				},
			},
		}
		if !diagram.IsArtefactTypeCategoryHidden {
			artefactTypeCategoryNode.Buttons[0].Icon = string(buttons.BUTTON_visibility_off)
			artefactTypeCategoryNode.Buttons[0].ToolTipText = "Hide artefact types from diagram"
		}

		artefactTypeCategoryNode.Buttons = append(artefactTypeCategoryNode.Buttons,
			&tree.Button{
				Name:            diagram.GetName() + " add artefactType",
				Icon:            string(buttons.BUTTON_add),
				HasToolTip:      true,
				ToolTipText:     "Add an ArtefactType",
				ToolTipPosition: tree.Right,
				OnClick: func() {
					artefactType := (&ArtefactType{Name: ""}).Stage(stage)
					artefactType.SetIsInRenameMode(true)

					artefactTypeShape := &ArtefactTypeShape{
						ArtefactType: artefactType,
						PositionAndSize: PositionAndSize{
							Width:  150,
							Height: 25,
							X:      float64(int(rand.Float32()*100) + 10),
							Y:      float64(int(rand.Float32()*100) + 10),
						},
					}
					artefactTypeShape.Stage(stage)
					diagram.ArtefactTypeShapes = append(diagram.ArtefactTypeShapes, artefactTypeShape)
					stage.Commit()
				},
			})

		for _, artefactType := range GetGongstrucsSorted[*ArtefactType](stager.stage) {

			isInDiagram := false
			var shape *ArtefactTypeShape
			for _, _shape := range diagram.ArtefactTypeShapes {
				if _shape.ArtefactType == artefactType {
					isInDiagram = true
					shape = _shape
					continue
				}
			}

			artefactTypeNode := &tree.Node{
				Name:              artefactType.Name,
				HasCheckboxButton: true,
				IsChecked:         isInDiagram,
				IsNodeClickable:   true,
				IsInEditMode:      artefactType.GetIsInRenameMode(),
				OnIsCheckedChanged: func(isChecked bool) {
					if isChecked {
						artefactTypeShape := &ArtefactTypeShape{
							ArtefactType: artefactType,
							PositionAndSize: PositionAndSize{
								Width:  150,
								Height: 25,
								X:      float64(int(rand.Float32()*100) + 10),
								Y:      float64(int(rand.Float32()*100) + 10),
							},
						}
						artefactTypeShape.Stage(stage)
						diagram.ArtefactTypeShapes = append(diagram.ArtefactTypeShapes, artefactTypeShape)
					} else {
						for idx, shape := range diagram.ArtefactTypeShapes {
							if shape.ArtefactType == artefactType {
								shape.Unstage(stage)
								diagram.ArtefactTypeShapes = slices.Delete(diagram.ArtefactTypeShapes, idx, idx+1)
								continue
							}
						}
					}
					stage.Commit()
				},
			}
			artefactTypeNode.OnClick = onNodeClicked(stager, artefactType)
			artefactTypeNode.OnNameChange = stager.onNameChange(artefactType)
			artefactTypeCategoryNode.Children = append(artefactTypeCategoryNode.Children, artefactTypeNode)

			addRenameButton(artefactType, artefactTypeNode, stager)

			if shape != nil {
				artefactTypeNode.Buttons = []*tree.Button{
					{
						Name:            diagram.GetName(),
						Icon:            string(buttons.BUTTON_visibility_off),
						ToolTipText:     "Hide from diagram",
						HasToolTip:      true,
						ToolTipPosition: tree.Right,
						OnClick: func() {
							shape.IsHidden = !shape.IsHidden
							stage.Commit()
						},
					},
				}
				if shape.IsHidden {
					artefactTypeNode.Buttons[0].Icon = string(buttons.BUTTON_visibility)
					artefactTypeNode.Buttons[0].ToolTipText = "Show on diagram"
				}
			}
		}

		artistCategoryNode := &tree.Node{
			Name:               "Artists",
			HasCheckboxButton:  false,
			IsExpanded:         diagram.IsArtistCategoryNodeExpanded,
			OnIsExpandedChange: stager.onIsExpandedChangeBool(&diagram.IsArtistCategoryNodeExpanded),
		}
		diagramNode.Children = append(diagramNode.Children, artistCategoryNode)

		artistCategoryNode.Buttons = []*tree.Button{
			{
				Name:            diagram.GetName(),
				Icon:            string(buttons.BUTTON_visibility),
				HasToolTip:      true,
				ToolTipText:     "Show artists on diagram",
				ToolTipPosition: tree.Right,
				OnClick: func() {
					diagram.IsArtistCategoryHidden = !diagram.IsArtistCategoryHidden
					stage.Commit()
				},
			},
		}
		if !diagram.IsArtistCategoryHidden {
			artistCategoryNode.Buttons[0].Icon = string(buttons.BUTTON_visibility_off)
			artistCategoryNode.Buttons[0].ToolTipText = "Hide artists from diagram"
		}

		artistCategoryNode.Buttons = append(artistCategoryNode.Buttons,
			&tree.Button{
				Name:            diagram.GetName() + " add artist",
				Icon:            string(buttons.BUTTON_add),
				HasToolTip:      true,
				ToolTipText:     "Add an Artist",
				ToolTipPosition: tree.Right,
				OnClick: func() {
					artist := (&Artist{Name: ""}).Stage(stage)
					artist.SetIsInRenameMode(true)

					artistShape := &ArtistShape{
						Artist: artist,
						PositionAndSize: PositionAndSize{
							Width:  80,
							Height: 30,
							X:      float64(int(rand.Float32()*100) + 10),
							Y:      float64(int(rand.Float32()*100) + 10),
						},
					}
					artistShape.Stage(stage)
					diagram.ArtistShapes = append(diagram.ArtistShapes, artistShape)
					stage.Commit()
				},
			})

		for _, element := range GetGongstrucsSorted[*Artist](stager.stage) {

			isInDiagram := false

			var shape *ArtistShape
			for _, _shape := range diagram.ArtistShapes {
				if _shape.Artist == element {
					isInDiagram = true
					shape = _shape
					continue
				}
			}

			node := &tree.Node{
				Name:              element.Name,
				HasCheckboxButton: true,
				IsChecked:         isInDiagram,
				IsNodeClickable:   true,
				IsInEditMode:      element.GetIsInRenameMode(),
				OnIsCheckedChanged: func(isChecked bool) {
					if isChecked {
						artistShape := &ArtistShape{
							Artist: element,
							PositionAndSize: PositionAndSize{
								Width:  80,
								Height: 30,
								X:      float64(int(rand.Float32()*100) + 10),
								Y:      float64(int(rand.Float32()*100) + 10),
							},
						}
						artistShape.Stage(stage)
						diagram.ArtistShapes = append(diagram.ArtistShapes, artistShape)
					} else {
						for idx, shape := range diagram.ArtistShapes {
							if shape.Artist == element {
								shape.Unstage(stage)
								diagram.ArtistShapes = slices.Delete(diagram.ArtistShapes, idx, idx+1)
								continue
							}
						}
					}
					stage.Commit()
				},
			}
			node.OnClick = onNodeClicked(stager, element)
			node.OnNameChange = stager.onNameChange(element)
			artistCategoryNode.Children = append(artistCategoryNode.Children, node)
			addRenameButton(element, node, stager)

			if shape != nil {
				visibilityButton := &tree.Button{
					Name:            diagram.GetName(),
					Icon:            string(buttons.BUTTON_visibility_off),
					ToolTipText:     "Hide from diagram",
					HasToolTip:      true,
					ToolTipPosition: tree.Right,
					OnClick: func() {
						shape.IsHidden = !shape.IsHidden
						stage.Commit()
					},
				}
				if shape.IsHidden {
					visibilityButton.Icon = string(buttons.BUTTON_visibility)
					visibilityButton.ToolTipText = "Show on diagram"
				}
				node.Buttons = append(node.Buttons, visibilityButton)
			}
		}

		influenceCategoryNode := &tree.Node{
			Name:               "Influences",
			HasCheckboxButton:  false,
			IsExpanded:         diagram.IsInfluenceCategoryNodeExpanded,
			OnIsExpandedChange: stager.onIsExpandedChangeBool(&diagram.IsInfluenceCategoryNodeExpanded),
		}
		diagramNode.Children = append(diagramNode.Children, influenceCategoryNode)

		influenceCategoryNode.Buttons = []*tree.Button{
			{
				Name:            diagram.GetName(),
				Icon:            string(buttons.BUTTON_visibility),
				HasToolTip:      true,
				ToolTipText:     "Show influences on diagram",
				ToolTipPosition: tree.Right,
				OnClick: func() {
					diagram.IsInfluenceCategoryHidden = !diagram.IsInfluenceCategoryHidden
					stage.Commit()
				},
			},
		}
		if !diagram.IsInfluenceCategoryHidden {
			influenceCategoryNode.Buttons[0].Icon = string(buttons.BUTTON_visibility_off)
			influenceCategoryNode.Buttons[0].ToolTipText = "Hide influences from diagram"
		}

		for _, influence := range GetGongstrucsSorted[*Influence](stager.stage) {

			isInDiagram := false
			var shape *InfluenceShape

			for _, _shape := range diagram.InfluenceShapes {
				if _shape.Influence == influence {
					isInDiagram = true
					shape = _shape
					continue
				}
			}

			influenceNode := &tree.Node{
				Name:              influence.Name,
				HasCheckboxButton: true,
				IsChecked:         isInDiagram,
				IsNodeClickable:   true,
				IsInEditMode:      influence.GetIsInRenameMode(),
				OnIsCheckedChanged: func(isChecked bool) {
					if isChecked {
						influenceShape := &InfluenceShape{
							Influence: influence,
						}
						influenceShape.Stage(stage)
						diagram.InfluenceShapes = append(diagram.InfluenceShapes, influenceShape)
					} else {
						for idx, shape := range diagram.InfluenceShapes {
							if shape.Influence == influence {
								shape.Unstage(stage)
								diagram.InfluenceShapes = slices.Delete(diagram.InfluenceShapes, idx, idx+1)
								continue
							}
						}
					}
					stage.Commit()
				},
			}
			influenceNode.OnClick = onNodeClicked(stager, influence)
			influenceNode.OnNameChange = stager.onNameChange(influence)
			influenceCategoryNode.Children = append(influenceCategoryNode.Children, influenceNode)
			addRenameButton(influence, influenceNode, stager)

			if shape != nil {
				influenceNode.Buttons = []*tree.Button{
					{
						Name:            diagram.GetName(),
						Icon:            string(buttons.BUTTON_visibility_off),
						ToolTipText:     "Hide from diagram",
						HasToolTip:      true,
						ToolTipPosition: tree.Right,
						OnClick: func() {
							shape.IsHidden = !shape.IsHidden
							stage.Commit()
						},
					},
				}
				if shape.IsHidden {
					influenceNode.Buttons[0].Icon = string(buttons.BUTTON_visibility)
					influenceNode.Buttons[0].ToolTipText = "Show on diagram"
				}
			}
		}
	}

	tree.StageBranch(stager.treeStage, tree_)

	stager.treeStage.Commit()
}
