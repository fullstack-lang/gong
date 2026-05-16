package main

import (
	"flag"
	"log"

	"github.com/fullstack-lang/gong/dsm/process/go/models"
)

var (
	outputPath = flag.String("output", "merged.go", "output file for the merged stage")
)

func main() {
	log.SetPrefix("merge: ")
	log.SetFlags(0)

	flag.Parse()

	if flag.NArg() == 0 {
		log.Fatal("no input files provided")
	}

	stage := models.NewStage("merge")

	for _, file := range flag.Args() {
		log.Printf("parsing %s", file)
		err := models.ParseAstFile(stage, file, false)
		if err != nil {
			log.Fatalf("error unmarshalling %s: %s", file, err.Error())
		}
	}

	stage.ComputeReverseMaps()
	stage.ComputeInstancesNb()
	stage.ComputeReferenceAndOrders()

	log.Printf("marshalling merged stage to %s", *outputPath)
	stage.MarshallFile(*outputPath, "github.com/fullstack-lang/gong/dsm/process/go/models", "main")
}
