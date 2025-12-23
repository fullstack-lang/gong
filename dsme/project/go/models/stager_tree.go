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
		projectNode.Impl = &tree.FunctionalNodeProxy{
			OnUpdate: stager.OnUpdateProject(project),
		}

		pbsNode := &tree.Node{
			Name:            "PBS",
			FontStyle:       tree.ITALIC,
			IsExpanded:      project.IsPBSNodeExpanded,
			IsNodeClickable: true,
		}
		projectNode.Children = append(projectNode.Children, pbsNode)
		pbsNode.Impl = &tree.FunctionalNodeProxy{
			OnUpdate: stager.OnUpdateExpansion(&project.IsPBSNodeExpanded),
		}

		addAddItemButton(stager, pbsNode, &project.RootProducts)

		for _, product := range project.RootProducts {
			stager.treePBS(product, pbsNode)
		}

		wbsNode := &tree.Node{
			Name:            "WBS",
			FontStyle:       tree.ITALIC,
			IsExpanded:      project.IsWBSNodeExpanded,
			IsNodeClickable: true,
		}
		projectNode.Children = append(projectNode.Children, wbsNode)
		wbsNode.Impl = &tree.FunctionalNodeProxy{
			OnUpdate: stager.OnUpdateExpansion(&project.IsWBSNodeExpanded),
		}

		addAddItemButton(stager, wbsNode, &project.RootTasks)

		for _, task := range project.RootTasks {
			stager.treeWBS(task, wbsNode)
		}

		diagramsNode := &tree.Node{
			Name:            "Diagrams",
			FontStyle:       tree.ITALIC,
			IsExpanded:      project.IsDiagramsNodeExpanded,
			IsNodeClickable: true,
		}
		projectNode.Children = append(projectNode.Children, diagramsNode)
		diagramsNode.Impl = &tree.FunctionalNodeProxy{
			OnUpdate: stager.OnUpdateExpansion(&project.IsDiagramsNodeExpanded),
		}

		addAddItemButton(stager, diagramsNode, &project.Diagrams)

		for _, diagram := range project.Diagrams {
			diagramNode := &tree.Node{
				Name:              diagram.Name,
				IsExpanded:        diagram.IsExpanded,
				IsNodeClickable:   true,
				HasCheckboxButton: true,
				IsChecked:         diagram.IsChecked,
			}
			diagramsNode.Children = append(diagramsNode.Children, diagramNode)
			diagramNode.Impl = &tree.FunctionalNodeProxy{
				OnUpdate: stager.OnUpdateDiagram(diagram),
			}

			// computes all products presents in the diagram
			diagram.map_Product_ProductShape = make(map[*Product]*ProductShape)
			for _, shape := range diagram.Product_Shapes {
				if shape.Product != nil {
					diagram.map_Product_ProductShape[shape.Product] = shape
				}
			}
			diagram.map_Task_TaskShape = make(map[*Task]*TaskShape)
			for _, shape := range diagram.Task_Shapes {
				if shape.Task != nil {
					diagram.map_Task_TaskShape[shape.Task] = shape
				}
			}

			diagram.map_Product_ProductCompositionShape = make(map[*Product]*ProductCompositionShape)
			for _, shape := range diagram.ProductComposition_Shapes {
				if shape.Product != nil {
					diagram.map_Product_ProductCompositionShape[shape.Product] = shape
				}
			}

			pbsNode := &tree.Node{
				Name:            "PBS",
				FontStyle:       tree.ITALIC,
				IsExpanded:      diagram.IsPBSNodeExpanded,
				IsNodeClickable: true,
			}
			diagramNode.Children = append(diagramNode.Children, pbsNode)
			pbsNode.Impl = &tree.FunctionalNodeProxy{
				OnUpdate: stager.OnUpdateExpansion(&diagram.IsPBSNodeExpanded),
			}

			for _, product := range project.RootProducts {
				stager.treePBSinDiagram(diagram, product, pbsNode)
			}

			diagram.map_Task_TaskCompositionShape = make(map[*Task]*TaskCompositionShape)
			for _, shape := range diagram.TaskComposition_Shapes {
				if shape.Task != nil {
					diagram.map_Task_TaskCompositionShape[shape.Task] = shape
				}
			}

			wbsNode := &tree.Node{
				Name:            "WBS",
				FontStyle:       tree.ITALIC,
				IsExpanded:      diagram.IsWBSNodeExpanded,
				IsNodeClickable: true,
			}
			diagramNode.Children = append(diagramNode.Children, wbsNode)
			wbsNode.Impl = &tree.FunctionalNodeProxy{
				OnUpdate: stager.OnUpdateExpansion(&diagram.IsWBSNodeExpanded),
			}

			for _, task := range project.RootTasks {
				stager.treeWBSinDiagram(diagram, task, wbsNode)
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

			// check if the item is a diagram, if so, set the IsEditable_ field to true
			if diagram, ok := any(item).(*Diagram); ok {
				diagram.IsEditable_ = true
			}

			stager.stage.Commit()
		},
	}
}

// Helper callbacks

func (stager *Stager) OnUpdateProject(project *Project) func(stage *tree.Stage, stagedNode, frontNode *tree.Node) {
	return func(stage *tree.Stage, stagedNode, frontNode *tree.Node) {
		if frontNode.IsExpanded != stagedNode.IsExpanded {
			stagedNode.IsExpanded = frontNode.IsExpanded
			project.IsExpanded = frontNode.IsExpanded
		} else {
			stager.probeForm.FillUpFormFromGongstruct(project, GetPointerToGongstructName[*Project]())
		}
		stager.stage.Commit()
	}
}

func (stager *Stager) OnUpdateExpansion(isExpanded *bool) func(stage *tree.Stage, stagedNode, frontNode *tree.Node) {
	return func(stage *tree.Stage, stagedNode, frontNode *tree.Node) {
		if frontNode.IsExpanded != stagedNode.IsExpanded {
			*isExpanded = !*isExpanded
		}
		stager.stage.Commit()
	}
}

func (stager *Stager) OnUpdateDiagram(diagram *Diagram) func(stage *tree.Stage, stagedNode, frontNode *tree.Node) {
	return func(stage *tree.Stage, stagedNode, frontNode *tree.Node) {
		if frontNode.IsChecked && !stagedNode.IsChecked {
			// reset all ddiagram selection
			for diagram_ := range *GetGongstructInstancesSet[Diagram](stager.stage) {
				diagram_.IsChecked = false
			}
			diagram.IsChecked = true
			stagedNode.IsChecked = frontNode.IsChecked
			stager.stage.Commit()
			return
		}
		if !frontNode.IsChecked && stagedNode.IsChecked {
			diagram.IsChecked = false
			stagedNode.IsChecked = frontNode.IsChecked
			// reset all ddiagram selection
			for diagram_ := range *GetGongstructInstancesSet[Diagram](stager.stage) {
				diagram_.IsChecked = false
			}
			stager.stage.Commit()
			return
		}
		if frontNode.IsExpanded != stagedNode.IsExpanded {
			stagedNode.IsExpanded = frontNode.IsExpanded
			diagram.IsExpanded = frontNode.IsExpanded
			stager.stage.Commit()
			return
		}
		if frontNode.Name != stagedNode.Name {
			diagram.Name = frontNode.Name
			diagram.IsInRenameMode = false
			stager.stage.Commit()
			return
		}
		stager.probeForm.FillUpFormFromGongstruct(diagram, "Diagram")
	}
}

// Append is a generic helper that appends an item to a slice via a pointer
func Append[T any](slice *[]T, item T) {
	*slice = append(*slice, item)
}
