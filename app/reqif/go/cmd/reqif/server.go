//go:build !js

package main

import (
	"flag"
	"log"
	"strconv"
	"strings"

	// insertion point for models import

	// the location of the following go package is important
	// They are NOT within the "github.com/fullstack-lang/gong/app/reqif/go/models"
	// because this package is within the "//go:embed models" directive in
	// the "github.com/fullstack-lang/gong/app/reqif/go" package
	// Therefore any change to those packge would pro
	"github.com/fullstack-lang/gong/app/reqif/go/datatypes"
	"github.com/fullstack-lang/gong/app/reqif/go/exporter"
	"github.com/fullstack-lang/gong/app/reqif/go/generator"
	"github.com/fullstack-lang/gong/app/reqif/go/namer"
	"github.com/fullstack-lang/gong/app/reqif/go/reqifz"
	"github.com/fullstack-lang/gong/app/reqif/go/specifications"
	"github.com/fullstack-lang/gong/app/reqif/go/specobjects"
	"github.com/fullstack-lang/gong/app/reqif/go/specrelations"
	"github.com/fullstack-lang/gong/app/reqif/go/spectypes"

	gongreqif_level1stack "github.com/fullstack-lang/gong/app/reqif/go/level1stack"
	gongreqif_models "github.com/fullstack-lang/gong/app/reqif/go/models"
)

var (
	pathToReqifFile       = flag.String("pathToReqifFile", "", "Path to the reqif file")
	pathToRenderingConf   = flag.String("pathToRenderingConf", "", "Path to the rendering conf file")
	pathToGoModelFile     = flag.String("pathToGoModelFile", "", "Path to the go model file")
	pathToXLFile          = flag.String("pathToXLFile", "", "Path to the go model file")
	pathToOutputReqifFile = flag.String("pathToOutputReqifFile", "", "Path to the output reqif file")
	logFile               = flag.String("logFile", "gongreqif.log", "path to the log file")
)

func executeServer(args []string) {
	if len(args) > 0 {
		pathToReqifOrReqifz := args[0]
		isReqifz, err := reqifz.ConvertReQIFzToReqIF(pathToReqifOrReqifz)
		if err != nil {
			log.Fatalln("Not able to convert reqif file", err.Error())
		}

		if isReqifz {
			*pathToReqifFile = strings.TrimSuffix(pathToReqifOrReqifz, ".reqifz") + ".reqif"
		}
	}

	// setup model stack with its probe
	stack := gongreqif_level1stack.NewLevel1Stack("reqif", unmarshallFromCode, marshallOnCommit, true, embeddedDiagrams)
	stack.Probe.Refresh()

	// insertion point for call to stager
	stager := gongreqif_models.NewStager(
		stack.R,
		stack.Stage,
		*pathToReqifFile,
		*pathToRenderingConf,
		*pathToOutputReqifFile,
		&datatypes.DataTypeTreeStageUpdater{},
		&spectypes.SpecTypesTreeStageUpdater{},
		&specobjects.SpecObjectsTreeStageUpdater{},
		&specrelations.SpecRelationsTreeStageUpdater{},
		&specifications.SpecificationsTreeStageUpdater{},
		&exporter.Exporter{},
		&namer.ObjectNamer{})

	if *pathToGoModelFile != "" {
		modelGenerator := &generator.ModelGenerator{
			PathToGoModelFile: *pathToGoModelFile,
		}
		stager.SetModelGenerator(modelGenerator)
		modelGenerator.GenerateModels(stager)
	}

	if *pathToXLFile != "" {
		gongreqif_models.SerializeStage(stack.Stage, *pathToXLFile)
	}

	log.Println("Server ready serve on localhost:" + strconv.Itoa(port))
	err := stack.R.Run(":" + strconv.Itoa(port))
	if err != nil {
		log.Fatalln(err.Error())
	}
}
