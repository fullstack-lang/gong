package main

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/dsm/structure/go/models"
	// injection point for ident package import declaration{{ImportPackageDeclaration}}
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var (
	_ time.Time
	_ = slices.Index[[]int, int]
)

// _ point for meta package dummy declaration{{ImportPackageDummyDeclaration}}

// When parsed, those maps will help with the renaming process
var _ map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// function will stage objects
func _(stage *models.Stage) {

	// insertion point for declaration of instances to stage

	__DiagramStructure__00000000_ := (&models.DiagramStructure{Name: `DiagramStructure`}).Stage(stage)
	__DiagramStructure__00000001_ := (&models.DiagramStructure{Name: ``}).Stage(stage)

	__Library__00000000_ := (&models.Library{Name: ``}).Stage(stage)

	__System__00000000_ := (&models.System{Name: ``}).Stage(stage)

	// insertion point for initialization of values

	__DiagramStructure__00000000_.Name = `DiagramStructure`
	__DiagramStructure__00000000_.ComputedPrefix = ``
	__DiagramStructure__00000000_.IsExpanded = false
	__DiagramStructure__00000000_.LayoutDirection = models.Vertical
	__DiagramStructure__00000000_.IsChecked = false
	__DiagramStructure__00000000_.IsEditable_ = true
	__DiagramStructure__00000000_.IsShowPrefix = false
	__DiagramStructure__00000000_.Width = 0.000000
	__DiagramStructure__00000000_.Height = 0.000000
	__DiagramStructure__00000000_.DefaultBoxWidth = 0.000000
	__DiagramStructure__00000000_.DefaultBoxHeigth = 0.000000
	__DiagramStructure__00000000_.IsPartsNodeExpanded = false

	__DiagramStructure__00000001_.Name = ``
	__DiagramStructure__00000001_.ComputedPrefix = ``
	__DiagramStructure__00000001_.IsExpanded = false
	__DiagramStructure__00000001_.LayoutDirection = models.Vertical
	__DiagramStructure__00000001_.IsChecked = false
	__DiagramStructure__00000001_.IsEditable_ = false
	__DiagramStructure__00000001_.IsShowPrefix = false
	__DiagramStructure__00000001_.Width = 0.000000
	__DiagramStructure__00000001_.Height = 0.000000
	__DiagramStructure__00000001_.DefaultBoxWidth = 0.000000
	__DiagramStructure__00000001_.DefaultBoxHeigth = 0.000000
	__DiagramStructure__00000001_.IsPartsNodeExpanded = false

	__Library__00000000_.Name = ``
	__Library__00000000_.IsSubLibrariesNodeExpanded = false
	__Library__00000000_.NbPixPerCharacter = 0.000000
	__Library__00000000_.LogoSVGFile = ``
	__Library__00000000_.ComputedPrefix = ``
	__Library__00000000_.IsExpanded = false
	__Library__00000000_.LayoutDirection = models.Vertical
	__Library__00000000_.IsRootLibrary = true
	__Library__00000000_.IsSystemsNodeExpanded = true

	__System__00000000_.Name = ``
	__System__00000000_.ComputedPrefix = ``
	__System__00000000_.IsExpanded = false
	__System__00000000_.LayoutDirection = models.Vertical
	__System__00000000_.IsPartsNodeExpanded = false
	__System__00000000_.IsSubSystemsNodeExpanded = false
	__System__00000000_.IsLinksNodeExpanded = false
	__System__00000000_.IsDiagramStructuresNodeExpanded = true

	// insertion point for setup of pointers
	__Library__00000000_.RootSystems = append(__Library__00000000_.RootSystems, __System__00000000_)
	__System__00000000_.DiagramStructures = append(__System__00000000_.DiagramStructures, __DiagramStructure__00000000_)
	__System__00000000_.DiagramStructures = append(__System__00000000_.DiagramStructures, __DiagramStructure__00000001_)
}
