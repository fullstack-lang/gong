// generated code - do not edit
package probe

import (
	gong_models "github.com/fullstack-lang/gong/go/models"
	gongtree_models "github.com/fullstack-lang/gong/lib/tree/go/models"

	"github.com/fullstack-lang/gong/lib/split/go/models"
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
	if nodeImplGongstruct.gongStruct.GetName() == "AsSplit" {
		updateAndCommitTable[models.AsSplit](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "AsSplitArea" {
		updateAndCommitTable[models.AsSplitArea](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Button" {
		updateAndCommitTable[models.Button](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Cursor" {
		updateAndCommitTable[models.Cursor](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Doc" {
		updateAndCommitTable[models.Doc](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "FavIcon" {
		updateAndCommitTable[models.FavIcon](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Form" {
		updateAndCommitTable[models.Form](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Load" {
		updateAndCommitTable[models.Load](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "LogoOnTheLeft" {
		updateAndCommitTable[models.LogoOnTheLeft](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "LogoOnTheRight" {
		updateAndCommitTable[models.LogoOnTheRight](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Slider" {
		updateAndCommitTable[models.Slider](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Split" {
		updateAndCommitTable[models.Split](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Svg" {
		updateAndCommitTable[models.Svg](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Table" {
		updateAndCommitTable[models.Table](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Title" {
		updateAndCommitTable[models.Title](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Tone" {
		updateAndCommitTable[models.Tone](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Tree" {
		updateAndCommitTable[models.Tree](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "View" {
		updateAndCommitTable[models.View](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Xlsx" {
		updateAndCommitTable[models.Xlsx](nodeImplGongstruct.probe)
	}

	// set color for node and reset all other nodes color
	for node := range *gongtree_models.GetGongstructInstancesSet[gongtree_models.Node](gongtreeStage) {
		node.BackgroundColor = ""
	}
	stagedNode.BackgroundColor = "lightgrey"
	gongtreeStage.Commit()

	nodeImplGongstruct.probe.tableStage.Commit()
}
