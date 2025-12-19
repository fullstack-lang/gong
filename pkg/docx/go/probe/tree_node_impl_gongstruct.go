// generated code - do not edit
package probe

import (
	gong_models "github.com/fullstack-lang/gong/go/models"
	gongtree_models "github.com/fullstack-lang/gong/lib/tree/go/models"

	"github.com/fullstack-lang/gong/pkg/docx/go/models"
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
	if nodeImplGongstruct.gongStruct.GetName() == "Body" {
		updateProbeTable[*models.Body](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Document" {
		updateProbeTable[*models.Document](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Docx" {
		updateProbeTable[*models.Docx](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "File" {
		updateProbeTable[*models.File](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Node" {
		updateProbeTable[*models.Node](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Paragraph" {
		updateProbeTable[*models.Paragraph](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "ParagraphProperties" {
		updateProbeTable[*models.ParagraphProperties](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "ParagraphStyle" {
		updateProbeTable[*models.ParagraphStyle](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Rune" {
		updateProbeTable[*models.Rune](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "RuneProperties" {
		updateProbeTable[*models.RuneProperties](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Table" {
		updateProbeTable[*models.Table](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "TableColumn" {
		updateProbeTable[*models.TableColumn](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "TableProperties" {
		updateProbeTable[*models.TableProperties](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "TableRow" {
		updateProbeTable[*models.TableRow](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "TableStyle" {
		updateProbeTable[*models.TableStyle](nodeImplGongstruct.probe)
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
