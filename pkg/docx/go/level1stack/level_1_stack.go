// do not modify, generated file
package level1stack

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/fullstack-lang/gong/pkg/docx/go/models"
	"github.com/fullstack-lang/gong/pkg/docx/go/probe"

	docx_go "github.com/fullstack-lang/gong/pkg/docx/go"

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

	stage.Marshall(file, "github.com/fullstack-lang/gong/pkg/docx/go/models", packageName)
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
			docx_go.GoModelsDir,
			docx_go.GoDiagramsDir,
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
		stage.ComputeReference()
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
	models.SetOrchestratorOnAfterUpdate[models.Body](stage)
	models.SetOrchestratorOnAfterUpdate[models.Document](stage)
	models.SetOrchestratorOnAfterUpdate[models.Docx](stage)
	models.SetOrchestratorOnAfterUpdate[models.File](stage)
	models.SetOrchestratorOnAfterUpdate[models.Node](stage)
	models.SetOrchestratorOnAfterUpdate[models.Paragraph](stage)
	models.SetOrchestratorOnAfterUpdate[models.ParagraphProperties](stage)
	models.SetOrchestratorOnAfterUpdate[models.ParagraphStyle](stage)
	models.SetOrchestratorOnAfterUpdate[models.Rune](stage)
	models.SetOrchestratorOnAfterUpdate[models.RuneProperties](stage)
	models.SetOrchestratorOnAfterUpdate[models.Table](stage)
	models.SetOrchestratorOnAfterUpdate[models.TableColumn](stage)
	models.SetOrchestratorOnAfterUpdate[models.TableProperties](stage)
	models.SetOrchestratorOnAfterUpdate[models.TableRow](stage)
	models.SetOrchestratorOnAfterUpdate[models.TableStyle](stage)
	models.SetOrchestratorOnAfterUpdate[models.Text](stage)

	return
}
