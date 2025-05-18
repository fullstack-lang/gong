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

	__Layer__000000_layer_1 := (&models.Layer{}).Stage(stage)

	__Rect__000000_rect_1 := (&models.Rect{}).Stage(stage)

	__SVG__000000_svg := (&models.SVG{}).Stage(stage)

	// Setup of values

	__Layer__000000_layer_1.Name = `layer 1`

	__Rect__000000_rect_1.Name = `rect 1`
	__Rect__000000_rect_1.X = 100.000000
	__Rect__000000_rect_1.Y = 100.000000
	__Rect__000000_rect_1.Width = 167.000000
	__Rect__000000_rect_1.Height = 100.000000
	__Rect__000000_rect_1.RX = 20.000000
	__Rect__000000_rect_1.Color = ``
	__Rect__000000_rect_1.FillOpacity = 0.000000
	__Rect__000000_rect_1.Stroke = `lightblue`
	__Rect__000000_rect_1.StrokeOpacity = 1.000000
	__Rect__000000_rect_1.StrokeWidth = 3.000000
	__Rect__000000_rect_1.StrokeDashArray = ``
	__Rect__000000_rect_1.StrokeDashArrayWhenSelected = ``
	__Rect__000000_rect_1.Transform = ``
	__Rect__000000_rect_1.IsSelectable = false
	__Rect__000000_rect_1.IsSelected = false
	__Rect__000000_rect_1.CanHaveLeftHandle = false
	__Rect__000000_rect_1.HasLeftHandle = false
	__Rect__000000_rect_1.CanHaveRightHandle = false
	__Rect__000000_rect_1.HasRightHandle = false
	__Rect__000000_rect_1.CanHaveTopHandle = false
	__Rect__000000_rect_1.HasTopHandle = false
	__Rect__000000_rect_1.IsScalingProportionally = false
	__Rect__000000_rect_1.CanHaveBottomHandle = false
	__Rect__000000_rect_1.HasBottomHandle = false
	__Rect__000000_rect_1.CanMoveHorizontaly = false
	__Rect__000000_rect_1.CanMoveVerticaly = false

	__SVG__000000_svg.Name = `svg`
	__SVG__000000_svg.IsEditable = false
	__SVG__000000_svg.IsSVGFrontEndFileGenerated = false
	__SVG__000000_svg.IsSVGBackEndFileGenerated = false
	__SVG__000000_svg.DefaultDirectoryForGeneratedImages = ``

	// Setup of pointers
	// setup of Layer instances pointers
	__Layer__000000_layer_1.Rects = append(__Layer__000000_layer_1.Rects, __Rect__000000_rect_1)
	// setup of Rect instances pointers
	// setup of SVG instances pointers
	__SVG__000000_svg.Layers = append(__SVG__000000_svg.Layers, __Layer__000000_layer_1)
}
