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
}