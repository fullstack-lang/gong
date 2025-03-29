package main

import (
	"time"

	"github.com/fullstack-lang/gong/lib/svg/go/models"
	// injection point for ident package import declaration{{ImportPackageDeclaration}}
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var _ time.Time

// Injection point for meta package dummy declaration{{ImportPackageDummyDeclaration}}

// When parsed, those maps will help with the renaming process
var _ map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// function will stage objects
func _(stage *models.Stage) {

	// Declaration of instances to stage

	__Animate__000000_move_x1 := (&models.Animate{}).Stage(stage)
	__Animate__000001_move_x2 := (&models.Animate{}).Stage(stage)

	__Layer__000000_Animate_test := (&models.Layer{}).Stage(stage)

	__Line__000000_End := (&models.Line{}).Stage(stage)
	__Line__000001_Moving_Line := (&models.Line{}).Stage(stage)
	__Line__000002_Start_Line := (&models.Line{}).Stage(stage)

	__SVG__000000_Animate_test := (&models.SVG{}).Stage(stage)

	// Setup of values

	__Animate__000000_move_x1.Name = `move x1`
	__Animate__000000_move_x1.AttributeName = `x1`
	__Animate__000000_move_x1.Values = ``
	__Animate__000000_move_x1.From = `50`
	__Animate__000000_move_x1.To = `600`
	__Animate__000000_move_x1.Dur = `5s`
	__Animate__000000_move_x1.RepeatCount = `indefinite`

	__Animate__000001_move_x2.Name = `move x2`
	__Animate__000001_move_x2.AttributeName = `x2`
	__Animate__000001_move_x2.Values = `50;600`
	__Animate__000001_move_x2.From = ``
	__Animate__000001_move_x2.To = ``
	__Animate__000001_move_x2.Dur = `5s`
	__Animate__000001_move_x2.RepeatCount = `indefinite`

	__Layer__000000_Animate_test.Display = true
	__Layer__000000_Animate_test.Name = `Animate test`

	__Line__000000_End.Name = `End`
	__Line__000000_End.X1 = 600.000000
	__Line__000000_End.Y1 = 300.000000
	__Line__000000_End.X2 = 600.000000
	__Line__000000_End.Y2 = 30.000000
	__Line__000000_End.Color = ``
	__Line__000000_End.FillOpacity = 0.000000
	__Line__000000_End.Stroke = `Black`
	__Line__000000_End.StrokeOpacity = 0.500000
	__Line__000000_End.StrokeWidth = 3.000000
	__Line__000000_End.StrokeDashArray = ``
	__Line__000000_End.StrokeDashArrayWhenSelected = ``
	__Line__000000_End.Transform = ``
	__Line__000000_End.MouseClickX = 0.000000
	__Line__000000_End.MouseClickY = 0.000000

	__Line__000001_Moving_Line.Name = `Moving Line`
	__Line__000001_Moving_Line.X1 = 50.000000
	__Line__000001_Moving_Line.Y1 = 300.000000
	__Line__000001_Moving_Line.X2 = 50.000000
	__Line__000001_Moving_Line.Y2 = 500.000000
	__Line__000001_Moving_Line.Color = ``
	__Line__000001_Moving_Line.FillOpacity = 0.000000
	__Line__000001_Moving_Line.Stroke = `blue`
	__Line__000001_Moving_Line.StrokeOpacity = 0.500000
	__Line__000001_Moving_Line.StrokeWidth = 3.000000
	__Line__000001_Moving_Line.StrokeDashArray = ``
	__Line__000001_Moving_Line.StrokeDashArrayWhenSelected = ``
	__Line__000001_Moving_Line.Transform = ``
	__Line__000001_Moving_Line.MouseClickX = 0.000000
	__Line__000001_Moving_Line.MouseClickY = 0.000000

	__Line__000002_Start_Line.Name = `Start Line`
	__Line__000002_Start_Line.X1 = 50.000000
	__Line__000002_Start_Line.Y1 = 300.000000
	__Line__000002_Start_Line.X2 = 50.000000
	__Line__000002_Start_Line.Y2 = 30.000000
	__Line__000002_Start_Line.Color = ``
	__Line__000002_Start_Line.FillOpacity = 0.000000
	__Line__000002_Start_Line.Stroke = `Black`
	__Line__000002_Start_Line.StrokeOpacity = 0.500000
	__Line__000002_Start_Line.StrokeWidth = 3.000000
	__Line__000002_Start_Line.StrokeDashArray = ``
	__Line__000002_Start_Line.StrokeDashArrayWhenSelected = ``
	__Line__000002_Start_Line.Transform = ``
	__Line__000002_Start_Line.MouseClickX = 0.000000
	__Line__000002_Start_Line.MouseClickY = 0.000000

	__SVG__000000_Animate_test.Name = `Animate test`
	__SVG__000000_Animate_test.IsEditable = false

	// Setup of pointers
	__Layer__000000_Animate_test.Lines = append(__Layer__000000_Animate_test.Lines, __Line__000002_Start_Line)
	__Layer__000000_Animate_test.Lines = append(__Layer__000000_Animate_test.Lines, __Line__000000_End)
	__Layer__000000_Animate_test.Lines = append(__Layer__000000_Animate_test.Lines, __Line__000001_Moving_Line)
	__Line__000001_Moving_Line.Animates = append(__Line__000001_Moving_Line.Animates, __Animate__000000_move_x1)
	__Line__000001_Moving_Line.Animates = append(__Line__000001_Moving_Line.Animates, __Animate__000001_move_x2)
	__SVG__000000_Animate_test.Layers = append(__SVG__000000_Animate_test.Layers, __Layer__000000_Animate_test)
}
