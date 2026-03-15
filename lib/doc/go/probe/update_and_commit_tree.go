package probe

import (
	"fmt"
	"sort"
	"strings"
	"time"

	gongtree_buttons "github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"

	gong_models "github.com/fullstack-lang/gong/go/models"

	"github.com/fullstack-lang/gong/lib/doc/go/models"
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
	sidebar.RootNodes = append(sidebar.RootNodes, topNode)

	notificationsResetButton := &tree.Button{
		Name:            "NotificationsResetButton",
		Icon:            string(gongtree_buttons.BUTTON_playlist_remove),
		HasToolTip:      true,
		ToolTipText:     "Reset notification table",
		ToolTipPosition: tree.Below,
	}
	topNode.Buttons = append(topNode.Buttons, notificationsResetButton)
	notificationsResetButton.OnUpdate = func(_ *tree.Stage, _ *tree.Button) {
		probe.ResetNotifications()
	}
	refreshButton := &tree.Button{
		Name:            "RefreshButton" + " " + string(gongtree_buttons.BUTTON_refresh),
		Icon:            string(gongtree_buttons.BUTTON_refresh),
		HasToolTip:      true,
		ToolTipText:     "Refresh probe",
		ToolTipPosition: tree.Below,
	}
	topNode.Buttons = append(topNode.Buttons, refreshButton)
	refreshButton.OnUpdate = func(_ *tree.Stage, _ *tree.Button) {
		probe.stageOfInterest.ComputeInstancesNb()
		probe.docStager.SetMap_GongStructName_InstancesNb(
			probe.stageOfInterest.Map_GongStructName_InstancesNb,
		)
		probe.Refresh()
	}

	if stageOfInterest.IsInDeltaMode() {
		probe.AddCommitNavigationNode(func(node models.GongNodeIF) {
			sidebar.RootNodes = append(sidebar.RootNodes, node.(*tree.Node))
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
		case "AttributeShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.AttributeShape](probe.stageOfInterest)
			count := 0
			for _attributeshape := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree.Node{Name: _attributeshape.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.OnUpdate = func(_ *tree.Stage, _, _ *tree.Node) {
					FillUpFormFromGongstruct(_attributeshape, probe)
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree.Stage, stagedNode, frontNode *tree.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.AttributeShape](probe)
				// set color for node and reset all other nodes color
				for node := range *tree.GetGongstructInstancesSet[tree.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "Classdiagram":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Classdiagram](probe.stageOfInterest)
			count := 0
			for _classdiagram := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree.Node{Name: _classdiagram.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.OnUpdate = func(_ *tree.Stage, _, _ *tree.Node) {
					FillUpFormFromGongstruct(_classdiagram, probe)
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree.Stage, stagedNode, frontNode *tree.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.Classdiagram](probe)
				// set color for node and reset all other nodes color
				for node := range *tree.GetGongstructInstancesSet[tree.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "DiagramPackage":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.DiagramPackage](probe.stageOfInterest)
			count := 0
			for _diagrampackage := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree.Node{Name: _diagrampackage.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.OnUpdate = func(_ *tree.Stage, _, _ *tree.Node) {
					FillUpFormFromGongstruct(_diagrampackage, probe)
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree.Stage, stagedNode, frontNode *tree.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.DiagramPackage](probe)
				// set color for node and reset all other nodes color
				for node := range *tree.GetGongstructInstancesSet[tree.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "GongEnumShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.GongEnumShape](probe.stageOfInterest)
			count := 0
			for _gongenumshape := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree.Node{Name: _gongenumshape.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.OnUpdate = func(_ *tree.Stage, _, _ *tree.Node) {
					FillUpFormFromGongstruct(_gongenumshape, probe)
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree.Stage, stagedNode, frontNode *tree.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.GongEnumShape](probe)
				// set color for node and reset all other nodes color
				for node := range *tree.GetGongstructInstancesSet[tree.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "GongEnumValueShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.GongEnumValueShape](probe.stageOfInterest)
			count := 0
			for _gongenumvalueshape := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree.Node{Name: _gongenumvalueshape.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.OnUpdate = func(_ *tree.Stage, _, _ *tree.Node) {
					FillUpFormFromGongstruct(_gongenumvalueshape, probe)
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree.Stage, stagedNode, frontNode *tree.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.GongEnumValueShape](probe)
				// set color for node and reset all other nodes color
				for node := range *tree.GetGongstructInstancesSet[tree.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "GongNoteLinkShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.GongNoteLinkShape](probe.stageOfInterest)
			count := 0
			for _gongnotelinkshape := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree.Node{Name: _gongnotelinkshape.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.OnUpdate = func(_ *tree.Stage, _, _ *tree.Node) {
					FillUpFormFromGongstruct(_gongnotelinkshape, probe)
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree.Stage, stagedNode, frontNode *tree.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.GongNoteLinkShape](probe)
				// set color for node and reset all other nodes color
				for node := range *tree.GetGongstructInstancesSet[tree.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "GongNoteShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.GongNoteShape](probe.stageOfInterest)
			count := 0
			for _gongnoteshape := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree.Node{Name: _gongnoteshape.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.OnUpdate = func(_ *tree.Stage, _, _ *tree.Node) {
					FillUpFormFromGongstruct(_gongnoteshape, probe)
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree.Stage, stagedNode, frontNode *tree.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.GongNoteShape](probe)
				// set color for node and reset all other nodes color
				for node := range *tree.GetGongstructInstancesSet[tree.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "GongStructShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.GongStructShape](probe.stageOfInterest)
			count := 0
			for _gongstructshape := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree.Node{Name: _gongstructshape.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.OnUpdate = func(_ *tree.Stage, _, _ *tree.Node) {
					FillUpFormFromGongstruct(_gongstructshape, probe)
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree.Stage, stagedNode, frontNode *tree.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.GongStructShape](probe)
				// set color for node and reset all other nodes color
				for node := range *tree.GetGongstructInstancesSet[tree.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "LinkShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.LinkShape](probe.stageOfInterest)
			count := 0
			for _linkshape := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree.Node{Name: _linkshape.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.OnUpdate = func(_ *tree.Stage, _, _ *tree.Node) {
					FillUpFormFromGongstruct(_linkshape, probe)
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree.Stage, stagedNode, frontNode *tree.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.LinkShape](probe)
				// set color for node and reset all other nodes color
				for node := range *tree.GetGongstructInstancesSet[tree.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		}

		nodeGongstruct.IsNodeClickable = true

		// add add button
		addButton := &tree.Button{
			Name:            gongStruct.Name + " " + string(gongtree_buttons.BUTTON_add),
			Icon:            string(gongtree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of " + gongStruct.GetName(),
			ToolTipPosition: tree.Right,
			OnUpdate: func(_ *tree.Stage, _ *tree.Button) {
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

	tree.StageBranch(probe.treeStage, sidebar)

	probe.treeStage.Commit()
}

func (probe *Probe) AddCommitNavigationNode(appendChildrenNodeFunc func(models.GongNodeIF)) {
	stageOfInterest := probe.stageOfInterest

	deltaNode := &tree.Node{}

	backwardButton := &tree.Button{
		Name:       "BackwardButton",
		Icon:       string(gongtree_buttons.BUTTON_undo),
		HasToolTip: true,
		ToolTipText: fmt.Sprintf("Go to previous commit (%d/%d)",
			len(stageOfInterest.GetBackwardCommits()), stageOfInterest.GetCommitsBehind()),
		ToolTipPosition: tree.Below,
	}
	deltaNode.Buttons = append(deltaNode.Buttons, backwardButton)
	backwardButton.Impl = &tree.FunctionalButtonProxy{
		OnUpdated: func(stage *tree.Stage,
			stagedButton, frontButton *tree.Button,
		) {
			err := stageOfInterest.ApplyBackwardCommit()
			if err != nil {
				panic(err)
			}
			probe.Refresh()
		},
	}

	if stageOfInterest.GetCommitsBehind() == len(stageOfInterest.GetBackwardCommits()) {
		backwardButton.IsDisabled = true
		backwardButton.ToolTipText = "No more previous commits"
	}

	forwardButton := &tree.Button{
		Name:       "ForwardButton",
		Icon:       string(gongtree_buttons.BUTTON_redo),
		HasToolTip: true,
		ToolTipText: fmt.Sprintf("Go to next commit (%d/%d)",
			len(stageOfInterest.GetBackwardCommits()), stageOfInterest.GetCommitsBehind()),
		ToolTipPosition: tree.Below,
	}
	deltaNode.Buttons = append(deltaNode.Buttons, forwardButton)
	forwardButton.Impl = &tree.FunctionalButtonProxy{
		OnUpdated: func(stage *tree.Stage,
			stagedButton, frontButton *tree.Button,
		) {
			err := stageOfInterest.ApplyForwardCommit()
			if err != nil {
				panic(err)
			}
			probe.Refresh()
		},
	}

	if stageOfInterest.GetCommitsBehind() == 0 {
		forwardButton.IsDisabled = true
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
		deltaNode.Buttons = append(deltaNode.Buttons, discardButton)
		discardButton.Impl = &tree.FunctionalButtonProxy{
			OnUpdated: func(stage *tree.Stage,
				stagedButton, frontButton *tree.Button,
			) {
				stageOfInterest.ResetHard()
				probe.Refresh()
			},
		}
	}


	orphansButton := &tree.Button{
		Name:            "OrphansButton",
		Icon:            string(gongtree_buttons.BUTTON_playlist_remove),
		HasToolTip:      true,
		ToolTipText:     "Discard all commits history (git orphan)",
		ToolTipPosition: tree.Below,
	}
	if len(stageOfInterest.GetBackwardCommits()) == 0 {
		orphansButton.IsDisabled = true
		orphansButton.ToolTipText = "No commits to orphan"
	}
	deltaNode.Buttons = append(deltaNode.Buttons, orphansButton)
	orphansButton.Impl = &tree.FunctionalButtonProxy{
		OnUpdated: func(stage *tree.Stage,
			stagedButton, frontButton *tree.Button,
		) {
			stageOfInterest.Orphans()
			probe.Refresh()
		},
	}

	logCommitsButton := &tree.Button{
		Name:            "LogCommitsButton",
		Icon:            string(gongtree_buttons.BUTTON_playlist_add),
		HasToolTip:      true,
		ToolTipText:     "Log commits to notification table",
		ToolTipPosition: tree.Below,
	}
	if len(stageOfInterest.GetBackwardCommits()) == 0 {
		logCommitsButton.IsDisabled = true
		logCommitsButton.ToolTipText = "No commits to log"
	}
	deltaNode.Buttons = append(deltaNode.Buttons, logCommitsButton)
	logCommitsButton.Impl = &tree.FunctionalButtonProxy{
		OnUpdated: func(stage *tree.Stage,
			stagedButton, frontButton *tree.Button,
		) {
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

	appendChildrenNodeFunc(deltaNode)
}
