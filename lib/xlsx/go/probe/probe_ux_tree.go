package probe

import (
	"fmt"
	"sort"
	"strings"
	"time"

	tree_buttons "github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree_models "github.com/fullstack-lang/gong/lib/tree/go/models"

	gong_models "github.com/fullstack-lang/gong/go/models"

	"github.com/fullstack-lang/gong/lib/xlsx/go/models"
)

func (probe *Probe) ux_tree() {
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

	probe.treeStage.Reset()

	stageOfInterest := probe.stageOfInterest

	// create tree
	sidebar := &tree_models.Tree{Name: "Sidebar"}
	topNode := &tree_models.Node{Name: fmt.Sprintf("%s", stageOfInterest.GetName())}
	sidebar.RootNodes = append(sidebar.RootNodes, topNode)

	notificationsResetButton := &tree_models.Button{
		Name:            "NotificationsResetButton",
		Icon:            string(tree_buttons.BUTTON_playlist_remove),
		HasToolTip:      true,
		ToolTipText:     "Reset notification table",
		ToolTipPosition: tree_models.Below,
		OnUpdate: func(_ *tree_models.Stage, _ *tree_models.Button) {
			probe.ResetNotifications()
		},
	}
	topNode.Buttons = append(topNode.Buttons, notificationsResetButton)
	refreshButton := &tree_models.Button{
		Name:            "RefreshButton" + " " + string(tree_buttons.BUTTON_refresh),
		Icon:            string(tree_buttons.BUTTON_refresh),
		HasToolTip:      true,
		ToolTipText:     "Refresh probe",
		ToolTipPosition: tree_models.Below,
		OnUpdate: func(_ *tree_models.Stage, _ *tree_models.Button) {
			probe.stageOfInterest.ComputeInstancesNb()
			probe.docStager.SetMap_GongStructName_InstancesNb(
				probe.stageOfInterest.Map_GongStructName_InstancesNb,
			)
			probe.Refresh()
		},
	}
	topNode.Buttons = append(topNode.Buttons, refreshButton)

	if stageOfInterest.IsInDeltaMode() {
		probe.AddCommitNavigationNode(func(node models.GongNodeIF) {
			sidebar.RootNodes = append(sidebar.RootNodes, node.(*tree_models.Node))
		})
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
		case "DisplaySelection":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.DisplaySelection](probe.stageOfInterest)
			count := 0
			for _displayselection := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _displayselection.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_displayselection, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.DisplaySelection](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "XLCell":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.XLCell](probe.stageOfInterest)
			count := 0
			for _xlcell := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _xlcell.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_xlcell, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.XLCell](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "XLFile":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.XLFile](probe.stageOfInterest)
			count := 0
			for _xlfile := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _xlfile.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_xlfile, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.XLFile](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "XLRow":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.XLRow](probe.stageOfInterest)
			count := 0
			for _xlrow := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _xlrow.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_xlrow, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.XLRow](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "XLSheet":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.XLSheet](probe.stageOfInterest)
			count := 0
			for _xlsheet := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _xlsheet.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_xlsheet, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.XLSheet](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
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
			OnUpdate: func(_ *tree_models.Stage, _ *tree_models.Button) {
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
		OnUpdate: func(_ *tree_models.Stage, _ *tree_models.Button) {
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

	forwardButton := &tree_models.Button{
		Name:       "ForwardButton",
		Icon:       string(tree_buttons.BUTTON_redo),
		HasToolTip: true,
		ToolTipText: fmt.Sprintf("Go to next commit (%d/%d)",
			len(stageOfInterest.GetBackwardCommits()), stageOfInterest.GetCommitsBehind()),
		ToolTipPosition: tree_models.Below,
		OnUpdate: func(_ *tree_models.Stage, _ *tree_models.Button) {
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
			OnUpdate: func(_ *tree_models.Stage, _ *tree_models.Button) {
				stageOfInterest.ResetHard()
				probe.Refresh()
			},
		}
		deltaNode.Buttons = append(deltaNode.Buttons, discardButton)
	}

	orphansButton := &tree_models.Button{
		Name:            "OrphansButton",
		Icon:            string(tree_buttons.BUTTON_playlist_remove),
		HasToolTip:      true,
		ToolTipText:     "Discard all commits history (git orphan)",
		ToolTipPosition: tree_models.Below,
		OnUpdate: func(_ *tree_models.Stage, _ *tree_models.Button) {
			stageOfInterest.Orphans()
			probe.Refresh()
		},
	}
	if len(stageOfInterest.GetBackwardCommits()) == 0 {
		orphansButton.IsDisabled = true
		orphansButton.ToolTipText = "No commits to orphan"
	}
	deltaNode.Buttons = append(deltaNode.Buttons, orphansButton)

	logCommitsButton := &tree_models.Button{
		Name:            "LogCommitsButton",
		Icon:            string(tree_buttons.BUTTON_playlist_add),
		HasToolTip:      true,
		ToolTipText:     "Log commits to notification table",
		ToolTipPosition: tree_models.Below,
		OnUpdate: func(_ *tree_models.Stage, _ *tree_models.Button) {
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
