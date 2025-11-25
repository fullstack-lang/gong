// do not modify, generated file
package fullstack

import (
	"github.com/fullstack-lang/gong/test/statemachines/go/controllers"
	"github.com/fullstack-lang/gong/test/statemachines/go/models"
	"github.com/fullstack-lang/gong/test/statemachines/go/orm"

	"github.com/gin-gonic/gin"

	// this will import the angular front end source code directory (versionned with git) in the vendor directory
	// this path will be included in the "tsconfig.json" front end compilation paths
	// to include this stack front end code
	// This is a level 1 gong application, no need to import the angular code
	// therefore, the following line that is necessary in level 2 applications, is commented
	// _ "github.com/fullstack-lang/gong/test/statemachines/ng-github.com-fullstack-lang-gong-test-statemachines"
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
	models.SetOrchestratorOnAfterUpdate[models.Architecture](stage)
	models.SetOrchestratorOnAfterUpdate[models.Diagram](stage)
	models.SetOrchestratorOnAfterUpdate[models.Kill](stage)
	models.SetOrchestratorOnAfterUpdate[models.Message](stage)
	models.SetOrchestratorOnAfterUpdate[models.MessageType](stage)
	models.SetOrchestratorOnAfterUpdate[models.Object](stage)
	models.SetOrchestratorOnAfterUpdate[models.Role](stage)
	models.SetOrchestratorOnAfterUpdate[models.State](stage)
	models.SetOrchestratorOnAfterUpdate[models.StateMachine](stage)
	models.SetOrchestratorOnAfterUpdate[models.StateShape](stage)
	models.SetOrchestratorOnAfterUpdate[models.Transition](stage)
	models.SetOrchestratorOnAfterUpdate[models.Transition_Shape](stage)

	return
}
