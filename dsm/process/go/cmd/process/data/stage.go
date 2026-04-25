package main

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/dsm/process/go/models"
	// injection point for ident package import declaration{{ImportPackageDeclaration}}
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var (
	_ time.Time
	_ = slices.Index[[]int, int]
)

// _ point for meta package dummy declaration{{ImportPackageDummyDeclaration}}

// When parsed, those maps will help with the renaming process
var _ map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// function will stage objects
func _(stage *models.Stage) {

	// insertion point for declaration of instances to stage

	__DiagramProcess__00000000_ := (&models.DiagramProcess{Name: `D1`}).Stage(stage)
	__DiagramProcess__00000005_ := (&models.DiagramProcess{Name: `DiagramProcess`}).Stage(stage)

	__Library__00000000_ := (&models.Library{Name: `Root`}).Stage(stage)

	__Process__00000000_ := (&models.Process{Name: `P0`}).Stage(stage)
	__Process__00000005_ := (&models.Process{Name: `P1`}).Stage(stage)

	__ProcessShape__00000003_ := (&models.ProcessShape{Name: `ProcessShape`}).Stage(stage)
	__ProcessShape__00000005_ := (&models.ProcessShape{Name: `ProcessShape`}).Stage(stage)

	// insertion point for initialization of values

	__DiagramProcess__00000000_.Name = `D1`
	__DiagramProcess__00000000_.ComputedPrefix = ``
	__DiagramProcess__00000000_.IsChecked = false
	__DiagramProcess__00000000_.IsEditable_ = true
	__DiagramProcess__00000000_.IsShowPrefix = false
	__DiagramProcess__00000000_.DefaultBoxWidth = 250.000000
	__DiagramProcess__00000000_.DefaultBoxHeigth = 70.000000
	__DiagramProcess__00000000_.Width = 33900.000000
	__DiagramProcess__00000000_.Height = 33900.000000
	__DiagramProcess__00000000_.IsProcesssNodeExpanded = false

	__DiagramProcess__00000005_.Name = `DiagramProcess`
	__DiagramProcess__00000005_.ComputedPrefix = ``
	__DiagramProcess__00000005_.IsChecked = true
	__DiagramProcess__00000005_.IsEditable_ = false
	__DiagramProcess__00000005_.IsShowPrefix = false
	__DiagramProcess__00000005_.DefaultBoxWidth = 250.000000
	__DiagramProcess__00000005_.DefaultBoxHeigth = 70.000000
	__DiagramProcess__00000005_.Width = 3300.000000
	__DiagramProcess__00000005_.Height = 3750.000000
	__DiagramProcess__00000005_.IsProcesssNodeExpanded = false

	__Library__00000000_.Name = `Root`
	__Library__00000000_.ComputedPrefix = ``
	__Library__00000000_.NbPixPerCharacter = 8.000000
	__Library__00000000_.LogoSVGFile = ``

	__Process__00000000_.Name = `P0`
	__Process__00000000_.ComputedPrefix = ``
	__Process__00000000_.IsSubProcessNodeExpanded = true

	__Process__00000005_.Name = `P1`
	__Process__00000005_.ComputedPrefix = ``
	__Process__00000005_.IsSubProcessNodeExpanded = false

	__ProcessShape__00000003_.Name = `ProcessShape`
	__ProcessShape__00000003_.IsExpanded = false
	__ProcessShape__00000003_.X = 100.000000
	__ProcessShape__00000003_.Y = 50.000000
	__ProcessShape__00000003_.Width = 887.000000
	__ProcessShape__00000003_.Height = 1000.000000
	__ProcessShape__00000003_.IsHidden = false

	__ProcessShape__00000005_.Name = `ProcessShape`
	__ProcessShape__00000005_.IsExpanded = false
	__ProcessShape__00000005_.X = 100.000000
	__ProcessShape__00000005_.Y = 50.000000
	__ProcessShape__00000005_.Width = 500.000000
	__ProcessShape__00000005_.Height = 1000.000000
	__ProcessShape__00000005_.IsHidden = false

	// insertion point for setup of pointers
	__DiagramProcess__00000000_.Process_Shapes = append(__DiagramProcess__00000000_.Process_Shapes, __ProcessShape__00000003_)
	__DiagramProcess__00000005_.Process_Shapes = append(__DiagramProcess__00000005_.Process_Shapes, __ProcessShape__00000005_)
	__Library__00000000_.RootProcesses = append(__Library__00000000_.RootProcesses, __Process__00000000_)
	__Library__00000000_.RootProcesses = append(__Library__00000000_.RootProcesses, __Process__00000005_)
	__Library__00000000_.ProcesssWhoseNodeIsExpanded = append(__Library__00000000_.ProcesssWhoseNodeIsExpanded, __Process__00000005_)
	__Process__00000000_.DiagramProcesss = append(__Process__00000000_.DiagramProcesss, __DiagramProcess__00000000_)
	__Process__00000005_.DiagramProcesss = append(__Process__00000005_.DiagramProcesss, __DiagramProcess__00000005_)
	__ProcessShape__00000003_.Process = __Process__00000000_
	__ProcessShape__00000005_.Process = __Process__00000005_
}
