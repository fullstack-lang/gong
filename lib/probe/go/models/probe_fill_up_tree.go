package models

import (
	"fmt"
	"strings"

	gongtree_buttons "github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

const SideBarTreeName = "gong"

func (probe *Probe2) fillUpTree() {
	// keep in memory which nodes have been unfolded / folded
	expandedNodesSet := make(map[string]any, 0)
	var _sidebar *tree.Tree
	for __sidebar := range probe.treeStage.Trees {
		_sidebar = __sidebar
	}
	if _sidebar != nil {
		for _, node := range _sidebar.RootNodes {
			if node.IsExpanded {
				expandedNodesSet[strings.Fields(node.Name)[0]] = true
			}
		}
	}

	probe.treeStage.Reset()

	// create tree
	sidebar := (&tree.Tree{Name: SideBarTreeName}).Stage(probe.treeStage)

	namedStructs := probe.stageOfInterest.GetNamedStructs()
	for _, namedStruct := range namedStructs {

		name := namedStruct.GetName() + " (" +
			fmt.Sprintf("%d", probe.stageOfInterest.GetMap_GongStructName_InstancesNb()[namedStruct.GetName()]) + ")"

		nodeNamedStruct := (&tree.Node{Name: name}).Stage(probe.treeStage)

		nodeNamedStruct.IsExpanded = false
		if _, ok := expandedNodesSet[strings.Fields(name)[0]]; ok {
			nodeNamedStruct.IsExpanded = true
		}

		for _, instance := range namedStruct.GetInstances() {
			nodeInstance := (&tree.Node{Name: instance.GetName()}).Stage(probe.treeStage)
			// nodeInstance.IsNodeClickable = true
			// nodeInstance.Impl = NewInstanceNodeCallback(_probe2, "Probe2", probe)

			nodeNamedStruct.Children = append(nodeNamedStruct.Children, nodeInstance)
		}

		nodeNamedStruct.IsNodeClickable = true
		nodeNamedStruct.Impl = NewTreeNodeImplNamedStruct(namedStruct, probe)

		// // add add button
		// addButton := (&tree.Button{
		// 	Name: gongStruct.Name + " " + string(gongtree_buttons.BUTTON_add),
		// 	Icon: string(gongtree_buttons.BUTTON_add)}).Stage(probe.treeStage)
		// nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)
		// addButton.Impl = NewButtonImplGongstruct(
		// 	gongStruct,
		// 	gongtree_buttons.BUTTON_add,
		// 	probe,
		// )

		sidebar.RootNodes = append(sidebar.RootNodes, nodeNamedStruct)
	}

	// Add a refresh button
	nodeRefreshButton := (&tree.Node{Name: ""}).Stage(probe.treeStage)
	sidebar.RootNodes = append(sidebar.RootNodes, nodeRefreshButton)
	refreshButton := (&tree.Button{
		Name: "RefreshButton" + " " + string(gongtree_buttons.BUTTON_refresh),
		Icon: string(gongtree_buttons.BUTTON_refresh)}).Stage(probe.treeStage)
	nodeRefreshButton.Buttons = append(nodeRefreshButton.Buttons, refreshButton)
	refreshButton.Impl = NewButtonImplRefresh(probe)

	probe.treeStage.Commit()
}

// type InstanceNodeCallback[T models.Gongstruct] struct {
// 	Instance       *T
// 	gongstructName string
// 	probe          *Probe
// }

// func NewInstanceNodeCallback[T models.Gongstruct](
// 	instance *T,
// 	gongstructName string,
// 	probe *Probe) (
// 	instanceNodeCallback *InstanceNodeCallback[T],
// ) {
// 	instanceNodeCallback = new(InstanceNodeCallback[T])

// 	instanceNodeCallback.probe = probe
// 	instanceNodeCallback.gongstructName = gongstructName
// 	instanceNodeCallback.Instance = instance

// 	return
// }

// func (instanceNodeCallback *InstanceNodeCallback[T]) OnAfterUpdate(
// 	gongtreeStage *tree.Stage,
// 	stagedNode, frontNode *tree.Node) {

// 	FillUpFormFromGongstruct(
// 		instanceNodeCallback.Instance,
// 		instanceNodeCallback.probe,
// 	)
// }
