package probe

import (
	"fmt"
	"sort"
	"strings"
	"time"

	gongtree_buttons "github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"

	gong_models "github.com/fullstack-lang/gong/go/models"

	"github.com/fullstack-lang/gong/lib/table/go/models"
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
		case "Cell":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Cell](probe.stageOfInterest)
			count := 0
			for _cell := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree.Node{Name: _cell.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.OnUpdate = func(_ *tree.Stage, _, _ *tree.Node) {
					FillUpFormFromGongstruct(_cell, probe)
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree.Stage, stagedNode, frontNode *tree.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.Cell](probe)
				// set color for node and reset all other nodes color
				for node := range *tree.GetGongstructInstancesSet[tree.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "CellBoolean":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.CellBoolean](probe.stageOfInterest)
			count := 0
			for _cellboolean := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree.Node{Name: _cellboolean.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.OnUpdate = func(_ *tree.Stage, _, _ *tree.Node) {
					FillUpFormFromGongstruct(_cellboolean, probe)
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree.Stage, stagedNode, frontNode *tree.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.CellBoolean](probe)
				// set color for node and reset all other nodes color
				for node := range *tree.GetGongstructInstancesSet[tree.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "CellFloat64":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.CellFloat64](probe.stageOfInterest)
			count := 0
			for _cellfloat64 := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree.Node{Name: _cellfloat64.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.OnUpdate = func(_ *tree.Stage, _, _ *tree.Node) {
					FillUpFormFromGongstruct(_cellfloat64, probe)
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree.Stage, stagedNode, frontNode *tree.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.CellFloat64](probe)
				// set color for node and reset all other nodes color
				for node := range *tree.GetGongstructInstancesSet[tree.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "CellIcon":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.CellIcon](probe.stageOfInterest)
			count := 0
			for _cellicon := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree.Node{Name: _cellicon.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.OnUpdate = func(_ *tree.Stage, _, _ *tree.Node) {
					FillUpFormFromGongstruct(_cellicon, probe)
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree.Stage, stagedNode, frontNode *tree.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.CellIcon](probe)
				// set color for node and reset all other nodes color
				for node := range *tree.GetGongstructInstancesSet[tree.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "CellInt":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.CellInt](probe.stageOfInterest)
			count := 0
			for _cellint := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree.Node{Name: _cellint.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.OnUpdate = func(_ *tree.Stage, _, _ *tree.Node) {
					FillUpFormFromGongstruct(_cellint, probe)
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree.Stage, stagedNode, frontNode *tree.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.CellInt](probe)
				// set color for node and reset all other nodes color
				for node := range *tree.GetGongstructInstancesSet[tree.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "CellString":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.CellString](probe.stageOfInterest)
			count := 0
			for _cellstring := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree.Node{Name: _cellstring.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.OnUpdate = func(_ *tree.Stage, _, _ *tree.Node) {
					FillUpFormFromGongstruct(_cellstring, probe)
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree.Stage, stagedNode, frontNode *tree.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.CellString](probe)
				// set color for node and reset all other nodes color
				for node := range *tree.GetGongstructInstancesSet[tree.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "CheckBox":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.CheckBox](probe.stageOfInterest)
			count := 0
			for _checkbox := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree.Node{Name: _checkbox.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.OnUpdate = func(_ *tree.Stage, _, _ *tree.Node) {
					FillUpFormFromGongstruct(_checkbox, probe)
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree.Stage, stagedNode, frontNode *tree.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.CheckBox](probe)
				// set color for node and reset all other nodes color
				for node := range *tree.GetGongstructInstancesSet[tree.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "DisplayedColumn":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.DisplayedColumn](probe.stageOfInterest)
			count := 0
			for _displayedcolumn := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree.Node{Name: _displayedcolumn.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.OnUpdate = func(_ *tree.Stage, _, _ *tree.Node) {
					FillUpFormFromGongstruct(_displayedcolumn, probe)
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree.Stage, stagedNode, frontNode *tree.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.DisplayedColumn](probe)
				// set color for node and reset all other nodes color
				for node := range *tree.GetGongstructInstancesSet[tree.Node](treeStagee) {
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
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree.Node{Name: _formdiv.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.OnUpdate = func(_ *tree.Stage, _, _ *tree.Node) {
					FillUpFormFromGongstruct(_formdiv, probe)
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree.Stage, stagedNode, frontNode *tree.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.FormDiv](probe)
				// set color for node and reset all other nodes color
				for node := range *tree.GetGongstructInstancesSet[tree.Node](treeStagee) {
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
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree.Node{Name: _formeditassocbutton.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.OnUpdate = func(_ *tree.Stage, _, _ *tree.Node) {
					FillUpFormFromGongstruct(_formeditassocbutton, probe)
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree.Stage, stagedNode, frontNode *tree.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.FormEditAssocButton](probe)
				// set color for node and reset all other nodes color
				for node := range *tree.GetGongstructInstancesSet[tree.Node](treeStagee) {
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
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree.Node{Name: _formfield.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.OnUpdate = func(_ *tree.Stage, _, _ *tree.Node) {
					FillUpFormFromGongstruct(_formfield, probe)
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree.Stage, stagedNode, frontNode *tree.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.FormField](probe)
				// set color for node and reset all other nodes color
				for node := range *tree.GetGongstructInstancesSet[tree.Node](treeStagee) {
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
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree.Node{Name: _formfielddate.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.OnUpdate = func(_ *tree.Stage, _, _ *tree.Node) {
					FillUpFormFromGongstruct(_formfielddate, probe)
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree.Stage, stagedNode, frontNode *tree.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.FormFieldDate](probe)
				// set color for node and reset all other nodes color
				for node := range *tree.GetGongstructInstancesSet[tree.Node](treeStagee) {
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
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree.Node{Name: _formfielddatetime.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.OnUpdate = func(_ *tree.Stage, _, _ *tree.Node) {
					FillUpFormFromGongstruct(_formfielddatetime, probe)
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree.Stage, stagedNode, frontNode *tree.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.FormFieldDateTime](probe)
				// set color for node and reset all other nodes color
				for node := range *tree.GetGongstructInstancesSet[tree.Node](treeStagee) {
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
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree.Node{Name: _formfieldfloat64.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.OnUpdate = func(_ *tree.Stage, _, _ *tree.Node) {
					FillUpFormFromGongstruct(_formfieldfloat64, probe)
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree.Stage, stagedNode, frontNode *tree.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.FormFieldFloat64](probe)
				// set color for node and reset all other nodes color
				for node := range *tree.GetGongstructInstancesSet[tree.Node](treeStagee) {
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
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree.Node{Name: _formfieldint.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.OnUpdate = func(_ *tree.Stage, _, _ *tree.Node) {
					FillUpFormFromGongstruct(_formfieldint, probe)
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree.Stage, stagedNode, frontNode *tree.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.FormFieldInt](probe)
				// set color for node and reset all other nodes color
				for node := range *tree.GetGongstructInstancesSet[tree.Node](treeStagee) {
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
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree.Node{Name: _formfieldselect.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.OnUpdate = func(_ *tree.Stage, _, _ *tree.Node) {
					FillUpFormFromGongstruct(_formfieldselect, probe)
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree.Stage, stagedNode, frontNode *tree.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.FormFieldSelect](probe)
				// set color for node and reset all other nodes color
				for node := range *tree.GetGongstructInstancesSet[tree.Node](treeStagee) {
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
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree.Node{Name: _formfieldstring.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.OnUpdate = func(_ *tree.Stage, _, _ *tree.Node) {
					FillUpFormFromGongstruct(_formfieldstring, probe)
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree.Stage, stagedNode, frontNode *tree.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.FormFieldString](probe)
				// set color for node and reset all other nodes color
				for node := range *tree.GetGongstructInstancesSet[tree.Node](treeStagee) {
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
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree.Node{Name: _formfieldtime.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.OnUpdate = func(_ *tree.Stage, _, _ *tree.Node) {
					FillUpFormFromGongstruct(_formfieldtime, probe)
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree.Stage, stagedNode, frontNode *tree.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.FormFieldTime](probe)
				// set color for node and reset all other nodes color
				for node := range *tree.GetGongstructInstancesSet[tree.Node](treeStagee) {
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
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree.Node{Name: _formgroup.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.OnUpdate = func(_ *tree.Stage, _, _ *tree.Node) {
					FillUpFormFromGongstruct(_formgroup, probe)
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree.Stage, stagedNode, frontNode *tree.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.FormGroup](probe)
				// set color for node and reset all other nodes color
				for node := range *tree.GetGongstructInstancesSet[tree.Node](treeStagee) {
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
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree.Node{Name: _formsortassocbutton.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.OnUpdate = func(_ *tree.Stage, _, _ *tree.Node) {
					FillUpFormFromGongstruct(_formsortassocbutton, probe)
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree.Stage, stagedNode, frontNode *tree.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.FormSortAssocButton](probe)
				// set color for node and reset all other nodes color
				for node := range *tree.GetGongstructInstancesSet[tree.Node](treeStagee) {
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
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree.Node{Name: _option.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.OnUpdate = func(_ *tree.Stage, _, _ *tree.Node) {
					FillUpFormFromGongstruct(_option, probe)
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree.Stage, stagedNode, frontNode *tree.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.Option](probe)
				// set color for node and reset all other nodes color
				for node := range *tree.GetGongstructInstancesSet[tree.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "Row":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Row](probe.stageOfInterest)
			count := 0
			for _row := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree.Node{Name: _row.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.OnUpdate = func(_ *tree.Stage, _, _ *tree.Node) {
					FillUpFormFromGongstruct(_row, probe)
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree.Stage, stagedNode, frontNode *tree.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.Row](probe)
				// set color for node and reset all other nodes color
				for node := range *tree.GetGongstructInstancesSet[tree.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "Table":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Table](probe.stageOfInterest)
			count := 0
			for _table := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree.Node{Name: _table.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.OnUpdate = func(_ *tree.Stage, _, _ *tree.Node) {
					FillUpFormFromGongstruct(_table, probe)
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree.Stage, stagedNode, frontNode *tree.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.Table](probe)
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
