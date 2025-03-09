package main

import (
	"time"

	"github.com/fullstack-lang/gong/lib/slider/go/models"

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
func _(stage *models.StageStruct) {

	// Declaration of instances to stage

	__Group__000000_g := (&models.Group{}).Stage(stage)

	__Layout__000000_l := (&models.Layout{}).Stage(stage)

	__Slider__000000_slider_1 := (&models.Slider{}).Stage(stage)

	// Setup of values

	__Group__000000_g.Name = `g`
	__Group__000000_g.Percentage = 100.000000

	__Layout__000000_l.Name = `l`

	__Slider__000000_slider_1.Name = `slider 1`
	__Slider__000000_slider_1.IsFloat64 = false
	__Slider__000000_slider_1.IsInt = true
	__Slider__000000_slider_1.MinInt = 20
	__Slider__000000_slider_1.MaxInt = 300
	__Slider__000000_slider_1.StepInt = 10
	__Slider__000000_slider_1.ValueInt = 20
	__Slider__000000_slider_1.MinFloat64 = 0.000000
	__Slider__000000_slider_1.MaxFloat64 = 0.000000
	__Slider__000000_slider_1.StepFloat64 = 0.000000
	__Slider__000000_slider_1.ValueFloat64 = 0.000000

	// Setup of pointers
	// setup of Group instances pointers
	__Group__000000_g.Sliders = append(__Group__000000_g.Sliders, __Slider__000000_slider_1)
	// setup of Layout instances pointers
	__Layout__000000_l.Groups = append(__Layout__000000_l.Groups, __Group__000000_g)
	// setup of Slider instances pointers
}
