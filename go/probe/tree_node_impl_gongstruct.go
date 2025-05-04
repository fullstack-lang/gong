// generated code - do not edit
package probe

import (
	gong_models "github.com/fullstack-lang/gong/go/models"
	gongtree_models "github.com/fullstack-lang/gong/lib/tree/go/models"

	"github.com/fullstack-lang/gong/go/models"
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
	if nodeImplGongstruct.gongStruct.GetName() == "GongBasicField" {
		updateAndCommitTable[models.GongBasicField](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "GongEnum" {
		updateAndCommitTable[models.GongEnum](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "GongEnumValue" {
		updateAndCommitTable[models.GongEnumValue](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "GongLink" {
		updateAndCommitTable[models.GongLink](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "GongNote" {
		updateAndCommitTable[models.GongNote](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "GongStruct" {
		updateAndCommitTable[models.GongStruct](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "GongTimeField" {
		updateAndCommitTable[models.GongTimeField](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Meta" {
		updateAndCommitTable[models.Meta](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "MetaReference" {
		updateAndCommitTable[models.MetaReference](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "ModelPkg" {
		updateAndCommitTable[models.ModelPkg](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "PointerToGongStructField" {
		updateAndCommitTable[models.PointerToGongStructField](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "SliceOfPointerToGongStructField" {
		updateAndCommitTable[models.SliceOfPointerToGongStructField](nodeImplGongstruct.probe)
	}

	// set color for node and reset all other nodes color
	for node := range *gongtree_models.GetGongstructInstancesSet[gongtree_models.Node](gongtreeStage) {
		node.BackgroundColor = ""
	}
	stagedNode.BackgroundColor = "lightgrey"
	gongtreeStage.Commit()

	nodeImplGongstruct.probe.tableStage.Commit()
}
