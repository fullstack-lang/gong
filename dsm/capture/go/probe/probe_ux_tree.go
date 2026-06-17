package probe

import (
	"fmt"
	"sort"
	"strings"
	"time"

	tree_buttons "github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree_models "github.com/fullstack-lang/gong/lib/tree/go/models"

	gong_models "github.com/fullstack-lang/gong/go/models"

	"github.com/fullstack-lang/gong/dsm/capture/go/models"
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
		case "AnalysisNeed":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.AnalysisNeed](probe.stageOfInterest)
			count := 0
			for _analysisneed := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _analysisneed.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_analysisneed, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.AnalysisNeed](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "Concept":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Concept](probe.stageOfInterest)
			count := 0
			for _concept := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _concept.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_concept, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.Concept](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "Concern":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Concern](probe.stageOfInterest)
			count := 0
			for _concern := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _concern.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_concern, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.Concern](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "ConcernCompositionShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.ConcernCompositionShape](probe.stageOfInterest)
			count := 0
			for _concerncompositionshape := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _concerncompositionshape.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_concerncompositionshape, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.ConcernCompositionShape](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "ConcernInputShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.ConcernInputShape](probe.stageOfInterest)
			count := 0
			for _concerninputshape := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _concerninputshape.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_concerninputshape, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.ConcernInputShape](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "ConcernOutputShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.ConcernOutputShape](probe.stageOfInterest)
			count := 0
			for _concernoutputshape := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _concernoutputshape.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_concernoutputshape, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.ConcernOutputShape](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "ConcernShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.ConcernShape](probe.stageOfInterest)
			count := 0
			for _concernshape := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _concernshape.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_concernshape, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.ConcernShape](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "Deliverable":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Deliverable](probe.stageOfInterest)
			count := 0
			for _deliverable := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _deliverable.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_deliverable, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.Deliverable](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "Diagram":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Diagram](probe.stageOfInterest)
			count := 0
			for _diagram := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _diagram.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_diagram, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.Diagram](probe)
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
		case "Note":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Note](probe.stageOfInterest)
			count := 0
			for _note := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _note.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_note, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.Note](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "NoteProductShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.NoteProductShape](probe.stageOfInterest)
			count := 0
			for _noteproductshape := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _noteproductshape.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_noteproductshape, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.NoteProductShape](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "NoteShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.NoteShape](probe.stageOfInterest)
			count := 0
			for _noteshape := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _noteshape.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_noteshape, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.NoteShape](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "NoteStakeholderShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.NoteStakeholderShape](probe.stageOfInterest)
			count := 0
			for _notestakeholdershape := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _notestakeholdershape.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_notestakeholdershape, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.NoteStakeholderShape](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "NoteTaskShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.NoteTaskShape](probe.stageOfInterest)
			count := 0
			for _notetaskshape := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _notetaskshape.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_notetaskshape, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.NoteTaskShape](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "ProductCompositionShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.ProductCompositionShape](probe.stageOfInterest)
			count := 0
			for _productcompositionshape := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _productcompositionshape.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_productcompositionshape, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.ProductCompositionShape](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "ProductShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.ProductShape](probe.stageOfInterest)
			count := 0
			for _productshape := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _productshape.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_productshape, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.ProductShape](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "Requirement":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Requirement](probe.stageOfInterest)
			count := 0
			for _requirement := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _requirement.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_requirement, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.Requirement](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "Stakeholder":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Stakeholder](probe.stageOfInterest)
			count := 0
			for _stakeholder := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _stakeholder.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_stakeholder, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.Stakeholder](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "StakeholderCompositionShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.StakeholderCompositionShape](probe.stageOfInterest)
			count := 0
			for _stakeholdercompositionshape := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _stakeholdercompositionshape.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_stakeholdercompositionshape, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.StakeholderCompositionShape](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "StakeholderConcernShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.StakeholderConcernShape](probe.stageOfInterest)
			count := 0
			for _stakeholderconcernshape := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _stakeholderconcernshape.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_stakeholderconcernshape, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.StakeholderConcernShape](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "StakeholderShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.StakeholderShape](probe.stageOfInterest)
			count := 0
			for _stakeholdershape := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _stakeholdershape.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_stakeholdershape, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.StakeholderShape](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "SupportLevel":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.SupportLevel](probe.stageOfInterest)
			count := 0
			for _supportlevel := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _supportlevel.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_supportlevel, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.SupportLevel](probe)
				// set color for node and reset all other nodes color
				for node := range *tree_models.GetGongstructInstancesSet[tree_models.Node](probe.treeStage) {
					node.BackgroundColor = ""
				}
				nodeGongstruct.BackgroundColor = "lightgrey"
				probe.treeStage.Commit()
			}
		case "Tool":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Tool](probe.stageOfInterest)
			count := 0
			for _tool := range set {
				if count >= probe.GetMaxElementsNbPerGongStructNode() {
					nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
					break
				}
				count++
				nodeInstance := &tree_models.Node{
					Name:            _tool.GetName(),
					IsNodeClickable: true,
					OnClick: func(frontNode *tree_models.Node) {
						FillUpFormFromGongstruct(_tool, probe)
					},
				}
				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
			nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
				nodeGongstruct.IsExpanded = isExpanded
				// no commit, it will be done in the refresh
			}
			nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
				updateProbeTable[*models.Tool](probe)
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
