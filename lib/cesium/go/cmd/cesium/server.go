//go:build !js

package main

import (
	"log"
	"strconv"

	// insertion point for models import{{modelsImportDirective}}
	cesium_stack "github.com/fullstack-lang/gong/lib/cesium/go/stack"
	cesium_static "github.com/fullstack-lang/gong/lib/cesium/go/static"
	cesium_models "github.com/fullstack-lang/gong/lib/cesium/go/models"
)

func executeServer() {

	// setup the static file server and get the controller
	r := cesium_static.ServeStaticFiles(logGINFlag)

	// setup model stack with its probe
	stack := cesium_stack.NewStack(r, "cesium", unmarshallFromCode, marshallOnCommit, "", embeddedDiagrams, true)
	stack.Probe.Refresh()
	stack.Stage.Commit()

	cesium_models.NewStager(r, stack.Stage, stack.Probe)

	log.Println("Server ready serve on localhost:" + strconv.Itoa(port))
	err := cesium_static.RunServer(r, ":" + strconv.Itoa(port))
	if err != nil {
		log.Fatalln(err.Error())
	}
}
