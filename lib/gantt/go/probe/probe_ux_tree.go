package probe

import (
	"fmt"
	"sort"
	"strings"
	"time"

	tree_buttons "github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree_models "github.com/fullstack-lang/gong/lib/tree/go/models"

	gong_models "github.com/fullstack-lang/gong/go/models"

	"github.com/fullstack-lang/gong/lib/gantt/go/models"
)

func (probe *Probe) ux_navigation_tree() {
	stageOfInterest := probe.stageOfInterest
	if !stageOfInterest.IsInDeltaMode() {
		return
	}
	probe.treeNavigationStage.Reset()

	sidebar := &tree_models.Tree{Name: "Sidebar"}

	probe.AddCommitNavigationNode(func(node models.GongNodeIF) {
		sidebar.RootNodes = append(sidebar.RootNodes, node.(*tree_models.Node))
	})

	tree_models.StageBranch(probe.treeNavigationStage, sidebar)

	probe.treeNavigationStage.Commit()
}

func (probe *Probe) ux_tree() {
	probe.ux_navigation_tree()
	probe.treeStage.Reset()
	stageOfInterest := probe.stageOfInterest

	// keep in memory which nodes have been unfolded / folded
	expandedNodesSet := make(map[string]any, 0)
	var _sidebar *tree_models.Tree
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

	// create tree
	sidebar := &tree_models.Tree{Name: "Sidebar"}
	topNode := &tree_models.Node{Name: fmt.Sprintf("%s", stageOfInterest.GetName())}
	sidebar.RootNodes = append(sidebar.RootNodes, topNode)

	refreshButton := &tree_models.Button{
		Name:            "RefreshButton" + " " + string(tree_buttons.BUTTON_refresh),
		Icon:            string(tree_buttons.BUTTON_refresh),
		HasToolTip:      true,
		ToolTipText:     "Refresh probe",
		ToolTipPosition: tree_models.Below,
		OnClick: func() {
			probe.stageOfInterest.ComputeInstancesNb()
			probe.docStager.SetMap_GongStructName_InstancesNb(
				probe.stageOfInterest.Map_GongStructName_InstancesNb,
			)
			probe.Refresh()
		},
	}
	topNode.Buttons = append(topNode.Buttons, refreshButton)

	exportExcelButton := &tree_models.Button{
		Name:            "ExportExcelButton" + " " + string(tree_buttons.BUTTON_download),
		Icon:            string(tree_buttons.BUTTON_download),
		HasToolTip:      true,
		ToolTipText:     "Export stage as Excel",
		ToolTipPosition: tree_models.Below,
		OnClick: func() {
			probe.ExportStageExcel()
		},
	}
	topNode.Buttons = append(topNode.Buttons, exportExcelButton)

	exportGoButton := &tree_models.Button{
		Name:            "ExportGoButton" + " " + string(tree_buttons.BUTTON_file_download),
		Icon:            string(tree_buttons.BUTTON_file_download),
		HasToolTip:      true,
		ToolTipText:     "Export stage as Go file",
		ToolTipPosition: tree_models.Below,
		OnClick: func() {
			probe.ExportStage()
		},
	}
	topNode.Buttons = append(topNode.Buttons, exportGoButton)

	resetButton := &tree_models.Button{
		Name:            "ResetButton" + " " + string(tree_buttons.BUTTON_reset_tv),
		Icon:            string(tree_buttons.BUTTON_reset_tv),
		HasToolTip:      true,
		ToolTipText:     "Reset stage",
		ToolTipPosition: tree_models.Below,
		OnClick: func() {
			probe.stageOfInterest.Reset()
			probe.stageOfInterest.Commit()
		},
	}
	topNode.Buttons = append(topNode.Buttons, resetButton)

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

		nodeGongstruct := &tree_models.Node{Name: name}
		nodeGongstruct.HasToolTip = true
		nodeGongstruct.ToolTipText = "Display table of all " + name + " instances"
		nodeGongstruct.ToolTipPosition = tree_models.Right

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
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _arrow.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_arrow, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.Arrow](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "Bar":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Bar](probe.stageOfInterest)
			count := 0
			for _bar := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _bar.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_bar, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.Bar](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "Gantt":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Gantt](probe.stageOfInterest)
			count := 0
			for _gantt := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _gantt.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_gantt, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.Gantt](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "Group":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Group](probe.stageOfInterest)
			count := 0
			for _group := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _group.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_group, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.Group](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "Lane":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Lane](probe.stageOfInterest)
			count := 0
			for _lane := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _lane.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_lane, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.Lane](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "LaneUse":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.LaneUse](probe.stageOfInterest)
			count := 0
			for _laneuse := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _laneuse.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_laneuse, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.LaneUse](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "Milestone":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Milestone](probe.stageOfInterest)
			count := 0
			for _milestone := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _milestone.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_milestone, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.Milestone](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		}

		nodeGongstruct.IsNodeClickable = true

		// add add button
		addButton := &tree_models.Button{
			Name:            gongStruct.Name + " " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of " + gongStruct.GetName(),
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				FillUpFormFromGongstructName(
					probe,
					gongStruct.Name,
					true,
				)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)
		sidebar.RootNodes = append(sidebar.RootNodes, nodeGongstruct)
	}

	tree_models.StageBranch(probe.treeStage, sidebar)

	probe.treeStage.Commit()
}

func (probe *Probe) AddCommitNavigationNode(appendChildrenNodeFunc func(models.GongNodeIF)) {
	stageOfInterest := probe.stageOfInterest

	deltaNode := &tree_models.Node{}

	backwardButton := &tree_models.Button{
		Name:       "BackwardButton",
		Icon:       string(tree_buttons.BUTTON_undo),
		HasToolTip: true,
		ToolTipText: fmt.Sprintf("Go to previous commit (%d/%d)",
			len(stageOfInterest.GetBackwardCommits()), stageOfInterest.GetCommitsBehind()),
		ToolTipPosition: tree_models.Below,
		OnClick: func() {
			err := stageOfInterest.ApplyBackwardCommit()
			if err != nil {
				panic(err)
			}
			probe.Refresh()
		},
	}
	deltaNode.Buttons = append(deltaNode.Buttons, backwardButton)

	if stageOfInterest.GetCommitsBehind() == len(stageOfInterest.GetBackwardCommits()) {
		backwardButton.IsDisabled = true
		backwardButton.ToolTipText = "No more previous commits"
	}

	if stageOfInterest.GetIsWithGenesisCommit() &&
		len(stageOfInterest.GetBackwardCommits()) > 0 &&
		stageOfInterest.GetCommitsBehind() >= len(stageOfInterest.GetBackwardCommits())-1 {
		backwardButton.IsDisabled = true
		backwardButton.ToolTipText = "Cannot rollback genesis commit"
	}

	forwardButton := &tree_models.Button{
		Name:       "ForwardButton",
		Icon:       string(tree_buttons.BUTTON_redo),
		HasToolTip: true,
		ToolTipText: fmt.Sprintf("Go to next commit (%d/%d)",
			len(stageOfInterest.GetBackwardCommits()), stageOfInterest.GetCommitsBehind()),
		ToolTipPosition: tree_models.Below,
		OnClick: func() {
			err := stageOfInterest.ApplyForwardCommit()
			if err != nil {
				panic(err)
			}
			probe.Refresh()
		},
	}
	deltaNode.Buttons = append(deltaNode.Buttons, forwardButton)

	if stageOfInterest.GetCommitsBehind() == 0 {
		forwardButton.IsDisabled = true
		forwardButton.ToolTipText = "No more next commits"
	}

	if stageOfInterest.GetCommitsBehind() > 0 {
		discardButton := &tree_models.Button{
			Name:            "DiscardButton",
			Icon:            string(tree_buttons.BUTTON_cancel),
			HasToolTip:      true,
			ToolTipText:     "Discard commits ahead (git reset --hard HEAD)",
			ToolTipPosition: tree_models.Below,
			OnClick: func() {
				stageOfInterest.ResetHard()
				probe.Refresh()
			},
		}
		deltaNode.Buttons = append(deltaNode.Buttons, discardButton)
	}

	squashButton := &tree_models.Button{
		Name:            "SquashButton",
		Icon:            string(tree_buttons.BUTTON_playlist_remove),
		HasToolTip:      true,
		ToolTipText:     "Squash all commits history",
		ToolTipPosition: tree_models.Below,
		OnClick: func() {
			stageOfInterest.Squash()
			probe.Refresh()
		},
	}
	if len(stageOfInterest.GetBackwardCommits()) == 0 {
		squashButton.IsDisabled = true
		squashButton.ToolTipText = "No commits to squash"
	}
	deltaNode.Buttons = append(deltaNode.Buttons, squashButton)

	logCommitsButton := &tree_models.Button{
		Name:            "LogCommitsButton",
		Icon:            string(tree_buttons.BUTTON_playlist_add),
		HasToolTip:      true,
		ToolTipText:     "Log commits to notification table",
		ToolTipPosition: tree_models.Below,
		OnClick: func() {
			var mergedCommits string
			for _, commit := range stageOfInterest.GetForwardCommits() {
				mergedCommits += commit
			}
			probe.AddNotification(
				time.Now(),
				"	// Forward commits:\n"+
					mergedCommits,
			)

			var reverseMergedCommits string
			for _, reverserCommit := range stageOfInterest.GetBackwardCommits() {
				reverseMergedCommits += reverserCommit
			}
			probe.AddNotification(
				time.Now(),
				"	// Backward commits:\n"+
					reverseMergedCommits,
			)

			probe.CommitNotificationTable()
		},
	}
	if len(stageOfInterest.GetBackwardCommits()) == 0 {
		logCommitsButton.IsDisabled = true
		logCommitsButton.ToolTipText = "No commits to log"
	}
	deltaNode.Buttons = append(deltaNode.Buttons, logCommitsButton)

	appendChildrenNodeFunc(deltaNode)
}
