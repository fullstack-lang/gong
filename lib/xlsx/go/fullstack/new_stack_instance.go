// do not modify, generated file
package fullstack

import (
	"github.com/fullstack-lang/gong/lib/xlsx/go/controllers"
	"github.com/fullstack-lang/gong/lib/xlsx/go/models"
	"github.com/fullstack-lang/gong/lib/xlsx/go/orm"

	"github.com/gin-gonic/gin"

	// this will import the angular front end source code directory (versionned with git) in the vendor directory
	// this path will be included in the "tsconfig.json" front end compilation paths
	// to include this stack front end code
	// This is a level 1 gong application, no need to import the angular code
	// therefore, the following line that is necessary in level 2 applications, is commented
	// _ "github.com/fullstack-lang/gong/lib/xlsx/ng-github.com-fullstack-lang-gong-lib-xlsx"
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
	models.SetOrchestratorOnAfterUpdate[models.DisplaySelection](stage)
	models.SetOrchestratorOnAfterUpdateWithMouseEvent[models.DisplaySelection](stage)
	models.SetOrchestratorOnAfterUpdate[models.XLCell](stage)
	models.SetOrchestratorOnAfterUpdateWithMouseEvent[models.XLCell](stage)
	models.SetOrchestratorOnAfterUpdate[models.XLFile](stage)
	models.SetOrchestratorOnAfterUpdateWithMouseEvent[models.XLFile](stage)
	models.SetOrchestratorOnAfterUpdate[models.XLRow](stage)
	models.SetOrchestratorOnAfterUpdateWithMouseEvent[models.XLRow](stage)
	models.SetOrchestratorOnAfterUpdate[models.XLSheet](stage)
	models.SetOrchestratorOnAfterUpdateWithMouseEvent[models.XLSheet](stage)

	return
}
