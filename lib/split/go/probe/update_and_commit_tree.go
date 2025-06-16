package probe

import (
	"fmt"
	"sort"
	"strings"
	"time"

	gongtree_buttons "github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"

	gong_models "github.com/fullstack-lang/gong/go/models"

	"github.com/fullstack-lang/gong/lib/split/go/models"
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
		case "AsSplit":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.AsSplit](probe.stageOfInterest)
			for _assplit := range set {
				nodeInstance := &tree.Node{Name: _assplit.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_assplit, "AsSplit", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "AsSplitArea":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.AsSplitArea](probe.stageOfInterest)
			for _assplitarea := range set {
				nodeInstance := &tree.Node{Name: _assplitarea.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_assplitarea, "AsSplitArea", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Button":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Button](probe.stageOfInterest)
			for _button := range set {
				nodeInstance := &tree.Node{Name: _button.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_button, "Button", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Cursor":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Cursor](probe.stageOfInterest)
			for _cursor := range set {
				nodeInstance := &tree.Node{Name: _cursor.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_cursor, "Cursor", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Doc":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Doc](probe.stageOfInterest)
			for _doc := range set {
				nodeInstance := &tree.Node{Name: _doc.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_doc, "Doc", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "FavIcon":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.FavIcon](probe.stageOfInterest)
			for _favicon := range set {
				nodeInstance := &tree.Node{Name: _favicon.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_favicon, "FavIcon", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Form":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Form](probe.stageOfInterest)
			for _form := range set {
				nodeInstance := &tree.Node{Name: _form.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_form, "Form", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Load":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Load](probe.stageOfInterest)
			for _load := range set {
				nodeInstance := &tree.Node{Name: _load.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_load, "Load", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "LogoOnTheLeft":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.LogoOnTheLeft](probe.stageOfInterest)
			for _logoontheleft := range set {
				nodeInstance := &tree.Node{Name: _logoontheleft.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_logoontheleft, "LogoOnTheLeft", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "LogoOnTheRight":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.LogoOnTheRight](probe.stageOfInterest)
			for _logoontheright := range set {
				nodeInstance := &tree.Node{Name: _logoontheright.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_logoontheright, "LogoOnTheRight", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Slider":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Slider](probe.stageOfInterest)
			for _slider := range set {
				nodeInstance := &tree.Node{Name: _slider.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_slider, "Slider", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Split":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Split](probe.stageOfInterest)
			for _split := range set {
				nodeInstance := &tree.Node{Name: _split.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_split, "Split", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Svg":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Svg](probe.stageOfInterest)
			for _svg := range set {
				nodeInstance := &tree.Node{Name: _svg.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_svg, "Svg", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Table":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Table](probe.stageOfInterest)
			for _table := range set {
				nodeInstance := &tree.Node{Name: _table.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_table, "Table", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Title":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Title](probe.stageOfInterest)
			for _title := range set {
				nodeInstance := &tree.Node{Name: _title.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_title, "Title", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Tone":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Tone](probe.stageOfInterest)
			for _tone := range set {
				nodeInstance := &tree.Node{Name: _tone.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_tone, "Tone", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Tree":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Tree](probe.stageOfInterest)
			for _tree := range set {
				nodeInstance := &tree.Node{Name: _tree.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_tree, "Tree", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "View":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.View](probe.stageOfInterest)
			for _view := range set {
				nodeInstance := &tree.Node{Name: _view.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_view, "View", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Xlsx":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Xlsx](probe.stageOfInterest)
			for _xlsx := range set {
				nodeInstance := &tree.Node{Name: _xlsx.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_xlsx, "Xlsx", probe)

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
