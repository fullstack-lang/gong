//go:build !js

package main

import (
	"log"
	"strconv"

	tree_models "github.com/fullstack-lang/gong/lib/tree/go/models"
	tree_stack "github.com/fullstack-lang/gong/lib/tree/go/stack"
	tree_static "github.com/fullstack-lang/gong/lib/tree/go/static"
)

func executeServer(args []string) {
	// setup the static file server and get the controller
	r := tree_static.ServeStaticFiles(logGINFlag)

	// setup stack
	stack := tree_stack.NewStack(r, tree_models.TreeStackDefaultName.ToString(), unmarshallFromCode, marshallOnCommit, "", embeddedDiagrams, true)
	stack.Probe.Refresh()
	stack.Stage.Commit()

	NewStager(r, stack.Stage)

	log.Printf("Server ready serve on localhost:" + strconv.Itoa(port))
	log.Println("Server ready serve on localhost:" + strconv.Itoa(port))
	err := r.Run(":" + strconv.Itoa(port))
	if err != nil {
		log.Fatalln(err.Error())
	}
}
