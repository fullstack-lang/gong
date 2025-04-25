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

	__Page__000000_First_Steps := (&models.Page{}).Stage(stage)

	// Setup of values

	__Chapter__000000_Getting_Started.Name = `Getting Started`
	__Chapter__000000_Getting_Started.MardownContent = `This is the introduction to the first chapter.
We'll cover the basics here.

This demonstrates how new lines are generated`

	__Chapter__000001_Advanced_Topics.Name = `Advanced Topics`
	__Chapter__000001_Advanced_Topics.MardownContent = `Moving onto more advanced concepts in the second chapter.`

	__Content__000000_My_Awesome_Go_Book.Name = `My Awesome Go Book`
	__Content__000000_My_Awesome_Go_Book.MardownContent = `Welcome to the main page of the book built with our custom Go SSG!`
	__Content__000000_My_Awesome_Go_Book.ContentPath = `content`
	__Content__000000_My_Awesome_Go_Book.OutputPath = `public`
	__Content__000000_My_Awesome_Go_Book.LayoutPath = `../../defaults/layouts`
	__Content__000000_My_Awesome_Go_Book.StaticPath = `../../defaults/static`
	__Content__000000_My_Awesome_Go_Book.Target = models.FILE

	__Page__000000_First_Steps.Name = `First Steps`
	__Page__000000_First_Steps.MardownContent = `Here is the detailed content for the first page within Chapter 1.

You can use standard Markdown formatting:

* Lists
* Are
* Easy

And include "code blocks".

Here is the logo embedded using HTML:

<img src="../../images/gong logo.svg" alt="My Logo" style="height: 200px; width: auto;">

You can place text around it.

| Header 1 | Header 2 | Header 3 |
| -------- | -------- | -------- |
| Cell 1-1 | Cell 1-2 | Cell 1-3 |
| Cell 2-1 | Cell 2-2 | Cell 2-3 |
| Cell 2-1 | Cell 2-2 | Cell 2-3 |

Legend of the table`

	// Setup of pointers
	// setup of Chapter instances pointers
	__Chapter__000000_Getting_Started.Pages = append(__Chapter__000000_Getting_Started.Pages, __Page__000000_First_Steps)
	// setup of Content instances pointers
	__Content__000000_My_Awesome_Go_Book.Chapters = append(__Content__000000_My_Awesome_Go_Book.Chapters, __Chapter__000000_Getting_Started)
	__Content__000000_My_Awesome_Go_Book.Chapters = append(__Content__000000_My_Awesome_Go_Book.Chapters, __Chapter__000001_Advanced_Topics)
	// setup of Page instances pointers
}
