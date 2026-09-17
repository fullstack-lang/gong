// do not modify, generated file
package fullstack

import (
	"net/http"

	"github.com/fullstack-lang/gong/lib/svg/go/controllers"
	"github.com/fullstack-lang/gong/lib/svg/go/models"
	"github.com/fullstack-lang/gong/lib/svg/go/orm"

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
	r any,
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

	if mux, ok := r.(*http.ServeMux); ok {
		controllers.Register(mux)
	}

	// Attempt to register the WASM socket. 
    // On Mac/Linux, this does absolutely nothing.
    // On WASM, it safely grabs the BackRepo and registers the JS callback
    registerWasmSocket(stackPath, backRepo)

	// add orchestration
	// insertion point
	stage.SetOrchestratorOnAfterUpdate[models.Animate]()
	stage.SetOrchestratorOnAfterUpdate[models.Circle]()
	stage.SetOrchestratorOnAfterUpdate[models.Condition]()
	stage.SetOrchestratorOnAfterUpdate[models.ControlPoint]()
	stage.SetOrchestratorOnAfterUpdate[models.Ellipse]()
	stage.SetOrchestratorOnAfterUpdate[models.FileToDownload]()
	stage.SetOrchestratorOnAfterUpdate[models.Layer]()
	stage.SetOrchestratorOnAfterUpdate[models.Line]()
	stage.SetOrchestratorOnAfterUpdate[models.Link]()
	stage.SetOrchestratorOnAfterUpdate[models.LinkAnchoredPath]()
	stage.SetOrchestratorOnAfterUpdate[models.LinkAnchoredText]()
	stage.SetOrchestratorOnAfterUpdate[models.Path]()
	stage.SetOrchestratorOnAfterUpdate[models.Point]()
	stage.SetOrchestratorOnAfterUpdate[models.Polygone]()
	stage.SetOrchestratorOnAfterUpdate[models.Polyline]()
	stage.SetOrchestratorOnAfterUpdate[models.Rect]()
	stage.SetOrchestratorOnAfterUpdate[models.RectAnchoredPath]()
	stage.SetOrchestratorOnAfterUpdate[models.RectAnchoredPngImage]()
	stage.SetOrchestratorOnAfterUpdate[models.RectAnchoredRect]()
	stage.SetOrchestratorOnAfterUpdate[models.RectAnchoredText]()
	stage.SetOrchestratorOnAfterUpdate[models.RectLinkLink]()
	stage.SetOrchestratorOnAfterUpdate[models.SVG]()
	stage.SetOrchestratorOnAfterUpdate[models.SvgText]()
	stage.SetOrchestratorOnAfterUpdate[models.Text]()

	return
}
