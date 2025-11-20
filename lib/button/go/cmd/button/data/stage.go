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

	const __write__local_time = "2025-11-20 05:41:56.599777 CET"
	const __write__utc_time__ = "2025-11-20 04:41:56.599777 UTC"

	const __commitId__ = "0000000027"

	// Declaration of instances to stage

	__Button__000000_One := (&models.Button{}).Stage(stage)
	__Button__000001_Two := (&models.Button{}).Stage(stage)
	__Button__000002_Three := (&models.Button{}).Stage(stage)
	__Button__000003_Four := (&models.Button{}).Stage(stage)
	__Button__000004_Five := (&models.Button{}).Stage(stage)
	__Button__000005_Un := (&models.Button{}).Stage(stage)

	__Group__000000_Group_1 := (&models.Group{}).Stage(stage)
	__Group__000001_ := (&models.Group{}).Stage(stage)

	__Layout__000000_layout := (&models.Layout{}).Stage(stage)

	// Setup of values

	__Button__000000_One.Name = `One`
	__Button__000000_One.Label = `One`
	__Button__000000_One.Icon = `draw`

	__Button__000001_Two.Name = `Two`
	__Button__000001_Two.Label = `Two`
	__Button__000001_Two.Icon = `add`

	__Button__000002_Three.Name = `Three`
	__Button__000002_Three.Label = `Three`
	__Button__000002_Three.Icon = `add_box`

	__Button__000003_Four.Name = `Four`
	__Button__000003_Four.Label = `Four`
	__Button__000003_Four.Icon = ``

	__Button__000004_Five.Name = `Five`
	__Button__000004_Five.Label = `Five`
	__Button__000004_Five.Icon = ``

	__Button__000005_Un.Name = `Un`
	__Button__000005_Un.Label = `Un`
	__Button__000005_Un.Icon = `hour`

	__Group__000000_Group_1.Name = `Group 1`
	__Group__000000_Group_1.Percentage = 50.000000
	__Group__000000_Group_1.NbColumns = 15

	__Group__000001_.Name = ``
	__Group__000001_.Percentage = 50.000000
	__Group__000001_.NbColumns = 5

	__Layout__000000_layout.Name = `layout`

	// Setup of pointers
	// setup of Button instances pointers
	// setup of Group instances pointers
	__Group__000000_Group_1.Buttons = append(__Group__000000_Group_1.Buttons, __Button__000000_One)
	__Group__000000_Group_1.Buttons = append(__Group__000000_Group_1.Buttons, __Button__000001_Two)
	__Group__000000_Group_1.Buttons = append(__Group__000000_Group_1.Buttons, __Button__000002_Three)
	__Group__000000_Group_1.Buttons = append(__Group__000000_Group_1.Buttons, __Button__000003_Four)
	__Group__000000_Group_1.Buttons = append(__Group__000000_Group_1.Buttons, __Button__000004_Five)
	__Group__000001_.Buttons = append(__Group__000001_.Buttons, __Button__000005_Un)
	// setup of Layout instances pointers
	__Layout__000000_layout.Groups = append(__Layout__000000_layout.Groups, __Group__000000_Group_1)
	__Layout__000000_layout.Groups = append(__Layout__000000_layout.Groups, __Group__000001_)
}

