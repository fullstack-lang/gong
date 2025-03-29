package main

import (
	"flag"
	"log"
	"strconv"

	xlsx_models "github.com/fullstack-lang/gong/lib/xlsx/go/models"
	xlsx_stack "github.com/fullstack-lang/gong/lib/xlsx/go/stack"
	xlsx_static "github.com/fullstack-lang/gong/lib/xlsx/go/static"
)

var (
	logGINFlag = flag.Bool("logGIN", false, "log mode for gin")

	unmarshallFromCode = flag.String("unmarshallFromCode", "", "unmarshall data from go file and '.go' (must be lowercased without spaces), If unmarshallFromCode arg is '', no unmarshalling")
	marshallOnCommit   = flag.String("marshallOnCommit", "", "on all commits, marshall staged data to a go file with the marshall name and '.go' (must be lowercased without spaces). If marshall arg is '', no marshalling")

	embeddedDiagrams = flag.Bool("embeddedDiagrams", false, "parse/analysis go/models and go/embeddedDiagrams")

	compareFlag = flag.String("compare", "", "compare to the other file")

	port = flag.Int("port", 8080, "port server")
)

func main() {

	log.SetPrefix("xlsx: ")
	log.SetFlags(0)

	// parse program arguments
	flag.Parse()

	// setup the static file server and get the controller
	r := xlsx_static.ServeStaticFiles(*logGINFlag)

	// setup stack
	stack := xlsx_stack.NewStack(r, "xlsx", *unmarshallFromCode, *marshallOnCommit, "", *embeddedDiagrams, true)
	stage := stack.Stage

	args := flag.Args()
	if len(args) < 1 {
		log.Panic("there should be at least one file argument")
	}

	sampleXLFile := new(xlsx_models.XLFile).Stage(stage)
	sampleXLFile.Open(stage, flag.Arg(0))

	if len(args) > 1 {
		sampleXLFile2 := new(xlsx_models.XLFile).Stage(stage)
		sampleXLFile2.Open(stage, flag.Arg(1))
	}

	stage.Commit()

	stack.Probe.Refresh()

	NewStager(r, stack.Stage)

	log.Println("Server ready serve on localhost:" + strconv.Itoa(*port))
	err := r.Run(":" + strconv.Itoa(*port))
	if err != nil {
		log.Fatalln(err.Error())
	}
}
