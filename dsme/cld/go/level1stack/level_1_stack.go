// do not modify, generated file
package level1stack

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/fullstack-lang/gong/dsme/cld/go/models"
	"github.com/fullstack-lang/gong/dsme/cld/go/probe"

	cld_go "github.com/fullstack-lang/gong/dsme/cld/go"

	"github.com/gin-gonic/gin"

	split_static "github.com/fullstack-lang/gong/lib/split/go/static"
)

// hook marhalling to stage
type BeforeCommitImplementation struct {
	marshallOnCommit string

	packageName string
}

func (impl *BeforeCommitImplementation) BeforeCommit(stage *models.Stage) {

	// the ".go" is not provided
	filename := impl.marshallOnCommit
	if !strings.HasSuffix(filename, ".go") {
		filename = filename + ".go"
	}

	file, err := os.Create(fmt.Sprintf("./%s", filename))
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()

	packageName := impl.packageName
	if packageName == "" {
		packageName = "main"
	}

	stage.Marshall(file, "github.com/fullstack-lang/gong/dsme/cld/go/models", packageName)
}

type Level1Stack struct {
	Stage *models.Stage
	Probe *probe.Probe
	R     *gin.Engine
}

func NewLevel1Stack(
	stackPath string,
	unmarshallFromCode string,
	marshallOnCommit string,
	withProbe bool,
	embeddedDiagrams bool,
) (miniStack *Level1Stack) {

	miniStack = new(Level1Stack)
	stage := models.NewStage(stackPath)
	miniStack.Stage = stage

	if unmarshallFromCode != "" {
		err := models.ParseAstFile(stage, unmarshallFromCode, true)

		// if the application is run with -unmarshallFromCode=xxx.go -marshallOnCommit
		// xxx.go might be absent the first time. However, this shall not be a show stopper.
		if err != nil {
			log.Println("no file to read " + err.Error())
		}

		stage.ComputeReverseMaps()
		stage.ComputeInstancesNb()
	} else {
		// in case the database is used, checkout the content to the stage
		stage.Checkout()
	}

	// hook automatic marshall to go code at every commit
	if marshallOnCommit != "" {
		hook := new(BeforeCommitImplementation)
		hook.marshallOnCommit = marshallOnCommit
		stage.OnInitCommitCallback = hook
	}

	miniStack.R = split_static.ServeStaticFiles(false)
	if withProbe {
		// if the application edits the diagrams via the probe, it is surmised
		// that the application is launched from "go/cmd/<appl>/". Therefore, to reach
		// "go/diagrams/diagrams.go", the path is "../../diagrams/diagrams.go"
		miniStack.Probe = probe.NewProbe(
			miniStack.R,
			cld_go.GoModelsDir,
			cld_go.GoDiagramsDir,
			embeddedDiagrams,
			stage,
		)
	}

	// add orchestration
	// insertion point
	models.SetOrchestratorOnAfterUpdate[models.Category1](stage)
	models.SetOrchestratorOnAfterUpdate[models.Category1Shape](stage)
	models.SetOrchestratorOnAfterUpdate[models.Category2](stage)
	models.SetOrchestratorOnAfterUpdate[models.Category2Shape](stage)
	models.SetOrchestratorOnAfterUpdate[models.Category3](stage)
	models.SetOrchestratorOnAfterUpdate[models.Category3Shape](stage)
	models.SetOrchestratorOnAfterUpdate[models.ControlPointShape](stage)
	models.SetOrchestratorOnAfterUpdate[models.Desk](stage)
	models.SetOrchestratorOnAfterUpdate[models.Diagram](stage)
	models.SetOrchestratorOnAfterUpdate[models.Influence](stage)
	models.SetOrchestratorOnAfterUpdate[models.InfluenceShape](stage)

	return
}
