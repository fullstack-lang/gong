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

	__Chapter__000000_Getting_Started := (&models.Chapter{}).Stage(stage)
	__Chapter__000001_Advanced_Topics := (&models.Chapter{}).Stage(stage)

	__Content__000000_My_Awesome_Go_Book := (&models.Content{}).Stage(stage)

	// Setup of values

	__Chapter__000000_Getting_Started.Name = `Getting Started`

	__Chapter__000001_Advanced_Topics.Name = `Advanced Topics`

	__Content__000000_My_Awesome_Go_Book.Name = `My Awesome Go Book`
	__Content__000000_My_Awesome_Go_Book.Text = `Welcome to the main page of the book built with our custom Go SSG!`
	__Content__000000_My_Awesome_Go_Book.ContentPath = `content`

	// Setup of pointers
	// setup of Chapter instances pointers
	// setup of Content instances pointers
	__Content__000000_My_Awesome_Go_Book.Chapters = append(__Content__000000_My_Awesome_Go_Book.Chapters, __Chapter__000000_Getting_Started)
	__Content__000000_My_Awesome_Go_Book.Chapters = append(__Content__000000_My_Awesome_Go_Book.Chapters, __Chapter__000001_Advanced_Topics)
}
