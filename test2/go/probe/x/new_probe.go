// generated code - do not edit
package x

import (
	"embed"

	"github.com/gin-gonic/gin"

	gongtable_fullstack "github.com/fullstack-lang/gongtable/go/fullstack"
	gongtree_fullstack "github.com/fullstack-lang/gongtree/go/fullstack"

	gong_fullstack "github.com/fullstack-lang/gong/go/fullstack"
	gong_models "github.com/fullstack-lang/gong/go/models"
	"github.com/fullstack-lang/gong/test2/go/models/x"
	orm_x "github.com/fullstack-lang/gong/test2/go/orm/x"

	gongdoc_load "github.com/fullstack-lang/gongdoc/go/load"
)

func NewProbe(
	r *gin.Engine,
	goModelsDir embed.FS,
	goDiagramsDir embed.FS,
	embeddedDiagrams bool,
	stackPath string,
	stageOfInterest *x.StageStruct,
	backRepoOfInterest *orm_x.BackRepoStruct) {

	gongStage, _ := gong_fullstack.NewStackInstance(r, stackPath)

	gong_models.LoadEmbedded(gongStage, goModelsDir)

	// treeForSelectingDate that is on the sidebar
	treeStage, _ := gongtree_fullstack.NewStackInstance(r, stackPath+"-sidebar")

	// stage for main table
	tableStage, _ := gongtable_fullstack.NewStackInstance(r, stackPath+"-table")
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

	gongdoc_load.Load(
		"",
		"github.com/fullstack-lang/gong/test2/go/models",
		goModelsDir,
		goDiagramsDir,
		r,
		embeddedDiagrams,
		&stageOfInterest.Map_GongStructName_InstancesNb)

}
