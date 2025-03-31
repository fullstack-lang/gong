package main

import (
	"flag"
	"log"
	"strconv"

	// insertion point for models import
	probe_models "github.com/fullstack-lang/gong/lib/probe/go/models"
	probe_stack "github.com/fullstack-lang/gong/lib/probe/go/stack"
	probe_static "github.com/fullstack-lang/gong/lib/probe/go/static"

	test_stack "github.com/fullstack-lang/gong/test/test/go/stack"
)

var (
	logGINFlag = flag.Bool("logGIN", false, "log mode for gin")

	unmarshallFromCode = flag.String("unmarshallFromCode", "", "unmarshall data from go file and '.go' (must be lowercased without spaces), If unmarshallFromCode arg is '', no unmarshalling")
	marshallOnCommit   = flag.String("marshallOnCommit", "", "on all commits, marshall staged data to a go file with the marshall name and '.go' (must be lowercased without spaces). If marshall arg is '', no marshalling")

	embeddedDiagrams = flag.Bool("embeddedDiagrams", false, "parse/analysis go/models and go/embeddedDiagrams")

	port = flag.Int("port", 8080, "port server")
)

func main() {

	log.SetPrefix("probe: ")
	log.SetFlags(0)

	// parse program arguments
	flag.Parse()

	// setup the static file server and get the controller
	r := probe_static.ServeStaticFiles(*logGINFlag)

	// setup model stack with its probe
	stack := probe_stack.NewStack(r, "probe", *unmarshallFromCode, *marshallOnCommit, "", *embeddedDiagrams, true)
	stack.Probe.Refresh()

	// no legacy probe and embedded diagram
	stackTest := test_stack.NewStack(r, "test", *unmarshallFromCode, *marshallOnCommit, "", true, false)

	// probe will create a split front end
	probe := probe_models.NewProbe2(r, stack.Stage, stackTest.Stage)

	// cmd stager will hosts the probe split (with name of the stack)
	NewStager(r, stack.Stage, probe)

	log.Println("Server ready serve on localhost:" + strconv.Itoa(*port))
	err := r.Run(":" + strconv.Itoa(*port))
	if err != nil {
		log.Fatalln(err.Error())
	}
}
