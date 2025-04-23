package main

import (
	"time"

	"github.com/fullstack-lang/gong/lib/ssg/go/models"

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

	__Chapter__000000_Chapter_1 := (&models.Chapter{}).Stage(stage)
	__Chapter__000001_Chapter_2 := (&models.Chapter{}).Stage(stage)

	__Content__000000_content := (&models.Content{}).Stage(stage)

	// Setup of values

	__Chapter__000000_Chapter_1.Name = `Chapter 1`

	__Chapter__000001_Chapter_2.Name = `Chapter 2`

	__Content__000000_content.Name = `content`

	// Setup of pointers
	// setup of Chapter instances pointers
	// setup of Content instances pointers
	__Content__000000_content.Chapters = append(__Content__000000_content.Chapters, __Chapter__000000_Chapter_1)
	__Content__000000_content.Chapters = append(__Content__000000_content.Chapters, __Chapter__000001_Chapter_2)
}
