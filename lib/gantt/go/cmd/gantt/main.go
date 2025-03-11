package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	gantt_models "github.com/fullstack-lang/gong/lib/gantt/go/models"
	gantt_stack "github.com/fullstack-lang/gong/lib/gantt/go/stack"
	gantt_static "github.com/fullstack-lang/gong/lib/gantt/go/static"

	svg_models "github.com/fullstack-lang/gong/lib/svg/go/models"
	svg_stack "github.com/fullstack-lang/gong/lib/svg/go/stack"

	"github.com/fullstack-lang/gong/lib/gantt/go/gantt2svg"
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

	// setup stackgantt
	stackgantt := gantt_stack.NewStack(r, gantt_models.GanttStackName.ToString(), *unmarshallFromCode, *marshallOnCommit, "", *embeddedDiagrams, true)
	stackgantt.Probe.Refresh()

	stacksvg := svg_stack.NewStack(r, gantt_models.SvgStackName.ToString(), "", "", "", true, true)
	stacksvg.Probe.Refresh()

	// set up the GanttSVGMapper that will intercept
	// commits on the gantt stage and that will
	// generate the svg
	ganttSVGMapper := new(gantt2svg.GanttSVGMapper)
	ganttSVGMapper.GanttOuputFile = *marshallOnCommit

	commitOnGanttStage := new(CommitFromFrontOnGanttStage)
	commitOnGanttStage.gongsvgStage = stacksvg.Stage
	commitOnGanttStage.ganttSVGMapper = ganttSVGMapper

	// hook on the commit from front
	stackgantt.Stage.OnInitCommitFromFrontCallback = commitOnGanttStage
	stackgantt.Stage.OnInitCommitFromBackCallback = commitOnGanttStage

	// initial publication
	ganttSVGMapper.GenerateSvg(stackgantt.Stage, stacksvg.Stage)

	log.Println("Server ready serve on localhost:" + strconv.Itoa(*port))
	err := r.Run(":" + strconv.Itoa(*port))
	if err != nil {
		log.Fatalln(err.Error())
	}
}

// hook marhalling to stage
type CommitFromFrontOnGanttStage struct {
	gongsvgStage   *svg_models.StageStruct
	ganttSVGMapper *gantt2svg.GanttSVGMapper
}

// BeforeCommit meets the interface for the commit on the gantt stage
// It performs 2 tasks
// 1 - update the SVG stack
// 2 - persists the data to the gantt file
func (beforeCommitFromFrontOnGanttStage *CommitFromFrontOnGanttStage) BeforeCommit(
	ganttStage *gantt_models.StageStruct) {
	file, err := os.Create(fmt.Sprintf("./%s.go", *marshallOnCommit))
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()

	// update the gantt stage with the back repo data that was updated from the front
	ganttStage.Checkout()

	// marshall to the file
	ganttStage.Marshall(file, "github.com/fullstack-lang/gong/lib/gantt/go/models", "main")
	beforeCommitFromFrontOnGanttStage.ganttSVGMapper.GenerateSvg(ganttStage, beforeCommitFromFrontOnGanttStage.gongsvgStage)
}
