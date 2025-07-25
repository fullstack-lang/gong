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

const SideBarTreeName = "gong"

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
	sidebar := &tree.Tree{Name: SideBarTreeName}

	nodeRefreshButton := &tree.Node{Name: fmt.Sprintf("Stage %s, # %d, %s",
		probe.stageOfInterest.GetName(),
		probe.stageOfInterest.GetCommitId(),
		probe.stageOfInterest.GetCommitTS().Local().Format(time.Kitchen))}
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
	setOfGongStructs := *gong_models.GetGongstructInstancesSet[gong_models.GongStruct](probe.gongStage)

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
		case "Classdiagram":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Classdiagram](probe.stageOfInterest)
			for _classdiagram := range set {
				nodeInstance := &tree.Node{Name: _classdiagram.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_classdiagram, "Classdiagram", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "DiagramPackage":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.DiagramPackage](probe.stageOfInterest)
			for _diagrampackage := range set {
				nodeInstance := &tree.Node{Name: _diagrampackage.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_diagrampackage, "DiagramPackage", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Field":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Field](probe.stageOfInterest)
			for _field := range set {
				nodeInstance := &tree.Node{Name: _field.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_field, "Field", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "GongEnumShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.GongEnumShape](probe.stageOfInterest)
			for _gongenumshape := range set {
				nodeInstance := &tree.Node{Name: _gongenumshape.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_gongenumshape, "GongEnumShape", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "GongEnumValueEntry":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.GongEnumValueEntry](probe.stageOfInterest)
			for _gongenumvalueentry := range set {
				nodeInstance := &tree.Node{Name: _gongenumvalueentry.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_gongenumvalueentry, "GongEnumValueEntry", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "GongStructShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.GongStructShape](probe.stageOfInterest)
			for _gongstructshape := range set {
				nodeInstance := &tree.Node{Name: _gongstructshape.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_gongstructshape, "GongStructShape", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Link":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Link](probe.stageOfInterest)
			for _link := range set {
				nodeInstance := &tree.Node{Name: _link.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_link, "Link", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "NoteShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.NoteShape](probe.stageOfInterest)
			for _noteshape := range set {
				nodeInstance := &tree.Node{Name: _noteshape.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_noteshape, "NoteShape", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "NoteShapeLink":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.NoteShapeLink](probe.stageOfInterest)
			for _noteshapelink := range set {
				nodeInstance := &tree.Node{Name: _noteshapelink.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_noteshapelink, "NoteShapeLink", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Position":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Position](probe.stageOfInterest)
			for _position := range set {
				nodeInstance := &tree.Node{Name: _position.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_position, "Position", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "UmlState":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.UmlState](probe.stageOfInterest)
			for _umlstate := range set {
				nodeInstance := &tree.Node{Name: _umlstate.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_umlstate, "UmlState", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Umlsc":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Umlsc](probe.stageOfInterest)
			for _umlsc := range set {
				nodeInstance := &tree.Node{Name: _umlsc.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_umlsc, "Umlsc", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Vertice":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Vertice](probe.stageOfInterest)
			for _vertice := range set {
				nodeInstance := &tree.Node{Name: _vertice.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_vertice, "Vertice", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
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

type InstanceNodeCallback[T models.Gongstruct] struct {
	Instance       *T
	gongstructName string
	probe          *Probe
}

func NewInstanceNodeCallback[T models.Gongstruct](
	instance *T,
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
