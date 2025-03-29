package main

import (
	"time"

	"github.com/fullstack-lang/gong/lib/load/go/models"
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

	__FileToDownload__000000_a := (&models.FileToDownload{}).Stage(stage)

	// Setup of values

	__FileToDownload__000000_a.Name = `a`
	__FileToDownload__000000_a.Content = `a`

	// Setup of pointers
	// setup of FileToDownload instances pointers
}
