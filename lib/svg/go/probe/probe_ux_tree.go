package probe

import (
	"fmt"
	"sort"
	"strings"
	"time"

	tree_buttons "github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree_models "github.com/fullstack-lang/gong/lib/tree/go/models"

	gong_models "github.com/fullstack-lang/gong/go/models"

	"github.com/fullstack-lang/gong/lib/svg/go/models"
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
		case "Animate":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Animate](probe.stageOfInterest)
			count := 0
			for _animate := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _animate.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_animate, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.Animate](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "Circle":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Circle](probe.stageOfInterest)
			count := 0
			for _circle := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _circle.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_circle, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.Circle](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "Condition":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Condition](probe.stageOfInterest)
			count := 0
			for _condition := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _condition.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_condition, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.Condition](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "ControlPoint":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.ControlPoint](probe.stageOfInterest)
			count := 0
			for _controlpoint := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _controlpoint.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_controlpoint, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.ControlPoint](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "Ellipse":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Ellipse](probe.stageOfInterest)
			count := 0
			for _ellipse := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _ellipse.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_ellipse, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.Ellipse](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "Layer":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Layer](probe.stageOfInterest)
			count := 0
			for _layer := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _layer.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_layer, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.Layer](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "Line":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Line](probe.stageOfInterest)
			count := 0
			for _line := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _line.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_line, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.Line](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "Link":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Link](probe.stageOfInterest)
			count := 0
			for _link := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _link.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_link, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.Link](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "LinkAnchoredText":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.LinkAnchoredText](probe.stageOfInterest)
			count := 0
			for _linkanchoredtext := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _linkanchoredtext.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_linkanchoredtext, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.LinkAnchoredText](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "Path":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Path](probe.stageOfInterest)
			count := 0
			for _path := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _path.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_path, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.Path](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "Point":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Point](probe.stageOfInterest)
			count := 0
			for _point := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _point.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_point, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.Point](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "Polygone":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Polygone](probe.stageOfInterest)
			count := 0
			for _polygone := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _polygone.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_polygone, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.Polygone](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "Polyline":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Polyline](probe.stageOfInterest)
			count := 0
			for _polyline := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _polyline.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_polyline, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.Polyline](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "Rect":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Rect](probe.stageOfInterest)
			count := 0
			for _rect := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _rect.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_rect, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.Rect](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "RectAnchoredJpgImage":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.RectAnchoredJpgImage](probe.stageOfInterest)
			count := 0
			for _rectanchoredjpgimage := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _rectanchoredjpgimage.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_rectanchoredjpgimage, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.RectAnchoredJpgImage](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "RectAnchoredPath":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.RectAnchoredPath](probe.stageOfInterest)
			count := 0
			for _rectanchoredpath := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _rectanchoredpath.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_rectanchoredpath, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.RectAnchoredPath](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "RectAnchoredPngImage":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.RectAnchoredPngImage](probe.stageOfInterest)
			count := 0
			for _rectanchoredpngimage := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _rectanchoredpngimage.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_rectanchoredpngimage, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.RectAnchoredPngImage](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "RectAnchoredRect":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.RectAnchoredRect](probe.stageOfInterest)
			count := 0
			for _rectanchoredrect := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _rectanchoredrect.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_rectanchoredrect, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.RectAnchoredRect](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "RectAnchoredSvgImage":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.RectAnchoredSvgImage](probe.stageOfInterest)
			count := 0
			for _rectanchoredsvgimage := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _rectanchoredsvgimage.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_rectanchoredsvgimage, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.RectAnchoredSvgImage](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "RectAnchoredText":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.RectAnchoredText](probe.stageOfInterest)
			count := 0
			for _rectanchoredtext := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _rectanchoredtext.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_rectanchoredtext, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.RectAnchoredText](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "RectLinkLink":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.RectLinkLink](probe.stageOfInterest)
			count := 0
			for _rectlinklink := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _rectlinklink.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_rectlinklink, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.RectLinkLink](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "SVG":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.SVG](probe.stageOfInterest)
			count := 0
			for _svg := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _svg.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_svg, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.SVG](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "SvgText":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.SvgText](probe.stageOfInterest)
			count := 0
			for _svgtext := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _svgtext.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_svgtext, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.SvgText](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](treeStagee) {
					node.BackgroundColor = ""
				}
				stagedNode.BackgroundColor = "lightgrey"
				treeStagee.Commit()
			}
		case "Text":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Text](probe.stageOfInterest)
			count := 0
			for _text := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _text.GetName(),
					IsNodeClickable: true,
					OnUpdate: func(_ *tree_models.Stage, _, _ *tree_models.Node) {
						FillUpFormFromGongstruct(_text, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnUpdate = func(treeStagee *tree_models.Stage, stagedNode, frontNode *tree_models.Node) {
				if stagedNode.IsExpanded != frontNode.IsExpanded {
					stagedNode.IsExpanded = frontNode.IsExpanded
					return
				}
				updateProbeTable[*models.Text](probe)
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
