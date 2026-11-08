//go:build !js

package main

import (
	"log"
	"strconv"

	// insertion point for models import{{modelsImportDirective}}
	maplibre_stack "github.com/fullstack-lang/gong/lib/maplibre/go/stack"
	maplibre_static "github.com/fullstack-lang/gong/lib/maplibre/go/static"
	maplibre_models "github.com/fullstack-lang/gong/lib/maplibre/go/models"
)

func executeServer() {

	// setup the static file server and get the controller
	r := maplibre_static.ServeStaticFiles(logGINFlag)

	// setup model stack with its probe
	stack := maplibre_stack.NewStack(r, "maplibre", unmarshallFromCode, marshallOnCommit, "", embeddedDiagrams, true)
	stack.Probe.Refresh()
	stack.Stage.Commit()

	maplibre_models.NewStager(r, stack.Stage, stack.Probe)

	log.Println("Server ready serve on localhost:" + strconv.Itoa(port))
	err := maplibre_static.RunServer(r, ":" + strconv.Itoa(port))
	if err != nil {
		log.Fatalln(err.Error())
	}
}
