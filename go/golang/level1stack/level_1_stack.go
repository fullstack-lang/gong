package level1stack

const DebouncedMarshallingLevel1StackInstanceTemplate = `// do not modify, generated file
package stack

import (
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"time"

	"{{PkgPathRoot}}/models"
	"{{PkgPathRoot}}/probe"

	{{pkgname}}_go "{{PkgPathRoot}}"

	"github.com/gin-gonic/gin"

	split_static "github.com/fullstack-lang/gong/lib/split/go/static"
)

// hook marhalling to stage
type BeforeCommitImplementation struct {
	marshallOnCommit string

	packageName string

	mu    sync.Mutex
	timer *time.Timer
}

const debounceDuration = 2 * time.Second

func (impl *BeforeCommitImplementation) BeforeCommit(stage *models.Stage) {
	impl.mu.Lock()
	defer impl.mu.Unlock()

	// If a timer is already running, stop it.
	if impl.timer != nil {
		impl.timer.Stop()
	}

	// Start a new timer. When it fires, it will execute performMarshalling
	// in a new goroutine.
	impl.timer = time.AfterFunc(debounceDuration, func() {
		go impl.performMarshalling(stage)
	})
}

func (impl *BeforeCommitImplementation) performMarshalling(stage *models.Stage) {
` + stackInstanceTemplateEpilogue

const BlockingMarshallingLevel1StackInstanceTemplate = `// do not modify, generated file
package level1stack

import (
	"fmt"
	"log"
	"os"
	"strings"

	"{{PkgPathRoot}}/models"
	"{{PkgPathRoot}}/probe"

	{{pkgname}}_go "{{PkgPathRoot}}"

	"github.com/gin-gonic/gin"

	split_static "github.com/fullstack-lang/gong/lib/split/go/static"
)

// hook marhalling to stage
type BeforeCommitImplementation struct {
	marshallOnCommit string

	packageName string
}

func (impl *BeforeCommitImplementation) BeforeCommit(stage *models.Stage) {
` + stackInstanceTemplateEpilogue

const stackInstanceTemplateEpilogue = `
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

	stage.Marshall(file, "{{PkgPathRoot}}/models", packageName)
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

	level1Stack = new(Level1Stack)
	stage := models.NewStage(stackPath)
	level1Stack.Stage = stage

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

	level1Stack.R = split_static.ServeStaticFiles(false)
	if withProbe {
		// if the application edits the diagrams via the probe, it is surmised
		// that the application is launched from "go/cmd/<appl>/". Therefore, to reach
		// "go/diagrams/diagrams.go", the path is "../../diagrams/diagrams.go"
		level1Stack.Probe = probe.NewProbe(
			level1Stack.R,
			{{pkgname}}_go.GoModelsDir,
			{{pkgname}}_go.GoDiagramsDir,
			embeddedDiagrams,
			stage,
		)

		stage.SetProbeIF(level1Stack.Probe)
	}

	// add orchestration
	// insertion point{{` + string(rune(ModelGongNLevel1tackInstanceSet)) + `}}

	return
}
`

type ModelGongNLevel1tackInstanceStructInsertionId int

const (
	ModelGongNLevel1tackInstanceSet ModelGongNLevel1tackInstanceStructInsertionId = iota
)

var ModelGongNLevel1tackInstanceStructSubTemplateCode map[string]string = // new line
map[string]string{
	string(rune(ModelGongNLevel1tackInstanceSet)): `
	models.SetOrchestratorOnAfterUpdate[models.{{Structname}}](stage)`,
}
