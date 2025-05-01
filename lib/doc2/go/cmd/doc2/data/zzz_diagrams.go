package main

import (
	"time"

	"github.com/fullstack-lang/gong/lib/doc2/go/models"

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

	__Classdiagram__000000_Default := (&models.Classdiagram{}).Stage(stage)

	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_01T10_36_09Z := (&models.DiagramPackage{}).Stage(stage)

	// Setup of values

	__Classdiagram__000000_Default.Name = `Default`
	__Classdiagram__000000_Default.IsInDrawMode = false

	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_01T10_36_09Z.Name = `Diagram Package created the 2025-05-01T10:36:09Z`
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_01T10_36_09Z.Path = ``
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_01T10_36_09Z.GongModelPath = ``
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_01T10_36_09Z.IsEditable = false
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_01T10_36_09Z.IsReloaded = false
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_01T10_36_09Z.AbsolutePathToDiagramPackage = ``

	// Setup of pointers
	// setup of Classdiagram instances pointers
	// setup of DiagramPackage instances pointers
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_01T10_36_09Z.Classdiagrams = append(__DiagramPackage__000000_Diagram_Package_created_the_2025_05_01T10_36_09Z.Classdiagrams, __Classdiagram__000000_Default)
}
