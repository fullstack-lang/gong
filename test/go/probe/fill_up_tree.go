package probe

import (
	"fmt"
	"sort"

	gongtree_buttons "github.com/fullstack-lang/gongtree/go/buttons"
	gongtree_models "github.com/fullstack-lang/gongtree/go/models"

	gong_models "github.com/fullstack-lang/gong/go/models"

	"github.com/fullstack-lang/gong/test/go/models"
)

func fillUpTree(
	playground *Playground,
) {
	playground.treeStage.Reset()

	// create tree
	sidebar := (&gongtree_models.Tree{Name: "gong"}).Stage(playground.treeStage)

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

		nodeGongstruct := (&gongtree_models.Node{Name: name}).Stage(playground.treeStage)


		nodeGongstruct.IsExpanded = false
		switch gongStruct.Name {
		// insertion point
		case "Astruct":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Astruct](playground.stageOfInterest)
			for astruct := range set {
				nodeInstance := (&gongtree_models.Node{Name: astruct.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(astruct, "Astruct", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "AstructBstruct2Use":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.AstructBstruct2Use](playground.stageOfInterest)
			for astructbstruct2use := range set {
				nodeInstance := (&gongtree_models.Node{Name: astructbstruct2use.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(astructbstruct2use, "AstructBstruct2Use", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "AstructBstructUse":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.AstructBstructUse](playground.stageOfInterest)
			for astructbstructuse := range set {
				nodeInstance := (&gongtree_models.Node{Name: astructbstructuse.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(astructbstructuse, "AstructBstructUse", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Bstruct":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Bstruct](playground.stageOfInterest)
			for bstruct := range set {
				nodeInstance := (&gongtree_models.Node{Name: bstruct.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(bstruct, "Bstruct", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Dstruct":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Dstruct](playground.stageOfInterest)
			for dstruct := range set {
				nodeInstance := (&gongtree_models.Node{Name: dstruct.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(dstruct, "Dstruct", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}	
		}

		nodeGongstruct.IsNodeClickable = true
		nodeGongstruct.Impl = NewTreeNodeImplGongstruct(gongStruct, playground)

		// add add button
		addButton := (&gongtree_models.Button{
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
	gongtreeStage *gongtree_models.StageStruct,
	stagedNode, frontNode *gongtree_models.Node) {

	FillUpFormFromGongstruct(
		instanceNodeCallback.Instance,
		instanceNodeCallback.playground,
	)
}
