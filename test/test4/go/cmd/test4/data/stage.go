package main

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/test/test4/go/models"
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

	__Astruct__00000000_ := (&models.Astruct{Name: `A0`}).Stage(stage)

	// insertion point for initialization of values

	__Astruct__00000000_.Name = `A0`

	// insertion point for setup of pointers
}
