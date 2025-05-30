package main

import (
	"time"

	"github.com/fullstack-lang/gong/lib/probe/go/models"

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

	__Probe2__000000_Probe := (&models.Probe2{}).Stage(stage)

	// Setup of values

	__Probe2__000000_Probe.Name = `Probe`

	// Setup of pointers
	// setup of Probe2 instances pointers
}
