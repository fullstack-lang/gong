package node2gongdoc

import (
	"fmt"
	"log"

	gong_models "github.com/fullstack-lang/gong/go/models"
	gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"
)

// NodeCB is the singloton callback implementation of CUD operations on node
type NodeCB struct {

	// diagramPackageNode is used for iterating on classdiagram nodes
	// and for getting the selected diagram
	diagramPackageNode *gongdoc_models.Node
	diagramPackage     *gongdoc_models.DiagramPackage

	// treeOfGongObjects is the root of all nodes related to gong objects
	// it is necessary for performing operation on all elements of the tree
	treeOfGongObjects *gongdoc_models.Tree

	// map_Children_Parent is a map to navigate from a children node to its parent node
	// it is set up once at the init phase
	map_Children_Parent map[*gongdoc_models.Node]*gongdoc_models.Node
}

// GetSelectedClassdiagram
func (nodesCb *NodeCB) GetSelectedClassdiagram() (classdiagram *gongdoc_models.Classdiagram) {

	classdiagram = nodesCb.diagramPackage.SelectedClassdiagram
	return
}

// OnAfterUpdate is called each time the end user interacts
// with any node. The front commit the state of the front node
// to the back ([frontNode] in the function).
// Therefore, the [stageNode] is now different from the [frontNode].
//
// This functiion and its siblings analyse which field of the
// node has changed and performs all necessary actions
func (nodeCb *NodeCB) OnAfterUpdate(
	stage *gongdoc_models.StageStruct,
	stagedNode, frontNode *gongdoc_models.Node) {

	if stagedNode.IsExpanded != frontNode.IsExpanded {
		// setting the value of the staged node	to the new value
		stagedNode.IsExpanded = frontNode.IsExpanded
		return
	}

	if stagedNode.Impl != nil {
		stagedNode.Impl.OnAfterUpdate(stage, stagedNode, frontNode)
		nodeCb.updateNodesStates(stage)
	}

}

// OnAfterCreate is another callback
func (nodeCb *NodeCB) OnAfterCreate(
	stage *gongdoc_models.StageStruct,
	node *gongdoc_models.Node) {

	log.Println("Node " + node.Name + " is created")

	node.HasCheckboxButton = true

	// check unicity of name, otherwise, add an index
	var hasNameCollision bool
	initialName := node.Name
	index := 0
	// loop until the name of the new diagram is not in collision with an existing
	// diagram name
	for index == 0 || hasNameCollision {
		index++
		hasNameCollision = false
		for classdiagram := range *gongdoc_models.GetGongstructInstancesSet[gongdoc_models.Classdiagram]() {
			if classdiagram.Name == node.Name {
				hasNameCollision = true
			}
		}
		if hasNameCollision {
			node.Name = initialName + fmt.Sprintf("_%d", index)
		}
	}

	classdiagram := (&gongdoc_models.Classdiagram{Name: node.Name}).Stage()
	nodeCb.diagramPackage.Classdiagrams = append(nodeCb.diagramPackage.Classdiagrams, classdiagram)
	node.IsInEditMode = false
	node.IsInDrawMode = false
	node.HasEditButton = false

	nodeCb.diagramPackageNode.Children = append(nodeCb.diagramPackageNode.Children, node)

	// set up the back pointer from the shape to the node
	classdiagramImpl := new(ClassdiagramImpl)
	classdiagramImpl.node = node
	classdiagramImpl.classdiagram = classdiagram
	classdiagramImpl.nodeCb = nodeCb
	node.Impl = classdiagramImpl

	nodeCb.updateNodesStates(stage)

}

// OnAfterDelete is called after a node is deleted
// notice that the fontNode only have its basic fields updated
// its pointers are not ok
func (nodeCb *NodeCB) OnAfterDelete(
	stage *gongdoc_models.StageStruct,
	stagedNode, frontNode *gongdoc_models.Node) {

	switch impl := stagedNode.Impl.(type) {
	case *ClassdiagramImpl:
		impl.OnAfterDelete(stage, stagedNode, frontNode)
	}

	nodeCb.updateNodesStates(stage)
}

func (nodeCb *NodeCB) FillUpDiagramNodeTree(diagramPackage *gongdoc_models.DiagramPackage) {

	// generate tree of diagrams
	gongdocTree := (&gongdoc_models.Tree{Name: "gongdoc"}).Stage()

	// add the root of class diagrams
	diagramPackageNode := (&gongdoc_models.Node{Name: "class diagrams"}).Stage()
	diagramPackageNode.IsExpanded = true
	diagramPackageNode.HasAddChildButton = diagramPackage.IsEditable
	gongdocTree.RootNodes = append(gongdocTree.RootNodes, diagramPackageNode)

	// add one node per class diagram
	for classdiagram := range *gongdoc_models.GetGongstructInstancesSet[gongdoc_models.Classdiagram]() {
		node := (&gongdoc_models.Node{Name: classdiagram.Name}).Stage()

		node.HasCheckboxButton = true
		node.HasDeleteButton = true
		node.HasEditButton = true
		diagramPackageNode.Children = append(diagramPackageNode.Children, node)

		// set up the back pointer from the shape to the node
		classdiagramImpl := new(ClassdiagramImpl)
		classdiagramImpl.node = node
		classdiagramImpl.classdiagram = classdiagram
		classdiagramImpl.nodeCb = nodeCb
		node.Impl = classdiagramImpl
	}

	// set callbacks on node updates
	nodeCb.diagramPackageNode = diagramPackageNode
}

func SetNodeBackPointer[T1 gong_models.Gongstruct](gong_instance *T1, backPointer gongdoc_models.NodeImplInterface) {
	gong_models.SetBackPointer(&gong_models.Stage, gong_instance, backPointer)
}
func GetNodeBackPointer[T1 gong_models.Gongstruct](gong_instance *T1) (backPointer gongdoc_models.NodeImplInterface) {
	tmp := gong_models.GetBackPointer(&gong_models.Stage, gong_instance)

	if tmp == nil {
		log.Fatal("backPointer is nil", gong_instance)
	}

	backPointer = tmp.(gongdoc_models.NodeImplInterface)

	return
}

func (nodeCb *NodeCB) FillUpTreeOfGongObjects() {

	// set up the gongTree to display elements
	gongTree := (&gongdoc_models.Tree{Name: "gong"}).Stage()
	nodeCb.treeOfGongObjects = gongTree

	gongstructRootNode := (&gongdoc_models.Node{Name: "gongstructs"}).Stage()
	gongstructRootNode.IsExpanded = true
	gongTree.RootNodes = append(gongTree.RootNodes, gongstructRootNode)
	for gongStruct := range *gong_models.GetGongstructInstancesSet[gong_models.GongStruct]() {

		nodeGongstruct := (&gongdoc_models.Node{Name: gongStruct.Name}).Stage()

		nodeGongstruct.HasCheckboxButton = true
		nodeGongstruct.IsExpanded = false

		// set up the back pointer from the shape to the node
		gongStructImpl := new(GongStructImpl)
		gongStructImpl.node = nodeGongstruct
		gongStructImpl.gongStruct = gongStruct
		gongStructImpl.nodeCb = nodeCb
		nodeGongstruct.Impl = gongStructImpl

		// append to the tree
		gongstructRootNode.Children = append(gongstructRootNode.Children, nodeGongstruct)

		SetNodeBackPointer(gongStruct, gongStructImpl)

		for _, field := range gongStruct.Fields {
			nodeGongField := (&gongdoc_models.Node{Name: field.GetName()}).Stage()

			nodeGongField.HasCheckboxButton = true

			fieldImpl := new(FieldImpl)
			fieldImpl.node = nodeGongstruct
			fieldImpl.field = field
			fieldImpl.nodeCb = nodeCb
			nodeGongField.Impl = fieldImpl

			// append to tree
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeGongField)

			switch fieldReal := field.(type) {
			case *gong_models.GongBasicField:
				SetNodeBackPointer(fieldReal, fieldImpl)
			case *gong_models.GongTimeField:
				SetNodeBackPointer(fieldReal, fieldImpl)
			case *gong_models.PointerToGongStructField:
				SetNodeBackPointer(fieldReal, fieldImpl)
			case *gong_models.SliceOfPointerToGongStructField:
				SetNodeBackPointer(fieldReal, fieldImpl)
			default:
			}
		}
	}

	gongenumRootNode := (&gongdoc_models.Node{Name: "gongenums"}).Stage()
	gongenumRootNode.IsExpanded = true
	gongTree.RootNodes = append(gongTree.RootNodes, gongenumRootNode)
	for gongEnum := range *gong_models.GetGongstructInstancesSet[gong_models.GongEnum]() {

		nodeGongEnum := (&gongdoc_models.Node{Name: gongEnum.Name}).Stage()
		nodeGongEnum.HasCheckboxButton = true
		nodeGongEnum.IsExpanded = false

		gongEnumImpl := new(GongEnumImpl)
		gongEnumImpl.node = nodeGongEnum
		gongEnumImpl.gongEnum = gongEnum
		gongEnumImpl.nodeCb = nodeCb
		nodeGongEnum.Impl = gongEnumImpl

		// append to tree
		gongenumRootNode.Children = append(gongenumRootNode.Children, nodeGongEnum)
		SetNodeBackPointer(gongEnum, gongEnumImpl)

		for _, gongEnumValue := range gongEnum.GongEnumValues {
			nodeGongEnumValue := (&gongdoc_models.Node{Name: gongEnumValue.GetName()}).Stage()

			nodeGongEnumValue.HasCheckboxButton = true

			gongEnumValueImpl := new(GongEnumValueImpl)
			gongEnumValueImpl.node = nodeGongEnum
			gongEnumValueImpl.gongEnumValue = gongEnumValue
			gongEnumValueImpl.nodeCb = nodeCb
			nodeGongEnumValue.Impl = gongEnumValueImpl

			// append to tree
			nodeGongEnum.Children = append(nodeGongEnum.Children, nodeGongEnumValue)
			SetNodeBackPointer(gongEnumValue, gongEnumValueImpl)
		}
	}

	gongNotesRootNode := (&gongdoc_models.Node{Name: "notes"}).Stage()
	gongNotesRootNode.IsExpanded = true
	gongTree.RootNodes = append(gongTree.RootNodes, gongNotesRootNode)
	for gongNote := range *gong_models.GetGongstructInstancesSet[gong_models.GongNote]() {

		node := (&gongdoc_models.Node{Name: gongNote.Name}).Stage()
		node.HasCheckboxButton = true

		node.IsExpanded = true

		gongNoteImpl := new(GongNoteImpl)
		gongNoteImpl.node = node
		gongNoteImpl.gongNote = gongNote
		gongNoteImpl.nodeCb = nodeCb
		node.Impl = gongNoteImpl

		// append to tree
		gongNotesRootNode.Children = append(gongNotesRootNode.Children, node)

		SetNodeBackPointer(gongNote, gongNoteImpl)

		for _, gongLink := range gongNote.Links {
			nodeGongLink := (&gongdoc_models.Node{Name: gongLink.Name}).Stage()
			nodeGongLink.HasCheckboxButton = true

			gongLinkImpl := new(GongLinkImpl)
			gongLinkImpl.node = node
			gongLinkImpl.gongLink = gongLink
			gongLinkImpl.nodeCb = nodeCb
			nodeGongLink.Impl = gongLinkImpl

			// append to tree
			node.Children = append(node.Children, nodeGongLink)

			SetNodeBackPointer(gongLink, gongLinkImpl)
		}
	}

	// generate the map to navigate from children to parents
	fieldName := gongdoc_models.GetAssociationName[gongdoc_models.Node]().Children[0].Name
	nodeCb.map_Children_Parent = gongdoc_models.GetSliceOfPointersReverseMap[gongdoc_models.Node, gongdoc_models.Node](fieldName)
}

func (nodeCb *NodeCB) updateNodesStates(stage *gongdoc_models.StageStruct) {

	nodeCb.updateDiagramsNodes(stage)

	// now manage object nodes accordign to the selected diagram

	// get the the selected diagram
	classdiagram := nodeCb.diagramPackage.SelectedClassdiagram

	for _, _node := range nodeCb.treeOfGongObjects.RootNodes {
		UncheckAndDisable(_node, classdiagram)
	}

	// no selected diagram yet
	if classdiagram == nil {
		gongdoc_models.Stage.Commit()
		return
	}

	nodeCb.updateGongObjectsNodes(stage, classdiagram)

	// log.Println("UpdateNodeStates, before commit, nb ", stage.BackRepo.GetLastCommitFromBackNb())
	stage.Commit()
	// log.Println("UpdateNodeStates, after  commit, nb ", stage.BackRepo.GetLastCommitFromBackNb())

}

func (nodesCb *NodeCB) updateDiagramsNodes(stage *gongdoc_models.StageStruct) {

	// compute wether one of the diagrams is in draw/edit mode
	// if so, all diagram check need to be disabled
	var inModificationMode bool
	for _, classdiagramNode := range nodesCb.diagramPackageNode.Children {
		if classdiagramNode.IsInDrawMode || classdiagramNode.IsInEditMode {
			inModificationMode = true
		}
	}

	// get the selected diagram and collect what are its referenced
	// gongstructs
	for _, classdiagramNode := range nodesCb.diagramPackageNode.Children {

		// reset the state of the classdiagram node
		classdiagramNode.HasEditButton = false
		classdiagramNode.HasDeleteButton = false
		classdiagramNode.HasDrawButton = false
		classdiagramNode.HasDrawOffButton = false

		classdiagramNode.IsCheckboxDisabled = inModificationMode

		if !classdiagramNode.IsChecked {
			classdiagramNode.IsInEditMode = false
			classdiagramNode.IsInDrawMode = false
			continue
		}

		editable := nodesCb.diagramPackage.IsEditable && !classdiagramNode.IsInEditMode && !classdiagramNode.IsInDrawMode

		classdiagramNode.HasEditButton = editable
		classdiagramNode.HasDeleteButton = editable
		classdiagramNode.HasDrawButton = editable
	}
}
