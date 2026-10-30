package cmd

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	xsd_level1stack "github.com/fullstack-lang/gong/app/xsd/go/level1stack"
	xsd_models "github.com/fullstack-lang/gong/app/xsd/go/models"
)

func process(args []string) (r *http.ServeMux, stack *xsd_level1stack.Level1Stack) {
	start := time.Now()
	if Verbose {
		fmt.Printf("xsd2gong generate start\n")
	}

	xsdFilePath := args[0]
	if Verbose {
		fmt.Printf("Reading root XSD file: %s\n", xsdFilePath)
	}

	// Setup model stack with its probe
	stack = xsd_level1stack.NewLevel1Stack("gongxsd", *unmarshallFromCode, *marshallOnCommit, true, true)
	stack.Probe.Refresh()

	xsd_models.NewStager(r, stack.Stage, stack.Probe)

	targetModelsDir := *outputModelFilePath
	if targetModelsDir == "" {
		targetModelsDir = filepath.Join(filepath.Dir(xsdFilePath), "go", "models")
	} else if strings.HasSuffix(targetModelsDir, ".go") {
		targetModelsDir = filepath.Dir(targetModelsDir)
	}

	sg, err := xsd_models.LoadSchemaGraph(stack.Stage, xsdFilePath, targetModelsDir)
	if err != nil {
		fmt.Printf("Error loading XSD schema graph: %v\n", err)
		os.Exit(1)
	}

	if Verbose {
		fmt.Printf("Loaded %d schema(s) into %d target Go package(s) in %s\n",
			len(sg.SchemasInOrder), len(sg.TargetPackages), targetModelsDir)
	}

	ctx, err := xsd_models.GenerateMultiPackage(sg, targetModelsDir)
	if err != nil {
		fmt.Printf("Error generating multi-package Go models: %v\n", err)
		os.Exit(1)
	}

	if err := ctx.GenerateDSMScaffold(); err != nil {
		fmt.Printf("Error generating DSM scaffold: %v\n", err)
		os.Exit(1)
	}

	if Verbose {
		fmt.Printf("xsd2gong generate took %s\n", time.Since(start))
	}

	return stack.R, stack
}
