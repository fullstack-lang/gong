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

	__Layer__000000_Path := (&models.Layer{}).Stage(stage)

	__Path__000000_Path := (&models.Path{}).Stage(stage)

	__Rect__000000_Host := (&models.Rect{}).Stage(stage)

	__RectAnchoredPath__000000_Anchored_Path := (&models.RectAnchoredPath{}).Stage(stage)

	__SVG__000000_svg := (&models.SVG{}).Stage(stage)

	// Setup of values

	__Layer__000000_Path.Name = `Path`

	__Path__000000_Path.Name = `Path`
	__Path__000000_Path.Definition = `M 0,41 30,11 20,1 l 30,0 V 31 L 40,21 10,51 Z`
	__Path__000000_Path.Color = ``
	__Path__000000_Path.FillOpacity = 0.000000
	__Path__000000_Path.Stroke = `blue`
	__Path__000000_Path.StrokeOpacity = 1.000000
	__Path__000000_Path.StrokeWidth = 1.000000
	__Path__000000_Path.StrokeDashArray = ``
	__Path__000000_Path.StrokeDashArrayWhenSelected = ``
	__Path__000000_Path.Transform = `scale(5 5)`

	__Rect__000000_Host.Name = `Host`
	__Rect__000000_Host.X = 0.000000
	__Rect__000000_Host.Y = 0.000000
	__Rect__000000_Host.Width = 0.000000
	__Rect__000000_Host.Height = 0.000000
	__Rect__000000_Host.RX = 0.000000
	__Rect__000000_Host.Color = ``
	__Rect__000000_Host.FillOpacity = 0.000000
	__Rect__000000_Host.Stroke = ``
	__Rect__000000_Host.StrokeOpacity = 0.000000
	__Rect__000000_Host.StrokeWidth = 0.000000
	__Rect__000000_Host.StrokeDashArray = ``
	__Rect__000000_Host.StrokeDashArrayWhenSelected = ``
	__Rect__000000_Host.Transform = ``
	__Rect__000000_Host.IsSelectable = false
	__Rect__000000_Host.IsSelected = false
	__Rect__000000_Host.CanHaveLeftHandle = false
	__Rect__000000_Host.HasLeftHandle = false
	__Rect__000000_Host.CanHaveRightHandle = false
	__Rect__000000_Host.HasRightHandle = false
	__Rect__000000_Host.CanHaveTopHandle = false
	__Rect__000000_Host.HasTopHandle = false
	__Rect__000000_Host.IsScalingProportionally = false
	__Rect__000000_Host.CanHaveBottomHandle = false
	__Rect__000000_Host.HasBottomHandle = false
	__Rect__000000_Host.CanMoveHorizontaly = false
	__Rect__000000_Host.CanMoveVerticaly = false
	__Rect__000000_Host.ChangeColorWhenHovered = false
	__Rect__000000_Host.ColorWhenHovered = ``
	__Rect__000000_Host.OriginalColor = ``
	__Rect__000000_Host.FillOpacityWhenHovered = 0.000000
	__Rect__000000_Host.OriginalFillOpacity = 0.000000
	__Rect__000000_Host.CheckboxHasToolTip = false
	__Rect__000000_Host.CheckboxToolTipText = ``

	__RectAnchoredPath__000000_Anchored_Path.Name = `Anchored Path`
	__RectAnchoredPath__000000_Anchored_Path.Definition = `M 0,41 30,11 20,1 l 30,0 V 31 L 40,21 10,51 Z`
	__RectAnchoredPath__000000_Anchored_Path.X_Offset = 0.000000
	__RectAnchoredPath__000000_Anchored_Path.Y_Offset = 0.000000
	__RectAnchoredPath__000000_Anchored_Path.ScalePropotionnally = true
	__RectAnchoredPath__000000_Anchored_Path.AppliedScaling = 6.000000
	__RectAnchoredPath__000000_Anchored_Path.Color = ``
	__RectAnchoredPath__000000_Anchored_Path.FillOpacity = 0.000000
	__RectAnchoredPath__000000_Anchored_Path.Stroke = `green`
	__RectAnchoredPath__000000_Anchored_Path.StrokeOpacity = 1.000000
	__RectAnchoredPath__000000_Anchored_Path.StrokeWidth = 1.000000
	__RectAnchoredPath__000000_Anchored_Path.StrokeDashArray = ``
	__RectAnchoredPath__000000_Anchored_Path.StrokeDashArrayWhenSelected = ``
	__RectAnchoredPath__000000_Anchored_Path.Transform = ``

	__SVG__000000_svg.Name = `svg`
	__SVG__000000_svg.IsEditable = false
	__SVG__000000_svg.IsSVGFrontEndFileGenerated = false
	__SVG__000000_svg.IsSVGBackEndFileGenerated = false
	__SVG__000000_svg.DefaultDirectoryForGeneratedImages = `../../diagrams/images`

	// Setup of pointers
	// setup of Layer instances pointers
	__Layer__000000_Path.Rects = append(__Layer__000000_Path.Rects, __Rect__000000_Host)
	__Layer__000000_Path.Paths = append(__Layer__000000_Path.Paths, __Path__000000_Path)
	// setup of Path instances pointers
	// setup of Rect instances pointers
	__Rect__000000_Host.RectAnchoredPaths = append(__Rect__000000_Host.RectAnchoredPaths, __RectAnchoredPath__000000_Anchored_Path)
	// setup of RectAnchoredPath instances pointers
	// setup of SVG instances pointers
	__SVG__000000_svg.Layers = append(__SVG__000000_svg.Layers, __Layer__000000_Path)
}
