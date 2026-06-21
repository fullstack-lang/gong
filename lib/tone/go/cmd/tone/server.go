//go:build !js

package main

import (
	"log"
	"strconv"

	//
	tone_models "github.com/fullstack-lang/gong/lib/tone/go/models"
	tone_stack "github.com/fullstack-lang/gong/lib/tone/go/stack"
	tone_static "github.com/fullstack-lang/gong/lib/tone/go/static"
)

func executeServer(args []string) {
	// setup the static file server and get the controller
	r := tone_static.ServeStaticFiles(logGINFlag)

	// setup model stack with its probe
	stack := tone_stack.NewStack(r, "tone", unmarshallFromCode, marshallOnCommit, "", embeddedDiagrams, true)
	stack.Probe.Refresh()
	stack.Stage.Commit()

	//
	tone_models.NewStager(r, stack.Stage)

	log.Println("Server ready serve on localhost:" + strconv.Itoa(port))
	err := r.Run(":" + strconv.Itoa(port))
	if err != nil {
		log.Fatalln(err.Error())
	}
}
