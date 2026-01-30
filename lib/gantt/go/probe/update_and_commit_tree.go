package probe

import (
	"fmt"
	"sort"
	"strings"

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

	stageOfInterest := probe.stageOfInterest

	// create tree
	sidebar := &tree.Tree{Name: "Sidebar"}

	topNode := &tree.Node{Name: fmt.Sprintf("%s", stageOfInterest.GetName())}

	if stageOfInterest.IsInDeltaMode() {
		topNode.Name += fmt.Sprintf(" (h %d, d %d)",
			len(stageOfInterest.GetBackwardCommits()),
			stageOfInterest.GetCommitsBehind())

		backwardButton := &tree.Button{
			Name:            "BackwardButton",
			Icon:            string(gongtree_buttons.BUTTON_arrow_back),
			HasToolTip:      true,
			ToolTipText:     "Go to previous commit",
			ToolTipPosition: tree.Below,
		}
		topNode.Buttons = append(topNode.Buttons, backwardButton)
		backwardButton.Impl = &tree.FunctionalButtonProxy{
			OnUpdated: func(stage *tree.Stage,
				stagedButton, frontButton *tree.Button) {
				err := stageOfInterest.ApplyBackwardCommit()
				if err != nil {
					panic(err)
				}
				probe.Refresh()
			},
		}

		if stageOfInterest.GetCommitsBehind() == len(stageOfInterest.GetBackwardCommits()) {
			backwardButton.IsDisabled = true
			backwardButton.Icon = string(gongtree_buttons.BUTTON_do_not_disturb)
			backwardButton.ToolTipText = "No more previous commits"
		}

		forwardButton := &tree.Button{
			Name:            "ForwardButton",
			Icon:            string(gongtree_buttons.BUTTON_arrow_forward),
			HasToolTip:      true,
			ToolTipText:     "Go to next commit",
			ToolTipPosition: tree.Below,
		}
		topNode.Buttons = append(topNode.Buttons, forwardButton)
		forwardButton.Impl = &tree.FunctionalButtonProxy{
			OnUpdated: func(stage *tree.Stage,
				stagedButton, frontButton *tree.Button) {
				err := stageOfInterest.ApplyForwardCommit()
				if err != nil {
					panic(err)
				}
				probe.Refresh()
			},
		}

		if stageOfInterest.GetCommitsBehind() == 0 {
			forwardButton.IsDisabled = true
			forwardButton.Icon = string(gongtree_buttons.BUTTON_do_not_disturb)
			forwardButton.ToolTipText = "No more next commits"
		}

		if stageOfInterest.GetCommitsBehind() > 0 {
			discardButton := &tree.Button{
				Name:            "DiscardButton",
				Icon:            string(gongtree_buttons.BUTTON_cancel),
				HasToolTip:      true,
				ToolTipText:     "Discard commits ahead (git reset --hard HEAD)",
				ToolTipPosition: tree.Below,
			}
			topNode.Buttons = append(topNode.Buttons, discardButton)
			discardButton.Impl = &tree.FunctionalButtonProxy{
				OnUpdated: func(stage *tree.Stage,
					stagedButton, frontButton *tree.Button) {
					stageOfInterest.ResetHard()
					probe.Refresh()
				},
			}
		}
	} else {
		topNode.Name += ""
	}

	sidebar.RootNodes = append(sidebar.RootNodes, topNode)
	refreshButton := &tree.Button{
		Name:            "RefreshButton" + " " + string(gongtree_buttons.BUTTON_refresh),
		Icon:            string(gongtree_buttons.BUTTON_refresh),
		HasToolTip:      true,
		ToolTipText:     "Refresh probe",
		ToolTipPosition: tree.Below,
	}
	topNode.Buttons = append(topNode.Buttons, refreshButton)
	refreshButton.Impl = &tree.FunctionalButtonProxy{
		OnUpdated: func(stage *tree.Stage,
			stagedButton, frontButton *tree.Button) {
			probe.stageOfInterest.ComputeInstancesNb()
			probe.docStager.SetMap_GongStructName_InstancesNb(
				probe.stageOfInterest.Map_GongStructName_InstancesNb,
			)
			probe.Refresh()
		},
	}

	notificationsResetButton := &tree.Button{
		Name:            "NotificationsResetButton",
		Icon:            string(gongtree_buttons.BUTTON_reset_tv),
		HasToolTip:      true,
		ToolTipText:     "Reset notification table",
		ToolTipPosition: tree.Below,
	}
	topNode.Buttons = append(topNode.Buttons, notificationsResetButton)
	notificationsResetButton.Impl = &tree.FunctionalButtonProxy{
		OnUpdated: func(stage *tree.Stage,
			stagedButton, frontButton *tree.Button) {
			probe.ResetNotifications()
		},
	}

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
			count := 0
			for _arrow := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree.Node{Name: _arrow.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_arrow, "Arrow", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Bar":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Bar](probe.stageOfInterest)
			count := 0
			for _bar := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree.Node{Name: _bar.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_bar, "Bar", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Gantt":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Gantt](probe.stageOfInterest)
			count := 0
			for _gantt := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree.Node{Name: _gantt.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_gantt, "Gantt", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Group":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Group](probe.stageOfInterest)
			count := 0
			for _group := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree.Node{Name: _group.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_group, "Group", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Lane":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Lane](probe.stageOfInterest)
			count := 0
			for _lane := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree.Node{Name: _lane.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_lane, "Lane", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "LaneUse":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.LaneUse](probe.stageOfInterest)
			count := 0
			for _laneuse := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree.Node{Name: _laneuse.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_laneuse, "LaneUse", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Milestone":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Milestone](probe.stageOfInterest)
			count := 0
			for _milestone := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree.Node{Name: _milestone.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_milestone, "Milestone", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
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
		addButton.Impl = &tree.FunctionalButtonProxy{
			OnUpdated: func(stage *tree.Stage, stagedButton, frontButton *tree.Button) {
				FillUpFormFromGongstructName(
					probe,
					gongStruct.Name,
					true,
				)
			},
		}

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
