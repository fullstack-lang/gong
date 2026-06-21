//go:build !js

package main

import (
	"log"
	"strconv"

	// insertion point for models import
	button_models "github.com/fullstack-lang/gong/lib/button/go/models"
	button_stack "github.com/fullstack-lang/gong/lib/button/go/stack"
	button_static "github.com/fullstack-lang/gong/lib/button/go/static"
)

func executeServer(args []string) {
	// setup the static file server and get the controller
	r := button_static.ServeStaticFiles(logGINFlag)

	// setup model stack with its probe
	stack := button_stack.NewStack(r, "button", unmarshallFromCode, marshallOnCommit, "", embeddedDiagrams, true)
	stack.Probe.Refresh()

	// insertion point for call to stager
	button_models.NewStager(r, stack.Stage)

	log.Println("Server ready serve on localhost:" + strconv.Itoa(port))
	err := r.Run(":" + strconv.Itoa(port))
	if err != nil {
		log.Fatalln(err.Error())
	}
}
