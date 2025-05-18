package main

import (
	"flag"
	"log"
	"strconv"

	// insertion point for models import
	test_models "github.com/fullstack-lang/gong/test/test/go/models"
	test_stack "github.com/fullstack-lang/gong/test/test/go/stack"
	test_static "github.com/fullstack-lang/gong/test/test/go/static"

	split "github.com/fullstack-lang/gong/lib/split/go/models"
	split_stack "github.com/fullstack-lang/gong/lib/split/go/stack"
)

var (
	logGINFlag = flag.Bool("logGIN", false, "log mode for gin")

	unmarshallFromCode = flag.String("unmarshallFromCode", "", "unmarshall data from go file and '.go' (must be lowercased without spaces), If unmarshallFromCode arg is '', no unmarshalling")
	marshallOnCommit   = flag.String("marshallOnCommit", "", "on all commits, marshall staged data to a go file with the marshall name and '.go' (must be lowercased without spaces). If marshall arg is '', no marshalling")

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

	// setup model stack with its probe
	stack := test_stack.NewStack(r, "test", *unmarshallFromCode, *marshallOnCommit, "", *embeddedDiagrams, true)
	stack.Probe.Refresh()

	// the root split name is "" by convention. Is is the same for all gong applications
	// that do not develop their specific angular component
	splitStage := split_stack.NewStack(r, "", "", "", "", false, false).Stage

	// insertion point for call to stager
	stager := test_models.NewStager(r, stack.Stage, splitStage)

	// creates two tabs, one for test
	split.StageBranch(splitStage, &split.View{
		Name:             "Test",
		RootAsSplitAreas: stager.GetAsSplitAreas(),
	})

	// one for the probe of the
	split.StageBranch(splitStage, &split.View{
		Name: "Probe of test",
		RootAsSplitAreas: []*split.AsSplitArea{
			(&split.AsSplitArea{
				Split: (&split.Split{
					StackName: stack.Stage.GetProbeSplitStageName(),
				}),
			}),
		},
	})

	// commit the split stage (this will initiate the front components)
	splitStage.Commit()

	log.Println("Server ready serve on localhost:" + strconv.Itoa(*port))
	err := r.Run(":" + strconv.Itoa(*port))
	if err != nil {
		log.Fatalln(err.Error())
	}
}
