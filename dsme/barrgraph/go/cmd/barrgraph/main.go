package main

import (
	_ "embed"
	"flag"
	"go/parser"
	"go/token"
	"log"
	"os"
	"strconv"

	"github.com/fullstack-lang/gong/dsme/barrgraph/go/level1stack"
	"github.com/fullstack-lang/gong/dsme/barrgraph/go/models"
)

var (
	unmarshallFromCode = flag.String("unmarshallFromCode", "", "unmarshall data from go file and '.go' (must be lowercased without spaces), If unmarshallFromCode arg is '', no unmarshalling")
	marshallOnCommit   = flag.String("marshallOnCommit", "", "on all commits, marshall staged data to a go file with the marshall name and '.go' (must be lowercased without spaces). If marshall arg is '', no marshalling")

	embeddedDiagrams = flag.Bool("embeddedDiagrams", true, "parse/analysis go/models and go/embeddedDiagrams")

	port = flag.Int("port", 8080, "port server")
)

//go:embed "data/cubism and abstract art.go"
var stageGo string

func main() {
	log.SetPrefix("barrgraph: ")
	log.SetFlags(log.Lmicroseconds)

	// parse program arguments
	flag.Parse()

	// setup
	// - model level1 stack with its probe
	// - unmarshall/marshall go file with stage data
	stack := level1stack.NewLevel1Stack("barrgraph", *unmarshallFromCode, *marshallOnCommit, true, *embeddedDiagrams)

	stage := stack.Stage

	if *unmarshallFromCode == "" {
		fset := token.NewFileSet()
		file, err := parser.ParseFile(fset, "", stageGo, parser.ParseComments)
		if err != nil {
			log.Panic("Unable to parse stageGo " + err.Error())
		}
		err = models.ParseAstFileFromAst(stage, file, fset, true)
		// if the application is run with -unmarshallFromCode=xxx.go -marshallOnCommit
		// xxx.go might be absent the first time. However, this shall not be a show stopper.
		if err != nil {
			log.Println("no file to read " + err.Error())
		}

		stage.ComputeReverseMaps()
		stage.ComputeInstancesNb()
		stage.ComputeReferenceAndOrders()
	}

	// refresh the probe, therefore we can see what has been unmarshalled
	stack.Probe.Refresh()

	// initiates the UX loop
	models.NewStager(
		stack.R,
		stack.Stage,
		stack.Probe,
	)

	portStr := os.Getenv("PORT")
	if portStr == "" {
		portStr = strconv.Itoa(*port) // Fallback on flag
	}
	*port, _ = strconv.Atoi(portStr)

	log.Println("Server ready serve on localhost:" + strconv.Itoa(*port))
	err := stack.R.Run(":" + strconv.Itoa(*port))
	if err != nil {
		log.Fatalln(err.Error())
	}
}
