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

	// insertion point for declaration of instances to stage

	__Process__00000000_ := (&models.Process{Name: `Process 1`}).Stage(stage)

	// insertion point for initialization of values

	__Process__00000000_.Name = `Process 1`

	// insertion point for setup of pointers

	__Library__00000000_ := (&models.Library{Name: ``}).Stage(stage)
	__Library__00000000_.Name = ``
	__Library__00000000_.ComputedPrefix = ``
	__Library__00000000_.IsInRenameMode = false
	__Library__00000000_.IsExpanded = false
	__Library__00000000_.NbPixPerCharacter = 0.000000
	__Library__00000000_.LogoSVGFile = ``
	stage.Commit()

	__Library__00000001_ := (&models.Library{Name: ``}).Stage(stage)
	// 
	__Library__00000000_.IsExpanded = true
	__Library__00000000_.SubLibraries = slices.Insert( __Library__00000000_.SubLibraries, 0, __Library__00000001_)
	__Library__00000001_.Name = ``
	__Library__00000001_.ComputedPrefix = ``
	__Library__00000001_.IsInRenameMode = true
	__Library__00000001_.IsExpanded = false
	__Library__00000001_.NbPixPerCharacter = 0.000000
	__Library__00000001_.LogoSVGFile = ``
	stage.Commit()

	// L1
	__Library__00000001_.Name = `L1`
	__Library__00000001_.IsInRenameMode = false
	stage.Commit()

	__DiagramProcess__00000000_ := (&models.DiagramProcess{Name: ``}).Stage(stage)
	// 
	__Library__00000000_.DiagramProcesss = slices.Insert( __Library__00000000_.DiagramProcesss, 0, __DiagramProcess__00000000_)
	__DiagramProcess__00000000_.Name = ``
	__DiagramProcess__00000000_.ComputedPrefix = ``
	__DiagramProcess__00000000_.IsInRenameMode = true
	__DiagramProcess__00000000_.IsExpanded = true
	__DiagramProcess__00000000_.IsChecked = true
	__DiagramProcess__00000000_.IsEditable_ = true
	__DiagramProcess__00000000_.IsShowPrefix = false
	__DiagramProcess__00000000_.DefaultBoxWidth = 0.000000
	__DiagramProcess__00000000_.DefaultBoxHeigth = 0.000000
	__DiagramProcess__00000000_.Width = 0.000000
	__DiagramProcess__00000000_.Height = 0.000000
	__DiagramProcess__00000000_.IsProcesssNodeExpanded = false
	stage.Commit()

	// D1
	__DiagramProcess__00000000_.Name = `D1`
	__DiagramProcess__00000000_.IsInRenameMode = false
	stage.Commit()

	// 
	__Library__00000000_.SubLibraries = slices.Delete( __Library__00000000_.SubLibraries, 0, 1)
	__Library__00000001_.Unstage(stage)
	stage.Commit()
}
