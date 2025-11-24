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
		case "AsSplit":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.AsSplit](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _assplit := range set {
				nodeInstance := &tree.Node{Name: _assplit.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_assplit, "AsSplit", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_assplit]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_assplit]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_assplit]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "AsSplitArea":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.AsSplitArea](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _assplitarea := range set {
				nodeInstance := &tree.Node{Name: _assplitarea.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_assplitarea, "AsSplitArea", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_assplitarea]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_assplitarea]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_assplitarea]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "Button":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Button](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _button := range set {
				nodeInstance := &tree.Node{Name: _button.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_button, "Button", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_button]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_button]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_button]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "Cursor":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Cursor](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _cursor := range set {
				nodeInstance := &tree.Node{Name: _cursor.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_cursor, "Cursor", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_cursor]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_cursor]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_cursor]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "FavIcon":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.FavIcon](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _favicon := range set {
				nodeInstance := &tree.Node{Name: _favicon.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_favicon, "FavIcon", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_favicon]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_favicon]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_favicon]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "Form":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Form](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _form := range set {
				nodeInstance := &tree.Node{Name: _form.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_form, "Form", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_form]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_form]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_form]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "Load":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Load](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _load := range set {
				nodeInstance := &tree.Node{Name: _load.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_load, "Load", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_load]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_load]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_load]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "LogoOnTheLeft":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.LogoOnTheLeft](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _logoontheleft := range set {
				nodeInstance := &tree.Node{Name: _logoontheleft.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_logoontheleft, "LogoOnTheLeft", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_logoontheleft]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_logoontheleft]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_logoontheleft]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "LogoOnTheRight":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.LogoOnTheRight](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _logoontheright := range set {
				nodeInstance := &tree.Node{Name: _logoontheright.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_logoontheright, "LogoOnTheRight", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_logoontheright]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_logoontheright]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_logoontheright]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "Markdown":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Markdown](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _markdown := range set {
				nodeInstance := &tree.Node{Name: _markdown.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_markdown, "Markdown", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_markdown]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_markdown]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_markdown]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "Slider":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Slider](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _slider := range set {
				nodeInstance := &tree.Node{Name: _slider.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_slider, "Slider", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_slider]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_slider]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_slider]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "Split":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Split](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _split := range set {
				nodeInstance := &tree.Node{Name: _split.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_split, "Split", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_split]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_split]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_split]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "Svg":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Svg](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _svg := range set {
				nodeInstance := &tree.Node{Name: _svg.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_svg, "Svg", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_svg]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_svg]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_svg]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "Table":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Table](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _table := range set {
				nodeInstance := &tree.Node{Name: _table.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_table, "Table", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_table]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_table]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_table]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "Title":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Title](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _title := range set {
				nodeInstance := &tree.Node{Name: _title.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_title, "Title", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_title]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_title]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_title]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "Tone":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Tone](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _tone := range set {
				nodeInstance := &tree.Node{Name: _tone.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_tone, "Tone", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_tone]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_tone]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_tone]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "Tree":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Tree](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _tree := range set {
				nodeInstance := &tree.Node{Name: _tree.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_tree, "Tree", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_tree]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_tree]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_tree]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "View":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.View](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _view := range set {
				nodeInstance := &tree.Node{Name: _view.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_view, "View", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_view]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_view]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_view]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "Xlsx":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Xlsx](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _xlsx := range set {
				nodeInstance := &tree.Node{Name: _xlsx.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_xlsx, "Xlsx", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_xlsx]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_xlsx]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_xlsx]; ok {
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
