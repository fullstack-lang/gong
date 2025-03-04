// generated code - do not edit
package probe

import (
	"log"

	gong_models "github.com/fullstack-lang/gong/go/models"
	gongtree_models "github.com/fullstack-lang/gong/lib/tree/go/models"

	"github.com/fullstack-lang/gong/lib/doc/go/models"
)

type TreeNodeImplGongstruct struct {
	gongStruct *gong_models.GongStruct
	probe *Probe
}

func NewTreeNodeImplGongstruct(
	gongStruct *gong_models.GongStruct,
	probe *Probe,
) (nodeImplGongstruct *TreeNodeImplGongstruct) {

	nodeImplGongstruct = new(TreeNodeImplGongstruct)
	nodeImplGongstruct.gongStruct = gongStruct
	nodeImplGongstruct.probe = probe
	return
}

func (nodeImplGongstruct *TreeNodeImplGongstruct) OnAfterUpdate(
	gongtreeStage *gongtree_models.StageStruct,
	stagedNode, frontNode *gongtree_models.Node) {

	// setting the value of the staged node	to the new value
	// otherwise, the expansion memory is lost
	if stagedNode.IsExpanded != frontNode.IsExpanded {
		stagedNode.IsExpanded = frontNode.IsExpanded
		return
	}

	// if node is unchecked
	if stagedNode.IsChecked && !frontNode.IsChecked {

	}

	// if node is checked, add gongstructshape
	if !stagedNode.IsChecked && frontNode.IsChecked {

	}

	// the node was selected. Therefore, one request the
	// table to route to the table
	log.Println("NodeImplGongstruct:OnAfterUpdate with: ", nodeImplGongstruct.gongStruct.GetName())

	// insertion point
	if nodeImplGongstruct.gongStruct.GetName() == "Classdiagram" {
		fillUpTable[models.Classdiagram](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "DiagramPackage" {
		fillUpTable[models.DiagramPackage](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Field" {
		fillUpTable[models.Field](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "GongEnumShape" {
		fillUpTable[models.GongEnumShape](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "GongEnumValueEntry" {
		fillUpTable[models.GongEnumValueEntry](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "GongStructShape" {
		fillUpTable[models.GongStructShape](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Link" {
		fillUpTable[models.Link](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "NoteShape" {
		fillUpTable[models.NoteShape](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "NoteShapeLink" {
		fillUpTable[models.NoteShapeLink](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Position" {
		fillUpTable[models.Position](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "UmlState" {
		fillUpTable[models.UmlState](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Umlsc" {
		fillUpTable[models.Umlsc](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Vertice" {
		fillUpTable[models.Vertice](nodeImplGongstruct.probe)
	}

	// set color for node and reset all other nodes color
	for node := range *gongtree_models.GetGongstructInstancesSet[gongtree_models.Node](gongtreeStage) {
		node.BackgroundColor = ""
	}
	stagedNode.BackgroundColor = "lightgrey"
	gongtreeStage.Commit()

	nodeImplGongstruct.probe.tableStage.Commit()
}
