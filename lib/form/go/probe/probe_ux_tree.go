package probe

import (
	"fmt"
	"sort"
	"strings"
	"time"

	tree_buttons "github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree_models "github.com/fullstack-lang/gong/lib/tree/go/models"

	gong_models "github.com/fullstack-lang/gong/go/models"

	"github.com/fullstack-lang/gong/lib/form/go/models"
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
		case "CheckBox":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.CheckBox](probe.stageOfInterest)
			count := 0
			for _checkbox := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _checkbox.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_checkbox, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.CheckBox](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "Form2":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Form2](probe.stageOfInterest)
			count := 0
			for _form2 := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _form2.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_form2, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.Form2](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "FormDiv":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.FormDiv](probe.stageOfInterest)
			count := 0
			for _formdiv := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _formdiv.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_formdiv, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.FormDiv](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "FormEditAssocButton":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.FormEditAssocButton](probe.stageOfInterest)
			count := 0
			for _formeditassocbutton := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _formeditassocbutton.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_formeditassocbutton, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.FormEditAssocButton](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "FormField":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.FormField](probe.stageOfInterest)
			count := 0
			for _formfield := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _formfield.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_formfield, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.FormField](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "FormFieldDate":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.FormFieldDate](probe.stageOfInterest)
			count := 0
			for _formfielddate := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _formfielddate.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_formfielddate, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.FormFieldDate](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "FormFieldDateTime":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.FormFieldDateTime](probe.stageOfInterest)
			count := 0
			for _formfielddatetime := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _formfielddatetime.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_formfielddatetime, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.FormFieldDateTime](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "FormFieldFloat64":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.FormFieldFloat64](probe.stageOfInterest)
			count := 0
			for _formfieldfloat64 := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _formfieldfloat64.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_formfieldfloat64, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.FormFieldFloat64](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "FormFieldInt":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.FormFieldInt](probe.stageOfInterest)
			count := 0
			for _formfieldint := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _formfieldint.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_formfieldint, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.FormFieldInt](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "FormFieldSelect":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.FormFieldSelect](probe.stageOfInterest)
			count := 0
			for _formfieldselect := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _formfieldselect.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_formfieldselect, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.FormFieldSelect](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "FormFieldString":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.FormFieldString](probe.stageOfInterest)
			count := 0
			for _formfieldstring := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _formfieldstring.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_formfieldstring, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.FormFieldString](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "FormFieldTime":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.FormFieldTime](probe.stageOfInterest)
			count := 0
			for _formfieldtime := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _formfieldtime.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_formfieldtime, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.FormFieldTime](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "FormGroup":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.FormGroup](probe.stageOfInterest)
			count := 0
			for _formgroup := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _formgroup.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_formgroup, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.FormGroup](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "FormSortAssocButton":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.FormSortAssocButton](probe.stageOfInterest)
			count := 0
			for _formsortassocbutton := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _formsortassocbutton.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_formsortassocbutton, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.FormSortAssocButton](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "Option":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Option](probe.stageOfInterest)
			count := 0
			for _option := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _option.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_option, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.Option](probe)
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
