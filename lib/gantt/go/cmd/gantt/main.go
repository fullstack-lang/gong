package main

import (
	"flag"
	"log"
	"strconv"

	// insertion point for models import
	gantt_models "github.com/fullstack-lang/gong/lib/gantt/go/models"
	gantt_stack "github.com/fullstack-lang/gong/lib/gantt/go/stack"
	gantt_static "github.com/fullstack-lang/gong/lib/gantt/go/static"

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

	log.SetPrefix("gantt: ")
	log.SetFlags(0)

	// parse program arguments
	flag.Parse()

	// setup the static file server and get the controller
	r := gantt_static.ServeStaticFiles(*logGINFlag)

	// setup model stack with its probe
	stack := gantt_stack.NewStack(r, "gantt", *unmarshallFromCode, *marshallOnCommit, "", *embeddedDiagrams, true)
	stack.Probe.Refresh()

	// the root split name is "" by convention. Is is the same for all gong applications
	// that do not develop their specific angular component
	splitStage := split_stack.NewStack(r, "", "", "", "", false, false).Stage

	stager := gantt_models.NewStager(r, stack.Stage, splitStage)

	// set up the GanttSVGMapper that will intercept
	// commits on the gantt stage and that will
	// generate the svg
	ganttSVGMapper := new(gantt_models.GanttSVGMapper)
	ganttSVGMapper.GanttOuputFile = *marshallOnCommit

	commitOnGanttStage := new(CommitFromFrontOnGanttStage)
	commitOnGanttStage.gongsvgStage = stager.GetSvgStage()
	commitOnGanttStage.ganttSVGMapper = ganttSVGMapper

	// hook on the commit from front
	stack.Stage.OnInitCommitFromFrontCallback = commitOnGanttStage
	stack.Stage.OnInitCommitFromBackCallback = commitOnGanttStage

	// initial publication
	ganttSVGMapper.GenerateSvg(stack.Stage, stager.GetSvgStage())

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

	// commit the split stage (this will initiate the front components)
	splitStage.Commit()

	log.Println("Server ready serve on localhost:" + strconv.Itoa(*port))
	err := r.Run(":" + strconv.Itoa(*port))
	if err != nil {
		log.Fatalln(err.Error())
	}
}
