package node2gongdoc

import (
	"log"
	"os"
	"path/filepath"

	gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"
	gongtree_models "github.com/fullstack-lang/gongtree/go/models"
)

type NodeImplClasssiagram struct {

	// one need to access the daigramPackage to set the selected diagram
	diagramPackage *gongdoc_models.DiagramPackage
	classdiagram   *gongdoc_models.Classdiagram

	// one needs to access the node of the diagram package to manage the childern nodes
	diagramPackageNode *gongtree_models.Node

	// one needs to perform computation of node confs after the update
	treeOfGongObjects *gongtree_models.Tree

	// peculiar
	IsInDrawMode bool
}

func NewNodeImplClasssiagram(
	diagramPackage *gongdoc_models.DiagramPackage,
	classdiagram *gongdoc_models.Classdiagram,
	diagramPackageNode *gongtree_models.Node,
	treeOfGongObjects *gongtree_models.Tree,
) (nodeImplClasssiagram *NodeImplClasssiagram) {

	nodeImplClasssiagram = new(NodeImplClasssiagram)
	nodeImplClasssiagram.diagramPackage = diagramPackage
	nodeImplClasssiagram.classdiagram = classdiagram
	nodeImplClasssiagram.diagramPackageNode = diagramPackageNode
	nodeImplClasssiagram.treeOfGongObjects = treeOfGongObjects

	return
}

func (nodeImplClasssiagram *NodeImplClasssiagram) OnAfterUpdate(
	gongtreeStage *gongtree_models.StageStruct,
	stagedNode,
	frontNode *gongtree_models.Node) {

	gongdocStage := nodeImplClasssiagram.diagramPackage.Stage_

	log.Println("NodeImplClasssiagram: OnAfterUpdate")

	// node has been checked by the end user
	if frontNode.IsChecked && !stagedNode.IsChecked {

		// setting the value of the staged node	to the new value
		// and commit to the database,
		// The front will detect that the backend has been commited
		// It will refresh and fetch the node with checked value
		stagedNode.IsChecked = true
		stagedNode.Commit(gongtreeStage)
		nodeImplClasssiagram.diagramPackage.SelectedClassdiagram = nodeImplClasssiagram.classdiagram

		for _, otherDiagramNode := range nodeImplClasssiagram.diagramPackageNode.Children {
			if otherDiagramNode == stagedNode {
				continue
			}

			// uncheck the other node
			if otherDiagramNode.IsChecked {
				// log.Println("Node " + node.Name + " is checked and should be unchecked")
				otherDiagramNode.IsChecked = false
				otherDiagramNode.Commit(gongtreeStage)
			}
		}
	}

	// in case the front change the name of the diagram
	// one need to indicate the change as an increase in the commit
	// from the back, otherwise, the front wont be able to detect
	// the change
	// change the name of the diagram
	if stagedNode.Name != frontNode.Name {

		// check that the name is a correct identifer in go
		for _, b := range frontNode.Name {
			if 'a' <= b && b <= 'z' || 'A' <= b && b <= 'Z' || b == '_' || '0' <= b && b <= '9' {
				// Avoid assigning a rune for the common case of an ascii character.
				continue
			} else {
				log.Println("The name of the diagram is not a correct identifier in go: " + frontNode.Name)
				stagedNode.Commit(gongtreeStage)
				return
			}
		}

		// rename the diagram file if it exists
		// remove the actual classdiagram file if it exsits
		oldClassdiagramFilePath := filepath.Join(nodeImplClasssiagram.diagramPackage.Path,
			"../diagrams", nodeImplClasssiagram.classdiagram.Name) + ".go"
		newClassdiagramFilePath := filepath.Join(nodeImplClasssiagram.diagramPackage.Path,
			"../diagrams", frontNode.Name) + ".go"

		if _, err := os.Stat(oldClassdiagramFilePath); err != nil {
			return
		}
		if err := os.Remove(oldClassdiagramFilePath); err != nil {
			log.Fatal("Error while renaming file " + oldClassdiagramFilePath + " : " + err.Error())
		}

		file, err := os.Create(newClassdiagramFilePath)
		if err != nil {
			log.Fatal("Cannot create diagram file" + err.Error())
		}
		defer file.Close()

		// checkout in order to get the latest version of the diagram before
		// modifying it updated by the front
		gongdocStage.Checkout()
		gongdocStage.Unstage()
		gongdoc_models.StageBranch(gongdocStage, nodeImplClasssiagram.classdiagram)
		nodeImplClasssiagram.classdiagram.Name = frontNode.Name

		gongdoc_models.SetupMapDocLinkRenaming(nodeImplClasssiagram.diagramPackage.ModelPkg.Stage_, gongdocStage)

		// save the diagram
		gongdocStage.Marshall(file, "github.com/fullstack-lang/gongdoc/go/models", "diagrams")

		// restore the original stage
		gongdocStage.Unstage()
		gongdocStage.Checkout()

		nodeImplClasssiagram.classdiagram.Name = frontNode.Name
		gongdocStage.Commit()

		stagedNode.Name = frontNode.Name
		stagedNode.IsInEditMode = false

	}

	// node was checked and user wants to uncheck it. This is not possible
	// from a application logic point of view
	// on need to commit the staged node for the front to reconstruct
	// the node as checked and overides the unchecking action
	if stagedNode.IsChecked && !frontNode.IsChecked {
		stagedNode.Commit(gongtreeStage)
	}

	computeNodeConfs(
		gongtreeStage,
		gongdocStage,
		nodeImplClasssiagram.diagramPackageNode,
		nodeImplClasssiagram.diagramPackage,
		nodeImplClasssiagram.treeOfGongObjects)

	gongdocStage.Commit()
	gongtreeStage.Commit()
}
