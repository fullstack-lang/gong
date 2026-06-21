//go:build !js

package main

import (
	"log"
	"strconv"

	// insertion point for models import

	load_models "github.com/fullstack-lang/gong/lib/load/go/models"
	load_stack "github.com/fullstack-lang/gong/lib/load/go/stack"
	load_static "github.com/fullstack-lang/gong/lib/load/go/static"

	split "github.com/fullstack-lang/gong/lib/split/go/models"
	split_stack "github.com/fullstack-lang/gong/lib/split/go/stack"
)

func executeServer(args []string) {
	// setup the static file server and get the controller
	r := load_static.ServeStaticFiles(logGINFlag)

	// setup model stack with its probe
	stack := load_stack.NewStack(r, "load", unmarshallFromCode, marshallOnCommit, "", embeddedDiagrams, true)
	stack.Probe.Refresh()
	stack.Stage.Commit()

	// the root split name is "" by convention. Is is the same for all gong applications
	// that do not develop their specific angular component
	splitStage := split_stack.NewStack(r, "", "", "", "", false, false).Stage

	stager := load_models.NewStager(r, stack.Stage, splitStage)

	// one for the probe of the
	split.StageBranch(splitStage, &split.View{
		Name: stack.Stage.GetName() + "with Probe",
		RootAsSplitAreas: []*split.AsSplitArea{
			(&split.AsSplitArea{
				Size: 50,
				AsSplit: (&split.AsSplit{
					Direction: split.Horizontal,
					AsSplitAreas: []*split.AsSplitArea{
						stager.GetAsSplitArea(),
					},
				}),
			}),
			(&split.AsSplitArea{
				Size: 50,
				Split: (&split.Split{
					StackName: stack.Stage.GetProbeSplitStageName(),
				}),
			}),
		},
	})

	// one for the probe of the
	split.StageBranch(splitStage, &split.View{
		Name: "with Probe",
		RootAsSplitAreas: []*split.AsSplitArea{
			(&split.AsSplitArea{
				Size: 50,
				Split: (&split.Split{
					StackName: stack.Stage.GetProbeSplitStageName(),
				}),
			}),
		},
	})

	// commit the split stage (this will initiate the front components)
	splitStage.Commit()

	log.Println("Server ready serve on localhost:" + strconv.Itoa(port))
	err := r.Run(":" + strconv.Itoa(port))
	if err != nil {
		log.Fatalln(err.Error())
	}
}
