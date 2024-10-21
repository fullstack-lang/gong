package main

import (
	"time"

	"github.com/fullstack-lang/gong/test3/go/models"

	// injection point for ident package import declaration{{ImportPackageDeclaration}}
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var _ time.Time

// Injection point for meta package dummy declaration{{ImportPackageDummyDeclaration}}

// When parsed, those maps will help with the renaming process
var _ map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// function will stage objects
func _(stage *models.StageStruct) {

	// Declaration of instances to stage

	__A__000000_a := (&models.A{}).Stage(stage)
	__A__000001_b := (&models.A{}).Stage(stage)

	// Setup of values

	__A__000000_a.Name = `a`

	__A__000001_b.Name = `b`

	// Setup of pointers
}
