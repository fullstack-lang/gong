package main

import (
	"flag"
	"log"
	"strconv"

	// insertion point for models import
	// markdown_models "github.com/fullstack-lang/gong/lib/markdown/go/models"
	markdown_stack "github.com/fullstack-lang/gong/lib/markdown/go/stack"
	markdown_static "github.com/fullstack-lang/gong/lib/markdown/go/static"
)

var (
	logGINFlag = flag.Bool("logGIN", false, "log mode for gin")

	unmarshallFromCode = flag.String("unmarshallFromCode", "", "unmarshall data from go file and '.go' (must be lowercased without spaces), If unmarshallFromCode arg is '', no unmarshalling")
	marshallOnCommit   = flag.String("marshallOnCommit", "", "on all commits, marshall staged data to a go file with the marshall name and '.go' (must be lowercased without spaces). If marshall arg is '', no marshalling")

	embeddedDiagrams = flag.Bool("embeddedDiagrams", false, "parse/analysis go/models and go/embeddedDiagrams")

	port = flag.Int("port", 8080, "port server")
)

func main() {

	log.SetPrefix("markdown: ")
	log.SetFlags(log.Lmicroseconds)

	// parse program arguments
	flag.Parse()

	// setup the static file server and get the controller
	r := markdown_static.ServeStaticFiles(*logGINFlag)

	// setup model stack with its probe
	stack := markdown_stack.NewStack(r, "markdown", *unmarshallFromCode, *marshallOnCommit, "", *embeddedDiagrams, true)
	stack.Probe.Refresh()

	NewStager(r, stack.Stage)

	log.Println("Server ready serve on localhost:" + strconv.Itoa(*port))
	err := r.Run(":" + strconv.Itoa(*port))
	if err != nil {
		log.Fatalln(err.Error())
	}
}
