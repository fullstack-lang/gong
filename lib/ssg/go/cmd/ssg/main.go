package main

import (
	"strconv"

	"flag"
	"log"

	// Import the time package

	// Use alias 'meta' to avoid conflict
	// <-- Import the standard extension package

	// Import the html renderer package

	// YAML parsing is handled by goldmark-meta

	// insertion point for models import

	ssg_models "github.com/fullstack-lang/gong/lib/ssg/go/models"
	ssg_stack "github.com/fullstack-lang/gong/lib/ssg/go/stack"
	ssg_static "github.com/fullstack-lang/gong/lib/ssg/go/static"
)

var (
	logGINFlag = flag.Bool("logGIN", false, "log mode for gin")

	unmarshallFromCode = flag.String("unmarshallFromCode", "", "unmarshall data from go file and '.go' (must be lowercased without spaces), If unmarshallFromCode arg is '', no unmarshalling")
	marshallOnCommit   = flag.String("marshallOnCommit", "", "on all commits, marshall staged data to a go file with the marshall name and '.go' (must be lowercased without spaces). If marshall arg is '', no marshalling")

	embeddedDiagrams = flag.Bool("embeddedDiagrams", false, "parse/analysis go/models and go/embeddedDiagrams")

	port = flag.Int("port", 8080, "port server")

	contentDir    = "content"
	staticDir     = "static"
	layoutDir     = "layouts"
	serverPort    = "8080"
	serveFlag     = flag.Bool("serve", false, "Serve the site on localhost after building (or without building if -no-build is used)")
	targetFlag    = flag.String("target", "file", "Build target: 'web' for server or 'file' for local filesystem (ignored if -no-build is used)")
	outputDirFlag = flag.String("output", "public", "Output directory for the generated site")
	noBuildFlag   = flag.Bool("no-build", false, "Skip the build process (requires -serve and existing output directory)")
)

func main() {

	log.SetPrefix("ssg: ")
	log.SetFlags(0)

	// parse program arguments
	flag.Parse()

	// setup the static file server and get the controller
	r := ssg_static.ServeStaticFiles(*logGINFlag)

	// setup model stack with its probe
	stack := ssg_stack.NewStack(r, "ssg", *unmarshallFromCode, *marshallOnCommit, "", *embeddedDiagrams, true)
	stack.Probe.Refresh()

	// insertion point for call to stager
	ssg_models.NewStager(r, stack.Stage)

	log.Println("Server ready serve on localhost:" + strconv.Itoa(*port))
	err := r.Run(":" + strconv.Itoa(*port))
	if err != nil {
		log.Fatalln(err.Error())
	}

}
