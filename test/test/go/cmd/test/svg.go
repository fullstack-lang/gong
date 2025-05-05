package main

import (
	"time"

	"github.com/fullstack-lang/gong/lib/svg/go/models"
	// injection point for ident package import declaration{{ImportPackageDeclaration}}
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var _ time.Time

// _ point for meta package dummy declaration{{ImportPackageDummyDeclaration}}

// When parsed, those maps will help with the renaming process
var _ map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// function will stage objects
func _(stage *models.Stage) {

	// Declaration of instances to stage

	__Circle__000000_ := (&models.Circle{}).Stage(stage)

	__Layer__000000_ := (&models.Layer{}).Stage(stage)

	__SVG__000000_ := (&models.SVG{}).Stage(stage)

	// Setup of values

	__Circle__000000_.Name = ``
	__Circle__000000_.CX = 200.000000
	__Circle__000000_.CY = 200.000000
	__Circle__000000_.Radius = 100.000000
	__Circle__000000_.Color = ``
	__Circle__000000_.FillOpacity = 0.000000
	__Circle__000000_.Stroke = `black`
	__Circle__000000_.StrokeOpacity = 1.000000
	__Circle__000000_.StrokeWidth = 2.000000
	__Circle__000000_.StrokeDashArray = ``
	__Circle__000000_.StrokeDashArrayWhenSelected = ``
	__Circle__000000_.Transform = ``

	__Layer__000000_.Display = true
	__Layer__000000_.Name = ``

	__SVG__000000_.Name = ``
	__SVG__000000_.IsEditable = false
	__SVG__000000_.IsSVGFileGenerated = false

	// Setup of pointers
	// setup of Circle instances pointers
	// setup of Layer instances pointers
	__Layer__000000_.Circles = append(__Layer__000000_.Circles, __Circle__000000_)
	// setup of SVG instances pointers
	__SVG__000000_.Layers = append(__SVG__000000_.Layers, __Layer__000000_)
}
