package main

import (
	"flag"
	"log"
	"strconv"

	// insertion point for models import

	split_static "github.com/fullstack-lang/gong/lib/split/go/static"

	table "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gong/test/statemachines/go/models"
	"github.com/fullstack-lang/gong/test/statemachines/go/probe"
	"github.com/fullstack-lang/gong/test/statemachines/go/stack"
)

var (
	logGINFlag = flag.Bool("logGIN", false, "log mode for gin")

	unmarshallFromCode = flag.String("unmarshallFromCode", "", "unmarshall data from go file and '.go' (must be lowercased without spaces), If unmarshallFromCode arg is '', no unmarshalling")
	marshallOnCommit   = flag.String("marshallOnCommit", "", "on all commits, marshall staged data to a go file with the marshall name and '.go' (must be lowercased without spaces). If marshall arg is '', no marshalling")

	embeddedDiagrams = flag.Bool("embeddedDiagrams", false, "parse/analysis go/models and go/embeddedDiagrams")

	port = flag.Int("port", 8080, "port server")
)

func main() {

	log.SetPrefix("gongstatemachines: ")
	log.SetFlags(log.Lmicroseconds)

	// parse program arguments
	flag.Parse()

	// setup the static file server and get the controller
	r := split_static.ServeStaticFiles(*logGINFlag)

	// setup model stack_ with its probe
	stack_ := stack.NewStack(r, "gongstatemachines", *unmarshallFromCode, *marshallOnCommit, "", *embeddedDiagrams, true)
	stack_.Probe.Refresh()

	// // the root split name is "" by convention. Is is the same for all gong applications
	// // that do not develop their specific angular component
	// splitStage := split_stack.NewStack(r, "", "", "", "", false, false).Stage

	stager := models.NewStager(
		r,
		stack_.Stage,
		stack_.Probe.GetDataEditor(),
		stack_.Probe.GetDiagramEditor(),
		&ProbeForm{
			probe: stack_.Probe,
		},
	)
	_ = stager

	log.Println("Server ready serve on localhost:" + strconv.Itoa(*port))
	err := r.Run(":" + strconv.Itoa(*port))
	if err != nil {
		log.Fatalln(err.Error())
	}
}

type ProbeForm struct {
	probe *probe.Probe
}

// FillUpNamedFormFromGongstruct implements models.ProbeFormInterface.
func (p *ProbeForm) FillUpFormFromGongstruct(instance any, formName string) {
	probe.FillUpFormFromGongstruct(instance, p.probe)
}

func (p *ProbeForm) GetFormStage() *table.Stage {
	return p.probe.GetFormStage()
}
