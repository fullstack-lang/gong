//go:build !js

package main

import (
	"log"
	"strconv"

	// insertion point for models import

	"github.com/fullstack-lang/gong/lib/doc/go/level1stack"

	split "github.com/fullstack-lang/gong/lib/split/go/models"
	split_stack "github.com/fullstack-lang/gong/lib/split/go/stack"
)

func executeServer(args []string) {
	// setup model stack with its probe
	stack := level1stack.NewLevel1Stack("doc", unmarshallFromCode, marshallOnCommit, true, embeddedDiagrams)

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

	log.Println("Server ready serve on localhost:" + strconv.Itoa(port))
	err := stack.R.Run(":" + strconv.Itoa(port))
	if err != nil {
		log.Fatalln(err.Error())
	}
}
