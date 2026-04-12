package probe

import (
	"fmt"
	"sort"
	"strings"
	"time"

	tree_buttons "github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree_models "github.com/fullstack-lang/gong/lib/tree/go/models"

	gong_models "github.com/fullstack-lang/gong/go/models"

	"github.com/fullstack-lang/gong/dsm/reqif/go/models"
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

	notificationsResetButton := &tree_models.Button{
		Name:            "NotificationsResetButton",
		Icon:            string(tree_buttons.BUTTON_playlist_remove),
		HasToolTip:      true,
		ToolTipText:     "Reset notification table",
		ToolTipPosition: tree_models.Below,
		OnClick: func() {
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
		OnClick: func() {
			probe.stageOfInterest.ComputeInstancesNb()
			probe.docStager.SetMap_GongStructName_InstancesNb(
				probe.stageOfInterest.Map_GongStructName_InstancesNb,
			)
			probe.Refresh()
		},
	}
	topNode.Buttons = append(topNode.Buttons, refreshButton)

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
		case "ALTERNATIVE_ID":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.ALTERNATIVE_ID](probe.stageOfInterest)
			count := 0
			for _alternative_id := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _alternative_id.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_alternative_id, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.ALTERNATIVE_ID](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "ATTRIBUTE_DEFINITION_BOOLEAN":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.ATTRIBUTE_DEFINITION_BOOLEAN](probe.stageOfInterest)
			count := 0
			for _attribute_definition_boolean := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _attribute_definition_boolean.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_attribute_definition_boolean, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.ATTRIBUTE_DEFINITION_BOOLEAN](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "ATTRIBUTE_DEFINITION_BOOLEAN_Rendering":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.ATTRIBUTE_DEFINITION_BOOLEAN_Rendering](probe.stageOfInterest)
			count := 0
			for _attribute_definition_boolean_rendering := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _attribute_definition_boolean_rendering.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_attribute_definition_boolean_rendering, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.ATTRIBUTE_DEFINITION_BOOLEAN_Rendering](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "ATTRIBUTE_DEFINITION_DATE":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.ATTRIBUTE_DEFINITION_DATE](probe.stageOfInterest)
			count := 0
			for _attribute_definition_date := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _attribute_definition_date.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_attribute_definition_date, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.ATTRIBUTE_DEFINITION_DATE](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "ATTRIBUTE_DEFINITION_DATE_Rendering":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.ATTRIBUTE_DEFINITION_DATE_Rendering](probe.stageOfInterest)
			count := 0
			for _attribute_definition_date_rendering := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _attribute_definition_date_rendering.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_attribute_definition_date_rendering, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.ATTRIBUTE_DEFINITION_DATE_Rendering](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "ATTRIBUTE_DEFINITION_ENUMERATION":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.ATTRIBUTE_DEFINITION_ENUMERATION](probe.stageOfInterest)
			count := 0
			for _attribute_definition_enumeration := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _attribute_definition_enumeration.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_attribute_definition_enumeration, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.ATTRIBUTE_DEFINITION_ENUMERATION](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "ATTRIBUTE_DEFINITION_ENUMERATION_Rendering":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.ATTRIBUTE_DEFINITION_ENUMERATION_Rendering](probe.stageOfInterest)
			count := 0
			for _attribute_definition_enumeration_rendering := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _attribute_definition_enumeration_rendering.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_attribute_definition_enumeration_rendering, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.ATTRIBUTE_DEFINITION_ENUMERATION_Rendering](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "ATTRIBUTE_DEFINITION_INTEGER":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.ATTRIBUTE_DEFINITION_INTEGER](probe.stageOfInterest)
			count := 0
			for _attribute_definition_integer := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _attribute_definition_integer.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_attribute_definition_integer, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.ATTRIBUTE_DEFINITION_INTEGER](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "ATTRIBUTE_DEFINITION_INTEGER_Rendering":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.ATTRIBUTE_DEFINITION_INTEGER_Rendering](probe.stageOfInterest)
			count := 0
			for _attribute_definition_integer_rendering := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _attribute_definition_integer_rendering.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_attribute_definition_integer_rendering, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.ATTRIBUTE_DEFINITION_INTEGER_Rendering](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "ATTRIBUTE_DEFINITION_REAL":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.ATTRIBUTE_DEFINITION_REAL](probe.stageOfInterest)
			count := 0
			for _attribute_definition_real := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _attribute_definition_real.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_attribute_definition_real, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.ATTRIBUTE_DEFINITION_REAL](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "ATTRIBUTE_DEFINITION_REAL_Rendering":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.ATTRIBUTE_DEFINITION_REAL_Rendering](probe.stageOfInterest)
			count := 0
			for _attribute_definition_real_rendering := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _attribute_definition_real_rendering.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_attribute_definition_real_rendering, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.ATTRIBUTE_DEFINITION_REAL_Rendering](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "ATTRIBUTE_DEFINITION_Rendering":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.ATTRIBUTE_DEFINITION_Rendering](probe.stageOfInterest)
			count := 0
			for _attribute_definition_rendering := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _attribute_definition_rendering.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_attribute_definition_rendering, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.ATTRIBUTE_DEFINITION_Rendering](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "ATTRIBUTE_DEFINITION_STRING":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.ATTRIBUTE_DEFINITION_STRING](probe.stageOfInterest)
			count := 0
			for _attribute_definition_string := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _attribute_definition_string.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_attribute_definition_string, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.ATTRIBUTE_DEFINITION_STRING](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "ATTRIBUTE_DEFINITION_STRING_Rendering":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.ATTRIBUTE_DEFINITION_STRING_Rendering](probe.stageOfInterest)
			count := 0
			for _attribute_definition_string_rendering := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _attribute_definition_string_rendering.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_attribute_definition_string_rendering, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.ATTRIBUTE_DEFINITION_STRING_Rendering](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "ATTRIBUTE_DEFINITION_XHTML":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.ATTRIBUTE_DEFINITION_XHTML](probe.stageOfInterest)
			count := 0
			for _attribute_definition_xhtml := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _attribute_definition_xhtml.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_attribute_definition_xhtml, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.ATTRIBUTE_DEFINITION_XHTML](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "ATTRIBUTE_DEFINITION_XHTML_Rendering":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.ATTRIBUTE_DEFINITION_XHTML_Rendering](probe.stageOfInterest)
			count := 0
			for _attribute_definition_xhtml_rendering := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _attribute_definition_xhtml_rendering.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_attribute_definition_xhtml_rendering, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.ATTRIBUTE_DEFINITION_XHTML_Rendering](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "ATTRIBUTE_VALUE_BOOLEAN":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.ATTRIBUTE_VALUE_BOOLEAN](probe.stageOfInterest)
			count := 0
			for _attribute_value_boolean := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _attribute_value_boolean.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_attribute_value_boolean, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.ATTRIBUTE_VALUE_BOOLEAN](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "ATTRIBUTE_VALUE_DATE":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.ATTRIBUTE_VALUE_DATE](probe.stageOfInterest)
			count := 0
			for _attribute_value_date := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _attribute_value_date.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_attribute_value_date, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.ATTRIBUTE_VALUE_DATE](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "ATTRIBUTE_VALUE_ENUMERATION":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.ATTRIBUTE_VALUE_ENUMERATION](probe.stageOfInterest)
			count := 0
			for _attribute_value_enumeration := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _attribute_value_enumeration.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_attribute_value_enumeration, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.ATTRIBUTE_VALUE_ENUMERATION](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "ATTRIBUTE_VALUE_INTEGER":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.ATTRIBUTE_VALUE_INTEGER](probe.stageOfInterest)
			count := 0
			for _attribute_value_integer := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _attribute_value_integer.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_attribute_value_integer, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.ATTRIBUTE_VALUE_INTEGER](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "ATTRIBUTE_VALUE_REAL":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.ATTRIBUTE_VALUE_REAL](probe.stageOfInterest)
			count := 0
			for _attribute_value_real := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _attribute_value_real.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_attribute_value_real, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.ATTRIBUTE_VALUE_REAL](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "ATTRIBUTE_VALUE_STRING":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.ATTRIBUTE_VALUE_STRING](probe.stageOfInterest)
			count := 0
			for _attribute_value_string := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _attribute_value_string.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_attribute_value_string, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.ATTRIBUTE_VALUE_STRING](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "ATTRIBUTE_VALUE_XHTML":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.ATTRIBUTE_VALUE_XHTML](probe.stageOfInterest)
			count := 0
			for _attribute_value_xhtml := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _attribute_value_xhtml.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_attribute_value_xhtml, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.ATTRIBUTE_VALUE_XHTML](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "A_ALTERNATIVE_ID":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.A_ALTERNATIVE_ID](probe.stageOfInterest)
			count := 0
			for _a_alternative_id := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _a_alternative_id.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_a_alternative_id, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.A_ALTERNATIVE_ID](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "A_ATTRIBUTE_DEFINITION_BOOLEAN_REF":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.A_ATTRIBUTE_DEFINITION_BOOLEAN_REF](probe.stageOfInterest)
			count := 0
			for _a_attribute_definition_boolean_ref := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _a_attribute_definition_boolean_ref.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_a_attribute_definition_boolean_ref, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.A_ATTRIBUTE_DEFINITION_BOOLEAN_REF](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "A_ATTRIBUTE_DEFINITION_DATE_REF":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.A_ATTRIBUTE_DEFINITION_DATE_REF](probe.stageOfInterest)
			count := 0
			for _a_attribute_definition_date_ref := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _a_attribute_definition_date_ref.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_a_attribute_definition_date_ref, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.A_ATTRIBUTE_DEFINITION_DATE_REF](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "A_ATTRIBUTE_DEFINITION_ENUMERATION_REF":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.A_ATTRIBUTE_DEFINITION_ENUMERATION_REF](probe.stageOfInterest)
			count := 0
			for _a_attribute_definition_enumeration_ref := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _a_attribute_definition_enumeration_ref.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_a_attribute_definition_enumeration_ref, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.A_ATTRIBUTE_DEFINITION_ENUMERATION_REF](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "A_ATTRIBUTE_DEFINITION_INTEGER_REF":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.A_ATTRIBUTE_DEFINITION_INTEGER_REF](probe.stageOfInterest)
			count := 0
			for _a_attribute_definition_integer_ref := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _a_attribute_definition_integer_ref.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_a_attribute_definition_integer_ref, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.A_ATTRIBUTE_DEFINITION_INTEGER_REF](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "A_ATTRIBUTE_DEFINITION_REAL_REF":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.A_ATTRIBUTE_DEFINITION_REAL_REF](probe.stageOfInterest)
			count := 0
			for _a_attribute_definition_real_ref := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _a_attribute_definition_real_ref.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_a_attribute_definition_real_ref, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.A_ATTRIBUTE_DEFINITION_REAL_REF](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "A_ATTRIBUTE_DEFINITION_STRING_REF":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.A_ATTRIBUTE_DEFINITION_STRING_REF](probe.stageOfInterest)
			count := 0
			for _a_attribute_definition_string_ref := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _a_attribute_definition_string_ref.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_a_attribute_definition_string_ref, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.A_ATTRIBUTE_DEFINITION_STRING_REF](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "A_ATTRIBUTE_DEFINITION_XHTML_REF":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.A_ATTRIBUTE_DEFINITION_XHTML_REF](probe.stageOfInterest)
			count := 0
			for _a_attribute_definition_xhtml_ref := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _a_attribute_definition_xhtml_ref.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_a_attribute_definition_xhtml_ref, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.A_ATTRIBUTE_DEFINITION_XHTML_REF](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "A_ATTRIBUTE_VALUE_BOOLEAN":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.A_ATTRIBUTE_VALUE_BOOLEAN](probe.stageOfInterest)
			count := 0
			for _a_attribute_value_boolean := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _a_attribute_value_boolean.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_a_attribute_value_boolean, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.A_ATTRIBUTE_VALUE_BOOLEAN](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "A_ATTRIBUTE_VALUE_DATE":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.A_ATTRIBUTE_VALUE_DATE](probe.stageOfInterest)
			count := 0
			for _a_attribute_value_date := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _a_attribute_value_date.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_a_attribute_value_date, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.A_ATTRIBUTE_VALUE_DATE](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "A_ATTRIBUTE_VALUE_ENUMERATION":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.A_ATTRIBUTE_VALUE_ENUMERATION](probe.stageOfInterest)
			count := 0
			for _a_attribute_value_enumeration := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _a_attribute_value_enumeration.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_a_attribute_value_enumeration, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.A_ATTRIBUTE_VALUE_ENUMERATION](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "A_ATTRIBUTE_VALUE_INTEGER":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.A_ATTRIBUTE_VALUE_INTEGER](probe.stageOfInterest)
			count := 0
			for _a_attribute_value_integer := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _a_attribute_value_integer.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_a_attribute_value_integer, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.A_ATTRIBUTE_VALUE_INTEGER](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "A_ATTRIBUTE_VALUE_REAL":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.A_ATTRIBUTE_VALUE_REAL](probe.stageOfInterest)
			count := 0
			for _a_attribute_value_real := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _a_attribute_value_real.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_a_attribute_value_real, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.A_ATTRIBUTE_VALUE_REAL](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "A_ATTRIBUTE_VALUE_STRING":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.A_ATTRIBUTE_VALUE_STRING](probe.stageOfInterest)
			count := 0
			for _a_attribute_value_string := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _a_attribute_value_string.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_a_attribute_value_string, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.A_ATTRIBUTE_VALUE_STRING](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "A_ATTRIBUTE_VALUE_XHTML":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.A_ATTRIBUTE_VALUE_XHTML](probe.stageOfInterest)
			count := 0
			for _a_attribute_value_xhtml := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _a_attribute_value_xhtml.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_a_attribute_value_xhtml, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.A_ATTRIBUTE_VALUE_XHTML](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "A_ATTRIBUTE_VALUE_XHTML_1":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.A_ATTRIBUTE_VALUE_XHTML_1](probe.stageOfInterest)
			count := 0
			for _a_attribute_value_xhtml_1 := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _a_attribute_value_xhtml_1.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_a_attribute_value_xhtml_1, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.A_ATTRIBUTE_VALUE_XHTML_1](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "A_CHILDREN":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.A_CHILDREN](probe.stageOfInterest)
			count := 0
			for _a_children := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _a_children.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_a_children, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.A_CHILDREN](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "A_CORE_CONTENT":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.A_CORE_CONTENT](probe.stageOfInterest)
			count := 0
			for _a_core_content := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _a_core_content.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_a_core_content, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.A_CORE_CONTENT](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "A_DATATYPES":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.A_DATATYPES](probe.stageOfInterest)
			count := 0
			for _a_datatypes := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _a_datatypes.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_a_datatypes, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.A_DATATYPES](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "A_DATATYPE_DEFINITION_BOOLEAN_REF":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.A_DATATYPE_DEFINITION_BOOLEAN_REF](probe.stageOfInterest)
			count := 0
			for _a_datatype_definition_boolean_ref := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _a_datatype_definition_boolean_ref.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_a_datatype_definition_boolean_ref, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.A_DATATYPE_DEFINITION_BOOLEAN_REF](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "A_DATATYPE_DEFINITION_DATE_REF":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.A_DATATYPE_DEFINITION_DATE_REF](probe.stageOfInterest)
			count := 0
			for _a_datatype_definition_date_ref := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _a_datatype_definition_date_ref.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_a_datatype_definition_date_ref, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.A_DATATYPE_DEFINITION_DATE_REF](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "A_DATATYPE_DEFINITION_ENUMERATION_REF":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.A_DATATYPE_DEFINITION_ENUMERATION_REF](probe.stageOfInterest)
			count := 0
			for _a_datatype_definition_enumeration_ref := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _a_datatype_definition_enumeration_ref.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_a_datatype_definition_enumeration_ref, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.A_DATATYPE_DEFINITION_ENUMERATION_REF](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "A_DATATYPE_DEFINITION_INTEGER_REF":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.A_DATATYPE_DEFINITION_INTEGER_REF](probe.stageOfInterest)
			count := 0
			for _a_datatype_definition_integer_ref := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _a_datatype_definition_integer_ref.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_a_datatype_definition_integer_ref, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.A_DATATYPE_DEFINITION_INTEGER_REF](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "A_DATATYPE_DEFINITION_REAL_REF":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.A_DATATYPE_DEFINITION_REAL_REF](probe.stageOfInterest)
			count := 0
			for _a_datatype_definition_real_ref := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _a_datatype_definition_real_ref.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_a_datatype_definition_real_ref, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.A_DATATYPE_DEFINITION_REAL_REF](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "A_DATATYPE_DEFINITION_STRING_REF":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.A_DATATYPE_DEFINITION_STRING_REF](probe.stageOfInterest)
			count := 0
			for _a_datatype_definition_string_ref := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _a_datatype_definition_string_ref.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_a_datatype_definition_string_ref, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.A_DATATYPE_DEFINITION_STRING_REF](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "A_DATATYPE_DEFINITION_XHTML_REF":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.A_DATATYPE_DEFINITION_XHTML_REF](probe.stageOfInterest)
			count := 0
			for _a_datatype_definition_xhtml_ref := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _a_datatype_definition_xhtml_ref.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_a_datatype_definition_xhtml_ref, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.A_DATATYPE_DEFINITION_XHTML_REF](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "A_EDITABLE_ATTS":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.A_EDITABLE_ATTS](probe.stageOfInterest)
			count := 0
			for _a_editable_atts := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _a_editable_atts.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_a_editable_atts, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.A_EDITABLE_ATTS](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "A_ENUM_VALUE_REF":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.A_ENUM_VALUE_REF](probe.stageOfInterest)
			count := 0
			for _a_enum_value_ref := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _a_enum_value_ref.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_a_enum_value_ref, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.A_ENUM_VALUE_REF](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "A_OBJECT":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.A_OBJECT](probe.stageOfInterest)
			count := 0
			for _a_object := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _a_object.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_a_object, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.A_OBJECT](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "A_PROPERTIES":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.A_PROPERTIES](probe.stageOfInterest)
			count := 0
			for _a_properties := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _a_properties.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_a_properties, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.A_PROPERTIES](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "A_RELATION_GROUP_TYPE_REF":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.A_RELATION_GROUP_TYPE_REF](probe.stageOfInterest)
			count := 0
			for _a_relation_group_type_ref := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _a_relation_group_type_ref.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_a_relation_group_type_ref, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.A_RELATION_GROUP_TYPE_REF](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "A_SOURCE_1":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.A_SOURCE_1](probe.stageOfInterest)
			count := 0
			for _a_source_1 := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _a_source_1.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_a_source_1, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.A_SOURCE_1](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "A_SOURCE_SPECIFICATION_1":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.A_SOURCE_SPECIFICATION_1](probe.stageOfInterest)
			count := 0
			for _a_source_specification_1 := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _a_source_specification_1.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_a_source_specification_1, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.A_SOURCE_SPECIFICATION_1](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "A_SPECIFICATIONS":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.A_SPECIFICATIONS](probe.stageOfInterest)
			count := 0
			for _a_specifications := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _a_specifications.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_a_specifications, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.A_SPECIFICATIONS](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "A_SPECIFICATION_TYPE_REF":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.A_SPECIFICATION_TYPE_REF](probe.stageOfInterest)
			count := 0
			for _a_specification_type_ref := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _a_specification_type_ref.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_a_specification_type_ref, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.A_SPECIFICATION_TYPE_REF](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "A_SPECIFIED_VALUES":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.A_SPECIFIED_VALUES](probe.stageOfInterest)
			count := 0
			for _a_specified_values := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _a_specified_values.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_a_specified_values, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.A_SPECIFIED_VALUES](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "A_SPEC_ATTRIBUTES":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.A_SPEC_ATTRIBUTES](probe.stageOfInterest)
			count := 0
			for _a_spec_attributes := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _a_spec_attributes.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_a_spec_attributes, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.A_SPEC_ATTRIBUTES](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "A_SPEC_OBJECTS":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.A_SPEC_OBJECTS](probe.stageOfInterest)
			count := 0
			for _a_spec_objects := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _a_spec_objects.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_a_spec_objects, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.A_SPEC_OBJECTS](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "A_SPEC_OBJECT_TYPE_REF":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.A_SPEC_OBJECT_TYPE_REF](probe.stageOfInterest)
			count := 0
			for _a_spec_object_type_ref := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _a_spec_object_type_ref.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_a_spec_object_type_ref, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.A_SPEC_OBJECT_TYPE_REF](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "A_SPEC_RELATIONS":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.A_SPEC_RELATIONS](probe.stageOfInterest)
			count := 0
			for _a_spec_relations := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _a_spec_relations.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_a_spec_relations, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.A_SPEC_RELATIONS](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "A_SPEC_RELATION_GROUPS":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.A_SPEC_RELATION_GROUPS](probe.stageOfInterest)
			count := 0
			for _a_spec_relation_groups := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _a_spec_relation_groups.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_a_spec_relation_groups, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.A_SPEC_RELATION_GROUPS](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "A_SPEC_RELATION_REF":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.A_SPEC_RELATION_REF](probe.stageOfInterest)
			count := 0
			for _a_spec_relation_ref := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _a_spec_relation_ref.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_a_spec_relation_ref, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.A_SPEC_RELATION_REF](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "A_SPEC_RELATION_TYPE_REF":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.A_SPEC_RELATION_TYPE_REF](probe.stageOfInterest)
			count := 0
			for _a_spec_relation_type_ref := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _a_spec_relation_type_ref.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_a_spec_relation_type_ref, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.A_SPEC_RELATION_TYPE_REF](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "A_SPEC_TYPES":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.A_SPEC_TYPES](probe.stageOfInterest)
			count := 0
			for _a_spec_types := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _a_spec_types.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_a_spec_types, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.A_SPEC_TYPES](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "A_THE_HEADER":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.A_THE_HEADER](probe.stageOfInterest)
			count := 0
			for _a_the_header := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _a_the_header.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_a_the_header, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.A_THE_HEADER](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "A_TOOL_EXTENSIONS":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.A_TOOL_EXTENSIONS](probe.stageOfInterest)
			count := 0
			for _a_tool_extensions := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _a_tool_extensions.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_a_tool_extensions, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.A_TOOL_EXTENSIONS](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "DATATYPE_DEFINITION_BOOLEAN":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.DATATYPE_DEFINITION_BOOLEAN](probe.stageOfInterest)
			count := 0
			for _datatype_definition_boolean := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _datatype_definition_boolean.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_datatype_definition_boolean, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.DATATYPE_DEFINITION_BOOLEAN](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "DATATYPE_DEFINITION_DATE":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.DATATYPE_DEFINITION_DATE](probe.stageOfInterest)
			count := 0
			for _datatype_definition_date := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _datatype_definition_date.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_datatype_definition_date, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.DATATYPE_DEFINITION_DATE](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "DATATYPE_DEFINITION_ENUMERATION":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.DATATYPE_DEFINITION_ENUMERATION](probe.stageOfInterest)
			count := 0
			for _datatype_definition_enumeration := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _datatype_definition_enumeration.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_datatype_definition_enumeration, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.DATATYPE_DEFINITION_ENUMERATION](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "DATATYPE_DEFINITION_INTEGER":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.DATATYPE_DEFINITION_INTEGER](probe.stageOfInterest)
			count := 0
			for _datatype_definition_integer := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _datatype_definition_integer.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_datatype_definition_integer, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.DATATYPE_DEFINITION_INTEGER](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "DATATYPE_DEFINITION_REAL":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.DATATYPE_DEFINITION_REAL](probe.stageOfInterest)
			count := 0
			for _datatype_definition_real := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _datatype_definition_real.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_datatype_definition_real, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.DATATYPE_DEFINITION_REAL](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "DATATYPE_DEFINITION_STRING":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.DATATYPE_DEFINITION_STRING](probe.stageOfInterest)
			count := 0
			for _datatype_definition_string := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _datatype_definition_string.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_datatype_definition_string, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.DATATYPE_DEFINITION_STRING](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "DATATYPE_DEFINITION_XHTML":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.DATATYPE_DEFINITION_XHTML](probe.stageOfInterest)
			count := 0
			for _datatype_definition_xhtml := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _datatype_definition_xhtml.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_datatype_definition_xhtml, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.DATATYPE_DEFINITION_XHTML](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "EMBEDDED_VALUE":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.EMBEDDED_VALUE](probe.stageOfInterest)
			count := 0
			for _embedded_value := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _embedded_value.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_embedded_value, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.EMBEDDED_VALUE](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "ENUM_VALUE":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.ENUM_VALUE](probe.stageOfInterest)
			count := 0
			for _enum_value := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _enum_value.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_enum_value, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.ENUM_VALUE](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "EmbeddedJpgImage":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.EmbeddedJpgImage](probe.stageOfInterest)
			count := 0
			for _embeddedjpgimage := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _embeddedjpgimage.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_embeddedjpgimage, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.EmbeddedJpgImage](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "EmbeddedPngImage":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.EmbeddedPngImage](probe.stageOfInterest)
			count := 0
			for _embeddedpngimage := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _embeddedpngimage.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_embeddedpngimage, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.EmbeddedPngImage](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "EmbeddedSvgImage":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.EmbeddedSvgImage](probe.stageOfInterest)
			count := 0
			for _embeddedsvgimage := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _embeddedsvgimage.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_embeddedsvgimage, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.EmbeddedSvgImage](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "Kill":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Kill](probe.stageOfInterest)
			count := 0
			for _kill := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _kill.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_kill, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.Kill](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "Map_identifier_bool":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Map_identifier_bool](probe.stageOfInterest)
			count := 0
			for _map_identifier_bool := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _map_identifier_bool.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_map_identifier_bool, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.Map_identifier_bool](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "RELATION_GROUP":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.RELATION_GROUP](probe.stageOfInterest)
			count := 0
			for _relation_group := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _relation_group.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_relation_group, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.RELATION_GROUP](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "RELATION_GROUP_TYPE":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.RELATION_GROUP_TYPE](probe.stageOfInterest)
			count := 0
			for _relation_group_type := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _relation_group_type.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_relation_group_type, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.RELATION_GROUP_TYPE](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "REQ_IF":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.REQ_IF](probe.stageOfInterest)
			count := 0
			for _req_if := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _req_if.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_req_if, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.REQ_IF](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "REQ_IF_CONTENT":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.REQ_IF_CONTENT](probe.stageOfInterest)
			count := 0
			for _req_if_content := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _req_if_content.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_req_if_content, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.REQ_IF_CONTENT](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "REQ_IF_HEADER":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.REQ_IF_HEADER](probe.stageOfInterest)
			count := 0
			for _req_if_header := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _req_if_header.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_req_if_header, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.REQ_IF_HEADER](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "REQ_IF_TOOL_EXTENSION":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.REQ_IF_TOOL_EXTENSION](probe.stageOfInterest)
			count := 0
			for _req_if_tool_extension := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _req_if_tool_extension.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_req_if_tool_extension, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.REQ_IF_TOOL_EXTENSION](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "SPECIFICATION":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.SPECIFICATION](probe.stageOfInterest)
			count := 0
			for _specification := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _specification.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_specification, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.SPECIFICATION](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "SPECIFICATION_Rendering":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.SPECIFICATION_Rendering](probe.stageOfInterest)
			count := 0
			for _specification_rendering := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _specification_rendering.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_specification_rendering, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.SPECIFICATION_Rendering](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "SPECIFICATION_TYPE":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.SPECIFICATION_TYPE](probe.stageOfInterest)
			count := 0
			for _specification_type := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _specification_type.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_specification_type, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.SPECIFICATION_TYPE](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "SPEC_HIERARCHY":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.SPEC_HIERARCHY](probe.stageOfInterest)
			count := 0
			for _spec_hierarchy := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _spec_hierarchy.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_spec_hierarchy, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.SPEC_HIERARCHY](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "SPEC_OBJECT":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.SPEC_OBJECT](probe.stageOfInterest)
			count := 0
			for _spec_object := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _spec_object.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_spec_object, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.SPEC_OBJECT](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "SPEC_OBJECT_TYPE":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.SPEC_OBJECT_TYPE](probe.stageOfInterest)
			count := 0
			for _spec_object_type := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _spec_object_type.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_spec_object_type, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.SPEC_OBJECT_TYPE](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "SPEC_OBJECT_TYPE_Rendering":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.SPEC_OBJECT_TYPE_Rendering](probe.stageOfInterest)
			count := 0
			for _spec_object_type_rendering := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _spec_object_type_rendering.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_spec_object_type_rendering, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.SPEC_OBJECT_TYPE_Rendering](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "SPEC_RELATION":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.SPEC_RELATION](probe.stageOfInterest)
			count := 0
			for _spec_relation := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _spec_relation.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_spec_relation, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.SPEC_RELATION](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "SPEC_RELATION_TYPE":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.SPEC_RELATION_TYPE](probe.stageOfInterest)
			count := 0
			for _spec_relation_type := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _spec_relation_type.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_spec_relation_type, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.SPEC_RELATION_TYPE](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "StaticWebSite":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.StaticWebSite](probe.stageOfInterest)
			count := 0
			for _staticwebsite := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _staticwebsite.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_staticwebsite, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.StaticWebSite](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "StaticWebSiteChapter":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.StaticWebSiteChapter](probe.stageOfInterest)
			count := 0
			for _staticwebsitechapter := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _staticwebsitechapter.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_staticwebsitechapter, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.StaticWebSiteChapter](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "StaticWebSiteGeneratedImage":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.StaticWebSiteGeneratedImage](probe.stageOfInterest)
			count := 0
			for _staticwebsitegeneratedimage := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _staticwebsitegeneratedimage.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_staticwebsitegeneratedimage, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.StaticWebSiteGeneratedImage](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "StaticWebSiteImage":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.StaticWebSiteImage](probe.stageOfInterest)
			count := 0
			for _staticwebsiteimage := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _staticwebsiteimage.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_staticwebsiteimage, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.StaticWebSiteImage](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "StaticWebSiteParagraph":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.StaticWebSiteParagraph](probe.stageOfInterest)
			count := 0
			for _staticwebsiteparagraph := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _staticwebsiteparagraph.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_staticwebsiteparagraph, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.StaticWebSiteParagraph](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "XHTML_CONTENT":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.XHTML_CONTENT](probe.stageOfInterest)
			count := 0
			for _xhtml_content := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _xhtml_content.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_xhtml_content, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.XHTML_CONTENT](probe)
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
