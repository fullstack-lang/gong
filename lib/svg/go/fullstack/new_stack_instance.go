// do not modify, generated file
package fullstack

import (
	"github.com/fullstack-lang/gong/lib/svg/go/controllers"
	"github.com/fullstack-lang/gong/lib/svg/go/models"
	"github.com/fullstack-lang/gong/lib/svg/go/orm"

	"github.com/gin-gonic/gin"

	// this will import the angular front end source code directory (versionned with git) in the vendor directory
	// this path will be included in the "tsconfig.json" front end compilation paths
	// to include this stack front end code
	_ "github.com/fullstack-lang/gong/lib/svg/ng-github.com-fullstack-lang-gong-lib-svg"
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
	models.SetOrchestratorOnAfterUpdate[models.Animate](stage)
	models.SetOrchestratorOnAfterUpdate[models.Circle](stage)
	models.SetOrchestratorOnAfterUpdate[models.Ellipse](stage)
	models.SetOrchestratorOnAfterUpdate[models.Layer](stage)
	models.SetOrchestratorOnAfterUpdate[models.Line](stage)
	models.SetOrchestratorOnAfterUpdate[models.Link](stage)
	models.SetOrchestratorOnAfterUpdate[models.LinkAnchoredText](stage)
	models.SetOrchestratorOnAfterUpdate[models.Path](stage)
	models.SetOrchestratorOnAfterUpdate[models.Point](stage)
	models.SetOrchestratorOnAfterUpdate[models.Polygone](stage)
	models.SetOrchestratorOnAfterUpdate[models.Polyline](stage)
	models.SetOrchestratorOnAfterUpdate[models.Rect](stage)
	models.SetOrchestratorOnAfterUpdate[models.RectAnchoredPath](stage)
	models.SetOrchestratorOnAfterUpdate[models.RectAnchoredRect](stage)
	models.SetOrchestratorOnAfterUpdate[models.RectAnchoredText](stage)
	models.SetOrchestratorOnAfterUpdate[models.RectLinkLink](stage)
	models.SetOrchestratorOnAfterUpdate[models.SVG](stage)
	models.SetOrchestratorOnAfterUpdate[models.SvgText](stage)
	models.SetOrchestratorOnAfterUpdate[models.Text](stage)

	return
}
