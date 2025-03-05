package main

import (
	"time"

	"github.com/fullstack-lang/gong/lib/svg/go/models"
	// injection point for ident package import declaration{{ImportPackageDeclaration}}
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var ___dummy__Stage_issue104 models.StageStruct
var ___dummy__Time_issue104 time.Time

// Injection point for meta package dummy declaration{{ImportPackageDummyDeclaration}}

// currently, DocLink renaming is not enabled in gopls
// the following map are devised to overcome this limitation
// those maps and the processing code will be eleminated when
// DocLink renaming will be enabled in gopls
// [Corresponding Issue](https://github.com/golang/go/issues/57559)
//
// When parsed, those maps will help with the renaming process
var map_DocLink_Identifier_issue104 map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// init might be handy if one want to have the data embedded in the binary
// but it has to properly reference the Injection gateway in the main package
// func init() {
// 	_ = __Dummy_time_variable
// 	InjectionGateway["issue104"] = issue104Injection
// }

// issue104Injection will stage objects of database "issue104"
func issue104Injection(stage *models.StageStruct) {

	// Declaration of instances to stage

	// Declarations of staged instances of AnchoredText

	// Declarations of staged instances of Animate

	// Declarations of staged instances of Circle

	// Declarations of staged instances of Ellipse

	// Declarations of staged instances of Layer
	__Layer__000000_104 := (&models.Layer{Name: `104`}).Stage(stage)

	// Declarations of staged instances of Line

	// Declarations of staged instances of Link

	// Declarations of staged instances of Path

	// Declarations of staged instances of Point

	// Declarations of staged instances of Polygone

	// Declarations of staged instances of Polyline

	// Declarations of staged instances of Rect
	__Rect__000000_104 := (&models.Rect{Name: `104`}).Stage(stage)

	// Declarations of staged instances of RectAnchoredRect
	__RectAnchoredRect__000000_Test_104 := (&models.RectAnchoredRect{Name: `Test 104`}).Stage(stage)

	// Declarations of staged instances of RectAnchoredText
	__RectAnchoredText__000000_test := (&models.RectAnchoredText{Name: `test`}).Stage(stage)

	// Declarations of staged instances of RectLinkLink

	// Declarations of staged instances of SVG
	__SVG__000000_104 := (&models.SVG{Name: `104`}).Stage(stage)

	// Declarations of staged instances of Text

	// Setup of values

	// Layer values setup
	__Layer__000000_104.Display = false
	__Layer__000000_104.Name = `104`

	// Rect values setup
	__Rect__000000_104.Name = `104`
	__Rect__000000_104.X = 127.000000
	__Rect__000000_104.Y = 153.000000
	__Rect__000000_104.Width = 300.000000
	__Rect__000000_104.Height = 200.000000
	__Rect__000000_104.RX = 4.000000
	__Rect__000000_104.Color = `bisque`
	__Rect__000000_104.FillOpacity = 50.000000
	__Rect__000000_104.Stroke = `lavender`
	__Rect__000000_104.StrokeWidth = 4.000000
	__Rect__000000_104.StrokeDashArray = ``
	__Rect__000000_104.StrokeDashArrayWhenSelected = ``
	__Rect__000000_104.Transform = ``
	__Rect__000000_104.IsSelectable = true
	__Rect__000000_104.IsSelected = false
	__Rect__000000_104.CanHaveLeftHandle = false
	__Rect__000000_104.HasLeftHandle = false
	__Rect__000000_104.CanHaveRightHandle = false
	__Rect__000000_104.HasRightHandle = false
	__Rect__000000_104.CanHaveTopHandle = false
	__Rect__000000_104.HasTopHandle = false
	__Rect__000000_104.CanHaveBottomHandle = false
	__Rect__000000_104.HasBottomHandle = false
	__Rect__000000_104.CanMoveHorizontaly = true
	__Rect__000000_104.CanMoveVerticaly = true

	// RectAnchoredRect values setup
	__RectAnchoredRect__000000_Test_104.Name = `Test 104`
	__RectAnchoredRect__000000_Test_104.X = 0.000000
	__RectAnchoredRect__000000_Test_104.Y = 0.000000
	__RectAnchoredRect__000000_Test_104.Width = 0.000000
	__RectAnchoredRect__000000_Test_104.Height = 50.000000
	__RectAnchoredRect__000000_Test_104.RX = 0.000000
	__RectAnchoredRect__000000_Test_104.X_Offset = 0.000000
	__RectAnchoredRect__000000_Test_104.Y_Offset = -50.000000
	__RectAnchoredRect__000000_Test_104.RectAnchorType = models.RECT_TOP_LEFT
	__RectAnchoredRect__000000_Test_104.WidthFollowRect = true
	__RectAnchoredRect__000000_Test_104.HeightFollowRect = false
	__RectAnchoredRect__000000_Test_104.Color = `lightgrey`
	__RectAnchoredRect__000000_Test_104.FillOpacity = 40.000000
	__RectAnchoredRect__000000_Test_104.Stroke = `black`
	__RectAnchoredRect__000000_Test_104.StrokeWidth = 1.000000
	__RectAnchoredRect__000000_Test_104.StrokeDashArray = ``
	__RectAnchoredRect__000000_Test_104.StrokeDashArrayWhenSelected = ``
	__RectAnchoredRect__000000_Test_104.Transform = ``

	// RectAnchoredText values setup
	__RectAnchoredText__000000_test.Name = `test`
	__RectAnchoredText__000000_test.Content = `test`
	__RectAnchoredText__000000_test.FontWeight = ``
	__RectAnchoredText__000000_test.FontSize = 14
	__RectAnchoredText__000000_test.X_Offset = 0.000000
	__RectAnchoredText__000000_test.Y_Offset = -100.000000
	__RectAnchoredText__000000_test.RectAnchorType = models.RECT_TOP
	__RectAnchoredText__000000_test.TextAnchorType = models.TEXT_ANCHOR_CENTER
	__RectAnchoredText__000000_test.Color = ``
	__RectAnchoredText__000000_test.FillOpacity = 0.000000
	__RectAnchoredText__000000_test.Stroke = `black`
	__RectAnchoredText__000000_test.StrokeWidth = 1.000000
	__RectAnchoredText__000000_test.StrokeDashArray = ``
	__RectAnchoredText__000000_test.StrokeDashArrayWhenSelected = ``
	__RectAnchoredText__000000_test.Transform = ``

	// SVG values setup
	__SVG__000000_104.Name = `104`
	__SVG__000000_104.IsEditable = true

	// Setup of pointers
	__Layer__000000_104.Rects = append(__Layer__000000_104.Rects, __Rect__000000_104)
	__Rect__000000_104.RectAnchoredTexts = append(__Rect__000000_104.RectAnchoredTexts, __RectAnchoredText__000000_test)
	__Rect__000000_104.RectAnchoredRects = append(__Rect__000000_104.RectAnchoredRects, __RectAnchoredRect__000000_Test_104)
	__SVG__000000_104.Layers = append(__SVG__000000_104.Layers, __Layer__000000_104)
}
