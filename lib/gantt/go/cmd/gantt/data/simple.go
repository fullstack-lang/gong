package main

import (
	"time"

	"github.com/fullstack-lang/gong/lib/gantt/go/models"
	// injection point for ident package import declaration{{ImportPackageDeclaration}}
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var ___dummy__Stage_simple models.StageStruct
var ___dummy__Time_simple time.Time

// Injection point for meta package dummy declaration{{ImportPackageDummyDeclaration}}

// currently, DocLink renaming is not enabled in gopls
// the following map are devised to overcome this limitation
// those maps and the processing code will be eleminated when
// DocLink renaming will be enabled in gopls
// [Corresponding Issue](https://github.com/golang/go/issues/57559)
//
// When parsed, those maps will help with the renaming process
var map_DocLink_Identifier_simple map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// init might be handy if one want to have the data embedded in the binary
// but it has to properly reference the Injection gateway in the main package
// func init() {
// 	_ = __Dummy_time_variable
// 	InjectionGateway["simple"] = simpleInjection
// }

// simpleInjection will stage objects of database "simple"
func simpleInjection(stage *models.StageStruct) {

	// Declaration of instances to stage

	// Declarations of staged instances of Arrow

	// Declarations of staged instances of Bar
	__Bar__000000_One_Task := (&models.Bar{Name: `One Task`}).Stage(stage)

	// Declarations of staged instances of Gantt
	__Gantt__000000_Test := (&models.Gantt{Name: `Test`}).Stage(stage)

	// Declarations of staged instances of Group

	// Declarations of staged instances of Lane
	__Lane__000000_Lane_1 := (&models.Lane{Name: `Lane-1`}).Stage(stage)

	// Declarations of staged instances of LaneUse

	// Declarations of staged instances of Milestone

	// Setup of values

	// Bar values setup
	__Bar__000000_One_Task.Name = `One Task`
	__Bar__000000_One_Task.Start, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2022-04-17 03:08:18.46153846 +0000 +0000")
	__Bar__000000_One_Task.End, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2023-08-23 14:23:59.999999968 +0000 +0000")
	__Bar__000000_One_Task.OptionnalColor = `lightsteelblue`
	__Bar__000000_One_Task.OptionnalStroke = ``
	__Bar__000000_One_Task.FillOpacity = 100.000000
	__Bar__000000_One_Task.StrokeWidth = 0.000000
	__Bar__000000_One_Task.StrokeDashArray = ``

	// Gantt values setup
	__Gantt__000000_Test.Name = `Test`
	__Gantt__000000_Test.ComputedStart, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2022-01-01 00:00:00 +0000 +0000")
	__Gantt__000000_Test.ComputedEnd, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2024-12-31 00:00:00 +0000 +0000")
	__Gantt__000000_Test.UseManualStartAndEndDates = false
	__Gantt__000000_Test.ManualStart, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2017-02-12 00:00:00 +0000 +0000")
	__Gantt__000000_Test.ManualEnd, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2023-02-12 00:00:00 +0000 +0000")
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

	// Setup of pointers
	__Gantt__000000_Test.Lanes = append(__Gantt__000000_Test.Lanes, __Lane__000000_Lane_1)
	__Lane__000000_Lane_1.Bars = append(__Lane__000000_Lane_1.Bars, __Bar__000000_One_Task)
}
