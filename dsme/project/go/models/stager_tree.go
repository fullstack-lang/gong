package models

import (
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"

	"github.com/fullstack-lang/gong/lib/tree/go/buttons"
)

func (stager *Stager) tree() {

	stager.treeStage.Reset()

	root := stager.root

	treeInstance := &tree.Tree{Name: "Project Tree"}

	allProjectsNode := &tree.Node{
		Name:       "** Tree of Projects **",
		IsExpanded: true,
	}
	treeInstance.RootNodes = append(treeInstance.RootNodes, allProjectsNode)

	for _, project := range root.Projects {
		projectNode := &tree.Node{
			Name:            project.Name,
			IsExpanded:      project.IsExpanded,
			IsNodeClickable: true,
		}
		treeInstance.RootNodes = append(treeInstance.RootNodes, projectNode)
		projectNode.Impl = &NodeProxy[*Project]{
			stager:   stager,
			node:     projectNode,
			instance: project,
		}

		pbsNode := &tree.Node{
			Name:            "PBS",
			FontStyle: tree.ITALIC,
			IsExpanded:      project.IsPBSNodeExpanded,
			IsNodeClickable: true,
		}
		projectNode.Children = append(projectNode.Children, pbsNode)
		pbsNode.Impl = &expandableNodeProxy{
			node:           pbsNode,
			stager:         stager,
			isNodeExpanded: &project.IsPBSNodeExpanded,
		}

		addAddItemButton(stager, pbsNode, &project.RootProducts)

		for _, product := range project.RootProducts {
			stager.treePBS(product, pbsNode)
		}

		wbsNode := &tree.Node{
			Name:            "WBS",
						FontStyle: tree.ITALIC,
			IsExpanded:      project.IsWBSNodeExpanded,
			IsNodeClickable: true,
		}
		projectNode.Children = append(projectNode.Children, wbsNode)
		wbsNode.Impl = &expandableNodeProxy{
			node:           wbsNode,
			stager:         stager,
			isNodeExpanded: &project.IsWBSNodeExpanded,
		}

		addAddItemButton(stager, wbsNode, &project.RootTasks)

		for _, task := range project.RootTasks {
			stager.treeWBS(task, wbsNode)
		}

		diagramsNode := &tree.Node{
			Name:            "Diagrams",
			FontStyle: tree.ITALIC,
			IsExpanded:      project.IsDiagramsNodeExpanded,
			IsNodeClickable: true,
		}
		projectNode.Children = append(projectNode.Children, diagramsNode)
		diagramsNode.Impl = &expandableNodeProxy{
			node:           diagramsNode,
			stager:         stager,
			isNodeExpanded: &project.IsDiagramsNodeExpanded,
		}

		addAddItemButton(stager, diagramsNode, &project.Diagrams)

		for _, diagram := range project.Diagrams {
			diagramNode := &tree.Node{
				Name:            diagram.Name,
				IsExpanded:      diagram.IsExpanded,
				IsNodeClickable: true,
				HasCheckboxButton: true,
				IsChecked: diagram.IsChecked,
			}
			diagramsNode.Children = append(diagramsNode.Children, diagramNode)
			diagramNode.Impl = &Diagram_Tree_DiagramProxy{
				stager:   stager,
				diagram: diagram,
			}

			for _, product := range project.RootProducts {
				stager.treePBSinDiagram(diagram, product, diagramNode)
			}
		}
	}

	if len(root.OrphanedProducts) > 0 {
		orphansProductNode := &tree.Node{Name: "Orphans Products", IsExpanded: true}
		treeInstance.RootNodes = append(treeInstance.RootNodes, orphansProductNode)
		for _, product := range root.OrphanedProducts {
			stager.treePBS(product, orphansProductNode)
		}
	}

	if len(root.OrphanedTasks) > 0 {
		orphansTaskNode := &tree.Node{Name: "Orphans Tasks", IsExpanded: true}
		treeInstance.RootNodes = append(treeInstance.RootNodes, orphansTaskNode)
		for _, task := range root.OrphanedTasks {
			stager.treeWBS(task, orphansTaskNode)
		}
	}

	tree.StageBranch(stager.treeStage, treeInstance)

	stager.treeStage.Commit()
}

func addAddItemButton[T Gongstruct, PT interface {
	*T
	GongstructIF
}](stager *Stager, node *tree.Node, items *[]*T) {

	var item PT
	addButton := &tree.Button{
		Name:            GetGongstructNameFromPointer(item) + " " + string(buttons.BUTTON_add),
		Icon:            string(buttons.BUTTON_add),
		ToolTipText:     "Add a " + GetGongstructNameFromPointer(item) + " to \"" + node.Name + "\"",
		HasToolTip:      true,
		ToolTipPosition: tree.Right,
	}
	node.Buttons = append(node.Buttons, addButton)
	addButton.Impl = &tree.FunctionalButtonProxy{
		OnUpdated: func(stage *tree.Stage, button *tree.Button, updatedButton *tree.Button) {
			item := PT(new(T))
			item.SetName("New" + GetGongstructNameFromPointer(item))
			item.StageVoid(stager.stage)
			*items = append(*items, item)

			stager.stage.Commit()
		},
	}
}

type NodeProxy[T ProjectElementType] struct {
	stager   *Stager
	node     *tree.Node
	instance T
}

// OnAfterUpdate implements models.NodeImplInterface.
func (p *NodeProxy[T]) OnAfterUpdate(stage *tree.Stage, stagedNode *tree.Node, frontNode *tree.Node) {

	if frontNode.IsExpanded != stagedNode.IsExpanded {
		stagedNode.IsExpanded = frontNode.IsExpanded
		p.instance.SetIsExpanded(!p.instance.GetIsExpanded())
		return
	}

	p.stager.probeForm.FillUpFormFromGongstruct(p.instance, GetPointerToGongstructName[T]())

	p.stager.stage.Commit()
}

// Append is a generic helper that appends an item to a slice via a pointer
func Append[T any](slice *[]T, item T) {
    *slice = append(*slice, item)
}