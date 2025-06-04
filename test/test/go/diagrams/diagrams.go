package diagrams

import (
	"time"

	"github.com/fullstack-lang/gong/lib/doc2/go/models"
	// injection point for ident package import declaration
	ref_models "github.com/fullstack-lang/gong/test/test/go/models"
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var _ time.Time

// _ point for meta package dummy declaration
var _ ref_models.Stage

// When parsed, those maps will help with the renaming process
var _ map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// function will stage objects
func _(stage *models.Stage) {

	// Declaration of instances to stage

	__Classdiagram__000000_Default := (&models.Classdiagram{}).Stage(stage)

	__DiagramPackage__000000_Diagram_Package_created_the_2025_06_03T03_09_51Z := (&models.DiagramPackage{}).Stage(stage)

	__GongEnumShape__000000_Default_AEnumType := (&models.GongEnumShape{}).Stage(stage)

	__GongEnumValueShape__000000_ENUM_VAL1 := (&models.GongEnumValueShape{}).Stage(stage)
	__GongEnumValueShape__000001_ENUM_VAL2 := (&models.GongEnumValueShape{}).Stage(stage)

	// Setup of values

	__Classdiagram__000000_Default.Name = `Default`
	__Classdiagram__000000_Default.Description = ``
	__Classdiagram__000000_Default.IsIncludedInStaticWebSite = false
	__Classdiagram__000000_Default.IsInRenameMode = false
	__Classdiagram__000000_Default.IsExpanded = true
	__Classdiagram__000000_Default.NodeGongStructsIsExpanded = false
	__Classdiagram__000000_Default.NodeGongStructNodeExpansion = ``
	__Classdiagram__000000_Default.NodeGongEnumsIsExpanded = true
	__Classdiagram__000000_Default.NodeGongEnumNodeExpansion = `[true]`
	__Classdiagram__000000_Default.NodeGongNotesIsExpanded = false
	__Classdiagram__000000_Default.NodeGongNoteNodeExpansion = ``

	__DiagramPackage__000000_Diagram_Package_created_the_2025_06_03T03_09_51Z.Name = `Diagram Package created the 2025-06-03T03:09:51Z`
	__DiagramPackage__000000_Diagram_Package_created_the_2025_06_03T03_09_51Z.Path = ``
	__DiagramPackage__000000_Diagram_Package_created_the_2025_06_03T03_09_51Z.GongModelPath = ``
	__DiagramPackage__000000_Diagram_Package_created_the_2025_06_03T03_09_51Z.AbsolutePathToDiagramPackage = ``

	__GongEnumShape__000000_Default_AEnumType.Name = `Default-AEnumType`
	__GongEnumShape__000000_Default_AEnumType.X = 88.000000
	__GongEnumShape__000000_Default_AEnumType.Y = 85.000000

	//gong:ident [ref_models.AEnumType] comment added to overcome the problem with the comment map association
	__GongEnumShape__000000_Default_AEnumType.Identifier = `ref_models.AEnumType`
	__GongEnumShape__000000_Default_AEnumType.IdentifierMeta = new(ref_models.AEnumType)
	__GongEnumShape__000000_Default_AEnumType.Width = 240.000000
	__GongEnumShape__000000_Default_AEnumType.Height = 103.000000
	__GongEnumShape__000000_Default_AEnumType.IsExpanded = false

	__GongEnumValueShape__000000_ENUM_VAL1.Name = `ENUM_VAL1`

	//gong:ident [ref_models.AEnumType.ENUM_VAL1] comment added to overcome the problem with the comment map association
	__GongEnumValueShape__000000_ENUM_VAL1.Identifier = `ref_models.AEnumType.ENUM_VAL1`
	__GongEnumValueShape__000000_ENUM_VAL1.IdentifierMeta = ref_models.ENUM_VAL1

	__GongEnumValueShape__000001_ENUM_VAL2.Name = `ENUM_VAL2`

	//gong:ident [ref_models.AEnumType.ENUM_VAL2] comment added to overcome the problem with the comment map association
	__GongEnumValueShape__000001_ENUM_VAL2.Identifier = `ref_models.AEnumType.ENUM_VAL2`
	__GongEnumValueShape__000001_ENUM_VAL2.IdentifierMeta = ref_models.ENUM_VAL2

	// Setup of pointers
	// setup of Classdiagram instances pointers
	__Classdiagram__000000_Default.GongEnumShapes = append(__Classdiagram__000000_Default.GongEnumShapes, __GongEnumShape__000000_Default_AEnumType)
	// setup of DiagramPackage instances pointers
	__DiagramPackage__000000_Diagram_Package_created_the_2025_06_03T03_09_51Z.Classdiagrams = append(__DiagramPackage__000000_Diagram_Package_created_the_2025_06_03T03_09_51Z.Classdiagrams, __Classdiagram__000000_Default)
	__DiagramPackage__000000_Diagram_Package_created_the_2025_06_03T03_09_51Z.SelectedClassdiagram = __Classdiagram__000000_Default
	// setup of GongEnumShape instances pointers
	__GongEnumShape__000000_Default_AEnumType.GongEnumValueShapes = append(__GongEnumShape__000000_Default_AEnumType.GongEnumValueShapes, __GongEnumValueShape__000000_ENUM_VAL1)
	__GongEnumShape__000000_Default_AEnumType.GongEnumValueShapes = append(__GongEnumShape__000000_Default_AEnumType.GongEnumValueShapes, __GongEnumValueShape__000001_ENUM_VAL2)
	// setup of GongEnumValueShape instances pointers
}
