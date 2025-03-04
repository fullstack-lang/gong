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

func fillUpTree(
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
	sidebar := (&tree.Tree{Name: "gong"}).Stage(probe.treeStage)

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

		nodeGongstruct := (&tree.Node{Name: name}).Stage(probe.treeStage)

		nodeGongstruct.IsExpanded = false
		if _, ok := expandedNodesSet[strings.Fields(name)[0]]; ok {
			nodeGongstruct.IsExpanded = true
		}

		switch gongStruct.Name {
		// insertion point
		case "Animate":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Animate](probe.stageOfInterest)
			for _animate := range set {
				nodeInstance := (&tree.Node{Name: _animate.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_animate, "Animate", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Circle":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Circle](probe.stageOfInterest)
			for _circle := range set {
				nodeInstance := (&tree.Node{Name: _circle.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_circle, "Circle", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Ellipse":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Ellipse](probe.stageOfInterest)
			for _ellipse := range set {
				nodeInstance := (&tree.Node{Name: _ellipse.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_ellipse, "Ellipse", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Layer":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Layer](probe.stageOfInterest)
			for _layer := range set {
				nodeInstance := (&tree.Node{Name: _layer.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_layer, "Layer", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Line":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Line](probe.stageOfInterest)
			for _line := range set {
				nodeInstance := (&tree.Node{Name: _line.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_line, "Line", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Link":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Link](probe.stageOfInterest)
			for _link := range set {
				nodeInstance := (&tree.Node{Name: _link.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_link, "Link", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "LinkAnchoredText":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.LinkAnchoredText](probe.stageOfInterest)
			for _linkanchoredtext := range set {
				nodeInstance := (&tree.Node{Name: _linkanchoredtext.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_linkanchoredtext, "LinkAnchoredText", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Path":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Path](probe.stageOfInterest)
			for _path := range set {
				nodeInstance := (&tree.Node{Name: _path.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_path, "Path", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Point":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Point](probe.stageOfInterest)
			for _point := range set {
				nodeInstance := (&tree.Node{Name: _point.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_point, "Point", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Polygone":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Polygone](probe.stageOfInterest)
			for _polygone := range set {
				nodeInstance := (&tree.Node{Name: _polygone.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_polygone, "Polygone", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Polyline":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Polyline](probe.stageOfInterest)
			for _polyline := range set {
				nodeInstance := (&tree.Node{Name: _polyline.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_polyline, "Polyline", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Rect":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Rect](probe.stageOfInterest)
			for _rect := range set {
				nodeInstance := (&tree.Node{Name: _rect.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_rect, "Rect", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "RectAnchoredPath":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.RectAnchoredPath](probe.stageOfInterest)
			for _rectanchoredpath := range set {
				nodeInstance := (&tree.Node{Name: _rectanchoredpath.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_rectanchoredpath, "RectAnchoredPath", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "RectAnchoredRect":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.RectAnchoredRect](probe.stageOfInterest)
			for _rectanchoredrect := range set {
				nodeInstance := (&tree.Node{Name: _rectanchoredrect.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_rectanchoredrect, "RectAnchoredRect", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "RectAnchoredText":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.RectAnchoredText](probe.stageOfInterest)
			for _rectanchoredtext := range set {
				nodeInstance := (&tree.Node{Name: _rectanchoredtext.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_rectanchoredtext, "RectAnchoredText", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "RectLinkLink":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.RectLinkLink](probe.stageOfInterest)
			for _rectlinklink := range set {
				nodeInstance := (&tree.Node{Name: _rectlinklink.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_rectlinklink, "RectLinkLink", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "SVG":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.SVG](probe.stageOfInterest)
			for _svg := range set {
				nodeInstance := (&tree.Node{Name: _svg.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_svg, "SVG", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "SvgText":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.SvgText](probe.stageOfInterest)
			for _svgtext := range set {
				nodeInstance := (&tree.Node{Name: _svgtext.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_svgtext, "SvgText", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Text":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Text](probe.stageOfInterest)
			for _text := range set {
				nodeInstance := (&tree.Node{Name: _text.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_text, "Text", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		}

		nodeGongstruct.IsNodeClickable = true
		nodeGongstruct.Impl = NewTreeNodeImplGongstruct(gongStruct, probe)

		// add add button
		addButton := (&tree.Button{
			Name: gongStruct.Name + " " + string(gongtree_buttons.BUTTON_add),
			Icon: string(gongtree_buttons.BUTTON_add)}).Stage(probe.treeStage)
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)
		addButton.Impl = NewButtonImplGongstruct(
			gongStruct,
			gongtree_buttons.BUTTON_add,
			probe,
		)

		sidebar.RootNodes = append(sidebar.RootNodes, nodeGongstruct)
	}

	// Add a refresh button
	nodeRefreshButton := (&tree.Node{Name: ""}).Stage(probe.treeStage)
	sidebar.RootNodes = append(sidebar.RootNodes, nodeRefreshButton)
	refreshButton := (&tree.Button{
		Name: "RefreshButton" + " " + string(gongtree_buttons.BUTTON_refresh),
		Icon: string(gongtree_buttons.BUTTON_refresh)}).Stage(probe.treeStage)
	nodeRefreshButton.Buttons = append(nodeRefreshButton.Buttons, refreshButton)
	refreshButton.Impl = NewButtonImplRefresh(probe)

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
	gongtreeStage *tree.StageStruct,
	stagedNode, frontNode *tree.Node) {

	FillUpFormFromGongstruct(
		instanceNodeCallback.Instance,
		instanceNodeCallback.probe,
	)
}
