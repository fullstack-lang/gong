package node2gongdoc

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

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
		nodeCb.computeNodesConfiguration(stage)
	}

}

// OnAfterCreate is another callback
func (nodeCb *NodeCB) OnAfterCreate(
	gongdocStage *gongdoc_models.StageStruct,
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

	classdiagram := (&gongdoc_models.Classdiagram{Name: node.Name}).Stage(nodeCb.diagramPackage.Stage_)
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

	filepath := filepath.Join(
		filepath.Join(classdiagramImpl.nodeCb.diagramPackage.AbsolutePathToDiagramPackage,
			"../diagrams"),
		classdiagramImpl.classdiagram.Name) + ".go"
	file, err := os.Create(filepath)
	if err != nil {
		log.Fatal("Cannot open diagram file" + err.Error())
	}
	defer file.Close()

	// save the diagram
	// checkout in order to get the latest version of the diagram before
	// modifying it updated by the front
	nodeCb.computeNodesConfiguration(gongdocStage)
	gongdocStage.Commit()

	// now save the diagram
	gongdocStage.Checkout()
	gongdocStage.Unstage()
	gongdoc_models.StageBranch(gongdocStage, classdiagramImpl.classdiagram)

	gongdoc_models.SetupMapDocLinkRenamingNew(gongdocStage, nodeCb.diagramPackage)

	gongdocStage.Marshall(file, "github.com/fullstack-lang/gongdoc/go/models", "diagrams")

	// restore the original stage
	gongdocStage.Unstage()
	gongdocStage.Checkout()

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

	nodeCb.computeNodesConfiguration(stage)
}

func (nodeCb *NodeCB) FillUpDiagramNodeTree(diagramPackage *gongdoc_models.DiagramPackage) {

	// generate tree of diagrams
	gongdocTree := (&gongdoc_models.Tree{Name: "gongdoc"}).Stage(nodeCb.diagramPackage.Stage_)

	// add the root of class diagrams
	diagramPackageNode := (&gongdoc_models.Node{Name: "class diagrams"}).Stage(nodeCb.diagramPackage.Stage_)
	diagramPackageNode.IsExpanded = true
	gongdocTree.RootNodes = append(gongdocTree.RootNodes, diagramPackageNode)

	// add one node per class diagram
	for classdiagram := range *gongdoc_models.GetGongstructInstancesSet[gongdoc_models.Classdiagram]() {
		node := (&gongdoc_models.Node{Name: classdiagram.Name}).Stage(nodeCb.diagramPackage.Stage_)

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

// computeNodesConfiguration computes both trees
func (nodeCb *NodeCB) computeNodesConfiguration(gongdocStage *gongdoc_models.StageStruct) {

	nodeCb.computeDiagramNodesConfigurations(gongdocStage)

	// now manage object nodes accordign to the selected diagram

	// get the the selected diagram
	classdiagram := nodeCb.diagramPackage.SelectedClassdiagram

	// if no diagram is selected, all gong nodes are to be disabled
	isCheckboxDisabled := true

	// is the classdiagram is not in drawing mode, all gong nodes are to be disabled
	// otherwise gong nodes are enabled by default
	if classdiagram != nil {
		isCheckboxDisabled = !classdiagram.IsInDrawMode
	}

	// now, compute wether each gong node to be checked / disabled
	for _, _node := range nodeCb.treeOfGongObjects.RootNodes {
		applyGongNodesConfiguration(_node, isCheckboxDisabled, false)
	}

	// no selected diagram yet
	if classdiagram == nil {
		gongdocStage.Commit()
		return
	}

	nodeCb.computeGongNodesConfiguration(gongdocStage, classdiagram)

	// log.Println("UpdateNodeStates, before commit, nb ", stage.BackRepo.GetLastCommitFromBackNb())
	gongdocStage.Commit()
	// log.Println("UpdateNodeStates, after  commit, nb ", stage.BackRepo.GetLastCommitFromBackNb())

}

func (nodesCb *NodeCB) computeDiagramNodesConfigurations(stage *gongdoc_models.StageStruct) {

	// compute wether one of the diagrams is in draw/edit mode
	// if so, all diagram check need to be disabled
	var inModificationMode bool
	for _, classdiagramNode := range nodesCb.diagramPackageNode.Children {
		if classdiagramNode.IsInDrawMode || classdiagramNode.IsInEditMode {
			inModificationMode = true
		}
	}

	nodesCb.diagramPackageNode.HasAddChildButton = !inModificationMode && nodesCb.diagramPackage.IsEditable

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
