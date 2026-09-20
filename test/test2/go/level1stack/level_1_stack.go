// do not modify, generated file
package level1stack

import (
	"fmt"
	"log"
	"strings"

	"github.com/fullstack-lang/gong/test/test2/go/models"
	"github.com/fullstack-lang/gong/test/test2/go/models/probe"

	embeddedgo "github.com/fullstack-lang/gong/test/test2/go"

	"net/http"

	split_static "github.com/fullstack-lang/gong/lib/split/go/static"
)

// hook marhalling to stage
type BeforeCommitImplementation struct {
	marshallOnCommit string

	packageName string
	stageSet *models.StageSet
}

func (impl *BeforeCommitImplementation) BeforeCommit(stage *models.Stage) {

	if stage.GetGongMarshallingMode() == models.GongMarshallingAppendCommit {
		stage.ComputeForwardAndBackwardCommits()
		stage.ComputeReferenceAndOrders()
	}

	// the ".go" is not provided
	filename := impl.marshallOnCommit
	if !strings.HasSuffix(filename, ".go") {
		filename = filename + ".go"
	}

	packageName := impl.packageName
	if packageName == "" {
		packageName = "main"
	}

	if impl.stageSet != nil {
		impl.stageSet.MarshallFile(fmt.Sprintf("./%s", filename), packageName)
	} else {
		stage.MarshallFile(fmt.Sprintf("./%s", filename), "github.com/fullstack-lang/gong/test/test2/go/models", packageName)
	}
}

type Level1Stack struct {
	Stage *models.Stage
	StageSet      *models.StageSet
	StageSetProbe *probe.StageSetProbe
	Probe *probe.Probe
	R     *http.ServeMux
}

func (stack *Level1Stack) Run(addr string) error {
	return split_static.RunServer(stack.R, addr)
}

func NewLevel1Stack(
	stackPath string,
	unmarshallFromCode string,
	marshallOnCommit string,
	withProbe bool,
	embeddedDiagrams bool,
) (level1Stack *Level1Stack) {
	return NewLevel1StackDelta(stackPath, unmarshallFromCode, marshallOnCommit, withProbe, embeddedDiagrams, false)
}

func NewLevel1StackDelta(
	stackPath string,
	unmarshallFromCode string,
	marshallOnCommit string,
	withProbe bool,
	embeddedDiagrams bool,
	deltaMode bool,
) (level1Stack *Level1Stack) {
	return newLevel1Stack(stackPath, unmarshallFromCode, marshallOnCommit, withProbe, embeddedDiagrams, deltaMode, false)
}

func NewLevel1StackStageSet(
	stackPath string,
	unmarshallFromCode string,
	marshallOnCommit string,
	withProbe bool,
	embeddedDiagrams bool,
) (level1Stack *Level1Stack) {
	return NewLevel1StackStageSetDelta(stackPath, unmarshallFromCode, marshallOnCommit, withProbe, embeddedDiagrams, false)
}

func NewLevel1StackStageSetDelta(
	stackPath string,
	unmarshallFromCode string,
	marshallOnCommit string,
	withProbe bool,
	embeddedDiagrams bool,
	deltaMode bool,
) (level1Stack *Level1Stack) {
	return newLevel1Stack(stackPath, unmarshallFromCode, marshallOnCommit, withProbe, embeddedDiagrams, deltaMode, true)
}

func newLevel1Stack(
	stackPath string,
	unmarshallFromCode string,
	marshallOnCommit string,
	withProbe bool,
	embeddedDiagrams bool,
	deltaMode bool,
	stageSetMode bool,
) (level1Stack *Level1Stack) {

	level1Stack = new(Level1Stack)
	stage := models.NewStage(stackPath)

	if deltaMode {
		stage.SetDeltaMode(true)
	}

	level1Stack.Stage = stage
	stageSet := models.NewStageSetFromStage(stage)
	level1Stack.StageSet = stageSet

	level1Stack.R = split_static.ServeStaticFiles(false)
	if withProbe {
		// if the application edits the diagrams via the probe, it is surmised
		// that the application is launched from "go/cmd/<appl>/". Therefore, to reach
		// "go/models/diagrams/diagrams.go", the path is "../../models/diagrams/diagrams.go"
		level1Stack.Probe = probe.NewProbe(
			level1Stack.R,
			embeddedgo.GoModelsDir,
			embeddedgo.GoDiagramsDir,
			embeddedDiagrams,
			stage,
		)

		stage.SetProbeIF(level1Stack.Probe)
		level1Stack.StageSetProbe = probe.NewStageSetProbe(
			level1Stack.R,
			stageSet,
		)
	}

	if unmarshallFromCode != "" {
		if stageSetMode {
			err := stageSet.ParseAstFile(unmarshallFromCode, true)

			// if the application is run with -unmarshallFromCode=xxx.go -marshallOnCommit
			// xxx.go might be absent the first time. However, this shall not be a show stopper.
			if err != nil {
				log.Println("no file to read " + err.Error())
			}

			stageSet.ComputeReverseMaps()
			stageSet.ComputeInstancesNb()
			stageSet.ComputeReferenceAndOrders()
		} else {
			err := stage.ParseAstFile(unmarshallFromCode, true)

			// if the application is run with -unmarshallFromCode=xxx.go -marshallOnCommit
			// xxx.go might be absent the first time. However, this shall not be a show stopper.
			if err != nil {
				log.Println("no file to read " + err.Error())
			}

			stage.ComputeReverseMaps()
			stage.ComputeInstancesNb()
			stage.ComputeReferenceAndOrders()
		}
	} else {
		// in case the database is used, checkout the content to the stage
		if stageSetMode {
			stageSet.Checkout()
		} else {
			stage.Checkout()
		}
	}

	// hook automatic marshall to go code at every commit
	if marshallOnCommit != "" {
		hook := new(BeforeCommitImplementation)
		hook.marshallOnCommit = marshallOnCommit
		if stageSetMode {
			hook.stageSet = stageSet
		}
		stage.OnInitCommitCallback = hook
	}

	// add orchestration
	// insertion point
	stage.SetOrchestratorOnAfterUpdate[models.A]()
	stage.SetOrchestratorOnAfterUpdate[models.B]()

	return
}
