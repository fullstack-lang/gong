package main

import (
	"time"

	"github.com/fullstack-lang/gongsvg/go/models"

	// injection point for ident package import declaration{{ImportPackageDeclaration}}
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var ___dummy__Stage_issue_triangle models.StageStruct
var ___dummy__Time_issue_triangle time.Time

// Injection point for meta package dummy declaration{{ImportPackageDummyDeclaration}}

// currently, DocLink renaming is not enabled in gopls
// the following map are devised to overcome this limitation
// those maps and the processing code will be eleminated when
// DocLink renaming will be enabled in gopls
// [Corresponding Issue](https://github.com/golang/go/issues/57559)
//
// When parsed, those maps will help with the renaming process
var map_DocLink_Identifier_issue_triangle map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// init might be handy if one want to have the data embedded in the binary
// but it has to properly reference the Injection gateway in the main package
// func init() {
// 	_ = __Dummy_time_variable
// 	InjectionGateway["issue_triangle"] = issue_triangleInjection
// }

// issue_triangleInjection will stage objects of database "issue_triangle"
func issue_triangleInjection(stage *models.StageStruct) {

	// Declaration of instances to stage

	// Declarations of staged instances of Animate

	// Declarations of staged instances of Circle

	// Declarations of staged instances of Ellipse

	// Declarations of staged instances of Layer
	__Layer__000000_ := (&models.Layer{Name: ``}).Stage(stage)

	// Declarations of staged instances of Line

	// Declarations of staged instances of Link

	// Declarations of staged instances of LinkAnchoredText

	// Declarations of staged instances of Path

	// Declarations of staged instances of Point

	// Declarations of staged instances of Polygone
	__Polygone__000000_ := (&models.Polygone{Name: ``}).Stage(stage)

	// Declarations of staged instances of Polyline
	__Polyline__000000_ := (&models.Polyline{Name: ``}).Stage(stage)

	// Declarations of staged instances of Rect

	// Declarations of staged instances of RectAnchoredPath

	// Declarations of staged instances of RectAnchoredRect

	// Declarations of staged instances of RectAnchoredText

	// Declarations of staged instances of RectLinkLink

	// Declarations of staged instances of SVG
	__SVG__000000_ := (&models.SVG{Name: ``}).Stage(stage)

	// Declarations of staged instances of Text

	// Setup of values

	// Layer values setup
	__Layer__000000_.Display = true
	__Layer__000000_.Name = ``

	// Polygone values setup
	__Polygone__000000_.Name = ``
	__Polygone__000000_.Points = ` 260.95 382.50 244.29 392.50 244.29 372.50`
	__Polygone__000000_.Color = `red`
	__Polygone__000000_.FillOpacity = 0.500000
	__Polygone__000000_.Stroke = `red`
	__Polygone__000000_.StrokeOpacity = 0.500000
	__Polygone__000000_.StrokeWidth = 0.000000
	__Polygone__000000_.StrokeDashArray = ``
	__Polygone__000000_.StrokeDashArrayWhenSelected = ``
	__Polygone__000000_.Transform = ``

	// Polyline values setup
	__Polyline__000000_.Name = ``
	__Polyline__000000_.Points = ` 30.00 382.50 30.20 382.50 30.67 382.50 31.62 382.50 33.28 382.50 35.89 382.50 39.55 382.50 44.14 382.50 49.28 382.50 54.53 382.50 59.54 382.50 64.14 382.50 68.30 382.50 72.08 382.50 75.55 382.50 78.80 382.50 81.88 382.50 84.85 382.50 87.75 382.50 90.59 382.50 93.40 382.50 96.19 382.50 98.97 382.50 101.73 382.50 104.49 382.50 107.24 382.50 109.99 382.50 112.74 382.50 115.49 382.50 118.24 382.50 120.99 382.50 123.74 382.50 126.48 382.50 129.23 382.50 131.98 382.50 134.73 382.50 137.47 382.50 140.22 382.50 142.97 382.50 145.71 382.50 148.46 382.50 151.21 382.50 153.96 382.50 156.70 382.50 159.45 382.50 162.20 382.50 164.94 382.50 167.69 382.50 170.44 382.50 173.19 382.50 175.93 382.50 178.68 382.50 181.43 382.50 184.17 382.50 186.92 382.50 189.66 382.50 192.40 382.50 195.14 382.50 197.88 382.50 200.61 382.50 203.34 382.50 206.05 382.50 208.76 382.50 211.45 382.50 214.12 382.50 216.76 382.50 219.38 382.50 221.95 382.50 224.47 382.50 226.92 382.50 229.30 382.50 231.60 382.50 233.80 382.50 235.88 382.50 237.84 382.50 239.67 382.50 241.35 382.50 242.89 382.50 244.29 382.50`
	__Polyline__000000_.Color = ``
	__Polyline__000000_.FillOpacity = 0.000000
	__Polyline__000000_.Stroke = `red`
	__Polyline__000000_.StrokeOpacity = 0.500000
	__Polyline__000000_.StrokeWidth = 20.000000
	__Polyline__000000_.StrokeDashArray = ``
	__Polyline__000000_.StrokeDashArrayWhenSelected = ``
	__Polyline__000000_.Transform = ``

	// SVG values setup
	__SVG__000000_.Name = ``
	__SVG__000000_.IsEditable = false

	// Setup of pointers
	__Layer__000000_.Polylines = append(__Layer__000000_.Polylines, __Polyline__000000_)
	__Layer__000000_.Polygones = append(__Layer__000000_.Polygones, __Polygone__000000_)
	__SVG__000000_.Layers = append(__SVG__000000_.Layers, __Layer__000000_)
}


