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

		node := &tree.Node{
			Name:              "Category 1",
			HasCheckboxButton: false,
			IsExpanded:        diagram.IsCategory1NodeExpanded,
		}
		node.Buttons = []*tree.Button{
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
			node.Buttons[0].Icon = string(buttons.BUTTON_visibility_off)
		}
		node.Impl = &expandableNodeProxy{
			node:           node,
			stager:         stager,
			isNodeExpanded: &diagram.IsCategory1NodeExpanded,
		}
		diagramNode.Children = append(diagramNode.Children, node)

		for _, category := range GetGongstrucsSorted[*Category1](stager.stage) {
			isInDiagram := false
			for _, shape := range diagram.Category1Shapes {
				if shape.Category1 == category {
					isInDiagram = true
					continue
				}
			}

			categoryInstanceNode := &tree.Node{
				Name:              category.Name,
				HasCheckboxButton: true,
				IsChecked:         isInDiagram,
			}
			categoryInstanceNode.Impl = &Category1NodeProxy{
				node:     categoryInstanceNode,
				diagram:  diagram,
				category: category,
				stager:   stager,
			}
			node.Children = append(node.Children, categoryInstanceNode)
		}

		category2Node := &tree.Node{
			Name:              "Category 2",
			HasCheckboxButton: false,
			IsExpanded:        diagram.IsCategory2NodeExpanded,
		}
		category2Node.Impl = &expandableNodeProxy{
			isNodeExpanded: &diagram.IsCategory2NodeExpanded,
			node:           category2Node,
			stager:         stager,
		}
		diagramNode.Children = append(diagramNode.Children, category2Node)
		category2Node.Buttons = []*tree.Button{
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
			category2Node.Buttons[0].Icon = string(buttons.BUTTON_visibility_off)
		}
		for _, category3 := range GetGongstrucsSorted[*Category3](stager.stage) {

			isInDiagram := false

			for _, shape := range diagram.Category3Shapes {
				if shape.Category3 == category3 {
					isInDiagram = true
					continue
				}
			}

			catNode := &tree.Node{
				Name:              category3.Name,
				HasCheckboxButton: true,
				IsChecked:         isInDiagram,
			}
			catNode.Impl = &Category3NodeProxy{
				node:     catNode,
				diagram:  diagram,
				category: category3,
				stager:   stager,
			}
			category2Node.Children = append(category2Node.Children, catNode)
		}

		category3Node := &tree.Node{
			Name:              "Category 3",
			HasCheckboxButton: false,
			IsExpanded:        diagram.IsCategory3NodeExpanded,
		}
		category3Node.Impl = &expandableNodeProxy{
			isNodeExpanded: &diagram.IsCategory3NodeExpanded,
			node:           category3Node,
			stager:         stager,
		}
		diagramNode.Children = append(diagramNode.Children, category3Node)

		category3Node.Buttons = []*tree.Button{
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
			category3Node.Buttons[0].Icon = string(buttons.BUTTON_visibility_off)
		}
		for _, category := range GetGongstrucsSorted[*Category2](stager.stage) {

			isInDiagram := false

			for _, shape := range diagram.Category2Shapes {
				if shape.Category2 == category {
					isInDiagram = true
					continue
				}
			}

			artistNode := &tree.Node{
				Name:              category.Name,
				HasCheckboxButton: true,
				IsChecked:         isInDiagram,
			}
			artistNode.Impl = &Category2NodeProxy{
				node:     artistNode,
				diagram:  diagram,
				category: category,
				stager:   stager,
			}
			category3Node.Children = append(category3Node.Children, artistNode)
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
