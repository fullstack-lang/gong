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

	__GongEnumShape__000000_Default_ButtonType := (&models.GongEnumShape{}).Stage(stage)

	__GongEnumValueShape__000000_DUPLICATE := (&models.GongEnumValueShape{}).Stage(stage)
	__GongEnumValueShape__000001_EDIT_CANCEL := (&models.GongEnumValueShape{}).Stage(stage)
	__GongEnumValueShape__000002_EDIT := (&models.GongEnumValueShape{}).Stage(stage)
	__GongEnumValueShape__000003_REMOVE := (&models.GongEnumValueShape{}).Stage(stage)
	__GongEnumValueShape__000004_RENAME_CANCEL := (&models.GongEnumValueShape{}).Stage(stage)

	// Setup of values

	__Classdiagram__000000_Default.Name = `Default`
	__Classdiagram__000000_Default.IsInDrawMode = true
	__Classdiagram__000000_Default.IsInRenameMode = false
	__Classdiagram__000000_Default.IsExpanded = true
	__Classdiagram__000000_Default.NodeGongStructsIsExpanded = true
	__Classdiagram__000000_Default.NodeGongStructNodeExpansionBinaryEncoding = 0
	__Classdiagram__000000_Default.NodeGongEnumsIsExpanded = false
	__Classdiagram__000000_Default.NodeGongEnumNodeExpansionBinaryEncoding = 6
	__Classdiagram__000000_Default.NodeGongNotesIsExpanded = false
	__Classdiagram__000000_Default.NodeGongNoteNodeExpansionBinaryEncoding = 0

	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_02T09_21_37Z.Name = `Diagram Package created the 2025-05-02T09:21:37Z`
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_02T09_21_37Z.Path = ``
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_02T09_21_37Z.GongModelPath = ``
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_02T09_21_37Z.IsEditable = false
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_02T09_21_37Z.IsReloaded = false
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_02T09_21_37Z.AbsolutePathToDiagramPackage = ``

	__GongEnumShape__000000_Default_ButtonType.Name = `Default-ButtonType`
	__GongEnumShape__000000_Default_ButtonType.X = 106.000000
	__GongEnumShape__000000_Default_ButtonType.Y = 105.000000

	//gong:ident [ref_models.ButtonType] comment added to overcome the problem with the comment map association
	__GongEnumShape__000000_Default_ButtonType.Identifier = `ref_models.ButtonType`
	__GongEnumShape__000000_Default_ButtonType.Width = 240.000000
	__GongEnumShape__000000_Default_ButtonType.Height = 163.000000
	__GongEnumShape__000000_Default_ButtonType.IsExpanded = false

	__GongEnumValueShape__000000_DUPLICATE.Name = `DUPLICATE`

	//gong:ident [ref_models.ButtonType.DUPLICATE] comment added to overcome the problem with the comment map association
	__GongEnumValueShape__000000_DUPLICATE.Identifier = `ref_models.ButtonType.DUPLICATE`

	__GongEnumValueShape__000001_EDIT_CANCEL.Name = `EDIT_CANCEL`

	//gong:ident [ref_models.ButtonType.EDIT_CANCEL] comment added to overcome the problem with the comment map association
	__GongEnumValueShape__000001_EDIT_CANCEL.Identifier = `ref_models.ButtonType.EDIT_CANCEL`

	__GongEnumValueShape__000002_EDIT.Name = `EDIT`

	//gong:ident [ref_models.ButtonType.EDIT] comment added to overcome the problem with the comment map association
	__GongEnumValueShape__000002_EDIT.Identifier = `ref_models.ButtonType.EDIT`

	__GongEnumValueShape__000003_REMOVE.Name = `REMOVE`

	//gong:ident [ref_models.ButtonType.REMOVE] comment added to overcome the problem with the comment map association
	__GongEnumValueShape__000003_REMOVE.Identifier = `ref_models.ButtonType.REMOVE`

	__GongEnumValueShape__000004_RENAME_CANCEL.Name = `RENAME_CANCEL`

	//gong:ident [ref_models.ButtonType.RENAME_CANCEL] comment added to overcome the problem with the comment map association
	__GongEnumValueShape__000004_RENAME_CANCEL.Identifier = `ref_models.ButtonType.RENAME_CANCEL`

	// Setup of pointers
	// setup of Classdiagram instances pointers
	__Classdiagram__000000_Default.GongEnumShapes = append(__Classdiagram__000000_Default.GongEnumShapes, __GongEnumShape__000000_Default_ButtonType)
	// setup of DiagramPackage instances pointers
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_02T09_21_37Z.Classdiagrams = append(__DiagramPackage__000000_Diagram_Package_created_the_2025_05_02T09_21_37Z.Classdiagrams, __Classdiagram__000000_Default)
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_02T09_21_37Z.SelectedClassdiagram = __Classdiagram__000000_Default
	// setup of GongEnumShape instances pointers
	__GongEnumShape__000000_Default_ButtonType.GongEnumValueShapes = append(__GongEnumShape__000000_Default_ButtonType.GongEnumValueShapes, __GongEnumValueShape__000000_DUPLICATE)
	__GongEnumShape__000000_Default_ButtonType.GongEnumValueShapes = append(__GongEnumShape__000000_Default_ButtonType.GongEnumValueShapes, __GongEnumValueShape__000001_EDIT_CANCEL)
	__GongEnumShape__000000_Default_ButtonType.GongEnumValueShapes = append(__GongEnumShape__000000_Default_ButtonType.GongEnumValueShapes, __GongEnumValueShape__000002_EDIT)
	__GongEnumShape__000000_Default_ButtonType.GongEnumValueShapes = append(__GongEnumShape__000000_Default_ButtonType.GongEnumValueShapes, __GongEnumValueShape__000003_REMOVE)
	__GongEnumShape__000000_Default_ButtonType.GongEnumValueShapes = append(__GongEnumShape__000000_Default_ButtonType.GongEnumValueShapes, __GongEnumValueShape__000004_RENAME_CANCEL)
	// setup of GongEnumValueShape instances pointers
}
