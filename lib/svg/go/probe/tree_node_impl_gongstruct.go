// generated code - do not edit
package probe

import (
	gong_models "github.com/fullstack-lang/gong/go/models"
	gongtree_models "github.com/fullstack-lang/gong/lib/tree/go/models"

	"github.com/fullstack-lang/gong/lib/svg/go/models"
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
	if nodeImplGongstruct.gongStruct.GetName() == "Animate" {
		updateProbeTable[*models.Animate](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Circle" {
		updateProbeTable[*models.Circle](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Condition" {
		updateProbeTable[*models.Condition](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "ControlPoint" {
		updateProbeTable[*models.ControlPoint](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Ellipse" {
		updateProbeTable[*models.Ellipse](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Layer" {
		updateProbeTable[*models.Layer](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Line" {
		updateProbeTable[*models.Line](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Link" {
		updateProbeTable[*models.Link](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "LinkAnchoredText" {
		updateProbeTable[*models.LinkAnchoredText](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Path" {
		updateProbeTable[*models.Path](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Point" {
		updateProbeTable[*models.Point](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Polygone" {
		updateProbeTable[*models.Polygone](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Polyline" {
		updateProbeTable[*models.Polyline](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Rect" {
		updateProbeTable[*models.Rect](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "RectAnchoredPath" {
		updateProbeTable[*models.RectAnchoredPath](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "RectAnchoredRect" {
		updateProbeTable[*models.RectAnchoredRect](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "RectAnchoredText" {
		updateProbeTable[*models.RectAnchoredText](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "RectLinkLink" {
		updateProbeTable[*models.RectLinkLink](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "SVG" {
		updateProbeTable[*models.SVG](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "SvgText" {
		updateProbeTable[*models.SvgText](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Text" {
		updateProbeTable[*models.Text](nodeImplGongstruct.probe)
	}

	// set color for node and reset all other nodes color
	for node := range *gongtree_models.GetGongstructInstancesSet[gongtree_models.Node](gongtreeStage) {
		node.BackgroundColor = ""
	}
	stagedNode.BackgroundColor = "lightgrey"
	gongtreeStage.Commit()
}
