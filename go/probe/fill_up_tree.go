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
	playground *Playground,
) {
	// keep in memory which nodes have been unfolded / folded
	expandedNodesSet := make(map[string]any, 0)
	var _sidebar *tree.Tree
	for __sidebar := range playground.treeStage.Trees {
		_sidebar = __sidebar
	}
	if _sidebar != nil {
		for _, node := range _sidebar.RootNodes {
			if node.IsExpanded {
				expandedNodesSet[strings.Fields(node.Name)[0]] = true
			}
		}
	}

	playground.treeStage.Reset()

	// create tree
	sidebar := (&tree.Tree{Name: "gong"}).Stage(playground.treeStage)

	// collect all gong struct to construe the true
	setOfGongStructs := *gong_models.GetGongstructInstancesSet[gong_models.GongStruct](playground.gongStage)

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
			fmt.Sprintf("%d", playground.stageOfInterest.Map_GongStructName_InstancesNb[gongStruct.Name]) + ")"

		nodeGongstruct := (&tree.Node{Name: name}).Stage(playground.treeStage)


		nodeGongstruct.IsExpanded = false
		if _, ok := expandedNodesSet[strings.Fields(name)[0]]; ok {
			nodeGongstruct.IsExpanded = true
		}
		
		switch gongStruct.Name {
		// insertion point
		case "GongBasicField":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.GongBasicField](playground.stageOfInterest)
			for _gongbasicfield := range set {
				nodeInstance := (&tree.Node{Name: _gongbasicfield.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_gongbasicfield, "GongBasicField", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "GongEnum":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.GongEnum](playground.stageOfInterest)
			for _gongenum := range set {
				nodeInstance := (&tree.Node{Name: _gongenum.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_gongenum, "GongEnum", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "GongEnumValue":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.GongEnumValue](playground.stageOfInterest)
			for _gongenumvalue := range set {
				nodeInstance := (&tree.Node{Name: _gongenumvalue.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_gongenumvalue, "GongEnumValue", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "GongLink":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.GongLink](playground.stageOfInterest)
			for _gonglink := range set {
				nodeInstance := (&tree.Node{Name: _gonglink.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_gonglink, "GongLink", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "GongNote":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.GongNote](playground.stageOfInterest)
			for _gongnote := range set {
				nodeInstance := (&tree.Node{Name: _gongnote.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_gongnote, "GongNote", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "GongStruct":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.GongStruct](playground.stageOfInterest)
			for _gongstruct := range set {
				nodeInstance := (&tree.Node{Name: _gongstruct.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_gongstruct, "GongStruct", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "GongTimeField":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.GongTimeField](playground.stageOfInterest)
			for _gongtimefield := range set {
				nodeInstance := (&tree.Node{Name: _gongtimefield.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_gongtimefield, "GongTimeField", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Meta":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Meta](playground.stageOfInterest)
			for _meta := range set {
				nodeInstance := (&tree.Node{Name: _meta.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_meta, "Meta", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "MetaReference":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.MetaReference](playground.stageOfInterest)
			for _metareference := range set {
				nodeInstance := (&tree.Node{Name: _metareference.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_metareference, "MetaReference", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "ModelPkg":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.ModelPkg](playground.stageOfInterest)
			for _modelpkg := range set {
				nodeInstance := (&tree.Node{Name: _modelpkg.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_modelpkg, "ModelPkg", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "PointerToGongStructField":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.PointerToGongStructField](playground.stageOfInterest)
			for _pointertogongstructfield := range set {
				nodeInstance := (&tree.Node{Name: _pointertogongstructfield.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_pointertogongstructfield, "PointerToGongStructField", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "SliceOfPointerToGongStructField":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.SliceOfPointerToGongStructField](playground.stageOfInterest)
			for _sliceofpointertogongstructfield := range set {
				nodeInstance := (&tree.Node{Name: _sliceofpointertogongstructfield.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_sliceofpointertogongstructfield, "SliceOfPointerToGongStructField", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}	
		}

		nodeGongstruct.IsNodeClickable = true
		nodeGongstruct.Impl = NewTreeNodeImplGongstruct(gongStruct, playground)

		// add add button
		addButton := (&tree.Button{
			Name: gongStruct.Name + " " + string(gongtree_buttons.BUTTON_add),
			Icon: string(gongtree_buttons.BUTTON_add)}).Stage(playground.treeStage)
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)
		addButton.Impl = NewButtonImplGongstruct(
			gongStruct,
			gongtree_buttons.BUTTON_add,
			playground,
		)

		sidebar.RootNodes = append(sidebar.RootNodes, nodeGongstruct)
	}
	playground.treeStage.Commit()
}

type InstanceNodeCallback[T models.Gongstruct] struct {
	Instance       *T
	gongstructName string
	playground     *Playground
}

func NewInstanceNodeCallback[T models.Gongstruct](
	instance *T,
	gongstructName string,
	playground *Playground) (
	instanceNodeCallback *InstanceNodeCallback[T],
) {
	instanceNodeCallback = new(InstanceNodeCallback[T])

	instanceNodeCallback.playground = playground
	instanceNodeCallback.gongstructName = gongstructName
	instanceNodeCallback.Instance = instance

	return
}

func (instanceNodeCallback *InstanceNodeCallback[T]) OnAfterUpdate(
	gongtreeStage *tree.StageStruct,
	stagedNode, frontNode *tree.Node) {

	FillUpFormFromGongstruct(
		instanceNodeCallback.Instance,
		instanceNodeCallback.playground,
	)
}
