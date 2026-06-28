// do not modify, generated file
package level1stack

import (
	"fmt"
	"log"
	"strings"

	"github.com/fullstack-lang/gong/app/xsd/go/models"
	"github.com/fullstack-lang/gong/app/xsd/go/probe"

	embeddedgo "github.com/fullstack-lang/gong/app/xsd/go"

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

	stage.MarshallFile(fmt.Sprintf("./%s", filename), "github.com/fullstack-lang/gong/app/xsd/go/models", packageName)
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
	models.SetOrchestratorOnAfterUpdate[models.All](stage)
	models.SetOrchestratorOnAfterUpdate[models.Annotation](stage)
	models.SetOrchestratorOnAfterUpdate[models.Attribute](stage)
	models.SetOrchestratorOnAfterUpdate[models.AttributeGroup](stage)
	models.SetOrchestratorOnAfterUpdate[models.Choice](stage)
	models.SetOrchestratorOnAfterUpdate[models.ComplexContent](stage)
	models.SetOrchestratorOnAfterUpdate[models.ComplexType](stage)
	models.SetOrchestratorOnAfterUpdate[models.Documentation](stage)
	models.SetOrchestratorOnAfterUpdate[models.Element](stage)
	models.SetOrchestratorOnAfterUpdate[models.Enumeration](stage)
	models.SetOrchestratorOnAfterUpdate[models.Extension](stage)
	models.SetOrchestratorOnAfterUpdate[models.Group](stage)
	models.SetOrchestratorOnAfterUpdate[models.Length](stage)
	models.SetOrchestratorOnAfterUpdate[models.MaxInclusive](stage)
	models.SetOrchestratorOnAfterUpdate[models.MaxLength](stage)
	models.SetOrchestratorOnAfterUpdate[models.MinInclusive](stage)
	models.SetOrchestratorOnAfterUpdate[models.MinLength](stage)
	models.SetOrchestratorOnAfterUpdate[models.Pattern](stage)
	models.SetOrchestratorOnAfterUpdate[models.Restriction](stage)
	models.SetOrchestratorOnAfterUpdate[models.Schema](stage)
	models.SetOrchestratorOnAfterUpdate[models.Sequence](stage)
	models.SetOrchestratorOnAfterUpdate[models.SimpleContent](stage)
	models.SetOrchestratorOnAfterUpdate[models.SimpleType](stage)
	models.SetOrchestratorOnAfterUpdate[models.TotalDigit](stage)
	models.SetOrchestratorOnAfterUpdate[models.Union](stage)
	models.SetOrchestratorOnAfterUpdate[models.WhiteSpace](stage)

	return
}
