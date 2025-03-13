package main

import (
	"time"

	"github.com/fullstack-lang/gong/lib/gantt/go/models"
	// injection point for ident package import declaration{{ImportPackageDeclaration}}
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var ___dummy__Stage_gantt1 models.StageStruct
var ___dummy__Time_gantt1 time.Time

// _ point for meta package dummy declaration{{ImportPackageDummyDeclaration}}

// currently, DocLink renaming is not enabled in gopls
// the following map are devised to overcome this limitation
// those maps and the processing code will be eleminated when
// DocLink renaming will be enabled in gopls
// [Corresponding Issue](https://github.com/golang/go/issues/57559)
//
// When parsed, those maps will help with the renaming process
var _ map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// init might be handy if one want to have the data embedded in the binary
// but it has to properly reference the _ gateway in the main package
// func init() {
// 	_ = __Dummy_time_variable
// 	InjectionGateway["gantt1"] = _
// }

// _ will stage objects of database "gantt1"
func _(stage *models.StageStruct) {

	// Declaration of instances to stage

	// Declarations of staged instances of Arrow
	__Arrow__000000_Lane_1_Bar_1_to_Lan_2_Bar_1 := (&models.Arrow{Name: `Lane 1 - Bar 1 to Lan 2 - Bar 1`}).Stage(stage)

	// Declarations of staged instances of Bar
	__Bar__000000_Bar_1 := (&models.Bar{Name: `Bar 1`}).Stage(stage)
	__Bar__000001_Bar_2 := (&models.Bar{Name: `Bar 2`}).Stage(stage)
	__Bar__000002_Bar_3 := (&models.Bar{Name: `Bar 3`}).Stage(stage)
	__Bar__000003_Bar_4 := (&models.Bar{Name: `Bar-4`}).Stage(stage)

	// Declarations of staged instances of Gantt
	__Gantt__000000_Test := (&models.Gantt{Name: `Test`}).Stage(stage)

	// Declarations of staged instances of Group

	// Declarations of staged instances of Lane
	__Lane__000000_Lane_1 := (&models.Lane{Name: `Lane-1`}).Stage(stage)
	__Lane__000001_Lane_2 := (&models.Lane{Name: `Lane-2`}).Stage(stage)
	__Lane__000002_Lane_3 := (&models.Lane{Name: `Lane-3`}).Stage(stage)
	__Lane__000003_Lane_4 := (&models.Lane{Name: `Lane-4`}).Stage(stage)

	// Declarations of staged instances of LaneUse

	// Declarations of staged instances of Milestone

	// Setup of values

	// Arrow values setup
	__Arrow__000000_Lane_1_Bar_1_to_Lan_2_Bar_1.Name = `Lane 1 - Bar 1 to Lan 2 - Bar 1`
	__Arrow__000000_Lane_1_Bar_1_to_Lan_2_Bar_1.OptionnalColor = ``
	__Arrow__000000_Lane_1_Bar_1_to_Lan_2_Bar_1.OptionnalStroke = ``

	// Bar values setup
	__Bar__000000_Bar_1.Name = `Bar 1`
	__Bar__000000_Bar_1.Start, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2020-04-01 00:00:00.000000001 +0000 UTC")
	__Bar__000000_Bar_1.End, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2024-03-31 23:59:59.999999984 +0000 UTC")
	__Bar__000000_Bar_1.ComputedDuration = 126230399999999983
	__Bar__000000_Bar_1.OptionnalColor = `red`
	__Bar__000000_Bar_1.OptionnalStroke = `red`
	__Bar__000000_Bar_1.FillOpacity = 0.000000
	__Bar__000000_Bar_1.StrokeWidth = 3.000000
	__Bar__000000_Bar_1.StrokeDashArray = `5 5`

	// Bar values setup
	__Bar__000001_Bar_2.Name = `Bar 2`
	__Bar__000001_Bar_2.Start, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2023-05-08 14:08:29.538461536 +0000 UTC")
	__Bar__000001_Bar_2.End, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2027-07-04 16:23:37.846153888 +0000 UTC")
	__Bar__000001_Bar_2.ComputedDuration = 131163308307692352
	__Bar__000001_Bar_2.OptionnalColor = `green`
	__Bar__000001_Bar_2.OptionnalStroke = ``
	__Bar__000001_Bar_2.FillOpacity = 0.000000
	__Bar__000001_Bar_2.StrokeWidth = 0.000000
	__Bar__000001_Bar_2.StrokeDashArray = ``

	// Bar values setup
	__Bar__000002_Bar_3.Name = `Bar 3`
	__Bar__000002_Bar_3.Start, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2021-05-27 08:33:58.153846152 +0000 UTC")
	__Bar__000002_Bar_3.End, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2023-08-30 21:55:56.3076924 +0000 UTC")
	__Bar__000002_Bar_3.ComputedDuration = 71328118153846248
	__Bar__000002_Bar_3.OptionnalColor = ``
	__Bar__000002_Bar_3.OptionnalStroke = ``
	__Bar__000002_Bar_3.FillOpacity = 0.000000
	__Bar__000002_Bar_3.StrokeWidth = 0.000000
	__Bar__000002_Bar_3.StrokeDashArray = ``

	// Bar values setup
	__Bar__000003_Bar_4.Name = `Bar-4`
	__Bar__000003_Bar_4.Start, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2023-10-06 01:55:12 +0000 UTC")
	__Bar__000003_Bar_4.End, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2025-08-04 15:23:48.923076864 +0000 UTC")
	__Bar__000003_Bar_4.ComputedDuration = 110956475076923016
	__Bar__000003_Bar_4.OptionnalColor = ``
	__Bar__000003_Bar_4.OptionnalStroke = ``
	__Bar__000003_Bar_4.FillOpacity = 0.000000
	__Bar__000003_Bar_4.StrokeWidth = 0.000000
	__Bar__000003_Bar_4.StrokeDashArray = ``

	// Gantt values setup
	__Gantt__000000_Test.Name = `Test`
	__Gantt__000000_Test.ComputedStart, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2020-01-01 00:00:00 +0000 UTC")
	__Gantt__000000_Test.ComputedEnd, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2027-12-31 00:00:00 +0000 UTC")
	__Gantt__000000_Test.ComputedDuration = 229019017846153887
	__Gantt__000000_Test.UseManualStartAndEndDates = false
	__Gantt__000000_Test.ManualStart, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2017-02-12 00:00:00 +0000 UTC")
	__Gantt__000000_Test.ManualEnd, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2023-02-12 00:00:00 +0000 UTC")
	__Gantt__000000_Test.LaneHeight = 80.000000
	__Gantt__000000_Test.RatioBarToLaneHeight = 0.700000
	__Gantt__000000_Test.YTopMargin = 10.000000
	__Gantt__000000_Test.XLeftText = 8.000000
	__Gantt__000000_Test.TextHeight = 16.000000
	__Gantt__000000_Test.XLeftLanes = 150.000000
	__Gantt__000000_Test.XRightMargin = 800.000000
	__Gantt__000000_Test.ArrowLengthToTheRightOfStartBar = 50.000000
	__Gantt__000000_Test.ArrowTipLenght = 15.000000
	__Gantt__000000_Test.TimeLine_Color = `"black"`
	__Gantt__000000_Test.TimeLine_FillOpacity = 1.000000
	__Gantt__000000_Test.TimeLine_Stroke = `"black"`
	__Gantt__000000_Test.TimeLine_StrokeWidth = 1.000000
	__Gantt__000000_Test.Group_Stroke = `blue`
	__Gantt__000000_Test.Group_StrokeWidth = 0.300000
	__Gantt__000000_Test.Group_StrokeDashArray = ``
	__Gantt__000000_Test.DateYOffset = 20.000000
	__Gantt__000000_Test.AlignOnStartEndOnYearStart = true

	// Lane values setup
	__Lane__000000_Lane_1.Name = `Lane-1`
	__Lane__000000_Lane_1.Order = 10

	// Lane values setup
	__Lane__000001_Lane_2.Name = `Lane-2`
	__Lane__000001_Lane_2.Order = 20

	// Lane values setup
	__Lane__000002_Lane_3.Name = `Lane-3`
	__Lane__000002_Lane_3.Order = 30

	// Lane values setup
	__Lane__000003_Lane_4.Name = `Lane-4`
	__Lane__000003_Lane_4.Order = 40

	// Setup of pointers
	__Arrow__000000_Lane_1_Bar_1_to_Lan_2_Bar_1.From = __Bar__000003_Bar_4
	__Arrow__000000_Lane_1_Bar_1_to_Lan_2_Bar_1.To = __Bar__000002_Bar_3
	__Gantt__000000_Test.Lanes = append(__Gantt__000000_Test.Lanes, __Lane__000000_Lane_1)
	__Gantt__000000_Test.Lanes = append(__Gantt__000000_Test.Lanes, __Lane__000001_Lane_2)
	__Gantt__000000_Test.Lanes = append(__Gantt__000000_Test.Lanes, __Lane__000002_Lane_3)
	__Gantt__000000_Test.Lanes = append(__Gantt__000000_Test.Lanes, __Lane__000003_Lane_4)
	__Gantt__000000_Test.Arrows = append(__Gantt__000000_Test.Arrows, __Arrow__000000_Lane_1_Bar_1_to_Lan_2_Bar_1)
	__Lane__000000_Lane_1.Bars = append(__Lane__000000_Lane_1.Bars, __Bar__000000_Bar_1)
	__Lane__000000_Lane_1.Bars = append(__Lane__000000_Lane_1.Bars, __Bar__000001_Bar_2)
	__Lane__000001_Lane_2.Bars = append(__Lane__000001_Lane_2.Bars, __Bar__000002_Bar_3)
	__Lane__000002_Lane_3.Bars = append(__Lane__000002_Lane_3.Bars, __Bar__000003_Bar_4)
}
