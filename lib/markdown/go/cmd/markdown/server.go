//go:build !js

package main

import (
	"log"
	"strconv"

	// insertion point for models import
	// markdown_models "github.com/fullstack-lang/gong/lib/markdown/go/models"
	markdown_stack "github.com/fullstack-lang/gong/lib/markdown/go/stack"
	markdown_static "github.com/fullstack-lang/gong/lib/markdown/go/static"
)

func executeServer(args []string) {
	// setup the static file server and get the controller
	r := markdown_static.ServeStaticFiles(logGINFlag)

	// setup model stack with its probe
	stack := markdown_stack.NewStack(r, "markdown", unmarshallFromCode, marshallOnCommit, "", embeddedDiagrams, true)
	stack.Probe.Refresh()
	stack.Stage.Commit()

	NewStager(r, stack.Stage)

	log.Println("Server ready serve on localhost:" + strconv.Itoa(port))
	err := r.Run(":" + strconv.Itoa(port))
	if err != nil {
		log.Fatalln(err.Error())
	}
}
