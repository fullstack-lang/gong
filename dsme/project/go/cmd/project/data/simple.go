package main

import (
	"time"

	"github.com/fullstack-lang/gong/dsme/project/go/models"
	// injection point for ident package import declaration{{ImportPackageDeclaration}}
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var _ time.Time

// _ point for meta package dummy declaration{{ImportPackageDummyDeclaration}}

// When parsed, those maps will help with the renaming process
var _ map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// function will stage objects
func _(stage *models.Stage) {

	// insertion point for declaration of instances to stage

	__Library__00000000_ := (&models.Library{Name: `Root Library`}).Stage(stage)
	__Library__00000001_ := (&models.Library{Name: `L1`}).Stage(stage)
	__Library__00000003_ := (&models.Library{Name: `L2`}).Stage(stage)

	// insertion point for initialization of values

	__Library__00000000_.Name = `Root Library`
	__Library__00000000_.ComputedPrefix = ``
	__Library__00000000_.IsInRenameMode = false
	__Library__00000000_.IsExpanded = false
	__Library__00000000_.NbPixPerCharacter = 8.000000

	__Library__00000001_.Name = `L1`
	__Library__00000001_.ComputedPrefix = `1`
	__Library__00000001_.IsInRenameMode = false
	__Library__00000001_.IsExpanded = false
	__Library__00000001_.NbPixPerCharacter = 0.000000

	__Library__00000003_.Name = `L2`
	__Library__00000003_.ComputedPrefix = `1.1`
	__Library__00000003_.IsInRenameMode = false
	__Library__00000003_.IsExpanded = false
	__Library__00000003_.NbPixPerCharacter = 0.000000

	// insertion point for setup of pointers
	__Library__00000000_.SubLibraries = append(__Library__00000000_.SubLibraries, __Library__00000001_)
	__Library__00000001_.SubLibraries = append(__Library__00000001_.SubLibraries, __Library__00000003_)
}
