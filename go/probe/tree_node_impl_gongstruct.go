// generated code - do not edit
package probe

import (
	"log"

	gong_models "github.com/fullstack-lang/gong/go/models"
	gongtree_models "github.com/fullstack-lang/gongtree/go/models"

	"github.com/fullstack-lang/gong/go/models"
)

type TreeNodeImplGongstruct struct {
	gongStruct *gong_models.GongStruct
	playground *Playground
}

func NewTreeNodeImplGongstruct(
	gongStruct *gong_models.GongStruct,
	playground *Playground,
) (nodeImplGongstruct *TreeNodeImplGongstruct) {

	nodeImplGongstruct = new(TreeNodeImplGongstruct)
	nodeImplGongstruct.gongStruct = gongStruct
	nodeImplGongstruct.playground = playground
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
	if nodeImplGongstruct.gongStruct.GetName() == "GongBasicField" {
		fillUpTable[models.GongBasicField](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "GongEnum" {
		fillUpTable[models.GongEnum](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "GongEnumValue" {
		fillUpTable[models.GongEnumValue](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "GongLink" {
		fillUpTable[models.GongLink](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "GongNote" {
		fillUpTable[models.GongNote](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "GongStruct" {
		fillUpTable[models.GongStruct](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "GongTimeField" {
		fillUpTable[models.GongTimeField](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Meta" {
		fillUpTable[models.Meta](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "MetaReference" {
		fillUpTable[models.MetaReference](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "ModelPkg" {
		fillUpTable[models.ModelPkg](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "PointerToGongStructField" {
		fillUpTable[models.PointerToGongStructField](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "SliceOfPointerToGongStructField" {
		fillUpTable[models.SliceOfPointerToGongStructField](nodeImplGongstruct.playground)
	}

	// set color for node and reset all other nodes color
	for node := range *gongtree_models.GetGongstructInstancesSet[gongtree_models.Node](gongtreeStage) {
		node.BackgroundColor = ""
	}
	stagedNode.BackgroundColor = "lightgrey"
	gongtreeStage.Commit()

	nodeImplGongstruct.playground.tableStage.Commit()
}
