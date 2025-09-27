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

	const __write__local_time = "2025-09-27 03:09:25.711821 CEST"
	const __write__utc_time__ = "2025-09-27 01:09:25.711821 UTC"

	const __commitId__ = "0000000007"

	// Declaration of instances to stage

	__FileToUpload__000000_erze_txt := (&models.FileToUpload{}).Stage(stage)

	// Setup of values

	__FileToUpload__000000_erze_txt.Name = `erze.txt`
	__FileToUpload__000000_erze_txt.Base64EncodedContent = `ZXJ6`

	// Setup of pointers
	// setup of FileToUpload instances pointers
}

