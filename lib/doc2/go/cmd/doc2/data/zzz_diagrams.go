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

	__Classdiagram__000000_Defaultxxx := (&models.Classdiagram{}).Stage(stage)

	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_02T09_21_37Z := (&models.DiagramPackage{}).Stage(stage)

	__GongStructShape__000000_Default_AttributeShape := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000001_Default_Classdiagram := (&models.GongStructShape{}).Stage(stage)

	// Setup of values

	__Classdiagram__000000_Defaultxxx.Name = `Defaultxxx`
	__Classdiagram__000000_Defaultxxx.IsInDrawMode = true
	__Classdiagram__000000_Defaultxxx.IsInRenameMode = false
	__Classdiagram__000000_Defaultxxx.IsExpanded = true
	__Classdiagram__000000_Defaultxxx.NodeNamedStructsIsExpanded = false
	__Classdiagram__000000_Defaultxxx.NodeGongEnumsIsExpanded = false
	__Classdiagram__000000_Defaultxxx.NodeGongNotesIsExpanded = false

	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_02T09_21_37Z.Name = `Diagram Package created the 2025-05-02T09:21:37Z`
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_02T09_21_37Z.Path = ``
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_02T09_21_37Z.GongModelPath = ``
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_02T09_21_37Z.IsEditable = false
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_02T09_21_37Z.IsReloaded = false
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_02T09_21_37Z.AbsolutePathToDiagramPackage = ``

	__GongStructShape__000000_Default_AttributeShape.Name = `Default-AttributeShape`
	__GongStructShape__000000_Default_AttributeShape.X = 23.000000
	__GongStructShape__000000_Default_AttributeShape.Y = 47.000000

	//gong:ident [ref_models.AttributeShape] comment added to overcome the problem with the comment map association
	__GongStructShape__000000_Default_AttributeShape.Identifier = `ref_models.AttributeShape`
	__GongStructShape__000000_Default_AttributeShape.ShowNbInstances = false
	__GongStructShape__000000_Default_AttributeShape.NbInstances = 0
	__GongStructShape__000000_Default_AttributeShape.Width = 240.000000
	__GongStructShape__000000_Default_AttributeShape.Height = 63.000000
	__GongStructShape__000000_Default_AttributeShape.IsSelected = false
	__GongStructShape__000000_Default_AttributeShape.IsExpanded = false

	__GongStructShape__000001_Default_Classdiagram.Name = `Default-Classdiagram`
	__GongStructShape__000001_Default_Classdiagram.X = 67.000000
	__GongStructShape__000001_Default_Classdiagram.Y = 90.000000

	//gong:ident [ref_models.Classdiagram] comment added to overcome the problem with the comment map association
	__GongStructShape__000001_Default_Classdiagram.Identifier = `ref_models.Classdiagram`
	__GongStructShape__000001_Default_Classdiagram.ShowNbInstances = false
	__GongStructShape__000001_Default_Classdiagram.NbInstances = 0
	__GongStructShape__000001_Default_Classdiagram.Width = 240.000000
	__GongStructShape__000001_Default_Classdiagram.Height = 63.000000
	__GongStructShape__000001_Default_Classdiagram.IsSelected = false
	__GongStructShape__000001_Default_Classdiagram.IsExpanded = false

	// Setup of pointers
	// setup of Classdiagram instances pointers
	__Classdiagram__000000_Defaultxxx.GongStructShapes = append(__Classdiagram__000000_Defaultxxx.GongStructShapes, __GongStructShape__000000_Default_AttributeShape)
	__Classdiagram__000000_Defaultxxx.GongStructShapes = append(__Classdiagram__000000_Defaultxxx.GongStructShapes, __GongStructShape__000001_Default_Classdiagram)
	// setup of DiagramPackage instances pointers
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_02T09_21_37Z.Classdiagrams = append(__DiagramPackage__000000_Diagram_Package_created_the_2025_05_02T09_21_37Z.Classdiagrams, __Classdiagram__000000_Defaultxxx)
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_02T09_21_37Z.SelectedClassdiagram = __Classdiagram__000000_Defaultxxx
	// setup of GongStructShape instances pointers
}
