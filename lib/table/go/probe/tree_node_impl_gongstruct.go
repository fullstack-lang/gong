// generated code - do not edit
package probe

import (
	gong_models "github.com/fullstack-lang/gong/go/models"
	gongtree_models "github.com/fullstack-lang/gong/lib/tree/go/models"

	"github.com/fullstack-lang/gong/lib/table/go/models"
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
	if nodeImplGongstruct.gongStruct.GetName() == "Cell" {
		updateAndCommitTable[models.Cell](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "CellBoolean" {
		updateAndCommitTable[models.CellBoolean](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "CellFloat64" {
		updateAndCommitTable[models.CellFloat64](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "CellIcon" {
		updateAndCommitTable[models.CellIcon](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "CellInt" {
		updateAndCommitTable[models.CellInt](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "CellString" {
		updateAndCommitTable[models.CellString](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "CheckBox" {
		updateAndCommitTable[models.CheckBox](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "DisplayedColumn" {
		updateAndCommitTable[models.DisplayedColumn](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "FormDiv" {
		updateAndCommitTable[models.FormDiv](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "FormEditAssocButton" {
		updateAndCommitTable[models.FormEditAssocButton](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "FormField" {
		updateAndCommitTable[models.FormField](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "FormFieldDate" {
		updateAndCommitTable[models.FormFieldDate](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "FormFieldDateTime" {
		updateAndCommitTable[models.FormFieldDateTime](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "FormFieldFloat64" {
		updateAndCommitTable[models.FormFieldFloat64](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "FormFieldInt" {
		updateAndCommitTable[models.FormFieldInt](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "FormFieldSelect" {
		updateAndCommitTable[models.FormFieldSelect](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "FormFieldString" {
		updateAndCommitTable[models.FormFieldString](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "FormFieldTime" {
		updateAndCommitTable[models.FormFieldTime](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "FormGroup" {
		updateAndCommitTable[models.FormGroup](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "FormSortAssocButton" {
		updateAndCommitTable[models.FormSortAssocButton](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Option" {
		updateAndCommitTable[models.Option](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Row" {
		updateAndCommitTable[models.Row](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Table" {
		updateAndCommitTable[models.Table](nodeImplGongstruct.probe)
	}

	// set color for node and reset all other nodes color
	for node := range *gongtree_models.GetGongstructInstancesSet[gongtree_models.Node](gongtreeStage) {
		node.BackgroundColor = ""
	}
	stagedNode.BackgroundColor = "lightgrey"
	gongtreeStage.Commit()

	nodeImplGongstruct.probe.tableStage.Commit()
}
