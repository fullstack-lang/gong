package node2gongdoc

import (
	"log"
	"os"
	"path/filepath"

	gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"
)

type ClassdiagramImpl struct {
	classdiagram *gongdoc_models.Classdiagram
	NodeImpl
}

func (classdiagramImpl *ClassdiagramImpl) OnAfterUpdate(
	gongdocStage *gongdoc_models.StageStruct,
	stagedNode, frontNode *gongdoc_models.Node) {

	// node has been checked by the end user
	if frontNode.IsChecked && !stagedNode.IsChecked {

		// setting the value of the staged node	to the new value
		// and commit to the database,
		// The front will detect that the backend has been commited
		// It will refresh and fetch the node with checked value
		stagedNode.IsChecked = true
		classdiagramImpl.nodeCb.diagramPackage.SelectedClassdiagram = classdiagramImpl.classdiagram

		for _, otherDiagramNode := range classdiagramImpl.nodeCb.diagramPackageNode.Children {
			if otherDiagramNode == stagedNode {
				continue
			}

			// uncheck the other node
			if otherDiagramNode.IsChecked {
				// log.Println("Node " + node.Name + " is checked and should be unchecked")
				otherDiagramNode.IsChecked = false
			}
		}
	}

	// node was checked and user wants to uncheck it. This is not possible
	// from a application logic point of view
	// on need to commit the staged node for the front to reconstruct
	// the node as checked and overides the unchecking action
	if stagedNode.IsChecked && !frontNode.IsChecked {
		stagedNode.Commit(gongdocStage)
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
				stagedNode.Commit(gongdocStage)
				return
			}
		}

		// rename the diagram file if it exists
		// remove the actual classdiagram file if it exsits
		oldClassdiagramFilePath := filepath.Join(classdiagramImpl.nodeCb.diagramPackage.Path, "../diagrams", classdiagramImpl.classdiagram.Name) + ".go"
		newClassdiagramFilePath := filepath.Join(classdiagramImpl.nodeCb.diagramPackage.Path, "../diagrams", frontNode.Name) + ".go"

		if _, err := os.Stat(oldClassdiagramFilePath); err == nil {
			if err := os.Remove(oldClassdiagramFilePath); err != nil {
				log.Println("Error while renaming file " + oldClassdiagramFilePath + " : " + err.Error())
			} else {
				file, err := os.Create(newClassdiagramFilePath)
				if err != nil {
					log.Fatal("Cannot open diagram file" + err.Error())
				}
				defer file.Close()

				classdiagramImpl.classdiagram.Name = frontNode.Name

				// checkout in order to get the latest version of the diagram before
				// modifying it updated by the front
				gongdocStage.Checkout()
				gongdocStage.Unstage()
				gongdoc_models.StageBranch(gongdocStage, classdiagramImpl.classdiagram)

				gongdoc_models.SetupMapDocLinkRenamingNew(gongdocStage, classdiagramImpl.nodeCb.diagramPackage)

				// save the diagram
				gongdocStage.Marshall(file, "github.com/fullstack-lang/gongdoc/go/models", "diagrams")

				// restore the original stage
				gongdocStage.Unstage()
				gongdocStage.Checkout()
			}
		}

		stagedNode.Name = frontNode.Name
		stagedNode.IsInEditMode = false

	}

	// the end user switch the edit mode
	if stagedNode.IsInEditMode != frontNode.IsInEditMode {
		stagedNode.IsInEditMode = frontNode.IsInEditMode

	}

	if stagedNode.IsInDrawMode != frontNode.IsInDrawMode {
		stagedNode.IsInDrawMode = frontNode.IsInDrawMode

		classdiagramImpl.classdiagram.IsInDrawMode = frontNode.IsInDrawMode

	}

	// marshall diagram to switch to saved state
	if !stagedNode.IsSaved && frontNode.IsSaved {

		// checkout in order to get the latest version of the diagram before
		// modifying it updated by the front
		gongdocStage.Checkout()
		gongdocStage.Unstage()
		gongdoc_models.StageBranch(gongdocStage, classdiagramImpl.classdiagram)

		filepath := filepath.Join(
			filepath.Join(classdiagramImpl.nodeCb.diagramPackage.AbsolutePathToDiagramPackage,
				"../diagrams"),
			classdiagramImpl.classdiagram.Name) + ".go"
		file, err := os.Create(filepath)
		if err != nil {
			log.Fatal("Cannot open diagram file" + err.Error())
		}
		defer file.Close()

		mapDocLinkRemaping := &gongdocStage.Map_DocLink_Renaming
		_ = mapDocLinkRemaping

		gongdoc_models.SetupMapDocLinkRenamingNew(gongdocStage, classdiagramImpl.nodeCb.diagramPackage)

		gongdocStage.Marshall(file, "github.com/fullstack-lang/gongdoc/go/models", "diagrams")

		// restore the original stage
		gongdocStage.Unstage()
		gongdocStage.Checkout()
		stagedNode.IsSaved = false
		gongdocStage.Commit()

	}
}

func (classdiagramImpl *ClassdiagramImpl) OnAfterDelete(
	stage *gongdoc_models.StageStruct,
	stagedNode, frontNode *gongdoc_models.Node) {

	// checkout the stage, it shall remove the link between
	// the parent node and the staged node because 0..1->0..N association
	// is stored in the staged node as a reverse pointer
	stage.Checkout()

	// remove the classdiagram node from the pkg element node
	classdiagramImpl.nodeCb.diagramPackage.Classdiagrams =
		remove(classdiagramImpl.nodeCb.diagramPackage.Classdiagrams, classdiagramImpl.classdiagram)
	gongdoc_models.UnstageBranch(stage, classdiagramImpl.classdiagram)

	// remove the actual classdiagram file if it exsits
	classdiagramFilePath := filepath.Join(classdiagramImpl.nodeCb.diagramPackage.Path, "../diagrams", classdiagramImpl.classdiagram.Name) + ".go"
	if _, err := os.Stat(classdiagramFilePath); err == nil {
		if err := os.Remove(classdiagramFilePath); err != nil {
			log.Println("Error while deleting file " + classdiagramFilePath + " : " + err.Error())
		}
	}
}

// remove node from slice
func remove[T gongdoc_models.Gongstruct](slice []*T, t *T) []*T {

	// get the index of the t
	rank := -1
	for i, t_ := range slice {
		if t_ == t {
			rank = i
		}
	}
	return append(slice[:rank], slice[rank+1:]...)
}
