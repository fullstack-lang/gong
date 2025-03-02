package main

import (
	"time"

	"github.com/fullstack-lang/gongsvg/go/models"

	// injection point for ident package import declaration{{ImportPackageDeclaration}}
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var ___dummy__Stage_bezier models.StageStruct
var ___dummy__Time_bezier time.Time

// Injection point for meta package dummy declaration{{ImportPackageDummyDeclaration}}

// currently, DocLink renaming is not enabled in gopls
// the following map are devised to overcome this limitation
// those maps and the processing code will be eleminated when
// DocLink renaming will be enabled in gopls
// [Corresponding Issue](https://github.com/golang/go/issues/57559)
//
// When parsed, those maps will help with the renaming process
var map_DocLink_Identifier_bezier map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// init might be handy if one want to have the data embedded in the binary
// but it has to properly reference the Injection gateway in the main package
// func init() {
// 	_ = __Dummy_time_variable
// 	InjectionGateway["bezier"] = bezierInjection
// }

// bezierInjection will stage objects of database "bezier"
func bezierInjection(stage *models.StageStruct) {

	// Declaration of instances to stage

	// Declarations of staged instances of Animate

	// Declarations of staged instances of Circle

	// Declarations of staged instances of Ellipse

	// Declarations of staged instances of Layer
	__Layer__000000_bezier := (&models.Layer{Name: `bezier`}).Stage(stage)

	// Declarations of staged instances of Line

	// Declarations of staged instances of Link
	__Link__000000_Test := (&models.Link{Name: `Test`}).Stage(stage)

	// Declarations of staged instances of LinkAnchoredText

	// Declarations of staged instances of Path
	__Path__000000_bezier := (&models.Path{Name: `bezier`}).Stage(stage)
	__Path__000001_bezier_2 := (&models.Path{Name: `bezier 2`}).Stage(stage)

	// Declarations of staged instances of Point

	// Declarations of staged instances of Polygone

	// Declarations of staged instances of Polyline

	// Declarations of staged instances of Rect
	__Rect__000000_End := (&models.Rect{Name: `End`}).Stage(stage)
	__Rect__000001_Start := (&models.Rect{Name: `Start`}).Stage(stage)

	// Declarations of staged instances of RectAnchoredPath

	// Declarations of staged instances of RectAnchoredRect

	// Declarations of staged instances of RectAnchoredText

	// Declarations of staged instances of RectLinkLink

	// Declarations of staged instances of SVG
	__SVG__000000_beizer := (&models.SVG{Name: `beizer`}).Stage(stage)

	// Declarations of staged instances of Text

	// Setup of values

	// Layer values setup
	__Layer__000000_bezier.Display = true
	__Layer__000000_bezier.Name = `bezier`

	// Link values setup
	__Link__000000_Test.Name = `Test`
	__Link__000000_Test.Type = models.LINK_TYPE_FLOATING_ORTHOGONAL
	__Link__000000_Test.IsBezierCurve = true
	__Link__000000_Test.StartOrientation = models.ORIENTATION_VERTICAL
	__Link__000000_Test.StartRatio = 0.530039
	__Link__000000_Test.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000000_Test.EndRatio = 0.000000
	__Link__000000_Test.CornerOffsetRatio = 4.240000
	__Link__000000_Test.CornerRadius = 20.000000
	__Link__000000_Test.HasEndArrow = true
	__Link__000000_Test.EndArrowSize = 15.000000
	__Link__000000_Test.HasStartArrow = false
	__Link__000000_Test.StartArrowSize = 0.000000
	__Link__000000_Test.Color = ``
	__Link__000000_Test.FillOpacity = 0.000000
	__Link__000000_Test.Stroke = `black`
	__Link__000000_Test.StrokeWidth = 5.000000
	__Link__000000_Test.StrokeDashArray = ``
	__Link__000000_Test.StrokeDashArrayWhenSelected = ``
	__Link__000000_Test.Transform = ``

	// Path values setup
	__Path__000000_bezier.Name = `bezier`
	__Path__000000_bezier.Definition = `M 0, 0 c 
100, 0 
100, 50
 200, 50`
	__Path__000000_bezier.Color = ``
	__Path__000000_bezier.FillOpacity = 0.000000
	__Path__000000_bezier.Stroke = `black`
	__Path__000000_bezier.StrokeWidth = 1.000000
	__Path__000000_bezier.StrokeDashArray = ``
	__Path__000000_bezier.StrokeDashArrayWhenSelected = ``
	__Path__000000_bezier.Transform = `scale(4 4)`

	// Path values setup
	__Path__000001_bezier_2.Name = `bezier 2`
	__Path__000001_bezier_2.Definition = `M 0, 0 c 
150, 0 
50, 50
 200, 50`
	__Path__000001_bezier_2.Color = ``
	__Path__000001_bezier_2.FillOpacity = 0.000000
	__Path__000001_bezier_2.Stroke = `black`
	__Path__000001_bezier_2.StrokeWidth = 1.000000
	__Path__000001_bezier_2.StrokeDashArray = ``
	__Path__000001_bezier_2.StrokeDashArrayWhenSelected = ``
	__Path__000001_bezier_2.Transform = `translate(0 200) scale(4 4)`

	// Rect values setup
	__Rect__000000_End.Name = `End`
	__Rect__000000_End.X = 513.000000
	__Rect__000000_End.Y = 706.000000
	__Rect__000000_End.Width = 100.000000
	__Rect__000000_End.Height = 100.000000
	__Rect__000000_End.RX = 0.000000
	__Rect__000000_End.Color = `grey`
	__Rect__000000_End.FillOpacity = 1.000000
	__Rect__000000_End.Stroke = ``
	__Rect__000000_End.StrokeWidth = 0.000000
	__Rect__000000_End.StrokeDashArray = ``
	__Rect__000000_End.StrokeDashArrayWhenSelected = ``
	__Rect__000000_End.Transform = ``
	__Rect__000000_End.IsSelectable = false
	__Rect__000000_End.IsSelected = false
	__Rect__000000_End.CanHaveLeftHandle = false
	__Rect__000000_End.HasLeftHandle = false
	__Rect__000000_End.CanHaveRightHandle = false
	__Rect__000000_End.HasRightHandle = false
	__Rect__000000_End.CanHaveTopHandle = false
	__Rect__000000_End.HasTopHandle = false
	__Rect__000000_End.IsScalingProportionally = false
	__Rect__000000_End.CanHaveBottomHandle = false
	__Rect__000000_End.HasBottomHandle = false
	__Rect__000000_End.CanMoveHorizontaly = true
	__Rect__000000_End.CanMoveVerticaly = true

	// Rect values setup
	__Rect__000001_Start.Name = `Start`
	__Rect__000001_Start.X = 36.000000
	__Rect__000001_Start.Y = 506.000000
	__Rect__000001_Start.Width = 100.000000
	__Rect__000001_Start.Height = 100.000000
	__Rect__000001_Start.RX = 0.000000
	__Rect__000001_Start.Color = `grey`
	__Rect__000001_Start.FillOpacity = 1.000000
	__Rect__000001_Start.Stroke = ``
	__Rect__000001_Start.StrokeWidth = 0.000000
	__Rect__000001_Start.StrokeDashArray = ``
	__Rect__000001_Start.StrokeDashArrayWhenSelected = ``
	__Rect__000001_Start.Transform = ``
	__Rect__000001_Start.IsSelectable = false
	__Rect__000001_Start.IsSelected = false
	__Rect__000001_Start.CanHaveLeftHandle = false
	__Rect__000001_Start.HasLeftHandle = false
	__Rect__000001_Start.CanHaveRightHandle = false
	__Rect__000001_Start.HasRightHandle = false
	__Rect__000001_Start.CanHaveTopHandle = false
	__Rect__000001_Start.HasTopHandle = false
	__Rect__000001_Start.IsScalingProportionally = false
	__Rect__000001_Start.CanHaveBottomHandle = false
	__Rect__000001_Start.HasBottomHandle = false
	__Rect__000001_Start.CanMoveHorizontaly = true
	__Rect__000001_Start.CanMoveVerticaly = true

	// SVG values setup
	__SVG__000000_beizer.Name = `beizer`
	__SVG__000000_beizer.IsEditable = true

	// Setup of pointers
	__Layer__000000_bezier.Rects = append(__Layer__000000_bezier.Rects, __Rect__000001_Start)
	__Layer__000000_bezier.Rects = append(__Layer__000000_bezier.Rects, __Rect__000000_End)
	__Layer__000000_bezier.Paths = append(__Layer__000000_bezier.Paths, __Path__000000_bezier)
	__Layer__000000_bezier.Paths = append(__Layer__000000_bezier.Paths, __Path__000001_bezier_2)
	__Layer__000000_bezier.Links = append(__Layer__000000_bezier.Links, __Link__000000_Test)
	__Link__000000_Test.Start = __Rect__000001_Start
	__Link__000000_Test.End = __Rect__000000_End
	__SVG__000000_beizer.Layers = append(__SVG__000000_beizer.Layers, __Layer__000000_bezier)
}


