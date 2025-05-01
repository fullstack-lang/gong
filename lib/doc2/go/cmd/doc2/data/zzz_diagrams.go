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

	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_01T10_49_57Z := (&models.DiagramPackage{}).Stage(stage)

	__Position__000000_Pos_Default_Classdiagram := (&models.Position{}).Stage(stage)
	__Position__000001_Pos_Default_Classdiagram := (&models.Position{}).Stage(stage)
	__Position__000002_Pos_Default_Classdiagram := (&models.Position{}).Stage(stage)
	__Position__000003_Pos_Default_Classdiagram := (&models.Position{}).Stage(stage)
	__Position__000004_Pos_Default_Classdiagram := (&models.Position{}).Stage(stage)
	__Position__000005_Pos_Default_Classdiagram := (&models.Position{}).Stage(stage)
	__Position__000006_Pos_Default_Classdiagram := (&models.Position{}).Stage(stage)
	__Position__000007_Pos_Default_Classdiagram := (&models.Position{}).Stage(stage)
	__Position__000008_Pos_Default_Classdiagram := (&models.Position{}).Stage(stage)

	// Setup of values

	__Classdiagram__000000_Default.Name = `Default`
	__Classdiagram__000000_Default.IsInDrawMode = false
	__Classdiagram__000000_Default.IsExpanded = true

	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_01T10_49_57Z.Name = `Diagram Package created the 2025-05-01T10:49:57Z`
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_01T10_49_57Z.Path = ``
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_01T10_49_57Z.GongModelPath = ``
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_01T10_49_57Z.IsEditable = false
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_01T10_49_57Z.IsReloaded = false
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_01T10_49_57Z.AbsolutePathToDiagramPackage = ``

	__Position__000000_Pos_Default_Classdiagram.X = 97.000000
	__Position__000000_Pos_Default_Classdiagram.Y = 63.000000
	__Position__000000_Pos_Default_Classdiagram.Name = `Pos-Default-Classdiagram`

	__Position__000001_Pos_Default_Classdiagram.X = 32.000000
	__Position__000001_Pos_Default_Classdiagram.Y = 88.000000
	__Position__000001_Pos_Default_Classdiagram.Name = `Pos-Default-Classdiagram`

	__Position__000002_Pos_Default_Classdiagram.X = 49.000000
	__Position__000002_Pos_Default_Classdiagram.Y = 52.000000
	__Position__000002_Pos_Default_Classdiagram.Name = `Pos-Default-Classdiagram`

	__Position__000003_Pos_Default_Classdiagram.X = 60.000000
	__Position__000003_Pos_Default_Classdiagram.Y = 30.000000
	__Position__000003_Pos_Default_Classdiagram.Name = `Pos-Default-Classdiagram`

	__Position__000004_Pos_Default_Classdiagram.X = 16.000000
	__Position__000004_Pos_Default_Classdiagram.Y = 98.000000
	__Position__000004_Pos_Default_Classdiagram.Name = `Pos-Default-Classdiagram`

	__Position__000005_Pos_Default_Classdiagram.X = 82.000000
	__Position__000005_Pos_Default_Classdiagram.Y = 53.000000
	__Position__000005_Pos_Default_Classdiagram.Name = `Pos-Default-Classdiagram`

	__Position__000006_Pos_Default_Classdiagram.X = 88.000000
	__Position__000006_Pos_Default_Classdiagram.Y = 80.000000
	__Position__000006_Pos_Default_Classdiagram.Name = `Pos-Default-Classdiagram`

	__Position__000007_Pos_Default_Classdiagram.X = 52.000000
	__Position__000007_Pos_Default_Classdiagram.Y = 83.000000
	__Position__000007_Pos_Default_Classdiagram.Name = `Pos-Default-Classdiagram`

	__Position__000008_Pos_Default_Classdiagram.X = 32.000000
	__Position__000008_Pos_Default_Classdiagram.Y = 84.000000
	__Position__000008_Pos_Default_Classdiagram.Name = `Pos-Default-Classdiagram`

	// Setup of pointers
	// setup of Classdiagram instances pointers
	// setup of DiagramPackage instances pointers
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_01T10_49_57Z.Classdiagrams = append(__DiagramPackage__000000_Diagram_Package_created_the_2025_05_01T10_49_57Z.Classdiagrams, __Classdiagram__000000_Default)
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_01T10_49_57Z.SelectedClassdiagram = __Classdiagram__000000_Default
	// setup of Position instances pointers
}
