package main

import (
	"flag"
	"log"
	"os"

	test4_models "github.com/fullstack-lang/gong/test/test4/go/models"
	test4_stack "github.com/fullstack-lang/gong/test/test4/go/stack"
	test4_static "github.com/fullstack-lang/gong/test/test4/go/static"

	"github.com/gin-gonic/gin"
)

var (
	logGINFlag = flag.Bool("logGIN", false, "log mode for gin")

	unmarshallFromCode = flag.String("unmarshallFromCode", "", "unmarshall data from go file and '.go' (must be lowercased without spaces), If unmarshallFromCode arg is '', no unmarshalling")
	marshallOnCommit   = flag.String("marshallOnCommit", "", "on all commits, marshall staged data to a go file with the marshall name and '.go' (must be lowercased without spaces). If marshall arg is '', no marshalling")

	embeddedDiagrams = flag.Bool("embeddedDiagrams", false, "parse/analysis go/models and go/embeddedDiagrams")

	port = flag.Int("port", 8080, "port server")
)

// setupApp initializes the Gin engine and Gong stacks without starting the server.
// Note: flag.Parse() must be called by the platform-specific main functions before calling this.
func setupApp() *gin.Engine {

	log.SetPrefix("test4: ")
	log.SetFlags(log.Lmicroseconds)

	if len(flag.Args()) > 0 {
		argument := os.Args[1]
		marshallOnCommit = &argument
		unmarshallFromCode = &argument
	}

	// setup the static file server and get the controller
	r := test4_static.ServeStaticFiles(*logGINFlag)

	// setup model stack with its probe
	stack := test4_stack.NewStack(r, "test4", *unmarshallFromCode, *marshallOnCommit, "", *embeddedDiagrams, false)
	// stack.Probe.Refresh()
	stack.Stage.Commit()

	test4_models.NewStager(r, stack.Stage)

	return r
}
