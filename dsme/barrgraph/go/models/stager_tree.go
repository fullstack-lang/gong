package models

import (
	"math/rand/v2"
	"slices"

	"github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) tree() {
	stage := stager.stage
	stager.treeStage.Reset()

	tree_ := &tree.Tree{
		Name: "Tree of objects",
	}

	stager.probeForm.AddCommitNavigationNode(func(gni GongNodeIF) {
		tree_.RootNodes = append(tree_.RootNodes, gni.(*tree.Node))
	})

	for _, diagram := range GetGongstrucsSorted[*Diagram](stager.stage) {

		diagramNode := &tree.Node{
			Name:              diagram.Name,
			HasCheckboxButton: true,
			IsExpanded:        diagram.IsNodeExpanded,
		}
		if stager.desk.SelectedDiagram == diagram {
			diagramNode.IsChecked = true
		}

		diagramNode.Buttons = []*tree.Button{
			{
				Name: diagram.GetName(),
				Icon: string(buttons.BUTTON_edit),
				Impl: &toggleButtonProxy{
					stager:      stager,
					toggleValue: &diagram.IsEditable,
				},
			},
		}
		if diagram.IsEditable {
			diagramNode.Buttons[0].Icon = string(buttons.BUTTON_edit_off)
		}

		tree_.RootNodes = append(tree_.RootNodes, diagramNode)

		diagramNode.OnUpdate = func(_ *tree.Stage, stagedNode, frontNode *tree.Node) {
			if frontNode.IsChecked && !stagedNode.IsChecked {
				stager.desk.SelectedDiagram = diagram
			}

			if frontNode.IsExpanded != stagedNode.IsExpanded {
				diagram.IsNodeExpanded = !diagram.IsNodeExpanded
			}

			stage.Commit()
		}

		movementCategoryNode := &tree.Node{
			Name:              "Movements",
			HasCheckboxButton: false,
			IsExpanded:        diagram.IsMovementCategoryNodeExpanded,
			OnUpdate: func(_ *tree.Stage, stagedNode, frontNode *tree.Node) {
				if frontNode.IsExpanded != stagedNode.IsExpanded {
					diagram.IsMovementCategoryNodeExpanded = !diagram.IsMovementCategoryNodeExpanded
					stage.Commit()
				}
			},
		}
		movementCategoryNode.Buttons = []*tree.Button{
			{
				Name: diagram.GetName(),
				Icon: string(buttons.BUTTON_visibility),
				OnUpdate: func(stage *tree.Stage, updatedButton *tree.Button) {
					diagram.IsMovementCategoryShown = !diagram.IsMovementCategoryShown
					stage.Commit()
				},
			},
		}
		if diagram.IsMovementCategoryShown {
			movementCategoryNode.Buttons[0].Icon = string(buttons.BUTTON_visibility_off)
		}
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
				OnUpdate: func(_ *tree.Stage, stagedNode, frontNode *tree.Node) {
					if frontNode.IsChecked && !stagedNode.IsChecked {
						movementShape := &MovementShape{
							Movement: movement,
							Width:    240,
							Height:   80,
							X:        float64(int(rand.Float32()*100) + 10),
							Y:        float64(int(rand.Float32()*100) + 10),
						}
						movementShape.Stage(stage)
						diagram.MovementShapes = append(diagram.MovementShapes, movementShape)
					}
					if !frontNode.IsChecked && stagedNode.IsChecked {
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
			movementCategoryNode.Children = append(movementCategoryNode.Children, movementNode)

			movementNode.Buttons = []*tree.Button{
				{
					Name:            diagram.GetName(),
					Icon:            string(buttons.BUTTON_visibility_off),
					ToolTipText:     "Hide from diagram",
					HasToolTip:      true,
					ToolTipPosition: tree.Right,
					OnUpdate: func(_ *tree.Stage, _ *tree.Button) {
						shape.IsHidden = !shape.IsHidden
						stage.Commit()
					},
				},
			}
			if shape != nil && shape.IsHidden {
				movementNode.Buttons[0].Icon = string(buttons.BUTTON_visibility)
				movementNode.Buttons[0].ToolTipText = "Show on diagram"
			}
		}

		artefactTypeCategoryNode := &tree.Node{
			Name:              "ArtefactTypes",
			HasCheckboxButton: false,
			IsExpanded:        diagram.IsArtefactTypeCategoryNodeExpanded,
			OnUpdate: func(_ *tree.Stage, stagedNode, frontNode *tree.Node) {
				if frontNode.IsExpanded != stagedNode.IsExpanded {
					diagram.IsArtefactTypeCategoryNodeExpanded = !diagram.IsArtefactTypeCategoryNodeExpanded
					stage.Commit()
				}
			},
		}
		diagramNode.Children = append(diagramNode.Children, artefactTypeCategoryNode)
		artefactTypeCategoryNode.Buttons = []*tree.Button{
			{
				Name: diagram.GetName(),
				Icon: string(buttons.BUTTON_visibility),
				OnUpdate: func(_ *tree.Stage, _ *tree.Button) {
					diagram.IsArtefactTypeCategoryShown = !diagram.IsArtefactTypeCategoryShown
					stage.Commit()
				},
			},
		}
		if diagram.IsArtefactTypeCategoryShown {
			artefactTypeCategoryNode.Buttons[0].Icon = string(buttons.BUTTON_visibility_off)
		}
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
				OnUpdate: func(_ *tree.Stage, stagedNode, frontNode *tree.Node) {
					if frontNode.IsChecked && !stagedNode.IsChecked {
						artefactTypeShape := &ArtefactTypeShape{
							ArtefactType: artefactType,
							Width:        150,
							Height:       25,
							X:            float64(int(rand.Float32()*100) + 10),
							Y:            float64(int(rand.Float32()*100) + 10),
						}
						artefactTypeShape.Stage(stage)
						diagram.ArtefactTypeShapes = append(diagram.ArtefactTypeShapes, artefactTypeShape)
					}
					if !frontNode.IsChecked && stagedNode.IsChecked {
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
			artefactTypeCategoryNode.Children = append(artefactTypeCategoryNode.Children, artefactTypeNode)

			artefactTypeNode.Buttons = []*tree.Button{
				{
					Name:            diagram.GetName(),
					Icon:            string(buttons.BUTTON_visibility_off),
					ToolTipText:     "Hide from diagram",
					HasToolTip:      true,
					ToolTipPosition: tree.Right,
					OnUpdate: func(_ *tree.Stage, _ *tree.Button) {
						shape.IsHidden = !shape.IsHidden
						stage.Commit()
					},
				},
			}
		}

		artistCategoryNode := &tree.Node{
			Name:              "Artists",
			HasCheckboxButton: false,
			IsExpanded:        diagram.IsArtistCategoryNodeExpanded,
			OnUpdate: func(_ *tree.Stage, stagedNode, frontNode *tree.Node) {
				if frontNode.IsExpanded != stagedNode.IsExpanded {
					diagram.IsArtistCategoryNodeExpanded = !diagram.IsArtistCategoryNodeExpanded
					stage.Commit()
				}
			},
		}
		diagramNode.Children = append(diagramNode.Children, artistCategoryNode)

		artistCategoryNode.Buttons = []*tree.Button{
			{
				Name: diagram.GetName(),
				Icon: string(buttons.BUTTON_visibility),
				Impl: &toggleButtonProxy{
					stager:      stager,
					toggleValue: &diagram.IsArtistCategoryShown,
				},
			},
		}
		if diagram.IsArtistCategoryShown {
			artistCategoryNode.Buttons[0].Icon = string(buttons.BUTTON_visibility_off)
		}
		for _, artist := range GetGongstrucsSorted[*Artist](stager.stage) {

			isInDiagram := false

			var shape *ArtistShape
			for _, _shape := range diagram.ArtistShapes {
				if _shape.Artist == artist {
					isInDiagram = true
					shape = _shape
					continue
				}
			}

			artistNode := &tree.Node{
				Name:              artist.Name,
				HasCheckboxButton: true,
				IsChecked:         isInDiagram,
				OnUpdate: func(_ *tree.Stage, stagedNode, frontNode *tree.Node) {
					if frontNode.IsChecked && !stagedNode.IsChecked {
						artistShape := &ArtistShape{
							Artist: artist,
							Width:  80,
							Height: 30,
							X:      float64(int(rand.Float32()*100) + 10),
							Y:      float64(int(rand.Float32()*100) + 10),
						}
						artistShape.Stage(stage)
						diagram.ArtistShapes = append(diagram.ArtistShapes, artistShape)
					}
					if !frontNode.IsChecked && stagedNode.IsChecked {
						for idx, shape := range diagram.ArtistShapes {
							if shape.Artist == artist {
								shape.Unstage(stage)
								diagram.ArtistShapes = slices.Delete(diagram.ArtistShapes, idx, idx+1)
								continue
							}
						}
					}

					stage.Commit()
				},
			}
			artistCategoryNode.Children = append(artistCategoryNode.Children, artistNode)

			artistNode.Buttons = []*tree.Button{
				{
					Name:            diagram.GetName(),
					Icon:            string(buttons.BUTTON_visibility_off),
					ToolTipText:     "Hide from diagram",
					HasToolTip:      true,
					ToolTipPosition: tree.Right,
					OnUpdate: func(_ *tree.Stage, _ *tree.Button) {
						shape.IsHidden = !shape.IsHidden
						stage.Commit()
					},
				},
			}
			if diagram.IsArtistCategoryShown {
				artistNode.Buttons[0].Icon = string(buttons.BUTTON_visibility)
				artistNode.Buttons[0].ToolTipText = "Show on diagram"
			}
		}

		influenceCategoryNode := &tree.Node{
			Name:              "Influences",
			HasCheckboxButton: false,
			IsExpanded:        diagram.IsInfluenceCategoryNodeExpanded,
			OnUpdate: func(_ *tree.Stage, stagedNode, frontNode *tree.Node) {
				if frontNode.IsExpanded != stagedNode.IsExpanded {
					diagram.IsInfluenceCategoryNodeExpanded = !diagram.IsInfluenceCategoryNodeExpanded
					stage.Commit()
				}
			},
		}
		diagramNode.Children = append(diagramNode.Children, influenceCategoryNode)

		influenceCategoryNode.Buttons = []*tree.Button{
			{
				Name: diagram.GetName(),
				Icon: string(buttons.BUTTON_visibility),
				OnUpdate: func(_ *tree.Stage, _ *tree.Button) {
					diagram.IsInfluenceCategoryShown = !diagram.IsInfluenceCategoryShown
					stage.Commit()
				},
			},
		}
		if diagram.IsInfluenceCategoryShown {
			influenceCategoryNode.Buttons[0].Icon = string(buttons.BUTTON_visibility_off)
		}
		for _, influence := range GetGongstrucsSorted[*Influence](stager.stage) {

			isInDiagram := false

			for _, shape := range diagram.InfluenceShapes {
				if shape.Influence == influence {
					isInDiagram = true
					continue
				}
			}

			influenceNode := &tree.Node{
				Name:              influence.Name,
				HasCheckboxButton: true,
				IsChecked:         isInDiagram,
				OnUpdate: func(_ *tree.Stage, stagedNode, frontNode *tree.Node) {
					if frontNode.IsChecked && !stagedNode.IsChecked {
						influenceShape := &InfluenceShape{
							Influence: influence,
						}
						influenceShape.Stage(stage)
						diagram.InfluenceShapes = append(diagram.InfluenceShapes, influenceShape)
					}
					if !frontNode.IsChecked && stagedNode.IsChecked {
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
			influenceCategoryNode.Children = append(influenceCategoryNode.Children, influenceNode)
		}
	}

	tree.StageBranch(stager.treeStage, tree_)

	stager.treeStage.Commit()
}
