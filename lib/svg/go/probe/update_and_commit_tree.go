package probe

import (
	"fmt"
	"sort"
	"strings"

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

	topNode := &tree.Node{Name: fmt.Sprintf("Stage %s", probe.stageOfInterest.GetName())}

	sidebar.RootNodes = append(sidebar.RootNodes, topNode)
	refreshButton := &tree.Button{
		Name:            "RefreshButton" + " " + string(gongtree_buttons.BUTTON_refresh),
		Icon:            string(gongtree_buttons.BUTTON_refresh),
		HasToolTip:      true,
		ToolTipText:     "Refresh probe",
		ToolTipPosition: tree.Left,
	}
	topNode.Buttons = append(topNode.Buttons, refreshButton)
	refreshButton.Impl = &tree.FunctionalButtonProxy{
		OnUpdated: func(stage *tree.Stage,
			stagedButton, frontButton *tree.Button) {
			probe.stageOfInterest.ComputeInstancesNb()
			probe.docStager.SetMap_GongStructName_InstancesNb(
				probe.stageOfInterest.Map_GongStructName_InstancesNb,
			)
			probe.Refresh()
		},
	}

	notificationsResetButton := &tree.Button{
		Name:            "NotificationsResetButton",
		Icon:            string(gongtree_buttons.BUTTON_reset_tv),
		HasToolTip:      true,
		ToolTipText:     "Reset notification table",
		ToolTipPosition: tree.Left,
	}
	topNode.Buttons = append(topNode.Buttons, notificationsResetButton)
	notificationsResetButton.Impl = &tree.FunctionalButtonProxy{
		OnUpdated: func(stage *tree.Stage,
			stagedButton, frontButton *tree.Button) {
			probe.ResetNotifications()
		},
	}

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
			for _animate := range set {
				nodeInstance := &tree.Node{Name: _animate.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_animate, "Animate", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Circle":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Circle](probe.stageOfInterest)
			for _circle := range set {
				nodeInstance := &tree.Node{Name: _circle.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_circle, "Circle", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Condition":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Condition](probe.stageOfInterest)
			for _condition := range set {
				nodeInstance := &tree.Node{Name: _condition.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_condition, "Condition", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "ControlPoint":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.ControlPoint](probe.stageOfInterest)
			for _controlpoint := range set {
				nodeInstance := &tree.Node{Name: _controlpoint.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_controlpoint, "ControlPoint", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Ellipse":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Ellipse](probe.stageOfInterest)
			for _ellipse := range set {
				nodeInstance := &tree.Node{Name: _ellipse.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_ellipse, "Ellipse", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Layer":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Layer](probe.stageOfInterest)
			for _layer := range set {
				nodeInstance := &tree.Node{Name: _layer.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_layer, "Layer", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Line":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Line](probe.stageOfInterest)
			for _line := range set {
				nodeInstance := &tree.Node{Name: _line.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_line, "Line", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Link":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Link](probe.stageOfInterest)
			for _link := range set {
				nodeInstance := &tree.Node{Name: _link.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_link, "Link", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "LinkAnchoredText":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.LinkAnchoredText](probe.stageOfInterest)
			for _linkanchoredtext := range set {
				nodeInstance := &tree.Node{Name: _linkanchoredtext.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_linkanchoredtext, "LinkAnchoredText", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Path":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Path](probe.stageOfInterest)
			for _path := range set {
				nodeInstance := &tree.Node{Name: _path.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_path, "Path", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Point":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Point](probe.stageOfInterest)
			for _point := range set {
				nodeInstance := &tree.Node{Name: _point.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_point, "Point", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Polygone":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Polygone](probe.stageOfInterest)
			for _polygone := range set {
				nodeInstance := &tree.Node{Name: _polygone.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_polygone, "Polygone", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Polyline":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Polyline](probe.stageOfInterest)
			for _polyline := range set {
				nodeInstance := &tree.Node{Name: _polyline.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_polyline, "Polyline", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Rect":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Rect](probe.stageOfInterest)
			for _rect := range set {
				nodeInstance := &tree.Node{Name: _rect.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_rect, "Rect", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "RectAnchoredPath":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.RectAnchoredPath](probe.stageOfInterest)
			for _rectanchoredpath := range set {
				nodeInstance := &tree.Node{Name: _rectanchoredpath.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_rectanchoredpath, "RectAnchoredPath", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "RectAnchoredRect":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.RectAnchoredRect](probe.stageOfInterest)
			for _rectanchoredrect := range set {
				nodeInstance := &tree.Node{Name: _rectanchoredrect.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_rectanchoredrect, "RectAnchoredRect", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "RectAnchoredText":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.RectAnchoredText](probe.stageOfInterest)
			for _rectanchoredtext := range set {
				nodeInstance := &tree.Node{Name: _rectanchoredtext.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_rectanchoredtext, "RectAnchoredText", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "RectLinkLink":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.RectLinkLink](probe.stageOfInterest)
			for _rectlinklink := range set {
				nodeInstance := &tree.Node{Name: _rectlinklink.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_rectlinklink, "RectLinkLink", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "SVG":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.SVG](probe.stageOfInterest)
			for _svg := range set {
				nodeInstance := &tree.Node{Name: _svg.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_svg, "SVG", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "SvgText":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.SvgText](probe.stageOfInterest)
			for _svgtext := range set {
				nodeInstance := &tree.Node{Name: _svgtext.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_svgtext, "SvgText", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Text":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Text](probe.stageOfInterest)
			for _text := range set {
				nodeInstance := &tree.Node{Name: _text.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_text, "Text", probe)

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
		addButton.Impl = &tree.FunctionalButtonProxy{
			OnUpdated: func(stage *tree.Stage, stagedButton, frontButton *tree.Button) {
				FillUpFormFromGongstructName(
					probe,
					gongStruct.Name,
					true,
				)
			},
		}

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
