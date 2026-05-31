package main

import (
	"embed"
	"flag"
	"log"
	"os"

	project_level1stack "github.com/fullstack-lang/gong/dsm/project/go/level1stack"
	project_models "github.com/fullstack-lang/gong/dsm/project/go/models"

	"github.com/gin-gonic/gin"
)

var (
	logGINFlag = flag.Bool("logGIN", false, "log mode for gin")

	unmarshallFromCode = flag.String("unmarshallFromCode", "", "unmarshall data from go file and '.go' (must be lowercased without spaces), If unmarshallFromCode arg is '', no unmarshalling")
	marshallOnCommit   = flag.String("marshallOnCommit", "", "on all commits, marshall staged data to a go file with the marshall name and '.go' (must be lowercased without spaces). If marshall arg is '', no marshalling")

	embeddedDiagrams = flag.Bool("embeddedDiagrams", true, "parse/analysis go/models and go/embeddedDiagrams")

	port = flag.Int("port", 8080, "port server")
)

//go:embed data
var stage_content embed.FS

// setupApp initializes the Gin engine and Gong stacks without starting the server.
// Note: flag.Parse() must be called by the platform-specific main functions before calling this.
func setupApp() (r *gin.Engine) {

	log.SetPrefix("test4: ")
	log.SetFlags(log.Lmicroseconds)

	if len(flag.Args()) > 0 {
		argument := os.Args[1]
		marshallOnCommit = &argument
		unmarshallFromCode = &argument
	}

	// // setup model stack with its probe
	// stack = test4_stack.NewStack(r, "test4", *unmarshallFromCode, *marshallOnCommit, "", *embeddedDiagrams, true)
	// stack.Probe.Refresh()

	stackproject := project_level1stack.NewLevel1StackDelta("project", "", "", true, true, true)

	// initiates the UX loop
	project_models.NewStager(
		stackproject.R,
		stackproject.Stage,
		stackproject.Probe,
	)
	r = stackproject.R

	return
}
