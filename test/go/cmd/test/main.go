package main

import (
	"flag"
	"log"
	"strconv"
	"time"

	"github.com/fullstack-lang/gong/test/go/models"
	test_stack "github.com/fullstack-lang/gong/test/go/stack"
	test_static "github.com/fullstack-lang/gong/test/go/static"
)

var (
	logGINFlag = flag.Bool("logGIN", false, "log mode for gin")

	unmarshallFromCode = flag.String("unmarshallFromCode", "", "unmarshall data from go file and '.go' (must be lowercased without spaces), If unmarshallFromCode arg is '', no unmarshalling")
	marshallOnCommit   = flag.String("marshallOnCommit", "", "on all commits, marshall staged data to a go file with the marshall name and '.go' (must be lowercased without spaces). If marshall arg is '', no marshalling")

	diagrams         = flag.Bool("diagrams", true, "parse/analysis go/models and go/diagrams")
	embeddedDiagrams = flag.Bool("embeddedDiagrams", false, "parse/analysis go/models and go/embeddedDiagrams")

	port = flag.Int("port", 8080, "port server")
)

func main() {

	log.SetPrefix("test: ")
	log.SetFlags(0)

	// parse program arguments
	flag.Parse()

	// setup the static file server and get the controller
	r := test_static.ServeStaticFiles(*logGINFlag)

	// setup stack
	stack := test_stack.NewStack(r, "test", *unmarshallFromCode, *marshallOnCommit, "", *embeddedDiagrams, true)
	stack.Probe.Refresh()

	go func() {

		time.Sleep(1 * time.Second)

		// get first element
		map_A := (*models.GetGongstructInstancesMap[models.Astruct](stack.Stage))
		if a, ok := map_A["A1"]; ok {
			for {

				time.Sleep(1 * time.Second)
				log.Println("a", a.Name)
				a.Name = a.Name + "*"
				stack.Stage.Commit()

			}
		}

	}()

	log.Printf("Server ready serve on localhost:" + strconv.Itoa(*port))
	err := r.Run(":" + strconv.Itoa(*port))
	if err != nil {
		log.Fatalln(err.Error())
	}

}
