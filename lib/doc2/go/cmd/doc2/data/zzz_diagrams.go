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

	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_02T09_21_37Z := (&models.DiagramPackage{}).Stage(stage)

	// Setup of values

	__Classdiagram__000000_Default.Name = `Default`
	__Classdiagram__000000_Default.IsInDrawMode = true
	__Classdiagram__000000_Default.IsInRenameMode = false
	__Classdiagram__000000_Default.IsExpanded = true
	__Classdiagram__000000_Default.NodeGongStructsIsExpanded = false
	__Classdiagram__000000_Default.NodeGongStructsBinaryEncoding = 0
	__Classdiagram__000000_Default.NodeGongEnumsIsExpanded = true
	__Classdiagram__000000_Default.NodeGongNotesIsExpanded = false

	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_02T09_21_37Z.Name = `Diagram Package created the 2025-05-02T09:21:37Z`
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_02T09_21_37Z.Path = ``
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_02T09_21_37Z.GongModelPath = ``
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_02T09_21_37Z.IsEditable = false
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_02T09_21_37Z.IsReloaded = false
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_02T09_21_37Z.AbsolutePathToDiagramPackage = ``

	// Setup of pointers
	// setup of Classdiagram instances pointers
	// setup of DiagramPackage instances pointers
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_02T09_21_37Z.Classdiagrams = append(__DiagramPackage__000000_Diagram_Package_created_the_2025_05_02T09_21_37Z.Classdiagrams, __Classdiagram__000000_Default)
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_02T09_21_37Z.SelectedClassdiagram = __Classdiagram__000000_Default
}
