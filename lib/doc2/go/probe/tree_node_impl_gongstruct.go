// generated code - do not edit
package probe

import (
	gong_models "github.com/fullstack-lang/gong/go/models"
	gongtree_models "github.com/fullstack-lang/gong/lib/tree/go/models"

	"github.com/fullstack-lang/gong/lib/doc2/go/models"
)

type TreeNodeImplGongstruct struct {
	gongStruct *gong_models.GongStruct
	probe      *Probe
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
	gongtreeStage *gongtree_models.Stage,
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
	// log.Println("NodeImplGongstruct:OnAfterUpdate with: ", nodeImplGongstruct.gongStruct.GetName())

	// insertion point
	if nodeImplGongstruct.gongStruct.GetName() == "AttributeShape" {
		updateAndCommitTable[models.AttributeShape](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Classdiagram" {
		updateAndCommitTable[models.Classdiagram](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "DiagramPackage" {
		updateAndCommitTable[models.DiagramPackage](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "GongEnumShape" {
		updateAndCommitTable[models.GongEnumShape](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "GongEnumValueShape" {
		updateAndCommitTable[models.GongEnumValueShape](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "GongNoteLinkShape" {
		updateAndCommitTable[models.GongNoteLinkShape](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "GongNoteShape" {
		updateAndCommitTable[models.GongNoteShape](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "GongStructShape" {
		updateAndCommitTable[models.GongStructShape](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "LinkShape" {
		updateAndCommitTable[models.LinkShape](nodeImplGongstruct.probe)
	}

	// set color for node and reset all other nodes color
	for node := range *gongtree_models.GetGongstructInstancesSet[gongtree_models.Node](gongtreeStage) {
		node.BackgroundColor = ""
	}
	stagedNode.BackgroundColor = "lightgrey"
	gongtreeStage.Commit()

	nodeImplGongstruct.probe.tableStage.Commit()
}
