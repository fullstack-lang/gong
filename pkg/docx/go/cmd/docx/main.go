package main

import (
	"flag"
	"log"
	"os"
	"strconv"

	"github.com/fullstack-lang/gong/pkg/docx/go/level1stack"
	"github.com/fullstack-lang/gong/pkg/docx/go/models"
)

var (
	unmarshallFromCode = flag.String("unmarshallFromCode", "", "unmarshall data from go file and '.go' (must be lowercased without spaces), If unmarshallFromCode arg is '', no unmarshalling")
	marshallOnCommit   = flag.String("marshallOnCommit", "", "on all commits, marshall staged data to a go file with the marshall name and '.go' (must be lowercased without spaces). If marshall arg is '', no marshalling")

	embeddedDiagrams = flag.Bool("embeddedDiagrams", false, "parse/analysis go/models and go/embeddedDiagrams")

	port = flag.Int("port", 8080, "port server")
)

func main() {

	log.SetPrefix("docx: ")
	log.SetFlags(log.Lmicroseconds)

	// parse program arguments
	flag.Parse()

	if flag.NArg() == 0 || flag.NArg() > 1 {
		flag.Usage()
		os.Exit(1)
	}

	// setup
	// - model level1 stack with its probe
	// - unmarshall/marshall go file with stage data
	stack := level1stack.NewLevel1Stack("docx", *unmarshallFromCode, *marshallOnCommit, true, *embeddedDiagrams)

	// refresh the probe, therefore we can see what has been unmarshalled
	stack.Probe.Refresh()

	filename := flag.Args()[0]

	// initiates the UX loop
	models.NewStager(
		stack.R,
		stack.Stage,
		stack.Probe,
		filename,
	)

	log.Println("Server ready serve on localhost:" + strconv.Itoa(*port))
	err := stack.R.Run(":" + strconv.Itoa(*port))
	if err != nil {
		log.Fatalln(err.Error())
	}
}
