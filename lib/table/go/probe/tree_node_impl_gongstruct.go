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
		updateProbeTable[*models.Cell](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "CellBoolean" {
		updateProbeTable[*models.CellBoolean](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "CellFloat64" {
		updateProbeTable[*models.CellFloat64](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "CellIcon" {
		updateProbeTable[*models.CellIcon](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "CellInt" {
		updateProbeTable[*models.CellInt](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "CellString" {
		updateProbeTable[*models.CellString](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "CheckBox" {
		updateProbeTable[*models.CheckBox](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "DisplayedColumn" {
		updateProbeTable[*models.DisplayedColumn](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "FormDiv" {
		updateProbeTable[*models.FormDiv](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "FormEditAssocButton" {
		updateProbeTable[*models.FormEditAssocButton](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "FormField" {
		updateProbeTable[*models.FormField](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "FormFieldDate" {
		updateProbeTable[*models.FormFieldDate](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "FormFieldDateTime" {
		updateProbeTable[*models.FormFieldDateTime](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "FormFieldFloat64" {
		updateProbeTable[*models.FormFieldFloat64](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "FormFieldInt" {
		updateProbeTable[*models.FormFieldInt](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "FormFieldSelect" {
		updateProbeTable[*models.FormFieldSelect](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "FormFieldString" {
		updateProbeTable[*models.FormFieldString](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "FormFieldTime" {
		updateProbeTable[*models.FormFieldTime](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "FormGroup" {
		updateProbeTable[*models.FormGroup](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "FormSortAssocButton" {
		updateProbeTable[*models.FormSortAssocButton](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Option" {
		updateProbeTable[*models.Option](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Row" {
		updateProbeTable[*models.Row](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Table" {
		updateProbeTable[*models.Table](nodeImplGongstruct.probe)
	}

	// set color for node and reset all other nodes color
	for node := range *gongtree_models.GetGongstructInstancesSet[gongtree_models.Node](gongtreeStage) {
		node.BackgroundColor = ""
	}
	stagedNode.BackgroundColor = "lightgrey"
	gongtreeStage.Commit()
}
