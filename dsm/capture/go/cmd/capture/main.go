//go:build !js

package main

import (
	"embed"
	"flag"
	"log"
	"os"
	"strconv"

	"github.com/fullstack-lang/gong/dsm/capture/go/level1stack"
	"github.com/fullstack-lang/gong/dsm/capture/go/models"
)

var (
	unmarshallFromCode = flag.String("unmarshallFromCode", "", "unmarshall data from go file and '.go' (must be lowercased without spaces), If unmarshallFromCode arg is '', no unmarshalling")
	marshallOnCommit   = flag.String("marshallOnCommit", "", "on all commits, marshall staged data to a go file with the marshall name and '.go' (must be lowercased without spaces). If marshall arg is '', no marshalling")

	embeddedDiagrams = flag.Bool("embeddedDiagrams", false, "parse/analysis go/models and go/embeddedDiagrams")

	port = flag.Int("port", 8080, "port server")
)

//go:embed data/*
var dataFS embed.FS

func main() {

	log.SetPrefix("capture: ")
	log.SetFlags(log.Lmicroseconds)

	// parse program arguments
	flag.Parse()

	if flag.NArg() > 0 {
		argument := os.Args[1]
		marshallOnCommit = &argument
		unmarshallFromCode = &argument
	}

	// setup
	// - set embedded data to models
	models.DataFS = &dataFS
	// - model level1 stack with its probe
	// - unmarshall/marshall go file with stage data
	stack := level1stack.NewLevel1Stack("capture", *unmarshallFromCode, *marshallOnCommit, true, *embeddedDiagrams)

	// refresh the probe, therefore we can see what has been unmarshalled
	stack.Probe.Refresh()

	// initiates the UX loop
	models.NewStager(
		stack.R,
		stack.Stage,
		stack.Probe,
	)

	log.Println("Server ready serve on localhost:" + strconv.Itoa(*port))
	err := stack.R.Run(":" + strconv.Itoa(*port))
	if err != nil {
		log.Fatalln(err.Error())
	}
}
