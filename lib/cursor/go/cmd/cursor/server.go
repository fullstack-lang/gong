//go:build !js

package main

import (
	"log"
	"strconv"

	cursor_stack "github.com/fullstack-lang/gong/lib/cursor/go/stack"
	cursor_static "github.com/fullstack-lang/gong/lib/cursor/go/static"
)

func executeServer(args []string) {
	// setup the static file server and get the controller
	r := cursor_static.ServeStaticFiles(logGINFlag)

	// setup stack
	stack := cursor_stack.NewStack(r, "cursor", unmarshallFromCode, marshallOnCommit, "", embeddedDiagrams, true)
	stack.Probe.Refresh()
	stack.Stage.Commit()

	log.Println("Server ready serve on localhost:" + strconv.Itoa(port))
	err := r.Run(":" + strconv.Itoa(port))
	if err != nil {
		log.Fatalln(err.Error())
	}
}
