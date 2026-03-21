package main

import (
	"time"

	"github.com/fullstack-lang/gong/dsme/project/go/models"
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

	// insertion point for declaration of instances to stage

	__Diagram__00000000_ := (&models.Diagram{Name: `Default Diagram`}).Stage(stage)

	__Library__00000000_ := (&models.Library{Name: `Root Library`}).Stage(stage)
	__Library__00000001_ := (&models.Library{Name: ``}).Stage(stage)

	// insertion point for initialization of values

	__Diagram__00000000_.Name = `Default Diagram`
	__Diagram__00000000_.IsExpanded = true
	__Diagram__00000000_.ComputedPrefix = ``
	__Diagram__00000000_.IsInRenameMode = false
	__Diagram__00000000_.IsChecked = true
	__Diagram__00000000_.IsEditable_ = true
	__Diagram__00000000_.ShowPrefix = false
	__Diagram__00000000_.DefaultBoxWidth = 250.000000
	__Diagram__00000000_.DefaultBoxHeigth = 70.000000
	__Diagram__00000000_.Width = 2200.000000
	__Diagram__00000000_.Height = 2200.000000
	__Diagram__00000000_.IsPBSNodeExpanded = false
	__Diagram__00000000_.IsWBSNodeExpanded = false
	__Diagram__00000000_.IsNotesNodeExpanded = false
	__Diagram__00000000_.IsResourcesNodeExpanded = false

	__Library__00000000_.Name = `Root Library`
	__Library__00000000_.IsExpanded = false
	__Library__00000000_.ComputedPrefix = ``
	__Library__00000000_.IsInRenameMode = false
	__Library__00000000_.NbPixPerCharacter = 8.000000

	__Library__00000001_.Name = ``
	__Library__00000001_.IsExpanded = true
	__Library__00000001_.ComputedPrefix = ``
	__Library__00000001_.IsInRenameMode = false
	__Library__00000001_.NbPixPerCharacter = 8.000000

	// insertion point for setup of pointers
	__Diagram__00000000_.OwningLibrary = __Library__00000000_
	__Library__00000000_.OwningLibrary = __Library__00000001_
	__Library__00000001_.OwningLibrary = __Library__00000000_
	__Library__00000001_.Diagrams = append(__Library__00000001_.Diagrams, __Diagram__00000000_)
}
