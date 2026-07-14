// do not modify, generated file
package fullstack

import (
	"github.com/fullstack-lang/gong/lib/threejs/go/controllers"
	"github.com/fullstack-lang/gong/lib/threejs/go/models"
	"github.com/fullstack-lang/gong/lib/threejs/go/orm"

	"github.com/gin-gonic/gin"

	// this will import the angular front end source code directory (versionned with git) in the vendor directory
	// this path will be included in the "tsconfig.json" front end compilation paths
	// to include this stack front end code
	_ "github.com/fullstack-lang/gong/lib/threejs/ng-github.com-fullstack-lang-gong-lib-threejs"
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

	// Attempt to register the WASM socket. 
    // On Mac/Linux, this does absolutely nothing.
    // On WASM, it safely grabs the BackRepo and registers the JS callback
    registerWasmSocket(stackPath, backRepo)

	// add orchestration
	// insertion point
	models.SetOrchestratorOnAfterUpdate[models.AmbiantLight](stage)
	models.SetOrchestratorOnAfterUpdate[models.BoxGeometry](stage)
	models.SetOrchestratorOnAfterUpdate[models.BufferGeometry](stage)
	models.SetOrchestratorOnAfterUpdate[models.Camera](stage)
	models.SetOrchestratorOnAfterUpdate[models.Canvas](stage)
	models.SetOrchestratorOnAfterUpdate[models.Curve](stage)
	models.SetOrchestratorOnAfterUpdate[models.CylinderGeometry](stage)
	models.SetOrchestratorOnAfterUpdate[models.DirectionalLight](stage)
	models.SetOrchestratorOnAfterUpdate[models.ExtrudeGeometry](stage)
	models.SetOrchestratorOnAfterUpdate[models.Mesh](stage)
	models.SetOrchestratorOnAfterUpdate[models.MeshMaterialBasic](stage)
	models.SetOrchestratorOnAfterUpdate[models.MeshPhysicalMaterial](stage)
	models.SetOrchestratorOnAfterUpdate[models.PlaneGeometry](stage)
	models.SetOrchestratorOnAfterUpdate[models.Shape](stage)
	models.SetOrchestratorOnAfterUpdate[models.SphereGeometry](stage)
	models.SetOrchestratorOnAfterUpdate[models.TorusGeometry](stage)
	models.SetOrchestratorOnAfterUpdate[models.Triangle](stage)
	models.SetOrchestratorOnAfterUpdate[models.TubeGeometry](stage)
	models.SetOrchestratorOnAfterUpdate[models.Vector2](stage)
	models.SetOrchestratorOnAfterUpdate[models.Vector3](stage)

	return
}
