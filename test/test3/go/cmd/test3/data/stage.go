package main

import (
	"time"

	"github.com/fullstack-lang/gong/test/test3/go/models"
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

	const __write__local_time = "2025-06-14 01:39:30.312921 CEST"
	const __write__utc_time__ = "2025-06-13 23:39:30.312921 UTC"

	const __commitId__ = "0000000005"

	// Declaration of instances to stage

	__A__000000_A1 := (&models.A{}).Stage(stage)
	__A__000001_A2 := (&models.A{}).Stage(stage)

	// Setup of values

	__A__000000_A1.Name = `A1`

	__A__000001_A2.Name = `A2`

	// Setup of pointers
	// setup of A instances pointers
	__A__000000_A1.As = append(__A__000000_A1.As, __A__000001_A2)
}

