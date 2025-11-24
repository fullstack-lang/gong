package probe

import (
	"fmt"
	"sort"
	"strings"
	"time"

	gongtree_buttons "github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"

	gong_models "github.com/fullstack-lang/gong/go/models"

	"github.com/fullstack-lang/gong/lib/gantt/go/models"
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
	nodeRefreshButton.Name +=
		fmt.Sprintf(" (C%d/U%d/D%d)", 
			len(probe.stageOfInterest.GetNew()), 
			len(probe.stageOfInterest.GetModified()), 
			len(probe.stageOfInterest.GetDeleted()),
		)
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
		case "Arrow":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Arrow](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _arrow := range set {
				nodeInstance := &tree.Node{Name: _arrow.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_arrow, "Arrow", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_arrow]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_arrow]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_arrow]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (C%d/U%d/D%d)", created, updated, deleted)
		case "Bar":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Bar](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _bar := range set {
				nodeInstance := &tree.Node{Name: _bar.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_bar, "Bar", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_bar]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_bar]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_bar]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (C%d/U%d/D%d)", created, updated, deleted)
		case "Gantt":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Gantt](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _gantt := range set {
				nodeInstance := &tree.Node{Name: _gantt.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_gantt, "Gantt", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_gantt]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_gantt]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_gantt]; ok {
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
		case "Lane":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Lane](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _lane := range set {
				nodeInstance := &tree.Node{Name: _lane.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_lane, "Lane", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_lane]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_lane]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_lane]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (C%d/U%d/D%d)", created, updated, deleted)
		case "LaneUse":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.LaneUse](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _laneuse := range set {
				nodeInstance := &tree.Node{Name: _laneuse.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_laneuse, "LaneUse", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_laneuse]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_laneuse]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_laneuse]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (C%d/U%d/D%d)", created, updated, deleted)
		case "Milestone":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Milestone](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _milestone := range set {
				nodeInstance := &tree.Node{Name: _milestone.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_milestone, "Milestone", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_milestone]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_milestone]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_milestone]; ok {
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
