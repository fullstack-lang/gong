package cmd

import (
	"encoding/xml"
	"fmt"
	"os"
	"strings"
	"time"

	xsd_level1stack "github.com/fullstack-lang/gong/dsm/xsd/go/level1stack"

	// insertion point for models import
	xsd_models "github.com/fullstack-lang/gong/dsm/xsd/go/models"

	"github.com/gin-gonic/gin"
)

func process(args []string) (r *gin.Engine, stack *xsd_level1stack.Level1Stack) {
	start := time.Now()
	if Verbose {
		fmt.Printf("generate start\n")
	}

	xsdFilePath := args[0]
	if Verbose {
		fmt.Printf("Reading file: %s\n", xsdFilePath)
	}

	// setup model stack with its probe
	stack = xsd_level1stack.NewLevel1Stack("gongxsd", *unmarshallFromCode, *marshallOnCommit, true, true)
	stack.Probe.Refresh()

	// insertion point for call to stager
	xsd_models.NewStager(r, stack.Stage, stack.Probe)

	content, err := os.ReadFile(xsdFilePath)
	if err != nil {
		fmt.Printf("Error reading XSD file: %v\n", err)
		os.Exit(1)
	}

	err = xml.Unmarshal(content, &xsd_models.SchemaSingloton)
	if err != nil {
		fmt.Printf("Error parsing XML: %v\n", err)
		os.Exit(1)
	}

	// suppress duplicates
	if strings.Contains(xsdFilePath, "dtc-11-04-05") {
		xsd_models.SchemaSingloton.FactorDuplicates()
		xsd_models.SchemaSingloton.RenameTypeAnonymousComplexType()
	}

	stack.Stage.StageBranchSchema(&xsd_models.SchemaSingloton)

	stack.Stage.ComputeReverseMaps()

	xsd_models.PostProcessing(stack.Stage)

	if Verbose {
		fmt.Printf("generating file %s\n", *outputModelFilePath)
	}

	// and XSD is a acyclic directed graph (ADG) and it can be interesting
	// to navigate this ADG. Therefore, one first set the interface links between
	// the XSD nodes
	xsd_models.SchemaSingloton.SetParentAndChildren(nil)
	xsd_models.Generate(stack.Stage, *outputModelFilePath)

	if Verbose {
		fmt.Printf("generate took %s\n", time.Since(start))
	}

	return stack.R, stack
}
