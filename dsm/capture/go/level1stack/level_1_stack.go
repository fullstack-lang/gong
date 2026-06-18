// do not modify, generated file
package level1stack

import (
	"fmt"
	"log"
	"strings"

	"github.com/fullstack-lang/gong/dsm/capture/go/models"
	"github.com/fullstack-lang/gong/dsm/capture/go/probe"

	embeddedgo "github.com/fullstack-lang/gong/dsm/capture/go"

	"github.com/gin-gonic/gin"

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

	stage.MarshallFile(fmt.Sprintf("./%s", filename), "github.com/fullstack-lang/gong/dsm/capture/go/models", packageName)
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
		// "go/diagrams/diagrams.go", the path is "../../diagrams/diagrams.go"
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
		err := models.ParseAstFile(stage, unmarshallFromCode, true)

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
	models.SetOrchestratorOnAfterUpdate[models.AnalysisNeed](stage)
	models.SetOrchestratorOnAfterUpdate[models.Concept](stage)
	models.SetOrchestratorOnAfterUpdate[models.ConceptShape](stage)
	models.SetOrchestratorOnAfterUpdate[models.Concern](stage)
	models.SetOrchestratorOnAfterUpdate[models.ConcernCompositionShape](stage)
	models.SetOrchestratorOnAfterUpdate[models.ConcernInputShape](stage)
	models.SetOrchestratorOnAfterUpdate[models.ConcernOutputShape](stage)
	models.SetOrchestratorOnAfterUpdate[models.ConcernShape](stage)
	models.SetOrchestratorOnAfterUpdate[models.Deliverable](stage)
	models.SetOrchestratorOnAfterUpdate[models.DeliverableCompositionShape](stage)
	models.SetOrchestratorOnAfterUpdate[models.DeliverableConceptShape](stage)
	models.SetOrchestratorOnAfterUpdate[models.DeliverableShape](stage)
	models.SetOrchestratorOnAfterUpdate[models.Diagram](stage)
	models.SetOrchestratorOnAfterUpdate[models.Library](stage)
	models.SetOrchestratorOnAfterUpdate[models.Note](stage)
	models.SetOrchestratorOnAfterUpdate[models.NoteDeliverableShape](stage)
	models.SetOrchestratorOnAfterUpdate[models.NoteShape](stage)
	models.SetOrchestratorOnAfterUpdate[models.NoteStakeholderShape](stage)
	models.SetOrchestratorOnAfterUpdate[models.NoteTaskShape](stage)
	models.SetOrchestratorOnAfterUpdate[models.Requirement](stage)
	models.SetOrchestratorOnAfterUpdate[models.RequirementShape](stage)
	models.SetOrchestratorOnAfterUpdate[models.Stakeholder](stage)
	models.SetOrchestratorOnAfterUpdate[models.StakeholderCompositionShape](stage)
	models.SetOrchestratorOnAfterUpdate[models.StakeholderConcernShape](stage)
	models.SetOrchestratorOnAfterUpdate[models.StakeholderShape](stage)
	models.SetOrchestratorOnAfterUpdate[models.SupportLevel](stage)
	models.SetOrchestratorOnAfterUpdate[models.Tool](stage)

	return
}
