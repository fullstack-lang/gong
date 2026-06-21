package commands

import (
	"log"
	"strconv"

	test_models "github.com/fullstack-lang/gong/test/test/go/models"
	test_stack "github.com/fullstack-lang/gong/test/test/go/stack"
	test_static "github.com/fullstack-lang/gong/test/test/go/static"

	split "github.com/fullstack-lang/gong/lib/split/go/models"
	split_stack "github.com/fullstack-lang/gong/lib/split/go/stack"
)

func executeServer(unmarshallFromCode string, marshallOnCommit string, port int, embeddedDiagrams bool, logGINFlag bool) {

	// setup the static file server and get the controller
	r := test_static.ServeStaticFiles(logGINFlag)

	// setup model stack with its probe
	stack := test_stack.NewStack(r, "test", unmarshallFromCode, marshallOnCommit, "", embeddedDiagrams, true)
	stack.Stage.Commit()
	stack.Probe.Refresh()

	// the root split name is "" by convention. Is is the same for all gong applications
	// that do not develop their specific angular component
	splitStage := split_stack.NewStack(r, "", "", "", "", false, false).Stage

	// set the title of the application (name of the tab)
	(&split.Title{Name: "Test"}).Stage(splitStage)
	(&split.FavIcon{Name: "Test",
		SVG: `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none">
			<circle cx="12" cy="12" r="10" fill="#ff6b6b"/>
			<path d="M8 14s1.5 2 4 2 4-2 4-2" stroke="white" stroke-width="2" stroke-linecap="round"/>
			<circle cx="9" cy="9" r="1" fill="white"/>
			<circle cx="15" cy="9" r="1" fill="white"/>
		</svg>`,
	}).Stage(splitStage)

	stager := test_models.NewStager(r, stack.Stage, splitStage)

	tmp := test_models.GetStructInstancesByOrderAuto[*test_models.Astruct](stack.Stage)

	log.Printf("%v\n", tmp)

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
				Size:    50,
				AsSplit: stack.Probe.GetDataEditor(),
			}),
		},
	})

	split.StageBranch(splitStage, &split.View{
		Name: "Diagram Editor",
		RootAsSplitAreas: []*split.AsSplitArea{
			stack.Probe.GetDiagramEditor(),
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
