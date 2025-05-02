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

	__GongEnumShape__000000_Default_GongEnumShapeType := (&models.GongEnumShape{}).Stage(stage)
	__GongEnumShape__000001_Default_MultiplicityType := (&models.GongEnumShape{}).Stage(stage)
	__GongEnumShape__000002_Default_ButtonType := (&models.GongEnumShape{}).Stage(stage)

	// Setup of values

	__Classdiagram__000000_Default.Name = `Default`
	__Classdiagram__000000_Default.IsInDrawMode = true
	__Classdiagram__000000_Default.IsInRenameMode = false
	__Classdiagram__000000_Default.IsExpanded = true
	__Classdiagram__000000_Default.NodeNamedStructsIsExpanded = false
	__Classdiagram__000000_Default.NodeGongEnumsIsExpanded = true
	__Classdiagram__000000_Default.NodeGongNotesIsExpanded = false

	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_02T09_21_37Z.Name = `Diagram Package created the 2025-05-02T09:21:37Z`
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_02T09_21_37Z.Path = ``
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_02T09_21_37Z.GongModelPath = ``
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_02T09_21_37Z.IsEditable = false
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_02T09_21_37Z.IsReloaded = false
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_02T09_21_37Z.AbsolutePathToDiagramPackage = ``

	__GongEnumShape__000000_Default_GongEnumShapeType.Name = `Default-GongEnumShapeType`
	__GongEnumShape__000000_Default_GongEnumShapeType.X = 27.000000
	__GongEnumShape__000000_Default_GongEnumShapeType.Y = 38.000000

	//gong:ident [ref_models.GongEnumShapeType] comment added to overcome the problem with the comment map association
	__GongEnumShape__000000_Default_GongEnumShapeType.Identifier = `ref_models.GongEnumShapeType`
	__GongEnumShape__000000_Default_GongEnumShapeType.Width = 240.000000
	__GongEnumShape__000000_Default_GongEnumShapeType.Height = 63.000000
	__GongEnumShape__000000_Default_GongEnumShapeType.IsExpanded = false

	__GongEnumShape__000001_Default_MultiplicityType.Name = `Default-MultiplicityType`
	__GongEnumShape__000001_Default_MultiplicityType.X = 45.000000
	__GongEnumShape__000001_Default_MultiplicityType.Y = 105.000000

	//gong:ident [ref_models.MultiplicityType] comment added to overcome the problem with the comment map association
	__GongEnumShape__000001_Default_MultiplicityType.Identifier = `ref_models.MultiplicityType`
	__GongEnumShape__000001_Default_MultiplicityType.Width = 240.000000
	__GongEnumShape__000001_Default_MultiplicityType.Height = 63.000000
	__GongEnumShape__000001_Default_MultiplicityType.IsExpanded = false

	__GongEnumShape__000002_Default_ButtonType.Name = `Default-ButtonType`
	__GongEnumShape__000002_Default_ButtonType.X = 74.000000
	__GongEnumShape__000002_Default_ButtonType.Y = 51.000000

	//gong:ident [ref_models.ButtonType] comment added to overcome the problem with the comment map association
	__GongEnumShape__000002_Default_ButtonType.Identifier = `ref_models.ButtonType`
	__GongEnumShape__000002_Default_ButtonType.Width = 240.000000
	__GongEnumShape__000002_Default_ButtonType.Height = 63.000000
	__GongEnumShape__000002_Default_ButtonType.IsExpanded = true

	// Setup of pointers
	// setup of Classdiagram instances pointers
	__Classdiagram__000000_Default.GongEnumShapes = append(__Classdiagram__000000_Default.GongEnumShapes, __GongEnumShape__000000_Default_GongEnumShapeType)
	__Classdiagram__000000_Default.GongEnumShapes = append(__Classdiagram__000000_Default.GongEnumShapes, __GongEnumShape__000001_Default_MultiplicityType)
	__Classdiagram__000000_Default.GongEnumShapes = append(__Classdiagram__000000_Default.GongEnumShapes, __GongEnumShape__000002_Default_ButtonType)
	// setup of DiagramPackage instances pointers
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_02T09_21_37Z.Classdiagrams = append(__DiagramPackage__000000_Diagram_Package_created_the_2025_05_02T09_21_37Z.Classdiagrams, __Classdiagram__000000_Default)
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_02T09_21_37Z.SelectedClassdiagram = __Classdiagram__000000_Default
	// setup of GongEnumShape instances pointers
}
