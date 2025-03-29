package main

import (
	"time"

	"github.com/fullstack-lang/gong/lib/gantt/go/models"
	// injection point for ident package import declaration{{ImportPackageDeclaration}}
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var ___dummy__Stage_migration models.Stage
var ___dummy__Time_migration time.Time

// Injection point for meta package dummy declaration{{ImportPackageDummyDeclaration}}

// currently, DocLink renaming is not enabled in gopls
// the following map are devised to overcome this limitation
// those maps and the processing code will be eleminated when
// DocLink renaming will be enabled in gopls
// [Corresponding Issue](https://github.com/golang/go/issues/57559)
//
// When parsed, those maps will help with the renaming process
var map_DocLink_Identifier_migration map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// init might be handy if one want to have the data embedded in the binary
// but it has to properly reference the Injection gateway in the main package
// func init() {
// 	_ = __Dummy_time_variable
// 	InjectionGateway["migration"] = migrationInjection
// }

// migrationInjection will stage objects of database "migration"
func migrationInjection(stage *models.Stage) {

	// Declaration of instances to stage

	// Declarations of staged instances of Arrow
	__Arrow__000000_Lane_1_Bar_1_to_Lan_2_Bar_1 := (&models.Arrow{Name: `Lane 1 - Bar 1 to Lan 2 - Bar 1`}).Stage(stage)

	// Declarations of staged instances of Bar
	__Bar__000000_Bar_3 := (&models.Bar{Name: `Bar 3`}).Stage(stage)
	__Bar__000001_Conversion_MMP := (&models.Bar{Name: `Conversion MMP`}).Stage(stage)
	__Bar__000002_Conversion_mod_le := (&models.Bar{Name: `Conversion modèle`}).Stage(stage)

	// Declarations of staged instances of Gantt
	__Gantt__000000_Migration := (&models.Gantt{Name: `Migration`}).Stage(stage)

	// Declarations of staged instances of Group
	__Group__000000_Sprint_2 := (&models.Group{Name: `Sprint 2`}).Stage(stage)

	// Declarations of staged instances of Lane
	__Lane__000000_Lane_2 := (&models.Lane{Name: `Lane-2`}).Stage(stage)
	__Lane__000001_Lane_4 := (&models.Lane{Name: `Lane-4`}).Stage(stage)
	__Lane__000002_Sprint_2_1 := (&models.Lane{Name: `Sprint 2.1`}).Stage(stage)
	__Lane__000003_Sprint_2_2 := (&models.Lane{Name: `Sprint 2.2`}).Stage(stage)

	// Declarations of staged instances of LaneUse

	// Declarations of staged instances of Milestone

	// Setup of values

	// Arrow values setup
	__Arrow__000000_Lane_1_Bar_1_to_Lan_2_Bar_1.Name = `Lane 1 - Bar 1 to Lan 2 - Bar 1`
	__Arrow__000000_Lane_1_Bar_1_to_Lan_2_Bar_1.OptionnalColor = ``
	__Arrow__000000_Lane_1_Bar_1_to_Lan_2_Bar_1.OptionnalStroke = ``

	// Bar values setup
	__Bar__000000_Bar_3.Name = `Bar 3`
	__Bar__000000_Bar_3.Start, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2021-05-27 08:33:58.15384616 +0000 UTC")
	__Bar__000000_Bar_3.End, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2021-11-27 02:39:30.461538464 +0000 UTC")
	__Bar__000000_Bar_3.OptionnalColor = ``
	__Bar__000000_Bar_3.OptionnalStroke = ``
	__Bar__000000_Bar_3.FillOpacity = 0.000000
	__Bar__000000_Bar_3.StrokeWidth = 0.000000
	__Bar__000000_Bar_3.StrokeDashArray = ``

	// Bar values setup
	__Bar__000001_Conversion_MMP.Name = `Conversion MMP`
	__Bar__000001_Conversion_MMP.Start, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2021-03-20 22:35:48.92307692 +0000 UTC")
	__Bar__000001_Conversion_MMP.End, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2021-08-16 06:07:45.230769216 +0000 UTC")
	__Bar__000001_Conversion_MMP.OptionnalColor = ``
	__Bar__000001_Conversion_MMP.OptionnalStroke = ``
	__Bar__000001_Conversion_MMP.FillOpacity = 0.000000
	__Bar__000001_Conversion_MMP.StrokeWidth = 0.000000
	__Bar__000001_Conversion_MMP.StrokeDashArray = ``

	// Bar values setup
	__Bar__000002_Conversion_mod_le.Name = `Conversion modèle`
	__Bar__000002_Conversion_mod_le.Start, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2020-10-10 03:12:44.307692308 +0000 UTC")
	__Bar__000002_Conversion_mod_le.End, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2021-04-14 23:35:37.846153832 +0000 UTC")
	__Bar__000002_Conversion_mod_le.OptionnalColor = ``
	__Bar__000002_Conversion_mod_le.OptionnalStroke = ``
	__Bar__000002_Conversion_mod_le.FillOpacity = 0.000000
	__Bar__000002_Conversion_mod_le.StrokeWidth = 1.000000
	__Bar__000002_Conversion_mod_le.StrokeDashArray = ``

	// Gantt values setup
	__Gantt__000000_Migration.Name = `Migration`
	__Gantt__000000_Migration.ComputedStart, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2020-01-01 00:00:00 +0000 UTC")
	__Gantt__000000_Migration.ComputedEnd, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2022-12-31 00:00:00 +0000 UTC")
	__Gantt__000000_Migration.UseManualStartAndEndDates = false
	__Gantt__000000_Migration.ManualStart, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2017-02-12 00:00:00 +0000 UTC")
	__Gantt__000000_Migration.ManualEnd, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2023-02-12 00:00:00 +0000 UTC")
	__Gantt__000000_Migration.LaneHeight = 80.000000
	__Gantt__000000_Migration.RatioBarToLaneHeight = 0.700000
	__Gantt__000000_Migration.YTopMargin = 10.000000
	__Gantt__000000_Migration.XLeftText = 8.000000
	__Gantt__000000_Migration.TextHeight = 16.000000
	__Gantt__000000_Migration.XLeftLanes = 150.000000
	__Gantt__000000_Migration.XRightMargin = 800.000000
	__Gantt__000000_Migration.ArrowLengthToTheRightOfStartBar = 50.000000
	__Gantt__000000_Migration.ArrowTipLenght = 15.000000
	__Gantt__000000_Migration.TimeLine_Color = `"black"`
	__Gantt__000000_Migration.TimeLine_FillOpacity = 1.000000
	__Gantt__000000_Migration.TimeLine_Stroke = `"black"`
	__Gantt__000000_Migration.TimeLine_StrokeWidth = 1.000000
	__Gantt__000000_Migration.Group_Stroke = `blue`
	__Gantt__000000_Migration.Group_StrokeWidth = 0.300000
	__Gantt__000000_Migration.Group_StrokeDashArray = ``
	__Gantt__000000_Migration.DateYOffset = 20.000000
	__Gantt__000000_Migration.AlignOnStartEndOnYearStart = true

	// Group values setup
	__Group__000000_Sprint_2.Name = `Sprint 2`

	// Lane values setup
	__Lane__000000_Lane_2.Name = `Lane-2`
	__Lane__000000_Lane_2.Order = 20

	// Lane values setup
	__Lane__000001_Lane_4.Name = `Lane-4`
	__Lane__000001_Lane_4.Order = 40

	// Lane values setup
	__Lane__000002_Sprint_2_1.Name = `Sprint 2.1`
	__Lane__000002_Sprint_2_1.Order = 10

	// Lane values setup
	__Lane__000003_Sprint_2_2.Name = `Sprint 2.2`
	__Lane__000003_Sprint_2_2.Order = 15

	// Setup of pointers
	__Arrow__000000_Lane_1_Bar_1_to_Lan_2_Bar_1.From = __Bar__000001_Conversion_MMP
	__Arrow__000000_Lane_1_Bar_1_to_Lan_2_Bar_1.To = __Bar__000000_Bar_3
	__Gantt__000000_Migration.Lanes = append(__Gantt__000000_Migration.Lanes, __Lane__000002_Sprint_2_1)
	__Gantt__000000_Migration.Lanes = append(__Gantt__000000_Migration.Lanes, __Lane__000003_Sprint_2_2)
	__Gantt__000000_Migration.Lanes = append(__Gantt__000000_Migration.Lanes, __Lane__000000_Lane_2)
	__Gantt__000000_Migration.Lanes = append(__Gantt__000000_Migration.Lanes, __Lane__000001_Lane_4)
	__Gantt__000000_Migration.Groups = append(__Gantt__000000_Migration.Groups, __Group__000000_Sprint_2)
	__Gantt__000000_Migration.Arrows = append(__Gantt__000000_Migration.Arrows, __Arrow__000000_Lane_1_Bar_1_to_Lan_2_Bar_1)
	__Group__000000_Sprint_2.GroupLanes = append(__Group__000000_Sprint_2.GroupLanes, __Lane__000002_Sprint_2_1)
	__Group__000000_Sprint_2.GroupLanes = append(__Group__000000_Sprint_2.GroupLanes, __Lane__000003_Sprint_2_2)
	__Lane__000000_Lane_2.Bars = append(__Lane__000000_Lane_2.Bars, __Bar__000000_Bar_3)
	__Lane__000002_Sprint_2_1.Bars = append(__Lane__000002_Sprint_2_1.Bars, __Bar__000002_Conversion_mod_le)
	__Lane__000003_Sprint_2_2.Bars = append(__Lane__000003_Sprint_2_2.Bars, __Bar__000001_Conversion_MMP)
}
