// do not modify, generated file
package fullstack

import (
	"github.com/fullstack-lang/gong/lib/gantt/go/controllers"
	"github.com/fullstack-lang/gong/lib/gantt/go/models"
	"github.com/fullstack-lang/gong/lib/gantt/go/orm"

	"github.com/gin-gonic/gin"

	// this will import the angular front end source code directory (versionned with git) in the vendor directory
	// this path will be included in the "tsconfig.json" front end compilation paths
	// to include this stack front end code
	// This is a level 1 gong application, no need to import the angular code
	// therefore, the following line that is necessary in level 2 applications, is commented
	// _ "github.com/fullstack-lang/gong/lib/gantt/ng-github.com-fullstack-lang-gong-lib-gantt"
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
	stage *models.Stage,
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
	models.SetOrchestratorOnAfterUpdate[models.Arrow](stage)
	models.SetOrchestratorOnAfterUpdateWithMouseEvent[models.Arrow](stage)
	models.SetOrchestratorOnAfterUpdate[models.Bar](stage)
	models.SetOrchestratorOnAfterUpdateWithMouseEvent[models.Bar](stage)
	models.SetOrchestratorOnAfterUpdate[models.Gantt](stage)
	models.SetOrchestratorOnAfterUpdateWithMouseEvent[models.Gantt](stage)
	models.SetOrchestratorOnAfterUpdate[models.Group](stage)
	models.SetOrchestratorOnAfterUpdateWithMouseEvent[models.Group](stage)
	models.SetOrchestratorOnAfterUpdate[models.Lane](stage)
	models.SetOrchestratorOnAfterUpdateWithMouseEvent[models.Lane](stage)
	models.SetOrchestratorOnAfterUpdate[models.LaneUse](stage)
	models.SetOrchestratorOnAfterUpdateWithMouseEvent[models.LaneUse](stage)
	models.SetOrchestratorOnAfterUpdate[models.Milestone](stage)
	models.SetOrchestratorOnAfterUpdateWithMouseEvent[models.Milestone](stage)

	return
}
