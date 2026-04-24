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

	__Library__00000000_ := (&models.Library{Name: `Root`}).Stage(stage)

	__Process__00000000_ := (&models.Process{Name: `P0`}).Stage(stage)
	__Process__00000001_ := (&models.Process{Name: `P1 (P0.A)`}).Stage(stage)
	__Process__00000002_ := (&models.Process{Name: `P2 (P0.B)`}).Stage(stage)
	__Process__00000003_ := (&models.Process{Name: `P3`}).Stage(stage)

	// insertion point for initialization of values

	__DiagramProcess__00000000_.Name = `D1`
	__DiagramProcess__00000000_.ComputedPrefix = ``
	__DiagramProcess__00000000_.IsChecked = true
	__DiagramProcess__00000000_.IsEditable_ = true
	__DiagramProcess__00000000_.IsShowPrefix = false
	__DiagramProcess__00000000_.DefaultBoxWidth = 250.000000
	__DiagramProcess__00000000_.DefaultBoxHeigth = 70.000000
	__DiagramProcess__00000000_.Width = 2400.000000
	__DiagramProcess__00000000_.Height = 2400.000000
	__DiagramProcess__00000000_.IsProcesssNodeExpanded = false

	__Library__00000000_.Name = `Root`
	__Library__00000000_.ComputedPrefix = ``
	__Library__00000000_.NbPixPerCharacter = 8.000000
	__Library__00000000_.LogoSVGFile = ``

	__Process__00000000_.Name = `P0`
	__Process__00000000_.ComputedPrefix = ``
	__Process__00000000_.IsSubProcessNodeExpanded = false

	__Process__00000001_.Name = `P1 (P0.A)`
	__Process__00000001_.ComputedPrefix = ``
	__Process__00000001_.IsSubProcessNodeExpanded = false

	__Process__00000002_.Name = `P2 (P0.B)`
	__Process__00000002_.ComputedPrefix = ``
	__Process__00000002_.IsSubProcessNodeExpanded = false

	__Process__00000003_.Name = `P3`
	__Process__00000003_.ComputedPrefix = ``
	__Process__00000003_.IsSubProcessNodeExpanded = false

	// insertion point for setup of pointers
	__Library__00000000_.RootProcesses = append(__Library__00000000_.RootProcesses, __Process__00000000_)
	__Library__00000000_.RootProcesses = append(__Library__00000000_.RootProcesses, __Process__00000003_)
	__Library__00000000_.ProcesssWhoseNodeIsExpanded = append(__Library__00000000_.ProcesssWhoseNodeIsExpanded, __Process__00000000_)
	__Process__00000000_.SubProcesses = append(__Process__00000000_.SubProcesses, __Process__00000001_)
	__Process__00000000_.SubProcesses = append(__Process__00000000_.SubProcesses, __Process__00000002_)
	__Process__00000000_.DiagramProcesss = append(__Process__00000000_.DiagramProcesss, __DiagramProcess__00000000_)
}
