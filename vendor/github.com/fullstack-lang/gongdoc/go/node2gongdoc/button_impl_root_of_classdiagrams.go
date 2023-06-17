package node2gongdoc

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"
	gongtree_models "github.com/fullstack-lang/gongtree/go/models"
)

type ButtonImplRootOfClassdiagrams struct {
	// one need to access the daigramPackage to set the selected diagram
	diagramPackage *gongdoc_models.DiagramPackage

	// one needs to access the node of the diagram package to manage the childern nodes
	diagramPackageNode *gongtree_models.Node

	// one needs to perform computation of node confs after the update
	treeOfGongObjects *gongtree_models.Tree

	// type of button
	Icon ButtonType
}

func NewButtonImplRootOfClassdiagrams(
	diagramPackage *gongdoc_models.DiagramPackage,
	diagramPackageNode *gongtree_models.Node,
	treeOfGongObjects *gongtree_models.Tree,
	icon ButtonType,
) (buttonImplRootOfClassdiagrams *ButtonImplRootOfClassdiagrams) {

	buttonImplRootOfClassdiagrams = new(ButtonImplRootOfClassdiagrams)

	buttonImplRootOfClassdiagrams.diagramPackage = diagramPackage
	buttonImplRootOfClassdiagrams.diagramPackageNode = diagramPackageNode
	buttonImplRootOfClassdiagrams.treeOfGongObjects = treeOfGongObjects
	buttonImplRootOfClassdiagrams.Icon = icon

	return
}

func (buttonImpl *ButtonImplRootOfClassdiagrams) ButtonUpdated(
	gongtreeStage *gongtree_models.StageStruct,
	stageButton, front *gongtree_models.Button) {

	gongdocStage := buttonImpl.diagramPackage.Stage_

	switch buttonImpl.Icon {
	case BUTTON_add:

		// check unicity of name, otherwise, add an index
		var hasNameCollision bool
		initialName := "Default"
		name := initialName
		index := 0
		// loop until the name of the new diagram is not in collision with an existing
		// diagram name
		for index == 0 || hasNameCollision {
			index++
			hasNameCollision = false
			for classdiagram := range *gongdoc_models.GetGongstructInstancesSet[gongdoc_models.Classdiagram](gongdocStage) {
				if classdiagram.Name == name {
					hasNameCollision = true
				}
			}
			if hasNameCollision {
				name = initialName + fmt.Sprintf("_%d", index)
			}
		}

		classdiagram := (&gongdoc_models.Classdiagram{Name: name}).Stage(gongdocStage)
		buttonImpl.diagramPackage.Classdiagrams =
			append(buttonImpl.diagramPackage.Classdiagrams, classdiagram)

		filepath := filepath.Join(
			filepath.Join(buttonImpl.diagramPackage.AbsolutePathToDiagramPackage,
				"../diagrams"),
			classdiagram.Name) + ".go"
		file, err := os.Create(filepath)
		if err != nil {
			log.Fatal("Cannot open diagram file" + err.Error())
		}
		defer file.Close()

		gongdocStage.Commit()

		// now save the diagram
		gongdocStage.Checkout()
		gongdocStage.Unstage()
		gongdoc_models.StageBranch(gongdocStage, classdiagram)

		gongdoc_models.SetupMapDocLinkRenaming(buttonImpl.diagramPackage.ModelPkg.Stage_, gongdocStage)

		gongdocStage.Marshall(file, "github.com/fullstack-lang/gongdoc/go/models", "diagrams")

		// restore the original stage
		gongdocStage.Unstage()
		gongdocStage.Checkout()

		classdiagramNode := NewClassdiagramNode(
			gongtreeStage,
			classdiagram,
			buttonImpl.diagramPackage,
			buttonImpl.diagramPackageNode,
			buttonImpl.treeOfGongObjects)
		buttonImpl.diagramPackageNode.Children =
			append(buttonImpl.diagramPackageNode.Children, classdiagramNode)
	default:
		log.Fatalln("Unkown button type", buttonImpl.Icon)
	}

	computeNodeConfs(
		gongtreeStage,
		gongdocStage,
		buttonImpl.diagramPackageNode,
		buttonImpl.diagramPackage,
		buttonImpl.treeOfGongObjects)

	gongdocStage.Commit()
	gongtreeStage.Commit()
}
