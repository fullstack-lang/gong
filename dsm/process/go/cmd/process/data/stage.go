package main

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/dsm/process/go/models"
	// injection point for ident package import declaration{{ImportPackageDeclaration}}
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var _ time.Time
var _ = slices.Index[[]int, int]

// _ point for meta package dummy declaration{{ImportPackageDummyDeclaration}}

// When parsed, those maps will help with the renaming process
var _ map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// function will stage objects
func _(stage *models.Stage) {

	__DiagramProcess__00000000_ := (&models.DiagramProcess{Name: `D1`}).Stage(stage)
	__DiagramProcess__00000001_ := (&models.DiagramProcess{Name: `D2`}).Stage(stage)
	__Library__00000000_ := (&models.Library{Name: `Root`}).Stage(stage)
	__ProcessShape__00000000_ := (&models.ProcessShape{Name: `-D1`}).Stage(stage)
	__ProcessShape__00000001_ := (&models.ProcessShape{Name: `-D1`}).Stage(stage)
	__Process__00000001_ := (&models.Process{Name: `P1`}).Stage(stage)
	__DiagramProcess__00000000_.Name = `D1`
	__DiagramProcess__00000000_.ComputedPrefix = ``
	__DiagramProcess__00000000_.IsChecked = false
	__DiagramProcess__00000000_.IsEditable_ = true
	__DiagramProcess__00000000_.IsShowPrefix = false
	__DiagramProcess__00000000_.DefaultBoxWidth = 0.000000
	__DiagramProcess__00000000_.DefaultBoxHeigth = 0.000000
	__DiagramProcess__00000000_.Width = 0.000000
	__DiagramProcess__00000000_.Height = 0.000000
	__DiagramProcess__00000000_.IsProcesssNodeExpanded = true
	__DiagramProcess__00000000_.Process_Shapes = append(__DiagramProcess__00000000_.Process_Shapes, __ProcessShape__00000000_)
	__DiagramProcess__00000000_.Process_Shapes = append(__DiagramProcess__00000000_.Process_Shapes, __ProcessShape__00000001_)
	__DiagramProcess__00000001_.Name = `D2`
	__DiagramProcess__00000001_.ComputedPrefix = ``
	__DiagramProcess__00000001_.IsChecked = true
	__DiagramProcess__00000001_.IsEditable_ = true
	__DiagramProcess__00000001_.IsShowPrefix = false
	__DiagramProcess__00000001_.DefaultBoxWidth = 0.000000
	__DiagramProcess__00000001_.DefaultBoxHeigth = 0.000000
	__DiagramProcess__00000001_.Width = 0.000000
	__DiagramProcess__00000001_.Height = 0.000000
	__DiagramProcess__00000001_.IsProcesssNodeExpanded = true
	__Library__00000000_.Name = `Root`
	__Library__00000000_.ComputedPrefix = ``
	__Library__00000000_.NbPixPerCharacter = 0.000000
	__Library__00000000_.LogoSVGFile = ``
	__Library__00000000_.DiagramProcesss = append(__Library__00000000_.DiagramProcesss, __DiagramProcess__00000000_)
	__Library__00000000_.DiagramProcesss = append(__Library__00000000_.DiagramProcesss, __DiagramProcess__00000001_)
	__Library__00000000_.RootProcesses = append(__Library__00000000_.RootProcesses, __Process__00000001_)
	__ProcessShape__00000000_.Name = `-D1`
	__ProcessShape__00000000_.IsExpanded = false
	__ProcessShape__00000000_.X = 128.340435
	__ProcessShape__00000000_.Y = 169.677953
	__ProcessShape__00000000_.Width = 0.000000
	__ProcessShape__00000000_.Height = 0.000000
	__ProcessShape__00000000_.IsHidden = false
	__ProcessShape__00000000_.Process = __Process__00000001_
	__ProcessShape__00000001_.Name = `-D1`
	__ProcessShape__00000001_.IsExpanded = false
	__ProcessShape__00000001_.X = 103.847267
	__ProcessShape__00000001_.Y = 191.884898
	__ProcessShape__00000001_.Width = 0.000000
	__ProcessShape__00000001_.Height = 0.000000
	__ProcessShape__00000001_.IsHidden = false
	__ProcessShape__00000001_.Process = nil
	__Process__00000001_.Name = `P1`
	__Process__00000001_.ComputedPrefix = ``
	stage.Commit()

	// D1
	__DiagramProcess__00000000_.IsChecked = true
	// D2
	__DiagramProcess__00000001_.IsChecked = false
	stage.Commit()

	// D1
	__DiagramProcess__00000000_.Process_Shapes = slices.Delete( __DiagramProcess__00000000_.Process_Shapes, 1, 2)
	__ProcessShape__00000001_.Unstage(stage)
	stage.Commit()

	__ProcessShape__00000002_ := (&models.ProcessShape{Name: `P1-D1`}).Stage(stage)
	// D1
	__DiagramProcess__00000000_.Process_Shapes = slices.Insert( __DiagramProcess__00000000_.Process_Shapes, 1, __ProcessShape__00000002_)
	__ProcessShape__00000002_.Name = `P1-D1`
	__ProcessShape__00000002_.IsExpanded = false
	__ProcessShape__00000002_.X = 102.568327
	__ProcessShape__00000002_.Y = 174.777678
	__ProcessShape__00000002_.Width = 0.000000
	__ProcessShape__00000002_.Height = 0.000000
	__ProcessShape__00000002_.IsHidden = false
	__ProcessShape__00000002_.Process = __Process__00000001_
	stage.Commit()

	// D1
	__DiagramProcess__00000000_.Process_Shapes = slices.Delete( __DiagramProcess__00000000_.Process_Shapes, 1, 2)
	__ProcessShape__00000002_.Unstage(stage)
	stage.Commit()

	// D1
	__DiagramProcess__00000000_.Process_Shapes = slices.Delete( __DiagramProcess__00000000_.Process_Shapes, 0, 1)
	__ProcessShape__00000000_.Unstage(stage)
	stage.Commit()

	// D1
	__DiagramProcess__00000000_.IsChecked = false
	// D2
	__DiagramProcess__00000001_.IsChecked = true
	stage.Commit()

	__ProcessShape__00000003_ := (&models.ProcessShape{Name: `P1-D2`}).Stage(stage)
	// D2
	__DiagramProcess__00000001_.Process_Shapes = slices.Insert( __DiagramProcess__00000001_.Process_Shapes, 0, __ProcessShape__00000003_)
	__ProcessShape__00000003_.Name = `P1-D2`
	__ProcessShape__00000003_.IsExpanded = false
	__ProcessShape__00000003_.X = 163.901066
	__ProcessShape__00000003_.Y = 129.483041
	__ProcessShape__00000003_.Width = 0.000000
	__ProcessShape__00000003_.Height = 0.000000
	__ProcessShape__00000003_.IsHidden = false
	__ProcessShape__00000003_.Process = __Process__00000001_
	stage.Commit()

	__ProcessShape__00000004_ := (&models.ProcessShape{Name: `P1-D2`}).Stage(stage)
	// D2
	__DiagramProcess__00000001_.Process_Shapes = slices.Insert( __DiagramProcess__00000001_.Process_Shapes, 1, __ProcessShape__00000004_)
	__ProcessShape__00000004_.Name = `P1-D2`
	__ProcessShape__00000004_.IsExpanded = false
	__ProcessShape__00000004_.X = 187.623755
	__ProcessShape__00000004_.Y = 114.162730
	__ProcessShape__00000004_.Width = 0.000000
	__ProcessShape__00000004_.Height = 0.000000
	__ProcessShape__00000004_.IsHidden = false
	__ProcessShape__00000004_.Process = __Process__00000001_
	stage.Commit()

	__ProcessShape__00000005_ := (&models.ProcessShape{Name: `P1-D2`}).Stage(stage)
	// D2
	__DiagramProcess__00000001_.Process_Shapes = slices.Insert( __DiagramProcess__00000001_.Process_Shapes, 2, __ProcessShape__00000005_)
	__ProcessShape__00000005_.Name = `P1-D2`
	__ProcessShape__00000005_.IsExpanded = false
	__ProcessShape__00000005_.X = 160.032097
	__ProcessShape__00000005_.Y = 171.426270
	__ProcessShape__00000005_.Width = 0.000000
	__ProcessShape__00000005_.Height = 0.000000
	__ProcessShape__00000005_.IsHidden = false
	__ProcessShape__00000005_.Process = __Process__00000001_
	stage.Commit()

	// D1
	__DiagramProcess__00000000_.IsChecked = true
	// D2
	__DiagramProcess__00000001_.IsChecked = false
	stage.Commit()

	// D2
	__DiagramProcess__00000001_.Process_Shapes = slices.Delete( __DiagramProcess__00000001_.Process_Shapes, 2, 3)
	__ProcessShape__00000005_.Unstage(stage)
	stage.Commit()

	// D2
	__DiagramProcess__00000001_.Process_Shapes = slices.Delete( __DiagramProcess__00000001_.Process_Shapes, 1, 2)
	__ProcessShape__00000004_.Unstage(stage)
	stage.Commit()

	// D1
	__DiagramProcess__00000000_.IsChecked = false
	// D2
	__DiagramProcess__00000001_.IsChecked = true
	stage.Commit()

	// D1
	__DiagramProcess__00000000_.IsChecked = true
	// D2
	__DiagramProcess__00000001_.IsChecked = false
	stage.Commit()

	// D1
	__DiagramProcess__00000000_.IsChecked = false
	// D2
	__DiagramProcess__00000001_.IsChecked = true
	stage.Commit()

	// P1-D2
	__ProcessShape__00000003_.Width = 100.000000
	__ProcessShape__00000003_.Height = 100.000000
	stage.Commit()

	// D1
	__DiagramProcess__00000000_.IsChecked = true
	// D2
	__DiagramProcess__00000001_.IsChecked = false
	stage.Commit()

	// D1
	__DiagramProcess__00000000_.IsChecked = false
	// D2
	__DiagramProcess__00000001_.IsChecked = true
	stage.Commit()

	// D1
	__DiagramProcess__00000000_.DefaultBoxWidth = 250.000000
	__DiagramProcess__00000000_.DefaultBoxHeigth = 70.000000
	// D2
	__DiagramProcess__00000001_.DefaultBoxWidth = 250.000000
	__DiagramProcess__00000001_.DefaultBoxHeigth = 70.000000
	// Root
	__Library__00000000_.NbPixPerCharacter = 8.000000
	stage.Commit()

	// D1
	__DiagramProcess__00000000_.IsChecked = true
	// D2
	__DiagramProcess__00000001_.IsChecked = false
	stage.Commit()

	// D1
	__DiagramProcess__00000000_.IsChecked = false
	// D2
	__DiagramProcess__00000001_.IsChecked = true
	stage.Commit()

	// D1
	__DiagramProcess__00000000_.Width = 200.000000
	__DiagramProcess__00000000_.Height = 200.000000
	// D2
	__DiagramProcess__00000001_.Width = 463.901066
	__DiagramProcess__00000001_.Height = 429.483041
	stage.Commit()

	// P1-D2
	__ProcessShape__00000003_.Width = 278.000000
	stage.Commit()

	// P1-D2
	__ProcessShape__00000003_.Width = 309.000000
	stage.Commit()

	// P1-D2
	__ProcessShape__00000003_.Width = 430.000000
	stage.Commit()

	// P1-D2
	__ProcessShape__00000003_.Width = 521.000000
	stage.Commit()

	// D1
	__DiagramProcess__00000000_.Width = 800.000000
	__DiagramProcess__00000000_.Height = 800.000000
	// D2
	__DiagramProcess__00000001_.Width = 1284.901066
	__DiagramProcess__00000001_.Height = 1029.483041
	stage.Commit()

	// Root
	__Library__00000000_.DiagramProcesss = slices.Delete( __Library__00000000_.DiagramProcesss, 0, 1)
	__DiagramProcess__00000000_.Unstage(stage)
	stage.Commit()

	// D2
	__DiagramProcess__00000001_.Width = 1584.901066
	__DiagramProcess__00000001_.Height = 1329.483041
	stage.Commit()

	// D2
	__DiagramProcess__00000001_.Process_Shapes = slices.Delete( __DiagramProcess__00000001_.Process_Shapes, 0, 1)
	__ProcessShape__00000003_.Unstage(stage)
	stage.Commit()

	__ProcessShape__00000006_ := (&models.ProcessShape{Name: `P1-D2`}).Stage(stage)
	// D2
	__DiagramProcess__00000001_.Process_Shapes = slices.Insert( __DiagramProcess__00000001_.Process_Shapes, 0, __ProcessShape__00000006_)
	__ProcessShape__00000006_.Name = `P1-D2`
	__ProcessShape__00000006_.IsExpanded = false
	__ProcessShape__00000006_.X = 167.190905
	__ProcessShape__00000006_.Y = 118.755068
	__ProcessShape__00000006_.Width = 250.000000
	__ProcessShape__00000006_.Height = 70.000000
	__ProcessShape__00000006_.IsHidden = false
	__ProcessShape__00000006_.Process = __Process__00000001_
	stage.Commit()

	// D2
	__DiagramProcess__00000001_.Process_Shapes = slices.Delete( __DiagramProcess__00000001_.Process_Shapes, 0, 1)
	__ProcessShape__00000006_.Unstage(stage)
	stage.Commit()

	__ProcessShape__00000007_ := (&models.ProcessShape{Name: `P1-D2`}).Stage(stage)
	// D2
	__DiagramProcess__00000001_.Process_Shapes = slices.Insert( __DiagramProcess__00000001_.Process_Shapes, 0, __ProcessShape__00000007_)
	__ProcessShape__00000007_.Name = `P1-D2`
	__ProcessShape__00000007_.IsExpanded = false
	__ProcessShape__00000007_.X = 121.522900
	__ProcessShape__00000007_.Y = 160.796912
	__ProcessShape__00000007_.Width = 250.000000
	__ProcessShape__00000007_.Height = 70.000000
	__ProcessShape__00000007_.IsHidden = false
	__ProcessShape__00000007_.Process = __Process__00000001_
	stage.Commit()
}