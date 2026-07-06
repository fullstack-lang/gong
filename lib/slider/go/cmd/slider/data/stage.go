package main

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/lib/slider/go/models"
	// injection point for ident package import declaration{{ImportPackageDeclaration}}
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var _ time.Time
var _ = slices.Index[[]int, int]

// function will stage objects
func _(stage *models.Stage) {

	// insertion point for declaration of instances to stage

	__Checkbox__00000000_ := (&models.Checkbox{Name: `Checkbox 1`}).Stage(stage)

	__Group__00000000_ := (&models.Group{Name: `group 1`}).Stage(stage)
	__Group__00000001_ := (&models.Group{Name: `Group 2`}).Stage(stage)

	__Layout__00000000_ := (&models.Layout{Name: `layout`}).Stage(stage)

	__Slider__00000000_ := (&models.Slider{Name: `slider int`}).Stage(stage)

	// insertion point for initialization of values

	__Checkbox__00000000_.Name = `Checkbox 1`
	__Checkbox__00000000_.ValueBool = true
	__Checkbox__00000000_.LabelForTrue = `Label For True`
	__Checkbox__00000000_.LabelForFalse = `Label for False`

	__Group__00000000_.Name = `group 1`
	__Group__00000000_.Percentage = 100.000000

	__Group__00000001_.Name = `Group 2`
	__Group__00000001_.Percentage = 50.000000

	__Layout__00000000_.Name = `layout`
	__Layout__00000000_.IsWithCustomGutterSize = true
	__Layout__00000000_.GutterSize = 1.000000

	__Slider__00000000_.Name = `slider int`
	__Slider__00000000_.IsFloat64 = false
	__Slider__00000000_.IsInt = true
	__Slider__00000000_.MinInt = 10
	__Slider__00000000_.MaxInt = 70
	__Slider__00000000_.StepInt = 5
	__Slider__00000000_.ValueInt = 10
	__Slider__00000000_.MinFloat64 = 0.000000
	__Slider__00000000_.MaxFloat64 = 0.000000
	__Slider__00000000_.StepFloat64 = 0.000000
	__Slider__00000000_.ValueFloat64 = 0.000000

	// insertion point for setup of pointers
	__Group__00000000_.Sliders = append(__Group__00000000_.Sliders, __Slider__00000000_)
	__Group__00000000_.Checkboxes = append(__Group__00000000_.Checkboxes, __Checkbox__00000000_)
	__Layout__00000000_.Groups = append(__Layout__00000000_.Groups, __Group__00000000_)
	__Layout__00000000_.Groups = append(__Layout__00000000_.Groups, __Group__00000001_)
}
