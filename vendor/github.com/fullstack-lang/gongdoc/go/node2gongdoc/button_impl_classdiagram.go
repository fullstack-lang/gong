package node2gongdoc

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"
	gongtree_models "github.com/fullstack-lang/gongtree/go/models"
)

type ButtonImplClassdiagram struct {
	// one need to access the diagramPackage to set the selected diagram
	diagramPackage *gongdoc_models.DiagramPackage
	classdiagram   *gongdoc_models.Classdiagram

	// one needs to access the node of the diagram package to manage the childern nodes
	diagramPackageNode *gongtree_models.Node

	// one needs to perform computation of node confs after the update
	treeOfGongObjects *gongtree_models.Tree

	classdiagramNode     *gongtree_models.Node
	nodeImplClassdiagram *NodeImplClasssiagram

	// type of button
	Icon ButtonType
}

func NewButtonImplClassdiagram(
	diagramPackage *gongdoc_models.DiagramPackage,
	classdiagram *gongdoc_models.Classdiagram,
	diagramPackageNode *gongtree_models.Node,
	treeOfGongObjects *gongtree_models.Tree,
	classdiagramNode *gongtree_models.Node,
	nodeImplClassdiagram *NodeImplClasssiagram,
	icon ButtonType,
) (buttonImplClassdiagram *ButtonImplClassdiagram) {

	buttonImplClassdiagram = new(ButtonImplClassdiagram)

	buttonImplClassdiagram.diagramPackage = diagramPackage
	buttonImplClassdiagram.classdiagram = classdiagram
	buttonImplClassdiagram.diagramPackageNode = diagramPackageNode
	buttonImplClassdiagram.treeOfGongObjects = treeOfGongObjects
	buttonImplClassdiagram.classdiagramNode = classdiagramNode
	buttonImplClassdiagram.nodeImplClassdiagram = nodeImplClassdiagram
	buttonImplClassdiagram.Icon = icon

	return
}

func (buttonImplClassdiagram *ButtonImplClassdiagram) ButtonUpdated(
	gongtreeStage *gongtree_models.StageStruct,
	stageButton, front *gongtree_models.Button) {

	gongdocStage := buttonImplClassdiagram.diagramPackage.Stage_

	log.Println("ButtonImplClassdiagramDraw, ButtonUpdated", front.Name)

	switch buttonImplClassdiagram.Icon {
	case BUTTON_draw:
		buttonImplClassdiagram.classdiagram.IsInDrawMode = true
		buttonImplClassdiagram.nodeImplClassdiagram.IsInDrawMode = true
	case BUTTON_edit_off:

		// edit of the drawing of the diagram
		if buttonImplClassdiagram.nodeImplClassdiagram.IsInDrawMode {

			selectedClassdiagramName := buttonImplClassdiagram.diagramPackage.SelectedClassdiagram.Name
			// reset the stage
			gongdocStage.Checkout()

			// reload
			fakeFrontDiagramPackage := (&gongdoc_models.DiagramPackage{})
			fakeFrontDiagramPackage.IsReloaded = buttonImplClassdiagram.diagramPackage.IsReloaded
			buttonImplClassdiagram.diagramPackage.IsReloaded = !buttonImplClassdiagram.diagramPackage.IsReloaded

			gongdocStage.OnAfterDiagramPackageUpdateCallback.OnAfterUpdate(gongdocStage,
				buttonImplClassdiagram.diagramPackage,
				fakeFrontDiagramPackage,
			)

			// get the diagram package on the stage and set the selected diagram package
			var diagramPackage *gongdoc_models.DiagramPackage
			for _, _diagramPackage := range *gongdoc_models.GetGongstructInstancesMap[gongdoc_models.DiagramPackage](gongdocStage) {
				diagramPackage = _diagramPackage
			}
			diagramPackage.SelectedClassdiagram =
				(*gongdoc_models.GetGongstructInstancesMap[gongdoc_models.Classdiagram](gongdocStage))[selectedClassdiagramName]

			gongdocStage.Commit()
			return
		}
		buttonImplClassdiagram.classdiagram.IsInDrawMode = false
		buttonImplClassdiagram.classdiagramNode.IsInEditMode = false
		buttonImplClassdiagram.nodeImplClassdiagram.IsInDrawMode = false
	case BUTTON_save:

		// save the diagram
		if buttonImplClassdiagram.nodeImplClassdiagram.IsInDrawMode {

			// checkout in order to get the latest version of the diagram before
			// modifying it updated by the front
			gongdocStage.Checkout()
			gongdocStage.Unstage()
			gongdoc_models.StageBranch(gongdocStage, buttonImplClassdiagram.classdiagram)

			filepath := filepath.Join(
				filepath.Join(buttonImplClassdiagram.diagramPackage.AbsolutePathToDiagramPackage,
					"../diagrams"),
				buttonImplClassdiagram.classdiagram.Name) + ".go"
			file, err := os.Create(filepath)
			if err != nil {
				log.Fatal("Cannot open diagram file" + err.Error())
			}
			defer file.Close()

			mapDocLinkRemaping := &gongdocStage.Map_DocLink_Renaming
			_ = mapDocLinkRemaping

			gongdoc_models.SetupMapDocLinkRenaming(buttonImplClassdiagram.diagramPackage.ModelPkg.Stage_, gongdocStage)

			gongdocStage.Marshall(file, "github.com/fullstack-lang/gongdoc/go/models", "diagrams")

			// restore the original stage
			gongdocStage.Unstage()
			gongdocStage.Checkout()
		}

		gongdocStage.Commit()

	case BUTTON_delete:
		// checkout the stage, it shall remove the link between
		// the parent node and the staged node because 0..1->0..N association
		// is stored in the staged node as a reverse pointer
		gongdocStage.Checkout()

		// remove the classdiagram node from the pkg element node
		buttonImplClassdiagram.diagramPackage.Classdiagrams =
			remove(buttonImplClassdiagram.diagramPackage.Classdiagrams, buttonImplClassdiagram.classdiagram)
		gongdoc_models.UnstageBranch(gongdocStage, buttonImplClassdiagram.classdiagram)

		// remove the actual classdiagram file if it exsits
		classdiagramFilePath := filepath.Join(buttonImplClassdiagram.diagramPackage.Path, "../diagrams", buttonImplClassdiagram.classdiagram.Name) + ".go"
		if _, err := os.Stat(classdiagramFilePath); err == nil {
			if err := os.Remove(classdiagramFilePath); err != nil {
				log.Println("Error while deleting file " + classdiagramFilePath + " : " + err.Error())
			}
		}

		// reload
		fakeFrontDiagramPackage := (&gongdoc_models.DiagramPackage{})
		fakeFrontDiagramPackage.IsReloaded = buttonImplClassdiagram.diagramPackage.IsReloaded
		buttonImplClassdiagram.diagramPackage.IsReloaded = !buttonImplClassdiagram.diagramPackage.IsReloaded

		gongdocStage.OnAfterDiagramPackageUpdateCallback.OnAfterUpdate(gongdocStage,
			buttonImplClassdiagram.diagramPackage,
			fakeFrontDiagramPackage,
		)
	case BUTTON_file_copy:
		var hasNameCollision bool
		initialName := buttonImplClassdiagram.classdiagram.Name
		newName := initialName
		index := 0
		// loop until the name of the new diagram is not in collision with an existing
		// diagram name
		for index == 0 || hasNameCollision {
			index++
			hasNameCollision = false
			for classdiagram := range *gongdoc_models.GetGongstructInstancesSet[gongdoc_models.Classdiagram](gongdocStage) {
				if classdiagram.Name == newName {
					hasNameCollision = true
				}
			}
			if hasNameCollision {
				newName = initialName + fmt.Sprintf("_%d", index)
			}
		}

		log.Println("New name is", newName)
		newClassdiagramFilePath := filepath.Join(buttonImplClassdiagram.diagramPackage.Path, "../diagrams", newName) + ".go"

		file, err := os.Create(newClassdiagramFilePath)
		if err != nil {
			log.Fatal("Cannot open diagram file" + err.Error())
		}
		defer file.Close()

		gongdocStage.Checkout()
		gongdocStage.Unstage()
		gongdoc_models.StageBranch(gongdocStage, buttonImplClassdiagram.classdiagram)
		buttonImplClassdiagram.classdiagram.Name = newName

		gongdoc_models.SetupMapDocLinkRenaming(buttonImplClassdiagram.diagramPackage.ModelPkg.Stage_, gongdocStage)
		gongdocStage.Marshall(file, "github.com/fullstack-lang/gongdoc/go/models", "diagrams")

		// marshall the new diagram to the new file
		gongdocStage.Checkout()
		gongdocStage.Unstage()
		gongdoc_models.StageBranch(gongdocStage, buttonImplClassdiagram.classdiagram)
		buttonImplClassdiagram.classdiagram.Name = newName

		gongdocStage.Checkout()
		gongdocStage.Commit()

		// reload
		fakeFrontDiagramPackage := (&gongdoc_models.DiagramPackage{})
		fakeFrontDiagramPackage.IsReloaded = buttonImplClassdiagram.diagramPackage.IsReloaded
		buttonImplClassdiagram.diagramPackage.IsReloaded = !buttonImplClassdiagram.diagramPackage.IsReloaded

		gongdocStage.OnAfterDiagramPackageUpdateCallback.OnAfterUpdate(gongdocStage,
			buttonImplClassdiagram.diagramPackage,
			fakeFrontDiagramPackage,
		)
	case BUTTON_edit:
		// change the name of the diagram
		buttonImplClassdiagram.classdiagramNode.IsInEditMode = true
	default:
		log.Fatalln("Unkown button type", buttonImplClassdiagram.Icon)
	}

	computeNodeConfs(
		gongtreeStage,
		gongdocStage,
		buttonImplClassdiagram.diagramPackageNode,
		buttonImplClassdiagram.diagramPackage,
		buttonImplClassdiagram.treeOfGongObjects)

	gongdocStage.Commit()
	gongtreeStage.Commit()
}
