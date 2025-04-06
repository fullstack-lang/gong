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
func _(stage *models.Stage) {

	// Declaration of instances to stage

	__Checkbox__000000_Checkbox_1 := (&models.Checkbox{}).Stage(stage)

	__Group__000000_group_1 := (&models.Group{}).Stage(stage)

	__Layout__000000_layout := (&models.Layout{}).Stage(stage)

	__Slider__000000_slider_int := (&models.Slider{}).Stage(stage)

	// Setup of values

	__Checkbox__000000_Checkbox_1.Name = `Checkbox 1`
	__Checkbox__000000_Checkbox_1.ValueBool = true
	__Checkbox__000000_Checkbox_1.LabelForTrue = `Label For True`
	__Checkbox__000000_Checkbox_1.LabelForFalse = `Label for False`

	__Group__000000_group_1.Name = `group 1`
	__Group__000000_group_1.Percentage = 100.000000

	__Layout__000000_layout.Name = `layout`

	__Slider__000000_slider_int.Name = `slider int`
	__Slider__000000_slider_int.IsFloat64 = false
	__Slider__000000_slider_int.IsInt = true
	__Slider__000000_slider_int.MinInt = 10
	__Slider__000000_slider_int.MaxInt = 70
	__Slider__000000_slider_int.StepInt = 5
	__Slider__000000_slider_int.ValueInt = 10
	__Slider__000000_slider_int.MinFloat64 = 0.000000
	__Slider__000000_slider_int.MaxFloat64 = 0.000000
	__Slider__000000_slider_int.StepFloat64 = 0.000000
	__Slider__000000_slider_int.ValueFloat64 = 0.000000

	// Setup of pointers
	// setup of Checkbox instances pointers
	// setup of Group instances pointers
	__Group__000000_group_1.Sliders = append(__Group__000000_group_1.Sliders, __Slider__000000_slider_int)
	__Group__000000_group_1.Checkboxes = append(__Group__000000_group_1.Checkboxes, __Checkbox__000000_Checkbox_1)
	// setup of Layout instances pointers
	__Layout__000000_layout.Groups = append(__Layout__000000_layout.Groups, __Group__000000_group_1)
	// setup of Slider instances pointers
}
