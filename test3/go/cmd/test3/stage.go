package main

import (
	"time"

	"github.com/fullstack-lang/gong/test3/go/models"
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
func _(stage *models.StageStruct) {

	// Declaration of instances to stage

	__A__000000_ := (&models.A{}).Stage(stage)
	__A__000001_ := (&models.A{}).Stage(stage)
	__A__000002_ := (&models.A{}).Stage(stage)
	__A__000003_ := (&models.A{}).Stage(stage)
	__A__000004_ddd := (&models.A{}).Stage(stage)

	// Setup of values

	__A__000000_.Name = `1`

	__A__000001_.Name = `2`

	__A__000002_.Name = `3`

	__A__000003_.Name = `4`

	__A__000004_ddd.Name = `ddd`

	// Setup of pointers
}
