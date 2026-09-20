package level1stack

import "fmt"

func GetLevel1StackTemplate(useSplitlite bool, hasStageSet bool) string {
	staticPkg := "split_static"
	staticImport := `split_static "github.com/fullstack-lang/gong/lib/split/go/static"`
	if useSplitlite {
		staticPkg = "splitlite_static"
		staticImport = `splitlite_static "github.com/fullstack-lang/gong/lib/splitlite/go/static"`
	}

	beforeCommitStructField := ""
	beforeCommitMarshall := `	stage.MarshallFile(fmt.Sprintf("./%s", filename), "{{PkgPathRoot}}/models", packageName)`
	level1StackFields := ""
	constructorsAndImpl := fmt.Sprintf(`func NewLevel1Stack(
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
	level1Stack.R = %s.ServeStaticFiles(false)
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
	// insertion point{{`+string(rune(ModelGongNLevel1tackInstanceSet))+`}}

	return
}`, staticPkg)

	if hasStageSet {
		beforeCommitStructField = "\n\tstageSet *models.StageSet"
		beforeCommitMarshall = `	if impl.stageSet != nil {
		impl.stageSet.MarshallFile(fmt.Sprintf("./%s", filename), packageName)
	} else {
		stage.MarshallFile(fmt.Sprintf("./%s", filename), "{{PkgPathRoot}}/models", packageName)
	}`
		level1StackFields = `
	StageSet      *models.StageSet
	StageSetProbe *probe.StageSetProbe`
		constructorsAndImpl = fmt.Sprintf(`func NewLevel1Stack(
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

	level1Stack.R = %s.ServeStaticFiles(false)
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
			embeddedgo.GoModelsDir,
			embeddedgo.GoDiagramsDir,
			embeddedDiagrams,
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
	// insertion point{{`+string(rune(ModelGongNLevel1tackInstanceSet))+`}}

	return
}`, staticPkg)
	}

	return fmt.Sprintf(`// do not modify, generated file
package level1stack

import (
	"fmt"
	"log"
	"strings"

	"{{PkgPathRoot}}/models"
	"{{PkgPathRoot}}/models/probe"

	embeddedgo "{{PkgPathRoot}}"

	"net/http"

	%s
)

// hook marhalling to stage
type BeforeCommitImplementation struct {
	marshallOnCommit string

	packageName string%s
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

%s
}

type Level1Stack struct {
	Stage *models.Stage%s
	Probe *probe.Probe
	R     *http.ServeMux
}

func (stack *Level1Stack) Run(addr string) error {
	return %s.RunServer(stack.R, addr)
}

%s
`,
		staticImport,
		beforeCommitStructField,
		beforeCommitMarshall,
		level1StackFields,
		staticPkg,
		constructorsAndImpl,
	)
}

var Level1StackInstanceTemplate = GetLevel1StackTemplate(false, false)
var Level1StackInstanceSplitliteTemplate = GetLevel1StackTemplate(true, false)

type ModelGongNLevel1tackInstanceStructInsertionId int

const (
	ModelGongNLevel1tackInstanceSet ModelGongNLevel1tackInstanceStructInsertionId = iota
)

var ModelGongNLevel1tackInstanceStructSubTemplateCode map[string]string = // new line
map[string]string{
	string(rune(ModelGongNLevel1tackInstanceSet)): `
	stage.SetOrchestratorOnAfterUpdate[models.{{Structname}}]()`,
}
