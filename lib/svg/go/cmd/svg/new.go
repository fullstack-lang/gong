package main

import (
	"time"

	"github.com/fullstack-lang/gong/lib/svg/go/models"
	// injection point for ident package import declaration{{ImportPackageDeclaration}}
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var ___dummy__Stage_new models.StageStruct
var ___dummy__Time_new time.Time

// Injection point for meta package dummy declaration{{ImportPackageDummyDeclaration}}

// currently, DocLink renaming is not enabled in gopls
// the following map are devised to overcome this limitation
// those maps and the processing code will be eleminated when
// DocLink renaming will be enabled in gopls
// [Corresponding Issue](https://github.com/golang/go/issues/57559)
//
// When parsed, those maps will help with the renaming process
var map_DocLink_Identifier_new map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// init might be handy if one want to have the data embedded in the binary
// but it has to properly reference the Injection gateway in the main package
// func init() {
// 	_ = __Dummy_time_variable
// 	InjectionGateway["new"] = newInjection
// }

// newInjection will stage objects of database "new"
func newInjection(stage *models.StageStruct) {

	// Declaration of instances to stage

	// Declarations of staged instances of Animate

	// Declarations of staged instances of Circle

	// Declarations of staged instances of Ellipse

	// Declarations of staged instances of Layer
	__Layer__000000_Paths := (&models.Layer{Name: `Paths`}).Stage(stage)

	// Declarations of staged instances of Line

	// Declarations of staged instances of Link

	// Declarations of staged instances of LinkAnchoredText

	// Declarations of staged instances of Path

	// Declarations of staged instances of Point

	// Declarations of staged instances of Polygone

	// Declarations of staged instances of Polyline

	// Declarations of staged instances of Rect
	__Rect__000000_Center_100_100 := (&models.Rect{Name: `Center 100 100`}).Stage(stage)
	__Rect__000001_Center_100_300 := (&models.Rect{Name: `Center 100 300`}).Stage(stage)
	__Rect__000002_Center_300_300 := (&models.Rect{Name: `Center 300 300`}).Stage(stage)

	// Declarations of staged instances of RectAnchoredPath
	__RectAnchoredPath__000000_Bottom_Bottom_Left := (&models.RectAnchoredPath{Name: `Bottom Bottom Left`}).Stage(stage)
	__RectAnchoredPath__000001_Bottom_Left_Left := (&models.RectAnchoredPath{Name: `Bottom Left Left`}).Stage(stage)
	__RectAnchoredPath__000002_Bottom_Right := (&models.RectAnchoredPath{Name: `Bottom Right`}).Stage(stage)
	__RectAnchoredPath__000003_Inside_at_the_Right := (&models.RectAnchoredPath{Name: `Inside at the Right`}).Stage(stage)
	__RectAnchoredPath__000004_Ref_anchored_path := (&models.RectAnchoredPath{Name: `Ref anchored path`}).Stage(stage)

	// Declarations of staged instances of RectAnchoredRect

	// Declarations of staged instances of RectAnchoredText

	// Declarations of staged instances of RectLinkLink

	// Declarations of staged instances of SVG
	__SVG__000000_SVG := (&models.SVG{Name: `SVG`}).Stage(stage)

	// Declarations of staged instances of Text

	// Setup of values

	// Layer values setup
	__Layer__000000_Paths.Display = true
	__Layer__000000_Paths.Name = `Paths`

	// Rect values setup
	__Rect__000000_Center_100_100.Name = `Center 100 100`
	__Rect__000000_Center_100_100.X = 146.000000
	__Rect__000000_Center_100_100.Y = 138.000000
	__Rect__000000_Center_100_100.Width = 254.000000
	__Rect__000000_Center_100_100.Height = 254.000000
	__Rect__000000_Center_100_100.RX = 4.000000
	__Rect__000000_Center_100_100.Color = ``
	__Rect__000000_Center_100_100.FillOpacity = 0.000000
	__Rect__000000_Center_100_100.Stroke = `blue`
	__Rect__000000_Center_100_100.StrokeWidth = 2.000000
	__Rect__000000_Center_100_100.StrokeDashArray = ``
	__Rect__000000_Center_100_100.StrokeDashArrayWhenSelected = ``
	__Rect__000000_Center_100_100.Transform = ``
	__Rect__000000_Center_100_100.IsSelectable = true
	__Rect__000000_Center_100_100.IsSelected = false
	__Rect__000000_Center_100_100.CanHaveLeftHandle = true
	__Rect__000000_Center_100_100.HasLeftHandle = false
	__Rect__000000_Center_100_100.CanHaveRightHandle = true
	__Rect__000000_Center_100_100.HasRightHandle = false
	__Rect__000000_Center_100_100.CanHaveTopHandle = true
	__Rect__000000_Center_100_100.HasTopHandle = false
	__Rect__000000_Center_100_100.IsScalingProportionally = true
	__Rect__000000_Center_100_100.CanHaveBottomHandle = true
	__Rect__000000_Center_100_100.HasBottomHandle = false
	__Rect__000000_Center_100_100.CanMoveHorizontaly = true
	__Rect__000000_Center_100_100.CanMoveVerticaly = true

	// Rect values setup
	__Rect__000001_Center_100_300.Name = `Center 100 300`
	__Rect__000001_Center_100_300.X = 469.000000
	__Rect__000001_Center_100_300.Y = 177.999969
	__Rect__000001_Center_100_300.Width = 380.000000
	__Rect__000001_Center_100_300.Height = 240.000000
	__Rect__000001_Center_100_300.RX = 0.000000
	__Rect__000001_Center_100_300.Color = `lavender`
	__Rect__000001_Center_100_300.FillOpacity = 0.800000
	__Rect__000001_Center_100_300.Stroke = `blue`
	__Rect__000001_Center_100_300.StrokeWidth = 1.000000
	__Rect__000001_Center_100_300.StrokeDashArray = ``
	__Rect__000001_Center_100_300.StrokeDashArrayWhenSelected = ``
	__Rect__000001_Center_100_300.Transform = ``
	__Rect__000001_Center_100_300.IsSelectable = true
	__Rect__000001_Center_100_300.IsSelected = false
	__Rect__000001_Center_100_300.CanHaveLeftHandle = false
	__Rect__000001_Center_100_300.HasLeftHandle = false
	__Rect__000001_Center_100_300.CanHaveRightHandle = false
	__Rect__000001_Center_100_300.HasRightHandle = false
	__Rect__000001_Center_100_300.CanHaveTopHandle = false
	__Rect__000001_Center_100_300.HasTopHandle = false
	__Rect__000001_Center_100_300.IsScalingProportionally = true
	__Rect__000001_Center_100_300.CanHaveBottomHandle = false
	__Rect__000001_Center_100_300.HasBottomHandle = false
	__Rect__000001_Center_100_300.CanMoveHorizontaly = true
	__Rect__000001_Center_100_300.CanMoveVerticaly = true

	// Rect values setup
	__Rect__000002_Center_300_300.Name = `Center 300 300`
	__Rect__000002_Center_300_300.X = 341.000000
	__Rect__000002_Center_300_300.Y = 480.000000
	__Rect__000002_Center_300_300.Width = 240.000000
	__Rect__000002_Center_300_300.Height = 240.000000
	__Rect__000002_Center_300_300.RX = 5.000000
	__Rect__000002_Center_300_300.Color = ``
	__Rect__000002_Center_300_300.FillOpacity = 0.000000
	__Rect__000002_Center_300_300.Stroke = `blue`
	__Rect__000002_Center_300_300.StrokeWidth = 1.000000
	__Rect__000002_Center_300_300.StrokeDashArray = ``
	__Rect__000002_Center_300_300.StrokeDashArrayWhenSelected = ``
	__Rect__000002_Center_300_300.Transform = ``
	__Rect__000002_Center_300_300.IsSelectable = true
	__Rect__000002_Center_300_300.IsSelected = false
	__Rect__000002_Center_300_300.CanHaveLeftHandle = false
	__Rect__000002_Center_300_300.HasLeftHandle = false
	__Rect__000002_Center_300_300.CanHaveRightHandle = true
	__Rect__000002_Center_300_300.HasRightHandle = false
	__Rect__000002_Center_300_300.CanHaveTopHandle = false
	__Rect__000002_Center_300_300.HasTopHandle = false
	__Rect__000002_Center_300_300.IsScalingProportionally = true
	__Rect__000002_Center_300_300.CanHaveBottomHandle = false
	__Rect__000002_Center_300_300.HasBottomHandle = false
	__Rect__000002_Center_300_300.CanMoveHorizontaly = true
	__Rect__000002_Center_300_300.CanMoveVerticaly = true

	// RectAnchoredPath values setup
	__RectAnchoredPath__000000_Bottom_Bottom_Left.Name = `Bottom Bottom Left`
	__RectAnchoredPath__000000_Bottom_Bottom_Left.Definition = `M 0,-960 V 0 h 960 v -960 z m 40,40 H 920 V -40 H 40 Z`
	__RectAnchoredPath__000000_Bottom_Bottom_Left.X_Offset = 0.000000
	__RectAnchoredPath__000000_Bottom_Bottom_Left.Y_Offset = 0.000000
	__RectAnchoredPath__000000_Bottom_Bottom_Left.RectAnchorType = models.RECT_BOTTOM_BOTTOM_LEFT
	__RectAnchoredPath__000000_Bottom_Bottom_Left.ScalePropotionnally = true
	__RectAnchoredPath__000000_Bottom_Bottom_Left.AppliedScaling = 0.250000
	__RectAnchoredPath__000000_Bottom_Bottom_Left.Color = ``
	__RectAnchoredPath__000000_Bottom_Bottom_Left.FillOpacity = 0.000000
	__RectAnchoredPath__000000_Bottom_Bottom_Left.Stroke = `darkturquoise`
	__RectAnchoredPath__000000_Bottom_Bottom_Left.StrokeWidth = 7.000000
	__RectAnchoredPath__000000_Bottom_Bottom_Left.StrokeDashArray = ``
	__RectAnchoredPath__000000_Bottom_Bottom_Left.StrokeDashArrayWhenSelected = ``
	__RectAnchoredPath__000000_Bottom_Bottom_Left.Transform = ``

	// RectAnchoredPath values setup
	__RectAnchoredPath__000001_Bottom_Left_Left.Name = `Bottom Left Left`
	__RectAnchoredPath__000001_Bottom_Left_Left.Definition = `M 0,-960 V 0 h 960 v -960 z m 40,40 H 920 V -40 H 40 Z`
	__RectAnchoredPath__000001_Bottom_Left_Left.X_Offset = 0.000000
	__RectAnchoredPath__000001_Bottom_Left_Left.Y_Offset = 0.000000
	__RectAnchoredPath__000001_Bottom_Left_Left.RectAnchorType = models.RECT_BOTTOM_LEFT_LEFT
	__RectAnchoredPath__000001_Bottom_Left_Left.ScalePropotionnally = true
	__RectAnchoredPath__000001_Bottom_Left_Left.AppliedScaling = 0.250000
	__RectAnchoredPath__000001_Bottom_Left_Left.Color = ``
	__RectAnchoredPath__000001_Bottom_Left_Left.FillOpacity = 0.000000
	__RectAnchoredPath__000001_Bottom_Left_Left.Stroke = `lavender`
	__RectAnchoredPath__000001_Bottom_Left_Left.StrokeWidth = 6.000000
	__RectAnchoredPath__000001_Bottom_Left_Left.StrokeDashArray = ``
	__RectAnchoredPath__000001_Bottom_Left_Left.StrokeDashArrayWhenSelected = ``
	__RectAnchoredPath__000001_Bottom_Left_Left.Transform = ``

	// RectAnchoredPath values setup
	__RectAnchoredPath__000002_Bottom_Right.Name = `Bottom Right`
	__RectAnchoredPath__000002_Bottom_Right.Definition = `M 0,-960 V 0 h 960 v -960 z m 40,40 H 920 V -40 H 40 Z`
	__RectAnchoredPath__000002_Bottom_Right.X_Offset = 0.000000
	__RectAnchoredPath__000002_Bottom_Right.Y_Offset = 0.000000
	__RectAnchoredPath__000002_Bottom_Right.RectAnchorType = models.RECT_BOTTOM_RIGHT
	__RectAnchoredPath__000002_Bottom_Right.ScalePropotionnally = true
	__RectAnchoredPath__000002_Bottom_Right.AppliedScaling = 0.250000
	__RectAnchoredPath__000002_Bottom_Right.Color = ``
	__RectAnchoredPath__000002_Bottom_Right.FillOpacity = 0.000000
	__RectAnchoredPath__000002_Bottom_Right.Stroke = `red`
	__RectAnchoredPath__000002_Bottom_Right.StrokeWidth = 8.000000
	__RectAnchoredPath__000002_Bottom_Right.StrokeDashArray = ``
	__RectAnchoredPath__000002_Bottom_Right.StrokeDashArrayWhenSelected = ``
	__RectAnchoredPath__000002_Bottom_Right.Transform = ``

	// RectAnchoredPath values setup
	__RectAnchoredPath__000003_Inside_at_the_Right.Name = `Inside at the Right`
	__RectAnchoredPath__000003_Inside_at_the_Right.Definition = `M 0,-960 V 0 h 960 v -960 z m 40,40 H 920 V -40 H 40 Z`
	__RectAnchoredPath__000003_Inside_at_the_Right.X_Offset = 0.000000
	__RectAnchoredPath__000003_Inside_at_the_Right.Y_Offset = 0.000000
	__RectAnchoredPath__000003_Inside_at_the_Right.RectAnchorType = models.RECT_BOTTOM_LEFT_LEFT
	__RectAnchoredPath__000003_Inside_at_the_Right.ScalePropotionnally = true
	__RectAnchoredPath__000003_Inside_at_the_Right.AppliedScaling = 0.250000
	__RectAnchoredPath__000003_Inside_at_the_Right.Color = ``
	__RectAnchoredPath__000003_Inside_at_the_Right.FillOpacity = 0.000000
	__RectAnchoredPath__000003_Inside_at_the_Right.Stroke = `grey`
	__RectAnchoredPath__000003_Inside_at_the_Right.StrokeWidth = 3.000000
	__RectAnchoredPath__000003_Inside_at_the_Right.StrokeDashArray = ``
	__RectAnchoredPath__000003_Inside_at_the_Right.StrokeDashArrayWhenSelected = ``
	__RectAnchoredPath__000003_Inside_at_the_Right.Transform = ``

	// RectAnchoredPath values setup
	__RectAnchoredPath__000004_Ref_anchored_path.Name = `Ref anchored path`
	__RectAnchoredPath__000004_Ref_anchored_path.Definition = `M 0,-960 V 0 h 960 v -960 z m 40,40 H 920 V -40 H 40 Z`
	__RectAnchoredPath__000004_Ref_anchored_path.X_Offset = 20.000000
	__RectAnchoredPath__000004_Ref_anchored_path.Y_Offset = 0.000000
	__RectAnchoredPath__000004_Ref_anchored_path.RectAnchorType = models.RECT_BOTTOM_LEFT
	__RectAnchoredPath__000004_Ref_anchored_path.ScalePropotionnally = true
	__RectAnchoredPath__000004_Ref_anchored_path.AppliedScaling = 0.264584
	__RectAnchoredPath__000004_Ref_anchored_path.Color = ``
	__RectAnchoredPath__000004_Ref_anchored_path.FillOpacity = 0.000000
	__RectAnchoredPath__000004_Ref_anchored_path.Stroke = `lightblue`
	__RectAnchoredPath__000004_Ref_anchored_path.StrokeWidth = 5.000000
	__RectAnchoredPath__000004_Ref_anchored_path.StrokeDashArray = ``
	__RectAnchoredPath__000004_Ref_anchored_path.StrokeDashArrayWhenSelected = ``
	__RectAnchoredPath__000004_Ref_anchored_path.Transform = ``

	// SVG values setup
	__SVG__000000_SVG.Name = `SVG`
	__SVG__000000_SVG.IsEditable = true

	// Setup of pointers
	__Layer__000000_Paths.Rects = append(__Layer__000000_Paths.Rects, __Rect__000000_Center_100_100)
	__Layer__000000_Paths.Rects = append(__Layer__000000_Paths.Rects, __Rect__000002_Center_300_300)
	__Layer__000000_Paths.Rects = append(__Layer__000000_Paths.Rects, __Rect__000001_Center_100_300)
	__Rect__000000_Center_100_100.RectAnchoredPaths = append(__Rect__000000_Center_100_100.RectAnchoredPaths, __RectAnchoredPath__000004_Ref_anchored_path)
	__Rect__000001_Center_100_300.RectAnchoredPaths = append(__Rect__000001_Center_100_300.RectAnchoredPaths, __RectAnchoredPath__000003_Inside_at_the_Right)
	__Rect__000002_Center_300_300.RectAnchoredPaths = append(__Rect__000002_Center_300_300.RectAnchoredPaths, __RectAnchoredPath__000002_Bottom_Right)
	__Rect__000002_Center_300_300.RectAnchoredPaths = append(__Rect__000002_Center_300_300.RectAnchoredPaths, __RectAnchoredPath__000001_Bottom_Left_Left)
	__Rect__000002_Center_300_300.RectAnchoredPaths = append(__Rect__000002_Center_300_300.RectAnchoredPaths, __RectAnchoredPath__000000_Bottom_Bottom_Left)
	__SVG__000000_SVG.Layers = append(__SVG__000000_SVG.Layers, __Layer__000000_Paths)
}
