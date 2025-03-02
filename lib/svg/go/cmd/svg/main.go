package main

import (
	"flag"
	"log"
	"strconv"

	gongsvg_models "github.com/fullstack-lang/gong/lib/svg/go/models"
	gongsvg_stack "github.com/fullstack-lang/gong/lib/svg/go/stack"
	gongsvg_static "github.com/fullstack-lang/gong/lib/svg/go/static"
)

var (
	logGINFlag = flag.Bool("logGIN", false, "log mode for gin")

	unmarshallFromCode = flag.String("unmarshallFromCode", "", "unmarshall data from go file and '.go' (must be lowercased without spaces), If unmarshallFromCode arg is '', no unmarshalling")
	marshallOnCommit   = flag.String("marshallOnCommit", "", "on all commits, marshall staged data to a go file with the marshall name and '.go' (must be lowercased without spaces). If marshall arg is '', no marshalling")

	embeddedDiagrams = flag.Bool("embeddedDiagrams", false, "parse/analysis go/models and go/embeddedDiagrams")

	port = flag.Int("port", 8080, "port server")
)

func main() {

	log.SetPrefix("gongsvg: ")
	log.SetFlags(0)

	// parse program arguments
	flag.Parse()

	// setup the static file server and get the controller
	r := gongsvg_static.ServeStaticFiles(*logGINFlag)

	// setup stack
	stack := gongsvg_stack.NewStack(r, gongsvg_models.StackNameDefault.ToString(), *unmarshallFromCode, *marshallOnCommit, "", *embeddedDiagrams, true)
	stack.Probe.Refresh()

	// get the unique svg and hook a callback for when it is edited or not
	// the is the only way to update the commit nb from the back
	var svg *gongsvg_models.SVG
	for svg_ := range *gongsvg_models.GetGongstructInstancesSet[gongsvg_models.SVG](stack.Stage) {
		svg = svg_
	}
	if svg != nil {
		svg.Impl = &SVGImpl{
			stack: stack,
			svg:   svg}

	}

	log.Printf("Server ready serve on localhost:" + strconv.Itoa(*port))
	err := r.Run(":" + strconv.Itoa(*port))
	if err != nil {
		log.Fatalln(err.Error())
	}
}

type SVGImpl struct {
	stack *gongsvg_stack.Stack
	svg   *gongsvg_models.SVG
}

// SVGUpdated implements models.SVGImplInterface.
func (sVGImpl *SVGImpl) SVGUpdated(updatedSVG *gongsvg_models.SVG) {
	sVGImpl.svg.IsEditable = updatedSVG.IsEditable
	sVGImpl.stack.Stage.Commit()
}
