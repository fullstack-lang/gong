//go:build !js

package main

import (
	"log"
	"strconv"

	"github.com/fullstack-lang/gong/lib/threejs/go/models"
	threejs_stack "github.com/fullstack-lang/gong/lib/threejs/go/stack"
	threejs_static "github.com/fullstack-lang/gong/lib/threejs/go/static"
)

func executeServer() {

	r := threejs_static.ServeStaticFiles(logGINFlag)

	// setup
	// - model stack with its probe
	// - unmarshall/marshall go file with stage data
	stack := threejs_stack.NewStack(r, "threejs", unmarshallFromCode, marshallOnCommit, "", embeddedDiagrams, true)

	// refresh the probe, therefore we can see what has been unmarshalled
	stack.Probe.Refresh()

	if generateStrangeFormsFlag {
		models.GenerateReferenceScene(stack.Stage)
	} else {
		stack.Stage.Commit()
	}

	// initiates the UX loop
	models.NewStager(
		r,
		stack.Stage,
		stack.Probe,
	)

	log.Println("Server ready serve on localhost:" + strconv.Itoa(port))
	err := r.Run(":" + strconv.Itoa(port))
	if err != nil {
		log.Fatalln(err.Error())
	}
}
