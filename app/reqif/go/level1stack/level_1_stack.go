// do not modify, generated file
package level1stack

import (
	"fmt"
	"log"
	"strings"

	"github.com/fullstack-lang/gong/app/reqif/go/models"
	"github.com/fullstack-lang/gong/app/reqif/go/probe"

	embeddedgo "github.com/fullstack-lang/gong/app/reqif/go"

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

	stage.MarshallFile(fmt.Sprintf("./%s", filename), "github.com/fullstack-lang/gong/app/reqif/go/models", packageName)
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
	models.SetOrchestratorOnAfterUpdate[models.ALTERNATIVE_ID](stage)
	models.SetOrchestratorOnAfterUpdate[models.ATTRIBUTE_DEFINITION_BOOLEAN](stage)
	models.SetOrchestratorOnAfterUpdate[models.ATTRIBUTE_DEFINITION_BOOLEAN_Rendering](stage)
	models.SetOrchestratorOnAfterUpdate[models.ATTRIBUTE_DEFINITION_DATE](stage)
	models.SetOrchestratorOnAfterUpdate[models.ATTRIBUTE_DEFINITION_DATE_Rendering](stage)
	models.SetOrchestratorOnAfterUpdate[models.ATTRIBUTE_DEFINITION_ENUMERATION](stage)
	models.SetOrchestratorOnAfterUpdate[models.ATTRIBUTE_DEFINITION_ENUMERATION_Rendering](stage)
	models.SetOrchestratorOnAfterUpdate[models.ATTRIBUTE_DEFINITION_INTEGER](stage)
	models.SetOrchestratorOnAfterUpdate[models.ATTRIBUTE_DEFINITION_INTEGER_Rendering](stage)
	models.SetOrchestratorOnAfterUpdate[models.ATTRIBUTE_DEFINITION_REAL](stage)
	models.SetOrchestratorOnAfterUpdate[models.ATTRIBUTE_DEFINITION_REAL_Rendering](stage)
	models.SetOrchestratorOnAfterUpdate[models.ATTRIBUTE_DEFINITION_Rendering](stage)
	models.SetOrchestratorOnAfterUpdate[models.ATTRIBUTE_DEFINITION_STRING](stage)
	models.SetOrchestratorOnAfterUpdate[models.ATTRIBUTE_DEFINITION_STRING_Rendering](stage)
	models.SetOrchestratorOnAfterUpdate[models.ATTRIBUTE_DEFINITION_XHTML](stage)
	models.SetOrchestratorOnAfterUpdate[models.ATTRIBUTE_DEFINITION_XHTML_Rendering](stage)
	models.SetOrchestratorOnAfterUpdate[models.ATTRIBUTE_VALUE_BOOLEAN](stage)
	models.SetOrchestratorOnAfterUpdate[models.ATTRIBUTE_VALUE_DATE](stage)
	models.SetOrchestratorOnAfterUpdate[models.ATTRIBUTE_VALUE_ENUMERATION](stage)
	models.SetOrchestratorOnAfterUpdate[models.ATTRIBUTE_VALUE_INTEGER](stage)
	models.SetOrchestratorOnAfterUpdate[models.ATTRIBUTE_VALUE_REAL](stage)
	models.SetOrchestratorOnAfterUpdate[models.ATTRIBUTE_VALUE_STRING](stage)
	models.SetOrchestratorOnAfterUpdate[models.ATTRIBUTE_VALUE_XHTML](stage)
	models.SetOrchestratorOnAfterUpdate[models.A_ALTERNATIVE_ID](stage)
	models.SetOrchestratorOnAfterUpdate[models.A_ATTRIBUTE_DEFINITION_BOOLEAN_REF](stage)
	models.SetOrchestratorOnAfterUpdate[models.A_ATTRIBUTE_DEFINITION_DATE_REF](stage)
	models.SetOrchestratorOnAfterUpdate[models.A_ATTRIBUTE_DEFINITION_ENUMERATION_REF](stage)
	models.SetOrchestratorOnAfterUpdate[models.A_ATTRIBUTE_DEFINITION_INTEGER_REF](stage)
	models.SetOrchestratorOnAfterUpdate[models.A_ATTRIBUTE_DEFINITION_REAL_REF](stage)
	models.SetOrchestratorOnAfterUpdate[models.A_ATTRIBUTE_DEFINITION_STRING_REF](stage)
	models.SetOrchestratorOnAfterUpdate[models.A_ATTRIBUTE_DEFINITION_XHTML_REF](stage)
	models.SetOrchestratorOnAfterUpdate[models.A_ATTRIBUTE_VALUE_BOOLEAN](stage)
	models.SetOrchestratorOnAfterUpdate[models.A_ATTRIBUTE_VALUE_DATE](stage)
	models.SetOrchestratorOnAfterUpdate[models.A_ATTRIBUTE_VALUE_ENUMERATION](stage)
	models.SetOrchestratorOnAfterUpdate[models.A_ATTRIBUTE_VALUE_INTEGER](stage)
	models.SetOrchestratorOnAfterUpdate[models.A_ATTRIBUTE_VALUE_REAL](stage)
	models.SetOrchestratorOnAfterUpdate[models.A_ATTRIBUTE_VALUE_STRING](stage)
	models.SetOrchestratorOnAfterUpdate[models.A_ATTRIBUTE_VALUE_XHTML](stage)
	models.SetOrchestratorOnAfterUpdate[models.A_ATTRIBUTE_VALUE_XHTML_1](stage)
	models.SetOrchestratorOnAfterUpdate[models.A_CHILDREN](stage)
	models.SetOrchestratorOnAfterUpdate[models.A_CORE_CONTENT](stage)
	models.SetOrchestratorOnAfterUpdate[models.A_DATATYPES](stage)
	models.SetOrchestratorOnAfterUpdate[models.A_DATATYPE_DEFINITION_BOOLEAN_REF](stage)
	models.SetOrchestratorOnAfterUpdate[models.A_DATATYPE_DEFINITION_DATE_REF](stage)
	models.SetOrchestratorOnAfterUpdate[models.A_DATATYPE_DEFINITION_ENUMERATION_REF](stage)
	models.SetOrchestratorOnAfterUpdate[models.A_DATATYPE_DEFINITION_INTEGER_REF](stage)
	models.SetOrchestratorOnAfterUpdate[models.A_DATATYPE_DEFINITION_REAL_REF](stage)
	models.SetOrchestratorOnAfterUpdate[models.A_DATATYPE_DEFINITION_STRING_REF](stage)
	models.SetOrchestratorOnAfterUpdate[models.A_DATATYPE_DEFINITION_XHTML_REF](stage)
	models.SetOrchestratorOnAfterUpdate[models.A_EDITABLE_ATTS](stage)
	models.SetOrchestratorOnAfterUpdate[models.A_ENUM_VALUE_REF](stage)
	models.SetOrchestratorOnAfterUpdate[models.A_OBJECT](stage)
	models.SetOrchestratorOnAfterUpdate[models.A_PROPERTIES](stage)
	models.SetOrchestratorOnAfterUpdate[models.A_RELATION_GROUP_TYPE_REF](stage)
	models.SetOrchestratorOnAfterUpdate[models.A_SOURCE_1](stage)
	models.SetOrchestratorOnAfterUpdate[models.A_SOURCE_SPECIFICATION_1](stage)
	models.SetOrchestratorOnAfterUpdate[models.A_SPECIFICATIONS](stage)
	models.SetOrchestratorOnAfterUpdate[models.A_SPECIFICATION_TYPE_REF](stage)
	models.SetOrchestratorOnAfterUpdate[models.A_SPECIFIED_VALUES](stage)
	models.SetOrchestratorOnAfterUpdate[models.A_SPEC_ATTRIBUTES](stage)
	models.SetOrchestratorOnAfterUpdate[models.A_SPEC_OBJECTS](stage)
	models.SetOrchestratorOnAfterUpdate[models.A_SPEC_OBJECT_TYPE_REF](stage)
	models.SetOrchestratorOnAfterUpdate[models.A_SPEC_RELATIONS](stage)
	models.SetOrchestratorOnAfterUpdate[models.A_SPEC_RELATION_GROUPS](stage)
	models.SetOrchestratorOnAfterUpdate[models.A_SPEC_RELATION_REF](stage)
	models.SetOrchestratorOnAfterUpdate[models.A_SPEC_RELATION_TYPE_REF](stage)
	models.SetOrchestratorOnAfterUpdate[models.A_SPEC_TYPES](stage)
	models.SetOrchestratorOnAfterUpdate[models.A_THE_HEADER](stage)
	models.SetOrchestratorOnAfterUpdate[models.A_TOOL_EXTENSIONS](stage)
	models.SetOrchestratorOnAfterUpdate[models.DATATYPE_DEFINITION_BOOLEAN](stage)
	models.SetOrchestratorOnAfterUpdate[models.DATATYPE_DEFINITION_DATE](stage)
	models.SetOrchestratorOnAfterUpdate[models.DATATYPE_DEFINITION_ENUMERATION](stage)
	models.SetOrchestratorOnAfterUpdate[models.DATATYPE_DEFINITION_INTEGER](stage)
	models.SetOrchestratorOnAfterUpdate[models.DATATYPE_DEFINITION_REAL](stage)
	models.SetOrchestratorOnAfterUpdate[models.DATATYPE_DEFINITION_STRING](stage)
	models.SetOrchestratorOnAfterUpdate[models.DATATYPE_DEFINITION_XHTML](stage)
	models.SetOrchestratorOnAfterUpdate[models.EMBEDDED_VALUE](stage)
	models.SetOrchestratorOnAfterUpdate[models.ENUM_VALUE](stage)
	models.SetOrchestratorOnAfterUpdate[models.EmbeddedJpgImage](stage)
	models.SetOrchestratorOnAfterUpdate[models.EmbeddedPngImage](stage)
	models.SetOrchestratorOnAfterUpdate[models.EmbeddedSvgImage](stage)
	models.SetOrchestratorOnAfterUpdate[models.Kill](stage)
	models.SetOrchestratorOnAfterUpdate[models.Map_identifier_bool](stage)
	models.SetOrchestratorOnAfterUpdate[models.RELATION_GROUP](stage)
	models.SetOrchestratorOnAfterUpdate[models.RELATION_GROUP_TYPE](stage)
	models.SetOrchestratorOnAfterUpdate[models.REQ_IF](stage)
	models.SetOrchestratorOnAfterUpdate[models.REQ_IF_CONTENT](stage)
	models.SetOrchestratorOnAfterUpdate[models.REQ_IF_HEADER](stage)
	models.SetOrchestratorOnAfterUpdate[models.REQ_IF_TOOL_EXTENSION](stage)
	models.SetOrchestratorOnAfterUpdate[models.SPECIFICATION](stage)
	models.SetOrchestratorOnAfterUpdate[models.SPECIFICATION_Rendering](stage)
	models.SetOrchestratorOnAfterUpdate[models.SPECIFICATION_TYPE](stage)
	models.SetOrchestratorOnAfterUpdate[models.SPEC_HIERARCHY](stage)
	models.SetOrchestratorOnAfterUpdate[models.SPEC_OBJECT](stage)
	models.SetOrchestratorOnAfterUpdate[models.SPEC_OBJECT_TYPE](stage)
	models.SetOrchestratorOnAfterUpdate[models.SPEC_OBJECT_TYPE_Rendering](stage)
	models.SetOrchestratorOnAfterUpdate[models.SPEC_RELATION](stage)
	models.SetOrchestratorOnAfterUpdate[models.SPEC_RELATION_TYPE](stage)
	models.SetOrchestratorOnAfterUpdate[models.StaticWebSite](stage)
	models.SetOrchestratorOnAfterUpdate[models.StaticWebSiteChapter](stage)
	models.SetOrchestratorOnAfterUpdate[models.StaticWebSiteGeneratedImage](stage)
	models.SetOrchestratorOnAfterUpdate[models.StaticWebSiteImage](stage)
	models.SetOrchestratorOnAfterUpdate[models.StaticWebSiteParagraph](stage)
	models.SetOrchestratorOnAfterUpdate[models.XHTML_CONTENT](stage)

	return
}
