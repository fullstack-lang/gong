package node2gongdoc

import (
	"log"

	gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"
	gongtree_models "github.com/fullstack-lang/gongtree/go/models"
)

func computeClassdiagramNodesConfigurations(
	diagramPackageNode *gongtree_models.Node,
	diagramPackage *gongdoc_models.DiagramPackage,
	gongdocStage *gongdoc_models.StageStruct,
	gongtreeStage *gongtree_models.StageStruct,
) {

	// compute wether one of the diagrams is in draw/edit mode
	// if so, all diagram check need to be disabled
	var inModificationMode bool
	for _, classdiagramNode := range diagramPackageNode.Children {

		nodeImplClasssiagram, ok := classdiagramNode.Impl.(*NodeImplClasssiagram)
		if !ok {
			log.Fatalln("not a good interface")
		}
		if nodeImplClasssiagram.IsInDrawMode {
			inModificationMode = true
		}
	}

	// get the selected diagram and collect what are its referenced
	// gongstructs
	for _, classdiagramNode := range diagramPackageNode.Children {

		// remove all buttons
		for _, _button := range classdiagramNode.Buttons {
			_button.Unstage(gongtreeStage)
		}
		classdiagramNode.Buttons = make([]*gongtree_models.Button, 0)

		nodeImplClasssiagram, ok := classdiagramNode.Impl.(*NodeImplClasssiagram)
		if !ok {
			log.Fatalln("not a good interface")
		}

		classdiagramNode.IsCheckboxDisabled = inModificationMode

		if !classdiagramNode.IsChecked {
			continue
		}

		selectedForEdit := diagramPackage.IsEditable && !nodeImplClasssiagram.IsInDrawMode && !classdiagramNode.IsInEditMode
		inDrawingMode := diagramPackage.IsEditable && nodeImplClasssiagram.IsInDrawMode
		inEditMode := diagramPackage.IsEditable && classdiagramNode.IsInEditMode

		AddNewButton(gongtreeStage, classdiagramNode, diagramPackageNode, BUTTON_draw, selectedForEdit)
		AddNewButton(gongtreeStage, classdiagramNode, diagramPackageNode, BUTTON_delete, selectedForEdit)
		AddNewButton(gongtreeStage, classdiagramNode, diagramPackageNode, BUTTON_file_copy, selectedForEdit)
		AddNewButton(gongtreeStage, classdiagramNode, diagramPackageNode, BUTTON_edit, selectedForEdit)

		AddNewButton(gongtreeStage, classdiagramNode, diagramPackageNode, BUTTON_edit_off, inDrawingMode || inEditMode)
		AddNewButton(gongtreeStage, classdiagramNode, diagramPackageNode, BUTTON_save, inDrawingMode)

	}
}

// AddNewButton add a button if confirmation is true
func AddNewButton(
	gongtreeStage *gongtree_models.StageStruct,
	classdiagramNode *gongtree_models.Node,
	diagramPackageNode *gongtree_models.Node,
	icon ButtonType, confirmation bool) {

	nodeImplClassdiagram, ok := classdiagramNode.Impl.(*NodeImplClasssiagram)
	if !ok {
		log.Fatalln("youre kidding me ?")
	}

	if confirmation == false {
		return
	}

	drawButton := (&gongtree_models.Button{
		Name: nodeImplClassdiagram.classdiagram.Name + " " + string(icon),
		Icon: string(icon)}).Stage(gongtreeStage)
	_ = drawButton
	classdiagramNode.Buttons = append(classdiagramNode.Buttons, drawButton)
	drawButton.Impl = NewButtonImplClassdiagram(
		nodeImplClassdiagram.diagramPackage,
		nodeImplClassdiagram.classdiagram,
		diagramPackageNode,
		nodeImplClassdiagram.treeOfGongObjects,
		classdiagramNode,
		nodeImplClassdiagram,
		icon,
	)

}
