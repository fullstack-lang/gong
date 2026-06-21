//go:build !js

package main

import (
	"log"
	"os"
	"path/filepath"
	"strconv"

	// insertion point for models import
	svg_models "github.com/fullstack-lang/gong/lib/svg/go/models"
	svg_stack "github.com/fullstack-lang/gong/lib/svg/go/stack"
	svg_static "github.com/fullstack-lang/gong/lib/svg/go/static"
)

func executeServer(args []string) {
	// setup the static file server and get the controller
	r := svg_static.ServeStaticFiles(logGINFlag)

	// setup model stack with its probe
	stack := svg_stack.NewStack(r, "svg", unmarshallFromCode, marshallOnCommit, "", embeddedDiagrams, true)
	stack.Probe.Refresh()
	stack.Stage.Commit()

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
		result, _, _ := svg.GenerateString()
		err := os.WriteFile(filepath.Join("../../diagrams/images", svg.Name+".svg"), []byte(result), 0644)
		if err != nil {
			log.Fatalln("Unable to generate file", err.Error())
		}
	}

	log.Println("Server ready serve on localhost:" + strconv.Itoa(port))
	err := r.Run(":" + strconv.Itoa(port))
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
