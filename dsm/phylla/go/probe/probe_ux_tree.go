package probe

import (
	"fmt"
	"sort"
	"strings"
	"time"

	tree_buttons "github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree_models "github.com/fullstack-lang/gong/lib/tree/go/models"

	gong_models "github.com/fullstack-lang/gong/go/models"

	"github.com/fullstack-lang/gong/dsm/phylla/go/models"
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
		ToolTipText:     "Export XL",
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
		ToolTipText:     "Export Go",
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
		case "Angle0Shape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Angle0Shape](probe.stageOfInterest)
			count := 0
			for _angle0shape := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _angle0shape.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_angle0shape, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.Angle0Shape](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "ArcNormalVectorShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.ArcNormalVectorShape](probe.stageOfInterest)
			count := 0
			for _arcnormalvectorshape := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _arcnormalvectorshape.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_arcnormalvectorshape, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.ArcNormalVectorShape](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "ArcNormalVectorShapeGrid":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.ArcNormalVectorShapeGrid](probe.stageOfInterest)
			count := 0
			for _arcnormalvectorshapegrid := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _arcnormalvectorshapegrid.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_arcnormalvectorshapegrid, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.ArcNormalVectorShapeGrid](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "AxesShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.AxesShape](probe.stageOfInterest)
			count := 0
			for _axesshape := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _axesshape.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_axesshape, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.AxesShape](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "BaseVectorShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.BaseVectorShape](probe.stageOfInterest)
			count := 0
			for _basevectorshape := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _basevectorshape.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_basevectorshape, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.BaseVectorShape](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "BaseVectorShapeGrid":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.BaseVectorShapeGrid](probe.stageOfInterest)
			count := 0
			for _basevectorshapegrid := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _basevectorshapegrid.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_basevectorshapegrid, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.BaseVectorShapeGrid](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "ChosenP1P2PairShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.ChosenP1P2PairShape](probe.stageOfInterest)
			count := 0
			for _chosenp1p2pairshape := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _chosenp1p2pairshape.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_chosenp1p2pairshape, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.ChosenP1P2PairShape](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "CircleGridShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.CircleGridShape](probe.stageOfInterest)
			count := 0
			for _circlegridshape := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _circlegridshape.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_circlegridshape, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.CircleGridShape](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "ClockAbstract":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.ClockAbstract](probe.stageOfInterest)
			count := 0
			for _clockabstract := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _clockabstract.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_clockabstract, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.ClockAbstract](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "ClockDiagram":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.ClockDiagram](probe.stageOfInterest)
			count := 0
			for _clockdiagram := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _clockdiagram.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_clockdiagram, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.ClockDiagram](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "ClockTopCurveShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.ClockTopCurveShape](probe.stageOfInterest)
			count := 0
			for _clocktopcurveshape := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _clocktopcurveshape.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_clocktopcurveshape, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.ClockTopCurveShape](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "EndArcShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.EndArcShape](probe.stageOfInterest)
			count := 0
			for _endarcshape := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _endarcshape.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_endarcshape, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.EndArcShape](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "EndArcShapeGrid":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.EndArcShapeGrid](probe.stageOfInterest)
			count := 0
			for _endarcshapegrid := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _endarcshapegrid.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_endarcshapegrid, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.EndArcShapeGrid](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "EndHalfwayArcShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.EndHalfwayArcShape](probe.stageOfInterest)
			count := 0
			for _endhalfwayarcshape := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _endhalfwayarcshape.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_endhalfwayarcshape, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.EndHalfwayArcShape](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "EndHalfwayArcShapeGrid":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.EndHalfwayArcShapeGrid](probe.stageOfInterest)
			count := 0
			for _endhalfwayarcshapegrid := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _endhalfwayarcshapegrid.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_endhalfwayarcshapegrid, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.EndHalfwayArcShapeGrid](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "ExplanationTextShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.ExplanationTextShape](probe.stageOfInterest)
			count := 0
			for _explanationtextshape := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _explanationtextshape.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_explanationtextshape, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.ExplanationTextShape](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "Eye3DShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Eye3DShape](probe.stageOfInterest)
			count := 0
			for _eye3dshape := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _eye3dshape.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_eye3dshape, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.Eye3DShape](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "EyeCornersSampledPoints3DShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.EyeCornersSampledPoints3DShape](probe.stageOfInterest)
			count := 0
			for _eyecornerssampledpoints3dshape := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _eyecornerssampledpoints3dshape.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_eyecornerssampledpoints3dshape, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.EyeCornersSampledPoints3DShape](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "EyeSampledPoints3DShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.EyeSampledPoints3DShape](probe.stageOfInterest)
			count := 0
			for _eyesampledpoints3dshape := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _eyesampledpoints3dshape.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_eyesampledpoints3dshape, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.EyeSampledPoints3DShape](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "EyeSeatBottomCurveShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.EyeSeatBottomCurveShape](probe.stageOfInterest)
			count := 0
			for _eyeseatbottomcurveshape := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _eyeseatbottomcurveshape.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_eyeseatbottomcurveshape, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.EyeSeatBottomCurveShape](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "EyeStoolBottomCurveShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.EyeStoolBottomCurveShape](probe.stageOfInterest)
			count := 0
			for _eyestoolbottomcurveshape := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _eyestoolbottomcurveshape.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_eyestoolbottomcurveshape, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.EyeStoolBottomCurveShape](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "EyeVolume3DShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.EyeVolume3DShape](probe.stageOfInterest)
			count := 0
			for _eyevolume3dshape := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _eyevolume3dshape.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_eyevolume3dshape, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.EyeVolume3DShape](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "GridPathShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.GridPathShape](probe.stageOfInterest)
			count := 0
			for _gridpathshape := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _gridpathshape.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_gridpathshape, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.GridPathShape](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "GrowthCurve2D":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.GrowthCurve2D](probe.stageOfInterest)
			count := 0
			for _growthcurve2d := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _growthcurve2d.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_growthcurve2d, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.GrowthCurve2D](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "GrowthCurve2DRibbon":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.GrowthCurve2DRibbon](probe.stageOfInterest)
			count := 0
			for _growthcurve2dribbon := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _growthcurve2dribbon.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_growthcurve2dribbon, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.GrowthCurve2DRibbon](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "GrowthCurve2DRibbonEndShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.GrowthCurve2DRibbonEndShape](probe.stageOfInterest)
			count := 0
			for _growthcurve2dribbonendshape := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _growthcurve2dribbonendshape.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_growthcurve2dribbonendshape, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.GrowthCurve2DRibbonEndShape](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "GrowthCurve2DRibbonStartShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.GrowthCurve2DRibbonStartShape](probe.stageOfInterest)
			count := 0
			for _growthcurve2dribbonstartshape := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _growthcurve2dribbonstartshape.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_growthcurve2dribbonstartshape, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.GrowthCurve2DRibbonStartShape](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "GrowthCurveRhombusGridShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.GrowthCurveRhombusGridShape](probe.stageOfInterest)
			count := 0
			for _growthcurverhombusgridshape := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _growthcurverhombusgridshape.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_growthcurverhombusgridshape, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.GrowthCurveRhombusGridShape](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "GrowthCurveRhombusShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.GrowthCurveRhombusShape](probe.stageOfInterest)
			count := 0
			for _growthcurverhombusshape := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _growthcurverhombusshape.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_growthcurverhombusshape, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.GrowthCurveRhombusShape](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "GrowthVectorShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.GrowthVectorShape](probe.stageOfInterest)
			count := 0
			for _growthvectorshape := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _growthvectorshape.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_growthvectorshape, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.GrowthVectorShape](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "InitialRhombusGridShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.InitialRhombusGridShape](probe.stageOfInterest)
			count := 0
			for _initialrhombusgridshape := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _initialrhombusgridshape.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_initialrhombusgridshape, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.InitialRhombusGridShape](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "InitialRhombusShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.InitialRhombusShape](probe.stageOfInterest)
			count := 0
			for _initialrhombusshape := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _initialrhombusshape.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_initialrhombusshape, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.InitialRhombusShape](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "Key3DShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Key3DShape](probe.stageOfInterest)
			count := 0
			for _key3dshape := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _key3dshape.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_key3dshape, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.Key3DShape](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "KeyHole3DShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.KeyHole3DShape](probe.stageOfInterest)
			count := 0
			for _keyhole3dshape := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _keyhole3dshape.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_keyhole3dshape, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.KeyHole3DShape](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "KeyHoleShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.KeyHoleShape](probe.stageOfInterest)
			count := 0
			for _keyholeshape := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _keyholeshape.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_keyholeshape, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.KeyHoleShape](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "Library":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Library](probe.stageOfInterest)
			count := 0
			for _library := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _library.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_library, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.Library](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "MidArcVectorShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.MidArcVectorShape](probe.stageOfInterest)
			count := 0
			for _midarcvectorshape := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _midarcvectorshape.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_midarcvectorshape, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.MidArcVectorShape](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "MidArcVectorShapeGrid":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.MidArcVectorShapeGrid](probe.stageOfInterest)
			count := 0
			for _midarcvectorshapegrid := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _midarcvectorshapegrid.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_midarcvectorshapegrid, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.MidArcVectorShapeGrid](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "OriginalPoints3DShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.OriginalPoints3DShape](probe.stageOfInterest)
			count := 0
			for _originalpoints3dshape := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _originalpoints3dshape.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_originalpoints3dshape, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.OriginalPoints3DShape](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "PartiallyGrowthCurve2DRibbon":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.PartiallyGrowthCurve2DRibbon](probe.stageOfInterest)
			count := 0
			for _partiallygrowthcurve2dribbon := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _partiallygrowthcurve2dribbon.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_partiallygrowthcurve2dribbon, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.PartiallyGrowthCurve2DRibbon](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "PartiallyGrowthCurve2DRibbonEndShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.PartiallyGrowthCurve2DRibbonEndShape](probe.stageOfInterest)
			count := 0
			for _partiallygrowthcurve2dribbonendshape := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _partiallygrowthcurve2dribbonendshape.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_partiallygrowthcurve2dribbonendshape, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.PartiallyGrowthCurve2DRibbonEndShape](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "PartiallyGrowthCurve2DRibbonStartShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.PartiallyGrowthCurve2DRibbonStartShape](probe.stageOfInterest)
			count := 0
			for _partiallygrowthcurve2dribbonstartshape := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _partiallygrowthcurve2dribbonstartshape.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_partiallygrowthcurve2dribbonstartshape, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.PartiallyGrowthCurve2DRibbonStartShape](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "PartiallyGrowthCurve2DTrajectory":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.PartiallyGrowthCurve2DTrajectory](probe.stageOfInterest)
			count := 0
			for _partiallygrowthcurve2dtrajectory := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _partiallygrowthcurve2dtrajectory.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_partiallygrowthcurve2dtrajectory, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.PartiallyGrowthCurve2DTrajectory](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "PartiallyGrowthCurve2DTrajectoryP1CurveShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.PartiallyGrowthCurve2DTrajectoryP1CurveShape](probe.stageOfInterest)
			count := 0
			for _partiallygrowthcurve2dtrajectoryp1curveshape := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _partiallygrowthcurve2dtrajectoryp1curveshape.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_partiallygrowthcurve2dtrajectoryp1curveshape, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.PartiallyGrowthCurve2DTrajectoryP1CurveShape](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "PartiallyGrowthCurve2DTrajectoryP1P2":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.PartiallyGrowthCurve2DTrajectoryP1P2](probe.stageOfInterest)
			count := 0
			for _partiallygrowthcurve2dtrajectoryp1p2 := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _partiallygrowthcurve2dtrajectoryp1p2.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_partiallygrowthcurve2dtrajectoryp1p2, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.PartiallyGrowthCurve2DTrajectoryP1P2](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape](probe.stageOfInterest)
			count := 0
			for _partiallygrowthcurve2dtrajectoryp1p2pairlineshape := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _partiallygrowthcurve2dtrajectoryp1p2pairlineshape.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_partiallygrowthcurve2dtrajectoryp1p2pairlineshape, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "PartiallyGrowthCurve2DTrajectoryP1PointShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.PartiallyGrowthCurve2DTrajectoryP1PointShape](probe.stageOfInterest)
			count := 0
			for _partiallygrowthcurve2dtrajectoryp1pointshape := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _partiallygrowthcurve2dtrajectoryp1pointshape.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_partiallygrowthcurve2dtrajectoryp1pointshape, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.PartiallyGrowthCurve2DTrajectoryP1PointShape](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "PartiallyGrowthCurve2DTrajectoryP2CurveShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.PartiallyGrowthCurve2DTrajectoryP2CurveShape](probe.stageOfInterest)
			count := 0
			for _partiallygrowthcurve2dtrajectoryp2curveshape := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _partiallygrowthcurve2dtrajectoryp2curveshape.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_partiallygrowthcurve2dtrajectoryp2curveshape, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.PartiallyGrowthCurve2DTrajectoryP2CurveShape](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "PartiallyGrowthCurve2DTrajectoryP2PointShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.PartiallyGrowthCurve2DTrajectoryP2PointShape](probe.stageOfInterest)
			count := 0
			for _partiallygrowthcurve2dtrajectoryp2pointshape := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _partiallygrowthcurve2dtrajectoryp2pointshape.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_partiallygrowthcurve2dtrajectoryp2pointshape, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.PartiallyGrowthCurve2DTrajectoryP2PointShape](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "PartiallyGrowthCurve2DTrajectoryShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.PartiallyGrowthCurve2DTrajectoryShape](probe.stageOfInterest)
			count := 0
			for _partiallygrowthcurve2dtrajectoryshape := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _partiallygrowthcurve2dtrajectoryshape.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_partiallygrowthcurve2dtrajectoryshape, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.PartiallyGrowthCurve2DTrajectoryShape](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "PartiallyRotatedSeatBottomCurveShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.PartiallyRotatedSeatBottomCurveShape](probe.stageOfInterest)
			count := 0
			for _partiallyrotatedseatbottomcurveshape := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _partiallyrotatedseatbottomcurveshape.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_partiallyrotatedseatbottomcurveshape, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.PartiallyRotatedSeatBottomCurveShape](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "PartiallyRotatedSeatTopCurveShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.PartiallyRotatedSeatTopCurveShape](probe.stageOfInterest)
			count := 0
			for _partiallyrotatedseattopcurveshape := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _partiallyrotatedseattopcurveshape.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_partiallyrotatedseattopcurveshape, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.PartiallyRotatedSeatTopCurveShape](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "PartiallyRotatedTorusShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.PartiallyRotatedTorusShape](probe.stageOfInterest)
			count := 0
			for _partiallyrotatedtorusshape := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _partiallyrotatedtorusshape.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_partiallyrotatedtorusshape, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.PartiallyRotatedTorusShape](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "PerpendicularVector":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.PerpendicularVector](probe.stageOfInterest)
			count := 0
			for _perpendicularvector := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _perpendicularvector.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_perpendicularvector, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.PerpendicularVector](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "PerpendicularVectorGrid":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.PerpendicularVectorGrid](probe.stageOfInterest)
			count := 0
			for _perpendicularvectorgrid := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _perpendicularvectorgrid.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_perpendicularvectorgrid, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.PerpendicularVectorGrid](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "PerpendicularVectorGridHalfway":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.PerpendicularVectorGridHalfway](probe.stageOfInterest)
			count := 0
			for _perpendicularvectorgridhalfway := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _perpendicularvectorgridhalfway.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_perpendicularvectorgridhalfway, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.PerpendicularVectorGridHalfway](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "PerpendicularVectorHalfway":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.PerpendicularVectorHalfway](probe.stageOfInterest)
			count := 0
			for _perpendicularvectorhalfway := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _perpendicularvectorhalfway.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_perpendicularvectorhalfway, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.PerpendicularVectorHalfway](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "PlantAbstract":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.PlantAbstract](probe.stageOfInterest)
			count := 0
			for _plantabstract := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _plantabstract.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_plantabstract, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.PlantAbstract](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "PlantCircumferenceShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.PlantCircumferenceShape](probe.stageOfInterest)
			count := 0
			for _plantcircumferenceshape := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _plantcircumferenceshape.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_plantcircumferenceshape, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.PlantCircumferenceShape](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "PlantDiagram":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.PlantDiagram](probe.stageOfInterest)
			count := 0
			for _plantdiagram := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _plantdiagram.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_plantdiagram, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.PlantDiagram](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "PointsAndLines3DShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.PointsAndLines3DShape](probe.stageOfInterest)
			count := 0
			for _pointsandlines3dshape := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _pointsandlines3dshape.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_pointsandlines3dshape, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.PointsAndLines3DShape](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "PxShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.PxShape](probe.stageOfInterest)
			count := 0
			for _pxshape := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _pxshape.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_pxshape, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.PxShape](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "Rendered3DShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Rendered3DShape](probe.stageOfInterest)
			count := 0
			for _rendered3dshape := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _rendered3dshape.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_rendered3dshape, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.Rendered3DShape](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "RhombusShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.RhombusShape](probe.stageOfInterest)
			count := 0
			for _rhombusshape := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _rhombusshape.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_rhombusshape, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.RhombusShape](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "RhombusStuff":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.RhombusStuff](probe.stageOfInterest)
			count := 0
			for _rhombusstuff := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _rhombusstuff.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_rhombusstuff, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.RhombusStuff](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "RotatedRhombusGridShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.RotatedRhombusGridShape](probe.stageOfInterest)
			count := 0
			for _rotatedrhombusgridshape := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _rotatedrhombusgridshape.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_rotatedrhombusgridshape, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.RotatedRhombusGridShape](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "RotatedRhombusShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.RotatedRhombusShape](probe.stageOfInterest)
			count := 0
			for _rotatedrhombusshape := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _rotatedrhombusshape.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_rotatedrhombusshape, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.RotatedRhombusShape](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "RotatedSampledPoints3DShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.RotatedSampledPoints3DShape](probe.stageOfInterest)
			count := 0
			for _rotatedsampledpoints3dshape := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _rotatedsampledpoints3dshape.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_rotatedsampledpoints3dshape, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.RotatedSampledPoints3DShape](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "RotatedSeatAndLegs3DShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.RotatedSeatAndLegs3DShape](probe.stageOfInterest)
			count := 0
			for _rotatedseatandlegs3dshape := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _rotatedseatandlegs3dshape.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_rotatedseatandlegs3dshape, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.RotatedSeatAndLegs3DShape](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "SampledPoints3DShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.SampledPoints3DShape](probe.stageOfInterest)
			count := 0
			for _sampledpoints3dshape := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _sampledpoints3dshape.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_sampledpoints3dshape, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.SampledPoints3DShape](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "Seat3DShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Seat3DShape](probe.stageOfInterest)
			count := 0
			for _seat3dshape := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _seat3dshape.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_seat3dshape, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.Seat3DShape](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "SeatAndLegs3DShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.SeatAndLegs3DShape](probe.stageOfInterest)
			count := 0
			for _seatandlegs3dshape := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _seatandlegs3dshape.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_seatandlegs3dshape, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.SeatAndLegs3DShape](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "SeatBottomCurveShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.SeatBottomCurveShape](probe.stageOfInterest)
			count := 0
			for _seatbottomcurveshape := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _seatbottomcurveshape.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_seatbottomcurveshape, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.SeatBottomCurveShape](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "SeatTopCurveShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.SeatTopCurveShape](probe.stageOfInterest)
			count := 0
			for _seattopcurveshape := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _seattopcurveshape.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_seattopcurveshape, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.SeatTopCurveShape](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "ShiftedBottomTopStartArcShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.ShiftedBottomTopStartArcShape](probe.stageOfInterest)
			count := 0
			for _shiftedbottomtopstartarcshape := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _shiftedbottomtopstartarcshape.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_shiftedbottomtopstartarcshape, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.ShiftedBottomTopStartArcShape](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "ShiftedBottomTopStartArcShapeGrid":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.ShiftedBottomTopStartArcShapeGrid](probe.stageOfInterest)
			count := 0
			for _shiftedbottomtopstartarcshapegrid := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _shiftedbottomtopstartarcshapegrid.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_shiftedbottomtopstartarcshapegrid, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.ShiftedBottomTopStartArcShapeGrid](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "ShiftedLeftGrowthCurve2DRibbon":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.ShiftedLeftGrowthCurve2DRibbon](probe.stageOfInterest)
			count := 0
			for _shiftedleftgrowthcurve2dribbon := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _shiftedleftgrowthcurve2dribbon.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_shiftedleftgrowthcurve2dribbon, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.ShiftedLeftGrowthCurve2DRibbon](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "ShiftedLeftGrowthCurve2DRibbonEndShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.ShiftedLeftGrowthCurve2DRibbonEndShape](probe.stageOfInterest)
			count := 0
			for _shiftedleftgrowthcurve2dribbonendshape := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _shiftedleftgrowthcurve2dribbonendshape.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_shiftedleftgrowthcurve2dribbonendshape, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.ShiftedLeftGrowthCurve2DRibbonEndShape](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "ShiftedLeftGrowthCurve2DRibbonStartShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.ShiftedLeftGrowthCurve2DRibbonStartShape](probe.stageOfInterest)
			count := 0
			for _shiftedleftgrowthcurve2dribbonstartshape := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _shiftedleftgrowthcurve2dribbonstartshape.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_shiftedleftgrowthcurve2dribbonstartshape, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.ShiftedLeftGrowthCurve2DRibbonStartShape](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "ShiftedLeftPartiallyGrowthCurve2DRibbon":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.ShiftedLeftPartiallyGrowthCurve2DRibbon](probe.stageOfInterest)
			count := 0
			for _shiftedleftpartiallygrowthcurve2dribbon := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _shiftedleftpartiallygrowthcurve2dribbon.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_shiftedleftpartiallygrowthcurve2dribbon, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.ShiftedLeftPartiallyGrowthCurve2DRibbon](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "ShiftedLeftPartiallyGrowthCurve2DRibbonEndShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.ShiftedLeftPartiallyGrowthCurve2DRibbonEndShape](probe.stageOfInterest)
			count := 0
			for _shiftedleftpartiallygrowthcurve2dribbonendshape := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _shiftedleftpartiallygrowthcurve2dribbonendshape.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_shiftedleftpartiallygrowthcurve2dribbonendshape, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.ShiftedLeftPartiallyGrowthCurve2DRibbonEndShape](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "ShiftedLeftPartiallyGrowthCurve2DRibbonStartShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.ShiftedLeftPartiallyGrowthCurve2DRibbonStartShape](probe.stageOfInterest)
			count := 0
			for _shiftedleftpartiallygrowthcurve2dribbonstartshape := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _shiftedleftpartiallygrowthcurve2dribbonstartshape.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_shiftedleftpartiallygrowthcurve2dribbonstartshape, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.ShiftedLeftPartiallyGrowthCurve2DRibbonStartShape](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "ShiftedLeftStackGrowthCurveEndArcShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.ShiftedLeftStackGrowthCurveEndArcShape](probe.stageOfInterest)
			count := 0
			for _shiftedleftstackgrowthcurveendarcshape := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _shiftedleftstackgrowthcurveendarcshape.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_shiftedleftstackgrowthcurveendarcshape, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.ShiftedLeftStackGrowthCurveEndArcShape](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "ShiftedLeftStackGrowthCurveStartArcShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.ShiftedLeftStackGrowthCurveStartArcShape](probe.stageOfInterest)
			count := 0
			for _shiftedleftstackgrowthcurvestartarcshape := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _shiftedleftstackgrowthcurvestartarcshape.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_shiftedleftstackgrowthcurvestartarcshape, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.ShiftedLeftStackGrowthCurveStartArcShape](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "ShiftedLeftStackNormalVector":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.ShiftedLeftStackNormalVector](probe.stageOfInterest)
			count := 0
			for _shiftedleftstacknormalvector := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _shiftedleftstacknormalvector.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_shiftedleftstacknormalvector, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.ShiftedLeftStackNormalVector](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "ShiftedLeftStackOfGrowthCurve":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.ShiftedLeftStackOfGrowthCurve](probe.stageOfInterest)
			count := 0
			for _shiftedleftstackofgrowthcurve := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _shiftedleftstackofgrowthcurve.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_shiftedleftstackofgrowthcurve, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.ShiftedLeftStackOfGrowthCurve](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "ShiftedLeftStackOfNormalVector":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.ShiftedLeftStackOfNormalVector](probe.stageOfInterest)
			count := 0
			for _shiftedleftstackofnormalvector := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _shiftedleftstackofnormalvector.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_shiftedleftstackofnormalvector, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.ShiftedLeftStackOfNormalVector](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "ShiftedRightGrowthCurve2DRibbon":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.ShiftedRightGrowthCurve2DRibbon](probe.stageOfInterest)
			count := 0
			for _shiftedrightgrowthcurve2dribbon := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _shiftedrightgrowthcurve2dribbon.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_shiftedrightgrowthcurve2dribbon, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.ShiftedRightGrowthCurve2DRibbon](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "ShiftedRightGrowthCurve2DRibbonEndShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.ShiftedRightGrowthCurve2DRibbonEndShape](probe.stageOfInterest)
			count := 0
			for _shiftedrightgrowthcurve2dribbonendshape := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _shiftedrightgrowthcurve2dribbonendshape.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_shiftedrightgrowthcurve2dribbonendshape, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.ShiftedRightGrowthCurve2DRibbonEndShape](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "ShiftedRightGrowthCurve2DRibbonStartShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.ShiftedRightGrowthCurve2DRibbonStartShape](probe.stageOfInterest)
			count := 0
			for _shiftedrightgrowthcurve2dribbonstartshape := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _shiftedrightgrowthcurve2dribbonstartshape.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_shiftedrightgrowthcurve2dribbonstartshape, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.ShiftedRightGrowthCurve2DRibbonStartShape](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "StackGrowthCurve2DEndHalfwayArcShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.StackGrowthCurve2DEndHalfwayArcShape](probe.stageOfInterest)
			count := 0
			for _stackgrowthcurve2dendhalfwayarcshape := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _stackgrowthcurve2dendhalfwayarcshape.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_stackgrowthcurve2dendhalfwayarcshape, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.StackGrowthCurve2DEndHalfwayArcShape](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "StackGrowthCurve2DRibbonEndShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.StackGrowthCurve2DRibbonEndShape](probe.stageOfInterest)
			count := 0
			for _stackgrowthcurve2dribbonendshape := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _stackgrowthcurve2dribbonendshape.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_stackgrowthcurve2dribbonendshape, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.StackGrowthCurve2DRibbonEndShape](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "StackGrowthCurve2DRibbonStartShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.StackGrowthCurve2DRibbonStartShape](probe.stageOfInterest)
			count := 0
			for _stackgrowthcurve2dribbonstartshape := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _stackgrowthcurve2dribbonstartshape.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_stackgrowthcurve2dribbonstartshape, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.StackGrowthCurve2DRibbonStartShape](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "StackGrowthCurve2DStartHalfwayArcShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.StackGrowthCurve2DStartHalfwayArcShape](probe.stageOfInterest)
			count := 0
			for _stackgrowthcurve2dstarthalfwayarcshape := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _stackgrowthcurve2dstarthalfwayarcshape.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_stackgrowthcurve2dstarthalfwayarcshape, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.StackGrowthCurve2DStartHalfwayArcShape](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "StackOfGrowthCurve2D":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.StackOfGrowthCurve2D](probe.stageOfInterest)
			count := 0
			for _stackofgrowthcurve2d := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _stackofgrowthcurve2d.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_stackofgrowthcurve2d, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.StackOfGrowthCurve2D](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "StackOfGrowthCurve2DByGrowthVector":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.StackOfGrowthCurve2DByGrowthVector](probe.stageOfInterest)
			count := 0
			for _stackofgrowthcurve2dbygrowthvector := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _stackofgrowthcurve2dbygrowthvector.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_stackofgrowthcurve2dbygrowthvector, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.StackOfGrowthCurve2DByGrowthVector](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "StackOfGrowthCurve2DRibbon":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.StackOfGrowthCurve2DRibbon](probe.stageOfInterest)
			count := 0
			for _stackofgrowthcurve2dribbon := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _stackofgrowthcurve2dribbon.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_stackofgrowthcurve2dribbon, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.StackOfGrowthCurve2DRibbon](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "StackOfPartiallyRotatedTorusShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.StackOfPartiallyRotatedTorusShape](probe.stageOfInterest)
			count := 0
			for _stackofpartiallyrotatedtorusshape := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _stackofpartiallyrotatedtorusshape.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_stackofpartiallyrotatedtorusshape, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.StackOfPartiallyRotatedTorusShape](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "StackOfRotatedGrowthCurve2D":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.StackOfRotatedGrowthCurve2D](probe.stageOfInterest)
			count := 0
			for _stackofrotatedgrowthcurve2d := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _stackofrotatedgrowthcurve2d.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_stackofrotatedgrowthcurve2d, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.StackOfRotatedGrowthCurve2D](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "StackOfRotatedGrowthCurve2DRibbon":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.StackOfRotatedGrowthCurve2DRibbon](probe.stageOfInterest)
			count := 0
			for _stackofrotatedgrowthcurve2dribbon := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _stackofrotatedgrowthcurve2dribbon.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_stackofrotatedgrowthcurve2dribbon, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.StackOfRotatedGrowthCurve2DRibbon](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "StackRotatedGrowthCurve2DEndArcShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.StackRotatedGrowthCurve2DEndArcShape](probe.stageOfInterest)
			count := 0
			for _stackrotatedgrowthcurve2dendarcshape := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _stackrotatedgrowthcurve2dendarcshape.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_stackrotatedgrowthcurve2dendarcshape, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.StackRotatedGrowthCurve2DEndArcShape](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "StackRotatedGrowthCurve2DRibbonEndShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.StackRotatedGrowthCurve2DRibbonEndShape](probe.stageOfInterest)
			count := 0
			for _stackrotatedgrowthcurve2dribbonendshape := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _stackrotatedgrowthcurve2dribbonendshape.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_stackrotatedgrowthcurve2dribbonendshape, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.StackRotatedGrowthCurve2DRibbonEndShape](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "StackRotatedGrowthCurve2DRibbonStartShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.StackRotatedGrowthCurve2DRibbonStartShape](probe.stageOfInterest)
			count := 0
			for _stackrotatedgrowthcurve2dribbonstartshape := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _stackrotatedgrowthcurve2dribbonstartshape.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_stackrotatedgrowthcurve2dribbonstartshape, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.StackRotatedGrowthCurve2DRibbonStartShape](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "StackRotatedGrowthCurve2DStartArcShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.StackRotatedGrowthCurve2DStartArcShape](probe.stageOfInterest)
			count := 0
			for _stackrotatedgrowthcurve2dstartarcshape := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _stackrotatedgrowthcurve2dstartarcshape.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_stackrotatedgrowthcurve2dstartarcshape, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.StackRotatedGrowthCurve2DStartArcShape](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "StartArcShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.StartArcShape](probe.stageOfInterest)
			count := 0
			for _startarcshape := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _startarcshape.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_startarcshape, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.StartArcShape](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "StartArcShapeGrid":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.StartArcShapeGrid](probe.stageOfInterest)
			count := 0
			for _startarcshapegrid := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _startarcshapegrid.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_startarcshapegrid, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.StartArcShapeGrid](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "StartHalfwayArcShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.StartHalfwayArcShape](probe.stageOfInterest)
			count := 0
			for _starthalfwayarcshape := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _starthalfwayarcshape.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_starthalfwayarcshape, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.StartHalfwayArcShape](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "StartHalfwayArcShapeGrid":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.StartHalfwayArcShapeGrid](probe.stageOfInterest)
			count := 0
			for _starthalfwayarcshapegrid := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _starthalfwayarcshapegrid.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_starthalfwayarcshapegrid, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.StartHalfwayArcShapeGrid](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "StoolAbstract":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.StoolAbstract](probe.stageOfInterest)
			count := 0
			for _stoolabstract := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _stoolabstract.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_stoolabstract, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.StoolAbstract](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "StoolDiagram":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.StoolDiagram](probe.stageOfInterest)
			count := 0
			for _stooldiagram := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _stooldiagram.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_stooldiagram, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.StoolDiagram](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "TopEndArcShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.TopEndArcShape](probe.stageOfInterest)
			count := 0
			for _topendarcshape := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _topendarcshape.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_topendarcshape, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.TopEndArcShape](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "TopEndArcShapeGrid":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.TopEndArcShapeGrid](probe.stageOfInterest)
			count := 0
			for _topendarcshapegrid := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _topendarcshapegrid.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_topendarcshapegrid, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.TopEndArcShapeGrid](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "TopEndHalfwayArcShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.TopEndHalfwayArcShape](probe.stageOfInterest)
			count := 0
			for _topendhalfwayarcshape := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _topendhalfwayarcshape.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_topendhalfwayarcshape, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.TopEndHalfwayArcShape](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "TopEndHalfwayArcShapeGrid":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.TopEndHalfwayArcShapeGrid](probe.stageOfInterest)
			count := 0
			for _topendhalfwayarcshapegrid := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _topendhalfwayarcshapegrid.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_topendhalfwayarcshapegrid, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.TopEndHalfwayArcShapeGrid](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "TopGrowthCurve2D":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.TopGrowthCurve2D](probe.stageOfInterest)
			count := 0
			for _topgrowthcurve2d := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _topgrowthcurve2d.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_topgrowthcurve2d, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.TopGrowthCurve2D](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "TopMidArcVectorShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.TopMidArcVectorShape](probe.stageOfInterest)
			count := 0
			for _topmidarcvectorshape := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _topmidarcvectorshape.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_topmidarcvectorshape, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.TopMidArcVectorShape](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "TopMidArcVectorShapeGrid":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.TopMidArcVectorShapeGrid](probe.stageOfInterest)
			count := 0
			for _topmidarcvectorshapegrid := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _topmidarcvectorshapegrid.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_topmidarcvectorshapegrid, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.TopMidArcVectorShapeGrid](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "TopStackGrowthCurve2DEndHalfwayArcShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.TopStackGrowthCurve2DEndHalfwayArcShape](probe.stageOfInterest)
			count := 0
			for _topstackgrowthcurve2dendhalfwayarcshape := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _topstackgrowthcurve2dendhalfwayarcshape.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_topstackgrowthcurve2dendhalfwayarcshape, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.TopStackGrowthCurve2DEndHalfwayArcShape](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "TopStackGrowthCurve2DStartHalfwayArcShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.TopStackGrowthCurve2DStartHalfwayArcShape](probe.stageOfInterest)
			count := 0
			for _topstackgrowthcurve2dstarthalfwayarcshape := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _topstackgrowthcurve2dstarthalfwayarcshape.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_topstackgrowthcurve2dstarthalfwayarcshape, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.TopStackGrowthCurve2DStartHalfwayArcShape](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "TopStackOfGrowthCurve2D":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.TopStackOfGrowthCurve2D](probe.stageOfInterest)
			count := 0
			for _topstackofgrowthcurve2d := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _topstackofgrowthcurve2d.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_topstackofgrowthcurve2d, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.TopStackOfGrowthCurve2D](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "TopStackOfRotatedGrowthCurve2D":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.TopStackOfRotatedGrowthCurve2D](probe.stageOfInterest)
			count := 0
			for _topstackofrotatedgrowthcurve2d := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _topstackofrotatedgrowthcurve2d.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_topstackofrotatedgrowthcurve2d, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.TopStackOfRotatedGrowthCurve2D](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "TopStackOfRotatedGrowthCurve2DEndArcShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.TopStackOfRotatedGrowthCurve2DEndArcShape](probe.stageOfInterest)
			count := 0
			for _topstackofrotatedgrowthcurve2dendarcshape := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _topstackofrotatedgrowthcurve2dendarcshape.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_topstackofrotatedgrowthcurve2dendarcshape, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.TopStackOfRotatedGrowthCurve2DEndArcShape](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "TopStackOfRotatedGrowthCurve2DStartArcShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.TopStackOfRotatedGrowthCurve2DStartArcShape](probe.stageOfInterest)
			count := 0
			for _topstackofrotatedgrowthcurve2dstartarcshape := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _topstackofrotatedgrowthcurve2dstartarcshape.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_topstackofrotatedgrowthcurve2dstartarcshape, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.TopStackOfRotatedGrowthCurve2DStartArcShape](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "TopStartArcShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.TopStartArcShape](probe.stageOfInterest)
			count := 0
			for _topstartarcshape := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _topstartarcshape.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_topstartarcshape, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.TopStartArcShape](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "TopStartArcShapeGrid":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.TopStartArcShapeGrid](probe.stageOfInterest)
			count := 0
			for _topstartarcshapegrid := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _topstartarcshapegrid.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_topstartarcshapegrid, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.TopStartArcShapeGrid](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "TopStartHalfwayArcShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.TopStartHalfwayArcShape](probe.stageOfInterest)
			count := 0
			for _topstarthalfwayarcshape := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _topstarthalfwayarcshape.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_topstarthalfwayarcshape, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.TopStartHalfwayArcShape](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "TopStartHalfwayArcShapeGrid":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.TopStartHalfwayArcShapeGrid](probe.stageOfInterest)
			count := 0
			for _topstarthalfwayarcshapegrid := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _topstarthalfwayarcshapegrid.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_topstarthalfwayarcshapegrid, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.TopStartHalfwayArcShapeGrid](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "Torus3DShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Torus3DShape](probe.stageOfInterest)
			count := 0
			for _torus3dshape := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _torus3dshape.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_torus3dshape, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.Torus3DShape](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "TorusEdge3DShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.TorusEdge3DShape](probe.stageOfInterest)
			count := 0
			for _torusedge3dshape := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _torusedge3dshape.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_torusedge3dshape, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.TorusEdge3DShape](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "TorusStackShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.TorusStackShape](probe.stageOfInterest)
			count := 0
			for _torusstackshape := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _torusstackshape.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_torusstackshape, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.TorusStackShape](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "VaseAbstract":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.VaseAbstract](probe.stageOfInterest)
			count := 0
			for _vaseabstract := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _vaseabstract.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_vaseabstract, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.VaseAbstract](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "VaseDiagram":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.VaseDiagram](probe.stageOfInterest)
			count := 0
			for _vasediagram := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _vasediagram.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_vasediagram, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.VaseDiagram](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "VerticalTorusStackShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.VerticalTorusStackShape](probe.stageOfInterest)
			count := 0
			for _verticaltorusstackshape := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _verticaltorusstackshape.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_verticaltorusstackshape, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.VerticalTorusStackShape](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "VolumeKey3DShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.VolumeKey3DShape](probe.stageOfInterest)
			count := 0
			for _volumekey3dshape := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _volumekey3dshape.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_volumekey3dshape, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.VolumeKey3DShape](probe)
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
