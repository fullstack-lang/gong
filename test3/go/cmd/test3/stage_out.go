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

	__A__000000_f := (&models.A{}).Stage(stage)
	__A__000001_f := (&models.A{}).Stage(stage)
	__A__000002_g := (&models.A{}).Stage(stage)
	__A__000003_q := (&models.A{}).Stage(stage)
	__A__000004_q := (&models.A{}).Stage(stage)
	__A__000005_q := (&models.A{}).Stage(stage)
	__A__000006_q := (&models.A{}).Stage(stage)

	// Setup of values

	__A__000000_f.Name = `f`

	__A__000001_f.Name = `f`

	__A__000002_g.Name = `g`

	__A__000003_q.Name = `q`

	__A__000004_q.Name = `q`

	__A__000005_q.Name = `q`

	__A__000006_q.Name = `q`

	// Setup of pointers
}
