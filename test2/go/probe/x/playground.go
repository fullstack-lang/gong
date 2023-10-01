// generated code - do not edit
package x

import (
	"github.com/gin-gonic/gin"

	gong "github.com/fullstack-lang/gong/go/models"
	"github.com/fullstack-lang/gong/test2/go/models/x"
	orm_x "github.com/fullstack-lang/gong/test2/go/orm/x"
	form "github.com/fullstack-lang/gongtable/go/models"
	tree "github.com/fullstack-lang/gongtree/go/models"
)

type Playground struct {
	r                  *gin.Engine
	stageOfInterest    *x.StageStruct
	backRepoOfInterest *orm_x.BackRepoStruct
	gongStage          *gong.StageStruct
	treeStage          *tree.StageStruct
	formStage          *form.StageStruct
	tableStage         *form.StageStruct
}

func NewPlayground(
	r *gin.Engine,
	stageOfInterest *x.StageStruct,
	backRepoOfInterest *orm_x.BackRepoStruct,
	gongStage *gong.StageStruct,
	treeStage *tree.StageStruct,
	formStage *form.StageStruct,
	tableStage *form.StageStruct,
) (playground *Playground) {
	playground = new(Playground)
	playground.r = r
	playground.stageOfInterest = stageOfInterest
	playground.backRepoOfInterest = backRepoOfInterest
	playground.gongStage = gongStage
	playground.treeStage = treeStage
	playground.formStage = formStage
	playground.tableStage = tableStage

	return
}
