//go:build !js

package main

import (
	"log"
	"strconv"

	// insertion point for models import

	slider_stack "github.com/fullstack-lang/gong/lib/slider/go/stack"
	slider_static "github.com/fullstack-lang/gong/lib/slider/go/static"
)

func executeServer(args []string) {
	// setup the static file server and get the controller
	r := slider_static.ServeStaticFiles(logGINFlag)

	// setup model stack with its probe
	stack := slider_stack.NewStack(r, "slider", unmarshallFromCode, marshallOnCommit, "", embeddedDiagrams, true)
	stack.Probe.Refresh()
	stack.Stage.Commit()

	// insertion point for call to stager
	NewStager(r, stack.Stage)

	log.Println("Server ready serve on localhost:" + strconv.Itoa(port))
	err := r.Run(":" + strconv.Itoa(port))
	if err != nil {
		log.Fatalln(err.Error())
	}
}
