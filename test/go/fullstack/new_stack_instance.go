// do not modify, generated file
package fullstack

import (
	"fmt"
	"log"
	"os"

	"github.com/fullstack-lang/gong/test/go/controllers"
	"github.com/fullstack-lang/gong/test/go/models"
	"github.com/fullstack-lang/gong/test/go/orm"
	"github.com/fullstack-lang/gong/test/go/probe"

	test_go "github.com/fullstack-lang/gong/test/go"

	"github.com/gin-gonic/gin"

	// this will import the angular front end source code directory (versionned with git) in the vendor directory
	// this path will be included in the "tsconfig.json" front end compilation paths
	// to include this stack front end code
	_ "github.com/fullstack-lang/gong/test/ng/projects"
)

// hook marhalling to stage
type BeforeCommitImplementation struct {
	marshallOnCommit string
}

func (impl *BeforeCommitImplementation) BeforeCommit(stage *models.StageStruct) {
	file, err := os.Create(fmt.Sprintf("./%s.go", impl.marshallOnCommit))
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()

	stage.Checkout()
	stage.Marshall(file, "github.com/fullstack-lang/gong/test/go/models", "main")
}

// NewStackInstance creates a new stack instance from the Stack Model
// and returns the backRepo of the stack instance (you can get the stage from backRepo.GetStage()
//
// - the stackPath is the unique identifier of the stack
// - the optional parameter filenames is for the name of the database filename
// if filenames is omitted, the database is persisted in memory
func NewStackInstance(
	r *gin.Engine,
	stackPath string,
	// filesnames is an optional parameter for the name of the database
	filenames ...string) (
	stage *models.StageStruct,
	backRepo *orm.BackRepoStruct) {

	stage = models.NewStage(stackPath)

	if len(filenames) == 0 {
		filenames = append(filenames, ":memory:")
	}

	backRepo = orm.NewBackRepo(stage, filenames[0])

	if stackPath != "" {
		controllers.GetController().AddBackRepo(backRepo, stackPath)
	}

	controllers.Register(r)

	// add orchestration
	// insertion point
	models.SetOrchestratorOnAfterUpdate[models.Astruct](stage)
	models.SetOrchestratorOnAfterUpdate[models.AstructBstruct2Use](stage)
	models.SetOrchestratorOnAfterUpdate[models.AstructBstructUse](stage)
	models.SetOrchestratorOnAfterUpdate[models.Bstruct](stage)
	models.SetOrchestratorOnAfterUpdate[models.Dstruct](stage)
	models.SetOrchestratorOnAfterUpdate[models.Fstruct](stage)

	return
}

// NewStage initializes and configures a new stage instance for a full-stack application.
// It sets up the backend repository, provides options for unmarshalling from Go code,
// automatic marshalling on commits, and initializing a probe for monitoring and visualization.
// The function returns a pointer to the initialized StageStruct.
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
// - *models.StageStruct: Pointer to the initialized stage instance.
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
func NewStage(
	r *gin.Engine,
	stackPath string,
	unmarshallFromCode string,
	marshallOnCommit string,
	dbFileName string,
	embeddedDiagrams bool,
	withProbe bool) (
	stage *models.StageStruct) {

	var backRepo *orm.BackRepoStruct

	if dbFileName == "" {
		stage, backRepo = NewStackInstance(r, stackPath)
	} else {
		stage, backRepo = NewStackInstance(r, stackPath, dbFileName)
	}

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
		hook := new(BeforeCommitImplementation)
		hook.marshallOnCommit = marshallOnCommit
		stage.OnInitCommitCallback = hook
	}

	if withProbe {
		probe.NewProbe(r, test_go.GoModelsDir, test_go.GoDiagramsDir,
			embeddedDiagrams, stackPath, stage, backRepo)
	}

	return
}
