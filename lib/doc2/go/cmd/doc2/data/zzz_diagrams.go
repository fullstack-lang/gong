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

	__AttributeShape__000000_Name := (&models.AttributeShape{}).Stage(stage)

	__Classdiagram__000000_Default := (&models.Classdiagram{}).Stage(stage)
	__Classdiagram__000001_Default_1 := (&models.Classdiagram{}).Stage(stage)

	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_02T09_21_37Z := (&models.DiagramPackage{}).Stage(stage)

	__GongStructShape__000000_Default_Classdiagram := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000001_Default_GongStructShape := (&models.GongStructShape{}).Stage(stage)

	__LinkShape__000000_GongStructShapes := (&models.LinkShape{}).Stage(stage)

	// Setup of values

	__AttributeShape__000000_Name.Name = `Name`

	//gong:ident [ref_models.Classdiagram.Name] comment added to overcome the problem with the comment map association
	__AttributeShape__000000_Name.Identifier = `ref_models.Classdiagram.Name`
	__AttributeShape__000000_Name.FieldTypeAsString = ``
	__AttributeShape__000000_Name.Structname = `Classdiagram`
	__AttributeShape__000000_Name.Fieldtypename = `string`

	__Classdiagram__000000_Default.Name = `Default`
	__Classdiagram__000000_Default.IsInDrawMode = true
	__Classdiagram__000000_Default.IsExpanded = true

	__Classdiagram__000001_Default_1.Name = `Default_1`
	__Classdiagram__000001_Default_1.IsInDrawMode = true
	__Classdiagram__000001_Default_1.IsExpanded = false

	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_02T09_21_37Z.Name = `Diagram Package created the 2025-05-02T09:21:37Z`
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_02T09_21_37Z.Path = ``
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_02T09_21_37Z.GongModelPath = ``
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_02T09_21_37Z.IsEditable = false
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_02T09_21_37Z.IsReloaded = false
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_02T09_21_37Z.AbsolutePathToDiagramPackage = ``

	__GongStructShape__000000_Default_Classdiagram.Name = `Default-Classdiagram`
	__GongStructShape__000000_Default_Classdiagram.X = 43.000000
	__GongStructShape__000000_Default_Classdiagram.Y = 128.000000

	//gong:ident [ref_models.Classdiagram] comment added to overcome the problem with the comment map association
	__GongStructShape__000000_Default_Classdiagram.Identifier = `ref_models.Classdiagram`
	__GongStructShape__000000_Default_Classdiagram.ShowNbInstances = false
	__GongStructShape__000000_Default_Classdiagram.NbInstances = 0
	__GongStructShape__000000_Default_Classdiagram.Width = 240.000000
	__GongStructShape__000000_Default_Classdiagram.Height = 78.000000
	__GongStructShape__000000_Default_Classdiagram.IsSelected = false
	__GongStructShape__000000_Default_Classdiagram.IsExpanded = true

	__GongStructShape__000001_Default_GongStructShape.Name = `Default-GongStructShape`
	__GongStructShape__000001_Default_GongStructShape.X = 753.000000
	__GongStructShape__000001_Default_GongStructShape.Y = 281.000000

	//gong:ident [ref_models.GongStructShape] comment added to overcome the problem with the comment map association
	__GongStructShape__000001_Default_GongStructShape.Identifier = `ref_models.GongStructShape`
	__GongStructShape__000001_Default_GongStructShape.ShowNbInstances = false
	__GongStructShape__000001_Default_GongStructShape.NbInstances = 0
	__GongStructShape__000001_Default_GongStructShape.Width = 240.000000
	__GongStructShape__000001_Default_GongStructShape.Height = 63.000000
	__GongStructShape__000001_Default_GongStructShape.IsSelected = false
	__GongStructShape__000001_Default_GongStructShape.IsExpanded = false

	__LinkShape__000000_GongStructShapes.Name = `GongStructShapes`

	//gong:ident [ref_models.Classdiagram.GongStructShapes] comment added to overcome the problem with the comment map association
	__LinkShape__000000_GongStructShapes.Identifier = `ref_models.Classdiagram.GongStructShapes`

	//gong:ident [ref_models.GongStructShape] comment added to overcome the problem with the comment map association
	__LinkShape__000000_GongStructShapes.Fieldtypename = `ref_models.GongStructShape`
	__LinkShape__000000_GongStructShapes.FieldOffsetX = 0.000000
	__LinkShape__000000_GongStructShapes.FieldOffsetY = 0.000000
	__LinkShape__000000_GongStructShapes.TargetMultiplicity = models.MANY
	__LinkShape__000000_GongStructShapes.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__000000_GongStructShapes.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__000000_GongStructShapes.SourceMultiplicity = models.MANY
	__LinkShape__000000_GongStructShapes.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__000000_GongStructShapes.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__000000_GongStructShapes.X = 678.500000
	__LinkShape__000000_GongStructShapes.Y = 155.500000
	__LinkShape__000000_GongStructShapes.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000000_GongStructShapes.StartRatio = 0.500000
	__LinkShape__000000_GongStructShapes.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000000_GongStructShapes.EndRatio = 0.500000
	__LinkShape__000000_GongStructShapes.CornerOffsetRatio = 4.865625

	// Setup of pointers
	// setup of AttributeShape instances pointers
	// setup of Classdiagram instances pointers
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000000_Default_Classdiagram)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000001_Default_GongStructShape)
	// setup of DiagramPackage instances pointers
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_02T09_21_37Z.Classdiagrams = append(__DiagramPackage__000000_Diagram_Package_created_the_2025_05_02T09_21_37Z.Classdiagrams, __Classdiagram__000000_Default)
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_02T09_21_37Z.Classdiagrams = append(__DiagramPackage__000000_Diagram_Package_created_the_2025_05_02T09_21_37Z.Classdiagrams, __Classdiagram__000001_Default_1)
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_02T09_21_37Z.SelectedClassdiagram = __Classdiagram__000000_Default
	// setup of GongStructShape instances pointers
	__GongStructShape__000000_Default_Classdiagram.AttributeShapes = append(__GongStructShape__000000_Default_Classdiagram.AttributeShapes, __AttributeShape__000000_Name)
	__GongStructShape__000000_Default_Classdiagram.LinkShapes = append(__GongStructShape__000000_Default_Classdiagram.LinkShapes, __LinkShape__000000_GongStructShapes)
	// setup of LinkShape instances pointers
}
