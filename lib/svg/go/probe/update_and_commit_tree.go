package probe

import (
	"fmt"
	"sort"
	"strings"
	"time"

	gongtree_buttons "github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"

	gong_models "github.com/fullstack-lang/gong/go/models"

	"github.com/fullstack-lang/gong/lib/svg/go/models"
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
		case "Animate":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Animate](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _animate := range set {
				nodeInstance := &tree.Node{Name: _animate.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_animate, "Animate", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_animate]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_animate]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_animate]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (C%d/U%d/D%d)", created, updated, deleted)
		case "Circle":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Circle](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _circle := range set {
				nodeInstance := &tree.Node{Name: _circle.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_circle, "Circle", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_circle]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_circle]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_circle]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (C%d/U%d/D%d)", created, updated, deleted)
		case "Condition":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Condition](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _condition := range set {
				nodeInstance := &tree.Node{Name: _condition.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_condition, "Condition", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_condition]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_condition]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_condition]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (C%d/U%d/D%d)", created, updated, deleted)
		case "ControlPoint":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.ControlPoint](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _controlpoint := range set {
				nodeInstance := &tree.Node{Name: _controlpoint.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_controlpoint, "ControlPoint", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_controlpoint]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_controlpoint]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_controlpoint]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (C%d/U%d/D%d)", created, updated, deleted)
		case "Ellipse":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Ellipse](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _ellipse := range set {
				nodeInstance := &tree.Node{Name: _ellipse.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_ellipse, "Ellipse", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_ellipse]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_ellipse]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_ellipse]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (C%d/U%d/D%d)", created, updated, deleted)
		case "Layer":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Layer](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _layer := range set {
				nodeInstance := &tree.Node{Name: _layer.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_layer, "Layer", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_layer]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_layer]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_layer]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (C%d/U%d/D%d)", created, updated, deleted)
		case "Line":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Line](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _line := range set {
				nodeInstance := &tree.Node{Name: _line.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_line, "Line", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_line]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_line]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_line]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (C%d/U%d/D%d)", created, updated, deleted)
		case "Link":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Link](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _link := range set {
				nodeInstance := &tree.Node{Name: _link.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_link, "Link", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_link]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_link]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_link]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (C%d/U%d/D%d)", created, updated, deleted)
		case "LinkAnchoredText":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.LinkAnchoredText](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _linkanchoredtext := range set {
				nodeInstance := &tree.Node{Name: _linkanchoredtext.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_linkanchoredtext, "LinkAnchoredText", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_linkanchoredtext]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_linkanchoredtext]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_linkanchoredtext]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (C%d/U%d/D%d)", created, updated, deleted)
		case "Path":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Path](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _path := range set {
				nodeInstance := &tree.Node{Name: _path.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_path, "Path", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_path]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_path]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_path]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (C%d/U%d/D%d)", created, updated, deleted)
		case "Point":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Point](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _point := range set {
				nodeInstance := &tree.Node{Name: _point.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_point, "Point", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_point]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_point]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_point]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (C%d/U%d/D%d)", created, updated, deleted)
		case "Polygone":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Polygone](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _polygone := range set {
				nodeInstance := &tree.Node{Name: _polygone.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_polygone, "Polygone", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_polygone]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_polygone]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_polygone]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (C%d/U%d/D%d)", created, updated, deleted)
		case "Polyline":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Polyline](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _polyline := range set {
				nodeInstance := &tree.Node{Name: _polyline.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_polyline, "Polyline", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_polyline]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_polyline]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_polyline]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (C%d/U%d/D%d)", created, updated, deleted)
		case "Rect":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Rect](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _rect := range set {
				nodeInstance := &tree.Node{Name: _rect.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_rect, "Rect", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_rect]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_rect]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_rect]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (C%d/U%d/D%d)", created, updated, deleted)
		case "RectAnchoredPath":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.RectAnchoredPath](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _rectanchoredpath := range set {
				nodeInstance := &tree.Node{Name: _rectanchoredpath.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_rectanchoredpath, "RectAnchoredPath", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_rectanchoredpath]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_rectanchoredpath]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_rectanchoredpath]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (C%d/U%d/D%d)", created, updated, deleted)
		case "RectAnchoredRect":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.RectAnchoredRect](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _rectanchoredrect := range set {
				nodeInstance := &tree.Node{Name: _rectanchoredrect.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_rectanchoredrect, "RectAnchoredRect", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_rectanchoredrect]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_rectanchoredrect]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_rectanchoredrect]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (C%d/U%d/D%d)", created, updated, deleted)
		case "RectAnchoredText":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.RectAnchoredText](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _rectanchoredtext := range set {
				nodeInstance := &tree.Node{Name: _rectanchoredtext.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_rectanchoredtext, "RectAnchoredText", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_rectanchoredtext]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_rectanchoredtext]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_rectanchoredtext]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (C%d/U%d/D%d)", created, updated, deleted)
		case "RectLinkLink":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.RectLinkLink](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _rectlinklink := range set {
				nodeInstance := &tree.Node{Name: _rectlinklink.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_rectlinklink, "RectLinkLink", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_rectlinklink]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_rectlinklink]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_rectlinklink]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (C%d/U%d/D%d)", created, updated, deleted)
		case "SVG":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.SVG](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _svg := range set {
				nodeInstance := &tree.Node{Name: _svg.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_svg, "SVG", probe)

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
			nodeGongstruct.Name += fmt.Sprintf(" (C%d/U%d/D%d)", created, updated, deleted)
		case "SvgText":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.SvgText](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _svgtext := range set {
				nodeInstance := &tree.Node{Name: _svgtext.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_svgtext, "SvgText", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_svgtext]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_svgtext]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_svgtext]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (C%d/U%d/D%d)", created, updated, deleted)
		case "Text":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Text](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _text := range set {
				nodeInstance := &tree.Node{Name: _text.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_text, "Text", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_text]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_text]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_text]; ok {
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
