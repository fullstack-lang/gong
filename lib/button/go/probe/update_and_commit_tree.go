package probe

import (
	"fmt"
	"sort"
	"strings"
	"time"

	gongtree_buttons "github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"

	gong_models "github.com/fullstack-lang/gong/go/models"

	"github.com/fullstack-lang/gong/lib/button/go/models"
)

func updateAndCommitTree(
	probe *Probe,
) {
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
	sidebar := &tree.Tree{Name: "Sidebar"}

	nodeRefreshButton := &tree.Node{Name: fmt.Sprintf("Stage %s, # %d, %s",
		probe.stageOfInterest.GetName(),
		probe.stageOfInterest.GetCommitId(),
		probe.stageOfInterest.GetCommitTS().Local().Format(time.Kitchen))}
	sidebar.RootNodes = append(sidebar.RootNodes, nodeRefreshButton)
	refreshButton := &tree.Button{
		Name:            "RefreshButton" + " " + string(gongtree_buttons.BUTTON_refresh),
		Icon:            string(gongtree_buttons.BUTTON_refresh),
		HasToolTip:      true,
		ToolTipText:     "Refresh probe",
		ToolTipPosition: tree.Left,
	}

	nodeRefreshButton.Buttons = append(nodeRefreshButton.Buttons, refreshButton)
	refreshButton.Impl = NewButtonImplRefresh(probe)

	// collect all gong struct to construe the true
	setOfGongStructs := *gong_models.GetGongstructInstancesSetFromPointerType[*gong_models.GongStruct](probe.gongStage)

	sliceOfGongStructsSorted := make([]*gong_models.GongStruct, len(setOfGongStructs))
	i := 0
	for k := range setOfGongStructs {
		sliceOfGongStructsSorted[i] = k
		i++
	}
	sort.Slice(sliceOfGongStructsSorted, func(i, j int) bool {
		return sliceOfGongStructsSorted[i].Name < sliceOfGongStructsSorted[j].Name
	})

	for _, gongStruct := range sliceOfGongStructsSorted {

		name := gongStruct.Name + " (" +
			fmt.Sprintf("%d", probe.stageOfInterest.Map_GongStructName_InstancesNb[gongStruct.Name]) + ")"

		nodeGongstruct := &tree.Node{Name: name}
		nodeGongstruct.HasToolTip = true
		nodeGongstruct.ToolTipText = "Display table of all " + name + " instances"
		nodeGongstruct.ToolTipPosition = tree.Right

		nodeGongstruct.IsExpanded = false
		if _, ok := expandedNodesSet[strings.Fields(name)[0]]; ok {
			nodeGongstruct.IsExpanded = true
		}

		switch gongStruct.Name {
		// insertion point
		case "Button":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Button](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _button := range set {
				nodeInstance := &tree.Node{Name: _button.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_button, "Button", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_button]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_button]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_button]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (C%d/U%d/D%d)", created, updated, deleted)
		case "Group":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Group](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _group := range set {
				nodeInstance := &tree.Node{Name: _group.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_group, "Group", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_group]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_group]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_group]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (C%d/U%d/D%d)", created, updated, deleted)
		case "Layout":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Layout](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _layout := range set {
				nodeInstance := &tree.Node{Name: _layout.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_layout, "Layout", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_layout]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_layout]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_layout]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (C%d/U%d/D%d)", created, updated, deleted)
		}

		nodeGongstruct.IsNodeClickable = true
		nodeGongstruct.Impl = NewTreeNodeImplGongstruct(gongStruct, probe)

		// add add button
		addButton := &tree.Button{
			Name:            gongStruct.Name + " " + string(gongtree_buttons.BUTTON_add),
			Icon:            string(gongtree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of " + gongStruct.GetName(),
			ToolTipPosition: tree.Right,
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)
		addButton.Impl = NewButtonImplGongstruct(
			gongStruct,
			gongtree_buttons.BUTTON_add,
			probe,
		)

		sidebar.RootNodes = append(sidebar.RootNodes, nodeGongstruct)
	}

	tree.StageBranch(probe.treeStage, sidebar)

	probe.treeStage.Commit()
}

type InstanceNodeCallback[T models.PointerToGongstruct] struct {
	Instance       T
	gongstructName string
	probe          *Probe
}

func NewInstanceNodeCallback[T models.PointerToGongstruct](
	instance T,
	gongstructName string,
	probe *Probe) (
	instanceNodeCallback *InstanceNodeCallback[T],
) {
	instanceNodeCallback = new(InstanceNodeCallback[T])

	instanceNodeCallback.probe = probe
	instanceNodeCallback.gongstructName = gongstructName
	instanceNodeCallback.Instance = instance

	return
}

func (instanceNodeCallback *InstanceNodeCallback[T]) OnAfterUpdate(
	gongtreeStage *tree.Stage,
	stagedNode, frontNode *tree.Node) {

	FillUpFormFromGongstruct(
		instanceNodeCallback.Instance,
		instanceNodeCallback.probe,
	)
}
