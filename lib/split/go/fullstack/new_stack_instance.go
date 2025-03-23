// do not modify, generated file
package fullstack

import (
	"github.com/fullstack-lang/gong/lib/split/go/controllers"
	"github.com/fullstack-lang/gong/lib/split/go/models"
	"github.com/fullstack-lang/gong/lib/split/go/orm"

	"github.com/gin-gonic/gin"

	// this will import the angular front end source code directory (versionned with git) in the vendor directory
	// this path will be included in the "tsconfig.json" front end compilation paths
	// to include this stack front end code
	_ "github.com/fullstack-lang/gong/lib/split/ng-github.com-fullstack-lang-gong-lib-split"
)

// NewStackInstance creates a new stack instance from the Stack Model
// and returns the backRepo of the stack instance (you can get the stage from backRepo.GetStage()
//
// - the stackPath is the unique identifier of the stack
// - the optional parameter filenames is for the name of the database filename
// if filenames is omitted, the database is persisted in memory
func NewStackInstance(
	r *gin.Engine,
	stackPath string,
	// filesnames is an optional parameter for the name of the database
	filenames ...string) (
	stage *models.StageStruct,
	backRepo *orm.BackRepoStruct) {

	stage = models.NewStage(stackPath)

	if len(filenames) == 0 {
		filenames = append(filenames, ":memory:")
	}

	backRepo = orm.NewBackRepo(stage, filenames[0])

	controllers.GetController().AddBackRepo(backRepo, stackPath)

	controllers.Register(r)

	// add orchestration
	// insertion point
	models.SetOrchestratorOnAfterUpdate[models.AsSplit](stage)
	models.SetOrchestratorOnAfterUpdate[models.AsSplitArea](stage)
	models.SetOrchestratorOnAfterUpdate[models.Button](stage)
	models.SetOrchestratorOnAfterUpdate[models.Cursor](stage)
	models.SetOrchestratorOnAfterUpdate[models.Doc](stage)
	models.SetOrchestratorOnAfterUpdate[models.Form](stage)
	models.SetOrchestratorOnAfterUpdate[models.Load](stage)
	models.SetOrchestratorOnAfterUpdate[models.Slider](stage)
	models.SetOrchestratorOnAfterUpdate[models.Split](stage)
	models.SetOrchestratorOnAfterUpdate[models.Svg](stage)
	models.SetOrchestratorOnAfterUpdate[models.Table](stage)
	models.SetOrchestratorOnAfterUpdate[models.Tone](stage)
	models.SetOrchestratorOnAfterUpdate[models.Tree](stage)
	models.SetOrchestratorOnAfterUpdate[models.View](stage)

	return
}
