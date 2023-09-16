// generated code - do not edit
package probe

import (
	"embed"

	"github.com/gin-gonic/gin"

	gongtree_fullstack "github.com/fullstack-lang/gongtree/go/fullstack"

	gongtable_fullstack "github.com/fullstack-lang/gongtable/go/fullstack"

	gong_fullstack "github.com/fullstack-lang/gong/go/fullstack"
	gong_models "github.com/fullstack-lang/gong/go/models"

	"github.com/fullstack-lang/gong/test/go/models"
	"github.com/fullstack-lang/gong/test/go/orm"
)

func NewProbe(
	r *gin.Engine,
	goModelsDir embed.FS,
	stackPath string,
	stageOfInterest *models.StageStruct,
	backRepoOfInterest *orm.BackRepoStruct) {

	gongStage, _ := gong_fullstack.NewStackInstance(r, stackPath)

	gong_models.LoadEmbedded(gongStage, goModelsDir)

	// treeForSelectingDate that is on the sidebar
	treeStage, _ := gongtree_fullstack.NewStackInstance(r, stackPath+"-sidebar")

	// stage for main table
	tableStage, _ := gongtable_fullstack.NewStackInstance(r, stackPath)
	tableStage.Commit()

	// stage for reusable form
	formStage, _ := gongtable_fullstack.NewStackInstance(r, stackPath+"-form")
	formStage.Commit()

	playground := NewPlayground(
		r,
		stageOfInterest,
		backRepoOfInterest,
		gongStage,
		treeStage,
		formStage,
		tableStage,
	)

	fillUpTree(playground)

}
