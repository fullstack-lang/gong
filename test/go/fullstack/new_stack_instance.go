package fullstack

import (
	"github.com/fullstack-lang/gong/test/go/controllers"
	"github.com/fullstack-lang/gong/test/go/models"
	"github.com/fullstack-lang/gong/test/go/orm"

	"github.com/gin-gonic/gin"

	// this will import the angular front end source code directory (versionned with git) in the vendor directory
	// this path will be included in the "tsconfig.json" front end compilation paths
	// to include this stack front end code
	_ "github.com/fullstack-lang/gong/test/ng/projects"
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
	stage *models.StageStruct) {

	// temporary
	if stackPath == "" {
		stage = models.GetDefaultStage()
	} else {
		stage = models.NewStage()
	}

	if len(filenames) == 0 {
		filenames = append(filenames, ":memory:")
	}

	backRepo := orm.NewBackRepo(stage, filenames[0])

	if stackPath != "" {
		controllers.GetController().AddBackRepo(backRepo, stackPath)
	}

	controllers.Register(r)

	return
}
