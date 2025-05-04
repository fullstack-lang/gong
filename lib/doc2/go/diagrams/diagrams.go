package diagrams

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

	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_04T22_31_26Z := (&models.DiagramPackage{}).Stage(stage)

	__GongStructShape__000000_Default_Classdiagram := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000001_Default_GongStructShape := (&models.GongStructShape{}).Stage(stage)

	__LinkShape__000000_GongStructShapes := (&models.LinkShape{}).Stage(stage)

	// Setup of values

	__Classdiagram__000000_Default.Name = `Default`
	__Classdiagram__000000_Default.IsInRenameMode = false
	__Classdiagram__000000_Default.IsExpanded = true
	__Classdiagram__000000_Default.NodeGongStructsIsExpanded = true
	__Classdiagram__000000_Default.NodeGongStructNodeExpansionBinaryEncoding = 34
	__Classdiagram__000000_Default.NodeGongEnumsIsExpanded = false
	__Classdiagram__000000_Default.NodeGongEnumNodeExpansionBinaryEncoding = 0
	__Classdiagram__000000_Default.NodeGongNotesIsExpanded = false
	__Classdiagram__000000_Default.NodeGongNoteNodeExpansionBinaryEncoding = 0

	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_04T22_31_26Z.Name = `Diagram Package created the 2025-05-04T22:31:26Z`
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_04T22_31_26Z.Path = ``
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_04T22_31_26Z.GongModelPath = ``
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_04T22_31_26Z.AbsolutePathToDiagramPackage = ``

	__GongStructShape__000000_Default_Classdiagram.Name = `Default-Classdiagram`
	__GongStructShape__000000_Default_Classdiagram.X = 40.000000
	__GongStructShape__000000_Default_Classdiagram.Y = 72.000000

	//gong:ident [ref_models.Classdiagram] comment added to overcome the problem with the comment map association
	__GongStructShape__000000_Default_Classdiagram.Identifier = `ref_models.Classdiagram`
	__GongStructShape__000000_Default_Classdiagram.ShowNbInstances = false
	__GongStructShape__000000_Default_Classdiagram.NbInstances = 0
	__GongStructShape__000000_Default_Classdiagram.Width = 240.000000
	__GongStructShape__000000_Default_Classdiagram.Height = 63.000000
	__GongStructShape__000000_Default_Classdiagram.IsSelected = false

	__GongStructShape__000001_Default_GongStructShape.Name = `Default-GongStructShape`
	__GongStructShape__000001_Default_GongStructShape.X = 655.000000
	__GongStructShape__000001_Default_GongStructShape.Y = 68.000000

	//gong:ident [ref_models.GongStructShape] comment added to overcome the problem with the comment map association
	__GongStructShape__000001_Default_GongStructShape.Identifier = `ref_models.GongStructShape`
	__GongStructShape__000001_Default_GongStructShape.ShowNbInstances = false
	__GongStructShape__000001_Default_GongStructShape.NbInstances = 0
	__GongStructShape__000001_Default_GongStructShape.Width = 240.000000
	__GongStructShape__000001_Default_GongStructShape.Height = 63.000000
	__GongStructShape__000001_Default_GongStructShape.IsSelected = false

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
	__LinkShape__000000_GongStructShapes.X = 400.500000
	__LinkShape__000000_GongStructShapes.Y = 96.000000
	__LinkShape__000000_GongStructShapes.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000000_GongStructShapes.StartRatio = 0.500000
	__LinkShape__000000_GongStructShapes.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000000_GongStructShapes.EndRatio = 0.500000
	__LinkShape__000000_GongStructShapes.CornerOffsetRatio = 1.380000

	// Setup of pointers
	// setup of Classdiagram instances pointers
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000000_Default_Classdiagram)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000001_Default_GongStructShape)
	// setup of DiagramPackage instances pointers
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_04T22_31_26Z.Classdiagrams = append(__DiagramPackage__000000_Diagram_Package_created_the_2025_05_04T22_31_26Z.Classdiagrams, __Classdiagram__000000_Default)
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_04T22_31_26Z.SelectedClassdiagram = __Classdiagram__000000_Default
	// setup of GongStructShape instances pointers
	__GongStructShape__000000_Default_Classdiagram.LinkShapes = append(__GongStructShape__000000_Default_Classdiagram.LinkShapes, __LinkShape__000000_GongStructShapes)
	// setup of LinkShape instances pointers
}
