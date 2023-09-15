// generated code - do not edit
package probe

import (
	"log"

	gong_models "github.com/fullstack-lang/gong/go/models"
	gongtree_models "github.com/fullstack-lang/gongtree/go/models"

	"github.com/fullstack-lang/gong/test/go/models"
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
	if nodeImplGongstruct.gongStruct.GetName() == "Astruct" {
		fillUpTable[models.Astruct](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "AstructBstruct2Use" {
		fillUpTable[models.AstructBstruct2Use](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "AstructBstructUse" {
		fillUpTable[models.AstructBstructUse](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Bstruct" {
		fillUpTable[models.Bstruct](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Dstruct" {
		fillUpTable[models.Dstruct](nodeImplGongstruct.playground)
	}

	nodeImplGongstruct.playground.tableStage.Commit()
}
