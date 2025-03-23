package main

import (
	"flag"
	"log"
	"strconv"

	load_models "github.com/fullstack-lang/gong/lib/load/go/models"
	load_stack "github.com/fullstack-lang/gong/lib/load/go/stack"
	load_static "github.com/fullstack-lang/gong/lib/load/go/static"

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

	log.SetPrefix("load: ")
	log.SetFlags(0)

	// parse program arguments
	flag.Parse()

	// setup the static file server and get the controller
	r := load_static.ServeStaticFiles(*logGINFlag)

	// setup model stack with its probe
	stack := load_stack.NewStack(r, "load", *unmarshallFromCode, *marshallOnCommit, "", *embeddedDiagrams, true)
	stack.Probe.Refresh()

	// setup root split stack and insert the probe at the root
	splitStage := split_stack.NewStack(r, "", "", "", "", false, false).Stage

	(&split.View{
		Name: "Download",
		RootAsSplitAreas: []*split.AsSplitArea{
			(&split.AsSplitArea{
				Load: (&split.Load{
					StackName: stack.Stage.GetName(),
				}).Stage(splitStage),
			}).Stage(splitStage),
		},
	}).Stage(splitStage)

	(&split.View{
		Name: "Probe",
		RootAsSplitAreas: []*split.AsSplitArea{
			(&split.AsSplitArea{
				Split: (&split.Split{
					StackName: stack.Stage.GetName() + load_models.ProbeSplitSuffix,
				}).Stage(splitStage),
			}).Stage(splitStage),
		},
	}).Stage(splitStage)

	splitStage.Commit()

	log.Println("Server ready serve on localhost:" + strconv.Itoa(*port))
	err := r.Run(":" + strconv.Itoa(*port))
	if err != nil {
		log.Fatalln(err.Error())
	}
}
