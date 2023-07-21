package golang

const NodeImplGongstructFileTemplate = `package data

import (
	"log"
	"strings"

	gong_models "github.com/fullstack-lang/gong/go/models"
	gongrouter_models "github.com/fullstack-lang/gongrouter/go/models"
	gongtree_models "github.com/fullstack-lang/gongtree/go/models"
)

type NodeImplGongstruct struct {
	gongStruct      *gong_models.GongStruct
	gongrouterStage *gongrouter_models.StageStruct
	tableRouter     *gongrouter_models.Outlet
}

func NewNodeImplGongstruct(
	gongStruct *gong_models.GongStruct,
	gongrouterStage *gongrouter_models.StageStruct,
	tableRouter *gongrouter_models.Outlet,
) (nodeImplGongstruct *NodeImplGongstruct) {

	nodeImplGongstruct = new(NodeImplGongstruct)
	nodeImplGongstruct.gongStruct = gongStruct
	nodeImplGongstruct.gongrouterStage = gongrouterStage
	nodeImplGongstruct.tableRouter = tableRouter
	return
}

func (nodeImplGongstruct *NodeImplGongstruct) OnAfterUpdate(
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
	nodeImplGongstruct.tableRouter.Path =
		"{{PkgPathRootWithoutSlashes}}-" +
			strings.ToLower(nodeImplGongstruct.gongStruct.Name) + "s"

	nodeImplGongstruct.gongrouterStage.Commit()

}
`
