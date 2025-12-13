package main

import (
	"flag"
	"log"
	"strconv"

	// insertion point for models import

	"github.com/fullstack-lang/gong/lib/doc/go/level1stack"

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

	log.SetPrefix("doc: ")
	log.SetFlags(0)

	// parse program arguments
	flag.Parse()

	// setup model stack with its probe
	stack := level1stack.NewLevel1Stack("doc", *unmarshallFromCode, *marshallOnCommit, true, *embeddedDiagrams)

	// since we do not use the default stager, we need to create the root split
	splitStage := split_stack.NewStack(stack.R, "", "", "", "", false, false).Stage

	probeSplitStageName := stack.Stage.GetProbeSplitStageName()

	split.StageBranch(splitStage, &split.View{
		Name: "Probe",
		RootAsSplitAreas: []*split.AsSplitArea{
			(&split.AsSplitArea{
				Split: (&split.Split{
					StackName: probeSplitStageName,
				}),
			}),
		},
	})

	splitStage.Commit()

	log.Println("Server ready serve on localhost:" + strconv.Itoa(*port))
	err := stack.R.Run(":" + strconv.Itoa(*port))
	if err != nil {
		log.Fatalln(err.Error())
	}
}
