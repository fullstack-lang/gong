package main

import (
	"time"

	"github.com/fullstack-lang/gong/lib/button/go/models"
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

	__Button__000000_One := (&models.Button{}).Stage(stage)

	__Group__000000_group := (&models.Group{}).Stage(stage)

	__Layout__000000_layout := (&models.Layout{}).Stage(stage)

	// Setup of values

	__Button__000000_One.Name = `One`
	__Button__000000_One.Label = `One`
	__Button__000000_One.Icon = `draw`

	__Group__000000_group.Name = `group`
	__Group__000000_group.Percentage = 100.000000

	__Layout__000000_layout.Name = `layout`

	// Setup of pointers
	// setup of Button instances pointers
	// setup of Group instances pointers
	__Group__000000_group.Buttons = append(__Group__000000_group.Buttons, __Button__000000_One)
	// setup of Layout instances pointers
	__Layout__000000_layout.Groups = append(__Layout__000000_layout.Groups, __Group__000000_group)
}
