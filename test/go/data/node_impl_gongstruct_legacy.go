package data

import (
	"log"
	"strings"

	gong_models "github.com/fullstack-lang/gong/go/models"
	gongrouter_models "github.com/fullstack-lang/gongrouter/go/models"
	gongtree_models "github.com/fullstack-lang/gongtree/go/models"
)

type NodeImplGongstructLegacy struct {
	gongStruct      *gong_models.GongStruct
	gongrouterStage *gongrouter_models.StageStruct
	tableRouter     *gongrouter_models.Outlet
}

func NewNodeImplGongstructLegacy(
	gongStruct *gong_models.GongStruct,
	gongrouterStage *gongrouter_models.StageStruct,
	tableRouter *gongrouter_models.Outlet,
) (nodeImplGongstructLegacy *NodeImplGongstructLegacy) {

	nodeImplGongstructLegacy = new(NodeImplGongstructLegacy)
	nodeImplGongstructLegacy.gongStruct = gongStruct
	nodeImplGongstructLegacy.gongrouterStage = gongrouterStage
	nodeImplGongstructLegacy.tableRouter = tableRouter
	return
}

func (nodeImplGongstructLegacy *NodeImplGongstructLegacy) OnAfterUpdate(
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
	// router to route to the table
	log.Println("NodeImplGongstruct:OnAfterUpdate")
	nodeImplGongstructLegacy.tableRouter.Path =
		"github_com_fullstack_lang_gong_test_go-" +
			strings.ToLower(nodeImplGongstructLegacy.gongStruct.Name) + "s"

	nodeImplGongstructLegacy.gongrouterStage.Commit()

}
