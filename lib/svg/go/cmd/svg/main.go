package main

import (
	"flag"
	"log"
	"path/filepath"
	"strconv"

	// insertion point for models import
	svg_models "github.com/fullstack-lang/gong/lib/svg/go/models"
	svg_stack "github.com/fullstack-lang/gong/lib/svg/go/stack"
	svg_static "github.com/fullstack-lang/gong/lib/svg/go/static"
)

var (
	logGINFlag = flag.Bool("logGIN", false, "log mode for gin")

	unmarshallFromCode = flag.String("unmarshallFromCode", "", "unmarshall data from go file and '.go' (must be lowercased without spaces), If unmarshallFromCode arg is '', no unmarshalling")
	marshallOnCommit   = flag.String("marshallOnCommit", "", "on all commits, marshall staged data to a go file with the marshall name and '.go' (must be lowercased without spaces). If marshall arg is '', no marshalling")

	embeddedDiagrams = flag.Bool("embeddedDiagrams", false, "parse/analysis go/models and go/embeddedDiagrams")

	port = flag.Int("port", 8080, "port server")
)

func main() {

	log.SetPrefix("svg: ")
	log.SetFlags(0)

	// parse program arguments
	flag.Parse()

	// setup the static file server and get the controller
	r := svg_static.ServeStaticFiles(*logGINFlag)

	// setup model stack with its probe
	stack := svg_stack.NewStack(r, "svg", *unmarshallFromCode, *marshallOnCommit, "", *embeddedDiagrams, true)
	stack.Probe.Refresh()

	// get the unique svg and hook a callback for when it is edited or not
	// the is the only way to update the commit nb from the back
	var svg *svg_models.SVG
	for svg_ := range *svg_models.GetGongstructInstancesSet[svg_models.SVG](stack.Stage) {
		svg = svg_
	}
	if svg != nil {
		svg.Impl = &SVGImpl{
			stack: stack,
			svg:   svg}

	}

	// insertion point for call to stager
	NewStager(r, stack.Stage)

	for svg := range *svg_models.GetGongstructInstancesSet[svg_models.SVG](stack.Stage) {
		err := svg.GenerateFile(filepath.Join("../../diagrams/images", svg.Name+".svg"))
		if err != nil {
			log.Fatalln("Unable to generate file", err.Error())
		}
	}

	log.Println("Server ready serve on localhost:" + strconv.Itoa(*port))
	err := r.Run(":" + strconv.Itoa(*port))
	if err != nil {
		log.Fatalln(err.Error())
	}
}

type SVGImpl struct {
	stack *svg_stack.Stack
	svg   *svg_models.SVG
}

// SVGUpdated implements models.SVGImplInterface.
func (sVGImpl *SVGImpl) SVGUpdated(updatedSVG *svg_models.SVG) {
	sVGImpl.svg.IsEditable = updatedSVG.IsEditable
	sVGImpl.stack.Stage.Commit()
}
