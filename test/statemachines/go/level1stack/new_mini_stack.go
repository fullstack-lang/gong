package level1stack

import (
	"fmt"
	"log"
	"os"
	"strings"

	statemachines_go "github.com/fullstack-lang/gong/test/statemachines/go"
	"github.com/fullstack-lang/gong/test/statemachines/go/models"
	"github.com/fullstack-lang/gong/test/statemachines/go/probe"
	"github.com/gin-gonic/gin"

	split_static "github.com/fullstack-lang/gong/lib/split/go/static"
)

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
		err := models.ParseAstFile(stage, unmarshallFromCode)

		// if the application is run with -unmarshallFromCode=xxx.go -marshallOnCommit
		// xxx.go might be absent the first time. However, this shall not be a show stopper.
		if err != nil {
			log.Println("no file to read " + err.Error())
		}

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
			statemachines_go.GoModelsDir,
			statemachines_go.GoDiagramsDir,
			embeddedDiagrams,
			stage,
		)
	}

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

	stage.Marshall(file, "github.com/fullstack-lang/gong/test/statemachines/go/models", packageName)
}
