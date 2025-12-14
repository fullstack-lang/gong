package models

import (
	"github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) TreeStageUpdate() {
	stager.treeStage.Reset()

	tree_ := &tree.Tree{
		Name: "Tree of objects",
	}

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

		diagramNode.Impl = &DiagramNodeProxy{
			stager:  stager,
			Node:    diagramNode,
			Diagram: diagram,
		}

		movementCategoryNode := &tree.Node{
			Name:              "Movements",
			HasCheckboxButton: false,
			IsExpanded:        diagram.IsCategory1NodeExpanded,
		}
		movementCategoryNode.Buttons = []*tree.Button{
			{
				Name: diagram.GetName(),
				Icon: string(buttons.BUTTON_visibility),
				Impl: &toggleButtonProxy{
					stager:      stager,
					toggleValue: &diagram.IsCategory1Shown,
				},
			},
		}
		if diagram.IsCategory1Shown {
			movementCategoryNode.Buttons[0].Icon = string(buttons.BUTTON_visibility_off)
		}
		movementCategoryNode.Impl = &expandableNodeProxy{
			node:           movementCategoryNode,
			stager:         stager,
			isNodeExpanded: &diagram.IsCategory1NodeExpanded,
		}
		diagramNode.Children = append(diagramNode.Children, movementCategoryNode)

		for _, movement := range GetGongstrucsSorted[*Category1](stager.stage) {
			isInDiagram := false
			for _, shape := range diagram.Category1Shapes {
				if shape.Category1 == movement {
					isInDiagram = true
					continue
				}
			}

			movementNode := &tree.Node{
				Name:              movement.Name,
				HasCheckboxButton: true,
				IsChecked:         isInDiagram,
			}
			movementNode.Impl = &Category1NodeProxy{
				node:     movementNode,
				diagram:  diagram,
				category: movement,
				stager:   stager,
			}
			movementCategoryNode.Children = append(movementCategoryNode.Children, movementNode)
		}

		artefactTypeCategoryNode := &tree.Node{
			Name:              "ArtefactTypes",
			HasCheckboxButton: false,
			IsExpanded:        diagram.IsCategory2NodeExpanded,
		}
		artefactTypeCategoryNode.Impl = &expandableNodeProxy{
			isNodeExpanded: &diagram.IsCategory2NodeExpanded,
			node:           artefactTypeCategoryNode,
			stager:         stager,
		}
		diagramNode.Children = append(diagramNode.Children, artefactTypeCategoryNode)
		artefactTypeCategoryNode.Buttons = []*tree.Button{
			{
				Name: diagram.GetName(),
				Icon: string(buttons.BUTTON_visibility),
				Impl: &toggleButtonProxy{
					stager:      stager,
					toggleValue: &diagram.IsCategory2Shown,
				},
			},
		}
		if diagram.IsCategory2Shown {
			artefactTypeCategoryNode.Buttons[0].Icon = string(buttons.BUTTON_visibility_off)
		}
		for _, artefactType := range GetGongstrucsSorted[*Category3](stager.stage) {

			isInDiagram := false

			for _, shape := range diagram.Category3Shapes {
				if shape.Category3 == artefactType {
					isInDiagram = true
					continue
				}
			}

			artefactTypeNode := &tree.Node{
				Name:              artefactType.Name,
				HasCheckboxButton: true,
				IsChecked:         isInDiagram,
			}
			artefactTypeNode.Impl = &Category3NodeProxy{
				node:     artefactTypeNode,
				diagram:  diagram,
				category: artefactType,
				stager:   stager,
			}
			artefactTypeCategoryNode.Children = append(artefactTypeCategoryNode.Children, artefactTypeNode)
		}

		artistCategoryNode := &tree.Node{
			Name:              "Artists",
			HasCheckboxButton: false,
			IsExpanded:        diagram.IsCategory3NodeExpanded,
		}
		artistCategoryNode.Impl = &expandableNodeProxy{
			isNodeExpanded: &diagram.IsCategory3NodeExpanded,
			node:           artistCategoryNode,
			stager:         stager,
		}
		diagramNode.Children = append(diagramNode.Children, artistCategoryNode)

		artistCategoryNode.Buttons = []*tree.Button{
			{
				Name: diagram.GetName(),
				Icon: string(buttons.BUTTON_visibility),
				Impl: &toggleButtonProxy{
					stager:      stager,
					toggleValue: &diagram.IsCategory3Shown,
				},
			},
		}
		if diagram.IsCategory3Shown {
			artistCategoryNode.Buttons[0].Icon = string(buttons.BUTTON_visibility_off)
		}
		for _, artist := range GetGongstrucsSorted[*Category2](stager.stage) {

			isInDiagram := false

			for _, shape := range diagram.Category2Shapes {
				if shape.Category2 == artist {
					isInDiagram = true
					continue
				}
			}

			artistNode := &tree.Node{
				Name:              artist.Name,
				HasCheckboxButton: true,
				IsChecked:         isInDiagram,
			}
			artistNode.Impl = &Category2NodeProxy{
				node:     artistNode,
				diagram:  diagram,
				category: artist,
				stager:   stager,
			}
			artistCategoryNode.Children = append(artistCategoryNode.Children, artistNode)
		}

		influenceCategoryNode := &tree.Node{
			Name:              "Influences",
			HasCheckboxButton: false,
			IsExpanded:        diagram.IsInfluenceCategoryNodeExpanded,
		}
		influenceCategoryNode.Impl = &expandableNodeProxy{
			isNodeExpanded: &diagram.IsInfluenceCategoryNodeExpanded,
			node:           influenceCategoryNode,
			stager:         stager,
		}
		diagramNode.Children = append(diagramNode.Children, influenceCategoryNode)

		influenceCategoryNode.Buttons = []*tree.Button{
			{
				Name: diagram.GetName(),
				Icon: string(buttons.BUTTON_visibility),
				Impl: &toggleButtonProxy{
					stager:      stager,
					toggleValue: &diagram.IsInfluenceCategoryShown,
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
			}
			influenceNode.Impl = &InfluenceNodeProxy{
				node:      influenceNode,
				diagram:   diagram,
				influence: influence,
				stager:    stager,
			}
			influenceCategoryNode.Children = append(influenceCategoryNode.Children, influenceNode)
		}
	}

	tree.StageBranch(stager.treeStage, tree_)

	stager.treeStage.Commit()
}

type DiagramNodeProxy struct {
	stager  *Stager
	Diagram *Diagram
	Node    *tree.Node
}

// OnAfterUpdate implements models.NodeImplInterface.
func (d *DiagramNodeProxy) OnAfterUpdate(stage *tree.Stage, stagedNode *tree.Node, frontNode *tree.Node) {

	if frontNode.IsChecked && !stagedNode.IsChecked {
		d.stager.desk.SelectedDiagram = d.Diagram
	}

	if frontNode.IsExpanded != stagedNode.IsExpanded {
		d.Diagram.IsNodeExpanded = !d.Diagram.IsNodeExpanded
	}

	d.stager.stage.Commit()
}
