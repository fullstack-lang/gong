package main

import (
	"time"

	"github.com/fullstack-lang/gong/test/statemachines/go/models"
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

	// Declaration of instances to stage

	__Architecture__000000_Gong_UX_loop_Architecture := (&models.Architecture{}).Stage(stage)

	// Setup of values

	__Architecture__000000_Gong_UX_loop_Architecture.Name = `Gong UX loop Architecture`
	__Architecture__000000_Gong_UX_loop_Architecture.NbPixPerCharacter = 0.000000

	// Setup of pointers
	// setup of Architecture instances pointers
}

