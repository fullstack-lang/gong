// do not modify, generated file
package stack

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/fullstack-lang/gong/lib/probe/go/fullstack"
	"github.com/fullstack-lang/gong/lib/probe/go/models"
	"github.com/fullstack-lang/gong/lib/probe/go/orm"
	"github.com/fullstack-lang/gong/lib/probe/go/probe"

	probe_go "github.com/fullstack-lang/gong/lib/probe/go"

	"github.com/gin-gonic/gin"
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

	stage.Checkout()
	stage.Marshall(file, "github.com/fullstack-lang/gong/lib/probe/go/models", packageName)
}

type Stack struct {
	Probe    *probe.Probe
	Stage    *models.Stage
	BackRepo *orm.BackRepoStruct

	hook *BeforeCommitImplementation
}

// NewStack initializes and configures a new stack instance for a full-stack application.
// It sets up the backend repository, provides options for unmarshalling from Go code,
// automatic marshalling on commits, and initializing a probe for monitoring and visualization.
// The function returns a pointer to the initialized Stage.
//
// Parameters:
//   - r *gin.Engine: A Gin engine instance for handling HTTP requests.
//   - stackPath string: A unique identifier for the stack instance, used for differentiating
//     between different instances if needed.
//   - unmarshallFromCode string: File path of a Go source code file. If provided, the function
//     will attempt to unmarshall the stage data from this file.
//   - marshallOnCommit string: If provided, sets up an automatic marshalling hook that writes
//     the stage data to a Go source code file with this name on every commit.
//   - dbFileName string: Name of the database file for persisting the stage data. If an empty
//     string is provided, the database will be persisted in memory.
//   - embeddedDiagrams bool: Flag indicating whether to embed diagrams in the probe.
//   - withProbe bool: If true, initializes a probe for monitoring and visualization.
//
// Returns:
// - *Stack: Pointer to the initialized stack instance.
//
// Behavior:
//  1. Initialize Stage and Backend Repository: Creates a new stage instance and a backend
//     repository. If dbFileName is provided, it is used as the database file name; otherwise,
//     the database is persisted in memory.
//  2. Optional Unmarshalling from Go Code: If unmarshallFromCode is provided, attempts to
//     unmarshall the stage data from the specified Go source code file. If the file is not
//     found, a log message is printed, but the function does not terminate.
//  3. Automatic Marshalling on Commit: If marshallOnCommit is provided, sets up a hook that
//     automatically marshals the stage data to a Go source code file with the specified name
//     on every commit.
//  4. Initialize Probe (Optional): If withProbe is true, initializes a probe for monitoring
//     and visualization.
//  5. Configure Orchestration: Configures orchestration for various model structures.
func NewStack(
	r *gin.Engine,
	stackPath string,
	unmarshallFromCode string,
	marshallOnCommit string,
	dbFileName string,
	embeddedDiagrams bool,
	withProbe bool) (
	stack *Stack) {

	stack = new(Stack)

	var backRepo *orm.BackRepoStruct
	var stage *models.Stage

	if dbFileName == "" {
		stage, backRepo = fullstack.NewStackInstance(r, stackPath)
	} else {
		stage, backRepo = fullstack.NewStackInstance(r, stackPath, dbFileName)
	}

	stack.Stage = stage
	stack.BackRepo = backRepo

	if unmarshallFromCode != "" {
		stage.Checkout()
		stage.Reset()
		stage.Commit()
		err := models.ParseAstFile(stage, unmarshallFromCode)

		// if the application is run with -unmarshallFromCode=xxx.go -marshallOnCommit
		// xxx.go might be absent the first time. However, this shall not be a show stopper.
		if err != nil {
			log.Println("no file to read " + err.Error())
		}

		stage.Commit()
	} else {
		// in case the database is used, checkout the content to the stage
		stage.Checkout()
	}

	// hook automatic marshall to go code at every commit
	if marshallOnCommit != "" {
		stack.hook = new(BeforeCommitImplementation)
		stack.hook.marshallOnCommit = marshallOnCommit
		stage.OnInitCommitCallback = stack.hook
	}

	if withProbe {
		if embeddedDiagrams {
			log.Panic("load of embedded diagram with doc2 not yet implemented")
		} else {
			// if the application edits the diagrams via the probe, it is surmised
			// that the application is launched from "go/cmd/<appl>/". Therefore, to reach
			// "go/diagrams/diagrams.go", the path is "../../diagrams/diagrams.go"
			stack.Probe = probe.NewProbe(r, probe_go.GoModelsDir, probe_go.GoDiagramsDir,
				embeddedDiagrams, stage, "../../diagrams/diagrams.go")
		}

	}

	return
}

func NewTranscientStack(r *gin.Engine, stackPath string, withProbe bool) (stack *Stack) {

	return NewStack(r, stackPath, "", "", "", true, withProbe)
}

// SetMarshallPackageName overrides the "main" package default generated
func (stack *Stack) SetMarshallPackageName(packageName string) {
	stack.hook.packageName = packageName
}
