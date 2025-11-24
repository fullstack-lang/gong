package probe

import (
	"fmt"
	"sort"
	"strings"
	"time"

	gongtree_buttons "github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"

	gong_models "github.com/fullstack-lang/gong/go/models"

	"github.com/fullstack-lang/gong/lib/doc2/go/models"
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
		fmt.Sprintf(" (%d/%d/%d)", 
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
		case "AttributeShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.AttributeShape](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _attributeshape := range set {
				nodeInstance := &tree.Node{Name: _attributeshape.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_attributeshape, "AttributeShape", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_attributeshape]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_attributeshape]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_attributeshape]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "Classdiagram":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Classdiagram](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _classdiagram := range set {
				nodeInstance := &tree.Node{Name: _classdiagram.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_classdiagram, "Classdiagram", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_classdiagram]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_classdiagram]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_classdiagram]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "DiagramPackage":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.DiagramPackage](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _diagrampackage := range set {
				nodeInstance := &tree.Node{Name: _diagrampackage.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_diagrampackage, "DiagramPackage", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_diagrampackage]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_diagrampackage]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_diagrampackage]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "GongEnumShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.GongEnumShape](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _gongenumshape := range set {
				nodeInstance := &tree.Node{Name: _gongenumshape.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_gongenumshape, "GongEnumShape", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_gongenumshape]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_gongenumshape]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_gongenumshape]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "GongEnumValueShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.GongEnumValueShape](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _gongenumvalueshape := range set {
				nodeInstance := &tree.Node{Name: _gongenumvalueshape.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_gongenumvalueshape, "GongEnumValueShape", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_gongenumvalueshape]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_gongenumvalueshape]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_gongenumvalueshape]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "GongNoteLinkShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.GongNoteLinkShape](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _gongnotelinkshape := range set {
				nodeInstance := &tree.Node{Name: _gongnotelinkshape.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_gongnotelinkshape, "GongNoteLinkShape", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_gongnotelinkshape]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_gongnotelinkshape]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_gongnotelinkshape]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "GongNoteShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.GongNoteShape](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _gongnoteshape := range set {
				nodeInstance := &tree.Node{Name: _gongnoteshape.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_gongnoteshape, "GongNoteShape", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_gongnoteshape]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_gongnoteshape]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_gongnoteshape]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "GongStructShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.GongStructShape](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _gongstructshape := range set {
				nodeInstance := &tree.Node{Name: _gongstructshape.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_gongstructshape, "GongStructShape", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_gongstructshape]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_gongstructshape]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_gongstructshape]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "LinkShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.LinkShape](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _linkshape := range set {
				nodeInstance := &tree.Node{Name: _linkshape.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_linkshape, "LinkShape", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_linkshape]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_linkshape]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_linkshape]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
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
