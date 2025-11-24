package probe

import (
	"fmt"
	"sort"
	"strings"
	"time"

	gongtree_buttons "github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"

	gong_models "github.com/fullstack-lang/gong/go/models"

	"github.com/fullstack-lang/gong/lib/markdown/go/models"
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

	// create tree
	sidebar := &tree.Tree{Name: "Sidebar"}

	nodeRefreshButton := &tree.Node{Name: fmt.Sprintf("Stage %s, # %d, %s",
		probe.stageOfInterest.GetName(),
		probe.stageOfInterest.GetCommitId(),
		probe.stageOfInterest.GetCommitTS().Local().Format(time.Kitchen))}
	nodeRefreshButton.Name +=
		fmt.Sprintf(" (C%d/U%d/D%d)", 
			len(probe.stageOfInterest.GetNew()), 
			len(probe.stageOfInterest.GetModified()), 
			len(probe.stageOfInterest.GetDeleted()),
		)
	sidebar.RootNodes = append(sidebar.RootNodes, nodeRefreshButton)
	refreshButton := &tree.Button{
		Name:            "RefreshButton" + " " + string(gongtree_buttons.BUTTON_refresh),
		Icon:            string(gongtree_buttons.BUTTON_refresh),
		HasToolTip:      true,
		ToolTipText:     "Refresh probe",
		ToolTipPosition: tree.Left,
	}

	nodeRefreshButton.Buttons = append(nodeRefreshButton.Buttons, refreshButton)
	refreshButton.Impl = NewButtonImplRefresh(probe)

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
		case "Content":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Content](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _content := range set {
				nodeInstance := &tree.Node{Name: _content.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_content, "Content", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_content]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_content]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_content]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (C%d/U%d/D%d)", created, updated, deleted)
		case "JpgImage":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.JpgImage](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _jpgimage := range set {
				nodeInstance := &tree.Node{Name: _jpgimage.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_jpgimage, "JpgImage", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_jpgimage]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_jpgimage]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_jpgimage]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (C%d/U%d/D%d)", created, updated, deleted)
		case "PngImage":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.PngImage](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _pngimage := range set {
				nodeInstance := &tree.Node{Name: _pngimage.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_pngimage, "PngImage", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_pngimage]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_pngimage]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_pngimage]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (C%d/U%d/D%d)", created, updated, deleted)
		case "SvgImage":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.SvgImage](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _svgimage := range set {
				nodeInstance := &tree.Node{Name: _svgimage.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_svgimage, "SvgImage", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_svgimage]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_svgimage]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_svgimage]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (C%d/U%d/D%d)", created, updated, deleted)
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
		addButton.Impl = NewButtonImplGongstruct(
			gongStruct,
			gongtree_buttons.BUTTON_add,
			probe,
		)

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
