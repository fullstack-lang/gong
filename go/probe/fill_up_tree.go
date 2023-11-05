package probe

import (
	"fmt"
	"sort"
	"strings"

	gongtree_buttons "github.com/fullstack-lang/gongtree/go/buttons"
	tree "github.com/fullstack-lang/gongtree/go/models"

	gong_models "github.com/fullstack-lang/gong/go/models"

	"github.com/fullstack-lang/gong/go/models"
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
		case "GongBasicField":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.GongBasicField](probe.stageOfInterest)
			for _gongbasicfield := range set {
				nodeInstance := (&tree.Node{Name: _gongbasicfield.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_gongbasicfield, "GongBasicField", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "GongEnum":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.GongEnum](probe.stageOfInterest)
			for _gongenum := range set {
				nodeInstance := (&tree.Node{Name: _gongenum.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_gongenum, "GongEnum", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "GongEnumValue":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.GongEnumValue](probe.stageOfInterest)
			for _gongenumvalue := range set {
				nodeInstance := (&tree.Node{Name: _gongenumvalue.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_gongenumvalue, "GongEnumValue", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "GongLink":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.GongLink](probe.stageOfInterest)
			for _gonglink := range set {
				nodeInstance := (&tree.Node{Name: _gonglink.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_gonglink, "GongLink", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "GongNote":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.GongNote](probe.stageOfInterest)
			for _gongnote := range set {
				nodeInstance := (&tree.Node{Name: _gongnote.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_gongnote, "GongNote", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "GongStruct":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.GongStruct](probe.stageOfInterest)
			for _gongstruct := range set {
				nodeInstance := (&tree.Node{Name: _gongstruct.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_gongstruct, "GongStruct", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "GongTimeField":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.GongTimeField](probe.stageOfInterest)
			for _gongtimefield := range set {
				nodeInstance := (&tree.Node{Name: _gongtimefield.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_gongtimefield, "GongTimeField", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Meta":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Meta](probe.stageOfInterest)
			for _meta := range set {
				nodeInstance := (&tree.Node{Name: _meta.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_meta, "Meta", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "MetaReference":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.MetaReference](probe.stageOfInterest)
			for _metareference := range set {
				nodeInstance := (&tree.Node{Name: _metareference.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_metareference, "MetaReference", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "ModelPkg":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.ModelPkg](probe.stageOfInterest)
			for _modelpkg := range set {
				nodeInstance := (&tree.Node{Name: _modelpkg.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_modelpkg, "ModelPkg", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "PointerToGongStructField":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.PointerToGongStructField](probe.stageOfInterest)
			for _pointertogongstructfield := range set {
				nodeInstance := (&tree.Node{Name: _pointertogongstructfield.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_pointertogongstructfield, "PointerToGongStructField", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "SliceOfPointerToGongStructField":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.SliceOfPointerToGongStructField](probe.stageOfInterest)
			for _sliceofpointertogongstructfield := range set {
				nodeInstance := (&tree.Node{Name: _sliceofpointertogongstructfield.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_sliceofpointertogongstructfield, "SliceOfPointerToGongStructField", probe)

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
