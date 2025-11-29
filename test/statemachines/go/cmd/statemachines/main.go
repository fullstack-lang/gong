package main

import (
	"flag"
	"log"
	"strconv"

	// insertion point for models import

	"github.com/fullstack-lang/gong/test/statemachines/go/level1stack"
	"github.com/fullstack-lang/gong/test/statemachines/go/models"
)

var (
	unmarshallFromCode = flag.String("unmarshallFromCode", "", "unmarshall data from go file and '.go' (must be lowercased without spaces), If unmarshallFromCode arg is '', no unmarshalling")
	marshallOnCommit   = flag.String("marshallOnCommit", "", "on all commits, marshall staged data to a go file with the marshall name and '.go' (must be lowercased without spaces). If marshall arg is '', no marshalling")

	embeddedDiagrams = flag.Bool("embeddedDiagrams", false, "parse/analysis go/models and go/embeddedDiagrams")

	port = flag.Int("port", 8080, "port server")
)

func main() {

	log.SetPrefix("gongstatemachines: ")
	log.SetFlags(log.Lmicroseconds)

	flag.Parse()

	stack := level1stack.NewLevel1Stack("gongstatemachines", *unmarshallFromCode, *marshallOnCommit, true, *embeddedDiagrams)

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
