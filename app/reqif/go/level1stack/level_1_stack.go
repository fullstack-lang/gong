// do not modify, generated file
package level1stack

import (
	"fmt"
	"log"
	"strings"

	"github.com/fullstack-lang/gong/app/reqif/go/models"
	"github.com/fullstack-lang/gong/app/reqif/go/models/probe"

	embeddedgo "github.com/fullstack-lang/gong/app/reqif/go"

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

	stage.MarshallFile(fmt.Sprintf("./%s", filename), "github.com/fullstack-lang/gong/app/reqif/go/models", packageName)
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
	stage.SetOrchestratorOnAfterUpdate[models.ALTERNATIVE_ID]()
	stage.SetOrchestratorOnAfterUpdate[models.ATTRIBUTE_DEFINITION_BOOLEAN]()
	stage.SetOrchestratorOnAfterUpdate[models.ATTRIBUTE_DEFINITION_BOOLEAN_Rendering]()
	stage.SetOrchestratorOnAfterUpdate[models.ATTRIBUTE_DEFINITION_DATE]()
	stage.SetOrchestratorOnAfterUpdate[models.ATTRIBUTE_DEFINITION_DATE_Rendering]()
	stage.SetOrchestratorOnAfterUpdate[models.ATTRIBUTE_DEFINITION_ENUMERATION]()
	stage.SetOrchestratorOnAfterUpdate[models.ATTRIBUTE_DEFINITION_ENUMERATION_Rendering]()
	stage.SetOrchestratorOnAfterUpdate[models.ATTRIBUTE_DEFINITION_INTEGER]()
	stage.SetOrchestratorOnAfterUpdate[models.ATTRIBUTE_DEFINITION_INTEGER_Rendering]()
	stage.SetOrchestratorOnAfterUpdate[models.ATTRIBUTE_DEFINITION_REAL]()
	stage.SetOrchestratorOnAfterUpdate[models.ATTRIBUTE_DEFINITION_REAL_Rendering]()
	stage.SetOrchestratorOnAfterUpdate[models.ATTRIBUTE_DEFINITION_Rendering]()
	stage.SetOrchestratorOnAfterUpdate[models.ATTRIBUTE_DEFINITION_STRING]()
	stage.SetOrchestratorOnAfterUpdate[models.ATTRIBUTE_DEFINITION_STRING_Rendering]()
	stage.SetOrchestratorOnAfterUpdate[models.ATTRIBUTE_DEFINITION_XHTML]()
	stage.SetOrchestratorOnAfterUpdate[models.ATTRIBUTE_DEFINITION_XHTML_Rendering]()
	stage.SetOrchestratorOnAfterUpdate[models.ATTRIBUTE_VALUE_BOOLEAN]()
	stage.SetOrchestratorOnAfterUpdate[models.ATTRIBUTE_VALUE_DATE]()
	stage.SetOrchestratorOnAfterUpdate[models.ATTRIBUTE_VALUE_ENUMERATION]()
	stage.SetOrchestratorOnAfterUpdate[models.ATTRIBUTE_VALUE_INTEGER]()
	stage.SetOrchestratorOnAfterUpdate[models.ATTRIBUTE_VALUE_REAL]()
	stage.SetOrchestratorOnAfterUpdate[models.ATTRIBUTE_VALUE_STRING]()
	stage.SetOrchestratorOnAfterUpdate[models.ATTRIBUTE_VALUE_XHTML]()
	stage.SetOrchestratorOnAfterUpdate[models.A_ALTERNATIVE_ID]()
	stage.SetOrchestratorOnAfterUpdate[models.A_ATTRIBUTE_DEFINITION_BOOLEAN_REF]()
	stage.SetOrchestratorOnAfterUpdate[models.A_ATTRIBUTE_DEFINITION_DATE_REF]()
	stage.SetOrchestratorOnAfterUpdate[models.A_ATTRIBUTE_DEFINITION_ENUMERATION_REF]()
	stage.SetOrchestratorOnAfterUpdate[models.A_ATTRIBUTE_DEFINITION_INTEGER_REF]()
	stage.SetOrchestratorOnAfterUpdate[models.A_ATTRIBUTE_DEFINITION_REAL_REF]()
	stage.SetOrchestratorOnAfterUpdate[models.A_ATTRIBUTE_DEFINITION_STRING_REF]()
	stage.SetOrchestratorOnAfterUpdate[models.A_ATTRIBUTE_DEFINITION_XHTML_REF]()
	stage.SetOrchestratorOnAfterUpdate[models.A_ATTRIBUTE_VALUE_BOOLEAN]()
	stage.SetOrchestratorOnAfterUpdate[models.A_ATTRIBUTE_VALUE_DATE]()
	stage.SetOrchestratorOnAfterUpdate[models.A_ATTRIBUTE_VALUE_ENUMERATION]()
	stage.SetOrchestratorOnAfterUpdate[models.A_ATTRIBUTE_VALUE_INTEGER]()
	stage.SetOrchestratorOnAfterUpdate[models.A_ATTRIBUTE_VALUE_REAL]()
	stage.SetOrchestratorOnAfterUpdate[models.A_ATTRIBUTE_VALUE_STRING]()
	stage.SetOrchestratorOnAfterUpdate[models.A_ATTRIBUTE_VALUE_XHTML]()
	stage.SetOrchestratorOnAfterUpdate[models.A_ATTRIBUTE_VALUE_XHTML_1]()
	stage.SetOrchestratorOnAfterUpdate[models.A_CHILDREN]()
	stage.SetOrchestratorOnAfterUpdate[models.A_CORE_CONTENT]()
	stage.SetOrchestratorOnAfterUpdate[models.A_DATATYPES]()
	stage.SetOrchestratorOnAfterUpdate[models.A_DATATYPE_DEFINITION_BOOLEAN_REF]()
	stage.SetOrchestratorOnAfterUpdate[models.A_DATATYPE_DEFINITION_DATE_REF]()
	stage.SetOrchestratorOnAfterUpdate[models.A_DATATYPE_DEFINITION_ENUMERATION_REF]()
	stage.SetOrchestratorOnAfterUpdate[models.A_DATATYPE_DEFINITION_INTEGER_REF]()
	stage.SetOrchestratorOnAfterUpdate[models.A_DATATYPE_DEFINITION_REAL_REF]()
	stage.SetOrchestratorOnAfterUpdate[models.A_DATATYPE_DEFINITION_STRING_REF]()
	stage.SetOrchestratorOnAfterUpdate[models.A_DATATYPE_DEFINITION_XHTML_REF]()
	stage.SetOrchestratorOnAfterUpdate[models.A_EDITABLE_ATTS]()
	stage.SetOrchestratorOnAfterUpdate[models.A_ENUM_VALUE_REF]()
	stage.SetOrchestratorOnAfterUpdate[models.A_OBJECT]()
	stage.SetOrchestratorOnAfterUpdate[models.A_PROPERTIES]()
	stage.SetOrchestratorOnAfterUpdate[models.A_RELATION_GROUP_TYPE_REF]()
	stage.SetOrchestratorOnAfterUpdate[models.A_SOURCE_1]()
	stage.SetOrchestratorOnAfterUpdate[models.A_SOURCE_SPECIFICATION_1]()
	stage.SetOrchestratorOnAfterUpdate[models.A_SPECIFICATIONS]()
	stage.SetOrchestratorOnAfterUpdate[models.A_SPECIFICATION_TYPE_REF]()
	stage.SetOrchestratorOnAfterUpdate[models.A_SPECIFIED_VALUES]()
	stage.SetOrchestratorOnAfterUpdate[models.A_SPEC_ATTRIBUTES]()
	stage.SetOrchestratorOnAfterUpdate[models.A_SPEC_OBJECTS]()
	stage.SetOrchestratorOnAfterUpdate[models.A_SPEC_OBJECT_TYPE_REF]()
	stage.SetOrchestratorOnAfterUpdate[models.A_SPEC_RELATIONS]()
	stage.SetOrchestratorOnAfterUpdate[models.A_SPEC_RELATION_GROUPS]()
	stage.SetOrchestratorOnAfterUpdate[models.A_SPEC_RELATION_REF]()
	stage.SetOrchestratorOnAfterUpdate[models.A_SPEC_RELATION_TYPE_REF]()
	stage.SetOrchestratorOnAfterUpdate[models.A_SPEC_TYPES]()
	stage.SetOrchestratorOnAfterUpdate[models.A_THE_HEADER]()
	stage.SetOrchestratorOnAfterUpdate[models.A_TOOL_EXTENSIONS]()
	stage.SetOrchestratorOnAfterUpdate[models.DATATYPE_DEFINITION_BOOLEAN]()
	stage.SetOrchestratorOnAfterUpdate[models.DATATYPE_DEFINITION_DATE]()
	stage.SetOrchestratorOnAfterUpdate[models.DATATYPE_DEFINITION_ENUMERATION]()
	stage.SetOrchestratorOnAfterUpdate[models.DATATYPE_DEFINITION_INTEGER]()
	stage.SetOrchestratorOnAfterUpdate[models.DATATYPE_DEFINITION_REAL]()
	stage.SetOrchestratorOnAfterUpdate[models.DATATYPE_DEFINITION_STRING]()
	stage.SetOrchestratorOnAfterUpdate[models.DATATYPE_DEFINITION_XHTML]()
	stage.SetOrchestratorOnAfterUpdate[models.EMBEDDED_VALUE]()
	stage.SetOrchestratorOnAfterUpdate[models.ENUM_VALUE]()
	stage.SetOrchestratorOnAfterUpdate[models.EmbeddedJpgImage]()
	stage.SetOrchestratorOnAfterUpdate[models.EmbeddedPngImage]()
	stage.SetOrchestratorOnAfterUpdate[models.EmbeddedSvgImage]()
	stage.SetOrchestratorOnAfterUpdate[models.Kill]()
	stage.SetOrchestratorOnAfterUpdate[models.Map_identifier_bool]()
	stage.SetOrchestratorOnAfterUpdate[models.RELATION_GROUP]()
	stage.SetOrchestratorOnAfterUpdate[models.RELATION_GROUP_TYPE]()
	stage.SetOrchestratorOnAfterUpdate[models.REQ_IF]()
	stage.SetOrchestratorOnAfterUpdate[models.REQ_IF_CONTENT]()
	stage.SetOrchestratorOnAfterUpdate[models.REQ_IF_HEADER]()
	stage.SetOrchestratorOnAfterUpdate[models.REQ_IF_TOOL_EXTENSION]()
	stage.SetOrchestratorOnAfterUpdate[models.SPECIFICATION]()
	stage.SetOrchestratorOnAfterUpdate[models.SPECIFICATION_Rendering]()
	stage.SetOrchestratorOnAfterUpdate[models.SPECIFICATION_TYPE]()
	stage.SetOrchestratorOnAfterUpdate[models.SPEC_HIERARCHY]()
	stage.SetOrchestratorOnAfterUpdate[models.SPEC_OBJECT]()
	stage.SetOrchestratorOnAfterUpdate[models.SPEC_OBJECT_TYPE]()
	stage.SetOrchestratorOnAfterUpdate[models.SPEC_OBJECT_TYPE_Rendering]()
	stage.SetOrchestratorOnAfterUpdate[models.SPEC_RELATION]()
	stage.SetOrchestratorOnAfterUpdate[models.SPEC_RELATION_TYPE]()
	stage.SetOrchestratorOnAfterUpdate[models.StaticWebSite]()
	stage.SetOrchestratorOnAfterUpdate[models.StaticWebSiteChapter]()
	stage.SetOrchestratorOnAfterUpdate[models.StaticWebSiteGeneratedImage]()
	stage.SetOrchestratorOnAfterUpdate[models.StaticWebSiteImage]()
	stage.SetOrchestratorOnAfterUpdate[models.StaticWebSiteParagraph]()
	stage.SetOrchestratorOnAfterUpdate[models.XHTML_CONTENT]()

	return
}
