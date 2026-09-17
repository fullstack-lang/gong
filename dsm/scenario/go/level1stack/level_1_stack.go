// do not modify, generated file
package level1stack

import (
	"fmt"
	"log"
	"strings"

	"github.com/fullstack-lang/gong/dsm/scenario/go/models"
	"github.com/fullstack-lang/gong/dsm/scenario/go/models/probe"

	embeddedgo "github.com/fullstack-lang/gong/dsm/scenario/go"

	"net/http"

	split_static "github.com/fullstack-lang/gong/lib/split/go/static"
)

// hook marhalling to stage
type BeforeCommitImplementation struct {
	marshallOnCommit string

	packageName string
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

	stage.MarshallFile(fmt.Sprintf("./%s", filename), "github.com/fullstack-lang/gong/dsm/scenario/go/models", packageName)
}

type Level1Stack struct {
	Stage *models.Stage
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

	level1Stack = new(Level1Stack)
	stage := models.NewStage(stackPath)

	if deltaMode {
		stage.SetDeltaMode(true)
	}

	level1Stack.Stage = stage

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
	}

	if unmarshallFromCode != "" {
		err := stage.ParseAstFile(unmarshallFromCode, true)

		// if the application is run with -unmarshallFromCode=xxx.go -marshallOnCommit
		// xxx.go might be absent the first time. However, this shall not be a show stopper.
		if err != nil {
			log.Println("no file to read " + err.Error())
		}

		stage.ComputeReverseMaps()
		stage.ComputeInstancesNb()
		stage.ComputeReferenceAndOrders()
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

	// add orchestration
	// insertion point
	stage.SetOrchestratorOnAfterUpdate[models.ActorState]()
	stage.SetOrchestratorOnAfterUpdate[models.ActorStateShape]()
	stage.SetOrchestratorOnAfterUpdate[models.ActorStateTransition]()
	stage.SetOrchestratorOnAfterUpdate[models.ActorStateTransitionShape]()
	stage.SetOrchestratorOnAfterUpdate[models.Analysis]()
	stage.SetOrchestratorOnAfterUpdate[models.ControlPointShape]()
	stage.SetOrchestratorOnAfterUpdate[models.Diagram]()
	stage.SetOrchestratorOnAfterUpdate[models.Document]()
	stage.SetOrchestratorOnAfterUpdate[models.DocumentUse]()
	stage.SetOrchestratorOnAfterUpdate[models.EvolutionDirection]()
	stage.SetOrchestratorOnAfterUpdate[models.EvolutionDirectionShape]()
	stage.SetOrchestratorOnAfterUpdate[models.Foo]()
	stage.SetOrchestratorOnAfterUpdate[models.GeoObject]()
	stage.SetOrchestratorOnAfterUpdate[models.GeoObjectUse]()
	stage.SetOrchestratorOnAfterUpdate[models.Group]()
	stage.SetOrchestratorOnAfterUpdate[models.GroupUse]()
	stage.SetOrchestratorOnAfterUpdate[models.Library]()
	stage.SetOrchestratorOnAfterUpdate[models.MapObject]()
	stage.SetOrchestratorOnAfterUpdate[models.MapObjectUse]()
	stage.SetOrchestratorOnAfterUpdate[models.Parameter]()
	stage.SetOrchestratorOnAfterUpdate[models.ParameterCategory]()
	stage.SetOrchestratorOnAfterUpdate[models.ParameterCategoryUse]()
	stage.SetOrchestratorOnAfterUpdate[models.ParameterShape]()
	stage.SetOrchestratorOnAfterUpdate[models.ParametersAggregate]()
	stage.SetOrchestratorOnAfterUpdate[models.ParametersAggregateShape]()
	stage.SetOrchestratorOnAfterUpdate[models.Position]()
	stage.SetOrchestratorOnAfterUpdate[models.Repository]()
	stage.SetOrchestratorOnAfterUpdate[models.Scenario]()
	stage.SetOrchestratorOnAfterUpdate[models.User]()
	stage.SetOrchestratorOnAfterUpdate[models.UserUse]()
	stage.SetOrchestratorOnAfterUpdate[models.Workspace]()

	return
}
